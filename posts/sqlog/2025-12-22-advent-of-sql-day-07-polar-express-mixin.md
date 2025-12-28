---
type: "sqlog"
title: "Advent of SQL 2025 Day 7: Polar Express Mixin"
slug: "advent-of-sql-2025-day-7"
date: 2025-12-22 15:30 +0530
series: ["advent-of-sql-2025"]
tags: ["sqlite", "sql", "advent-of-sql"]
---

## Advent of SQL, Day 7 - Polar Express

There were a few things, I had to dig up for converting the JSON `ARRAY[]` in the statements into strings for SQLite, we can't really use list of strings in SQLite.

Here's the command to convert that array of strings into string.

```
sed "s/ARRAY\['/'\[\"/g; s/','/\",\"/g; s/']/\"]'/g" day7-inserts.sql > day7-inserts-sqlite.sql
```
 
OK, once that's done, this can be safely run into a sqlite database.

```
$ sqlite3
SQLite version 3.45.1 2024-01-30 16:01:20
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
sqlite> .read day7-inserts-sqlite.sql
sqlite> .schema
CREATE TABLE passengers (
    passenger_id INT PRIMARY KEY,
    passenger_name TEXT,
    favorite_mixins TEXT[],
    car_id INT
);
CREATE TABLE cocoa_cars (
    car_id INT PRIMARY KEY,
    available_mixins TEXT[],
    total_stock INT
);
sqlite> 
sqlite> .mode table
sqlite> select * from passengers limit 20;
+--------------+----------------+--------------------------------------------------------------+--------+
| passenger_id | passenger_name |                       favorite_mixins                        | car_id |
+--------------+----------------+--------------------------------------------------------------+--------+
| 1            | Ava Johnson    | ["vanilla foam"]                                             | 2      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 2            | Mateo Cruz     | ["caramel drizzle","shaved chocolate","white chocolate"]     | 2      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 3            | Nia Grant      | ["shaved chocolate"]                                         | 5      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 4            | Hiro Tanaka    | ["shaved chocolate"]                                         | 2      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 5            | Layla Brooks   | ["crispy rice","dark chocolate","caramel drizzle","cinnamon" | 3      |
|              |                | ]                                                            |        |
+--------------+----------------+--------------------------------------------------------------+--------+
| 6            | Ravi Patel     | ["caramel drizzle","shaved chocolate","white chocolate"]     | 5      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 7            | Sofia Kim      | ["cinnamon"]                                                 | 9      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 8            | Jonah Wolfe    | ["cinnamon","dark chocolate"]                                | 7      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 9            | Elena Morales  | ["white chocolate","shaved chocolate","caramel drizzle"]     | 6      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 10           | Diego Ramos    | ["shaved chocolate"]                                         | 1      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 11           | Zara Sheikh    | ["vanilla foam","crispy rice","peppermint"]                  | 4      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 12           | Caleb Osei     | ["shaved chocolate","dark chocolate","white chocolate"]      | 8      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 13           | Mila Novak     | ["crispy rice","cinnamon"]                                   | 4      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 14           | Lucas Ford     | ["vanilla foam","white chocolate","cinnamon"]                | 4      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 15           | Yara Haddad    | ["white chocolate","dark chocolate"]                         | 2      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 16           | Omar Qureshi   | ["marshmallow"]                                              | 3      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 17           | Keiko Ito      | ["vanilla foam","marshmallow"]                               | 7      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 18           | Tariq Hassan   | ["dark chocolate","crispy rice","white chocolate","peppermin | 2      |
|              |                | t"]                                                          |        |
+--------------+----------------+--------------------------------------------------------------+--------+
| 19           | Mira Zhao      | ["caramel drizzle","marshmallow","cinnamon"]                 | 7      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 20           | Bianca Pereira | ["dark chocolate","peppermint"]                              | 5      |
+--------------+----------------+--------------------------------------------------------------+--------+
sqlite> 
sqlite> select * from cocoa_cars;
+--------+--------------------------------------------------------------+-------------+
| car_id |                       available_mixins                       | total_stock |
+--------+--------------------------------------------------------------+-------------+
| 5      | ["white chocolate","shaved chocolate"]                       | 412         |
+--------+--------------------------------------------------------------+-------------+
| 2      | ["cinnamon","marshmallow","caramel drizzle"]                 | 359         |
+--------+--------------------------------------------------------------+-------------+
| 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         |
|        | ate"]                                                        |             |
+--------+--------------------------------------------------------------+-------------+
| 4      | ["shaved chocolate","white chocolate"]                       | 338         |
+--------+--------------------------------------------------------------+-------------+
| 8      | ["vanilla foam","marshmallow"]                               | 263         |
+--------+--------------------------------------------------------------+-------------+
| 1      | ["peppermint","crispy rice"]                                 | 205         |
+--------+--------------------------------------------------------------+-------------+
| 6      | ["shaved chocolate","dark chocolate","crispy rice","cinnamon | 161         |
|        | ","peppermint"]                                              |             |
+--------+--------------------------------------------------------------+-------------+
| 7      | ["caramel drizzle","crispy rice","marshmallow","vanilla foam | 132         |
|        | ","cinnamon"]                                                |             |
+--------+--------------------------------------------------------------+-------------+
| 3      | ["vanilla foam","peppermint"]                                | 95          |
+--------+--------------------------------------------------------------+-------------+
sqlite> 
```

Here's your full SQL file:

