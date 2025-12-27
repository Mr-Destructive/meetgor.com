---
type: "sqlog"
title: "Advent of SQL 2025 Day 13: XML Travel Manifests"
slug: "advent-of-sql-2025-day-13"
date: 2025-12-27
series: ["advent-of-sql-2025"]
tags: ["sqlite", "sql", "advent-of-sql"]
---

## Advent of SQL - Day 13, XML Travel Manifests

Its day 13 of Advent of SQL, we have some xml to parse, which I don't think SQL can handle, but string manipulation to the rescue.

Let's get the SQL for the day:

```sql
DROP TABLE IF EXISTS travel_manifests;

CREATE TABLE travel_manifests (
    manifest_id INT PRIMARY KEY,
    vehicle_id TEXT,
    departure_time TIMESTAMP,
    manifest_xml XML
);

INSERT INTO travel_manifests (manifest_id, vehicle_id, departure_time, manifest_xml) VALUES
  (1, 'SLEIGH-01', '2025-12-22 06:00:00', '<manifest><passengers><passenger><name>Nia Grant</name><ticket_class>overnight</ticket_class></passenger><passenger><name>Sofia Kim</name><ticket_class>overnight</ticket_class><engine_check>ignored</engine_check></passenger><passenger><name>Jonah Wolfe</name><ticket_class>standard</ticket_class><snack_inventory>ignored</snack_inventory></passenger></passengers></manifest>'),
  (2, 'SLEIGH-07', '2025-12-23 13:00:00', '<manifest><passengers><passenger><name>Ravi Patel</name><ticket_class>priority</ticket_class></passenger><passenger><name>Keiko Ito</name><ticket_class>standard</ticket_class></passenger><passenger><name>Anya Pavlov</name><ticket_class>standard</ticket_class></passenger><passenger><name>Nia Grant</name><ticket_class>priority</ticket_class><snack_inventory>ignored</snack_inventory></passenger><passenger><name>Carter Lewis</name><ticket_class>priority</ticket_class><engine_check>ignored</engine_check></passenger><passenger><name>Layla Brooks</name><ticket_class>standard</ticket_class><reindeer_mood>ignored</reindeer_mood></passenger></passengers></manifest>'),
  (3, 'FLIGHT-NP-9', '2025-12-22 18:00:00', '<manifest><passengers><passenger><name>Keiko Ito</name><ticket_class>overnight</ticket_class></passenger><passenger><name>Jonah Wolfe</name><ticket_class>priority</ticket_class></passenger><passenger><name>Diego Ramos</name><ticket_class>overnight</ticket_class><snack_inventory>ignored</snack_inventory></passenger><passenger><name>Priya Das</name><ticket_class>priority</ticket_class></passenger><passenger><name>Layla Brooks</name><ticket_class>standard</ticket_class></passenger></passengers></manifest>'),
  (4, 'TRAIN-ICE-3', '2025-12-22 18:00:00', '<manifest><reindeer_mood>low</reindeer_mood><passengers><passenger><name>Isla Torres</name><ticket_class>overnight</ticket_class></passenger><passenger><name>Ravi Patel</name><ticket_class>overnight</ticket_class></passenger><passenger><name>Hiro Tanaka</name><ticket_class>standard</ticket_class></passenger><passenger><name>Priya Das</name><ticket_class>priority</ticket_class></passenger></passengers></manifest>'),
  (5, 'FLIGHT-NP-9', '2025-12-22 17:00:00', '<manifest><passengers><passenger><name>Nia Grant</name><ticket_class>priority</ticket_class></passenger><passenger><name>Mateo Cruz</name><ticket_class>standard</ticket_class><snowfall_inches>ignored</snowfall_inches></passenger></passengers></manifest>'),
  (6, 'CARGO-12', '2025-12-22 15:00:00', '<manifest><passengers><passenger><name>Carter Lewis</name><ticket_class>overnight</ticket_class></passenger><passenger><name>Sofia Kim</name><ticket_class>priority</ticket_class></passenger><passenger><name>Hiro Tanaka</name><ticket_class>priority</ticket_class></passenger><passenger><name>Lucas Ford</name><ticket_class>priority</ticket_class><snack_inventory>ignored</snack_inventory></passenger><passenger><name>Nia Grant</name><ticket_class>overnight</ticket_class></passenger></passengers></manifest>'),
  (7, 'SLEIGH-01', '2025-12-22 11:00:00', '<manifest><snack_inventory>unknown</snack_inventory><passengers><passenger><name>Priya Das</name><ticket_class>standard</ticket_class></passenger><passenger><name>Diego Ramos</name><ticket_class>standard</ticket_class></passenger><passenger><name>Lucas Ford</name><ticket_class>overnight</ticket_class></passenger><passenger><name>Carter Lewis</name><ticket_class>overnight</ticket_class><reindeer_mood>ignored</reindeer_mood></passenger><passenger><name>Hiro Tanaka</name><ticket_class>priority</ticket_class><weather_note>ignored</weather_note></passenger><passenger><name>Zara Sheikh</name><ticket_class>overnight</ticket_class></passenger></passengers></manifest>'),
  (8, 'CARGO-12', '2025-12-23 13:00:00', '<manifest><passengers><passenger><name>Jonah Wolfe</name><ticket_class>standard</ticket_class></passenger><passenger><name>Layla Brooks</name><ticket_class>overnight</ticket_class></passenger><passenger><name>Leo Becker</name><ticket_class>overnight</ticket_class><weather_note>ignored</weather_note></passenger><passenger><name>Ravi Patel</name><ticket_class>standard</ticket_class></passenger><passenger><name>Sofia Kim</name><ticket_class>priority</ticket_class><snack_inventory>ignored</snack_inventory></passenger><passenger><name>Elena Morales</name><ticket_class>overnight</ticket_class></passenger></passengers></manifest>'),
  (9, 'SLEIGH-01', '2025-12-23 10:00:00', '<manifest><passengers><passenger><name>Sofia Kim</name><ticket_class>overnight</ticket_class></passenger><passenger><name>Bianca Pereira</name><ticket_class>standard</ticket_class></passenger><passenger><name>Zara Sheikh</name><ticket_class>overnight</ticket_class></passenger><passenger><name>Elena Morales</name><ticket_class>overnight</ticket_class><reindeer_mood>ignored</reindeer_mood></passenger><passenger><name>Hiro Tanaka</name><ticket_class>priority</ticket_class></passenger><passenger><name>Keiko Ito</name><ticket_class>priority</ticket_class><snack_inventory>ignored</snack_inventory></passenger></passengers></manifest>'),
  (10, 'SLEIGH-01', '2025-12-22 21:00:00', '<manifest><snowfall_inches>low</snowfall_inches><passengers><passenger><name>Nia Grant</name><ticket_class>standard</ticket_class></passenger><passenger><name>Ava Johnson</name><ticket_class>standard</ticket_class></passenger><passenger><name>Priya Das</name><ticket_class>priority</ticket_class></passenger><passenger><name>Mateo Cruz</name><ticket_class>standard</ticket_class></passenger><passenger><name>Bianca Pereira</name><ticket_class>overnight</ticket_class></passenger><passenger><name>Leo Becker</name><ticket_class>overnight</ticket_class></passenger></passengers></manifest>'),
  (11, 'SLEIGH-07', '2025-12-23 10:00:00', '<manifest><passengers><passenger><name>Jonah Wolfe</name><ticket_class>priority</ticket_class></passenger><passenger><name>Bianca Pereira</name><ticket_class>priority</ticket_class></passenger><passenger><name>Anya Pavlov</name><ticket_class>standard</ticket_class></passenger></passengers></manifest>'),
  (12, 'SLEIGH-01', '2025-12-22 08:00:00', '<manifest><reindeer_mood>ok</reindeer_mood><passengers><passenger><name>Ravi Patel</name><ticket_class>priority</ticket_class><snack_inventory>ignored</snack_inventory></passenger><passenger><name>Bianca Pereira</name><ticket_class>priority</ticket_class></passenger><passenger><name>Keiko Ito</name><ticket_class>standard</ticket_class></passenger></passengers></manifest>'),
  (13, 'FLIGHT-NP-9', '2025-12-22 11:00:00', '<manifest><snowfall_inches>ok</snowfall_inches><passengers><passenger><name>Nia Grant</name><ticket_class>priority</ticket_class></passenger><passenger><name>Elena Morales</name><ticket_class>standard</ticket_class></passenger></passengers></manifest>'),
  (14, 'SLEIGH-01', '2025-12-22 14:00:00', '<manifest><engine_check>high</engine_check><passengers><passenger><name>Layla Brooks</name><ticket_class>standard</ticket_class><weather_note>ignored</weather_note></passenger><passenger><name>Anya Pavlov</name><ticket_class>overnight</ticket_class></passenger><passenger><name>Keiko Ito</name><ticket_class>standard</ticket_class></passenger><passenger><name>Zara Sheikh</name><ticket_class>standard</ticket_class><engine_check>ignored</engine_check></passenger><passenger><name>Sofia Kim</name><ticket_class>overnight</ticket_class></passenger></passengers></manifest>'),
  (15, 'FLIGHT-NP-9', '2025-12-22 14:00:00', '<manifest><snowfall_inches>ok</snowfall_inches><passengers><passenger><name>Priya Das</name><ticket_class>priority</ticket_class><snowfall_inches>ignored</snowfall_inches></passenger><passenger><name>Bianca Pereira</name><ticket_class>standard</ticket_class></passenger><passenger><name>Nia Grant</name><ticket_class>standard</ticket_class></passenger></passengers></manifest>');
```

