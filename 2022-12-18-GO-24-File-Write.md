---
article_html: "<h2>Introduction</h2>\n<p>In the 24th post of the series, we will be
  taking a look at how we can perform write operations to a file using golang. We
  will be using the <code>os</code> package in most operations along with <code>bufio</code>
  text manipulations. We will be performing write operations like appending, deleting,
  and replacing a file using golang. We will be heavily leveraging standard library
  packages like <code>os</code>, <code>bufio</code>, <code>bytes</code> and <code>fmt</code>.
  We will also be looking into overwriting and string formatting to a file.</p>\n<h2>Write
  to a File</h2>\n<p>The first part of this section is the basic write operation to
  a file, we assume we are writing to a fresh file and overriding the contents of
  the existing file. The next section will cover the appending of content to the file
  and so on. In this example, we will see how we perform basic writing operations
  to write a string, a slice of string to a file.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">HandleError</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">){</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">str</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;Hi, I am
  a gopher!\\n&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">OpenFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;abc.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">os</span><span class=\"p\">.</span><span class=\"nx\">O_WRONLY</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"mo\">0660</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"c1\">// f, err := os.Create(&quot;abc.txt&quot;)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">HandleError</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">_</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span
  class=\"nx\">Write</span><span class=\"p\">([]</span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">str</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">HandleError</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  cat abc.txt\n\n$ go run main.go\n\n$ cat abc.txt\nHi, I am a gopher!\n</pre></div>\n\n</pre>\n\n<p>So,
  we have used a simple golang script to write to a file that exists/has already been
  created. If you don't want errors while having to write on a file that does not
  exist, use the <a href=\"https://pkg.go.dev/os#Create\">Create</a> method instead
  which is similar to the <code>Open</code> method but creates a file if it doesn't
  exist. We use the <a href=\"https://pkg.go.dev/os#File.Write\">Write</a> method
  to overwrite contents to the file, it takes in a parameter as a slice of byte, so
  we typecast the string <code>str</code> into <code>[]byte</code> using the <code>[]byte(str)</code>
  syntax. Thereby we write the contents of the string into the file. We use the defer
  keyword for closing the file at the end of the script or the end of the main function
  scope.</p>\n<h3>Write slice of strings to file</h3>\n<p>We can even write slice
  of string to file using a for loop and appending each string with a new line character.</p>\n<pre
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">HandleError</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">){</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">f</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Create</span><span class=\"p\">(</span><span
  class=\"s\">&quot;abc.txt&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"c1\">//f, err := os.Open(&quot;abc.txt&quot;,
  os.O_WRONLY, 0660)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">langs</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"kt\">string</span><span
  class=\"p\">{</span><span class=\"s\">&quot;golang&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;python&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;rust&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;javascript&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;ruby&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">lang</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">langs</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">_</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">f</span><span
  class=\"p\">.</span><span class=\"nx\">WriteString</span><span class=\"p\">(</span><span
  class=\"nx\">lang</span><span class=\"w\"> </span><span class=\"o\">+</span><span
  class=\"w\"> </span><span class=\"s\">&quot;\\n&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">HandleError</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  cat abc.txt\n\n$ go run main.go\n\n$ cat abc.txt\ngolang\npython\nrust\njavascript\nruby\n</pre></div>\n\n</pre>\n\n<p>We
  have used the <a href=\"https://pkg.go.dev/os#File.WriteString\">WriteString</a>
  method which will take in a string as a parameter instead of a slice of bytes. So,
  we don't have to type cast into slice of bytes. So, as we can see we have written
  the string slice into a file.</p>\n<h3>Over Write</h3>\n<p>The minimal code to write
  to a file is the <a href=\"https://pkg.go.dev/os#WriteFile\">WriteFile</a> function
  in the <a href=\"https://pkg.go.dev/os\">os</a> package, it overrides the content
  of the file with the provided slice of bytes, the name of the file, and the necessary
  permission to write. The funciton additionally creates a file if it does not exist,
  which is one less reason for the error. Though it returns an error object, the error
  might be created due to not right permissions to write to the file, encoding issues,
  etc.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">data</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"kt\">byte</span><span
  class=\"p\">{</span><span class=\"mi\">115</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">111</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">109</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">101</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">65</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">WriteFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;test.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">data</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"mo\">0660</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">s</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Hello&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">WriteFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;test.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"p\">[]</span><span class=\"nb\">byte</span><span class=\"p\">(</span><span
  class=\"nx\">s</span><span class=\"p\">),</span><span class=\"w\"> </span><span
  class=\"mo\">0660</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\n2022/12/17 19:24:13 &lt;nil&gt;\n2022/12/17 19:24:13 &lt;nil&gt;\n\n$
  cat test.txt\nHello\n</pre></div>\n\n</pre>\n\n<p>So, we have used the <code>WriteFile</code>
  method two times in the script, it first takes in a slice of bytes as it is defined
  as <code>data</code> which corresponds to <code>115 -&gt; s</code>, <code>111 -&gt;
  o</code>, <code>65 -&gt; A</code>, ASCII mapped to strings. The slice of bytes can
  be taken as a string like <code>someA</code> as the literal value of the underlying
  slice of the byte. So, we take that slice of byte and parse it to the second parameter
  of the WriteFile function. The first parameter is a string path of the file we want
  to write the contents to, the third parameter is the file permission. We have set
  it as <code>0660</code> indicating read(4) + write(2) to the group and the user
  and no permission to the other users. The function will return an error if any,
  or else it simply overwrites the data in the file.</p>\n<p>In this case, we have
  called the <code>WriteFile</code> method with string <code>s</code> type cast to
  slice of bytes at the end of the script so we see the file has contents as <code>Hello</code>
  instead of <code>someA</code>. If we reverse the action, we don't see the <code>Hello</code>
  string in the file.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">s</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Hello&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">WriteFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;test.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"p\">[]</span><span class=\"nb\">byte</span><span class=\"p\">(</span><span
  class=\"nx\">s</span><span class=\"p\">),</span><span class=\"w\"> </span><span
  class=\"mo\">0660</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">data</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"kt\">byte</span><span
  class=\"p\">{</span><span class=\"mi\">115</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">111</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">109</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">101</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">65</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">WriteFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;test.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">data</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"mo\">0660</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\n2022/12/17 19:24:13 &lt;nil&gt;\n2022/12/17 19:24:13 &lt;nil&gt;\n\n$
  cat test.txt\nsomeA\n</pre></div>\n\n</pre>\n\n<p>As we can see the <code>Hello</code>
  has been overwritten by <code>someA</code>.</p>\n<h3>Write formatted string</h3>\n<p>We
  can even use fmt to write formatted strings to a file. Just like we can take inputs
  with <code>Scanf</code>, we can use <a href=\"https://pkg.go.dev/fmt#Fprint\">Fprint</a>
  and other similar functions like <a href=\"https://pkg.go.dev/fmt#Fprintf\">Fprintf</a>,
  and <a href=\"https://pkg.go.dev/fmt#Fprintln\">Fprintln</a> functions to print/add
  contents to the file.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">HandleErr</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"kt\">error</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">f</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Create</span><span class=\"p\">(</span><span
  class=\"s\">&quot;temp.txt&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">HandleErr</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">name</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">lang</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">exp</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;John&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;go&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"mi\">2</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">_</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Fprint</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;Hi,
  I am &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">name</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;\\n&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">HandleErr</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Fprintf</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;Language
  of choice: %s.\\n&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">lang</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">HandleErr</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">_</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Fprintln</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;Having&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">exp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;years of experience.&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">HandleErr</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">f</span><span
  class=\"p\">.</span><span class=\"nx\">Close</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  cat temp.txt\ncat: temp.txt: No such file or directory\n\n$ go run format.go\n\n$
  cat test.txt\nHi, I am John.\nLanguage of choice: go.\nHaving 2 years of experience.\n</pre></div>\n\n</pre>\n\n<p>So,
  we can see that we have used all three methods having their own use cases, we can
  use <code>Fprint</code> for simple strings, <code>Fprintf</code> for formatting
  the block of a string with multiple placeholders, and the <code>Fprintln</code>
  which works simply like <code>Fprint</code> but it adds a new line itself, we don't
  need to specify it explicitly.</p>\n<h3>Append</h3>\n<p>If we want to append text
  to a file, we can use the <a href=\"https://pkg.go.dev/os#OpenFile\">OpenFile</a>
  function and provide a few parameters to append the contents instead of overwriting.</p>\n<p>Here,
  we have two steps, open the file and then write the contents in the file. So while
  opening the file, we provide a few options as parameters to make the fine-tuned
  system call like only open for read, write or append modes. These options are defined
  as constant int values in the <a href=\"https://pkg.go.dev/os#pkg-constants\">os
  package</a>.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">HandleError</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"kt\">error</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">s</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;Hello&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">WriteFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;test.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"p\">[]</span><span class=\"nb\">byte</span><span class=\"p\">(</span><span
  class=\"nx\">s</span><span class=\"p\">),</span><span class=\"w\"> </span><span
  class=\"mo\">0660</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">HandleError</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">s</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"s\">&quot;World&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">f</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">OpenFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;test.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">os</span><span class=\"p\">.</span><span class=\"nx\">O_APPEND</span><span
  class=\"p\">|</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">O_WRONLY</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"mo\">0660</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">HandleError</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">_</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span
  class=\"nx\">WriteString</span><span class=\"p\">(</span><span class=\"nx\">s</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">HandleError</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">f</span><span
  class=\"p\">.</span><span class=\"nx\">Close</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\n\n$ cat test.txt\nHelloWorld\n</pre></div>\n\n</pre>\n\n<p>So, from
  the above example, we are able to append text into a file, we have first added the
  <code>Hello</code> world string into the file using the <code>WriteFile</code> method
  to indicate we overwrite the previous contents of the file. We then use the <a href=\"https://pkg.go.dev/os#OpenFile\">OpenFile</a>
  method to open a file provided in the first parameter as a string path. The second
  parameter is the options to be passed for performing operations on the opened file,
  we always have to use them <code>defer</code> to close the file or other resource-locking
  operations.</p>\n<p>We have specified the <code>os.O_WRONLY</code> and the <code>os.O_APPEND</code>
  options indicating we want to write to the file while the file is open and specifically
  append to the file. So this is fine-tuning the opened file operation. We can use
  the ReadFile or WriteFile operation which is just used for simple read and write
  operations respectively.</p>\n<p>We use the <a href=\"https://pkg.go.dev/os#File.WriteString\">WriteString</a>
  method, but we can even use the <a href=\"https://pkg.go.dev/os#File.Write\">Write</a>
  method to write a slice of byte instead. This is just used for exploring the different
  options in the file types of the os package.</p>\n<h3>Append at a specific line</h3>\n<p>We
  can also add content to a specific line or a portion of the file. There are no direct
  functions in golang to do the same, we will have to do some manual fine-tuning of
  file operations to append a particular text at a specific line.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;bufio&quot;</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"s\">&quot;bytes&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">HandleError</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"kt\">error</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"cm\">/* test.txt</span>\n<span class=\"cm\">
  \   Hi</span>\n<span class=\"cm\">    Hello</span>\n<span class=\"cm\">    World</span>\n<span
  class=\"cm\">    Gopher</span>\n<span class=\"cm\">    */</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">f</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">OpenFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;test.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">os</span><span class=\"p\">.</span><span class=\"nx\">O_RDWR</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"mo\">0660</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">f</span><span
  class=\"p\">.</span><span class=\"nx\">Close</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">HandleError</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">m</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">bufio</span><span
  class=\"p\">.</span><span class=\"nx\">NewScanner</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">bytes_till</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">0</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">// line to be
  appended</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">line_till</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">2</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">lines_after</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">lines_till</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">i</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">0</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">m</span><span class=\"p\">.</span><span class=\"nx\">Scan</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">line</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">m</span><span
  class=\"p\">.</span><span class=\"nx\">Text</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">i</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"nx\">line_till</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"nx\">bytes_till</span><span class=\"w\"> </span><span class=\"o\">+=</span><span
  class=\"w\"> </span><span class=\"nx\">bytes</span><span class=\"p\">.</span><span
  class=\"nx\">Count</span><span class=\"p\">([]</span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">line</span><span class=\"p\">),</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"kt\">byte</span><span
  class=\"p\">{})</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">i</span><span
  class=\"w\"> </span><span class=\"p\">&gt;</span><span class=\"w\"> </span><span
  class=\"mi\">0</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t\t\t</span><span class=\"nx\">lines_till</span><span class=\"w\">
  </span><span class=\"o\">+=</span><span class=\"w\"> </span><span class=\"s\">&quot;\\n&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span class=\"nx\">lines_till</span><span
  class=\"w\"> </span><span class=\"o\">+=</span><span class=\"w\"> </span><span class=\"nx\">line</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"p\">}</span><span
  class=\"w\"> </span><span class=\"k\">else</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"nx\">lines_after</span><span class=\"w\"> </span><span class=\"o\">+=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;\\n&quot;</span><span class=\"w\"> </span><span
  class=\"o\">+</span><span class=\"w\"> </span><span class=\"nx\">line</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">i</span><span
  class=\"w\"> </span><span class=\"o\">+=</span><span class=\"w\"> </span><span class=\"mi\">1</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">HandleError</span><span
  class=\"p\">(</span><span class=\"nx\">m</span><span class=\"p\">.</span><span class=\"nx\">Err</span><span
  class=\"p\">())</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">insert_text</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">lines_till</span><span class=\"w\"> </span><span
  class=\"o\">+</span><span class=\"w\"> </span><span class=\"s\">&quot;\\nInserted
  content&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">insert_text_bytes</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bytes</span><span class=\"p\">.</span><span
  class=\"nx\">Count</span><span class=\"p\">([]</span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">insert_text</span><span class=\"p\">),</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"kt\">byte</span><span
  class=\"p\">{})</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">lines_after_bytes</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bytes</span><span class=\"p\">.</span><span
  class=\"nx\">Count</span><span class=\"p\">([]</span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">lines_after</span><span class=\"p\">),</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"kt\">byte</span><span
  class=\"p\">{})</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">WriteFile</span><span class=\"p\">(</span><span class=\"s\">&quot;test.txt&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"p\">[]</span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">insert_text</span><span class=\"p\">),</span><span
  class=\"w\"> </span><span class=\"mo\">0660</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">HandleError</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">_</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">f</span><span
  class=\"p\">.</span><span class=\"nx\">WriteAt</span><span class=\"p\">([]</span><span
  class=\"nb\">byte</span><span class=\"p\">(</span><span class=\"nx\">lines_after</span><span
  class=\"p\">),</span><span class=\"w\"> </span><span class=\"nb\">int64</span><span
  class=\"p\">(</span><span class=\"nx\">lines_after_bytes</span><span class=\"p\">)</span><span
  class=\"o\">+</span><span class=\"nb\">int64</span><span class=\"p\">(</span><span
  class=\"nx\">insert_text_bytes</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">HandleError</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"cm\">/* test.txt</span>\n<span class=\"cm\">
  \   Hi</span>\n<span class=\"cm\">    Hello</span>\n<span class=\"cm\">    Inserted
  content</span>\n<span class=\"cm\">    World</span>\n<span class=\"cm\">    Gopher</span>\n<span
  class=\"cm\">    */</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  cat test.txt \nHi\nHello\nWorld\nGophers\n\n$ go run append.go \n\n$ cat test.txt
  \nHi\nHello\nInserted content\nWorld\nGophers\n</pre></div>\n\n</pre>\n\n<p>We have
  inserted <code>Inserted content</code> after the second line because the <code>line_till</code>
  the variable is set to <code>2</code></p>\n<p>So, in the above example, we have
  used a bunch of packages to append a string or any form of text to a particular
  line. We first read the contents of the file, by using the <code>OpenFile</code>
  method which will open the file with certain permissions. We need to close the file
  at the end of the script so we simply use the <code>defer</code> keyword before
  the <code>f.Close()</code> method call. We then start to scan the file buffer, by
  creating a <code>Scanner</code> object with the <code>NewScanner</code> method.
  Then, with the scanner object of the file content, we then can use the <code>Scan()</code>
  method to scan the file contents line by line. By converting the content at each
  line from a slice of bytes to a string using the <code>Text</code>, we append to
  a string <code>line</code>, this will be used for keeping the count of bytes for
  appending text before the newly inserted text.</p>\n<p>The <code>line_till</code>
  variable is used for the line number from which we want to append to the text after.</p>\n<p>We
  count the bytes for the current line and add it to the <code>bytes_till</code> variable
  indicating the number of bytes there before appending content. We have a simple
  if-else check for the first line that is for appending a new line of characters.
  We append the lines into a single string <code>lines_till</code>. The string <code>insert_text</code>
  is created by appending all the lines before the line number <code>line_till</code>
  with the actual content to be inserted. We calculate the number of bytes using the
  <a href=\"https://pkg.go.dev/bytes#Count\">Count</a> method in the bytes package.
  The separator is kept blank. The <code>lines_after</code> is also been created as
  a single string of lines after the line number in the file.</p>\n<p>We add the <code>insert_text</code>
  (lines before + inserted text) into the file using the <code>WriteFile</code> which
  will override the contents of the file. Then we append the <code>lines_after</code>
  string as a slice of bytes to the <code>insert_text_bytes + lines_after_bytes</code>
  so we get the byte number position to append the <code>lines_after</code> string.</p>\n<p>In
  short, we basically overwrite the file by creating two strings (slice of bytes)
  one which has the lines before the line number with the text to be inserted and
  the second string has all the lines after the line number.</p>\n<h2>Replace text
  in a file</h2>\n<p>Using the <a href=\"https://pkg.go.dev/bytes#Replace\">bytes.Replace</a>
  method, we can first read all the bytes and replace the old with the new text, and
  store them as a slice of bytes. We then write these slices of bytes to the file
  again, so we first read the contents into slices of bytes, replace the content of
  the byte and then overwrite the contents with the slice of bytes. It's quite simple.</p>\n<pre
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;bytes&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">HandleError</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"kt\">error</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">filename</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;test.txt&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">file</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">ReadFile</span><span class=\"p\">(</span><span
  class=\"nx\">filename</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">HandleError</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">old_text</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;Hello\\nWorld&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">new_text</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;Bye&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">new_content</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">bytes</span><span
  class=\"p\">.</span><span class=\"nx\">Replace</span><span class=\"p\">(</span><span
  class=\"nx\">file</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"p\">[]</span><span class=\"nb\">byte</span><span class=\"p\">(</span><span
  class=\"nx\">old_text</span><span class=\"p\">),</span><span class=\"w\"> </span><span
  class=\"p\">[]</span><span class=\"nb\">byte</span><span class=\"p\">(</span><span
  class=\"nx\">new_text</span><span class=\"p\">),</span><span class=\"w\"> </span><span
  class=\"o\">-</span><span class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">WriteFile</span><span class=\"p\">(</span><span
  class=\"nx\">filename</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">new_content</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"mo\">0660</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">HandleError</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  cat test.txt\nHi\nHello\nWorld\nGophers\n\n$ go run main.go\n\n\n$ cat test.txt\nHi\nBye\nGophers\n</pre></div>\n\n</pre>\n\n<p>As
  we can see we have replaced <code>Hello\\nWorld</code> it with <code>Bye</code>.
  So, the [Replace] method in the bytes package takes in the parameters like the slice
  of bytes which should be the actual contents of the file, the old text to be replaced
  with again as a slice of bytes, and the new text to replace also as the slice of
  bytes, the final parameter is the number of replacements to be made. Here <code>-1</code>
  indicates there are no limits on how many replacements can be done, it can be <code>1</code>,
  <code>2</code> for replacing the first n occurrence of the old text, depending on
  how many times you want to replace the content in the file.</p>\n<h2>Delete Text
  from a File</h2>\n<p>We can use the <a href=\"https://pkg.go.dev/os#File.Truncate\">os.Truncate</a>
  method to delete the contents of the file. The <code>Truncate</code> method takes
  in the parameters like the file path string and the size of the file to truncate
  or set to. If we set the second parameter to <code>0</code>, the file size will
  be zero and all the contents will be deleted or removed.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"cm\">/* test.txt</span>\n<span class=\"cm\">    Hi</span>\n<span class=\"cm\">
  \   Hello</span>\n<span class=\"cm\">    World</span>\n<span class=\"cm\">    Gophers</span>\n<span
  class=\"cm\">    */</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Truncate</span><span class=\"p\">(</span><span class=\"s\">&quot;test.txt&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"mi\">0</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"cm\">/* test.txt
  is empty</span>\n<span class=\"cm\">    */</span><span class=\"w\"></span>\n<span
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  cat test.txt\nHi\nHello\nWorld\nGophers\n\n$ go run delete.go\n\n$ cat test.txt\n</pre></div>\n\n</pre>\n\n<p>As
  we can see that the contents of the file are emptied if we set the second parameter(size)
  of the <code>Truncate</code> method as 0.</p>\n<p>We can also set the value of the
  size as the number of bytes to keep, so instead of <code>0</code> we can set it
  to <code>n</code> positive integer to only save the first <code>n</code> bytes in
  the file.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"cm\">/* test.txt</span>\n<span class=\"cm\">    Hi</span>\n<span class=\"cm\">
  \   Hello</span>\n<span class=\"cm\">    World</span>\n<span class=\"cm\">    Gophers</span>\n<span
  class=\"cm\">    */</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Truncate</span><span class=\"p\">(</span><span class=\"s\">&quot;test.txt&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"mi\">6</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"cm\">/* test.txt
  </span>\n<span class=\"cm\">    Hi</span>\n<span class=\"cm\">    Hel</span>\n<span
  class=\"cm\">    */</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  cat test.txt\nHi\nHello\nWorld\nGophers\n\n$ go run delete.go\n\n$ cat test.txt\nHi\nHel\n</pre></div>\n\n</pre>\n\n<p>So,
  in the above example, if we set the size parameter to the Truncate method as <code>6</code>,
  it will keep the size of the file to 6 bytes. So we only see <code>Hi\\nHel</code>,
  the new line is a single byte. The rest of the content is deleted. This is how we
  previously deleted all the bytes from the file by setting the size to <code>0</code>.</p>\n<p>That's
  it from this part. Reference for all the code examples and commands can be found
  in the <a href=\"https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/files/write/\">100
  days of Golang</a> GitHub repository.</p>\n<h2>Conclusion</h2>\n<p>So, from this
  section of the series, we were able to perform write operations on a file using
  golang. We used the packages from the standard library and performed operations
  like write, append, overwrite, delete and replace to a simple text file, but it
  could have been any file format.</p>\n<p>Thank you for reading. If you have any
  queries, questions, or feedback, you can let me know in the discussion below or
  on my social handles. Happy Coding :)</p>\n"
