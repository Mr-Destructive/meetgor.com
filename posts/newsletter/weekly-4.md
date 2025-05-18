---
type: "newsletter"
title: "Techstructive Weekly #4"
date: 2024-08-24
---

## Techstructive Weekly #4

It was an exhilarating week for me, learned a lot of stuff, wrote a lot of code, pushed a lot of bugs, and tweeted a lot. This week, I wrote a lot of Python scripts at my work for experimentation setup and testing a lot of things, I was learning and prototyping my idea of Audiofy with Appwrite cloud, and read some inspiring articles.

Let’s refresh the developer’s life a bit and think philosophically.

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.

## Quote/Thought of the week

> "Learning is not attained by chance, it must be sought for with ardor and attended to with diligence." – [Abigail Adams](https://en.wikiquote.org/wiki/Abigail_Adams)

Not every week is the same, if it is then you are not learning, you are not growing. Stepping outside of the comfort zone is the way to get out of that creative slump. There needs to be hard work and curiosity to drive learning, it won’t come by thinking about how worse the situation is, taking action is the key, if you fail, you learn something.

---

## Read

* A Letter to a friend who is thinking of starting something new:
    
    ![](https://substack-post-media.s3.amazonaws.com/public/images/7163f0b1-f8d3-40c0-8df9-43fead0a6260_1008x1008.png align="left")
    
    [The Sublime](https://sublimeinternet.substack.com/p/letter-to-a-friend-who-is-thinking-410?utm_source=substack&utm_campaign=post_embed&utm_medium=web)
    
    [letter to a friend who is thinking of starting something new](https://sublimeinternet.substack.com/p/letter-to-a-friend-who-is-thinking-410?utm_source=substack&utm_campaign=post_embed&utm_medium=web)
    
    [Hi hi…](https://sublimeinternet.substack.com/p/letter-to-a-friend-who-is-thinking-410?utm_source=substack&utm_campaign=post_embed&utm_medium=web)
    
    [Read more](https://sublimeinternet.substack.com/p/letter-to-a-friend-who-is-thinking-410?utm_source=substack&utm_campaign=post_embed&utm_medium=web)
    
    [4 months ago · 68 likes · 8 comments · Sari Azout](https://sublimeinternet.substack.com/p/letter-to-a-friend-who-is-thinking-410?utm_source=substack&utm_campaign=post_embed&utm_medium=web)
    
    This blog actually had me think over some of the questions I have by somehow never pondering over more deeply. Because to be honest, I am afraid, those answers will leave me in either fear, guilt, or anxiety. It’s better to answer those and accept them because nothing is perfect in this world.
    
* [17 Things you should know about Python Lists to not regret later](https://zlliu.medium.com/12-things-i-regret-not-knowing-earlier-about-python-lists-a71dd8a435e2): As a developer who uses Python at work, I find this extremely helpful, I use most of these almost every day, though some of them also are new to me like list unpacking, not used often much. Some of them are intuitive to use in some cases, which makes Python nice to write.
    
* [Go is my hammer and everything seems like a nail](https://www.maragu.dev/blog/go-is-my-hammer-and-everything-is-a-nail?ref=dailydev): I get that if someone loves to write in a language, it becomes instinct to use it everywhere, but this got me a bit of questioning, like really go for everything? Not sure, like surely could be used, but not necessarily should be used. Would you use Golang to code a frontend application? No right? It’s not meant to be, surely you could write HTML and say I coded a full-stack app in Golang, but really? Been there, done that. Doesn’t feel like touching the code again.
    
* [Keeping a daily working journal](https://blog.isquaredsoftware.com/2020/09/coding-career-advice-daily-work-journal/?ref=dailydev): I love this idea of journalling daily, I have a prototype of an idea that I made a few weeks back that could be fitted into this domain. Basically an organisation-level blogs for individuals.
    
* [Why you should make a New programming language](https://ntietz.com/blog/you-should-make-a-new-terrible-programming-language/): This blog actually hit me and gave me a hint of my old project of Substack Markdown Parser. So have started to read a couple of blogs and books mentioned in this blog to start understanding interpreters.
    
* [Crafting Interpreters: Introduction](https://craftinginterpreters.com/introduction.html):  As said in the above blog, I am reading this book blog to get some overview of the working of interpreters. I have created a parser for markdown to substack json and vice versa, that is some Python shenanigans. But want to make it a proper parser and interpreter-like tool.
    
* Appwrite: [Serverless Function 101 Best Practices](https://appwrite.io/blog/post/serverless-functions-best-practices): I have been playing with Appwrite’s functions the whole week on the side, it is well documented and this blog added one more silver lining to the rich documentation.
    

## Watched

* Appwrite Init Videos:
    
    * Function Local Development:
        
    * CLI Revamped
        
    * Functions:
        

* [HTMX CalmCode Crash Course](https://calmcode.io/course/htmx/introduction): A really good quick course on HTMX, highly recommend starting here to get a good understanding of why this library exists.
    
* Vim Register and Macros Trick:
    
    This is actually helpful and I knew some of the tricks but being able to get a reminder of that makes me click a few things that I need to change in my daily workflow.
    
* Be a full stack developer with Golang and React: Trolled
    

## Learnt

* A list of dictionaries in Python to Pandas Dataframe: This looks quite simple but is really handy and have never thought of it deeply.
    
    ```go
    import pandas as pd
    List of dictionaries representing bank statement transactions
    data = [
    {'Date': '2024-08-01', 'Description': 'Salary Deposit', 'Transaction Type': 'Credit', 'Amount': 3000.00, 'Balance': 3500.00},
    {'Date': '2024-08-03', 'Description': 'Grocery Store Purchase', 'Transaction Type': 'Debit', 'Amount': 150.00, 'Balance': 3350.00},
    {'Date': '2024-08-05', 'Description': 'ATM Withdrawal', 'Transaction Type': 'Debit', 'Amount': 200.00, 'Balance': 3150.00}
    ]
    Convert the list of dictionaries to a Pandas DataFrame
    df = pd.DataFrame(data)
    Display the DataFrame
    print(df)
    """
    Date          Description Transaction Type  Amount  Balance
    0  2024-08-01       Salary Deposit      Credit  3000.0   3500.0
    1  2024-08-03    Grocery Store Purchase Debit   150.0   3350.0
    2  2024-08-05        ATM Withdrawal     Debit   200.0   3150.0
    """
    ```
    
* Convert a dict with a number as a key and any as a value into a list of ordered numbers with elements as the value of the dict.
    
    ```go
    # Dictionary with string keys representing numbers and consistent list values
    data = {
    "3": [30, 31, 32],
    "1": [10],
    "2": [20, 21]
    }
    Convert the dictionary into a list of values, ordered by the numerical interpretation of the keys
    ordered_values = [value for key, value in sorted(data.items(), key=lambda x: int(x[0]))]
    Display the ordered list of values
    print(ordered_values)
    [ [10], [20, 21], [30, 31, 32] ]
    ```
    
* [Pip Install a git repo as a python package from a specific branch from the requriements.txt file](https://stackoverflow.com/questions/16584552/how-to-state-in-requirements-txt-a-direct-github-source): I was trying to install the python SDK for Appwrite but the latest release was not live yet, so decided to download the version branch from the [GitHub repo](https://github.com/appwrite/sdk-for-python/tree/1.6.x). However, that failed as in the Appwrtie function environment, or any cloud/serverless function ecosystem there won’t be git installed.
    
* [Appwrite Functions in Python](https://appwrite.io/docs/products/functions/develop): As I have said, I have been experimenting with the Appwrite function ecosystem throughout the week on the side. I was basically trying to use Python functions for some prototyping ideas. I must say the documentation is really good. I learnt how to parse binary data into a response, a structure for a sample Python function, and set an appwrite json locally for functions.
    
* [Table extraction from Images using GPT-4 Vision and Python](https://python.useinstructor.com/examples/extracting_tables/): This is something I have been researching at my work and found it interesting to read and experiment upon. I am amazed at how beautifully LLMs can parse and get back data.
    

## Tech News

* [Appwrite Init](https://appwrite.io/init): It was a great week with a ton of launches and great features. i was really looking for Golang support and it is really a game changer for the developer experience of quickly making a backend with great control.
    
* [Redis 8 and AI capabilities](https://redis.io/blog/introducing-another-era-of-fast/): Redis since it changed its license, gets the [biggest release ever](https://techcrunch.com/2024/08/23/after-changing-its-license-redis-drops-its-biggest-release-yet/#:~:text=Redis%2C%20the%20company%20behind%20the,the%20launch%20of%20Redis%208.). As any other thing launching in 2024 has the word **AI** in it.
    

For more news, follow the Hackernewsletter [https://mailchi.mp/hackernewsletter/713](https://mailchi.mp/hackernewsletter/713)

---

That’s it from this week, I hope you did well this week, and have a happy week and weekend ahead!

Follow me on [Twitter/X](https://x.com/meetgor21) for more of my thoughts and tech stuff.

Thank you for reading, let’s catch up in the next week.

Happy Coding :)

Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.
