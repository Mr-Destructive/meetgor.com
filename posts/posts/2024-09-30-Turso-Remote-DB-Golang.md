{
  "type": "posts",
  "title": "Connect LibSQL Database hosted on Turso in a Golang Application",
  "description": "Exploring how to connect and query a LibSQL database hosted on Turso/Cloud in a Golang Application using libsql-client.",
  "date": "2024-09-30 23:45:00",
  "status": "published",
  "slug": "turso-libsql-db-golang",
  "tags": [
    "go",
    "turso",
    "libsql"
  ],
  "series": [
    "lets-go-with-turso"
  ],
  "series_description": "Exploring Turso's LibSQL database with Golang.",
  "image_url": "https://meetgor-cdn.pages.dev/connect-turso-libsql-golang.png"
}

## Introduction

Welcome to the new series in Golang, Let's Go with Turso. In this series, we will learn how to interact with LibSQL databases with Golang. We will connect with a remote/local LibSQL database, create Embedded replicas, set up a local LibSQL database, and so much more as we explore and find out more features of LibSQL.

## Connect a LibSQL database in a Golang application

In this post, we will learn how to connect and query a LibSQL database hosted on Turso/Cloud in a Golang Application using libsql-client package. We will go from setting up golang project, installing turso-cli, creating a database on turso with the cli, connecting to the database with shell, and golang and finally, we can query the database using Golang.

If you want to check out the YouTube video, check this out:

[Conenct LibSQL Database hosted on Turso with Golang](https://www.youtube.com/watch?v=vBrvX0X0phw)

<iframe width="560" height="315" src="https://www.youtube.com/embed/vBrvX0X0phw" frameborder="0" allowfullscreen></iframe>

### Initializing a Golang project

Let's start with initializing a Golang project.

```bash
# go mod init <git-provider-domain>/<username>/<project-name>
# Example

go mod init github.com/mr-destructive/lets-go-with-turso

```

This will initialize the project in the current directory, creating a `go.mod` file with the specification of the Golang version and the packages that we will install and use in this module.

### Installing Turso CLI

```bash
# Linux/Windows
curl -sSfL https://get.tur.so/install.sh | bash
curl -sSfL https://get.tur.so/install.sh | bash

# macOS
brew install tursodatabase/tap/turso

```

This will install the Turso CLI. To verify that Turso CLI is installed properly, you can run the version command to check the setup.

```
turso --version
```

Once it is installed, we can now log in into Turso platform, simply by running the `auth signup` or `auth login` to Register or Log-in.

```
turso auth signup

# OR

turso auth login
```

This will redirect to the browser for the OAuth flow, once signed up and logged in, this will allow to interact with the Turso platform with the CLI that we downloaded.

To make sure we are logged in as the correct user, we can run the `auth whoami` command to get the currently logged-in user.

```
turso auth whoami
```

This will print the username if you are logged-in. If everything seems correct, we can move ahead with the database creation step.

### Creating a Remote LibSQL Database on Turso

To create a LibSQL database hosted on Turso, we will use the `turso db create` command.

```
turso db create

# OR

turso db create <name>
```

This will create a database with the specified name, even if you don't provide the name, it will give out a random friendly two-word name to your database. It will create a database on the nearest location available from your location.

This command will output the following:

```
Created database <db-name> at group default in 1.99s.

Start an interactive SQL shell with the following:
    turso db shell <db-name>

To see information about the database, including a connection URL, run:
    turso db show <db-name>

To get an authentication token for the database, run:
    turso db tokens create <db-name>
```

The next step, it shows to start an interactive shell, to see information about the database, and to generate an authentication token for the database.

We will move to the next part, which would be to create an authentication token for accessing the database from an external application.

### Generating and Storing Authentication Token for LibSQL Database

After we executed the `db create` command, and it created the database on the Turso cloud, there was a command hint for creating a `token` with the command `db tokens create`

So, this command will create a JWT authentication token, that will be used to connect and read/write to the database.

```bash
turso db tokens create <db-name>

# OR

turso db tokens create <db-name> --read-only

# OR

turso db tokens create <db-name> --expiration 30d
```

We can use the simple `db tokens create <db-name>` to create an authentication token for the database with (read + write access). You can copy that returned token into a environment variable, or wherever your application can read that token.

This could be stored in the environment variable as follows:

```bash
export TURSO_AUTH_TOKEN="<token>"
```

To make a `read-only` token, we can use the flag `--read-only`. This will be handy, if you only have a database as a local replica, and the only purpose of the database is for querying data.
This will prevent any write operation on the database.

We can also use the `--expiration` flag that will be used to set the duration of the token. By default the value for expiry is `never`, but that could be a little too risky if you are making a serious application. You can either set it to `7d` which will make the token expire after seven days.


Now, we can get the remote database URL and connect to the database. The URL could be obtained by running the command `db show <db-name>`

```
turso db show <db-name>
```

This will output the following:

```bash
Name:           <db-name>
URL:            libsql://<db-name>-<username>.turso.io
ID:             <db-id>   
Group:          default
Version:        0.24.22
Locations:      bom
Size:           4.1 kB
Archived:       No
Bytes Synced:   0 B
Is Schema:      No

Database Instances:
NAME        TYPE        LOCATION
bom         primary     bom
```

The above output shows the meta-information of the database. This also has the URL hosted on Turso. We can construct the URL using the name of the database and your username as `libsql://<db-name>-<username>.turso.io`, you can set this in an environment variable or in the configuration wherever you can access it from the application.

To set the URL of the database in your application, you can use the `TURSO_DB_URL` environment variable.

```bash
export TURSO_DATABASE_URL="libsql://<db-name>-<username>.turso.io"
```

So, we have a remote database URL, and the access token configured, these are the two pieces that we will need to connect, read and write to the libsql database.


### Installing LibSQL Client for Golang

So, we can install the [libsql-client-go](https://pkg.go.dev/github.com/tursodatabase/libsql-client-go/libsql) package for Golang which will be used as an SDK in Golang to interact with a remote LibSQL database.

```bash
go get github.com/tursodatabase/libsql-client-go/libsql
```

This will install the package `libsql` into the golang module. Now, we can use this in our golang application.

### Populating the LibSQL Database

Moving ahead, we have a database, but it doesn't have data! So let's create some tables and insert some rows. We can use the `db shell` command to open an interactive SQL shell on a remote LibSQL database.

```bash
turso db shell libsql://<db-name>-<username>.turso.io
```

This will be a default `sqlite3` like a shell, where we can execute SQL commands, like `.schema`, `.mode`, `.tables`, etc.

```bash
  .dump       Render database content as SQL
  .help       List of all available commands.
  .indexes    List indexes in a table or database
  .mode       Set output mode
  .quit       Exit this program
  .read       Execute commands from a file
  .schema     Show table schemas.
  .tables     List all existing tables in the database.
```

And definitely, we can use the normal SQL queries, to read, write and delete data from the database.

#### Creating a Table

First, let's create a simple table, called `posts` with columns like `id`, `title` as a `VARCHAR(100)`, `description` as a `VARCHAR(255)`, and `content` as `TEXT` which won't be `NULL`.

```sql
CREATE TABLE posts
  (
     id          INTEGER PRIMARY KEY,
     title       VARCHAR(100),
     description VARCHAR(255),
     content     TEXT NOT NULL
  ); 
```

This will create a table `posts` on the LibSQL database, yes this will mutate the primary LibSQL database which is hosted on Turso.

#### Inserting Rows

Now, since we have the `posts` table, we will insert some rows into the table.

```bash
INSERT INTO posts (title, description, content)
VALUES 
    ('test title', 'test description', 'test content');
```

Now, we have some rows populated in the `posts` table. We can add more tables, and rows into the database, as usual, but this is just an example, so I'll keep it short.

### Connecting to the LibSQL Database

Now, we have something to query from a database, after we connect to the database via the Golang application program.

First, we will grab two pieces to connect to the database.

```bash
export TURSO_DATABASE_URL="libsql://<db-name>-<username>.turso.io"
export TURSO_AUTH_TOKEN="<token>"
```

Now, let's start with the Golang program code.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    dbURL := os.Getenv("TURSO_DATABASE_URL")
    dbToken := os.Getenv("TURSO_AUTH_TOKEN")
    dbUrl := fmt.Sprintf("%s?authToken=%s", dbURL, dbToken)
}
```

This will be the basic config to grab the database URL and the authentication token, then we can construct the `dbURL` along with `dbToken` to construct the complete dbURL which will allow to access the database.

Moving ahead, we will import `database/sql` package that will be used to open the database connection and `github.com/tursodatabase/libsql-client-go/libsql` to connect to the remote LibSQL database.

```go
package main

import (
	"database/sql"
    "fmt"
    "os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	dbURL := os.Getenv("TURSO_DATABASE_URL")
	dbToken := os.Getenv("TURSO_AUTH_TOKEN")
	dbUrl := fmt.Sprintf("%s?authToken=%s", dbURL, dbToken)

	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
		os.Exit(1)
	}
	defer db.Close()

}
```

The `sql.Open` function will open the connection to the database, this will return a `sql.DB` object. The driver selected is `libsql` and the `dbURL` is the entire URL along with the authentication token.


```go
type Post struct {
	Id          int
	Title       string
	Description string
	Content     string
}

