---
type: sqlog
slug: sqlite-one-to-many-foreign-key
title: "SQLite SQL: One to Many Relation with Foreign Key"
date: "2025-09-16"
tags: ["sqlite", "sql"]
---

## One to Many Relation with Foreign Key

Back to the question that we raised in the [previous](https://www.meetgor.com/sqlog/sqlite-foreign-key-table-constraint/) post, "What would happen if there are more than one row in the primary table that references the same id in the foreign table?"

This is what is precisely called the `one-to-many` relation, or what the foreign key is used for.

- One row in the foreign(parent/other) table
- that can be referenced by many rows in the primary(child) table

So, taking the example from the previous post of author and posts,

Let's define the schema again:

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

the analogy can be stated as:

- One author can have many posts
- One post can have only one author

Right?

If you look at the schema, the posts table is having the foreign key `author_id` that is referencing the primary key `id` of the `users` table. So the author can be referenced by multiple posts, however one post record can be refer only one author. I mean that is the design, a post can have multiple authors if we wanted to but we then would have to change the design then, we check that in the next post.

Let's keep the `Foreign Key` pragma/setting constraint on throughout this post.

```sql
PRAGMA foreign_keys=on;
```

Let's insert a few users and posts to get what we mean by the `one-to-many` relation.

```sql
INSERT INTO users(name) VALUES('harry'), ('ron'), ('malfoy');
```
We are inserting 3 users, `harry`, `ron`, and `malfoy` which will have ids as `1`, `2`, `3` respectively.

```sql
SELECT * FROM users;
```

As you would see in the output, we have 3 users with ids `1`, `2`, and `3`.

Now, let's insert a few posts.

```sql
INSERT INTO posts(content, author_id) VALUES
('hi, I am harry', 1),
('Expecto Patronum', 1),
('hi, I am ron', 2),
('hi, I am malfoy', 3);
```

I am inserting the `2nd` post as `harry`'s because it will confuse with the author_id and the `post.id`

```sql
SELECT * FROM posts;
```

As you would see in the output, we have 4 posts with ids `1`, `2`, `3`, `4` the first 2 are of `harry`, then the next ron and malfoy respectively.

Below are some neat little queries to visualize it better.

```sql
SELECT * FROM posts JOIN users ON posts.author_id = users.id;
```

```sql
SELECT posts.*, users.name AS author FROM posts JOIN users ON posts.author_id = users.id;
```

This is the sqlog for the things that we did so far.

```sqlite
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
sqlite> PRAGMA foreign_keys=on;
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
sqlite> INSERT INTO posts(content, author_id) VALUES
('hi, I am harry', 1),
('Expecto Patronum', 1),
('hi, I am ron', 2),
('hi, I am malfoy', 3);
sqlite> SELECT * FROM posts;
+----+------------------+-----------+
| id |     content      | author_id |
+----+------------------+-----------+
| 1  | hi, I am harry   | 1         |
| 2  | Expecto Patronum | 1         |
| 3  | hi, I am ron     | 2         |
| 4  | hi, I am malfoy  | 3         |
+----+------------------+-----------+
sqlite> SELECT * FROM posts JOIN users ON posts.author_id = users.id;
+----+------------------+-----------+----+--------+
| id |     content      | author_id | id |  name  |
+----+------------------+-----------+----+--------+
| 1  | hi, I am harry   | 1         | 1  | harry  |
| 2  | Expecto Patronum | 1         | 1  | harry  |
| 3  | hi, I am ron     | 2         | 2  | ron    |
| 4  | hi, I am malfoy  | 3         | 3  | malfoy |
+----+------------------+-----------+----+--------+
sqlite> SELECT posts.*, users.name AS author_name FROM posts JOIN users ON posts.author_id = users.id;
+----+------------------+-----------+-------------+
| id |     content      | author_id | author_name |
+----+------------------+-----------+-------------+
| 1  | hi, I am harry   | 1         | harry       |
| 2  | Expecto Patronum | 1         | harry       |
| 3  | hi, I am ron     | 2         | ron         |
| 4  | hi, I am malfoy  | 3         | malfoy      |
+----+------------------+-----------+-------------+
sqlite>
```

Now, let's explore this:

- Harry has 2 posts, which means one user can author multiple post by design.
- One post can be always be author by only one user.

Let's assume all the `NULL` values are not allowed for now, we can see one post can only be by one author, that is a constraint of the `FOREIGN KEY`, as we are referencing only one foreign key in the posts table.

This is what we call `one-author-many-posts` or `one-to-many` relation. You can also call it `many-to-one` but then you have to change it to `many-posts-one-author` that sounds a little wired. Its basically the same thing but from the different perspective, I prefer saying `one-to-many` since most of the times, the foreign/parent table needs to exist first inorder for the parent/primary table to reference any key from it.

One is on the side of the parent table (foreign table) and the Many is usually on the primary or the place where the `foreign key` is placed (in this case it is placed in the posts table).


What about multiple foreign keys?

Yes, you can have multiple foreign keys, but those will be different keys, you can't have `author_id_1`, `author_id_2` and then refer the same foreign key as `users.id`, that is possible. However, we need to define separate table for mapping users and posts and that will be called as `many-to-many` relations as will check in the next post.

