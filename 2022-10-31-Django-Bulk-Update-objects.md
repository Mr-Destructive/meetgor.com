---
article_html: "<p>Let's say, I have a lots of objects which I want to update with
  a particular field or fields. We can use the <a href=\"https://docs.djangoproject.com/en/4.1/ref/models/querysets/#bulk-update\">bulk_update</a>
  method with the model name.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># blog/models.py</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">django.db</span>
  <span class=\"kn\">import</span> <span class=\"n\">models</span>\n\n<span class=\"n\">ARTICLE_STATUS</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"p\">(</span><span
  class=\"s2\">&quot;PUBLISHED&quot;</span><span class=\"p\">,</span> <span class=\"s2\">&quot;Published&quot;</span><span
  class=\"p\">),</span>\n    <span class=\"p\">(</span><span class=\"s2\">&quot;DRAFT&quot;</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;Draft&quot;</span><span class=\"p\">),</span>\n<span
  class=\"p\">]</span>\n\n<span class=\"k\">class</span> <span class=\"nc\">Article</span><span
  class=\"p\">(</span><span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">Model</span><span class=\"p\">):</span>\n    <span class=\"n\">title</span>
  <span class=\"o\">=</span> <span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">CharField</span><span class=\"p\">(</span><span class=\"n\">max_length</span><span
  class=\"o\">=</span><span class=\"mi\">128</span><span class=\"p\">)</span>\n    <span
  class=\"n\">description</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">CharField</span><span class=\"p\">(</span><span
  class=\"n\">max_length</span><span class=\"o\">=</span><span class=\"mi\">512</span><span
  class=\"p\">)</span>\n    <span class=\"n\">content</span> <span class=\"o\">=</span>
  <span class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">TextField</span><span
  class=\"p\">()</span>\n    <span class=\"n\">status</span> <span class=\"o\">=</span>
  <span class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">CharField</span><span
  class=\"p\">(</span><span class=\"n\">max_length</span><span class=\"o\">=</span><span
  class=\"mi\">16</span><span class=\"p\">,</span> <span class=\"n\">choices</span><span
  class=\"o\">=</span><span class=\"n\">ARTICLE_STATUS</span><span class=\"p\">,</span>
  <span class=\"n\">default</span><span class=\"o\">=</span><span class=\"s2\">&quot;DRAFT&quot;</span><span
  class=\"p\">)</span>\n\n    <span class=\"k\">def</span> <span class=\"fm\">__str__</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"p\">):</span>\n
  \       <span class=\"k\">return</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">title</span>\n</pre></div>\n\n</pre>\n\n<p>Let's say we have a simple
  model <code>Article</code> with a few typical attributes like <code>title</code>,
  <code>description</code>, <code>content</code>, and <code>status</code>. We have
  the status as a choice field from two options as <code>Draft</code> and <code>Published</code>.
  It could be a boolean field, but that looks too gross for a article status.</p>\n<pre
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">from</span> <span class=\"nn\">blog.models</span> <span class=\"kn\">import</span>
  <span class=\"n\">Article</span>\n\n<span class=\"n\">articles</span> <span class=\"o\">=</span>
  <span class=\"n\">Article</span><span class=\"o\">.</span><span class=\"n\">objects</span><span
  class=\"o\">.</span><span class=\"n\">filter</span><span class=\"p\">(</span><span
  class=\"n\">status</span><span class=\"o\">=</span><span class=\"s2\">&quot;draft&quot;</span><span
  class=\"p\">)</span>\n\n<span class=\"k\">for</span> <span class=\"n\">i</span>
  <span class=\"ow\">in</span> <span class=\"nb\">range</span><span class=\"p\">(</span><span
  class=\"nb\">len</span><span class=\"p\">(</span><span class=\"n\">articles</span><span
  class=\"p\">)):</span>\n    <span class=\"n\">articles</span><span class=\"p\">[</span><span
  class=\"n\">i</span><span class=\"p\">]</span><span class=\"o\">.</span><span class=\"n\">status</span>
  <span class=\"o\">=</span> <span class=\"s2\">&quot;published&quot;</span>\n\n<span
  class=\"n\">Article</span><span class=\"o\">.</span><span class=\"n\">objects</span><span
  class=\"o\">.</span><span class=\"n\">bulk_update</span><span class=\"p\">(</span><span
  class=\"n\">articles</span><span class=\"p\">,</span> <span class=\"p\">[</span><span
  class=\"s2\">&quot;status&quot;</span><span class=\"p\">,])</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above code, the <code>Articles</code> model is filtered by the status of <code>draft</code>.
  We iterate over the <code>QuerySet</code> which will contain the objects of the
  articles, by setting the object's properties to the value we want to set. We are
  jsut setting the value of the property of the object for each object.</p>\n<p>This
  just makes a changes to the <code>QuerySet</code>, by using the <code>bulk_update</code>
  method, the two parameters required are the <code>QuerySet</code> and the list of
  <code>fields</code> which are to be updated. The function returns the number of
  objects/records updated.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"o\">&gt;&gt;&gt;</span> <span class=\"kn\">from</span> <span class=\"nn\">blog.models</span>
  <span class=\"kn\">import</span> <span class=\"n\">Article</span>\n<span class=\"o\">&gt;&gt;&gt;</span>
  <span class=\"n\">articles</span> <span class=\"o\">=</span> <span class=\"n\">Article</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">filter</span><span class=\"p\">(</span><span class=\"n\">status</span><span
  class=\"o\">=</span><span class=\"s2\">&quot;DRAFT&quot;</span><span class=\"p\">)</span>\n<span
  class=\"o\">&gt;&gt;&gt;</span> <span class=\"n\">articles</span>\n<span class=\"o\">&lt;</span><span
  class=\"n\">QuerySet</span> <span class=\"p\">[</span><span class=\"o\">&lt;</span><span
  class=\"n\">Article</span><span class=\"p\">:</span> <span class=\"n\">test</span>
  <span class=\"mi\">1</span><span class=\"o\">&gt;</span><span class=\"p\">,</span>
  <span class=\"o\">&lt;</span><span class=\"n\">Article</span><span class=\"p\">:</span>
  <span class=\"n\">test</span> <span class=\"mi\">3</span><span class=\"o\">&gt;</span><span
  class=\"p\">]</span><span class=\"o\">&gt;</span>\n\n<span class=\"o\">&gt;&gt;&gt;</span>
  <span class=\"k\">for</span> <span class=\"n\">i</span> <span class=\"ow\">in</span>
  <span class=\"nb\">range</span><span class=\"p\">(</span><span class=\"nb\">len</span><span
  class=\"p\">(</span><span class=\"n\">articles</span><span class=\"p\">)):</span>\n<span
  class=\"o\">...</span>     <span class=\"n\">articles</span><span class=\"p\">[</span><span
  class=\"n\">i</span><span class=\"p\">]</span><span class=\"o\">.</span><span class=\"n\">status</span>
  <span class=\"o\">=</span> <span class=\"s2\">&quot;PUBLISHED&quot;</span>\n<span
  class=\"o\">...</span>\n<span class=\"o\">&gt;&gt;&gt;</span> <span class=\"n\">articles</span>\n<span
  class=\"o\">&lt;</span><span class=\"n\">QuerySet</span> <span class=\"p\">[</span><span
  class=\"o\">&lt;</span><span class=\"n\">Article</span><span class=\"p\">:</span>
  <span class=\"n\">test</span> <span class=\"mi\">1</span><span class=\"o\">&gt;</span><span
  class=\"p\">,</span> <span class=\"o\">&lt;</span><span class=\"n\">Article</span><span
  class=\"p\">:</span> <span class=\"n\">test</span> <span class=\"mi\">3</span><span
  class=\"o\">&gt;</span><span class=\"p\">]</span><span class=\"o\">&gt;</span>\n<span
  class=\"o\">&gt;&gt;&gt;</span>\n\n<span class=\"o\">&gt;&gt;&gt;</span> <span class=\"n\">Article</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">bulk_update</span><span class=\"p\">(</span><span class=\"n\">articles</span><span
  class=\"p\">,</span> <span class=\"p\">[</span><span class=\"s1\">&#39;status&#39;</span><span
  class=\"p\">,])</span>\n<span class=\"mi\">2</span>\n\n<span class=\"o\">&gt;&gt;&gt;</span>
  <span class=\"n\">Article</span><span class=\"o\">.</span><span class=\"n\">obejcts</span><span
  class=\"o\">.</span><span class=\"n\">get</span><span class=\"p\">(</span><span
  class=\"n\">title</span><span class=\"o\">=</span><span class=\"s2\">&quot;test
  1&quot;</span><span class=\"p\">)</span><span class=\"o\">.</span><span class=\"n\">status</span>\n<span
  class=\"s1\">&#39;PUBLISHED&#39;</span>\n\n<span class=\"o\">&gt;&gt;&gt;</span>
  <span class=\"n\">Article</span><span class=\"o\">.</span><span class=\"n\">objects</span><span
  class=\"o\">.</span><span class=\"n\">filter</span><span class=\"p\">(</span><span
  class=\"n\">status</span><span class=\"o\">=</span><span class=\"s2\">&quot;DRAFT&quot;</span><span
  class=\"p\">)</span>\n<span class=\"o\">&lt;</span><span class=\"n\">QuerySet</span>
  <span class=\"p\">[]</span><span class=\"o\">&gt;</span>\n<span class=\"o\">&gt;&gt;&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>As,
  we can see here there were two obejcts <code>test 1</code> and <code>test 2</code>
  objects with the status as <code>Draft</code>. By iterating over the queryset and
  assigning the status of the object to published, the query set was changed and modified
  locally.\nBy using the <code>bulk_update</code> method, we parsed the queryset and
  the list of attributes to be updated into the function. This gives us the number
  of objects which were updated, in this case <code>2</code>. We then look into the
  article actual record in the database and it has indeed updated to the value we
  set in this operation.</p>\n<div class='prevnext'>\n    <style type='text/css'>\n
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
  \   </style>\n<pre><code>&lt;a class='prev' href='/git-ssh-multiple-accounts'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Adding
  SSH Keys for Multiple Accounts in Git&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a
  class='next' href='/golang-test-output-json'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Golang:
  Test Output JSON&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot;
  height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
