---
title: "Faster python imports with lazy imports"
date: 2026-10-09
link: "https://lemire.me/blog/2026/10/09/faster-python-startup-with-lazy-imports/"
status: published
image_url: "https://lemire.me/blog/wp-content/uploads/2026/10/lazy-import-cover-1024x541.jpg"
source: newsletter
newsletter: techstructive-weekly-114
type: links
slug: faster-python-imports-with-lazy-imports
tags:
description: "When a Python program starts, it needs to load all its dependencies (import). With the upcoming version of Python, it is possible to use lazy imports instead. lazy import json lazy from decimal import Decimal In this instance, the name json is no longer the module, but a mere placeholder object. The actual import only … Continue reading Faster Python startup with lazy imports"
hash: a52b6144c804742cfc0ea3ff3acc17e51914e2d5f6c3cdb9a1d37c7ae4736813
---
My thoughts on [Faster python imports with lazy imports](https://lemire.me/blog/2026/10/09/faster-python-startup-with-lazy-imports/): Faster python imports with lazy imports

## Commentary

- Faster python imports with lazy imports
- This is cool, can gain a lot of performance by placing the imports smartly. I mean if you have internal packages and dependencies that take a lot of time, you can do the other thing while doing the dependent thing in parallel right?
- Really cool way to improve python projects since there is no other point of improving the time
