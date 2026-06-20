---
title: "Marek Suppa’s BLog: Making HTTP requests from a container that has no curl, using bash /dev/tcp"
date: 2026-06-19
link: "https://mareksuppa.com/til/bash-dev-tcp-http-without-curl/"
status: published
image_url: "https://mareksuppa.com/og/bash-dev-tcp-http-without-curl.png"
source: newsletter
newsletter: techstructive-weekly-99
type: links
slug: marek-suppa-s-blog-making-http-requests-from-a-container-that-has-no-curl-using-bash-dev-tcp
tags:
description: "Minimal container images often ship without curl, wget, or any HTTP client at all. Bash can open a TCP socket through /dev/tcp, which is enough to write a tiny HTTP/1.1 request by hand for quick checks."
hash: 0c4215e07e537a903b35bb606c5de76cf0d2f0da083f64ff5dc148d9b29ea8b7
---
My thoughts on [Marek Suppa’s BLog: Making HTTP requests from a container that has no curl, using bash /dev/tcp](https://mareksuppa.com/til/bash-dev-tcp-http-without-curl/): Marek Suppa’s BLog: Making HTTP requests from a container that has no curl, using bash /dev/tcp

## Commentary

- Marek Suppa’s BLog: Making HTTP requests from a container that has no curl, using bash /dev/tcp
- This is cool. Bash is some magical compiler that knows a little too much to be used right?
- It cannot read the request just parse it up so that it can send the request. It feels like everything was designed so perfectly right? Imagine an AI in this container, will it be able to get the request across?
