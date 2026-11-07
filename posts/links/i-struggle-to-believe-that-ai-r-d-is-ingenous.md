---
title: "I struggle to believe that AI R&D is ingenous"
date: 2026-11-06
link: "https://www.lesswrong.com/posts/DALNiKBjMGMF8grme/arjun-panickssery-s-shortform?commentId=fJtaHHHfx45tD7BSL"
status: published
image_url: "https://res.cloudinary.com/lesswrong-2-0/image/upload/v1654295382/new_mississippi_river_fjdmww.jpg"
source: newsletter
newsletter: techstructive-weekly-118
type: links
slug: i-struggle-to-believe-that-ai-r-d-is-ingenous
tags:
description: "Comment by Arjun Panickssery - I struggle to believe at an intuitive level that autonomous R&D is so ingenious. I have been monitoring a few side projects:

 1. Building a new SOTA engine for Scrabble (current SOTA is not really a neural-network method and doesn't beat the strongest human player)
 2. Building a bot to play Ultimate Tic-Tac-Toe
 3. Researching philosophy and formalizing some proofs for the AI Philosophy contest

In all of these cases, the model's judgment is just very poor in a basic way. For example:

 1. When reviewing literature, it fails to identify the specific papers or blog posts that are obviously the most important and useful.
 2. It will conduct experiments that have no information value because the amount of training data used was obviously too small to distinguish the architectures to be compared.
 3. In the philosophy project, where it built a wiki over many tokens and investigations, it kept going down directions which were extremely trivial and had little value.
 4. For Ultimate Tic-Tac-Toe, I saw a very clever idea in a blog post that I doubt Astra would ever have come up with; it actually located the blog post but didn't realize its importance or relevance (which could be discerned easily from his strong performance in an online arena and the recency of the post) over papers from previous years that are obviously less relevant.

I judged an Apart hackathon last weekend and reviewed eight papers, all of which scored between 60% and 95% on Pangram. Predictably, the papers were often very similar to each other, and kind of methodologically incomplete in a vague but similar way, where they seem to overlook the most straightforward analysis.

----------------------------------------

Maybe the important exception is certain hill-climbing tasks with fast feedback loops. For example, I improved the SOTA on \"listing all legal Scrabble moves\" by 10x just by prompting Astra over and over to try harder, to aim for <specific performance level>, to \"systematically try every me"
hash: 034f854d04f3f010ea007cbeb5be9c42472434b7828724e43270045b979cf6b0
---
My thoughts on [I struggle to believe that AI R&D is ingenous](https://www.lesswrong.com/posts/DALNiKBjMGMF8grme/arjun-panickssery-s-shortform?commentId=fJtaHHHfx45tD7BSL): I struggle to believe that AI R&D is ingenous

## Commentary

- I struggle to believe that AI R&D is ingenous
- True. Models can once in a while do something ingenious but the problem is you are trying to solve is too small for it then. Always, LLMs need to be given a ambitious goal, but that fails catastrophically if not paired with right tools, or reward/review mechanisms I think. If it can’t verify, If you can’t verify if you’ve build something workable or intended then how can it be done?
- It almost always need that silver lining of human experience and domain expertise to be able to steer and create a wow moment for you to be effective working with LLM Agents right?
