---
title: "Gitignore isn’t the only place to ignore files in git"
date: 2026-06-19
link: "https://nelson.cloud/.gitignore-isnt-the-only-way-to-ignore-files-in-git/"
status: published
image_url: "https://nelson.cloud/.gitignore-isnt-the-only-way-to-ignore-files-in-git/og.png"
source: newsletter
newsletter: techstructive-weekly-99
type: links
slug: gitignore-isn-t-the-only-place-to-ignore-files-in-git
tags:
description: "You can ignore files in .gitignore, .git/info/exclude, and ~/.config/git/ignore"
hash: 1770073c37a1d553f3b0cf655d0f1e8bdb41d5418e308266f487ecb3e6ba588b
---
My thoughts on [Gitignore isn’t the only place to ignore files in git](https://nelson.cloud/.gitignore-isnt-the-only-way-to-ignore-files-in-git/): Gitignore isn’t the only place to ignore files in git

## Commentary

- Gitignore isn’t the only place to ignore files in git
- Wow! This was clever. I didn’t knew you could ignore a previously tracked file from the git without mentioning it in the gitignore.
- There are two more places you could do that, those look internal but gets the job done without leaking on the public
- .git/info/exclude: This is for the local git repo, you add the filename and it ignores the file. Pretty neat.
- ~/.config/git/ignore: This is for global context. If you say .DS_STORE or whatever mac has, it will not track in any git repo of your device.
