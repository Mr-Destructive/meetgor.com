---
article_html: "<p>Let's  say you have multiple github accounts. One for your personal
  projects, one for your company that you work at, and one other remote repository
  account (let's say gitlab).</p>\n<p>You are juggling with multiple accounts, you
  should not waste much time and pick a SSH from those remote repository and pull
  it in your local machine, that makes the process just smooth and saves a ton of
  time.</p>\n<h3>Create a SSH Key</h3>\n<p>To create a SSH key, in linux you can use
  <code>ssh-keygen</code> command.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>ssh-keygen
  -t ed25519 -C <span class=\"s2\">&quot;alice@example.com&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above command will prompt you for two things</p>\n<ol>\n<li>The location where you
  want to store the key</li>\n<li>The passphrase for accessing the key</li>\n</ol>\n<h3>Add
  SSH Key to Github</h3>\n<p>Locate to the <code>ssh</code> folder and copy the generated
  <code>.pub</code> file to your <code>github</code> account.</p>\n<p>For example,
  if you have created the key at <code>~/.ssh/your_name</code> then copy the contents
  of the file <code>~/.ssh/your_name.pub</code> to your clipbaord.</p>\n<p>Navigate
  to your <code>github</code> account and in the settings, <code>SSH and GPG keys</code>
  tab, click on <code>Add SSH key</code> and copy the contents of your clipboard to
  the <code>Key</code> field.</p>\n<h3>Configuring the SSH keys for multiple accounts</h3>\n<pre
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
  \       \n<div class='language-bar'>config</div>\n<div class=\"highlight\"><pre><span></span>Host
  your_company\n  HostName github.com\n  User git\n  IdentityFile ~/.ssh/your_company\n\nHost
  your_name\n  HostName github.com\n  User git\n  IdentityFile ~/.ssh/your_name\n\nHost
  some_name\n  HostName gitlab.com\n  User git\n  IdentityFile ~/.ssh/some_name\n</pre></div>\n\n</pre>\n\n<p>You
  can change the <code>Host</code> config tag values in the <code>~/.ssh/conFig</code></p>\n<p>The
  next time you clone/create a repository on those remote git providers, you need
  to specify the ssh key for that account.</p>\n<p>For example, if you have a repository
  <code>github.com/StartUp_company/some_wired_project</code> then you can specify
  the remote as <code>git@your_company.com:StartUp_company/some_wired_project</code>.
  Here, the <code>git@your_company</code> is the <code>Host</code> value tag from
  the <code>~/.ssh/config</code>. If that repository is from your <code>your_company</code>
  organisation/user scope, you need to add the <code>git@your_company</code> tag,
  if that's your project, simply add <code>git@your_name</code> before the repository
  url i.e. <code>your_name/repo_name</code> which would set the origin as <code>git@your_name:your_name/repo_name</code>,
  here the 1st <code>your_name</code> is the tag from the <code>Host</code> config
  and the 2nd <code>your_name</code> is the github username.</p>\n<p>So, in summary
  if you wanted to use multiple accounts in the same machine, you can understand in
  the following example:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>ssh
  -T git@your_name\n\ngit clone https://github.com/your_name/repo_name\n</pre></div>\n\n</pre>\n\n<p>However,
  you will need to authenticate with the ssh keys in this way everytime you push/pull
  a repository. So for that, you can set the origin with the <code>git@your_name</code>
  tag as the host for automatically authenticating the ssh keys on every push/pull
  or other activities.</p>\n<p>Thanks for reading, Happy Coding :)</p>\n<div class='prevnext'>\n
  \   <style type='text/css'>\n    :root {\n      --prevnext-color-text: #eefbfe;\n
  \     --prevnext-color-angle: #ff6600;\n      --prevnext-subtitle-brightness: 3;\n
  \   }\n    [data-theme=\"light\"] {\n      --prevnext-color-text: #1f2022;\n      --prevnext-color-angle:
  #ffeb00;\n      --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"dark\"]
  {\n      --prevnext-color-text: #eefbfe;\n      --prevnext-color-angle: #ff6600;\n
  \     --prevnext-subtitle-brightness: 3;\n    }\n    .prevnext {\n      display:
  flex;\n      flex-direction: row;\n      justify-content: space-around;\n      align-items:
  flex-start;\n    }\n    .prevnext a {\n      display: flex;\n      align-items:
  center;\n      width: 100%;\n      text-decoration: none;\n    }\n    a.next {\n
  \     justify-content: flex-end;\n    }\n    .prevnext a:hover {\n      background:
  #00000006;\n    }\n    .prevnext-subtitle {\n      color: var(--prevnext-color-text);\n
  \     filter: brightness(var(--prevnext-subtitle-brightness));\n      font-size:
  .8rem;\n    }\n    .prevnext-title {\n      color: var(--prevnext-color-text);\n
  \     font-size: 1rem;\n    }\n    .prevnext-text {\n      max-width: 30vw;\n    }\n
  \   </style>\n<pre><code>&lt;a class='prev' href='/python-pipreqs'&gt;\n\n    &lt;svg
  width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot;
  xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n        &lt;path d=&quot;M13.5
  8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot; stroke-width=&quot;1.5&quot;
  stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt; &lt;/path&gt;\n
  \   &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n
  \       &lt;p class='prevnext-title'&gt;Python Pipreqs: Generate requirements file
  from the imported packages&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a class='next'
  href='/django-bulk-update-queryset'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Django
  Bulk Update QuerySet objects&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot;
  height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
cover: ''
date: 2024-03-22
datetime: 2024-03-22 00:00:00+00:00
description: Setting up SSH config for using multiple accounts for Git repositories.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/tils/2024-03-22-Git-SSH-Multiple-Accounts.md
html: "<p>Let's  say you have multiple github accounts. One for your personal projects,
  one for your company that you work at, and one other remote repository account (let's
  say gitlab).</p>\n<p>You are juggling with multiple accounts, you should not waste
  much time and pick a SSH from those remote repository and pull it in your local
  machine, that makes the process just smooth and saves a ton of time.</p>\n<h3>Create
  a SSH Key</h3>\n<p>To create a SSH key, in linux you can use <code>ssh-keygen</code>
  command.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>ssh-keygen
  -t ed25519 -C <span class=\"s2\">&quot;alice@example.com&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above command will prompt you for two things</p>\n<ol>\n<li>The location where you
  want to store the key</li>\n<li>The passphrase for accessing the key</li>\n</ol>\n<h3>Add
  SSH Key to Github</h3>\n<p>Locate to the <code>ssh</code> folder and copy the generated
  <code>.pub</code> file to your <code>github</code> account.</p>\n<p>For example,
  if you have created the key at <code>~/.ssh/your_name</code> then copy the contents
  of the file <code>~/.ssh/your_name.pub</code> to your clipbaord.</p>\n<p>Navigate
  to your <code>github</code> account and in the settings, <code>SSH and GPG keys</code>
  tab, click on <code>Add SSH key</code> and copy the contents of your clipboard to
  the <code>Key</code> field.</p>\n<h3>Configuring the SSH keys for multiple accounts</h3>\n<pre
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
  \       \n<div class='language-bar'>config</div>\n<div class=\"highlight\"><pre><span></span>Host
  your_company\n  HostName github.com\n  User git\n  IdentityFile ~/.ssh/your_company\n\nHost
  your_name\n  HostName github.com\n  User git\n  IdentityFile ~/.ssh/your_name\n\nHost
  some_name\n  HostName gitlab.com\n  User git\n  IdentityFile ~/.ssh/some_name\n</pre></div>\n\n</pre>\n\n<p>You
  can change the <code>Host</code> config tag values in the <code>~/.ssh/conFig</code></p>\n<p>The
  next time you clone/create a repository on those remote git providers, you need
  to specify the ssh key for that account.</p>\n<p>For example, if you have a repository
  <code>github.com/StartUp_company/some_wired_project</code> then you can specify
  the remote as <code>git@your_company.com:StartUp_company/some_wired_project</code>.
  Here, the <code>git@your_company</code> is the <code>Host</code> value tag from
  the <code>~/.ssh/config</code>. If that repository is from your <code>your_company</code>
  organisation/user scope, you need to add the <code>git@your_company</code> tag,
  if that's your project, simply add <code>git@your_name</code> before the repository
  url i.e. <code>your_name/repo_name</code> which would set the origin as <code>git@your_name:your_name/repo_name</code>,
  here the 1st <code>your_name</code> is the tag from the <code>Host</code> config
  and the 2nd <code>your_name</code> is the github username.</p>\n<p>So, in summary
  if you wanted to use multiple accounts in the same machine, you can understand in
  the following example:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>ssh
  -T git@your_name\n\ngit clone https://github.com/your_name/repo_name\n</pre></div>\n\n</pre>\n\n<p>However,
  you will need to authenticate with the ssh keys in this way everytime you push/pull
  a repository. So for that, you can set the origin with the <code>git@your_name</code>
  tag as the host for automatically authenticating the ssh keys on every push/pull
  or other activities.</p>\n<p>Thanks for reading, Happy Coding :)</p>\n<div class='prevnext'>\n
  \   <style type='text/css'>\n    :root {\n      --prevnext-color-text: #eefbfe;\n
  \     --prevnext-color-angle: #ff6600;\n      --prevnext-subtitle-brightness: 3;\n
  \   }\n    [data-theme=\"light\"] {\n      --prevnext-color-text: #1f2022;\n      --prevnext-color-angle:
  #ffeb00;\n      --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"dark\"]
  {\n      --prevnext-color-text: #eefbfe;\n      --prevnext-color-angle: #ff6600;\n
  \     --prevnext-subtitle-brightness: 3;\n    }\n    .prevnext {\n      display:
  flex;\n      flex-direction: row;\n      justify-content: space-around;\n      align-items:
  flex-start;\n    }\n    .prevnext a {\n      display: flex;\n      align-items:
  center;\n      width: 100%;\n      text-decoration: none;\n    }\n    a.next {\n
  \     justify-content: flex-end;\n    }\n    .prevnext a:hover {\n      background:
  #00000006;\n    }\n    .prevnext-subtitle {\n      color: var(--prevnext-color-text);\n
  \     filter: brightness(var(--prevnext-subtitle-brightness));\n      font-size:
  .8rem;\n    }\n    .prevnext-title {\n      color: var(--prevnext-color-text);\n
  \     font-size: 1rem;\n    }\n    .prevnext-text {\n      max-width: 30vw;\n    }\n
  \   </style>\n<pre><code>&lt;a class='prev' href='/python-pipreqs'&gt;\n\n    &lt;svg
  width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot;
  xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n        &lt;path d=&quot;M13.5
  8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot; stroke-width=&quot;1.5&quot;
  stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt; &lt;/path&gt;\n
  \   &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n
  \       &lt;p class='prevnext-title'&gt;Python Pipreqs: Generate requirements file
  from the imported packages&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a class='next'
  href='/django-bulk-update-queryset'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Django
  Bulk Update QuerySet objects&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot;
  height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
long_description: Let You are juggling with multiple accounts, you should not waste
  much time and pick a SSH from those remote repository and pull it in your local
  machine, that makes the process just smooth and saves a ton of time. To create a
  SSH key, in linux you c
now: 2025-05-01 18:11:33.315145
path: blog/tils/2024-03-22-Git-SSH-Multiple-Accounts.md
slug: git-ssh-multiple-accounts
status: published-til
tags:
- git
- github
templateKey: til
title: Adding SSH Keys for Multiple Accounts in Git
today: 2025-05-01
---

Let's  say you have multiple github accounts. One for your personal projects, one for your company that you work at, and one other remote repository account (let's say gitlab).

