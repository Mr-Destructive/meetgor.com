{"author":"meet","date":"2025-08-09","post_dir":"newsletter","status":"published","title": "Techstructive Weekly #54","type":"newsletter", "slug": "techstructive-weekly-54"}

## Week #54

What a week! LLM models all over the place, gpt this, gpt that.Anyways, I find myself in a better place by avoiding AI in some places and using at its full potential at certain place, it took some time to realise it, but its kind of working now.

I don’t use AI to build the side projects like one shot, but use AI to help me understand certain topic, do something myself and converse with it about what it means, and what gaps I have in it. Using AI (cursor) at my work to speed things up and staying calm and letting things happen.

This week I continued learning more SQL, for the past 3 weeks, I have decided to learn SQL, just because I am curious and lo behold found a bug that can be fixed in SQLC, patched it, created a pull request and done. Started a series for solving SQL problems on exercism SQLite course. Finally back on livestreaming. Finally a good week, after a while. Feeling like a human, after dreaded existential crisis due to AI maybe.

### Quote of the week

> "The self is not found in comfort, but in the friction of becoming."
> 
> — Rollo May, *The Courage to Create*

I think I had been stuck in a cycle of consumption, I think I am coming out of it, by creating, I have not written much. This week I am thinking of writing blindly. I have recorded a couple of more videos for SQLite problem solving, and planning to add more over the weekend. There are a few hackathons that I think I would be participating, and would have a few projects out of it to showcase it. That was friction and that gives me the dopamine to wake up in the morning.

## Created

