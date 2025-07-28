{
  "title": "Use Embedded Replicas of LibSQL Database hosted on Turso with a Golang Application",
  "status": "published",
  "slug": "turso-libsql-embedded-replicas-golang",
  "date": "2025-04-05T12:33:22Z"
}

<h2>Introduction</h2>
<p>Welcome to the Let's Go with Turso series. In this series, we will learn how to interact with LibSQL databases with Golang. In the past article of the series, we explored how to connect remote/local LibSQL database in golang.</p>
<p>With this section, we will specifally dive into understanding how to create, connect, and query local embedded replicas of LibSQL database hosted on Turso with a Golang application.</p>
<p>If you want to check out the YouTube video, check this out:</p>
<p><a href="https://www.youtube.com/watch?v=BitxB40rdVw">LibSQL Embedded Replicas on Turso in Golang</a></p>
<!-- raw HTML omitted -->
<h2>LibSQL Embedded Replicas on Turso in Golang</h2>
<h3>Embedded Replicas</h3>
<p>The embedded replica is a native feature of the libSQL database. With embedded replicas, we can have faster writes as with the local database and global read access from the remote database.</p>
<p>Embedded replica is like creating a copy of a primary remote database and using it for performing any operations from the applications as a local database and then it has a mechanism or standard to sync up with the primary remote database. The primary remote database serves as a single source of truth, however we can use the database locally as well. This however should be done carefully to avoid data corruption or stale data. To cope up with this stale or corruption of data, the periodic syncing can be used.</p>
<p>Embedded replicas have a 3 fold process</p>
<ul>
<li>Create a copy of the primary remote database</li>
<li>Perform any operations on the local database</li>
<li>Sync up with the primary remote database</li>
</ul>
<p><img src="https://meetgor-cdn.pages.dev/embedded-replicas-flow.png" alt="Embedded Replicas for LibSQL"></p>
<p>There are a few things to remember here:</p>
<ul>
<li>The local database can read it's newly written data, however other local replica databases can only view those changes once the sync has happened and the primary database has been updated from the local copy.</li>
<li>These would require to sync the local with the primary database first and then the other local database also needs to sync with the primary database.</li>
</ul>
<p>You can read more about it <a href="https://docs.turso.tech/features/embedded-replicas/introduction">here</a> on the Turso documentation.</p>
<p>Let's get started,</p>
<p>What we are going to do is create a step by step procedure to understand how to work with embedded replicas of LibSQL database.</p>
<ol>
<li>Perform the operations on the local LibSQL database</li>
<li>Create a Embedded Replica and sync up with the primary remote database</li>
<li>Fetch data from the primary remote database</li>
</ol>
<p><img src="https://meetgor-cdn.pages.dev/LibSQL_Embedded_Replicas_on_Turso_in_Golang.gif" alt="Embedded Replicas of LibSQL with Golang"></p>
<p>In this way, we can understand how to operate the Embedded Replicas as a whole from locally as well as remotely</p>
<h2>Initializing a Golang project</h2>
<p>Let's start with initializing a Golang project.</p>
<pre><code class="language-bash"># go mod init &lt;git-provider-domain&gt;/&lt;username&gt;/&lt;project-name&gt;
# Example

go mod init github.com/mr-destructive/lets-go-with-turso

</code></pre>
<p>This will initialize the project in the current directory, creating a <code>go.mod</code> file with the specification of the Golang version and the packages that we will install and use in this module.</p>
<h2>Installing go-libsql package</h2>
<p>We will need to install the <a href="https://github.com/tursodatabase/go-libsql">go-libsql</a> package, this is the driver for LibSQL for Golang. To install simply use the <code>go get</code> command to be installed as the dependency for the project</p>
<pre><code class="language-bash">go get github.com/tursodatabase/go-libsql
</code></pre>
<p>Once this is done, we can create a local LibSQL database.</p>
<h2>Creating a local LibSQL database</h2>
<p>Let's create a local LibSQL database. With the <code>go-libsql</code> package, we can connect to the local database. This will be done using the <code>libsql</code> driver. There is really no much difference than the normal SQLite database driver, this works just like SQLite, the only difference being the ability to connect to the remote database.</p>
<pre><code class="language-go">package main

