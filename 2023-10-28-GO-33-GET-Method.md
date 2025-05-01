---
article_html: "<h2>Introduction</h2>\n<p>In this section of the series, we will be
  exploring how to send a <code>GET</code> HTTP request in golang. We will be understanding
  how to send a basic GET request, create an HTTP request and customize the client,
  add headers, read the response body, etc in the following sections of this post.</p>\n<h2>What
  is a GET method?</h2>\n<p>A <a href=\"https://en.wikipedia.org/wiki/HTTP#Request_methods\">GET</a>
  method in the context of an HTTP request is an action that is used to obtain data/resources.
  The <code>GET</code> method is used in a web application to get a resource like
  an HTML page, image, video, media, etc.</p>\n<p>Some common usecases of the <code>GET</code>
  method are:</p>\n<ul>\n<li>Loading a webpage</li>\n<li>Getting an image, file or
  other resource</li>\n<li>API requests to retrieve data</li>\n<li>Search queries
  sending filters and parameters</li>\n</ul>\n<h2>Basic GET Request</h2>\n<p>To use
  the HTTP method <code>GET</code> in golang, the <a href=\"https://pkg.go.dev/net/http\">net/http</a>
  package has a <a href=\"https://pkg.go.dev/net/http#Get\">Get</a> method. This method
  simply takes in a URL as a string and returns the <a href=\"https://pkg.go.dev/net/http#Response\">response</a>
  or an error. Let's look at how to send a basic HTTP GET request in golang.</p>\n<pre
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
  class=\"c1\">// web/methods/get/main.go</span><span class=\"w\"></span>\n\n\n<span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">reqURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://www.google.com&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">reqURL</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Status</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Status Code:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run main.go\n\n&amp;{200
  OK 200 HTTP/2.0 2 0 map[Alt-Svc:[h3=&quot;:443&quot;; ma=2592000,h3-29=&quot;:443&quot;;
  ma=2592000] Cache-Control:[private, max-age=0] Content-Security-Policy-Report-Only:[object-src
  &#39;none&#39;;base-uri &#39;self&#39;;script-src &#39;nonce-pdz_s8Gr0owwMbX8I9qNEQ&#39;
  &#39;strict-dynamic&#39; &#39;report-sample&#39; &#39;unsafe-eval&#39; &#39;unsafe-inline&#39;
  https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp] Content-Type:[text/html;
  charset=ISO-8859-1] Date:[Fri, 27 Oct 2023 09:37:04 GMT] Expires:[-1] P3p:[CP=&quot;This
  is not a P3P policy! See g.co/p3phelp for more info.&quot;] Server:[gws] Set-Cookie:[1P_JAR=2023-10-27-09;
  expires=Sun, 26-Nov-2023 09:37:04 GMT; path=/; domain=.google.com; Secure AEC=Ackid1Q5FARA_9d7f7znggUdw6DoJA1DBpI17Z0SWxN519Dc64EqmYVHlFg;
  expires=Wed, 24-Apr-2024 09:37:04 GMT; path=/; domain=.google.com; Secure; HttpOnly;
  SameSite=lax NID=511=EToBPqckCVRE7Paehug1PgNBKqe7lFLI9d12xJrGbvP9r8tkFIRWciry3gsy8FZ8OUIK4gE4PD-irgNzg4Y1fVePLdyu0AJdY_HcqL6zQYok-Adn-y5TDPmMCNuDnrouBfoxtqVjYY_4RFOe8jalkYto5fQAwzWnNJyw8K0avsw;
  expires=Sat, 27-Apr-2024 09:37:04 GMT; path=/; domain=.google.com; HttpOnly] X-Frame-Options:[SAMEORIGIN]
  X-Xss-Protection:[0]] 0xc000197920 -1 [] false true map[] 0xc0000ee000 0xc0000d8420}\n\n200
  OK\n\nStatus Code: 200\n</pre></div>\n\n</pre>\n\n<p>In the above code, we have
  defined a URL string as <code>reqURL</code> and used the <a href=\"https://pkg.go.dev/net/http#Get\">Get</a>
  method to send a GET request. The <code>Get</code> is parsed with the <code>reqURL</code>
  string and the return values are stored as <code>resp</code> and <code>err</code>.
  We have added an error check after calling the <code>Get</code> method in order
  to avoid errors later in the code.</p>\n<p>The <code>Get</code> method as seen from
  the output has returned a <code>*http.Response</code> object, we can use the <code>Status</code>
  and <code>StatusCode</code> properties to get the status code of the response. In
  this case, the status code of the response was <code>200</code>. The response <code>resp</code>
  is an object of type <code>http.Response</code> i.e. it has fields like <code>Body</code>,
  <code>StatusCode</code>, <code>Headers</code>, <code>Proto</code>, etc. We can get
  each individual field from the <code>resp</code> object. We will later look into
  how to read the <code>Body</code> field from the response, it is not directly read
  as a string nor it is stored in other forms, rather it is streamed from the requested
  URL.</p>\n<h2>Creating a GET request</h2>\n<p>We can even construct a GET request
  using the <a href=\"https://pkg.go.dev/net/http#NewRequest\">NewRequest</a> method.
  This is a low-level way of creating a <code>GET</code> request. We mention the <code>method</code>,
  <code>URL</code>, and the <code>body</code>, in the case of <code>GET</code> request,
  there is nobody. So, the <code>NewRequest</code> is a general way of creating a
  <code>http</code> request.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\">// web/methods/get/newreq.go</span><span class=\"w\"></span>\n\n<span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">reqURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://www.google.com&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">MethodGet</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">reqURL</span><span
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
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>As
  we can see, we construct a <code>GET</code> request using the <code>NewRequest</code>
  method and then use the <a href=\"https://pkg.go.dev/net/http#Client.Do\">Do</a>
  method to send the request to the server. The <a href=\"https://pkg.go.dev/net/http#DefaultClient\">http.DefaultClient</a>
  is used as a client to send the request, if we want to customize this we can create
  a new instance object of <a href=\"https://pkg.go.dev/net/http#Client\">http.Client</a>
  and use it to send requests. We will be taking a look at clients in another part
  of this series when we want to persist a connection or avoid connecting multiple
  times to the same application/URL.</p>\n<p>For now, we will go ahead with the DefaultClient.
  This will trigger the request, in this case, a <code>GET</code> request to the specified
  URL in the <code>reqURL</code> string. The <code>Do</code> method returns either
  a <code>http.Response</code> or an <code>error</code> just like the <code>Get</code>
  method did.</p>\n<h2>Reading the Response Body</h2>\n<p>We saw some different ways
  to send a <code>GET</code> request, now the below example will demonstrate how to
  read the body of the response. The response body is read from a buffer rather than
  loading the entire response into memory. It makes it flexible to parse the data
  efficiently and as per the needs. We will see how we use the <a href=\"https://pkg.go.dev/io\">io</a>
  package's <a href=\"https://pkg.go.dev/io#ReadAll\">ReadAll</a> method can be used
  to read from the response buffer.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\">// web/methods/get/body.go</span><span class=\"w\"></span>\n\n<span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">reqURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://httpbin.org/html&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">reqURL</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// close the body object before returning the function</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">// this is to
  avoid the memory leak</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">// stream the data from the response body
  only once</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">//
  it is not buffered in the memory</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">body</span><span class=\"p\">,</span><span class=\"w\"> </span><span
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
  class=\"nx\">body</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In the
  above example, we are trying to get the body from the response to the request sent
  at <a href=\"https://httpbin.org/html\"><code>https://httpbin.org/html</code></a>.
  We have used the simple <code>Get</code> method instead of <code>NewRequest</code>
  and <code>Do</code> for simplicity. The response is stored in <code>resp</code>,
  we also have added <code>defer resp.Body.Close()</code> which is to say that we
  will close the body reader object when the function is returned/closed. So, this
  means that the <code>Body</code> is not readily available data, we need to obtain/stream
  the data from the server. We have to receive the body in bytes as a tcp request,
  the body is streamed in a buffer.</p>\n<p>The response body is streamed from the
  server, which means that it's not immediately available as a complete data set.
  We read the body in bytes as it arrives over the network, and it's stored in a buffer,
  which allows us to process the data efficiently.</p>\n<h3>Reading Body in bytes</h3>\n<p>We
  can even read the body in bytes i.e. by reading a chunk of the buffer at a time.
  We can use the <a href=\"https://pkg.go.dev/bytes#Buffer\">bytes.Buffer</a> container
  object to store the body. Then we can create a slice of bytes as <code>[]bytes</code>
  of a certain size and read the body into the chunk. By writing the chunk into the
  buffer, we get the entire body from the response.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\">// web/methods/get/body.go</span><span class=\"w\"></span>\n\n\n<span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;bytes&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">reqURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://httpbin.org/html&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">reqURL</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">    </span><span class=\"c1\">// create a empty buffer</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">buf</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nb\">new</span><span
  class=\"p\">(</span><span class=\"nx\">bytes</span><span class=\"p\">.</span><span
  class=\"nx\">Buffer</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">    </span><span class=\"c1\">// create a chunk buffer of a fixed size</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">chunk</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nb\">make</span><span
  class=\"p\">([]</span><span class=\"kt\">byte</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">1024</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"c1\">// Read into buffer</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">n</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Read</span><span class=\"p\">(</span><span class=\"nx\">chunk</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"k\">break</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"c1\">// append the chunk to the buffer</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">buf</span><span class=\"p\">.</span><span
  class=\"nx\">Write</span><span class=\"p\">(</span><span class=\"nx\">chunk</span><span
  class=\"p\">[:</span><span class=\"nx\">n</span><span class=\"p\">])</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Printf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s\\n&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">chunk</span><span class=\"p\">[:</span><span class=\"nx\">n</span><span
  class=\"p\">])</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">    </span><span
  class=\"c1\">// entire body stored in bytes</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">buf</span><span
  class=\"p\">.</span><span class=\"nx\">String</span><span class=\"p\">())</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, the body is read chunkwise buffers and obtained as a slice of
  bytes. We define the request as a <code>GET</code> request to the <a href=\"http://httpbin.org/html\"><code>httpbin.org/html</code></a>.
  We create a new Buffer as a slice of bytes with <a href=\"https://pkg.go.dev/bytes#Buffer\">bytes.Buffer</a>,
  then we define chunk as a container to stream the response body with a particular
  size. We have taken <code>1024</code> bytes as the size of the chunk. Then inside
  an infinite for loop, we read the body as <code>n, err :=</code> <a href=\"http://resp.Body.Read\"><code>resp.Body.Read</code></a><code>(chunk)</code>.
  The code will read the body into the chunk(slice of byte) and the return value will
  be the size of the bytes read or an error. Then we check if there is no error, and
  if there is an error, we break the loop indicating we have completed reading the
  entire body or something went wrong. Then we write the chunk into the buffer that
  we allocated earlier as <code>buf</code>. This is a slice of bytes, we are basically
  populating the buffer with more slices of bytes.</p>\n<p>The entire body is then
  stored in the buffer as a slice of bytes. So, we have to cast it into a string to
  see the contents. So, this is how we can read the contents of a body in a response
  in chunks.</p>\n<h3>Parsing the JSON body with structs</h3>\n<p>If we have the structure
  of the response body already decided, then we can define a struct for the response
  body and then we can <a href=\"https://doc.akka.io/docs/akka-http/current/common/unmarshalling.html#unmarshalling:~:text=Unmarshalling,type%20T.\">Unmarshal</a>
  / deserialize/unpickle. This means we can convert the bytes representation of the
  data into a Golang-specific structure which is called a high-level representation
  of the data. We can parse the JSON body into a defined struct using <a href=\"https://pkg.go.dev/encoding/json#Unmarshal\">Unmarshal</a>
  or <a href=\"https://pkg.go.dev/encoding/json#Decoder.Decode\">Decode</a> methods
  in the <a href=\"https://pkg.go.dev/encoding/json\">json</a> package.</p>\n<p>Let's
  look at both the methods.</p>\n<h4>Using Unmarshal</h4>\n<p>The <code>Unmarshal</code>
  method takes in two parameters i.e. the body in bytes and the reference of the object
  that we want to unmarshal into. The method returns an error if there is a discrepancy
  in the returned JSON or the structure defined it is unable to deserialize the JSON
  object into the defined structure.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Product</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">                 </span><span class=\"kt\">int</span><span
  class=\"w\">      </span><span class=\"s\">`json:&quot;id&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Title</span><span class=\"w\">              </span><span
  class=\"kt\">string</span><span class=\"w\">   </span><span class=\"s\">`json:&quot;title&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Description</span><span
  class=\"w\">        </span><span class=\"kt\">string</span><span class=\"w\">   </span><span
  class=\"s\">`json:&quot;description&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Price</span><span class=\"w\">              </span><span
  class=\"kt\">float64</span><span class=\"w\">  </span><span class=\"s\">`json:&quot;price&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">DiscountPercentage</span><span
  class=\"w\"> </span><span class=\"kt\">float64</span><span class=\"w\">  </span><span
  class=\"s\">`json:&quot;discountPercentage&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Rating</span><span class=\"w\">             </span><span
  class=\"kt\">float64</span><span class=\"w\">  </span><span class=\"s\">`json:&quot;rating&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Stock</span><span
  class=\"w\">              </span><span class=\"kt\">int</span><span class=\"w\">
  \     </span><span class=\"s\">`json:&quot;stock&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Brand</span><span class=\"w\">              </span><span
  class=\"kt\">string</span><span class=\"w\">   </span><span class=\"s\">`json:&quot;brand&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Category</span><span
  class=\"w\">           </span><span class=\"kt\">string</span><span class=\"w\">
  \  </span><span class=\"s\">`json:&quot;category&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">Thumbnail</span><span class=\"w\">          </span><span
  class=\"kt\">string</span><span class=\"w\">   </span><span class=\"s\">`json:&quot;thumbnail,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">Images</span><span
  class=\"w\">             </span><span class=\"p\">[]</span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;-&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">reqURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummyjson.com/products/1&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">reqURL</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">body</span><span class=\"p\">,</span><span
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
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">data</span><span
  class=\"w\"> </span><span class=\"nx\">Product</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">body</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">data</span><span class=\"p\">);</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">data</span><span class=\"p\">.</span><span class=\"nx\">Title</span><span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\n\n<span class=\"o\">{</span><span class=\"m\">1</span> iPhone <span
  class=\"m\">9</span> An apple mobile which is nothing like apple <span class=\"m\">549</span>
  <span class=\"m\">12</span>.96 <span class=\"m\">4</span>.69 <span class=\"m\">94</span>
  Apple smartphones https://cdn.dummyjson.com/product-images/1/thumbnail.jpg <span
  class=\"o\">[]}</span>\niPhone <span class=\"m\">9</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we have defined a structure called Product with fields such as
  <code>ID</code>, <code>Title</code>, <code>Description</code>, etc. We use the JSON
  tag to specify how each field should be encoded to or decoded from JSON. These tags
  guide the Golang JSON encoders and decoders to correctly map JSON data to struct
  fields and vice versa. The <code>omitempty</code> option in a struct tag instructs
  the encoder to omit the field from the JSON output if the field's value is the zero
  value for its type (e.g., 0 for integers, &quot;&quot; for strings, nil for pointers,
  slices, and maps). This is useful for reducing the size of the JSON output by excluding
  empty or default-valued fields.</p>\n<p>Conversely, the <code>-</code> option in
  a struct tag tells the encoder and decoder to completely ignore the field. It will
  not be included in encoded JSON, nor will it be populated when decoding JSON into
  a struct. This is particularly useful for excluding fields that are meant for internal
  use only and should not be exposed through JSON.</p>\n<p>Therefore, <code>omitempty</code>
  is used to control the inclusion of fields in the JSON output based on their values,
  while <code>-</code> is used to exclude fields from both encoding and decoding from
  JSON.</p>\n<p>We send the <code>GET</code> request to the api <code>https://dummyjson.com/products/1</code>.
  The response from the request is read into a slice of bytes with <a href=\"https://pkg.go.dev/io#ReadAll\">io.ReadAll</a>
  that takes in a <a href=\"https://pkg.go.dev/io#Reader\">io.Reader</a> object in
  this case it is the <code>resp.Body</code> and it returns a slice of byte and error
  if any issue arises while reading in the body. Further, we can use the <a href=\"https://pkg.go.dev/encoding/json#Unmarshal\">Unmarshal</a>
  method to parse the slice of body <code>body</code> into the struct <code>Product</code>
  with the variable <code>data</code>, the reference to <code>&amp;data</code> indicates
  that the method will directly mutate/change this variable and populate the object
  with the fields from the body.</p>\n<p>So in a gist, to convert the JSON body into
  a golang native structure with <code>Unmarshal</code> with the following steps:</p>\n<ul>\n<li>Read
  the body into a slice of bytes using <code>io.ReadAll</code></li>\n<li>Create an
  object of the struct</li>\n<li>Pass the body as a slice of bytes and the reference
  of that object (struct instance) into the Unmarshal method</li>\n<li>Access the
  object with the fields in the struct</li>\n</ul>\n<p>In the output response, we
  can see the object is populated with the fields from the body. The <code>Title</code>
  field is accessed using the <code>data.Title</code> as we do with a normal golang
  struct. The <code>Images</code> field is not populated because we have always ignored/omitted
  from the json tag with <code>-</code>.</p>\n<h4>Using Decoder</h4>\n<p>Similar to
  the <code>Unmarshal</code> we can use the <a href=\"https://pkg.go.dev/encoding/json#Decoder\">Decoder</a>
  to parse the body into a struct. However, the parameters it takes are a bit different
  and it is a two-step process. We first create a <a href=\"https://pkg.go.dev/encoding/json#Decoder\">Decoder</a>
  object using the <a href=\"https://pkg.go.dev/encoding/json#NewDecoder\">NewDecoder</a>
  method, which takes in a <code>io.Reader</code> object, luckily the body from the
  response is already in that structure, so we can directly pass that <code>resp.Body</code>
  into the <code>NewDecoder</code> method. The second step is to decode the data into
  the object, here as well, we need to create the object of the struct and parse the
  reference to the object to the <a href=\"https://pkg.go.dev/encoding/json#Decoder.Decode\">Decode</a>
  method. The <code>Decode</code> method converts the bytes parsed in the <code>resp.Body</code>
  from the <code>Decoder</code> object and populates the fields of the object provided
  in the reference struct.</p>\n<p>So the steps for deserializing the json object
  into the struct with the decode method are:</p>\n<ul>\n<li>Create a decoder with
  <code>NewDecoder</code> method and pass the <code>resp.Body</code> as the parameter
  which is an <code>io.Reader</code> object</li>\n<li>Create an object of the struct</li>\n<li>Decode
  the body into that object using the <code>decoder.Decode</code> method</li>\n<li>Access
  the object with the fields in the struct</li>\n</ul>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Product</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">                 </span><span class=\"kt\">int</span><span
  class=\"w\">     </span><span class=\"s\">`json:&quot;id&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Title</span><span class=\"w\">              </span><span
  class=\"kt\">string</span><span class=\"w\">  </span><span class=\"s\">`json:&quot;title&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Description</span><span
  class=\"w\">        </span><span class=\"kt\">string</span><span class=\"w\">  </span><span
  class=\"s\">`json:&quot;description&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Price</span><span class=\"w\">              </span><span
  class=\"kt\">float64</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;price&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">DiscountPercentage</span><span
  class=\"w\"> </span><span class=\"kt\">float64</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;discountPercentage&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Rating</span><span class=\"w\">             </span><span
  class=\"kt\">float64</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;rating&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Stock</span><span
  class=\"w\">              </span><span class=\"kt\">int</span><span class=\"w\">
  \    </span><span class=\"s\">`json:&quot;stock&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Brand</span><span class=\"w\">              </span><span
  class=\"kt\">string</span><span class=\"w\">  </span><span class=\"s\">`json:&quot;brand&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Category</span><span
  class=\"w\">           </span><span class=\"kt\">string</span><span class=\"w\">
  \ </span><span class=\"s\">`json:&quot;category&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">Thumbnail</span><span class=\"w\">          </span><span
  class=\"kt\">string</span><span class=\"w\">   </span><span class=\"s\">`json:&quot;thumbnail,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">Images</span><span
  class=\"w\">             </span><span class=\"p\">[]</span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;-&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">reqURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummyjson.com/products/1&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">reqURL</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">data</span><span class=\"w\"> </span><span class=\"nx\">Product</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">decoder</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">NewDecoder</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">decoder</span><span class=\"p\">.</span><span
  class=\"nx\">Decode</span><span class=\"p\">(</span><span class=\"o\">&amp;</span><span
  class=\"nx\">data</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">data</span><span class=\"p\">.</span><span class=\"nx\">Title</span><span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\n\n<span class=\"o\">{</span><span class=\"m\">1</span> iPhone <span
  class=\"m\">9</span> An apple mobile which is nothing like apple <span class=\"m\">549</span>
  <span class=\"m\">12</span>.96 <span class=\"m\">4</span>.69 <span class=\"m\">94</span>
  Apple smartphones https://cdn.dummyjson.com/product-images/1/thumbnail.jpg <span
  class=\"o\">[]}</span>\niPhone <span class=\"m\">9</span>\n</pre></div>\n\n</pre>\n\n<p>We
  have first defined the struct <code>Product</code> with the <code>json:&quot;id&quot;</code>
  tag. As explained earlier, we have used the json tags to identify the fields from
  the json data to the structures while encoding and decoding.\nIn the above example,
  we have sent a <code>GET</code> request to the api endpoint <code>https://dummyjson.com/products/1</code>,
  and we have created a new decoder with the <code>NewDecoder</code> method with the
  <code>resp.Body</code> as the parameter. The <code>data</code> is created as a <code>Product</code>
  instance. The reference to <code>data</code> is parsed to the <code>Decode</code>
  method from the <code>decoder</code> instance as <code>&amp;data</code>. This method
  will either return <code>nil</code> or an <code>error</code>. Thereafter, we can
  check for errors and then only access the data object with its populated fields
  from the response body.</p>\n<p>There is a certain difference between the <code>Unmarshal</code>
  and <code>Decode</code> methods. The difference is just a slight performance improvement
  in the <code>NewDecoder</code> and <code>Decode</code> methods. Though it is not
  significant, having a little info about it might be handy in your use case. Read
  here for more info : <a href=\"https://dev.to/jpoly1219/to-unmarshal-or-to-decode-json-processing-in-go-explained-5870\">To
  Unmarshal or Decode</a></p>\n<h2>Adding Headers to a GET Request</h2>\n<p>We can
  even add headers before sending a <code>GET</code> request to a URL. By creating
  a <code>Request</code> object with the <code>NewRequest</code> method and adding
  a <a href=\"https://pkg.go.dev/net/http#Header\">Header</a> with the <a href=\"https://pkg.go.dev/net/http#Header.Add\">Add</a>
  method. The <code>Add</code> method will take in two parameters i.e. the key of
  the header, and the value of the header both as strings.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\">// web/methods/get/header.go</span><span class=\"w\"></span>\n\n<span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">MethodGet</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://api.github.com/users/mr-destructive&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">.</span><span class=\"nx\">Header</span><span
  class=\"p\">.</span><span class=\"nx\">Add</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Authorization&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;token YOUR_TOKEN&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">DefaultClient</span><span class=\"p\">.</span><span
  class=\"nx\">Do</span><span class=\"p\">(</span><span class=\"nx\">req</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">body</span><span class=\"p\">,</span><span
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
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">body</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run web/methods/get/header.go\n\n{&quot;message&quot;:&quot;Bad
  credentials&quot;,&quot;documentation_url&quot;:&quot;https://docs.github.com/rest&quot;}\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we have created a <code>GET</code> request to the <a href=\"https://api.github.com/users/mr-destructive\"><code>https://api.github.com/users/mr-destructive</code></a>
  the last portion is the username, it could be any valid username. The request is
  to the GitHub API, so it might require API Key/Tokens in the headers, however, if
  there are certain endpoints that do not require Authorization headers might work
  just fine.</p>\n<p>So, the above code will give <code>401</code> error indicating
  the request has wrong or invalid credentials, if we remove the header, the request
  will work fine. This is just an example, but headers are useful in working with
  APIs.</p>\n<p>Without adding the header:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run web/methods/get/header.go\n\n{&quot;login&quot;:&quot;Mr-Destructive&quot;,&quot;id&quot;:40317114,&quot;node_id&quot;:&quot;MDQ6VXNlcjQwMzE3MTE0&quot;,&quot;avatar_url&quot;:&quot;https://avatars.githubusercontent.com/u/40317114?v=4&quot;,&quot;gravatar_id&quot;:&quot;&quot;,&quot;url&quot;:&quot;https://api.github.com/users/Mr-Destructive&quot;,\n...
  \n&quot;updated_at&quot;:&quot;2023-10-10T17:57:22Z&quot;}\n</pre></div>\n\n</pre>\n\n<p>That's
  it from the 33rd part of the series, all the source code for the examples are linked
  in the GitHub on the <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/get/\">100
  days of Golang</a> repository.</p>\n<p><a href=\"https://github.com/Mr-Destructive/100-days-of-golang\">100-days-of-golang</a></p>\n<h2>References</h2>\n<ul>\n<li><a
  href=\"https://dev.to/jpoly1219/to-unmarshal-or-to-decode-json-processing-in-go-explained-5870\">To
  Unmarshal or Decode</a></li>\n<li><a href=\"https://drstearns.github.io/tutorials/gojson/\">Golang
  JSON tutorial</a></li>\n<li><a href=\"https://www.sohamkamani.com/golang/omitempty/\">Golang
  OmitEmpty</a></li>\n</ul>\n<h2>Conclusion</h2>\n<p>From this article, we explored
  the <code>GET</code> HTTP method in golang. By using a few examples for creating
  a get request, adding headers, reading response body, the basic use cases were demonstrated.</p>\n<p>Hopefully,
  you found this section helpful, if you have any comments or feedback, please let
  me know in the comments section or on my social handles. Thank you for reading.
  Happy Coding :)</p>\n"
