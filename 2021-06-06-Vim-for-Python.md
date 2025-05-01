---
article_html: "<h2>Introduction</h2>\n<p>Vim is quite a powerful text editor which
  can add performance to the already fast typed language Python. Vim can be highly
  customizable and efficient to use as it has the power of <strong>adding custom plugins
  and plugins managers, key mappings</strong>, and the most critical weapon of vim
  - Access to the terminal straight away.\nThis is not a full-featured guide of using
  vim for python, it's just a quick setup for using python on vim blazingly fast!!</p>\n<h2>Plugin
  Managers</h2>\n<p>So let us start making Vim, the text editor suitable for a python
  programmer. Firstly we'll need the vim plugin manager. There are different plugin
  managers out there, each of them has the same purpose to install, upgrade and manage
  the plugins for vim. You can install any one of them and get up and running.</p>\n<ul>\n<li><a
  href=\"https://www.vim.org/scripts/script.php?script_id=4828\">Vim Plug</a></li>\n<li><a
  href=\"https://github.com/VundleVim/Vundle.vim\">Vundle</a></li>\n<li><a href=\"https://github.com/tpope/vim-pathogen\">Pathogen</a></li>\n</ul>\n<p>These
  are some of the finest and well-supported plugin managers in vim. You can use any
  of these plugin managers, and get started by installing some plugins.</p>\n<h2>JEDI-VIM-
  Auto completion</h2>\n<p>Firstly I will like to install Jedi for code completion
  in Python. The plugin can be simple and straightforward to install using any of
  the above plugin managers. Jedi-Vim provides some neat and clean** syntax analytics
  and autocompletion for Python in Vim**. You'll find the docs and installation process
  here  <a href=\"https://github.com/davidhalter/jedi-vim\">JEDI-VIM </a></p>\n<h2>NERDTree-File
  manager</h2>\n<p>Next, It would be great if we install a file manager for managing
  the files and folders in the code directories. We can simply use the Nerdtree plugin
  for this. NerdTree is quite a <strong>fantastic plugin for file management in Vim</strong>.
  It simply makes Vim feel like VS Code. The installation and docs can be found here
  \ <a href=\"https://github.com/preservim/nerdtree\">NERDTree</a>.</p>\n<p>Nerdtree
  commands can be longer to write, for that let's start mapping and for that, we can
  start editing our Vimrc.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">set</span> <span class=\"k\">number</span>\n<span class=\"nb\">syntax</span>
  enable\n<span class=\"k\">filetype</span> indent <span class=\"k\">on</span>\n<span
  class=\"k\">set</span> <span class=\"nb\">tabstop</span><span class=\"p\">=</span><span
  class=\"m\">4</span>\n<span class=\"k\">set</span> <span class=\"nb\">softtabstop</span><span
  class=\"p\">=</span><span class=\"m\">4</span>\n<span class=\"k\">set</span> <span
  class=\"nb\">autoindent</span> \n<span class=\"k\">set</span> <span class=\"nb\">encoding</span><span
  class=\"p\">=</span>utf<span class=\"m\">-8</span>\n</pre></div>\n\n</pre>\n\n<p>This
  can be some addition to your existing vimrc as you might have a setup for plugin
  managers. You can choose the color scheme of your choice, don't waste time selecting
  the color scheme. Feel free and modify the vimrc according to your knowledge and
  choice.</p>\n<h2>Keymappings</h2>\n<p>We move on to the Key mappings for NERDTree
  and other features. You can make mappings generally in the normal mode but there
  might be cases where you need to use maps for visual mode or insert mode as well,
  that entirely depends on the user :)</p>\n<p>To map in normal mode, we'll its command
  to be specific:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">nnoremap</span> <span class=\"p\">&lt;</span>C<span class=\"p\">-</span><span
  class=\"k\">n</span><span class=\"p\">&gt;</span> :NERDTree<span class=\"p\">&lt;</span>CR<span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>This will map CTRL+n to open
  the NERDTree file manager to the left, saving a bit of time and avoiding frustration.
  Feel free to add any keymap of your choice, this is just for demonstration.\nYou
  can further automate NERDTree for switching between tabs because it makes you type
  CTRL+w twice, you can reduce that to just typing w.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"nb\">nnoremap</span> <span class=\"k\">w</span>:<span class=\"p\">&lt;</span>C<span
  class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span>C<span
  class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<h2>Integrated
  Terminal Macros</h2>\n<p>We can open a terminal window like a split between the
  editor. We can simply use the command :terminal to split the window horizontally,
  where the upper split will be terminal and the down window will have the editor.
  This is quite a neat feature of Vim in that it blends with the terminal so well
  so that we can switch between the terminal and the editor very quickly. For that,
  you can create a macro if you need to fire up a terminal again and again.</p>\n<pre
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
  class=\"nb\">nnoremap</span> <span class=\"p\">&lt;</span>C<span class=\"p\">-</span><span
  class=\"k\">t</span><span class=\"p\">&gt;</span> :terminal<span class=\"p\">&lt;</span>CR<span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>If you place the above macro
  in your vimrc and then type Ctrl+t, the exact thing will happen to fire up a terminal
  split but with fewer keystrokes and without leaving the normal mode.\nAlso, the
  NERDTree macro can be also fruitful with this as it will make a full-blown IDE-like
  feeling inside of Vim.\n<img src=\"https://s6.gifyu.com/images/screenrecording.gif\"
  alt=\"Demonstrate macros for NERDTree and terminal split\" /></p>\n<h2>Running the
  Code with a snap</h2>\n<p>We can automate the process of running python scripts
  inside of vim. Instead of typing out the entire command for executing python script
  from vim. We can use keymaps for it as they can significantly boost the time required
  to run and debug the code.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">nnoremap</span> <span class=\"k\">py</span> :<span class=\"p\">!</span>python
  %\n</pre></div>\n\n</pre>\n\n<p>This is a small map but can save a lot of time and
  give some motivation to use vim as you run the code blazingly faster than other
  editors. I have used py, but it can cause some problems as p is already mapped for
  pasting. So it's better to use other key combinations such as ty, yh, or any other
  key combination of your choice. Try it out and add your own flavor that's how we
  all learn.</p>\n<p>So, that's the basic set-up for python on vim, you can make more
  custom mappings, find more plugins and test out which work out the best for your
  workflow. Happy Coding and Viming ;)</p>\n"
