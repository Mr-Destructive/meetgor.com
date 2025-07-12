{"author":"meet","date":"2025-07-12","post_dir":"newsletter","status":"published","title": "Techstructive Weekly #50","type":"newsletter", "slug": "techstructive-weekly-50"}

## Week #50

A bit of boring week, had a lot of hopes starting on Monday, but it just tumbled down. Also broke the 15 day long writing streak. It all just falls apart, I want to change things now. I work remotely, and not able to balance work-life and personal life, its just too exhausting. I get self-doubt creep up with this, am I really a sustainable developer, can I not continue development for the next 10 years even? It is starting to haunt me. I know it would be temporary, imposter syndrome feeling, but need to really change the flow of my daily routine to get a mindset shift.

But learnt a lot this week, about SQL continuing from the past week, started solving SQL 50 problems, which I am halfway through. Also need to ramp up my problem solving skills, need to build projects, there is lot of stuff to do. I just get overwhelmed with stuff every time.

I haven’t streamed for the past 3 weeks and it is giving me negative cloud of procrastination. Hopefully will be able to break the jinx this weekend. July has been good to me in the past, hoping for the best again.

### Quote of the week

> “It’s not about how hard you hit. It’s about how hard you can get hit and keep moving forward.”
> 
> – Rocky Balboa

Yes, I admit I might be hit, I have fallen, I no longer have the motivation to code for some patches of the week. But I have the spark back at times and flash of seconds. It might be a temporary burst of problems at my end, but every storm weathers out, every night ends and that’s when the comeback begins, the sun strikes and light appears, there we see a hope, and the light at the end of the tunnel, it might be small, but enough to keep moving us ahead.

I am willing to change my workflow and routine, I just need to give it some time. If habits are easy to build they are easy to break too. I know this with my 15 day writing streak that I broke, I didn’t had the mental energy to sit and write for an hour, and that’s the reason I am frustrated with my life, with my work. I am adamant to change.

By the end of July the max, the first year anniversary of this newsletter will be special one.

## Created

