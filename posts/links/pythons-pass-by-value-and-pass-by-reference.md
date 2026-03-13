---


























type: links
title: 'Python’s pass by value and pass by reference'
date: 2025-08-09
slug: pythons-pass-by-value-and-pass-by-reference
tags:
  - links
link: 'https://www.thepythoncodingstack.com/p/python-pass-by-value-reference-assignment'
status: published
description: 'Python’s pass by value and pass by reference'
source: newsletter
newsletter: 2025-08-09-techstructive-weekly-54


























---


## Commentary

- This is one hell of a reason, Python gets a little more confusing and less friendly.
- TLDR of the post is that if you pass a immutable variable/object to a function call in python, you need to return it back from the function (if the function modifies those immutable objects). Because the object is immutable it won’t get updated inside the function, it will be created a new, so we need to assign it to the modified version when the function returns.
- But for mutable objects, the function can modify it and we are passing it to the function, so the object will be updated.