```sql
DROP TABLE IF EXISTS passengers;
DROP TABLE IF EXISTS cocoa_cars;

CREATE TABLE passengers (
    passenger_id INT PRIMARY KEY,
    passenger_name TEXT,
    favorite_mixins TEXT[],
    car_id INT
);

CREATE TABLE cocoa_cars (
    car_id INT PRIMARY KEY,
    available_mixins TEXT[],
    total_stock INT
);

INSERT INTO passengers (passenger_id, passenger_name, favorite_mixins, car_id) VALUES
    (1, 'Ava Johnson', '["vanilla foam"]', 2),
    (2, 'Mateo Cruz', '["caramel drizzle","shaved chocolate","white chocolate"]', 2),
    (3, 'Nia Grant', '["shaved chocolate"]', 5),
    (4, 'Hiro Tanaka', '["shaved chocolate"]', 2),
    (5, 'Layla Brooks', '["crispy rice","dark chocolate","caramel drizzle","cinnamon"]', 3),
    (6, 'Ravi Patel', '["caramel drizzle","shaved chocolate","white chocolate"]', 5),
    (7, 'Sofia Kim', '["cinnamon"]', 9),
    (8, 'Jonah Wolfe', '["cinnamon","dark chocolate"]', 7),
    (9, 'Elena Morales', '["white chocolate","shaved chocolate","caramel drizzle"]', 6),
    (10, 'Diego Ramos', '["shaved chocolate"]', 1),
    (11, 'Zara Sheikh', '["vanilla foam","crispy rice","peppermint"]', 4),
    (12, 'Caleb Osei', '["shaved chocolate","dark chocolate","white chocolate"]', 8),
    (13, 'Mila Novak', '["crispy rice","cinnamon"]', 4),
    (14, 'Lucas Ford', '["vanilla foam","white chocolate","cinnamon"]', 4),
    (15, 'Yara Haddad', '["white chocolate","dark chocolate"]', 2),
    (16, 'Omar Qureshi', '["marshmallow"]', 3),
    (17, 'Keiko Ito', '["vanilla foam","marshmallow"]', 7),
    (18, 'Tariq Hassan', '["dark chocolate","crispy rice","white chocolate","peppermint"]', 2),
    (19, 'Mira Zhao', '["caramel drizzle","marshmallow","cinnamon"]', 7),
    (20, 'Bianca Pereira', '["dark chocolate","peppermint"]', 5),
    (21, 'Eva Schmidt', '["white chocolate","marshmallow"]', 5),
    (22, 'Rafael Silva', '["cinnamon","caramel drizzle"]', 3),
    (23, 'Nolan Murphy', '["caramel drizzle"]', 8),
    (24, 'Sara Johansson', '["marshmallow"]', 6),
    (25, 'Ingrid Nilsen', '["shaved chocolate","peppermint","marshmallow"]', 2),
    (26, 'Arjun Kapoor', '["dark chocolate"]', 2),
    (27, 'Nova Adams', '["cinnamon","dark chocolate"]', 9),
    (28, 'Felix Schneider', '["crispy rice","vanilla foam"]', 4),
    (29, 'Tim Cook', '["crispy rice","vanilla foam"]', 6),
    (30, 'Sophia Rossi', '["white chocolate","dark chocolate","peppermint","marshmallow"]', 4),
    (31, 'Liam OConnor', '["caramel drizzle"]', 1),
    (32, 'Olivia Dubois', '["shaved chocolate","peppermint"]', 2),
    (33, 'Emma Svensson', '["shaved chocolate"]', 2),
    (34, 'Noah Fischer', '["caramel drizzle"]', 2),
    (35, 'William Becker', '["crispy rice","dark chocolate"]', 4),
    (36, 'Isabella Laurent', '["dark chocolate","shaved chocolate"]', 8),
    (37, 'James Kim', '["shaved chocolate","marshmallow","peppermint","caramel drizzle"]', 7),
    (38, 'Mia Chen', '["vanilla foam","white chocolate","shaved chocolate"]', 1),
    (39, 'Benjamin Patel', '["peppermint"]', 7),
    (40, 'Charlotte Singh', '["marshmallow","shaved chocolate","white chocolate"]', 4),
    (41, 'Daniel Murphy', '["cinnamon","vanilla foam","marshmallow","white chocolate"]', 8),
    (42, 'Zoe Wilson', '["marshmallow","dark chocolate"]', 9),
    (43, 'Robert Smith', '["peppermint"]', 9),
    (44, 'Emily Johnson', '["marshmallow"]', 4),
    (45, 'David Brown', '["vanilla foam","dark chocolate"]', 8),
    (46, 'Sarah Davis', '["vanilla foam","peppermint"]', 3),
    (47, 'James Simon', '["peppermint","vanilla foam","cinnamon","shaved chocolate"]', 5),
    (48, 'Linda Lee', '["white chocolate","dark chocolate","marshmallow","vanilla foam"]', 5),
    (49, 'Carlos Mendez', '["peppermint","white chocolate"]', 6),
    (50, 'Fatima Noor', '["peppermint"]', 8),
    (51, 'Youssef El-Sayed', '["peppermint","marshmallow"]', 3),
    (52, 'Ian Landsman', '["marshmallow"]', 4),
    (53, 'Nolan Young', '["marshmallow","shaved chocolate","crispy rice","vanilla foam"]', 1),
    (54, 'Ava Martinez', '["vanilla foam"]', 9),
    (55, 'William Chen', '["crispy rice","shaved chocolate","caramel drizzle"]', 6),
    (56, 'Isabella Rodriguez', '["crispy rice","vanilla foam"]', 3),
    (57, 'Zachary Collins', '["dark chocolate","caramel drizzle","vanilla foam"]', 2),
    (58, 'Audrey Edwards', '["dark chocolate"]', 2),
    (59, 'Jason Stewart', '["white chocolate"]', 4),
    (60, 'Lucy Morris', '["cinnamon","caramel drizzle","peppermint"]', 4),
    (61, 'Cameron Rogers', '["crispy rice","cinnamon","shaved chocolate"]', 9),
    (62, 'Aria Blackwood', '["white chocolate","peppermint","caramel drizzle"]', 9),
    (63, 'Felix Whitmore', '["marshmallow","cinnamon","dark chocolate"]', 2),
    (64, 'Luna Hartley', '["white chocolate"]', 3),
    (65, 'Jasper Thorne', '["crispy rice","shaved chocolate","caramel drizzle"]', 6),
    (66, 'Nora Calloway', '["crispy rice","dark chocolate"]', 5),
    (67, 'Silas Merrick', '["marshmallow"]', 7),
    (68, 'Iris Pembroke', '["peppermint","white chocolate","cinnamon"]', 3),
    (69, 'Milo Ashford', '["cinnamon","dark chocolate","crispy rice"]', 7),
    (70, 'Clara Westbrook', '["marshmallow"]', 2),
    (71, 'Owen Fairchild', '["white chocolate","peppermint"]', 6),
    (72, 'Ruby Hawthorne', '["vanilla foam","cinnamon"]', 1),
    (73, 'Finn Lockhart', '["caramel drizzle","peppermint","cinnamon"]', 4),
    (74, 'Violet Sterling', '["marshmallow","caramel drizzle"]', 9),
    (75, 'August Blackwell', '["cinnamon","shaved chocolate","vanilla foam","marshmallow"]', 3),
    (76, 'Hazel Kincaid', '["peppermint","cinnamon","caramel drizzle","dark chocolate"]', 7),
    (77, 'Leo Greyson', '["crispy rice","cinnamon"]', 2),
    (78, 'Stella Beaumont', '["peppermint","dark chocolate","marshmallow","vanilla foam"]', 8),
    (79, 'Miles Brennan', '["crispy rice","shaved chocolate","marshmallow"]', 1),
    (80, 'Ivy Winslow', '["vanilla foam","caramel drizzle"]', 5),
    (81, 'Jack Carmichael', '["crispy rice"]', 6),
    (82, 'Scarlett Dalton', '["white chocolate","caramel drizzle","peppermint","vanilla foam"]', 5),
    (83, 'Oliver Ashby', '["crispy rice","peppermint"]', 2),
    (84, 'Aurora Whitfield', '["caramel drizzle","white chocolate","shaved chocolate","crispy rice"]', 9),
    (85, 'Noah Hastings', '["vanilla foam"]', 4),
    (86, 'Eliza Radcliffe', '["peppermint","vanilla foam","white chocolate"]', 9),
    (87, 'Liam Donovan', '["caramel drizzle","vanilla foam"]', 2),
    (88, 'Penelope Sinclair', '["caramel drizzle","marshmallow","white chocolate"]', 5),
    (89, 'Ethan Marlowe', '["vanilla foam","caramel drizzle","shaved chocolate"]', 5),
    (90, 'Charlotte Waverly', '["shaved chocolate","vanilla foam"]', 7),
    (91, 'Lucas Prescott', '["crispy rice","vanilla foam"]', 9),
    (92, 'Amelia Rosewood', '["crispy rice"]', 5),
    (93, 'Henry Treadwell', '["vanilla foam","caramel drizzle"]', 8),
    (94, 'Sophie Langford', '["dark chocolate","shaved chocolate","crispy rice","white chocolate"]', 3),
    (95, 'Benjamin Fairweather', '["crispy rice"]', 9),
    (96, 'Grace Aldridge', '["marshmallow","shaved chocolate","caramel drizzle"]', 5),
    (97, 'Samuel Kingsley', '["shaved chocolate","cinnamon"]', 1),
    (98, 'Eleanor Morrison', '["shaved chocolate"]', 8),
    (99, 'Daniel Lockwood', '["dark chocolate"]', 7),
    (100, 'Lucy Harrington', '["vanilla foam","dark chocolate"]', 7),
    (101, 'Matthew Sutherland', '["cinnamon","peppermint"]', 2),
    (102, 'Emma Gilmore', '["shaved chocolate","cinnamon","vanilla foam","caramel drizzle"]', 9),
    (103, 'Alexander Stratton', '["peppermint","shaved chocolate","vanilla foam","white chocolate"]', 8),
    (104, 'Abigail Worthington', '["dark chocolate","caramel drizzle"]', 8),
    (105, 'William Beauchamp', '["white chocolate","dark chocolate","marshmallow","caramel drizzle"]', 8),
    (106, 'Hannah Livingston', '["crispy rice","shaved chocolate","vanilla foam","caramel drizzle"]', 5),
    (107, 'James Garrison', '["shaved chocolate","crispy rice","white chocolate","peppermint"]', 5),
    (108, 'Sophia Brookshire', '["crispy rice","caramel drizzle"]', 6),
    (109, 'Theodore Hadley', '["cinnamon"]', 3),
    (110, 'Olivia Carrington', '["vanilla foam","cinnamon"]', 4),
    (111, 'Sebastian Ashworth', '["vanilla foam"]', 7),
    (112, 'Chloe Blackstone', '["white chocolate","dark chocolate","shaved chocolate"]', 1),
    (113, 'Nicholas Montague', '["vanilla foam","white chocolate"]', 1),
    (114, 'Madeline Ramsey', '["dark chocolate","peppermint","cinnamon","vanilla foam"]', 7),
    (115, 'Gabriel Winthrop', '["white chocolate","shaved chocolate","dark chocolate","marshmallow"]', 5),
    (116, 'Alice Merriweather', '["dark chocolate","peppermint","shaved chocolate","cinnamon"]', 7),
    (117, 'Isaac Kendrick', '["dark chocolate","cinnamon"]', 9),
    (118, 'Lillian Holbrook', '["vanilla foam"]', 1),
    (119, 'Caleb Bellamy', '["vanilla foam"]', 3),
    (120, 'Rose Drummond', '["cinnamon","peppermint","white chocolate","shaved chocolate"]', 6),
    (121, 'Elijah Wakefield', '["dark chocolate","caramel drizzle"]', 6),
    (122, 'Margaret Fairbanks', '["crispy rice","vanilla foam","cinnamon","peppermint"]', 8),
    (123, 'Julian Blackburn', '["white chocolate"]', 1),
    (124, 'Eva Templeton', '["shaved chocolate","marshmallow","vanilla foam"]', 1),
    (125, 'Nathan Whitley', '["shaved chocolate"]', 4),
    (126, 'Anna Westfield', '["cinnamon"]', 4),
    (127, 'Aaron Ashcroft', '["dark chocolate","marshmallow"]', 4),
    (128, 'Julia Pendleton', '["crispy rice","caramel drizzle","marshmallow","white chocolate"]', 2),
    (129, 'Connor Redmond', '["crispy rice","marshmallow"]', 1),
    (130, 'Grace Thornhill', '["vanilla foam","white chocolate","caramel drizzle"]', 4),
    (131, 'Zachary Stafford', '["shaved chocolate"]', 2),
    (132, 'Caroline Bannister', '["marshmallow","peppermint","cinnamon"]', 9),
    (133, 'Dylan Blakely', '["caramel drizzle","marshmallow","crispy rice","white chocolate"]', 6),
    (134, 'Katherine Underwood', '["vanilla foam"]', 8),
    (135, 'Tyler Braddock', '["vanilla foam"]', 6),
    (136, 'Victoria Harwood', '["cinnamon","vanilla foam","marshmallow","caramel drizzle"]', 9),
    (137, 'Ryan Beckett', '["white chocolate","dark chocolate","shaved chocolate"]', 7),
    (138, 'Elizabeth Chesterfield', '["caramel drizzle","shaved chocolate","vanilla foam"]', 2),
    (139, 'Jordan Waverly', '["dark chocolate","shaved chocolate","vanilla foam"]', 8),
    (140, 'Sarah Remington', '["caramel drizzle","peppermint","shaved chocolate","cinnamon"]', 3),
    (141, 'Brandon Locklear', '["shaved chocolate","caramel drizzle","vanilla foam","cinnamon"]', 6),
    (142, 'Rachel Wyndham', '["crispy rice","peppermint","white chocolate"]', 4),
    (143, 'Logan Sherwood', '["shaved chocolate"]', 7),
    (144, 'Amanda Fitzroy', '["white chocolate","shaved chocolate","caramel drizzle","dark chocolate"]', 8),
    (145, 'Jackson Thorpe', '["peppermint","marshmallow","cinnamon","dark chocolate"]', 7),
    (146, 'Rebecca Ashcombe', '["crispy rice","caramel drizzle"]', 8),
    (147, 'Cameron Gladstone', '["vanilla foam","caramel drizzle","cinnamon"]', 8),
    (148, 'Jessica Langston', '["crispy rice","white chocolate","marshmallow"]', 2),
    (149, 'Mason Fairmont', '["caramel drizzle","marshmallow"]', 9),
    (150, 'Emily Claridge', '["shaved chocolate","white chocolate"]', 8),
    (151, 'Hunter Bellingham', '["white chocolate","crispy rice","peppermint"]', 4),
    (152, 'Laura Thornbury', '["shaved chocolate","caramel drizzle","marshmallow"]', 5),
    (153, 'Wyatt Alderton', '["white chocolate"]', 3),
    (154, 'Claire Berkshire', '["peppermint","white chocolate","crispy rice"]', 5),
    (155, 'Cole Ashland', '["dark chocolate","marshmallow"]', 1),
    (156, 'Diana Brightwell', '["dark chocolate","white chocolate","shaved chocolate"]', 6),
    (157, 'Aiden Stanfield', '["peppermint","crispy rice"]', 8),
    (158, 'Natalie Warwick', '["marshmallow"]', 7),
    (159, 'Parker Blackmore', '["marshmallow","peppermint","white chocolate","vanilla foam"]', 5),
    (160, 'Morgan Steadman', '["shaved chocolate"]', 2),
    (161, 'Blake Dunwood', '["shaved chocolate","vanilla foam","white chocolate","dark chocolate"]', 5),
    (162, 'Taylor Woodridge', '["crispy rice","peppermint","white chocolate","caramel drizzle"]', 2),
    (163, 'Chase Ashbury', '["shaved chocolate","crispy rice"]', 2),
    (164, 'Madison Clearwater', '["shaved chocolate","cinnamon"]', 9),
    (165, 'Carter Brookfield', '["cinnamon"]', 1),
    (166, 'Ashley Fairhaven', '["dark chocolate","white chocolate","cinnamon","peppermint"]', 4),
    (167, 'Griffin Hartwell', '["crispy rice","dark chocolate","peppermint"]', 4),
    (168, 'Megan Redfield', '["shaved chocolate","vanilla foam","peppermint"]', 9),
    (169, 'Grayson Westmore', '["cinnamon","crispy rice"]', 3),
    (170, 'Nicole Ashridge', '["peppermint"]', 3),
    (171, 'Sawyer Hollingsworth', '["crispy rice","dark chocolate","peppermint"]', 8),
    (172, 'Alexis Thorndale', '["vanilla foam","crispy rice","dark chocolate"]', 9),
    (173, 'Declan Summerfield', '["dark chocolate","marshmallow","crispy rice","peppermint"]', 7),
    (174, 'Samantha Brightwood', '["crispy rice","peppermint","dark chocolate"]', 4),
    (175, 'Tristan Ashbrook', '["crispy rice"]', 1),
    (176, 'Melissa Ravenscroft', '["dark chocolate","white chocolate"]', 5),
    (177, 'Colton Hawthorne', '["vanilla foam","dark chocolate"]', 2),
    (178, 'Lauren Silverton', '["caramel drizzle","vanilla foam","cinnamon","dark chocolate"]', 2),
    (179, 'Landon Whitworth', '["caramel drizzle","vanilla foam"]', 8),
    (180, 'Kayla Mansfield', '["vanilla foam","peppermint","shaved chocolate"]', 2);

INSERT INTO cocoa_cars (car_id, available_mixins, total_stock) VALUES
    (5, '["white chocolate","shaved chocolate"]', 412),
    (2, '["cinnamon","marshmallow","caramel drizzle"]', 359),
    (9, '["crispy rice","peppermint","caramel drizzle","shaved chocolate"]', 354),
    (4, '["shaved chocolate","white chocolate"]', 338),
    (8, '["vanilla foam","marshmallow"]', 263),
    (1, '["peppermint","crispy rice"]', 205),
    (6, '["shaved chocolate","dark chocolate","crispy rice","cinnamon","peppermint"]', 161),
    (7, '["caramel drizzle","crispy rice","marshmallow","vanilla foam","cinnamon"]', 132),
    (3, '["vanilla foam","peppermint"]', 95);

```

