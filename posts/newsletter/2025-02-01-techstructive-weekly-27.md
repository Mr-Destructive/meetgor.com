---
type: "newsletter"
title: "Techstructive Weekly #27"
date: 2025-02-01
---


## Week #27

It was a bit of lazy day, not that lazy though, did a lot of things while being lazy. I had been doing a lot of things in the past few weeks, so this week I decided to take things a little slowly. If I just keep on being that productive, it would be a burnout soon.

So, I created a new kind of video using a presentation-style approach, wrote a bit of python, and debugged it a ton, and now I see why tests are handy. I don’t have to debug the issue every time I make a change here, If run tests, I know exactly which test failed and where to look out for so easily.

I started an initiative to write something for an hour, I did just once a day, the other day’s life took over me.

### Quote of the week

> *"Almost everything will work again if you unplug it for a few minutes, including you."*
> 
> – Anne Lamott

Yes, I just took this week slowly. I took deep gazes and deep writing, want to have silent walks next week. Recharging the creative batteries before they go out. It is important to charge up and slow down simultaneously in my opinion, rather than waiting for yourself to drain away.

## Created

- [What is a programming language](https://youtu.be/pHt-Sri4ExQ?si=2UqaE4gd3lCmqkDj)?  
  I created this video as just a way to express my thoughts and the knowledge I have been gathering over the past few years. It might look silly, but it solidifies what I already about the thing I think I know.Double click to interact with video

## Read

- [How AI might change programming?](https://registerspill.thorstenball.com/p/how-might-ai-change-programming)THis is such a insightful post, so many things to excite yourself with, so many questions to satisfy your hunger of curiosity with.[![](https://substackcdn.com/image/fetch/$s_!-2HZ!,w_56,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2Ff114450f-2313-48c6-a253-a1d476c21d93_1164x1164.png)Register SpillHow might AI change programming?AI will change programming. I’m convinced of it now…Read more5 months ago · 21 likes · 9 comments · Thorsten Ball](https://registerspill.thorstenball.com/p/how-might-ai-change-programming?utm_source=substack&utm_campaign=post_embed&utm_medium=web)
- [How I am moving away from Google Ecosystem](https://itsfoss.com/leaving-google-ecosystem/?ref=dailydev): This is quite straightforward post, I see a lot of things missing like f-droid instead of Play Store, but it was quite helpful. I don’t use Mail that extensively and I think Gmail is not a bad thing. Others are good too, not sure of their privacy reasons.
- [GO’s wired little iterators](https://mcyoung.xyz/2024/12/16/rangefuncs/?ref=dailydev)   
  I have just skimmed through this post, haven’t really read it entirely but I think it will be important to go through.
- [A year of writing](https://organizingautomation.substack.com/p/a-year-of-writing?ref=dailydev)  
  This was a practical post.
    - Publishing requires courage
    - Fear can be overcome by practise
    - Hobby of creation will over do the activity of consumption[![](https://substackcdn.com/image/fetch/$s_!CpXu!,w_56,c_limit,f_auto,q_auto:good,fl_progressive:steep/https%3A%2F%2Fsubstack-post-media.s3.amazonaws.com%2Fpublic%2Fimages%2F8334a217-b96b-4626-8000-3c096654ec6a_144x144.png)Organizing AutomationA year of writing As the year comes to an end, it is a natural time to reflect on all of the things I have done. One new hobby I picked up in 2024 was the writing this blog. I’m taking some time to reflect, for my own understanding and also to help out others who might be considering the same…Read more6 months ago · 2 likes · Milan van Stiphout](https://organizingautomation.substack.com/p/a-year-of-writing?utm_source=substack&utm_campaign=post_embed&utm_medium=web)

## Watched

- [Exploring DeepSeek’s approach to LLM on Computerphile](https://www.youtube.com/watch?v=gY4Z-9QlZ64)  
  This was a great video explaining the key difference on how Deepseek did the LLM game differently, a concept called MoE mixture of experts. A network where the LLM will branch out to a specific network where it can use the weights more efficiently instead of the entire weights. Nice thinking, this feels so high-level view, how exciting or frustrating it would be to do that in a low-level and actually hands-on with the actual model.Double click to interact with video
- [Martin Fowler on Refactoring Podcast](https://youtu.be/lurbDAEU0KM?si=31-08t72M-DmbBM2)  
  Golden hour, this just changed my perspective on tests and that too at the right time. I had a issue to solve, or rather refactor a logic a little, but testing those changes was getting too tiring. I saw this and one thing stuck to me, is about tests. Testing would be the perfect thing to avoid debugging, it will not solve bugs, but it will avoid their creation in the first place. Valuable advice about everything, every minute or second of this video is pure gold.Double click to interact with video

## Learnt

- Filtering a data frame using bitwise operators. I wanted to filter a data frame elements with the condition, let’s a document data has entire page content, I only want the upper 10% and bottom 10%, I can do something like this using the bitwise operator
  ```
  df[ 
      ((df.y2 > 0) & (df.y2 < 0.1*height))
      |
      ((df.y2 > 0.9*height) & (df.y2 < height))
  ]
  ```
  Nice stuff.

## Tech News

- [DeepSeek](https://github.com/deepseek-ai/DeepSeek-R1), I think you are not on earth if you have not heard it this week.

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-731) and for software development/coding articles, join [daily.dev](http://daily.dev/) .

That’s it from this 27th edition of my weekly learning, I hope you enjoyed it, and leave comments on what you think about some of my takes or any feedback.
