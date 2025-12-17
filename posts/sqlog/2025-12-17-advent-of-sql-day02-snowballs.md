---
type: "sqlog"
title: "Advent of SQL 2025 Day 2: Snowballs"
slug: "advent-of-sql-2025-day-2"
date: 2025-12-17
series: ["advent-of-sql-2025"]
tags: ["sqlite", "sql", "advent-of-sql"]
---

## SQLog: Advent of SQL Day 2

Here we are on the day 2 of Advent of SQL

As I said in the previous day this is in SQLite so I won't be doing it in the playground. So here is your SQLite playground :)

```sql
SELECT 1;
```

From now on no setup straight to the problem!

Let's download the .sql file for today's problem to see what data we are playing with.

```sql
DROP TABLE IF EXISTS snowball_categories;
DROP TABLE IF EXISTS snowball_inventory;

CREATE TABLE snowball_categories (
    id INT PRIMARY KEY,
    official_category TEXT
);

CREATE TABLE snowball_inventory (
    id INT PRIMARY KEY,
    batch_id TEXT,
    category_name TEXT,
    quantity INT,
    status TEXT
);
```

Well, neat and clean!

Straight pulling it in the sqlite shell.

```plaintext
$ sqlite3 
SQLite version 3.45.1 2024-01-30 16:01:20
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
sqlite> .read day2-inserts.sql
sqlite> .schema
CREATE TABLE snowball_categories (
    id INT PRIMARY KEY,
    official_category TEXT
);
CREATE TABLE snowball_inventory (
    id INT PRIMARY KEY,
    batch_id TEXT,
    category_name TEXT,
    quantity INT,
    status TEXT
);
sqlite> 
```

Worked right off the bat.

Straight to problem now.

If you are working in the playground, do add this full code with inserts in order to get a sense of what the data looks like, note that this below statements is not all data, just sharing to get you a taste of the problem.

```sql
DROP TABLE IF EXISTS snowball_categories;
DROP TABLE IF EXISTS snowball_inventory;

CREATE TABLE snowball_categories (
    id INT PRIMARY KEY,
    official_category TEXT
);

CREATE TABLE snowball_inventory (
    id INT PRIMARY KEY,
    batch_id TEXT,
    category_name TEXT,
    quantity INT,
    status TEXT
);

INSERT INTO snowball_categories (id, official_category) VALUES
    (1, 'frost-flight deluxe'),
    (2, 'north ridge compact'),
    (3, 'glacier sphere (xl)'),
    (4, 'polar precision microball'),
    (5, 'everfrost training round'),
    (6, 'arctic blast premium');

INSERT INTO snowball_inventory (id, batch_id, category_name, quantity, status) VALUES
    (1, 'BATCH-35443-J', 'frost-flight deluxe', 19, NULL),
    (2, 'BATCH-59767-M', 'frost-flight deluxe', 41, 'incomplete'),
    (3, 'BATCH-44795-B', 'frost-flight deluxe', 21, 'ready'),
    (4, 'BATCH-23396-C', 'north ridge compact', 0, 'incomplete'),
    (5, 'BATCH-88907-A', 'frost-flight deluxe', -2, 'incomplete'),
    (6, 'BATCH-42662-D', 'frost-flight deluxe', 47, 'needs review'),
    (7, 'BATCH-37460-V', 'north ridge compact', 43, 'ready'),
    (8, 'BATCH-21395-S', 'frost-flight deluxe', -2, 'ready'),
    (9, 'BATCH-36100-E', 'frost-flight deluxe', 46, 'ready'),
    (10, 'BATCH-64987-H', 'frost-flight deluxe', 43, NULL),
    (11, 'BATCH-57576-Z', 'melty deluxe', -5, 'ready'),
    (12, 'BATCH-56025-U', 'snowball v2', 11, 'ready'),
    (13, 'BATCH-86556-W', 'snowball v2', 12, 'ready'),
    (14, 'BATCH-83385-N', 'frost-flight deluxe', 38, 'incomplete'),
    (15, 'BATCH-85156-M', 'prototype x-12', 28, 'incomplete'),
    (16, 'BATCH-82135-F', 'north ridge compact', 32, 'incomplete'),
    (17, 'BATCH-10074-T', 'frost-flight deluxe', 49, 'needs review'),
    (18, 'BATCH-22676-L', 'frost-flight deluxe', 16, 'incomplete'),
    (19, 'BATCH-31174-R', 'north ridge compact', 33, 'incomplete'),
    (20, 'BATCH-41385-B', 'frost-flight deluxe', 4, 'ready'),
    (21, 'BATCH-50404-L', 'frost-flight deluxe', -4, 'needs review'),
    (22, 'BATCH-92240-F', 'north ridge compact', 20, 'ready'),
    (23, 'BATCH-29198-J', 'beta test sphere', 0, 'incomplete'),
    (24, 'BATCH-64987-H', 'glacier sphere (xl)', 18, 'needs review'),
    (25, 'BATCH-80008-A', 'frost-flight deluxe', 3, 'incomplete'),
    (26, 'BATCH-88907-A', 'polar precision microball', 48, 'incomplete'),
    (27, 'BATCH-55830-J', 'north ridge compact', 0, 'needs review'),
    (28, 'BATCH-69470-A', 'frost-flight deluxe', -3, 'incomplete'),
    (29, 'BATCH-46211-R', 'frost-flight deluxe', -3, 'ready'),
    (30, 'BATCH-18675-G', 'glacier sphere (xl)', -1, 'ready');
```

