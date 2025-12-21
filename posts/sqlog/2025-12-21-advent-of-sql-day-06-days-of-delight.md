---
type: "sqlog"
title: "Advent of SQL 2025 Day 6: Days of Delight"
slug: "advent-of-sql-2025-day-6"
date: 2025-12-21
series: ["advent-of-sql-2025"]
tags: ["sqlite", "sql", "advent-of-sql"]
---

## Advent of SQL Day 6: Days of Delight

It is day 6 of advent of SQL.

Let's jump straight into the sql for the day.

```sql
DROP TABLE IF EXISTS families;
DROP TABLE IF EXISTS deliveries_assigned;

CREATE TABLE families (
    id INT PRIMARY KEY,
    family_name TEXT
);

CREATE TABLE deliveries_assigned (
    id INT PRIMARY KEY,
    family_id INT,
    gift_date DATE,
    gift_name TEXT
);

INSERT INTO families (id, family_name) VALUES
    (1, 'Isla Martinez'),
    (2, 'Nolan Garcia'),
    (3, 'Yara Chen'),
    (4, 'Tariq Nguyen'),
    (5, 'Mila Hernandez'),
    (6, 'Casey Kim'),
    (7, 'Mateo Hernandez'),
    (8, 'Keiko Petrov'),
    (9, 'Ethan Flores'),
    (10, 'Mateo Nakamura'),
    (11, 'Maya Fernandez'),
    (12, 'Mila Davis'),
    (13, 'Yara Rossi'),
    (14, 'Nolan Phillips'),
    (15, 'Amina Perez');

INSERT INTO deliveries_assigned (id, family_id, gift_date, gift_name) VALUES
    (1, 1, '2025-12-01', 'roasted cashews'),
    (2, 1, '2025-12-02', 'cookie decorating kit'),
    (3, 1, '2025-12-03', 'dark chocolate assortment'),
    (4, 1, '2025-12-04', 'white chocolate candies'),
    (5, 1, '2025-12-05', 'reindeer headband'),
    (6, 1, '2025-12-06', 'holiday brownie bites'),
    (7, 1, '2025-12-07', 'shortbread cookie tin'),
    (8, 1, '2025-12-08', 'chocolate chip cookies'),
    (9, 1, '2025-12-11', 'holiday jam trio'),
    (10, 1, '2025-12-12', 'white chocolate popcorn'),
    (11, 1, '2025-12-14', 'holiday jam trio'),
    (12, 1, '2025-12-15', 'fudge bites'),
    (13, 1, '2025-12-16', 'holiday sticker sheet'),
    (14, 1, '2025-12-18', 'hot cocoa bombs'),
    (15, 1, '2025-12-19', 'honey roasted nuts'),
    (16, 1, '2025-12-20', 'holiday mug'),
    (17, 1, '2025-12-21', 'white chocolate candies'),
    (18, 1, '2025-12-22', 'puzzle book'),
    (19, 1, '2025-12-23', 'snowman plush'),
    (20, 1, '2025-12-24', 'scented hand cream'),
    (21, 1, '2025-12-25', 'vanilla bean wafers'),
    (22, 2, '2025-12-01', 'roasted cashews'),
    (23, 2, '2025-12-02', 'holiday brownie bites'),
    (24, 2, '2025-12-03', 'peppermint bark bites'),
    (25, 2, '2025-12-04', 'holiday jam trio'),
    (26, 2, '2025-12-05', 'festive notepad'),
    (27, 2, '2025-12-06', 'scented pine sachet'),
    (28, 2, '2025-12-07', 'holiday mug'),
    (29, 2, '2025-12-08', 'shortbread cookie tin'),
    (30, 2, '2025-12-09', 'dark chocolate assortment');

```

So, we have two tables:
1. `families` table
2. `deliveries_assigned` table

The first table, `family` just has the id and the name of the family.

The second table, `deliveries_assigned` has the id, family id, gift date, and gift name.

Let's look at the problem statement.

## Problem

> Generate a report that returns the dates and families that have no delivery assigned after December 14th, using the `families` and `deliveries_assigned`.
> 
> Each row in the report should be a date and family name that represents the dates in which families don't have a delivery assigned yet.
> 
> Label the columns as `unassigned_date` and `name`. Order the results by `unassigned_date` and `name`, respectively, both in ascending order.


