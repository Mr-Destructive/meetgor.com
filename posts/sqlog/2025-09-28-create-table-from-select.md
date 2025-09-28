---
type: sqlog
slug: sqlite-create-table-from-select
title: "SQLite SQL: Create Table with Select"
date: "2025-09-28"
tags: ["sqlite", "sql"]
---

## Creating table with Select Query

Sometimes you want to create a new table based on data that already exists,maybe you’re archiving old rows, generating a summary, cloning a table, or transforming data into a new structure. Doing this might require you to write a `CREATE TABLE` statement with all column definitions, then running one or more `INSERT INTO` statements to populate it. That’s a lot of work.

Well, SQL is more flexible then you might think, Instead of manually defining columns and inserting data, you write one statement that both builds the table and fills it with the rows returned by your `SELECT` query. This makes it incredibly useful for backups, and whatever you are doing (hopefully not taking down the prod db).

We have CTAS, `CREATE TABLE AS SELECT` statements.
This basically means to take whatever the SELECT gives, and turn it into a table.

## Creating CTAS

Lets's start with a simple example, we will use our goodol users table first and populate it with some data.

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER NOT NULL
);
```

This would create a table called `users` with the columns `id`, `name`, and `age`. The `id` column is the primary key, and the `name` and `age` columns are not null.

Let's insert some data into the table.

```sql
INSERT INTO users (name, age)
VALUES ('abc', 20), ('def', 30), ('ghi', 40);
```

This would insert three rows into the `users` table, with the names `abc`, `def`, and `ghi`, and the ages 20, 30, and 40.

Now, let's use the `CREATE TABLE AS SELECT` statement to create a new table called `users_copy` that is a copy of the `users` table.

```sql
CREATE TABLE users_copy AS SELECT * FROM users;
```
This would create a new table called `users_copy` that is a copy of the `users` table, with the same columns and data.

It will copy the same column structure and data from the `users` table to the `users_copy` table.

```sql
SELECT * FROM users_copy;
```

Let's check the schema of the new table `users_copy`

```sql
PRAGMA table_info(users_copy);
PRAGMA table_info(users);
```

This would print the schema of the `users_copy` table.

```
PRAGMA table_info(users_copy);
+-----+------+------+---------+------------+----+
| cid | name | type | notnull | dflt_value | pk |
+-----+------+------+---------+------------+----+
| 0   | id   | INT  | 0       |            | 0  |
| 1   | name | TEXT | 0       |            | 0  |
| 2   | age  | INT  | 0       |            | 0  |
+-----+------+------+---------+------------+----+


PRAGMA table_info(users);
+-----+------+---------+---------+------------+----+
| cid | name |  type   | notnull | dflt_value | pk |
+-----+------+---------+---------+------------+----+
| 0   | id   | INTEGER | 0       |            | 1  |
| 1   | name | TEXT    | 1       |            | 0  |
| 2   | age  | INTEGER | 1       |            | 0  |
+-----+------+---------+---------+------------+----+
```

You can see the difference here, the `users.id` column is a primary key in the `users` table, however in the `users_copy` table it isn't. Also the constraints like `notnull` are not reciprocated in the `users_copy` table.

So, is this a good way to create a copy of a table?

Maybe, its just copying the data and the bare-bone structure of the table and not the schema.

```sqlite

sqlite> CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER NOT NULL
);
sqlite> INSERT INTO users (name, age)
VALUES ('abc', 20), ('def', 30), ('ghi', 40);
sqlite> .mode table
sqlite> SELECT * FROM users;
+----+------+-----+
| id | name | age |
+----+------+-----+
| 1  | abc  | 20  |
| 2  | def  | 30  |
| 3  | ghi  | 40  |
+----+------+-----+
sqlite> CREATE TABLE users_copy AS SELECT * FROM users;
sqlite> SELECT * FROM users_copy;
+----+------+-----+
| id | name | age |
+----+------+-----+
| 1  | abc  | 20  |
| 2  | def  | 30  |
| 3  | ghi  | 40  |
+----+------+-----+
sqlite> PRAGMA table_info(users_copy);
+-----+------+------+---------+------------+----+
| cid | name | type | notnull | dflt_value | pk |
+-----+------+------+---------+------------+----+
| 0   | id   | INT  | 0       |            | 0  |
| 1   | name | TEXT | 0       |            | 0  |
| 2   | age  | INT  | 0       |            | 0  |
+-----+------+------+---------+------------+----+
sqlite> PRAGMA table_info(users);
+-----+------+---------+---------+------------+----+
| cid | name |  type   | notnull | dflt_value | pk |
+-----+------+---------+---------+------------+----+
| 0   | id   | INTEGER | 0       |            | 1  |
| 1   | name | TEXT    | 1       |            | 0  |
| 2   | age  | INTEGER | 1       |            | 0  |
+-----+------+---------+---------+------------+----+
sqlite>
```



## What CTAS do?

It will create a new table called `users_copy` that is a copy of the `users` table, with the same columns and data.

However, it just copies the data and structure but not the entire schema.

What I means is that if you have constraints and indexes, the select statement won't copy it over to the new table.

So, it will not copy the following;
- Constraints
- indexes
- Triggers

However, it will copy everything in the select statement's output.

So, you can add custom columns, computed columns, etc.

You can also filter the select statement with `WHERE` and `ORDER BY` statements in order to store the limited set of data that you wished to store.

## Copying only the structure

If you only wanted to only copy the structure and not the entire data, you can invalidate any rows from the select statement with the `WHERE` clause. This will just create columns in the new table. Giving you a fresh copy of the new table, without any data.

```sql
CREATE TABLE users_copy AS SELECT * FROM users WHERE 0;
```

This will create a new table called `users_copy` that is a copy of the `users` table, but without any data.
 
 ```sqlite

sqlite> CREATE TABLE users_copy AS SELECT * FROM users WHERE 0;
Parse error: table users_copy already exists
  CREATE TABLE users_copy AS SELECT * FROM users WHERE 0;
               ^--- error here
sqlite> CREATE TABLE IF NOT EXISTS users_copy AS SELECT * FROM users WHERE 0;
sqlite> SELECT * FROM users_copy;
+----+------+-----+
| id | name | age |
+----+------+-----+
| 1  | abc  | 20  |
| 2  | def  | 30  |
| 3  | ghi  | 40  |
+----+------+-----+
sqlite> DROP TABLE users_copy;
sqlite> CREATE TABLE IF NOT EXISTS users_copy AS SELECT * FROM users WHERE 0;
sqlite> SELECT * FROM users_copy;
sqlite> PRAGMA table_info(users_copy);
+-----+------+------+---------+------------+----+
| cid | name | type | notnull | dflt_value | pk |
+-----+------+------+---------+------------+----+
| 0   | id   | INT  | 0       |            | 0  |
| 1   | name | TEXT | 0       |            | 0  |
| 2   | age  | INT  | 0       |            | 0  |
+-----+------+------+---------+------------+----+
sqlite>
 ```
 As you can see, we need to actually drop the table in order to create a new table form scratch just like any normal table in sqlite.

 So, there are a few more gotchas and details of this, than I thought to be, will be taking a closer look at each one by one in the next few posts.

