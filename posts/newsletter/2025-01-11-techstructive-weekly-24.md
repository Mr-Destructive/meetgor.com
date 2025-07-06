---
type: "newsletter"
title: "Techstructive Weekly #24"
date: 2025-01-11
---

## Week #24

It was an exciting week. I was pretty excited from the end of the past week with a new beginning to my work-life goal. After completing one, I have a new challenge in front of me: diving into the world of cutting-edge LLMs, agents, and Workflow. This can’t be more exciting.

This week, I was researching and exploring the frameworks, libraries and tools to create Agentic Workflows with LLMs, I can’t share much, but I see myself creating content around the things I will learn at my work/day job.

Apart from the hustle, I [streamed again](https://www.youtube.com/live/hr3Xsuw0IDk?si=yjNVYVDyCkX5gBnl) keeping the 2-month streak, moreover, I am excited and have new ideas popping out frequently. I have also been writing a script for a project-based tutorial video, which was the stuff I wanted to create when I was learning to code myself. I will try to complete the video this weekend, it would be a long one, but I would be delighted to create it.

Double click to interact with video

### Quote of the week

> *“The future is not something we wait for, it's something we create." *
> 
> — Pierre Teilhard de Chardin

I can’t wait for next week to roll in through the implementation, some proof of concepts, some demonstrations, and some planning, this is the future I am creating. I take the ownership and I build it, and not wait, sit and watch it unfold. I create my future, I am responsible for the misery or the pride I will have at the end.

## Read

- [Simon Wilson Predictions for AI/LLMS for 1/3/6 years](https://simonwillison.net/2025/Jan/10/ai-predictions/): That sounds a bit scary to be honest, but exciting as well. It would be more of the adoption rather than autonomy of the LLMs that in the coming years will dictate the direction.
- [Living the future by the numbers](https://tailscale.com/blog/living-in-the-future?ref=dailydev): It sounds like enormous growth, the cost is a different thing as mentioned in the post, but yes, the performance is the order of magnitude higher which makes a point for the progress humans have made in computing.
- [Human writing in the age of AI](https://alvaromontoro.com/blog/68068/human-writing-in-the-age-of-ai): This is a truthful article, we are growing ourselves in the photocopies of photocopies (using ai to summarise content written by ai). At some point, there will come a time, when humans will need human content and that’s when there will be a search for a gold rush, the authentic content rush. It’s time now, to make hay while the sun shines, it might be the last time when humans are valued before they are not and again they are. Be in the dark standing to help others when the darkness seems to take over.
- [Building Effective Agents (Anthropic Blog)](https://www.anthropic.com/research/building-effective-agents): This is a great overview of LLM Agents, different types of Agents and how is the landscape of AI Agents is changing due to LLM. It’s not changing technically, but how simple concepts from AI are striking again like a revival. I remember learning about it in my bachelor’s degree classes, agents as some program that acts according to the conditions and interaction with the environment, I can relate it so well and sticks to me due to that.
- [FastAPI Documentation: Concurrency and Async/Await](https://fastapi.tiangolo.com/async/): This was clear and the example provided are indeed helpful and relatable. I read this to understand the architecture these frameworks for creating agentic workflows are using. I had a confusion about async and asyncio, so just read this and got absolutely cleared.

## Wrote

- [Code for dummy json patch API](https://github.com/Mr-Destructive/dummy-json-patch): This is a dummy, a demonstration-only API, will add a documentation page and more to it as I write the Golang HTTP PATCH method article. Hosted on netlify, took some time to understand the patch request and constructed the json-patch endpoint, will also add in a json merge patch endpoint.

## Watched

- [Dennis Ivy: Building with Django and HTMX](https://www.youtube.com/watch?v=rgegMr8ImKo&t=2607s)  
  I have not touched Django and HTMX as well for quite some months (maybe a year). I need to get back and create some quick projects for the community. I want to keep the guides and tutorials on Django from my article series more versatile and expansive. Talking about the stream, it was a great stream, loved that Dennis is back, learned a lot of stuff with htmx. A lot to learn when done yourself.

Double click to interact with video
- [Distributed and Async task Queues with Python and celery](https://youtu.be/v-Snbz3WmJU?si=qyiFzrQQrZ5MrrjJ) Pycon US 2024:  
  This was an in-depth guide to celery, I learned a lot of stuff, celery might confuse you if you have not paid close attention to the details while implementing the queues. I forgot what celery was doing internally, I had used it, and built background workers while I was an intern @ [Arka 360 (The Solar Labs](https://arka360.com/in)). I learned a ton of celery and Django, but it was 2 years back, forgot a few things, and after watching this it clicked everything and now makes sense.Double click to interact with video

## Learnt

- I learned a lot of things about the JSON PATCH method, I will be writing a detailed post about using JSON Patch in the article **100 days of Golang: HTTP PATCH Method, **which should be live after this newsletter.
    - The HTTP PATCH method is like a PUT request but for updating only specific fields and not the entire resource in that sense, you only send the fields to be updated compared to the PUT request where you have to send the entire resource (including the ones that you don’t want to update).
    - [JSON PATCH](https://datatracker.ietf.org/doc/html/rfc6902) is a specific type of PATCH method, where the payload is a JSON patch document, which includes
          - The operation to be performed
          - The field (path) to be updated
          - The value of the field/path to update to
    - [JSON Merge PATCH](https://datatracker.ietf.org/doc/html/rfc7386) is another type of PATCH method where the payload is a JSON body with the fields and values you want to update (like put but only the fields to update)
- [Replace github with gitingest](https://thoughts.waylonwalker.com/post/517) for getting a single string value for all the file contents, useful for LLM prompting.
- [Slash pages for specific sections on the blog like](https://thoughts.waylonwalker.com/post/494)
    - /now
    - /colophon
    - /links, etc

I am about to create a thoughts/link blog page for my site. Moving my site out of an SSG is the new year norm I believe, I moved my site from Jekyll to Markata 2 years ago, and now it seems the correct time. I don’t have any problems with Markata, I just want to use Golang for all the things I make, it feels good (not biased).

## Tech News

- [Appwrite Compute Capabilities for Appwrite Functions](https://appwrite.io/blog/post/introducing-new-compute-capabilities-appwrite-functions): Appwrite keeps pushing the bar for what is called an open source backend as a service.

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-728) and for software development/coding articles, join [daily.dev](http://daily.dev/) .

That’s it from this 24th edition of my weekly learning, hope you enjoyed it, and leave comments on what you think about some of my takes or any feedback.

[Leave a comment](%%half_magic_comments_url%%)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](%%share_url%%)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
