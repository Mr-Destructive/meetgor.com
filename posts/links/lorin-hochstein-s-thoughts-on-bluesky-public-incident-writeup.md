---
description: "Back on April 4, the social media site Bluesky suffered a pretty big outage. I was delighted to discover that one of their engineers, Jim Calabro, published a public writeup about it: April 2026 Ou…"
hash: 47588d795606b5de695e38b953c52b5319e33443e2f7a6c30ff487b83443c1ba
title: "Lorin Hochstein’s thoughts on BlueSky public Incident writeup"
link: "https://surfingcomplexity.blog/2026/04/12/thoughts-on-the-bluesky-public-incident-write-up/"
source: newsletter
type: links
tags: 
date: 2026-04-17
status: published
image_url: "https://surfingcomplexity.blog/wp-content/uploads/2026/04/image.png"
newsletter: techstructive-weekly-90
slug: lorin-hochstein-s-thoughts-on-bluesky-public-incident-writeup
---
My thoughts on [Lorin Hochstein’s thoughts on BlueSky public Incident writeup](https://surfingcomplexity.blog/2026/04/12/thoughts-on-the-bluesky-public-incident-write-up/): Lorin Hochstein’s thoughts on BlueSky public Incident writeup

## Commentary

- Lorin Hochstein’s thoughts on BlueSky public Incident writeup
- That is a brutal one. This is like concurrency, you thought you had one problem, but after adding concurrency, you now have two. Setting a limit on how many go-routines can be spawned in a group is critical here.
- TIME_WAIT, that is really a neat thing to learn. The TCP connection waits for that duration before sending another FIN so that the client can be sure of the delayed packet delivery, if any. And that TIME_WAIT actually caused them to fill up all the ephemeral ports. That is 28k ports, which sounds a lot, but after reading it through, that is surprisingly a low number. If you are Bluesky scale, you might need millions of ports.
- Diagnosis skill is something that is going to be super valuable going forward in the AI era. AI can help, but it would be too slow and can never reach the instinct-based debugging of humans (as of now, at least)