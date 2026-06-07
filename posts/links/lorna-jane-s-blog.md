---
tags: 
description: "Lorna Jane’s Blog:"
date: 2026-06-05
link: "https://lornajane.net/"
status: published
type: links
slug: lorna-jane-s-blog
hash: 31afb3a40f51d3bc95279b67a277648023f10facdd94bfd4314c62c1c383cb65
title: "Lorna Jane’s Blog:"
image_url: "https://beta.lornajane.net/wp-content/uploads/2023/11/LornaJane-OG1.png"
source: newsletter
newsletter: techstructive-weekly-97
---
My thoughts on [Lorna Jane’s Blog:](https://lornajane.net/): Lorna Jane’s Blog:

## Commentary

- Lorna Jane’s Blog: This blog is a gold mine for API lovers. It also covers a lot of developer relations and dev-ex. Really a perfect combination for something I want to read and talk about.
- Overlays in OpenAPI and avoiding oversharing APIsThis is a gem. We can trim off or add certain things to the spec in certain points of the pipeline of the API Spec. This is so powerful for removing internal APIs and only exposing user facing ones, while still having the rich documentation for humans (now AI agents) to understand the actual api underneath.Being able to do this programatically is amazing and makes a API pipeline so controllable and rich at the right place and time.
- Using Tags in OpenAPI SpecI was confused as to what the use of Tags was actually, might be for listing certain groups of related api. I didn’t think of grouping in the documentation page and elsewhere. I was just puzzled as to where it could probably be helpful. This article just made it clear.“Quite a few of the documentation tools will render the endpoints grouped by tag which is very valuable”. That is what I wanted to hear. Nobody made it quite obvious.Now I see the header like sections, I can now think, those are the tags, yes, so the APIs inside those tags might be related.
- What’s new in OpenAPI 3.1 SpecThe type having a array of possibilities is so cool. We can basically have a union like datatype in python here. Powerful.The webhook thing, i must say it always confuses and frightens me. I always associate webhook to clients, I get scared. But this article clears that tooSuppose we are writing the Stripe’s API Spec. The customer of stripe, some fancy startup needs a event to send to the user when it happens, so the URL is defined by the user. So we don’t add the URL, we specify the event name (by convention not necessarily enforced) and the method. If there is a method then there must be request and response to provide those.There are other changes as well like the override in ref and info block has a summary field now. By the way, there is a 3.2.0 released as well https://github.com/OAI/OpenAPI-Specification/releases/tag/3.2.0 by the time of this post. It feels good that there is still innovation and releases happening at for granted taken things like APIs.
- Tips for better documentation with OpenAPISummary and description are really undervalued. I cannot love a API if it doesn’t have relatable and to the point humanly descriptions and summary, even a simple one liner helps a ton.Does than will do is really a good thing to keep in mind to build trust