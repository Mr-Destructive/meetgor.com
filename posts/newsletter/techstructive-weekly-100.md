---
canonical_url: https://techstructively.substack.com/p/techstructive-weekly-100
tags: ["newsletter", "substack"]
date: 2026-06-26
slug: techstructive-weekly-100
status: published
source: newsletter
description: "A first century edition of the newsletter for the things created read, watched, and explored from the week 21st to 27th June 2026"
hash: c50acfb121b207286b9110127f6a329ab5e43432b7750d5c6448c9d11d60b08b
title: "Techstructive Weekly #100"
type: newsletter
---
## Week #100

Well! This feels unreal. 100 weeks have pased since I started writing this newsletter. I am still writing them. Unbroken saturday streak. Feels like time travel.

2 years is a lot of time. I don’t know what revolution we are witnessing after 2024, but since I started this newsletter, tech space has just fundamentally changed due to LLMs first, then AI-Agents and now Agentic Engineering. What comes next is what I am trying to figure out, how to work with them and how to be a better developer. That is why this was newsletter was created and it continues to do so.

I get a chance to feel myself in the loop by reading about people’s thoughts, their expereinces. But the best part is reflection. The amount of things I am able to retain and learn again from these posts are so much valuable. I am not just consuming the artilces or videos I reflect on what it gave me to think for, which direction I was wandering and where it left me. These questions are not something people or I can keep in my head while working in tech. its too overwhelming to do so. So every end of week, I try to dump out my thoughts in order to start again and align in a better direction to be a better developer or a human atleast.

So this another week in that, and I continue on this journey to explore and just tinker with stuff just becuase there was a curiosity to do so.

### Quote of the Week

