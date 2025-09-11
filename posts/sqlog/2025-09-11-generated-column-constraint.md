---
type: sqlog
slug: sqlite-generated-column-constraint
title: 'SQLite SQL: GENERATED column constraint'
date: "2025-09-11"
tags: ["sqlite", "sql"]
---

## The GENERATED column constraint

We saw certain limitations of the `DEFAULT` column constraint clause, that it can't generate dynamic values based on other columns or some other dynamic popularities. For overcomming those kind of constraints, there is one more specific constraint, the [GENERATED](https://sqlite.org/gencol.html) column constraint.

You can think of it as a `DEFAULT` but having more than static values, it is evaluated and computed each time a new row is inserted. However they come with one catch, you can't later update the values, you need to update the values of the dependent columns first, so it will generate **ALWAYS** a update is made to the row(relevant columns in that row).

Let's create a simple table to understand the `GENERATED` column constraint:

```sql
CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    content TEXT,
    word_count INTEGER GENERATED ALWAYS AS (length(content) - length(replace(content, ' ', '')) + 1),
    read_time_minutes INTEGER GENERATED ALWAYS AS (ceil(word_count / 200.0))
);
```

The syntax of the `GENERATED` constraint is 

```
column_name <type> GENERATED ALWAYS AS <expression>
```

Nothing fancy, just a content field as a text, then based on that value, the `word_count` and `read_time_minutes` is always calculated/generated based on the value of the `content` and the `read_time_minutes` is inturn dependent on the `word_count` value. So, the change in `content` will make a update on the `word_count` that will inturn update the `read_time_minutes`, like a chain reaction of sorts.

Let's insert a sample post, note, we can't insert a generated value.

```sql
INSERT INTO posts(content) VALUES('A sample post');
```

