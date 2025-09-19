---
type: sqlog
slug: sqlite-one-to-one-relations
title: "SQLite SQL: One to One Relations"
date: "2025-09-19"
tags: ["sqlite", "sql"]
---

## One to One Relations

I might missed this basic relationship model as I haven't really found it quite commonly used, but its still used in very specific examples.

That one relationship is `one-to-one`. As the name suggests, it maps one row to exactly one row.

Unlike the `one-to-many` relation which had one of the records/rows from the tables or entities connected by multiple ropes, both the records/rows from the tables or entities are connected by only one rope in `one-to-one` relation.

Where do you think this could be used?

Since it can't be applied to a user and a posts relation model, as that will be a too rigid constrained relationship. One user can only author one post, and one post can only be authored by one user.

So, can you think of a relation where one record or row is tied to exactly one other record or row in the other table?

Well, it could be user and his unique card, passport or even his subscriptions.

1. One user can only have one passport (or any other identity document).
2. One user can only have one credit card (for a specific bank)
3. One user can only have one subscription (to a specific service)

In those cases, I think the `one-to-one` relation serves well and is maybe the only way to get over the constraint.

You can definitely restrict the `one-to-many` relation to get this done, but might be a little wired. Will check that in other post.

## Creating a One to One Relation

How do we define a one to one relation, if we add a foreign key to the table that can refer multiple entities, so maybe if that foreign key is the primary key?

Ok, let me explain more clearly.

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE news_subscriptions (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    price INTEGER NOT NULL,
    status TEXT NOT NULL,
    expiry_date TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

In the above example, we have a `users` table and `news_subscriptions` table, where `news_subscriptions` table is a one to one relation to `users` table.

Here, the `user_id` is a foreign key in `news_subscriptions` table that can refer to the `id` of the `users` table.

This as you know is a `one-to-many` relation, as one user can have multiple news subscriptions.

Becuase here, one user can create multiple subscriptions, we don't want that right?

How can we restrict the many subscriptions to one user?

What if the `PRIMARY KEY` of the `subscriptions` table was the same as the `users` table?

That will restrict it to add any duplicate subscription for the same user.

Solved!

```sql
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS news_subscriptions;

CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE news_subscriptions (
    user_id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    price INTEGER NOT NULL,
    status TEXT NOT NULL,
    expiry_date TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

Here, instead of a separate `user_id` column, we have made the `user_id` the `PRIMARY KEY` of the `news_subscriptions` table.

NOTE: OOPs! don't forget to turn on the `PRAGMA foreign_keys` option.

> PRAGMA foreign_keys = ON;

This will tie the `news_subscriptions` table directly to an record/row that would require only one record/row in the `users` table.

Let's try inserting a few users:

```sql
INSERT INTO users(email, password) VALUES('harry', 'harry123'),
('ron', 'ron123'),
('malfoy', 'malfoy123');
```

This will insert 3 records/rows in the `users` table.

Now, let's try inserting a subscription:

```sql
INSERT INTO news_subscriptions(user_id, name, price, status, expiry_date)
VALUES(1, 'Wizardry Weekly', 500, 'active', '2025-12-31');
```

This will insert 1 record/row in the `news_subscriptions` table. This is a subscription for the user with `id` as `1`.

Now, if we try to insert another subscription for the same user:

```sql
INSERT INTO news_subscriptions(user_id, name, price, status, expiry_date)
VALUES(1, 'Muggle Cup', 250, 'active', '2025-12-31');
```

This will throw an error, as we have already inserted a subscription for the user with `id` as `1`.

Hence we successfully created a one to one relation between the `users` and `news_subscriptions` tables.


```
sqlite>
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE news_subscriptions (
    user_id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    price INTEGER NOT NULL,
    status TEXT NOT NULL,
    expiry_date TEXT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
sqlite> .mode table
sqlite> INSERT INTO users(email, password) VALUES('harry', 'harry123'), ('ron', 'ron123'), ('malfoy', 'malfoy123');
sqlite> SELECT * FROM users;
+----+--------+-----------+
| id | email  | password  |
+----+--------+-----------+
| 1  | harry  | harry123  |
| 2  | ron    | ron123    |
| 3  | malfoy | malfoy123 |
+----+--------+-----------+
sqlite> INSERT INTO news_subscriptions(user_id, name, price, status, expiry_date) VALUES(1, 'Wizardry Weekly', 500, 'active', '2025-12-31');
sqlite> SELECT * FROM news_subscriptions;
+---------+-----------------+-------+--------+-------------+
| user_id |      name       | price | status | expiry_date |
+---------+-----------------+-------+--------+-------------+
| 1       | Wizardry Weekly | 500   | active | 2025-12-31  |
+---------+-----------------+-------+--------+-------------+
sqlite> INSERT INTO news_subscriptions(user_id, name, price, status, expiry_date) VALUES(1, 'Muggle Cup', 250, 'active', '2025-12-31');
Runtime error: UNIQUE constraint failed: news_subscriptions.user_id (19)
sqlite>
```

Dead simple, this is how we create `one-to-one` relation.

We define the `primary key` of the child table as the `foreign key` of the parent table.

