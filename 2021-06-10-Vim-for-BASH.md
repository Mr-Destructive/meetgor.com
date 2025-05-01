---
article_html: "<h2>Vim and BASH?</h2>\n<p>Bash Scripting is a powerful skill to have
  as a programmer because we find Linux almost everywhere and to get through it you
  must have a command over its interface which is generally the BASH shell. Vim is
  a great option for doing this, or probably the best out there! Because Vim is pre-installed
  in almost every Linux distribution. This is not an in-depth setup for BASH on Vim,
  it is a simple editorial on starting up BASH scripting on the Vim editor. So without
  wasting time on &quot;Vim features&quot; let's dive in with the setup for BASH in
  Vim.</p>\n<h2>Boilerplate macro</h2>\n<p>Setting up a bash script doesn't require
  much code but still in some cases it can be a bit hassle and to avoid the repetitive
  task, one can easily set up a macro for the boilerplate BASH script.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">nnoremap</span> <span class=\"nb\">bs</span> <span class=\"k\">i</span>#<span
  class=\"p\">!</span><span class=\"sr\">/bin/</span>bash/<span class=\"p\">&lt;</span>ESC<span
  class=\"p\">&gt;</span><span class=\"k\">o</span>\n</pre></div>\n\n</pre>\n\n<p>Ok
  that was pretty dumb but it can scale pretty quickly and it will be nice to tailor
  it as per needs, here's some snippet with function pre-loaded.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">nnoremap</span> <span class=\"nb\">bs</span> <span class=\"k\">i</span>#<span
  class=\"p\">!</span><span class=\"sr\">/bin/</span>bash/<span class=\"p\">&lt;</span>ESC<span
  class=\"p\">&gt;</span><span class=\"k\">o</span>\n<span class=\"nb\">nnoremap</span>
  <span class=\"nb\">bs</span> <span class=\"k\">i</span>#<span class=\"p\">!</span><span
  class=\"sr\">/bin/</span>bash/<span class=\"p\">&lt;</span>ESC<span class=\"p\">&gt;</span><span
  class=\"k\">o</span><span class=\"p\">&lt;</span>ESC<span class=\"p\">&gt;</span>ofunction
  main<span class=\"p\">()</span>{<span class=\"p\">&lt;</span>ESC<span class=\"p\">&gt;</span><span
  class=\"k\">o</span><span class=\"p\">&lt;</span>ESC<span class=\"p\">&gt;</span><span
  class=\"k\">o</span>}<span class=\"p\">&lt;</span>ESC<span class=\"p\">&gt;</span>ki<span
  class=\"p\">&lt;</span>S<span class=\"p\">-</span>TAB<span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://s6.gifyu.com/images/bsclip.gif\" alt=\"type bs to load boiler-plate
  code\" /></p>\n<p>When the key bs is typed in normal mode, you enter into insert
  mode as per the command macro, then we type in the required text and escape to move
  to the next line and continue the same stuff. This could be extended further like
  making some input or printing out some text and any other formatted text that you
  could think it as repetition.</p>\n<h2>Sourcing Scripts</h2>\n<p>So, after creating
  the file, sourcing the script, and running it can be a bit slow for some people,
  as you have to go to the terminal and toggle in the permission to run the script
  and then run, But pull on your seatbelts as this is VIM! You can die due to slowness!</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">nnoremap</span> <span class=\"k\">sh</span> :<span class=\"p\">!</span>chmod
  <span class=\"p\">+</span><span class=\"k\">x</span> % &amp;&amp; source %\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://s6.gifyu.com/images/shclip.gif\" alt=\"type sh to run script\" /></p>\n<p>When
  the sh keys are typed in the normal mode, the preceding command after ! (bang) will
  be executed in the terminal, the &amp;&amp; keywords will execute the second command
  only when the first command is successfully executed.\nThat just can so fast! Imagine
  doing this for long scripts and especially for debugging, it will waste 2 minutes
  every time you leave the editor and for 10 times you do the debugging, you will
  carelessly was roughly 20 minutes! Improve your debugging skills surely :)</p>\n<h2>Plugins</h2>\n<p>There
  are very few plugins out there for BASH as for VIM, but it's quite to write scripts
  even without any plugins. One of the most supported and popular plugins for BASH
  in Vim is  <a href=\"https://www.vim.org/scripts/script.php?script_id=365\">Bash-Support-Vim</a>
  for auto-completion and <a href=\"https://www.shellcheck.net\">Shell-Check</a> for
  finding/correcting any bugs or error in the script .\nThe mentioned plugin is quite
  awesome and it can greatly improve the speed of scripting for BASH, some commands
  such as shortcuts for writing if-else, while, for loops, commenting and other aspects
  in the scripting. The thorough documentation for such commands is also provided
  by the  <a href=\"https://wolfgangmehner.github.io/vim-plugins/bashsupport.html\">plugin
  website</a>.\nThis can be used for autocompleting keywords and writing nested if-else
  and other logical operators in BASH scripting. Again, you can do absolutely fine
  without plugins in Vim as it is heavily customizable to the user's need and can
  be very rewarding to set up your own configuration for BASH. You can use standard
  Vim(barebones) for auto-completion as well with the command CTRL+N and CTRL-P to
  move down and up respectively.</p>\n<h2>Some More Tricks</h2>\n<p>BASH in Vim can
  be quite versatile to use as it provides some custom addons to make the script more
  functional and easier to understand. Some tricks such as using autocompletion can
  be quite inconvenient to use at once but it can get really smooth after some runs
  at writing the scripts.</p>\n<ul>\n<li>In BASH Scripts there are quite a lot of
  brackets to play with that's why to jump around swiftly around such parentheses
  or brackets you can use <strong>% to move from the opened to closed brackets or
  vice versa</strong>.</li>\n<li>You can execute any terminal command from Vim, be
  sure to be in command mode and press ! after the command, you would like to execute.
  This will run the command from the terminal and you don't have to leave the editor,
  it saves a ton of time and it's blazingly fast.</li>\n<li>With the above trick,
  you kind of have a superpower within Vim to make, build, source, run the files or
  scripts within Vim, that is not repetition but it can run bash within bash. Ok!
  that's was pretty fast. Don't die of quickness now!</li>\n</ul>\n<p>Writing BASH
  scripts in Vim can be also boosted by using some built-in commands such as adding
  comments for multiple lines at once and some unexplored stuff which can be learned
  in the way to understanding the flow of Vim and BASH together. Happy Coding and
  Viming :)</p>\n"
