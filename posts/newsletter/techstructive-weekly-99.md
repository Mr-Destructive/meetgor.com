---
title: "Techstructive Weekly #99"
tags: ["newsletter", "substack"]
hash: aa488f2a8756c0f55fb2804ca032fbfe72433980befab10ed177b3dc03491d47
date: 2026-06-19
slug: techstructive-weekly-99
type: newsletter
status: published
source: newsletter
canonical_url: https://techstructively.substack.com/p/techstructive-weekly-99
description: "Reading slowly, trying to find the learning rhythm back, understanding systems and interacting with users among the other things watched and learnt in the week of 14th to 20th June 2026"
---
## Week #99

What a week. It was a great week not gonna lie. Finally seeing the light of the day after 1 month of grind. Being able to see the satisfaction of a user is a different ball game. Speed is something I couldn’t deal with, it becomes overwhelming to manage. But sometimes its good to push yourselves to get something. Maybe nobody will appreciate or notice at that time, but you will know what worked for you, how can you adapt next time and learn the lessons.

I am feeling the energy now to explore things and just keep doing things once I found exciting back in 2020-2023 era. Somehow I lost that overtime. Maybe I can blame AI, but I should blame here myself for getting carried away with the hype and living in the expectation that AI will eventually make developers obsolete. Or I can put it as there won’t be a need to learn anything or something specific deeply. I was wrong. You don’t learn to compete with AI, you learn to be confident in yourself, the moat is not what you know, its what can you do with that, what patterns you can find that LLM can’t comprehend. Maybe its time to get back into learning mode. Forgetting for a moment that LLM can do it for me.

