---
article_html: "<p>I have been learning Golang for around 2 years now, and I have never
  paid attention to the sort package, can you believe that! Where was this package
  hiding?</p>\n<p>The <code>sort</code> package provides convenient methods for sorting
  slices and user-defined collections in Go. Here's a quick overview of what it can
  do:</p>\n<h2>Sorting Slices</h2>\n<p>To sort a slice of builtin types like ints,
  strings, etc, you can simply call <code>sort.Slice()</code>:</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"nx\">nums</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"kt\">int</span><span
  class=\"p\">{</span><span class=\"mi\">5</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"mi\">2</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"mi\">6</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"mi\">3</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"mi\">1</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"nx\">sort</span><span class=\"p\">.</span><span
  class=\"nx\">Slice</span><span class=\"p\">(</span><span class=\"nx\">nums</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"kd\">func</span><span
  class=\"p\">(</span><span class=\"nx\">i</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">j</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"kt\">bool</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">nums</span><span
  class=\"p\">[</span><span class=\"nx\">i</span><span class=\"p\">]</span><span class=\"w\">
  </span><span class=\"p\">&lt;</span><span class=\"w\"> </span><span class=\"nx\">nums</span><span
  class=\"p\">[</span><span class=\"nx\">j</span><span class=\"p\">]</span><span class=\"w\"></span>\n<span
  class=\"p\">})</span><span class=\"w\"></span>\n\n<span class=\"c1\">// or </span><span
  class=\"w\"></span>\n<span class=\"c1\">// sorts.Ints(nums)</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">nums</span><span class=\"p\">)</span><span
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
  go run main.go\n\n<span class=\"o\">[</span><span class=\"m\">1</span>, <span class=\"m\">2</span>,
  <span class=\"m\">3</span>, <span class=\"m\">5</span>, <span class=\"m\">6</span><span
  class=\"o\">]</span>\n</pre></div>\n\n</pre>\n\n<p>The sort is in-place, mutating
  the original slice. You can customize the sort order by providing a less(i, j int)
  bool function.</p>\n<h2>Sorting Struct Slices</h2>\n<p>To sort a slice of custom
  structs, add a Len(), Less(i, j int) bool, and Swap(i, j int) method:</p>\n<pre
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
  class=\"kd\">type</span><span class=\"w\"> </span><span class=\"nx\">Person</span><span
  class=\"w\"> </span><span class=\"kd\">struct</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nx\">Name</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">Age</span><span
  class=\"w\">  </span><span class=\"kt\">int</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"nx\">people</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"nx\">Person</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"p\">{</span><span class=\"s\">&quot;Bob&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"mi\">30</span><span
  class=\"p\">},</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"p\">{</span><span class=\"s\">&quot;John&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">20</span><span class=\"p\">},</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"p\">{</span><span
  class=\"s\">&quot;Alice&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"mi\">25</span><span class=\"p\">},</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"nx\">sort</span><span
  class=\"p\">.</span><span class=\"nx\">Slice</span><span class=\"p\">(</span><span
  class=\"nx\">people</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"kd\">func</span><span class=\"p\">(</span><span class=\"nx\">i</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">j</span><span class=\"w\">
  </span><span class=\"kt\">int</span><span class=\"p\">)</span><span class=\"w\">
  </span><span class=\"kt\">bool</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"k\">return</span><span
  class=\"w\"> </span><span class=\"nx\">people</span><span class=\"p\">[</span><span
  class=\"nx\">i</span><span class=\"p\">].</span><span class=\"nx\">Age</span><span
  class=\"w\"> </span><span class=\"p\">&lt;</span><span class=\"w\"> </span><span
  class=\"nx\">people</span><span class=\"p\">[</span><span class=\"nx\">j</span><span
  class=\"p\">].</span><span class=\"nx\">Age</span><span class=\"w\"></span>\n<span
  class=\"p\">})</span><span class=\"w\"> </span>\n\n<span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">person</span><span class=\"p\">)</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  go run main.go\n<span class=\"o\">[{</span><span class=\"s2\">&quot;Name&quot;</span>:<span
  class=\"s2\">&quot;Alice&quot;</span>,<span class=\"s2\">&quot;Age&quot;</span>:25<span
  class=\"o\">}</span>,<span class=\"o\">{</span><span class=\"s2\">&quot;Name&quot;</span>:<span
  class=\"s2\">&quot;Bob&quot;</span>,<span class=\"s2\">&quot;Age&quot;</span>:30<span
  class=\"o\">}</span>,<span class=\"o\">{</span><span class=\"s2\">&quot;Name&quot;</span>:<span
  class=\"s2\">&quot;John&quot;</span>,<span class=\"s2\">&quot;Age&quot;</span>:20<span
  class=\"o\">}]</span>\n</pre></div>\n\n</pre>\n\n<p>This will sort people by age.</p>\n<h2>Sorting
  Maps</h2>\n<p>Maps are inherently unordered in Go. We can't sort them, but we can
  iterate them in a sorted way. We need a separate slice of keys or values(whicher
  required), we will sort those keys/values and iterate over them in that order.:</p>\n<h3>Sort
  by Keys</h3>\n<p>To sort a map by keys, we can use the <code>sort.Strings()</code>
  function or any other sort function as per the data structure, there are functions
  like <a href=\"https://pkg.go.dev/sort#Ints\">sort.Ints</a>, <a href=\"https://pkg.go.dev/sort#Float64s\">sort.Float64</a>,
  or <a href=\"https://pkg.go.dev/sort#Slice\">sort.Slice</a>. We can create a new
  slice of keys from the map and then apply the sort on that newly created slice.
  After the slice of keys is created, we can iterate over it and access the map values
  in a order of sorted keys.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nx\">counts</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"kd\">map</span><span class=\"p\">[</span><span
  class=\"kt\">string</span><span class=\"p\">]</span><span class=\"kt\">int</span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"s\">&quot;hello&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"mi\">5</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"s\">&quot;world&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">2</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"s\">&quot;foo&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">3</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"nx\">keys</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nb\">make</span><span
  class=\"p\">([]</span><span class=\"kt\">string</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">0</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nb\">len</span><span class=\"p\">(</span><span class=\"nx\">counts</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n<span class=\"c1\">// extract keys
  </span><span class=\"w\"></span>\n<span class=\"k\">for</span><span class=\"w\">
  </span><span class=\"nx\">k</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">counts</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">keys</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nb\">append</span><span
  class=\"p\">(</span><span class=\"nx\">keys</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">k</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"> </span>\n\n<span class=\"c1\">// sort keys</span><span
  class=\"w\"></span>\n<span class=\"nx\">sort</span><span class=\"p\">.</span><span
  class=\"nx\">Strings</span><span class=\"p\">(</span><span class=\"nx\">keys</span><span
  class=\"p\">)</span><span class=\"w\"> </span>\n\n<span class=\"c1\">// iterate
  with the sorted keys slice</span><span class=\"w\"></span>\n<span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">k</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">keys</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">k</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">counts</span><span
  class=\"p\">[</span><span class=\"nx\">k</span><span class=\"p\">])</span><span
  class=\"w\"> </span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  go run main.go\n\nfoo <span class=\"m\">3</span>\nhello <span class=\"m\">5</span>\nworld
  <span class=\"m\">2</span>\n</pre></div>\n\n</pre>\n\n<p>This prints the map ordered
  by key. You can sort by value too.</p>\n<h3>Sort by Value</h3>\n<p>To iterate the
  map with sorting order of values, we can approach it in a similar way. We create
  a slice of keys. This time, we don't sort the keys, instead, we change the position
  of the slice of keys depending on the values. This changes the key order based on
  the sorted values in the map. By similarly iterating over the key slice, we iterate
  over the map in a sorted order of the values.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"nx\">counts</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"kd\">map</span><span class=\"p\">[</span><span
  class=\"kt\">string</span><span class=\"p\">]</span><span class=\"kt\">int</span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"s\">&quot;hello&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"mi\">5</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"s\">&quot;world&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">2</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"s\">&quot;foo&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">3</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"nx\">keys</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nb\">make</span><span
  class=\"p\">([]</span><span class=\"kt\">string</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">0</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nb\">len</span><span class=\"p\">(</span><span class=\"nx\">counts</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n<span class=\"c1\">// extract keys</span><span
  class=\"w\"></span>\n<span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">k</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">counts</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">keys</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nb\">append</span><span
  class=\"p\">(</span><span class=\"nx\">keys</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">k</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"c1\">// sort by value</span><span
  class=\"w\"></span>\n<span class=\"c1\">// i.e. change the order of key based on
  values sort order</span><span class=\"w\"></span>\n<span class=\"nx\">sort</span><span
  class=\"p\">.</span><span class=\"nx\">SliceStable</span><span class=\"p\">(</span><span
  class=\"nx\">keys</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"kd\">func</span><span class=\"p\">(</span><span class=\"nx\">i</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">j</span><span class=\"w\">
  </span><span class=\"kt\">int</span><span class=\"p\">)</span><span class=\"w\">
  </span><span class=\"kt\">bool</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">return</span><span
  class=\"w\"> </span><span class=\"nx\">counts</span><span class=\"p\">[</span><span
  class=\"nx\">keys</span><span class=\"p\">[</span><span class=\"nx\">i</span><span
  class=\"p\">]]</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"nx\">counts</span><span class=\"p\">[</span><span
  class=\"nx\">keys</span><span class=\"p\">[</span><span class=\"nx\">j</span><span
  class=\"p\">]]</span><span class=\"w\"></span>\n<span class=\"p\">})</span><span
  class=\"w\"></span>\n\n<span class=\"c1\">// iterate sorted</span><span class=\"w\"></span>\n<span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">_</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">k</span><span class=\"w\">
  </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">keys</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">k</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">counts</span><span class=\"p\">[</span><span class=\"nx\">k</span><span
  class=\"p\">])</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
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
  go run main.go\n\nworld <span class=\"m\">2</span>\nfoo <span class=\"m\">3</span>\nhello
  <span class=\"m\">5</span>\n</pre></div>\n\n</pre>\n\n<p>The <code>sort.SliceStable</code>
  function is used to sort the slice in place. It takes in a slice and the less function
  which is the actual logic for comparison in the values. This function sorts the
  key elements based on the values in the map, and thereby the key slice is shuffled
  by the sorted order of the values in the map.</p>\n<p>The sort package is super
  useful for quickly organizing Go data structures.\nThere are <a href=\"https://pkg.go.dev/sort#Find\">Find</a>
  which tries to return a index of the first element that satisfies a comparison condition
  with a flag or found or not with a boolean. There is also a <a href=\"https://pkg.go.dev/sort#Search\">Search</a>
  which is used specifically for searching elements in a sorted array, it also gives
  a first index of the occuring element. But that is a topic for another day!</p>\n<p>Thank
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
  \   </style>\n<pre><code>&lt;a class='prev' href='/vimscript-to-lua-keymapper'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Map
  Vimscript Keymaps to Lua with a single function&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a
  class='next' href='/golang-build-from-source-1-24-above'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Building
  Golang from Source v1.23 and Above&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot;
  height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
