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
    config = markata.get_plugin_config("series")
    if config is None:
        config["series"] = dict()
    if "series" not in config.keys():
        config["series"] = dict()
        config["series"]["filter"] = "templateKey in ['blog-post',] and status.lower()=='published'"

    description = markata.get_config("description") or ""
    url = markata.get_config("url") or ""
    template = Path(__file__).resolve().parents[1] / "layouts"/ "series_template.html"

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
    series = Path(markata.config["output_dir"]) / "series" / "index.html"
    if not home.exists() and series.exists():
        shutil.copy(str(series), str(home))


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

    posts = [post for post in posts if 'series' in post]
    series_list = []
    for post in posts:
        series_list.append(post['series'])
    series_list = list(set(series_list))
    
    count = len(posts)
    cards = [create_card(series, card_template) for series in series_list]
    cards.insert(0, "<ul>")
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


def create_card(series, template=None):
    if template is None:
        series_name = series.replace(" ","-").lower()
        return textwrap.dedent(
            f"""
            <li class='post'>
            <a href="/techstructive-blog/series/{series_name}/">
               <h2 id="title"> {series} </h2>
            </a>
            </li>
            """
        )
    try:
        with open(template) as f:
            template = Template(f.read())
    except FileNotFoundError:
        template = Template(template)