I also read this article “[Reading is not a competition](https://medium.com/@mentalgarden/reading-is-not-a-competition-thats-why-i-left-goodreads-4637bac30cdd)”, it was a quick read, the author was to the point and added his own personal experiences which makes it authentic and worth reading. I relate to this. reading should be a personal activity and not a gamification of numbers. People spend time reading to get more books done, but that is not the point. The point is to enjoy them. I read a bunch of books, not sure how many but in May-June I read at least 3 I think, one of them was “Kafka on the Shore” that book just hits hard. It is so deep and magical at the same time. I will continue to read more books just for fun, I do have some number in reading challenge, but I don’t care If I complete it or not. It is just for me to keep reminding me to keep going through in life, keep up the reading habit. The habit itself heals me.

### Quote of the Week

> “It’s like Tolstoy said. Happiness is an allegory, unhappiness a story.”
> 
> [― Haruki Murakami, Kafka on the Shore](https://www.goodreads.com/quotes/1977-it-s-like-tolstoy-said-happiness-is-an-allegory-unhappiness-a)

Yeah! Happiness is just in dreams something abstract. But unhappiness quietly becomes a story. How relatable and rightly put. In the context of the story this was so hard hitting. It feels right to use that analogy. Happiness for me at least is not a story maybe, it becomes a journey, and unhappiness is the start of it, it eventually lands maybe but unhappiness is a complete story. Either you failed or you missed it, if you are lucky you get success, for the rest of us it remains a dream something that we can’t feel, almost distant.

* * *

## Read

1. [Grey Beam Blog: DuckDB Internals - Part 1](https://www.greybeam.ai/blog/duckdb-internals-part-1)
   
   1. A very balanced article. Not too shallow, nor too deep. It was the right thing for me to get through. I understood the parts involved in the query parsing and storing the database. There are a lot of things done by the storage engine too, just to make the rest of the query analysis and execution process easier and efficient.
   2. I didn’t knew parquet was a column based storage, I thought it was some sort of a compression format. It also stores zones which help in retrieval. DuckDB leverages the strengths of Parquet files on its own architecture and the result is SQLite like nativeness with speed and more features.
   3. I am keen on reading the other 2 articles in the series. As I have used duckdb on my airspace case study, I want to know how it was able to pull off the huge mamoth data but sqlite was slow.
2. [AI Creativity and Discovery - Thoughts by Richard Sutton](https://x.com/RichardSSutton/status/2061216087744946656)
   
   - A very calm and no fuss article / tweet. It was to the point and calming, nothing like “developers have lost their moat” and whatever glazing people do these days, it was like a peaceful write up.
   - It gave me a thought, a freeing thought if you will, “AIs are just mimic” It was said by him, I am not quoting him for it, but that is so freeing and a well put in words. People think this is AGI and all jobs are replaceable and what not. But if we want to create something novel and good, AI cannot do both which is right. It can write good code, also novel code sometimes, but not at the same time. It cannot write something that is out of the ordinary, it can produce customized software, basic ones, like applying one thing pattern to other. That is solved, its grunt work even for human I can say that. But beyond that, it cannot really do much. It cannot create the 10008 RFC Query method spec for humans to follow. Humans still need to be in the process, the steering drivers.
   - Just feeling good after reading this. I am very positive on AI as a tool. But sometimes, people just go too far with that idea and leave everything on auto-pilot and panic mode. Humans still needs to be in the weeds. This kind of write-up is a example that I keep on reading stuff. Every once in a while I get this kind of words to read that makes me feel alive and charged back.

<!--THE END-->

03. [Mark Nottingham Blog: You want to define well-known URI](https://mnot.net/blog/2026/well_known_uris)
    
    1. This is some great advice.
    2. Don't use it as a shortcut for a real URL, If your protocol can transmit a full URL, transmit the full URL.
04. [Gitignore isn’t the only place to ignore files in git](https://nelson.cloud/.gitignore-isnt-the-only-way-to-ignore-files-in-git/)
    
    1. Wow! This was clever. I didn’t knew you could ignore a previously tracked file from the git without mentioning it in the gitignore.
    2. There are two more places you could do that, those look internal but gets the job done without leaking on the public
    3. `.git/info/exclude:`This is for the local git repo, you add the filename and it ignores the file. Pretty neat.
    4. `~/.config/git/ignore`: This is for global context. If you say .DS\_STORE or whatever mac has, it will not track in any git repo of your device.
05. [RFC 10008: The HTTP Query Method](https://www.rfc-editor.org/info/rfc10008/)
    
    1. This was proposed in 2019s, and got accepted in 2026, damm, we are still evolving HTTP. And some people claim “Software is solved(mostly)”.
    2. I like this approach of taking things time, protocols are something you can’t vibe code or generate a spec using plan mode. It needs discussions, sleeping on it, thinking through things, and a lot more thinking then the currrent LLMs context limits.
06. [American Express Blog on Cell based Architecture for payment processing](https://americanexpress.io/cell-based-architecture-for-resilient-payment-systems/)
    
    1. This feels one of those articles where its quite hard to understand the details. It like doing a high level design interview but just completed after the basic diagram, no edge case discussion no schema or examples.
    2. I am not saying this is not resourceful. Its a great piece of information, maybe not for my level of understanding. Maybe its too high level or just enough to give someone a hint about designing a broader system.
    3. I like some points though, which require some digesting
       
       > “for mission-critical systems like payments, we find that the benefits of a reduced blast radius and improved resiliency outweigh the additional complexity”
    4. This is so true, the trade off is there and rightly so, the payment processing comes over in priority over the complexity.
07. [Introducing the pkg.go dev API](https://go.dev/blog/pkgsite-api)
    
    1. This is a good entry into the ecosystem. Go team knows there is a need to expose the packages in a ai-agent-friendly way. And this seems to be a right move.
    2. By exposing this over api, an ai agent can know what the package’s structure or core features are without needing to download it. its debatable how well it is documented. But if it is then its a powerful way to use it for agents.
    3. I would like to explore this with the open api specification. Right now that looks very bare bones.
08. [Roman Imankulov Blog: A backdoor in a Linkedin Job offer](https://roman.pt/posts/linkedin-backdoor/)
    
    1. Damm! GitHub borrowed identity. You can do that? Like what? You can take my username and stick to some other repos’ commit! Whew!
    2. But that was a clever way to check it. Readonly tools and PI agent. Wow! I haven’t tried PI maybe its time to combine the PI and the 3B model by cohere locally.
    3. Security is a nightmare these days, AI has made it easier to create software, but at a cost of introducing more bugs and loopholes and scripts like that. Reviewing software is not a bottleneck, its a mandatory thing to do and might be the only thing to do before breaking into the prod.
09. [What job interviews thought me about Kubernetes](https://notnotp.com/notes/what-job-interviews-taught-me-about-kubernetes/)
    
    1. This is the trend which looks a bit dangerous and safer too. If every one is on K8s then its fine, but if there are people on other things then it becomes like the pascal or fortran thing, very few developers know the way out.
    2. This is also partly due to AI maybe. People get same suggestions and they accept it without second thought. Maybe there might be a few years where software development has stalled. Like no new frameworks or ways of working with things, just rewiring existing technology or frameworks to make the things moving.
10. [Marek Suppa’s BLog: Making HTTP requests from a container that has no curl, using bash /dev/tcp](https://mareksuppa.com/til/bash-dev-tcp-http-without-curl/)
    
    1. This is cool. Bash is some magical compiler that knows a little too much to be used right?
    2. It cannot read the request just parse it up so that it can send the request. It feels like everything was designed so perfectly right? Imagine an AI in this container, will it be able to get the request across?

## Watched

- [Theo Reacts: AI makes junior engineers less harmful](https://youtu.be/rTMRlqT8Q8c)
  
  - This is a hot take really. Its a hard pill to digest, a bitter lesson. Some people get bad with AI, they don’t learn more. I think I am falling into that territory. I need to up my game and actively learn new stuff. But again, this kind of is too much of a pace to keep up, what exactly to learn, what depth is enough, is just too much blurry these days. You can call it excuses but honestly for me its a battle to navigate this time.
  - Judgement and Escalation of work is something will be the key in the coming years maybe, but gathering experience and being in the deep tech will be helpful, no matter what the situation becomes.

<!--THE END-->

- [The Primeagen: I am done with Golang](https://youtu.be/WqSWZuGS9pc)
  
  - Yep, the iterators just bugged me. I was not able to wrap the head around why was it needed, when there was just one way to loop, for for everything, simple.
  - I have used go for maybe more than 3 years now, I do still like it, but the features they are adding, its like making it bloated and a bit confusing and looking like a another generic programming language. Its not that bad, (C++ sigh! its way way better then that, maybe its out of reach of that) but I do hope that it remains as simple as it is now. Adding all swiss army knives might be breaking the appreciation and the bond that developers have over the language.

<!--THE END-->

- [The Twin Prime conjecture](https://youtu.be/8HBDE-msUjw)
  
  - This is so fascinating. It was impossible to it, but somehow it was done.
  - Everything seems impossible until its done. Another quote to take away is “not knowing somethings are better”. Because in some cases, you need your mind to wander in that uncertainty for it to find new pathways and push beyond the comfort zone.
  - I also learnt about the sieve of erotosthenes. It is quite simple right, but people just make it complicated. It is one of the examples, of how simple things can lead to broad understanding of unrelated things.

Non tech video but helpful for indian audience on mathematics

- Basic Mathematics: Episode 1 by Anand Kumar in Super Infinity
  
  - This is so underrated. Like I know this sounds too basic, but look at how enthusiastically he explains those simple concepts, look at the relation he draws from each concept to other. Its just a beautiful thing to watch.
  - Fundamentals are going to be the thing that makes or break the youth at any era, AI or not, you need to know some things regardless of how smart things are around you.

## Learnt

- The query method in HTTP which was released in the RFC 10008 this month has been in discussion since 8 years.
  
  - This query method bridges the gap of complicated GET request for filtering by not using POST method.
  - Like if we have some query to filter in a get request, the query parameter can be used but there are limitations and it becomes a bit hard to decipher. Also some browsers or websites discard that piece of information that might be needed in the query.
  - And using POST method, is a bit conter-intuitive since you are not creating or modifying resources on the server, you are just filtering them right. People use POST just to get the request body.
  - QUERY method solves this, by adding a content type and based on that the request body is communicated to the server or the client. Really making the right balance in complex filtering options.
  - I am looking forward to building something with it over the weekend on a livestream.

## Tech News

- [Anthropic’s Flagship model “Fable” was banned by the US Government](https://www.anthropic.com/news/fable-mythos-access)
  
  - This lasted barely a day and shut down. Just like the name fable is a short lived story, the model did what it was meant to. Might be a marketing stunt again from Anthropic to get attention.
- [SpaceX to acquire Cursor](https://www.reuters.com/legal/transactional/spacex-buy-anysphere-60-billion-2026-06-16/)
  
  - This was coming and it looks closer. I am not sure what to expect out of it. The composer models are quite fast I can say, a bit expensive yes, but the speed is just next tier of its own.
  - If they get right compute engine, they might have the best moat.
- [Z.ai releases GLM 5.2](https://z.ai/blog/glm-5.2):
  
  - According to [artificial analysis it](https://artificialanalysis.ai/articles/glm-5-2-is-the-new-leading-open-weights-model-on-the-artificial-analysis-intelligence-index) is the leading open weights model.
- [Cohere releases their first ever developer model: North Mini Code](https://cohere.com/blog/north-mini-code)
  
  - This is pretty neat, usable and open source. 3B is manageable I think on any low-end laptops too. Especially people like me with only 8GB of RAM.

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/798) (#798th edition), and for software development/coding articles, join daily.dev.

* * *

That’s it from this 100th newsletter its 99th edition starting from 0, so 100 newsletters written. That feels unreal. 2 years gone!

Looking forward to a learning weekend and slowing things down by reading and getting mental peace. Being a developer in 2026 is easy, but being a good and effective developer in 2026 is exhaustive and not a piece of cake.

With that thought, I’ll conclude here.

See you in the next one.

Happy Coding :)