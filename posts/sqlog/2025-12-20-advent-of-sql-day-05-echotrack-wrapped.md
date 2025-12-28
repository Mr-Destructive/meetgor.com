---
type: "sqlog"
title: "Advent of SQL 2025 Day 5: EchoTrack Wrapped"
slug: "advent-of-sql-2025-day-5"
date: 2025-12-20 15:30 +0530
series: ["advent-of-sql-2025"]
tags: ["sqlite", "sql", "advent-of-sql"]
---

## Advent of SQL Day 5 - EchoTrack Wrapped

It is day 5 of advent of SQL.

Let's get rollin. It looks like a good problem. I am excited!

Here's the SQL to get started.

```sql
DROP TABLE IF EXISTS listening_logs;

CREATE TABLE listening_logs (
    id INTEGER PRIMARY KEY,
    user_name TEXT,
    artist TEXT,
    played_at TIMESTAMP,
    content_type TEXT
);

INSERT INTO listening_logs (id, user_name, artist, played_at, content_type) VALUES
    (1, 'Zoe Garcia', 'Arijit Singh', '2025-04-08 00:21:53', 'song'),
    (2, 'Zoe Garcia', 'Huberman Lab', '2025-11-10 19:18:47', 'podcast'),
    (3, 'Zoe Garcia', 'Huberman Lab', '2025-01-20 15:31:02', 'podcast'),
    (4, 'Zoe Garcia', 'Arijit Singh', '2025-01-06 17:33:11', 'song'),
    (5, 'Zoe Garcia', 'Candace', '2025-03-06 14:07:54', 'podcast'),
    (6, 'Zoe Garcia', 'Arijit Singh', '2025-06-05 17:57:59', 'song'),
    (7, 'Zoe Garcia', 'Huberman Lab', '2025-01-01 20:05:22', 'podcast'),
    (8, 'Zoe Garcia', 'Huberman Lab', '2025-11-01 12:04:03', 'podcast'),
    (9, 'Zoe Garcia', 'Arijit Singh', '2025-09-28 12:42:12', 'song'),
    (10, 'Zoe Garcia', 'The Ben Shapiro Show', '2025-09-15 01:05:15', 'podcast'),
    (11, 'Zoe Garcia', 'Arijit Singh', '2025-04-26 05:31:02', 'song'),
    (12, 'Zoe Garcia', 'Arijit Singh', '2025-10-13 17:34:03', 'song'),
    (13, 'Zoe Garcia', 'Mariah Carey', '2025-01-20 11:21:37', 'song'),
    (14, 'Zoe Garcia', 'Arijit Singh', '2025-11-28 03:55:31', 'song'),
    (15, 'Zoe Garcia', 'Arijit Singh', '2025-07-17 05:18:16', 'song'),
    (16, 'Zoe Garcia', 'Arijit Singh', '2025-08-20 02:07:45', 'song'),
    (17, 'Zoe Garcia', 'Kendrick Lamar', '2025-02-16 13:25:27', 'song'),
    (18, 'Zoe Garcia', 'Huberman Lab', '2025-08-13 19:55:00', 'podcast'),
    (19, 'Zoe Garcia', 'Bruno Mars', '2025-09-13 07:09:43', 'song'),
    (20, 'Zoe Garcia', 'Arijit Singh', '2025-04-12 06:30:44', 'song');
```

Let's open a SQLite shell and get started.

