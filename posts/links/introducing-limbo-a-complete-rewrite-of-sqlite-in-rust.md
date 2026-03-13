---

type: links
title: 'Introducing Limbo: A complete rewrite of SQLite in Rust'
date: 2024-12-14
slug: introducing-limbo-a-complete-rewrite-of-sqlite-in-rust
tags:
  - links
link: 'https://github.com/rusqlite/rusqlite'
status: published
description: 'Introducing Limbo: A complete rewrite of SQLite in Rust'
image_url: 'https://opengraph.githubassets.com/9675b6b5c55ab235338980e37a4216b3296e5b41b3ac3a0a9267b058b9f100ee/rusqlite/rusqlite'
source: newsletter
newsletter: 2024-12-14-techstructive-weekly-20
---


## Commentary

- This is pretty cool, I was thinking it might be already , yes definitely, but they are just not re-writing it, they are forking and adding features on top of it which is absolutely wild. Surely there won’t be time gains, but they now have a lot more control over what needs to be changed and included while adding more or even upstreaming from SQLite-core. People might call re-writing a waste of time (especially for such a well-developed and stable tool), but people forget they are making something from scratch gives you a whole different depth of understanding than just forking it. It will pay dividends slowly in the long run, pulling from upstream might be challenging and tedious but since SQLite is rock-solid, there won’t be any breaking changes that might get added to it, so a win-win for Turso in my opinion.
