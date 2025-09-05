---
type: "sqlog"
title: "SQLite: Randomly order rows"
date: 2025-08-24
---

Return rows or records in a table with random order in SQLite

Let's take a simple table for this:

```sql
CREATE TABLE user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT
);
INSERT INTO user (name) VALUES ('abc'), ('def'), ('ghi'), ('jkl'), ('mno');
```

```sql
SELECT * FROM user ORDER BY random();
```

This will return the rows in a random order

You can limit the number of rows

```sql
SELECT * FROM user ORDER BY random() LIMIT 5;
```

We can even randomly shuffle a subset of rows with limit and where clauses

```sql
SELECT * FROM user
WHERE id BETWEEN 1 and 10 ORDER BY random() limit 5;
```

This is cool if you want to get certain subset of samples but in no strict order.

