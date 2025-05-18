---
type: til
title: "LibSQL: Query a remote Turso database with cURL"
description: "Querying a libsql database hosted on turso with cURL. Using a remote sqlite-like db with turso api to query data."
status: published
slug: libsql-query-remote-db
tags: ["database", "curl", "libsql",]
date: 2023-11-15 17:30:00
---

If you are using a local [libsql](https://turso.tech/libsql) database, it is quite easy to query the database, but for a remote or a database on a hosted cloud platform like [turso](https://turso.tech/), we can use other clients or the api itself to query the data.

We can use the turso cli to get the authentication token for the database and then query the database.

## Turso CLI

Using the [turso-cli](https://docs.turso.tech/reference/turso-cli) to access the turso platform. We will use turso cli to create a libsql database, create authentication tokens, and query the db.

### Create a database (it's optional, you might already have a database)

```bash
turso db create
```

You will get a database on the turso cloud platform with some random interesting name like a passphrase.

Use the command `turso db list` and copy the URL

```graphql
DB_URL=dbname-orgname.turso.io
```

### Create an authentication token for a particular database

```bash
turso db tokens create db_name
```

Copy the JWT token and this will be used as a authentication token when accessing the remote database in the turso cloud.

```bash
TOKEN=abcdef.12345.wxyz
DB_URL=dbname-orgname.turso.io
```

* Querying the database using curl or other [api clients](https://docs.turso.tech/libsql/client-access)
    

```bash
curl -s -H "Authorization: bearer $TOKEN" \
     -d '{"statements": ["SELECT name FROM sqlite_master WHERE type=\"table\";"]}' \
     $DB_URL
```

We can use `curl` or any api client tools to send queries to the database hosted on the turso platform. We need to provide the JWT token in the `Authorization` header to connect to that particular database. The request's body is a JSON string with a list of statements to query the database.

```graphql
[
    {
        "results":
            {
                "columns": ["name"],
                "rows":[
                    ["libsql_wasm_func_table"], ["_litestream_seq"], ["_litestream_lock"], ["sqlite_sequence"], ["user"]
                ]
            }
     }
]
```

The result is a list of key-value pairs as `columns` and `rows` for each of the statements in the body. The columns are a list of column names requested in the query, and the rows are a list of rows where each row is a list of field values from the query.
