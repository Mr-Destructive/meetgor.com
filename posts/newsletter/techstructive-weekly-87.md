---
title: "Techstructive Weekly #87"
date: 2026-03-27
slug: techstructive-weekly-87
type: newsletter
status: published
source: newsletter
canonical_url: https://techstructively.substack.com/p/techstructive-weekly-87
description: "Building Flight Observatory, crunching data, learning SQL, reading about LLMs among the other things read, watched, learnt and created in the week from 22nd to 28th March 2026"
tags: ["newsletter", "substack"]
---

## Week #87

It was an active week. Did quite a lot of stuff. Most of which might be wasted, but lessons learnt. Pulling datasets in GBs from the internet is not a easy task, especially if its in chunks of JSON.

I completed CS50 SQL course and created the final project linked in the created section. It was fun, learnt a lot. I had watched the full video series a few months back. But never got the full time to complete the problem sets and the project. Locked in the past weekend, and completed the full 7 problem sets.

### Flight Observatory Project

I also worked on creating and making the flight observatory. Its a bit of a challenge, which I want to complete by hook or by crook. It is a dataset issue, the data is there in samples but that too is quite too much to pull off. I cannot find a easy way to grab all the data for Mumbai only region. Loading full is not feasible and cannot manage it locally. So, loading each file in memory and filtering based on positions and saving the relevant bits into compact CSV is the best I have gotten, but that is still a hard challenge to get 17k requests per day and that will be 12 per year, I want a decade of data, so that makes 17k\*12\*10 requests. Not a small number to deal with. Its not rate limiting but processing that much is quite time consuming. Even parallelism won’t help here.

Anyways…

Will check over the weekend if there is a possibility to pull it off quickly and cleanly.

### Quote of the week

