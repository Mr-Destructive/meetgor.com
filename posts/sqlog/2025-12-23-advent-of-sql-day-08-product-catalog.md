---
type: "sqlog"
title: "Advent of SQL 2025 Day 8: Product Catalog"
slug: "advent-of-sql-2025-day-8"
date: 2025-12-23 15:30 +0530
series: ["advent-of-sql-2025"]
tags: ["sqlite", "sql", "advent-of-sql"]
---

## Advent of SQL - Day 8, Product Catalog

Whopsies! This is day 8.

Let's get straigh...

HOOH! We need to clean up some SQL for running in SQLite.

```bash
sed -i 's/TIMESTAMP[[:space:]]*//g' day8-inserts-sqlite.sql
```

Just cleaning up `TIMESTAMP` in `INSERT` before the date value.

Here we go:
The SQL to run in SQLite.

```sql
DROP TABLE IF EXISTS price_changes;
DROP TABLE IF EXISTS products;

CREATE TABLE products (
    product_id INT PRIMARY KEY,
    product_name TEXT
);

CREATE TABLE price_changes (
    id INT PRIMARY KEY,
    product_id INT,
    price NUMERIC(10,2),
    effective_timestamp 
);

INSERT INTO products (product_id, product_name) VALUES
    (1, 'Deluxe Sled'),
    (2, 'Holiday Trail Mix Trio'),
    (3, 'Premium Cinnamon Roasted Almonds'),
    (4, 'Deluxe Wrapping Paper'),
    (5, 'Deluxe Roasted Cashews'),
    (6, 'Festive Cookware Set'),
    (7, 'Deluxe Mug'),
    (8, 'Premium Sled'),
    (9, 'Essential Sled'),
    (10, 'Family Snow Boots'),
    (11, 'Family Dark Chocolate Almonds'),
    (12, 'Premium Festive Scarf'),
    (13, 'Essential Cookie Decorating Kit'),
    (14, 'Festive White Chocolate Popcorn'),
    (15, 'Cozy Puzzle'),
    (16, 'Holiday Cheddar Popcorn'),
    (17, 'Premium Board Game'),
    (18, 'Deluxe Pecan Praline Bites'),
    (19, 'Cozy Almond Brittle'),
    (20, 'Winter Sled');

INSERT INTO price_changes (id, product_id, price, effective_timestamp) VALUES
    (1, 1, 148.28, '2025-12-01 05:25:35'),
    (2, 1, 148.63, '2025-12-02 18:15:33'),
    (3, 1, 126.78, '2025-12-02 18:40:38'),
    (4, 1, 119.12, '2025-12-03 10:14:35'),
    (5, 1, 98.57, '2025-12-04 04:14:31'),
    (6, 1, 88.49, '2025-12-06 19:02:40'),
    (7, 1, 80.88, '2025-12-07 10:43:54'),
    (8, 1, 78.88, '2025-12-08 06:45:39'),
    (9, 1, 80.24, '2025-12-08 16:11:11'),
    (10, 1, 73.9, '2025-12-10 14:33:43'),
    (11, 1, 88.2, '2025-12-12 02:21:09'),
    (12, 1, 99.03, '2025-12-12 02:58:14'),
    (13, 1, 100.18, '2025-12-14 15:58:03'),
    (14, 1, 106.91, '2025-12-16 01:51:05'),
    (15, 1, 109.25, '2025-12-16 16:01:53'),
    (16, 2, 29.54, '2025-12-03 14:21:10'),
    (17, 2, 34.33, '2025-12-03 19:14:31'),
    (18, 2, 39.08, '2025-12-04 06:13:48'),
    (19, 2, 32.71, '2025-12-04 18:33:17'),
    (20, 2, 31.71, '2025-12-05 22:36:14'),
    (21, 2, 28.88, '2025-12-06 02:42:02'),
    (22, 2, 23.14, '2025-12-07 09:46:54'),
    (23, 2, 25.65, '2025-12-07 10:03:45'),
    (24, 2, 27.06, '2025-12-07 14:39:15'),
    (25, 2, 24.48, '2025-12-07 20:08:05'),
    (26, 2, 26.84, '2025-12-09 07:44:32'),
    (27, 2, 27.39, '2025-12-13 06:25:19'),
    (28, 2, 26.6, '2025-12-14 10:16:34'),
    (29, 2, 21.38, '2025-12-15 16:20:20'),
    (30, 2, 17.75, '2025-12-16 09:28:13');
```

We can get started.

## Problem

> Generate a report, using the products and `price_changes` tables for leadership that returns the `product_name`, `current_price`, `previous_price`, and the difference between the current and previous prices


So what we need is 

- product_name
- current_price (latest)
- previous_price (just before the latest)
- price_difference = current - previous

Again we have to meddle with dates, maybe, maybe not!

### With CTEs and JOINs

Let's start with the simplest approach. We need 2 prices, the latest(highest date timestamp) and the 2nd latest (the 2nd highest date timestamp). We can get the first pretty easily, but what about the second?

Well, if we get the first, then the second should be easy to get right? Right? Because it will be just before it. Well not directly.

Let's first grab the max timestamp.

```sql
SELECT 
    product_id,
    MAX(effective_timestamp) AS latest_timestamp
FROM price_changes
GROUP BY product_id;
```

