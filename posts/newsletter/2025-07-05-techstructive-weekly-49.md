{"author":"meet","date":"2025-07-05","post_dir":"newsletter","status":"published","title": "Techstructive Weekly #49","type":"newsletter", "slug": "techstructive-weekly-49"}

## Week #49

I’ll keep this week short, I have a lot to learn, and too little time. I have been studying SQL in detail, practising problem sets over the week. Taking a few more courses on Redis, Python, and more SQL.

One interesting thing, I kept on with my creative writing endeavour and made 14 days streak, writing with an average of 1200 words a day. So, I am roughly at 20k words after 15 days on my novel. This could be my first novel, could be even ready by the end of July if I keep going after it. Sounds exciting to me. Can’t wait to complete it.

I also want to create a weekly wrap-up and learning for SQL. I am really ready to dive deeper into the database world. I think this is what I can master and become a depth-first person for Databases and SQL.

### Quote of the week

> “Every now and then a man's mind is stretched by a new idea or sensation, and never shrinks back to its former dimensions.”
> 
> — Oliver Wendell Holmes, Sr.

I kept on writing my novel, one hour each day, I had no idea about the whole story, just kept a timer and no distractions, ideas flew after a while and the word count reached 1k, that momentum is enough to carry it ahead and continue the story next day and the day after, just like the mind has been stretched and formed with new ideas and uncharted water. Here imagination and curiosity take control and it feels so empowering and refreshing to discover that something existed within but never tapped.

Keep your mind open to new ideas, experiment, and you never know something will click like a light bulb and could change your life forever.

## Read

