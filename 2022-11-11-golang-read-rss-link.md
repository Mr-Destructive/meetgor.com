---
article_html: "<h2>Reding Rss Feed</h2>\n<p>We can use golang's <a href=\"https://pkg.go.dev/encoding/xml\">encoding/xml</a>
  package to read a Rss feed. Though we have to be speicific of what type of structure
  the Rss feed has, so it is not dynamic but it works really well with structs. I
  have covered a few nuances of reading XML file in the <a href=\"https://www.meetgor.com/golang-config-file-read/#reading-xml-file\">config
  file reading</a> post of the 100 days of golang series.</p>\n<h3>Get request to
  Rss feed</h3>\n<p>We first need to send a <code>GET</code> request to the Rss feed,
  we can use the <a href=\"https://pkg.go.dev/net/http\">http</a> package to grab
  the response.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">url</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://meetgor.com/rss.xml&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">response</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">url</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">response</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">response</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  in the above example, we have used the <code>net/http</code> package to send a <code>GET</code>
  request with the <a href=\"https://pkg.go.dev/net/http#Get\">Get</a> funciton. The
  function takes in a string as a <code>URL</code> and returns either the object as
  response or an error. If there arose any error, we simply exit out of the program
  and log the error. If the error is <code>nil</code>, we return the response in the
  <code>response</code> variable. This builds up a good foundation for the next step
  to read the response body and fetching the actual bytes from the response.</p>\n<h3>Fetch
  the content from the Link</h3>\n<p>Since we have the <code>response</code> object,
  we can use the <a href=\"https://pkg.go.dev/io#ReadAll\">io.ReadAll</a> function
  to read the bytes in the response body. The function takes in the <a href=\"https://pkg.go.dev/io#Reader\">Reader</a>
  object in this case it is <a href=\"https://pkg.go.dev/io#ReadCloser\">ReadCloser</a>
  object as a http object. The function then returns the slice of bytes/int8. The
  slice then can be interpreted as string or other form data that can be used for
  parsing the xml from the response.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">url</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://meetgor.com/rss.xml&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">response</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">url</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">data</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">response</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nb\">string</span><span class=\"p\">(</span><span class=\"nx\">data</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Type -&gt; %T&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">data</span><span class=\"p\">)</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>&lt;rss&gt;\n    &lt;channel&gt;\n
  \       &lt;item&gt;\n        ...\n        ...\n        ...\n        &lt;/item&gt;\n
  \   &lt;/channel&gt;\n&lt;/rss&gt;\n\n\nType -&gt; []uint8 \n</pre></div>\n\n</pre>\n\n<p>So,
  we can see that the parsed content is indeed xml, it is type casted to string from
  the slice of bytes. This can be further used for the parsing the text as Rss structure
  and fetch the required details.</p>\n<h2>Parsing Rss with a struct</h2>\n<p>We can
  now move into creating a struct for individual tags required in the parsing.</p>\n<pre
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
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;encoding/xml&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Rss</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">XMLName</span><span class=\"w\"> </span><span class=\"nx\">xml</span><span
  class=\"p\">.</span><span class=\"nx\">Name</span><span class=\"w\"> </span><span
  class=\"s\">`xml:&quot;rss&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Channel</span><span class=\"w\"> </span><span class=\"nx\">Channel</span><span
  class=\"w\">  </span><span class=\"s\">`xml:&quot;channel&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Channel</span><span class=\"w\"> </span><span
  class=\"kd\">struct</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">XMLName</span><span
  class=\"w\">     </span><span class=\"nx\">xml</span><span class=\"p\">.</span><span
  class=\"nx\">Name</span><span class=\"w\"> </span><span class=\"s\">`xml:&quot;channel&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Title</span><span
  class=\"w\">       </span><span class=\"kt\">string</span><span class=\"w\">   </span><span
  class=\"s\">`xml:&quot;title&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Description</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\">   </span><span class=\"s\">`xml:&quot;description&quot;`</span><span
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
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">url</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://meetgor.com/rss.xml&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">response</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">url</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">data</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">response</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nb\">string</span><span class=\"p\">(</span><span class=\"nx\">data</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>If you would look at the <a href=\"https://meetgor.com/rss.xml\">rss
  feed</a>, you can see it has a structure of tags and elements. The <code>rss</code>
  tag is the root tag, followed by <code>channel</code> and other types of nested
  tags speicific for the type of information to be stored like <code>title</code>
  for the title in the feed, <code>link</code> for the link to the feed, etc.</p>\n<p>So,
  we create those as structure, the root structure is the <code>Rss</code> which we
  will create with a few attributes like <code>Channel</code> and the name of the
  current tag. In the <code>Rss</code> case the name of the tag/element is <code>rss</code>,
  so it is given the <code>xml.Name</code> as <code>xml:'rss'</code> in backticks
  indicating the type hint for the field. The next field is the <code>Channel</code>
  which is another type(custom type struct). We have defined <code>Channel</code>
  as a struct just after it that will hold information like the <code>title</code>,
  <code>description</code> of the website. We also have the <code>xml.Name</code>
  as <code>xml:&quot;channel&quot;</code> which indicates the current struct is representation
  of <code>channel</code> tag in the rss feed. Finally, we also have a custom type
  struct as <code>Item</code>. The <code>Item</code> struct has a few attributes like
  <code>Title</code>, <code>Link</code> and you can now start to see the pattern,
  you can customize it as per your requirements and speicifications.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;encoding/xml&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Rss</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">XMLName</span><span class=\"w\"> </span><span class=\"nx\">xml</span><span
  class=\"p\">.</span><span class=\"nx\">Name</span><span class=\"w\"> </span><span
  class=\"s\">`xml:&quot;rss&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Channel</span><span class=\"w\"> </span><span class=\"nx\">Channel</span><span
  class=\"w\">  </span><span class=\"s\">`xml:&quot;channel&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Channel</span><span class=\"w\"> </span><span
  class=\"kd\">struct</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">XMLName</span><span
  class=\"w\">     </span><span class=\"nx\">xml</span><span class=\"p\">.</span><span
  class=\"nx\">Name</span><span class=\"w\"> </span><span class=\"s\">`xml:&quot;channel&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Title</span><span
  class=\"w\">       </span><span class=\"kt\">string</span><span class=\"w\">   </span><span
  class=\"s\">`xml:&quot;title&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Description</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\">   </span><span class=\"s\">`xml:&quot;description&quot;`</span><span
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
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">url</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://meetgor.com/rss.xml&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">response</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">url</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">data</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">response</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"c1\">// New Code</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">d</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">Rss</span><span
  class=\"p\">{}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">xml</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">data</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">d</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">item</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">d</span><span class=\"p\">.</span><span class=\"nx\">Channel</span><span
  class=\"p\">.</span><span class=\"nx\">Item</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run main.go\n\nWhy and
  How to make and use Vim as a text editor and customizable IDE\nSetting up Vim for
  Python\nSetting up Vim for BASH Scripting\nVim: Keymapping Guide\n...\n...\n...\nDjango
  + HTMX CRUD application\nPGCLI: Postgres from the terminal\nGolang: Closures\nGolang:
  Interfaces\nGolang: Error Handling\nGolang: Paths\nGolang: File Reading\nGolang:
  JSON YAML TOML (config) File Reading.\n</pre></div>\n\n</pre>\n\n<p>So, here we
  have initialized the <code>Rss</code> struct as empty and then used the <a href=\"https://pkg.go.dev/encoding/xml#Unmarshal\">Unmarshal</a>
  method in the <code>xml</code> package. The Unmarshal method will parse the data
  as per the type of either int, float, bool or string, any other type of data will
  be discarded as interface or struct. We can usually parse any valid type of data
  into <code>Unmarshal</code> method and it generally gives a proper expected outcome.</p>\n<p>The
  Unmarshal method takes in the slice of byte and the second paramter as pointer to
  a struct or any variable that will store the parsed xml content from the slice of
  byte. The function just returns the error type, either <code>nil</code> in case
  of no errors, and returns the actual error obejct if there arise any type of error.</p>\n<p>So
  we parse the <code>data</code> which is a slice of byte to the funciton and the
  reference to the <code>d</code> object which is a empty <code>Rss</code> object.
  This will get us the data in the <code>d</code> object. We can then iterate over
  the object as per the struct and use the perform operations like type casting or
  converting types, etc to get your required data back.</p>\n<p>In the above example,
  we simply iterate over the <code>d.Channel.Item</code> which is a list of elements
  of tag <code>item</code> in the rss feed. Inside the for loop, we can access the
  object and simply print or perform any sort of operations. I have simply printed
  the list of articles with titles.</p>\n<p>Links for the code available on the <a
  href=\"https://github.com/Mr-Destructive/100-days-of-golang/blob/main/scripts/files/read/config_files/xml/rss.go\">100
  days of golang</a> GitHub repository.</p>\n<p>So, that's how we parse an XML feed
  in golang. Just plug and play if you have a similar type of structure of the Rss
  XML feed. Happy Coding :)</p>\n<div class='prevnext'>\n    <style type='text/css'>\n
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
  \   </style>\n<pre><code>&lt;a class='prev' href='/django-non-clustered-index-pg'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Create
  a Non-Clustered Index in Django with Postgres as DB&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a
  class='next' href='/python-pipreqs'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Python
  Pipreqs: Generate requirements file from the imported packages&lt;/p&gt;\n    &lt;/div&gt;\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
