---
type: "sqlog"
title: "SQLite SQL: Create Table Bare Bones Table"
slug: sqlite-create-table-bare-bones-table
date: 2025-09-05
tags: ["sqlite", "sql"]
---

What if you only wanted one column. If you create specify a column name in create table, then there will be actually 2 columns created, one that you specified and the other which is hidden as `rowid`, unless you mentioned rowid as the column.

```sql
CREATE TABLE users(name);
INSERT INTO users(name) VALUES ("abc"), ("def");

SELECT * FROM users;
```

You would think, you have one column `name` in the table `users`. Nope!

```sql
SELECT rowid, * from users;
```

```
+-------+------+
| rowid | name |
+-------+------+
| 1     | abc  |
| 2     | def  |
+-------+------+
```

Surprise!
Any user defined table in sqlite almost always has a [rowid column](https://www.sqlite.org/rowidtable.html)

That tempts me to create a table with one column, and see what happens, how can we uniquely identify rows in that case.

```sql
CREATE TABLE users(rowid);

INSERT INTO users DEFAULT VALUES;
INSERT INTO users DEFAULT VALUES;

.mode table

SELECT * FROM users;
```

NOTE: DEFAULT VALUES, this means, use the default values for the columns, if you didn't specify a type for the column, the default values are `NULL`.


```
sqlite> create table users(rowid);
sqlite> insert into users default values;
sqlite> .mode table
sqlite> select * from users;
+-------+
| rowid |
+-------+
|       |
+-------+
sqlite> insert into users default values;
sqlite> select * from users;
+-------+
| rowid |
+-------+
|       |
|       |
+-------+
sqlite> select distinct * from users;
+-------+
| rowid |
+-------+
|       |
+-------+
sqlite> select distinct rowid from users;
+-------+
| rowid |
+-------+
|       |
+-------+
```
Here, its kind of hard to distinguish each row, so we have kind of skewed up. There is only one row and its not anymore served as a unique identifier in the table. 
As you can insert multiple same values for the users table.

```sql
DROP TABLE users;
```

Let's create a table with a table that increments the rowid auto-incrementally.

```sql
CREATE TABLE users(rowid INTEGER PRIMARY KEY AUTOINCREMENT);
INSERT INTO users DEFAULT VALUES;
INSERT INTO users DEFAULT VALUES;
SELECT * FROM users;
```

Now it simply increments each time we add a row, that's a neat little table.

I think this is far we can push sqlite, just one column.

But wait, what's beneath this?

Like why rowid? Why not name? Well, SQLite always creates this underlying rowid for structural reasons. When we name our column rowid, we're just making that hidden rowid visible. But if you create an `INTEGER PRIMARY KEY` (regardless of its name), that becomes an alias for the underlying `rowid` and not just any primary key, specifically `INTEGER PRIMARY KEY`

But here, we ourselves defined the rowid, so when you create something like a `id` not `rowid` as a primary key, it will be a alias for that underlying `rowid`

In the first example, the `CREATE TABLE users(rowid)` this is a violation of sqlite, as we now have created a table which cannot be distinctly identified with a column, as there is only one column, and that too is exposed to the user, and hence in the default case, the `rowid` is hidden column, unless the user/developer has added a primary key to a certain row. The user/developer is taking away the ownership of managing the primary key (unique identifier for a row/entry in a table) to him/herself rather than relying on the `rowid` solely.
