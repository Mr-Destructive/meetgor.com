---
article_html: "<h2>Introduction</h2>\n<p>Keymapping or Key binding is a process in
  which a user or a system can assign certain keys or commands to perform a particular
  task or commands. This can be quite useful in Vim as well as in other text editors
  to avoid some repetitive tasks and save time. In Vim this can be a great power to
  have for programmers as it can be really very flexible to set up for any programming
  language and it's\nenvironment.</p>\n<p>In Vim, you can map a particular keystroke
  combination to trigger a command or any operation. You can also map a key in place
  of a weird key combination as well. You can map a key to integrate several commands
  and make it run at a single key or without much of a hassle. Key Mapping is quite
  powerful in Vim, only your creativity and imagination are the limits here. Vim can
  really shine in such aspects where the user has the freedom to tailor the editor
  and his development environment as per his/her needs.</p>\n<h2>Understanding the
  structure of Key binding</h2>\n<p>You can basically map a key depending on the current
  mode you are in. So that means you can have the same key mapped to different commands
  depending on the mode. That is really flexible and powerful. Vim allows you to basically
  map in almost every mode such as normal, insert, visual, command, and any other
  existing modes. Let us see what a basic key mapping looks like.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>{mode}{attribute}map {key}
  {command}\n</pre></div>\n\n</pre>\n\n<p>The key mapping is broken down by adding
  the mode to map the key, an optional attribute to change/modify the way the mapping
  behaves and we finally have the key and the command pair. If you want to test a
  map you can temporarily map in the command mode using the syntax for the map command
  or if you want a permanent mapping, you can add the mappings in your vimrc file.</p>\n<p>For
  a basic example let us map c to paste from the clipboard.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"p\">:</span>map <span class=\"k\">c</span> <span class=\"c\">&quot;+p&lt;CR&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above command maps the &quot;c&quot; key in the <strong>normal/visual mode</strong>
  by default to paste the contents from the clipboard. It is just an example of how
  a basic key mapping or binding can be used to save your time and efforts. This is
  a very small example to demonstrate the concept of mapping in Vim but it scales
  quite fantastically and smoothly.</p>\n<p>Now let's see how to view the mapping
  inside of the current vim buffer. We can simply type map, imap, or nmap in the command
  mode to view the mappings in the respective modes. This will display the maps which
  are currently present inside of the vim buffer, some might be temporary and most
  of them will be permanent.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>:map \n</pre></div>\n\n</pre>\n\n<p>If
  you type in map being in the command mode, it will list out the maps. If you want
  to stop seeing the list just press q, and you will be back in the editor. You can
  also view the maps in the normal, visual, insert, or any other modes by prefixing
  n, v, i, or other modes with map.</p>\n<h2>Key Mapping in Modes</h2>\n<p>To map
  a key binding explicitly in the normal mode or any other modes, we have to prefix
  the word such as n for normal, v for visual or i for insert, and so on. The map
  will only work in the provided mode so the same keys may have several different
  maps depending on the current mode, but that can create a bit of confusion so avoid
  doing that. Let's map a key in the normal mode for example,</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"p\">:</span>nmap <span class=\"p\">&lt;</span>C<span class=\"p\">-</span>s<span
  class=\"p\">&gt;</span> :<span class=\"k\">w</span><span class=\"p\">&lt;</span>CR<span
  class=\"p\">&gt;</span>    \n</pre></div>\n\n</pre>\n\n<p>This is pretty bad for
  Vim beginners as it will spoil the real Vim experience of saving with :w, If you
  feel comfortable using Ctrl + s to save a file, the above map will be perfect. It
  will map the key Ctrl + s by pressing :w and then pressing enter to save changes
  to the file.</p>\n<p>We can also map certain non-alphabetical keys such as CTRL,
  ALT, and others, it will help you in saving time to leave a mode and enter the mode
  again to edit the text. The following is a perfect example of such.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"p\">:</span>imap <span class=\"p\">&lt;</span>C<span class=\"p\">-</span><span
  class=\"k\">c</span><span class=\"p\">&gt;</span> <span class=\"p\">&lt;</span>ESC<span
  class=\"p\">&gt;</span>&quot;<span class=\"p\">+</span><span class=\"nb\">pa</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above command is <strong>mapped in insert mode</strong> as the mode is set to insert
  at the command's beginning. The CTRL+c keys are mapped to first Escaping out of
  insert mode and then pasting from the &quot;+ register which is the device clipboard
  here. In the end, we again get back to insert mode just to the right of the current
  cursor position. So, that is how we can map the keys to do the task and save a bit
  of time.</p>\n<h2>Adding Keymaps in vimrc file</h2>\n<p>So you can now map in any
  available modes in Vim and test it out in te=he current buffer, but if you notice
  when you close the buffer, your custom keymaps will not be functioning as they remain
  in the buffer until you close it. To make it permanent you need to type the exact
  map command to the vimrc file. Just don't add &quot;:&quot; as it is not a command
  mode, it should be only the pure command of the map.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span>nmap
  <span class=\"p\">&lt;</span>C<span class=\"p\">-</span>s<span class=\"p\">&gt;</span>
  :<span class=\"k\">w</span><span class=\"p\">&lt;</span>CR<span class=\"p\">&gt;</span>\nnmap
  <span class=\"p\">&lt;</span>C<span class=\"p\">-</span><span class=\"k\">p</span><span
  class=\"p\">&gt;</span> <span class=\"c\">&quot;+p</span>\nvmap <span class=\"p\">&lt;</span>C<span
  class=\"p\">-</span><span class=\"k\">y</span><span class=\"p\">&gt;</span> <span
  class=\"c\">&quot;+y</span>\n</pre></div>\n\n</pre>\n\n<p>The above commands are
  pasted in the vimrc file. These keymaps once sourced will stay in permanently unless
  you remove them from the file.</p>\n<h2>Keymap Modifications</h2>\n<p>We can also
  add extra arguments to the keymaps such as noremap and remap and others. In remap,
  the keymap is recursive which means the key pair will get mapped to the keys mapped
  already somewhere. The keymap will redefine its map depending on the already existing
  maps.  It is preferred to use noremap as it will not mix up the existing maps and
  have a fresh key binding. The mapping in Vim is quite versatile and there are quite
  a lot of modifications you could do to make your key binding feel and work according
  to you. So for further un depth understanding of modifications of keymap in Vim
  you can check out their documentation  <a href=\"https://vim.fandom.com/wiki/Mapping_keys_in_Vim_-_Tutorial_(Part_1)\">here</a>.</p>\n<h2>Keymapping
  Usecases</h2>\n<p>Now, it's entirely up to you to make keymaps according to your
  preference and choice. You can use your creativity and knowledge to make powerful
  and efficient keymaps.</p>\n<p>For some insights, I'd like to give in some cool
  and productive keymaps which I personally use very frequently in coding a particular
  programming language.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span>nmap
  <span class=\"p\">&lt;</span>C<span class=\"p\">-</span>s<span class=\"p\">&gt;</span>
  :<span class=\"k\">w</span><span class=\"p\">&lt;</span>CR<span class=\"p\">&gt;</span>\n<span
  class=\"nb\">nnoremap</span> <span class=\"k\">py</span> :<span class=\"p\">!</span>python
  %<span class=\"p\">&lt;</span>CR<span class=\"p\">&gt;</span>                          <span
  class=\"c\">&quot; Python run macro</span>\n<span class=\"nb\">nnoremap</span> cpp
  :<span class=\"p\">!</span><span class=\"k\">c</span><span class=\"p\">++</span>
  % <span class=\"p\">-</span><span class=\"k\">o</span> %:<span class=\"k\">r</span>
  &amp;&amp; ./%:<span class=\"k\">r</span> <span class=\"p\">&lt;</span>CR<span class=\"p\">&gt;</span>
  \          <span class=\"c\">&quot; C++ compile and run macro</span>\n<span class=\"nb\">nnoremap</span>
  <span class=\"k\">sh</span> :<span class=\"p\">!</span>chmod <span class=\"p\">+</span><span
  class=\"k\">x</span> % &amp;&amp; source %<span class=\"p\">&lt;</span>CR<span class=\"p\">&gt;</span>
  \           <span class=\"c\">&quot; Bash sourcing macro </span>\n<span class=\"nb\">nnoremap</span>
  <span class=\"k\">c</span> :<span class=\"p\">!</span>clang % <span class=\"p\">-</span><span
  class=\"k\">o</span> %:<span class=\"k\">r</span> &amp;&amp; ./%:<span class=\"k\">r</span>
  <span class=\"p\">&lt;</span>CR<span class=\"p\">&gt;</span>           <span class=\"c\">&quot;
  C compile and run macro </span>\n<span class=\"nb\">nnoremap</span> jv :<span class=\"p\">!</span>javac
  % &amp;&amp; java %:<span class=\"k\">r</span> <span class=\"p\">&lt;</span>CR<span
  class=\"p\">&gt;</span>              <span class=\"c\">&quot; Java compile and run
  macro</span>\n</pre></div>\n\n</pre>\n\n<p>The above mapping will run the file or
  compile and run the file you are currently editing in Vim. That is just an overview
  of how you can use key mappings in Vim for any programming language or tool. Vim
  allows you to modify and make the keys do the heavy lifting very easily and save
  a lot of time and frustration. This is how you implement DRY(don't repeat yourself)
  perfectly. I personally think this is just perfect to do certain things which are
  quite common and sometimes daunts you to do this again. So just have a look and
  play around with Vim keymaps, it is the power with which one can excel in saving
  time and energy. This is just flawless. OK, That was too much from my side.</p>\n<h2>Keymapping
  and Macros</h2>\n<p>Some Keymappings provided above are quite interesting for a
  person trying to understand macros and key binding. A macro is a bunch of commands
  packed in together to access it very efficiently, Not the exact definition, but
  still, it just means to simplify things and saves time. Keymapping is the same thing
  but to map with the keys.</p>\n<p>From some snippets and explanations of VIm keymaps,
  it must be easier now for a beginner to understand Keymappings in Vim and how to
  customize accordingly.</p>\n<p>I hope it helped, Thank you for listening to my understanding
  of Vim keymapping. Below are some of my keymaps currently in vim for windows. Some
  of them were already shown in the code snippets.\n<img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1623654888460/pBfJO2jkZ.png\"
  alt=\"vimkmap.png\" /></p>\n<p>So, Keymaps are quite powerful and can be a bit difficult
  to set up in a single try. Keep experimenting with keymaps and make Vim the powerhouse
  of productivity and customization. Happy Viming :)</p>\n"
