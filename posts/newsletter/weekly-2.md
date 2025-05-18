---
type: "newsletter"
title: "Techstructive Weekly #2"
date: 2024-08-10
---

## Week #2

This week I made some mistakes at my job (well they were the week before) but they got caught after the review. I learned how powerful reviews could be at catching flaws in code. I also caught a bug while pair programming, it was a mixed feeling as a developer for me this week.

In the end, I would recommend debugging your code thrice, yes not once, not twice but thrice, and also make sure to forcefully make the code flow reach the section of your changes, if that is a bit of an edge case thing.

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.

On the side, I have been playing with the Meta AI API (not the official API) to use it anonymously without logging in using Golang.

## Quote of the week

> “We all make mistakes, but it's how we come back from the mistakes that matters.”
> 
> \- Unknown

As I said, I made mistakes, but those were guiding principles and learning experiences to grow with. You can cry and give reasons, or accept your mistake and improve. Choose the latter, it will help you grow.

## Read

* [How broken is hiring in 2024?](https://cropp.blog/2024/08/job-searching-in-2024-is-horribly-broken?ref=dailydev)  
    Being a fresher in 2024, I have experienced it quite closely from the glimpse of the beginning of this phase after Q3 in 2023, where I applied almost 100 applications without a single interview—finally landed an interview and working in that company today. What a ride it has been, 100+ applications, 1 interview to crack the job, skills issues? maybe but I think it could be timing issues.
    
* [How Canva Saved $ 3M manually in Amazon S3 Costs with Glacier Instant Retrieval](https://aws.amazon.com/blogs/storage/how-canva-saves-over-3-million-annually-in-amazon-s3-costs/)  
    This is quite a smart solution, I would call, it the right tool for the right job, backed by data. They rightly used the right analytical metric to solve their problem.
    
* [How to build Extremely quickly](https://learnhowtolearn.org/how-to-build-extremely-quickly/?ref=dailydev)  
    Outlining is important in making anything, it gives me the boost to complete the unfinished project, as well as the momentum to carry on after a break or the next day.
    

## Watched

* Git Cherry Pick
    
    I haven’t actually used this feature of git, but in my previous internship was used by my seniors while making hotfixes in releases. Now I got what actually it was.
    
* Git Merge
    
* Backend Banter (Lane Wagner) with Lawrence Lockhart
    
    This is quite a good listen, it elaborates on juniors trying to land their first job, how you should network, and how different is learning for everyone. Must watch for beginners and freshers to be placed in a good mindset for the future as well as be helpful to your future junior colleagues.
    

## Learnt

* Structuring a [function call output](https://platform.openai.com/docs/guides/function-calling) in an Open AI API Call: I Learnt a few things like adding a structured output to a function call in an Open AI API call.
    
* Reading and Writing a PDF file with [pikepdf](https://pikepdf.readthedocs.io/en/latest/) (useful for splitting and manipulating files)
    
    ```go
    import pikepdf
    from pathlib import Path
    def split_pdf(input_pdf_path, output_dir):
    pdf = pikepdf.Pdf.open(input_pdf_path)
    output_dir = Path(output_dir)
    output_dir.mkdir(exist_ok=True)
    for i, page in enumerate(pdf.pages):
    output_pdf_path = output_dir / f'page_{i+1}.pdf'
    with pikepdf.Pdf.new() as output_pdf:
    output_pdf.pages.append(page)
    output_pdf.save(output_pdf_path)
    ```
    
* LLMs are taking over. I mean not quality-wise, but they seem to be everywhere, almost every company is trying to use AI Agents to make their business look smart (but is not actually). This is a harsh or soft truth that I have to accept and move ahead to leverage in things I would do and build.
    

## Tech News

* [Structured Response](https://platform.openai.com/docs/guides/structured-outputs) from Function Calls in Open AI API Calls
    
* [Appwrite Init](https://appwrite.io/init): Some special releases from the Appwrite Team in the upcoming weeks. Make sure to grab your ticket. Golang's support for cloud functions is confirmed!
    
* [Google LLM Gemini API Cost Cut by around 70%](https://developers.googleblog.com/en/gemini-15-flash-updates-google-ai-studio-gemini-api/#:~:text=Gemini%201.5%20Flash%20price%20decrease&text=To%20make%20this%20model%20even,tier%20as%20well%20as%20caching\).)
    

Follow [HackerNewsletter](https://mailchi.mp/hackernewsletter/711?e=ed0f2c4e4f) for more updates

That’s it from this week, I hope you did well this week, and have a happy week and weekend ahead!

Thank you for reading, let’s catch up in the next week.

Happy Coding :)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
