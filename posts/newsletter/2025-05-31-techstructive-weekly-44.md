{"author":"meet","date":"2025-05-31","post_dir":"newsletter","status":"published","title": "Techstructive Weekly #44","type":"newsletter", "slug": "techstructive-weekly-44"}

## Week #44

It was week of comeback. 1 videos and 2 live streams, back on track. Though lost the steam in the middle of the week. Still fresh and excited to get hands on projects this weekend.

Appwrite Sites? Android Apps for LLMs? AI Agents with some specific goals?

We’ll roll the die and let’s nature take where it wants to. We’ll start though, I realized it at the end of last week. No brainer right, once I start something, it quickly makes me addicted to completing it and the thing in between is curiosity and consistency. Not hamering buzzwords, but truly, I just started to record the video for pydantic AI and I created a few examples, within an hour I had 7 examples created. Some of which I had to discard, but nonetheless, enough material for a video to record.

Hit the record button and started speaking, some 1-2 hours later, was completed with 1 and 20 minutes of raw footage, next day edited it and published. I t turned out to be a 45 minute video, not bad, it only has got around 30 views in 4 days, but nevermind that. I learned a lot and got a good understanding of AI Agents with Pydantic AI.

A good week, looking to the next week for more projects created and feeling more satisfied with my efforts. Because that is what we can do, focus on the efforts, forget about the results.

### Quote of the week

> "You can’t connect the dots looking forward; you can only connect them looking backward."

I created the examples and then recorded the video (i.e. the thing I wanted to show). I knew what I wanted to explain in that video, a general direction, but after creating the examples, I really got the understanding of what I actually meant and what I am creating, that was working backwards and figuring it out on the go.

Simialrly in life, we don’t know the future, we can’t. We know the direction, but not the actual destination. We know what we have done, and by doing what you feel right right now is the way to move ahead, overthinking, perfectioning the thing, will just delay it. Just do it, follow the intuition.

---

## Created

