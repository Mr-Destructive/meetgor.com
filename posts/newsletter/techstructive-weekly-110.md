---
title: "Techstructive Weekly #110"
date: 2026-09-11
slug: techstructive-weekly-110
type: newsletter
status: published
source: newsletter
canonical_url: https://techstructively.substack.com/p/techstructive-weekly-110
description: "Created free llm router, annoyed with gemini models, among the other things read, watched and learnt in the week from 6th to 12th September 2026"
tags: ["newsletter", "substack"]
---

## Week #110

It was a hammering week. 3k lines shipped to prod, no issues. Agentic coding on the move. I am not sure, I should be proud or not, but not bragging about it, it was a win for me. The code was obviously reviewed from me :| and no one else even spent an second of eye on it, shipped, nothing broke. It was a model updation + some tracing and standardization of model routing. There were 2 major things at once, if you are updating models, that could lead to unknowns that you even don’t know, and that’s what happened only 3 instances among 3k requests in the week. The other being tracing and model standardization library addition. It worked in the end, there are improvements to be done, there always are. But a fair learning after recovering from sick 2 week leave. I call it a comeback motivation from nature. Thanking and grateful for it.

On to the week work, it was good work, I now like working with agents, there is a learning involved, if you are curious and like exploring codebases, its a great tool. I mostly use claude and occasional cursor, I am designing apis or building deeper integrations. I would like to explore more things in the weekend, but it gets a bit tiring. So might read some books and chill out.

### Quote of the Week

> “I thought there were no more ghosts there than those of absence and loss”
> 
> “And for a moment I thought there were no more ghosts there than those of absence and loss, and that the light that smiled on me was borrowed light, only real as long as I could hold it in my eyes, second by second.”
> 
> ― [Carlos Ruiz Zafón](https://www.goodreads.com/quotes/6648408-and-for-a-moment-i-thought-there-were-no-more) in The Shadow of the Wind

What a beautiful quote that hits timely. There are no ghosts, after a age in childhood, you don’t fear darkness anymore, as grow from child to adult, you start fearing the loneliness and absence of someone. For me there is a fear of losing my mother and father, everyone has, but its more scarier than my own death, because these two are the only one that can truly understand me, without me saying anything. Light is something they have provided me constantly and its absence is what brings me fear.

## Created

- [Free LLM Router](https://github.com/mr-Destructive/free-llm-router)
  
  - I discussed it last week and here is the repo today. I did manage to get it working after gpt not willing to implement the api, opencode to the rescue and it created it correctly with the freely available llms on the web
  - It uses llms like gemini, kilo, zen, ovh and llm7 providers. I know they will get blocked in some time, but its worth keeping a repo of a quick check on how somethings can work, just a free call without having the auth or a paywall.
  - Yes of course share private info in these chats, it might be used for training, but these days people are rarely concerned about security and privacy right?
  - This week I’ll try to deploy it on cloudflare or vercel cloud functions. Need to first verify if those environments will block the request, my suspicions is that they will.

## Read

1. [Native is the future of mobile at Shopify](https://shopify.engineering/back-to-native)
   
   1. This is cool, using ai to review the screen to make incremental changes. Helix is a nice idea to port one system to other, in this case React Native to Kotlin or native mobile application port.
   2. The premise of software development has changed and the technology is also changing, because it does still mater on which technology is used to make it a pleasant experience to users.
   3. The code can be generated at fly, so generating code is not the issue, maintaining two repos is also not a big deal but porting or migrating is.

## Watched

- [AGI is here, and its bad right?](https://youtu.be/4B4R2T4w7Kg)
  
  - This is really bad. The person working at Anthropic is saying it, that they are working irresponsibly.
  - If they know its wrong direction, why even move, you know I think its about the “power” comes with “ego” and ego dissolves the ability to think. Maybe they have become too arrogant to address the elephant in the room.
  - They can make AGI but what the price humans will pay for that? Losing the soul? Devastating the whole earth? Who knows, we can be optimistic but for what reason?

<!--THE END-->

- [You should stop pretending you understand the codebase](https://youtu.be/5KvY8CnBB3w)
  
  - You cannot know everything about the codebase. You can know something but not everything. If you think you know everything then you are probably wrong, or you are not shipping enough, or thinking and dreaming big enough.
  - Agents make it really easy to make the change but the context to make that is really critical.
  - The other thing that caught me was that you cannot re-create codebases from agents unless you have the context and the mental model of the existing one, truly a great piece of thought to keep in mind.

<!--THE END-->

- [You are the project](https://youtu.be/8faAKQeFkgk)
  
  - The project is not the point, the process of transforming yourself into someone you love to be is.
  - banger of a video, shows that AI cannot replace the satisfaction and the development of skill that the friction or the process of going through the actual act of doing it.

## Learnt

- Gemini models have moved from thinking budget to thinking tiers and its annoying as hell
  
  - This was a shock for me, I had kept 5k for some thing and expected it to work for over a year, and now some model is being deprecated and it sucks to switch from token count to levels like low,medium and high where there is no constraint and adherence to the level.
  - This is a bad strategy from Google, I don’t like it, there is no restriction on what low or high means, and don’t even select medium that is the worse of all.

## Tech News

- [DeepSeek releases V4.1 Flash](https://www.deepseek.com/en/news/deepseek-v4-1-flash/)
- [Tailwind labs is joining Shopify](https://tailwindcss.com/blog/tailwind-is-joining-shopify)
- [Meta releases Muse agent on the cloud](https://ai.meta.com/muse/) (only in US though)

* * *

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/809) (#809th edition), and for software development/coding articles, join daily.dev.

That’s it from the 110th Edition of techstructive weekly. I hope you found it helpful, and relaxing. If not please drop any suggestions, feedback or discussion about certain things you want to in the comments or drop me a message on my [socials](https://www.meetgor.com/contact).

Thank you for reading,

Until next week.

Happy Coding :)
