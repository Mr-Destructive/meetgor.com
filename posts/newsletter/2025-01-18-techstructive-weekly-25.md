---
type: "newsletter"
title: "Techstructive Weekly #25"
date: 2025-01-18
---

## Week #25

It was an exciting week, I did complete a lot of things on my work side, I had a bit of change in the routine of my work, an hour early, which makes me really productive and gets a lot of things done quite early and more importantly, I am now able to dedicate myself more time to learning new things.

This week, I didn’t release any videos, I was a bit exhausted and wanted to slow down a little, but I was not sitting idle, I was evaluating a project idea for some future videos, and the project is ready. Next week, expect some high-quality and in-depth videos.

Also, this week I learnt a ton of things about AI Agents and Workflows. I even created a proof of concept for creating agents using Langraph. I might create tutorials or guides on this soon.

### Quote of the week

> “The capacity to learn is a gift,   
> The ability to learn is a skill,  
> The willingness to learn is a choice.”
> 
> — Brian Herbert

Observe, Understand, Adapt, and Learn are quite the pillars to grow and become a better developer. As a software engineer or developers, we have to adapt to new changes just like these, for working with LLMs, AI Agents, and all the cutting-edge things happening, if this doesn’t excite you, I don’t what will, something should at least to continue in the long run.

## Read

- [Golang 1.24 Interactive Guide](https://antonz.org/go-1-24/)This was really well explained. Every change was accompanied by an interactive example that made it really a breeze to understand the concept and change being added to the language. Highly recommend checking it out.
- [Hot Take: **TILs are junk food**](https://antonz.org/til/?ref=dailydev)
  I don’t agree with this, but some extent this post is true. I don’t read tils, I only come upon TILs from the feed that are relevant to me, so I only follow people whom I share the same interests and tools being used. There is value in reading someone's else perspective about using and discovering the tools or any thing common.

- [**How HTTP/2 Works and How to Enable It in Go**](https://victoriametrics.com/blog/go-http2/?ref=dailydev)
  This is also another post that I took the time to read and was worth it. I honestly don’t know how HTTP 2 works. To some extent, I know how HTTP 1 works, but if someone went a bit deeper, I would start breaking sweat. I really need to implement HTTP from scratch to understand the network stack—one day or day one.
    - OH, the article, yes it talked in detail about what is the problem with HTTP 1 and how HTTP somewhat solves it.
    - It is about breaking down the data into frames and makes sure the client has received the frame even if the previous frame is delayed. The fastest frame is served so, it doesn’t block the latest requests.

## Wrote

- [Golang Web: PATCH Method](https://www.meetgor.com/golang-web-patch-method/) (37th post in 100 days of Golang Series)  
  I wrote a 3k-word article, it was fun, and it was so rewarding. I learned a lot about different types of PATCH requests, the MERGE PATCH Request, the JSON PATCH Request, and the normal PATCH which is kind of misplaced everywhere.
    - Merge PATCH is quite the essence of the PATCH Request, we only send the field / updated field in the request and we expect the server will update those fields only.
    - JSON PATCH is the more complicated or I should say more fine tuned request. If the API or the resources are JSON based, this type of PATCH Reuqest is quite suitable and recommended.
- [Golang Web: DELETE Method](https://www.meetgor.com/golang-web-delete-method/) (38th post in 100 days of Golang Series)
  This felt quite short, after writing the mammoth article on PATCH. It was simple yet I wanted to complete the whole HTTP Method Suite. Now the next article will be diving deeper into the HTTP Clients and creating APIs.
- A simple CRUD Application with Golang and Netlify Cloud Functions with Turso’s LibSQL Database

## Watched

- Picking a language in 2025
    - I have still not completed watching this video, but till this point, all the things mentioned are so relatable.
          - C++ Standaard library: I learned the very basics of C++ in my first few years of the self-taught developer route, however, I never found it intuitive to use the standard library. Now I realize the beauty of Golang and Python. Just simple yet powerful to do almost anything.
          - Learning a language is quite simple, but mastering and getting proficient at it takes years. I can relate to Golang and Python, as I am still struggling to date.

Double click to interact with video
- [AI Risjs No one is Talking About](https://youtu.be/pmtuMJDjh5A?si=l_XGGNMSMp1VOoEZ)

> Seeking the truth, the thing that AIs won’t ever have, the thing that makes human a human  
> — Teej Devries 2025

This was such a great short talk that covered a lot of questions and the right things to consider while adapting to LLMs. They are good, but still far from taking over, why on earth people want LLMs to take over int he first place?

Double click to interact with video
- The worst thing you can do for your career is to quit now:  
  Yes, that’s the worst part
  But the best thing you can do is to just give your best, not even the best, be a little curious, have the grit to seek the truth, (troll devin sometimes)  
  And expect it to work out, if that doesn’t, then the last choice would be to quit eventually, but not without a fight, not just yet.Double click to interact with video

## Learnt

- Creating AI Agents with Langraph
    - I have been learning about creating Agentic workflows. This week, I spent much time understanding the various components, such as the router, orchestrator, tools, and conditional edges.
    - I want to summarise what I learned in a few sentences:
          - The agent is someone who takes action on certain inputs by acting on the environment. In the case of AI Agents, the LLMs will be given some text (any data) and they will be provided some environment to act on.
          - Tools are the sets of actions that the Agent can take this case, LLMs can interact with data by ng APIs, custom code, another call, etc.
          - Router / Orchestrator / Oracle is the decision maker, the central unit deciding what to do when.
    - I might create a series around this sometime soon, I am currently overwhelmed by learning so many things and need to absorb them by writing more.
- Using diff for two strings in Linux
  ```
  diff <(<<<"123") <(<<<"1234")

  Output:

  1c1
  < 123
  ---
  > 1234
  ```
- Python Count number of occurrences of a particular element in a list
  Some subtle thing but really a handy way to count the occurrence of the element in a list.
  ```
  >> s = [True, False, False, True, False]
  >> s.count(True)
  # 2
  >> s.count(False)
  # 3

  >> nums = [1,1,2,3,4,4,4,5,5,5,5]
  >> nums.count(2)
  # 1
  >> nums.count(4)
  # 3
  ```

## Tech News

- [AI will start automating the work of midlevel soydevs according to Mark Zuckerberg](https://www.reddit.com/r/Futurology/comments/1hzvcbd/mark_zuckerberg_said_meta_will_start_automating/?rdt=64385)
    - Come on, really Devin will replace me? I might have fun with that though, good luck to the employers fixing and debugging bugs in the prod!
- [Microsoft trying to form an AI focused on developers](https://techcrunch.com/2025/01/13/microsoft-forms-new-internal-dev-focused-ai-org/)
    - This is unexpected, Microsoft might have very rarely cared about developers. Is this in a bad way? Like Copilot for developers, or a snatch away the code with spies?

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-728) and for software development/coding articles, join [daily.dev](http://daily.dev/) .

That’s it from this 25th edition of my weekly learning, I hope you enjoyed it, and leave comments on what you think about some of my takes or any feedback

[Leave a comment](%%half_magic_comments_url%%)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](%%share_url%%)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