> “Bird by bird, buddy. Just take it bird by bird”
> 
> — [Anne Lammott](https://www.goodreads.com/quotes/371038-thirty-years-ago-my-older-brother-who-was-ten-years) in Bird By Bird: Some instructions on writing and life

Bird by bird, one day at a time, next move best move (Sagar from Chess base india for the prop). You don’t have to live your life in a day, you don’t have to write the book in a day. You have to do one word after other, one paragraph other, one page , one chapter after other. That is how it goes. That is how our mind can keep up otherwise it will throw you off in random directions. Speed is not important, direction is, once you get direction, the speed is something comes naturally.

Just like this newsletter, one week at a time. It grew to 100 editions. Over two years. One article, one youtube video watched, one question, one bit of discovery led to 100s and thousands of links being explored. I created this web of link blog and I cannot thank myself enough for keeping up with the overwhelming amount of tech revolution. I kept on hanging, it was overwhelming but I still remember the joy of fixing a bug on my own on a computer windows 7 being in college 11th standard. It gives me the joy of solving problems and exploring. That will never leave me, one search at a time, one thing at a time. Don’t stress yourself, you are capable of doing anything just take the right amount of task at a time. Token by token like large lang… Damm that anaolgy.

## Created

- [HTTP Query Method](https://www.meetgor.com/posts/http-query-method/)
  
  - I wrote this blog to make myself clear on what the Query method is actually solving. And it was right. It serves as a bridging gap for limitation of GET request query parameters for filtering/searching/querying data and the unintuitiveness of using POST method for querying or searching data.
  - I wrote the problem first, and then built the intuition on why the HTTP Query method was proposed. It rightfully makes sense to have a separate method verb for obtaining or filtering data.
- [Live Stream: HTTP Query Method in Golang](https://www.youtube.com/live/5Bv83Rb5Qjo?si=qgNN894xkKP0jMEg)
  
  - I also did what I said last week on this newsletter to start streaming again. And I loved it. It was annoying with the laptop crashing on saturday, but then I limited myself to few resources while streaming and it worked perfectly. I completed a basic implementation of query method using the net/http standard library for basic filtering. It was hardcoded to filter certain values only, but that is not the problem statement of this. The idea of this stream was to understand how QUERY method plays the role differently from GET and POST and how headers from OPTIONS help in linking the requirements to the final query.
  - Looking forward to doing a couple of more streams on writing http server from scratch in Golang this weekend.

## Read

01. Jim’s Blog: [Blogging can just be stating the obvious](https://blog.jim-nielsen.com/2026/blogging-stating-the-obvious/)
    
    1. It might sound obvious at that moment for someone else, but for me if its quite complex to put thought into something then I write about it.
    2. Just like I wrote the Query method this week. It might be obvious that the query method does querying of data right? But why a separate method when we have GET? Think for a moment and you will find it intriguing to write about it once you figure that out in words.
    3. I might find this article like what am I talking about? That’s quite obvious now, it makes so little sense to write it, but to get it out of my brain help me crystalize what it was about. Writing blog is about connecting ideas, bringing them together and figuring them out as we learn.
02. [Mark Round’s Blog: Its only when you look back after 25 years on a blog](https://www.markround.com/blog/2026/06/17/25-its-only-when-you-look-back/)
    
    1. Wow! He has written everything he grew up in. I really the time, when sharing things was novel. But now, people look at it with a suspicion, is it ai slop? The value for writing honestly is more than ever.
    2. Writing is easier to produce but human sentiment is not. Human level understanding and intuition is a bit on a stretch. We need humans to make humans understand things. Maybe it might not be the case in few years, who knows.
    3. But this blog deserves praise. 25 years of consistent blogging whilst keeping up with his hobbies and job is not something everyone is able to do. And having the courage to say see you in 25 years is commendable.
    4. The change looks really bizare, like in 90s there were rams of hardly kbs. Now we struggle even having 10s of GBs of ram. Phew! Nothing is never enough. The blog running on PHP and what not was really relatable. I started my blog with jekyll on github pages, moved to python markata still on github pages and now its on custom go ssg on github pages with other apis and projects on cloudflare and vercel.
    5. I have written articles over 5 years now. I dont want to stop. I am writting again, after a year of blocking by hands due to AI hype and existential dread, but I have hopes now. I just care for myself and 1 human atleast who would find my writing joyful, maybe its a ai agent at some time, I don’t care even if nobody reads my blogs, I will have a repertoire of content as a archive of my growth just to myself.
03. Akash Tandon Blog: [Debugging in the age of ai agents](https://www.akashtandon.in/ai/2026-06-24-debugging-in-the-age-of-agents/)
    
    1. This is a bit hard and situational. Hallucination is a very thin line between the truth and the evidence. if you have designed the system to produce some context that might delude the agent into something misleading, it will hallucinate. This is not a ai problem but a human engineering problem.
    2. The other problem is genuine which it gives a confident wrong analysis. Its almost like a fantasy story it came up with, the moment we say “are you sure”, the story disappears and the agent wakes up to crash and produce more false evidences.
    3. Having a general understanding and the direction in a bug is so underated. If you can identify the problem, narrow it down to a certain scope or module, then the chances of Agent going haywire reduce, not saying it goes to 0, but you can manually then hop in and see for yourself. Programming is outsourced to AI agents, debugging still cannot be.
04. The Guardian: [Indian Factory workers told to film themselves while working](https://www.theguardian.com/global-development/2026/jun/24/indian-factory-workers-told-film-themselves-for-ai-robots)
    
    1. This is rough! I though only software development process was easily solvable with AI since we exactly know what we want (actually we dont at all, but we pretend we do). Software is testable, we can check if its right or wrong. But for manual work, its possible too but quite hard to verify it.
    2. That is why labs are gathering data now. Text generation has now hit a ceilink. We can only do certain things for certain tasks with text models and tool calling. This is now getting understood and there is a need to get image and video footage of replicating the same with robots in mechanics.
    3. I don’t like it, but seems inevitable. The industrial revolution might be nothing compared to this. This might give people a serious existential dread.
05. Vinicius’ Blog: [When I reject AI code even if it works](https://vinibrasil.com/when-i-reject-ai-code-even-if-it-works/)
    
    1. Yes, exactly. When YOU write the code, you have the context, you have the details, you know why exactly there were 2 parameters and not one, why it was abstracted from the module, why the code was written the way it was. But when we offload that to an agent, it has the context. No matter how much you read the ask mode answer, you still have lost the connection to the code, you have connection but scattered or not fully formed.
    2. This is might not seem like a issue to have confidence in the code at first when merging, but when things start to break over time, as they do, it becomes harder and harder to reason through it. The compounding effect of your lose connection to the code and the pile of changes from the Agent which might be unwanted and trash. We software developers have to pay that cognition debt to resolve it, its a trade off people are thinking coding is solved. It might not be just yet. The compounding will take some time, some of them have already realised it, some are about to hit the storm.
    3. This paragpraph sums it so well
       
       > Yes, coding agents can help you with this task with more than just writing code, but that doesn’t mean they can do it autonomously in a sustainable manner *yet*.
    4. They cannot just yet.
06. PostHog’s Blog: [Robbie Coomber wrote a 70x faster SQL parser while barely looking at the code (agentic coding and not vibe coding)](https://posthog.com/blog/sql-parser)
    
    1. OK, hot takes.This is a great article. It proves that engineering is still a need to get most out of ai agents. It wouldn’t have been possible without the author knowing how the sql parser works in existing library ANTLR and a other approach to do it efficiently.
    2. Technically I dont know how each of the two approaches work, but from reading it meant two approaches are opposite ends, one is faster and only for specific set of general operations, the other is more careful and gives a general purpose full parser maybe. Both the approaches worked for him, but he choose the faster specific use case one. This shows how easy it is to generate a working code, but maintaining becomes a question. I case of parsing a SQL query, this looks simple and without any maintenance cost I think. But in general cases, AI generated code is not magic and will need some before or after thought.
07. Stephen Smith’s Blog: [Linux kills strncpy](https://smist08.wordpress.com/2026/06/25/linux-kills-strncpy/)
    
    1. Oh! That is nasty. I didn’t knew strncpy was that dangerous. Copying data from the stack was a gigachad move from hackers. being a hacker is so underrated, like it takes some thinking to break code, though the direction is wrong so its not worth it.
    2. But the other point is that it took 5-6 years to remove this from the kernel. This shows how hard it could get if some code remains in the system for a long time. This is legacy code, if we are vibe coding things, ai generated code will at some point become this, and the problem is not that it becomes hard to remove things, but the knowledge of where the code should be and should not be removed is also gone. Engineers might not know the code anymore, so it remains on the whim of AIs to generate and clean up things. Adding the complexity.
08. CSVBase Blog: [Parquet is more than Turbo CSV](https://csvbase.com/blog/3)
    
    1. This was so cool. Last week I. read about the duckdb and its engine. I was intrigued by the design of duckdb like parquet. And this week this answers that question. Seems like I am sync with the universe, its giving me hints.
    2. So, parquet is column based storage but also row based chunking. So streaming is not possible if you need intermediate data. Since its broken down into columns, its not possible to stream half of it, either you get full or nothing. The metadata thing is so cool. it has sorted the different types of data, encodings, null and even time stamps and zones (god has anyone said time? what time). having rich metadata is not a tradeoff its a design that makes parquet really great for large data storage.
    3. It reduces the size to almost 50% from csv that is just bonkers. Imagine what it would be from json to parquet that I did and felt for aircraft flight observatory project a couple of months back. Its quite evident, parquet is a great format for storage.
09. [Pydantic takes a great initiative to take httpx maintenance with httpx2](https://github.com/pydantic/httpx2)
    
    1. This is a genuinely a great gesture from Pydantic, this means they truly value human and open source. This shows the care for development of such libraries. In the world of darkness, there are people still lighting up candles to show the path to the lost.
    2. Really looking forward to this project, I want to see why is it so great. Honestly I haven’t used httpx, I have only used requests. Maybe I have touched a bit of it in contributing to posting [here](https://github.com/darrenburns/posting/pull/231/changes)? I think its a good abstraction for working with http requests and responses, but why is it different from requests? is it opinionated?
10. Timon Web’s Blog: [Wagtail as Django Admin on Steroids](https://timonweb.com/wagtail/wagtail-as-django-admin-on-steroids/)
    
    1. This is true. And it shows how a niche thing can become a business. This was executed brilliantly. No doubt.
    2. Wagtail is just a wrapper around django admin, that’s it , the UI and the configuration. Nothing is magical, yet people use it to save the time and money from doing it themselves. Perfectly valid and genuine use case for a business to thrive.
    3. I am thinking is there anything to build like that?
11. Plik Blog: [Read more fiction](https://pilk.website/5/read-more-fiction)
    
    1. I dont need this advice. I relate to this so much. I already read non-fiction all the time, emails, technical reading the thing I read for my blog here, etc. That’s why I want to escape into a world. Escapism, copium, whatever you can call it. It heals me, I don’t care if you call me weak.
    2. I read, I feel good. I feel I can live my life, I feel I am not alone. When I write, there is the page, the blank editor ready to hear me out.

## Watched

- [The Pragmatic Engineer Gergely’s Talk: Slow down and Speed up](https://youtu.be/5wks1W-auKY)
  
  - ouch! Facebook sorry Meta had a bug in their AI chatbot? Was it that capable of sending useful data? Really? Was it intentional to get attention that their AI is somewhat capable? I didn’t knew this bug in instagram happened? Phew! Vibe coding is actually happening at the best engineering firm.
  - Token maxing was a thing yes, now people are facing the issue of cost for certainly. The AI bit is sitll a constant, its not going anywhere, its like a atium in the mistborn, everyone wants unlimited amount of tokens but nobody has the money to buy that. It will go down as another tool, I surely think, it wont be 100% developer replacing thing.
  - This video is a great showing just that, people will still be around figuring stuff out, so we need to slow down and brace for what is going on and about to happen.

<!--THE END-->

- [AI Loop Engineering take from Neetcode](https://youtu.be/vM6DNlsdpsg)
  
  - We needed this desparately. A great great courage by Neetcode I forgot his name(I think Navdeep right?) . But the point being, cutting the hype. People are posting contradictory things as evident from the video. Nobody knows the correct and reproducible way to work with LLMs that is the thing people and the labs advocating are sugarcoating.
  - I know its a great tool, I like to use it, but there’s a limit in what people should talk about and not spread degeneracy and misleading hopes of being 100x and being left behind. That is atrocious and not something I like to live with.
  - Thanks Navdeep for this post. We genuinely need content like this from others too. That actually glaze over the hype and spit out facts.

<!--THE END-->

- [Sentdex: There has been a situation in AI](https://youtu.be/DA2E8szXjoA)
  
  - OK, People are waking up. This is true. People dont want to vendor locked in. This is like cloud lock. But a far more muffled and hideous.
  - This guy is a legend, and mostly has thought many python developers of today the basics of ML in python. An absolute legend of pre-LLM era. Open source models are getting seriously close, like there is a understandable gap, but people can build meaningful prototypes from it. And that is already halved the cost not less right?
  - I have no idea about GLM 5.2 but if its this guy saying it, then hands down it should be. Cannot test it, but will try to check with openrouter if available for free tier.

<!--THE END-->

- [Daniel Hirsch Streaming: Coding an HTTP Server in C](https://youtu.be/WmZi1yrOD0Q)
  
  - This was a great learning experience. I wanted to understand what a socket connection really does in a http server. So its basically a file descriptor just like pipe, socket, or terminal descriptior in unix. That is just like a file streaming. That makes sense. The interface is unified for accessing the network and the file.
  - I feel like building this in Golang. Maybe this weekend it will be a stream to build a http server from scratch in Golang. I want to see how we connect to sockets and listen and then accpet and then process the request and response to the client. It feels great to just be able to do things low level.

<!--THE END-->

- [Tsoding Streams: Let’s talk about AppImage](https://youtu.be/sRGT90fqh4g)
  
  - Ok, this was good stream. It showed how ridiculous documentation can get. People are trying to sell things, but when the buyer is already sold, the user want to know how to use it, and AppImage developers did a bad job making that process easy. It was obsfucated from making obvious to create an appimage than it was to run a app image.
  - Maybe AppImages are great but they still have shared dependencies from the os level so its not quite isolated right?

<!--THE END-->

- [Whom does ChatGPT serve? Jonathan Pageau talks](https://youtu.be/yZUuKzDQSsI)
  
  - A great talk! I dont know the story but it felt relatable the urge to listen to someone and act blindly in instinct like animal instinct is increasing with the usage of AI.
  - We need to understand what we are doing and on what motivation are we doing it. The fundamental of action is being eroded by these AI text chunking machines.
  - Great points by the speaker here, we humans contradict ourselves when our mind is not aligned, like he gave examples of Elon Musk saying AI as existential risk to humans initially but then creating his own AI. WOW! What a full cricle here.
  - Must watch to get the sentiment of this topic, its a heavy topic to grasp and cannot be comprehended here.

## Learnt

- HEAD Method is not used for getting supported or accepted request format in http
  
  - I thought for using QUERY, I need to first send OPTIONS to checks if QUERY method is supported and if so, then send HEAD to get the Accept-Query header to check which format is supported for requests and then send the QUERY method finally.
  - But that was not it, We can directly send the Query method. The options and head are for discoverability and not like required to send it obviously I was just showing the right fundamental way of exploring APIs with the new Query method.
- Django filters: chained filter vs multiple fields in a filter
  
  - This one was interesting. I learned that chaining `.filter()` calls and combining multiple conditions in one `.filter()` are equivalent in many cases, they both produce a bit of situational sql depending on the relation in the fields passed to it.
- How to build a http server from scratch in C
  
  - Instead of relying on a high-level HTTP library, I learned how an HTTP server starts with sockets, accepts TCP connections, reads raw bytes, parses the HTTP request manually, and sends back a correctly formatted HTTP response.
- How to parse numbers from strings in C
- Messaging apps scrub the exif data of images before sharing

## Tech News

- [Anthropic releases Claude Tag on Slack](https://www.anthropic.com/news/introducing-claude-tag)
  
  - Ok, didn’t they release this like bot earlier? I thought it was a joke but then Karapathy’s tweet was seen and I was like what? Then I realized this is a instance of claude code (kind of) in a slack channel. I am not sure about this.
  - The amount of wasteful content in a channel might pollute the context and make it do random things. So, not sure it would be the right thing all the time, but can see it being useful in isolated specfic conversations.
- [Notion shuts down Notion Email after a year](https://www.neowin.net/news/notion-is-shutting-down-its-email-client-one-year-after-launch/)
  
  - I remember or not maybe, it was Notion who said, we know mail better then you, so we launched [Notion mail.](https://www.notion.com/releases/2025-04-15) Was it notion or someone else I am not sure, but this is contradictory.

* * *

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/799) (#799th edition), and for software development/coding articles, join daily.dev.

That’s it from the 100th Edition of techstructive weekly. I hope you found it helpful, and relaxing. If not please drop any suggestions, feedback or discussion about certain things you want to in the comments or drop me a message on my [socials](https://www.meetgor.com/contact).

Thank you for reading,

Until next week.

Happy Coding :)

[Leave a comment](https://techstructively.substack.com/p/techstructive-weekly-100/comments)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](https://techstructively.substack.com/p/techstructive-weekly-100?utm_source=substack&utm_medium=email&utm_content=share&action=share)

[Subscribe now](https://techstructively.substack.com/subscribe?)