cover: ''
date: 2022-12-18
datetime: 2022-12-18 00:00:00+00:00
description: Using system calls to write files in Golang. Performig writing, appending,
  deleting, replacing operations to a file
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2022-12-18-GO-24-File-Write.md
html: "<h2>Introduction</h2>\n<p>In the 24th post of the series, we will be taking
  a look at how we can perform write operations to a file using golang. We will be
  using the <code>os</code> package in most operations along with <code>bufio</code>
  text manipulations. We will be performing write operations like appending, deleting,
  and replacing a file using golang. We will be heavily leveraging standard library
  packages like <code>os</code>, <code>bufio</code>, <code>bytes</code> and <code>fmt</code>.
  We will also be looking into overwriting and string formatting to a file.</p>\n<h2>Write
  to a File</h2>\n<p>The first part of this section is the basic write operation to
  a file, we assume we are writing to a fresh file and overriding the contents of
  the existing file. The next section will cover the appending of content to the file
  and so on. In this example, we will see how we perform basic writing operations
  to write a string, a slice of string to a file.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">HandleError</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">){</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">str</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;Hi, I am
  a gopher!\\n&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">OpenFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;abc.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">os</span><span class=\"p\">.</span><span class=\"nx\">O_WRONLY</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"mo\">0660</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"c1\">// f, err := os.Create(&quot;abc.txt&quot;)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">HandleError</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">_</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span
  class=\"nx\">Write</span><span class=\"p\">([]</span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">str</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">HandleError</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  cat abc.txt\n\n$ go run main.go\n\n$ cat abc.txt\nHi, I am a gopher!\n</pre></div>\n\n</pre>\n\n<p>So,
  we have used a simple golang script to write to a file that exists/has already been
  created. If you don't want errors while having to write on a file that does not
  exist, use the <a href=\"https://pkg.go.dev/os#Create\">Create</a> method instead
  which is similar to the <code>Open</code> method but creates a file if it doesn't
  exist. We use the <a href=\"https://pkg.go.dev/os#File.Write\">Write</a> method
  to overwrite contents to the file, it takes in a parameter as a slice of byte, so
  we typecast the string <code>str</code> into <code>[]byte</code> using the <code>[]byte(str)</code>
  syntax. Thereby we write the contents of the string into the file. We use the defer
  keyword for closing the file at the end of the script or the end of the main function
  scope.</p>\n<h3>Write slice of strings to file</h3>\n<p>We can even write slice
  of string to file using a for loop and appending each string with a new line character.</p>\n<pre
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">HandleError</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">){</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">f</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Create</span><span class=\"p\">(</span><span
  class=\"s\">&quot;abc.txt&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"c1\">//f, err := os.Open(&quot;abc.txt&quot;,
  os.O_WRONLY, 0660)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">langs</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"kt\">string</span><span
  class=\"p\">{</span><span class=\"s\">&quot;golang&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;python&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;rust&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;javascript&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;ruby&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">lang</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">langs</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">_</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">f</span><span
  class=\"p\">.</span><span class=\"nx\">WriteString</span><span class=\"p\">(</span><span
  class=\"nx\">lang</span><span class=\"w\"> </span><span class=\"o\">+</span><span
  class=\"w\"> </span><span class=\"s\">&quot;\\n&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">HandleError</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  cat abc.txt\n\n$ go run main.go\n\n$ cat abc.txt\ngolang\npython\nrust\njavascript\nruby\n</pre></div>\n\n</pre>\n\n<p>We
  have used the <a href=\"https://pkg.go.dev/os#File.WriteString\">WriteString</a>
  method which will take in a string as a parameter instead of a slice of bytes. So,
  we don't have to type cast into slice of bytes. So, as we can see we have written
  the string slice into a file.</p>\n<h3>Over Write</h3>\n<p>The minimal code to write
  to a file is the <a href=\"https://pkg.go.dev/os#WriteFile\">WriteFile</a> function
  in the <a href=\"https://pkg.go.dev/os\">os</a> package, it overrides the content
  of the file with the provided slice of bytes, the name of the file, and the necessary
  permission to write. The funciton additionally creates a file if it does not exist,
  which is one less reason for the error. Though it returns an error object, the error
  might be created due to not right permissions to write to the file, encoding issues,
  etc.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">data</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"kt\">byte</span><span
  class=\"p\">{</span><span class=\"mi\">115</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">111</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">109</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">101</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">65</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">WriteFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;test.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">data</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"mo\">0660</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">s</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Hello&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">WriteFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;test.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"p\">[]</span><span class=\"nb\">byte</span><span class=\"p\">(</span><span
  class=\"nx\">s</span><span class=\"p\">),</span><span class=\"w\"> </span><span
  class=\"mo\">0660</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\n2022/12/17 19:24:13 &lt;nil&gt;\n2022/12/17 19:24:13 &lt;nil&gt;\n\n$
  cat test.txt\nHello\n</pre></div>\n\n</pre>\n\n<p>So, we have used the <code>WriteFile</code>
  method two times in the script, it first takes in a slice of bytes as it is defined
  as <code>data</code> which corresponds to <code>115 -&gt; s</code>, <code>111 -&gt;
  o</code>, <code>65 -&gt; A</code>, ASCII mapped to strings. The slice of bytes can
  be taken as a string like <code>someA</code> as the literal value of the underlying
  slice of the byte. So, we take that slice of byte and parse it to the second parameter
  of the WriteFile function. The first parameter is a string path of the file we want
  to write the contents to, the third parameter is the file permission. We have set
  it as <code>0660</code> indicating read(4) + write(2) to the group and the user
  and no permission to the other users. The function will return an error if any,
  or else it simply overwrites the data in the file.</p>\n<p>In this case, we have
  called the <code>WriteFile</code> method with string <code>s</code> type cast to
  slice of bytes at the end of the script so we see the file has contents as <code>Hello</code>
  instead of <code>someA</code>. If we reverse the action, we don't see the <code>Hello</code>
  string in the file.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">s</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Hello&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">WriteFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;test.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"p\">[]</span><span class=\"nb\">byte</span><span class=\"p\">(</span><span
  class=\"nx\">s</span><span class=\"p\">),</span><span class=\"w\"> </span><span
  class=\"mo\">0660</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">data</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"kt\">byte</span><span
  class=\"p\">{</span><span class=\"mi\">115</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">111</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">109</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">101</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">65</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">WriteFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;test.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">data</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"mo\">0660</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\n2022/12/17 19:24:13 &lt;nil&gt;\n2022/12/17 19:24:13 &lt;nil&gt;\n\n$
  cat test.txt\nsomeA\n</pre></div>\n\n</pre>\n\n<p>As we can see the <code>Hello</code>
  has been overwritten by <code>someA</code>.</p>\n<h3>Write formatted string</h3>\n<p>We
  can even use fmt to write formatted strings to a file. Just like we can take inputs
  with <code>Scanf</code>, we can use <a href=\"https://pkg.go.dev/fmt#Fprint\">Fprint</a>
  and other similar functions like <a href=\"https://pkg.go.dev/fmt#Fprintf\">Fprintf</a>,
  and <a href=\"https://pkg.go.dev/fmt#Fprintln\">Fprintln</a> functions to print/add
  contents to the file.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">HandleErr</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"kt\">error</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">f</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Create</span><span class=\"p\">(</span><span
  class=\"s\">&quot;temp.txt&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">HandleErr</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">name</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">lang</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">exp</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;John&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;go&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"mi\">2</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">_</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Fprint</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;Hi,
  I am &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">name</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;\\n&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">HandleErr</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Fprintf</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;Language
  of choice: %s.\\n&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">lang</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">HandleErr</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">_</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Fprintln</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;Having&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">exp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;years of experience.&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">HandleErr</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">f</span><span
  class=\"p\">.</span><span class=\"nx\">Close</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  cat temp.txt\ncat: temp.txt: No such file or directory\n\n$ go run format.go\n\n$
  cat test.txt\nHi, I am John.\nLanguage of choice: go.\nHaving 2 years of experience.\n</pre></div>\n\n</pre>\n\n<p>So,
  we can see that we have used all three methods having their own use cases, we can
  use <code>Fprint</code> for simple strings, <code>Fprintf</code> for formatting
  the block of a string with multiple placeholders, and the <code>Fprintln</code>
  which works simply like <code>Fprint</code> but it adds a new line itself, we don't
  need to specify it explicitly.</p>\n<h3>Append</h3>\n<p>If we want to append text
  to a file, we can use the <a href=\"https://pkg.go.dev/os#OpenFile\">OpenFile</a>
  function and provide a few parameters to append the contents instead of overwriting.</p>\n<p>Here,
  we have two steps, open the file and then write the contents in the file. So while
  opening the file, we provide a few options as parameters to make the fine-tuned
  system call like only open for read, write or append modes. These options are defined
  as constant int values in the <a href=\"https://pkg.go.dev/os#pkg-constants\">os
  package</a>.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">HandleError</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"kt\">error</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">s</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;Hello&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">WriteFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;test.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"p\">[]</span><span class=\"nb\">byte</span><span class=\"p\">(</span><span
  class=\"nx\">s</span><span class=\"p\">),</span><span class=\"w\"> </span><span
  class=\"mo\">0660</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">HandleError</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">s</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"s\">&quot;World&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">f</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">OpenFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;test.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">os</span><span class=\"p\">.</span><span class=\"nx\">O_APPEND</span><span
  class=\"p\">|</span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">O_WRONLY</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"mo\">0660</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">HandleError</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">_</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">f</span><span class=\"p\">.</span><span
  class=\"nx\">WriteString</span><span class=\"p\">(</span><span class=\"nx\">s</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">HandleError</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">f</span><span
  class=\"p\">.</span><span class=\"nx\">Close</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\n\n$ cat test.txt\nHelloWorld\n</pre></div>\n\n</pre>\n\n<p>So, from
  the above example, we are able to append text into a file, we have first added the
  <code>Hello</code> world string into the file using the <code>WriteFile</code> method
  to indicate we overwrite the previous contents of the file. We then use the <a href=\"https://pkg.go.dev/os#OpenFile\">OpenFile</a>
  method to open a file provided in the first parameter as a string path. The second
  parameter is the options to be passed for performing operations on the opened file,
  we always have to use them <code>defer</code> to close the file or other resource-locking
  operations.</p>\n<p>We have specified the <code>os.O_WRONLY</code> and the <code>os.O_APPEND</code>
  options indicating we want to write to the file while the file is open and specifically
  append to the file. So this is fine-tuning the opened file operation. We can use
  the ReadFile or WriteFile operation which is just used for simple read and write
  operations respectively.</p>\n<p>We use the <a href=\"https://pkg.go.dev/os#File.WriteString\">WriteString</a>
  method, but we can even use the <a href=\"https://pkg.go.dev/os#File.Write\">Write</a>
  method to write a slice of byte instead. This is just used for exploring the different
  options in the file types of the os package.</p>\n<h3>Append at a specific line</h3>\n<p>We
  can also add content to a specific line or a portion of the file. There are no direct
  functions in golang to do the same, we will have to do some manual fine-tuning of
  file operations to append a particular text at a specific line.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;bufio&quot;</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"s\">&quot;bytes&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">HandleError</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"kt\">error</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"cm\">/* test.txt</span>\n<span class=\"cm\">
  \   Hi</span>\n<span class=\"cm\">    Hello</span>\n<span class=\"cm\">    World</span>\n<span
  class=\"cm\">    Gopher</span>\n<span class=\"cm\">    */</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">f</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">OpenFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;test.txt&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">os</span><span class=\"p\">.</span><span class=\"nx\">O_RDWR</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"mo\">0660</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">f</span><span
  class=\"p\">.</span><span class=\"nx\">Close</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">HandleError</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">m</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">bufio</span><span
  class=\"p\">.</span><span class=\"nx\">NewScanner</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">bytes_till</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">0</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">// line to be
  appended</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">line_till</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">2</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">lines_after</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">lines_till</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">i</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">0</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">m</span><span class=\"p\">.</span><span class=\"nx\">Scan</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">line</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">m</span><span
  class=\"p\">.</span><span class=\"nx\">Text</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">i</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"nx\">line_till</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"nx\">bytes_till</span><span class=\"w\"> </span><span class=\"o\">+=</span><span
  class=\"w\"> </span><span class=\"nx\">bytes</span><span class=\"p\">.</span><span
  class=\"nx\">Count</span><span class=\"p\">([]</span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">line</span><span class=\"p\">),</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"kt\">byte</span><span
  class=\"p\">{})</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">i</span><span
  class=\"w\"> </span><span class=\"p\">&gt;</span><span class=\"w\"> </span><span
  class=\"mi\">0</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t\t\t</span><span class=\"nx\">lines_till</span><span class=\"w\">
  </span><span class=\"o\">+=</span><span class=\"w\"> </span><span class=\"s\">&quot;\\n&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span class=\"nx\">lines_till</span><span
  class=\"w\"> </span><span class=\"o\">+=</span><span class=\"w\"> </span><span class=\"nx\">line</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"p\">}</span><span
  class=\"w\"> </span><span class=\"k\">else</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"nx\">lines_after</span><span class=\"w\"> </span><span class=\"o\">+=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;\\n&quot;</span><span class=\"w\"> </span><span
  class=\"o\">+</span><span class=\"w\"> </span><span class=\"nx\">line</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">i</span><span
  class=\"w\"> </span><span class=\"o\">+=</span><span class=\"w\"> </span><span class=\"mi\">1</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">HandleError</span><span
  class=\"p\">(</span><span class=\"nx\">m</span><span class=\"p\">.</span><span class=\"nx\">Err</span><span
  class=\"p\">())</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">insert_text</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">lines_till</span><span class=\"w\"> </span><span
  class=\"o\">+</span><span class=\"w\"> </span><span class=\"s\">&quot;\\nInserted
  content&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">insert_text_bytes</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bytes</span><span class=\"p\">.</span><span
  class=\"nx\">Count</span><span class=\"p\">([]</span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">insert_text</span><span class=\"p\">),</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"kt\">byte</span><span
  class=\"p\">{})</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">lines_after_bytes</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bytes</span><span class=\"p\">.</span><span
  class=\"nx\">Count</span><span class=\"p\">([]</span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">lines_after</span><span class=\"p\">),</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"kt\">byte</span><span
  class=\"p\">{})</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">WriteFile</span><span class=\"p\">(</span><span class=\"s\">&quot;test.txt&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"p\">[]</span><span class=\"nb\">byte</span><span
  class=\"p\">(</span><span class=\"nx\">insert_text</span><span class=\"p\">),</span><span
  class=\"w\"> </span><span class=\"mo\">0660</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">HandleError</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">_</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">f</span><span
  class=\"p\">.</span><span class=\"nx\">WriteAt</span><span class=\"p\">([]</span><span
  class=\"nb\">byte</span><span class=\"p\">(</span><span class=\"nx\">lines_after</span><span
  class=\"p\">),</span><span class=\"w\"> </span><span class=\"nb\">int64</span><span
  class=\"p\">(</span><span class=\"nx\">lines_after_bytes</span><span class=\"p\">)</span><span
  class=\"o\">+</span><span class=\"nb\">int64</span><span class=\"p\">(</span><span
  class=\"nx\">insert_text_bytes</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">HandleError</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"cm\">/* test.txt</span>\n<span class=\"cm\">
  \   Hi</span>\n<span class=\"cm\">    Hello</span>\n<span class=\"cm\">    Inserted
  content</span>\n<span class=\"cm\">    World</span>\n<span class=\"cm\">    Gopher</span>\n<span
  class=\"cm\">    */</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  cat test.txt \nHi\nHello\nWorld\nGophers\n\n$ go run append.go \n\n$ cat test.txt
  \nHi\nHello\nInserted content\nWorld\nGophers\n</pre></div>\n\n</pre>\n\n<p>We have
  inserted <code>Inserted content</code> after the second line because the <code>line_till</code>
  the variable is set to <code>2</code></p>\n<p>So, in the above example, we have
  used a bunch of packages to append a string or any form of text to a particular
  line. We first read the contents of the file, by using the <code>OpenFile</code>
  method which will open the file with certain permissions. We need to close the file
  at the end of the script so we simply use the <code>defer</code> keyword before
  the <code>f.Close()</code> method call. We then start to scan the file buffer, by
  creating a <code>Scanner</code> object with the <code>NewScanner</code> method.
  Then, with the scanner object of the file content, we then can use the <code>Scan()</code>
  method to scan the file contents line by line. By converting the content at each
  line from a slice of bytes to a string using the <code>Text</code>, we append to
  a string <code>line</code>, this will be used for keeping the count of bytes for
  appending text before the newly inserted text.</p>\n<p>The <code>line_till</code>
  variable is used for the line number from which we want to append to the text after.</p>\n<p>We
  count the bytes for the current line and add it to the <code>bytes_till</code> variable
  indicating the number of bytes there before appending content. We have a simple
  if-else check for the first line that is for appending a new line of characters.
  We append the lines into a single string <code>lines_till</code>. The string <code>insert_text</code>
  is created by appending all the lines before the line number <code>line_till</code>
  with the actual content to be inserted. We calculate the number of bytes using the
  <a href=\"https://pkg.go.dev/bytes#Count\">Count</a> method in the bytes package.
  The separator is kept blank. The <code>lines_after</code> is also been created as
  a single string of lines after the line number in the file.</p>\n<p>We add the <code>insert_text</code>
  (lines before + inserted text) into the file using the <code>WriteFile</code> which
  will override the contents of the file. Then we append the <code>lines_after</code>
  string as a slice of bytes to the <code>insert_text_bytes + lines_after_bytes</code>
  so we get the byte number position to append the <code>lines_after</code> string.</p>\n<p>In
  short, we basically overwrite the file by creating two strings (slice of bytes)
  one which has the lines before the line number with the text to be inserted and
  the second string has all the lines after the line number.</p>\n<h2>Replace text
  in a file</h2>\n<p>Using the <a href=\"https://pkg.go.dev/bytes#Replace\">bytes.Replace</a>
  method, we can first read all the bytes and replace the old with the new text, and
  store them as a slice of bytes. We then write these slices of bytes to the file
  again, so we first read the contents into slices of bytes, replace the content of
  the byte and then overwrite the contents with the slice of bytes. It's quite simple.</p>\n<pre
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
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;bytes&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">HandleError</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"kt\">error</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">filename</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;test.txt&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">file</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">ReadFile</span><span class=\"p\">(</span><span
  class=\"nx\">filename</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">HandleError</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">old_text</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;Hello\\nWorld&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">new_text</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;Bye&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">new_content</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">bytes</span><span
  class=\"p\">.</span><span class=\"nx\">Replace</span><span class=\"p\">(</span><span
  class=\"nx\">file</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"p\">[]</span><span class=\"nb\">byte</span><span class=\"p\">(</span><span
  class=\"nx\">old_text</span><span class=\"p\">),</span><span class=\"w\"> </span><span
  class=\"p\">[]</span><span class=\"nb\">byte</span><span class=\"p\">(</span><span
  class=\"nx\">new_text</span><span class=\"p\">),</span><span class=\"w\"> </span><span
  class=\"o\">-</span><span class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">WriteFile</span><span class=\"p\">(</span><span
  class=\"nx\">filename</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">new_content</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"mo\">0660</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">HandleError</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  cat test.txt\nHi\nHello\nWorld\nGophers\n\n$ go run main.go\n\n\n$ cat test.txt\nHi\nBye\nGophers\n</pre></div>\n\n</pre>\n\n<p>As
  we can see we have replaced <code>Hello\\nWorld</code> it with <code>Bye</code>.
  So, the [Replace] method in the bytes package takes in the parameters like the slice
  of bytes which should be the actual contents of the file, the old text to be replaced
  with again as a slice of bytes, and the new text to replace also as the slice of
  bytes, the final parameter is the number of replacements to be made. Here <code>-1</code>
  indicates there are no limits on how many replacements can be done, it can be <code>1</code>,
  <code>2</code> for replacing the first n occurrence of the old text, depending on
  how many times you want to replace the content in the file.</p>\n<h2>Delete Text
  from a File</h2>\n<p>We can use the <a href=\"https://pkg.go.dev/os#File.Truncate\">os.Truncate</a>
  method to delete the contents of the file. The <code>Truncate</code> method takes
  in the parameters like the file path string and the size of the file to truncate
  or set to. If we set the second parameter to <code>0</code>, the file size will
  be zero and all the contents will be deleted or removed.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"cm\">/* test.txt</span>\n<span class=\"cm\">    Hi</span>\n<span class=\"cm\">
  \   Hello</span>\n<span class=\"cm\">    World</span>\n<span class=\"cm\">    Gophers</span>\n<span
  class=\"cm\">    */</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Truncate</span><span class=\"p\">(</span><span class=\"s\">&quot;test.txt&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"mi\">0</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"cm\">/* test.txt
  is empty</span>\n<span class=\"cm\">    */</span><span class=\"w\"></span>\n<span
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  cat test.txt\nHi\nHello\nWorld\nGophers\n\n$ go run delete.go\n\n$ cat test.txt\n</pre></div>\n\n</pre>\n\n<p>As
  we can see that the contents of the file are emptied if we set the second parameter(size)
  of the <code>Truncate</code> method as 0.</p>\n<p>We can also set the value of the
  size as the number of bytes to keep, so instead of <code>0</code> we can set it
  to <code>n</code> positive integer to only save the first <code>n</code> bytes in
  the file.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"cm\">/* test.txt</span>\n<span class=\"cm\">    Hi</span>\n<span class=\"cm\">
  \   Hello</span>\n<span class=\"cm\">    World</span>\n<span class=\"cm\">    Gophers</span>\n<span
  class=\"cm\">    */</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">os</span><span class=\"p\">.</span><span
  class=\"nx\">Truncate</span><span class=\"p\">(</span><span class=\"s\">&quot;test.txt&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"mi\">6</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"cm\">/* test.txt
  </span>\n<span class=\"cm\">    Hi</span>\n<span class=\"cm\">    Hel</span>\n<span
  class=\"cm\">    */</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
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
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>$
  cat test.txt\nHi\nHello\nWorld\nGophers\n\n$ go run delete.go\n\n$ cat test.txt\nHi\nHel\n</pre></div>\n\n</pre>\n\n<p>So,
  in the above example, if we set the size parameter to the Truncate method as <code>6</code>,
  it will keep the size of the file to 6 bytes. So we only see <code>Hi\\nHel</code>,
  the new line is a single byte. The rest of the content is deleted. This is how we
  previously deleted all the bytes from the file by setting the size to <code>0</code>.</p>\n<p>That's
  it from this part. Reference for all the code examples and commands can be found
  in the <a href=\"https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/files/write/\">100
  days of Golang</a> GitHub repository.</p>\n<h2>Conclusion</h2>\n<p>So, from this
  section of the series, we were able to perform write operations on a file using
  golang. We used the packages from the standard library and performed operations
  like write, append, overwrite, delete and replace to a simple text file, but it
  could have been any file format.</p>\n<p>Thank you for reading. If you have any
  queries, questions, or feedback, you can let me know in the discussion below or
  on my social handles. Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/golang-024-file-write.png
