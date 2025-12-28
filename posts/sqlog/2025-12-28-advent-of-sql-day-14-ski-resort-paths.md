---
type: "sqlog"
title: "Advent of SQL 2025 Day 14: Ski Resort Paths"
slug: "advent-of-sql-2025-day-14"
date: 2025-12-28 15:30 +0530
series: ["advent-of-sql-2025"]
tags: ["sqlite", "sql", "advent-of-sql"]
---

## Advent of SQL, Day 14 - Ski Resort Paths

Ok, almost to the penultimate day in the series. It is day 14 of Advent of SQL.

Let's grab the SQL for the day.

```sql
DROP TABLE IF EXISTS mountain_network;


CREATE TABLE mountain_network (
    id INTEGER PRIMARY KEY,
    from_node TEXT,
    to_node TEXT,
    node_type TEXT,    -- 'Lift' or 'Trail'
    difficulty TEXT    -- Only applicable for trails: 'green', 'blue', 'black', 'double_black'
);

INSERT INTO mountain_network (id, from_node, to_node, node_type, difficulty) VALUES
(1, 'Outlaw Express', 'Stairway Lift', 'Lift', NULL),
(2, 'Outlaw Express', 'Top Gun Bowl', 'Trail', 'black'),
(3, 'Top Gun Bowl', 'Top Gun', 'Trail', 'black'),
(4, 'Top Gun', 'Montoya', 'Trail', 'blue'),
(5, 'Montoya', 'Center Aisle', 'Trail', 'green'),
(6, 'Center Aisle', 'Lower Stampede', 'Trail', 'green'),
(7, 'Stairway Lift', 'Red''s Lift', 'Lift', NULL),
(8, 'Stairway Lift', 'Broadway', 'Trail', 'green'),
(9, 'Red''s Lift', 'Bearclaw', 'Trail', 'blue'),
(10, 'Bearclaw', 'Last Chance', 'Trail', 'blue'),
(11, 'Last Chance', 'Diamondback', 'Trail', 'blue'),
(12, 'Diamondback', 'Broadway', 'Trail', 'green'),
(13, 'Red''s Lift', 'Bishop''s Bowl', 'Trail', 'black'),
(14, 'Red''s Lift', 'Amy''s Ridge', 'Trail', 'blue'),
(15, 'Amy''s Ridge', 'Grizzly Bowl', 'Trail', 'black'),
(16, 'Flathead Lift', 'Amy''s Ridge', 'Trail', 'blue'),
(17, 'Jake''s Lift', 'Wildwood Lift', 'Lift', NULL),
(18, 'Wildwood Lift', 'Sidewinder', 'Trail', 'green'),
(19, 'Wildwood Lift', 'Brightside', 'Trail', 'blue'),
(20, 'Brightside', 'Moonrise', 'Trail', 'green'),
(21, 'Moonrise', 'Draw', 'Trail', 'green'),
(22, 'Moonrise', 'Lone Pine', 'Trail', 'blue'),
(23, 'Draw', 'Maverick', 'Trail', 'blue'),
(24, 'Draw', 'Broadway', 'Trail', 'green'),
(25, 'Broadway', 'Outlaw Trail', 'Trail', 'green'),
(26, 'Outlaw Trail', 'Center Aisle', 'Trail', 'green'),
(27, 'Center Aisle', 'Bandit', 'Trail', 'green'),
(28, 'Jake''s Lift', 'Maverick', 'Trail', 'blue');

```

That's it, we just have 1 table called `mountain_network` and it has `28` records.

Let's look at the problem statement to check what we need to do to make sense of those 28 rows. 


## Problem

> Find all the possible routes from Jake's Lift to Maverick. None of the possible routes will take more than 12 connections.


Oh! This is a graph like or network like problem.

Ouch! 

Relational databases looks like they would be a good fit for these stuff, but if the data is just in the single table, there's not really much help.

Let's see how we can think about it.

### JOINs and UNIONS

