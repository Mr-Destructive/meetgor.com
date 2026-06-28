---
title: "Parquet is more than Turbo CSV"
date: 2026-06-26
link: "https://csvbase.com/blog/3"
image_url: "https://csvbase.com/blog-static/tjentiste.jpg"
source: newsletter
newsletter: techstructive-weekly-100
slug: parquet-is-more-than-turbo-csv
tags: 
status: published
type: links
description: "Quicker, but also more convenient"
hash: d83d02bc085b5a64cb682346b1ce13c91a363ee953164c335cc04ffdc9f50200
---
My thoughts on [Parquet is more than Turbo CSV](https://csvbase.com/blog/3): Parquet is more than Turbo CSV

## Commentary

- CSVBase Blog: Parquet is more than Turbo CSV
- This was so cool. Last week I. read about the duckdb and its engine. I was intrigued by the design of duckdb like parquet. And this week this answers that question. Seems like I am sync with the universe, its giving me hints.
- So, parquet is column based storage but also row based chunking. So streaming is not possible if you need intermediate data. Since its broken down into columns, its not possible to stream half of it, either you get full or nothing. The metadata thing is so cool. it has sorted the different types of data, encodings, null and even time stamps and zones (god has anyone said time? what time). having rich metadata is not a tradeoff its a design that makes parquet really great for large data storage.
- It reduces the size to almost 50% from csv that is just bonkers. Imagine what it would be from json to parquet that I did and felt for aircraft flight observatory project a couple of months back. Its quite evident, parquet is a great format for storage.