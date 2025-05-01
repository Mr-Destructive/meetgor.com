---
article_html: "<h2>Introduction</h2>\n<p>Are you bored of writing all the keymaps
  from vimscript to lua? Try the below function to create all your keymaps to lua
  equivalent maps in Neovim.</p>\n<p>Take your vimscript keymaps and put them in lua
  don't write any lua for it ;)</p>\n<h2>The Lua Function</h2>\n<p>The below-provided
  snippet is a lua function that takes in a table of strings(list of strings), the
  strings will be your keymaps. The function then maps these keymaps using lua functions.
  You don't have to type out all the keymaps by yourself. It can also print out the
  lua equivalent function calls required to map the existing keymaps from vimscript
  to lua runtime in Neovim. Though it won't handle all the options, we have passed
  in a default value to the keymap.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kr\">function</span> <span class=\"nf\">key_mapper</span><span class=\"p\">(</span><span
  class=\"n\">keymaps</span><span class=\"p\">)</span>\n  <span class=\"kr\">for</span>
  <span class=\"n\">_</span><span class=\"p\">,</span> <span class=\"n\">keymap</span>
  <span class=\"kr\">in</span> <span class=\"nb\">ipairs</span><span class=\"p\">(</span><span
  class=\"n\">keymaps</span><span class=\"p\">)</span> <span class=\"kr\">do</span>\n
  \   <span class=\"kd\">local</span> <span class=\"n\">mode</span> <span class=\"o\">=</span>
  <span class=\"n\">keymap</span><span class=\"p\">:</span><span class=\"n\">sub</span><span
  class=\"p\">(</span><span class=\"mi\">1</span><span class=\"p\">,</span><span class=\"mi\">1</span><span
  class=\"p\">)</span>\n    <span class=\"kd\">local</span> <span class=\"n\">delimiter</span>
  <span class=\"o\">=</span> <span class=\"s2\">&quot; &quot;</span>\n    <span class=\"kd\">local</span>
  <span class=\"n\">lhs</span> <span class=\"o\">=</span> <span class=\"s1\">&#39;&#39;</span>\n
  \   <span class=\"kd\">local</span> <span class=\"n\">rhs_parts</span> <span class=\"o\">=</span>
  <span class=\"p\">{}</span>\n    <span class=\"kd\">local</span> <span class=\"n\">m</span>
  <span class=\"o\">=</span> <span class=\"mi\">0</span>\n    <span class=\"kd\">local</span>
  <span class=\"n\">options</span> <span class=\"o\">=</span> <span class=\"p\">{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">}</span>\n    <span class=\"kr\">for</span> <span class=\"n\">matches</span>
  <span class=\"kr\">in</span> <span class=\"p\">(</span><span class=\"n\">keymap</span><span
  class=\"o\">..</span><span class=\"n\">delimiter</span><span class=\"p\">):</span><span
  class=\"n\">gmatch</span><span class=\"p\">(</span><span class=\"s2\">&quot;(.-)&quot;</span><span
  class=\"o\">..</span><span class=\"n\">delimiter</span><span class=\"p\">)</span>
  <span class=\"kr\">do</span>\n      <span class=\"kr\">if</span> <span class=\"n\">m</span>
  <span class=\"o\">==</span> <span class=\"mi\">1</span> <span class=\"kr\">then</span>\n
  \       <span class=\"n\">lhs</span> <span class=\"o\">=</span> <span class=\"n\">matches</span>\n
  \     <span class=\"kr\">end</span>\n      <span class=\"kr\">if</span> <span class=\"n\">m</span>
  <span class=\"o\">&gt;=</span> <span class=\"mi\">2</span> <span class=\"kr\">then</span>\n
  \       <span class=\"nb\">table.insert</span><span class=\"p\">(</span><span class=\"n\">rhs_parts</span><span
  class=\"p\">,</span> <span class=\"n\">matches</span><span class=\"p\">)</span>\n
  \     <span class=\"kr\">end</span>\n      <span class=\"n\">m</span> <span class=\"o\">=</span>
  <span class=\"n\">m</span> <span class=\"o\">+</span> <span class=\"mi\">1</span>\n
  \   <span class=\"kr\">end</span>\n    <span class=\"n\">rhs</span> <span class=\"o\">=</span>
  <span class=\"s1\">&#39;&#39;</span>\n    <span class=\"kr\">for</span> <span class=\"n\">_</span><span
  class=\"p\">,</span> <span class=\"n\">p</span> <span class=\"kr\">in</span> <span
  class=\"nb\">ipairs</span><span class=\"p\">(</span><span class=\"n\">rhs_parts</span><span
  class=\"p\">)</span> <span class=\"kr\">do</span>\n      <span class=\"n\">rhs</span>
  <span class=\"o\">=</span> <span class=\"n\">rhs</span> <span class=\"o\">..</span>
  <span class=\"s2\">&quot; &quot;</span> <span class=\"o\">..</span> <span class=\"n\">p</span>\n
  \   <span class=\"kr\">end</span>\n    <span class=\"c1\">--print(&quot;vim.keymap.set(&quot;..&quot;\\&#39;&quot;..mode..&quot;\\&#39;&quot;..&quot;,
  &quot;..&quot;\\&#39;&quot;..lhs..&quot;\\&#39;&quot;..&quot;, &quot;..&quot;\\&#39;&quot;..rhs..&quot;\\&#39;&quot;..&quot;,
  &quot;..vim.inspect(options)..&quot;)&quot;)</span>\n    <span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">keymap</span><span class=\"p\">.</span><span
  class=\"n\">set</span><span class=\"p\">(</span><span class=\"n\">mode</span><span
  class=\"p\">,</span> <span class=\"n\">lhs</span><span class=\"p\">,</span> <span
  class=\"n\">rhs</span><span class=\"p\">,</span> <span class=\"n\">options</span><span
  class=\"p\">)</span>\n  <span class=\"kr\">end</span>\n<span class=\"kr\">end</span>\n</pre></div>\n\n</pre>\n\n<p>You
  can uncomment the print statement <strong>once</strong> to grab the keymaps and
  paste them into the config file. If you leave it uncommented, it might print every
  time you open up a new neovim instance. The function can be called like below:</p>\n<pre
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">key_mapper</span><span class=\"p\">({</span>\n  <span class=\"s1\">&#39;nnoremap
  cpp :!c++ % -o %:r &amp;&amp; %:r&lt;CR&gt;i&#39;</span><span class=\"p\">,</span>\n
  \ <span class=\"s1\">&#39;nnoremap c, :!gcc % -o %:r &amp;&amp; %:r&lt;CR&gt;&#39;</span><span
  class=\"p\">,</span>\n  <span class=\"s1\">&#39;nnoremap py :!python %&lt;cr&gt;&#39;</span><span
  class=\"p\">,</span>\n  <span class=\"s1\">&#39;nnoremap go :!go run %&lt;cr&gt;&#39;</span><span
  class=\"p\">,</span>\n  <span class=\"s1\">&#39;nnoremap sh :!bash %&lt;CR&gt;&#39;</span>\n<span
  class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657559501/blog-media/neovim-lua-keymapper.gif\"
  alt=\"Keymapper demonstration\" /></p>\n<p>We pass in a table of strings, these
  strings are just the vimscript keymaps. This function call will then map the keymaps
  into equivalent lua maps. You can customize it as per your needs.</p>\n<p>For further
  references, you can check out my <a href=\"https://github.com/Mr-Destructive/dotfiles\">dotfiles</a>
  on GitHub.</p>\n<h2>How the function works</h2>\n<p>The function is simply a text
  scrapping from lua strings. We extract the first character in the string for the
  mode, grab the strings which are space-separated and finally sort out which are
  lhs and rhs sides of the maps.</p>\n<p>We iterate over the table in lua with the
  help of ipairs function which allows us to iterate over an ordered list of items
  in a table. Using the gmatch function, we find a pattern to split the string with
  the space as the delimiter. Thereby, we can have separate sets of strings identified
  as rhs and lhs. We can store them in variables as strings as the lua functions require
  them as strings.</p>\n<p>We simply add those variables into the <a href=\"https://neovim.io/doc/user/lua.html#:~:text=set(%7Bmode%7D%2C%20%7Blhs%7D%2C%20%7Brhs%7D%2C%20%7Bopts%7D)%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20*vim.keymap.set()*\">vim.keymap.set</a>
  or <a href=\"https://neovim.io/doc/user/api.html#nvim_set_keymap():~:text=nvim_set_keymap(%7Bmode%7D%2C%20%7Blhs%7D%2C%20%7Brhs%7D%2C%20%7B*opts%7D)%20%20%20%20%20%20%20%20%20%20%20%20%20*nvim_set_keymap()*\">vim.api.nvim_set_keymap</a>
  functions. We by default set the value of <code>{noremap: True}</code> to avoid
  teh recursive mapping of the keys. These option parameter is the one which needs
  to be a bit more dynamic in terms of wide variety of keymaps.</p>\n<p>So, this is
  how we can convert the vimscript keymaps to lua in Neovim. Hope you found this useful.
  Thanks for reading. Happy Viming :)</p>\n<div class='prevnext'>\n    <style type='text/css'>\n
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
  \   </style>\n<pre><code>&lt;a class='prev' href='/golang-test-output-json'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Golang:
  Test Output JSON&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a class='next' href='/golang-sort-package-basic'&gt;\n\n
  \   &lt;div class='prevnext-text'&gt;\n        &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n
  \       &lt;p class='prevnext-title'&gt;Golang: Sort Package Introduction&lt;/p&gt;\n
  \   &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0
  0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