cover: ''
date: 2022-10-31
datetime: 2022-10-31 00:00:00+00:00
description: Using bulk_update to update multiple objects in one go.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/tils/2022-10-31-Django-Bulk-Update-objects.md
html: "<p>Let's say, I have a lots of objects which I want to update with a particular
  field or fields. We can use the <a href=\"https://docs.djangoproject.com/en/4.1/ref/models/querysets/#bulk-update\">bulk_update</a>
  method with the model name.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># blog/models.py</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">django.db</span>
  <span class=\"kn\">import</span> <span class=\"n\">models</span>\n\n<span class=\"n\">ARTICLE_STATUS</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"p\">(</span><span
  class=\"s2\">&quot;PUBLISHED&quot;</span><span class=\"p\">,</span> <span class=\"s2\">&quot;Published&quot;</span><span
  class=\"p\">),</span>\n    <span class=\"p\">(</span><span class=\"s2\">&quot;DRAFT&quot;</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;Draft&quot;</span><span class=\"p\">),</span>\n<span
  class=\"p\">]</span>\n\n<span class=\"k\">class</span> <span class=\"nc\">Article</span><span
  class=\"p\">(</span><span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">Model</span><span class=\"p\">):</span>\n    <span class=\"n\">title</span>
  <span class=\"o\">=</span> <span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">CharField</span><span class=\"p\">(</span><span class=\"n\">max_length</span><span
  class=\"o\">=</span><span class=\"mi\">128</span><span class=\"p\">)</span>\n    <span
  class=\"n\">description</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">CharField</span><span class=\"p\">(</span><span
  class=\"n\">max_length</span><span class=\"o\">=</span><span class=\"mi\">512</span><span
  class=\"p\">)</span>\n    <span class=\"n\">content</span> <span class=\"o\">=</span>
  <span class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">TextField</span><span
  class=\"p\">()</span>\n    <span class=\"n\">status</span> <span class=\"o\">=</span>
  <span class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">CharField</span><span
  class=\"p\">(</span><span class=\"n\">max_length</span><span class=\"o\">=</span><span
  class=\"mi\">16</span><span class=\"p\">,</span> <span class=\"n\">choices</span><span
  class=\"o\">=</span><span class=\"n\">ARTICLE_STATUS</span><span class=\"p\">,</span>
  <span class=\"n\">default</span><span class=\"o\">=</span><span class=\"s2\">&quot;DRAFT&quot;</span><span
  class=\"p\">)</span>\n\n    <span class=\"k\">def</span> <span class=\"fm\">__str__</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"p\">):</span>\n
  \       <span class=\"k\">return</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">title</span>\n</pre></div>\n\n</pre>\n\n<p>Let's say we have a simple
  model <code>Article</code> with a few typical attributes like <code>title</code>,
  <code>description</code>, <code>content</code>, and <code>status</code>. We have
  the status as a choice field from two options as <code>Draft</code> and <code>Published</code>.
  It could be a boolean field, but that looks too gross for a article status.</p>\n<pre
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">from</span> <span class=\"nn\">blog.models</span> <span class=\"kn\">import</span>
  <span class=\"n\">Article</span>\n\n<span class=\"n\">articles</span> <span class=\"o\">=</span>
  <span class=\"n\">Article</span><span class=\"o\">.</span><span class=\"n\">objects</span><span
  class=\"o\">.</span><span class=\"n\">filter</span><span class=\"p\">(</span><span
  class=\"n\">status</span><span class=\"o\">=</span><span class=\"s2\">&quot;draft&quot;</span><span
  class=\"p\">)</span>\n\n<span class=\"k\">for</span> <span class=\"n\">i</span>
  <span class=\"ow\">in</span> <span class=\"nb\">range</span><span class=\"p\">(</span><span
  class=\"nb\">len</span><span class=\"p\">(</span><span class=\"n\">articles</span><span
  class=\"p\">)):</span>\n    <span class=\"n\">articles</span><span class=\"p\">[</span><span
  class=\"n\">i</span><span class=\"p\">]</span><span class=\"o\">.</span><span class=\"n\">status</span>
  <span class=\"o\">=</span> <span class=\"s2\">&quot;published&quot;</span>\n\n<span
  class=\"n\">Article</span><span class=\"o\">.</span><span class=\"n\">objects</span><span
  class=\"o\">.</span><span class=\"n\">bulk_update</span><span class=\"p\">(</span><span
  class=\"n\">articles</span><span class=\"p\">,</span> <span class=\"p\">[</span><span
  class=\"s2\">&quot;status&quot;</span><span class=\"p\">,])</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above code, the <code>Articles</code> model is filtered by the status of <code>draft</code>.
  We iterate over the <code>QuerySet</code> which will contain the objects of the
  articles, by setting the object's properties to the value we want to set. We are
  jsut setting the value of the property of the object for each object.</p>\n<p>This
  just makes a changes to the <code>QuerySet</code>, by using the <code>bulk_update</code>
  method, the two parameters required are the <code>QuerySet</code> and the list of
  <code>fields</code> which are to be updated. The function returns the number of
  objects/records updated.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"o\">&gt;&gt;&gt;</span> <span class=\"kn\">from</span> <span class=\"nn\">blog.models</span>
  <span class=\"kn\">import</span> <span class=\"n\">Article</span>\n<span class=\"o\">&gt;&gt;&gt;</span>
  <span class=\"n\">articles</span> <span class=\"o\">=</span> <span class=\"n\">Article</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">filter</span><span class=\"p\">(</span><span class=\"n\">status</span><span
  class=\"o\">=</span><span class=\"s2\">&quot;DRAFT&quot;</span><span class=\"p\">)</span>\n<span
  class=\"o\">&gt;&gt;&gt;</span> <span class=\"n\">articles</span>\n<span class=\"o\">&lt;</span><span
  class=\"n\">QuerySet</span> <span class=\"p\">[</span><span class=\"o\">&lt;</span><span
  class=\"n\">Article</span><span class=\"p\">:</span> <span class=\"n\">test</span>
  <span class=\"mi\">1</span><span class=\"o\">&gt;</span><span class=\"p\">,</span>
  <span class=\"o\">&lt;</span><span class=\"n\">Article</span><span class=\"p\">:</span>
  <span class=\"n\">test</span> <span class=\"mi\">3</span><span class=\"o\">&gt;</span><span
  class=\"p\">]</span><span class=\"o\">&gt;</span>\n\n<span class=\"o\">&gt;&gt;&gt;</span>
  <span class=\"k\">for</span> <span class=\"n\">i</span> <span class=\"ow\">in</span>
  <span class=\"nb\">range</span><span class=\"p\">(</span><span class=\"nb\">len</span><span
  class=\"p\">(</span><span class=\"n\">articles</span><span class=\"p\">)):</span>\n<span
  class=\"o\">...</span>     <span class=\"n\">articles</span><span class=\"p\">[</span><span
  class=\"n\">i</span><span class=\"p\">]</span><span class=\"o\">.</span><span class=\"n\">status</span>
  <span class=\"o\">=</span> <span class=\"s2\">&quot;PUBLISHED&quot;</span>\n<span
  class=\"o\">...</span>\n<span class=\"o\">&gt;&gt;&gt;</span> <span class=\"n\">articles</span>\n<span
  class=\"o\">&lt;</span><span class=\"n\">QuerySet</span> <span class=\"p\">[</span><span
  class=\"o\">&lt;</span><span class=\"n\">Article</span><span class=\"p\">:</span>
  <span class=\"n\">test</span> <span class=\"mi\">1</span><span class=\"o\">&gt;</span><span
  class=\"p\">,</span> <span class=\"o\">&lt;</span><span class=\"n\">Article</span><span
  class=\"p\">:</span> <span class=\"n\">test</span> <span class=\"mi\">3</span><span
  class=\"o\">&gt;</span><span class=\"p\">]</span><span class=\"o\">&gt;</span>\n<span
  class=\"o\">&gt;&gt;&gt;</span>\n\n<span class=\"o\">&gt;&gt;&gt;</span> <span class=\"n\">Article</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">bulk_update</span><span class=\"p\">(</span><span class=\"n\">articles</span><span
  class=\"p\">,</span> <span class=\"p\">[</span><span class=\"s1\">&#39;status&#39;</span><span
  class=\"p\">,])</span>\n<span class=\"mi\">2</span>\n\n<span class=\"o\">&gt;&gt;&gt;</span>
  <span class=\"n\">Article</span><span class=\"o\">.</span><span class=\"n\">obejcts</span><span
  class=\"o\">.</span><span class=\"n\">get</span><span class=\"p\">(</span><span
  class=\"n\">title</span><span class=\"o\">=</span><span class=\"s2\">&quot;test
  1&quot;</span><span class=\"p\">)</span><span class=\"o\">.</span><span class=\"n\">status</span>\n<span
  class=\"s1\">&#39;PUBLISHED&#39;</span>\n\n<span class=\"o\">&gt;&gt;&gt;</span>
  <span class=\"n\">Article</span><span class=\"o\">.</span><span class=\"n\">objects</span><span
  class=\"o\">.</span><span class=\"n\">filter</span><span class=\"p\">(</span><span
  class=\"n\">status</span><span class=\"o\">=</span><span class=\"s2\">&quot;DRAFT&quot;</span><span
  class=\"p\">)</span>\n<span class=\"o\">&lt;</span><span class=\"n\">QuerySet</span>
  <span class=\"p\">[]</span><span class=\"o\">&gt;</span>\n<span class=\"o\">&gt;&gt;&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>As,
  we can see here there were two obejcts <code>test 1</code> and <code>test 2</code>
  objects with the status as <code>Draft</code>. By iterating over the queryset and
  assigning the status of the object to published, the query set was changed and modified
  locally.\nBy using the <code>bulk_update</code> method, we parsed the queryset and
  the list of attributes to be updated into the function. This gives us the number
  of objects which were updated, in this case <code>2</code>. We then look into the
  article actual record in the database and it has indeed updated to the value we
  set in this operation.</p>\n<div class='prevnext'>\n    <style type='text/css'>\n
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
  \   </style>\n<pre><code>&lt;a class='prev' href='/git-ssh-multiple-accounts'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Adding
  SSH Keys for Multiple Accounts in Git&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a
  class='next' href='/golang-test-output-json'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Golang:
  Test Output JSON&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot;
  height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
