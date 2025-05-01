---
article_html: "<h2>Introduction</h2>\n<p>In this section of the series, we will be
  exploring how to send a <code>POST</code> HTTP request in golang. We will understand
  how to send a basic POST request, create an HTTP request, and parse json, structs
  into the request body, add headers, etc in the following sections of this post.
  We will understand how to marshal the golang struct/types into JSON format, send
  files in the request, and handle form data with examples of each in this article.
  Let's answer a few questions first.</p>\n<h2>What is a POST request?</h2>\n<p>POST
  method is a type of request that is used to send data to a server(a machine on the
  internet).</p>\n<p>Imagine you are placing an order at a restaurant. With a GET
  request, it would be like asking the waiter, &quot;What kind of pizza do you have?&quot;
  The waiter would respond by telling you the menu options (the information retrieved
  from the server).</p>\n<p>However, a POST request is more like giving your completed
  order to the waiter. You tell them the specific pizza you want, its size, and any
  additional toppings (the data you send). The waiter then takes this information
  (POST request) back to the kitchen (the server) to process it (fulfill your order).</p>\n<p>In
  the world of web development, POST requests are often used for things like:</p>\n<p>Submitting
  forms (e.g., contact forms, login forms) Uploading files (e.g., photos, videos)
  Creating new accounts Sending data to be processed (e.g., online purchases)</p>\n<p>Here's
  an example of what the POST request might look like in this scenario:</p>\n<pre
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
  \       \n<div class='language-bar'>nginx</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">POST</span><span class=\"w\"> </span><span class=\"s\">/api/order</span><span
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
  class=\"w\"> </span><span class=\"mi\">2</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"err\">}</span><span class=\"s\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">            </span><span class=\"kn\">&quot;itemID&quot;:</span><span
  class=\"w\"> </span><span class=\"mi\">999</span><span class=\"s\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">            </span><span class=\"s\">&quot;name&quot;:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Burger&quot;,</span><span class=\"w\"></span>\n<span
  class=\"w\">            </span><span class=\"s\">&quot;quantity&quot;:</span><span
  class=\"w\"> </span><span class=\"mi\">1</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"err\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"s\">]</span><span class=\"w\"></span>\n<span
  class=\"err\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  this example:</p>\n<ul>\n<li>The <code>POST</code> method is used to send data to
  the server.</li>\n<li>The <code>/api/order</code> is the endpoint of the server.</li>\n<li>The
  <code>application/json</code> is the content type of the request.</li>\n<li>The
  <code>123</code> is the content length of the request.</li>\n<li>The <code>{&quot;userID&quot;:
  123, &quot;orderID&quot;: 456, &quot;items&quot;: [{&quot;itemID&quot;: 789, &quot;name&quot;:
  &quot;Pizza&quot;, &quot;quantity&quot;: 2}, {&quot;itemID&quot;: 999, &quot;name&quot;:
  &quot;Burger&quot;, &quot;quantity&quot;: 1}]}</code> is the body of the request.</li>\n</ul>\n<h2>Why
  the need for a POST request?</h2>\n<p>In the world of HTTP requests, we use the
  POST method to securely send data from a client (like a user's browser) to a server.
  This is crucial because the GET method, while convenient for retrieving data, has
  limitations:</p>\n<p>Imagine you are in registering for an event via Google form,
  you type in your details on the webpage like name, email, address, phone number,
  and other personal details. If the website/app was using the <code>GET</code> method
  to send the request to register or do any other authentication/privacy-related requests,
  it could expose the data in the URL itself. It would be something along the lines
  <a href=\"https://form.google.com/register/%3Cform-name%3E-%3Cid%3E/?name=John&amp;phone_number=1234567890\"><code>https://form.google.com/register/&lt;form-name&gt;-&lt;id&gt;/?name=John&amp;phone_number=1234567890</code></a>,
  if a user maliciously sniffs into your network and inspects the URL, your data will
  be exposed. That is the reason we need <code>POST</code> a method.</p>\n<h2>How
  a POST method works?</h2>\n<p>A <a href=\"https://www.rfc-editor.org/rfc/rfc9110#POST\">POST</a>
  request is used to send data to a server to create or update(there is a separate
  method for updating) a resource. The client(browser/other APIs) sends a POST request
  to the server's API endpoint with the data in the request body. This data can be
  in formats like JSON, XML, or form data. The server processes the POST request,
  validates and parses the data in the request body, makes any changes or creates
  resources based on that data, and returns a response. The response would contain
  a status code indicating the success or failure of the operation and may contain
  the newly created or updated resource in the response body. The client must check
  the response status code to verify the outcome and process the response accordingly.
  Unlike GET, POST can create new resources on the server. The body of a POST contains
  the data for creation while the URL identifies the resource to be created. Overall,
  POST transfers data to the server for processing, creation or updating of resources.</p>\n<p>The
  status code is usually <code>201</code> indicating the resource is successfully
  created or <code>200</code> for just indicating success.</p>\n<p>Some common steps
  for creating and sending a POST request as a developer include:</p>\n<ul>\n<li>\n<p>Defining
  the API endpoint</p>\n</li>\n<li>\n<p>Clarifying the data format (json, language
  native objects, xml , text, form-data, etc)</p>\n</li>\n<li>\n<p>Converting / Marshalling
  the data</p>\n</li>\n<li>\n<p>Attaching header for <code>Content-Type</code> as
  key and value as the format of the data type (e.g. <code>application/json</code>
  for json)</p>\n</li>\n<li>\n<p>Sending the request</p>\n</li>\n</ul>\n<p>The above
  steps are general for creating and sending a POST request, they are not specific
  to Golang. For golang specific steps, we need to dive a bit deeper, let's get started.</p>\n<h2>Basic
  POST method in Golang</h2>\n<p>To send a POST request in golang, we need to use
  the <code>http</code> package. The <code>http</code> package has the <code>Post</code>
  method, which takes in 3 parameters, namely the URL, the Content-Type, and the Body.
  The body can be <code>nil</code> if the URL endpoint doesn't necessarily require
  a body. The <code>Content-Type</code> is the string, since we are just touching
  on how the Post request is constructed, we will see what the <code>Content-Type</code>
  string value should be in the later sections.</p>\n<blockquote>\n<p><a href=\"http://http.Post\">http.Post</a>(URL,
  Content-Type, Body)</p>\n</blockquote>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in/api/users&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">// POST request</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Post</span><span class=\"p\">(</span><span
  class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"kc\">nil</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"c1\">// ideally the Content-Type header should
  be set to the relevant format</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// resp, err := http.Post(apiURL, &quot;application/json&quot;, nil)</span><span
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
  class=\"w\">    </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
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
  go run main.go\n\n<span class=\"m\">201</span>\n<span class=\"p\">&amp;</span><span
  class=\"o\">{</span>\n    <span class=\"m\">201</span> Created\n    <span class=\"m\">201</span>\n
  \   HTTP/2.0\n    <span class=\"m\">2</span>\n    <span class=\"m\">0</span>\n    map<span
  class=\"o\">[</span>\n        Access-Control-Allow-Origin:<span class=\"o\">[</span>*<span
  class=\"o\">]</span>\n        Cf-Cache-Status:<span class=\"o\">[</span>DYNAMIC<span
  class=\"o\">]</span>\n        Cf-Ray:<span class=\"o\">[</span>861cd9aec8223e4b-BOM<span
  class=\"o\">]</span>\n        Content-Length:<span class=\"o\">[</span><span class=\"m\">50</span><span
  class=\"o\">]</span>\n        Content-Type:<span class=\"o\">[</span>application/json<span
  class=\"p\">;</span> <span class=\"nv\">charset</span><span class=\"o\">=</span>utf-8<span
  class=\"o\">]</span>\n        Date:<span class=\"o\">[</span>Sat, <span class=\"m\">09</span>
  Mar <span class=\"m\">2024</span> <span class=\"m\">17</span>:40:28 GMT<span class=\"o\">]</span>\n
  \       Server:<span class=\"o\">[</span>cloudflare<span class=\"o\">]</span>\n
  \       ...\n        ...\n        ...\n        X-Powered-By:<span class=\"o\">[</span>Express<span
  class=\"o\">]</span>\n    <span class=\"o\">]</span>\n    <span class=\"o\">{</span>0xc00017c180<span
  class=\"o\">}</span>\n    <span class=\"m\">50</span>\n    <span class=\"o\">[]</span>\n
  \   <span class=\"nb\">false</span>\n    <span class=\"nb\">false</span>\n    map<span
  class=\"o\">[]</span>\n    0xc000156000\n    0xc00012a420\n<span class=\"o\">}</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above code is sending the <code>POST</code> request to the <a href=\"https://reqres.in/api/users\"><code>https://reqres.in/api/users</code></a>
  endpoint with an empty body and no specific format for <code>Content-Type</code>
  header. The response is according to the <a href=\"https://pkg.go.dev/net/http#Response\">Response</a>
  structure. We can see we got <code>201</code> status, which indicates the server
  received the POST request successfully, the API is a dummy api, so we don't care
  about the data we are processing, we are just using the API as a placeholder for
  sending the POST request.</p>\n<p>We can use <code>map[string]interface{}</code>
  it to pass the data in the request body. The <code>json.Marshal</code> method is
  used to convert the map into JSON format. We will look into the details shortly
  in the next few examples.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in/api/users&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">bodyMap</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"kt\">string</span><span class=\"p\">]</span><span
  class=\"kd\">interface</span><span class=\"p\">{}{</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"s\">&quot;name&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;morpheus&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"s\">&quot;job&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s\">&quot;leader&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">    </span><span
  class=\"nx\">requestBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Marshal</span><span class=\"p\">(</span><span class=\"nx\">bodyMap</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">body</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bytes</span><span class=\"p\">.</span><span
  class=\"nx\">NewBuffer</span><span class=\"p\">(</span><span class=\"nx\">requestBody</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">    </span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">Post</span><span class=\"p\">(</span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;application/json&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">defer</span><span class=\"w\"> </span><span
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
  go run main.go\n\n<span class=\"m\">201</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above code sends the <code>POST</code> request to the <a href=\"https://reqres.in/api/users\"><code>https://reqres.in/api/users</code></a>
  endpoint with the data in the request body in JSON format.</p>\n<h2>Creating a POST
  request in Golang</h2>\n<p>We can construct the POST request with the <a href=\"https://pkg.go.dev/net/http#NewRequest\">NewRequest</a>
  method. The method takes in 3 parameters, namely the <code>method</code> (e.g. <code>POST</code>,
  <code>GET</code>), the <code>URL</code> and the <code>body</code> (if there is any).
  We can then add extra information to the headers or the Request object after constructing
  the basic HTTP <a href=\"https://pkg.go.dev/net/http#Request\">Request</a> object.</p>\n<pre
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
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in/api/users&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">MethodPost</span><span
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
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"c1\">//fmt.Println(resp)</span><span class=\"w\"></span>\n<span
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
  go run main.go\n\n<span class=\"m\">201</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we have created an HTTP Request as the <code>POST</code> method,
  with <a href=\"https://reqres.in/api/users\"><code>https://reqres.in/api/users</code></a>
  as the URL, and no body. This constructs an HTTP Request object, which can be sent
  as the parameter to the <a href=\"http://http.DefaultClient.Do\">http.DefaultClient.Do</a>
  method, which is the default client for the request we sent in the earlier examples
  as <code>http.Get</code> or <a href=\"http://http.Post\"><code>http.Post</code></a>
  methods. We can implement a custom client as well, and then apply <code>Do</code>
  the method with the request parameters. The <code>Do</code> method returns the <code>Request</code>
  object or the <code>error</code> if any.</p>\n<p>More on the customizing Client
  will be explained in a separate post in the series.</p>\n<p>The response is also
  in the same format as the <a href=\"https://pkg.go.dev/net/http#Response\">Response</a>
  structure that we have seen earlier. This section of the series aims to construct
  a post request, and not to parse the response, we have already understood the parsing
  of the response in the <a href=\"https://www.meetgor.com/golang-web-get-method/#?:~:text=Parsing%20the%20JSON%20body%20with%20structs\">Get
  method</a> section of the series.</p>\n<h3>Parsing objects to JSON for POST method
  request</h3>\n<p>We might have a golang object that we want to send as a body to
  an API in the POST request, for that we need to convert the golang struct object
  to JSON. We can do this by using the <a href=\"https://pkg.go.dev/encoding/json#Marshal\">Marshal</a>
  or the <a href=\"https://pkg.go.dev/encoding/json#Encoder.Encode\">Encode</a> method
  for serialization of the golang struct object to JSON.</p>\n<h4>Using Marshal method</h4>\n<p>Marshaling
  is the process of converting data from a data structure into a format suitable for
  transmission over a network or for storage. It's commonly used to convert native
  objects in a programming language into a serialized format, typically a byte stream,
  that can be transmitted or stored efficiently. You might get a question here, what
  is the difference between <code>Marshalling</code> and <code>Serialization</code>?
  Well, Serialization, is a broader term that encompasses marshalling. It refers to
  the process of converting an object or data structure into a format that can be
  stored or transmitted and later reconstructed into the original object. Serialization
  may involve converting data into byte streams, XML, JSON, or other formats. So,
  in summary, marshaling specifically deals with converting native objects into a
  format suitable for transmission, while serialization encompasses the broader process
  of preparing data for storage or transmission.</p>\n<p>The <code>json</code> package
  has the <a href=\"https://pkg.go.dev/encoding/json#Marshal\">Marshal</a> method
  that converts the golang object into JSON. The <code>Marshal</code> method takes
  in a parameter as the struct object with type <code>any</code> and returns a byte
  slice <code>[]byte</code> and error (if any).</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;bytes&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">User</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Name</span><span class=\"w\">   </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;name&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Salary</span><span class=\"w\"> </span><span
  class=\"kt\">int</span><span class=\"w\">    </span><span class=\"s\">`json:&quot;salary&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Age</span><span
  class=\"w\">    </span><span class=\"kt\">int</span><span class=\"w\">    </span><span
  class=\"s\">`json:&quot;age&quot;`</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">user</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">User</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">Name</span><span
  class=\"p\">:</span><span class=\"w\">   </span><span class=\"s\">&quot;Alice&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">Salary</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">50000</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">Age</span><span class=\"p\">:</span><span
  class=\"w\">    </span><span class=\"mi\">25</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">apiURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://dummy.restapiexample.com/api/v1/create&quot;</span><span
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
  class=\"w\">\t</span><span class=\"nx\">resp</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Post</span><span class=\"p\">(</span><span
  class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;application/json&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">body</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  go run main.go\n\n<span class=\"o\">{</span><span class=\"s2\">&quot;name&quot;</span>:<span
  class=\"s2\">&quot;Alice&quot;</span>,<span class=\"s2\">&quot;salary&quot;</span>:50000,<span
  class=\"s2\">&quot;age&quot;</span>:25<span class=\"o\">}</span>\n<span class=\"m\">200</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we have created a struct <code>User</code> with fields <code>Name</code>,
  <code>Salary</code>, and <code>Age</code>, the json tags will help label each key
  in JSON with the tag for the respective fields in the struct. We create an object
  <code>user</code> of a type <code>User</code> with the values as <code>Alice</code>,
  <code>50000</code>, and <code>25</code> respectively.</p>\n<p>We call the <code>json.Marshal</code>
  method with the parameter <code>user</code> that represents the struct object <code>User</code>,
  the method returns a slice of bytes, or an error either or both could be nil. If
  we try to see the stringified representation of the byte slice, we can see something
  like <code>{&quot;name&quot;:&quot;Alice&quot;,&quot;salary&quot;:50000,&quot;age&quot;:25}</code>
  which is a JSON string for the user struct. We can't parse the byte slice as the
  body in the POST request, we need the <code>io.Reader</code> object, so we can load
  the byte slice <code>bodyBytes</code> into a buffer and parse that as a body for
  the POST request.</p>\n<p>We then send a <code>POST</code> request to the endpoint
  <a href=\"https://dummy.restapiexample.com/api/v1/create\"><code>https://dummy.restapiexample.com/api/v1/create</code></a>
  with the content type as <code>application/json</code> and with the body as <code>body</code>
  which was a <code>io.Reader</code> object as an in-memory buffer.</p>\n<p>In brief,
  we can summarize the marshaling of the golang object into JSON with <code>Marshal</code>
  function as the following steps:</p>\n<ul>\n<li>\n<p>Defining the structure as per
  the request body</p>\n</li>\n<li>\n<p>Creating the struct object for parsing the
  data as body to the request</p>\n</li>\n<li>\n<p>Calling the <code>json.Marshal</code>
  function to convert the object to JSON (parameter as the struct object <code>any</code>
  type)</p>\n</li>\n<li>\n<p>Loading the byte slice into a buffer with <code>bytes.NewBuffer()</code></p>\n</li>\n<li>\n<p>Sending
  the POST request to the endpoint with the body as the <code>io.Reader</code> object
  and content type as <code>application/json</code></p>\n</li>\n</ul>\n<h4>Using Encode
  method</h4>\n<p>We can even use the <a href=\"https://pkg.go.dev/encoding/json#Encoder.Encode\">Encoder.Encode</a>
  method to parse the golang struct object to JSON. Firstly, we should have the struct
  defined as per the request body that the particular API takes, we can make use of
  the json tags, omitempty, omit(-) options to make the marshaling process work accordingly.
  We can then create the object of that particular struct with the data we require
  to be created as a resource with the POST request on that API service.</p>\n<p>Thereafter
  we can create an empty buffer object with <a href=\"https://pkg.go.dev/bytes#Buffer\">bytes.Buffer</a>,
  this buffer object would be used to initialize the <a href=\"https://pkg.go.dev/encoding/json#Encoder\">Encoder</a>
  object with the <a href=\"https://pkg.go.dev/encoding/json#NewEncoder\">NewEncoder</a>
  method. This would give access to the <a href=\"https://pkg.go.dev/encoding/json#Encoder.Encode\">Encode</a>
  method, which is used to take in the struct object (<code>any</code> type) and this
  will populate the buffer we initialized with the <code>NewEncoder</code> method.</p>\n<p>Later
  we can access that buffer to parse it to the Post request as the body. Let's understand
  it better with an example.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">User</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Name</span><span class=\"w\">   </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Salary</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Age</span><span class=\"w\">    </span><span
  class=\"kt\">int</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">user</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">User</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">Name</span><span
  class=\"p\">:</span><span class=\"w\">   </span><span class=\"s\">&quot;Alice&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">Salary</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">50000</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">Age</span><span class=\"p\">:</span><span
  class=\"w\">    </span><span class=\"mi\">25</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">apiURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://dummy.restapiexample.com/api/v1/create&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">bodyBuffer</span><span class=\"w\"> </span><span
  class=\"nx\">bytes</span><span class=\"p\">.</span><span class=\"nx\">Buffer</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">encoder</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">NewEncoder</span><span class=\"p\">(</span><span
  class=\"o\">&amp;</span><span class=\"nx\">bodyBuffer</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">encoder</span><span
  class=\"p\">.</span><span class=\"nx\">Encode</span><span class=\"p\">(</span><span
  class=\"nx\">user</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">Post</span><span class=\"p\">(</span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;application/json&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">bodyBuffer</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Over
  here, we have created a struct <code>User</code> with fields <code>Name</code>,
  <code>Salary</code>, and <code>Age</code>, we initialize the <code>user</code> as
  the object of the <code>User</code> struct. Then we create a buffer <code>bodyBuffer</code>
  of type <code>bytes.Buffer</code> this is the actual buffer that we will send as
  the body. Further, we initialize the <code>Encoder</code> object as <code>encoder</code>
  with the <code>json.NewEncoder</code> method by parsing the reference of <code>bodyBuffer</code>
  as the parameter. Since <code>bytes.Buffer</code> implements the <code>io.Writer</code>
  interface, we can pass the <code>bodyBuffer</code> to the <code>NewEncoder</code>
  method. This will create the <code>Encoder</code> object which in turn will give
  us access to the <code>Encode</code> method, where we will parse the struct instance
  and it will populate the buffer with which we initialized the <code>Encoder</code>
  object earlier.</p>\n<p>Now, we have the <code>encode</code> object, this gives
  us the access to <code>Encode</code> method, we call the <code>Encode</code> method
  with the parameter of <code>user</code> which is a User struct instance/object.
  The Encode method will populate the <code>bodyBuffer</code> object or it will result
  in an error if anything goes wrong (the data is incorrectly parsed or is not in
  the required format).</p>\n<p>We can call the <code>Post</code> method with the
  initialized URL, the <code>Content-Type</code> as <code>application/json</code>
  since we have converted the struct instance to JSON object, and the body as the
  reference to the buffer as <code>&amp;bodyBuffer</code></p>\n<p>So, the steps for
  parsing struct instances into JSON objects with the <code>Encoder.Encode</code>
  method is as follows:</p>\n<ul>\n<li>\n<p>Defining the structure as per the request
  body</p>\n</li>\n<li>\n<p>Creating the struct object for parsing the data as body
  to the request</p>\n</li>\n<li>\n<p>Creating an empty <code>bytes.Buffer</code>
  object as an in-memory buffer</p>\n</li>\n<li>\n<p>Initializing the <code>Encoder</code>
  object with <code>NewEncoder</code> method by parsing the reference of <code>bodyBuffer</code>
  as the parameter</p>\n</li>\n<li>\n<p>Calling the <code>Encode</code> method with
  the parameter of struct instance/object</p>\n</li>\n<li>\n<p>Sending the POST request
  to the endpoint with the content type as <code>application/json</code> and body
  as the reference to the buffer</p>\n</li>\n</ul>\n<p>The results are the same as
  the above example just the way we have parsed the struct instance to JSON object
  is different.</p>\n<h3>Parsing JSON to POST request</h3>\n<p>We have seen how we
  can parse golang struct instances to JSON and then send the post request, but what
  if we had the JSON string already with us, and we want to send the request? Well,
  that's much easier, right? We already have parsed the JSON string to the Post request
  by loading the slice of bytes into a buffer, so we just need to convert the string
  to a slice of bytes which is quite an easy task, and then load that byte slice to
  the buffer.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// dummy api</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy.restapiexample.com/api/v1/create&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">// json data</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">data</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`{</span>\n<span
  class=\"s\">        &quot;name&quot;: &quot;Alice&quot;,</span>\n<span class=\"s\">
  \       &quot;job&quot;: &quot;Teacher&quot;</span>\n<span class=\"s\">    }`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">body</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">bytes</span><span
  class=\"p\">.</span><span class=\"nx\">NewBuffer</span><span class=\"p\">([]</span><span
  class=\"nb\">byte</span><span class=\"p\">(</span><span class=\"nx\">data</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"c1\">// POST request</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">Post</span><span class=\"p\">(</span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;application/json&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In the
  example above, we already have a JSON string <code>data</code> with keys as <code>name</code>
  and <code>job</code> but it is not JSON, it is a stringified JSON. We can convert
  the stringified JSON to a slice of bytes using the <code>[]byte</code> function.
  Further, we have used the <code>bytes.NewBuffer</code> method to load the byte slice
  into an <code>io.Reader</code> object. This object returned by the <code>bytes.NewBuffer</code>
  will serve as the body for the POST request.</p>\n<h3>Parsing JSON to objects in
  Golang from POST method response</h3>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nx\">Status</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;status&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Data</span><span class=\"w\">   </span><span
  class=\"nx\">User</span><span class=\"w\">   </span><span class=\"s\">`json:&quot;data&quot;`</span><span
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
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy.restapiexample.com/api/v1/create&quot;</span><span
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
  class=\"w\">\t</span><span class=\"nx\">resp</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Post</span><span class=\"p\">(</span><span
  class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;application/json&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">body</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">// Read response body</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">respBody</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"c1\">// unmarshalling process</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// converting JSON to Go specific data structure/types</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">userResponse</span><span class=\"w\"> </span><span
  class=\"nx\">UserResponse</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
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
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">userResponse</span><span class=\"p\">.</span><span class=\"nx\">Data</span><span
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
  \       \n<div class='language-bar'>nginx</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">{success</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"kn\">Alice</span><span class=\"w\"> </span><span class=\"mi\">50000</span><span
  class=\"w\"> </span><span class=\"mi\">25</span><span class=\"w\"> </span><span
  class=\"mi\">3239</span><span class=\"err\">}}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"kn\">Alice</span><span class=\"w\"> </span><span
  class=\"mi\">50000</span><span class=\"w\"> </span><span class=\"mi\">25</span><span
  class=\"w\"> </span><span class=\"mi\">577</span><span class=\"err\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The above example is a POST request
  with a struct instance being loaded as a JSON string and then sent as a buffer to
  the API endpoint, it also reads the response body with a specific structure <code>UserResponse</code>
  and unmarshalled the <code>resp.Body</code> from the <code>io.Reader</code> as <code>respBody</code>
  and then loads into <code>userResponse</code> object. This example gives an entire
  process of what we have understood in the JSON data parsing for a POST request.</p>\n<h3>Sending
  Form data in a POST request</h3>\n<p>We can also send data to a POST request in
  the form of a form, the form which we use in the HTML. Golang has a <code>net/url</code>
  package to parse the form data. The form data is sent in the <code>application/x-www-form-urlencoded</code>
  format.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/url&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">ResponseLogin</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Token</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;token&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// dummy api</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in/api/login&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">// Define form
  data</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">formData</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">url</span><span
  class=\"p\">.</span><span class=\"nx\">Values</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">formData</span><span
  class=\"p\">.</span><span class=\"nx\">Set</span><span class=\"p\">(</span><span
  class=\"s\">&quot;email&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;eve.holt@reqres.in&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">formData</span><span
  class=\"p\">.</span><span class=\"nx\">Set</span><span class=\"p\">(</span><span
  class=\"s\">&quot;password&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;cityslicka&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">// Encode the
  form data</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">formData</span><span class=\"p\">.</span><span class=\"nx\">Encode</span><span
  class=\"p\">())</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">reqBody</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">strings</span><span class=\"p\">.</span><span
  class=\"nx\">NewReader</span><span class=\"p\">(</span><span class=\"nx\">formData</span><span
  class=\"p\">.</span><span class=\"nx\">Encode</span><span class=\"p\">())</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">reqBody</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">// Make a POST request with form data</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Post</span><span class=\"p\">(</span><span
  class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;application/x-www-form-urlencoded&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">reqBody</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">// Print response status code</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Status Code:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">StatusCode</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"c1\">// Read response body</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">io</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">token</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">ResponseLogin</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">token</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">token</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  go run main.go\n\n<span class=\"nv\">email</span><span class=\"o\">=</span>eve.holt%40reqres.in<span
  class=\"p\">&amp;</span><span class=\"nv\">password</span><span class=\"o\">=</span>cityslicka\n<span
  class=\"p\">&amp;</span><span class=\"o\">{</span><span class=\"nv\">email</span><span
  class=\"o\">=</span>eve.holt%40reqres.in<span class=\"p\">&amp;</span><span class=\"nv\">password</span><span
  class=\"o\">=</span>cityslicka <span class=\"m\">0</span> -1<span class=\"o\">}</span>\nStatus
  Code: <span class=\"m\">200</span>\n<span class=\"o\">{</span>QpwL5tke4Pnpja7X4<span
  class=\"o\">}</span>\n</pre></div>\n\n</pre>\n\n<p>In the above example, we set
  a <code>formData</code> with the values of <code>email</code> and <code>password</code>
  which are <code>url.Values</code> object. The <code>url.Values</code> the object
  is used to store the key-value pairs of the form data. The <code>formData</code>
  is encoded with the <code>url.Encode</code> method, We load the encoded string to
  a <code>buffer</code> with <code>strings.NewReader</code> which implements the <code>io.Reader</code>
  interface, so that way we can pass that object as the body to the post request.</p>\n<p>We
  send the <code>POST</code> request to the endpoint <a href=\"https://reqres.in/api/login\"><code>https://reqres.in/api/login</code></a>
  with the content type as <code>application/x-www-form-urlencoded</code> and with
  the body as <code>reqBody</code> which implements the <code>io.Reader</code> interface
  as an in-memory buffer. The response from the request is read into the buffer with
  <code>io.ReadAll</code> method and we can <code>Unmarshal</code> the stream of bytes
  as a buffer into the <code>ResponseLogin</code> struct object.</p>\n<p>The output
  shows the <code>formData</code> as encoded string <code>email=eve.holt%</code><a
  href=\"http://40reqres.in\"><code>40reqres.in</code></a><code>&amp;password=cityslicka</code>
  as <code>@</code> is encoded to <code>%40</code>, then we wrap the <code>formData</code>
  in a <code>strings.NewReader</code> object which is a buffer that implements <code>io.Reader</code>
  interface, hence we can see the result as the object. The status code for the request
  is <code>200</code> indicating the server received the <code>form-data</code> in
  the body and upon unmarshalling, we get the token as a response to the POST request
  which was a dummy login API.</p>\n<p>This way we have parsed the form-data to the
  body of a POST request.</p>\n<h3>Sending File in a POST request</h3>\n<p>We have
  covered, parsing text, JSON, and form data, and now we need to move into sending
  files in a POST request. We can use the <code>multipart</code> package to parse
  files into the request body and set appropriate headers for reading the file from
  the API services.</p>\n<p>We first read the file contents <a href=\"http://os.Open\"><code>os.Open</code></a>
  which returns a reference to the <code>file</code> object or an error. We create
  an empty <code>bytes.Buffer</code> object as <code>body</code> which will be populated
  later. The <a href=\"https://pkg.go.dev/mime/multipart#NewWriter\">multipart.NewWriter</a>
  method takes in the <code>io.Writer</code> object which will be the <code>body</code>
  as it is an <code>bytes.Buffer</code> object that implements the <code>io.Writer</code>
  interface. This will initialize the <a href=\"https://pkg.go.dev/mime/multipart#Writer\">Writer</a>
  object in the <code>multipart</code> package.</p>\n<p>We create a <code>form-field</code>
  in the <code>Writer</code> object with the <a href=\"https://pkg.go.dev/mime/multipart#Writer.CreateFormFile\">CreateFormFile</a>
  method, which takes in the <code>fieldName</code> as the name of the field, and
  the <code>fileName</code> as the name of the file which will be read later in the
  multipart form. The method returns either the part or the error. The <code>part</code>
  is an object that implements the <code>io.Writer</code> interface.</p>\n<p>Since
  we have stored the file contents in the <code>file</code> object, we copy the contents
  into the <code>form-field</code> with the <a href=\"https://pkg.go.dev/io#Copy\">Copy</a>
  method. Since the <code>part</code> return from the <code>CreateFormFile</code>
  was implementing the <code>io.Writer</code> interface, we can use it to Copy the
  contents from source to destination. The source is the <code>io.Reader</code> object
  and the destination is the <code>io.Writer</code> object, the destination for the
  <code>Copy</code> method is the first parameter, the source is the second parameter.</p>\n<p>This
  Copy method will populate the buffer initialized earlier in the <code>NewWriter</code>
  method. This will give us a buffer that has the file contents in it. We can pass
  this buffer to the POST request with the <code>body</code> parameter. We also need
  to make sure we close the <code>Writer</code> object after copying the contents
  of the file. We can extract the type of file which will serve as the <code>Content-Type</code>
  of the request.</p>\n<p>Let's clear the explanation with an example.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;bytes&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;mime/multipart&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">ResponseFile</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Files</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"kt\">string</span><span class=\"p\">]</span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;files&quot;`</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">apiURL</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;http://postman-echo.com/post&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fileName</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;sample.csv&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">file</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"nx\">fileName</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">file</span><span
  class=\"p\">.</span><span class=\"nx\">Close</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">body</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">bytes</span><span class=\"p\">.</span><span class=\"nx\">Buffer</span><span
  class=\"p\">{}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">writer</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">multipart</span><span class=\"p\">.</span><span
  class=\"nx\">NewWriter</span><span class=\"p\">(</span><span class=\"nx\">body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">part</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">writer</span><span class=\"p\">.</span><span
  class=\"nx\">CreateFormFile</span><span class=\"p\">(</span><span class=\"s\">&quot;csvFile&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">fileName</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">Copy</span><span class=\"p\">(</span><span
  class=\"nx\">part</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">file</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">    </span><span
  class=\"nx\">contentType</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">writer</span><span class=\"p\">.</span><span
  class=\"nx\">FormDataContentType</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">contentType</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">writer</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">Post</span><span class=\"p\">(</span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">contentType</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Status
  Code:&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">StatusCode</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">respBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">token</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">ResponseFile</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">token</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">token</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">token</span><span
  class=\"p\">.</span><span class=\"nx\">Files</span><span class=\"p\">[</span><span
  class=\"nx\">fileName</span><span class=\"p\">])</span><span class=\"w\"></span>\n<span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>multipart/form-data<span
  class=\"p\">;</span> <span class=\"nv\">boundary</span><span class=\"o\">=</span>7e0eacfff890be395eba19c70415c908124b503a56f23ebeec0ab3c665ca\n\n\n--619671ea2c0aa47ca6664a7cda422169d73f3b8a089c659203f5413d03de\nContent-Disposition:
  form-data<span class=\"p\">;</span> <span class=\"nv\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;csvFile&quot;</span><span class=\"p\">;</span> <span class=\"nv\">filename</span><span
  class=\"o\">=</span><span class=\"s2\">&quot;sample.csv&quot;</span>\nContent-Type:
  application/octet-stream\n\nUser,City,Age,Country\nAlex Smith,Los Angeles,20,USA\nJohn
  Doe,New York,30,USA\nJane Smith,Paris,25,France\nBob Johnson,London,40,UK\n\n--619671ea2c0aa47ca6664a7cda422169d73f3b8a089c659203f5413d03de--\n\n\n\nStatus
  Code: <span class=\"m\">200</span>\n\n<span class=\"o\">{</span>map<span class=\"o\">[</span>sample.csv:data:application/octet-stream<span
  class=\"p\">;</span>base64,VXNlcixDaXR5LEFnZSxDb3VudHJ5CkFsZXggU21pdGgsTG9zIEFuZ2VsZXMsMjAsVVNBCkpvaG4gRG9lLE5ldyBZb3JrLDMwLFVTQQpKYW5lIFNtaXRoLFBhmlzLDI1LEZyYW5jZQpCb2IgSm9obnNvbixMb25kb24sNDAsVUsK<span
  class=\"o\">]}</span>\n\ndata:application/octet-stream<span class=\"p\">;</span>base64,VXNlcixDaXR5LEFnZSxDb3VudHJ5CkFsZXggU21pdGgsTG9zIEFuZ2VsZXMsMjAsVVNBCkpvaG4gRG9lLE5ldyBZb3JrLDMwLFVTQQpKYW5lIFNtaXRoLFBhmlzLDI1LEZyYW5jZQpCb2IgSm9obnNvbixMb25kb24sNDAsVUsK\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we first read the file <code>sample.csv</code> into the <code>file</code>
  object with <a href=\"http://os.Open\"><code>os.Open</code></a> method, this will
  return a reference to the file object or return an error if any arises while opening
  the file.</p>\n<p>Then we create an empty buffer <code>bytes.Buffer</code> object
  which will serve as the body of the post request later as it will get populated
  with the file contents in the form of <code>multipart/form-data</code>.</p>\n<p>We
  initialize the <code>Writer</code> object with <code>multipart.NewWriter</code>
  method which takes in the empty buffer as the parameter, we parse the <code>body</code>
  as the parameter. The method will return a reference to the <code>multipart.Writer</code>
  object.</p>\n<p>With the <code>Writer</code> object we access the <code>CreateFormFile</code>
  method which takes in the <code>fieldName</code> as the name of the field, and the
  <code>fileName</code> as the name of the file. The method will return either the
  part or an error. The <code>part</code> in this case, is the reference to the <code>io.Writer</code>
  object that will be used to write the contents from the uploaded file.</p>\n<p>Then,
  we can use the <code>io.Copy</code> method to copy the contents from the <code>io.Reader</code>
  object to the <code>io.Writer</code> object. The source is the <code>io.Reader</code>
  object and the destination is the <code>io.Writer</code> object. The first parameter
  is however the destination and the second parameter is the source. In the example,
  we call <code>io.Copy(part, file)</code> which will copy the contents of <code>file</code>
  to the <code>part</code> buffer.</p>\n<p>We get the <code>Content-Type</code> by
  calling the <a href=\"https://pkg.go.dev/mime/multipart#Writer.FormDataContentType\">Writer.FormDataContentType</a>
  method. This returns us <code>multipart/form-data; boundary=7e0eacfff890be395eba19c70415c908124b503a56f23ebeec0ab3c665ca</code>
  which will serve the <code>Content-Type</code> for the Post request.</p>\n<p>We
  need to make sure we close the <code>Writer</code> object with the <code>Close</code>
  method.</p>\n<p>We just print the <code>body.String()</code> to get a look at what
  the actual body looks like, we can see there is a form for the file as a <code>form-data</code>
  with keys like <code>Content-Type</code>, <code>Content-Disposition</code>, etc.
  The file has the <code>Content-Type</code> as <code>application/octet-stream</code>
  and the actual content is rendered in the output.</p>\n<p>The dummy API responds
  with a 200 status code and also sends the JSON data with the name of the file as
  the key and the value as the <code>base64</code> encoded value of the file contents.
  This indicates that we were able to upload the file to the server API using a POST
  request. Well done!</p>\n<p>I have also included some more examples of POST requests
  with files <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/post/file_2.go\">here</a>
  which extends the above example by taking the encoded values and decoding to get
  the actual contents of the file back.</p>\n<h2>Best Practices for POST method</h2>\n<p>Here
  are some of the best practices for the POST method which are followed to make sure
  you consume or create the POST request in the most secure, efficient, and graceful
  way.</p>\n<h3>Always Close the Response Body</h3>\n<p>Ensure that you close the
  response body after reading from it. Use <code>defer response.Body.Close()</code>
  to automatically close the body when the surrounding function returns. This is crucial
  for releasing associated resources like network connections or file descriptors.
  Failure to close the response body can lead to memory leaks, particularly with a
  large volume of requests. Properly closing the body prevents resource exhaustion
  and maintains efficient memory usage.</p>\n<h3>Client Customization</h3>\n<p>Utilize
  the <a href=\"https://pkg.go.dev/net/http#Client\">Client</a> struct to customize
  the HTTP client behavior. By using a custom client, you can set timeouts, headers,
  user agents, and other configurations without modifying the <code>DefaultClient</code>
  provided by the <code>http</code> package. This approach allows for flexibility
  and avoids repetitive adjustments to the client configuration for each request.</p>\n<h3>Set
  Content-Type Appropriately</h3>\n<p>Ensure that you set the <code>Content-Type</code>
  header according to the request payload. Correctly specifying the Content-Type is
  crucial for the server to interpret the request payload correctly. Failing to set
  the Content-Type header accurately may result in the server rejecting the request.
  Always verify and match the Content-Type header with the content being sent in the
  POST request to ensure smooth communication with the server.</p>\n<h2>Reference</h2>\n<ul>\n<li>\n<p><a
  href=\"https://www.postman.com/postman/workspace/postman-answers/documentation/13455110-00378d5c-5b08-4813-98da-bc47a2e6021d\">Postman
  POST API</a> (For POST request with file upload)</p>\n</li>\n<li>\n<p><a href=\"https://pkg.go.dev/net/http\">Golang
  net/http Package</a></p>\n</li>\n</ul>\n<h2>Conclusion</h2>\n<p>That's it from this
  post of the series, a post on the POST method in golang :)</p>\n<p>We have covered
  topics like creating basic post requests, Marshalling golang types into JSON format,
  parsing form data, sending a POST request with files, and best practices for the
  POST method. Hope you found this article helpful. If you have any queries, questions,
  or feedback, please let me know in the comments or on my social handles.</p>\n<p>Happy
  Coding :)</p>\n"
