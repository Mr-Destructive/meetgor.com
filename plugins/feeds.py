import datetime
from re import template
import shutil
import textwrap
from pathlib import Path, PosixPath
from typing import TYPE_CHECKING, Any, List, Optional, Union

from jinja2 import Environment, FileSystemLoader, Template, Undefined, exceptions

from markata import Markata, os
from markata.hookspec import hook_impl

if TYPE_CHECKING:
    from frontmatter import Post


class SilentUndefined(Undefined):
    def _fail_with_undefined_error(self, *args, **kwargs):
        return ""


class MarkataFilterError(RuntimeError):
    ...

template = (
    Path(__file__).resolve().parents[1] / "layouts" / "default_post_template.html"
)

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
    #template = markata.get_plugin_config("feeds").get("template") \
    #        or Path(__file__).resolve().parents[1] / "layouts" / "default_post_template.html"
    template = (
        Path(__file__).resolve().parents[1] / "layouts" / "default_post_template.html"
    )

    for page, page_conf in config.items():
        template = Path(__file__).resolve().parents[1] / "layouts" / "default_post_template.html"
        if page not in ["cache_expire", "config_key"]:
            if markata.get_plugin_config("feeds")[page].get("template"):
                template = Path(__file__).resolve().parents[1] / markata.get_plugin_config("feeds")[page].get("template")
            create_page(
                markata,
                page,
                description=description,
                url=url,
                template=template,
                **page_conf,
            )
            create_yearly_pages(
                markata,
                page,
                description=description,
                url=url,
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
    status="",
    template=template,
    card_template=None,
    filter=None,
    description=None,
    url=None,
    today=datetime.datetime.today(),
    title="Techstructive Blog",
    year=None,
    posts=None,
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

    with open(template) as f:
        template = Template(f.read())

    #if template != Path(__file__).resolve().parents[1] / "layouts" / "default_post_template.html":
    #    template = str(Path(__file__).resolve().parents[1] / markata.get_plugin_config("feeds")[page].get("template"))
    """
    with open(template) as f:
        env = Environment()
        env.loader = FileSystemLoader("markata/plugins/")
        if type(template) is PosixPath:
            template = template.name
            if not Path(template).is_file():
                template = Template(f.read(), undefined=SilentUndefined)
            else:
                template = str(Path(__file__).resolve().parents[1] / template)
        template = env.get_template(template)
    """

    if year:
        output_file = Path(markata.config["output_dir"]) / "blog" / str(year) /  "index.html"
        archive_file = Path(markata.config["output_dir"]) / "archive" / str(year) / "index.html"
    else:
        output_file = Path(markata.config["output_dir"]) / "blog" / "index.html"
        archive_file = Path(markata.config["output_dir"]) / "archive" / "index.html"
    canonical_url = f"{url}/{page}/"
    output_file.parent.mkdir(exist_ok=True, parents=True)
    archive_file.parent.mkdir(exist_ok=True, parents=True)

    if year and page in ["archive", "blog"]:
        years = give_years([])
    else:
        years = give_years(posts)

    with open(output_file, "w+") as f:
        f.write(
            template.render(
                body="".join(cards),
                count=count,
                url=url,
                description=description,
                title=title,
                years=years,
                canonical_url=canonical_url,
                today=datetime.datetime.today(),
            )
        )
    #cards.insert(0, "<ul>")
    #cards.append("</ul>")
    #for y in years.keys():
    #    cards.append(f"<li><a href='/blog/{y}/'>{y}</a></li>")

    with open(archive_file, "w+") as f:
        f.write(
            template.render(
                body="".join(cards),
                count=count,
                url=url,
                years=years,
                description=description,
                title=title,
                canonical_url=canonical_url,
                today=datetime.datetime.today(),
            )
        )

def give_years(posts):
    years = {}
    if not posts:
        markata = Markata()
        posts = [post for post in markata.articles if post["status"] == "published"]
    for post in posts:
        year = post["date"].year
        if year not in years:
            years[year] = []
        years[year].append(post)
    return sorted(years.keys())

def create_yearly_pages(
    markata,
    page,
    description=None,
    url=None,
    today=datetime.datetime.today(),
    title="Techstructive Blog",
    filter=None,
    year=None,
) -> None:
    # Group posts by year
    posts_by_year = {}
    for post in markata.articles:
        year = post["date"].year
        if year not in posts_by_year:
            posts_by_year[year] = []
        posts_by_year[year].append(post)

    def try_filter_date(x):
        try:
            return x["date"]
        except KeyError:
            return -1

    if filter is not None:
        posts_by_year = {
            year: reversed(sorted(posts, key=try_filter_date))
            for year, posts in posts_by_year.items()
        }
        try:
            posts_by_year = {
                year: [post for post in posts if eval(filter, post.to_dict(), {})]
                for year, posts in posts_by_year.items()
            }
        except BaseException as e:
            msg = textwrap.dedent(
                f"""
                    While processing page='{page}' markata hit the following exception
                    during filter='{filter}'
                    {e}
                    """
            )
            raise MarkataFilterError(msg)
    # Create a page for each year
    for year, posts in posts_by_year.items():
        create_page(
            markata, page, posts=posts, 
            description=description, url=url, title=title, year=year
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