## Problem

Let's check the challenge of day 2

> Using the `snowball_inventory` and `snowball_categories` tables, write a query that returns valid snowball categories with the count of valid snowballs per category. Your final table should have the columns `official_category` and `total_usable_snowballs`. Sort the output from fewest to most `total_usable_snowballs`.

So, we have two tables:

1. snowball categories
    
2. snowball inventory
    

The snowball categories looks a quite a small table with just the category name and the id which is not really the data, the name is.

```sql
SELECT * FROM snowball_categories;
```

Just 6 entries with some names of the categories, the column is official categories.

```sql
sqlite> SELECT * FROM snowball_categories;
+----+---------------------------+
| id |     official_category     |
+----+---------------------------+
| 1  | frost-flight deluxe       |
| 2  | north ridge compact       |
| 3  | glacier sphere (xl)       |
| 4  | polar precision microball |
| 5  | everfrost training round  |
| 6  | arctic blast premium      |
+----+---------------------------+
sqlite> 
```

And the other table has oddles of data:

```sql
SELECT * FROM snowball_inventory limit 10;
```

So, we have quite a few columns,

1. batch\_id
    
2. category\_name
    
3. quantity
    
4. status
    

It has 200,000 rows, that's quite a lot.

```plaintext
sqlite> SELECT count(*) FROM snowball_inventory;
+----------+
| count(*) |
+----------+
| 200000   |
+----------+

sqlite> SELECT * FROM snowball_inventory limit 10;
+----+---------------+---------------------+----------+--------------+
| id |   batch_id    |    category_name    | quantity |    status    |
+----+---------------+---------------------+----------+--------------+
| 1  | BATCH-35443-J | frost-flight deluxe | 19       |              |
| 2  | BATCH-59767-M | frost-flight deluxe | 41       | incomplete   |
| 3  | BATCH-44795-B | frost-flight deluxe | 21       | ready        |
| 4  | BATCH-23396-C | north ridge compact | 0        | incomplete   |
| 5  | BATCH-88907-A | frost-flight deluxe | -2       | incomplete   |
| 6  | BATCH-42662-D | frost-flight deluxe | 47       | needs review |
| 7  | BATCH-37460-V | north ridge compact | 43       | ready        |
| 8  | BATCH-21395-S | frost-flight deluxe | -2       | ready        |
| 9  | BATCH-36100-E | frost-flight deluxe | 46       | ready        |
| 10 | BATCH-64987-H | frost-flight deluxe | 43       |              |
+----+---------------+---------------------+----------+--------------+
sqlite> 
```

What we have to do?

> return valid snowball categories with the count of valid snowballs per category

That's a mouthful.

The problem is trying to get at counting the number of inventory items per category I think.

Because there are quite less categories and tons of inventory records.

Also the `category_name` in the `snowball_inventory` table is not trustworthy.

> Santa hurried to the snowball storage board, but the situation only got stranger. Whole batches appeared twice. Some batches claimed they had negative snowballs (“a bookkeeping accident,” the elves muttered). Others had a quantity of zero but were still marked “Ready.” And many batches referenced categories that didn’t appear anywhere in the official Snowball Category Guide

