---
newsletter: 2025-01-18-techstructive-weekly-25
date: 2025-01-18
link: 'https://victoriametrics.com/blog/go-http2/?ref=dailydev'
image_url: 'https://victoriametrics.com/blog/go-http2/go-http2-preview.webp'
source: newsletter
hash: 341385e3d7f07480c5c9467fe392114fcb6688b0289e1d94e14a23fff5d2f854
type: links
title: 'How HTTP/2 Works and How to Enable It in Go'
slug: how-http2-works-and-how-to-enable-it-in-go
tags: 
status: published
description: 'How HTTP/2 Works and How to Enable It in Go'
---
## Commentary

- This is also another post that I took the time to read and was worth it. I honestly don’t know how HTTP 2 works. To some extent, I know how HTTP 1 works, but if someone went a bit deeper, I would start breaking sweat. I really need to implement HTTP from scratch to understand the network stack—one day or day one.
- OH, the article, yes it talked in detail about what is the problem with HTTP 1 and how HTTP somewhat solves it.
- It is about breaking down the data into frames and makes sure the client has received the frame even if the previous frame is delayed. The fastest frame is served so, it doesn’t block the latest requests.