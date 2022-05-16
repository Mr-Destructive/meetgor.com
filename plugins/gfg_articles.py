import bs4
from urllib.request import urlopen
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
        config["gfg_articles"] = dict()
    if "gfg_articles" not in config.keys():
        config["gfg_articles"] = dict()
        config["gfg_articles"]["filter"] = "True"
    template = Path(__file__).resolve().parents[1] / "layouts" / "gfg_template.html"

    description = markata.get_config("description") or ""

    for page, page_conf in config.items():
        if page not in ["cache_expire", "config_key"]:
            create_page(
                markata,
                page,
                description=description,
                template=template,
                **page_conf,
            )

    home = Path(markata.config["output_dir"]) / "index.html"
    archive = Path(markata.config["output_dir"]) / "geeksforgeeks" / "index.html"
    if not home.exists() and archive.exists():
        shutil.copy(str(archive), str(home))

def get_posts():
    url = "https://auth.geeksforgeeks.org/user/meetgor/articles"
    html_page = urlopen(url)
    soup = bs4.BeautifulSoup(html_page, features="lxml")
    li = list(soup.select(".contribute-ol li a"))
    return li

def create_page(
    markata,
    page,
    tags=None,
    description=None,
    status="published-gfg",
    template=None,
    card_template=None,
    filter=None,
    url=None,
    today=datetime.datetime.today(),
    title="GeeksforGeeks Articles",
):

    def try_filter_date(x):
        try:
            return x["date"]
        except KeyError:
            return -1

    posts = get_posts()[:-7][::-1]
    count = len(posts)
    cards = [create_card(post, card_template) for post in posts]
    cards.insert(0, "<ul>")
    cards.append("</ul>")

    with open(template) as f:
        template = Template(f.read())
    output_file = Path(markata.config["output_dir"])/ "gfg" / "index.html"
    output_file.parent.mkdir(exist_ok=True, parents=True)
    canonical_url = f"/{url}/gfg/"

    with open(output_file, "w+", encoding="utf-8") as f:
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
        return textwrap.dedent(
            f"""
            <li class='post'>
            <a href="{post.get('href')}">
                <h3 id="title">{post.text}</h3>
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
