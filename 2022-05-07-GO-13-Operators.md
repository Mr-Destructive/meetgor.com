---
article_html: "<h2>Introduction</h2>\n<p>In this 13th part of the series, we will
  be exploring the fundamentals of operators in Golang. We will be exploring the basics
  of operators and the various types like Arithmetic, Bitwise, Comparison, Assignment
  operators in Golang.</p>\n<p>Operators are quite fundamentals in any programming
  language. Operators are basically expressions or a set of character(s) to perform
  certain fundamental tasks. They allow us to perform certain trivial operations with
  a simple expression or character. There are quite a few operators in Golang to perform
  various operations.</p>\n<h2>Types of Operators</h2>\n<p>Golang has a few types
  of operators, each type providing particular aspect of forming expressions and evaluate
  conditions.</p>\n<ol>\n<li>Bitwise Operators</li>\n<li>Logical Operators</li>\n<li>Arithmetic
  Operators</li>\n<li>Assignment Operators</li>\n<li>Comparison Operators</li>\n</ol>\n<h3>Bitwise
  Operators</h3>\n<p>Bitwise Operators are used in performing operations on binary
  numbers. We can perform operation on a bit level and hence they are known as bitwise
  operators. Some fundamental bitwise operators include, <code>AND</code>, <code>OR</code>,
  <code>NOT</code>, and <code>EXOR</code>. Using this operators, the bits in the operands
  can be manipulated and certain logical operations can be performed.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">x</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">3</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">y</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">5</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">// 3 -&gt; 011</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">// 5 -&gt; 101</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;X AND Y = &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">x</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"w\"> </span><span class=\"nx\">y</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;X OR
  Y = &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">x</span><span
  class=\"w\"> </span><span class=\"p\">|</span><span class=\"w\"> </span><span class=\"nx\">y</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;X EXOR Y = &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">x</span><span class=\"w\"> </span><span class=\"p\">^</span><span
  class=\"w\"> </span><span class=\"nx\">y</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;X Right
  Shift 1  = &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">x</span><span class=\"w\"> </span><span class=\"o\">&gt;&gt;</span><span
  class=\"w\"> </span><span class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;X Right
  Shift 2  = &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">x</span><span class=\"w\"> </span><span class=\"o\">&gt;&gt;</span><span
  class=\"w\"> </span><span class=\"mi\">2</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Y Left
  Shift 1 = &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">y</span><span class=\"w\"> </span><span class=\"o\">&lt;&lt;</span><span
  class=\"w\"> </span><span class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run bitwise/main.go\n\nX
  AND Y =  1\nX OR Y =  7\nX EXOR Y =  6\nX Right Shift 1  =  1\nX Right Shift 2  =
  \ 0\nY Left Shift 1 =  10\n</pre></div>\n\n</pre>\n\n<p>We use the <code>&amp;</code>
  (AND operator) for performing AND operations on two operands. Here we are logically
  ANDing <code>3</code> and <code>5</code> i.e. <code>011</code> with <code>101</code>
  so it becomes <code>001</code> in binary or 1 in decimal.</p>\n<p>Also, the <code>|</code>
  (OR operator) for performing logical OR operation on two operands. Here we are logically
  ORing <code>3</code> and <code>5</code> i.e. <code>011</code> with <code>101</code>
  so it becomes <code>111</code> in binary or 7 in decimal.</p>\n<p>Also the <code>^</code>
  (EXOR operator) for performing logical EXOR operation on two operands. Here we are
  logically EXORing <code>3</code> and <code>5</code> i.e. <code>011</code> with <code>101</code>
  so it becomes <code>110</code> in binary or 6 in decimal.</p>\n<p>We have a couple
  of more bitwise operators that allow us to shift bits in the binary representation
  of the number. We have two types of these shift operators, right sift and left shift
  operators. The main function of these operator is to shift a bit in either right
  or left direction.</p>\n<p>In the above example, we have shifted <code>3</code>
  i.e. <code>011</code> to right by one bit so it becomes <code>001</code>. If we
  would have given <code>x &gt;&gt; 2</code> it would have become <code>0</code> since
  the last bit was shifted to right and hence all bits were 0.</p>\n<p>Similarly,
  the left shift operator sifts the bits in the binary representation of the number
  to the left. So, in the example above, <code>5</code> i.e. <code>101</code> is shifted
  left by one bit so it becomes <code>1010</code> in binary i.e. 10 in decimal.</p>\n<p>This
  was a basic overview of bitwise operators in Golang. We can use these basic operators
  to perform low level operations on numbers.</p>\n<h3>Comparison Operators</h3>\n<p>This
  type of operators are quite important and widely used as they form the fundamentals
  of comparison of variables and forming boolean expressions. The comparison operator
  is used to compare two values or expressions.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">45</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">b</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">12</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Is A equal to B ? &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">==</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Is
  A not equal to B ? &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Is
  A greater than B ? &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"p\">&gt;</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Is
  A less than B ? &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Is
  A greater than or equal to B ? &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">&gt;=</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Is
  A less than or equal to B ? &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">&lt;=</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run comparison/main.go\n\nIs
  A equal to B ?  false\nIs A not equal to B ?  true\nIs A greater than B ?  true\nIs
  A less than B ?  false\nIs A greater than or equal to B ?  true\nIs A less than
  or equal to B ?  false\n</pre></div>\n\n</pre>\n\n<p>We use simple comparison operators
  like <code>==</code> or <code>!=</code> for comparing if two values are equal or
  not. The expression <code>a == b</code> will evaluate to <code>true</code> if the
  values of both variables or operands are equal. However, the expression <code>a
  != b</code> will evaluate to <code>true</code> if the values of both variables or
  operands are not equal.</p>\n<p>Similarly, we have the <code>&lt;</code> and <code>&gt;</code>
  operators which allow us to evaluate expression by comparing if the values are less
  than or grater than the other operand. So, the expression <code>a &gt; b</code>
  will evaluate to <code>true</code> if the value of <code>a</code> is greater than
  the value of <code>b</code>. Also the expression <code>a &lt; b</code> will evaluate
  to <code>true</code> if the value of <code>a</code> is less than the value of <code>b</code>.</p>\n<p>Finally,
  the operators <code>&lt;=</code> and <code>&gt;=</code> allow us to evaluate expression
  by comparing if the values are less than or equal to and greater than or equal to
  the other operand. So, the expression <code>a &gt;= b</code> will evaluate to <code>true</code>
  if the value of <code>a</code> is greater than or if it is equal to the value of
  <code>b</code>, else it would evaluate to <code>false</code>. Similarly, the expression
  <code>a &lt;= b</code> will evaluate to <code>true</code> if the value of <code>a</code>
  is less than or if it is equal to the value of <code>b</code>, else it would evaluate
  to <code>false</code>.</p>\n<p>These was a basic overview of comparison operators
  in golang.</p>\n<h3>Logical Operators</h3>\n<p>Next, we move on to the logical operators
  in Golang which allow to perform logical operations like <code>AND</code>, <code>OR</code>,
  and <code>NOT</code> with conditional statements or storing boolean expressions.</p>\n<pre
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">45</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">b</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;Something&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"p\">&gt;</span><span
  class=\"w\"> </span><span class=\"mi\">40</span><span class=\"w\"> </span><span
  class=\"o\">&amp;&amp;</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"w\"> </span><span class=\"o\">==</span><span class=\"w\"> </span><span class=\"s\">&quot;Something&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">a</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"mi\">40</span><span class=\"w\"> </span><span
  class=\"o\">&amp;&amp;</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"w\"> </span><span class=\"o\">==</span><span class=\"w\"> </span><span class=\"s\">&quot;Something&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">a</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"mi\">40</span><span class=\"w\"> </span><span
  class=\"o\">||</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"w\"> </span><span class=\"o\">==</span><span class=\"w\"> </span><span class=\"s\">&quot;Something&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">a</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"mi\">40</span><span class=\"w\"> </span><span
  class=\"o\">||</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"s\">&quot;Something&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(!(</span><span class=\"nx\">a</span><span class=\"w\"> </span><span
  class=\"p\">&lt;</span><span class=\"w\"> </span><span class=\"mi\">40</span><span
  class=\"w\"> </span><span class=\"o\">||</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"s\">&quot;Something&quot;</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run logical/main.go\n\ntrue\nfalse\ntrue\nfalse\ntrue\n</pre></div>\n\n</pre>\n\n<p>Here,
  we have used logical operators like <code>&amp;&amp;</code> for Logical AND, <code>||</code>
  for logical OR, and <code>!</code> for complementing the evaluated result. The <code>&amp;&amp;</code>
  operation only evaluates to <code>true</code> if both the expressions are <code>true</code>
  and <code>||</code> OR operator evaluates to <code>true</code> if either or both
  the expressions are <code>true</code>. The <code>!</code> operator is used to complement
  the evaluated expression from the preceding parenthesis.</p>\n<h3>Arithmetic Operators</h3>\n<p>Arithmetic
  operators are used for performing Arithmetic operations. We have few basic arithmetic
  operators like <code>+</code>, <code>-</code>, <code>*</code>, <code>/</code>, and
  <code>%</code> for adding, subtracting, multiplication, division, and modulus operation
  in golang.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">30</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">b</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">50</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;A + B = &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">a</span><span class=\"o\">+</span><span class=\"nx\">b</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;A - B = &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">a</span><span class=\"o\">-</span><span class=\"nx\">b</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;A * B = &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">a</span><span class=\"o\">*</span><span class=\"nx\">b</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;A / B = &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">a</span><span class=\"o\">/</span><span class=\"nx\">b</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;A % B = &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">a</span><span class=\"o\">%</span><span class=\"nx\">b</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run arithmetic/main.go\nA
  + B =  80\nA - B =  -20\nA * B =  1500\nA / B =  0\nA % B =  30\n</pre></div>\n\n</pre>\n\n<p>These
  are the basic mathematical operators in any programming language. We can use <code>+</code>
  to add two values, <code>-</code> to subtract two values, <code>*</code> to multiply
  to values, <code>/</code> for division of two values and finally <code>%</code>
  to get the remainder of a division of two values i.e. if we divide 30 by 50, the
  remainder is 30 and the quotient is 0.</p>\n<p>We also have a few other operators
  like <code>++</code> and <code>--</code> that help in incrementing and decrementing
  values by a unit value. Let's say we have a variable <code>k</code> which is set
  to <code>4</code> and we want to increment it by one, so we can definitely use <code>k
  = k + 1</code> but it looks kind of too long, we have a short notation for the same
  <code>k++</code> to do the same.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">k</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">3</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">j</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">20</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;k = &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">k</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;j =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">j</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">k</span><span class=\"o\">++</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">j</span><span class=\"o\">--</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;k = &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">k</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;j =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">j</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run arithmetic/main.go\n\nk
  =  3\nj =  20\n\nk =  4\nj =  19\n</pre></div>\n\n</pre>\n\n<p>So, we can see that
  the variable <code>k</code> is incremented by one and variable <code>j</code> is
  decremented by <code>1</code> using the <code>++</code> and <code>--</code> operator.</p>\n<h3>Assignment
  Operators</h3>\n<p>These types of operators are quite handy and can condense down
  large operations into simple expressions. These types of operators allow us to perform
  operation on the same operand. Let's say we have the variable <code>k</code> set
  to <code>20</code> initially, we want to add <code>30</code> to the variable <code>k</code>,
  we can do that by using <code>k = k + 30</code> but a more sophisticated way would
  be to use <code>k += 30</code> which adds <code>30</code> or any value provided
  the same variable assigned and operated on.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">a</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"mi\">100</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">b</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">20</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;a = &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">a</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;b =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">+=</span><span
  class=\"w\"> </span><span class=\"mi\">30</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;a =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">a</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">b</span><span class=\"w\"> </span><span class=\"o\">-=</span><span
  class=\"w\"> </span><span class=\"mi\">5</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;b =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">*=</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;a =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">a</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;b = &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">a</span><span class=\"w\"> </span><span
  class=\"o\">/=</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;a = &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">a</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;b =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">%=</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;a =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">a</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;b = &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run assignment/main.go\n\na
  =  100\nb =  20\n\na =  130\nb =  15\n\na =  1950\nb =  15\n\na =  130\nb =  15\n\na
  =  10\nb =  15\n</pre></div>\n\n</pre>\n\n<p>From the above example, we are able
  to perform operations by using shorthand notations like <code>+=</code> to add the
  value to the same operand. These also saves a bit of time and memory not much but
  considerable enough. This allow us to directly access and modify the contents of
  the provided operand in the register rather than assigning different registers and
  performing the operations.</p>\n<p>That's it from this part. Reference for all the
  code examples and commands can be found in the <a href=\"https://github.com/mr-destructive/100-days-of-golang/\">100
  days of Golang</a> GitHub repository.</p>\n<h2>Conclusion</h2>\n<p>So, from the
  following part of the series, we were able to learn the basics of operators in golang.
  Using some simple and easy to understand examples, we were able to explore different
  types of operators like arithmetic, logical, assignment and bitwise operators in
  golang. These are quite fundamental in programming in general, this lays a good
  foundation for working with larger and complex projects that deal with any kind
  of logic in it, without a doubt almost all of the applications do have a bit of
  logic attached to it. So, we need to know the basics of operators in golang.</p>\n"