long_description: In the 24th post of the series, we will be taking a look at how
  we can perform write operations to a file using golang. We will be using the  The
  first part of this section is the basic write operation to a file, we assume we
  are writing to a fresh f
now: 2025-05-01 18:11:33.313264
path: blog/posts/2022-12-18-GO-24-File-Write.md
prevnext: null
series:
- 100-days-of-golang
slug: golang-file-write
status: published
tags:
- go
templateKey: blog-post
title: 'Golang: File Write'
today: 2025-05-01
---

## Introduction

In the 24th post of the series, we will be taking a look at how we can perform write operations to a file using golang. We will be using the `os` package in most operations along with `bufio` text manipulations. We will be performing write operations like appending, deleting, and replacing a file using golang. We will be heavily leveraging standard library packages like `os`, `bufio`, `bytes` and `fmt`. We will also be looking into overwriting and string formatting to a file.

## Write to a File

The first part of this section is the basic write operation to a file, we assume we are writing to a fresh file and overriding the contents of the existing file. The next section will cover the appending of content to the file and so on. In this example, we will see how we perform basic writing operations to write a string, a slice of string to a file.

```go
package main

import (
	"log"
	"os"
)

func HandleError(err){
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
	str := "Hi, I am a gopher!\n"
	f, err := os.OpenFile("abc.txt", os.O_WRONLY, 0660)
    // f, err := os.Create("abc.txt")
	HandleError(err)
	_, err = f.Write([]byte(str))
	HandleError(err)
	defer f.Close()
}
```

