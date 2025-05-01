---
article_html: "<h2>Reading specific file types (JSON, YAML, TOML)</h2>\n<p>In the
  previous post, we have seen how to read files in golang, in this extended post of
  that part, we will look into reading some specific files used for configuration
  and storing data like JSON, YAML, TOML, CSV, etc.</p>\n<p>We will see how to read
  files and get individual pieces in the files. We'll use packages like <code>os</code>,
  <code>ioutil</code> and <code>encoding</code> to perform reading operations on file
  and file objects.</p>\n<h3>Reading a JSON File</h3>\n<p>Golang has built-in support
  for reading JSON files, but still, we can and need to have low-level controls on
  how to parse and extract content from the file.</p>\n<p>Let's say we have a <code>json</code>
  file named <code>blog.json</code>, we can use the <a href=\"https://pkg.go.dev/encoding/json\">encoding/json</a>
  package to convert the JSON data into a GO object (an object that is native and
  understandable to go). The <a href=\"https://pkg.go.dev/encoding/json#Unmarshal\">Unmarshal</a>
  function is used to convert the slice of bytes from the file, into a map object.</p>\n<pre
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
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;title&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;Golang Blog Series&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;date&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;22nd October
  2022&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nt\">&quot;tags&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"p\">[</span><span class=\"s2\">&quot;go&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s2\">&quot;files&quot;</span><span
  class=\"p\">],</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;words&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"mi\">1500</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nt\">&quot;published&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"kc\">true</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  above is a simple JSON file, this file has a few types of key-value pairs like string,
  list, integer, and boolean. But we can also have nested objects and a list of those
  nested objects.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">f</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">ReadFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;blog.json&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">data</span><span class=\"w\"> </span><span
  class=\"kd\">map</span><span class=\"p\">[</span><span class=\"kt\">string</span><span
  class=\"p\">]</span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">([]</span><span
  class=\"nb\">byte</span><span class=\"p\">(</span><span class=\"nx\">f</span><span
  class=\"p\">),</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">data</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">data</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">k</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">v</span><span class=\"w\">
  </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">data</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">k</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">v</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p><strong>I
  have removed the time stamp from the logs below so as to clearly see the output.
  We can use <code>fmt</code> to print the simple things while keeping consistent
  with the rest of the snippets in the series.</strong></p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run json.go\n\nmap[date:22nd
  October 2022 published:true tags:[go files] title:Golang Blog Series words:1500]\npublished
  : true\ntitle : Golang Blog Series\ndate : 22nd October 2022\ntags : [go files]\nwords
  : 1500\n</pre></div>\n\n</pre>\n\n<p>The file is read using the <a href=\"https://pkg.go.dev/os#ReadFile\">os.ReadFile</a>
  method, that takes in a string as a path to the file and returns a slice of bytes
  or an error if there was an issue in reading the file. The parsed slice of byte
  is than passed as the first argument to the <code>Unmarshal</code> method in the
  <code>encoding/json</code> package. The second parameter is the output reference
  where the parsed JSON will be stored or returned. The function does not return the
  parsed content instead returns an error if there arose any while parsing the JSON
  content.</p>\n<p>As we can see we have got a map of <code>string</code> with an
  <code>interface</code>. The interface is used because the value of the key can be
  anything. There is no fixed value like a <code>string</code>, <code>int</code>,
  <code>bool</code>, or a nested <code>map</code>, <code>slice</code>. Hence we have
  mapped the JSON object as a map of <code>string</code> with an <code>interface</code>.
  The type of the value is identified with the interface it has attached to it. Let's
  take a look what is the type of each value in the map.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>published : true\nbool\n\ntitle
  : Golang Blog Series\nstring\n\ndate : 22nd October 2022\nstring\n\ntags : [go files]\n[]interface
  {}\n\nwords : 1500\nfloat64\n</pre></div>\n\n</pre>\n\n<p>Here, we can see it has
  correctly identified the string type of the fields like bool in case of true or
  false, a string for string type of values, the fourth field however has a list interface
  attached to it. The default precedence for <code>float64</code> over integer is
  the reason the <code>1500</code> value is of type <code>float64</code>.</p>\n<h3>Reading
  a YAML File</h3>\n<p>Though there is no standard package for parsing/unmarshaling
  <code>YAML</code> files in golang, it's quite easy to use a third-party package
  and use it to read YAML files.</p>\n<p>The package <a href=\"https://pkg.go.dev/gopkg.in/yaml.v3\">gopkg.in/yaml.v3</a>
  is used for encoding and decoding YAML files. We'll be just using it for decoding
  a YAML file by reading it and converting the file object into native Go objects
  like maps, lists, strings, etc.</p>\n<p>The below steps can be used for setting
  up the package and installing the YAML package locally.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go mod init &lt;your_project_package_name&gt;\ngo
  get gopkg.in/yaml.v3\n</pre></div>\n\n</pre>\n\n<p>This should create two files
  namely <code>go.mod</code> and <code>go.sum</code> with the dependency of the <code>gopkg.in/yaml.v3</code>
  package.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>yml</div>\n<div class=\"highlight\"><pre><span></span>title:
  &quot;Golang Blog Series&quot;\ndate: &quot;22nd October 2022&quot;\ntags: [&quot;go&quot;,
  &quot;files&quot;]\npublished: false\nwords: 1500\n</pre></div>\n\n</pre>\n\n<p>The
  above file is a simple YAML config, we'll follow similar kind of examples for the
  dummy files used in the examples.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">yaml</span><span class=\"w\"> </span><span
  class=\"s\">&quot;gopkg.in/yaml.v3&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">f</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">ReadFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;blog.yaml&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">data</span><span class=\"w\"> </span><span
  class=\"kd\">map</span><span class=\"p\">[</span><span class=\"kt\">string</span><span
  class=\"p\">]</span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">yaml</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">data</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">data</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">k</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">v</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">data</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">k</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">v</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run yaml.go\n\nmap[date:22nd
  October 2022 published:false tags:[go files] title:Golang Blog Series words:1500]\npublished
  : false\nwords : 1500\ntitle : Golang Blog Series\ndate : 22nd October 2022\ntags
  : [go files]\n</pre></div>\n\n</pre>\n\n<p>The above code and output demonstrate
  the usage of the <code>yaml.v3</code> package for reading a YAML file.</p>\n<p>Firstly,
  we read the file into a single-string object with the <a href=\"https://pkg.go.dev/os#ReadFile\">os.ReadFile()</a>
  method. The method will return a <code>[]byte</code> (slice of byte) or an error.
  If there is an error, we simply log or panic out of the program, else we can use
  the <a href=\"https://pkg.go.dev/gopkg.in/yaml.v3#Unmarshal\">yaml.Unmarshal</a>
  method to convert the string/slice of the byte into a map or a pre-defined struct.
  In this example, we just keep things simple by storing the file content as <code>map
  [string, interface{}]</code>, i.e. a map of <code>string</code> and an <code>interface</code>.
  The key for YAML can be only a string or an integer. It can't have unrestricted
  data types like the value can have. Though if you want to be unrestrictive, you
  can use a map of <code>map[interface{}]interface{}</code> to make the key any shape
  you like to have.</p>\n<p>So, we have created a variable called <code>data</code>
  as a map of <code>string</code> and <code>interface{}</code>, basically key can
  be a string and the value can be any type of interface depending on the parsed literally
  from the file object. The <code>Unmarshal</code> function takes in two parameters,
  the first being the slice of byte i.e. the file contents, and the second being the
  output variable. Now, the method does not return the parsed YAML, it can return
  an error if there arose, so we need to set the second parameter as a pointer to
  the object into which we want to store the parsed YAML.</p>\n<p>In the example,
  we have called <code>Unmarshal(f, &amp;data)</code> which will fetch the contents
  from the slice of bytes <code>f</code> and output the parsed YAML from the slice
  of bytes into the memory location of <code>data</code> and hence using <code>&amp;data</code>
  indicating the pointer to the variable(fetch the memory address).</p>\n<p>So, that
  is how we obtain the map of keys and values from the YAML config, thereafter, you
  can iterate on the map, access the keys and values, type caste them as per requirement,
  and basically have control over what processing needs to be done to the parsed YAML
  content.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n\n<div class=\"highlight\"><pre><span></span>published : false\nbool\n\nwords
  : 1500\nint\n\ntitle : Golang Blog Series\nstring\n\ndate : 22nd October 2022\nstring\n\ntags
  : [go files]\n[]interface {}\n</pre></div>\n\n</pre>\n\n<p>I have just printed the
  types of the values in the above output as <code>log.Printf(&quot;%T&quot;, v)</code>,
  we can see the types are being correctly recognized and being parsed. The last object
  is indeed a <code>slice</code> so it has an interface of the slice(array) attached
  to it.</p>\n<h3>Reading a TOML file</h3>\n<p>YAML and TOML are almost identical,
  though toml has more restrictive configuration options, but is more readable than
  YAML, as YAML can get complicated pretty quickly. Though both of them have their
  pros and cons, YAML is used everywhere in the DevOps world, configs, whereas TOML
  is the format of choice for python packaging, and static site generation configs.</p>\n<p>Let's
  see how we can use golang to read TOML files.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go mod init &lt;your_project_package_name&gt;\n$
  go get github.com/pelletier/go-toml\n</pre></div>\n\n</pre>\n\n<p>The above commands
  are used for setting up a golang package or project and installing the <a href=\"https://pkg.go.dev/github.com/pelletier/go-toml\">go-toml</a>
  package. Once the above commands are done executing, it will generate <code>go.mod</code>
  and <code>go.sum</code> files used for storing dependencies and packages installed
  for the project locally.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>toml</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">[blog]</span><span class=\"w\"></span>\n<span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s\">&#39;techstructive-blog&#39;</span><span class=\"w\"></span>\n<span
  class=\"n\">tags</span><span class=\"o\">=</span><span class=\"p\">[</span><span
  class=\"s\">&#39;go&#39;</span><span class=\"p\">,</span><span class=\"s\">&#39;django&#39;</span><span
  class=\"p\">,</span><span class=\"s\">&#39;vim&#39;</span><span class=\"p\">]</span><span
  class=\"w\"></span>\n<span class=\"n\">author</span><span class=\"o\">=</span><span
  class=\"s\">&#39;meet gor&#39;</span><span class=\"w\"></span>\n<span class=\"n\">active</span><span
  class=\"o\">=</span><span class=\"kc\">true</span><span class=\"w\"></span>\n\n<span
  class=\"k\">[author]</span><span class=\"w\"></span>\n<span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s\">&#39;Meet Gor&#39;</span><span class=\"w\"></span>\n<span
  class=\"n\">github</span><span class=\"o\">=</span><span class=\"s\">&#39;mr-destructive&#39;</span><span
  class=\"w\"></span>\n<span class=\"n\">twitter</span><span class=\"o\">=</span><span
  class=\"s\">&#39;meetgor21&#39;</span><span class=\"w\"></span>\n<span class=\"n\">posts</span><span
  class=\"o\">=</span><span class=\"mi\">80</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  above is the sample file <code>blog.toml</code> which we will use to read in the
  go script below. The toml file has a similar structure as we have seen in the previous
  examples. We have different data types like string, boolean, integer, and list.</p>\n<pre
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
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">toml</span><span class=\"w\"> </span><span
  class=\"s\">&quot;github.com/pelletier/go-toml&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">f</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">ReadFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;blog.toml&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">data</span><span class=\"w\"> </span><span
  class=\"kd\">map</span><span class=\"p\">[</span><span class=\"kd\">interface</span><span
  class=\"p\">{}]</span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">toml</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">data</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">data</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">k</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">v</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">data</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">k</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">v</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t\t</span><span class=\"k\">switch</span><span
  class=\"w\"> </span><span class=\"nx\">t</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">v</span><span class=\"p\">.(</span><span
  class=\"kd\">type</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">case</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"kt\">string</span><span class=\"p\">]</span><span
  class=\"kd\">interface</span><span class=\"p\">{}:</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">a</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">t</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">a</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run toml.go\n\nmap[author:map[github:mr-destructive
  name:Meet Gor posts:80 twitter:meetgor21] blog:map[active:true author:meet gor\n\nname:techstructive-blog
  tags:[go django vim]]]\n\nblog : map[active:true author:meet gor name:techstructive-blog
  tags:[go django vim]]\nname : techstructive-blog\ntags : [go django vim]\nauthor
  : meet gor\nactive : true\n\nauthor : map[github:mr-destructive name:Meet Gor posts:80
  twitter:meetgor21]\n\nname : Meet Gor\ngithub : mr-destructive\ntwitter : meetgor21\nposts
  : 80\n</pre></div>\n\n</pre>\n\n<p>So, in the above example and output, the YAML
  file was read and the key-value pairs inside them were read. The first thing we
  do is read the file <code>blog.toml</code> with <code>ioutil</code> package, with
  the <code>ReadFile</code> function. The function takes in the string as the path
  to the file to be read and returns a slice of byte. We use this slice of byte as
  a parameter to the <a href=\"https://pkg.go.dev/github.com/pelletier/go-toml#Unmarshal\">Unmarshal</a>
  method. The second paramter for the <code>Unmarshal</code> is the output variable(usually
  a pointer to a variable), we have created a map of <code>interface{]</code> with
  an <code>interface</code> as we see there can be nested keys which hold the name
  of the config.</p>\n<p>The variable <code>data</code> is a map of <code>interface{}</code>
  to an <code>interface{}</code>, and we parse the memory address to the <code>data</code>
  variable to the <code>Unmarshal</code> method. Thereby the parsed <code>TOML</code>
  content is stored in the data variable.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>name : techstructive-blog\nstring\n\ntags
  : [go django vim]\n[]interface{}\n\nauthor : meet gor\nstring\n\nactive : true\nbool\n\nname
  : Meet Gor\nstring\n\ngithub : mr-destructive\nstring\n\ntwitter : meetgor21\nstring\n\nposts
  : 80\nint64\n</pre></div>\n\n</pre>\n\n<p>The above is a verbose output for the
  type of the values as parsed by golang, we have <code>string</code>, <code>bool</code>,
  <code>int64</code>, and a <code>slice</code> (list with interface{} attached with
  it). Only types like <code>string</code>, <code>bool</code>, <code>int</code>, <code>float64</code>
  can be parsed from the Unmarshal function, other than these types, the type will
  have an interface attached to it.</p>\n<p>In such cases, where the type of value
  is not among the 4 types(string, bool, int float), we can use a pre-defined struct
  to parse the content from the file. Though it would require a strict structure and
  predictable response from the parsed file.</p>\n<h3>Reading CSV file</h3>\n<p>We
  can even read a CSV file in golang, we have seen in the previous post, we have used
  custom delimiters in the parsing of the file.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>csv</div>\n<div class=\"highlight\"><pre><span></span>id,name,posts,exp\n21,jim,23,2\n33,kevin,39,1\n45,james,70,2\n56,chris,89,3\n</pre></div>\n\n</pre>\n\n<p>The
  above file is a sample csv file, though the size is too small, we can use it as
  an example.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;encoding/csv&quot;</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;temp.csv&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">check_error</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">reader</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">csv</span><span
  class=\"p\">.</span><span class=\"nx\">NewReader</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">n</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">reader</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">check_error</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">line</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">n</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">text</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">line</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">text</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run main.go\nid\nname\nposts\nexp\n21\njim\n23\n2\n33\nkevin\n39\n1\n45\njames\n70\n2\n56\nchris\n89\n3\n</pre></div>\n\n</pre>\n\n<p>The
  CSV package has a <a href=\"https://pkg.go.dev/encoding/csv#NewReader\">NewReader</a>
  method that accepts an <code>io.Reader</code> and returns a <code>Reader</code>
  object. After parsing the reader, we use the <a href=\"https://pkg.go.dev/encoding/csv#Reader.ReadAll\">ReadAll</a>
  method to return a 2d string or an error if there exists an error while parsing
  the content. You can get a detailed explanation of the CSV parsing and reading in
  the <a href=\"https://www.meetgor.com/golang-file-read/#Read%20File%20by%20a%20delimiter\">previous
  post</a>.</p>\n<h3>Reading CSV from URL</h3>\n<p>The CSV file can also be read from
  the URL, the content of the file is a <code>response.Body</code> in place of the
  file object reference, in the previous example, the <a href=\"https://pkg.go.dev/os#Open\">os.Open()</a>
  method returns a <a href=\"https://pkg.go.dev/os#File\">os.File</a> object.</p>\n<p>We
  use the <code>http.Get(string)</code> method to get the response from the URL for
  reading the CSV file present on the web.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">url</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://github.com/woocommerce/woocommerce/raw/master/sample-data/sample_products.csv&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">response</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">url</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">return</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">response</span><span class=\"p\">.</span><span
  class=\"nx\">Body</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">reader</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">csv</span><span class=\"p\">.</span><span
  class=\"nx\">NewReader</span><span class=\"p\">(</span><span class=\"nx\">response</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">n</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">reader</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">line</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">n</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">text</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">line</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">text</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run csv.go\n&lt;feff&gt;ID\nType\nSKU\nName\nPublished\nIs
  featured?\nVisibility in catalog\nShort description\nDescription\nDate sale price
  starts\nDate sale price ends\n...\n...\n...\n</pre></div>\n\n</pre>\n\n<p>So, that's
  how we can read a CSV file from the URL. By fetching the CSV URL <code>https://github.com/woocommerce/woocommerce/raw/master/sample-data/sample_products.csv</code>
  from the <a href=\"https://pkg.go.dev/net/http#Client.Get\">http.Get</a> method,
  this will get us the <a href=\"https://pkg.go.dev/net/http#Response\">response.Body</a>
  that contains the actual CSV file content. The response than can be parsed to the
  <code>csv.NewReader(*Os.File).ReadAll()</code> i.e. <a href=\"https://pkg.go.dev/encoding/csv#Reader.ReadAll\">reader.ReadAll()</a>.
  The function returns a multidimensional slice <code>[][]slice</code> that can be
  iterated and parsed as per requirement.</p>\n<h3>Reading XML file</h3>\n<p>XML is
  the de facto standard for RSS feeds, it is widely used in many places and is still
  all over the web. We'll see an example to read an XML file locally, but as we saw
  in the above example, we can also read an RSS link from the web.</p>\n<p>Just like
  CSV, we have <a href=\"https://pkg.go.dev/encoding/xml\">encoding/xml</a>, and the
  standard library has all the functions used for parsing the XML files.</p>\n<p>We
  will be using a local XML file called <code>rss.xml</code>, and reading the contents
  from the tags in the file.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>xml</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cp\">&lt;?xml version=&quot;1.0&quot; encoding=&quot;UTF-8&quot; ?&gt;</span>\n<span
  class=\"nt\">&lt;channel&gt;</span>\n<span class=\"nt\">&lt;title&gt;</span>Meet
  Gor<span class=\"nt\">&lt;/title&gt;</span>\n<span class=\"nt\">&lt;description&gt;</span>Techstructive
  Blog Feed<span class=\"nt\">&lt;/description&gt;</span>\n<span class=\"nt\">&lt;item&gt;</span>\n<span
  class=\"nt\">&lt;title&gt;</span>Why and How to make and use Vim as a text editor
  and customizable IDE<span class=\"nt\">&lt;/title&gt;</span>\n<span class=\"nt\">&lt;link&gt;</span>https://www.meetgor.com/vim-text-editor-ide<span
  class=\"nt\">&lt;/link&gt;</span>\n<span class=\"nt\">&lt;/item&gt;</span>\n<span
  class=\"nt\">&lt;item&gt;</span>\n<span class=\"nt\">&lt;title&gt;</span>Setting
  up Vim for Python<span class=\"nt\">&lt;/title&gt;</span>\n<span class=\"nt\">&lt;link&gt;</span>https://www.meetgor.com/vim-for-python<span
  class=\"nt\">&lt;/link&gt;</span>\n<span class=\"nt\">&lt;/item&gt;</span>\n<span
  class=\"nt\">&lt;/channel&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>The above example
  is a short part of my blog's <a href=\"https://www.meetgor.com/rss\">rss</a> feed.
  I have just trimmed the unwanted part and will be just using the tags that we want
  to fetch.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"s\">&quot;encoding/xml&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Channel</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">XMLName</span><span class=\"w\">     </span><span class=\"nx\">xml</span><span
  class=\"p\">.</span><span class=\"nx\">Name</span><span class=\"w\"> </span><span
  class=\"s\">`xml:&quot;channel&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Title</span><span class=\"w\">       </span><span class=\"kt\">string</span><span
  class=\"w\">   </span><span class=\"s\">`xml:&quot;title&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Description</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\">   </span><span class=\"s\">`xml:&quot;description&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Item</span><span
  class=\"w\">        </span><span class=\"p\">[]</span><span class=\"nx\">Item</span><span
  class=\"w\">   </span><span class=\"s\">`xml:&quot;item&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Item</span><span class=\"w\"> </span><span
  class=\"kd\">struct</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">XMLName</span><span
  class=\"w\"> </span><span class=\"nx\">xml</span><span class=\"p\">.</span><span
  class=\"nx\">Name</span><span class=\"w\"> </span><span class=\"s\">`xml:&quot;item&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Title</span><span
  class=\"w\">   </span><span class=\"kt\">string</span><span class=\"w\">   </span><span
  class=\"s\">`xml:&quot;title&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Link</span><span class=\"w\">    </span><span class=\"kt\">string</span><span
  class=\"w\">   </span><span class=\"s\">`xml:&quot;link&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">check_error</span><span class=\"p\">(</span><span
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
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">f</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">ReadFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;rss.xml&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">check_error</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">d</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">Channel</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">xml</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">d</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">check_error</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">item</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">d</span><span class=\"p\">.</span><span class=\"nx\">Item</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">item</span><span class=\"p\">.</span><span
  class=\"nx\">Title</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run xml.go\n\n{{ channel}
  Meet Gor Techstructive Blog Feed [{{ item} Why and How to make and use Vim as a
  text editor and customizable IDE https://www.meetgor.com/vim-text-editor-ide} {{
  item} Setting up Vim for Python https://www.meetgor.com/vim-for-python}]}\n\nWhy
  and How to make and use Vim as a text editor and customizable IDE\nSetting up Vim
  for Python\n</pre></div>\n\n</pre>\n\n<p>The above example uses a couple of <code>struct</code>
  like <code>Channel</code> and <code>Item</code> that stores the tag data like <code>title</code>,
  <code>description</code>, <code>link</code>, etc. Unlike the JSON, YAML, and toml
  files, XML can't parse the content directly we need a structure to parse into. And
  that's fine as XML is not much dynamic in terms of config, we usually have standard
  tags and elements which can be pre-defined in a struct type.</p>\n<p>In this example,
  the RSS feed has a <code>Channel</code> tag with <code>title</code>, <code>description</code>,
  and <code>item</code>.</p>\n<p><strong>NOTE: Use Title case for the fields of the
  structs. It will make the fields public, I spent a few hours debugging that really
  :)</strong></p>\n<p>So, we define the <code>Channel</code> struct with fields like
  <code>Title</code> as a string which is a tag in the file as <code>xml:&quot;title&quot;</code>.
  This means the title in the tag of the XML will be stored in the field as a string
  in the attribute name as <code>Title</code>. Similarly, we have fields like <code>Description</code>
  and <code>Item[]</code> this is a list or multiple of <code>item</code> tags that
  might be present in the XML file. The <code>XMLName</code> is used for identifying
  the parent tag for the struct, so we use <code>channel</code> for the first struct
  as it is the first tag appearing of the hierarchy in the XML file.</p>\n<p>So, we
  create an object for the root structure as <code>Channel{}</code> (an empty object
  instantiated). The <code>xml.Unmarshal</code> function is parsed with the content
  of the file as <code>data</code> which is a slice of byte as we have seen in the
  previous examples. The slice of byte is then used in the <code>Unmarshal</code>
  method as the first parameter and the reference of an empty <code>Channel</code>
  object as the second parameter. The second parameter will be to store the parsed
  XML content from the file.</p>\n<p>I have a few examples on the GitHub repository
  covering the reading of files from a URL for the CSV, and XML files. But, this concept
  in the example, can be applied to JSON, YAML, and other file formats as well.</p>\n<p>That's
  it from this part. Reference for all the code examples and commands can be found
  in the <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/tree/main/scripts/files/read/config_files\">100
  days of Golang</a> GitHub repository.</p>\n<h2>Conclusion</h2>\n<p>So, that's it
  from this post, we covered how to read specific configuration files like <code>JSON</code>,
  <code>CSV</code>, <code>YAML</code>, <code>TOML</code>, and <code>XML</code>. We
  saw how to read a local file and also touched on the concept to read contents from
  a file on the web with a URL. We also saw how we can use pre-defined structs to
  parse content from a file, especially for XML.</p>\n<p>Thank you for reading. If
  you have any queries, questions, or feedback, you can let me know in the discussion
  below or on my social handles. Happy Coding :)</p>\n"
