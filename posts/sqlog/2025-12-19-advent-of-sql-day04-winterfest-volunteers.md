---
type: "sqlog"
title: "Advent of SQL 2025 Day 4: WinterFest Volunteers"
slug: "advent-of-sql-2025-day-4"
date: 2025-12-19
series: ["advent-of-sql-2025"]
tags: ["sqlite", "sql", "advent-of-sql"]
---

## Advent of SQL Day 4 WinterFest Volunteers

It is day 4 of advent of SQL.

No fuss, straight to the problem, the elves and humans are getting dumber as the days progress.

Let's download the SQL inserts for the day.

And load it into a SQLite shell.

```sql
DROP TABLE IF EXISTS official_shifts;
DROP TABLE IF EXISTS last_minute_signups;

CREATE TABLE official_shifts (
    id INT PRIMARY KEY,
    volunteer_name TEXT,
    role TEXT,
    shift_time TEXT,
    age_group TEXT,
    code TEXT
);

CREATE TABLE last_minute_signups (
    id INT PRIMARY KEY,
    volunteer_name TEXT,
    assigned_task TEXT,
    time_slot TEXT
);

INSERT INTO official_shifts (id, volunteer_name, role, shift_time, age_group, code) VALUES
    (1, 'Jude Thompson', 'choir_assistant', '12:00 PM', 'senior', NULL),
    (2, 'Mateo Cruz', 'choir_assistant', '12:00 PM', 'senior', NULL),
    (3, 'Olivia Dubois', 'choir_assistant', '2:00 PM', 'teen', 'A1'),
    (4, 'Jeff Bezos', 'choir_assistant', '10:00 AM', 'adult', 'X7'),
    (5, 'Kian Rahimi', 'stage_setup', '12:00 PM', 'adult', 'X7'),
    (6, 'Haruto Sato', 'cocoa_station', '10:00 AM', 'adult', 'X7'),
    (7, 'Uma Singh', 'parking_support', '10:00 AM', 'adult', NULL),
    (8, 'Owen Scott', 'parking_support', '10:00 AM', 'adult', 'X7'),
    (9, 'Adil Rahman', 'stage_setup', '2:00 PM', 'adult', 'A1'),
    (10, 'Aaron Diaz', 'choir_assistant', '2:00 PM', 'senior', 'X7'),
    (11, 'Carter Lewis', 'cocoa_station', '10:00 AM', 'senior', 'B2'),
    (12, 'Anya Pavlov', 'stage_setup', '10:00 AM', 'senior', 'OLD'),
    (13, 'Ethan Brown', 'stage_setup', '2:00 PM', 'adult', 'A1'),
    (14, 'Lucia Fernandez', 'choir_assistant', '12:00 PM', 'senior', 'X7'),
    (15, 'Casey Morgan', 'choir_assistant', '12:00 PM', 'teen', 'OLD');

INSERT INTO last_minute_signups (id, volunteer_name, assigned_task, time_slot) VALUES
    (1, 'Jude Thompson', 'Choir', 'noon'),
    (2, 'Mateo Cruz', 'choir', 'noon'),
    (3, 'Olivia Dubois', 'choir', '2 PM'),
    (4, 'Jeff Bezos', 'choir assistant', '10AM'),
    (5, 'Kian Rahimi', 'stage setup', 'noon'),
    (6, 'Haruto Sato', 'cocoa station', '10AM'),
    (7, 'Uma Singh', 'parking_support', '10AM'),
    (8, 'Owen Scott', 'parking', '10AM'),
    (9, 'Adil Rahman', 'Stage-Setup', '2 PM'),
    (10, 'Aaron Diaz', 'Choir', '2 PM'),
    (11, 'Carter Lewis', 'Cocoa Station', '10AM'),
    (12, 'Anya Pavlov', 'stage_setup', '10AM'),
    (13, 'Olivia Brown', 'stage setup', '2 PM'),
    (14, 'Lena Fischer', 'cocoa station', '2 pm'),
    (15, 'Nolan Murphy', 'parking-support', '10AM');
```

Once the data is loaded, let's sneak peak.

```sql
SELECT * FROM official_shifts LIMIT 15;
```

```sql
SELECT * FROM last_minute_signups LIMIT 15;
```

Let's count how many rows in each table we have:

```sql
SELECT COUNT(*) FROM official_shifts LIMIT 15;
SELECT COUNT(*) FROM last_minute_signups LIMIT 15;
```

