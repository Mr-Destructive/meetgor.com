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
    config = markata.get_plugin_config("feeds")
    if config is None:
        config["tils"] = dict()
    if "tils" not in config.keys():
        config["tils"] = dict()
        config["tils"]["filter"] = "True"
        config["tils"]["filter"] = "templateKey == 'til'"

    description = markata.get_config("description") or ""
    url = markata.get_config("url") or ""

    template = Path(__file__).resolve().parents[1] / "layouts" / "tils_template.html"

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
    archive = Path(markata.config["output_dir"]) / "tils" / "index.html"
    if not home.exists() and archive.exists():
        shutil.copy(str(archive), str(home))


def create_page(
    markata,
    page,
    tags=None,
    status="published-til",
    template=None,
    card_template=None,
    filter=None,
    description=None,
    url=None,
    today=datetime.datetime.today(),
    title="Techstructive Tils",
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

    cards = [create_card(post, card_template) for post in posts]
    cards.insert(0, "<ul>")
    cards.append("</ul>")

    with open(template) as f:
        template = Template(f.read())
    output_file = Path(markata.config["output_dir"])/ page / "index.html"
    canonical_url = f"/{url}/{page}/"
    output_file.parent.mkdir(exist_ok=True, parents=True)

    with open(output_file, "w+") as f:
        f.write(
            template.render(
                body="".join(cards),
                url=url,
                description=description,
                title=title,
                canonical_url=canonical_url,
                today=datetime.datetime.today(),
            )
        )


def create_card(post, template=None):
    if template is None:
        if "date" in post.keys():
            return textwrap.dedent(
                f"""
                <li class='post'>
                <a href="/techstructive-blog/{post['slug']}/">
                    <h3 id="title">{post['title']} - {post['date'].year}-{post['date'].month}-{post['date'].day}</h3>
                </a>
                </li>
                """
            )
        else:
            return textwrap.dedent(
                f"""
                <li class='post'>
                <a href="/techstructive-blog/{post['slug']}/">
                    <h2 id="title">{post['title']}</h2>
                </a>
                </li>
                """
            )
    try:
        with open(template) as f:
            template = Template(f.read())
    except FileNotFoundError:
        template = Template(template)
    post["article_html"] = post.article_html

    return template.render(**post.to_dict())
