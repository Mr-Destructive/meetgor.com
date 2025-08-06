{"author":"meet","date":"2025-07-19","post_dir":"newsletter","status":"published","title": "Techstructive Weekly #51","type":"newsletter", "slug": "techstructive-weekly-51"}
## Week #51

Another hectic week, the work seems unstoppable and almost killing my back. Had to work almost 12 hours at stretch to get something in production. Not the best of times to be working at. Especially when life puts you in wired position, not just in work but in mental situation, you have to believe in yourself and amidst that every other bit of things happen.

### Random Thought

I was working away some bugs and fixing issues which was created for some other purpose at the time of initial prototype. But as the time evolved, it changed and deviated from the initial purpose, so fixed the current problem at hand. But people then point out your decisions. I felt grieved, lost, and most importantly wrong. But then things clarify, working all the possibilities of blunder out. Only problem was the decision made by the person in front of me when I asked the question in the past. But it was my fault in the end today. Not blaming anyone, just letting my thoughts out that people are too quick to judge and blame you when things go wrong rather than clarifying what they discussed first. World is cruel for the symphatic, for the humble, for the kind. It’s just is life. Can’t do much. But will you with one inspiration if you are lashed out at and the tides are against you.

> Crying is never a sign of your grief, its a sign that you are alive right from your birth.

Is it okay to cry? Probably yes, but don’t feel guilt and shame in crying out, cry alone but wipe out the tears and move on. At the end its only you who have to wipe your own tears, there are people who are around you who care, but crying in front of them might make them sad that’s what you would think. I too think that and cry alone. The journey is lonely, there are people that will help you, guide you, love you but in the end you need to put your foot ahead and walk yourself.

### Quote of the week

> “The oak fought the wind and was broken, the willow bent when it must and survived.”
> 
> — **Robert Jordan, **[**The Fires of Heaven**](https://www.goodreads.com/work/quotes/588034)

Such a great quote to describe my situation right now. I am just bowing down, accepting the flaws but not the things that were not mine. Just letting go of things, and let life show its wrath or the good behind the bad its happening. Because I can confidently say, every storm ends with a new beginning and there are definitely some things that flourish after the storm but some do get weathered out due to resistance.

## Read

1. [How I do it - Daniel Stenberg(Creator and Maintainer of cURL)](https://daniel.haxx.se/blog/2025/07/13/how-i-do-it)
    1. Such an honest and encouraging post. It shows how a single curiosity, a fun project can become a livelihood for someone or many people.
    2. He has made his curiosity take direction and kindled it with consistency.
    3. He also without shame and guilt admits the truth about maintaining a project or even working at something for quite a few years or even half a decade. Commendable spirit.
      > First let’s not pretend that it *always* feels fun and thrilling. Sometimes it actually feels a bit boring and *done. *There is no shame in that and it is not strange or odd. Such periods come and go. When they come, I might do less curl for a while. Or maybe find a corner of the project that is not important but could be fun to poke at. I have learned that these periods come and go.
2. [Reflections on OpenAI](https://calv.info/openai-reflections)
    1. A good type of post actually. Very rarely I have seen people reflect back on the company they worked and detailed the things he learnt and felt at the company without brag or hype especially for a company like OpenAI
    2. Everyone does a post where they talk about where they want to go next, or just complain about the things that made them leave the org. An important distinction for this post.
    3. It seem fascinating how bold and achieved is OpenAI, the risk and rewards had been blessed to OpenAI it seems, I am not saying they have not worked hard, just that it could have gone worse, if ChatGPT didn’t pick up.
3. [Kubernetes is not just for Black Fridays](https://ergaster.org/posts/2025/07/09-kubernetes-black-friday/)
    1. I mean this is correct, I don’t want to deny this. Having a dependable system that scales is a good thing, instead of hacking a script or a pipeline for deployment or CI, creating a heavy yet robust system is worth it.
    2. However, the only counterpoint I have is how you learn the things that the larger system does without using each component at its core.
4. [The great SQLite rewrite](https://bytes.dev/archives/407)
    1. It is a serious thing.
    2. People at Turso are seriously building it, it’s not a honey pot, last week I think they rewarded a thousand or more money to a person who actually found a bug in the Turso Rust version.

## Watched

- Philosophy is the last subject worth learning and studying amidst the AI hype
    - This might look like off-topic, but it really is relatable and worth thinking about
    - If you think, AI can do everything, what is there that makes human live?
    - What makes humans, human? the communication, the art, the craft, the way to express ideas? IF so why are we in a race to get that replaced with AI
    - So, he clearly distinguished the two things like theoretical and philosophical things. The former might be replaced by AI, but the later is the crux of what makes humans, human.
    - Worth thinking and reflecting upon if you are getting anxiety as a person with the new LLMs launching every day.

Double click to interact with video

## Learnt

- SQLC Issue with PostgreSQL Aggregate types
    - I was just trying to get back into open-source, just to distract myself and gain the lost fun back in programming especially after hard times at work.
    - I scampered through the issue list at sqlc and found this bug that caught me and wanted to fix.
    - Reproduced the issue, gave the filestructure to GPT, Claude and Grok and it gave me a few files that I could look for and find the bug after log slog and print debugging.
    - I reached a point after a few hours of interrupted debugging for a couple of days and found the core part where it misses the type.
    - There is no type inference done for the return type of the aggregate function having multiple schema type, the table schema is not being passed to the function that does it, so it just picks up the first map.
- Completed SQL Bolt course tutorial
    - Completed all the 18 exercises of the learning SQL starter guide on SQL Bolt
    - Learned a lot of interesting quirks about LEFT JOIN and RIGHT and NATURAL JOIN being only present with LEFT JOIN in the sqlite js with the wasm compatible in the browser.

## Tech News

- [Kimi AI K2](https://moonshotai.github.io/Kimi-K2/): This is probably the largest open weights model out in the wild, and that too probably the first that is really beating the state of the art models in tool calling and agentic coding. A huge standard leveller for open source community. There are quirks about the license but its pretty commendable in terms of what they have achieved.
- [Kiro IDE](https://kiro.dev/blog/introducing-kiro/): Amazon out of thin air just blew this up. Its almost like Cursor but with more controlled and flexible using preferences. It almost forces developers to plan and thoroughly consider the most tedious and boring part (sometimes the most exciting) of the development which is defining the problem and listing the features.
- [Cognition acquires Windsurf](https://cognition.ai/blog/windsurf): Devin will now use Windsurf, interesting deal this. How on earth are these two even partnering, doesn’t makes sense. At one point it looked they are rivals, even opposite mentality, but somehow they are collaborating now.
- [ChatGPT releases Agents](https://openai.com/index/introducing-chatgpt-agent/): This is like a thing that we all knew was coming but is late on the tables. Need to check it out if that is available for free users or not. If yes, then will be quite a fun thing to do and learn.

---

Didn’t consume much this week, produced a lot of shitty python and probably fixed a tons of bugs, but never got recognition. That’s quite a whirlwind to do, will be taking this weekend a slight self development route to upskill myself for learning SQL. I am finding it really interesting to learn SQL and contribute to SQLC, will continue on that.

[Leave a comment](%%half_magic_comments_url%%)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](%%share_url%%)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-754) (#754 edition) , and for software development/coding articles, join [daily.dev](http://daily.dev/).
