---
article_html: "<h2>Introduction</h2>\n<p>Have you ever used the Postgres database
  and did you know you don't have to launch PGAdmin every time you want to write SQL
  queries, you can write down those queries even inside a terminal with a python package.
  It's PGCLI, that can act as an editor window for writing SQL queries for Postgres
  Databases. We can simply execute SQL queries in a PSQL-like shell with additional
  features like autocompletion, text-editing modes, file input/output options, etc.</p>\n<p>If
  you are like CLIs and love to play with backend systems like Postgres(database)
  then chances are you will love this article. This article will cover the basics
  of using the PGCLI tool to perform simply as well as advanced options in interacting
  with your Postgres database.</p>\n<h2>PGCLI</h2>\n<p>PGCLi is a python package that
  acts as a CLI for executing SQL queries in a Postgres database. PGCLI allows us
  to interact with any Postgres database via the command line, it can be a remote
  database or a local database, you can access it right away from your terminal. The
  package acts as a CLI for accessing and executing any SQL queries, so we can simply
  provide the options/parameters to plug the database in and access the schema and
  its related data.</p>\n<h2>Installation</h2>\n<p>There are a lot of ways pgcli can
  be installed, my preferred way is with <code>pipx</code>, it just works well and
  you don't have to install it every time or mess up your global python packages environment.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pipx install pgcli\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659878333/blog-media/pgcli-install-pipx.png\"
  alt=\"PGCLI Install with pipx\" /></p>\n<p>We don't even require installing it when
  using pipx, as it will each time create a fresh instance of the pgcli in a virtual
  environment.</p>\n<p>We can now run the <code>pgcli</code> with parameters using
  the <code>pipx run</code> command. In the below command, we have used the <code>--help</code>
  options to list out all the commands and other options available with the <code>pgcli</code>
  command.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pipx run pgcli --help\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659878322/blog-media/pgcli-help.png\"
  alt=\"PGCLI Help\" /></p>\n<p>If this command is working fine, it means we are ready
  to connect and interact with a database with pgcli.</p>\n<h2>List local databases</h2>\n<p>We
  can list all the databases in our local system by providing the hostname and the
  username, in my case, I have to provide the <code>localhost</code> and <code>postgres</code>
  as the hostname and the username respectively. Further, by providing the <code>--list</code>
  option, we can list down all the available databases on our system, basically for
  a given host and user.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pipx run pgcli --host localhost
  --user postgres --list\n\nOR\n\npipx run pgcli -h localhost -U postgres -l\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659878596/blog-media/pgcli-list-db.png\"
  alt=\"PGCLI List all databases\" /></p>\n<p>As, we can see, the command lists out
  all the available databases on the local server. This can help in getting an overview
  of the Postgres databases present in your local system. You need to enter the password
  for the username to access the database.</p>\n<h2>Connect to a database</h2>\n<p>To
  connect to an individual Postgres database we have two options, we can either use
  a URL string of postgres database or parse parameters to the command individually.
  It's easier to use the URL option because it wraps a lot of things in a single string
  rather than writing each parameters.</p>\n<h3>Connect with the Postgres Connection
  URI</h3>\n<p>We can simply pass the postgres connection URI for accessing a postgres
  database with pgcli. The URI has a specific pattern in the order where we need to
  specify the values like hostname, password, username, port, and database name.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>postgresql://username:password@hostname:port/database_name\n</pre></div>\n\n</pre>\n\n<p>The
  above is the structure for a postgres connection URI, we parse in the <code>username</code>,
  <code>password</code> for that <code>username</code>, <code>hostname</code> with
  the <code>port</code>, and the <code>database name</code>. These parameters if parsed
  correctly can directly access the database.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pipx run pgcli postgresql://postgres:postgres@localhost:5432/techstructive_blog\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659878955/blog-media/pgcli-connect-uri.png\"
  alt=\"PGCLI Postgres Connection URI\" /></p>\n<p>This command will leave us in a
  prompt where we are basically in a PSQL shell in the provided database or host.
  For connecting to a local database, the hostname will likely be <code>localhost</code>
  and postgres database port is generally <code>5432</code> there might be exceptions
  to this. We can quit out of PGCLI with <code>CTRL + D</code> or <code>exit</code>.</p>\n<p>While
  using a remote database, the hostname and other details might be provided, for instance,
  we might have a Django application, in that project, we will have <code>database_name</code>,
  <code>port</code>, and details mentioned in the <code>DATABASES</code> settings,
  those fields can be checked and a postgres connection URI might be constructed.</p>\n<p>Let's
  take an example.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">DATABASES</span> <span class=\"o\">=</span> <span class=\"p\">{</span>\n
  \   <span class=\"s2\">&quot;default&quot;</span><span class=\"p\">:</span> <span
  class=\"p\">{</span>\n        <span class=\"s2\">&quot;ENGINE&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;django.db.backends.postgresql&quot;</span><span
  class=\"p\">,</span>\n        <span class=\"s2\">&quot;NAME&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;blog&quot;</span><span class=\"p\">,</span>\n        <span
  class=\"s2\">&quot;USER&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;meet&quot;</span><span
  class=\"p\">,</span>\n        <span class=\"s2\">&quot;PASSWORD&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;postgres&quot;</span><span class=\"p\">,</span>\n
  \       <span class=\"s2\">&quot;HOST&quot;</span><span class=\"p\">:</span> <span
  class=\"s2\">&quot;localhost&quot;</span><span class=\"p\">,</span>\n    <span class=\"p\">}</span>\n<span
  class=\"p\">}</span>\n</pre></div>\n\n</pre>\n\n<p>We have all the fields provided
  in the django project' settings file. We can use this to construct a postgres database
  connection URI.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pgcli postgresql://meet:postgres@localhost:5432/blog\n</pre></div>\n\n</pre>\n\n<p>So,
  that is how you can extract parameters from other technologies and frameworks for
  working with postgres URIs.</p>\n<h3>Connect with CLI options/paramters</h3>\n<p>We
  can parse the options like <code>hostname</code>, <code>password</code>, and <code>database
  name</code> manually with the CLI options that PGCLI provides.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pipx run pgcli --username
  postgres --host localhost --port 5432 --dbname techstructive_blog\n\nOR\n\npipx
  run pgcli -u postgres -h localhost -p 5432 -d techstructive_blog\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659879642/blog-media/pgcli-connect-paramters.png\"
  alt=\"PGCLI Connect Paramaters\" /></p>\n<p>You will be prompted for a password
  for the database, after entering the password, you will be inside the psql shell.
  In there, we can execute SQL queries and other specific psql or pgcli commands.</p>\n<p>Once
  we have connected to a database, we can now try to work with the SQL queries and
  explore some of the commands and features provided by pgcli.</p>\n<h2>List out all
  databases if connected to a server</h2>\n<p>Let's say we don't know the database
  name, we are just connected to a database server, we can get the list of the databases
  inside the shell with the <code>\\l+</code> command. This command is similar to
  the <code>--list</code> parameter in the PGCLI command, but the only difference
  is that we are executing it from the psql shell rather than our terminal shell.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># Let&#39;s say we connected
  to a database server\n\npipx run pgcli --username postgres --host localhost --port
  5432\n\n\n# List out all databases in that server\n\n\\l+\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659879988/blog-media/pgcli-list-databases.png\"
  alt=\"PGCLI List databases in a server\" /></p>\n<p>So, that's how we would get
  the gist of all the databases inside the database server to which we are connected
  using the <code>\\l+</code> command.</p>\n<h2>List all tables in the database</h2>\n<p>We
  can get a list of all the tables existing in the database we are logged in with
  the <code>\\dt</code> command, we can simply enter the command in the prompt we
  are directed to once we are logged in to the postgres database.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>\\dt\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659880232/blog-media/pgcli-list-tables.png\"
  alt=\"PGCLI List tables in a database\" /></p>\n<p>We can use <code>\\d &lt;table_name&gt;</code>
  to describe the details of the table provided. We get back the list of the attributes
  inside the table and the indexes of the relationships of the table as well as the
  referenced table details.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659880453/blog-media/pgcli-describe-table.png\"
  alt=\"PGCLI describe a table in a database\" /></p>\n<p>We also have other commands
  for displaying types, schemas, roles, views, etc. with the <code>\\d</code> prefix,
  a detailed list can be found in the <a href=\"https://www.pgcli.com/commands\">PGCLI
  documentation</a>.</p>\n<h2>Executing queries</h2>\n<p>We can execute SQL queries
  in the PGCLI prompt as normally we do in a PSQL shell. There are added benefits
  like certain modes of editing text and auto-completion. We'll execute a few queries
  and also try out other options for editing and saving queries.</p>\n<p>Normally
  for simple and easy tasks, we can write SQL queries right in the PSQL shell, it
  might be small and simple to construct a query. Let's take the simplest example
  to get all the attribute lists from a specific table in a database.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>SQL</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">SELECT</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"w\"> </span><span class=\"k\">FROM</span><span class=\"w\"> </span><span
  class=\"k\">table_name</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659880794/blog-media/pgcli-sql-query.png\"
  alt=\"PGCLI executing SQL queries\" /></p>\n<p>We get an excellent table display
  of output which we can view by scroll or arrow keys. We can even use <code>j</code>
  and <code>k</code> for moving vertically in an output table.</p>\n<h3>Wriing modes
  in PGCLI</h3>\n<p>We can use two modes in PGCLI prompt, one is the Emacs mode which
  is the default one and the other is Vi-mode for Vim fanboys. We can switch between
  these modes with <code>F4</code> key. The <code>Emacs-mode</code> is just a normal
  typing experience without any modes or shortcut macros and other stuff. Though it's
  nice that PGCLI offers a VI-mode, it is great for people using Vim, when using the
  Shell for editing a Query, this mode can be utilized.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659882586/blog-media/pgcli-vi-mode-demo.gif\"
  alt=\"PGCLI VI-mode\" /></p>\n<p>We can use <code>hjkl</code> keybinding for movement
  and other normal mode keymaps like <code>w</code> for moving a word, or <code>b</code>
  for moving a word in the backward direction, <code>e</code> to the end of the word,
  etc. We can use <code>i</code> or <code>a</code> for entering into insert mode,
  <code>dd</code> to delete the entire line, <code>cc</code> or <code>S</code> to
  delete the current line and enter into insert mode, <code>yy</code> to yank line,
  and <code>p</code> to paste, <code>u</code> for undo, etc.</p>\n<p>We can even execute
  shell commands by using the <code>!</code> in normal mode where we will be prompted
  for the command executed in the same window.</p>\n<h3>Autocompletion in PGCLI shell</h3>\n<p>PGCLI
  shell has an auto-completion feature that will provide some suggestions based on
  the recent queries or commands executed or the default options. You can toggle the
  autocompletion with the <code>F2</code> key, it is officially called as <code>Smart-Completion</code>
  as you would see in the bottom left corner of the shell.</p>\n<h3>Select an Editor
  to write queries</h3>\n<p>We can use different editors for writing SQL queries.
  We can write a query from an editor by using the <code>\\e</code> command. The default
  editor is chosen which will be set in one of your <code>$EDITOR</code> or <code>$VISUAL</code>
  environment variables. You can set them manually in the bashrc file by exporting
  the value of <code>VISUAL</code> or <code>EDITOR</code> as the name of your editor.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">export</span> <span class=\"nv\">VISUAL</span><span class=\"o\">=</span>vim\n\nOR\n\n<span
  class=\"nb\">export</span> <span class=\"nv\">EDITOR</span><span class=\"o\">=</span>vim\n\nOR\n\n<span
  class=\"nb\">export</span> <span class=\"nv\">VISUAL</span><span class=\"o\">=</span>gedit\n</pre></div>\n\n</pre>\n\n<p>Here,
  we can set <code>VISUAL</code> environment variable as the name of the editor or
  the path to it. The <code>VISUAL</code> environment variable is preferred first.
  For the details, we can check out the <a href=\"https://www.pgcli.com/editor\">editor
  command documentation</a>.</p>\n<p>Once, we have configured the editor, we will
  be inside the editor after we enter the <code>\\e</code> command. We can write queries
  inside the editor and quit after we are done, once the editor is closed, the text
  is read from the editor and entered in the psql shell. This will make us wait for
  the enter command for executing the query.</p>\n<p>We can even execute more than
  one query at a time if we use the <code>;</code> at the end of each query. Also,
  we can press <code>v</code> in VI-normal mode, to open the current query in the
  default editor.</p>\n<h3>Save the query output to a file</h3>\n<p>Let's say we have
  written a query inside the psql shell and we want the output to be saved in a file.
  We can achieve that with the <code>\\o</code> command, we need to specify the filename
  in which the output of the queries will be saved.</p>\n<p><strong>NOTE: Once we
  use the <code>\\o</code> command, all the queries will start appending the results
  in the provided file. To stop the behavior, you can use <code>\\o</code> (without
  any filename) to stop appending the results to a file.</strong></p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># Start recording the output
  of queries to file\n\\o filename\n\n\n# Stop recording the output of queries to
  file\n\\o\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659884242/blog-media/pgcli-output-query.gif\"
  alt=\"PGCLI Query output to file\" /></p>\n<h3>Execute query from a file</h3>\n<p>We
  can even execute queries stored in a file, we can use the <code>\\i</code> command
  followed by the file name. This will load in the queries from the file and will
  display the output of each query.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>\\i file_name\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659883225/blog-media/pgcli-file-query.gif\"
  alt=\"PGCLI query from a file\" /></p>\n<h2>Summary</h2>\n<p>In this little article,
  we were able to explore the basics of the PGCLI package in python, with that package
  we can interact with the Postgres database that can be on your local server or somewhere
  remotely. With the PSQL-like shell provided by the PGCLI tool, we were able to write
  SQL queries, execute commands, write/read queries from files, etc. This tool is
  quite handy for people dealing with postgres databases and who need to view data
  or schema in that database.</p>\n<p>I've personally used this tool in my current
  internship where I needed to create a few SQL queries to fetch particular data related
  to various tables and relationships. It saved a lot of time with auto-completion
  and integration with the text editor for writing longer queries, especially with
  JOINS and nested queries.</p>\n<h2>Conclusion</h2>\n<p>This was my take and experience
  of using the PGCLI tool and exploring the various features it provides that can
  be leveraged for quick interactions with a Postgres database. Hopefully, you have
  learned something from this post, if you have any queries(NOT SQL queries) or feedback,
  please let me know in the comments or on my social handles. Thank you for reading
  and Happy Coding :)</p>\n"
