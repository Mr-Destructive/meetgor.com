---
templateKey: blog-post
title: "Vim: Terminal Integration"
subtitle: "To feel and live in Vim with the terminal."
date: 2021-06-29 15:21:11 +0530
status: published
tags: ['vim', 'linux',]
slug: vim-plus-teminal
image_url: https://meetgor-cdn.pages.dev/vim-terminal-integration.webp
---

## Vim and Terminal!?
Vim was made to work with the command line. Many beginners do not understand what are the true capabilities of Vim, myself included:) Vim can run terminal commands without leaving the text editor, open an instance of a terminal, work with shell environments, and other things depending on the use case.

## Running Terminal/ shell commands from within Vim

You can run the commands from inside of Vim by just using `:!` before the command, this means you have to be in command mode. Just after being in command mode, the ! or bang operator will execute the command typed after it from the terminal(Linux/ macOS) or your default shell(Windows -> CMD/Powershell).
```
:!pwd
```
The above command from vim will redirect to the terminal and show the output of the command and return on pressing any key. In this case, it will execute the PWD command and just wait for the user to enter any key to return to Vim.

The following is an example of how it could be used from Vim in Windows using Powershell as the default shell.

![Animation.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1624885870237/Ie5C-3u1B.gif)

In Windows, dir is equivalent of ls for Linux. That was not the best example of how a terminal can be used at its best, You can also use a logical operator from within vim to run multiple commands at once. 

### Running programs/ code from Vim on terminal

This becomes quite a great feature for making Vim from a text editor to an IDE, this can be paired with Keymaps i.e when the user types certain keys, the command gets executed making the code run from the terminal. I have already used this feature to set up Vim for python, bash, and other programming languages. Also, I have written an article about  [keymapping](https://dev.to/mrdestructive/vim-keymapping-guide-3olb)  and Vim setup for  [Python](https://dev.to/mrdestructive/setting-up-vim-for-python-ej)  and  [Bash](https://techstructiveblog.hashnode.dev/vim-setup-for-bash-scripting), this will give you an idea of how to setup vim for any programming language. 

Vim can really shine in this kind of feature as it just becomes flawless and a smooth experience even for a beginner. We just have to run the compile the code and run its executable/ output file, rather for python and other interpreted languages, we have to just pass the file name to the interpreter and that's it.  

## Opening instance of Terminal within Vim.

Vim can also create an instance of the terminal within its window by making a split. This is quite similar to VS Code and other Text editors that have the functionality to create an instance of the terminal within itself. This feature is useful for developing complex systems and depending on the use case, it can be quite important and efficient as well. 

The terminal can be created in various ways the most preferred way is by typing in `:term` from Vim. 
This will create a horizontal split from the current editor and split it into half. You can change the size of the split using the mouse according to your preference. 

![vimtermsplit.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1624888468392/wR0JT8SBN.gif)

Here Vim has certain variables and shortcuts to make things even simpler, say you want to parse the current file to the terminal for execution. You can surely type the name manually or you can be a bit smarter and use % instead, the `%` symbol will parse the file name along with the extension in the terminal. Also `%:r` will parse filename without the extensions(.txt/.py/etc) to the terminal.

There are many things you can do with terminals surely, but with Vim that even goes further than the limits. Terminal/command line is quite important in any development environment as it is an interface for the user to interact with the Operating System. Vim is quite powerful and behaves as a gecko for programmers because it changes itself according to our needs flawlessly and **efficiently**.


![vimpython.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1624891655340/5f81Dpp_O.gif)

Integrating Terminal into a Text Editor truly lights up the environment for development. It becomes an easy and enjoyable experience to test out the code without wasting much time on the actual execution process. Surely it needs time to set up the environment to speed things, for that understanding of the programming and development environment is required. Happy Viming :)