long_description: 'Let Let In the above code, the  This just makes a changes to the  As,
  we can see here there were two obejcts '
now: 2025-05-01 18:11:33.315142
path: blog/tils/2022-10-31-Django-Bulk-Update-objects.md
slug: django-bulk-update-queryset
status: published-til
tags:
- django
- python
templateKey: til
title: Django Bulk Update QuerySet objects
today: 2025-05-01
---

Let's say, I have a lots of objects which I want to update with a particular field or fields. We can use the [bulk_update](https://docs.djangoproject.com/en/4.1/ref/models/querysets/#bulk-update) method with the model name.

```python
# blog/models.py

from django.db import models

ARTICLE_STATUS = [
    ("PUBLISHED", "Published"),
    ("DRAFT", "Draft"),
]

class Article(models.Model):
    title = models.CharField(max_length=128)
    description = models.CharField(max_length=512)
    content = models.TextField()
    status = models.CharField(max_length=16, choices=ARTICLE_STATUS, default="DRAFT")

    def __str__(self):
        return self.title

```

Let's say we have a simple model `Article` with a few typical attributes like `title`, `description`, `content`, and `status`. We have the status as a choice field from two options as `Draft` and `Published`. It could be a boolean field, but that looks too gross for a article status.


```python

from blog.models import Article

articles = Article.objects.filter(status="draft")

for i in range(len(articles)):
    articles[i].status = "published"

Article.objects.bulk_update(articles, ["status",])

```


