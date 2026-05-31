---
title: "How soon is now in PostgreSQL?"
link: "https://www.architecture-weekly.com/p/how-soon-is-now-in-postgresql"
source: newsletter
newsletter: techstructive-weekly-96
date: 2026-05-29
status: published
image_url: "https://substackcdn.com/image/fetch/$s_!4LmN!,w_1200,h_675,c_fill,f_jpg,q_auto:good,fl_progressive:steep,g_auto/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Feabeb60e-b28d-46a8-b1bd-c88c5c811cb8_800x436.png"
type: links
slug: how-soon-is-now-in-postgresql
tags: 
description: "On troubleshooting nasty bug in PostgreSQL and on good tests vs proper tests"
hash: cf7e617191b4df61d93d718d68b25b22f1816a1b63338857417953e1eb4abcab
---
My thoughts on [How soon is now in PostgreSQL?](https://www.architecture-weekly.com/p/how-soon-is-now-in-postgresql): How soon is now in PostgreSQL?

## Commentary

- How soon is now in PostgreSQL?
- Great article. A good thing to know that now doesn’t take a new time in a transaction, it keeps the same timestamp the better way to use a current time in the transaction is clock_timestamp() , that makes it uniquely timestamp each call inside a transaction.
- Neat little thing to know before you shoot yourself on the foot.