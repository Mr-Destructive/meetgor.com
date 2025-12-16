---
type: "sqlog"
title: "Advent of SQL 2025: Wish List"
slug: "ao-sql-2025-day-01-wish-list"
date: 2025-12-16
series: ["advent-of-sql-2025"]
tags: ["sqlite", "sql", "advent-of-sql"]
---

## Learning SQLite: Advent of SQL Day 1

I am trying to learn SQLite, I want to understand that database. It's quite simple yet the whole world uses it for various kinds of things ranging from developers' toy database to spaceships. What a tiny engineering marvel!

I am happy to see this happening: [Advent of SQL](https://databaseschool.com/series/advent-of-sql/videos/308)

What a better time to learn more. I guess I want to start by exploring all the specificities of the INSERT statement in SQLite after exploring most of the things of the CREATE TABLE statement.

But here I am jumping to this. Why? Because I want to solve something first before exploring other branches.

Today I am going to try to solve the day 1 part.

> **NOTE:** I would be using my local sqlite db for this or a playground on this for testing. I am not going to use the browser-based playground attached in the databaseschool.com app for a reason. I want to use SQLite. The database is some form of Postgres; I don't mind using it, but I want to do it in SQLite.

I have a playground on my blog for SQLite, you can try it out here:

```sql
SELECT 1;
```

It uses an embedded SQLite version (3.49.1) with sql-js as a wasm extension: [sql.js v1.13.0](https://github.com/sql-js/sql.js/releases/tag/v1.13.0)

Back to the problem elves!

## Setup

This is the first day, so advent calendar usually requires some setup or preparation for the rest of the days. Luckily it's optional for you if you are doing it in the playground of database school or in PostgreSQL Database.

We have some .sql files as input for creation and insertions of tables and rows in the database. It's for constructing the schema (tables) and populating the rows that the problem requires us to do.

The SQL looked something simpler like:

```sql
DROP TABLE IF EXISTS wish_list CASCADE;

CREATE TABLE wish_list (
   id          BIGSERIAL PRIMARY KEY,
   child_name  TEXT,
   raw_wish    TEXT
);

INSERT INTO wish_list (id, child_name, raw_wish) VALUES (1, 'James A.', ' BLUEY SUPERMARKET PLAY SET');
INSERT INTO wish_list (id, child_name, raw_wish) VALUES (2, 'Sade C.', 'lego star wars set ');
```

There are around 499,000 rows!

However when I tried to read directly into a SQLite shell as:

```
.read day1-wish-list.sql
```

It got an error for parsing the DROP TABLE statement:

```
sqlite> .read day1-wish-list.sql
Parse error near line 11: near "CASCADE": syntax error
  DROP TABLE IF EXISTS wish_list CASCADE;
                       error here ---^
```

Obviously it was designed for Postgres. It won't work in SQLite.

SQLite is minimal. It might not have everything that PostgreSQL has, but PostgreSQL might have everything that SQLite has (maybe but not as is).

So, we need to remove the CASCADE, which is an option to decide what to do with the related data rows when a relation is removed. In this case, it is cascading—deleting all the other related data points or records in the related tables. SQLite doesn't have options to modify the relations for the DROP TABLE statement. It has it for CREATE TABLE with the foreign key constraint.

Now we need to remove it. It can't be in the DROP TABLE statement for SQLite database:

```sql
DROP TABLE IF EXISTS wish_list;
```

Now, let's check by running the queries again:

```
.read day1-wish-list.sql
```

That works!

```
$ sqlite3
SQLite version 3.45.1 2024-01-30 16:01:20
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
sqlite> .read day1-wish-list_sqlite1.sql
sqlite> .schema
CREATE TABLE wish_list (
   id          BIGSERIAL PRIMARY KEY,
   child_name  TEXT,
   raw_wish    TEXT
);
```

But this looks weird:

```sql
CREATE TABLE wish_list (
   id          BIGSERIAL PRIMARY KEY,
   child_name  TEXT,
   raw_wish    TEXT
);
```

BIGSERIAL is not a datatype in SQLite. It might be in PostgreSQL. But does it matter? In SQLite, if the table is not STRICT, it doesn't matter what type the column is or even if it is NOT specified of any type. That's fine.

But I want to know what is the type of `wish_list.id`?

```sql
SELECT distinct(typeof(id)) FROM wish_list;
```

Here's the output:

```
sqlite> SELECT distinct(typeof(id)) FROM wish_list;
integer
```

It is integer, because of type affinity, I guess. Good work SQLite.

Looks like the data is fine.

Now to the problem.

## Problem

[Link to the challenge](https://databaseschool.com/series/advent-of-sql/videos/309#:~:text=Challenge%3A%20Using%20the,and%20get%20started!)

**Challenge:** Using the wish_list table, count how many times each cleaned toy name appears, from most requested to least requested. Return the results in two columns: wish and count. Make sure the wish results have no extra leading or trailing spaces and are all lowercase.

So, simply we need two columns:
- **wish** (the text)
- **count** (the number of times that wish is wished)

### Subtleties

> Some children had typed extra spaces. Some wrote in ALL CAPS. Some had letters that danced between cases like playful snowflakes.
> 
> I know there are some issues with spelling, the extra spaces, or the funny capitalization, but I just need to know what the children truly meant.

So, we need to either use lower caps or upper caps the wish text and trim off the space.

SCALAR FUNCTIONS!!

### Scalar Functions

[SQLite Core Functions](https://sqlite.org/lang_corefunc.html)

I read through the list of around ~70 of them, most of them are kind of the same with different parameters.

The one that I found relevant are:
- [LOWER](https://sqlite.org/lang_corefunc.html#lower)
- [TRIM](https://sqlite.org/lang_corefunc.html#trim)

That's it, right? Convert into LOWER (or UPPER) and TRIM off the spaces.

```sql
SELECT LOWER(TRIM(wish)) FROM wish_list;
```

Don't run just yet!

```sql
SELECT LOWER(TRIM(wish)) FROM wish_list LIMIT 100;
```

Looks good. Now to the next step.

### Grouping and Counting

We need to count them i.e. to group by the wish.

> GROUP BY: What group by does is that it condenses the rows of certain column into a singular column for instance if there are 10 entries for "lego star wars set" adding a group by wish will create a singular entry for that wish and we can then perform operations like sum, count, average and all of that on other rows

```sql
SELECT LOWER(TRIM(wish)) AS wish, count(*) AS count
FROM wish_list
GROUP BY wish;
```

Here we are grouping by wish because we don't want 10 entries of "lego star wars set" we just want one common entry to view the unique wishes.

Also by using `COUNT(*)` we are counting the instances of each row. As I said, the multiple rows with the same wish are squished into a single row. So now we can use aggregate functions like count, sum, in our case we want to count how many instances of those particular wish are.


### Ordering Results

Does that solve it? Mostly, just need the ORDER BY now.

Because we also need to order the results:

```sql
SELECT LOWER(TRIM(wish)) AS wish, count(*) AS count
FROM wish_list
GROUP BY wish
ORDER BY count DESC;
```

Perfect? Probably.

> What ORDER BY does is that it just determines which way the row should be aligned based on what column and how i.e. the column name and either ASC(ending) or DESC(ending).


Here we have ordered by count so that we can filter the most wished toy or least wished toy at the top and increasing or decreasing order of it.

### Results

Now with the mode table:

```
sqlite> .mode table
sqlite> SELECT LOWER(TRIM(raw_wish)) as wish, count(*) as count FROM wish_list group by wish order by count desc;
+-----------------------------+-------+
|            wish             | count |
+-----------------------------+-------+
| lego city f1 car            | 32893 |
| barbie dreamhouse           | 32785 |
| nerf blaster                | 32746 |
| lego star wars set          | 32611 |
| beyblade battle arena       | 29564 |
| magna-tiles pet playhouse   | 29529 |
| bluey supermarket play set  | 26292 |
| lego friends amusement park | 25982 |
| pokemon trainer box         | 25968 |
| duplo building set          | 23005 |
| mini brands fill the fridge | 22965 |
| electric toy train set      | 22885 |
| toniebox audio player       | 19529 |
| scooter                     | 19496 |
| vr headset                  | 16468 |
| squishmallows               | 16304 |
| shaved ice machine          | 16263 |
| drone for kids              | 13151 |
| coding robot                | 13025 |
| headphones                  | 13006 |
| interactive robot dog       | 9770  |
| fidget spinner              | 3590  |
| yo-yo                       | 3565  |
| slime kit                   | 3553  |
| littlest pet shop playset   | 3543  |
| chatter telephone           | 3527  |
| fingerlings robot monkey    | 3511  |
| rubik's revolution          | 3474  |
+-----------------------------+-------+
sqlite>
```

---

Day 1 done, moving on to day two by helping those pesky elves tomorrow. I am amazed at how stupider problems humans create with those elves as target. Just kidding, humans are elves :)

Happy Coding :)

Happy Squealing
