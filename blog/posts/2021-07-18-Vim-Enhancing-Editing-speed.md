---
templateKey: blog-post
title: "Vim: Enhancing Editing Speed"
subtitle: "Improving the way you edit code in Vim."
date: 2021-07-18 16:52:17 +0530
status: published
tags: ['vim',] 
slug: vim-editing-speed
image_url: https://cdn.hashnode.com/res/hashnode/image/upload/v1625651851743/RK-CxEtLT.png?w=1600&h=840&fit=crop&crop=entropy&auto=compress
---


## Introduction
Editing/ Writing is a crucial part of any text editor, the swiftness with which one can make changes in a file(s) or a structure is a bit dependent on the editor we use as well. Vim can be a bit hassle for beginners in the initial stage but it becomes second nature soon. It also depends majorly on the level of usage and the area of usage. If you are an advanced programmer, you will use these commands extensively, and might not be a big deal for you, But for a beginner, it might take some time to get used to the workflow in Vim.

This article won't be covering the basics of editing in Vim but the ways to save time on the basic level commands to improve efficiency. That being said, beginners can also read to be aware of the possibilities and tricks in Vim. 

Below are some quick basic commands for getting started in editing and improving the overall workflow. 

### Entering in into the Insert Mode

`i`  -> Enter into **Insert mode** from the cursor.

`I`   -> Enter into **Insert mode** from the beginning of the current line.

`a`   -> Enter into **Insert mode**  just after the cursor.

`A`   -> Enter into **Insert mode** from the end of the current line.

`o`   -> Enter into **Insert mode** below the current line.

`O`   -> Enter the **Insert mode** above the current line.

The above commands are purely to add text in the line or above or below the line without deleting anything.

The following commands will delete some characters/words/lines and then entering into the Insert mode.

`s`   -> delete the character under the cursor and enter into **Insert mode**.

`cw`  -> Delete the word under cursor and enter into **Insert mode**.

`S`   -> delete the entire line under the cursor and enter into **Insert mode**.


## Cut Copy Paste Commands

This set of commands are quite helpful as a programmer and it is used quite frequently. These commands can surely boost the time to cut-copy-paste but also they will provide more customization to the way you do these tasks.

### Cut/Delete Commands

The below-mentioned commands can be used in Normal or any Visual/ Selection mode as it depends whether you have selected the text or you want to work in Normal mode itself. 

`dd`  -> Delete the current entire line.

`dw`  -> Delete the word on the cursor.

`d0` -> Deletes the line from the current cursor position to the beginning of the line.

`D` or `d$` -> Deletes the line from the current cursor position to the end of the line.

`d`  -> Delete the selected text (only in Visual/Visual-Line/Visual-Block/Selection mode).

`x`  -> Delete the character under the cursor.

### Yank/Copy Commands

The following commands are used in the Normal mode as they perform the copying of text with words and lines only.

`yw` -> yank(copy) the word on the cursor.

`y$`  -> yank(copy) line till the end from the current cursor position.

`yy`  -> yank(copy) the current entire line to the unnamed register ("").

You have to be in Visual/ Visual line/ Selection mode to yank the text for the next set of yanking commands. 

`y`  -> yank(copy) the selected text to the unnamed register ("").

`"+y` -> yank(copy) the selected text to the system clipboard ("+ register).

### Paste Commands

`p`   -> Paste the content of the unnamed register("") below the cursor.

`P`   -> Paste the content of the unnamed register("") above the cursor.

`"+p`  -> Paster the content of system clipboard ("+ register) to the cursor.


### Replacing Text

Replacing is a great option for instant productivity, if you want to make some minor changes, you don't have to go into insert mode and delete and then edit the text. Instead, the replace commands such as `r` and `R` allow us to replace some characters being in Normal and Replace mode respectively. This can be used very heavily if you just want to replace it instead of adding/removing text.

`r`   -> replace the character under the cursor with the following key entered from the keyboard.

`R`   -> Enter into **Replace mode**( replace the character with the specified word from the keyboard). 

### Undoing and Redoing

We often make mistakes and want to revert to the changes we have to make and start from the last save. The following sets of commands will make us do exactly that.

`u`   -> Undo the last made changes before saving.

`U` -> Restore the changes in the entire line.

`<C-R>`   -> Redo the last undo (un-undo -> revert back changes).


### Search and Replacement

Some of the below-mentioned commands are a great set for bulk replacement and addition. We can either replace a particular pattern in the entire file or in specific parts of the file as mentioned and explained as follows:

`:%s/old/new`  -> Replace the word `old` with `new` in the entire file.

`:4s/old/new`  -> Replace the word `old` with the word `new` on line 4( where 4 can be any number of lines in the file).

`:2, 9s/old/new`  -> Replace the word `old` with `new` between the lines 2 and 9 inclusive(where 2 and 9 can be any number of lines in the file).

`:%s/^/This`  -> Add `This` to the beginning of each line in the file.

`:%s/$/That`  -> Append the word `That` to the end of each file.

You can notice that the `%` symbol here indicates the entire file. We can skip prefixing s with `%` to make changes only in the current line or any number of lines specified instead of it. This command is quite a lot customizable and powerful, I can't show each and every combination of this command. It should be used as per requirement and thus should be modified accordingly.


### Indenting Text

Indenting is quite important in certain languages like Python, YAML, Haskell, etc. This can get really frustrating if you even miss a single indentation, you have to format everything to a proper indentation scratch. But thanks to powerful Text-editors and IDEs which have made the indentation quite easy and even auto-correct the wrong indentation. Vim has some commands to make those indentations much easier and it also has a customizable number of spaces in its config file called vimrc.

`>>`  -> Indent or shift the current line to the right. (normal mode)

`<<`  -> Unindent shift the current line to the left. (normal mode)

`>`   -> Indent or shift the selected text to right. (Visual/Visual-line/VIsual-block/Select mode)

`<`  -> Unindent or shift the selected text to left. (Visual/Visual-line/VIsual-block/Select mode)
 
If your file is saved as a particular language that supports indentation, it will automatically indent lines for you but it cannot be reliable. So, we need to [IndentLine](https://github.com/Yggdroot/indentLine), and others as well. 

## Miscellaneous

`~` -> Convert the character under the cursor to upper case/ lower case.

`vip` -> Yank a entire paragraph ( till empty line).

`gu` -> Convert the selected text into lowercase. (Visual/Select Mode)

`gU` -> Convert the selected text into Uppercase. (Visual/Select Mode)



All of the above commands were somewhat basic and commonly used but if used along with other key shortcuts for movement can also improve the editing speed quite considerably. Just keep using these commands and you'll be amazed by the speed you've developed. Thanks for reading. Happy Coding :)

References:  [catswhocode](https://catswhocode.com/vim-commands/)  [The valuable dev](https://thevaluable.dev/vim-advanced/) [rtorr.com](https://vim.rtorr.com/)
                      
