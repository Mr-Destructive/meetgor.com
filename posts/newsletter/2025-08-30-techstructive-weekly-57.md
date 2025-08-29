{
    "author":"meet",
    "date":"2025-08-30",
    "post_dir":"newsletter",
    "status":"published",
    "title": "Techstructive Weekly #57",
    "type":"newsletter",
    "slug": "techstructive-weekly-57"
}

## Week #57

It was a fun week, really doubled down on learning SQLite and SQL. Shipped a bit of code and had fun creating improving metrics on the work side of things. Able to see 10% improvement in the things I have been working on the past few months. Really a good vibe week. I finally broke the barrier or imposter syndrome and whatever it might be called it while using the AI-assisted/vibe coding things.

### Use AI for what’s it god at: Throwaway code

I actually started to understand the areas where I need and can 100% use AI tools like throwaway scripts, CSV/JSON dumping, data extraction scripts things like that. I used Cursor-Agent to do the most of things and I was happy at the end of the day, it nearly one shot it, I had to do a few tweaks here and there but was mostly correct, nothing like starting to write the script from scratch after it giving up. It really helped me save at least half to an hour worth of work in 5 minutes almost 3 days in a row. I felt good, I don’t felt like doing that chores so nothing valuable is lost in that.

I can focus on the other priority items. So, this was a valuable lesson for me to use AI effectively where it shines, in generating shit lot of code that nobody cares but I need to get the job done, generate a few data points from sapheggeti logs and piles and oodles of data. I don’t care if it uses a list comprehension or a C-style for loop, I didn’t even bother looking at it, here readability has 0 value. So use AI to generate plots, structured data from logs, benchmarks (review the samples first though), among the other things.

I can breathe in comfort, after knowing that I develop that intuition as a developer, those things are still valuable and evergreen even if AI can spit out code, there needs to be the one who can guide it in a way to get the job done **effectively**.

---

Enough of shilling and rambling, I am back on track with the excitement and fire back, will be recording a few videos and might do a live stream if time management is done right this weekend.

### Quote of the week

