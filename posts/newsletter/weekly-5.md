---
type: "newsletter"
title: "Techstructive Weekly #5"
date: 2024-08-31
---

## Week #5

This week felt like a slog, with many challenges and frustrations. However, by the end of the week, I found my stride and got excited about the direction I’m heading. It's amazing how quickly things can shift from feeling like the end of the world to experiencing a burst of excitement.

I also managed to get the side project of the [Meta AI API wrapper](https://github.com/Mr-Destructive/meta-ai-golang) in Golang correctly, over the weekend will polish this project and also fix the bug in the Chat GPT anonymous client in Python.

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.

## Quote of the week

> **"Success is not final, failure is not fatal: It is the courage to continue that counts."**  
> — Winston Churchill

As a developer, you will experience both triumphs and challenges. Your journey might feel like a rollercoaster, swinging from highs to lows in a single day. Embrace the mantra: Build, Iterate, Release. There is no ultimate success or failure, and no product is ever perfect. What matters is the continuous improvement and learning you gain along the way.

## Read

* [Don't be an Alpha Geek - John Crickett](https://open.substack.com/pub/developingskills/p/dont-be-an-alpha-geek?utm_source=share&utm_medium=android&r=1hoe7f)  
    Just be empathetic and thoughtful about your actions and feedback. This will gradually creep into your all issues and make you a better developer, this is great advice.
    
* [Review your own PRs](https://sophiabits.com/blog/review-your-own-prs)
    
    This is good advice, I do it as a ritual on GitHub, 8 out of 10 times, I get the feedback myself. The editor myth is real, there are things that you don’t notice in your editor, in your environment, in your flow. But as soon as the context changes, the words, and the logic seem to be distant. Believe this advice, this saves a ton of time.
    
* [Lessons learned in 35 years of making Software](https://dev.jimgrey.net/2024/07/03/lessons-learned-in-35-years-of-making-software/): I have barely lived half of 35 years, the sheer amount of experience in 35 years of software is immense respect. But what is shared here, the words disheartened me that your code will eventually be deleted, will be replaced, will be evolved. This is a harsh reality but we need to accept and move ahead in life.
    

## Watched

* Recursion with [Boot.Dev](http://Boot.Dev):
    
    This has been well explained and visualized as well. The key part in understanding recursion for beginners is visualizing the call stack and going through the code step by step. Must watch for begineers.
    
* Arden Labs Podcast: Guest Samantha Coyle with host Bill Kennedy  
    

This had some good insights and provided some guidance on how to get internships through networking and building the way up the ladder as a software developer.

* Rust Command Line Arguemnts by Francesco Cuila:
    
    I watched this on the Live stream on Saturday, was a nice chilling stream with learning a thing or two in Rust.
    

## Learned

* Difference between cookies Add and Set in URL Values:
    
    [Add](https://pkg.go.dev/net/url#Values.Add): Appends the value to the key without replacing existing values (useful for handling multiple values for a single key).
    
    [Set](https://pkg.go.dev/net/url#Values.Set): Replaces the existing value for the key (ensures that only one value is associated with the key).
    
    I learnt this will working with the Meta AI API wrapper in Golang. The API uses payload as a URL encoded body and will append key-value pairs to the request body, the subtle difference can cause nil pointer access if not initialized and used the appropriate method correctly. I think I will write a blog post on this.
    
* Shuffling Two Lists keeping the order of the corresponding index the same:
    
    What I was doing was testing and evaluating some results on data, and that data was coming from a set of files in a folder, I wanted to randomly shuffle those values. I wanted to track the metrics from the data with the filename, so I created this little function that shuffles two or more lists in a random order and maintains a one-on-one index mapping.
    
    ```go
    import random
    List of file names
    file_names = ["file1.txt", "file2.txt", "file3.txt", "file4.txt", "file5.txt"]
    Corresponding sample data for each file
    sample_data = ["Data from file 1", "Data from file 2", "Data from file 3", "Data from file 4", "Data from file 5"]
    Create a list of indices
    indices = list(range(len(file_names)))
    Shuffle the indices
    random.shuffle(indices)
    Reorder both lists using the shuffled indices
    shuffled_file_names = [file_names[i] for i in indices]
    shuffled_sample_data = [sample_data[i] for i in indices]
    print(shuffled_file_names)
    print(shuffled_sample_data)
    """
    ['file3.txt', 'file5.txt', 'file1.txt', 'file2.txt', 'file4.txt']
    ['Data from file 3', 'Data from file 5', 'Data from file 1', 'Data from file 2', 'Data from file 4']
    """
    ```
    

This is the function, nice simple and modular.

```go
import random
def shuffle_lists(*lists):
"""
Shuffles two or more lists while keeping the order of corresponding elements the same.
Args:
*lists: Two or more lists to be shuffled.
Returns:
A tuple of shuffled lists with the same order of corresponding elements.
"""
# Check if all lists have the same length
if len(set(len(lst) for lst in lists)) != 1:
raise ValueError("All input lists must have the same length.")
# Create a list of indices
indices = list(range(len(lists[0])))
# Shuffle the indices
random.shuffle(indices)
# Reorder all lists using the shuffled indices
shuffled_lists = tuple([lst[i] for i in indices] for lst in lists)
return shuffled_lists
```

* LLMs are good at one single thing, Don’t ask too many things in a shot, make it sequential like a pipeline.
    
* Using Tabulate in Python to format a list of dictionaries pretty:
    
    This is really handy to work if you want to quickly visualise dictionaries and list like stuff.
    
    ```go
    from tabulate import tabulate
    List of books
    books = [
    {'Title': '1984', 'Author': 'George Orwell', 'Year': 1949, 'Genre': 'Dystopian'},
    {'Title': 'To Kill a Mockingbird', 'Author': 'Harper Lee', 'Year': 1960, 'Genre': 'Fiction'},
    {'Title': 'The Great Gatsby', 'Author': 'F. Scott Fitzgerald', 'Year': 1925, 'Genre': 'Classic'}
    ]
    Generate and print the table
    print(tabulate(books, headers="keys", tablefmt="grid"))
    ```
    
    Output:
    
    ```go
    +-----------------------+---------------------+--------+-----------+
    | Title                 | Author              |   Year | Genre     |
    +=======================+=====================+========+===========+
    | 1984                  | George Orwell       |   1949 | Dystopian |
    +-----------------------+---------------------+--------+-----------+
    | To Kill a Mockingbird | Harper Lee          |   1960 | Fiction   |
    +-----------------------+---------------------+--------+-----------+
    | The Great Gatsby      | F. Scott Fitzgerald |   1925 | Classic   |
    +-----------------------+---------------------+--------+-----------+
    ```
    

## Tech News

* [Hashnode Docs](https://hashnode.com/blog/announcing-docs-by-hashnode-the-content-engine-for-api-and-product-documentation): The content engine for API and product documentation.  
    This brings me in awe with the Hashnode speed and quality of development. The love with which they craft these products is truly commendable, they have the best free tiers for doing almost anything related to writing as a developer.
    
* [Artifacts on Antrhropic](https://www.anthropic.com/news/artifacts): Artifacts can be used to create, view, and iterate on platform-specific work directly within [Claude.ai](http://Claude.ai).
    
* Buildspace is closing: [https://buildspace.so/](https://buildspace.so/)
    
    In a tragic turn of events, the gradually growing community of builders is sadly shut down.
    

For more news, follow the Hackernewsletter [https://mailchi.mp/hackernewsletter/714](https://mailchi.mp/hackernewsletter/714) and for daily developer articles, join [daily.dev](http://daily.dev)

---

That’s it from this week, hope you did well this week, and have a happy week and weekend ahead!

Thank you for reading, let’s catch up in the next week.

Happy Coding :)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
