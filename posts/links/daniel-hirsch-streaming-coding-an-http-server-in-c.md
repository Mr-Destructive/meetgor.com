---
link: "https://youtu.be/WmZi1yrOD0Q"
status: published
image_url: "https://i.ytimg.com/vi/WmZi1yrOD0Q/hqdefault.jpg"
type: links
hash: 801cd5d8f5d0dfacfb1eccf2818cb1444c68904fa3b5782b7329a484670206c0
title: "Daniel Hirsch Streaming: Coding an HTTP Server in C"
date: 2026-06-26
source: newsletter
newsletter: techstructive-weekly-100
slug: daniel-hirsch-streaming-coding-an-http-server-in-c
tags: 
description: "Daniel Hirsch Streaming: Coding an HTTP Server in C"
---
My thoughts on [Daniel Hirsch Streaming: Coding an HTTP Server in C](https://youtu.be/WmZi1yrOD0Q): Daniel Hirsch Streaming: Coding an HTTP Server in C

## Commentary

- Daniel Hirsch Streaming: Coding an HTTP Server in C
- This was a great learning experience. I wanted to understand what a socket connection really does in a http server. So its basically a file descriptor just like pipe, socket, or terminal descriptior in unix. That is just like a file streaming. That makes sense. The interface is unified for accessing the network and the file.
- I feel like building this in Golang. Maybe this weekend it will be a stream to build a http server from scratch in Golang. I want to see how we connect to sockets and listen and then accpet and then process the request and response to the client. It feels great to just be able to do things low level.