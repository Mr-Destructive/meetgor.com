{"author":"meet","date":"2025-07-26","post_dir":"newsletter","status":"published","title": "Techstructive Weekly #52","type":"newsletter", "slug": "techstructive-weekly-52"}

## Week #52

Finally a rewarding week, after 2 week of slogs. I / We were able to make the client happy after having gone through hundreds of issues and wired experiments. I was involved when the bug was introduced, but I had clarified that is this what we want, and at that time the answer was yes! Still in the end I was the one biting teeth and on the burn.

Anyways, at least the problem got solved and we were saved.

I read a ton about SQL and SQLite this week. Want to more this week too. I have a few ideas and issues that I am finding interest to work towards and finally at the end of the week or weekend will have something to share.

I am considering a mini series of SQL tib-bits, not a thoughtful things, but have something in mind over the month.

Ops! almost missed it, it’s week #52, that is, one year anniversary of [Techstructive Weekly](https://techstructively.substack.com/p/week-0). I can’t believe that. I just kept writing and here we are with 50 weeks of constant reflection and hitting the publish button.

I am not going to stop this, instead I fancy myself writing more over the months and years, because some has said it.

> Writing is thinking and thinking is writing.

There might be 104th edition of this newsletter with nothing new, maybe it will be completely changed, who knows what the future holds, but one thing is certain, something will change in me because of this. A hope, and not a false one.

### Quote of the week

> "A year from now you may wish you had started today."
> 
> — Karen Lamb

I don’t regret even a slightest that I started writing this weekly a year back. I learnt to articulate better, knew things I forget in a day or two. Helped me organise my learning and reading resources, be mindful about the media that I consume, helped me get over the blank page fear and most importantly helped me cultivate the reflective mindset rather than what next. Because this period when I write this is important for planing and thinking next what to do.

So, what is something I wish I would do now?