cover: ''
date: 2022-07-11
datetime: 2022-07-11 00:00:00+00:00
description: Takeout the vimscript keymaps into lua with a single function call in
  Neovim
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/tils/2022-07-11-Neovim-Vim-keymaper.md
html: "<h2>Introduction</h2>\n<p>Are you bored of writing all the keymaps from vimscript
  to lua? Try the below function to create all your keymaps to lua equivalent maps
  in Neovim.</p>\n<p>Take your vimscript keymaps and put them in lua don't write any
  lua for it ;)</p>\n<h2>The Lua Function</h2>\n<p>The below-provided snippet is a
  lua function that takes in a table of strings(list of strings), the strings will
  be your keymaps. The function then maps these keymaps using lua functions. You don't
  have to type out all the keymaps by yourself. It can also print out the lua equivalent
  function calls required to map the existing keymaps from vimscript to lua runtime
  in Neovim. Though it won't handle all the options, we have passed in a default value
  to the keymap.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kr\">function</span> <span class=\"nf\">key_mapper</span><span class=\"p\">(</span><span
  class=\"n\">keymaps</span><span class=\"p\">)</span>\n  <span class=\"kr\">for</span>
  <span class=\"n\">_</span><span class=\"p\">,</span> <span class=\"n\">keymap</span>
  <span class=\"kr\">in</span> <span class=\"nb\">ipairs</span><span class=\"p\">(</span><span
  class=\"n\">keymaps</span><span class=\"p\">)</span> <span class=\"kr\">do</span>\n
  \   <span class=\"kd\">local</span> <span class=\"n\">mode</span> <span class=\"o\">=</span>
  <span class=\"n\">keymap</span><span class=\"p\">:</span><span class=\"n\">sub</span><span
  class=\"p\">(</span><span class=\"mi\">1</span><span class=\"p\">,</span><span class=\"mi\">1</span><span
  class=\"p\">)</span>\n    <span class=\"kd\">local</span> <span class=\"n\">delimiter</span>
  <span class=\"o\">=</span> <span class=\"s2\">&quot; &quot;</span>\n    <span class=\"kd\">local</span>
  <span class=\"n\">lhs</span> <span class=\"o\">=</span> <span class=\"s1\">&#39;&#39;</span>\n
  \   <span class=\"kd\">local</span> <span class=\"n\">rhs_parts</span> <span class=\"o\">=</span>
  <span class=\"p\">{}</span>\n    <span class=\"kd\">local</span> <span class=\"n\">m</span>
  <span class=\"o\">=</span> <span class=\"mi\">0</span>\n    <span class=\"kd\">local</span>
  <span class=\"n\">options</span> <span class=\"o\">=</span> <span class=\"p\">{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">}</span>\n    <span class=\"kr\">for</span> <span class=\"n\">matches</span>
  <span class=\"kr\">in</span> <span class=\"p\">(</span><span class=\"n\">keymap</span><span
  class=\"o\">..</span><span class=\"n\">delimiter</span><span class=\"p\">):</span><span
  class=\"n\">gmatch</span><span class=\"p\">(</span><span class=\"s2\">&quot;(.-)&quot;</span><span
  class=\"o\">..</span><span class=\"n\">delimiter</span><span class=\"p\">)</span>
  <span class=\"kr\">do</span>\n      <span class=\"kr\">if</span> <span class=\"n\">m</span>
  <span class=\"o\">==</span> <span class=\"mi\">1</span> <span class=\"kr\">then</span>\n
  \       <span class=\"n\">lhs</span> <span class=\"o\">=</span> <span class=\"n\">matches</span>\n
  \     <span class=\"kr\">end</span>\n      <span class=\"kr\">if</span> <span class=\"n\">m</span>
  <span class=\"o\">&gt;=</span> <span class=\"mi\">2</span> <span class=\"kr\">then</span>\n
  \       <span class=\"nb\">table.insert</span><span class=\"p\">(</span><span class=\"n\">rhs_parts</span><span
  class=\"p\">,</span> <span class=\"n\">matches</span><span class=\"p\">)</span>\n
  \     <span class=\"kr\">end</span>\n      <span class=\"n\">m</span> <span class=\"o\">=</span>
  <span class=\"n\">m</span> <span class=\"o\">+</span> <span class=\"mi\">1</span>\n
  \   <span class=\"kr\">end</span>\n    <span class=\"n\">rhs</span> <span class=\"o\">=</span>
  <span class=\"s1\">&#39;&#39;</span>\n    <span class=\"kr\">for</span> <span class=\"n\">_</span><span
  class=\"p\">,</span> <span class=\"n\">p</span> <span class=\"kr\">in</span> <span
  class=\"nb\">ipairs</span><span class=\"p\">(</span><span class=\"n\">rhs_parts</span><span
  class=\"p\">)</span> <span class=\"kr\">do</span>\n      <span class=\"n\">rhs</span>
  <span class=\"o\">=</span> <span class=\"n\">rhs</span> <span class=\"o\">..</span>
  <span class=\"s2\">&quot; &quot;</span> <span class=\"o\">..</span> <span class=\"n\">p</span>\n
  \   <span class=\"kr\">end</span>\n    <span class=\"c1\">--print(&quot;vim.keymap.set(&quot;..&quot;\\&#39;&quot;..mode..&quot;\\&#39;&quot;..&quot;,
  &quot;..&quot;\\&#39;&quot;..lhs..&quot;\\&#39;&quot;..&quot;, &quot;..&quot;\\&#39;&quot;..rhs..&quot;\\&#39;&quot;..&quot;,
  &quot;..vim.inspect(options)..&quot;)&quot;)</span>\n    <span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">keymap</span><span class=\"p\">.</span><span
  class=\"n\">set</span><span class=\"p\">(</span><span class=\"n\">mode</span><span
  class=\"p\">,</span> <span class=\"n\">lhs</span><span class=\"p\">,</span> <span
  class=\"n\">rhs</span><span class=\"p\">,</span> <span class=\"n\">options</span><span
  class=\"p\">)</span>\n  <span class=\"kr\">end</span>\n<span class=\"kr\">end</span>\n</pre></div>\n\n</pre>\n\n<p>You
  can uncomment the print statement <strong>once</strong> to grab the keymaps and
  paste them into the config file. If you leave it uncommented, it might print every
  time you open up a new neovim instance. The function can be called like below:</p>\n<pre
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">key_mapper</span><span class=\"p\">({</span>\n  <span class=\"s1\">&#39;nnoremap
  cpp :!c++ % -o %:r &amp;&amp; %:r&lt;CR&gt;i&#39;</span><span class=\"p\">,</span>\n
  \ <span class=\"s1\">&#39;nnoremap c, :!gcc % -o %:r &amp;&amp; %:r&lt;CR&gt;&#39;</span><span
  class=\"p\">,</span>\n  <span class=\"s1\">&#39;nnoremap py :!python %&lt;cr&gt;&#39;</span><span
  class=\"p\">,</span>\n  <span class=\"s1\">&#39;nnoremap go :!go run %&lt;cr&gt;&#39;</span><span
  class=\"p\">,</span>\n  <span class=\"s1\">&#39;nnoremap sh :!bash %&lt;CR&gt;&#39;</span>\n<span
  class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657559501/blog-media/neovim-lua-keymapper.gif\"
  alt=\"Keymapper demonstration\" /></p>\n<p>We pass in a table of strings, these
  strings are just the vimscript keymaps. This function call will then map the keymaps
  into equivalent lua maps. You can customize it as per your needs.</p>\n<p>For further
  references, you can check out my <a href=\"https://github.com/Mr-Destructive/dotfiles\">dotfiles</a>
  on GitHub.</p>\n<h2>How the function works</h2>\n<p>The function is simply a text
  scrapping from lua strings. We extract the first character in the string for the
  mode, grab the strings which are space-separated and finally sort out which are
  lhs and rhs sides of the maps.</p>\n<p>We iterate over the table in lua with the
  help of ipairs function which allows us to iterate over an ordered list of items
  in a table. Using the gmatch function, we find a pattern to split the string with
  the space as the delimiter. Thereby, we can have separate sets of strings identified
  as rhs and lhs. We can store them in variables as strings as the lua functions require
  them as strings.</p>\n<p>We simply add those variables into the <a href=\"https://neovim.io/doc/user/lua.html#:~:text=set(%7Bmode%7D%2C%20%7Blhs%7D%2C%20%7Brhs%7D%2C%20%7Bopts%7D)%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20*vim.keymap.set()*\">vim.keymap.set</a>
  or <a href=\"https://neovim.io/doc/user/api.html#nvim_set_keymap():~:text=nvim_set_keymap(%7Bmode%7D%2C%20%7Blhs%7D%2C%20%7Brhs%7D%2C%20%7B*opts%7D)%20%20%20%20%20%20%20%20%20%20%20%20%20*nvim_set_keymap()*\">vim.api.nvim_set_keymap</a>
  functions. We by default set the value of <code>{noremap: True}</code> to avoid
  teh recursive mapping of the keys. These option parameter is the one which needs
  to be a bit more dynamic in terms of wide variety of keymaps.</p>\n<p>So, this is
  how we can convert the vimscript keymaps to lua in Neovim. Hope you found this useful.
  Thanks for reading. Happy Viming :)</p>\n<div class='prevnext'>\n    <style type='text/css'>\n
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
  \   </style>\n<pre><code>&lt;a class='prev' href='/golang-test-output-json'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Golang:
  Test Output JSON&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a class='next' href='/golang-sort-package-basic'&gt;\n\n
  \   &lt;div class='prevnext-text'&gt;\n        &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n
  \       &lt;p class='prevnext-title'&gt;Golang: Sort Package Introduction&lt;/p&gt;\n
  \   &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0
  0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
