---
article_html: "<h2>Introduction</h2>\n<p>In the 22nd post of the series, we will be
  looking into the file-handling process in golang, in the next few posts, we will
  cover the operations on file using golang. In this first entry of the file handling
  sub-series, we will understand the <code>READ</code> operation with files. We will
  see different ways to read a file, it can be word by word, line by line, or even
  custom chink by chunk.</p>\n<p>While dealing with files, we will also use standard
  library packages such as <code>os</code>, <code>bufio</code>, etc. We'll also touch
  on how we can read files from a remote location. Using golang, we will have a low-level
  interaction with file management but golang also abstracts the most of heavy lifting
  and management of files for us, so it becomes quite easy to work with files.</p>\n<h2>Read
  the file as a single string (using os.ReadFile)</h2>\n<p>We can use the <a href=\"https://pkg.go.dev/os\">os</a>
  package in golang, in which we have access to the <a href=\"https://pkg.go.dev/os#ReadFile\">ReadFile</a>
  funciton. The <code>ReadFile</code> function takes in a parameter as a string which
  should be a file name, it returns a slice of bytes or an error. We have discussed
  the error handling in the previous part of the series. So, we have to use the comma
  ok error syntax to get the appropriate return value from the funciton. We can grab
  the slice of bytes as the text we want or an error if there are errors like a file
  doesn't exist, it's a folder, etc.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">text</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">ReadFile</span><span class=\"p\">(</span><span class=\"s\">&quot;sample.txt&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nb\">string</span><span class=\"p\">(</span><span class=\"nx\">text</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run main.go                                                                                                                                                         \n2022/10/23
  22:39:11 Golang is a programming language.                                                                                                                    \ncreated:
  2007                                                                                                                                                            \ntype:static
  \n</pre></div>\n\n</pre>\n\n<p>So, under the hood here, we have text as a slice
  of bytes. We can iterate over the text as a slice and get the character by character-in
  the content of the file. Though we don't directly interact with the file content,
  we are storing it in a variable. In technical words, the file is directly loaded
  into the memory at once. We thereby return a single string object containing the
  content of the file.</p>\n<h2>Read file line by line</h2>\n<p>We can even read a
  file line by line. Using the <a href=\"https://pkg.go.dev/bufio#NewScanner\">bufio.NewScanner()</a>,
  the function takes in a <a href=\"https://pkg.go.dev/io#Reader\">Reader</a> object
  in our case it will be a file object. The function returns a scanner object that
  can be used to read the text with a particular scanner method. The returned object
  can be used in the loop to iterate over the content, in our case, we use the <a
  href=\"https://pkg.go.dev/bufio#Scanner.Scan\">Scan</a> method to split the file
  into lines. But we can use other methods like <a href=\"https://pkg.go.dev/bufio#ScanWords\">ScanWords</a>
  for scanning words, <a href=\"https://pkg.go.dev/bufio#ScanRunes\">ScanRunes</a>
  for scanning character by character, <a href=\"https://pkg.go.dev/bufio#ScanBytes\">ScanBytes</a>
  for scanning byte by byte.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;bufio&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;sample.txt&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">line_list</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">string</span><span class=\"p\">{}</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">scanner</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bufio</span><span class=\"p\">.</span><span
  class=\"nx\">NewScanner</span><span class=\"p\">(</span><span class=\"nx\">f</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">scanner</span><span
  class=\"p\">.</span><span class=\"nx\">Scan</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">line</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">scanner</span><span class=\"p\">.</span><span
  class=\"nx\">Text</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">line_list</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nb\">append</span><span
  class=\"p\">(</span><span class=\"nx\">line_list</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">line</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">scanner</span><span
  class=\"p\">.</span><span class=\"nx\">Err</span><span class=\"p\">();</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">line</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">line_list</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">line</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run line.go\n\n2022/10/23
  22:39:50 Golang is a programming language.\n2022/10/23 22:39:50 created: 2007\n2022/10/23
  22:39:50 type:static\n</pre></div>\n\n</pre>\n\n<p>In the above example, the file
  is read in with the <code>bufio.NewScanner</code> and is iterated line by line with
  the help of the <code>Scan</code> function. The text in the line is scanned and
  stored in the variable <code>line</code> which is a string, this is further appended
  to the string slice <code>line_list</code>. Hence we can iterate over the file content
  line by line and store the results as a string array.</p>\n<p>Here we have used
  the <code>defer</code> keyword before calling the <code>f.Close()</code> method
  because we want to close the file after performing operations on the file. The defer
  will call the function at almost the end of the main function i.e. at the end of
  the program.</p>\n<h2>Read File by a delimiter</h2>\n<p>We can even read a file
  with a custom delimiter which can be used to read CSV or other delimiters. The <a
  href=\"https://pkg.go.dev/encoding/csv\">csv</a> package has a <a href=\"https://pkg.go.dev/encoding/csv#NewReader\">NewReader</a>
  function which takes in a object of file content, and it will return a <a href=\"https://pkg.go.dev/encoding/csv#Reader\">Reader</a>
  object. We can alter the attribute <code>Comma</code> in the <code>Reader</code>
  object and set it to any value we want as a delimiter. Thereafter we can read the
  whole content or read it as words, lines, bytes, or chunks as per your criteria.
  The extracted content will be split as a slice of the data separated by the delimiter
  set in the <code>Comma</code> attribute.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;encoding/csv&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;delimeter.txt&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">reader</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">csv</span><span class=\"p\">.</span><span
  class=\"nx\">NewReader</span><span class=\"p\">(</span><span class=\"nx\">f</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">reader</span><span class=\"p\">.</span><span class=\"nx\">Comma</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"sc\">&#39;:&#39;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span>\n<span class=\"w\">    </span><span
  class=\"nx\">data</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">reader</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">data</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ cat delimiter.txt\n\n10:22:2022\ngolang:21:read\n\n\n$
  go run delimiter.go\n\n2022/10/23 22:40:44 [[10 22 2022] [golang 21 read]]\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, the delimiter is set as <code>:</code> so by using the <code>Comma</code>
  attribute we can set the delimiter. By using the <code>NewReader</code> function,
  we fetch the reader object and by using the <code>ReadAll</code> function associated
  to the reader object, we read the contents. The content is fetched as a slice of
  strings which will be separated by the delimiter.</p>\n<h3>Reading File word by
  word</h3>\n<p>We can even use <a href=\"https://pkg.go.dev/bufio#ScanWords\">ScanWords</a>
  to read a file word by word. A word can be a collection of characters that are separated
  by space. Instead of reading it line by line, this function reads the file content
  after a space.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;bufio&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;sample.txt&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">scanner</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bufio</span><span class=\"p\">.</span><span
  class=\"nx\">NewScanner</span><span class=\"p\">(</span><span class=\"nx\">f</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">scanner</span><span class=\"p\">.</span><span class=\"nx\">Split</span><span
  class=\"p\">(</span><span class=\"nx\">bufio</span><span class=\"p\">.</span><span
  class=\"nx\">ScanWords</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">scanner</span><span class=\"p\">.</span><span
  class=\"nx\">Err</span><span class=\"p\">();</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">wordlist</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">string</span><span class=\"p\">{}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">scanner</span><span class=\"p\">.</span><span class=\"nx\">Scan</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">word</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">scanner</span><span
  class=\"p\">.</span><span class=\"nx\">Text</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">wordlist</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nb\">append</span><span
  class=\"p\">(</span><span class=\"nx\">wordlist</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">word</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">word</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">wordlist</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run word.go\n\n2022/10/23
  22:42:03 Golang\n2022/10/23 22:42:03 is\n2022/10/23 22:42:03 a\n2022/10/23 22:42:03
  programming\n2022/10/23 22:42:03 language.\n2022/10/23 22:42:03 created:\n2022/10/23
  22:42:03 2007\n2022/10/23 22:42:03 type:static\n\n2022/10/23 22:42:03 [Golang is
  a programming language. created: 2007 type:static] \n</pre></div>\n\n</pre>\n\n<p>With
  the <code>ScanWords</code> function, we can read the contents of the file word by
  word. The scanner object which has the actual content of the file is split by the
  <code>Split</code> function, the split criteria are used as Word, where the delimiter
  will be used as space. The <code>wordlist</code> is a slice of strings to which
  we append the string <code>word</code> that in turn is read from the <code>scanner.Text()</code>
  function.</p>\n<h3>Reading Files in chunks</h3>\n<p>We can even read files in chunks,
  a chunk is a collection/array of bytes. We can specify the number of bytes we want
  to read in one go and the file reader will scan the content as a slice of that number
  of bytes each iteration. The <a href=\"https://pkg.go.dev/bufio#Reader.Read\">Read</a>
  funciton takes in a slice of bytes and will return the number of bytes in the reader
  object.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;bufio&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;sample.txt&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">reader</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bufio</span><span class=\"p\">.</span><span
  class=\"nx\">NewReader</span><span class=\"p\">(</span><span class=\"nx\">f</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">chunk_size</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">16</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">chunk_list</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">string</span><span class=\"p\">{}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">buf</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nb\">make</span><span
  class=\"p\">([]</span><span class=\"kt\">byte</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">chunk_size</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">n</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">reader</span><span
  class=\"p\">.</span><span class=\"nx\">Read</span><span class=\"p\">(</span><span
  class=\"nx\">buf</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">EOF</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span class=\"k\">break</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">chunk_list</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nb\">append</span><span
  class=\"p\">(</span><span class=\"nx\">chunk_list</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">buf</span><span class=\"p\">[</span><span class=\"mi\">0</span><span
  class=\"p\">:</span><span class=\"nx\">n</span><span class=\"p\">]))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">chunk</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">chunk_list</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Print</span><span class=\"p\">(</span><span
  class=\"nx\">chunk</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run chunks.go\n\n2022/10/23
  22:44:41 Golang is a prog\n2022/10/23 22:44:41 ramming language\n2022/10/23 22:44:41
  .\ncreated: 2007\n2022/10/23 22:44:41 type:static\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we have opened the file and loaded the content into the <code>f</code>
  variable. the contents are read with the help of the <code>NewReader</code> function
  which returns a reader object which further can be used to read contents into chunks
  of bytes. The <code>chunk_size</code> defines the size we want to use for reading
  the content, <code>chunk_list</code> as a slice of strings which will hold the slice
  of chunks/bytes as a type caste into a slice of strings. With the <code>Read</code>
  function, the bytes are read into the function, and the buffer is split as per the
  chunk size obtained in the <code>Read</code> function. We append the slice of bytes
  into the sliced array and thereby we obtain the slice of strings.</p>\n<h3>Read
  file character by character</h3>\n<p>We can even read file each character at a time,
  using the <a href=\"https://pkg.go.dev/bufio#ScanRunes\">ScanRunes</a> function,
  this function scans a single rune/byte at a time. So, we can scan these runes one
  at a time and store them as a slice of bytes. Thereby we will have the content of
  the file stored as a slice of bytes.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;bufio&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;sample.txt&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">scanner</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bufio</span><span class=\"p\">.</span><span
  class=\"nx\">NewScanner</span><span class=\"p\">(</span><span class=\"nx\">f</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">scanner</span><span class=\"p\">.</span><span class=\"nx\">Split</span><span
  class=\"p\">(</span><span class=\"nx\">bufio</span><span class=\"p\">.</span><span
  class=\"nx\">ScanRunes</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">scanner</span><span class=\"p\">.</span><span
  class=\"nx\">Err</span><span class=\"p\">();</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">charlist</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">byte</span><span class=\"p\">{}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">scanner</span><span class=\"p\">.</span><span class=\"nx\">Scan</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">char</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">scanner</span><span class=\"p\">.</span><span
  class=\"nx\">Text</span><span class=\"p\">()[</span><span class=\"mi\">0</span><span
  class=\"p\">])</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">charlist</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nb\">append</span><span class=\"p\">(</span><span
  class=\"nx\">charlist</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">char</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">charlist</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run char.go\n\n2022/10/23
  22:48:55 [71 111 108 97 110 103 32 105 115 32 97 32 112 114 111 103 114 97 109 109
  105 110 103 32 108 97 110 103 117 97 103 101 46 10 99 114 101 97 116 101 100 58
  32 50 48 48 55 10 116 121 112 101 58 115 116 97 116 105 99 10]\n</pre></div>\n\n</pre>\n\n<p>We
  can see in the above example, the output is a slice of byte, hence these are <code>uint8</code>,
  we can cast them to <code>string</code> and obtain the equivalent ASCII representation
  of the bytes. The <code>ScanRunes</code> function allows us to read the content
  from the reader object as a rune as we split the reader object into the unit bytes/runes.</p>\n<p>That's
  it from this part. Reference for all the code examples and commands can be found
  in the <a href=\"https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/files/read/\">100
  days of Golang</a> GitHub repository.</p>\n<h2>Conclusion</h2>\n<p>In this section,
  we explored the functions and packages related to file reading. We saw how we can
  use packages like <code>os</code>, <code>bufio</code>, <code>encoding</code>, etc.
  to read files in a different way. We saw how to read files as a single string, line
  by line, word by word, character by character, in chunks, and also with a custom
  delimiter. Hopefully, the basics of file reading will have been cleared and with
  the examples, the syntactical construct was understood.</p>\n<p>Thank you for reading.
  If you have any queries, questions, or feedback, you can let me know in the discussion
  below or on my social handles. Happy Coding :)</p>\n"
