---
title: "Parquet is more than Turbo CSV"
date: 2026-06-26
link: "https://csvbase.com/blog/3"
status: published
image_url: "https://csvbase.com/blog-static/tjentiste.jpg"
source: newsletter
newsletter: techstructive-weekly-100
type: links
slug: parquet-is-more-than-turbo-csv
tags:
description: "Quicker, but also more convenient"
hash: 6261834358462991a2347dc40626cc5d3a3e3da5aeaa53314cd83c7e876390c2
---
My thoughts on [Parquet is more than Turbo CSV](https://csvbase.com/blog/3): Parquet is more than Turbo CSV

## Commentary

- CSVBase Blog: Parquet is more than Turbo CSV
- This was so cool. Last week I. read about the duckdb and its engine. I was intrigued by the design of duckdb like parquet. And this week this answers that question. Seems like I am sync with the universe, its giving me hints.
- So, parquet is column based storage but also row based chunking. So streaming is not possible if you need intermediate data. Since its broken down into columns, its not possible to stream half of it, either you get full or nothing. The metadata thing is so cool. it has sorted the different types of data, encodings, null and even time stamps and zones (god has anyone said time? what time). having rich metadata is not a tradeoff its a design that makes parquet really great for large data storage.
- It reduces the size to almost 50% from csv that is just bonkers. Imagine what it would be from json to parquet that I did and felt for aircraft flight observatory project a couple of months back. Its quite evident, parquet is a great format for storage.
