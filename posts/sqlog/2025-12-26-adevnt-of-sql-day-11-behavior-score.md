---
type: "sqlog"
title: "Advent of SQL 2025 Day 11: Behavior Score"
slug: "advent-of-sql-2025-day-11"
date: 2025-12-26 15:30 +0530
series: ["advent-of-sql-2025"]
tags: ["sqlite", "sql", "advent-of-sql"]
---

## Advent of SQL - Day 11, Behavior Score

All right, this is day 11 from Advent of SQL.

Let's pull in the data.

```sql
DROP TABLE IF EXISTS behavior_logs;

CREATE TABLE behavior_logs (
    id INT PRIMARY KEY,
    child_id INT,
    child_name TEXT,
    behavior_date DATE,
    score INT
);

INSERT INTO behavior_logs (id, child_id, child_name, behavior_date, score) VALUES
    (1, 1, 'Emma D.', '2025-12-01', 5),
    (2, 1, 'Emma D.', '2025-12-02', 1),
    (3, 1, 'Emma D.', '2025-12-03', 3),
    (4, 1, 'Emma D.', '2025-12-04', 5),
    (5, 1, 'Emma D.', '2025-12-05', 2),
    (6, 1, 'Emma D.', '2025-12-06', 2),
    (7, 1, 'Emma D.', '2025-12-07', 3),
    (8, 1, 'Emma D.', '2025-12-08', 5),
    (9, 1, 'Emma D.', '2025-12-09', 4),
    (10, 1, 'Emma D.', '2025-12-10', 5),
    (11, 1, 'Emma D.', '2025-12-11', 5),
    (12, 1, 'Emma D.', '2025-12-12', -1),
    (13, 1, 'Emma D.', '2025-12-13', 1),
    (14, 1, 'Emma D.', '2025-12-14', 1),
    (15, 1, 'Emma D.', '2025-12-15', -1),
    (16, 1, 'Emma D.', '2025-12-16', 3),
    (17, 1, 'Emma D.', '2025-12-17', -2),
    (18, 1, 'Emma D.', '2025-12-18', 1),
    (19, 1, 'Emma D.', '2025-12-19', 1),
    (20, 1, 'Emma D.', '2025-12-20', -2),
    (21, 2, 'Ava X.', '2025-12-01', 0),
    (22, 2, 'Ava X.', '2025-12-02', -1),
    (23, 2, 'Ava X.', '2025-12-03', 4),
    (24, 2, 'Ava X.', '2025-12-04', 0),
    (25, 2, 'Ava X.', '2025-12-05', 2),
    (26, 2, 'Ava X.', '2025-12-06', 3),
    (27, 2, 'Ava X.', '2025-12-07', 5),
    (28, 2, 'Ava X.', '2025-12-08', 2),
    (29, 2, 'Ava X.', '2025-12-09', 1),
    (30, 2, 'Ava X.', '2025-12-10', 5),
    (31, 2, 'Ava X.', '2025-12-11', 2),
    (32, 2, 'Ava X.', '2025-12-12', 5),
    (33, 2, 'Ava X.', '2025-12-13', 5),
    (34, 2, 'Ava X.', '2025-12-14', 2),
    (35, 2, 'Ava X.', '2025-12-15', 0),
    (36, 2, 'Ava X.', '2025-12-16', 0),
    (37, 2, 'Ava X.', '2025-12-17', 5),
    (38, 2, 'Ava X.', '2025-12-18', 4),
    (39, 2, 'Ava X.', '2025-12-19', 5),
    (40, 2, 'Ava X.', '2025-12-20', 5),
    (181, 10, 'Ava C.', '2025-12-01', 3),
    (182, 10, 'Ava C.', '2025-12-02', 0),
    (183, 10, 'Ava C.', '2025-12-03', 3),
    (184, 10, 'Ava C.', '2025-12-04', 5),
    (185, 10, 'Ava C.', '2025-12-05', 5),
    (186, 10, 'Ava C.', '2025-12-06', 4),
    (187, 10, 'Ava C.', '2025-12-07', 1),
    (188, 10, 'Ava C.', '2025-12-08', 4),
    (189, 10, 'Ava C.', '2025-12-09', 5),
    (190, 10, 'Ava C.', '2025-12-10', 5),
    (191, 10, 'Ava C.', '2025-12-11', 5),
    (192, 10, 'Ava C.', '2025-12-12', 0),
    (193, 10, 'Ava C.', '2025-12-13', 0),
    (194, 10, 'Ava C.', '2025-12-14', 3),
    (195, 10, 'Ava C.', '2025-12-15', 1),
    (196, 10, 'Ava C.', '2025-12-16', 3),
    (197, 10, 'Ava C.', '2025-12-17', -1),
    (198, 10, 'Ava C.', '2025-12-18', 0),
    (199, 10, 'Ava C.', '2025-12-19', 5),
    (200, 10, 'Ava C.', '2025-12-20', 4),
    (241, 13, 'Ava R.', '2025-12-01', 3),
    (242, 13, 'Ava R.', '2025-12-02', 2),
    (243, 13, 'Ava R.', '2025-12-03', 2),
    (244, 13, 'Ava R.', '2025-12-04', 1),
    (245, 13, 'Ava R.', '2025-12-05', -1),
    (246, 13, 'Ava R.', '2025-12-06', -1),
    (247, 13, 'Ava R.', '2025-12-07', 2),
    (248, 13, 'Ava R.', '2025-12-08', 5),
    (249, 13, 'Ava R.', '2025-12-09', 0),
    (250, 13, 'Ava R.', '2025-12-10', 5),
    (251, 13, 'Ava R.', '2025-12-11', 2),
    (252, 13, 'Ava R.', '2025-12-12', -1),
    (253, 13, 'Ava R.', '2025-12-13', 2),
    (254, 13, 'Ava R.', '2025-12-14', 3),
    (255, 13, 'Ava R.', '2025-12-15', 2),
    (256, 13, 'Ava R.', '2025-12-16', -1),
    (257, 13, 'Ava R.', '2025-12-17', -2),
    (258, 13, 'Ava R.', '2025-12-18', -4),
    (259, 13, 'Ava R.', '2025-12-19', -3),
    (260, 13, 'Ava R.', '2025-12-20', 2),
    (1961, 99, 'Ava X.', '2025-12-01', 2),
    (1962, 99, 'Ava X.', '2025-12-02', -2),
    (1963, 99, 'Ava X.', '2025-12-03', -1),
    (1964, 99, 'Ava X.', '2025-12-04', -2),
    (1965, 99, 'Ava X.', '2025-12-05', 3),
    (1966, 99, 'Ava X.', '2025-12-06', -1),
    (1967, 99, 'Ava X.', '2025-12-07', 0),
    (1968, 99, 'Ava X.', '2025-12-08', 1),
    (1969, 99, 'Ava X.', '2025-12-09', 0),
    (1970, 99, 'Ava X.', '2025-12-10', 0),
    (1971, 99, 'Ava X.', '2025-12-11', 3),
    (1972, 99, 'Ava X.', '2025-12-12', 4),
    (1973, 99, 'Ava X.', '2025-12-13', 4),
    (1974, 99, 'Ava X.', '2025-12-14', 0),
    (1975, 99, 'Ava X.', '2025-12-15', 3),
    (1976, 99, 'Ava X.', '2025-12-16', -1),
    (1977, 99, 'Ava X.', '2025-12-17', -1),
    (1978, 99, 'Ava X.', '2025-12-18', 3),
    (1979, 99, 'Ava X.', '2025-12-19', 3),
    (1980, 99, 'Ava X.', '2025-12-20', -3);
```

