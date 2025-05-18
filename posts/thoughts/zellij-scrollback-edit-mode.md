{"author":"meet","date":"2025-03-20","post_dir":"thoughts","published":"published","title":"Zellij Open Scrollback Edit mode","type":"thoughts"}

I was looking at a long list of logs, (debugging of course). 

I had a list of transactions
- Two lists of transactions, one as ground truth and other as predicted.  Transactions is a list of object(dictionary or map), each object has fields like date, amount, description, etc.
- Those objects in a list need not necessarily be in order, however I do want to compare them, how to do that?
- I decided to sort them based on date(one field in that object).

Then for each date, I group the transactions and this would narrow down the search space for comparison of one-one transaction, since now I can compare which of the ones are closely matching, the date will be a exact match the amount should also be, the description can be fuzzily matched.

However, for amount, I guess I was wrong, there could be a value like `10` and the other could have `10.01` and those python doesn't count equal atleast when compared as a string. I converted to float and compared rounded off numbers.

Now the problem kicked in and print statements flooded, dates everywhere.

Now I was using ghostty terminal, and there was definitely some scroll limit so I couldn't scroll all the way to start. So I opened up zellij.

I had a log with *MATCHED* text where I logged the date and amount where both transaction matched.

Now, I wanted a count of these matches. I could use search but it was not giving the count of occurrences, I can't keep counting with my mouth(that idea flew by though)

Now, that's where I accidentally hit `<Ctrl>S` and `E`
And I was in a editor, woah!

I was excited, I could finally copy and throw that in vim and get everything I want.
hehe
I tried +"y but it didn't copy to the clipboard, it yanked yes, but not in the system clipboard. That frustrated me and took my hope down, but I googled it and also gpted it (is that a word, I think we can say llmed it).

And yes we can set the `export $EDITOR` to the editor executable path and it would open that thing in that editor.

I did and it did work.

That mode is called [scrollback-edit](https://zellij.dev/news/edit-scrollback-compact/) mode. I should say a life saver mode, a log viewer and really cool.

I am probably too dumb and I know this exists in tmux, but I felt good and helped me solve my problem. So thank you whoever made that mode, its really helpful to debug with logs (debloging) Yes I am bad at naming things, but I like this more than vibe coding ;)
