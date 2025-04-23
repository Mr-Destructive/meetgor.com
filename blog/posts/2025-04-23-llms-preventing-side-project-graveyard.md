---
templateKey: blog-post
title: "Can Developers prevent the Side-Project Graveyard with LLMs"
description: "Could LLMs help us(developers) avoid setup and analysis paralysis problem and side project graveyard from piling up."
date: 2025-04-23 23:15:00
status: published
slug: llms-prevent-side-project-graveyard
tags: ['llms','thoughts',]
---


> Side projects no more in the Graveyard!

I have lost count of many times I have heard the word "Agentic" in 2025! Especially with the code editors and IDEs integrating AI natively. It might be a revolution how developers code.

It's quite terrifying to see the revolution and adoption in the community so quickly. There is hardly any week without new tools, models, approaches to how developers use and figure out the quirks and ways to leverage AI/LLMs into their workflow.

One thing that came to my mind is the thing of the side project graveyard. I have around 50 actual projects in the bin! Is that just me or it's a real thing, I think its common to have a personal side project graveyard. That's when you verify you are a true 100x developer right? and not a grug brained developer.

## LLMs with Abilities

The point I will be making in this post is "LLMs (Agentic IDEs) might help reduce that side project graveyard"
Why?
At this point, LLMs can
- Know generic stuff (knowledge)
- Reason
- Take Actions (Tool Calling)

That last bit there, has put LLMs from kind of useful state to actually human-like abilities to code (paired with reasoning).

I am not saying it is going to replace developers. Please keep that thought aside for a while. Like I do all the time, for now.

Just think of LLM as your assistant, you can instruct it to do a few things so that it takes the action and you can focus on the interesting bit.

Because lets be honest, till the point we try to setup and reach the interesting (actual thought provoking part) we already have exhausted a lot of energy. The term cognitive load hits here. We had too much back and forth between configs and stack overflow (even LLM conversations), too many commands written with hit and trial method, fixing dependencies, conflicts, updating package managers and what not. Now I am not saying these are bad or not needed, those are valuable things that truly can make a good developer a better one, through learning with the mistakes and debugging things. But at what cost?

Time and energy, at some point, we loose the motivation for the actual thing that we were excited for.
We might have the curious energy but physically or mentally we might have been out of steam.

Enter Agentic IDEs (Copilot, Cursor, Windsurf, you name it)

## Agentic IDEs

Those are really good! I had a suspicion for a while, but one weekend I gave it a shot and I was blown away!

For those who are not aware, Agentic IDEs are development environments where LLMs can interact not only with the user but with the code, terminal, and the system in general. Provided with the search tools, file system and its general knowledge, LLMs are like tiny bots that take action on behalf of you.

What Agentic IDEs can do
- Search and Read File context
- Run terminal commands
- View the output of terminal and decide what to do next

Some other obvious ones
- Tab auto-completions
- Inline chat
- Normal chat interface

One shot in a simple project (Link Reader and Aggregator). The project was simple, basic authentication, user can dump a bunch of links inside of a group (like collection). The Interface will be reading focused, so the user can then come back and read each of the content in the links in a simplified reader interface. Without any ads or distractions.
It handled the CRUD pretty well, the UI was not great or too fancy, and that was expected since I had never prompted it to make the best of UIs. But I wanted to test the capability of LLMs to handle CRUD and authentication. Those it handled fairly well, no doubt those are crucial and basic things to add, but it helps in getting it out of the way.

## One Shot experience

I attempted this on Windsurf with `GPT 4.1` since it was on free mode for the past week, and I think they have extended it to this week too. It might be the model that is good at coding, but given how good it was, other models might not be far sooner or later.

Now with that project, I have something to sketch on, something to fix, something to take ahead. Instead of staring at a blank page (editor), I already have a working or broken thing in front of me. Making broken things work is more easier than to making the thing, and making the thing to work.

I can let LLM debug the errors, so that I can write the core (maybe complicated) logic myself. Instead of me debugging and letting LLM write that complicated logic.

## Conclusion

At the rate at which things are moving, it seems a bit naive to comment on the scope of LLMs in improving the developer experience. However, so far it looks a massive improvements in the velocity, the friction removal, and even making things offload from the boilerplate and cookie-cutter things.

## A thought for Juniors

Though I also feel for juniors (I am one too) and students starting their developer journey now. I have developed a intuition about development to some extent, however for them its completely new, and intuition can't be built without experience, one has to try, fail and learn, no shortcut there. But again it's moving too fast for everyone of us that adopting LLMs into the workflow is tricky and awkward at times.

But software development is a field that has stood the times of change. Change is the only constant in software development. Adoption of LLMs is no brainer, we have to leverage it.

Fellow developers, we have a chance to reduce or keep the size of the side project graveyard the same! Let's leverage LLMs to the fullest and ship projects.

Happy Coding :)
Happy Vibing :|
Happy Debugging :( 