```
$ sqlite3
SQLite version 3.50.4 2025-07-30 19:33:53
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
sqlite> .read day5-inserts.sql
sqlite> .schema
CREATE TABLE listening_logs (
    id INTEGER PRIMARY KEY,
    user_name TEXT,
    artist TEXT,
    played_at TIMESTAMP,
    content_type TEXT
);
sqlite> .mode table
sqlite> SELECT * FROM listening_logs LIMIT 20;
+----+------------+----------------------+---------------------+--------------+
| id | user_name  |        artist        |      played_at      | content_type |
+----+------------+----------------------+---------------------+--------------+
| 1  | Zoe Garcia | Arijit Singh         | 2025-04-08 00:21:53 | song         |
| 2  | Zoe Garcia | Huberman Lab         | 2025-11-10 19:18:47 | podcast      |
| 3  | Zoe Garcia | Huberman Lab         | 2025-01-20 15:31:02 | podcast      |
| 4  | Zoe Garcia | Arijit Singh         | 2025-01-06 17:33:11 | song         |
| 5  | Zoe Garcia | Candace              | 2025-03-06 14:07:54 | podcast      |
| 6  | Zoe Garcia | Arijit Singh         | 2025-06-05 17:57:59 | song         |
| 7  | Zoe Garcia | Huberman Lab         | 2025-01-01 20:05:22 | podcast      |
| 8  | Zoe Garcia | Huberman Lab         | 2025-11-01 12:04:03 | podcast      |
| 9  | Zoe Garcia | Arijit Singh         | 2025-09-28 12:42:12 | song         |
| 10 | Zoe Garcia | The Ben Shapiro Show | 2025-09-15 01:05:15 | podcast      |
| 11 | Zoe Garcia | Arijit Singh         | 2025-04-26 05:31:02 | song         |
| 12 | Zoe Garcia | Arijit Singh         | 2025-10-13 17:34:03 | song         |
| 13 | Zoe Garcia | Mariah Carey         | 2025-01-20 11:21:37 | song         |
| 14 | Zoe Garcia | Arijit Singh         | 2025-11-28 03:55:31 | song         |
| 15 | Zoe Garcia | Arijit Singh         | 2025-07-17 05:18:16 | song         |
| 16 | Zoe Garcia | Arijit Singh         | 2025-08-20 02:07:45 | song         |
| 17 | Zoe Garcia | Kendrick Lamar       | 2025-02-16 13:25:27 | song         |
| 18 | Zoe Garcia | Huberman Lab         | 2025-08-13 19:55:00 | podcast      |
| 19 | Zoe Garcia | Bruno Mars           | 2025-09-13 07:09:43 | song         |
| 20 | Zoe Garcia | Arijit Singh         | 2025-04-12 06:30:44 | song         |
+----+------------+----------------------+---------------------+--------------+
sqlite>
sqlite> SELECT COUNT(*) FROM listening_logs;
+----------+
| COUNT(*) |
+----------+
| 18174    |
+----------+
sqlite>
```

OK! We have around 18k records in a single table! That's a lot but not not much!

Let's see what we have to do

## Problem

> Write a query that returns the top 3 artists per user. Order the results by the most played

Alright, this is quite a problem to solve, if you are thinking, it easy peasy, then hold on!

We clearly see, that we have 2 columns of our interest.
1. `user_name`
2. `artist`

We need to group for each user his most played artists, and we need to rank those top 3 artist per user.

And each entry is a song or podcast that the user has litened to.

We need to aggregate, group, and then rank, and then what?

How would you chunk out the top three?
It's time to put your SQL glasses and gloves on, it's getting colder!

### Counting Artists Per User

Let's take one step at a time, we need to first count how many times each artist has been played for that user.

```sql
SELECT
    user_name,
    artist,
    COUNT(*) AS play_count
FROM listening_logs
GROUP BY user_name, artist
ORDER BY user_name, play_count DESC, artist;
```
OK, simple right?

Select all the usernames, artist and then group by the username and the artist and count the number of times the user had played that artist. Simply then order by username, the playcount and the artist (if there is a tie in count, I think we can choose the artist name as the breaker)

But it gives all the artists, not just the top 3, it ranks them in the decreasing order of the number of plays, but we only want to list the top 3 per user.

That's tricky!

### With SELF JOIN

What if we can join the table with itself?

Then we can compare the number of times the user has played the artist, with the number of times the user has played another artist. Then we can remove the artist if it has played less than 3 times, this then will filter out the top 3 for us.

```sql
SELECT 
    a.user_name,
    a.artist AS current_artist,
    a.play_count AS current_plays,
    b.artist AS other_artist,
    b.play_count AS other_plays
FROM (
    SELECT user_name, artist, COUNT(*) AS play_count
    FROM listening_logs
    GROUP BY user_name, artist
) a
LEFT JOIN (
    SELECT user_name, artist, COUNT(*) AS play_count
    FROM listening_logs
    GROUP BY user_name, artist
) b ON a.user_name = b.user_name
ORDER BY a.user_name, a.play_count DESC, b.play_count DESC;
```

This would create a cross join of sorts between the same table.