cover: ''
date: 2022-11-01
datetime: 2022-11-01 00:00:00+00:00
description: Reading specific files used generally for configuration and storing data,
  as well as for web development. Reading file formats like JSON, YAML, TOML, CSV,
  and XML.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2022-11-01-GO-23-Config-File-Read.md
html: "<h2>Reading specific file types (JSON, YAML, TOML)</h2>\n<p>In the previous
  post, we have seen how to read files in golang, in this extended post of that part,
  we will look into reading some specific files used for configuration and storing
  data like JSON, YAML, TOML, CSV, etc.</p>\n<p>We will see how to read files and
  get individual pieces in the files. We'll use packages like <code>os</code>, <code>ioutil</code>
  and <code>encoding</code> to perform reading operations on file and file objects.</p>\n<h3>Reading
  a JSON File</h3>\n<p>Golang has built-in support for reading JSON files, but still,
  we can and need to have low-level controls on how to parse and extract content from
  the file.</p>\n<p>Let's say we have a <code>json</code> file named <code>blog.json</code>,
  we can use the <a href=\"https://pkg.go.dev/encoding/json\">encoding/json</a> package
  to convert the JSON data into a GO object (an object that is native and understandable
  to go). The <a href=\"https://pkg.go.dev/encoding/json#Unmarshal\">Unmarshal</a>
  function is used to convert the slice of bytes from the file, into a map object.</p>\n<pre
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
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;title&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;Golang Blog Series&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;date&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;22nd October
  2022&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nt\">&quot;tags&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"p\">[</span><span class=\"s2\">&quot;go&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s2\">&quot;files&quot;</span><span
  class=\"p\">],</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;words&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"mi\">1500</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nt\">&quot;published&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"kc\">true</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  above is a simple JSON file, this file has a few types of key-value pairs like string,
  list, integer, and boolean. But we can also have nested objects and a list of those
  nested objects.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">f</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">ReadFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;blog.json&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">data</span><span class=\"w\"> </span><span
  class=\"kd\">map</span><span class=\"p\">[</span><span class=\"kt\">string</span><span
  class=\"p\">]</span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">([]</span><span
  class=\"nb\">byte</span><span class=\"p\">(</span><span class=\"nx\">f</span><span
  class=\"p\">),</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">data</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">data</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">k</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">v</span><span class=\"w\">
  </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">data</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">k</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">v</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p><strong>I
  have removed the time stamp from the logs below so as to clearly see the output.
  We can use <code>fmt</code> to print the simple things while keeping consistent
  with the rest of the snippets in the series.</strong></p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run json.go\n\nmap[date:22nd
  October 2022 published:true tags:[go files] title:Golang Blog Series words:1500]\npublished
  : true\ntitle : Golang Blog Series\ndate : 22nd October 2022\ntags : [go files]\nwords
  : 1500\n</pre></div>\n\n</pre>\n\n<p>The file is read using the <a href=\"https://pkg.go.dev/os#ReadFile\">os.ReadFile</a>
  method, that takes in a string as a path to the file and returns a slice of bytes
  or an error if there was an issue in reading the file. The parsed slice of byte
  is than passed as the first argument to the <code>Unmarshal</code> method in the
  <code>encoding/json</code> package. The second parameter is the output reference
  where the parsed JSON will be stored or returned. The function does not return the
  parsed content instead returns an error if there arose any while parsing the JSON
  content.</p>\n<p>As we can see we have got a map of <code>string</code> with an
  <code>interface</code>. The interface is used because the value of the key can be
  anything. There is no fixed value like a <code>string</code>, <code>int</code>,
  <code>bool</code>, or a nested <code>map</code>, <code>slice</code>. Hence we have
  mapped the JSON object as a map of <code>string</code> with an <code>interface</code>.
  The type of the value is identified with the interface it has attached to it. Let's
  take a look what is the type of each value in the map.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>published : true\nbool\n\ntitle
  : Golang Blog Series\nstring\n\ndate : 22nd October 2022\nstring\n\ntags : [go files]\n[]interface
  {}\n\nwords : 1500\nfloat64\n</pre></div>\n\n</pre>\n\n<p>Here, we can see it has
  correctly identified the string type of the fields like bool in case of true or
  false, a string for string type of values, the fourth field however has a list interface
  attached to it. The default precedence for <code>float64</code> over integer is
  the reason the <code>1500</code> value is of type <code>float64</code>.</p>\n<h3>Reading
  a YAML File</h3>\n<p>Though there is no standard package for parsing/unmarshaling
  <code>YAML</code> files in golang, it's quite easy to use a third-party package
  and use it to read YAML files.</p>\n<p>The package <a href=\"https://pkg.go.dev/gopkg.in/yaml.v3\">gopkg.in/yaml.v3</a>
  is used for encoding and decoding YAML files. We'll be just using it for decoding
  a YAML file by reading it and converting the file object into native Go objects
  like maps, lists, strings, etc.</p>\n<p>The below steps can be used for setting
  up the package and installing the YAML package locally.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go mod init &lt;your_project_package_name&gt;\ngo
  get gopkg.in/yaml.v3\n</pre></div>\n\n</pre>\n\n<p>This should create two files
  namely <code>go.mod</code> and <code>go.sum</code> with the dependency of the <code>gopkg.in/yaml.v3</code>
  package.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>yml</div>\n<div class=\"highlight\"><pre><span></span>title:
  &quot;Golang Blog Series&quot;\ndate: &quot;22nd October 2022&quot;\ntags: [&quot;go&quot;,
  &quot;files&quot;]\npublished: false\nwords: 1500\n</pre></div>\n\n</pre>\n\n<p>The
  above file is a simple YAML config, we'll follow similar kind of examples for the
  dummy files used in the examples.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">yaml</span><span class=\"w\"> </span><span
  class=\"s\">&quot;gopkg.in/yaml.v3&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">f</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">ReadFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;blog.yaml&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">data</span><span class=\"w\"> </span><span
  class=\"kd\">map</span><span class=\"p\">[</span><span class=\"kt\">string</span><span
  class=\"p\">]</span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">yaml</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">data</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">data</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">k</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">v</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">data</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">k</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">v</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run yaml.go\n\nmap[date:22nd
  October 2022 published:false tags:[go files] title:Golang Blog Series words:1500]\npublished
  : false\nwords : 1500\ntitle : Golang Blog Series\ndate : 22nd October 2022\ntags
  : [go files]\n</pre></div>\n\n</pre>\n\n<p>The above code and output demonstrate
  the usage of the <code>yaml.v3</code> package for reading a YAML file.</p>\n<p>Firstly,
  we read the file into a single-string object with the <a href=\"https://pkg.go.dev/os#ReadFile\">os.ReadFile()</a>
  method. The method will return a <code>[]byte</code> (slice of byte) or an error.
  If there is an error, we simply log or panic out of the program, else we can use
  the <a href=\"https://pkg.go.dev/gopkg.in/yaml.v3#Unmarshal\">yaml.Unmarshal</a>
  method to convert the string/slice of the byte into a map or a pre-defined struct.
  In this example, we just keep things simple by storing the file content as <code>map
  [string, interface{}]</code>, i.e. a map of <code>string</code> and an <code>interface</code>.
  The key for YAML can be only a string or an integer. It can't have unrestricted
  data types like the value can have. Though if you want to be unrestrictive, you
  can use a map of <code>map[interface{}]interface{}</code> to make the key any shape
  you like to have.</p>\n<p>So, we have created a variable called <code>data</code>
  as a map of <code>string</code> and <code>interface{}</code>, basically key can
  be a string and the value can be any type of interface depending on the parsed literally
  from the file object. The <code>Unmarshal</code> function takes in two parameters,
  the first being the slice of byte i.e. the file contents, and the second being the
  output variable. Now, the method does not return the parsed YAML, it can return
  an error if there arose, so we need to set the second parameter as a pointer to
  the object into which we want to store the parsed YAML.</p>\n<p>In the example,
  we have called <code>Unmarshal(f, &amp;data)</code> which will fetch the contents
  from the slice of bytes <code>f</code> and output the parsed YAML from the slice
  of bytes into the memory location of <code>data</code> and hence using <code>&amp;data</code>
  indicating the pointer to the variable(fetch the memory address).</p>\n<p>So, that
  is how we obtain the map of keys and values from the YAML config, thereafter, you
  can iterate on the map, access the keys and values, type caste them as per requirement,
  and basically have control over what processing needs to be done to the parsed YAML
  content.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n\n<div class=\"highlight\"><pre><span></span>published : false\nbool\n\nwords
  : 1500\nint\n\ntitle : Golang Blog Series\nstring\n\ndate : 22nd October 2022\nstring\n\ntags
  : [go files]\n[]interface {}\n</pre></div>\n\n</pre>\n\n<p>I have just printed the
  types of the values in the above output as <code>log.Printf(&quot;%T&quot;, v)</code>,
  we can see the types are being correctly recognized and being parsed. The last object
  is indeed a <code>slice</code> so it has an interface of the slice(array) attached
  to it.</p>\n<h3>Reading a TOML file</h3>\n<p>YAML and TOML are almost identical,
  though toml has more restrictive configuration options, but is more readable than
  YAML, as YAML can get complicated pretty quickly. Though both of them have their
  pros and cons, YAML is used everywhere in the DevOps world, configs, whereas TOML
  is the format of choice for python packaging, and static site generation configs.</p>\n<p>Let's
  see how we can use golang to read TOML files.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go mod init &lt;your_project_package_name&gt;\n$
  go get github.com/pelletier/go-toml\n</pre></div>\n\n</pre>\n\n<p>The above commands
  are used for setting up a golang package or project and installing the <a href=\"https://pkg.go.dev/github.com/pelletier/go-toml\">go-toml</a>
  package. Once the above commands are done executing, it will generate <code>go.mod</code>
  and <code>go.sum</code> files used for storing dependencies and packages installed
  for the project locally.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>toml</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">[blog]</span><span class=\"w\"></span>\n<span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s\">&#39;techstructive-blog&#39;</span><span class=\"w\"></span>\n<span
  class=\"n\">tags</span><span class=\"o\">=</span><span class=\"p\">[</span><span
  class=\"s\">&#39;go&#39;</span><span class=\"p\">,</span><span class=\"s\">&#39;django&#39;</span><span
  class=\"p\">,</span><span class=\"s\">&#39;vim&#39;</span><span class=\"p\">]</span><span
  class=\"w\"></span>\n<span class=\"n\">author</span><span class=\"o\">=</span><span
  class=\"s\">&#39;meet gor&#39;</span><span class=\"w\"></span>\n<span class=\"n\">active</span><span
  class=\"o\">=</span><span class=\"kc\">true</span><span class=\"w\"></span>\n\n<span
  class=\"k\">[author]</span><span class=\"w\"></span>\n<span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s\">&#39;Meet Gor&#39;</span><span class=\"w\"></span>\n<span
  class=\"n\">github</span><span class=\"o\">=</span><span class=\"s\">&#39;mr-destructive&#39;</span><span
  class=\"w\"></span>\n<span class=\"n\">twitter</span><span class=\"o\">=</span><span
  class=\"s\">&#39;meetgor21&#39;</span><span class=\"w\"></span>\n<span class=\"n\">posts</span><span
  class=\"o\">=</span><span class=\"mi\">80</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  above is the sample file <code>blog.toml</code> which we will use to read in the
  go script below. The toml file has a similar structure as we have seen in the previous
  examples. We have different data types like string, boolean, integer, and list.</p>\n<pre
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
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">toml</span><span class=\"w\"> </span><span
  class=\"s\">&quot;github.com/pelletier/go-toml&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">f</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">ReadFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;blog.toml&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">data</span><span class=\"w\"> </span><span
  class=\"kd\">map</span><span class=\"p\">[</span><span class=\"kd\">interface</span><span
  class=\"p\">{}]</span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">toml</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">data</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">log</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">data</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">k</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">v</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">data</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">k</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">v</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t\t</span><span class=\"k\">switch</span><span
  class=\"w\"> </span><span class=\"nx\">t</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">v</span><span class=\"p\">.(</span><span
  class=\"kd\">type</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">case</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"kt\">string</span><span class=\"p\">]</span><span
  class=\"kd\">interface</span><span class=\"p\">{}:</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">a</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">t</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t\t</span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">a</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run toml.go\n\nmap[author:map[github:mr-destructive
  name:Meet Gor posts:80 twitter:meetgor21] blog:map[active:true author:meet gor\n\nname:techstructive-blog
  tags:[go django vim]]]\n\nblog : map[active:true author:meet gor name:techstructive-blog
  tags:[go django vim]]\nname : techstructive-blog\ntags : [go django vim]\nauthor
  : meet gor\nactive : true\n\nauthor : map[github:mr-destructive name:Meet Gor posts:80
  twitter:meetgor21]\n\nname : Meet Gor\ngithub : mr-destructive\ntwitter : meetgor21\nposts
  : 80\n</pre></div>\n\n</pre>\n\n<p>So, in the above example and output, the YAML
  file was read and the key-value pairs inside them were read. The first thing we
  do is read the file <code>blog.toml</code> with <code>ioutil</code> package, with
  the <code>ReadFile</code> function. The function takes in the string as the path
  to the file to be read and returns a slice of byte. We use this slice of byte as
  a parameter to the <a href=\"https://pkg.go.dev/github.com/pelletier/go-toml#Unmarshal\">Unmarshal</a>
  method. The second paramter for the <code>Unmarshal</code> is the output variable(usually
  a pointer to a variable), we have created a map of <code>interface{]</code> with
  an <code>interface</code> as we see there can be nested keys which hold the name
  of the config.</p>\n<p>The variable <code>data</code> is a map of <code>interface{}</code>
  to an <code>interface{}</code>, and we parse the memory address to the <code>data</code>
  variable to the <code>Unmarshal</code> method. Thereby the parsed <code>TOML</code>
  content is stored in the data variable.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>name : techstructive-blog\nstring\n\ntags
  : [go django vim]\n[]interface{}\n\nauthor : meet gor\nstring\n\nactive : true\nbool\n\nname
  : Meet Gor\nstring\n\ngithub : mr-destructive\nstring\n\ntwitter : meetgor21\nstring\n\nposts
  : 80\nint64\n</pre></div>\n\n</pre>\n\n<p>The above is a verbose output for the
  type of the values as parsed by golang, we have <code>string</code>, <code>bool</code>,
  <code>int64</code>, and a <code>slice</code> (list with interface{} attached with
  it). Only types like <code>string</code>, <code>bool</code>, <code>int</code>, <code>float64</code>
  can be parsed from the Unmarshal function, other than these types, the type will
  have an interface attached to it.</p>\n<p>In such cases, where the type of value
  is not among the 4 types(string, bool, int float), we can use a pre-defined struct
  to parse the content from the file. Though it would require a strict structure and
  predictable response from the parsed file.</p>\n<h3>Reading CSV file</h3>\n<p>We
  can even read a CSV file in golang, we have seen in the previous post, we have used
  custom delimiters in the parsing of the file.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>csv</div>\n<div class=\"highlight\"><pre><span></span>id,name,posts,exp\n21,jim,23,2\n33,kevin,39,1\n45,james,70,2\n56,chris,89,3\n</pre></div>\n\n</pre>\n\n<p>The
  above file is a sample csv file, though the size is too small, we can use it as
  an example.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;encoding/csv&quot;</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"s\">&quot;temp.csv&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">check_error</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">reader</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">csv</span><span
  class=\"p\">.</span><span class=\"nx\">NewReader</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">n</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">reader</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">check_error</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">line</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">n</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">text</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">line</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">text</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run main.go\nid\nname\nposts\nexp\n21\njim\n23\n2\n33\nkevin\n39\n1\n45\njames\n70\n2\n56\nchris\n89\n3\n</pre></div>\n\n</pre>\n\n<p>The
  CSV package has a <a href=\"https://pkg.go.dev/encoding/csv#NewReader\">NewReader</a>
  method that accepts an <code>io.Reader</code> and returns a <code>Reader</code>
  object. After parsing the reader, we use the <a href=\"https://pkg.go.dev/encoding/csv#Reader.ReadAll\">ReadAll</a>
  method to return a 2d string or an error if there exists an error while parsing
  the content. You can get a detailed explanation of the CSV parsing and reading in
  the <a href=\"https://www.meetgor.com/golang-file-read/#Read%20File%20by%20a%20delimiter\">previous
  post</a>.</p>\n<h3>Reading CSV from URL</h3>\n<p>The CSV file can also be read from
  the URL, the content of the file is a <code>response.Body</code> in place of the
  file object reference, in the previous example, the <a href=\"https://pkg.go.dev/os#Open\">os.Open()</a>
  method returns a <a href=\"https://pkg.go.dev/os#File\">os.File</a> object.</p>\n<p>We
  use the <code>http.Get(string)</code> method to get the response from the URL for
  reading the CSV file present on the web.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">url</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://github.com/woocommerce/woocommerce/raw/master/sample-data/sample_products.csv&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">response</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">url</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">return</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">response</span><span class=\"p\">.</span><span
  class=\"nx\">Body</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">reader</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">csv</span><span class=\"p\">.</span><span
  class=\"nx\">NewReader</span><span class=\"p\">(</span><span class=\"nx\">response</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">n</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">reader</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">line</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">n</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">text</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">line</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">text</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run csv.go\n&lt;feff&gt;ID\nType\nSKU\nName\nPublished\nIs
  featured?\nVisibility in catalog\nShort description\nDescription\nDate sale price
  starts\nDate sale price ends\n...\n...\n...\n</pre></div>\n\n</pre>\n\n<p>So, that's
  how we can read a CSV file from the URL. By fetching the CSV URL <code>https://github.com/woocommerce/woocommerce/raw/master/sample-data/sample_products.csv</code>
  from the <a href=\"https://pkg.go.dev/net/http#Client.Get\">http.Get</a> method,
  this will get us the <a href=\"https://pkg.go.dev/net/http#Response\">response.Body</a>
  that contains the actual CSV file content. The response than can be parsed to the
  <code>csv.NewReader(*Os.File).ReadAll()</code> i.e. <a href=\"https://pkg.go.dev/encoding/csv#Reader.ReadAll\">reader.ReadAll()</a>.
  The function returns a multidimensional slice <code>[][]slice</code> that can be
  iterated and parsed as per requirement.</p>\n<h3>Reading XML file</h3>\n<p>XML is
  the de facto standard for RSS feeds, it is widely used in many places and is still
  all over the web. We'll see an example to read an XML file locally, but as we saw
  in the above example, we can also read an RSS link from the web.</p>\n<p>Just like
  CSV, we have <a href=\"https://pkg.go.dev/encoding/xml\">encoding/xml</a>, and the
  standard library has all the functions used for parsing the XML files.</p>\n<p>We
  will be using a local XML file called <code>rss.xml</code>, and reading the contents
  from the tags in the file.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>xml</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cp\">&lt;?xml version=&quot;1.0&quot; encoding=&quot;UTF-8&quot; ?&gt;</span>\n<span
  class=\"nt\">&lt;channel&gt;</span>\n<span class=\"nt\">&lt;title&gt;</span>Meet
  Gor<span class=\"nt\">&lt;/title&gt;</span>\n<span class=\"nt\">&lt;description&gt;</span>Techstructive
  Blog Feed<span class=\"nt\">&lt;/description&gt;</span>\n<span class=\"nt\">&lt;item&gt;</span>\n<span
  class=\"nt\">&lt;title&gt;</span>Why and How to make and use Vim as a text editor
  and customizable IDE<span class=\"nt\">&lt;/title&gt;</span>\n<span class=\"nt\">&lt;link&gt;</span>https://www.meetgor.com/vim-text-editor-ide<span
  class=\"nt\">&lt;/link&gt;</span>\n<span class=\"nt\">&lt;/item&gt;</span>\n<span
  class=\"nt\">&lt;item&gt;</span>\n<span class=\"nt\">&lt;title&gt;</span>Setting
  up Vim for Python<span class=\"nt\">&lt;/title&gt;</span>\n<span class=\"nt\">&lt;link&gt;</span>https://www.meetgor.com/vim-for-python<span
  class=\"nt\">&lt;/link&gt;</span>\n<span class=\"nt\">&lt;/item&gt;</span>\n<span
  class=\"nt\">&lt;/channel&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>The above example
  is a short part of my blog's <a href=\"https://www.meetgor.com/rss\">rss</a> feed.
  I have just trimmed the unwanted part and will be just using the tags that we want
  to fetch.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"s\">&quot;encoding/xml&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Channel</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">XMLName</span><span class=\"w\">     </span><span class=\"nx\">xml</span><span
  class=\"p\">.</span><span class=\"nx\">Name</span><span class=\"w\"> </span><span
  class=\"s\">`xml:&quot;channel&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Title</span><span class=\"w\">       </span><span class=\"kt\">string</span><span
  class=\"w\">   </span><span class=\"s\">`xml:&quot;title&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Description</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\">   </span><span class=\"s\">`xml:&quot;description&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Item</span><span
  class=\"w\">        </span><span class=\"p\">[]</span><span class=\"nx\">Item</span><span
  class=\"w\">   </span><span class=\"s\">`xml:&quot;item&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Item</span><span class=\"w\"> </span><span
  class=\"kd\">struct</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">XMLName</span><span
  class=\"w\"> </span><span class=\"nx\">xml</span><span class=\"p\">.</span><span
  class=\"nx\">Name</span><span class=\"w\"> </span><span class=\"s\">`xml:&quot;item&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Title</span><span
  class=\"w\">   </span><span class=\"kt\">string</span><span class=\"w\">   </span><span
  class=\"s\">`xml:&quot;title&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Link</span><span class=\"w\">    </span><span class=\"kt\">string</span><span
  class=\"w\">   </span><span class=\"s\">`xml:&quot;link&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">check_error</span><span class=\"p\">(</span><span
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
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">f</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">ReadFile</span><span class=\"p\">(</span><span
  class=\"s\">&quot;rss.xml&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">check_error</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">f</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">d</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">Channel</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">xml</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">f</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">d</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">check_error</span><span class=\"p\">(</span><span
  class=\"nx\">err</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">item</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">d</span><span class=\"p\">.</span><span class=\"nx\">Item</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">item</span><span class=\"p\">.</span><span
  class=\"nx\">Title</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run xml.go\n\n{{ channel}
  Meet Gor Techstructive Blog Feed [{{ item} Why and How to make and use Vim as a
  text editor and customizable IDE https://www.meetgor.com/vim-text-editor-ide} {{
  item} Setting up Vim for Python https://www.meetgor.com/vim-for-python}]}\n\nWhy
  and How to make and use Vim as a text editor and customizable IDE\nSetting up Vim
  for Python\n</pre></div>\n\n</pre>\n\n<p>The above example uses a couple of <code>struct</code>
  like <code>Channel</code> and <code>Item</code> that stores the tag data like <code>title</code>,
  <code>description</code>, <code>link</code>, etc. Unlike the JSON, YAML, and toml
  files, XML can't parse the content directly we need a structure to parse into. And
  that's fine as XML is not much dynamic in terms of config, we usually have standard
  tags and elements which can be pre-defined in a struct type.</p>\n<p>In this example,
  the RSS feed has a <code>Channel</code> tag with <code>title</code>, <code>description</code>,
  and <code>item</code>.</p>\n<p><strong>NOTE: Use Title case for the fields of the
  structs. It will make the fields public, I spent a few hours debugging that really
  :)</strong></p>\n<p>So, we define the <code>Channel</code> struct with fields like
  <code>Title</code> as a string which is a tag in the file as <code>xml:&quot;title&quot;</code>.
  This means the title in the tag of the XML will be stored in the field as a string
  in the attribute name as <code>Title</code>. Similarly, we have fields like <code>Description</code>
  and <code>Item[]</code> this is a list or multiple of <code>item</code> tags that
  might be present in the XML file. The <code>XMLName</code> is used for identifying
  the parent tag for the struct, so we use <code>channel</code> for the first struct
  as it is the first tag appearing of the hierarchy in the XML file.</p>\n<p>So, we
  create an object for the root structure as <code>Channel{}</code> (an empty object
  instantiated). The <code>xml.Unmarshal</code> function is parsed with the content
  of the file as <code>data</code> which is a slice of byte as we have seen in the
  previous examples. The slice of byte is then used in the <code>Unmarshal</code>
  method as the first parameter and the reference of an empty <code>Channel</code>
  object as the second parameter. The second parameter will be to store the parsed
  XML content from the file.</p>\n<p>I have a few examples on the GitHub repository
  covering the reading of files from a URL for the CSV, and XML files. But, this concept
  in the example, can be applied to JSON, YAML, and other file formats as well.</p>\n<p>That's
  it from this part. Reference for all the code examples and commands can be found
  in the <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/tree/main/scripts/files/read/config_files\">100
  days of Golang</a> GitHub repository.</p>\n<h2>Conclusion</h2>\n<p>So, that's it
  from this post, we covered how to read specific configuration files like <code>JSON</code>,
  <code>CSV</code>, <code>YAML</code>, <code>TOML</code>, and <code>XML</code>. We
  saw how to read a local file and also touched on the concept to read contents from
  a file on the web with a URL. We also saw how we can use pre-defined structs to
  parse content from a file, especially for XML.</p>\n<p>Thank you for reading. If
  you have any queries, questions, or feedback, you can let me know in the discussion
  below or on my social handles. Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/golang-023-config-files.png
