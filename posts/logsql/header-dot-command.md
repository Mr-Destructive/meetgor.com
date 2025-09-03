---
type: "logsql"
title: "SQLite dot commands: header command"
slug: sqlite-dot-command-header
date: 2025-09-03
tags: ["sqlite", "sql"]
---

The `.headers` command is crucial as might effect the way the result set is displayed. The first row would be set to the name of the columns in the output of the relevant commands like `.output`, `.once`, or even your simple standard out queries if it is set on.

```
.headers on

OR

.header on
```

Some how either of them work. You need to set it to `on` to enable the headers in the output of the result set. And you can turn off with `.headers off`.

Which modes are effected with this command if set on or off?

Well we need to think about how effected means
1. Only added in the first row
2. Added in each row
3. No effect

---

1. Only added in the first row
    1. ascii
    2. column
    3. csv
    4. html
    5. list
    6. quote
    7. tabs
    8. tcl
2. Added in each row
    1. insert
3. No effect
    1. box
    2. json
    3. line
    4. markdown
    5. table


Reference table

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

Below are the outputs of each command

### Added on the first row

In the below modes, the headers command if set will add the column name as the header in the first row.

1. ASCII

```sqlite
.mode ascii
.headers on
select * from books;
.headers off
select * from books;
```

Output:
```
sqlite> .mode ascii
sqlite> .headers on
sqlite> select * from books;
idtitleauthorpagesrelease_dateprice1The HobbitJ.R.R. Tolkien3101937-09-2139.992The Fellowship of the RingJ.R.R. Tolkien4231954-07-2949.993The Two TowersJ.R.R. Tolkien3521954-11-1149.994The Return of the KingJ.R.R. Tolkien4161955-10-2049.99sqlite>
sqlite> .headers off
sqlite> select * from books;
1The HobbitJ.R.R. Tolkien3101937-09-2139.992The Fellowship of the RingJ.R.R. Tolkien4231954-07-2949.993The Two TowersJ.R.R. Tolkien3521954-11-1149.994The Return of the KingJ.R.R. Tolkien4161955-10-2049.99sqlite>
sqlite>
```

2. Column

```
.mode column
.headers on
select * from books;
.headers off
select * from books;
```

Output

```
sqlite> .mode column
sqlite> .headers on
sqlite> select * from books;
id  title                       author          pages  release_date  price
--  --------------------------  --------------  -----  ------------  -----
1   The Hobbit                  J.R.R. Tolkien  310    1937-09-21    39.99
2   The Fellowship of the Ring  J.R.R. Tolkien  423    1954-07-29    49.99
3   The Two Towers              J.R.R. Tolkien  352    1954-11-11    49.99
4   The Return of the King      J.R.R. Tolkien  416    1955-10-20    49.99
sqlite> .headers off
sqlite> select * from books;
1   The Hobbit                  J.R.R. Tolkien  310    1937-09-21    39.99
2   The Fellowship of the Ring  J.R.R. Tolkien  423    1954-07-29    49.99
3   The Two Towers              J.R.R. Tolkien  352    1954-11-11    49.99
4   The Return of the King      J.R.R. Tolkien  416    1955-10-20    49.99
sqlite>
```

3. CSV

```sqlite
.mode csv
.headers on
select * from books;
.headers off
select * from books;
```

Output:

```
sqlite> .mode csv
sqlite> .headers on
sqlite> select * from books;
id,title,author,pages,release_date,price
1,"The Hobbit","J.R.R. Tolkien",310,1937-09-21,39.99
2,"The Fellowship of the Ring","J.R.R. Tolkien",423,1954-07-29,49.99
3,"The Two Towers","J.R.R. Tolkien",352,1954-11-11,49.99
4,"The Return of the King","J.R.R. Tolkien",416,1955-10-20,49.99
sqlite> .headers off
sqlite> select * from books;
1,"The Hobbit","J.R.R. Tolkien",310,1937-09-21,39.99
2,"The Fellowship of the Ring","J.R.R. Tolkien",423,1954-07-29,49.99
3,"The Two Towers","J.R.R. Tolkien",352,1954-11-11,49.99
4,"The Return of the King","J.R.R. Tolkien",416,1955-10-20,49.99
sqlite>
```

4. HTML

```sqlite
.mode html
.headers on
select * from books;
.headers off
select * from books;
```

Output:

```
sqlite> .mode html
sqlite> .headers on
sqlite> select * from books;
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
sqlite> .headers off
sqlite> select * from books;
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
sqlite>
```

5. List

```sqlite
.mode list
.headers on
select * from books;
.headers off
select * from books;
```

Output:

