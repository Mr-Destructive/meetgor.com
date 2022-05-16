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
    config = markata.get_plugin_config("series_index")
    if config is None:
        config["series_index"] = dict()
    if "series" not in config.keys():
        config["series_index"] = dict()
        config["series_index"]["filter"] = "templateKey in ['blog-post',] and status.lower()=='published'"

    articles = [article for article in markata.iter_articles("series")]

    series_list = [post for post in articles]
    series_list = [post for post in series_list if 'series' in post]
    series_list = [post['series'] for post in series_list]
    series_list = list(set(series_list))

    for series in series_list:
        description = markata.get_config("description") or ""
        url = markata.get_config("url") or ""
        template = Path(__file__).resolve().parents[1] / "layouts" / "series_page.html"

        for page, page_conf in config.items():
            if page not in ["cache_expire", "config_key"]:
                create_page(
                    markata,
                    page,
                    description=description,
                    url=url,
                    series=series,
                    template=template,
                    **page_conf,
                )

        home = Path(markata.config["output_dir"]) / "index.html"
        series_name = series.replace(" ", "-").lower()
        series_page = Path(markata.config["output_dir"]) / "series" / series_name / "index.html"
        if not home.exists() and series_page.exists():
            shutil.copy(str(series_page), str(home))


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
    series=None,
    today=datetime.datetime.today(),
    title="Techstructive Blog",
):
    def try_filter_date(x):
        try:
            return x["date"]
        except KeyError:
            return -1

    if filter is not None:
        posts = sorted(markata.articles, key=try_filter_date)
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

    posts = [post for post in posts if 'series' in post]
    posts = [post for post in posts if post['series'] == series]
    count = len(posts)

    cards = [create_card(post, card_template) for post in posts]
    cards.insert(0, "<ul>")
    cards.append("</ul>")

    with open(template) as f:
        template = Template(f.read())
    series_name = series.replace(" ", "-").lower()
    series_description = series
    if "series_description" in posts[0]:
        series_description = posts[0]["series_description"]
    output_file = Path(markata.config["output_dir"]) / "series" / series_name / "index.html"
    canonical_url = f"/{url}/{page}/"
    output_file.parent.mkdir(exist_ok=True, parents=True)

    with open(output_file, "w+") as f:
        f.write(
            template.render(
                body="".join(cards),
                series=series,
                series_description=series_description,
                count=count,
                url=url,
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