Alright, the data is visible and we can head on to the problem statement.



```
$ sqlite3
SQLite version 3.45.1 2024-01-30 16:01:20
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
sqlite> .read day4_inserts.sql
sqlite> .schema
CREATE TABLE official_shifts (
    id INT PRIMARY KEY,
    volunteer_name TEXT,
    role TEXT,
    shift_time TEXT,
    age_group TEXT,
    code TEXT
);
CREATE TABLE last_minute_signups (
    id INT PRIMARY KEY,
    volunteer_name TEXT,
    assigned_task TEXT,
    time_slot TEXT
);

sqlite> SELECT * FROM official_shifts LIMIT 15;
1|Jude Thompson|choir_assistant|12:00 PM|senior|
2|Mateo Cruz|choir_assistant|12:00 PM|senior|
3|Olivia Dubois|choir_assistant|2:00 PM|teen|A1
4|Jeff Bezos|choir_assistant|10:00 AM|adult|X7
5|Kian Rahimi|stage_setup|12:00 PM|adult|X7
6|Haruto Sato|cocoa_station|10:00 AM|adult|X7
7|Uma Singh|parking_support|10:00 AM|adult|
8|Owen Scott|parking_support|10:00 AM|adult|X7
9|Adil Rahman|stage_setup|2:00 PM|adult|A1
10|Aaron Diaz|choir_assistant|2:00 PM|senior|X7
11|Carter Lewis|cocoa_station|10:00 AM|senior|B2
12|Anya Pavlov|stage_setup|10:00 AM|senior|OLD
13|Ethan Brown|stage_setup|2:00 PM|adult|A1
14|Lucia Fernandez|choir_assistant|12:00 PM|senior|X7
15|Casey Morgan|choir_assistant|12:00 PM|teen|OLD

sqlite> .mode table 
sqlite> SELECT * FROM official_shifts LIMIT 15;
+----+-----------------+-----------------+------------+-----------+------+
| id | volunteer_name  |      role       | shift_time | age_group | code |
+----+-----------------+-----------------+------------+-----------+------+
| 1  | Jude Thompson   | choir_assistant | 12:00 PM   | senior    |      |
| 2  | Mateo Cruz      | choir_assistant | 12:00 PM   | senior    |      |
| 3  | Olivia Dubois   | choir_assistant | 2:00 PM    | teen      | A1   |
| 4  | Jeff Bezos      | choir_assistant | 10:00 AM   | adult     | X7   |
| 5  | Kian Rahimi     | stage_setup     | 12:00 PM   | adult     | X7   |
| 6  | Haruto Sato     | cocoa_station   | 10:00 AM   | adult     | X7   |
| 7  | Uma Singh       | parking_support | 10:00 AM   | adult     |      |
| 8  | Owen Scott      | parking_support | 10:00 AM   | adult     | X7   |
| 9  | Adil Rahman     | stage_setup     | 2:00 PM    | adult     | A1   |
| 10 | Aaron Diaz      | choir_assistant | 2:00 PM    | senior    | X7   |
| 11 | Carter Lewis    | cocoa_station   | 10:00 AM   | senior    | B2   |
| 12 | Anya Pavlov     | stage_setup     | 10:00 AM   | senior    | OLD  |
| 13 | Ethan Brown     | stage_setup     | 2:00 PM    | adult     | A1   |
| 14 | Lucia Fernandez | choir_assistant | 12:00 PM   | senior    | X7   |
| 15 | Casey Morgan    | choir_assistant | 12:00 PM   | teen      | OLD  |
+----+-----------------+-----------------+------------+-----------+------+

sqlite> SELECT * FROM last_minute_signups LIMIT 15;
+----+----------------+-----------------+-----------+
| id | volunteer_name |  assigned_task  | time_slot |
+----+----------------+-----------------+-----------+
| 1  | Jude Thompson  | Choir           | noon      |
| 2  | Mateo Cruz     | choir           | noon      |
| 3  | Olivia Dubois  | choir           | 2 PM      |
| 4  | Jeff Bezos     | choir assistant | 10AM      |
| 5  | Kian Rahimi    | stage setup     | noon      |
| 6  | Haruto Sato    | cocoa station   | 10AM      |
| 7  | Uma Singh      | parking_support | 10AM      |
| 8  | Owen Scott     | parking         | 10AM      |
| 9  | Adil Rahman    | Stage-Setup     | 2 PM      |
| 10 | Aaron Diaz     | Choir           | 2 PM      |
| 11 | Carter Lewis   | Cocoa Station   | 10AM      |
| 12 | Anya Pavlov    | stage_setup     | 10AM      |
| 13 | Olivia Brown   | stage setup     | 2 PM      |
| 14 | Lena Fischer   | cocoa station   | 2 pm      |
| 15 | Nolan Murphy   | parking-support | 10AM      |
+----+----------------+-----------------+-----------+

sqlite> SELECT count(*) FROM official_shifts;
+----------+
| count(*) |
+----------+
| 250      |
+----------+

sqlite> SELECT count(*) FROM last_minute_signups;
+----------+
| count(*) |
+----------+
| 126      |
+----------+
sqlite> 
```

