---
article_html: "<h2>Vim and Terminal!?</h2>\n<p>Vim was made to work with the command
  line. Many beginners do not understand what are the true capabilities of Vim, myself
  included:) Vim can run terminal commands without leaving the text editor, open an
  instance of a terminal, work with shell environments, and other things depending
  on the use case.</p>\n<h2>Running Terminal/ shell commands from within Vim</h2>\n<p>You
  can run the commands from inside of Vim by just using <code>:!</code> before the
  command, this means you have to be in command mode. Just after being in command
  mode, the ! or bang operator will execute the command typed after it from the terminal(Linux/
  macOS) or your default shell(Windows -&gt; CMD/Powershell).</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>:!pwd\n</pre></div>\n\n</pre>\n\n<p>The
  above command from vim will redirect to the terminal and show the output of the
  command and return on pressing any key. In this case, it will execute the PWD command
  and just wait for the user to enter any key to return to Vim.</p>\n<p>The following
  is an example of how it could be used from Vim in Windows using Powershell as the
  default shell.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1624885870237/Ie5C-3u1B.gif\"
  alt=\"Animation.gif\" /></p>\n<p>In Windows, dir is equivalent of ls for Linux.
  That was not the best example of how a terminal can be used at its best, You can
  also use a logical operator from within vim to run multiple commands at once.</p>\n<h3>Running
  programs/ code from Vim on terminal</h3>\n<p>This becomes quite a great feature
  for making Vim from a text editor to an IDE, this can be paired with Keymaps i.e
  when the user types certain keys, the command gets executed making the code run
  from the terminal. I have already used this feature to set up Vim for python, bash,
  and other programming languages. Also, I have written an article about  <a href=\"https://dev.to/mrdestructive/vim-keymapping-guide-3olb\">keymapping</a>
  \ and Vim setup for  <a href=\"https://dev.to/mrdestructive/setting-up-vim-for-python-ej\">Python</a>
  \ and  <a href=\"https://techstructiveblog.hashnode.dev/vim-setup-for-bash-scripting\">Bash</a>,
  this will give you an idea of how to setup vim for any programming language.</p>\n<p>Vim
  can really shine in this kind of feature as it just becomes flawless and a smooth
  experience even for a beginner. We just have to run the compile the code and run
  its executable/ output file, rather for python and other interpreted languages,
  we have to just pass the file name to the interpreter and that's it.</p>\n<h2>Opening
  instance of Terminal within Vim.</h2>\n<p>Vim can also create an instance of the
  terminal within its window by making a split. This is quite similar to VS Code and
  other Text editors that have the functionality to create an instance of the terminal
  within itself. This feature is useful for developing complex systems and depending
  on the use case, it can be quite important and efficient as well.</p>\n<p>The terminal
  can be created in various ways the most preferred way is by typing in <code>:term</code>
  from Vim.\nThis will create a horizontal split from the current editor and split
  it into half. You can change the size of the split using the mouse according to
  your preference.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1624888468392/wR0JT8SBN.gif\"
  alt=\"vimtermsplit.gif\" /></p>\n<p>Here Vim has certain variables and shortcuts
  to make things even simpler, say you want to parse the current file to the terminal
  for execution. You can surely type the name manually or you can be a bit smarter
  and use % instead, the <code>%</code> symbol will parse the file name along with
  the extension in the terminal. Also <code>%:r</code> will parse filename without
  the extensions(.txt/.py/etc) to the terminal.</p>\n<p>There are many things you
  can do with terminals surely, but with Vim that even goes further than the limits.
  Terminal/command line is quite important in any development environment as it is
  an interface for the user to interact with the Operating System. Vim is quite powerful
  and behaves as a gecko for programmers because it changes itself according to our
  needs flawlessly and <strong>efficiently</strong>.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1624891655340/5f81Dpp_O.gif\"
  alt=\"vimpython.gif\" /></p>\n<p>Integrating Terminal into a Text Editor truly lights
  up the environment for development. It becomes an easy and enjoyable experience
  to test out the code without wasting much time on the actual execution process.
  Surely it needs time to set up the environment to speed things, for that understanding
  of the programming and development environment is required. Happy Viming :)</p>\n"
