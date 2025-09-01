---
type: "logsql"
title: "SQLite dot commands: Output mode"
slug: sqlite-mode-dot-command
date: 2025-09-01
tags: ["sqlite", "sql"]
---

## Output mode dot command

The SQLite shell is a great way to test out and run bunch of queries, but the output might be squished or cluttered. You might need to view the output in a specific way, and the creators of SQLite were already hearing your screams.

They created not 5, not 10, but 14 different output modes, and to extend it even further, you can even use any delimiter/separator as well.

The `.mode` will give you the currently configured/set output mode with some options which are default for certain type of modes.

The default mode is `list`, with the option of `escape` with the ascii character set.

```
current output mode: list --escape ascii
```

The list mode defined in the [docs](https://www.sqlite.org/cli.html#changing_output_formats) as:

> In list mode, each row of a query result is written on one line of output and each column within that row is separated by a specific separator string. The default separator is a pipe symbol ("|"). List mode is especially useful when you are going to send the output of a query to another program (such as AWK) for additional processing.

So, in short, the `list` output mode prints each row on a line with a `|` as the separator.

You can get all the information you need with the `.help mode` command

```
.mode ?MODE? ?OPTIONS?   Set output mode
   MODE is one of:
     ascii       Columns/rows delimited by 0x1F and 0x1E
     box         Tables using unicode box-drawing characters
     csv         Comma-separated values
     column      Output in columns.  (See .width)
     html        HTML <table> code
     insert      SQL insert statements for TABLE
     json        Results in a JSON array
     line        One value per line
     list        Values delimited by "|"
     markdown    Markdown table format
     qbox        Shorthand for "box --wrap 60 --quote"
     quote       Escape answers as for SQL
     table       ASCII-art table
     tabs        Tab-separated values
     tcl         TCL list elements
   OPTIONS: (for columnar modes or insert mode):
     --escape T     ctrl-char escape; T is one of: symbol, ascii, off
     --wrap N       Wrap output lines to no longer than N characters
     --wordwrap B   Wrap or not at word boundaries per B (on/off)
     --ww           Shorthand for "--wordwrap 1"
     --quote        Quote output text as SQL literals
     --noquote      Do not quote output text
     TABLE          The name of SQL table used for "insert" mode
sqlite>
```


```sql
CREATE TABLE IF NOT EXISTS books(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    pages INTEGER NOT NULL,
    release_date TEXT NOT NULL,
    price REAL NOT NULL
);

INSERT INTO books(title, author, pages, release_date, price)
VALUES
('The Hobbit', 'J.R.R. Tolkien', 310, '1937-09-21', 39.99),
('The Fellowship of the Ring', 'J.R.R. Tolkien', 423, '1954-07-29', 49.99),
('The Two Towers', 'J.R.R. Tolkien', 352, '1954-11-11', 49.99), 
('The Return of the King', 'J.R.R. Tolkien', 416, '1955-10-20', 49.99);

SELECT * FROM books;
```

```sqlite
1|The Hobbit|J.R.R. Tolkien|310|1937-09-21|39.99
2|The Fellowship of the Ring|J.R.R. Tolkien|423|1954-07-29|49.99
3|The Two Towers|J.R.R. Tolkien|352|1954-11-11|49.99
4|The Return of the King|J.R.R. Tolkien|416|1955-10-20|49.99
```

Not the best way to look at the data, but handy in some cases. We can change it, we have 14 different modes to choose from.
Its like choosing a taste flovour for an ice-cream, you want something different each time, or you stick to the simple vanilla one. Which one are you? I am a vanilla guy (I like the table mode, but use csv heavily too).

Let's change it to a table format with `.mode table` This will set the mode as `table`

> table       ASCII-art table

This will show the result sets in a ascii-like table structure with the `+++` and `---` to separate the rows and columns.

```sql
SELECT * FROM books;
```

```
+----+----------------------------+----------------+-------+--------------+-------+
| id |           title            |     author     | pages | release_date | price |
+----+----------------------------+----------------+-------+--------------+-------+
| 1  | The Hobbit                 | J.R.R. Tolkien | 310   | 1937-09-21   | 39.99 |
| 2  | The Fellowship of the Ring | J.R.R. Tolkien | 423   | 1954-07-29   | 49.99 |
| 3  | The Two Towers             | J.R.R. Tolkien | 352   | 1954-11-11   | 49.99 |
| 4  | The Return of the King     | J.R.R. Tolkien | 416   | 1955-10-20   | 49.99 |
+----+----------------------------+----------------+-------+--------------+-------+
```

That is soo nice looking! Crystal clear.

You can even see the current set mode precisely with the `.mode` command

Since I changed the mode to table, let's see what the mode is now set as?
```
sqlite> .mode
current output mode: table --wrap 60 --wordwrap off --noquote --escape ascii
```

It has a bunch of options like `wrap`, `wordwrap`, `noquote`, and `escape` etc. You can take a look at the help and tweak them as per your liking, I usually don't change the options, but it might be very rare to switch from the defaults.


## Output modes

### ASCII 

This mode separates the columns with by `0x1F` and rows with `0x1E`.

```
.mode ascii
SELECT * FROM books;
```

```
idtitleauthorpagesrelease_dateprice1The HobbitJ.R.R. Tolkien3101937-09-2139.992The Fellowship of the RingJ.R.R. Tolkien4231954-07-2949.993The Two TowersJ.R.R. Tolkien3521954-11-1149.994The Return of the KingJ.R.R. Tolkien4161955-10-2049.99
```

### Box

The box mode renders the result set using unicode box-drawing characters.

```
.mode box
SELECT * FROM books;
```

```
┌────┬────────────────────────────┬────────────────┬───────┬──────────────┬───────┐
│ id │           title            │     author     │ pages │ release_date │ price │
├────┼────────────────────────────┼────────────────┼───────┼──────────────┼───────┤
│ 1  │ The Hobbit                 │ J.R.R. Tolkien │ 310   │ 1937-09-21   │ 39.99 │
│ 2  │ The Fellowship of the Ring │ J.R.R. Tolkien │ 423   │ 1954-07-29   │ 49.99 │
│ 3  │ The Two Towers             │ J.R.R. Tolkien │ 352   │ 1954-11-11   │ 49.99 │
│ 4  │ The Return of the King     │ J.R.R. Tolkien │ 416   │ 1955-10-20   │ 49.99 │
└────┴────────────────────────────┴────────────────┴───────┴──────────────┴───────┘
```


### CSV

The csv mode simply outputs the result in the comma-separated values for columns and newline character for rows.

```
.mode csv
SELECT * FROM books;
```

```csv
1,"The Hobbit","J.R.R. Tolkien",310,1937-09-21,39.99
2,"The Fellowship of the Ring","J.R.R. Tolkien",423,1954-07-29,49.99
3,"The Two Towers","J.R.R. Tolkien",352,1954-11-11,49.99
4,"The Return of the King","J.R.R. Tolkien",416,1955-10-20,49.99
```


### Column

The column mode simply outputs the result in columns with certain width. We can set the option `.width` to change the width.

```
.mode column
SELECT * FROM books;
```

```
id  title                       author          pages  release_date  price
--  --------------------------  --------------  -----  ------------  -----
1   The Hobbit                  J.R.R. Tolkien  310    1937-09-21    39.99
2   The Fellowship of the Ring  J.R.R. Tolkien  423    1954-07-29    49.99
3   The Two Towers              J.R.R. Tolkien  352    1954-11-11    49.99
4   The Return of the King      J.R.R. Tolkien  416    1955-10-20    49.99
```


### HTML

The html output mode simply renders the result set / tables into an table element in html.

```
.mode html
SELECT * FROM books;
```

```html
<TR><TH>id</TH>
<TH>title</TH>
<TH>author</TH>
<TH>pages</TH>
<TH>release_date</TH>
<TH>price</TH>
</TR>
<TR><TD>1</TD>
<TD>The Hobbit</TD>
<TD>J.R.R. Tolkien</TD>
<TD>310</TD>
<TD>1937-09-21</TD>
<TD>39.99</TD>
</TR>
<TR><TD>2</TD>
<TD>The Fellowship of the Ring</TD>
<TD>J.R.R. Tolkien</TD>
<TD>423</TD>
<TD>1954-07-29</TD>
<TD>49.99</TD>
</TR>
<TR><TD>3</TD>
<TD>The Two Towers</TD>
<TD>J.R.R. Tolkien</TD>
<TD>352</TD>
<TD>1954-11-11</TD>
<TD>49.99</TD>
</TR>
<TR><TD>4</TD>
<TD>The Return of the King</TD>
<TD>J.R.R. Tolkien</TD>
<TD>416</TD>
<TD>1955-10-20</TD>
<TD>49.99</TD>
</TR>
```


### Insert

The insert mode simply outputs the result in SQL insert statements. This is really handy for bulk inserting or migrating data across databases or tables.

```
.mode insert
SELECT * FROM books;
```

```sql
INSERT INTO "table"(id,title,author,pages,release_date,price) VALUES(1,'The Hobbit','J.R.R. Tolkien',310,'1937-09-21',39.99000000000000198);
INSERT INTO "table"(id,title,author,pages,release_date,price) VALUES(2,'The Fellowship of the Ring','J.R.R. Tolkien',423,'1954-07-29',49.99000000000000198);
INSERT INTO "table"(id,title,author,pages,release_date,price) VALUES(3,'The Two Towers','J.R.R. Tolkien',352,'1954-11-11',49.99000000000000198);
INSERT INTO "table"(id,title,author,pages,release_date,price) VALUES(4,'The Return of the King','J.R.R. Tolkien',416,'1955-10-20',49.99000000000000198);
```

### JSON

The json mode simply outputs the result in json format.

```
.mode json
SELECT * FROM books;
```

```json
[
    {"id":1,"title":"The Hobbit","author":"J.R.R. Tolkien","pages":310,"release_date":"1937-09-21","price":39.99000000000000198},
    {"id":2,"title":"The Fellowship of the Ring","author":"J.R.R. Tolkien","pages":423,"release_date":"1954-07-29","price":49.99000000000000198},
    {"id":3,"title":"The Two Towers","author":"J.R.R. Tolkien","pages":352,"release_date":"1954-11-11","price":49.99000000000000198},
    {"id":4,"title":"The Return of the King","author":"J.R.R. Tolkien","pages":416,"release_date":"1955-10-20","price":49.99000000000000198}
]
```


### List (The default)

As we already know, the list mode is the default mode.

```
.mode list
SELECT * FROM books;
```

```
id|title|author|pages|release_date|price
1|The Hobbit|J.R.R. Tolkien|310|1937-09-21|39.99
2|The Fellowship of the Ring|J.R.R. Tolkien|423|1954-07-29|49.99
3|The Two Towers|J.R.R. Tolkien|352|1954-11-11|49.99
4|The Return of the King|J.R.R. Tolkien|416|1955-10-20|49.99
```

### Markdown

The markdown mode simply outputs the result in markdown format. We can use this output in a markdown file and it will nicely draw the tables out.

```
.mode markdown
SELECT * FROM books;
```
```md
| id |           title            |     author     | pages | release_date | price |
|----|----------------------------|----------------|-------|--------------|-------|
| 1  | The Hobbit                 | J.R.R. Tolkien | 310   | 1937-09-21   | 39.99 |
| 2  | The Fellowship of the Ring | J.R.R. Tolkien | 423   | 1954-07-29   | 49.99 |
| 3  | The Two Towers             | J.R.R. Tolkien | 352   | 1954-11-11   | 49.99 |
| 4  | The Return of the King     | J.R.R. Tolkien | 416   | 1955-10-20   | 49.99 |
```

### Qbox

The qbox mode simply outputs the result in box format with the option as `--wrap 60 --quote`.

```
.mode qbox
SELECT * FROM books;
```

```
┌────┬──────────────────────────────┬──────────────────┬───────┬──────────────┬───────┐
│ id │            title             │      author      │ pages │ release_date │ price │
├────┼──────────────────────────────┼──────────────────┼───────┼──────────────┼───────┤
│ 1  │ 'The Hobbit'                 │ 'J.R.R. Tolkien' │ 310   │ '1937-09-21' │ 39.99 │
│ 2  │ 'The Fellowship of the Ring' │ 'J.R.R. Tolkien' │ 423   │ '1954-07-29' │ 49.99 │
│ 3  │ 'The Two Towers'             │ 'J.R.R. Tolkien' │ 352   │ '1954-11-11' │ 49.99 │
│ 4  │ 'The Return of the King'     │ 'J.R.R. Tolkien' │ 416   │ '1955-10-20' │ 49.99 │
└────┴──────────────────────────────┴──────────────────┴───────┴──────────────┴───────┘
```

### Quote

The quote mode simply outputs the result in SQL-string quote format.

```
.mode quote
SELECT * FROM books;
```
```
'id','title','author','pages','release_date','price'
1,'The Hobbit','J.R.R. Tolkien',310,'1937-09-21',39.99000000000000198
2,'The Fellowship of the Ring','J.R.R. Tolkien',423,'1954-07-29',49.99000000000000198
3,'The Two Towers','J.R.R. Tolkien',352,'1954-11-11',49.99000000000000198
4,'The Return of the King','J.R.R. Tolkien',416,'1955-10-20',49.99000000000000198
```

### Table

The table mode simply outputs the result in table format. We saw this mode in the previous example.

```
.mode table
SELECT * FROM books;
```

```
+----+----------------------------+----------------+-------+--------------+-------+
| id |           title            |     author     | pages | release_date | price |
+----+----------------------------+----------------+-------+--------------+-------+
| 1  | The Hobbit                 | J.R.R. Tolkien | 310   | 1937-09-21   | 39.99 |
| 2  | The Fellowship of the Ring | J.R.R. Tolkien | 423   | 1954-07-29   | 49.99 |
| 3  | The Two Towers             | J.R.R. Tolkien | 352   | 1954-11-11   | 49.99 |
| 4  | The Return of the King     | J.R.R. Tolkien | 416   | 1955-10-20   | 49.99 |
+----+----------------------------+----------------+-------+--------------+-------+
```

### Tabs

The tabs mode simply outputs the result in tab-separated values.

```
.mode tabs
SELECT * FROM books;
```
```
id      title   author  pages   release_date    price
1       The Hobbit      J.R.R. Tolkien  310     1937-09-21      39.99
2       The Fellowship of the Ring      J.R.R. Tolkien  423     1954-07-29      49.99
3       The Two Towers  J.R.R. Tolkien  352     1954-11-11      49.99
4       The Return of the King  J.R.R. Tolkien  416     1955-10-20      49.99
```

### Tcl 

The Tcl mode simply outputs the result in Tcl format. Tcl or tool command language is a configuration friendly format.
```
.mode tcl
SELECT * FROM books;
```
```tcl
"id" "title" "author" "pages" "release_date" "price"
"1" "The Hobbit" "J.R.R. Tolkien" "310" "1937-09-21" "39.99"
"2" "The Fellowship of the Ring" "J.R.R. Tolkien" "423" "1954-07-29" "49.99"
"3" "The Two Towers" "J.R.R. Tolkien" "352" "1954-11-11" "49.99"
"4" "The Return of the King" "J.R.R. Tolkien" "416" "1955-10-20" "49.99"
```