You are juggling with multiple accounts, you should not waste much time and pick a SSH from those remote repository and pull it in your local machine, that makes the process just smooth and saves a ton of time.

### Create a SSH Key

To create a SSH key, in linux you can use `ssh-keygen` command.

```bash
ssh-keygen -t ed25519 -C "alice@example.com"
```

The above command will prompt you for two things

1. The location where you want to store the key
2. The passphrase for accessing the key


### Add SSH Key to Github

Locate to the `ssh` folder and copy the generated `.pub` file to your `github` account.

For example, if you have created the key at `~/.ssh/your_name` then copy the contents of the file `~/.ssh/your_name.pub` to your clipbaord.

Navigate to your `github` account and in the settings, `SSH and GPG keys` tab, click on `Add SSH key` and copy the contents of your clipboard to the `Key` field.


### Configuring the SSH keys for multiple accounts

```config
Host your_company
  HostName github.com
  User git
  IdentityFile ~/.ssh/your_company

Host your_name
  HostName github.com
  User git
  IdentityFile ~/.ssh/your_name

Host some_name
  HostName gitlab.com
  User git
  IdentityFile ~/.ssh/some_name
```

You can change the `Host` config tag values in the `~/.ssh/conFig`

The next time you clone/create a repository on those remote git providers, you need to specify the ssh key for that account.