Just one table, with some wild XML.

Let's see what do we want to do in the problem statement.

## Problem

> Using the `travel_manifests` table, extract the passenger information from the XML data and produce a report that shows all of the departure times for "CARGO" vehicles that have more than 20 passengers booked. Include in the results:
> 
> - The vehicle_id
> - The departure_time
> - The total number of passengers on that departure
> - Order the results by departure_time.


Ok, so we need the number of passengers in the records which are of type `CARGO` and have more than 20 passengers booked.

Interesting!

Let's look at one record.

```
sqlite> .schema
CREATE TABLE travel_manifests (
    manifest_id INT PRIMARY KEY,
    vehicle_id TEXT,
    departure_time TIMESTAMP,
    manifest_xml XML
);
sqlite> SELECT * FROM travel_manifests WHERE id = 1;
Parse error: no such column: id
  SELECT * FROM travel_manifests WHERE id = 1;
                         error here ---^
sqlite> SELECT * FROM travel_manifests LIMIT 1;
+-------------+------------+---------------------+--------------------------------------------------------------+
| manifest_id | vehicle_id |   departure_time    |                         manifest_xml                         |
+-------------+------------+---------------------+--------------------------------------------------------------+
| 1           | SLEIGH-01  | 2025-12-22 06:00:00 | <manifest><passengers><passenger><name>Nia Grant</name><tick |
|             |            |                     | et_class>overnight</ticket_class></passenger><passenger><nam |
|             |            |                     | e>Sofia Kim</name><ticket_class>overnight</ticket_class><eng |
|             |            |                     | ine_check>ignored</engine_check></passenger><passenger><name |
|             |            |                     | >Jonah Wolfe</name><ticket_class>standard</ticket_class><sna |
|             |            |                     | ck_inventory>ignored</snack_inventory></passenger></passenge |
|             |            |                     | rs></manifest>                                               |
+-------------+------------+---------------------+--------------------------------------------------------------+
sqlite> 

```