```
sqlite> SELECT product_id, MAX(effective_timestamp) AS max_ts
        FROM price_changes
        GROUP BY product_id;
+------------+---------------------+
| product_id |       max_ts        |
+------------+---------------------+
| 1          | 2025-12-16 16:01:53 |
| 2          | 2025-12-16 09:28:13 |
| 3          | 2025-12-15 03:20:11 |
| 4          | 2025-12-16 01:33:41 |
| 5          | 2025-12-12 10:11:48 |
| 6          | 2025-12-15 11:31:40 |
| 7          | 2025-12-16 03:00:51 |
| 8          | 2025-12-15 22:33:48 |
| 9          | 2025-12-15 20:05:34 |
| 10         | 2025-12-15 20:53:45 |
...
...
| 139        | 2025-12-16 04:46:33 |
| 140        | 2025-12-16 21:19:30 |
| 141        | 2025-12-15 09:50:36 |
| 142        | 2025-12-16 18:39:51 |
| 143        | 2025-12-15 07:27:06 |
| 144        | 2025-12-16 16:25:16 |
| 146        | 2025-12-13 07:07:19 |
| 148        | 2025-12-16 09:30:11 |
| 149        | 2025-12-13 16:40:21 |
| 150        | 2025-12-13 08:24:43 |
+------------+---------------------+
```

We got all the timestamp's for each product. But we wanted the prices. Well we can't grab the price here, since we are grouping!

We can use the timestamp as we are selecting the `MAX` aggregate function however among the many rows for single product how do we group the price? `MIN`, `MAX`, `AVG`, but that's not we want, we just want the price for that timestamp.

Well, we need to join to the same table for that timestamp and grab the price.

```sql
SELECT 
    price_changes.product_id,
    price_changes.price AS current_price,
    price_changes.effective_timestamp AS latest_ts
FROM price_changes
JOIN (
    SELECT 
        product_id, 
        MAX(effective_timestamp) AS latest_timestamp
    FROM price_changes
    GROUP BY product_id
) AS latest_price_change
  ON price_changes.product_id = latest_price_change.product_id
 AND price_changes.effective_timestamp = latest_price_change.latest_timestamp;
```

Here, we first specify what we want
- `product_id`
- `current_price` which is the price for the 
- `latest_timestamp` which is the latest time recorded for the product price.

We are grouping this by `product_id` since there are price recorded and various timestamps. We need to get the latest timestamp.
So, we do a nested query to join the latest timestamp and join the price with the same timestamp.
 
This condition `price_changes.effective_timestamp = latest_price_change.latest_timestamp` helps us get the `price` for the `latest_timestamp`. We first get each timestamp for the product and then find its max, so that's the inner query from which we joined this table to. Self join with the different thing to find.

This gives us 2 things.
- Product id
- Price for the latest timestamp 
 
We don't really want the timestamp in the final result, its just a criteria or a intermediate value to get the current and the previous prices for each product.

```
sqlite> 
sqlite> SELECT 
    price_changes.product_id,
    price_changes.price AS current_price,
    price_changes.effective_timestamp AS latest_ts
FROM price_changes
JOIN (
    SELECT 
        product_id, 
        MAX(effective_timestamp) AS latest_timestamp
    FROM price_changes
    GROUP BY product_id
) AS latest_price_change
  ON price_changes.product_id = latest_price_change.product_id
 AND price_changes.effective_timestamp = latest_price_change.latest_timestamp;
+------------+---------------+---------------------+
| product_id | current_price |      latest_ts      |
+------------+---------------+---------------------+
| 1          | 109.25        | 2025-12-16 16:01:53 |
| 2          | 17.75         | 2025-12-16 09:28:13 |
| 3          | 143.65        | 2025-12-15 03:20:11 |
| 4          | 98.51         | 2025-12-16 01:33:41 |
| 5          | 124.04        | 2025-12-12 10:11:48 |
| 6          | 84.14         | 2025-12-15 11:31:40 |
| 7          | 123.09        | 2025-12-16 03:00:51 |
| 8          | 221.06        | 2025-12-15 22:33:48 |
| 9          | 255.88        | 2025-12-15 20:05:34 |
| 10         | 57.99         | 2025-12-15 20:53:45 |
...
...
| 139        | 16.41         | 2025-12-16 04:46:33 |
| 140        | 173.05        | 2025-12-16 21:19:30 |
| 141        | 69.97         | 2025-12-15 09:50:36 |
| 142        | 35.05         | 2025-12-16 18:39:51 |
| 143        | 153.94        | 2025-12-15 07:27:06 |
| 144        | 118.21        | 2025-12-16 16:25:16 |
| 146        | 54.73         | 2025-12-13 07:07:19 |
| 148        | 107.81        | 2025-12-16 09:30:11 |
| 149        | 72.6          | 2025-12-13 16:40:21 |
| 150        | 138.66        | 2025-12-13 08:24:43 |
+------------+---------------+---------------------+
sqlite> 

```

Now, we need to the price before this max timestamp. How do we get it?

We need to again join? Yes...

We need to subquery the timestamp just before it. But how will we get the max timestamp for each product. Well that's what we wrote above.

We can convert that to a CTE.

```sql

WITH latest AS (
    SELECT price_changes.product_id,
           price_changes.price AS current_price,
           price_changes.effective_timestamp AS latest_timestamp
    FROM price_changes
    JOIN (                                                   
        SELECT product_id, MAX(effective_timestamp) AS max_timestamp
        FROM price_changes
        GROUP BY product_id
    ) latest_price_changes                             
      ON price_changes.product_id = latest_price_changes.product_id
     AND price_changes.effective_timestamp = latest_price_changes.max_timestamp
)
SELECT * FROM latest;
```

We just got the same thing, however we can now use `latest` as a temporary table in the query.