cover: ''
date: 2022-10-23
datetime: 2022-10-23 00:00:00+00:00
description: Using system calls to read files in Golang. Reading files using various
  styles like line by line, word, character, chunk and any other specifc delimiter.
  Using os, bufio, io, encoding packages for helper funcitons making system calls.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2022-10-23-GO-22-File-Reading.md
html: "<h2>Introduction</h2>\n<p>In the 22nd post of the series, we will be looking
  into the file-handling process in golang, in the next few posts, we will cover the
  operations on file using golang. In this first entry of the file handling sub-series,
  we will understand the <code>READ</code> operation with files. We will see different
  ways to read a file, it can be word by word, line by line, or even custom chink
  by chunk.</p>\n<p>While dealing with files, we will also use standard library packages
  such as <code>os</code>, <code>bufio</code>, etc. We'll also touch on how we can
  read files from a remote location. Using golang, we will have a low-level interaction
  with file management but golang also abstracts the most of heavy lifting and management
  of files for us, so it becomes quite easy to work with files.</p>\n<h2>Read the
  file as a single string (using os.ReadFile)</h2>\n<p>We can use the <a href=\"https://pkg.go.dev/os\">os</a>
  package in golang, in which we have access to the <a href=\"https://pkg.go.dev/os#ReadFile\">ReadFile</a>
  funciton. The <code>ReadFile</code> function takes in a parameter as a string which
  should be a file name, it returns a slice of bytes or an error. We have discussed
  the error handling in the previous part of the series. So, we have to use the comma
  ok error syntax to get the appropriate return value from the funciton. We can grab
  the slice of bytes as the text we want or an error if there are errors like a file
  doesn't exist, it's a folder, etc.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">text</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">ReadFile</span><span class=\"p\">(</span><span class=\"s\">&quot;sample.txt&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nb\">string</span><span class=\"p\">(</span><span class=\"nx\">text</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run main.go                                                                                                                                                         \n2022/10/23
  22:39:11 Golang is a programming language.                                                                                                                    \ncreated:
  2007                                                                                                                                                            \ntype:static
  \n</pre></div>\n\n</pre>\n\n<p>So, under the hood here, we have text as a slice
  of bytes. We can iterate over the text as a slice and get the character by character-in
  the content of the file. Though we don't directly interact with the file content,
  we are storing it in a variable. In technical words, the file is directly loaded
  into the memory at once. We thereby return a single string object containing the
  content of the file.</p>\n<h2>Read file line by line</h2>\n<p>We can even read a
  file line by line. Using the <a href=\"https://pkg.go.dev/bufio#NewScanner\">bufio.NewScanner()</a>,
  the function takes in a <a href=\"https://pkg.go.dev/io#Reader\">Reader</a> object
  in our case it will be a file object. The function returns a scanner object that
  can be used to read the text with a particular scanner method. The returned object
  can be used in the loop to iterate over the content, in our case, we use the <a
  href=\"https://pkg.go.dev/bufio#Scanner.Scan\">Scan</a> method to split the file
  into lines. But we can use other methods like <a href=\"https://pkg.go.dev/bufio#ScanWords\">ScanWords</a>
  for scanning words, <a href=\"https://pkg.go.dev/bufio#ScanRunes\">ScanRunes</a>
  for scanning character by character, <a href=\"https://pkg.go.dev/bufio#ScanBytes\">ScanBytes</a>
  for scanning byte by byte.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;bufio&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;sample.txt&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">line_list</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">string</span><span class=\"p\">{}</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">scanner</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bufio</span><span class=\"p\">.</span><span
  class=\"nx\">NewScanner</span><span class=\"p\">(</span><span class=\"nx\">f</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">scanner</span><span
  class=\"p\">.</span><span class=\"nx\">Scan</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">line</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">scanner</span><span class=\"p\">.</span><span
  class=\"nx\">Text</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">line_list</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nb\">append</span><span
  class=\"p\">(</span><span class=\"nx\">line_list</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">line</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">scanner</span><span
  class=\"p\">.</span><span class=\"nx\">Err</span><span class=\"p\">();</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">line</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">line_list</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">line</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run line.go\n\n2022/10/23
  22:39:50 Golang is a programming language.\n2022/10/23 22:39:50 created: 2007\n2022/10/23
  22:39:50 type:static\n</pre></div>\n\n</pre>\n\n<p>In the above example, the file
  is read in with the <code>bufio.NewScanner</code> and is iterated line by line with
  the help of the <code>Scan</code> function. The text in the line is scanned and
  stored in the variable <code>line</code> which is a string, this is further appended
  to the string slice <code>line_list</code>. Hence we can iterate over the file content
  line by line and store the results as a string array.</p>\n<p>Here we have used
  the <code>defer</code> keyword before calling the <code>f.Close()</code> method
  because we want to close the file after performing operations on the file. The defer
  will call the function at almost the end of the main function i.e. at the end of
  the program.</p>\n<h2>Read File by a delimiter</h2>\n<p>We can even read a file
  with a custom delimiter which can be used to read CSV or other delimiters. The <a
  href=\"https://pkg.go.dev/encoding/csv\">csv</a> package has a <a href=\"https://pkg.go.dev/encoding/csv#NewReader\">NewReader</a>
  function which takes in a object of file content, and it will return a <a href=\"https://pkg.go.dev/encoding/csv#Reader\">Reader</a>
  object. We can alter the attribute <code>Comma</code> in the <code>Reader</code>
  object and set it to any value we want as a delimiter. Thereafter we can read the
  whole content or read it as words, lines, bytes, or chunks as per your criteria.
  The extracted content will be split as a slice of the data separated by the delimiter
  set in the <code>Comma</code> attribute.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;encoding/csv&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;delimeter.txt&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">reader</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">csv</span><span class=\"p\">.</span><span
  class=\"nx\">NewReader</span><span class=\"p\">(</span><span class=\"nx\">f</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">reader</span><span class=\"p\">.</span><span class=\"nx\">Comma</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"sc\">&#39;:&#39;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span>\n<span class=\"w\">    </span><span
  class=\"nx\">data</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">reader</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">data</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ cat delimiter.txt\n\n10:22:2022\ngolang:21:read\n\n\n$
  go run delimiter.go\n\n2022/10/23 22:40:44 [[10 22 2022] [golang 21 read]]\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, the delimiter is set as <code>:</code> so by using the <code>Comma</code>
  attribute we can set the delimiter. By using the <code>NewReader</code> function,
  we fetch the reader object and by using the <code>ReadAll</code> function associated
  to the reader object, we read the contents. The content is fetched as a slice of
  strings which will be separated by the delimiter.</p>\n<h3>Reading File word by
  word</h3>\n<p>We can even use <a href=\"https://pkg.go.dev/bufio#ScanWords\">ScanWords</a>
  to read a file word by word. A word can be a collection of characters that are separated
  by space. Instead of reading it line by line, this function reads the file content
  after a space.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;bufio&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;sample.txt&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">scanner</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bufio</span><span class=\"p\">.</span><span
  class=\"nx\">NewScanner</span><span class=\"p\">(</span><span class=\"nx\">f</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">scanner</span><span class=\"p\">.</span><span class=\"nx\">Split</span><span
  class=\"p\">(</span><span class=\"nx\">bufio</span><span class=\"p\">.</span><span
  class=\"nx\">ScanWords</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">scanner</span><span class=\"p\">.</span><span
  class=\"nx\">Err</span><span class=\"p\">();</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">wordlist</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">string</span><span class=\"p\">{}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">scanner</span><span class=\"p\">.</span><span class=\"nx\">Scan</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">word</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">scanner</span><span
  class=\"p\">.</span><span class=\"nx\">Text</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">wordlist</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nb\">append</span><span
  class=\"p\">(</span><span class=\"nx\">wordlist</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">word</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">word</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">wordlist</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run word.go\n\n2022/10/23
  22:42:03 Golang\n2022/10/23 22:42:03 is\n2022/10/23 22:42:03 a\n2022/10/23 22:42:03
  programming\n2022/10/23 22:42:03 language.\n2022/10/23 22:42:03 created:\n2022/10/23
  22:42:03 2007\n2022/10/23 22:42:03 type:static\n\n2022/10/23 22:42:03 [Golang is
  a programming language. created: 2007 type:static] \n</pre></div>\n\n</pre>\n\n<p>With
  the <code>ScanWords</code> function, we can read the contents of the file word by
  word. The scanner object which has the actual content of the file is split by the
  <code>Split</code> function, the split criteria are used as Word, where the delimiter
  will be used as space. The <code>wordlist</code> is a slice of strings to which
  we append the string <code>word</code> that in turn is read from the <code>scanner.Text()</code>
  function.</p>\n<h3>Reading Files in chunks</h3>\n<p>We can even read files in chunks,
  a chunk is a collection/array of bytes. We can specify the number of bytes we want
  to read in one go and the file reader will scan the content as a slice of that number
  of bytes each iteration. The <a href=\"https://pkg.go.dev/bufio#Reader.Read\">Read</a>
  funciton takes in a slice of bytes and will return the number of bytes in the reader
  object.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;bufio&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;sample.txt&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">reader</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bufio</span><span class=\"p\">.</span><span
  class=\"nx\">NewReader</span><span class=\"p\">(</span><span class=\"nx\">f</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">chunk_size</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">16</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">chunk_list</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">string</span><span class=\"p\">{}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">buf</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nb\">make</span><span
  class=\"p\">([]</span><span class=\"kt\">byte</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">chunk_size</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">n</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">reader</span><span
  class=\"p\">.</span><span class=\"nx\">Read</span><span class=\"p\">(</span><span
  class=\"nx\">buf</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">EOF</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span class=\"k\">break</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">chunk_list</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nb\">append</span><span
  class=\"p\">(</span><span class=\"nx\">chunk_list</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">buf</span><span class=\"p\">[</span><span class=\"mi\">0</span><span
  class=\"p\">:</span><span class=\"nx\">n</span><span class=\"p\">]))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">chunk</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">chunk_list</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Print</span><span class=\"p\">(</span><span
  class=\"nx\">chunk</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run chunks.go\n\n2022/10/23
  22:44:41 Golang is a prog\n2022/10/23 22:44:41 ramming language\n2022/10/23 22:44:41
  .\ncreated: 2007\n2022/10/23 22:44:41 type:static\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we have opened the file and loaded the content into the <code>f</code>
  variable. the contents are read with the help of the <code>NewReader</code> function
  which returns a reader object which further can be used to read contents into chunks
  of bytes. The <code>chunk_size</code> defines the size we want to use for reading
  the content, <code>chunk_list</code> as a slice of strings which will hold the slice
  of chunks/bytes as a type caste into a slice of strings. With the <code>Read</code>
  function, the bytes are read into the function, and the buffer is split as per the
  chunk size obtained in the <code>Read</code> function. We append the slice of bytes
  into the sliced array and thereby we obtain the slice of strings.</p>\n<h3>Read
  file character by character</h3>\n<p>We can even read file each character at a time,
  using the <a href=\"https://pkg.go.dev/bufio#ScanRunes\">ScanRunes</a> function,
  this function scans a single rune/byte at a time. So, we can scan these runes one
  at a time and store them as a slice of bytes. Thereby we will have the content of
  the file stored as a slice of bytes.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;bufio&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;sample.txt&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">scanner</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bufio</span><span class=\"p\">.</span><span
  class=\"nx\">NewScanner</span><span class=\"p\">(</span><span class=\"nx\">f</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">scanner</span><span class=\"p\">.</span><span class=\"nx\">Split</span><span
  class=\"p\">(</span><span class=\"nx\">bufio</span><span class=\"p\">.</span><span
  class=\"nx\">ScanRunes</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">scanner</span><span class=\"p\">.</span><span
  class=\"nx\">Err</span><span class=\"p\">();</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">charlist</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">byte</span><span class=\"p\">{}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">scanner</span><span class=\"p\">.</span><span class=\"nx\">Scan</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">char</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">scanner</span><span class=\"p\">.</span><span
  class=\"nx\">Text</span><span class=\"p\">()[</span><span class=\"mi\">0</span><span
  class=\"p\">])</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">charlist</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nb\">append</span><span class=\"p\">(</span><span
  class=\"nx\">charlist</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">char</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">charlist</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run char.go\n\n2022/10/23
  22:48:55 [71 111 108 97 110 103 32 105 115 32 97 32 112 114 111 103 114 97 109 109
  105 110 103 32 108 97 110 103 117 97 103 101 46 10 99 114 101 97 116 101 100 58
  32 50 48 48 55 10 116 121 112 101 58 115 116 97 116 105 99 10]\n</pre></div>\n\n</pre>\n\n<p>We
  can see in the above example, the output is a slice of byte, hence these are <code>uint8</code>,
  we can cast them to <code>string</code> and obtain the equivalent ASCII representation
  of the bytes. The <code>ScanRunes</code> function allows us to read the content
  from the reader object as a rune as we split the reader object into the unit bytes/runes.</p>\n<p>That's
  it from this part. Reference for all the code examples and commands can be found
  in the <a href=\"https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/files/read/\">100
  days of Golang</a> GitHub repository.</p>\n<h2>Conclusion</h2>\n<p>In this section,
  we explored the functions and packages related to file reading. We saw how we can
  use packages like <code>os</code>, <code>bufio</code>, <code>encoding</code>, etc.
  to read files in a different way. We saw how to read files as a single string, line
  by line, word by word, character by character, in chunks, and also with a custom
  delimiter. Hopefully, the basics of file reading will have been cleared and with
  the examples, the syntactical construct was understood.</p>\n<p>Thank you for reading.
  If you have any queries, questions, or feedback, you can let me know in the discussion
  below or on my social handles. Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/golang-022-file-read.png