```
sqlite> .mode list
sqlite> .headers on
sqlite> select * from books;
id|title|author|pages|release_date|price
1|The Hobbit|J.R.R. Tolkien|310|1937-09-21|39.99
2|The Fellowship of the Ring|J.R.R. Tolkien|423|1954-07-29|49.99
3|The Two Towers|J.R.R. Tolkien|352|1954-11-11|49.99
4|The Return of the King|J.R.R. Tolkien|416|1955-10-20|49.99
sqlite> .headers off
sqlite> select * from books;
1|The Hobbit|J.R.R. Tolkien|310|1937-09-21|39.99
2|The Fellowship of the Ring|J.R.R. Tolkien|423|1954-07-29|49.99
3|The Two Towers|J.R.R. Tolkien|352|1954-11-11|49.99
4|The Return of the King|J.R.R. Tolkien|416|1955-10-20|49.99
sqlite>
```

6. Quote

```sqlite
.mode quote
.headers on
select * from books;
.headers off
select * from books;
```

Output:

```
sqlite> .mode quote
sqlite> .headers on
sqlite> select * from books;
'id','title','author','pages','release_date','price'
1,'The Hobbit','J.R.R. Tolkien',310,'1937-09-21',39.99000000000000198
2,'The Fellowship of the Ring','J.R.R. Tolkien',423,'1954-07-29',49.99000000000000198
3,'The Two Towers','J.R.R. Tolkien',352,'1954-11-11',49.99000000000000198
4,'The Return of the King','J.R.R. Tolkien',416,'1955-10-20',49.99000000000000198
sqlite> .headers off
sqlite> select * from books;
1,'The Hobbit','J.R.R. Tolkien',310,'1937-09-21',39.99000000000000198
2,'The Fellowship of the Ring','J.R.R. Tolkien',423,'1954-07-29',49.99000000000000198
3,'The Two Towers','J.R.R. Tolkien',352,'1954-11-11',49.99000000000000198
4,'The Return of the King','J.R.R. Tolkien',416,'1955-10-20',49.99000000000000198
sqlite>

```

7. Tabs

```sqlite
.mode tabs
.headers on
select * from books;
.headers off
select * from books;
```

Output:

```
sqlite> .mode tabs
sqlite> .headers on
sqlite> select * from books;
id      title   author  pages   release_date    price
1       The Hobbit      J.R.R. Tolkien  310     1937-09-21      39.99
2       The Fellowship of the Ring      J.R.R. Tolkien  423     1954-07-29      49.99
3       The Two Towers  J.R.R. Tolkien  352     1954-11-11      49.99
4       The Return of the King  J.R.R. Tolkien  416     1955-10-20      49.99
sqlite> .headers off
sqlite> select * from books;
1       The Hobbit      J.R.R. Tolkien  310     1937-09-21      39.99
2       The Fellowship of the Ring      J.R.R. Tolkien  423     1954-07-29      49.99
3       The Two Towers  J.R.R. Tolkien  352     1954-11-11      49.99
4       The Return of the King  J.R.R. Tolkien  416     1955-10-20      49.99
sqlite>
```

8. TCL

```sqlite
.mode tcl
.headers on
select * from books;
.headers off
select * from books;
```

Output:

```
sqlite> .mode tcl
sqlite> .headers on
sqlite> select * from books;
"id" "title" "author" "pages" "release_date" "price"
"1" "The Hobbit" "J.R.R. Tolkien" "310" "1937-09-21" "39.99"
"2" "The Fellowship of the Ring" "J.R.R. Tolkien" "423" "1954-07-29" "49.99"
"3" "The Two Towers" "J.R.R. Tolkien" "352" "1954-11-11" "49.99"
"4" "The Return of the King" "J.R.R. Tolkien" "416" "1955-10-20" "49.99"
sqlite> .headers off
sqlite> select * from books;
"1" "The Hobbit" "J.R.R. Tolkien" "310" "1937-09-21" "39.99"
"2" "The Fellowship of the Ring" "J.R.R. Tolkien" "423" "1954-07-29" "49.99"
"3" "The Two Towers" "J.R.R. Tolkien" "352" "1954-11-11" "49.99"
"4" "The Return of the King" "J.R.R. Tolkien" "416" "1955-10-20" "49.99"
sqlite>
```


### Added in each row

In this type, each row has the column name as some form of the row.

1. Insert

```sqlite
.mode insert
.headers on
select * from books;
.headers off
select * from books;
```

Output:

