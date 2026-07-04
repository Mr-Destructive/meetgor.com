---
title: "Techstructive Weekly #101"
date: 2026-07-03
slug: techstructive-weekly-101
type: newsletter
status: published
source: newsletter
canonical_url: https://techstructively.substack.com/p/techstructive-weekly-101
description: "A code-heavy week. Writing about the things learnt, read and watched over the week from 28th June to 4th July 2026"
tags: ["newsletter", "substack"]
---

## Week #101

I will be quick this week. I had barely anytime to read or watch content. I was back in code. Deep not careless as people are pretending to be. I am not writing loops or vibe coding into production. I am using chat mode in Cursor to spawning parallel agents to solve issues that I find it cumbersome for me to tackle. I don’t skim through the suggested code, I read line by line, understand the flow. Write some code myself, to understand what is actually happening and to make myself realize the complexity.

I feel alive again as a software developer. I think this is life, some weeks make you feel down, some week make you excited to get up and smash hands on the keyboards. This week has been the later and I can’t be grateful enough.

It’s likely I would carry this momentum forward in the week. Do a livestream to explore some specific thing to keep the kindling flicker of curiosity alive and afloat myself in the dreaded hype of AI. I will read some and write some essays or novels to freshen my mind. Last week I wrote a 3k word essay and I feel great. It was from my memories, it felt great to move it off my chest. I would do the same this week too. Saturday for streaming and Sunday for writing my heart out.

### Quote of the Week

