---
link: "https://earendil.com/posts/compaction-in-pi/"
status: published
newsletter: techstructive-weekly-107
slug: how-chat-compaction-works-in-pi-agent
tags: 
date: 2026-08-14
image_url: "https://earendil.com/static/og/posts/compaction-in-pi.png"
source: newsletter
type: links
description: "Why compaction is needed for large language models and how Pi implements it"
hash: ac80b0e606980ddfdcb72a330d74a7725ecd7b9dce938471e90a025a6726c459
title: "How chat compaction works in Pi Agent"
---
My thoughts on [How chat compaction works in Pi Agent](https://earendil.com/posts/compaction-in-pi/): How chat compaction works in Pi Agent

## Commentary

- How chat compaction works in Pi Agent
- This is interesting. I thought it might be some deterministic call. But it uses an LLM to summarize and hand-off the current conversation to other in order to make it coherent.
- I now understand that claude might be using haiku or some older sonnet models to compact conversations maybe. Also cursor might do something really smart since they know the bread and butter of chat for ai agents.