```
sqlite> .mode insert
sqlite> .headers on
sqlite> select * from books;
INSERT INTO "table"(id,title,author,pages,release_date,price) VALUES(1,'The Hobbit','J.R.R. Tolkien',310,'1937-09-21',39.99000000000000198);
INSERT INTO "table"(id,title,author,pages,release_date,price) VALUES(2,'The Fellowship of the Ring','J.R.R. Tolkien',423,'1954-07-29',49.99000000000000198);
INSERT INTO "table"(id,title,author,pages,release_date,price) VALUES(3,'The Two Towers','J.R.R. Tolkien',352,'1954-11-11',49.99000000000000198);
INSERT INTO "table"(id,title,author,pages,release_date,price) VALUES(4,'The Return of the King','J.R.R. Tolkien',416,'1955-10-20',49.99000000000000198);
sqlite> .headers off
sqlite> select * from books;
INSERT INTO "table" VALUES(1,'The Hobbit','J.R.R. Tolkien',310,'1937-09-21',39.99000000000000198);
INSERT INTO "table" VALUES(2,'The Fellowship of the Ring','J.R.R. Tolkien',423,'1954-07-29',49.99000000000000198);
INSERT INTO "table" VALUES(3,'The Two Towers','J.R.R. Tolkien',352,'1954-11-11',49.99000000000000198);
INSERT INTO "table" VALUES(4,'The Return of the King','J.R.R. Tolkien',416,'1955-10-20',49.99000000000000198);
sqlite>

```


### No effect

These modes show the headers or column names irrespective of the `.headers` flag, as for these modes the headers form quite a bit of structure to the output they are showing.

1. Box

```sqlite
.mode box
.headers on
select * from books;
.headers off
select * from books;
```

Output:

```
sqlite> .mode box
sqlite> .headers on
sqlite> select * from books;
┌────┬────────────────────────────┬────────────────┬───────┬──────────────┬───────┐
│ id │           title            │     author     │ pages │ release_date │ price │
├────┼────────────────────────────┼────────────────┼───────┼──────────────┼───────┤
│ 1  │ The Hobbit                 │ J.R.R. Tolkien │ 310   │ 1937-09-21   │ 39.99 │
│ 2  │ The Fellowship of the Ring │ J.R.R. Tolkien │ 423   │ 1954-07-29   │ 49.99 │
│ 3  │ The Two Towers             │ J.R.R. Tolkien │ 352   │ 1954-11-11   │ 49.99 │
│ 4  │ The Return of the King     │ J.R.R. Tolkien │ 416   │ 1955-10-20   │ 49.99 │
└────┴────────────────────────────┴────────────────┴───────┴──────────────┴───────┘
sqlite> .headers off
sqlite> select * from books;
┌────┬────────────────────────────┬────────────────┬───────┬──────────────┬───────┐
│ id │           title            │     author     │ pages │ release_date │ price │
├────┼────────────────────────────┼────────────────┼───────┼──────────────┼───────┤
│ 1  │ The Hobbit                 │ J.R.R. Tolkien │ 310   │ 1937-09-21   │ 39.99 │
│ 2  │ The Fellowship of the Ring │ J.R.R. Tolkien │ 423   │ 1954-07-29   │ 49.99 │
│ 3  │ The Two Towers             │ J.R.R. Tolkien │ 352   │ 1954-11-11   │ 49.99 │
│ 4  │ The Return of the King     │ J.R.R. Tolkien │ 416   │ 1955-10-20   │ 49.99 │
└────┴────────────────────────────┴────────────────┴───────┴──────────────┴───────┘
sqlite>
```

2. JSON

```sqlite
.mode json
.headers on
select * from books;
.headers off
select * from books;
```

Output:

```
sqlite> .mode json
sqlite> .headers on
sqlite> select * from books;
[{"id":1,"title":"The Hobbit","author":"J.R.R. Tolkien","pages":310,"release_date":"1937-09-21","price":39.99000000000000198},
{"id":2,"title":"The Fellowship of the Ring","author":"J.R.R. Tolkien","pages":423,"release_date":"1954-07-29","price":49.99000000000000198},
{"id":3,"title":"The Two Towers","author":"J.R.R. Tolkien","pages":352,"release_date":"1954-11-11","price":49.99000000000000198},
{"id":4,"title":"The Return of the King","author":"J.R.R. Tolkien","pages":416,"release_date":"1955-10-20","price":49.99000000000000198}]
sqlite> .headers off
sqlite> select * from books;
[{"id":1,"title":"The Hobbit","author":"J.R.R. Tolkien","pages":310,"release_date":"1937-09-21","price":39.99000000000000198},
{"id":2,"title":"The Fellowship of the Ring","author":"J.R.R. Tolkien","pages":423,"release_date":"1954-07-29","price":49.99000000000000198},
{"id":3,"title":"The Two Towers","author":"J.R.R. Tolkien","pages":352,"release_date":"1954-11-11","price":49.99000000000000198},
{"id":4,"title":"The Return of the King","author":"J.R.R. Tolkien","pages":416,"release_date":"1955-10-20","price":49.99000000000000198}]
sqlite>
```

