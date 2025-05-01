---
article_html: "<h2>Introduction</h2>\n<p>In the second part of our NGINX Survival
  Guide, we dive into the practical aspects of using NGINX to serve web applications.
  This section will guide you through the essential tasks of setting up a basic HTTP
  server, configuring NGINX to serve content from custom directories, and using it
  as a reverse proxy to forward requests to backend servers.</p>\n<p>NGINX is a versatile
  web server that can be used to serve applications in a variety of ways, from simple
  web servers to complex proxy configurations. NGINX can be used to serve static HTML
  content, proxy requests to a backend server, or load balance traffic across multiple
  servers. In this guide, we'll explore the different ways to use NGINX to serve applications,
  including setting up a simple HTTP server, serving content from custom directories,
  and using it to load balance traffic across multiple upstream servers.</p>\n<h2>Simple
  HTTP Server</h2>\n<p>NGINX serves as the default HTTP server on port 80 of your
  local machine if NGINX is properly installed and running on your system. If you
  head on to the localhost, you will see the default NGINX HTML page like the one
  below:</p>\n<p><img src=\"https://meetgor-cdn.pages.dev/nginx-survival-guide/nginx-default-page.png\"
  alt=\"NGINX Default Page\" /></p>\n<p>This is the default HTML page served by NGINX
  as per the configuration in the <code>/etc/nginx/nginx.conf</code> file. The default
  folder for NGINX to serve HTML content is located at <code>/usr/share/nginx/html/index.html</code>
  , If you change the contents of this file and restart NGINX, the http server will
  load the new HTML content.</p>\n<p>Let's first look, at how we can serve a simple
  http message within the configuration file in NGINX.</p>\n<h2>Serving simple text</h2>\n<p>We
  will try to write our simple HTTP server from scratch, so it would be nice to empty
  the existing <code>/etc/nginx/nginx.conf</code> file or use other ports to serve
  the content rather than the default <code>127.0.0.1:80</code> port.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"k\">http</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"kn\">server</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"kn\">listen</span><span class=\"w\"> </span><span class=\"mi\">8000</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">return</span><span class=\"w\"> </span><span class=\"mi\">200</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Hello,</span><span class=\"w\"> </span><span
  class=\"s\">World!\\n&quot;</span><span class=\"p\">;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  above config will serve the text <code>Hello, World!</code> when there is a request
  to the URL <code>127.0.0.1:8000</code> or <code>localhost:8000</code> You can change
  the port per your requirements and even add a <code>server_name</code> for your
  domain name.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  curl http://127.0.0.1:8000 \nHello, World!\n\n\n$ curl -i http://127.0.0.1:8000\nHTTP/1.1
  <span class=\"m\">200</span> OK\nServer: nginx/1.18.0 <span class=\"o\">(</span>Ubuntu<span
  class=\"o\">)</span>\nDate: Sat, <span class=\"m\">03</span> Feb <span class=\"m\">2024</span>
  <span class=\"m\">11</span>:41:16 GMT\nContent-Type: application/octet-stream\nContent-Length:
  <span class=\"m\">14</span>\nConnection: keep-alive\n\nHello, World!\n</pre></div>\n\n</pre>\n\n<p>As
  we can see the NGINX served the HTTP content when the request was made to port 8000
  on the localhost.</p>\n<h2>Serving from a custom path/folder</h2>\n<p>But things
  are not like these in the real world, we need to serve an entire directory of HTML
  pages. We need to add the <code>root</code> directive with the path to the folder
  where our HTML content resides. The path should have the <code>index.html</code>
  file as the starting point of the request.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"k\">http</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"kn\">server</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"kn\">listen</span><span class=\"w\"> </span><span class=\"mi\">8000</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">root</span><span class=\"w\"> </span><span class=\"s\">/srv/techstructive-blog</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">index</span><span class=\"w\"> </span><span class=\"s\">index.html</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p><strong>NOTE:
  The path to the HTML content needs to be accessible and the Nginx process should
  have the read permission to serve the contents.</strong></p>\n<p>It is commonly
  recommended to store HTML/static content files in directories such as <code>/srv</code>
  or <code>/var/www</code>. These paths follow conventions for serving static files
  and web applications in Unix-type operating systems. While it's not a strict requirement,
  adhering to these conventions can improve the organization and maintainability of
  web content.</p>\n<h2>Serving from a web server</h2>\n<p>If you already have a web
  server running in a port on your system, you could use Nginx as a gateway to the
  application instead of exposing your application to the internet.</p>\n<p>We could
  use the <a href=\"https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_pass\">proxy_pass</a>
  directive in the location setting to specify which URL to pass the request to, the
  <code>listen</code> will forward the request to the proxy specified in the location
  directive.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>nginx</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">http</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kn\">server</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"kn\">listen</span><span class=\"w\"> </span><span class=\"mi\">80</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"kn\">location</span><span class=\"w\"> </span><span class=\"s\">/</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"kn\">proxy_pass</span><span class=\"w\"> </span><span class=\"s\">http://localhost:8001</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, the NGINX listens to port 80 in the local system and sends the
  request to the localhost at port 8001. The <code>proxy_pass</code> is used to specify
  the URL to redirect the request to.</p>\n<ul>\n<li>\n<p><strong>listen 80:</strong>
  Nginx listens for incoming requests on port 80, the standard HTTP port.</p>\n</li>\n<li>\n<p><strong>location
  /:</strong> This directive matches all incoming requests, regardless of the path.</p>\n</li>\n<li>\n<p><strong>proxy_pass</strong><a
  href=\"http://localhost:8001\"><strong>http://localhost:8001</strong></a><strong>:</strong>
  Requests are forwarded to the web application running on <a href=\"http://localhost\">localhost</a>
  at port 8001.</p>\n</li>\n</ul>\n<p>This example configuration is a basic building
  block for setting up more complex proxy configurations with NGINX.</p>\n<h2>Serving
  from Multiple Upstream Servers</h2>\n<p>NGINX can also serve content from multiple
  upstream servers, balancing the load between them. This is useful for high-traffic
  applications that require multiple backend servers to handle the load.</p>\n<p>What
  are upstream servers, you might ask, well in the context of NGINX, upstream servers
  refer to backend servers that handle the actual processing of requests. NGINX acts
  as a gateway, forwarding incoming requests to these upstream servers. This setup
  allows NGINX to manage the traffic efficiently and distribute it among multiple
  servers, which can be particularly beneficial for high-traffic applications. . For
  example, you might have your application running on <a href=\"http://localhost:8001\"><code>localhost:8001</code></a>
  and <a href=\"http://localhost:8002\"><code>localhost:8002</code></a>.</p>\n<p>Here’s
  an example configuration:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"k\">http</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"kn\">upstream</span><span
  class=\"w\"> </span><span class=\"s\">myapp</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">server</span><span class=\"w\"> </span><span class=\"s\">backend1.example.com</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">server</span><span class=\"w\"> </span><span class=\"s\">backend2.example.com</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">server</span><span class=\"w\"> </span><span class=\"s\">backend3.example.com</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">    </span><span
  class=\"kn\">server</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"kn\">listen</span><span
  class=\"w\"> </span><span class=\"mi\">80</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">        </span><span class=\"kn\">location</span><span
  class=\"w\"> </span><span class=\"s\">/</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">            </span><span class=\"kn\">proxy_pass</span><span
  class=\"w\"> </span><span class=\"s\">http://myapp</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  this configuration:</p>\n<ul>\n<li>\n<p>The <code>upstream</code> block defines
  a named group of backend servers (<code>myapp</code>).</p>\n</li>\n<li>\n<p>The
  <code>server</code> block listens on port 80 and proxies requests to the upstream
  group defined earlier.</p>\n</li>\n<li>\n<p><code>upstream myapp</code>: This directive
  creates a group of backend servers named <code>myapp</code>.</p>\n</li>\n<li>\n<p><a
  href=\"http://backend1.example.com\"><code>server backend1.example.com</code></a>
  : These directives list the backend servers that will handle the requests. These
  can be specified by hostname, IP address, or combination.</p>\n</li>\n<li>\n<p><code>proxy_pass</code>
  <a href=\"http://myapp\"><code>http://myapp</code></a>: This directive tells NGINX
  to forward incoming requests to the <code>myapp</code> upstream group.</p>\n</li>\n</ul>\n<h3>Why
  Use Upstream Servers?</h3>\n<p>Using upstream servers has several advantages:</p>\n<ul>\n<li>\n<p>Scalability:
  By distributing requests across multiple servers, you can handle more traffic and
  scale your application horizontally.</p>\n</li>\n<li>\n<p>Fault Tolerance: If one
  of the backend servers goes down, NGINX can continue to serve requests using the
  remaining servers, ensuring high availability.</p>\n</li>\n<li>\n<p>Load Distribution:
  Upstream servers help in balancing the load, which can improve the performance and
  responsiveness of your web application.</p>\n</li>\n</ul>\n<p>The below configuration
  sets up NGINX to act as a gateway that distributes incoming traffic to multiple
  upstream servers. It defines an upstream block with servers at <a href=\"http://localhost:8001\"><code>localhost:8001</code></a>
  and <a href=\"http://localhost:8002\"><code>localhost:8002</code></a>, and forward
  requests to these servers.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"k\">http</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"kn\">upstream</span><span
  class=\"w\"> </span><span class=\"s\">myapp</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">server</span><span class=\"w\"> </span><span class=\"n\">localhost</span><span
  class=\"p\">:</span><span class=\"mi\">8001</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"kn\">server</span><span
  class=\"w\"> </span><span class=\"n\">localhost</span><span class=\"p\">:</span><span
  class=\"mi\">8002</span><span class=\"p\">;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"w\">    </span><span class=\"kn\">server</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">listen</span><span class=\"w\"> </span><span class=\"mi\">80</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n\n<span class=\"w\">        </span><span
  class=\"kn\">location</span><span class=\"w\"> </span><span class=\"s\">/</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \           </span><span class=\"kn\">proxy_pass</span><span class=\"w\"> </span><span
  class=\"s\">http://myapp</span><span class=\"p\">;</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  provided NGINX configuration sets up an upstream block named <code>myapp</code>
  with two backend servers running on <a href=\"http://localhost\"><code>localhost</code></a>
  at ports 8001 and 8002. The server block listens on port 80 and uses a location
  block to match all incoming requests to the root URL (<code>/</code>). These requests
  are forwarded to the <code>myapp</code> upstream group via the <code>proxy_pass</code>
  directive, allowing NGINX to distribute the requests between the two backend servers,
  effectively balancing the load and enhancing the application's performance and reliability.</p>\n<h2>Conclusion</h2>\n<p>From
  this part of ther series, we have learned how to set up a simple HTTP server to
  serve content from custom directories and using NGINX as a gateway to backend servers,
  which covered essential ways to utilize NGINX for serving web applications.</p>\n<p>That's
  it from this part of the series, we will look into detail how to use NGINX as a
  load balancer and reverse proxy, serving static files, and caching content in the
  next part of the series, where we'll dive deeper into advanced NGINX configurations.\nThank
  you for reading, hopefully you found this helpful. If you have any feedback, questions,
  or queries drop them below in the comments or reach me out directly on my social
  handles.</p>\n<p>Happy Coding :)</p>\n"