Changing `TEXT[]` to `TEXT` is not required, as SQLITE really doesn't bother with the types unless the table is strict. We'll take that flexibility here.

## Problem

> Get the stewards a list of all the passengers and the cocoa car(s) they can be served from that has at least one of their favorite mixins.
> 
> Remember only the top three most-stocked cocoa cars remained operational, so the passengers must be served from one of those cars.

Ok so, we have two tables.

1. `cocoa_cars`
2. `passengers`

What we need to do is to list down the passengers which the `cocoa_cars` can satisfy their atleast one mixin(food) with their car_ids. So, there are limited cars, just 9 cars in total right?

```sql
SELECT * FROM cocoa_cars ORDER BY car_id;
```
So, there are only 9 cars.
```
sqlite> SELECT * FROM cocoa_cars ORDER BY car_id;
+--------+--------------------------------------------------------------+-------------+
| car_id |                       available_mixins                       | total_stock |
+--------+--------------------------------------------------------------+-------------+
| 1      | ["peppermint","crispy rice"]                                 | 205         |
+--------+--------------------------------------------------------------+-------------+
| 2      | ["cinnamon","marshmallow","caramel drizzle"]                 | 359         |
+--------+--------------------------------------------------------------+-------------+
| 3      | ["vanilla foam","peppermint"]                                | 95          |
+--------+--------------------------------------------------------------+-------------+
| 4      | ["shaved chocolate","white chocolate"]                       | 338         |
+--------+--------------------------------------------------------------+-------------+
| 5      | ["white chocolate","shaved chocolate"]                       | 412         |
+--------+--------------------------------------------------------------+-------------+
| 6      | ["shaved chocolate","dark chocolate","crispy rice","cinnamon | 161         |
|        | ","peppermint"]                                              |             |
+--------+--------------------------------------------------------------+-------------+
| 7      | ["caramel drizzle","crispy rice","marshmallow","vanilla foam | 132         |
|        | ","cinnamon"]                                                |             |
+--------+--------------------------------------------------------------+-------------+
| 8      | ["vanilla foam","marshmallow"]                               | 263         |
+--------+--------------------------------------------------------------+-------------+
| 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         |
|        | ate"]                                                        |             |
+--------+--------------------------------------------------------------+-------------+
sqlite> 
```

However, the question states, we only need to consider the top three most stacked cars. Now we can order by `total_stack` here and find the top stacked 3 cars with the `LIMIT 3` clause.

```sql
SELECT * FROM cocoa_cars ORDER BY total_stock DESC LIMIT 3;
```

```
sqlite> 
sqlite> select * from cocoa_cars ORDER BY total_stock DESC;
+--------+--------------------------------------------------------------+-------------+
| car_id |                       available_mixins                       | total_stock |
+--------+--------------------------------------------------------------+-------------+
| 5      | ["white chocolate","shaved chocolate"]                       | 412         |
+--------+--------------------------------------------------------------+-------------+
| 2      | ["cinnamon","marshmallow","caramel drizzle"]                 | 359         |
+--------+--------------------------------------------------------------+-------------+
| 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         |
|        | ate"]                                                        |             |
+--------+--------------------------------------------------------------+-------------+
| 4      | ["shaved chocolate","white chocolate"]                       | 338         |
+--------+--------------------------------------------------------------+-------------+
| 8      | ["vanilla foam","marshmallow"]                               | 263         |
+--------+--------------------------------------------------------------+-------------+
| 1      | ["peppermint","crispy rice"]                                 | 205         |
+--------+--------------------------------------------------------------+-------------+
| 6      | ["shaved chocolate","dark chocolate","crispy rice","cinnamon | 161         |
|        | ","peppermint"]                                              |             |
+--------+--------------------------------------------------------------+-------------+
| 7      | ["caramel drizzle","crispy rice","marshmallow","vanilla foam | 132         |
|        | ","cinnamon"]                                                |             |
+--------+--------------------------------------------------------------+-------------+
| 3      | ["vanilla foam","peppermint"]                                | 95          |
+--------+--------------------------------------------------------------+-------------+
sqlite> select * from cocoa_cars ORDER BY total_stock DESC LIMIT 3;
+--------+--------------------------------------------------------------+-------------+
| car_id |                       available_mixins                       | total_stock |
+--------+--------------------------------------------------------------+-------------+
| 5      | ["white chocolate","shaved chocolate"]                       | 412         |
+--------+--------------------------------------------------------------+-------------+
| 2      | ["cinnamon","marshmallow","caramel drizzle"]                 | 359         |
+--------+--------------------------------------------------------------+-------------+
| 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         |
|        | ate"]                                                        |             |
+--------+--------------------------------------------------------------+-------------+
sqlite> 

```

