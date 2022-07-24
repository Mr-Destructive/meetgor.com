---
cover: ''
date: 2022-03-29
datetime: 2022-03-29 00:00:00+00:00
description: Automatically format python code in the current file after saving the
  file in Vim.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/tils/2022-03-29-Vim-Python-autoformat.md
long_description: 'If you are like me who writes Python very badly, it has empty lines
  with whitespaces, no proper format in assigning variables, not formatted according
  to  Install the  OR with pipx For a detailed guide about running packages with pipx
  head toward my '
now: 2022-07-24 13:37:38.457963
path: blog/tils/2022-03-29-Vim-Python-autoformat.md
prevnext: null
slug: vim-python-black-autoformat
status: published-til
tags:
- vim
- python
templateKey: til
title: Autoformat Python file with Black after saving in Vim
today: 2022-07-24
---

If you are like me who writes Python very badly, it has empty lines with whitespaces, no proper format in assigning variables, not formatted according to [PEP 8](https://peps.python.org/pep-0008/) standards, and you use Vim as your text editor then my friend you need a autocmd badly for it.

## Install Black in Python

Install the [black](https://pypi.org/project/black/) package in python globally or locally as per your preferences.

```
pip install black
```

OR with pipx

```
pipx install black
```

For a detailed guide about running packages with pipx head toward my article on [pipx](https://mr-destructive.github.io/techstructive-blog/pipx-intro/).

## Set up Autocmd in Vim

We can set up a autocmd. What is a autocmd? It is about running commands when certain events occur like running a command when a file is saved, a buffer is opened or closed, and so on. What we want is to run the black command from the shell when the current file is saved. 

So, we can create a autocmd as follows:

```vimscript
autocmd BufWritePost * !black %
```

Now, this will run when any type of file is saved, so we will make it specific to python by adding a `*.py` to add in the autocmd.

```vimscript
autocmd BufWritePost *.py !black %
```

This works, but it gives a prompt after the command has been executed, to run the command silently we can simply add the silent keyword before the execution of the command from the shell.

```vimscript
autocmd BufWritePost *.py silent !black %
```

This looks perfect! 

But still, we need to add a auto-group(`augroup`) that groups the autocmds and by adding `autocmd!` it will clear all the commands from the group. 

```vimscript
augroup python_format
    autocmd!
    autocmd BufWritePost *.py silent !black %
augroup end
```
We can now add it to the vimrc to work all the time.

## Using pipx 

If you have used pipx to install black, you need to setup the autocmd a bit differently. 

```vimscript
autocmd BufWritePost *.py silent !pipx run black %
```

It might be a bit slower than running with global installation, but it is a neat way to run python package. 

So, that's it we can now write clean and safe python code without breaking a sweat in Vim. Happy Coding :)