## Naive Approach

The first approach I see is scan through all the inventory records and check if the category is in the `snowball_category` table and its quantity is positive or more than 0, also I wonder if the status needs to be checked as `ready`.

> “We need to know what we actually have left,” Santa said. “Not puddles. Not phantom batches. Real, usable, throw-ready snowballs.”

**Real, throw-ready snowballs**

Let's check how to do that with simple sub query.

```sql
SELECT * FROM snowball_inventory
WHERE
   category_name IN (SELECT official_category FROM snowball_categories)
   AND quantity > 0
   AND status == 'ready';
```

OK!

What we did? Simply selected all the data from `snowball_inventory` table,

in which the category name is matching either of the 6 categories in the `snowball_categories` table.

Also the quantity is positive and not zero, also the status is set to be `ready`.

This looks naive to me for using subquery because for each 200,000 records we will scan the `snowball_categories` . Ew!

Wait its done! We need the count of each category! We need to group by the `category_name`

```sql
SELECT * FROM snowball_inventory
WHERE
   category_name IN (SELECT official_category FROM snowball_categories)
   AND quantity > 0
   AND status == 'ready'
GROUP BY category_name;
```

```sql
sqlite> SELECT * FROM snowball_inventory
WHERE
   category_name IN (SELECT official_category FROM snowball_categories)
   AND quantity > 0
   AND status == 'ready'
GROUP BY category_name;
+-----+---------------+---------------------------+----------+--------+
| id  |   batch_id    |       category_name       | quantity | status |
+-----+---------------+---------------------------+----------+--------+
| 240 | BATCH-55793-L | arctic blast premium      | 45       | ready  |
| 163 | BATCH-75333-O | everfrost training round  | 23       | ready  |
| 3   | BATCH-44795-B | frost-flight deluxe       | 21       | ready  |
| 39  | BATCH-22704-V | glacier sphere (xl)       | 37       | ready  |
| 7   | BATCH-37460-V | north ridge compact       | 43       | ready  |
| 125 | BATCH-81987-E | polar precision microball | 47       | ready  |
+-----+---------------+---------------------------+----------+--------+
```

That sort of looks wired right?

Why?

Because what happens to the quantity? Is that summed? averaged, or what just happened to the batch\_id?

No it takes one value out of the `200 000` rows for each category. That's not what we want right?

We want this

> return valid snowball categories with the count of valid snowballs per category

So, we just want the category name and the count of those category. Basically count per category.

```sql
SELECT category_name as official_category, sum(quantity) as total_usable_snowballs
FROM snowball_inventory 
WHERE
   category_name IN (SELECT official_category FROM snowball_categories)
   AND quantity > 0
   AND status == 'ready'
GROUP BY category_name;
```

Here we are only fetching the `category_name` and the `count` which will bundle up all the counts from the grouped of the category\_name.

```plaintext
sqlite> SELECT category_name as official_category, sum(quantity) as total_usable_snowballs 
FROM snowball_inventory 
WHERE
   category_name IN (SELECT official_category FROM snowball_categories)
   AND quantity > 0
   AND status == 'ready'
GROUP BY category_name;
+---------------------------+------------------------+
|     official_category     | total_usable_snowballs |
+---------------------------+------------------------+
| arctic blast premium      | 11470                  |
| everfrost training round  | 24248                  |
| frost-flight deluxe       | 952019                 |
| glacier sphere (xl)       | 165158                 |
| north ridge compact       | 594119                 |
| polar precision microball | 70773                  |
+---------------------------+------------------------+
sqlite> 

```

Not done yet!

We need to order by the count.

> Sort the output from fewest to most `total_usable_snowballs`.

```sql
SELECT category_name, sum(quantity)  as total_usable_snowballs
FROM snowball_inventory
WHERE
   category_name IN (SELECT official_category FROM snowball_categories)
   AND quantity > 0
   AND status == 'ready'
GROUP BY category_name
ORDER BY total_usable_snowballs;
```

We can provide `ORDER BY total_usable_snowballs ASC` but ascending is default. I prefer keeping things default, you can be explicit and mention it as `ASC` to make it clear and readable.

