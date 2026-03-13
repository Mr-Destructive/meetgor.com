# Techstructive Blog

- [Blog](https://meetgor.com/)
- Author: [Meet Gor](https://github.com/mr-destructive)

I write about the things I learn here. Some of the topics for most of the Articles/Blogs are:

- [Django](https://meetgor.com/tag/django)
- [Golang](https://meetgor.com/tag/go)
- [Python](https://meetgor.com/tag/python)
- [Vim](https://meetgor.com/tag/vim)
- And Programming/Development in general

This is a personal blog made with html/css/markdown, generated using Burrow(a SSG written in Golang) and hosted on GitHub pages.
- GitHub Pages for hosting
- Turso Database for storing/syncing posts
- Netlify Cloud Functions for serverless functions (adding and syncing posts, triggering builds)

## Content + CMS Sync

- Run `scripts/sync_from_author.sh` to sync latest markdown content from `/home/meet/code/blog/author` into this repo's `posts/` directory.
- CMS endpoint is available at `/cms` (served by `netlify/functions/cms`) and uses existing `api`, `sync-db`, and `trigger-build` functions.
- Run `python3 scripts/sync_newsletter_feed.py` to pull the latest Techstructive Weekly issues from Substack into `posts/newsletter/`.

## Turso Post Sync

Script:

```bash
./scripts/sync_posts.sh --sync-all --verbose
```

Build-time plugin (optional). The plugin only runs if `SYNC_DB_ON_BUILD=1`:

```bash
export SYNC_DB_ON_BUILD=1
export TURSO_DATABASE_NAME="libsql://..."
export TURSO_DATABASE_AUTH_TOKEN="..."
```

Optional plugin flags: `SYNC_DB_SYNC_ALL=1`, `SYNC_DB_VERBOSE=1`, `SYNC_DB_DRY_RUN=1`, `SYNC_DB_HOURS=6`.

### Techstructive Blog is created with a aim to learn programming and have fun in the process.

You'll also find some of my articles written at [GeeksforGeeks](https://auth.geeksforgeeks.org/user/meetgor/articles).
They mainly contain shell scripting(bash), Java, Vim and some random articles as well.

I hope you'll have fun reading my articles, if not please give feedback on the handles from [Contact](https://meetgor.com/blog/) section.
