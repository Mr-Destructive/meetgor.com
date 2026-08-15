---
title: "How chat compaction works in Pi Agent"
date: 2026-08-14
link: "https://earendil.com/posts/compaction-in-pi/"
status: published
image_url: "https://earendil.com/static/og/posts/compaction-in-pi.png"
source: newsletter
newsletter: techstructive-weekly-107
type: links
slug: how-chat-compaction-works-in-pi-agent
tags:
description: "Why compaction is needed for large language models and how Pi implements it"
hash: fdc17e382130a2f4abe812a1c9175eacd025d21e9f964949d775287a5510cd9f
---
My thoughts on [How chat compaction works in Pi Agent](https://earendil.com/posts/compaction-in-pi/): How chat compaction works in Pi Agent

## Commentary

- How chat compaction works in Pi Agent
- This is interesting. I thought it might be some deterministic call. But it uses an LLM to summarize and hand-off the current conversation to other in order to make it coherent.
- I now understand that claude might be using haiku or some older sonnet models to compact conversations maybe. Also cursor might do something really smart since they know the bread and butter of chat for ai agents.