1. [Which LLM writes the best analytical SQL](https://www.tinybird.co/blog-posts/which-llm-writes-the-best-sql)
    - This was a great tutorial. It highlighted the point that “SQL is still a skill”. LLMs can generate SQL, but not analytical and efficient queries yet.
    - The cost is a factor, but right now it seems not worth it.
    - I thought LLMs were replacing people writing SQL, but here we are still requiring the domain experts. The people who will roll up the sleeve and press CAPs Lock and head to write SQL.
    - GPT models are decent in a balance of latency with accuracy, then comes Claude with high accuracy but slight slow, Gemini is good, especially the Pro, but takes time, the flash models are bad with faster times.
    - So, the LLMs are not perfect yet, they might get a few years time down the road, but it looks far from now.
2. [How long context fail](https://www.dbreunig.com/2025/06/22/how-contexts-fail-and-how-to-fix-them.html)s
    - This was an insightful observation, must know things before working with LLMs having a large context window, or even a short for that matter.
    - The needle in the haystack problem is not solved, and LLMs are very sensitive to getting up in a rabbit hole.
          - Context Poisoning
          - Context Distraction
          - Context Confusion
          - Context Clash
    - All of these reasons look the same, but can mean different things in different “CONTEXT”.
3. [How we accidentally solved robotics by watching 1M hours of YouTube](https://ksagar.bearblog.dev/vjepa/)
    - This is hilariously funny.
    - They trained the model on a corpus of YouTube videos, because sometimes, text and images aren’t enough for these kinds of operations like robotics, where moments are essential for learning and iteration
4. [Claude Code is my Computer](https://steipete.me/posts/2025/claude-code-is-my-computer)
    - I see this with many developers, not just newbies or just soydevs, but also veteran and knowledgeable developers, too. They seem to have a liking towards Claude code.
    - I know the feeling I think after using Amp, Warp, Gemini CLI to some extent.
    - But this is a bit of avoiding the chores part, which is fine.
5. [Just fucking use Kubernetes](https://waylonwalker.com/just-fucking-use-kubernetes/)
    - This article is true and motivating to learn Kubernetes
    - I read an article which was the opposite of this, not in a bad way, just a way of presenting the idea to keep it simple and avoid the complexity if not needed.
    - This article, however, is a motivator in a couple of sense, first it says if you think the scale is low, your ambitions are low. What a statement.
    - And also it’s an under kill if you think Kubernetes is an overkill
    - Nice AI-generated slop, actually is good.
6. [Writing code was never the bottleneck](https://ordep.dev/posts/writing-code-was-never-the-bottleneck)
    - Yes, code was never the bottleneck, but it was not also the easiest thing. People plan and plan, and plan more, but the execution is not up the mark, and the planning was of no use. Coding is neglected among managers but empathy is needed to bridge the gap.
7. [Why I want to write again](https://ordep.dev/posts/why-i-want-to-write-again)
    - I also want to write again. It has been almost a year since I started this newsletter, and I have not been able to break into a habit of writing the blogs that I used to back then.
    - The author’s claims are so true, and I can’t agree more, but the habit is the key part that holds me back. There is never a smooth life; something or the other hinders the focus.

## Watched

- [Understanding B Trees: The data structure behind databases](https://www.youtube.com/watch?v=K1a2Bk8NrYQ)
    - This was super cool, the way it was visualized and taught.
    - I adore Brain Yu, he is a master teacher. I learnt Python and Django from him.
    - B-trees are basically trees with an equal number of leaf nodes. No unbalanced roots. The operations to make the insertion and deletion made it really awesome for understanding.

Double click to interact with video
- [Agentic Coding: The future of software development with agents](https://www.youtube.com/watch?v=nfOVgz_omlU)
    - See, again, Claude's code and agentic tools are getting adorned by experienced developers too. This is the thing that makes me concerned about my own opinion and thoughts.
    - It surely is powerful but it hasn’t clicked for me yet I think.

Double click to interact with video
- [Vercel Finally Caught up](https://www.youtube.com/watch?v=Vd98UhPGVfY)
    - Vercel now makes you pay for only the CPU usage and not the time.
    - Still expensive from Cloudflare but a good dip in the number

Double click to interact with video
- [I finally switched to PostgreSQL](https://www.youtube.com/watch?v=iIMIKgRvS1Q)
    - Planetscale just added PostgreSQL as a supported database
    - Convex now switched to PostgreSQL
    - So, convex becomes a free tier for Postgres on Planetscale

Double click to interact with video
- [CS50 SQL Lectures Playlist](https://www.youtube.com/playlist?list=PLhQjrBD2T382v1MBjNOhPu9SiJ1fsD4C0)
    - I am learning SQL with this playlist, its in depth and also has a problem set to explore on our own.
    - I have completed watching all the videos and just yesterday completed problem set 1, will be moving on to solving more.

## Learnt

- SQL
    - Nested Queries
          - With this we can use select and that result set becomes a value for a condition for the outer select that is handy
          - = for single value and IN for range of values(list)
    - JOINS
          - INNER JOINS are only for matching records between two tables
          - LEFT JOIN will join everything on the left (first) table, even if there is no matching record on the right
          - RIGHT JOIN will join everything on the right (second) table even if there is no matching record on the left
          - FULL JOIN will join both tables, populating the NULL values if either of them doesn’t match.
          - NATURAL JOIN is the most unnatural part of joins which will combine the tables with the common named column(s) and remove the duplicate column
    - VIEWS
          - This is so cool. We can basically create this for reducing redundancy, Its like a macro, a stored function kind of thing, where it doesn’t have memory of its own but has things that it can run on existing data. Super useful to know

## Tech News

- [Claude Code now can support hooks](https://docs.anthropic.com/en/docs/claude-code/hooks)
- [Google’s Veo 3 launches worldwide for Pro/Ultra Tier](https://gemini.google/overview/video-generation/)

---

A week a bit of no new models, nothing fancy launches. Finally, we get a week or two before it gets wild. It’s time to reflect and build something meaningful. It’s time to learn and plant the seeds that will help in the future.

[Leave a comment](https://techstructively.substack.com/p/techstructive-weekly-49/comments)

Thanks for reading Techstructive Weekly! This post is public, so feel free to share it.

[Share](%%share_url%%)

---

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-753) (#753 edition) , and for software development/coding articles, join [daily.dev](http://daily.dev/).

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
