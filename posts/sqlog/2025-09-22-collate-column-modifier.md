---
type: sqlog
slug: sqlite-collate-column-modifier
title: "SQLite SQL: Collate Column Modifier"
date: "2025-09-22"
tags: ["sqlite", "sql"]
---

## Collate Column Modifier

Collate is a modifier in SQL that specifies the column how to arrange and compare the values that would be populated in the rows.

So, for a given table, we can specify a few modifiers that would let SQL decide how to handle the values.

## Adding Collate Modifier

To add a collate modifier, we can use the `COLLATE` keyword, followed by the collation name.

There are 3 collations available in SQLite:

- BINARY
- NOCASE
- RTRIM

These basically define how the column value might be treated in later comparisons or populating in querying.

Let's quickly understand each with the basic example

## BINARY

The `BINARY` modifier is the **default collation** in SQLite.  
It compares values **byte by byte**, meaning it is case-sensitive and space-sensitive.

Let's create a demo table:

```sql
CREATE TABLE users (
    name TEXT COLLATE BINARY
);
```
Now, let's insert some values

```sql
INSERT INTO users (name) VALUES ('jane'), ('John'), ('JANE'), ('Jane   ');
```

Here, I have inserted 4 names in the `users` table. I have deliberately inserted a mix of cases. Let's now see how the values are populated in the table.

```sql
SELECT * FROM users;
```
We see that there are `4` rows rightly being populated.

```sql
SELECT * FROM users WHERE name = 'jane';
```
Now, we see that only the `jane` name is being populated. This is because the `BINARY` modifier is not case sensitive and space sensitive.

```sqlite

sqlite> CREATE TABLE users (
    name TEXT COLLATE BINARY
);
sqlite> INSERT INTO users (name) VALUES ('jane'), ('John'), ('JANE'), ('Jane   ');

sqlite> SELECT * FROM users;
+---------+
|  name   |
+---------+
| jane    |
| John    |
| JANE    |
| Jane    |
+---------+
sqlite> SELECT * FROM users WHERE name = 'jane';
+------+
| name |
+------+
| jane |
+------+
```

This is the modifier added by default to the columns that have `TEXT` affinity (not necessarily the `TEXT` type).

## NOCASE

The `NOCASE` modifier is used to ignore the case of the values in the column. When comparing the two values of the same column, if that values are comparable as strings, then the case is not considered.

Let's create a typical `users` table with a `name` column

```sql
CREATE TABLE users (
    name TEXT COLLATE NOCASE 
);
```

Here, we have added a `COLLATE NOCASE` modifier to the `name` column.

Let's insert some values:

```sql
INSERT INTO users (name) VALUES ('John'), ('jane'), ('JANE');
```

Here, I have inserted 3 names in the `users` table. I have deliberately inserted a mix of cases. Let's now see how the values are populated in the table. The `Jane` name is in lowercase as well as uppercase.

```sql
SELECT * FROM users;
```

We see that there are `3` rows rightly being populated.

Now, what would happen if the user queries for `jane`?

```sql
SELECT * FROM users WHERE name = 'jane';
```
This comes with both the `jane`s the `jane` with lowercase as well `JANE` with uppercase. This is because the `NOCASE` modifier is not case sensitive.

```sqlite
sqlite> CREATE TABLE users (
    name TEXT COLLATE NOCASE
);

sqlite> INSERT INTO users (name) VALUES ('John'), ('jane'), ('JANE');

sqlite> .mode table

sqlite> SELECT * FROM users;
+------+
| name |
+------+
| John |
| jane |
| JANE |
+------+

sqlite> SELECT * FROM users WHERE name = 'jane';
+------+
| name |
+------+
| jane |
| JANE |
+------+
sqlite>
```

That is the basic idea of how the `NOCASE` collation works.


## RTRIM

The `RTRIM` modifier is used to remove the trailing spaces from the values in the column.

Let's create a typical `users` table with a `name` column

```sql
CREATE TABLE users (
    name TEXT COLLATE RTRIM 
);
```

Here, we have added a `COLLATE RTRIM` modifier to the `name` column.

Let's insert some values:

```sql
INSERT INTO users (name) VALUES ('John'), ('jane'), ('JANE'), ('Jane   '), ('jane   ');
```
Here, I have inserted 5 names in the `users` table. I have deliberately inserted a mix of cases. Let's now see how the values are populated in the table.

```sql
SELECT * FROM users;
```

We see that there are `5` rows rightly being populated.

```sql
SELECT * FROM users WHERE name = 'jane';
```

We see 2 rows being populated. This is because the `RTRIM` modifier is removing the trailing spaces.

The first row is the `jane` name without any spaces, the 2nd row that we added. The second row is the `jane` name with the trailing space, the last one that was added.


```sql
sqlite> CREATE TABLE users (
    name TEXT COLLATE RTRIM
);
sqlite> INSERT INTO users (name) VALUES ('John'), ('jane'), ('JANE'), ('Jane   '), ('jane   ');
sqlite> SELECT * FROM users;
+---------+
|  name   |
+---------+
| John    |
| jane    |
| JANE    |
| Jane    |
| jane    |
+---------+
sqlite> SELECT * FROM users WHERE name = 'jane';
+---------+
|  name   |
+---------+
| jane    |
| jane    |
+---------+
sqlite> SELECT rowid, * FROM users WHERE name = 'jane';
+-------+---------+
| rowid |  name   |
+-------+---------+
| 2     | jane    |
| 5     | jane    |
+-------+---------+
sqlite>
```

That is the basic idea of how the `RTRIM` collation works.


