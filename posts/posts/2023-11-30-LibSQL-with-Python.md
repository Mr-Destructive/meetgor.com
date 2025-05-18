{
  "type": "posts",
  "title": "Connecting LibSQL database with Python",
  "description": "",
  "status": "draft",
  "date": "2023-11-30 20:45:00",
  "tags": [
    "libsql",
    "python"
  ],
  "slug": "libsql-python",
  "image_url": "https://meetgor-cdn.pages.dev/connect-libsql-with-python.png"
}

## Introduction

LibSQL is an **Open Contribution** fork of SQLite. Open Contribution means that it allows suggestions and contributions from the community as opposed to SQLite which is open source but doesn't accept community contributions.

## Installation of LibSQL Client(s)

There are two libraries for LibSQL to interact with Python, the [libsql-client](https://github.com/libsql/libsql-client-py/) and the [libsql-experimental-python](https://github.com/libsql/libsql-experimental-python). The former is the recommended client as it is stable, whereas the latter is in development and has the latest features from the libsql database engine, however, it is compatible with the [sqlite](https://docs.python.org/3/library/sqlite3.html) module.

There are differences in how to connect and fetch responses in each of them, but if you want to get started quickly you can safely turn with the `libsql-client` as it is recommended from the official documentation. If you want to connect an existing SQLite database to libsql, you can turn up the `libsql-experimental-pyhon` package.

We will explore both of them in this article. Let's dive straight into the installation of the package.

## LibSQL Client

So, Libsql-client is the Python package provided by Turso as a Python client to interact with a libsql/sqlite database.

### Installation

The installation of a Python package is as simple as a pip install.

```bash
 pip install libsql-client
```

### Connecting to the database

Connecting to a libsql database is as simple as SQLite for the local database file. However, if you have a libsql database over the edge(Turso), you can use the API provided by SQLD to connect to that database.

#### Connecting Local database file

To connect to a simple libsql database file, you can either create a file as `mydb.db` or just move into the code straight away.

```python
import libsql_client

client = libsql_client.create_client_sync("file:temp.db")
result = client.execute("SELECT 1;")

for row in result.rows:
    print(row)
client.close()
```

OR

```python
import libsql_client

with libsql_client.create_client_sync("file:temp.db") as client:
    result = client.execute("SELECT 1;")
    
    for row in result.rows:
        print(row)
```

So, in the above code, the `client` object is a client of the libsql database. We use the [create\_client\_sync](https://libsql.org/libsql-client-py/reference.html#create_client_sync) method that accepts a few parameters `url` i.e. the URL of the database we want to connect to, and the `auth_token` which we will see when we want to connect to a database on the edge.

There is also [create\_client](https://libsql.org/libsql-client-py/reference.html#create_client) method which is an async method. You can turn up to this method if you are looking for async connections.

We have used `file:temp.db` i.e. to connect a local file in the current directory named `temp.db` . This method will return a [Client](https://libsql.org/libsql-client-py/reference.html#Client) object in this case a [ClientSync](https://libsql.org/libsql-client-py/reference.html#ClientSync) object which is a wrapper around the `Client` object. We will be using the methods available for the Client class later while querying in detail. You can see that we have used the [execute](https://libsql.org/libsql-client-py/reference.html#ClientSync.execute) method from the client object.

#### Connecting to a cloud(turso) database

To connect to a libsql database on the edge i.e. turso you need to specify the auth token(JWT) in the `create_client` method.

You can create the token for accessing the database with the

```bash
turso db tokens create mydb-name
```

This command will create a `JWT_TOKEN` , store this token securely and this will be used to access the database which is hosted on the Turso cloud.

```bash
export JWT_TOKEN="YOURTOKEN"
OR
## Save it in the .env file 
## JWT_TOKEN=YOURTOKEN
```

This will store the token in the environment variable and can be later used from the local environment.

```python
import libsql_client

with libsql_client.create_client_sync(
    "libsql://dbname-orgname.turso.io",
    auth_token="secret.something.secret"
) as client:
    result = client.execute("SELECT 1;")
    
    for row in result.rows:
        print(row)
```

The `auth_token` is the JWT Token that we created in the previous step, you can load the token as an Environment variable with the following code:

```python
import os

JWT_TOKEN = os.environ.get("JWT_TOKEN")

## OR

# pip install python-dotenv
from dotenv import load_dotenv
load_dotenv()
JWT_TOKEN = os.environ.get("JWT_TOKEN")
```

This token can be used to authenticate the connection to the Turso database. The rest of the connection code remains as it is.

### Running Queries and Fetching results

Now, that we have the connection, we can move into querying the database through the client connection.

## LibSQL-Client (SQLite compatible)

### Installation

The installation of the [libsql-experimental-python](https://badge.fury.io/py/libsql-experimental) package is simple with the pip install command with the name of the package.

```bash
pip install libsql-experimental
```

Now, we can move into connecting to the database with Python script.

### Connecting to the database

Connecting to a libsql database is as simple as SQLite for the local database file.

### Syncing with a remote database

### Running Queries and Fetching results
