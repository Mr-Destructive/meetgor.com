---
slug: robbie-coomber-wrote-a-70x-faster-sql-parser-while-barely-looking-at-the-code-agentic-coding-and-not-vibe-coding
hash: 392d77245b289034db71f9ce77c575a22b0794c65506589ccf50676e77a108a5
title: "Robbie Coomber wrote a 70x faster SQL parser while barely looking at the code (agentic coding and not vibe coding)"
status: published
source: newsletter
tags: 
description: "After the success of using agents to  improve query performance through autoresearch , I wanted to try something more ambitious. I rewrote PostHog's…"
date: 2026-06-26
link: "https://posthog.com/blog/sql-parser"
image_url: "https://d36j3rcgc2qfsv.cloudfront.net/blogsql-parser.jpeg"
newsletter: techstructive-weekly-100
type: links
---
My thoughts on [Robbie Coomber wrote a 70x faster SQL parser while barely looking at the code (agentic coding and not vibe coding)](https://posthog.com/blog/sql-parser): Robbie Coomber wrote a 70x faster SQL parser while barely looking at the code (agentic coding and not vibe coding)

## Commentary

- PostHog’s Blog: Robbie Coomber wrote a 70x faster SQL parser while barely looking at the code (agentic coding and not vibe coding)
- OK, hot takes.This is a great article. It proves that engineering is still a need to get most out of ai agents. It wouldn’t have been possible without the author knowing how the sql parser works in existing library ANTLR and a other approach to do it efficiently.
- Technically I dont know how each of the two approaches work, but from reading it meant two approaches are opposite ends, one is faster and only for specific set of general operations, the other is more careful and gives a general purpose full parser maybe. Both the approaches worked for him, but he choose the faster specific use case one. This shows how easy it is to generate a working code, but maintaining becomes a question. I case of parsing a SQL query, this looks simple and without any maintenance cost I think. But in general cases, AI generated code is not magic and will need some before or after thought.