---
type: sqlog
slug: sqlite-foreign-key-table-constraint
title: 'SQLite SQL: Foreign KEY table constraint'
date: "2025-09-15"
tags: ["sqlite", "sql"]
---

## The Foreign KEY Table constraint

Foreign Keys are the fundamentals of any relational databases, as they are the ones that make the connection or the relations among the tables in our database system.

> Foreign key as the name suggest, this is a key referencing or pointing to a foreign (other) table, and that key could be a primary key for that table, hence referred to as key.

So, with a foreign key we can connect the data/records/row from other table to the table in which the foreign key is linked.

Think of it like a string(rope) attaching one record to the other. It is a link between two tables.

Now, how we define the `FOREIGN KEY` constraint is what the rope will be tied to and how.

Let's take a look at how the syntax of defining a `FOREIGN KEY` constraint looks like in SQLite.

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    content TEXT NOT NULL,
    author_id INTEGER,
    FOREIGN KEY (author_id) REFERENCES users (id)
);
```

You will need atleast two tables in order to map a foreign key constraint. In the above example, we have created a table `users` that will form the base table, and the table `posts` is what is linking the `id` from the `users` table to itself. 

This means one post can link to one user, we are using that reference in the current `posts` table with the alias of `author_id`. This way the `author_id` is the rope (string whatever you prefer saying, I think developers get confused when I say string) that connects a record in the post table to the record in the users table.

We can conclude a few things from this:
- We need to define what that `author_id` is for, for each record we insert into the `posts` table.
- If we don't define the `author_id` then will it be automatically populated? Nope!
- The `FOREIGN KEY` at its core is just like any other column, its just that how we set that column is what is important for the linking part.
- It depends on how the actual key i.e. the column in the foreign key is defined in the foreign table.

Might still not make sense, let me explain with a few example.

Let's insert a bunch of users.

```sql
INSERT INTO users(name) VALUES('harry'), ('ron'), ('malfoy');
```

```sql
SELECT * FROM users;
```

This has 3 records each assigned ids from 1 to 3.

Now, let's insert a bunch of posts.

```sql
INSERT INTO posts(content) VALUES ('hi, I am who?');
```

```sql
SELECT * FROM posts;
```

That `author_id` is `NULL` because we didn't inserted anythign against it.

```
sqlite> CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    content TEXT NOT NULL,
    author_id INTEGER,
    FOREIGN KEY (author_id) REFERENCES users (id)
);
sqlite> INSERT INTO users(name) VALUES('harry'), ('ron'), ('malfoy');
sqlite> .mode table
sqlite> SELECT * FROM users;
+----+--------+
| id |  name  |
+----+--------+
| 1  | harry  |
| 2  | ron    |
| 3  | malfoy |
+----+--------+
sqlite> INSERT INTO posts(content) VALUES ('hi, I am who?');
sqlite> SELECT * FROM posts;
+----+---------------+-----------+
| id |    content    | author_id |
+----+---------------+-----------+
| 1  | hi, I am who? |           |
+----+---------------+-----------+
sqlite> INSERT INTO users(name, id) VALUES('neville', NULL);
sqlite> SELECT * FROM posts;
+----+---------------+-----------+
| id |    content    | author_id |
+----+---------------+-----------+
| 1  | hi, I am who? |           |
+----+---------------+-----------+
sqlite> SELECT * FROM users;
+----+---------+
| id |  name   |
+----+---------+
| 1  | harry   |
| 2  | ron     |
| 3  | malfoy  |
| 4  | neville |
+----+---------+
sqlite>
```

We need to be careful while inserting anything in that `author_id` column, becuase think for a moment:
What do we insert in the `author_id` column?

Well, we need to have a valid `users(id)` right, that is the id in the users table, that is what we are referencing by that `FOREIGN KEY` constraint.

There are other questions that should arise now, like:

- What happens if we link something that doesn't exist in the foreign table? (in this case the `users` table)
- What happens if we what we have linked is deleted or changed in the foreign table?
- What happend if multiple rows refer the same id in the foreign table?

The `FOREIGN KEY` constraint will make sure we are linking the valid id and hence if we try to insert any invalid (non-existent) id of the foreign table (here the `users` table) we will fail this constraint and hence the record/row won't be inserted or updated.

As I said, this is a rope, the one end of the rope is always the current table record and the other one is what you attach it when inserting or updating a record.

So, for starters, we'll insert the users with the id that exist in the first place.

In our case, we have 4 authors created with the id from 1 to 4 and names as `harry`, `ron`, `malfoy` and `neville` respectively.

Now, let's insert a post with the `author_id` as 1, that is the id of the user with the name `harry` which will be the author of the post.

```sql
INSERT INTO posts(content, author_id) VALUES ('hi, I am harry', 1);
```

This will set the `author_id` column of the `posts` table to 1, which is the id of the `harry` user.

If you want to see some better results, just run some of these to get it clear:

```sql
SELECT * FROM posts;
```

Get both tables data:

```sql
SELECT * FROM posts JOIN users ON author_id = id;
```

Make column names clear and remove redundant columns:

```sql
SELECT posts.*, users.name FROM posts JOIN users ON posts.author_id = users.id;
```



```
sqlite> INSERT INTO posts(content, author_id) VALUES ('hi, I am harry', 1);
sqlite> SELECT * FROM posts;
+----+----------------+-----------+
| id |    content     | author_id |
+----+----------------+-----------+
| 1  | hi, I am who?  |           |
| 2  | hi, I am harry | 1         |
+----+----------------+-----------+
sqlite> SELECT * FROM posts JOIN users ON author_id = id;
Parse error: ambiguous column name: id
sqlite> SELECT * FROM posts JOIN users ON posts.author_id = users.id;
+----+----------------+-----------+----+-------+
| id |    content     | author_id | id | name  |
+----+----------------+-----------+----+-------+
| 2  | hi, I am harry | 1         | 1  | harry |
+----+----------------+-----------+----+-------+
sqlite> SELECT posts.*, users.name FROM posts JOIN users ON posts.author_id = users.id;
+----+----------------+-----------+-------+
| id |    content     | author_id | name  |
+----+----------------+-----------+-------+
| 2  | hi, I am harry | 1         | harry |
+----+----------------+-----------+-------+
sqlite> SELECT posts.*, users.name AS author FROM posts JOIN users ON posts.author_id = users.id;
+----+----------------+-----------+--------+
| id |    content     | author_id | author |
+----+----------------+-----------+--------+
| 2  | hi, I am harry | 1         | harry  |
+----+----------------+-----------+--------+
sqlite>
```

Now, try inserting a post with the `author_id` which doesn't exist in the `users` table i.e. the id of the users above 4.

```sql
INSERT INTO posts(content, author_id) VALUES ('hi, I am hermoine', 5);
```

That inserted successfully? How can that be allowed? The user with id `5` doesn't exist and we are referencing the author (users) that doesn't exist.

## Pragma Foreign keys

Well, we have been lied to all along, the SQLite is so flexible by default, you need to nudge it inorder to be a little strict:

We need to enable the foreign key constraint checking by setting the `PRAGMA foreign_keys=on;`

```
PRAGMA foreign_keys=on;
```
This setting is off by default due to backward compatibility reasons.

You can check if the foreign key constraint checking is enabled by running the following command:

```
PRAGMA foreign_keys;
```

If it returns `1` then its enabled, or `0` if not.

Now, now if we run the query where the author id is `5` which means the `user` referenced by that id doesn't exist, we will get the following error:

```sql
INSERT INTO posts(content, author_id) VALUES ('hi, I am hermoine', 5);
```

> Runtime error: FOREIGN KEY constraint failed

Exactly as we expected.

```
sqlite> INSERT INTO posts(content, author_id) VALUES ('hi, I am hermoine', 5);
sqlite> SELECT posts.*, users.name AS author FROM posts JOIN users ON posts.author_id = users.id;
+----+----------------+-----------+--------+
| id |    content     | author_id | author |
+----+----------------+-----------+--------+
| 2  | hi, I am harry | 1         | harry  |
+----+----------------+-----------+--------+
sqlite> SELECT * FROM posts;
+----+-------------------+-----------+
| id |      content      | author_id |
+----+-------------------+-----------+
| 1  | hi, I am who?     |           |
| 2  | hi, I am harry    | 1         |
| 3  | hi, I am hermoine | 5         |
+----+-------------------+-----------+
sqlite> PRAGMA foreign_keys=on
   ...> ;