cover: ''
date: 2021-06-29
datetime: 2021-06-29 00:00:00+00:00
description: Vim was made to work with the command line. Many beginners do not understand
  what are the true capabilities of Vim, myself included:) Vim can run terminal comma
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-06-29-Vim-Terminal.md
html: "<h2>Vim and Terminal!?</h2>\n<p>Vim was made to work with the command line.
  Many beginners do not understand what are the true capabilities of Vim, myself included:)
  Vim can run terminal commands without leaving the text editor, open an instance
  of a terminal, work with shell environments, and other things depending on the use
  case.</p>\n<h2>Running Terminal/ shell commands from within Vim</h2>\n<p>You can
  run the commands from inside of Vim by just using <code>:!</code> before the command,
  this means you have to be in command mode. Just after being in command mode, the
  ! or bang operator will execute the command typed after it from the terminal(Linux/
  macOS) or your default shell(Windows -&gt; CMD/Powershell).</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n\n<div class=\"highlight\"><pre><span></span>:!pwd\n</pre></div>\n\n</pre>\n\n<p>The
  above command from vim will redirect to the terminal and show the output of the
  command and return on pressing any key. In this case, it will execute the PWD command
  and just wait for the user to enter any key to return to Vim.</p>\n<p>The following
  is an example of how it could be used from Vim in Windows using Powershell as the
  default shell.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1624885870237/Ie5C-3u1B.gif\"
  alt=\"Animation.gif\" /></p>\n<p>In Windows, dir is equivalent of ls for Linux.
  That was not the best example of how a terminal can be used at its best, You can
  also use a logical operator from within vim to run multiple commands at once.</p>\n<h3>Running
  programs/ code from Vim on terminal</h3>\n<p>This becomes quite a great feature
  for making Vim from a text editor to an IDE, this can be paired with Keymaps i.e
  when the user types certain keys, the command gets executed making the code run
  from the terminal. I have already used this feature to set up Vim for python, bash,
  and other programming languages. Also, I have written an article about  <a href=\"https://dev.to/mrdestructive/vim-keymapping-guide-3olb\">keymapping</a>
  \ and Vim setup for  <a href=\"https://dev.to/mrdestructive/setting-up-vim-for-python-ej\">Python</a>
  \ and  <a href=\"https://techstructiveblog.hashnode.dev/vim-setup-for-bash-scripting\">Bash</a>,
  this will give you an idea of how to setup vim for any programming language.</p>\n<p>Vim
  can really shine in this kind of feature as it just becomes flawless and a smooth
  experience even for a beginner. We just have to run the compile the code and run
  its executable/ output file, rather for python and other interpreted languages,
  we have to just pass the file name to the interpreter and that's it.</p>\n<h2>Opening
  instance of Terminal within Vim.</h2>\n<p>Vim can also create an instance of the
  terminal within its window by making a split. This is quite similar to VS Code and
  other Text editors that have the functionality to create an instance of the terminal
  within itself. This feature is useful for developing complex systems and depending
  on the use case, it can be quite important and efficient as well.</p>\n<p>The terminal
  can be created in various ways the most preferred way is by typing in <code>:term</code>
  from Vim.\nThis will create a horizontal split from the current editor and split
  it into half. You can change the size of the split using the mouse according to
  your preference.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1624888468392/wR0JT8SBN.gif\"
  alt=\"vimtermsplit.gif\" /></p>\n<p>Here Vim has certain variables and shortcuts
  to make things even simpler, say you want to parse the current file to the terminal
  for execution. You can surely type the name manually or you can be a bit smarter
  and use % instead, the <code>%</code> symbol will parse the file name along with
  the extension in the terminal. Also <code>%:r</code> will parse filename without
  the extensions(.txt/.py/etc) to the terminal.</p>\n<p>There are many things you
  can do with terminals surely, but with Vim that even goes further than the limits.
  Terminal/command line is quite important in any development environment as it is
  an interface for the user to interact with the Operating System. Vim is quite powerful
  and behaves as a gecko for programmers because it changes itself according to our
  needs flawlessly and <strong>efficiently</strong>.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1624891655340/5f81Dpp_O.gif\"
  alt=\"vimpython.gif\" /></p>\n<p>Integrating Terminal into a Text Editor truly lights
  up the environment for development. It becomes an easy and enjoyable experience
  to test out the code without wasting much time on the actual execution process.
  Surely it needs time to set up the environment to speed things, for that understanding
  of the programming and development environment is required. Happy Viming :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/vim-terminal-integration.webp
long_description: Vim was made to work with the command line. Many beginners do not
  understand what are the true capabilities of Vim, myself included:) Vim can run
  terminal commands without leaving the text editor, open an instance of a terminal,
  work with shell envir
now: 2025-05-01 18:11:33.311667
path: blog/posts/2021-06-29-Vim-Terminal.md
prevnext: null
slug: vim-plus-teminal
status: published
subtitle: To feel and live in Vim with the terminal.
tags:
- vim
- linux
templateKey: blog-post
title: 'Vim: Terminal Integration'
today: 2025-05-01
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