```sql
WITH latest AS (
    SELECT price_changes.product_id,
           price_changes.price AS current_price,
           price_changes.effective_timestamp AS latest_timestamp
    FROM price_changes
    JOIN (                                                   
        SELECT product_id, MAX(effective_timestamp) AS max_timestamp
        FROM price_changes
        GROUP BY product_id
    ) latest_price_changes
      ON price_changes.product_id = latest_price_changes.product_id
     AND price_changes.effective_timestamp = latest_price_changes.max_timestamp
),  previous AS (
    SELECT price_changes.product_id,
           price_changes.price AS previous_price,
           MAX(price_changes.effective_timestamp) AS previous_timestamp
    FROM price_changes
    JOIN latest
      ON price_changes.product_id = latest.product_id
    WHERE price_changes.effective_timestamp < latest.latest_timestamp
    GROUP BY price_changes.product_id
)
SELECT * FROM previous JOIN latest ON previous.product_id = latest.product_id;
```


So, ok, this is getting long.

Just we added this 

```sql
SELECT 
    price_changes.product_id,
    price_changes.price AS previous_price,
    MAX(price_changes.effective_timestamp) AS previous_timestamp
FROM price_changes
JOIN latest
    ON price_changes.product_id = latest.product_id
WHERE 
    price_changes.effective_timestamp < latest.latest_timestamp
GROUP BY price_changes.product_id
```

This won't work as we need the latest table which we converted to CTE.

So, to get the 2nd highest or latest timestamp for each product, we do this.

```sql
WITH latest AS (
    SELECT price_changes.product_id,
           price_changes.price AS current_price,
           price_changes.effective_timestamp AS latest_timestamp
    FROM price_changes
    JOIN (                                                   
        SELECT product_id, MAX(effective_timestamp) AS max_timestamp
        FROM price_changes
        GROUP BY product_id
    ) latest_price_changes                             
      ON price_changes.product_id = latest_price_changes.product_id
     AND price_changes.effective_timestamp = latest_price_changes.max_timestamp
) 
    SELECT price_changes.product_id,
           price_changes.price AS previous_price,
           MAX(price_changes.effective_timestamp) AS previous_timestamp
    FROM price_changes
    JOIN latest
      ON price_changes.product_id = latest.product_id
    WHERE price_changes.effective_timestamp < latest.latest_timestamp
    GROUP BY price_changes.product_id;
```

We use:
 
```sql
 JOIN latest
      ON price_changes.product_id = latest.product_id
    WHERE price_changes.effective_timestamp < latest.latest_timestamp
```
To get match the timestamp which will exclude the latest timestamp and that way we can again get the `MAX(price_changes.effective_timestamp) AS previous_timestamp` for the subset of the timestamp.



This gives us all the previous timestamps i.e. the price just before the max timestamp for each product.

```
sqlite> WITH latest AS (
    SELECT price_changes.product_id,
           price_changes.price AS current_price,
           price_changes.effective_timestamp AS latest_timestamp
    FROM price_changes
    JOIN (                                                   
        SELECT product_id, MAX(effective_timestamp) AS max_timestamp
        FROM price_changes
        GROUP BY product_id
    ) latest_price_changes                             
      ON price_changes.product_id = latest_price_changes.product_id
     AND price_changes.effective_timestamp = latest_price_changes.max_timestamp
) 
    SELECT price_changes.product_id,
           price_changes.price AS previous_price,
           MAX(price_changes.effective_timestamp) AS previous_timestamp
    FROM price_changes
    JOIN latest
      ON price_changes.product_id = latest.product_id
    WHERE price_changes.effective_timestamp < latest.latest_timestamp
    GROUP BY price_changes.product_id;
+------------+----------------+---------------------+
| product_id | previous_price | previous_timestamp  |
+------------+----------------+---------------------+
| 1          | 106.91         | 2025-12-16 01:51:05 |
| 2          | 21.38          | 2025-12-15 16:20:20 |
| 3          | 159.65         | 2025-12-14 09:52:09 |
| 4          | 105.6          | 2025-12-14 14:09:20 |
| 5          | 129.23         | 2025-12-12 04:08:45 |
| 6          | 88.97          | 2025-12-15 02:13:04 |
| 7          | 127.14         | 2025-12-13 07:25:12 |
| 8          | 241.99         | 2025-12-14 19:31:40 |
| 9          | 259.56         | 2025-12-12 12:47:13 |
| 10         | 64.88          | 2025-12-15 10:52:34 |
...
...
| 140        | 157.02         | 2025-12-09 17:58:07 |
| 141        | 73.88          | 2025-12-13 14:35:53 |
| 142        | 30.25          | 2025-12-16 13:44:42 |
| 143        | 143.04         | 2025-12-11 10:08:13 |
| 144        | 114.22         | 2025-12-15 17:33:37 |
| 146        | 65.71          | 2025-12-10 15:50:14 |
| 148        | 101.09         | 2025-12-15 05:37:27 |
| 149        | 86.31          | 2025-12-13 13:18:14 |
| 150        | 123.61         | 2025-12-12 10:49:22 |
+------------+----------------+---------------------+
sqlite> 
```

Now, we need to join both of these on the same product id, to fetch both previous and current timestamp as well as the price.

```sql
SELECT * FROM previous JOIN latest ON previous.product_id = latest.product_id;
```

Simple

