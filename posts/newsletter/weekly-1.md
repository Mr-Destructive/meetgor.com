---
title: "Techstructive Weekly #1"
date: 2024-08-03
type: "newsletter"
---


# Week #1

This week, I had a ton of fun on the side as I finally made the MVP of a side project. An SSG with a Content Management System-like interface. an SSG with an editor that syncs up the posts from a database.

In this process, I learned about Cloudflare workers and golang. At work, I had some thinking and experimentation with LLMs.

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.

## Quote of the Week

> "The only limit to our realization of tomorrow is our doubts of today."
> 
> — Franklin D. Roosevelt

The LLM thing is affecting the mindset of people, more of which people think that it is going to replace humans, but really? Can a computer program that computes some numbers based on some other numbers even compare to human thoughts? That’s way far from what you are capable of, don’t waste the energy in thinking this.

## Wrote

Well, I didn’t write any articles this week, I am in the research phase for writing the Golang Patch Method, which is an interesting topic. It has way more things than I expected to be. This is what I love about writing articles/blogs; it helps me learn concepts in such detail.

Hopefully, next week, there will be an article to share in this newsletter.

## Read

* Pandas’ Dataframe is Column Major Data Structures
    

![](https://substack-post-media.s3.amazonaws.com/public/images/c5dc1fee-2d1e-4892-b219-4b96f6998ab5_288x288.png align="left")

[Daily Dose of Data Science](https://blog.dailydoseofds.com/p/what-happens-when-you-append-rows?utm_source=substack&utm_campaign=post_embed&utm_medium=web)

[What Happens When You Append Rows to a Pandas DataFrame](https://blog.dailydoseofds.com/p/what-happens-when-you-append-rows?utm_source=substack&utm_campaign=post_embed&utm_medium=web)

[Daily Dose of Data Science Free Book | Deep Dives…](https://blog.dailydoseofds.com/p/what-happens-when-you-append-rows?utm_source=substack&utm_campaign=post_embed&utm_medium=web)

[Read more](https://blog.dailydoseofds.com/p/what-happens-when-you-append-rows?utm_source=substack&utm_campaign=post_embed&utm_medium=web)

[5 months ago · 24 likes · 4 comments · Avi Chawla](https://blog.dailydoseofds.com/p/what-happens-when-you-append-rows?utm_source=substack&utm_campaign=post_embed&utm_medium=web)

* Deno and HTTP imports: [Article](https://deno.com/blog/http-imports)
    
* Pixel by Pixel, could create a vault of articles: [Pauline’s Blog](https://www.pawlean.com/vault)
    
* Row Scanning in Golang with reflect package: [Stackoverflow Question](https://stackoverflow.com/questions/56525471/how-to-use-rows-scan-of-gos-database-sql)
    

## Watched

* **Cloudflare Workers 101: A Crash Course on Cloudflare Workers**
    
* **Git Log**
    
* **How good are LLM? Think about it**
    

## Learnt

* **Cloudflare Workers Setup**: [Documentation](https://developers.cloudflare.com/workers/examples/)  
    It has good examples of some standard use cases and applications.
    
* **Getting table data using pymupdf**: [Article](https://artifex.com/blog/table-recognition-extraction-from-pdfs-pymupdf-python)  
    Getting some table position information of tables present in a pdf document. It can be really useful for the extraction of table-related data. For instance, I could get the width and height of the table to further use it in extraction techniques.
    
* **HTTP Patch method is not what you think it is**: [Article](https://imantumorang.com/posts/http-patch-method-ive-thought-the-wrong-way/)  
    I am also surprised that the PATCH method is different from than rest of the methods, we actually need to send commands/instructions to the server in order to make it understand what and how to update the resource-specific fields or entities.
    
* **Setting a Github Pages on GitHub Actions Workflow**:  
    I used this snippet in my [Tuxo SSG](https://github.com/Mr-Destructive/tuxo/blob/main/.github/workflows/cronjob.yml), which would serve as the output directory of the generated SSG files to GitHub Pages.
    
    ```go
        - name: GitHub Pages
          uses: crazy-max/ghaction-github-pages@v3
          with:
            target_branch: output-branch
            build_dir: my_app/
            jekyll: false
          env:
            GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
    ```
    

## Tech News

* [FastHTML](https://fastht.ml/): A Python framework for developing web applications. I would like to explore this in the upcoming weeks.
    
* [AI Studio](https://ai.meta.com/ai-studio/): Create fictional AI characters to chat with by Meta. This is pretty good use case for authors and fantasy writers. It could allow them to get to know their characters in a better way.
    

As always, I recommend going through the [HackerNews Newsletter](https://mailchi.mp/hackernewsletter/710) to get all the tech news.

That’s it from this week, hope you did well this week, and have a happy week and weekend ahead!

Thank you for reading, let’s catch up in the next week.

Happy Coding :)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
