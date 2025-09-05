{
    "author":"meet",
    "date":"2025-09-06",
    "post_dir":"newsletter",
    "status":"published",
    "title": "Techstructive Weekly #58",
    "type":"newsletter",
    "slug": "techstructive-weekly-58"
}

## Week #58

What a week, almost broke prod for half an hour! On Monday, can you believe it? I didn’t but that was a bug. Patched it and moved on. Had a consistent posting and learning sessions daily for an hour about SQL and SQLite. Feels good.

Had an situation where knowing javascript was a big deal and I don’t feel like writing javascript, it just doesn’t appeal to me, I guess I have to amp that up. Anyways, one more reason added for hating javascript, its unavoidable.

> The only painful thing about Javascript, its unavoidable

Yes, back to SQL, I am absolutely feeling energetic to talk and learn about SQLite daily. I think we’ll be 100 posts in no time. I am thinking about creating and completing a project this weekend. Let’s see if we are able to do it.

### Quote of the week

> To live is to wrestle with despair, yet never allow despair to have the last word
> 
> — [Cornel West](https://www.goodreads.com/quotes/19982-the-country-is-in-deep-trouble-we-ve-forgotten-that-a)

We all have problems, but we can’t let it define us, rather we will be defined on how we reacted to them. You can give up, but either you’ll be unsatisfied or below average. You have to go through the pain in order to bear the fruit, and even if the fruit doesn’t wag your way, so be it, if you enjoyed doing it (writing or reading in my case) then I believe that was worth it. Else what is the purpose of life? Who are we really and why do we even breathe. Just enjoy the thing you enjoy there is should be no measure for productivity when it comes to living.

## Wrote + Created

- 7 More posts in the SQLog Series or SQL Blog
    - [Change directory dot command](https://www.meetgor.com/sqlog/sqlite-cd-dot-command/)
    - [Output Mode dot command](https://www.meetgor.com/sqlog/sqlite-mode-dot-command/)
    - [Output Mode Separator dot command](https://www.meetgor.com/sqlog/sqlite-mode-dot-command-separators/)
    - [Header toggle dot command](https://www.meetgor.com/sqlog/sqlite-dot-command-header/)
    - [SQL: Create basic table](https://www.meetgor.com/sqlog/sqlite-create-table/)
    - [SQL: Bare bones table in sqlite](https://www.meetgor.com/sqlog/sqlite-create-table-bare-bones-table/)
- Updated my blog SQL codeblocks to support latest SQLite version 3.49 checkout [here](https://www.meetgor.com/sqlog/sqlite-scalar-function-3-valued-iif-and-if-scalar-function/)
- A demonstration video explaining about the blog I created [link](https://www.loom.com/share/2695ff73fac940d1bd2d3b6432f1cb2a?sid=9fceb6f8-ec01-44c0-83f1-f41a7171adc0)

## Learnt

- SQL week probably some 7, I am loosing count, let’s forget the week count from now
    - SQLite has a rowid as a default primary key for distinguishing each row for a table. If created a primary key, then rowid is an alias for that column
    - Different modes in sqlite and how they get affected with the headers dot command.
    - SQL is a relational algebra under the hood, select, relation, union, projection, intersections, difference, joins,product.

## Read

1. [The modern job hunt: Part 1](https://thedailywtf.com/articles/the-modern-job-hunt-part-1)
    - This was a heart warming and hope-giving post. For a person struggling with jobs, finding a new hope is something that this post absolutely gave in abundance.
    - The comparison of a walk to garbage collection is so funny. Job search as a searching algorithm.
    - I also like the two different contrasting thoughts and way of thinking about the situation, from
      > *This job search is a hopeless, unending slog!*TO*Yes, this will end. Everything ends.*
    - This hits home. Truly a great post, Eagerly awaiting for a second part. Subscribed to the RSS feed only to know the feed isn’t updated.
2. [An illustrated guide to OAUTH 2](https://www.ducktyped.org/p/an-illustrated-guide-to-oauth)
    - This made everything click, not quite everything. I want a better example. I will write a post where I would relate it to a much more simpler and broader example. The Chase bank and other finance provider is a bit not known to me, I want to make the example clear.
    - But really the explanations in this blog are really sketched out and make a lot of sense.
3. [Why Sync?](https://adamwiggins.com/posts/why-sync/)
    - I see a lot of sync things, like sync engines, local database sync, I even myself had this issue on blog sync with database and github repo.
    - My manager hit me, a quick sync? That’s a different sync of course. Sync!
4. [A programmers guide to stress](https://tkouleris.eu/blog/a-programmers-guide-to-stress)
    - Wow! What a banger of a post, what gold pieces of adivce
      > we cannot get rid of stress. What we can do is change our philosophy and the way we see things.
    - This blog advocates to meditate 10 minutes a day, it helps maybe. I need to try. But that comes after doing something, you act first, then if the result is not coming, you get anxious and frustrated if it failed, in that case, you need to calm down, you did whatever you could, but now you need to wait for it to resolve or learn from the failures.
    - Getting anxiety from the current problems is true, I cannot keep my mind away. I need to breathe and let it go, will try over the weekend and continue if it makes me happier.
5. [SQL Needed structure](http://scattered-thoughts.net/writing/sql-needed-structure/)
    - This is a bad take, What is more structured than SQL? If you are trying to get heirarchial data in SQL, you will be comprising redundancy. Like literally for the ease of the access, you will duplicate json blobs everywhere. I work in that environment and know the pains of those queries. I won’t recommend MongoDB ever to anyone. Just use structured data, it would be a one time investment on schema design, but it sustains and is much easier and predictable.
    - It could have been other things like sharding, I would have agreed, but this nope.
6. [Static Site enable a good time travel experience](https://hamatti.org/posts/static-sites-enable-a-good-time-travel-experience/)
    - A really cool observation. Anyone can pick a commit hash and simply run a command to obtain the static files. Or better just checkout to the output branch and go back to the date. Its like a time machine.
    - I love static site generators.
7. [You have to feel it](https://mitchellh.com/writing/feel-it)
    - Feel it, Just completing tasks, merging PRs, lists of metrics improved won’t help, if you don’t feel it.
    - Just creating for the completion is fine, sometimes needed, if you haven’t gone through the whole feeling, but not always, you need the drive to complete it in a sustainable way.
8. [How to learn System Design](https://medium.com/@himanshusingour7/how-i-learned-system-design-d7444d454367)
    - System design, something I can see myself struggling a lot if in a interview. Need to start this approach of creating one my approach in designing existing system.
9. [Authorisation explained](https://levelup.gitconnected.com/authentication-explained-when-to-use-basic-bearer-oauth2-jwt-sso-c3fb0aa083ef)
    - A nice walkthrough of authorisation types, role based, attribute based and access lists.
10. [Vibe Coding is Legacy code](https://blog.val.town/vibe-code)
    - This is the realisation I had last week too. I used agent like the code based agents to write a script that I won’t care about, some code that I won’t care about after a while. That type of code is a vibe.
11. [Reading with AI](https://arnoldkling.substack.com/p/reading-with-ai)
    - People are pushing AI in wired direction. By using that for activity like reading, you are defeating the whole point of reading. Non-ficition or ficiton, if you put down the effort of going through word by word, the context, the feel to the LLM, you are doing yourself some deservice.

### Interesting Links

- [30 Minutes with Strangers](https://pudding.cool/2025/06/hello-stranger/)

## Watched

- [Database System: Relational Model and Algebra](https://youtu.be/7NPIENPr-zk)
    - Prepared a detailed notes as I was watching
    - data model - how to define the relation, rules
    - schema - defining the specifics of the data model, details
          - relational (the primary kind, we only think of this as real dbs) key value - graph - document - column - array - and more
          - Initial idea: Writing cobol in 1970s to get data, literally telling the db to how to get the data. Not ideal as you don't know if that would be ideal way to get data in all cases, example artist and album, you write for each artist, find albums, but what if the number of albums by each artist keep growing and more people ar accessing different artists data together, so not sure how to get data for any query. In SQL however, we dont say how to get data, we say what and from where to get data
          - Relational Model
            - Structure
            - Integrity (constraints)
            - Manipulation
          - Its upon the database system how to query/mutate the data
          - Components
                  - Database storage(bits)
                  - Physical storage (pages, files, etc)
                  - Logical Schema (schema, constraints)
                  - External Schema (common table data)
                  - Application
          - Everything below application is a db system
          - Relation is a set, relation of attributes that represent entities
          - Tuple is a set of attribute values in relation. Primary key, uniquely identify a tuple in a relation. Foreign key, related tuples(attributes) across relations. Constraints, conditions must hold for any tuple in a relation.
          - Relational Algebra
                  - Select (where conditions)
                  - projection (select with what to extract optional modify the selected values)
                  - union (all, must have same number of attributes)
                  - intersection (same but common)
                  - difference (same but difference)
                  - product (cartesian product, cross join)
                  - join (natural, without params, common ones, same attribute name, can use params like using **on**)


- [SVGs are so cool](https://youtu.be/mN--sGH97dY)
    - I really didn’t knew SVGs can be that cool, you can scale, animate and even make interactive elements.
    - I need to add some SVGs on my blog. This is absolutely inspired me.

- [Vibe coding are three types](https://youtu.be/6TMPWvPG5GA)
    - I know and care about the code
    - I know the code but don’t care enough
    - I don’t know what I am doing
    - Developers, let’s focus on the first 2 and we should be fine, the 3rd one is a luxury as a hobby.


## Tech News

- [Atlassian to buy the Browser/Arc Company](https://techcrunch.com/2025/09/04/atlassian-to-buy-arc-developer-the-browser-company-for-610m/)
    - I don’t care.
- [Warp release Code: the terminal based coding assistant](https://www.warp.dev/blog/introducing-warp-code-prompt-to-prod)
    - This an another weapon added to the arsenal we are seeing in the arms race of the AI coding assistant. How many we have, I forgot the count.
- [Anthropic raises $13B Series F at $183B post-money valuation](https://www.anthropic.com/news/anthropic-raises-series-f-at-usd183b-post-money-valuation)
    - I don’t know what that means, but is a big thing according to the big tech bros.

Nothing happening for a few weeks now, finally there is a slow down in the models and all this hype. I want this hype to be minimum as possible for a stretched duration of time, as it helps learn and do things peacefully.

---

Happy Coding :)

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-761) (#761st edition) , and for software development/coding articles, join [daily.dev](http://daily.dev/).

[Leave a comment](%%half_magic_comments_url%%)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](%%share_url%%)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