cover: ''
date: 2023-10-28
datetime: 2023-10-28 00:00:00+00:00
description: Exploring the fundamentals of a GET method in golang. How to request
  a resource, parse body, headers, etc. in a HTTP GET request.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2023-10-28-GO-33-GET-Method.md
html: "<h2>Introduction</h2>\n<p>In this section of the series, we will be exploring
  how to send a <code>GET</code> HTTP request in golang. We will be understanding
  how to send a basic GET request, create an HTTP request and customize the client,
  add headers, read the response body, etc in the following sections of this post.</p>\n<h2>What
  is a GET method?</h2>\n<p>A <a href=\"https://en.wikipedia.org/wiki/HTTP#Request_methods\">GET</a>
  method in the context of an HTTP request is an action that is used to obtain data/resources.
  The <code>GET</code> method is used in a web application to get a resource like
  an HTML page, image, video, media, etc.</p>\n<p>Some common usecases of the <code>GET</code>
  method are:</p>\n<ul>\n<li>Loading a webpage</li>\n<li>Getting an image, file or
  other resource</li>\n<li>API requests to retrieve data</li>\n<li>Search queries
  sending filters and parameters</li>\n</ul>\n<h2>Basic GET Request</h2>\n<p>To use
  the HTTP method <code>GET</code> in golang, the <a href=\"https://pkg.go.dev/net/http\">net/http</a>
  package has a <a href=\"https://pkg.go.dev/net/http#Get\">Get</a> method. This method
  simply takes in a URL as a string and returns the <a href=\"https://pkg.go.dev/net/http#Response\">response</a>
  or an error. Let's look at how to send a basic HTTP GET request in golang.</p>\n<pre
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
  class=\"c1\">// web/methods/get/main.go</span><span class=\"w\"></span>\n\n\n<span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">reqURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://www.google.com&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">reqURL</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Status</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Status Code:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">StatusCode</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run main.go\n\n&amp;{200
  OK 200 HTTP/2.0 2 0 map[Alt-Svc:[h3=&quot;:443&quot;; ma=2592000,h3-29=&quot;:443&quot;;
  ma=2592000] Cache-Control:[private, max-age=0] Content-Security-Policy-Report-Only:[object-src
  &#39;none&#39;;base-uri &#39;self&#39;;script-src &#39;nonce-pdz_s8Gr0owwMbX8I9qNEQ&#39;
  &#39;strict-dynamic&#39; &#39;report-sample&#39; &#39;unsafe-eval&#39; &#39;unsafe-inline&#39;
  https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp] Content-Type:[text/html;
  charset=ISO-8859-1] Date:[Fri, 27 Oct 2023 09:37:04 GMT] Expires:[-1] P3p:[CP=&quot;This
  is not a P3P policy! See g.co/p3phelp for more info.&quot;] Server:[gws] Set-Cookie:[1P_JAR=2023-10-27-09;
  expires=Sun, 26-Nov-2023 09:37:04 GMT; path=/; domain=.google.com; Secure AEC=Ackid1Q5FARA_9d7f7znggUdw6DoJA1DBpI17Z0SWxN519Dc64EqmYVHlFg;
  expires=Wed, 24-Apr-2024 09:37:04 GMT; path=/; domain=.google.com; Secure; HttpOnly;
  SameSite=lax NID=511=EToBPqckCVRE7Paehug1PgNBKqe7lFLI9d12xJrGbvP9r8tkFIRWciry3gsy8FZ8OUIK4gE4PD-irgNzg4Y1fVePLdyu0AJdY_HcqL6zQYok-Adn-y5TDPmMCNuDnrouBfoxtqVjYY_4RFOe8jalkYto5fQAwzWnNJyw8K0avsw;
  expires=Sat, 27-Apr-2024 09:37:04 GMT; path=/; domain=.google.com; HttpOnly] X-Frame-Options:[SAMEORIGIN]
  X-Xss-Protection:[0]] 0xc000197920 -1 [] false true map[] 0xc0000ee000 0xc0000d8420}\n\n200
  OK\n\nStatus Code: 200\n</pre></div>\n\n</pre>\n\n<p>In the above code, we have
  defined a URL string as <code>reqURL</code> and used the <a href=\"https://pkg.go.dev/net/http#Get\">Get</a>
  method to send a GET request. The <code>Get</code> is parsed with the <code>reqURL</code>
  string and the return values are stored as <code>resp</code> and <code>err</code>.
  We have added an error check after calling the <code>Get</code> method in order
  to avoid errors later in the code.</p>\n<p>The <code>Get</code> method as seen from
  the output has returned a <code>*http.Response</code> object, we can use the <code>Status</code>
  and <code>StatusCode</code> properties to get the status code of the response. In
  this case, the status code of the response was <code>200</code>. The response <code>resp</code>
  is an object of type <code>http.Response</code> i.e. it has fields like <code>Body</code>,
  <code>StatusCode</code>, <code>Headers</code>, <code>Proto</code>, etc. We can get
  each individual field from the <code>resp</code> object. We will later look into
  how to read the <code>Body</code> field from the response, it is not directly read
  as a string nor it is stored in other forms, rather it is streamed from the requested
  URL.</p>\n<h2>Creating a GET request</h2>\n<p>We can even construct a GET request
  using the <a href=\"https://pkg.go.dev/net/http#NewRequest\">NewRequest</a> method.
  This is a low-level way of creating a <code>GET</code> request. We mention the <code>method</code>,
  <code>URL</code>, and the <code>body</code>, in the case of <code>GET</code> request,
  there is nobody. So, the <code>NewRequest</code> is a general way of creating a
  <code>http</code> request.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\">// web/methods/get/newreq.go</span><span class=\"w\"></span>\n\n<span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">reqURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://www.google.com&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">MethodGet</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">reqURL</span><span
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
  class=\"p\">(</span><span class=\"nx\">resp</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>As
  we can see, we construct a <code>GET</code> request using the <code>NewRequest</code>
  method and then use the <a href=\"https://pkg.go.dev/net/http#Client.Do\">Do</a>
  method to send the request to the server. The <a href=\"https://pkg.go.dev/net/http#DefaultClient\">http.DefaultClient</a>
  is used as a client to send the request, if we want to customize this we can create
  a new instance object of <a href=\"https://pkg.go.dev/net/http#Client\">http.Client</a>
  and use it to send requests. We will be taking a look at clients in another part
  of this series when we want to persist a connection or avoid connecting multiple
  times to the same application/URL.</p>\n<p>For now, we will go ahead with the DefaultClient.
  This will trigger the request, in this case, a <code>GET</code> request to the specified
  URL in the <code>reqURL</code> string. The <code>Do</code> method returns either
  a <code>http.Response</code> or an <code>error</code> just like the <code>Get</code>
  method did.</p>\n<h2>Reading the Response Body</h2>\n<p>We saw some different ways
  to send a <code>GET</code> request, now the below example will demonstrate how to
  read the body of the response. The response body is read from a buffer rather than
  loading the entire response into memory. It makes it flexible to parse the data
  efficiently and as per the needs. We will see how we use the <a href=\"https://pkg.go.dev/io\">io</a>
  package's <a href=\"https://pkg.go.dev/io#ReadAll\">ReadAll</a> method can be used
  to read from the response buffer.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\">// web/methods/get/body.go</span><span class=\"w\"></span>\n\n<span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">reqURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://httpbin.org/html&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">reqURL</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">// close the body object before returning the function</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">// this is to
  avoid the memory leak</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">// stream the data from the response body
  only once</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">//
  it is not buffered in the memory</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">body</span><span class=\"p\">,</span><span class=\"w\"> </span><span
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
  class=\"nx\">body</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In the
  above example, we are trying to get the body from the response to the request sent
  at <a href=\"https://httpbin.org/html\"><code>https://httpbin.org/html</code></a>.
  We have used the simple <code>Get</code> method instead of <code>NewRequest</code>
  and <code>Do</code> for simplicity. The response is stored in <code>resp</code>,
  we also have added <code>defer resp.Body.Close()</code> which is to say that we
  will close the body reader object when the function is returned/closed. So, this
  means that the <code>Body</code> is not readily available data, we need to obtain/stream
  the data from the server. We have to receive the body in bytes as a tcp request,
  the body is streamed in a buffer.</p>\n<p>The response body is streamed from the
  server, which means that it's not immediately available as a complete data set.
  We read the body in bytes as it arrives over the network, and it's stored in a buffer,
  which allows us to process the data efficiently.</p>\n<h3>Reading Body in bytes</h3>\n<p>We
  can even read the body in bytes i.e. by reading a chunk of the buffer at a time.
  We can use the <a href=\"https://pkg.go.dev/bytes#Buffer\">bytes.Buffer</a> container
  object to store the body. Then we can create a slice of bytes as <code>[]bytes</code>
  of a certain size and read the body into the chunk. By writing the chunk into the
  buffer, we get the entire body from the response.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\">// web/methods/get/body.go</span><span class=\"w\"></span>\n\n\n<span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;bytes&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">reqURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://httpbin.org/html&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">reqURL</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">    </span><span class=\"c1\">// create a empty buffer</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">buf</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nb\">new</span><span
  class=\"p\">(</span><span class=\"nx\">bytes</span><span class=\"p\">.</span><span
  class=\"nx\">Buffer</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">    </span><span class=\"c1\">// create a chunk buffer of a fixed size</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">chunk</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nb\">make</span><span
  class=\"p\">([]</span><span class=\"kt\">byte</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"mi\">1024</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"k\">for</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"c1\">// Read into buffer</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">n</span><span class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Read</span><span class=\"p\">(</span><span class=\"nx\">chunk</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"k\">break</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"c1\">// append the chunk to the buffer</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">buf</span><span class=\"p\">.</span><span
  class=\"nx\">Write</span><span class=\"p\">(</span><span class=\"nx\">chunk</span><span
  class=\"p\">[:</span><span class=\"nx\">n</span><span class=\"p\">])</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Printf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s\\n&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">chunk</span><span class=\"p\">[:</span><span class=\"nx\">n</span><span
  class=\"p\">])</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">    </span><span
  class=\"c1\">// entire body stored in bytes</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">buf</span><span
  class=\"p\">.</span><span class=\"nx\">String</span><span class=\"p\">())</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, the body is read chunkwise buffers and obtained as a slice of
  bytes. We define the request as a <code>GET</code> request to the <a href=\"http://httpbin.org/html\"><code>httpbin.org/html</code></a>.
  We create a new Buffer as a slice of bytes with <a href=\"https://pkg.go.dev/bytes#Buffer\">bytes.Buffer</a>,
  then we define chunk as a container to stream the response body with a particular
  size. We have taken <code>1024</code> bytes as the size of the chunk. Then inside
  an infinite for loop, we read the body as <code>n, err :=</code> <a href=\"http://resp.Body.Read\"><code>resp.Body.Read</code></a><code>(chunk)</code>.
  The code will read the body into the chunk(slice of byte) and the return value will
  be the size of the bytes read or an error. Then we check if there is no error, and
  if there is an error, we break the loop indicating we have completed reading the
  entire body or something went wrong. Then we write the chunk into the buffer that
  we allocated earlier as <code>buf</code>. This is a slice of bytes, we are basically
  populating the buffer with more slices of bytes.</p>\n<p>The entire body is then
  stored in the buffer as a slice of bytes. So, we have to cast it into a string to
  see the contents. So, this is how we can read the contents of a body in a response
  in chunks.</p>\n<h3>Parsing the JSON body with structs</h3>\n<p>If we have the structure
  of the response body already decided, then we can define a struct for the response
  body and then we can <a href=\"https://doc.akka.io/docs/akka-http/current/common/unmarshalling.html#unmarshalling:~:text=Unmarshalling,type%20T.\">Unmarshal</a>
  / deserialize/unpickle. This means we can convert the bytes representation of the
  data into a Golang-specific structure which is called a high-level representation
  of the data. We can parse the JSON body into a defined struct using <a href=\"https://pkg.go.dev/encoding/json#Unmarshal\">Unmarshal</a>
  or <a href=\"https://pkg.go.dev/encoding/json#Decoder.Decode\">Decode</a> methods
  in the <a href=\"https://pkg.go.dev/encoding/json\">json</a> package.</p>\n<p>Let's
  look at both the methods.</p>\n<h4>Using Unmarshal</h4>\n<p>The <code>Unmarshal</code>
  method takes in two parameters i.e. the body in bytes and the reference of the object
  that we want to unmarshal into. The method returns an error if there is a discrepancy
  in the returned JSON or the structure defined it is unable to deserialize the JSON
  object into the defined structure.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Product</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">                 </span><span class=\"kt\">int</span><span
  class=\"w\">      </span><span class=\"s\">`json:&quot;id&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Title</span><span class=\"w\">              </span><span
  class=\"kt\">string</span><span class=\"w\">   </span><span class=\"s\">`json:&quot;title&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Description</span><span
  class=\"w\">        </span><span class=\"kt\">string</span><span class=\"w\">   </span><span
  class=\"s\">`json:&quot;description&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Price</span><span class=\"w\">              </span><span
  class=\"kt\">float64</span><span class=\"w\">  </span><span class=\"s\">`json:&quot;price&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">DiscountPercentage</span><span
  class=\"w\"> </span><span class=\"kt\">float64</span><span class=\"w\">  </span><span
  class=\"s\">`json:&quot;discountPercentage&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Rating</span><span class=\"w\">             </span><span
  class=\"kt\">float64</span><span class=\"w\">  </span><span class=\"s\">`json:&quot;rating&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Stock</span><span
  class=\"w\">              </span><span class=\"kt\">int</span><span class=\"w\">
  \     </span><span class=\"s\">`json:&quot;stock&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Brand</span><span class=\"w\">              </span><span
  class=\"kt\">string</span><span class=\"w\">   </span><span class=\"s\">`json:&quot;brand&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Category</span><span
  class=\"w\">           </span><span class=\"kt\">string</span><span class=\"w\">
  \  </span><span class=\"s\">`json:&quot;category&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">Thumbnail</span><span class=\"w\">          </span><span
  class=\"kt\">string</span><span class=\"w\">   </span><span class=\"s\">`json:&quot;thumbnail,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">Images</span><span
  class=\"w\">             </span><span class=\"p\">[]</span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;-&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">reqURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummyjson.com/products/1&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">reqURL</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">body</span><span class=\"p\">,</span><span
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
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">data</span><span
  class=\"w\"> </span><span class=\"nx\">Product</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">body</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">data</span><span class=\"p\">);</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">data</span><span class=\"p\">.</span><span class=\"nx\">Title</span><span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\n\n<span class=\"o\">{</span><span class=\"m\">1</span> iPhone <span
  class=\"m\">9</span> An apple mobile which is nothing like apple <span class=\"m\">549</span>
  <span class=\"m\">12</span>.96 <span class=\"m\">4</span>.69 <span class=\"m\">94</span>
  Apple smartphones https://cdn.dummyjson.com/product-images/1/thumbnail.jpg <span
  class=\"o\">[]}</span>\niPhone <span class=\"m\">9</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we have defined a structure called Product with fields such as
  <code>ID</code>, <code>Title</code>, <code>Description</code>, etc. We use the JSON
  tag to specify how each field should be encoded to or decoded from JSON. These tags
  guide the Golang JSON encoders and decoders to correctly map JSON data to struct
  fields and vice versa. The <code>omitempty</code> option in a struct tag instructs
  the encoder to omit the field from the JSON output if the field's value is the zero
  value for its type (e.g., 0 for integers, &quot;&quot; for strings, nil for pointers,
  slices, and maps). This is useful for reducing the size of the JSON output by excluding
  empty or default-valued fields.</p>\n<p>Conversely, the <code>-</code> option in
  a struct tag tells the encoder and decoder to completely ignore the field. It will
  not be included in encoded JSON, nor will it be populated when decoding JSON into
  a struct. This is particularly useful for excluding fields that are meant for internal
  use only and should not be exposed through JSON.</p>\n<p>Therefore, <code>omitempty</code>
  is used to control the inclusion of fields in the JSON output based on their values,
  while <code>-</code> is used to exclude fields from both encoding and decoding from
  JSON.</p>\n<p>We send the <code>GET</code> request to the api <code>https://dummyjson.com/products/1</code>.
  The response from the request is read into a slice of bytes with <a href=\"https://pkg.go.dev/io#ReadAll\">io.ReadAll</a>
  that takes in a <a href=\"https://pkg.go.dev/io#Reader\">io.Reader</a> object in
  this case it is the <code>resp.Body</code> and it returns a slice of byte and error
  if any issue arises while reading in the body. Further, we can use the <a href=\"https://pkg.go.dev/encoding/json#Unmarshal\">Unmarshal</a>
  method to parse the slice of body <code>body</code> into the struct <code>Product</code>
  with the variable <code>data</code>, the reference to <code>&amp;data</code> indicates
  that the method will directly mutate/change this variable and populate the object
  with the fields from the body.</p>\n<p>So in a gist, to convert the JSON body into
  a golang native structure with <code>Unmarshal</code> with the following steps:</p>\n<ul>\n<li>Read
  the body into a slice of bytes using <code>io.ReadAll</code></li>\n<li>Create an
  object of the struct</li>\n<li>Pass the body as a slice of bytes and the reference
  of that object (struct instance) into the Unmarshal method</li>\n<li>Access the
  object with the fields in the struct</li>\n</ul>\n<p>In the output response, we
  can see the object is populated with the fields from the body. The <code>Title</code>
  field is accessed using the <code>data.Title</code> as we do with a normal golang
  struct. The <code>Images</code> field is not populated because we have always ignored/omitted
  from the json tag with <code>-</code>.</p>\n<h4>Using Decoder</h4>\n<p>Similar to
  the <code>Unmarshal</code> we can use the <a href=\"https://pkg.go.dev/encoding/json#Decoder\">Decoder</a>
  to parse the body into a struct. However, the parameters it takes are a bit different
  and it is a two-step process. We first create a <a href=\"https://pkg.go.dev/encoding/json#Decoder\">Decoder</a>
  object using the <a href=\"https://pkg.go.dev/encoding/json#NewDecoder\">NewDecoder</a>
  method, which takes in a <code>io.Reader</code> object, luckily the body from the
  response is already in that structure, so we can directly pass that <code>resp.Body</code>
  into the <code>NewDecoder</code> method. The second step is to decode the data into
  the object, here as well, we need to create the object of the struct and parse the
  reference to the object to the <a href=\"https://pkg.go.dev/encoding/json#Decoder.Decode\">Decode</a>
  method. The <code>Decode</code> method converts the bytes parsed in the <code>resp.Body</code>
  from the <code>Decoder</code> object and populates the fields of the object provided
  in the reference struct.</p>\n<p>So the steps for deserializing the json object
  into the struct with the decode method are:</p>\n<ul>\n<li>Create a decoder with
  <code>NewDecoder</code> method and pass the <code>resp.Body</code> as the parameter
  which is an <code>io.Reader</code> object</li>\n<li>Create an object of the struct</li>\n<li>Decode
  the body into that object using the <code>decoder.Decode</code> method</li>\n<li>Access
  the object with the fields in the struct</li>\n</ul>\n<pre class='wrapper'>\n\n<div
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
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Product</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">                 </span><span class=\"kt\">int</span><span
  class=\"w\">     </span><span class=\"s\">`json:&quot;id&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Title</span><span class=\"w\">              </span><span
  class=\"kt\">string</span><span class=\"w\">  </span><span class=\"s\">`json:&quot;title&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Description</span><span
  class=\"w\">        </span><span class=\"kt\">string</span><span class=\"w\">  </span><span
  class=\"s\">`json:&quot;description&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Price</span><span class=\"w\">              </span><span
  class=\"kt\">float64</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;price&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">DiscountPercentage</span><span
  class=\"w\"> </span><span class=\"kt\">float64</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;discountPercentage&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Rating</span><span class=\"w\">             </span><span
  class=\"kt\">float64</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;rating&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Stock</span><span
  class=\"w\">              </span><span class=\"kt\">int</span><span class=\"w\">
  \    </span><span class=\"s\">`json:&quot;stock&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Brand</span><span class=\"w\">              </span><span
  class=\"kt\">string</span><span class=\"w\">  </span><span class=\"s\">`json:&quot;brand&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Category</span><span
  class=\"w\">           </span><span class=\"kt\">string</span><span class=\"w\">
  \ </span><span class=\"s\">`json:&quot;category&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">Thumbnail</span><span class=\"w\">          </span><span
  class=\"kt\">string</span><span class=\"w\">   </span><span class=\"s\">`json:&quot;thumbnail,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">Images</span><span
  class=\"w\">             </span><span class=\"p\">[]</span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;-&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">reqURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummyjson.com/products/1&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Get</span><span class=\"p\">(</span><span
  class=\"nx\">reqURL</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">data</span><span class=\"w\"> </span><span class=\"nx\">Product</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">decoder</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">NewDecoder</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">decoder</span><span class=\"p\">.</span><span
  class=\"nx\">Decode</span><span class=\"p\">(</span><span class=\"o\">&amp;</span><span
  class=\"nx\">data</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">data</span><span class=\"p\">.</span><span class=\"nx\">Title</span><span
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  go run main.go\n\n<span class=\"o\">{</span><span class=\"m\">1</span> iPhone <span
  class=\"m\">9</span> An apple mobile which is nothing like apple <span class=\"m\">549</span>
  <span class=\"m\">12</span>.96 <span class=\"m\">4</span>.69 <span class=\"m\">94</span>
  Apple smartphones https://cdn.dummyjson.com/product-images/1/thumbnail.jpg <span
  class=\"o\">[]}</span>\niPhone <span class=\"m\">9</span>\n</pre></div>\n\n</pre>\n\n<p>We
  have first defined the struct <code>Product</code> with the <code>json:&quot;id&quot;</code>
  tag. As explained earlier, we have used the json tags to identify the fields from
  the json data to the structures while encoding and decoding.\nIn the above example,
  we have sent a <code>GET</code> request to the api endpoint <code>https://dummyjson.com/products/1</code>,
  and we have created a new decoder with the <code>NewDecoder</code> method with the
  <code>resp.Body</code> as the parameter. The <code>data</code> is created as a <code>Product</code>
  instance. The reference to <code>data</code> is parsed to the <code>Decode</code>
  method from the <code>decoder</code> instance as <code>&amp;data</code>. This method
  will either return <code>nil</code> or an <code>error</code>. Thereafter, we can
  check for errors and then only access the data object with its populated fields
  from the response body.</p>\n<p>There is a certain difference between the <code>Unmarshal</code>
  and <code>Decode</code> methods. The difference is just a slight performance improvement
  in the <code>NewDecoder</code> and <code>Decode</code> methods. Though it is not
  significant, having a little info about it might be handy in your use case. Read
  here for more info : <a href=\"https://dev.to/jpoly1219/to-unmarshal-or-to-decode-json-processing-in-go-explained-5870\">To
  Unmarshal or Decode</a></p>\n<h2>Adding Headers to a GET Request</h2>\n<p>We can
  even add headers before sending a <code>GET</code> request to a URL. By creating
  a <code>Request</code> object with the <code>NewRequest</code> method and adding
  a <a href=\"https://pkg.go.dev/net/http#Header\">Header</a> with the <a href=\"https://pkg.go.dev/net/http#Header.Add\">Add</a>
  method. The <code>Add</code> method will take in two parameters i.e. the key of
  the header, and the value of the header both as strings.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\">// web/methods/get/header.go</span><span class=\"w\"></span>\n\n<span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">func</span><span class=\"w\"> </span><span
  class=\"nx\">main</span><span class=\"p\">()</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">MethodGet</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://api.github.com/users/mr-destructive&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">.</span><span class=\"nx\">Header</span><span
  class=\"p\">.</span><span class=\"nx\">Add</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Authorization&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;token YOUR_TOKEN&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">DefaultClient</span><span class=\"p\">.</span><span
  class=\"nx\">Do</span><span class=\"p\">(</span><span class=\"nx\">req</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">body</span><span class=\"p\">,</span><span
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
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">body</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run web/methods/get/header.go\n\n{&quot;message&quot;:&quot;Bad
  credentials&quot;,&quot;documentation_url&quot;:&quot;https://docs.github.com/rest&quot;}\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we have created a <code>GET</code> request to the <a href=\"https://api.github.com/users/mr-destructive\"><code>https://api.github.com/users/mr-destructive</code></a>
  the last portion is the username, it could be any valid username. The request is
  to the GitHub API, so it might require API Key/Tokens in the headers, however, if
  there are certain endpoints that do not require Authorization headers might work
  just fine.</p>\n<p>So, the above code will give <code>401</code> error indicating
  the request has wrong or invalid credentials, if we remove the header, the request
  will work fine. This is just an example, but headers are useful in working with
  APIs.</p>\n<p>Without adding the header:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ go run web/methods/get/header.go\n\n{&quot;login&quot;:&quot;Mr-Destructive&quot;,&quot;id&quot;:40317114,&quot;node_id&quot;:&quot;MDQ6VXNlcjQwMzE3MTE0&quot;,&quot;avatar_url&quot;:&quot;https://avatars.githubusercontent.com/u/40317114?v=4&quot;,&quot;gravatar_id&quot;:&quot;&quot;,&quot;url&quot;:&quot;https://api.github.com/users/Mr-Destructive&quot;,\n...
  \n&quot;updated_at&quot;:&quot;2023-10-10T17:57:22Z&quot;}\n</pre></div>\n\n</pre>\n\n<p>That's
  it from the 33rd part of the series, all the source code for the examples are linked
  in the GitHub on the <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/get/\">100
  days of Golang</a> repository.</p>\n<p><a href=\"https://github.com/Mr-Destructive/100-days-of-golang\">100-days-of-golang</a></p>\n<h2>References</h2>\n<ul>\n<li><a
  href=\"https://dev.to/jpoly1219/to-unmarshal-or-to-decode-json-processing-in-go-explained-5870\">To
  Unmarshal or Decode</a></li>\n<li><a href=\"https://drstearns.github.io/tutorials/gojson/\">Golang
  JSON tutorial</a></li>\n<li><a href=\"https://www.sohamkamani.com/golang/omitempty/\">Golang
  OmitEmpty</a></li>\n</ul>\n<h2>Conclusion</h2>\n<p>From this article, we explored
  the <code>GET</code> HTTP method in golang. By using a few examples for creating
  a get request, adding headers, reading response body, the basic use cases were demonstrated.</p>\n<p>Hopefully,
  you found this section helpful, if you have any comments or feedback, please let
  me know in the comments section or on my social handles. Thank you for reading.
  Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/100-days-of-golang/golang-033-get-method.png
