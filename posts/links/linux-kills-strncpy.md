---
title: "Linux kills strncpy"
date: 2026-06-26
link: "https://smist08.wordpress.com/2026/06/25/linux-kills-strncpy/"
status: published
image_url: "https://i0.wp.com/smist08.wordpress.com/wp-content/uploads/2026/06/image-2.png?fit=1200%2C800&ssl=1"
source: newsletter
newsletter: techstructive-weekly-100
type: links
slug: linux-kills-strncpy
tags:
description: "Introduction The C string library is compact, fast and efficient. However, if not used correctly and carefully, it leads to buffer overrun errors which either cause programs to crash or worse allow…"
hash: f9e97b25bfa73b0949613e5226f7e62e59ba92ae4912acbb25a65667432f9558
---
My thoughts on [Linux kills strncpy](https://smist08.wordpress.com/2026/06/25/linux-kills-strncpy/): Linux kills strncpy

## Commentary

- Stephen Smith’s Blog: Linux kills strncpy
- Oh! That is nasty. I didn’t knew strncpy was that dangerous. Copying data from the stack was a gigachad move from hackers. being a hacker is so underrated, like it takes some thinking to break code, though the direction is wrong so its not worth it.
- But the other point is that it took 5-6 years to remove this from the kernel. This shows how hard it could get if some code remains in the system for a long time. This is legacy code, if we are vibe coding things, ai generated code will at some point become this, and the problem is not that it becomes hard to remove things, but the knowledge of where the code should be and should not be removed is also gone. Engineers might not know the code anymore, so it remains on the whim of AIs to generate and clean up things. Adding the complexity.