cover: ''
date: 2024-01-15
datetime: 2024-01-15 00:00:00+00:00
description: Understanding the fundamentals of the sort package in Golang. Sorting
  integers, slices, struct values, maps
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/tils/2024-01-15-go-sort-maps.md
html: "<p>I have been learning Golang for around 2 years now, and I have never paid
  attention to the sort package, can you believe that! Where was this package hiding?</p>\n<p>The
  <code>sort</code> package provides convenient methods for sorting slices and user-defined
  collections in Go. Here's a quick overview of what it can do:</p>\n<h2>Sorting Slices</h2>\n<p>To
  sort a slice of builtin types like ints, strings, etc, you can simply call <code>sort.Slice()</code>:</p>\n<pre
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
  class=\"nx\">nums</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"p\">[]</span><span class=\"kt\">int</span><span
  class=\"p\">{</span><span class=\"mi\">5</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"mi\">2</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"mi\">6</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"mi\">3</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"mi\">1</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"nx\">sort</span><span class=\"p\">.</span><span
  class=\"nx\">Slice</span><span class=\"p\">(</span><span class=\"nx\">nums</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"kd\">func</span><span
  class=\"p\">(</span><span class=\"nx\">i</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">j</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"kt\">bool</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">nums</span><span
  class=\"p\">[</span><span class=\"nx\">i</span><span class=\"p\">]</span><span class=\"w\">
  </span><span class=\"p\">&lt;</span><span class=\"w\"> </span><span class=\"nx\">nums</span><span
  class=\"p\">[</span><span class=\"nx\">j</span><span class=\"p\">]</span><span class=\"w\"></span>\n<span
  class=\"p\">})</span><span class=\"w\"></span>\n\n<span class=\"c1\">// or </span><span
  class=\"w\"></span>\n<span class=\"c1\">// sorts.Ints(nums)</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">nums</span><span class=\"p\">)</span><span
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
  go run main.go\n\n<span class=\"o\">[</span><span class=\"m\">1</span>, <span class=\"m\">2</span>,
  <span class=\"m\">3</span>, <span class=\"m\">5</span>, <span class=\"m\">6</span><span
  class=\"o\">]</span>\n</pre></div>\n\n</pre>\n\n<p>The sort is in-place, mutating
  the original slice. You can customize the sort order by providing a less(i, j int)
  bool function.</p>\n<h2>Sorting Struct Slices</h2>\n<p>To sort a slice of custom
  structs, add a Len(), Less(i, j int) bool, and Swap(i, j int) method:</p>\n<pre
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
  class=\"kd\">type</span><span class=\"w\"> </span><span class=\"nx\">Person</span><span
  class=\"w\"> </span><span class=\"kd\">struct</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nx\">Name</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">Age</span><span
  class=\"w\">  </span><span class=\"kt\">int</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"nx\">people</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"nx\">Person</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"p\">{</span><span class=\"s\">&quot;Bob&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"mi\">30</span><span
  class=\"p\">},</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"p\">{</span><span class=\"s\">&quot;John&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">20</span><span class=\"p\">},</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"p\">{</span><span
  class=\"s\">&quot;Alice&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"mi\">25</span><span class=\"p\">},</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"nx\">sort</span><span
  class=\"p\">.</span><span class=\"nx\">Slice</span><span class=\"p\">(</span><span
  class=\"nx\">people</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"kd\">func</span><span class=\"p\">(</span><span class=\"nx\">i</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">j</span><span class=\"w\">
  </span><span class=\"kt\">int</span><span class=\"p\">)</span><span class=\"w\">
  </span><span class=\"kt\">bool</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"k\">return</span><span
  class=\"w\"> </span><span class=\"nx\">people</span><span class=\"p\">[</span><span
  class=\"nx\">i</span><span class=\"p\">].</span><span class=\"nx\">Age</span><span
  class=\"w\"> </span><span class=\"p\">&lt;</span><span class=\"w\"> </span><span
  class=\"nx\">people</span><span class=\"p\">[</span><span class=\"nx\">j</span><span
  class=\"p\">].</span><span class=\"nx\">Age</span><span class=\"w\"></span>\n<span
  class=\"p\">})</span><span class=\"w\"> </span>\n\n<span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">person</span><span class=\"p\">)</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  go run main.go\n<span class=\"o\">[{</span><span class=\"s2\">&quot;Name&quot;</span>:<span
  class=\"s2\">&quot;Alice&quot;</span>,<span class=\"s2\">&quot;Age&quot;</span>:25<span
  class=\"o\">}</span>,<span class=\"o\">{</span><span class=\"s2\">&quot;Name&quot;</span>:<span
  class=\"s2\">&quot;Bob&quot;</span>,<span class=\"s2\">&quot;Age&quot;</span>:30<span
  class=\"o\">}</span>,<span class=\"o\">{</span><span class=\"s2\">&quot;Name&quot;</span>:<span
  class=\"s2\">&quot;John&quot;</span>,<span class=\"s2\">&quot;Age&quot;</span>:20<span
  class=\"o\">}]</span>\n</pre></div>\n\n</pre>\n\n<p>This will sort people by age.</p>\n<h2>Sorting
  Maps</h2>\n<p>Maps are inherently unordered in Go. We can't sort them, but we can
  iterate them in a sorted way. We need a separate slice of keys or values(whicher
  required), we will sort those keys/values and iterate over them in that order.:</p>\n<h3>Sort
  by Keys</h3>\n<p>To sort a map by keys, we can use the <code>sort.Strings()</code>
  function or any other sort function as per the data structure, there are functions
  like <a href=\"https://pkg.go.dev/sort#Ints\">sort.Ints</a>, <a href=\"https://pkg.go.dev/sort#Float64s\">sort.Float64</a>,
  or <a href=\"https://pkg.go.dev/sort#Slice\">sort.Slice</a>. We can create a new
  slice of keys from the map and then apply the sort on that newly created slice.
  After the slice of keys is created, we can iterate over it and access the map values
  in a order of sorted keys.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nx\">counts</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"kd\">map</span><span class=\"p\">[</span><span
  class=\"kt\">string</span><span class=\"p\">]</span><span class=\"kt\">int</span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"s\">&quot;hello&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"mi\">5</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"s\">&quot;world&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">2</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"s\">&quot;foo&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">3</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"nx\">keys</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nb\">make</span><span
  class=\"p\">([]</span><span class=\"kt\">string</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">0</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nb\">len</span><span class=\"p\">(</span><span class=\"nx\">counts</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n<span class=\"c1\">// extract keys
  </span><span class=\"w\"></span>\n<span class=\"k\">for</span><span class=\"w\">
  </span><span class=\"nx\">k</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">counts</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">keys</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nb\">append</span><span
  class=\"p\">(</span><span class=\"nx\">keys</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">k</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"> </span>\n\n<span class=\"c1\">// sort keys</span><span
  class=\"w\"></span>\n<span class=\"nx\">sort</span><span class=\"p\">.</span><span
  class=\"nx\">Strings</span><span class=\"p\">(</span><span class=\"nx\">keys</span><span
  class=\"p\">)</span><span class=\"w\"> </span>\n\n<span class=\"c1\">// iterate
  with the sorted keys slice</span><span class=\"w\"></span>\n<span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">k</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">keys</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">k</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">counts</span><span
  class=\"p\">[</span><span class=\"nx\">k</span><span class=\"p\">])</span><span
  class=\"w\"> </span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  go run main.go\n\nfoo <span class=\"m\">3</span>\nhello <span class=\"m\">5</span>\nworld
  <span class=\"m\">2</span>\n</pre></div>\n\n</pre>\n\n<p>This prints the map ordered
  by key. You can sort by value too.</p>\n<h3>Sort by Value</h3>\n<p>To iterate the
  map with sorting order of values, we can approach it in a similar way. We create
  a slice of keys. This time, we don't sort the keys, instead, we change the position
  of the slice of keys depending on the values. This changes the key order based on
  the sorted values in the map. By similarly iterating over the key slice, we iterate
  over the map in a sorted order of the values.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"nx\">counts</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"kd\">map</span><span class=\"p\">[</span><span
  class=\"kt\">string</span><span class=\"p\">]</span><span class=\"kt\">int</span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"s\">&quot;hello&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"mi\">5</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"s\">&quot;world&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">2</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"s\">&quot;foo&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">3</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"nx\">keys</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nb\">make</span><span
  class=\"p\">([]</span><span class=\"kt\">string</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">0</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nb\">len</span><span class=\"p\">(</span><span class=\"nx\">counts</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n<span class=\"c1\">// extract keys</span><span
  class=\"w\"></span>\n<span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">k</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">counts</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">keys</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nb\">append</span><span
  class=\"p\">(</span><span class=\"nx\">keys</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">k</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"c1\">// sort by value</span><span
  class=\"w\"></span>\n<span class=\"c1\">// i.e. change the order of key based on
  values sort order</span><span class=\"w\"></span>\n<span class=\"nx\">sort</span><span
  class=\"p\">.</span><span class=\"nx\">SliceStable</span><span class=\"p\">(</span><span
  class=\"nx\">keys</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"kd\">func</span><span class=\"p\">(</span><span class=\"nx\">i</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">j</span><span class=\"w\">
  </span><span class=\"kt\">int</span><span class=\"p\">)</span><span class=\"w\">
  </span><span class=\"kt\">bool</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">return</span><span
  class=\"w\"> </span><span class=\"nx\">counts</span><span class=\"p\">[</span><span
  class=\"nx\">keys</span><span class=\"p\">[</span><span class=\"nx\">i</span><span
  class=\"p\">]]</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"nx\">counts</span><span class=\"p\">[</span><span
  class=\"nx\">keys</span><span class=\"p\">[</span><span class=\"nx\">j</span><span
  class=\"p\">]]</span><span class=\"w\"></span>\n<span class=\"p\">})</span><span
  class=\"w\"></span>\n\n<span class=\"c1\">// iterate sorted</span><span class=\"w\"></span>\n<span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">_</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">k</span><span class=\"w\">
  </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">keys</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">k</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">counts</span><span class=\"p\">[</span><span class=\"nx\">k</span><span
  class=\"p\">])</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
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
  go run main.go\n\nworld <span class=\"m\">2</span>\nfoo <span class=\"m\">3</span>\nhello
  <span class=\"m\">5</span>\n</pre></div>\n\n</pre>\n\n<p>The <code>sort.SliceStable</code>
  function is used to sort the slice in place. It takes in a slice and the less function
  which is the actual logic for comparison in the values. This function sorts the
  key elements based on the values in the map, and thereby the key slice is shuffled
  by the sorted order of the values in the map.</p>\n<p>The sort package is super
  useful for quickly organizing Go data structures.\nThere are <a href=\"https://pkg.go.dev/sort#Find\">Find</a>
  which tries to return a index of the first element that satisfies a comparison condition
  with a flag or found or not with a boolean. There is also a <a href=\"https://pkg.go.dev/sort#Search\">Search</a>
  which is used specifically for searching elements in a sorted array, it also gives
  a first index of the occuring element. But that is a topic for another day!</p>\n<p>Thank
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
  \   </style>\n<pre><code>&lt;a class='prev' href='/vimscript-to-lua-keymapper'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Map
  Vimscript Keymaps to Lua with a single function&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a
  class='next' href='/golang-build-from-source-1-24-above'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Building
  Golang from Source v1.23 and Above&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot;
  height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
