---
article_html: "<h2>Introduction</h2>\n<p>We have explored GET, POST, PUT, and PATCH
  methods in our previous entries for this series. It is the final entry for all the
  HTTP methods which is going to be the <code>DELETE</code> method. In this entry
  of the series, we will take a view on how to construct and request an HTTP DELETE
  Method to an API Endpoint.</p>\n<p>The DELETE Method is quite simple. You just provide
  the URI of the resource. Most often, the request body is not needed. That request
  will simply delete the request entity from the server (in the database or wherever
  the resource is stored).</p>\n<p>Let’s understand it in a more detailed way.</p>\n<h2>What
  is a DELETE Method</h2>\n<p>The DELETE method requests that the server remove the
  association between the target resource and its URI (Uniform Resource Identifier).
  This doesn't necessarily mean the underlying data is physically deleted; it means
  the resource is no longer accessible through that specific URL. DELETE can also
  be used to remove relationships between resources, effectively &quot;delinking&quot;
  them.</p>\n<p>According to the RFC:</p>\n<blockquote>\n<p>The DELETE method requests
  that the origin server remove the association between the target resource and its
  current functionality.</p>\n</blockquote>\n<p>Examples:</p>\n<ul>\n<li>\n<p><strong>Social
  Media</strong> (Deleting a Tweet): When you delete a tweet, you're sending a DELETE
  request to the server. This removes the tweet from your timeline and makes it inaccessible
  via its URL. While the data might be archived or retained for a period, the key
  action is removing the public association between the tweet and its online presence.
  This is closer to a true deletion than the cart example.</p>\n</li>\n<li>\n<p><strong>E-Commerce</strong>
  (Removing an Item from a Cart): When you remove an item from your online shopping
  cart, you're sending a request (often a DELETE) to remove the item from your cart.
  The actual product remains available in the store's inventory. This is a clear example
  of delinking. You're deleting the link between your cart and the product, not the
  product itself.</p>\n</li>\n</ul>\n<p>Let’s start constructing a simple DELETE Request
  in Golang.</p>\n<h2>A Simple DELETE Request</h2>\n<p>We don’t have a specific method
  for <code>DELETE</code> in <code>net/http</code> as we have for <code>GET</code>
  and <code>POST</code>, so we need to create a request and use a client to send the
  request.</p>\n<h3>Constructing the URL</h3>\n<p>We would need to define the endpoint
  that we are hitting. We can directly use the API URL or construct the API URL on
  the fly, depending on the ID and dynamic parameters. DELETE requests usually delete
  a particular entity. We would generally have some form of identifier for that entity/object
  on the database, etc. So, in this case, it is the user's ID, so we can pass the
  post to the URL.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\">// define URL to hit the API</span><span class=\"w\"></span>\n<span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in/api/users/4&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"c1\">// OR</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">baseURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in&quot;</span><span
  class=\"w\"></span>\n<span class=\"nx\">userID</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">4</span><span
  class=\"w\"></span>\n<span class=\"nx\">apiURL</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Sprintf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s/api/users/%d&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">baseURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">userID</span><span class=\"p\">)</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>We
  can either directly define the URL or dynamically construct the URL, that is quite
  straightforward. The latter one is the one we usually use and design.</p>\n<p>The
  DELETE Request doesn’t usually require a request body, however, if your server requires
  some specifications, you can construct the body as we did with the previous examples
  in POST, PUT, or PATCH method requests.</p>\n<h3>Constructing and sending the DELETE
  Request</h3>\n<p>We can simply construct the request by specifying the http.MethodDelete
  as the request method, the URL to hit, and a body(optional) just like a <code>GET</code>
  request. Once we have the request, we can create the default client and send the
  request.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"c1\">// create a DELETE request</span><span class=\"w\"></span>\n<span class=\"nx\">req</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">MethodDelete</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"c1\">// construct the default http client and send the request</span><span
  class=\"w\"></span>\n<span class=\"nx\">client</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">Client</span><span
  class=\"p\">{}</span><span class=\"w\"></span>\n<span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>This
  is the normal code used for constructing an HTTP request in Golang, we create a
  request using the NewRequest function that takes in the method type, the URL to
  send the request, and the body if any. Then we need to create a http.Client for
  sending the request, we usually create a client with default values and send the
  request using the Do method on the created client using the request that we constructed
  earlier.</p>\n<h3>Fetching the Response</h3>\n<p>Once the request is sent, we can
  fetch the response and read the body as bytes, and check the status if that succeeded
  or failed.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Response Status:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">Status</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">io</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Response Body:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>We
  can grab the Status field for checking the status code and message for the request.
  Usually, the body would be empty since there is no resource we are expecting after
  deletion of the object. However, if the server is implemented in a way to return
  the deleted object, you can read the bytes of the body and unmarshal it to the desired
  struct.</p>\n<p>So, that is the entire code to create a simple Delete request with
  Go, simply construct the URL with the identifier for the resource to be deleted,
  create the request, and send the request, and if the status code is 204 (usually)
  then we can assume it succeeded.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">baseURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">userID</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">2</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">apiURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Sprintf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s/api/users/%d&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">baseURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">userID</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">MethodDelete</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">return</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">client</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">Client</span><span
  class=\"p\">{}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">client</span><span class=\"p\">.</span><span
  class=\"nx\">Do</span><span class=\"p\">(</span><span class=\"nx\">req</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">return</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Response Status:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Status</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">io</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">return</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Response Body:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nb\">string</span><span class=\"p\">(</span><span class=\"nx\">respBody</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<h2>Facts about the DELETE Method</h2>\n<ul>\n<li>\n<p>DELETE
  Method is idempotent: Similar requests will result in identical behavior or response
  since once the resource is deleted, the resource won’t exist and hence the behavior
  would not hinder any other parts.</p>\n</li>\n<li>\n<p>DELETE Method is not safe:
  The operation is quite dangerous as it could literally remove a resource from a
  database/storage. Hence it is called not safe as it is making changes on the server.</p>\n</li>\n</ul>\n<p>I
  have also included some more examples of DELETE requests <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/delete/\"><strong>here</strong></a>.</p>\n<p>That's
  it from the 36th part of the series, all the source code for the examples are linked
  in the GitHub on the <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/delete/\"><strong>100
  days of Golang</strong></a> repository.</p>\n<p><a href=\"https://github.com/Mr-Destructive/100-days-of-golang\"><strong>100-days-of-golang</strong></a></p>\n<h2>Conclusion</h2>\n<p>That
  would be it from the DELETE Method in Golang. We can use this method just like a
  normal <code>GET</code> request however a bit more carefully.</p>\n<p>Hope you found
  this article, helpful, leave some feedback or any suggestions if you have any. Thank
  you for reading.</p>\n<p>Happy Coding :)</p>\n"