```bash
$ cat abc.txt

$ go run main.go

$ cat abc.txt
Hi, I am a gopher!
```

So, we have used a simple golang script to write to a file that exists/has already been created. If you don't want errors while having to write on a file that does not exist, use the [Create](https://pkg.go.dev/os#Create) method instead which is similar to the `Open` method but creates a file if it doesn't exist. We use the [Write](https://pkg.go.dev/os#File.Write) method to overwrite contents to the file, it takes in a parameter as a slice of byte, so we typecast the string `str` into `[]byte` using the `[]byte(str)` syntax. Thereby we write the contents of the string into the file. We use the defer keyword for closing the file at the end of the script or the end of the main function scope.

### Write slice of strings to file

We can even write slice of string to file using a for loop and appending each string with a new line character.

```go
package main

import (
	"log"
	"os"
)

func HandleError(err){
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
	f, err := os.Create("abc.txt")
	//f, err := os.Open("abc.txt", os.O_WRONLY, 0660)
	langs := []string{"golang", "python", "rust", "javascript", "ruby"}
	for _, lang := range langs {
		_, err := f.WriteString(lang + "\n")
		HandleError(err)
	}
	defer f.Close()
}
```

```plaintext
$ cat abc.txt

$ go run main.go

$ cat abc.txt
golang
python
rust
javascript
ruby
```

We have used the [WriteString](https://pkg.go.dev/os#File.WriteString) method which will take in a string as a parameter instead of a slice of bytes. So, we don't have to type cast into slice of bytes. So, as we can see we have written the string slice into a file.

### Over Write

The minimal code to write to a file is the [WriteFile](https://pkg.go.dev/os#WriteFile) function in the [os](https://pkg.go.dev/os) package, it overrides the content of the file with the provided slice of bytes, the name of the file, and the necessary permission to write. The funciton additionally creates a file if it does not exist, which is one less reason for the error. Though it returns an error object, the error might be created due to not right permissions to write to the file, encoding issues, etc.

```go
package main

import (
	"log"
	"os"
)

func main() {
	data := []byte{115, 111, 109, 101, 65}
	err := os.WriteFile("test.txt", data, 0660)
    log.Println(err)

	s := "Hello"
	err = os.WriteFile("test.txt", []byte(s), 0660)
    log.Println(err)
}
```

```plaintext
$ go run main.go
2022/12/17 19:24:13 <nil>
2022/12/17 19:24:13 <nil>

$ cat test.txt
Hello
```

So, we have used the `WriteFile` method two times in the script, it first takes in a slice of bytes as it is defined as `data` which corresponds to `115 -> s`, `111 -> o`, `65 -> A`, ASCII mapped to strings. The slice of bytes can be taken as a string like `someA` as the literal value of the underlying slice of the byte. So, we take that slice of byte and parse it to the second parameter of the WriteFile function. The first parameter is a string path of the file we want to write the contents to, the third parameter is the file permission. We have set it as `0660` indicating read(4) + write(2) to the group and the user and no permission to the other users. The function will return an error if any, or else it simply overwrites the data in the file.

In this case, we have called the `WriteFile` method with string `s` type cast to slice of bytes at the end of the script so we see the file has contents as `Hello` instead of `someA`. If we reverse the action, we don't see the `Hello` string in the file.

```go
package main

import (
	"log"
	"os"
)

func main() {
	s := "Hello"
	err := os.WriteFile("test.txt", []byte(s), 0660)
    log.Println(err)

	data := []byte{115, 111, 109, 101, 65}
	err = os.WriteFile("test.txt", data, 0660)
    log.Println(err)
}
```

```plaintext
$ go run main.go
2022/12/17 19:24:13 <nil>
2022/12/17 19:24:13 <nil>

$ cat test.txt
someA
```

As we can see the `Hello` has been overwritten by `someA`.

### Write formatted string

We can even use fmt to write formatted strings to a file. Just like we can take inputs with `Scanf`, we can use [Fprint](https://pkg.go.dev/fmt#Fprint) and other similar functions like [Fprintf](https://pkg.go.dev/fmt#Fprintf), and [Fprintln](https://pkg.go.dev/fmt#Fprintln) functions to print/add contents to the file.

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	f, err := os.Create("temp.txt")
	HandleErr(err)
	name, lang, exp := "John", "go", 2
	_, err = fmt.Fprint(f, "Hi, I am ", name, "\n")
	HandleErr(err)
	_, err = fmt.Fprintf(f, "Language of choice: %s.\n", lang)
	HandleErr(err)
	_, err = fmt.Fprintln(f, "Having", exp, "years of experience.")
	HandleErr(err)
	defer f.Close()
}
```

```plaintext
$ cat temp.txt
cat: temp.txt: No such file or directory

$ go run format.go

$ cat test.txt
Hi, I am John.
Language of choice: go.
Having 2 years of experience.
```

So, we can see that we have used all three methods having their own use cases, we can use `Fprint` for simple strings, `Fprintf` for formatting the block of a string with multiple placeholders, and the `Fprintln` which works simply like `Fprint` but it adds a new line itself, we don't need to specify it explicitly.

### Append

If we want to append text to a file, we can use the [OpenFile](https://pkg.go.dev/os#OpenFile) function and provide a few parameters to append the contents instead of overwriting.

Here, we have two steps, open the file and then write the contents in the file. So while opening the file, we provide a few options as parameters to make the fine-tuned system call like only open for read, write or append modes. These options are defined as constant int values in the [os package](https://pkg.go.dev/os#pkg-constants).

```go
package main

import (
	"log"
	"os"
)

func HandleError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
	s := "Hello"
	err := os.WriteFile("test.txt", []byte(s), 0660)
    HandleError(err)

	s = "World"
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0660)
	HandleError(err)
	_, err = f.WriteString(s)
	HandleError(err)
	defer f.Close()
}
```

```plaintext
$ go run main.go

$ cat test.txt
HelloWorld
```

So, from the above example, we are able to append text into a file, we have first added the `Hello` world string into the file using the `WriteFile` method to indicate we overwrite the previous contents of the file. We then use the [OpenFile](https://pkg.go.dev/os#OpenFile) method to open a file provided in the first parameter as a string path. The second parameter is the options to be passed for performing operations on the opened file, we always have to use them `defer` to close the file or other resource-locking operations.

We have specified the `os.O_WRONLY` and the `os.O_APPEND` options indicating we want to write to the file while the file is open and specifically append to the file. So this is fine-tuning the opened file operation. We can use the ReadFile or WriteFile operation which is just used for simple read and write operations respectively.

We use the [WriteString](https://pkg.go.dev/os#File.WriteString) method, but we can even use the [Write](https://pkg.go.dev/os#File.Write) method to write a slice of byte instead. This is just used for exploring the different options in the file types of the os package.

### Append at a specific line

We can also add content to a specific line or a portion of the file. There are no direct functions in golang to do the same, we will have to do some manual fine-tuning of file operations to append a particular text at a specific line.

```go
package main

import (
    "bufio"
    "bytes"
    "log"
    "os"
)

func HandleError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    /* test.txt
    Hi
    Hello
    World
    Gopher
    */
	f, err = os.OpenFile("test.txt", os.O_RDWR, 0660)
	defer f.Close()
	HandleError(err)
	m := bufio.NewScanner(f)
	bytes_till := 0
	// line to be appended
	line_till := 2
	var lines_after string
	var lines_till string
	i := 0
	for m.Scan() {
		line := m.Text()
		if i < line_till {
			bytes_till += bytes.Count([]byte(line), []byte{})
			if i > 0 {
				lines_till += "\n"
			}
			lines_till += line
		} else {
			lines_after += "\n" + line
		}
		i += 1
	}
	HandleError(m.Err())
	insert_text := lines_till + "\nInserted content"
	insert_text_bytes := bytes.Count([]byte(insert_text), []byte{})
	lines_after_bytes := bytes.Count([]byte(lines_after), []byte{})

	err = os.WriteFile("test.txt", []byte(insert_text), 0660)
	HandleError(err)
	_, err = f.WriteAt([]byte(lines_after), int64(lines_after_bytes)+int64(insert_text_bytes))
	HandleError(err)
    /* test.txt
    Hi
    Hello
    Inserted content
    World
    Gopher
    */
}
```

```plaintext
$ cat test.txt 
Hi
Hello
World
Gophers

