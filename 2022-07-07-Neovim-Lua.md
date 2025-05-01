---
article_html: "<h2>Introduction</h2>\n<p>It has been a while since I have written
  a Vim article. Finally, I got some ideas after configuring my Neovim setup for Lua.
  I recently migrated to Ubuntu a couple of months back and it has been a cool change
  from Windows 7!</p>\n<p>In this article, we'll see how you can set up neovim for
  Lua. Since Neovim 0.5, it supports lua out of the box, so in the recent release
  0.7, it added more native support to lua making it a lot easier to configure and
  play with neovim. So, we will see how we can use lua to convert all the 200 liner
  vimscript into lua (We can even have packages and modules:) We will cover how to
  configure your keymaps, pull up all the plugins, vim options, and other customizations.</p>\n<h2>Why
  move to Lua?</h2>\n<p>I have used Vimscript for quite a while now, configured it
  as per my needs, and also made a few plugins like <a href=\"https://github.com/Mr-Destructive/frontmatter.vim\">frontmatter</a>,
  <a href=\"https://github.com/Mr-Destructive/dj.vim\">dj.vim</a>, and <a href=\"https://github.com/Mr-Destructive/commenter.vim\">commenter</a>
  which are quite clunky and not robust in terms of usage and customizability. Vimscript
  is good but it's a bit messy when you want extreme customization.</p>\n<p>I recently
  wanted to go full Neovim, I was kind of stuck in Vimscript and some of my plugins
  work for me but it might not work for others, it might be a piece of gibberish vimscript
  dumped. So, Why not have full native experience in Neovim if you can. It has now
  baked-in support for LSP, Debugging, Autocommands, and so much more. If you have
  Neovim 0.5+ you should be good to go full lua.</p>\n<h2>Backup Neovim Config</h2>\n<p>Firstly,
  keep your original neovim/vim init files safe, back them up, make a copy and save
  it with a different name like <code>nvim_config.vim</code>. If you already have
  all your config files backed up as an ansible script or dotfiles GitHub repository,
  you can proceed ahead.</p>\n<p>But don't keep the <code>init.vim</code> file as
  it is in the <code>~/.config/nvim</code> folder, it might override after we configure
  the lua scripts.</p>\n<h2>Basic Configuration</h2>\n<p>So, I assume you have set
  up Neovim, If not you need to follow some simple steps like downloading the package
  and making sure your neovim environment is working with vimscript first. The <a
  href=\"https://github.com/neovim/neovim/wiki/Installing-Neovim\">Neovim Wiki</a>
  provides great documentation on how to install neovim on various platforms using
  different methods.</p>\n<p>After your neovim is set up and you have a basic configuration,
  you can now start to migrate into lua.\nCreate a <code>init.lua</code> file in the
  same path as your <code>init.lua</code> file is i.e. at <code>~/.config/nvim</code>
  or <code>~/AppData/Local/nvim/</code> for Windows. That's why it is recommended
  to keep the initial configuration vimscript file in a safe place. While migrating
  from vimscript to lua, once the lua file is created and the next time you restart
  neovim, the default settings will be from <code>init.lua</code> and not <code>init.vim</code>,
  so be prepared.</p>\n<p>Firstly, you need to configure some options like <code>number</code>,
  <code>syntax highlighting</code>, <code>tabs</code>, and some <code>keymaps</code>
  of course. We can use the <code>vim.opt</code> method to set options in vim using
  lua syntax. So, certain corresponding vim options would be converted as follows:</p>\n<p>If
  you have the following kind of settings in your vimrc or init.vim:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>vimscript</div>\n<div class=\"highlight\"><pre><span></span>--
  vimscript\nset number\nset tabstop=4 \nset shiftwidth=4 \nset softtabstop=0 \nset
  expandtab \nset noswapfile\n</pre></div>\n\n</pre>\n\n<p>The above settings are
  migrated into lua syntax as follows:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">--lua</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">opt</span><span class=\"p\">.</span><span class=\"n\">number</span>
  <span class=\"o\">=</span> <span class=\"kc\">true</span>\n<span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">opt</span><span class=\"p\">.</span><span
  class=\"n\">tabstop</span> <span class=\"o\">=</span> <span class=\"mi\">4</span>\n<span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">opt</span><span
  class=\"p\">.</span><span class=\"n\">shiftwidth</span> <span class=\"o\">=</span>
  <span class=\"mi\">4</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">opt</span><span class=\"p\">.</span><span class=\"n\">softtabstop</span>
  <span class=\"o\">=</span> <span class=\"mi\">0</span>\n<span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">opt</span><span class=\"p\">.</span><span
  class=\"n\">expandtab</span> <span class=\"o\">=</span> <span class=\"kc\">true</span>\n<span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">opt</span><span
  class=\"p\">.</span><span class=\"n\">swapfile</span> <span class=\"o\">=</span>
  <span class=\"kc\">false</span>\n</pre></div>\n\n</pre>\n\n<p>You can set other
  options in your config file accordingly. If you get sick of writing <code>vim.opt.</code>
  again and again, you can use a variable set to <code>vim.opt</code> and then access
  that variable to set the options. Something of the lines as below:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kd\">local</span> <span class=\"n\">set</span> <span class=\"o\">=</span>
  <span class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">opt</span>\n\n<span
  class=\"n\">set</span><span class=\"p\">.</span><span class=\"n\">number</span>
  <span class=\"o\">=</span> <span class=\"kc\">true</span>\n<span class=\"n\">set</span><span
  class=\"p\">.</span><span class=\"n\">tabstop</span> <span class=\"o\">=</span>
  <span class=\"mi\">4</span>\n<span class=\"n\">set</span><span class=\"p\">.</span><span
  class=\"n\">shiftwidth</span> <span class=\"o\">=</span> <span class=\"mi\">4</span>\n<span
  class=\"n\">set</span><span class=\"p\">.</span><span class=\"n\">softtabstop</span>
  <span class=\"o\">=</span> <span class=\"mi\">0</span>\n<span class=\"n\">set</span><span
  class=\"p\">.</span><span class=\"n\">expandtab</span> <span class=\"o\">=</span>
  <span class=\"kc\">true</span>\n<span class=\"n\">set</span><span class=\"p\">.</span><span
  class=\"n\">swapfile</span> <span class=\"o\">=</span> <span class=\"kc\">false</span>\n</pre></div>\n\n</pre>\n\n<p>We
  can create a variable in lua like <code>local variable_name = something</code> so,
  we have created a variable <code>set</code> which is assigned to the value of <code>vim.opt</code>
  which is a method(function) in lua to set the options from the vimscript environment.
  Finally, access the <code>set</code> keyword to set the options. Using the <code>set</code>
  word is not necessary, you can use whatever you want.</p>\n<p>After setting up the
  basic options, you can source the file with <code>:so %</code> from the command
  mode. Just normally as you source the config files.</p>\n<h3>Using Lua in Command
  Mode</h3>\n<p>We can use the lua functions or any other commands from the command
  mode in neovim using the lua command. Just prefix the command with <code>:lua</code>
  and after that, you can use lua syntax like function calling or setting variables,
  logging things, etc.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657380885/blog-media/lua_in_command_mode.gif\"
  alt=\"Lua in Command Mode\" /></p>\n<h2>Adding Keymaps</h2>\n<p>Now, that we have
  some basic config setup, we can quickly get the keymaps. It's not that hard to make
  keymaps to set up in lua. To create keymaps, we have two options. One is compatible
  with Neovim and the other is compatible with Vim as well.</p>\n<ol>\n<li>vim.keymap.set
  OR</li>\n<li>vim.api.nvim_set_keymap</li>\n</ol>\n<p>Personally, I don't see a difference
  in terms of usage but <a href=\"https://github.com/neovim/neovim/blob/master/runtime/lua/vim/keymap.lua#L51\">vim.keymap.set</a>
  is a wrapper around <a href=\"https://github.com/neovim/neovim/blob/master/src/nvim/api/vim.c#L1451\">nvim_set_keymap</a>.
  So, it is really a matter of personal preference which you want to use.</p>\n<p>So,
  both the functions have quite similar parameters:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">keymap</span><span
  class=\"p\">.</span><span class=\"n\">set</span><span class=\"p\">({</span><span
  class=\"n\">mode</span><span class=\"p\">},</span> <span class=\"p\">{</span><span
  class=\"n\">lhs</span><span class=\"p\">},</span> <span class=\"p\">{</span><span
  class=\"n\">rhs</span><span class=\"p\">},</span> <span class=\"p\">{</span><span
  class=\"n\">options</span><span class=\"p\">})</span>\n\n<span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">api</span><span class=\"p\">.</span><span
  class=\"n\">nvim_set_keymap</span><span class=\"p\">({</span><span class=\"n\">mode</span><span
  class=\"p\">},</span> <span class=\"p\">{</span><span class=\"n\">lhs</span><span
  class=\"p\">},</span> <span class=\"p\">{</span><span class=\"n\">rhs</span><span
  class=\"p\">},</span> <span class=\"p\">{</span><span class=\"n\">options</span><span
  class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p>The advantage of <code>vim.keymap.set</code>
  over <code>vim.api.nvim_set_keymap</code> is that it allows directly calling lua
  functions rather than calling vimscripty way like <code>:lua function()</code>,
  so we directly can use lua code in the RHS part of the function parameter.</p>\n<p>Let's
  take a basic example mapping:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n\n<div class=\"highlight\"><pre><span></span>vim.keymap.set(&#39;n&#39;,
  &#39;Y&#39;, &#39;yy&#39;, {noremap = false})\n</pre></div>\n\n</pre>\n\n<p>Here,
  we have mapped <code>Shift+y</code> with the keys <code>yy</code> in <code>n</code>ormal
  mode. The first argument is the mode, it can be a single-mode like <code>'n'</code>,
  <code>'v'</code>, <code>'i'</code>, etc., or a multi-mode table like <code>{'n',
  'v'}</code>, <code>{'v', 'i'}</code>, etc.</p>\n<p>The next parameter is also a
  string but it should be the key for triggering the mapping. In this case, we have
  used <code>Y</code> which is <code>Shift + y</code>, it can be any key shortcut
  you want to map.</p>\n<p>The third parameter is the string which will be the command
  to be executed after the key has been used. Here we have used the keys <code>yy</code>,
  if the map is from a command mode, you will be using something like <code>':commands_to_be
  executed'</code> as the third parameter.</p>\n<p>The fourth parameter which is optional
  can contain <a href=\"https://neovim.io/doc/user/api.html#:~:text=nvim_set_keymap(%7Bmode%7D%2C%20%7Blhs%7D%2C%20%7Brhs%7D%2C%20%7B*opts%7D)%20%20%20%20%20%20%20%20%20%20%20%20%20*nvim_set_keymap()*\">special
  arguments</a>. We have set a default option which is <code>noremap</code> as true,
  the options are not string but lua tables instead, so it can be similar to python
  dictionary or a map kind of a structure with a key value pair.</p>\n<p>One more
  important aspect in keymapping might about the leader key, you can set your leader
  key by using the global vim options with <code>vim.g</code> and access <code>mapleader</code>
  to set it to the key you wish. This will make the <code>leader</code> key available
  to us and thereafter we can map the leader key in custom mappings.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>vim.g.mapleader = &quot;
  &quot;\n</pre></div>\n\n</pre>\n\n<p>Here, I have set my leader key to the <code>&lt;Space&gt;</code>
  key. Now, we can map keys to the existing keymaps in the vimscript. Let's map some
  basic keymaps first and then after setting up the plugins,we can move into plugin-specific
  mappings.</p>\n<p>You can also use <code>vim.api.nvim_set_keymap</code> function
  with the same parameters as well.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">keymap</span><span
  class=\"p\">.</span><span class=\"n\">set</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;w&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;:w&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">keymap</span><span class=\"p\">.</span><span class=\"n\">set</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;&lt;leader&gt;q&#39;</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;:q!&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span class=\"n\">noremap</span>
  <span class=\"o\">=</span> <span class=\"kc\">true</span><span class=\"p\">})</span>\n<span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">keymap</span><span
  class=\"p\">.</span><span class=\"n\">set</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;s&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;:so %&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">keymap</span><span class=\"p\">.</span><span class=\"n\">set</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;&lt;leader&gt;ev&#39;</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;:vsplit $MYVIMRC&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">keymap</span><span class=\"p\">.</span><span class=\"n\">set</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;&lt;leader&gt;sv&#39;</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;:w&lt;CR&gt;:so %&lt;CR&gt;:q&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n\n<span class=\"n\">OR</span>\n\n<span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">api</span><span class=\"p\">.</span><span
  class=\"n\">nvim_set_keymap</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;w&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;:w&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">api</span><span class=\"p\">.</span><span class=\"n\">nvim_set_keymap</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;&lt;leader&gt;q&#39;</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;:q!&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span class=\"n\">noremap</span>
  <span class=\"o\">=</span> <span class=\"kc\">true</span><span class=\"p\">})</span>\n<span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">api</span><span
  class=\"p\">.</span><span class=\"n\">nvim_set_keymap</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;s&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;:so %&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">api</span><span class=\"p\">.</span><span class=\"n\">nvim_set_keymap</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;&lt;leader&gt;ev&#39;</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;:vsplit $MYVIMRC&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">api</span><span class=\"p\">.</span><span class=\"n\">nvim_set_keymap</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;&lt;leader&gt;sv&#39;</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;:w&lt;CR&gt;:so %&lt;CR&gt;:q&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p>If, you don't like writing <code>vim.keymap.set</code>
  or <code>vim.api.nvim_set_keymap</code> again and again, you can create a simpler
  function for it. In lua a function is created just like a variable by specifying
  the scope of the function i.e. local followed by the <code>function</code> keyword
  and finally the name of the function and parenthesis. The function body is terminated
  by the <code>end</code> keyword.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kr\">function</span> <span class=\"nf\">map</span><span class=\"p\">(</span><span
  class=\"n\">mode</span><span class=\"p\">,</span> <span class=\"n\">lhs</span><span
  class=\"p\">,</span> <span class=\"n\">rhs</span><span class=\"p\">,</span> <span
  class=\"n\">opts</span><span class=\"p\">)</span>\n    <span class=\"kd\">local</span>
  <span class=\"n\">options</span> <span class=\"o\">=</span> <span class=\"p\">{</span>
  <span class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span>
  <span class=\"p\">}</span>\n    <span class=\"kr\">if</span> <span class=\"n\">opts</span>
  <span class=\"kr\">then</span>\n        <span class=\"n\">options</span> <span class=\"o\">=</span>
  <span class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">tbl_extend</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;force&quot;</span><span class=\"p\">,</span>
  <span class=\"n\">options</span><span class=\"p\">,</span> <span class=\"n\">opts</span><span
  class=\"p\">)</span>\n    <span class=\"kr\">end</span>\n    <span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">api</span><span class=\"p\">.</span><span
  class=\"n\">nvim_set_keymap</span><span class=\"p\">(</span><span class=\"n\">mode</span><span
  class=\"p\">,</span> <span class=\"n\">lhs</span><span class=\"p\">,</span> <span
  class=\"n\">rhs</span><span class=\"p\">,</span> <span class=\"n\">options</span><span
  class=\"p\">)</span>\n<span class=\"kr\">end</span>\n</pre></div>\n\n</pre>\n\n<p>Now,
  in this function <code>map</code>, we have passed in the same parameters like the
  <code>vim.keymap.set</code> function takes but we have just parsed the function
  in a shorter and safer way by adding <code>noremap = true</code> by default. So
  this is just a helper function or a verbose function for calling the vim.keymap.set
  function.</p>\n<p>To use this function, we can simply call <code>map</code> with
  the same arguments as given to the prior functions.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;w&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;:w&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;q&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;:q!&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;s&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;:so %&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>Notice
  here, though, we have not passed the <code>{noremap = true}</code> as it will be
  passed by default to the <code>vim.api.nvim_set_keymap</code> or <code>vim.keymap.set</code>
  function via the custom map function.</p>\n<p>If you want some more examples, here
  are some additional mapping specific to languages, meant for compiling or running
  scripts with neovim instance.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">-- vimscript</span>\n\n<span class=\"n\">nnoremap</span> <span class=\"n\">cpp</span>
  <span class=\"p\">:</span><span class=\"err\">!</span><span class=\"n\">c</span><span
  class=\"o\">++</span> <span class=\"o\">%</span> <span class=\"o\">-</span><span
  class=\"n\">o</span> <span class=\"o\">%</span><span class=\"p\">:</span><span class=\"n\">r</span>
  <span class=\"o\">&amp;&amp;</span> <span class=\"o\">%</span><span class=\"p\">:</span><span
  class=\"n\">r</span><span class=\"o\">&lt;</span><span class=\"n\">CR</span><span
  class=\"o\">&gt;</span>\n<span class=\"n\">nnoremap</span> <span class=\"n\">c</span><span
  class=\"p\">,</span> <span class=\"p\">:</span><span class=\"err\">!</span><span
  class=\"n\">gcc</span> <span class=\"o\">%</span> <span class=\"o\">-</span><span
  class=\"n\">o</span> <span class=\"o\">%</span><span class=\"p\">:</span><span class=\"n\">r</span>
  <span class=\"o\">&amp;&amp;</span> <span class=\"o\">%</span><span class=\"p\">:</span><span
  class=\"n\">r</span><span class=\"o\">&lt;</span><span class=\"n\">CR</span><span
  class=\"o\">&gt;</span>\n<span class=\"n\">nnoremap</span> <span class=\"n\">py</span>
  <span class=\"p\">:</span><span class=\"err\">!</span><span class=\"n\">python</span>
  <span class=\"o\">%&lt;</span><span class=\"n\">cr</span><span class=\"o\">&gt;</span>\n<span
  class=\"n\">nnoremap</span> <span class=\"n\">go</span> <span class=\"p\">:</span><span
  class=\"err\">!</span><span class=\"n\">go</span> <span class=\"n\">run</span> <span
  class=\"o\">%&lt;</span><span class=\"n\">cr</span><span class=\"o\">&gt;</span>\n<span
  class=\"n\">nnoremap</span> <span class=\"n\">sh</span> <span class=\"p\">:</span><span
  class=\"err\">!</span><span class=\"n\">bash</span> <span class=\"o\">%&lt;</span><span
  class=\"n\">CR</span><span class=\"o\">&gt;</span>\n\n\n<span class=\"c1\">-- lua</span>\n\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;cpp&#39;</span> <span class=\"s1\">&#39;:!c++
  % -o %:r &amp;&amp; %:r&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span class=\"n\">map</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;c,&#39;</span> <span class=\"s1\">&#39;:!gcc % -o %:r &amp;&amp;
  %:r&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span class=\"n\">map</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;py&#39;</span> <span class=\"s1\">&#39;:!python %&lt;cr&gt;&#39;</span><span
  class=\"p\">)</span>\n<span class=\"n\">map</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;go&#39;</span>
  <span class=\"s1\">&#39;:!go run %&lt;cr&gt;&#39;</span><span class=\"p\">)</span>\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;sh&#39;</span> <span class=\"s1\">&#39;:!bash
  %&lt;cr&gt;&#39;</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  this is how we can set up our keymaps in lua. You can customize this function as
  per your needs. These are just made for the understanding purpose of how to reduce
  the repetitive stuff in the setup.</p>\n<p><strong>If you are really stuck up and
  not feeling to convert those mappings into lua then I have a function that can do
  it for you, check out my dotfile repo -&gt; <a href=\"https://github.com/Mr-Destructive/dotfiles/blob/master/nvim/lua/destructive/options.lua#L9\">keymapper</a></strong></p>\n<h2>Adding
  Plugin Manager</h2>\n<p>Now, we really missing some plugins, aren't we? So, the
  neovim community has some good choices for using a new plugin manager written in
  core lua. It is usually a good idea to move into lua completely rather than switching
  to and fro between vimscript and lua.</p>\n<p>So, <a href=\"https://github.com/wbthomason/packer.nvim\">Packer</a>
  is the new plugin manager for Neovim in Lua, there is other plugin managers out
  there as well like <a href=\"https://github.com/savq/paq-nvim\">paq</a>. If you
  don't want to switch with the plugin manager, you can still use vim-based plugin
  managers like <a href=\"https://dev.to/vonheikemen/neovim-using-vim-plug-in-lua-3oom\">Vim-Plug</a>.</p>\n<p>So,
  let's install the Packer plugin manager into Neovim. We simply have to run the following
  command in the console and make sure the plugin manager is configured correctly.</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span># Linux\n\ngit clone --depth
  1 https://github.com/wbthomason/packer.nvim\\\n~/.local/share/nvim/site/pack/packer/start/packer.nvim\n\n\n#
  Windows\n\ngit clone https://github.com/wbthomason/packer.nvim &quot;$env:LOCALAPPDATA\\nvim-data\\site\\pack\\packer\\start\\packer.nvim&quot;\n</pre></div>\n\n</pre>\n\n<p>Now,
  if you open a new neovim instance and try to run the command <code>:PackerClean</code>,
  and no error pops out that means you have configured it correctly. You can move
  ahead to installing plugins now. Yeah! PLUG-IN time!</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kr\">return</span> <span class=\"nb\">require</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;packer&#39;</span><span class=\"p\">).</span><span class=\"n\">startup</span><span
  class=\"p\">(</span><span class=\"kr\">function</span><span class=\"p\">()</span>\n<span
  class=\"kr\">end</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>First
  try to source the file, if it throws out errors it shouldn't try to fix the installation
  path of Packer. If the command succeded we can finally pull up some plugins.</p>\n<p>Below
  are some of the plugins that you can use irrespective of what language preferences
  you would have. This includes basic dev-icons for the status line as well as the
  explorer window file icons. As usual, add your plugins and make them yours.</p>\n<pre
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kr\">return</span> <span class=\"nb\">require</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;packer&#39;</span><span class=\"p\">).</span><span class=\"n\">startup</span><span
  class=\"p\">(</span><span class=\"kr\">function</span><span class=\"p\">()</span>\n
  \ <span class=\"n\">use</span> <span class=\"s1\">&#39;wbthomason/packer.nvim&#39;</span>\n
  \ <span class=\"n\">use</span> <span class=\"s1\">&#39;tpope/vim-fugitive&#39;</span>\n
  \ <span class=\"n\">use</span> <span class=\"p\">{</span>\n    <span class=\"s1\">&#39;nvim-lualine/lualine.nvim&#39;</span><span
  class=\"p\">,</span>\n    <span class=\"n\">requires</span> <span class=\"o\">=</span>
  <span class=\"p\">{</span> <span class=\"s1\">&#39;kyazdani42/nvim-web-devicons&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">opt</span> <span class=\"o\">=</span> <span
  class=\"kc\">true</span> <span class=\"p\">}</span>\n  <span class=\"p\">}</span>\n
  \ <span class=\"n\">use</span> <span class=\"s1\">&#39;tiagofumo/vim-nerdtree-syntax-highlight&#39;</span>\n
  \ <span class=\"n\">use</span> <span class=\"s1\">&#39;kyazdani42/nvim-web-devicons&#39;</span>\n
  \ <span class=\"n\">use</span> <span class=\"s1\">&#39;vim-airline/vim-airline&#39;</span>\n
  \ <span class=\"n\">use</span> <span class=\"s1\">&#39;vim-airline/vim-airline-themes&#39;</span>\n<span
  class=\"kr\">end</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>After
  adding the list of your plugins, you need to source the file and then install the
  plugins with the command <code>:PackerInstall</code>. This function will install
  all the plugins after the file has been sourced so make sure you don't forget it.</p>\n<h2>Colors
  and Color Themes</h2>\n<p>You might fancy some good-looking and aesthetic setup
  for neovim of course! In Neovim, we already have a wide variety of configurations
  to set up like color schemes, GUI colors, terminal colors, etc. You can pick up
  a color scheme from a large list of awesome color schemes from <a href=\"https://github.com/topics/neovim-colorscheme\">GitHub</a>.</p>\n<p>After
  choosing the theme, plug it in the packer plugin list which we just created and
  source the file and finally run <code>:PackerInstall</code>. This should install
  the plugin. You can then set the colorscheme as per your preference, first view
  the color scheme temporarily on the current instance with the command <code>:colorscheme
  colorscheme_name</code>.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kr\">return</span> <span class=\"nb\">require</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;packer&#39;</span><span class=\"p\">).</span><span class=\"n\">startup</span><span
  class=\"p\">(</span><span class=\"kr\">function</span><span class=\"p\">()</span>\n
  \ <span class=\"n\">use</span> <span class=\"s1\">&#39;wbthomason/packer.nvim&#39;</span>\n
  \ <span class=\"c1\">-- </span>\n  <span class=\"n\">use</span> <span class=\"s1\">&#39;Mofiqul/dracula.nvim&#39;</span>\n
  \ <span class=\"c1\">--</span>\n<span class=\"kr\">end</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>You
  can then add the command to set it as the default color scheme for Neovim.</p>\n<pre
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">cmd</span> <span
  class=\"s\">[[silent! colorscheme dracula]]</span>\n</pre></div>\n\n</pre>\n\n<p>You
  can change the background color, text color, icons style and terminal and gui style
  separately with the augroup and setting the colorscheme commands.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">api</span><span
  class=\"p\">.</span><span class=\"n\">nvim_command</span><span class=\"p\">(</span><span
  class=\"s\">[[</span>\n<span class=\"s\">    augroup ChangeBackgroudColour</span>\n<span
  class=\"s\">        autocmd colorscheme * :hi normal termbg=#000030 termfg=#ffffff</span>\n<span
  class=\"s\">        autocmd colorscheme * :hi Directory ctermfg=#ffffff</span>\n<span
  class=\"s\">    augroup END</span>\n<span class=\"s\">]]</span><span class=\"p\">)</span>\n<span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">o</span><span
  class=\"p\">.</span><span class=\"n\">termguicolors</span> <span class=\"o\">=</span>
  <span class=\"kc\">true</span>\n</pre></div>\n\n</pre>\n\n<p>Here, I have used the
  background and foreground colors of the terminal variant of Neovim. Also for the
  Directory Explorer i.e. netrw, I have changed the terminal foreground. This you
  can configure as per your needs, Though this is still vimscripty, there are Autocommands
  and autogroups available sooner in Neovim.</p>\n<h2>Separating Configurations</h2>\n<p>If
  you wish to keep all your config in one file i.e. <code>init.lua</code> file, you
  can, though you can separate out things like <code>keymaps</code>, <code>plugins</code>,
  <code>custom_options</code> or if you have <code>lsp</code> configurations into
  separate lua packages or creating a separate module. This helps in loading up things
  as per requirement and also looks readable, making it a lot easier to test out things
  without breaking the whole config.</p>\n<p>Definitely, there will be personal preferences
  and pros and cons of both approaches, you can pick up whatever suits your style.</p>\n<h3>Creating
  separate packages</h3>\n<p>To create a separate package, we can simply add a file
  in the same folder as <code>init.vim</code> i.e. in the folder <code>~/.config/nvim</code>
  or equivalent for windows. The package name can be any valid filename with the <code>lua</code>
  extension.</p>\n<p>For instance, you can create a package for all your keymaps and
  load it in the <code>init.lua</code> as per the order you want to load them. It
  can be at the top, or else if you have certain base settings lower in the init file,
  these might not reflect or load up in the keymap package so better to load them
  after some of the core settings have been set.</p>\n<p>Let's create the package
  and dump all our maps into the keymap file package.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">-- ~/.config/nvim/keymap.lua</span>\n\n<span class=\"kr\">function</span>
  <span class=\"nf\">map</span><span class=\"p\">(</span><span class=\"n\">mode</span><span
  class=\"p\">,</span> <span class=\"n\">lhs</span><span class=\"p\">,</span> <span
  class=\"n\">rhs</span><span class=\"p\">,</span> <span class=\"n\">opts</span><span
  class=\"p\">)</span>\n    <span class=\"kd\">local</span> <span class=\"n\">options</span>
  <span class=\"o\">=</span> <span class=\"p\">{</span> <span class=\"n\">noremap</span>
  <span class=\"o\">=</span> <span class=\"kc\">true</span> <span class=\"p\">}</span>\n
  \   <span class=\"kr\">if</span> <span class=\"n\">opts</span> <span class=\"kr\">then</span>\n
  \       <span class=\"n\">options</span> <span class=\"o\">=</span> <span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">tbl_extend</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;force&quot;</span><span class=\"p\">,</span> <span class=\"n\">options</span><span
  class=\"p\">,</span> <span class=\"n\">opts</span><span class=\"p\">)</span>\n    <span
  class=\"kr\">end</span>\n    <span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">api</span><span class=\"p\">.</span><span class=\"n\">nvim_set_keymap</span><span
  class=\"p\">(</span><span class=\"n\">mode</span><span class=\"p\">,</span> <span
  class=\"n\">lhs</span><span class=\"p\">,</span> <span class=\"n\">rhs</span><span
  class=\"p\">,</span> <span class=\"n\">options</span><span class=\"p\">)</span>\n<span
  class=\"kr\">end</span>\n\n<span class=\"n\">map</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;w&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;:w&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;q&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;:q!&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;s&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;:so %&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;cpp&#39;</span> <span class=\"s1\">&#39;:!c++
  % -o %:r &amp;&amp; %:r&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span class=\"n\">map</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;c,&#39;</span> <span class=\"s1\">&#39;:!gcc % -o %:r &amp;&amp;
  %:r&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span class=\"n\">map</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;py&#39;</span> <span class=\"s1\">&#39;:!python %&lt;cr&gt;&#39;</span><span
  class=\"p\">)</span>\n<span class=\"n\">map</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;go&#39;</span>
  <span class=\"s1\">&#39;:!go run %&lt;cr&gt;&#39;</span><span class=\"p\">)</span>\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;sh&#39;</span> <span class=\"s1\">&#39;:!bash
  %&lt;cr&gt;&#39;</span><span class=\"p\">)</span>\n\n<span class=\"c1\">-- more
  keymaps</span>\n</pre></div>\n\n</pre>\n\n<p>So, this might work if you don't have
  any plugin-related keymaps as it would require those functions or objects available
  to execute. So, we might also want to separate plugins and load them first in our
  keymaps or in the init file.</p>\n<p>Now, there needs to be a way for grabbing a
  package. Yes, there is basically like import in python or any other programming
  language, lua has <code>require</code> keyword for importing packages. Since the
  <code>init</code> file and the <code>keymaps</code> are in the same folder path,
  we can simply say, <code>require(&quot;keymap&quot;)</code> in our <code>init.lua</code>
  file. Now, it depends on your config where you want to load this package. At the
  top i.e. at the beginning of neovim instance or after loading the plugins down.</p>\n<pre
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">-- init.lua</span>\n\n<span class=\"nb\">require</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;keymaps&quot;</span><span class=\"p\">)</span>\n\n<span class=\"c1\">--
  At the top</span>\n<span class=\"c1\">-- OR</span>\n<span class=\"c1\">-- After
  loading Packer plugins</span>\n</pre></div>\n\n</pre>\n\n<p>So, now you can separate
  all your configs as per your requirement. It is like splitting up a puzzle and again
  combining them. Similar package can be created for <code>plugins</code>, <code>options</code>
  or <code>lsp</code> configurations.</p>\n<h3>Creating a separate module</h3>\n<p>Now,
  we have seen how to create a lua package and loading in neovim. We also can create
  modules in neovim configuration. For instance, first, the init file is loaded, Other
  files might not be required hence it is not loaded by default, it is only loaded
  when <code>require</code>ed. So, we can create a module in lua with a folder, and
  inside of it, we can have packages as we had in the previous method. Also, every
  module can have a init file loaded first when we require that module. The rest of
  the packages in the module are thus made available thereafter.</p>\n<ul>\n<li>Module
  is a component not loaded by default</li>\n<li>Only loaded when required (literally
  require)</li>\n<li>Every module can have a <code>init.lua</code> file loaded first
  when required.</li>\n<li>Require a package in module from outside -&gt; <code>require('module_name.package_name')</code></li>\n</ul>\n<p>So,
  in neovim, we need to create a <code>lua</code> folder for placing all our modules
  so that the lua runtime is picked up correctly. Inside this lua folder, we can create
  a module basically a folder. This folder name can be anything you like, I like to
  keep it my nickname, you can use whatever you prefer.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span># ~/.config/nvim\n\n-- init.lua\n--
  lua/\n    -- module_name/\n        -- init.lua\n        -- package_name.lua\n        --
  keymaps.lua\n</pre></div>\n\n</pre>\n\n<p>Now, we can create packages in this module.
  You can move your keymaps package inside this folder. The keymaps package is nothing
  here but an example provided in the previous section for creating a package. You
  can basically dump all your keymaps in a single lua file and then import it from
  the init file. Similarly you can create a package inside a module and import it
  from the init file of the module(local init file <code>~/.config/nvim/lua/module_name/init.lua</code>)
  or the global init file(<code>~/.config/nvim/init.lua</code>).</p>\n<p>To load the
  packages, you have to use the same require statement irrespective of where you are
  importing from i.e. either from the module or from the global init file. The require
  statement would look like the following <code>require(&quot;module_name/package_name&quot;)</code>.
  Now, we can use the keymaps package to import from the module init file and then
  import the module from the global init file. To import a module, we can simply use
  the module name in the require string as <code>require(&quot;module_name&quot;)</code>.</p>\n<pre
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">-- ~/.config/nvim</span>\n\n<span class=\"c1\">-- lua/module_name/options.lua</span>\n\n<span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">opt</span><span
  class=\"p\">.</span><span class=\"n\">number</span> <span class=\"o\">=</span> <span
  class=\"kc\">true</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">opt</span><span class=\"p\">.</span><span class=\"n\">tabstop</span>
  <span class=\"o\">=</span> <span class=\"mi\">4</span>\n<span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">opt</span><span class=\"p\">.</span><span
  class=\"n\">swapfile</span> <span class=\"o\">=</span> <span class=\"kc\">false</span>\n\n\n<span
  class=\"c1\">-- lua/module_name/plugins.lua</span>\n\n<span class=\"nb\">require</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;module_name.options&quot;</span><span
  class=\"p\">)</span>\n<span class=\"kr\">return</span> <span class=\"nb\">require</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;packer&#39;</span><span class=\"p\">).</span><span
  class=\"n\">startup</span><span class=\"p\">(</span><span class=\"kr\">function</span><span
  class=\"p\">()</span>\n  <span class=\"n\">use</span> <span class=\"s1\">&#39;wbthomason/packer.nvim&#39;</span>\n
  \ <span class=\"c1\">--plugins</span>\n<span class=\"kr\">end</span><span class=\"p\">)</span>\n\n<span
  class=\"c1\">-- lua/module_name/keymap.lua</span>\n\n<span class=\"nb\">require</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;module_name.plugins&quot;</span><span
  class=\"p\">)</span>\n<span class=\"c1\">-- maps()</span>\n\n\n<span class=\"c1\">--
  lua/module_name/init.lua</span>\n\n<span class=\"nb\">require</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;module_name.keymaps)</span>\n\n\n<span class=\"s2\">-- init.lua</span>\n\n<span
  class=\"s2\">require(&quot;</span><span class=\"n\">module_name</span><span class=\"s2\">&quot;)</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  this is how we can create modules and packages for configurations in neovim using
  lua. This is also a kind of a structure for creating your own neovim plugin with
  lua!</p>\n<p>For further references, you can check out my <a href=\"https://github.com/Mr-Destructive/dotfiles\">dotfiles</a>.</p>\n<h3>References</h3>\n<ul>\n<li><a
  href=\"https://vonheikemen.github.io/devlog/tools/configuring-neovim-using-lua/\">Configure
  Neovim for Lua</a></li>\n<li><a href=\"https://alpha2phi.medium.com/neovim-for-beginners-init-lua-45ff91f741cb\">Neovim
  with Lua for beginners</a></li>\n<li><a href=\"https://www.youtube.com/c/TJDeVries/videos\">TJ
  Devries Youtube</a></li>\n</ul>\n<h2>Conclusion</h2>\n<p>So, that is just a basic
  overview of how you can get your neovim configured for lua. It is a great experience
  to just create such a personalized environment and play with it and have fun. You
  might have hopefully configured Neovim for Lua from this little guide. Maybe in
  the next article, I'll set up LSP in Neovim. If you have any queries or feedback,
  please let me know. Thank you for reading.</p>\n<p>Happy Viming :)</p>\n"