So, we have the following columns 

- `vehicle_id` which is I think related to filtering `CARGO` related vehicles only
- `departure_time` which we just return as is
- `manifest_xml`, oh! This is xml and it has passenger details. 

If we look carefully, we can see the xml looks like this:

```xml
<manifest>
    <passengers>
        <passenger>
            <name>Nia Grant</name> 
           <ticket_class>overnight</ticket_class>
        </passenger>
        <passenger>
            <name>Sofia Kim</name> 
           <ticket_class>overnight</ticket_class>
           <engine_check>ignored</engine_check>
       </passenger>
       <passenger>
           <name>Jonah Wolfe</name>
           <ticket_class>standard</ticket_class>
           <snack_inventory>ignored</snack_inventory>
       </passenger>
    </passengers>
</manifest>
```

We have `manifest` which has a property of `passengers` which is a list of `passenger` tags, inside of which, each element of `passenger` has its details like `name`, `ticket_class`, etc.

We only want the count of `passengers`, how can we get that? The dirtiest way to do is to count the occurances of `<passenger>` or `</passenger>` in the xml string. 

We can do that with counting the full length of the xml string, and then dividing by the number of times the string can be replaced(which is the dirty part, there could be hidden `<passenger>` string somewhere that might break this logic, but if it is a valid xml, it works). We count the number of characters left after we replace the string `<passenger>` with empty string `''` so that we can get the difference of the total number of character and the number of characters occupied by the string `<passenger>`. This difference if we divide by the length of `<passenger>` will give us the count of the number of times the `<passenger>` string is present in the xml string.