cover: ''
date: 2021-06-06
datetime: 2021-06-06 00:00:00+00:00
description: 'Vim is quite a powerful text editor which can add performance to the
  already fast typed language Python. Vim can be highly customizable and efficient
  to use as '
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-06-06-Vim-for-Python.md
html: "<h2>Introduction</h2>\n<p>Vim is quite a powerful text editor which can add
  performance to the already fast typed language Python. Vim can be highly customizable
  and efficient to use as it has the power of <strong>adding custom plugins and plugins
  managers, key mappings</strong>, and the most critical weapon of vim - Access to
  the terminal straight away.\nThis is not a full-featured guide of using vim for
  python, it's just a quick setup for using python on vim blazingly fast!!</p>\n<h2>Plugin
  Managers</h2>\n<p>So let us start making Vim, the text editor suitable for a python
  programmer. Firstly we'll need the vim plugin manager. There are different plugin
  managers out there, each of them has the same purpose to install, upgrade and manage
  the plugins for vim. You can install any one of them and get up and running.</p>\n<ul>\n<li><a
  href=\"https://www.vim.org/scripts/script.php?script_id=4828\">Vim Plug</a></li>\n<li><a
  href=\"https://github.com/VundleVim/Vundle.vim\">Vundle</a></li>\n<li><a href=\"https://github.com/tpope/vim-pathogen\">Pathogen</a></li>\n</ul>\n<p>These
  are some of the finest and well-supported plugin managers in vim. You can use any
  of these plugin managers, and get started by installing some plugins.</p>\n<h2>JEDI-VIM-
  Auto completion</h2>\n<p>Firstly I will like to install Jedi for code completion
  in Python. The plugin can be simple and straightforward to install using any of
  the above plugin managers. Jedi-Vim provides some neat and clean** syntax analytics
  and autocompletion for Python in Vim**. You'll find the docs and installation process
  here  <a href=\"https://github.com/davidhalter/jedi-vim\">JEDI-VIM </a></p>\n<h2>NERDTree-File
  manager</h2>\n<p>Next, It would be great if we install a file manager for managing
  the files and folders in the code directories. We can simply use the Nerdtree plugin
  for this. NerdTree is quite a <strong>fantastic plugin for file management in Vim</strong>.
  It simply makes Vim feel like VS Code. The installation and docs can be found here
  \ <a href=\"https://github.com/preservim/nerdtree\">NERDTree</a>.</p>\n<p>Nerdtree
  commands can be longer to write, for that let's start mapping and for that, we can
  start editing our Vimrc.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">set</span> <span class=\"k\">number</span>\n<span class=\"nb\">syntax</span>
  enable\n<span class=\"k\">filetype</span> indent <span class=\"k\">on</span>\n<span
  class=\"k\">set</span> <span class=\"nb\">tabstop</span><span class=\"p\">=</span><span
  class=\"m\">4</span>\n<span class=\"k\">set</span> <span class=\"nb\">softtabstop</span><span
  class=\"p\">=</span><span class=\"m\">4</span>\n<span class=\"k\">set</span> <span
  class=\"nb\">autoindent</span> \n<span class=\"k\">set</span> <span class=\"nb\">encoding</span><span
  class=\"p\">=</span>utf<span class=\"m\">-8</span>\n</pre></div>\n\n</pre>\n\n<p>This
  can be some addition to your existing vimrc as you might have a setup for plugin
  managers. You can choose the color scheme of your choice, don't waste time selecting
  the color scheme. Feel free and modify the vimrc according to your knowledge and
  choice.</p>\n<h2>Keymappings</h2>\n<p>We move on to the Key mappings for NERDTree
  and other features. You can make mappings generally in the normal mode but there
  might be cases where you need to use maps for visual mode or insert mode as well,
  that entirely depends on the user :)</p>\n<p>To map in normal mode, we'll its command
  to be specific:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">nnoremap</span> <span class=\"p\">&lt;</span>C<span class=\"p\">-</span><span
  class=\"k\">n</span><span class=\"p\">&gt;</span> :NERDTree<span class=\"p\">&lt;</span>CR<span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>This will map CTRL+n to open
  the NERDTree file manager to the left, saving a bit of time and avoiding frustration.
  Feel free to add any keymap of your choice, this is just for demonstration.\nYou
  can further automate NERDTree for switching between tabs because it makes you type
  CTRL+w twice, you can reduce that to just typing w.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"nb\">nnoremap</span> <span class=\"k\">w</span>:<span class=\"p\">&lt;</span>C<span
  class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span>C<span
  class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<h2>Integrated
  Terminal Macros</h2>\n<p>We can open a terminal window like a split between the
  editor. We can simply use the command :terminal to split the window horizontally,
  where the upper split will be terminal and the down window will have the editor.
  This is quite a neat feature of Vim in that it blends with the terminal so well
  so that we can switch between the terminal and the editor very quickly. For that,
  you can create a macro if you need to fire up a terminal again and again.</p>\n<pre
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
  class=\"nb\">nnoremap</span> <span class=\"p\">&lt;</span>C<span class=\"p\">-</span><span
  class=\"k\">t</span><span class=\"p\">&gt;</span> :terminal<span class=\"p\">&lt;</span>CR<span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>If you place the above macro
  in your vimrc and then type Ctrl+t, the exact thing will happen to fire up a terminal
  split but with fewer keystrokes and without leaving the normal mode.\nAlso, the
  NERDTree macro can be also fruitful with this as it will make a full-blown IDE-like
  feeling inside of Vim.\n<img src=\"https://s6.gifyu.com/images/screenrecording.gif\"
  alt=\"Demonstrate macros for NERDTree and terminal split\" /></p>\n<h2>Running the
  Code with a snap</h2>\n<p>We can automate the process of running python scripts
  inside of vim. Instead of typing out the entire command for executing python script
  from vim. We can use keymaps for it as they can significantly boost the time required
  to run and debug the code.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">nnoremap</span> <span class=\"k\">py</span> :<span class=\"p\">!</span>python
  %\n</pre></div>\n\n</pre>\n\n<p>This is a small map but can save a lot of time and
  give some motivation to use vim as you run the code blazingly faster than other
  editors. I have used py, but it can cause some problems as p is already mapped for
  pasting. So it's better to use other key combinations such as ty, yh, or any other
  key combination of your choice. Try it out and add your own flavor that's how we
  all learn.</p>\n<p>So, that's the basic set-up for python on vim, you can make more
  custom mappings, find more plugins and test out which work out the best for your
  workflow. Happy Coding and Viming ;)</p>\n"