This will create a record with the `id` as 1, the `content` as `A sample post`, and the `word_count` will be generated/calculated as the `(length(content) - length(replace(content, ' ', ''))+1)` as the number of spaces + 1. The `read_time_minutes` is generated with the `word_count` divided by 200 which is the average word read in a minute by a human (let's not go into the nity-grity) but you can get the point of this. These are just mathematical values, and those are dependent on the value of the other columns of the same record/row.


```sql
SELECT * FROM posts;
```

As, you can see here, the value of `word_count` is `3` and the value of `read_time_minutes` is `1`.

Let's insert a few more values and check how these generated values behave.


Let's insert a bit longer text.

```sql
INSERT INTO posts(content) VALUES('A SQLITE Post for the generated column constraint.');
```

```sql
SELECT * FROM posts;
```

As, you can see the value of the `word_count` is `8` and the value of `read_time_minutes` is `1`.

Let's update the `content` and check if the `word_count` and the `read_time_minutes` are updated.

```sql
UPDATE posts SET content = content || 'New content
some more content' where id = 2;

SELECT * FROM posts;
```

As you can see the value of the `word_count` is `11` and the value of `read_time_minutes` is `1` still. That is because, until we write `200` words it will ceil off to 1.

Let's add a record with more than `200` words, and check if the `read_time_minutes` is updated.

```sql
INSERT INTO posts(content) VALUES('New post');
SELECT * FROM posts;
```

The current value of the `word_count` is `2` and the `read_time_minutes` is `1`, let's update it by adding the value of `Adding some extra words here.`.

```sql
UPDATE posts 
    SET content = content || replace(
        printf('%200c', ' '),
        ' ',
        ' Adding some extra words here.'
    )
WHERE id = 3;
```

Ok, this might be too much, but let's break it down:

- First, we are updating the `content` by adding the value of `Adding some extra words here.` to the end of the existing value of `content`.
- In the replace function, the parameters are original string, the value to replace, and the value to replace it with.
- So in this case, we are taking the string `printf('%200c', ' ')`, which is a string of 200 spaces, basically adding a single space character, but it will be right-justified within a field of 200 characters.
- The second parameter is the character to replace, in this case it is single white space or `' '` character.
- The third parameter is the text to replace that `' '` space character with, here it is ` Adding some extra words here`.
- So, essentially we are taking `200` white spaces and adding ` Adding some extra words here.` for each space to the end of the existing value of `content`.

Hence we will have around ~`200 * 5` words i.e. 1000 words, this way we can check the value of the `read_time_minutes` is updated.

```sql
SELECT * FROM posts;
```

As you can see the value of the `word_count` is `1002` and the value of `read_time_minutes` is `6`. That is perfect. And it is getting updated for each updation to the `content` column.

Here's a step by step output log: Its a SQLog :)

```
sqlite> CREATE TABLE posts (
   id INTEGER PRIMARY KEY,
   content TEXT,
   word_count INTEGER GENERATED ALWAYS AS (length(content) - length(replace(content, ' ', '')) + 1),
   read_time_minutes INTEGER GENERATED ALWAYS AS (ceil(word_count / 200.0))
);
sqlite> .mode table

sqlite> INSERT INTO posts(content) VALUES('A sample post');

sqlite> SELECT * FROM posts;
+----+---------------+------------+-------------------+
| id |    content    | word_count | read_time_minutes |
+----+---------------+------------+-------------------+
| 1  | A sample post | 3          | 1                 |
+----+---------------+------------+-------------------+
sqlite> INSERT INTO posts(content) VALUES('A SQLITE Post for the generated column constraint.');

sqlite> SELECT * FROM posts;
+----+----------------------------------------------------+------------+-------------------+
| id |                      content                       | word_count | read_time_minutes |
+----+----------------------------------------------------+------------+-------------------+
| 1  | A sample post                                      | 3          | 1                 |
| 2  | A SQLITE Post for the generated column constraint. | 8          | 1                 |
+----+----------------------------------------------------+------------+-------------------+

sqlite> UPDATE posts SET content = content || 'New content
'  ...> some more content' where id = 2;

sqlite> SELECT * FROM posts;
+----+--------------------------------------------------------------+------------+-------------------+
| id |                           content                            | word_count | read_time_minutes |
+----+--------------------------------------------------------------+------------+-------------------+
| 1  | A sample post                                                | 3          | 1                 |
+----+--------------------------------------------------------------+------------+-------------------+
| 2  | A SQLITE Post for the generated column constraint.New conten | 11         | 1                 |
|    | t                                                            |            |                   |
|    | some more content                                            |            |                   |
+----+--------------------------------------------------------------+------------+-------------------+

sqlite> INSERT INTO posts(content) VALUES('New post');

sqlite> SELECT * FROM posts;
+----+--------------------------------------------------------------+------------+-------------------+
| id |                           content                            | word_count | read_time_minutes |
+----+--------------------------------------------------------------+------------+-------------------+
| 1  | A sample post                                                | 3          | 1                 |
+----+--------------------------------------------------------------+------------+-------------------+
| 2  | A SQLITE Post for the generated column constraint.New conten | 11         | 1                 |
|    | t                                                            |            |                   |
|    | some more content                                            |            |                   |
+----+--------------------------------------------------------------+------------+-------------------+
| 3  | New post                                                     | 2          | 1                 |
+----+--------------------------------------------------------------+------------+-------------------+

sqlite> UPDATE posts SET content = content || replace(printf('%200c', ' '), ' ', ' Adding some extra words here.') WHERE id = 3;

sqlite> SELECT * FROM posts;
+----+--------------------------------------------------------------+------------+-------------------+
| id |                           content                            | word_count | read_time_minutes |
+----+--------------------------------------------------------------+------------+-------------------+
| 1  | A sample post                                                | 3          | 1                 |
+----+--------------------------------------------------------------+------------+-------------------+
| 2  | A SQLITE Post for the generated column constraint.New conten | 11         | 1                 |
|    | t                                                            |            |                   |
|    | some more content                                            |            |                   |
+----+--------------------------------------------------------------+------------+-------------------+
| 3  | New post Adding some extra words here. Adding some extra wor | 1002       | 6                 |
|    | ds here. Adding some extra words here. Adding some extra wor |            |                   |
|    | ds here. Adding some extra words here. Adding some extra wor |            |                   |
|    |..............................................................|            |                   |
|    | ds here. Adding some extra words here. Adding some extra wor |            |                   |
|    | ds here.                                                     |            |                   |
+----+--------------------------------------------------------------+------------+-------------------+

```

## Generated columns cannot be updated manually

If you tried to update the columns with `GENERATED ALWAYS` constraint, you will get a `constraint failed` error. As the constraint itself says, it is **GENERATED ALWAYS**, hence not to be updated or inserted.

```sql
UPDATE posts SET word_count = 10 WHERE id = 1;
```

It will fail with `cannot UPDATE generated column "word_count"` as mentioned.


```sql
UPDATE posts SET read_time_minutes = 10 WHERE id = 1;
```

Similarly, this will also fail, as it is also a `GENERATED ALWAYS` constrained column.

```
sqlite> UPDATE posts SET word_count = 10 WHERE id = 1;
Parse error: cannot UPDATE generated column "word_count"

sqlite> UPDATE posts SET read_time_minutes = 10 WHERE id = 1;
Parse error: cannot UPDATE generated column "read_time_minutes"
```

## STORED VS VIRTUAL 

Now, the Generated columns can be `STORED` as in stored in the database to the disk file, or they could be `VIRTUAL` as in computed (generated,calculated each time). There are trade-offs and I will definitely add the meme here `IT ACTUALLY DEPENDS`!

The `STORED` column will take up space but the queries will be quick, since it doesn't need to calculate it each time, only it needs to do it at each updation (that can't be avoided).
The `VIRTUAL` column will not take up space but the queries will be a little slower, depending on the number of data to compute, since it doesn't store the column values, it will have to compute each time and it as well need to update it each time.