Ok, so we need to list the deliveries_assigned for each family first, to check what kind of pattern are we looking for.

```sql
SELECT * FROM deliveries_assigned WHERE family_id = 1;
```

```
sqlite> select * from deliveries_assigned where family_id=1;
+----+-----------+------------+---------------------------+
| id | family_id | gift_date  |         gift_name         |
+----+-----------+------------+---------------------------+
| 1  | 1         | 2025-12-01 | roasted cashews           |
| 2  | 1         | 2025-12-02 | cookie decorating kit     |
| 3  | 1         | 2025-12-03 | dark chocolate assortment |
| 4  | 1         | 2025-12-04 | white chocolate candies   |
| 5  | 1         | 2025-12-05 | reindeer headband         |
| 6  | 1         | 2025-12-06 | holiday brownie bites     |
| 7  | 1         | 2025-12-07 | shortbread cookie tin     |
| 8  | 1         | 2025-12-08 | chocolate chip cookies    |
| 9  | 1         | 2025-12-11 | holiday jam trio          |
| 10 | 1         | 2025-12-12 | white chocolate popcorn   |
| 11 | 1         | 2025-12-14 | holiday jam trio          |
| 12 | 1         | 2025-12-15 | fudge bites               |
| 13 | 1         | 2025-12-16 | holiday sticker sheet     |
| 14 | 1         | 2025-12-18 | hot cocoa bombs           |
| 15 | 1         | 2025-12-19 | honey roasted nuts        |
| 16 | 1         | 2025-12-20 | holiday mug               |
| 17 | 1         | 2025-12-21 | white chocolate candies   |
| 18 | 1         | 2025-12-22 | puzzle book               |
| 19 | 1         | 2025-12-23 | snowman plush             |
| 20 | 1         | 2025-12-24 | scented hand cream        |
| 21 | 1         | 2025-12-25 | vanilla bean wafers       |
+----+-----------+------------+---------------------------+
sqlite>
```

So, we are missing gifts for family id `1` on `09`, `10`, `13`, and `17`. But we are only asked for gifts after `14` (December 14th).

> Generate a report that returns the dates and families that have **no delivery assigned after December 14th**, using the families and deliveries_assigned.

So, we can discard `09`, `10`, and `13` as they are before December 14th.

```sql
SELECT * FROM deliveries_assigned WHERE family_id = 1 AND gift_date > '2025-12-14';
```

It gives the right dates where the gifts are assigned after December 14th. But the problem is we need to get dates which are missing in this `deliveries_assigned` table record for each family.

Finding something missing is kind of wired, because you don't have what is missing. Especially for dates, like dates are very painful.

We need to find among the sequential order of the dates, when some of the dates are missing, that is simple here, but you can see it could be quite cubersome if we have to manually add each date in the list for comparing with.

### JOINs with NOT EXISTS

So, the basic dirty solution is to check the missing dates for each family, one by one.

```sql
SELECT 
    families.family_name AS name,
    dates.column1 AS unassigned_date
FROM families
JOIN (
    VALUES 
        ('2025-12-15'), ('2025-12-16'), ('2025-12-17'),
        ('2025-12-18'), ('2025-12-19'), ('2025-12-20'),
        ('2025-12-21'), ('2025-12-22'), ('2025-12-23'),
        ('2025-12-24'), ('2025-12-25')
) AS dates ON 1=1
WHERE NOT EXISTS (
    SELECT 1 
    FROM deliveries_assigned
    WHERE deliveries_assigned.family_id = families.id 
    AND deliveries_assigned.gift_date = dates.column1
)
ORDER BY unassigned_date, name;
```

Let's break it down:
1. We have a `families` table with id and family name.
2. We have a `deliveries_assigned` table with id, family id, gift date, and gift name.
3. We create a list of dates from `2025-12-15` to `2025-12-25` using the `VALUES` keyword.
   - This just appends one date after other and names the columns as `column1` with the table as `dates`.
   - The `ON` condition is `1=1` to make sure the `WHERE` condition is true for the JOIN to happen.