## Problem

Here's the challenge for day 4

> Using the official_shifts and last_minute_signups tables, create a combined de-duplicated volunteer list.
> 
> Ensure the list has standardized role labels of Stage Setup, Cocoa Station, Parking Support, Choir Assistant, Snow Shoveling, Handwarmer Handout.
> 
> Make sure that the timeslot formats follow John's official shifts format.

What we have here is a official shift table which could have been entered by the system. However the `last_minute_shift` is messy and has been added from a sheet, so we need to clean it up and combine those two tables data into a single de-duplicated list of volunteers.

Let's see what we got

```sql
SELECT * FROM last_minute_signups;
```

Ok, we have 126 records and the columns are:
1.  `volunteer_name` which doesn't look bad, 
2. `assigned_task` which looks wonky
3. `time_slot` is just wild, we have wired definition of times there.

And let's look at the `official_shifts`

```sql
SELECT * FROM offical_shifts;
```

This looks neat and tidy, nothing looking off from each other.

So we need to make sure we are cleaning up the `last_minute_signups` before we merge them.

```sql
SELECT DISTINCT assigned_task FROM last_minute_signups;
```
Ok, so casing is one thing we can see, `-` and ` ` space are the things to normalize, and then some inconsistent naming convention like `choir` and `choir assistant`, then `parking_support` and `parking`. We need to clean'em up.

We can search for 
```sql
SELECT 
    id,
    volunteer_name,
    time_slot,
    CASE 
        WHEN assigned_task LIKE 'choir%' THEN 'choir_assistant'
        WHEN assigned_task LIKE 'stage%' THEN 'stage_setup' 
        WHEN assigned_task LIKE '%cocoa%' THEN 'cocoa_station'
        WHEN assigned_task LIKE 'parking%' THEN 'parking_support'
        WHEN assigned_task LIKE 'hand%' THEN 'hand_warmer'
        WHEN assigned_task LIKE '%shovel%' THEN 'snow_showel'  
   END AS assigned_task
FROM last_minute_signups;
``` 

We just do a case match for:
- `LIKE 'choir%'` which will match any case (`Choir`, `choir`) and also anything after `choir....` like `choir assistant`
- `LIKE 'stage%'` which will match any case (`Stage`, `stage`) and also anything after `stage...` like `Stage-Setup`, `stage    setup` or `stage_setup`.
- `LIKE '%cocoa'` which will match any case (`Cocoa`, `cocoa`) and also anything before or after `...cocoa...` like `Cocoa Station`, `cocoa station`, etc.
- `LIKE 'parking%'` which will match any case (`Parking`, `parking`) and also anything after `parking...` like `parking-support` or `parking_support`, etc.
- `LIKE 'hand%'` which will match any case (`Hand`, `hand`) and also anything after `hand...` like `handwarmer handout`, `handwarmers`, `Handwarmer-Handout`, etc.
- `LIKE '%shovel%'` which will match any case (`Shovel`, `shovel`) and also anything before and after `...shovel...`  like `Snow-Shoveling`, `shovel`, `snow shoveling`, etc.

Ok, now this looks unified for the `assigned_task`.

