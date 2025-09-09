---
type: sqlog
slug: sqlite-unique-column-constraint
title: 'SQLite SQL: UNIQUE column constraint'
date: "2025-09-09"
tags: ["sqlite", "sql"]
---

## Unique Constraint

The `UNIQUE` constraint ensures that a column in a table contains no duplicate values. Like, if you add an `UNIQUE` constraint to a column, then no two rows can have the exact/same value for that column.

The documentation is more nuanced.

> A `UNIQUE` constraint is similar to a `PRIMARY KEY` constraint, **except** that a single table may have any number of `UNIQUE` constraints. For each `UNIQUE` constraint on the table, each row must contain a unique combination of values in the columns identified by the `UNIQUE` constraint. For the purposes of `UNIQUE` constraints, `NULL` values are considered distinct from all other values, including other NULLs. As with `PRIMARY KEY`s, a `UNIQUE` table-constraint clause must contain only column names — the use of expressions in an indexed-column of a `UNIQUE` table-constraint is not supported.

There are three points to note, I'll simplify the points:
1. If you add an `UNIQUE` constraint to a column, then no two rows can have the exact/same value for that column in that table.
2. The `UNIQUE` constraint considers `NULL` values as distinct i.e. it cannot distinguish between two rows having `NULL` values, hence two rows will pass the `UNIQUE` constraint. (Better to use `UNIQUE` and `NOT NULL`)
3. The `UNIQUE` constraint is tied to one or more columns, so you cannot add custom expressions to the `UNIQUE` constraint. (Better to use `UNIQUE INDEX` to create custom expression-based uniqueness or index for querying data.)


## Adding a UNIQUE constraint

To add a unique constraint to a column, simply add the `UNIQUE` keyword to the column definition:

Let's keep the table simple with just three fields: `name`, `email`, and `age`.

```sql
CREATE TABLE users (
    name TEXT,
    email TEXT UNIQUE,
    age INTEGER CHECK (age > 0)
);
```

Now, lets insert a few records

```sql
INSERT INTO users (name, email, age) VALUES('alice', 'alice@wonderland.com', 12);
INSERT INTO users (name, email, age) VALUES('bob', 'bob@wonderland.com', 13);
```

Now, lets try to insert a record with the existing email value:

```sql
INSERT INTO users (name, email, age) VALUES('charlie', 'alice@wonderland.com', 14);
```

This will immediately fail the `UNIQUE` constraint on the email, and the record will not be inserted.

```
sqlite> INSERT INTO users (name, email, age) VALUES('charlie', 'alice@wonderland.com', 14);
Runtime error: UNIQUE constraint failed: users.email (19)
```

Notice, as the constraint is only on the email column, it's not on the name or age columns.

```sql
INSERT INTO users (name, email, age) VALUES('alice', 'charlie@wonderland.com', 12);

```

This will succeed, as the `UNIQUE` constraint is only on the email column. Hence, we can conclude that the `UNIQUE` constraint is on a column level when stated next to the column while creating the table.

```sql
SELECT * FROM users;
```

```
+-------+------------------------+-----+
| name  |         email          | age |
+-------+------------------------+-----+
| alice | alice@wonderland.com   | 12  |
| bob   | bob@wonderland.com     | 13  |
| alice | charlie@wonderland.com | 12  |
+-------+------------------------+-----+
sqlite>
```

Let's understand the problem first with `NULL` values in UNIQUE constraint.

## Adding a UNIQUE constraint with NULL values

What if I insert a record with a `NULL` value for the `email` column?

```sql
INSERT INTO users (name, email, age) VALUES('charlie', NULL, 14);
--OR
INSERT INTO users (name, age) VALUES('charlie', 14);
```

That went in, it created the record, which is fine, so far, as there was no previous entry with `NULL` email in the table, or rather the `email` column.

```sql
SELECT * FROM users;
```

However, what happens if we try to insert another record with a `NULL` value for the `email` column?

```sql
INSERT INTO users (name, email, age) VALUES('dave', NULL, 15);
-- OR
INSERT INTO users (name, age) VALUES('dave', 15);
```

This goes in as well, Hhh? It shouldn't have right? If `alice@wonderland.com` is not allowed twice then why should `NULL` be allowed?

That's what the documentation clearly stated:

> For the purposes of `UNIQUE` constraints, `NULL` values are considered distinct from all other values, including other NULLs.

That is clear as crystal, SQLite won't distinguish between two `NULL` values in the `email` column. Hence, we can't have duplicate `NULL` values in the `email` column.

For that we need to set email as `NOT NULL` as well as `UNIQUE` inorder to avoid populating `NULL` values in the `email` column.


## Adding UNIQUE Constraint on multiple columns

If you wondered, can we add unique constraint on multiple columns? Yes, we can and there are subtle variations on how you interpret those as, multiple as in two or more columns unique independently or the combination of two or more columns being unique.

1. Unique constraint independent of each other
2. Combinational Unique constraint

### Unique constraint independent of each other

Let's say we want to add constraint to the user table, which could have more than one column that needs to be unique, like example, the `phone_number` for two users cannot be same, and so on.

```sql
DROP TABLE users;
CREATE TABLE users (
    name TEXT,
    email TEXT UNIQUE,
    phone_number TEXT UNIQUE,
    government_id TEXT UNIQUE,
    age INTEGER CHECK (age > 0),
    bio TEXT
);
```
In the above table schema, we have three unique constraints, `email`, `phone_number`, and `government_id`. Each unique constraint is independent of each other. Let's take a look at a few cases in order to understand it better.

```sql
INSERT INTO users (name, email, phone_number, government_id, age, bio) 
    VALUES('alice', 'alice@wonderland.com', '1234567890', 'ABC123', 12, 'I am alice');
INSERT INTO users (name, email, phone_number, government_id, age, bio)
    VALUES('bob', 'bob@wonderland.com', '1234567891', 'DEF456', 13, 'I am bob');
```

