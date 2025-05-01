---
article_html: "<h2>Introduction</h2>\n<p>We all might have used the curl command but
  might be unaware of it. It's super old\nand still serves a great purpose. It has
  been available since 1996 and still is\nwidely used in many embedded technologies,
  web API testing, CLI applications,\netc. In this article, we'll see some basics
  of using the curl command along with\nits applications.</p>\n<h2>What is the curl
  command?</h2>\n<p>Curl or cURL command is the utility or tool to access the internet
  from the command\nline interface using various protocols. This looks trivial but
  it can blow up\nyour mind! Most people use this tool for fetching and processing
  the\ndata from the servers/internet from their terminal without the browser but\nthere
  is a lot more to it. It is used in various embedded devices for accessing\nthe network
  in a lightweight and accessible way. Let's see how you can use the curl\ncommand
  from the very basics.</p>\n<h2>Why do we need it?</h2>\n<p>Before we talk about
  how to use the curl command let's talk about why might we need\nthat? There are
  a lot of reasons and it even depends on the application you are\nusing.  You can
  use curl to test your API, well there are other tools like\nPOSTMAN, Insomnia, etc
  but for keeping things simple you can quickly get in\nwith curl and test some endpoints.
  \ You might require curl for creating some\nCLI applications that require fetching/posting
  to an URL over the internet.\nIf you are using the terminal, curl integrates really
  very well with the shell\nprogramming languages like BASH, ZSH, etc So, after making
  WHY out of the way,\nlet's start with the actual content.</p>\n<h2>Structure of
  curl command</h2>\n<p><strong>curl or Client URL is a command-line utility that
  helps in accessing/posting\ndata with various protocols over the internet.</strong>
  It basically serves as a\nbare-bones browser URL search bar.  You can't render those
  pages like the\nactual GUI, and all but you can get is the HTML source code, JSON
  response,\netc.  That's still quite powerful and used in tons of applications.</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>curl URL arguments \n</pre></div>\n\n</pre>\n\n<p>The
  above is a basic structure of the curl command. We see the argument\nstructure in-depth
  in the next section. Firstly, let's take a simple curl command with just the URL
  is given.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>curl
  <span class=\"s2\">&quot;https://github.com&quot;</span>   \n</pre></div>\n\n</pre>\n\n<p>From
  this query to <code>github.com</code>, you are literally going to <code>GitHub.com</code>
  and getting a response as the entire HTML source code of the page.\nIf you don't
  want to spam the output in the terminal, you can redirect the output to a file.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>curl
  <span class=\"s2\">&quot;https://github.com&quot;</span> &gt;temp.html\n</pre></div>\n\n</pre>\n\n<p>With
  this command, we store the output of the command in the file temp.html, it can be
  any other file you like.</p>\n<h3>Arguments</h3>\n<p>It turns out that you can even
  parse in certain arguments to the <code>curl</code> command to get some desired
  and modified results. Let's take a look at some of them.\nThe <a href=\"https://curl.se/docs/manpage.html\">entire
  list of arguments</a> is quite huge\nand baffling, but this shows how customizable
  the command is.</p>\n<ul>\n<li><code>-s</code> (silent the progress bar)</li>\n<li><code>-X</code>
  (web requests <code>POST, GET, etc</code> to the URL)</li>\n<li><code>-o</code>
  (output to a file)</li>\n<li><code>-H</code> ( provide Header to the request)</li>\n<li><code>-d</code>
  (providing the data e.g. in POST request)</li>\n</ul>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>curl
  -s -o <span class=\"s2\">&quot;https://github.com&quot;</span> temp.html\n</pre></div>\n\n</pre>\n\n<p>This
  command doesn't load the progress bar and simply outputs the response in a\nfile,
  making the execution process in the terminal clean.</p>\n<h3>Integration with other
  commands</h3>\n<p>As said, the <code>curl</code> command can be well integrated
  with the other commands using piping in shell, assigning to variables, and so on.</p>\n<p>Let's
  see how we can convert the <code>JSON</code> response to a BASH variable.</p>\n<pre
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
  class=\"nv\">resp</span><span class=\"o\">=</span><span class=\"k\">$(</span>curl
  -H <span class=\"s2\">&quot;api-key: N2vDzMyEeYGTxjUTePhC8bYd&quot;</span> https://dev.to/api/users/me<span
  class=\"k\">)</span>\n\n<span class=\"nb\">echo</span> <span class=\"nv\">$resp</span>\n</pre></div>\n\n</pre>\n\n<p>Here,
  we are fetching the <code>JSON</code> response from the <code>dev.to</code> <a href=\"https://developers.forem.com/api/\">API</a>,The
  wired string <code>N2vDzMyEeYGTxjUTePhC8bYd</code> is my <a href=\"https://dev.to/settings/account\">dev.to
  API token</a>(don't worry I have revoked it:) ) we have provided an argument <code>-H</code>
  that is a Header for accepting a <code>Json</code> response.\nWe can store the contents
  of the curl command by using the <code>$( )</code> and assigning that to the variable
  name of your choice.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nv\">username</span><span class=\"o\">=</span><span class=\"k\">$(</span>curl
  -H <span class=\"s2\">&quot;api-key: N2vDzMyEeYGTxjUTePhC8bYd&quot;</span> https://dev.to/api/users/me
  <span class=\"p\">|</span> grep -o -P <span class=\"s1\">&#39;(?&lt;=username&quot;:&quot;).*(?=&quot;,&quot;name)&#39;</span><span
  class=\"k\">)</span>\n</pre></div>\n\n</pre>\n\n<p>Here, we have stored the username
  from a <code>JSON</code> response to the variable username. We have piped the curl
  command so that we can work with that <code>JSON</code> response and modify the
  contents and then store the final results in a variable.\nIn this case, we are using
  <code>grep</code> to filter out the content between the key <code>username</code>
  and <code>name</code>, thus we get the value we desired. To see the value you can
  always run the echo command as below:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">echo</span> <span class=\"nv\">$username</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  that's how the <code>curl</code> command integrates flawlessly with BASH and other
  shell programming languages.</p>\n<h2>Where is it used?</h2>\n<p><code>curl</code>
  is actually used in API testing, CLI applications, Web Scrapping, etc. It's a great
  tool for terminal lovers. Let's see where we can use the curl command actually to
  make some good projects.</p>\n<h3>API Testing</h3>\n<p>We can use, <code>curl</code>
  to test an API, it might be API you would have made or to simply test and play with
  other API available publicly. You can get an in-depth guide about <a href=\"https://www.codepedia.org/ama/how-to-test-a-rest-api-from-command-line-with-curl/\">Testing
  a REST API with curl</a>.\nActually, curl can do more than just testing, I have
  made a <a href=\"https://gist.github.com/Mr-Destructive/80860664b1014ef0b94092d68ead1044\">bash
  script</a> that actually posts some data over a database through the API so that
  I don't have to do that manually. That is the kind of automation I love to do and
  curl! Just did that.</p>\n<h3>Web Scrapping</h3>\n<p>Web-scrapping is usually trending
  with Python, but I have done that with BASH.\nThat might be an outdated idea but
  is a good task to learn the basics of\nWeb-scrapping with BASH ;). I must say that
  sed, awk, grep are the tools are\npowerful like heck in doing these tricks. I have
  made this\n<a href=\"https://mr-destructive.github.io/techstructive-blog/bash/2021/07/15/BASH-Crypto-Coingecko.html\">crypto-currency</a>\nand\n<a
  href=\"https://mr-destructive.github.io/techstructive-blog/bash/2021/07/27/BASH-script-dictionary-scrap.html\">dictionary</a>\nscrapper
  with BASH. Web-scrapping can be done with the curl command by fetching to\nan API
  if any or any website. We need to search and find the particular fields,\nclasses,
  or ids the elements the required data might be into and then extract\nand filter
  using the tools like grep, sed or awk.</p>\n<h3>CLI Applications</h3>\n<p>We can
  make CLI applications like creating a terminal view of existing\napplications using
  their APIs or website. I recently made a CLI for\n<a href=\"https://github.com/Mr-Destructive/crossposter\">cross-posting
  articles</a> to\ndev. to, hashnode and medium. That is a project still in progress(tons
  of bugs)\nbut still serving a decent job. Definitely <code>curl</code> might not
  be the only command\nthat works here, but the project might look so incomplete without
  <code>curl</code>.</p>\n<p><strong>There might be other applications as well, who
  knows there is a lot to do with this command.</strong> If you know one, please let
  everyone know in the comments.</p>\n<h3>References:</h3>\n<p>Special Thanks to the
  creator of the curl command - <a href=\"https://github.com/bagder\">Magnus Daniel
  Stenberg</a> and the developers who are still contributing and maintaining the great
  project.</p>\n<h3>Conclusion</h3>\n<p>So, from this article, we were able to understand
  the basics of the <code>curl</code> command and understand its applications in actual
  programming stuff. Hope you liked it. Thanks for reading and until then Happy Coding
  :)</p>\n"
