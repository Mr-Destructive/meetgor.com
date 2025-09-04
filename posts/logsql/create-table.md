---
type: "logsql"
title: "SQLite SQL: Create Table Basic"
slug: sqlite-create-table
date: 2025-09-04
tags: ["sqlite", "sql"]
---

We have explored enough sqlite dot commands, now I want to explore the actual nuts and bolts of SQL in SQLite, what better place then to create tables.

## CREATE TABLE Statement

The `CREATE TABLE` as the name suggest will create a table in a sqlite database. What is a table?

> Table is a set of relations

The syntax for the create table might look something like this:

```sql
CREATE [TEMP | TEMPORARY] TABLE [IF NOT EXISTS] [schema-name.]table-name
( column-def ( , column-def )* [, table-constraint]* )
[WITHOUT ROWID] [STRICT]
```
Check [this diagram](https://sqlite.org/lang_createtable.html) to make things clear for you, might even confuse you if you are absolutely new, so just stick around we will explore it one piece at a time, and as usual follow a exhaustive, all combination exploration of this statemnet.

There are a lot of things, but we'll see each of them one by one, not all at once.

The important and mandatory things to provide are table name and atleast one column name. Not even the type of the column is needed, just the name is sufficient, since the type is really not a thing in sqlite tables unless you add `STRICT` as the table option.

So, the bare bones table creation statement might look something like this:

```sql
CREATE TABLE users(name);
```

Dead! Just the table name and one column name, can you even go with no columns (Except the rowid one that is automagically added?). You can actually, but you will have to specify the column as rowid (let’s keep that aside for now).
This is wired, but was juse exploring the possibilities of what the bare minimum is required for creating a table in sqlite.

Now, if we want to insert this into the `users` table, we can simply do this:

```sql
INSERT INTO users(name) VALUES ("abc"), ("def"), (100), (89.8);
SELECT * FROM users;
```
Simple right?

Note that you can insert anything in the id column, not just numbers, but you can't insert anything in the name column, since that is a mandatory column. That is wired but that's what sqlite has a duct typing if not opted out of it.

