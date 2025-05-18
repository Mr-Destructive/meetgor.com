{"author":"meet","date":"2025-03-21","post_dir":"newsletter","status":"published","title":"Techstructive Weekly #34","type":"newsletter"}

 
## Week #34
 
 It was a pretty busy and productive week. After three hectic weeks of research, hacking, scripting, implementing, Google Sheets, and Python Vulture coding an API, I made a few releases. Yes, that is the normal role of a software engineer at a startup. I get to wear multiple hats: analysing, writing, debugging, scripting, testing, and obviously communicating effectively.
 
 I was able to complete the implementation successfully. The release somehow happened on a Friday evening, which I am skeptical about, but it happened without any hiccups, a few minor ones. It will all be well if it ends well. 
 
 Last weekend, I finally launched my own SSG and Blog, which I built with it. I am trying to move the site from `meetgor.com` to `dev.meetgor.com`. The site is running on GitHub pages and is built with the Burrow SSG.
 
 A week, productive, and happy about whatever I learned so far. A lot of Python scripting, with pair programming with LLMs, google Sheets for finding accuracy numbers, API shenanigans, and some serious mode all the time. I like being serious, that’s my trait, I am not bragging, just saying.
 
 Next week, I will be hopefully writing more and uploading a video or two about some stuff that I have been tinkering with.
 
 ## Quote of the week
 
 > "It’s not whether you get knocked down, it’s whether you get up."
 >
 > -- Vince Lombardi
 
 I have fallen several times, in the past few weeks, but I have always got up and kept going. It was hard, but facing guilt is more hard than the pain of falling. I can sit and cry, but that won’t get me anywhere, I am here to change my destiny, not to wait and let it unfold. I cannot change winds, but I surely can control the way I move.
 
 Things are a bit different in tech right now, LLMs are everywhere, and they probably can do at par with a junior developer, but one thing is for sure they can’t. The curiosity, the internal fire that keeps a person alive, awake at night, forgetting the surroundings when debugging, this LLM can’t do yet! If you fall, we all do, get up and forget everything that happened, remember you can drive it, drive it better, change the way you react to failure, and  don’t let it take over you.
 
 ---
 
 ## Wrote
 
 Wrote a few things, want to write more.
 
 * [Thoughts: Zellij Scrollback edit mode](https://dev.meetgor.com/thoughts/zellij-open-scrollback-edit-mode/)
 * [TIL: Format JSON file in Vim with JQ](https://dev.meetgor.com/til/format-json-in-vim-with-jq/)
 
 I will be writing more since I have set up my blog at dev.meetgor.com, it is served by my own Static Site generator using Golang and Netlify Cloud Functions hosted on GitHub Pages. I still need a few things to sort out which will do on the live stream.
 
 I created this SSG in Golang from scratch (except for  the markdown parsing) on live streams. This also has a CMS or an API that can add posts to the database and GitHub pages can sync up with the database and save the post to the GitHub repo.
 
 * [Creating SSG from Scratch in Golang Playlist](https://www.youtube.com/playlist?list=PLMVgNvnU9WlGRy0FySl6Ot9M5Rtb7qopu)
 * [Burrow: SSG + CMS in Golang (not a mature and ready-to-use project yet)](https://github.com/mr-destructive/burrow)
 
 Feels really good about how the SSG came together, it's really pluggable SSG with flexible JSON as well as YAML compatible frontmatter and configs. It has an editor that creates a post and stores it in the database, the GitHub workflow cronjob picks up the last edited or created post in the last hour and creates or updates the file in the GitHub repo.
 
 ## Read
 
 * [AI LLM Agents are just graphs, frameworks over-complicate things](https://zacharyhuang.substack.com/p/llm-agent-internal-as-a-graph-tutorial)
     * This article shines like gold, it’s quite comprehensive and explains exactly what agents are. The code is precisely 100 lines of code which is surprising to me, was that intentional, or was that a good given luck.
     * LLM decides for you
     * The tools are rule-based or just procedural steps to execute
     * Agents are like graphs, node,s and edges, one pointing to different directions after taking action or deciding what to do next.
     * Pocket Flow
     * [LLM Agents are simply Graph — Tutorial For Dummies](https://zacharyhuang.substack.com/p/llm-agent-internal-as-a-graph-tutorial)
         * Ever wondered how AI agents actually work behind the scenes? This guide breaks down how agent systems are built as simple graphs - explained in the most beginner-friendly way possible…
         * Read more
         * 4 days ago · 59 likes · 1 comment · Zachary Huang
 * [Learn Go Templates: A practical guide to layouts, data rendering in Golang](https://evolveasdev.com/blogs/guide/learn-go-templates-a-practical-guide-to-layouts-data-binding-and-rendering?ref=dailydev)
     * It was a great article. It covered a lot of things in detail as well as digestible examples. I liked the explanation of template partials and functions. I might be using them in my ssg in Golang (burrow). There was one missing nuance, that is using a variable if inside a loop range and referencing with the $ to access the outer variables in templates.
 
 ## Watched
 
 1.  [AI is here, and yes you are screwed, as a junior: Article review by THE PRIMEAGEN](https://youtu.be/LXUw0xSib-g):
     * I am here for the hot takes, I read this article last week and shared some thoughts I was completely agreed to the article as I am not sure about certain things, I lack some common sense I think, but yes THE PRIMEAGEN in GOD mode opened my eyes
     * If you are with people who are working for money, is usually good according to their experiences, which is fair, it sometimes it depends on the interaction with the person rather than the nature of the person. We sometimes get used to the nature of the person after some time.
     * Use LLM but understand what are you doing exactly
     * Copilot was great bump in the productivity of developers when it launched but since now we are used to the generated code, it seems like we get bad completion or it’s bad.
     * A good analogy for LLM's ability to code, the more precise you want to do certain thing, the better the generation is, and the broader the scope of the task, it messes up. Like give me the next 2 lines of code in the current position, that is good, LLM is really good at that (not always). But if you give i want to build this project, in an x and y way with a and b technology, that is where it might get a little screwed (Devin not mentioned but I heard the screams)
     * AI SaaS clones are shit and not true things, it is more than code, relation, care the bond. If someone is earning money for the SaaS clone they build with AI, then they might be building a silly problem for a silly user.
     * The point of juniors is that in the future they will pay off the investment they made.
     * That was too much maybe, a reaction of a reaction, no it’s a note for me.
 2.  [Backend Banter with Mitchell Hashimoto](https://youtu.be/586_BAMMOQ8):
     * I have watched it halfway through and so far some hot takes
     * Rust ecosystem is a little messed up to interact with for him at least
     * Zig fixes the pains of C, the build ecosystem (handles the compilation for us), and package management (like go)
     * Project-based learning for the win!!
     * He started to build Ghostty as a fun toy project, then found some pain points and tried to fix those with an ecosystem of libraries to build an  abstraction layer of apps on top of terminal emulation.
 
 ## Learnt
 
 * Scrollback editor mode in Zellij
     * This mode is GOD Mode for Vim nerds, it is really powerful that we can explore the logs from the editor. I used it to grep and search quickly through a long list of logs and filter, count, and extract some things out of the sea of logs. Highly recommend checking it out, might be handy in certain situations.
 * Format a JSON file in Vim with jq
     * I always had to open up VS Cod\* for formatting the file, which is currently opened in neovim, since I have a low-spec laptop, VSC crashes down the machine, and sometimes also go the browser and pastes in a json frontmatter, but that sucks too. So this command using jq really solved my problem. I use jq frequently when I am outside of neovim, in the shell, however, when I am in the editor, I want it to be formatted immediately.
 * Using the difflib from Python standard library with SequencMatcher for fuzzy matching
     * I had been using rapidfuzz and other fuzzy matching tools for string matching. However, there is difflib in the Python standard library, which is really handy for quick simpler things. I found this while asking GPT for a particular thing and it gave me this difflib SequenceMatcher function to compare the two strings, and I was unsure of it at first, but I typed it and surprisingly my LSP picked up difflib, so I checked what is difflib and man that was a python standard library.
 
 ## Tech News
 
 * [Open AI releases new Audio Text to Speech API playground](https://openai.com/index/introducing-our-next-generation-audio-models/) → [openai.fm](openai.fm)
     * This looks good, really good, podcasts, videos are a bit risky at voice now. Clearly, voice is taking over with AI, text > images > voice > ??? This is going big.
 
 * [Claude adds search on the web with generated output](https://www.anthropic.com/news/web-search) (only for Pro Users though)
 
 I have no opinion on that since I have not tried it, but it seems like a great feature and must have since all other major LLMs are masters at this, especially Preplexity, Gemini, ChatGPT, etc.
 
 For more interesting articles, check out the [hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-738) for the week edition #338, for even more software development/coding articles, join [daily.dev](https://daily.dev/).
 
 That’s it from this 34th edition of my weekly learning, I hope you enjoyed it, and leave comments on what you think about some of my takes or any feedback.
 
 ---
 
 Thanks for reading Techstructive Weekly! 
 This post is public so feel free to share it. [TECHSTRUCTIVE WEEKLY #34](https://techstructively.substack.com/p/techstructive-weekly-34)
 
 Thanks for reading Techstructive Weekly! [SUBSCRIBE](https://techstructively.substack.com/subscribe) for free to receive new posts and support my work.
 
