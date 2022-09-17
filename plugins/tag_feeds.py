from markata.hookspec import hook_impl
from string import Template
from pathlib import Path
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
    config = markata.get_plugin_config("tag_feeds")
    if config is None:
        config["tags"] = dict()
        config["tags"]["filter"] = "templateKey in ['blog-post',] and status.lower()=='published'"
    tags = config.get("feed_for_each", None)
    if type(tags) is not list:
        tags = [tags]

    description = markata.get_config("description") or ""
    url = markata.get_config("url") or ""
    template = Path(__file__).resolve().parent / "default_post_template.html"

    for tag in tags:

        for page, page_conf in config.items():
            if page not in ["cache_expire", "config_key",]:
                create_page(
                    markata,
                    page,
                    description=description,
                    tag=tag,
                    url=url,
                    template=template,
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

    posts = [post for post in markata.articles if tag in post.to_dict()['tags']] 
    if filter is not None:
        posts = reversed(sorted(markata.articles, key=try_filter_date))
        try:
            posts = [post for post in posts if eval(filter, post.to_dict(), {}) if tag in post.tags]
        except BaseException as e:
            msg = textwrap.dedent(
                f"""
                    While processing page='{page}' markata hit the following exception
                    during filter='{filter}'
                    {e}
                    """
            )
            raise MarkataFilterError(msg)

    count = len(posts)
    cards = [create_card(post, card_template) for post in posts]
    cards.insert(0, "<ul class='taglist'>")
    cards.append("</ul>")

    with open(template) as f:
        template = Template(f.read())
    output_file = Path(markata.config["output_dir"]) / "tag" / tag / "index.html"
    canonical_url = f"/{url}/{tag}/"
    output_file.parent.mkdir(exist_ok=True, parents=True)

    with open(output_file, "w+") as f:
        f.write(
            template.render(
                body="".join(cards),
                url=url,
                description=description,
                title=title,
                tag=tag,
                count=count,
                canonical_url=canonical_url,
                today=datetime.datetime.today(),
            )
        )


def create_card(post, tag, template=None):
    if template is None:
        if "date" and "image_url" in post.keys():
            return textwrap.dedent(
                f"""
                <li class='post'>
                <img src="{post['image_url']}" class="cover-image" >
                <a href="/{post['slug']}/">
                   <h2 id="title"> {post['title']} </h2>
                </a>
                </li>
                """
            )
        else:
            return textwrap.dedent(
                f"""
                <li class='post'>
                <a href="/{post['slug']}/">
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
