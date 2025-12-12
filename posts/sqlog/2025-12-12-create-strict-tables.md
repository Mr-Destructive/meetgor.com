---
type: "sqlog"
title: "SQLITE SQL: Create Table with STRICT Option"
slug: "sqlite-create-strict-table"
date: 2025-12-12
tags: ["sqlite", "sql"]
---

We have seen how to create `TABLE` loosely Not adhering to the types. Because if we create a table with column name of type text, and insert a integer, it will happily store it as text. It is very flexible as we saw. So, in such `CREATE TABLE` statement in SQLite, without the `STRICT` constraint, the types don't matter.


If you created a table with column type as XYZ it will accept it, because it really doesn't see that. It will see the data coming in from the insert statement and store it whatever it thinks is the best one for that piece of data. Look at the below example:

```sql
CREATE TABLE t1 (n xyz);
INSERT INTO t1 values(4);
INSERT INTO t1 values("gg");
SELECT rowid, n, typeof(n) FROM t1;
```


```
rowid    n    typeof(n)
1        4    integer
2       gg    text
```


See? The column type, it doesn't matter.

Unless it's strict or any constraints, or generated conditions have been added.

The STRICT table option

Let's quote from the documentation what it means


> *   Every column definition must specify a datatype for that column.
>
> *   The freedom to specify a column without a datatype is removed.
>
> *   The datatype must be one of the following:
>
>     *   INT
>
>     *   INTEGER
>
>     *   REAL
>
>     *   TEXT
>
>     *   BLOB
>
>     *   ANY
>
> *   The [PRAGMA integrity\_check](https://sqlite.org/pragma.html#pragma_integrity_check) and [PRAGMA quick\_check](https://sqlite.org/pragma.html#pragma_quick_check) commands check the type of the content of all columns in STRICT tables and show errors if anything is amiss.
>

There are other nuances of the STRICT table options and the kind of constraint that you put on the columns, but that requires studying very specific examples. We'll check those nuances later.

For now though, we need to understand how to create a strictly typed table, and what the strict option adds to the table.

```sql
CREATE TABLE users(
    name TEXT,
    age  INT,
    credits REAL,
    profile_pic BLOB
) STRICT;
```

So, we have all the actual possible types we can use in a table column when defining a table. If you don't provide an column type, or provide any other type than `TEXT`, `INT` or `INTEGER`, `REAL`, `BLOB`, or `ANY` (don't put any, you lose the purpose of strict) it won't compile and execute the table creation. You need to provide a valid type among the 5 types.

However if you try to create a strict table with wrong column type or no column type.

```sql
CREATE TABLE t1 (t) STRICT;
-- Error: missing datatype for t1.t

CREATE TABLE t1 (t something) STRICT;
-- Error: unknown datatype for t1.t: "something"
```



Without STRICT it works as usual:

```sql
CREATE TABLE t1 (t  something);
INSERT INTO t1 values(123), ('abc'), (X''), (123.45);
SELECT t, typeof(t) FROM t1;
```


```
t    typeof(t)
123    integer
abc    text
    blob
123.45    real
```


Now back to the original example:

Insert a couple of rows:

```sql
-- All are NULL Values
INSERT INTO users DEFAULT VALUES;

INSERT INTO users (name, age, credits, profile_pic)
VALUES (
    'Alice',
    30,
    100.0,
    X'89504E470D0A1A0A'
);
```


This will insert two rows, the first one, all the columns will be `NULL`. If you look at the type of these statement. Those will be as per the table schema, consistent for all rows.

```
name    typeof(name)    age    typeof(age)    credits    typeof(credits)    profile_pic    typeof(profile_pic)
null        null        null        null
Alice    text    30    integer    100    real    137,80,78,71,13,10,26,10    blob
```

This has rightly added NULL type for the null values but when the data is in the row, it forces that type stated in the schema of the table.


Now, if we try to mess up the column data, it won't work

```sql
INSERT INTO users (name, age, credits, profile_pic)
VALUES (34, '4', 8, 123);
-- Error: cannot store INT value in BLOB column users.profile_pic

INSERT INTO users (name, age, credits, profile_pic)
VALUES (34, '4', 8, '');
-- Error: cannot store TEXT value in BLOB column users.profile_pic

INSERT INTO users (name, age, credits, profile_pic)
VALUES (34, 'abc', 8, X'');
-- Error: cannot store TEXT value in INT column users.age
```

This will work, as type affinity and the conversion is possible within the column types here.

```
INSERT INTO users (name, age, credits, profile_pic)
VALUES (CAST(34 AS INT), '3', 8, X'');
```

But if some data is not able to convert into that strict type, it will fail the constraint of strict column type.

For instance
- `123` or `""` is not force convertible to BLOB which is binary large object. We need to parse it with X'' strings for some raw data to make it a BLOB like object in SQLite.
- `abc`  is not convertible/casteble to INTEGER or REAL Value.

So, the strict type is actually strict as we see the pattern repeating in SQLite.

> It is flexible till you allow it to be, you can at anytime change the lever and make it strict

This is true for column-row level type checking with the STRICT table option while creating table.