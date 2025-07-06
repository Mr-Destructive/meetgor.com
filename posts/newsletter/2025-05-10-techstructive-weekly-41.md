{
    "author":"meet",
    "date":"2025-05-10",
    "post_dir":"newsletter",
    "status":"published",
    "title": "Techstructive Weekly #41",
    "type":"newsletter",
    "slug": "techstructive-weekly-41"
}

## Week #41

A slow-moving week overall. I had some spiritual realizations, not major, but enough to clear mental clutter and ease the overthinking through action, without worrying about results

The past weekend was too boring; I was not able to do anything fancy. Didn’t live stream, nor create a video. I think videos are not my thing for the moment, It takes too much energy to edit and polish my voice, and I don’t see enough return on that investment right now.

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.

That gives me some time to crunch some articles and blog posts. Will be focusing on that and see if I can get more value than videos, that’s my theory for my skills, but I need to experiment.

### AI Slop

Ok, we cannot spend any week without rambling about AI and LLMs, can we? This week, I read a few articles about vibe coding and the jazz people have created that it is going to replace developers. To that, I say: hell no. Yes, it can write code, well-documented even, but can it solve the whole problem? Nope. Not yet! It’s way far from that, however, if you give it all the pieces, and you think the solution and give it nudges in certain directions, then? Yes, to some extent.

We are on a cliff right now, we as developers need to learn to use LLMs to our advantage, but on the other side of the cliff is the part which is almost LLM writing the code, which looking at the speed of accuracy of code being written by different models, its not far but not quite there yet. It is at the stage where only a developer can steer the wheel; once a layman can steer the wheel, developers are doomed, but till then, till the laymen realize that developers are doomed, we are going to rule, take out their money for sure.

### Quote of the week

> “It is never the tool that decides. It's the hands-and the heart-of the one who wields it.”
> 
> ― Kevin Sands

AI right now feels like that tool, if you give it to a layman, normal person, they will call it shit for coding, but if you give it to a good developer, he/she will say that I want that in my tool belt always. Creating software is a craft, an art, and it is suited for those learning and caring about the craft and tools, not those looking for expensive or shiny tools. Try, fail, learn, improve, that is LLM at the moment for coding:

- Don’t blindly accept suggestions, take time to reflect, it has already saved you a bunch of minutes by writing, may as well read it and understand and think
- If you don’t know a thing that it added, search and start a chat with LLM with search documentation tool
- If it fails, retry, but think why is it failing, read the error message, if you are experienced, you can fix it quickly than an LLM, don’t waste time and tokens, do it yourself.

In essence, having a general common sense of using LLM tools is vital in today’s software development ecosystem. Might not be relevant to software, but I am not sure.

## Read

- [What even is Vibe coding](https://ashley.dev/posts/what-even-is-vibe-coding/)
    - I completely agree with the entire post, this is the best take on vibe coding. I have experienced this and also mentioned it in my [post](https://meetgor.bearblog.dev/developers-side-project-graveyard-with-llms/). AI can handle the boring part, so that the developer can focus on the important, the heart of the craft. Whatever AI creates before this needs to be seen with a grain of salt, i.e., needs testing and review, it’s not a real person writing code after all.
    - It’s more of a tool than a replacement. LLMs are not there yet, it’s far from reality and close to boilerplate or already solved problems, not innovation.
    - If AI is giving you the entire code, then it is likely that the idea that you have for the software is already kind of there, not entirely, but vaguely there. For real innovations in the software, you still would need a developer.
- I read the book The War of Art, Why?
    - I asked Chat GPT, based on my conversation and everything you know about me, to give me a book recommendation. And it gave the war of art
    - And Oh my god! I loved the book, it was a quick one. I don’t usually read non-fiction, self-help books because I already have too many technical things in my work and side projects that I can’t bear reading those in my spare time.
    - But that book was a huge shift in mindset. We are all battling resistance. The bigger the resistance, the bigger the calling (the good task or intuition). We have to defeat that enemy, and it’s not easy, but with a war, it can be defeated. And it’s not an enemy once you defeat it, it’s done; rather, it’s a daily war, it will come in the places you are comfortable.
    - This prompt was a pretty good use of LLMs, to be honest. If you can tune it to your needs, LLMs are a blessing; if you try using it for your replacement, you are destined to suffer.

## Watched

- A case against SQL
    - I agree to the points that SQL is a bad abstraction of programs, we can use strings to query the database
    - We should use the actual interface to access the database to query or mutate.
    - But on the scale of complexity and ease of use, SQL outweighs the former
    - We can think of SQL as Python and the native way as C; both are viable, but one is easier to write, and the latter is like shooting in the foot.

Double click to interact with video

## Learnt

- Removing duplicates from a list of dictionaries
  ```
  def remove_duplicate_dicts(dict_list):
      seen = set()
      unique_dicts = []
      for d in dict_list:
          dict_tuple = tuple(sorted(d.items()))
          if dict_tuple not in seen:
              seen.add(dict_tuple)
              unique_dicts.append(d)
      return unique_dicts

  ```

I didn’t code much this week. A bit of slow week, I don’t quite like this type of weeks, but have to go through them. However, I played with Lovable and Bolt for my website overhaul, and Lovable seems to be great, but very limited credits. Bolt is quite clunky and gets interrupted in the middle. very bad experience with Bolt.

## Tech News

- [Gemini 2.5 better model for coding tasks](https://developers.googleblog.com/en/gemini-2-5-pro-io-improved-coding-performance): Gemini is really taking the lead here. Anthropic might lose the bet on the code quality vs cost race here if they don’t launch a great model in a couple of months (We all know they will)

- [OpenAI’s Sora becomes available for free users for image generation](https://help.openai.com/en/articles/10877094-creating-images-on-sora): I heard this yesterday, and I thought it was not there already seems it was there since 1st April. Not sure.
- [Open AI buys Windsurf for $3 billion](https://www.bloomberg.com/news/articles/2025-05-06/openai-reaches-agreement-to-buy-startup-windsurf-for-3-billion): I watched an video of Thoe comparing the usage of these AI IDEs/editors, and Windsurf seems to be low at 4% adoption compared to almost 40% of VS Code and 30% of Other IDEs, and around 20-30% of Cursor usage. Which makes people think, why windsurf? Maybe that is a production-ready enterprise-grade IDE? I have used it and had no problems using it, but Cursor is good too.

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
