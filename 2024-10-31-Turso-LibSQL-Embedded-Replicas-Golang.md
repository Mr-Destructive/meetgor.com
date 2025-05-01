---
article_html: "<h2>Introduction</h2>\n<p>Welcome to the Let's Go with Turso series.
  In this series, we will learn how to interact with LibSQL databases with Golang.
  In the past article of the series, we explored how to connect remote/local LibSQL
  database in golang.</p>\n<p>With this section, we will specifally dive into understanding
  how to create, connect, and query local embedded replicas of LibSQL database hosted
  on Turso with a Golang application.</p>\n<p>If you want to check out the YouTube
  video, check this out:</p>\n<p><a href=\"https://www.youtube.com/watch?v=BitxB40rdVw\">LibSQL
  Embedded Replicas on Turso in Golang</a></p>\n<iframe width=\"560\" height=\"315\"
  src=\"https://www.youtube.com/embed/vBrvX0X0phw\" frameborder=\"0\" allowfullscreen></iframe>\n<h2>LibSQL
  Embedded Replicas on Turso in Golang</h2>\n<h3>Embedded Replicas</h3>\n<p>The embedded
  replica is a native feature of the libSQL database. With embedded replicas, we can
  have faster writes as with the local database and global read access from the remote
  database.</p>\n<p>Embedded replica is like creating a copy of a primary remote database
  and using it for performing any operations from the applications as a local database
  and then it has a mechanism or standard to sync up with the primary remote database.
  The primary remote database serves as a single source of truth, however we can use
  the database locally as well. This however should be done carefully to avoid data
  corruption or stale data. To cope up with this stale or corruption of data, the
  periodic syncing can be used.</p>\n<p>Embedded replicas have a 3 fold process</p>\n<ul>\n<li>Create
  a copy of the primary remote database</li>\n<li>Perform any operations on the local
  database</li>\n<li>Sync up with the primary remote database</li>\n</ul>\n<p><img
  src=\"https://meetgor-cdn.pages.dev/embedded-replicas-flow.png\" alt=\"Embedded
  Replicas for LibSQL\" /></p>\n<p>There are a few things to remember here:</p>\n<ul>\n<li>The
  local database can read it's newly written data, however other local replica databases
  can only view those changes once the sync has happened and the primary database
  has been updated from the local copy.</li>\n<li>These would require to sync the
  local with the primary database first and then the other local database also needs
  to sync with the primary database.</li>\n</ul>\n<p>You can read more about it <a
  href=\"https://docs.turso.tech/features/embedded-replicas/introduction\">here</a>
  on the Turso documentation.</p>\n<p>Let's get started,</p>\n<p>What we are going
  to do is create a step by step procedure to understand how to work with embedded
  replicas of LibSQL database.</p>\n<ol>\n<li>Perform the operations on the local
  LibSQL database</li>\n<li>Create a Embedded Replica and sync up with the primary
  remote database</li>\n<li>Fetch data from the primary remote database</li>\n</ol>\n<p><img
  src=\"https://meetgor-cdn.pages.dev/LibSQL_Embedded_Replicas_on_Turso_in_Golang.gif\"
  alt=\"Embedded Replicas of LibSQL with Golang\" /></p>\n<p>In this way, we can understand
  how to operate the Embedded Replicas as a whole from locally as well as remotely</p>\n<h2>Initializing
  a Golang project</h2>\n<p>Let's start with initializing a Golang project.</p>\n<pre
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
  class=\"c1\"># go mod init &lt;git-provider-domain&gt;/&lt;username&gt;/&lt;project-name&gt;</span>\n<span
  class=\"c1\"># Example</span>\n\ngo mod init github.com/mr-destructive/lets-go-with-turso\n</pre></div>\n\n</pre>\n\n<p>This
  will initialize the project in the current directory, creating a <code>go.mod</code>
  file with the specification of the Golang version and the packages that we will
  install and use in this module.</p>\n<h2>Installing go-libsql package</h2>\n<p>We
  will need to install the <a href=\"https://github.com/tursodatabase/go-libsql\">go-libsql</a>
  package, this is the driver for LibSQL for Golang. To install simply use the <code>go
  get</code> command to be installed as the dependency for the project</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>go
  get github.com/tursodatabase/go-libsql\n</pre></div>\n\n</pre>\n\n<p>Once this is
  done, we can create a local LibSQL database.</p>\n<h2>Creating a local LibSQL database</h2>\n<p>Let's
  create a local LibSQL database. With the <code>go-libsql</code> package, we can
  connect to the local database. This will be done using the <code>libsql</code> driver.
  There is really no much difference than the normal SQLite database driver, this
  works just like SQLite, the only difference being the ability to connect to the
  remote database.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;database/sql&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">_</span><span class=\"w\"> </span><span class=\"s\">&quot;github.com/tursodatabase/go-libsql&quot;</span><span
  class=\"w\"></span>\n<span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">dbName</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;file:./local.db&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">db</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">sql</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;libsql&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">dbName</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to open db %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">db</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The above code will simply connect
  to the local LibSQL database which resides as the <code>local.db</code> file. Now,
  to demonstrate that it is an actual sqlite-like database, we can execute queries
  on the connected database.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;database/sql&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">_</span><span class=\"w\"> </span><span class=\"s\">&quot;github.com/tursodatabase/go-libsql&quot;</span><span
  class=\"w\"></span>\n<span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">dbName</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;file:./local.db&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">db</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">sql</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;libsql&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">dbName</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to open db %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">db</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">rows</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">db</span><span class=\"p\">.</span><span
  class=\"nx\">Query</span><span class=\"p\">(</span><span class=\"s\">&quot;SELECT
  ABS(RANDOM()%5) FROM generate_series(1,10)&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to query %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">rows</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">rows</span><span
  class=\"p\">.</span><span class=\"nx\">Next</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">i</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">rows</span><span class=\"p\">.</span><span
  class=\"nx\">Scan</span><span class=\"p\">(</span><span class=\"o\">&amp;</span><span
  class=\"nx\">i</span><span class=\"p\">);</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to scan %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">i</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Output:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run main.go\n\n3\n0\n4\n4\n2\n1\n2\n3\n4\n1\n\n$
  go run main.go\n\n0\n2\n1\n2\n3\n2\n2\n1\n0\n2\n</pre></div>\n\n</pre>\n\n<p>Just
  a simple <code>SELECT ABS(RANDOM()%5) FROM generate_series(1,10)</code> query will
  be executed on the connected database. This query will basically genrate a random
  number between <code>-4</code> and <code>4</code> and take the absolute value i.e.
  now between 0 and 4, this will genrate 10 such numbers.</p>\n<p>Now, we can move
  into the actual embedded replica creation of the primary remote database and use
  it as a local database to perform operations.</p>\n<h2>Creating an Embedded Replica
  of Turso's LibSQL database</h2>\n<p>We need a few things to specify before we create
  the embedded replica, first being the primary remote database, that typically is
  a libsql database hosted on turso or self hosted. We also need to provide the local
  database path, where the local-replica will be stored or we can say the copy of
  the primary libsql database will be created.</p>\n<p>We will be using the <a href=\"https://pkg.go.dev/github.com/levydsa/libsql-go#NewEmbeddedReplicaConnector\">LibSQL.NewEmbeddedReplicaConnector</a>
  to create a local replica of a libsql database. The function takes in two paramters,
  the first paramter being the local database filename path to create the copy into,
  and the second paramter being the primary database URL. The function returns a connector
  object or an error if any. The connector object is then further used with <a href=\"https://pkg.go.dev/database/sql#OpenDB\">OpenDB</a>
  function to create a database connection. The <code>OpenDB</code> function returns
  a reference of database connections which we'll use to connect and perform operations
  on the database.\nThe same <code>connector</code> object could be used to sync with
  the primary database after performing operations on the local database with the
  <a href=\"https://pkg.go.dev/github.com/levydsa/libsql-go#Connector.Sync\">Sync</a>
  method. This will pull or push the changes from the local database to the primary
  database.</p>\n<p>We can configure the syncing mechanism while creating the embedded
  replica with the additional parameters to the <code>NewEmbeddedReplicaConnector</code>
  function. There are <a href=\"https://pkg.go.dev/github.com/levydsa/libsql-go#Option\">Options</a>
  to include for the paramters that could be passed like:</p>\n<ul>\n<li><code>WithAuthToken(string)</code>:
  This will be used to authenticate with the primary database.</li>\n<li><code>WithSyncInterval(time.Duration)</code>:
  This will be used to specify the interval of syncing between the local and primary
  database.</li>\n<li><code>WithEncrytion(string)</code>: This will be used to encrypt
  the local database.</li>\n<li><code>WithReadYourWrites(bool)</code>: This will be
  used to specify if the local database can read the newly written changes or not.</li>\n</ul>\n<p>So,
  let's create a exmaple to create a embedded replica, make some changes by creating
  tables and then syncing the local with primary, finally appending some data to the
  local and reading those.</p>\n<h4>Create the Embedded Replica</h4>\n<p>We first
  need to create a copy of the primary database, as said, we will configure the 2
  paramters that we need to create the embedded replica with <code>NewEmbeddedReplicaConnector</code>.
  Then once we have the connector object, we open up a database connection, at that
  point we can further run queries on the local replica that was just created from
  the primary remote LibSQL database.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;database/sql&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;path/filepath&quot;</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;github.com/tursodatabase/go-libsql&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">dbName</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;local.db&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// this is
  not required, but can be used to create a temporary directory and then delete it
  later</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">dir</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">MkdirTemp</span><span class=\"p\">(</span><span
  class=\"s\">&quot;&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;libsql-*&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Error creating temporary directory:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">os</span><span class=\"p\">.</span><span class=\"nx\">Exit</span><span
  class=\"p\">(</span><span class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">os</span><span class=\"p\">.</span><span class=\"nx\">RemoveAll</span><span
  class=\"p\">(</span><span class=\"nx\">dir</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"c1\">// first
  paramter required for creating NewEmbeddedReplicaConnector</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">dbPath</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">filepath</span><span
  class=\"p\">.</span><span class=\"nx\">Join</span><span class=\"p\">(</span><span
  class=\"nx\">dir</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">dbName</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">dbPath</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">    </span><span
  class=\"c1\">// second paramter required for creating NewEmbeddedReplicaConnector</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">dbURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Getenv</span><span class=\"p\">(</span><span
  class=\"s\">&quot;TURSO_DATABASE_URL&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">dbAuthToken</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Getenv</span><span class=\"p\">(</span><span
  class=\"s\">&quot;TURSO_AUTH_TOKEN&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">connector</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">libsql</span><span
  class=\"p\">.</span><span class=\"nx\">NewEmbeddedReplicaConnector</span><span class=\"p\">(</span><span
  class=\"nx\">dbPath</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">dbURL</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">libsql</span><span class=\"p\">.</span><span class=\"nx\">WithAuthToken</span><span
  class=\"p\">(</span><span class=\"nx\">dbAuthToken</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">connector</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to open db %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">connector</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">    </span><span
  class=\"c1\">// open a database connection from the connector object</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">db</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">sql</span><span
  class=\"p\">.</span><span class=\"nx\">OpenDB</span><span class=\"p\">(</span><span
  class=\"nx\">connector</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Connected
  to database&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to open db %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">db</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In the above code, we first create
  a temporary directory with the help of <a href=\"https://pkg.go.dev/os#MkdirTemp\">MkdirTemp</a>,
  this is not required, but would be easier for cleanup later. We then the path for
  the local database to be created. The combined path string <code>dbPath</code> will
  serve as the first paramter to the <code>NewEmbeddedReplicaConnector</code>. Then
  we have taken the <code>dbURL</code> and the <code>dbAuthToken</code> from the environment
  variables <code>TURSO_DATABASE_URL</code> and <code>TURSO_AUTH_TOKEN</code> respectively.
  The <code>dbURL</code> will serve as the second paramter for the <code>NewEmbeddedReplicaConnector</code>
  that is the URL of the primary remote LibSQL database. The function <code>NewEmbeddedReplicaConnector</code>
  will return the <code>Connector</code> object if successfull in creation of the
  replica, else return <code>err</code> if it fails. The connector object needs to
  be closed at the end of the program, so we use the <code>defer connector.Close()</code>
  that will close the connection to the primary database at the end of the program.
  The <code>sql.OpenDB</code> is used to create the connection with the local database
  that will be created from the <code>connector</code> object. Finally we also need
  to close the local database connection at the end of the program.</p>\n<p>Further,
  we will try to query the local replica and create tables and append data to it.</p>\n<h3>Adding
  data to teh local replica</h3>\n<p>Once we have the <code>db</code> connection to
  the local database, we can noramlly query the database as we did in the previous
  example, of querying the local LibSQL database. Let's start by creating a table
  <code>posts</code> to the local replica, this will basically create the schema in
  the local database.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"w\">    </span><span class=\"o\">...</span><span class=\"p\">.</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">createPostTableQuery</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`CREATE
  TABLE IF NOT EXISTS posts(</span>\n<span class=\"s\">        id INTEGER PRIMARY
  KEY,</span>\n<span class=\"s\">        title VARCHAR(100),</span>\n<span class=\"s\">
  \       description VARCHAR(255),</span>\n<span class=\"s\">        content TEXT</span>\n<span
  class=\"s\">    );`</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">db</span><span
  class=\"p\">.</span><span class=\"nx\">Exec</span><span class=\"p\">(</span><span
  class=\"nx\">createPostTableQuery</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to create table %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  <code>createPostTableQuery</code> will be the <code>SQL</code> to create the table
  <code>posts</code> if it doesn't already exist in the database (local replica).
  Then with the help of <a href=\"https://pkg.go.dev/database/sql#DB.Exec\">db.Exec</a>
  function, we can execute the query and return back the rows if it had created any.
  In this case it won't as we have just added a table.</p>\n<p>Then, we can either
  sync the database to the primary, but let's populate the table <code>posts</code>
  with some data before syncing with the primary db.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"w\">\t</span><span class=\"nx\">createPostQuery</span><span class=\"w\">
  </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`INSERT
  INTO posts(title, description, content) </span>\n<span class=\"s\">        VALUES(?,
  ?, ?)`</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">_</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">db</span><span
  class=\"p\">.</span><span class=\"nx\">Exec</span><span class=\"p\">(</span><span
  class=\"nx\">createPostQuery</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;test title&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;test description&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;test content&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to insert %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>We
  have created the <code>createPostQuery</code> similarly to insert into the <code>posts</code>
  table in the local replica. The values are added with the placeholders in the <code>Exec</code>
  function as positional parameters. Once we have executed the query, this will populate
  the <code>posts</code> table in the lcoal replica.</p>\n<p>We can now finally sync
  with the primary remote LibSQL database to make sure that the primary database also
  has these migrations applied.</p>\n<h3>Syncing the local replica</h3>\n<p>Remember,
  <code>connector</code> is for primary database and <code>db</code> is for the local
  replica. So, we will sync the remote database from the replica that was created
  with the <code>connector.Sync</code></p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"w\">\t</span><span class=\"nx\">_</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">connector</span><span
  class=\"p\">.</span><span class=\"nx\">Sync</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to sync %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Printf</span><span class=\"p\">(</span><span class=\"s\">&quot;Successfully
  synced %s db\\n&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">dbPath</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">rows</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">db</span><span
  class=\"p\">.</span><span class=\"nx\">Query</span><span class=\"p\">(</span><span
  class=\"s\">&quot;SELECT * FROM posts&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to query %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">rows</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">rows</span><span
  class=\"p\">.</span><span class=\"nx\">Next</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">id</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">title</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">description</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">content</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">rows</span><span class=\"p\">.</span><span
  class=\"nx\">Scan</span><span class=\"p\">(</span><span class=\"o\">&amp;</span><span
  class=\"nx\">id</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">title</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">description</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">content</span><span class=\"p\">);</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to scan %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">id</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">title</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">description</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">content</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Output:</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go                                                            \n\n/tmp/libsql-349052144/local.db\n<span
  class=\"p\">&amp;</span><span class=\"o\">{</span>0x2eec9d0 &lt;nil&gt; &lt;nil&gt;<span
  class=\"o\">}</span>\nConnected to database\nSuccessfully synced /tmp/libsql-349052144/local.db
  db\n<span class=\"m\">1</span> <span class=\"nb\">test</span> title <span class=\"nb\">test</span>
  description <span class=\"nb\">test</span> content\n</pre></div>\n\n</pre>\n\n<p>Once
  we have synced the local replica, we can now query the database i.e. the local replica,
  with the changes, also note that this could also be done without syncing the database,
  but the primary database won't have the applied changes.</p>\n<p>We finally Query
  the local replica with the query <code>SELECT * FROM posts</code> and print out
  the results. This has the 1 record in the <code>posts</code> table that we inserted.</p>\n<p>So,
  that's how we basically create a local replica from a remote LibSQL database hosted
  on Turso. We first create the path for the local database to be copied, then provide
  the primary database URL and credentials, then request a copy of the primary database,
  then we perform any mutation or operations on the local copy and finally sync up
  with the remote primary database to persist the data from that replica (acting like
  a session of database operation).</p>\n<p>That wraps the article for now.</p>\n<p>For
  all the code related to this article, you can check out the <a href=\"https://github.com/mr-destructive/lets-go-with-turso\">Let's
  Go with Turso</a> GitHub repo for all the examples and additional examples for using
  LibSQL with Golang.</p>\n<h2>Conclusion</h2>\n<p>So, that is a wrap for this part
  of the series, we have explored how to create a local embedded replica from a remote
  LibSQL database hosted on Turso with Golang. In the next part of the series, we
  will explore how to setup a local LibSQL database server and then connect it with
  Golang.</p>\n<p>Thank you for reading this post, If you have any questions, feedback,
  and suggestions, feel free to drop them in the comments.</p>\n<p>Happy Coding :)</p>\n"
