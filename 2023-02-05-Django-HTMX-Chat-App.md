---
article_html: "<h1>Django + HTMX Chat application</h1>\n<h2>Introduction</h2>\n<p>In
  this article, we will be creating a Django project, which will be a chat-room kind
  of application. The user needs to authenticate to the app and then and there he/she
  can create or join rooms, every room will have some name and URL associated with
  it. So, the user simply needs to enter the name of the room, which will be unique.
  The user can then simply enter the messages in the chat room. This is a core chat
  application that uses web sockets.</p>\n<p>The unique thing about this app will
  be that, we don't have to write a javascript client. It will all be handled by some
  HTMX magic. The web socket in the backend will definitely handle using Django channels.</p>\n<p>Demo:</p>\n<p><img
  src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-app-demo.webm\" alt=\"Demonstration
  of the Chat App\" /></p>\n<p><a href=\"https://github.com/Mr-Destructive/django-htmx-chat\">GitHub
  Repository</a></p>\n<h3>Requirements:</h3>\n<ul>\n<li>\n<p>Django</p>\n</li>\n<li>\n<p>Django-channels</p>\n</li>\n<li>\n<p>daphne</p>\n</li>\n<li>\n<p>HTMX</p>\n</li>\n<li>\n<p>SQLite
  or any relational database</p>\n</li>\n</ul>\n<p>Also if we want to use the application
  on a large and production scale:</p>\n<ul>\n<li>\n<p>Redis</p>\n</li>\n<li>\n<p>channels_redis</p>\n</li>\n</ul>\n<p>The
  code for this chat app is provided in the <a href=\"https://github.com/Mr-Destructive/django-htmx-chat\">GitHub
  repository</a>.</p>\n<h2>Setup for Django project</h2>\n<p>We will create a simple
  Django project to start with. The project can have two apps, one for auth and the
  other for the chat. You can customize the way you want your existing project accordingly.
  This project is just a demonstration of the use of the chat application with websockets
  and Django channels.</p>\n<p>I'll call the project <code>backchat</code>, you can
  call it whatever you want. We will create a virtual environment and install Django
  in that virtual environment</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>virtualenv
  .venv\n\nFor Linux/macOS:\n<span class=\"nb\">source</span> .venv/bin/activate\n\nFor
  Windows:\n.venv<span class=\"se\">\\s</span>cripts<span class=\"se\">\\a</span>ctivate\n\npip
  install django\ndjango-admin startproject backchat\n<span class=\"nb\">cd</span>
  backchat\n</pre></div>\n\n</pre>\n\n<p>This will set up a base Django project. We
  can now work on the actual implementation of the Django project. Firstly, we will
  start with authentication.</p>\n<h2>Adding basic Authentication and Authorization</h2>\n<h3>Creating
  the accounts app</h3>\n<p>We can separate the authentication of the user from the
  rest of the project, by creating a separate app called <code>user</code> or <code>accounts</code>
  app.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>python
  manage.py startapp accounts\n</pre></div>\n\n</pre>\n\n<h3>Creating a base user
  model</h3>\n<p>We'll start by inheriting the <a href=\"https://docs.djangoproject.com/en/4.1/topics/auth/customizing/#using-a-custom-user-model-when-starting-a-project\">AbstractUser</a>
  the model provided in the <code>django.contrib.auth.models</code> module. The model
  has base fields like <code>username</code> and <code>password</code> which are required
  fields, and <code>email</code>, <code>first_name</code>, <code>last_name</code>,
  etc. which are not mandatory. It is better to create a custom model by inheriting
  the <code>AbstractUser</code> because if in the longer run, we need to make custom
  fields in the user model, it becomes a breeze.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># accounts/models.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">djnago.contrib.auth.models</span> <span class=\"kn\">import</span>
  <span class=\"n\">AbstractUser</span>\n\n\n<span class=\"k\">class</span> <span
  class=\"nc\">User</span><span class=\"p\">(</span><span class=\"n\">AbstractUser</span><span
  class=\"p\">):</span>\n    <span class=\"k\">pass</span>\n</pre></div>\n\n</pre>\n\n<p>This
  creates a basic custom user rather than using the Django built-in user. Next, we
  need to make sure, Django understands the default user is the one we defined in
  the <code>accounts</code> app and not the default <code>User</code>. So, we just
  need to add a field called <code>AUTH_USER_MODEL</code> in the <code>settings.py</code>
  file. The value of this field will be the app name or the module name and the model
  in that python module that we need our custom user model to be set as the default
  user model.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># backchat/settings.py</span>\n\n\n<span class=\"n\">INSALLED_APPS</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"o\">...</span>\n
  \   <span class=\"o\">...</span>\n    <span class=\"s2\">&quot;accounts&quot;</span><span
  class=\"p\">,</span>\n<span class=\"p\">]</span>\n\n<span class=\"c1\"># Append
  to the end of file</span>\n<span class=\"n\">AUTH_USER_MODEL</span> <span class=\"o\">=</span>
  <span class=\"s1\">&#39;accounts.User&#39;</span>\n</pre></div>\n\n</pre>\n\n<h3>Initial
  migrations for the Django project</h3>\n<p>Now, this will get picked up as the default
  user model while referring to any processing related to the user. We can move into
  migrating the changes for the basic Django project and the user model.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>python
  manage.py makemigrations\npython manage.py migrate\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-accounts-migrations.png\" alt=\"initial
  migration for base django and user model\" /></p>\n<h3>Creating register view</h3>\n<p>Further,
  we can add the views like register and profile for the accounts app that can be
  used for the basic authentication. The Login and Logout views are provided in the
  <code>django.contrib.auth.views</code> module, we only have to write our own register
  view. I will be using function-based views to keep it simple to understand but it
  can simply be a class-based view as well.</p>\n<p>To define a view first, we need
  to have form, a user registration form. The form will inherit from the <a href=\"https://docs.djangoproject.com/en/4.1/topics/auth/default/#django.contrib.auth.forms.UserCreationForm\">UserCreationForm</a>
  form which will make the bulk of the task for us at the user registration. We can
  simply then override the Meta class with the fields that we want to display, so
  hence we just keep the <code>username</code> and the <code>password</code> fields.
  The form can be customized by adding in the widget attribute and specifying the
  classes used in them.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># accounts/forms.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">accounts.models</span> <span class=\"kn\">import</span> <span class=\"n\">User</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.contrib.auth.forms</span> <span
  class=\"kn\">import</span> <span class=\"n\">UserCreationForm</span>\n\n<span class=\"k\">class</span>
  <span class=\"nc\">UserRegisterForm</span><span class=\"p\">(</span><span class=\"n\">UserCreationForm</span><span
  class=\"p\">):</span>\n\n    <span class=\"k\">class</span> <span class=\"nc\">Meta</span><span
  class=\"p\">:</span>\n        <span class=\"n\">model</span><span class=\"o\">=</span>
  <span class=\"n\">User</span>\n        <span class=\"n\">fields</span> <span class=\"o\">=</span>
  <span class=\"p\">[</span><span class=\"s1\">&#39;username&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;password1&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;password2&#39;</span><span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>This will give us the <code>UserRegisterForm</code>
  form that can be used for displaying in the register view that we will create in
  the next step.</p>\n<p>We will have to create the register view that will render
  the form for user registration and will also process the form submission.</p>\n<pre
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># accounts/views.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.contrib</span> <span class=\"kn\">import</span> <span class=\"n\">messages</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.shortcuts</span> <span class=\"kn\">import</span>
  <span class=\"n\">redirect</span><span class=\"p\">,</span> <span class=\"n\">render</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">accounts.forms</span> <span class=\"kn\">import</span>
  <span class=\"n\">UserRegisterForm</span>\n\n<span class=\"k\">def</span> <span
  class=\"nf\">register</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">):</span>\n    <span class=\"k\">if</span> <span class=\"n\">request</span><span
  class=\"o\">.</span><span class=\"n\">method</span> <span class=\"o\">==</span>
  <span class=\"s2\">&quot;POST&quot;</span><span class=\"p\">:</span>\n        <span
  class=\"n\">form</span> <span class=\"o\">=</span> <span class=\"n\">UserRegisterForm</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">POST</span><span class=\"p\">)</span>\n        <span class=\"k\">if</span>
  <span class=\"n\">form</span><span class=\"o\">.</span><span class=\"n\">is_valid</span><span
  class=\"p\">():</span>\n            <span class=\"n\">form</span><span class=\"o\">.</span><span
  class=\"n\">save</span><span class=\"p\">()</span>\n            <span class=\"n\">username</span>
  <span class=\"o\">=</span> <span class=\"n\">form</span><span class=\"o\">.</span><span
  class=\"n\">cleaned_data</span><span class=\"o\">.</span><span class=\"n\">get</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;username&quot;</span><span class=\"p\">)</span>\n
  \           <span class=\"n\">messages</span><span class=\"o\">.</span><span class=\"n\">success</span><span
  class=\"p\">(</span>\n                <span class=\"n\">request</span><span class=\"p\">,</span>
  <span class=\"sa\">f</span><span class=\"s2\">&quot;Your account has been created!
  You are now able to log in&quot;</span>\n            <span class=\"p\">)</span>\n
  \           <span class=\"k\">return</span> <span class=\"n\">redirect</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;login&quot;</span><span class=\"p\">)</span>\n
  \   <span class=\"k\">else</span><span class=\"p\">:</span>\n        <span class=\"n\">form</span>
  <span class=\"o\">=</span> <span class=\"n\">UserRegisterForm</span><span class=\"p\">()</span>\n
  \       <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s2\">&quot;accounts/register.html&quot;</span><span
  class=\"p\">,</span> <span class=\"p\">{</span><span class=\"s2\">&quot;form&quot;</span><span
  class=\"p\">:</span> <span class=\"n\">form</span><span class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above register view has two cases, one for the user requesting the registration
  form and the second request when the user submits the form. So, when the user makes
  a get request, we load an empty form <code>UserRegisterForm</code> and render the
  <code>register</code> template with the form. We will make the templates later.</p>\n<p>So,
  the template is just rendered when the user wants to register and when the user
  submits the form(sends a post request) we get the details from the post request
  and make it an instance of <code>UserRegisterForm</code> and save the form if it
  is valid. We then redirect the user to the login view (will use the default one
  in the next section). We parse the message as the indication that the user was created.</p>\n<h3>Adding
  URLs for Authentication and Authorization</h3>\n<p>Once, we have the register view
  setup, we can also add login and logout views in the app. But, we don't have to
  write it ourselves, we can override them if needed, but we'll keep the default ones.
  We will use the <a href=\"https://docs.djangoproject.com/en/4.1/topics/auth/default/#django.contrib.auth.views.LoginView\">LoginView</a>
  and the <a href=\"https://docs.djangoproject.com/en/4.1/topics/auth/default/#django.contrib.auth.views.LogoutView\">LogoutView</a>
  view which are class-based views provided in the <code>django.contrib.auth.views</code>
  module. We will provide the respective templates for each of these views.</p>\n<pre
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># accounts/urls.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.urls</span> <span class=\"kn\">import</span> <span class=\"n\">path</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.contrib.auth</span> <span class=\"kn\">import</span>
  <span class=\"n\">views</span> <span class=\"k\">as</span> <span class=\"n\">auth_views</span>\n<span
  class=\"kn\">import</span> <span class=\"nn\">user.views</span> <span class=\"k\">as</span>
  <span class=\"nn\">user_views</span>\n\n<span class=\"n\">urlpatterns</span> <span
  class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;register/&quot;</span><span class=\"p\">,</span>
  <span class=\"n\">user_views</span><span class=\"o\">.</span><span class=\"n\">register</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;register&quot;</span><span class=\"p\">),</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span>\n        <span class=\"s2\">&quot;login/&quot;</span><span
  class=\"p\">,</span>\n        <span class=\"n\">auth_views</span><span class=\"o\">.</span><span
  class=\"n\">LoginView</span><span class=\"o\">.</span><span class=\"n\">as_view</span><span
  class=\"p\">(</span><span class=\"n\">template_name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;accounts/login.html&quot;</span><span class=\"p\">),</span>\n
  \       <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s2\">&quot;login&quot;</span><span
  class=\"p\">,</span>\n    <span class=\"p\">),</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span>\n        <span class=\"s2\">&quot;logout/&quot;</span><span
  class=\"p\">,</span>\n        <span class=\"n\">auth_views</span><span class=\"o\">.</span><span
  class=\"n\">LogoutView</span><span class=\"o\">.</span><span class=\"n\">as_view</span><span
  class=\"p\">(</span><span class=\"n\">template_name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;accounts/logout.html&quot;</span><span class=\"p\">),</span>\n
  \       <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s2\">&quot;logout&quot;</span><span
  class=\"p\">,</span>\n    <span class=\"p\">),</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>We
  have named the URLs as <code>register</code>, <code>login</code>, and <code>logout</code>
  so that we can use them while rendering the links for them in the templates. Now,
  we also need to include the URLs from the accounts app in the project URLs. We can
  do that by using the <code>include</code> method and specifying the app name with
  the module where the urlpatterns are located.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># backchat/urls.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.contrib</span> <span class=\"kn\">import</span> <span class=\"n\">admin</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.urls</span> <span class=\"kn\">import</span>
  <span class=\"n\">include</span><span class=\"p\">,</span> <span class=\"n\">path</span>\n\n<span
  class=\"n\">urlpatterns</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;admin/&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">admin</span><span class=\"o\">.</span><span
  class=\"n\">site</span><span class=\"o\">.</span><span class=\"n\">urls</span><span
  class=\"p\">),</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;auth/&quot;</span><span class=\"p\">,</span> <span class=\"n\">include</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;accounts.urls&quot;</span><span class=\"p\">)),</span>\n<span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>So, we have routed the <code>/auth</code>
  path to include all the URLs in the accounts app. So, the login view will be at
  the URL <code>/auth/login/</code> , and so on.</p>\n<p>Also, we need to add the
  <code>LOGIN_REDIRECT_URL</code> and <code>LOGIN_URL</code> for specifying the url
  name which will be redirected once the user is logged in and the default login url
  name respectively.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># backchat / settings.py</span>\n\n\n<span class=\"n\">LOGIN_REDIRECT_URL</span>
  <span class=\"o\">=</span> <span class=\"s2\">&quot;index&quot;</span>\n<span class=\"n\">LOGIN_URL</span>
  <span class=\"o\">=</span> <span class=\"s2\">&quot;login&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>We
  are now almost done with the view and routing part of the accounts app and can move
  into the creation of templates.</p>\n<h3>Adding Templates for authentication views</h3>\n<p>We
  need a few templates that we have been referencing in the views and the URLs of
  the accounts app in the project. So there are a couple of ways to use templates
  in a Django project. I prefer to have a single templates folder in the root of the
  project and have subfolders as the app which will have the templates specific to
  those apps.</p>\n<p>I usually create a <code>base.html</code> file in the templates
  folder and use that for inheriting other templates. So, we will have to change one
  setting in the project to make sure it looks at <code>templates/</code> from the
  root of the project.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># backchat/settings.py</span>\n\nimport os\n\n...\n...\n\n<span class=\"nv\">TEMPLATES</span>
  <span class=\"o\">=</span> <span class=\"o\">[</span>\n    <span class=\"o\">{</span>\n
  \       <span class=\"s2\">&quot;BACKEND&quot;</span>: <span class=\"s2\">&quot;django.template.backends.django.DjangoTemplates&quot;</span>,\n
  \       <span class=\"s2\">&quot;DIRS&quot;</span>: <span class=\"o\">[</span> os.path.join<span
  class=\"o\">(</span>BASE_DIR, <span class=\"s2\">&quot;templates&quot;</span><span
  class=\"o\">)</span>, <span class=\"o\">]</span>,\n        <span class=\"s2\">&quot;APP_DIRS&quot;</span>:
  True,\n        <span class=\"s2\">&quot;OPTIONS&quot;</span>: <span class=\"o\">{</span>\n
  \           <span class=\"s2\">&quot;context_processors&quot;</span>: <span class=\"o\">[</span>\n
  \               <span class=\"s2\">&quot;django.template.context_processors.debug&quot;</span>,\n
  \               <span class=\"s2\">&quot;django.template.context_processors.request&quot;</span>,\n
  \               <span class=\"s2\">&quot;django.contrib.auth.context_processors.auth&quot;</span>,\n
  \               <span class=\"s2\">&quot;django.contrib.messages.context_processors.messages&quot;</span>,\n
  \           <span class=\"o\">]</span>,\n        <span class=\"o\">}</span>,\n    <span
  class=\"o\">}</span>,\n<span class=\"o\">]</span>\n</pre></div>\n\n</pre>\n\n<p>Then
  create the folder in the same path as you see your <code>manage.py</code> file.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>mkdir
  templates\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-mkdir-templates.png\"
  alt=\"Template Set Up\" /></p>\n<h4>Creating the base template</h4>\n<p>The below
  will be the base template used for the chat application, you can custmize it as
  per your needs.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cp\">&lt;!DOCTYPE html&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">meta</span>
  <span class=\"na\">charset</span><span class=\"o\">=</span><span class=\"s\">&quot;utf-8&quot;</span>
  <span class=\"p\">/&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>Chat App<span class=\"p\">&lt;/</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>\n        {% load static %}\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">script</span> <span class=\"na\">src</span><span class=\"o\">=</span><span
  class=\"s\">&quot;https://unpkg.com/htmx.org@1.8.5&quot;</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">script</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span
  class=\"nt\">head</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">body</span><span class=\"p\">&gt;</span>\n        {% if user.is_authenticated
  %}\n            <span class=\"p\">&lt;</span><span class=\"nt\">a</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% url &#39;logout&#39; %}&quot;</span><span
  class=\"p\">&gt;</span>Logout<span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n        {% else %}\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% url &#39;login&#39; %}&quot;</span><span class=\"p\">&gt;</span>Login<span
  class=\"p\">&lt;/</span><span class=\"nt\">a</span><span class=\"p\">&gt;</span>\n
  \       {% endif %}\n        <span class=\"p\">&lt;</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>Back Chat<span class=\"p\">&lt;/</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>\n        {% block base %}\n        {% endblock %}\n    <span
  class=\"p\">&lt;/</span><span class=\"nt\">body</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;/</span><span class=\"nt\">html</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>We
  have included the <a href=\"https://htmx.org/docs/#installing\">htmx</a> package
  as the script into this template as we will be using it in the actual part of chat
  application.</p>\n<h4>Creating the Register Template</h4>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>#
  templates / accounts / register.html\n\n\n{% extends &#39;base.html&#39; %}\n{%
  block base %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">div</span> <span
  class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;content-section&quot;</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">form</span>
  <span class=\"na\">method</span><span class=\"o\">=</span><span class=\"s\">&quot;POST&quot;</span><span
  class=\"p\">&gt;</span>\n            {% csrf_token %}\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">fieldset</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;form-group&quot;</span><span class=\"p\">&gt;</span>\n                <span
  class=\"p\">&lt;</span><span class=\"nt\">legend</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;border-bottom mb-4&quot;</span><span
  class=\"p\">&gt;</span>Register Now<span class=\"p\">&lt;/</span><span class=\"nt\">legend</span><span
  class=\"p\">&gt;</span>\n                {{ form.as_p }}\n            <span class=\"p\">&lt;/</span><span
  class=\"nt\">fieldset</span><span class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">div</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;form-group&quot;</span><span class=\"p\">&gt;</span>\n                <span
  class=\"p\">&lt;</span><span class=\"nt\">button</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;btn btn-outline-info&quot;</span> <span
  class=\"na\">type</span><span class=\"o\">=</span><span class=\"s\">&quot;submit&quot;</span><span
  class=\"p\">&gt;</span>Sign Up<span class=\"p\">&lt;/</span><span class=\"nt\">button</span><span
  class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;/</span><span class=\"nt\">form</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">div</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;border-top
  pt-3&quot;</span><span class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">small</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;text-muted&quot;</span><span class=\"p\">&gt;</span>\n\t\t    Already
  Have An Account? <span class=\"p\">&lt;</span><span class=\"nt\">a</span> <span
  class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;ml-2&quot;</span>
  <span class=\"na\">href</span><span class=\"o\">=</span><span class=\"s\">&quot;{%
  url &#39;login&#39; %}&quot;</span><span class=\"p\">&gt;</span>Log In<span class=\"p\">&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;/</span><span
  class=\"nt\">small</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;/</span><span
  class=\"nt\">div</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span
  class=\"nt\">div</span><span class=\"p\">&gt;</span>\n{% endblock %}\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-register-page.png\" alt=\"User
  Registration Page\" /></p>\n<h4>Creating the Login Template</h4>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>#
  templates / accounts / login.html    \n\n\n{% extends &#39;base.html&#39; %}\n{%
  block base %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">div</span> <span
  class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;content-section&quot;</span>
  <span class=\"na\">id</span><span class=\"o\">=</span><span class=\"s\">&quot;login&quot;</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">form</span>
  <span class=\"na\">method</span><span class=\"o\">=</span><span class=\"s\">&quot;POST&quot;</span>
  <span class=\"p\">&gt;</span>\n            {% csrf_token %}\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">fieldset</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;form-group&quot;</span><span class=\"p\">&gt;</span>\n                <span
  class=\"p\">&lt;</span><span class=\"nt\">legend</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;border-bottom mb-4&quot;</span><span
  class=\"p\">&gt;</span>LOG IN<span class=\"p\">&lt;/</span><span class=\"nt\">legend</span><span
  class=\"p\">&gt;</span>\n                {{ form.as_p }}\n            <span class=\"p\">&lt;/</span><span
  class=\"nt\">fieldset</span><span class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">div</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;form-group&quot;</span><span class=\"p\">&gt;</span>\n                <span
  class=\"p\">&lt;</span><span class=\"nt\">button</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;btn btn-outline-info&quot;</span> <span
  class=\"na\">type</span><span class=\"o\">=</span><span class=\"s\">&quot;submit&quot;</span><span
  class=\"p\">&gt;</span>Log In<span class=\"p\">&lt;/</span><span class=\"nt\">button</span><span
  class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;/</span><span class=\"nt\">form</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">div</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;border-top
  pt-3&quot;</span><span class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">small</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;text-muted&quot;</span><span class=\"p\">&gt;</span>\n                Register
  Here <span class=\"p\">&lt;</span><span class=\"nt\">a</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;ml-2&quot;</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% url &#39;register&#39; %}&quot;</span><span
  class=\"p\">&gt;</span>Sign Up<span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;/</span><span class=\"nt\">small</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n{% endblock %}\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-login-page.png\"
  alt=\"User Login Page\" /></p>\n<p>Creating the Logout Template</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>#
  templates / accounts / logout.html    \n\n\n{% extends &#39;base.html&#39; %}\n{%
  block base %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span
  class=\"p\">&gt;</span>You have been logged out<span class=\"p\">&lt;/</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">div</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;border-top pt-3&quot;</span><span class=\"p\">&gt;</span>\n        <span
  class=\"p\">&lt;</span><span class=\"nt\">small</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;text-muted&quot;</span><span class=\"p\">&gt;</span>\n
  \           <span class=\"p\">&lt;</span><span class=\"nt\">a</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% url &#39;login&#39; %}&quot;</span><span
  class=\"p\">&gt;</span>Log In Again<span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;/</span><span class=\"nt\">small</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n{% endblock %}\n</pre></div>\n\n</pre>\n\n<h2>Install and
  setup channels</h2>\n<p>We will be using channels to create long-running connections,
  it is a wrapper around Django's asynchronous components and allows us to incorporate
  other protocols like web sockets and other asynchronous protocols.</p>\n<p>So, we
  will be using the Django channels package that will allow us to use the WebSocket
  protocol in the chat application. <a href=\"https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API\">WebSocket</a>
  is a communication protocol(set of rules and standards to be followed) that allows
  both the client and server to send and receive messages(communication).</p>\n<p>To
  install Django channels, we can use pip and install channels with daphne which will
  be used as the web server gateway interface for asynchronous web applications.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>pip
  install -U channels<span class=\"o\">[</span><span class=\"s2\">&quot;daphne&quot;</span><span
  class=\"o\">]</span>\n</pre></div>\n\n</pre>\n\n<p>So this will install the latest
  version of channels and daphne. We now have <a href=\"https://channels.readthedocs.io/en/stable/installation.html\">channels</a>
  in our Django project, we simply need to configure the <code>CHANNEL_LAYER</code>
  config for specifying the backend used that can be <code>Redis</code>, <code>In-Memory</code>,
  or others. We need to add <code>channels</code> , <code>daphne</code> to the <code>INSALLED_APPS</code>
  config of the project. Make sure the <code>daphne</code> app is on top of the applications
  list.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># backchat/settings.py</span>\n\n...\n...\n\n<span class=\"nv\">INSALLED_APPS</span>
  <span class=\"o\">=</span> <span class=\"o\">[</span>\n    <span class=\"s2\">&quot;daphne&quot;</span>,\n
  \   ...\n    ...\n    <span class=\"s2\">&quot;channels&quot;</span>,\n<span class=\"o\">]</span>\n\n\n<span
  class=\"nv\">ASGI_APPLICATION</span> <span class=\"o\">=</span> <span class=\"s2\">&quot;backchat.asgi.application&quot;</span>\n\n...\n...\n\n<span
  class=\"c1\"># For InMemory channels</span>\n\n<span class=\"nv\">CHANNEL_LAYERS</span>
  <span class=\"o\">=</span> <span class=\"o\">{</span>\n    <span class=\"s1\">&#39;default&#39;</span>:
  <span class=\"o\">{</span>\n        <span class=\"s2\">&quot;BACKEND&quot;</span>:
  <span class=\"s2\">&quot;channels.layers.InMemoryChannelLayer&quot;</span>,\n    <span
  class=\"o\">}</span>\n<span class=\"o\">}</span>\n\n\n<span class=\"c1\"># For Redis</span>\n\n<span
  class=\"nv\">CHANNEL_LAYERS</span> <span class=\"o\">=</span> <span class=\"o\">{</span>\n
  \   <span class=\"s2\">&quot;default&quot;</span>: <span class=\"o\">{</span>\n
  \       <span class=\"s2\">&quot;BACKEND&quot;</span>: <span class=\"s2\">&quot;asgi_redis.RedisChannelLayer&quot;</span>,\n
  \       <span class=\"s2\">&quot;CONFIG&quot;</span>: <span class=\"o\">{</span>\n
  \           <span class=\"s2\">&quot;hosts&quot;</span>: <span class=\"o\">[(</span><span
  class=\"s2\">&quot;redis-host-url&quot;</span>, <span class=\"m\">6379</span><span
  class=\"o\">)]</span>,\n        <span class=\"o\">}</span>,\n    <span class=\"o\">}</span>,\n<span
  class=\"o\">}</span>\n</pre></div>\n\n</pre>\n\n<h3>Redis Configuration (Optional)</h3>\n<p>You
  can either use the <a href=\"https://channels.readthedocs.io/en/latest/topics/channel_layers.html\">InMemoryChannelLayer</a>
  or you can use them <code>RedisChannelLayer</code> for the backend of the chat app.
  There are other types of backends like <code>Amazon SQS</code> services, <code>RabbitMQ</code>,
  <code>Kafka</code>, <code>Google Cloud Pub/Sub</code>, etc. I will be creating the
  app with only the <code>InMemoryChannelLayer</code> but will provide a guide for
  redis as well, both are quite similar and only have a few nuances.</p>\n<p>We need
  to install <a href=\"https://github.com/django/channels_redis/\">channels_redis</a>
  it for integrating redis in the Django project with channels.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>pip
  install channels_redis\n</pre></div>\n\n</pre>\n\n<p>So, this will make the <code>channels_redis</code>
  package available in the project, we use this package for real-time storage, in
  the case of the chat app, we might use it for storing messages or room details,
  etc.</p>\n<h2>Creating the Chat App</h2>\n<p>Further, we now can create another
  app for handling the rooms and chat application logic. This app will have its own
  models, views, and URLs. Also, we will define consumers and routers just like views
  and URLs but for asynchronous connections. More on that soon.</p>\n<p>So, let's
  create the <code>chat</code> app.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>python
  manage.py startapp chat\n</pre></div>\n\n</pre>\n\n<p>Then we will add the chat
  app to the <code>INSALLED_APPS</code> config.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># backchat/settings.py</span>\n\n<span class=\"n\">INSALLED_APPS</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"o\">...</span>\n
  \   <span class=\"o\">...</span><span class=\"p\">,</span>\n    <span class=\"s2\">&quot;chat&quot;</span><span
  class=\"p\">,</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  this will make sure to load the chat app in the project. Whenever we run any commands
  like migrations or running the server, it keeps the app in the <code>INSALLED_APPS</code>
  checked up.</p>\n<h3>Defining models</h3>\n<p>This is optional, but we'll do it
  for since we are making a Django app. We already have an auth system configured,
  adding rooms and authorizing the users will become easier then.</p>\n<p>So, let's
  create the models for the chat app which will be the <code>Room</code>.</p>\n<pre
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># chat/models.py</span>\n\n\n<span class=\"kn\">from</span> <span class=\"nn\">django.db</span>
  <span class=\"kn\">import</span> <span class=\"n\">models</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">accounts.models</span> <span class=\"kn\">import</span> <span
  class=\"n\">User</span>\n\n<span class=\"k\">class</span> <span class=\"nc\">Room</span><span
  class=\"p\">(</span><span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">Model</span><span class=\"p\">):</span>\n    <span class=\"n\">name</span>
  <span class=\"o\">=</span> <span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">CharField</span><span class=\"p\">(</span><span class=\"n\">max_length</span><span
  class=\"o\">=</span><span class=\"mi\">128</span><span class=\"p\">)</span>\n    <span
  class=\"n\">slug</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">SlugField</span><span class=\"p\">(</span><span
  class=\"n\">unique</span><span class=\"o\">=</span><span class=\"kc\">True</span><span
  class=\"p\">)</span>\n    <span class=\"n\">users</span> <span class=\"o\">=</span>
  <span class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">ManyToManyField</span><span
  class=\"p\">(</span><span class=\"n\">User</span><span class=\"p\">)</span>\n\n
  \   <span class=\"k\">def</span> <span class=\"fm\">__str__</span><span class=\"p\">(</span><span
  class=\"bp\">self</span><span class=\"p\">):</span>\n        <span class=\"k\">return</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">name</span>\n\n\n<span
  class=\"k\">class</span> <span class=\"nc\">Message</span><span class=\"p\">(</span><span
  class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">Model</span><span
  class=\"p\">):</span>\n    <span class=\"n\">room</span> <span class=\"o\">=</span>
  <span class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">ForeignKey</span><span
  class=\"p\">(</span><span class=\"n\">Room</span><span class=\"p\">,</span> <span
  class=\"n\">on_delete</span><span class=\"o\">=</span><span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">CASCADE</span><span class=\"p\">)</span>\n
  \   <span class=\"n\">user</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">ForeignKey</span><span class=\"p\">(</span><span
  class=\"n\">User</span><span class=\"p\">,</span> <span class=\"n\">on_delete</span><span
  class=\"o\">=</span><span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">CASCADE</span><span class=\"p\">)</span>\n    <span class=\"n\">message</span>
  <span class=\"o\">=</span> <span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">TextField</span><span class=\"p\">()</span>\n    <span class=\"n\">created_at</span>
  <span class=\"o\">=</span> <span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">DateTimeField</span><span class=\"p\">(</span><span class=\"n\">auto_now_add</span><span
  class=\"o\">=</span><span class=\"kc\">True</span><span class=\"p\">)</span>\n\n
  \   <span class=\"k\">def</span> <span class=\"fm\">__str__</span><span class=\"p\">(</span><span
  class=\"bp\">self</span><span class=\"p\">):</span>\n        <span class=\"k\">return</span>
  <span class=\"p\">(</span>\n            <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room</span><span class=\"o\">.</span><span class=\"n\">name</span> <span
  class=\"o\">+</span> <span class=\"s2\">&quot; - &quot;</span> <span class=\"o\">+</span>\n
  \           <span class=\"nb\">str</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">user</span><span class=\"o\">.</span><span
  class=\"n\">username</span><span class=\"p\">)</span> <span class=\"o\">+</span>
  <span class=\"s2\">&quot; : &quot;</span> <span class=\"o\">+</span>\n            <span
  class=\"nb\">str</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">message</span><span class=\"p\">)</span>\n
  \       <span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>So, we simply have
  the name which will be taken from the user, and the slug which will be served as
  the URL to the room. The users are set as <a href=\"https://docs.djangoproject.com/en/4.1/ref/models/fields/#django.db.models.ManyToManyField\">ManyToManyField</a>
  since one room can have multiple users and a user can be in multiple rooms. Also
  we are creating the model <code>Message</code> that will be storing the room and
  the user as the foreign key and the actual text as the message, we can improve the
  security by encrypting the text but it's not the point of this article.</p>\n<p>We
  have set the <code>created_at</code> the field which will set the time when the
  object was created. Finally, the dunder string method is used for representing the
  message object as a price of the concatenation of strings of room name, username,
  and the message. This is useful for admin stuff as it makes it easier to read the
  object rather than the default id.</p>\n<p>Now, once the models are designed we
  can migrate the models into the database.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>python manage.py makemigrations\npython
  manage.py migrate\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-migrations.png\"
  alt=\"Chat app migrations\" /></p>\n<p>We now have a decent model structure for
  the chat application. We can now start the crux of the application i.e. consumers
  and routing with channels.</p>\n<h3>Writing consumers and routers for WebSockets</h3>\n<p>So,
  we start with the bare bones provided in the tutorial on the channel <a href=\"https://channels.readthedocs.io/en/stable/tutorial/part_3.html\">documentation</a>.
  We create a class (consumer) called <code>ChatConsumer</code> which inherits from
  the <code>AsyncWebsocketConsumer</code> provided by the <code>channels.generic.websocket</code>
  module. This has a few methods like <code>connect</code>, <code>disconnect</code>,
  and <code>receive</code>. These are the building blocks of a consumer. We define
  these methods as they will be used for communication via the WebSocket protocol
  through the channels interface.</p>\n<p>In the following block of code, we are essentially
  doing the following:</p>\n<ul>\n<li>\n<p>Accepting connection on the requested room
  name</p>\n</li>\n<li>\n<p>Sending and Receiving messages on the room/group</p>\n</li>\n<li>\n<p>Closing
  the WebSocket connection and removing the client from the room/group</p>\n</li>\n</ul>\n<pre
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># chat/consumers.py</span>\n\n\n<span class=\"kn\">import</span> <span
  class=\"nn\">json</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">asgiref.sync</span>
  <span class=\"kn\">import</span> <span class=\"n\">sync_to_async</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">channels.generic.websocket</span> <span class=\"kn\">import</span>
  <span class=\"n\">AsyncWebsocketConsumer</span>\n\n<span class=\"kn\">from</span>
  <span class=\"nn\">chat.models</span> <span class=\"kn\">import</span> <span class=\"n\">Room</span><span
  class=\"p\">,</span> <span class=\"n\">Message</span>\n\n\n<span class=\"k\">class</span>
  <span class=\"nc\">ChatConsumer</span><span class=\"p\">(</span><span class=\"n\">AsyncWebsocketConsumer</span><span
  class=\"p\">):</span>\n    <span class=\"k\">async</span> <span class=\"k\">def</span>
  <span class=\"nf\">connect</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"p\">):</span>\n        <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_name</span> <span class=\"o\">=</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">scope</span><span class=\"p\">[</span><span
  class=\"s2\">&quot;url_route&quot;</span><span class=\"p\">][</span><span class=\"s2\">&quot;kwargs&quot;</span><span
  class=\"p\">][</span><span class=\"s2\">&quot;room_slug&quot;</span><span class=\"p\">]</span>\n
  \       <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">room_group_name</span>
  <span class=\"o\">=</span> <span class=\"sa\">f</span><span class=\"s2\">&quot;chat_</span><span
  class=\"si\">{</span><span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_name</span><span class=\"si\">}</span><span class=\"s2\">&quot;</span>\n
  \       <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">user</span>
  <span class=\"o\">=</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">scope</span><span class=\"p\">[</span><span class=\"s2\">&quot;user&quot;</span><span
  class=\"p\">]</span>\n\n        <span class=\"k\">await</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">channel_layer</span><span class=\"o\">.</span><span
  class=\"n\">group_add</span><span class=\"p\">(</span>\n            <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">room_group_name</span><span class=\"p\">,</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">channel_name</span>\n
  \       <span class=\"p\">)</span>\n\n        <span class=\"k\">await</span> <span
  class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">accept</span><span
  class=\"p\">()</span>\n\n    <span class=\"k\">async</span> <span class=\"k\">def</span>
  <span class=\"nf\">disconnect</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"p\">,</span> <span class=\"n\">close_code</span><span class=\"p\">):</span>\n
  \       <span class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">channel_layer</span><span class=\"o\">.</span><span class=\"n\">group_discard</span><span
  class=\"p\">(</span>\n            <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_group_name</span><span class=\"p\">,</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">channel_name</span>\n        <span class=\"p\">)</span>\n\n
  \   <span class=\"k\">async</span> <span class=\"k\">def</span> <span class=\"nf\">receive</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"p\">,</span> <span
  class=\"n\">text_data</span><span class=\"p\">):</span>\n        <span class=\"n\">text_data_json</span>
  <span class=\"o\">=</span> <span class=\"n\">json</span><span class=\"o\">.</span><span
  class=\"n\">loads</span><span class=\"p\">(</span><span class=\"n\">text_data</span><span
  class=\"p\">)</span>\n        <span class=\"n\">message</span> <span class=\"o\">=</span>
  <span class=\"n\">text_data_json</span><span class=\"p\">[</span><span class=\"s2\">&quot;message&quot;</span><span
  class=\"p\">]</span>\n        <span class=\"n\">username</span> <span class=\"o\">=</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">user</span><span
  class=\"o\">.</span><span class=\"n\">username</span>\n        \n        <span class=\"k\">await</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">channel_layer</span><span
  class=\"o\">.</span><span class=\"n\">group_send</span><span class=\"p\">(</span>\n
  \           <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">room_group_name</span><span
  class=\"p\">,</span> \n            <span class=\"p\">{</span>\n                <span
  class=\"s2\">&quot;type&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;chat_message&quot;</span><span
  class=\"p\">,</span>\n                <span class=\"s2\">&quot;message&quot;</span><span
  class=\"p\">:</span> <span class=\"n\">message</span><span class=\"p\">,</span>\n
  \               <span class=\"s2\">&quot;username&quot;</span><span class=\"p\">:</span>
  <span class=\"n\">username</span><span class=\"p\">,</span>\n            <span class=\"p\">}</span>\n
  \       <span class=\"p\">)</span>\n\n    <span class=\"k\">async</span> <span class=\"k\">def</span>
  <span class=\"nf\">chat_message</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"p\">,</span> <span class=\"n\">event</span><span class=\"p\">):</span>\n
  \       <span class=\"n\">message</span> <span class=\"o\">=</span> <span class=\"n\">event</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;message&quot;</span><span class=\"p\">]</span>\n
  \       <span class=\"n\">username</span> <span class=\"o\">=</span> <span class=\"n\">event</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;username&quot;</span><span class=\"p\">]</span>\n\n
  \       <span class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">send</span><span class=\"p\">(</span>\n            <span class=\"n\">text_data</span><span
  class=\"o\">=</span><span class=\"n\">json</span><span class=\"o\">.</span><span
  class=\"n\">dumps</span><span class=\"p\">(</span>\n                <span class=\"p\">{</span>\n
  \                   <span class=\"s2\">&quot;message&quot;</span><span class=\"p\">:</span>
  <span class=\"n\">message</span><span class=\"p\">,</span>\n                    <span
  class=\"s2\">&quot;username&quot;</span><span class=\"p\">:</span> <span class=\"n\">username</span>\n
  \               <span class=\"p\">}</span>\n            <span class=\"p\">)</span>\n
  \       <span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<h4>Accept the WebSocket
  connection</h4>\n<p>Here, room and group more or less mean the same thing but are
  different in different contexts. For instance, the group refers to the collection
  of clients which are connected to a channel(communication between consumer and web
  socket) and the Room is referring to the collection of clients connected to a single
  connection stream like a literal room. So we can say, the group is a technical term
  and the room is more of a layman's term for the same thing.</p>\n<p>The method <code>connect</code>
  is called when the client establishes a websocket connection. When that happens,
  the function gets the room slug from the URL of the client and stores <code>room_name</code>
  which is a string. It creates a separate variable called <code>room_group_name</code>
  by prepending the <code>chat_</code> string to the <code>room_name</code>, it also
  gets the currently logged-in user from the request. It then adds the <code>channel_name</code>
  to the group named <code>room_group_name</code>. The <code>channel_name</code> is
  a unique identifier to the connection/consumer in the channel. By adding the <code>channel_name</code>,
  the consumer then can broadcast the message to all the channels within the group.
  Finally, the function accepts the connection, and a <strong>webcoket connection
  is established from both ends, connection is sent from the client and is now accepted
  from the backend.</strong></p>\n<h4>Disconnect from the WebSocket connection</h4>\n<p>When
  the client sends a close connection request to the server, the <code>disconnect</code>
  method is triggered and it basically removes the <code>channel_name</code> from
  the group i.e. the group name <code>room_group_name</code> whatever the string it
  happens to be. So, it basically removes the client from the broadcast group and
  hence it can't receive or send messages through the websocket since it has been
  closed from both ends.</p>\n<p>If you would have paid attention to the <code>close_code</code>
  in-method parameter, it is not being used currently. However, we can use it for
  checking why the connection was closed, as the <code>close_code</code> is a numeric
  value just like the status code in the web request for letting the server know the
  reason for disconnecting from the client.</p>\n<h4>Receive a message from the WebSocket
  connection</h4>\n<p>The <code>recieve</code> method is the core of the consumer
  as it is responsible for all the logic and parsing of the message and broadcasting
  of the messages from the clients to the group channels. The function takes in a
  parameter called <code>text_data</code> and it is sent from the client through websocket,
  so it is JSON content. We need to get the actual message from the JSON object or
  any other piece of content from the client. So, we deserialize (convert the JSON
  object to python objects) the received payload, and get the value from the key <code>message</code>.
  The key is the input name or id from the client sending the request through the
  web socket, so it can be different depending on the frontend template(we'll see
  the front end soon as well).</p>\n<p>We get the user from the scope of the consumer
  as we previously initialized it in the connect method. This can be used for understanding
  which user has sent the message, it will be used later on as we send/broadcast the
  message to the group.</p>\n<p>Now, the final piece in the receive method is the
  <code>channel_layer.group_send</code> method, this method as the name suggests is
  used to send or broadcast the received message to the entire group. The method has
  two parameters:</p>\n<ol>\n<li>\n<p>The name of the group</p>\n</li>\n<li>\n<p>The
  JSON body containing the message and other details</p>\n</li>\n</ol>\n<p>The method
  is not directly sending the message but it has a type key in the JSON body which
  will be the function name to call. The function will simply call the other funciton
  mentioned in the type key in the dict. The following keys in the dict will be the
  parameters of that funciton. In this case, the funciton specified in the <code>type</code>
  key is <code>chat_message</code> which takes in the <code>event</code> as the parameter.
  This event will have all the parameters from the <code>group_send</code> method.</p>\n<p>So,
  the <code>chat_message</code> will load in this message, username, and the room
  name and then call the <code>send</code> method which actually sends the message
  as a JSON payload to the WebSocket connection which will be received by all the
  clients in the same group as provided in the <code>room_group_name</code> string.</p>\n<h3>Adding
  Routers for WebSocket connections</h3>\n<p>So, till this point have consumers, which
  are just like views in terms of channels. Now, we need some URL routes to map these
  consumers to a path. So, we will create a file/module called <code>routing.py</code>
  which will look quite similar to the <code>urls.py</code> file. It will have a list
  called <code>websocket_urlpatterns</code> just like <code>urlpatterns</code> with
  the list of <code>path</code>. These paths however are not <code>http</code> routes
  but will serve for the <code>WebSocket</code> path.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># chat / routing.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.urls</span> <span class=\"kn\">import</span> <span class=\"n\">path</span>\n\n<span
  class=\"kn\">from</span> <span class=\"nn\">chat</span> <span class=\"kn\">import</span>
  <span class=\"n\">consumers</span>\n\n<span class=\"n\">websocket_urlpatterns</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;chat/&lt;str:room_slug&gt;/&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">consumers</span><span class=\"o\">.</span><span
  class=\"n\">ChatConsumer</span><span class=\"o\">.</span><span class=\"n\">as_asgi</span><span
  class=\"p\">()),</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above code block, we have defined a URL for the web socket with the path as
  <code>/chat/&lt;room_slug&gt;</code> where room_name will be the <code>slug</code>
  for the room. The path is bound with the consumer-defined in the <code>consumers.py</code>
  module as <code>ChatConsumer</code>. The <code>as_asgi</code> method is used for
  converting a view into an ASGI-compatible view for the WebSocket interface.</p>\n<h3>Setting
  up ASGI Application</h3>\n<p>We are using the ASGI application config rather than
  a typical WSGI application which only works one request at a time. We want the chat
  application to be asynchronous because multiple clients might send and receive the
  messages at the same time, we don't want to make a client wait before the server
  process a message from another client, that's just the reason we are using WebSocket
  protocol.</p>\n<p>So, we need to also make sure, it makes the http request and also
  add our websocket config from the chat app we created in the previous sections.
  So, inside the <code>asgi.py</code> file in the project config module, we need to
  make some changes to include the chat application configurations.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># backchat / asgi.py</span>\n\n\n<span class=\"kn\">import</span> <span
  class=\"nn\">os</span>\n<span class=\"kn\">from</span> <span class=\"nn\">django.core.asgi</span>
  <span class=\"kn\">import</span> <span class=\"n\">get_asgi_application</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">channels.auth</span> <span class=\"kn\">import</span>
  <span class=\"n\">AuthMiddlewareStack</span>\n<span class=\"kn\">from</span> <span
  class=\"nn\">channels.routing</span> <span class=\"kn\">import</span> <span class=\"n\">ProtocolTypeRouter</span><span
  class=\"p\">,</span> <span class=\"n\">URLRouter</span>\n\n<span class=\"n\">os</span><span
  class=\"o\">.</span><span class=\"n\">environ</span><span class=\"o\">.</span><span
  class=\"n\">setdefault</span><span class=\"p\">(</span><span class=\"s1\">&#39;DJANGO_SETTINGS_MODULE&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;backchat.settings&#39;</span><span
  class=\"p\">)</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">chat</span>
  <span class=\"kn\">import</span> <span class=\"n\">routing</span>\n\n\n<span class=\"n\">application</span>
  <span class=\"o\">=</span> <span class=\"n\">ProtocolTypeRouter</span><span class=\"p\">({</span>\n
  \   <span class=\"s1\">&#39;http&#39;</span><span class=\"p\">:</span> <span class=\"n\">get_asgi_application</span><span
  class=\"p\">(),</span>\n    <span class=\"s2\">&quot;websocket&quot;</span><span
  class=\"p\">:</span><span class=\"n\">AuthMiddlewareStack</span><span class=\"p\">(</span>\n
  \       <span class=\"n\">URLRouter</span><span class=\"p\">(</span>\n            <span
  class=\"n\">routing</span><span class=\"o\">.</span><span class=\"n\">websocket_urlpatterns</span>\n
  \       <span class=\"p\">)</span>\n    <span class=\"p\">)</span>\n<span class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p>We
  will override the <code>application</code> config which is a component used for
  routing different types of protocols for the <code>ASGI</code> application. We have
  set the two keys, <code>http</code> and <code>websocket</code> in our application.
  The <code>http</code> type of requests will be served with the <code>get_asgi_application</code>
  application which is used for running the application in an ASGI environment.</p>\n<p>For
  <code>websocket</code> type of requests, we are setting the <a href=\"https://channels.readthedocs.io/en/latest/topics/authentication.html\">AuthMiddlewareStack</a>
  which helps in authenticating the users requesting the WebSocket connection and
  allows only authorized users to make a connection to the application. The <a href=\"https://channels.readthedocs.io/en/stable/topics/routing.html\">URLRouter</a>
  is used for mapping the list of URL patterns with the incoming request. So, this
  basically serves the request URL with the appropriate consumer in the application.
  We are parsing in the <code>websocket_urlpatterns</code> as the list of URLs that
  will be used for the WebSocket connections.</p>\n<p>Now, we run the server, we should
  be seeing the <code>ASGI</code> server serving our application rather than the plain
  WSGI application.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ python manage.py runserver\n\nWatching
  for file changes with StatReloader\nPerforming system checks...\n\nSystem check
  identified no issues (0 silenced).\nFebruary 05, 2023 - 09:22:45\nDjango version
  4.1.5, using settings &#39;backchat.settings&#39;\nStarting ASGI/Daphne version
  4.0.0 development server at http://127.0.0.1:8000/\nQuit the server with CONTROL-C.\n</pre></div>\n\n</pre>\n\n<p>The
  application is not complete yet, it might not have most components working functional
  yet. So, we'll now get into making the user interfaces for the application, to create,
  join, and view rooms in the application.</p>\n<h3>Adding Views for Chat Rooms</h3>\n<p>We
  will have a couple of views like create room page, the join room page, and the chat
  room page. We will define each view as a distinct view and all of them will require
  authenticated users.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># chat / views.py</span>\n\n\n<span class=\"kn\">import</span> <span
  class=\"nn\">string</span>\n<span class=\"kn\">import</span> <span class=\"nn\">random</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.contrib.auth.decorators</span>
  <span class=\"kn\">import</span> <span class=\"n\">login_required</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.shortcuts</span> <span class=\"kn\">import</span>
  <span class=\"n\">render</span><span class=\"p\">,</span> <span class=\"n\">reverse</span><span
  class=\"p\">,</span> <span class=\"n\">redirect</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">django.utils.text</span> <span class=\"kn\">import</span> <span
  class=\"n\">slugify</span>\n<span class=\"kn\">from</span> <span class=\"nn\">chat.models</span>
  <span class=\"kn\">import</span> <span class=\"n\">Room</span>\n\n\n<span class=\"nd\">@login_required</span>\n<span
  class=\"k\">def</span> <span class=\"nf\">index</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"n\">slug</span><span
  class=\"p\">):</span>\n    <span class=\"n\">room</span> <span class=\"o\">=</span>
  <span class=\"n\">Room</span><span class=\"o\">.</span><span class=\"n\">objects</span><span
  class=\"o\">.</span><span class=\"n\">get</span><span class=\"p\">(</span><span
  class=\"n\">slug</span><span class=\"o\">=</span><span class=\"n\">slug</span><span
  class=\"p\">)</span>\n    <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;chat/room.html&#39;</span><span class=\"p\">,</span> <span class=\"p\">{</span><span
  class=\"s1\">&#39;name&#39;</span><span class=\"p\">:</span> <span class=\"n\">room</span><span
  class=\"o\">.</span><span class=\"n\">name</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;slug&#39;</span><span class=\"p\">:</span> <span class=\"n\">room</span><span
  class=\"o\">.</span><span class=\"n\">slug</span><span class=\"p\">})</span>\n\n<span
  class=\"nd\">@login_required</span>\n<span class=\"k\">def</span> <span class=\"nf\">room_create</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"k\">if</span> <span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">method</span> <span class=\"o\">==</span> <span class=\"s2\">&quot;POST&quot;</span><span
  class=\"p\">:</span>\n        <span class=\"n\">room_name</span> <span class=\"o\">=</span>
  <span class=\"n\">request</span><span class=\"o\">.</span><span class=\"n\">POST</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;room_name&quot;</span><span class=\"p\">]</span>\n
  \       <span class=\"n\">uid</span> <span class=\"o\">=</span> <span class=\"nb\">str</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;&#39;</span><span class=\"o\">.</span><span
  class=\"n\">join</span><span class=\"p\">(</span><span class=\"n\">random</span><span
  class=\"o\">.</span><span class=\"n\">choices</span><span class=\"p\">(</span><span
  class=\"n\">string</span><span class=\"o\">.</span><span class=\"n\">ascii_letters</span>
  <span class=\"o\">+</span> <span class=\"n\">string</span><span class=\"o\">.</span><span
  class=\"n\">digits</span><span class=\"p\">,</span> <span class=\"n\">k</span><span
  class=\"o\">=</span><span class=\"mi\">4</span><span class=\"p\">)))</span>\n        <span
  class=\"n\">room_slug</span> <span class=\"o\">=</span> <span class=\"n\">slugify</span><span
  class=\"p\">(</span><span class=\"n\">room_name</span> <span class=\"o\">+</span>
  <span class=\"s2\">&quot;_&quot;</span> <span class=\"o\">+</span> <span class=\"n\">uid</span><span
  class=\"p\">)</span>\n        <span class=\"n\">room</span> <span class=\"o\">=</span>
  <span class=\"n\">Room</span><span class=\"o\">.</span><span class=\"n\">objects</span><span
  class=\"o\">.</span><span class=\"n\">create</span><span class=\"p\">(</span><span
  class=\"n\">name</span><span class=\"o\">=</span><span class=\"n\">room_name</span><span
  class=\"p\">,</span> <span class=\"n\">slug</span><span class=\"o\">=</span><span
  class=\"n\">room_slug</span><span class=\"p\">)</span>\n        <span class=\"k\">return</span>
  <span class=\"n\">redirect</span><span class=\"p\">(</span><span class=\"n\">reverse</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;chat&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">kwargs</span><span class=\"o\">=</span><span class=\"p\">{</span><span
  class=\"s1\">&#39;slug&#39;</span><span class=\"p\">:</span> <span class=\"n\">room</span><span
  class=\"o\">.</span><span class=\"n\">slug</span><span class=\"p\">}))</span>\n
  \   <span class=\"k\">else</span><span class=\"p\">:</span>\n        <span class=\"k\">return</span>
  <span class=\"n\">render</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;chat/create.html&#39;</span><span class=\"p\">)</span>\n\n<span
  class=\"nd\">@login_required</span>\n<span class=\"k\">def</span> <span class=\"nf\">room_join</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"k\">if</span> <span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">method</span> <span class=\"o\">==</span> <span class=\"s2\">&quot;POST&quot;</span><span
  class=\"p\">:</span>\n        <span class=\"n\">room_slug</span> <span class=\"o\">=</span>
  <span class=\"n\">request</span><span class=\"o\">.</span><span class=\"n\">POST</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;room_slug&quot;</span><span class=\"p\">]</span>\n
  \       <span class=\"n\">room</span> <span class=\"o\">=</span> <span class=\"n\">Room</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">get</span><span class=\"p\">(</span><span class=\"n\">slug</span><span
  class=\"o\">=</span><span class=\"n\">room_slug</span><span class=\"p\">)</span>\n
  \       <span class=\"k\">return</span> <span class=\"n\">redirect</span><span class=\"p\">(</span><span
  class=\"n\">reverse</span><span class=\"p\">(</span><span class=\"s1\">&#39;chat&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">kwargs</span><span class=\"o\">=</span><span
  class=\"p\">{</span><span class=\"s1\">&#39;slug&#39;</span><span class=\"p\">:</span>
  <span class=\"n\">room</span><span class=\"o\">.</span><span class=\"n\">slug</span><span
  class=\"p\">}))</span>\n    <span class=\"k\">else</span><span class=\"p\">:</span>\n
  \       <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;chat/join.html&#39;</span><span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>In the above views module, we
  have added 3 views namely <code>index</code> as the room page, <code>room_create</code>
  for the room creation page, and the <code>room_join</code> for the room join page.
  The index view is a simple get request to the provided slug of the room, it gets
  the slug from the URL from the request and fetches an object of the room associated
  with that slug. Then it renders the room template with the context variables like
  the name of the room and the slug associated with that room.</p>\n<p>The <code>room_create</code>
  view is a simple two-case view that either can render the room creation page or
  process the submitted form and create the room. Just like we used in the <code>register</code>
  view in the accounts app. When the user will send a <code>GET</code> request to
  the URL which we will map to <code>/create/</code> shortly after this, the user
  will be given a form. So, we will render the <code>create.html</code> template.
  We will create the html template shortly.\nIf the user has sent a <code>POST</code>
  request to the view via the <code>/create</code> URL, we will fetch the name field
  in the sent request and create a unique identifier with the name of the room. We
  will slugify the concatenation of the name with the uid and will set it as the slug
  of the room. We will then simply create the room and redirect the user to the <code>room</code>
  page.</p>\n<p>The <code>room_join</code> view also is a two-case view, where the
  user can either request the join room form or send a slug with the form submission.
  If the user is requesting a form, we will render the <code>join.html</code> template.
  If the user is submitting the form, we will fetch the room based on the slug provided
  and redirect the user to the <code>room</code> page.</p>\n<p>So, the <code>room_join</code>
  and <code>room_create</code> views are quite similar, we are fetching an already
  existing room in the case of the join view and creating a new instance of room in
  the create view. Now, we will connect the views to the URLs and finally get to the
  templates.</p>\n<h3>Connecting Views and URLs</h3>\n<p>We have three views to route
  to a URL. But we will also have one additional URL which will be the home page for
  the application, on that page the user can choose to either join or create a room.
  We have the room creation, join the room and the room view to be mapped in this
  URL routing of the app.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># chat / urls.py</span>\n\n\n<span class=\"kn\">from</span> <span class=\"nn\">django.urls</span>
  <span class=\"kn\">import</span> <span class=\"n\">path</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">django.views.generic</span> <span class=\"kn\">import</span>
  <span class=\"n\">TemplateView</span>\n<span class=\"kn\">from</span> <span class=\"nn\">chat</span>
  <span class=\"kn\">import</span> <span class=\"n\">views</span>\n\n\n<span class=\"n\">urlpatterns</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;&quot;</span><span class=\"p\">,</span>
  <span class=\"n\">TemplateView</span><span class=\"o\">.</span><span class=\"n\">as_view</span><span
  class=\"p\">(</span><span class=\"n\">template_name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;base.html&quot;</span><span class=\"p\">),</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;index&#39;</span><span class=\"p\">),</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;room/&lt;str:slug&gt;/&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">index</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;chat&#39;</span><span class=\"p\">),</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;create/&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">room_create</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;room-create&#39;</span><span class=\"p\">),</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;join/&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">room_join</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;room-join&#39;</span><span class=\"p\">),</span>\n<span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>So, the first route is the home
  page view called <code>index</code>, we have used the <a href=\"https://docs.djangoproject.com/en/4.1/ref/class-based-views/base/#templateview\">TemplateView</a>
  which will simply render the template provided. We don't have to create a separate
  view for just rendering a template. We already have defined the <code>base.html</code>
  template while setting up the <code>accounts</code> app. This will be the same template,
  we will add some more content to the template later on. The URL mapped is the <code>/</code>
  since we will add the URLs of this app to the project URLs with an empty <code>&quot;&quot;</code>
  path.</p>\n<p>The second route is used for the room index page, i.e. where the user
  will be able to send and receive messages. The path is assigned as <code>/room/&lt;str:slug&gt;/</code>
  indicating a parameter called slug of type string will be used in accessing a particular
  room. The URL will be bound to the index view that we created in the views file,
  which fetches the room based on the slug, so here is where the slug will be coming
  from. The name of the URL-View route will be <code>chat</code> but you can keep
  it as per your requirements. The URL name is really handy for use in the templates.</p>\n<p>The
  third route is for the room creation page. The <code>/create/</code> URL will be
  bound to the <code>room_create</code> view, as we discussed, it will serve two purposes,
  one to render the form for creating the room, and the other for sending a <code>POST</code>
  request to the same path for the creating the Room with the name provided. The name
  is not required but helps in identifying and making it user-friendly. The name of
  this URL is set as <code>room-create</code>.</p>\n<p>The final route is for joining
  the room, the <code>/join/</code> URL will be triggering the <code>room_join</code>
  view. Similar to the <code>room-create</code> URL, the URL will render the join
  room form on a <code>GET</code> request, fetch the room with the provided slug and
  redirect to the room page. Here, the slug field in the form will be required for
  actually finding the matching room. The name of the URL route is set as <code>room-join</code>.</p>\n<p>We
  now add the URLs from the chat app to the project URLs. This will make the <code>/</code>
  as the entry point for the chat application URLs.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># backchat / urls.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.contrib</span> <span class=\"kn\">import</span> <span class=\"n\">admin</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.urls</span> <span class=\"kn\">import</span>
  <span class=\"n\">path</span><span class=\"p\">,</span> <span class=\"n\">include</span>\n\n<span
  class=\"n\">urlpatterns</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;admin/&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">admin</span><span class=\"o\">.</span><span
  class=\"n\">site</span><span class=\"o\">.</span><span class=\"n\">urls</span><span
  class=\"p\">),</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;auth/&quot;</span><span class=\"p\">,</span> <span class=\"n\">include</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;accounts.urls&#39;</span><span class=\"p\">)),</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">include</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;chat.urls&#39;</span><span class=\"p\">)),</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>Hence
  the process is completed for the backend to process the message, it then is dependent
  on the client to process and render the message.</p>\n<p>Till HTMX was not a thing!</p>\n<p>We
  won't have to write a single line of javascript to receive and handle the WebSocket
  connection!</p>\n<h3>Creating Templates and adding htmx</h3>\n<p>We now move into
  the actual frontend or creating the template for actually working with the rooms
  and user interaction. We will have three pieces of templates, a room creates the
  page, a room join page, and a room chat page. As these template names suggest, they
  will be used for creating a room with the name, joining the room with the slug,
  and the room chat page where the user will send and receive messages.</p>\n<p>Let/s
  modify the base template first.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cp\">&lt;!DOCTYPE html&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">meta</span>
  <span class=\"na\">charset</span><span class=\"o\">=</span><span class=\"s\">&quot;utf-8&quot;</span>
  <span class=\"p\">/&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>Chat App<span class=\"p\">&lt;/</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>\n        {% load static %}\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">script</span> <span class=\"na\">src</span><span class=\"o\">=</span><span
  class=\"s\">&quot;https://unpkg.com/htmx.org@1.8.5&quot;</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">script</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span
  class=\"nt\">head</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% url &#39;index&#39; %}&quot;</span><span class=\"p\">&gt;</span>Home<span
  class=\"p\">&lt;/</span><span class=\"nt\">a</span><span class=\"p\">&gt;</span>\n
  \       {% if user.is_authenticated %}\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% url &#39;logout&#39; %}&quot;</span><span class=\"p\">&gt;</span>Logout<span
  class=\"p\">&lt;/</span><span class=\"nt\">a</span><span class=\"p\">&gt;</span>\n
  \       {% else %}\n            <span class=\"p\">&lt;</span><span class=\"nt\">a</span>
  <span class=\"na\">href</span><span class=\"o\">=</span><span class=\"s\">&quot;{%
  url &#39;login&#39; %}&quot;</span><span class=\"p\">&gt;</span>Login<span class=\"p\">&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n        {% endif %}\n    <span
  class=\"p\">&lt;</span><span class=\"nt\">body</span> <span class=\"na\">hx-ext</span><span
  class=\"o\">=</span><span class=\"s\">&quot;ws&quot;</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;</span><span class=\"nt\">h1</span><span class=\"p\">&gt;</span>Back
  Chat<span class=\"p\">&lt;/</span><span class=\"nt\">h1</span><span class=\"p\">&gt;</span>\n
  \       {% block base %}\n            <span class=\"p\">&lt;</span><span class=\"nt\">a</span>
  <span class=\"na\">href</span><span class=\"o\">=</span><span class=\"s\">&quot;{%
  url &#39;room-join&#39; %}&quot;</span><span class=\"p\">&gt;</span>Join Room<span
  class=\"p\">&lt;/</span><span class=\"nt\">a</span><span class=\"p\">&gt;</span>\n
  \           <span class=\"p\">&lt;</span><span class=\"nt\">a</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% url &#39;room-create&#39; %}&quot;</span><span
  class=\"p\">&gt;</span>Create Room<span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n        {% endblock %}\n    <span class=\"p\">&lt;/</span><span
  class=\"nt\">body</span><span class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span
  class=\"nt\">html</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-home-page.png\" alt=\"Chat
  App Home Page\" /></p>\n<h4>Create Room Template</h4>\n<p>We will have to render
  the form with a field like <code>name</code> for setting it as the name of the room,
  it is not required but again, it makes it easier for the user to find the name of
  the room a bit more friendly than random characters.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>#
  templates / chat / create.html\n\n\n{% extends &#39;base.html&#39; %}\n\n{% block
  base %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">form</span> <span class=\"na\">method</span><span
  class=\"o\">=</span><span class=\"s\">&#39;post&#39;</span> <span class=\"na\">action</span><span
  class=\"o\">=</span><span class=\"s\">&#39;{% url &#39;</span><span class=\"na\">room-create</span><span
  class=\"err\">&#39;</span> <span class=\"err\">%}&#39;</span><span class=\"p\">&gt;</span>\n
  \       {% csrf_token %}\n        <span class=\"p\">&lt;</span><span class=\"nt\">input</span>
  <span class=\"na\">name</span><span class=\"o\">=</span><span class=\"s\">&#39;room_name&#39;</span>
  <span class=\"na\">id</span><span class=\"o\">=</span><span class=\"s\">&#39;room_name&#39;</span>
  <span class=\"na\">placeholder</span><span class=\"o\">=</span><span class=\"s\">&#39;Room
  Name&#39;</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">input</span> <span class=\"na\">type</span><span class=\"o\">=</span><span
  class=\"s\">&#39;submit&#39;</span> <span class=\"na\">id</span><span class=\"o\">=</span><span
  class=\"s\">&quot;submit&quot;</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span
  class=\"nt\">form</span><span class=\"p\">&gt;</span>\n{% endblock %}\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-create-room-page.png\" alt=\"Chat
  Room Create Page\" /></p>\n<p>The template is inherited from the <code>base.html</code>
  template and we render a form with the <code>room_name</code> input. The form will
  be submitted to the URL named <code>room-create</code> hence it will send a <code>POST</code>
  request to the server where it will create the room and further process the request.</p>\n<h4>Join
  Room Template</h4>\n<p>The join room template is similar to the create room template
  except it gets the slug of the room from the user rather than the name which is
  not a unique one to join the room.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>#
  templates / chat / join.html\n\n\n{% extends &#39;base.html&#39; %}\n\n{% block
  base %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">form</span> <span class=\"na\">method</span><span
  class=\"o\">=</span><span class=\"s\">&#39;post&#39;</span> <span class=\"na\">action</span><span
  class=\"o\">=</span><span class=\"s\">&#39;{% url &#39;</span><span class=\"na\">room-join</span><span
  class=\"err\">&#39;</span> <span class=\"err\">%}&#39;</span><span class=\"p\">&gt;</span>\n
  \       {% csrf_token %}\n        <span class=\"p\">&lt;</span><span class=\"nt\">input</span>
  <span class=\"na\">name</span><span class=\"o\">=</span><span class=\"s\">&#39;room_slug&#39;</span>
  <span class=\"na\">id</span><span class=\"o\">=</span><span class=\"s\">&#39;room_slug&#39;</span>
  <span class=\"na\">required</span><span class=\"o\">=</span><span class=\"s\">&#39;true&#39;</span>
  <span class=\"na\">placeholder</span><span class=\"o\">=</span><span class=\"s\">&#39;Room
  Code&#39;</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">input</span> <span class=\"na\">type</span><span class=\"o\">=</span><span
  class=\"s\">&#39;submit&#39;</span> <span class=\"na\">id</span><span class=\"o\">=</span><span
  class=\"s\">&quot;submit&quot;</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span
  class=\"nt\">form</span><span class=\"p\">&gt;</span>\n{% endblock %}\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-join-room-page.png\" alt=\"Chat
  Room Join Page\" /></p>\n<p>The form will be submitted to the URL named <code>room-join</code>
  hence it will send a <code>POST</code> request to the server where it will fetch
  the room and further process the request.</p>\n<h3>Room Template (HTMX code)</h3>\n<p>Now,
  time for the actual ingredient in the application, some HTMX magic!</p>\n<p>This
  template, as the two templates above inherit from the base template, that's the
  same thing. But it has a special <code>div</code> with the attribute <a href=\"https://htmx.org/attributes/hx-ws/\">hx-ws</a>
  which is used for using attributes related to the web socket in the htmx library.
  The <code>connect</code> value is used for connecting to a WebSocket. The value
  of the attribute must be set to the URL of the WebSocket to be connected. In our
  case, it is the URL path from the <code>routing</code> app as <code>/chat/&lt;room_slug&gt;/</code>.
  This simply will connect the client to the WebSocket from the backend. The other
  important attribute is the <code>send</code> which is used for sending a message
  to the connected web socket.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>#
  templates / chat / room.html\n\n\n{% extends &#39;base.html&#39; %}\n\n{% block
  base %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>{{
  name }}<span class=\"p\">&lt;/</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;</span><span class=\"nt\">div</span> <span class=\"na\">hx-ws</span><span
  class=\"o\">=</span><span class=\"s\">&quot;connect:/chat/{{ slug }}/&quot;</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">form</span>
  <span class=\"na\">hx-ws</span><span class=\"o\">=</span><span class=\"s\">&quot;send:submit&quot;</span><span
  class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span class=\"nt\">input</span>
  <span class=\"na\">name</span><span class=\"o\">=</span><span class=\"s\">&quot;message&quot;</span><span
  class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span class=\"nt\">input</span>
  <span class=\"na\">type</span><span class=\"o\">=</span><span class=\"s\">&quot;submit&quot;</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;/</span><span class=\"nt\">form</span><span
  class=\"p\">&gt;</span>\n     <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n     <span class=\"p\">&lt;</span><span class=\"nt\">div</span>
  <span class=\"na\">id</span><span class=\"o\">=</span><span class=\"s\">&#39;messages&#39;</span><span
  class=\"p\">&gt;&lt;/</span><span class=\"nt\">div</span><span class=\"p\">&gt;</span>\n{%
  endblock %}\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-room-page.png\"
  alt=\"Chat Room Page\" /></p>\n<p>NOTE: The template has a div with the id <code>messages</code>
  which will be very important for sending the messages from the WebSocket to the
  client, so more on that when we use the HTMX part.</p>\n<p>For testing this template,
  you can create a room, and that will redirect you to the room template as we have
  seen in the views for the room creation. If you see something like <code>WebSocket
  CONNECT</code> it means, that the application has been able to establish a WebSocket
  connection to the backend, and we can be ready to accept messages and other stuff.</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>HTTP GET /chat/room/def_teas/
  200 [0.03, 127.0.0.1:38660]\nWebSocket HANDSHAKING /chat/def_teas/ [127.0.0.1:38666]\nWebSocket
  CONNECT /chat/def_teas/ [127.0.0.1:38666]\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-asgi-server.png\"
  alt=\"Django ASGI server websocket connection\" /></p>\n<p>Till this point, we should
  have a running and almost complete application, though we just have a minor part
  missing that will be the most important part.</p>\n<h3>Sending HTML response from
  backend for htmx</h3>\n<p>We will be sending a fragment of HTML from the backend
  when the user sends a message, to broadcast it to the group. Let's make some changes
  to the application, especially to the receive method in the <code>ChatConsumer</code>
  of the chat application.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># chat / consumers.py</span>\n    \n\n    <span class=\"o\">...</span>\n
  \   <span class=\"o\">...</span>\n\n    <span class=\"k\">async</span> <span class=\"k\">def</span>
  <span class=\"nf\">receive</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"p\">,</span> <span class=\"n\">text_data</span><span class=\"p\">):</span>\n
  \       <span class=\"n\">text_data_json</span> <span class=\"o\">=</span> <span
  class=\"n\">json</span><span class=\"o\">.</span><span class=\"n\">loads</span><span
  class=\"p\">(</span><span class=\"n\">text_data</span><span class=\"p\">)</span>\n
  \       <span class=\"n\">message</span> <span class=\"o\">=</span> <span class=\"n\">text_data_json</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;message&quot;</span><span class=\"p\">]</span>\n
  \       <span class=\"n\">user</span> <span class=\"o\">=</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">user</span>\n        <span class=\"n\">username</span>
  <span class=\"o\">=</span> <span class=\"n\">user</span><span class=\"o\">.</span><span
  class=\"n\">username</span>\n\n        <span class=\"k\">await</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">channel_layer</span><span class=\"o\">.</span><span
  class=\"n\">group_send</span><span class=\"p\">(</span>\n            <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">room_group_name</span><span class=\"p\">,</span>
  \n            <span class=\"p\">{</span>\n                <span class=\"s2\">&quot;type&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;chat_message&quot;</span><span class=\"p\">,</span>\n
  \               <span class=\"s2\">&quot;message&quot;</span><span class=\"p\">:</span>
  <span class=\"n\">message</span><span class=\"p\">,</span>\n                <span
  class=\"s2\">&quot;username&quot;</span><span class=\"p\">:</span> <span class=\"n\">username</span><span
  class=\"p\">,</span>\n            <span class=\"p\">}</span>\n        <span class=\"p\">)</span>\n\n
  \   <span class=\"k\">async</span> <span class=\"k\">def</span> <span class=\"nf\">chat_message</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"p\">,</span> <span
  class=\"n\">event</span><span class=\"p\">):</span>\n        <span class=\"n\">message</span>
  <span class=\"o\">=</span> <span class=\"n\">event</span><span class=\"p\">[</span><span
  class=\"s2\">&quot;message&quot;</span><span class=\"p\">]</span>\n        <span
  class=\"n\">username</span> <span class=\"o\">=</span> <span class=\"n\">event</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;username&quot;</span><span class=\"p\">]</span>\n\n
  \       <span class=\"c1\"># This is the crucial part of the application</span>\n
  \       <span class=\"n\">message_html</span> <span class=\"o\">=</span> <span class=\"sa\">f</span><span
  class=\"s2\">&quot;&lt;div hx-swap-oob=&#39;beforeend:#messages&#39;&gt;&lt;p&gt;&lt;b&gt;</span><span
  class=\"si\">{</span><span class=\"n\">username</span><span class=\"si\">}</span><span
  class=\"s2\">&lt;/b&gt;: </span><span class=\"si\">{</span><span class=\"n\">message</span><span
  class=\"si\">}</span><span class=\"s2\">&lt;/p&gt;&lt;/div&gt;&quot;</span>\n        <span
  class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">send</span><span class=\"p\">(</span>\n            <span class=\"n\">text_data</span><span
  class=\"o\">=</span><span class=\"n\">json</span><span class=\"o\">.</span><span
  class=\"n\">dumps</span><span class=\"p\">(</span>\n                <span class=\"p\">{</span>\n
  \                   <span class=\"s2\">&quot;message&quot;</span><span class=\"p\">:</span>
  <span class=\"n\">message_html</span><span class=\"p\">,</span>\n                    <span
  class=\"s2\">&quot;username&quot;</span><span class=\"p\">:</span> <span class=\"n\">username</span>\n
  \               <span class=\"p\">}</span>\n            <span class=\"p\">)</span>\n
  \       <span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-room-msg-1.png\"
  alt=\"Chat Room Message\" />\n<img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-room-msg-2.png\"
  alt=\"Chat Room Message 2 Users\" />\n<img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-room-msg-3.png\"
  alt=\"Chat Room Message\" /></p>\n<p>In the above snippet, we are just changing
  the final message object to include some HTML just simple. The HTML however has
  home htmx attributes like <a href=\"https://htmx.org/attributes/hx-swap-oob/\">hx-swap-oob</a>
  which will just update the specified DOM element to the content in the div. In this
  case, the DOM element is <code>#message</code> which is the id message present in
  the room template. We basically add the username and the message into the same id
  by appending it to the element. That's it, it would work and it would start showing
  the messages from the connected clients and broadcast them as well.</p>\n<p>There
  are some things to keep in mind while using htmx in the long run especially when
  the htmx 2.0 is released, it will have <code>ws</code> as a separate extension.
  It will have a bit of a different syntax than above. I have tried the latest version
  but doesn't seem to work. I'll just leave a few snippets for your understanding
  of the problem.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>#
  templates / chat / room.html\n\n\n{% extends &#39;base.html&#39; %}\n\n{% block
  base %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>{{
  name }}<span class=\"p\">&lt;/</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;</span><span class=\"nt\">div</span> <span class=\"na\">hx-ext</span><span
  class=\"o\">=</span><span class=\"s\">&quot;ws&quot;</span> <span class=\"na\">ws-connect</span><span
  class=\"o\">=</span><span class=\"s\">&quot;/chat/{{ slug }}/&quot;</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">form</span>
  <span class=\"na\">ws-send</span><span class=\"p\">&gt;</span>\n            <span
  class=\"p\">&lt;</span><span class=\"nt\">input</span> <span class=\"na\">name</span><span
  class=\"o\">=</span><span class=\"s\">&quot;message&quot;</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;/</span><span class=\"nt\">form</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;</span><span class=\"nt\">div</span> <span class=\"na\">id</span><span
  class=\"o\">=</span><span class=\"s\">&#39;messages&#39;</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">div</span><span class=\"p\">&gt;</span>\n{% endblock %}\n</pre></div>\n\n</pre>\n\n<p>I
  have added, the <code>hx-ext</code> as <code>ws</code> which is a htmx <a href=\"https://htmx.org/extensions/web-sockets/\">extension
  for websocket</a>. This extension has websocket-specific attributes like <code>ws-connect</code>
  and <code>ws-send</code>. I have tried a combination like changing the htmx versions,
  adding submit value as the <code>ws-send</code> attribute, etc, but no results yet.
  I have opened a <a href=\"https://github.com/bigskysoftware/htmx/discussions/1231\">discussion</a>
  on GitHub for this issue, you can express your solution or views there.</p>\n<h3>Adding
  some utility features for the chat app</h3>\n<p>We can save messages, add and remove
  the users from the room according to the connection, and other stuff that can make
  this a fully-fledged app. So, I have made a few changes to the chat consumers for
  saving the messages and also updating the room with the users in the room.</p>\n<pre
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># chat / consumers.py</span>\n\n\n<span class=\"kn\">import</span>
  <span class=\"nn\">json</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">asgiref.sync</span>
  <span class=\"kn\">import</span> <span class=\"n\">sync_to_async</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">channels.generic.websocket</span> <span class=\"kn\">import</span>
  <span class=\"n\">AsyncWebsocketConsumer</span>\n\n<span class=\"kn\">from</span>
  <span class=\"nn\">chat.models</span> <span class=\"kn\">import</span> <span class=\"n\">Room</span><span
  class=\"p\">,</span> <span class=\"n\">Message</span>\n\n\n<span class=\"k\">class</span>
  <span class=\"nc\">ChatConsumer</span><span class=\"p\">(</span><span class=\"n\">AsyncWebsocketConsumer</span><span
  class=\"p\">):</span>\n    <span class=\"k\">async</span> <span class=\"k\">def</span>
  <span class=\"nf\">connect</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"p\">):</span>\n        <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_name</span> <span class=\"o\">=</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">scope</span><span class=\"p\">[</span><span
  class=\"s2\">&quot;url_route&quot;</span><span class=\"p\">][</span><span class=\"s2\">&quot;kwargs&quot;</span><span
  class=\"p\">][</span><span class=\"s2\">&quot;room_slug&quot;</span><span class=\"p\">]</span>\n
  \       <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">room_group_name</span>
  <span class=\"o\">=</span> <span class=\"s2\">&quot;chat_</span><span class=\"si\">%s</span><span
  class=\"s2\">&quot;</span> <span class=\"o\">%</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">room_name</span>\n        <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">user</span> <span class=\"o\">=</span> <span
  class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">scope</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;user&quot;</span><span class=\"p\">]</span>\n\n
  \       <span class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">channel_layer</span><span class=\"o\">.</span><span class=\"n\">group_add</span><span
  class=\"p\">(</span>\n            <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_group_name</span><span class=\"p\">,</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">channel_name</span>\n        <span class=\"p\">)</span>\n\n
  \       <span class=\"c1\"># Add the user when the client connects</span>\n        <span
  class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">add_user</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">room_name</span><span class=\"p\">,</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">user</span><span
  class=\"p\">)</span>\n\n        <span class=\"k\">await</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">accept</span><span class=\"p\">()</span>\n\n
  \   <span class=\"k\">async</span> <span class=\"k\">def</span> <span class=\"nf\">disconnect</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"p\">,</span> <span
  class=\"n\">close_code</span><span class=\"p\">):</span>\n\n        <span class=\"c1\">#
  Remove the user when the client disconnects</span>\n        <span class=\"k\">await</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">remove_user</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_name</span><span class=\"p\">,</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">user</span><span class=\"p\">)</span>\n\n
  \       <span class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">channel_layer</span><span class=\"o\">.</span><span class=\"n\">group_discard</span><span
  class=\"p\">(</span>\n            <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_group_name</span><span class=\"p\">,</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">channel_name</span>\n        <span class=\"p\">)</span>\n\n
  \   <span class=\"k\">async</span> <span class=\"k\">def</span> <span class=\"nf\">receive</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"p\">,</span> <span
  class=\"n\">text_data</span><span class=\"p\">):</span>\n        <span class=\"n\">text_data_json</span>
  <span class=\"o\">=</span> <span class=\"n\">json</span><span class=\"o\">.</span><span
  class=\"n\">loads</span><span class=\"p\">(</span><span class=\"n\">text_data</span><span
  class=\"p\">)</span>\n        <span class=\"n\">message</span> <span class=\"o\">=</span>
  <span class=\"n\">text_data_json</span><span class=\"p\">[</span><span class=\"s2\">&quot;message&quot;</span><span
  class=\"p\">]</span>\n        <span class=\"n\">user</span> <span class=\"o\">=</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">user</span>\n
  \       <span class=\"n\">username</span> <span class=\"o\">=</span> <span class=\"n\">user</span><span
  class=\"o\">.</span><span class=\"n\">username</span>\n        <span class=\"n\">room</span>
  <span class=\"o\">=</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_name</span>\n\n        <span class=\"c1\"># Save the message on
  recieving</span>\n        <span class=\"k\">await</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">save_message</span><span class=\"p\">(</span><span
  class=\"n\">room</span><span class=\"p\">,</span> <span class=\"n\">user</span><span
  class=\"p\">,</span> <span class=\"n\">message</span><span class=\"p\">)</span>\n\n
  \       <span class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">channel_layer</span><span class=\"o\">.</span><span class=\"n\">group_send</span><span
  class=\"p\">(</span>\n            <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_group_name</span><span class=\"p\">,</span> \n            <span
  class=\"p\">{</span>\n                <span class=\"s2\">&quot;type&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;chat_message&quot;</span><span class=\"p\">,</span>\n
  \               <span class=\"s2\">&quot;message&quot;</span><span class=\"p\">:</span>
  <span class=\"n\">message</span><span class=\"p\">,</span>\n                <span
  class=\"s2\">&quot;username&quot;</span><span class=\"p\">:</span> <span class=\"n\">username</span><span
  class=\"p\">,</span>\n            <span class=\"p\">}</span>\n        <span class=\"p\">)</span>\n\n
  \   <span class=\"k\">async</span> <span class=\"k\">def</span> <span class=\"nf\">chat_message</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"p\">,</span> <span
  class=\"n\">event</span><span class=\"p\">):</span>\n        <span class=\"n\">message</span>
  <span class=\"o\">=</span> <span class=\"n\">event</span><span class=\"p\">[</span><span
  class=\"s2\">&quot;message&quot;</span><span class=\"p\">]</span>\n        <span
  class=\"n\">username</span> <span class=\"o\">=</span> <span class=\"n\">event</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;username&quot;</span><span class=\"p\">]</span>\n\n\n
  \       <span class=\"n\">message_html</span> <span class=\"o\">=</span> <span class=\"sa\">f</span><span
  class=\"s2\">&quot;&lt;div hx-swap-oob=&#39;beforeend:#messages&#39;&gt;&lt;p&gt;&lt;b&gt;</span><span
  class=\"si\">{</span><span class=\"n\">username</span><span class=\"si\">}</span><span
  class=\"s2\">&lt;/b&gt;: </span><span class=\"si\">{</span><span class=\"n\">message</span><span
  class=\"si\">}</span><span class=\"s2\">&lt;/p&gt;&lt;/div&gt;&quot;</span>\n        <span
  class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">send</span><span class=\"p\">(</span>\n            <span class=\"n\">text_data</span><span
  class=\"o\">=</span><span class=\"n\">json</span><span class=\"o\">.</span><span
  class=\"n\">dumps</span><span class=\"p\">(</span>\n                <span class=\"p\">{</span>\n
  \                   <span class=\"s2\">&quot;message&quot;</span><span class=\"p\">:</span>
  <span class=\"n\">message_html</span><span class=\"p\">,</span>\n                    <span
  class=\"s2\">&quot;username&quot;</span><span class=\"p\">:</span> <span class=\"n\">username</span>\n
  \               <span class=\"p\">}</span>\n            <span class=\"p\">)</span>\n
  \       <span class=\"p\">)</span>\n\n    <span class=\"nd\">@sync_to_async</span>\n
  \   <span class=\"k\">def</span> <span class=\"nf\">save_message</span><span class=\"p\">(</span><span
  class=\"bp\">self</span><span class=\"p\">,</span> <span class=\"n\">room</span><span
  class=\"p\">,</span> <span class=\"n\">user</span><span class=\"p\">,</span> <span
  class=\"n\">message</span><span class=\"p\">):</span>\n        <span class=\"n\">room</span>
  <span class=\"o\">=</span> <span class=\"n\">Room</span><span class=\"o\">.</span><span
  class=\"n\">objects</span><span class=\"o\">.</span><span class=\"n\">get</span><span
  class=\"p\">(</span><span class=\"n\">slug</span><span class=\"o\">=</span><span
  class=\"n\">room</span><span class=\"p\">)</span>\n        <span class=\"n\">Message</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">create</span><span class=\"p\">(</span><span class=\"n\">room</span><span
  class=\"o\">=</span><span class=\"n\">room</span><span class=\"p\">,</span> <span
  class=\"n\">user</span><span class=\"o\">=</span><span class=\"n\">user</span><span
  class=\"p\">,</span> <span class=\"n\">message</span><span class=\"o\">=</span><span
  class=\"n\">message</span><span class=\"p\">)</span>\n\n    <span class=\"nd\">@sync_to_async</span>\n
  \   <span class=\"k\">def</span> <span class=\"nf\">add_user</span><span class=\"p\">(</span><span
  class=\"bp\">self</span><span class=\"p\">,</span> <span class=\"n\">room</span><span
  class=\"p\">,</span> <span class=\"n\">user</span><span class=\"p\">):</span>\n
  \       <span class=\"n\">room</span> <span class=\"o\">=</span> <span class=\"n\">Room</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">get</span><span class=\"p\">(</span><span class=\"n\">slug</span><span
  class=\"o\">=</span><span class=\"n\">room</span><span class=\"p\">)</span>\n        <span
  class=\"k\">if</span> <span class=\"n\">user</span> <span class=\"ow\">not</span>
  <span class=\"ow\">in</span> <span class=\"n\">room</span><span class=\"o\">.</span><span
  class=\"n\">users</span><span class=\"o\">.</span><span class=\"n\">all</span><span
  class=\"p\">():</span>\n            <span class=\"n\">room</span><span class=\"o\">.</span><span
  class=\"n\">users</span><span class=\"o\">.</span><span class=\"n\">add</span><span
  class=\"p\">(</span><span class=\"n\">user</span><span class=\"p\">)</span>\n            <span
  class=\"n\">room</span><span class=\"o\">.</span><span class=\"n\">save</span><span
  class=\"p\">()</span>\n\n    <span class=\"nd\">@sync_to_async</span>\n    <span
  class=\"k\">def</span> <span class=\"nf\">remove_user</span><span class=\"p\">(</span><span
  class=\"bp\">self</span><span class=\"p\">,</span> <span class=\"n\">room</span><span
  class=\"p\">,</span> <span class=\"n\">user</span><span class=\"p\">):</span>\n
  \       <span class=\"n\">room</span> <span class=\"o\">=</span> <span class=\"n\">Room</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">get</span><span class=\"p\">(</span><span class=\"n\">slug</span><span
  class=\"o\">=</span><span class=\"n\">room</span><span class=\"p\">)</span>\n        <span
  class=\"k\">if</span> <span class=\"n\">user</span> <span class=\"ow\">in</span>
  <span class=\"n\">room</span><span class=\"o\">.</span><span class=\"n\">users</span><span
  class=\"o\">.</span><span class=\"n\">all</span><span class=\"p\">():</span>\n            <span
  class=\"n\">room</span><span class=\"o\">.</span><span class=\"n\">users</span><span
  class=\"o\">.</span><span class=\"n\">remove</span><span class=\"p\">(</span><span
  class=\"n\">user</span><span class=\"p\">)</span>\n            <span class=\"n\">room</span><span
  class=\"o\">.</span><span class=\"n\">save</span><span class=\"p\">()</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  we have created a few methods like <code>save_message</code>, <code>add_user</code>,
  and <code>remove_user</code> which all are <code>synchronous</code> methods but
  we are using an asynchronous web server, so we add in the <code>sync_to_async</code>
  decorator which wraps a synchronous method to an asynchronous method. Inside the
  methods, we simply perform the database operations like creating a message object,
  and adding or removing the user from the room.</p>\n<p>That's only a few features
  that I have added, you can add to this application and customize them as per your
  needs.</p>\n<p>The code for this chat app is provided in the <a href=\"https://github.com/Mr-Destructive/django-htmx-chat\">GitHub
  repository</a>.</p>\n<h2>Conclusion</h2>\n<p>So, from this post, we were able to
  create a simple chat app (frontendless) with Django and htmx. We used Django channels
  and HTMX to make a chat application without the need to write javascript for the
  client-side connection. Hope you found this tutorial helpful, do give your feedback
  and thoughts on it, I'll be eager to improve this post. Thank you for your patient
  listening. Happy Coding :)</p>\n"
cover: ''
date: 2023-02-05
datetime: 2023-02-05 00:00:00+00:00
description: Building a async, websocket based chat application using django, channels
  and htmx. With this chat application, the user can create and join rooms and send
  and recieve messages.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2023-02-05-Django-HTMX-Chat-App.md
html: "<h1>Django + HTMX Chat application</h1>\n<h2>Introduction</h2>\n<p>In this
  article, we will be creating a Django project, which will be a chat-room kind of
  application. The user needs to authenticate to the app and then and there he/she
  can create or join rooms, every room will have some name and URL associated with
  it. So, the user simply needs to enter the name of the room, which will be unique.
  The user can then simply enter the messages in the chat room. This is a core chat
  application that uses web sockets.</p>\n<p>The unique thing about this app will
  be that, we don't have to write a javascript client. It will all be handled by some
  HTMX magic. The web socket in the backend will definitely handle using Django channels.</p>\n<p>Demo:</p>\n<p><img
  src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-app-demo.webm\" alt=\"Demonstration
  of the Chat App\" /></p>\n<p><a href=\"https://github.com/Mr-Destructive/django-htmx-chat\">GitHub
  Repository</a></p>\n<h3>Requirements:</h3>\n<ul>\n<li>\n<p>Django</p>\n</li>\n<li>\n<p>Django-channels</p>\n</li>\n<li>\n<p>daphne</p>\n</li>\n<li>\n<p>HTMX</p>\n</li>\n<li>\n<p>SQLite
  or any relational database</p>\n</li>\n</ul>\n<p>Also if we want to use the application
  on a large and production scale:</p>\n<ul>\n<li>\n<p>Redis</p>\n</li>\n<li>\n<p>channels_redis</p>\n</li>\n</ul>\n<p>The
  code for this chat app is provided in the <a href=\"https://github.com/Mr-Destructive/django-htmx-chat\">GitHub
  repository</a>.</p>\n<h2>Setup for Django project</h2>\n<p>We will create a simple
  Django project to start with. The project can have two apps, one for auth and the
  other for the chat. You can customize the way you want your existing project accordingly.
  This project is just a demonstration of the use of the chat application with websockets
  and Django channels.</p>\n<p>I'll call the project <code>backchat</code>, you can
  call it whatever you want. We will create a virtual environment and install Django
  in that virtual environment</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>virtualenv
  .venv\n\nFor Linux/macOS:\n<span class=\"nb\">source</span> .venv/bin/activate\n\nFor
  Windows:\n.venv<span class=\"se\">\\s</span>cripts<span class=\"se\">\\a</span>ctivate\n\npip
  install django\ndjango-admin startproject backchat\n<span class=\"nb\">cd</span>
  backchat\n</pre></div>\n\n</pre>\n\n<p>This will set up a base Django project. We
  can now work on the actual implementation of the Django project. Firstly, we will
  start with authentication.</p>\n<h2>Adding basic Authentication and Authorization</h2>\n<h3>Creating
  the accounts app</h3>\n<p>We can separate the authentication of the user from the
  rest of the project, by creating a separate app called <code>user</code> or <code>accounts</code>
  app.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>python
  manage.py startapp accounts\n</pre></div>\n\n</pre>\n\n<h3>Creating a base user
  model</h3>\n<p>We'll start by inheriting the <a href=\"https://docs.djangoproject.com/en/4.1/topics/auth/customizing/#using-a-custom-user-model-when-starting-a-project\">AbstractUser</a>
  the model provided in the <code>django.contrib.auth.models</code> module. The model
  has base fields like <code>username</code> and <code>password</code> which are required
  fields, and <code>email</code>, <code>first_name</code>, <code>last_name</code>,
  etc. which are not mandatory. It is better to create a custom model by inheriting
  the <code>AbstractUser</code> because if in the longer run, we need to make custom
  fields in the user model, it becomes a breeze.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># accounts/models.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">djnago.contrib.auth.models</span> <span class=\"kn\">import</span>
  <span class=\"n\">AbstractUser</span>\n\n\n<span class=\"k\">class</span> <span
  class=\"nc\">User</span><span class=\"p\">(</span><span class=\"n\">AbstractUser</span><span
  class=\"p\">):</span>\n    <span class=\"k\">pass</span>\n</pre></div>\n\n</pre>\n\n<p>This
  creates a basic custom user rather than using the Django built-in user. Next, we
  need to make sure, Django understands the default user is the one we defined in
  the <code>accounts</code> app and not the default <code>User</code>. So, we just
  need to add a field called <code>AUTH_USER_MODEL</code> in the <code>settings.py</code>
  file. The value of this field will be the app name or the module name and the model
  in that python module that we need our custom user model to be set as the default
  user model.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># backchat/settings.py</span>\n\n\n<span class=\"n\">INSALLED_APPS</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"o\">...</span>\n
  \   <span class=\"o\">...</span>\n    <span class=\"s2\">&quot;accounts&quot;</span><span
  class=\"p\">,</span>\n<span class=\"p\">]</span>\n\n<span class=\"c1\"># Append
  to the end of file</span>\n<span class=\"n\">AUTH_USER_MODEL</span> <span class=\"o\">=</span>
  <span class=\"s1\">&#39;accounts.User&#39;</span>\n</pre></div>\n\n</pre>\n\n<h3>Initial
  migrations for the Django project</h3>\n<p>Now, this will get picked up as the default
  user model while referring to any processing related to the user. We can move into
  migrating the changes for the basic Django project and the user model.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>python
  manage.py makemigrations\npython manage.py migrate\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-accounts-migrations.png\" alt=\"initial
  migration for base django and user model\" /></p>\n<h3>Creating register view</h3>\n<p>Further,
  we can add the views like register and profile for the accounts app that can be
  used for the basic authentication. The Login and Logout views are provided in the
  <code>django.contrib.auth.views</code> module, we only have to write our own register
  view. I will be using function-based views to keep it simple to understand but it
  can simply be a class-based view as well.</p>\n<p>To define a view first, we need
  to have form, a user registration form. The form will inherit from the <a href=\"https://docs.djangoproject.com/en/4.1/topics/auth/default/#django.contrib.auth.forms.UserCreationForm\">UserCreationForm</a>
  form which will make the bulk of the task for us at the user registration. We can
  simply then override the Meta class with the fields that we want to display, so
  hence we just keep the <code>username</code> and the <code>password</code> fields.
  The form can be customized by adding in the widget attribute and specifying the
  classes used in them.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># accounts/forms.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">accounts.models</span> <span class=\"kn\">import</span> <span class=\"n\">User</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.contrib.auth.forms</span> <span
  class=\"kn\">import</span> <span class=\"n\">UserCreationForm</span>\n\n<span class=\"k\">class</span>
  <span class=\"nc\">UserRegisterForm</span><span class=\"p\">(</span><span class=\"n\">UserCreationForm</span><span
  class=\"p\">):</span>\n\n    <span class=\"k\">class</span> <span class=\"nc\">Meta</span><span
  class=\"p\">:</span>\n        <span class=\"n\">model</span><span class=\"o\">=</span>
  <span class=\"n\">User</span>\n        <span class=\"n\">fields</span> <span class=\"o\">=</span>
  <span class=\"p\">[</span><span class=\"s1\">&#39;username&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;password1&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;password2&#39;</span><span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>This will give us the <code>UserRegisterForm</code>
  form that can be used for displaying in the register view that we will create in
  the next step.</p>\n<p>We will have to create the register view that will render
  the form for user registration and will also process the form submission.</p>\n<pre
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># accounts/views.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.contrib</span> <span class=\"kn\">import</span> <span class=\"n\">messages</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.shortcuts</span> <span class=\"kn\">import</span>
  <span class=\"n\">redirect</span><span class=\"p\">,</span> <span class=\"n\">render</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">accounts.forms</span> <span class=\"kn\">import</span>
  <span class=\"n\">UserRegisterForm</span>\n\n<span class=\"k\">def</span> <span
  class=\"nf\">register</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">):</span>\n    <span class=\"k\">if</span> <span class=\"n\">request</span><span
  class=\"o\">.</span><span class=\"n\">method</span> <span class=\"o\">==</span>
  <span class=\"s2\">&quot;POST&quot;</span><span class=\"p\">:</span>\n        <span
  class=\"n\">form</span> <span class=\"o\">=</span> <span class=\"n\">UserRegisterForm</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">POST</span><span class=\"p\">)</span>\n        <span class=\"k\">if</span>
  <span class=\"n\">form</span><span class=\"o\">.</span><span class=\"n\">is_valid</span><span
  class=\"p\">():</span>\n            <span class=\"n\">form</span><span class=\"o\">.</span><span
  class=\"n\">save</span><span class=\"p\">()</span>\n            <span class=\"n\">username</span>
  <span class=\"o\">=</span> <span class=\"n\">form</span><span class=\"o\">.</span><span
  class=\"n\">cleaned_data</span><span class=\"o\">.</span><span class=\"n\">get</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;username&quot;</span><span class=\"p\">)</span>\n
  \           <span class=\"n\">messages</span><span class=\"o\">.</span><span class=\"n\">success</span><span
  class=\"p\">(</span>\n                <span class=\"n\">request</span><span class=\"p\">,</span>
  <span class=\"sa\">f</span><span class=\"s2\">&quot;Your account has been created!
  You are now able to log in&quot;</span>\n            <span class=\"p\">)</span>\n
  \           <span class=\"k\">return</span> <span class=\"n\">redirect</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;login&quot;</span><span class=\"p\">)</span>\n
  \   <span class=\"k\">else</span><span class=\"p\">:</span>\n        <span class=\"n\">form</span>
  <span class=\"o\">=</span> <span class=\"n\">UserRegisterForm</span><span class=\"p\">()</span>\n
  \       <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s2\">&quot;accounts/register.html&quot;</span><span
  class=\"p\">,</span> <span class=\"p\">{</span><span class=\"s2\">&quot;form&quot;</span><span
  class=\"p\">:</span> <span class=\"n\">form</span><span class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p>The
  above register view has two cases, one for the user requesting the registration
  form and the second request when the user submits the form. So, when the user makes
  a get request, we load an empty form <code>UserRegisterForm</code> and render the
  <code>register</code> template with the form. We will make the templates later.</p>\n<p>So,
  the template is just rendered when the user wants to register and when the user
  submits the form(sends a post request) we get the details from the post request
  and make it an instance of <code>UserRegisterForm</code> and save the form if it
  is valid. We then redirect the user to the login view (will use the default one
  in the next section). We parse the message as the indication that the user was created.</p>\n<h3>Adding
  URLs for Authentication and Authorization</h3>\n<p>Once, we have the register view
  setup, we can also add login and logout views in the app. But, we don't have to
  write it ourselves, we can override them if needed, but we'll keep the default ones.
  We will use the <a href=\"https://docs.djangoproject.com/en/4.1/topics/auth/default/#django.contrib.auth.views.LoginView\">LoginView</a>
  and the <a href=\"https://docs.djangoproject.com/en/4.1/topics/auth/default/#django.contrib.auth.views.LogoutView\">LogoutView</a>
  view which are class-based views provided in the <code>django.contrib.auth.views</code>
  module. We will provide the respective templates for each of these views.</p>\n<pre
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># accounts/urls.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.urls</span> <span class=\"kn\">import</span> <span class=\"n\">path</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.contrib.auth</span> <span class=\"kn\">import</span>
  <span class=\"n\">views</span> <span class=\"k\">as</span> <span class=\"n\">auth_views</span>\n<span
  class=\"kn\">import</span> <span class=\"nn\">user.views</span> <span class=\"k\">as</span>
  <span class=\"nn\">user_views</span>\n\n<span class=\"n\">urlpatterns</span> <span
  class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;register/&quot;</span><span class=\"p\">,</span>
  <span class=\"n\">user_views</span><span class=\"o\">.</span><span class=\"n\">register</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;register&quot;</span><span class=\"p\">),</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span>\n        <span class=\"s2\">&quot;login/&quot;</span><span
  class=\"p\">,</span>\n        <span class=\"n\">auth_views</span><span class=\"o\">.</span><span
  class=\"n\">LoginView</span><span class=\"o\">.</span><span class=\"n\">as_view</span><span
  class=\"p\">(</span><span class=\"n\">template_name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;accounts/login.html&quot;</span><span class=\"p\">),</span>\n
  \       <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s2\">&quot;login&quot;</span><span
  class=\"p\">,</span>\n    <span class=\"p\">),</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span>\n        <span class=\"s2\">&quot;logout/&quot;</span><span
  class=\"p\">,</span>\n        <span class=\"n\">auth_views</span><span class=\"o\">.</span><span
  class=\"n\">LogoutView</span><span class=\"o\">.</span><span class=\"n\">as_view</span><span
  class=\"p\">(</span><span class=\"n\">template_name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;accounts/logout.html&quot;</span><span class=\"p\">),</span>\n
  \       <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s2\">&quot;logout&quot;</span><span
  class=\"p\">,</span>\n    <span class=\"p\">),</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>We
  have named the URLs as <code>register</code>, <code>login</code>, and <code>logout</code>
  so that we can use them while rendering the links for them in the templates. Now,
  we also need to include the URLs from the accounts app in the project URLs. We can
  do that by using the <code>include</code> method and specifying the app name with
  the module where the urlpatterns are located.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># backchat/urls.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.contrib</span> <span class=\"kn\">import</span> <span class=\"n\">admin</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.urls</span> <span class=\"kn\">import</span>
  <span class=\"n\">include</span><span class=\"p\">,</span> <span class=\"n\">path</span>\n\n<span
  class=\"n\">urlpatterns</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;admin/&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">admin</span><span class=\"o\">.</span><span
  class=\"n\">site</span><span class=\"o\">.</span><span class=\"n\">urls</span><span
  class=\"p\">),</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;auth/&quot;</span><span class=\"p\">,</span> <span class=\"n\">include</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;accounts.urls&quot;</span><span class=\"p\">)),</span>\n<span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>So, we have routed the <code>/auth</code>
  path to include all the URLs in the accounts app. So, the login view will be at
  the URL <code>/auth/login/</code> , and so on.</p>\n<p>Also, we need to add the
  <code>LOGIN_REDIRECT_URL</code> and <code>LOGIN_URL</code> for specifying the url
  name which will be redirected once the user is logged in and the default login url
  name respectively.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># backchat / settings.py</span>\n\n\n<span class=\"n\">LOGIN_REDIRECT_URL</span>
  <span class=\"o\">=</span> <span class=\"s2\">&quot;index&quot;</span>\n<span class=\"n\">LOGIN_URL</span>
  <span class=\"o\">=</span> <span class=\"s2\">&quot;login&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>We
  are now almost done with the view and routing part of the accounts app and can move
  into the creation of templates.</p>\n<h3>Adding Templates for authentication views</h3>\n<p>We
  need a few templates that we have been referencing in the views and the URLs of
  the accounts app in the project. So there are a couple of ways to use templates
  in a Django project. I prefer to have a single templates folder in the root of the
  project and have subfolders as the app which will have the templates specific to
  those apps.</p>\n<p>I usually create a <code>base.html</code> file in the templates
  folder and use that for inheriting other templates. So, we will have to change one
  setting in the project to make sure it looks at <code>templates/</code> from the
  root of the project.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># backchat/settings.py</span>\n\nimport os\n\n...\n...\n\n<span class=\"nv\">TEMPLATES</span>
  <span class=\"o\">=</span> <span class=\"o\">[</span>\n    <span class=\"o\">{</span>\n
  \       <span class=\"s2\">&quot;BACKEND&quot;</span>: <span class=\"s2\">&quot;django.template.backends.django.DjangoTemplates&quot;</span>,\n
  \       <span class=\"s2\">&quot;DIRS&quot;</span>: <span class=\"o\">[</span> os.path.join<span
  class=\"o\">(</span>BASE_DIR, <span class=\"s2\">&quot;templates&quot;</span><span
  class=\"o\">)</span>, <span class=\"o\">]</span>,\n        <span class=\"s2\">&quot;APP_DIRS&quot;</span>:
  True,\n        <span class=\"s2\">&quot;OPTIONS&quot;</span>: <span class=\"o\">{</span>\n
  \           <span class=\"s2\">&quot;context_processors&quot;</span>: <span class=\"o\">[</span>\n
  \               <span class=\"s2\">&quot;django.template.context_processors.debug&quot;</span>,\n
  \               <span class=\"s2\">&quot;django.template.context_processors.request&quot;</span>,\n
  \               <span class=\"s2\">&quot;django.contrib.auth.context_processors.auth&quot;</span>,\n
  \               <span class=\"s2\">&quot;django.contrib.messages.context_processors.messages&quot;</span>,\n
  \           <span class=\"o\">]</span>,\n        <span class=\"o\">}</span>,\n    <span
  class=\"o\">}</span>,\n<span class=\"o\">]</span>\n</pre></div>\n\n</pre>\n\n<p>Then
  create the folder in the same path as you see your <code>manage.py</code> file.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>mkdir
  templates\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-mkdir-templates.png\"
  alt=\"Template Set Up\" /></p>\n<h4>Creating the base template</h4>\n<p>The below
  will be the base template used for the chat application, you can custmize it as
  per your needs.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cp\">&lt;!DOCTYPE html&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">meta</span>
  <span class=\"na\">charset</span><span class=\"o\">=</span><span class=\"s\">&quot;utf-8&quot;</span>
  <span class=\"p\">/&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>Chat App<span class=\"p\">&lt;/</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>\n        {% load static %}\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">script</span> <span class=\"na\">src</span><span class=\"o\">=</span><span
  class=\"s\">&quot;https://unpkg.com/htmx.org@1.8.5&quot;</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">script</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span
  class=\"nt\">head</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">body</span><span class=\"p\">&gt;</span>\n        {% if user.is_authenticated
  %}\n            <span class=\"p\">&lt;</span><span class=\"nt\">a</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% url &#39;logout&#39; %}&quot;</span><span
  class=\"p\">&gt;</span>Logout<span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n        {% else %}\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% url &#39;login&#39; %}&quot;</span><span class=\"p\">&gt;</span>Login<span
  class=\"p\">&lt;/</span><span class=\"nt\">a</span><span class=\"p\">&gt;</span>\n
  \       {% endif %}\n        <span class=\"p\">&lt;</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>Back Chat<span class=\"p\">&lt;/</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>\n        {% block base %}\n        {% endblock %}\n    <span
  class=\"p\">&lt;/</span><span class=\"nt\">body</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;/</span><span class=\"nt\">html</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>We
  have included the <a href=\"https://htmx.org/docs/#installing\">htmx</a> package
  as the script into this template as we will be using it in the actual part of chat
  application.</p>\n<h4>Creating the Register Template</h4>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>#
  templates / accounts / register.html\n\n\n{% extends &#39;base.html&#39; %}\n{%
  block base %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">div</span> <span
  class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;content-section&quot;</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">form</span>
  <span class=\"na\">method</span><span class=\"o\">=</span><span class=\"s\">&quot;POST&quot;</span><span
  class=\"p\">&gt;</span>\n            {% csrf_token %}\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">fieldset</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;form-group&quot;</span><span class=\"p\">&gt;</span>\n                <span
  class=\"p\">&lt;</span><span class=\"nt\">legend</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;border-bottom mb-4&quot;</span><span
  class=\"p\">&gt;</span>Register Now<span class=\"p\">&lt;/</span><span class=\"nt\">legend</span><span
  class=\"p\">&gt;</span>\n                {{ form.as_p }}\n            <span class=\"p\">&lt;/</span><span
  class=\"nt\">fieldset</span><span class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">div</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;form-group&quot;</span><span class=\"p\">&gt;</span>\n                <span
  class=\"p\">&lt;</span><span class=\"nt\">button</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;btn btn-outline-info&quot;</span> <span
  class=\"na\">type</span><span class=\"o\">=</span><span class=\"s\">&quot;submit&quot;</span><span
  class=\"p\">&gt;</span>Sign Up<span class=\"p\">&lt;/</span><span class=\"nt\">button</span><span
  class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;/</span><span class=\"nt\">form</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">div</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;border-top
  pt-3&quot;</span><span class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">small</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;text-muted&quot;</span><span class=\"p\">&gt;</span>\n\t\t    Already
  Have An Account? <span class=\"p\">&lt;</span><span class=\"nt\">a</span> <span
  class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;ml-2&quot;</span>
  <span class=\"na\">href</span><span class=\"o\">=</span><span class=\"s\">&quot;{%
  url &#39;login&#39; %}&quot;</span><span class=\"p\">&gt;</span>Log In<span class=\"p\">&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;/</span><span
  class=\"nt\">small</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;/</span><span
  class=\"nt\">div</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span
  class=\"nt\">div</span><span class=\"p\">&gt;</span>\n{% endblock %}\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-register-page.png\" alt=\"User
  Registration Page\" /></p>\n<h4>Creating the Login Template</h4>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>#
  templates / accounts / login.html    \n\n\n{% extends &#39;base.html&#39; %}\n{%
  block base %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">div</span> <span
  class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;content-section&quot;</span>
  <span class=\"na\">id</span><span class=\"o\">=</span><span class=\"s\">&quot;login&quot;</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">form</span>
  <span class=\"na\">method</span><span class=\"o\">=</span><span class=\"s\">&quot;POST&quot;</span>
  <span class=\"p\">&gt;</span>\n            {% csrf_token %}\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">fieldset</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;form-group&quot;</span><span class=\"p\">&gt;</span>\n                <span
  class=\"p\">&lt;</span><span class=\"nt\">legend</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;border-bottom mb-4&quot;</span><span
  class=\"p\">&gt;</span>LOG IN<span class=\"p\">&lt;/</span><span class=\"nt\">legend</span><span
  class=\"p\">&gt;</span>\n                {{ form.as_p }}\n            <span class=\"p\">&lt;/</span><span
  class=\"nt\">fieldset</span><span class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">div</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;form-group&quot;</span><span class=\"p\">&gt;</span>\n                <span
  class=\"p\">&lt;</span><span class=\"nt\">button</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;btn btn-outline-info&quot;</span> <span
  class=\"na\">type</span><span class=\"o\">=</span><span class=\"s\">&quot;submit&quot;</span><span
  class=\"p\">&gt;</span>Log In<span class=\"p\">&lt;/</span><span class=\"nt\">button</span><span
  class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;/</span><span class=\"nt\">form</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">div</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;border-top
  pt-3&quot;</span><span class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">small</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;text-muted&quot;</span><span class=\"p\">&gt;</span>\n                Register
  Here <span class=\"p\">&lt;</span><span class=\"nt\">a</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;ml-2&quot;</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% url &#39;register&#39; %}&quot;</span><span
  class=\"p\">&gt;</span>Sign Up<span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;/</span><span class=\"nt\">small</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n{% endblock %}\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-login-page.png\"
  alt=\"User Login Page\" /></p>\n<p>Creating the Logout Template</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>#
  templates / accounts / logout.html    \n\n\n{% extends &#39;base.html&#39; %}\n{%
  block base %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span
  class=\"p\">&gt;</span>You have been logged out<span class=\"p\">&lt;/</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">div</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;border-top pt-3&quot;</span><span class=\"p\">&gt;</span>\n        <span
  class=\"p\">&lt;</span><span class=\"nt\">small</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;text-muted&quot;</span><span class=\"p\">&gt;</span>\n
  \           <span class=\"p\">&lt;</span><span class=\"nt\">a</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% url &#39;login&#39; %}&quot;</span><span
  class=\"p\">&gt;</span>Log In Again<span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;/</span><span class=\"nt\">small</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n{% endblock %}\n</pre></div>\n\n</pre>\n\n<h2>Install and
  setup channels</h2>\n<p>We will be using channels to create long-running connections,
  it is a wrapper around Django's asynchronous components and allows us to incorporate
  other protocols like web sockets and other asynchronous protocols.</p>\n<p>So, we
  will be using the Django channels package that will allow us to use the WebSocket
  protocol in the chat application. <a href=\"https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API\">WebSocket</a>
  is a communication protocol(set of rules and standards to be followed) that allows
  both the client and server to send and receive messages(communication).</p>\n<p>To
  install Django channels, we can use pip and install channels with daphne which will
  be used as the web server gateway interface for asynchronous web applications.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>pip
  install -U channels<span class=\"o\">[</span><span class=\"s2\">&quot;daphne&quot;</span><span
  class=\"o\">]</span>\n</pre></div>\n\n</pre>\n\n<p>So this will install the latest
  version of channels and daphne. We now have <a href=\"https://channels.readthedocs.io/en/stable/installation.html\">channels</a>
  in our Django project, we simply need to configure the <code>CHANNEL_LAYER</code>
  config for specifying the backend used that can be <code>Redis</code>, <code>In-Memory</code>,
  or others. We need to add <code>channels</code> , <code>daphne</code> to the <code>INSALLED_APPS</code>
  config of the project. Make sure the <code>daphne</code> app is on top of the applications
  list.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># backchat/settings.py</span>\n\n...\n...\n\n<span class=\"nv\">INSALLED_APPS</span>
  <span class=\"o\">=</span> <span class=\"o\">[</span>\n    <span class=\"s2\">&quot;daphne&quot;</span>,\n
  \   ...\n    ...\n    <span class=\"s2\">&quot;channels&quot;</span>,\n<span class=\"o\">]</span>\n\n\n<span
  class=\"nv\">ASGI_APPLICATION</span> <span class=\"o\">=</span> <span class=\"s2\">&quot;backchat.asgi.application&quot;</span>\n\n...\n...\n\n<span
  class=\"c1\"># For InMemory channels</span>\n\n<span class=\"nv\">CHANNEL_LAYERS</span>
  <span class=\"o\">=</span> <span class=\"o\">{</span>\n    <span class=\"s1\">&#39;default&#39;</span>:
  <span class=\"o\">{</span>\n        <span class=\"s2\">&quot;BACKEND&quot;</span>:
  <span class=\"s2\">&quot;channels.layers.InMemoryChannelLayer&quot;</span>,\n    <span
  class=\"o\">}</span>\n<span class=\"o\">}</span>\n\n\n<span class=\"c1\"># For Redis</span>\n\n<span
  class=\"nv\">CHANNEL_LAYERS</span> <span class=\"o\">=</span> <span class=\"o\">{</span>\n
  \   <span class=\"s2\">&quot;default&quot;</span>: <span class=\"o\">{</span>\n
  \       <span class=\"s2\">&quot;BACKEND&quot;</span>: <span class=\"s2\">&quot;asgi_redis.RedisChannelLayer&quot;</span>,\n
  \       <span class=\"s2\">&quot;CONFIG&quot;</span>: <span class=\"o\">{</span>\n
  \           <span class=\"s2\">&quot;hosts&quot;</span>: <span class=\"o\">[(</span><span
  class=\"s2\">&quot;redis-host-url&quot;</span>, <span class=\"m\">6379</span><span
  class=\"o\">)]</span>,\n        <span class=\"o\">}</span>,\n    <span class=\"o\">}</span>,\n<span
  class=\"o\">}</span>\n</pre></div>\n\n</pre>\n\n<h3>Redis Configuration (Optional)</h3>\n<p>You
  can either use the <a href=\"https://channels.readthedocs.io/en/latest/topics/channel_layers.html\">InMemoryChannelLayer</a>
  or you can use them <code>RedisChannelLayer</code> for the backend of the chat app.
  There are other types of backends like <code>Amazon SQS</code> services, <code>RabbitMQ</code>,
  <code>Kafka</code>, <code>Google Cloud Pub/Sub</code>, etc. I will be creating the
  app with only the <code>InMemoryChannelLayer</code> but will provide a guide for
  redis as well, both are quite similar and only have a few nuances.</p>\n<p>We need
  to install <a href=\"https://github.com/django/channels_redis/\">channels_redis</a>
  it for integrating redis in the Django project with channels.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>pip
  install channels_redis\n</pre></div>\n\n</pre>\n\n<p>So, this will make the <code>channels_redis</code>
  package available in the project, we use this package for real-time storage, in
  the case of the chat app, we might use it for storing messages or room details,
  etc.</p>\n<h2>Creating the Chat App</h2>\n<p>Further, we now can create another
  app for handling the rooms and chat application logic. This app will have its own
  models, views, and URLs. Also, we will define consumers and routers just like views
  and URLs but for asynchronous connections. More on that soon.</p>\n<p>So, let's
  create the <code>chat</code> app.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>python
  manage.py startapp chat\n</pre></div>\n\n</pre>\n\n<p>Then we will add the chat
  app to the <code>INSALLED_APPS</code> config.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># backchat/settings.py</span>\n\n<span class=\"n\">INSALLED_APPS</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"o\">...</span>\n
  \   <span class=\"o\">...</span><span class=\"p\">,</span>\n    <span class=\"s2\">&quot;chat&quot;</span><span
  class=\"p\">,</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  this will make sure to load the chat app in the project. Whenever we run any commands
  like migrations or running the server, it keeps the app in the <code>INSALLED_APPS</code>
  checked up.</p>\n<h3>Defining models</h3>\n<p>This is optional, but we'll do it
  for since we are making a Django app. We already have an auth system configured,
  adding rooms and authorizing the users will become easier then.</p>\n<p>So, let's
  create the models for the chat app which will be the <code>Room</code>.</p>\n<pre
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># chat/models.py</span>\n\n\n<span class=\"kn\">from</span> <span class=\"nn\">django.db</span>
  <span class=\"kn\">import</span> <span class=\"n\">models</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">accounts.models</span> <span class=\"kn\">import</span> <span
  class=\"n\">User</span>\n\n<span class=\"k\">class</span> <span class=\"nc\">Room</span><span
  class=\"p\">(</span><span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">Model</span><span class=\"p\">):</span>\n    <span class=\"n\">name</span>
  <span class=\"o\">=</span> <span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">CharField</span><span class=\"p\">(</span><span class=\"n\">max_length</span><span
  class=\"o\">=</span><span class=\"mi\">128</span><span class=\"p\">)</span>\n    <span
  class=\"n\">slug</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">SlugField</span><span class=\"p\">(</span><span
  class=\"n\">unique</span><span class=\"o\">=</span><span class=\"kc\">True</span><span
  class=\"p\">)</span>\n    <span class=\"n\">users</span> <span class=\"o\">=</span>
  <span class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">ManyToManyField</span><span
  class=\"p\">(</span><span class=\"n\">User</span><span class=\"p\">)</span>\n\n
  \   <span class=\"k\">def</span> <span class=\"fm\">__str__</span><span class=\"p\">(</span><span
  class=\"bp\">self</span><span class=\"p\">):</span>\n        <span class=\"k\">return</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">name</span>\n\n\n<span
  class=\"k\">class</span> <span class=\"nc\">Message</span><span class=\"p\">(</span><span
  class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">Model</span><span
  class=\"p\">):</span>\n    <span class=\"n\">room</span> <span class=\"o\">=</span>
  <span class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">ForeignKey</span><span
  class=\"p\">(</span><span class=\"n\">Room</span><span class=\"p\">,</span> <span
  class=\"n\">on_delete</span><span class=\"o\">=</span><span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">CASCADE</span><span class=\"p\">)</span>\n
  \   <span class=\"n\">user</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">ForeignKey</span><span class=\"p\">(</span><span
  class=\"n\">User</span><span class=\"p\">,</span> <span class=\"n\">on_delete</span><span
  class=\"o\">=</span><span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">CASCADE</span><span class=\"p\">)</span>\n    <span class=\"n\">message</span>
  <span class=\"o\">=</span> <span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">TextField</span><span class=\"p\">()</span>\n    <span class=\"n\">created_at</span>
  <span class=\"o\">=</span> <span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">DateTimeField</span><span class=\"p\">(</span><span class=\"n\">auto_now_add</span><span
  class=\"o\">=</span><span class=\"kc\">True</span><span class=\"p\">)</span>\n\n
  \   <span class=\"k\">def</span> <span class=\"fm\">__str__</span><span class=\"p\">(</span><span
  class=\"bp\">self</span><span class=\"p\">):</span>\n        <span class=\"k\">return</span>
  <span class=\"p\">(</span>\n            <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room</span><span class=\"o\">.</span><span class=\"n\">name</span> <span
  class=\"o\">+</span> <span class=\"s2\">&quot; - &quot;</span> <span class=\"o\">+</span>\n
  \           <span class=\"nb\">str</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">user</span><span class=\"o\">.</span><span
  class=\"n\">username</span><span class=\"p\">)</span> <span class=\"o\">+</span>
  <span class=\"s2\">&quot; : &quot;</span> <span class=\"o\">+</span>\n            <span
  class=\"nb\">str</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">message</span><span class=\"p\">)</span>\n
  \       <span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>So, we simply have
  the name which will be taken from the user, and the slug which will be served as
  the URL to the room. The users are set as <a href=\"https://docs.djangoproject.com/en/4.1/ref/models/fields/#django.db.models.ManyToManyField\">ManyToManyField</a>
  since one room can have multiple users and a user can be in multiple rooms. Also
  we are creating the model <code>Message</code> that will be storing the room and
  the user as the foreign key and the actual text as the message, we can improve the
  security by encrypting the text but it's not the point of this article.</p>\n<p>We
  have set the <code>created_at</code> the field which will set the time when the
  object was created. Finally, the dunder string method is used for representing the
  message object as a price of the concatenation of strings of room name, username,
  and the message. This is useful for admin stuff as it makes it easier to read the
  object rather than the default id.</p>\n<p>Now, once the models are designed we
  can migrate the models into the database.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>python manage.py makemigrations\npython
  manage.py migrate\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-migrations.png\"
  alt=\"Chat app migrations\" /></p>\n<p>We now have a decent model structure for
  the chat application. We can now start the crux of the application i.e. consumers
  and routing with channels.</p>\n<h3>Writing consumers and routers for WebSockets</h3>\n<p>So,
  we start with the bare bones provided in the tutorial on the channel <a href=\"https://channels.readthedocs.io/en/stable/tutorial/part_3.html\">documentation</a>.
  We create a class (consumer) called <code>ChatConsumer</code> which inherits from
  the <code>AsyncWebsocketConsumer</code> provided by the <code>channels.generic.websocket</code>
  module. This has a few methods like <code>connect</code>, <code>disconnect</code>,
  and <code>receive</code>. These are the building blocks of a consumer. We define
  these methods as they will be used for communication via the WebSocket protocol
  through the channels interface.</p>\n<p>In the following block of code, we are essentially
  doing the following:</p>\n<ul>\n<li>\n<p>Accepting connection on the requested room
  name</p>\n</li>\n<li>\n<p>Sending and Receiving messages on the room/group</p>\n</li>\n<li>\n<p>Closing
  the WebSocket connection and removing the client from the room/group</p>\n</li>\n</ul>\n<pre
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># chat/consumers.py</span>\n\n\n<span class=\"kn\">import</span> <span
  class=\"nn\">json</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">asgiref.sync</span>
  <span class=\"kn\">import</span> <span class=\"n\">sync_to_async</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">channels.generic.websocket</span> <span class=\"kn\">import</span>
  <span class=\"n\">AsyncWebsocketConsumer</span>\n\n<span class=\"kn\">from</span>
  <span class=\"nn\">chat.models</span> <span class=\"kn\">import</span> <span class=\"n\">Room</span><span
  class=\"p\">,</span> <span class=\"n\">Message</span>\n\n\n<span class=\"k\">class</span>
  <span class=\"nc\">ChatConsumer</span><span class=\"p\">(</span><span class=\"n\">AsyncWebsocketConsumer</span><span
  class=\"p\">):</span>\n    <span class=\"k\">async</span> <span class=\"k\">def</span>
  <span class=\"nf\">connect</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"p\">):</span>\n        <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_name</span> <span class=\"o\">=</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">scope</span><span class=\"p\">[</span><span
  class=\"s2\">&quot;url_route&quot;</span><span class=\"p\">][</span><span class=\"s2\">&quot;kwargs&quot;</span><span
  class=\"p\">][</span><span class=\"s2\">&quot;room_slug&quot;</span><span class=\"p\">]</span>\n
  \       <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">room_group_name</span>
  <span class=\"o\">=</span> <span class=\"sa\">f</span><span class=\"s2\">&quot;chat_</span><span
  class=\"si\">{</span><span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_name</span><span class=\"si\">}</span><span class=\"s2\">&quot;</span>\n
  \       <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">user</span>
  <span class=\"o\">=</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">scope</span><span class=\"p\">[</span><span class=\"s2\">&quot;user&quot;</span><span
  class=\"p\">]</span>\n\n        <span class=\"k\">await</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">channel_layer</span><span class=\"o\">.</span><span
  class=\"n\">group_add</span><span class=\"p\">(</span>\n            <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">room_group_name</span><span class=\"p\">,</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">channel_name</span>\n
  \       <span class=\"p\">)</span>\n\n        <span class=\"k\">await</span> <span
  class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">accept</span><span
  class=\"p\">()</span>\n\n    <span class=\"k\">async</span> <span class=\"k\">def</span>
  <span class=\"nf\">disconnect</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"p\">,</span> <span class=\"n\">close_code</span><span class=\"p\">):</span>\n
  \       <span class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">channel_layer</span><span class=\"o\">.</span><span class=\"n\">group_discard</span><span
  class=\"p\">(</span>\n            <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_group_name</span><span class=\"p\">,</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">channel_name</span>\n        <span class=\"p\">)</span>\n\n
  \   <span class=\"k\">async</span> <span class=\"k\">def</span> <span class=\"nf\">receive</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"p\">,</span> <span
  class=\"n\">text_data</span><span class=\"p\">):</span>\n        <span class=\"n\">text_data_json</span>
  <span class=\"o\">=</span> <span class=\"n\">json</span><span class=\"o\">.</span><span
  class=\"n\">loads</span><span class=\"p\">(</span><span class=\"n\">text_data</span><span
  class=\"p\">)</span>\n        <span class=\"n\">message</span> <span class=\"o\">=</span>
  <span class=\"n\">text_data_json</span><span class=\"p\">[</span><span class=\"s2\">&quot;message&quot;</span><span
  class=\"p\">]</span>\n        <span class=\"n\">username</span> <span class=\"o\">=</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">user</span><span
  class=\"o\">.</span><span class=\"n\">username</span>\n        \n        <span class=\"k\">await</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">channel_layer</span><span
  class=\"o\">.</span><span class=\"n\">group_send</span><span class=\"p\">(</span>\n
  \           <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">room_group_name</span><span
  class=\"p\">,</span> \n            <span class=\"p\">{</span>\n                <span
  class=\"s2\">&quot;type&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;chat_message&quot;</span><span
  class=\"p\">,</span>\n                <span class=\"s2\">&quot;message&quot;</span><span
  class=\"p\">:</span> <span class=\"n\">message</span><span class=\"p\">,</span>\n
  \               <span class=\"s2\">&quot;username&quot;</span><span class=\"p\">:</span>
  <span class=\"n\">username</span><span class=\"p\">,</span>\n            <span class=\"p\">}</span>\n
  \       <span class=\"p\">)</span>\n\n    <span class=\"k\">async</span> <span class=\"k\">def</span>
  <span class=\"nf\">chat_message</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"p\">,</span> <span class=\"n\">event</span><span class=\"p\">):</span>\n
  \       <span class=\"n\">message</span> <span class=\"o\">=</span> <span class=\"n\">event</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;message&quot;</span><span class=\"p\">]</span>\n
  \       <span class=\"n\">username</span> <span class=\"o\">=</span> <span class=\"n\">event</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;username&quot;</span><span class=\"p\">]</span>\n\n
  \       <span class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">send</span><span class=\"p\">(</span>\n            <span class=\"n\">text_data</span><span
  class=\"o\">=</span><span class=\"n\">json</span><span class=\"o\">.</span><span
  class=\"n\">dumps</span><span class=\"p\">(</span>\n                <span class=\"p\">{</span>\n
  \                   <span class=\"s2\">&quot;message&quot;</span><span class=\"p\">:</span>
  <span class=\"n\">message</span><span class=\"p\">,</span>\n                    <span
  class=\"s2\">&quot;username&quot;</span><span class=\"p\">:</span> <span class=\"n\">username</span>\n
  \               <span class=\"p\">}</span>\n            <span class=\"p\">)</span>\n
  \       <span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<h4>Accept the WebSocket
  connection</h4>\n<p>Here, room and group more or less mean the same thing but are
  different in different contexts. For instance, the group refers to the collection
  of clients which are connected to a channel(communication between consumer and web
  socket) and the Room is referring to the collection of clients connected to a single
  connection stream like a literal room. So we can say, the group is a technical term
  and the room is more of a layman's term for the same thing.</p>\n<p>The method <code>connect</code>
  is called when the client establishes a websocket connection. When that happens,
  the function gets the room slug from the URL of the client and stores <code>room_name</code>
  which is a string. It creates a separate variable called <code>room_group_name</code>
  by prepending the <code>chat_</code> string to the <code>room_name</code>, it also
  gets the currently logged-in user from the request. It then adds the <code>channel_name</code>
  to the group named <code>room_group_name</code>. The <code>channel_name</code> is
  a unique identifier to the connection/consumer in the channel. By adding the <code>channel_name</code>,
  the consumer then can broadcast the message to all the channels within the group.
  Finally, the function accepts the connection, and a <strong>webcoket connection
  is established from both ends, connection is sent from the client and is now accepted
  from the backend.</strong></p>\n<h4>Disconnect from the WebSocket connection</h4>\n<p>When
  the client sends a close connection request to the server, the <code>disconnect</code>
  method is triggered and it basically removes the <code>channel_name</code> from
  the group i.e. the group name <code>room_group_name</code> whatever the string it
  happens to be. So, it basically removes the client from the broadcast group and
  hence it can't receive or send messages through the websocket since it has been
  closed from both ends.</p>\n<p>If you would have paid attention to the <code>close_code</code>
  in-method parameter, it is not being used currently. However, we can use it for
  checking why the connection was closed, as the <code>close_code</code> is a numeric
  value just like the status code in the web request for letting the server know the
  reason for disconnecting from the client.</p>\n<h4>Receive a message from the WebSocket
  connection</h4>\n<p>The <code>recieve</code> method is the core of the consumer
  as it is responsible for all the logic and parsing of the message and broadcasting
  of the messages from the clients to the group channels. The function takes in a
  parameter called <code>text_data</code> and it is sent from the client through websocket,
  so it is JSON content. We need to get the actual message from the JSON object or
  any other piece of content from the client. So, we deserialize (convert the JSON
  object to python objects) the received payload, and get the value from the key <code>message</code>.
  The key is the input name or id from the client sending the request through the
  web socket, so it can be different depending on the frontend template(we'll see
  the front end soon as well).</p>\n<p>We get the user from the scope of the consumer
  as we previously initialized it in the connect method. This can be used for understanding
  which user has sent the message, it will be used later on as we send/broadcast the
  message to the group.</p>\n<p>Now, the final piece in the receive method is the
  <code>channel_layer.group_send</code> method, this method as the name suggests is
  used to send or broadcast the received message to the entire group. The method has
  two parameters:</p>\n<ol>\n<li>\n<p>The name of the group</p>\n</li>\n<li>\n<p>The
  JSON body containing the message and other details</p>\n</li>\n</ol>\n<p>The method
  is not directly sending the message but it has a type key in the JSON body which
  will be the function name to call. The function will simply call the other funciton
  mentioned in the type key in the dict. The following keys in the dict will be the
  parameters of that funciton. In this case, the funciton specified in the <code>type</code>
  key is <code>chat_message</code> which takes in the <code>event</code> as the parameter.
  This event will have all the parameters from the <code>group_send</code> method.</p>\n<p>So,
  the <code>chat_message</code> will load in this message, username, and the room
  name and then call the <code>send</code> method which actually sends the message
  as a JSON payload to the WebSocket connection which will be received by all the
  clients in the same group as provided in the <code>room_group_name</code> string.</p>\n<h3>Adding
  Routers for WebSocket connections</h3>\n<p>So, till this point have consumers, which
  are just like views in terms of channels. Now, we need some URL routes to map these
  consumers to a path. So, we will create a file/module called <code>routing.py</code>
  which will look quite similar to the <code>urls.py</code> file. It will have a list
  called <code>websocket_urlpatterns</code> just like <code>urlpatterns</code> with
  the list of <code>path</code>. These paths however are not <code>http</code> routes
  but will serve for the <code>WebSocket</code> path.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># chat / routing.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.urls</span> <span class=\"kn\">import</span> <span class=\"n\">path</span>\n\n<span
  class=\"kn\">from</span> <span class=\"nn\">chat</span> <span class=\"kn\">import</span>
  <span class=\"n\">consumers</span>\n\n<span class=\"n\">websocket_urlpatterns</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;chat/&lt;str:room_slug&gt;/&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">consumers</span><span class=\"o\">.</span><span
  class=\"n\">ChatConsumer</span><span class=\"o\">.</span><span class=\"n\">as_asgi</span><span
  class=\"p\">()),</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above code block, we have defined a URL for the web socket with the path as
  <code>/chat/&lt;room_slug&gt;</code> where room_name will be the <code>slug</code>
  for the room. The path is bound with the consumer-defined in the <code>consumers.py</code>
  module as <code>ChatConsumer</code>. The <code>as_asgi</code> method is used for
  converting a view into an ASGI-compatible view for the WebSocket interface.</p>\n<h3>Setting
  up ASGI Application</h3>\n<p>We are using the ASGI application config rather than
  a typical WSGI application which only works one request at a time. We want the chat
  application to be asynchronous because multiple clients might send and receive the
  messages at the same time, we don't want to make a client wait before the server
  process a message from another client, that's just the reason we are using WebSocket
  protocol.</p>\n<p>So, we need to also make sure, it makes the http request and also
  add our websocket config from the chat app we created in the previous sections.
  So, inside the <code>asgi.py</code> file in the project config module, we need to
  make some changes to include the chat application configurations.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># backchat / asgi.py</span>\n\n\n<span class=\"kn\">import</span> <span
  class=\"nn\">os</span>\n<span class=\"kn\">from</span> <span class=\"nn\">django.core.asgi</span>
  <span class=\"kn\">import</span> <span class=\"n\">get_asgi_application</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">channels.auth</span> <span class=\"kn\">import</span>
  <span class=\"n\">AuthMiddlewareStack</span>\n<span class=\"kn\">from</span> <span
  class=\"nn\">channels.routing</span> <span class=\"kn\">import</span> <span class=\"n\">ProtocolTypeRouter</span><span
  class=\"p\">,</span> <span class=\"n\">URLRouter</span>\n\n<span class=\"n\">os</span><span
  class=\"o\">.</span><span class=\"n\">environ</span><span class=\"o\">.</span><span
  class=\"n\">setdefault</span><span class=\"p\">(</span><span class=\"s1\">&#39;DJANGO_SETTINGS_MODULE&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;backchat.settings&#39;</span><span
  class=\"p\">)</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">chat</span>
  <span class=\"kn\">import</span> <span class=\"n\">routing</span>\n\n\n<span class=\"n\">application</span>
  <span class=\"o\">=</span> <span class=\"n\">ProtocolTypeRouter</span><span class=\"p\">({</span>\n
  \   <span class=\"s1\">&#39;http&#39;</span><span class=\"p\">:</span> <span class=\"n\">get_asgi_application</span><span
  class=\"p\">(),</span>\n    <span class=\"s2\">&quot;websocket&quot;</span><span
  class=\"p\">:</span><span class=\"n\">AuthMiddlewareStack</span><span class=\"p\">(</span>\n
  \       <span class=\"n\">URLRouter</span><span class=\"p\">(</span>\n            <span
  class=\"n\">routing</span><span class=\"o\">.</span><span class=\"n\">websocket_urlpatterns</span>\n
  \       <span class=\"p\">)</span>\n    <span class=\"p\">)</span>\n<span class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p>We
  will override the <code>application</code> config which is a component used for
  routing different types of protocols for the <code>ASGI</code> application. We have
  set the two keys, <code>http</code> and <code>websocket</code> in our application.
  The <code>http</code> type of requests will be served with the <code>get_asgi_application</code>
  application which is used for running the application in an ASGI environment.</p>\n<p>For
  <code>websocket</code> type of requests, we are setting the <a href=\"https://channels.readthedocs.io/en/latest/topics/authentication.html\">AuthMiddlewareStack</a>
  which helps in authenticating the users requesting the WebSocket connection and
  allows only authorized users to make a connection to the application. The <a href=\"https://channels.readthedocs.io/en/stable/topics/routing.html\">URLRouter</a>
  is used for mapping the list of URL patterns with the incoming request. So, this
  basically serves the request URL with the appropriate consumer in the application.
  We are parsing in the <code>websocket_urlpatterns</code> as the list of URLs that
  will be used for the WebSocket connections.</p>\n<p>Now, we run the server, we should
  be seeing the <code>ASGI</code> server serving our application rather than the plain
  WSGI application.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>$ python manage.py runserver\n\nWatching
  for file changes with StatReloader\nPerforming system checks...\n\nSystem check
  identified no issues (0 silenced).\nFebruary 05, 2023 - 09:22:45\nDjango version
  4.1.5, using settings &#39;backchat.settings&#39;\nStarting ASGI/Daphne version
  4.0.0 development server at http://127.0.0.1:8000/\nQuit the server with CONTROL-C.\n</pre></div>\n\n</pre>\n\n<p>The
  application is not complete yet, it might not have most components working functional
  yet. So, we'll now get into making the user interfaces for the application, to create,
  join, and view rooms in the application.</p>\n<h3>Adding Views for Chat Rooms</h3>\n<p>We
  will have a couple of views like create room page, the join room page, and the chat
  room page. We will define each view as a distinct view and all of them will require
  authenticated users.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># chat / views.py</span>\n\n\n<span class=\"kn\">import</span> <span
  class=\"nn\">string</span>\n<span class=\"kn\">import</span> <span class=\"nn\">random</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.contrib.auth.decorators</span>
  <span class=\"kn\">import</span> <span class=\"n\">login_required</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.shortcuts</span> <span class=\"kn\">import</span>
  <span class=\"n\">render</span><span class=\"p\">,</span> <span class=\"n\">reverse</span><span
  class=\"p\">,</span> <span class=\"n\">redirect</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">django.utils.text</span> <span class=\"kn\">import</span> <span
  class=\"n\">slugify</span>\n<span class=\"kn\">from</span> <span class=\"nn\">chat.models</span>
  <span class=\"kn\">import</span> <span class=\"n\">Room</span>\n\n\n<span class=\"nd\">@login_required</span>\n<span
  class=\"k\">def</span> <span class=\"nf\">index</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"n\">slug</span><span
  class=\"p\">):</span>\n    <span class=\"n\">room</span> <span class=\"o\">=</span>
  <span class=\"n\">Room</span><span class=\"o\">.</span><span class=\"n\">objects</span><span
  class=\"o\">.</span><span class=\"n\">get</span><span class=\"p\">(</span><span
  class=\"n\">slug</span><span class=\"o\">=</span><span class=\"n\">slug</span><span
  class=\"p\">)</span>\n    <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;chat/room.html&#39;</span><span class=\"p\">,</span> <span class=\"p\">{</span><span
  class=\"s1\">&#39;name&#39;</span><span class=\"p\">:</span> <span class=\"n\">room</span><span
  class=\"o\">.</span><span class=\"n\">name</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;slug&#39;</span><span class=\"p\">:</span> <span class=\"n\">room</span><span
  class=\"o\">.</span><span class=\"n\">slug</span><span class=\"p\">})</span>\n\n<span
  class=\"nd\">@login_required</span>\n<span class=\"k\">def</span> <span class=\"nf\">room_create</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"k\">if</span> <span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">method</span> <span class=\"o\">==</span> <span class=\"s2\">&quot;POST&quot;</span><span
  class=\"p\">:</span>\n        <span class=\"n\">room_name</span> <span class=\"o\">=</span>
  <span class=\"n\">request</span><span class=\"o\">.</span><span class=\"n\">POST</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;room_name&quot;</span><span class=\"p\">]</span>\n
  \       <span class=\"n\">uid</span> <span class=\"o\">=</span> <span class=\"nb\">str</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;&#39;</span><span class=\"o\">.</span><span
  class=\"n\">join</span><span class=\"p\">(</span><span class=\"n\">random</span><span
  class=\"o\">.</span><span class=\"n\">choices</span><span class=\"p\">(</span><span
  class=\"n\">string</span><span class=\"o\">.</span><span class=\"n\">ascii_letters</span>
  <span class=\"o\">+</span> <span class=\"n\">string</span><span class=\"o\">.</span><span
  class=\"n\">digits</span><span class=\"p\">,</span> <span class=\"n\">k</span><span
  class=\"o\">=</span><span class=\"mi\">4</span><span class=\"p\">)))</span>\n        <span
  class=\"n\">room_slug</span> <span class=\"o\">=</span> <span class=\"n\">slugify</span><span
  class=\"p\">(</span><span class=\"n\">room_name</span> <span class=\"o\">+</span>
  <span class=\"s2\">&quot;_&quot;</span> <span class=\"o\">+</span> <span class=\"n\">uid</span><span
  class=\"p\">)</span>\n        <span class=\"n\">room</span> <span class=\"o\">=</span>
  <span class=\"n\">Room</span><span class=\"o\">.</span><span class=\"n\">objects</span><span
  class=\"o\">.</span><span class=\"n\">create</span><span class=\"p\">(</span><span
  class=\"n\">name</span><span class=\"o\">=</span><span class=\"n\">room_name</span><span
  class=\"p\">,</span> <span class=\"n\">slug</span><span class=\"o\">=</span><span
  class=\"n\">room_slug</span><span class=\"p\">)</span>\n        <span class=\"k\">return</span>
  <span class=\"n\">redirect</span><span class=\"p\">(</span><span class=\"n\">reverse</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;chat&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">kwargs</span><span class=\"o\">=</span><span class=\"p\">{</span><span
  class=\"s1\">&#39;slug&#39;</span><span class=\"p\">:</span> <span class=\"n\">room</span><span
  class=\"o\">.</span><span class=\"n\">slug</span><span class=\"p\">}))</span>\n
  \   <span class=\"k\">else</span><span class=\"p\">:</span>\n        <span class=\"k\">return</span>
  <span class=\"n\">render</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;chat/create.html&#39;</span><span class=\"p\">)</span>\n\n<span
  class=\"nd\">@login_required</span>\n<span class=\"k\">def</span> <span class=\"nf\">room_join</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"k\">if</span> <span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">method</span> <span class=\"o\">==</span> <span class=\"s2\">&quot;POST&quot;</span><span
  class=\"p\">:</span>\n        <span class=\"n\">room_slug</span> <span class=\"o\">=</span>
  <span class=\"n\">request</span><span class=\"o\">.</span><span class=\"n\">POST</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;room_slug&quot;</span><span class=\"p\">]</span>\n
  \       <span class=\"n\">room</span> <span class=\"o\">=</span> <span class=\"n\">Room</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">get</span><span class=\"p\">(</span><span class=\"n\">slug</span><span
  class=\"o\">=</span><span class=\"n\">room_slug</span><span class=\"p\">)</span>\n
  \       <span class=\"k\">return</span> <span class=\"n\">redirect</span><span class=\"p\">(</span><span
  class=\"n\">reverse</span><span class=\"p\">(</span><span class=\"s1\">&#39;chat&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">kwargs</span><span class=\"o\">=</span><span
  class=\"p\">{</span><span class=\"s1\">&#39;slug&#39;</span><span class=\"p\">:</span>
  <span class=\"n\">room</span><span class=\"o\">.</span><span class=\"n\">slug</span><span
  class=\"p\">}))</span>\n    <span class=\"k\">else</span><span class=\"p\">:</span>\n
  \       <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;chat/join.html&#39;</span><span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>In the above views module, we
  have added 3 views namely <code>index</code> as the room page, <code>room_create</code>
  for the room creation page, and the <code>room_join</code> for the room join page.
  The index view is a simple get request to the provided slug of the room, it gets
  the slug from the URL from the request and fetches an object of the room associated
  with that slug. Then it renders the room template with the context variables like
  the name of the room and the slug associated with that room.</p>\n<p>The <code>room_create</code>
  view is a simple two-case view that either can render the room creation page or
  process the submitted form and create the room. Just like we used in the <code>register</code>
  view in the accounts app. When the user will send a <code>GET</code> request to
  the URL which we will map to <code>/create/</code> shortly after this, the user
  will be given a form. So, we will render the <code>create.html</code> template.
  We will create the html template shortly.\nIf the user has sent a <code>POST</code>
  request to the view via the <code>/create</code> URL, we will fetch the name field
  in the sent request and create a unique identifier with the name of the room. We
  will slugify the concatenation of the name with the uid and will set it as the slug
  of the room. We will then simply create the room and redirect the user to the <code>room</code>
  page.</p>\n<p>The <code>room_join</code> view also is a two-case view, where the
  user can either request the join room form or send a slug with the form submission.
  If the user is requesting a form, we will render the <code>join.html</code> template.
  If the user is submitting the form, we will fetch the room based on the slug provided
  and redirect the user to the <code>room</code> page.</p>\n<p>So, the <code>room_join</code>
  and <code>room_create</code> views are quite similar, we are fetching an already
  existing room in the case of the join view and creating a new instance of room in
  the create view. Now, we will connect the views to the URLs and finally get to the
  templates.</p>\n<h3>Connecting Views and URLs</h3>\n<p>We have three views to route
  to a URL. But we will also have one additional URL which will be the home page for
  the application, on that page the user can choose to either join or create a room.
  We have the room creation, join the room and the room view to be mapped in this
  URL routing of the app.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># chat / urls.py</span>\n\n\n<span class=\"kn\">from</span> <span class=\"nn\">django.urls</span>
  <span class=\"kn\">import</span> <span class=\"n\">path</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">django.views.generic</span> <span class=\"kn\">import</span>
  <span class=\"n\">TemplateView</span>\n<span class=\"kn\">from</span> <span class=\"nn\">chat</span>
  <span class=\"kn\">import</span> <span class=\"n\">views</span>\n\n\n<span class=\"n\">urlpatterns</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;&quot;</span><span class=\"p\">,</span>
  <span class=\"n\">TemplateView</span><span class=\"o\">.</span><span class=\"n\">as_view</span><span
  class=\"p\">(</span><span class=\"n\">template_name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;base.html&quot;</span><span class=\"p\">),</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;index&#39;</span><span class=\"p\">),</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;room/&lt;str:slug&gt;/&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">index</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;chat&#39;</span><span class=\"p\">),</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;create/&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">room_create</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;room-create&#39;</span><span class=\"p\">),</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;join/&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">room_join</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;room-join&#39;</span><span class=\"p\">),</span>\n<span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>So, the first route is the home
  page view called <code>index</code>, we have used the <a href=\"https://docs.djangoproject.com/en/4.1/ref/class-based-views/base/#templateview\">TemplateView</a>
  which will simply render the template provided. We don't have to create a separate
  view for just rendering a template. We already have defined the <code>base.html</code>
  template while setting up the <code>accounts</code> app. This will be the same template,
  we will add some more content to the template later on. The URL mapped is the <code>/</code>
  since we will add the URLs of this app to the project URLs with an empty <code>&quot;&quot;</code>
  path.</p>\n<p>The second route is used for the room index page, i.e. where the user
  will be able to send and receive messages. The path is assigned as <code>/room/&lt;str:slug&gt;/</code>
  indicating a parameter called slug of type string will be used in accessing a particular
  room. The URL will be bound to the index view that we created in the views file,
  which fetches the room based on the slug, so here is where the slug will be coming
  from. The name of the URL-View route will be <code>chat</code> but you can keep
  it as per your requirements. The URL name is really handy for use in the templates.</p>\n<p>The
  third route is for the room creation page. The <code>/create/</code> URL will be
  bound to the <code>room_create</code> view, as we discussed, it will serve two purposes,
  one to render the form for creating the room, and the other for sending a <code>POST</code>
  request to the same path for the creating the Room with the name provided. The name
  is not required but helps in identifying and making it user-friendly. The name of
  this URL is set as <code>room-create</code>.</p>\n<p>The final route is for joining
  the room, the <code>/join/</code> URL will be triggering the <code>room_join</code>
  view. Similar to the <code>room-create</code> URL, the URL will render the join
  room form on a <code>GET</code> request, fetch the room with the provided slug and
  redirect to the room page. Here, the slug field in the form will be required for
  actually finding the matching room. The name of the URL route is set as <code>room-join</code>.</p>\n<p>We
  now add the URLs from the chat app to the project URLs. This will make the <code>/</code>
  as the entry point for the chat application URLs.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># backchat / urls.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.contrib</span> <span class=\"kn\">import</span> <span class=\"n\">admin</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.urls</span> <span class=\"kn\">import</span>
  <span class=\"n\">path</span><span class=\"p\">,</span> <span class=\"n\">include</span>\n\n<span
  class=\"n\">urlpatterns</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;admin/&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">admin</span><span class=\"o\">.</span><span
  class=\"n\">site</span><span class=\"o\">.</span><span class=\"n\">urls</span><span
  class=\"p\">),</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;auth/&quot;</span><span class=\"p\">,</span> <span class=\"n\">include</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;accounts.urls&#39;</span><span class=\"p\">)),</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">include</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;chat.urls&#39;</span><span class=\"p\">)),</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>Hence
  the process is completed for the backend to process the message, it then is dependent
  on the client to process and render the message.</p>\n<p>Till HTMX was not a thing!</p>\n<p>We
  won't have to write a single line of javascript to receive and handle the WebSocket
  connection!</p>\n<h3>Creating Templates and adding htmx</h3>\n<p>We now move into
  the actual frontend or creating the template for actually working with the rooms
  and user interaction. We will have three pieces of templates, a room creates the
  page, a room join page, and a room chat page. As these template names suggest, they
  will be used for creating a room with the name, joining the room with the slug,
  and the room chat page where the user will send and receive messages.</p>\n<p>Let/s
  modify the base template first.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cp\">&lt;!DOCTYPE html&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">meta</span>
  <span class=\"na\">charset</span><span class=\"o\">=</span><span class=\"s\">&quot;utf-8&quot;</span>
  <span class=\"p\">/&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>Chat App<span class=\"p\">&lt;/</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>\n        {% load static %}\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">script</span> <span class=\"na\">src</span><span class=\"o\">=</span><span
  class=\"s\">&quot;https://unpkg.com/htmx.org@1.8.5&quot;</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">script</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span
  class=\"nt\">head</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% url &#39;index&#39; %}&quot;</span><span class=\"p\">&gt;</span>Home<span
  class=\"p\">&lt;/</span><span class=\"nt\">a</span><span class=\"p\">&gt;</span>\n
  \       {% if user.is_authenticated %}\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% url &#39;logout&#39; %}&quot;</span><span class=\"p\">&gt;</span>Logout<span
  class=\"p\">&lt;/</span><span class=\"nt\">a</span><span class=\"p\">&gt;</span>\n
  \       {% else %}\n            <span class=\"p\">&lt;</span><span class=\"nt\">a</span>
  <span class=\"na\">href</span><span class=\"o\">=</span><span class=\"s\">&quot;{%
  url &#39;login&#39; %}&quot;</span><span class=\"p\">&gt;</span>Login<span class=\"p\">&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n        {% endif %}\n    <span
  class=\"p\">&lt;</span><span class=\"nt\">body</span> <span class=\"na\">hx-ext</span><span
  class=\"o\">=</span><span class=\"s\">&quot;ws&quot;</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;</span><span class=\"nt\">h1</span><span class=\"p\">&gt;</span>Back
  Chat<span class=\"p\">&lt;/</span><span class=\"nt\">h1</span><span class=\"p\">&gt;</span>\n
  \       {% block base %}\n            <span class=\"p\">&lt;</span><span class=\"nt\">a</span>
  <span class=\"na\">href</span><span class=\"o\">=</span><span class=\"s\">&quot;{%
  url &#39;room-join&#39; %}&quot;</span><span class=\"p\">&gt;</span>Join Room<span
  class=\"p\">&lt;/</span><span class=\"nt\">a</span><span class=\"p\">&gt;</span>\n
  \           <span class=\"p\">&lt;</span><span class=\"nt\">a</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% url &#39;room-create&#39; %}&quot;</span><span
  class=\"p\">&gt;</span>Create Room<span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n        {% endblock %}\n    <span class=\"p\">&lt;/</span><span
  class=\"nt\">body</span><span class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span
  class=\"nt\">html</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-home-page.png\" alt=\"Chat
  App Home Page\" /></p>\n<h4>Create Room Template</h4>\n<p>We will have to render
  the form with a field like <code>name</code> for setting it as the name of the room,
  it is not required but again, it makes it easier for the user to find the name of
  the room a bit more friendly than random characters.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>#
  templates / chat / create.html\n\n\n{% extends &#39;base.html&#39; %}\n\n{% block
  base %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">form</span> <span class=\"na\">method</span><span
  class=\"o\">=</span><span class=\"s\">&#39;post&#39;</span> <span class=\"na\">action</span><span
  class=\"o\">=</span><span class=\"s\">&#39;{% url &#39;</span><span class=\"na\">room-create</span><span
  class=\"err\">&#39;</span> <span class=\"err\">%}&#39;</span><span class=\"p\">&gt;</span>\n
  \       {% csrf_token %}\n        <span class=\"p\">&lt;</span><span class=\"nt\">input</span>
  <span class=\"na\">name</span><span class=\"o\">=</span><span class=\"s\">&#39;room_name&#39;</span>
  <span class=\"na\">id</span><span class=\"o\">=</span><span class=\"s\">&#39;room_name&#39;</span>
  <span class=\"na\">placeholder</span><span class=\"o\">=</span><span class=\"s\">&#39;Room
  Name&#39;</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">input</span> <span class=\"na\">type</span><span class=\"o\">=</span><span
  class=\"s\">&#39;submit&#39;</span> <span class=\"na\">id</span><span class=\"o\">=</span><span
  class=\"s\">&quot;submit&quot;</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span
  class=\"nt\">form</span><span class=\"p\">&gt;</span>\n{% endblock %}\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-create-room-page.png\" alt=\"Chat
  Room Create Page\" /></p>\n<p>The template is inherited from the <code>base.html</code>
  template and we render a form with the <code>room_name</code> input. The form will
  be submitted to the URL named <code>room-create</code> hence it will send a <code>POST</code>
  request to the server where it will create the room and further process the request.</p>\n<h4>Join
  Room Template</h4>\n<p>The join room template is similar to the create room template
  except it gets the slug of the room from the user rather than the name which is
  not a unique one to join the room.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>#
  templates / chat / join.html\n\n\n{% extends &#39;base.html&#39; %}\n\n{% block
  base %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">form</span> <span class=\"na\">method</span><span
  class=\"o\">=</span><span class=\"s\">&#39;post&#39;</span> <span class=\"na\">action</span><span
  class=\"o\">=</span><span class=\"s\">&#39;{% url &#39;</span><span class=\"na\">room-join</span><span
  class=\"err\">&#39;</span> <span class=\"err\">%}&#39;</span><span class=\"p\">&gt;</span>\n
  \       {% csrf_token %}\n        <span class=\"p\">&lt;</span><span class=\"nt\">input</span>
  <span class=\"na\">name</span><span class=\"o\">=</span><span class=\"s\">&#39;room_slug&#39;</span>
  <span class=\"na\">id</span><span class=\"o\">=</span><span class=\"s\">&#39;room_slug&#39;</span>
  <span class=\"na\">required</span><span class=\"o\">=</span><span class=\"s\">&#39;true&#39;</span>
  <span class=\"na\">placeholder</span><span class=\"o\">=</span><span class=\"s\">&#39;Room
  Code&#39;</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">input</span> <span class=\"na\">type</span><span class=\"o\">=</span><span
  class=\"s\">&#39;submit&#39;</span> <span class=\"na\">id</span><span class=\"o\">=</span><span
  class=\"s\">&quot;submit&quot;</span><span class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span
  class=\"nt\">form</span><span class=\"p\">&gt;</span>\n{% endblock %}\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-join-room-page.png\" alt=\"Chat
  Room Join Page\" /></p>\n<p>The form will be submitted to the URL named <code>room-join</code>
  hence it will send a <code>POST</code> request to the server where it will fetch
  the room and further process the request.</p>\n<h3>Room Template (HTMX code)</h3>\n<p>Now,
  time for the actual ingredient in the application, some HTMX magic!</p>\n<p>This
  template, as the two templates above inherit from the base template, that's the
  same thing. But it has a special <code>div</code> with the attribute <a href=\"https://htmx.org/attributes/hx-ws/\">hx-ws</a>
  which is used for using attributes related to the web socket in the htmx library.
  The <code>connect</code> value is used for connecting to a WebSocket. The value
  of the attribute must be set to the URL of the WebSocket to be connected. In our
  case, it is the URL path from the <code>routing</code> app as <code>/chat/&lt;room_slug&gt;/</code>.
  This simply will connect the client to the WebSocket from the backend. The other
  important attribute is the <code>send</code> which is used for sending a message
  to the connected web socket.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>#
  templates / chat / room.html\n\n\n{% extends &#39;base.html&#39; %}\n\n{% block
  base %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>{{
  name }}<span class=\"p\">&lt;/</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;</span><span class=\"nt\">div</span> <span class=\"na\">hx-ws</span><span
  class=\"o\">=</span><span class=\"s\">&quot;connect:/chat/{{ slug }}/&quot;</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">form</span>
  <span class=\"na\">hx-ws</span><span class=\"o\">=</span><span class=\"s\">&quot;send:submit&quot;</span><span
  class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span class=\"nt\">input</span>
  <span class=\"na\">name</span><span class=\"o\">=</span><span class=\"s\">&quot;message&quot;</span><span
  class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span class=\"nt\">input</span>
  <span class=\"na\">type</span><span class=\"o\">=</span><span class=\"s\">&quot;submit&quot;</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;/</span><span class=\"nt\">form</span><span
  class=\"p\">&gt;</span>\n     <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n     <span class=\"p\">&lt;</span><span class=\"nt\">div</span>
  <span class=\"na\">id</span><span class=\"o\">=</span><span class=\"s\">&#39;messages&#39;</span><span
  class=\"p\">&gt;&lt;/</span><span class=\"nt\">div</span><span class=\"p\">&gt;</span>\n{%
  endblock %}\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-room-page.png\"
  alt=\"Chat Room Page\" /></p>\n<p>NOTE: The template has a div with the id <code>messages</code>
  which will be very important for sending the messages from the WebSocket to the
  client, so more on that when we use the HTMX part.</p>\n<p>For testing this template,
  you can create a room, and that will redirect you to the room template as we have
  seen in the views for the room creation. If you see something like <code>WebSocket
  CONNECT</code> it means, that the application has been able to establish a WebSocket
  connection to the backend, and we can be ready to accept messages and other stuff.</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>HTTP GET /chat/room/def_teas/
  200 [0.03, 127.0.0.1:38660]\nWebSocket HANDSHAKING /chat/def_teas/ [127.0.0.1:38666]\nWebSocket
  CONNECT /chat/def_teas/ [127.0.0.1:38666]\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-asgi-server.png\"
  alt=\"Django ASGI server websocket connection\" /></p>\n<p>Till this point, we should
  have a running and almost complete application, though we just have a minor part
  missing that will be the most important part.</p>\n<h3>Sending HTML response from
  backend for htmx</h3>\n<p>We will be sending a fragment of HTML from the backend
  when the user sends a message, to broadcast it to the group. Let's make some changes
  to the application, especially to the receive method in the <code>ChatConsumer</code>
  of the chat application.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># chat / consumers.py</span>\n    \n\n    <span class=\"o\">...</span>\n
  \   <span class=\"o\">...</span>\n\n    <span class=\"k\">async</span> <span class=\"k\">def</span>
  <span class=\"nf\">receive</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"p\">,</span> <span class=\"n\">text_data</span><span class=\"p\">):</span>\n
  \       <span class=\"n\">text_data_json</span> <span class=\"o\">=</span> <span
  class=\"n\">json</span><span class=\"o\">.</span><span class=\"n\">loads</span><span
  class=\"p\">(</span><span class=\"n\">text_data</span><span class=\"p\">)</span>\n
  \       <span class=\"n\">message</span> <span class=\"o\">=</span> <span class=\"n\">text_data_json</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;message&quot;</span><span class=\"p\">]</span>\n
  \       <span class=\"n\">user</span> <span class=\"o\">=</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">user</span>\n        <span class=\"n\">username</span>
  <span class=\"o\">=</span> <span class=\"n\">user</span><span class=\"o\">.</span><span
  class=\"n\">username</span>\n\n        <span class=\"k\">await</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">channel_layer</span><span class=\"o\">.</span><span
  class=\"n\">group_send</span><span class=\"p\">(</span>\n            <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">room_group_name</span><span class=\"p\">,</span>
  \n            <span class=\"p\">{</span>\n                <span class=\"s2\">&quot;type&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;chat_message&quot;</span><span class=\"p\">,</span>\n
  \               <span class=\"s2\">&quot;message&quot;</span><span class=\"p\">:</span>
  <span class=\"n\">message</span><span class=\"p\">,</span>\n                <span
  class=\"s2\">&quot;username&quot;</span><span class=\"p\">:</span> <span class=\"n\">username</span><span
  class=\"p\">,</span>\n            <span class=\"p\">}</span>\n        <span class=\"p\">)</span>\n\n
  \   <span class=\"k\">async</span> <span class=\"k\">def</span> <span class=\"nf\">chat_message</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"p\">,</span> <span
  class=\"n\">event</span><span class=\"p\">):</span>\n        <span class=\"n\">message</span>
  <span class=\"o\">=</span> <span class=\"n\">event</span><span class=\"p\">[</span><span
  class=\"s2\">&quot;message&quot;</span><span class=\"p\">]</span>\n        <span
  class=\"n\">username</span> <span class=\"o\">=</span> <span class=\"n\">event</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;username&quot;</span><span class=\"p\">]</span>\n\n
  \       <span class=\"c1\"># This is the crucial part of the application</span>\n
  \       <span class=\"n\">message_html</span> <span class=\"o\">=</span> <span class=\"sa\">f</span><span
  class=\"s2\">&quot;&lt;div hx-swap-oob=&#39;beforeend:#messages&#39;&gt;&lt;p&gt;&lt;b&gt;</span><span
  class=\"si\">{</span><span class=\"n\">username</span><span class=\"si\">}</span><span
  class=\"s2\">&lt;/b&gt;: </span><span class=\"si\">{</span><span class=\"n\">message</span><span
  class=\"si\">}</span><span class=\"s2\">&lt;/p&gt;&lt;/div&gt;&quot;</span>\n        <span
  class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">send</span><span class=\"p\">(</span>\n            <span class=\"n\">text_data</span><span
  class=\"o\">=</span><span class=\"n\">json</span><span class=\"o\">.</span><span
  class=\"n\">dumps</span><span class=\"p\">(</span>\n                <span class=\"p\">{</span>\n
  \                   <span class=\"s2\">&quot;message&quot;</span><span class=\"p\">:</span>
  <span class=\"n\">message_html</span><span class=\"p\">,</span>\n                    <span
  class=\"s2\">&quot;username&quot;</span><span class=\"p\">:</span> <span class=\"n\">username</span>\n
  \               <span class=\"p\">}</span>\n            <span class=\"p\">)</span>\n
  \       <span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-room-msg-1.png\"
  alt=\"Chat Room Message\" />\n<img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-room-msg-2.png\"
  alt=\"Chat Room Message 2 Users\" />\n<img src=\"https://meetgor-cdn.pages.dev/django-htmx/chat-room-msg-3.png\"
  alt=\"Chat Room Message\" /></p>\n<p>In the above snippet, we are just changing
  the final message object to include some HTML just simple. The HTML however has
  home htmx attributes like <a href=\"https://htmx.org/attributes/hx-swap-oob/\">hx-swap-oob</a>
  which will just update the specified DOM element to the content in the div. In this
  case, the DOM element is <code>#message</code> which is the id message present in
  the room template. We basically add the username and the message into the same id
  by appending it to the element. That's it, it would work and it would start showing
  the messages from the connected clients and broadcast them as well.</p>\n<p>There
  are some things to keep in mind while using htmx in the long run especially when
  the htmx 2.0 is released, it will have <code>ws</code> as a separate extension.
  It will have a bit of a different syntax than above. I have tried the latest version
  but doesn't seem to work. I'll just leave a few snippets for your understanding
  of the problem.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>#
  templates / chat / room.html\n\n\n{% extends &#39;base.html&#39; %}\n\n{% block
  base %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>{{
  name }}<span class=\"p\">&lt;/</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;</span><span class=\"nt\">div</span> <span class=\"na\">hx-ext</span><span
  class=\"o\">=</span><span class=\"s\">&quot;ws&quot;</span> <span class=\"na\">ws-connect</span><span
  class=\"o\">=</span><span class=\"s\">&quot;/chat/{{ slug }}/&quot;</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">form</span>
  <span class=\"na\">ws-send</span><span class=\"p\">&gt;</span>\n            <span
  class=\"p\">&lt;</span><span class=\"nt\">input</span> <span class=\"na\">name</span><span
  class=\"o\">=</span><span class=\"s\">&quot;message&quot;</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;/</span><span class=\"nt\">form</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;</span><span class=\"nt\">div</span> <span class=\"na\">id</span><span
  class=\"o\">=</span><span class=\"s\">&#39;messages&#39;</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">div</span><span class=\"p\">&gt;</span>\n{% endblock %}\n</pre></div>\n\n</pre>\n\n<p>I
  have added, the <code>hx-ext</code> as <code>ws</code> which is a htmx <a href=\"https://htmx.org/extensions/web-sockets/\">extension
  for websocket</a>. This extension has websocket-specific attributes like <code>ws-connect</code>
  and <code>ws-send</code>. I have tried a combination like changing the htmx versions,
  adding submit value as the <code>ws-send</code> attribute, etc, but no results yet.
  I have opened a <a href=\"https://github.com/bigskysoftware/htmx/discussions/1231\">discussion</a>
  on GitHub for this issue, you can express your solution or views there.</p>\n<h3>Adding
  some utility features for the chat app</h3>\n<p>We can save messages, add and remove
  the users from the room according to the connection, and other stuff that can make
  this a fully-fledged app. So, I have made a few changes to the chat consumers for
  saving the messages and also updating the room with the users in the room.</p>\n<pre
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
  \       \n<div class='language-bar'>python</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\"># chat / consumers.py</span>\n\n\n<span class=\"kn\">import</span>
  <span class=\"nn\">json</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">asgiref.sync</span>
  <span class=\"kn\">import</span> <span class=\"n\">sync_to_async</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">channels.generic.websocket</span> <span class=\"kn\">import</span>
  <span class=\"n\">AsyncWebsocketConsumer</span>\n\n<span class=\"kn\">from</span>
  <span class=\"nn\">chat.models</span> <span class=\"kn\">import</span> <span class=\"n\">Room</span><span
  class=\"p\">,</span> <span class=\"n\">Message</span>\n\n\n<span class=\"k\">class</span>
  <span class=\"nc\">ChatConsumer</span><span class=\"p\">(</span><span class=\"n\">AsyncWebsocketConsumer</span><span
  class=\"p\">):</span>\n    <span class=\"k\">async</span> <span class=\"k\">def</span>
  <span class=\"nf\">connect</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"p\">):</span>\n        <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_name</span> <span class=\"o\">=</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">scope</span><span class=\"p\">[</span><span
  class=\"s2\">&quot;url_route&quot;</span><span class=\"p\">][</span><span class=\"s2\">&quot;kwargs&quot;</span><span
  class=\"p\">][</span><span class=\"s2\">&quot;room_slug&quot;</span><span class=\"p\">]</span>\n
  \       <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">room_group_name</span>
  <span class=\"o\">=</span> <span class=\"s2\">&quot;chat_</span><span class=\"si\">%s</span><span
  class=\"s2\">&quot;</span> <span class=\"o\">%</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">room_name</span>\n        <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">user</span> <span class=\"o\">=</span> <span
  class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">scope</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;user&quot;</span><span class=\"p\">]</span>\n\n
  \       <span class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">channel_layer</span><span class=\"o\">.</span><span class=\"n\">group_add</span><span
  class=\"p\">(</span>\n            <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_group_name</span><span class=\"p\">,</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">channel_name</span>\n        <span class=\"p\">)</span>\n\n
  \       <span class=\"c1\"># Add the user when the client connects</span>\n        <span
  class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">add_user</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">room_name</span><span class=\"p\">,</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">user</span><span
  class=\"p\">)</span>\n\n        <span class=\"k\">await</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">accept</span><span class=\"p\">()</span>\n\n
  \   <span class=\"k\">async</span> <span class=\"k\">def</span> <span class=\"nf\">disconnect</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"p\">,</span> <span
  class=\"n\">close_code</span><span class=\"p\">):</span>\n\n        <span class=\"c1\">#
  Remove the user when the client disconnects</span>\n        <span class=\"k\">await</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">remove_user</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_name</span><span class=\"p\">,</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">user</span><span class=\"p\">)</span>\n\n
  \       <span class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">channel_layer</span><span class=\"o\">.</span><span class=\"n\">group_discard</span><span
  class=\"p\">(</span>\n            <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_group_name</span><span class=\"p\">,</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">channel_name</span>\n        <span class=\"p\">)</span>\n\n
  \   <span class=\"k\">async</span> <span class=\"k\">def</span> <span class=\"nf\">receive</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"p\">,</span> <span
  class=\"n\">text_data</span><span class=\"p\">):</span>\n        <span class=\"n\">text_data_json</span>
  <span class=\"o\">=</span> <span class=\"n\">json</span><span class=\"o\">.</span><span
  class=\"n\">loads</span><span class=\"p\">(</span><span class=\"n\">text_data</span><span
  class=\"p\">)</span>\n        <span class=\"n\">message</span> <span class=\"o\">=</span>
  <span class=\"n\">text_data_json</span><span class=\"p\">[</span><span class=\"s2\">&quot;message&quot;</span><span
  class=\"p\">]</span>\n        <span class=\"n\">user</span> <span class=\"o\">=</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">user</span>\n
  \       <span class=\"n\">username</span> <span class=\"o\">=</span> <span class=\"n\">user</span><span
  class=\"o\">.</span><span class=\"n\">username</span>\n        <span class=\"n\">room</span>
  <span class=\"o\">=</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_name</span>\n\n        <span class=\"c1\"># Save the message on
  recieving</span>\n        <span class=\"k\">await</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">save_message</span><span class=\"p\">(</span><span
  class=\"n\">room</span><span class=\"p\">,</span> <span class=\"n\">user</span><span
  class=\"p\">,</span> <span class=\"n\">message</span><span class=\"p\">)</span>\n\n
  \       <span class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">channel_layer</span><span class=\"o\">.</span><span class=\"n\">group_send</span><span
  class=\"p\">(</span>\n            <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">room_group_name</span><span class=\"p\">,</span> \n            <span
  class=\"p\">{</span>\n                <span class=\"s2\">&quot;type&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;chat_message&quot;</span><span class=\"p\">,</span>\n
  \               <span class=\"s2\">&quot;message&quot;</span><span class=\"p\">:</span>
  <span class=\"n\">message</span><span class=\"p\">,</span>\n                <span
  class=\"s2\">&quot;username&quot;</span><span class=\"p\">:</span> <span class=\"n\">username</span><span
  class=\"p\">,</span>\n            <span class=\"p\">}</span>\n        <span class=\"p\">)</span>\n\n
  \   <span class=\"k\">async</span> <span class=\"k\">def</span> <span class=\"nf\">chat_message</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"p\">,</span> <span
  class=\"n\">event</span><span class=\"p\">):</span>\n        <span class=\"n\">message</span>
  <span class=\"o\">=</span> <span class=\"n\">event</span><span class=\"p\">[</span><span
  class=\"s2\">&quot;message&quot;</span><span class=\"p\">]</span>\n        <span
  class=\"n\">username</span> <span class=\"o\">=</span> <span class=\"n\">event</span><span
  class=\"p\">[</span><span class=\"s2\">&quot;username&quot;</span><span class=\"p\">]</span>\n\n\n
  \       <span class=\"n\">message_html</span> <span class=\"o\">=</span> <span class=\"sa\">f</span><span
  class=\"s2\">&quot;&lt;div hx-swap-oob=&#39;beforeend:#messages&#39;&gt;&lt;p&gt;&lt;b&gt;</span><span
  class=\"si\">{</span><span class=\"n\">username</span><span class=\"si\">}</span><span
  class=\"s2\">&lt;/b&gt;: </span><span class=\"si\">{</span><span class=\"n\">message</span><span
  class=\"si\">}</span><span class=\"s2\">&lt;/p&gt;&lt;/div&gt;&quot;</span>\n        <span
  class=\"k\">await</span> <span class=\"bp\">self</span><span class=\"o\">.</span><span
  class=\"n\">send</span><span class=\"p\">(</span>\n            <span class=\"n\">text_data</span><span
  class=\"o\">=</span><span class=\"n\">json</span><span class=\"o\">.</span><span
  class=\"n\">dumps</span><span class=\"p\">(</span>\n                <span class=\"p\">{</span>\n
  \                   <span class=\"s2\">&quot;message&quot;</span><span class=\"p\">:</span>
  <span class=\"n\">message_html</span><span class=\"p\">,</span>\n                    <span
  class=\"s2\">&quot;username&quot;</span><span class=\"p\">:</span> <span class=\"n\">username</span>\n
  \               <span class=\"p\">}</span>\n            <span class=\"p\">)</span>\n
  \       <span class=\"p\">)</span>\n\n    <span class=\"nd\">@sync_to_async</span>\n
  \   <span class=\"k\">def</span> <span class=\"nf\">save_message</span><span class=\"p\">(</span><span
  class=\"bp\">self</span><span class=\"p\">,</span> <span class=\"n\">room</span><span
  class=\"p\">,</span> <span class=\"n\">user</span><span class=\"p\">,</span> <span
  class=\"n\">message</span><span class=\"p\">):</span>\n        <span class=\"n\">room</span>
  <span class=\"o\">=</span> <span class=\"n\">Room</span><span class=\"o\">.</span><span
  class=\"n\">objects</span><span class=\"o\">.</span><span class=\"n\">get</span><span
  class=\"p\">(</span><span class=\"n\">slug</span><span class=\"o\">=</span><span
  class=\"n\">room</span><span class=\"p\">)</span>\n        <span class=\"n\">Message</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">create</span><span class=\"p\">(</span><span class=\"n\">room</span><span
  class=\"o\">=</span><span class=\"n\">room</span><span class=\"p\">,</span> <span
  class=\"n\">user</span><span class=\"o\">=</span><span class=\"n\">user</span><span
  class=\"p\">,</span> <span class=\"n\">message</span><span class=\"o\">=</span><span
  class=\"n\">message</span><span class=\"p\">)</span>\n\n    <span class=\"nd\">@sync_to_async</span>\n
  \   <span class=\"k\">def</span> <span class=\"nf\">add_user</span><span class=\"p\">(</span><span
  class=\"bp\">self</span><span class=\"p\">,</span> <span class=\"n\">room</span><span
  class=\"p\">,</span> <span class=\"n\">user</span><span class=\"p\">):</span>\n
  \       <span class=\"n\">room</span> <span class=\"o\">=</span> <span class=\"n\">Room</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">get</span><span class=\"p\">(</span><span class=\"n\">slug</span><span
  class=\"o\">=</span><span class=\"n\">room</span><span class=\"p\">)</span>\n        <span
  class=\"k\">if</span> <span class=\"n\">user</span> <span class=\"ow\">not</span>
  <span class=\"ow\">in</span> <span class=\"n\">room</span><span class=\"o\">.</span><span
  class=\"n\">users</span><span class=\"o\">.</span><span class=\"n\">all</span><span
  class=\"p\">():</span>\n            <span class=\"n\">room</span><span class=\"o\">.</span><span
  class=\"n\">users</span><span class=\"o\">.</span><span class=\"n\">add</span><span
  class=\"p\">(</span><span class=\"n\">user</span><span class=\"p\">)</span>\n            <span
  class=\"n\">room</span><span class=\"o\">.</span><span class=\"n\">save</span><span
  class=\"p\">()</span>\n\n    <span class=\"nd\">@sync_to_async</span>\n    <span
  class=\"k\">def</span> <span class=\"nf\">remove_user</span><span class=\"p\">(</span><span
  class=\"bp\">self</span><span class=\"p\">,</span> <span class=\"n\">room</span><span
  class=\"p\">,</span> <span class=\"n\">user</span><span class=\"p\">):</span>\n
  \       <span class=\"n\">room</span> <span class=\"o\">=</span> <span class=\"n\">Room</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">get</span><span class=\"p\">(</span><span class=\"n\">slug</span><span
  class=\"o\">=</span><span class=\"n\">room</span><span class=\"p\">)</span>\n        <span
  class=\"k\">if</span> <span class=\"n\">user</span> <span class=\"ow\">in</span>
  <span class=\"n\">room</span><span class=\"o\">.</span><span class=\"n\">users</span><span
  class=\"o\">.</span><span class=\"n\">all</span><span class=\"p\">():</span>\n            <span
  class=\"n\">room</span><span class=\"o\">.</span><span class=\"n\">users</span><span
  class=\"o\">.</span><span class=\"n\">remove</span><span class=\"p\">(</span><span
  class=\"n\">user</span><span class=\"p\">)</span>\n            <span class=\"n\">room</span><span
  class=\"o\">.</span><span class=\"n\">save</span><span class=\"p\">()</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  we have created a few methods like <code>save_message</code>, <code>add_user</code>,
  and <code>remove_user</code> which all are <code>synchronous</code> methods but
  we are using an asynchronous web server, so we add in the <code>sync_to_async</code>
  decorator which wraps a synchronous method to an asynchronous method. Inside the
  methods, we simply perform the database operations like creating a message object,
  and adding or removing the user from the room.</p>\n<p>That's only a few features
  that I have added, you can add to this application and customize them as per your
  needs.</p>\n<p>The code for this chat app is provided in the <a href=\"https://github.com/Mr-Destructive/django-htmx-chat\">GitHub
  repository</a>.</p>\n<h2>Conclusion</h2>\n<p>So, from this post, we were able to
  create a simple chat app (frontendless) with Django and htmx. We used Django channels
  and HTMX to make a chat application without the need to write javascript for the
  client-side connection. Hope you found this tutorial helpful, do give your feedback
  and thoughts on it, I'll be eager to improve this post. Thank you for your patient
  listening. Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/django-htmx-chat-cover.png
long_description: In this article, we will be creating a Django project, which will
  be a chat-room kind of application. The user needs to authenticate to the app and
  then and there he/she can create or join rooms, every room will have some name and
  URL associated with
now: 2025-05-01 18:11:33.312528
path: blog/posts/2023-02-05-Django-HTMX-Chat-App.md
prevnext: null
series:
- Django-Series
slug: django-htmx-chat-app
status: published
tags:
- django
- htmx
- python
templateKey: blog-post
title: Creating a Chat Application with Django and HTMX
today: 2025-05-01
---

# Django + HTMX Chat application

## Introduction

In this article, we will be creating a Django project, which will be a chat-room kind of application. The user needs to authenticate to the app and then and there he/she can create or join rooms, every room will have some name and URL associated with it. So, the user simply needs to enter the name of the room, which will be unique. The user can then simply enter the messages in the chat room. This is a core chat application that uses web sockets.

The unique thing about this app will be that, we don't have to write a javascript client. It will all be handled by some HTMX magic. The web socket in the backend will definitely handle using Django channels.

Demo:

![Demonstration of the Chat App](https://meetgor-cdn.pages.dev/django-htmx/chat-app-demo.webm)

[GitHub Repository](https://github.com/Mr-Destructive/django-htmx-chat)

### Requirements:

* Django
    
* Django-channels
    
* daphne
    
* HTMX
    
* SQLite or any relational database
    

Also if we want to use the application on a large and production scale:

* Redis
    
* channels_redis

The code for this chat app is provided in the [GitHub repository](https://github.com/Mr-Destructive/django-htmx-chat).

## Setup for Django project

We will create a simple Django project to start with. The project can have two apps, one for auth and the other for the chat. You can customize the way you want your existing project accordingly. This project is just a demonstration of the use of the chat application with websockets and Django channels.

I'll call the project `backchat`, you can call it whatever you want. We will create a virtual environment and install Django in that virtual environment

```bash
virtualenv .venv

For Linux/macOS:
source .venv/bin/activate

For Windows:
.venv\scripts\activate

pip install django
django-admin startproject backchat
cd backchat
```

This will set up a base Django project. We can now work on the actual implementation of the Django project. Firstly, we will start with authentication.

## Adding basic Authentication and Authorization

### Creating the accounts app

We can separate the authentication of the user from the rest of the project, by creating a separate app called `user` or `accounts` app.

```bash
python manage.py startapp accounts
```

### Creating a base user model

We'll start by inheriting the [AbstractUser](https://docs.djangoproject.com/en/4.1/topics/auth/customizing/#using-a-custom-user-model-when-starting-a-project) the model provided in the `django.contrib.auth.models` module. The model has base fields like `username` and `password` which are required fields, and `email`, `first_name`, `last_name`, etc. which are not mandatory. It is better to create a custom model by inheriting the `AbstractUser` because if in the longer run, we need to make custom fields in the user model, it becomes a breeze.

```python
# accounts/models.py


from djnago.contrib.auth.models import AbstractUser


class User(AbstractUser):
    pass
```

This creates a basic custom user rather than using the Django built-in user. Next, we need to make sure, Django understands the default user is the one we defined in the `accounts` app and not the default `User`. So, we just need to add a field called `AUTH_USER_MODEL` in the `settings.py` file. The value of this field will be the app name or the module name and the model in that python module that we need our custom user model to be set as the default user model.

```python
# backchat/settings.py


INSALLED_APPS = [
    ...
    ...
    "accounts",
]

# Append to the end of file
AUTH_USER_MODEL = 'accounts.User'
```

### Initial migrations for the Django project

Now, this will get picked up as the default user model while referring to any processing related to the user. We can move into migrating the changes for the basic Django project and the user model.

```bash
python manage.py makemigrations
python manage.py migrate
```

![initial migration for base django and user model](https://meetgor-cdn.pages.dev/django-htmx/chat-accounts-migrations.png)

### Creating register view

Further, we can add the views like register and profile for the accounts app that can be used for the basic authentication. The Login and Logout views are provided in the `django.contrib.auth.views` module, we only have to write our own register view. I will be using function-based views to keep it simple to understand but it can simply be a class-based view as well.

To define a view first, we need to have form, a user registration form. The form will inherit from the [UserCreationForm](https://docs.djangoproject.com/en/4.1/topics/auth/default/#django.contrib.auth.forms.UserCreationForm) form which will make the bulk of the task for us at the user registration. We can simply then override the Meta class with the fields that we want to display, so hence we just keep the `username` and the `password` fields. The form can be customized by adding in the widget attribute and specifying the classes used in them.

```python
# accounts/forms.py


from accounts.models import User
from django.contrib.auth.forms import UserCreationForm

class UserRegisterForm(UserCreationForm):

    class Meta:
        model= User
        fields = ['username', 'password1', 'password2']
```

This will give us the `UserRegisterForm` form that can be used for displaying in the register view that we will create in the next step.

We will have to create the register view that will render the form for user registration and will also process the form submission.

```python
# accounts/views.py


from django.contrib import messages
from django.shortcuts import redirect, render
from accounts.forms import UserRegisterForm

def register(request):
    if request.method == "POST":
        form = UserRegisterForm(request.POST)
        if form.is_valid():
            form.save()
            username = form.cleaned_data.get("username")
            messages.success(
                request, f"Your account has been created! You are now able to log in"
            )
            return redirect("login")
    else:
        form = UserRegisterForm()
        return render(request, "accounts/register.html", {"form": form})
```

The above register view has two cases, one for the user requesting the registration form and the second request when the user submits the form. So, when the user makes a get request, we load an empty form `UserRegisterForm` and render the `register` template with the form. We will make the templates later.

So, the template is just rendered when the user wants to register and when the user submits the form(sends a post request) we get the details from the post request and make it an instance of `UserRegisterForm` and save the form if it is valid. We then redirect the user to the login view (will use the default one in the next section). We parse the message as the indication that the user was created.

### Adding URLs for Authentication and Authorization

Once, we have the register view setup, we can also add login and logout views in the app. But, we don't have to write it ourselves, we can override them if needed, but we'll keep the default ones. We will use the [LoginView](https://docs.djangoproject.com/en/4.1/topics/auth/default/#django.contrib.auth.views.LoginView) and the [LogoutView](https://docs.djangoproject.com/en/4.1/topics/auth/default/#django.contrib.auth.views.LogoutView) view which are class-based views provided in the `django.contrib.auth.views` module. We will provide the respective templates for each of these views.

```python
# accounts/urls.py


from django.urls import path
from django.contrib.auth import views as auth_views
import user.views as user_views

urlpatterns = [
    path("register/", user_views.register, name="register"),
    path(
        "login/",
        auth_views.LoginView.as_view(template_name="accounts/login.html"),
        name="login",
    ),
    path(
        "logout/",
        auth_views.LogoutView.as_view(template_name="accounts/logout.html"),
        name="logout",
    ),
]
```

We have named the URLs as `register`, `login`, and `logout` so that we can use them while rendering the links for them in the templates. Now, we also need to include the URLs from the accounts app in the project URLs. We can do that by using the `include` method and specifying the app name with the module where the urlpatterns are located.

```python
# backchat/urls.py


from django.contrib import admin
from django.urls import include, path

urlpatterns = [
    path("admin/", admin.site.urls),
    path("auth/", include("accounts.urls")),
]
```

So, we have routed the `/auth` path to include all the URLs in the accounts app. So, the login view will be at the URL `/auth/login/` , and so on.

Also, we need to add the `LOGIN_REDIRECT_URL` and `LOGIN_URL` for specifying the url name which will be redirected once the user is logged in and the default login url name respectively.

```python
# backchat / settings.py


LOGIN_REDIRECT_URL = "index"
LOGIN_URL = "login"
```

We are now almost done with the view and routing part of the accounts app and can move into the creation of templates.

### Adding Templates for authentication views

We need a few templates that we have been referencing in the views and the URLs of the accounts app in the project. So there are a couple of ways to use templates in a Django project. I prefer to have a single templates folder in the root of the project and have subfolders as the app which will have the templates specific to those apps.

I usually create a `base.html` file in the templates folder and use that for inheriting other templates. So, we will have to change one setting in the project to make sure it looks at `templates/` from the root of the project.

```bash
# backchat/settings.py

import os

...
...

TEMPLATES = [
    {
        "BACKEND": "django.template.backends.django.DjangoTemplates",
        "DIRS": [ os.path.join(BASE_DIR, "templates"), ],
        "APP_DIRS": True,
        "OPTIONS": {
            "context_processors": [
                "django.template.context_processors.debug",
                "django.template.context_processors.request",
                "django.contrib.auth.context_processors.auth",
                "django.contrib.messages.context_processors.messages",
            ],
        },
    },
]
```

Then create the folder in the same path as you see your `manage.py` file.

```bash
mkdir templates
```

![Template Set Up](https://meetgor-cdn.pages.dev/django-htmx/chat-mkdir-templates.png)

#### Creating the base template

The below will be the base template used for the chat application, you can custmize it as per your needs.

```html
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8" />
        <title>Chat App</title>
        {% load static %}
        <script src="https://unpkg.com/htmx.org@1.8.5"></script>
    </head>
    <body>
        {% if user.is_authenticated %}
            <a href="{% url 'logout' %}">Logout</a>
        {% else %}
            <a href="{% url 'login' %}">Login</a>
        {% endif %}
        <h1>Back Chat</h1>
        {% block base %}
        {% endblock %}
    </body>
</html>
```

We have included the [htmx](https://htmx.org/docs/#installing) package as the script into this template as we will be using it in the actual part of chat application.

#### Creating the Register Template

```html
# templates / accounts / register.html


{% extends 'base.html' %}
{% block base %}
    <div class="content-section">
        <form method="POST">
            {% csrf_token %}
            <fieldset class="form-group">
                <legend class="border-bottom mb-4">Register Now</legend>
                {{ form.as_p }}
            </fieldset>
            <div class="form-group">
                <button class="btn btn-outline-info" type="submit">Sign Up</button>
            </div>
        </form>
        <div class="border-top pt-3">
            <small class="text-muted">
		    Already Have An Account? <a class="ml-2" href="{% url 'login' %}">Log In</a>
            </small>
        </div>
    </div>
{% endblock %}
```

![User Registration Page](https://meetgor-cdn.pages.dev/django-htmx/chat-register-page.png)

#### Creating the Login Template

```html
# templates / accounts / login.html    


{% extends 'base.html' %}
{% block base %}
    <div class="content-section" id="login">
        <form method="POST" >
            {% csrf_token %}
            <fieldset class="form-group">
                <legend class="border-bottom mb-4">LOG IN</legend>
                {{ form.as_p }}
            </fieldset>
            <div class="form-group">
                <button class="btn btn-outline-info" type="submit">Log In</button>
            </div>
        </form>
        <div class="border-top pt-3">
            <small class="text-muted">
                Register Here <a class="ml-2" href="{% url 'register' %}">Sign Up</a>
            </small>
        </div>
    </div>
{% endblock %}
```

![User Login Page](https://meetgor-cdn.pages.dev/django-htmx/chat-login-page.png)

Creating the Logout Template

```html
# templates / accounts / logout.html    


{% extends 'base.html' %}
{% block base %}
    <h2>You have been logged out</h2>
    <div class="border-top pt-3">
        <small class="text-muted">
            <a href="{% url 'login' %}">Log In Again</a>
        </small>
    </div>
{% endblock %}
```

## Install and setup channels

We will be using channels to create long-running connections, it is a wrapper around Django's asynchronous components and allows us to incorporate other protocols like web sockets and other asynchronous protocols.

So, we will be using the Django channels package that will allow us to use the WebSocket protocol in the chat application. [WebSocket](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API) is a communication protocol(set of rules and standards to be followed) that allows both the client and server to send and receive messages(communication).

To install Django channels, we can use pip and install channels with daphne which will be used as the web server gateway interface for asynchronous web applications.

```bash
pip install -U channels["daphne"]
```

So this will install the latest version of channels and daphne. We now have [channels](https://channels.readthedocs.io/en/stable/installation.html) in our Django project, we simply need to configure the `CHANNEL_LAYER` config for specifying the backend used that can be `Redis`, `In-Memory`, or others. We need to add `channels` , `daphne` to the `INSALLED_APPS` config of the project. Make sure the `daphne` app is on top of the applications list.

```bash
# backchat/settings.py

...
...

INSALLED_APPS = [
    "daphne",
    ...
    ...
    "channels",
]


ASGI_APPLICATION = "backchat.asgi.application"

...
...

# For InMemory channels

CHANNEL_LAYERS = {
    'default': {
        "BACKEND": "channels.layers.InMemoryChannelLayer",
    }
}


# For Redis

CHANNEL_LAYERS = {
    "default": {
        "BACKEND": "asgi_redis.RedisChannelLayer",
        "CONFIG": {
            "hosts": [("redis-host-url", 6379)],
        },
    },
}
```

### Redis Configuration (Optional)

You can either use the [InMemoryChannelLayer](https://channels.readthedocs.io/en/latest/topics/channel_layers.html) or you can use them `RedisChannelLayer` for the backend of the chat app. There are other types of backends like `Amazon SQS` services, `RabbitMQ`, `Kafka`, `Google Cloud Pub/Sub`, etc. I will be creating the app with only the `InMemoryChannelLayer` but will provide a guide for redis as well, both are quite similar and only have a few nuances.

We need to install [channels_redis](https://github.com/django/channels_redis/) it for integrating redis in the Django project with channels.

```bash
pip install channels_redis
```

So, this will make the `channels_redis` package available in the project, we use this package for real-time storage, in the case of the chat app, we might use it for storing messages or room details, etc.

## Creating the Chat App

Further, we now can create another app for handling the rooms and chat application logic. This app will have its own models, views, and URLs. Also, we will define consumers and routers just like views and URLs but for asynchronous connections. More on that soon.

So, let's create the `chat` app.

```bash
python manage.py startapp chat
```

Then we will add the chat app to the `INSALLED_APPS` config.

```python
# backchat/settings.py

INSALLED_APPS = [
    ...
    ...,
    "chat",
]
```

So, this will make sure to load the chat app in the project. Whenever we run any commands like migrations or running the server, it keeps the app in the `INSALLED_APPS` checked up.

### Defining models

This is optional, but we'll do it for since we are making a Django app. We already have an auth system configured, adding rooms and authorizing the users will become easier then.

So, let's create the models for the chat app which will be the `Room`.

```python
# chat/models.py


from django.db import models
from accounts.models import User

class Room(models.Model):
    name = models.CharField(max_length=128)
    slug = models.SlugField(unique=True)
    users = models.ManyToManyField(User)

    def __str__(self):
        return self.name


class Message(models.Model):
    room = models.ForeignKey(Room, on_delete=models.CASCADE)
    user = models.ForeignKey(User, on_delete=models.CASCADE)
    message = models.TextField()
    created_at = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return (
            self.room.name + " - " +
            str(self.user.username) + " : " +
            str(self.message)
        )
```

So, we simply have the name which will be taken from the user, and the slug which will be served as the URL to the room. The users are set as [ManyToManyField](https://docs.djangoproject.com/en/4.1/ref/models/fields/#django.db.models.ManyToManyField) since one room can have multiple users and a user can be in multiple rooms. Also we are creating the model `Message` that will be storing the room and the user as the foreign key and the actual text as the message, we can improve the security by encrypting the text but it's not the point of this article.

We have set the `created_at` the field which will set the time when the object was created. Finally, the dunder string method is used for representing the message object as a price of the concatenation of strings of room name, username, and the message. This is useful for admin stuff as it makes it easier to read the object rather than the default id.

Now, once the models are designed we can migrate the models into the database.
```
python manage.py makemigrations
python manage.py migrate
```

![Chat app migrations](https://meetgor-cdn.pages.dev/django-htmx/chat-migrations.png)

We now have a decent model structure for the chat application. We can now start the crux of the application i.e. consumers and routing with channels.

### Writing consumers and routers for WebSockets

So, we start with the bare bones provided in the tutorial on the channel [documentation](https://channels.readthedocs.io/en/stable/tutorial/part_3.html). We create a class (consumer) called `ChatConsumer` which inherits from the `AsyncWebsocketConsumer` provided by the `channels.generic.websocket` module. This has a few methods like `connect`, `disconnect`, and `receive`. These are the building blocks of a consumer. We define these methods as they will be used for communication via the WebSocket protocol through the channels interface.

In the following block of code, we are essentially doing the following:

* Accepting connection on the requested room name
    
* Sending and Receiving messages on the room/group
    
* Closing the WebSocket connection and removing the client from the room/group
    

```python
# chat/consumers.py


import json

from asgiref.sync import sync_to_async
from channels.generic.websocket import AsyncWebsocketConsumer

from chat.models import Room, Message


class ChatConsumer(AsyncWebsocketConsumer):
    async def connect(self):
        self.room_name = self.scope["url_route"]["kwargs"]["room_slug"]
        self.room_group_name = f"chat_{self.room_name}"
        self.user = self.scope["user"]

        await self.channel_layer.group_add(
            self.room_group_name, self.channel_name
        )

        await self.accept()

    async def disconnect(self, close_code):
        await self.channel_layer.group_discard(
            self.room_group_name, self.channel_name
        )

    async def receive(self, text_data):
        text_data_json = json.loads(text_data)
        message = text_data_json["message"]
        username = self.user.username
        
        await self.channel_layer.group_send(
            self.room_group_name, 
            {
                "type": "chat_message",
                "message": message,
                "username": username,
            }
        )

    async def chat_message(self, event):
        message = event["message"]
        username = event["username"]

        await self.send(
            text_data=json.dumps(
                {
                    "message": message,
                    "username": username
                }
            )
        )
```

#### Accept the WebSocket connection

Here, room and group more or less mean the same thing but are different in different contexts. For instance, the group refers to the collection of clients which are connected to a channel(communication between consumer and web socket) and the Room is referring to the collection of clients connected to a single connection stream like a literal room. So we can say, the group is a technical term and the room is more of a layman's term for the same thing.

The method `connect` is called when the client establishes a websocket connection. When that happens, the function gets the room slug from the URL of the client and stores `room_name` which is a string. It creates a separate variable called `room_group_name` by prepending the `chat_` string to the `room_name`, it also gets the currently logged-in user from the request. It then adds the `channel_name` to the group named `room_group_name`. The `channel_name` is a unique identifier to the connection/consumer in the channel. By adding the `channel_name`, the consumer then can broadcast the message to all the channels within the group. Finally, the function accepts the connection, and a **webcoket connection is established from both ends, connection is sent from the client and is now accepted from the backend.**

#### Disconnect from the WebSocket connection

When the client sends a close connection request to the server, the `disconnect` method is triggered and it basically removes the `channel_name` from the group i.e. the group name `room_group_name` whatever the string it happens to be. So, it basically removes the client from the broadcast group and hence it can't receive or send messages through the websocket since it has been closed from both ends.

If you would have paid attention to the `close_code` in-method parameter, it is not being used currently. However, we can use it for checking why the connection was closed, as the `close_code` is a numeric value just like the status code in the web request for letting the server know the reason for disconnecting from the client.

#### Receive a message from the WebSocket connection

The `recieve` method is the core of the consumer as it is responsible for all the logic and parsing of the message and broadcasting of the messages from the clients to the group channels. The function takes in a parameter called `text_data` and it is sent from the client through websocket, so it is JSON content. We need to get the actual message from the JSON object or any other piece of content from the client. So, we deserialize (convert the JSON object to python objects) the received payload, and get the value from the key `message`. The key is the input name or id from the client sending the request through the web socket, so it can be different depending on the frontend template(we'll see the front end soon as well).

We get the user from the scope of the consumer as we previously initialized it in the connect method. This can be used for understanding which user has sent the message, it will be used later on as we send/broadcast the message to the group.

Now, the final piece in the receive method is the `channel_layer.group_send` method, this method as the name suggests is used to send or broadcast the received message to the entire group. The method has two parameters:

1. The name of the group
    
2. The JSON body containing the message and other details
    

The method is not directly sending the message but it has a type key in the JSON body which will be the function name to call. The function will simply call the other funciton mentioned in the type key in the dict. The following keys in the dict will be the parameters of that funciton. In this case, the funciton specified in the `type` key is `chat_message` which takes in the `event` as the parameter. This event will have all the parameters from the `group_send` method.

So, the `chat_message` will load in this message, username, and the room name and then call the `send` method which actually sends the message as a JSON payload to the WebSocket connection which will be received by all the clients in the same group as provided in the `room_group_name` string.

### Adding Routers for WebSocket connections

So, till this point have consumers, which are just like views in terms of channels. Now, we need some URL routes to map these consumers to a path. So, we will create a file/module called `routing.py` which will look quite similar to the `urls.py` file. It will have a list called `websocket_urlpatterns` just like `urlpatterns` with the list of `path`. These paths however are not `http` routes but will serve for the `WebSocket` path.


```python
# chat / routing.py


from django.urls import path

from chat import consumers

websocket_urlpatterns = [
    path('chat/<str:room_slug>/', consumers.ChatConsumer.as_asgi()),
]
```

In the above code block, we have defined a URL for the web socket with the path as `/chat/<room_slug>` where room_name will be the `slug` for the room. The path is bound with the consumer-defined in the `consumers.py` module as `ChatConsumer`. The `as_asgi` method is used for converting a view into an ASGI-compatible view for the WebSocket interface.

### Setting up ASGI Application

We are using the ASGI application config rather than a typical WSGI application which only works one request at a time. We want the chat application to be asynchronous because multiple clients might send and receive the messages at the same time, we don't want to make a client wait before the server process a message from another client, that's just the reason we are using WebSocket protocol.

So, we need to also make sure, it makes the http request and also add our websocket config from the chat app we created in the previous sections. So, inside the `asgi.py` file in the project config module, we need to make some changes to include the chat application configurations.

```python
# backchat / asgi.py


import os
from django.core.asgi import get_asgi_application
from channels.auth import AuthMiddlewareStack
from channels.routing import ProtocolTypeRouter, URLRouter

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'backchat.settings')

from chat import routing


application = ProtocolTypeRouter({
    'http': get_asgi_application(),
    "websocket":AuthMiddlewareStack(
        URLRouter(
            routing.websocket_urlpatterns
        )
    )
})
```

We will override the `application` config which is a component used for routing different types of protocols for the `ASGI` application. We have set the two keys, `http` and `websocket` in our application. The `http` type of requests will be served with the `get_asgi_application` application which is used for running the application in an ASGI environment.

For `websocket` type of requests, we are setting the [AuthMiddlewareStack](https://channels.readthedocs.io/en/latest/topics/authentication.html) which helps in authenticating the users requesting the WebSocket connection and allows only authorized users to make a connection to the application. The [URLRouter](https://channels.readthedocs.io/en/stable/topics/routing.html) is used for mapping the list of URL patterns with the incoming request. So, this basically serves the request URL with the appropriate consumer in the application. We are parsing in the `websocket_urlpatterns` as the list of URLs that will be used for the WebSocket connections.

Now, we run the server, we should be seeing the `ASGI` server serving our application rather than the plain WSGI application.

```
$ python manage.py runserver

Watching for file changes with StatReloader
Performing system checks...

System check identified no issues (0 silenced).
February 05, 2023 - 09:22:45
Django version 4.1.5, using settings 'backchat.settings'
Starting ASGI/Daphne version 4.0.0 development server at http://127.0.0.1:8000/
Quit the server with CONTROL-C.
```


The application is not complete yet, it might not have most components working functional yet. So, we'll now get into making the user interfaces for the application, to create, join, and view rooms in the application.

### Adding Views for Chat Rooms

We will have a couple of views like create room page, the join room page, and the chat room page. We will define each view as a distinct view and all of them will require authenticated users.

```python
# chat / views.py


import string
import random
from django.contrib.auth.decorators import login_required
from django.shortcuts import render, reverse, redirect
from django.utils.text import slugify
from chat.models import Room


@login_required
def index(request, slug):
    room = Room.objects.get(slug=slug)
    return render(request, 'chat/room.html', {'name': room.name, 'slug': room.slug})

@login_required
def room_create(request):
    if request.method == "POST":
        room_name = request.POST["room_name"]
        uid = str(''.join(random.choices(string.ascii_letters + string.digits, k=4)))
        room_slug = slugify(room_name + "_" + uid)
        room = Room.objects.create(name=room_name, slug=room_slug)
        return redirect(reverse('chat', kwargs={'slug': room.slug}))
    else:
        return render(request, 'chat/create.html')

@login_required
def room_join(request):
    if request.method == "POST":
        room_slug = request.POST["room_slug"]
        room = Room.objects.get(slug=room_slug)
        return redirect(reverse('chat', kwargs={'slug': room.slug}))
    else:
        return render(request, 'chat/join.html')
```

In the above views module, we have added 3 views namely `index` as the room page, `room_create` for the room creation page, and the `room_join` for the room join page. The index view is a simple get request to the provided slug of the room, it gets the slug from the URL from the request and fetches an object of the room associated with that slug. Then it renders the room template with the context variables like the name of the room and the slug associated with that room.

The `room_create` view is a simple two-case view that either can render the room creation page or process the submitted form and create the room. Just like we used in the `register` view in the accounts app. When the user will send a `GET` request to the URL which we will map to `/create/` shortly after this, the user will be given a form. So, we will render the `create.html` template. We will create the html template shortly. 
If the user has sent a `POST` request to the view via the `/create` URL, we will fetch the name field in the sent request and create a unique identifier with the name of the room. We will slugify the concatenation of the name with the uid and will set it as the slug of the room. We will then simply create the room and redirect the user to the `room` page.

The `room_join` view also is a two-case view, where the user can either request the join room form or send a slug with the form submission. If the user is requesting a form, we will render the `join.html` template. If the user is submitting the form, we will fetch the room based on the slug provided and redirect the user to the `room` page.

So, the `room_join` and `room_create` views are quite similar, we are fetching an already existing room in the case of the join view and creating a new instance of room in the create view. Now, we will connect the views to the URLs and finally get to the templates.

### Connecting Views and URLs

We have three views to route to a URL. But we will also have one additional URL which will be the home page for the application, on that page the user can choose to either join or create a room. We have the room creation, join the room and the room view to be mapped in this URL routing of the app.

```python
# chat / urls.py


from django.urls import path
from django.views.generic import TemplateView
from chat import views


urlpatterns = [
    path("", TemplateView.as_view(template_name="base.html"), name='index'),
    path("room/<str:slug>/", views.index, name='chat'),
    path("create/", views.room_create, name='room-create'),
    path("join/", views.room_join, name='room-join'),
]
```

So, the first route is the home page view called `index`, we have used the [TemplateView](https://docs.djangoproject.com/en/4.1/ref/class-based-views/base/#templateview) which will simply render the template provided. We don't have to create a separate view for just rendering a template. We already have defined the `base.html` template while setting up the `accounts` app. This will be the same template, we will add some more content to the template later on. The URL mapped is the `/` since we will add the URLs of this app to the project URLs with an empty `""` path.

The second route is used for the room index page, i.e. where the user will be able to send and receive messages. The path is assigned as `/room/<str:slug>/` indicating a parameter called slug of type string will be used in accessing a particular room. The URL will be bound to the index view that we created in the views file, which fetches the room based on the slug, so here is where the slug will be coming from. The name of the URL-View route will be `chat` but you can keep it as per your requirements. The URL name is really handy for use in the templates.

The third route is for the room creation page. The `/create/` URL will be bound to the `room_create` view, as we discussed, it will serve two purposes, one to render the form for creating the room, and the other for sending a `POST` request to the same path for the creating the Room with the name provided. The name is not required but helps in identifying and making it user-friendly. The name of this URL is set as `room-create`.

The final route is for joining the room, the `/join/` URL will be triggering the `room_join` view. Similar to the `room-create` URL, the URL will render the join room form on a `GET` request, fetch the room with the provided slug and redirect to the room page. Here, the slug field in the form will be required for actually finding the matching room. The name of the URL route is set as `room-join`.

We now add the URLs from the chat app to the project URLs. This will make the `/` as the entry point for the chat application URLs.

```python
# backchat / urls.py


from django.contrib import admin
from django.urls import path, include

urlpatterns = [
    path("admin/", admin.site.urls),
    path("auth/", include('accounts.urls')),
    path("", include('chat.urls')),
]
```

Hence the process is completed for the backend to process the message, it then is dependent on the client to process and render the message.

Till HTMX was not a thing!

We won't have to write a single line of javascript to receive and handle the WebSocket connection!

### Creating Templates and adding htmx

We now move into the actual frontend or creating the template for actually working with the rooms and user interaction. We will have three pieces of templates, a room creates the page, a room join page, and a room chat page. As these template names suggest, they will be used for creating a room with the name, joining the room with the slug, and the room chat page where the user will send and receive messages.

Let/s modify the base template first.

```html
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8" />
        <title>Chat App</title>
        {% load static %}
        <script src="https://unpkg.com/htmx.org@1.8.5"></script>
    </head>
    <a href="{% url 'index' %}">Home</a>
        {% if user.is_authenticated %}
            <a href="{% url 'logout' %}">Logout</a>
        {% else %}
            <a href="{% url 'login' %}">Login</a>
        {% endif %}
    <body hx-ext="ws">
        <h1>Back Chat</h1>
        {% block base %}
            <a href="{% url 'room-join' %}">Join Room</a>
            <a href="{% url 'room-create' %}">Create Room</a>
        {% endblock %}
    </body>
</html>
```

![Chat App Home Page](https://meetgor-cdn.pages.dev/django-htmx/chat-home-page.png)

#### Create Room Template

We will have to render the form with a field like `name` for setting it as the name of the room, it is not required but again, it makes it easier for the user to find the name of the room a bit more friendly than random characters.

```html
# templates / chat / create.html


{% extends 'base.html' %}

{% block base %}
    <form method='post' action='{% url 'room-create' %}'>
        {% csrf_token %}
        <input name='room_name' id='room_name' placeholder='Room Name'>
        <input type='submit' id="submit">
    </form>
{% endblock %}
```

![Chat Room Create Page](https://meetgor-cdn.pages.dev/django-htmx/chat-create-room-page.png)

The template is inherited from the `base.html` template and we render a form with the `room_name` input. The form will be submitted to the URL named `room-create` hence it will send a `POST` request to the server where it will create the room and further process the request.

#### Join Room Template

The join room template is similar to the create room template except it gets the slug of the room from the user rather than the name which is not a unique one to join the room.

```html
# templates / chat / join.html


{% extends 'base.html' %}

{% block base %}
    <form method='post' action='{% url 'room-join' %}'>
        {% csrf_token %}
        <input name='room_slug' id='room_slug' required='true' placeholder='Room Code'>
        <input type='submit' id="submit">
    </form>
{% endblock %}
```

![Chat Room Join Page](https://meetgor-cdn.pages.dev/django-htmx/chat-join-room-page.png)

The form will be submitted to the URL named `room-join` hence it will send a `POST` request to the server where it will fetch the room and further process the request.

### Room Template (HTMX code)

Now, time for the actual ingredient in the application, some HTMX magic!

This template, as the two templates above inherit from the base template, that's the same thing. But it has a special `div` with the attribute [hx-ws](https://htmx.org/attributes/hx-ws/) which is used for using attributes related to the web socket in the htmx library. The `connect` value is used for connecting to a WebSocket. The value of the attribute must be set to the URL of the WebSocket to be connected. In our case, it is the URL path from the `routing` app as `/chat/<room_slug>/`. This simply will connect the client to the WebSocket from the backend. The other important attribute is the `send` which is used for sending a message to the connected web socket.

```html
# templates / chat / room.html


{% extends 'base.html' %}

{% block base %}
    <h2>{{ name }}</h2>
    <div hx-ws="connect:/chat/{{ slug }}/">
        <form hx-ws="send:submit">
            <input name="message">
            <input type="submit">
        </form>
     </div>
     <div id='messages'></div>
{% endblock %}
```

![Chat Room Page](https://meetgor-cdn.pages.dev/django-htmx/chat-room-page.png)

NOTE: The template has a div with the id `messages` which will be very important for sending the messages from the WebSocket to the client, so more on that when we use the HTMX part.

For testing this template, you can create a room, and that will redirect you to the room template as we have seen in the views for the room creation. If you see something like `WebSocket CONNECT` it means, that the application has been able to establish a WebSocket connection to the backend, and we can be ready to accept messages and other stuff.

```
HTTP GET /chat/room/def_teas/ 200 [0.03, 127.0.0.1:38660]
WebSocket HANDSHAKING /chat/def_teas/ [127.0.0.1:38666]
WebSocket CONNECT /chat/def_teas/ [127.0.0.1:38666]
```

![Django ASGI server websocket connection](https://meetgor-cdn.pages.dev/django-htmx/chat-asgi-server.png)

Till this point, we should have a running and almost complete application, though we just have a minor part missing that will be the most important part.

### Sending HTML response from backend for htmx

We will be sending a fragment of HTML from the backend when the user sends a message, to broadcast it to the group. Let's make some changes to the application, especially to the receive method in the `ChatConsumer` of the chat application.

```python
# chat / consumers.py
    

    ...
    ...

    async def receive(self, text_data):
        text_data_json = json.loads(text_data)
        message = text_data_json["message"]
        user = self.user
        username = user.username

        await self.channel_layer.group_send(
            self.room_group_name, 
            {
                "type": "chat_message",
                "message": message,
                "username": username,
            }
        )

    async def chat_message(self, event):
        message = event["message"]
        username = event["username"]

        # This is the crucial part of the application
        message_html = f"<div hx-swap-oob='beforeend:#messages'><p><b>{username}</b>: {message}</p></div>"
        await self.send(
            text_data=json.dumps(
                {
                    "message": message_html,
                    "username": username
                }
            )
        )
```

![Chat Room Message](https://meetgor-cdn.pages.dev/django-htmx/chat-room-msg-1.png)
![Chat Room Message 2 Users](https://meetgor-cdn.pages.dev/django-htmx/chat-room-msg-2.png)
![Chat Room Message](https://meetgor-cdn.pages.dev/django-htmx/chat-room-msg-3.png)


In the above snippet, we are just changing the final message object to include some HTML just simple. The HTML however has home htmx attributes like [hx-swap-oob](https://htmx.org/attributes/hx-swap-oob/) which will just update the specified DOM element to the content in the div. In this case, the DOM element is `#message` which is the id message present in the room template. We basically add the username and the message into the same id by appending it to the element. That's it, it would work and it would start showing the messages from the connected clients and broadcast them as well.

There are some things to keep in mind while using htmx in the long run especially when the htmx 2.0 is released, it will have `ws` as a separate extension. It will have a bit of a different syntax than above. I have tried the latest version but doesn't seem to work. I'll just leave a few snippets for your understanding of the problem.


```html
# templates / chat / room.html


{% extends 'base.html' %}

{% block base %}
    <h2>{{ name }}</h2>
    <div hx-ext="ws" ws-connect="/chat/{{ slug }}/">
        <form ws-send>
            <input name="message">
        </form>
    </div>
    <div id='messages'></div>
{% endblock %}
```

I have added, the `hx-ext` as `ws` which is a htmx [extension for websocket](https://htmx.org/extensions/web-sockets/). This extension has websocket-specific attributes like `ws-connect` and `ws-send`. I have tried a combination like changing the htmx versions, adding submit value as the `ws-send` attribute, etc, but no results yet. I have opened a [discussion](https://github.com/bigskysoftware/htmx/discussions/1231) on GitHub for this issue, you can express your solution or views there.

### Adding some utility features for the chat app

We can save messages, add and remove the users from the room according to the connection, and other stuff that can make this a fully-fledged app. So, I have made a few changes to the chat consumers for saving the messages and also updating the room with the users in the room.

```python
# chat / consumers.py


import json

from asgiref.sync import sync_to_async
from channels.generic.websocket import AsyncWebsocketConsumer

from chat.models import Room, Message


class ChatConsumer(AsyncWebsocketConsumer):
    async def connect(self):
        self.room_name = self.scope["url_route"]["kwargs"]["room_slug"]
        self.room_group_name = "chat_%s" % self.room_name
        self.user = self.scope["user"]

        await self.channel_layer.group_add(
            self.room_group_name, self.channel_name
        )

        # Add the user when the client connects
        await self.add_user(self.room_name, self.user)

        await self.accept()

    async def disconnect(self, close_code):

        # Remove the user when the client disconnects
        await self.remove_user(self.room_name, self.user)

        await self.channel_layer.group_discard(
            self.room_group_name, self.channel_name
        )

    async def receive(self, text_data):
        text_data_json = json.loads(text_data)
        message = text_data_json["message"]
        user = self.user
        username = user.username
        room = self.room_name

        # Save the message on recieving
        await self.save_message(room, user, message)

        await self.channel_layer.group_send(
            self.room_group_name, 
            {
                "type": "chat_message",
                "message": message,
                "username": username,
            }
        )

    async def chat_message(self, event):
        message = event["message"]
        username = event["username"]


        message_html = f"<div hx-swap-oob='beforeend:#messages'><p><b>{username}</b>: {message}</p></div>"
        await self.send(
            text_data=json.dumps(
                {
                    "message": message_html,
                    "username": username
                }
            )
        )

    @sync_to_async
    def save_message(self, room, user, message):
        room = Room.objects.get(slug=room)
        Message.objects.create(room=room, user=user, message=message)

    @sync_to_async
    def add_user(self, room, user):
        room = Room.objects.get(slug=room)
        if user not in room.users.all():
            room.users.add(user)
            room.save()

    @sync_to_async
    def remove_user(self, room, user):
        room = Room.objects.get(slug=room)
        if user in room.users.all():
            room.users.remove(user)
            room.save()
```

So, we have created a few methods like `save_message`, `add_user`, and `remove_user` which all are `synchronous` methods but we are using an asynchronous web server, so we add in the `sync_to_async` decorator which wraps a synchronous method to an asynchronous method. Inside the methods, we simply perform the database operations like creating a message object, and adding or removing the user from the room.

That's only a few features that I have added, you can add to this application and customize them as per your needs.

The code for this chat app is provided in the [GitHub repository](https://github.com/Mr-Destructive/django-htmx-chat).

## Conclusion

So, from this post, we were able to create a simple chat app (frontendless) with Django and htmx. We used Django channels and HTMX to make a chat application without the need to write javascript for the client-side connection. Hope you found this tutorial helpful, do give your feedback and thoughts on it, I'll be eager to improve this post. Thank you for your patient listening. Happy Coding :)