cover: ''
date: 2022-08-07
datetime: 2022-08-07 00:00:00+00:00
description: Running SQL queries in a Postgres database with PGCLI. Exploring the
  python package PGCLI, that can run SQL queries for a postgres database environment
  from the command line.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2022-08-07-PGCLI-postgres-from-terminal.md
html: "<h2>Introduction</h2>\n<p>Have you ever used the Postgres database and did
  you know you don't have to launch PGAdmin every time you want to write SQL queries,
  you can write down those queries even inside a terminal with a python package. It's
  PGCLI, that can act as an editor window for writing SQL queries for Postgres Databases.
  We can simply execute SQL queries in a PSQL-like shell with additional features
  like autocompletion, text-editing modes, file input/output options, etc.</p>\n<p>If
  you are like CLIs and love to play with backend systems like Postgres(database)
  then chances are you will love this article. This article will cover the basics
  of using the PGCLI tool to perform simply as well as advanced options in interacting
  with your Postgres database.</p>\n<h2>PGCLI</h2>\n<p>PGCLi is a python package that
  acts as a CLI for executing SQL queries in a Postgres database. PGCLI allows us
  to interact with any Postgres database via the command line, it can be a remote
  database or a local database, you can access it right away from your terminal. The
  package acts as a CLI for accessing and executing any SQL queries, so we can simply
  provide the options/parameters to plug the database in and access the schema and
  its related data.</p>\n<h2>Installation</h2>\n<p>There are a lot of ways pgcli can
  be installed, my preferred way is with <code>pipx</code>, it just works well and
  you don't have to install it every time or mess up your global python packages environment.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pipx install pgcli\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659878333/blog-media/pgcli-install-pipx.png\"
  alt=\"PGCLI Install with pipx\" /></p>\n<p>We don't even require installing it when
  using pipx, as it will each time create a fresh instance of the pgcli in a virtual
  environment.</p>\n<p>We can now run the <code>pgcli</code> with parameters using
  the <code>pipx run</code> command. In the below command, we have used the <code>--help</code>
  options to list out all the commands and other options available with the <code>pgcli</code>
  command.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pipx run pgcli --help\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659878322/blog-media/pgcli-help.png\"
  alt=\"PGCLI Help\" /></p>\n<p>If this command is working fine, it means we are ready
  to connect and interact with a database with pgcli.</p>\n<h2>List local databases</h2>\n<p>We
  can list all the databases in our local system by providing the hostname and the
  username, in my case, I have to provide the <code>localhost</code> and <code>postgres</code>
  as the hostname and the username respectively. Further, by providing the <code>--list</code>
  option, we can list down all the available databases on our system, basically for
  a given host and user.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pipx run pgcli --host localhost
  --user postgres --list\n\nOR\n\npipx run pgcli -h localhost -U postgres -l\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659878596/blog-media/pgcli-list-db.png\"
  alt=\"PGCLI List all databases\" /></p>\n<p>As, we can see, the command lists out
  all the available databases on the local server. This can help in getting an overview
  of the Postgres databases present in your local system. You need to enter the password
  for the username to access the database.</p>\n<h2>Connect to a database</h2>\n<p>To
  connect to an individual Postgres database we have two options, we can either use
  a URL string of postgres database or parse parameters to the command individually.
  It's easier to use the URL option because it wraps a lot of things in a single string
  rather than writing each parameters.</p>\n<h3>Connect with the Postgres Connection
  URI</h3>\n<p>We can simply pass the postgres connection URI for accessing a postgres
  database with pgcli. The URI has a specific pattern in the order where we need to
  specify the values like hostname, password, username, port, and database name.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>postgresql://username:password@hostname:port/database_name\n</pre></div>\n\n</pre>\n\n<p>The
  above is the structure for a postgres connection URI, we parse in the <code>username</code>,
  <code>password</code> for that <code>username</code>, <code>hostname</code> with
  the <code>port</code>, and the <code>database name</code>. These parameters if parsed
  correctly can directly access the database.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pipx run pgcli postgresql://postgres:postgres@localhost:5432/techstructive_blog\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659878955/blog-media/pgcli-connect-uri.png\"
  alt=\"PGCLI Postgres Connection URI\" /></p>\n<p>This command will leave us in a
  prompt where we are basically in a PSQL shell in the provided database or host.
  For connecting to a local database, the hostname will likely be <code>localhost</code>
  and postgres database port is generally <code>5432</code> there might be exceptions
  to this. We can quit out of PGCLI with <code>CTRL + D</code> or <code>exit</code>.</p>\n<p>While
  using a remote database, the hostname and other details might be provided, for instance,
  we might have a Django application, in that project, we will have <code>database_name</code>,
  <code>port</code>, and details mentioned in the <code>DATABASES</code> settings,
  those fields can be checked and a postgres connection URI might be constructed.</p>\n<p>Let's
  take an example.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">DATABASES</span> <span class=\"o\">=</span> <span class=\"p\">{</span>\n
  \   <span class=\"s2\">&quot;default&quot;</span><span class=\"p\">:</span> <span
  class=\"p\">{</span>\n        <span class=\"s2\">&quot;ENGINE&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;django.db.backends.postgresql&quot;</span><span
  class=\"p\">,</span>\n        <span class=\"s2\">&quot;NAME&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;blog&quot;</span><span class=\"p\">,</span>\n        <span
  class=\"s2\">&quot;USER&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;meet&quot;</span><span
  class=\"p\">,</span>\n        <span class=\"s2\">&quot;PASSWORD&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;postgres&quot;</span><span class=\"p\">,</span>\n
  \       <span class=\"s2\">&quot;HOST&quot;</span><span class=\"p\">:</span> <span
  class=\"s2\">&quot;localhost&quot;</span><span class=\"p\">,</span>\n    <span class=\"p\">}</span>\n<span
  class=\"p\">}</span>\n</pre></div>\n\n</pre>\n\n<p>We have all the fields provided
  in the django project' settings file. We can use this to construct a postgres database
  connection URI.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pgcli postgresql://meet:postgres@localhost:5432/blog\n</pre></div>\n\n</pre>\n\n<p>So,
  that is how you can extract parameters from other technologies and frameworks for
  working with postgres URIs.</p>\n<h3>Connect with CLI options/paramters</h3>\n<p>We
  can parse the options like <code>hostname</code>, <code>password</code>, and <code>database
  name</code> manually with the CLI options that PGCLI provides.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>pipx run pgcli --username
  postgres --host localhost --port 5432 --dbname techstructive_blog\n\nOR\n\npipx
  run pgcli -u postgres -h localhost -p 5432 -d techstructive_blog\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659879642/blog-media/pgcli-connect-paramters.png\"
  alt=\"PGCLI Connect Paramaters\" /></p>\n<p>You will be prompted for a password
  for the database, after entering the password, you will be inside the psql shell.
  In there, we can execute SQL queries and other specific psql or pgcli commands.</p>\n<p>Once
  we have connected to a database, we can now try to work with the SQL queries and
  explore some of the commands and features provided by pgcli.</p>\n<h2>List out all
  databases if connected to a server</h2>\n<p>Let's say we don't know the database
  name, we are just connected to a database server, we can get the list of the databases
  inside the shell with the <code>\\l+</code> command. This command is similar to
  the <code>--list</code> parameter in the PGCLI command, but the only difference
  is that we are executing it from the psql shell rather than our terminal shell.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># Let&#39;s say we connected
  to a database server\n\npipx run pgcli --username postgres --host localhost --port
  5432\n\n\n# List out all databases in that server\n\n\\l+\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659879988/blog-media/pgcli-list-databases.png\"
  alt=\"PGCLI List databases in a server\" /></p>\n<p>So, that's how we would get
  the gist of all the databases inside the database server to which we are connected
  using the <code>\\l+</code> command.</p>\n<h2>List all tables in the database</h2>\n<p>We
  can get a list of all the tables existing in the database we are logged in with
  the <code>\\dt</code> command, we can simply enter the command in the prompt we
  are directed to once we are logged in to the postgres database.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>\\dt\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659880232/blog-media/pgcli-list-tables.png\"
  alt=\"PGCLI List tables in a database\" /></p>\n<p>We can use <code>\\d &lt;table_name&gt;</code>
  to describe the details of the table provided. We get back the list of the attributes
  inside the table and the indexes of the relationships of the table as well as the
  referenced table details.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659880453/blog-media/pgcli-describe-table.png\"
  alt=\"PGCLI describe a table in a database\" /></p>\n<p>We also have other commands
  for displaying types, schemas, roles, views, etc. with the <code>\\d</code> prefix,
  a detailed list can be found in the <a href=\"https://www.pgcli.com/commands\">PGCLI
  documentation</a>.</p>\n<h2>Executing queries</h2>\n<p>We can execute SQL queries
  in the PGCLI prompt as normally we do in a PSQL shell. There are added benefits
  like certain modes of editing text and auto-completion. We'll execute a few queries
  and also try out other options for editing and saving queries.</p>\n<p>Normally
  for simple and easy tasks, we can write SQL queries right in the PSQL shell, it
  might be small and simple to construct a query. Let's take the simplest example
  to get all the attribute lists from a specific table in a database.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>SQL</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">SELECT</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"w\"> </span><span class=\"k\">FROM</span><span class=\"w\"> </span><span
  class=\"k\">table_name</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659880794/blog-media/pgcli-sql-query.png\"
  alt=\"PGCLI executing SQL queries\" /></p>\n<p>We get an excellent table display
  of output which we can view by scroll or arrow keys. We can even use <code>j</code>
  and <code>k</code> for moving vertically in an output table.</p>\n<h3>Wriing modes
  in PGCLI</h3>\n<p>We can use two modes in PGCLI prompt, one is the Emacs mode which
  is the default one and the other is Vi-mode for Vim fanboys. We can switch between
  these modes with <code>F4</code> key. The <code>Emacs-mode</code> is just a normal
  typing experience without any modes or shortcut macros and other stuff. Though it's
  nice that PGCLI offers a VI-mode, it is great for people using Vim, when using the
  Shell for editing a Query, this mode can be utilized.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659882586/blog-media/pgcli-vi-mode-demo.gif\"
  alt=\"PGCLI VI-mode\" /></p>\n<p>We can use <code>hjkl</code> keybinding for movement
  and other normal mode keymaps like <code>w</code> for moving a word, or <code>b</code>
  for moving a word in the backward direction, <code>e</code> to the end of the word,
  etc. We can use <code>i</code> or <code>a</code> for entering into insert mode,
  <code>dd</code> to delete the entire line, <code>cc</code> or <code>S</code> to
  delete the current line and enter into insert mode, <code>yy</code> to yank line,
  and <code>p</code> to paste, <code>u</code> for undo, etc.</p>\n<p>We can even execute
  shell commands by using the <code>!</code> in normal mode where we will be prompted
  for the command executed in the same window.</p>\n<h3>Autocompletion in PGCLI shell</h3>\n<p>PGCLI
  shell has an auto-completion feature that will provide some suggestions based on
  the recent queries or commands executed or the default options. You can toggle the
  autocompletion with the <code>F2</code> key, it is officially called as <code>Smart-Completion</code>
  as you would see in the bottom left corner of the shell.</p>\n<h3>Select an Editor
  to write queries</h3>\n<p>We can use different editors for writing SQL queries.
  We can write a query from an editor by using the <code>\\e</code> command. The default
  editor is chosen which will be set in one of your <code>$EDITOR</code> or <code>$VISUAL</code>
  environment variables. You can set them manually in the bashrc file by exporting
  the value of <code>VISUAL</code> or <code>EDITOR</code> as the name of your editor.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">export</span> <span class=\"nv\">VISUAL</span><span class=\"o\">=</span>vim\n\nOR\n\n<span
  class=\"nb\">export</span> <span class=\"nv\">EDITOR</span><span class=\"o\">=</span>vim\n\nOR\n\n<span
  class=\"nb\">export</span> <span class=\"nv\">VISUAL</span><span class=\"o\">=</span>gedit\n</pre></div>\n\n</pre>\n\n<p>Here,
  we can set <code>VISUAL</code> environment variable as the name of the editor or
  the path to it. The <code>VISUAL</code> environment variable is preferred first.
  For the details, we can check out the <a href=\"https://www.pgcli.com/editor\">editor
  command documentation</a>.</p>\n<p>Once, we have configured the editor, we will
  be inside the editor after we enter the <code>\\e</code> command. We can write queries
  inside the editor and quit after we are done, once the editor is closed, the text
  is read from the editor and entered in the psql shell. This will make us wait for
  the enter command for executing the query.</p>\n<p>We can even execute more than
  one query at a time if we use the <code>;</code> at the end of each query. Also,
  we can press <code>v</code> in VI-normal mode, to open the current query in the
  default editor.</p>\n<h3>Save the query output to a file</h3>\n<p>Let's say we have
  written a query inside the psql shell and we want the output to be saved in a file.
  We can achieve that with the <code>\\o</code> command, we need to specify the filename
  in which the output of the queries will be saved.</p>\n<p><strong>NOTE: Once we
  use the <code>\\o</code> command, all the queries will start appending the results
  in the provided file. To stop the behavior, you can use <code>\\o</code> (without
  any filename) to stop appending the results to a file.</strong></p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span># Start recording the output
  of queries to file\n\\o filename\n\n\n# Stop recording the output of queries to
  file\n\\o\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659884242/blog-media/pgcli-output-query.gif\"
  alt=\"PGCLI Query output to file\" /></p>\n<h3>Execute query from a file</h3>\n<p>We
  can even execute queries stored in a file, we can use the <code>\\i</code> command
  followed by the file name. This will load in the queries from the file and will
  display the output of each query.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>\\i file_name\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659883225/blog-media/pgcli-file-query.gif\"
  alt=\"PGCLI query from a file\" /></p>\n<h2>Summary</h2>\n<p>In this little article,
  we were able to explore the basics of the PGCLI package in python, with that package
  we can interact with the Postgres database that can be on your local server or somewhere
  remotely. With the PSQL-like shell provided by the PGCLI tool, we were able to write
  SQL queries, execute commands, write/read queries from files, etc. This tool is
  quite handy for people dealing with postgres databases and who need to view data
  or schema in that database.</p>\n<p>I've personally used this tool in my current
  internship where I needed to create a few SQL queries to fetch particular data related
  to various tables and relationships. It saved a lot of time with auto-completion
  and integration with the text editor for writing longer queries, especially with
  JOINS and nested queries.</p>\n<h2>Conclusion</h2>\n<p>This was my take and experience
  of using the PGCLI tool and exploring the various features it provides that can
  be leveraged for quick interactions with a Postgres database. Hopefully, you have
  learned something from this post, if you have any queries(NOT SQL queries) or feedback,
  please let me know in the comments or on my social handles. Thank you for reading
  and Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/python-pkg-pgcli-postgres-from-terminal.png
