---
type: "newsletter"
title: "Techstructive Weekly #16"
date: 2024-11-16
---

## Week #16

This week was really productive, was able to make the code I have been writing for the past 2 months and experiments that have been doing into a deployable and ready to launch code. I finally feel a bit satisfied by my efforts, feel the motivation back, the curiosity back. I had an intense focus or flow state this week, I did not create any content this week as I was out on Sunday and was really tired on Saturday.

But this weekend, I am planning to create some 4 short tidbits of Golang videos. Also planning for a livestream, let’s see how that goes.

### Stats for the week

* Made one release with some minor changes and bug fixes
    

### Plans for the next week

* Tidbit Video for Golang 1.24 release features
    
* Init the SSG with DB project live stream
    

## Quote of the week

> "The key to success is not in what you do once, but in what you do repeatedly."

This is the vibe of the weeks, I have been doing things repeatedly over the span of 6 weeks. I have kept up with my writing routine, some days are rough but some days I am tough. It is important to go through the bad days in order to get the fulfilment and bring the best of us.

## Read

* [How Cosiine Similarity Works](https://tomhazledine.com/cosine-similarity/?ref=dailydev): This is a good read for understanding the cosine or vector similarity. This is really important to understand in the context of today’s fast-paced AI and LLM world, understanding the tokens and the embedding similarity between them. It might not seem oblivious but having an understanding of the essence of the logic behind the things happening around you makes life a lot easier.
    
* [INodes in Linux and FIlesystems](https://www.redhat.com/en/blog/inodes-linux-filesystem): I found this while researching for INODE busting from the video of the Git course walkthrough from Primeagen on [Boot.dev](http://Boot.dev)
    
* [Avoid Select \* in SQL](https://x.com/hnasr/status/1856745402399359315): This is really helpful, always need to be explicit instead of being lazy or over-cautious in case would need that field in the future. There are two things here, one is being explicit and the other is dynamic for future changes, just in case the other fields might be required later.  
    

## Watched

* [Python 3.13 Replace Function](https://youtu.be/H2G_BsF6HT4?si=RM0qPLwvu_UYqcwM)  
    This is really neat, I see myself using this to mutate objects in tight data-oriented programs or applications.
    
* [Turso LibSQL Video from Fireship](https://youtu.be/PGpL5hYpY1o?si=yPfAOsB9l8DhuVAt)
    
    Finally Fireship creating a video on Turso’ s LibSQL, this deserved attention and its finally getting the love it deserves.
    

* [Have you heard of Orion? It is worse than Gippty!](https://www.youtube.com/watch?v=ZehQ4XQs9NA)
    
    This is such a funny and satisfying video, it brings such relief to hear each word. programmers are not replaceable because we write bad code that works :)
    

## Learned

* [What is INode (Index Nodes) used in Linux file systems](https://www.redhat.com/en/blog/inodes-linux-filesystem): The was really interesting piece of concept in Linux, if you have been using Linux for a while and have kept a directory with loads of files, then something wired happens, not sure what but something off happens, I have experienced it. And this seems to be the reason for it, pretty neat of git using this to store commit hashes in this way.
    
* How to iterate over the df rows and mutate the values: We can use the iterrows to iterate over the rows and use the `at` with the index and the column name to mutate the field by setting it to the required value.
    

```go
import pandas as pd
df = pd.DataFrame({
'A': [1, 2, 3],
'B': [4, 5, 6]
})
for index, row in df.iterrows():
df.at[index, 'A'] = row['A'] * 2
print(df)
A  B
0  2  4
1  4  5
2  6  6
```

## Tech News

* [Can we have a moment of silence for the ones who thought Programmers were replaceable](https://the-decoder.com/openais-new-orion-model-reportedly-shows-small-gains-over-gpt-4/)?
    

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-722) and for daily developer articles, join [daily.dev](http://daily.dev)

That’s it from this week, hope you did well this week, and have a happy week and weekend ahead!

[Leave a comment](https://techstructively.substack.com/p/techstructive-weekly-16/comments)

Thank you for reading, let’s catch up in the next week.

Happy Coding :)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](https://techstructively.substack.com/p/techstructive-weekly-16?utm_source=substack&utm_medium=email&utm_content=share&action=share)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