cover: ''
date: 2024-03-10
datetime: 2024-03-10 00:00:00+00:00
description: Exploring the fundamentals of a POST method in golang. How to send a
  basic POST request, parse json, structs, files, form-data into the request body.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2024-03-09-Golang-POST-Method.md
html: "<h2>Introduction</h2>\n<p>In this section of the series, we will be exploring
  how to send a <code>POST</code> HTTP request in golang. We will understand how to
  send a basic POST request, create an HTTP request, and parse json, structs into
  the request body, add headers, etc in the following sections of this post. We will
  understand how to marshal the golang struct/types into JSON format, send files in
  the request, and handle form data with examples of each in this article. Let's answer
  a few questions first.</p>\n<h2>What is a POST request?</h2>\n<p>POST method is
  a type of request that is used to send data to a server(a machine on the internet).</p>\n<p>Imagine
  you are placing an order at a restaurant. With a GET request, it would be like asking
  the waiter, &quot;What kind of pizza do you have?&quot; The waiter would respond
  by telling you the menu options (the information retrieved from the server).</p>\n<p>However,
  a POST request is more like giving your completed order to the waiter. You tell
  them the specific pizza you want, its size, and any additional toppings (the data
  you send). The waiter then takes this information (POST request) back to the kitchen
  (the server) to process it (fulfill your order).</p>\n<p>In the world of web development,
  POST requests are often used for things like:</p>\n<p>Submitting forms (e.g., contact
  forms, login forms) Uploading files (e.g., photos, videos) Creating new accounts
  Sending data to be processed (e.g., online purchases)</p>\n<p>Here's an example
  of what the POST request might look like in this scenario:</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"k\">POST</span><span class=\"w\"> </span><span class=\"s\">/api/order</span><span
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
  class=\"w\"> </span><span class=\"mi\">2</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"err\">}</span><span class=\"s\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">            </span><span class=\"kn\">&quot;itemID&quot;:</span><span
  class=\"w\"> </span><span class=\"mi\">999</span><span class=\"s\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">            </span><span class=\"s\">&quot;name&quot;:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Burger&quot;,</span><span class=\"w\"></span>\n<span
  class=\"w\">            </span><span class=\"s\">&quot;quantity&quot;:</span><span
  class=\"w\"> </span><span class=\"mi\">1</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"err\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"s\">]</span><span class=\"w\"></span>\n<span
  class=\"err\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  this example:</p>\n<ul>\n<li>The <code>POST</code> method is used to send data to
  the server.</li>\n<li>The <code>/api/order</code> is the endpoint of the server.</li>\n<li>The
  <code>application/json</code> is the content type of the request.</li>\n<li>The
  <code>123</code> is the content length of the request.</li>\n<li>The <code>{&quot;userID&quot;:
  123, &quot;orderID&quot;: 456, &quot;items&quot;: [{&quot;itemID&quot;: 789, &quot;name&quot;:
  &quot;Pizza&quot;, &quot;quantity&quot;: 2}, {&quot;itemID&quot;: 999, &quot;name&quot;:
  &quot;Burger&quot;, &quot;quantity&quot;: 1}]}</code> is the body of the request.</li>\n</ul>\n<h2>Why
  the need for a POST request?</h2>\n<p>In the world of HTTP requests, we use the
  POST method to securely send data from a client (like a user's browser) to a server.
  This is crucial because the GET method, while convenient for retrieving data, has
  limitations:</p>\n<p>Imagine you are in registering for an event via Google form,
  you type in your details on the webpage like name, email, address, phone number,
  and other personal details. If the website/app was using the <code>GET</code> method
  to send the request to register or do any other authentication/privacy-related requests,
  it could expose the data in the URL itself. It would be something along the lines
  <a href=\"https://form.google.com/register/%3Cform-name%3E-%3Cid%3E/?name=John&amp;phone_number=1234567890\"><code>https://form.google.com/register/&lt;form-name&gt;-&lt;id&gt;/?name=John&amp;phone_number=1234567890</code></a>,
  if a user maliciously sniffs into your network and inspects the URL, your data will
  be exposed. That is the reason we need <code>POST</code> a method.</p>\n<h2>How
  a POST method works?</h2>\n<p>A <a href=\"https://www.rfc-editor.org/rfc/rfc9110#POST\">POST</a>
  request is used to send data to a server to create or update(there is a separate
  method for updating) a resource. The client(browser/other APIs) sends a POST request
  to the server's API endpoint with the data in the request body. This data can be
  in formats like JSON, XML, or form data. The server processes the POST request,
  validates and parses the data in the request body, makes any changes or creates
  resources based on that data, and returns a response. The response would contain
  a status code indicating the success or failure of the operation and may contain
  the newly created or updated resource in the response body. The client must check
  the response status code to verify the outcome and process the response accordingly.
  Unlike GET, POST can create new resources on the server. The body of a POST contains
  the data for creation while the URL identifies the resource to be created. Overall,
  POST transfers data to the server for processing, creation or updating of resources.</p>\n<p>The
  status code is usually <code>201</code> indicating the resource is successfully
  created or <code>200</code> for just indicating success.</p>\n<p>Some common steps
  for creating and sending a POST request as a developer include:</p>\n<ul>\n<li>\n<p>Defining
  the API endpoint</p>\n</li>\n<li>\n<p>Clarifying the data format (json, language
  native objects, xml , text, form-data, etc)</p>\n</li>\n<li>\n<p>Converting / Marshalling
  the data</p>\n</li>\n<li>\n<p>Attaching header for <code>Content-Type</code> as
  key and value as the format of the data type (e.g. <code>application/json</code>
  for json)</p>\n</li>\n<li>\n<p>Sending the request</p>\n</li>\n</ul>\n<p>The above
  steps are general for creating and sending a POST request, they are not specific
  to Golang. For golang specific steps, we need to dive a bit deeper, let's get started.</p>\n<h2>Basic
  POST method in Golang</h2>\n<p>To send a POST request in golang, we need to use
  the <code>http</code> package. The <code>http</code> package has the <code>Post</code>
  method, which takes in 3 parameters, namely the URL, the Content-Type, and the Body.
  The body can be <code>nil</code> if the URL endpoint doesn't necessarily require
  a body. The <code>Content-Type</code> is the string, since we are just touching
  on how the Post request is constructed, we will see what the <code>Content-Type</code>
  string value should be in the later sections.</p>\n<blockquote>\n<p><a href=\"http://http.Post\">http.Post</a>(URL,
  Content-Type, Body)</p>\n</blockquote>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in/api/users&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">// POST request</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Post</span><span class=\"p\">(</span><span
  class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"kc\">nil</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"c1\">// ideally the Content-Type header should
  be set to the relevant format</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// resp, err := http.Post(apiURL, &quot;application/json&quot;, nil)</span><span
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
  class=\"w\">    </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
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
  go run main.go\n\n<span class=\"m\">201</span>\n<span class=\"p\">&amp;</span><span
  class=\"o\">{</span>\n    <span class=\"m\">201</span> Created\n    <span class=\"m\">201</span>\n
  \   HTTP/2.0\n    <span class=\"m\">2</span>\n    <span class=\"m\">0</span>\n    map<span
  class=\"o\">[</span>\n        Access-Control-Allow-Origin:<span class=\"o\">[</span>*<span
  class=\"o\">]</span>\n        Cf-Cache-Status:<span class=\"o\">[</span>DYNAMIC<span
  class=\"o\">]</span>\n        Cf-Ray:<span class=\"o\">[</span>861cd9aec8223e4b-BOM<span
  class=\"o\">]</span>\n        Content-Length:<span class=\"o\">[</span><span class=\"m\">50</span><span
  class=\"o\">]</span>\n        Content-Type:<span class=\"o\">[</span>application/json<span
  class=\"p\">;</span> <span class=\"nv\">charset</span><span class=\"o\">=</span>utf-8<span
  class=\"o\">]</span>\n        Date:<span class=\"o\">[</span>Sat, <span class=\"m\">09</span>
  Mar <span class=\"m\">2024</span> <span class=\"m\">17</span>:40:28 GMT<span class=\"o\">]</span>\n
  \       Server:<span class=\"o\">[</span>cloudflare<span class=\"o\">]</span>\n
  \       ...\n        ...\n        ...\n        X-Powered-By:<span class=\"o\">[</span>Express<span
  class=\"o\">]</span>\n    <span class=\"o\">]</span>\n    <span class=\"o\">{</span>0xc00017c180<span
  class=\"o\">}</span>\n    <span class=\"m\">50</span>\n    <span class=\"o\">[]</span>\n
  \   <span class=\"nb\">false</span>\n    <span class=\"nb\">false</span>\n    map<span
  class=\"o\">[]</span>\n    0xc000156000\n    0xc00012a420\n<span class=\"o\">}</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above code is sending the <code>POST</code> request to the <a href=\"https://reqres.in/api/users\"><code>https://reqres.in/api/users</code></a>
  endpoint with an empty body and no specific format for <code>Content-Type</code>
  header. The response is according to the <a href=\"https://pkg.go.dev/net/http#Response\">Response</a>
  structure. We can see we got <code>201</code> status, which indicates the server
  received the POST request successfully, the API is a dummy api, so we don't care
  about the data we are processing, we are just using the API as a placeholder for
  sending the POST request.</p>\n<p>We can use <code>map[string]interface{}</code>
  it to pass the data in the request body. The <code>json.Marshal</code> method is
  used to convert the map into JSON format. We will look into the details shortly
  in the next few examples.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in/api/users&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">bodyMap</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"kt\">string</span><span class=\"p\">]</span><span
  class=\"kd\">interface</span><span class=\"p\">{}{</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"s\">&quot;name&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;morpheus&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"s\">&quot;job&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s\">&quot;leader&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">    </span><span
  class=\"nx\">requestBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Marshal</span><span class=\"p\">(</span><span class=\"nx\">bodyMap</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">body</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">bytes</span><span class=\"p\">.</span><span
  class=\"nx\">NewBuffer</span><span class=\"p\">(</span><span class=\"nx\">requestBody</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">    </span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">Post</span><span class=\"p\">(</span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;application/json&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">defer</span><span class=\"w\"> </span><span
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
  go run main.go\n\n<span class=\"m\">201</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above code sends the <code>POST</code> request to the <a href=\"https://reqres.in/api/users\"><code>https://reqres.in/api/users</code></a>
  endpoint with the data in the request body in JSON format.</p>\n<h2>Creating a POST
  request in Golang</h2>\n<p>We can construct the POST request with the <a href=\"https://pkg.go.dev/net/http#NewRequest\">NewRequest</a>
  method. The method takes in 3 parameters, namely the <code>method</code> (e.g. <code>POST</code>,
  <code>GET</code>), the <code>URL</code> and the <code>body</code> (if there is any).
  We can then add extra information to the headers or the Request object after constructing
  the basic HTTP <a href=\"https://pkg.go.dev/net/http#Request\">Request</a> object.</p>\n<pre
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
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in/api/users&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">MethodPost</span><span
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
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"c1\">//fmt.Println(resp)</span><span class=\"w\"></span>\n<span
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
  go run main.go\n\n<span class=\"m\">201</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we have created an HTTP Request as the <code>POST</code> method,
  with <a href=\"https://reqres.in/api/users\"><code>https://reqres.in/api/users</code></a>
  as the URL, and no body. This constructs an HTTP Request object, which can be sent
  as the parameter to the <a href=\"http://http.DefaultClient.Do\">http.DefaultClient.Do</a>
  method, which is the default client for the request we sent in the earlier examples
  as <code>http.Get</code> or <a href=\"http://http.Post\"><code>http.Post</code></a>
  methods. We can implement a custom client as well, and then apply <code>Do</code>
  the method with the request parameters. The <code>Do</code> method returns the <code>Request</code>
  object or the <code>error</code> if any.</p>\n<p>More on the customizing Client
  will be explained in a separate post in the series.</p>\n<p>The response is also
  in the same format as the <a href=\"https://pkg.go.dev/net/http#Response\">Response</a>
  structure that we have seen earlier. This section of the series aims to construct
  a post request, and not to parse the response, we have already understood the parsing
  of the response in the <a href=\"https://www.meetgor.com/golang-web-get-method/#?:~:text=Parsing%20the%20JSON%20body%20with%20structs\">Get
  method</a> section of the series.</p>\n<h3>Parsing objects to JSON for POST method
  request</h3>\n<p>We might have a golang object that we want to send as a body to
  an API in the POST request, for that we need to convert the golang struct object
  to JSON. We can do this by using the <a href=\"https://pkg.go.dev/encoding/json#Marshal\">Marshal</a>
  or the <a href=\"https://pkg.go.dev/encoding/json#Encoder.Encode\">Encode</a> method
  for serialization of the golang struct object to JSON.</p>\n<h4>Using Marshal method</h4>\n<p>Marshaling
  is the process of converting data from a data structure into a format suitable for
  transmission over a network or for storage. It's commonly used to convert native
  objects in a programming language into a serialized format, typically a byte stream,
  that can be transmitted or stored efficiently. You might get a question here, what
  is the difference between <code>Marshalling</code> and <code>Serialization</code>?
  Well, Serialization, is a broader term that encompasses marshalling. It refers to
  the process of converting an object or data structure into a format that can be
  stored or transmitted and later reconstructed into the original object. Serialization
  may involve converting data into byte streams, XML, JSON, or other formats. So,
  in summary, marshaling specifically deals with converting native objects into a
  format suitable for transmission, while serialization encompasses the broader process
  of preparing data for storage or transmission.</p>\n<p>The <code>json</code> package
  has the <a href=\"https://pkg.go.dev/encoding/json#Marshal\">Marshal</a> method
  that converts the golang object into JSON. The <code>Marshal</code> method takes
  in a parameter as the struct object with type <code>any</code> and returns a byte
  slice <code>[]byte</code> and error (if any).</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;bytes&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">User</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Name</span><span class=\"w\">   </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;name&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Salary</span><span class=\"w\"> </span><span
  class=\"kt\">int</span><span class=\"w\">    </span><span class=\"s\">`json:&quot;salary&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Age</span><span
  class=\"w\">    </span><span class=\"kt\">int</span><span class=\"w\">    </span><span
  class=\"s\">`json:&quot;age&quot;`</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">user</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">User</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">Name</span><span
  class=\"p\">:</span><span class=\"w\">   </span><span class=\"s\">&quot;Alice&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">Salary</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">50000</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">Age</span><span class=\"p\">:</span><span
  class=\"w\">    </span><span class=\"mi\">25</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">apiURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://dummy.restapiexample.com/api/v1/create&quot;</span><span
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
  class=\"w\">\t</span><span class=\"nx\">resp</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Post</span><span class=\"p\">(</span><span
  class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;application/json&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">body</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  go run main.go\n\n<span class=\"o\">{</span><span class=\"s2\">&quot;name&quot;</span>:<span
  class=\"s2\">&quot;Alice&quot;</span>,<span class=\"s2\">&quot;salary&quot;</span>:50000,<span
  class=\"s2\">&quot;age&quot;</span>:25<span class=\"o\">}</span>\n<span class=\"m\">200</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we have created a struct <code>User</code> with fields <code>Name</code>,
  <code>Salary</code>, and <code>Age</code>, the json tags will help label each key
  in JSON with the tag for the respective fields in the struct. We create an object
  <code>user</code> of a type <code>User</code> with the values as <code>Alice</code>,
  <code>50000</code>, and <code>25</code> respectively.</p>\n<p>We call the <code>json.Marshal</code>
  method with the parameter <code>user</code> that represents the struct object <code>User</code>,
  the method returns a slice of bytes, or an error either or both could be nil. If
  we try to see the stringified representation of the byte slice, we can see something
  like <code>{&quot;name&quot;:&quot;Alice&quot;,&quot;salary&quot;:50000,&quot;age&quot;:25}</code>
  which is a JSON string for the user struct. We can't parse the byte slice as the
  body in the POST request, we need the <code>io.Reader</code> object, so we can load
  the byte slice <code>bodyBytes</code> into a buffer and parse that as a body for
  the POST request.</p>\n<p>We then send a <code>POST</code> request to the endpoint
  <a href=\"https://dummy.restapiexample.com/api/v1/create\"><code>https://dummy.restapiexample.com/api/v1/create</code></a>
  with the content type as <code>application/json</code> and with the body as <code>body</code>
  which was a <code>io.Reader</code> object as an in-memory buffer.</p>\n<p>In brief,
  we can summarize the marshaling of the golang object into JSON with <code>Marshal</code>
  function as the following steps:</p>\n<ul>\n<li>\n<p>Defining the structure as per
  the request body</p>\n</li>\n<li>\n<p>Creating the struct object for parsing the
  data as body to the request</p>\n</li>\n<li>\n<p>Calling the <code>json.Marshal</code>
  function to convert the object to JSON (parameter as the struct object <code>any</code>
  type)</p>\n</li>\n<li>\n<p>Loading the byte slice into a buffer with <code>bytes.NewBuffer()</code></p>\n</li>\n<li>\n<p>Sending
  the POST request to the endpoint with the body as the <code>io.Reader</code> object
  and content type as <code>application/json</code></p>\n</li>\n</ul>\n<h4>Using Encode
  method</h4>\n<p>We can even use the <a href=\"https://pkg.go.dev/encoding/json#Encoder.Encode\">Encoder.Encode</a>
  method to parse the golang struct object to JSON. Firstly, we should have the struct
  defined as per the request body that the particular API takes, we can make use of
  the json tags, omitempty, omit(-) options to make the marshaling process work accordingly.
  We can then create the object of that particular struct with the data we require
  to be created as a resource with the POST request on that API service.</p>\n<p>Thereafter
  we can create an empty buffer object with <a href=\"https://pkg.go.dev/bytes#Buffer\">bytes.Buffer</a>,
  this buffer object would be used to initialize the <a href=\"https://pkg.go.dev/encoding/json#Encoder\">Encoder</a>
  object with the <a href=\"https://pkg.go.dev/encoding/json#NewEncoder\">NewEncoder</a>
  method. This would give access to the <a href=\"https://pkg.go.dev/encoding/json#Encoder.Encode\">Encode</a>
  method, which is used to take in the struct object (<code>any</code> type) and this
  will populate the buffer we initialized with the <code>NewEncoder</code> method.</p>\n<p>Later
  we can access that buffer to parse it to the Post request as the body. Let's understand
  it better with an example.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">User</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Name</span><span class=\"w\">   </span><span class=\"kt\">string</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Salary</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Age</span><span class=\"w\">    </span><span
  class=\"kt\">int</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">user</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">User</span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">Name</span><span
  class=\"p\">:</span><span class=\"w\">   </span><span class=\"s\">&quot;Alice&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">Salary</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">50000</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">Age</span><span class=\"p\">:</span><span
  class=\"w\">    </span><span class=\"mi\">25</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">apiURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://dummy.restapiexample.com/api/v1/create&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">bodyBuffer</span><span class=\"w\"> </span><span
  class=\"nx\">bytes</span><span class=\"p\">.</span><span class=\"nx\">Buffer</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">encoder</span><span class=\"w\"> </span><span
  class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">NewEncoder</span><span class=\"p\">(</span><span
  class=\"o\">&amp;</span><span class=\"nx\">bodyBuffer</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">encoder</span><span
  class=\"p\">.</span><span class=\"nx\">Encode</span><span class=\"p\">(</span><span
  class=\"nx\">user</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">Post</span><span class=\"p\">(</span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;application/json&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">bodyBuffer</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Over
  here, we have created a struct <code>User</code> with fields <code>Name</code>,
  <code>Salary</code>, and <code>Age</code>, we initialize the <code>user</code> as
  the object of the <code>User</code> struct. Then we create a buffer <code>bodyBuffer</code>
  of type <code>bytes.Buffer</code> this is the actual buffer that we will send as
  the body. Further, we initialize the <code>Encoder</code> object as <code>encoder</code>
  with the <code>json.NewEncoder</code> method by parsing the reference of <code>bodyBuffer</code>
  as the parameter. Since <code>bytes.Buffer</code> implements the <code>io.Writer</code>
  interface, we can pass the <code>bodyBuffer</code> to the <code>NewEncoder</code>
  method. This will create the <code>Encoder</code> object which in turn will give
  us access to the <code>Encode</code> method, where we will parse the struct instance
  and it will populate the buffer with which we initialized the <code>Encoder</code>
  object earlier.</p>\n<p>Now, we have the <code>encode</code> object, this gives
  us the access to <code>Encode</code> method, we call the <code>Encode</code> method
  with the parameter of <code>user</code> which is a User struct instance/object.
  The Encode method will populate the <code>bodyBuffer</code> object or it will result
  in an error if anything goes wrong (the data is incorrectly parsed or is not in
  the required format).</p>\n<p>We can call the <code>Post</code> method with the
  initialized URL, the <code>Content-Type</code> as <code>application/json</code>
  since we have converted the struct instance to JSON object, and the body as the
  reference to the buffer as <code>&amp;bodyBuffer</code></p>\n<p>So, the steps for
  parsing struct instances into JSON objects with the <code>Encoder.Encode</code>
  method is as follows:</p>\n<ul>\n<li>\n<p>Defining the structure as per the request
  body</p>\n</li>\n<li>\n<p>Creating the struct object for parsing the data as body
  to the request</p>\n</li>\n<li>\n<p>Creating an empty <code>bytes.Buffer</code>
  object as an in-memory buffer</p>\n</li>\n<li>\n<p>Initializing the <code>Encoder</code>
  object with <code>NewEncoder</code> method by parsing the reference of <code>bodyBuffer</code>
  as the parameter</p>\n</li>\n<li>\n<p>Calling the <code>Encode</code> method with
  the parameter of struct instance/object</p>\n</li>\n<li>\n<p>Sending the POST request
  to the endpoint with the content type as <code>application/json</code> and body
  as the reference to the buffer</p>\n</li>\n</ul>\n<p>The results are the same as
  the above example just the way we have parsed the struct instance to JSON object
  is different.</p>\n<h3>Parsing JSON to POST request</h3>\n<p>We have seen how we
  can parse golang struct instances to JSON and then send the post request, but what
  if we had the JSON string already with us, and we want to send the request? Well,
  that's much easier, right? We already have parsed the JSON string to the Post request
  by loading the slice of bytes into a buffer, so we just need to convert the string
  to a slice of bytes which is quite an easy task, and then load that byte slice to
  the buffer.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// dummy api</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy.restapiexample.com/api/v1/create&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">// json data</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">data</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`{</span>\n<span
  class=\"s\">        &quot;name&quot;: &quot;Alice&quot;,</span>\n<span class=\"s\">
  \       &quot;job&quot;: &quot;Teacher&quot;</span>\n<span class=\"s\">    }`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">body</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">bytes</span><span
  class=\"p\">.</span><span class=\"nx\">NewBuffer</span><span class=\"p\">([]</span><span
  class=\"nb\">byte</span><span class=\"p\">(</span><span class=\"nx\">data</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"c1\">// POST request</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">Post</span><span class=\"p\">(</span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;application/json&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In the
  example above, we already have a JSON string <code>data</code> with keys as <code>name</code>
  and <code>job</code> but it is not JSON, it is a stringified JSON. We can convert
  the stringified JSON to a slice of bytes using the <code>[]byte</code> function.
  Further, we have used the <code>bytes.NewBuffer</code> method to load the byte slice
  into an <code>io.Reader</code> object. This object returned by the <code>bytes.NewBuffer</code>
  will serve as the body for the POST request.</p>\n<h3>Parsing JSON to objects in
  Golang from POST method response</h3>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nx\">Status</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;status&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Data</span><span class=\"w\">   </span><span
  class=\"nx\">User</span><span class=\"w\">   </span><span class=\"s\">`json:&quot;data&quot;`</span><span
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
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy.restapiexample.com/api/v1/create&quot;</span><span
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
  class=\"w\">\t</span><span class=\"nx\">resp</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Post</span><span class=\"p\">(</span><span
  class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;application/json&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">body</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">// Read response body</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">respBody</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"c1\">// unmarshalling process</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// converting JSON to Go specific data structure/types</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">userResponse</span><span class=\"w\"> </span><span
  class=\"nx\">UserResponse</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
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
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">userResponse</span><span class=\"p\">.</span><span class=\"nx\">Data</span><span
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
  \       \n<div class='language-bar'>nginx</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">{success</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"kn\">Alice</span><span class=\"w\"> </span><span class=\"mi\">50000</span><span
  class=\"w\"> </span><span class=\"mi\">25</span><span class=\"w\"> </span><span
  class=\"mi\">3239</span><span class=\"err\">}}</span><span class=\"w\"></span>\n<span
  class=\"p\">{</span><span class=\"kn\">Alice</span><span class=\"w\"> </span><span
  class=\"mi\">50000</span><span class=\"w\"> </span><span class=\"mi\">25</span><span
  class=\"w\"> </span><span class=\"mi\">577</span><span class=\"err\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The above example is a POST request
  with a struct instance being loaded as a JSON string and then sent as a buffer to
  the API endpoint, it also reads the response body with a specific structure <code>UserResponse</code>
  and unmarshalled the <code>resp.Body</code> from the <code>io.Reader</code> as <code>respBody</code>
  and then loads into <code>userResponse</code> object. This example gives an entire
  process of what we have understood in the JSON data parsing for a POST request.</p>\n<h3>Sending
  Form data in a POST request</h3>\n<p>We can also send data to a POST request in
  the form of a form, the form which we use in the HTML. Golang has a <code>net/url</code>
  package to parse the form data. The form data is sent in the <code>application/x-www-form-urlencoded</code>
  format.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/url&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">ResponseLogin</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Token</span><span class=\"w\"> </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;token&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// dummy api</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in/api/login&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">// Define form
  data</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">formData</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">url</span><span
  class=\"p\">.</span><span class=\"nx\">Values</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">formData</span><span
  class=\"p\">.</span><span class=\"nx\">Set</span><span class=\"p\">(</span><span
  class=\"s\">&quot;email&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;eve.holt@reqres.in&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">formData</span><span
  class=\"p\">.</span><span class=\"nx\">Set</span><span class=\"p\">(</span><span
  class=\"s\">&quot;password&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;cityslicka&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">// Encode the
  form data</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">formData</span><span class=\"p\">.</span><span class=\"nx\">Encode</span><span
  class=\"p\">())</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">reqBody</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">strings</span><span class=\"p\">.</span><span
  class=\"nx\">NewReader</span><span class=\"p\">(</span><span class=\"nx\">formData</span><span
  class=\"p\">.</span><span class=\"nx\">Encode</span><span class=\"p\">())</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">reqBody</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">// Make a POST request with form data</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Post</span><span class=\"p\">(</span><span
  class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"s\">&quot;application/x-www-form-urlencoded&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">reqBody</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">// Print response status code</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Status Code:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">StatusCode</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"c1\">// Read response body</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">io</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">token</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">ResponseLogin</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">token</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">token</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  go run main.go\n\n<span class=\"nv\">email</span><span class=\"o\">=</span>eve.holt%40reqres.in<span
  class=\"p\">&amp;</span><span class=\"nv\">password</span><span class=\"o\">=</span>cityslicka\n<span
  class=\"p\">&amp;</span><span class=\"o\">{</span><span class=\"nv\">email</span><span
  class=\"o\">=</span>eve.holt%40reqres.in<span class=\"p\">&amp;</span><span class=\"nv\">password</span><span
  class=\"o\">=</span>cityslicka <span class=\"m\">0</span> -1<span class=\"o\">}</span>\nStatus
  Code: <span class=\"m\">200</span>\n<span class=\"o\">{</span>QpwL5tke4Pnpja7X4<span
  class=\"o\">}</span>\n</pre></div>\n\n</pre>\n\n<p>In the above example, we set
  a <code>formData</code> with the values of <code>email</code> and <code>password</code>
  which are <code>url.Values</code> object. The <code>url.Values</code> the object
  is used to store the key-value pairs of the form data. The <code>formData</code>
  is encoded with the <code>url.Encode</code> method, We load the encoded string to
  a <code>buffer</code> with <code>strings.NewReader</code> which implements the <code>io.Reader</code>
  interface, so that way we can pass that object as the body to the post request.</p>\n<p>We
  send the <code>POST</code> request to the endpoint <a href=\"https://reqres.in/api/login\"><code>https://reqres.in/api/login</code></a>
  with the content type as <code>application/x-www-form-urlencoded</code> and with
  the body as <code>reqBody</code> which implements the <code>io.Reader</code> interface
  as an in-memory buffer. The response from the request is read into the buffer with
  <code>io.ReadAll</code> method and we can <code>Unmarshal</code> the stream of bytes
  as a buffer into the <code>ResponseLogin</code> struct object.</p>\n<p>The output
  shows the <code>formData</code> as encoded string <code>email=eve.holt%</code><a
  href=\"http://40reqres.in\"><code>40reqres.in</code></a><code>&amp;password=cityslicka</code>
  as <code>@</code> is encoded to <code>%40</code>, then we wrap the <code>formData</code>
  in a <code>strings.NewReader</code> object which is a buffer that implements <code>io.Reader</code>
  interface, hence we can see the result as the object. The status code for the request
  is <code>200</code> indicating the server received the <code>form-data</code> in
  the body and upon unmarshalling, we get the token as a response to the POST request
  which was a dummy login API.</p>\n<p>This way we have parsed the form-data to the
  body of a POST request.</p>\n<h3>Sending File in a POST request</h3>\n<p>We have
  covered, parsing text, JSON, and form data, and now we need to move into sending
  files in a POST request. We can use the <code>multipart</code> package to parse
  files into the request body and set appropriate headers for reading the file from
  the API services.</p>\n<p>We first read the file contents <a href=\"http://os.Open\"><code>os.Open</code></a>
  which returns a reference to the <code>file</code> object or an error. We create
  an empty <code>bytes.Buffer</code> object as <code>body</code> which will be populated
  later. The <a href=\"https://pkg.go.dev/mime/multipart#NewWriter\">multipart.NewWriter</a>
  method takes in the <code>io.Writer</code> object which will be the <code>body</code>
  as it is an <code>bytes.Buffer</code> object that implements the <code>io.Writer</code>
  interface. This will initialize the <a href=\"https://pkg.go.dev/mime/multipart#Writer\">Writer</a>
  object in the <code>multipart</code> package.</p>\n<p>We create a <code>form-field</code>
  in the <code>Writer</code> object with the <a href=\"https://pkg.go.dev/mime/multipart#Writer.CreateFormFile\">CreateFormFile</a>
  method, which takes in the <code>fieldName</code> as the name of the field, and
  the <code>fileName</code> as the name of the file which will be read later in the
  multipart form. The method returns either the part or the error. The <code>part</code>
  is an object that implements the <code>io.Writer</code> interface.</p>\n<p>Since
  we have stored the file contents in the <code>file</code> object, we copy the contents
  into the <code>form-field</code> with the <a href=\"https://pkg.go.dev/io#Copy\">Copy</a>
  method. Since the <code>part</code> return from the <code>CreateFormFile</code>
  was implementing the <code>io.Writer</code> interface, we can use it to Copy the
  contents from source to destination. The source is the <code>io.Reader</code> object
  and the destination is the <code>io.Writer</code> object, the destination for the
  <code>Copy</code> method is the first parameter, the source is the second parameter.</p>\n<p>This
  Copy method will populate the buffer initialized earlier in the <code>NewWriter</code>
  method. This will give us a buffer that has the file contents in it. We can pass
  this buffer to the POST request with the <code>body</code> parameter. We also need
  to make sure we close the <code>Writer</code> object after copying the contents
  of the file. We can extract the type of file which will serve as the <code>Content-Type</code>
  of the request.</p>\n<p>Let's clear the explanation with an example.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;bytes&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;mime/multipart&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;os&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">ResponseFile</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Files</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"kt\">string</span><span class=\"p\">]</span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;files&quot;`</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">apiURL</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;http://postman-echo.com/post&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fileName</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;sample.csv&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">file</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">os</span><span
  class=\"p\">.</span><span class=\"nx\">Open</span><span class=\"p\">(</span><span
  class=\"nx\">fileName</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">file</span><span
  class=\"p\">.</span><span class=\"nx\">Close</span><span class=\"p\">()</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">body</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">bytes</span><span class=\"p\">.</span><span class=\"nx\">Buffer</span><span
  class=\"p\">{}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">writer</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">multipart</span><span class=\"p\">.</span><span
  class=\"nx\">NewWriter</span><span class=\"p\">(</span><span class=\"nx\">body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">part</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">writer</span><span class=\"p\">.</span><span
  class=\"nx\">CreateFormFile</span><span class=\"p\">(</span><span class=\"s\">&quot;csvFile&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">fileName</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">_</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">Copy</span><span class=\"p\">(</span><span
  class=\"nx\">part</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">file</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">    </span><span
  class=\"nx\">contentType</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">writer</span><span class=\"p\">.</span><span
  class=\"nx\">FormDataContentType</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">contentType</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">writer</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">Post</span><span class=\"p\">(</span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">contentType</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Status
  Code:&quot;</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">StatusCode</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">respBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">token</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">ResponseFile</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">token</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">token</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">token</span><span
  class=\"p\">.</span><span class=\"nx\">Files</span><span class=\"p\">[</span><span
  class=\"nx\">fileName</span><span class=\"p\">])</span><span class=\"w\"></span>\n<span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>multipart/form-data<span
  class=\"p\">;</span> <span class=\"nv\">boundary</span><span class=\"o\">=</span>7e0eacfff890be395eba19c70415c908124b503a56f23ebeec0ab3c665ca\n\n\n--619671ea2c0aa47ca6664a7cda422169d73f3b8a089c659203f5413d03de\nContent-Disposition:
  form-data<span class=\"p\">;</span> <span class=\"nv\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;csvFile&quot;</span><span class=\"p\">;</span> <span class=\"nv\">filename</span><span
  class=\"o\">=</span><span class=\"s2\">&quot;sample.csv&quot;</span>\nContent-Type:
  application/octet-stream\n\nUser,City,Age,Country\nAlex Smith,Los Angeles,20,USA\nJohn
  Doe,New York,30,USA\nJane Smith,Paris,25,France\nBob Johnson,London,40,UK\n\n--619671ea2c0aa47ca6664a7cda422169d73f3b8a089c659203f5413d03de--\n\n\n\nStatus
  Code: <span class=\"m\">200</span>\n\n<span class=\"o\">{</span>map<span class=\"o\">[</span>sample.csv:data:application/octet-stream<span
  class=\"p\">;</span>base64,VXNlcixDaXR5LEFnZSxDb3VudHJ5CkFsZXggU21pdGgsTG9zIEFuZ2VsZXMsMjAsVVNBCkpvaG4gRG9lLE5ldyBZb3JrLDMwLFVTQQpKYW5lIFNtaXRoLFBhmlzLDI1LEZyYW5jZQpCb2IgSm9obnNvbixMb25kb24sNDAsVUsK<span
  class=\"o\">]}</span>\n\ndata:application/octet-stream<span class=\"p\">;</span>base64,VXNlcixDaXR5LEFnZSxDb3VudHJ5CkFsZXggU21pdGgsTG9zIEFuZ2VsZXMsMjAsVVNBCkpvaG4gRG9lLE5ldyBZb3JrLDMwLFVTQQpKYW5lIFNtaXRoLFBhmlzLDI1LEZyYW5jZQpCb2IgSm9obnNvbixMb25kb24sNDAsVUsK\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we first read the file <code>sample.csv</code> into the <code>file</code>
  object with <a href=\"http://os.Open\"><code>os.Open</code></a> method, this will
  return a reference to the file object or return an error if any arises while opening
  the file.</p>\n<p>Then we create an empty buffer <code>bytes.Buffer</code> object
  which will serve as the body of the post request later as it will get populated
  with the file contents in the form of <code>multipart/form-data</code>.</p>\n<p>We
  initialize the <code>Writer</code> object with <code>multipart.NewWriter</code>
  method which takes in the empty buffer as the parameter, we parse the <code>body</code>
  as the parameter. The method will return a reference to the <code>multipart.Writer</code>
  object.</p>\n<p>With the <code>Writer</code> object we access the <code>CreateFormFile</code>
  method which takes in the <code>fieldName</code> as the name of the field, and the
  <code>fileName</code> as the name of the file. The method will return either the
  part or an error. The <code>part</code> in this case, is the reference to the <code>io.Writer</code>
  object that will be used to write the contents from the uploaded file.</p>\n<p>Then,
  we can use the <code>io.Copy</code> method to copy the contents from the <code>io.Reader</code>
  object to the <code>io.Writer</code> object. The source is the <code>io.Reader</code>
  object and the destination is the <code>io.Writer</code> object. The first parameter
  is however the destination and the second parameter is the source. In the example,
  we call <code>io.Copy(part, file)</code> which will copy the contents of <code>file</code>
  to the <code>part</code> buffer.</p>\n<p>We get the <code>Content-Type</code> by
  calling the <a href=\"https://pkg.go.dev/mime/multipart#Writer.FormDataContentType\">Writer.FormDataContentType</a>
  method. This returns us <code>multipart/form-data; boundary=7e0eacfff890be395eba19c70415c908124b503a56f23ebeec0ab3c665ca</code>
  which will serve the <code>Content-Type</code> for the Post request.</p>\n<p>We
  need to make sure we close the <code>Writer</code> object with the <code>Close</code>
  method.</p>\n<p>We just print the <code>body.String()</code> to get a look at what
  the actual body looks like, we can see there is a form for the file as a <code>form-data</code>
  with keys like <code>Content-Type</code>, <code>Content-Disposition</code>, etc.
  The file has the <code>Content-Type</code> as <code>application/octet-stream</code>
  and the actual content is rendered in the output.</p>\n<p>The dummy API responds
  with a 200 status code and also sends the JSON data with the name of the file as
  the key and the value as the <code>base64</code> encoded value of the file contents.
  This indicates that we were able to upload the file to the server API using a POST
  request. Well done!</p>\n<p>I have also included some more examples of POST requests
  with files <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/post/file_2.go\">here</a>
  which extends the above example by taking the encoded values and decoding to get
  the actual contents of the file back.</p>\n<h2>Best Practices for POST method</h2>\n<p>Here
  are some of the best practices for the POST method which are followed to make sure
  you consume or create the POST request in the most secure, efficient, and graceful
  way.</p>\n<h3>Always Close the Response Body</h3>\n<p>Ensure that you close the
  response body after reading from it. Use <code>defer response.Body.Close()</code>
  to automatically close the body when the surrounding function returns. This is crucial
  for releasing associated resources like network connections or file descriptors.
  Failure to close the response body can lead to memory leaks, particularly with a
  large volume of requests. Properly closing the body prevents resource exhaustion
  and maintains efficient memory usage.</p>\n<h3>Client Customization</h3>\n<p>Utilize
  the <a href=\"https://pkg.go.dev/net/http#Client\">Client</a> struct to customize
  the HTTP client behavior. By using a custom client, you can set timeouts, headers,
  user agents, and other configurations without modifying the <code>DefaultClient</code>
  provided by the <code>http</code> package. This approach allows for flexibility
  and avoids repetitive adjustments to the client configuration for each request.</p>\n<h3>Set
  Content-Type Appropriately</h3>\n<p>Ensure that you set the <code>Content-Type</code>
  header according to the request payload. Correctly specifying the Content-Type is
  crucial for the server to interpret the request payload correctly. Failing to set
  the Content-Type header accurately may result in the server rejecting the request.
  Always verify and match the Content-Type header with the content being sent in the
  POST request to ensure smooth communication with the server.</p>\n<h2>Reference</h2>\n<ul>\n<li>\n<p><a
  href=\"https://www.postman.com/postman/workspace/postman-answers/documentation/13455110-00378d5c-5b08-4813-98da-bc47a2e6021d\">Postman
  POST API</a> (For POST request with file upload)</p>\n</li>\n<li>\n<p><a href=\"https://pkg.go.dev/net/http\">Golang
  net/http Package</a></p>\n</li>\n</ul>\n<h2>Conclusion</h2>\n<p>That's it from this
  post of the series, a post on the POST method in golang :)</p>\n<p>We have covered
  topics like creating basic post requests, Marshalling golang types into JSON format,
  parsing form data, sending a POST request with files, and best practices for the
  POST method. Hope you found this article helpful. If you have any queries, questions,
  or feedback, please let me know in the comments or on my social handles.</p>\n<p>Happy
  Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/100-days-of-golang/golang-034-post-method.png