No hiccups! Good to go.

We just have one table today.

```sql
 SELECT * FROM behavior_logs;
```

```
sqlite> .read day11-inserts.sql
sqlite> .mode table
sqlite> .schema
CREATE TABLE behavior_logs (
    id INT PRIMARY KEY,
    child_id INT,
    child_name TEXT,
    behavior_date DATE,
    score INT
);
sqlite> SELECT * FROM behavior_logs LIMIT 20;
+----+----------+------------+---------------+-------+
| id | child_id | child_name | behavior_date | score |
+----+----------+------------+---------------+-------+
| 1  | 1        | Emma D.    | 2025-12-01    | 5     |
| 2  | 1        | Emma D.    | 2025-12-02    | 1     |
| 3  | 1        | Emma D.    | 2025-12-03    | 3     |
| 4  | 1        | Emma D.    | 2025-12-04    | 5     |
| 5  | 1        | Emma D.    | 2025-12-05    | 2     |
| 6  | 1        | Emma D.    | 2025-12-06    | 2     |
| 7  | 1        | Emma D.    | 2025-12-07    | 3     |
| 8  | 1        | Emma D.    | 2025-12-08    | 5     |
| 9  | 1        | Emma D.    | 2025-12-09    | 4     |
| 10 | 1        | Emma D.    | 2025-12-10    | 5     |
| 11 | 1        | Emma D.    | 2025-12-11    | 5     |
| 12 | 1        | Emma D.    | 2025-12-12    | -1    |
| 13 | 1        | Emma D.    | 2025-12-13    | 1     |
| 14 | 1        | Emma D.    | 2025-12-14    | 1     |
| 15 | 1        | Emma D.    | 2025-12-15    | -1    |
| 16 | 1        | Emma D.    | 2025-12-16    | 3     |
| 17 | 1        | Emma D.    | 2025-12-17    | -2    |
| 18 | 1        | Emma D.    | 2025-12-18    | 1     |
| 19 | 1        | Emma D.    | 2025-12-19    | 1     |
| 20 | 1        | Emma D.    | 2025-12-20    | -2    |
+----+----------+------------+---------------+-------+
sqlite> SELECT * FROM behavior_logs WHERE child_name LIKE 'Ava';
sqlite> SELECT * FROM behavior_logs WHERE child_name LIKE 'Ava%';
+------+----------+------------+---------------+-------+
|  id  | child_id | child_name | behavior_date | score |
+------+----------+------------+---------------+-------+
| 21   | 2        | Ava X.     | 2025-12-01    | 0     |
| 22   | 2        | Ava X.     | 2025-12-02    | -1    |
| 23   | 2        | Ava X.     | 2025-12-03    | 4     |
| 24   | 2        | Ava X.     | 2025-12-04    | 0     |
| 25   | 2        | Ava X.     | 2025-12-05    | 2     |
| 26   | 2        | Ava X.     | 2025-12-06    | 3     |
| 27   | 2        | Ava X.     | 2025-12-07    | 5     |
| 28   | 2        | Ava X.     | 2025-12-08    | 2     |
| 29   | 2        | Ava X.     | 2025-12-09    | 1     |
| 30   | 2        | Ava X.     | 2025-12-10    | 5     |
| 31   | 2        | Ava X.     | 2025-12-11    | 2     |
| 32   | 2        | Ava X.     | 2025-12-12    | 5     |
| 33   | 2        | Ava X.     | 2025-12-13    | 5     |
| 34   | 2        | Ava X.     | 2025-12-14    | 2     |
| 35   | 2        | Ava X.     | 2025-12-15    | 0     |
| 36   | 2        | Ava X.     | 2025-12-16    | 0     |
| 37   | 2        | Ava X.     | 2025-12-17    | 5     |
| 38   | 2        | Ava X.     | 2025-12-18    | 4     |
| 39   | 2        | Ava X.     | 2025-12-19    | 5     |
| 40   | 2        | Ava X.     | 2025-12-20    | 5     |
| 181  | 10       | Ava C.     | 2025-12-01    | 3     |
| 182  | 10       | Ava C.     | 2025-12-02    | 0     |
| 183  | 10       | Ava C.     | 2025-12-03    | 3     |
| 184  | 10       | Ava C.     | 2025-12-04    | 5     |
| 185  | 10       | Ava C.     | 2025-12-05    | 5     |
| 186  | 10       | Ava C.     | 2025-12-06    | 4     |
| 187  | 10       | Ava C.     | 2025-12-07    | 1     |
| 188  | 10       | Ava C.     | 2025-12-08    | 4     |
| 189  | 10       | Ava C.     | 2025-12-09    | 5     |
| 190  | 10       | Ava C.     | 2025-12-10    | 5     |
| 191  | 10       | Ava C.     | 2025-12-11    | 5     |
| 192  | 10       | Ava C.     | 2025-12-12    | 0     |
| 193  | 10       | Ava C.     | 2025-12-13    | 0     |
| 194  | 10       | Ava C.     | 2025-12-14    | 3     |
| 195  | 10       | Ava C.     | 2025-12-15    | 1     |
| 196  | 10       | Ava C.     | 2025-12-16    | 3     |
| 197  | 10       | Ava C.     | 2025-12-17    | -1    |
| 198  | 10       | Ava C.     | 2025-12-18    | 0     |
| 199  | 10       | Ava C.     | 2025-12-19    | 5     |
| 200  | 10       | Ava C.     | 2025-12-20    | 4     |
| 241  | 13       | Ava R.     | 2025-12-01    | 3     |
| 242  | 13       | Ava R.     | 2025-12-02    | 2     |
| 243  | 13       | Ava R.     | 2025-12-03    | 2     |
| 244  | 13       | Ava R.     | 2025-12-04    | 1     |
| 245  | 13       | Ava R.     | 2025-12-05    | -1    |
| 246  | 13       | Ava R.     | 2025-12-06    | -1    |
| 247  | 13       | Ava R.     | 2025-12-07    | 2     |
| 248  | 13       | Ava R.     | 2025-12-08    | 5     |
| 249  | 13       | Ava R.     | 2025-12-09    | 0     |
| 250  | 13       | Ava R.     | 2025-12-10    | 5     |
| 251  | 13       | Ava R.     | 2025-12-11    | 2     |
| 252  | 13       | Ava R.     | 2025-12-12    | -1    |
| 253  | 13       | Ava R.     | 2025-12-13    | 2     |
| 254  | 13       | Ava R.     | 2025-12-14    | 3     |
| 255  | 13       | Ava R.     | 2025-12-15    | 2     |
| 256  | 13       | Ava R.     | 2025-12-16    | -1    |
| 257  | 13       | Ava R.     | 2025-12-17    | -2    |
| 258  | 13       | Ava R.     | 2025-12-18    | -4    |
| 259  | 13       | Ava R.     | 2025-12-19    | -3    |
| 260  | 13       | Ava R.     | 2025-12-20    | 2     |
| 1961 | 99       | Ava X.     | 2025-12-01    | 2     |
| 1962 | 99       | Ava X.     | 2025-12-02    | -2    |
| 1963 | 99       | Ava X.     | 2025-12-03    | -1    |
| 1964 | 99       | Ava X.     | 2025-12-04    | -2    |
| 1965 | 99       | Ava X.     | 2025-12-05    | 3     |
| 1966 | 99       | Ava X.     | 2025-12-06    | -1    |
| 1967 | 99       | Ava X.     | 2025-12-07    | 0     |
| 1968 | 99       | Ava X.     | 2025-12-08    | 1     |
| 1969 | 99       | Ava X.     | 2025-12-09    | 0     |
| 1970 | 99       | Ava X.     | 2025-12-10    | 0     |
| 1971 | 99       | Ava X.     | 2025-12-11    | 3     |
| 1972 | 99       | Ava X.     | 2025-12-12    | 4     |
| 1973 | 99       | Ava X.     | 2025-12-13    | 4     |
| 1974 | 99       | Ava X.     | 2025-12-14    | 0     |
| 1975 | 99       | Ava X.     | 2025-12-15    | 3     |
| 1976 | 99       | Ava X.     | 2025-12-16    | -1    |
| 1977 | 99       | Ava X.     | 2025-12-17    | -1    |
| 1978 | 99       | Ava X.     | 2025-12-18    | 3     |
| 1979 | 99       | Ava X.     | 2025-12-19    | 3     |
| 1980 | 99       | Ava X.     | 2025-12-20    | -3    |
+------+----------+------------+---------------+-------+
sqlite> 
```

