---



























title: Techstructive Weekly 69
type: newsletter
date: 2025-11-22
status: published
tags:
- ai
- devops
- git
- newsletter
- security
source_url: 'https://techstructively.substack.com/p/techstructive-weekly-69'


























---

## Week #69
It was a good week. That’s what I can say. I continued writing. I experimented with quite a lot of things. VLLMs, new models, new approaches, tactics, and read a lot of articles as usual.

As I was in the GCP Cloud Log Explorer TUI in the bubbletea golang app, I vibe-coded it, I wanted to feel it. No, I am not using that to publish on GitHub. I would make one from bare hands, what to put in the model, on the screen, when to update what, which keybindings to add where, I will think about each of them and code from scratch. My goal is to learn and build, not just build.

I wrote for the whole week, 21 days in the streak, and feeling good and motivated to complete the month on a high note, the project isn’t complete, but it put it on the right stage for me to complete it.

### Quote of the week
**“You can always edit a bad page. You can’t edit a blank page.”**— *[Jodi Picoul](https://www.goodreads.com/quotes/568141-you-can-always-edit-a-bad-page-you-can-t-edit)*I’ll go with this one. I have written for the past 20 days, and I can surely say I am a bad first drafter. I have no story, just an idea of a scene or a plot for an entire story. I am lost in the start. I still put pen to paper, wrote ~1000 words daily, and here I have some things to move the story, from point A to B.

Another quote might be

“As you start to walk the way appears”

— [Rumi](https://www.goodreads.com/quotes/811906-as-you-start-to-walk-on-the-way-the-way)I had one idea, but the start was blank. I still wrote some garbage. It was enough to help me break the barrier and helped in thinking out. After writing so much shit, not slop, I have a good understanding of what works and what the issues are. As the first quote goes, you can’t edit an empty page. There is no story to edit, there are no flaws to fix, no plotholes to patch, what is that for? You have to fail, understand the missing points, connect the dots, and move forward.

## Read
[A project is not a bundle ot tasks](https://secondthoughts.ai/p/a-project-is-not-a-bundle-of-tasks)- This is one of the best and practically grounded takes on the progress on AI. If it can’t make a decision that lasts longer, it might collapse eventually. AI Coding has improved over the last year for sure, but is it there? Not really. A developer still has to understand the logic, think about the AI answers, there has to be a human intuition and judgement to make its way to a sustainable solution.

- As per the growth, developers might be overgrown by 2030, is that really that quick? Maybe the plateau is almost here. It can code, sure, but it doesn’t create anything. It can create something like they are created by humans, not something out of the blue.

[Alien Authors](https://behan.substack.com/p/alien-authors)- Wow, this is just wow. The analogy of writers managing a bunch of written text from AI to create their own story is quite relatable to a developer managing a bunch of parallel agents to generate code and design a feature or come up with a solution.

- The whole story was really cleverly put and relatable. Though I would argue it’s quite a bad one for making one artist happy. I never want an artist to generate art with AI and tweak and tweak to please someone. Art is something that comes fully from the heart, otherwise, it’s not art. It’s pseudo art, artificial art, forced art.

[Why Software Development fell to AI first](https://davegriffith.substack.com/p/why-software-development-fell-to)It makes sense now,

- It is instantly verifiable&gt; if you can run it, it has some potential

- The medium is the message &gt; Text in &gt; text out

- There is verification, and it can check itself &gt; if there are tests, it can run, it can itself verify if the generated code works or not, and tweak accordingly

- These are compelling enough reasons to understand that software is quite suited for AI to bite its jaws from.

[The 18th November 2025 Cloudflare Outage Report](https://blog.cloudflare.com/18-november-2025-outage/)- Ok, that was quite a big outage. It was the biggest Cloudflare outage since 2019. Four hours. It was down, everything was down. It got a few eyes on the engineering debt.

- I wonder if Cloudflare was really required in their tech stack who were affected in this outage. Really? I don’t think so, just for reverse-proxy, verification, 10 users, do you need Cloudflare? No, dog, you don’t.

This is what is rightly pointed out [here](https://gist.github.com/jbreckmckye/32587f2907e473dd06d68b0362fb0048). It was a good thing, it put eyes on things which we take for granted or overlook, and just use the defaults, the frictionless tools.[How bcrypt can be unsafe for more than 72 characters](https://blog.enamya.me/posts/bcrypt-limitation)- Oh, that is wired, use Argon guys, if you aren’t just storing passwords.

- Nice to know that bcrypt is not safe for passwords greater than 72 characters, who would even store such a long password?

- But that is the thing, subtle decisions, like this is not a password, so we can use bcrypt, and bam, you would be wrong

[Fizzbuzz without conditions or booleans](https://evanhahn.com/fizz-buzz-without-conditionals-or-booleans/)- A nice pattern like loop for 3 and 5 divisibility, it won’t scale i think for other problems

[Make your own website](https://michaelenger.com/blog/make-your-own-website/)- Yes, it helps you understand what you need, what you are actually writing. I have built an SSG and learnt a lot of things.

- It puts you under control, it helps you think broader, and not rely on third-party things all the time.

[I can’t recommend Grafana to everyone](https://henrikgerdes.me/blog/2025-11-grafana-mess/)- Change is the fundamental in software, if something out there exist, it might not tomorrow, there is no gurantee, even if someone says so, might not.

- There are rare gems like SQLite, cURL, linux kernel and some of the fundamental tools that just work and don’t change or won’t change because that is what they do are supposed to do. No more no less, more they are adding, but it might come at the cost of backwards compatibility.

## Watched
[The probelm with AI Slop](https://youtu.be/vrTrOCQZoQE)- I am also quite happy, if we use LLMs to train on their own generated data, it will be a stangnation. Artists are going to thrive here, but that’s too dumb of a mistake these AI companies are to make. They can take all possible measures to make the people make use these chatbots more and more.

[Gemini 3 is the best model ever made?](https://youtu.be/39PdgOYjBMg)- Its a good model, it seems its quite heavy and removes the subtle mistakes and biases it has. I love hove Google makes a solid general purpose models. Unlike OpenAI, whose naming conventions are all over the place, its like a slop generation to me.

- But Gemini models hit different, they just solve what have been given to them. Quite a good upgrade from the 2.5 models.

- Waiting for Gemini 3 flash version

[TOON vs JSON](https://youtu.be/nTMP_rLZOYM)- Its just CSV in disguise. YAML in some other way, nothing really surprising.

Building SSH TUI with Wish and Bubble Tea

- It was quite magical to see the TUI on a SSH, imagine you can just use a lot of applications with a nice interface from anywhere. That is a superpower.

- Charm has a really good intention and they have done a great job in solving the problems.

## Learnt
GCP Log Explorer Admin API Routes

- Understand the Admin API, the request-response structure. The model design for which parameters to keep and are relevant for the TUI state.

- The log item components, it has the timestamp, the log level, the text, and the labels. It helps us making the design for the TUI as these are the lego blocks for it.

Construct a BubleTea Application

- The Model, View, and Update architecture is so ideal. It just fits well. So intuitive to use and feel.

Parse TOON-like structure to JSON

- Its no brainer that toon is a mix of YAML and CSV. KV fields are just YAML and List elements are CSV. Combine them, and here is your TOON.

- For input it’s an easy thing, but for making it generate the TOON Structure is no easy thing.

## Tech News
Google Releases [Gemini 3,](https://blog.google/products/gemini/gemini-3/) [Antigravity](https://antigravity.google/), and [Nano Banana](https://blog.google/technology/ai/nano-banana-pro/)- This is quite a hell of a release week. New Model, an AI Coding assistant IDE, and an image generation/editing tool.

[OpenAI Releases GPT 5.1](https://openai.com/index/gpt-5-1/)- Is there a day when these companies held themselves back from releasing their own slight better models just to get a few eyeballs.

- No, they don’t. They just game the numbers of the evals and make everyone believe that they are the king and worth using and making them the default.

- I am sick of these tactics from both OpenAI and Anthropic.

[Meta Releases SAM - Segment Anything Model 3rd version](https://ai.meta.com/sam3/)

This week and the next week or so are quite busy, I’ll be travelling out from home. So might miss out on next week’s edition. Its wedding season, and I have no power to refuse the invitation, forcefully or willingly, I’ll have to do my duty as a family member. Anyway, will catch up some other week for a fresh or exhausted fortnight.

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/772/) (#772nd edition), and for software development/coding articles, join [daily.dev](http://daily.dev/).