cover: ''
date: 2022-05-07
datetime: 2022-05-07 00:00:00+00:00
description: Understanding the basics of operators like arithmetic, logical, bitwise,
  assignment operators in Golang
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2022-05-07-GO-13-Operators.md
html: "<h2>Introduction</h2>\n<p>In this 13th part of the series, we will be exploring
  the fundamentals of operators in Golang. We will be exploring the basics of operators
  and the various types like Arithmetic, Bitwise, Comparison, Assignment operators
  in Golang.</p>\n<p>Operators are quite fundamentals in any programming language.
  Operators are basically expressions or a set of character(s) to perform certain
  fundamental tasks. They allow us to perform certain trivial operations with a simple
  expression or character. There are quite a few operators in Golang to perform various
  operations.</p>\n<h2>Types of Operators</h2>\n<p>Golang has a few types of operators,
  each type providing particular aspect of forming expressions and evaluate conditions.</p>\n<ol>\n<li>Bitwise
  Operators</li>\n<li>Logical Operators</li>\n<li>Arithmetic Operators</li>\n<li>Assignment
  Operators</li>\n<li>Comparison Operators</li>\n</ol>\n<h3>Bitwise Operators</h3>\n<p>Bitwise
  Operators are used in performing operations on binary numbers. We can perform operation
  on a bit level and hence they are known as bitwise operators. Some fundamental bitwise
  operators include, <code>AND</code>, <code>OR</code>, <code>NOT</code>, and <code>EXOR</code>.
  Using this operators, the bits in the operands can be manipulated and certain logical
  operations can be performed.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">x</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">3</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">y</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">5</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">// 3 -&gt; 011</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">// 5 -&gt; 101</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;X AND Y = &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">x</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"w\"> </span><span class=\"nx\">y</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;X OR
  Y = &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">x</span><span
  class=\"w\"> </span><span class=\"p\">|</span><span class=\"w\"> </span><span class=\"nx\">y</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;X EXOR Y = &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">x</span><span class=\"w\"> </span><span class=\"p\">^</span><span
  class=\"w\"> </span><span class=\"nx\">y</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;X Right
  Shift 1  = &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">x</span><span class=\"w\"> </span><span class=\"o\">&gt;&gt;</span><span
  class=\"w\"> </span><span class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;X Right
  Shift 2  = &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">x</span><span class=\"w\"> </span><span class=\"o\">&gt;&gt;</span><span
  class=\"w\"> </span><span class=\"mi\">2</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Y Left
  Shift 1 = &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">y</span><span class=\"w\"> </span><span class=\"o\">&lt;&lt;</span><span
  class=\"w\"> </span><span class=\"mi\">1</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run bitwise/main.go\n\nX
  AND Y =  1\nX OR Y =  7\nX EXOR Y =  6\nX Right Shift 1  =  1\nX Right Shift 2  =
  \ 0\nY Left Shift 1 =  10\n</pre></div>\n\n</pre>\n\n<p>We use the <code>&amp;</code>
  (AND operator) for performing AND operations on two operands. Here we are logically
  ANDing <code>3</code> and <code>5</code> i.e. <code>011</code> with <code>101</code>
  so it becomes <code>001</code> in binary or 1 in decimal.</p>\n<p>Also, the <code>|</code>
  (OR operator) for performing logical OR operation on two operands. Here we are logically
  ORing <code>3</code> and <code>5</code> i.e. <code>011</code> with <code>101</code>
  so it becomes <code>111</code> in binary or 7 in decimal.</p>\n<p>Also the <code>^</code>
  (EXOR operator) for performing logical EXOR operation on two operands. Here we are
  logically EXORing <code>3</code> and <code>5</code> i.e. <code>011</code> with <code>101</code>
  so it becomes <code>110</code> in binary or 6 in decimal.</p>\n<p>We have a couple
  of more bitwise operators that allow us to shift bits in the binary representation
  of the number. We have two types of these shift operators, right sift and left shift
  operators. The main function of these operator is to shift a bit in either right
  or left direction.</p>\n<p>In the above example, we have shifted <code>3</code>
  i.e. <code>011</code> to right by one bit so it becomes <code>001</code>. If we
  would have given <code>x &gt;&gt; 2</code> it would have become <code>0</code> since
  the last bit was shifted to right and hence all bits were 0.</p>\n<p>Similarly,
  the left shift operator sifts the bits in the binary representation of the number
  to the left. So, in the example above, <code>5</code> i.e. <code>101</code> is shifted
  left by one bit so it becomes <code>1010</code> in binary i.e. 10 in decimal.</p>\n<p>This
  was a basic overview of bitwise operators in Golang. We can use these basic operators
  to perform low level operations on numbers.</p>\n<h3>Comparison Operators</h3>\n<p>This
  type of operators are quite important and widely used as they form the fundamentals
  of comparison of variables and forming boolean expressions. The comparison operator
  is used to compare two values or expressions.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">45</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">b</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">12</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Is A equal to B ? &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">==</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Is
  A not equal to B ? &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Is
  A greater than B ? &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"p\">&gt;</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Is
  A less than B ? &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Is
  A greater than or equal to B ? &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">&gt;=</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Is
  A less than or equal to B ? &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">&lt;=</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run comparison/main.go\n\nIs
  A equal to B ?  false\nIs A not equal to B ?  true\nIs A greater than B ?  true\nIs
  A less than B ?  false\nIs A greater than or equal to B ?  true\nIs A less than
  or equal to B ?  false\n</pre></div>\n\n</pre>\n\n<p>We use simple comparison operators
  like <code>==</code> or <code>!=</code> for comparing if two values are equal or
  not. The expression <code>a == b</code> will evaluate to <code>true</code> if the
  values of both variables or operands are equal. However, the expression <code>a
  != b</code> will evaluate to <code>true</code> if the values of both variables or
  operands are not equal.</p>\n<p>Similarly, we have the <code>&lt;</code> and <code>&gt;</code>
  operators which allow us to evaluate expression by comparing if the values are less
  than or grater than the other operand. So, the expression <code>a &gt; b</code>
  will evaluate to <code>true</code> if the value of <code>a</code> is greater than
  the value of <code>b</code>. Also the expression <code>a &lt; b</code> will evaluate
  to <code>true</code> if the value of <code>a</code> is less than the value of <code>b</code>.</p>\n<p>Finally,
  the operators <code>&lt;=</code> and <code>&gt;=</code> allow us to evaluate expression
  by comparing if the values are less than or equal to and greater than or equal to
  the other operand. So, the expression <code>a &gt;= b</code> will evaluate to <code>true</code>
  if the value of <code>a</code> is greater than or if it is equal to the value of
  <code>b</code>, else it would evaluate to <code>false</code>. Similarly, the expression
  <code>a &lt;= b</code> will evaluate to <code>true</code> if the value of <code>a</code>
  is less than or if it is equal to the value of <code>b</code>, else it would evaluate
  to <code>false</code>.</p>\n<p>These was a basic overview of comparison operators
  in golang.</p>\n<h3>Logical Operators</h3>\n<p>Next, we move on to the logical operators
  in Golang which allow to perform logical operations like <code>AND</code>, <code>OR</code>,
  and <code>NOT</code> with conditional statements or storing boolean expressions.</p>\n<pre
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">45</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">b</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;Something&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"p\">&gt;</span><span
  class=\"w\"> </span><span class=\"mi\">40</span><span class=\"w\"> </span><span
  class=\"o\">&amp;&amp;</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"w\"> </span><span class=\"o\">==</span><span class=\"w\"> </span><span class=\"s\">&quot;Something&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">a</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"mi\">40</span><span class=\"w\"> </span><span
  class=\"o\">&amp;&amp;</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"w\"> </span><span class=\"o\">==</span><span class=\"w\"> </span><span class=\"s\">&quot;Something&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">a</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"mi\">40</span><span class=\"w\"> </span><span
  class=\"o\">||</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"w\"> </span><span class=\"o\">==</span><span class=\"w\"> </span><span class=\"s\">&quot;Something&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">a</span><span class=\"w\"> </span><span class=\"p\">&lt;</span><span
  class=\"w\"> </span><span class=\"mi\">40</span><span class=\"w\"> </span><span
  class=\"o\">||</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"s\">&quot;Something&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(!(</span><span class=\"nx\">a</span><span class=\"w\"> </span><span
  class=\"p\">&lt;</span><span class=\"w\"> </span><span class=\"mi\">40</span><span
  class=\"w\"> </span><span class=\"o\">||</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"s\">&quot;Something&quot;</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run logical/main.go\n\ntrue\nfalse\ntrue\nfalse\ntrue\n</pre></div>\n\n</pre>\n\n<p>Here,
  we have used logical operators like <code>&amp;&amp;</code> for Logical AND, <code>||</code>
  for logical OR, and <code>!</code> for complementing the evaluated result. The <code>&amp;&amp;</code>
  operation only evaluates to <code>true</code> if both the expressions are <code>true</code>
  and <code>||</code> OR operator evaluates to <code>true</code> if either or both
  the expressions are <code>true</code>. The <code>!</code> operator is used to complement
  the evaluated expression from the preceding parenthesis.</p>\n<h3>Arithmetic Operators</h3>\n<p>Arithmetic
  operators are used for performing Arithmetic operations. We have few basic arithmetic
  operators like <code>+</code>, <code>-</code>, <code>*</code>, <code>/</code>, and
  <code>%</code> for adding, subtracting, multiplication, division, and modulus operation
  in golang.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">30</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">b</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">50</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;A + B = &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">a</span><span class=\"o\">+</span><span class=\"nx\">b</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;A - B = &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">a</span><span class=\"o\">-</span><span class=\"nx\">b</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;A * B = &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">a</span><span class=\"o\">*</span><span class=\"nx\">b</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;A / B = &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">a</span><span class=\"o\">/</span><span class=\"nx\">b</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;A % B = &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">a</span><span class=\"o\">%</span><span class=\"nx\">b</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run arithmetic/main.go\nA
  + B =  80\nA - B =  -20\nA * B =  1500\nA / B =  0\nA % B =  30\n</pre></div>\n\n</pre>\n\n<p>These
  are the basic mathematical operators in any programming language. We can use <code>+</code>
  to add two values, <code>-</code> to subtract two values, <code>*</code> to multiply
  to values, <code>/</code> for division of two values and finally <code>%</code>
  to get the remainder of a division of two values i.e. if we divide 30 by 50, the
  remainder is 30 and the quotient is 0.</p>\n<p>We also have a few other operators
  like <code>++</code> and <code>--</code> that help in incrementing and decrementing
  values by a unit value. Let's say we have a variable <code>k</code> which is set
  to <code>4</code> and we want to increment it by one, so we can definitely use <code>k
  = k + 1</code> but it looks kind of too long, we have a short notation for the same
  <code>k++</code> to do the same.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">k</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">3</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">j</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">20</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;k = &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">k</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;j =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">j</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">k</span><span class=\"o\">++</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">j</span><span class=\"o\">--</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;k = &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">k</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;j =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">j</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run arithmetic/main.go\n\nk
  =  3\nj =  20\n\nk =  4\nj =  19\n</pre></div>\n\n</pre>\n\n<p>So, we can see that
  the variable <code>k</code> is incremented by one and variable <code>j</code> is
  decremented by <code>1</code> using the <code>++</code> and <code>--</code> operator.</p>\n<h3>Assignment
  Operators</h3>\n<p>These types of operators are quite handy and can condense down
  large operations into simple expressions. These types of operators allow us to perform
  operation on the same operand. Let's say we have the variable <code>k</code> set
  to <code>20</code> initially, we want to add <code>30</code> to the variable <code>k</code>,
  we can do that by using <code>k = k + 30</code> but a more sophisticated way would
  be to use <code>k += 30</code> which adds <code>30</code> or any value provided
  the same variable assigned and operated on.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">a</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"mi\">100</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">b</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">20</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;a = &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">a</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;b =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">+=</span><span
  class=\"w\"> </span><span class=\"mi\">30</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;a =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">a</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">b</span><span class=\"w\"> </span><span class=\"o\">-=</span><span
  class=\"w\"> </span><span class=\"mi\">5</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;b =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">*=</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;a =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">a</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;b = &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">a</span><span class=\"w\"> </span><span
  class=\"o\">/=</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;a = &quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">a</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;b =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">b</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">a</span><span class=\"w\"> </span><span class=\"o\">%=</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;a =
  &quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">a</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;b = &quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">b</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run assignment/main.go\n\na
  =  100\nb =  20\n\na =  130\nb =  15\n\na =  1950\nb =  15\n\na =  130\nb =  15\n\na
  =  10\nb =  15\n</pre></div>\n\n</pre>\n\n<p>From the above example, we are able
  to perform operations by using shorthand notations like <code>+=</code> to add the
  value to the same operand. These also saves a bit of time and memory not much but
  considerable enough. This allow us to directly access and modify the contents of
  the provided operand in the register rather than assigning different registers and
  performing the operations.</p>\n<p>That's it from this part. Reference for all the
  code examples and commands can be found in the <a href=\"https://github.com/mr-destructive/100-days-of-golang/\">100
  days of Golang</a> GitHub repository.</p>\n<h2>Conclusion</h2>\n<p>So, from the
  following part of the series, we were able to learn the basics of operators in golang.
  Using some simple and easy to understand examples, we were able to explore different
  types of operators like arithmetic, logical, assignment and bitwise operators in
  golang. These are quite fundamental in programming in general, this lays a good
  foundation for working with larger and complex projects that deal with any kind
  of logic in it, without a doubt almost all of the applications do have a bit of
  logic attached to it. So, we need to know the basics of operators in golang.</p>\n"
