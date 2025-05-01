---
article_html: "<p>Bash or shell won't be much popular and powerful if they didn't
  have some tools and utilities baked in. But even further they are supported natively
  in Bash, which just makes every task and challenge quite hassle-free to deal with.
  In this part of the series, I'll try to cover an overview of some quite powerful
  and robust tools and utilities in Bash(shell in general) and also some of the advanced
  topics like dictionaries and positional parameters. Enough talk let's dive in.</p>\n<p>The
  topics to be covered in this part include the following:</p>\n<ul>\n<li>Hash tables/dictionaries
  in BASH</li>\n<li>Positional parameters</li>\n<li>Aliases in BASH</li>\n<li>Some
  Tools and utilities\n<ul>\n<li>grep/sed/awk</li>\n<li>cat/tac/head/tail</li>\n<li>cURL</li>\n<li>find</li>\n<li>bc</li>\n<li>wc</li>\n</ul>\n</li>\n</ul>\n<h2>Bash
  dictionaries</h2>\n<p>Bash dictionaries or hash tables are just like any other hash
  tables or keymaps in other programming languages. Bash dictionaries are quite similar
  to arrays but they have a key instead of the index(0,1,2...) and a value just like
  arrays. This can be quite useful for storing passwords with emails or usernames
  or any other way in which a value can be accessed only via a unique key.</p>\n<p>To
  declare a dictionary/ hash table, we can simply write <code>declare -A name</code>,
  this will declare an empty hash map for us. Further, we can populate the hash map
  with keys and values using the same syntax as of array just instead of numbers we
  can also have strings.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">declare</span> -A fruits\nfruits<span
  class=\"o\">[</span><span class=\"s2\">&quot;apple&quot;</span><span class=\"o\">]=</span><span
  class=\"s2\">&quot;red&quot;</span>\nfruits<span class=\"o\">[</span><span class=\"s2\">&quot;mango&quot;</span><span
  class=\"o\">]=</span><span class=\"s2\">&quot;yellow&quot;</span>\nfruits<span class=\"o\">[</span><span
  class=\"s2\">&quot;grapes&quot;</span><span class=\"o\">]=</span><span class=\"s2\">&quot;green&quot;</span>\n\n<span
  class=\"nb\">read</span> -p <span class=\"s2\">&quot;Enter the name of fruit : &quot;</span>
  name \n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;The fruit is </span><span
  class=\"nv\">$name</span><span class=\"s2\"> and its color is </span><span class=\"si\">${</span><span
  class=\"nv\">fruits</span><span class=\"p\">[</span><span class=\"nv\">$name</span><span
  class=\"p\">]</span><span class=\"si\">}</span><span class=\"s2\"> &quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626167875237/A2TxYPNoS.png\"
  alt=\"possh.png\" /></p>\n<p>The above example depicts a way to declare, define
  and access the key values in a dictionary. The example may look silly but you get
  the idea. We can also access every key or value using the <code>@</code> variable
  and access the number of key-value pairs using the <code>#</code> variable just
  like an array.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">declare</span> -A fruits\nfruits<span
  class=\"o\">[</span><span class=\"s2\">&quot;apple&quot;</span><span class=\"o\">]=</span><span
  class=\"s2\">&quot;red&quot;</span>\nfruits<span class=\"o\">[</span><span class=\"s2\">&quot;mango&quot;</span><span
  class=\"o\">]=</span><span class=\"s2\">&quot;yellow&quot;</span>\nfruits<span class=\"o\">[</span><span
  class=\"s2\">&quot;grapes&quot;</span><span class=\"o\">]=</span><span class=\"s2\">&quot;green&quot;</span>\n\n<span
  class=\"k\">for</span> i <span class=\"k\">in</span> <span class=\"s2\">&quot;</span><span
  class=\"si\">${</span><span class=\"p\">!fruits[@]</span><span class=\"si\">}</span><span
  class=\"s2\">&quot;</span><span class=\"p\">;</span>\n<span class=\"k\">do</span>\n\t<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"nv\">$i</span><span
  class=\"s2\"> : </span><span class=\"si\">${</span><span class=\"nv\">fruits</span><span
  class=\"p\">[</span><span class=\"nv\">$i</span><span class=\"p\">]</span><span
  class=\"si\">}</span><span class=\"s2\">&quot;</span>\n<span class=\"k\">done</span>\n\n<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;There are </span><span class=\"si\">${#</span><span
  class=\"nv\">fruits</span><span class=\"p\">[@]</span><span class=\"si\">}</span><span
  class=\"s2\"> key-value pairs.&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626171398570/14jJl2eFs.png\"
  alt=\"possh.png\" /></p>\n<p>OK, this is very tricky they are the same variables
  but used slightly differently. Firstly in the range-based for loop <code>${!fruits[@]}</code>,
  focus on the <code>!</code> before the hash table name, this will expand(depict)
  the hash map's keys. This is used to access every key from the hash table and we
  can also see <code>#</code> at the beginning of the hash map name as it is used
  to represent the entire hash map further in the <code>{#fruits[@]}</code> we can
  also use <code>{#fruits[*]}</code>.  Inside the for loop, <code>i</code> will be
  the key, and <code>{fruits[$i]}</code> will be the value for that <code>i</code>
  th key.</p>\n<p>Also, you can notice the bash interpreter automatically arranges
  the map in the alphabetical order of the values and not keys. This is quite a neat
  little feature that can come in handy a lot of times.</p>\n<p>If you want to delete
  or add any key-value pairs we can do that by the following commands:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">declare</span> -A fruits\nfruits<span
  class=\"o\">[</span><span class=\"s2\">&quot;apple&quot;</span><span class=\"o\">]=</span><span
  class=\"s2\">&quot;red&quot;</span>\nfruits<span class=\"o\">[</span><span class=\"s2\">&quot;mango&quot;</span><span
  class=\"o\">]=</span><span class=\"s2\">&quot;yellow&quot;</span>\nfruits<span class=\"o\">[</span><span
  class=\"s2\">&quot;grapes&quot;</span><span class=\"o\">]=</span><span class=\"s2\">&quot;green&quot;</span>\n\n<span
  class=\"k\">for</span> i <span class=\"k\">in</span> <span class=\"s2\">&quot;</span><span
  class=\"si\">${</span><span class=\"p\">!fruits[@]</span><span class=\"si\">}</span><span
  class=\"s2\">&quot;</span><span class=\"p\">;</span>\n<span class=\"k\">do</span>\n\t<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"nv\">$i</span><span
  class=\"s2\"> : </span><span class=\"si\">${</span><span class=\"nv\">fruits</span><span
  class=\"p\">[</span><span class=\"nv\">$i</span><span class=\"p\">]</span><span
  class=\"si\">}</span><span class=\"s2\">&quot;</span>\n<span class=\"k\">done</span>\n<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;There are </span><span class=\"si\">${#</span><span
  class=\"nv\">fruits</span><span class=\"p\">[@]</span><span class=\"si\">}</span><span
  class=\"s2\"> key-value pairs.&quot;</span>\n\n<span class=\"nb\">unset</span> fruits<span
  class=\"o\">[</span><span class=\"s2\">&quot;mango&quot;</span><span class=\"o\">]</span>
  \n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;There are </span><span
  class=\"si\">${#</span><span class=\"nv\">fruits</span><span class=\"p\">[@]</span><span
  class=\"si\">}</span><span class=\"s2\"> key-value pairs.&quot;</span>\nfruits<span
  class=\"o\">[</span><span class=\"s2\">&quot;strawberry&quot;</span><span class=\"o\">]=</span><span
  class=\"s2\">&quot;pink&quot;</span>\n\n<span class=\"k\">for</span> i <span class=\"k\">in</span>
  <span class=\"s2\">&quot;</span><span class=\"si\">${</span><span class=\"p\">!fruits[@]</span><span
  class=\"si\">}</span><span class=\"s2\">&quot;</span><span class=\"p\">;</span>\n<span
  class=\"k\">do</span>\n\t<span class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span
  class=\"nv\">$i</span><span class=\"s2\"> : </span><span class=\"si\">${</span><span
  class=\"nv\">fruits</span><span class=\"p\">[</span><span class=\"nv\">$i</span><span
  class=\"p\">]</span><span class=\"si\">}</span><span class=\"s2\">&quot;</span>\n<span
  class=\"k\">done</span>\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;There
  are </span><span class=\"si\">${#</span><span class=\"nv\">fruits</span><span class=\"p\">[@]</span><span
  class=\"si\">}</span><span class=\"s2\"> key-value pairs.&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626172120966/yCeXpaM9w.png\"
  alt=\"possh.png\" /></p>\n<p>The above code example is a bit complex but easy enough
  to understand. We can delete the key-value pair using the unset command and pass
  in the key along with the hash map name. We can create a key-value pair by simple
  command as depicted in the above example. This was a basic overview of hash maps/dictionaries
  in BASH.</p>\n<h2>Positional parameters</h2>\n<p>We often use user input from within
  the script but there is another way to pass in parameters outside of the script
  using positional parameters. It basically allows us to pass in arguments or parameters
  from the command prompt/ shell and inside of the script, we can access them via
  Positional Parameters ( $1, $2, $3....$9, ${10} and so on).</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;first
  parameter : &quot;</span> <span class=\"nv\">$1</span>\n<span class=\"nb\">echo</span>
  <span class=\"s2\">&quot;second parameter : &quot;</span> <span class=\"nv\">$2</span>\n<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;eleventh parameter : &quot;</span>
  <span class=\"si\">${</span><span class=\"nv\">11</span><span class=\"si\">}</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626159559317/XSyVwkC9U.png\"
  alt=\"possh.png\" /></p>\n<p>You can see from the above script we have passed the
  parameters from the command line just after typing the filename. The positional
  parameter $0 is the file name and from 1 onwards the parameters are optional to
  run if only your script needs any parameters or input to work with. The numbers
  are just random and just used for demonstration. The 11th parameter or double-digit
  parameter starting from 10 onwards, you need to encapsulate the number in {curly
  braces}<code>${number}</code> because it won't interpret <code>$10</code> or any
  other number as just <code>$1</code> and print 0.</p>\n<p>So we can pass a list
  of parameters that should be space-separated. We can pass any relevant information
  such as a string, number, or file names as well.</p>\n<p>If we want to access all
  the parameters passed to the script, we can use <code>@</code> variable. You may
  know this symbol from the array section of part-II, it is used to access every element
  in the array. But here it is used to access every parameter passed to the script
  just behaving like a list of values.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;The
  parameters passed are: &quot;</span> <span class=\"nv\">$@</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626160205206/eH6BD1r_Yu.png\"
  alt=\"possh.png\" /></p>\n<p>To get the number of the parameters passed to the script,
  we can use <code>#</code> variable. This is also a variable used in the array section
  for accessing the number of elements in the array, in this case, the number of parameters
  in the list.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;The
  parameters passed are: &quot;</span> <span class=\"nv\">$#</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626160630991/jVvJqtCqs.png\"
  alt=\"possh.png\" /></p>\n<p>Positional parameters allow to not take explicit input
  from the user from the script. This might not be used for the simple scripts but
  for administration purposes for the admins and users who know what does the script
  needs and it allows them to pass in arguments without designing the input system.</p>\n<h2>Bash
  aliases</h2>\n<p>Bash aliases are a great way of reducing the command length and
  making it much easier to type and work with the scripts or any development-related
  work. Alias is a file called bash_aliases inside the .bashrc folder that contains
  our shortcut commands, it has a particular order to map certain commands with others.</p>\n<p>Let's
  see what is an alias first and then we'll see how to set it up.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">alias</span> <span class=\"nv\">cdc</span><span class=\"o\">=</span><span
  class=\"s1\">&#39;cd C:/Users/acer/Desktop/New\\ folder/Code/&#39;</span>\n</pre></div>\n\n</pre>\n\n<p>This
  will make it viable to just type cdc and I will be in this directory instead of
  printing all of the jargon. The command we need to use to replace the big command
  is <code>cdc</code>. The right command is the variable assigned the command and
  the left or its value is the command to be replaced with it.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626163036098/lDwlSdiry.gif\"
  alt=\"shalias.gif\" /></p>\n<p>This command will definitely defer on your machine
  and OS as the filesystems are different in each major operating system. We can quickly
  make other such alias or shortcuts so to speak for making the development process
  faster and efficient.</p>\n<p>Now let us see how we set up this environment for
  bash alias, it's quite straightforward. You need to create a hidden file named &quot;bashrc&quot;
  i.e the file name will be <code>.bashrc</code>. This file has to be in the root
  directory (the folder to which bash defaults). I do not mean the <code>root</code>
  directory in Linux but the repository in which your bash interpreter opens. Once
  you have created the file put any alias in the file and source the file using the
  command:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">source</span> .bashrc\n</pre></div>\n\n</pre>\n\n<p>And that would
  do the trick, you can now test your macro or shortcut by opening a new instance
  of the terminal. If this doesn't work for you, then you can check  <a href=\"https://opensource.com/article/19/7/bash-aliases\">this
  article</a>  for a broader understanding of the setup.</p>\n<h2>Bash tools &amp;
  utilities</h2>\n<p>What would you call BASH without grep or sed man! It's a sad
  life:( BASH comes with some absolutely powerful and handy tools and utilities such
  as grep, sed, awk, at, wc, find, tar, gzip, which, make, ping, cURL, wget, ssh,
  .... my words there is an unstoppable long list. Really they are quite important
  and lay the foundation for some quite complex tasks. Some web servers can become
  redundant if some of the tools went missing. Let us find why they are so powerful.</p>\n<h3>grep</h3>\n<p>GREP
  or global regular expression print is a tool or command that can find patterns using
  regular expressions in files/strings or any other piece of data. It's not just printing
  or searching for the text, besides all that it can also edit the file and store
  the output in the desired file or any variable by proving some arguments to it.
  Grep supports Pearl's regular expression as well. There is a lot of customization
  options and arguments available in grep that can just do anything. It becomes an
  irreplaceable tool for some complex tasks.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626179054053/9ttkV-MZZ.png\"
  alt=\"possh.png\" /></p>\n<p>The above code finds the pattern &quot;more text&quot;
  in the file specified and prints the line to the screen, but we can modify the output
  we want, extract the output in a file and do all kinds of wizardry with this tool.
  This is just a basic, to get started example but trust me it's more than you think,
  this tool is used widely for web scrapping and pattern matching in quite a lot of
  use cases.</p>\n<h3>sed</h3>\n<p>SED or stream editor is another beast in BASH's
  toolkit, this is just a flawless tool. No words for this. This is a great tool but
  still underrated. This can actually edit the text inside the terminal, no graphical
  environment, no interface at all just commands, but it can do what a huge text editor
  can't! Save time, just edit text without opening anything except a terminal, becomes
  unbeatable in large files. This is surely a tiny little application that can skyrocket
  the efficiency development process.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626179410079/BkvdAqkfDS.png\"
  alt=\"possh.png\" /></p>\n<p>The above example replaces the word <code>more</code>
  with <code>less</code> using the sed command if you type 3g or nth line preceded
  by g, you will replace the word on the nth line only. In this case, only<code>g</code>
  will replace every occurrence of the word.\nThis is again a basic example of a sed
  command, its more if you go deeper, its more than a tool, its kind of a text-editor
  for wizards ;)</p>\n<h3>awk</h3>\n<p>awk or Aho, Weinberger, and Kernighan XD are
  the names of the developers of this application. This is another mind-blowing tool
  that can programmatically do a lot of stuff. This is like a programming language
  to a whole new level that can extrapolate and extract data from files and other
  forms of data. This is quite a great option if you want to quite neatly do something.
  It has great support libraries and functions that can even perform complex mathematical
  and scientific operations.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626180322400/hWkEVhPl_.png\"
  alt=\"possh.png\" /></p>\n<p>These are the topics for separate articles because
  it is insufficient to explain everything at once.</p>\n<h3>cat / tac / head / tail</h3>\n<p>CAT
  or concatenate is a tool used for printing out files, create files, sorting the
  contents of files, editing files, and a plethora of stuff. This command is generally
  used for printing the file but there is more to it like creating a file directly
  in the terminal, merging two files, and a ton of other operations.</p>\n<p>TAC or
  reverse of CAT is a tool used for everything that CAT can do but in reverse:) This
  is a wired tool but still quite useful sometimes.</p>\n<p>Head is a tool that will
  print or edit the text in the first 10 lines of the file, it can be used to extrapolate
  multiple files with similar content.\nTail is a tool that will print or edit the
  text in the last 10 lines of the file, it can be used just like head but for the
  last few lines.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626180451092/Z5VUpIxCm.png\"
  alt=\"possh.png\" /></p>\n<p>It turns out, you can not only print the first or last
  10 lines but n lines by passing the -n as an argument, there is a ton of stuff to
  discover, this just drives me crazy.</p>\n<h3>cURL</h3>\n<p>cURL or client URL is
  a tool that can be used to transfer data via various network protocols. You might
  not believe but it is used in cars, televisions, routers, and other embedded systems
  for exchanging relevant data via appropriate protocols.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626181263235/JPocJwoEd.png\"
  alt=\"possh.png\" /></p>\n<p>This example depicts how we can fetch data from an
  API using cURL and extract data in JSON format and use it for relevant tasks.\nThis
  is again one of the best utility out there as it becomes quite remarkable and vintage.
  Despite being almost 30 years old, it shines bright in the tech world.</p>\n<h3>find</h3>\n<p>Find
  as the name suggests it is used to find files among the folders and directories
  in a file system. it becomes quite helpful in complex projects where the directory
  structure is deep and large.\n<img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626181386566/zpP9Yaom4.png\"
  alt=\"possh.png\" /></p>\n<p>The command <code>find *.txt</code> finds every txt
  file available on this directory. As simple as it can get. This is surely looking
  silly and idiotic but it finds its glory in large and complicated codebases.</p>\n<h3>bc</h3>\n<p>bc
  or basic calculator is a utility tool for performing mathematical and arithmetical
  operations in the terminal, this commands gets integrated with other commands such
  as awk really well and can be used for further extending the limits of what the
  command line development can do.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626182601487/z8X4KeDGG.png\"
  alt=\"possh.png\" /></p>\n<p>AWW! I could hear the excitement. That just added new
  dimensions into BASH. Just creativity and resonance to anything is the limit here.
  I am using  <a href=\"http://repl.it/\">REPL.IT</a>  here for using bash as I do
  not have it on my windows machine :( But that command is truly insane.</p>\n<h3>wc</h3>\n<p>wc
  or word count is a utility tool for counting and analyzing the size or count of
  characters, words, lines, or files in a given file structure. This is quite a handy
  tool for monitoring and keeping track of a system, also for general development
  purposes.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626182319602/r8UidHV2z.png\"
  alt=\"possh.png\" /></p>\n<p>The above command prints out the word and lines in
  the provided file. This command <code>wc</code> can even compute the size of files
  and even more properties of files.\nThose were some of the quite powerful commands,
  tools, or utilities in BASH/shell. There are plenty of other commands not covered
  here because this an extremely large topic and even making separate articles or
  resources there will certainly and surely be some things that will get missed out,
  that's the beauty Linux or in general Computer Science has.\nOk, that was a lot,
  but I hope you got some insights for learning more BASH or Linux in general. This
  is a wide topic and can't be covered entirely in a single article.</p>\n<p>Now that
  is it from this part, everything cannot be covered in any number of parts but at
  least it will help someone to get started in BASH scripting and its specifications
  for development. Have a Blast learning BASH. Happy Coding :)</p>\n"
cover: ''
date: 2021-07-12
datetime: 2021-07-12 00:00:00+00:00
description: 'Bash or shell won The topics to be covered in this part include the
  following: Hash tables/dictionaries in BASH Positional parameters Aliases in BASH
  Some Tools'
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-07-12-BASH-scripting-P3.md
html: "<p>Bash or shell won't be much popular and powerful if they didn't have some
  tools and utilities baked in. But even further they are supported natively in Bash,
  which just makes every task and challenge quite hassle-free to deal with. In this
  part of the series, I'll try to cover an overview of some quite powerful and robust
  tools and utilities in Bash(shell in general) and also some of the advanced topics
  like dictionaries and positional parameters. Enough talk let's dive in.</p>\n<p>The
  topics to be covered in this part include the following:</p>\n<ul>\n<li>Hash tables/dictionaries
  in BASH</li>\n<li>Positional parameters</li>\n<li>Aliases in BASH</li>\n<li>Some
  Tools and utilities\n<ul>\n<li>grep/sed/awk</li>\n<li>cat/tac/head/tail</li>\n<li>cURL</li>\n<li>find</li>\n<li>bc</li>\n<li>wc</li>\n</ul>\n</li>\n</ul>\n<h2>Bash
  dictionaries</h2>\n<p>Bash dictionaries or hash tables are just like any other hash
  tables or keymaps in other programming languages. Bash dictionaries are quite similar
  to arrays but they have a key instead of the index(0,1,2...) and a value just like
  arrays. This can be quite useful for storing passwords with emails or usernames
  or any other way in which a value can be accessed only via a unique key.</p>\n<p>To
  declare a dictionary/ hash table, we can simply write <code>declare -A name</code>,
  this will declare an empty hash map for us. Further, we can populate the hash map
  with keys and values using the same syntax as of array just instead of numbers we
  can also have strings.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">declare</span> -A fruits\nfruits<span
  class=\"o\">[</span><span class=\"s2\">&quot;apple&quot;</span><span class=\"o\">]=</span><span
  class=\"s2\">&quot;red&quot;</span>\nfruits<span class=\"o\">[</span><span class=\"s2\">&quot;mango&quot;</span><span
  class=\"o\">]=</span><span class=\"s2\">&quot;yellow&quot;</span>\nfruits<span class=\"o\">[</span><span
  class=\"s2\">&quot;grapes&quot;</span><span class=\"o\">]=</span><span class=\"s2\">&quot;green&quot;</span>\n\n<span
  class=\"nb\">read</span> -p <span class=\"s2\">&quot;Enter the name of fruit : &quot;</span>
  name \n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;The fruit is </span><span
  class=\"nv\">$name</span><span class=\"s2\"> and its color is </span><span class=\"si\">${</span><span
  class=\"nv\">fruits</span><span class=\"p\">[</span><span class=\"nv\">$name</span><span
  class=\"p\">]</span><span class=\"si\">}</span><span class=\"s2\"> &quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626167875237/A2TxYPNoS.png\"
  alt=\"possh.png\" /></p>\n<p>The above example depicts a way to declare, define
  and access the key values in a dictionary. The example may look silly but you get
  the idea. We can also access every key or value using the <code>@</code> variable
  and access the number of key-value pairs using the <code>#</code> variable just
  like an array.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">declare</span> -A fruits\nfruits<span
  class=\"o\">[</span><span class=\"s2\">&quot;apple&quot;</span><span class=\"o\">]=</span><span
  class=\"s2\">&quot;red&quot;</span>\nfruits<span class=\"o\">[</span><span class=\"s2\">&quot;mango&quot;</span><span
  class=\"o\">]=</span><span class=\"s2\">&quot;yellow&quot;</span>\nfruits<span class=\"o\">[</span><span
  class=\"s2\">&quot;grapes&quot;</span><span class=\"o\">]=</span><span class=\"s2\">&quot;green&quot;</span>\n\n<span
  class=\"k\">for</span> i <span class=\"k\">in</span> <span class=\"s2\">&quot;</span><span
  class=\"si\">${</span><span class=\"p\">!fruits[@]</span><span class=\"si\">}</span><span
  class=\"s2\">&quot;</span><span class=\"p\">;</span>\n<span class=\"k\">do</span>\n\t<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"nv\">$i</span><span
  class=\"s2\"> : </span><span class=\"si\">${</span><span class=\"nv\">fruits</span><span
  class=\"p\">[</span><span class=\"nv\">$i</span><span class=\"p\">]</span><span
  class=\"si\">}</span><span class=\"s2\">&quot;</span>\n<span class=\"k\">done</span>\n\n<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;There are </span><span class=\"si\">${#</span><span
  class=\"nv\">fruits</span><span class=\"p\">[@]</span><span class=\"si\">}</span><span
  class=\"s2\"> key-value pairs.&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626171398570/14jJl2eFs.png\"
  alt=\"possh.png\" /></p>\n<p>OK, this is very tricky they are the same variables
  but used slightly differently. Firstly in the range-based for loop <code>${!fruits[@]}</code>,
  focus on the <code>!</code> before the hash table name, this will expand(depict)
  the hash map's keys. This is used to access every key from the hash table and we
  can also see <code>#</code> at the beginning of the hash map name as it is used
  to represent the entire hash map further in the <code>{#fruits[@]}</code> we can
  also use <code>{#fruits[*]}</code>.  Inside the for loop, <code>i</code> will be
  the key, and <code>{fruits[$i]}</code> will be the value for that <code>i</code>
  th key.</p>\n<p>Also, you can notice the bash interpreter automatically arranges
  the map in the alphabetical order of the values and not keys. This is quite a neat
  little feature that can come in handy a lot of times.</p>\n<p>If you want to delete
  or add any key-value pairs we can do that by the following commands:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">declare</span> -A fruits\nfruits<span
  class=\"o\">[</span><span class=\"s2\">&quot;apple&quot;</span><span class=\"o\">]=</span><span
  class=\"s2\">&quot;red&quot;</span>\nfruits<span class=\"o\">[</span><span class=\"s2\">&quot;mango&quot;</span><span
  class=\"o\">]=</span><span class=\"s2\">&quot;yellow&quot;</span>\nfruits<span class=\"o\">[</span><span
  class=\"s2\">&quot;grapes&quot;</span><span class=\"o\">]=</span><span class=\"s2\">&quot;green&quot;</span>\n\n<span
  class=\"k\">for</span> i <span class=\"k\">in</span> <span class=\"s2\">&quot;</span><span
  class=\"si\">${</span><span class=\"p\">!fruits[@]</span><span class=\"si\">}</span><span
  class=\"s2\">&quot;</span><span class=\"p\">;</span>\n<span class=\"k\">do</span>\n\t<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"nv\">$i</span><span
  class=\"s2\"> : </span><span class=\"si\">${</span><span class=\"nv\">fruits</span><span
  class=\"p\">[</span><span class=\"nv\">$i</span><span class=\"p\">]</span><span
  class=\"si\">}</span><span class=\"s2\">&quot;</span>\n<span class=\"k\">done</span>\n<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;There are </span><span class=\"si\">${#</span><span
  class=\"nv\">fruits</span><span class=\"p\">[@]</span><span class=\"si\">}</span><span
  class=\"s2\"> key-value pairs.&quot;</span>\n\n<span class=\"nb\">unset</span> fruits<span
  class=\"o\">[</span><span class=\"s2\">&quot;mango&quot;</span><span class=\"o\">]</span>
  \n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;There are </span><span
  class=\"si\">${#</span><span class=\"nv\">fruits</span><span class=\"p\">[@]</span><span
  class=\"si\">}</span><span class=\"s2\"> key-value pairs.&quot;</span>\nfruits<span
  class=\"o\">[</span><span class=\"s2\">&quot;strawberry&quot;</span><span class=\"o\">]=</span><span
  class=\"s2\">&quot;pink&quot;</span>\n\n<span class=\"k\">for</span> i <span class=\"k\">in</span>
  <span class=\"s2\">&quot;</span><span class=\"si\">${</span><span class=\"p\">!fruits[@]</span><span
  class=\"si\">}</span><span class=\"s2\">&quot;</span><span class=\"p\">;</span>\n<span
  class=\"k\">do</span>\n\t<span class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span
  class=\"nv\">$i</span><span class=\"s2\"> : </span><span class=\"si\">${</span><span
  class=\"nv\">fruits</span><span class=\"p\">[</span><span class=\"nv\">$i</span><span
  class=\"p\">]</span><span class=\"si\">}</span><span class=\"s2\">&quot;</span>\n<span
  class=\"k\">done</span>\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;There
  are </span><span class=\"si\">${#</span><span class=\"nv\">fruits</span><span class=\"p\">[@]</span><span
  class=\"si\">}</span><span class=\"s2\"> key-value pairs.&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626172120966/yCeXpaM9w.png\"
  alt=\"possh.png\" /></p>\n<p>The above code example is a bit complex but easy enough
  to understand. We can delete the key-value pair using the unset command and pass
  in the key along with the hash map name. We can create a key-value pair by simple
  command as depicted in the above example. This was a basic overview of hash maps/dictionaries
  in BASH.</p>\n<h2>Positional parameters</h2>\n<p>We often use user input from within
  the script but there is another way to pass in parameters outside of the script
  using positional parameters. It basically allows us to pass in arguments or parameters
  from the command prompt/ shell and inside of the script, we can access them via
  Positional Parameters ( $1, $2, $3....$9, ${10} and so on).</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;first
  parameter : &quot;</span> <span class=\"nv\">$1</span>\n<span class=\"nb\">echo</span>
  <span class=\"s2\">&quot;second parameter : &quot;</span> <span class=\"nv\">$2</span>\n<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;eleventh parameter : &quot;</span>
  <span class=\"si\">${</span><span class=\"nv\">11</span><span class=\"si\">}</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626159559317/XSyVwkC9U.png\"
  alt=\"possh.png\" /></p>\n<p>You can see from the above script we have passed the
  parameters from the command line just after typing the filename. The positional
  parameter $0 is the file name and from 1 onwards the parameters are optional to
  run if only your script needs any parameters or input to work with. The numbers
  are just random and just used for demonstration. The 11th parameter or double-digit
  parameter starting from 10 onwards, you need to encapsulate the number in {curly
  braces}<code>${number}</code> because it won't interpret <code>$10</code> or any
  other number as just <code>$1</code> and print 0.</p>\n<p>So we can pass a list
  of parameters that should be space-separated. We can pass any relevant information
  such as a string, number, or file names as well.</p>\n<p>If we want to access all
  the parameters passed to the script, we can use <code>@</code> variable. You may
  know this symbol from the array section of part-II, it is used to access every element
  in the array. But here it is used to access every parameter passed to the script
  just behaving like a list of values.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;The
  parameters passed are: &quot;</span> <span class=\"nv\">$@</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626160205206/eH6BD1r_Yu.png\"
  alt=\"possh.png\" /></p>\n<p>To get the number of the parameters passed to the script,
  we can use <code>#</code> variable. This is also a variable used in the array section
  for accessing the number of elements in the array, in this case, the number of parameters
  in the list.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;The
  parameters passed are: &quot;</span> <span class=\"nv\">$#</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626160630991/jVvJqtCqs.png\"
  alt=\"possh.png\" /></p>\n<p>Positional parameters allow to not take explicit input
  from the user from the script. This might not be used for the simple scripts but
  for administration purposes for the admins and users who know what does the script
  needs and it allows them to pass in arguments without designing the input system.</p>\n<h2>Bash
  aliases</h2>\n<p>Bash aliases are a great way of reducing the command length and
  making it much easier to type and work with the scripts or any development-related
  work. Alias is a file called bash_aliases inside the .bashrc folder that contains
  our shortcut commands, it has a particular order to map certain commands with others.</p>\n<p>Let's
  see what is an alias first and then we'll see how to set it up.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">alias</span> <span class=\"nv\">cdc</span><span class=\"o\">=</span><span
  class=\"s1\">&#39;cd C:/Users/acer/Desktop/New\\ folder/Code/&#39;</span>\n</pre></div>\n\n</pre>\n\n<p>This
  will make it viable to just type cdc and I will be in this directory instead of
  printing all of the jargon. The command we need to use to replace the big command
  is <code>cdc</code>. The right command is the variable assigned the command and
  the left or its value is the command to be replaced with it.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626163036098/lDwlSdiry.gif\"
  alt=\"shalias.gif\" /></p>\n<p>This command will definitely defer on your machine
  and OS as the filesystems are different in each major operating system. We can quickly
  make other such alias or shortcuts so to speak for making the development process
  faster and efficient.</p>\n<p>Now let us see how we set up this environment for
  bash alias, it's quite straightforward. You need to create a hidden file named &quot;bashrc&quot;
  i.e the file name will be <code>.bashrc</code>. This file has to be in the root
  directory (the folder to which bash defaults). I do not mean the <code>root</code>
  directory in Linux but the repository in which your bash interpreter opens. Once
  you have created the file put any alias in the file and source the file using the
  command:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">source</span> .bashrc\n</pre></div>\n\n</pre>\n\n<p>And that would
  do the trick, you can now test your macro or shortcut by opening a new instance
  of the terminal. If this doesn't work for you, then you can check  <a href=\"https://opensource.com/article/19/7/bash-aliases\">this
  article</a>  for a broader understanding of the setup.</p>\n<h2>Bash tools &amp;
  utilities</h2>\n<p>What would you call BASH without grep or sed man! It's a sad
  life:( BASH comes with some absolutely powerful and handy tools and utilities such
  as grep, sed, awk, at, wc, find, tar, gzip, which, make, ping, cURL, wget, ssh,
  .... my words there is an unstoppable long list. Really they are quite important
  and lay the foundation for some quite complex tasks. Some web servers can become
  redundant if some of the tools went missing. Let us find why they are so powerful.</p>\n<h3>grep</h3>\n<p>GREP
  or global regular expression print is a tool or command that can find patterns using
  regular expressions in files/strings or any other piece of data. It's not just printing
  or searching for the text, besides all that it can also edit the file and store
  the output in the desired file or any variable by proving some arguments to it.
  Grep supports Pearl's regular expression as well. There is a lot of customization
  options and arguments available in grep that can just do anything. It becomes an
  irreplaceable tool for some complex tasks.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626179054053/9ttkV-MZZ.png\"
  alt=\"possh.png\" /></p>\n<p>The above code finds the pattern &quot;more text&quot;
  in the file specified and prints the line to the screen, but we can modify the output
  we want, extract the output in a file and do all kinds of wizardry with this tool.
  This is just a basic, to get started example but trust me it's more than you think,
  this tool is used widely for web scrapping and pattern matching in quite a lot of
  use cases.</p>\n<h3>sed</h3>\n<p>SED or stream editor is another beast in BASH's
  toolkit, this is just a flawless tool. No words for this. This is a great tool but
  still underrated. This can actually edit the text inside the terminal, no graphical
  environment, no interface at all just commands, but it can do what a huge text editor
  can't! Save time, just edit text without opening anything except a terminal, becomes
  unbeatable in large files. This is surely a tiny little application that can skyrocket
  the efficiency development process.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626179410079/BkvdAqkfDS.png\"
  alt=\"possh.png\" /></p>\n<p>The above example replaces the word <code>more</code>
  with <code>less</code> using the sed command if you type 3g or nth line preceded
  by g, you will replace the word on the nth line only. In this case, only<code>g</code>
  will replace every occurrence of the word.\nThis is again a basic example of a sed
  command, its more if you go deeper, its more than a tool, its kind of a text-editor
  for wizards ;)</p>\n<h3>awk</h3>\n<p>awk or Aho, Weinberger, and Kernighan XD are
  the names of the developers of this application. This is another mind-blowing tool
  that can programmatically do a lot of stuff. This is like a programming language
  to a whole new level that can extrapolate and extract data from files and other
  forms of data. This is quite a great option if you want to quite neatly do something.
  It has great support libraries and functions that can even perform complex mathematical
  and scientific operations.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626180322400/hWkEVhPl_.png\"
  alt=\"possh.png\" /></p>\n<p>These are the topics for separate articles because
  it is insufficient to explain everything at once.</p>\n<h3>cat / tac / head / tail</h3>\n<p>CAT
  or concatenate is a tool used for printing out files, create files, sorting the
  contents of files, editing files, and a plethora of stuff. This command is generally
  used for printing the file but there is more to it like creating a file directly
  in the terminal, merging two files, and a ton of other operations.</p>\n<p>TAC or
  reverse of CAT is a tool used for everything that CAT can do but in reverse:) This
  is a wired tool but still quite useful sometimes.</p>\n<p>Head is a tool that will
  print or edit the text in the first 10 lines of the file, it can be used to extrapolate
  multiple files with similar content.\nTail is a tool that will print or edit the
  text in the last 10 lines of the file, it can be used just like head but for the
  last few lines.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626180451092/Z5VUpIxCm.png\"
  alt=\"possh.png\" /></p>\n<p>It turns out, you can not only print the first or last
  10 lines but n lines by passing the -n as an argument, there is a ton of stuff to
  discover, this just drives me crazy.</p>\n<h3>cURL</h3>\n<p>cURL or client URL is
  a tool that can be used to transfer data via various network protocols. You might
  not believe but it is used in cars, televisions, routers, and other embedded systems
  for exchanging relevant data via appropriate protocols.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626181263235/JPocJwoEd.png\"
  alt=\"possh.png\" /></p>\n<p>This example depicts how we can fetch data from an
  API using cURL and extract data in JSON format and use it for relevant tasks.\nThis
  is again one of the best utility out there as it becomes quite remarkable and vintage.
  Despite being almost 30 years old, it shines bright in the tech world.</p>\n<h3>find</h3>\n<p>Find
  as the name suggests it is used to find files among the folders and directories
  in a file system. it becomes quite helpful in complex projects where the directory
  structure is deep and large.\n<img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626181386566/zpP9Yaom4.png\"
  alt=\"possh.png\" /></p>\n<p>The command <code>find *.txt</code> finds every txt
  file available on this directory. As simple as it can get. This is surely looking
  silly and idiotic but it finds its glory in large and complicated codebases.</p>\n<h3>bc</h3>\n<p>bc
  or basic calculator is a utility tool for performing mathematical and arithmetical
  operations in the terminal, this commands gets integrated with other commands such
  as awk really well and can be used for further extending the limits of what the
  command line development can do.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626182601487/z8X4KeDGG.png\"
  alt=\"possh.png\" /></p>\n<p>AWW! I could hear the excitement. That just added new
  dimensions into BASH. Just creativity and resonance to anything is the limit here.
  I am using  <a href=\"http://repl.it/\">REPL.IT</a>  here for using bash as I do
  not have it on my windows machine :( But that command is truly insane.</p>\n<h3>wc</h3>\n<p>wc
  or word count is a utility tool for counting and analyzing the size or count of
  characters, words, lines, or files in a given file structure. This is quite a handy
  tool for monitoring and keeping track of a system, also for general development
  purposes.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1626182319602/r8UidHV2z.png\"
  alt=\"possh.png\" /></p>\n<p>The above command prints out the word and lines in
  the provided file. This command <code>wc</code> can even compute the size of files
  and even more properties of files.\nThose were some of the quite powerful commands,
  tools, or utilities in BASH/shell. There are plenty of other commands not covered
  here because this an extremely large topic and even making separate articles or
  resources there will certainly and surely be some things that will get missed out,
  that's the beauty Linux or in general Computer Science has.\nOk, that was a lot,
  but I hope you got some insights for learning more BASH or Linux in general. This
  is a wide topic and can't be covered entirely in a single article.</p>\n<p>Now that
  is it from this part, everything cannot be covered in any number of parts but at
  least it will help someone to get started in BASH scripting and its specifications
  for development. Have a Blast learning BASH. Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/bash-scripting-guide-p3.webp
long_description: 'Bash or shell won The topics to be covered in this part include
  the following: Hash tables/dictionaries in BASH Positional parameters Aliases in
  BASH Some Tools and utilities grep/sed/awk cat/tac/head/tail cURL find bc wc Bash
  dictionaries or hash ta'
now: 2025-05-01 18:11:33.314569
path: blog/posts/2021-07-12-BASH-scripting-P3.md
prevnext: null
series:
- BASH Scripting Guide
slug: bash-guide-p3
status: published
subtitle: Exploring some advance features, underrated tools and utilities in BASH.
tags:
- bash
templateKey: blog—post
title: BASH Scripting Guide - PART - 3
today: 2025-05-01
---

Bash or shell won't be much popular and powerful if they didn't have some tools and utilities baked in. But even further they are supported natively in Bash, which just makes every task and challenge quite hassle-free to deal with. In this part of the series, I'll try to cover an overview of some quite powerful and robust tools and utilities in Bash(shell in general) and also some of the advanced topics like dictionaries and positional parameters. Enough talk let's dive in.

The topics to be covered in this part include the following:

- Hash tables/dictionaries in BASH
- Positional parameters
- Aliases in BASH
- Some Tools and utilities
    - grep/sed/awk
    - cat/tac/head/tail
    - cURL
    - find
    - bc
    - wc


## Bash dictionaries
Bash dictionaries or hash tables are just like any other hash tables or keymaps in other programming languages. Bash dictionaries are quite similar to arrays but they have a key instead of the index(0,1,2...) and a value just like arrays. This can be quite useful for storing passwords with emails or usernames or any other way in which a value can be accessed only via a unique key. 

To declare a dictionary/ hash table, we can simply write `declare -A name`, this will declare an empty hash map for us. Further, we can populate the hash map with keys and values using the same syntax as of array just instead of numbers we can also have strings. 

```bash
#!/bin/bash

declare -A fruits
fruits["apple"]="red"
fruits["mango"]="yellow"
fruits["grapes"]="green"

read -p "Enter the name of fruit : " name 
echo "The fruit is $name and its color is ${fruits[$name]} "

```

![possh.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626167875237/A2TxYPNoS.png)

The above example depicts a way to declare, define and access the key values in a dictionary. The example may look silly but you get the idea. We can also access every key or value using the `@` variable and access the number of key-value pairs using the `#` variable just like an array. 
```bash
#!/bin/bash

declare -A fruits
fruits["apple"]="red"
fruits["mango"]="yellow"
fruits["grapes"]="green"

for i in "${!fruits[@]}";
do
	echo "$i : ${fruits[$i]}"
done

echo "There are ${#fruits[@]} key-value pairs."
```

![possh.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626171398570/14jJl2eFs.png)

OK, this is very tricky they are the same variables but used slightly differently. Firstly in the range-based for loop `${!fruits[@]}`, focus on the `!` before the hash table name, this will expand(depict) the hash map's keys. This is used to access every key from the hash table and we can also see `#` at the beginning of the hash map name as it is used to represent the entire hash map further in the `{#fruits[@]}` we can also use `{#fruits[*]}`.  Inside the for loop, `i` will be the key, and `{fruits[$i]}` will be the value for that `i` th key.

Also, you can notice the bash interpreter automatically arranges the map in the alphabetical order of the values and not keys. This is quite a neat little feature that can come in handy a lot of times.

If you want to delete or add any key-value pairs we can do that by the following commands:
```bash
#!/bin/bash

declare -A fruits
fruits["apple"]="red"
fruits["mango"]="yellow"
fruits["grapes"]="green"

for i in "${!fruits[@]}";
do
	echo "$i : ${fruits[$i]}"
done
echo "There are ${#fruits[@]} key-value pairs."

unset fruits["mango"] 
echo "There are ${#fruits[@]} key-value pairs."
fruits["strawberry"]="pink"

for i in "${!fruits[@]}";
do
	echo "$i : ${fruits[$i]}"
done
echo "There are ${#fruits[@]} key-value pairs."

```

![possh.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626172120966/yCeXpaM9w.png)

The above code example is a bit complex but easy enough to understand. We can delete the key-value pair using the unset command and pass in the key along with the hash map name. We can create a key-value pair by simple command as depicted in the above example. This was a basic overview of hash maps/dictionaries in BASH.

##  Positional parameters

We often use user input from within the script but there is another way to pass in parameters outside of the script using positional parameters. It basically allows us to pass in arguments or parameters from the command prompt/ shell and inside of the script, we can access them via Positional Parameters ( $1, $2, $3....$9, ${10} and so on).

```bash
#!/bin/bash

echo "first parameter : " $1
echo "second parameter : " $2
echo "eleventh parameter : " ${11}

```
![possh.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626159559317/XSyVwkC9U.png)

You can see from the above script we have passed the parameters from the command line just after typing the filename. The positional parameter $0 is the file name and from 1 onwards the parameters are optional to run if only your script needs any parameters or input to work with. The numbers are just random and just used for demonstration. The 11th parameter or double-digit parameter starting from 10 onwards, you need to encapsulate the number in {curly braces}`${number}` because it won't interpret `$10` or any other number as just `$1` and print 0. 

So we can pass a list of parameters that should be space-separated. We can pass any relevant information such as a string, number, or file names as well. 

If we want to access all the parameters passed to the script, we can use `@` variable. You may know this symbol from the array section of part-II, it is used to access every element in the array. But here it is used to access every parameter passed to the script just behaving like a list of values.

```bash
#!/bin/bash

echo "The parameters passed are: " $@

```

![possh.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626160205206/eH6BD1r_Yu.png)
 
To get the number of the parameters passed to the script, we can use `#` variable. This is also a variable used in the array section for accessing the number of elements in the array, in this case, the number of parameters in the list.

```bash
#!/bin/bash

echo "The parameters passed are: " $#

```
![possh.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626160630991/jVvJqtCqs.png)

Positional parameters allow to not take explicit input from the user from the script. This might not be used for the simple scripts but for administration purposes for the admins and users who know what does the script needs and it allows them to pass in arguments without designing the input system.


## Bash aliases

Bash aliases are a great way of reducing the command length and making it much easier to type and work with the scripts or any development-related work. Alias is a file called bash_aliases inside the .bashrc folder that contains our shortcut commands, it has a particular order to map certain commands with others. 

Let's see what is an alias first and then we'll see how to set it up.
```bash
alias cdc='cd C:/Users/acer/Desktop/New\ folder/Code/'
```
This will make it viable to just type cdc and I will be in this directory instead of printing all of the jargon. The command we need to use to replace the big command is `cdc`. The right command is the variable assigned the command and the left or its value is the command to be replaced with it.

![shalias.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1626163036098/lDwlSdiry.gif)

This command will definitely defer on your machine and OS as the filesystems are different in each major operating system. We can quickly make other such alias or shortcuts so to speak for making the development process faster and efficient. 

Now let us see how we set up this environment for bash alias, it's quite straightforward. You need to create a hidden file named "bashrc" i.e the file name will be `.bashrc`. This file has to be in the root directory (the folder to which bash defaults). I do not mean the `root` directory in Linux but the repository in which your bash interpreter opens. Once you have created the file put any alias in the file and source the file using the command:
```bash
source .bashrc
```
And that would do the trick, you can now test your macro or shortcut by opening a new instance of the terminal. If this doesn't work for you, then you can check  [this article](https://opensource.com/article/19/7/bash-aliases)  for a broader understanding of the setup.

## Bash tools & utilities  

What would you call BASH without grep or sed man! It's a sad life:( BASH comes with some absolutely powerful and handy tools and utilities such as grep, sed, awk, at, wc, find, tar, gzip, which, make, ping, cURL, wget, ssh, .... my words there is an unstoppable long list. Really they are quite important and lay the foundation for some quite complex tasks. Some web servers can become redundant if some of the tools went missing. Let us find why they are so powerful.

### grep
GREP or global regular expression print is a tool or command that can find patterns using regular expressions in files/strings or any other piece of data. It's not just printing or searching for the text, besides all that it can also edit the file and store the output in the desired file or any variable by proving some arguments to it. Grep supports Pearl's regular expression as well. There is a lot of customization options and arguments available in grep that can just do anything. It becomes an irreplaceable tool for some complex tasks. 

![possh.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626179054053/9ttkV-MZZ.png)

The above code finds the pattern "more text" in the file specified and prints the line to the screen, but we can modify the output we want, extract the output in a file and do all kinds of wizardry with this tool. This is just a basic, to get started example but trust me it's more than you think, this tool is used widely for web scrapping and pattern matching in quite a lot of use cases.


### sed
SED or stream editor is another beast in BASH's toolkit, this is just a flawless tool. No words for this. This is a great tool but still underrated. This can actually edit the text inside the terminal, no graphical environment, no interface at all just commands, but it can do what a huge text editor can't! Save time, just edit text without opening anything except a terminal, becomes unbeatable in large files. This is surely a tiny little application that can skyrocket the efficiency development process. 

![possh.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626179410079/BkvdAqkfDS.png)

The above example replaces the word `more` with `less` using the sed command if you type 3g or nth line preceded by g, you will replace the word on the nth line only. In this case, only`g` will replace every occurrence of the word.
This is again a basic example of a sed command, its more if you go deeper, its more than a tool, its kind of a text-editor for wizards ;) 

### awk
awk or Aho, Weinberger, and Kernighan XD are the names of the developers of this application. This is another mind-blowing tool that can programmatically do a lot of stuff. This is like a programming language to a whole new level that can extrapolate and extract data from files and other forms of data. This is quite a great option if you want to quite neatly do something. It has great support libraries and functions that can even perform complex mathematical and scientific operations.

![possh.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626180322400/hWkEVhPl_.png)

These are the topics for separate articles because it is insufficient to explain everything at once.

### cat / tac / head / tail

CAT or concatenate is a tool used for printing out files, create files, sorting the contents of files, editing files, and a plethora of stuff. This command is generally used for printing the file but there is more to it like creating a file directly in the terminal, merging two files, and a ton of other operations. 

TAC or reverse of CAT is a tool used for everything that CAT can do but in reverse:) This is a wired tool but still quite useful sometimes.

Head is a tool that will print or edit the text in the first 10 lines of the file, it can be used to extrapolate multiple files with similar content. 
Tail is a tool that will print or edit the text in the last 10 lines of the file, it can be used just like head but for the last few lines.

![possh.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626180451092/Z5VUpIxCm.png)

It turns out, you can not only print the first or last 10 lines but n lines by passing the -n as an argument, there is a ton of stuff to discover, this just drives me crazy.

### cURL
cURL or client URL is a tool that can be used to transfer data via various network protocols. You might not believe but it is used in cars, televisions, routers, and other embedded systems for exchanging relevant data via appropriate protocols. 

![possh.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626181263235/JPocJwoEd.png)

This example depicts how we can fetch data from an API using cURL and extract data in JSON format and use it for relevant tasks.
This is again one of the best utility out there as it becomes quite remarkable and vintage. Despite being almost 30 years old, it shines bright in the tech world.

### find
Find as the name suggests it is used to find files among the folders and directories in a file system. it becomes quite helpful in complex projects where the directory structure is deep and large. 
![possh.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626181386566/zpP9Yaom4.png)

The command `find *.txt` finds every txt file available on this directory. As simple as it can get. This is surely looking silly and idiotic but it finds its glory in large and complicated codebases. 

### bc
bc or basic calculator is a utility tool for performing mathematical and arithmetical operations in the terminal, this commands gets integrated with other commands such as awk really well and can be used for further extending the limits of what the command line development can do.


![possh.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626182601487/z8X4KeDGG.png)

AWW! I could hear the excitement. That just added new dimensions into BASH. Just creativity and resonance to anything is the limit here. I am using  [REPL.IT](http://repl.it/)  here for using bash as I do not have it on my windows machine :( But that command is truly insane.

### wc
wc or word count is a utility tool for counting and analyzing the size or count of characters, words, lines, or files in a given file structure. This is quite a handy tool for monitoring and keeping track of a system, also for general development purposes.

![possh.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626182319602/r8UidHV2z.png)

The above command prints out the word and lines in the provided file. This command `wc` can even compute the size of files and even more properties of files.
Those were some of the quite powerful commands, tools, or utilities in BASH/shell. There are plenty of other commands not covered here because this an extremely large topic and even making separate articles or resources there will certainly and surely be some things that will get missed out, that's the beauty Linux or in general Computer Science has.
Ok, that was a lot, but I hope you got some insights for learning more BASH or Linux in general. This is a wide topic and can't be covered entirely in a single article. 

Now that is it from this part, everything cannot be covered in any number of parts but at least it will help someone to get started in BASH scripting and its specifications for development. Have a Blast learning BASH. Happy Coding :)