We need to find all the routes from a start node(record) i.e. `Jake's Lift` and find all ways that lead to the node(record) i.e. `Maverick`. The table gives a list of edges i.e. from which node to which node there is a way or a trail or lift.

So, we can do a simple select to check if the `from_node` has `Jake's Lift` and the `to_node` has `Maverick` right?

But that would make it too long.

Like we would have to look for 12 consecutive edges and branch off at each direction.

For the first level it would look simple like this:

```sql
SELECT 
    mountain_network.from_node || ' -> ' || mountain_network.to_node AS path, 
    1 AS connections
FROM mountain_network
WHERE 
    mountain_network.from_node = 'Jake''s Lift'
    AND mountain_network.to_node = 'Maverick';
```

```
sqlite> SELECT mountain_network.from_node || ' -> ' || mountain_network.to_node AS path, 1 AS connections
FROM mountain_network
WHERE mountain_network.from_node = 'Jake''s Lift' AND mountain_network.to_node = 'Maverick'
   ...> ;
+-------------------------+-------------+
|          path           | connections |
+-------------------------+-------------+
| Jake's Lift -> Maverick | 1           |
+-------------------------+-------------+
sqlite>
```

Luckily, we have one route directly from `Jake's Lift` to `Maverick`. But that might not be for all levels. Like we need to check then, from each of the start nodes that begin at `Jake's Lift`, then for all the nodes that begin at that then it branches of further till we have a dead end only at `Maverick`.

Phew that's going to be a long one.

```sql
SELECT T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node, 2                     FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
WHERE 
    T1.from_node = 'Jake''s Lift'
    AND T2.to_node = 'Maverick';
```

That is the first branch off from `Jake's Lift`. If we remove the `T2.to_node = 'Maverick'` you would see all nodes from the `Jake's Lift`, not necessarily ending at `Maverick`

```sql
SELECT T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node as path,
 2 as connections
FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
WHERE T1.from_node = 'Jake''s Lift';
``` 

```
sqlite> SELECT T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node as path,
 2 as connections
FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
WHERE T1.from_node = 'Jake''s Lift';
+--------------------------------------------+-------------+
|                    path                    | connections |
+--------------------------------------------+-------------+
| Jake's Lift -> Wildwood Lift -> Brightside | 2           |
| Jake's Lift -> Wildwood Lift -> Sidewinder | 2           |
+--------------------------------------------+-------------+
sqlite>
```

Similarly we can do it for 3 and 4, and till 12 connections.

```sql
SELECT 
    T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node || ' -> ' || T3.to_node as path,
    3 as connections
FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
JOIN mountain_network T3 ON T2.to_node = T3.from_node
WHERE T1.from_node = 'Jake''s Lift';
```

```
sqlite> SELECT T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node || ' -> ' || T3.to_node as path, 3 as connections
FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
JOIN mountain_network T3 ON T2.to_node = T3.from_node
WHERE T1.from_node = 'Jake''s Lift'
   ...> ;
+--------------------------------------------------------+-------------+
|                          path                          | connections |
+--------------------------------------------------------+-------------+
| Jake's Lift -> Wildwood Lift -> Brightside -> Moonrise | 3           |
+--------------------------------------------------------+-------------+
sqlite>
```


Yikes.

That gives a long long query. I cannot write.