image_url: https://meetgor-cdn.pages.dev/golang-013-operators.png
long_description: In this 13th part of the series, we will be exploring the fundamentals
  of operators in Golang. We will be exploring the basics of operators and the various
  types like Arithmetic, Bitwise, Comparison, Assignment operators in Golang. Operators
  are quit
now: 2025-05-01 18:11:33.313445
path: blog/posts/2022-05-07-GO-13-Operators.md
prevnext: null
series:
- 100-days-of-golang
slug: golang-operators
status: published
tags:
- go
templateKey: blog-post
title: 'Golang: Operators'
today: 2025-05-01
---

## Introduction 

In this 13th part of the series, we will be exploring the fundamentals of operators in Golang. We will be exploring the basics of operators and the various types like Arithmetic, Bitwise, Comparison, Assignment operators in Golang.

Operators are quite fundamentals in any programming language. Operators are basically expressions or a set of character(s) to perform certain fundamental tasks. They allow us to perform certain trivial operations with a simple expression or character. There are quite a few operators in Golang to perform various operations.

## Types of Operators

Golang has a few types of operators, each type providing particular aspect of forming expressions and evaluate conditions.

1. Bitwise Operators
2. Logical Operators
3. Arithmetic Operators
4. Assignment Operators
5. Comparison Operators

