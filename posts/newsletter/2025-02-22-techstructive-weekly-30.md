---
type: "newsletter"
title: "Techstructive Weekly #30"
date: 2025-02-22
---

## Week #30

What a week captain! Realized this on Wednesday truly. A bit of exhausting yet exciting week, it was a wild ride, true but learnt a lot. A lot of mistakes popping up at once.

I had to scamper and get things done, not just do it, but do it correct, not only do it correct , but also do it more, that was exaggerated, but yes, it was a bit taxing at some point, but the reward? I don’t care, I did my best (to some extent) and I am leaving the result to nature, really can’t control that now. A lot of intense debugging sessions, pair programming, side scripting, a lot of fun conversations and moments, it was worth it.

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.

I again didn’t stream and upload any videos or articles this week, and probably won’t do this weekend either. A tough month February 2025.

### Quote of the week

> **"It’s not whether you get knocked down, it’s whether you get up."**
> 
> — Vince Lombardi

Yes, this is how I feel right now. Knock me down, you’ll see me again standing up, struggling, yes, quiting, nope, never. I mean I am not saying I don’t give up, but I know when to give up and when to press on. I know which battles to fight and which to leave, it’s not worth it sometimes.

## Read

- [Why Blog if nobody reads it](https://news.ycombinator.com/item?id=42992159#:~:text=Why%20blog%20if%20nobody%20reads,Hacker%20News&text=Writing%20on%20a%20blog%20is,ve%20written%20in%20the%20past.)?
  This is a good post, why write articles when no one reads them. According to spirituality or philosophically, do work without caring for the result. But is writing articles is work? or hobby? Either of them, just do it because you have or need to or just so. Really there is no answer, will this universe be affected if you write or not write that post? Probably not, will it impact you, probably yes in a good way, will drain some time but I can guarantee after 3 years of blogging, you will have a better understanding (at least something) from what you had before writing. Just write it.

## Watched

- Nothing

A bit of hectic week as said, not much time for

## Learnt

- If you want to send a pure json response from API, and if for some reason, you code deals with numpy or dataframe in that case, then you might get into problems if not handled correctly. If you don’t convert or cast the lower level values (int64, char, etc used by numpy low level C) into python data types then you might get errors while decoding back to json. So make sure you cast back to higher level types. Had to debug for 3 hours to figure out where the heck the issue is coming from.
- What is the better approach to find a sub-string match with a list of keywords.
  So, let’s say I have a string
  `this is a large language model`
  and a sub-string to be found
  `language model`
  If language model is present then it is good, but let’s say the string is truncated to
  `this is a large language mod`
  In this case `language mod `is partially in `language model` it should still be allowed, however it’s a bit tricky and have tons of edge cases covering various strings.

## Tech News

- [LlamaCon](https://techcrunch.com/2025/02/18/meta-announces-llamacon-its-first-generative-ai-dev-conference/): Meta has announced the first tech conference for AI/LLM.
- [Grok 3 Released](https://techcrunch.com/2025/02/17/elon-musks-ai-company-xai-releases-its-latest-flagship-ai-grok-3/): Elon Musk’s AI company, xAI, late on Monday released its latest flagship AI model, Grok 3, and unveiled new capabilities for the Grok iOS and web apps. This comes in twitter / X as a chat access. LLM is now like a search button everywhere.

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-734) and for software development/coding articles, join [daily.dev](http://daily.dev/).

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
