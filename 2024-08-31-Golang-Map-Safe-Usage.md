---
article_html: "<h2>Introduction</h2>\n<p>This week, I was working on one of the API
  wrapper packages for golang, and that dealt with sending post requests with URL
  encoded values, setting cookies, and all the fun stuff. However, while I was constructing
  the body, I was using <a href=\"https://pkg.go.dev/net/url#Values\">url.Value</a>
  type to construct the body, and use that to add and set key-value pairs. However,
  I was getting a wired <code>nil</code> pointer reference error in some of the parts,
  I thought it was because of some of the variables I set manually. However, by debugging
  closer, I found out a common pitfall or bad practice of just declaring a type but
  initializing it and that causing <code>nil</code> reference errors.</p>\n<p>In this
  post, I will cover, what are maps, how to create maps, and especially how to properly
  declare and initialize them. Create a proper distinction between the declaration
  and initialization of maps or any similar data type in golang.</p>\n<h2>What is
  a map in Golang?</h2>\n<p>A <a href=\"https://go.dev/src/runtime/map.go\">map</a>
  or a hashmap in golang is a basic data type that allows us to store key-value pairs.
  Under the hood, it is a header map-like data structure that holds buckets, which
  are basically pointers to bucket arrays (contiguous memory). It has hash codes that
  store the actual key-value pairs, and pointers to new buckets if the current overflows
  with the number of keys. This is a really smart data structure that provides almost
  constant time access.</p>\n<h2>How to create maps in Golang</h2>\n<p>To create a
  simple map in golang, you can take an example of a letter frequency counter using
  a map of string and integer. The map will store the letters as keys and their frequency
  as values.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nx\">words</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;hello how are you&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">letters</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"kt\">string</span><span class=\"p\">]</span><span
  class=\"kt\">int</span><span class=\"p\">{}</span><span class=\"w\"></span>\n\n<span
  class=\"w\">    </span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">word</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">words</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nx\">wordCount</span><span class=\"p\">[</span><span class=\"nx\">word</span><span
  class=\"p\">]</span><span class=\"o\">++</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"w\">    </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Word
  counts:&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">word</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">count</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">wordCount</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;%s: %d\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">word</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">count</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\n\nWord counts:\ne: <span class=\"m\">2</span>\n : <span class=\"m\">3</span>\nw:
  <span class=\"m\">1</span>\nr: <span class=\"m\">1</span>\ny: <span class=\"m\">1</span>\nu:
  <span class=\"m\">1</span>\nh: <span class=\"m\">2</span>\nl: <span class=\"m\">2</span>\no:
  <span class=\"m\">3</span>\na: <span class=\"m\">1</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  by initializing the map as <code>map[string]int{}</code> you will get an empty map.
  This can be then used to populate the keys and values, we iterate over the string,
  and for each character (rune) we cast that byte of character into the string and
  increment the value, the zero value for int is <code>0</code>, so by default if
  the key is not present, it will be zero, it is a bit of double-edged swords though,
  we never know the key is present in the map with the value <code>0</code> or the
  key is not present but the default value is <code>0</code>. For that, you need to
  check if the <code>key</code> exists in the map or not.</p>\n<p>For further details,
  you can check out my <a href=\"https://www.meetgor.com/golang-maps/\">Golang Maps</a>
  post in detail.</p>\n<h2>Difference between declaration and initialization</h2>\n<p>There
  is a difference in declaring and initializing any variable in a programming language
  and has to do a lot more with the implementation of the underlying type. In the
  case of primary data types like <code>int</code>, <code>string</code>, <code>float</code>,
  etc. there are default/zero values, so that is the same as the declaration and initialization
  of the variables. However, in the case of maps and slices, the declaration is just
  making sure the variable is available to the scope of the program, however for initialization
  is setting it to its default/zero value or the actual value that should be assigned.</p>\n<p>So,
  declaration simply makes the variable available within the scope of the program.
  For maps and slices, declaring a variable without initialization sets it to <code>nil</code>,
  meaning it points to no allocated memory and cannot be used directly.</p>\n<p>Whereas
  the <code>initialization</code> allocates memory and sets the variable to a usable
  state. For maps and slices, you need to explicitly initialize them using syntax
  like <code>myMap = make(map[keyType]valueType)</code> or <code>slice = []type{}</code>.
  Without this initialization, attempting to use the map or slice will lead to runtime
  errors, such as panics for accessing or modifying a nil map or slice.</p>\n<p>Let's
  looks at the values of a map when it is declared/initialized/not initialized.</p>\n<p>Imagine
  you're building a configuration manager that reads settings from a map. The map
  will be declared globally but initialized only when the configuration is loaded.</p>\n<ol>\n<li>Declared
  but not initialized</li>\n</ol>\n<p>The below code demonstrates a map access that
  is not initialized.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"c1\">// Global map to store configuration settings</span><span
  class=\"w\"></span>\n<span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">configSettings</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"kt\">string</span><span class=\"p\">]</span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"c1\">// Declared
  but not initialized</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// Attempt to get a configuration setting before initializing the map</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">serverPort</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">getConfigSetting</span><span
  class=\"p\">(</span><span class=\"s\">&quot;server_port&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Printf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Server port: %s\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">serverPort</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">getConfigSetting</span><span
  class=\"p\">(</span><span class=\"nx\">key</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">configSettings</span><span class=\"w\"> </span><span
  class=\"o\">==</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Configuration settings map is not initialized&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">value</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">exists</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">configSettings</span><span class=\"p\">[</span><span
  class=\"nx\">key</span><span class=\"p\">]</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"p\">!</span><span class=\"nx\">exists</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"s\">&quot;Setting
  not found&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">value</span><span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\nServer port: Setting not found\n</pre></div>\n\n</pre>\n\n<ol start=\"2\">\n<li>Declared
  and Initialized at the same time</li>\n</ol>\n<p>The below code demonstrates a map
  access that is initialized at the same time.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"c1\">// Global map to store configuration settings</span><span
  class=\"w\"></span>\n<span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">configSettings</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"kd\">map</span><span class=\"p\">[</span><span
  class=\"kt\">string</span><span class=\"p\">]</span><span class=\"kt\">string</span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;server_port&quot;</span><span class=\"p\">:</span><span class=\"w\">
  \ </span><span class=\"s\">&quot;8080&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"s\">&quot;database_url&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s\">&quot;localhost:5432&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">serverPort</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">getConfigSetting</span><span
  class=\"p\">(</span><span class=\"s\">&quot;server_port&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Printf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Server port: %s\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">serverPort</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">getConfigSetting</span><span
  class=\"p\">(</span><span class=\"nx\">key</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">value</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">exists</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">configSettings</span><span
  class=\"p\">[</span><span class=\"nx\">key</span><span class=\"p\">]</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"p\">!</span><span class=\"nx\">exists</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"s\">&quot;Setting
  not found&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">value</span><span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\nServer port: <span class=\"m\">8080</span>\n</pre></div>\n\n</pre>\n\n<ol
  start=\"3\">\n<li>Declared and later initialized</li>\n</ol>\n<p>The below code
  demonstrates a map access that is initialized later.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"c1\">// Global map to store configuration settings</span><span
  class=\"w\"></span>\n<span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">configSettings</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"kt\">string</span><span class=\"p\">]</span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"c1\">// declared
  but not initialized</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// Initialize configuration settings</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">initializeConfigSettings</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// if the function
  is not called, the map will be nil</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"c1\">// Get a configuration setting safely</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">serverPort</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">getConfigSetting</span><span
  class=\"p\">(</span><span class=\"s\">&quot;server_port&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Printf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Server port: %s\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">serverPort</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">initializeConfigSettings</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">configSettings</span><span class=\"w\"> </span><span class=\"o\">==</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">configSettings</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nb\">make</span><span class=\"p\">(</span><span
  class=\"kd\">map</span><span class=\"p\">[</span><span class=\"kt\">string</span><span
  class=\"p\">]</span><span class=\"kt\">string</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"c1\">// Properly initialize the map</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">configSettings</span><span
  class=\"p\">[</span><span class=\"s\">&quot;server_port&quot;</span><span class=\"p\">]</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"s\">&quot;8080&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">configSettings</span><span
  class=\"p\">[</span><span class=\"s\">&quot;database_url&quot;</span><span class=\"p\">]</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"s\">&quot;localhost:5432&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Configuration settings initialized&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">getConfigSetting</span><span
  class=\"p\">(</span><span class=\"nx\">key</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">configSettings</span><span class=\"w\"> </span><span
  class=\"o\">==</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Configuration settings map is not initialized&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">value</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">exists</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">configSettings</span><span class=\"p\">[</span><span
  class=\"nx\">key</span><span class=\"p\">]</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"p\">!</span><span class=\"nx\">exists</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"s\">&quot;Setting
  not found&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">value</span><span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\nConfiguration settings initialized\nServer port: <span class=\"m\">8080</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above code, we declared the global map <code>configSettings</code> but didn't
  initialize it at that point, until we wanted to access the map. We initialize the
  map in the main function, this main function could be other specific parts of the
  code, and the global variable <code>configSettings</code> a map from another part
  of the code, by initializing it in the required scope, we prevent it from causing
  nil pointer access errors. We only initialize the map if it is <code>nil</code>
  i.e. it has not been initialized elsewhere in the code. This prevents overriding
  the map/flushing out the config set from other parts of the scope.</p>\n<h2>Pitfalls
  in access of un-initialized maps</h2>\n<p>But since it deals with pointers, it comes
  with its own pitfalls like nil pointers access when the map is not initialized.</p>\n<p>Let's
  take a look at an example, a real case where this might happen.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/url&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">vals</span><span
  class=\"w\"> </span><span class=\"nx\">url</span><span class=\"p\">.</span><span
  class=\"nx\">Values</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nx\">vals</span><span class=\"p\">.</span><span class=\"nx\">Add</span><span
  class=\"p\">(</span><span class=\"s\">&quot;foo&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;bar&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">vals</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>This
  will result in a runtime panic.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>nginx</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">$</span><span class=\"w\"> </span><span class=\"s\">go</span><span class=\"w\">
  </span><span class=\"s\">run</span><span class=\"w\"> </span><span class=\"s\">main.go</span><span
  class=\"w\"></span>\n<span class=\"s\">panic:</span><span class=\"w\"> </span><span
  class=\"s\">assignment</span><span class=\"w\"> </span><span class=\"s\">to</span><span
  class=\"w\"> </span><span class=\"s\">entry</span><span class=\"w\"> </span><span
  class=\"s\">in</span><span class=\"w\"> </span><span class=\"s\">nil</span><span
  class=\"w\"> </span><span class=\"s\">map</span><span class=\"w\"></span>\n\n<span
  class=\"s\">goroutine</span><span class=\"w\"> </span><span class=\"mi\">1</span><span
  class=\"w\"> </span><span class=\"s\">[running]:</span><span class=\"w\"></span>\n<span
  class=\"s\">net/url.Values.Add(...)</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"s\">/usr/local/go/src/net/url/url.go:902</span><span
  class=\"w\"></span>\n<span class=\"s\">main.main()</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"s\">/home/meet/code/playground/go/main.go:10</span><span
  class=\"w\"> </span><span class=\"s\">+0x2d</span><span class=\"w\"></span>\n<span
  class=\"s\">exit</span><span class=\"w\"> </span><span class=\"s\">status</span><span
  class=\"w\"> </span><span class=\"mi\">2</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>This
  is because the <a href=\"https://pkg.go.dev/net/url#Values\">url.Values</a> is a
  map of string and a list of string values. Since the underlying type is a map for
  <code>Values</code>, and in the example, we only have declared the variable <code>vals</code>
  with the type <code>url.Values</code>, it will point to a <code>nil</code> reference,
  hence the message on adding the value to the type. So, it is a good practice to
  use <code>make</code> while declaring or initializing a map data type. If you are
  not sure the underlying type is <code>map</code> then you could use <code>Type{}</code>
  to initialize an empty value of that type.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;net/url&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nx\">vals</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nb\">make</span><span class=\"p\">(</span><span
  class=\"nx\">url</span><span class=\"p\">.</span><span class=\"nx\">Values</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"c1\">// OR</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"c1\">// vals := url.Values{}</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">vals</span><span class=\"p\">.</span><span class=\"nx\">Add</span><span
  class=\"p\">(</span><span class=\"s\">&quot;foo&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;bar&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">vals</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run urlvals.go\nmap<span class=\"o\">[</span>foo:<span class=\"o\">[</span>bar<span
  class=\"o\">]]</span>\n<span class=\"nv\">foo</span><span class=\"o\">=</span>bar\n</pre></div>\n\n</pre>\n\n<p>It
  is also recommended by the <a href=\"https://go.dev/blog/maps\">golang team</a>
  to use the make function while initializing a map. So, either use <code>make</code>
  for maps, slices, and channels, or initialize the empty value variable with <code>Type{}</code>.
  Both of them work similarly, but the latter is more generally applicable to structs
  as well.</p>\n<h2>Conclusion</h2>\n<p>Understanding the difference between declaring
  and initializing maps in Golang is essential for any developer, not just in golang,
  but in general. As we've explored, simply declaring a map variable without initializing
  it can lead to runtime errors, such as panics when attempting to access or modify
  a nil map. Initializing a map ensures that it is properly allocated in memory and
  ready for use, thereby avoiding these pitfalls.</p>\n<p>By following best practices—such
  as using the make function or initializing with Type{}—you can prevent common issues
  related to uninitialized maps. Always ensure that maps and slices are explicitly
  initialized before use to safeguard against unexpected nil pointer dereferences</p>\n<p>Thank
  you for reading this post, If you have any questions, feedback, and suggestions,
  feel free to drop them in the comments.</p>\n<p>Happy Coding :)</p>\n"
