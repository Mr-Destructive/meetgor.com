---
hash: 2a11091b3cc426e432c518062816909734dabe063dd2662d86011ba161f32c2f
date: 2026-04-17
slug: techstructive-weekly-90
status: published
source: newsletter
canonical_url: https://techstructively.substack.com/p/techstructive-weekly-90
description: "Learning Java, yes? Coping with AI against hand coded mastery, keeping code review skills in check among the other interesting tech-things read, watched,learnt in the week from 12th to 18th April 2026"
tags: ["newsletter", "substack"]
title: "Techstructive Weekly #90"
type: newsletter
---
## Week #90

Another rough slog of a week. But fighting nonetheless. Hopes are the glasses we have strapped in, darkness cannot blind us.

I have been upskilling myself with a few things, locked in with handcrafted coding. I am spending at least 2 hours in the morning to solve problems without chatgippty in my way. I don’t use agents and no AI. I think there needs to be a restraint in building and improving code skills. So, planning to do weekly livestreams for that public building and tracking my progress, also it helps me be accountable of what I say and commit to.

This week I am learning Spring Boot and Java skills, I don’t know why, don’t ask me yet. But I don’t learn anything without reasons, it might be curiosity, it might be a requirement. Learning is always a skill of a developer, and each year I want to learn something different. Maybe 2026 would be a year I would learn Java for enterprise-grade? Wow, that will be a great flex to make while people are prumpting with their AI-buddies.

### Quote of the week