```sql
WITH latest AS (
    SELECT price_changes.product_id,
           price_changes.price AS current_price,
           price_changes.effective_timestamp AS latest_timestamp
    FROM price_changes
    JOIN (                                                   
        SELECT product_id, MAX(effective_timestamp) AS max_timestamp
        FROM price_changes
        GROUP BY product_id
    ) latest_price_changes                             
      ON price_changes.product_id = latest_price_changes.product_id
     AND price_changes.effective_timestamp = latest_price_changes.max_timestamp
),  previous AS (
    SELECT price_changes.product_id,
           price_changes.price AS previous_price,
           MAX(price_changes.effective_timestamp) AS previous_timestamp
    FROM price_changes
    JOIN latest
      ON price_changes.product_id = latest.product_id
    WHERE price_changes.effective_timestamp < latest.latest_timestamp
    GROUP BY price_changes.product_id 
)
SELECT * FROM previous JOIN latest ON previous.product_id = latest.product_id;
```

```
sqlite> WITH latest AS (
    SELECT price_changes.product_id,
           price_changes.price AS current_price,
           price_changes.effective_timestamp AS latest_timestamp
    FROM price_changes
    JOIN (                                                   
        SELECT product_id, MAX(effective_timestamp) AS max_timestamp
        FROM price_changes
        GROUP BY product_id
    ) latest_price_changes                             
      ON price_changes.product_id = latest_price_changes.product_id
     AND price_changes.effective_timestamp = latest_price_changes.max_timestamp
),  previous AS (
    SELECT price_changes.product_id,
           price_changes.price AS previous_price,
           MAX(price_changes.effective_timestamp) AS previous_timestamp
    FROM price_changes
    JOIN latest
      ON price_changes.product_id = latest.product_id
    WHERE price_changes.effective_timestamp < latest.latest_timestamp
    GROUP BY price_changes.product_id
)
SELECT * FROM previous JOIN latest ON previous.product_id = latest.product_id;
+------------+----------------+---------------------+------------+---------------+---------------------+
| product_id | previous_price | previous_timestamp  | product_id | current_price |  latest_timestamp   |
+------------+----------------+---------------------+------------+---------------+---------------------+
| 1          | 106.91         | 2025-12-16 01:51:05 | 1          | 109.25        | 2025-12-16 16:01:53 |
| 2          | 21.38          | 2025-12-15 16:20:20 | 2          | 17.75         | 2025-12-16 09:28:13 |
| 3          | 159.65         | 2025-12-14 09:52:09 | 3          | 143.65        | 2025-12-15 03:20:11 |
| 4          | 105.6          | 2025-12-14 14:09:20 | 4          | 98.51         | 2025-12-16 01:33:41 |
| 5          | 129.23         | 2025-12-12 04:08:45 | 5          | 124.04        | 2025-12-12 10:11:48 |
| 6          | 88.97          | 2025-12-15 02:13:04 | 6          | 84.14         | 2025-12-15 11:31:40 |
| 7          | 127.14         | 2025-12-13 07:25:12 | 7          | 123.09        | 2025-12-16 03:00:51 |
| 8          | 241.99         | 2025-12-14 19:31:40 | 8          | 221.06        | 2025-12-15 22:33:48 |
| 9          | 259.56         | 2025-12-12 12:47:13 | 9          | 255.88        | 2025-12-15 20:05:34 |
| 10         | 64.88          | 2025-12-15 10:52:34 | 10         | 57.99         | 2025-12-15 20:53:45 |
...
...
| 139        | 15.26          | 2025-12-12 03:43:32 | 139        | 16.41         | 2025-12-16 04:46:33 |
| 140        | 157.02         | 2025-12-09 17:58:07 | 140        | 173.05        | 2025-12-16 21:19:30 |
| 141        | 73.88          | 2025-12-13 14:35:53 | 141        | 69.97         | 2025-12-15 09:50:36 |
| 142        | 30.25          | 2025-12-16 13:44:42 | 142        | 35.05         | 2025-12-16 18:39:51 |
| 143        | 143.04         | 2025-12-11 10:08:13 | 143        | 153.94        | 2025-12-15 07:27:06 |
| 144        | 114.22         | 2025-12-15 17:33:37 | 144        | 118.21        | 2025-12-16 16:25:16 |
| 146        | 65.71          | 2025-12-10 15:50:14 | 146        | 54.73         | 2025-12-13 07:07:19 |
| 148        | 101.09         | 2025-12-15 05:37:27 | 148        | 107.81        | 2025-12-16 09:30:11 |
| 149        | 86.31          | 2025-12-13 13:18:14 | 149        | 72.6          | 2025-12-13 16:40:21 |
| 150        | 123.61         | 2025-12-12 10:49:22 | 150        | 138.66        | 2025-12-13 08:24:43 |
+------------+----------------+---------------------+------------+---------------+---------------------+
sqlite> 


```

Now, we are getting somewhere, we just need to find the difference right?

Yes, but more JOINs

We need the product_name from the `product` table. Almost forgot that table exists?

```sql
SELECT 
    *
FROM 
    products 
JOIN latest 
    ON products.product_id = latest.product_id 
LEFT JOIN previous 
    ON products.product_id = previous.product_id;
```

We need both right? So, we need to fetch the product id and join for the latest and the previous table from the CTE, joining on the `product_id`.


