---
article_html: "<h2>Introduction</h2>\n<p>We have looked at the defer keyword in golang
  in the <a href=\"https://www.meetgor.com/golang-defer/\">previous</a> part of the
  series, in this section, we will understand how we can use anonymous functions in
  golang. We will explore how to declare and use anonymous functions with a few examples.</p>\n<h2>What
  are Anonymous Functions</h2>\n<p>Anonymous functions are quite simple to understand,
  we don't define a function, we declare it and call it instantly. An anonymous function
  doesn't have a name so hence it is called an anonymous function. As a normal function
  it can take in parameters and return values. With anonymous functions, we can bind
  the operations to a variable or a constant as a literal(value). If an anonymous
  function takes in a parameter, it needs to be parsed immediately after the end of
  the function body. We will see how we define the syntax and specifications of the
  anonymous functions in golang.</p>\n<h2>Simple Anonymous functions</h2>\n<p>To create
  a simple anonymous function we use the same function syntax without giving it a
  name like <code>func() {}</code>, inside the function body i.e. <code>{}</code>,
  you can define the operations that need to be performed.</p>\n<p>Here, we have created
  an anonymous function that simply calls the <code>fmt.Println</code> function with
  a string. So, we have made things a little too much as we can even directly call
  the <code>fmt.Println</code> function from the main function, instead we have called
  an anonymous function that in turn calls the <code>fmt.Println</code> function.
  It might not make sense to use an anonymous function here, but it can be helpful
  for other complex tasks that you want to isolate the logic without creating a dedicated
  function for the process.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">  </span><span class=\"kd\">func</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Hello, Anonymous functions&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"p\">}()</span><span class=\"w\"></span>\n\n<span class=\"p\">}</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run anonymous_function.go\n\nHello,
  Anonymous functions\n</pre></div>\n\n</pre>\n\n<p>So, this is how we create a basic
  anonymous function in golang, this can be further used for returning values from
  the function or passing parameters into the function and also assigning the function
  call to a variable or a constant.</p>\n<h2>Assigning anonymous function to variables</h2>\n<p>We
  can assign the call to the anonymous function to a variable or a constant and call
  the function as many times as we require. So, we can basically store the function
  logic in a variable and call it whenever we require the function with the <code>()</code>
  parenthesis as an indication to call the function.</p>\n<p>In the following example,
  we have used the variable <code>draw</code> to store the function call which prints
  <code>Drawing</code> with the help of the <code>fmt.Println</code> function. The
  draw variable now contains the function and not its value. So be careful here, the
  anonymous function which we have defined as the <code>draw</code> variable's literal
  value, it's like we are giving a name to this anonymous function and the name will
  be the variable name so we have created the function <code>draw</code> which is
  an anonymous function stored in a variable.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">  </span><span class=\"nx\">draw</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"kd\">func</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Drawing&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nx\">draw</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"nx\">draw</span><span class=\"p\">()</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run anonymous_function.go\n\nDrawing\nDrawing\n</pre></div>\n\n</pre>\n\n<p>We
  see that we call the variable <code>draw</code> as a function call by appending
  <code>()</code> parenthesis to it as <code>draw()</code> this thereby calls the
  anonymous function stored inside the variable value.</p>\n<p>The main thing to note
  here is that we are not adding <code>()</code> at the time of declaring the anonymous
  function, as it will call the function directly. The problem then will arise because
  the function is not returning anything so we can't assign it to a variable.</p>\n<pre
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
  class=\"nx\">draw</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"kd\">func</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \ </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Drawing&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}()</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>functions/anonymous_functions.go:10:11:
  func() {…}() (no value) used as value\n</pre></div>\n\n</pre>\n\n<p>So, we are trying
  to assign a variable to a function call that has no return value. This has no value,
  not even nil, so we get an error for assigning a variable to nothing.</p>\n<p>Though
  if you had a return value from the function, we can directly assign the function
  call to the variable as it has returned a valid value defined in the function literal.</p>\n<h2>Parsing
  parameters</h2>\n<p>We can even parse parameters to the anonymous functions as any
  other normal function. We define the name of the variable followed by the type of
  the variable in the parenthesis to use these parameters inside the anonymous function.</p>\n<p>We
  need to keep in mind that these function parameters can either be passed with the
  variable call or directly at the time of defining the function.</p>\n<p>In the below
  example, we have created a variable <code>call</code> and are assigning it to an
  anonymous function that takes in a parameter <code>name</code> which is a <code>string</code>,
  and prints some text to the console. So, we call the funciton <code>call</code>
  with a parameter as a string, as we have <code>&quot;Meet&quot;</code> and <code>person
  := &quot;Chris&quot;</code> as a string passed to the <code>call</code> function.</p>\n<pre
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
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">  </span><span class=\"nx\">call</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"kd\">func</span><span
  class=\"p\">(</span><span class=\"nx\">name</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Calling,&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">name</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">  </span><span class=\"nx\">call</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Meet&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">person</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;Chris&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">call</span><span
  class=\"p\">(</span><span class=\"nx\">person</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run anonymous_function.go\n\nCalling,
  Meet\nCalling, Chris\n</pre></div>\n\n</pre>\n\n<p>Here, we can see that the function
  <code>call</code> prints the text that we have defined to call the <code>fmt.Println</code>
  function. We parse the function with the string literal as the anonymous function
  takes a parameter in the form of a string. We can parse multiple parameters to the
  anonymous function as we can with the normal function like slices, maps, arrays,
  struct, etc.</p>\n<h2>Returning values</h2>\n<p>We can even return values from the
  anonymous function if we want to instantly call the function and save the <code>returned</code>
  value in the variable. We can return single or multiple values as per our requirements
  just like any normal function in golang.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">  </span><span class=\"nx\">is_divisible</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"kd\">func</span><span
  class=\"p\">(</span><span class=\"nx\">x</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">y</span><span class=\"w\">
  </span><span class=\"kt\">int</span><span class=\"p\">)</span><span class=\"w\">
  </span><span class=\"kt\">bool</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">res</span><span class=\"w\"> </span><span class=\"kt\">bool</span><span
  class=\"w\"> </span>\n<span class=\"w\">    </span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">x</span><span class=\"w\"> </span><span class=\"o\">%</span><span
  class=\"w\"> </span><span class=\"nx\">y</span><span class=\"w\"> </span><span class=\"o\">==</span><span
  class=\"w\"> </span><span class=\"mi\">0</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">      </span><span class=\"nx\">res</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"kc\">true</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"> </span><span class=\"k\">else</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">      </span><span class=\"nx\">res</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"kc\">false</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">res</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"nx\">res</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">  </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;%T\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">is_divisible</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">is_divisible</span><span
  class=\"p\">(</span><span class=\"mi\">10</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">5</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"nx\">is_divisible</span><span class=\"p\">(</span><span
  class=\"mi\">33</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"mi\">5</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">  </span><span class=\"nx\">divisblity_check</span><span class=\"w\">
  </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">is_divisible</span><span
  class=\"p\">(</span><span class=\"mi\">45</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">5</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Printf</span><span class=\"p\">(</span><span class=\"s\">&quot;%T :
  %t\\n&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">divisblity_check</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">divisblity_check</span><span
  class=\"p\">)</span><span class=\"w\"> </span>\n\n<span class=\"p\">}</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run anonymous_function.go\n\nfunc(int,
  int) bool \ntrue\nfalse\nbool : true\n</pre></div>\n\n</pre>\n\n<p>As we can see
  the function has returned a boolean value and we store it in a variable <code>divisblity_check</code>,
  the variable which contains the function i.e. <code>is_divisible</code> can then
  be called, and thereby we get the returned value in the variable as a boolean as
  we print the type and the value of the variable <code>divisblity_check</code>. We
  also can see that the type of the variable <code>is_divisible</code> is of type
  function literal which takes <code>int, int</code> and has a return value as <code>bool</code>.</p>\n<p>We
  can also do one more interesting here, which we were restricted previously in the
  above examples. We can directly call the function and store it as a value rather
  than the function itself. So, we can only use the value returned from the function
  but can't call the function later.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nx\">is_divisible</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"kd\">func</span><span class=\"p\">(</span><span
  class=\"nx\">x</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">y</span><span class=\"w\">
  </span><span class=\"kt\">int</span><span class=\"p\">)</span><span class=\"w\">
  </span><span class=\"kt\">bool</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">res</span><span class=\"w\"> </span><span class=\"kt\">bool</span><span
  class=\"w\"> </span>\n<span class=\"w\">  </span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">x</span><span class=\"w\"> </span><span class=\"o\">%</span><span
  class=\"w\"> </span><span class=\"nx\">y</span><span class=\"w\"> </span><span class=\"o\">==</span><span
  class=\"w\"> </span><span class=\"mi\">0</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">res</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"kc\">true</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"p\">}</span><span
  class=\"w\"> </span><span class=\"k\">else</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">res</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"kc\">false</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">res</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"nx\">res</span><span class=\"w\"></span>\n<span class=\"p\">}(</span><span
  class=\"mi\">13</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"mi\">4</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;%T\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">is_divisible</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Printf</span><span class=\"p\">(</span><span class=\"nx\">is_divisible</span><span
  class=\"p\">)</span><span class=\"w\"> </span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run anonymous_function.go\n\nbool\nfalse\n</pre></div>\n\n</pre>\n\n<p>So,
  in the above-modified example, we have passed in the parameter instead of a callable
  function. This will store the returned value of the function call. So, we will store
  the boolean value in the <code>is_divisible</code> and we will have to pass the
  integer values to the function which we have parsed as <code>(13, 4)</code> to the
  anonymous function call.</p>\n<p>In the below example, we have created an anonymous
  function that takes in three parameters like <code>(string, int, string)</code>
  and returns a string. We have used <code>fmt.Sprintf</code> function to format the
  variable and store it in a variable, we then return the string. This anonymous function
  is then directly called and we store the returned value instead of the function.</p>\n<p>So,
  in this example, the <code>format_email</code> variable will store a string instead
  of acting it as a function as a callable object.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">  </span><span class=\"nx\">format_email</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"kd\">func</span><span
  class=\"p\">(</span><span class=\"nx\">name</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">age</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">company</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">email</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Sprintf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s.%d@%s.com&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">name</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">age</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">company</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"nx\">email</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"p\">}(</span><span class=\"s\">&quot;john&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">25</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;gophersoft&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">  </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">format_email</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Printf</span><span class=\"p\">(</span><span class=\"s\">&quot;%T\\n&quot;</span><span
  class=\"p\">,</span><span class=\"nx\">format_email</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run anonymous_function.go\n\njohn.25@gophersoft.com
  \  \nstring\n</pre></div>\n\n</pre>\n\n<p>As we can see the <code>format_email</code>
  variable only returns a string instead of a callable object. We have directly parsed
  the parameters to the anonymous function and hence it instantly calls it and we
  return the string.</p>\n<h2>Passing Anonymous function to function parameters</h2>\n<p>We
  can even pass an anonymous function to a parameter to a function. This can be quite
  helpful for writing some simple logic inside a complex script. We can pass the function
  type as a parameter to the parameter and theere we can parse the actual data and
  perform the desired operation.</p>\n<p>The below example is a bit of code to write
  but makes a lot of sense to understand anonymous functions in golang. The function
  <code>get_csv</code> is a function which takes in three parameters <code>int, string,
  func(string)[] string</code>. The third parameter is a function literal that takes
  in a string and returns a slice of string. So, we are basically converting a string
  <code>&quot;kevin,21,john,33&quot;</code> into a slice of the slice like <code>[[kevin
  21] [john 33]]</code>, this is basically separating values with <code>,</code> comma
  as the delimiter and then processing slices to create a single slice.</p>\n<pre
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"kn\">import</span><span
  class=\"w\"> </span><span class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">get_csv</span><span
  class=\"p\">(</span><span class=\"nx\">index</span><span class=\"w\"> </span><span
  class=\"kt\">int</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">str</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">t</span><span class=\"w\">
  </span><span class=\"kd\">func</span><span class=\"p\">(</span><span class=\"nx\">csv</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">)[]</span><span
  class=\"kt\">string</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"p\">[][]</span><span class=\"kt\">string</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">s</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">t</span><span
  class=\"p\">(</span><span class=\"nx\">str</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">res</span><span class=\"w\"> </span><span
  class=\"p\">[][]</span><span class=\"kt\">string</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">i</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">1</span><span class=\"p\">;</span><span class=\"w\">
  </span><span class=\"nx\">i</span><span class=\"p\">&lt;</span><span class=\"nb\">len</span><span
  class=\"p\">(</span><span class=\"nx\">s</span><span class=\"p\">);</span><span
  class=\"w\"> </span><span class=\"nx\">i</span><span class=\"o\">+=</span><span
  class=\"mi\">2</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">data</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">string</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">data</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nb\">append</span><span class=\"p\">(</span><span
  class=\"nx\">data</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">s</span><span class=\"p\">[</span><span class=\"nx\">i</span><span
  class=\"o\">-</span><span class=\"mi\">1</span><span class=\"p\">],</span><span
  class=\"w\"> </span><span class=\"nx\">s</span><span class=\"p\">[</span><span class=\"nx\">i</span><span
  class=\"p\">])</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">res</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nb\">append</span><span class=\"p\">(</span><span
  class=\"nx\">res</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">data</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"nx\">res</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">  </span><span
  class=\"nx\">csv_slice</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"nx\">csv</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"p\">[]</span><span class=\"kt\">string</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">return</span><span
  class=\"w\"> </span><span class=\"nx\">strings</span><span class=\"p\">.</span><span
  class=\"nx\">Split</span><span class=\"p\">(</span><span class=\"nx\">csv</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;,&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nx\">csv_data</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">get_csv</span><span class=\"p\">(</span><span
  class=\"mi\">2</span><span class=\"p\">,</span><span class=\"s\">&quot;kevin,21,john,33,george,24&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">csv_slice</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">csv_data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">s</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">csv_data</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Printf</span><span class=\"p\">(</span><span class=\"s\">&quot;%s -
  %s\\n&quot;</span><span class=\"p\">,</span><span class=\"nx\">s</span><span class=\"p\">[</span><span
  class=\"mi\">0</span><span class=\"p\">],</span><span class=\"nx\">s</span><span
  class=\"p\">[</span><span class=\"mi\">1</span><span class=\"p\">])</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"p\">}</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run functions/anonymous_function.go\n\n[[kevin
  21] [john 33] [george 24]]\n\nkevin - 21\njohn - 33\ngeorge - 24\n</pre></div>\n\n</pre>\n\n<p>Let's
  break down the code one by one, we will start with the main function, where we have
  <code>csv_slice</code> as a function literal and is an anonymous function that takes
  in a string and returns a slice of string. The function returns a call to the function
  <a href=\"https://pkg.go.dev/strings#Split\">strings.Split</a> taking in the string
  from the function parameter. We then call the function <code>get_csv</code> with
  parameters <code>(2, &quot;kevin,21....&quot;, csv_slice)</code>, this function
  is defined outside the main. The function takes in three parameters as discussed
  and parsed from the main function and it returns a slice of type string.</p>\n<p>So,
  inside the <code>get_csv</code> function, we define <code>s</code> as the function
  call to <code>t(str)</code> which if you look carefully is a function call to <code>csv_slice</code>
  with parameter as a string. This function call returns a slice of strings separated
  by <code>,</code>. So that's all the logic we need to understand anonymous function
  from parameters. We have used a function literal to call the function from another
  function. In this case, the funciton is an anonymous function assigned to a variable.</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>&quot;kevin,21,john,33,george,24&quot;\n
  \           ||\n            \\/\n[kevin 21 john 33 george 24]\n            ||\n
  \           \\/\n[[kevin 21] [john 33] [george 24]]\n</pre></div>\n\n</pre>\n\n<p>Further,
  after we have <code>s</code> which would look like <code>[kevin 21 john 33 george
  24]</code> as each individual element. Thereafter we create an empty <a href=\"https://www.geeksforgeeks.org/slice-of-slices-in-golang/\">slice
  of slice</a> string as <code>res</code> and operate a loop through the slice jumping
  2 indexes at a time. Why? because we want a slice of two elements combined. So,
  we also create a slice of string named <code>data</code> and we add the two components
  to it like <code>[kevin 21]</code> with the <a href=\"https://pkg.go.dev/builtin#append\">append</a>
  function, and this slice is appended to a slice of slices <code>res</code> as <code>[[kevin
  21]]</code> thereby iterating over the entire slice and creating the desired slice
  data. We return the <code>res</code> from the function. This get's us back to the
  main function which simply iterates over the slice and we get the desired data from
  it.</p>\n<p>So, this is how we convert an extremely easy task to extremely complicated
  code :)</p>\n<p>That's it from this part. Reference for all the code examples and
  commands can be found in the <a href=\"https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/functions/anonymous_function.go\">100
  days of Golang</a> GitHub repository.</p>\n<h2>Conclusion</h2>\n<p>That is it from
  this part, we covered a little bit about anonymous functions in golang. Anonymous
  functions are simply function literals that can be used to do a lot of quick operations
  without needing to write an explicit function in the program. Further, in the next
  part look into closures which are a bit complex to understand and require the understanding
  of anonymous functions.</p>\n<p>Thank you for reading, if you have any queries,
  feedback, or questions, you can freely ask me on my social handles. Happy Coding
  :)</p>\n"
