---
title: "Techstructive Weekly #96"
date: 2026-05-29
slug: techstructive-weekly-96
type: newsletter
status: published
source: newsletter
canonical_url: https://techstructively.substack.com/p/techstructive-weekly-96
description: "Reading and writing some stuff, feeling a it burned out after productive maxing with AI agents, among the other things observed and watched in the week from 24th to 30th May 2026"
tags: ["newsletter", "substack"]
---

## Week #96

It was a work heavy week. Trying to build a better habit and sticked to it the entire week. Feeling great again, a bit stress and ai psychosis maybe but coping with it, thanks to the good habits. I started writing an page on paper, not my thoughts, a story, maybe childish memories and essays, but it helped me feel better, just 1 page a day. I also started reading consistently 10-20 pages a day. I finishes two books over the weekend and started another 500 page book after a 3 month gap. I read books around 150-200 page range. Its great so far, will update next week how it went.

I am finally thinking to live stream this weekend at last. Maybe it won’t happen depending on my energy levels. I want to practice coding again, the old school way. I want to learn elixir or more golang.

This week I interacted heavily with AI, claude and cursor as harness. I love them both, I like the intelligence of Claude and the ergonomic and swiftness of cursor. I use cursor when my 5 hour limit exhausts in 1 hour. I think I am learning how to work with them. I make mistakes and waste tokens, but that’s how you develop a intuition on what will work and what won’t next time around. The models are not changing that rapidly each month now, from the past couple of months I believe. But that is life now, no pair programming, its pair ai to work and stare at.

### Quote of the week