```
sqlite> SELECT 
    id,
    volunteer_name,
    time_slot,
    CASE 
        WHEN assigned_task LIKE 'choir%' THEN 'choir_assistant'
        WHEN assigned_task LIKE 'stage%' THEN 'stage_setup'
        WHEN assigned_task LIKE '%cocoa%' THEN 'cocoa_station'
        WHEN assigned_task LIKE 'parking%' THEN 'parking_support'
        WHEN assigned_task LIKE 'hand%' THEN 'hand_warmer'
        WHEN assigned_task LIKE '%shovel%' THEN 'snow_showel' 
    END AS assigned_task 
FROM last_minute_signups;

+-----+-------------------+-----------+-----------------+
| id  |  volunteer_name   | time_slot |  assigned_task  |
+-----+-------------------+-----------+-----------------+
| 1   | Jude Thompson     | noon      | choir_assistant |
| 2   | Mateo Cruz        | noon      | choir_assistant |
| 3   | Olivia Dubois     | 2 PM      | choir_assistant |
| 4   | Jeff Bezos        | 10AM      | choir_assistant |
| 5   | Kian Rahimi       | noon      | stage_setup     |
| 6   | Haruto Sato       | 10AM      | cocoa_station   |
| 7   | Uma Singh         | 10AM      | parking_support |
| 8   | Owen Scott        | 10AM      | parking_support |
| 9   | Adil Rahman       | 2 PM      | stage_setup     |
| 10  | Aaron Diaz        | 2 PM      | choir_assistant |
| 11  | Carter Lewis      | 10AM      | cocoa_station   |
| 12  | Anya Pavlov       | 10AM      | stage_setup     |
| 13  | Olivia Brown      | 2 PM      | stage_setup     |
| 14  | Lena Fischer      | 2 pm      | cocoa_station   |
| 15  | Nolan Murphy      | 10AM      | parking_support |
+-----+-------------------+-----------+-----------------+

```



We need to make it for the time slot too.

```sql
SELECT DISTINCT time_slot FROM last_minute_signups
```

```
sqlite> SELECT DISTINCT time_slot FROM last_minute_signups;
+-----------+
| time_slot |
+-----------+
| noon      |
| 2 PM      |
| 10AM      |
| 2 pm      |
| 10 am     |
+-----------+
sqlite> 
```

Since we have to follow the `official_shifts` let's check over there.
```sql
SELECT distinct shift_time FROM official_shifts;
```

```
sqlite> SELECT distinct shift_time FROM official_shifts;
+------------+
| shift_time |
+------------+
| 12:00 PM   |
| 2:00 PM    |
| 10:00 AM   |
+------------+
```

Ok we have only 3 times to change.

Alright, we have some small things to standardize.

```sql
SELECT 
    DISTINCT CASE 
        WHEN time_slot LIKE '2%' THEN '2:00 PM'
        WHEN time_slot LIKE 'noon' THEN '12:00 PM'
        WHEN time_slot LIKE '10%' THEN '10:00 AM'
    END AS time_slot 
FROM last_minute_signups;
```

```
sqlite> SELECT 
    DISTINCT CASE 
        WHEN time_slot LIKE '2%' THEN '2:00 PM'
        WHEN time_slot LIKE '10%' THEN '10:00 AM'
        WHEN time_slot LIKE 'noon' THEN '12:00 PM'
    END AS time_slot 
FROM last_minute_signups;

+-----------+
| time_slot |
+-----------+
| 12:00 PM  |
| 2:00 PM   |
| 10:00 AM  |
+-----------+
sqlite> 
```

So, we have simply standardize the time_slots.

- `LIKE '2%'` will match any case but we need `LIKE` to match the `%` rest of the stuff after `2`.
- `LIKE '10%'` will match any case but we need `LIKE` to match the `%` rest of the stuff after `10`. We can't keep it `LIKE '1%'` as it will match `1:00` as well
- `LIKE 'noon'` will match any case of `noon` like `NOON` or `Noon`, etc. And we need to cast it to the `HH:MM AM or PM` format.

So, now we can combine them.

```sql
SELECT id, volunteer_name,
    CASE                                                                                                                                                               
        WHEN assigned_task LIKE 'choir%' THEN 'choir_assistant'
        WHEN assigned_task LIKE 'stage%' THEN 'stage_setup'
        WHEN assigned_task LIKE '%cocoa%' THEN 'cocoa_station'
        WHEN assigned_task LIKE 'parking%' THEN 'parking_support'
        WHEN assigned_task LIKE 'hand%' THEN 'hand_warmer'
        WHEN assigned_task LIKE '%shovel%' THEN 'snow_showel' 
    END AS role, 
    CASE
        WHEN time_slot LIKE '2%' THEN '2:00 PM'
        WHEN time_slot LIKE 'noon' THEN '12:00 PM'
        WHEN time_slot LIKE '10%' THEN '10:00 PM'
    END AS shift_time 
FROM last_minute_signups;
```
Just changed the column names from `assigned_task` to `role` and `time_slot` to `shift_time` as per the name convention in the `official_shifts` table.
Phew! its a long statement.