- Select the required columns (user_name, artist and count)
Look at this part
This is `a`
```sql
-- a
SELECT user_name, artist, COUNT(*) AS play_count
FROM listening_logs
GROUP BY user_name, artist
ORDER BY user_name, play_count DESC, artist
```

Then we need to join this with itself

- Select the required columns from the same table (user_name, artist and count)
This is `b`
```sql
-- b
SELECT user_name, artist, COUNT(*) AS play_count
FROM listening_logs
GROUP BY user_name, artist
ORDER BY user_name, play_count DESC, artist
```
Both `a` and `b` are the same, just that we want a cross join of sorts.

And then we need to join `a` and `b`
- Join with itself on the user_name
- Order by user_name, play_count desc, artist

```sql
SELECT 
    a.user_name,
    a.artist AS current_artist,
    a.play_count AS current_plays,
    b.artist AS other_artist,
    b.play_count AS other_plays
FROM (
    SELECT user_name, artist, COUNT(*) AS play_count
    FROM listening_logs
    GROUP BY user_name, artist
) a
LEFT JOIN (
    SELECT user_name, artist, COUNT(*) AS play_count
    FROM listening_logs
    GROUP BY user_name, artist
) b ON a.user_name = b.user_name
ORDER BY a.user_name, a.play_count DESC, b.play_count DESC;
```

Now, we want to keep the rows where the `b` table has play count greater than `a` table, or if they are equal, then if the `b` table has artist less than `a` table.

To do that we can continue the `JOIN` condition with `AND` and add `b.play_count > a.play_count` and `b.artist < a.artist` in case of a tie.
The idea here is subtle:
- For a given artist `a`, we count how many artists `b` (for the same user) have more plays, or the same plays but come earlier alphabetically.
- If fewer than 3 artists beat `a`, then `a` must be in the top 3.

So the query becomes this:

```sql
SELECT 
    a.user_name,
    a.artist,
    a.play_count
FROM (
    SELECT user_name, artist, COUNT(*) AS play_count
    FROM listening_logs
    GROUP BY user_name, artist
) a
LEFT JOIN (
    SELECT user_name, artist, COUNT(*) AS play_count
    FROM listening_logs
    GROUP BY user_name, artist
) b ON a.user_name = b.user_name
   AND (
       b.play_count > a.play_count
       OR (b.play_count = a.play_count AND b.artist < a.artist)
   )
GROUP BY a.user_name, a.artist, a.play_count
HAVING COUNT(b.artist) < 3
ORDER BY a.user_name, a.play_count DESC, a.artist;
```
```
sqlite> SELECT
    a.user_name,
    a.artist,
    a.play_count
FROM (
    SELECT user_name, artist, COUNT(*) AS play_count
    FROM listening_logs
    GROUP BY user_name, artist
) a
LEFT JOIN (
    SELECT user_name, artist, COUNT(*) AS play_count
    FROM listening_logs
    GROUP BY user_name, artist
) b ON a.user_name = b.user_name
   AND (
       b.play_count > a.play_count
       OR (b.play_count = a.play_count AND b.artist < a.artist)
   )
GROUP BY a.user_name, a.artist, a.play_count
HAVING COUNT(b.artist) < 3
ORDER BY a.user_name, a.play_count DESC, a.artist;
+-------------------+--------------------------------------------+------------+
|     user_name     |                   artist                   | play_count |
+-------------------+--------------------------------------------+------------+
| Abigail Hernandez | Ed Sheeran                                 | 78         |
| Abigail Hernandez | Rotten Mango                               | 15         |
| Abigail Hernandez | Billie Eilish                              | 4          |
| Adrian Cox        | Kendrick Lamar                             | 128        |
| Adrian Cox        | Stuff You Should Know                      | 30         |
| Adrian Cox        | Fuerza Regida                              | 6          |
| Alex Rivera       | Ed Sheeran                                 | 274        |
| Alex Rivera       | Call Her Daddy (Alex Cooper)               | 42         |
| Alex Rivera       | Green Day                                  | 11         |
| Anders Nilsson    | Snow Patrol                                | 101        |
| Anders Nilsson    | SmartLess                                  | 29         |
| Anders Nilsson    | Blink-182                                  | 5          |
| Anthony King      | Pentatonix                                 | 114        |
| Anthony King      | The Tucker Carlson Show                    | 14         |
| Anthony King      | Angels & Airwaves                          | 5          |
...
...
| Zara Sheikh       | Green Day                                  | 138        |
| Zara Sheikh       | This Past Weekend w Theo Von               | 20         |
| Zara Sheikh       | The Beatles                                | 7          |
| Zoe Garcia        | Arijit Singh                               | 50         |
| Zoe Garcia        | Huberman Lab                               | 14         |
| Zoe Garcia        | Kendrick Lamar                             | 5          |
| Zoe Wilson        | Pentatonix                                 | 65         |
| Zoe Wilson        | The Mel Robbins Podcast                    | 14         |
| Zoe Wilson        | Angels & Airwaves                          | 4          |
| Zuri Okafor       | Kendrick Lamar                             | 96         |
| Zuri Okafor       | The Tim Dillon Show                        | 14         |
| Zuri Okafor       | Ed Sheeran                                 | 5          |
+-------------------+--------------------------------------------+------------+
Run Time: real 0.088 user 0.082510 sys 0.003999
```
This is the final query, it looks long, it might be not the best way to do it, but its definitely not the worst way to do it.