cover: ''
date: 2021-11-05
datetime: 2021-11-05 00:00:00+00:00
description: We all might have used the curl command but might be unaware of it. It
  Curl or cURL command is the utility or tool to access the internet from the command
  Befor
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-11-05-CURL-Basics.md
html: "<h2>Introduction</h2>\n<p>We all might have used the curl command but might
  be unaware of it. It's super old\nand still serves a great purpose. It has been
  available since 1996 and still is\nwidely used in many embedded technologies, web
  API testing, CLI applications,\netc. In this article, we'll see some basics of using
  the curl command along with\nits applications.</p>\n<h2>What is the curl command?</h2>\n<p>Curl
  or cURL command is the utility or tool to access the internet from the command\nline
  interface using various protocols. This looks trivial but it can blow up\nyour mind!
  Most people use this tool for fetching and processing the\ndata from the servers/internet
  from their terminal without the browser but\nthere is a lot more to it. It is used
  in various embedded devices for accessing\nthe network in a lightweight and accessible
  way. Let's see how you can use the curl\ncommand from the very basics.</p>\n<h2>Why
  do we need it?</h2>\n<p>Before we talk about how to use the curl command let's talk
  about why might we need\nthat? There are a lot of reasons and it even depends on
  the application you are\nusing.  You can use curl to test your API, well there are
  other tools like\nPOSTMAN, Insomnia, etc but for keeping things simple you can quickly
  get in\nwith curl and test some endpoints.  You might require curl for creating
  some\nCLI applications that require fetching/posting to an URL over the internet.\nIf
  you are using the terminal, curl integrates really very well with the shell\nprogramming
  languages like BASH, ZSH, etc So, after making WHY out of the way,\nlet's start
  with the actual content.</p>\n<h2>Structure of curl command</h2>\n<p><strong>curl
  or Client URL is a command-line utility that helps in accessing/posting\ndata with
  various protocols over the internet.</strong> It basically serves as a\nbare-bones
  browser URL search bar.  You can't render those pages like the\nactual GUI, and
  all but you can get is the HTML source code, JSON response,\netc.  That's still
  quite powerful and used in tons of applications.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>curl URL arguments \n</pre></div>\n\n</pre>\n\n<p>The
  above is a basic structure of the curl command. We see the argument\nstructure in-depth
  in the next section. Firstly, let's take a simple curl command with just the URL
  is given.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>curl
  <span class=\"s2\">&quot;https://github.com&quot;</span>   \n</pre></div>\n\n</pre>\n\n<p>From
  this query to <code>github.com</code>, you are literally going to <code>GitHub.com</code>
  and getting a response as the entire HTML source code of the page.\nIf you don't
  want to spam the output in the terminal, you can redirect the output to a file.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>curl
  <span class=\"s2\">&quot;https://github.com&quot;</span> &gt;temp.html\n</pre></div>\n\n</pre>\n\n<p>With
  this command, we store the output of the command in the file temp.html, it can be
  any other file you like.</p>\n<h3>Arguments</h3>\n<p>It turns out that you can even
  parse in certain arguments to the <code>curl</code> command to get some desired
  and modified results. Let's take a look at some of them.\nThe <a href=\"https://curl.se/docs/manpage.html\">entire
  list of arguments</a> is quite huge\nand baffling, but this shows how customizable
  the command is.</p>\n<ul>\n<li><code>-s</code> (silent the progress bar)</li>\n<li><code>-X</code>
  (web requests <code>POST, GET, etc</code> to the URL)</li>\n<li><code>-o</code>
  (output to a file)</li>\n<li><code>-H</code> ( provide Header to the request)</li>\n<li><code>-d</code>
  (providing the data e.g. in POST request)</li>\n</ul>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>curl
  -s -o <span class=\"s2\">&quot;https://github.com&quot;</span> temp.html\n</pre></div>\n\n</pre>\n\n<p>This
  command doesn't load the progress bar and simply outputs the response in a\nfile,
  making the execution process in the terminal clean.</p>\n<h3>Integration with other
  commands</h3>\n<p>As said, the <code>curl</code> command can be well integrated
  with the other commands using piping in shell, assigning to variables, and so on.</p>\n<p>Let's
  see how we can convert the <code>JSON</code> response to a BASH variable.</p>\n<pre
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
  class=\"nv\">resp</span><span class=\"o\">=</span><span class=\"k\">$(</span>curl
  -H <span class=\"s2\">&quot;api-key: N2vDzMyEeYGTxjUTePhC8bYd&quot;</span> https://dev.to/api/users/me<span
  class=\"k\">)</span>\n\n<span class=\"nb\">echo</span> <span class=\"nv\">$resp</span>\n</pre></div>\n\n</pre>\n\n<p>Here,
  we are fetching the <code>JSON</code> response from the <code>dev.to</code> <a href=\"https://developers.forem.com/api/\">API</a>,The
  wired string <code>N2vDzMyEeYGTxjUTePhC8bYd</code> is my <a href=\"https://dev.to/settings/account\">dev.to
  API token</a>(don't worry I have revoked it:) ) we have provided an argument <code>-H</code>
  that is a Header for accepting a <code>Json</code> response.\nWe can store the contents
  of the curl command by using the <code>$( )</code> and assigning that to the variable
  name of your choice.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nv\">username</span><span class=\"o\">=</span><span class=\"k\">$(</span>curl
  -H <span class=\"s2\">&quot;api-key: N2vDzMyEeYGTxjUTePhC8bYd&quot;</span> https://dev.to/api/users/me
  <span class=\"p\">|</span> grep -o -P <span class=\"s1\">&#39;(?&lt;=username&quot;:&quot;).*(?=&quot;,&quot;name)&#39;</span><span
  class=\"k\">)</span>\n</pre></div>\n\n</pre>\n\n<p>Here, we have stored the username
  from a <code>JSON</code> response to the variable username. We have piped the curl
  command so that we can work with that <code>JSON</code> response and modify the
  contents and then store the final results in a variable.\nIn this case, we are using
  <code>grep</code> to filter out the content between the key <code>username</code>
  and <code>name</code>, thus we get the value we desired. To see the value you can
  always run the echo command as below:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">echo</span> <span class=\"nv\">$username</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  that's how the <code>curl</code> command integrates flawlessly with BASH and other
  shell programming languages.</p>\n<h2>Where is it used?</h2>\n<p><code>curl</code>
  is actually used in API testing, CLI applications, Web Scrapping, etc. It's a great
  tool for terminal lovers. Let's see where we can use the curl command actually to
  make some good projects.</p>\n<h3>API Testing</h3>\n<p>We can use, <code>curl</code>
  to test an API, it might be API you would have made or to simply test and play with
  other API available publicly. You can get an in-depth guide about <a href=\"https://www.codepedia.org/ama/how-to-test-a-rest-api-from-command-line-with-curl/\">Testing
  a REST API with curl</a>.\nActually, curl can do more than just testing, I have
  made a <a href=\"https://gist.github.com/Mr-Destructive/80860664b1014ef0b94092d68ead1044\">bash
  script</a> that actually posts some data over a database through the API so that
  I don't have to do that manually. That is the kind of automation I love to do and
  curl! Just did that.</p>\n<h3>Web Scrapping</h3>\n<p>Web-scrapping is usually trending
  with Python, but I have done that with BASH.\nThat might be an outdated idea but
  is a good task to learn the basics of\nWeb-scrapping with BASH ;). I must say that
  sed, awk, grep are the tools are\npowerful like heck in doing these tricks. I have
  made this\n<a href=\"https://mr-destructive.github.io/techstructive-blog/bash/2021/07/15/BASH-Crypto-Coingecko.html\">crypto-currency</a>\nand\n<a
  href=\"https://mr-destructive.github.io/techstructive-blog/bash/2021/07/27/BASH-script-dictionary-scrap.html\">dictionary</a>\nscrapper
  with BASH. Web-scrapping can be done with the curl command by fetching to\nan API
  if any or any website. We need to search and find the particular fields,\nclasses,
  or ids the elements the required data might be into and then extract\nand filter
  using the tools like grep, sed or awk.</p>\n<h3>CLI Applications</h3>\n<p>We can
  make CLI applications like creating a terminal view of existing\napplications using
  their APIs or website. I recently made a CLI for\n<a href=\"https://github.com/Mr-Destructive/crossposter\">cross-posting
  articles</a> to\ndev. to, hashnode and medium. That is a project still in progress(tons
  of bugs)\nbut still serving a decent job. Definitely <code>curl</code> might not
  be the only command\nthat works here, but the project might look so incomplete without
  <code>curl</code>.</p>\n<p><strong>There might be other applications as well, who
  knows there is a lot to do with this command.</strong> If you know one, please let
  everyone know in the comments.</p>\n<h3>References:</h3>\n<p>Special Thanks to the
  creator of the curl command - <a href=\"https://github.com/bagder\">Magnus Daniel
  Stenberg</a> and the developers who are still contributing and maintaining the great
  project.</p>\n<h3>Conclusion</h3>\n<p>So, from this article, we were able to understand
  the basics of the <code>curl</code> command and understand its applications in actual
  programming stuff. Hope you liked it. Thanks for reading and until then Happy Coding
  :)</p>\n"
