---
article_html: "<h2>Introduction</h2>\n<p>NGINX is a tool that can be used as a web
  server, reverse proxy, load balancer, streaming media files, application gateway,
  content caching, and so much more. It can be said to be a Swiss army knife for optimizing
  and securing your web application deployment.</p>\n<p>The series &quot;NGINX Survival
  Guide&quot; will start from the basics and cover the bare minimum required for a
  backend developer to get going with NGINX. I will use Docker widely throughout this
  course as it is a great combination with NGINX to server web applications. However,
  you can use NGINX without docker, and spawn multiple servers.</p>\n<p>The series
  will cover the terminologies of NGINX, configuring NGINX servers, load balancing
  multiple servers, using it as a reverse proxy, and as an API gateway, there will
  be tiny details and some tidbits of doing certain things in a certain constrained
  environment which will make the learning more valuable.</p>\n<h2>What is NGINX</h2>\n<p>NGINX
  (pronounced &quot;engine-x&quot;) is not just a web server, it is a powerful and
  versatile open-source software that wears many hats in the internet world. At its
  core, it functions as a <strong>lightning-fast web server</strong>, its secret weapon
  lies in its <strong>event-driven architecture</strong>, where it handles requests
  asynchronously, allowing it to serve countless users simultaneously without breaking
  a sweat.</p>\n<blockquote>\n<p>NGINX is a popular choice for powering some of the
  <strong>biggest websites and platforms in the world</strong>, demonstrating its
  reliability and scalability.</p>\n</blockquote>\n<p>NGINX's <strong>configurable
  nature</strong> lets you tailor its behavior to your specific needs, whether managing
  traffic flow with load balancing, caching frequently accessed content for faster
  delivery, or even acting as a gateway for your APIs.</p>\n<p>This versatility makes
  NGINX a <strong>powerful tool for building efficient, secure, and scalable web applications</strong>,
  regardless of size or complexity. Hence the need to learn it as a developer and
  especially important for a backend developer.</p>\n<h3>Why NGINX is must learn for
  backend developers</h3>\n<p>Nginx is a highly efficient and performant web server.
  Understanding its configuration and management allows a backend developer to optimize
  server performance, handle large volumes of traffic, and reduce latency.</p>\n<p>In
  microservices architectures, Nginx can serve as an API gateway, managing and routing
  requests between different services. Nginx provides caching mechanisms that enhance
  performance by serving cached content, reducing the load on backend servers.</p>\n<p>Having
  strong fundamentals in NGINX can indeed provide a competitive edge in the job market
  and make a backend developer more versatile in handling various aspects of backend
  web development.</p>\n<h3>Who is using NGINX?</h3>\n<p>Big Tech Companies are using
  NGINX like DropBox, Netfilx, and Cloudflare, among others. Cloudflare used NGINX
  before but it was not enough for them, so they developed their web server/edge proxy
  suited to their needs called Pingora.</p>\n<ul>\n<li>\n<p>Dropbox - <a href=\"https://dropbox.tech/infrastructure/optimizing-web-servers-for-high-throughput-and-low-latency\">Optimizing
  web servers for high throughput and low latency</a></p>\n</li>\n<li>\n<p>Cloudflare
  - <a href=\"https://blog.cloudflare.com/how-we-built-pingora-the-proxy-that-connects-cloudflare-to-the-internet/\">How
  Cloudflare outgrown NGINX and made way to Pingora</a></p>\n</li>\n<li>\n<p>Netflix
  - <a href=\"https://www.nginx.com/blog/tag/netflix/\">NGINX Netflix archives</a></p>\n</li>\n</ul>\n<h2>Installing
  NGINX</h2>\n<h3>Linux</h3>\n<p>There are comprehensive guides for your specific
  flavor/package manager/preferences in the <a href=\"https://docs.nginx.com/nginx/admin-guide/installing-nginx/installing-nginx-open-source/\">official
  documentation</a> of NGINX.</p>\n<p>A couple of common types of installation medium
  instructions are as follows:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># APT</span>\nsudo apt update\nsudo apt install nginx\n\n<span class=\"c1\">#
  YUM</span>\nsudo yum install epel-release\nsudo yum update\nsudo yum install nginx\n</pre></div>\n\n</pre>\n\n<p>Check
  the status of the NGINX service to ensure the installation was successful or not
  with the command:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sudo
  systemctl status nginx\n</pre></div>\n\n</pre>\n\n<p>If this doesn't have any errors
  or fatal messages, the nginx server is up and running on port 80 i.e. on <code>127.0.0.1</code>
  on the system.</p>\n<h3>MacOS</h3>\n<p>The installation on MacOS for NGINX is pretty
  simple with homebrew. The following <a href=\"https://dev.to/arjavdave/installing-nginx-on-mac-46ac\">article</a>
  walks through the steps of the installation:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>brew
  update\nbrew install nginx\nnginx\n</pre></div>\n\n</pre>\n\n<p>If you do not want
  to install it from homebrew, this <a href=\"https://gist.github.com/beatfactor/a093e872824f770a2a0174345cacf171\">gist</a>
  can help install it from the source.</p>\n<h3>Windows</h3>\n<p>For Windows installation,
  you can follow the <a href=\"https://nginx.org/en/docs/windows.html\">guide</a>
  from the official documentation.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># INSTALL the https://nginx.org/en/download.html</span>\n<span class=\"c1\">#
  A Zip file with the name nginx-version.zip will be downlaoded</span>\n<span class=\"c1\">#
  COPY it to the desired location and use that path while unzipping</span>\n<span
  class=\"nb\">cd</span> c:<span class=\"se\">\\</span>\nunzip nginx-1.25.3.zip\n<span
  class=\"nb\">cd</span> nginx-1.25.3\nstart nginx\n</pre></div>\n\n</pre>\n\n<p>You
  can check the status of NGINX if the installation was successful or not with the
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>tasklist
  /fi <span class=\"s2\">&quot;imagename eq nginx.exe&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>This
  should be from the installation section.</p>\n<h2>Understanding the default config</h2>\n<p>When
  you have completed the installation of nginx, you can see the default nginx configuration
  in the file path as <code>/etc/nginx/nginx.conf</code> in Linux/macOS or <code>C:\\nginx\\conf\\nginx</code>
  in Windows.</p>\n<p>The entire nginx config file consists of directives and contexts.
  The key value pairs mentioned below are as <code>key value;</code> are called <strong>directives</strong>
  and the named object i.e. <code>name {}</code> are called <strong>Contexts or Blocks</strong>.</p>\n<p>Block:
  A block directive has the same structure as a simple directive, but instead of the
  semicolon it ends with a set of additional instructions surrounded by braces (<code>{</code>
  and <code>}</code>). You can call it a group of directives.</p>\n<p>Context: If
  a block directive can have other directives inside braces, it is called a context.
  You can call it a group of blocks and directives.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"k\">user</span><span class=\"w\"> </span><span class=\"s\">www-data</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"k\">worker_processes</span><span
  class=\"w\"> </span><span class=\"s\">auto</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"k\">pid</span><span class=\"w\"> </span><span
  class=\"s\">/run/nginx.pid</span><span class=\"p\">;</span><span class=\"w\"></span>\n<span
  class=\"k\">include</span><span class=\"w\"> </span><span class=\"n\">/etc/nginx/modules-enabled/*.conf</span>;<span
  class=\"w\"></span>\n\n<span class=\"k\">events</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"kn\">worker_connections</span><span class=\"w\"> </span><span class=\"mi\">768</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\"># multi_accept on;</span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"k\">http</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">##</span>\n<span
  class=\"w\">\t</span><span class=\"c1\"># Basic Settings</span>\n<span class=\"w\">\t</span><span
  class=\"c1\">##</span>\n\n<span class=\"w\">\t</span><span class=\"kn\">sendfile</span><span
  class=\"w\"> </span><span class=\"no\">on</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kn\">tcp_nopush</span><span
  class=\"w\"> </span><span class=\"no\">on</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kn\">types_hash_max_size</span><span
  class=\"w\"> </span><span class=\"mi\">2048</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\"># server_tokens
  off;</span>\n\n<span class=\"w\">\t</span><span class=\"c1\"># server_names_hash_bucket_size
  64;</span>\n<span class=\"w\">\t</span><span class=\"c1\"># server_name_in_redirect
  off;</span>\n\n<span class=\"w\">\t</span><span class=\"kn\">include</span><span
  class=\"w\"> </span><span class=\"s\">/etc/nginx/mime.types</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kn\">default_type</span><span
  class=\"w\"> </span><span class=\"s\">application/octet-stream</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">##</span>\n<span
  class=\"w\">\t</span><span class=\"c1\"># SSL Settings</span>\n<span class=\"w\">\t</span><span
  class=\"c1\">##</span>\n\n<span class=\"w\">\t</span><span class=\"kn\">ssl_protocols</span><span
  class=\"w\"> </span><span class=\"s\">TLSv1</span><span class=\"w\"> </span><span
  class=\"s\">TLSv1.1</span><span class=\"w\"> </span><span class=\"s\">TLSv1.2</span><span
  class=\"w\"> </span><span class=\"s\">TLSv1.3</span><span class=\"p\">;</span><span
  class=\"w\"> </span><span class=\"c1\"># Dropping SSLv3, ref: POODLE</span>\n<span
  class=\"w\">\t</span><span class=\"kn\">ssl_prefer_server_ciphers</span><span class=\"w\">
  </span><span class=\"no\">on</span><span class=\"p\">;</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">##</span>\n<span class=\"w\">\t</span><span
  class=\"c1\"># Logging Settings</span>\n<span class=\"w\">\t</span><span class=\"c1\">##</span>\n\n<span
  class=\"w\">\t</span><span class=\"kn\">access_log</span><span class=\"w\"> </span><span
  class=\"s\">/var/log/nginx/access.log</span><span class=\"p\">;</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kn\">error_log</span><span class=\"w\"> </span><span
  class=\"s\">/var/log/nginx/error.log</span><span class=\"p\">;</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">##</span>\n<span class=\"w\">\t</span><span
  class=\"c1\"># Gzip Settings</span>\n<span class=\"w\">\t</span><span class=\"c1\">##</span>\n\n<span
  class=\"w\">\t</span><span class=\"kn\">gzip</span><span class=\"w\"> </span><span
  class=\"no\">on</span><span class=\"p\">;</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\"># gzip_vary on;</span>\n<span class=\"w\">\t</span><span
  class=\"c1\"># gzip_proxied any;</span>\n<span class=\"w\">\t</span><span class=\"c1\">#
  gzip_comp_level 6;</span>\n<span class=\"w\">\t</span><span class=\"c1\"># gzip_buffers
  16 8k;</span>\n<span class=\"w\">\t</span><span class=\"c1\"># gzip_http_version
  1.1;</span>\n<span class=\"w\">\t</span><span class=\"c1\"># gzip_types text/plain
  text/css application/json application/javascript text/xml application/xml application/xml+rss
  text/javascript;</span>\n\n<span class=\"w\">\t</span><span class=\"c1\">##</span>\n<span
  class=\"w\">\t</span><span class=\"c1\"># Virtual Host Configs</span>\n<span class=\"w\">\t</span><span
  class=\"c1\">##</span>\n\n<span class=\"w\">\t</span><span class=\"kn\">include</span><span
  class=\"w\"> </span><span class=\"s\">/etc/nginx/conf.d/*.conf</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kn\">include</span><span
  class=\"w\"> </span><span class=\"s\">/etc/nginx/sites-enabled/*</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n\n<span
  class=\"c1\">#mail {</span>\n<span class=\"c1\">#\t# See sample authentication script
  at:</span>\n<span class=\"c1\">#\t# http://wiki.nginx.org/ImapAuthenticateWithApachePhpScript</span>\n<span
  class=\"c1\">#</span>\n<span class=\"c1\">#\t# auth_http localhost/auth.php;</span>\n<span
  class=\"c1\">#\t# pop3_capabilities &quot;TOP&quot; &quot;USER&quot;;</span>\n<span
  class=\"c1\">#\t# imap_capabilities &quot;IMAP4rev1&quot; &quot;UIDPLUS&quot;;</span>\n<span
  class=\"c1\">#</span>\n<span class=\"c1\">#\tserver {</span>\n<span class=\"c1\">#\t\tlisten
  \    localhost:110;</span>\n<span class=\"c1\">#\t\tprotocol   pop3;</span>\n<span
  class=\"c1\">#\t\tproxy      on;</span>\n<span class=\"c1\">#\t}</span>\n<span class=\"c1\">#</span>\n<span
  class=\"c1\">#\tserver {</span>\n<span class=\"c1\">#\t\tlisten     localhost:143;</span>\n<span
  class=\"c1\">#\t\tprotocol   imap;</span>\n<span class=\"c1\">#\t\tproxy      on;</span>\n<span
  class=\"c1\">#\t}</span>\n<span class=\"c1\">#}</span>\n</pre></div>\n\n</pre>\n\n<p>Let's
  break down what some of the directives and context mean here:</p>\n<p>There are
  a few contexts like the global context, the events context, the HTTP context, etc.</p>\n<ul>\n<li>\n<p><code>user</code>
  : The user directive is the user that NGINX will run as, it is often set to www-data.</p>\n</li>\n<li>\n<p><code>pid</code>
  : The pid directive defines the file where the process ID of the main process is
  stored.</p>\n</li>\n<li>\n<p><code>events</code> : The events block has some directives,
  one of them is the <code>worker_connections</code> and others, this block dictates
  how NGINX handles connections and events. It's like setting the rules for how NGINX
  efficiently manages traffic and resources.</p>\n</li>\n<li>\n<p><code>http</code>
  : The HTTP block is used for defining how NGINX handles HTTP-specific settings and
  behaviors. It's like the control panel for orchestrating web traffic and content
  delivery.</p>\n</li>\n</ul>\n<p>I will cover the customization and tweaking of the
  nginx config throughout the series, but for now, understanding the bare bones of
  the default config would be a good spot to be in.</p>\n<h2>Start/Stop/Restarting
  NGINX</h2>\n<p>Starting, stopping, and reloading are the key operations that should
  be known to a developer while developing web applications with nginx. Some of the
  ways to perform these operations with the command line interface are as follows:</p>\n<h3>Start
  NGINX</h3>\n<p>The start command for nginx is usually with the <code>systemctl</code>
  command as :</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sudo
  systemctl start nginx\n</pre></div>\n\n</pre>\n\n<p>This will start the nginx daemon
  and the default port <code>80</code> will be served with the mentioned configuration
  in the <code>nginx.conf</code> file.</p>\n<h3>Stop NGINX</h3>\n<p>To stop the nginx
  daemon, you can simply use the <code>systemctl</code> command or with the nginx
  signal flag as:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sudo
  systemctl stop nginx\n<span class=\"c1\"># OR</span>\nsudo nginx -s stop\n</pre></div>\n\n</pre>\n\n<p>The
  nginx command can take in various signal commands such as :</p>\n<ul>\n<li>\n<p><code>stop</code>
  : To stop the NGINX server (quick shutdown that terminates the server in the middle
  of serving a connection)</p>\n</li>\n<li>\n<p><code>quit</code> : To stop the NGINX
  server gracefully (NGINX will finish serving the open connections before it shuts
  down)</p>\n</li>\n<li>\n<p><code>reload</code> : To reload the NGINX server (generally
  done to pick up changes in config)</p>\n</li>\n<li>\n<p><code>reopen</code> : To
  reopen the log files of an NGINX server (when you want to rotate the log files without
  stopping the NGINX server)</p>\n</li>\n</ul>\n<h3>Reloading</h3>\n<p>If you have
  changed the configuration of the nginx, you will need to restart/reload the nginx
  daemon for it to pick up the latest changes. This can be achieved with the <code>systemctl</code>
  command or with the nginx signal flag of the reload command as:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sudo
  nginx -s reload\n<span class=\"c1\"># OR</span>\nsudo systemctl restart nginx\n</pre></div>\n\n</pre>\n\n<p>These
  are the basic operations we need to perform while working with NGINX in web applications,
  we will see some specific commands in the later parts of the series.</p>\n<h2>Using
  NGINX with Docker</h2>\n<p>NGINX is usually used for deploying web servers, it usually
  acts as a gateway to interacting with the actual web applications and services,
  so it makes sense to use it with docker. We can simply spawn a nginx web server
  with docker using the following command:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>docker
  run -it --rm -d -p <span class=\"m\">8080</span>:80 --name web nginx\n</pre></div>\n\n</pre>\n\n<p>Then
  we can also mount volumes to make your project directory be loaded in the nginx
  web app instead of the default one with the following command:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>docker
  run -it --rm -d -p <span class=\"m\">8080</span>:80 --name web -v ~/mywebapp:/usr/share/nginx/html
  nginx\n</pre></div>\n\n</pre>\n\n<p>In the above command the <code>-v your_path:default_path</code>
  is the flag used to mount your computer's path to the actual container's path in
  this case the default folder where NGINX looks for the index.html file to server
  the content.</p>\n<p>NGINX server with programming languages using Dockerfile -&gt;
  <a href=\"https://github.com/hoalongnatsu/Dockerfile\">GitHub repo</a></p>\n<h2>Conclusion</h2>\n<p>From
  the first introduction of NGINX, we were able to set up NGINX in our system and
  be able to understand the fundamentals of interacting and operating with the NGINX
  server itself. In the next post of this series, we will be diving into spawning
  multiple web servers and making it load-balanced among them.</p>\n<p>Thank you for
  reading, hopefully you found this helpful. If you have any feedback, questions,
  or queries drop them below in the comments or reach me out directly on my social
  handles.</p>\n<p>Happy Coding :)</p>\n"