### Bitwise Operators

Bitwise Operators are used in performing operations on binary numbers. We can perform operation on a bit level and hence they are known as bitwise operators. Some fundamental bitwise operators include, `AND`, `OR`, `NOT`, and `EXOR`. Using this operators, the bits in the operands can be manipulated and certain logical operations can be performed. 

```go
package main

import "fmt"

func main() {
	x := 3
	y := 5
	// 3 -> 011
	// 5 -> 101
	fmt.Println("X AND Y = ", x & y)
	fmt.Println("X OR Y = ", x | y)
	fmt.Println("X EXOR Y = ", x ^ y)
	fmt.Println("X Right Shift 1  = ", x >> 1)
	fmt.Println("X Right Shift 2  = ", x >> 2)
	fmt.Println("Y Left Shift 1 = ", y << 1)
}
```

```
$ go run bitwise/main.go

X AND Y =  1
X OR Y =  7
X EXOR Y =  6
X Right Shift 1  =  1
X Right Shift 2  =  0
Y Left Shift 1 =  10

```

We use the `&` (AND operator) for performing AND operations on two operands. Here we are logically ANDing `3` and `5` i.e. `011` with `101` so it becomes `001` in binary or 1 in decimal.

Also, the `|` (OR operator) for performing logical OR operation on two operands. Here we are logically ORing `3` and `5` i.e. `011` with `101` so it becomes `111` in binary or 7 in decimal.

