{"author":"meet","date":"2025-04-11","post_dir":"til","published":"published","tags":["git"],"title":"Add hunks: only line specific changes with git add -p","type":"til"}

I had made the changes for a future release, next set of changes so to speak in git terms. Then, I encountered a bug while testing the previous changes. Now, I could have 
1. Open VS Code and add that little `+` icon in the gutter to fix the bug
2. Create a separate branch with stashed changes for this new feature
3. Use GitSigns in NeoVim
4. Use `git add -p` to stage [hunks](https://git-scm.com/book/en/v2/Git-Tools-Interactive-Staging)

The last one was a gem given by ChatGPT, and then I went a bit rabbit hole for reading about it and it turns out it's real! 

It's called staging patches, and it has an interactive mode to select or reject each hunk (patch, line of changes) at a time. So by entering the command `git add -p` this will let you go through each file and individual change (patch/hunk) and take action interactively. You can use the below commands for moving one hunk or patch at a time: 
- `y` for adding (staging) the current hunk (patch)
- `n` for not staging the current hunk

Or you can even use the below commands for moving one file at once:
- `a` for adding (staging) all the patches in the file (just like `git add file.txt`)
- `d` for not staging the file at all, just skip the file for patching

If there are multiple hunks (patches) in the file, the prompt will show the other options like `[y,n,q,a,d,j,J,g,/,e,?]` 

```
(1/7) Stage this hunk [y,n,q,a,d,j,J,g,/,e,?]? g
1: -33,6 +33,7 +from app.utils.utils import get_additional_param
2: -345,6 +346,7 + strict_check = str_to_bool(get_additional_param("stri
....
....
7: -875,7 +878,7 - if strict_check:
```

This shows the `g`, `j`, `J`, `/`, and `e` which are quite helpful for navigating across multiple patches (hunks) in the file. I haven't explored all of them, but the 4, I showed earlier, are the ones that I think I will be using extensively now, within the terminal. No need of fancy editor shenanigans and redundant branching tricks. 

Git Good!