```sql
WITH latest AS (                                        
    SELECT price_changes.product_id,
           price_changes.price AS current_price,
           price_changes.effective_timestamp AS latest_timestamp
    FROM price_changes
    JOIN (                                                   
        SELECT product_id, MAX(effective_timestamp) AS max_timestamp
        FROM price_changes
        GROUP BY product_id
    ) latest_price_changes                             
      ON price_changes.product_id = latest_price_changes.product_id
     AND price_changes.effective_timestamp = latest_price_changes.max_timestamp
),  previous AS (
    SELECT price_changes.product_id,
           price_changes.price AS previous_price,
           MAX(price_changes.effective_timestamp) AS previous_timestamp
    FROM price_changes
    JOIN latest
      ON price_changes.product_id = latest.product_id
    WHERE price_changes.effective_timestamp < latest.latest_timestamp
    GROUP BY price_changes.product_id
)
SELECT 
    *
FROM 
    products 
JOIN latest 
    ON products.product_id = latest.product_id 
LEFT JOIN previous 
    ON products.product_id = previous.product_id;
```


```
sqlite> WITH latest AS (                                        
    SELECT price_changes.product_id,
           price_changes.price AS current_price,
           price_changes.effective_timestamp AS latest_timestamp
    FROM price_changes
    JOIN (                                                   
        SELECT product_id, MAX(effective_timestamp) AS max_timestamp
        FROM price_changes
        GROUP BY product_id
    ) latest_price_changes                             
      ON price_changes.product_id = latest_price_changes.product_id
     AND price_changes.effective_timestamp = latest_price_changes.max_timestamp
),  previous AS (
    SELECT price_changes.product_id,
           price_changes.price AS previous_price,
           MAX(price_changes.effective_timestamp) AS previous_timestamp
    FROM price_changes
    JOIN latest
      ON price_changes.product_id = latest.product_id
    WHERE price_changes.effective_timestamp < latest.latest_timestamp
    GROUP BY price_changes.product_id
)
SELECT * FROM products JOIN latest ON products.product_id = latest.product_id 
   ...> LEFT JOIN previous ON products.product_id = previous.product_id;
+------------+-----------------------------------+------------+---------------+---------------------+------------+----------------+---------------------+
| product_id |           product_name            | product_id | current_price |  latest_timestamp   | product_id | previous_price | previous_timestamp  |
+------------+-----------------------------------+------------+---------------+---------------------+------------+----------------+---------------------+
| 1          | Deluxe Sled                       | 1          | 109.25        | 2025-12-16 16:01:53 | 1          | 106.91         | 2025-12-16 01:51:05 |
| 2          | Holiday Trail Mix Trio            | 2          | 17.75         | 2025-12-16 09:28:13 | 2          | 21.38          | 2025-12-15 16:20:20 |
| 3          | Premium Cinnamon Roasted Almonds  | 3          | 143.65        | 2025-12-15 03:20:11 | 3          | 159.65         | 2025-12-14 09:52:09 |
| 4          | Deluxe Wrapping Paper             | 4          | 98.51         | 2025-12-16 01:33:41 | 4          | 105.6          | 2025-12-14 14:09:20 |
| 5          | Deluxe Roasted Cashews            | 5          | 124.04        | 2025-12-12 10:11:48 | 5          | 129.23         | 2025-12-12 04:08:45 |
| 6          | Festive Cookware Set              | 6          | 84.14         | 2025-12-15 11:31:40 | 6          | 88.97          | 2025-12-15 02:13:04 |
| 7          | Deluxe Mug                        | 7          | 123.09        | 2025-12-16 03:00:51 | 7          | 127.14         | 2025-12-13 07:25:12 |
| 8          | Premium Sled                      | 8          | 221.06        | 2025-12-15 22:33:48 | 8          | 241.99         | 2025-12-14 19:31:40 |
| 9          | Essential Sled                    | 9          | 255.88        | 2025-12-15 20:05:34 | 9          | 259.56         | 2025-12-12 12:47:13 |
| 10         | Family Snow Boots                 | 10         | 57.99         | 2025-12-15 20:53:45 | 10         | 64.88          | 2025-12-15 10:52:34 |
...
...
| 140        | Classic Mug                       | 140        | 173.05        | 2025-12-16 21:19:30 | 140        | 157.02         | 2025-12-09 17:58:07 |
| 141        | Family Fruit Assortment           | 141        | 69.97         | 2025-12-15 09:50:36 | 141        | 73.88          | 2025-12-13 14:35:53 |
| 142        | Classic Ornament                  | 142        | 35.05         | 2025-12-16 18:39:51 | 142        | 30.25          | 2025-12-16 13:44:42 |
| 143        | Essential Ornament                | 143        | 153.94        | 2025-12-15 07:27:06 | 143        | 143.04         | 2025-12-11 10:08:13 |
| 144        | Premium Trail Mix Trio            | 144        | 118.21        | 2025-12-16 16:25:16 | 144        | 114.22         | 2025-12-15 17:33:37 |
| 146        | Premium Book Collection           | 146        | 54.73         | 2025-12-13 07:07:19 | 146        | 65.71          | 2025-12-10 15:50:14 |
| 148        | Cozy Trail Mix Trio               | 148        | 107.81        | 2025-12-16 09:30:11 | 148        | 101.09         | 2025-12-15 05:37:27 |
| 149        | Family Cheddar Popcorn            | 149        | 72.6          | 2025-12-13 16:40:21 | 149        | 86.31          | 2025-12-13 13:18:14 |
| 150        | Holiday Headphones                | 150        | 138.66        | 2025-12-13 08:24:43 | 150        | 123.61         | 2025-12-12 10:49:22 |
+------------+-----------------------------------+------------+---------------+---------------------+------------+----------------+---------------------+
sqlite> 
```

Now, we just want
- product_name
- previous_price
- current_price
- difference of current_price and previous price

So,