> "You may say that it was the book, but it's how you read a book that is most valuable, rather than any power it might have itself."
> 
> — [Michiko Aoyama](https://www.goodreads.com/author/quotes/18144498.Michiko_Aoyama)
> 
> [(Author of What You Are Looking For Is in the Library)](https://www.goodreads.com/author/quotes/18144498.Michiko_Aoyama)

I read this book over the weekend, and really found it stuck with the weekly things I did. For instance, the AI thing, it might be powerful but if we don’t know what it’s capable of really and where and when to use it, it feels worthless, it feels forced, it feels hyped (not gonna lie, its over-hyped though).

## Created

- Added Interactive SQL Code-blocks / Playground on my blog
    - Try it here
```sql 
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL
);
INSERT INTO users (name, email) VALUES
    ('John', 'john@example.com'),
    ('Alice', 'alice@example.com'),
    ('Bob', 'bob@example.com');
SELECT * FROM users;
```

- Added [RSS feeds](https://www.meetgor.com/feeds/) specific to tags and post-types
- LogSQL Entries for each day of the week
    - I wrote a bunch of articles around SQLite specific things like IIF, random, unistr and unistr_quote scalar functions, the read dot command, and specific things about the version upgrades.

1. [Random Scalar Function](https://www.meetgor.com/logsql/sqlite-scalar-function-random/)
2. [Randomly order rows](https://www.meetgor.com/logsql/sqlite-randomly-order-rows/)
3. [Once dot command with html tables](https://www.meetgor.com/logsql/sqlite-dot-commands-html-tables-with-output-and-once/)
4. [SQLite unistr and unistr_quote](https://www.meetgor.com/logsql/sqlite-functions-unistr-and-unistr-quote/)
5. [SQLite’s IIF and IF scalar functions](https://www.meetgor.com/logsql/sqlite-scalar-function-3-valued-iif-and-if-scalar-function/)
6. [Read dot command](https://www.meetgor.com/logsql/sqlite-scalar-function-3-valued-iif-and-if-scalar-function/)
7. [Read dot command is recursive](https://www.meetgor.com/logsql/sqlite-read-dot-command-is-recursive/)

## Read

- [Learning Peal in one day and the importance of building strong foundations](https://guilhermenl.dev/articles/9096ed7725d387606d713e7964e2b3ac06f9bebd2650080b9ca070f0106f5c70)
    - Great article, it shows the author cares about his craft, I can relate to the feeling of learning a lot of things.
    - Doing things faster can lead to cutting corners, that’s a good take. I think we can agree to that, we are loosing the ability to know something deeper. Learning is becoming easy, but not deeper, its shallow, brittle and hallucinated. Failure-based learning is more valuable then productive driven success.
- [Job Churn](https://seths.blog/2025/08/job-churn/)
    - Hope, is what the article gives. If the history has thought us something, then it will repeat itself and we will have jobs disappearing and appearing again, in a different form and domains.
- [AI Slop in Open Source Contributions](https://blog.stuartspence.ca/2025-08-declining-ai-slop-mr.html)
    - I agree to this post, its quite obvious that if the author of the MR/PR has just submitted a vibe contribution (without even looking or testing) he/she doesn’t even care about the project, its a red flag.
    - I have contributed to a few open source projects and have been on the other side (contributor side) and I use AI to understand the problem, help me write a fix, but I do that in a worktree separate from the branch that I am working, I nit pick details that I think are important, it produces a lot of things, I examine each line and word carefully and have a lot of to and fro for understanding why it added it, usually the reasons are lame and on the lines of of course why not! I delete without thinking twice.
    - I think people are against using AI as is and not as an assistance for helping you contribute and solve the actual problem, if AI had solved it, why even the maintainer would require you to create a PR and all the overhead in between understanding your thoughts on this.
- [OCR is Legacy tech](https://www.cloudsquid.io/blog/ocr-is-legacy-tech)
    - Maybe it is, but not completely. There will be need of low cost easy document extractions. Yes, if the cost factor goes down, OCR would be obsolete.
    - I work in these domain, and can say, a lot of things are changing quite rapidly, LLMs and VLLMs are taking the extraction industry by storm.
    - Its quite easy to one shot an extraction with VLLMs given the intricated details might be missed by OCR and with VLLMs, its like a brain powered visual detector. It really understands the context and its knowledge really comes in the clutch to parse almost any type of document.
    - Really excited to see where my journey leads in this development
- [Grok 2.5 is not really open source](https://www.zdnet.com/article/no-grok-2-5-has-not-been-open-sourced-heres-how-you-can-tell/)
    - The License is not truly open source
    - Its not open weights as there are restrictions on its usage
    - So its basically weight available model, and you can’t really get any value from the weights.

## Watched

- [Teej Devries on the Wookash Podcast](https://youtu.be/RNDqPoZvdL8)
    - I like Teej Approach, read the f-ing manual. Writing cover letter, I am not sure about the people might get negative impression of AI-written, but I can see the effort put in by the candidate or the desperation of getting the job.
    - Going full time on streaming is really bold decision and he is really a great person to do just that, a real practical teacher. A true inspiration.

Double click to interact with video
- [Please Stop using booleans in database designs](https://youtu.be/xIRL3klHM9I)
    - This is great advice, I get the point, having updated_at is much better than having is_complete, etc.
    - The enum example also makes sense. Its basically like having a breakdown of the field into required pieces.
  Double click to interact with video

## Learnt

- SQLite Week 6
    - I learnt a lot while creating the bits and pieces in my daily logsql blog posts. Especially the details of the dot commands, scalar functions and new updates in the SQLite 3.47+
    - The IIF is more rich now, as it can take in from 2 to N arguments, where N can be conditions and pair of conditions to get the value if the condition is true, this is new, the older one was 3 way condition, the first condition, the second as the true value and the 3rd being the false value (it can be nested). However this both are now in the same function, maintaining backwards compatibility. It also added an alias as IF making it easier to come from other databases.
    - The unicode and unistr and quoting the emojis in SQLite. I am suspecting because of LLMs these two functions were added, as LLMs are generating a lot of emojis people might be tempted to store something in the database and use it later, so escaping the unicode and quoting it is required to make it work with the unicodes.
    - The read and the output commands are really interesting, the former is even recursive, found a bit of easter egg with it today.

## Tech News

- [Zed adds Gemini CLI](https://zed.dev/blog/bring-your-own-agent-to-zed) as a support for [A2A](https://agentclientprotocol.com/overview/introduction) protocol
- [Claude for Chrome](https://www.anthropic.com/news/claude-for-chrome)
- X AI makes Grok 2.5 open source (not really, its just weight available model)

---

Happy Coding :)

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-760) (#760th edition) , and for software development/coding articles, join [daily.dev](http://daily.dev/).

[Leave a comment](%%half_magic_comments_url%%)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.
Subscribe to receive new posts and support my work.
