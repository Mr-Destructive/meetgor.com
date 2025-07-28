{
  "title": "PGCLI: Postgres from the terminal",
  "status": "published",
  "slug": "pgcli-pg-from-terminal",
  "date": "2025-04-05T12:33:37Z"
}

<h2>Introduction</h2>
<p>Have you ever used the Postgres database and did you know you don't have to launch PGAdmin every time you want to write SQL queries, you can write down those queries even inside a terminal with a python package. It's PGCLI, that can act as an editor window for writing SQL queries for Postgres Databases. We can simply execute SQL queries in a PSQL-like shell with additional features like autocompletion, text-editing modes, file input/output options, etc.</p>
<p>If you are like CLIs and love to play with backend systems like Postgres(database) then chances are you will love this article. This article will cover the basics of using the PGCLI tool to perform simply as well as advanced options in interacting with your Postgres database.</p>
<h2>PGCLI</h2>
<p>PGCLi is a python package that acts as a CLI for executing SQL queries in a Postgres database. PGCLI allows us to interact with any Postgres database via the command line, it can be a remote database or a local database, you can access it right away from your terminal. The package acts as a CLI for accessing and executing any SQL queries, so we can simply provide the options/parameters to plug the database in and access the schema and its related data.</p>
<h2>Installation</h2>
<p>There are a lot of ways pgcli can be installed, my preferred way is with <code>pipx</code>, it just works well and you don't have to install it every time or mess up your global python packages environment.</p>
<pre><code>pipx install pgcli
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659878333/blog-media/pgcli-install-pipx.png" alt="PGCLI Install with pipx"></p>
<p>We don't even require installing it when using pipx, as it will each time create a fresh instance of the pgcli in a virtual environment.</p>
<p>We can now run the <code>pgcli</code> with parameters using the <code>pipx run</code> command. In the below command, we have used the <code>--help</code> options to list out all the commands and other options available with the <code>pgcli</code> command.</p>
<pre><code>pipx run pgcli --help
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659878322/blog-media/pgcli-help.png" alt="PGCLI Help"></p>
<p>If this command is working fine, it means we are ready to connect and interact with a database with pgcli.</p>
<h2>List local databases</h2>
<p>We can list all the databases in our local system by providing the hostname and the username, in my case, I have to provide the <code>localhost</code> and <code>postgres</code> as the hostname and the username respectively. Further, by providing the <code>--list</code> option, we can list down all the available databases on our system, basically for a given host and user.</p>
<pre><code>pipx run pgcli --host localhost --user postgres --list

OR

pipx run pgcli -h localhost -U postgres -l
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659878596/blog-media/pgcli-list-db.png" alt="PGCLI List all databases"></p>
<p>As, we can see, the command lists out all the available databases on the local server. This can help in getting an overview of the Postgres databases present in your local system. You need to enter the password for the username to access the database.</p>
<h2>Connect to a database</h2>
<p>To connect to an individual Postgres database we have two options, we can either use a URL string of postgres database or parse parameters to the command individually. It's easier to use the URL option because it wraps a lot of things in a single string rather than writing each parameters.</p>
<h3>Connect with the Postgres Connection URI</h3>
<p>We can simply pass the postgres connection URI for accessing a postgres database with pgcli. The URI has a specific pattern in the order where we need to specify the values like hostname, password, username, port, and database name.</p>
<pre><code>postgresql://username:password@hostname:port/database_name
</code></pre>
<p>The above is the structure for a postgres connection URI, we parse in the <code>username</code>, <code>password</code> for that <code>username</code>, <code>hostname</code> with the <code>port</code>, and the <code>database name</code>. These parameters if parsed correctly can directly access the database.</p>
<pre><code>pipx run pgcli postgresql://postgres:postgres@localhost:5432/techstructive_blog
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659878955/blog-media/pgcli-connect-uri.png" alt="PGCLI Postgres Connection URI"></p>
<p>This command will leave us in a prompt where we are basically in a PSQL shell in the provided database or host. For connecting to a local database, the hostname will likely be <code>localhost</code> and postgres database port is generally <code>5432</code> there might be exceptions to this. We can quit out of PGCLI with <code>CTRL + D</code> or <code>exit</code>.</p>
<p>While using a remote database, the hostname and other details might be provided, for instance, we might have a Django application, in that project, we will have <code>database_name</code>, <code>port</code>, and details mentioned in the <code>DATABASES</code> settings, those fields can be checked and a postgres connection URI might be constructed.</p>
<p>Let's take an example.</p>
<pre><code class="language-python">DATABASES = {
    &quot;default&quot;: {
        &quot;ENGINE&quot;: &quot;django.db.backends.postgresql&quot;,
        &quot;NAME&quot;: &quot;blog&quot;,
        &quot;USER&quot;: &quot;meet&quot;,
        &quot;PASSWORD&quot;: &quot;postgres&quot;,
        &quot;HOST&quot;: &quot;localhost&quot;,
    }
}
</code></pre>
<p>We have all the fields provided in the django project' settings file. We can use this to construct a postgres database connection URI.</p>
<pre><code>pgcli postgresql://meet:postgres@localhost:5432/blog
</code></pre>
<p>So, that is how you can extract parameters from other technologies and frameworks for working with postgres URIs.</p>
<h3>Connect with CLI options/paramters</h3>
<p>We can parse the options like <code>hostname</code>, <code>password</code>, and <code>database name</code> manually with the CLI options that PGCLI provides.</p>
<pre><code>pipx run pgcli --username postgres --host localhost --port 5432 --dbname techstructive_blog