$ go run append.go 

$ cat test.txt 
Hi
Hello
Inserted content
World
Gophers
```

We have inserted `Inserted content` after the second line because the `line_till` the variable is set to `2`

So, in the above example, we have used a bunch of packages to append a string or any form of text to a particular line. We first read the contents of the file, by using the `OpenFile` method which will open the file with certain permissions. We need to close the file at the end of the script so we simply use the `defer` keyword before the `f.Close()` method call. We then start to scan the file buffer, by creating a `Scanner` object with the `NewScanner` method. Then, with the scanner object of the file content, we then can use the `Scan()` method to scan the file contents line by line. By converting the content at each line from a slice of bytes to a string using the `Text`, we append to a string `line`, this will be used for keeping the count of bytes for appending text before the newly inserted text.

The `line_till` variable is used for the line number from which we want to append to the text after.

We count the bytes for the current line and add it to the `bytes_till` variable indicating the number of bytes there before appending content. We have a simple if-else check for the first line that is for appending a new line of characters. We append the lines into a single string `lines_till`. The string `insert_text` is created by appending all the lines before the line number `line_till` with the actual content to be inserted. We calculate the number of bytes using the [Count](https://pkg.go.dev/bytes#Count) method in the bytes package. The separator is kept blank. The `lines_after` is also been created as a single string of lines after the line number in the file.

We add the `insert_text` (lines before + inserted text) into the file using the `WriteFile` which will override the contents of the file. Then we append the `lines_after` string as a slice of bytes to the `insert_text_bytes + lines_after_bytes` so we get the byte number position to append the `lines_after` string.

In short, we basically overwrite the file by creating two strings (slice of bytes) one which has the lines before the line number with the text to be inserted and the second string has all the lines after the line number.

## Replace text in a file

Using the [bytes.Replace](https://pkg.go.dev/bytes#Replace) method, we can first read all the bytes and replace the old with the new text, and store them as a slice of bytes. We then write these slices of bytes to the file again, so we first read the contents into slices of bytes, replace the content of the byte and then overwrite the contents with the slice of bytes. It's quite simple.

```go
package main

