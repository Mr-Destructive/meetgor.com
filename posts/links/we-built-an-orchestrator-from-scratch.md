---
date: 2024-11-30T00:00:00Z
description: We built an orchestrator from scratch
image_url: https://i.ytimg.com/vi/dy2RJdDEvO0/hqdefault.jpg
link: https://www.youtube.com/watch?v=dy2RJdDEvO0
newsletter: 2024-11-30-techstructive-weekly-18
slug: we-built-an-orchestrator-from-scratch
source: newsletter
status: published
tags:
    - links
title: We built an orchestrator from scratch
type: links
---


## Commentary

- : Fly.ioWhy Flyio built their orchestrator, kind of. They use VMs which Kubernetes is not ideal for as it is designed for orchestrating containers, Nomad has a quirk of assigning a minimal number of VMs/machines for the users, which is not secure enough for Fly.io, so reasonable enough that it is worth for them to write and Orchestrator from scratch (well not entirely from scratch)
