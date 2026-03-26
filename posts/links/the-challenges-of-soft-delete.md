---
description: 'The challenges of soft delete'
newsletter: 2026-01-24-techstructive-weekly-78
slug: the-challenges-of-soft-delete
tags: 
link: 'https://atlas9.dev/blog/soft-delete.html'
source: newsletter
hash: 864ea27757f1ade20100f110f17446f598fbe04d4b293abce87af182010a4055
type: links
title: 'The challenges of soft delete'
date: 2026-01-24
status: published
---
## Commentary

- Nice read. I had experienced it in my first internship. This problem of dead objects. Especially if you are using Django and Postgres. It looked easy to add a field of soft deletion. But the resulting queries could create bottlenecks.
- Since then I haven’t quite gotten the chance to explore this, this article showed me the different ways to implement the soft deletion.