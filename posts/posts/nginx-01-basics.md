{
  "title": "NGINX Basics and Setup",
  "status": "published",
  "slug": "nginx-01-basics",
  "date": "2025-04-05T12:33:28Z"
}

<h2>Introduction</h2>
<p>NGINX is a tool that can be used as a web server, reverse proxy, load balancer, streaming media files, application gateway, content caching, and so much more. It can be said to be a Swiss army knife for optimizing and securing your web application deployment.</p>
<p>The series &quot;NGINX Survival Guide&quot; will start from the basics and cover the bare minimum required for a backend developer to get going with NGINX. I will use Docker widely throughout this course as it is a great combination with NGINX to server web applications. However, you can use NGINX without docker, and spawn multiple servers.</p>
<p>The series will cover the terminologies of NGINX, configuring NGINX servers, load balancing multiple servers, using it as a reverse proxy, and as an API gateway, there will be tiny details and some tidbits of doing certain things in a certain constrained environment which will make the learning more valuable.</p>
<h2>What is NGINX</h2>
<p>NGINX (pronounced &quot;engine-x&quot;) is not just a web server, it is a powerful and versatile open-source software that wears many hats in the internet world. At its core, it functions as a <strong>lightning-fast web server</strong>, its secret weapon lies in its <strong>event-driven architecture</strong>, where it handles requests asynchronously, allowing it to serve countless users simultaneously without breaking a sweat.</p>
<blockquote>
<p>NGINX is a popular choice for powering some of the <strong>biggest websites and platforms in the world</strong>, demonstrating its reliability and scalability.</p>
</blockquote>
<p>NGINX's <strong>configurable nature</strong> lets you tailor its behavior to your specific needs, whether managing traffic flow with load balancing, caching frequently accessed content for faster delivery, or even acting as a gateway for your APIs.</p>
<p>This versatility makes NGINX a <strong>powerful tool for building efficient, secure, and scalable web applications</strong>, regardless of size or complexity. Hence the need to learn it as a developer and especially important for a backend developer.</p>
<h3>Why NGINX is must learn for backend developers</h3>
<p>Nginx is a highly efficient and performant web server. Understanding its configuration and management allows a backend developer to optimize server performance, handle large volumes of traffic, and reduce latency.</p>
<p>In microservices architectures, Nginx can serve as an API gateway, managing and routing requests between different services. Nginx provides caching mechanisms that enhance performance by serving cached content, reducing the load on backend servers.</p>
<p>Having strong fundamentals in NGINX can indeed provide a competitive edge in the job market and make a backend developer more versatile in handling various aspects of backend web development.</p>
<h3>Who is using NGINX?</h3>
<p>Big Tech Companies are using NGINX like DropBox, Netfilx, and Cloudflare, among others. Cloudflare used NGINX before but it was not enough for them, so they developed their web server/edge proxy suited to their needs called Pingora.</p>
<ul>
<li>
<p>Dropbox - <a href="https://dropbox.tech/infrastructure/optimizing-web-servers-for-high-throughput-and-low-latency">Optimizing web servers for high throughput and low latency</a></p>
</li>
<li>
<p>Cloudflare - <a href="https://blog.cloudflare.com/how-we-built-pingora-the-proxy-that-connects-cloudflare-to-the-internet/">How Cloudflare outgrown NGINX and made way to Pingora</a></p>
</li>
<li>
<p>Netflix - <a href="https://www.nginx.com/blog/tag/netflix/">NGINX Netflix archives</a></p>
</li>
</ul>
<h2>Installing NGINX</h2>
<h3>Linux</h3>
<p>There are comprehensive guides for your specific flavor/package manager/preferences in the <a href="https://docs.nginx.com/nginx/admin-guide/installing-nginx/installing-nginx-open-source/">official documentation</a> of NGINX.</p>
<p>A couple of common types of installation medium instructions are as follows:</p>
<pre><code class="language-bash"># APT
sudo apt update
sudo apt install nginx

# YUM
sudo yum install epel-release
sudo yum update
sudo yum install nginx
</code></pre>
<p>Check the status of the NGINX service to ensure the installation was successful or not with the command:</p>
<pre><code class="language-bash">sudo systemctl status nginx
</code></pre>
<p>If this doesn't have any errors or fatal messages, the nginx server is up and running on port 80 i.e. on <code>127.0.0.1</code> on the system.</p>
<h3>MacOS</h3>
<p>The installation on MacOS for NGINX is pretty simple with homebrew. The following <a href="https://dev.to/arjavdave/installing-nginx-on-mac-46ac">article</a> walks through the steps of the installation:</p>
<pre><code class="language-bash">brew update
brew install nginx
nginx
</code></pre>
<p>If you do not want to install it from homebrew, this <a href="https://gist.github.com/beatfactor/a093e872824f770a2a0174345cacf171">gist</a> can help install it from the source.</p>
<h3>Windows</h3>
<p>For Windows installation, you can follow the <a href="https://nginx.org/en/docs/windows.html">guide</a> from the official documentation.</p>
<pre><code class="language-bash"># INSTALL the https://nginx.org/en/download.html
# A Zip file with the name nginx-version.zip will be downlaoded
# COPY it to the desired location and use that path while unzipping
cd c:\
unzip nginx-1.25.3.zip
cd nginx-1.25.3
start nginx
</code></pre>
<p>You can check the status of NGINX if the installation was successful or not with the command:</p>
<pre><code class="language-bash">tasklist /fi &quot;imagename eq nginx.exe&quot;
</code></pre>
<p>This should be from the installation section.</p>
<h2>Understanding the default config</h2>
<p>When you have completed the installation of nginx, you can see the default nginx configuration in the file path as <code>/etc/nginx/nginx.conf</code> in Linux/macOS or `C:
ginx</p>
