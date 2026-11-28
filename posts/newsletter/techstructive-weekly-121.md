---
title: "Techstructive Weekly #121"
date: 2026-11-27
slug: techstructive-weekly-121
type: newsletter
status: published
source: newsletter
canonical_url: https://techstructively.substack.com/p/techstructive-weekly-121
description: "Building new things, understanding different domains, seeing the mathematical revolution with AI among the other things read, watched and learnt in the week from 22nd to 28th November 2026"
tags: ["newsletter", "substack"]
---

## Week #121

It was a exciting week. I created a lot of things (slop + brainstormed slop). I don’t know if that is slop or not, but I think it is because I haven’t reviewed all the code nor do I want to, because it was working, with some limitations of course.

I learned a lot about one domain, it was not related to my work, but I have seen the pain in a user and want to solve it. But before understanding the problem (yes, pain is different from the problem) I want to understand the right mindset in why its a pain right? I knew a little about the domain, as I worked on a client with touching the sector in some or the other way, and had created almost 95% accurate system of automating a validation system. The next step for me was to understand the terms and the overall high-level overview of the industry. After getting the picture I built a prototype of it over the weekend, and it was great.

Next week and this weekend is going to be some fun project tuning.

PS: I was also watching Chess Championship (Gukesh vs Sindarov, first 3 games). It was a rollercoaster, first game a calm draw, 2nd game a draw in 17 moves like what happened there? The third game was so fabulous. Gukesh blitizing out moves till 20 and then still finding great moves as black. In the end a draw. Wow!

### Quote of the Week