```
sqlite> SELECT id, volunteer_name, CASE                                                                                                                                                               WHEN assigned_task LIKE 'choir%' THEN 'choir_assistant'
        WHEN assigned_task LIKE 'stage%' THEN 'stage_setup'
        WHEN assigned_task LIKE '%cocoa%' THEN 'cocoa_station'
        WHEN assigned_task LIKE 'parking%' THEN 'parking_support'
        WHEN assigned_task LIKE 'hand%' THEN 'hand_warmer'
        WHEN assigned_task LIKE '%shovel%' THEN 'snow_showel' 
    END AS assigned_task, CASE WHEN time_slot LIKE '2%' THEN '2:00 PM' WHEN time_slot LIKE 'noon' THEN '12:00 PM' WHEN time_slot LIKE '10%' THEN '10:00 AM' END AS time_slot FROM last_minute_signups;

+-----+-------------------+-----------------+-----------+
| id  |  volunteer_name   |  assigned_task  | time_slot |
+-----+-------------------+-----------------+-----------+
| 1   | Jude Thompson     | choir_assistant | 12:00 PM  |
| 2   | Mateo Cruz        | choir_assistant | 12:00 PM  |
| 3   | Olivia Dubois     | choir_assistant | 2:00 PM   |
| 4   | Jeff Bezos        | choir_assistant | 10:00 AM  |
| 5   | Kian Rahimi       | stage_setup     | 12:00 PM  |
| 6   | Haruto Sato       | cocoa_station   | 10:00 AM  |
| 7   | Uma Singh         | parking_support | 10:00 AM  |
| 8   | Owen Scott        | parking_support | 10:00 AM  |
| 9   | Adil Rahman       | stage_setup     | 2:00 PM   |
| 10  | Aaron Diaz        | choir_assistant | 2:00 PM   |
| 11  | Carter Lewis      | cocoa_station   | 10:00 AM  |
| 12  | Anya Pavlov       | stage_setup     | 10:00 AM  |
| 13  | Olivia Brown      | stage_setup     | 2:00 PM   |
| 14  | Lena Fischer      | cocoa_station   | 2:00 PM   |
| 15  | Nolan Murphy      | parking_support | 10:00 AM  |
+-----+-------------------+-----------------+-----------+
```

SO, now we have the table of `last_minute_signups` cleaned up, just with select, we can update them if needed.

We now need to combine the both tables, cleaned up `last_minute_signups` and the `official_shifts`, we can use `UNION` to take out the duplicates from the two selection.
REMEMBER to order the rows correctly in both the tables.
- volunteer_name
- role
- shift_time

I don't think name should be same, but I am keeping it same for clarity.

Why `UNION`
- Because we have data in both the tables.
- We don't have a relation in both of the tables, those are the same tables just that the columns are not cleaned or in proper format.
- We want to grab all of them from one, all from other table, and remove the duplicates, that's a definition of `UNION`

We can't use `UNION ALL` as it would include all the rows from both the tables without removing duplicates.

```sql
SELECT volunteer_name, 
    CASE 
        WHEN assigned_task LIKE 'choir%' THEN 'choir_assistant'
        WHEN assigned_task LIKE 'stage%' THEN 'stage_setup'
        WHEN assigned_task LIKE '%cocoa%' THEN 'cocoa_station'
        WHEN assigned_task LIKE 'parking%' THEN 'parking_support'
        WHEN assigned_task LIKE 'hand%' THEN 'hand_warmer'
        WHEN assigned_task LIKE '%shovel%' THEN 'snow_showel' 
    END AS assigned_task,
    CASE 
        WHEN time_slot LIKE '2%' THEN '2:00 PM'
        WHEN time_slot LIKE 'noon' THEN '12:00 PM'
        WHEN time_slot LIKE '10%' THEN '10:00 AM'
    END AS time_slot
FROM last_minute_signups 
UNION  
SELECT 
    volunteer_name,
    role,
    shift_time
FROM official_shifts 
ORDER BY volunteer_name;

```

Ok, that is a mess, isn't it?

