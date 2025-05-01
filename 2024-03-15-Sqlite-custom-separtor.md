---
article_html: "<p>explaination about how i learned how to write a command for importing
  a csv\ninto a sqlite3 db with cusotm Separator with a single command inline</p>\n<h2>Introduction</h2>\n<p>I
  was exploring some Issues that I can solve given the context of the problem and
  the intution for the solution.</p>\n<p>I have some github repositories that I always
  look for to learn more. Some of them are:</p>\n<ul>\n<li><a href=\"https://github.com/turbot/steampipe\">Steampipe</a></li>\n<li><a
  href=\"https://github.com/mindsdb/mindsdb\">MindsDB</a></li>\n<li><a href=\"https://github.com/tursodatabase\">Turso</a></li>\n</ul>\n<p>So,
  navigating around for a few minutes, I landed on this <a href=\"https://github.com/tursodatabase/turso-cli/issues/811\">issue</a>.</p>\n<p>The
  issue is really well explained in terms of what the feature is to be added, how
  the usage should be, which for a contributor is half the job done.</p>\n<h2>What?</h2>\n<p>The
  <a href=\"https://github.com/tursodatabase/turso-cli\">turso-cli</a> is a CLI for
  <a href=\"https://github.com/tursodatabase/turso\">Turso</a> that is used to manage
  libsql databases with the <a href=\"https://turso.tech\">turso</a> platform.</p>\n<p>This
  issue is about adding a flag to the <code>db create</code> command to allow the
  user to pass in a custom separator while imporing a csv file into a sqlite3 database.</p>\n<p>The
  only puzzle piece left is to answer the question <code>how</code>?</p>\n<h2>How?</h2>\n<p>So,
  I went in to the Codebase and found where the <code>db create</code> command has
  been handled and landed on this <a href=\"https://sourcegraph.com/github.com/tursodatabase/turso-cli/-/blob/internal/cmd/db_create.go\">file</a></p>\n<p>While
  controbuting to open source, I try to do the small steps and solve try to maintain
  the excitment with that progress. Because if you cannot find the solution, you try
  to procastinate and in the end nothing gets accomplished. So, breaking down the
  problem into small chunks is much helpful than solving the entire problem.</p>\n<p>In
  this case, I  tried to add the flag in the cli which is pretty straight forward.
  We just add one more entry in the list of flags in the <code>db create</code> command.
  This step is just to add a functional CLI for taking the input of csv-separator
  string, however, this won't do anything for the functionality part of things, just
  allow the user to specify the separator/delimitor to use for parsing the CSV file.</p>\n<p>Currently,
  the command that <code>turso db create</code> uses under the hood for creating a
  db from a csv file is:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sqlite3
  <span class=\"s2\">&quot;-csv&quot;</span> <span class=\"s2\">&quot;dbName&quot;</span>
  <span class=\"s2\">&quot;.import &lt;FileName&gt; &lt;TableName&gt;&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above command is found in the <a href=\"https://sourcegraph.com/github.com/tursodatabase/turso-cli/-/blob/internal/cmd/group_flag.go\">group
  flag</a> file. To parse in the separator, we can use another string as <code>.separator
  &quot;;&quot;</code>, here the <code>;</code> is the character that should be used
  as the separator for the csv file into the db.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sqlite3
  <span class=\"s2\">&quot;-csv&quot;</span> <span class=\"s2\">&quot;dbname&quot;</span>
  <span class=\"s2\">&quot;.separator&quot;</span> <span class=\"s2\">&quot;;&quot;</span>
  <span class=\"s2\">&quot;.import &lt;FileName&gt; &lt;TableName&gt;&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>Thank
  you, Happy Coding :)</p>\n<div class='prevnext'>\n    <style type='text/css'>\n
  \   :root {\n      --prevnext-color-text: #eefbfe;\n      --prevnext-color-angle:
  #ff6600;\n      --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"light\"]
  {\n      --prevnext-color-text: #1f2022;\n      --prevnext-color-angle: #ffeb00;\n
  \     --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"dark\"] {\n      --prevnext-color-text:
  #eefbfe;\n      --prevnext-color-angle: #ff6600;\n      --prevnext-subtitle-brightness:
  3;\n    }\n    .prevnext {\n      display: flex;\n      flex-direction: row;\n      justify-content:
  space-around;\n      align-items: flex-start;\n    }\n    .prevnext a {\n      display:
  flex;\n      align-items: center;\n      width: 100%;\n      text-decoration: none;\n
  \   }\n    a.next {\n      justify-content: flex-end;\n    }\n    .prevnext a:hover
  {\n      background: #00000006;\n    }\n    .prevnext-subtitle {\n      color: var(--prevnext-color-text);\n
  \     filter: brightness(var(--prevnext-subtitle-brightness));\n      font-size:
  .8rem;\n    }\n    .prevnext-title {\n      color: var(--prevnext-color-text);\n
  \     font-size: 1rem;\n    }\n    .prevnext-text {\n      max-width: 30vw;\n    }\n
  \   </style>\n<pre><code>&lt;a class='prev' href='/python-chain-assignment-gotcha'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Gotcha
  with Chained Assignment in Python&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a
  class='next' href='/docker-port-forward'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Docker
  Port Forwarding&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot;
  height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
cover: ''
date: 2024-03-15
datetime: 2024-03-15 00:00:00+00:00
description: Explorng SQLite CLI with inline CSV import command with custom separator
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/tils/2024-03-15-Sqlite-custom-separtor.md
html: "<p>explaination about how i learned how to write a command for importing a
  csv\ninto a sqlite3 db with cusotm Separator with a single command inline</p>\n<h2>Introduction</h2>\n<p>I
  was exploring some Issues that I can solve given the context of the problem and
  the intution for the solution.</p>\n<p>I have some github repositories that I always
  look for to learn more. Some of them are:</p>\n<ul>\n<li><a href=\"https://github.com/turbot/steampipe\">Steampipe</a></li>\n<li><a
  href=\"https://github.com/mindsdb/mindsdb\">MindsDB</a></li>\n<li><a href=\"https://github.com/tursodatabase\">Turso</a></li>\n</ul>\n<p>So,
  navigating around for a few minutes, I landed on this <a href=\"https://github.com/tursodatabase/turso-cli/issues/811\">issue</a>.</p>\n<p>The
  issue is really well explained in terms of what the feature is to be added, how
  the usage should be, which for a contributor is half the job done.</p>\n<h2>What?</h2>\n<p>The
  <a href=\"https://github.com/tursodatabase/turso-cli\">turso-cli</a> is a CLI for
  <a href=\"https://github.com/tursodatabase/turso\">Turso</a> that is used to manage
  libsql databases with the <a href=\"https://turso.tech\">turso</a> platform.</p>\n<p>This
  issue is about adding a flag to the <code>db create</code> command to allow the
  user to pass in a custom separator while imporing a csv file into a sqlite3 database.</p>\n<p>The
  only puzzle piece left is to answer the question <code>how</code>?</p>\n<h2>How?</h2>\n<p>So,
  I went in to the Codebase and found where the <code>db create</code> command has
  been handled and landed on this <a href=\"https://sourcegraph.com/github.com/tursodatabase/turso-cli/-/blob/internal/cmd/db_create.go\">file</a></p>\n<p>While
  controbuting to open source, I try to do the small steps and solve try to maintain
  the excitment with that progress. Because if you cannot find the solution, you try
  to procastinate and in the end nothing gets accomplished. So, breaking down the
  problem into small chunks is much helpful than solving the entire problem.</p>\n<p>In
  this case, I  tried to add the flag in the cli which is pretty straight forward.
  We just add one more entry in the list of flags in the <code>db create</code> command.
  This step is just to add a functional CLI for taking the input of csv-separator
  string, however, this won't do anything for the functionality part of things, just
  allow the user to specify the separator/delimitor to use for parsing the CSV file.</p>\n<p>Currently,
  the command that <code>turso db create</code> uses under the hood for creating a
  db from a csv file is:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sqlite3
  <span class=\"s2\">&quot;-csv&quot;</span> <span class=\"s2\">&quot;dbName&quot;</span>
  <span class=\"s2\">&quot;.import &lt;FileName&gt; &lt;TableName&gt;&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above command is found in the <a href=\"https://sourcegraph.com/github.com/tursodatabase/turso-cli/-/blob/internal/cmd/group_flag.go\">group
  flag</a> file. To parse in the separator, we can use another string as <code>.separator
  &quot;;&quot;</code>, here the <code>;</code> is the character that should be used
  as the separator for the csv file into the db.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sqlite3
  <span class=\"s2\">&quot;-csv&quot;</span> <span class=\"s2\">&quot;dbname&quot;</span>
  <span class=\"s2\">&quot;.separator&quot;</span> <span class=\"s2\">&quot;;&quot;</span>
  <span class=\"s2\">&quot;.import &lt;FileName&gt; &lt;TableName&gt;&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>Thank
  you, Happy Coding :)</p>\n<div class='prevnext'>\n    <style type='text/css'>\n
  \   :root {\n      --prevnext-color-text: #eefbfe;\n      --prevnext-color-angle:
  #ff6600;\n      --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"light\"]
  {\n      --prevnext-color-text: #1f2022;\n      --prevnext-color-angle: #ffeb00;\n
  \     --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"dark\"] {\n      --prevnext-color-text:
  #eefbfe;\n      --prevnext-color-angle: #ff6600;\n      --prevnext-subtitle-brightness:
  3;\n    }\n    .prevnext {\n      display: flex;\n      flex-direction: row;\n      justify-content:
  space-around;\n      align-items: flex-start;\n    }\n    .prevnext a {\n      display:
  flex;\n      align-items: center;\n      width: 100%;\n      text-decoration: none;\n
  \   }\n    a.next {\n      justify-content: flex-end;\n    }\n    .prevnext a:hover
  {\n      background: #00000006;\n    }\n    .prevnext-subtitle {\n      color: var(--prevnext-color-text);\n
  \     filter: brightness(var(--prevnext-subtitle-brightness));\n      font-size:
  .8rem;\n    }\n    .prevnext-title {\n      color: var(--prevnext-color-text);\n
  \     font-size: 1rem;\n    }\n    .prevnext-text {\n      max-width: 30vw;\n    }\n
  \   </style>\n<pre><code>&lt;a class='prev' href='/python-chain-assignment-gotcha'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Gotcha
  with Chained Assignment in Python&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a
  class='next' href='/docker-port-forward'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Docker
  Port Forwarding&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot;
  height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
long_description: explaination about how i learned how to write a command for importing
  a csv I was exploring some Issues that I can solve given the context of the problem
  and the intution for the solution. I have some github repositories that I always
  look for to lea
now: 2025-05-01 18:11:33.315182
path: blog/tils/2024-03-15-Sqlite-custom-separtor.md
slug: sqlite-inline-custom-separator
status: published-til
tags:
- sqlite
templateKey: til
title: SQLite importing CSV with custom separator
today: 2025-05-01
---

explaination about how i learned how to write a command for importing a csv 
into a sqlite3 db with cusotm Separator with a single command inline

## Introduction

I was exploring some Issues that I can solve given the context of the problem and the intution for the solution.

I have some github repositories that I always look for to learn more. Some of them are:

- [Steampipe](https://github.com/turbot/steampipe)
- [MindsDB](https://github.com/mindsdb/mindsdb)
- [Turso](https://github.com/tursodatabase)

So, navigating around for a few minutes, I landed on this [issue](https://github.com/tursodatabase/turso-cli/issues/811).

The issue is really well explained in terms of what the feature is to be added, how the usage should be, which for a contributor is half the job done.

## What?

The [turso-cli](https://github.com/tursodatabase/turso-cli) is a CLI for [Turso](https://github.com/tursodatabase/turso) that is used to manage libsql databases with the [turso](https://turso.tech) platform.

This issue is about adding a flag to the `db create` command to allow the user to pass in a custom separator while imporing a csv file into a sqlite3 database.

The only puzzle piece left is to answer the question `how`?

## How?

So, I went in to the Codebase and found where the `db create` command has been handled and landed on this [file](https://sourcegraph.com/github.com/tursodatabase/turso-cli/-/blob/internal/cmd/db_create.go)

While controbuting to open source, I try to do the small steps and solve try to maintain the excitment with that progress. Because if you cannot find the solution, you try to procastinate and in the end nothing gets accomplished. So, breaking down the problem into small chunks is much helpful than solving the entire problem.

In this case, I  tried to add the flag in the cli which is pretty straight forward. We just add one more entry in the list of flags in the `db create` command. This step is just to add a functional CLI for taking the input of csv-separator string, however, this won't do anything for the functionality part of things, just allow the user to specify the separator/delimitor to use for parsing the CSV file.

Currently, the command that `turso db create` uses under the hood for creating a db from a csv file is:

```bash
sqlite3 "-csv" "dbName" ".import <FileName> <TableName>"
```

The above command is found in the [group flag](https://sourcegraph.com/github.com/tursodatabase/turso-cli/-/blob/internal/cmd/group_flag.go) file. To parse in the separator, we can use another string as `.separator ";"`, here the `;` is the character that should be used as the separator for the csv file into the db.

```bash
sqlite3 "-csv" "dbname" ".separator" ";" ".import <FileName> <TableName>"
```

Thank you, Happy Coding :)
<div class='prevnext'>
    <style type='text/css'>
    :root {
      --prevnext-color-text: #eefbfe;
      --prevnext-color-angle: #ff6600;
      --prevnext-subtitle-brightness: 3;
    }
    [data-theme="light"] {
      --prevnext-color-text: #1f2022;
      --prevnext-color-angle: #ffeb00;
      --prevnext-subtitle-brightness: 3;
    }
    [data-theme="dark"] {
      --prevnext-color-text: #eefbfe;
      --prevnext-color-angle: #ff6600;
      --prevnext-subtitle-brightness: 3;
    }
    .prevnext {
      display: flex;
      flex-direction: row;
      justify-content: space-around;
      align-items: flex-start;
    }
    .prevnext a {
      display: flex;
      align-items: center;
      width: 100%;
      text-decoration: none;
    }
    a.next {
      justify-content: flex-end;
    }
    .prevnext a:hover {
      background: #00000006;
    }
    .prevnext-subtitle {
      color: var(--prevnext-color-text);
      filter: brightness(var(--prevnext-subtitle-brightness));
      font-size: .8rem;
    }
    .prevnext-title {
      color: var(--prevnext-color-text);
      font-size: 1rem;
    }
    .prevnext-text {
      max-width: 30vw;
    }
    </style>
    
    <a class='prev' href='/python-chain-assignment-gotcha'>
    
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M13.5 8.25L9.75 12L13.5 15.75" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"> </path>
        </svg>
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>prev</p>
            <p class='prevnext-title'>Gotcha with Chained Assignment in Python</p>
        </div>
    </a>
    
    <a class='next' href='/docker-port-forward'>
    
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>next</p>
            <p class='prevnext-title'>Docker Port Forwarding</p>
        </div>
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.5 15.75L14.25 12L10.5 8.25" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path>
        </svg>
    </a>
  </div>