- Exercism SQLite Track: I started a series of solving SQLite problems on exercism and have plans to solve leetcode 50 sql, CS50 sql problemsets too. It will take time, but I would have learnt SQL then no other resource but practise and also upped my communication and teaching skills.
    - [Hello world and Local environment setup](https://youtu.be/IaqlpeWexig)Double click to interact with video
    - [Two Fer](https://youtu.be/-sZZfSTUrko)
  Double click to interact with video

- OpenAI’s GPT-OSS Models review for 20 billion parameter model
    - [Video](https://www.youtube.com/watch?v=3aiJN2uGmZk) and [article](https://www.meetgor.com/thoughts/openai-open-weight-oss-model/)
  Double click to interact with video
- Open source contribution to SQLC
    - [Bugfix on DROP/ALTER/INSERT table values for case-sensitive table name](https://github.com/sqlc-dev/sqlc/pull/4045)
          - For issue > [#3936](https://github.com/sqlc-dev/sqlc/issues/3936): Bug on SQLC generate for queries relating to dropping, and update queries having case sensitive table name whereas, the select, insert and delete queries are case-insensitive to table names.
- I finally livestreamed after a month or two about Golang project.
    - Pocket clone, simplifying the initial prototype. No AI or vibing, just me and the editor. I use AI autocompletion, the next word autocomplete (calling it copilot level like thing pre cursor era)Double click to interact with video

## Read

1. [The fastest way to detect vowel in a string (Python)](https://austinhenley.com/blog/vowels.html)
    1. Wow, this dude just found 11 legit (almost 13) ways to detect vowels in a string in python.Such a great depth, the benchmarks feels so intuitive as why each way performs the way it does.
    2. Here are all the ways it did it
          1. For loop: Simple, readable. Fastest for small strings
          2. C-Styled for loop: Uses or comparisons, but surprisingly much slower
          3. Nested for loop: Totally exhaustive, but slow
          4. Set intersection: Clever and clean. Great when strings are long or vowels are sparse
          5. Generator expression: Pythonic one-liner. Reasonably fast, readable
          6. Recursion: Functional but inefficient. Crashes on long strings
          7. Regex search: Shockingly fast. Calls C-level code internally
          8. Regex replace: Works but inefficient. Doesn’t short-circuit
          9. Filter: Readable but wasteful because it processes the whole string
          10. Map: Similar to filter but slightly better
          11. Prime Numbers: Extremely creative. Maps characters to primes, uses GCD. Way too slow to be practical
    3. Would like to do something in Golang, it sounds so fun that I can’t stop thinking about so many ways to do so trivial things.
2. [Don’t tell engineers what to do?](https://code.dblock.org/2025/07/30/dont-tell-engineers-what-to-do.html)
    1. This post, hits me and in a good way. It basically helped me get out of the self constrained mindset that I was doing wrong, instead I think I am doing my job. That is to listen and implement, the moment I raised objections or pointed out a few things, I was turned down as a kid on a shop. I am not embarrassed about it, just that the mindset in which I am working might be a little old or startupy. I am not saying I am perfect, just that there is a way to address engineers and their concerns.
    2. Everything in the post that has been said, is totally relatable and worth reading for me. I read and felt a deep satisfaction that I might be just right and move on with the things.
3. [Python’s pass by value and pass by reference](https://www.thepythoncodingstack.com/p/python-pass-by-value-reference-assignment)
    1. This is one hell of a reason, Python gets a little more confusing and less friendly.
    2. TLDR of the post is that if you pass a immutable variable/object to a function call in python, you need to return it back from the function (if the function modifies those immutable objects). Because the object is immutable it won’t get updated inside the function, it will be created a new, so we need to assign it to the modified version when the function returns.
    3. But for mutable objects, the function can modify it and we are passing it to the function, so the object will be updated.
4. [The art of saying yes: How do I do so many things](https://brittanyellich.com/say-yes-do-all-the-things/)
    1. Totally relatable post, I do a lot of things, because I am curious and excited to learn about it.
    2. Starting small, building momentum, building a habit, is really what got me into this exact newsletter, one brings other, I kept writing one week after other and here I am writing the 54th edition. Not all weeks are fair, the same intensity or energy, some are dull, frustrating, just barely pushing past, but there are mostly the weeks where I know what I am doing, why I am here writing the post, full of curiosity, eager to learn more, and break prod.
5. [Flashy, Fancy shortcuts aren’t always suitable](https://www.thepythoncodingstack.com/p/flashy-fancy-shortcuts-arent-always)
    1. Using or to handle mutable default values in Python functions can create unexpected new objects when passed an empty list, breaking the link to the original. The standard if None check is more reliable and readable, ensuring the original list is modified correctly.
6. [100 Rust exercises to learn Rust](https://blog.jetbrains.com/education/2025/07/28/rust-exercises-rustrover/)
    1. So good. I will be using this to learn rust within a week it seems. I skimmed through a few posts, and found it really nice and interesting to go through them, clear and nice examples.
7. [Stop building AI tools backwards](https://hazelweakly.me/blog/stop-building-ai-tools-backwards/)
    1. AI should augment, not replace, human learning through effortful retrieval and collaborative iteration. Buttons that auto-do tasks deskill users and erode trust, prioritize interactions that spark learning.
8. [How to become a confident software engineer](https://newsletter.francofernando.com/p/how-to-become-a-confident-software)
    1. Valuable advice and practical too
          1. Pick and learn a single programming language properly (like I am doing with golang)
          2. Write unit tests and CI (Ah, now it makes sense)
                  1. I have not written serious tests in my 1 year long career, rest apart from the 4 year tech journey. Now it makes sense, there is a purpose to give me confidence to build reliable software.
          3. Make refactoring a habbit (encourage to do, but done or not is sometimes not in our hands especially for juniors)
          4. Pair with people (I do frequently whenever the need arise, and learn a few git commands and browser speed up shortcuts which help me learn more)
          5. Read books (I am currently reading Learning SQL to learn SQL in depth)
          6. Teach what you learn (doing it by youtube or articles I post)

## Watched

- [You can just do things](https://youtu.be/dgr-WAUgELw): Aaron Francis at Laracon US 2025
    - You don’t need to be extraordinary just be out there to do things. Being kind, and expressing interest and curiosity is the way to get things done and build meaningful connections.

Double click to interact with video
- [RSS is not dead](https://youtu.be/hSgpW0j2eOk)
    - RSS is not dead. Wow! I just had this vision of building a read later app a few weeks back, when the Pocket was deprecated. I want to build a combination of Read it later and RSS reader, a sync engine is something that might be used here if using cross-platform.
    - Now I understand the conversation here, if someone is fetching the posts and storing metadata then it makes sense to have a paid solution.
      [Bluesky post](https://bsky.app/profile/waylonwalker.com/post/3lvobbewtcc2f)
    - But the other point I am afraid of is the content scraping, is it ethical to do that? Like I know its not stealing, but crawling a legit site and using it on other platform is not good for link building for the author of the original article right? I am not sure, need a bit of thinking and ethical considerations.

Double click to interact with video
- [The painful truth about startups](https://youtu.be/lWsZT-2pQL4)
    - OK, it was a harsh journey. Success is not overnight, its just a preparation for a overnight success. Uploadthing, round, picthing, and so many thing, and then finally T3 chat (chat thing, better chat), that made it.
    - The salary looks like it will be life worth income. Man! india has a really low income-expense ratio.
    - Just nerd out about the things I love, that’s what I am doing here.

Double click to interact with video

## Learnt

- SQL Week #3
    - String concatenation in SQLite
          - printf function that exists in SQLite. That is such a great thing to have, C like function, just drop the placeholder for the appropriate type and it will be a formatted string, neat and tidy.
          - || operator for simple concatenation. You can just use “Hello” || “ World!” to get a string “Hello World!”. It’s compact but for large number of strings, or readability, it gets a little tricky.
    - I finally wrapped my head around autoincrement and sqlite_sequence table.
          - Autoincrement is a constraint that will force the newly inserted rows to use primary key id values greater than any existing or deleted primary key id values.
          - If the highest row created is deleted before inserting a new record, the sqlite_sequence comes in handy to fetch the max value of the primary key id across the table, since the value is stored in a separate table called the sqlite_sequence with table_name and the seq columns.
          - The sqlite engine decides to get the max of the seq value and the current max row id (it can effectively get it using B+ trees, as it will be the leftmost node or right-most node, however its stored)
          - I will be creating a separate blog on this, a deep dive on the various cases we can run into if the sqlite_sequence table gets altered.
    - USE common table expressions in sqlite
          - This is a way to create a temporary table (like only valid till the query completes running) and use it in the main query(can have nested queries too)
          - The syntax looks like this :
            ```
            WITH <temp-table-name> AS (SELECT something from somewhere)
            SELECT something, <temp-table-name>.something from elsewhere 
            ```
          - This is something I studied while solving the difference of [sum of squares and squares of sum problem in Exercism SQLite track](https://exercism.org/tracks/sqlite/exercises/difference-of-squares).
    - SQL’s IIF is equivalent for CASE WHEN THEN
          - The IIF is like an handy if else block to use when having nested conditions.
          - This works and looks neat for smaller expressions, like one or two condition max, after that its better to use CASE
          - The syntax looks something like:
            ```
            SELECT IIF(something > 10, "YES", "NO") as answer FROM somewhere;
            ```

### Interesting Links

- https://github.com/zakirullin/gpt-go
- https://blog.jetbrains.com/education/2025/07/28/rust-exercises-rustrover/

## Tech News

Phew! can we start counting the number of models released this week, this was that was kind of week.

- OpenAI releases GPT-OSS (open weight models under Apache 2.0 license)
- OpenAI releases GPT 5 chat, mini, nano with and without thinking capabilities
- Anthropic drops Claude 4.1
- Google’s Genie 3
- Horizon Alpha and Beta models on openrouter
- OpenAI releases Study mode in ChatGPT and Google releases Guided learning mode in Gemini
- Appwrite goes public (generally available)
- Google’s Jules AI Agent goes public (generally available)

I think we can say August start week was a huge bag of announcements and model releases.

That’s it from this week! A fun exciting week, not at work, but having fun on the side.

See you in the next one!

Happy Coding :)

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-757) (#757th edition) , and for software development/coding articles, join [daily.dev](http://daily.dev/).

[Leave a comment](%%half_magic_comments_url%%)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](%%share_url%%)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
