{
  "title": "Connect LibSQL Database hosted on Turso in a Golang Application",
  "status": "published",
  "slug": "turso-libsql-db-golang",
  "date": "2025-04-05T12:33:23Z"
}

<h2>Introduction</h2>
<p>Welcome to the new series in Golang, Let's Go with Turso. In this series, we will learn how to interact with LibSQL databases with Golang. We will connect with a remote/local LibSQL database, create Embedded replicas, set up a local LibSQL database, and so much more as we explore and find out more features of LibSQL.</p>
<h2>Connect a LibSQL database in a Golang application</h2>
<p>In this post, we will learn how to connect and query a LibSQL database hosted on Turso/Cloud in a Golang Application using libsql-client package. We will go from setting up golang project, installing turso-cli, creating a database on turso with the cli, connecting to the database with shell, and golang and finally, we can query the database using Golang.</p>
<p>If you want to check out the YouTube video, check this out:</p>
<p><a href="https://www.youtube.com/watch?v=vBrvX0X0phw">Conenct LibSQL Database hosted on Turso with Golang</a></p>
<!-- raw HTML omitted -->
<h3>Initializing a Golang project</h3>
<p>Let's start with initializing a Golang project.</p>
<pre><code class="language-bash"># go mod init &lt;git-provider-domain&gt;/&lt;username&gt;/&lt;project-name&gt;
# Example

go mod init github.com/mr-destructive/lets-go-with-turso

</code></pre>
<p>This will initialize the project in the current directory, creating a <code>go.mod</code> file with the specification of the Golang version and the packages that we will install and use in this module.</p>
<h3>Installing Turso CLI</h3>
<pre><code class="language-bash"># Linux/Windows
curl -sSfL https://get.tur.so/install.sh | bash
curl -sSfL https://get.tur.so/install.sh | bash

# macOS
brew install tursodatabase/tap/turso

</code></pre>
<p>This will install the Turso CLI. To verify that Turso CLI is installed properly, you can run the version command to check the setup.</p>
<pre><code>turso --version
</code></pre>
<p>Once it is installed, we can now log in into Turso platform, simply by running the <code>auth signup</code> or <code>auth login</code> to Register or Log-in.</p>
<pre><code>turso auth signup

# OR

turso auth login
</code></pre>
<p>This will redirect to the browser for the OAuth flow, once signed up and logged in, this will allow to interact with the Turso platform with the CLI that we downloaded.</p>
<p>To make sure we are logged in as the correct user, we can run the <code>auth whoami</code> command to get the currently logged-in user.</p>
<pre><code>turso auth whoami
</code></pre>
<p>This will print the username if you are logged-in. If everything seems correct, we can move ahead with the database creation step.</p>
<h3>Creating a Remote LibSQL Database on Turso</h3>
<p>To create a LibSQL database hosted on Turso, we will use the <code>turso db create</code> command.</p>
<pre><code>turso db create

# OR

turso db create &lt;name&gt;
</code></pre>
<p>This will create a database with the specified name, even if you don't provide the name, it will give out a random friendly two-word name to your database. It will create a database on the nearest location available from your location.</p>
<p>This command will output the following:</p>
<pre><code>Created database &lt;db-name&gt; at group default in 1.99s.

Start an interactive SQL shell with the following:
    turso db shell &lt;db-name&gt;

To see information about the database, including a connection URL, run:
    turso db show &lt;db-name&gt;

To get an authentication token for the database, run:
    turso db tokens create &lt;db-name&gt;
</code></pre>
<p>The next step, it shows to start an interactive shell, to see information about the database, and to generate an authentication token for the database.</p>
<p>We will move to the next part, which would be to create an authentication token for accessing the database from an external application.</p>
<h3>Generating and Storing Authentication Token for LibSQL Database</h3>
<p>After we executed the <code>db create</code> command, and it created the database on the Turso cloud, there was a command hint for creating a <code>token</code> with the command <code>db tokens create</code></p>
<p>So, this command will create a JWT authentication token, that will be used to connect and read/write to the database.</p>
<pre><code class="language-bash">turso db tokens create &lt;db-name&gt;

# OR

turso db tokens create &lt;db-name&gt; --read-only

# OR