cover: ''
date: 2022-07-09
datetime: 2022-07-09 00:00:00+00:00
description: Setting up or Migrating to Lua in Neovim Configuration. Using Lua in
  Neovim for configurations and moving out of Vimscript. Setting up Keymaps, Plugins,
  customizations
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2022-07-07-Neovim-Lua.md
html: "<h2>Introduction</h2>\n<p>It has been a while since I have written a Vim article.
  Finally, I got some ideas after configuring my Neovim setup for Lua. I recently
  migrated to Ubuntu a couple of months back and it has been a cool change from Windows
  7!</p>\n<p>In this article, we'll see how you can set up neovim for Lua. Since Neovim
  0.5, it supports lua out of the box, so in the recent release 0.7, it added more
  native support to lua making it a lot easier to configure and play with neovim.
  So, we will see how we can use lua to convert all the 200 liner vimscript into lua
  (We can even have packages and modules:) We will cover how to configure your keymaps,
  pull up all the plugins, vim options, and other customizations.</p>\n<h2>Why move
  to Lua?</h2>\n<p>I have used Vimscript for quite a while now, configured it as per
  my needs, and also made a few plugins like <a href=\"https://github.com/Mr-Destructive/frontmatter.vim\">frontmatter</a>,
  <a href=\"https://github.com/Mr-Destructive/dj.vim\">dj.vim</a>, and <a href=\"https://github.com/Mr-Destructive/commenter.vim\">commenter</a>
  which are quite clunky and not robust in terms of usage and customizability. Vimscript
  is good but it's a bit messy when you want extreme customization.</p>\n<p>I recently
  wanted to go full Neovim, I was kind of stuck in Vimscript and some of my plugins
  work for me but it might not work for others, it might be a piece of gibberish vimscript
  dumped. So, Why not have full native experience in Neovim if you can. It has now
  baked-in support for LSP, Debugging, Autocommands, and so much more. If you have
  Neovim 0.5+ you should be good to go full lua.</p>\n<h2>Backup Neovim Config</h2>\n<p>Firstly,
  keep your original neovim/vim init files safe, back them up, make a copy and save
  it with a different name like <code>nvim_config.vim</code>. If you already have
  all your config files backed up as an ansible script or dotfiles GitHub repository,
  you can proceed ahead.</p>\n<p>But don't keep the <code>init.vim</code> file as
  it is in the <code>~/.config/nvim</code> folder, it might override after we configure
  the lua scripts.</p>\n<h2>Basic Configuration</h2>\n<p>So, I assume you have set
  up Neovim, If not you need to follow some simple steps like downloading the package
  and making sure your neovim environment is working with vimscript first. The <a
  href=\"https://github.com/neovim/neovim/wiki/Installing-Neovim\">Neovim Wiki</a>
  provides great documentation on how to install neovim on various platforms using
  different methods.</p>\n<p>After your neovim is set up and you have a basic configuration,
  you can now start to migrate into lua.\nCreate a <code>init.lua</code> file in the
  same path as your <code>init.lua</code> file is i.e. at <code>~/.config/nvim</code>
  or <code>~/AppData/Local/nvim/</code> for Windows. That's why it is recommended
  to keep the initial configuration vimscript file in a safe place. While migrating
  from vimscript to lua, once the lua file is created and the next time you restart
  neovim, the default settings will be from <code>init.lua</code> and not <code>init.vim</code>,
  so be prepared.</p>\n<p>Firstly, you need to configure some options like <code>number</code>,
  <code>syntax highlighting</code>, <code>tabs</code>, and some <code>keymaps</code>
  of course. We can use the <code>vim.opt</code> method to set options in vim using
  lua syntax. So, certain corresponding vim options would be converted as follows:</p>\n<p>If
  you have the following kind of settings in your vimrc or init.vim:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>vimscript</div>\n<div class=\"highlight\"><pre><span></span>--
  vimscript\nset number\nset tabstop=4 \nset shiftwidth=4 \nset softtabstop=0 \nset
  expandtab \nset noswapfile\n</pre></div>\n\n</pre>\n\n<p>The above settings are
  migrated into lua syntax as follows:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">--lua</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">opt</span><span class=\"p\">.</span><span class=\"n\">number</span>
  <span class=\"o\">=</span> <span class=\"kc\">true</span>\n<span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">opt</span><span class=\"p\">.</span><span
  class=\"n\">tabstop</span> <span class=\"o\">=</span> <span class=\"mi\">4</span>\n<span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">opt</span><span
  class=\"p\">.</span><span class=\"n\">shiftwidth</span> <span class=\"o\">=</span>
  <span class=\"mi\">4</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">opt</span><span class=\"p\">.</span><span class=\"n\">softtabstop</span>
  <span class=\"o\">=</span> <span class=\"mi\">0</span>\n<span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">opt</span><span class=\"p\">.</span><span
  class=\"n\">expandtab</span> <span class=\"o\">=</span> <span class=\"kc\">true</span>\n<span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">opt</span><span
  class=\"p\">.</span><span class=\"n\">swapfile</span> <span class=\"o\">=</span>
  <span class=\"kc\">false</span>\n</pre></div>\n\n</pre>\n\n<p>You can set other
  options in your config file accordingly. If you get sick of writing <code>vim.opt.</code>
  again and again, you can use a variable set to <code>vim.opt</code> and then access
  that variable to set the options. Something of the lines as below:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kd\">local</span> <span class=\"n\">set</span> <span class=\"o\">=</span>
  <span class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">opt</span>\n\n<span
  class=\"n\">set</span><span class=\"p\">.</span><span class=\"n\">number</span>
  <span class=\"o\">=</span> <span class=\"kc\">true</span>\n<span class=\"n\">set</span><span
  class=\"p\">.</span><span class=\"n\">tabstop</span> <span class=\"o\">=</span>
  <span class=\"mi\">4</span>\n<span class=\"n\">set</span><span class=\"p\">.</span><span
  class=\"n\">shiftwidth</span> <span class=\"o\">=</span> <span class=\"mi\">4</span>\n<span
  class=\"n\">set</span><span class=\"p\">.</span><span class=\"n\">softtabstop</span>
  <span class=\"o\">=</span> <span class=\"mi\">0</span>\n<span class=\"n\">set</span><span
  class=\"p\">.</span><span class=\"n\">expandtab</span> <span class=\"o\">=</span>
  <span class=\"kc\">true</span>\n<span class=\"n\">set</span><span class=\"p\">.</span><span
  class=\"n\">swapfile</span> <span class=\"o\">=</span> <span class=\"kc\">false</span>\n</pre></div>\n\n</pre>\n\n<p>We
  can create a variable in lua like <code>local variable_name = something</code> so,
  we have created a variable <code>set</code> which is assigned to the value of <code>vim.opt</code>
  which is a method(function) in lua to set the options from the vimscript environment.
  Finally, access the <code>set</code> keyword to set the options. Using the <code>set</code>
  word is not necessary, you can use whatever you want.</p>\n<p>After setting up the
  basic options, you can source the file with <code>:so %</code> from the command
  mode. Just normally as you source the config files.</p>\n<h3>Using Lua in Command
  Mode</h3>\n<p>We can use the lua functions or any other commands from the command
  mode in neovim using the lua command. Just prefix the command with <code>:lua</code>
  and after that, you can use lua syntax like function calling or setting variables,
  logging things, etc.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1657380885/blog-media/lua_in_command_mode.gif\"
  alt=\"Lua in Command Mode\" /></p>\n<h2>Adding Keymaps</h2>\n<p>Now, that we have
  some basic config setup, we can quickly get the keymaps. It's not that hard to make
  keymaps to set up in lua. To create keymaps, we have two options. One is compatible
  with Neovim and the other is compatible with Vim as well.</p>\n<ol>\n<li>vim.keymap.set
  OR</li>\n<li>vim.api.nvim_set_keymap</li>\n</ol>\n<p>Personally, I don't see a difference
  in terms of usage but <a href=\"https://github.com/neovim/neovim/blob/master/runtime/lua/vim/keymap.lua#L51\">vim.keymap.set</a>
  is a wrapper around <a href=\"https://github.com/neovim/neovim/blob/master/src/nvim/api/vim.c#L1451\">nvim_set_keymap</a>.
  So, it is really a matter of personal preference which you want to use.</p>\n<p>So,
  both the functions have quite similar parameters:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">keymap</span><span
  class=\"p\">.</span><span class=\"n\">set</span><span class=\"p\">({</span><span
  class=\"n\">mode</span><span class=\"p\">},</span> <span class=\"p\">{</span><span
  class=\"n\">lhs</span><span class=\"p\">},</span> <span class=\"p\">{</span><span
  class=\"n\">rhs</span><span class=\"p\">},</span> <span class=\"p\">{</span><span
  class=\"n\">options</span><span class=\"p\">})</span>\n\n<span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">api</span><span class=\"p\">.</span><span
  class=\"n\">nvim_set_keymap</span><span class=\"p\">({</span><span class=\"n\">mode</span><span
  class=\"p\">},</span> <span class=\"p\">{</span><span class=\"n\">lhs</span><span
  class=\"p\">},</span> <span class=\"p\">{</span><span class=\"n\">rhs</span><span
  class=\"p\">},</span> <span class=\"p\">{</span><span class=\"n\">options</span><span
  class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p>The advantage of <code>vim.keymap.set</code>
  over <code>vim.api.nvim_set_keymap</code> is that it allows directly calling lua
  functions rather than calling vimscripty way like <code>:lua function()</code>,
  so we directly can use lua code in the RHS part of the function parameter.</p>\n<p>Let's
  take a basic example mapping:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n\n<div class=\"highlight\"><pre><span></span>vim.keymap.set(&#39;n&#39;,
  &#39;Y&#39;, &#39;yy&#39;, {noremap = false})\n</pre></div>\n\n</pre>\n\n<p>Here,
  we have mapped <code>Shift+y</code> with the keys <code>yy</code> in <code>n</code>ormal
  mode. The first argument is the mode, it can be a single-mode like <code>'n'</code>,
  <code>'v'</code>, <code>'i'</code>, etc., or a multi-mode table like <code>{'n',
  'v'}</code>, <code>{'v', 'i'}</code>, etc.</p>\n<p>The next parameter is also a
  string but it should be the key for triggering the mapping. In this case, we have
  used <code>Y</code> which is <code>Shift + y</code>, it can be any key shortcut
  you want to map.</p>\n<p>The third parameter is the string which will be the command
  to be executed after the key has been used. Here we have used the keys <code>yy</code>,
  if the map is from a command mode, you will be using something like <code>':commands_to_be
  executed'</code> as the third parameter.</p>\n<p>The fourth parameter which is optional
  can contain <a href=\"https://neovim.io/doc/user/api.html#:~:text=nvim_set_keymap(%7Bmode%7D%2C%20%7Blhs%7D%2C%20%7Brhs%7D%2C%20%7B*opts%7D)%20%20%20%20%20%20%20%20%20%20%20%20%20*nvim_set_keymap()*\">special
  arguments</a>. We have set a default option which is <code>noremap</code> as true,
  the options are not string but lua tables instead, so it can be similar to python
  dictionary or a map kind of a structure with a key value pair.</p>\n<p>One more
  important aspect in keymapping might about the leader key, you can set your leader
  key by using the global vim options with <code>vim.g</code> and access <code>mapleader</code>
  to set it to the key you wish. This will make the <code>leader</code> key available
  to us and thereafter we can map the leader key in custom mappings.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>vim.g.mapleader = &quot;
  &quot;\n</pre></div>\n\n</pre>\n\n<p>Here, I have set my leader key to the <code>&lt;Space&gt;</code>
  key. Now, we can map keys to the existing keymaps in the vimscript. Let's map some
  basic keymaps first and then after setting up the plugins,we can move into plugin-specific
  mappings.</p>\n<p>You can also use <code>vim.api.nvim_set_keymap</code> function
  with the same parameters as well.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">keymap</span><span
  class=\"p\">.</span><span class=\"n\">set</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;w&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;:w&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">keymap</span><span class=\"p\">.</span><span class=\"n\">set</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;&lt;leader&gt;q&#39;</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;:q!&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span class=\"n\">noremap</span>
  <span class=\"o\">=</span> <span class=\"kc\">true</span><span class=\"p\">})</span>\n<span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">keymap</span><span
  class=\"p\">.</span><span class=\"n\">set</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;s&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;:so %&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">keymap</span><span class=\"p\">.</span><span class=\"n\">set</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;&lt;leader&gt;ev&#39;</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;:vsplit $MYVIMRC&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">keymap</span><span class=\"p\">.</span><span class=\"n\">set</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;&lt;leader&gt;sv&#39;</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;:w&lt;CR&gt;:so %&lt;CR&gt;:q&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n\n<span class=\"n\">OR</span>\n\n<span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">api</span><span class=\"p\">.</span><span
  class=\"n\">nvim_set_keymap</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;w&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;:w&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">api</span><span class=\"p\">.</span><span class=\"n\">nvim_set_keymap</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;&lt;leader&gt;q&#39;</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;:q!&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span class=\"n\">noremap</span>
  <span class=\"o\">=</span> <span class=\"kc\">true</span><span class=\"p\">})</span>\n<span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">api</span><span
  class=\"p\">.</span><span class=\"n\">nvim_set_keymap</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;s&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;:so %&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">api</span><span class=\"p\">.</span><span class=\"n\">nvim_set_keymap</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;&lt;leader&gt;ev&#39;</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;:vsplit $MYVIMRC&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">api</span><span class=\"p\">.</span><span class=\"n\">nvim_set_keymap</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;&lt;leader&gt;sv&#39;</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;:w&lt;CR&gt;:so %&lt;CR&gt;:q&lt;CR&gt;&#39;</span><span class=\"p\">,{</span><span
  class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span><span
  class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p>If, you don't like writing <code>vim.keymap.set</code>
  or <code>vim.api.nvim_set_keymap</code> again and again, you can create a simpler
  function for it. In lua a function is created just like a variable by specifying
  the scope of the function i.e. local followed by the <code>function</code> keyword
  and finally the name of the function and parenthesis. The function body is terminated
  by the <code>end</code> keyword.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kr\">function</span> <span class=\"nf\">map</span><span class=\"p\">(</span><span
  class=\"n\">mode</span><span class=\"p\">,</span> <span class=\"n\">lhs</span><span
  class=\"p\">,</span> <span class=\"n\">rhs</span><span class=\"p\">,</span> <span
  class=\"n\">opts</span><span class=\"p\">)</span>\n    <span class=\"kd\">local</span>
  <span class=\"n\">options</span> <span class=\"o\">=</span> <span class=\"p\">{</span>
  <span class=\"n\">noremap</span> <span class=\"o\">=</span> <span class=\"kc\">true</span>
  <span class=\"p\">}</span>\n    <span class=\"kr\">if</span> <span class=\"n\">opts</span>
  <span class=\"kr\">then</span>\n        <span class=\"n\">options</span> <span class=\"o\">=</span>
  <span class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">tbl_extend</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;force&quot;</span><span class=\"p\">,</span>
  <span class=\"n\">options</span><span class=\"p\">,</span> <span class=\"n\">opts</span><span
  class=\"p\">)</span>\n    <span class=\"kr\">end</span>\n    <span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">api</span><span class=\"p\">.</span><span
  class=\"n\">nvim_set_keymap</span><span class=\"p\">(</span><span class=\"n\">mode</span><span
  class=\"p\">,</span> <span class=\"n\">lhs</span><span class=\"p\">,</span> <span
  class=\"n\">rhs</span><span class=\"p\">,</span> <span class=\"n\">options</span><span
  class=\"p\">)</span>\n<span class=\"kr\">end</span>\n</pre></div>\n\n</pre>\n\n<p>Now,
  in this function <code>map</code>, we have passed in the same parameters like the
  <code>vim.keymap.set</code> function takes but we have just parsed the function
  in a shorter and safer way by adding <code>noremap = true</code> by default. So
  this is just a helper function or a verbose function for calling the vim.keymap.set
  function.</p>\n<p>To use this function, we can simply call <code>map</code> with
  the same arguments as given to the prior functions.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;w&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;:w&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;q&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;:q!&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;s&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;:so %&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>Notice
  here, though, we have not passed the <code>{noremap = true}</code> as it will be
  passed by default to the <code>vim.api.nvim_set_keymap</code> or <code>vim.keymap.set</code>
  function via the custom map function.</p>\n<p>If you want some more examples, here
  are some additional mapping specific to languages, meant for compiling or running
  scripts with neovim instance.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">-- vimscript</span>\n\n<span class=\"n\">nnoremap</span> <span class=\"n\">cpp</span>
  <span class=\"p\">:</span><span class=\"err\">!</span><span class=\"n\">c</span><span
  class=\"o\">++</span> <span class=\"o\">%</span> <span class=\"o\">-</span><span
  class=\"n\">o</span> <span class=\"o\">%</span><span class=\"p\">:</span><span class=\"n\">r</span>
  <span class=\"o\">&amp;&amp;</span> <span class=\"o\">%</span><span class=\"p\">:</span><span
  class=\"n\">r</span><span class=\"o\">&lt;</span><span class=\"n\">CR</span><span
  class=\"o\">&gt;</span>\n<span class=\"n\">nnoremap</span> <span class=\"n\">c</span><span
  class=\"p\">,</span> <span class=\"p\">:</span><span class=\"err\">!</span><span
  class=\"n\">gcc</span> <span class=\"o\">%</span> <span class=\"o\">-</span><span
  class=\"n\">o</span> <span class=\"o\">%</span><span class=\"p\">:</span><span class=\"n\">r</span>
  <span class=\"o\">&amp;&amp;</span> <span class=\"o\">%</span><span class=\"p\">:</span><span
  class=\"n\">r</span><span class=\"o\">&lt;</span><span class=\"n\">CR</span><span
  class=\"o\">&gt;</span>\n<span class=\"n\">nnoremap</span> <span class=\"n\">py</span>
  <span class=\"p\">:</span><span class=\"err\">!</span><span class=\"n\">python</span>
  <span class=\"o\">%&lt;</span><span class=\"n\">cr</span><span class=\"o\">&gt;</span>\n<span
  class=\"n\">nnoremap</span> <span class=\"n\">go</span> <span class=\"p\">:</span><span
  class=\"err\">!</span><span class=\"n\">go</span> <span class=\"n\">run</span> <span
  class=\"o\">%&lt;</span><span class=\"n\">cr</span><span class=\"o\">&gt;</span>\n<span
  class=\"n\">nnoremap</span> <span class=\"n\">sh</span> <span class=\"p\">:</span><span
  class=\"err\">!</span><span class=\"n\">bash</span> <span class=\"o\">%&lt;</span><span
  class=\"n\">CR</span><span class=\"o\">&gt;</span>\n\n\n<span class=\"c1\">-- lua</span>\n\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;cpp&#39;</span> <span class=\"s1\">&#39;:!c++
  % -o %:r &amp;&amp; %:r&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span class=\"n\">map</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;c,&#39;</span> <span class=\"s1\">&#39;:!gcc % -o %:r &amp;&amp;
  %:r&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span class=\"n\">map</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;py&#39;</span> <span class=\"s1\">&#39;:!python %&lt;cr&gt;&#39;</span><span
  class=\"p\">)</span>\n<span class=\"n\">map</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;go&#39;</span>
  <span class=\"s1\">&#39;:!go run %&lt;cr&gt;&#39;</span><span class=\"p\">)</span>\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;sh&#39;</span> <span class=\"s1\">&#39;:!bash
  %&lt;cr&gt;&#39;</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  this is how we can set up our keymaps in lua. You can customize this function as
  per your needs. These are just made for the understanding purpose of how to reduce
  the repetitive stuff in the setup.</p>\n<p><strong>If you are really stuck up and
  not feeling to convert those mappings into lua then I have a function that can do
  it for you, check out my dotfile repo -&gt; <a href=\"https://github.com/Mr-Destructive/dotfiles/blob/master/nvim/lua/destructive/options.lua#L9\">keymapper</a></strong></p>\n<h2>Adding
  Plugin Manager</h2>\n<p>Now, we really missing some plugins, aren't we? So, the
  neovim community has some good choices for using a new plugin manager written in
  core lua. It is usually a good idea to move into lua completely rather than switching
  to and fro between vimscript and lua.</p>\n<p>So, <a href=\"https://github.com/wbthomason/packer.nvim\">Packer</a>
  is the new plugin manager for Neovim in Lua, there is other plugin managers out
  there as well like <a href=\"https://github.com/savq/paq-nvim\">paq</a>. If you
  don't want to switch with the plugin manager, you can still use vim-based plugin
  managers like <a href=\"https://dev.to/vonheikemen/neovim-using-vim-plug-in-lua-3oom\">Vim-Plug</a>.</p>\n<p>So,
  let's install the Packer plugin manager into Neovim. We simply have to run the following
  command in the console and make sure the plugin manager is configured correctly.</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span># Linux\n\ngit clone --depth
  1 https://github.com/wbthomason/packer.nvim\\\n~/.local/share/nvim/site/pack/packer/start/packer.nvim\n\n\n#
  Windows\n\ngit clone https://github.com/wbthomason/packer.nvim &quot;$env:LOCALAPPDATA\\nvim-data\\site\\pack\\packer\\start\\packer.nvim&quot;\n</pre></div>\n\n</pre>\n\n<p>Now,
  if you open a new neovim instance and try to run the command <code>:PackerClean</code>,
  and no error pops out that means you have configured it correctly. You can move
  ahead to installing plugins now. Yeah! PLUG-IN time!</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kr\">return</span> <span class=\"nb\">require</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;packer&#39;</span><span class=\"p\">).</span><span class=\"n\">startup</span><span
  class=\"p\">(</span><span class=\"kr\">function</span><span class=\"p\">()</span>\n<span
  class=\"kr\">end</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>First
  try to source the file, if it throws out errors it shouldn't try to fix the installation
  path of Packer. If the command succeded we can finally pull up some plugins.</p>\n<p>Below
  are some of the plugins that you can use irrespective of what language preferences
  you would have. This includes basic dev-icons for the status line as well as the
  explorer window file icons. As usual, add your plugins and make them yours.</p>\n<pre
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kr\">return</span> <span class=\"nb\">require</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;packer&#39;</span><span class=\"p\">).</span><span class=\"n\">startup</span><span
  class=\"p\">(</span><span class=\"kr\">function</span><span class=\"p\">()</span>\n
  \ <span class=\"n\">use</span> <span class=\"s1\">&#39;wbthomason/packer.nvim&#39;</span>\n
  \ <span class=\"n\">use</span> <span class=\"s1\">&#39;tpope/vim-fugitive&#39;</span>\n
  \ <span class=\"n\">use</span> <span class=\"p\">{</span>\n    <span class=\"s1\">&#39;nvim-lualine/lualine.nvim&#39;</span><span
  class=\"p\">,</span>\n    <span class=\"n\">requires</span> <span class=\"o\">=</span>
  <span class=\"p\">{</span> <span class=\"s1\">&#39;kyazdani42/nvim-web-devicons&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">opt</span> <span class=\"o\">=</span> <span
  class=\"kc\">true</span> <span class=\"p\">}</span>\n  <span class=\"p\">}</span>\n
  \ <span class=\"n\">use</span> <span class=\"s1\">&#39;tiagofumo/vim-nerdtree-syntax-highlight&#39;</span>\n
  \ <span class=\"n\">use</span> <span class=\"s1\">&#39;kyazdani42/nvim-web-devicons&#39;</span>\n
  \ <span class=\"n\">use</span> <span class=\"s1\">&#39;vim-airline/vim-airline&#39;</span>\n
  \ <span class=\"n\">use</span> <span class=\"s1\">&#39;vim-airline/vim-airline-themes&#39;</span>\n<span
  class=\"kr\">end</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>After
  adding the list of your plugins, you need to source the file and then install the
  plugins with the command <code>:PackerInstall</code>. This function will install
  all the plugins after the file has been sourced so make sure you don't forget it.</p>\n<h2>Colors
  and Color Themes</h2>\n<p>You might fancy some good-looking and aesthetic setup
  for neovim of course! In Neovim, we already have a wide variety of configurations
  to set up like color schemes, GUI colors, terminal colors, etc. You can pick up
  a color scheme from a large list of awesome color schemes from <a href=\"https://github.com/topics/neovim-colorscheme\">GitHub</a>.</p>\n<p>After
  choosing the theme, plug it in the packer plugin list which we just created and
  source the file and finally run <code>:PackerInstall</code>. This should install
  the plugin. You can then set the colorscheme as per your preference, first view
  the color scheme temporarily on the current instance with the command <code>:colorscheme
  colorscheme_name</code>.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kr\">return</span> <span class=\"nb\">require</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;packer&#39;</span><span class=\"p\">).</span><span class=\"n\">startup</span><span
  class=\"p\">(</span><span class=\"kr\">function</span><span class=\"p\">()</span>\n
  \ <span class=\"n\">use</span> <span class=\"s1\">&#39;wbthomason/packer.nvim&#39;</span>\n
  \ <span class=\"c1\">-- </span>\n  <span class=\"n\">use</span> <span class=\"s1\">&#39;Mofiqul/dracula.nvim&#39;</span>\n
  \ <span class=\"c1\">--</span>\n<span class=\"kr\">end</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>You
  can then add the command to set it as the default color scheme for Neovim.</p>\n<pre
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">cmd</span> <span
  class=\"s\">[[silent! colorscheme dracula]]</span>\n</pre></div>\n\n</pre>\n\n<p>You
  can change the background color, text color, icons style and terminal and gui style
  separately with the augroup and setting the colorscheme commands.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">api</span><span
  class=\"p\">.</span><span class=\"n\">nvim_command</span><span class=\"p\">(</span><span
  class=\"s\">[[</span>\n<span class=\"s\">    augroup ChangeBackgroudColour</span>\n<span
  class=\"s\">        autocmd colorscheme * :hi normal termbg=#000030 termfg=#ffffff</span>\n<span
  class=\"s\">        autocmd colorscheme * :hi Directory ctermfg=#ffffff</span>\n<span
  class=\"s\">    augroup END</span>\n<span class=\"s\">]]</span><span class=\"p\">)</span>\n<span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">o</span><span
  class=\"p\">.</span><span class=\"n\">termguicolors</span> <span class=\"o\">=</span>
  <span class=\"kc\">true</span>\n</pre></div>\n\n</pre>\n\n<p>Here, I have used the
  background and foreground colors of the terminal variant of Neovim. Also for the
  Directory Explorer i.e. netrw, I have changed the terminal foreground. This you
  can configure as per your needs, Though this is still vimscripty, there are Autocommands
  and autogroups available sooner in Neovim.</p>\n<h2>Separating Configurations</h2>\n<p>If
  you wish to keep all your config in one file i.e. <code>init.lua</code> file, you
  can, though you can separate out things like <code>keymaps</code>, <code>plugins</code>,
  <code>custom_options</code> or if you have <code>lsp</code> configurations into
  separate lua packages or creating a separate module. This helps in loading up things
  as per requirement and also looks readable, making it a lot easier to test out things
  without breaking the whole config.</p>\n<p>Definitely, there will be personal preferences
  and pros and cons of both approaches, you can pick up whatever suits your style.</p>\n<h3>Creating
  separate packages</h3>\n<p>To create a separate package, we can simply add a file
  in the same folder as <code>init.vim</code> i.e. in the folder <code>~/.config/nvim</code>
  or equivalent for windows. The package name can be any valid filename with the <code>lua</code>
  extension.</p>\n<p>For instance, you can create a package for all your keymaps and
  load it in the <code>init.lua</code> as per the order you want to load them. It
  can be at the top, or else if you have certain base settings lower in the init file,
  these might not reflect or load up in the keymap package so better to load them
  after some of the core settings have been set.</p>\n<p>Let's create the package
  and dump all our maps into the keymap file package.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">-- ~/.config/nvim/keymap.lua</span>\n\n<span class=\"kr\">function</span>
  <span class=\"nf\">map</span><span class=\"p\">(</span><span class=\"n\">mode</span><span
  class=\"p\">,</span> <span class=\"n\">lhs</span><span class=\"p\">,</span> <span
  class=\"n\">rhs</span><span class=\"p\">,</span> <span class=\"n\">opts</span><span
  class=\"p\">)</span>\n    <span class=\"kd\">local</span> <span class=\"n\">options</span>
  <span class=\"o\">=</span> <span class=\"p\">{</span> <span class=\"n\">noremap</span>
  <span class=\"o\">=</span> <span class=\"kc\">true</span> <span class=\"p\">}</span>\n
  \   <span class=\"kr\">if</span> <span class=\"n\">opts</span> <span class=\"kr\">then</span>\n
  \       <span class=\"n\">options</span> <span class=\"o\">=</span> <span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">tbl_extend</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;force&quot;</span><span class=\"p\">,</span> <span class=\"n\">options</span><span
  class=\"p\">,</span> <span class=\"n\">opts</span><span class=\"p\">)</span>\n    <span
  class=\"kr\">end</span>\n    <span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">api</span><span class=\"p\">.</span><span class=\"n\">nvim_set_keymap</span><span
  class=\"p\">(</span><span class=\"n\">mode</span><span class=\"p\">,</span> <span
  class=\"n\">lhs</span><span class=\"p\">,</span> <span class=\"n\">rhs</span><span
  class=\"p\">,</span> <span class=\"n\">options</span><span class=\"p\">)</span>\n<span
  class=\"kr\">end</span>\n\n<span class=\"n\">map</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;w&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;:w&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;q&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;:q!&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;&lt;leader&gt;s&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;:so %&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;cpp&#39;</span> <span class=\"s1\">&#39;:!c++
  % -o %:r &amp;&amp; %:r&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span class=\"n\">map</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;c,&#39;</span> <span class=\"s1\">&#39;:!gcc % -o %:r &amp;&amp;
  %:r&lt;CR&gt;&#39;</span><span class=\"p\">)</span>\n<span class=\"n\">map</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;py&#39;</span> <span class=\"s1\">&#39;:!python %&lt;cr&gt;&#39;</span><span
  class=\"p\">)</span>\n<span class=\"n\">map</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;n&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;go&#39;</span>
  <span class=\"s1\">&#39;:!go run %&lt;cr&gt;&#39;</span><span class=\"p\">)</span>\n<span
  class=\"n\">map</span><span class=\"p\">(</span><span class=\"s1\">&#39;n&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;sh&#39;</span> <span class=\"s1\">&#39;:!bash
  %&lt;cr&gt;&#39;</span><span class=\"p\">)</span>\n\n<span class=\"c1\">-- more
  keymaps</span>\n</pre></div>\n\n</pre>\n\n<p>So, this might work if you don't have
  any plugin-related keymaps as it would require those functions or objects available
  to execute. So, we might also want to separate plugins and load them first in our
  keymaps or in the init file.</p>\n<p>Now, there needs to be a way for grabbing a
  package. Yes, there is basically like import in python or any other programming
  language, lua has <code>require</code> keyword for importing packages. Since the
  <code>init</code> file and the <code>keymaps</code> are in the same folder path,
  we can simply say, <code>require(&quot;keymap&quot;)</code> in our <code>init.lua</code>
  file. Now, it depends on your config where you want to load this package. At the
  top i.e. at the beginning of neovim instance or after loading the plugins down.</p>\n<pre
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">-- init.lua</span>\n\n<span class=\"nb\">require</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;keymaps&quot;</span><span class=\"p\">)</span>\n\n<span class=\"c1\">--
  At the top</span>\n<span class=\"c1\">-- OR</span>\n<span class=\"c1\">-- After
  loading Packer plugins</span>\n</pre></div>\n\n</pre>\n\n<p>So, now you can separate
  all your configs as per your requirement. It is like splitting up a puzzle and again
  combining them. Similar package can be created for <code>plugins</code>, <code>options</code>
  or <code>lsp</code> configurations.</p>\n<h3>Creating a separate module</h3>\n<p>Now,
  we have seen how to create a lua package and loading in neovim. We also can create
  modules in neovim configuration. For instance, first, the init file is loaded, Other
  files might not be required hence it is not loaded by default, it is only loaded
  when <code>require</code>ed. So, we can create a module in lua with a folder, and
  inside of it, we can have packages as we had in the previous method. Also, every
  module can have a init file loaded first when we require that module. The rest of
  the packages in the module are thus made available thereafter.</p>\n<ul>\n<li>Module
  is a component not loaded by default</li>\n<li>Only loaded when required (literally
  require)</li>\n<li>Every module can have a <code>init.lua</code> file loaded first
  when required.</li>\n<li>Require a package in module from outside -&gt; <code>require('module_name.package_name')</code></li>\n</ul>\n<p>So,
  in neovim, we need to create a <code>lua</code> folder for placing all our modules
  so that the lua runtime is picked up correctly. Inside this lua folder, we can create
  a module basically a folder. This folder name can be anything you like, I like to
  keep it my nickname, you can use whatever you prefer.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span># ~/.config/nvim\n\n-- init.lua\n--
  lua/\n    -- module_name/\n        -- init.lua\n        -- package_name.lua\n        --
  keymaps.lua\n</pre></div>\n\n</pre>\n\n<p>Now, we can create packages in this module.
  You can move your keymaps package inside this folder. The keymaps package is nothing
  here but an example provided in the previous section for creating a package. You
  can basically dump all your keymaps in a single lua file and then import it from
  the init file. Similarly you can create a package inside a module and import it
  from the init file of the module(local init file <code>~/.config/nvim/lua/module_name/init.lua</code>)
  or the global init file(<code>~/.config/nvim/init.lua</code>).</p>\n<p>To load the
  packages, you have to use the same require statement irrespective of where you are
  importing from i.e. either from the module or from the global init file. The require
  statement would look like the following <code>require(&quot;module_name/package_name&quot;)</code>.
  Now, we can use the keymaps package to import from the module init file and then
  import the module from the global init file. To import a module, we can simply use
  the module name in the require string as <code>require(&quot;module_name&quot;)</code>.</p>\n<pre
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
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">-- ~/.config/nvim</span>\n\n<span class=\"c1\">-- lua/module_name/options.lua</span>\n\n<span
  class=\"n\">vim</span><span class=\"p\">.</span><span class=\"n\">opt</span><span
  class=\"p\">.</span><span class=\"n\">number</span> <span class=\"o\">=</span> <span
  class=\"kc\">true</span>\n<span class=\"n\">vim</span><span class=\"p\">.</span><span
  class=\"n\">opt</span><span class=\"p\">.</span><span class=\"n\">tabstop</span>
  <span class=\"o\">=</span> <span class=\"mi\">4</span>\n<span class=\"n\">vim</span><span
  class=\"p\">.</span><span class=\"n\">opt</span><span class=\"p\">.</span><span
  class=\"n\">swapfile</span> <span class=\"o\">=</span> <span class=\"kc\">false</span>\n\n\n<span
  class=\"c1\">-- lua/module_name/plugins.lua</span>\n\n<span class=\"nb\">require</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;module_name.options&quot;</span><span
  class=\"p\">)</span>\n<span class=\"kr\">return</span> <span class=\"nb\">require</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;packer&#39;</span><span class=\"p\">).</span><span
  class=\"n\">startup</span><span class=\"p\">(</span><span class=\"kr\">function</span><span
  class=\"p\">()</span>\n  <span class=\"n\">use</span> <span class=\"s1\">&#39;wbthomason/packer.nvim&#39;</span>\n
  \ <span class=\"c1\">--plugins</span>\n<span class=\"kr\">end</span><span class=\"p\">)</span>\n\n<span
  class=\"c1\">-- lua/module_name/keymap.lua</span>\n\n<span class=\"nb\">require</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;module_name.plugins&quot;</span><span
  class=\"p\">)</span>\n<span class=\"c1\">-- maps()</span>\n\n\n<span class=\"c1\">--
  lua/module_name/init.lua</span>\n\n<span class=\"nb\">require</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;module_name.keymaps)</span>\n\n\n<span class=\"s2\">-- init.lua</span>\n\n<span
  class=\"s2\">require(&quot;</span><span class=\"n\">module_name</span><span class=\"s2\">&quot;)</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  this is how we can create modules and packages for configurations in neovim using
  lua. This is also a kind of a structure for creating your own neovim plugin with
  lua!</p>\n<p>For further references, you can check out my <a href=\"https://github.com/Mr-Destructive/dotfiles\">dotfiles</a>.</p>\n<h3>References</h3>\n<ul>\n<li><a
  href=\"https://vonheikemen.github.io/devlog/tools/configuring-neovim-using-lua/\">Configure
  Neovim for Lua</a></li>\n<li><a href=\"https://alpha2phi.medium.com/neovim-for-beginners-init-lua-45ff91f741cb\">Neovim
  with Lua for beginners</a></li>\n<li><a href=\"https://www.youtube.com/c/TJDeVries/videos\">TJ
  Devries Youtube</a></li>\n</ul>\n<h2>Conclusion</h2>\n<p>So, that is just a basic
  overview of how you can get your neovim configured for lua. It is a great experience
  to just create such a personalized environment and play with it and have fun. You
  might have hopefully configured Neovim for Lua from this little guide. Maybe in
  the next article, I'll set up LSP in Neovim. If you have any queries or feedback,
  please let me know. Thank you for reading.</p>\n<p>Happy Viming :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/vimtolua.png