```sql
SELECT T1.from_node || ' -> ' || T1.to_node AS path, 1 AS connections
FROM mountain_network T1
WHERE T1.from_node = 'Jake''s Lift' AND T1.to_node = 'Maverick'

UNION ALL

SELECT T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node, 2
FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
WHERE T1.from_node = 'Jake''s Lift' AND T2.to_node = 'Maverick'

UNION ALL

SELECT T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node || ' -> ' || T3.to_node, 3
FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
JOIN mountain_network T3 ON T2.to_node = T3.from_node
WHERE T1.from_node = 'Jake''s Lift' AND T3.to_node = 'Maverick'

UNION ALL

SELECT T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node || ' -> ' || T3.to_node || ' -> ' || T4.to_node, 4
FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
JOIN mountain_network T3 ON T2.to_node = T3.from_node
JOIN mountain_network T4 ON T3.to_node = T4.from_node
WHERE T1.from_node = 'Jake''s Lift' AND T4.to_node = 'Maverick'

UNION ALL

SELECT T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node || ' -> ' || T3.to_node || ' -> ' || T4.to_node || ' -> ' || T5.to_node, 5
FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
JOIN mountain_network T3 ON T2.to_node = T3.from_node
JOIN mountain_network T4 ON T3.to_node = T4.from_node
JOIN mountain_network T5 ON T4.to_node = T5.from_node
WHERE T1.from_node = 'Jake''s Lift' AND T5.to_node = 'Maverick'

UNION ALL

SELECT T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node || ' -> ' || T3.to_node || ' -> ' || T4.to_node || ' -> ' || T5.to_node || ' -> ' || T6.to_node, 6
FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
JOIN mountain_network T3 ON T2.to_node = T3.from_node
JOIN mountain_network T4 ON T3.to_node = T4.from_node
JOIN mountain_network T5 ON T4.to_node = T5.from_node
JOIN mountain_network T6 ON T5.to_node = T6.from_node
WHERE T1.from_node = 'Jake''s Lift' AND T6.to_node = 'Maverick'


UNION ALL

SELECT T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node || ' -> ' || T3.to_node || ' -> ' || T4.to_node || ' -> ' || T5.to_node || ' -> ' || T6.to_node || ' -> ' || T7.to_node, 7
FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
JOIN mountain_network T3 ON T2.to_node = T3.from_node
JOIN mountain_network T4 ON T3.to_node = T4.from_node
JOIN mountain_network T5 ON T4.to_node = T5.from_node
JOIN mountain_network T6 ON T5.to_node = T6.from_node
JOIN mountain_network T7 ON T6.to_node = T7.from_node
WHERE T1.from_node = 'Jake''s Lift' AND T7.to_node = 'Maverick'

UNION ALL

SELECT T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node || ' -> ' || T3.to_node || ' -> ' || T4.to_node || ' -> ' || T5.to_node || ' -> ' || T6.to_node || ' -> ' || T7.to_node || ' -> ' || T8.to_node, 8
FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
JOIN mountain_network T3 ON T2.to_node = T3.from_node
JOIN mountain_network T4 ON T3.to_node = T4.from_node
JOIN mountain_network T5 ON T4.to_node = T5.from_node
JOIN mountain_network T6 ON T5.to_node = T6.from_node
JOIN mountain_network T7 ON T6.to_node = T7.from_node
JOIN mountain_network T8 ON T7.to_node = T8.from_node
WHERE T1.from_node = 'Jake''s Lift' AND T8.to_node = 'Maverick'

UNION ALL

SELECT T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node || ' -> ' || T3.to_node || ' -> ' || T4.to_node || ' -> ' || T5.to_node || ' -> ' || T6.to_node || ' -> ' || T7.to_node || ' -> ' || T8.to_node || ' -> ' || T9.to_node, 9
FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
JOIN mountain_network T3 ON T2.to_node = T3.from_node
JOIN mountain_network T4 ON T3.to_node = T4.from_node
JOIN mountain_network T5 ON T4.to_node = T5.from_node
JOIN mountain_network T6 ON T5.to_node = T6.from_node
JOIN mountain_network T7 ON T6.to_node = T7.from_node
JOIN mountain_network T8 ON T7.to_node = T8.from_node
JOIN mountain_network T9 ON T8.to_node = T9.from_node
WHERE T1.from_node = 'Jake''s Lift' AND T9.to_node = 'Maverick'

UNION ALL

SELECT T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node || ' -> ' || T3.to_node || ' -> ' || T4.to_node || ' -> ' || T5.to_node || ' -> ' || T6.to_node || ' -> ' || T7.to_node || ' -> ' || T8.to_node || ' -> ' || T9.to_node || ' -> ' || T10.to_node, 10
FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
JOIN mountain_network T3 ON T2.to_node = T3.from_node
JOIN mountain_network T4 ON T3.to_node = T4.from_node
JOIN mountain_network T5 ON T4.to_node = T5.from_node
JOIN mountain_network T6 ON T5.to_node = T6.from_node
JOIN mountain_network T7 ON T6.to_node = T7.from_node
JOIN mountain_network T8 ON T7.to_node = T8.from_node
JOIN mountain_network T9 ON T8.to_node = T9.from_node
JOIN mountain_network T10 ON T9.to_node = T10.from_node
WHERE T1.from_node = 'Jake''s Lift' AND T10.to_node = 'Maverick'

UNION ALL

SELECT T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node || ' -> ' || T3.to_node || ' -> ' || T4.to_node || ' -> ' || T5.to_node || ' -> ' || T6.to_node || ' -> ' || T7.to_node || ' -> ' || T8.to_node || ' -> ' || T9.to_node || ' -> ' || T10.to_node || ' -> ' || T11.to_node, 11
FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
JOIN mountain_network T3 ON T2.to_node = T3.from_node
JOIN mountain_network T4 ON T3.to_node = T4.from_node
JOIN mountain_network T5 ON T4.to_node = T5.from_node
JOIN mountain_network T6 ON T5.to_node = T6.from_node
JOIN mountain_network T7 ON T6.to_node = T7.from_node
JOIN mountain_network T8 ON T7.to_node = T8.from_node
JOIN mountain_network T9 ON T8.to_node = T9.from_node
JOIN mountain_network T10 ON T9.to_node = T10.from_node
JOIN mountain_network T11 ON T10.to_node = T11.from_node
WHERE T1.from_node = 'Jake''s Lift' AND T11.to_node = 'Maverick'

UNION ALL

SELECT T1.from_node || ' -> ' || T1.to_node || ' -> ' || T2.to_node || ' -> ' || T3.to_node || ' -> ' || T4.to_node || ' -> ' || T5.to_node || ' -> ' || T6.to_node || ' -> ' || T7.to_node || ' -> ' || T8.to_node || ' -> ' || T9.to_node || ' -> ' || T10.to_node || ' -> ' || T11.to_node || ' -> ' || T12.to_node, 12
FROM mountain_network T1
JOIN mountain_network T2 ON T1.to_node = T2.from_node
JOIN mountain_network T3 ON T2.to_node = T3.from_node
JOIN mountain_network T4 ON T3.to_node = T4.from_node
JOIN mountain_network T5 ON T4.to_node = T5.from_node
JOIN mountain_network T6 ON T5.to_node = T6.from_node
JOIN mountain_network T7 ON T6.to_node = T7.from_node
JOIN mountain_network T8 ON T7.to_node = T8.from_node
JOIN mountain_network T9 ON T8.to_node = T9.from_node
JOIN mountain_network T10 ON T9.to_node = T10.from_node
JOIN mountain_network T11 ON T10.to_node = T11.from_node
JOIN mountain_network T12 ON T11.to_node = T12.from_node
WHERE T1.from_node = 'Jake''s Lift' AND T12.to_node = 'Maverick';
```

