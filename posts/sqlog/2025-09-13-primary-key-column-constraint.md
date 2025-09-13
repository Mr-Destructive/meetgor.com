---
type: sqlog
slug: sqlite-primary-key-column-constraint
title: 'SQLite SQL: PRIMARY KEY column constraint'
date: "2025-09-13"
tags: ["sqlite", "sql"]
---

## The PRIMARY KEY column constraint

The PRIMARY KEY constraint is not actually a column constraint, it is a table constraint. But in this section, we'll only learn about how to handle single column-level `PRIMARY KEY`.

The [PRIMARY KEY](https://sqlite.org/lang_createtable.html#the_primary_key) column constraint is a constraint that ensures that a column contains unique values and is the `PRIMARY` way to distinguish between all the rows of that table.

From the documentation:

> Each table in SQLite may have at most one PRIMARY KEY. If the keywords PRIMARY KEY are added to a column definition, then the primary key for the table consists of that single column. Or, if a PRIMARY KEY clause is specified as a table-constraint, then the primary key of the table consists of the list of columns specified as part of the PRIMARY KEY clause. The PRIMARY KEY clause must contain only column names — the use of expressions in an indexed-column of a PRIMARY KEY is not supported. An error is raised if more than one PRIMARY KEY clause appears in a CREATE TABLE statement. The PRIMARY KEY is optional for ordinary tables but is required for WITHOUT ROWID tables.
> If a table has a single column primary key and the declared type of that column is "INTEGER" and the table is not a WITHOUT ROWID table, then the column is known as an INTEGER PRIMARY KEY. See below for a description of the special properties and behaviors associated with an INTEGER PRIMARY KEY.
> Each row in a table with a primary key must have a unique combination of values in its primary key columns. For the purposes of determining the uniqueness of primary key values, NULL values are considered distinct from all other values, including other NULLs. If an INSERT or UPDATE statement attempts to modify the table content so that two or more rows have identical primary key values, that is a constraint violation.
> According to the SQL standard, PRIMARY KEY should always imply NOT NULL. Unfortunately, due to a bug in some early versions, this is not the case in SQLite. Unless the column is an INTEGER PRIMARY KEY or the table is a WITHOUT ROWID table or a STRICT table or the column is declared NOT NULL, SQLite allows NULL values in a PRIMARY KEY column. SQLite could be fixed to conform to the standard, but doing so might break legacy applications. Hence, it has been decided to merely document the fact that SQLite allows NULLs in most PRIMARY KEY columns.

Here is what I think in short:

- Used to identify the rows of a table. Hence must be unique
- At most one primary key per table (default rowid, multiple primary key not allowed)
- Primary key can be a combination of multiple columns or even single column
- Only column names allowed (not expressions)
- Required for WITHOUT ROWID
- INTEGER PRIMARY KEY special behavior
- NULL behavior (SQLite quirk vs SQL standard)

## Creating a simple table

By default, if you don't  create a column in a table as a `PRIMARY KEY`, it is automatically assumed that `rowid` will be used as the unique identifier for that table.

Let's create a simple table with columns like `username` and `email`, non of which are `PRIMARY KEY` explicitly mentioned.

```sql
CREATE TABLE users (
    username TEXT NOT NULL,
    email TEXT NOT NULL
);
```
Now, let's insert a few columns and check what we can observe.

```sql
INSERT INTO users(username, email) VALUES('jim', 'jim@abc.com');
INSERT INTO users(username, email) VALUES('jack', 'jack@abc.com');
SELECT * FROM users;
```

As, you can see, we have populated the table with `username` and `email` columns.

But wait, there should be `rowid` right? It is there but its a hidden column (do not touch it). 

```sql
SELECT rowid, * FROM users;
```
You can see the `rowid` column in the output.

```
sqlite> CREATE TABLE users (
    username TEXT NOT NULL,
    email TEXT NOT NULL
);
sqlite> .mode table
sqlite> INSERT INTO users(username, email) VALUES('jim', 'jim@abc.com');
sqlite> INSERT INTO users(username, email) VALUES('jack', 'jack@abc.com');
sqlite> SELECT * FROM users;
+----------+--------------+
| username |    email     |
+----------+--------------+
| jim      | jim@abc.com  |
| jack     | jack@abc.com |
+----------+--------------+
sqlite> SELECT rowid, * FROM users;
+-------+----------+--------------+
| rowid | username |    email     |
+-------+----------+--------------+
| 1     | jim      | jim@abc.com  |
| 2     | jack     | jack@abc.com |
+-------+----------+--------------+
sqlite>
```

The rowid is used as a primary key for the table, this means it will take the max of the current state of the table and use the next number as its rowid.

We can verify that if we try to insert a duplicate value/record/row in the `users` table:

Let's insert the `jack` as the user and check if it inserts it, as it would have errored if the username or the email was the primary key, as this value is duplicated. But the `username` and `email` are not the `PRIMARY KEY` in this table, hence it will happily insert it.

```sql
INSERT INTO users(username, email) VALUES('jack', 'jack@abc.com');
SELECT rowid, * FROM users;
```

As, you can see it inserted the value and the `rowid` is now `3`.

```
sqlite> SELECT rowid, * FROM users;
+-------+----------+--------------+
| rowid | username |    email     |
+-------+----------+--------------+
| 1     | jim      | jim@abc.com  |
| 2     | jack     | jack@abc.com |
+-------+----------+--------------+
sqlite> INSERT INTO users(username, email) VALUES('jack', 'jack@abc.com');
sqlite> SELECT rowid, * FROM users;
+-------+----------+--------------+
| rowid | username |    email     |
+-------+----------+--------------+
| 1     | jim      | jim@abc.com  |
| 2     | jack     | jack@abc.com |
| 3     | jack     | jack@abc.com |
+-------+----------+--------------+
sqlite>
```

What would happen if the middle values are deleted?

let's delete the row with `rowid` `2`:

```sql
DELETE FROM users WHERE rowid = 2;
```

```sql
SELECT rowid, * FROM users;
```

As, you can see, the `rowid` `2` is deleted. The `rowid` will still take the max of the current state of the table, and will use the next number as the `rowid` which in this state should be `4`.

Let's insert a new record and check if the `rowid` is the max of the current state of the table.

```sql
INSERT INTO users(username, email) VALUES('jill', 'jill@abc.com');
```

```sql
SELECT rowid, * FROM users;
```

As you can see, the `rowid` is now `4`.

Let's delete that and insert one more.

```sql
DELETE FROM users WHERE rowid = 4;
```

```sql
SELECT rowid, * FROM users;
```

Now, that the max row of the table is `3` but we had `4` earlier, we need to check what could be the next rowid?

```sql
INSERT INTO users(username, email) VALUES('jill', 'jill@abc.com');
SELECT rowid, * FROM users;
```

Ok, as you can see, the `rowid` is now `4`. It is reusing the previous `rowid` value, which might be an issue, very rarely but this is a expected behaviour.

Which gives us the answer that there is maybe a quirk or a limitation of `rowid` that it can insert values which are max of the rowid in the current state.

```
sqlite> SELECT rowid, * FROM users;
+-------+----------+--------------+
| rowid | username |    email     |
+-------+----------+--------------+
| 1     | jim      | jim@abc.com  |
| 2     | jack     | jack@abc.com |
| 3     | jack     | jack@abc.com |
+-------+----------+--------------+
sqlite> DELETE FROM users WHERE rowid = 2;
sqlite> SELECT rowid, * FROM users;
+-------+----------+--------------+
| rowid | username |    email     |
+-------+----------+--------------+
| 1     | jim      | jim@abc.com  |
| 3     | jack     | jack@abc.com |
+-------+----------+--------------+
sqlite> INSERT INTO users(username, email) VALUES('jill', 'jill@abc.com');
sqlite> SELECT rowid, * FROM users;
+-------+----------+--------------+
| rowid | username |    email     |
+-------+----------+--------------+
| 1     | jim      | jim@abc.com  |
| 3     | jack     | jack@abc.com |
| 4     | jill     | jill@abc.com |
+-------+----------+--------------+
sqlite> DELETE FROM users WHERE rowid = 4;
sqlite> SELECT rowid, * FROM users;
+-------+----------+--------------+
| rowid | username |    email     |
+-------+----------+--------------+
| 1     | jim      | jim@abc.com  |
| 3     | jack     | jack@abc.com |
+-------+----------+--------------+
sqlite> INSERT INTO users(username, email) VALUES('joe', 'joe@abc.com');
sqlite> SELECT rowid, * FROM users;
+-------+----------+--------------+
| rowid | username |    email     |
+-------+----------+--------------+
| 1     | jim      | jim@abc.com  |
| 3     | jack     | jack@abc.com |
| 4     | joe      | joe@abc.com  |
+-------+----------+--------------+
sqlite>
```

Ok, that was a wired tangent, but worth noting.
We will check how to fix it with the `AUTOINCREMENT` clause in the future posts.

Now, let's create a `PRIMARY KEY` column that we can safely call a primary key.

## Creating a column as primary key

To create a column as a `PRIMARY KEY` we need to simply add the `PRIMARY KEY` constraint to the column. 

Let's create a simple table with columns like `username` and `email`, and here `email` is the primary key.

```sql
CREATE TABLE users (
    username TEXT NOT NULL,
    email TEXT PRIMARY KEY
);
```

Mentioning `PRIMARY KEY` with email means that each record needs to have a unique `email`.

If you attach a `PRIMARY KEY` to a column, the `UNIQUE` constraint is also added automatically. So, saying `email TEXT PRIMARY KEY UNIQUE` is redundant.

Let's try inserting default values to the users.

```sql
INSERT INTO users DEFAULT VALUES;
```
As you can see, this will fail as the `username` is a `NOT NULL` column.

OK, let's just satisfy that `NOT NULL` constraint for username.

```sql
INSERT INTO users(username) VALUES('james');
```
OK, that inserted the value.

```sql
SELECT * FROM users;
```

But, the email is empty! Becuase we didn't specify it? Yes, it is a text field, and a primary key, so you need to handle that. What about the rowid?

```sql
SELECT rowid, * FROM users;
```

As you can see, the rowid is `1` and email is empty.

That's worth noting that `rowid` will be populated independently of the `PRIMARY KEY` column. But if speicifed as `WITHOUT ROWID` in the table creation, it won't be there.

Now, let's insert the same record again. We have said the `email` is a `PRIMARY KEY`, so it will be unique right? RIGHT?

```sql
INSERT INTO users(username) VALUES('james');
```

Surprisingly, it inserted that too.

```sql
SELECT * FROM users;
```
The email is empty for both the rows, but the username is same, so the `email` is NULL, and you guessed it right, `NULL` is identified as `UNIQUE`. That is a bad schema design. The users table had a text key as a primary key, and we need to add `NOT NULL` to the `PRIMARY KEY` column.

```
sqlite> CREATE TABLE users (
    username TEXT NOT NULL,
    email TEXT PRIMARY KEY
);
sqlite> INSERT INTO users DEFAULT VALUES;
Runtime error: NOT NULL constraint failed: users.username (19)
sqlite> INSERT INTO users(username) VALUES('james');
sqlite> SELECT * FROM users;
+----------+-------+
| username | email |
+----------+-------+
| james    |       |
+----------+-------+
sqlite> INSERT INTO users(username) VALUES('james');
sqlite> SELECT * FROM users;
+----------+-------+
| username | email |
+----------+-------+
| james    |       |
| james    |       |
+----------+-------+
sqlite> INSERT INTO users(username) VALUES('jack');
sqlite> SELECT * FROM users;
+----------+-------+
| username | email |
+----------+-------+
| james    |       |
| james    |       |
| jack     |       |
+----------+-------+
sqlite>
```

> NOTE: NOT NULL is only added automatically for `INTEGER PRIMARY KEY` column, which means the column becomes an alias for the `rowid` table.
 
## PRIMARY KEY with NOT NULL

It turns out that `NOT NULL` is only added automatically for `INTEGER PRIMARY KEY` column, so for rest of the combination of `PRIMARY KEY` columns we need to explicitly add the `NOT NULL` constraint. 

So, let's add that:

```sql
DROP TABLE users;
CREATE TABLE users (
    username TEXT NOT NULL,
    email TEXT PRIMARY KEY NOT NULL
);
```

Now, let's try inserting default values to the users.

```sql
INSERT INTO users DEFAULT VALUES;
```

As you can see, this will fail as the `username` is a `NOT NULL` column. As we knew in the first attempt.

```sql
INSERT INTO users(username) VALUES('james');
```
OK, that is what we expected, it should fail, as have not specified the `email` column, as that is a primary key and it cannot be null. I mean it could be if we wanted to, but that is not expected to behave.

```sql
INSERT INTO users(username, email) VALUES('james', 'james@abc.com');
```
Ok, this should insert it successfully. 

```sql
SELECT rowid, * FROM users;
```
As you can see, we have populated the table with `username` and `email` columns. This is the perfect thing to setup.

The `rowid` will still be there as discussed, the `email` now is the `PRIMARY KEY` which means any duplicate string email will be constraint violation.

```sql
INSERT INTO users(username, email) VALUES('james', 'james@abc.com');
```
As you can see, this will fail, as we have already inserted a record with the same email. This is expected as we noted that the `PRIMARY KEY` automatically adds the `UNIQUE` constraint to the column.

Similary, we can check a few other things.

Like if we try to insert a record with the same username but different email:

```sql
INSERT INTO users(username, email) VALUES('james', 'james_new@abc.com');
```

This will succeed, as we have not specified the `username` as primary key, and it is ok to not have unique values for it.


```
sqlite> CREATE TABLE users (
    username TEXT NOT NULL,
    email TEXT PRIMARY KEY NOT NULL
);

sqlite> INSERT INTO users DEFAULT VALUES;
Runtime error: NOT NULL constraint failed: users.username (19)

sqlite> INSERT INTO users(username) VALUES('james');
Runtime error: NOT NULL constraint failed: users.email (19)

sqlite> INSERT INTO users(username, email) VALUES('james', 'james@abc.com');

sqlite> SELECT rowid, * FROM users;
+-------+----------+---------------+
| rowid | username |     email     |
+-------+----------+---------------+
| 1     | james    | james@abc.com |
+-------+----------+---------------+

sqlite> INSERT INTO users(username, email) VALUES('james', 'james@abc.com');
Runtime error: UNIQUE constraint failed: users.email (19)

sqlite> INSERT INTO users(username, email) VALUES('james', 'james_new@abc.com');

sqlite> SELECT rowid, * FROM users;
+-------+----------+-------------------+
| rowid | username |       email       |
+-------+----------+-------------------+
| 1     | james    | james@abc.com     |
| 2     | james    | james_new@abc.com |
+-------+----------+-------------------+

sqlite> INSERT INTO users(username, email) VALUES('jill', 'james_new@abc.com');
Runtime error: UNIQUE constraint failed: users.email (19)

sqlite> INSERT INTO users(username, email) VALUES('jill', 'jill@abc.com');

sqlite> SELECT rowid, * FROM users;
+-------+----------+-------------------+
| rowid | username |       email       |
+-------+----------+-------------------+
| 1     | james    | james@abc.com     |
| 2     | james    | james_new@abc.com |
| 3     | jill     | jill@abc.com      |
+-------+----------+-------------------+
sqlite>
```

That is how we primary key works, atleast the basic of it.

## INTEGER PRIMARY KEY

This is a special case where the column defined as `PRIMARY KEY` with INTEGER type becomes an alias to the `rowid` and hence we will have both `UNIQUE` and `NOT NULL` constraints added to it.

Let's create a table with `INTEGER PRIMARY KEY` column.

```sql
CREATE TABLE users (
    user_id INTEGER PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL
);
```
In this table `users` we have `user_id` which is `INTEGER PRIMARY KEY` column. It is an alias for the `rowid` table as we'll see.

Now, let's insert some data and check what we can observe.

```sql
INSERT INTO users DEFAULT VALUES;
```
Obivously, this will fail, as we have not specified the `username` and `email` as `NOT NULL`, and we haven't specified in the insertion.

Let's specify both and check if it can violate any constraints.

```sql
INSERT INTO users(username, email) VALUES('james', 'james@abc.com');
```

As you can see, we successfully inserted the record, as we have specified the `username` and `email` as `NOT NULL` and the primary key which is `user_id` and it will automatically get populated. The `user_id` will be automatically populated as it will be a alias to the `rowid` column, and we know it will populate automatically for each inserted value for each row, by incrementing the max row id at that point.

```sql
SELECT * FROM users;
```

As you can see, the `user_id` is `1` and `username` is `james` and `email` is `james@abc.com`. But what is `rowid`?

Let's check that too.

```sql
SELECT rowid, * FROM users;
```

As you can see, the `user_id` is `1` but the rowid has become the `user_id` which is basically an alias for the `rowid`.

Let's see if we can add more records.

Let's try adding the same record again.

```sql
INSERT INTO users(username, email) VALUES('james', 'james@abc.com');
```
Interestingly, this works, because we have not specified the `user_id` as `PRIMARY KEY` and inserting duplicate username and email should not violate any constraints unless we have specified as `UNIQUE` or others.

```
sqlite> CREATE TABLE users (
    user_id INTEGER PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL
);
sqlite> INSERT INTO users DEFAULT VALUES;
Runtime error: NOT NULL constraint failed: users.username (19)
sqlite> INSERT INTO users(username) VALUES('james');
Runtime error: NOT NULL constraint failed: users.email (19)
sqlite> INSERT INTO users(username, email) VALUES('james', 'james@abc.com');
sqlite> SELECT * FROM users;
+---------+----------+---------------+
| user_id | username |     email     |
+---------+----------+---------------+
| 1       | james    | james@abc.com |
+---------+----------+---------------+
sqlite> SELECT rowid, * FROM users;
+---------+---------+----------+---------------+
| user_id | user_id | username |     email     |
+---------+---------+----------+---------------+
| 1       | 1       | james    | james@abc.com |
+---------+---------+----------+---------------+
sqlite> INSERT INTO users(username, email) VALUES('james', 'james@abc.com');
sqlite> SELECT rowid, * FROM users;
+---------+---------+----------+---------------+
| user_id | user_id | username |     email     |
+---------+---------+----------+---------------+
| 1       | 1       | james    | james@abc.com |
| 2       | 2       | james    | james@abc.com |
+---------+---------+----------+---------------+
sqlite> INSERT INTO users(username, email) VALUES('jill', 'jill@abc.com');
sqlite> SELECT rowid, * FROM users;
+---------+---------+----------+---------------+
| user_id | user_id | username |     email     |
+---------+---------+----------+---------------+
| 1       | 1       | james    | james@abc.com |
| 2       | 2       | james    | james@abc.com |
| 3       | 3       | jill     | jill@abc.com  |
+---------+---------+----------+---------------+
sqlite>
```
So that is how `INTEGER PRIMARY KEY` works.

You can add `PRIMARY KEY` to any column, but you need to ensure that it doesn't violate the constraint of allowing duplicate enties.

We know that the `PRIMARY KEY` is not a column specific constraint. It is table wide constraint, we'll check out the table wide constraint that is combinational constraint in the next section.
