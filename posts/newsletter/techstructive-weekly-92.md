---
title: "Techstructive Weekly #92"
date: 2026-05-01
slug: techstructive-weekly-92
type: newsletter
status: published
source: newsletter
canonical_url: https://techstructively.substack.com/p/techstructive-weekly-92
description: "Back to work week, reading ai discourse, watching the end of subsidized token era, among the other things read, watched and learnt in the week from 26th April to 2nd May 2026"
tags: ["newsletter", "substack"]
---

## Week #92

I enjoyed working this week, I am surprised I am saying that. The same thing a month back was dreading me, and now its making me wake up each morning. AI is kind of a progression and also a setback in some sense. I don’t know, but the usage will steer where the world moves with it. There will be a divide, of course, and if I am correct, it will create the biggest divide the world has ever seen.

But till that happens, just sip in the beauty of the nature, and read some handwritten articles, human voice, and some hot takes on the internet about tech.

### Quote of the week

> “To know that we know what we know, and that we do not know what we do not know, that is true knowledge.”
> 
> — Confucius
> 
> ― [Henry David Thoreau, Walden or, Life in the Woods](https://www.goodreads.com/quotes/383876-confucius-said-to-know-that-we-know-what-we-know)

How true is that? If you know something, you are eager to express or impress, but when you don’t, you shy away or fear away. The later part is what makes us humans. The latter, the complete part, is what distinguishes us from the AI bots. They will hallucinate and pretend the opposite when they don’t know something, they are unpredictable even in the first known part.

Knowledge is not learning something, its about intuition and learning to adapt. Being honest is the best thing to do to yourself.

* * *

## Read

01. [To my students](http://ozark.hendrix.edu/~yorgey/forest/00FD/index.xml)
    
    - Read this article, if you are a student or not. You will find answers and sentences that matter to your soul and not your ego.
    
    > Above all, be motivated by love instead of fear.
    
    - Everyone is in fear, unsure of the future, don’t be. Recollect the feeling when you started to do the thing in the first place, I bet it was not money, if it was, and still is, then I don’t know, if it was love, keep on steering.
    
    > Have the courage to go slowly, especially when everyone else is telling you that you need to go fast and cut corners
    
    - Vibe coding, running agents, and shipping without reviewing, these are hype and will fade with time. In professional development, at least, this won’t be the standard as far as I know.
02. [Ghostty is leaving GitHub](https://mitchellh.com/writing/ghostty-leaving-github)
    
    - This is bad. Oh my god, he has spent almost two decades on GitHub. I might not have lived consciously for that many years.
    - Man, this person has breathed and lived through GitHub. I can resonate with the line “When I went through tough breakups, I lost myself in open source... on GitHub”. I have done that to some extent, too. Not at his scale, but when I didn’t get any internships, got rejected in interviews, I found SQLC, Turbot, and a handful of organisations to fill my hunger for learning and programming. GitHub was the center.
    - I don’t rely on GitHub since I don’t maintain or do any hustles anymore, just because life is tough and so is the AI hype. I can relate to the outages happening these days. One of my projects [flight-observatory](https://github.com/Mr-Destructive/flight-observatory), had a few failing actions notifications. I thought it might be 403 or some upstream external API issues, but GitHub was the reason behind it.
03. [You can beat binary search](https://lemire.me/blog/2026/04/27/you-can-beat-the-binary-search/): Quad Search
    
    - This is creative. Instead of just halving the list, we do a buckets of 16, quads, and compute which of the buckets the element can be. Since we have the max of each bucket. We can find the bucket and parallelly then compare each element in it.
    - Surprisingly, this is better for large arrays, and with multiple cores this will outperform binary search easily. Because we are not diving the array into halves anymore, we are picking the most possible region where the element could be and parallel searching it in those.
04. [Full Text search with DuckDB](https://peterdohertys.website/blog-posts/full-text-search-w-duckdb.html)
    
    - This is quite similar to what we do in SQLite. Pretty simple. I like how DuckDB just keeps on making transitioning to modern databases, but still keeping the backwards compatibility and keeping it simple and dead easy to use just like SQLite.
05. [I got sick of remembering port numbers](https://gregraiz.com/blog/local-vibe/)
    
    - He got sick of remembering port numbers, so he just reinvented ngrok for local?
    - Yes, it is different than assigning names to services, but I feel that it’s doing that similarly. It’s a clever trick.
    - This is the phase where people are building anything they want; the idea of personal software and developer tooling is just skyrocketing.
06. [AI, Tokens and the gathering storm](https://www.pootlepress.com/2026/04/ai-tokens-and-the-gathering-storm/)
    
    - This is quite evident from the recent events and rug pull of Anthropics subsidized tokens, and GitHub’s move from tier based to usage based billing. Its on the wall, its the end of the subsidized token era.
    - The next few months might actually tell us how expensive are these tokens.
    - Local LLMs be ready for shinny demo, the dark days might be about to end.
07. [I asked Local LLM to add 23 numbers, I got 7 different wrong answers.](https://viggy28.dev/article/local-llm-seven-wrong-answers/)
    
    - My last sentence was jinxed. LLMs can’t do math yet. Sigh!
    - I did explore this last year and they failed miserably, till date this hasn’t improved. I think we can create a skill for this shall we, I don’t want it to use a python interpretter for such a trivial thing to add. Can we make them add? It would consume a bit more tokens but let’s see. I have some ideas here.
    - But the point being, Local LLMs are not quite the hype that the propreitary models live up to.
08. [Agentic Coding Fatigue](https://www.0xsid.com/blog/agentic-coding-fatigue)
    
    - True. This is me reading to myself. Writing code gives us the clarity, becuase we were in the weeds, we know why each of the if else was written, we knew the edge cases, but now? We just run the slot machine to “fix it” prompt and cross our fingers so that it works.
    - The balance of understanding and the need to understanding, is wide and is growing faster than ever. I don’t know which one to lean onto. The former sounds like there won’t be any harm. But the software is such a term a thing, that nobody really cares how well you understand it. Only the output matters, not the usage of design patterns or Python or Golang.
09. [Do I belong in Tech Anymore?](https://ky.fyi/posts/ai-burnout)
    
    - This one is really rough to read through. Not that it’s bad writing or thoughts, its is dead real and truthful. I myself am not able to express the true things that I have to go through, maybe I am not as privileged or have made my situation like this.
    - The principles that the author listed are gold, make sure we point it here too
      
      - Things that are worth doing are worth doing well.
      - Things that are done well require time and effort.
      - You make meaning through the doing.
      - Ideas are common; effort is not.
      - There are no shortcuts.
    - Frame it in gold, or memorise them, it is better to remember this than to be a slave to an agent or a corporation that is.
10. [People who don’t use AI will get left behind](https://migrainebrain.bearblog.dev/people-who-dont-use-ai-will-be-left-behind/)
    
    - I will frame it as “People who use only AI will definitely get left behind”
    - People would stop learning and thinking if they just learnt to use AI and not anything else around it.
11. [Have you seen the new Excel?](https://idiallo.com/blog/have-you-seen-the-new-xl-ai-parody)
    
    - Learn Excel or get left behind, learn AI or get left behind. Everyone is on a roll with hypes. I do agree, excel has just revolutionized the way people think of software or even any form of work.
    - It has abstracted the code in such a way that people rely on the results and not the code. Superb product.
    - I worry AI is doing the same, but for text generation, for code generation. It might be the de facto thing to produce custom software on the fly.
    - Until that happens, keep learning Excel ;)
12. [Where the goblins came from](https://openai.com/index/where-the-goblins-came-from/)
    
    - This was an interesting nerdy read. I liked they are open about it. Though they don’t back away from shoe-horning codex glazing. I don’t doubt they use codex, but its quite a bit of token-hungry and slow model.
    - The problem of the goblin is really interesting. The goblin behavior wasn’t a bug, according to them, it came from reinforcement learning rewards that favored playful, creature-based metaphors in the “Nerdy” personality.
    - Those rewarded patterns spread and amplified across training loops, even outside the original context. It resulted in a small stylistic quirk becoming a generalized model habit, showing how local incentives can unintentionally shape global behavior.

## Watched

- [Making Sense of the AI hype](https://youtu.be/HgNKa9UlRF8) - Wading through AI -Episode 2
  
  - The topic of engineer writing the blog and the manager or the sales person making the video is a reality and the only needle making the AI labs floating.
  - Its all gimmick play. Models aren’t quite capable of creating a novel thing like a compiler from scratch, that is, way far currently.

<!--THE END-->

- [AI isn’t taking jobs its worse](https://youtu.be/NZa5lApeFic)
  
  - Anthropic is eating the money, by creating fear. Like organisations need things quick so they burn tokens, for that they need to cut down on cost, which surprise is got by firing employees. So, the money eventually flows to Anthropic by cutting the job, which is hype honestly speaking. Not sustainable, they themselves have proven it.

## Learnt

- String Pool in Java
  
  - String Pool is a special memory area inside the heap where unique string literals are stored. The idea is instead of creating a new object every time, it reuses existing strings from this pool when possible.
  - A string in a string pool is a deduplicated storage of string literals inside the heap.

<!--THE END-->

- Cursor agent &gt; Claude code
  
  - I like cursor agent cli, it is intuitive and doesn’t block the user, Claude code is like I(the agent) am the owner. I don’t like that.
  - Cursor has the ability to view the output of the tool calls with Ctrl+o and let the screen continue, whereas the Claude just blocks the view.
  - I hate anthropic that is one reason. Maybe, but honestly speaking, claude code is blaoted.

### Random Tidbits

- https://www.numberempire.com/
  
  - A cool webpage to see different patterns for a given number

## Tech News

- [Copilot is moving to usage-based billing: The end of a subsidized token era](https://github.blog/news-insights/company-news/github-copilot-is-moving-to-usage-based-billing/)
  
  - This is a sign guys, if you still think, AI is not hype. The hype was at its peak, and now its wading. The subsidized token era is over. The real cost is on the labs and the providers that are now realising the importance of sustenance.
- [Warp is now open source: Another attempt to win the race?](https://www.warp.dev/blog/warp-is-now-open-source)
  
  - This looks like a desparate attempt to get into the lead of the race, I am not against it by any chance, its a solid move.
  - More organisations should do this now, espcially the ones that created this right? Yes I am talking about A…. Never mind.
- [HERMES.md in git commit messages causes requests to route to extra usage billing instead of plan quota](https://github.com/anthropics/claude-code/issues/53262)
  
  - Again. Never mind. Nobody would be using their coding agent.
  - They don’t leave a chance to not to hate them.

* * *

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/792) (#792nd edition), and for software development/coding articles, join daily.dev.

* * *

That’s it from this week. It was a bit of a work-heavy routine. Finally feeling back, maybe not peacefully good but warmly good. Looking up to slowing down with manual coding and building some impactful things in the era of slop.

Till then,

Happy Coding :)
