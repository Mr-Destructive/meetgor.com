---
type: "sqlog"
title: "Advent of SQL 2025 Day 9: Evergreen Market Orders"
slug: "advent-of-sql-2025-day-9"
date: 2025-12-24
series: ["advent-of-sql-2025"]
tags: ["sqlite", "sql", "advent-of-sql"]
---

## Advent of SQL, Day 9 - Evergreen Market Orders

We are on day 9 of advent of SQL, and I feel good so far.

Let's see what we learn today?

Let's get the inserts for the day.

```
sqlite> .read day9-inserts.sql
sqlite> .schema
CREATE TABLE orders (
    id           INT PRIMARY KEY,
    customer_id  INT,
    created_at   TIMESTAMP,
    order_data   JSONB
);
sqlite> .mode table
sqlite> SELECT * FROM orders limit 10;
+----+-------------+---------------------+--------------------------------------------------------------+
| id | customer_id |     created_at      |                          order_data                          |
+----+-------------+---------------------+--------------------------------------------------------------+
| 1  | 1           | 2025-11-21 13:08:22 | {"shipping": {"method": "standard"}, "gift": {"wrapped": tru |
|    |             |                     | e}}                                                          |
+----+-------------+---------------------+--------------------------------------------------------------+
| 2  | 1           | 2025-11-21 18:42:58 | {"shipping": {"method": "overnight"}, "risk": {"flag": "high |
|    |             |                     | "}, "gift": {"wrapped": false}}                              |
+----+-------------+---------------------+--------------------------------------------------------------+
| 3  | 1           | 2025-11-21 21:01:46 | {"shipping": {"method": "standard"}, "gift": {"wrapped": fal |
|    |             |                     | se}}                                                         |
+----+-------------+---------------------+--------------------------------------------------------------+
| 4  | 1           | 2025-11-24 13:17:27 | {"shipping": {"method": "standard"}, "gift": {"wrapped": tru |
|    |             |                     | e}}                                                          |
+----+-------------+---------------------+--------------------------------------------------------------+
| 5  | 1           | 2025-11-24 21:09:46 | {"shipping": {"method": "overnight"}, "gift": {"wrapped": fa |
|    |             |                     | lse}}                                                        |
+----+-------------+---------------------+--------------------------------------------------------------+
| 6  | 1           | 2025-11-25 07:24:55 | {"shipping": {"method": "standard"}, "risk": {"flag": "mediu |
|    |             |                     | m"}, "gift": {"wrapped": true}}                              |
+----+-------------+---------------------+--------------------------------------------------------------+
| 7  | 1           | 2025-11-25 17:42:36 | {"shipping": {"method": "standard"}, "gift": {"wrapped": fal |
|    |             |                     | se}}                                                         |
+----+-------------+---------------------+--------------------------------------------------------------+
| 8  | 1           | 2025-11-27 02:34:24 | {"shipping": {"method": "express"}, "gift": {"wrapped": true |
|    |             |                     | }}                                                           |
+----+-------------+---------------------+--------------------------------------------------------------+
| 9  | 1           | 2025-11-30 22:43:54 | {"shipping": {"method": "overnight"}, "gift": {"wrapped": tr |
|    |             |                     | ue}}                                                         |
+----+-------------+---------------------+--------------------------------------------------------------+
| 10 | 1           | 2025-12-01 04:03:33 | {"shipping": {"method": "express"}, "risk": {"flag": "medium |
|    |             |                     | "}, "gift": {"wrapped": false}}                              |
+----+-------------+---------------------+--------------------------------------------------------------+
sqlite> 
```

Looks like we will deal with JSON today, seems exciting. I haven't dealt with JSON in SQLite yet, today will break it.

Let's get some sample inserts for you to play in the browser.

LIMITING TO 20, there are more than 400 rows!

