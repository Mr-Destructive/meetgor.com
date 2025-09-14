---
type: sqlog
slug: sqlite-primary-key-table-constraint
title: 'SQLite SQL: PRIMARY KEY table constraint'
date: "2025-09-14"
tags: ["sqlite", "sql"]
---

## The PRIMARY KEY Table constraint

The PRIMARY KEY constraint is not actually a column constraint, it is a table constraint. In the previous section we learnt about how to handle single column-level `PRIMARY KEY`. In this section we'll understand how to use `PRIMARY KEY` as table constraint, with that we can use multiple columns to combine the key.

Since, `PRIMARY KEY` is a table level constraint, it only can be defined once for the table, as it needs to be unique for each record inserted in that table. So, we can define it with the column or define it at the end, doesn't matter, but it needs to be defined only once.

```sql
CREATE TABLE users (
    user_id INTEGER PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL
);
```
This is equivalent to:

```
CREATE TABLE users (
    user_id INTEGER,
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    PRIMARY KEY (user_id)
);
```

Both of these are setting the column `user_id` as the `PRIMARY KEY` for the table `users`.

This will do the same thing, we saw the first one in the previous blog post, in this we will do something different.

## Multiple column PRIMARY KEY

The schem design looks a little fragile, we can have `username` and `email` as duplicates right? But we have `user_id` as unique.

Rather, we want a combination of the `username` and `email` as unique, do we care about the `user_id` as the `PRIMARY KEY`?

It depends:
- If your application logic only wants to use `username` and `email` as unique, then the `user_id` is not required.
- If your applicaation logic only requires the `user_id` then it doesn't matter if the `username` of `email` are duplicated, the `user_id` will be unique.

So, assuming the 1st scenario, let's change the users table schema:

```sql
CREATE TABLE users (
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    PRIMARY KEY (username, email)
);
```

Here we have defined the `username` and `email` as `PRIMARY KEY` for the table `users`. This will take a combination of the `username` and the `email` coluns, and add a `UNIQUE` constraint to it, here there is one more addition that we did which is `NOT NULL` since both the columns are `TEXT` we need to add teh `NOT NULL` constraint, for `INTEGER` and `PRIMARY KEY` it will be handled for us.

Let's try inserting a few default values:

```sql
INSERT INTO users DEFAULT VALUES;
```

This won't work as the `username` and `email` columns are `NULL` and there is no `NOT NULL` constraint defined on them.

Let's try inserting actual values:

```sql
INSERT INTO users(username, email) VALUES('john', 'john@abc.com');
```

This successfully inserts the row into the table.

And if we try to insert it again, the duplicate combination of `username` and `email` will throw an error:

```sql
INSERT INTO users(username, email) VALUES('john', 'john@abc.com');
```

So, we either need the `username` or `email` to be unique, but not necessarily both.

```sql
INSERT INTO users(username, email) VALUES('john', 'john_new@abc.com');
```

This successfully inserts the row into the table, as the `username` is same but `email` is unique. Since the PRIMARY key is a combination of `username` and `email` we can allow the combinational unique values into a new record.

```sql
SELECT * FROM users;
```

Here's the SQLog :)

```
sqlite> CREATE TABLE users (
    username TEXT NOT NULL,
    email TEXT NOT NULL,
    PRIMARY KEY (username, email)
);
sqlite> INSERT INTO users DEFAULT VALUES;
Runtime error: NOT NULL constraint failed: users.username (19)
sqlite> INSERT INTO users(username, email) VALUES('john', 'john@abc.com');
sqlite> SELECT * FROM users;
+----------+--------------+
| username |    email     |
+----------+--------------+
| john     | john@abc.com |
+----------+--------------+
sqlite> SELECT rowid, * FROM users;
+-------+----------+--------------+
| rowid | username |    email     |
+-------+----------+--------------+
| 1     | john     | john@abc.com |
+-------+----------+--------------+
sqlite> INSERT INTO users(username, email) VALUES('john', 'john@abc.com');
Runtime error: UNIQUE constraint failed: users.username, users.email (19)
sqlite> INSERT INTO users(username, email) VALUES('john', 'john_new@abc.com');
sqlite> SELECT * FROM users;
+----------+------------------+
| username |      email       |
+----------+------------------+
| john     | john@abc.com     |
| john     | john_new@abc.com |
+----------+------------------+
sqlite> SELECT rowid, * FROM users;
+-------+----------+------------------+
| rowid | username |      email       |
+-------+----------+------------------+
| 1     | john     | john@abc.com     |
| 2     | john     | john_new@abc.com |
+-------+----------+------------------+
sqlite>
```


I think that is all we need about `PRIMARY KEY` for the moment, if you need any specific example, we will be exploring `AUTOINCREMENT`, `NOT NULL` cases for `PRIMARY KEY`s and other schema design principles in separate sections.