cover: ''
date: 2024-01-14
datetime: 2024-01-14 00:00:00+00:00
description: 'Exploring NGINX Fundamentals: A Guide for Backend Developers, from the
  Importance of Learning NGINX to Installation and Server Setup'
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2024-01-14-NGINX-Survival-Guide-P1-basics.md
html: "<h2>Introduction</h2>\n<p>NGINX is a tool that can be used as a web server,
  reverse proxy, load balancer, streaming media files, application gateway, content
  caching, and so much more. It can be said to be a Swiss army knife for optimizing
  and securing your web application deployment.</p>\n<p>The series &quot;NGINX Survival
  Guide&quot; will start from the basics and cover the bare minimum required for a
  backend developer to get going with NGINX. I will use Docker widely throughout this
  course as it is a great combination with NGINX to server web applications. However,
  you can use NGINX without docker, and spawn multiple servers.</p>\n<p>The series
  will cover the terminologies of NGINX, configuring NGINX servers, load balancing
  multiple servers, using it as a reverse proxy, and as an API gateway, there will
  be tiny details and some tidbits of doing certain things in a certain constrained
  environment which will make the learning more valuable.</p>\n<h2>What is NGINX</h2>\n<p>NGINX
  (pronounced &quot;engine-x&quot;) is not just a web server, it is a powerful and
  versatile open-source software that wears many hats in the internet world. At its
  core, it functions as a <strong>lightning-fast web server</strong>, its secret weapon
  lies in its <strong>event-driven architecture</strong>, where it handles requests
  asynchronously, allowing it to serve countless users simultaneously without breaking
  a sweat.</p>\n<blockquote>\n<p>NGINX is a popular choice for powering some of the
  <strong>biggest websites and platforms in the world</strong>, demonstrating its
  reliability and scalability.</p>\n</blockquote>\n<p>NGINX's <strong>configurable
  nature</strong> lets you tailor its behavior to your specific needs, whether managing
  traffic flow with load balancing, caching frequently accessed content for faster
  delivery, or even acting as a gateway for your APIs.</p>\n<p>This versatility makes
  NGINX a <strong>powerful tool for building efficient, secure, and scalable web applications</strong>,
  regardless of size or complexity. Hence the need to learn it as a developer and
  especially important for a backend developer.</p>\n<h3>Why NGINX is must learn for
  backend developers</h3>\n<p>Nginx is a highly efficient and performant web server.
  Understanding its configuration and management allows a backend developer to optimize
  server performance, handle large volumes of traffic, and reduce latency.</p>\n<p>In
  microservices architectures, Nginx can serve as an API gateway, managing and routing
  requests between different services. Nginx provides caching mechanisms that enhance
  performance by serving cached content, reducing the load on backend servers.</p>\n<p>Having
  strong fundamentals in NGINX can indeed provide a competitive edge in the job market
  and make a backend developer more versatile in handling various aspects of backend
  web development.</p>\n<h3>Who is using NGINX?</h3>\n<p>Big Tech Companies are using
  NGINX like DropBox, Netfilx, and Cloudflare, among others. Cloudflare used NGINX
  before but it was not enough for them, so they developed their web server/edge proxy
  suited to their needs called Pingora.</p>\n<ul>\n<li>\n<p>Dropbox - <a href=\"https://dropbox.tech/infrastructure/optimizing-web-servers-for-high-throughput-and-low-latency\">Optimizing
  web servers for high throughput and low latency</a></p>\n</li>\n<li>\n<p>Cloudflare
  - <a href=\"https://blog.cloudflare.com/how-we-built-pingora-the-proxy-that-connects-cloudflare-to-the-internet/\">How
  Cloudflare outgrown NGINX and made way to Pingora</a></p>\n</li>\n<li>\n<p>Netflix
  - <a href=\"https://www.nginx.com/blog/tag/netflix/\">NGINX Netflix archives</a></p>\n</li>\n</ul>\n<h2>Installing
  NGINX</h2>\n<h3>Linux</h3>\n<p>There are comprehensive guides for your specific
  flavor/package manager/preferences in the <a href=\"https://docs.nginx.com/nginx/admin-guide/installing-nginx/installing-nginx-open-source/\">official
  documentation</a> of NGINX.</p>\n<p>A couple of common types of installation medium
  instructions are as follows:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># APT</span>\nsudo apt update\nsudo apt install nginx\n\n<span class=\"c1\">#
  YUM</span>\nsudo yum install epel-release\nsudo yum update\nsudo yum install nginx\n</pre></div>\n\n</pre>\n\n<p>Check
  the status of the NGINX service to ensure the installation was successful or not
  with the command:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sudo
  systemctl status nginx\n</pre></div>\n\n</pre>\n\n<p>If this doesn't have any errors
  or fatal messages, the nginx server is up and running on port 80 i.e. on <code>127.0.0.1</code>
  on the system.</p>\n<h3>MacOS</h3>\n<p>The installation on MacOS for NGINX is pretty
  simple with homebrew. The following <a href=\"https://dev.to/arjavdave/installing-nginx-on-mac-46ac\">article</a>
  walks through the steps of the installation:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>brew
  update\nbrew install nginx\nnginx\n</pre></div>\n\n</pre>\n\n<p>If you do not want
  to install it from homebrew, this <a href=\"https://gist.github.com/beatfactor/a093e872824f770a2a0174345cacf171\">gist</a>
  can help install it from the source.</p>\n<h3>Windows</h3>\n<p>For Windows installation,
  you can follow the <a href=\"https://nginx.org/en/docs/windows.html\">guide</a>
  from the official documentation.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># INSTALL the https://nginx.org/en/download.html</span>\n<span class=\"c1\">#
  A Zip file with the name nginx-version.zip will be downlaoded</span>\n<span class=\"c1\">#
  COPY it to the desired location and use that path while unzipping</span>\n<span
  class=\"nb\">cd</span> c:<span class=\"se\">\\</span>\nunzip nginx-1.25.3.zip\n<span
  class=\"nb\">cd</span> nginx-1.25.3\nstart nginx\n</pre></div>\n\n</pre>\n\n<p>You
  can check the status of NGINX if the installation was successful or not with the
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>tasklist
  /fi <span class=\"s2\">&quot;imagename eq nginx.exe&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>This
  should be from the installation section.</p>\n<h2>Understanding the default config</h2>\n<p>When
  you have completed the installation of nginx, you can see the default nginx configuration
  in the file path as <code>/etc/nginx/nginx.conf</code> in Linux/macOS or <code>C:\\nginx\\conf\\nginx</code>
  in Windows.</p>\n<p>The entire nginx config file consists of directives and contexts.
  The key value pairs mentioned below are as <code>key value;</code> are called <strong>directives</strong>
  and the named object i.e. <code>name {}</code> are called <strong>Contexts or Blocks</strong>.</p>\n<p>Block:
  A block directive has the same structure as a simple directive, but instead of the
  semicolon it ends with a set of additional instructions surrounded by braces (<code>{</code>
  and <code>}</code>). You can call it a group of directives.</p>\n<p>Context: If
  a block directive can have other directives inside braces, it is called a context.
  You can call it a group of blocks and directives.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"k\">user</span><span class=\"w\"> </span><span class=\"s\">www-data</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"k\">worker_processes</span><span
  class=\"w\"> </span><span class=\"s\">auto</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"k\">pid</span><span class=\"w\"> </span><span
  class=\"s\">/run/nginx.pid</span><span class=\"p\">;</span><span class=\"w\"></span>\n<span
  class=\"k\">include</span><span class=\"w\"> </span><span class=\"n\">/etc/nginx/modules-enabled/*.conf</span>;<span
  class=\"w\"></span>\n\n<span class=\"k\">events</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"kn\">worker_connections</span><span class=\"w\"> </span><span class=\"mi\">768</span><span
  class=\"p\">;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\"># multi_accept on;</span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"k\">http</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">##</span>\n<span
  class=\"w\">\t</span><span class=\"c1\"># Basic Settings</span>\n<span class=\"w\">\t</span><span
  class=\"c1\">##</span>\n\n<span class=\"w\">\t</span><span class=\"kn\">sendfile</span><span
  class=\"w\"> </span><span class=\"no\">on</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kn\">tcp_nopush</span><span
  class=\"w\"> </span><span class=\"no\">on</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kn\">types_hash_max_size</span><span
  class=\"w\"> </span><span class=\"mi\">2048</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\"># server_tokens
  off;</span>\n\n<span class=\"w\">\t</span><span class=\"c1\"># server_names_hash_bucket_size
  64;</span>\n<span class=\"w\">\t</span><span class=\"c1\"># server_name_in_redirect
  off;</span>\n\n<span class=\"w\">\t</span><span class=\"kn\">include</span><span
  class=\"w\"> </span><span class=\"s\">/etc/nginx/mime.types</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kn\">default_type</span><span
  class=\"w\"> </span><span class=\"s\">application/octet-stream</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"c1\">##</span>\n<span
  class=\"w\">\t</span><span class=\"c1\"># SSL Settings</span>\n<span class=\"w\">\t</span><span
  class=\"c1\">##</span>\n\n<span class=\"w\">\t</span><span class=\"kn\">ssl_protocols</span><span
  class=\"w\"> </span><span class=\"s\">TLSv1</span><span class=\"w\"> </span><span
  class=\"s\">TLSv1.1</span><span class=\"w\"> </span><span class=\"s\">TLSv1.2</span><span
  class=\"w\"> </span><span class=\"s\">TLSv1.3</span><span class=\"p\">;</span><span
  class=\"w\"> </span><span class=\"c1\"># Dropping SSLv3, ref: POODLE</span>\n<span
  class=\"w\">\t</span><span class=\"kn\">ssl_prefer_server_ciphers</span><span class=\"w\">
  </span><span class=\"no\">on</span><span class=\"p\">;</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">##</span>\n<span class=\"w\">\t</span><span
  class=\"c1\"># Logging Settings</span>\n<span class=\"w\">\t</span><span class=\"c1\">##</span>\n\n<span
  class=\"w\">\t</span><span class=\"kn\">access_log</span><span class=\"w\"> </span><span
  class=\"s\">/var/log/nginx/access.log</span><span class=\"p\">;</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"kn\">error_log</span><span class=\"w\"> </span><span
  class=\"s\">/var/log/nginx/error.log</span><span class=\"p\">;</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">##</span>\n<span class=\"w\">\t</span><span
  class=\"c1\"># Gzip Settings</span>\n<span class=\"w\">\t</span><span class=\"c1\">##</span>\n\n<span
  class=\"w\">\t</span><span class=\"kn\">gzip</span><span class=\"w\"> </span><span
  class=\"no\">on</span><span class=\"p\">;</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\"># gzip_vary on;</span>\n<span class=\"w\">\t</span><span
  class=\"c1\"># gzip_proxied any;</span>\n<span class=\"w\">\t</span><span class=\"c1\">#
  gzip_comp_level 6;</span>\n<span class=\"w\">\t</span><span class=\"c1\"># gzip_buffers
  16 8k;</span>\n<span class=\"w\">\t</span><span class=\"c1\"># gzip_http_version
  1.1;</span>\n<span class=\"w\">\t</span><span class=\"c1\"># gzip_types text/plain
  text/css application/json application/javascript text/xml application/xml application/xml+rss
  text/javascript;</span>\n\n<span class=\"w\">\t</span><span class=\"c1\">##</span>\n<span
  class=\"w\">\t</span><span class=\"c1\"># Virtual Host Configs</span>\n<span class=\"w\">\t</span><span
  class=\"c1\">##</span>\n\n<span class=\"w\">\t</span><span class=\"kn\">include</span><span
  class=\"w\"> </span><span class=\"s\">/etc/nginx/conf.d/*.conf</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"kn\">include</span><span
  class=\"w\"> </span><span class=\"s\">/etc/nginx/sites-enabled/*</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n\n<span
  class=\"c1\">#mail {</span>\n<span class=\"c1\">#\t# See sample authentication script
  at:</span>\n<span class=\"c1\">#\t# http://wiki.nginx.org/ImapAuthenticateWithApachePhpScript</span>\n<span
  class=\"c1\">#</span>\n<span class=\"c1\">#\t# auth_http localhost/auth.php;</span>\n<span
  class=\"c1\">#\t# pop3_capabilities &quot;TOP&quot; &quot;USER&quot;;</span>\n<span
  class=\"c1\">#\t# imap_capabilities &quot;IMAP4rev1&quot; &quot;UIDPLUS&quot;;</span>\n<span
  class=\"c1\">#</span>\n<span class=\"c1\">#\tserver {</span>\n<span class=\"c1\">#\t\tlisten
  \    localhost:110;</span>\n<span class=\"c1\">#\t\tprotocol   pop3;</span>\n<span
  class=\"c1\">#\t\tproxy      on;</span>\n<span class=\"c1\">#\t}</span>\n<span class=\"c1\">#</span>\n<span
  class=\"c1\">#\tserver {</span>\n<span class=\"c1\">#\t\tlisten     localhost:143;</span>\n<span
  class=\"c1\">#\t\tprotocol   imap;</span>\n<span class=\"c1\">#\t\tproxy      on;</span>\n<span
  class=\"c1\">#\t}</span>\n<span class=\"c1\">#}</span>\n</pre></div>\n\n</pre>\n\n<p>Let's
  break down what some of the directives and context mean here:</p>\n<p>There are
  a few contexts like the global context, the events context, the HTTP context, etc.</p>\n<ul>\n<li>\n<p><code>user</code>
  : The user directive is the user that NGINX will run as, it is often set to www-data.</p>\n</li>\n<li>\n<p><code>pid</code>
  : The pid directive defines the file where the process ID of the main process is
  stored.</p>\n</li>\n<li>\n<p><code>events</code> : The events block has some directives,
  one of them is the <code>worker_connections</code> and others, this block dictates
  how NGINX handles connections and events. It's like setting the rules for how NGINX
  efficiently manages traffic and resources.</p>\n</li>\n<li>\n<p><code>http</code>
  : The HTTP block is used for defining how NGINX handles HTTP-specific settings and
  behaviors. It's like the control panel for orchestrating web traffic and content
  delivery.</p>\n</li>\n</ul>\n<p>I will cover the customization and tweaking of the
  nginx config throughout the series, but for now, understanding the bare bones of
  the default config would be a good spot to be in.</p>\n<h2>Start/Stop/Restarting
  NGINX</h2>\n<p>Starting, stopping, and reloading are the key operations that should
  be known to a developer while developing web applications with nginx. Some of the
  ways to perform these operations with the command line interface are as follows:</p>\n<h3>Start
  NGINX</h3>\n<p>The start command for nginx is usually with the <code>systemctl</code>
  command as :</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sudo
  systemctl start nginx\n</pre></div>\n\n</pre>\n\n<p>This will start the nginx daemon
  and the default port <code>80</code> will be served with the mentioned configuration
  in the <code>nginx.conf</code> file.</p>\n<h3>Stop NGINX</h3>\n<p>To stop the nginx
  daemon, you can simply use the <code>systemctl</code> command or with the nginx
  signal flag as:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sudo
  systemctl stop nginx\n<span class=\"c1\"># OR</span>\nsudo nginx -s stop\n</pre></div>\n\n</pre>\n\n<p>The
  nginx command can take in various signal commands such as :</p>\n<ul>\n<li>\n<p><code>stop</code>
  : To stop the NGINX server (quick shutdown that terminates the server in the middle
  of serving a connection)</p>\n</li>\n<li>\n<p><code>quit</code> : To stop the NGINX
  server gracefully (NGINX will finish serving the open connections before it shuts
  down)</p>\n</li>\n<li>\n<p><code>reload</code> : To reload the NGINX server (generally
  done to pick up changes in config)</p>\n</li>\n<li>\n<p><code>reopen</code> : To
  reopen the log files of an NGINX server (when you want to rotate the log files without
  stopping the NGINX server)</p>\n</li>\n</ul>\n<h3>Reloading</h3>\n<p>If you have
  changed the configuration of the nginx, you will need to restart/reload the nginx
  daemon for it to pick up the latest changes. This can be achieved with the <code>systemctl</code>
  command or with the nginx signal flag of the reload command as:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sudo
  nginx -s reload\n<span class=\"c1\"># OR</span>\nsudo systemctl restart nginx\n</pre></div>\n\n</pre>\n\n<p>These
  are the basic operations we need to perform while working with NGINX in web applications,
  we will see some specific commands in the later parts of the series.</p>\n<h2>Using
  NGINX with Docker</h2>\n<p>NGINX is usually used for deploying web servers, it usually
  acts as a gateway to interacting with the actual web applications and services,
  so it makes sense to use it with docker. We can simply spawn a nginx web server
  with docker using the following command:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>docker
  run -it --rm -d -p <span class=\"m\">8080</span>:80 --name web nginx\n</pre></div>\n\n</pre>\n\n<p>Then
  we can also mount volumes to make your project directory be loaded in the nginx
  web app instead of the default one with the following command:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>docker
  run -it --rm -d -p <span class=\"m\">8080</span>:80 --name web -v ~/mywebapp:/usr/share/nginx/html
  nginx\n</pre></div>\n\n</pre>\n\n<p>In the above command the <code>-v your_path:default_path</code>
  is the flag used to mount your computer's path to the actual container's path in
  this case the default folder where NGINX looks for the index.html file to server
  the content.</p>\n<p>NGINX server with programming languages using Dockerfile -&gt;
  <a href=\"https://github.com/hoalongnatsu/Dockerfile\">GitHub repo</a></p>\n<h2>Conclusion</h2>\n<p>From
  the first introduction of NGINX, we were able to set up NGINX in our system and
  be able to understand the fundamentals of interacting and operating with the NGINX
  server itself. In the next post of this series, we will be diving into spawning
  multiple web servers and making it load-balanced among them.</p>\n<p>Thank you for
  reading, hopefully you found this helpful. If you have any feedback, questions,
  or queries drop them below in the comments or reach me out directly on my social
  handles.</p>\n<p>Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/nginx-survival-guide/nginx-sg-1.png