```sql
WITH latest AS (                                        
    SELECT price_changes.product_id,
           price_changes.price AS current_price,
           price_changes.effective_timestamp AS latest_timestamp
    FROM price_changes
    JOIN (                                                   
        SELECT product_id, MAX(effective_timestamp) AS max_timestamp
        FROM price_changes
        GROUP BY product_id
    ) latest_price_changes                             
      ON price_changes.product_id = latest_price_changes.product_id
     AND price_changes.effective_timestamp = latest_price_changes.max_timestamp
),  previous AS (
    SELECT price_changes.product_id,
           price_changes.price AS previous_price,
           MAX(price_changes.effective_timestamp) AS previous_timestamp
    FROM price_changes
    JOIN latest
      ON price_changes.product_id = latest.product_id
    WHERE price_changes.effective_timestamp < latest.latest_timestamp
    GROUP BY price_changes.product_id
)
SELECT 
    products.product_name, 
    latest.current_price, 
    previous.previous_price, 
    (latest.current_price - previous.previous_price) as price_difference
FROM 
    products 
JOIN latest 
    ON products.product_id = latest.product_id 
LEFT JOIN previous 
    ON products.product_id = previous.product_id;
```


```
sqlite> WITH latest AS (                                        
    SELECT price_changes.product_id,
           price_changes.price AS current_price,
           price_changes.effective_timestamp AS latest_timestamp
    FROM price_changes
    JOIN (                                                   
        SELECT product_id, MAX(effective_timestamp) AS max_timestamp
        FROM price_changes
        GROUP BY product_id
    ) latest_price_changes                             
      ON price_changes.product_id = latest_price_changes.product_id
     AND price_changes.effective_timestamp = latest_price_changes.max_timestamp
),  previous AS (
    SELECT price_changes.product_id,
           price_changes.price AS previous_price,
           MAX(price_changes.effective_timestamp) AS previous_timestamp
    FROM price_changes
    JOIN latest
      ON price_changes.product_id = latest.product_id
    WHERE price_changes.effective_timestamp < latest.latest_timestamp
    GROUP BY price_changes.product_id
)
SELECT 
    products.product_name, latest.current_price, previous.previous_price, (latest.current_price - previous.previous_price) as price_difference
FROM 
    products 
JOIN latest 
    ON products.product_id = latest.product_id 
LEFT JOIN previous 
    ON products.product_id = previous.product_id;

+-----------------------------------+---------------+----------------+-------------------+
|           product_name            | current_price | previous_price | price_difference  |
+-----------------------------------+---------------+----------------+-------------------+
| Deluxe Sled                       | 109.25        | 106.91         | 2.34              |
| Holiday Trail Mix Trio            | 17.75         | 21.38          | -3.63             |
| Premium Cinnamon Roasted Almonds  | 143.65        | 159.65         | -16.0             |
| Deluxe Wrapping Paper             | 98.51         | 105.6          | -7.08999999999999 |
| Deluxe Roasted Cashews            | 124.04        | 129.23         | -5.18999999999998 |
| Festive Cookware Set              | 84.14         | 88.97          | -4.83             |
| Deluxe Mug                        | 123.09        | 127.14         | -4.05             |
| Premium Sled                      | 221.06        | 241.99         | -20.93            |
| Essential Sled                    | 255.88        | 259.56         | -3.68000000000001 |
...
...
| Classic Ornament                  | 35.05         | 30.25          | 4.8               |
| Essential Ornament                | 153.94        | 143.04         | 10.9              |
| Premium Trail Mix Trio            | 118.21        | 114.22         | 3.98999999999999  |
| Premium Book Collection           | 54.73         | 65.71          | -10.98            |
| Cozy Trail Mix Trio               | 107.81        | 101.09         | 6.72              |
| Family Cheddar Popcorn            | 72.6          | 86.31          | -13.71            |
| Holiday Headphones                | 138.66        | 123.61         | 15.05             |
+-----------------------------------+---------------+----------------+-------------------+
sqlite> 
```


Oh ok! That is it!

CTEs and some nasty number of JOINs.

Can we do it some other ways?

Of course we can, and we will


### ROW NUMBER - Window Function

