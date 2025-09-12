---
type: sqlog
slug: sqlite-check-column-constraint
title: 'SQLite SQL: CHECK column constraint'
date: "2025-09-12"
tags: ["sqlite", "sql"]
---

## The CHECK column constraint

The [CHECK](https://sqlite.org/lang_createtable.html#check_constraints) clause is a column constraint that allows us to define certain conditions that we want to evaluate before inserting the column for the row and populate it accordingly. In this post, we will check, what this check is and when it is performed.

## What it checks

The `CHECK` constraint basically acts a validator, we define the condition and if the condition is not met (is false) then the row is not inserted or updated.

Let's take a simplest example with our good-ol users table:

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER CHECK(age >= 18)
);
```

Let's insert a user:

```sql
INSERT INTO users (name, age) VALUES ('John', 21);
```

```sql
SELECT * FROM users;
```
This rightly populates with the user `John` with age `21`.

If we try to insert a user with age less than 18, it will fail:

```sql
INSERT INTO users (name, age) VALUES ('Alice', 17);
```
This should fail the constraint as the age is less than 18.

```
sqlite> INSERT INTO users (name, age) VALUES ('Alice', 17);
Runtime error: CHECK constraint failed: age >= 18 (19)
sqlite>
```
As you would see, the insertion query failed as a `CHECK` constraint failed.

Let's add one more valid user:

```sql
INSERT INTO users (name, age) VALUES ('Bob', 20);
```

This is fine, and now, let's try to update the age of a user:

```sql
UPDATE users SET age = 17 WHERE id = 2;
```

This should fail the constraint as the age is less than 18.

```
sqlite> INSERT INTO users (name, age) VALUES ('Bob', 20);
sqlite> SELECT * FROM users;
+----+------+-----+
| id | name | age |
+----+------+-----+
| 1  | John | 21  |
| 2  | Bob  | 20  |
+----+------+-----+

sqlite> UPDATE users SET age = 17 where id = 2;
Runtime error: CHECK constraint failed: age >= 18 (19)
sqlite>
```

As you would see, the update query failed as a `CHECK` constraint failed.

## When it is checked?

As we saw it is checked before the row is inserted or updated.

## Disabling CHECK constraint

We can disable the `CHECK` constraint on all the tables in the database by setting the `PRAGMA` of `ignore_check_constraints` to `TRUE`:

```sql
PRAGMA ignore_check_constraints = TRUE;
```

This will disable all the `CHECK` constraints on all the tables in the database.

Let's try again updating the user's age with less then 18:

```sql
UPDATE users SET age = 17 WHERE id = 2;
```

This time, it readily updates the age to less then 18, without failing the constraint, which means the constraint was disable or shunned down.

```
sqlite> UPDATE users SET age = 17 where id = 2;
Runtime error: CHECK constraint failed: age >= 18 (19)
sqlite> PRAGMA ignore_check_constraints = TRUE;
sqlite> SELECT * FROM users;
+----+------+-----+
| id | name | age |
+----+------+-----+
| 1  | John | 21  |
| 2  | Bob  | 20  |
+----+------+-----+
sqlite> UPDATE users SET age = 17 where id = 2;
sqlite> SELECT * FROM users;
+----+------+-----+
| id | name | age |
+----+------+-----+
| 1  | John | 21  |
| 2  | Bob  | 17  |
+----+------+-----+
sqlite>
```

Similarly the inserts will be allowed if the constraint is disabled.

```sql
INSERT INTO users (name, age) VALUES ('Harry', 14);
```

```
sqlite> INSERT INTO users (name, age) VALUES ('Harry', 14);
sqlite> SELECT * FROM users;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | John  | 21  |
| 2  | Bob   | 17  |
| 3  | Harry | 14  |
+----+-------+-----+
sqlite>
```

## Table level check constraint

The `CHECK` constraint can be applied to a table too, in fact it doesn't really matter you add it to a column or a table, as it will be evaluated for insertion or updation of the column and it's not tied to a column.

Let's take a better example:

```sql
CREATE TABLE accounts (
    id INTEGER PRIMARY KEY,
    account_type TEXT NOT NULL,
    balance INTEGER NOT NULL,
    CHECK (
        (account_type = 'savings' AND balance >= 0)
        OR (account_type = 'loan' AND balance <= 0)
    )
);
```

Here, we have a `CHECK` constraint that checks if the account type is `savings` and balance is greater than or equal to 0, or if the account type is `loan` and balance is less than or equal to 0.

So we are combining two columns to form a condition and check if the row is actually valid or not to be inserted.

Let's insert a few rows:

```sql
INSERT INTO accounts (account_type, balance) VALUES ('savings', 1000);
INSERT INTO accounts (account_type, balance) VALUES ('loan', -1000);
```
This will insert readily the `savings` and the `loan` account as the constraint passes.

However if we try to do the opposite and make the constraint check fail.

```sql
INSERT INTO accounts (account_type, balance) VALUES ('savings', -1000);
```

```sql
INSERT INTO accounts (account_type, balance) VALUES ('loan', 1000);
```

It will fail as the check condition is not TRUE anymore for both the cases.


```
sqlite> INSERT INTO accounts (account_type, balance) VALUES ('savings', 1000);
sqlite> INSERT INTO accounts (account_type, balance) VALUES ('loan', -1000);
sqlite> INSERT INTO accounts (account_type, balance) VALUES ('loan', 1000);
Runtime error: CHECK constraint failed: (account_type = 'savings' AND balance >= 0)
        OR (account_type = 'loan' AND balance <= 0) (19)