long_description: In this section of the series, we will be exploring how to send
  a  POST method is a type of request that is used to send data to a server(a machine
  on the internet). Imagine you are placing an order at a restaurant. With a GET request,
  it would be li
now: 2025-05-01 18:11:33.313704
path: blog/posts/2024-03-09-Golang-POST-Method.md
prevnext: null
series:
- 100-days-of-golang
slug: golang-web-post-method
status: published
tags:
- go
templateKey: blog-post
title: 'Golang Web: POST Method'
today: 2025-05-01
---

## Introduction

In this section of the series, we will be exploring how to send a `POST` HTTP request in golang. We will understand how to send a basic POST request, create an HTTP request, and parse json, structs into the request body, add headers, etc in the following sections of this post. We will understand how to marshal the golang struct/types into JSON format, send files in the request, and handle form data with examples of each in this article. Let's answer a few questions first.

## What is a POST request?

POST method is a type of request that is used to send data to a server(a machine on the internet).

Imagine you are placing an order at a restaurant. With a GET request, it would be like asking the waiter, "What kind of pizza do you have?" The waiter would respond by telling you the menu options (the information retrieved from the server).

However, a POST request is more like giving your completed order to the waiter. You tell them the specific pizza you want, its size, and any additional toppings (the data you send). The waiter then takes this information (POST request) back to the kitchen (the server) to process it (fulfill your order).

