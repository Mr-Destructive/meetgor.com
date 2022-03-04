from typing import Union, List
from markata.hookspec import hook_impl
from string import Template
from pathlib import Path
from more_itertools import unique_everseen
import datetime
import shutil
import textwrap
from pathlib import Path

from jinja2 import Template

from markata.hookspec import hook_impl


class MarkataFilterError(RuntimeError):
    ...

@hook_impl
def pre_render(markata):
    config = markata.get_plugin_config("tags")
    if config is None:
        config["tags"] = dict()
    if "tags" not in config.keys():
        config["tags"] = dict()
        config["tags"]["filter"] = "templateKey == 'blog-post'"

    description = markata.get_config("description") or ""
    url = markata.get_config("url") or ""

    articles = [article for article in markata.iter_articles("tags")]

    tags = [ tag for post in articles for tag in post['tags']]
    tags = list(set(tags))

    for tag in tags:
        description = markata.get_config("description") or ""
        url = markata.get_config("url") or ""
        template = Path(__file__).parent / "tags_template.html"

        for page, page_conf in config.items():
            if page not in ['cache_expire', 'config_key']:
                create_page(
                    markata,
                    page,
                    tag=tag,
                    description=description,
                    url=url,
                    template=template,
                )

        
        home = Path(markata.config["output_dir"]) / "index.html"
        archive = Path(markata.config["output_dir"]) / tag / "index.html"
        if not archive.exists():
            shutil.copy(str(archive), str(home))


def create_page(
    markata,
    page,
    tag,
    status="published",
    template=None,
    card_template=None,
    filter=None,
    description=None,
    url=None,
    today=datetime.datetime.today(),
    title="Techstructive Blog",
):

    def try_filter_date(x):
        try:
            return x["date"]
        except KeyError:
            return -1

    posts = [post for post in markata.iter_articles("tags") if tag in post['tags']]

    cards = [create_card(post, card_template) for post in posts]
    cards.insert(0, "<ul>")
    cards.append("</ul>")

    with open(template) as f:
        template = Template(f.read())
    output_file = Path(markata.config["output_dir"])/ tag / "index.html"
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
                canonical_url=canonical_url,
                today=datetime.datetime.today(),
            )
        )


def create_card(post, template=None):
    if template is None:
        if "date" and "image_url" in post.keys():
            return textwrap.dedent(
                f"""
                <li class='post'>
                <img src="{post['image_url']}" class="cover-image" >
                <a href="/techstructive-blog/{post['slug']}/">
                   <h2 id="title"> {post['title']} </h2>
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
