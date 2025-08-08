{"author":"meet","date":"2025-08-02","post_dir":"newsletter","status":"published","title": "Techstructive Weekly #53","type":"newsletter", "slug": "techstructive-weekly-53"}

## Week #53

A pretty slow moving week, but a lot of consumption. I watched and read a lot of resources on SQL, databases and backend in general. I want to understand SQL to learn what is the fundamental unit in CRUD apps, which I think is a SQL query.

Writing SQL query by hand in 2025 might be obsolete, but that 2% or 5% of the queries are pain in the ass to get LLMs write, that’s when the knowledge the depth, the legs of the PI shaped learning helps and make you feel alive, after weeks of existential crisis. I’ll keep learning SQL for the time being.

Admitting that I didn’t push any code to prod this week, it was a bit of struggle for me to get things in shape after loosing touch, I did a ton of debugging, but didn’t push any code. I’ll be honest about it, as there are days where you are just a helping hand or a gazing at the things unfold before you, because it’s ain’t my time or space to do it. Next week, probably will be breaking prod.

### Quote of the week

> “Without your involvement you can’t succeed. With your involvement you can’t fail.”
> 
> — APJ Abdul Kalam

Immerse yourself in the process, because you can’t win without, and the only thing in the way is yourself. I have involved myself in programming completely for 3 years, and have not failed in a literal sense, I have a great job, have good motivation to work towards, yes I made mistakes and continue to learn from them and make more different ones, that’s the process. If you put out, you know, you learn, you get better, the more you put out, the more luck doubles up each time.

Most people get exhausted after the upteenth attempt, but they might be just one more step away from success, success should never be the only goal, if it is, what after you achieve it? The process, the madness to do it, should be, the curiosity that will flicker in the toughest of times will carry the weakest to the peaks and the absence of it can make the strongest tumble down the cliff.

## Created