> The default is `VIRTUAL` if not specified.

The commonality of `STORED` and `VIRTUAL` is that they are both updated at the updation of the column it relates to or the updation of the row, the only difference is that one stores it the other calculates it for each query made to it.

We'll see the difference how is that the case, it generates each time in the case of the `VIRTUAL` column `GENERATED` constraint.

I'll show the most basic example continuing with the same table, but with the `VIRTUAL` clause on the `GENERATED ALWAYS` constraint.

### Virtual generated column

By default, it is `VIRTUAL` it doesn't matter if you add it or not (let's add it to make it clear).

I'll specify the `word_count` as the `VIRTUAL` column with the `GENERATED ALWAYS` constraint with the same formula to get the word count and at the end specify the `VIRTUAL` keyword (its optional).

> Let's skip the `read_time_minutes` for now, you can add it won't make it any difference (it was a little extreme example to use 200 words to update its value)

```sql
CREATE TABLE posts_virtual (
   id INTEGER PRIMARY KEY,
   content TEXT,
   word_count INTEGER GENERATED ALWAYS AS (
     length(content) - length(replace(content, ' ', '')) + 1
   ) VIRTUAL
);
```
Now, we'll add some data to the table and check the value of the `word_count` column.

```sql
INSERT INTO posts_virtual(content) VALUES('A sample post');
INSERT INTO posts_virtual(content) VALUES('A SQLITE Post for the virtual generated column constraint.');
SELECT * FROM posts_virtual;
```

This we already know, the plain old `word_count` is `3` and `9` for the two posts respectively.

But let's see how is the `word_count` column is actually stored and how slow or fast it actually is.

```sql
PRAGMA table_info(posts_virtual);
```

HH? There is no `word_count` column here, it is not stored, it is a `VIRTUAL` column.

```
sqlite> PRAGMA table_info(posts_virtual);
+-----+---------+---------+---------+------------+----+
| cid |  name   |  type   | notnull | dflt_value | pk |
+-----+---------+---------+---------+------------+----+
| 0   | id      | INTEGER | 0       |            | 1  |
| 1   | content | TEXT    | 0       |            | 0  |
+-----+---------+---------+---------+------------+----+
sqlite>
```

