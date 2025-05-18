---
title: "Techstructive Weekly #10"
type: "newsletter"
date: 2024-10-05
---

## Week #10

Wrapping up the third quarter with a burst of energy! This week felt more mentally demanding, but as the week closed, I realized that even small progress is still progress. Sometimes, numbers don’t tell the whole story—they can shift perceptions or even mislead. I believe in valuing the journey more than the destination, and this week’s coding, learning, and creating was worth more than any metric can show.

This week, I did some programming, not much, but the weekend was super productive. I wrote an article, created a video, and also learned a lot of things about golang, docker, and SQLite.

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.

### Stats for the week

* Articles written: 1
    
* Videos Created: 1
    
* Concepts learned: 2
    
* Code Lines Written: ~500
    
* Bugs Caught: 4
    
* Articles Read: 3
    
* Videos Watched: 3
    
* Books Pages Read: 37 (fiction+non-fiction)
    

### Plans for the next week

* Creating a video about SQLC library with LibSQL database
    
* Crossposting and cleaning up the article on remote LibSQL database connection in Golang
    
* Contribute to one of the open source projects: Steampipe, MindsDB, Turso, Appwrite
    

## Quote of the week

> **“Not everything that can be counted counts, and not everything that counts can be counted.”**  
> — William Bruce Cameron

This sums up my week really well. I worked hard to get things up and running, making tweaks, and shifting directions. But in the end, it boiled down to numbers. For me, though, the process was more important than the final metrics. The process helped me reshape my mindset, shift my perspective, and adapt to different conditions as the problem evolved. I suppose I learned to embrace change and gain more insights than the numbers suggest.

## Read

* [Interfaces in Golang: ByteSizeGo](https://www.bytesizego.com/blog/golang-interfaces): Interfaces are really important in making certain applications, but more important is how to use them effectively and whether to use them or not in your specific problems.
    
* [NotebookLM has got everyone in a daze about LLMs and podcasts](https://simonwillison.net/2024/Sep/29/notebooklm-audio-overview/): I just dropped my portfolio link and it just spits out some compliments after compliments about my journey as a developer. It actually said certain things that I myself haven’t thought of someone viewing me, a couple of things were rubbish for sure, but it was pretty much factual.
    

## Wrote

* [Connecting a remote LibSQL database hosted on Turso in Golang](https://www.meetgor.com/turso-libsql-db-golang/)  
    This is a write-up for my first-ever YouTube video which after 3 weeks has got around 200 views which is fascinating. Hmm, numbers are wired. After creating a video, the process of writing the article seems so smooth and easy as all the hard work is done, it’s just about to flush out the idea into words.
    

## Created

* Setup LibSQL Database server, SQLD locally using Docker and Connecting with Golang:
    
    I also went ahead and created my 3rd video in the series “Let’s Go with Turso” covering how to setup a LibSQL database which is called SQLD locally using Docker and then connecting and querying with a Golang program.
    

## Watched

* Tips for handling timezones in Postgres:  
    It is quite weird and frustrating to manage timezones as a developer it seems.
    

I have watched very less YouTube for the past two weeks, it is the effect of creating content, consuming less, and creating more.

## Learnt

* Doker: Mounting volumes with the directory in the system and point it to a location in the container for persistence
    
    To create a mounted volume, you need to specify the host directory path and the directory path inside the container.
    
    ```go
    docker run -v /path/on/host:/path/in/container image-name
    ```
    
    I used it to mount sqlite db from the local file path to the container of sqld to mount the default db used to serve the libsql database like this:
    
    ```go
    docker run -p 8080:8080 -ti \
    -v $(pwd)/data.sqld/:/var/lib/sqld/data/ \
    ghcr.io/tursodatabase/libsql-server:latest
    ```
    
* I am learning about memory management with the [course on](https://www.boot.dev/courses/learn-memory-management) [boot.dev](http://boot.dev)[:](https://www.boot.dev/courses/learn-memory-management) This course has been made by [TJ devries](https://www.boot.dev/teachers/tj-devries), If you don’t know him, you are missing out on some serious Neovim and Ocaml Shenanigans on Twitter and Twitch. I am not gonna miss this masterpiece while it’s fresh, I know the basics of C, and I should be able to learn the more advanced nuances of memory management.
    

## Tech News

* NotebookLM is the buzz, people are creating actual podcasts from certain pieces of aggregated content, pretty impressive to be honest
    
    [https://notebooklm.google/](https://notebooklm.google/)
    
* [Meta AI launches Movie Gen model](https://ai.meta.com/blog/movie-gen-media-foundation-models-generative-ai-video/) that generates realistic audio and video. AI trend seems to be never-ending, each week there is one or the other model being released, taking over the other by the barest of margin but looking like a huge growth.
    

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-718) (#718 has not yet been published but will soon as usual), and for daily developer articles, join [daily.dev](http://daily.dev)

That’s it from this week, hope you did well this week, and have a happy week and weekend ahead!

Thank you for reading, let’s catch up in the next week.

Happy Coding :)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
