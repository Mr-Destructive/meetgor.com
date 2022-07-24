---
cover: ''
date: 2021-06-26
datetime: 2021-06-26 00:00:00+00:00
description: '![ OK First things first, unmap arrow keys and make a habit of using
  h,j,k, and l. Yes, this would not make any sense in the initial stage but that will
  make no'
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-06-26-Vim-Enhancing-Movement-speed.md
image_url: https://techstructiveblog.hashnode.dev/_next/image?url=https%3A%2F%2Fcdn.hashnode.com%2Fres%2Fhashnode%2Fimage%2Fupload%2Fv1624692961710%2FGNj64mCHz.png%3Fw%3D1600%26h%3D840%26fit%3Dcrop%26crop%3Dentropy%26auto%3Dcompress%2Cformat%26format%3Dwebp&w=1920&q=75
long_description: '![ OK First things first, unmap arrow keys and make a habit of
  using h,j,k, and l. Yes, this would not make any sense in the initial stage but
  that will make no sense for not using it later. This is quite the common movement
  that every programmer use'
now: 2022-07-24 13:37:38.454890
path: blog/posts/2021-06-26-Vim-Enhancing-Movement-speed.md
slug: vim-movement-speed
status: published
subtitle: Learning to Fly and Glid in Vim!
tags:
- vim
templateKey: blog-post
title: 'Vim: Enhancing Movement Speed'
today: 2022-07-24
---

![]({{ page.image | relative_url }})
## Introduction
OK! Vim and movement are like bread and butter or failed brakes. To become a proficient Vim user, you need to move in Vim very effectively. You don't have to think about doing certain things, your fingertips should automatically move without wasting time thinking about it. I am sure, it takes time and effort but OH! it is so rewarding in the end.  

### Why H J K L?
First things first, unmap arrow keys and make a habit of using h,j,k, and l. Yes, this would not make any sense in the initial stage but that will make no sense for not using it later. 
The thing with H J K L is that they are quite flexible to work with and if you use it with numbers you would navigate much faster than arrow keys. Such as `10j` will move you 10 lines down at a time in normal mode. These keys are used with many of the other key shortcuts and commands which just make it most important to begin learning to move around in Vim with H J K L. 

## Moving Horizontally.
This is quite the common movement that every programmer uses most of the time. This is also a much-neglected part when movement in Vim is concerned. To the basics, we use the following commands: 

 `w`   ->  **move forward by a word (considering punctuations as separate words).**

 `W`   ->  **move forward by a word (punctuations ignored).**

`b`  ->  **move backward by a word (considering punctuations as separate words).**

`B`  ->  **move backward by a word (punctuations ignored).**

`e`  ->  **move to end of a word (considering punctuations as separate words).**

`E`  ->  **move to end of a word (punctuations ignored).**

`0`  ->  **move to the beginning of a sentence.**

`$`  -   **move to the end of a sentence.**

Those are the most useful and common commands for moving across the line. Don't forget to use the number before the command to repeat the task for that number of times. Like for example, if you would like to go 6 words ahead type in `6w`. This can improve your thinking and typing as well, good signs of a programmer Eh!
 

## Moving Vertically.

To move vertically we can imagine moving within a file or the block of code. For moving in a file, the following are some useful commands.

`gg`  -> **move to the beginning of a file.**

`G`  ->  **move to the end of a file.**

`Ctrl + e`  ->  **move the screen down without moving the cursor.**

`Ctrl + y`  -> **move the screen up without moving the cursor.**

`Ctrl + f`  -> **move forward one entire screen.**

`Ctrl + b`  -> **move backward one entire screen.**

`Ctrl + d`  -> **move forward half screen.**

`Ctrl + u`  -> **move backward half screen.**

This just was moving around the screen and now a bit programmatic movement. We will see some keystrokes to move in code blocks or code snippets very efficiently.

`gd`  -> **move to the local declaration of any code.**

`gD`  -> **move to the global declaration of any code.**

`%`  -> **move between pairs of ( ), { }, [ ] or any other type of such braces.**

`{`  -> **move to the next paragraph/ code block/ function/ etc)**

`}`  -> **move to the previous paragraph/ code block/ functions/ etc)**

`fa`  -> **move to the next occurrence of the character 'a' in a sentence.**

`Fa`  -> **move to the previous occurrence of the character 'a' in a sentence.**

`ta`  -> **jump to before of the next occurrence of the character 'a' in a sentence.**  

`Ta`  -> **jump to after of the previous occurrence of the character 'a' in a sentence.**

The above might be quite handy key shortcuts in moving in a large code file. Even in files with complex variable names and structures, this can be quite handy. 