What about extra info?

```sql
PRAGMA table_xinfo(posts_virtual);
```
Yes, that is indeed there, and its hidden? Nope, its `generated` a special type.
- `0` is not hidden
- `1` is hidden
- `2` is generated

```
sqlite> PRAGMA table_xinfo(posts_virtual);
+-----+------------+---------+---------+------------+----+--------+
| cid |    name    |  type   | notnull | dflt_value | pk | hidden |
+-----+------------+---------+---------+------------+----+--------+
| 0   | id         | INTEGER | 0       |            | 1  | 0      |
| 1   | content    | TEXT    | 0       |            | 0  | 0      |
| 2   | word_count | INTEGER | 0       |            | 0  | 2      |
+-----+------------+---------+---------+------------+----+--------+
sqlite>
```

It's time to benchmark it then, let's bulk insert 100000 posts.


```sql
WITH RECURSIVE cnt(x) AS (
  SELECT 1
  UNION ALL
  SELECT x+1 FROM cnt WHERE x < 100000
)
INSERT INTO posts_virtual(content)
SELECT 'This is a sample post number ' || x || ' with some words repeated multiple times.'
FROM cnt;
```

If we see now the `posts_virtual` table has over 100000 rows.

```sql
SELECT COUNT(*) FROM posts_virtual;
```

We already had `2` records, and we added `100000` more, so the total is `100002`.

Let's check the sum of the `word_count` to be sure it inserted it right.

```sql
SELECT SUM(word_count) FROM posts_virtual;
```

Ok that shows the sum of the `word_count` is `1300012`.

Let's check with the timer on and evaluate the query.

```
sqlite> WITH RECURSIVE cnt(x) AS (
  SELECT 1
  UNION ALL
  SELECT x+1 FROM cnt WHERE x < 100000
)
INSERT INTO posts_virtual(content)
SELECT 'This is a sample post number ' || x || ' with some words repeated multiple times.'
FROM cnt;

sqlite> PRAGMA table_xinfo(posts_virtual);
+-----+------------+---------+---------+------------+----+--------+
| cid |    name    |  type   | notnull | dflt_value | pk | hidden |
+-----+------------+---------+---------+------------+----+--------+
| 0   | id         | INTEGER | 0       |            | 1  | 0      |
| 1   | content    | TEXT    | 0       |            | 0  | 0      |
| 2   | word_count | INTEGER | 0       |            | 0  | 2      |
+-----+------------+---------+---------+------------+----+--------+

sqlite> .timer on

sqlite> SELECT COUNT(*) FROM posts_virtual;
+----------+
| COUNT(*) |
+----------+
| 100002   |
+----------+
Run Time: real 0.001 user 0.000352 sys 0.000103

sqlite> SELECT SUM(word_count) FROM posts_virtual;
+-----------------+
| SUM(word_count) |
+-----------------+
| 1300012         |
+-----------------+
Run Time: real 0.041 user 0.040682 sys 0.000418

sqlite> PRAGMA page_count;
+------------+
| page_count |
+------------+
| 2086       |
+------------+
Run Time: real 0.000 user 0.000000 sys 0.000078

sqlite> PRAGMA page_size;
+-----------+
| page_size |
+-----------+
| 4096      |
+-----------+
Run Time: real 0.000 user 0.000025 sys 0.000006

sqlite> UPDATE posts_virtual SET content = content || ' extra text' WHERE id % 1000 = 0;
Run Time: real 0.010 user 0.009353 sys 0.000161
```

I think we'll leave it here and move onto creating the `STORED` column as a generated value

### Stored generated column

Similar to the `VIRTUAL` column, we'll create the `STORED` column with the `GENERATED ALWAYS` constraint.

This will be the same table named as `posts_stored` but with the `word_count` being a `STORED` column rather than a `VIRTUAL` column.

