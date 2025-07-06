{"author":"meet","date":"2025-06-21","post_dir":"newsletter","status":"published","title": "Techstructive Weekly #47","type":"newsletter", "slug": "techstructive-weekly-47"}

## Week #47

It was a exciting and exhaustive week, in a good way. I was able to almost complete a working AI Chatbot Application, start writing routine for an hour, get inspiration for more project ideas, and a lot of learning in tid-bits around LLMs.

This can’t be more exciting. A week or two before I was drenched in anxiety and negativity around LLMs as replacement for developers, but that’s far from reality and considering them as tools sounds so fun. I feel excited so much so that I feel I have no time to squeeze the energy.

This weekend I was not able to livestream and record any videos due to network issues and also I had a bit of cold. This week hopefully will be doing 2 3-hour live streams on creating projects.

- Mozilla Firefox’s Pocket clone
- LiteLLM Implementation in Golang

Not both of them, either of them will be working this weekend.

I haven’t created any videos, haven’t planned any video as of yet, but have a lot of ideas for writing. I think I will be creating a timer based blog writing, I have to force myself to write a blog. I have lost too many months writing a blog, can’t lose them anymore.

### Quote of the week

> “AI is thinking faster, so we need to think slow”
> 
> — I can’t find who quoted this, but I read it somewhere

This is such a great quote, whoever quoted it, awesome job. AI is producing slop and hit sometimes, it might not be correct every time, since perfection is the enemy of progress. So is math, thinking a problem for LLMs currently, they think a lot, hard, too hard at times, and spiral into overthinking just like humans, but meaningless and not required from machines.

We humans are so funny, we expect LLMs to work perfectly but have trained on the dataset that is imperfect.

We need to sift through the LLMs response, which could be slop or a gem, which ever it is, we need to think a bit slow than what it gave. Reflection is key in the case of AI-assisted anything.

---

## Created