cover: ''
date: 2025-01-12
datetime: 2025-01-12 00:00:00+00:00
description: Exploring the fundamentals of a DELETE method in golang. How to request
  a resource, parse body, headers, etc. in a HTTP DELETE request.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2025-01-12-GO-37-DELETE-Method.md
html: "<h2>Introduction</h2>\n<p>We have explored GET, POST, PUT, and PATCH methods
  in our previous entries for this series. It is the final entry for all the HTTP
  methods which is going to be the <code>DELETE</code> method. In this entry of the
  series, we will take a view on how to construct and request an HTTP DELETE Method
  to an API Endpoint.</p>\n<p>The DELETE Method is quite simple. You just provide
  the URI of the resource. Most often, the request body is not needed. That request
  will simply delete the request entity from the server (in the database or wherever
  the resource is stored).</p>\n<p>Let’s understand it in a more detailed way.</p>\n<h2>What
  is a DELETE Method</h2>\n<p>The DELETE method requests that the server remove the
  association between the target resource and its URI (Uniform Resource Identifier).
  This doesn't necessarily mean the underlying data is physically deleted; it means
  the resource is no longer accessible through that specific URL. DELETE can also
  be used to remove relationships between resources, effectively &quot;delinking&quot;
  them.</p>\n<p>According to the RFC:</p>\n<blockquote>\n<p>The DELETE method requests
  that the origin server remove the association between the target resource and its
  current functionality.</p>\n</blockquote>\n<p>Examples:</p>\n<ul>\n<li>\n<p><strong>Social
  Media</strong> (Deleting a Tweet): When you delete a tweet, you're sending a DELETE
  request to the server. This removes the tweet from your timeline and makes it inaccessible
  via its URL. While the data might be archived or retained for a period, the key
  action is removing the public association between the tweet and its online presence.
  This is closer to a true deletion than the cart example.</p>\n</li>\n<li>\n<p><strong>E-Commerce</strong>
  (Removing an Item from a Cart): When you remove an item from your online shopping
  cart, you're sending a request (often a DELETE) to remove the item from your cart.
  The actual product remains available in the store's inventory. This is a clear example
  of delinking. You're deleting the link between your cart and the product, not the
  product itself.</p>\n</li>\n</ul>\n<p>Let’s start constructing a simple DELETE Request
  in Golang.</p>\n<h2>A Simple DELETE Request</h2>\n<p>We don’t have a specific method
  for <code>DELETE</code> in <code>net/http</code> as we have for <code>GET</code>
  and <code>POST</code>, so we need to create a request and use a client to send the
  request.</p>\n<h3>Constructing the URL</h3>\n<p>We would need to define the endpoint
  that we are hitting. We can directly use the API URL or construct the API URL on
  the fly, depending on the ID and dynamic parameters. DELETE requests usually delete
  a particular entity. We would generally have some form of identifier for that entity/object
  on the database, etc. So, in this case, it is the user's ID, so we can pass the
  post to the URL.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\">// define URL to hit the API</span><span class=\"w\"></span>\n<span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in/api/users/4&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"c1\">// OR</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">baseURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in&quot;</span><span
  class=\"w\"></span>\n<span class=\"nx\">userID</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">4</span><span
  class=\"w\"></span>\n<span class=\"nx\">apiURL</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Sprintf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s/api/users/%d&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">baseURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">userID</span><span class=\"p\">)</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>We
  can either directly define the URL or dynamically construct the URL, that is quite
  straightforward. The latter one is the one we usually use and design.</p>\n<p>The
  DELETE Request doesn’t usually require a request body, however, if your server requires
  some specifications, you can construct the body as we did with the previous examples
  in POST, PUT, or PATCH method requests.</p>\n<h3>Constructing and sending the DELETE
  Request</h3>\n<p>We can simply construct the request by specifying the http.MethodDelete
  as the request method, the URL to hit, and a body(optional) just like a <code>GET</code>
  request. Once we have the request, we can create the default client and send the
  request.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"c1\">// create a DELETE request</span><span class=\"w\"></span>\n<span class=\"nx\">req</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">MethodDelete</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"c1\">// construct the default http client and send the request</span><span
  class=\"w\"></span>\n<span class=\"nx\">client</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">Client</span><span
  class=\"p\">{}</span><span class=\"w\"></span>\n<span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>This
  is the normal code used for constructing an HTTP request in Golang, we create a
  request using the NewRequest function that takes in the method type, the URL to
  send the request, and the body if any. Then we need to create a http.Client for
  sending the request, we usually create a client with default values and send the
  request using the Do method on the created client using the request that we constructed
  earlier.</p>\n<h3>Fetching the Response</h3>\n<p>Once the request is sent, we can
  fetch the response and read the body as bytes, and check the status if that succeeded
  or failed.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Response Status:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">Status</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">io</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Response Body:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>We
  can grab the Status field for checking the status code and message for the request.
  Usually, the body would be empty since there is no resource we are expecting after
  deletion of the object. However, if the server is implemented in a way to return
  the deleted object, you can read the bytes of the body and unmarshal it to the desired
  struct.</p>\n<p>So, that is the entire code to create a simple Delete request with
  Go, simply construct the URL with the identifier for the resource to be deleted,
  create the request, and send the request, and if the status code is 204 (usually)
  then we can assume it succeeded.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">baseURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://reqres.in&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">userID</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">2</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">apiURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Sprintf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s/api/users/%d&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">baseURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">userID</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">MethodDelete</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">return</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">client</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">Client</span><span
  class=\"p\">{}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">client</span><span class=\"p\">.</span><span
  class=\"nx\">Do</span><span class=\"p\">(</span><span class=\"nx\">req</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">return</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Response Status:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Status</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">io</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"k\">return</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Response Body:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nb\">string</span><span class=\"p\">(</span><span class=\"nx\">respBody</span><span
  class=\"p\">))</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<h2>Facts about the DELETE Method</h2>\n<ul>\n<li>\n<p>DELETE
  Method is idempotent: Similar requests will result in identical behavior or response
  since once the resource is deleted, the resource won’t exist and hence the behavior
  would not hinder any other parts.</p>\n</li>\n<li>\n<p>DELETE Method is not safe:
  The operation is quite dangerous as it could literally remove a resource from a
  database/storage. Hence it is called not safe as it is making changes on the server.</p>\n</li>\n</ul>\n<p>I
  have also included some more examples of DELETE requests <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/delete/\"><strong>here</strong></a>.</p>\n<p>That's
  it from the 36th part of the series, all the source code for the examples are linked
  in the GitHub on the <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/delete/\"><strong>100
  days of Golang</strong></a> repository.</p>\n<p><a href=\"https://github.com/Mr-Destructive/100-days-of-golang\"><strong>100-days-of-golang</strong></a></p>\n<h2>Conclusion</h2>\n<p>That
  would be it from the DELETE Method in Golang. We can use this method just like a
  normal <code>GET</code> request however a bit more carefully.</p>\n<p>Hope you found
  this article, helpful, leave some feedback or any suggestions if you have any. Thank
  you for reading.</p>\n<p>Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/100-days-of-golang/golang-037-delete-method.png