cover: ''
date: 2024-07-21
datetime: 2024-07-21 00:00:00+00:00
description: 'NGINX Fundamentals: Setting Up Simple HTTP Servers, Serving Custom Content,
  multiple upstream servers'
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2024-07-21-NGINX-Survival-Guide-P2-web-server.md
html: "<h2>Introduction</h2>\n<p>In the second part of our NGINX Survival Guide, we
  dive into the practical aspects of using NGINX to serve web applications. This section
  will guide you through the essential tasks of setting up a basic HTTP server, configuring
  NGINX to serve content from custom directories, and using it as a reverse proxy
  to forward requests to backend servers.</p>\n<p>NGINX is a versatile web server
  that can be used to serve applications in a variety of ways, from simple web servers
  to complex proxy configurations. NGINX can be used to serve static HTML content,
  proxy requests to a backend server, or load balance traffic across multiple servers.
  In this guide, we'll explore the different ways to use NGINX to serve applications,
  including setting up a simple HTTP server, serving content from custom directories,
  and using it to load balance traffic across multiple upstream servers.</p>\n<h2>Simple
  HTTP Server</h2>\n<p>NGINX serves as the default HTTP server on port 80 of your
  local machine if NGINX is properly installed and running on your system. If you
  head on to the localhost, you will see the default NGINX HTML page like the one
  below:</p>\n<p><img src=\"https://meetgor-cdn.pages.dev/nginx-survival-guide/nginx-default-page.png\"
  alt=\"NGINX Default Page\" /></p>\n<p>This is the default HTML page served by NGINX
  as per the configuration in the <code>/etc/nginx/nginx.conf</code> file. The default
  folder for NGINX to serve HTML content is located at <code>/usr/share/nginx/html/index.html</code>
  , If you change the contents of this file and restart NGINX, the http server will
  load the new HTML content.</p>\n<p>Let's first look, at how we can serve a simple
  http message within the configuration file in NGINX.</p>\n<h2>Serving simple text</h2>\n<p>We
  will try to write our simple HTTP server from scratch, so it would be nice to empty
  the existing <code>/etc/nginx/nginx.conf</code> file or use other ports to serve
  the content rather than the default <code>127.0.0.1:80</code> port.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"k\">http</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"kn\">server</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"kn\">listen</span><span class=\"w\"> </span><span class=\"mi\">8000</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">return</span><span class=\"w\"> </span><span class=\"mi\">200</span><span
  class=\"w\"> </span><span class=\"s\">&quot;Hello,</span><span class=\"w\"> </span><span
  class=\"s\">World!\\n&quot;</span><span class=\"p\">;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  above config will serve the text <code>Hello, World!</code> when there is a request
  to the URL <code>127.0.0.1:8000</code> or <code>localhost:8000</code> You can change
  the port per your requirements and even add a <code>server_name</code> for your
  domain name.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>$
  curl http://127.0.0.1:8000 \nHello, World!\n\n\n$ curl -i http://127.0.0.1:8000\nHTTP/1.1
  <span class=\"m\">200</span> OK\nServer: nginx/1.18.0 <span class=\"o\">(</span>Ubuntu<span
  class=\"o\">)</span>\nDate: Sat, <span class=\"m\">03</span> Feb <span class=\"m\">2024</span>
  <span class=\"m\">11</span>:41:16 GMT\nContent-Type: application/octet-stream\nContent-Length:
  <span class=\"m\">14</span>\nConnection: keep-alive\n\nHello, World!\n</pre></div>\n\n</pre>\n\n<p>As
  we can see the NGINX served the HTTP content when the request was made to port 8000
  on the localhost.</p>\n<h2>Serving from a custom path/folder</h2>\n<p>But things
  are not like these in the real world, we need to serve an entire directory of HTML
  pages. We need to add the <code>root</code> directive with the path to the folder
  where our HTML content resides. The path should have the <code>index.html</code>
  file as the starting point of the request.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"k\">http</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"kn\">server</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"kn\">listen</span><span class=\"w\"> </span><span class=\"mi\">8000</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">root</span><span class=\"w\"> </span><span class=\"s\">/srv/techstructive-blog</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">index</span><span class=\"w\"> </span><span class=\"s\">index.html</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p><strong>NOTE:
  The path to the HTML content needs to be accessible and the Nginx process should
  have the read permission to serve the contents.</strong></p>\n<p>It is commonly
  recommended to store HTML/static content files in directories such as <code>/srv</code>
  or <code>/var/www</code>. These paths follow conventions for serving static files
  and web applications in Unix-type operating systems. While it's not a strict requirement,
  adhering to these conventions can improve the organization and maintainability of
  web content.</p>\n<h2>Serving from a web server</h2>\n<p>If you already have a web
  server running in a port on your system, you could use Nginx as a gateway to the
  application instead of exposing your application to the internet.</p>\n<p>We could
  use the <a href=\"https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_pass\">proxy_pass</a>
  directive in the location setting to specify which URL to pass the request to, the
  <code>listen</code> will forward the request to the proxy specified in the location
  directive.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>nginx</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">http</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kn\">server</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"kn\">listen</span><span class=\"w\"> </span><span class=\"mi\">80</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"kn\">location</span><span class=\"w\"> </span><span class=\"s\">/</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t\t</span><span
  class=\"kn\">proxy_pass</span><span class=\"w\"> </span><span class=\"s\">http://localhost:8001</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, the NGINX listens to port 80 in the local system and sends the
  request to the localhost at port 8001. The <code>proxy_pass</code> is used to specify
  the URL to redirect the request to.</p>\n<ul>\n<li>\n<p><strong>listen 80:</strong>
  Nginx listens for incoming requests on port 80, the standard HTTP port.</p>\n</li>\n<li>\n<p><strong>location
  /:</strong> This directive matches all incoming requests, regardless of the path.</p>\n</li>\n<li>\n<p><strong>proxy_pass</strong><a
  href=\"http://localhost:8001\"><strong>http://localhost:8001</strong></a><strong>:</strong>
  Requests are forwarded to the web application running on <a href=\"http://localhost\">localhost</a>
  at port 8001.</p>\n</li>\n</ul>\n<p>This example configuration is a basic building
  block for setting up more complex proxy configurations with NGINX.</p>\n<h2>Serving
  from Multiple Upstream Servers</h2>\n<p>NGINX can also serve content from multiple
  upstream servers, balancing the load between them. This is useful for high-traffic
  applications that require multiple backend servers to handle the load.</p>\n<p>What
  are upstream servers, you might ask, well in the context of NGINX, upstream servers
  refer to backend servers that handle the actual processing of requests. NGINX acts
  as a gateway, forwarding incoming requests to these upstream servers. This setup
  allows NGINX to manage the traffic efficiently and distribute it among multiple
  servers, which can be particularly beneficial for high-traffic applications. . For
  example, you might have your application running on <a href=\"http://localhost:8001\"><code>localhost:8001</code></a>
  and <a href=\"http://localhost:8002\"><code>localhost:8002</code></a>.</p>\n<p>Here’s
  an example configuration:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"k\">http</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"kn\">upstream</span><span
  class=\"w\"> </span><span class=\"s\">myapp</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">server</span><span class=\"w\"> </span><span class=\"s\">backend1.example.com</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">server</span><span class=\"w\"> </span><span class=\"s\">backend2.example.com</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">server</span><span class=\"w\"> </span><span class=\"s\">backend3.example.com</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">    </span><span
  class=\"kn\">server</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"kn\">listen</span><span
  class=\"w\"> </span><span class=\"mi\">80</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">        </span><span class=\"kn\">location</span><span
  class=\"w\"> </span><span class=\"s\">/</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">            </span><span class=\"kn\">proxy_pass</span><span
  class=\"w\"> </span><span class=\"s\">http://myapp</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  this configuration:</p>\n<ul>\n<li>\n<p>The <code>upstream</code> block defines
  a named group of backend servers (<code>myapp</code>).</p>\n</li>\n<li>\n<p>The
  <code>server</code> block listens on port 80 and proxies requests to the upstream
  group defined earlier.</p>\n</li>\n<li>\n<p><code>upstream myapp</code>: This directive
  creates a group of backend servers named <code>myapp</code>.</p>\n</li>\n<li>\n<p><a
  href=\"http://backend1.example.com\"><code>server backend1.example.com</code></a>
  : These directives list the backend servers that will handle the requests. These
  can be specified by hostname, IP address, or combination.</p>\n</li>\n<li>\n<p><code>proxy_pass</code>
  <a href=\"http://myapp\"><code>http://myapp</code></a>: This directive tells NGINX
  to forward incoming requests to the <code>myapp</code> upstream group.</p>\n</li>\n</ul>\n<h3>Why
  Use Upstream Servers?</h3>\n<p>Using upstream servers has several advantages:</p>\n<ul>\n<li>\n<p>Scalability:
  By distributing requests across multiple servers, you can handle more traffic and
  scale your application horizontally.</p>\n</li>\n<li>\n<p>Fault Tolerance: If one
  of the backend servers goes down, NGINX can continue to serve requests using the
  remaining servers, ensuring high availability.</p>\n</li>\n<li>\n<p>Load Distribution:
  Upstream servers help in balancing the load, which can improve the performance and
  responsiveness of your web application.</p>\n</li>\n</ul>\n<p>The below configuration
  sets up NGINX to act as a gateway that distributes incoming traffic to multiple
  upstream servers. It defines an upstream block with servers at <a href=\"http://localhost:8001\"><code>localhost:8001</code></a>
  and <a href=\"http://localhost:8002\"><code>localhost:8002</code></a>, and forward
  requests to these servers.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"k\">http</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"kn\">upstream</span><span
  class=\"w\"> </span><span class=\"s\">myapp</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">server</span><span class=\"w\"> </span><span class=\"n\">localhost</span><span
  class=\"p\">:</span><span class=\"mi\">8001</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"kn\">server</span><span
  class=\"w\"> </span><span class=\"n\">localhost</span><span class=\"p\">:</span><span
  class=\"mi\">8002</span><span class=\"p\">;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"w\">    </span><span class=\"kn\">server</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"kn\">listen</span><span class=\"w\"> </span><span class=\"mi\">80</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n\n<span class=\"w\">        </span><span
  class=\"kn\">location</span><span class=\"w\"> </span><span class=\"s\">/</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \           </span><span class=\"kn\">proxy_pass</span><span class=\"w\"> </span><span
  class=\"s\">http://myapp</span><span class=\"p\">;</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  provided NGINX configuration sets up an upstream block named <code>myapp</code>
  with two backend servers running on <a href=\"http://localhost\"><code>localhost</code></a>
  at ports 8001 and 8002. The server block listens on port 80 and uses a location
  block to match all incoming requests to the root URL (<code>/</code>). These requests
  are forwarded to the <code>myapp</code> upstream group via the <code>proxy_pass</code>
  directive, allowing NGINX to distribute the requests between the two backend servers,
  effectively balancing the load and enhancing the application's performance and reliability.</p>\n<h2>Conclusion</h2>\n<p>From
  this part of ther series, we have learned how to set up a simple HTTP server to
  serve content from custom directories and using NGINX as a gateway to backend servers,
  which covered essential ways to utilize NGINX for serving web applications.</p>\n<p>That's
  it from this part of the series, we will look into detail how to use NGINX as a
  load balancer and reverse proxy, serving static files, and caching content in the
  next part of the series, where we'll dive deeper into advanced NGINX configurations.\nThank
  you for reading, hopefully you found this helpful. If you have any feedback, questions,
  or queries drop them below in the comments or reach me out directly on my social
  handles.</p>\n<p>Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/nginx-survival-guide/nginx-sg-2.png
