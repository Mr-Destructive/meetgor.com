---
title: "Techstructive Weekly #115"
date: 2026-10-16
slug: techstructive-weekly-115
type: newsletter
status: published
source: newsletter
canonical_url: https://techstructively.substack.com/p/techstructive-weekly-115
description: "Reading watching the tech world come back with thinking about AI, among the other things learnt and observed in the week of 11th to 17th October 2026"
tags: ["newsletter", "substack"]
---

## Week #115

Its Navratri this week and the coming week. Its a good time to be alive actually. The nostalgia of songs, the memories attached to it, it just makes me feel good. I decided to read “Forest of Enchantment” this week. And oh boy it is a good book. I will devour it this weekend.

Apart from the reading, I also did a few writing sessions this week, as I usually do since the last 2 years. I feel great, writing is something that just calms my mind. I will continue to do so.

And for the tech side, it was a draining week, I worked on solving issues, experiments and a lot of conversations with ai, it almost drains me and makes me anxious. I don’t want to be doing this it feels, but then something I learn about a specific api or thing and it ignites me to learn and grow, setting me into motion again.

### Quote of the Week

> Ramcharitmanas aeihi nama, sunata shravana paiye vishrama
> 
> “The very name of Ramcharitmanas(the manasa lake of Shri Rama’s journey) brings solace to those who hear its story.”
> 
> — Tulsidas, Ramcharitmanas, Balkand

Beautiful. This is just beautiful. I am not saying I have read the Ramcharitmanas, I am midst reading it almost consistently one paragraph atleast. It gives me joy that nobody can replace, its just a peaceful read.

I don’t have any other quotes, actually I have too many after reading some portion of the first chapter, but this quote just makes me smile. Hope it does too you as well

* * *

## Read

1. [Were software engineers blue collar workers all along](https://ptrbrynt.com/posts/were-we-blue-collar-workers-all-along)
   
   1. Very very deep thinking here. This gives me hope, if someone writes positive amongst the fear and hype, this is blessing, wonderful work by the author
      
      > But if we learn how to use the looms, the spreadsheets, and the chainsaws, we might just be okay. And the good news is that there are still people sewing and chopping wood by hand, because they want to. It may no longer be something you can get paid for, but it’s certainly something you can continue to enjoy.
   2. Yes, this is the piece that we as software developers need to think on.
2. [Software Centaur age may last decades](https://www.seangoedecke.com/softwares-centaur-age-may-last-decades/)
   
   1. Beautiful observations. I was thinking about the same on chess, like it is still being played competitively right?
3. [I am sorry but you still have to think](https://itsallaboutthebit.com/i-am-sorry-but-you-still-have-to-think/)
   
   1. Again, asking the right questions, thinking deeply about the problem is still a thing to do. I mean this is just the same problem from a different question right?
4. [AI is a tool, using it well takes skills](https://yyhh.org/blog/2026/10/ai-is-a-tool/)
   
   1. Skill Issues basically. I have skill issues while using LLMs. The post is true, without the context, without the right question, the LLM is a piece of shit. Just like excel is for non-computer users, pen to a illiterate person. You cannot get value if you don’t have the right questions.
5. [Now that information is free, what’s college campus for?](https://commonedge.org/now-that-information-is-free-whats-a-college-campus-for/)
   
   1. This is a beautiful writeup. I have watched talks from IISc professor discussing this with students about the same. If college was just to gain information, then it was already dead in times of google, LLM has just personified the information, but the bigger question is still not distilled in people. Its about connecting with humans, building communication skills, empathy and care in the world.
   2. I am not sure with the current pace of technology, education might be able to entertain students, the people need to have open discussions and debate about how to create meaningful time while being in college and what the students really need to learn.
6. [The Emperor Pelican has no clothes](https://arcfu.com/post/the-emperor-pelican-has-no-clothes/)
   
   1. This is true, they seemed to have stagnated on what they can produce, the creativeness is not novel, its just unpredictability is too noisy and not quite helpful.
   2. Artists, your jobs are safe.

## Watched

- Stop being skeptical about AI: Charity Majors in Pragmatic Engineer Podcast
  
  - Ok, we have moved from “being not sure” on how to write code with AI, to “when to ship” the code. The iteration has become so fast that producing plausible and working code is no more a bottleneck, how and what should be spitted out is.

<!--THE END-->

- [Talk to your agent like a human: By Anthropic take from Primeagen](https://youtu.be/ASY_H2nLG8g)
  
  - This is atrocious. AI cannot be called a “being” that needs to be respected with communication. I know its a fair call on catching fraudsters, but what is the point of presenting it like a “being” Dario?
  - The point of being impatient when a human doesn’t reply to you immediately is a bit relatable. ChatGPT / Claude have spoiled this trait of humans. The way of communications, or I should say the parameters of those are changing with these technologies, especially multi-modal, if they keep improving it will become really hard to communicate with AI and humans alike.

<!--THE END-->

- [AI is decompiling every game](https://youtu.be/s95Q7D0WR40): The Standup
  
  - For fun, this so so good, but for understanding and actual joy of doing it, its not a thing you should be bothered about. But Adobe and SaaS definitely needs to be now.

I know I don’t share non-tech stuff here, but just how peaceful this video is

## Learnt

- [Pydantic model before validator](https://pydantic.dev/docs/validation/2.9/concepts/validators/#before-after-wrap-and-plain-validators)
  
  - The before validator(@model\_validator(mode=”before”)) receives the raw input data before Pydantic validates the model fields, so you can transform a string or dictionary into the structure your model expects.
  - This is cool because I wanted to do something like this `"text___key:value::key2:value2"`
    
    and convert to a python data
    
    ```
    {
        "text": "text",
        "key": "value",
        "key2": "value2",
    }
    ```

## Tech News

- [Meta releases Muse Spark Open](https://research.meta.ai/blog/muse-spark-open)
  
  - Oh! This looks sick. Now, we are coming closer to reaching usable models on consumer hardware, this is still the best hardware you can get, not ordinary hardware of Apple, but still local “usable” models are coming.
- [Mistral releases timestral: Open weight model for timeseries](https://huggingface.co/collections/mistralai/timestral)
  
  - Hm, maybe we just tune a model to certain specific domain now?

<!--THE END-->

- [Google drops Gemini 4 Argon](https://deepmind.google/models/gemini/)
  
  - Meh! Like come on, what is wrong with Google, they are just good at understanding the context, but taking actions on them is just not the model’s expertise.
- [SQLite release 3.54](https://sqlite.org/releaselog/3_54_0.html).0
  
  - Subtle improvements like the CLI default header and title options are really handy and makes out of the box experience smooth and intuitive.

* * *

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/814/) (#814th edition), and for software development/coding articles, join daily.dev.

That’s it from the 114th Edition of techstructive weekly. I hope you found it helpful, and relaxing. If not please drop any suggestions, feedback or discussion about certain things you want to in the comments or drop me a message on my [socials](https://www.meetgor.com/contact/).

Thank you for reading,

Until next week. happy Navratri and Shubh Dusherra!!

Happy Coding :)
