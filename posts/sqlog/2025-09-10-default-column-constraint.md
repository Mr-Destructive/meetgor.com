---
type: sqlog
slug: sqlite-default-column-constraint
title: 'SQLite SQL: DEFAULT column constraint'
date: "2025-09-10"
tags: ["sqlite", "sql"]
---

## DEFAULT Constraint

The `DEFAULT` constraint in SQLite is used to specify a default value for a column, if not provided, then the column will be populated with the `NULL` value as the default value. 

According to the docs:

> The DEFAULT clause specifies a default value to use for the column if no value is explicitly provided by the user when doing an INSERT. If there is no explicit DEFAULT clause attached to a column definition, then the default value of the column is NULL. 

NOTE: This constraint doesn't guard a `NULL` value, it only prevents the initial insert (or new record creation) of a `NULL` value, you can update the column later with a `NULL` value. It only is used to override the default `NULL` value, by providing a different one.

The `DEFAULT` constraint can only take constant values, and it can't deduce or dynamically store values at the time of creation, if you want that kind of default values, then the other constraint [GENERATED ALWAYS](https://sqlite.org/lang_createtable.html#the_generated_always_as_clause) would help.

> Also, DEFAULT might not be a constraint, its just a clause or a column constraint clause.

## Creating a Table with a DEFAULT Constraint

The general syntax for creating a default constraint is basically the `DEFAULT` keyword, followed by the actual literal value.

In the below example, let's create a dummy users table, with the `country_code` and `followers` as the columns, and their default values as `US` and `0` respectively. Note that the value can be wrapped in a `()` but it's not a expression so we can avoid it. However, writing the value in the `()` makes it more redable I think.

```sql
CREATE TABLE users (
    username TEXT NOT NULL UNIQUE,
    country_code TEXT DEFAULT 'US',
    followers INTEGER DEFAULT 0
);
```

Once we have a table with columns having default values, we can avoid inserting the values each time, as the default values will be populated.

```sql
INSERT INTO users(username) VALUES('alice');
```

Now, we have inserted a row with the `username` as `alice`, and the `country_code` and `followers` will be populated as the default values as they are not specified as `US` and `0` respectively.

```sql 
SELECT * FROM users;
```

It doens't mean that you can't add any other values, the default values will be used when you don't provide any value while inserting the row. In other words, the problem of seeing empty values (`NULL` values) when you don't specify will be gone, unless you update the record later explicitly with the `NULL` or empty values.

```sql
INSERT INTO users(username, country_code) VALUES('bob', 'UK');
INSERT INTO users(username, country_code, followers) VALUES('ronaldo', 'PT', 100000);
```

In the above queries, we can set the `country_code` or the `followers` if we want to insert a specific value in the `country_code` or the `followers` column. 

```sql
SELECT * FROM users;
```
As you can see, it populates the values if provided, else uses the default values.


## Inserting default values 

If you want to create a record with all default values, you can specify the `DEFAULT VALUES` as the statement clause in the `INSERT` statement.

For this to work, we need to have a default value for every column. If not specified it will be `NULL`, but if there are columns with `NOT NULL` or other constraint that needs to have some value from the user then we won't be able to insert the record without providing at least one value.

```sql
INSERT INTO users DEFAULT VALUES;
```
This above statement will fail because we don't have any default value for the `username` column and it has a `NOT NULL` constraint, as discussed in the [NOT NULL Constraint blog post](https://www.meetgor.com/sqlog/sqlite-not-null-column-constraint/), we need to provide a value for that column in order to create a new record.

So, it will only we suited, for tables having all values with some or the other values default or generated.

Like the example below:

```sql
CREATE TABLE documents (
    id INTEGER PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    title TEXT DEFAULT 'Untitled',
    content TEXT DEFAULT 'Here goes the content'
);
```

In the above table, all the columns have one or the other `default` value, like `PRIMARY KEY` this is a alias for the underlying `rowid` column (if not mentioned explicitly or with combinational columns). The default value for `created_at` at is `CURRENT_TIMESTAMP` which will look in a second, the `title` and `content` has default value of `Untitled` and `Here foes the content` respectively.

```sql
INSERT INTO documents DEFAULT VALUES;
```

Now, we can insert the `DEFAULT VALUES` without any value specified in the `VALUES` or columns mentioned in the statement.

```sql
SELECT * FROM documents;
```

As you can see, now we can insert multiple rows with default values, and the unique id will keep the records in the `documents` table unique.

## Defaults for Date and Time related columns

There are special values or constant expression that can be used as default expressions like 

- `CURRENT_TIME`  value in the format `HH:MM:SS`
- `CURRENT_DATE` value in the format `YYYY-MM-DD`
- `CURRENT_TIMESTAMP` value in the format `YYYY-MM-DD HH:MM:SS`

These values can be used as default values relating to date and time related columns which are mostly text related fields.

```sql
CREATE TABLE documents (
    id INTEGER PRIMARY KEY,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    title TEXT DEFAULT 'Untitled',
    content TEXT DEFAULT 'Here goes the content'
);
```

The various values can be used here, like for instance, `created_at`, `updated_at` which can either be `CURRENT_TIMESTAMP` or just `CURRENT_DATE` depending on the use case.

```sql
INSERT INTO documents DEFAULT VALUES;
```

Here, it will insert a record with the `created_at` and `updated_at` as the current date and time values.

```sql
SELECT * FROM documents;
```

```
sqlite> SELECT * FROM documents;
+----+---------------------+---------------------+----------+-----------------------+
| id |     created_at      |     updated_at      |  title   |        content        |
+----+---------------------+---------------------+----------+-----------------------+
| 1  | 2025-09-10 17:04:43 | 2025-09-10 17:04:43 | Untitled | Here goes the content |
+----+---------------------+---------------------+----------+-----------------------+
sqlite>
```

As you can see, now we can insert multiple rows with default values, and the unique id will keep the records in the `documents` table unique.

That's the basic we can use the `DEFAULT` constraint, you can add any expression as long as it produces a constant value for each of the records.
 