long_description: In the second part of our NGINX Survival Guide, we dive into the
  practical aspects of using NGINX to serve web applications. This section will guide
  you through the essential tasks of setting up a basic HTTP server, configuring NGINX
  to serve content
now: 2025-05-01 18:11:33.314653
path: blog/posts/2024-07-21-NGINX-Survival-Guide-P2-web-server.md
prevnext: null
series:
- nginx-survival-guide
slug: nginx-02-web-servers
status: published
tags:
- nginx
- web-development
templateKey: blog-post
title: 'NGINX Survival Guide: Serving Web Applications'
today: 2025-05-01
---

## Introduction

In the second part of our NGINX Survival Guide, we dive into the practical aspects of using NGINX to serve web applications. This section will guide you through the essential tasks of setting up a basic HTTP server, configuring NGINX to serve content from custom directories, and using it as a reverse proxy to forward requests to backend servers.

NGINX is a versatile web server that can be used to serve applications in a variety of ways, from simple web servers to complex proxy configurations. NGINX can be used to serve static HTML content, proxy requests to a backend server, or load balance traffic across multiple servers. In this guide, we'll explore the different ways to use NGINX to serve applications, including setting up a simple HTTP server, serving content from custom directories, and using it to load balance traffic across multiple upstream servers.

## Simple HTTP Server

