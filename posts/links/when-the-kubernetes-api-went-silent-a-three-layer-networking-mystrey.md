---
title: "When the Kubernetes API went silent: A three layer networking mystrey"
date: 2026-11-13
link: "https://cloud.google.com/blog/products/networking/the-three-layer-gke-network-mystery"
status: published
image_url: "https://storage.googleapis.com/gweb-cloudblog-publish/images/Networking.max-2500x2500.jpg"
source: newsletter
newsletter: techstructive-weekly-119
type: links
slug: when-the-kubernetes-api-went-silent-a-three-layer-networking-mystrey
tags:
description: "A tricky Google Cloud support ticket that spanned GKE, Private Service Connect (PSC), Networking, Hostnet, Kernel, and the Linux kernel itself."
hash: db77d153e8cfebd59d1200c772d53515dd83dc83696594d0009b608f7836b506
---
My thoughts on [When the Kubernetes API went silent: A three layer networking mystrey](https://cloud.google.com/blog/products/networking/the-three-layer-gke-network-mystery): When the Kubernetes API went silent: A three layer networking mystrey

## Commentary

- When the Kubernetes API went silent: A three layer networking mystrey
- This was quite a complex thing to debug and understand. Like combination of packets arriving out of order, Linux TCP bug, and a networking component dropping packets needed to recover the SYN-ACK.
- I discovered that I didn’t knew a lot of fundamentals of the networking stack. Some things I learnt were like NICs, TCP and Linux sockets to Kubernetes, proxies and Google Cloud's VPC, etc.
- The issue then I can understand was something like that made connections stuck in an unexpected state, and when a new connection tried to reuse the same network port, Google's networking layer dropped a packet needed to recover. So, like its not about single failure that you can point out, rather its a nested things to uncover from those. Like from a Pod and its Kubernetes node, through virtual NICs, VPC networking, load balancers and proxies, and control plane.
