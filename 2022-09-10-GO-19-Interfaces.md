---
article_html: "<h2>Introduction</h2>\n<p>In the 19th post of the series, we will be
  taking a look into interfaces in golang. Interfaces allow us to create function
  signatures common to different structs or types. So, we can allow multiple structs
  to have a common interface(method) that can have different implementations.</p>\n<h2>What
  are Interfaces</h2>\n<p>Interface as the name suggests is a way to create methods
  that are common to different structures or types but can have different implementations.
  It's an interface to define the method or function signatures but not the implementation.
  Let's take an example of <code>Laptop</code> and <code>Phone</code> having the functionality
  of wifi. We can connect to wifi more or the less in a similar way on both devices.
  The implementation behind the functionality might be different but they share the
  same operation. The WiFi can act as an interface for both devices to connect to
  the internet.</p>\n<h2>Define an Interface</h2>\n<p>To declare an interface in golang,
  we can use the <code>interface</code> keyword in golang. An interface is created
  with the type keyword, providing the name of the interface and defining the function
  declaration. Inside the interface which acts as a struct of general method signatures.
  The method signatures usually consist of the name of the function with its parameters
  if any and the return type of the function.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Player</span><span class=\"w\"> </span><span
  class=\"kd\">struct</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">name</span><span
  class=\"w\">   </span><span class=\"kt\">string</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">health</span><span class=\"w\"> </span><span
  class=\"kt\">int</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Mob</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">name</span><span class=\"w\">     </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">health</span><span
  class=\"w\">   </span><span class=\"kt\">int</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">is_passive</span><span class=\"w\"> </span><span
  class=\"kt\">bool</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Creature</span><span class=\"w\"> </span><span class=\"kd\">interface</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">intro</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">//attack() int</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">//heal() int</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">player</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">Player</span><span class=\"p\">{</span><span
  class=\"nx\">name</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s\">&quot;Steve&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">mob</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">Mob</span><span
  class=\"p\">{</span><span class=\"nx\">name</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Zombie&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">player</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">mob</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run main.go\n\n{Steve
  0}\n{Zombie 0 false}\n</pre></div>\n\n</pre>\n\n<p>In this above example, we have
  created an interface called <code>Creature</code>. There are a few structs that
  we have to define like <code>Player</code> and <code>Mob</code>, these two methods
  have a few attributes like <code>name</code> as <code>string</code> and <code>health</code>
  as <code>int</code> which are common in both structs but the <code>Mob</code> struct
  has an additional attribute <code>is_passive</code> as a <code>boolean</code> value.
  The <code>Creature</code> is an interface that will define certain function signatures,
  here we have declared <code>intro</code>, <code>attack</code>, and <code>heal</code>
  as the methods bound to the Creature interface. This means, that any object which
  satisfies the Creature interface, can define the methods associated with the interface.</p>\n<h2>Defining
  Interfaces</h2>\n<p>Once we have declared the interface method signatures, we can
  move into defining the functionality of these methods depending on the struct. If
  we want a different working method for different types of struct objects passed
  we can define those for each type of struct that we want. Here, we have two types
  of structs namely <code>Creature</code> and <code>Mob</code>, based on the struct
  we can define the <code>intro</code> method for them individually.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Creature</span><span class=\"w\"> </span><span
  class=\"kd\">interface</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">intro</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">attack</span><span
  class=\"p\">(</span><span class=\"o\">*</span><span class=\"kt\">int</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">type</span><span class=\"w\"> </span><span class=\"nx\">Player</span><span
  class=\"w\"> </span><span class=\"kd\">struct</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">name</span><span class=\"w\">   </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">health</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Mob</span><span class=\"w\"> </span><span
  class=\"kd\">struct</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">name</span><span
  class=\"w\">     </span><span class=\"kt\">string</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">health</span><span class=\"w\">   </span><span
  class=\"kt\">int</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">category</span><span class=\"w\"> </span><span class=\"kt\">bool</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"p\">(</span><span
  class=\"nx\">p</span><span class=\"w\"> </span><span class=\"nx\">Player</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"nx\">intro</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Player has spawned&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">p</span><span
  class=\"p\">.</span><span class=\"nx\">name</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"p\">(</span><span class=\"nx\">p</span><span class=\"w\">
  </span><span class=\"nx\">Player</span><span class=\"p\">)</span><span class=\"w\">
  </span><span class=\"nx\">attack</span><span class=\"p\">(</span><span class=\"nx\">m_health</span><span
  class=\"w\"> </span><span class=\"o\">*</span><span class=\"kt\">int</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Player has attacked!&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"o\">*</span><span class=\"nx\">m_health</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">m_health</span><span
  class=\"w\"> </span><span class=\"o\">-</span><span class=\"w\"> </span><span class=\"mi\">50</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">return</span><span
  class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">m_health</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"p\">(</span><span
  class=\"nx\">m</span><span class=\"w\"> </span><span class=\"nx\">Mob</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"nx\">intro</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;A wild %s has appeared!\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">m</span><span class=\"p\">.</span><span
  class=\"nx\">name</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"nx\">m</span><span class=\"p\">.</span><span class=\"nx\">name</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"p\">(</span><span
  class=\"nx\">m</span><span class=\"w\"> </span><span class=\"nx\">Mob</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"nx\">attack</span><span
  class=\"p\">(</span><span class=\"nx\">p_health</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"kt\">int</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;%s has attacked you! -%d\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">m</span><span class=\"p\">.</span><span
  class=\"nx\">name</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"mi\">30</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"o\">*</span><span class=\"nx\">p_health</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">p_health</span><span class=\"w\"> </span><span class=\"o\">-</span><span
  class=\"w\"> </span><span class=\"mi\">30</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">p_health</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">player</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">Player</span><span class=\"p\">{</span><span
  class=\"nx\">name</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s\">&quot;Steve&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">health</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"mi\">100</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">mob</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">Mob</span><span
  class=\"p\">{</span><span class=\"nx\">name</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Zombie&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">health</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">140</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">player</span><span class=\"p\">.</span><span class=\"nx\">intro</span><span
  class=\"p\">())</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">mob</span><span class=\"p\">.</span><span
  class=\"nx\">intro</span><span class=\"p\">())</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">mob</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">player</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">player</span><span class=\"p\">.</span><span class=\"nx\">attack</span><span
  class=\"p\">(</span><span class=\"o\">&amp;</span><span class=\"nx\">mob</span><span
  class=\"p\">.</span><span class=\"nx\">health</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">mob</span><span class=\"p\">.</span><span class=\"nx\">attack</span><span
  class=\"p\">(</span><span class=\"o\">&amp;</span><span class=\"nx\">player</span><span
  class=\"p\">.</span><span class=\"nx\">health</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">mob</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">player</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run main.go\n\nPlayer
  has spawned\nSteve\nA wild Zombie has appeared!\nZombie\n{Zombie 140 false}\n{Steve
  100}\nPlayer has attacked!\n90\nZombie has attacked you! -30\n70\n{Zombie 90 false}\n{Steve
  70}\n</pre></div>\n\n</pre>\n\n<p>As we can see, the method <code>intro()</code>
  is bound to both the struct depending on what struct signature is associated with
  the method. The method <code>intro</code> takes in the object struct associated
  as per the call and returns <code>string</code> as defined in the method signature.</p>\n<p>The
  <code>attack</code> method in the <code>Creature</code> interface is also implemented
  separately for the two structs. For the <code>Player</code> method, we simply take
  in a pointer to an integer and return an <code>int</code>. The parameter is the
  pointer to the mob health, and it returns the modified health. We take in a pointer
  to the mob or player health so as to parse in the actual value and not the copy
  of the value. If we modify the value, we want to reflect those changes in the actual
  object. So that is how we can use interfaces to construct dynamic operations on
  objects as well as different types of structs.</p>\n<p>We have seen a simple example
  of how to declare and define interfaces for given type structs. Also, we can pass
  by value as well as by pointers so as to define the behavior of the method whether
  to dynamically modify or change the values of the object associated with it.</p>\n<h2>Examples
  of Interfaces</h2>\n<p>There are quite some use cases of interfaces, in object-oriented
  programming, the above example fits the polymorphism feature quite well. The ability
  to reuse certain method signatures and define the functions as per requirement brings
  flexibility to the code structure.</p>\n<p>We will see a few examples for understanding
  how we can use interfaces in various ways.</p>\n<h3>Type Switch Interface</h3>\n<p>We
  can use an empty interface to check for the type of variable we have parsed. Using
  this empty interface we can create a kind of dynamic parameter to a function.</p>\n<pre
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strconv&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">parse_int</span><span class=\"p\">(</span><span class=\"nx\">n</span><span
  class=\"w\"> </span><span class=\"kd\">interface</span><span class=\"p\">{})</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">switch</span><span class=\"w\"> </span><span class=\"nx\">n</span><span
  class=\"p\">.(</span><span class=\"kd\">type</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">case</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"p\">:</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"p\">(</span><span
  class=\"nx\">n</span><span class=\"p\">).(</span><span class=\"kt\">int</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"o\">*</span><span class=\"w\">
  </span><span class=\"p\">(</span><span class=\"nx\">n</span><span class=\"p\">).(</span><span
  class=\"kt\">int</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">case</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">:</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">s</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">strconv</span><span class=\"p\">.</span><span
  class=\"nx\">Atoi</span><span class=\"p\">(</span><span class=\"nx\">n</span><span
  class=\"p\">.(</span><span class=\"kt\">string</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">return</span><span
  class=\"w\"> </span><span class=\"nx\">s</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">case</span><span class=\"w\"> </span><span
  class=\"kt\">float64</span><span class=\"p\">:</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"nb\">int</span><span class=\"p\">(</span><span class=\"nx\">n</span><span
  class=\"p\">.(</span><span class=\"kt\">float64</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">default</span><span
  class=\"p\">:</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">n</span><span
  class=\"p\">.(</span><span class=\"kt\">int</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">num</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">parse_int</span><span
  class=\"p\">(</span><span class=\"mi\">4</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">num</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">num</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">parse_int</span><span class=\"p\">(</span><span
  class=\"s\">&quot;4&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">num</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">num</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">parse_int</span><span class=\"p\">(</span><span
  class=\"mf\">4.1243</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">num</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run main.go\n\n16\n4\n4\n</pre></div>\n\n</pre>\n\n<p>Here,
  we can see we have an interface as a parameter to the function <code>parse_int</code>,
  the return type is <code>int</code>, so the incoming parameter can be any valid
  type. But if we don't convert the given type into an appropriate int, it will result
  in an error as we are returning the int value of the parsed parameter. We are taking
  the parameter as <code>interface{}</code> which is an empty interface, this will
  contain the parameter parsed as an interface type. That's why we need to convert
  the interface object into an int or the parsed type of the interface.</p>\n<h3>Interface
  Slice</h3>\n<p>We can even create a slice of interfaces, which means we can initialize
  or group together various objects of different structs in a single slice. This might
  be helpful for calling functions associated with different objects via the interface
  very easily and in a much cleaner way.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Creature</span><span class=\"w\"> </span><span
  class=\"kd\">interface</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">intro</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">type</span><span class=\"w\"> </span><span class=\"nx\">Player</span><span
  class=\"w\"> </span><span class=\"kd\">struct</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">name</span><span class=\"w\">   </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">health</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Mob</span><span class=\"w\"> </span><span
  class=\"kd\">struct</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">name</span><span
  class=\"w\">     </span><span class=\"kt\">string</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">health</span><span class=\"w\">   </span><span
  class=\"kt\">int</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">category</span><span class=\"w\"> </span><span class=\"kt\">bool</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"p\">(</span><span
  class=\"nx\">p</span><span class=\"w\"> </span><span class=\"nx\">Player</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"nx\">intro</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Player has spawned&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">p</span><span
  class=\"p\">.</span><span class=\"nx\">name</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"p\">(</span><span class=\"nx\">m</span><span class=\"w\">
  </span><span class=\"nx\">Mob</span><span class=\"p\">)</span><span class=\"w\">
  </span><span class=\"nx\">intro</span><span class=\"p\">()</span><span class=\"w\">
  </span><span class=\"kt\">string</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">name</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">m</span><span
  class=\"p\">.</span><span class=\"nx\">name</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"s\">&quot;&quot;</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">name</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">m</span><span class=\"p\">.</span><span class=\"nx\">name</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"> </span><span class=\"k\">else</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">name</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Mob&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Printf</span><span class=\"p\">(</span><span class=\"s\">&quot;A wild
  %s has appeared!\\n&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">name</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"nx\">m</span><span class=\"p\">.</span><span class=\"nx\">name</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">entity</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"nx\">Creature</span><span class=\"p\">{</span><span class=\"nx\">Player</span><span
  class=\"p\">{},</span><span class=\"w\"> </span><span class=\"nx\">Mob</span><span
  class=\"p\">{},</span><span class=\"w\"> </span><span class=\"nx\">Mob</span><span
  class=\"p\">{},</span><span class=\"w\"> </span><span class=\"nx\">Player</span><span
  class=\"p\">{}}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">_</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">obj</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">entity</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">obj</span><span class=\"p\">.</span><span
  class=\"nx\">intro</span><span class=\"p\">())</span><span class=\"w\"></span>\n<span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>Player has spawned\nA wild
  Zombie has appeared!\nA wild Zombie has appeared!\nPlayer has spawned\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we can see that the entity variable is created as a slice of
  interfaces <code>Creature</code>, i.e. various objects associated with the <code>Creature</code>
  interface can be contained in a single slice. There are 2 instances of Player and
  Mob each in the slice. We can further iterate over the slice as a range-based loop
  and thereby the functions associated with the interfaces can be called. Here, we
  have called the <code>intro</code> function.</p>\n<p>So, there are a lot of things
  that can be done with interfaces, we can create multiple interfaces for a single
  object struct and nest interfaces. Based on the use case of the program, interfaces
  can be used to reduce the boilerplate code as well as improve the readability of
  the code.</p>\n<p>That's it from this part. Reference for all the code examples
  and commands can be found in the <a href=\"https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/interfaces/main.go\">100
  days of Golang</a> GitHub repository.</p>\n<h2>Conclusion</h2>\n<p>From this part
  of the series, we were able to understand the basics of interfaces using a few examples.
  We explored how interfaces can be used to bring in polymorphism in golang, also
  we can improve the readability of the code. The boilerplate code can be considerably
  reduced by using interfaces when dealing with structs and types. Hopefully, you
  found this post helpful and understood even the basics of interfaces in golang.
  Thank you for reading, if you have any queries, questions, or feedback, you can
  ping me on my social handles or in the comments. Happy Coding :)</p>\n"