```plaintext
sqlite> SELECT category_name as official_category, sum(quantity) as total_usable_snowballs
FROM snowball_inventory 
WHERE
   category_name IN (SELECT official_category FROM snowball_categories)
   AND quantity > 0
   AND status == 'ready'
GROUP BY category_name
ORDER BY total_usable_snowballs;
+---------------------------+------------------------+
|     official_category     | total_usable_snowballs |
+---------------------------+------------------------+
| arctic blast premium      | 11470                  |
| everfrost training round  | 24248                  |
| polar precision microball | 70773                  |
| glacier sphere (xl)       | 165158                 |
| north ridge compact       | 594119                 |
| frost-flight deluxe       | 952019                 |
+---------------------------+------------------------+

sqlite> 
```

There it is!

The solution to the day 2.

But Santa, I am using subqueries. Is that fine?

## Joins?

We can use joins here, since we just require the count of each category.

Which JOIN though?

LEFT, RIGHT, INNER?

any really?

You don't choose the join based on the problem, you define your output columns and then choose the type of join that would give you the result.

If the thing that is to be searched from left to right and then its a left join, i.e. join everything from the left. and so on.

Here,

I need all the `official_category` column right, which is in the `snowball_categories` table.

If I assume the `snowball_categories` is on the left, I can join everything for that row in the left to match all the `category_name` rows in the `snowball_inventory` table which would be in the right. Like so.

```sql
SELECT 
    snowball_categories.official_category as official_category,
    SUM(snowball_inventory.quantity) as total_usable_snowballs
FROM snowball_categories
LEFT JOIN snowball_inventory
     ON snowball_categories.official_category = snowball_inventory.category_name
     AND snowball_inventory.quantity > 0
     AND snowball_inventory.status == 'ready'
GROUP BY official_category
ORDER BY total_usable_snowballs;
```

We JOIN on the condition of `snowball_categories.official_category = snowball_inventory.category_name` which is to say, the `category_name` in the `snowball_inventory` table should match the `official category` from the `snowball_categories` table. Also the other conditions like `quantity` should be more than `0` and the `status` should be `ready`.

We will still need to group by and order by as the joins will basically be the filtering criteria in which to select or reject the bad categories. However to group by the category name and obtain the count of each category we need to group by the name of the category.

```plaintext
sqlite> SELECT snowball_categories.official_category AS official_category, SUM(snowball_inventory.quantity) AS total_usable_snowballs
FROM snowball_categories
LEFT JOIN snowball_inventory
     ON snowball_categories.official_category = snowball_inventory.category_name
     AND snowball_inventory.quantity > 0
     AND snowball_inventory.status == 'ready'
GROUP BY official_category
ORDER BY total_usable_snowballs;
+---------------------------+------------------------+
|     official_category     | total_usable_snowballs |
+---------------------------+------------------------+
| arctic blast premium      | 11470                  |
| everfrost training round  | 24248                  |
| polar precision microball | 70773                  |
| glacier sphere (xl)       | 165158                 |
| north ridge compact       | 594119                 |
| frost-flight deluxe       | 952019                 |
+---------------------------+------------------------+
```

This is it! The same thing that we did for the subquery. Just the filtering part is different.

How about INNER JOIN?

`INNER JOIN` is different from `LEFT JOIN` as it will by default take only rows that match the condition, whereas `LEFT JOIN` will include records/rows from the `LEFT` table even if they are empty or `NULL` . Similar is `RIGHT JOIN`, it will include everything from the RIGHT table even if that is `NULL` so that all the records from the `RIGHT` table are in the result set.

For INNER JOIN if we flip the tables end it won't matter, as it only relies on the condition and not on the order of where the tables are placed (left or right).

```sql
SELECT
  snowball_categories.official_category AS official_category,
  SUM(snowball_inventory.quantity) AS total_usable_snowballs
FROM snowball_inventory
JOIN snowball_categories
  ON snowball_categories.official_category = snowball_inventory.category_name
  AND snowball_inventory.quantity > 0
  AND snowball_inventory.status = 'ready'
GROUP BY official_category
ORDER BY total_usable_snowballs;
```

This is with INNER JOIN nothing changed here? did it?

except the word `LEFT` is no longer there, the default join is `INNER JOIN` .

In our case i don't think we have `NULL` values in the `snowball_categories` table or `snowball_inventory` table with the category name columns. So LEFT AND INNER JOIN won't be different in terms of the results.