In the world of web development, POST requests are often used for things like:

Submitting forms (e.g., contact forms, login forms) Uploading files (e.g., photos, videos) Creating new accounts Sending data to be processed (e.g., online purchases)

Here's an example of what the POST request might look like in this scenario:

```nginx
POST /api/order HTTP/1.1
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
            "quantity": 2
        },
        {
            "itemID": 999,
            "name": "Burger",
            "quantity": 1
        }
    ]
}
```

In this example:

* The `POST` method is used to send data to the server.
* The `/api/order` is the endpoint of the server.
* The `application/json` is the content type of the request.
* The `123` is the content length of the request.
* The `{"userID": 123, "orderID": 456, "items": [{"itemID": 789, "name": "Pizza", "quantity": 2}, {"itemID": 999, "name": "Burger", "quantity": 1}]}` is the body of the request.


## Why the need for a POST request?

In the world of HTTP requests, we use the POST method to securely send data from a client (like a user's browser) to a server. This is crucial because the GET method, while convenient for retrieving data, has limitations:

Imagine you are in registering for an event via Google form, you type in your details on the webpage like name, email, address, phone number, and other personal details. If the website/app was using the `GET` method to send the request to register or do any other authentication/privacy-related requests, it could expose the data in the URL itself. It would be something along the lines [`https://form.google.com/register/<form-name>-<id>/?name=John&phone_number=1234567890`](https://form.google.com/register/%3Cform-name%3E-%3Cid%3E/?name=John&phone_number=1234567890), if a user maliciously sniffs into your network and inspects the URL, your data will be exposed. That is the reason we need `POST` a method.

## How a POST method works?

A [POST](https://www.rfc-editor.org/rfc/rfc9110#POST) request is used to send data to a server to create or update(there is a separate method for updating) a resource. The client(browser/other APIs) sends a POST request to the server's API endpoint with the data in the request body. This data can be in formats like JSON, XML, or form data. The server processes the POST request, validates and parses the data in the request body, makes any changes or creates resources based on that data, and returns a response. The response would contain a status code indicating the success or failure of the operation and may contain the newly created or updated resource in the response body. The client must check the response status code to verify the outcome and process the response accordingly. Unlike GET, POST can create new resources on the server. The body of a POST contains the data for creation while the URL identifies the resource to be created. Overall, POST transfers data to the server for processing, creation or updating of resources.

The status code is usually `201` indicating the resource is successfully created or `200` for just indicating success.

Some common steps for creating and sending a POST request as a developer include:

* Defining the API endpoint

* Clarifying the data format (json, language native objects, xml , text, form-data, etc)

* Converting / Marshalling the data

* Attaching header for `Content-Type` as key and value as the format of the data type (e.g. `application/json` for json)

* Sending the request


The above steps are general for creating and sending a POST request, they are not specific to Golang. For golang specific steps, we need to dive a bit deeper, let's get started.

## Basic POST method in Golang

To send a POST request in golang, we need to use the `http` package. The `http` package has the `Post` method, which takes in 3 parameters, namely the URL, the Content-Type, and the Body. The body can be `nil` if the URL endpoint doesn't necessarily require a body. The `Content-Type` is the string, since we are just touching on how the Post request is constructed, we will see what the `Content-Type` string value should be in the later sections.

> [http.Post](http://http.Post)(URL, Content-Type, Body)

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	apiURL := "https://reqres.in/api/users"

	// POST request
	resp, err := http.Post(apiURL, "", nil)
    // ideally the Content-Type header should be set to the relevant format
	// resp, err := http.Post(apiURL, "application/json", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
    fmt.Println(resp)
	defer resp.Body.Close()
}
```

```bash

$ go run main.go

201
&{
    201 Created
    201
    HTTP/2.0
    2
    0
    map[
        Access-Control-Allow-Origin:[*]
        Cf-Cache-Status:[DYNAMIC]
        Cf-Ray:[861cd9aec8223e4b-BOM]
        Content-Length:[50]
        Content-Type:[application/json; charset=utf-8]
        Date:[Sat, 09 Mar 2024 17:40:28 GMT]
        Server:[cloudflare]
        ...
        ...
        ...
        X-Powered-By:[Express]
    ]
    {0xc00017c180}
    50
    []
    false
    false
    map[]
    0xc000156000
    0xc00012a420
}
```

The above code is sending the `POST` request to the [`https://reqres.in/api/users`](https://reqres.in/api/users) endpoint with an empty body and no specific format for `Content-Type` header. The response is according to the [Response](https://pkg.go.dev/net/http#Response) structure. We can see we got `201` status, which indicates the server received the POST request successfully, the API is a dummy api, so we don't care about the data we are processing, we are just using the API as a placeholder for sending the POST request.

We can use `map[string]interface{}` it to pass the data in the request body. The `json.Marshal` method is used to convert the map into JSON format. We will look into the details shortly in the next few examples.

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
    apiURL := "https://reqres.in/api/users"
    bodyMap := map[string]interface{}{
        "name": "morpheus",
        "job": "leader",
    }

    requestBody, err := json.Marshal(bodyMap)
    if err != nil {
        panic(err)
    }
    body := bytes.NewBuffer(requestBody)

    resp, err := http.Post(apiURL, "application/json", body)
    if err != nil {
        panic(err)
    }
    fmt.Println(resp.StatusCode)
    defer resp.Body.Close()
}
```

```bash
$ go run main.go

201
```

The above code sends the `POST` request to the [`https://reqres.in/api/users`](https://reqres.in/api/users) endpoint with the data in the request body in JSON format.

## Creating a POST request in Golang

We can construct the POST request with the [NewRequest](https://pkg.go.dev/net/http#NewRequest) method. The method takes in 3 parameters, namely the `method` (e.g. `POST`, `GET`), the `URL` and the `body` (if there is any). We can then add extra information to the headers or the Request object after constructing the basic HTTP [Request](https://pkg.go.dev/net/http#Request) object.

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	apiURL := "https://reqres.in/api/users"

	req, err := http.NewRequest(http.MethodPost, apiURL, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	//fmt.Println(resp)
	defer resp.Body.Close()
}
```

```bash
$ go run main.go

201
```

In the above example, we have created an HTTP Request as the `POST` method, with [`https://reqres.in/api/users`](https://reqres.in/api/users) as the URL, and no body. This constructs an HTTP Request object, which can be sent as the parameter to the [http.DefaultClient.Do](http://http.DefaultClient.Do) method, which is the default client for the request we sent in the earlier examples as `http.Get` or [`http.Post`](http://http.Post) methods. We can implement a custom client as well, and then apply `Do` the method with the request parameters. The `Do` method returns the `Request` object or the `error` if any.

More on the customizing Client will be explained in a separate post in the series.

The response is also in the same format as the [Response](https://pkg.go.dev/net/http#Response) structure that we have seen earlier. This section of the series aims to construct a post request, and not to parse the response, we have already understood the parsing of the response in the [Get method](https://www.meetgor.com/golang-web-get-method/#?:~:text=Parsing%20the%20JSON%20body%20with%20structs) section of the series.

### Parsing objects to JSON for POST method request

We might have a golang object that we want to send as a body to an API in the POST request, for that we need to convert the golang struct object to JSON. We can do this by using the [Marshal](https://pkg.go.dev/encoding/json#Marshal) or the [Encode](https://pkg.go.dev/encoding/json#Encoder.Encode) method for serialization of the golang struct object to JSON.

#### Using Marshal method

Marshaling is the process of converting data from a data structure into a format suitable for transmission over a network or for storage. It's commonly used to convert native objects in a programming language into a serialized format, typically a byte stream, that can be transmitted or stored efficiently. You might get a question here, what is the difference between `Marshalling` and `Serialization`? Well, Serialization, is a broader term that encompasses marshalling. It refers to the process of converting an object or data structure into a format that can be stored or transmitted and later reconstructed into the original object. Serialization may involve converting data into byte streams, XML, JSON, or other formats. So, in summary, marshaling specifically deals with converting native objects into a format suitable for transmission, while serialization encompasses the broader process of preparing data for storage or transmission.

The `json` package has the [Marshal](https://pkg.go.dev/encoding/json#Marshal) method that converts the golang object into JSON. The `Marshal` method takes in a parameter as the struct object with type `any` and returns a byte slice `[]byte` and error (if any).

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Age    int    `json:"age"`
}

func main() {
	user := User{
		Name:   "Alice",
		Salary: 50000,
		Age:    25,
	}
	apiURL := "https://dummy.restapiexample.com/api/v1/create"

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
	resp, err := http.Post(apiURL, "application/json", body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
}
```

```bash
$ go run main.go

{"name":"Alice","salary":50000,"age":25}
200
```

In the above example, we have created a struct `User` with fields `Name`, `Salary`, and `Age`, the json tags will help label each key in JSON with the tag for the respective fields in the struct. We create an object `user` of a type `User` with the values as `Alice`, `50000`, and `25` respectively.

We call the `json.Marshal` method with the parameter `user` that represents the struct object `User`, the method returns a slice of bytes, or an error either or both could be nil. If we try to see the stringified representation of the byte slice, we can see something like `{"name":"Alice","salary":50000,"age":25}` which is a JSON string for the user struct. We can't parse the byte slice as the body in the POST request, we need the `io.Reader` object, so we can load the byte slice `bodyBytes` into a buffer and parse that as a body for the POST request.

We then send a `POST` request to the endpoint [`https://dummy.restapiexample.com/api/v1/create`](https://dummy.restapiexample.com/api/v1/create) with the content type as `application/json` and with the body as `body` which was a `io.Reader` object as an in-memory buffer.

In brief, we can summarize the marshaling of the golang object into JSON with `Marshal` function as the following steps:

* Defining the structure as per the request body

* Creating the struct object for parsing the data as body to the request

* Calling the `json.Marshal` function to convert the object to JSON (parameter as the struct object `any` type)

* Loading the byte slice into a buffer with `bytes.NewBuffer()`

* Sending the POST request to the endpoint with the body as the `io.Reader` object and content type as `application/json`


#### Using Encode method

We can even use the [Encoder.Encode](https://pkg.go.dev/encoding/json#Encoder.Encode) method to parse the golang struct object to JSON. Firstly, we should have the struct defined as per the request body that the particular API takes, we can make use of the json tags, omitempty, omit(-) options to make the marshaling process work accordingly. We can then create the object of that particular struct with the data we require to be created as a resource with the POST request on that API service.

Thereafter we can create an empty buffer object with [bytes.Buffer](https://pkg.go.dev/bytes#Buffer), this buffer object would be used to initialize the [Encoder](https://pkg.go.dev/encoding/json#Encoder) object with the [NewEncoder](https://pkg.go.dev/encoding/json#NewEncoder) method. This would give access to the [Encode](https://pkg.go.dev/encoding/json#Encoder.Encode) method, which is used to take in the struct object (`any` type) and this will populate the buffer we initialized with the `NewEncoder` method.

Later we can access that buffer to parse it to the Post request as the body. Let's understand it better with an example.

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name   string
	Salary int
	Age    int
}

func main() {
	user := User{
		Name:   "Alice",
		Salary: 50000,
		Age:    25,
	}
	apiURL := "https://dummy.restapiexample.com/api/v1/create"

	var bodyBuffer bytes.Buffer
	var encoder = json.NewEncoder(&bodyBuffer)
	err := encoder.Encode(user)
    if err != nil {
        panic(err)
    }

	resp, err := http.Post(apiURL, "application/json", &bodyBuffer)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp)
	defer resp.Body.Close()
}
```

Over here, we have created a struct `User` with fields `Name`, `Salary`, and `Age`, we initialize the `user` as the object of the `User` struct. Then we create a buffer `bodyBuffer` of type `bytes.Buffer` this is the actual buffer that we will send as the body. Further, we initialize the `Encoder` object as `encoder` with the `json.NewEncoder` method by parsing the reference of `bodyBuffer` as the parameter. Since `bytes.Buffer` implements the `io.Writer` interface, we can pass the `bodyBuffer` to the `NewEncoder` method. This will create the `Encoder` object which in turn will give us access to the `Encode` method, where we will parse the struct instance and it will populate the buffer with which we initialized the `Encoder` object earlier.

Now, we have the `encode` object, this gives us the access to `Encode` method, we call the `Encode` method with the parameter of `user` which is a User struct instance/object. The Encode method will populate the `bodyBuffer` object or it will result in an error if anything goes wrong (the data is incorrectly parsed or is not in the required format).

We can call the `Post` method with the initialized URL, the `Content-Type` as `application/json` since we have converted the struct instance to JSON object, and the body as the reference to the buffer as `&bodyBuffer`

So, the steps for parsing struct instances into JSON objects with the `Encoder.Encode` method is as follows:

* Defining the structure as per the request body

* Creating the struct object for parsing the data as body to the request

* Creating an empty `bytes.Buffer` object as an in-memory buffer

* Initializing the `Encoder` object with `NewEncoder` method by parsing the reference of `bodyBuffer` as the parameter

* Calling the `Encode` method with the parameter of struct instance/object

* Sending the POST request to the endpoint with the content type as `application/json` and body as the reference to the buffer

The results are the same as the above example just the way we have parsed the struct instance to JSON object is different.

### Parsing JSON to POST request

We have seen how we can parse golang struct instances to JSON and then send the post request, but what if we had the JSON string already with us, and we want to send the request? Well, that's much easier, right? We already have parsed the JSON string to the Post request by loading the slice of bytes into a buffer, so we just need to convert the string to a slice of bytes which is quite an easy task, and then load that byte slice to the buffer.

```go
package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	// dummy api
	apiURL := "https://dummy.restapiexample.com/api/v1/create"

	// json data
	data := `{
        "name": "Alice",
        "job": "Teacher"
    }`
	body := bytes.NewBuffer([]byte(data))

	// POST request
	resp, err := http.Post(apiURL, "application/json", body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp)
	defer resp.Body.Close()
}
```

In the example above, we already have a JSON string `data` with keys as `name` and `job` but it is not JSON, it is a stringified JSON. We can convert the stringified JSON to a slice of bytes using the `[]byte` function. Further, we have used the `bytes.NewBuffer` method to load the byte slice into an `io.Reader` object. This object returned by the `bytes.NewBuffer` will serve as the body for the POST request.

### Parsing JSON to objects in Golang from POST method response

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
	Status string `json:"status"`
	Data   User   `json:"data"`
}

func main() {
	user := User{
		Name:   "Alice",
		Salary: 50000,
		Age:    "25",
	}
	apiURL := "https://dummy.restapiexample.com/api/v1/create"

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
	resp, err := http.Post(apiURL, "application/json", body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp)
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// unmarshalling process
	// converting JSON to Go specific data structure/types
	var userResponse UserResponse
	if err := json.Unmarshal(respBody, &userResponse); err != nil {
		panic(err)
	}
	fmt.Println(userResponse)
	fmt.Println(userResponse.Data)
}
```

```nginx

{success {Alice 50000 25 3239}}
{Alice 50000 25 577}
```

The above example is a POST request with a struct instance being loaded as a JSON string and then sent as a buffer to the API endpoint, it also reads the response body with a specific structure `UserResponse` and unmarshalled the `resp.Body` from the `io.Reader` as `respBody` and then loads into `userResponse` object. This example gives an entire process of what we have understood in the JSON data parsing for a POST request.

### Sending Form data in a POST request

We can also send data to a POST request in the form of a form, the form which we use in the HTML. Golang has a `net/url` package to parse the form data. The form data is sent in the `application/x-www-form-urlencoded` format.

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ResponseLogin struct {
	Token string `json:"token"`
}

func main() {
	// dummy api
	apiURL := "https://reqres.in/api/login"

	// Define form data
	formData := url.Values{}
	formData.Set("email", "eve.holt@reqres.in")
	formData.Set("password", "cityslicka")

	// Encode the form data
	fmt.Println(formData.Encode())
	reqBody := strings.NewReader(formData.Encode())
	fmt.Println(reqBody)

	// Make a POST request with form data
	resp, err := http.Post(apiURL, "application/x-www-form-urlencoded", reqBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print response status code
	fmt.Println("Status Code:", resp.StatusCode)

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
	token := ResponseLogin{}

	json.Unmarshal(respBody, &token)
	fmt.Println(token)
}
```

```bash
$ go run main.go

email=eve.holt%40reqres.in&password=cityslicka
&{email=eve.holt%40reqres.in&password=cityslicka 0 -1}
Status Code: 200
{QpwL5tke4Pnpja7X4}
```

In the above example, we set a `formData` with the values of `email` and `password` which are `url.Values` object. The `url.Values` the object is used to store the key-value pairs of the form data. The `formData` is encoded with the `url.Encode` method, We load the encoded string to a `buffer` with `strings.NewReader` which implements the `io.Reader` interface, so that way we can pass that object as the body to the post request.

We send the `POST` request to the endpoint [`https://reqres.in/api/login`](https://reqres.in/api/login) with the content type as `application/x-www-form-urlencoded` and with the body as `reqBody` which implements the `io.Reader` interface as an in-memory buffer. The response from the request is read into the buffer with `io.ReadAll` method and we can `Unmarshal` the stream of bytes as a buffer into the `ResponseLogin` struct object.

The output shows the `formData` as encoded string `email=eve.holt%`[`40reqres.in`](http://40reqres.in)`&password=cityslicka` as `@` is encoded to `%40`, then we wrap the `formData` in a `strings.NewReader` object which is a buffer that implements `io.Reader` interface, hence we can see the result as the object. The status code for the request is `200` indicating the server received the `form-data` in the body and upon unmarshalling, we get the token as a response to the POST request which was a dummy login API.

This way we have parsed the form-data to the body of a POST request.

### Sending File in a POST request

We have covered, parsing text, JSON, and form data, and now we need to move into sending files in a POST request. We can use the `multipart` package to parse files into the request body and set appropriate headers for reading the file from the API services.

We first read the file contents [`os.Open`](http://os.Open) which returns a reference to the `file` object or an error. We create an empty `bytes.Buffer` object as `body` which will be populated later. The [multipart.NewWriter](https://pkg.go.dev/mime/multipart#NewWriter) method takes in the `io.Writer` object which will be the `body` as it is an `bytes.Buffer` object that implements the `io.Writer` interface. This will initialize the [Writer](https://pkg.go.dev/mime/multipart#Writer) object in the `multipart` package.

We create a `form-field` in the `Writer` object with the [CreateFormFile](https://pkg.go.dev/mime/multipart#Writer.CreateFormFile) method, which takes in the `fieldName` as the name of the field, and the `fileName` as the name of the file which will be read later in the multipart form. The method returns either the part or the error. The `part` is an object that implements the `io.Writer` interface.

Since we have stored the file contents in the `file` object, we copy the contents into the `form-field` with the [Copy](https://pkg.go.dev/io#Copy) method. Since the `part` return from the `CreateFormFile` was implementing the `io.Writer` interface, we can use it to Copy the contents from source to destination. The source is the `io.Reader` object and the destination is the `io.Writer` object, the destination for the `Copy` method is the first parameter, the source is the second parameter.

This Copy method will populate the buffer initialized earlier in the `NewWriter` method. This will give us a buffer that has the file contents in it. We can pass this buffer to the POST request with the `body` parameter. We also need to make sure we close the `Writer` object after copying the contents of the file. We can extract the type of file which will serve as the `Content-Type` of the request.

Let's clear the explanation with an example.

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type ResponseFile struct {
	Files map[string]string `json:"files"`
}

func main() {
	apiURL := "http://postman-echo.com/post"
	fileName := "sample.csv"

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("csvFile", fileName)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

    contentType := writer.FormDataContentType()
    fmt.Println(contentType)

	writer.Close()

	resp, err := http.Post(apiURL, contentType, body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Status Code:", resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	token := ResponseFile{}
	json.Unmarshal(respBody, &token)
	fmt.Println(token)
	fmt.Println(token.Files[fileName])
}
```

```bash
multipart/form-data; boundary=7e0eacfff890be395eba19c70415c908124b503a56f23ebeec0ab3c665ca


--619671ea2c0aa47ca6664a7cda422169d73f3b8a089c659203f5413d03de
Content-Disposition: form-data; name="csvFile"; filename="sample.csv"
Content-Type: application/octet-stream

User,City,Age,Country
Alex Smith,Los Angeles,20,USA
John Doe,New York,30,USA
Jane Smith,Paris,25,France
Bob Johnson,London,40,UK

--619671ea2c0aa47ca6664a7cda422169d73f3b8a089c659203f5413d03de--



Status Code: 200

{map[sample.csv:data:application/octet-stream;base64,VXNlcixDaXR5LEFnZSxDb3VudHJ5CkFsZXggU21pdGgsTG9zIEFuZ2VsZXMsMjAsVVNBCkpvaG4gRG9lLE5ldyBZb3JrLDMwLFVTQQpKYW5lIFNtaXRoLFBhmlzLDI1LEZyYW5jZQpCb2IgSm9obnNvbixMb25kb24sNDAsVUsK]}

data:application/octet-stream;base64,VXNlcixDaXR5LEFnZSxDb3VudHJ5CkFsZXggU21pdGgsTG9zIEFuZ2VsZXMsMjAsVVNBCkpvaG4gRG9lLE5ldyBZb3JrLDMwLFVTQQpKYW5lIFNtaXRoLFBhmlzLDI1LEZyYW5jZQpCb2IgSm9obnNvbixMb25kb24sNDAsVUsK
```

In the above example, we first read the file `sample.csv` into the `file` object with [`os.Open`](http://os.Open) method, this will return a reference to the file object or return an error if any arises while opening the file.

Then we create an empty buffer `bytes.Buffer` object which will serve as the body of the post request later as it will get populated with the file contents in the form of `multipart/form-data`.

We initialize the `Writer` object with `multipart.NewWriter` method which takes in the empty buffer as the parameter, we parse the `body` as the parameter. The method will return a reference to the `multipart.Writer` object.

With the `Writer` object we access the `CreateFormFile` method which takes in the `fieldName` as the name of the field, and the `fileName` as the name of the file. The method will return either the part or an error. The `part` in this case, is the reference to the `io.Writer` object that will be used to write the contents from the uploaded file.

Then, we can use the `io.Copy` method to copy the contents from the `io.Reader` object to the `io.Writer` object. The source is the `io.Reader` object and the destination is the `io.Writer` object. The first parameter is however the destination and the second parameter is the source. In the example, we call `io.Copy(part, file)` which will copy the contents of `file` to the `part` buffer.

We get the `Content-Type` by calling the [Writer.FormDataContentType](https://pkg.go.dev/mime/multipart#Writer.FormDataContentType) method. This returns us `multipart/form-data; boundary=7e0eacfff890be395eba19c70415c908124b503a56f23ebeec0ab3c665ca` which will serve the `Content-Type` for the Post request.

We need to make sure we close the `Writer` object with the `Close` method.

We just print the `body.String()` to get a look at what the actual body looks like, we can see there is a form for the file as a `form-data` with keys like `Content-Type`, `Content-Disposition`, etc. The file has the `Content-Type` as `application/octet-stream` and the actual content is rendered in the output.

The dummy API responds with a 200 status code and also sends the JSON data with the name of the file as the key and the value as the `base64` encoded value of the file contents. This indicates that we were able to upload the file to the server API using a POST request. Well done!

I have also included some more examples of POST requests with files [here](https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/post/file_2.go) which extends the above example by taking the encoded values and decoding to get the actual contents of the file back.

## Best Practices for POST method

Here are some of the best practices for the POST method which are followed to make sure you consume or create the POST request in the most secure, efficient, and graceful way.

### Always Close the Response Body

Ensure that you close the response body after reading from it. Use `defer response.Body.Close()` to automatically close the body when the surrounding function returns. This is crucial for releasing associated resources like network connections or file descriptors. Failure to close the response body can lead to memory leaks, particularly with a large volume of requests. Properly closing the body prevents resource exhaustion and maintains efficient memory usage.

### Client Customization

Utilize the [Client](https://pkg.go.dev/net/http#Client) struct to customize the HTTP client behavior. By using a custom client, you can set timeouts, headers, user agents, and other configurations without modifying the `DefaultClient` provided by the `http` package. This approach allows for flexibility and avoids repetitive adjustments to the client configuration for each request.

### Set Content-Type Appropriately

Ensure that you set the `Content-Type` header according to the request payload. Correctly specifying the Content-Type is crucial for the server to interpret the request payload correctly. Failing to set the Content-Type header accurately may result in the server rejecting the request. Always verify and match the Content-Type header with the content being sent in the POST request to ensure smooth communication with the server.

## Reference

* [Postman POST API](https://www.postman.com/postman/workspace/postman-answers/documentation/13455110-00378d5c-5b08-4813-98da-bc47a2e6021d) (For POST request with file upload)

* [Golang net/http Package](https://pkg.go.dev/net/http)


## Conclusion

That's it from this post of the series, a post on the POST method in golang :)

We have covered topics like creating basic post requests, Marshalling golang types into JSON format, parsing form data, sending a POST request with files, and best practices for the POST method. Hope you found this article helpful. If you have any queries, questions, or feedback, please let me know in the comments or on my social handles.

Happy Coding :)