For example, if you have a repository `github.com/StartUp_company/some_wired_project` then you can specify the remote as `git@your_company.com:StartUp_company/some_wired_project`. Here, the `git@your_company` is the `Host` value tag from the `~/.ssh/config`. If that repository is from your `your_company` organisation/user scope, you need to add the `git@your_company` tag, if that's your project, simply add `git@your_name` before the repository url i.e. `your_name/repo_name` which would set the origin as `git@your_name:your_name/repo_name`, here the 1st `your_name` is the tag from the `Host` config and the 2nd `your_name` is the github username.

So, in summary if you wanted to use multiple accounts in the same machine, you can understand in the following example:

```bash
ssh -T git@your_name

git clone https://github.com/your_name/repo_name
```

However, you will need to authenticate with the ssh keys in this way everytime you push/pull a repository. So for that, you can set the origin with the `git@your_name` tag as the host for automatically authenticating the ssh keys on every push/pull or other activities.

Thanks for reading, Happy Coding :)
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
    
    <a class='prev' href='/python-pipreqs'>
    
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M13.5 8.25L9.75 12L13.5 15.75" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"> </path>
        </svg>
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>prev</p>
            <p class='prevnext-title'>Python Pipreqs: Generate requirements file from the imported packages</p>
        </div>
    </a>
    
    <a class='next' href='/django-bulk-update-queryset'>
    
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>next</p>
            <p class='prevnext-title'>Django Bulk Update QuerySet objects</p>
        </div>
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.5 15.75L14.25 12L10.5 8.25" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path>
        </svg>
    </a>
  </div>