cover: ''
date: 2022-11-11
datetime: 2022-11-11 00:00:00+00:00
description: Reading Rss Feed with a Rss XML Link/URL in golang using encoding package
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/tils/2022-11-11-golang-read-rss-link.md
html: "<h2>Reding Rss Feed</h2>\n<p>We can use golang's <a href=\"https://pkg.go.dev/encoding/xml\">encoding/xml</a>
  package to read a Rss feed. Though we have to be speicific of what type of structure
  the Rss feed has, so it is not dynamic but it works really well with structs. I
  have covered a few nuances of reading XML file in the <a href=\"https://www.meetgor.com/golang-config-file-read/#reading-xml-file\">config
  file reading</a> post of the 100 days of golang series.</p>\n<h3>Get request to
  Rss feed</h3>\n<p>We first need to send a <code>GET</code> request to the Rss feed,
  we can use the <a href=\"https://pkg.go.dev/net/http\">http</a> package to grab
  the response.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">url</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://meetgor.com/rss.xml&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">response</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">url</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">response</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">response</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  in the above example, we have used the <code>net/http</code> package to send a <code>GET</code>
  request with the <a href=\"https://pkg.go.dev/net/http#Get\">Get</a> funciton. The
  function takes in a string as a <code>URL</code> and returns either the object as
  response or an error. If there arose any error, we simply exit out of the program
  and log the error. If the error is <code>nil</code>, we return the response in the
  <code>response</code> variable. This builds up a good foundation for the next step
  to read the response body and fetching the actual bytes from the response.</p>\n<h3>Fetch
  the content from the Link</h3>\n<p>Since we have the <code>response</code> object,
  we can use the <a href=\"https://pkg.go.dev/io#ReadAll\">io.ReadAll</a> function
  to read the bytes in the response body. The function takes in the <a href=\"https://pkg.go.dev/io#Reader\">Reader</a>
  object in this case it is <a href=\"https://pkg.go.dev/io#ReadCloser\">ReadCloser</a>
  object as a http object. The function then returns the slice of bytes/int8. The
  slice then can be interpreted as string or other form data that can be used for
  parsing the xml from the response.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">url</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://meetgor.com/rss.xml&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">response</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">url</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">data</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">response</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nb\">string</span><span class=\"p\">(</span><span class=\"nx\">data</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Type -&gt; %T&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">data</span><span class=\"p\">)</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>&lt;rss&gt;\n    &lt;channel&gt;\n
  \       &lt;item&gt;\n        ...\n        ...\n        ...\n        &lt;/item&gt;\n
  \   &lt;/channel&gt;\n&lt;/rss&gt;\n\n\nType -&gt; []uint8 \n</pre></div>\n\n</pre>\n\n<p>So,
  we can see that the parsed content is indeed xml, it is type casted to string from
  the slice of bytes. This can be further used for the parsing the text as Rss structure
  and fetch the required details.</p>\n<h2>Parsing Rss with a struct</h2>\n<p>We can
  now move into creating a struct for individual tags required in the parsing.</p>\n<pre
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
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;encoding/xml&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Rss</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">XMLName</span><span class=\"w\"> </span><span class=\"nx\">xml</span><span
  class=\"p\">.</span><span class=\"nx\">Name</span><span class=\"w\"> </span><span
  class=\"s\">`xml:&quot;rss&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Channel</span><span class=\"w\"> </span><span class=\"nx\">Channel</span><span
  class=\"w\">  </span><span class=\"s\">`xml:&quot;channel&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Channel</span><span class=\"w\"> </span><span
  class=\"kd\">struct</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">XMLName</span><span
  class=\"w\">     </span><span class=\"nx\">xml</span><span class=\"p\">.</span><span
  class=\"nx\">Name</span><span class=\"w\"> </span><span class=\"s\">`xml:&quot;channel&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Title</span><span
  class=\"w\">       </span><span class=\"kt\">string</span><span class=\"w\">   </span><span
  class=\"s\">`xml:&quot;title&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Description</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\">   </span><span class=\"s\">`xml:&quot;description&quot;`</span><span
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
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">url</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://meetgor.com/rss.xml&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">response</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">url</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">data</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">response</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"nx\">log</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nb\">string</span><span class=\"p\">(</span><span class=\"nx\">data</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>If you would look at the <a href=\"https://meetgor.com/rss.xml\">rss
  feed</a>, you can see it has a structure of tags and elements. The <code>rss</code>
  tag is the root tag, followed by <code>channel</code> and other types of nested
  tags speicific for the type of information to be stored like <code>title</code>
  for the title in the feed, <code>link</code> for the link to the feed, etc.</p>\n<p>So,
  we create those as structure, the root structure is the <code>Rss</code> which we
  will create with a few attributes like <code>Channel</code> and the name of the
  current tag. In the <code>Rss</code> case the name of the tag/element is <code>rss</code>,
  so it is given the <code>xml.Name</code> as <code>xml:'rss'</code> in backticks
  indicating the type hint for the field. The next field is the <code>Channel</code>
  which is another type(custom type struct). We have defined <code>Channel</code>
  as a struct just after it that will hold information like the <code>title</code>,
  <code>description</code> of the website. We also have the <code>xml.Name</code>
  as <code>xml:&quot;channel&quot;</code> which indicates the current struct is representation
  of <code>channel</code> tag in the rss feed. Finally, we also have a custom type
  struct as <code>Item</code>. The <code>Item</code> struct has a few attributes like
  <code>Title</code>, <code>Link</code> and you can now start to see the pattern,
  you can customize it as per your requirements and speicifications.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;encoding/xml&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Rss</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">XMLName</span><span class=\"w\"> </span><span class=\"nx\">xml</span><span
  class=\"p\">.</span><span class=\"nx\">Name</span><span class=\"w\"> </span><span
  class=\"s\">`xml:&quot;rss&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Channel</span><span class=\"w\"> </span><span class=\"nx\">Channel</span><span
  class=\"w\">  </span><span class=\"s\">`xml:&quot;channel&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Channel</span><span class=\"w\"> </span><span
  class=\"kd\">struct</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">XMLName</span><span
  class=\"w\">     </span><span class=\"nx\">xml</span><span class=\"p\">.</span><span
  class=\"nx\">Name</span><span class=\"w\"> </span><span class=\"s\">`xml:&quot;channel&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Title</span><span
  class=\"w\">       </span><span class=\"kt\">string</span><span class=\"w\">   </span><span
  class=\"s\">`xml:&quot;title&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Description</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\">   </span><span class=\"s\">`xml:&quot;description&quot;`</span><span
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
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">url</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://meetgor.com/rss.xml&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">response</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">url</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">data</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">response</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"c1\">// New Code</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">d</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">Rss</span><span
  class=\"p\">{}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">xml</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">data</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">d</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">item</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">d</span><span class=\"p\">.</span><span class=\"nx\">Channel</span><span
  class=\"p\">.</span><span class=\"nx\">Item</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run main.go\n\nWhy and
  How to make and use Vim as a text editor and customizable IDE\nSetting up Vim for
  Python\nSetting up Vim for BASH Scripting\nVim: Keymapping Guide\n...\n...\n...\nDjango
  + HTMX CRUD application\nPGCLI: Postgres from the terminal\nGolang: Closures\nGolang:
  Interfaces\nGolang: Error Handling\nGolang: Paths\nGolang: File Reading\nGolang:
  JSON YAML TOML (config) File Reading.\n</pre></div>\n\n</pre>\n\n<p>So, here we
  have initialized the <code>Rss</code> struct as empty and then used the <a href=\"https://pkg.go.dev/encoding/xml#Unmarshal\">Unmarshal</a>
  method in the <code>xml</code> package. The Unmarshal method will parse the data
  as per the type of either int, float, bool or string, any other type of data will
  be discarded as interface or struct. We can usually parse any valid type of data
  into <code>Unmarshal</code> method and it generally gives a proper expected outcome.</p>\n<p>The
  Unmarshal method takes in the slice of byte and the second paramter as pointer to
  a struct or any variable that will store the parsed xml content from the slice of
  byte. The function just returns the error type, either <code>nil</code> in case
  of no errors, and returns the actual error obejct if there arise any type of error.</p>\n<p>So
  we parse the <code>data</code> which is a slice of byte to the funciton and the
  reference to the <code>d</code> object which is a empty <code>Rss</code> object.
  This will get us the data in the <code>d</code> object. We can then iterate over
  the object as per the struct and use the perform operations like type casting or
  converting types, etc to get your required data back.</p>\n<p>In the above example,
  we simply iterate over the <code>d.Channel.Item</code> which is a list of elements
  of tag <code>item</code> in the rss feed. Inside the for loop, we can access the
  object and simply print or perform any sort of operations. I have simply printed
  the list of articles with titles.</p>\n<p>Links for the code available on the <a
  href=\"https://github.com/Mr-Destructive/100-days-of-golang/blob/main/scripts/files/read/config_files/xml/rss.go\">100
  days of golang</a> GitHub repository.</p>\n<p>So, that's how we parse an XML feed
  in golang. Just plug and play if you have a similar type of structure of the Rss
  XML feed. Happy Coding :)</p>\n<div class='prevnext'>\n    <style type='text/css'>\n
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
  \   </style>\n<pre><code>&lt;a class='prev' href='/django-non-clustered-index-pg'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Create
  a Non-Clustered Index in Django with Postgres as DB&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a
  class='next' href='/python-pipreqs'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Python
  Pipreqs: Generate requirements file from the imported packages&lt;/p&gt;\n    &lt;/div&gt;\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