long_description: It has been a while since I have written a Vim article. Finally,
  I got some ideas after configuring my Neovim setup for Lua. I recently migrated
  to Ubuntu a couple of months back and it has been a cool change from Windows 7 In
  this article, we I have
now: 2025-05-01 18:11:33.313059
path: blog/posts/2022-07-07-Neovim-Lua.md
prevnext: null
slug: neovim-vimscript-to-lua
status: published
tags:
- vim
- lua
templateKey: blog-post
title: Configure Neovim in Lua
today: 2025-05-01
---

## Introduction

It has been a while since I have written a Vim article. Finally, I got some ideas after configuring my Neovim setup for Lua. I recently migrated to Ubuntu a couple of months back and it has been a cool change from Windows 7! 

In this article, we'll see how you can set up neovim for Lua. Since Neovim 0.5, it supports lua out of the box, so in the recent release 0.7, it added more native support to lua making it a lot easier to configure and play with neovim. So, we will see how we can use lua to convert all the 200 liner vimscript into lua (We can even have packages and modules:) We will cover how to configure your keymaps, pull up all the plugins, vim options, and other customizations.

## Why move to Lua?

I have used Vimscript for quite a while now, configured it as per my needs, and also made a few plugins like [frontmatter](https://github.com/Mr-Destructive/frontmatter.vim), [dj.vim](https://github.com/Mr-Destructive/dj.vim), and [commenter](https://github.com/Mr-Destructive/commenter.vim) which are quite clunky and not robust in terms of usage and customizability. Vimscript is good but it's a bit messy when you want extreme customization. 

I recently wanted to go full Neovim, I was kind of stuck in Vimscript and some of my plugins work for me but it might not work for others, it might be a piece of gibberish vimscript dumped. So, Why not have full native experience in Neovim if you can. It has now baked-in support for LSP, Debugging, Autocommands, and so much more. If you have Neovim 0.5+ you should be good to go full lua.

## Backup Neovim Config

Firstly, keep your original neovim/vim init files safe, back them up, make a copy and save it with a different name like `nvim_config.vim`. If you already have all your config files backed up as an ansible script or dotfiles GitHub repository, you can proceed ahead. 

But don't keep the `init.vim` file as it is in the `~/.config/nvim` folder, it might override after we configure the lua scripts.

## Basic Configuration

So, I assume you have set up Neovim, If not you need to follow some simple steps like downloading the package and making sure your neovim environment is working with vimscript first. The [Neovim Wiki](https://github.com/neovim/neovim/wiki/Installing-Neovim) provides great documentation on how to install neovim on various platforms using different methods.

After your neovim is set up and you have a basic configuration, you can now start to migrate into lua.
Create a `init.lua` file in the same path as your `init.lua` file is i.e. at `~/.config/nvim` or `~/AppData/Local/nvim/` for Windows. That's why it is recommended to keep the initial configuration vimscript file in a safe place. While migrating from vimscript to lua, once the lua file is created and the next time you restart neovim, the default settings will be from `init.lua` and not `init.vim`, so be prepared.

Firstly, you need to configure some options like `number`, `syntax highlighting`, `tabs`, and some `keymaps` of course. We can use the `vim.opt` method to set options in vim using lua syntax. So, certain corresponding vim options would be converted as follows:

If you have the following kind of settings in your vimrc or init.vim:

```vimscript
-- vimscript
set number
set tabstop=4 
set shiftwidth=4 
set softtabstop=0 
set expandtab 
set noswapfile
```
The above settings are migrated into lua syntax as follows:

```lua
--lua
vim.opt.number = true
vim.opt.tabstop = 4
vim.opt.shiftwidth = 4
vim.opt.softtabstop = 0
vim.opt.expandtab = true
vim.opt.swapfile = false
```

You can set other options in your config file accordingly. If you get sick of writing `vim.opt.` again and again, you can use a variable set to `vim.opt` and then access that variable to set the options. Something of the lines as below:

```lua
local set = vim.opt

set.number = true
set.tabstop = 4
set.shiftwidth = 4
set.softtabstop = 0
set.expandtab = true
set.swapfile = false
```

We can create a variable in lua like `local variable_name = something` so, we have created a variable `set` which is assigned to the value of `vim.opt` which is a method(function) in lua to set the options from the vimscript environment. Finally, access the `set` keyword to set the options. Using the `set` word is not necessary, you can use whatever you want. 

After setting up the basic options, you can source the file with `:so %` from the command mode. Just normally as you source the config files.

### Using Lua in Command Mode

We can use the lua functions or any other commands from the command mode in neovim using the lua command. Just prefix the command with `:lua` and after that, you can use lua syntax like function calling or setting variables, logging things, etc.

![Lua in Command Mode](https://res.cloudinary.com/techstructive-blog/image/upload/v1657380885/blog-media/lua_in_command_mode.gif)

## Adding Keymaps

Now, that we have some basic config setup, we can quickly get the keymaps. It's not that hard to make keymaps to set up in lua. To create keymaps, we have two options. One is compatible with Neovim and the other is compatible with Vim as well.

1. vim.keymap.set OR 
2. vim.api.nvim_set_keymap

Personally, I don't see a difference in terms of usage but [vim.keymap.set](https://github.com/neovim/neovim/blob/master/runtime/lua/vim/keymap.lua#L51) is a wrapper around [nvim_set_keymap](https://github.com/neovim/neovim/blob/master/src/nvim/api/vim.c#L1451). So, it is really a matter of personal preference which you want to use. 

So, both the functions have quite similar parameters:

```lua
vim.keymap.set({mode}, {lhs}, {rhs}, {options})

vim.api.nvim_set_keymap({mode}, {lhs}, {rhs}, {options})
```

The advantage of `vim.keymap.set` over `vim.api.nvim_set_keymap` is that it allows directly calling lua functions rather than calling vimscripty way like `:lua function()`, so we directly can use lua code in the RHS part of the function parameter.

Let's take a basic example mapping:

```
vim.keymap.set('n', 'Y', 'yy', {noremap = false})
```

Here, we have mapped `Shift+y` with the keys `yy` in `n`ormal mode. The first argument is the mode, it can be a single-mode like `'n'`, `'v'`, `'i'`, etc., or a multi-mode table like `{'n', 'v'}`, `{'v', 'i'}`, etc. 

The next parameter is also a string but it should be the key for triggering the mapping. In this case, we have used `Y` which is `Shift + y`, it can be any key shortcut you want to map.

The third parameter is the string which will be the command to be executed after the key has been used. Here we have used the keys `yy`, if the map is from a command mode, you will be using something like `':commands_to_be executed'` as the third parameter.

The fourth parameter which is optional can contain [special arguments](https://neovim.io/doc/user/api.html#:~:text=nvim_set_keymap(%7Bmode%7D%2C%20%7Blhs%7D%2C%20%7Brhs%7D%2C%20%7B*opts%7D)%20%20%20%20%20%20%20%20%20%20%20%20%20*nvim_set_keymap()*). We have set a default option which is `noremap` as true, the options are not string but lua tables instead, so it can be similar to python dictionary or a map kind of a structure with a key value pair.


One more important aspect in keymapping might about the leader key, you can set your leader key by using the global vim options with `vim.g` and access `mapleader` to set it to the key you wish. This will make the `leader` key available to us and thereafter we can map the leader key in custom mappings.

```
vim.g.mapleader = " "
```

Here, I have set my leader key to the `<Space>` key. Now, we can map keys to the existing keymaps in the vimscript. Let's map some basic keymaps first and then after setting up the plugins,we can move into plugin-specific mappings.

You can also use `vim.api.nvim_set_keymap` function with the same parameters as well. 

```lua
vim.keymap.set('n', '<leader>w', ':w<CR>',{noremap = true})
vim.keymap.set('n', '<leader>q', ':q!<CR>',{noremap = true})
vim.keymap.set('n', '<leader>s', ':so %<CR>',{noremap = true})
vim.keymap.set('n', '<leader>ev', ':vsplit $MYVIMRC<CR>',{noremap = true})
vim.keymap.set('n', '<leader>sv', ':w<CR>:so %<CR>:q<CR>',{noremap = true})

OR

vim.api.nvim_set_keymap('n', '<leader>w', ':w<CR>',{noremap = true})
vim.api.nvim_set_keymap('n', '<leader>q', ':q!<CR>',{noremap = true})
vim.api.nvim_set_keymap('n', '<leader>s', ':so %<CR>',{noremap = true})
vim.api.nvim_set_keymap('n', '<leader>ev', ':vsplit $MYVIMRC<CR>',{noremap = true})
vim.api.nvim_set_keymap('n', '<leader>sv', ':w<CR>:so %<CR>:q<CR>',{noremap = true})
```

If, you don't like writing `vim.keymap.set` or `vim.api.nvim_set_keymap` again and again, you can create a simpler function for it. In lua a function is created just like a variable by specifying the scope of the function i.e. local followed by the `function` keyword and finally the name of the function and parenthesis. The function body is terminated by the `end` keyword.

```lua
function map(mode, lhs, rhs, opts)
    local options = { noremap = true }
    if opts then
        options = vim.tbl_extend("force", options, opts)
    end
    vim.api.nvim_set_keymap(mode, lhs, rhs, options)
end
```
Now, in this function `map`, we have passed in the same parameters like the `vim.keymap.set` function takes but we have just parsed the function in a shorter and safer way by adding `noremap = true` by default. So this is just a helper function or a verbose function for calling the vim.keymap.set function.

To use this function, we can simply call `map` with the same arguments as given to the prior functions.

```lua
map('n', '<leader>w', ':w<CR>')
map('n', '<leader>q', ':q!<CR>')
map('n', '<leader>s', ':so %<CR>')
```

Notice here, though, we have not passed the `{noremap = true}` as it will be passed by default to the `vim.api.nvim_set_keymap` or `vim.keymap.set` function via the custom map function.

If you want some more examples, here are some additional mapping specific to languages, meant for compiling or running scripts with neovim instance. 

```lua
-- vimscript

nnoremap cpp :!c++ % -o %:r && %:r<CR>
nnoremap c, :!gcc % -o %:r && %:r<CR>
nnoremap py :!python %<cr>
nnoremap go :!go run %<cr>
nnoremap sh :!bash %<CR>


-- lua

map('n', 'cpp' ':!c++ % -o %:r && %:r<CR>')
map('n', 'c,' ':!gcc % -o %:r && %:r<CR>')
map('n', 'py' ':!python %<cr>')
map('n', 'go' ':!go run %<cr>')
map('n', 'sh' ':!bash %<cr>')

```

So, this is how we can set up our keymaps in lua. You can customize this function as per your needs. These are just made for the understanding purpose of how to reduce the repetitive stuff in the setup.

**If you are really stuck up and not feeling to convert those mappings into lua then I have a function that can do it for you, check out my dotfile repo -> [keymapper](https://github.com/Mr-Destructive/dotfiles/blob/master/nvim/lua/destructive/options.lua#L9)**

## Adding Plugin Manager

Now, we really missing some plugins, aren't we? So, the neovim community has some good choices for using a new plugin manager written in core lua. It is usually a good idea to move into lua completely rather than switching to and fro between vimscript and lua.

So, [Packer](https://github.com/wbthomason/packer.nvim) is the new plugin manager for Neovim in Lua, there is other plugin managers out there as well like [paq](https://github.com/savq/paq-nvim). If you don't want to switch with the plugin manager, you can still use vim-based plugin managers like [Vim-Plug](https://dev.to/vonheikemen/neovim-using-vim-plug-in-lua-3oom).

So, let's install the Packer plugin manager into Neovim. We simply have to run the following command in the console and make sure the plugin manager is configured correctly.

```
# Linux

git clone --depth 1 https://github.com/wbthomason/packer.nvim\
~/.local/share/nvim/site/pack/packer/start/packer.nvim


# Windows

git clone https://github.com/wbthomason/packer.nvim "$env:LOCALAPPDATA\nvim-data\site\pack\packer\start\packer.nvim"
```

Now, if you open a new neovim instance and try to run the command `:PackerClean`, and no error pops out that means you have configured it correctly. You can move ahead to installing plugins now. Yeah! PLUG-IN time! 

```lua
return require('packer').startup(function()
end)
```

First try to source the file, if it throws out errors it shouldn't try to fix the installation path of Packer. If the command succeded we can finally pull up some plugins.

Below are some of the plugins that you can use irrespective of what language preferences you would have. This includes basic dev-icons for the status line as well as the explorer window file icons. As usual, add your plugins and make them yours.

```lua

return require('packer').startup(function()
  use 'wbthomason/packer.nvim'
  use 'tpope/vim-fugitive'
  use {
    'nvim-lualine/lualine.nvim',
    requires = { 'kyazdani42/nvim-web-devicons', opt = true }
  }
  use 'tiagofumo/vim-nerdtree-syntax-highlight'
  use 'kyazdani42/nvim-web-devicons'
  use 'vim-airline/vim-airline'
  use 'vim-airline/vim-airline-themes'
end)
```

After adding the list of your plugins, you need to source the file and then install the plugins with the command `:PackerInstall`. This function will install all the plugins after the file has been sourced so make sure you don't forget it.

## Colors and Color Themes

You might fancy some good-looking and aesthetic setup for neovim of course! In Neovim, we already have a wide variety of configurations to set up like color schemes, GUI colors, terminal colors, etc. You can pick up a color scheme from a large list of awesome color schemes from [GitHub](https://github.com/topics/neovim-colorscheme). 

After choosing the theme, plug it in the packer plugin list which we just created and source the file and finally run `:PackerInstall`. This should install the plugin. You can then set the colorscheme as per your preference, first view the color scheme temporarily on the current instance with the command `:colorscheme colorscheme_name`. 

```lua
return require('packer').startup(function()
  use 'wbthomason/packer.nvim'
  -- 
  use 'Mofiqul/dracula.nvim'
  --
end)
```

You can then add the command to set it as the default color scheme for Neovim.

```lua
vim.cmd [[silent! colorscheme dracula]]
```

You can change the background color, text color, icons style and terminal and gui style separately with the augroup and setting the colorscheme commands.

```lua
vim.api.nvim_command([[
    augroup ChangeBackgroudColour
        autocmd colorscheme * :hi normal termbg=#000030 termfg=#ffffff
        autocmd colorscheme * :hi Directory ctermfg=#ffffff
    augroup END
]])
vim.o.termguicolors = true
```

Here, I have used the background and foreground colors of the terminal variant of Neovim. Also for the Directory Explorer i.e. netrw, I have changed the terminal foreground. This you can configure as per your needs, Though this is still vimscripty, there are Autocommands and autogroups available sooner in Neovim.

## Separating Configurations

If you wish to keep all your config in one file i.e. `init.lua` file, you can, though you can separate out things like `keymaps`, `plugins`, `custom_options` or if you have `lsp` configurations into separate lua packages or creating a separate module. This helps in loading up things as per requirement and also looks readable, making it a lot easier to test out things without breaking the whole config. 

Definitely, there will be personal preferences and pros and cons of both approaches, you can pick up whatever suits your style.

### Creating separate packages 

To create a separate package, we can simply add a file in the same folder as `init.vim` i.e. in the folder `~/.config/nvim` or equivalent for windows. The package name can be any valid filename with the `lua` extension. 

For instance, you can create a package for all your keymaps and load it in the `init.lua` as per the order you want to load them. It can be at the top, or else if you have certain base settings lower in the init file, these might not reflect or load up in the keymap package so better to load them after some of the core settings have been set.

Let's create the package and dump all our maps into the keymap file package.

```lua
-- ~/.config/nvim/keymap.lua

function map(mode, lhs, rhs, opts)
    local options = { noremap = true }
    if opts then
        options = vim.tbl_extend("force", options, opts)
    end
    vim.api.nvim_set_keymap(mode, lhs, rhs, options)
end

map('n', '<leader>w', ':w<CR>')
map('n', '<leader>q', ':q!<CR>')
map('n', '<leader>s', ':so %<CR>')
map('n', 'cpp' ':!c++ % -o %:r && %:r<CR>')
map('n', 'c,' ':!gcc % -o %:r && %:r<CR>')
map('n', 'py' ':!python %<cr>')
map('n', 'go' ':!go run %<cr>')
map('n', 'sh' ':!bash %<cr>')

-- more keymaps

```

So, this might work if you don't have any plugin-related keymaps as it would require those functions or objects available to execute. So, we might also want to separate plugins and load them first in our keymaps or in the init file.

Now, there needs to be a way for grabbing a package. Yes, there is basically like import in python or any other programming language, lua has `require` keyword for importing packages. Since the `init` file and the `keymaps` are in the same folder path, we can simply say, `require("keymap")` in our `init.lua` file. Now, it depends on your config where you want to load this package. At the top i.e. at the beginning of neovim instance or after loading the plugins down.

```lua
-- init.lua

require("keymaps")

-- At the top
-- OR
-- After loading Packer plugins
```

So, now you can separate all your configs as per your requirement. It is like splitting up a puzzle and again combining them. Similar package can be created for `plugins`, `options` or `lsp` configurations.

### Creating a separate module

Now, we have seen how to create a lua package and loading in neovim. We also can create modules in neovim configuration. For instance, first, the init file is loaded, Other files might not be required hence it is not loaded by default, it is only loaded when `require`ed. So, we can create a module in lua with a folder, and inside of it, we can have packages as we had in the previous method. Also, every module can have a init file loaded first when we require that module. The rest of the packages in the module are thus made available thereafter.

- Module is a component not loaded by default
- Only loaded when required (literally require)
- Every module can have a `init.lua` file loaded first when required.
- Require a package in module from outside -> `require('module_name.package_name')`

So, in neovim, we need to create a `lua` folder for placing all our modules so that the lua runtime is picked up correctly. Inside this lua folder, we can create a module basically a folder. This folder name can be anything you like, I like to keep it my nickname, you can use whatever you prefer. 

```
# ~/.config/nvim

-- init.lua
-- lua/
    -- module_name/
        -- init.lua
        -- package_name.lua
        -- keymaps.lua
```

Now, we can create packages in this module. You can move your keymaps package inside this folder. The keymaps package is nothing here but an example provided in the previous section for creating a package. You can basically dump all your keymaps in a single lua file and then import it from the init file. Similarly you can create a package inside a module and import it from the init file of the module(local init file `~/.config/nvim/lua/module_name/init.lua`) or the global init file(`~/.config/nvim/init.lua`). 

To load the packages, you have to use the same require statement irrespective of where you are importing from i.e. either from the module or from the global init file. The require statement would look like the following `require("module_name/package_name")`. Now, we can use the keymaps package to import from the module init file and then import the module from the global init file. To import a module, we can simply use the module name in the require string as `require("module_name")`.

```lua
-- ~/.config/nvim

-- lua/module_name/options.lua

vim.opt.number = true
vim.opt.tabstop = 4
vim.opt.swapfile = false


-- lua/module_name/plugins.lua

require("module_name.options")
return require('packer').startup(function()
  use 'wbthomason/packer.nvim'
  --plugins
end)

-- lua/module_name/keymap.lua

require("module_name.plugins")
-- maps()


-- lua/module_name/init.lua

require("module_name.keymaps)


-- init.lua

require("module_name")

```

So, this is how we can create modules and packages for configurations in neovim using lua. This is also a kind of a structure for creating your own neovim plugin with lua!

For further references, you can check out my [dotfiles](https://github.com/Mr-Destructive/dotfiles). 
### References

- [Configure Neovim for Lua](https://vonheikemen.github.io/devlog/tools/configuring-neovim-using-lua/)
- [Neovim with Lua for beginners](https://alpha2phi.medium.com/neovim-for-beginners-init-lua-45ff91f741cb)
- [TJ Devries Youtube](https://www.youtube.com/c/TJDeVries/videos)

## Conclusion

So, that is just a basic overview of how you can get your neovim configured for lua. It is a great experience to just create such a personalized environment and play with it and have fun. You might have hopefully configured Neovim for Lua from this little guide. Maybe in the next article, I'll set up LSP in Neovim. If you have any queries or feedback, please let me know. Thank you for reading.

Happy Viming :)