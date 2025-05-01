---
article_html: "<h2>Introduction</h2>\n<p>In this section of the series, we will be
  exploring how to send a <code>PUT</code> HTTP request in golang. We will understand
  how to send a basic PUT request, create an HTTP request, update a resource on a
  server, parsing the content from struct to json, headers, etc in the following section
  of this post.</p>\n<h2>What is a PUT Method</h2>\n<p>A PUT method is a type of request
  that is used to update or modify an entire resource on a server/database.</p>\n<p>Imagine
  you have ordered a pizza at a restaurant and realized you want to change the toppings
  after you've already placed the order. With a PUT request, it's like informing the
  waiter about the changes you want to make to your existing order. You specify the
  updated toppings or any modifications (the data you send). The waiter then takes
  this updated information (PUT request) back to the kitchen (the server) to apply
  the changes to your order.</p>\n<p>Let's say you created a order.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>nginx</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">PUT</span><span class=\"w\"> </span><span class=\"s\">/api/order/456</span><span
  class=\"w\"> </span><span class=\"s\">HTTP/1.1</span><span class=\"w\"></span>\n<span
  class=\"s\">Host:</span><span class=\"w\"> </span><span class=\"s\">example.com</span><span
  class=\"w\"></span>\n<span class=\"s\">Content-Type:</span><span class=\"w\"> </span><span
  class=\"s\">application/json</span><span class=\"w\"></span>\n<span class=\"s\">Content-Length:</span><span
  class=\"w\"> </span><span class=\"mi\">123</span><span class=\"w\"></span>\n\n<span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"kn\">&quot;userID&quot;:</span><span class=\"w\"> </span><span class=\"mi\">123</span><span
  class=\"s\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;orderID&quot;:</span><span class=\"w\"> </span><span class=\"mi\">456</span><span
  class=\"s\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;items&quot;:</span><span class=\"w\"> </span><span class=\"s\">[</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">            </span><span class=\"kn\">&quot;itemID&quot;:</span><span
  class=\"w\"> </span><span class=\"mi\">789</span><span class=\"s\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">            </span><span class=\"s\">&quot;name&quot;:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Pizza&quot;,</span><span class=\"w\"></span>\n<span
  class=\"w\">            </span><span class=\"s\">&quot;quantity&quot;:</span><span
  class=\"w\"> </span><span class=\"mi\">2</span><span class=\"s\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">            </span><span class=\"s\">&quot;toppings&quot;:</span><span
  class=\"w\"> </span><span class=\"s\">[&quot;Mushrooms&quot;]</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"err\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"s\">]</span><span class=\"w\"></span>\n<span
  class=\"err\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  the context of web development, PUT requests are often used for actions such as:</p>\n<ul>\n<li>\n<p>Updating
  existing records or resources</p>\n</li>\n<li>\n<p>Modifying specific parts of an
  existing resource</p>\n</li>\n<li>\n<p>Replacing an entire resource with updated
  data</p>\n</li>\n</ul>\n<p>Here's an example of what the PUT request might look
  like in this scenario:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"k\">PUT</span><span class=\"w\"> </span><span class=\"s\">/api/order/456</span><span
  class=\"w\"> </span><span class=\"s\">HTTP/1.1</span><span class=\"w\"></span>\n<span
  class=\"s\">Host:</span><span class=\"w\"> </span><span class=\"s\">example.com</span><span
  class=\"w\"></span>\n<span class=\"s\">Content-Type:</span><span class=\"w\"> </span><span
  class=\"s\">application/json</span><span class=\"w\"></span>\n<span class=\"s\">Content-Length:</span><span
  class=\"w\"> </span><span class=\"mi\">155</span><span class=\"w\"></span>\n\n<span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"kn\">&quot;userID&quot;:</span><span class=\"w\"> </span><span class=\"mi\">123</span><span
  class=\"s\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;orderID&quot;:</span><span class=\"w\"> </span><span class=\"mi\">456</span><span
  class=\"s\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;items&quot;:</span><span class=\"w\"> </span><span class=\"s\">[</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">            </span><span class=\"kn\">&quot;itemID&quot;:</span><span
  class=\"w\"> </span><span class=\"mi\">789</span><span class=\"s\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">            </span><span class=\"s\">&quot;name&quot;:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Pizza&quot;,</span><span class=\"w\"></span>\n<span
  class=\"w\">            </span><span class=\"s\">&quot;quantity&quot;:</span><span
  class=\"w\"> </span><span class=\"mi\">2</span><span class=\"s\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">            </span><span class=\"s\">&quot;toppings&quot;:</span><span
  class=\"w\"> </span><span class=\"s\">[&quot;Mushrooms&quot;,</span><span class=\"w\">
  </span><span class=\"s\">&quot;Olives&quot;]</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"err\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"s\">]</span><span class=\"w\"></span>\n<span
  class=\"err\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  this example:</p>\n<ul>\n<li>\n<p>The PUT method is used to update the resource
  identified by <code>/api/order/456</code>.</p>\n</li>\n<li>\n<p>The application/json
  is the content type of the request.</p>\n</li>\n<li>\n<p>The 155 is the content
  length of the request.</p>\n</li>\n<li>\n<p>The body contains the updated details
  of the order, including the addition of toppings to the pizza.</p>\n</li>\n</ul>\n<p>PUT
  requests are crucial for maintaining and updating data in applications where accuracy
  and consistency are paramount, ensuring that resources are kept current and reflect
  the latest changes made by users or systems</p>\n<h2>Why the need of PUT Method</h2>\n<p>In
  the world of HTTP requests, we use the PUT method to update or modify an entire
  resource on a server or database. This method is crucial because the POST method,
  while convenient for creating new data, is not intended for updating existing resources
  according to standard conventions. While it's possible to misuse the POST method
  for updates internally, doing so can lead to confusion and inconsistencies in how
  requests are understood and processed.</p>\n<h2>How PUT Method request works</h2>\n<p>A
  <a href=\"https://www.rfc-editor.org/rfc/rfc9110#PUT\">PUT</a> request is utilized
  to send data to a server for the purpose of updating a resource. When a client (such
  as a browser or other APIs) sends a PUT request to the server's API endpoint, it
  includes data in the request body, typically formatted as JSON, XML, or form data.</p>\n<p>The
  server processes the PUT request by first identifying the resource to be updated,
  either through the URL or data provided in the request body. It then validates,
  parses, and applies the data from the request body to make modifications to the
  resource. Following this, the server returns a response that includes a status code
  indicating the success or failure of the operation. Optionally, the response may
  also include the updated resource in the response body.</p>\n<p>Unlike the POST
  method, which is primarily used for creating new resources, PUT is specifically
  designed for updating existing resources on the server. The request body of a PUT
  contains the data necessary for the update, while the URL identifies the specific
  resource to be updated.</p>\n<p>In summary, PUT requests facilitate the transfer
  of data to the server specifically for updating resources, ensuring that changes
  to existing data are accurately processed and reflected.</p>\n<h2>Basic PUT Method</h2>\n<p>To
  send a <code>PUT</code> request to an API in golang, we need to create a <code>http.Request</code>
  object. For <code>POST</code> method, the <code>http</code> package had the <code>Post</code>
  function defined, however for <code>PUT</code> method, there is no separate function.
  The Go philosophy is right now against adding all the method functions. There have
  been a couple of discussions on this on <a href=\"https://github.com/golang/go/issues/22841\">GitHub</a>,
  but it is not been adopted as of 2024.</p>\n<p>So, we need to create a <code>http.Request</code>
  object for <code>PUT</code> method.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in/api/users/5&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">MethodPut</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">DefaultClient</span><span class=\"p\">.</span><span class=\"nx\">Do</span><span
  class=\"p\">(</span><span class=\"nx\">req</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">.</span><span class=\"nx\">Close</span><span class=\"p\">()</span><span
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
  go run main.go\n\n<span class=\"m\">200</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above code sends a <code>PUT</code> request to the <a href=\"https://reqres.in/api/users/5\"><code>https://reqres.in/api/users/5</code></a>
  endpoint. The resource we are trying to update is fetched with the identifier <code>5</code>
  which could probably be <code>id</code> for the user in the database of the server.</p>\n<h2>PUT
  Method with JSON</h2>\n<p>Marshaling and encoding are essential in Go for preparing
  structured data, such as from a struct, into JSON format suitable for HTTP requests
  like PUT. This conversion ensures data integrity and compatibility between Go types
  and JSON representations. It's crucial when updating resources on servers, as APIs
  often require specific data formats for processing updates correctly. Marshaling
  transforms Go structs into JSON bytes, while encoding further prepares them as request
  bodies, facilitating seamless communication with web services. This process ensures
  data consistency and adherence to API specifications, maintaining robust communication
  in distributed systems.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;bytes&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">User</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Name</span><span class=\"w\">   </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;name&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Salary</span><span class=\"w\"> </span><span
  class=\"kt\">int</span><span class=\"w\">    </span><span class=\"s\">`json:&quot;salary&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Age</span><span
  class=\"w\">    </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;age&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">     </span><span class=\"kt\">int</span><span
  class=\"w\">    </span><span class=\"s\">`json:&quot;id,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">type</span><span class=\"w\"> </span><span class=\"nx\">UserResponse</span><span
  class=\"w\"> </span><span class=\"kd\">struct</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Status</span><span class=\"w\">  </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;status&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Message</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;message&quot;`</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">user</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">User</span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">Name</span><span class=\"p\">:</span><span class=\"w\">   </span><span
  class=\"s\">&quot;Alice&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">Salary</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">50000</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">Age</span><span
  class=\"p\">:</span><span class=\"w\">    </span><span class=\"s\">&quot;25&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy.restapiexample.com/api/v1/update/11&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">// marshalling
  process</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">//
  converting Go specific data structure/types to JSON</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">bodyBytes</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Marshal</span><span class=\"p\">(</span><span
  class=\"nx\">user</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">bodyBytes</span><span class=\"p\">))</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">// reading json into a buffer/in-memory</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">body</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">bytes</span><span
  class=\"p\">.</span><span class=\"nx\">NewBuffer</span><span class=\"p\">(</span><span
  class=\"nx\">bodyBytes</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">// post request</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">MethodPut</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">DefaultClient</span><span class=\"p\">.</span><span class=\"nx\">Do</span><span
  class=\"p\">(</span><span class=\"nx\">req</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">StatusCode</span><span
  class=\"w\"> </span><span class=\"o\">==</span><span class=\"w\"> </span><span class=\"mi\">429</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;too many requests&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">io</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">.</span><span class=\"nx\">Close</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">// unmarshalling
  process</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">//
  converting JSON to Go specific data structure/types</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">userResponse</span><span class=\"w\"> </span><span class=\"nx\">UserResponse</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">userResponse</span><span class=\"p\">);</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">userResponse</span><span class=\"p\">)</span><span
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
  go run json.go                                                                 \n<span
  class=\"o\">{</span><span class=\"s2\">&quot;name&quot;</span>:<span class=\"s2\">&quot;Alice&quot;</span>,<span
  class=\"s2\">&quot;salary&quot;</span>:50000,<span class=\"s2\">&quot;age&quot;</span>:<span
  class=\"s2\">&quot;25&quot;</span><span class=\"o\">}</span>\n<span class=\"m\">200</span>
  \                                                                             \n<span
  class=\"o\">{</span><span class=\"s2\">&quot;status&quot;</span>:<span class=\"s2\">&quot;success&quot;</span>,<span
  class=\"s2\">&quot;data&quot;</span>:<span class=\"o\">[]</span>,<span class=\"s2\">&quot;message&quot;</span>:<span
  class=\"s2\">&quot;Successfully! Record has been updated.&quot;</span><span class=\"o\">}</span>\n<span
  class=\"o\">{</span>success Successfully! Record has been updated.<span class=\"o\">}</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the provided Go code example, the <code>json.Marshal</code> function is used to
  convert a Go struct (<code>User</code>) into a JSON formatted byte slice (<code>[]byte</code>).
  Here's a breakdown of the steps involved:</p>\n<ul>\n<li>\n<p>Struct Definition:
  Define a Go struct with json tags.</p>\n</li>\n<li>\n<p>Marshalling: Use json.Marshal
  to convert the struct into JSON byte slice.</p>\n</li>\n<li>\n<p>Buffer Creation:
  Wrap the JSON byte slice into an in-memory buffer (bytes.Buffer).</p>\n</li>\n<li>\n<p>Request
  Sending: Send a PUT request with the buffer as the request body and set appropriate
  headers.</p>\n</li>\n</ul>\n<p>Let's explore it step by step in detail:</p>\n<p>When
  using the PUT method in Go to update a resource on a server, you often need to send
  data in JSON format as the request body. Here's how you can achieve this using marshaling
  and encoding:</p>\n<ol>\n<li>Define the Struct</li>\n</ol>\n<pre class='wrapper'>\n\n<div
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
  class=\"kd\">type</span><span class=\"w\"> </span><span class=\"nx\">User</span><span
  class=\"w\"> </span><span class=\"kd\">struct</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">Name</span><span class=\"w\">   </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;name&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">Salary</span><span class=\"w\"> </span><span
  class=\"kt\">int</span><span class=\"w\">    </span><span class=\"s\">`json:&quot;salary&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">Age</span><span
  class=\"w\">    </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;age&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nx\">ID</span><span class=\"w\">     </span><span class=\"kt\">int</span><span
  class=\"w\">    </span><span class=\"s\">`json:&quot;id,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Define
  a Go struct (<code>User</code>) that represents the data structure you want to send
  in JSON format. The json tags specify how each field should be serialized into JSON.</p>\n<ol
  start=\"2\">\n<li>Create an Object</li>\n</ol>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nx\">user</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">User</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">Name</span><span
  class=\"p\">:</span><span class=\"w\">   </span><span class=\"s\">&quot;Alice&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">Salary</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">50000</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">Age</span><span class=\"p\">:</span><span
  class=\"w\">    </span><span class=\"s\">&quot;25&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Create
  an instance of the User struct (user) with sample data. This data will be marshaled
  into JSON format to send in the <code>PUT</code> request body.</p>\n<ol start=\"3\">\n<li>Marshal
  the Struct</li>\n</ol>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nx\">bodyBytes</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Marshal</span><span class=\"p\">(</span><span class=\"nx\">user</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Use
  json.Marshal(user) to convert the user struct into a JSON byte slice (bodyBytes).
  This byte slice contains the serialized JSON representation of the User struct.</p>\n<ol
  start=\"4\">\n<li>Create a Buffer</li>\n</ol>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nx\">body</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bytes</span><span class=\"p\">.</span><span
  class=\"nx\">NewBuffer</span><span class=\"p\">(</span><span class=\"nx\">bodyBytes</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Use
  <code>bytes.NewBuffer(bodyBytes)</code> to create an in-memory buffer (<code>body</code>)
  containing the JSON byte slice (<code>bodyBytes</code>). The buffer implements the
  <code>io.Reader</code> interface needed for the PUT request body.</p>\n<ol start=\"5\">\n<li>Create
  a PUT Request</li>\n</ol>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">MethodPut</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">apiURL</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Use
  http.NewRequest to create a new PUT request to the specified URL with the JSON buffer
  (<code>body</code>) as the request body. Set appropriate headers if needed (e.g.,
  Content-Type as application/json).</p>\n<ol start=\"6\">\n<li>Send the Request</li>\n</ol>\n<pre
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
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">DefaultClient</span><span class=\"p\">.</span><span class=\"nx\">Do</span><span
  class=\"p\">(</span><span class=\"nx\">req</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Use
  <a href=\"http://http.DefaultClient.Do\"><code>http.DefaultClient.Do</code></a><code>(req)</code>
  to execute the PUT request and obtain the response. Handle any errors that may occur
  during the request execution.</p>\n<ol start=\"7\">\n<li>Process the Response</li>\n</ol>\n<pre
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
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">io</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Use
  <code>io.ReadAll(resp.Body)</code> to read and store the response body from the
  server. Handle any errors encountered during the reading process.</p>\n<ol start=\"8\">\n<li>Unmarshal
  the Response</li>\n</ol>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">userResponse</span><span
  class=\"w\"> </span><span class=\"nx\">UserResponse</span><span class=\"w\"></span>\n<span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">userResponse</span><span class=\"p\">);</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Use
  <code>json.Unmarshal(respBody, &amp;userResponse)</code> to deserialize the JSON
  response body into a Go struct. This allows you to work with the response data in
  a structured manner.</p>\n<p>The parsing of files and form data is also possible
  with <code>PUT</code> requests, however, that has been covered in the <a href=\"https://meetgor.com/golang-web-post-method\">POST
  Method</a>. Those snippets would be handy in these request method as well.</p>\n<p>I
  have also included some more examples of PUT requests <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/put/\">here</a>.</p>\n<p>That's
  it from the 35th part of the series, all the source code for the examples are linked
  in the GitHub on the <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/put/\">100
  days of Golang</a> repository.</p>\n<p><a href=\"https://github.com/Mr-Destructive/100-days-of-golang\">100-days-of-golang</a></p>\n<h2>Conclusion</h2>\n<p>That's
  it from this post of the series, a post on the PUT method in golang :)</p>\n<p>We
  have covered topics like creating basic PUT requests and marshaling golang types
  into JSON format. Hope you found this article helpful. If you have any queries,
  questions, or feedback, please let me know in the comments or on my social handles.
  Thank you for reading.</p>\n<p>Happy Coding :)</p>\n"