```plaintext
sqlite> SELECT
  snowball_categories.official_category AS official_category,
  SUM(snowball_inventory.quantity) AS total_usable_snowballs
FROM snowball_inventory
JOIN snowball_categories
  ON snowball_categories.official_category = snowball_inventory.category_name
  AND snowball_inventory.quantity > 0
  AND snowball_inventory.status = 'ready'
GROUP BY official_category
ORDER BY total_usable_snowballs;
+---------------------------+------------------------+
|     official_category     | total_usable_snowballs |
+---------------------------+------------------------+
| arctic blast premium      | 11470                  |
| everfrost training round  | 24248                  |
| polar precision microball | 70773                  |
| glacier sphere (xl)       | 165158                 |
| north ridge compact       | 594119                 |
| frost-flight deluxe       | 952019                 |
+---------------------------+------------------------+
sqlite> 
```

However if we decide to make the `snowball_inventory` table as the `Left` table then does it make a difference?

Well we don't want that to be the left table if we are doing a `LEFT` JOIN as it might list all the wrong categories in the result list as well even if they are not in the `official_category` those will be 0 or NULL but that would make a mess.

We instead can do a RIGHT JOIN.

With the tables order as

1. snowball\_inventory
    
2. snowball\_categories
    

And we will RIGHT JOIN that is select all the columns from the right to be in the result set even if the condition is not true. It will be NULL. But the case is not true for here, as there are only 6 non-NULL category names in the snowball\_categories table.

```sql

SELECT                                                       
  snowball_categories.official_category AS official_category,
  SUM(snowball_inventory.quantity) AS total_usable_snowballs
FROM snowball_inventory
RIGHT JOIN snowball_categories
  ON snowball_categories.official_category = snowball_inventory.category_name
  AND snowball_inventory.quantity > 0
  AND snowball_inventory.status = 'ready'
GROUP BY official_category
ORDER BY total_usable_snowballs;
```

```plaintext
sqlite> SELECT                                                       
  snowball_categories.official_category AS official_category,
  SUM(snowball_inventory.quantity) AS total_usable_snowballs
FROM snowball_inventory
RIGHT JOIN snowball_categories
  ON snowball_categories.official_category = snowball_inventory.category_name
  AND snowball_inventory.quantity > 0
  AND snowball_inventory.status = 'ready'
GROUP BY official_category
ORDER BY total_usable_snowballs;
+---------------------------+------------------------+
|     official_category     | total_usable_snowballs |
+---------------------------+------------------------+
| arctic blast premium      | 11470                  |
| everfrost training round  | 24248                  |
| polar precision microball | 70773                  |
| glacier sphere (xl)       | 165158                 |
| north ridge compact       | 594119                 |
| frost-flight deluxe       | 952019                 |
+---------------------------+------------------------+
```

However if we try to use `LEFT JOIN` with `snowball_inventory` as the `LEFT` table, we might get something wired.

```sql
SELECT                                                
  snowball_categories.official_category AS official_category,
  SUM(snowball_inventory.quantity) AS total_usable_snowballs
FROM snowball_inventory
LEFT JOIN snowball_categories
  ON snowball_categories.official_category = snowball_inventory.category_name
  AND snowball_inventory.quantity > 0
  AND snowball_inventory.status = 'ready'
GROUP BY official_category
ORDER BY total_usable_snowballs;
```

We have 6 and 1 more category as empty ?

That category is for all the rows that didn't met the criteria of JOIN but still as asked to `LEFT JOIN` it will try to show the results for all the rows in the left table which is the `snowball_inventory` table.

```plaintext

sqlite> SELECT                                                
  snowball_categories.official_category AS official_category,
  SUM(snowball_inventory.quantity) AS total_usable_snowballs
FROM snowball_inventory
LEFT JOIN snowball_categories
  ON snowball_categories.official_category = snowball_inventory.category_name
  AND snowball_inventory.quantity > 0
  AND snowball_inventory.status = 'ready'
GROUP BY official_category
ORDER BY total_usable_snowballs;
+---------------------------+------------------------+
|     official_category     | total_usable_snowballs |
+---------------------------+------------------------+
| arctic blast premium      | 11470                  |
| everfrost training round  | 24248                  |
| polar precision microball | 70773                  |
| glacier sphere (xl)       | 165158                 |
| north ridge compact       | 594119                 |
| frost-flight deluxe       | 952019                 |
|                           | 1692814                |
+---------------------------+------------------------+
```