```sql
DROP TABLE IF EXISTS orders;

CREATE TABLE orders (
    id           INT PRIMARY KEY,
    customer_id  INT,
    created_at   TIMESTAMP,
    order_data   JSONB
);

INSERT INTO orders (id, customer_id, created_at, order_data) VALUES
    (1, 1, '2025-11-21 13:08:22', '{"shipping": {"method": "standard"}, "gift": {"wrapped": true}}'),
    (2, 1, '2025-11-21 18:42:58', '{"shipping": {"method": "overnight"}, "risk": {"flag": "high"}, "gift": {"wrapped": false}}'),
    (3, 1, '2025-11-21 21:01:46', '{"shipping": {"method": "standard"}, "gift": {"wrapped": false}}'),
    (4, 1, '2025-11-24 13:17:27', '{"shipping": {"method": "standard"}, "gift": {"wrapped": true}}'),
    (5, 1, '2025-11-24 21:09:46', '{"shipping": {"method": "overnight"}, "gift": {"wrapped": false}}'),
    (6, 1, '2025-11-25 07:24:55', '{"shipping": {"method": "standard"}, "risk": {"flag": "medium"}, "gift": {"wrapped": true}}'),
    (7, 1, '2025-11-25 17:42:36', '{"shipping": {"method": "standard"}, "gift": {"wrapped": false}}'),
    (8, 1, '2025-11-27 02:34:24', '{"shipping": {"method": "express"}, "gift": {"wrapped": true}}'),
    (9, 1, '2025-11-30 22:43:54', '{"shipping": {"method": "overnight"}, "gift": {"wrapped": true}}'),
    (10, 1, '2025-12-01 04:03:33', '{"shipping": {"method": "express"}, "risk": {"flag": "medium"}, "gift": {"wrapped": false}}'),
    (11, 1, '2025-12-02 05:19:10', '{"shipping": {"method": "overnight"}, "risk": {"flag": "low"}, "gift": {"wrapped": true}}'),
    (12, 1, '2025-12-03 16:25:56', '{"shipping": {"method": "overnight"}, "risk": {"flag": "medium"}, "gift": {"wrapped": true}}'),
    (13, 1, '2025-12-10 19:34:28', '{"shipping": {"method": "standard"}, "gift": {"wrapped": false}}'),
    (14, 1, '2025-12-16 19:23:53', '{"shipping": {"method": "express"}, "gift": {"wrapped": false}}'),
    (15, 2, '2025-11-23 19:11:23', '{"shipping": {"method": "overnight"}, "gift": {"wrapped": true}}'),
    (16, 2, '2025-11-28 15:23:27', '{"shipping": {"method": "express"}, "gift": {"wrapped": true}}'),
    (17, 2, '2025-11-30 12:05:36', '{"shipping": {"method": "overnight"}, "risk": {"flag": "low"}, "gift": {"wrapped": true}}'),
    (18, 2, '2025-12-03 07:03:06', '{"shipping": {"method": "standard"}, "gift": {"wrapped": false}}'),
    (19, 2, '2025-12-07 13:55:13', '{"shipping": {"method": "express"}, "risk": {"flag": "high"}, "gift": {"wrapped": false}}'),
    (20, 2, '2025-12-08 07:17:31', '{"shipping": {"method": "standard"}, "gift": {"wrapped": false}}');
```

```sql
SELECT * FROM orders;
```

Let's get to the problem now.

## Problem

> Build a report using the orders table that shows the latest order for each customer, along with their requested shipping method, gift wrap choice (as true or false), and the risk flag in separate columns.
> 
> Order the report by the most recent order first so Evergreen Market can reach out to them ASAP.

Ok, so we need for each customer the latest order with the following details:
- shipping method
- gift wrap choice
- risk flag

These all I think are in the same column as a JSON string or blob. We need to extract those out from the column.

Let's first check the `orders` table.

It has a few columns:
- id
- customer_id
- created_at
- order_data

We do require the `order_data` as that is the column that has JSON.

Also the problem said to give the most recent order, so we need to order by `created_at` in a reverse way, the latest first. Also we need it per customer, so we need to group by `customer_id`.

Let's see how to get the data inside JSON in SQLite.

### JSON Extract