Also the `^` (EXOR operator) for performing logical EXOR operation on two operands. Here we are logically EXORing `3` and `5` i.e. `011` with `101` so it becomes `110` in binary or 6 in decimal.

We have a couple of more bitwise operators that allow us to shift bits in the binary representation of the number. We have two types of these shift operators, right sift and left shift operators. The main function of these operator is to shift a bit in either right or left direction. 

In the above example, we have shifted `3` i.e. `011` to right by one bit so it becomes `001`. If we would have given `x >> 2` it would have become `0` since the last bit was shifted to right and hence all bits were 0.

Similarly, the left shift operator sifts the bits in the binary representation of the number to the left. So, in the example above, `5` i.e. `101` is shifted left by one bit so it becomes `1010` in binary i.e. 10 in decimal. 

This was a basic overview of bitwise operators in Golang. We can use these basic operators to perform low level operations on numbers.

### Comparison Operators

This type of operators are quite important and widely used as they form the fundamentals of comparison of variables and forming boolean expressions. The comparison operator is used to compare two values or expressions. 

```go
package main

import "fmt"

func main() {
	a := 45
	b := 12
	fmt.Println("Is A equal to B ? ", a == b)
	fmt.Println("Is A not equal to B ? ", a != b)
	fmt.Println("Is A greater than B ? ", a > b)
	fmt.Println("Is A less than B ? ", a < b)
	fmt.Println("Is A greater than or equal to B ? ", a >= b)
	fmt.Println("Is A less than or equal to B ? ", a <= b)

```