OR

pipx run pgcli -u postgres -h localhost -p 5432 -d techstructive_blog
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659879642/blog-media/pgcli-connect-paramters.png" alt="PGCLI Connect Paramaters"></p>
<p>You will be prompted for a password for the database, after entering the password, you will be inside the psql shell. In there, we can execute SQL queries and other specific psql or pgcli commands.</p>
<p>Once we have connected to a database, we can now try to work with the SQL queries and explore some of the commands and features provided by pgcli.</p>
<h2>List out all databases if connected to a server</h2>
<p>Let's say we don't know the database name, we are just connected to a database server, we can get the list of the databases inside the shell with the <code>\l+</code> command. This command is similar to the <code>--list</code> parameter in the PGCLI command, but the only difference is that we are executing it from the psql shell rather than our terminal shell.</p>
<pre><code># Let's say we connected to a database server

pipx run pgcli --username postgres --host localhost --port 5432


# List out all databases in that server

\l+
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659879988/blog-media/pgcli-list-databases.png" alt="PGCLI List databases in a server"></p>
<p>So, that's how we would get the gist of all the databases inside the database server to which we are connected using the <code>\l+</code> command.</p>
<h2>List all tables in the database</h2>
<p>We can get a list of all the tables existing in the database we are logged in with the <code>\dt</code> command, we can simply enter the command in the prompt we are directed to once we are logged in to the postgres database.</p>
<pre><code>\dt
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659880232/blog-media/pgcli-list-tables.png" alt="PGCLI List tables in a database"></p>
<p>We can use <code>\d &lt;table_name&gt;</code> to describe the details of the table provided. We get back the list of the attributes inside the table and the indexes of the relationships of the table as well as the referenced table details.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659880453/blog-media/pgcli-describe-table.png" alt="PGCLI describe a table in a database"></p>
<p>We also have other commands for displaying types, schemas, roles, views, etc. with the <code>\d</code> prefix, a detailed list can be found in the <a href="https://www.pgcli.com/commands">PGCLI documentation</a>.</p>
<h2>Executing queries</h2>
<p>We can execute SQL queries in the PGCLI prompt as normally we do in a PSQL shell. There are added benefits like certain modes of editing text and auto-completion. We'll execute a few queries and also try out other options for editing and saving queries.</p>
<p>Normally for simple and easy tasks, we can write SQL queries right in the PSQL shell, it might be small and simple to construct a query. Let's take the simplest example to get all the attribute lists from a specific table in a database.</p>
<pre><code class="language-SQL">SELECT * FROM table_name
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659880794/blog-media/pgcli-sql-query.png" alt="PGCLI executing SQL queries"></p>
<p>We get an excellent table display of output which we can view by scroll or arrow keys. We can even use <code>j</code> and <code>k</code> for moving vertically in an output table.</p>
<h3>Wriing modes in PGCLI</h3>
<p>We can use two modes in PGCLI prompt, one is the Emacs mode which is the default one and the other is Vi-mode for Vim fanboys. We can switch between these modes with <code>F4</code> key. The <code>Emacs-mode</code> is just a normal typing experience without any modes or shortcut macros and other stuff. Though it's nice that PGCLI offers a VI-mode, it is great for people using Vim, when using the Shell for editing a Query, this mode can be utilized.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659882586/blog-media/pgcli-vi-mode-demo.gif" alt="PGCLI VI-mode"></p>
<p>We can use <code>hjkl</code> keybinding for movement and other normal mode keymaps like <code>w</code> for moving a word, or <code>b</code> for moving a word in the backward direction, <code>e</code> to the end of the word, etc. We can use <code>i</code> or <code>a</code> for entering into insert mode, <code>dd</code> to delete the entire line, <code>cc</code> or <code>S</code> to delete the current line and enter into insert mode, <code>yy</code> to yank line, and <code>p</code> to paste, <code>u</code> for undo, etc.</p>
<p>We can even execute shell commands by using the <code>!</code> in normal mode where we will be prompted for the command executed in the same window.</p>
<h3>Autocompletion in PGCLI shell</h3>
<p>PGCLI shell has an auto-completion feature that will provide some suggestions based on the recent queries or commands executed or the default options. You can toggle the autocompletion with the <code>F2</code> key, it is officially called as <code>Smart-Completion</code> as you would see in the bottom left corner of the shell.</p>
<h3>Select an Editor to write queries</h3>
<p>We can use different editors for writing SQL queries. We can write a query from an editor by using the <code></code> command. The default editor is chosen which will be set in one of your <code>$EDITOR</code> or <code>$VISUAL</code> environment variables. You can set them manually in the bashrc file by exporting the value of <code>VISUAL</code> or <code>EDITOR</code> as the name of your editor.</p>
<pre><code class="language-bash">export VISUAL=vim

