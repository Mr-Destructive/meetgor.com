---
article_html: "<h2>Introduction</h2>\n<p>We as programmers always fiddle with commenting
  out code for code testing, documenting the function of code, and most importantly
  debugging. So you can't wait to comment on a large chunk of code manually, as it
  is quite a tedious thing to do. Let's do it effectively in Vim.</p>\n<p>In this
  part of the series, I'll cover how to comment/uncomment chunks/blocks of code effectively
  in Vim. We will see and use some commands, keybindings for doing so, and also we
  would add certain components to our vimrc file as well to design some custom key
  mappings.  Let's get faster with Vim.</p>\n<h2>How to comment multiple lines effectively</h2>\n<p>To
  comment on multiple lines of code, we can use the Visual Block mode to select the
  lines, and then after entering into insert mode, we can comment a single line and
  it would be reflected on all the selected lines.</p>\n<ol>\n<li>\n<p>Press <code>CTRL+V</code>
  and Select the line using j and k</p>\n</li>\n<li>\n<p>After Selecting the lines,
  Press <code>Escape</code></p>\n</li>\n<li>\n<p>Press <code>Shift + I</code>, to
  enter insert mode</p>\n</li>\n<li>\n<p>Enter the comment code (<code>//</code>,
  <code>#</code>, or other)</p>\n</li>\n</ol>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1633518136135/06dfBTq2T.gif\"
  alt=\"vimcoment.gif\" /></p>\n<p>So, using just simple steps you can comment out
  large chunks of code quite easily and effectively. If you are using some other language
  that has multiple characters for commenting like <code>//</code>, <code>- -</code>,
  etc, you can type in any number of characters while being in insert mode after selecting
  the lines.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1633520509953/0q-k2ZHC7.gif\"
  alt=\"vimcppcom.gif\" /></p>\n<p>This might look a bit wired on the first try but
  just try it every day, It is a life-saving and very satisfying experience once applied
  in a real-world scenario.</p>\n<h2>How to uncomment multiple lines effectively</h2>\n<p>Now,
  as we have seen to comment out a large chunk of code, we can even uncomment the
  code very easily. It's even simpler than commenting the code.</p>\n<ol>\n<li>\n<p>Press
  <code>CTRL + V</code> to enter Visual Block mode</p>\n</li>\n<li>\n<p>Select the
  commented characters</p>\n</li>\n<li>\n<p>Press <code>d</code> to delete the comments</p>\n</li>\n<li>\n<p>Press
  <code>Escape</code></p>\n</li>\n</ol>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1633518156818/GJzRPTI3I.gif\"
  alt=\"vimuncoment.gif\" /></p>\n<p>We can simply use the CTRL + V to select the
  comment, and then press d to delete all the comment characters.</p>\n<p><strong>We
  are using the Visual Block mode as we only want the comment to be selected and not
  the entire code associated with the lines.</strong></p>\n<h2>Using Multiline Comments
  for Programming languages</h2>\n<p>Now you might say, why use multiple single-line
  comments when we can use multiline comments in almost all programming languages.
  Well, Of course, you can do that, it's easier for reading the code if syntax highlighting
  is accurate and greys out the commented part. We can simply add those characters
  to the start of the block and at the end of the block.</p>\n<p>But in Vim, we can
  customize that too, just imagine when you just select the chunk/block of code that
  you need to comment out and then simply press a few keystrokes (just 2) and the
  multiline comments are automatically (programmatically) added as per the programming
  language extension of the file.</p>\n<p>Isn't that cool? Well, you just need to
  copy-paste the below code to your Vimrc file and source it and you are good to go.</p>\n<pre
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
  class=\"k\">function</span><span class=\"p\">!</span> Comment<span class=\"p\">()</span>\n
  \   <span class=\"k\">let</span> ext <span class=\"p\">=</span> tolower<span class=\"p\">(</span>expand<span
  class=\"p\">(</span><span class=\"s1\">&#39;%:e&#39;</span><span class=\"p\">))</span>\n
  \   <span class=\"k\">if</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;py&#39;</span>
  \n        <span class=\"k\">let</span> cmt1 <span class=\"p\">=</span> <span class=\"s2\">&quot;&#39;&#39;&#39;&quot;</span>\n\t
  \   <span class=\"k\">let</span> cmt2 <span class=\"p\">=</span> <span class=\"s2\">&quot;&#39;&#39;&#39;&quot;</span>
  \  \n    <span class=\"k\">elseif</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;cpp&#39;</span>
  <span class=\"p\">||</span> ext <span class=\"p\">==</span><span class=\"s1\">&#39;java&#39;</span>
  <span class=\"p\">||</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;css&#39;</span>
  <span class=\"p\">||</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;js&#39;</span>
  <span class=\"p\">||</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;c&#39;</span>
  <span class=\"p\">||</span> ext <span class=\"p\">==</span><span class=\"s1\">&#39;cs&#39;</span>
  <span class=\"p\">||</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;rs&#39;</span>
  <span class=\"p\">||</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;go&#39;</span>\n\t
  \   <span class=\"k\">let</span> cmt1 <span class=\"p\">=</span> <span class=\"s1\">&#39;/*&#39;</span>\n\t
  \   <span class=\"k\">let</span> cmt2 <span class=\"p\">=</span> <span class=\"s1\">&#39;*/&#39;</span>\n
  \   <span class=\"k\">elseif</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;sh&#39;</span>\n\t
  \   <span class=\"k\">let</span> cmt1 <span class=\"p\">=</span> <span class=\"s2\">&quot;:&#39;&quot;</span>\n\t
  \   <span class=\"k\">let</span> cmt2 <span class=\"p\">=</span> <span class=\"s2\">&quot;&#39;&quot;</span>\n
  \   <span class=\"k\">elseif</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;html&#39;</span>\n\t
  \   <span class=\"k\">let</span> cmt1 <span class=\"p\">=</span> <span class=\"s2\">&quot;&lt;!--&quot;</span>\n\t
  \   <span class=\"k\">let</span> cmt2 <span class=\"p\">=</span> <span class=\"s2\">&quot;--&gt;&quot;</span>\n
  \   <span class=\"k\">elseif</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;hs&#39;</span>\n\t
  \   <span class=\"k\">let</span> cmt1 <span class=\"p\">=</span> <span class=\"s2\">&quot;{-&quot;</span>\n\t
  \   <span class=\"k\">let</span> cmt2 <span class=\"p\">=</span> <span class=\"s2\">&quot;-}&quot;</span>\n
  \   <span class=\"k\">elseif</span> ext <span class=\"p\">==</span> <span class=\"s2\">&quot;rb&quot;</span>\n\t
  \   <span class=\"k\">let</span> cmt1 <span class=\"p\">=</span> <span class=\"s2\">&quot;=begin&quot;</span>\n\t
  \   <span class=\"k\">let</span> cmt2 <span class=\"p\">=</span> <span class=\"s2\">&quot;=end&quot;</span>\n
  \   <span class=\"k\">endif</span>\n    exe line<span class=\"p\">(</span><span
  class=\"s2\">&quot;&#39;&lt;&quot;</span><span class=\"p\">)</span>.<span class=\"s2\">&quot;normal
  O&quot;</span>. cmt1 <span class=\"p\">|</span> exe line<span class=\"p\">(</span><span
  class=\"s2\">&quot;&#39;&gt;&quot;</span><span class=\"p\">)</span>.<span class=\"s2\">&quot;normal
  o&quot;</span>. cmt2 \n<span class=\"k\">endfunction</span>\n\n<span class=\"k\">function</span><span
  class=\"p\">!</span> UnComment<span class=\"p\">()</span>\n    exe line<span class=\"p\">(</span><span
  class=\"s2\">&quot;&#39;&lt;&quot;</span><span class=\"p\">)</span>.<span class=\"s2\">&quot;normal
  dd&quot;</span> <span class=\"p\">|</span> exe line<span class=\"p\">(</span><span
  class=\"s2\">&quot;&#39;&gt;&quot;</span><span class=\"p\">)</span>.<span class=\"s2\">&quot;normal
  dd&quot;</span>   \n<span class=\"k\">endfunction</span>\n\n\n<span class=\"nb\">vnoremap</span>
  <span class=\"p\">,</span><span class=\"k\">m</span> :<span class=\"p\">&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;</span><span
  class=\"k\">call</span> Comment<span class=\"p\">()&lt;</span>CR<span class=\"p\">&gt;</span>\n<span
  class=\"nb\">vnoremap</span> <span class=\"k\">m</span><span class=\"p\">,</span>
  :<span class=\"p\">&lt;</span><span class=\"k\">c</span><span class=\"p\">-</span><span
  class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span class=\"k\">c</span><span
  class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;</span><span
  class=\"k\">call</span> UnComment<span class=\"p\">()&lt;</span>CR<span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>The
  below screencast is an example of <code>HTML</code> snippet in a file that is getting
  commented using mapping with the keys <code>,m</code> you can put any other keybinding
  you like.\n<img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1633595891674/hbhrbtRHd.gif\"
  alt=\"htmcm.gif\" /></p>\n<hr />\n<p>Similarly for the next screencast is of an
  <code>Javascript</code> snippet in a file which is getting commented using a mapping
  with the keys <code>,m</code> and uncommented using <code>m,</code> again you can
  put any other keybinding you like.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1633595919104/xGTh5ztWu.gif\"
  alt=\"jscom.gif\" /></p>\n<hr />\n<p>The following screencast is of a shell script(BASH)
  snippet.\n<img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1633596156121/tbGHQBSSA.gif\"
  alt=\"shcom.gif\" /></p>\n<hr />\n<h3>Multiline Comments in various Programming
  Languages:</h3>\n<h4>1. C / C++ / Java / Javascript / CSS / C# / Rust / Go / PHP
  / Swift / Dart / Kotlin</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>/*\n*/\n</pre></div>\n\n</pre>\n\n<h4>2.
  Python</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n\n<div class=\"highlight\"><pre><span></span>&#39;&#39;&#39;\n&#39;&#39;&#39;\n</pre></div>\n\n</pre>\n\n<p>You
  can even use <code>&quot;&quot;&quot;</code> double quotes instead of single quotes</p>\n<h4>3.
  BASH (Shell Scripting)</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>: &#39;\n&#39;\n</pre></div>\n\n</pre>\n\n<p>You
  can even use <code>: &quot;</code> and <code>&quot;</code> double quotes instead
  of single quotes</p>\n<h4>4. Haskell</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>{-\n-}\n</pre></div>\n\n</pre>\n\n<h4>5.
  Ruby</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n\n<div class=\"highlight\"><pre><span></span>=begin\n=end\n</pre></div>\n\n</pre>\n\n<h4>6.
  HTML</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n\n<div class=\"highlight\"><pre><span></span>&lt;!--\n--&gt;\n</pre></div>\n\n</pre>\n\n<h4>7.
  Julia</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n\n<div class=\"highlight\"><pre><span></span>#=\n=#\n</pre></div>\n\n</pre>\n\n<h3>Understanding
  the Commands / Keymapping</h3>\n<p><strong>NOTE : You need to go from the top to
  bottom while commenting on the block of code, otherwise, it would be a mismatch
  in commenting for specific language syntax. While uncommenting the order doesn't
  matter.</strong></p>\n<h4>Getting the extension (filetype)</h4>\n<p>In Vim, we can
  get the file extension i.e. we can get the programming language associated with
  the current file. To do that we can use, <code>expand('%:e')</code>.</p>\n<p>This
  will give us the file extension of the current file. Just for simplicity, <code>%</code>
  means the current file, added to it is <code>:e</code> for excluding the filename
  and keeping the extension. We convert the extension into lowercase just for keeping
  things safe and programmatic and store it in a variable <code>ext</code>.</p>\n<h4>Checking
  for programming language</h4>\n<p>We then can then use an if-else ladder to check
  for the programming languages and assign two variables <code>cmt1</code> for the
  initial characters in the multiline comment and <code>cmt2</code> for enclosing
  the comment.</p>\n<h4>Typing in the characters</h4>\n<p>We can use the function
  <code>line(&quot;'&lt;&quot;)</code> to get the line number of the previous visual
  selection. Similarly, <code>line(&quot;'&gt;&quot;)</code> for the ending line.
  We are using the <code>exe</code> command to execute the function <code>line</code>
  and so we have to use a concatenation of the commands even to write the raw commands
  like <code>i</code> to insert mode, <code>o</code> to insert mode but a line below
  the cursor. So, we use <code>normal</code> command for that. This command indicates
  the interpreter to execute the following commands from the normal mode.</p>\n<p>We
  have to enclose the <code>normal</code> command in double-quotes/single quotes.
  We can simply use the variable again with concatenation.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span>exe
  line<span class=\"p\">(</span><span class=\"s2\">&quot;&#39;&gt;&quot;</span><span
  class=\"p\">)</span>.<span class=\"s2\">&quot;normal o&quot;</span>. cmt2 \n</pre></div>\n\n</pre>\n\n<p>The
  above command will fetch the last line's number of the previous visual selection
  followed by entering <code>o</code> from the normal mode and concatenated with the
  value of the variable <code>cmt2</code> which we have already initialized in the
  <code>Comment</code> function. We are using <code>|</code> for running multiple
  commands as we also need to include the comment at the beginning of the visual selection.</p>\n<p>For
  uncommenting the code, we are simply deleting the entire first and the last line
  in the visual selection. For that, we have used <code>dd</code> from the normal
  mode.</p>\n<h3>Conclusion</h3>\n<p>So, from the following type of tutorial, we were
  able to set up our Vim editor for efficient code commenting/ uncommenting using
  some commands, key shortcuts, and configuring the vimrc for making custom keymappings.
  We were also able to understand the multiline comments in various programming languages
  and use them in Vim very effectively with a simple addon to the config vimrc file.</p>\n<p>Thank
  you for reading, hope you found this article helpful. If you have any queries or
  wanna add multiline comments for some more programming languages please let me know
  in the comments or contact section.</p>\n<p>Happy Coding :)</p>\n<h3>References</h3>\n<ul>\n<li>\n<p><a
  href=\"https://stackoverflow.com/questions/1676632/whats-a-quick-way-to-comment-uncomment-lines-in-vim/1676690\">StackOverflow
  - Commenting lines in Vim </a></p>\n</li>\n<li>\n<p><a href=\"https://dev.to/grepliz/3-ways-to-comment-out-blocks-of-code-in-vi-6j4\">Liz
  Lam - 3 ways to comment code in Vim</a></p>\n</li>\n<li>\n<p><a href=\"https://vi.stackexchange.com/questions/9644/how-to-use-a-variable-in-the-expression-of-a-normal-command\">StackExchange
  - Use variable in normal command</a></p>\n</li>\n</ul>\n"
cover: ''
date: 2021-10-07
datetime: 2021-10-07 00:00:00+00:00
description: 'We as programmers always fiddle with commenting out code for code testing,
  documenting the function of code, and most importantly debugging. So you can In
  this '
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-10-07-Vim-P-comments.md
html: "<h2>Introduction</h2>\n<p>We as programmers always fiddle with commenting out
  code for code testing, documenting the function of code, and most importantly debugging.
  So you can't wait to comment on a large chunk of code manually, as it is quite a
  tedious thing to do. Let's do it effectively in Vim.</p>\n<p>In this part of the
  series, I'll cover how to comment/uncomment chunks/blocks of code effectively in
  Vim. We will see and use some commands, keybindings for doing so, and also we would
  add certain components to our vimrc file as well to design some custom key mappings.
  \ Let's get faster with Vim.</p>\n<h2>How to comment multiple lines effectively</h2>\n<p>To
  comment on multiple lines of code, we can use the Visual Block mode to select the
  lines, and then after entering into insert mode, we can comment a single line and
  it would be reflected on all the selected lines.</p>\n<ol>\n<li>\n<p>Press <code>CTRL+V</code>
  and Select the line using j and k</p>\n</li>\n<li>\n<p>After Selecting the lines,
  Press <code>Escape</code></p>\n</li>\n<li>\n<p>Press <code>Shift + I</code>, to
  enter insert mode</p>\n</li>\n<li>\n<p>Enter the comment code (<code>//</code>,
  <code>#</code>, or other)</p>\n</li>\n</ol>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1633518136135/06dfBTq2T.gif\"
  alt=\"vimcoment.gif\" /></p>\n<p>So, using just simple steps you can comment out
  large chunks of code quite easily and effectively. If you are using some other language
  that has multiple characters for commenting like <code>//</code>, <code>- -</code>,
  etc, you can type in any number of characters while being in insert mode after selecting
  the lines.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1633520509953/0q-k2ZHC7.gif\"
  alt=\"vimcppcom.gif\" /></p>\n<p>This might look a bit wired on the first try but
  just try it every day, It is a life-saving and very satisfying experience once applied
  in a real-world scenario.</p>\n<h2>How to uncomment multiple lines effectively</h2>\n<p>Now,
  as we have seen to comment out a large chunk of code, we can even uncomment the
  code very easily. It's even simpler than commenting the code.</p>\n<ol>\n<li>\n<p>Press
  <code>CTRL + V</code> to enter Visual Block mode</p>\n</li>\n<li>\n<p>Select the
  commented characters</p>\n</li>\n<li>\n<p>Press <code>d</code> to delete the comments</p>\n</li>\n<li>\n<p>Press
  <code>Escape</code></p>\n</li>\n</ol>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1633518156818/GJzRPTI3I.gif\"
  alt=\"vimuncoment.gif\" /></p>\n<p>We can simply use the CTRL + V to select the
  comment, and then press d to delete all the comment characters.</p>\n<p><strong>We
  are using the Visual Block mode as we only want the comment to be selected and not
  the entire code associated with the lines.</strong></p>\n<h2>Using Multiline Comments
  for Programming languages</h2>\n<p>Now you might say, why use multiple single-line
  comments when we can use multiline comments in almost all programming languages.
  Well, Of course, you can do that, it's easier for reading the code if syntax highlighting
  is accurate and greys out the commented part. We can simply add those characters
  to the start of the block and at the end of the block.</p>\n<p>But in Vim, we can
  customize that too, just imagine when you just select the chunk/block of code that
  you need to comment out and then simply press a few keystrokes (just 2) and the
  multiline comments are automatically (programmatically) added as per the programming
  language extension of the file.</p>\n<p>Isn't that cool? Well, you just need to
  copy-paste the below code to your Vimrc file and source it and you are good to go.</p>\n<pre
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
  class=\"k\">function</span><span class=\"p\">!</span> Comment<span class=\"p\">()</span>\n
  \   <span class=\"k\">let</span> ext <span class=\"p\">=</span> tolower<span class=\"p\">(</span>expand<span
  class=\"p\">(</span><span class=\"s1\">&#39;%:e&#39;</span><span class=\"p\">))</span>\n
  \   <span class=\"k\">if</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;py&#39;</span>
  \n        <span class=\"k\">let</span> cmt1 <span class=\"p\">=</span> <span class=\"s2\">&quot;&#39;&#39;&#39;&quot;</span>\n\t
  \   <span class=\"k\">let</span> cmt2 <span class=\"p\">=</span> <span class=\"s2\">&quot;&#39;&#39;&#39;&quot;</span>
  \  \n    <span class=\"k\">elseif</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;cpp&#39;</span>
  <span class=\"p\">||</span> ext <span class=\"p\">==</span><span class=\"s1\">&#39;java&#39;</span>
  <span class=\"p\">||</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;css&#39;</span>
  <span class=\"p\">||</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;js&#39;</span>
  <span class=\"p\">||</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;c&#39;</span>
  <span class=\"p\">||</span> ext <span class=\"p\">==</span><span class=\"s1\">&#39;cs&#39;</span>
  <span class=\"p\">||</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;rs&#39;</span>
  <span class=\"p\">||</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;go&#39;</span>\n\t
  \   <span class=\"k\">let</span> cmt1 <span class=\"p\">=</span> <span class=\"s1\">&#39;/*&#39;</span>\n\t
  \   <span class=\"k\">let</span> cmt2 <span class=\"p\">=</span> <span class=\"s1\">&#39;*/&#39;</span>\n
  \   <span class=\"k\">elseif</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;sh&#39;</span>\n\t
  \   <span class=\"k\">let</span> cmt1 <span class=\"p\">=</span> <span class=\"s2\">&quot;:&#39;&quot;</span>\n\t
  \   <span class=\"k\">let</span> cmt2 <span class=\"p\">=</span> <span class=\"s2\">&quot;&#39;&quot;</span>\n
  \   <span class=\"k\">elseif</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;html&#39;</span>\n\t
  \   <span class=\"k\">let</span> cmt1 <span class=\"p\">=</span> <span class=\"s2\">&quot;&lt;!--&quot;</span>\n\t
  \   <span class=\"k\">let</span> cmt2 <span class=\"p\">=</span> <span class=\"s2\">&quot;--&gt;&quot;</span>\n
  \   <span class=\"k\">elseif</span> ext <span class=\"p\">==</span> <span class=\"s1\">&#39;hs&#39;</span>\n\t
  \   <span class=\"k\">let</span> cmt1 <span class=\"p\">=</span> <span class=\"s2\">&quot;{-&quot;</span>\n\t
  \   <span class=\"k\">let</span> cmt2 <span class=\"p\">=</span> <span class=\"s2\">&quot;-}&quot;</span>\n
  \   <span class=\"k\">elseif</span> ext <span class=\"p\">==</span> <span class=\"s2\">&quot;rb&quot;</span>\n\t
  \   <span class=\"k\">let</span> cmt1 <span class=\"p\">=</span> <span class=\"s2\">&quot;=begin&quot;</span>\n\t
  \   <span class=\"k\">let</span> cmt2 <span class=\"p\">=</span> <span class=\"s2\">&quot;=end&quot;</span>\n
  \   <span class=\"k\">endif</span>\n    exe line<span class=\"p\">(</span><span
  class=\"s2\">&quot;&#39;&lt;&quot;</span><span class=\"p\">)</span>.<span class=\"s2\">&quot;normal
  O&quot;</span>. cmt1 <span class=\"p\">|</span> exe line<span class=\"p\">(</span><span
  class=\"s2\">&quot;&#39;&gt;&quot;</span><span class=\"p\">)</span>.<span class=\"s2\">&quot;normal
  o&quot;</span>. cmt2 \n<span class=\"k\">endfunction</span>\n\n<span class=\"k\">function</span><span
  class=\"p\">!</span> UnComment<span class=\"p\">()</span>\n    exe line<span class=\"p\">(</span><span
  class=\"s2\">&quot;&#39;&lt;&quot;</span><span class=\"p\">)</span>.<span class=\"s2\">&quot;normal
  dd&quot;</span> <span class=\"p\">|</span> exe line<span class=\"p\">(</span><span
  class=\"s2\">&quot;&#39;&gt;&quot;</span><span class=\"p\">)</span>.<span class=\"s2\">&quot;normal
  dd&quot;</span>   \n<span class=\"k\">endfunction</span>\n\n\n<span class=\"nb\">vnoremap</span>
  <span class=\"p\">,</span><span class=\"k\">m</span> :<span class=\"p\">&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;</span><span
  class=\"k\">call</span> Comment<span class=\"p\">()&lt;</span>CR<span class=\"p\">&gt;</span>\n<span
  class=\"nb\">vnoremap</span> <span class=\"k\">m</span><span class=\"p\">,</span>
  :<span class=\"p\">&lt;</span><span class=\"k\">c</span><span class=\"p\">-</span><span
  class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span class=\"k\">c</span><span
  class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;&lt;</span><span
  class=\"k\">c</span><span class=\"p\">-</span><span class=\"k\">w</span><span class=\"p\">&gt;</span><span
  class=\"k\">call</span> UnComment<span class=\"p\">()&lt;</span>CR<span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>The
  below screencast is an example of <code>HTML</code> snippet in a file that is getting
  commented using mapping with the keys <code>,m</code> you can put any other keybinding
  you like.\n<img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1633595891674/hbhrbtRHd.gif\"
  alt=\"htmcm.gif\" /></p>\n<hr />\n<p>Similarly for the next screencast is of an
  <code>Javascript</code> snippet in a file which is getting commented using a mapping
  with the keys <code>,m</code> and uncommented using <code>m,</code> again you can
  put any other keybinding you like.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1633595919104/xGTh5ztWu.gif\"
  alt=\"jscom.gif\" /></p>\n<hr />\n<p>The following screencast is of a shell script(BASH)
  snippet.\n<img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1633596156121/tbGHQBSSA.gif\"
  alt=\"shcom.gif\" /></p>\n<hr />\n<h3>Multiline Comments in various Programming
  Languages:</h3>\n<h4>1. C / C++ / Java / Javascript / CSS / C# / Rust / Go / PHP
  / Swift / Dart / Kotlin</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>/*\n*/\n</pre></div>\n\n</pre>\n\n<h4>2.
  Python</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n\n<div class=\"highlight\"><pre><span></span>&#39;&#39;&#39;\n&#39;&#39;&#39;\n</pre></div>\n\n</pre>\n\n<p>You
  can even use <code>&quot;&quot;&quot;</code> double quotes instead of single quotes</p>\n<h4>3.
  BASH (Shell Scripting)</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>: &#39;\n&#39;\n</pre></div>\n\n</pre>\n\n<p>You
  can even use <code>: &quot;</code> and <code>&quot;</code> double quotes instead
  of single quotes</p>\n<h4>4. Haskell</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>{-\n-}\n</pre></div>\n\n</pre>\n\n<h4>5.
  Ruby</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n\n<div class=\"highlight\"><pre><span></span>=begin\n=end\n</pre></div>\n\n</pre>\n\n<h4>6.
  HTML</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n\n<div class=\"highlight\"><pre><span></span>&lt;!--\n--&gt;\n</pre></div>\n\n</pre>\n\n<h4>7.
  Julia</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
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
  \       \n\n<div class=\"highlight\"><pre><span></span>#=\n=#\n</pre></div>\n\n</pre>\n\n<h3>Understanding
  the Commands / Keymapping</h3>\n<p><strong>NOTE : You need to go from the top to
  bottom while commenting on the block of code, otherwise, it would be a mismatch
  in commenting for specific language syntax. While uncommenting the order doesn't
  matter.</strong></p>\n<h4>Getting the extension (filetype)</h4>\n<p>In Vim, we can
  get the file extension i.e. we can get the programming language associated with
  the current file. To do that we can use, <code>expand('%:e')</code>.</p>\n<p>This
  will give us the file extension of the current file. Just for simplicity, <code>%</code>
  means the current file, added to it is <code>:e</code> for excluding the filename
  and keeping the extension. We convert the extension into lowercase just for keeping
  things safe and programmatic and store it in a variable <code>ext</code>.</p>\n<h4>Checking
  for programming language</h4>\n<p>We then can then use an if-else ladder to check
  for the programming languages and assign two variables <code>cmt1</code> for the
  initial characters in the multiline comment and <code>cmt2</code> for enclosing
  the comment.</p>\n<h4>Typing in the characters</h4>\n<p>We can use the function
  <code>line(&quot;'&lt;&quot;)</code> to get the line number of the previous visual
  selection. Similarly, <code>line(&quot;'&gt;&quot;)</code> for the ending line.
  We are using the <code>exe</code> command to execute the function <code>line</code>
  and so we have to use a concatenation of the commands even to write the raw commands
  like <code>i</code> to insert mode, <code>o</code> to insert mode but a line below
  the cursor. So, we use <code>normal</code> command for that. This command indicates
  the interpreter to execute the following commands from the normal mode.</p>\n<p>We
  have to enclose the <code>normal</code> command in double-quotes/single quotes.
  We can simply use the variable again with concatenation.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>vim</div>\n<div class=\"highlight\"><pre><span></span>exe
  line<span class=\"p\">(</span><span class=\"s2\">&quot;&#39;&gt;&quot;</span><span
  class=\"p\">)</span>.<span class=\"s2\">&quot;normal o&quot;</span>. cmt2 \n</pre></div>\n\n</pre>\n\n<p>The
  above command will fetch the last line's number of the previous visual selection
  followed by entering <code>o</code> from the normal mode and concatenated with the
  value of the variable <code>cmt2</code> which we have already initialized in the
  <code>Comment</code> function. We are using <code>|</code> for running multiple
  commands as we also need to include the comment at the beginning of the visual selection.</p>\n<p>For
  uncommenting the code, we are simply deleting the entire first and the last line
  in the visual selection. For that, we have used <code>dd</code> from the normal
  mode.</p>\n<h3>Conclusion</h3>\n<p>So, from the following type of tutorial, we were
  able to set up our Vim editor for efficient code commenting/ uncommenting using
  some commands, key shortcuts, and configuring the vimrc for making custom keymappings.
  We were also able to understand the multiline comments in various programming languages
  and use them in Vim very effectively with a simple addon to the config vimrc file.</p>\n<p>Thank
  you for reading, hope you found this article helpful. If you have any queries or
  wanna add multiline comments for some more programming languages please let me know
  in the comments or contact section.</p>\n<p>Happy Coding :)</p>\n<h3>References</h3>\n<ul>\n<li>\n<p><a
  href=\"https://stackoverflow.com/questions/1676632/whats-a-quick-way-to-comment-uncomment-lines-in-vim/1676690\">StackOverflow
  - Commenting lines in Vim </a></p>\n</li>\n<li>\n<p><a href=\"https://dev.to/grepliz/3-ways-to-comment-out-blocks-of-code-in-vi-6j4\">Liz
  Lam - 3 ways to comment code in Vim</a></p>\n</li>\n<li>\n<p><a href=\"https://vi.stackexchange.com/questions/9644/how-to-use-a-variable-in-the-expression-of-a-normal-command\">StackExchange
  - Use variable in normal command</a></p>\n</li>\n</ul>\n"
image_url: https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643288865/blogmedia/jbs2hcaiplfvwsgrcfgl.png
long_description: We as programmers always fiddle with commenting out code for code
  testing, documenting the function of code, and most importantly debugging. So you
  can In this part of the series, I To comment on multiple lines of code, we can use
  the Visual Block mo
now: 2025-05-01 18:11:33.313530
path: blog/posts/2021-10-07-Vim-P-comments.md
prevnext: null
slug: vim-un-comment-p1
status: published
subtitle: A short guide to comment/uncomment chunks of code effectively in Vim
tags:
- vim
templateKey: blog-post
title: 'Comment/Uncomment Code: Vim for Programmers'
today: 2025-05-01
---

## Introduction

We as programmers always fiddle with commenting out code for code testing, documenting the function of code, and most importantly debugging. So you can't wait to comment on a large chunk of code manually, as it is quite a tedious thing to do. Let's do it effectively in Vim.

In this part of the series, I'll cover how to comment/uncomment chunks/blocks of code effectively in Vim. We will see and use some commands, keybindings for doing so, and also we would add certain components to our vimrc file as well to design some custom key mappings.  Let's get faster with Vim.
  
## How to comment multiple lines effectively

To comment on multiple lines of code, we can use the Visual Block mode to select the lines, and then after entering into insert mode, we can comment a single line and it would be reflected on all the selected lines.

1. Press `CTRL+V` and Select the line using j and k

2. After Selecting the lines, Press `Escape`

3. Press `Shift + I`, to enter insert mode

4. Enter the comment code (`//`, `#`, or other)


![vimcoment.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1633518136135/06dfBTq2T.gif)

So, using just simple steps you can comment out large chunks of code quite easily and effectively. If you are using some other language that has multiple characters for commenting like `//`, `- -`, etc, you can type in any number of characters while being in insert mode after selecting the lines.

 
![vimcppcom.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1633520509953/0q-k2ZHC7.gif)

This might look a bit wired on the first try but just try it every day, It is a life-saving and very satisfying experience once applied in a real-world scenario.


## How to uncomment multiple lines effectively

Now, as we have seen to comment out a large chunk of code, we can even uncomment the code very easily. It's even simpler than commenting the code.

1. Press `CTRL + V` to enter Visual Block mode

2. Select the commented characters

3. Press `d` to delete the comments

4. Press `Escape`

![vimuncoment.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1633518156818/GJzRPTI3I.gif)

We can simply use the CTRL + V to select the comment, and then press d to delete all the comment characters. 

**We are using the Visual Block mode as we only want the comment to be selected and not the entire code associated with the lines.**

## Using Multiline Comments for Programming languages

Now you might say, why use multiple single-line comments when we can use multiline comments in almost all programming languages. Well, Of course, you can do that, it's easier for reading the code if syntax highlighting is accurate and greys out the commented part. We can simply add those characters to the start of the block and at the end of the block.  

But in Vim, we can customize that too, just imagine when you just select the chunk/block of code that you need to comment out and then simply press a few keystrokes (just 2) and the multiline comments are automatically (programmatically) added as per the programming language extension of the file.

Isn't that cool? Well, you just need to copy-paste the below code to your Vimrc file and source it and you are good to go. 

```vim
function! Comment()
    let ext = tolower(expand('%:e'))
    if ext == 'py' 
        let cmt1 = "'''"
	    let cmt2 = "'''"   
    elseif ext == 'cpp' || ext =='java' || ext == 'css' || ext == 'js' || ext == 'c' || ext =='cs' || ext == 'rs' || ext == 'go'
	    let cmt1 = '/*'
	    let cmt2 = '*/'
    elseif ext == 'sh'
	    let cmt1 = ":'"
	    let cmt2 = "'"
    elseif ext == 'html'
	    let cmt1 = "<!--"
	    let cmt2 = "-->"
    elseif ext == 'hs'
	    let cmt1 = "{-"
	    let cmt2 = "-}"
    elseif ext == "rb"
	    let cmt1 = "=begin"
	    let cmt2 = "=end"
    endif
    exe line("'<")."normal O". cmt1 | exe line("'>")."normal o". cmt2 
endfunction

function! UnComment()
    exe line("'<")."normal dd" | exe line("'>")."normal dd"   
endfunction


vnoremap ,m :<c-w><c-w><c-w><c-w><c-w>call Comment()<CR>
vnoremap m, :<c-w><c-w><c-w><c-w><c-w>call UnComment()<CR>

```
The below screencast is an example of `HTML` snippet in a file that is getting commented using mapping with the keys `,m` you can put any other keybinding you like. 
![htmcm.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1633595891674/hbhrbtRHd.gif)

---
Similarly for the next screencast is of an `Javascript` snippet in a file which is getting commented using a mapping with the keys `,m` and uncommented using `m,` again you can put any other keybinding you like. 

![jscom.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1633595919104/xGTh5ztWu.gif)

---
The following screencast is of a shell script(BASH) snippet.
![shcom.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1633596156121/tbGHQBSSA.gif)

---
### Multiline Comments in various Programming Languages:

#### 1. C / C++ / Java / Javascript / CSS / C# / Rust / Go / PHP / Swift / Dart / Kotlin
```
/*
*/
```
#### 2. Python
```
'''
'''
```
You can even use `"""` double quotes instead of single quotes

#### 3. BASH (Shell Scripting)
```
: '
'
```
You can even use `: "` and `"` double quotes instead of single quotes

#### 4. Haskell

```
{-
-}
```
#### 5. Ruby

```
=begin
=end
```

#### 6. HTML

```
<!--
-->
```

#### 7. Julia

```
#=
=#
```

### Understanding the Commands / Keymapping

**NOTE : You need to go from the top to bottom while commenting on the block of code, otherwise, it would be a mismatch in commenting for specific language syntax. While uncommenting the order doesn't matter.**


#### Getting the extension (filetype)

In Vim, we can get the file extension i.e. we can get the programming language associated with the current file. To do that we can use, `expand('%:e')`.

This will give us the file extension of the current file. Just for simplicity, `%` means the current file, added to it is `:e` for excluding the filename and keeping the extension. We convert the extension into lowercase just for keeping things safe and programmatic and store it in a variable `ext`. 

#### Checking for programming language
We then can then use an if-else ladder to check for the programming languages and assign two variables `cmt1` for the initial characters in the multiline comment and `cmt2` for enclosing the comment. 

#### Typing in the characters

We can use the function `line("'<")` to get the line number of the previous visual selection. Similarly, `line("'>")` for the ending line. We are using the `exe` command to execute the function `line` and so we have to use a concatenation of the commands even to write the raw commands like `i` to insert mode, `o` to insert mode but a line below the cursor. So, we use `normal` command for that. This command indicates the interpreter to execute the following commands from the normal mode. 

We have to enclose the `normal` command in double-quotes/single quotes. We can simply use the variable again with concatenation.

```vim
exe line("'>")."normal o". cmt2 
```  

The above command will fetch the last line's number of the previous visual selection followed by entering `o` from the normal mode and concatenated with the value of the variable `cmt2` which we have already initialized in the `Comment` function. We are using `|` for running multiple commands as we also need to include the comment at the beginning of the visual selection. 

For uncommenting the code, we are simply deleting the entire first and the last line in the visual selection. For that, we have used `dd` from the normal mode.  


### Conclusion

So, from the following type of tutorial, we were able to set up our Vim editor for efficient code commenting/ uncommenting using some commands, key shortcuts, and configuring the vimrc for making custom keymappings. We were also able to understand the multiline comments in various programming languages and use them in Vim very effectively with a simple addon to the config vimrc file.

Thank you for reading, hope you found this article helpful. If you have any queries or wanna add multiline comments for some more programming languages please let me know in the comments or contact section. 

Happy Coding :)

### References

- [StackOverflow - Commenting lines in Vim ](https://stackoverflow.com/questions/1676632/whats-a-quick-way-to-comment-uncomment-lines-in-vim/1676690)

- [Liz Lam - 3 ways to comment code in Vim](https://dev.to/grepliz/3-ways-to-comment-out-blocks-of-code-in-vi-6j4)
 
- [StackExchange - Use variable in normal command](https://vi.stackexchange.com/questions/9644/how-to-use-a-variable-in-the-expression-of-a-normal-command)