> “The woods are lovely, dark and deep, but I have promises to keep and miles before I go to sleep.”
> 
> — [Robert Frost, Stopping by Woods on a Snowy Evening](https://www.goodreads.com/quotes/2377-these-woods-are-lovely-dark-and-deep-but-i-have)

This is something I read from the daily quotes. And it just made me smile. Just beautiful right. The path is dark and deep, but I have to walk anyway. That is how it is. The path appears gloomy and dark, but we often and most often come as a better person when we go through it. That is the way nature operates. You have to change in order to be happy, that is a state, to transition into that, you need to shed yourself from the current state. Change is the only constant. We’ll walk it through anyways.

## Created

- HTTP Server in Golang from (almost scratch)
  
  - As I said last week, I wanted to build a http server in golang from scratch and I did it.
  - It was not from scratch but it was doing deep from the high-level. I enjoyed it. It was refreshing and I was able to understand the http request-response parsing. The next step is to understand the socket file descriptor inner working I think.

## Read

1. Manu Martínez-Almeida Blog: [Building Go-Gin: Simple over Easy](https://manualmeida.dev/articles/gin-simple-over-easy/)
   
   1. Ok this is really good adivce. Building things which are simple over easy. There is a subtle difference between simple and easy. Something is simple to understand like http semantics, working with apis on postman feels easy, and python is easy to learn, but golang is a simple language to understand. There is a difference, yet very fine line is blurred for most of us.
   2. When I say easy, I mean the effort required to do something. Python looks easy on the outside, but when dealing with intricate dict mutation and internals, the easiness is no longer the ground. Its no longer simple, it was never simple, it made it look easy. Not saying Python is a bad language, but its not something you want if you are looking for lower abstraction.
   3. When I say simple, it means its requires very little or almost no cognition or effort to understand something or do it. When I look at a Golang codebase, it is not easy to understand it, there is a lot of code, but its simple. There is a difference. The size of things matter yes, but when everything is dead simple, it sometimes or most often is verbose. More verbosity doesn’t necessarily mean simple. Java is another example when I hear the word verbose. But what I am saying it simple is intuitive. Easy is a effort thing.
2. Carson Gross Blog: [University in the AI Era](https://htmx.org/essays/universities-and-ai/#automate-everything)
   
   1. This is very grounded and honest take on AI. I have hardly seen a professor express and address the concerns about AI. Really appreciated effort.
   2. This statement that University are in a great position to distance from AI with the physical interaction making it harder to AI access is good, but not sure till when that might hold true. If robotics sees a boom like text based LLMs we are not months away from creating a mess.
   3. The increasing of ambitious homework is the right and exciting move. I want to see myself how would I put effort to get to the goal. Its right that professors accept students will use AI and they intentionally increase the level of things they can produce.
   4. Homework is no longer a signal that builds on the above point. I like that, because I could see myself skimming over the homework to do something more interesting maybe? As I always have been to do programming.
   5. Visualizations are cheap. This is a underrated feature of AI, people are not trained enough to produce good quality visualizations. Maybe its moving too fast but there should be some ground where people are trained to create these helpful visualizations of things they want to learn to form a deeper understanding of a concept.
   6. I just love the take that hand coding is still a valiant thing and thought-after thing. Yes, it looks obsolete right now. AI agents can produce 40K lines of code in a day or more. But the human behind it needs to understand the value of those lines before judging the quality of it. Coding and programming are still valuable and not obsolete. Just that you need some edge in other domain to get the most leverage out of it. Now just knowing how to invert a bnary tree is not helping anyone, people need to understand the problem, the domain in order to apply and solve problems first.
3. [The daily brief by Zerodha: A deep dive into the textile industry](https://thedailybrief.zerodha.com/p/a-deep-dive-into-the-textile-industry)
   
   1. What a great read. It shows what the textile industry is, what the past, the current and the future looks like.
   2. The part that india is behind in the most profitable chunk of the process is so hurting. Also its a bit concerning that we are stagnant just because nobody wants to get out of the comfort zone. Its a huge risk, I understand, it might put end or stake someones livelihood out of the normal. But does that stability stall the progress of the nation? Its a bit debatable but India has surely gotten a bit stalled on this I think.
   3. The question of finding the next billion jobs is a bigger one. Its not easy to answer that, and predict what could come next. Does AI affect this in any way, does it make it worse or better. I don’t see any way it would make it better.
4. SpyGlass : [Natural Born Bloggers](https://spyglass.org/om/)
   
   1. A heartfelt blog in the memory of Om Malik. I am not aware of this person, but whoever it is, he is immortal due to his writing (at least for sometime for someone). I think writing is a great act you can do.
   2. I like that kind of personality. being able to communicate and make relatable analogies for people as a communication process.
   3. I cannot relate but deeply fathom that it would be hard to hear and discuss opinions of someone who you considered valuable.
5. Dan Kinksy: [HackerRank opensourced their ATS, the results are quite inconsistent](https://danunparsed.com/p/hackerrank-open-source-ats)
   
   1. This is hilarious. Why would a company like hackerrank make such a software. I think it was vibe coded? Maybe. who knows.
   2. But building AI-systems is not just prompt, its more about the layers above and below it. It needs to deterministic, otherwise its not called a software. ATS is a software, and making it random would break the point of the software, either call it ATS or just call it a slot machine.
6. Dan’s Webnotes: [CSS Tricks for markdown blogs](https://dan-webnotes.com/posts/2026-06-28-css-tricks-markdown-blogs/)
   
   1. This is so neat and useful. I want to build these things on my blog too. The blog itself is so aesthetically pleasing. I like minimal style.
7. Igor Kulman’s blog: I [do not feel like a programmer anymore](https://blog.kulman.sk/i-do-not-feel-like-a-programmer-anymore/)
   
   1. People are actually more brave then I thought. This is true. A very honest, open hearted blog.
   2. Tuning prompts and reviewing code is not the same as solving a deeper problem and sitting on it enough to get the bulb glowing. The moments of costly dopamine (opposite of cheap dopamine?) is getting harder to get in coding. Writing prompts, seeing AI solve problems, and guessing the outcomes doesn’t give any motivating feelings that would move us forward.
   3. This all cycle is so messed up in the software world. People are re-inventing how we develop software. I don’t know what is going to happen, but sooner or later, the gold old days should come back that’s my wish. Otherwise I don’t like it very much, I don’t hate it always but its a bit hard-pilled medicine to digest.

## Watched

- [Mario Zecher on Pi: I hated every coding agent, So I built my own](https://youtu.be/Dli5slNaJu0)
  
  - This is a nice little opinion, and he built a non-opinionated ai agent. he didn’t like the way other harnesses modified and obscured the context. He built one that does what it shows. Simple.
  - I haven’t used PI to be honest, because I don’t have a paid api key to work with. Thinking of using it with a local small open llm model.

<!--THE END-->

- [Corey Schafer: Python AsyncIO, a complete guide to asynchronous programming with visualisations](https://youtu.be/oAkLSJNr5zY)
  
  - Just beautiful. Amazing stuff. i didn’t knew asynchronous means to pause and do other stuff while we wait for IO. I thought it was executing things in parallel.
  - Now I get the things, Its still requires me to build projects to fully understand it, but the idea is clear now. Its not magic anymore. The visualisations were really well made and made everything feel understandable instantly.
  - I think I am going to take some time to digest this. Run loop, the tasks, the threads and the groups all need to be understood in a slow way. the person making this video made a great effort in showing what async is in python. But on the other end, I also need to put in effort to understand it deeply.

## Learnt

- PostgreSQL: you cannot use the aggregate function column with the having clause
  
  - SQLite allows it. PostgreSQL doesn’t. Wired.
  - basically you cannot say `select sum(price) as total from transactions group by order having total > 50`
  - You need to use the aggregate function again like so `select sum(price) as total from transactions group by order having sum(price) > 50`
- Kubectl on GCP
  
  - kubectl get pods gives the actual running instances (change on restart).
  - kubectl get deploy gives us the template + replica count that owns those pods

## Tech News

- [Fable is now allowed by US government](https://www.anthropic.com/news/redeploying-fable-5)
- [Anthropic releases Sonnet 5](https://www.google.com/url?sa=t&source=web&rct=j&opi=89978449&url=https%3A%2F%2Fplatform.claude.com%2Fdocs%2Fen%2Fabout-claude%2Fmodels%2Fwhats-new-sonnet-5&ved=2ahUKEwjex42O57aVAxVzmFYBHSOqO_kQFnoECBwQAQ&usg=AOvVaw0hqaSLqa5NbxOI_SGp4Dz6)
- [Google releases Nano Banana Lite](https://www.google.com/url?sa=t&source=web&rct=j&opi=89978449&url=https%3A%2F%2Fcloud.google.com%2Fblog%2Fproducts%2Fai-machine-learning%2Fnano-banana-2-lite-and-gemini-omni-flash-available&ved=2ahUKEwibpKiG57aVAxXMlFYBHUEsEIsQFnoECBwQAQ&usg=AOvVaw3k-oY8bvxIfnwif8mZkvIF)

* * *

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/800) (#800th edition), and for software development/coding articles, join daily.dev.

That’s it from the 100th Edition of techstructive weekly. I hope you found it helpful, and relaxing. If not please drop any suggestions, feedback or discussion about certain things you want to in the comments or drop me a message on my [socials](https://www.meetgor.com/contact).

Thank you for reading,

Until next week.

Happy Coding :)