Let's get to the problem of day 11

## Problem

> Calculate the 7-day rolling average behavior score for each child. Identify any child whose rolling average drops below 0. For those children with a rolling average below 0, return the `child_id`, `child_name`, `behavior_date` (this will be the latest date in the 7-day rolling average), and the calculated 7-day rolling average. Only include results with a `behavior_date` of `December 7, 2025` or later, ensuring that each rolling average is based on a full 7 days of data.
> 
> Order the results by `behavior_date` and then `child_name`.

So, we need to do what?

- Group by child_id
- Compute the rolling average for the past 7 days (so if there is a data for 20 days from 1st to 20th December, we'll only consider the average from 14th to 20th December)
- Order by behavior_date and child_name.

### Using Simple Join

We have to first grab the rolling average which is only for the past 7 days per child.

So, to do that, we can self join the `behavior_log` table on the condition that the behavior date (from right table) is between the current behavior date (from left table) and 6 days past that for each child.

```sql
SELECT 
    behavior_logs.child_id,
    behavior_logs.child_name,
    behavior_logs.behavior_date,
    AVG(week_logs.score) AS rolling_avg
FROM 
    behavior_logs
JOIN 
    behavior_logs week_logs
    ON behavior_logs.child_id = week_logs.child_id
    AND week_logs.behavior_date BETWEEN DATE(behavior_logs.behavior_date, '-6 days') AND behavior_logs.behavior_date
GROUP BY 
    behavior_logs.child_id;
```

We basically joined the table `behavior_logs` with itself, i.e. the right and left tables are the same. So, we join on the condition of same child_id, and then look for dates between the past 6 days and the current behavior date (we included the 7th day in the behavior date, so hence subtracting 6 days from that date).

We group by `child_id`, so that we get a single row for each `child_id`. We compute the `AVG(score)` to get the average score for each day in the range of the past 7 days.

You see the issue, we need to compute the average for each week, and not just the last, it will give us the last most score for a child. We need to also group by the behavior date so as to keep week unique and have multiple entries for a child.

```sql
SELECT 
    behavior_logs.child_id,
    behavior_logs.child_name,
    behavior_logs.behavior_date,
    AVG(week_logs.score) AS rolling_avg
FROM 
    behavior_logs
JOIN 
    behavior_logs week_logs
    ON behavior_logs.child_id = week_logs.child_id
    AND week_logs.behavior_date BETWEEN DATE(behavior_logs.behavior_date, '-6 days') AND behavior_logs.behavior_date
GROUP BY 
    behavior_logs.child_id, behavior_logs.behavior_date;
```

Now, we have grouped by `behavior_date` so that each child can have multiple entries for the weeks that we have logs for. Now we can further filter it.

We also need to filter when the `rolling_average` is less than 0, that is the child had a bad week overall. So that will be in a `HAVING` condition and not a `WHERE` condition since `AVG` is a aggregate function, we can't reference the `rolling_average` in the where clause, it won't be defined there yet. So, we use `HAVING` to filter `rolling_average` as less than `0`.

```sql
SELECT 
    behavior_logs.child_id,
    behavior_logs.child_name,
    behavior_logs.behavior_date,
    AVG(week_logs.score) AS rolling_avg
FROM 
    behavior_logs
JOIN 
    behavior_logs week_logs
    ON behavior_logs.child_id = week_logs.child_id
    AND week_logs.behavior_date BETWEEN DATE(behavior_logs.behavior_date, '-6 days') AND behavior_logs.behavior_date
GROUP BY 
    behavior_logs.child_id, behavior_logs.behavior_date
HAVING 
    rolling_avg < 0;
```

We are now down to the rows that only have `bad` weeks for the child in various weeks. 

There is one catch however, we can't use week periods before `7th December` because we don't have enough data to compute the rolling average for the 7 day score. Hence we only include the records having logs after `7th December`.

This again would come in the `HAVING` clause as we are deciding the final behavior date from the week and not the individual logs from the table.

```sql
SELECT 
    behavior_logs.child_id,
    behavior_logs.child_name,
    behavior_logs.behavior_date,
    AVG(week_logs.score) AS rolling_avg
FROM 
    behavior_logs
JOIN 
    behavior_logs week_logs
    ON behavior_logs.child_id = week_logs.child_id
    AND week_logs.behavior_date BETWEEN DATE(behavior_logs.behavior_date, '-6 days') AND behavior_logs.behavior_date
GROUP BY 
    behavior_logs.child_id, behavior_logs.behavior_date
HAVING 
    rolling_avg < 0
    AND behavior_logs.behavior_date >= '2025-12-07'

```

Now, the final piece is the order.

We need to order it by the `behavior_date` and the `child_name` as requested.

```sql
SELECT 
    behavior_logs.child_id,
    behavior_logs.child_name,
    behavior_logs.behavior_date,
    AVG(week_logs.score) AS rolling_avg
FROM 
    behavior_logs
JOIN 
    behavior_logs week_logs
    ON behavior_logs.child_id = week_logs.child_id
    AND week_logs.behavior_date BETWEEN DATE(behavior_logs.behavior_date, '-6 days') AND behavior_logs.behavior_date
GROUP BY 
    behavior_logs.child_id, behavior_logs.behavior_date
HAVING 
    rolling_avg < 0
    AND behavior_logs.behavior_date >= '2025-12-07'
ORDER BY 
    behavior_logs.behavior_date, behavior_logs.child_name;
```

This is a simple solution.
Easy to understand, and explain, not quite short and crisp.


We can even do this in a sub query without a JOIN like so

### Using Sub-query

We take the conditions from the `JOIN` and replace that in a sub-query like so.

```sql
SELECT
    current_logs.child_id,
    current_logs.child_name,
    current_logs.behavior_date,
    (
        SELECT AVG(past_logs.score)
        FROM behavior_logs past_logs
        WHERE past_logs.child_id = current_logs.child_id
          AND past_logs.behavior_date
              BETWEEN date(current_logs.behavior_date, '-6 days')
                  AND current_logs.behavior_date
    ) AS rolling_avg
FROM behavior_logs current_logs
WHERE current_logs.behavior_date >= '2025-12-07'
  AND rolling_avg < 0
ORDER BY current_logs.behavior_date, current_logs.child_name;

```

Since its a subquery, we don't need to group by or add a having clause to filter the rolling average and the behavior date.

### Using Sub-query and Window Function

We can take the above query and instead of join, we can write a sub-query to compute the rolling average using window function.

```sql
SELECT
    child_id,
    child_name,
    behavior_date,
    AVG(score) OVER (
        PARTITION BY child_id
        ORDER BY behavior_date
        ROWS BETWEEN 6 PRECEDING AND CURRENT ROW
    ) AS rolling_avg
FROM behavior_logs

```

We use the `AVG(score) OVER ()` this is a partition. We partition or create a window for each child and we order by the behavior date and then create a sliding window for the past 7 days. 

The `ROWS BETWEEN 6 PRECEDING AND CURRENT ROW` defines a 7-row sliding window for each child. Every row gets its own window, and each window is separate per child because of PARTITION BY `child_id`.

So we get a full per day rolling average for each child with this query.

This had around 2400 rows.

Now, we need filter it down to the only rows where the rolling average is less than `0` and we don't include the average for days before the `7th December` as we don't have enough days before that to compute the 7-day rolling average.

But this is a query in itself we can't reference the `rolling_average` in the where clause as its not available there. We haven't grouped by anything explicitly so we can't use `HAVING`, we need to wrap it in a sub query to get the rolling average as well as the behavior date condition.

```sql
SELECT
    child_id,
    child_name,
    behavior_date,
    rolling_avg
FROM (
    SELECT
        child_id,
        child_name,
        behavior_date,
        AVG(score) OVER (
            PARTITION BY child_id
            ORDER BY behavior_date
            ROWS BETWEEN 6 PRECEDING AND CURRENT ROW
        ) AS rolling_avg
    FROM behavior_logs
)
WHERE behavior_date >= '2025-12-07'
  AND rolling_avg < 0;
```

This will filter down the rows.

Now, we also need to order by the `behavior_date` and then `child_name`.

```sql
SELECT
    child_id,
    child_name,
    behavior_date,
    rolling_avg
FROM (
    SELECT
        child_id,
        child_name,
        behavior_date,
        AVG(score) OVER (
            PARTITION BY child_id
            ORDER BY behavior_date
            ROWS BETWEEN 6 PRECEDING AND CURRENT ROW
        ) AS rolling_avg
    FROM behavior_logs
)
WHERE behavior_date >= '2025-12-07'
  AND rolling_avg < 0
ORDER BY behavior_date, child_name;
```

So, that is again the same stuff just with window function.

We can even do this with that subquery wrapped in a CTE (it just looks and reads good, nothing different really, other than you can use the CTE in the same query multiple times, here we don't need to do it)

### Using CTE and Window Function

We just take the above sub-query and wrap it in a CTE.

```sql
WITH rolling AS (
    SELECT
        child_id,
        child_name,
        behavior_date,
        AVG(score) OVER (
            PARTITION BY child_id
            ORDER BY behavior_date
            ROWS BETWEEN 6 PRECEDING AND CURRENT ROW
        ) AS rolling_avg
    FROM behavior_logs
)
SELECT
    child_id,
    child_name,
    behavior_date,
    rolling_avg
FROM rolling
WHERE behavior_date >= '2025-12-07'
  AND rolling_avg < 0
ORDER BY behavior_date, child_name;

```

The rest remains the same, we just reference the `rolling` as the CTE and grab the necessary details in the query using that CTE.

That should be it for the Day 11!

Some cool CTE, Window function and pratical use case for computing rolling averages, loved it!

On to the day 12!