- SqueaLite Browser Editor → [squealite.vercel.app](https://squealite.vercel.app/)
    - Just wanted a quick and easy way to write SQL queries and test out on a database.
    - Turns out sql.js is a great solution for it, it’s a WASM library to interact with SQL in the browser with JS. I am using it to create a database for that session for the user in the browser and let the user use it as a playground database.
    - I added (AI-assisted) editor, result panel, multiple query tabs, loading and saving queries and databases, among other visual aided table things.
    - Planning to add two more sections
          - Problems → Like leetcode styled run and submit challenges.
          - Exercises → Independent and specific playground for understanding specific concepts.
          - Per user database for playground and exercises using Turso libsql database.
- [Analogies are not always good: Comparing VACUUM in Postgres with free in C](https://substack.com/home/post/p-167661543)
    - Well, this is a off-topic write-up I felt writing because no one seems to have compared VACUUM in Postgres and free in C.
    - Those 2 might look unrelated but for understanding it makes a lot of sense. Actually I learnt that VACUUM is not like free, but VACUUM FULL does look like free.
    - My point is drawing analogies is good but not when you are not sure of either of them, it could look good on the high level talks, but getting a little deeper will shatter the comparison if not thought carefully.

## Read

1. [Tools: Code is all you need](https://lucumr.pocoo.org/2025/7/3/tools/)
    - MCP is not worth the context I think, and this proves it. I have no experience but just from the understanding it would make sense, since a lot of things need to be made clear before using them right?
    - Imagine a human navigating this, you will have to make him aware that these are the tools, explain what each tool does and then provide them. But not to a professional, so we can conclude that LLMs are not professionals yet!
    - MCP is a bit verbose and might not be right for all kinds of models with less context window or limited capabilities.
2. [Thoughts on motivation and my 40 year-old career](https://charity.wtf/2025/07/09/thoughts-on-motivation-and-my-40-year-career/)
    - This was a intense read, a bit of self reflection and some insights into business vs wage job.
    - I think engineering is hard to explain by the looks of the post, its a bit of complex thing to be working at, not hard, but you need time and patience to get used to it.
3. [I still care about the code](https://martinfowler.com/articles/exploring-gen-ai/i-still-care-about-the-code.html)
    - This is so well put, LLMs are not compilers or interpreters, they still are unreliable. I mean ~~software is unreliable~~, **untested software is unreliable** but LLM generated code is on the extreme end of black holes.
    - Hallucination is a things, I am not talking about on-surface wrong things, I am talking about deep and little details, that right now only human developers can craft.
4. [How AI is changing the Software Hiring in 2025](https://www.finalroundai.com/blog/software-developer-skills-ctos-want-in-2025)
    - Being low on confidence and not fake ego are so valuable, being humble has never been rarer and rewarding then ever.
5. [Waiting for a Follow-up](https://theweeklydbrief.beehiiv.com/p/waiting-on-a-follow-up-here-s-what-i-do-96fc6d7b522603e6)
    - Man, this week, I am reading and watching stuff, almost as I do them, it was the analogy thing, now I am waiting for a follow-up, it ended but still, feeling like I am walking into coincidences right and left this week.
    - But true, looking outward helps. It creates a sense of motion and drive us from the overthinking and procrastination that would have consumed and stalled us.
6. [How Cloudflare’s R2 actually works](https://developers.cloudflare.com/r2/how-r2-works/)
    - It now makes sense that how R2 is ideal for storing and frequently accessing large amounts of binary or media like data.
7. [How to get a job at Canonical](https://ubuntu.com/blog/how-to-get-a-job-at-canonical)
    - Care, Attention to detail and skills are quite the standard in any big tech company.
    - Interesting and exciting to see Canonical not using AI, at least there is hope.

## Watched

- [Learning Pydantic AI in Marimo](https://www.youtube.com/live/ujQjqqBka-8?si=SDf9b7waWu30c5yR)
    - This was the video of the week, I learnt a lot of stuff. One being, marimo is something I need to try as fast as possible and make it my go to tool for anything prototyping in python.
    - The discussion around Agent and LLMs is so beautiful.
    - Adding `Type_1 | str` is so so useful, It just opens up so many possibilities. The entire Agent concept could be made possible due to this.
    - Because if you are expecting a type in return but there is no enough information to return, the model has to forcefully return a value which might not be great, by just adding a | str it would then ask the user to fill in the details. Extremely well designed library.

Double click to interact with video
- [“Software is changing again by Anderj Karapathy” Review by the Primeagen](https://youtu.be/vDWaKVmqznQ?si=fiu9XxFxCBN8n92K)
    - Analogies man, they are everywhere.
    - Andrej drew a few analogies like OS and Electricity with AI and LLMs, which makes sense but Primeagen rightly said that he could come up with any thing and compare it with anything if he has too, but understanding deeper just breaks the analogies, which was the reason I was writing the before mentioned post, all of this clicked to me.

Double click to interact with video
- [Why Programs use Stack, Heap and other memory segments](https://youtu.be/EXIxAPITb7U?si=WenZsZqxtafWFBVG)
    - Now, I got a better understanding of what heap and stack actually means.
    - Stack would be useful for quickly adding function calls, and heap for accessing long term variables on the memory which might be fragmented.

Double click to interact with video

## Learnt

- SQL
    - SQLite doesn’t have TRUNCATE
    - .read is a handy way of reading sql queries in a shell environment
    - CASE WHEN <CONDITION> THEN <value> ELSE <value> END
          - This is the statement I used to filter values in a aggregate functions like COUNT, SUM, etc.
          - Really helpful to know this as it could be used with certain conditions to count specific elements but not for the entire query. Or there are different conditions for each column to get out.
- PostgreSQL
    - Timestamp and Datetime
          - Use named timezones as far as possible and stick to UTC
          - If you are using both date and timestamp better use datetime
    - UUID
          - Are efficient compared to text type
    - Integers | Numeric | Floating point
          - Integers are fast, range from 2, 4 and 8 bytes, super accurate and fast
          - Numeric > extremely precise but slower
          - Floating point > fast with a little less precision
    - Don’t use money type when dealing with money

## Tech News

- [XAI releases Grok 4](https://docs.x.ai/docs/models/grok-4-0709): The model is the smartest model to date, and also the most expensive. Interesting quirk is that, if you ask for its opinions, it dials up Elon Musk’s Tweets to find his opinion. Wired quirk but maybe *unintentional*
- [Cursor Pricing drama](https://techcrunch.com/2025/07/07/cursor-apologizes-for-unclear-pricing-changes-that-upset-users/): Cursor just updated / clarified their pricing for the individual plan and being costing a bit high for unlimited usage.
- [Mistral updates the Devstral](https://mistral.ai/news/devstral-2507) (Coding LLM) small and medium variants: This makes the Local Coding model race a bit heated, as Devstral Small stands neck to neck with Gemini and GPT 4.1 models.
- [Perplexity launches Comet browser](https://comet.perplexity.ai/): Now AI is getting into everything. I am not sure when this will end. People are just too eager to cram AI in the silliest of things.

As expected, there was a model release this week, if you have been following the trend, we are nearly getting a new model every two weeks. And for the past two weeks, there was no major model drop and this was anticipated. Next week Open AI might be cooking something.

That’s it from the week 50, we are close on the one year anniversary of this newsletter. Wow! a whole year has been gone in writing these newsletters, didn’t expect it to be that quick.

[Leave a comment](%%half_magic_comments_url%%)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](%%share_url%%)

---

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-753) (#753 edition) , and for software development/coding articles, join [daily.dev](http://daily.dev/).

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
