{
  "title": "Connecting LibSQL database with Python",
  "status": "published",
  "slug": "libsql-python",
  "date": "2025-04-05T12:33:29Z"
}

<h2>Introduction</h2>
<p>LibSQL is an <strong>Open Contribution</strong> fork of SQLite. Open Contribution means that it allows suggestions and contributions from the community as opposed to SQLite which is open source but doesn't accept community contributions.</p>
<h2>Installation of LibSQL Client(s)</h2>
<p>There are two libraries for LibSQL to interact with Python, the <a href="https://github.com/libsql/libsql-client-py/">libsql-client</a> and the <a href="https://github.com/libsql/libsql-experimental-python">libsql-experimental-python</a>. The former is the recommended client as it is stable, whereas the latter is in development and has the latest features from the libsql database engine, however, it is compatible with the <a href="https://docs.python.org/3/library/sqlite3.html">sqlite</a> module.</p>
<p>There are differences in how to connect and fetch responses in each of them, but if you want to get started quickly you can safely turn with the <code>libsql-client</code> as it is recommended from the official documentation. If you want to connect an existing SQLite database to libsql, you can turn up the <code>libsql-experimental-pyhon</code> package.</p>
<p>We will explore both of them in this article. Let's dive straight into the installation of the package.</p>
<h2>LibSQL Client</h2>
<p>So, Libsql-client is the Python package provided by Turso as a Python client to interact with a libsql/sqlite database.</p>
<h3>Installation</h3>
<p>The installation of a Python package is as simple as a pip install.</p>
<pre><code class="language-bash"> pip install libsql-client
</code></pre>
<h3>Connecting to the database</h3>
<p>Connecting to a libsql database is as simple as SQLite for the local database file. However, if you have a libsql database over the edge(Turso), you can use the API provided by SQLD to connect to that database.</p>
<h4>Connecting Local database file</h4>
<p>To connect to a simple libsql database file, you can either create a file as <code>mydb.db</code> or just move into the code straight away.</p>
<pre><code class="language-python">import libsql_client

client = libsql_client.create_client_sync(&quot;file:temp.db&quot;)
result = client.execute(&quot;SELECT 1;&quot;)

for row in result.rows:
    print(row)
client.close()
</code></pre>
<p>OR</p>
<pre><code class="language-python">import libsql_client

with libsql_client.create_client_sync(&quot;file:temp.db&quot;) as client:
    result = client.execute(&quot;SELECT 1;&quot;)
    
    for row in result.rows:
        print(row)
</code></pre>
<p>So, in the above code, the <code>client</code> object is a client of the libsql database. We use the <a href="https://libsql.org/libsql-client-py/reference.html#create_client_sync">create_client_sync</a> method that accepts a few parameters <code>url</code> i.e. the URL of the database we want to connect to, and the <code>auth_token</code> which we will see when we want to connect to a database on the edge.</p>
<p>There is also <a href="https://libsql.org/libsql-client-py/reference.html#create_client">create_client</a> method which is an async method. You can turn up to this method if you are looking for async connections.</p>
<p>We have used <code>file:temp.db</code> i.e. to connect a local file in the current directory named <code>temp.db</code> . This method will return a <a href="https://libsql.org/libsql-client-py/reference.html#Client">Client</a> object in this case a <a href="https://libsql.org/libsql-client-py/reference.html#ClientSync">ClientSync</a> object which is a wrapper around the <code>Client</code> object. We will be using the methods available for the Client class later while querying in detail. You can see that we have used the <a href="https://libsql.org/libsql-client-py/reference.html#ClientSync.execute">execute</a> method from the client object.</p>
<h4>Connecting to a cloud(turso) database</h4>
<p>To connect to a libsql database on the edge i.e. turso you need to specify the auth token(JWT) in the <code>create_client</code> method.</p>
<p>You can create the token for accessing the database with the</p>
<pre><code class="language-bash">turso db tokens create mydb-name
</code></pre>
<p>This command will create a <code>JWT_TOKEN</code> , store this token securely and this will be used to access the database which is hosted on the Turso cloud.</p>
<pre><code class="language-bash">export JWT_TOKEN=&quot;YOURTOKEN&quot;
OR
## Save it in the .env file 
## JWT_TOKEN=YOURTOKEN
</code></pre>
<p>This will store the token in the environment variable and can be later used from the local environment.</p>
<pre><code class="language-python">import libsql_client

with libsql_client.create_client_sync(
    &quot;libsql://dbname-orgname.turso.io&quot;,
    auth_token=&quot;secret.something.secret&quot;
) as client:
    result = client.execute(&quot;SELECT 1;&quot;)
    
    for row in result.rows:
        print(row)
</code></pre>
<p>The <code>auth_token</code> is the JWT Token that we created in the previous step, you can load the token as an Environment variable with the following code:</p>
<pre><code class="language-python">import os

JWT_TOKEN = os.environ.get(&quot;JWT_TOKEN&quot;)

## OR

# pip install python-dotenv
from dotenv import load_dotenv
load_dotenv()
JWT_TOKEN = os.environ.get(&quot;JWT_TOKEN&quot;)
</code></pre>
<p>This token can be used to authenticate the connection to the Turso database. The rest of the connection code remains as it is.</p>
<h3>Running Queries and Fetching results</h3>
<p>Now, that we have the connection, we can move into querying the database through the client connection.</p>
<h2>LibSQL-Client (SQLite compatible)</h2>
<h3>Installation</h3>
<p>The installation of the <a href="https://badge.fury.io/py/libsql-experimental">libsql-experimental-python</a> package is simple with the pip install command with the name of the package.</p>
<pre><code class="language-bash">pip install libsql-experimental
</code></pre>
<p>Now, we can move into connecting to the database with Python script.</p>
<h3>Connecting to the database</h3>
<p>Connecting to a libsql database is as simple as SQLite for the local database file.</p>
<h3>Syncing with a remote database</h3>
<h3>Running Queries and Fetching results</h3>