```
$ go run comparison/main.go

Is A equal to B ?  false
Is A not equal to B ?  true
Is A greater than B ?  true
Is A less than B ?  false
Is A greater than or equal to B ?  true
Is A less than or equal to B ?  false
```

We use simple comparison operators like `==` or `!=` for comparing if two values are equal or not. The expression `a == b` will evaluate to `true` if the values of both variables or operands are equal. However, the expression `a != b` will evaluate to `true` if the values of both variables or operands are not equal.

Similarly, we have the `<` and `>` operators which allow us to evaluate expression by comparing if the values are less than or grater than the other operand. So, the expression `a > b` will evaluate to `true` if the value of `a` is greater than the value of `b`. Also the expression `a < b` will evaluate to `true` if the value of `a` is less than the value of `b`. 

Finally, the operators `<=` and `>=` allow us to evaluate expression by comparing if the values are less than or equal to and greater than or equal to the other operand. So, the expression `a >= b` will evaluate to `true` if the value of `a` is greater than or if it is equal to the value of `b`, else it would evaluate to `false`. Similarly, the expression `a <= b` will evaluate to `true` if the value of `a` is less than or if it is equal to the value of `b`, else it would evaluate to `false`.

These was a basic overview of comparison operators in golang.

### Logical Operators

Next, we move on to the logical operators in Golang which allow to perform logical operations like `AND`, `OR`, and `NOT` with conditional statements or storing boolean expressions. 