cover: ''
date: 2024-06-15
datetime: 2024-06-15 00:00:00+00:00
description: Exploring the fundamentals of a PUT method in golang. How to request
  a resource, parse body, headers, etc. in a HTTP PUT request.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2024-06-15-GO-35-PUT-Method.md
html: "<h2>Introduction</h2>\n<p>In this section of the series, we will be exploring
  how to send a <code>PUT</code> HTTP request in golang. We will understand how to
  send a basic PUT request, create an HTTP request, update a resource on a server,
  parsing the content from struct to json, headers, etc in the following section of
  this post.</p>\n<h2>What is a PUT Method</h2>\n<p>A PUT method is a type of request
  that is used to update or modify an entire resource on a server/database.</p>\n<p>Imagine
  you have ordered a pizza at a restaurant and realized you want to change the toppings
  after you've already placed the order. With a PUT request, it's like informing the
  waiter about the changes you want to make to your existing order. You specify the
  updated toppings or any modifications (the data you send). The waiter then takes
  this updated information (PUT request) back to the kitchen (the server) to apply
  the changes to your order.</p>\n<p>Let's say you created a order.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>nginx</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">PUT</span><span class=\"w\"> </span><span class=\"s\">/api/order/456</span><span
  class=\"w\"> </span><span class=\"s\">HTTP/1.1</span><span class=\"w\"></span>\n<span
  class=\"s\">Host:</span><span class=\"w\"> </span><span class=\"s\">example.com</span><span
  class=\"w\"></span>\n<span class=\"s\">Content-Type:</span><span class=\"w\"> </span><span
  class=\"s\">application/json</span><span class=\"w\"></span>\n<span class=\"s\">Content-Length:</span><span
  class=\"w\"> </span><span class=\"mi\">123</span><span class=\"w\"></span>\n\n<span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"kn\">&quot;userID&quot;:</span><span class=\"w\"> </span><span class=\"mi\">123</span><span
  class=\"s\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;orderID&quot;:</span><span class=\"w\"> </span><span class=\"mi\">456</span><span
  class=\"s\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;items&quot;:</span><span class=\"w\"> </span><span class=\"s\">[</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">            </span><span class=\"kn\">&quot;itemID&quot;:</span><span
  class=\"w\"> </span><span class=\"mi\">789</span><span class=\"s\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">            </span><span class=\"s\">&quot;name&quot;:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Pizza&quot;,</span><span class=\"w\"></span>\n<span
  class=\"w\">            </span><span class=\"s\">&quot;quantity&quot;:</span><span
  class=\"w\"> </span><span class=\"mi\">2</span><span class=\"s\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">            </span><span class=\"s\">&quot;toppings&quot;:</span><span
  class=\"w\"> </span><span class=\"s\">[&quot;Mushrooms&quot;]</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"err\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"s\">]</span><span class=\"w\"></span>\n<span
  class=\"err\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  the context of web development, PUT requests are often used for actions such as:</p>\n<ul>\n<li>\n<p>Updating
  existing records or resources</p>\n</li>\n<li>\n<p>Modifying specific parts of an
  existing resource</p>\n</li>\n<li>\n<p>Replacing an entire resource with updated
  data</p>\n</li>\n</ul>\n<p>Here's an example of what the PUT request might look
  like in this scenario:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"k\">PUT</span><span class=\"w\"> </span><span class=\"s\">/api/order/456</span><span
  class=\"w\"> </span><span class=\"s\">HTTP/1.1</span><span class=\"w\"></span>\n<span
  class=\"s\">Host:</span><span class=\"w\"> </span><span class=\"s\">example.com</span><span
  class=\"w\"></span>\n<span class=\"s\">Content-Type:</span><span class=\"w\"> </span><span
  class=\"s\">application/json</span><span class=\"w\"></span>\n<span class=\"s\">Content-Length:</span><span
  class=\"w\"> </span><span class=\"mi\">155</span><span class=\"w\"></span>\n\n<span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"kn\">&quot;userID&quot;:</span><span class=\"w\"> </span><span class=\"mi\">123</span><span
  class=\"s\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;orderID&quot;:</span><span class=\"w\"> </span><span class=\"mi\">456</span><span
  class=\"s\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;items&quot;:</span><span class=\"w\"> </span><span class=\"s\">[</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">            </span><span class=\"kn\">&quot;itemID&quot;:</span><span
  class=\"w\"> </span><span class=\"mi\">789</span><span class=\"s\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">            </span><span class=\"s\">&quot;name&quot;:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Pizza&quot;,</span><span class=\"w\"></span>\n<span
  class=\"w\">            </span><span class=\"s\">&quot;quantity&quot;:</span><span
  class=\"w\"> </span><span class=\"mi\">2</span><span class=\"s\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">            </span><span class=\"s\">&quot;toppings&quot;:</span><span
  class=\"w\"> </span><span class=\"s\">[&quot;Mushrooms&quot;,</span><span class=\"w\">
  </span><span class=\"s\">&quot;Olives&quot;]</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"err\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"s\">]</span><span class=\"w\"></span>\n<span
  class=\"err\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  this example:</p>\n<ul>\n<li>\n<p>The PUT method is used to update the resource
  identified by <code>/api/order/456</code>.</p>\n</li>\n<li>\n<p>The application/json
  is the content type of the request.</p>\n</li>\n<li>\n<p>The 155 is the content
  length of the request.</p>\n</li>\n<li>\n<p>The body contains the updated details
  of the order, including the addition of toppings to the pizza.</p>\n</li>\n</ul>\n<p>PUT
  requests are crucial for maintaining and updating data in applications where accuracy
  and consistency are paramount, ensuring that resources are kept current and reflect
  the latest changes made by users or systems</p>\n<h2>Why the need of PUT Method</h2>\n<p>In
  the world of HTTP requests, we use the PUT method to update or modify an entire
  resource on a server or database. This method is crucial because the POST method,
  while convenient for creating new data, is not intended for updating existing resources
  according to standard conventions. While it's possible to misuse the POST method
  for updates internally, doing so can lead to confusion and inconsistencies in how
  requests are understood and processed.</p>\n<h2>How PUT Method request works</h2>\n<p>A
  <a href=\"https://www.rfc-editor.org/rfc/rfc9110#PUT\">PUT</a> request is utilized
  to send data to a server for the purpose of updating a resource. When a client (such
  as a browser or other APIs) sends a PUT request to the server's API endpoint, it
  includes data in the request body, typically formatted as JSON, XML, or form data.</p>\n<p>The
  server processes the PUT request by first identifying the resource to be updated,
  either through the URL or data provided in the request body. It then validates,
  parses, and applies the data from the request body to make modifications to the
  resource. Following this, the server returns a response that includes a status code
  indicating the success or failure of the operation. Optionally, the response may
  also include the updated resource in the response body.</p>\n<p>Unlike the POST
  method, which is primarily used for creating new resources, PUT is specifically
  designed for updating existing resources on the server. The request body of a PUT
  contains the data necessary for the update, while the URL identifies the specific
  resource to be updated.</p>\n<p>In summary, PUT requests facilitate the transfer
  of data to the server specifically for updating resources, ensuring that changes
  to existing data are accurately processed and reflected.</p>\n<h2>Basic PUT Method</h2>\n<p>To
  send a <code>PUT</code> request to an API in golang, we need to create a <code>http.Request</code>
  object. For <code>POST</code> method, the <code>http</code> package had the <code>Post</code>
  function defined, however for <code>PUT</code> method, there is no separate function.
  The Go philosophy is right now against adding all the method functions. There have
  been a couple of discussions on this on <a href=\"https://github.com/golang/go/issues/22841\">GitHub</a>,
  but it is not been adopted as of 2024.</p>\n<p>So, we need to create a <code>http.Request</code>
  object for <code>PUT</code> method.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in/api/users/5&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">MethodPut</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">DefaultClient</span><span class=\"p\">.</span><span class=\"nx\">Do</span><span
  class=\"p\">(</span><span class=\"nx\">req</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">.</span><span class=\"nx\">Close</span><span class=\"p\">()</span><span
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
  go run main.go\n\n<span class=\"m\">200</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above code sends a <code>PUT</code> request to the <a href=\"https://reqres.in/api/users/5\"><code>https://reqres.in/api/users/5</code></a>
  endpoint. The resource we are trying to update is fetched with the identifier <code>5</code>
  which could probably be <code>id</code> for the user in the database of the server.</p>\n<h2>PUT
  Method with JSON</h2>\n<p>Marshaling and encoding are essential in Go for preparing
  structured data, such as from a struct, into JSON format suitable for HTTP requests
  like PUT. This conversion ensures data integrity and compatibility between Go types
  and JSON representations. It's crucial when updating resources on servers, as APIs
  often require specific data formats for processing updates correctly. Marshaling
  transforms Go structs into JSON bytes, while encoding further prepares them as request
  bodies, facilitating seamless communication with web services. This process ensures
  data consistency and adherence to API specifications, maintaining robust communication
  in distributed systems.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;bytes&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">User</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Name</span><span class=\"w\">   </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;name&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Salary</span><span class=\"w\"> </span><span
  class=\"kt\">int</span><span class=\"w\">    </span><span class=\"s\">`json:&quot;salary&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Age</span><span
  class=\"w\">    </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;age&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">     </span><span class=\"kt\">int</span><span
  class=\"w\">    </span><span class=\"s\">`json:&quot;id,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">type</span><span class=\"w\"> </span><span class=\"nx\">UserResponse</span><span
  class=\"w\"> </span><span class=\"kd\">struct</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Status</span><span class=\"w\">  </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;status&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Message</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;message&quot;`</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">user</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">User</span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">Name</span><span class=\"p\">:</span><span class=\"w\">   </span><span
  class=\"s\">&quot;Alice&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">Salary</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">50000</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">Age</span><span
  class=\"p\">:</span><span class=\"w\">    </span><span class=\"s\">&quot;25&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy.restapiexample.com/api/v1/update/11&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">// marshalling
  process</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">//
  converting Go specific data structure/types to JSON</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">bodyBytes</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Marshal</span><span class=\"p\">(</span><span
  class=\"nx\">user</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">bodyBytes</span><span class=\"p\">))</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">// reading json into a buffer/in-memory</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">body</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">bytes</span><span
  class=\"p\">.</span><span class=\"nx\">NewBuffer</span><span class=\"p\">(</span><span
  class=\"nx\">bodyBytes</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">// post request</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">MethodPut</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">DefaultClient</span><span class=\"p\">.</span><span class=\"nx\">Do</span><span
  class=\"p\">(</span><span class=\"nx\">req</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">StatusCode</span><span
  class=\"w\"> </span><span class=\"o\">==</span><span class=\"w\"> </span><span class=\"mi\">429</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;too many requests&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">return</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">io</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">defer</span><span class=\"w\"> </span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">.</span><span class=\"nx\">Close</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">// unmarshalling
  process</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">//
  converting JSON to Go specific data structure/types</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">userResponse</span><span class=\"w\"> </span><span class=\"nx\">UserResponse</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">userResponse</span><span class=\"p\">);</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">userResponse</span><span class=\"p\">)</span><span
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
  go run json.go                                                                 \n<span
  class=\"o\">{</span><span class=\"s2\">&quot;name&quot;</span>:<span class=\"s2\">&quot;Alice&quot;</span>,<span
  class=\"s2\">&quot;salary&quot;</span>:50000,<span class=\"s2\">&quot;age&quot;</span>:<span
  class=\"s2\">&quot;25&quot;</span><span class=\"o\">}</span>\n<span class=\"m\">200</span>
  \                                                                             \n<span
  class=\"o\">{</span><span class=\"s2\">&quot;status&quot;</span>:<span class=\"s2\">&quot;success&quot;</span>,<span
  class=\"s2\">&quot;data&quot;</span>:<span class=\"o\">[]</span>,<span class=\"s2\">&quot;message&quot;</span>:<span
  class=\"s2\">&quot;Successfully! Record has been updated.&quot;</span><span class=\"o\">}</span>\n<span
  class=\"o\">{</span>success Successfully! Record has been updated.<span class=\"o\">}</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the provided Go code example, the <code>json.Marshal</code> function is used to
  convert a Go struct (<code>User</code>) into a JSON formatted byte slice (<code>[]byte</code>).
  Here's a breakdown of the steps involved:</p>\n<ul>\n<li>\n<p>Struct Definition:
  Define a Go struct with json tags.</p>\n</li>\n<li>\n<p>Marshalling: Use json.Marshal
  to convert the struct into JSON byte slice.</p>\n</li>\n<li>\n<p>Buffer Creation:
  Wrap the JSON byte slice into an in-memory buffer (bytes.Buffer).</p>\n</li>\n<li>\n<p>Request
  Sending: Send a PUT request with the buffer as the request body and set appropriate
  headers.</p>\n</li>\n</ul>\n<p>Let's explore it step by step in detail:</p>\n<p>When
  using the PUT method in Go to update a resource on a server, you often need to send
  data in JSON format as the request body. Here's how you can achieve this using marshaling
  and encoding:</p>\n<ol>\n<li>Define the Struct</li>\n</ol>\n<pre class='wrapper'>\n\n<div
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
  class=\"kd\">type</span><span class=\"w\"> </span><span class=\"nx\">User</span><span
  class=\"w\"> </span><span class=\"kd\">struct</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">Name</span><span class=\"w\">   </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;name&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">Salary</span><span class=\"w\"> </span><span
  class=\"kt\">int</span><span class=\"w\">    </span><span class=\"s\">`json:&quot;salary&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">Age</span><span
  class=\"w\">    </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;age&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nx\">ID</span><span class=\"w\">     </span><span class=\"kt\">int</span><span
  class=\"w\">    </span><span class=\"s\">`json:&quot;id,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Define
  a Go struct (<code>User</code>) that represents the data structure you want to send
  in JSON format. The json tags specify how each field should be serialized into JSON.</p>\n<ol
  start=\"2\">\n<li>Create an Object</li>\n</ol>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nx\">user</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">User</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">Name</span><span
  class=\"p\">:</span><span class=\"w\">   </span><span class=\"s\">&quot;Alice&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">Salary</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">50000</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">Age</span><span class=\"p\">:</span><span
  class=\"w\">    </span><span class=\"s\">&quot;25&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Create
  an instance of the User struct (user) with sample data. This data will be marshaled
  into JSON format to send in the <code>PUT</code> request body.</p>\n<ol start=\"3\">\n<li>Marshal
  the Struct</li>\n</ol>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nx\">bodyBytes</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Marshal</span><span class=\"p\">(</span><span class=\"nx\">user</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Use
  json.Marshal(user) to convert the user struct into a JSON byte slice (bodyBytes).
  This byte slice contains the serialized JSON representation of the User struct.</p>\n<ol
  start=\"4\">\n<li>Create a Buffer</li>\n</ol>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nx\">body</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bytes</span><span class=\"p\">.</span><span
  class=\"nx\">NewBuffer</span><span class=\"p\">(</span><span class=\"nx\">bodyBytes</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Use
  <code>bytes.NewBuffer(bodyBytes)</code> to create an in-memory buffer (<code>body</code>)
  containing the JSON byte slice (<code>bodyBytes</code>). The buffer implements the
  <code>io.Reader</code> interface needed for the PUT request body.</p>\n<ol start=\"5\">\n<li>Create
  a PUT Request</li>\n</ol>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">MethodPut</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">apiURL</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Use
  http.NewRequest to create a new PUT request to the specified URL with the JSON buffer
  (<code>body</code>) as the request body. Set appropriate headers if needed (e.g.,
  Content-Type as application/json).</p>\n<ol start=\"6\">\n<li>Send the Request</li>\n</ol>\n<pre
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
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">DefaultClient</span><span class=\"p\">.</span><span class=\"nx\">Do</span><span
  class=\"p\">(</span><span class=\"nx\">req</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Use
  <a href=\"http://http.DefaultClient.Do\"><code>http.DefaultClient.Do</code></a><code>(req)</code>
  to execute the PUT request and obtain the response. Handle any errors that may occur
  during the request execution.</p>\n<ol start=\"7\">\n<li>Process the Response</li>\n</ol>\n<pre
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
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">io</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Use
  <code>io.ReadAll(resp.Body)</code> to read and store the response body from the
  server. Handle any errors encountered during the reading process.</p>\n<ol start=\"8\">\n<li>Unmarshal
  the Response</li>\n</ol>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">userResponse</span><span
  class=\"w\"> </span><span class=\"nx\">UserResponse</span><span class=\"w\"></span>\n<span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">userResponse</span><span class=\"p\">);</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Use
  <code>json.Unmarshal(respBody, &amp;userResponse)</code> to deserialize the JSON
  response body into a Go struct. This allows you to work with the response data in
  a structured manner.</p>\n<p>The parsing of files and form data is also possible
  with <code>PUT</code> requests, however, that has been covered in the <a href=\"https://meetgor.com/golang-web-post-method\">POST
  Method</a>. Those snippets would be handy in these request method as well.</p>\n<p>I
  have also included some more examples of PUT requests <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/put/\">here</a>.</p>\n<p>That's
  it from the 35th part of the series, all the source code for the examples are linked
  in the GitHub on the <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/put/\">100
  days of Golang</a> repository.</p>\n<p><a href=\"https://github.com/Mr-Destructive/100-days-of-golang\">100-days-of-golang</a></p>\n<h2>Conclusion</h2>\n<p>That's
  it from this post of the series, a post on the PUT method in golang :)</p>\n<p>We
  have covered topics like creating basic PUT requests and marshaling golang types
  into JSON format. Hope you found this article helpful. If you have any queries,
  questions, or feedback, please let me know in the comments or on my social handles.
  Thank you for reading.</p>\n<p>Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/100-days-of-golang/golang-035-put-method.png