- [Creating AI Agents with Pydantic AI](https://youtu.be/_m2YpvsdxSA?si=gkOMw1SkVHKPeXLK)
    - Understanding the basics of Pydantic AI
    - Using the output type to get the structured response back from LLMs and AI Agents
    - Using tools with LLMs called Agents with Pydantic AI
    - Using local as well as cloud provider models in Pydantic AIDouble click to interact with video

### Live Streamed

Completing the database sync and admin routes for CMS in the Static Site Generator in Golang

On Saturday, I was expecting to create Appwrite sites, but was surprised to find it in a waitlist, so decided to move on and complete the CMS for my SSG

I created the script to sync the database and posts on the github repo, it had a issue with the SQL query and I fixed it later after the stream.

Double click to interact with videoOn Sunday, I continued with the CMS part, and added the workflow with the sync script to write the contents to the GitHub repo.

Also added authentication to the Sync DB and Trigger build GitHub actions to only the authors in the site with Netlify Cloud functions.

Finally completing the SSG + CMS in Golang called Burrow.

Double click to interact with video

## Read

- [The Copilot Dillusion:](https://deplet.ing/the-copilot-delusion/)
    - This was like a word of caution of getting too much reliant on AI for coding
    - Coding with AI is all well and good, but the thinking part still is in the heads of the developer. It can’t think beyond a certain capability.
    - Till now, its fair to say from this post that, AI is most safer in the hands of developers than laymen for coding. It’s not gone that far that managers can vibe code and ship everyday.
    - Maybe someday it will, but there will be the need of developers in pushing it and nudging it to get the most of it.
- [Am I online?](https://antonz.org/is-online/) A useful way to identify your servers are connected to the internet.
- [Read out the thing you have written](https://martinfowler.com/bliki/SayYourWriting.html):
    - This is like a thought or tip from Martin Fowler
    - He suggests to read out (at least lip movement) the thing we have written. I really like it and have not noticed it yet, but yes that is true.
    - We get a different perspective and feel for the draft that we have written. Its almost like reviewing your code locally vs on GitHub. You get something while reading out that you don’t just reading in mind.
- [Explaing Transformers in simple words:](https://janvikalra.substack.com/p/explaining-gpt-4s-secret-sauce-transformers)
    - This was a simple yet effective explanation of the transformer model architecture
    - It skipped the middle part, it explained the input and output part well though, enough for someone to get curious and fall in the rabbit hole of exploration.

[![](https://substackcdn.com/image/fetch/$s_!USUs!,w_56,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F01061d86-5a43-45c4-be84-06c48dcede90_704x704.png)janvi kalraExplaining GPT-4’s Secret Sauce: TransformersTransformers are the secret sauce that makes Chat-GPT, DALL-E, and other GPT-based systems so powerful. No, I'm not talking about Optimus Prime and his Autobot pals 🚗🤖 - I'm talking about the neural network architecture. Transformers are the "T" in GPT-4 (Generative Pretrained Transformers v4), and their development has enabled machines to understand …Read more2 years ago · 13 likes · 1 comment · Janvi Kalra](https://janvikalra.substack.com/p/explaining-gpt-4s-secret-sauce-transformers?utm_source=substack&utm_campaign=post_embed&utm_medium=web)
- [Let’s Fing Go:](https://notrab.dev/blog/lfg)
    - This was a reflection and a way for switching and learning Go from Javascript. It is evident that Golang is the almost the perfect tool for backend servers and quick simple applications. Not saying its bad for complex applications, but one must be equipped with the right set of tooling and mindset to head into a large scale application, rather than complain about writing bunch of err != nil, skill issues.

### Interesting Things

- [Radicle](https://radicle.xyz/): The GitHub alternative
    - This week we had a GitHub outage, and on twitter I read about someone complaining, how has someone not figured out a GitHub alternative.
    - The tweeter mentioned that it was just a

## Watched

- [Inside an LLM](https://youtu.be/wjZofJX0v4M)
    - This 3 videos in the series were so much valuable. It helped me understand the mathematics and the architecture behind LLMs. Its quite fascinating
    - I am now thinking about explaining these concepts to laymen because I want them to understand what they are actually interfacing with are just mathematical numbers and nothing human like robots.

Double click to interact with video
- [MIT Lecture: LLMs introduction](https://youtu.be/ZNodOsz94cc)
    - It had me till the part “LLMs can do math”, no please no.
    - All the parts of the videos were great, it touched upon almost everything about LLM and the capabilities it has, helped understand the difference of the actual model and the interface we are interacting with.

Double click to interact with video
- [API Gateways](https://youtu.be/7-6F3b14baA)
    - This was a concise explanation of API Gateways
    - API Gateways are simply a way to route your app to the microservices you have, and also some gluecode like middleware and rate limiting which might be repeated across most of the microservices.

Double click to interact with video
- Primeagent tried AI Video Services
    - Just for entertainment, this was funny though (some jokes went over my head though)
    - Veo 3 and Luma something was great at almost realistic and relatable videos
    - There were services that were clearly there and others were pieces of shit.
    - Nothing in between because average sucks.Double click to interact with video

## Learnt

- Comparing values in Python. We can also compare any type with the equality operator.
    - I thought I might get an error if the value is None and I am comparing the value with a string, but I was wrong
    - In Python, the equality operator can allow any type of expression to be compared with any value.

## Tech News

- [Anthropic finally launches the voice mode in its model](https://support.anthropic.com/en/articles/11101966-using-voice-mode-on-claude-mobile-apps)
    - Finally anthropic is adding voice mode. I love the way that anthropic comes at everything from behind and takes everything at it head, hopefully here will be the same.
    - The code part, the thinking part, the editing part of code is so reliable in Claude, however the search part I am not sure of.
- [Mistral launches AI Agents API](https://mistral.ai/news/agents-api)
    - Every company is launching remote agents
    - Jules from Google, Codex from OpenAI
    - Now mistral, they have similar things
          - Python code environment
          - GitHub integration
- [Deepseek R1 with the new update to its model](https://huggingface.co/deepseek-ai/DeepSeek-R1-0528)
    - This seems to be inching slightly up and a less heavier model than devstral or other models runnable in a single GPU or limited RAM
    - Let’s see how community tests and where it ranks over the week

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-748) (#748 edition), and for software development/coding articles, join [daily.dev](http://daily.dev/).

---

So that’s a wrap from the #44 edition of the newsletter, which has turning out to be my reflection of the week, which it is. Some questions that we can discuss about

- Agents are great, still people hate when it is mentioned, what do you think? I meant the tools and LLMs in the loop
- Are you using Remote agents while coding? some quick tasks or prototypes?
- Are you using local agents like claude code, amp, wrap? how do you handle costs?
- What are your thoughts on AI Video generation? do you feel scared about it

Let’s try to wrap our heads around these questions, since we need sometime and mental space to explore the huge impact AI is having on us. TIll then, keep coding.

That’s it from this 44th edition of my weekly learning. I hope you enjoyed it, and leave comments on what you think about some of my takes or any feedback.

[Leave a comment](%%half_magic_comments_url%%)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](%%share_url%%)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
