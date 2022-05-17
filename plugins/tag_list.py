import datetime
import shutil
import textwrap
from pathlib import Path

from jinja2 import Template

from markata.hookspec import hook_impl


class MarkataFilterError(RuntimeError):
    ...


@hook_impl
def save(markata):
    config = markata.get_plugin_config("tag_list")
    if config is None:
        config["tags"] = dict()
    if "tags" not in config.keys():
        config["tags"] = dict()
        config["tags"]["filter"] = "templateKey in ['blog-post',] and status.lower()=='published'"

    description = markata.get_config("description") or ""
    url = markata.get_config("url") or ""
    template = Path(__file__).resolve().parents[1] / "layouts"/ "tags_page.html"

    articles = [article for article in markata.iter_articles("tags")]

    tags = [ tag for post in articles for tag in post['tags']]
    tags = list(set(tags))

    for tag in tags:

        for page, page_conf in config.items():
            if page not in ["cache_expire", "config_key"]:
                create_page(
                    markata,
                    page,
                    description=description,
                    url=url,
                    template=template,
                    **page_conf,
                )

        home = Path(markata.config["output_dir"]) / "index.html"
        tags_page = Path(markata.config["output_dir"]) / "tags" / "index.html"
        if not home.exists() and tags_page.exists():
            shutil.copy(str(tags_page), str(home))


def create_page(
    markata,
    page,
    tags=None,
    status="published",
    template=None,
    card_template=None,
    filter=None,
    description=None,
    url=None,
    tag=None,
    today=datetime.datetime.today(),
    title="Techstructive Blog",
):
    def try_filter_date(x):
        try:
            return x["date"]
        except KeyError:
            return -1

    if filter is not None:
        posts = reversed(sorted(markata.articles, key=try_filter_date))
        try:
            posts = [post for post in posts if eval(filter, post.to_dict(), {})]
        except BaseException as e:
            msg = textwrap.dedent(
                f"""
                    While processing page='{page}' markata hit the following exception
                    during filter='{filter}'
                    {e}
                    """
            )
            raise MarkataFilterError(msg)

    articles = [article for article in markata.iter_articles("tags")]
    tags = [ tag for post in articles for tag in post['tags']]
    tags = list(set(tags))

    count = len(posts)
    cards = [create_card(tag, card_template) for tag in tags]
    cards.insert(0, "<ul class='taglist'>")
    cards.append("</ul>")

    with open(template) as f:
        template = Template(f.read())
    output_file = Path(markata.config["output_dir"]) / page / "index.html"
    canonical_url = f"/{page}/"
    output_file.parent.mkdir(exist_ok=True, parents=True)

    with open(output_file, "w+") as f:
        f.write(
            template.render(
                body="".join(cards),
                count=count,
                url=url,
                description=description,
                title=title,
                canonical_url=canonical_url,
                today=datetime.datetime.today(),
            )
        )


def create_card(tag, template=None):
    if template is None:
        tag_name = tag.replace(" ","-").lower()
        return textwrap.dedent(
            f"""
            <li class='post'>
            <a href="/techstructive-blog/tag/{tag_name}/">
               <h2 id="title">#{tag}</h2>
            </a>
            </li>
            """
        )
    try:
        with open(template) as f:
            template = Template(f.read())
    except FileNotFoundError:
        template = Template(template)