4. We use the `NOT EXISTS` keyword to check if the `deliveries_assigned` table has a record for each date in the list.
   - we use `NOT EXISTS` because we want to check if there is no date for each date in the assigned list.
5. We order the results by `unassigned_date` and `name`, respectively, both in ascending order.

So, that is not the best way to solve this, I think.

### Recursive CTEs

We can generate a table full of dates, and then cross join the table with the `families` table. This will give us all possible combinations of dates and families like the gift should ideally be there for each family from 1st to 25th of December(however we are only interested from 15th to 25th December). Then once we have that full table of combination, we can check from the `deliveries_assigned` table and inner join for each family and filter out the rows which have family_id as `NULL` because some dates will be missing for that family.

First we'll create a recursive table of dates from `2025-12-15` to `2025-12-25`

```sql
WITH RECURSIVE dates(gift_date) AS (
    SELECT '2025-12-15'
    UNION ALL
    SELECT date(gift_date, '+1 day')
    FROM dates
    WHERE gift_date < '2025-12-25'
)
SELECT * FROM dates;
```

This will give us the dates from `2025-12-15` to `2025-12-25` in a recursive table.

What is a recursive table?

> A recursive table is a table that is defined as a combination of itself.

So, we can create a recursive table of dates from `2025-12-15` to `2025-12-25`, The base case is `2025-12-15`, and the recursive case is `SELECT date(gift_date, '+1 day') FROM dates WHERE gift_date < '2025-12-25'` which means it will call it recursively for `2025-12-16` on the first call inside it because of `+1 day` as the interval in the `date` function.