cover: ''
date: 2021-06-10
datetime: 2021-06-10 00:00:00+00:00
description: Bash Scripting is a powerful skill to have as a programmer because we
  find Linux almost everywhere and to get through it you must have a command over
  its interf
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-06-10-Vim-for-BASH.md
html: "<h2>Vim and BASH?</h2>\n<p>Bash Scripting is a powerful skill to have as a
  programmer because we find Linux almost everywhere and to get through it you must
  have a command over its interface which is generally the BASH shell. Vim is a great
  option for doing this, or probably the best out there! Because Vim is pre-installed
  in almost every Linux distribution. This is not an in-depth setup for BASH on Vim,
  it is a simple editorial on starting up BASH scripting on the Vim editor. So without
  wasting time on &quot;Vim features&quot; let's dive in with the setup for BASH in
  Vim.</p>\n<h2>Boilerplate macro</h2>\n<p>Setting up a bash script doesn't require
  much code but still in some cases it can be a bit hassle and to avoid the repetitive
  task, one can easily set up a macro for the boilerplate BASH script.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">nnoremap</span> <span class=\"nb\">bs</span> <span class=\"k\">i</span>#<span
  class=\"p\">!</span><span class=\"sr\">/bin/</span>bash/<span class=\"p\">&lt;</span>ESC<span
  class=\"p\">&gt;</span><span class=\"k\">o</span>\n</pre></div>\n\n</pre>\n\n<p>Ok
  that was pretty dumb but it can scale pretty quickly and it will be nice to tailor
  it as per needs, here's some snippet with function pre-loaded.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">nnoremap</span> <span class=\"nb\">bs</span> <span class=\"k\">i</span>#<span
  class=\"p\">!</span><span class=\"sr\">/bin/</span>bash/<span class=\"p\">&lt;</span>ESC<span
  class=\"p\">&gt;</span><span class=\"k\">o</span>\n<span class=\"nb\">nnoremap</span>
  <span class=\"nb\">bs</span> <span class=\"k\">i</span>#<span class=\"p\">!</span><span
  class=\"sr\">/bin/</span>bash/<span class=\"p\">&lt;</span>ESC<span class=\"p\">&gt;</span><span
  class=\"k\">o</span><span class=\"p\">&lt;</span>ESC<span class=\"p\">&gt;</span>ofunction
  main<span class=\"p\">()</span>{<span class=\"p\">&lt;</span>ESC<span class=\"p\">&gt;</span><span
  class=\"k\">o</span><span class=\"p\">&lt;</span>ESC<span class=\"p\">&gt;</span><span
  class=\"k\">o</span>}<span class=\"p\">&lt;</span>ESC<span class=\"p\">&gt;</span>ki<span
  class=\"p\">&lt;</span>S<span class=\"p\">-</span>TAB<span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://s6.gifyu.com/images/bsclip.gif\" alt=\"type bs to load boiler-plate
  code\" /></p>\n<p>When the key bs is typed in normal mode, you enter into insert
  mode as per the command macro, then we type in the required text and escape to move
  to the next line and continue the same stuff. This could be extended further like
  making some input or printing out some text and any other formatted text that you
  could think it as repetition.</p>\n<h2>Sourcing Scripts</h2>\n<p>So, after creating
  the file, sourcing the script, and running it can be a bit slow for some people,
  as you have to go to the terminal and toggle in the permission to run the script
  and then run, But pull on your seatbelts as this is VIM! You can die due to slowness!</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">nnoremap</span> <span class=\"k\">sh</span> :<span class=\"p\">!</span>chmod
  <span class=\"p\">+</span><span class=\"k\">x</span> % &amp;&amp; source %\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://s6.gifyu.com/images/shclip.gif\" alt=\"type sh to run script\" /></p>\n<p>When
  the sh keys are typed in the normal mode, the preceding command after ! (bang) will
  be executed in the terminal, the &amp;&amp; keywords will execute the second command
  only when the first command is successfully executed.\nThat just can so fast! Imagine
  doing this for long scripts and especially for debugging, it will waste 2 minutes
  every time you leave the editor and for 10 times you do the debugging, you will
  carelessly was roughly 20 minutes! Improve your debugging skills surely :)</p>\n<h2>Plugins</h2>\n<p>There
  are very few plugins out there for BASH as for VIM, but it's quite to write scripts
  even without any plugins. One of the most supported and popular plugins for BASH
  in Vim is  <a href=\"https://www.vim.org/scripts/script.php?script_id=365\">Bash-Support-Vim</a>
  for auto-completion and <a href=\"https://www.shellcheck.net\">Shell-Check</a> for
  finding/correcting any bugs or error in the script .\nThe mentioned plugin is quite
  awesome and it can greatly improve the speed of scripting for BASH, some commands
  such as shortcuts for writing if-else, while, for loops, commenting and other aspects
  in the scripting. The thorough documentation for such commands is also provided
  by the  <a href=\"https://wolfgangmehner.github.io/vim-plugins/bashsupport.html\">plugin
  website</a>.\nThis can be used for autocompleting keywords and writing nested if-else
  and other logical operators in BASH scripting. Again, you can do absolutely fine
  without plugins in Vim as it is heavily customizable to the user's need and can
  be very rewarding to set up your own configuration for BASH. You can use standard
  Vim(barebones) for auto-completion as well with the command CTRL+N and CTRL-P to
  move down and up respectively.</p>\n<h2>Some More Tricks</h2>\n<p>BASH in Vim can
  be quite versatile to use as it provides some custom addons to make the script more
  functional and easier to understand. Some tricks such as using autocompletion can
  be quite inconvenient to use at once but it can get really smooth after some runs
  at writing the scripts.</p>\n<ul>\n<li>In BASH Scripts there are quite a lot of
  brackets to play with that's why to jump around swiftly around such parentheses
  or brackets you can use <strong>% to move from the opened to closed brackets or
  vice versa</strong>.</li>\n<li>You can execute any terminal command from Vim, be
  sure to be in command mode and press ! after the command, you would like to execute.
  This will run the command from the terminal and you don't have to leave the editor,
  it saves a ton of time and it's blazingly fast.</li>\n<li>With the above trick,
  you kind of have a superpower within Vim to make, build, source, run the files or
  scripts within Vim, that is not repetition but it can run bash within bash. Ok!
  that's was pretty fast. Don't die of quickness now!</li>\n</ul>\n<p>Writing BASH
  scripts in Vim can be also boosted by using some built-in commands such as adding
  comments for multiple lines at once and some unexplored stuff which can be learned
  in the way to understanding the flow of Vim and BASH together. Happy Coding and
  Viming :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/vim-for-bash.webp