> “What I cannot create, I do not understand.”
> 
> — [Richard Feynman](https://www.goodreads.com/quotes/8414-what-i-cannot-create-i-do-not-understand)

This is my mind right now. If I cannot hand code it, I don’t understand it. Your counter point might be, no one cares if you are anyone understands it or not, its code, not a form of art that is valued. Yes, that is true, users don’t care a damm if I hand coded it or the clankers generated it, but when it breaks, who is responsible. Yes, say it. It is a developer. No one can blame the clanker, people exist to blame people. I am not saying to blame it or not, but to take action and a decision from there on to resolve the issue, is only possible if I know the ins-and-out of the program, that only happens if I am in the weeds of it. Reiewing code can do it, you think, sure. Do it, and don’t complain after 3 hours of glazzing on the clankers’ 3k line of code slop.

I need to create in order to understand, I need to write it to know what it feels like to break and make. Before reviewing the code, I need to know how it works and what can break.

* * *

## Read

1. [Frustration Driven Development](https://swizec.com/blog/frustration-driven-development/)
   
   - Wow! No engineering is good engineering. Remove the problem from the root, well said.
     
     > Your job is not doing the work, your job is removing work.
   - How hard does that hit? In the AI world? Automating our own job, yeah, that’s fun, right? Maybe!
   - But this is a great point, great software engineers use swear words and get the problem out of the way, and do not work around the problem. (as codex does, and Claude code just wraps around the problem when it sees it)
2. [Who will be the senior engineer in 2035?](https://theengineeringmanager.substack.com/p/who-will-be-the-senior-engineers)
   
   - Totally valid point. If AI is becoming the cheap junior dev, who are we upskilling for the future?
   - The gap is already wide. In a couple of years, it might grow like crazy, and the divide of expectation and reality might hit people and developers alike.
   - There is and will be an expectation of in-depth understanding of code, but by using the AI-Agents, the quality will deteriorate drastically; developers won’t read and write code, and the instincts and the muscle memory to write code without assistance will be gone.
   - I am thinking of doing weekly streams now, to hand-code certain things. Maybe that will keep me up my ante of vim flexing and away from vibe coding.
3. [Lorin Hochstein’s thoughts on BlueSky public Incident writeup](https://surfingcomplexity.blog/2026/04/12/thoughts-on-the-bluesky-public-incident-write-up/)
   
   - That is a brutal one. This is like concurrency, you thought you had one problem, but after adding concurrency, you now have two. Setting a limit on how many go-routines can be spawned in a group is critical here.
   - TIME\_WAIT, that is really a neat thing to learn. The TCP connection waits for that duration before sending another FIN so that the client can be sure of the delayed packet delivery, if any. And that TIME\_WAIT actually caused them to fill up all the ephemeral ports. That is 28k ports, which sounds a lot, but after reading it through, that is surprisingly a low number. If you are Bluesky scale, you might need millions of ports.
   - Diagnosis skill is something that is going to be super valuable going forward in the AI era. AI can help, but it would be too slow and can never reach the instinct-based debugging of humans (as of now, at least)
4. [Six Characters: Decoding PNR Number and e-ticket system (Ajitem)](https://ajitem.com/blog/iron-core-part-2-six-characters/)
   
   - This is a treat. People should write these kinds of blogs. This is curiosity at its best.
   - PNR is not a unique global identifier, it’s specific to the airline or the entity handling it, it’s for passenger name record.
   - It’s quite ingenious how the currency conversion works, without breaking and keeping simple and straightforward.
5. [AI will be met with violence, and nothing good will come out of it](https://www.thealgorithmicbridge.com/p/ai-will-be-met-with-violence-and)
   
   - Spicy take. True in most senses.
   - The CEO things are pretty messed up, people might fight and get violent, which is inevitable in any direction we go.
   - The start of the post was quite well written, if that holds true, we can actually see where this will go.
6. [I’ve been blogging wrong](https://www.leoniemonigatti.com/blog/ive-been-blogging-wrong.html)
   
   - People are waking up and finding the strength to write authentic content. Sharing human experience, which was the sole purpose of blogs. But it has taken the shape of vanity metrics and technical jargon. We are so back!
7. [Llama Parsebench: First Document parsing benchmark for AI Agents](https://www.llamaindex.ai/blog/parsebench)
   
   - I am a bit surprised that there are no document parsing benchmarks yet? What? Where are the Chinese labs and Msitral and all the OCR benchmarks then. Oh, they might be just OCR is it? Well then that makes sense. For document-specific parsing.
   - Another industry or field demolishing with the report card calculator ready.
8. [Now is the best time to write code by hand](https://sitebloom.ch/writing/now-is-the-best-time-to-write-code-by-hand/)
   
   - Yeah! People are firing now. This is happening. Writing code by hand can be a hobby but it will be a skill that pays like COBOL or PASCAL developers are paid today. Trust in code will be more from human than in agents, that is the bet we are making if that holds, software developers are going to be PHP developers with Lambos in their garage.

## Watched

- [Data Structures Explained by Nic Barker: HashMap](https://youtu.be/y11XNXi9dgs)
  
  - This was a solid explantion of a hash map. I actually thought it was magic, but I knew it was some hash function or hash code, but this actually clears a lot of those magical things.
  - I learnt that we can have linked list sort of a structure to tackle collisions on hashes. That is a really clever way of solving a problem.

<!--THE END-->

- [You are still better at editing documents than AI](https://youtu.be/FqB_4QY6x6g)
  
  - This is a great explanation of why the software interface is skewed for agents. It was never meant for them.
  - The right set of tools and context is really important. It might also depend on the model, since its training data, if it doesn’t have the right examples on how to manipulate docx or xml files, it might skrew up big time. But I have not seen those kinds of issues from proprietary models yet.

<!--THE END-->

- [Vim has a 0 day?](https://youtu.be/zMpn9ICagdE)
  
  - This is evident, and quite not surprising to me. AI agents and slopy review system will make this worse.
  - Vim codebase was hand-chiselled for more than 30 years, it still has these vulnerabilities, now take AI into the game, and it could easily make it worse in a matter of months if not years.
  - Reviewing code and testing it will be quite a skill to have. Security essentials will be key in moving out of this kind of mess. I also read about the diagnosis skills in the read section, right from the Bluesky incident report, so that will be another skill to hone.

<!--THE END-->

- [Anthropic Dunking and PI pilling](https://youtu.be/3DNkDIVKtK8)
  
  - Oh! That was a pure entertainment. Grilling Anthropic is my new pastime hobby, or I should say its Theo’s duty.
  - Anthropic is reduced to Syltherin and worse, for some time last year I was considering them Gryfindor, oh my god, how idiotic I was. I thought they shared the safety scores, and made it specific to developers so it might be nice.
  - But the CEO and the company are some evil propaganda to get funding. Good luck getting rejections.

## Learnt

- Difference in this and super in Java OOP
  
  - I understood that `this` refers to the current object instance and is primarily used to access instance variables, methods, or constructors within the same class. `super` refers to the immediate parent class and is used to access overridden methods or parent constructors. The important learning was not just usage, but when ambiguity arises (e.g., variable shadowing or method overriding) and how these keywords resolve it.
- Internal working of HashSet and its relation to HashMap
  
  - I realised that HashSet is backed by a HashMap, where elements are stored as keys and the hashvalue as its value. This means all properties of hashing, collision handling, and performance characteristics come from HashMap.
- Design Patterns from Spring Boot
  
  - While studying Spring Boot, I learnt that Dependency Injection (DI) is a way to provide an object with the dependencies it needs from the outside, instead of the object creating them itself. This separates what a class does from how its dependencies are created. Without DI, a class directly creates its dependencies using new, which tightly couples it to specific implementations. This makes code harder to test, extend, or replace. DI removes this by depending on abstractions instead of concrete implementations. Its pretty good pattern to be honest.
- While brushing up on concepts on OOP, I also learnt about interfaces and abstract classes. An interface defines what a class must do, and an abstract class defines what a class is, along with some shared behavior.
- The software industry is in extreme ends
  
  - On one hand, we have AI-pilled managers and bosses. On the other hand, we have Java coorporate coder (hand-crafted coders).
  - Like if someone looks at the job market, it’s a chaos of expectations and reality. How can a field have 3+ years of GenAI experience when the field itself is publicly known to be less than 3 years old?
  - I must say, not everything is doom and gloom, there are still companies expecting people to hand-chisel code, and that is where Java enterprise comes in. Maybe that can change my trajectory, the next two weeks are the decisive ones.

## Tech News

- [OpenSource is in danger](https://thenewstack.io/cal-com-codebase-security-ai/)
  
  - This is bad. Yes, we had around 5 supply chain attacks in a period of 2 weeks. That was brutal, but one cannot blame it to open source, and it cannot end.
  - It is the pillar on which the tech giants and the smaller firms are standing on, it cannot be something that you go private for. Its shows the lack of trust on the 80% of the people doing good things, but those 20% of the people yes that is brutal.

<!--THE END-->

- [Google releases Agent CLI for building Android Apps](https://android-developers.googleblog.com/2026/04/build-android-apps-3x-faster-using-any-agent.html)
  
  - This is actually something useful. Opening Android Studio is like baking a grilled sandwich. I have to open and close my laptop 20 times and cannot build the app.
  - With the agent and the tool chaining, this could help me produce a slop of apps.

<!--THE END-->

- [Anthropic drops Claude Opus 4.7](https://www.anthropic.com/news/claude-opus-4-7)
  
  - Oh another 0.1 bump. Nothing much. This is fine, we get impressed, we suspect, and we reject it after 20 days.

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/790) (#790th edition), and for software development/coding articles, join [daily.dev](http://daily.dev/).

* * *

That’s it from the heavy Java-pilled edition. Never thought I would be learning java, but here I am. never thought code can be generated that correctly and that fast, yet here we are. Life has a lot of surprises for us, so don’t get discouraged, you might not know what happens next.

Until that happens for good, keep your hand on the keyboard and brace yourself with the nerves of reviewing code.

Happy Coding :)

[Leave a comment](https://techstructively.substack.com/p/techstructive-weekly-90/comments)

[Share Techstructive Weekly](https://techstructively.substack.com/?utm_source=substack&utm_medium=email&utm_content=share&action=share)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.