3. Line

```sqlite
.mode line
.headers on
select * from books;
.headers off
select * from books;
```

Output:
```
sqlite> .mode line
sqlite> .headers on
sqlite> select * from books;
          id = 1
       title = The Hobbit
      author = J.R.R. Tolkien
       pages = 310
release_date = 1937-09-21
       price = 39.99

          id = 2
       title = The Fellowship of the Ring
      author = J.R.R. Tolkien
       pages = 423
release_date = 1954-07-29
       price = 49.99

          id = 3
       title = The Two Towers
      author = J.R.R. Tolkien
       pages = 352
release_date = 1954-11-11
       price = 49.99

          id = 4
       title = The Return of the King
      author = J.R.R. Tolkien
       pages = 416
release_date = 1955-10-20
       price = 49.99


sqlite> .headers off
sqlite> select * from books;
          id = 1
       title = The Hobbit
      author = J.R.R. Tolkien
       pages = 310
release_date = 1937-09-21
       price = 39.99

          id = 2
       title = The Fellowship of the Ring
      author = J.R.R. Tolkien
       pages = 423
release_date = 1954-07-29
       price = 49.99

          id = 3
       title = The Two Towers
      author = J.R.R. Tolkien
       pages = 352
release_date = 1954-11-11
       price = 49.99

          id = 4
       title = The Return of the King
      author = J.R.R. Tolkien
       pages = 416
release_date = 1955-10-20
       price = 49.99
sqlite>
```

4. Markdown

```sqlite
.mode markdown
.headers on
select * from books;
.headers off
select * from books;
```

Output:

```
sqlite> .mode markdown
sqlite> .headers on
sqlite> select * from books;
| id |           title            |     author     | pages | release_date | price |
|----|----------------------------|----------------|-------|--------------|-------|
| 1  | The Hobbit                 | J.R.R. Tolkien | 310   | 1937-09-21   | 39.99 |
| 2  | The Fellowship of the Ring | J.R.R. Tolkien | 423   | 1954-07-29   | 49.99 |
| 3  | The Two Towers             | J.R.R. Tolkien | 352   | 1954-11-11   | 49.99 |
| 4  | The Return of the King     | J.R.R. Tolkien | 416   | 1955-10-20   | 49.99 |
sqlite> .headers off
sqlite> select * from books;
| id |           title            |     author     | pages | release_date | price |
|----|----------------------------|----------------|-------|--------------|-------|
| 1  | The Hobbit                 | J.R.R. Tolkien | 310   | 1937-09-21   | 39.99 |
| 2  | The Fellowship of the Ring | J.R.R. Tolkien | 423   | 1954-07-29   | 49.99 |
| 3  | The Two Towers             | J.R.R. Tolkien | 352   | 1954-11-11   | 49.99 |
| 4  | The Return of the King     | J.R.R. Tolkien | 416   | 1955-10-20   | 49.99 |
sqlite>
```

5. Table

```sqlite
.mode table
.headers on
select * from books;
.headers off
select * from books;
```

Output:

```
sqlite> .mode table
sqlite> .headers on
sqlite> select * from books;
+----+----------------------------+----------------+-------+--------------+-------+
| id |           title            |     author     | pages | release_date | price |
+----+----------------------------+----------------+-------+--------------+-------+
| 1  | The Hobbit                 | J.R.R. Tolkien | 310   | 1937-09-21   | 39.99 |
| 2  | The Fellowship of the Ring | J.R.R. Tolkien | 423   | 1954-07-29   | 49.99 |
| 3  | The Two Towers             | J.R.R. Tolkien | 352   | 1954-11-11   | 49.99 |
| 4  | The Return of the King     | J.R.R. Tolkien | 416   | 1955-10-20   | 49.99 |
+----+----------------------------+----------------+-------+--------------+-------+
sqlite> .headers off
sqlite> select * from books;
+----+----------------------------+----------------+-------+--------------+-------+
| id |           title            |     author     | pages | release_date | price |
+----+----------------------------+----------------+-------+--------------+-------+
| 1  | The Hobbit                 | J.R.R. Tolkien | 310   | 1937-09-21   | 39.99 |
| 2  | The Fellowship of the Ring | J.R.R. Tolkien | 423   | 1954-07-29   | 49.99 |
| 3  | The Two Towers             | J.R.R. Tolkien | 352   | 1954-11-11   | 49.99 |
| 4  | The Return of the King     | J.R.R. Tolkien | 416   | 1955-10-20   | 49.99 |
+----+----------------------------+----------------+-------+--------------+-------+
sqlite>

```