long_description: We can use golang We first need to send a  So, in the above example,
  we have used the  Since we have the  So, we can see that the parsed content is indeed
  xml, it is type casted to string from the slice of bytes. This can be further used
  for the pars
now: 2025-05-01 18:11:33.315152
path: blog/tils/2022-11-11-golang-read-rss-link.md
slug: golang-read-rss-feed
status: published-til
tags:
- go
templateKey: til
title: Read a Rss Feed with a URL in Golang
today: 2025-05-01
---

## Reding Rss Feed

We can use golang's [encoding/xml](https://pkg.go.dev/encoding/xml) package to read a Rss feed. Though we have to be speicific of what type of structure the Rss feed has, so it is not dynamic but it works really well with structs. I have covered a few nuances of reading XML file in the [config file reading](https://www.meetgor.com/golang-config-file-read/#reading-xml-file) post of the 100 days of golang series.

### Get request to Rss feed

We first need to send a `GET` request to the Rss feed, we can use the [http](https://pkg.go.dev/net/http) package to grab the response.

```go
package main

import (
	"log"
	"net/http"
)

func main() {

	url := "https://meetgor.com/rss.xml"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

    log.Println(response.Body)
	defer response.Body.Close()

}
```

So, in the above example, we have used the `net/http` package to send a `GET` request with the [Get](https://pkg.go.dev/net/http#Get) funciton. The function takes in a string as a `URL` and returns either the object as response or an error. If there arose any error, we simply exit out of the program and log the error. If the error is `nil`, we return the response in the `response` variable. This builds up a good foundation for the next step to read the response body and fetching the actual bytes from the response.

### Fetch the content from the Link

Since we have the `response` object, we can use the [io.ReadAll](https://pkg.go.dev/io#ReadAll) function to read the bytes in the response body. The function takes in the [Reader](https://pkg.go.dev/io#Reader) object in this case it is [ReadCloser](https://pkg.go.dev/io#ReadCloser) object as a http object. The function then returns the slice of bytes/int8. The slice then can be interpreted as string or other form data that can be used for parsing the xml from the response.

```go
package main

import (
	"log"
	"net/http"
    "io"
)

func main() {

	url := "https://meetgor.com/rss.xml"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

    log.Println(string(data))
    log.Printf("Type -> %T", data)
}
```

```
<rss>
    <channel>
        <item>
        ...
        ...
        ...
        </item>
    </channel>
</rss>


Type -> []uint8 

```

So, we can see that the parsed content is indeed xml, it is type casted to string from the slice of bytes. This can be further used for the parsing the text as Rss structure and fetch the required details.

## Parsing Rss with a struct

We can now move into creating a struct for individual tags required in the parsing.

```go
package main

import (
    "encoding/xml"
	"io"
	"log"
	"net/http"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

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

func main() {

	url := "https://meetgor.com/rss.xml"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

    log.Println(string(data))
}
```

If you would look at the [rss feed](https://meetgor.com/rss.xml), you can see it has a structure of tags and elements. The `rss` tag is the root tag, followed by `channel` and other types of nested tags speicific for the type of information to be stored like `title` for the title in the feed, `link` for the link to the feed, etc. 

So, we create those as structure, the root structure is the `Rss` which we will create with a few attributes like `Channel` and the name of the current tag. In the `Rss` case the name of the tag/element is `rss`, so it is given the `xml.Name` as `xml:'rss'` in backticks indicating the type hint for the field. The next field is the `Channel` which is another type(custom type struct). We have defined `Channel` as a struct just after it that will hold information like the `title`, `description` of the website. We also have the `xml.Name` as `xml:"channel"` which indicates the current struct is representation of `channel` tag in the rss feed. Finally, we also have a custom type struct as `Item`. The `Item` struct has a few attributes like `Title`, `Link` and you can now start to see the pattern, you can customize it as per your requirements and speicifications.

```go
package main

import (
    "encoding/xml"
	"io"
	"log"
	"net/http"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

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

func main() {

	url := "https://meetgor.com/rss.xml"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

    // New Code

	d := Rss{}
	err = xml.Unmarshal(data, &d)

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range d.Channel.Item {
		log.Println(item.Title)
	}
}
```

```
$ go run main.go

Why and How to make and use Vim as a text editor and customizable IDE
Setting up Vim for Python
Setting up Vim for BASH Scripting
Vim: Keymapping Guide
...
...
...
Django + HTMX CRUD application
PGCLI: Postgres from the terminal
Golang: Closures
Golang: Interfaces
Golang: Error Handling
Golang: Paths
Golang: File Reading
Golang: JSON YAML TOML (config) File Reading.
```

So, here we have initialized the `Rss` struct as empty and then used the [Unmarshal](https://pkg.go.dev/encoding/xml#Unmarshal) method in the `xml` package. The Unmarshal method will parse the data as per the type of either int, float, bool or string, any other type of data will be discarded as interface or struct. We can usually parse any valid type of data into `Unmarshal` method and it generally gives a proper expected outcome.

The Unmarshal method takes in the slice of byte and the second paramter as pointer to a struct or any variable that will store the parsed xml content from the slice of byte. The function just returns the error type, either `nil` in case of no errors, and returns the actual error obejct if there arise any type of error.

So we parse the `data` which is a slice of byte to the funciton and the reference to the `d` object which is a empty `Rss` object. This will get us the data in the `d` object. We can then iterate over the object as per the struct and use the perform operations like type casting or converting types, etc to get your required data back.

In the above example, we simply iterate over the `d.Channel.Item` which is a list of elements of tag `item` in the rss feed. Inside the for loop, we can access the object and simply print or perform any sort of operations. I have simply printed the list of articles with titles. 

Links for the code available on the [100 days of golang](https://github.com/Mr-Destructive/100-days-of-golang/blob/main/scripts/files/read/config_files/xml/rss.go) GitHub repository.

So, that's how we parse an XML feed in golang. Just plug and play if you have a similar type of structure of the Rss XML feed. Happy Coding :)
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
    
    <a class='prev' href='/django-non-clustered-index-pg'>
    
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M13.5 8.25L9.75 12L13.5 15.75" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"> </path>
        </svg>
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>prev</p>
            <p class='prevnext-title'>Create a Non-Clustered Index in Django with Postgres as DB</p>
        </div>
    </a>
    
    <a class='next' href='/python-pipreqs'>
    
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>next</p>
            <p class='prevnext-title'>Python Pipreqs: Generate requirements file from the imported packages</p>
        </div>
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.5 15.75L14.25 12L10.5 8.25" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path>
        </svg>
    </a>
  </div>