long_description: In the previous post, we have seen how to read files in golang,
  in this extended post of that part, we will look into reading some specific files
  used for configuration and storing data like JSON, YAML, TOML, CSV, etc. We will
  see how to read files a
now: 2025-05-01 18:11:33.312532
path: blog/posts/2022-11-01-GO-23-Config-File-Read.md
prevnext: null
series:
- 100-days-of-golang
slug: golang-config-file-read
status: published
tags:
- go
templateKey: blog-post
title: 'Golang: JSON YAML TOML (config) File Reading.'
today: 2025-05-01
---

## Reading specific file types (JSON, YAML, TOML)

In the previous post, we have seen how to read files in golang, in this extended post of that part, we will look into reading some specific files used for configuration and storing data like JSON, YAML, TOML, CSV, etc.

We will see how to read files and get individual pieces in the files. We'll use packages like `os`, `ioutil` and `encoding` to perform reading operations on file and file objects.

### Reading a JSON File

Golang has built-in support for reading JSON files, but still, we can and need to have low-level controls on how to parse and extract content from the file.

Let's say we have a `json` file named `blog.json`, we can use the [encoding/json](https://pkg.go.dev/encoding/json) package to convert the JSON data into a GO object (an object that is native and understandable to go). The [Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) function is used to convert the slice of bytes from the file, into a map object.