import (
	&quot;database/sql&quot;
	&quot;fmt&quot;
	&quot;os&quot;

	_ &quot;github.com/tursodatabase/go-libsql&quot;
)

func main() {
	dbName := &quot;file:./local.db&quot;

	db, err := sql.Open(&quot;libsql&quot;, dbName)
	if err != nil {
		fmt.Fprintf(os.Stderr, &quot;failed to open db %s&quot;, err)
		os.Exit(1)
	}
	defer db.Close()
}
</code></pre>
<p>The above code will simply connect to the local LibSQL database which resides as the <code>local.db</code> file. Now, to demonstrate that it is an actual sqlite-like database, we can execute queries on the connected database.</p>
<pre><code class="language-go">package main

import (
	&quot;database/sql&quot;
	&quot;fmt&quot;
	&quot;os&quot;

	_ &quot;github.com/tursodatabase/go-libsql&quot;
)

func main() {
	dbName := &quot;file:./local.db&quot;

	db, err := sql.Open(&quot;libsql&quot;, dbName)
	if err != nil {
		fmt.Fprintf(os.Stderr, &quot;failed to open db %s&quot;, err)
		os.Exit(1)
	}
	defer db.Close()
	rows, err := db.Query(&quot;SELECT ABS(RANDOM()%5) FROM generate_series(1,10)&quot;)
	if err != nil {
		fmt.Fprintf(os.Stderr, &quot;failed to query %s&quot;, err)
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var i int
		if err := rows.Scan(&amp;i); err != nil {
			fmt.Fprintf(os.Stderr, &quot;failed to scan %s&quot;, err)
			os.Exit(1)
		}
		fmt.Println(i)
	}

}
</code></pre>
<p>Output:</p>
<pre><code>$ go run main.go

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
</code></pre>
<p>Just a simple <code>SELECT ABS(RANDOM()%5) FROM generate_series(1,10)</code> query will be executed on the connected database. This query will basically genrate a random number between <code>-4</code> and <code>4</code> and take the absolute value i.e. now between 0 and 4, this will genrate 10 such numbers.</p>
<p>Now, we can move into the actual embedded replica creation of the primary remote database and use it as a local database to perform operations.</p>
<h2>Creating an Embedded Replica of Turso's LibSQL database</h2>
<p>We need a few things to specify before we create the embedded replica, first being the primary remote database, that typically is a libsql database hosted on turso or self hosted. We also need to provide the local database path, where the local-replica will be stored or we can say the copy of the primary libsql database will be created.</p>
<p>We will be using the <a href="https://pkg.go.dev/github.com/levydsa/libsql-go#NewEmbeddedReplicaConnector">LibSQL.NewEmbeddedReplicaConnector</a> to create a local replica of a libsql database. The function takes in two paramters, the first paramter being the local database filename path to create the copy into, and the second paramter being the primary database URL. The function returns a connector object or an error if any. The connector object is then further used with <a href="https://pkg.go.dev/database/sql#OpenDB">OpenDB</a> function to create a database connection. The <code>OpenDB</code> function returns a reference of database connections which we'll use to connect and perform operations on the database.
The same <code>connector</code> object could be used to sync with the primary database after performing operations on the local database with the <a href="https://pkg.go.dev/github.com/levydsa/libsql-go#Connector.Sync">Sync</a> method. This will pull or push the changes from the local database to the primary database.</p>
<p>We can configure the syncing mechanism while creating the embedded replica with the additional parameters to the <code>NewEmbeddedReplicaConnector</code> function. There are <a href="https://pkg.go.dev/github.com/levydsa/libsql-go#Option">Options</a> to include for the paramters that could be passed like:</p>
<ul>
<li><code>WithAuthToken(string)</code>: This will be used to authenticate with the primary database.</li>
<li><code>WithSyncInterval(time.Duration)</code>: This will be used to specify the interval of syncing between the local and primary database.</li>
<li><code>WithEncrytion(string)</code>: This will be used to encrypt the local database.</li>
<li><code>WithReadYourWrites(bool)</code>: This will be used to specify if the local database can read the newly written changes or not.</li>
</ul>
<p>So, let's create a exmaple to create a embedded replica, make some changes by creating tables and then syncing the local with primary, finally appending some data to the local and reading those.</p>
<h4>Create the Embedded Replica</h4>
<p>We first need to create a copy of the primary database, as said, we will configure the 2 paramters that we need to create the embedded replica with <code>NewEmbeddedReplicaConnector</code>. Then once we have the connector object, we open up a database connection, at that point we can further run queries on the local replica that was just created from the primary remote LibSQL database.</p>
<pre><code class="language-go">package main