Well we have [json_extract](https://sqlite.org/json1.html#jex) which can give us the value of the key from the given json data string.

The function takes in the column name from the table which would be the column containing the json data and then the second parameter is the path to the key, in this case if we want to get the `method` from the `shipping` key, we can use `$.shipping.method` which means from the root `$` get the `shipping` key, and inside that (`shipping` key) give the value inside the `method` key.

If the path is not present, in our case the `risk`  key is very rarely present in the actual original json data, the function skips the further key lookup, it returns `NULL`.

```sql
SELECT json_extract(orders.order_data, '$.shipping.method') FROM orders LIMIT 5;
```

```
sqlite> select json_extract(orders.order_data, '$.shipping.method') FROM orders LIMIT 5;
+------------------------------------------------------+
| json_extract(orders.order_data, '$.shipping.method') |
+------------------------------------------------------+
| standard                                             |
| overnight                                            |
| standard                                             |
| standard                                             |
| overnight                                            |
+------------------------------------------------------+
sqlite> select *, json_extract(orders.order_data, '$.shipping.method') FROM orders LIMIT 5;
+----+-------------+---------------------+--------------------------------------------------------------+------------------------------------------------------+
| id | customer_id |     created_at      |                          order_data                          | json_extract(orders.order_data, '$.shipping.method') |
+----+-------------+---------------------+--------------------------------------------------------------+------------------------------------------------------+
| 1  | 1           | 2025-11-21 13:08:22 | {"shipping": {"method": "standard"}, "gift": {"wrapped": tru | standard                                             |
|    |             |                     | e}}                                                          |                                                      |
+----+-------------+---------------------+--------------------------------------------------------------+------------------------------------------------------+
| 2  | 1           | 2025-11-21 18:42:58 | {"shipping": {"method": "overnight"}, "risk": {"flag": "high | overnight                                            |
|    |             |                     | "}, "gift": {"wrapped": false}}                              |                                                      |
+----+-------------+---------------------+--------------------------------------------------------------+------------------------------------------------------+
| 3  | 1           | 2025-11-21 21:01:46 | {"shipping": {"method": "standard"}, "gift": {"wrapped": fal | standard                                             |
|    |             |                     | se}}                                                         |                                                      |
+----+-------------+---------------------+--------------------------------------------------------------+------------------------------------------------------+
| 4  | 1           | 2025-11-24 13:17:27 | {"shipping": {"method": "standard"}, "gift": {"wrapped": tru | standard                                             |
|    |             |                     | e}}                                                          |                                                      |
+----+-------------+---------------------+--------------------------------------------------------------+------------------------------------------------------+
| 5  | 1           | 2025-11-24 21:09:46 | {"shipping": {"method": "overnight"}, "gift": {"wrapped": fa | overnight                                            |
|    |             |                     | lse}}                                                        |                                                      |
+----+-------------+---------------------+--------------------------------------------------------------+------------------------------------------------------+
sqlite> 

```

Let's grab the other two.

```sql
SELECT
    json_extract(order_data, '$.shipping.method'),
    json_extract(order_data, '$.gift.wrapped'),
    json_extract(order_data, '$.risk.flag') 
FROM orders;
```

```
sqlite> select json_extract(order_data, '$.shipping.method'), json_extract(order_data, '$.gift.wrapped'), json_extract(order_data, '$.risk.flag') FROM orders LIMIT 5;
+-----------------------------------------------+--------------------------------------------+-----------------------------------------+
| json_extract(order_data, '$.shipping.method') | json_extract(order_data, '$.gift.wrapped') | json_extract(order_data, '$.risk.flag') |
+-----------------------------------------------+--------------------------------------------+-----------------------------------------+
| standard                                      | 1                                          |                                         |
| overnight                                     | 0                                          | high                                    |
| standard                                      | 0                                          |                                         |
| standard                                      | 1                                          |                                         |
| overnight                                     | 0                                          |                                         |
+-----------------------------------------------+--------------------------------------------+-----------------------------------------+
sqlite> 
```

Hmm, interesting, it automatically converted `false` to `0`, and `true` to `1` as you can see in the actual json data it was `true` and `false`. Nice, but could go the other way as well, if you were checking for true or false instead of 0 and 1. Both are same in SQLite, they are treated as truth and false values.

We have all the things we need, now?

Well, let's group by the `customer_id` and order by the `created_at` to get the info about each customer with its latest order.

```sql
SELECT *, json_extract(order_data, '$.shipping.method'), json_extract(order_data, '$.gift.wrapped'), json_extract(order_data, '$.risk.flag') FROM orders GROUP BY customer_id ORDER by created_at LIMIT 5;

```

But there I see the problem, how can we group by customer_id and find the latest order for that customer! The order by will only work after the grouping has been done, so its not necessarily the per order for each customer.

Hmm!

We need to make a subquery to get the latest order for each customer it seems.

```sql
SELECT 
    orders.customer_id,
    orders.created_at,
    json_extract(orders.order_data, '$.shipping.method') AS shipping_method,
    json_extract(orders.order_data, '$.gift.wrapped') AS gift_wrap,
    json_extract(orders.order_data, '$.risk.flag') AS risk_flag
FROM orders
WHERE orders.created_at = (
    SELECT MAX(created_at)
    FROM orders AS latest_order
    WHERE orders.customer_id = latest_order.customer_id
)
ORDER BY orders.created_at DESC;

```
We just added this 

```sql
WHERE orders.created_at = (
    SELECT MAX(created_at)
    FROM orders AS latest_order
    WHERE orders.customer_id = latest_order.customer_id
)
```
This will filter only the latest created order for each customer and then we can grab the details from json in each row per customer after this isolates the row with the latest created at time.

```
sqlite> SELECT 
    orders.customer_id,
    orders.created_at,
    json_extract(orders.order_data, '$.shipping.method') AS shipping_method,
    json_extract(orders.order_data, '$.gift.wrapped') AS gift_wrap,
    json_extract(orders.order_data, '$.risk.flag') AS risk_flag
FROM orders
WHERE orders.created_at = (
    SELECT MAX(created_at)
    FROM orders AS latest_order
    WHERE orders.customer_id = latest_order.customer_id
)
ORDER BY orders.created_at DESC;

+-------------+---------------------+-----------------+-----------+-----------+
| customer_id |     created_at      | shipping_method | gift_wrap | risk_flag |
+-------------+---------------------+-----------------+-----------+-----------+
| 32          | 2025-12-17 21:17:39 | overnight       | 0         |           |
| 15          | 2025-12-17 19:21:33 | express         | 0         | medium    |
| 50          | 2025-12-17 14:47:54 | express         | 1         | low       |
| 43          | 2025-12-17 14:23:46 | express         | 1         |           |
| 27          | 2025-12-17 14:05:13 | standard        | 1         |           |
| 3           | 2025-12-17 14:02:28 | standard        | 1         | high      |
| 49          | 2025-12-17 13:28:49 | express         | 1         | high      |
| 36          | 2025-12-17 11:11:29 | overnight       | 1         |           |
| 31          | 2025-12-17 08:05:46 | express         | 0         |           |
| 16          | 2025-12-17 07:32:36 | express         | 0         |           |
| 38          | 2025-12-17 06:05:12 | standard        | 1         |           |
| 44          | 2025-12-17 05:28:54 | standard        | 1         |           |
| 9           | 2025-12-17 04:33:08 | express         | 1         |           |
| 23          | 2025-12-17 03:01:49 | express         | 0         |           |
| 21          | 2025-12-16 23:53:14 | overnight       | 1         |           |
| 25          | 2025-12-16 20:49:58 | overnight       | 1         | high      |
| 46          | 2025-12-16 19:38:37 | standard        | 0         |           |
| 1           | 2025-12-16 19:23:53 | express         | 0         |           |
| 28          | 2025-12-16 18:20:55 | standard        | 0         | low       |
| 40          | 2025-12-16 17:54:05 | express         | 0         |           |
| 13          | 2025-12-16 16:11:16 | standard        | 1         |           |
| 24          | 2025-12-16 14:19:45 | overnight       | 0         |           |
| 11          | 2025-12-16 11:20:31 | standard        | 1         | medium    |
| 17          | 2025-12-16 08:19:36 | standard        | 0         |           |
| 4           | 2025-12-16 04:38:51 | express         | 0         |           |
| 34          | 2025-12-16 02:11:57 | express         | 0         |           |
| 30          | 2025-12-15 15:32:04 | overnight       | 0         | medium    |
| 48          | 2025-12-15 13:03:59 | standard        | 1         |           |
| 41          | 2025-12-15 13:00:00 | standard        | 0         | high      |
| 45          | 2025-12-15 11:37:57 | standard        | 0         |           |
| 7           | 2025-12-14 23:39:47 | express         | 0         |           |
| 35          | 2025-12-14 22:46:36 | express         | 1         | high      |
| 47          | 2025-12-14 20:53:07 | standard        | 0         |           |
| 22          | 2025-12-14 12:38:58 | standard        | 0         | medium    |
| 12          | 2025-12-14 07:59:28 | standard        | 1         | medium    |
| 18          | 2025-12-14 04:55:34 | overnight       | 0         | low       |
| 20          | 2025-12-14 04:54:07 | overnight       | 0         |           |
| 14          | 2025-12-13 07:44:19 | standard        | 1         |           |
| 6           | 2025-12-13 07:03:12 | overnight       | 1         |           |
| 10          | 2025-12-13 04:23:37 | standard        | 0         | medium    |
| 19          | 2025-12-13 03:29:15 | standard        | 0         |           |
| 8           | 2025-12-12 12:42:18 | overnight       | 0         |           |
| 26          | 2025-12-11 17:35:46 | standard        | 0         | low       |
| 37          | 2025-12-11 13:55:35 | overnight       | 1         |           |
| 33          | 2025-12-09 12:30:54 | express         | 1         |           |
| 2           | 2025-12-08 07:17:31 | standard        | 0         |           |
| 42          | 2025-12-08 02:48:12 | overnight       | 0         | medium    |
| 5           | 2025-12-06 17:53:53 | overnight       | 1         |           |
| 39          | 2025-12-06 14:38:29 | overnight       | 1         |           |
| 29          | 2025-12-03 05:10:32 | overnight       | 1         | high      |
+-------------+---------------------+-----------------+-----------+-----------+
sqlite> 
```

Now, that is it! I guess, a quick solution for getting this data sorted.

However there is one more way without needing the subquery.

### ROW NUMBER - Window Function

We can create partition for each customer with the `customer_id` and then sort by the `created_at` date in a latest first order (descending) and simply take the first row for extraction. This is kind of the same thing, but a kind of 'gives a good feel to me' solution. Elite Mindset!

```sql
WITH ranked_orders AS (
    SELECT 
        orders.customer_id,
        orders.created_at,
        json_extract(orders.order_data, '$.shipping.method') AS shipping_method,
        json_extract(orders.order_data, '$.gift.wrapped') AS gift_wrap,
        json_extract(orders.order_data, '$.risk.flag') AS risk_flag,
        ROW_NUMBER() OVER (PARTITION BY orders.customer_id ORDER BY orders.created_at DESC) AS row_num
    FROM orders
)
SELECT 
    customer_id,
    created_at,
    shipping_method,
    gift_wrap,
    risk_flag
FROM ranked_orders
WHERE row_num = 1
ORDER BY created_at DESC;
```

We are just doing the same thing as explained.

This thing

`ROW_NUMBER() OVER (PARTITION BY orders.customer_id ORDER BY orders.created_at DESC) AS row_num`

It will partition the table `orders` with `customer_id` and order by `created_at` latest first. And we assign each row as the `row_number` and grab the 1st row per customer to grab the latest order (filter with WHERE when calling the cte, or querying. 


However we created the CTE `ranked_orders` or you can call it `latest_orders`, to get the `row_num` filter to only the first `1` and for that we grab the already extracted json data as well the other columns.

Simple!

That is it!

Day 9 was easy peasy!

Onwards day 10! Catch you tomorrow!
