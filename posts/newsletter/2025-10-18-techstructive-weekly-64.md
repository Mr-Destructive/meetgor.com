{
    "author":"meet",
    "date":"2025-10-18",
    "post_dir":"newsletter",
    "status":"published",
    "title": "Techstructive Weekly #64",
    "type":"newsletter",
    "slug": "techstructive-weekly-64"
}


## Week #64

It was a pretty busy week, not from a work perspective but from the chaotic household perspective. It is Diwali this weekend and next week. So, the preparations, the cleaning of homes, were done past weekend. It actually took my whole Sunday, sweet Sunday, and even half of Saturday. Hh! I read a lot of books, though. I found a great book that fits the vibe I wanted while being in the atmosphere of Diwali.

At work, I was busy with simple log aggregation and number crunching for all the hard work we have been doing for the past 4-5 months. The fruits and results of that work were satisfying on a personal level, but there are always improvements and gaps people find, and the harsh lessons I have learnt from the little life I have spent. I know how to stay positive and happy, I don’t need someone’s appreciation for that.

### Loggers Clog

I almost gave up trying to get logs from the GCP Log Explorer. The UI is super slow. I have to be careful I don’t keep multiple tabs open, editors open, otherwise it clogs the CPU, and I have to restart the system. I have 8 GB of RAM, who on earth uses Slack, Google, and VS****, do? They can’t bear just being in the idle state.

I finally decided, I need to write my own TUI for GCP Log Explorer. I searched but found none. I asked LLMs to find me if any, but no result. I had to sit down and build it. But then, I don’t have a GCP account. I know I can create a free trial, but I want to save it for some other day, when I will probably start my startup. Maybe. So, how do I really test it? Testing with the Work credentials will be like bombarding too much, and I would probably livestream building that. So, what do I do? Use other Log providers and mock GCP Log Explorer? Maybe...

I’ll have to sit and decide what and how to build that Logger TUI. This weekend its hard to start it, but next weekend probably sounds like it.

### Quote of the week

> “Darkness isn’t real. It’s just where light hasn’t reached yet. Be that light”

Diwali is a celebration of light, we light our homes and surroundings and the places we consider our home, be it work or the palce where we live. Its not about buying new things, its about opening the doors that you are keeping yourself shut, its the time to dust off things and bring in the memories, the joy, the power, the flame, the light back.

---

## Read

