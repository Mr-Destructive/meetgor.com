---
type: sqlog
slug: sqlite-create-table-column-constraint
title: 'SQLite SQL: Create Table with column constraints'
date: "2025-09-07"
tags: ["sqlite", "sql"]
---

## Create Table with column constraints

We have seen the table creation with column types, we now need to look into the constraints or limitations that we can add to those columns, in order to make the data-model more robust and add some basic validations before doing any updation/mutation on the data itself.

To add a constraint, we can specify it at the end of the column name, after the type, so there are a list of constraints that could be added to column, which are relevant to specific scenarios of the actual model of the column.
- `NOT NULL`: This will prevent insertion of `NULL` or empty values in the column field, even updating an existing filled value with `NULL` will prevent it from inserting a `NULL` value.
- `UNIQUE`: This will prevent insertion of duplicate value for an column field, same applies for updating an existing value with a duplicate will lead to violating the constraint.
- `DEFAULT`: This will add a default value to the column field, if not specified in the insertion of the value.
- `GENERATED ALWAYS AS`: A expression used to generate the value of the column field, this can be derived from the existing columns, or a constant value.
- `CHECK`: A custom check (numeric expression) can be added for that column before updating or inserting its value in the column.
- `PRIMARY KEY`: This will set the column as the primary identifier for the table, so essentially it will be like a `rowid` for the table, acting like a unique identifier among the other rows.
- `FOREIGN KEY`: This will set the column as a foreign key, and it will be a reference to the primary key of another table.

Let's look at a basic example of each, I would go into details in each subsequent posts in the future as their is a lot of nitty-gritty details to be covered.

## NOT NULL Constraint

The `NOT NULL` constraint is a column-level constraint that can be added to a column in order to avoid getting a `NULL` or empty value being populated/inserted/updated in that column field for any row in the table.

```sql
CREATE TABLE llm_models(
    name TEXT NOT NULL,
    model_type TEXT,
    model_code TEXT
);
```

```sql
INSERT INTO llm_models DEFAULT VALUES;
```

OOPs! `Runtime error: NOT NULL constraint failed: llm_models.name (19)`. The default value is `NULL` or empty if not explicitly constrained on the column. So, it tried to insert the `NULL` as the `name` of the `llm_models` record/row, and while doing so sqlite hit a error, at runtime, that the constraint of the column `llm_models.name` cannot be `NULL`. 

So, one of the solution to avoid this, could be explicitly passing a value, or setting a default value in the column setting (we will see that in the later section as the default constraint).

Let's sepcify the value of the `name` column:

```sql
INSERT INTO llm_models(name) VALUES('abc');
```

This time it successfully ran, the thing to note here is that the `model_type` or `model_code` column is populated as `NULL` becuase they don't have any constraint, especially they don't have the `NOT NULL` constraint, and hence we need to provide a value to the `name` column but the `model_type` or `model_code` columns becomes optional.

SQLite also treats `''` empty string as not null, i.e. you can insert `''` into the `name` column and it will work fine. This is a bit wired but again a flexibility being a double edged sword. We need to understand what is the difference between `NULL` and empty string in this case, those are both different values.

## UNIQUE Constraint


Now, we will add a constraint to the `model_code` column that it should be unique, we don't want it to be a primary key just yet, but each row/record in this table should have a unique model_code.

We cannot alter the table to add a constraint in sqlite, but is allowed in other databases like postgres, mysql, etc. however that is a bit risky as the previous data might become redundant and also would fail the constraints updated. So ideally we need to re-create the table, you can however alter table columns with a combination of constraints in order to maintain the backwards data correction.

Let's drop the table `llm_models` so that we can start afresh.

```sql
DROP TABLE llm_models;
CREATE TABLE llm_models(
    name TEXT NOT NULL,
    model_type TEXT,
    model_code TEXT UNIQUE
);
```

This way, we make sure the `model_code` column is unique, and we don't have any duplicate values in the column. When each new record is added to the table `llm_models` the `model_code` needs to be unique and even updation would fail if the new value is duplicated.