image_url: https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643289075/blogmedia/bqnrrfaeaqfaj7hezxjx.png
long_description: We all might have used the curl command but might be unaware of
  it. It Curl or cURL command is the utility or tool to access the internet from the
  command Before we talk about how to use the curl command let The above is a basic
  structure of the curl
now: 2025-05-01 18:11:33.311581
path: blog/posts/2021-11-05-CURL-Basics.md
prevnext: null
slug: curl-basics
status: published
subtitle: An short simple introductory guide to the curl command
tags:
- bash
- linux
- networking
templateKey: blog-post
title: Basics of curl command
today: 2025-05-01
---

## Introduction

We all might have used the curl command but might be unaware of it. It's super old
and still serves a great purpose. It has been available since 1996 and still is
widely used in many embedded technologies, web API testing, CLI applications,
etc. In this article, we'll see some basics of using the curl command along with
its applications.

## What is the curl command?

Curl or cURL command is the utility or tool to access the internet from the command
line interface using various protocols. This looks trivial but it can blow up
your mind! Most people use this tool for fetching and processing the
data from the servers/internet from their terminal without the browser but
there is a lot more to it. It is used in various embedded devices for accessing
the network in a lightweight and accessible way. Let's see how you can use the curl
command from the very basics.


## Why do we need it?

