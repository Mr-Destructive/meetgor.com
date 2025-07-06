---
type: "newsletter"
title: "Techstructive Weekly #23"
date: 2025-01-04
---


We are in 2025! Happy New Year

It feels good to carry this newsletter through the year from July 2024 to January 2025 and counting ….

## Week #23

It was a huge roaring start to the year, I created a video on the 1st, and that too a big video, I put in everything I could know about the thing I wanted to talk about in little as I could, and it's gaining views and is likely to spike as the Golang 1.24 release date approaches. 200 IQ move there.

I also had a few moments of enjoyment and satisfaction with my work, I was delivering the metrics report for the quarter and felt really good about the things I have done, still not good but way better than where it was.

I spent my Saturday and Sunday on live stream and understanding the weak package and weak pointers. I again spent half of the 1st of January on preparing the material for it and finally released the video with a sense of accomplishment. Couldn’t have asked for a better start to the year.

### Quote of the week

> “Learning and failures will always remain,
> 
> Guiding you through both joy and pain”
> 
> — Meet Rajesh Gor

This week, I failed and tried again. I failed again, but I kept trying and finally, I understood the weak package. It’s a reminder that this principle holds true in life as well.

At times, I might have been in the wrong mindset while learning, but I ended up discovering more than I expected. The satisfaction I feel now is beyond words.

Learning stays with you for a lifetime, shaping who you are. Failures, on the other hand, often come with taunts and criticism when things aren't going your way. But in those moments, it’s crucial to let your learning guide you to turn those failures into stepping stones for growth.

## Created

- [What’s new in Golang 1.24: Weak Package](https://youtu.be/ehXaekF9lD4?si=MJ2rVohdvkkRyF9N)  
  An introduction to the weak package in the standard library, I explored the difference between strong and weak references and some garbage collector-related quirks in Golang. The usage of weak pointers and some examples in different cases.Double click to interact with video
- [2024 Tech Retrospective/Review Article](https://www.meetgor.com/2024-review/):
  This has been my ritual since 2022, I love reflecting on my successes, failures, and just things that I did in the past year. It helps me realize how far I have come and motivates me to keep going ahead. About the actual review, 2024 was a bittersweet year for me, the first half kept me dangling by the thread but the latter half made me fly in the sky. This was my lesson that when no one believes in you, you still have to belive in yourself, you have to have hope, you need to have the courage to keep going ahead and not care too much about the result. Hope you read and [enjoy my rambling](https://www.meetgor.com/2024-review/) :)
- Live streamed about Static site generator and making plugin-based system in Golang.   
  [Saturday’s Stream](https://www.youtube.com/live/SdpvQk3f2Yg?si=XYZkrQHlp-TRwh-i)  
  [Sunday’s Stream](https://www.youtube.com/live/B-jHx1hLtv0?si=Tvf7Fzca0uXGyZVI)Double click to interact with videoDouble click to interact with video
  I took a break from Advent of Code problems, it was getting a bit tough to manage time and resist completing this project, so I abandoned this Aoc 2024 for the moment and picked up this project back. I converted the basic SSG into a pluggable component-based SSG. This weekend I will be completing the Database integration part.

## Read

- [Are Pointers faster than Values](https://blog.boot.dev/golang/pointers-faster-than-values/?ref=dailydev):   
  Hell no, that is not faster, it is just that stack is faster in that case, but the moment you make you point the gun to yourself, you will shoot on the foot.
- [Docker Networking Basics](https://adventofdocker.com/posts/day-10-docker-networking/):  
  Learning about docker is never enough, I learned about the -P which will forward all ports from the container to the host machine at random. That is wired but a good quirk to know, you never know when that might be handy.
- Peoples 2024 review posts
    - [Aron Francis’ 2024 year in review](https://aaronfrancis.com/2024/year-in-review-2024-ebfbb78c)  
      This is the man I admire, and wow he has done quite a bit to make quite an impact on developers, the courses are soo valuable, the articles are soo motivating, as humble as he gets.
    - [Haimantika Mitras: Wrapping up 2024](https://newsletter.haimantika.com/p/wrapping-up-2024)
      It’s a bit hard to face situations like hers in the festive month. It feels a bit disturbing and yet she came out with flying colors and kept the enthusiasm on, a lot of inspiration to draw from.
    - [Florencia Luz Duarte’s: Goodbye 2024](https://unicornio.dev/en/blog/bye-2024/)  
      THe story felt a bit similar to mine, landing a full-time job after a hunt. It feels good to be in those shoes. It might be hard, but I am emphatic about it, as I know the feelings of rejection after you worked hard. Also the nice little tidbits like having 24 hours in a day, we can do only a few things, needn’t do all the things is a reminder to ease out and take things one by one.

## Watched

- [Don’t serve static files from your backend serve, use serverless storage cloud providers](https://youtu.be/aybSXT9ZJ8w?si=t_laTOvNBqDPv11E)Double click to interact with video
- [It’s time for a change for The Primeagen](https://youtu.be/BiZ1CLT3nEM?si=_UU5WfX62a23cnWI)
  He quit Netflix and now feels a bit of change in his routine. It is a heck of a commitment to stream for 6 years almost every single day. That is enthusiasm, that is contagious curiosity. I have a huge respect for him.Double click to interact with video

## Learnt

- Understood a lot about Weak pointers and Garbage collection in Golang
    - Weak pointers allow to observe a value/object without stopping the garbage collector from collecting or cleaning up that object.
    - The garbage collector has a graph like structure that keeps track of the states of the variables in the programs in the mark phase, in order to observe and clean them in the sweep phase.
    - Weak pointers use unique pointers which are the direct references to all the objects used in the program. The weak pointer internally points to this unique pointer which in turn points to the actual object, so when the garbage collector sees that the actual object has no direct reference it cleans that memory where the object is stored making the unique pointer nil. So, when the weak pointer is accessed to check the value, it will be nil in a safe way.
- [How to send messages on a slack channel with the webhooks](https://api.slack.com/messaging/webhooks)  
  This was one thing I studied for implementing or updating a few things at my work. I was excited to get the data out and by doing this I was satisfied by looking at my work, the metrics shine in the alerts channel.

## Tech News

- [70% of go developers (who filled the survey) use LLM while developing with Go](https://www.infoworld.com/article/3630940/go-teams-struggle-with-coding-standards-survey.html)
  That is an insane number, are you saying to me every 7 people out of 10 use LLM while working with Go? That is understandable considering how repeatable Golang gets sometimes, but the adoption rate is bonkers.

Not much news, I mean who will release a product on 31st December, developers are partying, right?

[![8 March | image tagged in new year computer guy,international women's day,8 march,memes | made w/ Imgflip meme maker](https://substackcdn.com/image/fetch/$s_!TAY-!,w_1456,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Fee72ed20-f923-47b8-8e34-4af668c2a2d6_150x150.jpeg)](https://substackcdn.com/image/fetch/$s_!TAY-!,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Fee72ed20-f923-47b8-8e34-4af668c2a2d6_150x150.jpeg)Yes, they party a lot!

Anyways:)

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-727) and for software development/coding articles, join [daily.dev](http://daily.dev/) .

That’s it from this 23rd edition of my weekly learning, hope you enjoyed it, and leave comments on what you think about some of my takes or any feedback.

[Leave a comment](%%half_magic_comments_url%%)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](%%share_url%%)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