1. [Notes on switching to Helix from Vim](https://jvns.ca/blog/2025/10/10/notes-on-switching-to-helix-from-vim/)
    - I like this approach of adapting to tooling. The author had clearly listed the reason, the main crux of why he chose Helix over Vim in the first place.
    - I like that the language server is built in, just like PyCharm is for Python. VS Code, you have to install plugins or extensions. Vim as well you’ll have to configure LSPs and plugins. Helix is a good middle ground, I think.
    - No tool is perfect, and it obviously has some quirks and things that might not please everyone, you just have to get used to them.
2. [Syntax highlighting is a waste of the information channel](https://buttondown.com/hillelwayne/archive/syntax-highlighting-is-a-waste-of-an-information/)
    - This is unique thinking. I like it, I can see how it can help developers understand and read code. But having that specific mode toggled is quite challenging and worth digging into.
    - Tree sitter is quite the thing that we can use here, but toggling different modes in different contexts is quite absurd, I think.
    - How would you know which mode works in the given context? I guess you’ll have to try a few things and get a sense of what you are trying to learn from the code.
3. [LLMs are getting better at character-level manipulation](https://blog.burkert.me/posts/llm_evolution_character_manipulation/)
    - Its evident from the test that newer and larger models are better at generalizing Base64 encoding and decoding. So that implies they will get better at character-level manipulation and analysis.
    - Sadly the how many r’s in strawberry problem will be solvable by LLMs
    - Thinking is out of the equation, the crux here is the tokenisation, the better sense of the word you have, the better it understands, but the fine balance between less and more context is critical, and I think it is still being fine tuned to get a sweet spot.
4. [Secret Life of Local First Value](https://marcobambini.substack.com/p/the-secret-life-of-a-local-first)
    - This is lovely. So well explained what CRDTs are. It’s like a log of what happened in a row of a table. Like column-level details of updation/insertion and deletion. It makes sense now.
    - The metadata table is the crux of this structure. What would happen if the database itself crashes? That is unlikely, I think. SQLite cannot crash at least locally. Nice thinking here.
5. [There are no programmers in Star Trek](https://www.i-programmer.info/news/99-professional/18368-there-are-no-programmers-in-star-trek.html)
    - This makes me sad, if there are no real programmers, how are their services operating? If computers just work and do what they are told, why is there even a service? Why just not complete everything and be done with it?
    - I wonder if 2027 be a year like that, where you say something and vibe coded mess some how works, how far are we from the reality to here?
6. [Just the grind won’t get you there](https://swizec.com/blog/the-grind-wont-get-you-there/)
    - True, you need to speak up. Just little conversations, little efforts, small nudges go a long, long way. You never know what you would be leading to. Just keep your mind open and create a positive environment around you.
7. [Craft, not fame, makes your story worth telling](https://herbertlui.net/craft-not-fame-makes-your-story-worth-telling/)
    - This is not a technical read, but worth mentioning here.
      > If you’re concerned that your story is too boring, put some effort into making it matterEverybody has a story to tell. You just need to pick the right ones, and to give it some meaning
    - So true and well said here. If you really want to tell a story, you will write one, else, you will find excuses or make a horrendous mess of AI and average slop

## Watched

- [Claude Haiku 4.5: The best model from Anthropic for cost to intelligence figures](https://youtu.be/iES9r7AZP1s)
    - Maybe its a big leap in terms of intelligence to cost ratio. We might get more cheaper models with more intelligence.
    - Open Models especially the China Models are pushing these AI Labs in a better direction. GLM and Kimi K2 have forced Anthropic to release this to stay competitive and able to sustain their growth.
    - Phew! What times to live in as a developer.

- [What is Apache Iceberg](https://www.youtube.com/watch?v=TsmhRZElPvM&t=34s&pp=ygUOaWNlYmVyZyBhcGFjaGU%3D)
    - This was a good overview of how Iceberg is actually a bridge from the old data lakes to the modern microserver architecture.

- [What is Apache Kafka and where it’s headed?](https://youtu.be/9CrlA0Wasvk)
    - Why everyone needs to inject AI to help someone use AI?
    - WHY? Kafka is just a stream processing library, why just not keep it as is, why try to slop it with AI and sell it like a AI support and what not.


## Learnt

- Always test after resolving and committing merge conflicts
    - I just wrote a little thought about it [here](https://www.meetgor.com/thoughts/test-after-resolving-merge-conflicts/).
    - Yes, we accidentally pushed certain changes that were kind of broken state and led to around 100 requests being errored out. Just because of some bad state after a merge conflict.
    - GitHub’s feature to merge conflicts is nice and cool, but the syntax highlighting and warning couldn’t be neglected. The actual pipeline errored out but went unnoticed due to the large number of changes being pushed in the PR.
    - It was a small bug, nothing major, but it shows how small things can lead to disastrous outcomes if not tackled quickly.
- Vim Macros Optimizations
    - If you have a macro to run in a more than 1000-line log or text file, don’t save the file as the instruction in the macro
    - This led to my system crashing after I ran a log cleanup macro 200 times on a 2000-line file
    - I just then removed the command that saved the file (wrote to the disk), which saves an io operation, and it exponentially saves time and memory.
    - I kind of wish a TUI existed that would open the current log in Vim, that’s one idea I have as a feature to have in my project I discussed above.

## Tech News

- [Amp is now free](https://ampcode.com/news/amp-free) (with Ads)
    - This is big. I am making this my default AI coding assistant from next week. I would be trying out a few things.
- [Nanochat](https://github.com/karpathy/nanochat)
    - Looks like a great way to learn about how models are trained. A very thought-through learning resource as well as a production-grade model builder.
    - Unfortunately, I don’t have GPUs, so I cannot run it to test my hypothesis.

---

Happy Diwali!

That’s it for this week. Hope you have a great Diwali, happy new year (Indian business new year, Vikram Samvat 2082)

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-767) (#767th edition) , and for software development/coding articles, join [daily.dev](http://daily.dev/).

Happy Coding :)
