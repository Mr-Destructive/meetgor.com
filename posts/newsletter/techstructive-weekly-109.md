---
title: "Techstructive Weekly #109"
date: 2026-09-04
slug: techstructive-weekly-109
type: newsletter
status: published
source: newsletter
canonical_url: https://techstructively.substack.com/p/techstructive-weekly-109
description: "Not much reading, but watching and learning a lot of things, working with AI to build and refactor systems, among the other things did in the week from 30th August to 5th September 2026"
tags: ["newsletter", "substack"]
---

## Week #109

A bit of back to work week. Enjoyed the process maybe. Not sure, still figuring out the agentic engineering environment. I am actually lucky to have peers and colleagues that respect each other and understand the current trend of agentic work. “You don’t have to know everything, but you have to own it” that’s a bit hard to digest pill, If you know what I mean. Generating code, but not knowing the details but knowing enough of it to steer the issues. I felt good, not comfortable but not stressed or disrespected the least. I understand not everyone might have the privilege to be in this ai-change-friendly environment.

While being excited for the week, I found a couple of use cases for playing with LLMs.

- There is no openrouter like app on android or ios right?
- There is a scope for a python package or api that gives you free llm routing (like any free llm interface)
- Something to make with the needle2 14MB model which is a function tool calling for android

Would be looking forward in the weekend to do some of these.

### Quote of the Week

> “Presents are made for the pleasure of the one who gives them, not for the merit of those who receive them”
> 
> [— Carlos Ruiz Zafón, The Shadow of the Wind](https://www.goodreads.com/work/quotes/3209783https://www.goodreads.com/quotes/133882-presents-are-made-for-the-pleasure-of-who-gives-them)

I am reading the shadow of the wind, and father of the main character gives his son a present, the son is not happy. But the father is. Presents truly are for the giver and the by product is the receiver is happy too. What a beautiful little conversation it was, but it shows much about life.

## Read

1. [AI written code is still your code, are you ok with that?](https://martiansoftware.com/articles/ai-written-code-is-still-yours)
   
   1. Maybe sometimes not always right? If we generate the code, its the responsibility of the developer to own it that’s the rule.
   2. I find it a bit frustrating, quite less from last year. However some frustration is still there, since I haven’t found the reason for why each word in the code exists. To find that it takes a lot of effort, if its python, it can be easy with a REPL but other languages, not gonna take easy route, even with python in large code-bases and especially with LLM integration, it becomes harder and harder to make changes reliably and tested.

Didn’t read much, was a bit restless and read some book and worked the rest half of the week.

## Watched

- [Sriniously: HTTP for backend engineers](https://youtu.be/a3C1DMswClQ)
  
  - It was a great explanation of the protocol and the actual details
  - Unlike the theoretical stuff, the speaker actually showed each concept with a live demo. The headers part was really great, the methods and request response schema was touched perfectly.

<!--THE END-->

- [Dennis Hirsch Coding LS in C](https://youtu.be/HqXr5uXOb5o)
  
  - This was cool to see the file path. I didn’t we can get that low level with so little code in C. I want to do this in Go, but should I.
  - The feeling for doing it despite of AI is cool, but is it worth it? Should I be doing it? I am so overwhelmed with so many stuff.

<!--THE END-->

- [Omarchy Quattro](https://youtu.be/F7fe9pa8OeE)
  
  - Wow! This is so cool. It feels so lowkey, anybody who had used or customised his linux distro knows its a bit mechanical to do all of it. And omarchy specifically this release is the agentic way of doing that, not a big deal. But the vision to make someone onboard into linux is the big deal.
  - The effort to add functionality and the customisability is flawless and effortless.

<!--THE END-->

- PlanetScale: Postgres and MySQL Indexes are completely different
  
  - This is cool to know.
  - Postgres creates a B Tree for each index with the actual key and the tuple id , which makes look up of other data very trivial
  - MySQL creates a B Tree with the actual key and the primary key to look up the data separately in the other B-Tree. I might be a little less efficient compared to Postgres, but I am not sure what the trade-off here is, a good question to answer for the weekend. I love databases.

## Learnt

- Was reading the Learning SQL book and found some bits of information elsewhere to learn
  
  - SQLite has 2k max columns limit per table
  - PostgreSQL has 1600 max columns
  - MySQL has 1017 columns
  - DuckDB has unlimited * (maybe a bit of very high number)
  - Oracle has 1000 columns

## Tech News

- [OpenAI releases GPT 6 Astra](https://blog.google/innovation-and-ai/models-and-research/gemini-models/3-8-flash-and-3-8-flash-cyber/)
- [Google drops Gemini 3.8 Flash](https://blog.google/innovation-and-ai/models-and-research/gemini-models/3-8-flash-and-3-8-flash-cyber/)
- [Anthropic drops Fable and Mythos 5.1](https://www.anthropic.com/claude-fable-and-mythos-5-1)

* * *

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/808) (#808 edition), and for software development/coding articles, join daily.dev.

That’s it from the 109th Edition of techstructive weekly. I hope you found it helpful, and relaxing. If not please drop any suggestions, feedback or discussion about certain things you want to in the comments or drop me a message on my [socials](https://www.meetgor.com/contact).

Thank you for reading,

Until next week.

Happy Coding :)
