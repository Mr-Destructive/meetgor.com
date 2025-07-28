{
  "title": "NGINX Survival Guide: Serving Web Applications",
  "status": "published",
  "slug": "nginx-02-web-servers",
  "date": "2025-04-05T12:33:26Z"
}

<h2>Introduction</h2>
<p>In the second part of our NGINX Survival Guide, we dive into the practical aspects of using NGINX to serve web applications. This section will guide you through the essential tasks of setting up a basic HTTP server, configuring NGINX to serve content from custom directories, and using it as a reverse proxy to forward requests to backend servers.</p>
<p>NGINX is a versatile web server that can be used to serve applications in a variety of ways, from simple web servers to complex proxy configurations. NGINX can be used to serve static HTML content, proxy requests to a backend server, or load balance traffic across multiple servers. In this guide, we'll explore the different ways to use NGINX to serve applications, including setting up a simple HTTP server, serving content from custom directories, and using it to load balance traffic across multiple upstream servers.</p>
<h2>Simple HTTP Server</h2>
<p>NGINX serves as the default HTTP server on port 80 of your local machine if NGINX is properly installed and running on your system. If you head on to the localhost, you will see the default NGINX HTML page like the one below:</p>
<p><img src="https://meetgor-cdn.pages.dev/nginx-survival-guide/nginx-default-page.png" alt="NGINX Default Page"></p>
<p>This is the default HTML page served by NGINX as per the configuration in the <code>/etc/nginx/nginx.conf</code> file. The default folder for NGINX to serve HTML content is located at <code>/usr/share/nginx/html/index.html</code> , If you change the contents of this file and restart NGINX, the http server will load the new HTML content.</p>
<p>Let's first look, at how we can serve a simple http message within the configuration file in NGINX.</p>
<h2>Serving simple text</h2>
<p>We will try to write our simple HTTP server from scratch, so it would be nice to empty the existing <code>/etc/nginx/nginx.conf</code> file or use other ports to serve the content rather than the default <code>127.0.0.1:80</code> port.</p>
<pre><code class="language-nginx">http {
    server {
        listen 8000;
        return 200 &quot;Hello, World!
&quot;;
    }
}
</code></pre>
<p>The above config will serve the text <code>Hello, World!</code> when there is a request to the URL <code>127.0.0.1:8000</code> or <code>localhost:8000</code> You can change the port per your requirements and even add a <code>server_name</code> for your domain name.</p>
<pre><code class="language-bash">$ curl http://127.0.0.1:8000 
Hello, World!


$ curl -i http://127.0.0.1:8000
HTTP/1.1 200 OK
Server: nginx/1.18.0 (Ubuntu)
Date: Sat, 03 Feb 2024 11:41:16 GMT
Content-Type: application/octet-stream
Content-Length: 14
Connection: keep-alive