cover: ''
date: 2022-09-10
datetime: 2022-09-10 00:00:00+00:00
description: Understanding the baiscs of interfaces in Golang. Exploring the concept
  of polymorphism in golang with the help of interfaces and structs.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2022-09-10-GO-19-Interfaces.md
html: "<h2>Introduction</h2>\n<p>In the 19th post of the series, we will be taking
  a look into interfaces in golang. Interfaces allow us to create function signatures
  common to different structs or types. So, we can allow multiple structs to have
  a common interface(method) that can have different implementations.</p>\n<h2>What
  are Interfaces</h2>\n<p>Interface as the name suggests is a way to create methods
  that are common to different structures or types but can have different implementations.
  It's an interface to define the method or function signatures but not the implementation.
  Let's take an example of <code>Laptop</code> and <code>Phone</code> having the functionality
  of wifi. We can connect to wifi more or the less in a similar way on both devices.
  The implementation behind the functionality might be different but they share the
  same operation. The WiFi can act as an interface for both devices to connect to
  the internet.</p>\n<h2>Define an Interface</h2>\n<p>To declare an interface in golang,
  we can use the <code>interface</code> keyword in golang. An interface is created
  with the type keyword, providing the name of the interface and defining the function
  declaration. Inside the interface which acts as a struct of general method signatures.
  The method signatures usually consist of the name of the function with its parameters
  if any and the return type of the function.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Player</span><span class=\"w\"> </span><span
  class=\"kd\">struct</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">name</span><span
  class=\"w\">   </span><span class=\"kt\">string</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">health</span><span class=\"w\"> </span><span
  class=\"kt\">int</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Mob</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">name</span><span class=\"w\">     </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">health</span><span
  class=\"w\">   </span><span class=\"kt\">int</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">is_passive</span><span class=\"w\"> </span><span
  class=\"kt\">bool</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Creature</span><span class=\"w\"> </span><span class=\"kd\">interface</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">intro</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">//attack() int</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">//heal() int</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">player</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">Player</span><span class=\"p\">{</span><span
  class=\"nx\">name</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s\">&quot;Steve&quot;</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">mob</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">Mob</span><span
  class=\"p\">{</span><span class=\"nx\">name</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Zombie&quot;</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">player</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">mob</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run main.go\n\n{Steve
  0}\n{Zombie 0 false}\n</pre></div>\n\n</pre>\n\n<p>In this above example, we have
  created an interface called <code>Creature</code>. There are a few structs that
  we have to define like <code>Player</code> and <code>Mob</code>, these two methods
  have a few attributes like <code>name</code> as <code>string</code> and <code>health</code>
  as <code>int</code> which are common in both structs but the <code>Mob</code> struct
  has an additional attribute <code>is_passive</code> as a <code>boolean</code> value.
  The <code>Creature</code> is an interface that will define certain function signatures,
  here we have declared <code>intro</code>, <code>attack</code>, and <code>heal</code>
  as the methods bound to the Creature interface. This means, that any object which
  satisfies the Creature interface, can define the methods associated with the interface.</p>\n<h2>Defining
  Interfaces</h2>\n<p>Once we have declared the interface method signatures, we can
  move into defining the functionality of these methods depending on the struct. If
  we want a different working method for different types of struct objects passed
  we can define those for each type of struct that we want. Here, we have two types
  of structs namely <code>Creature</code> and <code>Mob</code>, based on the struct
  we can define the <code>intro</code> method for them individually.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Creature</span><span class=\"w\"> </span><span
  class=\"kd\">interface</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">intro</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">attack</span><span
  class=\"p\">(</span><span class=\"o\">*</span><span class=\"kt\">int</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">type</span><span class=\"w\"> </span><span class=\"nx\">Player</span><span
  class=\"w\"> </span><span class=\"kd\">struct</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">name</span><span class=\"w\">   </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">health</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Mob</span><span class=\"w\"> </span><span
  class=\"kd\">struct</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">name</span><span
  class=\"w\">     </span><span class=\"kt\">string</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">health</span><span class=\"w\">   </span><span
  class=\"kt\">int</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">category</span><span class=\"w\"> </span><span class=\"kt\">bool</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"p\">(</span><span
  class=\"nx\">p</span><span class=\"w\"> </span><span class=\"nx\">Player</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"nx\">intro</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Player has spawned&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">p</span><span
  class=\"p\">.</span><span class=\"nx\">name</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"p\">(</span><span class=\"nx\">p</span><span class=\"w\">
  </span><span class=\"nx\">Player</span><span class=\"p\">)</span><span class=\"w\">
  </span><span class=\"nx\">attack</span><span class=\"p\">(</span><span class=\"nx\">m_health</span><span
  class=\"w\"> </span><span class=\"o\">*</span><span class=\"kt\">int</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Player has attacked!&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"o\">*</span><span class=\"nx\">m_health</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">m_health</span><span
  class=\"w\"> </span><span class=\"o\">-</span><span class=\"w\"> </span><span class=\"mi\">50</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">return</span><span
  class=\"w\"> </span><span class=\"o\">*</span><span class=\"nx\">m_health</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"p\">(</span><span
  class=\"nx\">m</span><span class=\"w\"> </span><span class=\"nx\">Mob</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"nx\">intro</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;A wild %s has appeared!\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">m</span><span class=\"p\">.</span><span
  class=\"nx\">name</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"nx\">m</span><span class=\"p\">.</span><span class=\"nx\">name</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"p\">(</span><span
  class=\"nx\">m</span><span class=\"w\"> </span><span class=\"nx\">Mob</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"nx\">attack</span><span
  class=\"p\">(</span><span class=\"nx\">p_health</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"kt\">int</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;%s has attacked you! -%d\\n&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">m</span><span class=\"p\">.</span><span
  class=\"nx\">name</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"mi\">30</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"o\">*</span><span class=\"nx\">p_health</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"o\">*</span><span
  class=\"nx\">p_health</span><span class=\"w\"> </span><span class=\"o\">-</span><span
  class=\"w\"> </span><span class=\"mi\">30</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"o\">*</span><span class=\"nx\">p_health</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">player</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">Player</span><span class=\"p\">{</span><span
  class=\"nx\">name</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s\">&quot;Steve&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">health</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"mi\">100</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">mob</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">Mob</span><span
  class=\"p\">{</span><span class=\"nx\">name</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Zombie&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">health</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">140</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">player</span><span class=\"p\">.</span><span class=\"nx\">intro</span><span
  class=\"p\">())</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">mob</span><span class=\"p\">.</span><span
  class=\"nx\">intro</span><span class=\"p\">())</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">mob</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">player</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">player</span><span class=\"p\">.</span><span class=\"nx\">attack</span><span
  class=\"p\">(</span><span class=\"o\">&amp;</span><span class=\"nx\">mob</span><span
  class=\"p\">.</span><span class=\"nx\">health</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">mob</span><span class=\"p\">.</span><span class=\"nx\">attack</span><span
  class=\"p\">(</span><span class=\"o\">&amp;</span><span class=\"nx\">player</span><span
  class=\"p\">.</span><span class=\"nx\">health</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">mob</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">player</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run main.go\n\nPlayer
  has spawned\nSteve\nA wild Zombie has appeared!\nZombie\n{Zombie 140 false}\n{Steve
  100}\nPlayer has attacked!\n90\nZombie has attacked you! -30\n70\n{Zombie 90 false}\n{Steve
  70}\n</pre></div>\n\n</pre>\n\n<p>As we can see, the method <code>intro()</code>
  is bound to both the struct depending on what struct signature is associated with
  the method. The method <code>intro</code> takes in the object struct associated
  as per the call and returns <code>string</code> as defined in the method signature.</p>\n<p>The
  <code>attack</code> method in the <code>Creature</code> interface is also implemented
  separately for the two structs. For the <code>Player</code> method, we simply take
  in a pointer to an integer and return an <code>int</code>. The parameter is the
  pointer to the mob health, and it returns the modified health. We take in a pointer
  to the mob or player health so as to parse in the actual value and not the copy
  of the value. If we modify the value, we want to reflect those changes in the actual
  object. So that is how we can use interfaces to construct dynamic operations on
  objects as well as different types of structs.</p>\n<p>We have seen a simple example
  of how to declare and define interfaces for given type structs. Also, we can pass
  by value as well as by pointers so as to define the behavior of the method whether
  to dynamically modify or change the values of the object associated with it.</p>\n<h2>Examples
  of Interfaces</h2>\n<p>There are quite some use cases of interfaces, in object-oriented
  programming, the above example fits the polymorphism feature quite well. The ability
  to reuse certain method signatures and define the functions as per requirement brings
  flexibility to the code structure.</p>\n<p>We will see a few examples for understanding
  how we can use interfaces in various ways.</p>\n<h3>Type Switch Interface</h3>\n<p>We
  can use an empty interface to check for the type of variable we have parsed. Using
  this empty interface we can create a kind of dynamic parameter to a function.</p>\n<pre
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strconv&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">parse_int</span><span class=\"p\">(</span><span class=\"nx\">n</span><span
  class=\"w\"> </span><span class=\"kd\">interface</span><span class=\"p\">{})</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">switch</span><span class=\"w\"> </span><span class=\"nx\">n</span><span
  class=\"p\">.(</span><span class=\"kd\">type</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">case</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"p\">:</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"p\">(</span><span
  class=\"nx\">n</span><span class=\"p\">).(</span><span class=\"kt\">int</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"o\">*</span><span class=\"w\">
  </span><span class=\"p\">(</span><span class=\"nx\">n</span><span class=\"p\">).(</span><span
  class=\"kt\">int</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">case</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">:</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">s</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">strconv</span><span class=\"p\">.</span><span
  class=\"nx\">Atoi</span><span class=\"p\">(</span><span class=\"nx\">n</span><span
  class=\"p\">.(</span><span class=\"kt\">string</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">return</span><span
  class=\"w\"> </span><span class=\"nx\">s</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">case</span><span class=\"w\"> </span><span
  class=\"kt\">float64</span><span class=\"p\">:</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"nb\">int</span><span class=\"p\">(</span><span class=\"nx\">n</span><span
  class=\"p\">.(</span><span class=\"kt\">float64</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">default</span><span
  class=\"p\">:</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">n</span><span
  class=\"p\">.(</span><span class=\"kt\">int</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">num</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">parse_int</span><span
  class=\"p\">(</span><span class=\"mi\">4</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">num</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">num</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">parse_int</span><span class=\"p\">(</span><span
  class=\"s\">&quot;4&quot;</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">num</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">num</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">parse_int</span><span class=\"p\">(</span><span
  class=\"mf\">4.1243</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">num</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run main.go\n\n16\n4\n4\n</pre></div>\n\n</pre>\n\n<p>Here,
  we can see we have an interface as a parameter to the function <code>parse_int</code>,
  the return type is <code>int</code>, so the incoming parameter can be any valid
  type. But if we don't convert the given type into an appropriate int, it will result
  in an error as we are returning the int value of the parsed parameter. We are taking
  the parameter as <code>interface{}</code> which is an empty interface, this will
  contain the parameter parsed as an interface type. That's why we need to convert
  the interface object into an int or the parsed type of the interface.</p>\n<h3>Interface
  Slice</h3>\n<p>We can even create a slice of interfaces, which means we can initialize
  or group together various objects of different structs in a single slice. This might
  be helpful for calling functions associated with different objects via the interface
  very easily and in a much cleaner way.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Creature</span><span class=\"w\"> </span><span
  class=\"kd\">interface</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">intro</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">type</span><span class=\"w\"> </span><span class=\"nx\">Player</span><span
  class=\"w\"> </span><span class=\"kd\">struct</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">name</span><span class=\"w\">   </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">health</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">Mob</span><span class=\"w\"> </span><span
  class=\"kd\">struct</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">name</span><span
  class=\"w\">     </span><span class=\"kt\">string</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">health</span><span class=\"w\">   </span><span
  class=\"kt\">int</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">category</span><span class=\"w\"> </span><span class=\"kt\">bool</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"p\">(</span><span
  class=\"nx\">p</span><span class=\"w\"> </span><span class=\"nx\">Player</span><span
  class=\"p\">)</span><span class=\"w\"> </span><span class=\"nx\">intro</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Player has spawned&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">return</span><span class=\"w\"> </span><span class=\"nx\">p</span><span
  class=\"p\">.</span><span class=\"nx\">name</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"p\">(</span><span class=\"nx\">m</span><span class=\"w\">
  </span><span class=\"nx\">Mob</span><span class=\"p\">)</span><span class=\"w\">
  </span><span class=\"nx\">intro</span><span class=\"p\">()</span><span class=\"w\">
  </span><span class=\"kt\">string</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">name</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">m</span><span
  class=\"p\">.</span><span class=\"nx\">name</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"s\">&quot;&quot;</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">name</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">m</span><span class=\"p\">.</span><span class=\"nx\">name</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"> </span><span class=\"k\">else</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">name</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Mob&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Printf</span><span class=\"p\">(</span><span class=\"s\">&quot;A wild
  %s has appeared!\\n&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">name</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"nx\">m</span><span class=\"p\">.</span><span class=\"nx\">name</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">entity</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"nx\">Creature</span><span class=\"p\">{</span><span class=\"nx\">Player</span><span
  class=\"p\">{},</span><span class=\"w\"> </span><span class=\"nx\">Mob</span><span
  class=\"p\">{},</span><span class=\"w\"> </span><span class=\"nx\">Mob</span><span
  class=\"p\">{},</span><span class=\"w\"> </span><span class=\"nx\">Player</span><span
  class=\"p\">{}}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"k\">for</span><span class=\"w\"> </span><span class=\"nx\">_</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">obj</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"k\">range</span><span
  class=\"w\"> </span><span class=\"nx\">entity</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">obj</span><span class=\"p\">.</span><span
  class=\"nx\">intro</span><span class=\"p\">())</span><span class=\"w\"></span>\n<span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>Player has spawned\nA wild
  Zombie has appeared!\nA wild Zombie has appeared!\nPlayer has spawned\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we can see that the entity variable is created as a slice of
  interfaces <code>Creature</code>, i.e. various objects associated with the <code>Creature</code>
  interface can be contained in a single slice. There are 2 instances of Player and
  Mob each in the slice. We can further iterate over the slice as a range-based loop
  and thereby the functions associated with the interfaces can be called. Here, we
  have called the <code>intro</code> function.</p>\n<p>So, there are a lot of things
  that can be done with interfaces, we can create multiple interfaces for a single
  object struct and nest interfaces. Based on the use case of the program, interfaces
  can be used to reduce the boilerplate code as well as improve the readability of
  the code.</p>\n<p>That's it from this part. Reference for all the code examples
  and commands can be found in the <a href=\"https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/interfaces/main.go\">100
  days of Golang</a> GitHub repository.</p>\n<h2>Conclusion</h2>\n<p>From this part
  of the series, we were able to understand the basics of interfaces using a few examples.
  We explored how interfaces can be used to bring in polymorphism in golang, also
  we can improve the readability of the code. The boilerplate code can be considerably
  reduced by using interfaces when dealing with structs and types. Hopefully, you
  found this post helpful and understood even the basics of interfaces in golang.
  Thank you for reading, if you have any queries, questions, or feedback, you can
  ping me on my social handles or in the comments. Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/golang-019-interfaces.png