turso db tokens create &lt;db-name&gt; --expiration 30d
</code></pre>
<p>We can use the simple <code>db tokens create &lt;db-name&gt;</code> to create an authentication token for the database with (read + write access). You can copy that returned token into a environment variable, or wherever your application can read that token.</p>
<p>This could be stored in the environment variable as follows:</p>
<pre><code class="language-bash">export TURSO_AUTH_TOKEN=&quot;&lt;token&gt;&quot;
</code></pre>
<p>To make a <code>read-only</code> token, we can use the flag <code>--read-only</code>. This will be handy, if you only have a database as a local replica, and the only purpose of the database is for querying data.
This will prevent any write operation on the database.</p>
<p>We can also use the <code>--expiration</code> flag that will be used to set the duration of the token. By default the value for expiry is <code>never</code>, but that could be a little too risky if you are making a serious application. You can either set it to <code>7d</code> which will make the token expire after seven days.</p>
<p>Now, we can get the remote database URL and connect to the database. The URL could be obtained by running the command <code>db show &lt;db-name&gt;</code></p>
<pre><code>turso db show &lt;db-name&gt;
</code></pre>
<p>This will output the following:</p>
<pre><code class="language-bash">Name:           &lt;db-name&gt;
URL:            libsql://&lt;db-name&gt;-&lt;username&gt;.turso.io
ID:             &lt;db-id&gt;   
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
</code></pre>
<p>The above output shows the meta-information of the database. This also has the URL hosted on Turso. We can construct the URL using the name of the database and your username as <code>libsql://&lt;db-name&gt;-&lt;username&gt;.turso.io</code>, you can set this in an environment variable or in the configuration wherever you can access it from the application.</p>
<p>To set the URL of the database in your application, you can use the <code>TURSO_DB_URL</code> environment variable.</p>
<pre><code class="language-bash">export TURSO_DATABASE_URL=&quot;libsql://&lt;db-name&gt;-&lt;username&gt;.turso.io&quot;
</code></pre>
<p>So, we have a remote database URL, and the access token configured, these are the two pieces that we will need to connect, read and write to the libsql database.</p>
<h3>Installing LibSQL Client for Golang</h3>
<p>So, we can install the <a href="https://pkg.go.dev/github.com/tursodatabase/libsql-client-go/libsql">libsql-client-go</a> package for Golang which will be used as an SDK in Golang to interact with a remote LibSQL database.</p>
<pre><code class="language-bash">go get github.com/tursodatabase/libsql-client-go/libsql
</code></pre>
<p>This will install the package <code>libsql</code> into the golang module. Now, we can use this in our golang application.</p>
<h3>Populating the LibSQL Database</h3>
<p>Moving ahead, we have a database, but it doesn't have data! So let's create some tables and insert some rows. We can use the <code>db shell</code> command to open an interactive SQL shell on a remote LibSQL database.</p>
<pre><code class="language-bash">turso db shell libsql://&lt;db-name&gt;-&lt;username&gt;.turso.io
</code></pre>
<p>This will be a default <code>sqlite3</code> like a shell, where we can execute SQL commands, like <code>.schema</code>, <code>.mode</code>, <code>.tables</code>, etc.</p>
<pre><code class="language-bash">  .dump       Render database content as SQL
  .help       List of all available commands.
  .indexes    List indexes in a table or database
  .mode       Set output mode
  .quit       Exit this program
  .read       Execute commands from a file
  .schema     Show table schemas.
  .tables     List all existing tables in the database.
</code></pre>
<p>And definitely, we can use the normal SQL queries, to read, write and delete data from the database.</p>
<h4>Creating a Table</h4>
<p>First, let's create a simple table, called <code>posts</code> with columns like <code>id</code>, <code>title</code> as a <code>VARCHAR(100)</code>, <code>description</code> as a <code>VARCHAR(255)</code>, and <code>content</code> as <code>TEXT</code> which won't be <code>NULL</code>.</p>
<pre><code class="language-sql">CREATE TABLE posts
  (
     id          INTEGER PRIMARY KEY,
     title       VARCHAR(100),
     description VARCHAR(255),
     content     TEXT NOT NULL
  ); 
</code></pre>
<p>This will create a table <code>posts</code> on the LibSQL database, yes this will mutate the primary LibSQL database which is hosted on Turso.</p>
<h4>Inserting Rows</h4>
<p>Now, since we have the <code>posts</code> table, we will insert some rows into the table.</p>
<pre><code class="language-bash">INSERT INTO posts (title, description, content)
VALUES 
    ('test title', 'test description', 'test content');
</code></pre>
<p>Now, we have some rows populated in the <code>posts</code> table. We can add more tables, and rows into the database, as usual, but this is just an example, so I'll keep it short.</p>
<h3>Connecting to the LibSQL Database</h3>
<p>Now, we have something to query from a database, after we connect to the database via the Golang application program.</p>
<p>First, we will grab two pieces to connect to the database.</p>
<pre><code class="language-bash">export TURSO_DATABASE_URL=&quot;libsql://&lt;db-name&gt;-&lt;username&gt;.turso.io&quot;
export TURSO_AUTH_TOKEN=&quot;&lt;token&gt;&quot;
</code></pre>
<p>Now, let's start with the Golang program code.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
    &quot;os&quot;
)

func main() {
    dbURL := os.Getenv(&quot;TURSO_DATABASE_URL&quot;)
    dbToken := os.Getenv(&quot;TURSO_AUTH_TOKEN&quot;)
    dbUrl := fmt.Sprintf(&quot;%s?authToken=%s&quot;, dbURL, dbToken)
}
</code></pre>
<p>This will be the basic config to grab the database URL and the authentication token, then we can construct the <code>dbURL</code> along with <code>dbToken</code> to construct the complete dbURL which will allow to access the database.</p>
<p>Moving ahead, we will import <code>database/sql</code> package that will be used to open the database connection and <code>github.com/tursodatabase/libsql-client-go/libsql</code> to connect to the remote LibSQL database.</p>
<pre><code class="language-go">package main

