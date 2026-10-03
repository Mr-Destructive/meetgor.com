---
title: "Postgres timestamp vs timestamptz"
date: 2026-10-02
link: "https://bookofrevenue.com/blog/6ab81e9a97a13f0001f7e4e1/postgres-at-time-zone-u-does-not-do-what-you-think-it-does"
status: published
image_url: "https://bookofrevenue.com/assets/images/bb5d5899f499a7ad369d587be86672c2-blog-og-image.jpg"
source: newsletter
newsletter: techstructive-weekly-113
type: links
slug: postgres-timestamp-vs-timestamptz
tags:
description: "And why you may need to repeat AT TIME ZONE 'UTC' twice.





SUMMARY










 * AT TIME ZONE 'UTC' converts the data type from timestamptz to timestamp (example).
 * Now you are inadvertently using the timestamp without time zone (aka timestamp) data type, which is markedly discouraged.
 * There are tons of footguns with timestamp.
 * For example, the equality of timestamp and timestamptz will always be false..... unless your Postgres' default timezone is UTC. Read more.
 * Adding a month wit"
hash: 5f1bd31326047770f32ef8e3379523df7aef4ea2c58cb3862e3fe4c0beec06c2
---
My thoughts on [Postgres timestamp vs timestamptz](https://bookofrevenue.com/blog/6ab81e9a97a13f0001f7e4e1/postgres-at-time-zone-u-does-not-do-what-you-think-it-does): Postgres timestamp vs timestamptz

## Commentary

- Postgres timestamp vs timestamptz
- Oh timezones and developers. Love the ever love-hate relationship. There is no love actually. Dealing with time zones as a developer especially in databases, is a classic brain fader.
