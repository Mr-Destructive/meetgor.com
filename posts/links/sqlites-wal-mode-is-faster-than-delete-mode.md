---
date: 2025-07-26T00:00:00Z
description: SQLite’s WAL Mode is faster than DELETE Mode
image_url: https://i.ytimg.com/vi/qf0GqRz-c74/hqdefault.jpg
link: https://youtu.be/qf0GqRz-c74?si=HZ_1yav_DFOzyiOn
newsletter: 2025-07-26-techstructive-weekly-52
slug: sqlites-wal-mode-is-faster-than-delete-mode
source: newsletter
status: published
tags:
    - links
title: SQLite’s WAL Mode is faster than DELETE Mode
type: links
---


## Commentary

- This is so well explained, first showed everything what each one is and then the benchmark just makes everything clear.
- The WAL mode basically writes the changes in a separate file and merges to the original db file whenever required, hence there is no overhead when reading or writing multiple writers or readers.
- The delete mode is like a backup, a journal, it keeps pages of the data that are to be changed and after it is committed it deletes the file, that clearly looks slow.