> “The fragrance of a flower spreads only in the direction of the wind, but the goodness of a person spreads in all directions”
> 
> — [Chanakya](https://www.goodreads.com/quotes/7153244-the-fragrance-of-flowers-spreads-only-in-the-direction-of)

What a beautiful line. It also is about your natural talent. It will only catch certain attention. But a relentless pursuit of something greater good, will inspire everyone around you. You don’t have to please everyone, but your actions will affect everyone. The example I can give is, if you are a good software developer, you can solve a problem for specific use case, but if you are a good person you can solve problems for poor or needy people. I am not saying to leave your job and start a full-time service to people, but just saying to have a kindness mindset whilst pursuing your life.

## Created

- [Agent Weasley](https://agent-weasely.dev-meetgor.workers.dev/)
  
  - Ok, this is just vibe coded, but some architecture decisions were taken in rush. Developed by codex+opencode from free tier. Hosted on cloudflare and some shenanigans to make tool call work.
  - I want to develop it truly for agentic-first experience, anybody can just prompt in and they should get back results. I wanted it to work like a headless mode, maybe you can call it a agent on cloud, but it can be more than that. Once you remove the auth layer and restrictions it becomes a bit chaotic, so I might fix the auth security first.
  - I have one-off ideas, so built this, I think it should help anyone get ideas out of their head.
- The mail-client project
  
  - I am working on, there is no live view or anything, its just a mail client that I want to build. Its a domain specific mail client.
  - People might criticize me about building a mail client in 2026, but its the perfect opportunity to build a mail-client for a domain that is way in the weeds of manual work and making the process smooth is the goal.

## Read

1. [Cynthia dunlop on writing books and blog](https://www.ssp.sh/blog/cynthia-dunlop-writing-books/)
   
   1. That is a beautiful post. Very encouraging and inspiring.
   2. Some beautiful quotes.
      
      > “*I’ll likely turn to an LLM for **some critical feedback”***
      
      This is true like, I use it all the time, like this post even goes into the chatgpt for picking up things that I might have said and looks off or is wrong or even is misplaced. I never apply some suggestions(I am lazy and ignorant, but they are making this newsletter perfect, I don’t want the newsletter to be perfect, not the AI slop fiesta.
      
      > “*I will sacrifice sleep rather than skip some morning exercise.”*
   3. Wow. Like just 200 IQ moment here. I also tried this for a month, but recoiled into bed once the weather got a bit chilly like rainy. Its a bit tough to not get 8 hour sleep. I mean I am habituated to that and breaking that looks a bit tricky and cozy.
      
      > “*Start **writing in your own words** and develop your own style.* Please don’t let LLMs homogenize it”
   4. Good advice. And this opens up a point which is LLMs make the text homogeneous which is like average or muffled in the voice. Letting your mind churn the words onto the keyboard or a pen and paper makes it distinct, it adds soul to the whatever you are writing. AI-generated writing is just soul-less.
2. [What’s left for software developers](https://www.mooreds.com/wordpress/archives/3757)
   
   1. That is a good post. When software generation gets cheaper, all the things around it gets valuable like user-experience, marketing, domain expertise and so on. The product market fit gap is going to climb high as the production of software gets cheaper, the ability to build the “right” thing is the moat now (which always was, but its more than ever now).
3. [New set of problems for engineers to solve](https://www.katelynlesse.com/blog/new-problems-for-engineers-to-solve)
   
   1. Ok, multi-agent systems are truly creating a havoc in the software space. Its adding vulnerabilities and holes in the system. The blast radius is a good point since it can create a big mess for humans who are now starting to get off-hands in the code. So its a loss-loss situation if its in auto-pilot mode.
   2. But the hope lies in understanding that and building on top of it. The problem space just got bigger. “It doesn’t get easier, you just get faster” that is the thing happening here. Problems are getting harder, and we are faster due to ai-agents, a equal ramp-up.
4. [GitHub Actions accept invalid YAM](https://yossarian.net/til/post/github-actions-accepts-invalid-yaml/)L
   
   1. Of course they do, and when its incorrect they don’t even let you commit right?
   2. Is this slop or how did this even happen?

## Watched

- [Mitchel Hashimoto on the standup](https://youtu.be/PEuGhjZnScI): Standup #80
  
  - It was nice episode on discussing why he build rex after building ghostty. Rex was the realisation that oh! terminals are gaining popularity maybe we need to make terminals headless.
  - A good approach and direction of rex and ghostty. It is really thinking through the problem.
  - Apart from that, the prime’s point on the 12 ascii logic was straight up a banger, like that is what i like about being able to smell code and dig into the right parts, even if its ai generated or not, if you have written enough, you’ll know what is happening where.

<!--THE END-->

- [Google cooking with Neon?](https://youtu.be/HQlk9RGx2dk)
  
  - Oh ok, I might take some words back, when I said, google was dropping models like shit and they “were” shit, hydrogen, carbon, argon, but not neon, maybe. I have not tried it, people are saying its almost better than GPT and very coherent compared to opus. Not taking the word but maybe google has found the ingredients now?

<!--THE END-->

- [Riemann Hypothesis](https://youtu.be/d6c6uIyieoo)
  
  - That is some absurd hypothesis. Like how can one come up with that? I can’t even understand why it matters? Like what the heck.
  - And how would one prove it, like ok, if it was that simple to prove all non-trivial zeroes lie on the 1/2 point then why has that taken more than 150 years? Like that just sounds ridiculous and amazing at the same time.

<!--THE END-->

- Why does Riemann Hypothesis
  
  - Wow, I didn’t knew that was possible. That is a great way to learn things. I just applied it with chatgpt chat, and I had a lot of slop to read, finding the important piece is also a skill in these conversations. maybe we’ll get better over time.
  - But the way it explained and visualised it in a understandable way that was really awesome.

<!--THE END-->

- [How did they crack the Riemann hypothesis](https://youtu.be/PjGlt3y-Nco?)
  
  - Oh, that is a very simple yet mechanical thing to do right? Taking a existing approach and trying to fit the numbers so that it almost feels the right curve or the steep steps. Nice. But clanking a 10k page long proof, that is a different ball game for humans to uncover.
  - The new job of mathematicians is to clean the slop, same for developers!

## Learnt

- How to architect a mail client (custom api)
  
  - You have a mail id, 3rd part mail. The mail arrives on their servers right? So you have to ask for the new mails. In this case you can use webhook, but that is not we want. We can have webhooks, but the point is not all mail providers might provide that right? So we need to create a abstraction beyond a webhook.
  - I call it “collecting mailboxes”, like you will have to daily or some particular or repeated interval go to the mailbox and check if any mails are there for us. I think that is a design decision and not a single way to do things.
  - I learnt a bit about [webhooks](https://www.svix.com/resources/webhook-university/fundamentals/what-is-a-webhook/) in the process.

## Tech News

- [Levent Alpoge with Claude claim to have solved Riemann Hypothesis](https://alpo.ge/remark.pdf)
- [The new Docker branding](https://www.docker.com/blog/meet-the-new-docker-built-to-run-on-trust/)
- [Fly.io introduces the free plan](https://fly.io/changelog/introducing-the-free-plan/)

* * *

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/720) (#820th edition), and for software development/coding articles, join daily.dev.

That’s it from the 121st Edition of techstructive weekly. I hope you found it helpful, and relaxing. If not please drop any suggestions, feedback or discussion about certain things you want to in the comments or drop me a message on my [socials](https://www.meetgor.com/contact).

Thank you for reading,

Until next week.

Happy Coding :)
