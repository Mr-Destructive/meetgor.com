---
date: 2025-06-14
link: 'https://medium.com/javarevisited/how-i-finally-understood-docker-and-kubernetes-5debb13cacfe'
status: published
source: newsletter
newsletter: 2025-06-14-techstructive-weekly-46
hash: 9e6692727719bcc85eb45dd016700b6f2222af5530e78113f5c04313484d0bd1
title: 'How I finally understood Docker and Kubernetes'
slug: how-i-finally-understood-docker-and-kubernetes
tags: 
description: 'How I finally understood Docker and Kubernetes'
type: links
---
## Commentary

- This was my pick of the week. I understood the reason why Kubernetes exists
- I knew the concept of Docker (it could be because I have used it extensively in the past to create projects as well in my internships to deploy APIs and apps)
- But the concept of Kubernetes is like a black-box. But the author’s explanation style and simple example made it clear.
- Kubernetes is like
- Container Image > Deployment > Pod > Service
- Container Image is the actual image of your app that you want to run, maybe it has multiple of those.
- Deployment is like defining what and how many (other things too) to run.
- Pod is like the actual unit of containers; in itself, it has no control, it just runs whatever was given to it.
- Service is like the layer that exposes it to the world, maybe the network, the other containers, which are like a configurable exposure of the network.