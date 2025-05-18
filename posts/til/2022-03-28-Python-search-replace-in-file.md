{
    "type": "til",
    "title": "Python: Search and Replace in File",
    "description": "Perform search-replace operation in a file using python",
    "date": "2022-03-28 18:59:30",
    "status": "published",
    "slug": "python-search-replace-file",
    "tags": ["python"]
}

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

