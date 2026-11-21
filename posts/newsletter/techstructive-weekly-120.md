---
title: "Techstructive Weekly #120"
date: 2026-11-20
slug: techstructive-weekly-120
type: newsletter
status: published
source: newsletter
canonical_url: https://techstructively.substack.com/p/techstructive-weekly-120
description: "Read, watched and learnt about AI engineering amongst the other things while observing the choas of the internet with CVEs, RCEs discovered like crazy in the week from 15th to 21st November 2026"
tags: ["newsletter", "substack"]
---

## Week #120

It was a action packed week, I mean not physical, but the drama around tech and things. Nvidia blockbusters, Outages and RCE after RCEs in major software tools and products, wierd model releases from Gemini and OpenAI, and more RCEs. This is exciting and scary too. On one hand, I feel, software development is here to stay with the slop and mess the AIs are creating, and on the other hand I am scared too by getting hit by one of the RCEs at some point. Also there is a fear that the vulnerabilities will eventually plateau and all software will be perfect? Maybe that’s a scifi possibility. But having personal agents was a fictional scene too, a year ago! It just turned tides within few months right?

So I have been thinking to enjoying this ride while it lasts, I think I need to build products and stay in the loop. I think I have stopped paying attention to the skills and tools in my toolbelt. Partly because its changing too fast to understand and even consider anything a tool, like a year back, tool calling was blowing up, then we had mcp, then skills and graphs, loops like what are we converging at? It just keeps on changing and the pace is just crazy.

### Quote of the Week