long_description: Bash Scripting is a powerful skill to have as a programmer because
  we find Linux almost everywhere and to get through it you must have a command over
  its interface which is generally the BASH shell. Vim is a great option for doing
  this, or probably t
now: 2025-05-01 18:11:33.312261
path: blog/posts/2021-06-10-Vim-for-BASH.md
prevnext: null
slug: vim-for-bash
status: published
subtitle: VIM + BASH = Killer Speed Automation
tags:
- bash
- vim
templateKey: blog-post
title: Setting up Vim for BASH Scripting
today: 2025-05-01
---

## Vim and BASH? 
Bash Scripting is a powerful skill to have as a programmer because we find Linux almost everywhere and to get through it you must have a command over its interface which is generally the BASH shell. Vim is a great option for doing this, or probably the best out there! Because Vim is pre-installed in almost every Linux distribution. This is not an in-depth setup for BASH on Vim, it is a simple editorial on starting up BASH scripting on the Vim editor. So without wasting time on "Vim features" let's dive in with the setup for BASH in Vim.

## Boilerplate macro
Setting up a bash script doesn't require much code but still in some cases it can be a bit hassle and to avoid the repetitive task, one can easily set up a macro for the boilerplate BASH script.

```vim
nnoremap bs i#!/bin/bash/<ESC>o
```
Ok that was pretty dumb but it can scale pretty quickly and it will be nice to tailor it as per needs, here's some snippet with function pre-loaded.