- Log the things and observations about LLMs and different models or even versions. Different techniques about coding with LLMs that work and don’t work
- Monthly Open source contribution reflections
- Daily Work log (private google doc) inspired from [Arpit’s ritual](https://arpitbhayani.me/blogs/worklog)

Will I be able to do all? Probably not, but even something better at the end of the year.

## Read

Here are some of the articles and tid-bits around the internet I read in this week:

1. [SQLite Internals: Pages and B-Trees](https://fly.io/blog/sqlite-internals-btree/)
    1. This is quite interesting and helpful in making things clear
    2. Every piece of data is stored in pages, a page is the unit of data in SQLite. Each page has parts like divided each for storing its metadata and the actual data.
    3. Each type has certain number of bytes to be stored, so there is a identifier for that, so it makes retrieval and storing efficient.
2. [How I keep up with AI progress and you should too](https://blog.nilenso.com/blog/2025/06/23/how-i-keep-up-with-ai-progress/)
    1. Simon Wilson, Andrej Karapathy and official sources are quite the right streams of sources for learning and keeping up with AI and tech in general.
    2. These people are really hands-on and have a decades of experience and wisdom.
3. [Rickrolling Turso](https://avi.im/blag/2025/rickrolling-turso/)
    1. Avi has been hacking on Turso, he is a database nerd.
    2. This was a fun way to demonstrate how to hack into a rust database, sqlite internals.
4. [Why LLMs struggle with analytics, and how tinybird solve it](https://www.tinybird.co/blog-posts/why-llms-struggle-with-analytics-and-how-we-fixed-that)
  > It ain’t about generating perfect queries, its about getting insights about data.
    1. That’s a great way to put it.
    2. With LLM context is really important and especially for critical things like database queries. Passing only the relevant info and cutting the fluff out is the key, but executing SQL queries in itself is that problem for the analyser.
    3. So LLM is just solving that step but for writing the actual query from the natural language.
5. [Simon Wilson’s commentary on GitHub Spark](https://simonwillison.net/2025/Jul/24/github-spark/)
    1. As per this and more experimentation, it could be said Spark is the most competent and well engineered AI Assisted Code Environment.
    2. Why? Because it’s GitHub, it has integrated the LLM in such a way that it seamlessly blends with GitHub Actions, Pages, Repositories, Secrets and Environments, Google had a great chance with FireStudio, but failed there.
    3. GitHub seems to have taken a better and more engineered approach and not rushed it.
6. [The Business case for Vanilla JS](https://lewiscampbell.tech/blog/250430.html)
    1. Just write plain javascript and shut it up. Why are we really bottling ourselves with the gazillions of frameworks and mental load of understanding each other’s semantics and constraints and styles.
    2. Why is there a need to create abstraction on top of abstractions. I get it, that writing javascript is cubersome each time. But the amount of abstraction and the cost of it adds is wild. We have a separate steps and ecosystem for doing just these, man!
7. [Memcached Architecture](https://hnasr.substack.com/p/memcached-architecture)
    1. Simple and straightforward explanation of the memcache
    2. Simple in-memory key-value store with slab-based memory management to avoid fragmentation. It has a threaded architecture and per-slab LRU for efficient concurrency and eviction, Also client-managed sharding enables distributed caching without server communication.
8. [Rate Limiting explained in 252 words](https://www.systemdesignbutsimple.com/p/rate-limiting-in-1-diagram-and-252-words)
  > how many requests a user or client can make in a given period
  And not control of how much a user can consume over a longer period

## Watched

Some random youtube video I watched at night after work during the week and weekends:

- [DHH on Lex Fridman Podcast](https://youtu.be/vagyIcmIGOQ?si=WTY3YTRFF5RUL8O4)
    - So many hot takes and might change my perspective on software engineering
    - Simplify development is important than elegance, eg. Ruby over Python, SQL over ORMs
    - Using editor to chisel out code, and not let AI rip it.
    - Finally someone to speak about the craft of coding, the fun of writing code alone, in the darkness, letting us spare time and let us cook.
    - Engineering managers are meaningless for startups.

Double click to interact with video
- [SQLite’s WAL Mode is faster than DELETE Mode](https://youtu.be/qf0GqRz-c74?si=HZ_1yav_DFOzyiOn)
    - This is so well explained, first showed everything what each one is and then the benchmark just makes everything clear.
    - The WAL mode basically writes the changes in a separate file and merges to the original db file whenever required, hence there is no overhead when reading or writing multiple writers or readers.
    - The delete mode is like a backup, a journal, it keeps pages of the data that are to be changed and after it is committed it deletes the file, that clearly looks slow.

Double click to interact with video

## Learnt

A few things I learnt while at my work and majorly side-questing after work:

- DUAL Table in Oracle and MySQL Databases
    - Dual is like a dummy table that you can reference in Oracle or MySQL databases as in these databases you cannot use query without a `FROM` clause.
    - So `FROM` clause is requried, and hence the dummy table is needed to satisfy that quirky semantic of those databases.
- SQL Sequence table
    - The `sqlite_sequence` table is internal table that helps sqlite keep track of the auto increment counter for each table’s primary key.
    - It has 2 columns, name containing the name of the table and `seq` denoting the max value of the id (auto-increment value), usually the last inserted id of the row in the table.
    - Key things to note are:
          - Only one entry for each table, so this will only create the entry for the table if it has a integer primary key with auto-increment constraint.
          - If you change it manually, SQLite will still select the max of the row primary key by getting the full-scan on that table.

## Tech News

Some of the things happening in tech, hot and smoking.

- [Qwen-3 Coder](https://qwenlm.github.io/blog/qwen3-coder/): It seems we have now a whole bunch of terminal based ai code war. Claude Code, Codex (does anyone use it?), Warp, Ampcode, Gemini CLI, now Qwen.
- [GitHub releases Spark for public preview](https://github.blog/changelog/2025-07-23-github-spark-in-public-preview-for-copilot-pro-subscribers/): This looks solid and well engineered product, will be trying out this weekend, finger crossed, will get time and will be streaming it.
- [Google is testing an vibe coding app called Opal](https://techcrunch.com/2025/07/25/google-is-testing-a-vibe-coding-app-called-opal/): Another one. We have no shortage of these, Firebase Studio, Lovable, V0, Bolt, and what nots.

I’ll leave you here, it’s been a great week atleast on the learning side, should be more fun next week and over the weekend.

---

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-755) (#755th edition) , and for software development/coding articles, join [daily.dev](http://daily.dev/).

[Leave a comment](%%half_magic_comments_url%%)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](%%share_url%%)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