import (
	"bytes"
	"log"
	"os"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	filename := "test.txt"
	file, err := os.ReadFile(filename)
	HandleError(err)
	old_text := "Hello\nWorld"
	new_text := "Bye"
	new_content := bytes.Replace(file, []byte(old_text), []byte(new_text), -1)
	err = os.WriteFile(filename, new_content, 0660)
	HandleError(err)
}
```

```plaintext
$ cat test.txt
Hi
Hello
World
Gophers

$ go run main.go


$ cat test.txt
Hi
Bye
Gophers
```

As we can see we have replaced `Hello\nWorld` it with `Bye`. So, the \[Replace\] method in the bytes package takes in the parameters like the slice of bytes which should be the actual contents of the file, the old text to be replaced with again as a slice of bytes, and the new text to replace also as the slice of bytes, the final parameter is the number of replacements to be made. Here `-1` indicates there are no limits on how many replacements can be done, it can be `1`, `2` for replacing the first n occurrence of the old text, depending on how many times you want to replace the content in the file.

## Delete Text from a File

We can use the [os.Truncate](https://pkg.go.dev/os#File.Truncate) method to delete the contents of the file. The `Truncate` method takes in the parameters like the file path string and the size of the file to truncate or set to. If we set the second parameter to `0`, the file size will be zero and all the contents will be deleted or removed.

```go
package main

