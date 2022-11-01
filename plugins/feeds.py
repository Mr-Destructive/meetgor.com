import datetime
import shutil
import textwrap
from pathlib import Path, PosixPath
from typing import TYPE_CHECKING, Any, List, Optional, Union

from jinja2 import Environment, FileSystemLoader, Template, Undefined

from markata import Markata
from markata.hookspec import hook_impl

if TYPE_CHECKING:
    from frontmatter import Post


class SilentUndefined(Undefined):
    def _fail_with_undefined_error(self, *args, **kwargs):
        return ""


class MarkataFilterError(RuntimeError):
    ...


@hook_impl
def save(markata: Markata) -> None:
    """
    Creates a new feed page for each page in the config.
    """
    config = markata.get_plugin_config("feeds")
    if config is None:
        config["feeds"] = dict()
    if "archive" not in config.keys():
        config["archive"] = dict()
        config["archive"][
            "filter"
        ] = "templateKey in ['blog-post',] and status.lower()=='published'"

    description = markata.get_config("description") or ""
    url = markata.get_config("url") or ""
    template = (
        Path(__file__).resolve().parents[1] / "layouts" / "default_post_template.html"
    )

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
    archive = Path(markata.config["output_dir"]) / "blog" / "index.html"
    if not home.exists() and archive.exists():
        shutil.copy(str(archive), str(home))


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
    today=datetime.datetime.today(),
    title="Techstructive Blog",
) -> None:
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

    count = len(posts)
    cards = [create_card(post, card_template) for post in posts]
    cards.insert(0, "<ul class='post-list'>")
    cards.append("</ul>")

    #with open(template) as f:
    #    template = Template(f.read())


    with open(template) as f:
        env = Environment()
        env.loader = FileSystemLoader("markata/plugins/")
        if type(template) is PosixPath:
            template = template.name
            if not Path(template).is_file():
                template = Template(f.read(), undefined=SilentUndefined)
        print(template)
        template = env.get_template(template)

    output_file = Path(markata.config["output_dir"]) / "blog" / "index.html"
    archive_file = Path(markata.config["output_dir"]) / "archive" / "index.html"
    canonical_url = f"{url}/{page}/"
    output_file.parent.mkdir(exist_ok=True, parents=True)
    archive_file.parent.mkdir(exist_ok=True, parents=True)

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

    with open(archive_file, "w+") as f:
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


def create_card(post, template=None):
    if template is None:
        if "date" and "image_url" in post.keys():
            return textwrap.dedent(
                f"""
                <li class='post'>
                <img loading="lazy" src="{post['image_url']}" class="cover-image" >
                <a href="/{post['slug']}/">
                   <h2 id="title"> {post['title']} </h2>
                   <span>{ post['date'].strftime('%d-%m-%Y')  }</span>
                </a>
                </li>
                """
            )
        else:
            return textwrap.dedent(
                f"""
                <li class='post'>
                <a href="/{post['slug']}/">
                   <h2 id="title"> {post['title']} </h2>
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