```
+--------------------------------------------------------------+-------------+
|                             path                             | connections |
+--------------------------------------------------------------+-------------+
| Jake's Lift -> Maverick                                      | 1           |
+--------------------------------------------------------------+-------------+
| Jake's Lift -> Wildwood Lift -> Brightside -> Moonrise -> Dr | 5           |
| aw -> Maverick                                               |             |
+--------------------------------------------------------------+-------------+
sqlite>
```


Ok, we get the result, but that looks like a terrifying query.

Can we do better?

I looked at the tutorial and we have something like `RECURSIVE CTE`

Wow!

This is quite challenging to explain.


### Recursive CTE

We can create a recursive CTE that draws the path until we have not found the stop node that is, `Maverick` and the number of connections is not more than `12`.

So here's how it goes.


```sql
WITH RECURSIVE ski_paths AS (
    SELECT
        from_node,
        to_node,
        CAST(from_node || ' -> ' || to_node AS TEXT) AS full_path,
        1 AS connections
    FROM mountain_network
    WHERE from_node = 'Jake''s Lift'

    UNION ALL

    SELECT
        mn.from_node,
        mn.to_node,
        sp.full_path || ' -> ' || mn.to_node,
        sp.connections + 1
    FROM ski_paths sp
    JOIN mountain_network mn ON sp.to_node = mn.from_node
    WHERE sp.connections < 12
      AND sp.to_node != 'Maverick'
)
SELECT full_path, connections
FROM ski_paths
WHERE to_node = 'Maverick'
ORDER BY connections ASC;
```
We first define the base case i.e. what to select first. We get the first row where we start, the `Jake's Lift` and from there on we create a recursive select statement by referencing the CTE within it.