```sql
INSERT INTO llm_models(name) VALUES('abc');
```

Works fine, now notice here that the `model_code` is not provided and therefore it is `NULL`, let's add one more row with the same `name` and check if that fails the unique constraint on the `model_code` as it will now have the model_code duplicated `NULL` in both the rows, that should fail right? RIGHT?

```sql
INSERT INTO llm_models(name) VALUES('abc');
```

Nopes!

`NULL` is not a value to be called as unique. NULL is like undefined, you can't really distinguish between one `NULL` and other `NULL`s. Hence, the constraint is often associated with the `NOT NULL` as well to prevent the `NULL` confusion.


Let's add the `model_code` now.

```sql
INSERT INTO llm_models(name, model_code) VALUES ('abc', 'llm-1');
```

Ok, we added a record with `model_code` as `llm-1`. Now, let's add the same record again.

```sql
INSERT INTO llm_models(name, model_code) VALUES ('abc', 'llm-1');
```

And, it broke, it gave a `Runtime error: UNIQUE constraint failed: llm_models.model_code (19)` error. So, we can't add the same `model_code` twice.

```
sqlite> CREATE TABLE llm_models(
    name TEXT NOT NULL,
    model_type TEXT,
    model_code TEXT UNIQUE
);
sqlite> INSERT INTO llm_models(name) VALUES('abc');
sqlite> .mode table
sqlite> SELECT * FROM llm_models;
+------+------------+------------+
| name | model_type | model_code |
+------+------------+------------+
| abc  |            |            |
+------+------------+------------+
sqlite> INSERT INTO llm_models(name) VALUES('abc');
sqlite> SELECT * FROM llm_models;
+------+------------+------------+
| name | model_type | model_code |
+------+------------+------------+
| abc  |            |            |
| abc  |            |            |
+------+------------+------------+
sqlite> INSERT INTO llm_models(name, model_code) VALUES('abc', 'llm-1');
sqlite> INSERT INTO llm_models(name, model_code) VALUES('abc', 'llm-1');
Runtime error: UNIQUE constraint failed: llm_models.model_code (19)
sqlite> SELECT * FROM llm_models;
+------+------------+------------+
| name | model_type | model_code |
+------+------------+------------+
| abc  |            |            |
| abc  |            |            |
| abc  |            | llm-1      |
+------+------------+------------+
sqlite>
```

So some tips for using the `UNIQUE` constraints in sqlite:
- You can't use `UNIQUE` constraint on the `NULL` values.
- Prefer using `NOT NULL` in addition to the `UNIQUE` constraint.


## DEFAULT Constraint

The default constraint as the name suggest, will set a default value for a column if no value is provided in the insert statement. However, you need to provide the default value while creating the table. You can generate a default value (for that we can use the `GENERATED ALWAYS` constraint).

```sql
DROP TABLE llm_models;
CREATE TABLE llm_models(
    name TEXT NOT NULL,
    model_type TEXT DEFAULT 'text',
    model_code TEXT UNIQUE NOT NULL
);
```

I have added a constraint of `model_type` as `DEFAULT 'text'` which will set the value of `model_type` column to `text` if no value is provided in the insert statement.

Also, the `model_code` column is updated with the constraints of `UNIQUE` and `NOT NULL`

```sql
INSERT INTO llm_models(name, model_code) VALUES('abc', 'llm-1');
SELECT * FROM llm_models;
```

And that works as expected, it populates the value of `model_type` as `text` by default.

```
sqlite> CREATE TABLE llm_models(
    name TEXT NOT NULL,
    model_type TEXT DEFAULT 'text',
    model_code TEXT UNIQUE NOT NULL
);
sqlite> INSERT INTO llm_models(name, model_code) VALUES('abc', 'llm-1');
sqlite> SELECT * FROM llm_models;
+------+------------+------------+
| name | model_type | model_code |
+------+------------+------------+
| abc  | text       | llm-1      |
+------+------------+------------+
sqlite> INSERT INTO llm_models(name, model_code) VALUES('abc', 'llm-1');
Runtime error: UNIQUE constraint failed: llm_models.model_code (19)
sqlite>
```