long_description: Have you ever used the Postgres database and did you know you don
  If you are like CLIs and love to play with backend systems like Postgres(database)
  then chances are you will love this article. This article will cover the basics
  of using the PGCLI to
now: 2025-05-01 18:11:33.312266
path: blog/posts/2022-08-07-PGCLI-postgres-from-terminal.md
prevnext: null
slug: pgcli-pg-from-terminal
status: published
tags:
- postgres
- python
- DBMS
templateKey: blog-post
title: 'PGCLI: Postgres from the terminal'
today: 2025-05-01
---

## Introduction

Have you ever used the Postgres database and did you know you don't have to launch PGAdmin every time you want to write SQL queries, you can write down those queries even inside a terminal with a python package. It's PGCLI, that can act as an editor window for writing SQL queries for Postgres Databases. We can simply execute SQL queries in a PSQL-like shell with additional features like autocompletion, text-editing modes, file input/output options, etc.

If you are like CLIs and love to play with backend systems like Postgres(database) then chances are you will love this article. This article will cover the basics of using the PGCLI tool to perform simply as well as advanced options in interacting with your Postgres database.

## PGCLI 

PGCLi is a python package that acts as a CLI for executing SQL queries in a Postgres database. PGCLI allows us to interact with any Postgres database via the command line, it can be a remote database or a local database, you can access it right away from your terminal. The package acts as a CLI for accessing and executing any SQL queries, so we can simply provide the options/parameters to plug the database in and access the schema and its related data.

