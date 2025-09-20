---
type: sqlog
slug: sqlite-many-to-many-table-without-rowid
title: "SQLite SQL: Many to Many Table without RowID"
date: "2025-09-20"
tags: ["sqlite", "sql"]
---

## Many to Many Relation

We saw the basic example of [Many-To-Many](https://www.meetgor.com/sqlog/sqlite-many-to-many-relations/) Relation in the second last post from this, there we just focused on the concept of the relation and not so much on the structure of the junction table.

Let's take a look at the schema again:

```sql
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL
);

CREATE TABLE author_post (
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    PRIMARY KEY (user_id, post_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (post_id) REFERENCES posts (id)
);
```

Let's now populate the tables.

```sql
-- adding authors/users
INSERT INTO users(name) VALUES ('Glauber'), ('Jamie'), ('Pekka');

-- adding posts written by pekka
INSERT INTO posts(title, content) VALUES ('Switching to Zig from Rust', 'I love C');
INSERT INTO posts(title, content) VALUES ('RAG in SQLite', 'AI first database');
INSERT INTO author_post(user_id, post_id) VALUES (3, 1), (3, 2);

-- adding posts written by glauber
INSERT INTO posts(title, content) VALUES ('Rewriting SQLite', 'We are no more a sqlite-fork');
INSERT INTO posts(title, content) VALUES ('Offline Writes in SQLite', 'Lets sync');
INSERT INTO author_post(user_id, post_id) VALUES (1, 3), (1, 4);

-- adding a post co-authored by pekka and glauber
INSERT INTO posts(title, content) VALUES('Limbo', 'SQLite in Rust');
INSERT INTO author_post(user_id, post_id) VALUES (3, 5);
INSERT INTO author_post(user_id, post_id) VALUES (1, 5);
```

This has now populated the following tables:

```sql
SELECT
    p.id,
    p.content AS post,
    GROUP_CONCAT(u.name, ', ') AS authors
FROM posts p
JOIN author_post up ON p.id = up.post_id
JOIN users u ON u.id = up.user_id
GROUP BY p.id;
```

This has 5 posts, two written by `Pekka`, two written by `Glauber`, and one co-authored by both `Pekka` and `Glauber`.

Here's the SQLog:

```
sqlite> CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL
);

CREATE TABLE author_post (
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    PRIMARY KEY (user_id, post_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (post_id) REFERENCES posts (id)
);
sqlite> .mode table
sqlite> INSERT INTO users(name) VALUES ('Glauber'), ('Jamie'), ('Pekka');
sqlite> INSERT INTO posts(title, content) VALUES ('Switching to Zig from Rust', 'I love C');
INSERT INTO posts(title, content) VALUES ('RAG in SQLite', 'AI first database');
SELECT * FROM posts;
+----+----------------------------+-------------------+
| id |           title            |      content      |
+----+----------------------------+-------------------+
| 1  | Switching to Zig from Rust | I love C          |
| 2  | RAG in SQLite              | AI first database |
+----+----------------------------+-------------------+
sqlite> INSERT INTO posts(title, content) VALUES ('Rewriting SQLite', 'We are no more a sqlite-fork');
INSERT INTO posts(title, content) VALUES ('Offline Writes in SQLite', 'Lets sync');
SELECT * FROM posts;
+----+----------------------------+------------------------------+
| id |           title            |           content            |
+----+----------------------------+------------------------------+
| 1  | Switching to Zig from Rust | I love C                     |
| 2  | RAG in SQLite              | AI first database            |
| 3  | Rewriting SQLite           | We are no more a sqlite-fork |
| 4  | Offline Writes in SQLite   | Lets sync                    |
+----+----------------------------+------------------------------+
sqlite> INSERT INTO author_post(user_id, post_id) VALUES (1, 3), (1, 4);
sqlite> INSERT INTO author_post(user_id, post_id) VALUES (3, 1), (3, 2);
sqlite> INSERT INTO posts(title, content) VALUES('Limbo', 'SQLite in Rust');
sqlite> INSERT INTO author_post(user_id, post_id) VALUES (3, 5);
INSERT INTO author_post(user_id, post_id) VALUES (1, 5);
sqlite> SELECT
    p.id,
    p.content AS post,
    GROUP_CONCAT(u.name, ', ') AS authors
FROM posts p
JOIN author_post up ON p.id = up.post_id
JOIN users u ON u.id = up.user_id
GROUP BY p.id;
+----+------------------------------+----------------+
| id |             post             |    authors     |
+----+------------------------------+----------------+
| 1  | I love C                     | Pekka          |
| 2  | AI first database            | Pekka          |
| 3  | We are no more a sqlite-fork | Glauber        |
| 4  | Lets sync                    | Glauber        |
| 5  | SQLite in Rust               | Pekka, Glauber |
+----+------------------------------+----------------+
sqlite> SELECT * FROM author_post;
+---------+---------+
| user_id | post_id |
+---------+---------+
| 1       | 3       |
| 1       | 4       |
| 3       | 1       |
| 3       | 2       |
| 3       | 5       |
| 1       | 5       |
+---------+---------+

sqlite> SELECT rowid, * FROM author_post;
+-------+---------+---------+
| rowid | user_id | post_id |
+-------+---------+---------+
| 1     | 1       | 3       |
| 2     | 1       | 4       |
| 3     | 3       | 1       |
| 4     | 3       | 2       |
| 5     | 3       | 5       |
| 6     | 1       | 5       |
+-------+---------+---------+
sqlite>
```

The interesting part here is this `rowid`

```sql
SELECT rowid, * FROM author_post;
```

This table `author_post` has primary key identified from the combination of the `user_id` and `post_id` column. However, storing `rowid` makes no sense, as it is just redundant, right?

Why would you not want to have the `rowid`?

The answer is very naive and subtle, but might impact the database querying if the `author_post` has millions and billions of rows. The rowid space will be wasted for no reason.

Hence, we can avoid creating the rowid using the `WITHOUT ROWID` parameter or option while creating the table.

```sql
DROP TABLE author_post;

CREATE TABLE author_post (
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    PRIMARY KEY (user_id, post_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (post_id) REFERENCES posts (id)
) WITHOUT ROWID;
```

The rest of the commands from inserts to selection and all of the things we saw above will remain the same.

```sql
-- adding authors/users
INSERT INTO users(name) VALUES ('Glauber'), ('Jamie'), ('Pekka');

-- adding posts written by pekka
INSERT INTO posts(title, content) VALUES ('Switching to Zig from Rust', 'I love C');
INSERT INTO posts(title, content) VALUES ('RAG in SQLite', 'AI first database');
INSERT INTO author_post(user_id, post_id) VALUES (3, 1), (3, 2);

-- adding posts written by glauber
INSERT INTO posts(title, content) VALUES ('Rewriting SQLite', 'We are no more a sqlite-fork');
INSERT INTO posts(title, content) VALUES ('Offline Writes in SQLite', 'Lets sync');
INSERT INTO author_post(user_id, post_id) VALUES (1, 3), (1, 4);

-- adding a post co-authored by pekka and glauber
INSERT INTO posts(title, content) VALUES('Limbo', 'SQLite in Rust');
INSERT INTO author_post(user_id, post_id) VALUES (3, 5);
INSERT INTO author_post(user_id, post_id) VALUES (1, 5);
```

However, when you query the `author_post` table, the `rowid` will not be returned as it doesn't exists.

```sql
SELECT * FROM author_post;
```

```sql
SELECT rowid, * FROM author_post;
```
This would return error as the `rowid` doesn't exists.

We will still need the `rowid` in `posts` and `users` table as those are the `PRIMARY KEY` columns.

Both the tables `users` and `posts` will have `rowid` referenced or aliased as id of their respective columns as we know because of the special case of `INTEGER PRIMARY KEY`

```sql
SELECT rowid, _rowid_, * FROM users;
```

```sql
SELECT rowid, _rowid_, * FROM posts;
```

Only in the case of `author_post` or any junction table, where the primary key is referenced as a combination of other two foreign keys, we don't need the `rowid` as the junction table doesn't have any unique data in it, it is just the mapping of two existing data points/records in the other two tables in the database.