We can even use some window functions because we are doing things per product so we can leverage [ROW_NUMBER](https://sqlite.org/windowfunctions.html#:~:text=in%20window%20functions%3A-,row_number(),-The%20number%20of) in our case.

> ROW_NUMBER: The number of the row within the current partition. Rows are numbered starting from 1 in the order defined by the ORDER BY clause in the window definition, or in arbitrary order otherwise.

So, we can partition by `product_id` in the `price_changes` table and order by the timestamp (latest first i.e. descending) and grab the first two prices.

```sql
WITH ranked_prices AS (
    SELECT price_changes.*,
           ROW_NUMBER() OVER (
               PARTITION BY price_changes.product_id
               ORDER BY price_changes.effective_timestamp DESC
           ) AS row_number
    FROM price_changes
)
SELECT
    *
FROM ranked_prices;
```

Here, we just partitioned the price table based on the latest to oldest timestamp (descending) and added the other column `row_number` that will be further used to get the `previous` and the `current` price and time.

```
sqlite> WITH ranked_prices AS (
    SELECT price_changes.*,
           ROW_NUMBER() OVER (
               PARTITION BY price_changes.product_id
               ORDER BY price_changes.effective_timestamp DESC
           ) AS row_number
    FROM price_changes
)
SELECT
    *
FROM ranked_prices;
+------+------------+--------+---------------------+------------+
|  id  | product_id | price  | effective_timestamp | row_number |
+------+------------+--------+---------------------+------------+
| 15   | 1          | 109.25 | 2025-12-16 16:01:53 | 1          |
| 14   | 1          | 106.91 | 2025-12-16 01:51:05 | 2          |
| 13   | 1          | 100.18 | 2025-12-14 15:58:03 | 3          |
| 12   | 1          | 99.03  | 2025-12-12 02:58:14 | 4          |
| 11   | 1          | 88.2   | 2025-12-12 02:21:09 | 5          |
| 10   | 1          | 73.9   | 2025-12-10 14:33:43 | 6          |
| 9    | 1          | 80.24  | 2025-12-08 16:11:11 | 7          |
| 8    | 1          | 78.88  | 2025-12-08 06:45:39 | 8          |
| 7    | 1          | 80.88  | 2025-12-07 10:43:54 | 9          |
| 6    | 1          | 88.49  | 2025-12-06 19:02:40 | 10         |
| 5    | 1          | 98.57  | 2025-12-04 04:14:31 | 11         |
| 4    | 1          | 119.12 | 2025-12-03 10:14:35 | 12         |
| 3    | 1          | 126.78 | 2025-12-02 18:40:38 | 13         |
| 2    | 1          | 148.63 | 2025-12-02 18:15:33 | 14         |
| 1    | 1          | 148.28 | 2025-12-01 05:25:35 | 15         |
...
...
| 1260 | 149        | 167.1  | 2025-12-01 06:36:43 | 15         |
| 1288 | 150        | 138.66 | 2025-12-13 08:24:43 | 1          |
| 1287 | 150        | 123.61 | 2025-12-12 10:49:22 | 2          |
| 1286 | 150        | 141.8  | 2025-12-10 07:06:06 | 3          |
| 1285 | 150        | 122.16 | 2025-12-09 20:01:54 | 4          |
| 1284 | 150        | 122.06 | 2025-12-06 22:27:41 | 5          |
| 1283 | 150        | 128.6  | 2025-12-06 13:03:18 | 6          |
| 1282 | 150        | 154.72 | 2025-12-05 08:15:08 | 7          |
| 1281 | 150        | 170.3  | 2025-12-04 15:25:50 | 8          |
| 1280 | 150        | 156.51 | 2025-12-03 19:11:12 | 9          |
| 1279 | 150        | 161.93 | 2025-12-02 02:47:10 | 10         |
| 1278 | 150        | 174.36 | 2025-12-02 01:39:14 | 11         |
| 1277 | 150        | 180.17 | 2025-12-01 20:36:02 | 12         |
| 1276 | 150        | 164.35 | 2025-12-01 07:09:29 | 13         |
| 1275 | 150        | 141    | 2025-12-01 01:29:46 | 14         |
+------+------------+--------+---------------------+------------+

```



```sql
WITH ranked_prices AS (
    SELECT price_changes.*,
           ROW_NUMBER() OVER (
               PARTITION BY price_changes.product_id
               ORDER BY price_changes.effective_timestamp DESC
           ) AS row_number
    FROM price_changes
)
SELECT
    products.product_name,
    current_price.price AS current_price,
    previous_price.price AS previous_price,
    current_price.price - previous_price.price AS price_difference
FROM products
LEFT JOIN ranked_prices AS current_price
       ON products.product_id = current_price.product_id
      AND current_price.row_number = 1
LEFT JOIN ranked_prices AS previous_price
       ON products.product_id = previous_price.product_id
      AND previous_price.row_number = 2;
```
This is pretty simple, we just get
- the `current_price` as `row_number = 1` from the `ranked_prices` CTE
- the `previous_price` as `row_number = 2` from the `ranked_prices` CTE

We got all we needed as:
- `products.product_name`
- `current_price.price AS current_price`
- `previous_price.price AS previous_price`
- `current_price.price - previous_price.price AS price_difference`

So, this is how `ROW_NUMBER` can be used here.


### With LEAD LAG - Window Functions

We can also use [LEAD](https://sqlite.org/windowfunctions.html#:~:text=does%20not%20exist.-,lead(expr),-lead(expr%2C%20offset) and [LAG](https://sqlite.org/windowfunctions.html#:~:text=a%20part%20of.-,lag(expr),-lag(expr%2C%20offset) window functions here. We can specifically use `LAG` as it was a challenge to get the second latest timestamped price.

We will partition the `price_changes` table on the `product_id` and order it by the `effective_timestamp`. This will give us the row before the timestamp of the current one, so the first row for each product can be empty (if sorted in ascending order of timestamp since there is no before row for the first row)

```sql
 WITH lagged_prices AS (
    SELECT
        product_id,
        price AS current_price,
        LAG(price) OVER (
            PARTITION BY product_id
            ORDER BY effective_timestamp
        ) AS previous_price,
        effective_timestamp
    FROM price_changes
)
SELECT * FROM lagged_prices;
```

Here, we simply selected the data that we need, `product_id`, `price` as the `current_price` since we can order by latest timestamp. 

```
sqlite> WITH lagged_prices AS (
    SELECT
        product_id,
        price AS current_price,
        LAG(price) OVER (
            PARTITION BY product_id
            ORDER BY effective_timestamp
        ) AS previous_price,
        effective_timestamp
    FROM price_changes
)
SELECT * FROM lagged_prices;
+------------+---------------+----------------+---------------------+
| product_id | current_price | previous_price | effective_timestamp |
+------------+---------------+----------------+---------------------+
| 1          | 148.28        |                | 2025-12-01 05:25:35 |
| 1          | 148.63        | 148.28         | 2025-12-02 18:15:33 |
| 1          | 126.78        | 148.63         | 2025-12-02 18:40:38 |
| 1          | 119.12        | 126.78         | 2025-12-03 10:14:35 |
| 1          | 98.57         | 119.12         | 2025-12-04 04:14:31 |
| 1          | 88.49         | 98.57          | 2025-12-06 19:02:40 |
| 1          | 80.88         | 88.49          | 2025-12-07 10:43:54 |
| 1          | 78.88         | 80.88          | 2025-12-08 06:45:39 |
| 1          | 80.24         | 78.88          | 2025-12-08 16:11:11 |
| 1          | 73.9          | 80.24          | 2025-12-10 14:33:43 |
| 1          | 88.2          | 73.9           | 2025-12-12 02:21:09 |
| 1          | 99.03         | 88.2           | 2025-12-12 02:58:14 |
| 1          | 100.18        | 99.03          | 2025-12-14 15:58:03 |
| 1          | 106.91        | 100.18         | 2025-12-16 01:51:05 |
| 1          | 109.25        | 106.91         | 2025-12-16 16:01:53 |
| 2          | 29.54         |                | 2025-12-03 14:21:10 |
| 2          | 34.33         | 29.54          | 2025-12-03 19:14:31 |
| 2          | 39.08         | 34.33          | 2025-12-04 06:13:48 |
| 2          | 32.71         | 39.08          | 2025-12-04 18:33:17 |
| 2          | 31.71         | 32.71          | 2025-12-05 22:36:14 |
| 2          | 28.88         | 31.71          | 2025-12-06 02:42:02 |
...
...
| 149        | 101.03        | 98.4           | 2025-12-10 00:37:46 |
| 149        | 95.68         | 101.03         | 2025-12-13 01:10:53 |
| 149        | 86.31         | 95.68          | 2025-12-13 13:18:14 |
| 149        | 72.6          | 86.31          | 2025-12-13 16:40:21 |
| 150        | 141           |                | 2025-12-01 01:29:46 |
| 150        | 164.35        | 141            | 2025-12-01 07:09:29 |
| 150        | 180.17        | 164.35         | 2025-12-01 20:36:02 |
| 150        | 174.36        | 180.17         | 2025-12-02 01:39:14 |
| 150        | 161.93        | 174.36         | 2025-12-02 02:47:10 |
| 150        | 156.51        | 161.93         | 2025-12-03 19:11:12 |
| 150        | 170.3         | 156.51         | 2025-12-04 15:25:50 |
| 150        | 154.72        | 170.3          | 2025-12-05 08:15:08 |
| 150        | 128.6         | 154.72         | 2025-12-06 13:03:18 |
| 150        | 122.06        | 128.6          | 2025-12-06 22:27:41 |
| 150        | 122.16        | 122.06         | 2025-12-09 20:01:54 |
| 150        | 141.8         | 122.16         | 2025-12-10 07:06:06 |
| 150        | 123.61        | 141.8          | 2025-12-12 10:49:22 |
| 150        | 138.66        | 123.61         | 2025-12-13 08:24:43 |
+------------+---------------+----------------+---------------------+

```


Then we can join. The conditions are important to filter.

- we get the lagged time stamp where we want the max of the timestamp since the latest timestamp will have the lagging timestamp for it from the CTE of lagged_prices and current_price as well.


```sql
SELECT
    products.product_name,
    lagged_prices.current_price,
    lagged_prices.previous_price,
    lagged_prices.current_price - lagged_prices.previous_price AS price_difference
FROM products
JOIN lagged_prices
  ON products.product_id = lagged_prices.product_id
WHERE lagged_prices.effective_timestamp = (
    SELECT MAX(effective_timestamp)
    FROM price_changes
    WHERE product_id = products.product_id
);

```


```sql
WITH lagged_prices AS (
    SELECT
        product_id,
        price AS current_price,
        LAG(price) OVER (
            PARTITION BY product_id
            ORDER BY effective_timestamp
        ) AS previous_price,
        effective_timestamp
    FROM price_changes
)
SELECT
    products.product_name,
    lagged_prices.current_price,
    lagged_prices.previous_price,
    lagged_prices.current_price - lagged_prices.previous_price AS price_difference
FROM products
JOIN lagged_prices
  ON products.product_id = lagged_prices.product_id
WHERE lagged_prices.effective_timestamp = (
    SELECT MAX(effective_timestamp)
    FROM price_changes
    WHERE product_id = products.product_id
);

```
This yields the same result.

This is all big, any smaller queries?

Not quite small.

### LIMIT OFFSET and CTE

Well here's a little short, but quite dirty.
```sql
SELECT *,
       current_price - previous_price AS price_difference
FROM (
    SELECT
        products.product_name,
        (SELECT price
         FROM price_changes
         WHERE price_changes.product_id = products.product_id
         ORDER BY effective_timestamp DESC
         LIMIT 1) AS current_price,
        (SELECT price
         FROM price_changes
         WHERE price_changes.product_id = products.product_id
         ORDER BY effective_timestamp DESC
         LIMIT 1 OFFSET 1) AS previous_price
    FROM products
);
```

This basically grabs the  price from the latest timestamp and the 2nd timestamp with offset and wraps it in a CTE to compute the difference.

Pretty slick if you ask me.

But hey! I am done for this! 

It was a great problem.

Getting data behind certain data is quite relatable and challenging enough to explore window functions and what not.

That's it from day 8.

See you tomorrow for day 9!