long_description: NGINX is a tool that can be used as a web server, reverse proxy,
  load balancer, streaming media files, application gateway, content caching, and
  so much more. It can be said to be a Swiss army knife for optimizing and securing
  your web application de
now: 2025-05-01 18:11:33.312780
path: blog/posts/2024-01-14-NGINX-Survival-Guide-P1-basics.md
prevnext: null
series:
- nginx-survival-guide
series_description: 'NGINX Survival Guide: A Guide for Backend Developers, understanding
  the bare minimum fundamentals to get going with NGINX.'
slug: nginx-01-basics
status: published
tags:
- nginx
- web-development
templateKey: blog-post
title: NGINX Basics and Setup
today: 2025-05-01
---

## Introduction

NGINX is a tool that can be used as a web server, reverse proxy, load balancer, streaming media files, application gateway, content caching, and so much more. It can be said to be a Swiss army knife for optimizing and securing your web application deployment.

The series "NGINX Survival Guide" will start from the basics and cover the bare minimum required for a backend developer to get going with NGINX. I will use Docker widely throughout this course as it is a great combination with NGINX to server web applications. However, you can use NGINX without docker, and spawn multiple servers.

The series will cover the terminologies of NGINX, configuring NGINX servers, load balancing multiple servers, using it as a reverse proxy, and as an API gateway, there will be tiny details and some tidbits of doing certain things in a certain constrained environment which will make the learning more valuable.

