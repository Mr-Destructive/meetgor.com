from pathlib import Path
from typing import TYPE_CHECKING
from slugify import slugify

from markata.hookspec import hook_impl

if TYPE_CHECKING:
    from markata import Markata


@hook_impl(tryfirst=True)
def pre_render(markata: "Markata") -> None:
    """
    Sets the article slug if one is not already set in the frontmatter.
    """
    for article in markata.iter_articles(description="creating slugs"):
        article["slug"] = slugify(article.get("slug", Path(article["path"]).stem))