long_description: 'Are you bored of writing all the keymaps from vimscript to lua?
  Try the below function to create all your keymaps to lua equivalent maps in Neovim.
  Take your vimscript keymaps and put them in lua don The below-provided snippet is
  a lua function that '
now: 2025-05-01 18:11:33.315135
path: blog/tils/2022-07-11-Neovim-Vim-keymaper.md
slug: vimscript-to-lua-keymapper
status: published-til
tags:
- vim
- lua
templateKey: til
title: Map Vimscript Keymaps to Lua with a single function
today: 2025-05-01
---

## Introduction

Are you bored of writing all the keymaps from vimscript to lua? Try the below function to create all your keymaps to lua equivalent maps in Neovim.

Take your vimscript keymaps and put them in lua don't write any lua for it ;)

## The Lua Function

The below-provided snippet is a lua function that takes in a table of strings(list of strings), the strings will be your keymaps. The function then maps these keymaps using lua functions. You don't have to type out all the keymaps by yourself. It can also print out the lua equivalent function calls required to map the existing keymaps from vimscript to lua runtime in Neovim. Though it won't handle all the options, we have passed in a default value to the keymap.

```lua
function key_mapper(keymaps)
  for _, keymap in ipairs(keymaps) do
    local mode = keymap:sub(1,1)
    local delimiter = " "
    local lhs = ''
    local rhs_parts = {}
    local m = 0
    local options = {noremap = true}
    for matches in (keymap..delimiter):gmatch("(.-)"..delimiter) do
      if m == 1 then
        lhs = matches
      end
      if m >= 2 then
        table.insert(rhs_parts, matches)
      end
      m = m + 1
    end
    rhs = ''
    for _, p in ipairs(rhs_parts) do
      rhs = rhs .. " " .. p
    end
    --print("vim.keymap.set(".."\'"..mode.."\'"..", ".."\'"..lhs.."\'"..", ".."\'"..rhs.."\'"..", "..vim.inspect(options)..")")
    vim.keymap.set(mode, lhs, rhs, options)
  end
end
```