The date function is a [function](https://sqlite.org/lang_datefunc.html#modifiers) which takes in a date and we can add modifiers to it to manipulate or extract parts of the date. Here we have added the modifier as `+1 day` which will increment the day by one. We then call that in the `dates` CTE again till we have the `date` as less then `2025-12-25`. Till then we will have created all the `dates` from `15` till `25` including `25`. We can just type in the values manually as we did in the first dirty solution, but I wanted to see how we can generate dates dynamically in sqlite.

```
sqlite> WITH RECURSIVE dates(gift_date) AS (
    SELECT '2025-12-15'
    UNION ALL
    SELECT date(gift_date, '+1 day')
    FROM dates
    WHERE gift_date < '2025-12-25'
)
SELECT * FROM dates;
+------------+
| gift_date  |
+------------+
| 2025-12-15 |
| 2025-12-16 |
| 2025-12-17 |
| 2025-12-18 |
| 2025-12-19 |
| 2025-12-20 |
| 2025-12-21 |
| 2025-12-22 |
| 2025-12-23 |
| 2025-12-24 |
| 2025-12-25 |
+------------+
Run Time: real 0.000 user 0.000225 sys 0.000009
sqlite>
```

Now we'll cross join the `families` table with the `dates` table.

```sql
SELECT families.id AS family_id, families.family_name as family_name, dates.gift_date as gift_date
FROM families
CROSS JOIN dates;
```

OK! Wait include the `dates` CTE above too, was just simplifying the query.

```sql
WITH RECURSIVE dates(gift_date) AS (
    SELECT '2025-12-15'
    UNION ALL
    SELECT date(gift_date, '+1 day')
    FROM dates
    WHERE gift_date < '2025-12-25'
)
SELECT families.id AS family_id, families.family_name as family_name, dates.gift_date as gift_date
FROM families
CROSS JOIN dates;
```
This will give us the full table of combination of dates and families.

```sql
SELECT COUNT(*) FROM families;
```

```sql
sqlite> SELECT COUNT(*) FROM families;
+----------+
| COUNT(*) |
+----------+
| 250      |
+----------+
Run Time: real 0.000 user 0.000113 sys 0.000006
sqlite>
```

That is :

- There are 11 dates right? `2025-12-15` to `2025-12-25`, -> `15` (1), `16` (2), `17` (3), `18` (4), `19` (5), `20` (6), `21` (7), `22` (8), `23` (9), `24` (10), `25` (11)
- There are I think `250` families.
- So, a `CROSS JOIN` will give us `11 * 250 = 2750` rows.


```sql
SELECT COUNT(*) FROM (WITH RECURSIVE dates(gift_date) AS (
    SELECT '2025-12-15'
    UNION ALL
    SELECT date(gift_date, '+1 day')
    FROM dates
    WHERE gift_date < '2025-12-25'
)
SELECT families.id AS family_id, families.family_name as family_name, dates.gift_date as gift_date
FROM families
CROSS JOIN dates) as count;
```

11 dates for each family.

```
sqlite> SELECT COUNT(*) FROM (WITH RECURSIVE dates(gift_date) AS (
    SELECT '2025-12-15'
    UNION ALL
    SELECT date(gift_date, '+1 day')
    FROM dates
    WHERE gift_date < '2025-12-25'
)
SELECT families.id AS family_id, families.family_name as family_name, dates.gift_date as gift_date
FROM families
CROSS JOIN dates) as count;
+----------+
| COUNT(*) |
+----------+
| 2750     |
+----------+
Run Time: real 0.001 user 0.000407 sys 0.000000
sqlite>
```

Now, we simply have to join the `deliveries_assigned` table with the above table.

Why?
Because we need to map which dates are assigned and which dates are missing.

We need to do a which join?

LEFT, RIGHT or INNER

LEFT

WHy?

Because, the left table will have all the dates, and the right table will have the assigned dates. **We need all the records from the `left` table (the combination, cross join table)**

We need all the rows in the `left` or the combination(cross join) of date and family table as to map which dates are assigned and which dates are missing.

If the `left` table has a record, then it means the date is assigned. If the `right` table has a record i.e. the gift is assigned, then it means the date is not assigned.

Hence, we can simply then filter out the relations with `NULL` as the `family_id` in the `deliveries_assigned` table.

```sql
WITH RECURSIVE dates(gift_date) AS (
    SELECT '2025-12-15'
    UNION ALL
    SELECT date(gift_date, '+1 day')
    FROM dates
    WHERE gift_date < '2025-12-25'
),
combination AS (
    SELECT families.id AS family_id, families.family_name as family_name, dates.gift_date as gift_date
    FROM families
    CROSS JOIN dates
)
SELECT
    *
FROM combination
LEFT JOIN deliveries_assigned
    ON deliveries_assigned.family_id = combination.family_id
    AND deliveries_assigned.gift_date = combination.gift_date
```

```
sqlite> WITH RECURSIVE dates(gift_date) AS (
    SELECT '2025-12-15'
    UNION ALL
    SELECT date(gift_date, '+1 day')
    FROM dates
    WHERE gift_date < '2025-12-25'
),
combination AS (
    SELECT families.id AS family_id, families.family_name as family_name, dates.gift_date as gift_date
    FROM families
    CROSS JOIN dates
)
SELECT
    *
FROM combination
LEFT JOIN deliveries_assigned
    ON deliveries_assigned.family_id = combination.family_id
    AND deliveries_assigned.gift_date = combination.gift_date LIMIT 30;
+-----------+---------------+------------+----+-----------+------------+-------------------------+
| family_id |  family_name  | gift_date  | id | family_id | gift_date  |        gift_name        |
+-----------+---------------+------------+----+-----------+------------+-------------------------+
| 1         | Isla Martinez | 2025-12-15 | 12 | 1         | 2025-12-15 | fudge bites             |
| 1         | Isla Martinez | 2025-12-16 | 13 | 1         | 2025-12-16 | holiday sticker sheet   |
| 1         | Isla Martinez | 2025-12-17 |    |           |            |                         |
| 1         | Isla Martinez | 2025-12-18 | 14 | 1         | 2025-12-18 | hot cocoa bombs         |
| 1         | Isla Martinez | 2025-12-19 | 15 | 1         | 2025-12-19 | honey roasted nuts      |
| 1         | Isla Martinez | 2025-12-20 | 16 | 1         | 2025-12-20 | holiday mug             |
| 1         | Isla Martinez | 2025-12-21 | 17 | 1         | 2025-12-21 | white chocolate candies |
| 1         | Isla Martinez | 2025-12-22 | 18 | 1         | 2025-12-22 | puzzle book             |
| 1         | Isla Martinez | 2025-12-23 | 19 | 1         | 2025-12-23 | snowman plush           |
| 1         | Isla Martinez | 2025-12-24 | 20 | 1         | 2025-12-24 | scented hand cream      |
| 1         | Isla Martinez | 2025-12-25 | 21 | 1         | 2025-12-25 | vanilla bean wafers     |
| 2         | Nolan Garcia  | 2025-12-15 | 36 | 2         | 2025-12-15 | mini marshmallow tubes  |
| 2         | Nolan Garcia  | 2025-12-16 | 37 | 2         | 2025-12-16 | white chocolate candies |
| 2         | Nolan Garcia  | 2025-12-17 | 38 | 2         | 2025-12-17 | gingerbread cookie kit  |
| 2         | Nolan Garcia  | 2025-12-18 | 39 | 2         | 2025-12-18 | family card game        |
| 2         | Nolan Garcia  | 2025-12-19 |    |           |            |                         |
| 2         | Nolan Garcia  | 2025-12-20 | 40 | 2         | 2025-12-20 | santa hat               |
| 2         | Nolan Garcia  | 2025-12-21 | 41 | 2         | 2025-12-21 | holiday sticker sheet   |
| 2         | Nolan Garcia  | 2025-12-22 |    |           |            |                         |
| 2         | Nolan Garcia  | 2025-12-23 | 42 | 2         | 2025-12-23 | pecan praline bites     |
| 2         | Nolan Garcia  | 2025-12-24 |    |           |            |                         |
| 2         | Nolan Garcia  | 2025-12-25 | 43 | 2         | 2025-12-25 | santa hat               |
| 3         | Yara Chen     | 2025-12-15 | 57 | 3         | 2025-12-15 | peppermint bark bites   |
| 3         | Yara Chen     | 2025-12-16 |    |           |            |                         |
| 3         | Yara Chen     | 2025-12-17 |    |           |            |                         |
| 3         | Yara Chen     | 2025-12-18 | 58 | 3         | 2025-12-18 | cheddar popcorn         |
| 3         | Yara Chen     | 2025-12-19 |    |           |            |                         |
| 3         | Yara Chen     | 2025-12-20 | 59 | 3         | 2025-12-20 | festive notepad         |
| 3         | Yara Chen     | 2025-12-21 | 60 | 3         | 2025-12-21 | fruit assortment        |
| 3         | Yara Chen     | 2025-12-22 |    |           |            |                         |
+-----------+---------------+------------+----+-----------+------------+-------------------------+
...
...
| 249       | Jude Bautista     | 2025-12-24 |      |           |            |                            |
| 249       | Jude Bautista     | 2025-12-25 | 5073 | 249       | 2025-12-25 | almond brittle             |
| 250       | Bianca Muller     | 2025-12-15 | 5086 | 250       | 2025-12-15 | cocoa mix bundle           |
| 250       | Bianca Muller     | 2025-12-16 | 5087 | 250       | 2025-12-16 | cookie decorating kit      |
| 250       | Bianca Muller     | 2025-12-17 | 5088 | 250       | 2025-12-17 | shortbread cookie tin      |
| 250       | Bianca Muller     | 2025-12-18 |      |           |            |                            |
| 250       | Bianca Muller     | 2025-12-19 | 5089 | 250       | 2025-12-19 | snowflake candle           |
| 250       | Bianca Muller     | 2025-12-20 |      |           |            |                            |
| 250       | Bianca Muller     | 2025-12-21 | 5090 | 250       | 2025-12-21 | trail mix trio             |
| 250       | Bianca Muller     | 2025-12-22 | 5091 | 250       | 2025-12-22 | shortbread cookie tin      |
| 250       | Bianca Muller     | 2025-12-23 |      |           |            |                            |
| 250       | Bianca Muller     | 2025-12-24 |      |           |            |                            |
| 250       | Bianca Muller     | 2025-12-25 | 5092 | 250       | 2025-12-25 | gingerbread cookie kit     |
+-----------+-------------------+------------+------+-----------+------------+----------------------------+
Run Time: real 0.028 user 0.013392 sys 0.013096
sqlite>
```

This will give all the gifts that have been assigned as well as not assigned for each family on each date from `2025-12-15` to `2025-12-25`.

Now, you can see, what we need and what we don't we simply can get the gifts or records with the `family_id` in the `deliveries_assigned` table as `NULL`, since there was no record for the family_id in the `deliveries_assigned` table for that date.


```sql
WITH RECURSIVE dates(gift_date) AS (
    SELECT '2025-12-15'
    UNION ALL
    SELECT date(gift_date, '+1 day')
    FROM dates
    WHERE gift_date < '2025-12-25'
),
combination AS (
    SELECT families.id AS family_id, families.family_name as family_name, dates.gift_date as gift_date
    FROM families
    CROSS JOIN dates
)
SELECT
    *
FROM combination
LEFT JOIN deliveries_assigned
    ON deliveries_assigned.family_id = combination.family_id
    AND deliveries_assigned.gift_date = combination.gift_date
WHERE deliveries_assigned.family_id IS NULL
```
So, a simple `WHERE` clause with deliveries_assigned.family_id `IS NULL` will give us the missing dates for each family.

```
sqlite> WITH RECURSIVE dates(gift_date) AS (
    SELECT '2025-12-15'
    UNION ALL
    SELECT date(gift_date, '+1 day')
    FROM dates
    WHERE gift_date < '2025-12-25'
),
combination AS (
    SELECT families.id AS family_id, families.family_name as family_name, dates.gift_date as gift_date
    FROM families
    CROSS JOIN dates
)
SELECT
    *
FROM combination
LEFT JOIN deliveries_assigned
    ON deliveries_assigned.family_id = combination.family_id
    AND deliveries_assigned.gift_date = combination.gift_date
WHERE deliveries_assigned.family_id IS NULL
   ...> ;
+-----------+-------------------+------------+----+-----------+-----------+-----------+
| family_id |    family_name    | gift_date  | id | family_id | gift_date | gift_name |
+-----------+-------------------+------------+----+-----------+-----------+-----------+
| 1         | Isla Martinez     | 2025-12-17 |    |           |           |           |
| 2         | Nolan Garcia      | 2025-12-19 |    |           |           |           |
| 2         | Nolan Garcia      | 2025-12-22 |    |           |           |           |
| 2         | Nolan Garcia      | 2025-12-24 |    |           |           |           |
| 3         | Yara Chen         | 2025-12-16 |    |           |           |           |
| 3         | Yara Chen         | 2025-12-17 |    |           |           |           |
| 3         | Yara Chen         | 2025-12-19 |    |           |           |           |
| 3         | Yara Chen         | 2025-12-22 |    |           |           |           |
| 4         | Tariq Nguyen      | 2025-12-16 |    |           |           |           |
...
...
| 247       | Malik Kim         | 2025-12-18 |    |           |           |           |
| 247       | Malik Kim         | 2025-12-20 |    |           |           |           |
| 247       | Malik Kim         | 2025-12-21 |    |           |           |           |
| 248       | Tariq Flores      | 2025-12-16 |    |           |           |           |
| 248       | Tariq Flores      | 2025-12-17 |    |           |           |           |
| 248       | Tariq Flores      | 2025-12-18 |    |           |           |           |
| 249       | Jude Bautista     | 2025-12-16 |    |           |           |           |
| 249       | Jude Bautista     | 2025-12-20 |    |           |           |           |
| 249       | Jude Bautista     | 2025-12-21 |    |           |           |           |
| 249       | Jude Bautista     | 2025-12-23 |    |           |           |           |
| 249       | Jude Bautista     | 2025-12-24 |    |           |           |           |
| 250       | Bianca Muller     | 2025-12-18 |    |           |           |           |
| 250       | Bianca Muller     | 2025-12-20 |    |           |           |           |
| 250       | Bianca Muller     | 2025-12-23 |    |           |           |           |
| 250       | Bianca Muller     | 2025-12-24 |    |           |           |           |
+-----------+-------------------+------------+----+-----------+-----------+-----------+
Run Time: real 0.014 user 0.011508 sys 0.002046
sqlite>
```

All right, we now need to order by `unassigned_date` and `name` which are the `dates` from the `combination` table and the `family_name` from the `families` table.

```sql
WITH RECURSIVE dates(gift_date) AS (
    SELECT '2025-12-15'
    UNION ALL
    SELECT date(gift_date, '+1 day')
    FROM dates
    WHERE gift_date < '2025-12-25'
),
combination AS (
    SELECT families.id AS family_id, families.family_name as family_name, dates.gift_date as gift_date
    FROM families
    CROSS JOIN dates
)
SELECT
    combination.gift_date as unassigned_date,
    family_name
FROM combination
LEFT JOIN deliveries_assigned
    ON deliveries_assigned.family_id = combination.family_id
    AND deliveries_assigned.gift_date = combination.gift_date
WHERE deliveries_assigned.family_id IS NULL
ORDER BY unassigned_date, family_name;
```

```
+-----------------+-------------------+
| unassigned_date |    family_name    |
+-----------------+-------------------+
| 2025-12-15      | Adil Rossi        |
| 2025-12-15      | Aisha Connor      |
| 2025-12-15      | Amina Perez       |
| 2025-12-15      | Amina Wong        |
| 2025-12-15      | Andre Flores      |
| 2025-12-15      | Anya Singh        |
| 2025-12-15      | Arjun Wong        |
| 2025-12-15      | Bianca Connor     |
| 2025-12-15      | Caleb Petrov      |
| 2025-12-15      | Caleb Roberts     |
| 2025-12-15      | Carmen Carter     |
| 2025-12-15      | Carmen Garcia     |
| 2025-12-15      | Casey Flores      |
| 2025-12-15      | Chi Hughes        |
| 2025-12-15      | Clara Johnson     |
| 2025-12-15      | Dara Bautista     |
| 2025-12-15      | David Ramirez     |
| 2025-12-15      | Elias Petrov      |
| 2025-12-15      | Elias Petrov      |
| 2025-12-15      | Ethan Flores      |
| 2025-12-15      | Eva Gonzalez      |
...
...
| 2025-12-25      | Owen Park         |
| 2025-12-25      | Priya Khan        |
| 2025-12-25      | Rafael Singh      |
| 2025-12-25      | Ravi Abdallah     |
| 2025-12-25      | Ravi Mitchell     |
| 2025-12-25      | Rosa Turner       |
| 2025-12-25      | Sara Jensen       |
| 2025-12-25      | Sara Lopez        |
| 2025-12-25      | Sara Rossi        |
| 2025-12-25      | Sarah Phillips    |
| 2025-12-25      | Seth Garcia       |
| 2025-12-25      | Sienna Lopez      |
| 2025-12-25      | Sofia Nakamura    |
| 2025-12-25      | Tariq Nguyen      |
| 2025-12-25      | Uma Ali           |
| 2025-12-25      | Uma Phillips      |
| 2025-12-25      | Yara Chen         |
| 2025-12-25      | Yara Rossi        |
| 2025-12-25      | Yusuf Ali         |
| 2025-12-25      | Yusuf Hansen      |
| 2025-12-25      | Yusuf Perez       |
| 2025-12-25      | Yusuf Rossi       |
| 2025-12-25      | Zara Khan         |
+-----------------+-------------------+
Run Time: real 0.008 user 0.004613 sys 0.002845
sqlite>
```

Phew!
Ok, that looks a mamoth query.

```sql
WITH RECURSIVE dates(gift_date) AS (
    SELECT '2025-12-15'
    UNION ALL
    SELECT date(gift_date, '+1 day')
    FROM dates
    WHERE gift_date < '2025-12-25'
),
combination AS (
    SELECT families.id AS family_id, families.family_name as family_name, dates.gift_date as gift_date
    FROM families
    CROSS JOIN dates
)
SELECT
    combination.gift_date as unassigned_date,
    family_name
FROM combination
LEFT JOIN deliveries_assigned
    ON deliveries_assigned.family_id = combination.family_id
    AND deliveries_assigned.gift_date = combination.gift_date
WHERE deliveries_assigned.family_id IS NULL
ORDER BY unassigned_date, family_name;
```

So, to recap

- Generate a table of dates from `2025-12-15` to `2025-12-25` using a recursive CTE.
- Create a table of combinations of `families` and `dates` using a cross join.
- Left join the `deliveries_assigned` table with the combination table.
- Filter out the rows where the `family_id` in the `deliveries_assigned` table is `NULL`.
- Order the results by `unassigned_date` and `name`, respectively, both in ascending order.

Simple right?

That's it from day 6. 

It getting serious out there!



