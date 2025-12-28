---
type: "sqlog"
title: "Advent of SQL 2025 Day 10: Misdelivered Presents"
slug: "advent-of-sql-2025-day-10"
date: 2025-12-25 15:30 +0530
series: ["advent-of-sql-2025"]
tags: ["sqlite", "sql", "advent-of-sql"]
---

## Advent of SQL, Day 10 - Misdelivered Presents

It's already day 10? We just need 5 more days now! Whoa! that flew by swiftly.

Let's pull in the data.

This is the SQL for day 10 in SQLite.

```sql
DROP TABLE IF EXISTS misdelivered_presents;
DROP TABLE IF EXISTS deliveries;

CREATE TABLE deliveries (
    id INT PRIMARY KEY,
    child_name TEXT,
    delivery_location TEXT,
    gift_name TEXT,
    scheduled_at TIMESTAMP
);

CREATE TABLE misdelivered_presents (
    id INT PRIMARY KEY,
    child_name TEXT,
    delivery_location TEXT,
    gift_name TEXT,
    scheduled_at TIMESTAMP,
    flagged_at TIMESTAMP,
    reason TEXT
);

INSERT INTO deliveries (id, child_name, delivery_location, gift_name, scheduled_at) VALUES
    (1, 'Omar Q.', '45 Maple Street', 'storybook collection', '2025-12-24 21:09:00'),
    (2, 'Sofia K.', '77 Snowflake Road', 'plush reindeer', '2025-12-24 18:35:00'),
    (3, 'Mila N.', 'The Vibes', 'storybook collection', '2025-12-24 21:09:00'),
    (4, 'Elias M.', 'Frost Hollow Cabin', 'board game', '2025-12-24 20:31:00'),
    (5, 'Ravi P.', '45 Maple Street', 'wooden train set', '2025-12-24 18:23:00'),
    (6, 'Jonah W.', '77 Snowflake Road', 'plush reindeer', '2025-12-24 20:34:00'),
    (7, 'Ava J.', '123 Evergreen Lane', 'board game', '2025-12-24 21:03:00'),
    (8, 'Omar Q.', '77 Snowflake Road', 'board game', '2025-12-24 18:56:00'),
    (9, 'Nia G.', 'Frost Hollow Cabin', 'teddy bear', '2025-12-24 21:27:00'),
    (10, 'Zara S.', 'North Pole Annex', 'wooden train set', '2025-12-24 20:58:00'),
    (11, 'Ravi P.', 'Frost Hollow Cabin', 'puzzle box', '2025-12-24 18:39:00'),
    (12, 'Jonah W.', '123 Evergreen Lane', 'puzzle box', '2025-12-24 18:23:00'),
    (13, 'Ravi P.', 'North Pole Annex', 'storybook collection', '2025-12-24 21:36:00'),
    (14, 'Lena F.', 'North Pole Annex', 'teddy bear', '2025-12-24 21:26:00'),
    (15, 'Ava J.', 'North Pole Annex', 'snow globe', '2025-12-24 18:31:00'),
    (16, 'Elias M.', 'Frost Hollow Cabin', 'robot toy', '2025-12-24 20:21:00'),
    (17, 'Sofia K.', 'Frost Hollow Cabin', 'teddy bear', '2025-12-24 20:27:00'),
    (18, 'Jonah W.', '77 Snowflake Road', 'storybook collection', '2025-12-24 20:49:00'),
    (19, 'Jonah W.', 'Frost Hollow Cabin', 'art supplies', '2025-12-24 21:38:00'),
    (20, 'Jonah W.', '123 Evergreen Lane', 'storybook collection', '2025-12-24 19:11:00');

INSERT INTO misdelivered_presents
(id, child_name, delivery_location, gift_name, scheduled_at, flagged_at, reason)
VALUES
    (601, 'Priya D.', 'The Vibes', 'plush reindeer', '2025-12-24 14:00:00', '2025-12-24 14:05:00', 'Invalid delivery location'),
    (602, 'Lena F.', 'Abandoned Lighthouse', 'board game', '2025-12-22 06:00:00', '2025-12-22 06:05:00', 'Invalid delivery location'),
    (603, 'Caleb O.', 'Drifting Igloo', 'board game', '2025-12-24 06:00:00', '2025-12-24 06:05:00', 'Invalid delivery location'),
    (604, 'Mateo C.', 'The Vibes', 'art supplies', '2025-12-22 04:00:00', '2025-12-22 04:05:00', 'Invalid delivery location'),
    (605, 'Hiro T.', 'The Vibes', 'robot toy', '2025-12-24 08:00:00', '2025-12-24 08:05:00', 'Invalid delivery location'),
    (606, 'Priya D.', 'Volcano Rim', 'puzzle box', '2025-12-22 08:00:00', '2025-12-22 08:05:00', 'Invalid delivery location'),
    (607, 'Nia G.', 'Abandoned Lighthouse', 'board game', '2025-12-24 01:00:00', '2025-12-24 01:05:00', 'Invalid delivery location'),
    (608, 'Elias M.', 'Drifting Igloo', 'board game', '2025-12-24 01:00:00', '2025-12-24 01:05:00', 'Invalid delivery location'),
    (609, 'Ravi P.', 'Volcano Rim', 'board game', '2025-12-24 02:00:00', '2025-12-24 02:05:00', 'Invalid delivery location'),
    (610, 'Hiro T.', 'Abandoned Lighthouse', 'science kit', '2025-12-23 20:00:00', '2025-12-23 20:05:00', 'Invalid delivery location'),
    (611, 'Priya D.', 'Drifting Igloo', 'puzzle box', '2025-12-22 21:00:00', '2025-12-22 21:05:00', 'Invalid delivery location'),
    (612, 'Hiro T.', 'Volcano Rim', 'art supplies', '2025-12-23 09:00:00', '2025-12-23 09:05:00', 'Invalid delivery location'),
    (613, 'Jonah W.', 'Abandoned Lighthouse', 'board game', '2025-12-24 01:00:00', '2025-12-24 01:05:00', 'Invalid delivery location'),
    (614, 'Omar Q.', 'Volcano Rim', 'art supplies', '2025-12-22 01:00:00', '2025-12-22 01:05:00', 'Invalid delivery location'),
    (615, 'Omar Q.', 'Drifting Igloo', 'science kit', '2025-12-23 20:00:00', '2025-12-23 20:05:00', 'Invalid delivery location'),
    (616, 'Omar Q.', 'Abandoned Lighthouse', 'teddy bear', '2025-12-24 12:00:00', '2025-12-24 12:05:00', 'Invalid delivery location'),
    (617, 'Zara S.', 'Volcano Rim', 'wooden train set', '2025-12-24 12:00:00', '2025-12-24 12:05:00', 'Invalid delivery location'),
    (618, 'Omar Q.', 'Abandoned Lighthouse', 'teddy bear', '2025-12-23 15:00:00', '2025-12-23 15:05:00', 'Invalid delivery location'),
    (619, 'Caleb O.', 'The Vibes', 'teddy bear', '2025-12-24 14:00:00', '2025-12-24 14:05:00', 'Invalid delivery location'),
    (620, 'Nia G.', 'Abandoned Lighthouse', 'board game', '2025-12-23 03:00:00', '2025-12-23 03:05:00', 'Invalid delivery location');
```