## Search and navigation.
Searching is quite a time-consuming task, especially when the code is quite complex and has a lot of variables and all. Vim shines in many of such aspects where people think it's dead. It rises from the ashes to produce a performance-driven experience like any other modern IDEs though it requires a bit of research:) Here are some commands that will make searching and navigating around it quite a lot easier.

`*` -> **next occurrence of the word under the cursor.**

`#` -> **previous occurrence of the word under the cursor.**

`n`  -> **next occurrence of the word searched pattern.**

`N`  -> **previous occurrence of the word searched pattern.**


The above commands will also work if you search the pattern from the command mode.
`/pattern`  or `?pattern` Enter and navigate to the next(`*` or `n`) and previous(`#` or `N`) occurrence of that pattern match.


## Moving across files.

Moving across files without any plugins or file explorer is often considered tricky or impossible for some people but there is a way. You can switch between files using the following commands:

`Ctrl + O`  ->   **move in the previously opened file.**

and 

`Ctrl + I`  ->  **move in the next file.**

We also can use `Ctrl + ^ ` to move the previous two opened files.


If you want to switch from buffers, you can use `:bn` for moving into the next buffer, and `:bp` to move in the previous buffer. You always have an option to move from a buffer using the file name `:b filename` or using the index as `:bindex`.


## Moving between Tabs.

People rarely use Tabs as far as I have seen, but they are quite useful and provide the polish off just as robust IDEs. 

`:tabnew filename`  ->  **create a Tab of a file.**

`gt`  -> **move to the next tab.**

`ngt`  -> **move to the nth tab.**

`gT`  -> **move to the previous tab**

`:tabo`  -> **close all the tabs except the current one.**

`:tabc`  -> **close the tab.**

`:tabm n`  -> **move the current tab to nth position.**


## Movement in Marks.

Marks are used for some quite large files and code-bases. It is used to move from one mark(kind of a bookmark) to another using few key commands, marks are generally created when you would go to a particular code block or any part of the file again and again. Some of the quick navigation using maps are the following.

`mn`  ->  **set the current position as mark 'n'.**

`<backtick>n`  ->  **jump to the position of mark 'n'.**

` `0`  -> **jump to the position where vim was last exited.**

` `"`  ->  **jump to the position when the last edit was made in the file.**


## Split Windows Movement

`Ctrl + w + r`  -> **move the split down.**

`Ctrl + w + R`  -> **move the split up.**

`Ctrl + w + h`  ->  **jump to the left split.**

`Ctrl + w + j`  ->  **jump to the split down.**

`Ctrl + w + k`  -> **jump to the upper split.**

`Ctrl + w + l`  ->  **jump to the left split.**

You can use Caps H J K L to move the leftmost, bottom, uppermost, rightmost split respectively. 
That just was quick to make you enough faster than previous hassles.


That was probably it, these were some tricks and shortcuts to move around Vim pretty effectively and smoothly. Moving around Vim can be quite complicated at once, but it's just finding the key shortcuts to make you feel and glide in VIm. There might be quite a lot of shortcuts missing, If you have any quicker shortcuts, Please let me know in the comments.  Happy Viming :)
<div class='prevnext'>
    <style type='text/css'>
    :root {
      --prevnext-color-text: #eefbfe;
      --prevnext-color-angle: #ff6600;
      --prevnext-subtitle-brightness: 3;
    }
    [data-theme="light"] {
      --prevnext-color-text: #1f2022;
      --prevnext-color-angle: #ffeb00;
      --prevnext-subtitle-brightness: 3;
    }
    [data-theme="dark"] {
      --prevnext-color-text: #eefbfe;
      --prevnext-color-angle: #ff6600;
      --prevnext-subtitle-brightness: 3;
    }
    .prevnext {
      display: flex;
      flex-direction: row;
      justify-content: space-around;
      align-items: flex-start;
    }
    .prevnext a {
      display: flex;
      align-items: center;
      width: 100%;
      text-decoration: none;
    }
    a.next {
      justify-content: flex-end;
    }
    .prevnext a:hover {
      background: #00000006;
    }
    .prevnext-subtitle {
      color: var(--prevnext-color-text);
      filter: brightness(var(--prevnext-subtitle-brightness));
      font-size: .8rem;
    }
    .prevnext-title {
      color: var(--prevnext-color-text);
      font-size: 1rem;
    }
    .prevnext-text {
      max-width: 30vw;
    }
    </style>
    
    <a class='prev' href='/pipenv-intro'>
    
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M13.5 8.25L9.75 12L13.5 15.75" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"> </path>
        </svg>
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>prev</p>
            <p class='prevnext-title'>Pipenv: Python's Official Package Management tool</p>
        </div>
    </a>
    
    <a class='next' href='/golang-structs'>
    
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>next</p>
            <p class='prevnext-title'>Golang: Structs</p>
        </div>
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.5 15.75L14.25 12L10.5 8.25" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path>
        </svg>
    </a>
  </div>