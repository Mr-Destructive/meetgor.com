{"author":"meet","date":"2025-06-07","post_dir":"newsletter","status":"published","title": "Techstructive Weekly #45","type":"newsletter", "slug": "techstructive-weekly-45"}

## Week #45

This week was exciting at first, but slowly it went on the extreme bad spectrum for me, leaving a spot for taking a bold decision.

Almost on the verge of a burnout or a big decision. Its been 3 months now, the first sparkle was created and now I think its high time to take the big decision. It’s no longer rash, but it’s a long thought decision.

Anyways, I don’t want to talk about it. Let’s see in the next few weeks where the wind take me.

I explored the AI Agents by creating a local agent in Python from scratch with the Meta AI LLM. It was a fun one, list, read, and edit files tools provided, I also wanted to add shell command execution. Will be doing it next weekend. Also had explored the LLM-plugin for Meta AI LLM wrapper.

### Quote of the week

> "We were born in order to see and listen to the world. It’s a powerful notion, with the potential to subtly reshape our view of everything."
> 
> — Mentioned in [Sweet Bean Paste by Durian Sukegawa](https://www.goodreads.com/work/quotes/26402563)

See and Listen is a very deep thing to do, it just changes everything, it eases the soul, removes the burden and expectations from one’s shoulder. The quote/sentence is simple, yet it hits home for a person feeling lost in life.

---

## Created

- AI Coding Agent with tools from scratch in Python
    - Meta AI API client
    - Tools like read/list/edit files
    - Used the blog from ampcode (Sourcegraph): [How to build Agents](https://ampcode.com/how-to-build-an-agent)
    - Livestreamed it
  Double click to interact with video
- Created LLM plugin for Meta AI API
    - LLM is a cli created by Simon Wilson, its pretty easy to use and supports a whole bunch of models and features like tool calling and image parsing, etc.
  Double click to interact with video

## Read

- [My AI Skeptic friends are all nuts](https://fly.io/blog/youre-all-nuts/)
    - This was a hot post on Hacker News for a week, rightly so, it is partial and follows an experienced developer talking about how LLMs are useful as a tool and not a job replacer.
    - I love one point, which is this quote:
      > Often, LLMs will drop you precisely at that golden moment where shit almost works, and development means tweaking code and immediately seeing things work better. That dopamine hit is why I code.
      I have talked about it previously, and it resonates with me too. I haven’t tried it to the extreme yet. But this is what it is supposed to be doing, if done and used correctly.
    - Caring about the craft is what the author rightly meant to keep LLM out of the loop. He compares with the woodworking; if you treat woodworking as a hobby, then you should care about the tiny details and fine refinements. That is not something for LLM to take care.
    - Junior level output of these LLM is a bit scary. It makes juniors a place of bother.
- [AI changes everything](https://lucumr.pocoo.org/2025/6/4/changes/)
    - Everyone seems to be positive about AI, I am not against it, its good, but its moving too fast.
      > I encourage you not meet that moment with cynicism or fear: meet it with curiosity, responsibility and the conviction that this future will be bright and worth embracing.

- [AI is rotting my brain](https://blog.anniesexton.com/AI-is-rotting-my-brain-203b150f3dfa80d89d5dcee47c79fecd)
    - “We couldn’t stop making art even if we tried,” that hit me, it makes me feel alive and have purpose in life again.
    - Yes, learning about LLM and how it is working is super fun, its quite interesting and rewarding. Then that makes the point that LLMs are not bad, it’s the adoption by people that matters.

- [My Engineering Craft regressed](https://lemmy.ml/post/30100312?ref=dailydev)
    - This is literally me. I can completely get the words of the person. Its exciting to work on side projects and contribute to open source projects, but really no one cares, but I have done it so far for personal pleasure.
    - If something makes me learn new things and go from knowing nothing to fixing a bug, the feeling of finally being able to go through that and see the green tick is so satisfying that the “who cares” question is just out of the equation. The feeling cannot be compared with anything.
    - I also feel a bit sluggish when solving LeetCode, I hate it for some reason. I have barely done any problems on leetcode, let alone grind it for weeks and months. I don’t quite understand the purpose of it. Just to get the job? To improve problem-solving skills? Maybe, but that’s just not how my brain likes to solve problems; my brain works when it sees the problem.
- [First User Framework](https://vishnubharathi.codes/blog/first-user-framework)
    - If there is no user, that is not a solution; it is a problem you are inventing yourself.
    - If you are developing something for yourself, don’t do it right away. If you have the time and capacity to keep up with the learning and work-life balance, then do it without thinking. But if you are already burnt out, maybe just don’t, it might lead you nowhere.
    - If you have one user, start building it. If you have at least a user, there is already a feedback loop developed once you build it. The instant feedback input and the validation part makes it addictive.

- [Styling the icons for HTML date and time types: Cassidy Williams Blog](https://cassidoo.co/post/input-type-date/?ref=dailydev)
    - This was a specific tip, but goes on to make me understand the input types are so rich in HTML. We can do so many things. Some of them might be browser based, and might be specific to certain versions, but still being able to do with bare bones HTML is amazing.

## Watched

- [AI is coming for your job](https://youtu.be/RIvIpILrNXE):
    - It’s far from that, as per Simon Wilson, the people with the domain expertise leveraging these tools and LLMs will be at the forefront and take the most advantage of the assistance.
    - It has cut the development or creation speed to almost half or even lower. This takes off the tedious efforts or menial work and puts humans heading the creative space.

Double click to interact with video
- [Theo Browne on Development and Career](https://youtu.be/oWGtP1iNlZM)
    - Be curious, be active
    - Content is bullshit, don’t consume too much, be actively building and enthusiasitic about the craft

Double click to interact with video

## Learnt

- People want LLMs everywhere, but the moment they add LLMs (or any shiny new thing), the cost of maintaining it creeps up, and in the end, we have to think everything back and put duct tape to limit their use of those.
    - I am annoyed by that, it was never explained why there was a need to add LLM everywhere. If everyone is following the trend, then I understand the reasoning behind it, but if we didn’t consider the consequences, then it becomes a nightmare.
    - Its like release > release > rollback > rollback
    - What was the use of that? wasted time, energy, money, and most importantly, trust.

- Managing API is critical, a couple of changes and boom, the backward compatibility strikes in front
    - Versioning is good, but it shows the level of lack of commitment of the company or developer to support or make breaking changes
    - I had worked with an API, one that I was really new to the system so it was broken and not well thought out, but later now it was revamped for another use case.
    - Somehow, the older API was used again, and by the time the new API was being used, the older API almost completely broke off.
    - I learnt that making changes in this type of environment is critical and needs testing and versioning in the worst case.

## Tech News

- [Anthropic releases Claude Explains](https://www.anthropic.com/claude-explains): An AI-generated blog about its user interaction and usage.
    - The site seems to be taken down by Anthropic, maybe it created a backlash.
- [Cursor launches 1.0](https://www.cursor.com/en/changelog): This makes Cursor a big product.

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-749) (#749 edition), and for software development/coding articles, join [daily.dev](http://daily.dev/).

---

That’s it for this week. A bit of harsh week (Friday deploys ruined everything).

let’s hope for better for the next week.

Thanks for reading :)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.

Thanks for reading Techstructive Weekly! This post is public, so feel free to share it.

[Share](%%share_url%%)

[Leave a comment](%%half_magic_comments_url%%)
