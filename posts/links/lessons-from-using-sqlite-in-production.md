---
newsletter: techstructive-weekly-89
type: links
slug: lessons-from-using-sqlite-in-production
tags: 
date: 2026-04-10
status: published
description: "We run a production Rails store on SQLite — not Postgres, not MySQL. A single file on a Docker volume. It works surprisingly well until two containers try to..."
hash: b1fd0bda58f01adef92ad38b49a20a8e2462544819df1547d7a649ebc8cdb942
title: "Lessons from using SQLite in production"
link: "https://ultrathink.art/blog/sqlite-in-production-lessons"
image_url: "https://ultrathink.art/assets/og-default-bb1b2d7c.png"
source: newsletter
---
My thoughts on [Lessons from using SQLite in production](https://ultrathink.art/blog/sqlite-in-production-lessons): Lessons from using SQLite in production

## Commentary

- Lessons from using SQLite in production
- Useful bit here. The WAL mode is absolutely a clutch, most people don’t know of, and just lament about it being a single-writer constrained DB. It’s a super-powerful and very versatile lightweight database.
- Folks at Turso are making it even better with read-write replicas and syncing, having a daemon for SQLite. The future might be apps full of SQLite.