import (
	&quot;database/sql&quot;
    &quot;fmt&quot;
    &quot;os&quot;

	_ &quot;github.com/tursodatabase/libsql-client-go/libsql&quot;
)

func main() {
	dbURL := os.Getenv(&quot;TURSO_DATABASE_URL&quot;)
	dbToken := os.Getenv(&quot;TURSO_AUTH_TOKEN&quot;)
	dbUrl := fmt.Sprintf(&quot;%s?authToken=%s&quot;, dbURL, dbToken)

	db, err := sql.Open(&quot;libsql&quot;, dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, &quot;failed to open db %s: %s&quot;, dbUrl, err)
		os.Exit(1)
	}
	defer db.Close()

}
</code></pre>
<p>The <code>sql.Open</code> function will open the connection to the database, this will return a <code>sql.DB</code> object. The driver selected is <code>libsql</code> and the <code>dbURL</code> is the entire URL along with the authentication token.</p>
<pre><code class="language-go">type Post struct {
	Id          int
	Title       string
	Description string
	Content     string
}

rows, err := db.Query(&quot;SELECT * FROM posts&quot;)
if err != nil {
    fmt.Fprintf(os.Stderr, &quot;failed to query: %s&quot;, err)
    os.Exit(1)
}

for rows.Next() {
    var post Post
    if err := rows.Scan(&amp;post.Id, &amp;post.Title, &amp;post.Description, &amp;post.Content); err != nil {
        fmt.Fprintf(os.Stderr, &quot;failed to scan: %s&quot;, err)
        os.Exit(1)
    }
    fmt.Println(post)
}
defer rows.Close()
</code></pre>
<p>Now, let's query some data from the database. We can construct the <code>Post</code> struct that will be used to grab the required fields like <code>Id</code>, <code>Title</code>, <code>Description</code>, and <code>Content</code> from the <code>posts</code> table in the database.</p>
<p>Then, we will use the <code>db.Query</code> function to query the database. This function takes in a query and returns a <code>sql.Rows</code> object. We will iterate over all the <code>rows</code> returned from the database, with the <code>rows.Next()</code> that will fetch each row. Then we can <code>row.Scan</code> the row object with the appropriate and respective fields returned in the row. In this case, the <code>Id</code>, <code>Title</code>, <code>Description</code>, and the <code>Content</code> is fetched and stored into the <code>post</code> fields.</p>
<p>We have fetched the rows and we can do operations on them as required, this was just a basic example. So the entire code can be found below.</p>
<pre><code class="language-go">package main

import (
	&quot;database/sql&quot;
	&quot;fmt&quot;
	&quot;os&quot;

	_ &quot;github.com/tursodatabase/libsql-client-go/libsql&quot;
)

type Post struct {
	Id          int
	Title       string
	Description string
	Content     string
}

func main() {
	dbURL := os.Getenv(&quot;TURSO_DATABASE_URL&quot;)
	dbToken := os.Getenv(&quot;TURSO_AUTH_TOKEN&quot;)
	dbUrl := fmt.Sprintf(&quot;%s?authToken=%s&quot;, dbURL, dbToken)

	db, err := sql.Open(&quot;libsql&quot;, dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, &quot;failed to open db %s: %s&quot;, dbUrl, err)
		os.Exit(1)
	}
	defer db.Close()

	rows, err := db.Query(&quot;SELECT * FROM posts&quot;)
	if err != nil {
		fmt.Fprintf(os.Stderr, &quot;failed to query: %s&quot;, err)
		os.Exit(1)
	}

	for rows.Next() {
		var post Post
		if err := rows.Scan(&amp;post.Id, &amp;post.Title, &amp;post.Description, &amp;post.Content); err != nil {
			fmt.Fprintf(os.Stderr, &quot;failed to scan: %s&quot;, err)
			os.Exit(1)
		}
		fmt.Println(post)
	}
	defer rows.Close()

}
</code></pre>
<p>The output of the above code will result in all the rows present in the post table of the LibSQL database.</p>
<pre><code class="language-bash">$ go run remote.go

{1 test title test description test content}
{2 test title test description test content}
{3 sample post libsql tutorial create db, connect, create tables, insert rows, sync}
{4 test title test description test content}
</code></pre>
<p>I have added a few more rows to the post table, as you can see we have successfully connected, inserted, and read from the post table in the LibSQL database hosted on Turso.</p>
<p>For all the code related to this article, you can check out the <a href="https://github.com/mr-destructive/lets-go-with-turso">Let's Go with Turso</a> GitHub repo for all the examples and additional examples for using LibSQL with Golang.</p>
<h2>Conclusion</h2>
<p>So, that is a wrap for this part of the series, we have explored how to connect a remote LibSQL database hosted on Turso with Golang. In the next part of the series, we will explore how to create embedded replicas on Turso's LibSQL database in Golang.</p>
<p>Thank you for reading this post, If you have any questions, feedback, and suggestions, feel free to drop them in the comments.</p>
<p>Happy Coding :)</p>