long_description: We have explored GET, POST, PUT, and PATCH methods in our previous
  entries for this series. It is the final entry for all the HTTP methods which is
  going to be the  The DELETE Method is quite simple. You just provide the URI of
  the resource. Most oft
now: 2025-05-01 18:11:33.314830
path: blog/posts/2025-01-12-GO-37-DELETE-Method.md
prevnext: null
series:
- 100-days-of-golang
slug: golang-web-delete-method
status: published
tags:
- go
templateKey: blog-post
title: 'Golang Web: DELETE Method'
today: 2025-05-01
---

## Introduction

We have explored GET, POST, PUT, and PATCH methods in our previous entries for this series. It is the final entry for all the HTTP methods which is going to be the `DELETE` method. In this entry of the series, we will take a view on how to construct and request an HTTP DELETE Method to an API Endpoint.

The DELETE Method is quite simple. You just provide the URI of the resource. Most often, the request body is not needed. That request will simply delete the request entity from the server (in the database or wherever the resource is stored).

Let’s understand it in a more detailed way.

## What is a DELETE Method

The DELETE method requests that the server remove the association between the target resource and its URI (Uniform Resource Identifier). This doesn't necessarily mean the underlying data is physically deleted; it means the resource is no longer accessible through that specific URL. DELETE can also be used to remove relationships between resources, effectively "delinking" them.