Let's take an example from the above. The length of the xml string is 374.

```sql
SELECT LENGTH(manifest_xml) FROM travel_manifests LIMIT 1;
```

```
sqlite> SELECT LENGTH(manifest_xml) FROM travel_manifests LIMIT 1;
+----------------------+
| LENGTH(manifest_xml) |
+----------------------+
| 374                  |
+----------------------+
sqlite> 
```

Let's replace the occurances of `<passenger> with empty string in the `manifest_xml` string, like so:

```sql
SELECT REPLACE(manifest_xml, '<passenger>', '') FROM travel_manifests LIMIT 1;
```

Now, we can see the string `<passenger>` is gone from the returned result set. We can try getting its length now.

```sql
SELECT 
    LENGTH(
        REPLACE(manifest_xml, '<passenger>', '')
    )
FROM travel_manifests LIMIT 1;
```

Now, it says `341` why? Because we removed (replaced with empty string) the occurences of `<passenger>`.

Let's get the length of `'<passenger>'` string, which should be `11` right? 

Spell it `p-a-s-s-e-n-g-e-r` as `pass` + `enger` (4+5=9) and 2 for `<>` so 11. Sometimes I don't know how to do math, I use SQL.

```sql
SELECT LENGTH('<passenger>');
```

There it is `11`.

Now, if you compute the difference of the actual length of XML with the removed parts of the `<passenger>` what do we get?

```sql
SELECT LENGTH(manifest_xml) - LENGTH(REPLACE(manifest_xml, '<passenger>', '')) FROM travel_manifests LIMIT 1;
```

WE got `33`, why would you ask because `11` times 3 is `33`. We found three instances of `<passenger>`,  so we just need to divide by the length of `<passenger>` or hard code it as `11` doesn't matter.

We would get the number of occurances of `<passenger>` which will give the number of passenger in the xml string.

```sql
SELECT
    (
        LENGTH(manifest_xml) - LENGTH(REPLACE(manifest_xml, '<passenger>', ''))
    ) / LENGTH('<passenger>')
FROM travel_manifests LIMIT 1;
```

Phew, its `3`!

That was a lot for a simple stuff. But hey its fun!

