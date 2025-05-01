---
article_html: "<h2>Populating a Python dict having a table-like structure to a CSV</h2>\n<p>Today,
  I want to share with you a neat trick I recently discovered for populating a CSV
  file with data in a table-like structure using Python.</p>\n<h3>Writing Key-Value
  Pairs to a CSV Row</h3>\n<p>Firstly, let's discuss the <code>write_key_value</code>
  function. This function allows us to write key-value pairs to a CSV row. It's particularly
  useful when dealing with metrics or data that can be represented as simple pairs.</p>\n<pre
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
  class=\"c1\"># Function to populate a CSV row with key-value pairs</span>\n<span
  class=\"k\">def</span> <span class=\"nf\">write_key_value</span><span class=\"p\">(</span><span
  class=\"n\">writer</span><span class=\"p\">,</span> <span class=\"n\">dictionary</span><span
  class=\"p\">):</span>\n    <span class=\"k\">for</span> <span class=\"n\">key</span><span
  class=\"p\">,</span> <span class=\"n\">value</span> <span class=\"ow\">in</span>
  <span class=\"n\">dictionary</span><span class=\"o\">.</span><span class=\"n\">items</span><span
  class=\"p\">():</span>\n        <span class=\"n\">writer</span><span class=\"o\">.</span><span
  class=\"n\">writerow</span><span class=\"p\">([</span><span class=\"n\">key</span><span
  class=\"p\">,</span> <span class=\"n\">value</span><span class=\"p\">])</span>\n</pre></div>\n\n</pre>\n\n<h3>Writing
  a Table-Like Structure to a CSV File</h3>\n<p>Now, let's dive into the <code>write_table</code>
  function, which handles more complex scenarios where the data follows a table-like
  structure. This function takes into account different types of metrics and adjusts
  the CSV table structure accordingly.</p>\n<p>Assuming you have a structure of the
  dictionary like:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"n\">data</span> <span class=\"o\">=</span> <span class=\"p\">{</span>\n
  \   <span class=\"s2\">&quot;Students&quot;</span><span class=\"p\">:</span> <span
  class=\"p\">{</span>\n        <span class=\"s2\">&quot;John Doe&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;course&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Mathematics&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;grade&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;A&quot;</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;attendance&quot;</span><span class=\"p\">:</span> <span class=\"mi\">95</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;assignments_completed&quot;</span><span
  class=\"p\">:</span> <span class=\"mi\">15</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;student_id&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;JD001&quot;</span>\n
  \       <span class=\"p\">},</span>\n        <span class=\"s2\">&quot;Alice Smith&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;course&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Physics&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;grade&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;B+&quot;</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;attendance&quot;</span><span class=\"p\">:</span> <span class=\"mi\">85</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;assignments_completed&quot;</span><span
  class=\"p\">:</span> <span class=\"mi\">12</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;student_id&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;AS002&quot;</span>\n
  \       <span class=\"p\">},</span>\n        <span class=\"s2\">&quot;Bob Johnson&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;course&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Computer Science&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;grade&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;A-&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;attendance&quot;</span><span class=\"p\">:</span>
  <span class=\"mi\">90</span><span class=\"p\">,</span>\n            <span class=\"s2\">&quot;assignments_completed&quot;</span><span
  class=\"p\">:</span> <span class=\"mi\">14</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;student_id&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;BJ003&quot;</span>\n
  \       <span class=\"p\">}</span>\n    <span class=\"p\">}</span>\n<span class=\"p\">}</span>\n</pre></div>\n\n</pre>\n\n<p>And
  you want to write it to a CSV file, like this:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>csv</div>\n<div class=\"highlight\"><pre><span></span>student,
  course, grade, attendance, assignments_completed, student_id\nJohn Doe, Mathematics,
  A, 95, 15, JD001\nAlice Smith, Physics, B+, 85, 12, AS002\nBob Johnson, Computer
  Science, A-, 90, 14, BJ003\n</pre></div>\n\n</pre>\n\n<p>We can create a function
  <code>write_table</code> that will take in the <code>dictionary</code> as the actual
  data. We want to store the keys of the inner dictionary to be the header/columns
  of the csv file. As we can see the keys of the inner dict i.e. the value for the
  key <code>John Doe</code> is a dict with the keys <code>course</code>, <code>grade</code>,
  <code>attendance</code>, etc. which remain the same for the all the keys in the
  dictionary.</p>\n<p>So, we can first create a <code>row_keys</code> variable to
  store the keys of the actual dictionary this will be the first column rows in the
  csv.</p>\n<p>Further we check if the <code>row_keys</code> is a dict and then we
  append it with the <code>index_key</code> which will be the first column in the
  csv. Since all the keys remain the same for the inner-dict, we can pick the first
  dict and create the <code>header</code> with the inner-dict keys.</p>\n<p>So, we
  can write the list <code>header</code> to the csv file.</p>\n<p>Then for each key
  in the <code>row_keys</code> we can create a list <code>row</code> with the key
  and the values of the inner-dict.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># Function to populate a CSV with a table-like structure</span>\n<span
  class=\"k\">def</span> <span class=\"nf\">write_table</span><span class=\"p\">(</span><span
  class=\"n\">writer</span><span class=\"p\">,</span> <span class=\"n\">dictionary</span><span
  class=\"p\">,</span> <span class=\"n\">index_key</span><span class=\"p\">):</span>\n\n
  \   <span class=\"n\">row_keys</span> <span class=\"o\">=</span> <span class=\"nb\">list</span><span
  class=\"p\">(</span><span class=\"n\">dictionary</span><span class=\"o\">.</span><span
  class=\"n\">keys</span><span class=\"p\">())</span>\n\n    <span class=\"k\">if</span>
  <span class=\"n\">row_keys</span> <span class=\"ow\">and</span> <span class=\"n\">data</span><span
  class=\"p\">[</span><span class=\"n\">row_keys</span><span class=\"p\">[</span><span
  class=\"mi\">0</span><span class=\"p\">]]</span> <span class=\"ow\">is</span> <span
  class=\"ow\">not</span> <span class=\"kc\">None</span><span class=\"p\">:</span>\n
  \       <span class=\"n\">headers</span> <span class=\"o\">=</span> <span class=\"p\">[</span><span
  class=\"n\">index_key</span><span class=\"p\">]</span> <span class=\"o\">+</span>
  <span class=\"nb\">list</span><span class=\"p\">(</span>\n            <span class=\"n\">dictionary</span><span
  class=\"p\">[</span><span class=\"n\">row_keys</span><span class=\"p\">[</span><span
  class=\"mi\">0</span><span class=\"p\">]]</span><span class=\"o\">.</span><span
  class=\"n\">keys</span><span class=\"p\">()</span>\n        <span class=\"p\">)</span>\n
  \   <span class=\"k\">else</span><span class=\"p\">:</span>\n        <span class=\"k\">return</span>\n
  \   <span class=\"n\">writer</span><span class=\"o\">.</span><span class=\"n\">writerow</span><span
  class=\"p\">(</span><span class=\"n\">headers</span><span class=\"p\">)</span>\n
  \   <span class=\"k\">for</span> <span class=\"n\">key</span> <span class=\"ow\">in</span>
  <span class=\"n\">row_keys</span><span class=\"p\">:</span>\n        <span class=\"n\">row</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span><span class=\"n\">key</span><span
  class=\"p\">]</span> <span class=\"o\">+</span> <span class=\"nb\">list</span><span
  class=\"p\">(</span><span class=\"n\">dictionary</span><span class=\"p\">[</span><span
  class=\"n\">key</span><span class=\"p\">]</span><span class=\"o\">.</span><span
  class=\"n\">values</span><span class=\"p\">())</span>\n        <span class=\"n\">writer</span><span
  class=\"o\">.</span><span class=\"n\">writerow</span><span class=\"p\">(</span><span
  class=\"n\">row</span><span class=\"p\">)</span>\n\n\n<span class=\"k\">with</span>
  <span class=\"nb\">open</span><span class=\"p\">(</span><span class=\"s1\">&#39;data.csv&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;w&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">newline</span><span class=\"o\">=</span><span class=\"s1\">&#39;&#39;</span><span
  class=\"p\">)</span> <span class=\"k\">as</span> <span class=\"n\">csvfile</span><span
  class=\"p\">:</span>\n    <span class=\"n\">writer</span> <span class=\"o\">=</span>
  <span class=\"n\">csv</span><span class=\"o\">.</span><span class=\"n\">writer</span><span
  class=\"p\">(</span><span class=\"n\">csvfile</span><span class=\"p\">)</span>\n
  \   <span class=\"k\">for</span> <span class=\"n\">key</span> <span class=\"ow\">in</span>
  <span class=\"n\">data</span><span class=\"p\">:</span>\n        <span class=\"n\">write_table</span><span
  class=\"p\">(</span><span class=\"n\">writer</span><span class=\"p\">,</span> <span
  class=\"n\">data</span><span class=\"p\">[</span><span class=\"n\">key</span><span
  class=\"p\">],</span> <span class=\"n\">key</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<h3>Example
  Usage</h3>\n<p>To illustrate how these functions can be used, let's consider a scenario
  where we have various types of metrics to populate into a CSV file. We handle key-value
  paired metrics separately and then populate the CSV with table-like metrics.</p>\n<pre
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
  class=\"kn\">import</span> <span class=\"nn\">csv</span>\n\n<span class=\"n\">data</span>
  <span class=\"o\">=</span> <span class=\"p\">{</span>\n    <span class=\"s2\">&quot;Students&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n        <span class=\"s2\">&quot;John
  Doe&quot;</span><span class=\"p\">:</span> <span class=\"p\">{</span>\n            <span
  class=\"s2\">&quot;course&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;Mathematics&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;grade&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;A&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;attendance&quot;</span><span class=\"p\">:</span>
  <span class=\"mi\">95</span><span class=\"p\">,</span>\n            <span class=\"s2\">&quot;assignments_completed&quot;</span><span
  class=\"p\">:</span> <span class=\"mi\">15</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;student_id&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;JD001&quot;</span>\n
  \       <span class=\"p\">},</span>\n        <span class=\"s2\">&quot;Alice Smith&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;course&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Physics&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;grade&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;B+&quot;</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;attendance&quot;</span><span class=\"p\">:</span> <span class=\"mi\">85</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;assignments_completed&quot;</span><span
  class=\"p\">:</span> <span class=\"mi\">12</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;student_id&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;AS002&quot;</span>\n
  \       <span class=\"p\">},</span>\n        <span class=\"s2\">&quot;Bob Johnson&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;course&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Computer Science&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;grade&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;A-&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;attendance&quot;</span><span class=\"p\">:</span>
  <span class=\"mi\">90</span><span class=\"p\">,</span>\n            <span class=\"s2\">&quot;assignments_completed&quot;</span><span
  class=\"p\">:</span> <span class=\"mi\">14</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;student_id&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;BJ003&quot;</span>\n
  \       <span class=\"p\">}</span>\n    <span class=\"p\">},</span>\n    <span class=\"s2\">&quot;Countries&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n        <span class=\"s2\">&quot;USA&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;capital&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Washington, D.C.&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;population&quot;</span><span
  class=\"p\">:</span> <span class=\"mi\">331000000</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;area_sq_km&quot;</span><span class=\"p\">:</span>
  <span class=\"mi\">9833517</span><span class=\"p\">,</span>\n            <span class=\"s2\">&quot;official_languages&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">[</span><span class=\"s2\">&quot;English&quot;</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;Spanish&quot;</span><span class=\"p\">],</span>\n
  \           <span class=\"s2\">&quot;currency&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;United States Dollar (USD)&quot;</span>\n        <span
  class=\"p\">},</span>\n        <span class=\"s2\">&quot;India&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;capital&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;New Delhi&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;population&quot;</span><span class=\"p\">:</span>
  <span class=\"mi\">1380004385</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;area_sq_km&quot;</span><span class=\"p\">:</span> <span class=\"mi\">3287263</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;official_languages&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">[</span><span class=\"s2\">&quot;Hindi&quot;</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;English&quot;</span><span class=\"p\">],</span>\n
  \           <span class=\"s2\">&quot;currency&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;Indian Rupee (INR)&quot;</span>\n        <span class=\"p\">},</span>\n
  \       <span class=\"s2\">&quot;Brazil&quot;</span><span class=\"p\">:</span> <span
  class=\"p\">{</span>\n            <span class=\"s2\">&quot;capital&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Brasília&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;population&quot;</span><span class=\"p\">:</span>
  <span class=\"mi\">212559417</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;area_sq_km&quot;</span><span class=\"p\">:</span> <span class=\"mi\">8515770</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;official_languages&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">[</span><span class=\"s2\">&quot;Portuguese&quot;</span><span
  class=\"p\">],</span>\n            <span class=\"s2\">&quot;currency&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Brazilian Real (BRL)&quot;</span>\n
  \       <span class=\"p\">}</span>\n    <span class=\"p\">}</span>\n<span class=\"p\">}</span>\n\n<span
  class=\"k\">def</span> <span class=\"nf\">populate_table</span><span class=\"p\">(</span><span
  class=\"n\">writer</span><span class=\"p\">,</span> <span class=\"n\">data</span><span
  class=\"p\">,</span> <span class=\"n\">index_key</span><span class=\"p\">):</span>\n
  \   <span class=\"n\">row_keys</span> <span class=\"o\">=</span> <span class=\"nb\">list</span><span
  class=\"p\">(</span><span class=\"n\">data</span><span class=\"o\">.</span><span
  class=\"n\">keys</span><span class=\"p\">())</span>\n    <span class=\"k\">if</span>
  <span class=\"n\">row_keys</span> <span class=\"ow\">and</span> <span class=\"n\">data</span><span
  class=\"p\">[</span><span class=\"n\">row_keys</span><span class=\"p\">[</span><span
  class=\"mi\">0</span><span class=\"p\">]]</span> <span class=\"ow\">is</span> <span
  class=\"ow\">not</span> <span class=\"kc\">None</span><span class=\"p\">:</span>\n
  \       <span class=\"n\">headers</span> <span class=\"o\">=</span> <span class=\"p\">[</span><span
  class=\"n\">index_key</span><span class=\"p\">]</span> <span class=\"o\">+</span>
  <span class=\"nb\">list</span><span class=\"p\">(</span><span class=\"n\">data</span><span
  class=\"p\">[</span><span class=\"n\">row_keys</span><span class=\"p\">[</span><span
  class=\"mi\">0</span><span class=\"p\">]]</span><span class=\"o\">.</span><span
  class=\"n\">keys</span><span class=\"p\">())</span>\n    <span class=\"k\">else</span><span
  class=\"p\">:</span>\n        <span class=\"k\">return</span>\n    <span class=\"n\">writer</span><span
  class=\"o\">.</span><span class=\"n\">writerow</span><span class=\"p\">(</span><span
  class=\"n\">headers</span><span class=\"p\">)</span>\n    <span class=\"k\">for</span>
  <span class=\"n\">key</span> <span class=\"ow\">in</span> <span class=\"n\">row_keys</span><span
  class=\"p\">:</span>\n        <span class=\"n\">row</span> <span class=\"o\">=</span>
  <span class=\"p\">[</span><span class=\"n\">key</span><span class=\"p\">]</span>
  <span class=\"o\">+</span> <span class=\"nb\">list</span><span class=\"p\">(</span><span
  class=\"n\">data</span><span class=\"p\">[</span><span class=\"n\">key</span><span
  class=\"p\">]</span><span class=\"o\">.</span><span class=\"n\">values</span><span
  class=\"p\">())</span>\n        <span class=\"n\">writer</span><span class=\"o\">.</span><span
  class=\"n\">writerow</span><span class=\"p\">(</span><span class=\"n\">row</span><span
  class=\"p\">)</span>\n\n\n<span class=\"k\">with</span> <span class=\"nb\">open</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;data.csv&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;w&#39;</span><span class=\"p\">,</span> <span class=\"n\">newline</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;&#39;</span><span class=\"p\">)</span>
  <span class=\"k\">as</span> <span class=\"n\">csvfile</span><span class=\"p\">:</span>\n
  \   <span class=\"n\">writer</span> <span class=\"o\">=</span> <span class=\"n\">csv</span><span
  class=\"o\">.</span><span class=\"n\">writer</span><span class=\"p\">(</span><span
  class=\"n\">csvfile</span><span class=\"p\">)</span>\n    <span class=\"k\">for</span>
  <span class=\"n\">key</span> <span class=\"ow\">in</span> <span class=\"n\">data</span><span
  class=\"p\">:</span>\n        <span class=\"n\">populate_table</span><span class=\"p\">(</span><span
  class=\"n\">writer</span><span class=\"p\">,</span> <span class=\"n\">data</span><span
  class=\"p\">[</span><span class=\"n\">key</span><span class=\"p\">],</span> <span
  class=\"n\">key</span><span class=\"p\">)</span>\n        <span class=\"n\">writer</span><span
  class=\"o\">.</span><span class=\"n\">writerow</span><span class=\"p\">([])</span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n<div class='language-bar'>csv</div>\n<div class=\"highlight\"><pre><span></span>Students,course,grade,attendance,assignments_completed,student_id\nJohn
  Doe,Mathematics,A,95,15,JD001\nAlice Smith,Physics,B+,85,12,AS002\nBob Johnson,Computer
  Science,A-,90,14,BJ003\n\nCountries,capital,population,area_sq_km,official_languages,currency\nUSA,&quot;Washington,
  D.C.&quot;,331000000,9833517,&quot;[&#39;English&#39;, &#39;Spanish&#39;]&quot;,United
  States Dollar (USD)\nIndia,New Delhi,1380004385,3287263,&quot;[&#39;Hindi&#39;,
  &#39;English&#39;]&quot;,Indian Rupee (INR)\nBrazil,Brasília,212559417,8515770,[&#39;Portuguese&#39;],Brazilian
  Real (BRL)\n</pre></div>\n\n</pre>\n\n<p>Thank you, Happy Coding :)</p>\n<div class='prevnext'>\n
  \   <style type='text/css'>\n    :root {\n      --prevnext-color-text: #eefbfe;\n
  \     --prevnext-color-angle: #ff6600;\n      --prevnext-subtitle-brightness: 3;\n
  \   }\n    [data-theme=\"light\"] {\n      --prevnext-color-text: #1f2022;\n      --prevnext-color-angle:
  #ffeb00;\n      --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"dark\"]
  {\n      --prevnext-color-text: #eefbfe;\n      --prevnext-color-angle: #ff6600;\n
  \     --prevnext-subtitle-brightness: 3;\n    }\n    .prevnext {\n      display:
  flex;\n      flex-direction: row;\n      justify-content: space-around;\n      align-items:
  flex-start;\n    }\n    .prevnext a {\n      display: flex;\n      align-items:
  center;\n      width: 100%;\n      text-decoration: none;\n    }\n    a.next {\n
  \     justify-content: flex-end;\n    }\n    .prevnext a:hover {\n      background:
  #00000006;\n    }\n    .prevnext-subtitle {\n      color: var(--prevnext-color-text);\n
  \     filter: brightness(var(--prevnext-subtitle-brightness));\n      font-size:
  .8rem;\n    }\n    .prevnext-title {\n      color: var(--prevnext-color-text);\n
  \     font-size: 1rem;\n    }\n    .prevnext-text {\n      max-width: 30vw;\n    }\n
  \   </style>\n<pre><code>&lt;a class='prev' href='/vim-python-black-autoformat'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Autoformat
  Python file with Black after saving in Vim&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a
  class='next' href='/python-search-replace-file'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Python:
  Search and Replace in File&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot;
  height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
cover: ''
date: 2024-03-20
datetime: 2024-03-20 00:00:00+00:00
description: Exploring how to write python dict/key-value pairs and a table-like structure
  to a CSV file.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/tils/2024-03-20-Python-dict-to-CSV-table.md
html: "<h2>Populating a Python dict having a table-like structure to a CSV</h2>\n<p>Today,
  I want to share with you a neat trick I recently discovered for populating a CSV
  file with data in a table-like structure using Python.</p>\n<h3>Writing Key-Value
  Pairs to a CSV Row</h3>\n<p>Firstly, let's discuss the <code>write_key_value</code>
  function. This function allows us to write key-value pairs to a CSV row. It's particularly
  useful when dealing with metrics or data that can be represented as simple pairs.</p>\n<pre
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
  class=\"c1\"># Function to populate a CSV row with key-value pairs</span>\n<span
  class=\"k\">def</span> <span class=\"nf\">write_key_value</span><span class=\"p\">(</span><span
  class=\"n\">writer</span><span class=\"p\">,</span> <span class=\"n\">dictionary</span><span
  class=\"p\">):</span>\n    <span class=\"k\">for</span> <span class=\"n\">key</span><span
  class=\"p\">,</span> <span class=\"n\">value</span> <span class=\"ow\">in</span>
  <span class=\"n\">dictionary</span><span class=\"o\">.</span><span class=\"n\">items</span><span
  class=\"p\">():</span>\n        <span class=\"n\">writer</span><span class=\"o\">.</span><span
  class=\"n\">writerow</span><span class=\"p\">([</span><span class=\"n\">key</span><span
  class=\"p\">,</span> <span class=\"n\">value</span><span class=\"p\">])</span>\n</pre></div>\n\n</pre>\n\n<h3>Writing
  a Table-Like Structure to a CSV File</h3>\n<p>Now, let's dive into the <code>write_table</code>
  function, which handles more complex scenarios where the data follows a table-like
  structure. This function takes into account different types of metrics and adjusts
  the CSV table structure accordingly.</p>\n<p>Assuming you have a structure of the
  dictionary like:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"n\">data</span> <span class=\"o\">=</span> <span class=\"p\">{</span>\n
  \   <span class=\"s2\">&quot;Students&quot;</span><span class=\"p\">:</span> <span
  class=\"p\">{</span>\n        <span class=\"s2\">&quot;John Doe&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;course&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Mathematics&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;grade&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;A&quot;</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;attendance&quot;</span><span class=\"p\">:</span> <span class=\"mi\">95</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;assignments_completed&quot;</span><span
  class=\"p\">:</span> <span class=\"mi\">15</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;student_id&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;JD001&quot;</span>\n
  \       <span class=\"p\">},</span>\n        <span class=\"s2\">&quot;Alice Smith&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;course&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Physics&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;grade&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;B+&quot;</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;attendance&quot;</span><span class=\"p\">:</span> <span class=\"mi\">85</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;assignments_completed&quot;</span><span
  class=\"p\">:</span> <span class=\"mi\">12</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;student_id&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;AS002&quot;</span>\n
  \       <span class=\"p\">},</span>\n        <span class=\"s2\">&quot;Bob Johnson&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;course&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Computer Science&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;grade&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;A-&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;attendance&quot;</span><span class=\"p\">:</span>
  <span class=\"mi\">90</span><span class=\"p\">,</span>\n            <span class=\"s2\">&quot;assignments_completed&quot;</span><span
  class=\"p\">:</span> <span class=\"mi\">14</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;student_id&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;BJ003&quot;</span>\n
  \       <span class=\"p\">}</span>\n    <span class=\"p\">}</span>\n<span class=\"p\">}</span>\n</pre></div>\n\n</pre>\n\n<p>And
  you want to write it to a CSV file, like this:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>csv</div>\n<div class=\"highlight\"><pre><span></span>student,
  course, grade, attendance, assignments_completed, student_id\nJohn Doe, Mathematics,
  A, 95, 15, JD001\nAlice Smith, Physics, B+, 85, 12, AS002\nBob Johnson, Computer
  Science, A-, 90, 14, BJ003\n</pre></div>\n\n</pre>\n\n<p>We can create a function
  <code>write_table</code> that will take in the <code>dictionary</code> as the actual
  data. We want to store the keys of the inner dictionary to be the header/columns
  of the csv file. As we can see the keys of the inner dict i.e. the value for the
  key <code>John Doe</code> is a dict with the keys <code>course</code>, <code>grade</code>,
  <code>attendance</code>, etc. which remain the same for the all the keys in the
  dictionary.</p>\n<p>So, we can first create a <code>row_keys</code> variable to
  store the keys of the actual dictionary this will be the first column rows in the
  csv.</p>\n<p>Further we check if the <code>row_keys</code> is a dict and then we
  append it with the <code>index_key</code> which will be the first column in the
  csv. Since all the keys remain the same for the inner-dict, we can pick the first
  dict and create the <code>header</code> with the inner-dict keys.</p>\n<p>So, we
  can write the list <code>header</code> to the csv file.</p>\n<p>Then for each key
  in the <code>row_keys</code> we can create a list <code>row</code> with the key
  and the values of the inner-dict.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># Function to populate a CSV with a table-like structure</span>\n<span
  class=\"k\">def</span> <span class=\"nf\">write_table</span><span class=\"p\">(</span><span
  class=\"n\">writer</span><span class=\"p\">,</span> <span class=\"n\">dictionary</span><span
  class=\"p\">,</span> <span class=\"n\">index_key</span><span class=\"p\">):</span>\n\n
  \   <span class=\"n\">row_keys</span> <span class=\"o\">=</span> <span class=\"nb\">list</span><span
  class=\"p\">(</span><span class=\"n\">dictionary</span><span class=\"o\">.</span><span
  class=\"n\">keys</span><span class=\"p\">())</span>\n\n    <span class=\"k\">if</span>
  <span class=\"n\">row_keys</span> <span class=\"ow\">and</span> <span class=\"n\">data</span><span
  class=\"p\">[</span><span class=\"n\">row_keys</span><span class=\"p\">[</span><span
  class=\"mi\">0</span><span class=\"p\">]]</span> <span class=\"ow\">is</span> <span
  class=\"ow\">not</span> <span class=\"kc\">None</span><span class=\"p\">:</span>\n
  \       <span class=\"n\">headers</span> <span class=\"o\">=</span> <span class=\"p\">[</span><span
  class=\"n\">index_key</span><span class=\"p\">]</span> <span class=\"o\">+</span>
  <span class=\"nb\">list</span><span class=\"p\">(</span>\n            <span class=\"n\">dictionary</span><span
  class=\"p\">[</span><span class=\"n\">row_keys</span><span class=\"p\">[</span><span
  class=\"mi\">0</span><span class=\"p\">]]</span><span class=\"o\">.</span><span
  class=\"n\">keys</span><span class=\"p\">()</span>\n        <span class=\"p\">)</span>\n
  \   <span class=\"k\">else</span><span class=\"p\">:</span>\n        <span class=\"k\">return</span>\n
  \   <span class=\"n\">writer</span><span class=\"o\">.</span><span class=\"n\">writerow</span><span
  class=\"p\">(</span><span class=\"n\">headers</span><span class=\"p\">)</span>\n
  \   <span class=\"k\">for</span> <span class=\"n\">key</span> <span class=\"ow\">in</span>
  <span class=\"n\">row_keys</span><span class=\"p\">:</span>\n        <span class=\"n\">row</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span><span class=\"n\">key</span><span
  class=\"p\">]</span> <span class=\"o\">+</span> <span class=\"nb\">list</span><span
  class=\"p\">(</span><span class=\"n\">dictionary</span><span class=\"p\">[</span><span
  class=\"n\">key</span><span class=\"p\">]</span><span class=\"o\">.</span><span
  class=\"n\">values</span><span class=\"p\">())</span>\n        <span class=\"n\">writer</span><span
  class=\"o\">.</span><span class=\"n\">writerow</span><span class=\"p\">(</span><span
  class=\"n\">row</span><span class=\"p\">)</span>\n\n\n<span class=\"k\">with</span>
  <span class=\"nb\">open</span><span class=\"p\">(</span><span class=\"s1\">&#39;data.csv&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;w&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">newline</span><span class=\"o\">=</span><span class=\"s1\">&#39;&#39;</span><span
  class=\"p\">)</span> <span class=\"k\">as</span> <span class=\"n\">csvfile</span><span
  class=\"p\">:</span>\n    <span class=\"n\">writer</span> <span class=\"o\">=</span>
  <span class=\"n\">csv</span><span class=\"o\">.</span><span class=\"n\">writer</span><span
  class=\"p\">(</span><span class=\"n\">csvfile</span><span class=\"p\">)</span>\n
  \   <span class=\"k\">for</span> <span class=\"n\">key</span> <span class=\"ow\">in</span>
  <span class=\"n\">data</span><span class=\"p\">:</span>\n        <span class=\"n\">write_table</span><span
  class=\"p\">(</span><span class=\"n\">writer</span><span class=\"p\">,</span> <span
  class=\"n\">data</span><span class=\"p\">[</span><span class=\"n\">key</span><span
  class=\"p\">],</span> <span class=\"n\">key</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<h3>Example
  Usage</h3>\n<p>To illustrate how these functions can be used, let's consider a scenario
  where we have various types of metrics to populate into a CSV file. We handle key-value
  paired metrics separately and then populate the CSV with table-like metrics.</p>\n<pre
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
  class=\"kn\">import</span> <span class=\"nn\">csv</span>\n\n<span class=\"n\">data</span>
  <span class=\"o\">=</span> <span class=\"p\">{</span>\n    <span class=\"s2\">&quot;Students&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n        <span class=\"s2\">&quot;John
  Doe&quot;</span><span class=\"p\">:</span> <span class=\"p\">{</span>\n            <span
  class=\"s2\">&quot;course&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;Mathematics&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;grade&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;A&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;attendance&quot;</span><span class=\"p\">:</span>
  <span class=\"mi\">95</span><span class=\"p\">,</span>\n            <span class=\"s2\">&quot;assignments_completed&quot;</span><span
  class=\"p\">:</span> <span class=\"mi\">15</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;student_id&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;JD001&quot;</span>\n
  \       <span class=\"p\">},</span>\n        <span class=\"s2\">&quot;Alice Smith&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;course&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Physics&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;grade&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;B+&quot;</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;attendance&quot;</span><span class=\"p\">:</span> <span class=\"mi\">85</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;assignments_completed&quot;</span><span
  class=\"p\">:</span> <span class=\"mi\">12</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;student_id&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;AS002&quot;</span>\n
  \       <span class=\"p\">},</span>\n        <span class=\"s2\">&quot;Bob Johnson&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;course&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Computer Science&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;grade&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;A-&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;attendance&quot;</span><span class=\"p\">:</span>
  <span class=\"mi\">90</span><span class=\"p\">,</span>\n            <span class=\"s2\">&quot;assignments_completed&quot;</span><span
  class=\"p\">:</span> <span class=\"mi\">14</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;student_id&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;BJ003&quot;</span>\n
  \       <span class=\"p\">}</span>\n    <span class=\"p\">},</span>\n    <span class=\"s2\">&quot;Countries&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n        <span class=\"s2\">&quot;USA&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;capital&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Washington, D.C.&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;population&quot;</span><span
  class=\"p\">:</span> <span class=\"mi\">331000000</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;area_sq_km&quot;</span><span class=\"p\">:</span>
  <span class=\"mi\">9833517</span><span class=\"p\">,</span>\n            <span class=\"s2\">&quot;official_languages&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">[</span><span class=\"s2\">&quot;English&quot;</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;Spanish&quot;</span><span class=\"p\">],</span>\n
  \           <span class=\"s2\">&quot;currency&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;United States Dollar (USD)&quot;</span>\n        <span
  class=\"p\">},</span>\n        <span class=\"s2\">&quot;India&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;capital&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;New Delhi&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;population&quot;</span><span class=\"p\">:</span>
  <span class=\"mi\">1380004385</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;area_sq_km&quot;</span><span class=\"p\">:</span> <span class=\"mi\">3287263</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;official_languages&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">[</span><span class=\"s2\">&quot;Hindi&quot;</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;English&quot;</span><span class=\"p\">],</span>\n
  \           <span class=\"s2\">&quot;currency&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;Indian Rupee (INR)&quot;</span>\n        <span class=\"p\">},</span>\n
  \       <span class=\"s2\">&quot;Brazil&quot;</span><span class=\"p\">:</span> <span
  class=\"p\">{</span>\n            <span class=\"s2\">&quot;capital&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Brasília&quot;</span><span class=\"p\">,</span>\n
  \           <span class=\"s2\">&quot;population&quot;</span><span class=\"p\">:</span>
  <span class=\"mi\">212559417</span><span class=\"p\">,</span>\n            <span
  class=\"s2\">&quot;area_sq_km&quot;</span><span class=\"p\">:</span> <span class=\"mi\">8515770</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;official_languages&quot;</span><span
  class=\"p\">:</span> <span class=\"p\">[</span><span class=\"s2\">&quot;Portuguese&quot;</span><span
  class=\"p\">],</span>\n            <span class=\"s2\">&quot;currency&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Brazilian Real (BRL)&quot;</span>\n
  \       <span class=\"p\">}</span>\n    <span class=\"p\">}</span>\n<span class=\"p\">}</span>\n\n<span
  class=\"k\">def</span> <span class=\"nf\">populate_table</span><span class=\"p\">(</span><span
  class=\"n\">writer</span><span class=\"p\">,</span> <span class=\"n\">data</span><span
  class=\"p\">,</span> <span class=\"n\">index_key</span><span class=\"p\">):</span>\n
  \   <span class=\"n\">row_keys</span> <span class=\"o\">=</span> <span class=\"nb\">list</span><span
  class=\"p\">(</span><span class=\"n\">data</span><span class=\"o\">.</span><span
  class=\"n\">keys</span><span class=\"p\">())</span>\n    <span class=\"k\">if</span>
  <span class=\"n\">row_keys</span> <span class=\"ow\">and</span> <span class=\"n\">data</span><span
  class=\"p\">[</span><span class=\"n\">row_keys</span><span class=\"p\">[</span><span
  class=\"mi\">0</span><span class=\"p\">]]</span> <span class=\"ow\">is</span> <span
  class=\"ow\">not</span> <span class=\"kc\">None</span><span class=\"p\">:</span>\n
  \       <span class=\"n\">headers</span> <span class=\"o\">=</span> <span class=\"p\">[</span><span
  class=\"n\">index_key</span><span class=\"p\">]</span> <span class=\"o\">+</span>
  <span class=\"nb\">list</span><span class=\"p\">(</span><span class=\"n\">data</span><span
  class=\"p\">[</span><span class=\"n\">row_keys</span><span class=\"p\">[</span><span
  class=\"mi\">0</span><span class=\"p\">]]</span><span class=\"o\">.</span><span
  class=\"n\">keys</span><span class=\"p\">())</span>\n    <span class=\"k\">else</span><span
  class=\"p\">:</span>\n        <span class=\"k\">return</span>\n    <span class=\"n\">writer</span><span
  class=\"o\">.</span><span class=\"n\">writerow</span><span class=\"p\">(</span><span
  class=\"n\">headers</span><span class=\"p\">)</span>\n    <span class=\"k\">for</span>
  <span class=\"n\">key</span> <span class=\"ow\">in</span> <span class=\"n\">row_keys</span><span
  class=\"p\">:</span>\n        <span class=\"n\">row</span> <span class=\"o\">=</span>
  <span class=\"p\">[</span><span class=\"n\">key</span><span class=\"p\">]</span>
  <span class=\"o\">+</span> <span class=\"nb\">list</span><span class=\"p\">(</span><span
  class=\"n\">data</span><span class=\"p\">[</span><span class=\"n\">key</span><span
  class=\"p\">]</span><span class=\"o\">.</span><span class=\"n\">values</span><span
  class=\"p\">())</span>\n        <span class=\"n\">writer</span><span class=\"o\">.</span><span
  class=\"n\">writerow</span><span class=\"p\">(</span><span class=\"n\">row</span><span
  class=\"p\">)</span>\n\n\n<span class=\"k\">with</span> <span class=\"nb\">open</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;data.csv&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;w&#39;</span><span class=\"p\">,</span> <span class=\"n\">newline</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;&#39;</span><span class=\"p\">)</span>
  <span class=\"k\">as</span> <span class=\"n\">csvfile</span><span class=\"p\">:</span>\n
  \   <span class=\"n\">writer</span> <span class=\"o\">=</span> <span class=\"n\">csv</span><span
  class=\"o\">.</span><span class=\"n\">writer</span><span class=\"p\">(</span><span
  class=\"n\">csvfile</span><span class=\"p\">)</span>\n    <span class=\"k\">for</span>
  <span class=\"n\">key</span> <span class=\"ow\">in</span> <span class=\"n\">data</span><span
  class=\"p\">:</span>\n        <span class=\"n\">populate_table</span><span class=\"p\">(</span><span
  class=\"n\">writer</span><span class=\"p\">,</span> <span class=\"n\">data</span><span
  class=\"p\">[</span><span class=\"n\">key</span><span class=\"p\">],</span> <span
  class=\"n\">key</span><span class=\"p\">)</span>\n        <span class=\"n\">writer</span><span
  class=\"o\">.</span><span class=\"n\">writerow</span><span class=\"p\">([])</span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n<div class='language-bar'>csv</div>\n<div class=\"highlight\"><pre><span></span>Students,course,grade,attendance,assignments_completed,student_id\nJohn
  Doe,Mathematics,A,95,15,JD001\nAlice Smith,Physics,B+,85,12,AS002\nBob Johnson,Computer
  Science,A-,90,14,BJ003\n\nCountries,capital,population,area_sq_km,official_languages,currency\nUSA,&quot;Washington,
  D.C.&quot;,331000000,9833517,&quot;[&#39;English&#39;, &#39;Spanish&#39;]&quot;,United
  States Dollar (USD)\nIndia,New Delhi,1380004385,3287263,&quot;[&#39;Hindi&#39;,
  &#39;English&#39;]&quot;,Indian Rupee (INR)\nBrazil,Brasília,212559417,8515770,[&#39;Portuguese&#39;],Brazilian
  Real (BRL)\n</pre></div>\n\n</pre>\n\n<p>Thank you, Happy Coding :)</p>\n<div class='prevnext'>\n
  \   <style type='text/css'>\n    :root {\n      --prevnext-color-text: #eefbfe;\n
  \     --prevnext-color-angle: #ff6600;\n      --prevnext-subtitle-brightness: 3;\n
  \   }\n    [data-theme=\"light\"] {\n      --prevnext-color-text: #1f2022;\n      --prevnext-color-angle:
  #ffeb00;\n      --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"dark\"]
  {\n      --prevnext-color-text: #eefbfe;\n      --prevnext-color-angle: #ff6600;\n
  \     --prevnext-subtitle-brightness: 3;\n    }\n    .prevnext {\n      display:
  flex;\n      flex-direction: row;\n      justify-content: space-around;\n      align-items:
  flex-start;\n    }\n    .prevnext a {\n      display: flex;\n      align-items:
  center;\n      width: 100%;\n      text-decoration: none;\n    }\n    a.next {\n
  \     justify-content: flex-end;\n    }\n    .prevnext a:hover {\n      background:
  #00000006;\n    }\n    .prevnext-subtitle {\n      color: var(--prevnext-color-text);\n
  \     filter: brightness(var(--prevnext-subtitle-brightness));\n      font-size:
  .8rem;\n    }\n    .prevnext-title {\n      color: var(--prevnext-color-text);\n
  \     font-size: 1rem;\n    }\n    .prevnext-text {\n      max-width: 30vw;\n    }\n
  \   </style>\n<pre><code>&lt;a class='prev' href='/vim-python-black-autoformat'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Autoformat
  Python file with Black after saving in Vim&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a
  class='next' href='/python-search-replace-file'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Python:
  Search and Replace in File&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot;
  height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
long_description: 'Today, I want to share with you a neat trick I recently discovered
  for populating a CSV file with data in a table-like structure using Python. Firstly,
  let Now, let Assuming you have a structure of the dictionary like: And you want
  to write it to a C'
now: 2025-05-01 18:11:33.315172
path: blog/tils/2024-03-20-Python-dict-to-CSV-table.md
slug: python-dict-to-csv-table
status: published-til
tags:
- python
templateKey: til
title: Turn Python dictionary into a neat CSV table
today: 2025-05-01
---

## Populating a Python dict having a table-like structure to a CSV

Today, I want to share with you a neat trick I recently discovered for populating a CSV file with data in a table-like structure using Python.

### Writing Key-Value Pairs to a CSV Row

Firstly, let's discuss the `write_key_value` function. This function allows us to write key-value pairs to a CSV row. It's particularly useful when dealing with metrics or data that can be represented as simple pairs.

```python
# Function to populate a CSV row with key-value pairs
def write_key_value(writer, dictionary):
    for key, value in dictionary.items():
        writer.writerow([key, value])
```

### Writing a Table-Like Structure to a CSV File

Now, let's dive into the `write_table` function, which handles more complex scenarios where the data follows a table-like structure. This function takes into account different types of metrics and adjusts the CSV table structure accordingly.

Assuming you have a structure of the dictionary like:

```python
data = {
    "Students": {
        "John Doe": {
            "course": "Mathematics",
            "grade": "A",
            "attendance": 95,
            "assignments_completed": 15,
            "student_id": "JD001"
        },
        "Alice Smith": {
            "course": "Physics",
            "grade": "B+",
            "attendance": 85,
            "assignments_completed": 12,
            "student_id": "AS002"
        },
        "Bob Johnson": {
            "course": "Computer Science",
            "grade": "A-",
            "attendance": 90,
            "assignments_completed": 14,
            "student_id": "BJ003"
        }
    }
}
```

And you want to write it to a CSV file, like this:

```csv
student, course, grade, attendance, assignments_completed, student_id
John Doe, Mathematics, A, 95, 15, JD001
Alice Smith, Physics, B+, 85, 12, AS002
Bob Johnson, Computer Science, A-, 90, 14, BJ003
```

We can create a function `write_table` that will take in the `dictionary` as the actual data. We want to store the keys of the inner dictionary to be the header/columns of the csv file. As we can see the keys of the inner dict i.e. the value for the key `John Doe` is a dict with the keys `course`, `grade`, `attendance`, etc. which remain the same for the all the keys in the dictionary.

So, we can first create a `row_keys` variable to store the keys of the actual dictionary this will be the first column rows in the csv. 

Further we check if the `row_keys` is a dict and then we append it with the `index_key` which will be the first column in the csv. Since all the keys remain the same for the inner-dict, we can pick the first dict and create the `header` with the inner-dict keys.

So, we can write the list `header` to the csv file.

Then for each key in the `row_keys` we can create a list `row` with the key and the values of the inner-dict.


```python
# Function to populate a CSV with a table-like structure
def write_table(writer, dictionary, index_key):

    row_keys = list(dictionary.keys())

    if row_keys and data[row_keys[0]] is not None:
        headers = [index_key] + list(
            dictionary[row_keys[0]].keys()
        )
    else:
        return
    writer.writerow(headers)
    for key in row_keys:
        row = [key] + list(dictionary[key].values())
        writer.writerow(row)


with open('data.csv', 'w', newline='') as csvfile:
    writer = csv.writer(csvfile)
    for key in data:
        write_table(writer, data[key], key)
```

### Example Usage

To illustrate how these functions can be used, let's consider a scenario where we have various types of metrics to populate into a CSV file. We handle key-value paired metrics separately and then populate the CSV with table-like metrics.

```python
import csv

data = {
    "Students": {
        "John Doe": {
            "course": "Mathematics",
            "grade": "A",
            "attendance": 95,
            "assignments_completed": 15,
            "student_id": "JD001"
        },
        "Alice Smith": {
            "course": "Physics",
            "grade": "B+",
            "attendance": 85,
            "assignments_completed": 12,
            "student_id": "AS002"
        },
        "Bob Johnson": {
            "course": "Computer Science",
            "grade": "A-",
            "attendance": 90,
            "assignments_completed": 14,
            "student_id": "BJ003"
        }
    },
    "Countries": {
        "USA": {
            "capital": "Washington, D.C.",
            "population": 331000000,
            "area_sq_km": 9833517,
            "official_languages": ["English", "Spanish"],
            "currency": "United States Dollar (USD)"
        },
        "India": {
            "capital": "New Delhi",
            "population": 1380004385,
            "area_sq_km": 3287263,
            "official_languages": ["Hindi", "English"],
            "currency": "Indian Rupee (INR)"
        },
        "Brazil": {
            "capital": "Brasília",
            "population": 212559417,
            "area_sq_km": 8515770,
            "official_languages": ["Portuguese"],
            "currency": "Brazilian Real (BRL)"
        }
    }
}

def populate_table(writer, data, index_key):
    row_keys = list(data.keys())
    if row_keys and data[row_keys[0]] is not None:
        headers = [index_key] + list(data[row_keys[0]].keys())
    else:
        return
    writer.writerow(headers)
    for key in row_keys:
        row = [key] + list(data[key].values())
        writer.writerow(row)


with open('data.csv', 'w', newline='') as csvfile:
    writer = csv.writer(csvfile)
    for key in data:
        populate_table(writer, data[key], key)
        writer.writerow([])
```

```csv
Students,course,grade,attendance,assignments_completed,student_id
John Doe,Mathematics,A,95,15,JD001
Alice Smith,Physics,B+,85,12,AS002
Bob Johnson,Computer Science,A-,90,14,BJ003

Countries,capital,population,area_sq_km,official_languages,currency
USA,"Washington, D.C.",331000000,9833517,"['English', 'Spanish']",United States Dollar (USD)
India,New Delhi,1380004385,3287263,"['Hindi', 'English']",Indian Rupee (INR)
Brazil,Brasília,212559417,8515770,['Portuguese'],Brazilian Real (BRL)
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
    
    <a class='prev' href='/vim-python-black-autoformat'>
    
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M13.5 8.25L9.75 12L13.5 15.75" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"> </path>
        </svg>
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>prev</p>
            <p class='prevnext-title'>Autoformat Python file with Black after saving in Vim</p>
        </div>
    </a>
    
    <a class='next' href='/python-search-replace-file'>
    
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>next</p>
            <p class='prevnext-title'>Python: Search and Replace in File</p>
        </div>
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.5 15.75L14.25 12L10.5 8.25" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path>
        </svg>
    </a>
  </div>