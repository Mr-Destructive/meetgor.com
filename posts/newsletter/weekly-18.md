---
type: "newsletter"
title: "Techstructive Weekly #18"
date: 2024-11-30
---

## Week #18

It was a great week, if I could call it a hard-working week, it would be appropriate. I started the week (the weekends) by streaming, creating a video, and planning a few more videos for the next week. Then starting the week with deep intense programming sessions, a couple of flow states during the week made the mind feel accomplished and excited.

After releasing the code that I had been working on for so long in the past week, this week was about taking a new approach to the initial implementation, and it was smooth, it felt good, I was able to change the code that I had written faster and iterate on it quickly. This gave me the confidence that I can write production-ready and readable code. I am not that bad, there is a lot to improve on but quite a good feeling to live with and continue the work.

### Quote of the week

> “The only way to do great work is to love what you do”  
> — Steve Jobs

This week, I did the things that I loved, working with SSGs, and databases, writing code, and teaching what I know. It might not be great, but these are the things that I love, and other is no one forcing me to do this, this comes out of my will.

## Created

* VIDEO: [What’s new in Golang 1.24: Omitzero Tag in Marshalling JSON](https://www.youtube.com/watch?v=RFUz4_axOZg)  
    This video was about the specific tag which is being added to the encoding/json package for Marshalling from Golang structs to JSON (serialization.)
    

* [TIL: Gotcha with Chained Assignment in Python](https://www.meetgor.com/python-chain-assignment-gotcha/)  
    I learned this point while implementing a feature at work, I was in a deep flow state, but this thing bugged me and broke out of the state rightly so, otherwise would have gone in wired directions without doing anything significant.  
    Learning python might be simple, but mastering it is a task, there are soo many things to learn about and be aware about. Python is simple but at times could be dangerous if not careful.
    

![](https://substack-post-media.s3.amazonaws.com/public/images/9786b965-3564-4542-8687-a55727236ab5_325x325.png align="left")

[Meet's Substack](https://meetgor.substack.com/p/til-gotcha-with-chained-assignment?utm_source=substack&utm_campaign=post_embed&utm_medium=web)

[TIL: Gotcha with Chained Assignment in Python](https://meetgor.substack.com/p/til-gotcha-with-chained-assignment?utm_source=substack&utm_campaign=post_embed&utm_medium=web)

[I was writing some Python code and wanted to initialize a few variables to an empty list. Instead of creating separate lists for each variable, I decided to use chained assignments like this…](https://meetgor.substack.com/p/til-gotcha-with-chained-assignment?utm_source=substack&utm_campaign=post_embed&utm_medium=web)

[Read more](https://meetgor.substack.com/p/til-gotcha-with-chained-assignment?utm_source=substack&utm_campaign=post_embed&utm_medium=web)

[19 days ago · Meet](https://meetgor.substack.com/p/til-gotcha-with-chained-assignment?utm_source=substack&utm_campaign=post_embed&utm_medium=web)

* STREAM: SSG from Scratch in Golang (2 streams)  
    I streamed on the weekend and started where I left off by creating static site generator in Golang. This part was about loading the config and posts from the folders and files.  
    Then on Sunday, I continued the stream for completing the most basic SSG in golang by adding the rendering of the templates.
    
    * [Loading Config and Posts](https://www.youtube.com/watch?v=smWEhhypbK4)
        

* * [Loading and Rendering Templates](https://www.youtube.com/watch?v=p_av6C8Lji0)
        

## Read

* [Database Indexing Explained](https://computersciencesimplified.substack.com/p/database-indexing-explained): from Computer Science Simplified Substack  
    This was a great post, explaining the necessary context and fundamentals required to understand the indexing. The last part was a bit heavy, but overall a good and in-depth article on learning indexing.
    
* [From laid off to hired: Software engineering guide](https://newsletter.eng-leadership.com/p/from-laid-off-to-hired-a-software)  
    This is a good motivational post on getting hired in the era of 2024, this is a new kind of world where you are not only hired based on what you say, but also what your impact says, your brand, your online presence, and actual skills.
    
* [How I increased my visibility: Kent C. Dodds](https://kentcdodds.com/blog/how-i-increased-my-visibility)  
    To increase visibility, you need to put it out there. To be seen by others, you have to show your skills and struggles in public. This is to a point article showing practical experience in his journey.
    
* [Wafris: Rearchitecting Redis to SQLite](https://wafris.org/blog/rearchitecting-for-sqlite)  
    I watched a video on the Database School where a SaaS (Firewall for Web applications) provider moved from Redis to Sqlite. This is such a nice technical post, there are so many things to dive into and how SQLite is the database that just works and not just works but also shines.
    
* [Time spent programming is often time well spent](https://stanbright.com/time-spent-programming/)
    
    > Programming is often more fun than the alternative uses of my time.
    
    Programming is really fun, the joy of creating something and the satisfaction of completing or figuring out by yourself hits different and cannot be replaced with anything else.
    

## Watched

* [Moving from Redis to SQLite with Mike Buckbee](https://www.youtube.com/watch?v=EwDuYId5v8k): Aron Francis in the Database School  
    As mentioned in the above post, this video was absolutely fun to watch, so much technical stuff in a digestible way. There is so much to learn about networking, the quirks of Redis, SQLite and its adaptability, IP Addresses, and interesting problems to solve. Totally worth the hour.
    
* [We built an orchestrator from scratch](https://www.youtube.com/watch?v=dy2RJdDEvO0): [Fly.io](http://Fly.io)  
    Why Flyio built their orchestrator, kind of. They use VMs which Kubernetes is not ideal for as it is designed for orchestrating containers, Nomad has a quirk of assigning a minimal number of VMs/machines for the users, which is not secure enough for [Fly.io](http://Fly.io), so reasonable enough that it is worth for them to write and Orchestrator from scratch (well not entirely from scratch)
    
* [Why would anyone use Functional programming](https://www.youtube.com/watch?v=rYR0eJdGBmQ): The Coding Gopher  
    Functional programming languages are designed to think about the what and not the how. So we just focus on the what things to implement and not care about how. So, there are things like functions, no variables no mutations (get out of here python).
    
* [Let’s Prototype a Javascript JIT Compiler](https://www.youtube.com/watch?v=8mxubNQC5O8): Andreas Kling  
    To be honest, I didn’t get most of the technical things, however, I get the point, he is using AI wisely, whenever he uses AI, HE IS IN CONTROL and he doesn’t let AI take over. I have not completed this video, I said last week it would be on my watched list, but I got lost in other stuff and lost interest in the livestream.
    

## Learned

* Difference between omit empty and omit zero while marshaling JSON bytes in Golang  
    From the documentation
    
    > The "omitempty" option specifies that the field should be omitted from the encoding if the field has an empty value, defined as false, 0, a nil pointer, a nil interface value, and any empty array, slice, map, or string.
    
      
    The omit empty is wired, it has its definition of what is empty (which I don’t disagree with but it could be confusing at times). With omit zero, we can control what a zero value is for certain specific data types like structs and time values.  
    Will write a post on it soon.
    
* How to render templates in Golang  
    While doing the live stream and building the static site generator from scratch, I explored the difference of Parse and ParseFS which are two different ways of rendering templates.  
    
    * Template
        
        ```go
        // templates/hello.html
        <h1>Hello, {{.Name}}</h1>
        ```
        
    * ParseFS
        
        ```go
        package main
        import (
        "embed"
        "html/template"
        "log"
        "os"
        )
        // Embed the templates directory
        //go:embed templates/*
        var templatesFS embed.FS
        func main() {
        tmpl, _ := template.ParseFS(templatesFS, "templates/hello.html")
        data := map[string]string{"Name": "Meetgor.com"}
        _ = tmpl.Execute(os.Stdout, data)
        }
        ```
        
    * ParseFiles
        

```go
package main
import (
"html/template"
"log"
"os"
)
func main() {
tmpl, _ := template.ParseFiles("templates/hello.html")
data := map[string]string{"Name": "Meetgor.com"}
_ = tmpl.Execute(os.Stdout, data)
}
```

* `ParseFiles` reads templates from the file system, while `ParseFS` reading templates from an embedded file system.
    
* `ParseFiles` requires physical files on disk, whereas `ParseFS` uses files embedded in the Go binary.
    
* `ParseFiles` allows template modifications without recompiling, while `ParseFS` requiring recompiling to update templates.
    

* Chained assignment references the same value to all the assigned variables in Python  
    Python might be easier to learn as said earlier, but mastering it really hard, I leart that chaining assignment is not ideal when you want to iniitialise individual variables to different values (same values but separate on their own moving ahead in the code)
    
    ```go
    a = b = c = [1,2,3]
    b.append(4)
    a = [1,2,3.4]
    b = [1,2,3,4]
    c = [1,2,3,4]
    ```
    

## Tech News

* [Alibaba Releases QwQ: An Open Source Model as Competitor to OpenAI’s O1 Reasoning Models](https://techcrunch.com/2024/11/27/alibaba-releases-an-open-challenger-to-openais-o1-reasoning-model/)  
    Every week, some random thing is happening in the world of AI, someone beats the other buzzing model by little margins. Exciting times in 2025.
    
* [Spotify cuts developer access to some developer recommendation APIs](https://techcrunch.com/2024/11/27/spotify-cuts-developer-access-to-several-of-its-recommendation-features/)  
    This is bad, and looks like the possible issues moving ahead in this AI driven internet.
    
* [BlueSky Open API means anyone can scrape your data for AI training](https://techcrunch.com/2024/11/27/blueskys-open-api-means-anyone-can-scrape-your-data-for-ai-training/)  
    Everything will have side effects, it depends on how the people are actually using it to their own and others’ advantage.
    

For more news, follow the [Hackernewsletter](https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-723) and for daily developer articles, join [daily.dev](http://daily.dev)

## Conclusion

This week was all about refinement and iteration, and it felt great to see the progress I made. From streaming and content creation to diving into deep coding challenges, I learned a lot about the intricacies of Golang and Python.

It is Advent of Code from next week, so I am planning to stream daily for solving AoC with Golang in the first iteration and OCAML in the second, let’s see how it goes.

If you want to tune in, slide into Twitch or youtube

* [https://www.twitch.tv/meet\_gor](https://www.twitch.tv/meet_gor)
    

* [https://www.youtube.com/@meet-technically](https://www.youtube.com/@meet-technically)
    

That’s it from this week, hope you did well this week, and have a happy week and weekend ahead!

[Leave a comment](https://techstructively.substack.com/p/techstructive-weekly-17/comments)

Thank you for reading, let’s catch up in the next week.

Happy Coding :)

Thanks for reading Techstructive Weekly! This post is public so feel free to share it.

[Share](https://techstructively.substack.com/p/techstructive-weekly-17?utm_source=substack&utm_medium=email&utm_content=share&action=share&token=eyJ1c2VyX2lkIjo5MDE1NzgwMywicG9zdF9pZCI6MTUyMDI0Njk0LCJpYXQiOjE3MzI4OTY2NTYsImV4cCI6MTczNTQ4ODY1NiwiaXNzIjoicHViLTI4MjQwMzciLCJzdWIiOiJwb3N0LXJlYWN0aW9uIn0.rtsuau2_H8vZZwASSY9gzTiaAyWxWXK4YrcvET-TcZE)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