## What is NGINX

NGINX (pronounced "engine-x") is not just a web server, it is a powerful and versatile open-source software that wears many hats in the internet world. At its core, it functions as a **lightning-fast web server**, its secret weapon lies in its **event-driven architecture**, where it handles requests asynchronously, allowing it to serve countless users simultaneously without breaking a sweat.

> NGINX is a popular choice for powering some of the **biggest websites and platforms in the world**, demonstrating its reliability and scalability.

NGINX's **configurable nature** lets you tailor its behavior to your specific needs, whether managing traffic flow with load balancing, caching frequently accessed content for faster delivery, or even acting as a gateway for your APIs.

This versatility makes NGINX a **powerful tool for building efficient, secure, and scalable web applications**, regardless of size or complexity. Hence the need to learn it as a developer and especially important for a backend developer.

### Why NGINX is must learn for backend developers

Nginx is a highly efficient and performant web server. Understanding its configuration and management allows a backend developer to optimize server performance, handle large volumes of traffic, and reduce latency.

In microservices architectures, Nginx can serve as an API gateway, managing and routing requests between different services. Nginx provides caching mechanisms that enhance performance by serving cached content, reducing the load on backend servers.

Having strong fundamentals in NGINX can indeed provide a competitive edge in the job market and make a backend developer more versatile in handling various aspects of backend web development.

