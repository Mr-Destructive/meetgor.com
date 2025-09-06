---
type: "sqlog"
title: "SQLite SQL: Create Table with column types"
slug: sqlite-create-table-column-types
date: 2025-09-06
tags: ["sqlite", "sql"]
---

## Creating a table with column types

If you want to store values in a column with a specific type, you need to specify one of the following types in order to store that value in relatively type specific data.

- `TEXT` (string or characters, bunch of text)
- `BLOB` (binary data, images, raw files, etc)
- `INTEGER` (numbers which are whole, like 1, 67, 986, etc)
- `REAL` (floating or decimal point numbers, like 3.14, 85.98, 999.99, etc)
- `NULL` (empty or not defined)

Type affinity is what determines how a data is being stored and treated in sqlite, so we'll need to cover that separately. For now, let's assume that the data you give in is indeed valdiated before sending over to the database (that rarely is the case though) but for now we can avoid that discussion, and focus on creation of the structure and relations of the data.

If you are curious, you can learn about types and affinity [here](https://www.sqlite.org/datatype3.html).

Let's see how to create a table with a few columns having a specific type.

Let's say we are creating a database system and a schema for storing LLMs and their history. How interesting!

```sql
CREATE TABLE llm_models(
    name TEXT,
    model_type TEXT,
    description TEXT,
    no_parameters REAL,
    f1_score REAL,
    input_token_limit INTEGER,
    output_token_limit INTEGER,
    context_window_tokens INTEGER,
    release_date TEXT,
    knowledge_cutoff_date TEXT,
    weights_file BLOB,
    license TEXT
);
```

As you can see, we have a few columns with specific types. The name, model type, description, licence could be plain text. The number of parameters and f1_score and others could be real, as those exact number of the parameters cannot be discrete values. Similarly, the `input_token_limit`, `output_token_limit`, `context_window_tokens` can be a discrete value hence used as integer. The `release_date` and `knowledge_cutoff_date` is a date value so could be used as numeric-affinity but we need to store it as text for now. The `weights_file` is a binary file, hence storing it as a blob i.e. binary large object.

So, if we were to insert into these the default values, what do you think what will it populate with?

```sql
INSERT INTO llm_models DEFAULT VALUES;
```

```sql
SELECT * FROM llm_models;
```

```
sqlite> select * from llm_models;
+------+------------+-------------+---------------+-------------------+--------------------+-----------------------+--------------+-----------------------+--------------+---------+
| name | model_type | description | no_parameters | input_token_limit | output_token_limit | context_window_tokens | release_date | knowledge_cutoff_date | weights_file | license |
+------+------------+-------------+---------------+-------------------+--------------------+-----------------------+--------------+-----------------------+--------------+---------+
|      |            |             |               |                   |                    |                       |              |                       |              |         |
+------+------------+-------------+---------------+-------------------+--------------------+-----------------------+--------------+-----------------------+--------------+---------+
sqlite>
```

Well all is empty! That's becuase it is `None` or `NULL` by default, and we don't have any constraint on the columns as well as on the tables.

We need to add some constraints on the columns as well as the tables. 

There are options to add contraints on the column like 

- `NOT NULL`
- `UNIQUE`
- `PRIMARY KEY`
- `CHECK`
- `DEFAULT`
- `COLLATE`
- `FOREIGN KEY`

We'll explore those separately for now understand those as some restrictions on how the column values can be stored. For instance in our example, we can't store a record or a llm model without a name so we need to have it as `NOT NULL` which will prevent a record from being inserted as empty or NULL essentially.

It turns out we can also add in a `STRICT` paramter to the end of the table to make sure the table doesn't allow `TEXT` values to be stored in `INTEGER` column.

There are a lot of details I think I am going to leave it right here. We only need to add these for our conveniences, sqlite is so flexible that it can be a little intimidating to add these values and not get benefit right away, but imagine this. You have a quick experiment to run and test out how the data looks like, your data model works, you just need your crud app to be ready, no time in wasting which will be the primary key, which will be a not null column, what the default value is for each column, it can be skipped with SQLite. That's a double-edged sword, yes in some sense, but more powerful for most of the applications. It assumes you don't need validations unless you specify it, and believe me there are tons of validations you can add in sqlite to make it robust.