> “Any sufficiently advanced technology is indistinguishable from magic.”
> 
> — Arthur C. Clarke, [Profiles of the Future: An Inquiry into the Limits of the Possible](https://www.goodreads.com/work/quotes/163451)

Yes. Doesn’t AI Agents feel like that, just give them a skill.md and boom, you have some level of augmentation in your workflow. Just like in the beginning of January of 2026, we had the moment with openclaw or whatever else it was renamed to right? The possibility of a ai-agent doing things for us felt like magic, like a sci-fi world thing, but just in a blink of an eye, we had witnessed that right? And it faded away? Not quite, the promise of personal agent was not just limited to the silly things people can do, but it expanded on the long running software engineering tasks and goals to build something wild, or even unattainable, like reverse engineering products and games and mash ups and what all have we seen in 2026, with a month more to go! Its beyond magic it feels to me.

## Created

- A bot that can do almost anything hosted on cloudflare workers with GitHub integration
  
  - I am still testing this out. I wanted to have something like GrokBot but with free llms and any llm that I wish to use. The agent can create repos and basically crank up a webpage or a report of some sort after some research or a prototype.

## Read

1. Dbushell Blog: [We don’t need chatbots](https://dbushell.com/2026/11/19/we-dont-need-chatbots/)
   
   1. Man this is so funny to read. Its true that we don’t need chatbots we need jarvis right? Who can talk in real time. I don’t think so the voice mode and those are that great at it yet.
2. [All AI generated code should be opensourced](https://jakub.jankiewicz.org/blog/ai-code-open-source/)
   
   1. How do you know if its AI-generated? At this point, from 2026, I can assume, 99% of the code is AI-generated, what is the difference in opensource vs not opensourcing it if the models themselves are getting trained on it somehow?
3. Pudlog: [Building in public is dead](https://pudlog.substack.com/p/is-building-in-public-dead)
   
   1. Where it started and where it went. I was excited to learn about the formal verification software he was building and it turned out to be 60% startup founder advice. I don’t hate it but it just falsified the promise of the article, after a exciting buildup.
4. [People don’t subscribe to blogs](https://micjm.com/blog/front-page-of-hacker-news-by-the-numbers/)
   
   1. Yes they don’t, nor do I. But I do occasionally search for them. I don’t follow a rss feed. Because I didn’t knew it was a thing until I had a more people to follow.
5. [Folscaps work](https://foolscap.works/)
   
   1. This is interesting product. I does everything for your book except to write it. Brilliant idea. I can think of it as a peer or a assistant giving me all the details and the nuances of the things I want.
   2. But the counter point here is the friction is what creates those connections in the story. Like if i want to look up something I wrote in other chapter or other scene, while reading some other parts I might discover a flaw or a idea might strike, whereas when its handed over in a instant, it becomes harder and harder to stick to the story.

## Watched

- [The CS50 Show: Season 1, Episode 3](https://youtu.be/KSZz4DPIj50), GTA 6 and the Nvidia blockbuster
  
  - Hmm This was a good one. Like the explanation of hashing, upload and download speeds, the CUDA and the moat of NVIDIA was amazing.
  - This is a great form of show, I like it and I would like to make my own too, but it becomes a bit exhaustive after writing this newsletter out.
  - The GTA 6 is 200 GBs! That takes around 1.5 hours of time if download speed is 300 Mbps, which is too high, its around 70-100 Mbps at my location. It would take around 4-5 hours to download that.
  - The vulnerability of crypt and the libraries is so funny like a * can come from no password, or if its NULL, and a condition is set when the password is beyond limits of 255 characters, the error is NULL, hence it matches and we get root acccess. How dumb is that if you think, but how intricate it can be because it has been since 2016. My god, 10 years in uncovering this?
  - The parallel processing makes so much sense now, They have 200-21k cores rather than 4 or 16 core CPUs. They are rightly called GPUs, they are monsters.

<!--THE END-->

- Creator of PStack (poteto) and Matt discussing the agentic engineering
  
  - This was really insightful. The concept of kitchen and chefs was profound. The concept of non-blocking the agents so that they can verify the changes is the key thing.
  - We all know agents became good enough in December 2025 so that we can almost let it generate what we want. But the year went by and we are just figured out some good ways to verify the changes, first we saw harness, loop engineering, goals, ultra glazing modes, graphs, mcp and what not, finally settling to skills and computer use right?
  - But its so much chaos, like everyday something new appears, or make it redundant, so much so that we have to entirely switch everything. The advice that “you will get left behind if you don’t learn prompt engineering” was so bullshit. In fact the opposite was true, the more late you were the more fresh perspective and no-biased approaches you would have.
  - The approach of sharpening the skills and toolset is also great and now I think about it, it makes me regret that I have wasted not learning or building any skills (like the md file I mean, not skill skill). But the counter point is that the change was so fast that I had to doubt and even today I doubt, will skill.md be relevant next year? or even 3 months from now?
  - Grok bots looks like a interesting concept and I am trying to build and use it for myself for some ridiculous ideas that I have using cloudflare and github and free llms.

<!--THE END-->

- [FInally a language for vibecoder, bend](https://youtu.be/c7k3ehM3oWI)
  
  - The bend language is like a law or a property that you want your software to poses. It was ridiculous to write it when we hand chiseled code, but when agent can write the code for us, and we just say the laws, it becomes a powerful thing I believe.
  - Its like a perfect blend of specification and programming.

<!--THE END-->

- Theo: I’m gonna miss coding
  
  - Ah! Yeah. We stand on the shoulders of the giants. The people who wrote in-depth articles, blog posts, documentation for libraries and frameworks, shit posts, opinions, good practices and bad practices, recorded podcasts with famous people, created courses and learnt the hard way.

<!--THE END-->

- Satya Nadella on Pragmatic Engineering Podcast with Gregley
  
  - Ok, I think among the tech giants, I like the approach of Microsoft more. because it has some empathy and humanity. Just because he said, “I'm hoping to ensure is we create that opportunity for folks because at the end of the day, right, you know, it's about economic return on your professional careers that matters” That is a pure empathetical statement. Beautiful.
  - This is a good take. This is the values that Microsoft brings to their employees and as a team. The team re-structuring and voluntary exit is not just exit but also a new beginning of someone’s role, which is a honest take in these rapid changing times. I think it is fair then to be brutal to lay people off or be delusional till the end of the times.

<!--THE END-->

- M[icrosoft Engineer breaks down the incident at Microsoft that has been caused by months long hidden bugs](https://youtu.be/iRnZ81QL3oU)
  
  - Oh ok, that is a bit disastrous. We cannot blame slopiness of code here, because it could be obvious to humans but not to AI right? But the problem got more intense as they flicked the lever and restarted the systems which made even worse. Because those systems also started going down and one by one each region went down. One of the simplest yet weirdest incidents in the history of computing in 21st century? It took 36 hours to recover.

<!--THE END-->

- [The Standup: Vibecoding golf games](https://youtu.be/IXoSzZbZ4jw)
  
  - Wow. I am inspired by these. But man playing games feels like ancient times.

## Learnt

- Microsoft has a Graph Mail API that helps in sending and receiving mails or syncing calendar events

## Tech News

- Microsoft had an outage for nearly 36 hours and also leads to [chlorine poisoning for 15 guests on a island](https://www1.wdr.de/nrw/rheinland/rhein-erft-kreis/technischer-defekt-chlor-phantasialand-1-100.html)
- [PostgreSQL 19 released](https://www.postgresql.org/docs/release/19.0/)
- [NGINX HTTP2 RCE vulnerability](https://www.wiz.io/blog/nginx-http2-rce-vulnerabilities)
- Google is dropping models like shit, and they seem to be shit, [Neon](https://blog.google/innovation-and-ai/models-and-research/gemini-models/introducing-gemini-4-neon/), Hydrogen, what is going on?
- [OpenAI releases GPT 6.2 Astra and Sol](https://openai.com/index/introducing-gpt-6-2-astra-sol/): I don’t know what the differences are since I cannot test them.

* * *

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/819) (#819th edition), and for software development/coding articles, join daily.dev.

That’s it from the 120th Edition of techstructive weekly. I hope you found it helpful, and relaxing. If not please drop any suggestions, feedback or discussion about certain things you want to in the comments or drop me a message on my [socials](https://www.meetgor.com/contact).

Thank you for reading,

Until next week.

Happy Coding :)

[Leave a comment](https://techstructively.substack.com/p/techstructive-weekly-120/comments)

[Share Techstructive Weekly](https://techstructively.substack.com/?utm_source=substack&utm_medium=email&utm_content=share&action=share)

[Subscribe now](https://techstructively.substack.com/subscribe?)
