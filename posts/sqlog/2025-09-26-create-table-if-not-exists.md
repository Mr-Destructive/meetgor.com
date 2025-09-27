---
type: sqlog
slug: sqlite-create-table-if-not-exists
title: "SQLite SQL: Create Table If Not Exists"
date: "2025-09-26"
tags: ["sqlite", "sql"]
---

## Creating table if not already exists

The `CREATE TABLE` has one clause that we can add to create table if it doesn't exist already. So this comes handy if you already have created a table and instead of throwing an error it simply gracefully handles the query and doesn't re-create the table. 

Let's first creat a table, as usual, we'll use the goodol `users` table.

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);
```
Let's insert some data into the table.

```sql
INSERT INTO users (name) VALUES ('abc'), ('def');
```

Ok, the table `users` exist and also has some data in it.

```sql
SELECT * FROM users;
```

Now, if we try to create the same table again, it will throw an error.

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);
```

It will result in `users` table already exists error.

You cannot re-create a table that already exist, you need to drop the table and create again, or alter any columns you wanted to if that is possible.

Or, if you just want to create a table only if it doesn't already exists, then you can add the `IF NOT EXISTS` clause:

```sql
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);
```

Now, if we try to create the table again, it will not throw any error.

However, it won't re-create the `users` table, it will just skip the creation.

The table would be as is.

```sqlite

sqlite> CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);
sqlite> INSERT INTO users (name) VALUES ('abc'), ('def');
sqlite> .mode table
sqlite> SELECT * FROM users;
+----+------+
| id | name |
+----+------+
| 1  | abc  |
| 2  | def  |
+----+------+
sqlite> CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);
Parse error: table users already exists
  CREATE TABLE users (     id INTEGER PRIMARY KEY,     name TEXT NOT NULL,     e
               ^--- error here
sqlite> CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    password TEXT NOT NULL
);
sqlite> SELECT * FROM users;
+----+------+
| id | name |
+----+------+
| 1  | abc  |
| 2  | def  |
+----+------+
sqlite> .schema users
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);
sqlite>
```

## Where can we use this?

- Accidental cases: Avoiding creating a table that already exists.
- Code Generation tools: If you are using schema-based code generation tools like [sqlc](https://github.com/sqlc-dev/sqlc), you might keep on re-running the generation code again and again, to avoid the error statements, you can use this to stop table from re-creating or erroing queries.