Now, we have all the cars that we can start assigning the passengers.
How?

We need to select and assign the passenger the `car_id` which contains one or more of their `favourite_mixins`.
Now, this is the tricky part.

We are in SQLite!

We have a `favourite_list` for `Mateo Cruz` as a string like `["caramel drizzle","shaved chocolate","white chocolate"]` which we need to match against these 3 car `available_mixins`:
- `["white chocolate","shaved chocolate"]` on car_id `5`
- `["cinnamon","marshmallow","caramel drizzle"]` on car_id `2`
- `["crispy rice","peppermint","caramel drizzle","shaved chocolate"]` on car_id `9`


So here all three cars have at least one right?
- car_id `5` has `shaved_chocolate`
- car_id `2` has `caramel drizzle`
- car_id `9` has 2 of his 3 `favourite_list` dishes `caramel drizzle`, and `shaved chocolate"`.

So we should ideally return for `Matro Cruz` the car_ids `[5, 2, 9]` or as separate rows doesn't matter as much I think. But the first one looks cool!

So, how do we do it?

First let's start with what we had!

The top 3 cars as

```sql
SELECT * FROM cocoa_cars ORDER BY total_stock DESC LIMIT 3;
```

This gives us the table that we can use to compare for each car_id if we have any one of the `favourite_list` for one passenger at a time.

But how do we split the string of mixins and favourite_list?