cover: ''
date: 2024-08-31
datetime: 2024-08-31 00:00:00+00:00
description: Walkthrough the differences and pitfalls in declaring and initializing
  maps in Golang
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2024-08-31-Golang-Map-Safe-Usage.md
html: "<h2>Introduction</h2>\n<p>This week, I was working on one of the API wrapper
  packages for golang, and that dealt with sending post requests with URL encoded
  values, setting cookies, and all the fun stuff. However, while I was constructing
  the body, I was using <a href=\"https://pkg.go.dev/net/url#Values\">url.Value</a>
  type to construct the body, and use that to add and set key-value pairs. However,
  I was getting a wired <code>nil</code> pointer reference error in some of the parts,
  I thought it was because of some of the variables I set manually. However, by debugging
  closer, I found out a common pitfall or bad practice of just declaring a type but
  initializing it and that causing <code>nil</code> reference errors.</p>\n<p>In this
  post, I will cover, what are maps, how to create maps, and especially how to properly
  declare and initialize them. Create a proper distinction between the declaration
  and initialization of maps or any similar data type in golang.</p>\n<h2>What is
  a map in Golang?</h2>\n<p>A <a href=\"https://go.dev/src/runtime/map.go\">map</a>
  or a hashmap in golang is a basic data type that allows us to store key-value pairs.
  Under the hood, it is a header map-like data structure that holds buckets, which
  are basically pointers to bucket arrays (contiguous memory). It has hash codes that
  store the actual key-value pairs, and pointers to new buckets if the current overflows
  with the number of keys. This is a really smart data structure that provides almost
  constant time access.</p>\n<h2>How to create maps in Golang</h2>\n<p>To create a
  simple map in golang, you can take an example of a letter frequency counter using
  a map of string and integer. The map will store the letters as keys and their frequency
  as values.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nx\">words</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;hello how are you&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">letters</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"kt\">string</span><span class=\"p\">]</span><span
  class=\"kt\">int</span><span class=\"p\">{}</span><span class=\"w\"></span>\n\n<span
  class=\"w\">    </span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">word</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">words</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nx\">wordCount</span><span class=\"p\">[</span><span class=\"nx\">word</span><span
  class=\"p\">]</span><span class=\"o\">++</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"w\">    </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Word
  counts:&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">word</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">count</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">wordCount</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;%s: %d\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">word</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">count</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\n\nWord counts:\ne: <span class=\"m\">2</span>\n : <span class=\"m\">3</span>\nw:
  <span class=\"m\">1</span>\nr: <span class=\"m\">1</span>\ny: <span class=\"m\">1</span>\nu:
  <span class=\"m\">1</span>\nh: <span class=\"m\">2</span>\nl: <span class=\"m\">2</span>\no:
  <span class=\"m\">3</span>\na: <span class=\"m\">1</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  by initializing the map as <code>map[string]int{}</code> you will get an empty map.
  This can be then used to populate the keys and values, we iterate over the string,
  and for each character (rune) we cast that byte of character into the string and
  increment the value, the zero value for int is <code>0</code>, so by default if
  the key is not present, it will be zero, it is a bit of double-edged swords though,
  we never know the key is present in the map with the value <code>0</code> or the
  key is not present but the default value is <code>0</code>. For that, you need to
  check if the <code>key</code> exists in the map or not.</p>\n<p>For further details,
  you can check out my <a href=\"https://www.meetgor.com/golang-maps/\">Golang Maps</a>
  post in detail.</p>\n<h2>Difference between declaration and initialization</h2>\n<p>There
  is a difference in declaring and initializing any variable in a programming language
  and has to do a lot more with the implementation of the underlying type. In the
  case of primary data types like <code>int</code>, <code>string</code>, <code>float</code>,
  etc. there are default/zero values, so that is the same as the declaration and initialization
  of the variables. However, in the case of maps and slices, the declaration is just
  making sure the variable is available to the scope of the program, however for initialization
  is setting it to its default/zero value or the actual value that should be assigned.</p>\n<p>So,
  declaration simply makes the variable available within the scope of the program.
  For maps and slices, declaring a variable without initialization sets it to <code>nil</code>,
  meaning it points to no allocated memory and cannot be used directly.</p>\n<p>Whereas
  the <code>initialization</code> allocates memory and sets the variable to a usable
  state. For maps and slices, you need to explicitly initialize them using syntax
  like <code>myMap = make(map[keyType]valueType)</code> or <code>slice = []type{}</code>.
  Without this initialization, attempting to use the map or slice will lead to runtime
  errors, such as panics for accessing or modifying a nil map or slice.</p>\n<p>Let's
  looks at the values of a map when it is declared/initialized/not initialized.</p>\n<p>Imagine
  you're building a configuration manager that reads settings from a map. The map
  will be declared globally but initialized only when the configuration is loaded.</p>\n<ol>\n<li>Declared
  but not initialized</li>\n</ol>\n<p>The below code demonstrates a map access that
  is not initialized.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"c1\">// Global map to store configuration settings</span><span
  class=\"w\"></span>\n<span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">configSettings</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"kt\">string</span><span class=\"p\">]</span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"c1\">// Declared
  but not initialized</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// Attempt to get a configuration setting before initializing the map</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">serverPort</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">getConfigSetting</span><span
  class=\"p\">(</span><span class=\"s\">&quot;server_port&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Printf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Server port: %s\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">serverPort</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">getConfigSetting</span><span
  class=\"p\">(</span><span class=\"nx\">key</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">configSettings</span><span class=\"w\"> </span><span
  class=\"o\">==</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Configuration settings map is not initialized&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">value</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">exists</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">configSettings</span><span class=\"p\">[</span><span
  class=\"nx\">key</span><span class=\"p\">]</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"p\">!</span><span class=\"nx\">exists</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"s\">&quot;Setting
  not found&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">value</span><span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\nServer port: Setting not found\n</pre></div>\n\n</pre>\n\n<ol start=\"2\">\n<li>Declared
  and Initialized at the same time</li>\n</ol>\n<p>The below code demonstrates a map
  access that is initialized at the same time.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"c1\">// Global map to store configuration settings</span><span
  class=\"w\"></span>\n<span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">configSettings</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"kd\">map</span><span class=\"p\">[</span><span
  class=\"kt\">string</span><span class=\"p\">]</span><span class=\"kt\">string</span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;server_port&quot;</span><span class=\"p\">:</span><span class=\"w\">
  \ </span><span class=\"s\">&quot;8080&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"s\">&quot;database_url&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s\">&quot;localhost:5432&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">serverPort</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">getConfigSetting</span><span
  class=\"p\">(</span><span class=\"s\">&quot;server_port&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Printf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Server port: %s\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">serverPort</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">getConfigSetting</span><span
  class=\"p\">(</span><span class=\"nx\">key</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">value</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">exists</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">configSettings</span><span
  class=\"p\">[</span><span class=\"nx\">key</span><span class=\"p\">]</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"p\">!</span><span class=\"nx\">exists</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"s\">&quot;Setting
  not found&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">value</span><span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\nServer port: <span class=\"m\">8080</span>\n</pre></div>\n\n</pre>\n\n<ol
  start=\"3\">\n<li>Declared and later initialized</li>\n</ol>\n<p>The below code
  demonstrates a map access that is initialized later.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"c1\">// Global map to store configuration settings</span><span
  class=\"w\"></span>\n<span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">configSettings</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"kt\">string</span><span class=\"p\">]</span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"c1\">// declared
  but not initialized</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// Initialize configuration settings</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">initializeConfigSettings</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// if the function
  is not called, the map will be nil</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"c1\">// Get a configuration setting safely</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">serverPort</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">getConfigSetting</span><span
  class=\"p\">(</span><span class=\"s\">&quot;server_port&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Printf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Server port: %s\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">serverPort</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">initializeConfigSettings</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">configSettings</span><span class=\"w\"> </span><span class=\"o\">==</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">configSettings</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nb\">make</span><span class=\"p\">(</span><span
  class=\"kd\">map</span><span class=\"p\">[</span><span class=\"kt\">string</span><span
  class=\"p\">]</span><span class=\"kt\">string</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"c1\">// Properly initialize the map</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">configSettings</span><span
  class=\"p\">[</span><span class=\"s\">&quot;server_port&quot;</span><span class=\"p\">]</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"s\">&quot;8080&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">configSettings</span><span
  class=\"p\">[</span><span class=\"s\">&quot;database_url&quot;</span><span class=\"p\">]</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"s\">&quot;localhost:5432&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Configuration settings initialized&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">getConfigSetting</span><span
  class=\"p\">(</span><span class=\"nx\">key</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">configSettings</span><span class=\"w\"> </span><span
  class=\"o\">==</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Configuration settings map is not initialized&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">value</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">exists</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">configSettings</span><span class=\"p\">[</span><span
  class=\"nx\">key</span><span class=\"p\">]</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"p\">!</span><span class=\"nx\">exists</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"s\">&quot;Setting
  not found&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">value</span><span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\nConfiguration settings initialized\nServer port: <span class=\"m\">8080</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above code, we declared the global map <code>configSettings</code> but didn't
  initialize it at that point, until we wanted to access the map. We initialize the
  map in the main function, this main function could be other specific parts of the
  code, and the global variable <code>configSettings</code> a map from another part
  of the code, by initializing it in the required scope, we prevent it from causing
  nil pointer access errors. We only initialize the map if it is <code>nil</code>
  i.e. it has not been initialized elsewhere in the code. This prevents overriding
  the map/flushing out the config set from other parts of the scope.</p>\n<h2>Pitfalls
  in access of un-initialized maps</h2>\n<p>But since it deals with pointers, it comes
  with its own pitfalls like nil pointers access when the map is not initialized.</p>\n<p>Let's
  take a look at an example, a real case where this might happen.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/url&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">vals</span><span
  class=\"w\"> </span><span class=\"nx\">url</span><span class=\"p\">.</span><span
  class=\"nx\">Values</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nx\">vals</span><span class=\"p\">.</span><span class=\"nx\">Add</span><span
  class=\"p\">(</span><span class=\"s\">&quot;foo&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;bar&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">vals</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>This
  will result in a runtime panic.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>nginx</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">$</span><span class=\"w\"> </span><span class=\"s\">go</span><span class=\"w\">
  </span><span class=\"s\">run</span><span class=\"w\"> </span><span class=\"s\">main.go</span><span
  class=\"w\"></span>\n<span class=\"s\">panic:</span><span class=\"w\"> </span><span
  class=\"s\">assignment</span><span class=\"w\"> </span><span class=\"s\">to</span><span
  class=\"w\"> </span><span class=\"s\">entry</span><span class=\"w\"> </span><span
  class=\"s\">in</span><span class=\"w\"> </span><span class=\"s\">nil</span><span
  class=\"w\"> </span><span class=\"s\">map</span><span class=\"w\"></span>\n\n<span
  class=\"s\">goroutine</span><span class=\"w\"> </span><span class=\"mi\">1</span><span
  class=\"w\"> </span><span class=\"s\">[running]:</span><span class=\"w\"></span>\n<span
  class=\"s\">net/url.Values.Add(...)</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"s\">/usr/local/go/src/net/url/url.go:902</span><span
  class=\"w\"></span>\n<span class=\"s\">main.main()</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"s\">/home/meet/code/playground/go/main.go:10</span><span
  class=\"w\"> </span><span class=\"s\">+0x2d</span><span class=\"w\"></span>\n<span
  class=\"s\">exit</span><span class=\"w\"> </span><span class=\"s\">status</span><span
  class=\"w\"> </span><span class=\"mi\">2</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>This
  is because the <a href=\"https://pkg.go.dev/net/url#Values\">url.Values</a> is a
  map of string and a list of string values. Since the underlying type is a map for
  <code>Values</code>, and in the example, we only have declared the variable <code>vals</code>
  with the type <code>url.Values</code>, it will point to a <code>nil</code> reference,
  hence the message on adding the value to the type. So, it is a good practice to
  use <code>make</code> while declaring or initializing a map data type. If you are
  not sure the underlying type is <code>map</code> then you could use <code>Type{}</code>
  to initialize an empty value of that type.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;net/url&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nx\">vals</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nb\">make</span><span class=\"p\">(</span><span
  class=\"nx\">url</span><span class=\"p\">.</span><span class=\"nx\">Values</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"c1\">// OR</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"c1\">// vals := url.Values{}</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nx\">vals</span><span class=\"p\">.</span><span class=\"nx\">Add</span><span
  class=\"p\">(</span><span class=\"s\">&quot;foo&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;bar&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">vals</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run urlvals.go\nmap<span class=\"o\">[</span>foo:<span class=\"o\">[</span>bar<span
  class=\"o\">]]</span>\n<span class=\"nv\">foo</span><span class=\"o\">=</span>bar\n</pre></div>\n\n</pre>\n\n<p>It
  is also recommended by the <a href=\"https://go.dev/blog/maps\">golang team</a>
  to use the make function while initializing a map. So, either use <code>make</code>
  for maps, slices, and channels, or initialize the empty value variable with <code>Type{}</code>.
  Both of them work similarly, but the latter is more generally applicable to structs
  as well.</p>\n<h2>Conclusion</h2>\n<p>Understanding the difference between declaring
  and initializing maps in Golang is essential for any developer, not just in golang,
  but in general. As we've explored, simply declaring a map variable without initializing
  it can lead to runtime errors, such as panics when attempting to access or modify
  a nil map. Initializing a map ensures that it is properly allocated in memory and
  ready for use, thereby avoiding these pitfalls.</p>\n<p>By following best practices—such
  as using the make function or initializing with Type{}—you can prevent common issues
  related to uninitialized maps. Always ensure that maps and slices are explicitly
  initialized before use to safeguard against unexpected nil pointer dereferences</p>\n<p>Thank
  you for reading this post, If you have any questions, feedback, and suggestions,
  feel free to drop them in the comments.</p>\n<p>Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/golang-safe-map-usage.png