> “A library is a good place to go when you feel unhappy, for there, in a book, you may find encouragement and comfort. A library is a good place to go when you feel bewildered or undecided, for there, in a book, you may have your question answered. Books are good company, in sad times and happy times, for books are people - people who have managed to stay alive by hiding between the covers of a book.”
> 
> — [E.B. White](https://www.goodreads.com/quotes/383616-a-library-is-a-good-place-to-go-when-you)

Ah! Library! If only Mumbai had a few more. Still being able to read free books in this digital world is a blessing. I can’t express in words how good I feel when reading a book, the characters, the moments all just make you feel a little seen, a little taste of relation and hint of emotion. That is what I need when escaping from this AI-slop and productive-focused work. I will continue to devour books.

## Read

1. [We should be more tired than the LLM Model](https://vickiboykis.com/2026/05/28/we-should-be-more-tired-than-the-model/)
   
   1. Of course we are. But not by slowing down and doing the trad coding, but by spiraling into Agentic code gambling. I might coin that term. Its quite exhaustive to be honest. Not that I love it, I am forced to work this way, with AI, we are asked to be more productive, more smarter. How to address that query?
   2. When to switch from agentic coding to trad coding. Do we even have a choice at this point? The expectations and the reality is just staggering with a divide. In the middle are developers, crumpled with expectations of gains and productivity.
   3. I just wish we all had a choice, but life is harsh and cruel at times, there is no option than to face it. I guess that’s how you learn lessons and be a man with scars and experience. Its rough out there, not gonna lie, exciting yes, but dreaded by expectations.
2. [Can we have a day off?](https://mlsu.io/posts/day-off/)
   
   1. Beautiful. What a sarcastic blog.
   2. If we are 10x productive, then can’t we just take friday off? Yes, we should be technically, but that’s not how it works right? Or I should say, they don’t want us to. If you are productive then we can do more work, that is the new north star metric right. I hope that doesn’t make us literally reach the north pole. Cold.
   3. This won’t be sustainable I can see, working 10x productively and still having the same day structure of rest and work, is not going to make people happier or healthier. It might have impact later in their lives.
   4. I am not joking, I have started exercising, just walking the least an hour. It has helped me stay a bit more attentive and focused. Its not easy to work 12 hours and still get the motivation to work each day, day after day, 5 day a week. Even sometimes, “you can work over the weekend maybe” and “take next week off”. That is just like saying, “suffer now” but rest later, no remedy just take the cure later, no precautions.
3. [Various LLM Smells](https://shvbsle.in/various-llm-smells/)
   
   1. We all have known “em dashes” and “delve” as LLM smells, but this goes beyond it and lists a few more. The voice of the language just gives away if you have generated enough and read enough. Why both? Well, you must know if something is LLM-generated in front of your eyes, you cannot trust anything anymore, if someone says this is AI generated, how can you know if you have not generated some and read that?
   2. The website generated by LLM oh! The buttons, the pills, the cards are all give-aways of slop. If you have generated and seen that first hand then only you can smell it.
   3. I would say “To be able to smell, you need to cook and taste and smell it yourself”
4. [Tony Olieverio: The worst interview I ever had](https://www.oliverio.dev/blog/the-worst-job-interview-i-had)
   
   1. Wow! I thought it might be catastrophic system design or brutal hard leetcode. Sigh! It was a cultural fit question.
   2. How can one reject a candidate if he was honest in his terms, and even if it was not fit for them, why not put it right in the start without interactive element. I think I am wrong here. That rejection might be a great thing for the company as they saw he was not the right fit (maybe). But that candidate had to take the toll of that decision. Its not easy for some kind of people (like me, introvert max) to express their emotions, or even experiences verbally to a person.
5. [Orchid Files: I’m tired of talking to AI](https://orchidfiles.com/im-tired-of-ai-generated-answers/)
   
   1. Wow! The title should be 2026 in a gist.
   2. I have sometimes done myself this at work. But not in real-world conversations of some other person. Using AI is good but it cannot do what a human conversation can. It cannot inspire a deep cord in the nervous system that a human connection can.
   3. Yes it can give you ideas and inspiration, but that won’t stick like a real human connection can.
6. [How soon is now in PostgreSQL?](https://www.architecture-weekly.com/p/how-soon-is-now-in-postgresql)
   
   1. Great article. A good thing to know that `now` doesn’t take a new time in a transaction, it keeps the same timestamp the better way to use a current time in the transaction is `clock_timestamp()` , that makes it uniquely timestamp each call inside a transaction.
   2. Neat little thing to know before you shoot yourself on the foot.
7. [Opaque types in Python](https://blog.glyph.im/2026/05/opaque-types-in-python.html)
   
   1. This is neat. To encapsulate the data and not overwhelming the user of the software to all the details, just providing/exposing the abstract thing to customize or compose the object is better. It lowers the maintenance toll on the developer and prevent the code from bloat and convoluted logic, makes it backwards incompatible.
   2. A good idea to keep in mind while designing configuration.
8. [What I’ve learned building online mini games in Elixir and Swift](https://calvinflegal.com/2026/05/24/what-ive-learned-so-far-building-online-mini-games-with-elixir-and-swift.html)
   
   1. This is worth reading. I want to start learning elixir, I had the intuition to do so, as I set out for doing traditional coding. I want to learn a language and build good software. Not technological decision fatigue right?
   2. But elixir is good with AI today, might not be tomorrow, LLMs might improve on other languages which it can validate more. Maybe we don’t ever need programming languages, who knows. I don’t quite believe programming languages but it might not matter much.

## Watched

- [Theo’s take on “Cursor just crushed Claude Code with Composer 2.5”](https://youtu.be/UvUzpSlXKtg)
  
  - This was about to happen as Cursor has the best data and focuses on Coding. Unlike Anthropic who loved coding initially, it has become a side quest which they think they are the crown holders for eternity.
  - The speed thing is just crazy, I cannot be more happy, like its cheap, its fast, its cray smart too. If I can get work done in multiple iterations (and it takes multiple iterations to be honest, no model can 90% one shot it, no ways, models are just hallucination gimmick, the limited the context the worse the output)
  - Speed might be double edged sword too, since it can produce slop 2x the speed maybe, because if something gets quicker and cheaper, it becomes very hard to resist using it again and again. But we don’t have the mental cognition, or bandwidth or capacity to review and understand what it just spitted out. So reviewing and evaluation and understanding becomes a bottleneck.
  - But still its great to have this option of speed of low-level grunt work. Just to get it out of your way.For intense tasks, I don’t think developers would prefer this or any fast or dumber model.

<!--THE END-->

- [Theo’s take on “Claude vs Codex vs Cursor”](https://youtu.be/JMYspR42HFM)
  
  - A good observation here. I didn’t see the claude code’s wired animation until now. Yes its glitchy but I think its not that gimicky right? Is it just me that finds claude code not as crazy as theo just mentioned, to get attention. However I do feel that they definitely have made it to feel “productive”. Yes. All coding agents do feel like it. Honestly. It happened to me with codex, because I was exploring these and claude code is not for free users, I started using codex extensively in the month of Feb-March and it was really addictive.
  - The codex GPT models are really good at doing work, good work. And that just made me go wild and build wired things, that I might have thrown out as a ambitious project. Amp has been the same kind of a drug dealer too, its a lotttt faster than codex and claude code, but not great accuracy wise (I am talking from a free user perspective).
  - Cursor is a separate league, its just perfect for me. It has good ergonomics, no screen flicker like claude code and codex. It is just amazing to work with as a reliable agent. I have used it for I think 8 months now, and have just loved it. I hate claude code but have to deal with it just because it is smarter. And also it has the annoying 5 hour cap, the cursor feels limitless, though I am subscribed from my org’s behalf.

<!--THE END-->

- [The Standup: Adam from Open Code and recovering from AI Psychosis](https://youtu.be/cVUVfn8OF5k)
  
  - Oh that is AI Psychosis. I thought it was gambling addiction XD
  - The problem of “I don’t know what I want to build, let me prompt it anyways” is real, that is a gateway drug to the slot machine and gambling like traits for programming. It starts by some thing it created by a simple prompt and we are in awe, a dopamine hit to extend that prompt and try again with this tweak, by checking that, by asking it to patch it and all of that instruction to make it work. But it spirals into a loop.
  - Everyone is being suffering from that, only few are lucky or I should say earned position, luxury of being able to take time off. Not everyone can take a week break and find it better, rather its extremely double the tasks from normal. People are just crazy on how developers can be used. They want us to use AI as a native tool. I understand that, but it cannot happen over-night and especially not in times of deadline of client requirements. Juggling between client expectation and AI psychosis is a pain, a literal head ache.
  - I don’t when will this era be over , its too mentally taxing and if not broken it can lead to a bigger burnout that nobody talks about.

<!--THE END-->

- [Aaron Francis: Agent orchestration shouldn’t be that hard](https://youtu.be/WAKGhlzpYgs)
  
  - Well, this was a moment for me to be honest. I started using claude code in full capacity, and for the first time I saw it spawning sub agents.
  - Like 10s of them for each thing I requested. It felt like it will collapse into context junk and my 5 hour limit might just go in a puff. And it went eventually but after completing the core tasks, the sub agents that were spawned were completed and the main agent had summarized the findings, which was a surprise to me. That was the first time I saw agent orchestration.
  - I don’t if that is a big problem right? I mean agent orchestration. I don’t know if people can afford to have multiple agents as in codex+claude code+amp. All are hell expensive. I can barely purchase my domain and keep that going for a year, rest aside the ai subscription is not cheap.
  - But for someone using all of them, yeah there should be a way, a singular way to consistently manage those agents and subagents. Solo looks pefect for them.

<!--THE END-->

- [Coding is no longer the bottleneck as English is the new programming language](https://youtu.be/WxQ-1_BIiBo)
  
  - Hot take here. I thought he was going to agree to that, what a clickbait. Not saying it was bad, I loved that rage. I wanted it more. This is the awareness that we need to create.
  - Subject matter expert is something that LLMs might struggle with and English is not a tool, its a medium to express it. If you cannot express the intent (which you won’t be, because you cannot understand everything and don’t have the experience of that thing) its not going to understand it. Simple right? If I say “bring me water” that’s vague right? Where to bring it, how to bring it, for what. The context is missing, the intent is not conveyed. Just like that, you cannot express in english for what you cannot understand.
  - English might be a good abstraction, LLMs have lowered the abstraction, but the problem remains to understand the problem. That is the last thing that will be replaced by LLM and I bet never will.

## Tech News

- [Anthropic Drops Opus 4.8](https://www.anthropic.com/news/claude-opus-4-8)
  
  - They admit its not a great improvement, but its more “honest”. Which translates their previous models were dishonest, or less honest. This model has similar hallucination or dishonest percent like mythos in that chart. A bold claim.
  - Also they are planning to release Mythos in coming “weeks”. A month earlier it was “Risky” to share in the open world and now its ok? In a month it was ready?

* * *

That’s it from a long week. It was not that long, but a bit too much to do without a peace of mind. Will take the weekends as usual to recharge but if I have the energy to speak or engage in a live stream would gladly start one.

Thanks for tuning in this week, hope you have a great week-and-end ahead!

Happy Coding :)
