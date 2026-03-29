---
title: "Ctrl + C in psql gives me the heebi-jeebies"
date: 2026-03-27
link: "https://neon.com/blog/ctrl-c-in-psql-gives-me-the-heebie-jeebies"
status: published
image_url: "https://neondatabase.wpenginepowered.com/wp-content/uploads/2026/03/neon-ctrl-c.jpg"
source: newsletter
newsletter: techstructive-weekly-87
type: links
slug: ctrl-c-in-psql-gives-me-the-heebi-jeebies
tags:
description: "There are a few different reasons to hit the brakes on a Postgres query. Maybe it’s taking too long to finish. Maybe you realised you forgot to create an index that will make it orders of magnitude quicker. Maybe there’s some reason the results are no longer needed. Or maybe you, or your LLM buddy, […]"
hash: 463a6babd95663433d821e8b2f76608489a383539c217b216483ba9fd85fa773
---
My thoughts on [Ctrl + C in psql gives me the heebi-jeebies](https://neon.com/blog/ctrl-c-in-psql-gives-me-the-heebie-jeebies): Ctrl + C in psql gives me the heebi-jeebies

## Commentary

- Ctrl + C in psql gives me the heebi-jeebies
- This is really interesting. I like the way he calls it heebi-jeebies. It really is.
- Like the TLS is not there for the cancel request, so your psql connection sends the unencrypted database secret in the wild, and somehow if intercepted by anyone in the same network, it can launch a Denial of Service attack.
- The Neon Proxy and Elephant shark(the wireshark but for Postgres) have a workaround by noting the secret with the initial connection and when the psql sends it with the plain text the secret it intercepts it and kills the right session. Wired stuff but kind of no choice, that would require a bit of a refactor on the protocol.