Before we talk about how to use the curl command let's talk about why might we need
that? There are a lot of reasons and it even depends on the application you are
using.  You can use curl to test your API, well there are other tools like
POSTMAN, Insomnia, etc but for keeping things simple you can quickly get in
with curl and test some endpoints.  You might require curl for creating some
CLI applications that require fetching/posting to an URL over the internet.
If you are using the terminal, curl integrates really very well with the shell
programming languages like BASH, ZSH, etc So, after making WHY out of the way,
let's start with the actual content. 


## Structure of curl command

**curl or Client URL is a command-line utility that helps in accessing/posting
data with various protocols over the internet.** It basically serves as a
bare-bones browser URL search bar.  You can't render those pages like the
actual GUI, and all but you can get is the HTML source code, JSON response,
etc.  That's still quite powerful and used in tons of applications. 

```
curl URL arguments 
```

The above is a basic structure of the curl command. We see the argument
structure in-depth in the next section. Firstly, let's take a simple curl command with just the URL is given.

```bash
curl "https://github.com"   
```
From this query to `github.com`, you are literally going to `GitHub.com` and getting a response as the entire HTML source code of the page.
If you don't want to spam the output in the terminal, you can redirect the output to a file.

```bash
curl "https://github.com" >temp.html
```
With this command, we store the output of the command in the file temp.html, it can be any other file you like. 