Hello, World!
</code></pre>
<p>As we can see the NGINX served the HTTP content when the request was made to port 8000 on the localhost.</p>
<h2>Serving from a custom path/folder</h2>
<p>But things are not like these in the real world, we need to serve an entire directory of HTML pages. We need to add the <code>root</code> directive with the path to the folder where our HTML content resides. The path should have the <code>index.html</code> file as the starting point of the request.</p>
<pre><code class="language-nginx">http {
    server {
        listen 8000;
        root /srv/techstructive-blog;
        index index.html;
    }
}
</code></pre>
<p><strong>NOTE: The path to the HTML content needs to be accessible and the Nginx process should have the read permission to serve the contents.</strong></p>
<p>It is commonly recommended to store HTML/static content files in directories such as <code>/srv</code> or <code>/var/www</code>. These paths follow conventions for serving static files and web applications in Unix-type operating systems. While it's not a strict requirement, adhering to these conventions can improve the organization and maintainability of web content.</p>
<h2>Serving from a web server</h2>
<p>If you already have a web server running in a port on your system, you could use Nginx as a gateway to the application instead of exposing your application to the internet.</p>
<p>We could use the <a href="https://nginx.org/en/docs/http/ngx_http_proxy_module.html#proxy_pass">proxy_pass</a> directive in the location setting to specify which URL to pass the request to, the <code>listen</code> will forward the request to the proxy specified in the location directive.</p>
<pre><code class="language-nginx">http {
	server {
		listen 80;
		location / {
			proxy_pass http://localhost:8001;
		}
	}
}
</code></pre>
<p>In the above example, the NGINX listens to port 80 in the local system and sends the request to the localhost at port 8001. The <code>proxy_pass</code> is used to specify the URL to redirect the request to.</p>
<ul>
<li>
<p><strong>listen 80:</strong> Nginx listens for incoming requests on port 80, the standard HTTP port.</p>
</li>
<li>
<p><strong>location /:</strong> This directive matches all incoming requests, regardless of the path.</p>
</li>
<li>
<p><strong>proxy_pass</strong><a href="http://localhost:8001"><strong>http://localhost:8001</strong></a><strong>:</strong> Requests are forwarded to the web application running on <a href="http://localhost">localhost</a> at port 8001.</p>
</li>
</ul>
<p>This example configuration is a basic building block for setting up more complex proxy configurations with NGINX.</p>
<h2>Serving from Multiple Upstream Servers</h2>
<p>NGINX can also serve content from multiple upstream servers, balancing the load between them. This is useful for high-traffic applications that require multiple backend servers to handle the load.</p>
<p>What are upstream servers, you might ask, well in the context of NGINX, upstream servers refer to backend servers that handle the actual processing of requests. NGINX acts as a gateway, forwarding incoming requests to these upstream servers. This setup allows NGINX to manage the traffic efficiently and distribute it among multiple servers, which can be particularly beneficial for high-traffic applications. . For example, you might have your application running on <a href="http://localhost:8001"><code>localhost:8001</code></a> and <a href="http://localhost:8002"><code>localhost:8002</code></a>.</p>
<p>Here’s an example configuration:</p>
<pre><code class="language-nginx">http {
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
</code></pre>
<p>In this configuration:</p>
<ul>
<li>
<p>The <code>upstream</code> block defines a named group of backend servers (<code>myapp</code>).</p>
</li>
<li>
<p>The <code>server</code> block listens on port 80 and proxies requests to the upstream group defined earlier.</p>
</li>
<li>
<p><code>upstream myapp</code>: This directive creates a group of backend servers named <code>myapp</code>.</p>
</li>
<li>
<p><a href="http://backend1.example.com"><code>server backend1.example.com</code></a> : These directives list the backend servers that will handle the requests. These can be specified by hostname, IP address, or combination.</p>
</li>
<li>
<p><code>proxy_pass</code> <a href="http://myapp"><code>http://myapp</code></a>: This directive tells NGINX to forward incoming requests to the <code>myapp</code> upstream group.</p>
</li>
</ul>
<h3>Why Use Upstream Servers?</h3>
<p>Using upstream servers has several advantages:</p>
<ul>
<li>
<p>Scalability: By distributing requests across multiple servers, you can handle more traffic and scale your application horizontally.</p>
</li>
<li>
<p>Fault Tolerance: If one of the backend servers goes down, NGINX can continue to serve requests using the remaining servers, ensuring high availability.</p>
</li>
<li>
<p>Load Distribution: Upstream servers help in balancing the load, which can improve the performance and responsiveness of your web application.</p>
</li>
</ul>
<p>The below configuration sets up NGINX to act as a gateway that distributes incoming traffic to multiple upstream servers. It defines an upstream block with servers at <a href="http://localhost:8001"><code>localhost:8001</code></a> and <a href="http://localhost:8002"><code>localhost:8002</code></a>, and forward requests to these servers.</p>
<pre><code class="language-nginx">http {
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
</code></pre>
<p>The provided NGINX configuration sets up an upstream block named <code>myapp</code> with two backend servers running on <a href="http://localhost"><code>localhost</code></a> at ports 8001 and 8002. The server block listens on port 80 and uses a location block to match all incoming requests to the root URL (<code>/</code>). These requests are forwarded to the <code>myapp</code> upstream group via the <code>proxy_pass</code> directive, allowing NGINX to distribute the requests between the two backend servers, effectively balancing the load and enhancing the application's performance and reliability.</p>
<h2>Conclusion</h2>
<p>From this part of ther series, we have learned how to set up a simple HTTP server to serve content from custom directories and using NGINX as a gateway to backend servers, which covered essential ways to utilize NGINX for serving web applications.</p>
<p>That's it from this part of the series, we will look into detail how to use NGINX as a load balancer and reverse proxy, serving static files, and caching content in the next part of the series, where we'll dive deeper into advanced NGINX configurations.
Thank you for reading, hopefully you found this helpful. If you have any feedback, questions, or queries drop them below in the comments or reach me out directly on my social handles.</p>
<p>Happy Coding :)</p>