### With Window Functions

Ok! I don't know window functions, but I searched and found that we could partition things before we group them or order them in the final result set. That's what we want right?

We had grouped the logs for each username and artist and counted the number of plays. Now we want to rank the artists for each user in the decreasing order of number of plays. 

So, we start with the same thing:

```sql
SELECT
    user_name,
    artist,
    COUNT(*) AS play_count
FROM listening_logs
GROUP BY user_name, artist
ORDER BY user_name, play_count DESC, artist;
```

You can see that we have `user_name` column, what if we could separate out the users, and then we could rank the artists for each user separately.

For that we can use `ROW_NUMBER` window function. This function needs a `PARTITION BY` clause which lets us create separate partitions based on certain columns, and then we can rank them using the `ORDER BY` clause as we do with ordinary statement. That becomes the row_number, which we can use as a rank to rank each artist for a user based on the number of time the user has played him/her.

```sql
SELECT 
    user_name,
    artist,
    COUNT(*) AS play_count,
    ROW_NUMBER() OVER (
        PARTITION BY user_name 
        ORDER BY COUNT(*) DESC, artist
    ) AS ranks
FROM listening_logs
GROUP BY user_name, artist
```

Here, `ROW_NUMBER` is a window function that assigns a rank to each row in the result set. It partitions the result set by `user_name` and orders the rows in the decreasing order of `COUNT(*)` and `artist` name.

So, imagine this table:

