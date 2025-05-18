---
type: "newsletter"
title: "Techstructive Weekly #19"
date: 2024-12-07
---

## Week #19

It was a pretty good week, with 2 videos published, 1 short created, and 4 live streams done, and 3 days of AOC solved. Phew!

On Saturday last week, I experimented with Netlify cloud functions in Golang and also for Turso’s LibSQL Database, and that went pretty well. Then started a livestream for the SSG in Golang, and completed the feed loading. Then on Sunday it was AOC Day 1, where I live-streamed it, created the video, and also scheduled the Netlify cloud functions video for Monday.

Then came Monday, derailed for a while but back to business on Tuesday somewhat. The week kept going and I kept solving problems for AoC in Golang and it gave me confidence to improve my problem-solving skills.

Pretty good week and I hope to make it even better next week.

### Quote of the week

> **"Do the best you can until you know better. Then when you know better, do better."**
> 
> — Maya Angelou

I have been creating videos almost consistently for the past 3 months and I have created 9 videos. I didn’t knew anything about video editing, I was slow at first, but I did the best I could. Now I know some tricks to get the editing done faster I do that to produce more videos and do it quickly. Start → Learn → Improve → Repeat.

## Created

* [Create and Deploy Netlify Serverless Functions in Golang](https://youtu.be/BY2Z2Em7OdA)
    
    On the weekend, I experimented with how to deploy a Golang serverless function for the LibSQL database, and through that, I got to understand the cloud function deployment on Netlify with Golang. So, created a video for that on Sunday and published it on Monday.
    
* Live streaming Advent of Code 2024 in Golang → [day1](https://www.youtube.com/live/3K02tEEBgto?si=m36J5UKzch1tjQ0X) | [day 2](https://www.youtube.com/live/4GwypzeIJAs?si=h5xt8bMeqDxVo19u) | [day 3](https://www.youtube.com/live/-rjLSk85M4Q?si=JYz1g7WEJ5dQsogo)
    
    After completing the basic SSG with Golang, on Sunday I live-streamed solving the day 1 puzzle for Advent of Code 2024 in Golang.
    
* [Advent of Code Day 1 in Golang video](https://youtu.be/4U97gLyz0Ss?si=KxGvQMnZjpONtPHS)
    
    After solving the problem for day 1 on live stream, created a short video explaining the problem and solution.
    

## Read

* [AOC in Golang Day 1: Missing ABS() for Integers](https://www.bytesizego.com/blog/aoc-day1-golang)
    
    This was a really well-researched and interesting article about why there is no Absolute function for integers. It makes sense now, but I still feel it could handle the general use cases so it shouldn’t be an issue. But yes there is a genuine reason for that to not exist in the standard library.
    
* [Command line tools I Like](https://rwblickhan.org/newsletters/command-line-tools-i-like-2022/): [rwblickhan.org](http://rwblickhan.org)  
    This was a nice post on the aggregated tools that the author uses most of the time and are quite handy. I also use most of the tools everyday and would like to create a dev container for the setup of my preferred editor/tools and shells and plugins and all that good stuff.
    
* [How I stopped worrying and learned to love Go Interfaces](https://dev.to/githaiga22/how-i-stopped-worrying-and-learned-to-love-go-interfaces-3m7p)
    
    This was well explained and to the point of why interfaces are confusing for most of the developers. They are really handy while working with APIs and I agree sometimes it could be annoying to fit the problems using interfaces, I have come to realize it is not that interfaces are bad, it’s just that we are not clear with the solution to the structure of the program that we are constructing.
    
* [The tech Utopia fantasy is over](https://blog.avas.space/tech-utopia-fantasy/)  
    I haven’t completed reading this post, but till what i have read it feels like we are moving away from the natural and individuality and focusing on the profits and metric games.
    
* [Avoid overusing the go init function](https://itnext.io/avoid-the-go-init-function-74f7f28e9154)  
    It is true, that we overuse certain things like this adding logic in the init functions that makes it hard to test and reproduce the behavior that is expected for a library or a package.
    

## Watched

* [Advent of Neovim](https://www.youtube.com/watch?v=TQn2hJeHQbM)  
    This is a good video explaining why he uses Neovim and how installing a program from a source helps you to appreciate and understand the project more.
    
    There is also a hidden gem for setting up an appname to a different name to isolate the versions or flavors of Neovim.
    

* [Configure Neovim’s options](https://www.youtube.com/watch?v=F1CQVXA5gf0)  
    This was another gem by Teej Devries for setting language/file type-specific configurations or options in Neovim.
    

* [Aaron Francis on The Software Huddle](https://youtu.be/Xdkwc26763M?si=kiI3dxdb1CobMwUs)
    
    He is a all round good guy. This podcast is really inspiring to learn and be consistent at it, he has gone from Accountant to launching his own company as a developer, speaker, educator, and database master.
    

## Learnt

* You can swap variables in Golang just like in Python
    

```go
a := 5
b := 10
a, b = b, a
```

This is pretty handy, I encountered this while trying to solve the advent of code problems. I thought can we do this in Golang as I have lazy habits from Python?

* [Slices.Delete](https://pkg.go.dev/slices#Delete) in Golang is Wired
    
    I mean just look at this
    

```go
s := []int{1,2,3}
slices.Delete(s, 1, 2)
fmt.Println(s)
// [1, 3, 3]
s = []int{1,2,3}
s = slices.Delete(s, 1, 2)
fmt.Println(s)
// [1,3]
```

I mean, If the function is returning the modified slice, why are we mutating the original one? Do one or the other, not both.

The safer route will be like this then:

```go
s = []int{1,2,3}
newS := slices.Delete(slices.Clone(s), 1, 2)
fmt.Println(s)
// [1, 2, 3]
fmt.Println(newS)
// [1,3]
```

I need to dive deep into why this is the way it is. Looks pretty confusing to mean.  
AFTER A WHILE:

[This article](https://medium.com/google-cloud/go-slices-deleting-items-and-memory-usage-81419317db3d) explains a bit clearly why the original slice is useless after the operation and we need to pass a copy of that in order to avoid it’s mutation.

## Tech News

* [Amazon’s New LLM Models](https://www.aboutamazon.com/news/aws/amazon-nova-artificial-intelligence-bedrock-aws) are the cheapest among the available ones
    
* [Open AI launches ChatGPT pro](https://openai.com/index/introducing-chatgpt-pro/) for $200 per month
    
* [Bluesky keeps growing in numbers](https://www.gsmarena.com/x_alternative_bluesky_reaches_24_million_users-news-65632.php)
    

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-724) and for software development/coding articles, join [daily.dev](http://daily.dev)

[Leave a comment](https://techstructively.substack.com/p/techstructive-weekly-19/comments)

[I ha](https://dly.to/LVQFgrjOUhf)ve been on bluesky as well, and I must say the engagement and the reach is much better compared to Twitter. I guess I’ll be posting more on Bluesky now.

You can connect with me on [meetgor.bsky.social](http://meetgor.bsky.social) on Bluesky.

Thank you for reading, let’s catch up in the next week.

Happy Coding :)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](https://techstructively.substack.com/p/techstructive-weekly-19?utm_source=substack&utm_medium=email&utm_content=share&action=share)
