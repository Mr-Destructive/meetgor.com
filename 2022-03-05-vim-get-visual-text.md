---
cover: ''
date: 2022-03-05
datetime: 2022-03-05 00:00:00+00:00
description: Store the selected text in a variable in Vim Script
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/tils/2022-03-05-vim-get-visual-text.md
long_description: We can get the selected text in a variable in Vim Script using registers.
  Lets break down the command Here, we are entering normal mode and selecting text
  which was previously selected and yank the contents into a register in this case
  we are using (
now: 2022-07-24 13:37:38.457948
path: blog/tils/2022-03-05-vim-get-visual-text.md
prevnext: null
slug: vim-get-visual-text
status: published-til
tags:
- vim
templateKey: til
title: 'Vim: Get the Text from Visual Selection'
today: 2022-07-24
---

## Using Registers 

We can get the selected text in a variable in Vim Script using registers. 

```vimscript
normal gv"xy
let context = getreg("x")
```

Lets break down the command

```
normal mode -> gv -> (y)ank text -> to the "x" register
               |                  |
               |               Copy the contents into x register(or any register you like)    
               |                  
            Select the previously selected text   
```

Here, we are entering normal mode and selecting text which was previously selected and yank the contents into a register in this case we are using (x) register, it can be any register. 
Now to get the contents of that register we can use the function `getreg("register_name")` or use `"xp"` to paste the contents of the `x` register or more generally for any register(`"<register-name>p`). 

Hence we can store the contents of the selected text in a variable for further processing or manipulation.

To quickly test this snippet from command mode, you can try the following steps:

Select a text and press Escape, we just want the `gv` command to refresh and get it's contents to the latest visual selection.

```vimscript
:normal! gv"xy
```

```vimscript
:let foo = getreg("x")
```

```vimscript
:echo foo
```

The echo command will simply print the text which we have selected in the file. 

![Visual Select Text](https://res.cloudinary.com/techstructive-blog/image/upload/v1646483173/blog-media/wlrxgtmegtycilyhvyiz.gif)