NGINX serves as the default HTTP server on port 80 of your local machine if NGINX is properly installed and running on your system. If you head on to the localhost, you will see the default NGINX HTML page like the one below:

![NGINX Default Page](https://meetgor-cdn.pages.dev/nginx-survival-guide/nginx-default-page.png)

This is the default HTML page served by NGINX as per the configuration in the `/etc/nginx/nginx.conf` file. The default folder for NGINX to serve HTML content is located at `/usr/share/nginx/html/index.html` , If you change the contents of this file and restart NGINX, the http server will load the new HTML content.

Let's first look, at how we can serve a simple http message within the configuration file in NGINX.

## Serving simple text

We will try to write our simple HTTP server from scratch, so it would be nice to empty the existing `/etc/nginx/nginx.conf` file or use other ports to serve the content rather than the default `127.0.0.1:80` port.

```nginx
http {
    server {
        listen 8000;
        return 200 "Hello, World!\n";
    }
}
```

The above config will serve the text `Hello, World!` when there is a request to the URL `127.0.0.1:8000` or `localhost:8000` You can change the port per your requirements and even add a `server_name` for your domain name.

```bash
$ curl http://127.0.0.1:8000 
Hello, World!


$ curl -i http://127.0.0.1:8000
HTTP/1.1 200 OK
Server: nginx/1.18.0 (Ubuntu)
Date: Sat, 03 Feb 2024 11:41:16 GMT
Content-Type: application/octet-stream
Content-Length: 14
Connection: keep-alive

Hello, World!
```

As we can see the NGINX served the HTTP content when the request was made to port 8000 on the localhost.

## Serving from a custom path/folder

But things are not like these in the real world, we need to serve an entire directory of HTML pages. We need to add the `root` directive with the path to the folder where our HTML content resides. The path should have the `index.html` file as the starting point of the request.

```nginx
http {
    server {
        listen 8000;
        root /srv/techstructive-blog;
        index index.html;
    }
}
```

**NOTE: The path to the HTML content needs to be accessible and the Nginx process should have the read permission to serve the contents.**

It is commonly recommended to store HTML/static content files in directories such as `/srv` or `/var/www`. These paths follow conventions for serving static files and web applications in Unix-type operating systems. While it's not a strict requirement, adhering to these conventions can improve the organization and maintainability of web content.

## Serving from a web server

If you already have a web server running in a port on your system, you could use Nginx as a gateway to the application instead of exposing your application to the internet.

We could use the [proxy\_pass](https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_pass) directive in the location setting to specify which URL to pass the request to, the `listen` will forward the request to the proxy specified in the location directive.

```nginx
http {
	server {
		listen 80;
		location / {
			proxy_pass http://localhost:8001;
		}
	}
}
```

In the above example, the NGINX listens to port 80 in the local system and sends the request to the localhost at port 8001. The `proxy_pass` is used to specify the URL to redirect the request to.

* **listen 80:** Nginx listens for incoming requests on port 80, the standard HTTP port.
    
* **location /:** This directive matches all incoming requests, regardless of the path.
    
* **proxy\_pass**[**http://localhost:8001**](http://localhost:8001)**:** Requests are forwarded to the web application running on [localhost](http://localhost) at port 8001.
    

This example configuration is a basic building block for setting up more complex proxy configurations with NGINX.

## Serving from Multiple Upstream Servers

NGINX can also serve content from multiple upstream servers, balancing the load between them. This is useful for high-traffic applications that require multiple backend servers to handle the load.

What are upstream servers, you might ask, well in the context of NGINX, upstream servers refer to backend servers that handle the actual processing of requests. NGINX acts as a gateway, forwarding incoming requests to these upstream servers. This setup allows NGINX to manage the traffic efficiently and distribute it among multiple servers, which can be particularly beneficial for high-traffic applications. . For example, you might have your application running on [`localhost:8001`](http://localhost:8001) and [`localhost:8002`](http://localhost:8002).

Here’s an example configuration:

```nginx
http {
    upstream myapp {
        server backend1.example.com;
        server backend2.example.com;
        server backend3.example.com;
    }

    server {
        listen 80;

        location / {
            proxy_pass http://myapp;
        }
    }
}
```

In this configuration:

* The `upstream` block defines a named group of backend servers (`myapp`).
    
* The `server` block listens on port 80 and proxies requests to the upstream group defined earlier.
    
* `upstream myapp`: This directive creates a group of backend servers named `myapp`.
    
* [`server backend1.example.com`](http://backend1.example.com) : These directives list the backend servers that will handle the requests. These can be specified by hostname, IP address, or combination.
    
* `proxy_pass` [`http://myapp`](http://myapp): This directive tells NGINX to forward incoming requests to the `myapp` upstream group.
    

### Why Use Upstream Servers?

Using upstream servers has several advantages:

* Scalability: By distributing requests across multiple servers, you can handle more traffic and scale your application horizontally.
    
* Fault Tolerance: If one of the backend servers goes down, NGINX can continue to serve requests using the remaining servers, ensuring high availability.
    
* Load Distribution: Upstream servers help in balancing the load, which can improve the performance and responsiveness of your web application.
    

The below configuration sets up NGINX to act as a gateway that distributes incoming traffic to multiple upstream servers. It defines an upstream block with servers at [`localhost:8001`](http://localhost:8001) and [`localhost:8002`](http://localhost:8002), and forward requests to these servers.

```nginx
http {
    upstream myapp {
        server localhost:8001;
        server localhost:8002;
    }

    server {
        listen 80;

        location / {
            proxy_pass http://myapp;
        }
    }
}
```

The provided NGINX configuration sets up an upstream block named `myapp` with two backend servers running on [`localhost`](http://localhost) at ports 8001 and 8002. The server block listens on port 80 and uses a location block to match all incoming requests to the root URL (`/`). These requests are forwarded to the `myapp` upstream group via the `proxy_pass` directive, allowing NGINX to distribute the requests between the two backend servers, effectively balancing the load and enhancing the application's performance and reliability.

## Conclusion

From this part of ther series, we have learned how to set up a simple HTTP server to serve content from custom directories and using NGINX as a gateway to backend servers, which covered essential ways to utilize NGINX for serving web applications.

That's it from this part of the series, we will look into detail how to use NGINX as a load balancer and reverse proxy, serving static files, and caching content in the next part of the series, where we'll dive deeper into advanced NGINX configurations.
Thank you for reading, hopefully you found this helpful. If you have any feedback, questions, or queries drop them below in the comments or reach me out directly on my social handles.

Happy Coding :)