In the above code, the `Articles` model is filtered by the status of `draft`. We iterate over the `QuerySet` which will contain the objects of the articles, by setting the object's properties to the value we want to set. We are jsut setting the value of the property of the object for each object.

This just makes a changes to the `QuerySet`, by using the `bulk_update` method, the two parameters required are the `QuerySet` and the list of `fields` which are to be updated. The function returns the number of objects/records updated.

```python
>>> from blog.models import Article
>>> articles = Article.objects.filter(status="DRAFT")
>>> articles
<QuerySet [<Article: test 1>, <Article: test 3>]>

>>> for i in range(len(articles)):
...     articles[i].status = "PUBLISHED"
...
>>> articles
<QuerySet [<Article: test 1>, <Article: test 3>]>
>>>

>>> Article.objects.bulk_update(articles, ['status',])
2

>>> Article.obejcts.get(title="test 1").status
'PUBLISHED'

>>> Article.objects.filter(status="DRAFT")
<QuerySet []>
>>>
```

As, we can see here there were two obejcts `test 1` and `test 2` objects with the status as `Draft`. By iterating over the queryset and assigning the status of the object to published, the query set was changed and modified locally.
By using the `bulk_update` method, we parsed the queryset and the list of attributes to be updated into the function. This gives us the number of objects which were updated, in this case `2`. We then look into the article actual record in the database and it has indeed updated to the value we set in this operation.
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
    
    <a class='prev' href='/git-ssh-multiple-accounts'>
    
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M13.5 8.25L9.75 12L13.5 15.75" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"> </path>
        </svg>
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>prev</p>
            <p class='prevnext-title'>Adding SSH Keys for Multiple Accounts in Git</p>
        </div>
    </a>
    
    <a class='next' href='/golang-test-output-json'>
    
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>next</p>
            <p class='prevnext-title'>Golang: Test Output JSON</p>
        </div>
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.5 15.75L14.25 12L10.5 8.25" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path>
        </svg>
    </a>
  </div>