OR

export EDITOR=vim

OR

export VISUAL=gedit
</code></pre>
<p>Here, we can set <code>VISUAL</code> environment variable as the name of the editor or the path to it. The <code>VISUAL</code> environment variable is preferred first. For the details, we can check out the <a href="https://www.pgcli.com/editor">editor command documentation</a>.</p>
<p>Once, we have configured the editor, we will be inside the editor after we enter the <code></code> command. We can write queries inside the editor and quit after we are done, once the editor is closed, the text is read from the editor and entered in the psql shell. This will make us wait for the enter command for executing the query.</p>
<p>We can even execute more than one query at a time if we use the <code>;</code> at the end of each query. Also, we can press <code>v</code> in VI-normal mode, to open the current query in the default editor.</p>
<h3>Save the query output to a file</h3>
<p>Let's say we have written a query inside the psql shell and we want the output to be saved in a file. We can achieve that with the <code>\o</code> command, we need to specify the filename in which the output of the queries will be saved.</p>
<p><strong>NOTE: Once we use the <code>\o</code> command, all the queries will start appending the results in the provided file. To stop the behavior, you can use <code>\o</code> (without any filename) to stop appending the results to a file.</strong></p>
<pre><code># Start recording the output of queries to file
\o filename


# Stop recording the output of queries to file
\o
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659884242/blog-media/pgcli-output-query.gif" alt="PGCLI Query output to file"></p>
<h3>Execute query from a file</h3>
<p>We can even execute queries stored in a file, we can use the <code>\i</code> command followed by the file name. This will load in the queries from the file and will display the output of each query.</p>
<pre><code>\i file_name
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659883225/blog-media/pgcli-file-query.gif" alt="PGCLI query from a file"></p>
<h2>Summary</h2>
<p>In this little article, we were able to explore the basics of the PGCLI package in python, with that package we can interact with the Postgres database that can be on your local server or somewhere remotely. With the PSQL-like shell provided by the PGCLI tool, we were able to write SQL queries, execute commands, write/read queries from files, etc. This tool is quite handy for people dealing with postgres databases and who need to view data or schema in that database.</p>
<p>I've personally used this tool in my current internship where I needed to create a few SQL queries to fetch particular data related to various tables and relationships. It saved a lot of time with auto-completion and integration with the text editor for writing longer queries, especially with JOINS and nested queries.</p>
<h2>Conclusion</h2>
<p>This was my take and experience of using the PGCLI tool and exploring the various features it provides that can be leveraged for quick interactions with a Postgres database. Hopefully, you have learned something from this post, if you have any queries(NOT SQL queries) or feedback, please let me know in the comments or on my social handles. Thank you for reading and Happy Coding :)</p>