- Wrote
    - [TIL: TMUX Scroll mode select and copy](https://www.meetgor.com/til/tmux-scroll-mode-select-and-copy/)
          - I was in dilemma, I sometimes use Ghostty or default terminal, sometimes with multiple windows, sometimes different tabs, sometimes zellij and sometimes tmux.
          - I just forced myself to use TMUX and ended on the problems or flows that I have been avoiding to fix,failing to find and navigate around like a wizard, but took a few minutes to get the mess out, and here we go.
          - This was just a few key strokes away namely, `prefix + [ `to enter the scroll mode, Ctrl + Space to enter selection mode, and Enter or Ctrl + J to copy to clipboard. Rock and rolling with log driven debugging.
    - [Thoughts: Kubernetes isn’t for you](https://www.meetgor.com/thoughts/kubernetes-isn-t-for-you)
          - I agree to this post, in certain situations. Especially for me if I am starting out in tech, I probably would not give advice to learn Kubernetes. Just get dirty with the normal deployment options, understand the pain points and then switch to K8s or others as needed.
          - This post lists down the points that people use Kubernetes just for bragging or showing that they know what they are doing, or just using it to show that they are cooking. But in reality if you are a small scale team, with couple of services, using Kubernetes makes no sense.
    - [Thoughts: Just fucking use kubernetes](https://www.meetgor.com/thoughts/just-fucking-use-kubernetes/)
          - I agreed to this post, too, as I said it depends. If you purpose is to learn Kubernetes, be a DevOps guy, be a SRE, or some cool tech guy, nerding, curios about it, JUST F-ING LEARN KUBERNETES.
          - This post made really good attempts to pull the trigger for me, I’ll quote waylon walker for a giga-chad quote
            > “But my App is small, So is you ambition”
          - That just hits home, for me at least. If you think your app is a toy project, you are probably right and wrong, how do you know that if you haven’t deployed yet and people haven’t used it yet.

## Read

1. [Overthinking GIS](https://scottsexton.co/post/overthinking_gis/)
    - This just triggered some neurons in my brains, some horses running. I got a bit of interest in exploring GIS data. Want to play with it, and find interesting details.
    - This post highlighted the ways to use GIS data to get the usability of a land, basically which are fertile and usable for farming I think. But the technicality in which it was explained was clicking the right knobs at the right time with the perfect steps and images.
    - Would be certainly writing a post in this style sometimes soon.
2. [HTTP VS Websockets: The breakthrough moment that clicked](https://tobennaoduah.substack.com/p/websockets-vs-http-the-breakthrough)
    - Such a honest, humble and insightful post. HTTP as a delivery truck, Websockets as a telephone! Such a relatable example, this shows the author got it right and is able to connect it well.
    - I can’t agree more to this tip
      > Try building a tiny app with both. Make a little dashboard that gets real time updates on something simple like stock prices or server load. Build one with HTTP polling. Build one with WebSockets. You will feel the difference, not just see it.
    - Back when I was a freshman in college, I tried doing this kind of thing for my world atlas chat app game, and failed badly, calling database for each message to send over the HTTP, what a idiot I was, but then, google searched about this and found plethora of articles about websockets and what not.
    - While making that, I realised the pain points getting solved with websockets and everything just clicked.
3. [More than Code](https://deadprogrammersociety.com/2025/07/more-than-code.html)
    - Reading code in the age of AI, is gold and the only thing I think people spend most of the time, with other being vibe-debugging which might be fair less then vibe-reading.
    - Reading code should become like instinct, just by gazing, you should be able to smell bad code and sniff the bugs out. This is not easy, it comes with practise and years of slog-debugging, first debugging the human code then go to vibe-debugging.
4. [Learning GRPC Completely in Golang](https://www.pixperk.tech/blog/learn-grpc-completely-in-golang)
    - This explains what GRPC, protocol-buffers and the connecting technologies behind them in Golang.
    - This gives a great overview of what and how the APIs are created.
5. [Death by AI](https://davebarry.substack.com/p/death-by-ai)
    - Woah! AI Overview mode in Google is terse at certain things. But for programming, I think it works a charm 95% of the times. yes there are pretty bad hallucinations too due to reddit and slop debates.
    - Haven’t thought that it would mistaken a person with other person’s name and call it dead! Hillarious.
6. [Why I do programming](https://esafev.com/notes/why-i-do-programming/)
  > For me, programming has always been more than a skill. It’s a way to explore, to tinker, and to satisfy curiosity.
    - This is what programming is about, using curiosity to find the solutions to the problems. Curiosity is like a fire that keeps you warm in the winter.
7. [Why GenAI Infrastructure feels backward](https://soypetetech.substack.com/p/why-genai-infrastructure-feels-backwards)
    - I agree to this, this post can’t read this entire thing, but can understand the point of view
    - Python and Javascript seems to be taking the forefront in the infrastructure side of things. All AI-Labs first launch SDKs and Packages for these two ecosystems, but never the others which are suited for them like Go or Rust.
    - It’s time to change and make a difference in this revolution.
8. [What are JSON Web tokens (JWT)](https://blog.algomaster.io/p/json-web-tokens)
    - Find me a better JWT tutorial then this, I’ll wait.
    - Explained the problem then developed the intuition for the reason why JWTs exist. Very well explained, detailed and the diagram made it perfect for visualising the flow. It also mentioned the best practises, ticking all the boxes for a great article.
    - Must read for beginner getting to understand JWT Authentication after learning Session based authentication.

## Watched

- [Anatomy of a Request: A deep dive of a http request processing from the Backend side](https://youtu.be/s0r3Aky9I5g)
    - Woah! That is a ton of computation.
    - On Client: Creating the payload, encryption (write copy), loading in kernel space, sending the data
    - On backend: Received the data, reading to the user space, decryption, decoding (serialization) of the body.
    - So many steps are there, the speaker rightly said, its a fascinating field, the more you go deeper, the more stuff is there to explore and learn.

Double click to interact with video
- [Writing a Text Editor - Computerphile](https://youtu.be/g2hiVp6oPZc)
    - Interesting that text editors use Gap buffers. Its like a temporary register (block of memory) used to append text characters while the user edits (adds) to the file and then it gets saved, the remaining empty part is truncated. Clever data structure.
    - He explained it so nicely, the approach, the problem and then a new intuition, again a limitation, then a proper intuition.
    - Vi uses linked linked like data structure for editing, very interesting.

Double click to interact with video
- [Why do databases store data in B+ Trees](https://youtu.be/09E-tVAUqQw)
    - Everything is about intuition it seems, you see a problem, you think for a while, you think of a ideal scenario and you just scramble up a solution by adding the good parts and discarding the bad parts.
    - This was the same, for why database use B+ Trees. Because we need to optimise for any arbitrary access for data as well as for range queries.

Double click to interact with video
- [SQLite: How it works: Richard Hipp](https://youtu.be/ZSKLA81tBis)
    - What a banger of a presentation and talk. Explained so much, in depth, in such a short time. It helped me understand what SQLite actually is, it’s a parser + virtual machine to run the core part and basically the fopen function in C to actually perform the operation.
    - One unique insight here is
          - Reading 10 files content from disk is slower than reading those file contents from SQLite
          - Why? Because the database file is opened once and the reading happens in that instance only, data is stored in pages (fragments of memory), so it’s just a matter of reading bytes at a specific order.
          - But reading 10 different files on disk will make you use fopen 10 times, and that is slow!
          - 200 IQ move from SQLite team, have never seen such a beautiful solution to almost all the problems in the data world.

Double click to interact with video
- [The real reason you can’t get a job](https://youtu.be/SPwPpsXpZfg)
    - More work, more luck
    - Curiosity to learn, outperforms desperation to get money

Double click to interact with video
- [Simple Joy of programming Course announcement](https://youtu.be/EV13CNrq4ZA)
    - It’s a great commitment to teaching the fundamentals.
    - Not just fundamentals but building on top of the strong foundation.

Double click to interact with video
- [Sync Engines and Local Data](https://youtu.be/1uVR5X7HpI8): Discussion of different database sync providers
    - This actually made me a bit curious about sync engines further. I was in confusion when I heard about it from Theo as he used it for T3 chat. It didn’t made sense at that time. It still doesn’t, as why syncing is required in a chat app, all the data comes from the backend, there is no processing on the frontend?

Double click to interact with video

## Learnt

- Getting into scroll mode in TMUX, selecting text and copying with vi-like keybindings.
- You can’t do a unpacking in a sql query when using nested query for more than one column returned
    - Like example i have a query like this
      ```
      SELECT 
        s.id, 
        s.package_id,
        address_id, 
        a.address AS from_address, 
        (
          SELECT 
            address,
          FROM 
            addresses 
          WHERE 
            id = p.to_address_id
        ) AS to_address 
      FROM 
        scans AS s 
        INNER JOIN packages AS p ON p.id = s.package_id 
        INNER JOIN addresses AS a on a.id = p.from_address_id;

      ```
    - But let’s say for some reason I wanted to also get the to_address_type like a column from the addresses table
    - You might try to over-optimise the queries and try something like this
      ```
      SELECT 
        s.id, 
        s.package_id, 
        action, 
        contents, 
        address_id, 
        a.address AS from_address, 
        (
          SELECT 
            address, 
            type 
          FROM 
            addresses 
          WHERE 
            id = p.to_address_id
        ) AS (to_address, to_address_type) 
      FROM 
        scans AS s 
        INNER JOIN packages AS p ON p.id = s.package_id 
        INNER JOIN addresses AS a on a.id = p.from_address_id 
      WHERE 
        address = '900 Somerville Avenue' 
        AND s.action = 'Drop';

      ```
    - And ERROR, you can’t do that
    - This bit right here
      ```
      (SELECT address, type FROM addresses WHERE id=p.to_address_id)
      AS
      (to_address, to_address_type)
      ```
    - This is not feasible in SQL, you can’t unpack multiple columns from a subquery directly and alias them inline in a single SELECT clause.
    - Well I have to do it this way then, duhh

  ```
  SELECT 
    s.id, 
    s.package_id, 
    action, 
    contents, 
    address_id, 
    a.address AS from_address, 
    (
      SELECT 
        address 
      FROM 
        addresses 
      WHERE 
        id = p.to_address_id
    ) AS to_address, 
    (
      SELECT 
        type 
      FROM 
        addresses 
      WHERE 
        id = p.to_address_id
    ) AS to_address_type 
  FROM 
    scans AS s 
    INNER JOIN packages AS p ON p.id = s.package_id 
    INNER JOIN addresses AS a on a.id = p.from_address_id 

  ```
    - What a long query!
    - By the way, this is one of the questions in [CS50 SQL Problem set 1 packages](https://cs50.harvard.edu/sql/psets/1/packages/) section.
    - Any better way to do this? drop them in the comments or hit me up on my socials, will be completing more challenges this weekend.

## Tech News

- [Google Gemini 2.5 Deep Thinking mode on app](https://blog.google/products/gemini/gemini-2-5-deep-think/)
    - This looks something interesting, Google keeps pushing the boundaries for what LLMs are capable and making it more accessible.
- [NotebookLM launches Video overview mode for notebooks summaries](https://blog.google/technology/google-labs/notebooklm-video-overviews-studio-upgrades/)
    - This is a good feature, not great! I tried it and it’s decent, cuts in between, just text and arrows matched up like a presentation deck for the given problem. Not much value for a deeper dive, but really great to get a gist and overview.
    - I can see this being used for news aggregation or summarising things from a lots of similar sources.
- [ZAI release GLM 4.5](https://z.ai/blog/glm-4.5)
    - Another chinese model that is a serious contender for coding or general purpose tasks. That too local and open source models.
- Qwen releases 3rd version of Instruct, Thinking and Coder models
    - They launched quite a few variations, 235B for thinking and instruct
    - Also 480 and 30 Billion variant for the Coder model
- [POE introduces the API for AI Models](https://poe.com/blog/introducing-the-poe-api)
    - I have used POE in the past, and this looks like a great addition to test out a few SOTA models for free with some limitations via the API.

---

Phew! a lot of open source models are cooking! Everything is intuition and curiosity, and AI is still a thing to wrap our heads around, but the fundamentals are the same.

See you in the next one!

Happy Coding :)

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-756) (#756th edition) , and for software development/coding articles, join [daily.dev](http://daily.dev/).

[Leave a comment](%%half_magic_comments_url%%)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](%%share_url%%)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