## Timer

Phew! Let's see which one of them is the best one

Let's see the timings:

```plaintext
.timer on
```

### SUBQUERY

```plaintext
sqlite> SELECT category_name, SUM(quantity)AS total_usable_snowballs
FROM snowball_inventory
WHERE
   category_name IN (SELECT official_category FROM snowball_categories)
   AND quantity > 0
   AND status == 'ready'
GROUP BY category_name
ORDER BY total_usable_snowballs;
+---------------------------+------------------------+
|       category_name       | total_usable_snowballs |
+---------------------------+------------------------+
| arctic blast premium      | 11470                  |
| everfrost training round  | 24248                  |
| polar precision microball | 70773                  |
| glacier sphere (xl)       | 165158                 |
| north ridge compact       | 594119                 |
| frost-flight deluxe       | 952019                 |
+---------------------------+------------------------+
Run Time: real 0.077 user 0.073886 sys 0.002942

sqlite> 
```

### LEFT JOIN

```plaintext
sqlite>  SELECT snowball_categories.official_category AS official_category, SUM(snowball_inventory.quantity) AS total_usable_snowballs
FROM snowball_categories
LEFT JOIN snowball_inventory
     ON snowball_categories.official_category = snowball_inventory.category_name 
     AND snowball_inventory.quantity > 0
     AND snowball_inventory.status == 'ready'
GROUP BY official_category
ORDER BY total_usable_snowballs;
+---------------------------+------------------------+
|     official_category     | total_usable_snowballs |
+---------------------------+------------------------+
| arctic blast premium      | 11470                  |
| everfrost training round  | 24248                  |
| polar precision microball | 70773                  |
| glacier sphere (xl)       | 165158                 |
| north ridge compact       | 594119                 |
| frost-flight deluxe       | 952019                 |
+---------------------------+------------------------+
Run Time: real 0.182 user 0.171433 sys 0.009907
sqlite> 
```

### INNER JOIN

```plaintext
sqlite>  SELECT
  snowball_categories.official_category AS official_category,
  SUM(snowball_inventory.quantity) AS total_usable_snowballs
FROM snowball_inventory
JOIN snowball_categories
  ON snowball_categories.official_category = snowball_inventory.category_name
  AND snowball_inventory.quantity > 0
  AND snowball_inventory.status = 'ready'
GROUP BY official_category
ORDER BY total_usable_snowballs;
+---------------------------+------------------------+
|     official_category     | total_usable_snowballs |
+---------------------------+------------------------+
| arctic blast premium      | 11470                  |
| everfrost training round  | 24248                  |
| polar precision microball | 70773                  |
| glacier sphere (xl)       | 165158                 |
| north ridge compact       | 594119                 |
| frost-flight deluxe       | 952019                 |
+---------------------------+------------------------+
Run Time: real 0.072 user 0.069623 sys 0.001929
sqlite> 
```

### RIGHT JOIN

```plaintext
sqlite> SELECT                                                       
  snowball_categories.official_category AS official_category,
  SUM(snowball_inventory.quantity) AS total_usable_snowballs
FROM snowball_inventory
RIGHT JOIN snowball_categories
  ON snowball_categories.official_category = snowball_inventory.category_name
  AND snowball_inventory.quantity > 0
  AND snowball_inventory.status = 'ready'
GROUP BY official_category
ORDER BY total_usable_snowballs;
+---------------------------+------------------------+
|     official_category     | total_usable_snowballs |
+---------------------------+------------------------+
| arctic blast premium      | 11470                  |
| everfrost training round  | 24248                  |
| polar precision microball | 70773                  |
| glacier sphere (xl)       | 165158                 |
| north ridge compact       | 594119                 |
| frost-flight deluxe       | 952019                 |
+---------------------------+------------------------+
Run Time: real 0.151 user 0.149690 sys 0.001957

sqlite> 
```

Time wise ranking:

1. INNER JOIN
    
2. SUBQUERY
    
3. RIGHT JOIN
    
4. LEFT JOIN
    

That's quite quick, not much can be measured here. I am sure there are other ways to do this. But stopping here for today. Explored JOINs a little.

See you for the day 3
