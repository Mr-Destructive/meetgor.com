{"author":"meet","date":"2025-04-04","post_dir":"newsletter","status":"published","title":"Techstructive Weekly #36","type":"newsletter"}

 
 ## Week #36
 
 Was quite a work-holic week, did a lot of things this week. Debugging, prompting, scripting, flying through a random codebase and number crunching. I felt good, I am genuinely curious about Agents now and want to experiment a lot this weekend. Will have a video or two about it.
 
 Finally it was a week, a little slow down after a 2 back to back weeks of AI launching and breaking the expectations. The image generation, gemma models, gemini 2.5 and mistral OCR, and what not have been released in the last stretch of first quarter of 2025. Phew! it was a heck of a quarter, that went in a jiffy.
 
 ### Quote of the week
 
 > “Its hard to beat talent when talent is driven by curiosity, but impossible to beat talent when it is driven by hard work“
 > 
 > — Meet Gor
 
 Believe me, talent and hard work is deadly combo. You don’t want to be competing with that kind of group of people, you might easily slip out, if you are just hard working or just talent won’t be enough, there needs to be a fire, a purpose, a curiosity burning that will lift even the one without feet and help fly without wings.
 
 ---
 
 ## Created
 
 * Pydanitc-ai-compatible interface for Meta AI LLM (llama)
     
     I created a pydantic class that will take the Meta AI LLM API and use it as a model instead of an API. I also need to make it tool-callable, since the api stucture is not designed for tool-calling, it needs to be done in a hacky way. But if that works, I can make the LLM available as an agent for free to anyone via the Meta AI API. That is a revolution.
     
 * API for converting a image URL into base 64 string representation in Golang and Javascript → https://imgbase.korogi5431.workers.dev/?url=
     
     I created this to understand the process of encoding the image blob to base 64, first I get the Image URL, then fetch that image, sometimes, the Image might be encrypted or locked in due to authentication, so it will fail, however if the image is available publicly we can fetch it and load the bytes to encoded in base 64. This way, if the resource i.e. URI breaks or in future can’t reference it, the image link will be broken, can’t view the image. With the base 64 encoded value, the image is constructed in place, so we are not locked in with the network call. There are trade-offs, the string is quite big!
     
 
 ## Read
 
 * [A](https://www.piglei.com/articles/en-programmer-reading-list-part-one/?ref=dailydev) [programmer’s](https://www.piglei.com/articles/en-programmer-reading-list-part-one/?ref=dailydev) [Reading List: 100 Articles I Enjoyed Part 1](https://www.piglei.com/articles/en-programmer-reading-list-part-one/?ref=dailydev) (1-50)Have not read any of those, but seems a great place to bookmark for never reading (bookmarks are bad design actually in 2025)
     
 * [The 13 software engineering laws](https://open.substack.com/pub/zaidesanton/p/the-13-software-engineering-laws): This is a great read, full of controversial sometimes. Deadlines, answering wrong is the better than asking the wrong question, that was funny.
     
 
 ## Watched
 
 * [*I Ranked every LLM based on vibes*](https://youtu.be/3yrAK2hMWw8?si=rOnFCMbfbGn-wU_j)
     
 
 {% embed https://youtu.be/3yrAK2hMWw8?si=rOnFCMbfbGn-wU\_j %}
 
 That is really cool, I didn’t knew gemini flash was that good, need to check that out really. I haven’t tried o3 and not sure would be quick enough to get results out. I have been using claude and gpt extensively for qick fixes and even brainstorming ideas. Pretty good tier list to be honest.
 
 * [Understanding MCP from scratch](https://youtu.be/CDjjaTALI68?si=Dkse1dPU96MherDQ)
     
     This cleared a lot of things.
     
 * * RAG = Context + LLMs
         
     * * Agents =Tools + LLMs
             
         * MCP Server/Client = Context + Tools + LLM
             
 * {%embed https://youtu.be/CDjjaTALI68?si=Dkse1dPU96MherDQ %}
     
 * [Raising an Agent - Episode 1](https://youtu.be/Cor-t9xC1ck?si=8SuaiF8kHdI8hrQw)
     
     Fascinating. A sourcegraph AI Agent is around the corner. Wow this could be the first one to actually replicate the editing experience. We know LLMs can’t really drive the code, so let them do the chore work while we think. That is the approach Sourcegraph will be taking, they are not completely saying LLMs are bad, they are infact bullish on LLMs, claude can do almost anything provided the tools, so LLMs with tools and context is a big brain move.
     
 
 {% embed https://youtu.be/Cor-t9xC1ck?si=8SuaiF8kHdI8hrQw %}
 
 ## Learnt
 
 * How to convert a image URL to base 64 image (obviously loading the actual image bytes first)
     
     So, we need to convert the bytes into base64 encoded letters (A-Z,a-z,0-9, and + and /). This is really cool to represent blobs of bits, this reduces the space of storing images, however its quite big for high quality images, as it will have a lot of information, like very granular pixel details.
     
 * Creating a model wrapper of Meta AI API: for pydantic compatible api
     
 * Creating a Cloud function on Appwrite Cloud in Golang
     
 * Creating a Cloud function worker on Cloudflare in Javascript
     
 
 ## Tech News
 
 * [*Turso* launches offline sync](https://turso.tech/blog/turso-offline-sync-public-beta): Sync your data to the primary database next time whenever you have the internet connectivity. This will allow the data to be saved on the local copy(replica) and get pushed to the primary database which will sync the primary and the replica. Really exited to try it out.
     
 * [Pico LM](https://www.picolm.io/): Demystifying how Language models learn. This is really cool, I want to understand how it is trained and what improves or what deteriorates the accuracy.
     
 * [Anthropic launches program for Universities and Students](https://www.anthropic.com/news/introducing-claude-for-education) - Claude for Education
     
     This is adoption, survival of the fittest, finding the ways in which people will reap benefits. Anthropic is solid step ahead of all the LLM providers in that race too, it is already solid at tool-calling, and now its spreading its legs.
     
 
 For more interesting articles, check out the [hacker-newsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-740/) for the week edition [#740](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-740/), for even more software development/coding articles, join daily.dev.
 
 ---
 
 A promising week ahead, lot of experiments to be done, a lot to be recorded, don’t have the mental energy but I will and need to push through to make myself better. I am really curious about AI Agents, I am emphasizing this again, because this is a field that needs attention. Because Attention is all you need!
 
 See you next time :)
 
 Thanks for reading [*Techstructive Weekly*](https://techstructively.substack.com/)! This post is public so feel free to share it.
 
 Thanks for reading [*Techstructive Weekly*](https://techstructively.substack.com/)! Subscribe for free to receive new posts and support my work.
