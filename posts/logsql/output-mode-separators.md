---
type: "logsql"
title: "SQLite dot commands: Output mode separator command"
slug: sqlite-mode-dot-command-separators
date: 2025-09-02
tags: ["sqlite", "sql"]
---

## Using the separator for the ouput

If you wanted to use a specific separator for columns and rows while displaying the result set / table, you can use the `.separator` dot command which can take 2 arguments, first as the separator for the column and the second for the row.

So, if we set use `.separator "|" "---"` then it will split the columns with `|` and each row with `---`. 

```
1|The Hobbit|J.R.R. Tolkien|310|1937-09-21|39.99---2|The Fellowship of the Ring|J.R.R. Tolkien|423|1954-07-29|49.99---3|The Two Towers|J.R.R. Tolkien|352|1954-11-11|49.99---4|The Return of the King|J.R.R. Tolkien|416|1955-10-20|49.99---
```

The output looks wired but I was giving a example.

The row separator is by default a `\n` character or `\r\n` on windows, which is for the list or any other mode. However if you want to add those, you need to specify it in the string like below:

```
.separator "|" "\n---"
```

```
>sqlite>.separator "|" "\n---"
sqlite> select * from books;
1|The Hobbit|J.R.R. Tolkien|310|1937-09-21|39.99
---2|The Fellowship of the Ring|J.R.R. Tolkien|423|1954-07-29|49.99
---3|The Two Towers|J.R.R. Tolkien|352|1954-11-11|49.99
---4|The Return of the King|J.R.R. Tolkien|416|1955-10-20|49.99
---sqlite>
```

The `\n---` will add it at the end of each row, so the first row won't have it if you are seeing the above output and confused, then that's the expected result.

You can also use `\r` or `\t` for the row separator, and for the column separator you can use `\n` or `\t` or `\r`.

The modes that will respect the separator setting:
- list
- quote

Sadly only two, since for changing the row separator you need to pass the column separator first, so that defeats the purpose of the change in the row separator, you'll have to override the column separator first and then the row.

So, for the mode `tabs`, if you want to keep the columns separator as is, but change the row separator then you can do something like this:
```
.mode tabs
.separator "\t" "\n\n"
```

This will keep the column separator as is and change the row separator to `\n\n`

```
sqlite> .mode tabs
sqlite> .separator "\t" "\n\n"
sqlite> select * from books;
id      title   author  pages   release_date    price

1       The Hobbit      J.R.R. Tolkien  310     1937-09-21      39.99

2       The Fellowship of the Ring      J.R.R. Tolkien  423     1954-07-29      49.99

3       The Two Towers  J.R.R. Tolkien  352     1954-11-11      49.99

4       The Return of the King  J.R.R. Tolkien  416     1955-10-20      49.99

sqlite>
```
Similarly for csv you would keep the separator as , and then pass the row separator

Nice tricks to know.

References:
- [SQLite dot commands: Changing output formats](https://www.sqlite.org/cli.html#changing_output_formats)