image_url: https://meetgor-cdn.pages.dev/vim-for-python.webp
long_description: Vim is quite a powerful text editor which can add performance to
  the already fast typed language Python. Vim can be highly customizable and efficient
  to use as it has the power of  So let us start making Vim, the text editor suitable
  for a python pro
now: 2025-05-01 18:11:33.312343
path: blog/posts/2021-06-06-Vim-for-Python.md
prevnext: null
slug: vim-for-python
status: published
subtitle: For the python programmers who need speed!
tags:
- python
- vim
templateKey: blog-post
title: Setting up Vim for Python
today: 2025-05-01
---

## Introduction
Vim is quite a powerful text editor which can add performance to the already fast typed language Python. Vim can be highly customizable and efficient to use as it has the power of **adding custom plugins and plugins managers, key mappings**, and the most critical weapon of vim - Access to the terminal straight away.
This is not a full-featured guide of using vim for python, it's just a quick setup for using python on vim blazingly fast!!

## Plugin Managers
So let us start making Vim, the text editor suitable for a python programmer. Firstly we'll need the vim plugin manager. There are different plugin managers out there, each of them has the same purpose to install, upgrade and manage the plugins for vim. You can install any one of them and get up and running.

-  [Vim Plug](https://www.vim.org/scripts/script.php?script_id=4828) 
-  [Vundle](https://github.com/VundleVim/Vundle.vim) 
-  [Pathogen](https://github.com/tpope/vim-pathogen) 

These are some of the finest and well-supported plugin managers in vim. You can use any of these plugin managers, and get started by installing some plugins.

## JEDI-VIM- Auto completion 
Firstly I will like to install Jedi for code completion in Python. The plugin can be simple and straightforward to install using any of the above plugin managers. Jedi-Vim provides some neat and clean** syntax analytics and autocompletion for Python in Vim**. You'll find the docs and installation process here  [JEDI-VIM ](https://github.com/davidhalter/jedi-vim) 

## NERDTree-File manager
Next, It would be great if we install a file manager for managing the files and folders in the code directories. We can simply use the Nerdtree plugin for this. NerdTree is quite a **fantastic plugin for file management in Vim**. It simply makes Vim feel like VS Code. The installation and docs can be found here  [NERDTree](https://github.com/preservim/nerdtree).

Nerdtree commands can be longer to write, for that let's start mapping and for that, we can start editing our Vimrc. 

```vim
set number
syntax enable
filetype indent on
set tabstop=4
set softtabstop=4
set autoindent 
set encoding=utf-8
``` 
This can be some addition to your existing vimrc as you might have a setup for plugin managers. You can choose the color scheme of your choice, don't waste time selecting the color scheme. Feel free and modify the vimrc according to your knowledge and choice. 

## Keymappings
We move on to the Key mappings for NERDTree and other features. You can make mappings generally in the normal mode but there might be cases where you need to use maps for visual mode or insert mode as well, that entirely depends on the user :)

To map in normal mode, we'll its command to be specific:


```vim
nnoremap <C-n> :NERDTree<CR>
``` 

This will map CTRL+n to open the NERDTree file manager to the left, saving a bit of time and avoiding frustration. Feel free to add any keymap of your choice, this is just for demonstration. 
You can further automate NERDTree for switching between tabs because it makes you type CTRL+w twice, you can reduce that to just typing w.

```vim
nnoremap w:<C-w><C-w>
``` 

## Integrated Terminal Macros
We can open a terminal window like a split between the editor. We can simply use the command :terminal to split the window horizontally, where the upper split will be terminal and the down window will have the editor. This is quite a neat feature of Vim in that it blends with the terminal so well so that we can switch between the terminal and the editor very quickly. For that, you can create a macro if you need to fire up a terminal again and again.
```vim
nnoremap <C-t> :terminal<CR>
```
If you place the above macro in your vimrc and then type Ctrl+t, the exact thing will happen to fire up a terminal split but with fewer keystrokes and without leaving the normal mode. 
Also, the NERDTree macro can be also fruitful with this as it will make a full-blown IDE-like feeling inside of Vim.
![Demonstrate macros for NERDTree and terminal split](https://s6.gifyu.com/images/screenrecording.gif)

## Running the Code with a snap

We can automate the process of running python scripts inside of vim. Instead of typing out the entire command for executing python script from vim. We can use keymaps for it as they can significantly boost the time required to run and debug the code. 


```vim
nnoremap py :!python %
``` 

This is a small map but can save a lot of time and give some motivation to use vim as you run the code blazingly faster than other editors. I have used py, but it can cause some problems as p is already mapped for pasting. So it's better to use other key combinations such as ty, yh, or any other key combination of your choice. Try it out and add your own flavor that's how we all learn.

So, that's the basic set-up for python on vim, you can make more custom mappings, find more plugins and test out which work out the best for your workflow. Happy Coding and Viming ;)