long_description: 'In the 19th post of the series, we will be taking a look into interfaces
  in golang. Interfaces allow us to create function signatures common to different
  structs or types. So, we can allow multiple structs to have a common interface(method)
  that can '
now: 2025-05-01 18:11:33.314918
path: blog/posts/2022-09-10-GO-19-Interfaces.md
prevnext: null
series:
- 100-days-of-golang
slug: golang-interfaces
status: published
tags:
- go
templateKey: blog-post
title: 'Golang: Interfaces'
today: 2025-05-01
---

## Introduction

In the 19th post of the series, we will be taking a look into interfaces in golang. Interfaces allow us to create function signatures common to different structs or types. So, we can allow multiple structs to have a common interface(method) that can have different implementations.

## What are Interfaces

Interface as the name suggests is a way to create methods that are common to different structures or types but can have different implementations. It's an interface to define the method or function signatures but not the implementation. Let's take an example of `Laptop` and `Phone` having the functionality of wifi. We can connect to wifi more or the less in a similar way on both devices. The implementation behind the functionality might be different but they share the same operation. The WiFi can act as an interface for both devices to connect to the internet.

## Define an Interface

To declare an interface in golang, we can use the `interface` keyword in golang. An interface is created with the type keyword, providing the name of the interface and defining the function declaration. Inside the interface which acts as a struct of general method signatures. The method signatures usually consist of the name of the function with its parameters if any and the return type of the function.