```
+-------------------+-------------------------------------+------------+-------+
|     user_name     |               artist                | play_count | ranks |
+-------------------+-------------------------------------+------------+-------+
| Abigail Hernandez | Ed Sheeran                          | 78         | 1     |
| Abigail Hernandez | Rotten Mango                        | 15         | 2     |
| Abigail Hernandez | Billie Eilish                       | 4          | 3     |
| Abigail Hernandez | Hans Zimmer                         | 4          | 4     |
| Abigail Hernandez | John Legend                         | 3          | 5     |
| Abigail Hernandez | John Williams                       | 3          | 6     |
| Abigail Hernandez | The Beatles                         | 3          | 7     |
| Abigail Hernandez | The Rolling Stones                  | 3          | 8     |
| Abigail Hernandez | Angels & Airwaves                   | 2          | 9     |
| Abigail Hernandez | Bad Bunny                           | 2          | 10    |
| Abigail Hernandez | Beyonce                             | 2          | 11    |
| Abigail Hernandez | Coldplay                            | 2          | 12    |
| Abigail Hernandez | Foo Fighters                        | 2          | 13    |
| Abigail Hernandez | Fuerza Regida                       | 2          | 14    |
| Abigail Hernandez | Kendrick Lamar                      | 2          | 15    |
| Abigail Hernandez | Ludovico Einaudi                    | 2          | 16    |
| Abigail Hernandez | Mariah Carey                        | 2          | 17    |
| Abigail Hernandez | Pentatonix                          | 2          | 18    |
| Abigail Hernandez | SmartLess                           | 2          | 19    |
| Abigail Hernandez | The Weeknd                          | 2          | 20    |
| Abigail Hernandez | Adele                               | 1          | 21    |
| Abigail Hernandez | Ariana Grande                       | 1          | 22    |
| Abigail Hernandez | Armchair Expert With Dax Shepard    | 1          | 23    |
| Abigail Hernandez | Bruno Mars                          | 1          | 24    |
| Abigail Hernandez | Candace                             | 1          | 25    |
| Abigail Hernandez | Crime, Conspiracy, Cults and Murder | 1          | 26    |
| Abigail Hernandez | Green Day                           | 1          | 27    |
| Abigail Hernandez | Matt and Shanes Secret Podcast      | 1          | 28    |
| Abigail Hernandez | On Purpose With Jay Shetty          | 1          | 29    |
| Abigail Hernandez | Snow Patrol                         | 1          | 30    |
| Abigail Hernandez | Sufjan Stevens                      | 1          | 31    |
| Abigail Hernandez | Taylor Swift                        | 1          | 32    |
| Abigail Hernandez | The Mel Robbins Podcast             | 1          | 33    |
| Abigail Hernandez | Unseen                              | 1          | 34    |
| Adrian Cox        | Kendrick Lamar                      | 128        | 1     |
| Adrian Cox        | Stuff You Should Know               | 30         | 2     |
| Adrian Cox        | Fuerza Regida                       | 6          | 3     |
| Adrian Cox        | Pentatonix                          | 6          | 4     |
| Adrian Cox        | Taylor Swift                        | 6          | 5     |
| Adrian Cox        | Snow Patrol                         | 5          | 6     |
+-------------------+-------------------------------------+------------+-------+
```

Now, we have ranked the artist for each user, there are `34` artist played by `Abigail Hernandez` so there are `34` ranks. Now, we need to filter out the top 3 artist for each user. That would be simple right?

Just add the `WHERE` clause:

```sql
SELECT
    user_name,
    artist,
    COUNT(*) AS play_count,
    ROW_NUMBER() OVER (
        PARTITION BY user_name 
        ORDER BY COUNT(*) DESC, artist
    ) AS ranks
FROM listening_logs
WHERE ranks <= 3
GROUP BY user_name, artist
ORDER BY user_name, play_count DESC, artist;
```

Well not really!

```
sqlite> SELECT
    user_name,
    artist,
    COUNT(*) AS play_count,
    ROW_NUMBER() OVER (
        PARTITION BY user_name 
        ORDER BY COUNT(*) DESC, artist
    )
FROM listening_logs
WHERE ranks <= 3
GROUP BY user_name, artist
ORDER BY user_name, play_count DESC, artist;
Run Time: real 0.000 user 0.000110 sys 0.000000
Parse error: misuse of aliased window function ranks
```

We can't filter the window function column with a `WHERE` clause. At the time the `WHERE` clause is evaluated, the `SELECT` list (including ranks) has not been computed yet. So ranks doesn't exist yet!

So, now what, so close yet so far!

Let's wrap the `SELECT` in a subquery:

```sql
SELECT user_name, artist, play_count
FROM (
    SELECT 
        user_name,
        artist,
        COUNT(*) AS play_count,
        ROW_NUMBER() OVER (
            PARTITION BY user_name
            ORDER BY COUNT(*) DESC, artist
        ) AS ranks
    FROM listening_logs
    GROUP BY user_name, artist
) ranked
WHERE ranks <= 3
ORDER BY user_name, play_count DESC, artist;
```

And there we have it!

We wrapped the `SELECT` in a subquery, and now we can filter the window function column with a `WHERE` clause.

We can even do it with a `CTE` i.e. common table expression. That's just a subquery with a name.

Not sure when it could be handy, but here it is.

```sql
WITH ranked AS (
    SELECT
        user_name,
        artist,
        COUNT(*) AS play_count,
        ROW_NUMBER() OVER (
            PARTITION BY user_name
            ORDER BY COUNT(*) DESC, artist
        ) AS ranks
    FROM listening_logs
    GROUP BY user_name, artist
)
SELECT user_name, artist, play_count
FROM ranked
WHERE ranks <= 3
ORDER BY user_name, play_count DESC, artist;
```

So, that is day 5, ok that is getting a little tricky now!

Off to day 6.

