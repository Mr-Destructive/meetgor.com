---
title: "Techstructive Weekly #118"
date: 2026-11-06
slug: techstructive-weekly-118
type: newsletter
status: published
source: newsletter
canonical_url: https://techstructively.substack.com/p/techstructive-weekly-118
description: "Reading and watching the software industry fall into security apocalypse, among the other things learnt and observed in the week from 1st to 7th November 2026 (a perfect start month)"
tags: ["newsletter", "substack"]
---

## Week #118

It was a normal week. Worked with agents to get improvements and fixes for things. I will be quick in this one, its festive time. Diwali weekend. I was busy with the cleaning and dusting and other chores the previous week, this week it will be the actual celebration, nothing much, but expecting a cheerful time with the family and friends.

This week on work side, I learned to have the high level overview of all the apis and software I am working and touching, it might be a bit hard in some cases. But it needs to be there, if we are not writing code, we must know what is being shipped and what currently exists.

I also wrote one essay around ~1.7k words. I might keep writing and publish them in a single burst.

### Quote of the Week

> “Be as a bird perched on a frail branch that she feels bending beneath her, still she sings away all the same, knowing she has wings.“
> 
> — [Victor Hugo](https://www.goodreads.com/quotes/117240-be-a-bird-perched-on-a-frail-branch-that-she)

What a beautiful quote. You know you are dealing with a problem under your actions, but you still do your best, knowing that you have the choice to move away at any time. You have the wings, the legs to fly away. If you are stuck in a situation, you can’t just cry and lament about it, you have to give your best, since you always have the chance to move away from it, not by quitting but taking a conscious decision that the problem is not for me.

## Created

- [Story seed generator](https://claude.ai/share/f03ccc0f-c802-408f-8bf0-98134dfa2066): https://claude.ai/share/f03ccc0f-c802-408f-8bf0-98134dfa2066
  
  - This was vibe coded, one shot on claude, was too lazy and occupied to publish it as a webpage somewhere, I haven’t actually made a pipeline that will turn my artifacts into a collection just like tools.simonwilison has.

## Read

1. [Don’t call it a decision Model: TypeSafe AI on Jev](https://typesafe.ai/blog/dont-call-it-a-decision-model)
   
   1. This is a interesting blog from the jev creator’s themselves, stating that Jev is not a decision model, its a model that can support a decision or create the direction for the eventual decision in a logic of choices. That’s a fair balanced take.
   2. The next question then I have is how does it fit with the agentic development while building systems. Does it make the system more stochastic if its paried with decisions or does it improve in all the cases? Or in a nutshell, is there a slice of software that jev or any judgemental model can make obsolete with its choices?
2. [I struggle to believe that AI R&D is ingenous](https://www.lesswrong.com/posts/DALNiKBjMGMF8grme/arjun-panickssery-s-shortform?commentId=fJtaHHHfx45tD7BSL)
   
   1. True. Models can once in a while do something ingenious but the problem is you are trying to solve is too small for it then. Always, LLMs need to be given a ambitious goal, but that fails catastrophically if not paired with right tools, or reward/review mechanisms I think. If it can’t verify, If you can’t verify if you’ve build something workable or intended then how can it be done?
   2. It almost always need that silver lining of human experience and domain expertise to be able to steer and create a wow moment for you to be effective working with LLM Agents right?

## Watched

- [CS50 Show Season 1 Episode 2: PHP-ocalypse, Git and GitHub](https://youtu.be/C_UwDG2JKBQ)
  
  - This was interesting to see the growth of AI-driven attack and vulnerability finding. Some bugs were obscure to find, but due to AI-Agents, everything is possible. I don’t get the point of not-using AI-Agents for defending. Like what’s the point in resisting the weapon that you are going to get shot down with eventually?
  - Also the git and github thing, I thought it was about how GitHub is losing the crown. Though it was something different, infact it felt it was suggesting to understand git more. Hmmm? Maybe we need to think again on how software is version controlled. Since Git is becoming kind of a blocker for agents and co-ordination.

<!--THE END-->

- [VPN is just NAT](https://youtu.be/MCbMko4IXKw)
  
  - Beautiful explanation. How the network works, it makes it so nicely understandable. The packing up of http request, then encrypting it, then the network layer and all of that opens up for exploration.
  - The private and the public network concept was so interesting to understand. Like the router is like a gateway that other devices in the network reach for when they want to communicate with other network devices.
  - And of course the concept of VPN is so dangerous, now I look at it. Wow! They have the power to interpret anything right? Like passwords, personal details and whatever we are consuming or browsing. That is a bit security concern, you never know when these VPN bros start taking and selling your data.

<!--THE END-->

- [The private key mistake: TLS and the number 4?](https://youtu.be/NDgS__J1U-M?)
  
  - Damm! the number 4 appears again. The part that they were able to generate numbers like 4,156, and others did bug me in the initial reaction too. Like how the heck their system generated that small of a number in particularly less than 1000 all 5-6 times? If that is the case then that system is not random right? Not even pseudo random I would say.

<!--THE END-->

- [Agents sucked at UI until now?](https://youtu.be/xq0eB8vxeqk)
  
  - Hill climbing. I think I am too poor to be able to use all of those models. Like I still don’t get how can one justify the 20$ or 100$ plan for personal use? Like how do you guys have that much money to spend each month? If you can generate some revenue then its justified like for theo or some people, but for others?
  - It feels a bit exciting that these things are possible but access is becoming a barrier to be able to spend that much money into those frontier models.
  - The loop engineering is becoming true, but there are nuances in how it is percieved and steered. Like right skills to rate it, then work towards the goal (not that goal but the eventual intention) are making the progress along with model alignment to make impossible task just a matter of prompts now.

<!--THE END-->

- [The easy way to learn nothing](https://youtu.be/wYk1IpGEBIk)
  
  - Yes, this is true and relatable. Once we get the drug of AI, learning feels like a obligation to do, almost like a hurdle or a forced thing to do. But what will this turn out humans in? Brainless puppet masters?

<!--THE END-->

- [Who is richer? Sonic or Mario](https://youtu.be/vCAD3W43lIk): Numberphile Game fictional comparison of game character earnings
  
  - This was a interesting comparison. Though some things were just biases or assumptions from Tom. I didn’t quite like the 50% of the coins like nobody plays that badly, right, come on!
  - Mario is richer than Sonic just because he has full solid coins and sonic has rings which have holes and make it less worthy of collecting. Interesting observation. Despite having more copies sold, he lost. Math is interesting.

## Tech News

- [Canva laysoff 1200 people](https://www.canva.com/newsroom/news/canva-team-update/)
- [OpenAI launches GPT Images 2.6](https://openai.com/index/introducing-images-2-6/)

* * *

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/817) (#817th edition), and for software development/coding articles, join daily.dev.

That’s it from the 118th Edition of techstructive weekly. I hope you found it helpful, and relaxing. If not please drop any suggestions, feedback or discussion about certain things you want to in the comments or drop me a message on my [socials](https://www.meetgor.com/contact).

Thank you for reading,

Until next week.

Happy Coding :)

Also

Happy Diwali! May your life be filled with joy and love. Wishing you and your family a happy and cheerful Diwali!