According to the RFC:

> The DELETE method requests that the origin server remove the association between the target resource and its current functionality.

Examples:

* **Social Media** (Deleting a Tweet): When you delete a tweet, you're sending a DELETE request to the server. This removes the tweet from your timeline and makes it inaccessible via its URL. While the data might be archived or retained for a period, the key action is removing the public association between the tweet and its online presence. This is closer to a true deletion than the cart example.
    
* **E-Commerce** (Removing an Item from a Cart): When you remove an item from your online shopping cart, you're sending a request (often a DELETE) to remove the item from your cart. The actual product remains available in the store's inventory. This is a clear example of delinking. You're deleting the link between your cart and the product, not the product itself.
    

Let’s start constructing a simple DELETE Request in Golang.

## A Simple DELETE Request

We don’t have a specific method for `DELETE` in `net/http` as we have for `GET` and `POST`, so we need to create a request and use a client to send the request.

### Constructing the URL

We would need to define the endpoint that we are hitting. We can directly use the API URL or construct the API URL on the fly, depending on the ID and dynamic parameters. DELETE requests usually delete a particular entity. We would generally have some form of identifier for that entity/object on the database, etc. So, in this case, it is the user's ID, so we can pass the post to the URL.

```go
// define URL to hit the API
apiURL := "https://reqres.in/api/users/4"

// OR

baseURL := "https://reqres.in"
userID := 4
apiURL := fmt.Sprintf("%s/api/users/%d", baseURL, userID)
```