long_description: 'In the 22nd post of the series, we will be looking into the file-handling
  process in golang, in the next few posts, we will cover the operations on file using
  golang. In this first entry of the file handling sub-series, we will understand
  the  While '
now: 2025-05-01 18:11:33.315039
path: blog/posts/2022-10-23-GO-22-File-Reading.md
prevnext: null
series:
- 100-days-of-golang
slug: golang-file-read
status: published
tags:
- go
templateKey: blog-post
title: 'Golang: File Reading'
today: 2025-05-01
---

## Introduction

In the 22nd post of the series, we will be looking into the file-handling process in golang, in the next few posts, we will cover the operations on file using golang. In this first entry of the file handling sub-series, we will understand the `READ` operation with files. We will see different ways to read a file, it can be word by word, line by line, or even custom chink by chunk. 

While dealing with files, we will also use standard library packages such as `os`, `bufio`, etc. We'll also touch on how we can read files from a remote location. Using golang, we will have a low-level interaction with file management but golang also abstracts the most of heavy lifting and management of files for us, so it becomes quite easy to work with files.


## Read the file as a single string (using os.ReadFile)

We can use the [os](https://pkg.go.dev/os) package in golang, in which we have access to the [ReadFile](https://pkg.go.dev/os#ReadFile) funciton. The `ReadFile` function takes in a parameter as a string which should be a file name, it returns a slice of bytes or an error. We have discussed the error handling in the previous part of the series. So, we have to use the comma ok error syntax to get the appropriate return value from the funciton. We can grab the slice of bytes as the text we want or an error if there are errors like a file doesn't exist, it's a folder, etc. 

```go
package main

import (
	"os"
	"log"
)

func main() {

	text, err := os.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(text))
}
```

```
$ go run main.go                                                                                                                                                         
2022/10/23 22:39:11 Golang is a programming language.                                                                                                                    
created: 2007                                                                                                                                                            
type:static 
```

So, under the hood here, we have text as a slice of bytes. We can iterate over the text as a slice and get the character by character-in the content of the file. Though we don't directly interact with the file content, we are storing it in a variable. In technical words, the file is directly loaded into the memory at once. We thereby return a single string object containing the content of the file.


## Read file line by line 

We can even read a file line by line. Using the [bufio.NewScanner()](https://pkg.go.dev/bufio#NewScanner), the function takes in a [Reader](https://pkg.go.dev/io#Reader) object in our case it will be a file object. The function returns a scanner object that can be used to read the text with a particular scanner method. The returned object can be used in the loop to iterate over the content, in our case, we use the [Scan](https://pkg.go.dev/bufio#Scanner.Scan) method to split the file into lines. But we can use other methods like [ScanWords](https://pkg.go.dev/bufio#ScanWords) for scanning words, [ScanRunes](https://pkg.go.dev/bufio#ScanRunes) for scanning character by character, [ScanBytes](https://pkg.go.dev/bufio#ScanBytes) for scanning byte by byte. 

```go
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	f, err := os.Open("sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	line_list := []string{}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line_list = append(line_list, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, line := range line_list {
		log.Println(line)
	}
}
```

```
$ go run line.go

2022/10/23 22:39:50 Golang is a programming language.
2022/10/23 22:39:50 created: 2007
2022/10/23 22:39:50 type:static
```

In the above example, the file is read in with the `bufio.NewScanner` and is iterated line by line with the help of the `Scan` function. The text in the line is scanned and stored in the variable `line` which is a string, this is further appended to the string slice `line_list`. Hence we can iterate over the file content line by line and store the results as a string array.

Here we have used the `defer` keyword before calling the `f.Close()` method because we want to close the file after performing operations on the file. The defer will call the function at almost the end of the main function i.e. at the end of the program.

## Read File by a delimiter

We can even read a file with a custom delimiter which can be used to read CSV or other delimiters. The [csv](https://pkg.go.dev/encoding/csv) package has a [NewReader](https://pkg.go.dev/encoding/csv#NewReader) function which takes in a object of file content, and it will return a [Reader](https://pkg.go.dev/encoding/csv#Reader) object. We can alter the attribute `Comma` in the `Reader` object and set it to any value we want as a delimiter. Thereafter we can read the whole content or read it as words, lines, bytes, or chunks as per your criteria. The extracted content will be split as a slice of the data separated by the delimiter set in the `Comma` attribute.

```go
package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {

	f, err := os.Open("delimeter.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	reader.Comma = ':'
	
    data, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)
}
```

```
$ cat delimiter.txt

10:22:2022
golang:21:read


$ go run delimiter.go

2022/10/23 22:40:44 [[10 22 2022] [golang 21 read]]
```

In the above example, the delimiter is set as `:` so by using the `Comma` attribute we can set the delimiter. By using the `NewReader` function, we fetch the reader object and by using the `ReadAll` function associated to the reader object, we read the contents. The content is fetched as a slice of strings which will be separated by the delimiter.


### Reading File word by word

We can even use [ScanWords](https://pkg.go.dev/bufio#ScanWords) to read a file word by word. A word can be a collection of characters that are separated by space. Instead of reading it line by line, this function reads the file content after a space. 

```go
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wordlist := []string{}
	for scanner.Scan() {
		word := scanner.Text()
		wordlist = append(wordlist, word)
		log.Println(word)
	}
	log.Println(wordlist)

}
```

```
$ go run word.go

2022/10/23 22:42:03 Golang
2022/10/23 22:42:03 is
2022/10/23 22:42:03 a
2022/10/23 22:42:03 programming
2022/10/23 22:42:03 language.
2022/10/23 22:42:03 created:
2022/10/23 22:42:03 2007
2022/10/23 22:42:03 type:static

2022/10/23 22:42:03 [Golang is a programming language. created: 2007 type:static] 
```

With the `ScanWords` function, we can read the contents of the file word by word. The scanner object which has the actual content of the file is split by the `Split` function, the split criteria are used as Word, where the delimiter will be used as space. The `wordlist` is a slice of strings to which we append the string `word` that in turn is read from the `scanner.Text()` function.

### Reading Files in chunks

We can even read files in chunks, a chunk is a collection/array of bytes. We can specify the number of bytes we want to read in one go and the file reader will scan the content as a slice of that number of bytes each iteration. The [Read](https://pkg.go.dev/bufio#Reader.Read) funciton takes in a slice of bytes and will return the number of bytes in the reader object.

```go
package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	chunk_size := 16
	chunk_list := []string{}
	buf := make([]byte, chunk_size)

	for {
		n, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		chunk_list = append(chunk_list, string(buf[0:n]))
	}
	for _, chunk := range chunk_list {
		log.Print(chunk)
	}
}
```

```
$ go run chunks.go

2022/10/23 22:44:41 Golang is a prog
2022/10/23 22:44:41 ramming language
2022/10/23 22:44:41 .
created: 2007
2022/10/23 22:44:41 type:static
```

In the above example, we have opened the file and loaded the content into the `f` variable. the contents are read with the help of the `NewReader` function which returns a reader object which further can be used to read contents into chunks of bytes. The `chunk_size` defines the size we want to use for reading the content, `chunk_list` as a slice of strings which will hold the slice of chunks/bytes as a type caste into a slice of strings. With the `Read` function, the bytes are read into the function, and the buffer is split as per the chunk size obtained in the `Read` function. We append the slice of bytes into the sliced array and thereby we obtain the slice of strings.


### Read file character by character

We can even read file each character at a time, using the [ScanRunes](https://pkg.go.dev/bufio#ScanRunes) function, this function scans a single rune/byte at a time. So, we can scan these runes one at a time and store them as a slice of bytes. Thereby we will have the content of the file stored as a slice of bytes.

```go
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	charlist := []byte{}
	for scanner.Scan() {
		char := byte(scanner.Text()[0])
		charlist = append(charlist, char)
	}
	log.Println(charlist)

}
```

```
$ go run char.go

2022/10/23 22:48:55 [71 111 108 97 110 103 32 105 115 32 97 32 112 114 111 103 114 97 109 109 105 110 103 32 108 97 110 103 117 97 103 101 46 10 99 114 101 97 116 101 100 58 32 50 48 48 55 10 116 121 112 101 58 115 116 97 116 105 99 10]
```

We can see in the above example, the output is a slice of byte, hence these are `uint8`, we can cast them to `string` and obtain the equivalent ASCII representation of the bytes. The `ScanRunes` function allows us to read the content from the reader object as a rune as we split the reader object into the unit bytes/runes.

That's it from this part. Reference for all the code examples and commands can be found in the [100 days of Golang](https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/files/read/) GitHub repository.

## Conclusion

In this section, we explored the functions and packages related to file reading. We saw how we can use packages like `os`, `bufio`, `encoding`, etc. to read files in a different way. We saw how to read files as a single string, line by line, word by word, character by character, in chunks, and also with a custom delimiter. Hopefully, the basics of file reading will have been cleared and with the examples, the syntactical construct was understood. 

Thank you for reading. If you have any queries, questions, or feedback, you can let me know in the discussion below or on my social handles. Happy Coding :)