```
sqlite> SELECT (manifest_xml) FROM travel_manifests LIMIT 1;
+--------------------------------------------------------------+
|                         manifest_xml                         |
+--------------------------------------------------------------+
| <manifest><passengers><passenger><name>Nia Grant</name><tick |
| et_class>overnight</ticket_class></passenger><passenger><nam |
| e>Sofia Kim</name><ticket_class>overnight</ticket_class><eng |
| ine_check>ignored</engine_check></passenger><passenger><name |
| >Jonah Wolfe</name><ticket_class>standard</ticket_class><sna |
| ck_inventory>ignored</snack_inventory></passenger></passenge |
| rs></manifest>                                               |
+--------------------------------------------------------------+

sqlite> SELECT REPLACE(manifest_xml, '<passenger>', '') FROM travel_manifests LIMIT 1;
+--------------------------------------------------------------+
|           REPLACE(manifest_xml, '<passenger>', '')           |
+--------------------------------------------------------------+
| <manifest><passengers><name>Nia Grant</name><ticket_class>ov |
| ernight</ticket_class></passenger><name>Sofia Kim</name><tic |
| ket_class>overnight</ticket_class><engine_check>ignored</eng |
| ine_check></passenger><name>Jonah Wolfe</name><ticket_class> |
| standard</ticket_class><snack_inventory>ignored</snack_inven |
| tory></passenger></passengers></manifest>                    |
+--------------------------------------------------------------+
sqlite> 
sqlite> SELECT LENGTH(REPLACE(manifest_xml, '<passenger>', '')) FROM travel_manifests LIMIT 1;
+--------------------------------------------------+
| LENGTH(REPLACE(manifest_xml, '<passenger>', '')) |
+--------------------------------------------------+
| 341                                              |
+--------------------------------------------------+
sqlite> SELECT LENGTH(manifest_xml) - LENGTH(REPLACE(manifest_xml, '<passenger>', '')) FROM travel_manifests LIMIT 1;
+--------------------------------------------------------------+
| LENGTH(manifest_xml) - LENGTH(REPLACE(manifest_xml, '<passen |
+--------------------------------------------------------------+
| 33                                                           |
+--------------------------------------------------------------+
sqlite> SELECT LENGTH('<passenger>') FROM travel_manifests LIMIT 1;
+-----------------------+
| LENGTH('<passenger>') |
+-----------------------+
| 11                    |
+-----------------------+

sqlite> SELECT (LENGTH(manifest_xml) - LENGTH(REPLACE(manifest_xml, '<passenger>', '')))/LENGTH('<passenger>') FROM travel_manifests LIMIT 1;
+--------------------------------------------------------------+
| (LENGTH(manifest_xml) - LENGTH(REPLACE(manifest_xml, '<passe |
+--------------------------------------------------------------+
| 3                                                            |
+--------------------------------------------------------------+
sqlite> 

```

Now, let construct the query to get the number of passengers.

```sql
SELECT
    vehicle_id,
    departure_time,
    (
        LENGTH(manifest_xml)
        - LENGTH(REPLACE(manifest_xml, '<passenger>', ''))
    ) / LENGTH('<passenger>') AS passengers_in_manifest
FROM travel_manifests
WHERE vehicle_id LIKE 'CARGO-%';
```

We need to wrap it in a CTE to grab and group by the vehicle_id I believe, as there are similar entries.

Also we need to group the records with the same departure time, so that we can combine the number of passengers for that vehicle.


```sql
WITH passenger_counts AS (
    SELECT
        vehicle_id,
        departure_time,
        (
            LENGTH(manifest_xml)
            - LENGTH(REPLACE(manifest_xml, '<passenger>', ''))
        ) / LENGTH('<passenger>') AS passengers_in_manifest
    FROM travel_manifests
    WHERE vehicle_id LIKE 'CARGO-%'
)
SELECT
    vehicle_id,
    departure_time,
    SUM(passengers_in_manifest) AS total_passengers
FROM passenger_counts
GROUP BY vehicle_id, departure_time
HAVING SUM(passengers_in_manifest) > 20
ORDER BY departure_time;
```

We count the number of passengers in the CTE and use it as a filter in the outer query as `SUM(passengers_in_manifest) > 20` which will give the right condition for us to get the result. We have to use `HAVING` as we need to do that after grouping the same `vehicle_id` and records across same `departure_time`.

We also use the `vehicle_id LIKE 'CARGO-%'` in the CTE to filter it right at the inner query to avoid looping in all the queries for computing the number of passengers. 


That solves this problem.


That's it from day 13 of Advent of SQL.

There are other ways, but its the same parsing, We can use JOINs and stuff, but hey that was not the point of this.

Anyways! See you tomorrow for day 14!