```sql
SELECT * FROM users;
```

That is the simple case, However what would happen if one user has the same `phone_number`, and `government_id` but different `email`?

```sql
INSERT INTO users (name, email, phone_number, government_id, age, bio)
    VALUES('alice', 'new_alice@wonderland.com', '1234567890', 'ABC123', 12, 'I am alice');
```

```
sqlite> INSERT INTO users (name, email, phone_number, government_id, age, bio)
    VALUES('alice', 'new_alice@wonderland.com', '1234567890', 'ABC123', 12, 'I am alice');

Runtime error: UNIQUE constraint failed: users.government_id (19)
```

Nope, it is independent of each constraint, no two rows can have the same `government_id` or `phone_number` and even the `email` column.

Now, what if we want to allow `phone_number`, `government_id` and `email_id` to have a unique combination? Like the three of the columns together should be unique across the table, now with that they won't be unique across the single column, they are combined.

### Combinational Unique constraint

In a combinational unique constraint, we can add two or more columns like `UNIQUE(colum_1, column_2, ... column_n)` as a combinational constraint.

In this type of constraint:
- Not necessarily each column needs to be unique, the combination of the `N` columns should be unique

```sql
DROP TABLE users;
CREATE TABLE users (
    name TEXT,
    email TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    government_id TEXT NOT NULL,
    age INTEGER CHECK (age > 0),
    bio TEXT,
    UNIQUE(email, phone_number, government_id)
);

-- OR

-- CREATE TABLE users(
--     name TEXT,
--     email TEXT NOT NULL,
--     phone_number TEXT NOT NULL,
--     age INTEGER CHECK (age > 0),
--     bio TEXT,
--     UNIQUE(email, phone_number)
-- );
```

I am just making it up, its not secure, like combination of `email`, `phone_number` and `government_id` is not secure, you can debate about that, but it's  not a data-integrity and schema-design class, this post is about understanding the `UNIQUE` column constraint with combination of two or more columns.

```sql
INSERT INTO users (name, email, phone_number, government_id, age, bio) 
    VALUES('alice', 'alice@wonderland.com', '1234567890', 'ABC123', 12, 'I am alice');
INSERT INTO users (name, email, phone_number, government_id, age, bio)
    VALUES('bob', 'bob@wonderland.com', '1234567891', 'DEF456', 13, 'I am bob');
INSERT INTO users (name, email, phone_number, government_id, age, bio)
    VALUES('alice', 'new_alice@wonderland.com', '1234567890', 'ABC123', 12, 'I am alice');
```

```sql
SELECT * FROM users;
```

That works! We have `1234567890` as the `phone_number` and `ABC123` as the `government_id` which is same, but `email` is different. So, the combination of `email`, `phone_number` and `government_id` is unique.


```
sqlite> SELECT * FROM users;
+-------+--------------------------+--------------+---------------+-----+------------+
| name  |          email           | phone_number | government_id | age |    bio     |
+-------+--------------------------+--------------+---------------+-----+------------+
| alice | alice@wonderland.com     | 1234567890   | ABC123        | 12  | I am alice |
| bob   | bob@wonderland.com       | 1234567891   | DEF456        | 13  | I am bob   |
| alice | new_alice@wonderland.com | 1234567890   | ABC123        | 12  | I am alice |
+-------+--------------------------+--------------+---------------+-----+------------+
sqlite> INSERT INTO users (name, email, phone_number, government_id, age, bio)
    VALUES('bob', 'new_alice@wonderland.com', '0987654321', 'ABC123', 12, 'I am alice');
sqlite> SELECT * FROM users;
+-------+--------------------------+--------------+---------------+-----+------------+
| name  |          email           | phone_number | government_id | age |    bio     |
+-------+--------------------------+--------------+---------------+-----+------------+
| alice | alice@wonderland.com     | 1234567890   | ABC123        | 12  | I am alice |
| bob   | bob@wonderland.com       | 1234567891   | DEF456        | 13  | I am bob   |
| alice | new_alice@wonderland.com | 1234567890   | ABC123        | 12  | I am alice |
| bob   | new_alice@wonderland.com | 0987654321   | ABC123        | 12  | I am alice |
+-------+--------------------------+--------------+---------------+-----+------------+
sqlite>
```

```sql
INSERT INTO users (name, email, phone_number, government_id, age, bio)
    VALUES('alice', 'new_alice@wonderland.com', '1234567890', 'ABC123', 12, 'I am alice');
```

```
sqlite> INSERT INTO users (name, email, phone_number, government_id, age, bio)
    VALUES('bob', 'new_alice@wonderland.com', '0987654321', 'ABC123', 12, 'I am alice');
Runtime error: UNIQUE constraint failed: users.email, users.phone_number, users.government_id (19)
sqlite>
```

Now this will fail, as the `email`, `phone_number`, `government_id` combination is already unique.


## Primary key vs Unique Constraint

Both of these do the same thing, but `UNIQUE` can be added to multiple columns however, the `PRIMARY KEY` should only be one, but it can contain multiple columns. In other words, UNIQUE constraint is column-level, however each table can only have one `PRIMARY KEY` that will be used to indentify the uniqueness among the rows.

```
PRIMARY KEY = UNIQUE + NOT NULL
```
Primary key is also a shorthand for `UNIQUE` and `NOT NULL`, if you add a `PRIMARY KEY` to a column or combination of columns for a table, you needn't specify `UNIQUE` and `NOT NULL`, it already does that for you.

So, that is the general difference of `UNIQUE` and `PRIMARY KEY`, we'll explore the `PRIMARY KEY` in later posts.