> “One machine can do the work of fifty ordinary men. No machine can do the work of one extraordinary man. ”
> 
> ― [Elbert Hubbard](https://www.goodreads.com/quotes/112927-one-machine-can-do-the-work-of-fifty-ordinary-men)

This hits home for the week. I had Codex, amp, and Gemini run for me for the flight observatory. But none of them could do a single thing without my instincts and pulling the cords. It was almost as if it can know a lot of stuff but cannot move a penny, only when steered in the right direction by the right person. Like an aeroplane in the hands of a pilot. It cannot do it properly without the steering and right guidance, the right mindset, and experience. Be extraordinary, don’t let AI limit you, let it push you.

* * *

## Created

- [Case Study for Mumbai Airport with ADS-B Historical Data](https://dev.meetgor.com/flight-observatory/case-study/mumbai/)
  
  - This is half baked, and made an AI-mistake that it took 900 second snapshot when the data was available for 5 seconds. Just because I told you I want the data to be quickly downloaded.
  - Have spent rather a lot of time wasted in this downloading part, I want to analyse the data, but its quite hard to pull it off.
- EPL Dataset Exploration for CS50 SQL Final Project
  
  - [Demonstration Video](https://youtu.be/JwxuWOYDkw4)
  - This was a great side quest for the final project. All done easily with a simple dataset. Downloaded the dataset of matches of EPL seasons from 1994 till 2025 and showed them to a sqlite databsae.
  - I got a great deal of info, who won the most epl titles in the past 3 decades. Which team scored most goals, and the silver lining of all was the validation query that I ran “which team never lost a game in a season”, gave the perfect answer. Arsenal, the invincibles.
- [Tags and a separate RSS Feed on the blog](https://www.meetgor.com/feeds/)
  
  - I wanted a single layer for my other places like s3g to pull my content from a reliable and non-blocking place. So this was much needed.
  - Now you can fetch tag wise, post type wise and even all content feeds.

## Read

01. [A eulogy for Vim](https://drewdevault.com/2026/03/25/2026-03-25-Forking-vim.html)
    
    - A really heartfelt post. I was not expecting AI, but yes, it matters.
    - Bram Moolenaar, an absolute legend, the creator of Vim. He died a couple of years back. The author of this post just wrote his heart out. Every humanly possible connection is relatable (not for me for Bram, but yes can see that). He was kind, he supported poor people in Uganda.
    - This post just gives me a hope that people can live still. AI will probably eat the world, but people like him will cease to be eaten up. And that is liberating, that is moving thought. People need to be aware what is happening in the world.
    - Vim Classic, a tribute to Bram, really great stuff. Will support if I can with any power.
02. [DatabaseMaxing with Preston Thrope](https://pthorpe92.dev/databasemaxxing/)
    
    - This is a great guide. Inspirational and eye opener. I wanted to do it, year back. But still procrastinating. I had written about learning SQL, but now in the phase of building low level stuff like interpreter and compilers.
    - I am thinking of building a markdown parser like pydantic for the web. I know that is wired but it can open up a lot of possibilities.
03. [Yes, I am bored of reading and listening about AI too](https://blog.jakesaunders.dev/is-anybody-else-bored-of-talking-about-ai/)
    
    - I hear you. Managers and literally everyone are somehow shoehorned into software development for asking how many tokens they are using.
    - Do they even know what a token is? Do they know wha a database engine and storage is, their difference? The difference between a row and a column storage nuances? Why? Why not ask those, instead of obsessing over the metrics. Its a wired and frustrating time to be a developer in one perspective.
    - Yeah! I know its all great and wonderful time to be a developer as well. But the bad parts just suck out the joy of it. The analogy of discussing what tool to use and how is just infuriating. Just show what you build and how, not with what and nerd sniping the markdown file specifications.
04. [So where are all the ai apps?](https://www.answer.ai/posts/2026-03-12-so-where-are-all-the-ai-apps.html)
    
    - Good questions and modest exploration. I don’t think the increase in productivity of building of software has any relation to packages being developed. Since the first trend that we would see is that people will make software for themselves. Solo software or personal software as many have called it, apps are just that.
    - Yes, that comparison of packages being shipped is a good point but not fully true and valid. Its kind of very early (oh, we are 3 years into AI now). But people are still figuring it out.
    - We say what just happened with LiteLLM on PyPI. That might be a nail in a coffin for shipping pacakges. Or even installing them.
05. [To live in a world without AI](https://pype.dev/to-live-in-a-world-without-ai/)
    
    - We have a flurry of posts like these. Good to see that. It gives hope that there are people who think what I fear is happening.
    - The thought from [waylonwalker](https://thoughts.waylonwalker.com/post/954) is truly a good callout. I hope we stay on the thinking loops rather than sitback and let clankers do the rug pull.
    - I am really startled that, this AI dystopia has already arrived and if it goes any further the world looks like a bad world from a fictional book. Phew! Turn of a year or two and bamm into a dystopia of robots.
06. [Every Kubernetes concept has a story](https://x.com/livingdevops/status/2037430761150984475)
    
    - A beautiful post. I was confused midway, should I laugh or learn?
    - I don’t think using a database in a container is a good idea in any situation. Unless you have a session-only requirement.
    - The pains of one thing are actually a concept in Kubernetes that is a good lesson to learn and learn each concept with a situation in mind.
07. [Markdown ate the world](https://matduggan.com/markdown-ate-the-world/)
    
    - This is really intriguing, I though markdown was old than docx. But its not it was there after 2004, that is quite a recent thing. Just over two decades. And it has changed the way people write.
    - I think majority of the non-tech people still use docx and whatever doc format is for Microsoft Word. Its just works you know. And if it doesn’t work it just doesn’t work. Nothing in between. But markdown always works. Anyways, who can convince them. Microsoft has a deep foot in the minds of people.
    - With LLMs, markdown just became a standard now. Everyone knows it, its simple with just enough structure to separate it out from plain text. Just the right balance but dead simple.
08. [Why I vibe in Go and not in Python and Rust](https://lifelog.my/episode/why-i-vibe-in-go-not-rust-or-python)
    
    - I can feel this might be written by AI, but still it has a valid point. Go doesn’t gets in the way. Python never does (only in prod), Rust just blocks you.
09. [The machines didn’t take your craft](https://www.davidabram.dev/musings/the-machine-didnt-take-your-craft/)
    
    - A good one, the last paragraph hits home.
      
      > No matter what tools arrive, no matter how powerful they become, they will always remain tools. They won’t replace our reason nor values. You will still choose what is worth building. And as long you reason, nothing essential has been lost.
    - This is a good call to not abandon your craft, just because a faster method exist, never attach yourself to the tools. They will keep on changing.
10. [The diminished art of coding](https://nolanlawson.com/2026/03/22/the-diminished-art-of-coding/)
    
    - Very good. Really good. Its written from heart and humans value this. This is what I call writing. Writing not for attention, writing for connection.
      
      > If you’ve never taken an interest in poetry, or painting, or dance, or whatever, now would be a good time. In an era where the internet is increasingly full of bots pumping their bland bot ideas into everybody’s brains, seeking out distinctly human forms of expression has become vital
    - A good point. In this era, connecting yourself with nature or art as we say is vital. Especially for developers, we need to get out of our heads sometimes, we need to ponder the blank, the boredom.
11. [You are not your job](https://jry.io/writing/you-are-not-your-job/)
    
    - Gold. Just sheer empathy and gratitude for the person writing.
      
      > - “You evaluate someone's warmth first to gauge their intent before ability”
      > - “The people who love you don't love you because you're good at your job. They love you because of something else entirely. Maybe it's your humor. Maybe it's that you actually listen. Maybe it's that you remember things about their lives and ask about them. Maybe it's simply that you show up. You're present. You don't extract a conversation and then disappear.”
      > - “The harder version is asking yourself: if my job title disappeared tomorrow, would I still be me? Would the people who matter still love me? If the answer is yes, you’re in the right place.
      >   
      >   If the answer is no - if your identity is not cleanly separated from what you do for money - your relationship to yourself may need an update.
      >   
      >   You are not your job. You’re a person first. Your ability to connect, be present, and make people feel understood is what makes you irreplaceable to the people around you, which is the only market that counts.”
    - bangers after bangers. Really, I am framing this post. Its like a healing potion. A compass for life in tough times, such as the current ones.
12. [Somethings just take time](https://lucumr.pocoo.org/2026/3/20/some-things-just-take-time/)
    
    - Yes, some things do take time, and friction is what helps it grow. Rightly said.
    - I like the idea of building communities and trust, that is what I had done without having extensive technical skills at my job, I was there just showed up daily, and … it ends wiredly but not all time ends up the right way. Maybe it was my mistake, but anyways, moving on to a new chapter.
    - I don’t know why everyone wants to churn software, what is the hurry, do you know what you want to build?
13. [Ctrl + C in psql gives me the heebi-jeebies](https://neon.com/blog/ctrl-c-in-psql-gives-me-the-heebie-jeebies)
    
    - This is really interesting. I like the way he calls it heebi-jeebies. It really is.
    - Like the TLS is not there for the cancel request, so your psql connection sends the unencrypted database secret in the wild, and somehow if intercepted by anyone in the same network, it can launch a Denial of Service attack.
    - The Neon Proxy and Elephant shark(the wireshark but for Postgres) have a workaround by noting the secret with the initial connection and when the psql sends it with the plain text the secret it intercepts it and kills the right session. Wired stuff but kind of no choice, that would require a bit of a refactor on the protocol.

## Watched

- [Concurrency Patterns in Golang](https://youtu.be/rDRa23k70CU)
  
  - This was a great video explaining go routines and concurrency patterns. Loved the analogy of gophers as senders and receivers as passing the buckets in channels, the buffered and unbuffered channels as the gophers in between, really well explained.
  - Also, the patterns and concepts for those concurrent go routines and channels were well explained.

<!--THE END-->

- [A bad day to use python](https://youtu.be/mx3g7XoPVNQ)
  
  - Wired, how can a pypi credentials be comprimised that too for a founder and that leads to a release without the developers realising it?
  - Very weird, litellm was a great package, I like what they do, I want to make that for Golang. Never able to make one. Thankful that I was not the one with Golang yet.

<!--THE END-->

- [OSI and TCP Best Explanation](https://youtu.be/3b_TAYtzuho):
  
  - A very enthusiastic and clear explanation of the IP Model.
  - The difference is really explained really elegantly. The confusion and the separation of layers is pitched right way, removing all the why so questions.
  - Highly recommend watching it to get a good grasp on the fundamentals of the Network Model.

<!--THE END-->

- [Composer 2 and Cursor Drama on Kimi](https://youtu.be/QGnKTRtEH50)
  
  - Damm! This is a bit sad, but not bad. Like Kimi is happy with the parternship, Cursor played around the limitation. Which is not a good thing, and can lead to other major closed source labs to follow this trend which would make open source models like a free fruit to grab.
  - Not a good spirit, but cannot say anything about it if both of them are happy.

<!--THE END-->

- [Kafka and RabbitMQ differences and uses](https://youtu.be/1HOVtQ-_fcE)
  
  - A good explanation of Kafka and RabbitMQ, the difference was there I didn’t knew it. They both are message brokers but one is a smart and other is a storage place with other operations to do things around it.
  - I still don’t know which one to use when, its time to really use them

## Learnt

- Downloading a dataset individually from raw files is harder than downloading with S3 or R2 files with RClone
- Always match concurrency to DB capacity.
- Use buffered channels to spawn goroutines.
- Early blocking reduces wasted memory and load on DB.
- Usage of Keep-Alive header in HTTP Requests
  
  - TCP Connections Are Expensive. Every HTTP request over TCP (or HTTPS over TLS) starts with a 3-way handshake, SYN where client asks to open a connection, then SYN-ACK where the server acknowledges, and then ACK where the client confirms.If HTTPS, we also have a TLS handshake, exchanging keys, certificates, negotiating ciphers. Each handshake takes milliseconds, adds CPU cycles, and adds network round-trips.
  - When we add the HTTP header `Connection: keep-alive`the TCP connection stays open for multiple requests between client and server.
    
    Backend (or proxy) maintains a connection pool, reusing the same socket.
    
    Reducing the Latency (no repeated handshake) and the CPU overhead.

### Interesting Tidbits

- https://flighty.com/airports
  
  - A great UI for viewing map and the traffic.
- https://github.com/ssrajadh/sentrysearch
  
  - Interesting idea for semantic search based on embeddings.
  - Gemini can understand the semantics of a video from its embedding model and we can search parts of videos. This is really cool, and this is where AI expands into real world use cases.
- https://translate.kagi.com/
  
  - This is hillarious. It translates sentence to Linkedin like voice. Amazing, whoever made it.

## Tech News

- [LiteLLM package compromised](https://github.com/BerriAI/litellm/issues/24512)
  
  - Phew! That is a wild gun. PyPI attacks have been quite a lot in the past few years. There needs to be some layer of security with it. The ecosystem needs to improve.

<!--THE END-->

- [Bye Bye Sora, did anyone use it?](https://x.com/soraofficialapp/status/2036532795984715896)
  
  - Yeah! Like who actually used it? I just saw bunch of ghibli trends and generic slop on linkedin and twooter. That’s it nothing useful.
  - Disney ended the deal, which means either it was not great, or it was too expensive for them. The former is more likely.

* * *

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/787) (#787th edition), and for software development/coding articles, join [daily.dev](http://daily.dev/).

That’s it from this week, it was quite exploratory week. Learnt and built a lot of stuff without actually shipping, but that is where real learning happens I think. Lets ship full guns blazing next week.

Happy Coding :)

[Leave a comment](https://techstructively.substack.com/p/techstructive-weekly-87/comments)

[Subscribe now](https://techstructively.substack.com/subscribe?)

[Share](https://techstructively.substack.com/p/techstructive-weekly-87?utm_source=substack&utm_medium=email&utm_content=share&action=share)