long_description: I have been learning Golang for around 2 years now, and I have never
  paid attention to the sort package, can you believe that The  To sort a slice of
  builtin types like ints, strings, etc, you can simply call  The sort is in-place,
  mutating the origi
now: 2025-05-01 18:11:33.315132
path: blog/tils/2024-01-15-go-sort-maps.md
slug: golang-sort-package-basic
status: published-til
tags:
- go
templateKey: til
title: 'Golang: Sort Package Introduction'
today: 2025-05-01
---

I have been learning Golang for around 2 years now, and I have never paid attention to the sort package, can you believe that! Where was this package hiding?

The `sort` package provides convenient methods for sorting slices and user-defined collections in Go. Here's a quick overview of what it can do:

## Sorting Slices

To sort a slice of builtin types like ints, strings, etc, you can simply call `sort.Slice()`:

```go
nums := []int{5, 2, 6, 3, 1}

sort.Slice(nums, func(i, j int) bool {
    return nums[i] < nums[j]
})

// or 
// sorts.Ints(nums)

fmt.Println(nums)
```

```bash
$ go run main.go

[1, 2, 3, 5, 6]
```
The sort is in-place, mutating the original slice. You can customize the sort order by providing a less(i, j int) bool function.

## Sorting Struct Slices

To sort a slice of custom structs, add a Len(), Less(i, j int) bool, and Swap(i, j int) method:

```go
type Person struct {
  Name string
  Age  int
}

people := []Person{
  {"Bob", 30},
  {"John", 20},
  {"Alice", 25},
}

sort.Slice(people, func(i, j int) bool {
  return people[i].Age < people[j].Age
}) 

fmt.Println(person)
```

```bash
$ go run main.go
[{"Name":"Alice","Age":25},{"Name":"Bob","Age":30},{"Name":"John","Age":20}]
```

This will sort people by age.

## Sorting Maps

Maps are inherently unordered in Go. We can't sort them, but we can iterate them in a sorted way. We need a separate slice of keys or values(whicher required), we will sort those keys/values and iterate over them in that order.:

### Sort by Keys

To sort a map by keys, we can use the `sort.Strings()` function or any other sort function as per the data structure, there are functions like [sort.Ints](https://pkg.go.dev/sort#Ints), [sort.Float64](https://pkg.go.dev/sort#Float64s), or [sort.Slice](https://pkg.go.dev/sort#Slice). We can create a new slice of keys from the map and then apply the sort on that newly created slice. After the slice of keys is created, we can iterate over it and access the map values in a order of sorted keys.

```go
counts := map[string]int{
  "hello": 5,
  "world": 2,
  "foo": 3,
}

keys := make([]string, 0, len(counts))
// extract keys 
for k := range counts {
  keys = append(keys, k)
} 

// sort keys
sort.Strings(keys) 

// iterate with the sorted keys slice
for _, k := range keys {
  fmt.Println(k, counts[k]) 
}
```

```bash
$ go run main.go

foo 3
hello 5
world 2
```

This prints the map ordered by key. You can sort by value too.

### Sort by Value

To iterate the map with sorting order of values, we can approach it in a similar way. We create a slice of keys. This time, we don't sort the keys, instead, we change the position of the slice of keys depending on the values. This changes the key order based on the sorted values in the map. By similarly iterating over the key slice, we iterate over the map in a sorted order of the values.

```go
counts := map[string]int{
  "hello": 5,
  "world": 2,
  "foo": 3,
}

keys := make([]string, 0, len(counts))
// extract keys
for k := range counts {
  keys = append(keys, k)
}

// sort by value
// i.e. change the order of key based on values sort order
sort.SliceStable(keys, func(i, j int) bool {
    return counts[keys[i]] < counts[keys[j]]
})

// iterate sorted
for _, k := range keys {
  fmt.Println(k, counts[k])
}
```

```bash
$ go run main.go

world 2
foo 3
hello 5
```

The `sort.SliceStable` function is used to sort the slice in place. It takes in a slice and the less function which is the actual logic for comparison in the values. This function sorts the key elements based on the values in the map, and thereby the key slice is shuffled by the sorted order of the values in the map.

The sort package is super useful for quickly organizing Go data structures.
There are [Find](https://pkg.go.dev/sort#Find) which tries to return a index of the first element that satisfies a comparison condition with a flag or found or not with a boolean. There is also a [Search](https://pkg.go.dev/sort#Search) which is used specifically for searching elements in a sorted array, it also gives a first index of the occuring element. But that is a topic for another day!

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
    
    <a class='prev' href='/vimscript-to-lua-keymapper'>
    
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M13.5 8.25L9.75 12L13.5 15.75" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"> </path>
        </svg>
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>prev</p>
            <p class='prevnext-title'>Map Vimscript Keymaps to Lua with a single function</p>
        </div>
    </a>
    
    <a class='next' href='/golang-build-from-source-1-24-above'>
    
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>next</p>
            <p class='prevnext-title'>Building Golang from Source v1.23 and Above</p>
        </div>
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.5 15.75L14.25 12L10.5 8.25" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path>
        </svg>
    </a>
  </div>