rows, err := db.Query("SELECT * FROM posts")
if err != nil {
    fmt.Fprintf(os.Stderr, "failed to query: %s", err)
    os.Exit(1)
}

for rows.Next() {
    var post Post
    if err := rows.Scan(&post.Id, &post.Title, &post.Description, &post.Content); err != nil {
        fmt.Fprintf(os.Stderr, "failed to scan: %s", err)
        os.Exit(1)
    }
    fmt.Println(post)
}
defer rows.Close()
```

Now, let's query some data from the database. We can construct the `Post` struct that will be used to grab the required fields like `Id`, `Title`, `Description`, and `Content` from the `posts` table in the database.

Then, we will use the `db.Query` function to query the database. This function takes in a query and returns a `sql.Rows` object. We will iterate over all the `rows` returned from the database, with the `rows.Next()` that will fetch each row. Then we can `row.Scan` the row object with the appropriate and respective fields returned in the row. In this case, the `Id`, `Title`, `Description`, and the `Content` is fetched and stored into the `post` fields.

We have fetched the rows and we can do operations on them as required, this was just a basic example. So the entire code can be found below.

```go
package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Post struct {
	Id          int
	Title       string
	Description string
	Content     string
}

func main() {
	dbURL := os.Getenv("TURSO_DATABASE_URL")
	dbToken := os.Getenv("TURSO_AUTH_TOKEN")
	dbUrl := fmt.Sprintf("%s?authToken=%s", dbURL, dbToken)

	db, err := sql.Open("libsql", dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", dbUrl, err)
		os.Exit(1)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to query: %s", err)
		os.Exit(1)
	}

	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Description, &post.Content); err != nil {
			fmt.Fprintf(os.Stderr, "failed to scan: %s", err)
			os.Exit(1)
		}
		fmt.Println(post)
	}
	defer rows.Close()

}
```
The output of the above code will result in all the rows present in the post table of the LibSQL database.

```bash
$ go run remote.go

{1 test title test description test content}
{2 test title test description test content}
{3 sample post libsql tutorial create db, connect, create tables, insert rows, sync}
{4 test title test description test content}
```

I have added a few more rows to the post table, as you can see we have successfully connected, inserted, and read from the post table in the LibSQL database hosted on Turso.

For all the code related to this article, you can check out the [Let's Go with Turso](https://github.com/mr-destructive/lets-go-with-turso) GitHub repo for all the examples and additional examples for using LibSQL with Golang.

## Conclusion

So, that is a wrap for this part of the series, we have explored how to connect a remote LibSQL database hosted on Turso with Golang. In the next part of the series, we will explore how to create embedded replicas on Turso's LibSQL database in Golang.

Thank you for reading this post, If you have any questions, feedback, and suggestions, feel free to drop them in the comments.

Happy Coding :)