long_description: In this section of the series, we will be exploring how to send
  a  A PUT method is a type of request that is used to update or modify an entire
  resource on a server/database. Imagine you have ordered a pizza at a restaurant
  and realized you want to c
now: 2025-05-01 18:11:33.313870
path: blog/posts/2024-06-15-GO-35-PUT-Method.md
prevnext: null
series:
- 100-days-of-golang
slug: golang-web-put-method
status: published
tags:
- go
templateKey: blog-post
title: 'Golang Web: PUT Method'
today: 2025-05-01
---

## Introduction

In this section of the series, we will be exploring how to send a `PUT` HTTP request in golang. We will understand how to send a basic PUT request, create an HTTP request, update a resource on a server, parsing the content from struct to json, headers, etc in the following section of this post.

## What is a PUT Method

A PUT method is a type of request that is used to update or modify an entire resource on a server/database.

Imagine you have ordered a pizza at a restaurant and realized you want to change the toppings after you've already placed the order. With a PUT request, it's like informing the waiter about the changes you want to make to your existing order. You specify the updated toppings or any modifications (the data you send). The waiter then takes this updated information (PUT request) back to the kitchen (the server) to apply the changes to your order.

Let's say you created a order.

```nginx
PUT /api/order/456 HTTP/1.1
Host: example.com
Content-Type: application/json
Content-Length: 123

{
    "userID": 123,
    "orderID": 456,
    "items": [
        {
            "itemID": 789,
            "name": "Pizza",
            "quantity": 2,
            "toppings": ["Mushrooms"]
        }
    ]
}
```