cover: ''
date: 2022-07-24
datetime: 2022-07-24 00:00:00+00:00
description: Understanding the concept of the Anonymous functions in golang
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2022-07-24-GO-17-Anonymous-functions.md
html: "<h2>Introduction</h2>\n<p>We have looked at the defer keyword in golang in
  the <a href=\"https://www.meetgor.com/golang-defer/\">previous</a> part of the series,
  in this section, we will understand how we can use anonymous functions in golang.
  We will explore how to declare and use anonymous functions with a few examples.</p>\n<h2>What
  are Anonymous Functions</h2>\n<p>Anonymous functions are quite simple to understand,
  we don't define a function, we declare it and call it instantly. An anonymous function
  doesn't have a name so hence it is called an anonymous function. As a normal function
  it can take in parameters and return values. With anonymous functions, we can bind
  the operations to a variable or a constant as a literal(value). If an anonymous
  function takes in a parameter, it needs to be parsed immediately after the end of
  the function body. We will see how we define the syntax and specifications of the
  anonymous functions in golang.</p>\n<h2>Simple Anonymous functions</h2>\n<p>To create
  a simple anonymous function we use the same function syntax without giving it a
  name like <code>func() {}</code>, inside the function body i.e. <code>{}</code>,
  you can define the operations that need to be performed.</p>\n<p>Here, we have created
  an anonymous function that simply calls the <code>fmt.Println</code> function with
  a string. So, we have made things a little too much as we can even directly call
  the <code>fmt.Println</code> function from the main function, instead we have called
  an anonymous function that in turn calls the <code>fmt.Println</code> function.
  It might not make sense to use an anonymous function here, but it can be helpful
  for other complex tasks that you want to isolate the logic without creating a dedicated
  function for the process.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">  </span><span class=\"kd\">func</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Hello, Anonymous functions&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"p\">}()</span><span class=\"w\"></span>\n\n<span class=\"p\">}</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run anonymous_function.go\n\nHello,
  Anonymous functions\n</pre></div>\n\n</pre>\n\n<p>So, this is how we create a basic
  anonymous function in golang, this can be further used for returning values from
  the function or passing parameters into the function and also assigning the function
  call to a variable or a constant.</p>\n<h2>Assigning anonymous function to variables</h2>\n<p>We
  can assign the call to the anonymous function to a variable or a constant and call
  the function as many times as we require. So, we can basically store the function
  logic in a variable and call it whenever we require the function with the <code>()</code>
  parenthesis as an indication to call the function.</p>\n<p>In the following example,
  we have used the variable <code>draw</code> to store the function call which prints
  <code>Drawing</code> with the help of the <code>fmt.Println</code> function. The
  draw variable now contains the function and not its value. So be careful here, the
  anonymous function which we have defined as the <code>draw</code> variable's literal
  value, it's like we are giving a name to this anonymous function and the name will
  be the variable name so we have created the function <code>draw</code> which is
  an anonymous function stored in a variable.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">  </span><span class=\"nx\">draw</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"kd\">func</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Drawing&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nx\">draw</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"nx\">draw</span><span class=\"p\">()</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run anonymous_function.go\n\nDrawing\nDrawing\n</pre></div>\n\n</pre>\n\n<p>We
  see that we call the variable <code>draw</code> as a function call by appending
  <code>()</code> parenthesis to it as <code>draw()</code> this thereby calls the
  anonymous function stored inside the variable value.</p>\n<p>The main thing to note
  here is that we are not adding <code>()</code> at the time of declaring the anonymous
  function, as it will call the function directly. The problem then will arise because
  the function is not returning anything so we can't assign it to a variable.</p>\n<pre
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
  class=\"nx\">draw</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"kd\">func</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \ </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Drawing&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}()</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>functions/anonymous_functions.go:10:11:
  func() {…}() (no value) used as value\n</pre></div>\n\n</pre>\n\n<p>So, we are trying
  to assign a variable to a function call that has no return value. This has no value,
  not even nil, so we get an error for assigning a variable to nothing.</p>\n<p>Though
  if you had a return value from the function, we can directly assign the function
  call to the variable as it has returned a valid value defined in the function literal.</p>\n<h2>Parsing
  parameters</h2>\n<p>We can even parse parameters to the anonymous functions as any
  other normal function. We define the name of the variable followed by the type of
  the variable in the parenthesis to use these parameters inside the anonymous function.</p>\n<p>We
  need to keep in mind that these function parameters can either be passed with the
  variable call or directly at the time of defining the function.</p>\n<p>In the below
  example, we have created a variable <code>call</code> and are assigning it to an
  anonymous function that takes in a parameter <code>name</code> which is a <code>string</code>,
  and prints some text to the console. So, we call the funciton <code>call</code>
  with a parameter as a string, as we have <code>&quot;Meet&quot;</code> and <code>person
  := &quot;Chris&quot;</code> as a string passed to the <code>call</code> function.</p>\n<pre
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
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">  </span><span class=\"nx\">call</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"kd\">func</span><span
  class=\"p\">(</span><span class=\"nx\">name</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Calling,&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">name</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">  </span><span class=\"nx\">call</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Meet&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">person</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;Chris&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">call</span><span
  class=\"p\">(</span><span class=\"nx\">person</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run anonymous_function.go\n\nCalling,
  Meet\nCalling, Chris\n</pre></div>\n\n</pre>\n\n<p>Here, we can see that the function
  <code>call</code> prints the text that we have defined to call the <code>fmt.Println</code>
  function. We parse the function with the string literal as the anonymous function
  takes a parameter in the form of a string. We can parse multiple parameters to the
  anonymous function as we can with the normal function like slices, maps, arrays,
  struct, etc.</p>\n<h2>Returning values</h2>\n<p>We can even return values from the
  anonymous function if we want to instantly call the function and save the <code>returned</code>
  value in the variable. We can return single or multiple values as per our requirements
  just like any normal function in golang.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">  </span><span class=\"nx\">is_divisible</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"kd\">func</span><span
  class=\"p\">(</span><span class=\"nx\">x</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">y</span><span class=\"w\">
  </span><span class=\"kt\">int</span><span class=\"p\">)</span><span class=\"w\">
  </span><span class=\"kt\">bool</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">res</span><span class=\"w\"> </span><span class=\"kt\">bool</span><span
  class=\"w\"> </span>\n<span class=\"w\">    </span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">x</span><span class=\"w\"> </span><span class=\"o\">%</span><span
  class=\"w\"> </span><span class=\"nx\">y</span><span class=\"w\"> </span><span class=\"o\">==</span><span
  class=\"w\"> </span><span class=\"mi\">0</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">      </span><span class=\"nx\">res</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"kc\">true</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"> </span><span class=\"k\">else</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">      </span><span class=\"nx\">res</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"kc\">false</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">res</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"nx\">res</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">  </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;%T\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">is_divisible</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">is_divisible</span><span
  class=\"p\">(</span><span class=\"mi\">10</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">5</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"nx\">is_divisible</span><span class=\"p\">(</span><span
  class=\"mi\">33</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"mi\">5</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">  </span><span class=\"nx\">divisblity_check</span><span class=\"w\">
  </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">is_divisible</span><span
  class=\"p\">(</span><span class=\"mi\">45</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">5</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Printf</span><span class=\"p\">(</span><span class=\"s\">&quot;%T :
  %t\\n&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">divisblity_check</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">divisblity_check</span><span
  class=\"p\">)</span><span class=\"w\"> </span>\n\n<span class=\"p\">}</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run anonymous_function.go\n\nfunc(int,
  int) bool \ntrue\nfalse\nbool : true\n</pre></div>\n\n</pre>\n\n<p>As we can see
  the function has returned a boolean value and we store it in a variable <code>divisblity_check</code>,
  the variable which contains the function i.e. <code>is_divisible</code> can then
  be called, and thereby we get the returned value in the variable as a boolean as
  we print the type and the value of the variable <code>divisblity_check</code>. We
  also can see that the type of the variable <code>is_divisible</code> is of type
  function literal which takes <code>int, int</code> and has a return value as <code>bool</code>.</p>\n<p>We
  can also do one more interesting here, which we were restricted previously in the
  above examples. We can directly call the function and store it as a value rather
  than the function itself. So, we can only use the value returned from the function
  but can't call the function later.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nx\">is_divisible</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"kd\">func</span><span class=\"p\">(</span><span
  class=\"nx\">x</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">y</span><span class=\"w\">
  </span><span class=\"kt\">int</span><span class=\"p\">)</span><span class=\"w\">
  </span><span class=\"kt\">bool</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">res</span><span class=\"w\"> </span><span class=\"kt\">bool</span><span
  class=\"w\"> </span>\n<span class=\"w\">  </span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">x</span><span class=\"w\"> </span><span class=\"o\">%</span><span
  class=\"w\"> </span><span class=\"nx\">y</span><span class=\"w\"> </span><span class=\"o\">==</span><span
  class=\"w\"> </span><span class=\"mi\">0</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">res</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"kc\">true</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"p\">}</span><span
  class=\"w\"> </span><span class=\"k\">else</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">res</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"kc\">false</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">res</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"nx\">res</span><span class=\"w\"></span>\n<span class=\"p\">}(</span><span
  class=\"mi\">13</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"mi\">4</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Printf</span><span
  class=\"p\">(</span><span class=\"s\">&quot;%T\\n&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">is_divisible</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Printf</span><span class=\"p\">(</span><span class=\"nx\">is_divisible</span><span
  class=\"p\">)</span><span class=\"w\"> </span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run anonymous_function.go\n\nbool\nfalse\n</pre></div>\n\n</pre>\n\n<p>So,
  in the above-modified example, we have passed in the parameter instead of a callable
  function. This will store the returned value of the function call. So, we will store
  the boolean value in the <code>is_divisible</code> and we will have to pass the
  integer values to the function which we have parsed as <code>(13, 4)</code> to the
  anonymous function call.</p>\n<p>In the below example, we have created an anonymous
  function that takes in three parameters like <code>(string, int, string)</code>
  and returns a string. We have used <code>fmt.Sprintf</code> function to format the
  variable and store it in a variable, we then return the string. This anonymous function
  is then directly called and we store the returned value instead of the function.</p>\n<p>So,
  in this example, the <code>format_email</code> variable will store a string instead
  of acting it as a function as a callable object.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">  </span><span class=\"nx\">format_email</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"kd\">func</span><span
  class=\"p\">(</span><span class=\"nx\">name</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">age</span><span class=\"w\"> </span><span class=\"kt\">int</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">company</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">)</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">email</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Sprintf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s.%d@%s.com&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">name</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">age</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">company</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"nx\">email</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"p\">}(</span><span class=\"s\">&quot;john&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">25</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;gophersoft&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">  </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">format_email</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Printf</span><span class=\"p\">(</span><span class=\"s\">&quot;%T\\n&quot;</span><span
  class=\"p\">,</span><span class=\"nx\">format_email</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run anonymous_function.go\n\njohn.25@gophersoft.com
  \  \nstring\n</pre></div>\n\n</pre>\n\n<p>As we can see the <code>format_email</code>
  variable only returns a string instead of a callable object. We have directly parsed
  the parameters to the anonymous function and hence it instantly calls it and we
  return the string.</p>\n<h2>Passing Anonymous function to function parameters</h2>\n<p>We
  can even pass an anonymous function to a parameter to a function. This can be quite
  helpful for writing some simple logic inside a complex script. We can pass the function
  type as a parameter to the parameter and theere we can parse the actual data and
  perform the desired operation.</p>\n<p>The below example is a bit of code to write
  but makes a lot of sense to understand anonymous functions in golang. The function
  <code>get_csv</code> is a function which takes in three parameters <code>int, string,
  func(string)[] string</code>. The third parameter is a function literal that takes
  in a string and returns a slice of string. So, we are basically converting a string
  <code>&quot;kevin,21,john,33&quot;</code> into a slice of the slice like <code>[[kevin
  21] [john 33]]</code>, this is basically separating values with <code>,</code> comma
  as the delimiter and then processing slices to create a single slice.</p>\n<pre
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"kn\">import</span><span
  class=\"w\"> </span><span class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">get_csv</span><span
  class=\"p\">(</span><span class=\"nx\">index</span><span class=\"w\"> </span><span
  class=\"kt\">int</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">str</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">t</span><span class=\"w\">
  </span><span class=\"kd\">func</span><span class=\"p\">(</span><span class=\"nx\">csv</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"p\">)[]</span><span
  class=\"kt\">string</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"p\">[][]</span><span class=\"kt\">string</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nx\">s</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">t</span><span
  class=\"p\">(</span><span class=\"nx\">str</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">res</span><span class=\"w\"> </span><span
  class=\"p\">[][]</span><span class=\"kt\">string</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"k\">for</span><span class=\"w\"> </span><span
  class=\"nx\">i</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"mi\">1</span><span class=\"p\">;</span><span class=\"w\">
  </span><span class=\"nx\">i</span><span class=\"p\">&lt;</span><span class=\"nb\">len</span><span
  class=\"p\">(</span><span class=\"nx\">s</span><span class=\"p\">);</span><span
  class=\"w\"> </span><span class=\"nx\">i</span><span class=\"o\">+=</span><span
  class=\"mi\">2</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">data</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"kt\">string</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">data</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nb\">append</span><span class=\"p\">(</span><span
  class=\"nx\">data</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">s</span><span class=\"p\">[</span><span class=\"nx\">i</span><span
  class=\"o\">-</span><span class=\"mi\">1</span><span class=\"p\">],</span><span
  class=\"w\"> </span><span class=\"nx\">s</span><span class=\"p\">[</span><span class=\"nx\">i</span><span
  class=\"p\">])</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">res</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nb\">append</span><span class=\"p\">(</span><span
  class=\"nx\">res</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">data</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"k\">return</span><span class=\"w\"> </span><span
  class=\"nx\">res</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n\n<span class=\"w\">  </span><span
  class=\"nx\">csv_slice</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"nx\">csv</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"p\">)</span><span class=\"w\"> </span><span
  class=\"p\">[]</span><span class=\"kt\">string</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">return</span><span
  class=\"w\"> </span><span class=\"nx\">strings</span><span class=\"p\">.</span><span
  class=\"nx\">Split</span><span class=\"p\">(</span><span class=\"nx\">csv</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;,&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nx\">csv_data</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">get_csv</span><span class=\"p\">(</span><span
  class=\"mi\">2</span><span class=\"p\">,</span><span class=\"s\">&quot;kevin,21,john,33,george,24&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">csv_slice</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">csv_data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">s</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"k\">range</span><span class=\"w\"> </span><span
  class=\"nx\">csv_data</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Printf</span><span class=\"p\">(</span><span class=\"s\">&quot;%s -
  %s\\n&quot;</span><span class=\"p\">,</span><span class=\"nx\">s</span><span class=\"p\">[</span><span
  class=\"mi\">0</span><span class=\"p\">],</span><span class=\"nx\">s</span><span
  class=\"p\">[</span><span class=\"mi\">1</span><span class=\"p\">])</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"p\">}</span><span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>go run functions/anonymous_function.go\n\n[[kevin
  21] [john 33] [george 24]]\n\nkevin - 21\njohn - 33\ngeorge - 24\n</pre></div>\n\n</pre>\n\n<p>Let's
  break down the code one by one, we will start with the main function, where we have
  <code>csv_slice</code> as a function literal and is an anonymous function that takes
  in a string and returns a slice of string. The function returns a call to the function
  <a href=\"https://pkg.go.dev/strings#Split\">strings.Split</a> taking in the string
  from the function parameter. We then call the function <code>get_csv</code> with
  parameters <code>(2, &quot;kevin,21....&quot;, csv_slice)</code>, this function
  is defined outside the main. The function takes in three parameters as discussed
  and parsed from the main function and it returns a slice of type string.</p>\n<p>So,
  inside the <code>get_csv</code> function, we define <code>s</code> as the function
  call to <code>t(str)</code> which if you look carefully is a function call to <code>csv_slice</code>
  with parameter as a string. This function call returns a slice of strings separated
  by <code>,</code>. So that's all the logic we need to understand anonymous function
  from parameters. We have used a function literal to call the function from another
  function. In this case, the funciton is an anonymous function assigned to a variable.</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>&quot;kevin,21,john,33,george,24&quot;\n
  \           ||\n            \\/\n[kevin 21 john 33 george 24]\n            ||\n
  \           \\/\n[[kevin 21] [john 33] [george 24]]\n</pre></div>\n\n</pre>\n\n<p>Further,
  after we have <code>s</code> which would look like <code>[kevin 21 john 33 george
  24]</code> as each individual element. Thereafter we create an empty <a href=\"https://www.geeksforgeeks.org/slice-of-slices-in-golang/\">slice
  of slice</a> string as <code>res</code> and operate a loop through the slice jumping
  2 indexes at a time. Why? because we want a slice of two elements combined. So,
  we also create a slice of string named <code>data</code> and we add the two components
  to it like <code>[kevin 21]</code> with the <a href=\"https://pkg.go.dev/builtin#append\">append</a>
  function, and this slice is appended to a slice of slices <code>res</code> as <code>[[kevin
  21]]</code> thereby iterating over the entire slice and creating the desired slice
  data. We return the <code>res</code> from the function. This get's us back to the
  main function which simply iterates over the slice and we get the desired data from
  it.</p>\n<p>So, this is how we convert an extremely easy task to extremely complicated
  code :)</p>\n<p>That's it from this part. Reference for all the code examples and
  commands can be found in the <a href=\"https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/functions/anonymous_function.go\">100
  days of Golang</a> GitHub repository.</p>\n<h2>Conclusion</h2>\n<p>That is it from
  this part, we covered a little bit about anonymous functions in golang. Anonymous
  functions are simply function literals that can be used to do a lot of quick operations
  without needing to write an explicit function in the program. Further, in the next
  part look into closures which are a bit complex to understand and require the understanding
  of anonymous functions.</p>\n<p>Thank you for reading, if you have any queries,
  feedback, or questions, you can freely ask me on my social handles. Happy Coding
  :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/golang-017-anonymous-functions.png
