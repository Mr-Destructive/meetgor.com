{
    "author":"meet",
    "date":"2025-10-11",
    "post_dir":"newsletter",
    "status":"published",
    "title": "Techstructive Weekly #63",
    "type":"newsletter",
    "slug": "techstructive-weekly-63"
}


## Week #63

It was a fun week. After a long, a long I have had the time to rip-off a code base hands on and do something beyond pushing a feature, it was about reforming how a thing is approached. I felt really good. There are always rough days, but these days, the days that give you pain but that pain is satisfying. At the end of each day of the week, I had gone and came out of bed brimming full of ideas and questions to talk and sit through. Felt I was back to the grind.

I also did a 2 hour-ish live stream. I finally started the project of creating a unified interface for LLM APIs in Golang. Felt satisfied and rejoiced in love of programming with bare hands.

Looking forward for a better weekend, a relaxing and busy week ahead of Diwali.

### Quote of the week

> “A world is a fine place and worth the fighting for”
> 
> — [Ernest Hemingway, For Whom the Bell Tolls](https://www.goodreads.com/quotes/2120666-the-world-is-a-fine-place-and-worth-the-fighting)
> 
> I read it in the “Days at the Torunka Cafe”

This quote, just came at the right time, the right week. What a luck! The world is a fine place and worth fighting for. I had worse weeks, but the past two weeks have been very gentle and rewarding to me. I am blessed to have a job, a great and supportive family, what else a man needs? Rest all is materialistic shine.

You live and breathe for the happiness of others, and those others none other than the one that give you the support in bad days, so you have the strength to live in the world that is worth fighting for.

## Created

Finally a [livestream after ages](https://www.youtube.com/live/fJS2Crq-8qU). Started implementation of Gollum, a unified interface for LLM APIs in Golang. Let’s see how I can wrap it up.

- [Creating LLM API interface in Golang](https://youtu.be/fJS2Crq-8qU)

Double click to interact with video

## Read

1. [Development gets better with age](https://www.allthingsdistributed.com/2025/10/better-with-age.html)
    - This one was really soothing one.
      > The older developer isn’t worried about the barrage of new model announcements and feature releases that come out every week. He’s seen that before. New tech, same patterns.
    - And this one
      > Have an in-depth conversation with your customer, listen, dive deep into their challenges, suggest architectures, migrations, and tools. And sometimes, the solution will be generative AI.
2. [Asynchronous Work is the enemy of project based success](https://www.arturhenriques.com/p/asynchronous-work-is-the-enemy-of)
    - Maybe this is true. Collaboration and Communication is key, I don’t think its about being asynchronous from each other. Its about being able to flexibly decide that, give enough brain space to understand and sit with the problem, to let developers cook (as in think and do their own stuff).
3. [Give me AI slop over human sludge any day](https://world.hey.com/dhh/give-me-ai-slop-over-human-sludge-any-day-8c4b747d)
    - True. I can see this over X, Youtube, people create cringe worthy content. AI content yet lifeless, can said to be informational the least. That also has cringe element but a little lesser I think.
4. [Why we need junior engineers](https://www.infoworld.com/article/4065771/why-we-need-junior-developers.html)
    - Yes, there is no other way round. The fresh perspective is really needed to a new revolution in technology. The existing mindset, the mental model might not allow the freedom and might offer a bit of resistance to adoption and leverage of those tools.
5. [How I learned to stop worrying and started trusting and using AI Agents](https://metalbear.com/blog/claude-experience/)
    - Yes, this is true. A year worth time in 2025, we can say, AI Agents or AI-assisted coding has become really powerful and steerable.
    - I also was skeptical on those AI-vibe coding apps. But then came in Cursor and took us by storm. WIndsurf, and what not VS code clones took over.
6. [My approach to building large technical projects - Mitchel Hashimoto](https://mitchellh.com/writing/building-large-technical-projects)
    - Agreed to this. It is truly valuable. Have something to make you go ahead, let that good dopamine hits. If you are doing the backend, add unit test, print stuff out to see the progress.
    - Yes, sometimes experience hurts. And having the perfectionist mindset might be a little harsh.
7. [Python splitlines does a lot more than new lines](https://yossarian.net/til/post/python-s-splitlines-does-a-lot-more-than-just-newlines/)
    - Wow, I knew this, but realising it again as I forgot how split and splitlines is like a magic function. You never know you need that level of magic at times, but sometimes you do.
    - This is something I love and hate about python.

## Watched

- MCP was a mistake
    - Yes it is a mistake. LLMs were launching left and right. People barely had the time to respond and understand LLMs, that these AI labs were launching tool calling, image generation, thinking and what not.
    - But then came a unified way to expose the tools, but it added more complexity than decreasing. Context rot and what not. People started adding more MCP servers, assuming that it would do more stuff and correct stuff. Which is wrong, the more information you give someone, the more it confuses itself.

- Theo ranks every vibe coding app
    - Its kind of a good take. Vibes. I really didn’t like Bolt, maybe because I was not using Chrome? But it just kept on failing. Haven’t tried V0, I had tried the very first version of V0, it was very frotnend focused. Those were some magical days.
    - Gemini and Qwen are good too, who can’t afford Claude and expensive models. We are in experimental phase, all are launching their CLI equivalents of the models. HH! What a time to be alive.


## Learnt

- Using [dc as a mathematical utility in linux](https://en.wikipedia.org/wiki/Dc_(computer_program))
    - dc or desk calculator, a very inappropriate name.
    - It is a calculator utility in linux. Its pretty basic and has only 20 odd commands, but it can do any maths operations. Apt very apt for a unix philosophy, do one thing very precisely and correctly.
- Reviewing AI generated code is no joke
    - I produced AI slop for a feature, or rather improvement feature.
    - It was so terrible or glazzing that I had to launch a v2. But then after 3 days spending my head banging at cursor chat window and letting it rip it off and analysing it with bare hands.
    - It took some time, but it had created something meaningful after all. Was it worth it, maybe. Would I have taken the same time? At the best case, yes, no maybe. But that is not the best part of LLM vibe coding.
    - It is about the rest of the stuff, the benchmarking, the rip off why things work, grilling it, questioning its own produced slop, against each other. Producing throwaway test and scripts to test its own garbage.
    - That is so quick, the feedback is almost instant feedback. I am genuinely blessed and grateful to have privileged to be in a company to be able to use such productive tools like Cursor, Gemini, and others.
- PDF Forensics
    - For creating the mentioned features, I had to dig deep, read a lot of slop to find the gems of actual features.
    - I discovered that we can detect suspicious pdfs with embedded files in metadata or the actual pdf content. We can even use certain techniques to detect overlaying text in the content of the pdf with the location.

## Tech News

- [Build Agents and workflows with OpenAI Agents](https://openai.com/index/introducing-agentkit/)
    - How many startups, ideas and products killed. People create products and open ai comes with their own platform ready for it. Except for the API, none of their products get used that way like killing off the thing I think.

It was great week. Looking for a next week which is Diwali weekend. Preparation for the festival of light will be underway. The cleaning of homes, shops, lighting, lights, new beginning, sweets, family and the nostalgia of memories of childhood.

See you next week.

---

Happy Coding :)

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-766) (#766th edition) , and for software development/coding articles, join [daily.dev](http://daily.dev/).