```
sqlite> SELECT volunteer_name, 
    CASE 
        WHEN assigned_task LIKE 'choir%' THEN 'choir_assistant'
        WHEN assigned_task LIKE 'stage%' THEN 'stage_setup'
        WHEN assigned_task LIKE '%cocoa%' THEN 'cocoa_station'
        WHEN assigned_task LIKE 'parking%' THEN 'parking_support'
        WHEN assigned_task LIKE 'hand%' THEN 'hand_warmer'
        WHEN assigned_task LIKE '%shovel%' THEN 'snow_showel' 
    END AS assigned_task,
    CASE 
        WHEN time_slot LIKE '2%' THEN '2:00 PM'
        WHEN time_slot LIKE 'noon' THEN '12:00 PM'
        WHEN time_slot LIKE '10%' THEN '10:00 AM'
    END AS time_slot
FROM last_minute_signups 
UNION 
SELECT 
    volunteer_name,
    role,
    shift_time
FROM official_shifts 
ORDER BY volunteer_name;

+-------------------+-----------------+-----------+
|  volunteer_name   |  assigned_task  | time_slot |
+-------------------+-----------------+-----------+
| Aaron Carter      | parking_support | 2:00 PM   |
| Aaron Diaz        | choir_assistant | 2:00 PM   |
| Aaron Diaz        | choir_assistant | 2:00 PM   |
| Aaron Evans       | cocoa_station   | 2:00 PM   |
| Aaron Francis     | hand_warmer     | 2:00 PM   |
| Abigail Hernandez | stage_setup     | 10:00 AM  |
| Adam King         | stage_setup     | 10:00 AM  |
| Adil Foster       | stage_setup     | 2:00 PM   |
| Adil Rahman       | stage_setup     | 2:00 PM   |
| Adil Rahman       | stage_setup     | 2:00 PM   |
| Adrian Cox        | cocoa_station   | 10:00 AM  |
| Aisha Bennett     | cocoa_station   | 12:00 PM  |
| Aisha Khan        | choir_assistant | 12:00 PM  |
| Aisha Khan        | choir_assistant | 12:00 PM  |
| Aisha Mohammed    | cocoa_station   | 2:00 PM   |
+-------------------+-----------------+-----------+
```

There we have it.

Let's count the number of distinct volunteers in the shifts.

```sql
SELECT COUNT(*) FROM (SELECT volunteer_name, role, shift_time FROM official_shifts UNION  SELECT volunteer_name, CASE 
        WHEN assigned_task LIKE '%choir%' THEN 'choir_assistant'
        WHEN assigned_task LIKE '%stage%' THEN 'stage_setup'
        WHEN assigned_task LIKE '%cocoa%' THEN 'cocoa_station'
        WHEN assigned_task LIKE '%parking%' THEN 'parking_support'
        WHEN assigned_task LIKE '%hand%' THEN 'hand_warmer'
        WHEN assigned_task LIKE '%shovel%' THEN 'snow_showel' 
    END AS role, CASE WHEN time_slot LIKE '2%' THEN '2:00 PM' WHEN time_slot LIKE 'noon' THEN '12:00 PM' WHEN time_slot LIKE '10%' THEN '10:00 AM' END AS shift_time FROM last_minute_signups) AS volunteers;
```

Just counted the full union of the statement using `SELECT COUNT(*) FROM <THE FULL QUERY> AS volunteers)

```
sqlite> SELECT COUNT(*) FROM (SELECT volunteer_name, role, shift_time FROM official_shifts UNION  SELECT volunteer_name, CASE 
        WHEN assigned_task LIKE '%choir%' THEN 'choir_assistant'
        WHEN assigned_task LIKE '%stage%' THEN 'stage_setup'
        WHEN assigned_task LIKE '%cocoa%' THEN 'cocoa_station'
        WHEN assigned_task LIKE '%parking%' THEN 'parking_support'
        WHEN assigned_task LIKE '%hand%' THEN 'hand_warmer'
        WHEN assigned_task LIKE '%shovel%' THEN 'snow_showel' 
    END AS role, CASE WHEN time_slot LIKE '2%' THEN '2:00 PM' WHEN time_slot LIKE 'noon' THEN '12:00 PM' WHEN time_slot LIKE '10%' THEN '10:00 AM' END AS shift_time FROM last_minute_signups) AS volunteers;

+------------+
| volunteers |
+------------+
| 284        |
+------------+
sqlite> 

```
So, we have around `284` rows. Looks good.

Pinebrook can see the volunteer list now. The cleaned one.

Off to day 5!
