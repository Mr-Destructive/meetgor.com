{
    "author":"meet",
    "date":"2025-09-20",
    "post_dir":"newsletter",
    "status":"published",
    "title": "Techstructive Weekly #60",
    "type":"newsletter",
    "slug": "techstructive-weekly-60"
}

## Weekly #60

Another great week, consistently learning SQL, developing another streak for solving problems on FreeCodeCamp with Javascript and Python, reading instead of doomscrolling.

I have been learning about SQL since maybe 10 weeks now and finding myself in a good position, for the past 4 weeks I have written posts every day and it has helped me get back to the technical writing flow.

On the work side, it was a fun week, tinkering with a system that is already working decently and wanted it to get the most of it, turning the bits and knobs of a system is so cool.

This weekend, I think I would be simply starting a golang project to get tokenizer or something from scratch. I just want to relax and code something from scratch just for fun and learning purposes. Let’s see if I can livestream it.

### Quote of the week

> "The first principle is that you must not fool yourself, and you are the easiest person to fool."
> 
> — Richard Feynman

I have been fooling myself that I know stuff. Until I started to learn SQL, and boy o boy, I feel like I don’t even know a drop in an ocean. Database is a vast field. Such is life, there is a lot to learn, once to learn something, you get a little egoistic. That is where you have to wake up and be grounded and humbled, realise that what you know might be temporary or is already known or created by some other human. Respect it and move forward to gain more and provide your experience and guidance to others, the same way the people have done to make you learn it.

## Created

Wrote more blog posts on SQLite and Relations:

1. [SQLite: PRIMARY KEY column constraint](https://www.meetgor.com/sqlog/sqlite-primary-key-column-constraint)
2. [SQLite: PRIMARY KEY table constraint](https://www.meetgor.com/sqlog/sqlite-primary-key-table-constraint)
3. [SQLite: Foreign KEY table constraint](https://www.meetgor.com/sqlog/sqlite-foreign-key-table-constraint)
4. [SQLite: One to Many Relation with Foreign Key](https://www.meetgor.com/sqlog/sqlite-one-to-many-foreign-key)
5. [SQLite: Self Join Relations](https://www.meetgor.com/sqlog/sqlite-self-join-relations)
6. [SQLite: Many to Many Relations](https://www.meetgor.com/sqlog/sqlite-many-to-many-relations)
7. [SQLite: One to One Relations](https://www.meetgor.com/sqlog/sqlite-one-to-one-relations)

This makes it the 32 posts in total, have been writing daily for a month, and could see a lot of confidence and new energy surging within me. Writing really liberates the soul, it gives purpose.

## Read

1. [My Favorite Postgres 18 feature: Virtual generated columns](https://tselai.com/virtual-gencolumns):
    - I agree to this, there are pros and cons of both. Stored makes write heavier but are read efficient. Virtual makes it write easier and read heavier. You have trade-offs, you need to decide based on the computation that impacts how you want the column to be generated.
    - I don’t like the notion of JSON flattening in Postgres. Postgres is not a database that would be ideal for that kind of data. I know there are tons and tons of support for JSON, but tables and JSON, I can’t bare it at once. Those two are just separate entities for me. Maybe they are useful in one-off values, not not much. Switch to NoSQL if you have that lengthy data.
2. [Boring is good](https://jenson.org/boring/)
    - Yes, this post summarises the current trend in LLMs well. The hype is about people adopting to the value of LLMs, but soon they’ll will realise that they were on the wrong path (one example is they thinking, LLMs could replace developers) and then we will settle on the thing they are good at. I know, throwaway code, temporary code, and simple stuff that you know you need to do, you know how to do, but not worth the time to manually type it in and craft it.
    - I am really excited and positive about the SLMs, the small language models, I want to use it to just be a google search but simple and not ripping out my entire project into a react and python boilerplate mess.
3. [Work Hard, have fun, go home](https://www.bonnycode.com/posts/work-hard-have-fun-go-home/)
    - If you have fun you can work sustainably, else no matter how trendy or shinny the technology you would use, you won’t endure it long enough on the verge of burnout and the feeling of unsatisfied will crumble your efforts.
    - Maintaining balance is key, but flowing with the excitement and energy is also helpful not always but can be done to improve the fun and memorable parts.
4. [Building a lexical analyzer from scratch in C](https://devlogs.xyz/blog/building-a-lexical-analyzer-from-scratch)
    - This actually cleared up how to write lexers from scratch. I was wondering if that was bunch of if-else to parse each tokens, but we do have to group the kind of tokens and then write specific conditions on how to parse them. Now that makes a lot of sense.
    - Worth exploring more by writing my own markdown parser, even adding more features and syntax. If your soul screams to write your own flavour of markdown, let the muse take over you. This month or one day, not this weekend though.
5. [Creating a static site for all my bookmarks](https://alexwlchan.net/2025/bookmarks-static-site/)
    - I am very much this kind of person, I want to access my bookmarks and the linkblog too.
    - I am developing it, but can’t make it polished enough to be usable. I had created linkblog.netlify.app. This is work in progress, anyone can add but I would like to add the authentication, but that again creates a friction element, this all constraint bother me then.
6. [Python can open web browser for you](https://koaning.io/posts/python-can-open-a-webbrowser-for-you/)
    - Python -m webbrowser <link>
    - That is so cool, can’t think of other automation that can be done here.
    - Neat and handy
7. [Representing Graphs](https://thepalindrome.org/p/representing-graphs)
    - Edge List: List (tuple) of nodes and optionally weights.
    - Adjacency Matrix: Matrix of each node with the all other nodes, really great way to describe the graph, has everything that is needed to understand about a graph.
    - Adjacency List: Map of each node with a list of nodes that are connected with it, if weighted, then can add a list of tuples representing the node and the weight. A little tricky but the author says this format is the most used and is optimal for general use cases, so seems good enough.
  [![](https://substackcdn.com/image/fetch/$s_!5Jm3!,w_56,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Ff8b68cf8-d3f4-42f6-b8dd-cccde036005f_720x720.png)The PalindromeRepresenting GraphsHello there…Read more6 days ago · 51 likes · Alberto Gonzalez](https://thepalindrome.org/p/representing-graphs?utm_source=substack&utm_campaign=post_embed&utm_medium=web)
8. [Go Struct Alignment: A Practical Guide](https://medium.com/@Realblank/go-struct-alignment-a-practical-guide-e725c8e1d14e)
    - I have read this and it makes sense, a bit wired but nice. Writing structs should be carefully planned, so just add the largest ones at the top and cram all the smaller ones thereafter. The rule of thumb to follow if you have any memory-heavy or scarce use case.
    - Handy functions like Sizeof, Alignof, Offsetof are used to get the total byte size, memory alignment requirement, the field start position of the struct or any fields. Use it to craft the proper and perfect structure by tinkering and aligning.
9. [Myopic Focus](https://thedailywtf.com/articles/myopic-focus)
    - Wow! Fantastical and Tech stories, are my love.
    - Gist of the story is people are very myopic(viewing in very hindsight) about system and think of the existing workflow as ritualistic that makes it harder for others to make it better or even improve the quality or realibility.
    - The story goes like > A new developer chops replaced fragile ID pool logic with robust UUIDs, eliminating inevitable crashes. But their manager fixated solely on declining unit test numbers, demanded rollback. It made the project a bit risky to sustain, but the fault is not of developers. Its the myopic view that causes it.
10. [What AI chatbots are doing under-the-hood, LLMs from scratch part 1](https://www.gilesthomas.com/2025/08/what-ai-chatbots-are-doing-under-the-hood)
    - This is a great series, I am reading more about it this weekend. LLMs are things that produce logits, which is like a list of probabilities from a vocabulary and for each token the highest probability token should have been chosen, however it is random the temperature determines how random the selection is, 0 means choose the first(highest probable, no creativeness), 1 means choose any random one. Between these 0 and 1 you can experiment to find the sweet spot for your needs.
11. [Magical Systems Thinking](https://worksinprogress.co/issue/magical-systems-thinking/)
    - I love this post, it hits home for me
    - Systems should grow from the simplest possible solutions and then branch off from the possibilities and situations.
    - Creating patches to the existing systems will only survive for limited time or none at all.
      > Sundar Pichai estimated in late 2024 that over 25 percent of Google’s code was AI generated; as of mid-2025, the figure for Anthropic is 80–90 percent.
    - The comparison to AI slop is great and fits well here. Vibe coded mess is no exception to systems thinking. We have seen LLMs want to patch and keep on patching existing code mess, we want to start from scratch, the urge is right, but we worry about whether we will lose the progress? The progress is fake progress.
    - The code LLM produces will always be throwaway code, as it is no subconsciously written, call me philosophical, but code is art and it needs attention form a human soul in order for it to work, not technically but overall in order to complete its purpose.
      > NEW SYSTEMS CREATE NEW PROBLEMS’ and ‘THE SYSTEM ALWAYS KICKS BACK’. As systems become more complex, they become more chaotic, not less. The best solution remains humility, and a simple system that works.
12. [The Sad, Sad world of Tech Blogging during the Era or Technological Stagnation](https://freddiedeboer.substack.com/p/the-sad-sad-world-of-tech-blogging)
    - There is nothing really ground breaking in mobile phones that is true, in LLMs too this seems to be getting closer to. The 90% closeness is achieved, the 10% will be like a forever process, people making predictions of AGI by 2027, and all that hype-crap is non-sensesical hypothesis.
    - Does this mean, tech world is plateaued? Maybe not but unless we have something wild, nothing is going to change marginally just like gpt 4 and 5 bump, they know it but want us to feel like its a jump
  [![](https://substackcdn.com/image/fetch/$s_!no2m!,w_56,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fbucketeer-e05bbc84-baa3-437e-9518-adb32be77984.s3.amazonaws.com%2Fpublic%2Fimages%2F0bc5fd66-6f8a-4d34-add5-3eff35a4e30e_512x512.png)Freddie deBoerThe Sad, Sad World of Tech Blogging During an Era of Technological StagnationI have a piece for Vital City arguing that, to succeed in the negotiation and compromise that will be required to build a ton of new housing, YIMBYs have to stop engaging in convenient and false caricatures of who opposes new construction. Check it out…Read morea day ago · 120 likes · 54 comments · Freddie deBoer](https://freddiedeboer.substack.com/p/the-sad-sad-world-of-tech-blogging?utm_source=substack&utm_campaign=post_embed&utm_medium=web)

## Watched

- [David Heinemeier Hansson (DHH) on Rails World 2025: Opening Note](https://youtu.be/gcwzWzC7gUA)
    - I can’t fathom here, how did we end up backwards? It takes some thinking to deploy to prod, yaml manifests, and what not to deploy a simple API, how the heck people are accepting those?
    - It is like a egoistic culture to have complexity and assume that it will work, but they are only adding complexity upon layers of complexity that developers will have to clean up themselves.
    - Rails is doing a phenomenal job in the developer ecosystem.

Double click to interact with video
- [OpenAI dropped another Codex](https://youtu.be/j9wvCrON3XA)
    - I agree to this, it is just a slight improvement to the GPT-5 experience. Not much here.
    - I kind of hate to admit that free users can’t use the web codex, the cloud instance of the agent. So, basically free users are devoid of the code experience of the GPT models.
    - Google might win that race. Gemini CLI is so cool, they also have Jules.

Double click to interact with video
- [Finding and Exploiting Buffer overflow password cracking in C](https://youtu.be/W1okWEzGAGo)
    - This guy is producing great stream, bangers after bangers. This one not quite remarkable but enduring nonetheless. I will vouch for persistence over hype and fuss.

Double click to interact with video

## Learnt

- Javascript Tidbits (TILS)
    - Rounding to a float using `toFixed(n) `this will round off to n digits of a decimal/floating number. Python has round(), Javascript has toFixed, kind of confusing name but fine.
    - Math.floor() to convert a decimal to an integer, this I used it to perform integer division. I come from python and `5//3` would give `1` but Javascript `//` is a comment literal so need to use Math.floor instead.
- SQL Logs
    - Learned to use Foreign keys and create
          - One to One relations
          - One to Many relations
          - Many to Many relations
          - Self referencing relations
    - In SQLite, foreign key constraint is not enabled by default, you need to explicitly set it as `PRAGMA foreign_key=on`. SQLite is very flexible, though it sometimes feels like a double-edged sword, so flexible by default, strict only when said, kind of good, but one can easily shot on the foot.
    - I tweeted about it [here](https://x.com/MeetGor21/status/1967818642768470447). Generated columns are quite new. This was added in:
          - SQLite [3.31](https://www.sqlite.org/releaselog/3_31_0.html) (in Jan 2020)
          - PostgreSQL [v18](https://www.postgresql.org/docs/18/release-18.html#RELEASE-18-UTILITY) (in Sep 2025)
          - MySQL [v5.7](https://dev.mysql.com/doc/refman/5.7/en/mysql-nutshell.html) (in Oct 2015)

## 

## Tech News

- Java 25 Released
    - How can one miss this update?
    - [Compact source files & instance main methods (JEP 512)](https://openjdk.org/jeps/512)
    - Finally the boilerplate is reduced in Java, one less thing to hate Java for. I never hated it for that, but never mind.
- OpenAI releases Codex, model, CLI, and cloud agent, or everything
    - Hhh! How many products Open AI has with Codex? I know the name is nice, but move on with it and be creative, the o- series was less confusing then this mess.
- [Meta AI live demo of glasses that went bad](https://techcrunch.com/2025/09/17/meta-unveils-new-smart-glasses-with-a-display-and-wristband-controller/)
    - Atlease they tried, that’s what I would say, they aren’t faking that it works like magic unlike others.
- [Notion Launches Agents in 3.0](https://techcrunch.com/2025/09/18/notion-launches-agents-for-data-analysis-and-task-automation/)
    - Every product is becoming something with AI.
- [Duocon 2025:](https://duocon.duolingo.com/)
    - This is an App I am admiring for its UI and UX, so clean and intuitive
    - Addictive design is a different thing, but its utility and worth using
    - I was awaiting the [Chess feature since May](https://blog.duolingo.com/chess-course/) and it finally launched this week. Can’t wait to up my chess skills

That was a heck of a long week. Lot of tech things happening. It will get even more in the next couple of months, as it is October and there are already a lot of developer related things, conferences, hactoberfests, and among other things.

---

[Leave a comment](%%half_magic_comments_url%%)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](%%share_url%%)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.

---

Happy Coding :)

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-763) (#763rd edition) , and for software development/coding articles, join [daily.dev](http://daily.dev/).