```go

package main

import "fmt"


type Player struct {
	name   string
	health int
}

type Mob struct {
	name     string
	health   int
	is_passive bool
}

type Creature interface {
	intro() string
	//attack() int
	//heal() int
}

func main() {
	player := Player{name: "Steve"}
	mob := Mob{name: "Zombie"}
	fmt.Println(player)
	fmt.Println(mob)
}
```

```
go run main.go

{Steve 0}
{Zombie 0 false}
```

In this above example, we have created an interface called `Creature`. There are a few structs that we have to define like `Player` and `Mob`, these two methods have a few attributes like `name` as `string` and `health` as `int` which are common in both structs but the `Mob` struct has an additional attribute `is_passive` as a `boolean` value. The `Creature` is an interface that will define certain function signatures, here we have declared `intro`, `attack`, and `heal` as the methods bound to the Creature interface. This means, that any object which satisfies the Creature interface, can define the methods associated with the interface.

## Defining Interfaces

Once we have declared the interface method signatures, we can move into defining the functionality of these methods depending on the struct. If we want a different working method for different types of struct objects passed we can define those for each type of struct that we want. Here, we have two types of structs namely `Creature` and `Mob`, based on the struct we can define the `intro` method for them individually.

```go
package main

import "fmt"

type Creature interface {
	intro() string
	attack(*int) int
}

type Player struct {
	name   string
	health int
}

type Mob struct {
	name     string
	health   int
	category bool
}

func (p Player) intro() string {
	fmt.Println("Player has spawned")
	return p.name
}

func (p Player) attack(m_health *int) int {
	fmt.Println("Player has attacked!")
	*m_health = *m_health - 50
	return *m_health
}

func (m Mob) intro() string {
	fmt.Printf("A wild %s has appeared!\n", m.name)
	return m.name
}
func (m Mob) attack(p_health *int) int {
	fmt.Printf("%s has attacked you! -%d\n", m.name, 30)
	*p_health = *p_health - 30
	return *p_health
}

func main() {
	player := Player{name: "Steve", health: 100}
	mob := Mob{name: "Zombie", health: 140}
	fmt.Println(player.intro())
	fmt.Println(mob.intro())
	fmt.Println(mob)
	fmt.Println(player)
	fmt.Println(player.attack(&mob.health))
	fmt.Println(mob.attack(&player.health))
	fmt.Println(mob)
	fmt.Println(player)
}
```