## Installation

There are a lot of ways pgcli can be installed, my preferred way is with `pipx`, it just works well and you don't have to install it every time or mess up your global python packages environment. 

```
pipx install pgcli
```

![PGCLI Install with pipx](https://res.cloudinary.com/techstructive-blog/image/upload/v1659878333/blog-media/pgcli-install-pipx.png)

We don't even require installing it when using pipx, as it will each time create a fresh instance of the pgcli in a virtual environment.

We can now run the `pgcli` with parameters using the `pipx run` command. In the below command, we have used the `--help` options to list out all the commands and other options available with the `pgcli` command.

```
pipx run pgcli --help
```

![PGCLI Help](https://res.cloudinary.com/techstructive-blog/image/upload/v1659878322/blog-media/pgcli-help.png)


If this command is working fine, it means we are ready to connect and interact with a database with pgcli.

## List local databases

We can list all the databases in our local system by providing the hostname and the username, in my case, I have to provide the `localhost` and `postgres` as the hostname and the username respectively. Further, by providing the `--list` option, we can list down all the available databases on our system, basically for a given host and user.

```
pipx run pgcli --host localhost --user postgres --list

OR

pipx run pgcli -h localhost -U postgres -l
```

![PGCLI List all databases](https://res.cloudinary.com/techstructive-blog/image/upload/v1659878596/blog-media/pgcli-list-db.png)

As, we can see, the command lists out all the available databases on the local server. This can help in getting an overview of the Postgres databases present in your local system. You need to enter the password for the username to access the database.

## Connect to a database

To connect to an individual Postgres database we have two options, we can either use a URL string of postgres database or parse parameters to the command individually. It's easier to use the URL option because it wraps a lot of things in a single string rather than writing each parameters.

### Connect with the Postgres Connection URI

We can simply pass the postgres connection URI for accessing a postgres database with pgcli. The URI has a specific pattern in the order where we need to specify the values like hostname, password, username, port, and database name. 

```
postgresql://username:password@hostname:port/database_name
```

The above is the structure for a postgres connection URI, we parse in the `username`, `password` for that `username`, `hostname` with the `port`, and the `database name`. These parameters if parsed correctly can directly access the database.

```
pipx run pgcli postgresql://postgres:postgres@localhost:5432/techstructive_blog
```

![PGCLI Postgres Connection URI](https://res.cloudinary.com/techstructive-blog/image/upload/v1659878955/blog-media/pgcli-connect-uri.png)

This command will leave us in a prompt where we are basically in a PSQL shell in the provided database or host. For connecting to a local database, the hostname will likely be `localhost` and postgres database port is generally `5432` there might be exceptions to this. We can quit out of PGCLI with `CTRL + D` or `exit`.

While using a remote database, the hostname and other details might be provided, for instance, we might have a Django application, in that project, we will have `database_name`, `port`, and details mentioned in the `DATABASES` settings, those fields can be checked and a postgres connection URI might be constructed.

Let's take an example.

```python
DATABASES = {
    "default": {
        "ENGINE": "django.db.backends.postgresql",
        "NAME": "blog",
        "USER": "meet",
        "PASSWORD": "postgres",
        "HOST": "localhost",
    }
}
```

We have all the fields provided in the django project' settings file. We can use this to construct a postgres database connection URI.

```
pgcli postgresql://meet:postgres@localhost:5432/blog
```

So, that is how you can extract parameters from other technologies and frameworks for working with postgres URIs.

### Connect with CLI options/paramters

We can parse the options like `hostname`, `password`, and `database name` manually with the CLI options that PGCLI provides. 

```
pipx run pgcli --username postgres --host localhost --port 5432 --dbname techstructive_blog

OR

pipx run pgcli -u postgres -h localhost -p 5432 -d techstructive_blog
```

![PGCLI Connect Paramaters](https://res.cloudinary.com/techstructive-blog/image/upload/v1659879642/blog-media/pgcli-connect-paramters.png)


You will be prompted for a password for the database, after entering the password, you will be inside the psql shell. In there, we can execute SQL queries and other specific psql or pgcli commands.

Once we have connected to a database, we can now try to work with the SQL queries and explore some of the commands and features provided by pgcli.

## List out all databases if connected to a server

Let's say we don't know the database name, we are just connected to a database server, we can get the list of the databases inside the shell with the `\l+` command. This command is similar to the `--list` parameter in the PGCLI command, but the only difference is that we are executing it from the psql shell rather than our terminal shell.

```
# Let's say we connected to a database server

pipx run pgcli --username postgres --host localhost --port 5432


# List out all databases in that server

\l+
```

![PGCLI List databases in a server](https://res.cloudinary.com/techstructive-blog/image/upload/v1659879988/blog-media/pgcli-list-databases.png)


So, that's how we would get the gist of all the databases inside the database server to which we are connected using the `\l+` command.


## List all tables in the database

We can get a list of all the tables existing in the database we are logged in with the `\dt` command, we can simply enter the command in the prompt we are directed to once we are logged in to the postgres database.

```
\dt
```

![PGCLI List tables in a database](https://res.cloudinary.com/techstructive-blog/image/upload/v1659880232/blog-media/pgcli-list-tables.png)

We can use `\d <table_name>` to describe the details of the table provided. We get back the list of the attributes inside the table and the indexes of the relationships of the table as well as the referenced table details.

![PGCLI describe a table in a database](https://res.cloudinary.com/techstructive-blog/image/upload/v1659880453/blog-media/pgcli-describe-table.png)

We also have other commands for displaying types, schemas, roles, views, etc. with the `\d` prefix, a detailed list can be found in the [PGCLI documentation](https://www.pgcli.com/commands).

## Executing queries

We can execute SQL queries in the PGCLI prompt as normally we do in a PSQL shell. There are added benefits like certain modes of editing text and auto-completion. We'll execute a few queries and also try out other options for editing and saving queries.

Normally for simple and easy tasks, we can write SQL queries right in the PSQL shell, it might be small and simple to construct a query. Let's take the simplest example to get all the attribute lists from a specific table in a database.

```SQL
SELECT * FROM table_name
```

![PGCLI executing SQL queries](https://res.cloudinary.com/techstructive-blog/image/upload/v1659880794/blog-media/pgcli-sql-query.png)

We get an excellent table display of output which we can view by scroll or arrow keys. We can even use `j` and `k` for moving vertically in an output table.

### Wriing modes in PGCLI

We can use two modes in PGCLI prompt, one is the Emacs mode which is the default one and the other is Vi-mode for Vim fanboys. We can switch between these modes with `F4` key. The `Emacs-mode` is just a normal typing experience without any modes or shortcut macros and other stuff. Though it's nice that PGCLI offers a VI-mode, it is great for people using Vim, when using the Shell for editing a Query, this mode can be utilized.

![PGCLI VI-mode](https://res.cloudinary.com/techstructive-blog/image/upload/v1659882586/blog-media/pgcli-vi-mode-demo.gif)

We can use `hjkl` keybinding for movement and other normal mode keymaps like `w` for moving a word, or `b` for moving a word in the backward direction, `e` to the end of the word, etc. We can use `i` or `a` for entering into insert mode, `dd` to delete the entire line, `cc` or `S` to delete the current line and enter into insert mode, `yy` to yank line, and `p` to paste, `u` for undo, etc.

We can even execute shell commands by using the `!` in normal mode where we will be prompted for the command executed in the same window.

### Autocompletion in PGCLI shell

PGCLI shell has an auto-completion feature that will provide some suggestions based on the recent queries or commands executed or the default options. You can toggle the autocompletion with the `F2` key, it is officially called as `Smart-Completion` as you would see in the bottom left corner of the shell.

### Select an Editor to write queries

We can use different editors for writing SQL queries. We can write a query from an editor by using the `\e` command. The default editor is chosen which will be set in one of your `$EDITOR` or `$VISUAL` environment variables. You can set them manually in the bashrc file by exporting the value of `VISUAL` or `EDITOR` as the name of your editor.

```bash
export VISUAL=vim

OR

export EDITOR=vim

OR

export VISUAL=gedit
```

Here, we can set `VISUAL` environment variable as the name of the editor or the path to it. The `VISUAL` environment variable is preferred first. For the details, we can check out the [editor command documentation](https://www.pgcli.com/editor).

Once, we have configured the editor, we will be inside the editor after we enter the `\e` command. We can write queries inside the editor and quit after we are done, once the editor is closed, the text is read from the editor and entered in the psql shell. This will make us wait for the enter command for executing the query. 

We can even execute more than one query at a time if we use the `;` at the end of each query. Also, we can press `v` in VI-normal mode, to open the current query in the default editor.

### Save the query output to a file

Let's say we have written a query inside the psql shell and we want the output to be saved in a file. We can achieve that with the `\o` command, we need to specify the filename in which the output of the queries will be saved. 

**NOTE: Once we use the `\o` command, all the queries will start appending the results in the provided file. To stop the behavior, you can use `\o` (without any filename) to stop appending the results to a file.**

```
# Start recording the output of queries to file
\o filename


# Stop recording the output of queries to file
\o
```

![PGCLI Query output to file](https://res.cloudinary.com/techstructive-blog/image/upload/v1659884242/blog-media/pgcli-output-query.gif)

### Execute query from a file

We can even execute queries stored in a file, we can use the `\i` command followed by the file name. This will load in the queries from the file and will display the output of each query.

```
\i file_name
```

![PGCLI query from a file](https://res.cloudinary.com/techstructive-blog/image/upload/v1659883225/blog-media/pgcli-file-query.gif)

## Summary

In this little article, we were able to explore the basics of the PGCLI package in python, with that package we can interact with the Postgres database that can be on your local server or somewhere remotely. With the PSQL-like shell provided by the PGCLI tool, we were able to write SQL queries, execute commands, write/read queries from files, etc. This tool is quite handy for people dealing with postgres databases and who need to view data or schema in that database.

I've personally used this tool in my current internship where I needed to create a few SQL queries to fetch particular data related to various tables and relationships. It saved a lot of time with auto-completion and integration with the text editor for writing longer queries, especially with JOINS and nested queries.

## Conclusion

This was my take and experience of using the PGCLI tool and exploring the various features it provides that can be leveraged for quick interactions with a Postgres database. Hopefully, you have learned something from this post, if you have any queries(NOT SQL queries) or feedback, please let me know in the comments or on my social handles. Thank you for reading and Happy Coding :)