### Arguments 

It turns out that you can even parse in certain arguments to the `curl` command to get some desired and modified results. Let's take a look at some of them.
The [entire list of arguments](https://curl.se/docs/manpage.html) is quite huge
and baffling, but this shows how customizable the command is. 

- `-s` (silent the progress bar)
- `-X` (web requests `POST, GET, etc` to the URL)
- `-o` (output to a file)
- `-H` ( provide Header to the request)
- `-d` (providing the data e.g. in POST request)

```bash
curl -s -o "https://github.com" temp.html
```

This command doesn't load the progress bar and simply outputs the response in a
file, making the execution process in the terminal clean.

### Integration with other commands 

As said, the `curl` command can be well integrated with the other commands using piping in shell, assigning to variables, and so on.

Let's see how we can convert the `JSON` response to a BASH variable.

```bash
resp=$(curl -H "api-key: N2vDzMyEeYGTxjUTePhC8bYd" https://dev.to/api/users/me)

echo $resp
```   
Here, we are fetching the `JSON` response from the `dev.to` [API](https://developers.forem.com/api/),The wired string `N2vDzMyEeYGTxjUTePhC8bYd` is my [dev.to API token](https://dev.to/settings/account)(don't worry I have revoked it:) ) we have provided an argument `-H` that is a Header for accepting a `Json` response. 
We can store the contents of the curl command by using the `$( )` and assigning that to the variable name of your choice.

```bash
username=$(curl -H "api-key: N2vDzMyEeYGTxjUTePhC8bYd" https://dev.to/api/users/me | grep -o -P '(?<=username":").*(?=","name)')
```
Here, we have stored the username from a `JSON` response to the variable username. We have piped the curl command so that we can work with that `JSON` response and modify the contents and then store the final results in a variable.
In this case, we are using `grep` to filter out the content between the key `username` and `name`, thus we get the value we desired. To see the value you can always run the echo command as below:
```bash
echo $username
```   
So, that's how the `curl` command integrates flawlessly with BASH and other shell programming languages. 

## Where is it used?

`curl` is actually used in API testing, CLI applications, Web Scrapping, etc. It's a great tool for terminal lovers. Let's see where we can use the curl command actually to make some good projects.

### API Testing

We can use, `curl` to test an API, it might be API you would have made or to simply test and play with other API available publicly. You can get an in-depth guide about [Testing a REST API with curl](https://www.codepedia.org/ama/how-to-test-a-rest-api-from-command-line-with-curl/).
Actually, curl can do more than just testing, I have made a [bash script](https://gist.github.com/Mr-Destructive/80860664b1014ef0b94092d68ead1044) that actually posts some data over a database through the API so that I don't have to do that manually. That is the kind of automation I love to do and curl! Just did that.

### Web Scrapping

Web-scrapping is usually trending with Python, but I have done that with BASH.
That might be an outdated idea but is a good task to learn the basics of
Web-scrapping with BASH ;). I must say that sed, awk, grep are the tools are
powerful like heck in doing these tricks. I have made this
[crypto-currency](https://mr-destructive.github.io/techstructive-blog/bash/2021/07/15/BASH-Crypto-Coingecko.html)
and
[dictionary](https://mr-destructive.github.io/techstructive-blog/bash/2021/07/27/BASH-script-dictionary-scrap.html)
scrapper with BASH. Web-scrapping can be done with the curl command by fetching to
an API if any or any website. We need to search and find the particular fields,
classes, or ids the elements the required data might be into and then extract
and filter using the tools like grep, sed or awk.


### CLI Applications

We can make CLI applications like creating a terminal view of existing
applications using their APIs or website. I recently made a CLI for
[cross-posting articles](https://github.com/Mr-Destructive/crossposter) to
dev. to, hashnode and medium. That is a project still in progress(tons of bugs)
but still serving a decent job. Definitely `curl` might not be the only command
that works here, but the project might look so incomplete without `curl`.

**There might be other applications as well, who knows there is a lot to do with this command.** If you know one, please let everyone know in the comments.

### References:

Special Thanks to the creator of the curl command - [Magnus Daniel Stenberg](https://github.com/bagder) and the developers who are still contributing and maintaining the great project.
 
### Conclusion

So, from this article, we were able to understand the basics of the `curl` command and understand its applications in actual programming stuff. Hope you liked it. Thanks for reading and until then Happy Coding :)