You can uncomment the print statement **once** to grab the keymaps and paste them into the config file. If you leave it uncommented, it might print every time you open up a new neovim instance. The function can be called like below:

```lua
key_mapper({
  'nnoremap cpp :!c++ % -o %:r && %:r<CR>i',
  'nnoremap c, :!gcc % -o %:r && %:r<CR>',
  'nnoremap py :!python %<cr>',
  'nnoremap go :!go run %<cr>',
  'nnoremap sh :!bash %<CR>'
})
```

![Keymapper demonstration](https://res.cloudinary.com/techstructive-blog/image/upload/v1657559501/blog-media/neovim-lua-keymapper.gif)

We pass in a table of strings, these strings are just the vimscript keymaps. This function call will then map the keymaps into equivalent lua maps. You can customize it as per your needs.

For further references, you can check out my [dotfiles](https://github.com/Mr-Destructive/dotfiles) on GitHub.

## How the function works

The function is simply a text scrapping from lua strings. We extract the first character in the string for the mode, grab the strings which are space-separated and finally sort out which are lhs and rhs sides of the maps.

We iterate over the table in lua with the help of ipairs function which allows us to iterate over an ordered list of items in a table. Using the gmatch function, we find a pattern to split the string with the space as the delimiter. Thereby, we can have separate sets of strings identified as rhs and lhs. We can store them in variables as strings as the lua functions require them as strings.

We simply add those variables into the [vim.keymap.set](https://neovim.io/doc/user/lua.html#:~:text=set(%7Bmode%7D%2C%20%7Blhs%7D%2C%20%7Brhs%7D%2C%20%7Bopts%7D)%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20%20*vim.keymap.set()*) or [vim.api.nvim_set_keymap](https://neovim.io/doc/user/api.html#nvim_set_keymap():~:text=nvim_set_keymap(%7Bmode%7D%2C%20%7Blhs%7D%2C%20%7Brhs%7D%2C%20%7B*opts%7D)%20%20%20%20%20%20%20%20%20%20%20%20%20*nvim_set_keymap()*) functions. We by default set the value of `{noremap: True}` to avoid teh recursive mapping of the keys. These option parameter is the one which needs to be a bit more dynamic in terms of wide variety of keymaps.

So, this is how we can convert the vimscript keymaps to lua in Neovim. Hope you found this useful. Thanks for reading. Happy Viming :)
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
    
    <a class='prev' href='/golang-test-output-json'>
    
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M13.5 8.25L9.75 12L13.5 15.75" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"> </path>
        </svg>
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>prev</p>
            <p class='prevnext-title'>Golang: Test Output JSON</p>
        </div>
    </a>
    
    <a class='next' href='/golang-sort-package-basic'>
    
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>next</p>
            <p class='prevnext-title'>Golang: Sort Package Introduction</p>
        </div>
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.5 15.75L14.25 12L10.5 8.25" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path>
        </svg>
    </a>
  </div>