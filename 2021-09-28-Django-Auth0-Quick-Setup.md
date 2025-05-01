---
article_html: "<h2>Introduction</h2>\n<p>This is a guide and a walkthrough of how
  to quickly set up a base Django project with Auth0 as integration for authentication
  and authorization. I will walk you through the Django setup and how to use and integrate
  the functionalities of the Auth0.  I will also discuss how why you should be using
  Auth0 and why I love it.</p>\n<p>The script takes <code>2:44</code> minutes time
  to do everything from scratch. From installing virtualenv in python to integrating
  the Auth0 application.</p>\n<p>Here's how the script works:</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1632830813802/MOvedPYdt.gif\"
  alt=\"authodj.gif\" /></p>\n<h3>Contents</h3>\n<ul>\n<li><a href=\"#what-is-auth0\">What
  is Auth0</a></li>\n<li><a href=\"#why-i-love-auth0\">Why I love Auth0</a></li>\n<li><a
  href=\"#set-up-a-django-project\">Set up a Django Project</a></li>\n<li><a href=\"#integrate-auth0-to-a-django-project\">Integrate
  Auth0 to a Django project</a></li>\n<li><a href=\"#creating-a-bash-script-for-integrating-auth0\">Creating
  a BASH Script for integrating Auth0</a>\n<ul>\n<li><a href=\"#appending-to-a-file\">Appending
  to a file</a></li>\n<li><a href=\"#adding-text-before-a-particular-line-using-sed\">Adding
  text before a particular line using <code>sed</code> </a></li>\n<li><a href=\"#appending-to-a-line-using-sed\">Appending
  to a line using <code>sed</code></a></li>\n</ul>\n</li>\n<li><a href=\"#complete-bash-script\">Complete
  BASH Script</a></li>\n<li><a href=\"#conclusion\">Conclusion</a></li>\n</ul>\n<h2>What
  is Auth0</h2>\n<p>Auth0 (<code>Auth zero</code>) is a platform that provides easy
  authentication and authorization for a number of platforms in various programming
  languages and frameworks. The easy-to-follow documentation, availability for almost
  all web frameworks across platforms make it a big bonus for developers. They actually
  make the Developer experience flawless and beginner-friendly.</p>\n<p>According
  to Auth0,</p>\n<blockquote>\n<p>They make your login box awesome</p>\n</blockquote>\n<p>And
  how true is that they make things pretty convenient and wicked fast to integrate
  a smooth functional backend for authentication and authorization. Of course, there
  are more things they offer than just making authentication systems but it is by
  far what the world knows them for.</p>\n<h2>Why I love Auth0</h2>\n<p>Auth0 is a
  generous company that provides a free tier for a limited capacity of authentication
  and that might be more than enough for a developer getting its feet wet in the web
  development (backend).</p>\n<p>They even provide a nice user interface out of the
  box for login/signup and even a dashboard ready-made, which is quite a lot of heavy
  lifting already done for you. Also, there is a dashboard for analyzing the number
  of sign-in/logins into the particular app. This provides the admin/developer of
  the app to get a closer look at the user registered in a day/week/months, number
  of active users, and so on.</p>\n<p>So, who would not love it? I am willing to write
  and use their service for some of my projects. I already have used one for the Hashnode
  x Auth0 Hackathon, I made <a href=\"https://github.com/Mr-Destructive/devquotes\">devquotes</a>
  using the authentication of Auth0 in my Django application.</p>\n<h2>Set up a Django
  Project</h2>\n<p>If you are reading this you already know how to set up a Django
  project, I assume. But nevertheless, I can just include a quick introduction on
  how to do it. I have a script to do this.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/usr/bin/env bash</span>\n\nmkdir <span class=\"nv\">$1</span>\n<span
  class=\"nb\">cd</span> <span class=\"nv\">$1</span>\npip install virtualenv\nvirtualenv
  env\n<span class=\"nb\">source</span> env<span class=\"se\">\\b</span>in<span class=\"se\">\\a</span>ctivate\n\npip
  install django\ndjango-admin startproject <span class=\"nv\">$1</span> .\nclear\n</pre></div>\n\n</pre>\n\n<p>You
  can check out  <a href=\"https://techstructiveblog.hashnode.dev/django-quick-setup-script\">Django
  Quick Setup Script</a>  for the details of this script and also a more in-depth
  guide of Django project setup.</p>\n<p>But if you want to understand the basics
  of the Django project setup here is a little guide about it:</p>\n<p>Firstly, create
  a virtual environment, it's not mandatory but it keeps things simple and easy for
  your project in correspondence to the entire OS. So in python, we have a module
  to create the virtual environment pretty easily,</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>pip
  install virtualenv\n</pre></div>\n\n</pre>\n\n<p>You can use <code>pip3</code> or
  <code>pip -m</code>, or however you install normal python modules. This just installs
  the python virtual environment, we need to create one in the current folder, so
  for that navigate to the folder where you want to create the project and enter the
  following command:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>virtualenv
  venv\n</pre></div>\n\n</pre>\n\n<p>Here, <code>venv</code> can be anything like
  <code>env</code> just for your understanding and simplicity it's a standard name
  kept for the same. After this, you will see a folder of the same name i.e. <code>venv</code>
  or any other name you have used. This is the folder where python will keep every
  installation private to the local folder itself. Now, we need to activate the virtual
  environment, for that we can use the command :</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># for Linux/macOS :</span>\n<span class=\"nb\">source</span> venv/bin/activate\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n<div class='language-bar'>batch</div>\n<div class=\"highlight\"><pre><span></span>#
  for Windows:\nvenv\\Scripts\\activate\n</pre></div>\n\n</pre>\n\n<p>After this,
  your command prompt will have a <code>(venv)</code> attached to its start. This
  indicates you are in a virtual environment, things you do here, may it be module
  installation or any configuration related to python will stay in the local folder
  itself.</p>\n<p>After the virtual environment is set up and activated, you can install
  Django and get started with it. Firstly, install Django using pip:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>pip
  install django\n</pre></div>\n\n</pre>\n\n<p>After the installation is completed,
  you can start a Django project in the current folder using the command:</p>\n<pre
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>django-admin
  startproject name\n</pre></div>\n\n</pre>\n\n<p>Here name can be your project name.
  After this, you will see one new folder and one file pop up.\nNamely, the <code>project
  named</code> folder and <code>manage.py</code> file. So you don't have to touch
  the <code>manage.py</code> file but we use it in most of the commands to use the
  Django functionalities.</p>\n<p>You can now run your basic server using the command
  :</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>python
  manage.py runserver\n</pre></div>\n\n</pre>\n\n<p>There is a base installed/setup
  of the Django project. Moving on in integrating the Auth0 login functionality in
  our webpage.</p>\n<h2>Integrate the Auth0 app in your project</h2>\n<p>So, for integrating
  the Auth0 app for your web application, you need to have an Auth0 account, you can
  signup here. After this you can create an Auth0 application for any type of application,
  we have a couple of options:</p>\n<ul>\n<li>Native Application</li>\n<li>Single
  Page Application</li>\n<li>Regular Web Application</li>\n<li>Machine to Machine
  Application</li>\n</ul>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1632748408741/sUeS8AgrX.png\"
  alt=\"image.png\" /></p>\n<p>You can choose as per your needs, but mostly it would
  be a Regular Web application if you have a backend like Django, Nodejs, PHP, or
  other types of frameworks and languages. So, moving ahead we have created an application
  for the Django framework as a backend. Now, we have a <code>Settings</code> tab
  in the application dashboard, where we have different credentials for the Auth0
  app to talk to our application.</p>\n<p>The credentials needed to be stored safely
  are:</p>\n<ul>\n<li>domain</li>\n<li>Client ID (Key)</li>\n<li>Client Secret</li>\n</ul>\n<p>This
  has to be secured for our local application which will go into production when ready.
  You can use several options like dotenv, environment variables, and so on when the
  application is being deployed but for now, let's hardcode them in our Django project.</p>\n<p>Now,
  you can follow the simple straightforward procedure to copy-paste your credentials
  from the  <a href=\"https://auth0.com/docs/quickstart/webapp/django/01-login#logout\">Auth0
  official documentation</a>. It's quite straightforward to follow the steps even
  for a beginner.</p>\n<p>After the Auth0 app has been configured following the procedure
  in the documentation, you need to integrate several files like dashboard and index
  templates into your custom templates.</p>\n<p>Following additional changes are also
  to be made if you have a user-defined app for your Django project.</p>\n<p>In the
  <code>auth0login</code> app, <code>view.py</code> file:</p>\n<ol>\n<li>The <code>index</code>
  function renders the base file for your project if the user is logged in.</li>\n<li>The
  <code>dashboard</code> function renders the baked version of your profile/dashboard
  of users on your app.</li>\n</ol>\n<p>You would also need to add the root URIs of
  your app that you will be using for testing or in production. For example, we can
  add <code>http://127.0.0.1:8000</code> to allow and use Auth0 in our development
  environment locally.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1632828981455/2gc4h7tTh.png\"
  alt=\"image.png\" /></p>\n<p>You also need to specify the callback URLs for your
  application which is <code>/complete/auth0</code> by default.</p>\n<h2>Creating
  a BASH Script for integrating Auth0</h2>\n<p>So, we can now dive into creating the
  BASH Script to set up the Django x Auth0 application in minutes. The script is quite
  large, like 200 lines but don't worry! Its automation reduces the pain of integrating
  a User Authorization flawlessly. I am also thinking of adding the <code>cURL</code>
  command and parsing in the Client ids, keys, and secret keys, etc.</p>\n<h3>Appending
  to a file</h3>\n<p>We can use the <code>cat</code> command to append text to a file,
  using the syntax as below:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>cat
  <span class=\"s\">&lt;&lt; EOF &gt;&gt; filename</span>\n<span class=\"s\">text</span>\n<span
  class=\"s\">more text</span>\n<span class=\"s\">EOF</span>\n</pre></div>\n\n</pre>\n\n<p>Remember
  here EOF is just a label to stop the command and save it to the file.</p>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1632826339707/_g_RXP3NE.png\"
  alt=\"image.png\" /></p>\n<p>So, we can see that we were able to append to a file,
  multiple lines using the cat command.</p>\n<p>We have used this concept in adding
  configuration and credentials to the <code>settings.py</code> or the <code>urls.py</code>
  files.</p>\n<h3>Adding text before a particular line using <code>sed</code></h3>\n<p><code>sed</code>
  is a great command, and there is nothing you can't do with it, OK there might be
  exceptions. We can get write to a file directly (don't display the output) and specify
  the line number before which we want to append the text. We can then add the text
  we want and followed by the filename.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>sed
  -i <span class=\"s1\">&#39;33 i sometext here&#39;</span> filename\n</pre></div>\n\n</pre>\n\n<p>Here,
  <code>33</code> is the line number in the file which we want to insert before. We
  have used <code>'&quot;'</code> to add a <code>'</code> inside a <code>'</code>,
  this might feel a bit wired but that is how it is in BASH.</p>\n<p>Let's say you
  want to add <code>print('Hello, World!')</code> to a particular line, we have to
  enclose <code>'</code> with these <code>&quot;'</code>( double and single quotes),</p>\n<pre
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>sed
  -i <span class=\"s1\">&#39;2i print(&#39;</span><span class=\"s2\">&quot;&#39;Hello,
  World&#39;&quot;</span><span class=\"s1\">&#39;)&#39;</span> hello.py\n</pre></div>\n\n</pre>\n\n<p>This
  will add the line <code>print('Hello World')</code> to the file <code>hello.py</code></p>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1632824742835/Uj8AF07UG.png\"
  alt=\"image.png\" /></p>\n<h3>Appending to a line using sed</h3>\n<p>We can even
  append text to a particular line using sed, we can use some escape characters and
  regex to add the text from the end of the line.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>sed
  -i <span class=\"s1\">&#39;2i s/$/ textgoes here /&#39;</span> filename\n</pre></div>\n\n</pre>\n\n<p>Here
  2 is any number of line you want to add text to, next <code>i</code> a prompt for
  inserting text and then we have regex like <code>s/$/ /</code>, this will put the
  text enclosed in <code>/ /</code> to the end of the line as indicated by <code>$</code>.
  \ We have the filename at its usual place as before.</p>\n<p>So, lets say, I want
  to add a comment to the second line in the previous example, I can use the following
  command to do it:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>sed
  -i <span class=\"s1\">&#39;2 s/$/ # another comment/&#39;</span> hello.py\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1632825067925/0eU2mkCDI.png\"
  alt=\"image.png\" /></p>\n<p>We have used these commands to add the <code>include</code>
  function in the <code>urls.py</code> in the project folder.</p>\n<p>So those were
  all the operations we used for doing some automated editing for the Auth0 app integration
  to our Django project.</p>\n<p>Below is the entire script and is also uploaded on
  <a href=\"https://github.com/Mr-Destructive/django-auth0-quick-setup\">GitHub</a>.</p>\n<h2>Complete
  BASH Script</h2>\n<p>You can run the file by parsing the name of your project.</p>\n<pre
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>bash
  script.sh mywebsite\n</pre></div>\n\n</pre>\n\n<p>Wait for some 2-3 minutes, and
  the script will produce the Django application with the Auth0 app integrated. You
  will have to enter the credentials manually wherever applicable.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/usr/bin/env bash</span>\n\nmkdir <span class=\"nv\">$1</span>\n<span
  class=\"nb\">cd</span> <span class=\"nv\">$1</span>\npip install virtualenv\nvirtualenv
  venv\n<span class=\"nb\">source</span> venv/Scripts/activate\n\npip install django\ndjango-admin
  startproject <span class=\"nv\">$1</span> .\n\ncat <span class=\"s\">&lt;&lt; EOF
  &gt;&gt; requirements.txt</span>\n<span class=\"s\">social-auth-app-django~=3.1
  </span>\n<span class=\"s\">python-jose~=3.0 </span>\n<span class=\"s\">python-dotenv~=0.9</span>\n<span
  class=\"s\">EOF</span>\n\npip install -r requirements.txt\n\npip freeze &gt; requirements.txt\n\npython
  manage.py startapp auth0login\n\ntouch auth0login/urls.py\nmkdir auth0login/templates\ntouch
  auth0login/templates/index.html\ntouch auth0login/templates/dashboard.html\n\nsed
  -i <span class=\"s1\">&#39;40 i \\    &#39;</span><span class=\"s2\">&quot;&#39;&quot;</span><span
  class=\"s1\">&#39;social_django&#39;</span><span class=\"s2\">&quot;&#39;&quot;</span><span
  class=\"s1\">&#39;,&#39;</span> <span class=\"nv\">$1</span>/settings.py\nsed -i
  <span class=\"s1\">&#39;41 i \\    &#39;</span><span class=\"s2\">&quot;&#39;&quot;</span><span
  class=\"s1\">&#39;auth0login&#39;</span><span class=\"s2\">&quot;&#39;&quot;</span><span
  class=\"s1\">&#39;,&#39;</span> <span class=\"nv\">$1</span>/settings.py\nsed -i
  <span class=\"s1\">&#39;21 i \\    path(&#39;</span><span class=\"s2\">&quot;&#39;&#39;&quot;</span><span
  class=\"s1\">&#39;, include(&#39;</span><span class=\"s2\">&quot;&#39;auth0login.urls&#39;&quot;</span><span
  class=\"s1\">&#39;)),&#39;</span> <span class=\"nv\">$1</span>/urls.py\nsed -i <span
  class=\"s1\">&#39;17 s/$/, include/&#39;</span> <span class=\"nv\">$1</span>/urls.py
  \n\ncat <span class=\"s\">&lt;&lt; EOF &gt;&gt; $1/settings.py</span>\n<span class=\"s\">SOCIAL_AUTH_TRAILING_SLASH
  = False  # Remove trailing slash from routes</span>\n<span class=\"s\">SOCIAL_AUTH_AUTH0_DOMAIN
  = &#39;YOUR_DOMAIN&#39;</span>\n<span class=\"s\">SOCIAL_AUTH_AUTH0_KEY = &#39;YOUR_CLIENT_ID&#39;</span>\n<span
  class=\"s\">SOCIAL_AUTH_AUTH0_SECRET = &#39;YOUR_CLIENT_SECRET&#39;</span>\n<span
  class=\"s\">EOF</span>\n\ncat <span class=\"s\">&lt;&lt; EOF &gt;&gt; $1/settings.py
  </span>\n<span class=\"s\">SOCIAL_AUTH_AUTH0_SCOPE = [</span>\n<span class=\"s\">
  \   &#39;openid&#39;,</span>\n<span class=\"s\">    &#39;profile&#39;,</span>\n<span
  class=\"s\">    &#39;email&#39;</span>\n<span class=\"s\">]</span>\n<span class=\"s\">EOF</span>\n\npython
  manage.py migrate\n\ncat <span class=\"s\">&lt;&lt; EOF &gt;&gt;auth0login/auth0backend.py</span>\n\n<span
  class=\"s\">from urllib import request</span>\n<span class=\"s\">from jose import
  jwt</span>\n<span class=\"s\">from social_core.backends.oauth import BaseOAuth2</span>\n\n\n<span
  class=\"s\">class Auth0(BaseOAuth2):</span>\n<span class=\"s\">    &quot;&quot;&quot;Auth0
  OAuth authentication backend&quot;&quot;&quot;</span>\n<span class=\"s\">    name
  = &#39;auth0&#39;</span>\n<span class=\"s\">    SCOPE_SEPARATOR = &#39; &#39;</span>\n<span
  class=\"s\">    ACCESS_TOKEN_METHOD = &#39;POST&#39;</span>\n<span class=\"s\">
  \   REDIRECT_STATE = False</span>\n<span class=\"s\">    EXTRA_DATA = [</span>\n<span
  class=\"s\">        (&#39;picture&#39;, &#39;picture&#39;),</span>\n<span class=\"s\">
  \       (&#39;email&#39;, &#39;email&#39;)</span>\n<span class=\"s\">    ]</span>\n\n<span
  class=\"s\">    def authorization_url(self):</span>\n<span class=\"s\">        return
  &#39;https://&#39; + self.setting(&#39;DOMAIN&#39;) + &#39;/authorize&#39;</span>\n\n<span
  class=\"s\">    def access_token_url(self):</span>\n<span class=\"s\">        return
  &#39;https://&#39; + self.setting(&#39;DOMAIN&#39;) + &#39;/oauth/token&#39;</span>\n\n<span
  class=\"s\">    def get_user_id(self, details, response):</span>\n<span class=\"s\">
  \       &quot;&quot;&quot;Return current user id.&quot;&quot;&quot;</span>\n<span
  class=\"s\">        return details[&#39;user_id&#39;]</span>\n\n<span class=\"s\">
  \   def get_user_details(self, response):</span>\n<span class=\"s\">        # Obtain
  JWT and the keys to validate the signature</span>\n<span class=\"s\">        id_token
  = response.get(&#39;id_token&#39;)</span>\n<span class=\"s\">        jwks = request.urlopen(&#39;https://&#39;
  + self.setting(&#39;DOMAIN&#39;) + &#39;/.well-known/jwks.json&#39;)</span>\n<span
  class=\"s\">        issuer = &#39;https://&#39; + self.setting(&#39;DOMAIN&#39;)
  + &#39;/&#39;</span>\n<span class=\"s\">        audience = self.setting(&#39;KEY&#39;)
  \ # CLIENT_ID</span>\n<span class=\"s\">        payload = jwt.decode(id_token, jwks.read(),
  algorithms=[&#39;RS256&#39;], audience=audience, issuer=issuer)</span>\n\n<span
  class=\"s\">        return {&#39;username&#39;: payload[&#39;nickname&#39;],</span>\n<span
  class=\"s\">                &#39;first_name&#39;: payload[&#39;name&#39;],</span>\n<span
  class=\"s\">                &#39;picture&#39;: payload[&#39;picture&#39;],</span>\n<span
  class=\"s\">                &#39;user_id&#39;: payload[&#39;sub&#39;],</span>\n<span
  class=\"s\">                &#39;email&#39;: payload[&#39;email&#39;]}</span>\n\n<span
  class=\"s\">EOF</span>\n\ncat <span class=\"s\">&lt;&lt; EOF &gt;&gt; $1/settings.py</span>\n\n<span
  class=\"s\">AUTHENTICATION_BACKENDS = {</span>\n<span class=\"s\">    #&#39;YOUR_DJANGO_APP_NAME.auth0backend.Auth0&#39;,</span>\n<span
  class=\"s\">    &#39;django.contrib.auth.backends.ModelBackend&#39;</span>\n<span
  class=\"s\">}</span>\n\n<span class=\"s\">EOF</span>\n\ncat <span class=\"s\">&lt;&lt;
  EOF &gt;&gt; $1/settings.py</span>\n\n<span class=\"s\">LOGIN_URL = &#39;/login/auth0&#39;</span>\n<span
  class=\"s\">LOGIN_REDIRECT_URL = &#39;/dashboard&#39;</span>\n<span class=\"s\">EOF</span>\n\ncat
  &gt; auth0login/views.py<span class=\"s\">&lt;&lt;EOF</span>\n\n<span class=\"s\">from
  django.shortcuts import render, redirect</span>\n<span class=\"s\">from django.contrib.auth.decorators
  import login_required</span>\n<span class=\"s\">from django.contrib.auth import
  logout as log_out</span>\n<span class=\"s\">from django.conf import settings</span>\n<span
  class=\"s\">from django.http import HttpResponseRedirect</span>\n<span class=\"s\">from
  urllib.parse import urlencode</span>\n<span class=\"s\">import json</span>\n\n<span
  class=\"s\">def index(request):</span>\n<span class=\"s\">    user = request.user</span>\n<span
  class=\"s\">    if user.is_authenticated:</span>\n<span class=\"s\">        return
  redirect(dashboard)</span>\n<span class=\"s\">    else:</span>\n<span class=\"s\">
  \       return render(request, &#39;index.html&#39;)</span>\n\n\n<span class=\"s\">@login_required</span>\n<span
  class=\"s\">def dashboard(request):</span>\n<span class=\"s\">    user = request.user</span>\n<span
  class=\"s\">    auth0user = user.social_auth.get(provider=&#39;auth0&#39;)</span>\n<span
  class=\"s\">    userdata = {</span>\n<span class=\"s\">        &#39;user_id&#39;:
  auth0user.uid,</span>\n<span class=\"s\">        &#39;name&#39;: user.first_name,</span>\n<span
  class=\"s\">        &#39;picture&#39;: auth0user.extra_data[&#39;picture&#39;],</span>\n<span
  class=\"s\">        &#39;email&#39;: auth0user.extra_data[&#39;email&#39;],</span>\n<span
  class=\"s\">    }</span>\n\n<span class=\"s\">    return render(request, &#39;dashboard.html&#39;,
  {</span>\n<span class=\"s\">        &#39;auth0User&#39;: auth0user,</span>\n<span
  class=\"s\">        &#39;userdata&#39;: json.dumps(userdata, indent=4)</span>\n<span
  class=\"s\">    })</span>\n\n<span class=\"s\">def logout(request):</span>\n<span
  class=\"s\">    log_out(request)</span>\n<span class=\"s\">    return_to = urlencode({&#39;returnTo&#39;:
  request.build_absolute_uri(&#39;/&#39;)})</span>\n<span class=\"s\">    logout_url
  = &#39;https://%s/v2/logout?client_id=%s&amp;%s&#39; % \\</span>\n<span class=\"s\">
  \                (settings.SOCIAL_AUTH_AUTH0_DOMAIN, settings.SOCIAL_AUTH_AUTH0_KEY,
  return_to)</span>\n<span class=\"s\">    return HttpResponseRedirect(logout_url)</span>\n\n<span
  class=\"s\">EOF</span>\n\ncat <span class=\"s\">&lt;&lt; EOF &gt;&gt; auth0login/templates/index.html</span>\n\n<span
  class=\"s\">&lt;div class=&quot;login-box auth0-box before&quot;&gt;</span>\n<span
  class=\"s\">    &lt;img src=&quot;https://i.cloudup.com/StzWWrY34s.png&quot; /&gt;</span>\n<span
  class=\"s\">    &lt;h3&gt;Auth0 Example&lt;/h3&gt;</span>\n<span class=\"s\">    &lt;p&gt;Zero
  friction identity infrastructure, built for developers&lt;/p&gt;</span>\n<span class=\"s\">
  \   &lt;a class=&quot;btn btn-primary btn-lg btn-login btn-block&quot; href=&quot;/login/auth0&quot;&gt;Log
  In&lt;/a&gt;</span>\n<span class=\"s\">&lt;/div&gt;</span>\n<span class=\"s\">EOF</span>\n\ncat
  <span class=\"s\">&lt;&lt; EOF &gt;&gt; auth0login/templates/dashboard.html</span>\n\n<span
  class=\"s\">&lt;div class=&quot;logged-in-box auth0-box logged-in&quot;&gt;</span>\n<span
  class=\"s\">    &lt;h1 id=&quot;logo&quot;&gt;&lt;img src=&quot;//cdn.auth0.com/samples/auth0_logo_final_blue_RGB.png&quot;
  /&gt;&lt;/h1&gt;</span>\n<span class=\"s\">    &lt;img class=&quot;avatar&quot;
  src=&quot;{{ auth0User.extra_data.picture }}&quot;/&gt;</span>\n<span class=\"s\">
  \   &lt;h2&gt;Welcome {{ user.username }}&lt;/h2&gt;</span>\n<span class=\"s\">
  \   &lt;pre&gt;{{ userdata }}&lt;/pre&gt;</span>\n<span class=\"s\">&lt;/div&gt;</span>\n<span
  class=\"s\">EOF</span>\n\ncat <span class=\"s\">&lt;&lt; EOF &gt;&gt; auth0login/urls.py</span>\n<span
  class=\"s\">from django.urls import path, include</span>\n<span class=\"s\">from
  . import views</span>\n\n<span class=\"s\">urlpatterns = [</span>\n<span class=\"s\">
  \   path(&#39;&#39;, views.index),</span>\n<span class=\"s\">    path(&#39;dashboard&#39;,
  views.dashboard),</span>\n<span class=\"s\">    path(&#39;logout&#39;, views.logout),</span>\n<span
  class=\"s\">    path(&#39;&#39;, include(&#39;django.contrib.auth.urls&#39;)),</span>\n<span
  class=\"s\">    path(&#39;&#39;, include(&#39;social_django.urls&#39;)),</span>\n<span
  class=\"s\">]</span>\n\n<span class=\"s\">EOF</span>\n\npython manage.py makemigrations\npython
  manage.py migrate\n</pre></div>\n\n</pre>\n\n<h2>Conclusion</h2>\n<p>Ok, so this
  was it, a quite big script but that's how automation can be. We were able to set
  up a Django base application with a ready app of Auth0 to extend the functionality.
  This was just a basic script also you can extend the functionalities like adding
  a curl command to fetch the credentials and make it more automated but that was
  not the aim of this article.</p>\n<p>If you had any issues using the script please
  let me know, I'll be happy to fix those. Thanks for reading. Happy Coding :)</p>\n"
cover: ''
date: 2021-09-28
datetime: 2021-09-28 00:00:00+00:00
description: This is a guide and a walkthrough of how to quickly set up a base Django
  project with Auth0 as integration for authentication and authorization. I will walk
  you
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-09-28-Django-Auth0-Quick-Setup.md
html: "<h2>Introduction</h2>\n<p>This is a guide and a walkthrough of how to quickly
  set up a base Django project with Auth0 as integration for authentication and authorization.
  I will walk you through the Django setup and how to use and integrate the functionalities
  of the Auth0.  I will also discuss how why you should be using Auth0 and why I love
  it.</p>\n<p>The script takes <code>2:44</code> minutes time to do everything from
  scratch. From installing virtualenv in python to integrating the Auth0 application.</p>\n<p>Here's
  how the script works:</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1632830813802/MOvedPYdt.gif\"
  alt=\"authodj.gif\" /></p>\n<h3>Contents</h3>\n<ul>\n<li><a href=\"#what-is-auth0\">What
  is Auth0</a></li>\n<li><a href=\"#why-i-love-auth0\">Why I love Auth0</a></li>\n<li><a
  href=\"#set-up-a-django-project\">Set up a Django Project</a></li>\n<li><a href=\"#integrate-auth0-to-a-django-project\">Integrate
  Auth0 to a Django project</a></li>\n<li><a href=\"#creating-a-bash-script-for-integrating-auth0\">Creating
  a BASH Script for integrating Auth0</a>\n<ul>\n<li><a href=\"#appending-to-a-file\">Appending
  to a file</a></li>\n<li><a href=\"#adding-text-before-a-particular-line-using-sed\">Adding
  text before a particular line using <code>sed</code> </a></li>\n<li><a href=\"#appending-to-a-line-using-sed\">Appending
  to a line using <code>sed</code></a></li>\n</ul>\n</li>\n<li><a href=\"#complete-bash-script\">Complete
  BASH Script</a></li>\n<li><a href=\"#conclusion\">Conclusion</a></li>\n</ul>\n<h2>What
  is Auth0</h2>\n<p>Auth0 (<code>Auth zero</code>) is a platform that provides easy
  authentication and authorization for a number of platforms in various programming
  languages and frameworks. The easy-to-follow documentation, availability for almost
  all web frameworks across platforms make it a big bonus for developers. They actually
  make the Developer experience flawless and beginner-friendly.</p>\n<p>According
  to Auth0,</p>\n<blockquote>\n<p>They make your login box awesome</p>\n</blockquote>\n<p>And
  how true is that they make things pretty convenient and wicked fast to integrate
  a smooth functional backend for authentication and authorization. Of course, there
  are more things they offer than just making authentication systems but it is by
  far what the world knows them for.</p>\n<h2>Why I love Auth0</h2>\n<p>Auth0 is a
  generous company that provides a free tier for a limited capacity of authentication
  and that might be more than enough for a developer getting its feet wet in the web
  development (backend).</p>\n<p>They even provide a nice user interface out of the
  box for login/signup and even a dashboard ready-made, which is quite a lot of heavy
  lifting already done for you. Also, there is a dashboard for analyzing the number
  of sign-in/logins into the particular app. This provides the admin/developer of
  the app to get a closer look at the user registered in a day/week/months, number
  of active users, and so on.</p>\n<p>So, who would not love it? I am willing to write
  and use their service for some of my projects. I already have used one for the Hashnode
  x Auth0 Hackathon, I made <a href=\"https://github.com/Mr-Destructive/devquotes\">devquotes</a>
  using the authentication of Auth0 in my Django application.</p>\n<h2>Set up a Django
  Project</h2>\n<p>If you are reading this you already know how to set up a Django
  project, I assume. But nevertheless, I can just include a quick introduction on
  how to do it. I have a script to do this.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/usr/bin/env bash</span>\n\nmkdir <span class=\"nv\">$1</span>\n<span
  class=\"nb\">cd</span> <span class=\"nv\">$1</span>\npip install virtualenv\nvirtualenv
  env\n<span class=\"nb\">source</span> env<span class=\"se\">\\b</span>in<span class=\"se\">\\a</span>ctivate\n\npip
  install django\ndjango-admin startproject <span class=\"nv\">$1</span> .\nclear\n</pre></div>\n\n</pre>\n\n<p>You
  can check out  <a href=\"https://techstructiveblog.hashnode.dev/django-quick-setup-script\">Django
  Quick Setup Script</a>  for the details of this script and also a more in-depth
  guide of Django project setup.</p>\n<p>But if you want to understand the basics
  of the Django project setup here is a little guide about it:</p>\n<p>Firstly, create
  a virtual environment, it's not mandatory but it keeps things simple and easy for
  your project in correspondence to the entire OS. So in python, we have a module
  to create the virtual environment pretty easily,</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>pip
  install virtualenv\n</pre></div>\n\n</pre>\n\n<p>You can use <code>pip3</code> or
  <code>pip -m</code>, or however you install normal python modules. This just installs
  the python virtual environment, we need to create one in the current folder, so
  for that navigate to the folder where you want to create the project and enter the
  following command:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>virtualenv
  venv\n</pre></div>\n\n</pre>\n\n<p>Here, <code>venv</code> can be anything like
  <code>env</code> just for your understanding and simplicity it's a standard name
  kept for the same. After this, you will see a folder of the same name i.e. <code>venv</code>
  or any other name you have used. This is the folder where python will keep every
  installation private to the local folder itself. Now, we need to activate the virtual
  environment, for that we can use the command :</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># for Linux/macOS :</span>\n<span class=\"nb\">source</span> venv/bin/activate\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n<div class='language-bar'>batch</div>\n<div class=\"highlight\"><pre><span></span>#
  for Windows:\nvenv\\Scripts\\activate\n</pre></div>\n\n</pre>\n\n<p>After this,
  your command prompt will have a <code>(venv)</code> attached to its start. This
  indicates you are in a virtual environment, things you do here, may it be module
  installation or any configuration related to python will stay in the local folder
  itself.</p>\n<p>After the virtual environment is set up and activated, you can install
  Django and get started with it. Firstly, install Django using pip:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>pip
  install django\n</pre></div>\n\n</pre>\n\n<p>After the installation is completed,
  you can start a Django project in the current folder using the command:</p>\n<pre
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>django-admin
  startproject name\n</pre></div>\n\n</pre>\n\n<p>Here name can be your project name.
  After this, you will see one new folder and one file pop up.\nNamely, the <code>project
  named</code> folder and <code>manage.py</code> file. So you don't have to touch
  the <code>manage.py</code> file but we use it in most of the commands to use the
  Django functionalities.</p>\n<p>You can now run your basic server using the command
  :</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>python
  manage.py runserver\n</pre></div>\n\n</pre>\n\n<p>There is a base installed/setup
  of the Django project. Moving on in integrating the Auth0 login functionality in
  our webpage.</p>\n<h2>Integrate the Auth0 app in your project</h2>\n<p>So, for integrating
  the Auth0 app for your web application, you need to have an Auth0 account, you can
  signup here. After this you can create an Auth0 application for any type of application,
  we have a couple of options:</p>\n<ul>\n<li>Native Application</li>\n<li>Single
  Page Application</li>\n<li>Regular Web Application</li>\n<li>Machine to Machine
  Application</li>\n</ul>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1632748408741/sUeS8AgrX.png\"
  alt=\"image.png\" /></p>\n<p>You can choose as per your needs, but mostly it would
  be a Regular Web application if you have a backend like Django, Nodejs, PHP, or
  other types of frameworks and languages. So, moving ahead we have created an application
  for the Django framework as a backend. Now, we have a <code>Settings</code> tab
  in the application dashboard, where we have different credentials for the Auth0
  app to talk to our application.</p>\n<p>The credentials needed to be stored safely
  are:</p>\n<ul>\n<li>domain</li>\n<li>Client ID (Key)</li>\n<li>Client Secret</li>\n</ul>\n<p>This
  has to be secured for our local application which will go into production when ready.
  You can use several options like dotenv, environment variables, and so on when the
  application is being deployed but for now, let's hardcode them in our Django project.</p>\n<p>Now,
  you can follow the simple straightforward procedure to copy-paste your credentials
  from the  <a href=\"https://auth0.com/docs/quickstart/webapp/django/01-login#logout\">Auth0
  official documentation</a>. It's quite straightforward to follow the steps even
  for a beginner.</p>\n<p>After the Auth0 app has been configured following the procedure
  in the documentation, you need to integrate several files like dashboard and index
  templates into your custom templates.</p>\n<p>Following additional changes are also
  to be made if you have a user-defined app for your Django project.</p>\n<p>In the
  <code>auth0login</code> app, <code>view.py</code> file:</p>\n<ol>\n<li>The <code>index</code>
  function renders the base file for your project if the user is logged in.</li>\n<li>The
  <code>dashboard</code> function renders the baked version of your profile/dashboard
  of users on your app.</li>\n</ol>\n<p>You would also need to add the root URIs of
  your app that you will be using for testing or in production. For example, we can
  add <code>http://127.0.0.1:8000</code> to allow and use Auth0 in our development
  environment locally.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1632828981455/2gc4h7tTh.png\"
  alt=\"image.png\" /></p>\n<p>You also need to specify the callback URLs for your
  application which is <code>/complete/auth0</code> by default.</p>\n<h2>Creating
  a BASH Script for integrating Auth0</h2>\n<p>So, we can now dive into creating the
  BASH Script to set up the Django x Auth0 application in minutes. The script is quite
  large, like 200 lines but don't worry! Its automation reduces the pain of integrating
  a User Authorization flawlessly. I am also thinking of adding the <code>cURL</code>
  command and parsing in the Client ids, keys, and secret keys, etc.</p>\n<h3>Appending
  to a file</h3>\n<p>We can use the <code>cat</code> command to append text to a file,
  using the syntax as below:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>cat
  <span class=\"s\">&lt;&lt; EOF &gt;&gt; filename</span>\n<span class=\"s\">text</span>\n<span
  class=\"s\">more text</span>\n<span class=\"s\">EOF</span>\n</pre></div>\n\n</pre>\n\n<p>Remember
  here EOF is just a label to stop the command and save it to the file.</p>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1632826339707/_g_RXP3NE.png\"
  alt=\"image.png\" /></p>\n<p>So, we can see that we were able to append to a file,
  multiple lines using the cat command.</p>\n<p>We have used this concept in adding
  configuration and credentials to the <code>settings.py</code> or the <code>urls.py</code>
  files.</p>\n<h3>Adding text before a particular line using <code>sed</code></h3>\n<p><code>sed</code>
  is a great command, and there is nothing you can't do with it, OK there might be
  exceptions. We can get write to a file directly (don't display the output) and specify
  the line number before which we want to append the text. We can then add the text
  we want and followed by the filename.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>sed
  -i <span class=\"s1\">&#39;33 i sometext here&#39;</span> filename\n</pre></div>\n\n</pre>\n\n<p>Here,
  <code>33</code> is the line number in the file which we want to insert before. We
  have used <code>'&quot;'</code> to add a <code>'</code> inside a <code>'</code>,
  this might feel a bit wired but that is how it is in BASH.</p>\n<p>Let's say you
  want to add <code>print('Hello, World!')</code> to a particular line, we have to
  enclose <code>'</code> with these <code>&quot;'</code>( double and single quotes),</p>\n<pre
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>sed
  -i <span class=\"s1\">&#39;2i print(&#39;</span><span class=\"s2\">&quot;&#39;Hello,
  World&#39;&quot;</span><span class=\"s1\">&#39;)&#39;</span> hello.py\n</pre></div>\n\n</pre>\n\n<p>This
  will add the line <code>print('Hello World')</code> to the file <code>hello.py</code></p>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1632824742835/Uj8AF07UG.png\"
  alt=\"image.png\" /></p>\n<h3>Appending to a line using sed</h3>\n<p>We can even
  append text to a particular line using sed, we can use some escape characters and
  regex to add the text from the end of the line.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>sed
  -i <span class=\"s1\">&#39;2i s/$/ textgoes here /&#39;</span> filename\n</pre></div>\n\n</pre>\n\n<p>Here
  2 is any number of line you want to add text to, next <code>i</code> a prompt for
  inserting text and then we have regex like <code>s/$/ /</code>, this will put the
  text enclosed in <code>/ /</code> to the end of the line as indicated by <code>$</code>.
  \ We have the filename at its usual place as before.</p>\n<p>So, lets say, I want
  to add a comment to the second line in the previous example, I can use the following
  command to do it:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>sed
  -i <span class=\"s1\">&#39;2 s/$/ # another comment/&#39;</span> hello.py\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1632825067925/0eU2mkCDI.png\"
  alt=\"image.png\" /></p>\n<p>We have used these commands to add the <code>include</code>
  function in the <code>urls.py</code> in the project folder.</p>\n<p>So those were
  all the operations we used for doing some automated editing for the Auth0 app integration
  to our Django project.</p>\n<p>Below is the entire script and is also uploaded on
  <a href=\"https://github.com/Mr-Destructive/django-auth0-quick-setup\">GitHub</a>.</p>\n<h2>Complete
  BASH Script</h2>\n<p>You can run the file by parsing the name of your project.</p>\n<pre
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>bash
  script.sh mywebsite\n</pre></div>\n\n</pre>\n\n<p>Wait for some 2-3 minutes, and
  the script will produce the Django application with the Auth0 app integrated. You
  will have to enter the credentials manually wherever applicable.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/usr/bin/env bash</span>\n\nmkdir <span class=\"nv\">$1</span>\n<span
  class=\"nb\">cd</span> <span class=\"nv\">$1</span>\npip install virtualenv\nvirtualenv
  venv\n<span class=\"nb\">source</span> venv/Scripts/activate\n\npip install django\ndjango-admin
  startproject <span class=\"nv\">$1</span> .\n\ncat <span class=\"s\">&lt;&lt; EOF
  &gt;&gt; requirements.txt</span>\n<span class=\"s\">social-auth-app-django~=3.1
  </span>\n<span class=\"s\">python-jose~=3.0 </span>\n<span class=\"s\">python-dotenv~=0.9</span>\n<span
  class=\"s\">EOF</span>\n\npip install -r requirements.txt\n\npip freeze &gt; requirements.txt\n\npython
  manage.py startapp auth0login\n\ntouch auth0login/urls.py\nmkdir auth0login/templates\ntouch
  auth0login/templates/index.html\ntouch auth0login/templates/dashboard.html\n\nsed
  -i <span class=\"s1\">&#39;40 i \\    &#39;</span><span class=\"s2\">&quot;&#39;&quot;</span><span
  class=\"s1\">&#39;social_django&#39;</span><span class=\"s2\">&quot;&#39;&quot;</span><span
  class=\"s1\">&#39;,&#39;</span> <span class=\"nv\">$1</span>/settings.py\nsed -i
  <span class=\"s1\">&#39;41 i \\    &#39;</span><span class=\"s2\">&quot;&#39;&quot;</span><span
  class=\"s1\">&#39;auth0login&#39;</span><span class=\"s2\">&quot;&#39;&quot;</span><span
  class=\"s1\">&#39;,&#39;</span> <span class=\"nv\">$1</span>/settings.py\nsed -i
  <span class=\"s1\">&#39;21 i \\    path(&#39;</span><span class=\"s2\">&quot;&#39;&#39;&quot;</span><span
  class=\"s1\">&#39;, include(&#39;</span><span class=\"s2\">&quot;&#39;auth0login.urls&#39;&quot;</span><span
  class=\"s1\">&#39;)),&#39;</span> <span class=\"nv\">$1</span>/urls.py\nsed -i <span
  class=\"s1\">&#39;17 s/$/, include/&#39;</span> <span class=\"nv\">$1</span>/urls.py
  \n\ncat <span class=\"s\">&lt;&lt; EOF &gt;&gt; $1/settings.py</span>\n<span class=\"s\">SOCIAL_AUTH_TRAILING_SLASH
  = False  # Remove trailing slash from routes</span>\n<span class=\"s\">SOCIAL_AUTH_AUTH0_DOMAIN
  = &#39;YOUR_DOMAIN&#39;</span>\n<span class=\"s\">SOCIAL_AUTH_AUTH0_KEY = &#39;YOUR_CLIENT_ID&#39;</span>\n<span
  class=\"s\">SOCIAL_AUTH_AUTH0_SECRET = &#39;YOUR_CLIENT_SECRET&#39;</span>\n<span
  class=\"s\">EOF</span>\n\ncat <span class=\"s\">&lt;&lt; EOF &gt;&gt; $1/settings.py
  </span>\n<span class=\"s\">SOCIAL_AUTH_AUTH0_SCOPE = [</span>\n<span class=\"s\">
  \   &#39;openid&#39;,</span>\n<span class=\"s\">    &#39;profile&#39;,</span>\n<span
  class=\"s\">    &#39;email&#39;</span>\n<span class=\"s\">]</span>\n<span class=\"s\">EOF</span>\n\npython
  manage.py migrate\n\ncat <span class=\"s\">&lt;&lt; EOF &gt;&gt;auth0login/auth0backend.py</span>\n\n<span
  class=\"s\">from urllib import request</span>\n<span class=\"s\">from jose import
  jwt</span>\n<span class=\"s\">from social_core.backends.oauth import BaseOAuth2</span>\n\n\n<span
  class=\"s\">class Auth0(BaseOAuth2):</span>\n<span class=\"s\">    &quot;&quot;&quot;Auth0
  OAuth authentication backend&quot;&quot;&quot;</span>\n<span class=\"s\">    name
  = &#39;auth0&#39;</span>\n<span class=\"s\">    SCOPE_SEPARATOR = &#39; &#39;</span>\n<span
  class=\"s\">    ACCESS_TOKEN_METHOD = &#39;POST&#39;</span>\n<span class=\"s\">
  \   REDIRECT_STATE = False</span>\n<span class=\"s\">    EXTRA_DATA = [</span>\n<span
  class=\"s\">        (&#39;picture&#39;, &#39;picture&#39;),</span>\n<span class=\"s\">
  \       (&#39;email&#39;, &#39;email&#39;)</span>\n<span class=\"s\">    ]</span>\n\n<span
  class=\"s\">    def authorization_url(self):</span>\n<span class=\"s\">        return
  &#39;https://&#39; + self.setting(&#39;DOMAIN&#39;) + &#39;/authorize&#39;</span>\n\n<span
  class=\"s\">    def access_token_url(self):</span>\n<span class=\"s\">        return
  &#39;https://&#39; + self.setting(&#39;DOMAIN&#39;) + &#39;/oauth/token&#39;</span>\n\n<span
  class=\"s\">    def get_user_id(self, details, response):</span>\n<span class=\"s\">
  \       &quot;&quot;&quot;Return current user id.&quot;&quot;&quot;</span>\n<span
  class=\"s\">        return details[&#39;user_id&#39;]</span>\n\n<span class=\"s\">
  \   def get_user_details(self, response):</span>\n<span class=\"s\">        # Obtain
  JWT and the keys to validate the signature</span>\n<span class=\"s\">        id_token
  = response.get(&#39;id_token&#39;)</span>\n<span class=\"s\">        jwks = request.urlopen(&#39;https://&#39;
  + self.setting(&#39;DOMAIN&#39;) + &#39;/.well-known/jwks.json&#39;)</span>\n<span
  class=\"s\">        issuer = &#39;https://&#39; + self.setting(&#39;DOMAIN&#39;)
  + &#39;/&#39;</span>\n<span class=\"s\">        audience = self.setting(&#39;KEY&#39;)
  \ # CLIENT_ID</span>\n<span class=\"s\">        payload = jwt.decode(id_token, jwks.read(),
  algorithms=[&#39;RS256&#39;], audience=audience, issuer=issuer)</span>\n\n<span
  class=\"s\">        return {&#39;username&#39;: payload[&#39;nickname&#39;],</span>\n<span
  class=\"s\">                &#39;first_name&#39;: payload[&#39;name&#39;],</span>\n<span
  class=\"s\">                &#39;picture&#39;: payload[&#39;picture&#39;],</span>\n<span
  class=\"s\">                &#39;user_id&#39;: payload[&#39;sub&#39;],</span>\n<span
  class=\"s\">                &#39;email&#39;: payload[&#39;email&#39;]}</span>\n\n<span
  class=\"s\">EOF</span>\n\ncat <span class=\"s\">&lt;&lt; EOF &gt;&gt; $1/settings.py</span>\n\n<span
  class=\"s\">AUTHENTICATION_BACKENDS = {</span>\n<span class=\"s\">    #&#39;YOUR_DJANGO_APP_NAME.auth0backend.Auth0&#39;,</span>\n<span
  class=\"s\">    &#39;django.contrib.auth.backends.ModelBackend&#39;</span>\n<span
  class=\"s\">}</span>\n\n<span class=\"s\">EOF</span>\n\ncat <span class=\"s\">&lt;&lt;
  EOF &gt;&gt; $1/settings.py</span>\n\n<span class=\"s\">LOGIN_URL = &#39;/login/auth0&#39;</span>\n<span
  class=\"s\">LOGIN_REDIRECT_URL = &#39;/dashboard&#39;</span>\n<span class=\"s\">EOF</span>\n\ncat
  &gt; auth0login/views.py<span class=\"s\">&lt;&lt;EOF</span>\n\n<span class=\"s\">from
  django.shortcuts import render, redirect</span>\n<span class=\"s\">from django.contrib.auth.decorators
  import login_required</span>\n<span class=\"s\">from django.contrib.auth import
  logout as log_out</span>\n<span class=\"s\">from django.conf import settings</span>\n<span
  class=\"s\">from django.http import HttpResponseRedirect</span>\n<span class=\"s\">from
  urllib.parse import urlencode</span>\n<span class=\"s\">import json</span>\n\n<span
  class=\"s\">def index(request):</span>\n<span class=\"s\">    user = request.user</span>\n<span
  class=\"s\">    if user.is_authenticated:</span>\n<span class=\"s\">        return
  redirect(dashboard)</span>\n<span class=\"s\">    else:</span>\n<span class=\"s\">
  \       return render(request, &#39;index.html&#39;)</span>\n\n\n<span class=\"s\">@login_required</span>\n<span
  class=\"s\">def dashboard(request):</span>\n<span class=\"s\">    user = request.user</span>\n<span
  class=\"s\">    auth0user = user.social_auth.get(provider=&#39;auth0&#39;)</span>\n<span
  class=\"s\">    userdata = {</span>\n<span class=\"s\">        &#39;user_id&#39;:
  auth0user.uid,</span>\n<span class=\"s\">        &#39;name&#39;: user.first_name,</span>\n<span
  class=\"s\">        &#39;picture&#39;: auth0user.extra_data[&#39;picture&#39;],</span>\n<span
  class=\"s\">        &#39;email&#39;: auth0user.extra_data[&#39;email&#39;],</span>\n<span
  class=\"s\">    }</span>\n\n<span class=\"s\">    return render(request, &#39;dashboard.html&#39;,
  {</span>\n<span class=\"s\">        &#39;auth0User&#39;: auth0user,</span>\n<span
  class=\"s\">        &#39;userdata&#39;: json.dumps(userdata, indent=4)</span>\n<span
  class=\"s\">    })</span>\n\n<span class=\"s\">def logout(request):</span>\n<span
  class=\"s\">    log_out(request)</span>\n<span class=\"s\">    return_to = urlencode({&#39;returnTo&#39;:
  request.build_absolute_uri(&#39;/&#39;)})</span>\n<span class=\"s\">    logout_url
  = &#39;https://%s/v2/logout?client_id=%s&amp;%s&#39; % \\</span>\n<span class=\"s\">
  \                (settings.SOCIAL_AUTH_AUTH0_DOMAIN, settings.SOCIAL_AUTH_AUTH0_KEY,
  return_to)</span>\n<span class=\"s\">    return HttpResponseRedirect(logout_url)</span>\n\n<span
  class=\"s\">EOF</span>\n\ncat <span class=\"s\">&lt;&lt; EOF &gt;&gt; auth0login/templates/index.html</span>\n\n<span
  class=\"s\">&lt;div class=&quot;login-box auth0-box before&quot;&gt;</span>\n<span
  class=\"s\">    &lt;img src=&quot;https://i.cloudup.com/StzWWrY34s.png&quot; /&gt;</span>\n<span
  class=\"s\">    &lt;h3&gt;Auth0 Example&lt;/h3&gt;</span>\n<span class=\"s\">    &lt;p&gt;Zero
  friction identity infrastructure, built for developers&lt;/p&gt;</span>\n<span class=\"s\">
  \   &lt;a class=&quot;btn btn-primary btn-lg btn-login btn-block&quot; href=&quot;/login/auth0&quot;&gt;Log
  In&lt;/a&gt;</span>\n<span class=\"s\">&lt;/div&gt;</span>\n<span class=\"s\">EOF</span>\n\ncat
  <span class=\"s\">&lt;&lt; EOF &gt;&gt; auth0login/templates/dashboard.html</span>\n\n<span
  class=\"s\">&lt;div class=&quot;logged-in-box auth0-box logged-in&quot;&gt;</span>\n<span
  class=\"s\">    &lt;h1 id=&quot;logo&quot;&gt;&lt;img src=&quot;//cdn.auth0.com/samples/auth0_logo_final_blue_RGB.png&quot;
  /&gt;&lt;/h1&gt;</span>\n<span class=\"s\">    &lt;img class=&quot;avatar&quot;
  src=&quot;{{ auth0User.extra_data.picture }}&quot;/&gt;</span>\n<span class=\"s\">
  \   &lt;h2&gt;Welcome {{ user.username }}&lt;/h2&gt;</span>\n<span class=\"s\">
  \   &lt;pre&gt;{{ userdata }}&lt;/pre&gt;</span>\n<span class=\"s\">&lt;/div&gt;</span>\n<span
  class=\"s\">EOF</span>\n\ncat <span class=\"s\">&lt;&lt; EOF &gt;&gt; auth0login/urls.py</span>\n<span
  class=\"s\">from django.urls import path, include</span>\n<span class=\"s\">from
  . import views</span>\n\n<span class=\"s\">urlpatterns = [</span>\n<span class=\"s\">
  \   path(&#39;&#39;, views.index),</span>\n<span class=\"s\">    path(&#39;dashboard&#39;,
  views.dashboard),</span>\n<span class=\"s\">    path(&#39;logout&#39;, views.logout),</span>\n<span
  class=\"s\">    path(&#39;&#39;, include(&#39;django.contrib.auth.urls&#39;)),</span>\n<span
  class=\"s\">    path(&#39;&#39;, include(&#39;social_django.urls&#39;)),</span>\n<span
  class=\"s\">]</span>\n\n<span class=\"s\">EOF</span>\n\npython manage.py makemigrations\npython
  manage.py migrate\n</pre></div>\n\n</pre>\n\n<h2>Conclusion</h2>\n<p>Ok, so this
  was it, a quite big script but that's how automation can be. We were able to set
  up a Django base application with a ready app of Auth0 to extend the functionality.
  This was just a basic script also you can extend the functionalities like adding
  a curl command to fetch the credentials and make it more automated but that was
  not the aim of this article.</p>\n<p>If you had any issues using the script please
  let me know, I'll be happy to fix those. Thanks for reading. Happy Coding :)</p>\n"
image_url: https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643287941/blogmedia/aggaqpjljlcysdopvoj3.png
long_description: 'This is a guide and a walkthrough of how to quickly set up a base
  Django project with Auth0 as integration for authentication and authorization. I
  will walk you through the Django setup and how to use and integrate the functionalities
  of the Auth0.  '
now: 2025-05-01 18:11:33.314175
path: blog/posts/2021-09-28-Django-Auth0-Quick-Setup.md
prevnext: null
slug: djagno-auth0-script
status: published
subtitle: A simple BASH script to set up a basic Django project with authentication
  integrated using Auth0.
tags:
- django
- bash
- python
templateKey: blog-post
title: Django + Auth0 Quick Setup
today: 2025-05-01
---

## Introduction


This is a guide and a walkthrough of how to quickly set up a base Django project with Auth0 as integration for authentication and authorization. I will walk you through the Django setup and how to use and integrate the functionalities of the Auth0.  I will also discuss how why you should be using Auth0 and why I love it.

The script takes `2:44` minutes time to do everything from scratch. From installing virtualenv in python to integrating the Auth0 application. 

Here's how the script works:

![authodj.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1632830813802/MOvedPYdt.gif)

### Contents

- [What is Auth0](#what-is-auth0)
- [Why I love Auth0](#why-i-love-auth0)
- [Set up a Django Project](#set-up-a-django-project)
- [Integrate Auth0 to a Django project](#integrate-auth0-to-a-django-project)
- [Creating a BASH Script for integrating Auth0](#creating-a-bash-script-for-integrating-auth0)
     - [Appending to a file](#appending-to-a-file)
     - [Adding text before a particular line using `sed` ](#adding-text-before-a-particular-line-using-sed)
     - [Appending to a line using `sed`](#appending-to-a-line-using-sed)
- [Complete BASH Script](#complete-bash-script)
- [Conclusion](#conclusion)

## What is Auth0

Auth0 (`Auth zero`) is a platform that provides easy authentication and authorization for a number of platforms in various programming languages and frameworks. The easy-to-follow documentation, availability for almost all web frameworks across platforms make it a big bonus for developers. They actually make the Developer experience flawless and beginner-friendly. 

According to Auth0,
> They make your login box awesome

And how true is that they make things pretty convenient and wicked fast to integrate a smooth functional backend for authentication and authorization. Of course, there are more things they offer than just making authentication systems but it is by far what the world knows them for.

## Why I love Auth0

Auth0 is a generous company that provides a free tier for a limited capacity of authentication and that might be more than enough for a developer getting its feet wet in the web development (backend). 

They even provide a nice user interface out of the box for login/signup and even a dashboard ready-made, which is quite a lot of heavy lifting already done for you. Also, there is a dashboard for analyzing the number of sign-in/logins into the particular app. This provides the admin/developer of the app to get a closer look at the user registered in a day/week/months, number of active users, and so on. 

So, who would not love it? I am willing to write and use their service for some of my projects. I already have used one for the Hashnode x Auth0 Hackathon, I made [devquotes](https://github.com/Mr-Destructive/devquotes) using the authentication of Auth0 in my Django application. 

## Set up a Django Project

If you are reading this you already know how to set up a Django project, I assume. But nevertheless, I can just include a quick introduction on how to do it. I have a script to do this. 

```bash
#!/usr/bin/env bash

mkdir $1
cd $1
pip install virtualenv
virtualenv env
source env\bin\activate

pip install django
django-admin startproject $1 .
clear
```
You can check out  [Django Quick Setup Script](https://techstructiveblog.hashnode.dev/django-quick-setup-script)  for the details of this script and also a more in-depth guide of Django project setup.

But if you want to understand the basics of the Django project setup here is a little guide about it:

Firstly, create a virtual environment, it's not mandatory but it keeps things simple and easy for your project in correspondence to the entire OS. So in python, we have a module to create the virtual environment pretty easily,

```shell
pip install virtualenv
```
You can use `pip3` or `pip -m`, or however you install normal python modules. This just installs the python virtual environment, we need to create one in the current folder, so for that navigate to the folder where you want to create the project and enter the following command:

```shell
virtualenv venv
``` 

Here, `venv` can be anything like `env` just for your understanding and simplicity it's a standard name kept for the same. After this, you will see a folder of the same name i.e. `venv` or any other name you have used. This is the folder where python will keep every installation private to the local folder itself. Now, we need to activate the virtual environment, for that we can use the command :

```bash
# for Linux/macOS :
source venv/bin/activate
```

```batch
# for Windows:
venv\Scripts\activate
```
After this, your command prompt will have a `(venv)` attached to its start. This indicates you are in a virtual environment, things you do here, may it be module installation or any configuration related to python will stay in the local folder itself.

After the virtual environment is set up and activated, you can install Django and get started with it. Firstly, install Django using pip:

```shell
pip install django
```
After the installation is completed, you can start a Django project in the current folder using the command: 

```shell
django-admin startproject name
```
Here name can be your project name. After this, you will see one new folder and one file pop up.
Namely, the `project named` folder and `manage.py` file. So you don't have to touch the `manage.py` file but we use it in most of the commands to use the Django functionalities. 

You can now run your basic server using the command : 
```shell
python manage.py runserver
```
There is a base installed/setup of the Django project. Moving on in integrating the Auth0 login functionality in our webpage.


## Integrate the Auth0 app in your project

So, for integrating the Auth0 app for your web application, you need to have an Auth0 account, you can signup here. After this you can create an Auth0 application for any type of application, we have a couple of options:

- Native Application
- Single Page Application
- Regular Web Application
- Machine to Machine Application

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1632748408741/sUeS8AgrX.png)

You can choose as per your needs, but mostly it would be a Regular Web application if you have a backend like Django, Nodejs, PHP, or other types of frameworks and languages. So, moving ahead we have created an application for the Django framework as a backend. Now, we have a `Settings` tab in the application dashboard, where we have different credentials for the Auth0 app to talk to our application.

The credentials needed to be stored safely are:
- domain
- Client ID (Key)
- Client Secret

This has to be secured for our local application which will go into production when ready. You can use several options like dotenv, environment variables, and so on when the application is being deployed but for now, let's hardcode them in our Django project. 

Now, you can follow the simple straightforward procedure to copy-paste your credentials from the  [Auth0 official documentation](https://auth0.com/docs/quickstart/webapp/django/01-login#logout). It's quite straightforward to follow the steps even for a beginner. 

After the Auth0 app has been configured following the procedure in the documentation, you need to integrate several files like dashboard and index templates into your custom templates.

Following additional changes are also to be made if you have a user-defined app for your Django project.

 In the `auth0login` app, `view.py` file:
1. The `index` function renders the base file for your project if the user is logged in.
2. The `dashboard` function renders the baked version of your profile/dashboard of users on your app.

You would also need to add the root URIs of your app that you will be using for testing or in production. For example, we can add `http://127.0.0.1:8000` to allow and use Auth0 in our development environment locally.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1632828981455/2gc4h7tTh.png)

You also need to specify the callback URLs for your application which is `/complete/auth0` by default.

## Creating a BASH Script for integrating Auth0

So, we can now dive into creating the BASH Script to set up the Django x Auth0 application in minutes. The script is quite large, like 200 lines but don't worry! Its automation reduces the pain of integrating a User Authorization flawlessly. I am also thinking of adding the `cURL` command and parsing in the Client ids, keys, and secret keys, etc. 

### Appending to a file

We can use the `cat` command to append text to a file, using the syntax as below:

```shell
cat << EOF >> filename
text
more text
EOF
```
Remember here EOF is just a label to stop the command and save it to the file. 


![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1632826339707/_g_RXP3NE.png)

So, we can see that we were able to append to a file, multiple lines using the cat command.  

We have used this concept in adding configuration and credentials to the `settings.py` or the `urls.py` files.

### Adding text before a particular line using `sed` 

`sed` is a great command, and there is nothing you can't do with it, OK there might be exceptions. We can get write to a file directly (don't display the output) and specify the line number before which we want to append the text. We can then add the text we want and followed by the filename.

```shell
sed -i '33 i sometext here' filename
```
Here, `33` is the line number in the file which we want to insert before. We have used `'"'` to add a `'` inside a `'`, this might feel a bit wired but that is how it is in BASH. 

Let's say you want to add `print('Hello, World!')` to a particular line, we have to enclose `'` with these `"'`( double and single quotes),

```shell
sed -i '2i print('"'Hello, World'"')' hello.py
```
This will add the line `print('Hello World')` to the file `hello.py`


![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1632824742835/Uj8AF07UG.png)

### Appending to a line using sed

We can even append text to a particular line using sed, we can use some escape characters and regex to add the text from the end of the line.

```shell
sed -i '2i s/$/ textgoes here /' filename
``` 
Here 2 is any number of line you want to add text to, next `i` a prompt for inserting text and then we have regex like `s/$/ /`, this will put the text enclosed in `/ /` to the end of the line as indicated by `$`.  We have the filename at its usual place as before.

So, lets say, I want to add a comment to the second line in the previous example, I can use the following command to do it:

```shell
sed -i '2 s/$/ # another comment/' hello.py

```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1632825067925/0eU2mkCDI.png)

We have used these commands to add the `include` function in the `urls.py` in the project folder.

So those were all the operations we used for doing some automated editing for the Auth0 app integration to our Django project.
 
Below is the entire script and is also uploaded on [GitHub](https://github.com/Mr-Destructive/django-auth0-quick-setup). 

## Complete BASH Script

You can run the file by parsing the name of your project.

```shell
bash script.sh mywebsite
```

Wait for some 2-3 minutes, and the script will produce the Django application with the Auth0 app integrated. You will have to enter the credentials manually wherever applicable.


```bash
#!/usr/bin/env bash

mkdir $1
cd $1
pip install virtualenv
virtualenv venv
source venv/Scripts/activate

pip install django
django-admin startproject $1 .

cat << EOF >> requirements.txt
social-auth-app-django~=3.1 
python-jose~=3.0 
python-dotenv~=0.9
EOF

pip install -r requirements.txt

pip freeze > requirements.txt

python manage.py startapp auth0login

touch auth0login/urls.py
mkdir auth0login/templates
touch auth0login/templates/index.html
touch auth0login/templates/dashboard.html

sed -i '40 i \    '"'"'social_django'"'"',' $1/settings.py
sed -i '41 i \    '"'"'auth0login'"'"',' $1/settings.py
sed -i '21 i \    path('"''"', include('"'auth0login.urls'"')),' $1/urls.py
sed -i '17 s/$/, include/' $1/urls.py 

cat << EOF >> $1/settings.py
SOCIAL_AUTH_TRAILING_SLASH = False  # Remove trailing slash from routes
SOCIAL_AUTH_AUTH0_DOMAIN = 'YOUR_DOMAIN'
SOCIAL_AUTH_AUTH0_KEY = 'YOUR_CLIENT_ID'
SOCIAL_AUTH_AUTH0_SECRET = 'YOUR_CLIENT_SECRET'
EOF

cat << EOF >> $1/settings.py 
SOCIAL_AUTH_AUTH0_SCOPE = [
    'openid',
    'profile',
    'email'
]
EOF

python manage.py migrate

cat << EOF >>auth0login/auth0backend.py

from urllib import request
from jose import jwt
from social_core.backends.oauth import BaseOAuth2


class Auth0(BaseOAuth2):
    """Auth0 OAuth authentication backend"""
    name = 'auth0'
    SCOPE_SEPARATOR = ' '
    ACCESS_TOKEN_METHOD = 'POST'
    REDIRECT_STATE = False
    EXTRA_DATA = [
        ('picture', 'picture'),
        ('email', 'email')
    ]

    def authorization_url(self):
        return 'https://' + self.setting('DOMAIN') + '/authorize'

    def access_token_url(self):
        return 'https://' + self.setting('DOMAIN') + '/oauth/token'

    def get_user_id(self, details, response):
        """Return current user id."""
        return details['user_id']

    def get_user_details(self, response):
        # Obtain JWT and the keys to validate the signature
        id_token = response.get('id_token')
        jwks = request.urlopen('https://' + self.setting('DOMAIN') + '/.well-known/jwks.json')
        issuer = 'https://' + self.setting('DOMAIN') + '/'
        audience = self.setting('KEY')  # CLIENT_ID
        payload = jwt.decode(id_token, jwks.read(), algorithms=['RS256'], audience=audience, issuer=issuer)

        return {'username': payload['nickname'],
                'first_name': payload['name'],
                'picture': payload['picture'],
                'user_id': payload['sub'],
                'email': payload['email']}

EOF

cat << EOF >> $1/settings.py

AUTHENTICATION_BACKENDS = {
    #'YOUR_DJANGO_APP_NAME.auth0backend.Auth0',
    'django.contrib.auth.backends.ModelBackend'
}

EOF

cat << EOF >> $1/settings.py

LOGIN_URL = '/login/auth0'
LOGIN_REDIRECT_URL = '/dashboard'
EOF

cat > auth0login/views.py<<EOF

from django.shortcuts import render, redirect
from django.contrib.auth.decorators import login_required
from django.contrib.auth import logout as log_out
from django.conf import settings
from django.http import HttpResponseRedirect
from urllib.parse import urlencode
import json

def index(request):
    user = request.user
    if user.is_authenticated:
        return redirect(dashboard)
    else:
        return render(request, 'index.html')


@login_required
def dashboard(request):
    user = request.user
    auth0user = user.social_auth.get(provider='auth0')
    userdata = {
        'user_id': auth0user.uid,
        'name': user.first_name,
        'picture': auth0user.extra_data['picture'],
        'email': auth0user.extra_data['email'],
    }

    return render(request, 'dashboard.html', {
        'auth0User': auth0user,
        'userdata': json.dumps(userdata, indent=4)
    })

def logout(request):
    log_out(request)
    return_to = urlencode({'returnTo': request.build_absolute_uri('/')})
    logout_url = 'https://%s/v2/logout?client_id=%s&%s' % \
                 (settings.SOCIAL_AUTH_AUTH0_DOMAIN, settings.SOCIAL_AUTH_AUTH0_KEY, return_to)
    return HttpResponseRedirect(logout_url)

EOF

cat << EOF >> auth0login/templates/index.html

<div class="login-box auth0-box before">
    <img src="https://i.cloudup.com/StzWWrY34s.png" />
    <h3>Auth0 Example</h3>
    <p>Zero friction identity infrastructure, built for developers</p>
    <a class="btn btn-primary btn-lg btn-login btn-block" href="/login/auth0">Log In</a>
</div>
EOF

cat << EOF >> auth0login/templates/dashboard.html

<div class="logged-in-box auth0-box logged-in">
    <h1 id="logo"><img src="//cdn.auth0.com/samples/auth0_logo_final_blue_RGB.png" /></h1>
    <img class="avatar" src="{{ auth0User.extra_data.picture }}"/>
    <h2>Welcome {{ user.username }}</h2>
    <pre>{{ userdata }}</pre>
</div>
EOF

cat << EOF >> auth0login/urls.py
from django.urls import path, include
from . import views

urlpatterns = [
    path('', views.index),
    path('dashboard', views.dashboard),
    path('logout', views.logout),
    path('', include('django.contrib.auth.urls')),
    path('', include('social_django.urls')),
]

EOF

python manage.py makemigrations
python manage.py migrate

```
## Conclusion

Ok, so this was it, a quite big script but that's how automation can be. We were able to set up a Django base application with a ready app of Auth0 to extend the functionality. This was just a basic script also you can extend the functionalities like adding a curl command to fetch the credentials and make it more automated but that was not the aim of this article. 

If you had any issues using the script please let me know, I'll be happy to fix those. Thanks for reading. Happy Coding :)