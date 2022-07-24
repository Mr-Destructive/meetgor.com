---
cover: ''
date: 2022-03-28
datetime: 2022-03-28 00:00:00+00:00
description: Perform search-replace operation in a file using python
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/tils/2022-03-28-Python-search-replace-in-file.md
long_description: Using simple python semantics, we can perform search and replace
  in a file. Firstly, we will define the file name, along with the words to search
  and replace. After defining the sets of variables, we will open the file in  We
  will store the entire fi
now: 2022-07-24 13:37:38.457943
path: blog/tils/2022-03-28-Python-search-replace-in-file.md
prevnext: null
slug: python-search-replace-file
status: published-til
tags:
- python
templateKey: til
title: 'Python: Search and Replace in File'
today: 2022-07-24
---

## Searching and Replacing the text in a File

Using simple python semantics, we can perform search and replace in a file. Firstly, we will define the file name, along with the words to search and replace. After defining the sets of variables, we will open the file in `r+` mode i.e. we can perform read as well as write operations in the file.

We will store the entire file contents using the [read](https://docs.python.org/3/tutorial/inputoutput.html#reading-and-writing-files) function, the contents of file are now stored in the form of a string. We further can set the position of the cursor or the current position in the file using the [seek](https://python-reference.readthedocs.io/en/latest/docs/file/seek.html) function. The seek function takes in a optional parameter as the position to set for reading/writing of file. Using the [truncate](https://python-reference.readthedocs.io/en/latest/docs/file/truncate.html) function, we can clear all the contents of the file.

Finally, we generate the content by replacing the words i.e. the old word with the new word from the string which we store the contents. After replacing the content in the string, we write the string variable into the file and hence the substitution was performed in the text file.

```python
file_name = 'temp.txt'
old_text = 'foo'
new_text = 'python'

with open(file_name, "r+") as fname:
    lines = fname.read()
    fname.seek(0)
    fname.truncate(0)
    subs = lines.replace(old_text, new_text)
    fname.write(subs)
```

![Python Search Replace in File](https://res.cloudinary.com/techstructive-blog/image/upload/v1648479344/blog-media/cstvfdlazyfriwvnilju.png)