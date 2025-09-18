---
type: sqlog
slug: sqlite-self-join-relations
title: "SQLite SQL: Self Join Relations"
date: "2025-09-17"
tags: ["sqlite", "sql"]
---

## Self Join Relations

In the last post, we had defined a `one-to-many` relation between the `users` table and the `posts` table.

Now, let's create a single table where we can define a `one-to-many` relationship. Yes, within a single table we can reference the primary key and it can act like a link. Since it isn't really about the table, it's more about the row relation form each other, we an connect any rows.

Let's take a example of `comments`.

```sql
CREATE TABLE comments (
    id INTEGER PRIMARY KEY,
    author TEXT NOT NULL,
    content TEXT NOT NULL,
    parent_id REFERENCES comments (id)
);
```
Here, we have created a table `comments`, it has a simple `id` as a `INTEGER PRIMARY KEY` (we know what that means right? check out 2 articles back to find out.). There are `author` and `content` columns that just hold some data, it could be any thing that actually stores the data. The final column is vital as it makes the `comments` table usable.

We have defined `parent_id` as a `REFERENCES comments (id)`, this is the key to the `self-join` relationship. It will refer any other comment with its id.

So, let's say, you have created a first comment `A`, here we don't have any previous comment, so that can be NULL. However, if someone comments `B` on my comment, then the parent_id will be of `A`. 

```sql
INSERT INTO comments (id, author, content, parent_id)
VALUES (1, 'Harry Potter', 'I think Snape is hiding something.', NULL);

INSERT INTO comments (id, author, content, parent_id)
VALUES (2, 'Ron Weasley', 'What exactly? ', 1);

INSERT INTO comments (id, author, content, parent_id)
VALUES (3, 'Draco Malfoy', 'Potter, you just want attention.', 1);

INSERT INTO comments (id, author, content, parent_id)
VALUES (4, 'Harry Potter', 'Attention is all you need', 3);

INSERT INTO comments (id, author, content, parent_id)
VALUES (5, 'Ron Weasley', 'Calm down guys', 3);
```

Ok, we have inserted a lot at once, let's break it down

- Harry adds the first comment `I think Snape is hiding something.` (parent_id will be NULL, as that is the first comment)
- Ron replies to the first comment `What exactly?` (parent_id will be 1, as that is the id of the first comment)
- Draco replies to the first comment `Potter, you just want attention.` (parent_id will be 1, as that is the id of the first comment)
- Harry replies to Draco's comment `Attention is all you need` (parent_id will be 3, as that is the id of Draco's comment)
- Ron replies to Draco's comment `Calm down guys` (parent_id will be 3, as that is the id of Draco's comment)

This is really cool, as it makes use of the same id or primary key of the table, referencing some other row.

Let's query the data and help you understand better.

To get the first comment (base comment), we can use the following query

```sql
SELECT id, author, content FROM comments WHERE parent_id IS NULL;
```

To get the replies for a given comment id, we can use the following query

```sql
SELECT id, content, parent_id FROM comments WHERE parent_id = 1;
```

To get all the comments for a given author, we can use the following query:

```sql
SELECT id, author, content FROM comments WHERE author = 'Harry Potter';
```

Here's a SQLog:

```sqlite
sqlite> CREATE TABLE comments (
    id INTEGER PRIMARY KEY,
    author TEXT NOT NULL,
    content TEXT NOT NULL,
    parent_id REFERENCES comments (id)
);
sqlite> .mode table

sqlite> INSERT INTO comments (id, author, content, parent_id)
VALUES (1, 'Harry Potter', 'I think Snape is hiding something.', NULL),
(2, 'Ron Weasley', 'What exactly? ', 1),
(3, 'Draco Malfoy', 'Potter, you just want attention.', 1),
(4, 'Harry Potter', 'Attention is all you need', 3),
(5, 'Ron Weasley', 'Calm down guys', 3);

sqlite> SELECT * FROM comments;
+----+--------------+------------------------------------+-----------+
| id |    author    |              content               | parent_id |
+----+--------------+------------------------------------+-----------+
| 1  | Harry Potter | I think Snape is hiding something. |           |
| 2  | Ron Weasley  | What exactly?                      | 1         |
| 3  | Draco Malfoy | Potter, you just want attention.   | 1         |
| 4  | Harry Potter | Attention is all you need          | 3         |
| 5  | Ron Weasley  | Calm down guys                     | 3         |
+----+--------------+------------------------------------+-----------+

sqlite> SELECT id, author, content FROM comments WHERE parent_id IS NULL;
+----+--------------+------------------------------------+
| id |    author    |              content               |
+----+--------------+------------------------------------+
| 1  | Harry Potter | I think Snape is hiding something. |
+----+--------------+------------------------------------+

sqlite> SELECT id, content, parent_id FROM comments WHERE author = 'Harry Potter';
+----+------------------------------------+-----------+
| id |              content               | parent_id |
+----+------------------------------------+-----------+
| 1  | I think Snape is hiding something. |           |
| 4  | Attention is all you need          | 3         |
+----+------------------------------------+-----------+

sqlite> SELECT id, content, parent_id FROM comments WHERE parent_id = 1;
+----+----------------------------------+-----------+
| id |             content              | parent_id |
+----+----------------------------------+-----------+
| 2  | What exactly?                    | 1         |
| 3  | Potter, you just want attention. | 1         |
+----+----------------------------------+-----------+
sqlite>
```

As you can see, we have created a `one-to-many` relation with the same single table.

Here, one comment can have mulitple comments. This makes it a bit recursive but it is what it is.

That's how certain systems are.

Self referencing relations!