```go
package main

import "fmt"

func main() {
	a := 45
	b := "Something"
	fmt.Println(a > 40 && b == "Something")
	fmt.Println(a < 40 && b == "Something")
	fmt.Println(a < 40 || b == "Something")
	fmt.Println(a < 40 || b != "Something")
	fmt.Println(!(a < 40 || b != "Something"))
}
```

```
$ go run logical/main.go

true
false
true
false
true
```

Here, we have used logical operators like `&&` for Logical AND, `||` for logical OR, and `!` for complementing the evaluated result. The `&&` operation only evaluates to `true` if both the expressions are `true` and `||` OR operator evaluates to `true` if either or both the expressions are `true`. The `!` operator is used to complement the evaluated expression from the preceding parenthesis.

### Arithmetic Operators

Arithmetic operators are used for performing Arithmetic operations. We have few basic arithmetic operators like `+`, `-`, `*`, `/`, and `%` for adding, subtracting, multiplication, division, and modulus operation in golang. 

```go
package main

import "fmt"

func main() {
	a := 30
	b := 50
	fmt.Println("A + B = ", a+b)
	fmt.Println("A - B = ", a-b)
	fmt.Println("A * B = ", a*b)
	fmt.Println("A / B = ", a/b)
	fmt.Println("A % B = ", a%b)
}
```

