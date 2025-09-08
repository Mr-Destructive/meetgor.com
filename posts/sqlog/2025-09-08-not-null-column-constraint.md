---
type: sqlog
slug: sqlite-not-null-column-constraint
title: 'SQLite SQL: NOT NULL column constraint'
date: "2025-09-08"
tags: ["sqlite", "sql"]
---

## NOT NULL Column constraint

The `NOT NULL` constraint is a column-level constraint that will prevent the insertion/updation of a `NULL` value in that column record.

What is `NULL`?

In SQLite, specifically, a `NULL` is a value indicating that a value is not known or just not there.

> NOTE: NULL is not equal to `''` empty string, `0` or `0.0` (zero), it's just cannot be compared to any value.

Some additional notes on the `NOT NULL` constraint:
- If a column has a `NOT NULL` constraint, and it doesn't have any other constraint like `DEFAULT`, `GENERATED`, `PRIMARY KEY`, or `FOREIGN KEY`, and there is no provided value, it will fail the constraint (i.e. error out the query).
- Only if any of the other constraint populate the column while inserting the `NOT NULL` constraint will be satisfied.

So, to take a simple iterative example, let's take the `users` table.

```sql
CREATE TABLE users(
    name TEXT NOT NULL,
    age INTEGER CHECK(age >= 0) NOT NULL,
    username TEXT UNIQUE NOT NULL,
    language TEXT DEFAULT 'en',
    country_code TEXT DEFAULT 'US' NOT NULL,
    email TEXT UNIQUE GENERATED ALWAYS AS (username || '@myapp.com')
);
```
Here we have 5 fields

- `name` which cannot be empty(`NULL`)
- `age` which cannot be less than 0 or `NULL`
- `username` which needs to be unique and cannot be `NULL`
- `language` which has a default value of `en` but it can be `NULL`
- `country_code` which has a default value of `US` but it cannot be `NULL`
- `email` which has a default value of `username || '@myapp.com'` and it will be unique.


Let's look at a few examples:

### Correct values

```sql
INSERT INTO users (name, age, username) VALUES ('Alice', 25, 'alice');
SELECT * FROM users;
```

### Duplicate username

Won't allow duplicate username and `NULL` name fields

```sql
INSERT INTO users (age, username) VALUES (25, 'alice');
SELECT * FROM users;
```

### Border case for age

Just par case for the `age` field as `0`

```sql
INSERT INTO users (name, age, username) VALUES ('Bob', 0, 'bob');
SELECT * FROM users;
```

### Invalid age check constraint

The `age` is inserted as `-1` which will fail the `CHECK` constraint.

```sql
INSERT INTO users (name, age, username) VALUES ('Bob', -1, 'bob');
SELECT * FROM users;
```

### Age as null

The `age` is inserted as `NULL` which will fail the `NOT NULL` constraint

```sql
INSERT INTO users (name, username) VALUES ('Bob', 'bob');
SELECT * FROM users;
```

### Correct Values

Will correctly populate the values.

```sql
INSERT INTO users (name, age, username) VALUES ('John', 30, 'john');
INSERT INTO users (name, age, username) VALUES ('Jim', 18, 'jim'); 
```

### Duplicate email

The email inserted will lead to duplicate and fail the `UNIQUE` constraint.

```sql
 INSERT INTO users (name, age, username, language) VALUES ('Jim', 18, 'jim', 'fr');
```

### Different country code

Will correctly populate the values.

```sql
INSERT INTO users (name, age, username, language) VALUES ('Jimmy', 18, 'jimmy', 'fr');
```

So, that is some combination of `NOT NULL` constraints along with other constraints and the insert values.

```
sqlite> INSERT INTO users (name, age, username) VALUES ('Alice', 25, 'alice');
sqlite> .mode table

sqlite> SELECT * FROM users;
+-------+-----+----------+----------+--------------+-----------------+
| name  | age | username | language | country_code |      email      |
+-------+-----+----------+----------+--------------+-----------------+
| Alice | 25  | alice    | en       | US           | alice@myapp.com |
+-------+-----+----------+----------+--------------+-----------------+

sqlite> INSERT INTO users (age, username) VALUES (25, 'alice');
Runtime error: NOT NULL constraint failed: users.name (19)

sqlite> INSERT INTO users (name, age, username) VALUES ('Bob', 0, 'bob');

sqlite> SELECT * FROM users;
+-------+-----+----------+----------+--------------+-----------------+
| name  | age | username | language | country_code |      email      |
+-------+-----+----------+----------+--------------+-----------------+
| Alice | 25  | alice    | en       | US           | alice@myapp.com |
| Bob   | 0   | bob      | en       | US           | bob@myapp.com   |
+-------+-----+----------+----------+--------------+-----------------+

sqlite> INSERT INTO users (name, age, username) VALUES ('Bob', -1, 'bob');
Runtime error: CHECK constraint failed: age >= 0 (19)

sqlite> INSERT INTO users (name, username) VALUES ('Bob', 'bob');
Runtime error: NOT NULL constraint failed: users.age (19)

sqlite> INSERT INTO users (name, age, username) VALUES ('John', 30, 'john');
sqlite> INSERT INTO users (name, age, username) VALUES ('Jim', 18, 'jim');

sqlite> INSERT INTO users (name, age, username, language) VALUES ('Jim', 18, 'jim', 'fr');
Runtime error: UNIQUE constraint failed: users.email (19)

sqlite> INSERT INTO users (name, age, username, language) VALUES ('Jimmy', 18, 'jimmy', 'fr');
+-------+-----+----------+----------+--------------+-----------------+
| name  | age | username | language | country_code |      email      |
+-------+-----+----------+----------+--------------+-----------------+
| Alice | 25  | alice    | en       | US           | alice@myapp.com |
| Bob   | 0   | bob      | en       | US           | bob@myapp.com   |
| John  | 30  | john     | en       | US           | john@myapp.com  |
| Jim   | 18  | jim      | en       | US           | jim@myapp.com   |
| Jimmy | 18  | jimmy    | fr       | US           | jimmy@myapp.com |
+-------+-----+----------+----------+--------------+-----------------+

sqlite> INSERT INTO users (name, age, username) VALUES ('Robin', 24, 'robin');
sqlite> INSERT INTO users (name, age, username) VALUES ('Robinson', 24, 'robin');
Runtime error: UNIQUE constraint failed: users.email (19)

sqlite> SELECT * FROM users;
+-------+-----+----------+----------+--------------+-----------------+
| name  | age | username | language | country_code |      email      |
+-------+-----+----------+----------+--------------+-----------------+
| Alice | 25  | alice    | en       | US           | alice@myapp.com |
| Bob   | 0   | bob      | en       | US           | bob@myapp.com   |
| John  | 30  | john     | en       | US           | john@myapp.com  |
| Jim   | 18  | jim      | en       | US           | jim@myapp.com   |
| Jimmy | 18  | jimmy    | fr       | US           | jimmy@myapp.com |
| Robin | 24  | robin    | en       | US           | robin@myapp.com |
+-------+-----+----------+----------+--------------+-----------------+
sqlite>
```


