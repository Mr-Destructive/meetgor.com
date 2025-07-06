---
type: "newsletter"
title: "Techstructive Weekly #33"
date: 2025-03-22
---

## Week #33

Finally a bit of stress reliving week. After constant hustle at work, this week felt a little less stressful (not stress as that stress, just too much cognitive load to handle). However I made 2 releases happen, a few experiments and brainstorming for cracking a problem.

### Work 

This week made me realise, I sometimes do things. I was asked to figure out a metric that we need to provide to a customer based on their customers data, they basically need to understand the data of the customer, however the data is too much and they forward to us to extract and analyse on top of that. So, my job was to figure out the requirement of them, they need to understand the pattern in the data of the customer and that is not quite easy.

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.

I did some analysing of our competitor data and it looked skewed, asked to ChatGPT, I cannot stress the use of interpreter mode enough. That just gave the LLM the superpower to understand its own mistake instead of us figuring out, and that saves a ton of back and forth. I passed the data and context to LLM and it gave somethings which really opened a eye and was then able to take further steps wit the guidance of my manager.

That was a rewarding thing, I don’t know what, that feels like a new way of solving problems. I feel excited when that quick prompt is able to guide me in the direction of the actual problem than sifting through the data.

### Quote of the week

> “You can’t solve a problem with the same thinking that created it.”
> 
> — Albert Einstein

Think out loud, but don’t let that hold you back in the actual work. We need a approach, a mindset fresh from the problem, if you created a problem, its more likely you have the same mindset to solve it, however that won’t get you out, you need to take a step back and understand the cause.

## Read

- [Yes you are screwed](https://ghuntley.com/screwed/): That was a eye-opening article. In the age of LLM, as a junior developer, we need to act, respond with curiosity rather than doubt or ignorance.
- [Porting Typescript compiler to Golang](https://devblogs.microsoft.com/typescript/typescript-native-port/?utm_source=hackernewsletter&utm_medium=email&utm_term=fav): That is givign them 10x boost in time, that is insane for compilers, I mean time doesn’t really matter in compilers but for developer productivity that helps a lot.
- [Building Websites with LLMs](https://blog.jim-nielsen.com/2025/lots-of-little-html-pages): This was a really a nice post, need to add this to my blog as well. I want to write more and the more I write I will create more sections like TILs, Thoughts, Link Blog, Articles, Series, Tags, and we can go on, I want to keep my existing site simple yet, I want to organise them in a susinct way.

## Wrote

- [Notes on Gemma 3](https://meetgor.substack.com/p/notes-on-gemma-3?utm_source=substack&utm_content=feed%3Arecommended%3Acopy_link)
  Inspired from [Simon wilson](https://simonwillison.net/2025/Mar/12/gemma-3/), I also did some experiments on the new model, on my phone as well as on laptop. Running a LLM on phone feels like magic.[![](https://substackcdn.com/image/fetch/$s_!Zyre!,w_56,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F9786b965-3564-4542-8687-a55727236ab5_325x325.png)Meet's SubstackNotes on Gemma 3Google announced a new family of models, which are open-sourced weights and have 4 variations of the parameters, 1, 4, 12, and 27 billion parameters. Leaving the 1 billion parameter, the rest 3 parameters are multi-modal, they support images and short videos…Read more4 months ago · Meet](https://meetgor.substack.com/p/notes-on-gemma-3?utm_source=substack&utm_campaign=post_embed&utm_medium=web)

I also streamed on weekend, one stream for wrapping the static site generator with cron job on github, and the next stream for contributing to posting with file uploads.

## Watched

- [Thinking/Planning vs Building](https://youtu.be/rosMfs3pZ_0?si=z7NM3q_xRIHYe0Ek): This was a cool video, I was always with hands-on approach when it comes to developing anything. However since the last year, having the company of my manager, I got into the habit of thinking and sketching out the plan for the implementation, maybe that is different, we don’t create a document, we just brainstorm on the idea. Making the logic clear.  
  I also like the direct code implementation, to get a taste of what works and what is the actual problem to solve than to guess and later realise the problem was to pick stones than to carve the stone.Double click to interact with video
- LLM Chat Apps are driving Theo insane: Hire Theo please, OpenAI, Claude, Meta. the design is really falwed, too unintutive and buggy at time. They are multi billion dollar company, they can’t afford this mistakes.Double click to interact with video
- [Overview of Marimo](https://youtu.be/3N6lInzq5MI?si=kEqeaHpAiQaadVFN):
  Marimo is a powerful tool to have in these times. LLMs are chained and this will help build better workflows, rather help in experimentation and iteration. Having to work with graph related scripts, with graph designed platform is the best way to move ahead.Double click to interact with video

## Learnt

- [Using marimo](https://marimo.io/) to create,edit and work with python based notebooks. Marimo is a modern alternative to jupyter notebooks, you can port a ipynb file to marimo, as marimo stores directly as a python file. That is really neat. Marimo’s core feature, or I should say selling point is the reactive system. If you change a value of one variable in a cell, then all the cells referring that variable will be mutated, or changed. That is actually cool, and makes notebooks intuitive to use. I’ll be exploring more of those these weekend.
- Debugging a textual app with the textual dev console: I was livestreaming the implementation of file upload in posting app and needed to debug the process, and the following command actually made it super easy to do that:
  ```
  > textual console 

  > textual run --dev posting.__main__:cli
  ```
- [HTTPX request structure](https://stackoverflow.com/questions/72103585/how-to-pass-file-object-to-httpx-request-in-fastapi-endpoint) for uploading files in the request. As I was trying to figure out the fix for uploading file in posting app, httpx is the library used under-the-hood to send and save request. Hence I felt the need to look up how to parse files into the httpx request object.
- [Curl uses `@` for indicating a file upload](https://stackoverflow.com/questions/12667797/using-curl-to-upload-post-data-with-files) in a request form data. That is like a decision they have made, a choice, that could be something else as well. However the community is supportive and adheres to the convention that is followed in all the places, rather the systems are designed around it.

## Tech News

- [Open AI Releases Tools](https://openai.com/index/new-tools-for-building-agents): They added the response API, which is suitable for Tool based calls. Web and File Search, as well as Computer use for automating stuff. This is pretty significant thing to be honest, in the direction of agentic revolution. These tools and ease of adoption will boost the landscape of this agentic hype.
- [Google releases Gemma 3:](https://blog.google/technology/developers/gemma-3/) Google released gemma 3 with 4 variation models of 1b, 4b, 12b, and 27b parameters. 4,12 and 27 are multi-modal supporting image parsing, have 128k context window, also support tool calling. I wrote some [notes on this](https://meetgor.substack.com/p/notes-on-gemma-3).

Phew, I did a good comeback after 2 weeks of complete slump. That was it for a bit busy week in tech. It has been quite a lot, models released as usual, transformative adoption of tools for agentic workflows.

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
