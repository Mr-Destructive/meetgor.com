---
date: 2026-01-03T00:00:00Z
description: 'CMU Database System #4 Memory Management and Buffer pools'
image_url: https://i.ytimg.com/vi/8-2yv4z0VZc/hqdefault.jpg
link: https://www.youtube.com/watch?v=8-2yv4z0VZc
newsletter: 2026-01-03-techstructive-weekly-75
slug: cmu-database-system-4-memory-management-and-buffer-pools
hash: c80bc1c91838818655c2c9fdec9a9e7afbcd2b16b8698f4b03ef975b14fd3f16
source: newsletter
status: published
tags: 
title: 'CMU Database System #4 Memory Management and Buffer pools'
type: links
---
## Commentary

- Ok OS is not our friend, we need to manage our memory ourselves. This went wild, I thought managing memory was like shooting yourselves on the foot, but not for DBs.
- So we load the database file, from the disk into memory not as full, but chunks called frames, where each page is contained in the buffer pool. Interesting, this is done in the actual ram or the memory not full at once.
- So this makes it the different algorithms to decide which frames/pages to keep and evict (remove)
- There is a difference in lock and latches, a lock is something that protect the database logical content from other transactions i.e. the data to write or avoid corrupted reading
- However a latch is something that helps in preventing the database internals from other operations, its only for an operation not a query. Its like a mutex.
- We can’t rely on OS, as OS doesn’t know what are we querying.
- There are like half a dozen implementation of replacement caches like LRU, Clock, LFU, LRU-K, ARC, etc.