```
go run main.go

Player has spawned
Steve
A wild Zombie has appeared!
Zombie
{Zombie 140 false}
{Steve 100}
Player has attacked!
90
Zombie has attacked you! -30
70
{Zombie 90 false}
{Steve 70}
```

As we can see, the method `intro()` is bound to both the struct depending on what struct signature is associated with the method. The method `intro` takes in the object struct associated as per the call and returns `string` as defined in the method signature. 

The `attack` method in the `Creature` interface is also implemented separately for the two structs. For the `Player` method, we simply take in a pointer to an integer and return an `int`. The parameter is the pointer to the mob health, and it returns the modified health. We take in a pointer to the mob or player health so as to parse in the actual value and not the copy of the value. If we modify the value, we want to reflect those changes in the actual object. So that is how we can use interfaces to construct dynamic operations on objects as well as different types of structs.

We have seen a simple example of how to declare and define interfaces for given type structs. Also, we can pass by value as well as by pointers so as to define the behavior of the method whether to dynamically modify or change the values of the object associated with it.

## Examples of Interfaces

There are quite some use cases of interfaces, in object-oriented programming, the above example fits the polymorphism feature quite well. The ability to reuse certain method signatures and define the functions as per requirement brings flexibility to the code structure.

We will see a few examples for understanding how we can use interfaces in various ways.

