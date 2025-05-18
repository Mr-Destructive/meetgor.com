{
  "type": "posts",
  "title": "NGINX Basics and Setup",
  "description": "Exploring NGINX Fundamentals: A Guide for Backend Developers, from the Importance of Learning NGINX to Installation and Server Setup",
  "date": "2024-01-14 18:15:00",
  "status": "published",
  "slug": "nginx-01-basics",
  "tags": [
    "nginx",
    "web-development"
  ],
  "series": [
    "nginx-survival-guide"
  ],
  "series_description": "NGINX Survival Guide: A Guide for Backend Developers, understanding the bare minimum fundamentals to get going with NGINX.",
  "image_url": "https://meetgor-cdn.pages.dev/nginx-survival-guide/nginx-sg-1.png"
}


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

When you have completed the installation of nginx, you can see the default nginx configuration in the file path as `/etc/nginx/nginx.conf` in Linux/macOS or `C:
ginx