- Bring Your Own Keys AI Chat App ([BYOK AI Chat](https://byok-chat-dev.vercel.app/))
    - GitHub Link: [https://github.com/Mr-Destructive/byok-ai-chat](https://github.com/Mr-Destructive/byok-ai-chat)
    - Vibe Coded the Frontend with Lovable / Windsurf
    - Half baked backend API with Claude and Grok AI
    - Minor Improvements and Bug Fixes with Jules Agent
    - Backend in Python (FastAPI), PostgreSQL Database
    - Frontend with React
    - Vercel + Render Deployment
    - LiteLLM for model provider
    - Yet to add proper authentication (probably with Clerk)

## Read

1. [The art of saying no](https://duncan.dev/post/art-of-saying-no): This post is so accurate in describing the feeling while doing AI assisted development. We have to constantly sift through hundreds of suggestions or ideas to get the one that we are looking for, and that is not easy and comforting at first. However, just like debugging intuitively, accepting suggestions and vibing would become a intuition based activity, it just takes a few misses and hits, some trials and explorations of how these LLMs work.
2. [Never Just](https://www.neverjust.net/): This is absolutely relatable, when we are in complex problem solving situation, and someone says “why don’t we just make it simple”. The person saying this either doesn’t know the problem or knows too much that he simply can do it instead of asking it to others. If this just comes from the surface, it feels a bit hurting, it should be reframed to make it more informative and actionable instead of attacking.
3. [A Linear Algebra Trick for computing Fibonacci Numbers Fast](https://blog.codingconfessions.com/p/a-linear-algebra-trick-for-fibonacci-numbers): An interesting way to compute Fibonacci numbers. A nice matrix multiplication trick.
4. [A tale of two Claudes](https://steveklabnik.com/writing/a-tale-of-two-claudes/): A completely honest and realistic take.
    1. Claude cannot work with Tailwind 4
    2. Claude works well with compiler and gnarly debugging memory related issues
    3. This is really cool to be aware of, we can use LLMs to guide us in the situations when we are not sure. Because it will speed things up and probably reach conclusions much faster and we can then decide if that conclusion was right or not, rather than spending hours in the gnarly bugs.
    4. Someone said it, LLMs are thinking too fast, so we should think slow, what a perfect sentence.
5. [Feedback is not attack](https://ashley.dev/posts/feedback-is-not-an-attack/): Feedback is quite a huge subject and can affect one’s relation in positive or negative way. Its not just about giving feedback, its also about empathy and being able to receive and accept the feedback well.
6. [What would Kubernetes 2.0 be like](https://matduggan.com/what-would-a-kubernetes-2-0-look-like/): YAML should be replaced with HCL, that is a golden point. No like seriously, reading a long yaml file just gives me headache. I prefer JSON Instead. Also the helm package manager is written in Golang, I want to dive deeper into K8s.
7. [This AI Agent should have been a SQL Query](https://www.morling.dev/blog/this-ai-agent-should-have-been-sql-query/): This one was more like Apache Flink comparison and walkthrough, maybe as a example but it was good.

## Watched

- [The State of Authentication](https://youtu.be/lxslnp-ZEMw?si=NHEnmYD5AVQUHOM0)
    - This is so messed up. There are actually three parts in Authentication. I thought Authentication and Authorization, that’s it
    - Authentication is basically “Am I who I say I am?”
    - Authorisation is what I am allowed to do (if i am who I say I am)\
    - Now the third part of Authentication
    - Auth UI: Lo behold, UIs are not my thing, and yes Auth0 and some other probably some other services too.

Double click to interact with video
- [Andrej Karapathy: Software is Changing](https://www.youtube.com/watch?v=LCEmiRjPEtQ)
    - We are in Software 3.0 Era
    - Software 1.0 was code | Software 2.0 was Neural Networks
    - Now we are in the LLM space, were we can send prompts to neural networks that can generate programs, that is wild idea
    - We need to create systems around LLMs, like interface to interact LLMs, He puts it like we are in the 1960s era of Computing. Where the computers were not personal, they were mainframe, large machines that fit on factories.
    - However the LLMs today are like those big gigantic computers that live in the cloud running in multiple clusters of GPUs, though we have local LLMs, the quality and accuracy of high parameter models is not near the local small sized models (they are improving). So we are reliant on Cloud inference like ChatGPT, Claude, Mistral, LLaMA, Gemini, etc.
    - He relates the transition of LLMs with the augmentation to an autonomous agent. This is compared to Iron Man’s suite, which is a mixture of human control along with AI assistance. When he is outside the augmented mode, its his instinct that help him navigate, however, in the agent mode, the AI decides the track.
    - Also the vibe coding term was emphasised which has a Wikipedia entry. Vibe coding is easy to do because we don’t care about the code, we care about the outcome, we can call it product driven prompting. And it makes sense when he said, we need to increase the Human-verification and AI-generation loop. It doesn’t mean, you’ll tell AI to one-shot the entire thing, instead go bit by bit, one thing and one action and feedback at a time.

Double click to interact with video
- [99% of the AI Startups will be Dead by 2026](https://youtu.be/I10_O47P7Zs?si=umAME2_EUvj97uY8)
    - Obivious, Startup is about survival and if all of them are wrappers around just a LLM they might die. But only if they are just that and nothing else. Turns out most startups might just be that.
    - However, many unicorns and valuable startups are built due to trust and alignment of the problem solved by the product and faced by their users, so its about competition there.
    - AI is here to stay and evolve, if the product keeps adapting, it will survive no matter what, the trend is, if people like it, people will pay to use it for comfort.

Double click to interact with video
- Why Internet went down for 2.5 hours on 12th-13th June 2025 (IST Timings)
    - Half of the internet was down for almost 2.5 hours, due to Google Cloud outage.
    - Cloudflare was down due to one service of theirs was reliant on Google Cloud, and that spiralled to all other services. Whoa! dependency hell is real.

Double click to interact with video

## Learnt

- LLaMA actaully stands for [Large Language Model Meta AI](https://en.wikipedia.org/wiki/Llama_(language_model))
- A wired random prompt that surprisingly is not random: [Tweet](https://x.com/MeetGor21/status/1935636326243549328)
    - This is not 100% reproducible but quite the gist of it is similar
    - Pick a number between 1 and 100
          - ChatGPT, LLaMA, Claude, Mistral gives 27
          - Grok AI gives 42
    - Pick a random letter between a and z
          - ChatGPT, Mistral gives M or K
          - Grok AI, LLaMA gives K
    - This is a wired behaviour from LLMs, and it also makes sense since it will try to mimic what it has seen most of the times in its training process.
- Render is the new Heroku for free application deployment
    - I used it to deploy my FastAPI backend for BYOK AI Chat App.
    - Its similar to the old Heroku hobby tier plan, which was vindicated in 2022.
    - I am not sure how reliable these Platform-As-A-Service are anymore after the Heroku sunset of hobby plan, but nonetheless its what we get from a free tier and its worth appreciating such companies are providing free tiers.
    - I am wondering can Appwrite take over its spot, its more of a microservice oriented setup but still it has a lot of potential to compete in the space.
- [Copy and Paste Markdown in Google Docs](https://one-tip-a-week.beehiiv.com/p/one-tip-a-week-paste-markdown-and-copy-to-markdown-in-google-docs)
    - This was a neat trick to learn, especially if you are writing for documentation and personal notes. Also handy if you want to quickly add something special without fluff in the doc. The keyboard centric approach helps in this markdown style.

## Tech News

- [Duolingo launches Chess as a course](https://blog.duolingo.com/chess-course/)
  > [It wasn’t built by a huge team or chess experts. It started with a PM, a designer, and Cursor, an AI coding assistant. No engineers. No dev team.](https://x.com/shweta_ai/status/1935377570091921617)
- Google Cloud Outage on 12-13th June 2025 for 2.5 hours
- [Your Brain on ChatGPT: Accumulation of Cognitive Debt when Using an AI Assistant for Essay Writing Task](https://www.media.mit.edu/publications/your-brain-on-chatgpt/)

---

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-751) (#751 edition) not arrived yet, and for software development/coding articles, join [daily.dev](http://daily.dev/).

[Leave a comment](https://techstructively.substack.com/p/techstructive-weekly-47/comments)

---

That’s it for this week. Hoped for a good week, got one, hoping for a better week.

Thanks for reading :)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](%%share_url%%)