sqlite> INSERT INTO posts(content, author_id) VALUES ('hi, I am hermoine', 5);
Runtime error: FOREIGN KEY constraint failed (19)
sqlite>
```
However if we insert the `posts` record with a valid `author_id` then:

```sql
INSERT INTO posts(content, author_id) VALUES ('hi, I am ron', 2);
```
This will readily insert the record, since the `author_id` of 2 exists in the `users` table.

```sql
SELECT * FROM posts;
```
Neat.

```
sqlite> INSERT INTO posts(content, author_id) VALUES ('hi, I am ron', 2);
sqlite> SELECT * FROM posts;
+----+-------------------+-----------+
| id |      content      | author_id |
+----+-------------------+-----------+
| 1  | hi, I am who?     |           |
| 2  | hi, I am harry    | 1         |
| 3  | hi, I am hermoine | 5         |
| 4  | hi, I am ron      | 2         |
+----+-------------------+-----------+
sqlite>
```

Now, back to those three questions:

- What happens if we link something that doesn't exist in the foreign table? (in this case the `users` table)
- What happens if we what we have linked is deleted or changed in the foreign table?
- What happend if multiple rows refer the same id in the foreign table?


Let's tackle one by one 

> NOTE: We'll assume PRAGMA foreign_keys=on; from here on

## Invalid Foreign Key

We already covered this in the previous discussion, if we link something that doesn't exist in the foreign table (in this case the `users` table) then we will fail the `FOREIGN KEY` constraint and hence the record/row won't be inserted or updated. But for this to fail we need to make sure the `PRAGMA foreign_keys` is on or toggled on.

## Deleted or Updated Foreign Key

What happens, if we inserted a record to the primary table (here the `posts` table) the `author_id` was existing and valid at the time of insertion, however after a while the `user` got deleted, what happens to the record that the `posts` record/row  that still references the deleted user id that is the `author_id` alias.


Enter `ON DELETE` and `ON UPDATE` clause on the `FOREIGN KEY` constraint.

It turns out, we can define the behaviour of the `FOREIGN KEY` constraint when the referenced row is deleted or updated in the foreign table (here the `users` table).

So, we need to define, what happens to the record/row in the primary table (here the `posts` table) when the referenced row in the foreign table (here the `users` table) is deleted or updated.

There are 5 options to choose from:
1. `NO ACTION`: Do nothing (default)
2. `RESTRICT`: Stop the foreign key table from deleting the record/key in the foreign table
3. `SET NULL`: Set the primary table's record foreign key to `NULL`
4. `SET DEFAULT`: Set the primary table's record foreign key to the default value.
5. `CASCADE`: Delete the primary table's record/row if the referenced row in the foreign table is deleted.

Phew! This opens a lot of options to play with:

We already know the `NO ACTION`, it will just let it happen, nothing will be done. It's not recommended though, but SQLite is flexible, how many times I have to say it? Double edged sword.

### Restrict

Let's drop the `posts` table and start a fresh with the `ON DELETE` and `ON UPDATE` as `RESTRICT` option on the `FOREIGN KEY`  constraint.

#### Restrict on delete

We can restrict the deletion of the foreign records referred by the primary record. We use the `ON DELETE RESTRICT` as the option on the `FOREIGN KEY` constraint.

```sql
DROP TABLE posts;

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    content TEXT NOT NULL,
    author_id INTEGER,
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE RESTRICT
);
```
This will now restrict the deletion of the foreign records referred by the primary record. In this case, if a record/row in `users` table which is refereced in the `posts` table is deleted, it will prevent it, it won't be deleted. However, any other record in the `users` table can be deleted.

Let's first insert a few rows into the `posts` table since we dropped it.

```sql
INSERT INTO posts(content, author_id) VALUES ('hi, I am harry', 1);
INSERT INTO posts(content, author_id) VALUES ('hi, I am ron', 2);
```

These are valid records, since the `author_id` of 1 and 2 exist in the `users` table.

However the following query will fail, since the `author_id` of 5 doesn't exist in the `users` table. We already know this:

```sql
INSERT INTO posts(content, author_id) VALUES ('hi, I am hermoine', 5);
```

Now, Let's look at all the `posts` and their authors

```sql
SELECT * FROM posts;
```

This is the list of `users` that we have:

```sql
SELECT * FROM users;
```

Now, let's try deleting a record in the `users` table.

We know that `users` with id `1` and `2` are referenced in the `posts` table.

```sql
DELETE FROM users WHERE id = 2;
```

This will fail, since we are trying to delete a record that is referenced in the `posts` table.

We can check if the `users` with `id` 2 is still there:

```sql
SELECT * FROM users;
```
It is indeed ther, it restricted the deletion.

However if we delete the `users` where id is `3` which is not referenced in the `posts` table, it will succeed.

```sql
DELETE FROM users WHERE id = 3;
```
This deleted the `users` with `id` 3.

```sql
SELECT * FROM users;
```
As you can see the dirt `malfoy` was deleted.

```
sqlite> DROP TABLE posts;

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    content TEXT NOT NULL,
    author_id INTEGER,
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE RESTRICT
);
sqlite> SELECT * FROM posts;
sqlite> INSERT INTO posts(content, author_id) VALUES ('hi, I am harry', 1);
sqlite> INSERT INTO posts(content, author_id) VALUES ('hi, I am ron', 2);
sqlite> SELECT * FROM posts;
+----+----------------+-----------+
| id |    content     | author_id |
+----+----------------+-----------+
| 1  | hi, I am harry | 1         |
| 2  | hi, I am ron   | 2         |
+----+----------------+-----------+
sqlite> INSERT INTO posts(content, author_id) VALUES ('hi, I am hermoine', 5);
Runtime error: FOREIGN KEY constraint failed (19)
sqlite> SELECT * FROM users;
+----+---------+
| id |  name   |
+----+---------+
| 1  | harry   |
| 2  | ron     |
| 3  | malfoy  |
| 4  | neville |
+----+---------+
sqlite> DELETE FROM users where id = 2;
Runtime error: FOREIGN KEY constraint failed (19)
sqlite> DELETE FROM users where id = 3;
sqlite> SELECT * FROM users;
+----+---------+
| id |  name   |
+----+---------+
| 1  | harry   |
| 2  | ron     |
| 4  | neville |
+----+---------+
sqlite>
```

What about updates?

#### Restrict on update

What if we update the users record will it restrict it, nope, unless we set `ON UPDATE` to `RESTRICT`:

> This option of `UPDATE ON` is usually not needed, as this is referring to the `UPDATE ON` the `FOREIGN KEY` and not on the entire foreign table columns, so only if you update the foreign key it will restrict or any action will be performed.

Which means, if you update the `name` of the `users` table of a record which is referenced in the `posts` table, it will allow it, however when you update the `users` id (which doen't happen unless you've skewed up or it's by design), it will prevent that updation. Note, updating foreign keys which are usually the primary keys of the foreign table is very rare, and not done usually, as it might corrupt the existing data.

So, unless you have a great usecase for updating the foreign key of a table, `UPDATE ON` is never used.

```sql
DROP TABLE posts;

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    content TEXT NOT NULL,
    author_id INTEGER,
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);
```

Here, we are adding the `ON UPDATE RESTRICT` option on the `FOREIGN KEY` constraint, the `ON DELETE RESTRICT` option won't affect the update option, so it's only triggered when we try to delete the record (hard delete).

Let's add few more rows as usual:

```sql
INSERT INTO posts(content, author_id) VALUES ('hi, I am harry', 1);
INSERT INTO posts(content, author_id) VALUES ('hi, I am ron', 2);
```

Now, let's try updating the `users` table:

```sql
UPDATE users SET name = 'harry potter' WHERE id = 1;
```

This is allowed, nothing got restricted, as we didn't update the `id` in the `users` table which is a foreign key in the `posts` table, unless we update the `id` of the `users` table, the `update` option will allow any updates of the columns in the foreign table (users table)

```sql
SELECT * FROM users;
```

This readily updated the row with `id` 1 to `harry potter` as it didn't mutate/change/update the foreign key `id` in the `users` table which is referenced as `author_id` in the `posts` table.


However, if we try to update the `id` of the `users` table, it will fail.

```sql
UPDATE users SET id = 3 WHERE id = 2;
```

We are updating the id of the `users` table to 3 which is not the id of the `users` table referenced in the `posts` table, so it will fail. The existing row of id = 2 is `ron`, hence it already has a posts entry which makes it a reference, as we have added a constraint to restrict on updation of the id the foreign key.

Now, let's try to update the `id` of the `users` table to the id of the `users` table not referenced in the `posts` table

```sql
UPDATE users SET id = 5 WHERE id = 4;
```

The row with id = 4 in the `users` table is `neville` which doesn't have any posts entry, so it will succeed.
This will succeed, as we are updating the `id` of the `users` table to the id of the `users` table not referenced in the `posts` table


```
sqlite> DROP TABLE posts;

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    content TEXT NOT NULL,
    author_id INTEGER,
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);
sqlite> SELECT * FROM posts;
sqlite> INSERT INTO posts(content, author_id) VALUES ('hi, I am harry', 1);
sqlite> INSERT INTO posts(content, author_id) VALUES ('hi, I am ron', 2);
sqlite> SELECT * FROM posts;
+----+----------------+-----------+
| id |    content     | author_id |
+----+----------------+-----------+
| 1  | hi, I am harry | 1         |
| 2  | hi, I am ron   | 2         |
+----+----------------+-----------+
sqlite> SELECT * FROM users;
+----+---------+
| id |  name   |
+----+---------+
| 1  | harry   |
| 2  | ron     |
| 4  | neville |
+----+---------+
sqlite> SELECT posts.*, users.name AS author FROM posts JOIN users ON posts.author_id = users.id;
+----+----------------+-----------+--------+
| id |    content     | author_id | author |
+----+----------------+-----------+--------+
| 1  | hi, I am harry | 1         | harry  |
| 2  | hi, I am ron   | 2         | ron    |
+----+----------------+-----------+--------+
sqlite> UPDATE users SET name = 'harry potter' where id = 1;
sqlite> SELECT * FROM users;
+----+--------------+
| id |     name     |
+----+--------------+
| 1  | harry potter |
| 2  | ron          |
| 4  | neville      |
+----+--------------+
sqlite> SELECT posts.*, users.name AS author FROM posts JOIN users ON posts.author_id = users.id;
+----+----------------+-----------+--------------+
| id |    content     | author_id |    author    |
+----+----------------+-----------+--------------+
| 1  | hi, I am harry | 1         | harry potter |
| 2  | hi, I am ron   | 2         | ron          |
+----+----------------+-----------+--------------+
sqlite> UPDATE users SET name = 'harry' where id = 1;
sqlite> UPDATE users SET id = 3 where id = 2;
Runtime error: FOREIGN KEY constraint failed (19)
sqlite> UPDATE users SET id = 5 where id = 4;
sqlite> SELECT * FROM users;
+----+---------+
| id |  name   |
+----+---------+
| 1  | harry   |
| 2  | ron     |
| 5  | neville |
+----+---------+
sqlite>
```

### Set NULL

For the rest of this post, we'll only be looking at the `ON DELETE` constraint, the `ON UPDATE` is very rarely used and we have explored already when it is used.

In this option, when the record in the foreign table is deleted, the foreign key reference of the  record in the primary table will be set to `NULL`.

For our example, if we set `ON DELETE SET NULL` to the `FOREIGN KEY` constraint, when the `users` record with id = 2 is deleted, the `author_id` of the `posts` record with id = 3 will be set to `NULL`.

```sql
DROP TABLE posts;

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    content TEXT NOT NULL,
    author_id INTEGER,
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE SET NULL
);
```

Now, let's add few more rows as usual:
NOTE: I have added one more post for harry to avoid the 1-1 mapping of posts and users in the data

```sql
INSERT INTO posts(content, author_id) VALUES ('hi, I am harry', 1);
INSERT INTO posts(content, author_id) VALUES ('Expecto Patronum', 1);
INSERT INTO posts(content, author_id) VALUES ('hi, I am ron', 2);
```

Let's take a look at the posts table.

```sql
SELECT * FROM posts;
```
 
Now, let's delete the `users` record with id = 2:

```sql
DELETE FROM users WHERE id = 2;
```
This will allow and should, as there is no restriction on the `posts` table, the `FOREIGN KEY` constraint is `SET NULL` when the reference in the `users` table is deleted.

```sql
SELECT * FROM posts;
```
As you can see, the `author_id` of the `posts` record with id = 3 has been set to `NULL`

```
sqlite> DROP TABLE posts;

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    content TEXT NOT NULL,
    author_id INTEGER,
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE SET NULL
);
sqlite> INSERT INTO posts(content, author_id) VALUES ('hi, I am harry', 1);
INSERT INTO posts(content, author_id) VALUES ('Expecto Patronum', 1);
INSERT INTO posts(content, author_id) VALUES ('hi, I am ron', 2);
sqlite> SELECT * FROM posts;
+----+------------------+-----------+
| id |     content      | author_id |
+----+------------------+-----------+
| 1  | hi, I am harry   | 1         |
| 2  | Expecto Patronum | 1         |
| 3  | hi, I am ron     | 2         |
+----+------------------+-----------+
sqlite> SELECT posts.*, users.name AS author FROM posts JOIN users ON posts.author_id = users.id;
+----+------------------+-----------+--------+
| id |     content      | author_id | author |
+----+------------------+-----------+--------+
| 1  | hi, I am harry   | 1         | harry  |
| 2  | Expecto Patronum | 1         | harry  |
| 3  | hi, I am ron     | 2         | ron    |
+----+------------------+-----------+--------+
sqlite> SELECT * FROM users;
+----+---------+
| id |  name   |
+----+---------+
| 1  | harry   |
| 2  | ron     |
| 5  | neville |
+----+---------+
sqlite> DELETE FROM users WHERE id = 2;
sqlite> SELECT * FROM posts;
+----+------------------+-----------+
| id |     content      | author_id |
+----+------------------+-----------+
| 1  | hi, I am harry   | 1         |
| 2  | Expecto Patronum | 1         |
| 3  | hi, I am ron     |           |
+----+------------------+-----------+
sqlite>
```


### Set Default

In this option, when the record in the foreign table is deleted, the foreign key reference of the  record in the primary table will be set to the default value of the foreign key.

NOTE: We have not set the default value for the `author_id` column in the `posts` table, so it will be set to `NULL` by default. We can add normal constraints like `DEFAULT`, `UNIQUE`, `NOT NULL`, `CHECK`, etc on the foreign key as and when required.

```sql
DROP TABLE posts;

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    content TEXT NOT NULL,
    author_id INTEGER DEFAULT 1,
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE SET DEFAULT
);
```
We are setting default foreign key as `1` which means the `users` record with id = 1 will be the default reference for the `posts` table. If you don't provide the `author_id`, all post will be authored by `harry` what a funny quirk of the posts.

Let's add rows as usual:

```sql
INSERT INTO posts(content, author_id) VALUES ('hi, I am harry', 1);
INSERT INTO posts(content, author_id) VALUES ('Expecto Patronum', 1);
INSERT INTO posts(content, author_id) VALUES ('hi, I am ron', 2);
```

Now, let's take a look at the posts table:

```sql
SELECT * FROM posts;
```

We can now try deleting the `users` record with id = 2, Since it is already tied / referenced by the `posts` table, the users record will be deleted, but the `posts` record with `author_id` 2 will be set to the default value `1` which is the id of the `users` record with name `harry`.

```sql
DELETE FROM users WHERE id = 2;
```
It ran successfully, but the `posts` record with `author_id` 2 will be set to the default value `1` which is the id of the `users` record with name `harry`.

```sql
SELECT * FROM posts;
```

Let's view more closely:

```sql
SELECT posts.*, users.name AS author FROM posts JOIN users ON posts.author_id = users.id;
```
 
 View the SQLog :)

```
sqlite> DROP TABLE posts;

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    content TEXT NOT NULL,
    author_id INTEGER DEFAULT 1,
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE SET DEFAULT
);
sqlite> SELECT * FROM users;
+----+---------+
| id |  name   |
+----+---------+
| 1  | harry   |
| 2  | ron     |
| 5  | neville |
+----+---------+
sqlite> INSERT INTO posts(content, author_id) VALUES ('hi, I am harry', 1);
INSERT INTO posts(content, author_id) VALUES ('Expecto Patronum', 1);
INSERT INTO posts(content, author_id) VALUES ('hi, I am ron', 2);
sqlite> SELECT * FROM posts;
+----+------------------+-----------+
| id |     content      | author_id |
+----+------------------+-----------+
| 1  | hi, I am harry   | 1         |
| 2  | Expecto Patronum | 1         |
| 3  | hi, I am ron     | 2         |
+----+------------------+-----------+
sqlite> SELECT posts.*, users.name AS author FROM posts JOIN users ON posts.author_id = users.id;
+----+------------------+-----------+--------+
| id |     content      | author_id | author |
+----+------------------+-----------+--------+
| 1  | hi, I am harry   | 1         | harry  |
| 2  | Expecto Patronum | 1         | harry  |
| 3  | hi, I am ron     | 2         | ron    |
+----+------------------+-----------+--------+
sqlite> DELETE FROM users where id = 2;
sqlite> SELECT * FROM posts;
+----+------------------+-----------+
| id |     content      | author_id |
+----+------------------+-----------+
| 1  | hi, I am harry   | 1         |
| 2  | Expecto Patronum | 1         |
| 3  | hi, I am ron     | 1         |
+----+------------------+-----------+
sqlite> SELECT posts.*, users.name AS author FROM posts JOIN users ON posts.author_id = users.id;
+----+------------------+-----------+--------+
| id |     content      | author_id | author |
+----+------------------+-----------+--------+
| 1  | hi, I am harry   | 1         | harry  |
| 2  | Expecto Patronum | 1         | harry  |
| 3  | hi, I am ron     | 1         | harry  |
+----+------------------+-----------+--------+
sqlite>
```
Let's take a look at the `CASCADE` option

### Cascade

If we set `ON DELETE CASCADE` to the `FOREIGN KEY` constraint, when the `users` record with id = 2 is deleted, the `posts` record with `author_id` 2 will be deleted as well.

We can set the `CASCADE` option to the `FOREIGN KEY` constraint as follows:

```sql
DROP TABLE posts;

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    content TEXT NOT NULL,
    author_id INTEGER,
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE CASCADE
);
```

Now, let's add few more rows as usual

```sql
INSERT INTO posts(content, author_id) VALUES ('hi, I am harry', 1);
INSERT INTO posts(content, author_id) VALUES ('Expecto Patronum', 1);
INSERT INTO posts(content, author_id) VALUES ('hi, I am ron', 2);
```

We now have 3 posts, associated with 2 users.

```sql
SELECT * FROM posts;
```

Now, let's delete the `users` record with id = 2

```sql
DELETE FROM users WHERE id = 2;
```

This will not only delete the `users` record with id = 2, but it will also delete the `posts` record with `author_id` 2.

```sql
SELECT * FROM posts;
```

Cascade as the name suggest, when the foreign key in the foreign table is deleted, the primary row is cascaded or deleted.

```sql
sqlite> DROP TABLE posts;

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    content TEXT NOT NULL,
    author_id INTEGER,
    FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE CASCADE
);
sqlite> SELECT * FROM users;
+----+---------+
| id |  name   |
+----+---------+
| 1  | harry   |
| 2  | ron     |
| 5  | neville |
+----+---------+
sqlite> INSERT INTO posts(content, author_id) VALUES ('hi, I am harry', 1);
INSERT INTO posts(content, author_id) VALUES ('Expecto Patronum', 1);
INSERT INTO posts(content, author_id) VALUES ('hi, I am ron', 2);
sqlite> SELECT * FROM posts;
+----+------------------+-----------+
| id |     content      | author_id |
+----+------------------+-----------+
| 1  | hi, I am harry   | 1         |
| 2  | Expecto Patronum | 1         |
| 3  | hi, I am ron     | 2         |
+----+------------------+-----------+
sqlite> DELETE FROM users where id = 2;
sqlite> SELECT * FROM posts;
+----+------------------+-----------+
| id |     content      | author_id |
+----+------------------+-----------+
| 1  | hi, I am harry   | 1         |
| 2  | Expecto Patronum | 1         |
+----+------------------+-----------+
sqlite>
```

## Many Foreign Keys

What would happen if there are more than one row in the primary table that references the same id in the foreign table?

That is a question for another post.

