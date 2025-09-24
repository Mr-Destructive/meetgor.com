---
type: sqlog
slug: sqlite-create-temp-table
title: "SQLite SQL: Create Temporary Table"
date: "2025-09-24"
tags: ["sqlite", "sql"]
---

## Temporary Table in SQLite

A Temporary table as the name suggests, is a temporary table. It only resisdes on the database until the current session of the database.

In case of sqlite, the temporary table is available in the session until the file reader or session driver closes the connection.

The table is created in a separate database file called `temp` that is stored in your temporary path.

## Create a temporary table

To create the temporary table, simply use the `TEMP` or `TEMPORARY` keyword before the `TABLE` in `CREATE TABLE` expression. So, `CREATE TEMP TABLE` or `CREATE TEMPORARY TABLE` will be the notation to create a temporary table in sqlite.

```sql
CREATE TEMP TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT
);
```
This will create a temporary table called `users` in the `temp` database.

You can insert, update, delete, query and relate to other tables just like a normal table.

```sql
INSERT INTO users (name) VALUES ('abc'), ('def'), ('ghi');
```

```sql
SELECT * FROM users;
```

```sql
UPDATE users SET name = 'xyz' WHERE name = 'abc';
SELECT * FROM users;
```

```sql
DELETE FROM users WHERE name = 'def';
```

```sql
CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    title TEXT,
    content TEXT,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users (id)
);
```

```sqlite
$ sqlite3 myusers.db

sqlite> CREATE TEMP TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT
);

sqlite> .tables
temp.users

sqlite> INSERT INTO users (name) VALUES ('abc'), ('def'), ('ghi');
sqlite> .mode table
sqlite> SELECT * FROM users;
+----+------+
| id | name |
+----+------+
| 1  | abc  |
| 2  | def  |
| 3  | ghi  |
+----+------+
sqlite> UPDATE users SET name = 'xyz' WHERE name = 'abc';
SELECT * FROM users;
+----+------+
| id | name |
+----+------+
| 1  | xyz  |
| 2  | def  |
| 3  | ghi  |
+----+------+
sqlite> DELETE FROM users WHERE name = 'def';
sqlite> SELECT * FROM users;
+----+------+
| id | name |
+----+------+
| 1  | xyz  |
| 3  | ghi  |
+----+------+
sqlite> CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    title TEXT,
    content TEXT,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

sqlite> .tables
posts temp.users

sqlite> INSERT INTO posts(title, content, user_id) VALUES('banger post', 'ai slop', 1);
sqlite> INSERT INTO posts(title, content, user_id) VALUES('ai slop', 'banger content', 3);
sqlite> SELECT * FROM posts;
+----+-------------+----------------+---------+
| id |    title    |    content     | user_id |
+----+-------------+----------------+---------+
| 1  | banger post | ai slop        | 1       |
| 2  | ai slop     | banger content | 3       |
+----+-------------+----------------+---------+
sqlite>

```

So, this creates a temporary table, but what if we already have a `users` table in the database?

> You can notice, when I query all the tables with the dot-command `.tables`, it adds a prefix of `temp.` to the temporary table. That is an indication of the table being stored in a temporary database file. Hence if there were two `users` table created, we would be able to identify it with the `temp` and the `main` as the database name.

The `main` is optional, but if you want to refer to the `temporary` table then explicitly mention it as `temp.users` or `temp.<table_name>`

However, if you try to perform anything on the actual table, you will have to explicitly mention the `main` keyword, as the temp table will take precedence over the table in the main database.

I reponed the database as `myusers.db` , now the `users` table which was the temporary table, is gone.

```sql
.tables
```

This only shows, `posts` table, but the temporary table was destroyed.

Let's create back the `users` table as a temporary table in the database.

```sql
CREATE TEMP TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT
);
```

This will create the table in the `temp` database.

Now, if we check the tables, we will see the `temp.users` table.

```sql
.tables
```
Let's create a new table `users` in the main database.

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT
);
```
This would create a `users` table that will be stored permanently in the main database i.e. the `myusers.db` in my case to the file.

Let's view what the `.tables` output, it shows 3 tables now

- `posts`
- `temp.users`
- `users`

Now, here, if you insert bunch of different values into the `users` table.

```sql
INSERT INTO users (name) VALUES ('abc'), ('def'), ('ghi');
```

And query the `users` table

```sql
SELECT * FROM users;
```

Which table is that inserted and queried?

The `temp.users` table, because the precedence of temporary table is higher than the main table.

So, let's try to query the `temp.users` table.

```sql
SELECT * FROM temp.users;
```

As you can see the table has the records. 

But if we try to query the `main.users` table, it will have no records.

```sql
SELECT * FROM main.users;
```
This is empty.
As, expected, just the `users` table will refer to th temporary table and not the main table.

So, let's insert different values in the `main.users` table.

```sql
INSERT INTO main.users (name) VALUES ('pqr'), ('stu');
```

And query the `main.users` table.

```sql
SELECT * FROM main.users;
```
This will show the inserted records in the `main.users` table.

Here's the SQLog :)

```sqlite
$ sqlite3 myusers.db
sqlite> .tables
posts

sqlite> CREATE TEMP TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT
);

sqlite> .tables
posts       temp.users

sqlite> CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT
);

sqlite> .tables
posts       temp.users  users

sqlite> .schema temp.users

CREATE TABLE temp.users (
    id INTEGER PRIMARY KEY,
    name TEXT
);

sqlite> .schema main.users
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT
);

sqlite> SELECT * FROM users;
sqlite> SELECT * FROM temp.users;
sqlite> INSERT INTO users (name) VALUES ('abc'), ('def'), ('ghi');
sqlite> SELECT * FROM users;
+----+------+
| id | name |
+----+------+
| 1  | abc  |
| 2  | def  |
| 3  | ghi  |
+----+------+
sqlite> SELECT * FROM temp.users;
+----+------+
| id | name |
+----+------+
| 1  | abc  |
| 2  | def  |
| 3  | ghi  |
+----+------+
sqlite> SELECT * FROM main.users;
sqlite> INSERT INTO main.users (name) VALUES ('pqr'), ('stu');
sqlite> SELECT * FROM main.users;
+----+------+
| id | name |
+----+------+
| 1  | pqr  |
| 2  | stu  |
+----+------+
sqlite> SELECT * FROM temp.users;
+----+------+
| id | name |
+----+------+
| 1  | abc  |
| 2  | def  |
| 3  | ghi  |
+----+------+
sqlite> SELECT * FROM users;
+----+------+
| id | name |
+----+------+
| 1  | abc  |
| 2  | def  |
| 3  | ghi  |
+----+------+
sqlite>
```

So, to sum up:
- `CREATE TEMP TABLE` or `CREATE TEMPORARY TABLE` will create a temporary table in sqlite.
- The temporary table is available in the session until the file reader or session driver closes the connection.
- The temporary table is stored on the separate temporary file
- The temporary table is prefixed with the `temp` table.
- If there are two tables, one temporary and one permanent, with the same name, then the temporary table will be prefered unless `main` is prefixed to the table name.