The recursive bit is this one 

```sql
 SELECT
        mn.from_node,
        mn.to_node,
        sp.full_path || ' -> ' || mn.to_node,
        sp.connections + 1
    FROM ski_paths sp
    JOIN mountain_network mn ON sp.to_node = mn.from_node
    WHERE sp.connections < 12
      AND sp.to_node != 'Maverick'
```

We are referencing the CTE (table) within it as the recursive condition to select from the nodes where the end node or the `to_node` is not Maverick and connections haven't gone beyond `12`, we try to find the path.

This would give us all the nodes.

Let's see first, by checking the CTE as is.

```sql
WITH RECURSIVE ski_paths AS (
    SELECT
        from_node,
        to_node,
        CAST(from_node || ' -> ' || to_node AS TEXT) AS full_path,
        1 AS connections
    FROM mountain_network
    WHERE from_node = 'Jake''s Lift'

    UNION ALL

    SELECT
        mn.from_node,
        mn.to_node,
        sp.full_path || ' -> ' || mn.to_node,
        sp.connections + 1
    FROM ski_paths sp
    JOIN mountain_network mn ON sp.to_node = mn.from_node
    WHERE sp.connections < 12
      AND sp.to_node != 'Maverick'
)
SELECT * from ski_paths;

```

```
sqlite> WITH RECURSIVE ski_paths AS (
    SELECT
        from_node,
        to_node,
        CAST(from_node || ' -> ' || to_node AS TEXT) AS full_path,
        1 AS connections
    FROM mountain_network
    WHERE from_node = 'Jake''s Lift'

    UNION ALL

    SELECT
        mn.from_node,
        mn.to_node,
        sp.full_path || ' -> ' || mn.to_node,
        sp.connections + 1
    FROM ski_paths sp
    JOIN mountain_network mn ON sp.to_node = mn.from_node
    WHERE sp.connections < 12
      AND sp.to_node != 'Maverick'
)
SELECT * from ski_paths;
+---------------+----------------+--------------------------------------------------------------+-------------+
|   from_node   |    to_node     |                          full_path                           | connections |
+---------------+----------------+--------------------------------------------------------------+-------------+
| Jake's Lift   | Wildwood Lift  | Jake's Lift -> Wildwood Lift                                 | 1           |
+---------------+----------------+--------------------------------------------------------------+-------------+
| Jake's Lift   | Maverick       | Jake's Lift -> Maverick                                      | 1           |
+---------------+----------------+--------------------------------------------------------------+-------------+
| Wildwood Lift | Brightside     | Jake's Lift -> Wildwood Lift -> Brightside                   | 2           |
+---------------+----------------+--------------------------------------------------------------+-------------+
| Wildwood Lift | Sidewinder     | Jake's Lift -> Wildwood Lift -> Sidewinder                   | 2           |
+---------------+----------------+--------------------------------------------------------------+-------------+
| Brightside    | Moonrise       | Jake's Lift -> Wildwood Lift -> Brightside -> Moonrise       | 3           |
+---------------+----------------+--------------------------------------------------------------+-------------+
| Moonrise      | Draw           | Jake's Lift -> Wildwood Lift -> Brightside -> Moonrise -> Dr | 4           |
|               |                | aw                                                           |             |
+---------------+----------------+--------------------------------------------------------------+-------------+
| Moonrise      | Lone Pine      | Jake's Lift -> Wildwood Lift -> Brightside -> Moonrise -> Lo | 4           |
|               |                | ne Pine                                                      |             |
+---------------+----------------+--------------------------------------------------------------+-------------+
| Draw          | Broadway       | Jake's Lift -> Wildwood Lift -> Brightside -> Moonrise -> Dr | 5           |
|               |                | aw -> Broadway                                               |             |
+---------------+----------------+--------------------------------------------------------------+-------------+
| Draw          | Maverick       | Jake's Lift -> Wildwood Lift -> Brightside -> Moonrise -> Dr | 5           |
|               |                | aw -> Maverick                                               |             |
+---------------+----------------+--------------------------------------------------------------+-------------+
| Broadway      | Outlaw Trail   | Jake's Lift -> Wildwood Lift -> Brightside -> Moonrise -> Dr | 6           |
|               |                | aw -> Broadway -> Outlaw Trail                               |             |
+---------------+----------------+--------------------------------------------------------------+-------------+
| Outlaw Trail  | Center Aisle   | Jake's Lift -> Wildwood Lift -> Brightside -> Moonrise -> Dr | 7           |
|               |                | aw -> Broadway -> Outlaw Trail -> Center Aisle               |             |
+---------------+----------------+--------------------------------------------------------------+-------------+
| Center Aisle  | Bandit         | Jake's Lift -> Wildwood Lift -> Brightside -> Moonrise -> Dr | 8           |
|               |                | aw -> Broadway -> Outlaw Trail -> Center Aisle -> Bandit     |             |
+---------------+----------------+--------------------------------------------------------------+-------------+
| Center Aisle  | Lower Stampede | Jake's Lift -> Wildwood Lift -> Brightside -> Moonrise -> Dr | 8           |
|               |                | aw -> Broadway -> Outlaw Trail -> Center Aisle -> Lower Stam |             |
|               |                | pede                                                         |             |
+---------------+----------------+--------------------------------------------------------------+-------------+
sqlite>
```

