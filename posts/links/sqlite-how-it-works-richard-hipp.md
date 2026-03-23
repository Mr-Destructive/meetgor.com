---
link: https://youtu.be/ZSKLA81tBis
newsletter: 2025-08-02-techstructive-weekly-53
source: newsletter
status: published
tags: 
type: links
image_url: https://i.ytimg.com/vi/ZSKLA81tBis/hqdefault.jpg
slug: sqlite-how-it-works-richard-hipp
title: 'SQLite: How it works: Richard Hipp'
hash: 6a05f7fcb055e48c2f75ffbb8f6fe3973d3c7341bad4ec61445b117b9aa394d2
date: 2025-08-02T00:00:00Z
description: 'SQLite: How it works: Richard Hipp'
---
## Commentary

- What a banger of a presentation and talk. Explained so much, in depth, in such a short time. It helped me understand what SQLite actually is, it’s a parser + virtual machine to run the core part and basically the fopen function in C to actually perform the operation.
- One unique insight here is
- Reading 10 files content from disk is slower than reading those file contents from SQLite
- Why? Because the database file is opened once and the reading happens in that instance only, data is stored in pages (fragments of memory), so it’s just a matter of reading bytes at a specific order.
- But reading 10 different files on disk will make you use fopen 10 times, and that is slow!
- 200 IQ move from SQLite team, have never seen such a beautiful solution to almost all the problems in the data world.