import (
	"log"
	"os"
)

func main() {
    /* test.txt
    Hi
    Hello
    World
    Gophers
    */
	err := os.Truncate("test.txt", 0)
	if err != nil {
		log.Fatal(err)
	}
    /* test.txt is empty
    */
}
```

```plaintext
$ cat test.txt
Hi
Hello
World
Gophers

$ go run delete.go

$ cat test.txt
```

As we can see that the contents of the file are emptied if we set the second parameter(size) of the `Truncate` method as 0.

We can also set the value of the size as the number of bytes to keep, so instead of `0` we can set it to `n` positive integer to only save the first `n` bytes in the file.

```go
package main

import (
	"log"
	"os"
)

func main() {
    /* test.txt
    Hi
    Hello
    World
    Gophers
    */
	err := os.Truncate("test.txt", 6)
	if err != nil {
		log.Fatal(err)
	}
    /* test.txt 
    Hi
    Hel
    */
}
```

```plaintext
$ cat test.txt
Hi
Hello
World
Gophers

$ go run delete.go

$ cat test.txt
Hi
Hel
```

So, in the above example, if we set the size parameter to the Truncate method as `6`, it will keep the size of the file to 6 bytes. So we only see `Hi\nHel`, the new line is a single byte. The rest of the content is deleted. This is how we previously deleted all the bytes from the file by setting the size to `0`.

That's it from this part. Reference for all the code examples and commands can be found in the [100 days of Golang](https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/files/write/) GitHub repository.

## Conclusion

So, from this section of the series, we were able to perform write operations on a file using golang. We used the packages from the standard library and performed operations like write, append, overwrite, delete and replace to a simple text file, but it could have been any file format.

Thank you for reading. If you have any queries, questions, or feedback, you can let me know in the discussion below or on my social handles. Happy Coding :)