```sql
CREATE TABLE posts_stored (
   id INTEGER PRIMARY KEY,
   content TEXT,
   word_count INTEGER GENERATED ALWAYS AS (
     length(content) - length(replace(content, ' ', '')) + 1
   ) STORED
);
```
Now, we'll add some data to the table.

```sql
INSERT INTO posts_stored(content) VALUES('A sample post');
INSERT INTO posts_stored(content) VALUES('A SQLITE Post for the virtual generated column constraint.');
SELECT * FROM posts_stored;
```
This is cool, working as expected, we have already seen it, we need to take a look at how this values are stored now.

Let's take a look at the `table_info` for the `posts_stored` table.

```sql
PRAGMA table_info(posts_stored);
```
It still doesn't have the `GENERATED` column `word_count`.

Let's take a look at the `table_xinfo` for the `posts_stored` table.

```sql
PRAGMA table_xinfo(posts_stored);
```

Ok, wiredly it has the `GENERATED ALWAYS` column `word_count` as well but the `hidden` value is `3`.

We knew it was `2` for hidden but `3` is for? You guessed it right, stored and generated column.

Now, let's also populate the `posts_stored` table with some data, bulk inserting the data.

Let's add `100000` rows in it.

```sql
WITH RECURSIVE cnt(x) AS (
  SELECT 1
  UNION ALL
  SELECT x+1 FROM cnt WHERE x < 100000
)
INSERT INTO posts_stored(content)
SELECT 'This is a sample post number ' || x || ' with some words repeated multiple times.'
FROM cnt;
```
This is the same query except its inserting in the `posts_stored` table.

Let's check if it has inserted the right number of rows.

```sql
SELECT COUNT(*) FROM posts_stored;
```

Indeed there are `100000` rows in the `posts_stored` table.

Now, let's verify that the inserted values are rightly populated.

```sql
SELECT SUM(word_count) FROM posts_stored;
```

This gives `1300012` as the sum of the `word_count` from the `posts_stored` table. Which matches the sum of the `word_count` from the `posts_virtual` table.

```sql
sqlite> CREATE TABLE posts_stored (
   id INTEGER PRIMARY KEY,
   content TEXT,
   word_count INTEGER GENERATED ALWAYS AS (
     length(content) - length(replace(content, ' ', '')) + 1
   ) STORED
);
Run Time: real 0.000 user 0.000258 sys 0.000027

sqlite> INSERT INTO posts_stored(content) VALUES('A sample post');
sqlite> INSERT INTO posts_stored(content) VALUES('A SQLITE Post for the virtual generated column constraint.');

SELECT * FROM posts_stored;
+----+------------------------------------------------------------+------------+
| id |                          content                           | word_count |
+----+------------------------------------------------------------+------------+
| 1  | A sample post                                              | 3          |
| 2  | A SQLITE Post for the virtual generated column constraint. | 9          |
+----+------------------------------------------------------------+------------+
Run Time: real 0.000 user 0.000224 sys 0.000000

sqlite> PRAGMA table_info(posts_stored);
+-----+---------+---------+---------+------------+----+
| cid |  name   |  type   | notnull | dflt_value | pk |
+-----+---------+---------+---------+------------+----+
| 0   | id      | INTEGER | 0       |            | 1  |
| 1   | content | TEXT    | 0       |            | 0  |
+-----+---------+---------+---------+------------+----+

sqlite> PRAGMA table_xinfo(posts_stored);
+-----+------------+---------+---------+------------+----+--------+
| cid |    name    |  type   | notnull | dflt_value | pk | hidden |
+-----+------------+---------+---------+------------+----+--------+
| 0   | id         | INTEGER | 0       |            | 1  | 0      |
| 1   | content    | TEXT    | 0       |            | 0  | 0      |
| 2   | word_count | INTEGER | 0       |            | 0  | 3      |
+-----+------------+---------+---------+------------+----+--------+
Run Time: real 0.000 user 0.000013 sys 0.000322

sqlite> SELECT COUNT(*) FROM posts_stored;
+----------+
| COUNT(*) |
+----------+
| 2        |
+----------+
Run Time: real 0.000 user 0.000108 sys 0.000012

sqlite> WITH RECURSIVE cnt(x) AS (
  SELECT 1
  UNION ALL
  SELECT x+1 FROM cnt WHERE x < 100000
)
INSERT INTO posts_stored(content)
SELECT 'This is a sample post number ' || x || ' with some words repeated multiple times.'
FROM cnt;
Run Time: real 0.071 user 0.067858 sys 0.002929

sqlite> SELECT COUNT(*) FROM posts_stored;
+----------+
| COUNT(*) |
+----------+
| 100002   |
+----------+
Run Time: real 0.000 user 0.000367 sys 0.000000

sqlite> SELECT SUM(word_count) FROM posts_stored;
+-----------------+
| SUM(word_count) |
+-----------------+
| 1300012         |
+-----------------+
Run Time: real 0.008 user 0.006492 sys 0.000941
```