cover: ''
date: 2024-10-31
datetime: 2024-10-31 00:00:00+00:00
description: Understanding and exploring how to create, connect, and query local embedded
  replicas of LibSQL database hosted on Turso with a Golang application.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2024-10-31-Turso-LibSQL-Embedded-Replicas-Golang.md
html: "<h2>Introduction</h2>\n<p>Welcome to the Let's Go with Turso series. In this
  series, we will learn how to interact with LibSQL databases with Golang. In the
  past article of the series, we explored how to connect remote/local LibSQL database
  in golang.</p>\n<p>With this section, we will specifally dive into understanding
  how to create, connect, and query local embedded replicas of LibSQL database hosted
  on Turso with a Golang application.</p>\n<p>If you want to check out the YouTube
  video, check this out:</p>\n<p><a href=\"https://www.youtube.com/watch?v=BitxB40rdVw\">LibSQL
  Embedded Replicas on Turso in Golang</a></p>\n<iframe width=\"560\" height=\"315\"
  src=\"https://www.youtube.com/embed/vBrvX0X0phw\" frameborder=\"0\" allowfullscreen></iframe>\n<h2>LibSQL
  Embedded Replicas on Turso in Golang</h2>\n<h3>Embedded Replicas</h3>\n<p>The embedded
  replica is a native feature of the libSQL database. With embedded replicas, we can
  have faster writes as with the local database and global read access from the remote
  database.</p>\n<p>Embedded replica is like creating a copy of a primary remote database
  and using it for performing any operations from the applications as a local database
  and then it has a mechanism or standard to sync up with the primary remote database.
  The primary remote database serves as a single source of truth, however we can use
  the database locally as well. This however should be done carefully to avoid data
  corruption or stale data. To cope up with this stale or corruption of data, the
  periodic syncing can be used.</p>\n<p>Embedded replicas have a 3 fold process</p>\n<ul>\n<li>Create
  a copy of the primary remote database</li>\n<li>Perform any operations on the local
  database</li>\n<li>Sync up with the primary remote database</li>\n</ul>\n<p><img
  src=\"https://meetgor-cdn.pages.dev/embedded-replicas-flow.png\" alt=\"Embedded
  Replicas for LibSQL\" /></p>\n<p>There are a few things to remember here:</p>\n<ul>\n<li>The
  local database can read it's newly written data, however other local replica databases
  can only view those changes once the sync has happened and the primary database
  has been updated from the local copy.</li>\n<li>These would require to sync the
  local with the primary database first and then the other local database also needs
  to sync with the primary database.</li>\n</ul>\n<p>You can read more about it <a
  href=\"https://docs.turso.tech/features/embedded-replicas/introduction\">here</a>
  on the Turso documentation.</p>\n<p>Let's get started,</p>\n<p>What we are going
  to do is create a step by step procedure to understand how to work with embedded
  replicas of LibSQL database.</p>\n<ol>\n<li>Perform the operations on the local
  LibSQL database</li>\n<li>Create a Embedded Replica and sync up with the primary
  remote database</li>\n<li>Fetch data from the primary remote database</li>\n</ol>\n<p><img
  src=\"https://meetgor-cdn.pages.dev/LibSQL_Embedded_Replicas_on_Turso_in_Golang.gif\"
  alt=\"Embedded Replicas of LibSQL with Golang\" /></p>\n<p>In this way, we can understand
  how to operate the Embedded Replicas as a whole from locally as well as remotely</p>\n<h2>Initializing
  a Golang project</h2>\n<p>Let's start with initializing a Golang project.</p>\n<pre
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
  class=\"c1\"># go mod init &lt;git-provider-domain&gt;/&lt;username&gt;/&lt;project-name&gt;</span>\n<span
  class=\"c1\"># Example</span>\n\ngo mod init github.com/mr-destructive/lets-go-with-turso\n</pre></div>\n\n</pre>\n\n<p>This
  will initialize the project in the current directory, creating a <code>go.mod</code>
  file with the specification of the Golang version and the packages that we will
  install and use in this module.</p>\n<h2>Installing go-libsql package</h2>\n<p>We
  will need to install the <a href=\"https://github.com/tursodatabase/go-libsql\">go-libsql</a>
  package, this is the driver for LibSQL for Golang. To install simply use the <code>go
  get</code> command to be installed as the dependency for the project</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>go
  get github.com/tursodatabase/go-libsql\n</pre></div>\n\n</pre>\n\n<p>Once this is
  done, we can create a local LibSQL database.</p>\n<h2>Creating a local LibSQL database</h2>\n<p>Let's
  create a local LibSQL database. With the <code>go-libsql</code> package, we can
  connect to the local database. This will be done using the <code>libsql</code> driver.
  There is really no much difference than the normal SQLite database driver, this
  works just like SQLite, the only difference being the ability to connect to the
  remote database.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;database/sql&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">_</span><span class=\"w\"> </span><span class=\"s\">&quot;github.com/tursodatabase/go-libsql&quot;</span><span
  class=\"w\"></span>\n<span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">dbName</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;file:./local.db&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">db</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">sql</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;libsql&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">dbName</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to open db %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">db</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The above code will simply connect
  to the local LibSQL database which resides as the <code>local.db</code> file. Now,
  to demonstrate that it is an actual sqlite-like database, we can execute queries
  on the connected database.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;database/sql&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">_</span><span class=\"w\"> </span><span class=\"s\">&quot;github.com/tursodatabase/go-libsql&quot;</span><span
  class=\"w\"></span>\n<span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">dbName</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;file:./local.db&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">db</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">sql</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;libsql&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">dbName</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to open db %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">db</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">rows</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">db</span><span class=\"p\">.</span><span
  class=\"nx\">Query</span><span class=\"p\">(</span><span class=\"s\">&quot;SELECT
  ABS(RANDOM()%5) FROM generate_series(1,10)&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to query %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">rows</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">rows</span><span
  class=\"p\">.</span><span class=\"nx\">Next</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">i</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">rows</span><span class=\"p\">.</span><span
  class=\"nx\">Scan</span><span class=\"p\">(</span><span class=\"o\">&amp;</span><span
  class=\"nx\">i</span><span class=\"p\">);</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to scan %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">i</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Output:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run main.go\n\n3\n0\n4\n4\n2\n1\n2\n3\n4\n1\n\n$
  go run main.go\n\n0\n2\n1\n2\n3\n2\n2\n1\n0\n2\n</pre></div>\n\n</pre>\n\n<p>Just
  a simple <code>SELECT ABS(RANDOM()%5) FROM generate_series(1,10)</code> query will
  be executed on the connected database. This query will basically genrate a random
  number between <code>-4</code> and <code>4</code> and take the absolute value i.e.
  now between 0 and 4, this will genrate 10 such numbers.</p>\n<p>Now, we can move
  into the actual embedded replica creation of the primary remote database and use
  it as a local database to perform operations.</p>\n<h2>Creating an Embedded Replica
  of Turso's LibSQL database</h2>\n<p>We need a few things to specify before we create
  the embedded replica, first being the primary remote database, that typically is
  a libsql database hosted on turso or self hosted. We also need to provide the local
  database path, where the local-replica will be stored or we can say the copy of
  the primary libsql database will be created.</p>\n<p>We will be using the <a href=\"https://pkg.go.dev/github.com/levydsa/libsql-go#NewEmbeddedReplicaConnector\">LibSQL.NewEmbeddedReplicaConnector</a>
  to create a local replica of a libsql database. The function takes in two paramters,
  the first paramter being the local database filename path to create the copy into,
  and the second paramter being the primary database URL. The function returns a connector
  object or an error if any. The connector object is then further used with <a href=\"https://pkg.go.dev/database/sql#OpenDB\">OpenDB</a>
  function to create a database connection. The <code>OpenDB</code> function returns
  a reference of database connections which we'll use to connect and perform operations
  on the database.\nThe same <code>connector</code> object could be used to sync with
  the primary database after performing operations on the local database with the
  <a href=\"https://pkg.go.dev/github.com/levydsa/libsql-go#Connector.Sync\">Sync</a>
  method. This will pull or push the changes from the local database to the primary
  database.</p>\n<p>We can configure the syncing mechanism while creating the embedded
  replica with the additional parameters to the <code>NewEmbeddedReplicaConnector</code>
  function. There are <a href=\"https://pkg.go.dev/github.com/levydsa/libsql-go#Option\">Options</a>
  to include for the paramters that could be passed like:</p>\n<ul>\n<li><code>WithAuthToken(string)</code>:
  This will be used to authenticate with the primary database.</li>\n<li><code>WithSyncInterval(time.Duration)</code>:
  This will be used to specify the interval of syncing between the local and primary
  database.</li>\n<li><code>WithEncrytion(string)</code>: This will be used to encrypt
  the local database.</li>\n<li><code>WithReadYourWrites(bool)</code>: This will be
  used to specify if the local database can read the newly written changes or not.</li>\n</ul>\n<p>So,
  let's create a exmaple to create a embedded replica, make some changes by creating
  tables and then syncing the local with primary, finally appending some data to the
  local and reading those.</p>\n<h4>Create the Embedded Replica</h4>\n<p>We first
  need to create a copy of the primary database, as said, we will configure the 2
  paramters that we need to create the embedded replica with <code>NewEmbeddedReplicaConnector</code>.
  Then once we have the connector object, we open up a database connection, at that
  point we can further run queries on the local replica that was just created from
  the primary remote LibSQL database.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;database/sql&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;path/filepath&quot;</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;github.com/tursodatabase/go-libsql&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">dbName</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;local.db&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// this is
  not required, but can be used to create a temporary directory and then delete it
  later</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">dir</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">MkdirTemp</span><span class=\"p\">(</span><span
  class=\"s\">&quot;&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;libsql-*&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Error creating temporary directory:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">os</span><span class=\"p\">.</span><span class=\"nx\">Exit</span><span
  class=\"p\">(</span><span class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">os</span><span class=\"p\">.</span><span class=\"nx\">RemoveAll</span><span
  class=\"p\">(</span><span class=\"nx\">dir</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"c1\">// first
  paramter required for creating NewEmbeddedReplicaConnector</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">dbPath</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">filepath</span><span
  class=\"p\">.</span><span class=\"nx\">Join</span><span class=\"p\">(</span><span
  class=\"nx\">dir</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">dbName</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">dbPath</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">    </span><span
  class=\"c1\">// second paramter required for creating NewEmbeddedReplicaConnector</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">dbURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Getenv</span><span class=\"p\">(</span><span
  class=\"s\">&quot;TURSO_DATABASE_URL&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">dbAuthToken</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Getenv</span><span class=\"p\">(</span><span
  class=\"s\">&quot;TURSO_AUTH_TOKEN&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">connector</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">libsql</span><span
  class=\"p\">.</span><span class=\"nx\">NewEmbeddedReplicaConnector</span><span class=\"p\">(</span><span
  class=\"nx\">dbPath</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">dbURL</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">libsql</span><span class=\"p\">.</span><span class=\"nx\">WithAuthToken</span><span
  class=\"p\">(</span><span class=\"nx\">dbAuthToken</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">connector</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to open db %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">connector</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">    </span><span
  class=\"c1\">// open a database connection from the connector object</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">db</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">sql</span><span
  class=\"p\">.</span><span class=\"nx\">OpenDB</span><span class=\"p\">(</span><span
  class=\"nx\">connector</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Connected
  to database&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to open db %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">db</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In the above code, we first create
  a temporary directory with the help of <a href=\"https://pkg.go.dev/os#MkdirTemp\">MkdirTemp</a>,
  this is not required, but would be easier for cleanup later. We then the path for
  the local database to be created. The combined path string <code>dbPath</code> will
  serve as the first paramter to the <code>NewEmbeddedReplicaConnector</code>. Then
  we have taken the <code>dbURL</code> and the <code>dbAuthToken</code> from the environment
  variables <code>TURSO_DATABASE_URL</code> and <code>TURSO_AUTH_TOKEN</code> respectively.
  The <code>dbURL</code> will serve as the second paramter for the <code>NewEmbeddedReplicaConnector</code>
  that is the URL of the primary remote LibSQL database. The function <code>NewEmbeddedReplicaConnector</code>
  will return the <code>Connector</code> object if successfull in creation of the
  replica, else return <code>err</code> if it fails. The connector object needs to
  be closed at the end of the program, so we use the <code>defer connector.Close()</code>
  that will close the connection to the primary database at the end of the program.
  The <code>sql.OpenDB</code> is used to create the connection with the local database
  that will be created from the <code>connector</code> object. Finally we also need
  to close the local database connection at the end of the program.</p>\n<p>Further,
  we will try to query the local replica and create tables and append data to it.</p>\n<h3>Adding
  data to teh local replica</h3>\n<p>Once we have the <code>db</code> connection to
  the local database, we can noramlly query the database as we did in the previous
  example, of querying the local LibSQL database. Let's start by creating a table
  <code>posts</code> to the local replica, this will basically create the schema in
  the local database.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"w\">    </span><span class=\"o\">...</span><span class=\"p\">.</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">createPostTableQuery</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`CREATE
  TABLE IF NOT EXISTS posts(</span>\n<span class=\"s\">        id INTEGER PRIMARY
  KEY,</span>\n<span class=\"s\">        title VARCHAR(100),</span>\n<span class=\"s\">
  \       description VARCHAR(255),</span>\n<span class=\"s\">        content TEXT</span>\n<span
  class=\"s\">    );`</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">db</span><span
  class=\"p\">.</span><span class=\"nx\">Exec</span><span class=\"p\">(</span><span
  class=\"nx\">createPostTableQuery</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to create table %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  <code>createPostTableQuery</code> will be the <code>SQL</code> to create the table
  <code>posts</code> if it doesn't already exist in the database (local replica).
  Then with the help of <a href=\"https://pkg.go.dev/database/sql#DB.Exec\">db.Exec</a>
  function, we can execute the query and return back the rows if it had created any.
  In this case it won't as we have just added a table.</p>\n<p>Then, we can either
  sync the database to the primary, but let's populate the table <code>posts</code>
  with some data before syncing with the primary db.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"w\">\t</span><span class=\"nx\">createPostQuery</span><span class=\"w\">
  </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`INSERT
  INTO posts(title, description, content) </span>\n<span class=\"s\">        VALUES(?,
  ?, ?)`</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">_</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">db</span><span
  class=\"p\">.</span><span class=\"nx\">Exec</span><span class=\"p\">(</span><span
  class=\"nx\">createPostQuery</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;test title&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;test description&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;test content&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to insert %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>We
  have created the <code>createPostQuery</code> similarly to insert into the <code>posts</code>
  table in the local replica. The values are added with the placeholders in the <code>Exec</code>
  function as positional parameters. Once we have executed the query, this will populate
  the <code>posts</code> table in the lcoal replica.</p>\n<p>We can now finally sync
  with the primary remote LibSQL database to make sure that the primary database also
  has these migrations applied.</p>\n<h3>Syncing the local replica</h3>\n<p>Remember,
  <code>connector</code> is for primary database and <code>db</code> is for the local
  replica. So, we will sync the remote database from the replica that was created
  with the <code>connector.Sync</code></p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"w\">\t</span><span class=\"nx\">_</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">connector</span><span
  class=\"p\">.</span><span class=\"nx\">Sync</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to sync %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Printf</span><span class=\"p\">(</span><span class=\"s\">&quot;Successfully
  synced %s db\\n&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">dbPath</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">rows</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">db</span><span
  class=\"p\">.</span><span class=\"nx\">Query</span><span class=\"p\">(</span><span
  class=\"s\">&quot;SELECT * FROM posts&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to query %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">rows</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">rows</span><span
  class=\"p\">.</span><span class=\"nx\">Next</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">id</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">title</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">description</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">content</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">rows</span><span class=\"p\">.</span><span
  class=\"nx\">Scan</span><span class=\"p\">(</span><span class=\"o\">&amp;</span><span
  class=\"nx\">id</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">title</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">description</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">content</span><span class=\"p\">);</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Fprintf</span><span
  class=\"p\">(</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Stderr</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;failed to scan %s&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Exit</span><span class=\"p\">(</span><span
  class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">id</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">title</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">description</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">content</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Output:</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go                                                            \n\n/tmp/libsql-349052144/local.db\n<span
  class=\"p\">&amp;</span><span class=\"o\">{</span>0x2eec9d0 &lt;nil&gt; &lt;nil&gt;<span
  class=\"o\">}</span>\nConnected to database\nSuccessfully synced /tmp/libsql-349052144/local.db
  db\n<span class=\"m\">1</span> <span class=\"nb\">test</span> title <span class=\"nb\">test</span>
  description <span class=\"nb\">test</span> content\n</pre></div>\n\n</pre>\n\n<p>Once
  we have synced the local replica, we can now query the database i.e. the local replica,
  with the changes, also note that this could also be done without syncing the database,
  but the primary database won't have the applied changes.</p>\n<p>We finally Query
  the local replica with the query <code>SELECT * FROM posts</code> and print out
  the results. This has the 1 record in the <code>posts</code> table that we inserted.</p>\n<p>So,
  that's how we basically create a local replica from a remote LibSQL database hosted
  on Turso. We first create the path for the local database to be copied, then provide
  the primary database URL and credentials, then request a copy of the primary database,
  then we perform any mutation or operations on the local copy and finally sync up
  with the remote primary database to persist the data from that replica (acting like
  a session of database operation).</p>\n<p>That wraps the article for now.</p>\n<p>For
  all the code related to this article, you can check out the <a href=\"https://github.com/mr-destructive/lets-go-with-turso\">Let's
  Go with Turso</a> GitHub repo for all the examples and additional examples for using
  LibSQL with Golang.</p>\n<h2>Conclusion</h2>\n<p>So, that is a wrap for this part
  of the series, we have explored how to create a local embedded replica from a remote
  LibSQL database hosted on Turso with Golang. In the next part of the series, we
  will explore how to setup a local LibSQL database server and then connect it with
  Golang.</p>\n<p>Thank you for reading this post, If you have any questions, feedback,
  and suggestions, feel free to drop them in the comments.</p>\n<p>Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/embedded-replicas-libsql-turso-go.png
long_description: Welcome to the Let With this section, we will specifally dive into
  understanding how to create, connect, and query local embedded replicas of LibSQL
  database hosted on Turso with a Golang application. If you want to check out the
  YouTube video, check
now: 2025-05-01 18:11:33.314262
path: blog/posts/2024-10-31-Turso-LibSQL-Embedded-Replicas-Golang.md
prevnext: null
series:
- lets-go-with-turso
slug: turso-libsql-embedded-replicas-golang
status: published
tags:
- go
- turso
- libsql
templateKey: blog-post
title: Use Embedded Replicas of LibSQL Database hosted on Turso with a Golang Application
today: 2025-05-01
---

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

	fmt.Printf("Successfully synced %s db\n", dbPath)
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