### Type Switch Interface

We can use an empty interface to check for the type of variable we have parsed. Using this empty interface we can create a kind of dynamic parameter to a function.  

```go
package main

import (
	"fmt"
	"strconv"
)

func parse_int(n interface{}) int {
	switch n.(type) {
	case int:
		return (n).(int) * (n).(int)
	case string:
		s, _ := strconv.Atoi(n.(string))
		return s
	case float64:
		return int(n.(float64))
	default:
		return n.(int)
	}
}

func main() {
	num := parse_int(4)
	fmt.Println(num)
	num = parse_int("4")
	fmt.Println(num)
	num = parse_int(4.1243)
	fmt.Println(num)
}

```

```
go run main.go

16
4
4
```

Here, we can see we have an interface as a parameter to the function `parse_int`, the return type is `int`, so the incoming parameter can be any valid type. But if we don't convert the given type into an appropriate int, it will result in an error as we are returning the int value of the parsed parameter. We are taking the parameter as `interface{}` which is an empty interface, this will contain the parameter parsed as an interface type. That's why we need to convert the interface object into an int or the parsed type of the interface.

### Interface Slice

We can even create a slice of interfaces, which means we can initialize or group together various objects of different structs in a single slice. This might be helpful for calling functions associated with different objects via the interface very easily and in a much cleaner way.