This is all setup for the comparison of both the tables.

### The difference

Now, let's check the difference between the `STORED` and `VIRTUAL` column.

Let's sum the `word_count` from each of the tables `posts_stored` and `posts_virtual`

First, lets set the timer on, which will give the breakdown of the time for each ran query.

```
.timer on
```
The timer will give three time durations, the real, user and the system.
The breakdown is like this:
- real time : The full clock time form the start to the end of the query
- user time : The CPU time in the user space (I/O) operations on the sqlite side, computations on the obtained data, or preprocessing the data before inserting.
- system time: The CPU time in the kernel space (I/O) operations on the host side, it could be reading or writing to the actual db file, etc.

And then run the query for the `posts_stored` table.

```sql
SELECT SUM(word_count) FROM posts_stored;
```

This is giving the result in around `~0.008` seconds might be approximately `0.01` in certain cases.
We can even get the breakdown of the time for each ran query.
- full query execution time: `0.008`
- user time: `0.006492`
- system time: `0.000941`

So, here in this case, the system and user time are around the same, however there is some system time for reading from the database file.

And then run the query for the `posts_virtual` table.

```sql
SELECT SUM(word_count) FROM posts_virtual;
```

This is giving the result in around `~0.039` seconds might be approximately `0.04` in certain cases.

We can even get the breakdown of the time for each ran query.
- full query execution time: `0.039`
- user time: `0.039366`
- system time: `0.000000` this can go sometimes to `0.000652`, almost negligible, but if it happens, in that case the full query time will increase by a bit.

So, it is taking a quite lot of time in the user space by the CPU right?
And almost 0 time in the kernel level (disk operation) to read from the database file.

That is so clear, as this is a virtual column, and we are only reading a single column, and that too is a virtual column, we don't need to hit a read to the database file. Clever right? But that costs some CPU to compute the sum for the `100002` rows. 

This is the trade-off

- Time spent on the CPU on the computation vs the time spent on the disk operation to read the data.
- These can be different depending on the query.

```sql
sqlite> SELECT SUM(word_count) FROM posts_stored;
+-----------------+
| SUM(word_count) |
+-----------------+
| 1300012         |
+-----------------+
Run Time: real 0.008 user 0.006492 sys 0.000941

sqlite> SELECT SUM(word_count) FROM posts_virtual;
+-----------------+
| SUM(word_count) |
+-----------------+
| 1300012         |
+-----------------+
Run Time: real 0.039 user 0.039366 sys 0.000000
```

You can clearly see the difference, the `STORED` column is almost `4` to `5` times faster than the `VIRTUAL` column.

### Key observations

1. If you have lot of processing to do in the generated column, better to go with the `STORED` column.
2. If you have a lot of data to compute and not necessarily complex, better to go with the `VIRTUAL` column.
3. It really depends on the data too, and we cannot say which one is better over the other without knowing the structure and the design of the queries.


There are lot of details as well beyond this, it will require specific guides on how to do certain things and busting certain myths in the documentation.
