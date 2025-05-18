{
  "type": "posts",
  "title": "Use Embedded Replicas of LibSQL Database hosted on Turso with a Golang Application",
  "description": "Understanding and exploring how to create, connect, and query local embedded replicas of LibSQL database hosted on Turso with a Golang application.",
  "date": "2024-10-31 21:45:00",
  "status": "published",
  "slug": "turso-libsql-embedded-replicas-golang",
  "tags": [
    "go",
    "turso",
    "libsql"
  ],
  "series": [
    "lets-go-with-turso"
  ],
  "image_url": "https://meetgor-cdn.pages.dev/embedded-replicas-libsql-turso-go.png"
}

## Introduction

Welcome to the Let's Go with Turso series. In this series, we will learn how to interact with LibSQL databases with Golang. In the past article of the series, we explored how to connect remote/local LibSQL database in golang.

With this section, we will specifally dive into understanding how to create, connect, and query local embedded replicas of LibSQL database hosted on Turso with a Golang application.

If you want to check out the YouTube video, check this out:

[LibSQL Embedded Replicas on Turso in Golang](https://www.youtube.com/watch?v=BitxB40rdVw)

<iframe width="560" height="315" src="https://www.youtube.com/embed/vBrvX0X0phw" frameborder="0" allowfullscreen></iframe>

## LibSQL Embedded Replicas on Turso in Golang

### Embedded Replicas
The embedded replica is a native feature of the libSQL database. With embedded replicas, we can have faster writes as with the local database and global read access from the remote database.

Embedded replica is like creating a copy of a primary remote database and using it for performing any operations from the applications as a local database and then it has a mechanism or standard to sync up with the primary remote database. The primary remote database serves as a single source of truth, however we can use the database locally as well. This however should be done carefully to avoid data corruption or stale data. To cope up with this stale or corruption of data, the periodic syncing can be used.

Embedded replicas have a 3 fold process
- Create a copy of the primary remote database
- Perform any operations on the local database
- Sync up with the primary remote database

![Embedded Replicas for LibSQL](https://meetgor-cdn.pages.dev/embedded-replicas-flow.png)

There are a few things to remember here:
- The local database can read it's newly written data, however other local replica databases can only view those changes once the sync has happened and the primary database has been updated from the local copy.
- These would require to sync the local with the primary database first and then the other local database also needs to sync with the primary database.

You can read more about it [here](https://docs.turso.tech/features/embedded-replicas/introduction) on the Turso documentation.

Let's get started,

What we are going to do is create a step by step procedure to understand how to work with embedded replicas of LibSQL database.

1. Perform the operations on the local LibSQL database
2. Create a Embedded Replica and sync up with the primary remote database
3. Fetch data from the primary remote database

![Embedded Replicas of LibSQL with Golang](https://meetgor-cdn.pages.dev/LibSQL_Embedded_Replicas_on_Turso_in_Golang.gif)

In this way, we can understand how to operate the Embedded Replicas as a whole from locally as well as remotely

## Initializing a Golang project

Let's start with initializing a Golang project.

```bash
# go mod init <git-provider-domain>/<username>/<project-name>
# Example

go mod init github.com/mr-destructive/lets-go-with-turso

```

This will initialize the project in the current directory, creating a `go.mod` file with the specification of the Golang version and the packages that we will install and use in this module.

## Installing go-libsql package

We will need to install the [go-libsql](https://github.com/tursodatabase/go-libsql) package, this is the driver for LibSQL for Golang. To install simply use the `go get` command to be installed as the dependency for the project

```bash
go get github.com/tursodatabase/go-libsql
```

Once this is done, we can create a local LibSQL database.

## Creating a local LibSQL database

Let's create a local LibSQL database. With the `go-libsql` package, we can connect to the local database. This will be done using the `libsql` driver. There is really no much difference than the normal SQLite database driver, this works just like SQLite, the only difference being the ability to connect to the remote database.

```go
package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/go-libsql"
)

func main() {
	dbName := "file:./local.db"

	db, err := sql.Open("libsql", dbName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s", err)
		os.Exit(1)
	}
	defer db.Close()
}
```

The above code will simply connect to the local LibSQL database which resides as the `local.db` file. Now, to demonstrate that it is an actual sqlite-like database, we can execute queries on the connected database.


```go
package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/go-libsql"
)

func main() {
	dbName := "file:./local.db"

	db, err := sql.Open("libsql", dbName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s", err)
		os.Exit(1)
	}
	defer db.Close()
	rows, err := db.Query("SELECT ABS(RANDOM()%5) FROM generate_series(1,10)")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to query %s", err)
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var i int
		if err := rows.Scan(&i); err != nil {
			fmt.Fprintf(os.Stderr, "failed to scan %s", err)
			os.Exit(1)
		}
		fmt.Println(i)
	}

}
```
Output:

```
$ go run main.go

3
0
4
4
2
1
2
3
4
1

$ go run main.go

0
2
1
2
3
2
2
1
0
2
```

Just a simple `SELECT ABS(RANDOM()%5) FROM generate_series(1,10)` query will be executed on the connected database. This query will basically genrate a random number between `-4` and `4` and take the absolute value i.e. now between 0 and 4, this will genrate 10 such numbers.

Now, we can move into the actual embedded replica creation of the primary remote database and use it as a local database to perform operations.

## Creating an Embedded Replica of Turso's LibSQL database

We need a few things to specify before we create the embedded replica, first being the primary remote database, that typically is a libsql database hosted on turso or self hosted. We also need to provide the local database path, where the local-replica will be stored or we can say the copy of the primary libsql database will be created.

We will be using the [LibSQL.NewEmbeddedReplicaConnector](https://pkg.go.dev/github.com/levydsa/libsql-go#NewEmbeddedReplicaConnector) to create a local replica of a libsql database. The function takes in two paramters, the first paramter being the local database filename path to create the copy into, and the second paramter being the primary database URL. The function returns a connector object or an error if any. The connector object is then further used with [OpenDB](https://pkg.go.dev/database/sql#OpenDB) function to create a database connection. The `OpenDB` function returns a reference of database connections which we'll use to connect and perform operations on the database.
The same `connector` object could be used to sync with the primary database after performing operations on the local database with the [Sync](https://pkg.go.dev/github.com/levydsa/libsql-go#Connector.Sync) method. This will pull or push the changes from the local database to the primary database.

We can configure the syncing mechanism while creating the embedded replica with the additional parameters to the `NewEmbeddedReplicaConnector` function. There are [Options](https://pkg.go.dev/github.com/levydsa/libsql-go#Option) to include for the paramters that could be passed like:

- `WithAuthToken(string)`: This will be used to authenticate with the primary database.
- `WithSyncInterval(time.Duration)`: This will be used to specify the interval of syncing between the local and primary database.
- `WithEncrytion(string)`: This will be used to encrypt the local database.
- `WithReadYourWrites(bool)`: This will be used to specify if the local database can read the newly written changes or not.

So, let's create a exmaple to create a embedded replica, make some changes by creating tables and then syncing the local with primary, finally appending some data to the local and reading those.

#### Create the Embedded Replica

We first need to create a copy of the primary database, as said, we will configure the 2 paramters that we need to create the embedded replica with `NewEmbeddedReplicaConnector`. Then once we have the connector object, we open up a database connection, at that point we can further run queries on the local replica that was just created from the primary remote LibSQL database.

```go
package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tursodatabase/go-libsql"
)

func main() {

	dbName := "local.db"
    // this is not required, but can be used to create a temporary directory and then delete it later
	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		fmt.Println("Error creating temporary directory:", err)
		os.Exit(1)
	}
	defer os.RemoveAll(dir)

    // first paramter required for creating NewEmbeddedReplicaConnector
	dbPath := filepath.Join(dir, dbName)
	fmt.Println(dbPath)

    // second paramter required for creating NewEmbeddedReplicaConnector
	dbURL := os.Getenv("TURSO_DATABASE_URL")
	dbAuthToken := os.Getenv("TURSO_AUTH_TOKEN")

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, dbURL, libsql.WithAuthToken(dbAuthToken))
	fmt.Println(connector)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s", err)
		os.Exit(1)
	}
	defer connector.Close()

    // open a database connection from the connector object
	db := sql.OpenDB(connector)
	fmt.Println("Connected to database")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s", err)
		os.Exit(1)
	}
	defer db.Close()
}
```

In the above code, we first create a temporary directory with the help of [MkdirTemp](https://pkg.go.dev/os#MkdirTemp), this is not required, but would be easier for cleanup later. We then the path for the local database to be created. The combined path string `dbPath` will serve as the first paramter to the `NewEmbeddedReplicaConnector`. Then we have taken the `dbURL` and the `dbAuthToken` from the environment variables `TURSO_DATABASE_URL` and `TURSO_AUTH_TOKEN` respectively. The `dbURL` will serve as the second paramter for the `NewEmbeddedReplicaConnector` that is the URL of the primary remote LibSQL database. The function `NewEmbeddedReplicaConnector` will return the `Connector` object if successfull in creation of the replica, else return `err` if it fails. The connector object needs to be closed at the end of the program, so we use the `defer connector.Close()` that will close the connection to the primary database at the end of the program. The `sql.OpenDB` is used to create the connection with the local database that will be created from the `connector` object. Finally we also need to close the local database connection at the end of the program.

Further, we will try to query the local replica and create tables and append data to it.

### Adding data to teh local replica

Once we have the `db` connection to the local database, we can noramlly query the database as we did in the previous example, of querying the local LibSQL database. Let's start by creating a table `posts` to the local replica, this will basically create the schema in the local database.

```go
    ....

	createPostTableQuery := `CREATE TABLE IF NOT EXISTS posts(
        id INTEGER PRIMARY KEY,
        title VARCHAR(100),
        description VARCHAR(255),
        content TEXT
    );`

	_, err = db.Exec(createPostTableQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create table %s", err)
		os.Exit(1)
	}
```

The `createPostTableQuery` will be the `SQL` to create the table `posts` if it doesn't already exist in the database (local replica). Then with the help of [db.Exec](https://pkg.go.dev/database/sql#DB.Exec) function, we can execute the query and return back the rows if it had created any. In this case it won't as we have just added a table.

Then, we can either sync the database to the primary, but let's populate the table `posts` with some data before syncing with the primary db.

```go

	createPostQuery := `INSERT INTO posts(title, description, content) 
        VALUES(?, ?, ?)`

	_, err = db.Exec(createPostQuery, "test title", "test description", "test content")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to insert %s", err)
		os.Exit(1)
	}

```

We have created the `createPostQuery` similarly to insert into the `posts` table in the local replica. The values are added with the placeholders in the `Exec` function as positional parameters. Once we have executed the query, this will populate the `posts` table in the lcoal replica.

We can now finally sync with the primary remote LibSQL database to make sure that the primary database also has these migrations applied.

### Syncing the local replica

Remember, `connector` is for primary database and `db` is for the local replica. So, we will sync the remote database from the replica that was created with the `connector.Sync`

```go

	_, err = connector.Sync()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to sync %s", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully synced %s db
", dbPath)
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to query %s", err)
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var title string
		var description string
		var content string
		if err := rows.Scan(&id, &title, &description, &content); err != nil {
			fmt.Fprintf(os.Stderr, "failed to scan %s", err)
			os.Exit(1)
		}
		fmt.Println(id, title, description, content)
	}

```

Output:

```bash

$ go run main.go                                                            

/tmp/libsql-349052144/local.db
&{0x2eec9d0 <nil> <nil>}
Connected to database
Successfully synced /tmp/libsql-349052144/local.db db
1 test title test description test content
```

Once we have synced the local replica, we can now query the database i.e. the local replica, with the changes, also note that this could also be done without syncing the database, but the primary database won't have the applied changes.

We finally Query the local replica with the query `SELECT * FROM posts` and print out the results. This has the 1 record in the `posts` table that we inserted.

So, that's how we basically create a local replica from a remote LibSQL database hosted on Turso. We first create the path for the local database to be copied, then provide the primary database URL and credentials, then request a copy of the primary database, then we perform any mutation or operations on the local copy and finally sync up with the remote primary database to persist the data from that replica (acting like a session of database operation).

That wraps the article for now.

For all the code related to this article, you can check out the [Let's Go with Turso](https://github.com/mr-destructive/lets-go-with-turso) GitHub repo for all the examples and additional examples for using LibSQL with Golang.


## Conclusion

So, that is a wrap for this part of the series, we have explored how to create a local embedded replica from a remote LibSQL database hosted on Turso with Golang. In the next part of the series, we will explore how to setup a local LibSQL database server and then connect it with Golang.

Thank you for reading this post, If you have any questions, feedback, and suggestions, feel free to drop them in the comments.

Happy Coding :)