sqlite> INSERT INTO accounts (account_type, balance) VALUES ('savings', -1000);
Runtime error: CHECK constraint failed: (account_type = 'savings' AND balance >= 0)
        OR (account_type = 'loan' AND balance <= 0) (19)
sqlite>
```
As you can see the `CHECK` condition failed for both the cases.


## Check constraint values

Note, here the value will be either `TRUE` or `FALSE`. But what about other values? NULL, 0, ?

That brings us to the fundamental of how `CHECK` considers the values as `TRUE` and `FALSE`.

From the documentation:

> If the result is zero (integer value 0 or real value 0.0), then a constraint violation has occurred.
> If the CHECK expression evaluates to NULL, or any other non-zero value, it is not a constraint violation


So, its not really `TRUE` or `FALSE` but:
- `TRUE`: is any non-zero value or `NULL` (if a condition is true, it is set as `1`)
- `FALSE`: is `0` (if a condition is false, it is set as `0`)

So, let's create a few tables with different `CHECK` constraints having different values for the check constraint:

```sql
CREATE TABLE test (
    val INTEGER,
    CHECK(val)
);
```
Now, let's insert some values

The below query will evaluate to false, hence failing the constraint.
```sql
INSERT INTO test(val) VALUES (0);
```

The value is `1` which is a non-zero value, hence passing the constraint.
```sql
INSERT INTO test(val) VALUES (1);
```

Here the value is `0.000` which is still zero, hence failing the constraint.
```sql
INSERT INTO test(val) VALUES (0.000);
```

The value here is `0.0001` which is a non-zero value, hence passing the constraint.
```sql
INSERT INTO test(val) VALUES (0.0001);
```

The `FALSE` is alias for `0` hence failing the constraint.
```sql
INSERT INTO test(val) VALUES (FALSE);
```

Similarly, `TRUE` is an alias for `1` hence passing the constraint.

```sql
INSERT INTO test(val) VALUES (TRUE);
```


The empty string is casted as numeric, it evaluates to `0`, hence failing the constraint.
```sql
INSERT INTO test(val) VALUES ('');
```

Similarly, the string `'a'` is casted as numeric, it evaluates to `0`, hence failing the constraint.
```sql
INSERT INTO test(val) VALUES ('a');
```

Here's the log, the SQLog :)

```sql
sqlite> CREATE TABLE test (
    val INTEGER,
    CHECK(val)
);
sqlite> INSERT INTO test(val) VALUES (0);
Runtime error: CHECK constraint failed: val (19)
sqlite> INSERT INTO test(val) VALUES (1);
sqlite> INSERT INTO test(val) VALUES (0.000);
Runtime error: CHECK constraint failed: val (19)
sqlite> INSERT INTO test(val) VALUES (0.0001);
sqlite> INSERT INTO test(val) VALUES (FALSE);
Runtime error: CHECK constraint failed: val (19)
sqlite> INSERT INTO test(val) VALUES (TRUE);
sqlite> INSERT INTO test(val) VALUES ('');
Runtime error: CHECK constraint failed: val (19)
sqlite> INSERT INTO test(val) VALUES ('a');
Runtime error: CHECK constraint failed: val (19)
sqlite> INSERT INTO test(val) VALUES (124);
sqlite> SELECT * FROM test;
+--------+
|  val   |
+--------+
| 1      |
| 0.0001 |
| 1      |
| 124    |
+--------+
sqlite>

sqlite> INSERT INTO test(val) VALUES (NULL);
sqlite> INSERT INTO test DEFAULT VALUES;
sqlite> SELECT * FROM test;
+--------+
|  val   |
+--------+
| 1      |
| 0.0001 |
| 1      |
| 124    |
|        |
|        |
+--------+
sqlite>
```

Now, you might be wondering why the heck `'a'` failed? Well, becuase casting strings in numeric form, will evaluate to `0`

```sql
SELECT CAST('a' AS NUMERIC);
```

A wired quirk but worth noting, so won't work with strings, need to cast or perform better checks for string related values.

```
sqlite> INSERT INTO test(val) VALUES ('a');
Runtime error: CHECK constraint failed: val (19)
sqlite> SELECT CAST('a' AS NUMERIC);
+----------------------+
| CAST('a' AS NUMERIC) |
+----------------------+
| 0                    |
+----------------------+
sqlite>
```

So that is it from the basic walkthrough of the `CHECK` constraint.

Some TLDRs:
1. The `CHECK` clause is a column constraint that allows us to define certain conditions that we want to evaluate on the column(s) before inserting/updating the row.
2. The constraint is table wide, as it boils down to inserting/updating or not inserting/updating the row.
3. The `CHECK` constraint is evaluated before the row is inserted/updated.
4. The value of the conditions is evaluated as a `NUMERIC` value so any value is `TRUE` if it is not `0`, even `NULL` is true, and `FALSE` if it is `0`.
5. The `CHECK` constraint can be disabled using the `PRAGMA` command `PRAGMA ignore_check_constraints = TRUE;` or `PRAGMA ignore_check_constraints = ON;` `ON`, `TRUE`, either works.
 
