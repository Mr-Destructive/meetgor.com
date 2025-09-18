---
type: sqlog
slug: sqlite-many-to-many-relations
title: "SQLite SQL: Many to Many Relations"
date: "2025-09-18"
tags: ["sqlite", "sql"]
---

## Many to Many Relations

We have seen [one-to-many](https://www.meetgor.com/sqlog/sqlite-one-to-many-foreign-key/) relationship and also the [self-join](https://www.meetgor.com/sqlog/sqlite-self-join-relations) relationship.

Let's extend it with `many-to-many` relationship. Here one row can refer to many other rows, but that other row can also be referred by multiple rows.

That was confusing, let's take the simple example of `users` and `posts` as usual.

- In `One-To-Many` relation, one post can only be written by one user. However, one user can author multiple posts.
- In `Many-To-Many` relation, one user can author multiple posts and one post can be written by multiple users.

That is the distinction, it is not about the relation, it is more about the design that your system needs.

Let's start creating one simple example:

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
```

We have created two tables, `users` and `posts` each representing their own data. However, we haven't referenced the users from the posts table, because that might not work out well.

Why?
Well, there are few considerations in a `many-to-many` relationship.

- In a `one-to-many` relation, we could simply add a `user_id` column in posts. That works because each post has exactly one author.
- But in a `many-to-many` relation, a single post can have multiple authors. If we put a `user_id` inside posts, we’d be forced to choose only one user, right?

What if we tried multiple columns like `user_id_1`, `user_id_2` and so on?
You can see that won't work or rather create a lot of confusion and redundancy. What if a post has 4 authors? Or 6? Not practical but could be an exception.

When we say, `many-to-many` we truly have to mean  `MANY` posts can be written by `ONE` author and `MANY` authors can write one `post`.

This is where we need to create a separate table usually called as the junction or an association table. There are [other](https://en.wikipedia.org/wiki/Associative_entity) names too.

For this we can create a table `author_post` (not the best of names but it's usually the table and it's relation as the name). This table will refer the `PRIMARY KEY`s of both the `users` and the `posts` table to map it to a single `post`, and the `PRIMARY KEY` of this table becomes the combinational pair of those two primary keys.

```sql
CREATE TABLE author_post (
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    PRIMARY KEY (user_id, post_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (post_id) REFERENCES posts(id)
);
```

Here, the `user_id` and `post_id` are the `PRIMARY KEY`s of the `users` and the `posts` table respectively, however those are the `FOREIGN KEY`s for this `author_post` table. The `PRIMARY KEY` of the `author_post` table becomes, the combination of both `user_id` and `post_id` as `PRIMARY KEY(user_id, post_id)`.

Let's take a example to make it clear as usual:

Let's say `Pekka`(sir) is writing very technical details in a post and `Glauber` (sir) is writing the practical details in a post, both are important.They write individual posts, both of them have their audiences on the same blog. But, there is a topic they want to write together, how can they co-author a `posts` with the previous design of `one-to-many` relation?

In the case of `one-to-many` relation, the `users` let's say `Pekka` with `id` as `3` and `Glauber` as `id` with `1`.  They wrote their own posts with id `98` and `99` respectively. Now, for the `100` th `post` they want to write it together. If we create a `posts` record with the `author_id` as `3`, then it will just refer the user with the id `3` which is `Pekka`, then the user with `id` as `1` can't be referred to the `posts` record.

If we tried to create 2 posts with different `user_id` it might be redundant to create 2 posts, and those might lead to inconsistent number of post counts, since each post will have unique identifier.

So, essentially, co-authoring in this `one-to-many` relationship, is not possible, we need to do something to make them write together.

The `many-to-many` relation will help them write a `post` together.

They both can write the `posts` record with just the `title` and `content`. Let's say the `post` was created with `id` as `100`

Them, while publishing the `post`, they can provide the `user_id`s, so here, two records of `author_post` will be created, one with `user_id` as `3`and `post_id` as `100`, similarly another record will be created with `user_id` as `1` and `post_id` as `100`. Notice how we are duplicating the `post_id` (it is same, as `100`), but the `user_id` is different in each case. Hence the constraint is satisfied, the `user_id` and `post_id` the combination of those two is unique. Hence those two can write a `post` together now, the `100` th post is co-authored by `Pekka` and `Glauber`.

And finally `Pekka` and `Glauber` decided to write a [post](https://turso.tech/blog/introducing-limbo-a-complete-rewrite-of-sqlite-in-rust)

NOTE: This is just a made up example, I found it easier to relate and come up with a example of the similar blog post kind of table, so took that and wrote this, sorry if anyone got hurt.

Let's look at actual queries:

Create three users `Glauber`, `Jamie` and `Pekka`.

```sql
INSERT INTO users(name) VALUES ('Glauber'), ('Jamie'), ('Pekka');
```
These users are the authors with the ids `1`, `2` and `3` respectively.

```sql
SELECT * FROM users;
```
As we can see in this users table.

Then, let's Pekka write the posts. He wrote 2 posts, one with the title `Switching to Zig from Rust` and the other with the title `RAG in SQLite`.
This is just for fun.

```sql
INSERT INTO posts(title, content) VALUES ('Switching to Zig from Rust', 'I love C');
INSERT INTO posts(title, content) VALUES ('RAG in SQLite', 'AI first database');
SELECT * FROM posts;
```
Here, the `id` of the first post is `1` and the `id` of the second post is `2`. So, `Pekka` has authored posts with id `1` and `2` but that is not mapped yet.

Let's map it in the `author_post` table, we know that the `PRIMARY KEY` for the `author_post` is a combination of `user_id` or `author_id` and the `post_id`, we know that `Pekka` has a `user_id` of `3` and he has written the `post` with `id`s `1` and `2`, so we'll simply insert those in to the `author_post` table as individual records.

```sql
INSERT INTO author_post(user_id, post_id) VALUES (3, 1), (3, 2);
```
This will create two distinct records rightly so, for each posts.

Then, Glauber decided to write a couple of posts. Let's say he wrote 2 posts, one with the title `Rewriting SQLite` and the other with the title `Offline Writes in SQLite`.

```sql
INSERT INTO posts(title, content) VALUES ('Rewriting SQLite', 'We are no more a sqlite-fork');
INSERT INTO posts(title, content) VALUES ('Offline Writes in SQLite', 'Lets sync');
SELECT * FROM posts;
```
Here, the posts created by `Glauber` have ids `3` and `4`. Now, we need to map them in the `author_post` table.

We know that `Glauber` has a `user_id` of `1` and he has written the `post` with `id`s `3` and `4`, so we'll simply insert those in to the `author_post` table as individual records.

```sql
INSERT INTO author_post(user_id, post_id) VALUES (1, 3), (1, 4);
```

So, this will create two distinct entries in the `author_post` table for the two posts created by `users` with `id` as `1`.

Now, the question of how they co-author a post? That is a piece of cake.

Let's say they wrote a post with the title `Limbo`.

```sql
INSERT INTO posts(title, content) VALUES('Limbo', 'SQLite in Rust');
```

Now, this will create a new post with `id` as `5`, let's map it in the `author_post` table

```sql
INSERT INTO author_post(user_id, post_id) VALUES (3, 5);
INSERT INTO author_post(user_id, post_id) VALUES (1, 5);
```

Here, we create two records but the `post_id` value is the same, the `user_id` value changes for both the posts, one for `Glauber` and one for `Pekka` with ids 1 and 3 respectively.

SQueaLicios!

Now, you will say, isn't it still doing redundant insertions of two `author_post` records?

Well ... yes ... but. 

Yes, that is true, it looks like we are doing redundant as we will create records in the `author_post` table (let's call it the junction table or the middle or connector table). But there is no redundant data, every row gives us unique insight into what it is actually storing. We aren't storing the original `posts` record twice or multiple twice, we are just `referencing` it with the `user_id` and the `post_id`

Another question would be, how the heck do we query this `post` which is co-authored? Or that has multiple authors in it, easy peasy, sqly.

Bear with some SQL shenanigans, some joins and wizardry.

```sql
SELECT 
    p.id,
    p.content AS post,
    GROUP_CONCAT(u.name, ', ') AS authors
FROM posts p
JOIN author_post up ON p.id = up.post_id
JOIN users u ON u.id = up.user_id
GROUP BY p.id
HAVING COUNT(u.id) > 1;
```
Don't worry if you don't get this query as a whole, just understand three things.
- What are we querying? The post id, the post content, and the author names concatenated with `,`
- From where? The `posts`, `users` and the `author_post` table.
- How? Querying all the posts, joining the `posts` with the id in the `author_post` table, then also fetching the `name` of the `author` from the `users` name. We are grouping by post id in order to keep it per post, and not duplicate the posts in case of multiple authors, the last filtering is the `HAVING` that is done as a aggregate of the count of users in the `author_post` table having more than one author in the mapping.

What about all the posts?

We just remove the filter of `HAVING` clause to include all the posts, rather than only posts having more than 1 author.

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


The heavy SQLog:

```sql
sqlite> CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL
);
sqlite> CREATE TABLE author_post (
    user_id INTEGER NOT NULL,
    post_id INTEGER NOT NULL,
    PRIMARY KEY (user_id, post_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (post_id) REFERENCES posts(id)
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
sqlite> INSERT INTO author_post(user_id, post_id) VALUES (3, 1), (3, 2);
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
GROUP BY p.id
HAVING COUNT(u.id) > 1;
+----+----------------+----------------+
| id |      post      |    authors     |
+----+----------------+----------------+
| 5  | SQLite in Rust | Pekka, Glauber |
+----+----------------+----------------+
sqlite> SELECT * FROM posts;
+----+----------------------------+------------------------------+
| id |           title            |           content            |
+----+----------------------------+------------------------------+
| 1  | Switching to Zig from Rust | I love C                     |
| 2  | RAG in SQLite              | AI first database            |
| 3  | Rewriting SQLite           | We are no more a sqlite-fork |
| 4  | Offline Writes in SQLite   | Lets sync                    |
| 5  | Limbo                      | SQLite in Rust               |
+----+----------------------------+------------------------------+
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
sqlite>
```

That’s the basic idea of a many-to-many relation in SQLite, instead of stuffing multiple IDs into one table, we create a separate mapping table. Each row in that table ties one user to one post. If multiple users are tied to the same post, we just add more mapping rows.