### Who is using NGINX?

Big Tech Companies are using NGINX like DropBox, Netfilx, and Cloudflare, among others. Cloudflare used NGINX before but it was not enough for them, so they developed their web server/edge proxy suited to their needs called Pingora.

* Dropbox - [Optimizing web servers for high throughput and low latency](https://dropbox.tech/infrastructure/optimizing-web-servers-for-high-throughput-and-low-latency)

* Cloudflare - [How Cloudflare outgrown NGINX and made way to Pingora](https://blog.cloudflare.com/how-we-built-pingora-the-proxy-that-connects-cloudflare-to-the-internet/)

* Netflix - [NGINX Netflix archives](https://www.nginx.com/blog/tag/netflix/)


## Installing NGINX

### Linux

There are comprehensive guides for your specific flavor/package manager/preferences in the [official documentation](https://docs.nginx.com/nginx/admin-guide/installing-nginx/installing-nginx-open-source/) of NGINX.

A couple of common types of installation medium instructions are as follows:

```bash
# APT
sudo apt update
sudo apt install nginx

# YUM
sudo yum install epel-release
sudo yum update
sudo yum install nginx
```

Check the status of the NGINX service to ensure the installation was successful or not with the command:

```bash
sudo systemctl status nginx
```

If this doesn't have any errors or fatal messages, the nginx server is up and running on port 80 i.e. on `127.0.0.1` on the system.

### MacOS

The installation on MacOS for NGINX is pretty simple with homebrew. The following [article](https://dev.to/arjavdave/installing-nginx-on-mac-46ac) walks through the steps of the installation:

```bash
brew update
brew install nginx
nginx
```

If you do not want to install it from homebrew, this [gist](https://gist.github.com/beatfactor/a093e872824f770a2a0174345cacf171) can help install it from the source.

### Windows

For Windows installation, you can follow the [guide](https://nginx.org/en/docs/windows.html) from the official documentation.

```bash
# INSTALL the https://nginx.org/en/download.html
# A Zip file with the name nginx-version.zip will be downlaoded
# COPY it to the desired location and use that path while unzipping
cd c:\
unzip nginx-1.25.3.zip
cd nginx-1.25.3
start nginx
```

You can check the status of NGINX if the installation was successful or not with the command:

```bash
tasklist /fi "imagename eq nginx.exe"
```

This should be from the installation section.

## Understanding the default config

When you have completed the installation of nginx, you can see the default nginx configuration in the file path as `/etc/nginx/nginx.conf` in Linux/macOS or `C:\nginx\conf\nginx` in Windows.

The entire nginx config file consists of directives and contexts. The key value pairs mentioned below are as `key value;` are called **directives** and the named object i.e. `name {}` are called **Contexts or Blocks**.

Block: A block directive has the same structure as a simple directive, but instead of the semicolon it ends with a set of additional instructions surrounded by braces (`{` and `}`). You can call it a group of directives.

Context: If a block directive can have other directives inside braces, it is called a context. You can call it a group of blocks and directives.

```nginx
user www-data;
worker_processes auto;
pid /run/nginx.pid;
include /etc/nginx/modules-enabled/*.conf;

events {
	worker_connections 768;
	# multi_accept on;
}

http {

	##
	# Basic Settings
	##

	sendfile on;
	tcp_nopush on;
	types_hash_max_size 2048;
	# server_tokens off;

	# server_names_hash_bucket_size 64;
	# server_name_in_redirect off;

	include /etc/nginx/mime.types;
	default_type application/octet-stream;

	##
	# SSL Settings
	##

	ssl_protocols TLSv1 TLSv1.1 TLSv1.2 TLSv1.3; # Dropping SSLv3, ref: POODLE
	ssl_prefer_server_ciphers on;

	##
	# Logging Settings
	##

	access_log /var/log/nginx/access.log;
	error_log /var/log/nginx/error.log;

	##
	# Gzip Settings
	##

	gzip on;

	# gzip_vary on;
	# gzip_proxied any;
	# gzip_comp_level 6;
	# gzip_buffers 16 8k;
	# gzip_http_version 1.1;
	# gzip_types text/plain text/css application/json application/javascript text/xml application/xml application/xml+rss text/javascript;

	##
	# Virtual Host Configs
	##

	include /etc/nginx/conf.d/*.conf;
	include /etc/nginx/sites-enabled/*;
}


#mail {
#	# See sample authentication script at:
#	# http://wiki.nginx.org/ImapAuthenticateWithApachePhpScript
#
#	# auth_http localhost/auth.php;
#	# pop3_capabilities "TOP" "USER";
#	# imap_capabilities "IMAP4rev1" "UIDPLUS";
#
#	server {
#		listen     localhost:110;
#		protocol   pop3;
#		proxy      on;
#	}
#
#	server {
#		listen     localhost:143;
#		protocol   imap;
#		proxy      on;
#	}
#}
```

Let's break down what some of the directives and context mean here:

There are a few contexts like the global context, the events context, the HTTP context, etc.

* `user` : The user directive is the user that NGINX will run as, it is often set to www-data.

* `pid` : The pid directive defines the file where the process ID of the main process is stored.

* `events` : The events block has some directives, one of them is the `worker_connections` and others, this block dictates how NGINX handles connections and events. It's like setting the rules for how NGINX efficiently manages traffic and resources.

* `http` : The HTTP block is used for defining how NGINX handles HTTP-specific settings and behaviors. It's like the control panel for orchestrating web traffic and content delivery.


I will cover the customization and tweaking of the nginx config throughout the series, but for now, understanding the bare bones of the default config would be a good spot to be in.

## Start/Stop/Restarting NGINX

Starting, stopping, and reloading are the key operations that should be known to a developer while developing web applications with nginx. Some of the ways to perform these operations with the command line interface are as follows:

### Start NGINX

The start command for nginx is usually with the `systemctl` command as :

```bash
sudo systemctl start nginx
```

This will start the nginx daemon and the default port `80` will be served with the mentioned configuration in the `nginx.conf` file.

### Stop NGINX

To stop the nginx daemon, you can simply use the `systemctl` command or with the nginx signal flag as:

```bash
sudo systemctl stop nginx
# OR
sudo nginx -s stop
```

The nginx command can take in various signal commands such as :

* `stop` : To stop the NGINX server (quick shutdown that terminates the server in the middle of serving a connection)

* `quit` : To stop the NGINX server gracefully (NGINX will finish serving the open connections before it shuts down)

* `reload` : To reload the NGINX server (generally done to pick up changes in config)

* `reopen` : To reopen the log files of an NGINX server (when you want to rotate the log files without stopping the NGINX server)


### Reloading

If you have changed the configuration of the nginx, you will need to restart/reload the nginx daemon for it to pick up the latest changes. This can be achieved with the `systemctl` command or with the nginx signal flag of the reload command as:

```bash
sudo nginx -s reload
# OR
sudo systemctl restart nginx
```

These are the basic operations we need to perform while working with NGINX in web applications, we will see some specific commands in the later parts of the series.

## Using NGINX with Docker

NGINX is usually used for deploying web servers, it usually acts as a gateway to interacting with the actual web applications and services, so it makes sense to use it with docker. We can simply spawn a nginx web server with docker using the following command:

```bash
docker run -it --rm -d -p 8080:80 --name web nginx
```

Then we can also mount volumes to make your project directory be loaded in the nginx web app instead of the default one with the following command:

```bash
docker run -it --rm -d -p 8080:80 --name web -v ~/mywebapp:/usr/share/nginx/html nginx
```

In the above command the `-v your_path:default_path` is the flag used to mount your computer's path to the actual container's path in this case the default folder where NGINX looks for the index.html file to server the content.

NGINX server with programming languages using Dockerfile -&gt; [GitHub repo](https://github.com/hoalongnatsu/Dockerfile)

## Conclusion

From the first introduction of NGINX, we were able to set up NGINX in our system and be able to understand the fundamentals of interacting and operating with the NGINX server itself. In the next post of this series, we will be diving into spawning multiple web servers and making it load-balanced among them.

Thank you for reading, hopefully you found this helpful. If you have any feedback, questions, or queries drop them below in the comments or reach me out directly on my social handles.

Happy Coding :)