import (
	&quot;database/sql&quot;
	&quot;fmt&quot;
	&quot;os&quot;
	&quot;path/filepath&quot;

	&quot;github.com/tursodatabase/go-libsql&quot;
)

func main() {

	dbName := &quot;local.db&quot;
    // this is not required, but can be used to create a temporary directory and then delete it later
	dir, err := os.MkdirTemp(&quot;&quot;, &quot;libsql-*&quot;)
	if err != nil {
		fmt.Println(&quot;Error creating temporary directory:&quot;, err)
		os.Exit(1)
	}
	defer os.RemoveAll(dir)

    // first paramter required for creating NewEmbeddedReplicaConnector
	dbPath := filepath.Join(dir, dbName)
	fmt.Println(dbPath)

    // second paramter required for creating NewEmbeddedReplicaConnector
	dbURL := os.Getenv(&quot;TURSO_DATABASE_URL&quot;)
	dbAuthToken := os.Getenv(&quot;TURSO_AUTH_TOKEN&quot;)

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, dbURL, libsql.WithAuthToken(dbAuthToken))
	fmt.Println(connector)
	if err != nil {
		fmt.Fprintf(os.Stderr, &quot;failed to open db %s&quot;, err)
		os.Exit(1)
	}
	defer connector.Close()

    // open a database connection from the connector object
	db := sql.OpenDB(connector)
	fmt.Println(&quot;Connected to database&quot;)
	if err != nil {
		fmt.Fprintf(os.Stderr, &quot;failed to open db %s&quot;, err)
		os.Exit(1)
	}
	defer db.Close()
}
</code></pre>
<p>In the above code, we first create a temporary directory with the help of <a href="https://pkg.go.dev/os#MkdirTemp">MkdirTemp</a>, this is not required, but would be easier for cleanup later. We then the path for the local database to be created. The combined path string <code>dbPath</code> will serve as the first paramter to the <code>NewEmbeddedReplicaConnector</code>. Then we have taken the <code>dbURL</code> and the <code>dbAuthToken</code> from the environment variables <code>TURSO_DATABASE_URL</code> and <code>TURSO_AUTH_TOKEN</code> respectively. The <code>dbURL</code> will serve as the second paramter for the <code>NewEmbeddedReplicaConnector</code> that is the URL of the primary remote LibSQL database. The function <code>NewEmbeddedReplicaConnector</code> will return the <code>Connector</code> object if successfull in creation of the replica, else return <code>err</code> if it fails. The connector object needs to be closed at the end of the program, so we use the <code>defer connector.Close()</code> that will close the connection to the primary database at the end of the program. The <code>sql.OpenDB</code> is used to create the connection with the local database that will be created from the <code>connector</code> object. Finally we also need to close the local database connection at the end of the program.</p>
<p>Further, we will try to query the local replica and create tables and append data to it.</p>
<h3>Adding data to teh local replica</h3>
<p>Once we have the <code>db</code> connection to the local database, we can noramlly query the database as we did in the previous example, of querying the local LibSQL database. Let's start by creating a table <code>posts</code> to the local replica, this will basically create the schema in the local database.</p>
<pre><code class="language-go">    ....

	createPostTableQuery := `CREATE TABLE IF NOT EXISTS posts(
        id INTEGER PRIMARY KEY,
        title VARCHAR(100),
        description VARCHAR(255),
        content TEXT
    );`

	_, err = db.Exec(createPostTableQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, &quot;failed to create table %s&quot;, err)
		os.Exit(1)
	}
