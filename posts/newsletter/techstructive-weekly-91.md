---
title: "Techstructive Weekly #91"
date: 2026-04-24
slug: techstructive-weekly-91
type: newsletter
status: published
source: newsletter
canonical_url: https://techstructively.substack.com/p/techstructive-weekly-91
description: "Learning Java, React internals, AI doomsday for programmers, model drop week, among the other things read, watched, learnt in the week from 19th to 25th April 2026"
tags: ["newsletter", "substack"]
---

## Week #91

It was a comeback sort of week. I still kept learning Java and fundamentals. I drilled down on dependency injection, object internals, some react too in the mix. I found LLMs are good at summarizing things, some are good, some are hallucinating, some are just give up. Yes, and I have opinions in this edition about that.

The articles are a bit AI-pilled I must say, those are getting a little into my head, expect little or no article from AI next week. I had enough for the month. But yes this trend is not stopping. We’ll have to figure out to deal with AI now.

### Quote of the week

> “The difficulty lies not so much in developing new ideas as in escaping from old ones.”
> 
> ― [John Maynard Keynes](https://www.goodreads.com/quotes/5099424-the-difficulty-lies-not-so-much-in-developing-new-ideas)

Its quite relatable here, Its not hard to develop new ideas now, we do it all the time, but now the problem is getting out of the mindset of living without AI or old mindset. It might not be completely AI but a mix of balance of conscious understanding of what to do and what to avoid.

Being a developer in this AI world might require us to shed certain ideas, the perfect code, the developer experience, the best design and all of that nitty gritty. It matters still but there seems to be no time for that, if LLMs are able to generate 10K lines in an hour, how can a human review that in a day? Its not easy, nor its impossible. But there would be a shift in this mindset of how we look at programming a few years back at least. And that change is very hard to make once it gets molded into the cortex.

* * *

## Read

1. [Lights-Out Codebases and the End of the IC role](https://vascoduarte.substack.com/p/ai-assisted-coding-why-a-distinguished)
   
   - Ouch! That one is brutal and eye-opening. Yes, looking at this, it just feels like software developer is no more a thing. But when I look at real world, nothing seems to be changed. Maybe its a bit slow, maybe people realize it slowly. But something is surely a divide.
   - Lights out codebases, this just makes total sense, I agree to that, because if LLMs just keep getting better, keep producing more code, more right code, human’s capacity to review that is out of hands. The chess engine analogy hits home. You don’t want to review the chess engine’s moves by a grandmaster right? Like what? See LLMs as the chess engine and the grand master as the developer. Yes, analysing a game between a human and a engine can be done with a grandmaster, but no one reviews code like a hobby. LLMs would react that point sooner or later where you don’t need to worry about the generated code, and at that point of time, why would one need a human developer to review the code.
   - Writing that passage got me some goosebumps, like its overthinking, but its grounded and factual. Not sure what lies for developers as a role ahead!
2. [Software Engineering might not be a lifetime career](https://www.seangoedecke.com/software-engineering-may-no-longer-be-a-lifetime-career/)
   
   - Maybe this is a bit negative, but factual, actually just tying up with the above post. It might not be slow, it might just wipe out the need for software workers in a single year who knows.
   - But world just changes drastically man, till 2025 people thought, software development is the hardest thing to replace with AI or robots, yet here we are the first ones to get down probably.
   - What would those people do? Farming, carpenter? use the software, yeah! we might become consumers rather than producers. No one knows.
3. [AI Resistance is growing](https://stephvee.ca/blog/artificial%20intelligence/ai-resistance-is-growing/) - read hackernews for better context [hackernews post](https://news.ycombinator.com/item?id=47839951)
   
   - The thread is civil but opinionated, with strong technical depth on poisoning feasibility mixed with philosophical debate.
   - Pro-AI commenters often acknowledge risks and call for better incentives, while AI-dommers focus on erosion of trust, job effects, and creative control, yeah! that is the typical conversations moving forward.
4. [Developer Taste and AI Slop](https://strategizeyourcareer.com/p/developer-taste-ai-slop)
   
   - This was an eye opening post. It is helping me understand, to differentiate generating vs developing. The code was and is never the bottleneck, what to build is. Even if that comes to you from other members upwards, a developer still has to think.
   - Anything can be made now, but if that is the real thing to make, worth it, or even needed is the question that would have an edge.
5. [My dear friend you just built kubernetes](https://www.macchaffee.com/blog/2024/you-have-built-a-kubernetes/)
   
   - This one is quite funny. After resisting the complexity of kubernetes, people actually re-invent the wheel.
   - To use any tool or software, we must know exactly what we are trying to do and what that tool or software is actually doing. Not necessarily it would fit, but it would help you reach the ultimate goal a bit quicker.
6. [Your company is a skill now](https://arlo.substack.com/p/your-company-is-a-skill-now)
   
   - This s a bit scary. I talked about it when the open claw and moltbook thing was getting hyped. We were just a bunch of markdown files away from creating a dystopian world.
   - Markdown files are just instructions and context relevant for doing certain task and not doing the rest of the side quests. I know that sounds too simple, but that is what it is. But its not that simple. Putting ideas and limited context to words is not simple as you think it is.
   - Maybe creating or translating ideas to skills will actually be a skill in some months, who knows!

## Watched

- [Should you be a carpenter? Wading through AI - Part 1](https://www.youtube.com/watch?v=RJyPVLMyyuA)
  
  - Ok, the person who is into the weeds of AI models, himself is confused as to what to learn and what to choose as a career. This is a bit hard to swallow. Are we done then?
  - Nobody know the answer whether software is a safe career anymore. At this point, we need to define what a career actually is.
- [I quit my GitHub Job because AI breaks software](https://www.youtube.com/watch?v=aEL384kTznk)
  
  - Ok, that is positive but people can get jobs just like that? I think its quite rough out there.
- [Microsoft accidentally told the truth about AI](https://www.youtube.com/watch?v=4CIlTOnc6I8)
  
  - Ops! Is it happening? AI Costs now trying to get better of the AI Labs. We saw with Claude subscriptions, now GitHub, cost cutting in place. First employees and now people will see the actual problem hidden underneath.
  - The tech debt after vibing everything. Maybe it won’t matter, because AI can clean it up, right?
  - But damm, the man just spoke a lot of things, they just gave me hope to be human. Just watch this, if you watch anything at all.
- [Chai Aur Code - React Series](https://www.youtube.com/playlist?list=PLu71SKxNbfoDqgPchmvIsL4hTnJIrtige)
  
  - What a beautiful series. It feels nice to learn things. I never though React was that simple. It just injects html and js. Wow!
  - Props are just attributes that we pass around the html elements, this is just art of teaching at best.
- [Did Anthropic just killed Figma?](https://www.youtube.com/watch?v=wDgq9aiuL-w)
  
  - Probably! This is a good product but not sure how sustainable it becomes. I mean I can use it if I have unlimited tokens and wallet to spend, otherwise its a big money burning machine.
  - And for that maybe figma is a good choice, old school. At this point, I am afraid what am I even going to see more from anthropic
- [We need to talk about GStack](https://www.youtube.com/watch?v=Rzi7oFTzjac)
  
  - A bit of wasted hour I would feel. Never no major thing to notice, except just rambling. It was not worth the time.
  - Gstack is basically I think bash script to give skills to a agent. Nothing much, and maybe it helps you produce 40k lines of code a day. Who cares.

## Learnt

- Dependency Injection with Spring Framework
  
  - This was a learning curve for understanding the core behind Spring Framework. Basically instead of changing values in object type, you just change it at the configuration level. So we don’t have to hardcode values. Its like the level of abstraction to define where to configure or inject the dependency. Hence the name.
  - Its quite clever to do this out of the box. In python I haven’t seen stuff like this. In django perhaps there is a way to just use [settings.py](http://settings.py/) and modify everything there without actually changing the model or view.
- LRU Cache
  
  - I tried to implement this in Java, and learnt to use Hashmap along with a linked list (doubly). Since we need to keep track of most and least recently used items, we need to be able to track it and more importantly for deleting and accessing in O(1) time, this structure just makes it quite great for this.
- Java equals and hashCode differences
  
  - The equals check if two objects are equal, their values and not reference value.
  - The hashcode method returns the hashcode of the object, its like the identity of the object.
  - The equals and hashcode needs to be consistent, so If you override the equals method, you need to make sure you keep the hashcode aligned with the equals too.
- React under the hood
  
  - Yes, I learnt a lot from the chair aur react playlist, and still learning.
  - React is a library, and Angular is a framework. React basically maintains its own DOM to effectively update the data.
- Table parsing for complex financial documents is still unsolved problem
  
  - Ok, I thought somewhere it might be solved but nope. Tables are a bit of a pain. Trust me when I say “pain”. A human himself cannot understand the layout, let alone LLMs. Poor LLMs. Even after giving them eyes (not realy) VLLMs are dumb too. I haven’t seen a great model like Gemini from other labs, if that breaks in, it should be a game over. But I don’t see those labs solving this problem.
- AI assisted/written articles needn’t be all slop
  
  - I just read this article sereis
  - [https://nitinsingh717.substack.com/s/threads](https://nitinsingh717.substack.com/s/threads)
  - Ok, please here me out. This is written by AI (but but). Its not slop. I honestly read it, all three posts. And it genuinely helped me understand concurrency better.
  - AI is not bad, but if you don’t have experience in the thing that you are generating, the taste (wow, everything is connecting now), then its slop, because slop is not AI-generated content, its AI-generated content that the author didn’t spent time thinking or reading about it.
  - And from the three posts I read till now, it genuinely felt that the model clearly understood the sentiment and the intuition of the author and it translated it perfectly.
  - I even discussed it with the author, and he admitted the same, humans are above LLMs for now.

[![](https://substackcdn.com/image/fetch/$s_!OSSG!,w_56,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Ffef8448b-7d22-408c-9fb9-6a4c072fb859_1024x1024.png)Level Up Your Programming with Nitin  
\
The Need: Why Sequential Thinking Breaks the Modern World  
\
Series: Thinking in Threads | Blog: 1 | Read time: ~14 min…  
\
Read more  
\
a month ago · 2 likes · 4 comments · Nitin Singh](https://nitinsingh717.substack.com/p/the-need-why-sequential-thinking?utm_source=substack&utm_campaign=post_embed&utm_medium=web)

- LLMs are a great tool for learning
  
  - Not all LLMs are equal, each has its own advantage.
  - I have found Grok to be best for search summarisation, Gemini for shallow explanations and quick searches, Claude for indepth reasoning and creative exploration. ChatGPT for general stuff.
  - Claude models are just great for agentic work, but they don’t quite search well. Grok is reallly great at searching and getting factual data in the context. These are opinions and not facts. But I have used all four enough and continue to do so, knowing what each of one can do and where I should throw which task.

## Tech News

- [OpenAI release GPT 5.5 - doubled the price of 5.4](https://openai.com/index/introducing-gpt-5-5/)
  
  - Another release. So, its plateaued now. Models are stagnant with the knowledge, and only breakthrough would be to make it sustainable for inference, rest is looking RL and post optimizations.
- [OpenAI releases ChatGPT Images 2.0](https://openai.com/index/introducing-chatgpt-images-2-0/)
  
  - Oh! This is good. Like the level of detail on the image I generated for Java OOP was just top notch. The text is just flawless.
  - Nano Banana was not that good at that precision. The crown has been stolen from Google to OpenAI again.
- [Canonical releases Ubunutu 26.04 LTS - Resolute Raccoon](https://ubuntu.com/blog/canonical-releases-ubuntu-26-04-lts-resolute-raccoon)
  
  - I just installed it on my work laptop and it was flawless. It has no major changes, but just software upgrade. However more rust is starting to get into the linux core.
- [Kimi Releases K2.6 - Open Source model are still keeping up with proprietary models, a good sign](https://www.kimi.com/blog/kimi-k2-6)
  
  - Wow. Open source models are not dropping the knees. The fight is on, the open source models are not as accurate or reliable as closed source but they are getting closer. The weight is a big factor for their adoption, if they can squeeze a great model in a relative smaller size, it can be a win for those labs.
- [DeepSeek releases v4 preview with pro and flash](https://x.com/deepseek_ai/status/2047516922263285776)
  
  - The fight is on. It keeps on pressing the closed source models to get out better models and we are stuck in a infinite loop.
- [Anthropic still shooting self on the foot - Mythos is over-hyped](https://www.theregister.com/2026/04/22/anthropic_mythos_hype_nothingburger/)
  
  - Yeah! Anthropic is not leaving their steak of shooting themselves on the foot. It has been over a month and the saga continues.

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/791) (#791st edition), and for software development/coding articles, join daily.dev.

* * *

That’s it from this week. It was a back in track, old vibes week. The next week is a crucial one, a make or break situation for me. If I am making a change, this is the moment. Its going to be hard, but I think I can give my best. The worst has already happened, there is nothing that can go worse than this situation for me. Anyways, looking forward to a positive week ahead.

Until then,

Happy Coding :)