long_description: We have looked at the defer keyword in golang in the  Anonymous
  functions are quite simple to understand, we don To create a simple anonymous function
  we use the same function syntax without giving it a name like  Here, we have created
  an anonymous f
now: 2025-05-01 18:11:33.313539
path: blog/posts/2022-07-24-GO-17-Anonymous-functions.md
prevnext: null
series:
- 100-days-of-golang
slug: golang-anonymous-functions
status: published
tags:
- go
templateKey: blog-post
title: 'Golang: Anonymous Functions'
today: 2025-05-01
---

## Introduction

We have looked at the defer keyword in golang in the [previous](https://www.meetgor.com/golang-defer/) part of the series, in this section, we will understand how we can use anonymous functions in golang. We will explore how to declare and use anonymous functions with a few examples.

## What are Anonymous Functions

Anonymous functions are quite simple to understand, we don't define a function, we declare it and call it instantly. An anonymous function doesn't have a name so hence it is called an anonymous function. As a normal function it can take in parameters and return values. With anonymous functions, we can bind the operations to a variable or a constant as a literal(value). If an anonymous function takes in a parameter, it needs to be parsed immediately after the end of the function body. We will see how we define the syntax and specifications of the anonymous functions in golang. 

## Simple Anonymous functions

To create a simple anonymous function we use the same function syntax without giving it a name like `func() {}`, inside the function body i.e. `{}`, you can define the operations that need to be performed. 

Here, we have created an anonymous function that simply calls the `fmt.Println` function with a string. So, we have made things a little too much as we can even directly call the `fmt.Println` function from the main function, instead we have called an anonymous function that in turn calls the `fmt.Println` function. It might not make sense to use an anonymous function here, but it can be helpful for other complex tasks that you want to isolate the logic without creating a dedicated function for the process.

```go
package main

import "fmt"

func main() {

  func() {
    fmt.Println("Hello, Anonymous functions")
  }()

}
```

```
go run anonymous_function.go

Hello, Anonymous functions
```

So, this is how we create a basic anonymous function in golang, this can be further used for returning values from the function or passing parameters into the function and also assigning the function call to a variable or a constant.

## Assigning anonymous function to variables

We can assign the call to the anonymous function to a variable or a constant and call the function as many times as we require. So, we can basically store the function logic in a variable and call it whenever we require the function with the `()` parenthesis as an indication to call the function.

In the following example, we have used the variable `draw` to store the function call which prints `Drawing` with the help of the `fmt.Println` function. The draw variable now contains the function and not its value. So be careful here, the anonymous function which we have defined as the `draw` variable's literal value, it's like we are giving a name to this anonymous function and the name will be the variable name so we have created the function `draw` which is an anonymous function stored in a variable.


```go
package main

import "fmt"

func main() {

  draw := func() {
    fmt.Println("Drawing")
  }
  draw()
  draw()
}
```

```
go run anonymous_function.go

Drawing
Drawing
```

We see that we call the variable `draw` as a function call by appending `()` parenthesis to it as `draw()` this thereby calls the anonymous function stored inside the variable value.

The main thing to note here is that we are not adding `()` at the time of declaring the anonymous function, as it will call the function directly. The problem then will arise because the function is not returning anything so we can't assign it to a variable.

```go
draw := func() {
  fmt.Println("Drawing")
}()
```

```
functions/anonymous_functions.go:10:11: func() {…}() (no value) used as value
```

So, we are trying to assign a variable to a function call that has no return value. This has no value, not even nil, so we get an error for assigning a variable to nothing.

Though if you had a return value from the function, we can directly assign the function call to the variable as it has returned a valid value defined in the function literal.

## Parsing parameters

We can even parse parameters to the anonymous functions as any other normal function. We define the name of the variable followed by the type of the variable in the parenthesis to use these parameters inside the anonymous function.

We need to keep in mind that these function parameters can either be passed with the variable call or directly at the time of defining the function.

In the below example, we have created a variable `call` and are assigning it to an anonymous function that takes in a parameter `name` which is a `string`, and prints some text to the console. So, we call the funciton `call` with a parameter as a string, as we have `"Meet"` and `person := "Chris"` as a string passed to the `call` function.

```go
package main

import "fmt"

func main() {

  call := func(name string) {
    fmt.Println("Calling,", name)
  }

  call("Meet")
  person := "Chris"
  call(person)

}
```

```
go run anonymous_function.go

Calling, Meet
Calling, Chris
```

Here, we can see that the function `call` prints the text that we have defined to call the `fmt.Println` function. We parse the function with the string literal as the anonymous function takes a parameter in the form of a string. We can parse multiple parameters to the anonymous function as we can with the normal function like slices, maps, arrays, struct, etc.

## Returning values

We can even return values from the anonymous function if we want to instantly call the function and save the `returned` value in the variable. We can return single or multiple values as per our requirements just like any normal function in golang.


```go
package main

import "fmt"

func main() {

  is_divisible := func(x int, y int) bool{
    var res bool 
    if x % y == 0 {
      res = true
    } else{
      res = false
    }
    fmt.Println(res)
    return res
  }

  fmt.Printf("%T\n", is_divisible)
  is_divisible(10, 5)
  is_divisible(33, 5)

  divisblity_check := is_divisible(45, 5)
  fmt.Printf("%T : %t\n", divisblity_check, divisblity_check) 

}
```

```
go run anonymous_function.go

func(int, int) bool 
true
false
bool : true
```

As we can see the function has returned a boolean value and we store it in a variable `divisblity_check`, the variable which contains the function i.e. `is_divisible` can then be called, and thereby we get the returned value in the variable as a boolean as we print the type and the value of the variable `divisblity_check`. We also can see that the type of the variable `is_divisible` is of type function literal which takes `int, int` and has a return value as `bool`.

We can also do one more interesting here, which we were restricted previously in the above examples. We can directly call the function and store it as a value rather than the function itself. So, we can only use the value returned from the function but can't call the function later.

```go
is_divisible := func(x int, y int) bool{
  var res bool 
  if x % y == 0 {
    res = true
  } else{
    res = false
  }
  fmt.Println(res)
  return res
}(13, 4)

fmt.Printf("%T\n", is_divisible)
fmt.Printf(is_divisible) 
```

```
go run anonymous_function.go

bool
false
```

So, in the above-modified example, we have passed in the parameter instead of a callable function. This will store the returned value of the function call. So, we will store the boolean value in the `is_divisible` and we will have to pass the integer values to the function which we have parsed as `(13, 4)` to the anonymous function call.


In the below example, we have created an anonymous function that takes in three parameters like `(string, int, string)` and returns a string. We have used `fmt.Sprintf` function to format the variable and store it in a variable, we then return the string. This anonymous function is then directly called and we store the returned value instead of the function.

So, in this example, the `format_email` variable will store a string instead of acting it as a function as a callable object.

```go
package main

import "fmt"

func main() {

  format_email := func(name string, age int, company string) string{
    email := fmt.Sprintf("%s.%d@%s.com", name, age, company)
    return email
  }("john", 25, "gophersoft")

  fmt.Println(format_email)
  fmt.Printf("%T\n",format_email)

}
```

```
go run anonymous_function.go

john.25@gophersoft.com   
string
```

As we can see the `format_email` variable only returns a string instead of a callable object. We have directly parsed the parameters to the anonymous function and hence it instantly calls it and we return the string.

## Passing Anonymous function to function parameters

We can even pass an anonymous function to a parameter to a function. This can be quite helpful for writing some simple logic inside a complex script. We can pass the function type as a parameter to the parameter and theere we can parse the actual data and perform the desired operation.

The below example is a bit of code to write but makes a lot of sense to understand anonymous functions in golang. The function `get_csv` is a function which takes in three parameters `int, string, func(string)[] string`. The third parameter is a function literal that takes in a string and returns a slice of string. So, we are basically converting a string `"kevin,21,john,33"` into a slice of the slice like `[[kevin 21] [john 33]]`, this is basically separating values with `,` comma as the delimiter and then processing slices to create a single slice.

```go
package main

import "fmt"
import "strings"

func get_csv(index int, str string, t func(csv string)[]string) [][]string{
  s := t(str)
  var res [][]string
  for i := 1; i<len(s); i+=2 {
    var data []string
    data = append(data, s[i-1], s[i])
    res = append(res, data)
  }
  return res
}

func main() {

  csv_slice := func (csv string) []string{
    return strings.Split(csv, ",")
  }
  csv_data := get_csv(2,"kevin,21,john,33,george,24", csv_slice)
  fmt.Println(csv_data)
  for _, s := range csv_data{
    fmt.Printf("%s - %s\n",s[0],s[1])
  }
```

```
go run functions/anonymous_function.go

[[kevin 21] [john 33] [george 24]]

kevin - 21
john - 33
george - 24
```

Let's break down the code one by one, we will start with the main function, where we have `csv_slice` as a function literal and is an anonymous function that takes in a string and returns a slice of string. The function returns a call to the function [strings.Split](https://pkg.go.dev/strings#Split) taking in the string from the function parameter. We then call the function `get_csv` with parameters `(2, "kevin,21....", csv_slice)`, this function is defined outside the main. The function takes in three parameters as discussed and parsed from the main function and it returns a slice of type string. 

So, inside the `get_csv` function, we define `s` as the function call to `t(str)` which if you look carefully is a function call to `csv_slice` with parameter as a string. This function call returns a slice of strings separated by `,`. So that's all the logic we need to understand anonymous function from parameters. We have used a function literal to call the function from another function. In this case, the funciton is an anonymous function assigned to a variable. 

```
"kevin,21,john,33,george,24"
            ||
            \/
[kevin 21 john 33 george 24]
            ||
            \/
[[kevin 21] [john 33] [george 24]]

```

Further, after we have `s` which would look like `[kevin 21 john 33 george 24]` as each individual element. Thereafter we create an empty [slice of slice](https://www.geeksforgeeks.org/slice-of-slices-in-golang/) string as `res` and operate a loop through the slice jumping 2 indexes at a time. Why? because we want a slice of two elements combined. So, we also create a slice of string named `data` and we add the two components to it like `[kevin 21]` with the [append](https://pkg.go.dev/builtin#append) function, and this slice is appended to a slice of slices `res` as `[[kevin 21]]` thereby iterating over the entire slice and creating the desired slice data. We return the `res` from the function. This get's us back to the main function which simply iterates over the slice and we get the desired data from it.

So, this is how we convert an extremely easy task to extremely complicated code :)

That's it from this part. Reference for all the code examples and commands can be found in the [100 days of Golang](https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/functions/anonymous_function.go) GitHub repository.

## Conclusion

That is it from this part, we covered a little bit about anonymous functions in golang. Anonymous functions are simply function literals that can be used to do a lot of quick operations without needing to write an explicit function in the program. Further, in the next part look into closures which are a bit complex to understand and require the understanding of anonymous functions.

Thank you for reading, if you have any queries, feedback, or questions, you can freely ask me on my social handles. Happy Coding :)