---
type: "newsletter"
title: "Techstructive Weekly #15"
date: 2024-11-09
---

## Week #15

A wonderful week, bringing back the joy of programming, after some slumpy weeks, a refreshing and mind challenging week. The start was drowzy after a long weekend, but as the week progressed, I had to shift gears and complete the things that took some deep concentration and thinking, finally the results were worth the efforts. Everything in its place, the reward for the focus and hard work.

After a live stream and video creation in the end of week last week, I didn’t do much on Sunday, just took a break and spent the diwali removing the negativity in life. I am not mentioning it enough but I continue the writing routine, 35 days and counting, the writing streak is un-broken, I completed one section of writing of 21 days and now moving on to another section of 17 days, combining a total of (9+21+5(till date)) 35 articles. This is the motivation for me waking up everyday, writing will get me to the places I can dream of.

### Stats for the week

* FIxed a ton of bugs
    
* Brainstormed on refactoring code
    

### Plans for the next week

* Create 2 articles
    
* Record 1 videos
    
* Live Stream at least once
    

## Quote of the week

> "Strength doesn’t come from what you can do. It comes from overcoming the things you once thought you couldn’t."
> 
> — Rikki Rogers

I once thought before the week, that this week will be the make or break week for the things I have been doing for the past months, and finally in the end was able to break shakles and get into deep programming flow state, it felt good, it felt energetic, removing the guilt and imposter syndrome that has been creeping in.

Strength is not what you have, it is about learning what you don’t. It is about taking the challenge head on.

## Read

* [Writing for developers](https://rmoff.net/2023/07/19/blog-writing-for-developers/):  
    This is really a good read, writing brings clarity and clarity is what shapes you to pick up correct directions while developing anything. All the time, I start to write an article, I ran off a tangent on one or the other interesting thing that I had no idea about and assumed I knew it.
    
* [How to become a Good Backend Engineer](https://twitter.com/hnasr/status/1852537428227375482)
    
    This is really well summarised and insightful read for me at least, it gives a birds-eye view and ignites a curiosity to dive deeper into a specific topic or domain, which is essential for any developer.
    

## Watched

* [Building ChatGPT like LLM Model on a small scale from scratch](https://youtu.be/kCc8FmEb1nY?si=UV9hrQh2Uw8dc57Z):
    
    I have not completed it but just understanding the architecture and what the LLM does is vital to work and adapt to the changes in the evolving world.
    

## Learned

* [TikToken](https://github.com/openai/tiktoken) the tokeniser that powers GPT and its family of models: I want to explore this and create a project around it. I learnt that is the model that actually determines the number of tokens passed around the chats while interacting with the LLMs.
    
* Using the `next` method in python to get the first non-empty or truth value from a dict or any expression.
    
    ```go
    row = {
        "name": {
            "value": "123",
            "position": [40, 40],
        },
        "city": {
            "value": "mumbai",
            "position": [50, 30],
        }
    }
    row_position = next((attrs["position"] for attrs in row.values() if isinstance(attrs, dict) and attrs.get("value")), None)
    ```
    
* Combining Pandas Dataframes
    

Let’s say I have a list of dataframes, i want to combine them with certain number but not the contents, just append the next df to the end of the current df. I used the [pd.concat](https://pandas.pydata.org/docs/reference/api/pandas.concat.html) function to combine the slices of pandas dataframe in a single list.

```go
PAGE_LIMIT = 2
PAGES = 5
for i in range(0, PAGES, PAGE_LIMIT):
if i + PAGE_LIMIT <= num_of_pages:
df_batch = pd.concat(df_list[i:i + PAGE_LIMIT], axis=0)
else:
df_batch = pd.concat(df_list[i:], axis=0)
```

* Deleting elements from the list given the indices
    
    I know this could be easily gotten from GPT but feels good to do it yourself sometimes.
    

```go
def delete_elements_by_indices(lst, indices):
indices_set = set(indices)
return [item for idx, item in enumerate(lst) if idx not in indices_set]
lst = [10, 20, 30, 40, 50]
indices_to_delete = [1, 3]
result = delete_elements_by_indices(lst, indices_to_delete)
print(result)
Output: [10, 30, 50]
```

Fun tidbit,

Chat GPT search is just awesome, it recognises me? really?

![](https://substack-post-media.s3.amazonaws.com/public/images/ee8f7e63-6ee2-4257-823c-179855c99054_835x449.png align="left")

Even my youtube channel which is barely 2 months old

![](https://substack-post-media.s3.amazonaws.com/public/images/413b9325-6eff-4bda-bf79-eb862863e7a6_835x449.png align="left")

## Tech News

* [Hertz.dev](http://Hertz.dev) launches the conversational audio generator model
    
* [Stripe launches Workbench](https://stripe.com/blog/workbench-a-new-way-to-debug-monitor-and-grow-your-stripe-integration): A new way to debug, monitor, and grow your Stripe integration
    

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-721) and for daily developer articles, join [daily.dev](http://daily.dev)

That’s it from this week, hope you did well this week, and have a happy week and weekend ahead!

[Leave a comment](https://techstructively.substack.com/p/techstructive-weekly-15/comments)

Thank you for reading, let’s catch up in the next week.

Happy Coding :)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](https://techstructively.substack.com/p/techstructive-weekly-15?utm_source=substack&utm_medium=email&utm_content=share&action=share)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