We can either directly define the URL or dynamically construct the URL, that is quite straightforward. The latter one is the one we usually use and design.

The DELETE Request doesn’t usually require a request body, however, if your server requires some specifications, you can construct the body as we did with the previous examples in POST, PUT, or PATCH method requests.

### Constructing and sending the DELETE Request

We can simply construct the request by specifying the http.MethodDelete as the request method, the URL to hit, and a body(optional) just like a `GET` request. Once we have the request, we can create the default client and send the request.

```go
// create a DELETE request
req, err := http.NewRequest(http.MethodDelete, apiURL, nil)
if err != nil {
	log.Fatal(err)
}

// construct the default http client and send the request
client := &http.Client{}
resp, err := client.Do(req)
if err != nil {
	log.Fatal(err)
}
```

This is the normal code used for constructing an HTTP request in Golang, we create a request using the NewRequest function that takes in the method type, the URL to send the request, and the body if any. Then we need to create a http.Client for sending the request, we usually create a client with default values and send the request using the Do method on the created client using the request that we constructed earlier.

### Fetching the Response

Once the request is sent, we can fetch the response and read the body as bytes, and check the status if that succeeded or failed.

```go
fmt.Println("Response Status:", resp.Status)
respBody, err := io.ReadAll(resp.Body)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Response Body:", string(respBody))
```

We can grab the Status field for checking the status code and message for the request. Usually, the body would be empty since there is no resource we are expecting after deletion of the object. However, if the server is implemented in a way to return the deleted object, you can read the bytes of the body and unmarshal it to the desired struct.

So, that is the entire code to create a simple Delete request with Go, simply construct the URL with the identifier for the resource to be deleted, create the request, and send the request, and if the status code is 204 (usually) then we can assume it succeeded.

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	baseURL := "https://reqres.in"
	userID := 2
	apiURL := fmt.Sprintf("%s/api/users/%d", baseURL, userID)

	req, err := http.NewRequest(http.MethodDelete, apiURL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response Status:", resp.Status)
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Response Body:", string(respBody))
}
```

## Facts about the DELETE Method

* DELETE Method is idempotent: Similar requests will result in identical behavior or response since once the resource is deleted, the resource won’t exist and hence the behavior would not hinder any other parts.
    
* DELETE Method is not safe: The operation is quite dangerous as it could literally remove a resource from a database/storage. Hence it is called not safe as it is making changes on the server.
    

I have also included some more examples of DELETE requests [**here**](https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/delete/).

That's it from the 36th part of the series, all the source code for the examples are linked in the GitHub on the [**100 days of Golang**](https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/delete/) repository.

[**100-days-of-golang**](https://github.com/Mr-Destructive/100-days-of-golang)

## Conclusion

That would be it from the DELETE Method in Golang. We can use this method just like a normal `GET` request however a bit more carefully.

Hope you found this article, helpful, leave some feedback or any suggestions if you have any. Thank you for reading.

Happy Coding :)