```json
{
    "title": "Golang Blog Series",
    "date": "22nd October 2022",
    "tags": ["go", "files"],
    "words": 1500,
    "published": true
}
```

The above is a simple JSON file, this file has a few types of key-value pairs like string, list, integer, and boolean. But we can also have nested objects and a list of those nested objects.

```go
package main

import (
	"encoding/json"
	"log"
    "os"
)

func main() {

	f, err := os.ReadFile("blog.json")
	if err != nil {
		log.Println(err)
	}
	var data map[string]interface{}
	json.Unmarshal([]byte(f), &data)

	log.Println(data)
	for k, v := range data {
		log.Println(k, ":", v)
	}

}
```

**I have removed the time stamp from the logs below so as to clearly see the output. We can use `fmt` to print the simple things while keeping consistent with the rest of the snippets in the series.**

```
$ go run json.go

map[date:22nd October 2022 published:true tags:[go files] title:Golang Blog Series words:1500]
published : true
title : Golang Blog Series
date : 22nd October 2022
tags : [go files]
words : 1500
```
The file is read using the [os.ReadFile](https://pkg.go.dev/os#ReadFile) method, that takes in a string as a path to the file and returns a slice of bytes or an error if there was an issue in reading the file. The parsed slice of byte is than passed as the first argument to the `Unmarshal` method in the `encoding/json` package. The second parameter is the output reference where the parsed JSON will be stored or returned. The function does not return the parsed content instead returns an error if there arose any while parsing the JSON content.

As we can see we have got a map of `string` with an `interface`. The interface is used because the value of the key can be anything. There is no fixed value like a `string`, `int`, `bool`, or a nested `map`, `slice`. Hence we have mapped the JSON object as a map of `string` with an `interface`. The type of the value is identified with the interface it has attached to it. Let's take a look what is the type of each value in the map.

```
published : true
bool

title : Golang Blog Series
string

date : 22nd October 2022
string

tags : [go files]
[]interface {}

words : 1500
float64
```

Here, we can see it has correctly identified the string type of the fields like bool in case of true or false, a string for string type of values, the fourth field however has a list interface attached to it. The default precedence for `float64` over integer is the reason the `1500` value is of type `float64`. 


### Reading a YAML File

Though there is no standard package for parsing/unmarshaling `YAML` files in golang, it's quite easy to use a third-party package and use it to read YAML files.

The package [gopkg.in/yaml.v3](https://pkg.go.dev/gopkg.in/yaml.v3) is used for encoding and decoding YAML files. We'll be just using it for decoding a YAML file by reading it and converting the file object into native Go objects like maps, lists, strings, etc.

The below steps can be used for setting up the package and installing the YAML package locally.

```
go mod init <your_project_package_name>
go get gopkg.in/yaml.v3
```

This should create two files namely `go.mod` and `go.sum` with the dependency of the `gopkg.in/yaml.v3` package.

```yml
title: "Golang Blog Series"
date: "22nd October 2022"
tags: ["go", "files"]
published: false
words: 1500
```

The above file is a simple YAML config, we'll follow similar kind of examples for the dummy files used in the examples.

```go
package main

import (
	"log"
    "os"

	yaml "gopkg.in/yaml.v3"
)

func main() {

	f, err := os.ReadFile("blog.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}

	err = yaml.Unmarshal(f, &data)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(data)
	for k, v := range data {
		log.Println(k, ":", v)
	}
}
```

```
$ go run yaml.go

map[date:22nd October 2022 published:false tags:[go files] title:Golang Blog Series words:1500]
published : false
words : 1500
title : Golang Blog Series
date : 22nd October 2022
tags : [go files]
```

The above code and output demonstrate the usage of the `yaml.v3` package for reading a YAML file.

Firstly, we read the file into a single-string object with the [os.ReadFile()](https://pkg.go.dev/os#ReadFile) method. The method will return a `[]byte` (slice of byte) or an error. If there is an error, we simply log or panic out of the program, else we can use the [yaml.Unmarshal](https://pkg.go.dev/gopkg.in/yaml.v3#Unmarshal) method to convert the string/slice of the byte into a map or a pre-defined struct. In this example, we just keep things simple by storing the file content as `map [string, interface{}]`, i.e. a map of `string` and an `interface`. The key for YAML can be only a string or an integer. It can't have unrestricted data types like the value can have. Though if you want to be unrestrictive, you can use a map of `map[interface{}]interface{}` to make the key any shape you like to have.

So, we have created a variable called `data` as a map of `string` and `interface{}`, basically key can be a string and the value can be any type of interface depending on the parsed literally from the file object. The `Unmarshal` function takes in two parameters, the first being the slice of byte i.e. the file contents, and the second being the output variable. Now, the method does not return the parsed YAML, it can return an error if there arose, so we need to set the second parameter as a pointer to the object into which we want to store the parsed YAML.

In the example, we have called `Unmarshal(f, &data)` which will fetch the contents from the slice of bytes `f` and output the parsed YAML from the slice of bytes into the memory location of `data` and hence using `&data` indicating the pointer to the variable(fetch the memory address).

So, that is how we obtain the map of keys and values from the YAML config, thereafter, you can iterate on the map, access the keys and values, type caste them as per requirement, and basically have control over what processing needs to be done to the parsed YAML content.

```
published : false
bool

words : 1500
int

title : Golang Blog Series
string

date : 22nd October 2022
string

tags : [go files]
[]interface {}
```

I have just printed the types of the values in the above output as `log.Printf("%T", v)`, we can see the types are being correctly recognized and being parsed. The last object is indeed a `slice` so it has an interface of the slice(array) attached to it.

### Reading a TOML file

YAML and TOML are almost identical, though toml has more restrictive configuration options, but is more readable than YAML, as YAML can get complicated pretty quickly. Though both of them have their pros and cons, YAML is used everywhere in the DevOps world, configs, whereas TOML is the format of choice for python packaging, and static site generation configs.

Let's see how we can use golang to read TOML files.

```
$ go mod init <your_project_package_name>
$ go get github.com/pelletier/go-toml
```

The above commands are used for setting up a golang package or project and installing the [go-toml](https://pkg.go.dev/github.com/pelletier/go-toml) package. Once the above commands are done executing, it will generate `go.mod` and `go.sum` files used for storing dependencies and packages installed for the project locally.

```toml
[blog]
name='techstructive-blog'
tags=['go','django','vim']
author='meet gor'
active=true

[author]
name='Meet Gor'
github='mr-destructive'
twitter='meetgor21'
posts=80
```

The above is the sample file `blog.toml` which we will use to read in the go script below. The toml file has a similar structure as we have seen in the previous examples. We have different data types like string, boolean, integer, and list.

```go
package main

import (
	"log"
    "os"

	toml "github.com/pelletier/go-toml"
)

func main() {

	f, err := os.ReadFile("blog.toml")

	if err != nil {
		log.Fatal(err)
	}

	var data map[interface{}]interface{}

	err = toml.Unmarshal(f, &data)
	log.Println(data)

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range data {
		log.Println(k, ":", v)

		switch t := v.(type) {
		case map[string]interface{}:
			for a, b := range t {
				log.Println(a, ":", b)
			}
		}
	}
}
```

```
$ go run toml.go

map[author:map[github:mr-destructive name:Meet Gor posts:80 twitter:meetgor21] blog:map[active:true author:meet gor

name:techstructive-blog tags:[go django vim]]]

blog : map[active:true author:meet gor name:techstructive-blog tags:[go django vim]]
name : techstructive-blog
tags : [go django vim]
author : meet gor
active : true

author : map[github:mr-destructive name:Meet Gor posts:80 twitter:meetgor21]

name : Meet Gor
github : mr-destructive
twitter : meetgor21
posts : 80
```

So, in the above example and output, the YAML file was read and the key-value pairs inside them were read. The first thing we do is read the file `blog.toml` with `ioutil` package, with the `ReadFile` function. The function takes in the string as the path to the file to be read and returns a slice of byte. We use this slice of byte as a parameter to the [Unmarshal](https://pkg.go.dev/github.com/pelletier/go-toml#Unmarshal) method. The second paramter for the `Unmarshal` is the output variable(usually a pointer to a variable), we have created a map of `interface{]` with an `interface` as we see there can be nested keys which hold the name of the config.

The variable `data` is a map of `interface{}` to an `interface{}`, and we parse the memory address to the `data` variable to the `Unmarshal` method. Thereby the parsed `TOML` content is stored in the data variable.

```
name : techstructive-blog
string

tags : [go django vim]
[]interface{}

author : meet gor
string

active : true
bool

name : Meet Gor
string

github : mr-destructive
string

twitter : meetgor21
string

posts : 80
int64
```

The above is a verbose output for the type of the values as parsed by golang, we have `string`, `bool`, `int64`, and a `slice` (list with interface{} attached with it). Only types like `string`, `bool`, `int`, `float64` can be parsed from the Unmarshal function, other than these types, the type will have an interface attached to it.
 
 In such cases, where the type of value is not among the 4 types(string, bool, int float), we can use a pre-defined struct to parse the content from the file. Though it would require a strict structure and predictable response from the parsed file.

### Reading CSV file

We can even read a CSV file in golang, we have seen in the previous post, we have used custom delimiters in the parsing of the file.

```csv
id,name,posts,exp
21,jim,23,2
33,kevin,39,1
45,james,70,2
56,chris,89,3
```

The above file is a sample csv file, though the size is too small, we can use it as an example.

```go
package main

import (
	"encoding/csv"
    "log"
    "os"
)

func main() {
	f, err := os.Open("temp.csv")
	check_error(err)

	reader := csv.NewReader(f)

	n, err := reader.ReadAll()
	check_error(err)
	for _, line := range n {
		for _, text := range line {
			log.Println(text)
		}
	}
}
```

```
$ go run main.go
id
name
posts
exp
21
jim
23
2
33
kevin
39
1
45
james
70
2
56
chris
89
3
```

The CSV package has a [NewReader](https://pkg.go.dev/encoding/csv#NewReader) method that accepts an `io.Reader` and returns a `Reader` object. After parsing the reader, we use the [ReadAll](https://pkg.go.dev/encoding/csv#Reader.ReadAll) method to return a 2d string or an error if there exists an error while parsing the content. You can get a detailed explanation of the CSV parsing and reading in the [previous post](https://www.meetgor.com/golang-file-read/#Read%20File%20by%20a%20delimiter).


### Reading CSV from URL

The CSV file can also be read from the URL, the content of the file is a `response.Body` in place of the file object reference, in the previous example, the [os.Open()](https://pkg.go.dev/os#Open) method returns a [os.File](https://pkg.go.dev/os#File) object. 

We use the `http.Get(string)` method to get the response from the URL for reading the CSV file present on the web.


```go
package main

import (
	"encoding/csv"
	"log"
	"net/http"
)

func main() {

	url := "https://github.com/woocommerce/woocommerce/raw/master/sample-data/sample_products.csv"
	response, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return
	}

	defer response.Body.Close()

	reader := csv.NewReader(response.Body)
	n, err := reader.ReadAll()

	if err != nil {
		log.Println(err)
	}

	for _, line := range n {
		for _, text := range line {
			log.Println(text)
		}
	}
}
```

```
$ go run csv.go
<feff>ID
Type
SKU
Name
Published
Is featured?
Visibility in catalog
Short description
Description
Date sale price starts
Date sale price ends
...
...
...
```
So, that's how we can read a CSV file from the URL. By fetching the CSV URL `https://github.com/woocommerce/woocommerce/raw/master/sample-data/sample_products.csv` from the [http.Get](https://pkg.go.dev/net/http#Client.Get) method, this will get us the [response.Body](https://pkg.go.dev/net/http#Response) that contains the actual CSV file content. The response than can be parsed to the `csv.NewReader(*Os.File).ReadAll()` i.e. [reader.ReadAll()](https://pkg.go.dev/encoding/csv#Reader.ReadAll). The function returns a multidimensional slice `[][]slice` that can be iterated and parsed as per requirement.

### Reading XML file

XML is the de facto standard for RSS feeds, it is widely used in many places and is still all over the web. We'll see an example to read an XML file locally, but as we saw in the above example, we can also read an RSS link from the web.

Just like CSV, we have [encoding/xml](https://pkg.go.dev/encoding/xml), and the standard library has all the functions used for parsing the XML files.

We will be using a local XML file called `rss.xml`, and reading the contents from the tags in the file.

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<channel>
<title>Meet Gor</title>
<description>Techstructive Blog Feed</description>
<item>
<title>Why and How to make and use Vim as a text editor and customizable IDE</title>
<link>https://www.meetgor.com/vim-text-editor-ide</link>
</item>
<item>
<title>Setting up Vim for Python</title>
<link>https://www.meetgor.com/vim-for-python</link>
</item>
</channel>
```

The above example is a short part of my blog's [rss](https://www.meetgor.com/rss) feed. I have just trimmed the unwanted part and will be just using the tags that we want to fetch. 

```go
package main

import (
	"encoding/xml"
	"log"
	"os"
)

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Item        []Item   `xml:"item"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
}

func check_error(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {

	f, err := os.ReadFile("rss.xml")
    check_error(err)
	defer f.Close()

	d := Channel{}
	err = xml.Unmarshal(f, &d)
    check_error(err)

	for _, item := range d.Item {
		log.Println(item.Title)
	}
}
```

```
$ go run xml.go

{{ channel} Meet Gor Techstructive Blog Feed [{{ item} Why and How to make and use Vim as a text editor and customizable IDE https://www.meetgor.com/vim-text-editor-ide} {{ item} Setting up Vim for Python https://www.meetgor.com/vim-for-python}]}

Why and How to make and use Vim as a text editor and customizable IDE
Setting up Vim for Python
```

The above example uses a couple of `struct` like `Channel` and `Item` that stores the tag data like `title`, `description`, `link`, etc. Unlike the JSON, YAML, and toml files, XML can't parse the content directly we need a structure to parse into. And that's fine as XML is not much dynamic in terms of config, we usually have standard tags and elements which can be pre-defined in a struct type.

In this example, the RSS feed has a `Channel` tag with `title`, `description`, and `item`. 

**NOTE: Use Title case for the fields of the structs. It will make the fields public, I spent a few hours debugging that really :)**

So, we define the `Channel` struct with fields like `Title` as a string which is a tag in the file as `xml:"title"`. This means the title in the tag of the XML will be stored in the field as a string in the attribute name as `Title`. Similarly, we have fields like `Description` and `Item[]` this is a list or multiple of `item` tags that might be present in the XML file. The `XMLName` is used for identifying the parent tag for the struct, so we use `channel` for the first struct as it is the first tag appearing of the hierarchy in the XML file.

So, we create an object for the root structure as `Channel{}` (an empty object instantiated). The `xml.Unmarshal` function is parsed with the content of the file as `data` which is a slice of byte as we have seen in the previous examples. The slice of byte is then used in the `Unmarshal` method as the first parameter and the reference of an empty `Channel` object as the second parameter. The second parameter will be to store the parsed XML content from the file.
 
I have a few examples on the GitHub repository covering the reading of files from a URL for the CSV, and XML files. But, this concept in the example, can be applied to JSON, YAML, and other file formats as well.

That's it from this part. Reference for all the code examples and commands can be found in the [100 days of Golang](https://github.com/Mr-Destructive/100-days-of-golang/tree/main/scripts/files/read/config_files) GitHub repository.

## Conclusion

So, that's it from this post, we covered how to read specific configuration files like `JSON`, `CSV`, `YAML`, `TOML`, and `XML`. We saw how to read a local file and also touched on the concept to read contents from a file on the web with a URL. We also saw how we can use pre-defined structs to parse content from a file, especially for XML.

Thank you for reading. If you have any queries, questions, or feedback, you can let me know in the discussion below or on my social handles. Happy Coding :)