</code></pre>
<p>The <code>createPostTableQuery</code> will be the <code>SQL</code> to create the table <code>posts</code> if it doesn't already exist in the database (local replica). Then with the help of <a href="https://pkg.go.dev/database/sql#DB.Exec">db.Exec</a> function, we can execute the query and return back the rows if it had created any. In this case it won't as we have just added a table.</p>
<p>Then, we can either sync the database to the primary, but let's populate the table <code>posts</code> with some data before syncing with the primary db.</p>
<pre><code class="language-go">
	createPostQuery := `INSERT INTO posts(title, description, content) 
        VALUES(?, ?, ?)`

	_, err = db.Exec(createPostQuery, &quot;test title&quot;, &quot;test description&quot;, &quot;test content&quot;)
	if err != nil {
		fmt.Fprintf(os.Stderr, &quot;failed to insert %s&quot;, err)
		os.Exit(1)
	}

</code></pre>
<p>We have created the <code>createPostQuery</code> similarly to insert into the <code>posts</code> table in the local replica. The values are added with the placeholders in the <code>Exec</code> function as positional parameters. Once we have executed the query, this will populate the <code>posts</code> table in the lcoal replica.</p>
<p>We can now finally sync with the primary remote LibSQL database to make sure that the primary database also has these migrations applied.</p>
<h3>Syncing the local replica</h3>
<p>Remember, <code>connector</code> is for primary database and <code>db</code> is for the local replica. So, we will sync the remote database from the replica that was created with the <code>connector.Sync</code></p>
<pre><code class="language-go">
	_, err = connector.Sync()
	if err != nil {
		fmt.Fprintf(os.Stderr, &quot;failed to sync %s&quot;, err)
		os.Exit(1)
	}

	fmt.Printf(&quot;Successfully synced %s db
&quot;, dbPath)
	rows, err := db.Query(&quot;SELECT * FROM posts&quot;)
	if err != nil {
		fmt.Fprintf(os.Stderr, &quot;failed to query %s&quot;, err)
		os.Exit(1)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var title string
		var description string
		var content string
		if err := rows.Scan(&amp;id, &amp;title, &amp;description, &amp;content); err != nil {
			fmt.Fprintf(os.Stderr, &quot;failed to scan %s&quot;, err)
			os.Exit(1)
		}
		fmt.Println(id, title, description, content)
	}

</code></pre>
<p>Output:</p>
<pre><code class="language-bash">
$ go run main.go                                                            

/tmp/libsql-349052144/local.db
&amp;{0x2eec9d0 &lt;nil&gt; &lt;nil&gt;}
Connected to database
Successfully synced /tmp/libsql-349052144/local.db db
1 test title test description test content
</code></pre>
<p>Once we have synced the local replica, we can now query the database i.e. the local replica, with the changes, also note that this could also be done without syncing the database, but the primary database won't have the applied changes.</p>
<p>We finally Query the local replica with the query <code>SELECT * FROM posts</code> and print out the results. This has the 1 record in the <code>posts</code> table that we inserted.</p>
<p>So, that's how we basically create a local replica from a remote LibSQL database hosted on Turso. We first create the path for the local database to be copied, then provide the primary database URL and credentials, then request a copy of the primary database, then we perform any mutation or operations on the local copy and finally sync up with the remote primary database to persist the data from that replica (acting like a session of database operation).</p>
<p>That wraps the article for now.</p>
<p>For all the code related to this article, you can check out the <a href="https://github.com/mr-destructive/lets-go-with-turso">Let's Go with Turso</a> GitHub repo for all the examples and additional examples for using LibSQL with Golang.</p>
<h2>Conclusion</h2>
<p>So, that is a wrap for this part of the series, we have explored how to create a local embedded replica from a remote LibSQL database hosted on Turso with Golang. In the next part of the series, we will explore how to setup a local LibSQL database server and then connect it with Golang.</p>
<p>Thank you for reading this post, If you have any questions, feedback, and suggestions, feel free to drop them in the comments.</p>
<p>Happy Coding :)</p>
