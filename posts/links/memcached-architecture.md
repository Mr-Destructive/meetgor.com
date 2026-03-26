---
title: 'Memcached Architecture'
tags: 
status: published
description: 'Memcached Architecture'
source: newsletter
newsletter: 2025-07-26-techstructive-weekly-52
type: links
date: 2025-07-26
slug: memcached-architecture
link: 'https://hnasr.substack.com/p/memcached-architecture'
hash: 65621340c0d1ee8b97addf2aeb087b38faa3095d9ffe4bda3093b42397e7d61e
---
## Commentary

- Simple and straightforward explanation of the memcache
- Simple in-memory key-value store with slab-based memory management to avoid fragmentation. It has a threaded architecture and per-slab LRU for efficient concurrency and eviction, Also client-managed sharding enables distributed caching without server communication.