In the context of web development, PUT requests are often used for actions such as:

* Updating existing records or resources
    
* Modifying specific parts of an existing resource
    
* Replacing an entire resource with updated data
    

Here's an example of what the PUT request might look like in this scenario:

```nginx
PUT /api/order/456 HTTP/1.1
Host: example.com
Content-Type: application/json
Content-Length: 155

{
    "userID": 123,
    "orderID": 456,
    "items": [
        {
            "itemID": 789,
            "name": "Pizza",
            "quantity": 2,
            "toppings": ["Mushrooms", "Olives"]
        }
    ]
}
```

In this example:

* The PUT method is used to update the resource identified by `/api/order/456`.
    
* The application/json is the content type of the request.
    
* The 155 is the content length of the request.
    
* The body contains the updated details of the order, including the addition of toppings to the pizza.
    

PUT requests are crucial for maintaining and updating data in applications where accuracy and consistency are paramount, ensuring that resources are kept current and reflect the latest changes made by users or systems

## Why the need of PUT Method

In the world of HTTP requests, we use the PUT method to update or modify an entire resource on a server or database. This method is crucial because the POST method, while convenient for creating new data, is not intended for updating existing resources according to standard conventions. While it's possible to misuse the POST method for updates internally, doing so can lead to confusion and inconsistencies in how requests are understood and processed.

