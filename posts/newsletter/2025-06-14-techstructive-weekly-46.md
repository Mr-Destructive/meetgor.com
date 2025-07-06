{"author":"meet","date":"2025-06-14","post_dir":"newsletter","status":"published","title": "Techstructive Weekly #46","type":"newsletter", "slug": "techstructive-weekly-46"}

## Week #46

Life takes turns pretty quickly. I was here last week, complaining about life and here I am today, feeling grateful for whatever happened to me.

I took decisions last week and by mid-week, I was seeing results, like actual results. And I was just thrilled to experience it, when you decide to change, the universe listens and blesses you with opportunities. I am grateful for the things blessed to me through the week. Even if I fail next week, I have a head start to keep going and break the shackles one day.

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.

Throughout the week, I dabbled in LLMs. Learning about different providers and models.

I spent 20 minutes playing this game before writing this newsletter. I was reading the Hacker News newsletter, and I saw this [link](https://midword.com/) and wow, the game is so cool. A binary search to find a word. A wordle but not quite that.

### Quote of the week

> **“Practice doesn’t make perfect. Practice reduces the imperfection.”**  
> — *Toba Beta*

I have been in the craft of programming for over 5 years, and now I know that it’s about figuring things out. I was naive and all over the place when I started, but slowly I learnt from mistakes and started to feel the patterns in solving problems. I am not perfect, never will be, but will have less muddier and novice like traits in me.

That’s what it means to live a developer’s life, really, adopting and adopting till your hunger and curiosity to know things is satisfied, it will be on you to keep it on till your last breathe.

---

## Created

- [Bring your own keys (BYOK) AI Chat App](https://github.com/Mr-Destructive/byok-ai-chat)
    - Created entirely from Claude and Lovable AI
    - Added Litellm for providers and model selection
    - FastAPI for backend, Bun/React for Frontend
    - Thinking of using a per-user database with Turso
- [LiteLLM listing providers and models](https://meetgor.substack.com/p/litellm-list-all-llm-providers-and?r=1hoe7f)
    - This is a rough draft version of the post.
    - I have worked together the way to list down the providers and models
    - Looking at what else could be done and exploring different ways to query info about a specific model

## Read

- [How I finally understood Docker and Kubernetes](https://medium.com/javarevisited/how-i-finally-understood-docker-and-kubernetes-5debb13cacfe)
    - This was my pick of the week. I understood the reason why Kubernetes exists
    - I knew the concept of Docker (it could be because I have used it extensively in the past to create projects as well in my internships to deploy APIs and apps)
    - But the concept of Kubernetes is like a black-box. But the author’s explanation style and simple example made it clear.
    - Kubernetes is like
          - Container Image > Deployment > Pod > Service
    - Container Image is the actual image of your app that you want to run, maybe it has multiple of those.
    - Deployment is like defining what and how many (other things too) to run.
    - Pod is like the actual unit of containers; in itself, it has no control, it just runs whatever was given to it.
    - Service is like the layer that exposes it to the world, maybe the network, the other containers, which are like a configurable exposure of the network.
- [I am disappointed in AI discourse](https://steveklabnik.com/writing/i-am-disappointed-in-the-ai-discourse/?ref=dailydev)
    - Wow, at least someone feels like me. This was so relatable to read on a Sunday afternoon.
    - This feels relatable to feel like how can people draw conclusions and biases when the technology is evolving everyday.
    - People are either fully bullish on AI, i.e. by 2026, no developers needed, to some saying AI is still crap. Both of them don’t know what AI is capable yet. It is not that bad, compared to a year ago, it’s really good at generating almost accurate code, but not quite the touch that it can be untouched without a developer glancing and ficing it.
- [The Software Engineering Identity Crisis](https://annievella.com/posts/the-software-engineering-identity-crisis/)
    - The quote that sums this well is this:
      > Perhaps the most valuable skill in this new landscape isn’t prompt engineering or systems architecture, but **adaptability** - the willingness to evolve, to learn new skills, and to find your unique place in a rapidly changing field.
    - Adoption is a skill that you can hone and be an engineer of the AI revolution
    - The author thinks that there will be engineering in terms of managing and overlooking Agents or AI Systems. Also a good point mentioned is “The scope of engineering is not shrinking, it’s expanding”. I can see this point coming true in some sense as the level of understanding to drive these AI systems is better suitable to programmers and engineers. Since they already deal with the pain of solving problems. AI is no different.
    - The first part that hits me the most is the loss of joy. I talked about it in the previous weeks. I don’t want to repeat that rambling, but yes, that somewhat feels a bit uncomfortable to digest.
    - An excellent post in navigating this AI landscape while maintaining the core feelings of a developer and where it breaks.
- [Programming with Agents: Sketch.dev](https://sketch.dev/blog/programming-with-agents)
    - Great post on agents, good and expected take on agent definition
    - David also narrates the need for agents and why they are suddenly in the hype. The function calling thing just blew the AI hype and at that time, the LLMs were not ready or trained for it. But in 2025, those LLMs are optimised for it, making it a great ecosystem to work towards.
    - This actually raises questions: Do you really need the craft? The IDE? If the agent can do it for you at the speed of a prompt? Yes, LLMs are not gods, they need assistance, and when they go berserk, its on the developer to hone on his tools and hack it out of the mess.
- [The Gentle Singularity: Sam Altman’s Blog](https://blog.samaltman.com/the-gentle-singularity)
    - Sam thinks that GPTs are powerful than any human lived on earth, is that really accurate? Maybe, in terms of knowledge, but that’s not truly knowledge. It does have billions of weights that somewhat represent the knowledge, but can it make sense of it? No, not yet. With tools and reasoning, maybe, but not quite without proper instructions.
    - The other stuff is just sci-fi future prediction. I don’t think that is true, but his vision is a bit daunting if that is slightly true.
    - He also mentions how many watts are consumed per GPT call, which is hilariously alarming.
- [Anyone can cook: How 37 Signals hired a junior developer](https://world.hey.com/jorge/anyone-can-cook-c6346f84)
    - Intrinsic motivation is greater than an educational degree; this is true for any professional. Because the degree won’t sustain you longer, if there is a fire within, that will carry on.
- [AI IDEs Free Tier War](https://ainativedev.io/news/ide-free-tier-war-windsurf)
    - Windsurf is surprisingly cheap, but they don’t have access to models like Claude 4 and others.
    - GPT 4.1 is good and all, but that’s a little supbar with the standard of Claude
    - I don’t know if that’s just me or I feel confident in copy-pasting a file from claude than from GPT

## Watched

- Janvi Kalra From Software Engineer to AI Engineer at OpenAI
  [The Pragmatic Engineer](https://open.substack.com/pub/pragmaticengineer)
    - This was a really inspirational interview. Its giving me hope in continuing what I am doing currently; being excited and willing to develop stuff, doing the due-deligence to research and solve things.
    - I think I now understand my flaw; I am not specific. I am all over the place.
    - She was very specific in boiling down what she was interested in, listing down the 50+ companies in that space, and getting the interviews. It shows the care and love for the craft.
    - I am not sure how to do that, really, because I think I would miss out on certain roles where I might feel excited. I am really a bad problem solver, I don’t know what I should focus on. I think I need to write more in order to nail it down in the coming weeks. Let’s do that and let’s see where it takes me.Double click to interact with video

- [My AI Skeptic Friends are Nuts: Review by Theo](https://youtu.be/uqRF4IszorU?si=wUzhdfrD9W5OD2uI)
    - I read the article last week, and surely it was AI bullish, but the GraphQL, Web3 bubble value to hype ratio just made me relate to this AI hype.
    - The hype is too high, but the value bar is high too; we need to get the value, which is a bit easier, but avoiding the hype and fluff is a bit tricky. I think by being hands-on you can separate the fluff from value.

Double click to interact with video
- [Are we feeling it now? The AI Model Fatigue](https://youtu.be/YwsHRMNZjjU)
    - There are too many models to keep track of, I had raised this concern in the past two weeks, but no one listens. Here we are in an LLM model apocalypse.
    - Maybe it’s a positive one, but we need to slow down a little.Double click to interact with video

- [Dopamine Driven Development](https://youtu.be/AWZ9AgjPfwo?si=Il7M0Ba0w8nWjh7q)
    - Passing tests, GitHub action tick, first try especially, different error message all gives dopamine, and that is not a cheap dopamine, its a value to effort cost.

Double click to interact with video

## Learnt

- Using LLMLite to list different models and providers
    - We can use `models_by_provider`, which is a dictionary of provider keys and a list of strings representing model names provided by the provider
- Over 3 years of only Python, I didn’t know ABC are abstract classes
    - I am ashamed that ABC and abstract_method are a thing in Python
    - Never really used them, never made sense, but yes, now it makes sense to use and see them

```
from abc import ABC, abstract_method

class Shape(ABC):
    @abstract_method
    def area():
        pass
   
    @abstract_method
    def perimeter():
        pass

```

## Tech News

- [Mistral releases Magistral](https://mistral.ai/news/magistral), a thinking/reasoning model
    - It thinks like crazy, I saw from a clip.
- [Open AI releases o3-pro and drops the prices by 80%](https://help.openai.com/en/articles/9624314-model-release-notes)
    - mini - standard - pro is the ecosystem Open AI has created for its thinking models
    - o1 - o3 - o4
    - By the time o3 standard or pro was released, o4-mini was already to the same level
- [Did Apple just give up on AI?](https://www.apple.com/newsroom/2025/06/apple-supercharges-its-tools-and-technologies-for-developers)
    - AI and Apple can’t be on the same boat at least in 2025, it seems.
    - Sad to see it, but they are truly behind in tech right now. Their hardware is good, but they have to grow significantly to stay even competitive.

Phew, what a week, nothing much in releases, but the cloudflare x google cloud outage just wrecked havoc in almost half IT companies. Many unicorns faced this downtime for a while. That’s a wild week till June 14th

---

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-750) (#750 edition), and for software development/coding articles, join [daily.dev](http://daily.dev/).

[Leave a comment](%%half_magic_comments_url%%)

---

That’s it for this week. Hoped for a good week, got one, hoping for a better week.

Thanks for reading :)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