```vim
nnoremap bs i#!/bin/bash/<ESC>o
nnoremap bs i#!/bin/bash/<ESC>o<ESC>ofunction main(){<ESC>o<ESC>o}<ESC>ki<S-TAB>

```
![type bs to load boiler-plate code](https://s6.gifyu.com/images/bsclip.gif)

When the key bs is typed in normal mode, you enter into insert mode as per the command macro, then we type in the required text and escape to move to the next line and continue the same stuff. This could be extended further like making some input or printing out some text and any other formatted text that you could think it as repetition. 

## Sourcing Scripts
So, after creating the file, sourcing the script, and running it can be a bit slow for some people, as you have to go to the terminal and toggle in the permission to run the script and then run, But pull on your seatbelts as this is VIM! You can die due to slowness!

```vim
nnoremap sh :!chmod +x % && source %
```

![type sh to run script](https://s6.gifyu.com/images/shclip.gif)

When the sh keys are typed in the normal mode, the preceding command after ! (bang) will be executed in the terminal, the && keywords will execute the second command only when the first command is successfully executed.
 That just can so fast! Imagine doing this for long scripts and especially for debugging, it will waste 2 minutes every time you leave the editor and for 10 times you do the debugging, you will carelessly was roughly 20 minutes! Improve your debugging skills surely :)

## Plugins
There are very few plugins out there for BASH as for VIM, but it's quite to write scripts even without any plugins. One of the most supported and popular plugins for BASH in Vim is  [Bash-Support-Vim](https://www.vim.org/scripts/script.php?script_id=365) for auto-completion and [Shell-Check](https://www.shellcheck.net) for finding/correcting any bugs or error in the script . 
The mentioned plugin is quite awesome and it can greatly improve the speed of scripting for BASH, some commands such as shortcuts for writing if-else, while, for loops, commenting and other aspects in the scripting. The thorough documentation for such commands is also provided by the  [plugin website](https://wolfgangmehner.github.io/vim-plugins/bashsupport.html). 
This can be used for autocompleting keywords and writing nested if-else and other logical operators in BASH scripting. Again, you can do absolutely fine without plugins in Vim as it is heavily customizable to the user's need and can be very rewarding to set up your own configuration for BASH. You can use standard Vim(barebones) for auto-completion as well with the command CTRL+N and CTRL-P to move down and up respectively.


## Some More Tricks
BASH in Vim can be quite versatile to use as it provides some custom addons to make the script more functional and easier to understand. Some tricks such as using autocompletion can be quite inconvenient to use at once but it can get really smooth after some runs at writing the scripts.
- In BASH Scripts there are quite a lot of brackets to play with that's why to jump around swiftly around such parentheses or brackets you can use **% to move from the opened to closed brackets or vice versa**.
- You can execute any terminal command from Vim, be sure to be in command mode and press ! after the command, you would like to execute. This will run the command from the terminal and you don't have to leave the editor, it saves a ton of time and it's blazingly fast.
- With the above trick, you kind of have a superpower within Vim to make, build, source, run the files or scripts within Vim, that is not repetition but it can run bash within bash. Ok! that's was pretty fast. Don't die of quickness now!

Writing BASH scripts in Vim can be also boosted by using some built-in commands such as adding comments for multiple lines at once and some unexplored stuff which can be learned in the way to understanding the flow of Vim and BASH together. Happy Coding and Viming :)