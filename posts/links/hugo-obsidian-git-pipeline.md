---
date: 2025-08-16
slug: hugo-obsidian-git-pipeline
description: 'Hugo + Obsidian + Git Pipeline'
newsletter: 2025-08-16-techstructive-weekly-55
type: links
tags: 
link: 'https://hugo.insanelogs.xyz/posts/hugo-pipeline/'
status: published
source: newsletter
hash: 49858537cf465e19669fcec4cdac8491616543e8732a9d2187444c9871b27a77
title: 'Hugo + Obsidian + Git Pipeline'
---
## Commentary

- This blog gave me a idea to simplify my static site generator, instead of syncing with the database from github, I can just simply use a sqlite file and sync to the database an vice versa. Right now the problem is on the inconsistency in github content vs the database, so after having an in-memory or local file that can reside at any time on the repo, it would be easier to pull, push changes to the remote repository.