We can use the [json_each](https://sqlite.org/json1.html#jeach) function which takes in any valid json string (could be raw, could be a column name) for each row of the the table. And it returns back a lot of things.

Let's just try to select everything from the `json_each` with the favourite_mixins column.

```sql
SELECT * FROM json_each(favorite_mixins) FROM passengers;
```

```sqlite> SELECT * FROM json_each(favorite_mixins) FROM passengers;
Parse error: near "FROM": syntax error
  SELECT * FROM json_each(favorite_mixins) FROM passengers;
                             error here ---^
sqlite> 
```
Ops!
Why?
Because we have given it the whole column, it can only take one cell at a time, so we need to give it that cell for each row.


```sql
SELECT * FROM json_each((SELECT favorite_mixins FROM passengers));
```

Here we try to give it only the `favorite_mixins` from the passengers column.

```
sqlite> SELECT * FROM json_each((SELECT favorite_mixins FROM passengers));
+-----+--------------+------+--------------+----+--------+---------+------+
| key |    value     | type |     atom     | id | parent | fullkey | path |
+-----+--------------+------+--------------+----+--------+---------+------+
| 0   | vanilla foam | text | vanilla foam | 2  |        | $[0]    | $    |
+-----+--------------+------+--------------+----+--------+---------+------+
sqlite> select * from passengers limit 5;
+--------------+----------------+--------------------------------------------------------------+--------+
| passenger_id | passenger_name |                       favorite_mixins                        | car_id |
+--------------+----------------+--------------------------------------------------------------+--------+
| 1            | Ava Johnson    | ["vanilla foam"]                                             | 2      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 2            | Mateo Cruz     | ["caramel drizzle","shaved chocolate","white chocolate"]     | 2      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 3            | Nia Grant      | ["shaved chocolate"]                                         | 5      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 4            | Hiro Tanaka    | ["shaved chocolate"]                                         | 2      |
+--------------+----------------+--------------------------------------------------------------+--------+
| 5            | Layla Brooks   | ["crispy rice","dark chocolate","caramel drizzle","cinnamon" | 3      |
|              |                | ]                                                            |        |
+--------------+----------------+--------------------------------------------------------------+--------+
sqlite> 
```
But it only gave it for the first column, unluckily it had just one mixin.

What is happening?

Well!

This actually selects only the first row's `favorite_mixins` because `json_each()` processes one value at a time. Since favorite_mixins is a JSON array, SQLite expects a single array value per row. When we try to pass the entire column, it only processes the first row of `favorite_mixins`

Let's try to use the json_each for passengers with more than one favorite mixins.

```sql
SELECT * 
FROM json_each(
     (SELECT favorite_mixins FROM passengers WHERE passenger_name = 'Mateo Cruz')
);
```


```
sqlite> SELECT * 
FROM json_each((SELECT favorite_mixins FROM passengers WHERE passenger_name = 'Ava Johnson'));
+-----+--------------+------+--------------+----+--------+---------+------+
| key |    value     | type |     atom     | id | parent | fullkey | path |
+-----+--------------+------+--------------+----+--------+---------+------+
| 0   | vanilla foam | text | vanilla foam | 2  |        | $[0]    | $    |
+-----+--------------+------+--------------+----+--------+---------+------+
sqlite> SELECT * 
FROM json_each((SELECT favorite_mixins FROM passengers WHERE passenger_name = 'Mateo Cruz'));
+-----+------------------+------+------------------+----+--------+---------+------+
| key |      value       | type |       atom       | id | parent | fullkey | path |
+-----+------------------+------+------------------+----+--------+---------+------+
| 0   | caramel drizzle  | text | caramel drizzle  | 2  |        | $[0]    | $    |
| 1   | shaved chocolate | text | shaved chocolate | 19 |        | $[1]    | $    |
| 2   | white chocolate  | text | white chocolate  | 37 |        | $[2]    | $    |
+-----+------------------+------+------------------+----+--------+---------+------+
sqlite> 
```

Ok now that is neat, it returned 3 rows for the 3 favorite items for the passenger `Mateo Cruz`

Now what?

How do we get it for all the passengers? 

How about JOINs, since the passenger's data will remain the same, we just change the mixins for each of their favorite list.


```sql
SELECT *
FROM passengers
JOIN json_each(passengers.favorite_mixins) AS mixin 
ORDER BY passengers.passenger_name;
```

```
sqlite> sqlite> SELECT *
FROM passengers
JOIN json_each(passengers.favorite_mixins) AS mixin 
ORDER BY passengers.passenger_name;
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| passenger_id |     passenger_name     |                       favorite_mixins                        | car_id | key |      value       | type |       atom       | id | parent | fullkey | path |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 127          | Aaron Ashcroft         | ["dark chocolate","marshmallow"]                             | 4      | 0   | dark chocolate   | text | dark chocolate   | 2  |        | $[0]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 127          | Aaron Ashcroft         | ["dark chocolate","marshmallow"]                             | 4      | 1   | marshmallow      | text | marshmallow      | 18 |        | $[1]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 104          | Abigail Worthington    | ["dark chocolate","caramel drizzle"]                         | 8      | 0   | dark chocolate   | text | dark chocolate   | 2  |        | $[0]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 104          | Abigail Worthington    | ["dark chocolate","caramel drizzle"]                         | 8      | 1   | caramel drizzle  | text | caramel drizzle  | 18 |        | $[1]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 157          | Aiden Stanfield        | ["peppermint","crispy rice"]                                 | 8      | 0   | peppermint       | text | peppermint       | 2  |        | $[0]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 157          | Aiden Stanfield        | ["peppermint","crispy rice"]                                 | 8      | 1   | crispy rice      | text | crispy rice      | 13 |        | $[1]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 103          | Alexander Stratton     | ["peppermint","shaved chocolate","vanilla foam","white choco | 8      | 0   | peppermint       | text | peppermint       | 2  |        | $[0]    | $    |
|              |                        | late"]                                                       |        |     |                  |      |                  |    |        |         |      |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 103          | Alexander Stratton     | ["peppermint","shaved chocolate","vanilla foam","white choco | 8      | 1   | shaved chocolate | text | shaved chocolate | 13 |        | $[1]    | $    |
|              |                        | late"]                                                       |        |     |                  |      |                  |    |        |         |      |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 103          | Alexander Stratton     | ["peppermint","shaved chocolate","vanilla foam","white choco | 8      | 2   | vanilla foam     | text | vanilla foam     | 31 |        | $[2]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
...
...
| 153          | Wyatt Alderton         | ["white chocolate"]                                          | 3      | 0   | white chocolate  | text | white chocolate  | 2  |        | $[0]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 15           | Yara Haddad            | ["white chocolate","dark chocolate"]                         | 2      | 0   | white chocolate  | text | white chocolate  | 2  |        | $[0]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 15           | Yara Haddad            | ["white chocolate","dark chocolate"]                         | 2      | 1   | dark chocolate   | text | dark chocolate   | 19 |        | $[1]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 51           | Youssef El-Sayed       | ["peppermint","marshmallow"]                                 | 3      | 0   | peppermint       | text | peppermint       | 2  |        | $[0]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 51           | Youssef El-Sayed       | ["peppermint","marshmallow"]                                 | 3      | 1   | marshmallow      | text | marshmallow      | 13 |        | $[1]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 57           | Zachary Collins        | ["dark chocolate","caramel drizzle","vanilla foam"]          | 2      | 0   | dark chocolate   | text | dark chocolate   | 2  |        | $[0]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 57           | Zachary Collins        | ["dark chocolate","caramel drizzle","vanilla foam"]          | 2      | 1   | caramel drizzle  | text | caramel drizzle  | 18 |        | $[1]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 57           | Zachary Collins        | ["dark chocolate","caramel drizzle","vanilla foam"]          | 2      | 2   | vanilla foam     | text | vanilla foam     | 35 |        | $[2]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 131          | Zachary Stafford       | ["shaved chocolate"]                                         | 2      | 0   | shaved chocolate | text | shaved chocolate | 2  |        | $[0]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 11           | Zara Sheikh            | ["vanilla foam","crispy rice","peppermint"]                  | 4      | 0   | vanilla foam     | text | vanilla foam     | 2  |        | $[0]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 11           | Zara Sheikh            | ["vanilla foam","crispy rice","peppermint"]                  | 4      | 1   | crispy rice      | text | crispy rice      | 16 |        | $[1]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 11           | Zara Sheikh            | ["vanilla foam","crispy rice","peppermint"]                  | 4      | 2   | peppermint       | text | peppermint       | 28 |        | $[2]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 42           | Zoe Wilson             | ["marshmallow","dark chocolate"]                             | 9      | 0   | marshmallow      | text | marshmallow      | 2  |        | $[0]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+
| 42           | Zoe Wilson             | ["marshmallow","dark chocolate"]                             | 9      | 1   | dark chocolate   | text | dark chocolate   | 14 |        | $[1]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+

```

It just automagically joins the relevant rows for the expanded rows from the json_each when we pass the column which is a json list.

But we don't want everything, we just want the passenger and the mixin names. And let's also include the index just to see the data.

```sql
SELECT passengers.passenger_name, mixin.key AS mixin_index, mixin.value AS mixin
FROM passengers
JOIN json_each(passengers.favorite_mixins) AS mixin
ORDER BY passengers.passenger_name;
```

```
sqlite> SELECT passengers.passenger_name, mixin.key AS mixin_index, mixin.value AS mixin
FROM passengers
JOIN json_each(passengers.favorite_mixins) AS mixin
ORDER BY passengers.passenger_name;
+------------------------+-------------+------------------+
|     passenger_name     | mixin_index |      mixin       |
+------------------------+-------------+------------------+
| Aaron Ashcroft         | 0           | dark chocolate   |
| Aaron Ashcroft         | 1           | marshmallow      |
| Abigail Worthington    | 0           | dark chocolate   |
| Abigail Worthington    | 1           | caramel drizzle  |
| Aiden Stanfield        | 0           | peppermint       |
| Aiden Stanfield        | 1           | crispy rice      |
| Alexander Stratton     | 0           | peppermint       |
| Alexander Stratton     | 1           | shaved chocolate |
| Alexander Stratton     | 2           | vanilla foam     |
| Alexander Stratton     | 3           | white chocolate  |
...
...
| Yara Haddad            | 1           | dark chocolate   |
| Youssef El-Sayed       | 0           | peppermint       |
| Youssef El-Sayed       | 1           | marshmallow      |
| Zachary Collins        | 0           | dark chocolate   |
| Zachary Collins        | 1           | caramel drizzle  |
| Zachary Collins        | 2           | vanilla foam     |
| Zachary Stafford       | 0           | shaved chocolate |
| Zara Sheikh            | 0           | vanilla foam     |
| Zara Sheikh            | 1           | crispy rice      |
| Zara Sheikh            | 2           | peppermint       |
| Zoe Wilson             | 0           | marshmallow      |
| Zoe Wilson             | 1           | dark chocolate   |
+------------------------+-------------+------------------+

```

That looks good.

Now what?

We also need to do it for the cocoa_cars with `available_mixins`

```sql
SELECT 
    *
FROM cocoa_cars 
JOIN 
    json_each(available_mixins) AS mixin
ORDER BY total_stock DESC;
```
OR

```sql
SELECT 
    car_id,
    mixin.key AS mixin_index,
    mixin.value AS mixin 
FROM cocoa_cars 
JOIN
    json_each(available_mixins) as mixin 
ORDER BY total_stock DESC;
```

```
sqlite> SELECT car_id, mixin.key as mixin_index, mixin.value as mixin from cocoa_cars join json_each(available_mixins) as mixin ORDER BY total_stock DESC;
+--------+-------------+------------------+
| car_id | mixin_index |      mixin       |
+--------+-------------+------------------+
| 5      | 0           | white chocolate  |
| 5      | 1           | shaved chocolate |
| 2      | 0           | cinnamon         |
| 2      | 1           | marshmallow      |
| 2      | 2           | caramel drizzle  |
| 9      | 0           | crispy rice      |
| 9      | 1           | peppermint       |
| 9      | 2           | caramel drizzle  |
| 9      | 3           | shaved chocolate |
| 4      | 0           | shaved chocolate |
| 4      | 1           | white chocolate  |
| 8      | 0           | vanilla foam     |
| 8      | 1           | marshmallow      |
| 1      | 0           | peppermint       |
| 1      | 1           | crispy rice      |
| 6      | 0           | shaved chocolate |
| 6      | 1           | dark chocolate   |
| 6      | 2           | crispy rice      |
| 6      | 3           | cinnamon         |
| 6      | 4           | peppermint       |
| 7      | 0           | caramel drizzle  |
| 7      | 1           | crispy rice      |
| 7      | 2           | marshmallow      |
| 7      | 3           | vanilla foam     |
| 7      | 4           | cinnamon         |
| 3      | 0           | vanilla foam     |
| 3      | 1           | peppermint       |
+--------+-------------+------------------+
sqlite> 
```

But hold on! We only needed it for the top 3 stocked cars.

So, how do we do it, we can (bad and dirty practise) limit by the count of the numbers of mixin in 5, 2 and 9 but that is bad.

We need to dynamically get this table. 

Help in! CTEs

```sql
WITH stocked_cars as (
    SELECT * FROM cocoa_cars ORDER BY total_stock DESC LIMIT 3
)
SELECT 
    *
FROM stocked_cars
JOIN json_each(stocked_cars.available_mixins);
```

OR

```sql
WITH stocked_cars as (
    SELECT * FROM cocoa_cars ORDER BY total_stock DESC LIMIT 3
)
SELECT 
    car_id,
    car_mixins.key as car_mixin_index,
    car_mixins.value as car_mixin
FROM stocked_cars
JOIN json_each(stocked_cars.available_mixins) as car_mixins;
```


```
sqlite> WITH stocked_cars as (
(x1...> select * from cocoa_cars ORDER BY total_stock DESC LIMIT 3)
   ...> select * from stocked_cars JOIN json_each(stocked_cars.available_mixins);
+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| car_id |                       available_mixins                       | total_stock | key |      value       | type |       atom       | id | parent | fullkey | path |
+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 5      | ["white chocolate","shaved chocolate"]                       | 412         | 0   | white chocolate  | text | white chocolate  | 2  |        | $[0]    | $    |
+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 5      | ["white chocolate","shaved chocolate"]                       | 412         | 1   | shaved chocolate | text | shaved chocolate | 19 |        | $[1]    | $    |
+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 2      | ["cinnamon","marshmallow","caramel drizzle"]                 | 359         | 0   | cinnamon         | text | cinnamon         | 2  |        | $[0]    | $    |
+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 2      | ["cinnamon","marshmallow","caramel drizzle"]                 | 359         | 1   | marshmallow      | text | marshmallow      | 11 |        | $[1]    | $    |
+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 2      | ["cinnamon","marshmallow","caramel drizzle"]                 | 359         | 2   | caramel drizzle  | text | caramel drizzle  | 23 |        | $[2]    | $    |
+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 0   | crispy rice      | text | crispy rice      | 2  |        | $[0]    | $    |
|        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 1   | peppermint       | text | peppermint       | 14 |        | $[1]    | $    |
|        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 2   | caramel drizzle  | text | caramel drizzle  | 25 |        | $[2]    | $    |
|        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 3   | shaved chocolate | text | shaved chocolate | 42 |        | $[3]    | $    |
|        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
sqlite> 
sqlite> WITH stocked_cars as (
    SELECT * FROM cocoa_cars ORDER BY total_stock DESC LIMIT 3
)
SELECT 
    car_id,
    car_mixins.key as car_mixin_index,
    car_mixins.value as car_mixin
FROM stocked_cars
JOIN json_each(stocked_cars.available_mixins) as car_mixins;
+--------+-----------------+------------------+
| car_id | car_mixin_index |    car_mixin     |
+--------+-----------------+------------------+
| 5      | 0               | white chocolate  |
| 5      | 1               | shaved chocolate |
| 2      | 0               | cinnamon         |
| 2      | 1               | marshmallow      |
| 2      | 2               | caramel drizzle  |
| 9      | 0               | crispy rice      |
| 9      | 1               | peppermint       |
| 9      | 2               | caramel drizzle  |
| 9      | 3               | shaved chocolate |
+--------+-----------------+------------------+
sqlite> 

```

Now what next?

We simply need to combine all of it

- Grab the top 3 cars 
- Grab the passenger favorite mixins (expand with json_each)
- Grab the car available mixins (expand with json_each and use the top 3 car as cte)
- JOIN them when the car_mixin has one mixin from a passenger favorite mixin.

```sql
WITH stocked_cars as (
    SELECT * FROM cocoa_cars ORDER BY total_stock DESC LIMIT 3
)
SELECT passengers.passenger_name, stocked_cars.car_id
FROM passengers  
JOIN json_each(passengers.favorite_mixins) as passenger_mixin
JOIN stocked_cars  
JOIN json_each(stocked_cars.available_mixins) AS available_mixin
    ON passenger_mixin.value = available_mixin.value;
```

OR 

Select some more rows for visual confirmation.

```sql
WITH stocked_cars as (
    SELECT * FROM cocoa_cars ORDER BY total_stock DESC LIMIT 3
)
SELECT passengers.passenger_name, passengers.favorite_mixins, stocked_cars.available_mixins, passenger_mixin.value as passenger_mixin, available_mixin.value as available_mixin, stocked_cars.car_id
FROM passengers
JOIN json_each(passengers.favorite_mixins) as passenger_mixin
JOIN stocked_cars
JOIN json_each(stocked_cars.available_mixins) AS available_mixin
    ON passenger_mixin.value = available_mixin.value;

```

OR

select everything 

```sql
WITH stocked_cars as (
    SELECT * FROM cocoa_cars ORDER BY total_stock DESC LIMIT 3
)
SELECT *
FROM passengers
JOIN json_each(passengers.favorite_mixins) as passenger_mixin
JOIN stocked_cars
JOIN json_each(stocked_cars.available_mixins) AS available_mixin
    ON passenger_mixin.value = available_mixin.value;

```


```
sqlite> WITH stocked_cars as (
    SELECT * FROM cocoa_cars ORDER BY total_stock DESC LIMIT 3
)
SELECT passengers.passenger_name, stocked_cars.car_id
FROM passengers  
JOIN json_each(passengers.favorite_mixins) as passenger_mixin
JOIN stocked_cars  
JOIN json_each(stocked_cars.available_mixins) AS available_mixin
    ON passenger_mixin.value = available_mixin.value;

+------------------------+--------+
|     passenger_name     | car_id |
+------------------------+--------+
| Mateo Cruz             | 5      |
| Mateo Cruz             | 5      |
| Nia Grant              | 5      |
| Hiro Tanaka            | 5      |
| Ravi Patel             | 5      |
| Ravi Patel             | 5      |
| Elena Morales          | 5      |
| Elena Morales          | 5      |
| Diego Ramos            | 5      |
| Caleb Osei             | 5      |
| Caleb Osei             | 5      |
| Lucas Ford             | 5      |
| Yara Haddad            | 5      |
| Tariq Hassan           | 5      |
| Eva Schmidt            | 5      |
| Ingrid Nilsen          | 5      |
| Sophia Rossi           | 5      |
| Olivia Dubois          | 5      |
| Emma Svensson          | 5      |
| Isabella Laurent       | 5      |
| James Kim              | 5      |
...
...
| Griffin Hartwell       | 9      |
| Griffin Hartwell       | 9      |
| Megan Redfield         | 9      |
| Megan Redfield         | 9      |
| Grayson Westmore       | 9      |
| Nicole Ashridge        | 9      |
| Sawyer Hollingsworth   | 9      |
| Sawyer Hollingsworth   | 9      |
| Alexis Thorndale       | 9      |
| Declan Summerfield     | 9      |
| Declan Summerfield     | 9      |
| Samantha Brightwood    | 9      |
| Samantha Brightwood    | 9      |
| Tristan Ashbrook       | 9      |
| Lauren Silverton       | 9      |
| Landon Whitworth       | 9      |
| Kayla Mansfield        | 9      |
| Kayla Mansfield        | 9      |
+------------------------+--------+
sqlite>


sqlite> WITH stocked_cars as (
    SELECT * FROM cocoa_cars ORDER BY total_stock DESC LIMIT 3
)
SELECT passengers.passenger_name, passengers.favorite_mixins, stocked_cars.available_mixins, passenger_mixin.value as passenger_mixin, available_mixin.value as available_mixin, stocked_cars.car_id
FROM passengers
JOIN json_each(passengers.favorite_mixins) as passenger_mixin
JOIN stocked_cars
JOIN json_each(stocked_cars.available_mixins) AS available_mixin
    ON passenger_mixin.value = available_mixin.value;
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
|     passenger_name     |                       favorite_mixins                        |                       available_mixins                       | passenger_mixin  | available_mixin  | car_id |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Mateo Cruz             | ["caramel drizzle","shaved chocolate","white chocolate"]     | ["white chocolate","shaved chocolate"]                       | shaved chocolate | shaved chocolate | 5      |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Mateo Cruz             | ["caramel drizzle","shaved chocolate","white chocolate"]     | ["white chocolate","shaved chocolate"]                       | white chocolate  | white chocolate  | 5      |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Nia Grant              | ["shaved chocolate"]                                         | ["white chocolate","shaved chocolate"]                       | shaved chocolate | shaved chocolate | 5      |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Hiro Tanaka            | ["shaved chocolate"]                                         | ["white chocolate","shaved chocolate"]                       | shaved chocolate | shaved chocolate | 5      |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Ravi Patel             | ["caramel drizzle","shaved chocolate","white chocolate"]     | ["white chocolate","shaved chocolate"]                       | shaved chocolate | shaved chocolate | 5      |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Ravi Patel             | ["caramel drizzle","shaved chocolate","white chocolate"]     | ["white chocolate","shaved chocolate"]                       | white chocolate  | white chocolate  | 5      |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Elena Morales          | ["white chocolate","shaved chocolate","caramel drizzle"]     | ["white chocolate","shaved chocolate"]                       | white chocolate  | white chocolate  | 5      |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Elena Morales          | ["white chocolate","shaved chocolate","caramel drizzle"]     | ["white chocolate","shaved chocolate"]                       | shaved chocolate | shaved chocolate | 5      |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Diego Ramos            | ["shaved chocolate"]                                         | ["white chocolate","shaved chocolate"]                       | shaved chocolate | shaved chocolate | 5      |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Caleb Osei             | ["shaved chocolate","dark chocolate","white chocolate"]      | ["white chocolate","shaved chocolate"]                       | shaved chocolate | shaved chocolate | 5      |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Caleb Osei             | ["shaved chocolate","dark chocolate","white chocolate"]      | ["white chocolate","shaved chocolate"]                       | white chocolate  | white chocolate  | 5      |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Lucas Ford             | ["vanilla foam","white chocolate","cinnamon"]                | ["white chocolate","shaved chocolate"]                       | white chocolate  | white chocolate  | 5      |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Yara Haddad            | ["white chocolate","dark chocolate"]                         | ["white chocolate","shaved chocolate"]                       | white chocolate  | white chocolate  | 5      |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Tariq Hassan           | ["dark chocolate","crispy rice","white chocolate","peppermin | ["white chocolate","shaved chocolate"]                       | white chocolate  | white chocolate  | 5      |
|                        | t"]                                                          |                                                              |                  |                  |        |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Eva Schmidt            | ["white chocolate","marshmallow"]                            | ["white chocolate","shaved chocolate"]                       | white chocolate  | white chocolate  | 5      |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
...
...
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Sawyer Hollingsworth   | ["crispy rice","dark chocolate","peppermint"]                | ["crispy rice","peppermint","caramel drizzle","shaved chocol | crispy rice      | crispy rice      | 9      |
|                        |                                                              | ate"]                                                        |                  |                  |        |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Sawyer Hollingsworth   | ["crispy rice","dark chocolate","peppermint"]                | ["crispy rice","peppermint","caramel drizzle","shaved chocol | peppermint       | peppermint       | 9      |
|                        |                                                              | ate"]                                                        |                  |                  |        |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Alexis Thorndale       | ["vanilla foam","crispy rice","dark chocolate"]              | ["crispy rice","peppermint","caramel drizzle","shaved chocol | crispy rice      | crispy rice      | 9      |
|                        |                                                              | ate"]                                                        |                  |                  |        |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Declan Summerfield     | ["dark chocolate","marshmallow","crispy rice","peppermint"]  | ["crispy rice","peppermint","caramel drizzle","shaved chocol | crispy rice      | crispy rice      | 9      |
|                        |                                                              | ate"]                                                        |                  |                  |        |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Declan Summerfield     | ["dark chocolate","marshmallow","crispy rice","peppermint"]  | ["crispy rice","peppermint","caramel drizzle","shaved chocol | peppermint       | peppermint       | 9      |
|                        |                                                              | ate"]                                                        |                  |                  |        |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Samantha Brightwood    | ["crispy rice","peppermint","dark chocolate"]                | ["crispy rice","peppermint","caramel drizzle","shaved chocol | crispy rice      | crispy rice      | 9      |
|                        |                                                              | ate"]                                                        |                  |                  |        |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Samantha Brightwood    | ["crispy rice","peppermint","dark chocolate"]                | ["crispy rice","peppermint","caramel drizzle","shaved chocol | peppermint       | peppermint       | 9      |
|                        |                                                              | ate"]                                                        |                  |                  |        |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Tristan Ashbrook       | ["crispy rice"]                                              | ["crispy rice","peppermint","caramel drizzle","shaved chocol | crispy rice      | crispy rice      | 9      |
|                        |                                                              | ate"]                                                        |                  |                  |        |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Lauren Silverton       | ["caramel drizzle","vanilla foam","cinnamon","dark chocolate | ["crispy rice","peppermint","caramel drizzle","shaved chocol | caramel drizzle  | caramel drizzle  | 9      |
|                        | "]                                                           | ate"]                                                        |                  |                  |        |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Landon Whitworth       | ["caramel drizzle","vanilla foam"]                           | ["crispy rice","peppermint","caramel drizzle","shaved chocol | caramel drizzle  | caramel drizzle  | 9      |
|                        |                                                              | ate"]                                                        |                  |                  |        |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Kayla Mansfield        | ["vanilla foam","peppermint","shaved chocolate"]             | ["crispy rice","peppermint","caramel drizzle","shaved chocol | peppermint       | peppermint       | 9      |
|                        |                                                              | ate"]                                                        |                  |                  |        |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
| Kayla Mansfield        | ["vanilla foam","peppermint","shaved chocolate"]             | ["crispy rice","peppermint","caramel drizzle","shaved chocol | shaved chocolate | shaved chocolate | 9      |
|                        |                                                              | ate"]                                                        |                  |                  |        |
+------------------------+--------------------------------------------------------------+--------------------------------------------------------------+------------------+------------------+--------+
sqlite> 





sqlite> WITH stocked_cars as (
    SELECT * FROM cocoa_cars ORDER BY total_stock DESC LIMIT 3
)
SELECT *                                             
FROM passengers
JOIN json_each(passengers.favorite_mixins) as passenger_mixin
JOIN stocked_cars
JOIN json_each(stocked_cars.available_mixins) AS available_mixin
    ON passenger_mixin.value = available_mixin.value;
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| passenger_id |     passenger_name     |                       favorite_mixins                        | car_id | key |      value       | type |       atom       | id | parent | fullkey | path | car_id |                       available_mixins                       | total_stock | key |      value       | type |       atom       | id | parent | fullkey | path |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 2            | Mateo Cruz             | ["caramel drizzle","shaved chocolate","white chocolate"]     | 2      | 1   | shaved chocolate | text | shaved chocolate | 19 |        | $[1]    | $    | 5      | ["white chocolate","shaved chocolate"]                       | 412         | 1   | shaved chocolate | text | shaved chocolate | 19 |        | $[1]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 2            | Mateo Cruz             | ["caramel drizzle","shaved chocolate","white chocolate"]     | 2      | 2   | white chocolate  | text | white chocolate  | 37 |        | $[2]    | $    | 5      | ["white chocolate","shaved chocolate"]                       | 412         | 0   | white chocolate  | text | white chocolate  | 2  |        | $[0]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 3            | Nia Grant              | ["shaved chocolate"]                                         | 5      | 0   | shaved chocolate | text | shaved chocolate | 2  |        | $[0]    | $    | 5      | ["white chocolate","shaved chocolate"]                       | 412         | 1   | shaved chocolate | text | shaved chocolate | 19 |        | $[1]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 4            | Hiro Tanaka            | ["shaved chocolate"]                                         | 2      | 0   | shaved chocolate | text | shaved chocolate | 2  |        | $[0]    | $    | 5      | ["white chocolate","shaved chocolate"]                       | 412         | 1   | shaved chocolate | text | shaved chocolate | 19 |        | $[1]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 6            | Ravi Patel             | ["caramel drizzle","shaved chocolate","white chocolate"]     | 5      | 1   | shaved chocolate | text | shaved chocolate | 19 |        | $[1]    | $    | 5      | ["white chocolate","shaved chocolate"]                       | 412         | 1   | shaved chocolate | text | shaved chocolate | 19 |        | $[1]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 6            | Ravi Patel             | ["caramel drizzle","shaved chocolate","white chocolate"]     | 5      | 2   | white chocolate  | text | white chocolate  | 37 |        | $[2]    | $    | 5      | ["white chocolate","shaved chocolate"]                       | 412         | 0   | white chocolate  | text | white chocolate  | 2  |        | $[0]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 9            | Elena Morales          | ["white chocolate","shaved chocolate","caramel drizzle"]     | 6      | 0   | white chocolate  | text | white chocolate  | 2  |        | $[0]    | $    | 5      | ["white chocolate","shaved chocolate"]                       | 412         | 0   | white chocolate  | text | white chocolate  | 2  |        | $[0]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 9            | Elena Morales          | ["white chocolate","shaved chocolate","caramel drizzle"]     | 6      | 1   | shaved chocolate | text | shaved chocolate | 19 |        | $[1]    | $    | 5      | ["white chocolate","shaved chocolate"]                       | 412         | 1   | shaved chocolate | text | shaved chocolate | 19 |        | $[1]    | $    |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 10           | Diego Ramos            | ["shaved chocolate"]                                         | 1      | 0   | shaved chocolate | text | shaved chocolate | 2  |        | $[0]    | $    | 5      | ["white chocolate","shaved chocolate"]                       | 412         | 1   | shaved chocolate | text | shaved chocolate | 19 |        | $[1]    | $    |
...
...
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 171          | Sawyer Hollingsworth   | ["crispy rice","dark chocolate","peppermint"]                | 8      | 0   | crispy rice      | text | crispy rice      | 2  |        | $[0]    | $    | 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 0   | crispy rice      | text | crispy rice      | 2  |        | $[0]    | $    |
|              |                        |                                                              |        |     |                  |      |                  |    |        |         |      |        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 171          | Sawyer Hollingsworth   | ["crispy rice","dark chocolate","peppermint"]                | 8      | 2   | peppermint       | text | peppermint       | 30 |        | $[2]    | $    | 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 1   | peppermint       | text | peppermint       | 14 |        | $[1]    | $    |
|              |                        |                                                              |        |     |                  |      |                  |    |        |         |      |        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 172          | Alexis Thorndale       | ["vanilla foam","crispy rice","dark chocolate"]              | 9      | 1   | crispy rice      | text | crispy rice      | 16 |        | $[1]    | $    | 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 0   | crispy rice      | text | crispy rice      | 2  |        | $[0]    | $    |
|              |                        |                                                              |        |     |                  |      |                  |    |        |         |      |        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 173          | Declan Summerfield     | ["dark chocolate","marshmallow","crispy rice","peppermint"]  | 7      | 2   | crispy rice      | text | crispy rice      | 30 |        | $[2]    | $    | 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 0   | crispy rice      | text | crispy rice      | 2  |        | $[0]    | $    |
|              |                        |                                                              |        |     |                  |      |                  |    |        |         |      |        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 173          | Declan Summerfield     | ["dark chocolate","marshmallow","crispy rice","peppermint"]  | 7      | 3   | peppermint       | text | peppermint       | 42 |        | $[3]    | $    | 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 1   | peppermint       | text | peppermint       | 14 |        | $[1]    | $    |
|              |                        |                                                              |        |     |                  |      |                  |    |        |         |      |        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 174          | Samantha Brightwood    | ["crispy rice","peppermint","dark chocolate"]                | 4      | 0   | crispy rice      | text | crispy rice      | 2  |        | $[0]    | $    | 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 0   | crispy rice      | text | crispy rice      | 2  |        | $[0]    | $    |
|              |                        |                                                              |        |     |                  |      |                  |    |        |         |      |        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 174          | Samantha Brightwood    | ["crispy rice","peppermint","dark chocolate"]                | 4      | 1   | peppermint       | text | peppermint       | 14 |        | $[1]    | $    | 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 1   | peppermint       | text | peppermint       | 14 |        | $[1]    | $    |
|              |                        |                                                              |        |     |                  |      |                  |    |        |         |      |        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 175          | Tristan Ashbrook       | ["crispy rice"]                                              | 1      | 0   | crispy rice      | text | crispy rice      | 2  |        | $[0]    | $    | 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 0   | crispy rice      | text | crispy rice      | 2  |        | $[0]    | $    |
|              |                        |                                                              |        |     |                  |      |                  |    |        |         |      |        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 178          | Lauren Silverton       | ["caramel drizzle","vanilla foam","cinnamon","dark chocolate | 2      | 0   | caramel drizzle  | text | caramel drizzle  | 2  |        | $[0]    | $    | 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 2   | caramel drizzle  | text | caramel drizzle  | 25 |        | $[2]    | $    |
|              |                        | "]                                                           |        |     |                  |      |                  |    |        |         |      |        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 179          | Landon Whitworth       | ["caramel drizzle","vanilla foam"]                           | 8      | 0   | caramel drizzle  | text | caramel drizzle  | 2  |        | $[0]    | $    | 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 2   | caramel drizzle  | text | caramel drizzle  | 25 |        | $[2]    | $    |
|              |                        |                                                              |        |     |                  |      |                  |    |        |         |      |        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 180          | Kayla Mansfield        | ["vanilla foam","peppermint","shaved chocolate"]             | 2      | 1   | peppermint       | text | peppermint       | 16 |        | $[1]    | $    | 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 1   | peppermint       | text | peppermint       | 14 |        | $[1]    | $    |
|              |                        |                                                              |        |     |                  |      |                  |    |        |         |      |        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
| 180          | Kayla Mansfield        | ["vanilla foam","peppermint","shaved chocolate"]             | 2      | 2   | shaved chocolate | text | shaved chocolate | 27 |        | $[2]    | $    | 9      | ["crispy rice","peppermint","caramel drizzle","shaved chocol | 354         | 3   | shaved chocolate | text | shaved chocolate | 42 |        | $[3]    | $    |
|              |                        |                                                              |        |     |                  |      |                  |    |        |         |      |        | ate"]                                                        |             |     |                  |      |                  |    |        |         |      |
+--------------+------------------------+--------------------------------------------------------------+--------+-----+------------------+------+------------------+----+--------+---------+------+--------+--------------------------------------------------------------+-------------+-----+------------------+------+------------------+----+--------+---------+------+
sqlite> 
```

Now what we did here?
- Create the top 3 stocked cars as CTE
- The main query first fetches each passenger with its favorite mixin expanded 
- It joined the table of available_mixin when?
- The mixin from passenger is equal to the mixin in the available mixin in the car
- Hence we get the car_id as the mixin served for that passenger.

You can see we have multiple rows for each user, we might not want that. Though nothing is wrong with it, but the report looks quite long, especially if there were more than a couple of mixins for the passenger or the number of cars were more.

We can group by the `passenger_name` and sort of concatenate the 

```sql
WITH stocked_cars as (
    SELECT * FROM cocoa_cars ORDER BY total_stock DESC LIMIT 3
)
SELECT 
     passengers.passenger_name,
     '[' || GROUP_CONCAT(DISTINCT stocked_cars.car_id) || ']' AS cocoa_cars
FROM passengers
JOIN json_each(passengers.favorite_mixins) as passenger_mixin
JOIN stocked_cars
JOIN json_each(stocked_cars.available_mixins) AS available_mixin
    ON passenger_mixin.value = available_mixin.value
GROUP BY passengers.passenger_name;
```

Here, we have added,

```
 '[' || GROUP_CONCAT(DISTINCT stocked_cars.car_id) || ']' 
```

AND to group by 

```
GROUP BY passengers.passenger_name;
```

So this will squish down all the separate rows of the passenger favorite mixins and the cars that have it, with a single row and we defined how we want to squish the different `car_id`s with a [GROUP_CONCAT](https://www.sqlite.org/lang_aggfunc.html#group_concat). This function can concatentate (join or attach or combine) together the multiple strings with a specific separator (by default the separate is `,`).
We also use the `||` concatenation operator to add `[` at the start and `]` at the end of the list of `car_ids`.  

> The group_concat() function returns a string which is the concatenation of all non-NULL values of X. If parameter Y is present then it is used as the separator between instances of X. A comma (",") is used as the separator if Y is omitted.



```
sqlite> WITH stocked_cars as (
    SELECT * FROM cocoa_cars ORDER BY total_stock DESC LIMIT 3
)
SELECT passengers.passenger_name, '[' || GROUP_CONCAT(DISTINCT stocked_cars.car_id) || ']' AS cocoa_cars
FROM passengers
JOIN json_each(passengers.favorite_mixins) as passenger_mixin
JOIN stocked_cars
JOIN json_each(stocked_cars.available_mixins) AS available_mixin
    ON passenger_mixin.value = available_mixin.value
   ...> GROUP BY passengers.passenger_name;
+------------------------+------------+
|     passenger_name     | cocoa_cars |
+------------------------+------------+
| Aaron Ashcroft         | [2]        |
| Abigail Worthington    | [2,9]      |
| Aiden Stanfield        | [9]        |
| Alexander Stratton     | [5,9]      |
| Alexis Thorndale       | [9]        |
| Alice Merriweather     | [5,2,9]    |
| Amanda Fitzroy         | [5,2,9]    |
| Amelia Rosewood        | [9]        |
| Anna Westfield         | [2]        |
| Aria Blackwood         | [5,2,9]    |
| Ashley Fairhaven       | [5,2,9]    |
| August Blackwell       | [5,2,9]    |
| Aurora Whitfield       | [5,2,9]    |
| Benjamin Fairweather   | [9]        |
| Benjamin Patel         | [9]        |
| Bianca Pereira         | [9]        |
| Blake Dunwood          | [5,9]      |
| Brandon Locklear       | [5,2,9]    |
| Caleb Osei             | [5,9]      |
| Cameron Gladstone      | [2,9]      |
| Cameron Rogers         | [5,2,9]    |
| Carlos Mendez          | [5,9]      |
| Caroline Bannister     | [2,9]      |
| Carter Brookfield      | [2]        |
| Charlotte Singh        | [5,2,9]    |
| Charlotte Waverly      | [5,9]      |
| Chase Ashbury          | [5,9]      |
| Chloe Blackstone       | [5,9]      |
| Claire Berkshire       | [5,9]      |
| Clara Westbrook        | [2]        |
| Cole Ashland           | [2]        |
| Connor Redmond         | [2,9]      |
| Daniel Murphy          | [5,2]      |
| Declan Summerfield     | [2,9]      |
| Diana Brightwell       | [5,9]      |
| Diego Ramos            | [5,9]      |
| Dylan Blakely          | [5,2,9]    |
| Eleanor Morrison       | [5,9]      |
| Elena Morales          | [5,2,9]    |
| Elijah Wakefield       | [2,9]      |
| Eliza Radcliffe        | [5,9]      |
| Elizabeth Chesterfield | [5,2,9]    |
| Emily Claridge         | [5,9]      |
| Emily Johnson          | [2]        |
| Emma Gilmore           | [5,2,9]    |
| Emma Svensson          | [5,9]      |
| Ethan Marlowe          | [5,2,9]    |
| Eva Schmidt            | [5,2]      |
| Eva Templeton          | [5,2,9]    |
| Fatima Noor            | [9]        |
| Felix Schneider        | [9]        |
| Felix Whitmore         | [2]        |
| Finn Lockhart          | [2,9]      |
| Gabriel Winthrop       | [5,2,9]    |
| Grace Aldridge         | [5,2,9]    |
| Grace Thornhill        | [5,2,9]    |
| Grayson Westmore       | [2,9]      |
| Griffin Hartwell       | [9]        |
| Hannah Livingston      | [5,2,9]    |
| Hazel Kincaid          | [2,9]      |
| Henry Treadwell        | [2,9]      |
| Hiro Tanaka            | [5,9]      |
| Hunter Bellingham      | [5,9]      |
| Ian Landsman           | [2]        |
| Ingrid Nilsen          | [5,2,9]    |
| Iris Pembroke          | [5,2,9]    |
| Isaac Kendrick         | [2]        |
| Isabella Laurent       | [5,9]      |
| Isabella Rodriguez     | [9]        |
| Ivy Winslow            | [2,9]      |
| Jack Carmichael        | [9]        |
| Jackson Thorpe         | [2,9]      |
| James Garrison         | [5,9]      |
| James Kim              | [5,2,9]    |
| James Simon            | [5,2,9]    |
| Jason Stewart          | [5]        |
| Jasper Thorne          | [5,2,9]    |
| Jessica Langston       | [5,2,9]    |
| Jonah Wolfe            | [2]        |
| Jordan Waverly         | [5,9]      |
| Julia Pendleton        | [5,2,9]    |
| Julian Blackburn       | [5]        |
| Kayla Mansfield        | [5,9]      |
| Keiko Ito              | [2]        |
| Landon Whitworth       | [2,9]      |
| Laura Thornbury        | [5,2,9]    |
| Lauren Silverton       | [2,9]      |
| Layla Brooks           | [2,9]      |
| Leo Greyson            | [2,9]      |
| Liam Donovan           | [2,9]      |
| Liam OConnor           | [2,9]      |
| Linda Lee              | [5,2]      |
| Logan Sherwood         | [5,9]      |
| Lucas Ford             | [5,2]      |
| Lucas Prescott         | [9]        |
| Lucy Morris            | [2,9]      |
| Luna Hartley           | [5]        |
| Madeline Ramsey        | [2,9]      |
| Madison Clearwater     | [5,2,9]    |
| Margaret Fairbanks     | [2,9]      |
| Mason Fairmont         | [2,9]      |
| Mateo Cruz             | [5,2,9]    |
| Matthew Sutherland     | [2,9]      |
| Megan Redfield         | [5,9]      |
| Melissa Ravenscroft    | [5]        |
| Mia Chen               | [5,9]      |
| Mila Novak             | [2,9]      |
| Miles Brennan          | [5,2,9]    |
| Milo Ashford           | [2,9]      |
| Mira Zhao              | [2,9]      |
| Morgan Steadman        | [5,9]      |
| Natalie Warwick        | [2]        |
| Nathan Whitley         | [5,9]      |
| Nia Grant              | [5,9]      |
| Nicholas Montague      | [5]        |
| Nicole Ashridge        | [9]        |
| Noah Fischer           | [2,9]      |
| Nolan Murphy           | [2,9]      |
| Nolan Young            | [5,2,9]    |
| Nora Calloway          | [9]        |
| Nova Adams             | [2]        |
| Oliver Ashby           | [9]        |
| Olivia Carrington      | [2]        |
| Olivia Dubois          | [5,9]      |
| Omar Qureshi           | [2]        |
| Owen Fairchild         | [5,9]      |
| Parker Blackmore       | [5,2,9]    |
| Penelope Sinclair      | [5,2,9]    |
| Rachel Wyndham         | [5,9]      |
| Rafael Silva           | [2,9]      |
| Ravi Patel             | [5,2,9]    |
| Rebecca Ashcombe       | [2,9]      |
| Robert Smith           | [9]        |
| Rose Drummond          | [5,2,9]    |
| Ruby Hawthorne         | [2]        |
| Ryan Beckett           | [5,9]      |
| Samantha Brightwood    | [9]        |
| Samuel Kingsley        | [5,2,9]    |
| Sara Johansson         | [2]        |
| Sarah Davis            | [9]        |
| Sarah Remington        | [5,2,9]    |
| Sawyer Hollingsworth   | [9]        |
| Scarlett Dalton        | [5,2,9]    |
| Silas Merrick          | [2]        |
| Sofia Kim              | [2]        |
| Sophia Brookshire      | [2,9]      |
| Sophia Rossi           | [5,2,9]    |
| Sophie Langford        | [5,9]      |
| Stella Beaumont        | [2,9]      |
| Tariq Hassan           | [5,9]      |
| Taylor Woodridge       | [5,2,9]    |
| Theodore Hadley        | [2]        |
| Tim Cook               | [9]        |
| Tristan Ashbrook       | [9]        |
| Victoria Harwood       | [2,9]      |
| Violet Sterling        | [2,9]      |
| William Beauchamp      | [5,2,9]    |
| William Becker         | [9]        |
| William Chen           | [5,2,9]    |
| Wyatt Alderton         | [5]        |
| Yara Haddad            | [5]        |
| Youssef El-Sayed       | [2,9]      |
| Zachary Collins        | [2,9]      |
| Zachary Stafford       | [5,9]      |
| Zara Sheikh            | [9]        |
| Zoe Wilson             | [2]        |
+------------------------+------------+
sqlite> 
```

I think we are done.

We did it, it was a bit different one.

Some wired hack here and there but we made it!

Day 7 done and dusted.

On to the day 8.
