---
title: Test after resolving merge conflicts
type: thoughts
post_dir: thoughts
date: 2025-10-17
status: published
---

Its a beautiful sunny Monday, you had prepared changes well tested on Friday. But today, just for today, Bill had to merge their hotfix into your changes as well. Well, they did.

And I didn't bother about it. It had some merge conflicts which Bill resolves happily (I thought). Everything is green, but some things are red, due to the size of PR I thought, (+3079 -789) changes. Phew!

Merged and deployed.

10 Minutes later, you see all sorts of red shades representing the errors.

> I will advice one thing, test after you resolve and commit the merge conflicts. PLEASE!

Bill missed a comma at the parameter to a function while merging and broke in production. I know that was a human error, but adding tests won't break a sweat to already added cost of AI slop of PR summary.

Am I mad about it, no. I myself had broken prod multiple times a week. But reasons were far from testable (yes that is my excuse). But that was quite obvious. It could have gone unnoticed and degraded the quality of the software. Tests are annoying but sometimes necessary to protect and defend a developer (yes you).
 
