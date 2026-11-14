---
title: "Techstructive Weekly #119"
date: 2026-11-13
slug: techstructive-weekly-119
type: newsletter
status: published
source: newsletter
canonical_url: https://techstructively.substack.com/p/techstructive-weekly-119
description: "Reading, learning and watching abut AI and the software apocalypse continues, among the other things I observed in the week from 8th to 14th November 2026"
tags: ["newsletter", "substack"]
---

## Week #119

It was a slow week. A week and weekend of Diwali. I spent some time reading and enjoying with family. Hence you will see a lot less read and watched content here this and next week too.

### Quote of the Week

> “Our leverage isn’t just in what we know, but in everything we’ve experienced that made us who we are.“
> 
> — [Shahar Glazner](https://shahar.blog/about/), in blog post [“on being a software engineer in 2026”](https://shahar.blog/2026/11/08/on-being-an-engineer-in-2026/)

I was reflecting back on the memories this week, while recollecting the year (FYI, it was a end of a year 2082 in vikram samvant). So just nostalgic and something hit my mind. What we are as a human or whatever I have experienced so far, is what makes me valuable, and not what I know or can learn about. The memories and the ability to relate things and being able to take conscious decisions is going to be a vital thing.

## Read

1. [When the Kubernetes API went silent: A three layer networking mystrey](https://cloud.google.com/blog/products/networking/the-three-layer-gke-network-mystery)
   
   1. This was quite a complex thing to debug and understand. Like combination of packets arriving out of order, Linux TCP bug, and a networking component dropping packets needed to recover the SYN-ACK.
   2. I discovered that I didn’t knew a lot of fundamentals of the networking stack. Some things I learnt were like NICs, TCP and Linux sockets to Kubernetes, proxies and Google Cloud's VPC, etc.
   3. The issue then I can understand was something like that made connections stuck in an unexpected state, and when a new connection tried to reuse the same network port, Google's networking layer dropped a packet needed to recover. So, like its not about single failure that you can point out, rather its a nested things to uncover from those. Like from a Pod and its Kubernetes node, through virtual NICs, VPC networking, load balancers and proxies, and control plane.
2. [Our software engineering interviews are changing a little](https://blog.janestreet.com/software-engineering-interviews-are-changing-a-little/)
   
   1. Ok, nothing is changing, just a little less scrutiny on specific code related questions. Developers will still be evaluated to solve problems, the instincts to grapple with a problem, the thinking and the approaches to build software. Like when to delegate and when to go deep is going to be a key distinction in becoming a better or remaining a developer.
3. [GraphQL HTTP specification](https://graphql.org/blog/2026-11-12-announcing-the-graphql-over-http-spec/)
   
   1. This is interesting. They might be supporting Query method too, this fits what GraphQL can add to it.
4. Shahar Glazner’s Blog: [On being engineer in 2026](https://shahar.blog/2026/11/08/on-being-an-engineer-in-2026/)
   
   1. Banger of a post. “our leverage isn’t just in what we know, but in everything we’ve experienced that made us who we are” What a line in the end. Hits home.
   2. We have been seeing the AI labs saying “We will not need software developers in 6 months”, “Software is mostly solved”, (say that in primeagent’s accent please)
   3. But that was in 2024, we are in 2026-27 phase, what is happening? Why are software developers still productive when they were meant to be obsolete? I think the easiest answer is no non-technical person wants to take the effort to build software. As simple as it can get, we have the most bizzare urge to learn something that will make something work, and that is what makes a software developers, nerds. Nobody is ready to just prompt it and assume it will work, the quirks, the underlying patterns is something you can’t ask a agent to do, because what will you ask? “Code as a software developer”, well good luck.
5. [What should we learn now?](https://ducktape.blot.im/what-should-we-learn-now)
   
   1. Hmm! Interesting one. Slow down when the people behind are flying in the planes right? Sounds ridiculous but that is I think a sensible advice. On anchoring the stand and advicing to learn things at our own pace. Nobody knows what $work is going to be, but hopefully it will be some problem to solve.

## Watched

- [Engineers are gardeners now](https://youtu.be/GtJKFhztgdM)
  
  - This is a bit hard to digest pill. Gardening is not about building or growing plants everytime, its about being able to setup or maintain the plants around or in the garden. Watering, laying the foundation of the soil, taking care of the soil is not infected or infertile, and among the other things that I assume a gardener does. Being a software gardener is going to be a bit unsatisfying in my opinion compared to the current role. Like what we will be doing as gardeners, the effect of it might not be immediate, it might not be at all, considering how much the acceleration is there in the software world right now.
  - But that point of being able to understand and unblock the system is key. being able to understand the system and then take actions to unblock or solve it, is going to be one heck of a skill.

<!--THE END-->

- Gemini is kind of cooking now?
  
  - What a model is that? That can reason through the cybersecurity audits. It might have been deliberate but it seems a good change.

<!--THE END-->

- [Crypto bros are having a time now due to lack of crypto or security?](https://youtu.be/KMOyP_ArDMo)
  
  - This is really funny. A binance infrastructure leaked by a single devops guy’s laptop, a employee’s work laptop used by him for interviewing and installing some software. Like what is that?
  - How can this level of casualness happen in a company that is one of the major for cyybersecurity.

## Tech News

- [Google releases its decision model](https://blog.google/innovation-and-ai/technology/developers-tools/gemini-decisions-api/)
- Google release Carbon and [Hydrogen](https://antigravity.google/blog/gemini-4-hydrogen-in-google-antigravity/): What’s up with the name change, but nothing seems to be change apart from that.
- Nvidia CEO quits
- [Anthropic releases Fable and Mythos 5.2](https://www.anthropic.com/claude-fable-and-mythos-5-2)
- [Meta releases a small model Muse Ember](https://research.meta.ai/blog/muse-ember)

* * *

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/818) (#818th edition), and for software development/coding articles, join daily.dev.

That’s it from the 119th Edition of techstructive weekly. I hope you found it helpful, and relaxing. If not please drop any suggestions, feedback or discussion about certain things you want to in the comments or drop me a message on my [socials](https://www.meetgor.com/contact).

Thank you for reading,

Until next week.

Happy Coding :)