And if we try to insert the same `model_code` value again, it will fail with the `UNIQUE` constraint.

## Generated Constraint

The [Generated](https://sqlite.org/gencol.html) constraint can add a value to a column based on evaluated expression from the values of the other columns in that row.


```sql
DROP TABLE llm_models;
CREATE TABLE llm_models(
    name TEXT NOT NULL,
    model_type TEXT DEFAULT 'text',
    model_code TEXT UNIQUE NOT NULL,
    display_name TEXT GENERATED ALWAYS AS (name || ' (' || model_code || ')') STORED
);
```

We have added a `GENERATED ALWAYS AS STORED` constraint which will store the result of `name || ' (' || model_code || ')'` in the `display_name` column.

```sql
INSERT INTO llm_models(name, model_code) VALUES('abc', 'llm-1');
SELECT * FROM llm_models;
```

There is a variant of `VIRTUAL` instead of `STORED`, we will explore this separately, I am just getting my hands on and getting familiarity with all the constraints available in SQL.

```sql
sqlite> CREATE TABLE llm_models(
    name TEXT NOT NULL,
    model_type TEXT DEFAULT 'text',
    model_code TEXT UNIQUE NOT NULL,
    display_name TEXT GENERATED ALWAYS AS (name || ' (' || model_code || ')') STORED
);
sqlite> .mode table
sqlite> INSERT INTO llm_models(name, model_code) VALUES('abc', 'llm-1');
sqlite> SELECT * FROM llm_models;
+------+------------+------------+--------------+
| name | model_type | model_code | display_name |
+------+------------+------------+--------------+
| abc  | text       | llm-1      | abc (llm-1)  |
+------+------------+------------+--------------+
```

Neat! It populated, or I should say `generated` the `display_name` column by evaluating the `name || ' (' || model_code || ')'` expression.

## Check Constraint

The `CHECK` constraint is a column-level constraint that can be added to a column in order to add a custom check (numeric expression) before updating or inserting its value in the column. Just a simple expression or a validation check before inserting the value to the column in the row.

Let's take the `model_type` constraint with the `CHECK` constraint by limiting the values of `model_type` to `text`, `conversational`, `multimodal`, `code`.

```sql
DROP TABLE llm_models;
CREATE TABLE llm_models(
    name TEXT NOT NULL,
    model_type TEXT CHECK(model_type IN ('text', 'conversational', 'multimodal', 'code')) DEFAULT 'text',
    model_code TEXT UNIQUE NOT NULL,
    display_name TEXT GENERATED ALWAYS AS (name || ' (' || model_code || ')') STORED
);
```
Here, we have added the `CHECK` constraint to the `model_type` column. And we have limited the values of `model_type` to `text`, `conversational`, `multimodal`, `code`. It could be anything, but its just a true or false check that will either fail or work.

```sql
INSERT INTO llm_models(name, model_code) VALUES('abc', 'llm-1');
SELECT * FROM llm_models;
```

This will add the default value as `text` for the `model_type`.

```sql
INSERT INTO llm_models(name, model_code, model_type) VALUES('abc', 'llm-1', 'visual');
SELECT * FROM llm_models;
```

Here we try to add the column `model_type` as `visual` which is not allowed by the `CHECK` constraint.


```sql
INSERT INTO llm_models(name, model_code, model_type) VALUES('abc', 'llm-1', 'code');
SELECT * FROM llm_models;
```

And we try to add the column `model_type` as `code` which is not allowed by the `CHECK` constraint.

```
sqlite> CREATE TABLE llm_models(
    name TEXT NOT NULL,
    model_type TEXT CHECK(model_type IN ('text', 'conversational', 'multimodal', 'code')) DEFAULT 'text',
    model_code TEXT UNIQUE NOT NULL,
    display_name TEXT GENERATED ALWAYS AS (name || ' (' || model_code || ')') STORED
);
sqlite> INSERT INTO llm_models(name, model_code) VALUES('abc', 'llm-1');
sqlite> SELECT * FROM llm_models
   ...> ;
+------+------------+------------+--------------+
| name | model_type | model_code | display_name |
+------+------------+------------+--------------+
| abc  | text       | llm-1      | abc (llm-1)  |
+------+------------+------------+--------------+
sqlite> INSERT INTO llm_models(name, model_code, model_type) VALUES('abc', 'llm-1', 'visual');
Runtime error: CHECK constraint failed: model_type IN ('text', 'conversational', 'multimodal', 'code') (19)
sqlite> INSERT INTO llm_models(name, model_code, model_type) VALUES('abc', 'llm-1', 'code');
Runtime error: UNIQUE constraint failed: llm_models.model_code (19)
sqlite> INSERT INTO llm_models(name, model_code, model_type) VALUES('abc', 'llm-2', 'code');
sqlite> SELECT * FROM llm_models;
+------+------------+------------+--------------+
| name | model_type | model_code | display_name |
+------+------------+------------+--------------+
| abc  | text       | llm-1      | abc (llm-1)  |
| abc  | code       | llm-2      | abc (llm-2)  |
+------+------------+------------+--------------+
sqlite>
```

Right! We checked that it failed the constraint `model_type IN ('text', 'conversational', 'multimodal', 'code')` and we checked that it failed the constraint `UNIQUE` constraint.


## Primary Key Constraint

The primary key constraint is a column-level constraint that is used to identify a unique row in a table.
It is a unique constraint that can be used to identify a row in a table using a single column.

We can provide it right after the column name or at the bottom of all the table constraints.

```sql
DROP TABLE llm_models;
CREATE TABLE llm_models(
    name TEXT NOT NULL,
    model_type TEXT CHECK(model_type IN ('text', 'conversational', 'multimodal', 'code')) DEFAULT 'text',
    model_code TEXT UNIQUE NOT NULL,
    display_name TEXT GENERATED ALWAYS AS (name || ' (' || model_code || ')') STORED,
    PRIMARY KEY(model_code)
);
```

OR

```sql
DROP TABLE llm_models;
CREATE TABLE llm_models(
    name TEXT NOT NULL,
    model_type TEXT CHECK(model_type IN ('text', 'conversational', 'multimodal', 'code')) DEFAULT 'text',
    model_code TEXT UNIQUE NOT NULL PRIMARY KEY,
    display_name TEXT GENERATED ALWAYS AS (name || ' (' || model_code || ')') STORED
);
```

The `PRIMARY KEY` is a table level constraint, as it needs to unique across all the rows in the table.

## Foreign Key Constraint

The foreign key constraint is a column-level constraint that is used to establish a relationship between two tables.

We will now create a separate table called `companies` for this example, this will be a table that stores the company information those will be creating or developing the LLM models.

```sql
CREATE TABLE companies(
    id INTEGER PRIMARY KEY,
    company_name TEXT UNIQUE NOT NULL,
    founded_year INTEGER
);
```
Let's insert a few companies

```sql
INSERT INTO companies (company_name, founded_year) VALUES('ABC', 1998);
INSERT INTO companies (company_name, founded_year) VALUES('ClosedAI', 2016);
INSERT INTO companies (company_name, founded_year) VALUES('Beta', 2014);
SELECT * FROM companies;
```

Original companies right?
Now we will update the `llm_models` table that will store the LLM models information. We will add a foreign key, that will be a reference i.e. to make a relation for the row that will reference another row. So, the foreign key indicates a key (field/column) in another table that's why named as `FOREIGN KEY`. This foreign key will be referencing the `companies.id` column from the `companies` table.

```sql
CREATE TABLE llm_models(
    name TEXT NOT NULL,
    model_type TEXT CHECK(model_type IN ('text', 'conversational', 'multimodal', 'code')) DEFAULT 'text',
    model_code TEXT UNIQUE NOT NULL PRIMARY KEY,
    display_name TEXT GENERATED ALWAYS AS (name || ' (' || model_code || ')') STORED,
    company_id INTEGER,
    FOREIGN KEY(company_id) REFERENCES companies(id)
);
```
We insert the `llm_models` with the foreign key of 2 indicating the company `ClosedAI`.

```sql
INSERT INTO llm_models(name, model_code, company_id) VALUES('abc', 'llm-1', 2);
```

Now, if we see, this has successfully inserted the row.

```sql
SELECT * FROM llm_models;
```

We can join the two tables together, this will give us the combined sets form the tables.

```sql
SELECT * FROM llm_models as llm INNER JOIN companies as comp ON comp.id = llm.company_id;
```

That was a bit unclear, let's only keep what we needed.

```sql
SELECT llm.*, comp.company_name  FROM llm_models as llm INNER JOIN companies as comp ON comp.id = llm.company_id;
```

Cool, that is sweat.

```
sqlite> CREATE TABLE companies(
    id INTEGER PRIMARY KEY,
    company_name TEXT UNIQUE NOT NULL,
    founded_year INTEGER
);
sqlite> INSERT INTO companies (company_name, founded_year) VALUES('ABC', 1998);
INSERT INTO companies (company_name, founded_year) VALUES('ClosedAI', 2016);
INSERT INTO companies (company_name, founded_year) VALUES('Beta', 2014);
SELECT * FROM companies;
+----+--------------+--------------+
| id | company_name | founded_year |
+----+--------------+--------------+
| 1  | ABC          | 1998         |
| 2  | ClosedAI     | 2016         |
| 3  | Beta         | 2014         |
+----+--------------+--------------+
sqlite>
sqlite> CREATE TABLE llm_models(
    name TEXT NOT NULL,
    model_type TEXT CHECK(model_type IN ('text', 'conversational', 'multimodal', 'code')) DEFAULT 'text',
    model_code TEXT UNIQUE NOT NULL PRIMARY KEY,
    display_name TEXT GENERATED ALWAYS AS (name || ' (' || model_code || ')') STORED,
    company_id INTEGER,
    FOREIGN KEY(company_id) REFERENCES companies(id)

);
sqlite> SELECT * FROM llm_models;
sqlite> INSERT INTO llm_models(name, model_code, model_type, company_id) VALUES('abc', 'llm-1', 'code', 2);
sqlite> SELECT * FROM llm_models;
+------+------------+------------+--------------+------------+
| name | model_type | model_code | display_name | company_id |
+------+------------+------------+--------------+------------+
| abc  | code       | llm-1      | abc (llm-1)  | 2          |
+------+------------+------------+--------------+------------+
sqlite> SELECT * FROM llm_models as llm INNER JOIN companies as comp ON comp.id = llm.company_id;
+------+------------+------------+--------------+------------+----+--------------+--------------+
| name | model_type | model_code | display_name | company_id | id | company_name | founded_year |
+------+------------+------------+--------------+------------+----+--------------+--------------+
| abc  | code       | llm-1      | abc (llm-1)  | 2          | 2  | ClosedAI     | 2016         |
+------+------------+------------+--------------+------------+----+--------------+--------------+
sqlite> SELECT llm.*, comp.company_name  FROM llm_models as llm INNER JOIN companies as comp ON comp.id = llm.company_id;
+------+------------+------------+--------------+------------+--------------+
| name | model_type | model_code | display_name | company_id | company_name |
+------+------------+------------+--------------+------------+--------------+
| abc  | code       | llm-1      | abc (llm-1)  | 2          | ClosedAI     |
+------+------------+------------+--------------+------------+--------------+
sqlite>
```

That is it, a high level overview of adding constraints in SQLite, I will dig deep into each constraint in separate posts over the coming week.