```
$ go run arithmetic/main.go
A + B =  80
A - B =  -20
A * B =  1500
A / B =  0
A % B =  30
```

These are the basic mathematical operators in any programming language. We can use `+` to add two values, `-` to subtract two values, `*` to multiply to values, `/` for division of two values and finally `%` to get the remainder of a division of two values i.e. if we divide 30 by 50, the remainder is 30 and the quotient is 0. 

We also have a few other operators like `++` and `--` that help in incrementing and decrementing values by a unit value. Let's say we have a variable `k` which is set to `4` and we want to increment it by one, so we can definitely use `k = k + 1` but it looks kind of too long, we have a short notation for the same `k++` to do the same.

```go
package main

import "fmt"

func main() {
	k := 3
	j := 20
	fmt.Println("k = ", k)
	fmt.Println("j = ", j)
	k++
	j--
	fmt.Println("k = ", k)
	fmt.Println("j = ", j)
}
```

```
$ go run arithmetic/main.go

k =  3
j =  20

k =  4
j =  19
```

So, we can see that the variable `k` is incremented by one and variable `j` is decremented by `1` using the `++` and `--` operator.

### Assignment Operators

These types of operators are quite handy and can condense down large operations into simple expressions. These types of operators allow us to perform operation on the same operand. Let's say we have the variable `k` set to `20` initially, we want to add `30` to the variable `k`, we can do that by using `k = k + 30` but a more sophisticated way would be to use `k += 30` which adds `30` or any value provided the same variable assigned and operated on.

```go
package main

import "fmt"

func main() {
	var a int = 100
	b := 20
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	a += 30
	fmt.Println("a = ", a)
	b -= 5
	fmt.Println("b = ", b)
	a *= b
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	a /= b
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	a %= b
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
```

```
$ go run assignment/main.go

a =  100
b =  20

a =  130
b =  15

a =  1950
b =  15

a =  130
b =  15

a =  10
b =  15
```

From the above example, we are able to perform operations by using shorthand notations like `+=` to add the value to the same operand. These also saves a bit of time and memory not much but considerable enough. This allow us to directly access and modify the contents of the provided operand in the register rather than assigning different registers and performing the operations.

That's it from this part. Reference for all the code examples and commands can be found in the [100 days of Golang](https://github.com/mr-destructive/100-days-of-golang/) GitHub repository.

## Conclusion

So, from the following part of the series, we were able to learn the basics of operators in golang. Using some simple and easy to understand examples, we were able to explore different types of operators like arithmetic, logical, assignment and bitwise operators in golang. These are quite fundamental in programming in general, this lays a good foundation for working with larger and complex projects that deal with any kind of logic in it, without a doubt almost all of the applications do have a bit of logic attached to it. So, we need to know the basics of operators in golang.