```go
package main

import "fmt"

type Creature interface {
	intro() string
}

type Player struct {
	name   string
	health int
}

type Mob struct {
	name     string
	health   int
	category bool
}

func (p Player) intro() string {
	fmt.Println("Player has spawned")
	return p.name
}

func (m Mob) intro() string {
	var name string
	if m.name != "" {
		name = m.name
	} else {
		name = "Mob"
	}
	fmt.Printf("A wild %s has appeared!\n", name)
	return m.name
}


func main() {
	entity := []Creature{Player{}, Mob{}, Mob{}, Player{}}

	for _, obj := range entity {
		fmt.Println(obj.intro())
	}
}
```

```
Player has spawned
A wild Zombie has appeared!
A wild Zombie has appeared!
Player has spawned
```

In the above example, we can see that the entity variable is created as a slice of interfaces `Creature`, i.e. various objects associated with the `Creature` interface can be contained in a single slice. There are 2 instances of Player and Mob each in the slice. We can further iterate over the slice as a range-based loop and thereby the functions associated with the interfaces can be called. Here, we have called the `intro` function.

So, there are a lot of things that can be done with interfaces, we can create multiple interfaces for a single object struct and nest interfaces. Based on the use case of the program, interfaces can be used to reduce the boilerplate code as well as improve the readability of the code.

That's it from this part. Reference for all the code examples and commands can be found in the [100 days of Golang](https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/interfaces/main.go) GitHub repository.

## Conclusion

From this part of the series, we were able to understand the basics of interfaces using a few examples. We explored how interfaces can be used to bring in polymorphism in golang, also we can improve the readability of the code. The boilerplate code can be considerably reduced by using interfaces when dealing with structs and types. Hopefully, you found this post helpful and understood even the basics of interfaces in golang. Thank you for reading, if you have any queries, questions, or feedback, you can ping me on my social handles or in the comments. Happy Coding :)