## How PUT Method request works

A [PUT](https://www.rfc-editor.org/rfc/rfc9110#PUT) request is utilized to send data to a server for the purpose of updating a resource. When a client (such as a browser or other APIs) sends a PUT request to the server's API endpoint, it includes data in the request body, typically formatted as JSON, XML, or form data.

The server processes the PUT request by first identifying the resource to be updated, either through the URL or data provided in the request body. It then validates, parses, and applies the data from the request body to make modifications to the resource. Following this, the server returns a response that includes a status code indicating the success or failure of the operation. Optionally, the response may also include the updated resource in the response body.

Unlike the POST method, which is primarily used for creating new resources, PUT is specifically designed for updating existing resources on the server. The request body of a PUT contains the data necessary for the update, while the URL identifies the specific resource to be updated.

In summary, PUT requests facilitate the transfer of data to the server specifically for updating resources, ensuring that changes to existing data are accurately processed and reflected.

## Basic PUT Method

To send a `PUT` request to an API in golang, we need to create a `http.Request` object. For `POST` method, the `http` package had the `Post` function defined, however for `PUT` method, there is no separate function. The Go philosophy is right now against adding all the method functions. There have been a couple of discussions on this on [GitHub](https://github.com/golang/go/issues/22841), but it is not been adopted as of 2024.

So, we need to create a `http.Request` object for `PUT` method.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
	apiURL := "https://reqres.in/api/users/5"

	req, err := http.NewRequest(http.MethodPut, apiURL, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)

	defer resp.Body.Close()
}
```

```bash
$ go run main.go

200
```

The above code sends a `PUT` request to the [`https://reqres.in/api/users/5`](https://reqres.in/api/users/5) endpoint. The resource we are trying to update is fetched with the identifier `5` which could probably be `id` for the user in the database of the server.

## PUT Method with JSON

Marshaling and encoding are essential in Go for preparing structured data, such as from a struct, into JSON format suitable for HTTP requests like PUT. This conversion ensures data integrity and compatibility between Go types and JSON representations. It's crucial when updating resources on servers, as APIs often require specific data formats for processing updates correctly. Marshaling transforms Go structs into JSON bytes, while encoding further prepares them as request bodies, facilitating seamless communication with web services. This process ensures data consistency and adherence to API specifications, maintaining robust communication in distributed systems.

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Age    string `json:"age"`
	ID     int    `json:"id,omitempty"`
}

type UserResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	user := User{
		Name:   "Alice",
		Salary: 50000,
		Age:    "25",
	}
	apiURL := "https://dummy.restapiexample.com/api/v1/update/11"

	// marshalling process
	// converting Go specific data structure/types to JSON
	bodyBytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyBytes))

	// reading json into a buffer/in-memory
	body := bytes.NewBuffer(bodyBytes)

	// post request
	req, err := http.NewRequest(http.MethodPut, apiURL, body)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 429 {
		fmt.Println("too many requests")
		return
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(respBody))
	defer resp.Body.Close()

	// unmarshalling process
	// converting JSON to Go specific data structure/types
	var userResponse UserResponse
	if err := json.Unmarshal(respBody, &userResponse); err != nil {
		panic(err)
	}
	fmt.Println(userResponse)
}
```

```bash
$ go run json.go                                                                 
{"name":"Alice","salary":50000,"age":"25"}
200                                                                              
{"status":"success","data":[],"message":"Successfully! Record has been updated."}
{success Successfully! Record has been updated.}
```

In the provided Go code example, the `json.Marshal` function is used to convert a Go struct (`User`) into a JSON formatted byte slice (`[]byte`). Here's a breakdown of the steps involved:

* Struct Definition: Define a Go struct with json tags.
    
* Marshalling: Use json.Marshal to convert the struct into JSON byte slice.
    
* Buffer Creation: Wrap the JSON byte slice into an in-memory buffer (bytes.Buffer).
    
* Request Sending: Send a PUT request with the buffer as the request body and set appropriate headers.
    

Let's explore it step by step in detail:

When using the PUT method in Go to update a resource on a server, you often need to send data in JSON format as the request body. Here's how you can achieve this using marshaling and encoding:

1. Define the Struct
    

```go
type User struct {
    Name   string `json:"name"`
    Salary int    `json:"salary"`
    Age    string `json:"age"`
    ID     int    `json:"id,omitempty"`
}
```

Define a Go struct (`User`) that represents the data structure you want to send in JSON format. The json tags specify how each field should be serialized into JSON.

2. Create an Object
    

```go
user := User{
    Name:   "Alice",
    Salary: 50000,
    Age:    "25",
}
```

Create an instance of the User struct (user) with sample data. This data will be marshaled into JSON format to send in the `PUT` request body.

3. Marshal the Struct
    

```go
bodyBytes, err := json.Marshal(user)
if err != nil {
    panic(err)
}
```

Use json.Marshal(user) to convert the user struct into a JSON byte slice (bodyBytes). This byte slice contains the serialized JSON representation of the User struct.

4. Create a Buffer
    

```go
body := bytes.NewBuffer(bodyBytes)
```

Use `bytes.NewBuffer(bodyBytes)` to create an in-memory buffer (`body`) containing the JSON byte slice (`bodyBytes`). The buffer implements the `io.Reader` interface needed for the PUT request body.

5. Create a PUT Request
    

```go
req, err := http.NewRequest(http.MethodPut, apiURL, body)
if err != nil {
    panic(err)
}
```

Use http.NewRequest to create a new PUT request to the specified URL with the JSON buffer (`body`) as the request body. Set appropriate headers if needed (e.g., Content-Type as application/json).

6. Send the Request
    

```go
resp, err := http.DefaultClient.Do(req)
if err != nil {
    panic(err)
}
```

Use [`http.DefaultClient.Do`](http://http.DefaultClient.Do)`(req)` to execute the PUT request and obtain the response. Handle any errors that may occur during the request execution.

7. Process the Response
    

```go
respBody, err := io.ReadAll(resp.Body)
if err != nil {
    panic(err)
}
```

Use `io.ReadAll(resp.Body)` to read and store the response body from the server. Handle any errors encountered during the reading process.

8. Unmarshal the Response
    

```go
var userResponse UserResponse
if err := json.Unmarshal(respBody, &userResponse); err != nil {
    panic(err)
}
```

Use `json.Unmarshal(respBody, &userResponse)` to deserialize the JSON response body into a Go struct. This allows you to work with the response data in a structured manner.

The parsing of files and form data is also possible with `PUT` requests, however, that has been covered in the [POST Method](https://meetgor.com/golang-web-post-method). Those snippets would be handy in these request method as well.

I have also included some more examples of PUT requests [here](https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/put/).

That's it from the 35th part of the series, all the source code for the examples are linked in the GitHub on the [100 days of Golang](https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/put/) repository.

[100-days-of-golang](https://github.com/Mr-Destructive/100-days-of-golang)

## Conclusion

That's it from this post of the series, a post on the PUT method in golang :)

We have covered topics like creating basic PUT requests and marshaling golang types into JSON format. Hope you found this article helpful. If you have any queries, questions, or feedback, please let me know in the comments or on my social handles. Thank you for reading.

Happy Coding :)