long_description: 'This week, I was working on one of the API wrapper packages for
  golang, and that dealt with sending post requests with URL encoded values, setting
  cookies, and all the fun stuff. However, while I was constructing the body, I was
  using  In this post, '
now: 2025-05-01 18:11:33.312524
path: blog/posts/2024-08-31-Golang-Map-Safe-Usage.md
prevnext: null
slug: golang-safely-using-maps
status: published
tags:
- go
templateKey: blog-post
title: 'Safely using Maps in Golang: Differences in declaration and initialization'
today: 2025-05-01
---

## Introduction

This week, I was working on one of the API wrapper packages for golang, and that dealt with sending post requests with URL encoded values, setting cookies, and all the fun stuff. However, while I was constructing the body, I was using [url.Value](https://pkg.go.dev/net/url#Values) type to construct the body, and use that to add and set key-value pairs. However, I was getting a wired `nil` pointer reference error in some of the parts, I thought it was because of some of the variables I set manually. However, by debugging closer, I found out a common pitfall or bad practice of just declaring a type but initializing it and that causing `nil` reference errors.

In this post, I will cover, what are maps, how to create maps, and especially how to properly declare and initialize them. Create a proper distinction between the declaration and initialization of maps or any similar data type in golang.

## What is a map in Golang?

A [map](https://go.dev/src/runtime/map.go) or a hashmap in golang is a basic data type that allows us to store key-value pairs. Under the hood, it is a header map-like data structure that holds buckets, which are basically pointers to bucket arrays (contiguous memory). It has hash codes that store the actual key-value pairs, and pointers to new buckets if the current overflows with the number of keys. This is a really smart data structure that provides almost constant time access.

## How to create maps in Golang

To create a simple map in golang, you can take an example of a letter frequency counter using a map of string and integer. The map will store the letters as keys and their frequency as values.

```go
package main

import "fmt"

func main() {
    words := "hello how are you"
    letters := map[string]int{}

    for _, word := range words {
        wordCount[word]++
    }

    fmt.Println("Word counts:")
    for word, count := range wordCount {
        fmt.Printf("%s: %d\n", word, count)
    }
}
```

```bash
$ go run main.go

Word counts:
e: 2
 : 3
w: 1
r: 1
y: 1
u: 1
h: 2
l: 2
o: 3
a: 1
```

So, by initializing the map as `map[string]int{}` you will get an empty map. This can be then used to populate the keys and values, we iterate over the string, and for each character (rune) we cast that byte of character into the string and increment the value, the zero value for int is `0`, so by default if the key is not present, it will be zero, it is a bit of double-edged swords though, we never know the key is present in the map with the value `0` or the key is not present but the default value is `0`. For that, you need to check if the `key` exists in the map or not.

For further details, you can check out my [Golang Maps](https://www.meetgor.com/golang-maps/) post in detail.

## Difference between declaration and initialization

There is a difference in declaring and initializing any variable in a programming language and has to do a lot more with the implementation of the underlying type. In the case of primary data types like `int`, `string`, `float`, etc. there are default/zero values, so that is the same as the declaration and initialization of the variables. However, in the case of maps and slices, the declaration is just making sure the variable is available to the scope of the program, however for initialization is setting it to its default/zero value or the actual value that should be assigned.

So, declaration simply makes the variable available within the scope of the program. For maps and slices, declaring a variable without initialization sets it to `nil`, meaning it points to no allocated memory and cannot be used directly.

Whereas the `initialization` allocates memory and sets the variable to a usable state. For maps and slices, you need to explicitly initialize them using syntax like `myMap = make(map[keyType]valueType)` or `slice = []type{}`. Without this initialization, attempting to use the map or slice will lead to runtime errors, such as panics for accessing or modifying a nil map or slice.

Let's looks at the values of a map when it is declared/initialized/not initialized.

Imagine you're building a configuration manager that reads settings from a map. The map will be declared globally but initialized only when the configuration is loaded.

1. Declared but not initialized
    

The below code demonstrates a map access that is not initialized.

```go
package main

import (
	"fmt"
	"log"
)

// Global map to store configuration settings
var configSettings map[string]string // Declared but not initialized

func main() {
	// Attempt to get a configuration setting before initializing the map
	serverPort := getConfigSetting("server_port")
	fmt.Printf("Server port: %s\n", serverPort)
}

func getConfigSetting(key string) string {
	if configSettings == nil {
		log.Fatal("Configuration settings map is not initialized")
	}
	value, exists := configSettings[key]
	if !exists {
		return "Setting not found"
	}
	return value
}
```

```bash
$ go run main.go
Server port: Setting not found
```

2. Declared and Initialized at the same time
    

The below code demonstrates a map access that is initialized at the same time.

```go
package main

import (
	"fmt"
	"log"
)

// Global map to store configuration settings
var configSettings = map[string]string{
	"server_port":  "8080",
	"database_url": "localhost:5432",
}

func main() {
	serverPort := getConfigSetting("server_port")
	fmt.Printf("Server port: %s\n", serverPort)
}

func getConfigSetting(key string) string {
	value, exists := configSettings[key]
	if !exists {
		return "Setting not found"
	}
	return value
}
```

```bash
$ go run main.go
Server port: 8080
```

3. Declared and later initialized
    

The below code demonstrates a map access that is initialized later.

```go
package main

import (
	"fmt"
	"log"
)

// Global map to store configuration settings
var configSettings map[string]string // declared but not initialized

func main() {
	// Initialize configuration settings
	initializeConfigSettings()
    // if the function is not called, the map will be nil

	// Get a configuration setting safely
	serverPort := getConfigSetting("server_port")
	fmt.Printf("Server port: %s\n", serverPort)
}

func initializeConfigSettings() {
	if configSettings == nil {
		configSettings = make(map[string]string) // Properly initialize the map
		configSettings["server_port"] = "8080"
		configSettings["database_url"] = "localhost:5432"
		fmt.Println("Configuration settings initialized")
	}
}

func getConfigSetting(key string) string {
	if configSettings == nil {
		log.Fatal("Configuration settings map is not initialized")
	}
	value, exists := configSettings[key]
	if !exists {
		return "Setting not found"
	}
	return value
}
```

```bash
$ go run main.go
Configuration settings initialized
Server port: 8080
```


In the above code, we declared the global map `configSettings` but didn't initialize it at that point, until we wanted to access the map. We initialize the map in the main function, this main function could be other specific parts of the code, and the global variable `configSettings` a map from another part of the code, by initializing it in the required scope, we prevent it from causing nil pointer access errors. We only initialize the map if it is `nil` i.e. it has not been initialized elsewhere in the code. This prevents overriding the map/flushing out the config set from other parts of the scope.

## Pitfalls in access of un-initialized maps

But since it deals with pointers, it comes with its own pitfalls like nil pointers access when the map is not initialized.

Let's take a look at an example, a real case where this might happen.

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
        var vals url.Values
        vals.Add("foo", "bar")
        fmt.Println(vals)
}
```

This will result in a runtime panic.

```nginx
$ go run main.go
panic: assignment to entry in nil map

goroutine 1 [running]:
net/url.Values.Add(...)
        /usr/local/go/src/net/url/url.go:902
main.main()
        /home/meet/code/playground/go/main.go:10 +0x2d
exit status 2
```

This is because the [url.Values](https://pkg.go.dev/net/url#Values) is a map of string and a list of string values. Since the underlying type is a map for `Values`, and in the example, we only have declared the variable `vals` with the type `url.Values`, it will point to a `nil` reference, hence the message on adding the value to the type. So, it is a good practice to use `make` while declaring or initializing a map data type. If you are not sure the underlying type is `map` then you could use `Type{}` to initialize an empty value of that type.

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
        vals := make(url.Values)
        // OR
        // vals := url.Values{}
        vals.Add("foo", "bar")
        fmt.Println(vals)
}
```

```bash
$ go run urlvals.go
map[foo:[bar]]
foo=bar
```

It is also recommended by the [golang team](https://go.dev/blog/maps) to use the make function while initializing a map. So, either use `make` for maps, slices, and channels, or initialize the empty value variable with `Type{}`. Both of them work similarly, but the latter is more generally applicable to structs as well.

## Conclusion

Understanding the difference between declaring and initializing maps in Golang is essential for any developer, not just in golang, but in general. As we've explored, simply declaring a map variable without initializing it can lead to runtime errors, such as panics when attempting to access or modify a nil map. Initializing a map ensures that it is properly allocated in memory and ready for use, thereby avoiding these pitfalls.

By following best practices—such as using the make function or initializing with Type{}—you can prevent common issues related to uninitialized maps. Always ensure that maps and slices are explicitly initialized before use to safeguard against unexpected nil pointer dereferences

Thank you for reading this post, If you have any questions, feedback, and suggestions, feel free to drop them in the comments.

Happy Coding :)