cover: ''
date: 2021-06-14
datetime: 2021-06-14 00:00:00+00:00
description: Keymapping or Key binding is a process in which a user or a system can
  assign certain keys or commands to perform a particular task or commands. This can
  be qui
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-06-14-Vim-Keymapping.md
html: "<h2>Introduction</h2>\n<p>Keymapping or Key binding is a process in which a
  user or a system can assign certain keys or commands to perform a particular task
  or commands. This can be quite useful in Vim as well as in other text editors to
  avoid some repetitive tasks and save time. In Vim this can be a great power to have
  for programmers as it can be really very flexible to set up for any programming
  language and it's\nenvironment.</p>\n<p>In Vim, you can map a particular keystroke
  combination to trigger a command or any operation. You can also map a key in place
  of a weird key combination as well. You can map a key to integrate several commands
  and make it run at a single key or without much of a hassle. Key Mapping is quite
  powerful in Vim, only your creativity and imagination are the limits here. Vim can
  really shine in such aspects where the user has the freedom to tailor the editor
  and his development environment as per his/her needs.</p>\n<h2>Understanding the
  structure of Key binding</h2>\n<p>You can basically map a key depending on the current
  mode you are in. So that means you can have the same key mapped to different commands
  depending on the mode. That is really flexible and powerful. Vim allows you to basically
  map in almost every mode such as normal, insert, visual, command, and any other
  existing modes. Let us see what a basic key mapping looks like.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>{mode}{attribute}map {key}
  {command}\n</pre></div>\n\n</pre>\n\n<p>The key mapping is broken down by adding
  the mode to map the key, an optional attribute to change/modify the way the mapping
  behaves and we finally have the key and the command pair. If you want to test a
  map you can temporarily map in the command mode using the syntax for the map command
  or if you want a permanent mapping, you can add the mappings in your vimrc file.</p>\n<p>For
  a basic example let us map c to paste from the clipboard.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"p\">:</span>map <span class=\"k\">c</span> <span class=\"c\">&quot;+p&lt;CR&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above command maps the &quot;c&quot; key in the <strong>normal/visual mode</strong>
  by default to paste the contents from the clipboard. It is just an example of how
  a basic key mapping or binding can be used to save your time and efforts. This is
  a very small example to demonstrate the concept of mapping in Vim but it scales
  quite fantastically and smoothly.</p>\n<p>Now let's see how to view the mapping
  inside of the current vim buffer. We can simply type map, imap, or nmap in the command
  mode to view the mappings in the respective modes. This will display the maps which
  are currently present inside of the vim buffer, some might be temporary and most
  of them will be permanent.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>:map \n</pre></div>\n\n</pre>\n\n<p>If
  you type in map being in the command mode, it will list out the maps. If you want
  to stop seeing the list just press q, and you will be back in the editor. You can
  also view the maps in the normal, visual, insert, or any other modes by prefixing
  n, v, i, or other modes with map.</p>\n<h2>Key Mapping in Modes</h2>\n<p>To map
  a key binding explicitly in the normal mode or any other modes, we have to prefix
  the word such as n for normal, v for visual or i for insert, and so on. The map
  will only work in the provided mode so the same keys may have several different
  maps depending on the current mode, but that can create a bit of confusion so avoid
  doing that. Let's map a key in the normal mode for example,</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"p\">:</span>nmap <span class=\"p\">&lt;</span>C<span class=\"p\">-</span>s<span
  class=\"p\">&gt;</span> :<span class=\"k\">w</span><span class=\"p\">&lt;</span>CR<span
  class=\"p\">&gt;</span>    \n</pre></div>\n\n</pre>\n\n<p>This is pretty bad for
  Vim beginners as it will spoil the real Vim experience of saving with :w, If you
  feel comfortable using Ctrl + s to save a file, the above map will be perfect. It
  will map the key Ctrl + s by pressing :w and then pressing enter to save changes
  to the file.</p>\n<p>We can also map certain non-alphabetical keys such as CTRL,
  ALT, and others, it will help you in saving time to leave a mode and enter the mode
  again to edit the text. The following is a perfect example of such.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"p\">:</span>imap <span class=\"p\">&lt;</span>C<span class=\"p\">-</span><span
  class=\"k\">c</span><span class=\"p\">&gt;</span> <span class=\"p\">&lt;</span>ESC<span
  class=\"p\">&gt;</span>&quot;<span class=\"p\">+</span><span class=\"nb\">pa</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above command is <strong>mapped in insert mode</strong> as the mode is set to insert
  at the command's beginning. The CTRL+c keys are mapped to first Escaping out of
  insert mode and then pasting from the &quot;+ register which is the device clipboard
  here. In the end, we again get back to insert mode just to the right of the current
  cursor position. So, that is how we can map the keys to do the task and save a bit
  of time.</p>\n<h2>Adding Keymaps in vimrc file</h2>\n<p>So you can now map in any
  available modes in Vim and test it out in te=he current buffer, but if you notice
  when you close the buffer, your custom keymaps will not be functioning as they remain
  in the buffer until you close it. To make it permanent you need to type the exact
  map command to the vimrc file. Just don't add &quot;:&quot; as it is not a command
  mode, it should be only the pure command of the map.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span>nmap
  <span class=\"p\">&lt;</span>C<span class=\"p\">-</span>s<span class=\"p\">&gt;</span>
  :<span class=\"k\">w</span><span class=\"p\">&lt;</span>CR<span class=\"p\">&gt;</span>\nnmap
  <span class=\"p\">&lt;</span>C<span class=\"p\">-</span><span class=\"k\">p</span><span
  class=\"p\">&gt;</span> <span class=\"c\">&quot;+p</span>\nvmap <span class=\"p\">&lt;</span>C<span
  class=\"p\">-</span><span class=\"k\">y</span><span class=\"p\">&gt;</span> <span
  class=\"c\">&quot;+y</span>\n</pre></div>\n\n</pre>\n\n<p>The above commands are
  pasted in the vimrc file. These keymaps once sourced will stay in permanently unless
  you remove them from the file.</p>\n<h2>Keymap Modifications</h2>\n<p>We can also
  add extra arguments to the keymaps such as noremap and remap and others. In remap,
  the keymap is recursive which means the key pair will get mapped to the keys mapped
  already somewhere. The keymap will redefine its map depending on the already existing
  maps.  It is preferred to use noremap as it will not mix up the existing maps and
  have a fresh key binding. The mapping in Vim is quite versatile and there are quite
  a lot of modifications you could do to make your key binding feel and work according
  to you. So for further un depth understanding of modifications of keymap in Vim
  you can check out their documentation  <a href=\"https://vim.fandom.com/wiki/Mapping_keys_in_Vim_-_Tutorial_(Part_1)\">here</a>.</p>\n<h2>Keymapping
  Usecases</h2>\n<p>Now, it's entirely up to you to make keymaps according to your
  preference and choice. You can use your creativity and knowledge to make powerful
  and efficient keymaps.</p>\n<p>For some insights, I'd like to give in some cool
  and productive keymaps which I personally use very frequently in coding a particular
  programming language.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span>nmap
  <span class=\"p\">&lt;</span>C<span class=\"p\">-</span>s<span class=\"p\">&gt;</span>
  :<span class=\"k\">w</span><span class=\"p\">&lt;</span>CR<span class=\"p\">&gt;</span>\n<span
  class=\"nb\">nnoremap</span> <span class=\"k\">py</span> :<span class=\"p\">!</span>python
  %<span class=\"p\">&lt;</span>CR<span class=\"p\">&gt;</span>                          <span
  class=\"c\">&quot; Python run macro</span>\n<span class=\"nb\">nnoremap</span> cpp
  :<span class=\"p\">!</span><span class=\"k\">c</span><span class=\"p\">++</span>
  % <span class=\"p\">-</span><span class=\"k\">o</span> %:<span class=\"k\">r</span>
  &amp;&amp; ./%:<span class=\"k\">r</span> <span class=\"p\">&lt;</span>CR<span class=\"p\">&gt;</span>
  \          <span class=\"c\">&quot; C++ compile and run macro</span>\n<span class=\"nb\">nnoremap</span>
  <span class=\"k\">sh</span> :<span class=\"p\">!</span>chmod <span class=\"p\">+</span><span
  class=\"k\">x</span> % &amp;&amp; source %<span class=\"p\">&lt;</span>CR<span class=\"p\">&gt;</span>
  \           <span class=\"c\">&quot; Bash sourcing macro </span>\n<span class=\"nb\">nnoremap</span>
  <span class=\"k\">c</span> :<span class=\"p\">!</span>clang % <span class=\"p\">-</span><span
  class=\"k\">o</span> %:<span class=\"k\">r</span> &amp;&amp; ./%:<span class=\"k\">r</span>
  <span class=\"p\">&lt;</span>CR<span class=\"p\">&gt;</span>           <span class=\"c\">&quot;
  C compile and run macro </span>\n<span class=\"nb\">nnoremap</span> jv :<span class=\"p\">!</span>javac
  % &amp;&amp; java %:<span class=\"k\">r</span> <span class=\"p\">&lt;</span>CR<span
  class=\"p\">&gt;</span>              <span class=\"c\">&quot; Java compile and run
  macro</span>\n</pre></div>\n\n</pre>\n\n<p>The above mapping will run the file or
  compile and run the file you are currently editing in Vim. That is just an overview
  of how you can use key mappings in Vim for any programming language or tool. Vim
  allows you to modify and make the keys do the heavy lifting very easily and save
  a lot of time and frustration. This is how you implement DRY(don't repeat yourself)
  perfectly. I personally think this is just perfect to do certain things which are
  quite common and sometimes daunts you to do this again. So just have a look and
  play around with Vim keymaps, it is the power with which one can excel in saving
  time and energy. This is just flawless. OK, That was too much from my side.</p>\n<h2>Keymapping
  and Macros</h2>\n<p>Some Keymappings provided above are quite interesting for a
  person trying to understand macros and key binding. A macro is a bunch of commands
  packed in together to access it very efficiently, Not the exact definition, but
  still, it just means to simplify things and saves time. Keymapping is the same thing
  but to map with the keys.</p>\n<p>From some snippets and explanations of VIm keymaps,
  it must be easier now for a beginner to understand Keymappings in Vim and how to
  customize accordingly.</p>\n<p>I hope it helped, Thank you for listening to my understanding
  of Vim keymapping. Below are some of my keymaps currently in vim for windows. Some
  of them were already shown in the code snippets.\n<img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1623654888460/pBfJO2jkZ.png\"
  alt=\"vimkmap.png\" /></p>\n<p>So, Keymaps are quite powerful and can be a bit difficult
  to set up in a single try. Keep experimenting with keymaps and make Vim the powerhouse
  of productivity and customization. Happy Viming :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/vim-keymapping-guide.webp
long_description: 'Keymapping or Key binding is a process in which a user or a system
  can assign certain keys or commands to perform a particular task or commands. This
  can be quite useful in Vim as well as in other text editors to avoid some repetitive
  tasks and save '
now: 2025-05-01 18:11:33.313055
path: blog/posts/2021-06-14-Vim-Keymapping.md
prevnext: null
slug: vim-keymaps
status: published
subtitle: A comprehensive guide for getting started with key bindings in Vim
tags:
- vim
templateKey: blog-post
title: 'Vim: Keymapping Guide'
today: 2025-05-01
---

## Introduction

Keymapping or Key binding is a process in which a user or a system can assign certain keys or commands to perform a particular task or commands. This can be quite useful in Vim as well as in other text editors to avoid some repetitive tasks and save time. In Vim this can be a great power to have for programmers as it can be really very flexible to set up for any programming language and it's 
 environment.

In Vim, you can map a particular keystroke combination to trigger a command or any operation. You can also map a key in place of a weird key combination as well. You can map a key to integrate several commands and make it run at a single key or without much of a hassle. Key Mapping is quite powerful in Vim, only your creativity and imagination are the limits here. Vim can really shine in such aspects where the user has the freedom to tailor the editor and his development environment as per his/her needs.

## Understanding the structure of Key binding
You can basically map a key depending on the current mode you are in. So that means you can have the same key mapped to different commands depending on the mode. That is really flexible and powerful. Vim allows you to basically map in almost every mode such as normal, insert, visual, command, and any other existing modes. Let us see what a basic key mapping looks like.


```
{mode}{attribute}map {key} {command}
```
The key mapping is broken down by adding the mode to map the key, an optional attribute to change/modify the way the mapping behaves and we finally have the key and the command pair. If you want to test a map you can temporarily map in the command mode using the syntax for the map command or if you want a permanent mapping, you can add the mappings in your vimrc file.

For a basic example let us map c to paste from the clipboard.

```vim
:map c "+p<CR>
```

The above command maps the "c" key in the **normal/visual mode** by default to paste the contents from the clipboard. It is just an example of how a basic key mapping or binding can be used to save your time and efforts. This is a very small example to demonstrate the concept of mapping in Vim but it scales quite fantastically and smoothly.

Now let's see how to view the mapping inside of the current vim buffer. We can simply type map, imap, or nmap in the command mode to view the mappings in the respective modes. This will display the maps which are currently present inside of the vim buffer, some might be temporary and most of them will be permanent. 
```
:map 
```
If you type in map being in the command mode, it will list out the maps. If you want to stop seeing the list just press q, and you will be back in the editor. You can also view the maps in the normal, visual, insert, or any other modes by prefixing n, v, i, or other modes with map.

## Key Mapping in Modes
To map a key binding explicitly in the normal mode or any other modes, we have to prefix the word such as n for normal, v for visual or i for insert, and so on. The map will only work in the provided mode so the same keys may have several different maps depending on the current mode, but that can create a bit of confusion so avoid doing that. Let's map a key in the normal mode for example,
```vim
:nmap <C-s> :w<CR>    
```

This is pretty bad for Vim beginners as it will spoil the real Vim experience of saving with :w, If you feel comfortable using Ctrl + s to save a file, the above map will be perfect. It will map the key Ctrl + s by pressing :w and then pressing enter to save changes to the file. 

We can also map certain non-alphabetical keys such as CTRL, ALT, and others, it will help you in saving time to leave a mode and enter the mode again to edit the text. The following is a perfect example of such.

```vim
:imap <C-c> <ESC>"+pa
```
The above command is **mapped in insert mode** as the mode is set to insert at the command's beginning. The CTRL+c keys are mapped to first Escaping out of insert mode and then pasting from the "+ register which is the device clipboard here. In the end, we again get back to insert mode just to the right of the current cursor position. So, that is how we can map the keys to do the task and save a bit of time. 

## Adding Keymaps in vimrc file
So you can now map in any available modes in Vim and test it out in te=he current buffer, but if you notice when you close the buffer, your custom keymaps will not be functioning as they remain in the buffer until you close it. To make it permanent you need to type the exact map command to the vimrc file. Just don't add ":" as it is not a command mode, it should be only the pure command of the map.

```vim
nmap <C-s> :w<CR>
nmap <C-p> "+p
vmap <C-y> "+y
``` 
The above commands are pasted in the vimrc file. These keymaps once sourced will stay in permanently unless you remove them from the file.

## Keymap Modifications
 We can also add extra arguments to the keymaps such as noremap and remap and others. In remap, the keymap is recursive which means the key pair will get mapped to the keys mapped already somewhere. The keymap will redefine its map depending on the already existing maps.  It is preferred to use noremap as it will not mix up the existing maps and have a fresh key binding. The mapping in Vim is quite versatile and there are quite a lot of modifications you could do to make your key binding feel and work according to you. So for further un depth understanding of modifications of keymap in Vim you can check out their documentation  [here](https://vim.fandom.com/wiki/Mapping_keys_in_Vim_-_Tutorial_(Part_1)).


## Keymapping Usecases
Now, it's entirely up to you to make keymaps according to your preference and choice. You can use your creativity and knowledge to make powerful and efficient keymaps.

For some insights, I'd like to give in some cool and productive keymaps which I personally use very frequently in coding a particular programming language. 

```vim
nmap <C-s> :w<CR>
nnoremap py :!python %<CR>                          " Python run macro
nnoremap cpp :!c++ % -o %:r && ./%:r <CR>           " C++ compile and run macro
nnoremap sh :!chmod +x % && source %<CR>            " Bash sourcing macro 
nnoremap c :!clang % -o %:r && ./%:r <CR>           " C compile and run macro 
nnoremap jv :!javac % && java %:r <CR>              " Java compile and run macro
```


The above mapping will run the file or compile and run the file you are currently editing in Vim. That is just an overview of how you can use key mappings in Vim for any programming language or tool. Vim allows you to modify and make the keys do the heavy lifting very easily and save a lot of time and frustration. This is how you implement DRY(don't repeat yourself) perfectly. I personally think this is just perfect to do certain things which are quite common and sometimes daunts you to do this again. So just have a look and play around with Vim keymaps, it is the power with which one can excel in saving time and energy. This is just flawless. OK, That was too much from my side. 

## Keymapping and Macros
Some Keymappings provided above are quite interesting for a person trying to understand macros and key binding. A macro is a bunch of commands packed in together to access it very efficiently, Not the exact definition, but still, it just means to simplify things and saves time. Keymapping is the same thing but to map with the keys. 

From some snippets and explanations of VIm keymaps, it must be easier now for a beginner to understand Keymappings in Vim and how to customize accordingly.

I hope it helped, Thank you for listening to my understanding of Vim keymapping. Below are some of my keymaps currently in vim for windows. Some of them were already shown in the code snippets.
![vimkmap.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1623654888460/pBfJO2jkZ.png)

 
So, Keymaps are quite powerful and can be a bit difficult to set up in a single try. Keep experimenting with keymaps and make Vim the powerhouse of productivity and customization. Happy Viming :)