```sql
SELECT * FROM deliveries;
SELECT * FROM misdelivered_presents;
```

We have two tables. Almost the same with critical logical distinction among them and one extra column.

Let's see the problem to check it.


## Problem

> Clean-up the deliveries table to remove any records where the delivery_location is 'Volcano Rim', 'Drifting Igloo', 'Abandoned Lighthouse', 'The Vibes'.
> 
> Move those records to the misdelivered_presents with all the same columns as deliveries plus a flagged_at column with the current time and a reason column with "Invalid delivery location" listed as the reason for each moved record.
> 
> Make sure your final step shows the misdelivered_presents records that you just moved (i.e. don't include any existing records from the misdelivered_presents table).

Ok, this looks some easy problem.

- SELECT some data
- INSERT that data into other table
- DELETE that data from the original table
- SELECT the newly inserted data in the other table

Right?

Unless!

> Santa turned to you.
>
> “I don’t want this done in five steps,” he said. “And I don’t want any re-selecting. Move the problem presents out of the delivery system, log them in the vault, and show me exactly what you moved.”

Ouch Santa! Don't be lazy! Be smart he says! Huhh!


Ok, at least let's check both the tables, how many rows they have, and the rows that we want to move around.

```sql
SELECT COUNT(*) AS delivery_count FROM deliveries;
SELECT COUNT(*) AS misdelivered_present_count FROM misdelivered_presents;

SELECT 
    COUNT(*) AS misdelivered_deliveries_count
FROM deliveries 
WHERE 
    delivery_location IN (
        'Volcano Rim',
        'Drifting Igloo',
        'Abandoned Lighthouse',
        'The Vibes'
    );

SELECT 
    COUNT(*) AS misdelivered_present_count
FROM misdelivered_presents
WHERE 
    delivery_location IN (
        'Volcano Rim',
        'Drifting Igloo',
        'Abandoned Lighthouse',
        'The Vibes'
    );
```

```
sqlite> SELECT COUNT(*) FROM deliveries;
+----------+
| COUNT(*) |
+----------+
| 600      |
+----------+
sqlite> SELECT COUNT(*) FROM misdelivered_presents;
+----------+
| COUNT(*) |
+----------+
| 50       |
+----------+
sqlite> SELECT COUNT(*) FROM deliveries WHERE delivery_location IN ('Volcano Rim', 'Drifting Igloo', 'Abandoned Lighthouse', 'The Vibes');
+----------+
| COUNT(*) |
+----------+
| 103      |
+----------+
sqlite> SELECT COUNT(*) FROM misdelivered_presents WHERE delivery_location  IN ('Volcano Rim', 'Drifting Igloo', 'Abandoned Lighthouse', 'The Vibes');
+----------+
| COUNT(*) |
+----------+
| 50       |
+----------+
sqlite> 
```

OK, so we need to delete the `103` records from the `deliveries` table and move them into the `misdelivered_presents` at the same time.

How are we going to do that in SQLite.

We can try
- DELETE FROM deliveries with a SELECT from deliveries 
- Then INSERT the deleted data into misdelivered_presents

Could that work?

Let's see

```sql
WITH moved AS (
    DELETE 
        FROM deliveries
        WHERE delivery_location IN (
            'Volcano Rim',
            'Drifting Igloo',
            'Abandoned Lighthouse',
            'The Vibes'
        )
    RETURNING 
        id, 
        child_name, 
        delivery_location, 
        gift_name, 
        scheduled_at, 
        datetime('now') AS flagged_at, 
        'Invalid delivery location' AS reason
)
SELECT * FROM moved;
```

```

```

Ops! Can't generate a CTE with delete in it.

That's nasty. Thought that could shove it in the insert into the misdelivered table.

But it should be other way then?

Insert first and then use the data to delete?

```sql
WITH inserted_data AS (
  INSERT INTO misdelivered_presents (id, child_name, delivery_location, gift_name, scheduled_at)
  SELECT id, child_name, delivery_location, gift_name, scheduled_at
  FROM deliveries
  WHERE delivery_location IN ('Volcano Rim', 'Drifting Igloo', 'Abandoned Lighthouse', 'The Vibes')
)
DELETE FROM deliveries
WHERE delivery_location IN ('Volcano Rim', 'Drifting Igloo', 'Abandoned Lighthouse', 'The Vibes');

```

```
sqlite> WITH inserted_data AS (
  INSERT INTO misdelivered_presents (id, child_name, delivery_location, gift_name, scheduled_at)
  SELECT id, child_name, delivery_location, gift_name, scheduled_at
  FROM deliveries
  WHERE delivery_location IN ('Volcano Rim', 'Drifting Igloo', 'Abandoned Lighthouse', 'The Vibes')
)
DELETE FROM deliveries
WHERE delivery_location IN ('Volcano Rim', 'Drifting Igloo', 'Abandoned Lighthouse', 'The Vibes');
Parse error: near "INSERT": syntax error
  rt the selected rows into misdelivered_presents   INSERT INTO misdelivered_pre
                                      error here ---^
sqlite> 

```

Tried it, didn't work

Quoting the documentation here.

> Common Table Expressions or CTEs act like temporary views that exist only for the duration of a single SQL statement. There are two kinds of common table expressions: "ordinary" and "recursive". Ordinary common table expressions are helpful for making queries easier to understand by factoring subqueries out of the main SQL statement. Recursive common table expressions provide the ability to do hierarchical or recursive queries of trees and graphs, a capability that is not otherwise available in the SQL language.

I think it doesn't support CTEs with delete or insert! Sigh!

Now?

BEGIN COMMIT? Atomic Transactions?

Yeah!

Santa wanted it in one go right? That's not possible in SQLite but at least everything will happen or nothing will.

```sql
BEGIN;

WITH misdelivered_deliveries AS (
    SELECT * FROM deliveries 
    WHERE delivery_location IN (
        'Volcano Rim', 
        'Drifting Igloo', 
        'Abandoned Lighthouse', 
        'The Vibes')
)
INSERT INTO misdelivered_presents (
    id, 
    child_name,
    delivery_location, 
    gift_name, 
    scheduled_at, 
    flagged_at, 
    reason
)
SELECT 
    id,
    child_name,
    delivery_location,
    gift_name,
    scheduled_at,
    DATETIME('now'),
    'Invalid delivery location'
FROM misdelivered_deliveries
RETURNING *;

DELETE FROM deliveries 
WHERE id IN (
    SELECT id FROM deliveries 
    WHERE delivery_location IN ('Volcano Rim', 'Drifting Igloo', 'Abandoned Lighthouse', 'The Vibes')
);

COMMIT;
```

So here, we defined the CTE that selects data from `deliveries` and uses it to insert into the `misdelivered_deliveries` table. Then a separate query to delete that data from the `deliveries` table.

Yeah! I mean I don't think there could be other way to it!

We might or can use triggers to insert into one table when inserted in one table. But I think that is too much of a farfetched solution. We might create a trigger and instantly delete it as it could populate unwanted data if kept in the database.

### Trigger to insert when deleted

We can create a `TRIGGER` to insert into `misdelivered_presents` when something is deleted from `deliveries` table. We will separately have to delete the  records from the `deliveries` table. But the insert will happen automatically after the deletion.

Opening a fresh instance of the database!

```sql
CREATE TRIGGER move_misdelivered_presents
BEFORE DELETE ON deliveries
FOR EACH ROW
WHEN OLD.delivery_location IN ('Volcano Rim', 'Drifting Igloo', 'Abandoned Lighthouse', 'The Vibes')
BEGIN
    INSERT INTO misdelivered_presents (
        id, child_name, delivery_location, gift_name, 
        scheduled_at, flagged_at, reason
    )
    VALUES (
        OLD.id, OLD.child_name, OLD.delivery_location, OLD.gift_name, 
        OLD.scheduled_at, DATETIME('now'), 'Invalid delivery location'
    );
END;
```
This will create the trigger to insert the row into `misdelivered_presents` table when deleted from the `deliveries` table.

```sql
DELETE FROM deliveries 
WHERE delivery_location IN ('Volcano Rim', 'Drifting Igloo', 'Abandoned Lighthouse', 'The Vibes')
RETURNING *;
```

Ops! We are returning from the DELETE statement which is wrong as the problem stated to select from the `misdelivered_presents` table.

> Make sure your final step shows the misdelivered_presents records that you just moved (i.e. don't include any existing records from the misdelivered_presents table).

This is invalid then!


Though its not technically atomic. It happens before the delete, so it can mess up things.



```
sqlite> .read day10-inserts.sql
sqlite> .mode table 
sqlite> CREATE TRIGGER move_misdelivered_presents
BEFORE DELETE ON deliveries
FOR EACH ROW
WHEN OLD.delivery_location IN ('Volcano Rim', 'Drifting Igloo', 'Abandoned Lighthouse', 'The Vibes')
BEGIN
    INSERT INTO misdelivered_presents (
        id, child_name, delivery_location, gift_name, 
        scheduled_at, flagged_at, reason
    )
    VALUES (
        OLD.id, OLD.child_name, OLD.delivery_location, OLD.gift_name, 
        OLD.scheduled_at, DATETIME('now'), 'Invalid delivery location'
    );
END;
sqlite> DELETE FROM deliveries 
WHERE delivery_location IN ('Volcano Rim', 'Drifting Igloo', 'Abandoned Lighthouse', 'The Vibes')
RETURNING *;
+-----+------------+----------------------+----------------------+---------------------+
| id  | child_name |  delivery_location   |      gift_name       |    scheduled_at     |
+-----+------------+----------------------+----------------------+---------------------+
| 3   | Mila N.    | The Vibes            | storybook collection | 2025-12-24 21:09:00 |
| 22  | Lena F.    | Abandoned Lighthouse | plush reindeer       | 2025-12-24 19:08:00 |
| 23  | Mila N.    | Abandoned Lighthouse | storybook collection | 2025-12-24 20:42:00 |
| 29  | Mateo C.   | Volcano Rim          | plush reindeer       | 2025-12-24 21:44:00 |
| 31  | Nia G.     | Drifting Igloo       | robot toy            | 2025-12-24 19:57:00 |
...
...
| 582 | Zara S.    | The Vibes            | teddy bear           | 2025-12-24 21:20:00 |
| 585 | Layla B.   | Abandoned Lighthouse | wooden train set     | 2025-12-24 18:39:00 |
| 587 | Nia G.     | Volcano Rim          | storybook collection | 2025-12-24 18:35:00 |
| 596 | Omar Q.    | Abandoned Lighthouse | puzzle box           | 2025-12-24 19:28:00 |
+-----+------------+----------------------+----------------------+---------------------+
sqlite> DROP TRIGGER move_misdelivered_presents;
sqlite> SELECT COUNT(*) FROM deliveries;
+----------+
| COUNT(*) |
+----------+
| 497      |
+----------+
sqlite> SELECT COUNT(*) FROM misdelivered_presents;
+----------+
| COUNT(*) |
+----------+
| 153      |
+----------+
sqlite> 

```

So, it gets the result but I don't think the `TRIGGER` is a solution to this.

In SQLite, Atomic transaction using begin and end is the only thing to go for right?

Some one prove Santa that it can't be done in SQLite in one query? Please!

That's it from day 10, I have spent enough time on this, banging my head on the sqlite shell! 

Off to day 11 tomorrow!
