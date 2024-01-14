---
templateKey: blog-post
title: "NGINX Basics and Setup"
description: "Exploring NGINX Fundamentals: A Guide for Backend Developers, from the Importance of Learning NGINX to Installation and Server Setup"
date: 2024-01-14 18:15:00
status: published
slug: nginx-01-basics
tags: ['nginx', 'web-development']
series: ['nginx-survival-guide',]
series_description: "NGINX Survival Guide: A Guide for Backend Developers, understanding the bare minimum fundamentals to get going with NGINX."
image_url: https://meetgor-cdn.pages.dev/nginx-survival-guide/nginx-sg-1.png
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
