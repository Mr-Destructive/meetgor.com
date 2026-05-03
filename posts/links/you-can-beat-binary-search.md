---
tags: 
description: "We sometimes have to look for a value in a sorted array. The simplest algorithm consists in just going through the values one by one, until we encounter the value, or exhaust the array. We sometimes call this algorithm a linear search. In C++, you can get the desired effect with the std::find function. For … Continue reading You can beat the binary search"
hash: e2cf4c0c3b1d3ffad9a95a700cf14bf4f0bae9352563a22b4cb11f8aebb60eeb
title: "You can beat binary search"
date: 2026-05-01
link: "https://lemire.me/blog/2026/04/27/you-can-beat-the-binary-search/"
status: published
source: newsletter
newsletter: techstructive-weekly-92
type: links
slug: you-can-beat-binary-search
image_url: "https://lemire.me/blog/wp-content/uploads/2026/04/Gemini_Generated_Image_rsv3zvrsv3zvrsv3-1024x572.png"
---
My thoughts on [You can beat binary search](https://lemire.me/blog/2026/04/27/you-can-beat-the-binary-search/): You can beat binary search

## Commentary

- You can beat binary search: Quad Search
- This is creative. Instead of just halving the list, we do a buckets of 16, quads, and compute which of the buckets the element can be. Since we have the max of each bucket. We can find the bucket and parallelly then compare each element in it.
- Surprisingly, this is better for large arrays, and with multiple cores this will outperform binary search easily. Because we are not diving the array into halves anymore, we are picking the most possible region where the element could be and parallel searching it in those.