You can see we have got quite a lot of paths, this are all the paths that start from `Jake's Lift` and have less than `12` connections.

So, now we can simply filter with the `WHERE` clause in the case where the `to_node = 'Maverick'` and we would get the result.



```
sqlite> WITH RECURSIVE ski_paths AS (
    SELECT
        from_node,
        to_node,
        CAST(from_node || ' -> ' || to_node AS TEXT) AS full_path,
        1 AS connections
    FROM mountain_network
    WHERE from_node = 'Jake''s Lift'

    UNION ALL

    SELECT
        mn.from_node,
        mn.to_node,
        sp.full_path || ' -> ' || mn.to_node,
        sp.connections + 1
    FROM ski_paths sp
    JOIN mountain_network mn ON sp.to_node = mn.from_node
    WHERE sp.connections < 12
      AND sp.to_node != 'Maverick'
)
SELECT full_path, connections
FROM ski_paths
WHERE to_node = 'Maverick'
ORDER BY connections ASC;
+--------------------------------------------------------------+-------------+
|                          full_path                           | connections |
+--------------------------------------------------------------+-------------+
| Jake's Lift -> Maverick                                      | 1           |
+--------------------------------------------------------------+-------------+
| Jake's Lift -> Wildwood Lift -> Brightside -> Moonrise -> Dr | 5           |
| aw -> Maverick                                               |             |
+--------------------------------------------------------------+-------------+
sqlite>
```

Ok, that gave the result correctly.

And this is it!

Sweet problem to learn about Recursive CTE!

Off to the final day of advent of SQL 2025!
