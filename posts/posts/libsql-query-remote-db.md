{
  "title": "LibSQL: Query a remote Turso database with cURL",
  "status": "published",
  "slug": "libsql-query-remote-db",
  "date": "2025-04-05T12:33:29Z"
}

<p>If you are using a local <a href="https://turso.tech/libsql">libsql</a> database, it is quite easy to query the database, but for a remote or a database on a hosted cloud platform like <a href="https://turso.tech/">turso</a>, we can use other clients or the api itself to query the data.</p>
<p>We can use the turso cli to get the authentication token for the database and then query the database.</p>
<h2>Turso CLI</h2>
<p>Using the <a href="https://docs.turso.tech/reference/turso-cli">turso-cli</a> to access the turso platform. We will use turso cli to create a libsql database, create authentication tokens, and query the db.</p>
<h3>Create a database (it's optional, you might already have a database)</h3>
<pre><code class="language-bash">turso db create
</code></pre>
<p>You will get a database on the turso cloud platform with some random interesting name like a passphrase.</p>
<p>Use the command <code>turso db list</code> and copy the URL</p>
<pre><code class="language-graphql">DB_URL=dbname-orgname.turso.io
</code></pre>
<h3>Create an authentication token for a particular database</h3>
<pre><code class="language-bash">turso db tokens create db_name
</code></pre>
<p>Copy the JWT token and this will be used as a authentication token when accessing the remote database in the turso cloud.</p>
<pre><code class="language-bash">TOKEN=abcdef.12345.wxyz
DB_URL=dbname-orgname.turso.io
</code></pre>
<ul>
<li>Querying the database using curl or other <a href="https://docs.turso.tech/libsql/client-access">api clients</a></li>
</ul>
<pre><code class="language-bash">curl -s -H &quot;Authorization: bearer $TOKEN&quot; \
     -d '{&quot;statements&quot;: [&quot;SELECT name FROM sqlite_master WHERE type=\&quot;table\&quot;;&quot;]}' \
     $DB_URL
</code></pre>
<p>We can use <code>curl</code> or any api client tools to send queries to the database hosted on the turso platform. We need to provide the JWT token in the <code>Authorization</code> header to connect to that particular database. The request's body is a JSON string with a list of statements to query the database.</p>
<pre><code class="language-graphql">[
    {
        &quot;results&quot;:
            {
                &quot;columns&quot;: [&quot;name&quot;],
                &quot;rows&quot;:[
                    [&quot;libsql_wasm_func_table&quot;], [&quot;_litestream_seq&quot;], [&quot;_litestream_lock&quot;], [&quot;sqlite_sequence&quot;], [&quot;user&quot;]
                ]
            }
     }
]
</code></pre>
<p>The result is a list of key-value pairs as <code>columns</code> and <code>rows</code> for each of the statements in the body. The columns are a list of column names requested in the query, and the rows are a list of rows where each row is a list of field values from the query.</p>