long_description: 'In this section of the series, we will be exploring how to send
  a  A  Some common usecases of the  Loading a webpage Getting an image, file or other
  resource API requests to retrieve data Search queries sending filters and parameters
  To use the HTTP '
now: 2025-05-01 18:11:33.314180
path: blog/posts/2023-10-28-GO-33-GET-Method.md
prevnext: null
series:
- 100-days-of-golang
slug: golang-web-get-method
status: published
tags:
- go
templateKey: blog-post
title: 'Golang Web: GET Method'
today: 2025-05-01
---

## Introduction

In this section of the series, we will be exploring how to send a `GET` HTTP request in golang. We will be understanding how to send a basic GET request, create an HTTP request and customize the client, add headers, read the response body, etc in the following sections of this post.

## What is a GET method?

A [GET](https://en.wikipedia.org/wiki/HTTP#Request_methods) method in the context of an HTTP request is an action that is used to obtain data/resources. The `GET` method is used in a web application to get a resource like an HTML page, image, video, media, etc.

Some common usecases of the `GET` method are:

- Loading a webpage
- Getting an image, file or other resource
- API requests to retrieve data
- Search queries sending filters and parameters

## Basic GET Request

To use the HTTP method `GET` in golang, the [net/http](https://pkg.go.dev/net/http) package has a [Get](https://pkg.go.dev/net/http#Get) method. This method simply takes in a URL as a string and returns the [response](https://pkg.go.dev/net/http#Response) or an error. Let's look at how to send a basic HTTP GET request in golang.

```go
// web/methods/get/main.go


package main

import (
	"fmt"
	"net/http"
)

func main() {
	reqURL := "https://www.google.com"
	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
    fmt.Println(resp.Status)
    fmt.Println("Status Code:", resp.StatusCode)
}
```

```
$ go run main.go

&{200 OK 200 HTTP/2.0 2 0 map[Alt-Svc:[h3=":443"; ma=2592000,h3-29=":443"; ma=2592000] Cache-Control:[private, max-age=0] Content-Security-Policy-Report-Only:[object-src 'none';base-uri 'self';script-src 'nonce-pdz_s8Gr0owwMbX8I9qNEQ' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp] Content-Type:[text/html; charset=ISO-8859-1] Date:[Fri, 27 Oct 2023 09:37:04 GMT] Expires:[-1] P3p:[CP="This is not a P3P policy! See g.co/p3phelp for more info."] Server:[gws] Set-Cookie:[1P_JAR=2023-10-27-09; expires=Sun, 26-Nov-2023 09:37:04 GMT; path=/; domain=.google.com; Secure AEC=Ackid1Q5FARA_9d7f7znggUdw6DoJA1DBpI17Z0SWxN519Dc64EqmYVHlFg; expires=Wed, 24-Apr-2024 09:37:04 GMT; path=/; domain=.google.com; Secure; HttpOnly; SameSite=lax NID=511=EToBPqckCVRE7Paehug1PgNBKqe7lFLI9d12xJrGbvP9r8tkFIRWciry3gsy8FZ8OUIK4gE4PD-irgNzg4Y1fVePLdyu0AJdY_HcqL6zQYok-Adn-y5TDPmMCNuDnrouBfoxtqVjYY_4RFOe8jalkYto5fQAwzWnNJyw8K0avsw; expires=Sat, 27-Apr-2024 09:37:04 GMT; path=/; domain=.google.com; HttpOnly] X-Frame-Options:[SAMEORIGIN] X-Xss-Protection:[0]] 0xc000197920 -1 [] false true map[] 0xc0000ee000 0xc0000d8420}

200 OK

Status Code: 200
```

In the above code, we have defined a URL string as `reqURL` and used the [Get](https://pkg.go.dev/net/http#Get) method to send a GET request. The `Get` is parsed with the `reqURL` string and the return values are stored as `resp` and `err`. We have added an error check after calling the `Get` method in order to avoid errors later in the code.

The `Get` method as seen from the output has returned a `*http.Response` object, we can use the `Status` and `StatusCode` properties to get the status code of the response. In this case, the status code of the response was `200`. The response `resp` is an object of type `http.Response` i.e. it has fields like `Body`, `StatusCode`, `Headers`, `Proto`, etc. We can get each individual field from the `resp` object. We will later look into how to read the `Body` field from the response, it is not directly read as a string nor it is stored in other forms, rather it is streamed from the requested URL.

## Creating a GET request

We can even construct a GET request using the [NewRequest](https://pkg.go.dev/net/http#NewRequest) method. This is a low-level way of creating a `GET` request. We mention the `method`, `URL`, and the `body`, in the case of `GET` request, there is nobody. So, the `NewRequest` is a general way of creating a `http` request.

```go
// web/methods/get/newreq.go

package main

import (
	"fmt"
	"net/http"
)

func main() {
	reqURL := "https://www.google.com"
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
```

As we can see, we construct a `GET` request using the `NewRequest` method and then use the [Do](https://pkg.go.dev/net/http#Client.Do) method to send the request to the server. The [http.DefaultClient](https://pkg.go.dev/net/http#DefaultClient) is used as a client to send the request, if we want to customize this we can create a new instance object of [http.Client](https://pkg.go.dev/net/http#Client) and use it to send requests. We will be taking a look at clients in another part of this series when we want to persist a connection or avoid connecting multiple times to the same application/URL.

For now, we will go ahead with the DefaultClient. This will trigger the request, in this case, a `GET` request to the specified URL in the `reqURL` string. The `Do` method returns either a `http.Response` or an `error` just like the `Get` method did.

## Reading the Response Body

We saw some different ways to send a `GET` request, now the below example will demonstrate how to read the body of the response. The response body is read from a buffer rather than loading the entire response into memory. It makes it flexible to parse the data efficiently and as per the needs. We will see how we use the [io](https://pkg.go.dev/io) package's [ReadAll](https://pkg.go.dev/io#ReadAll) method can be used to read from the response buffer.

```go
// web/methods/get/body.go

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	reqURL := "https://httpbin.org/html"
	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}
	// close the body object before returning the function
	// this is to avoid the memory leak
	defer resp.Body.Close()

	// stream the data from the response body only once
	// it is not buffered in the memory
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
```

In the above example, we are trying to get the body from the response to the request sent at [`https://httpbin.org/html`](https://httpbin.org/html). We have used the simple `Get` method instead of `NewRequest` and `Do` for simplicity. The response is stored in `resp`, we also have added `defer resp.Body.Close()` which is to say that we will close the body reader object when the function is returned/closed. So, this means that the `Body` is not readily available data, we need to obtain/stream the data from the server. We have to receive the body in bytes as a tcp request, the body is streamed in a buffer.

The response body is streamed from the server, which means that it's not immediately available as a complete data set. We read the body in bytes as it arrives over the network, and it's stored in a buffer, which allows us to process the data efficiently.

### Reading Body in bytes

We can even read the body in bytes i.e. by reading a chunk of the buffer at a time. We can use the [bytes.Buffer](https://pkg.go.dev/bytes#Buffer) container object to store the body. Then we can create a slice of bytes as `[]bytes` of a certain size and read the body into the chunk. By writing the chunk into the buffer, we get the entire body from the response.

```go
// web/methods/get/body.go


package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	reqURL := "https://httpbin.org/html"
	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

    // create a empty buffer
	buf := new(bytes.Buffer)

    // create a chunk buffer of a fixed size
	chunk := make([]byte, 1024)

	for {
		// Read into buffer
		n, err := resp.Body.Read(chunk)
		if err != nil {
			break
		}
        // append the chunk to the buffer
		buf.Write(chunk[:n])
		fmt.Printf("%s\n", chunk[:n])
	}

    // entire body stored in bytes
	fmt.Println(buf.String())
}
```

In the above example, the body is read chunkwise buffers and obtained as a slice of bytes. We define the request as a `GET` request to the [`httpbin.org/html`](http://httpbin.org/html). We create a new Buffer as a slice of bytes with [bytes.Buffer](https://pkg.go.dev/bytes#Buffer), then we define chunk as a container to stream the response body with a particular size. We have taken `1024` bytes as the size of the chunk. Then inside an infinite for loop, we read the body as `n, err :=` [`resp.Body.Read`](http://resp.Body.Read)`(chunk)`. The code will read the body into the chunk(slice of byte) and the return value will be the size of the bytes read or an error. Then we check if there is no error, and if there is an error, we break the loop indicating we have completed reading the entire body or something went wrong. Then we write the chunk into the buffer that we allocated earlier as `buf`. This is a slice of bytes, we are basically populating the buffer with more slices of bytes.

The entire body is then stored in the buffer as a slice of bytes. So, we have to cast it into a string to see the contents. So, this is how we can read the contents of a body in a response in chunks.

### Parsing the JSON body with structs

If we have the structure of the response body already decided, then we can define a struct for the response body and then we can [Unmarshal](https://doc.akka.io/docs/akka-http/current/common/unmarshalling.html#unmarshalling:~:text=Unmarshalling,type%20T.) / deserialize/unpickle. This means we can convert the bytes representation of the data into a Golang-specific structure which is called a high-level representation of the data. We can parse the JSON body into a defined struct using [Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) or [Decode](https://pkg.go.dev/encoding/json#Decoder.Decode) methods in the [json](https://pkg.go.dev/encoding/json) package.

Let's look at both the methods.

#### Using Unmarshal

The `Unmarshal` method takes in two parameters i.e. the body in bytes and the reference of the object that we want to unmarshal into. The method returns an error if there is a discrepancy in the returned JSON or the structure defined it is unable to deserialize the JSON object into the defined structure.

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Product struct {
	ID                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	Price              float64  `json:"price"`
	DiscountPercentage float64  `json:"discountPercentage"`
	Rating             float64  `json:"rating"`
	Stock              int      `json:"stock"`
	Brand              string   `json:"brand"`
	Category           string   `json:"category"`
    Thumbnail          string   `json:"thumbnail,omitempty"`
    Images             []string `json:"-"`
}

func main() {
	reqURL := "https://dummyjson.com/products/1"
	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data Product
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	fmt.Println(data)
	fmt.Println(data.Title)
}
```

```bash
$ go run main.go

{1 iPhone 9 An apple mobile which is nothing like apple 549 12.96 4.69 94 Apple smartphones https://cdn.dummyjson.com/product-images/1/thumbnail.jpg []}
iPhone 9
```

In the above example, we have defined a structure called Product with fields such as `ID`, `Title`, `Description`, etc. We use the JSON tag to specify how each field should be encoded to or decoded from JSON. These tags guide the Golang JSON encoders and decoders to correctly map JSON data to struct fields and vice versa. The `omitempty` option in a struct tag instructs the encoder to omit the field from the JSON output if the field's value is the zero value for its type (e.g., 0 for integers, "" for strings, nil for pointers, slices, and maps). This is useful for reducing the size of the JSON output by excluding empty or default-valued fields.

Conversely, the `-` option in a struct tag tells the encoder and decoder to completely ignore the field. It will not be included in encoded JSON, nor will it be populated when decoding JSON into a struct. This is particularly useful for excluding fields that are meant for internal use only and should not be exposed through JSON.

Therefore, `omitempty` is used to control the inclusion of fields in the JSON output based on their values, while `-` is used to exclude fields from both encoding and decoding from JSON.

We send the `GET` request to the api `https://dummyjson.com/products/1`. The response from the request is read into a slice of bytes with [io.ReadAll](https://pkg.go.dev/io#ReadAll) that takes in a [io.Reader](https://pkg.go.dev/io#Reader) object in this case it is the `resp.Body` and it returns a slice of byte and error if any issue arises while reading in the body. Further, we can use the [Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) method to parse the slice of body `body` into the struct `Product` with the variable `data`, the reference to `&data` indicates that the method will directly mutate/change this variable and populate the object with the fields from the body.


So in a gist, to convert the JSON body into a golang native structure with `Unmarshal` with the following steps:

- Read the body into a slice of bytes using `io.ReadAll`
- Create an object of the struct
- Pass the body as a slice of bytes and the reference of that object (struct instance) into the Unmarshal method
- Access the object with the fields in the struct

In the output response, we can see the object is populated with the fields from the body. The `Title` field is accessed using the `data.Title` as we do with a normal golang struct. The `Images` field is not populated because we have always ignored/omitted from the json tag with `-`.

#### Using Decoder

Similar to the `Unmarshal` we can use the [Decoder](https://pkg.go.dev/encoding/json#Decoder) to parse the body into a struct. However, the parameters it takes are a bit different and it is a two-step process. We first create a [Decoder](https://pkg.go.dev/encoding/json#Decoder) object using the [NewDecoder](https://pkg.go.dev/encoding/json#NewDecoder) method, which takes in a `io.Reader` object, luckily the body from the response is already in that structure, so we can directly pass that `resp.Body` into the `NewDecoder` method. The second step is to decode the data into the object, here as well, we need to create the object of the struct and parse the reference to the object to the [Decode](https://pkg.go.dev/encoding/json#Decoder.Decode) method. The `Decode` method converts the bytes parsed in the `resp.Body` from the `Decoder` object and populates the fields of the object provided in the reference struct.

So the steps for deserializing the json object into the struct with the decode method are:

- Create a decoder with `NewDecoder` method and pass the `resp.Body` as the parameter which is an `io.Reader` object
- Create an object of the struct
- Decode the body into that object using the `decoder.Decode` method
- Access the object with the fields in the struct

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID                 int     `json:"id"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	Price              float64 `json:"price"`
	DiscountPercentage float64 `json:"discountPercentage"`
	Rating             float64 `json:"rating"`
	Stock              int     `json:"stock"`
	Brand              string  `json:"brand"`
	Category           string  `json:"category"`
    Thumbnail          string   `json:"thumbnail,omitempty"`
    Images             []string `json:"-"`
}

func main() {
	reqURL := "https://dummyjson.com/products/1"
	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data Product
	decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&data)
    if err != nil {
        panic(err)
    }

	fmt.Println(data)
	fmt.Println(data.Title)
}
```

```bash
$ go run main.go

{1 iPhone 9 An apple mobile which is nothing like apple 549 12.96 4.69 94 Apple smartphones https://cdn.dummyjson.com/product-images/1/thumbnail.jpg []}
iPhone 9
```


We have first defined the struct `Product` with the `json:"id"` tag. As explained earlier, we have used the json tags to identify the fields from the json data to the structures while encoding and decoding.
In the above example, we have sent a `GET` request to the api endpoint `https://dummyjson.com/products/1`, and we have created a new decoder with the `NewDecoder` method with the `resp.Body` as the parameter. The `data` is created as a `Product` instance. The reference to `data` is parsed to the `Decode` method from the `decoder` instance as `&data`. This method will either return `nil` or an `error`. Thereafter, we can check for errors and then only access the data object with its populated fields from the response body.

There is a certain difference between the `Unmarshal` and `Decode` methods. The difference is just a slight performance improvement in the `NewDecoder` and `Decode` methods. Though it is not significant, having a little info about it might be handy in your use case. Read here for more info : [To Unmarshal or Decode](https://dev.to/jpoly1219/to-unmarshal-or-to-decode-json-processing-in-go-explained-5870)

## Adding Headers to a GET Request

We can even add headers before sending a `GET` request to a URL. By creating a `Request` object with the `NewRequest` method and adding a [Header](https://pkg.go.dev/net/http#Header) with the [Add](https://pkg.go.dev/net/http#Header.Add) method. The `Add` method will take in two parameters i.e. the key of the header, and the value of the header both as strings.

```go
// web/methods/get/header.go

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/users/mr-destructive", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "token YOUR_TOKEN")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
```

```
$ go run web/methods/get/header.go

{"message":"Bad credentials","documentation_url":"https://docs.github.com/rest"}
```

In the above example, we have created a `GET` request to the [`https://api.github.com/users/mr-destructive`](https://api.github.com/users/mr-destructive) the last portion is the username, it could be any valid username. The request is to the GitHub API, so it might require API Key/Tokens in the headers, however, if there are certain endpoints that do not require Authorization headers might work just fine.

So, the above code will give `401` error indicating the request has wrong or invalid credentials, if we remove the header, the request will work fine. This is just an example, but headers are useful in working with APIs.

Without adding the header:

```
$ go run web/methods/get/header.go

{"login":"Mr-Destructive","id":40317114,"node_id":"MDQ6VXNlcjQwMzE3MTE0","avatar_url":"https://avatars.githubusercontent.com/u/40317114?v=4","gravatar_id":"","url":"https://api.github.com/users/Mr-Destructive",
... 
"updated_at":"2023-10-10T17:57:22Z"}
```

That's it from the 33rd part of the series, all the source code for the examples are linked in the GitHub on the [100 days of Golang](https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/get/) repository.

[100-days-of-golang](https://github.com/Mr-Destructive/100-days-of-golang)

## References

- [To Unmarshal or Decode](https://dev.to/jpoly1219/to-unmarshal-or-to-decode-json-processing-in-go-explained-5870)
- [Golang JSON tutorial](https://drstearns.github.io/tutorials/gojson/)
- [Golang OmitEmpty](https://www.sohamkamani.com/golang/omitempty/)

## Conclusion

From this article, we explored the `GET` HTTP method in golang. By using a few examples for creating a get request, adding headers, reading response body, the basic use cases were demonstrated.

Hopefully, you found this section helpful, if you have any comments or feedback, please let me know in the comments section or on my social handles. Thank you for reading. Happy Coding :)