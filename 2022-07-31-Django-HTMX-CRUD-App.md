---
article_html: "<h2>Introduction</h2>\n<p>Gone are the days of writing Ajax requests
  with javascript, just add a few parameters to the HTML content tags and you will
  be ready for sending requests to your backend. So, we are going back in time and
  correcting the way we think of APIs and client/server-side rendering. We are referring
  to the <a href=\"https://en.wikipedia.org/wiki/Hypermedia\">Hypermedia model</a>
  for levering the server-side processing of data. Let's get our feets wet with this
  ancient but revolutionary methodology of development with <a href=\"https://htmx.org/\">HTMX</a>.</p>\n<p>Yes,
  HTMX can be used for the API/server-side calls directly in the HTML. We will be
  exploring the basis of HTMX by creating a basic CRUD application.</p>\n<h2>What
  is HTMX?</h2>\n<p>The first question that might come up is what and why HTMX? Htmx
  is a great library, it's a javascript library but wait. It is a javascript library
  designed to allow us to write less or no javascript at all. It acts as a way to
  send AJAX requests without you writing any javascript. It uses native browser features
  directly from HTML.</p>\n<p>So, we can use HTMX to create interactive templates
  in our Django application. We can dynamically call and fetch data from the server
  by using simple HTML attributes like <code>hx-get</code>, <code>hx-post</code>,
  etc. We'll cover those in this article.</p>\n<p>You can check the source code used
  in this article on this <a href=\"https://github.com/Mr-Destructive/htmx-blog-django\">GitHub
  repository</a>.</p>\n<h2>Setup Django Project</h2>\n<p>We'll be creating a Django
  project from scratch and designing a basic blog kind of app. We will be creating
  a quite simple project with a couple of apps like <code>user</code> for authentication
  and <code>article</code> for the CRUD part of our blog application.</p>\n<p>To set
  up a django project, we can run the following commands to quickly get up and running
  with a base django project.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>mkdir htmx_blog\npython3
  -m venv .venv\nsource .venv/bin/activate\npip install django\ndjango-admin startproject
  htmx_blog .\n</pre></div>\n\n</pre>\n\n<p>I have a base user model that I use for
  a simple authentication system in some basic django projects, you can define your
  own user app or get the app from <a href=\"https://github.com/Mr-Destructive/django-todo/tree/master/user\">here</a>.</p>\n<p>So,
  that being said, we will be using the user model for the article model which we
  will be defined next. By creating a basic signup functionality, you are good to
  go!</p>\n<h3>Create the Article app</h3>\n<p>We will need at least an app to work
  with htmx as we will define models, views, and URLs later as we configure the htmx.</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>django-admin startapp article\n</pre></div>\n\n</pre>\n\n<p>After
  the app has been created, you can add those app labels into the <code>INSTALLED_APPS</code>
  config in the <code>settings.py</code> file. The <code>user</code> app and the <code>article</code>
  app need to be added to the installed apps for the django to pick those up for various
  contexts related to the project.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span># htmx_blog/settings.py\n\nINSTALLED_APPS
  = [\n    ...\n    ...\n    ...\n\n    &#39;article&#39;,  \n    &#39;user&#39;,\n]\n</pre></div>\n\n</pre>\n\n<p>We
  are sone with the base setup, we also would require a few more configs for the proper
  working of the project.</p>\n<h3>Setup Templates and Static files</h3>\n<p>Templates
  will play an important role in the htmx part, so it is equally important to configure
  them properly before dabbling into the htmx and client-side rendering of data.</p>\n<p>I
  like to keep all the templates in a single folder in the <code>BASE_DIR</code> with
  separate sub-folders for specific apps. Also a single <code>static</code> folder
  with <code>css</code>, <code>js</code>, and <code>images</code> as the sub-folfers
  for a larger project.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>mkdir templates static\n</pre></div>\n\n</pre>\n\n<p>Further,
  configure the created static and templates in the settings.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"n\">TEMPLATES</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"p\">{</span>\n        <span class=\"s1\">&#39;BACKEND&#39;</span><span
  class=\"p\">:</span> <span class=\"s1\">&#39;django.template.backends.django.DjangoTemplates&#39;</span><span
  class=\"p\">,</span>\n        <span class=\"s1\">&#39;DIRS&#39;</span><span class=\"p\">:</span>
  <span class=\"p\">[</span><span class=\"n\">os</span><span class=\"o\">.</span><span
  class=\"n\">path</span><span class=\"o\">.</span><span class=\"n\">join</span><span
  class=\"p\">(</span><span class=\"n\">BASE_DIR</span><span class=\"p\">,</span>
  <span class=\"s2\">&quot;templates&quot;</span><span class=\"p\">)],</span>\n        <span
  class=\"s1\">&#39;APP_DIRS&#39;</span><span class=\"p\">:</span> <span class=\"kc\">True</span><span
  class=\"p\">,</span>\n        <span class=\"s1\">&#39;OPTIONS&#39;</span><span class=\"p\">:</span>
  <span class=\"p\">{</span>\n            <span class=\"s1\">&#39;context_processors&#39;</span><span
  class=\"p\">:</span> <span class=\"p\">[</span>\n                <span class=\"s1\">&#39;django.template.context_processors.debug&#39;</span><span
  class=\"p\">,</span>\n                <span class=\"s1\">&#39;django.template.context_processors.request&#39;</span><span
  class=\"p\">,</span>\n                <span class=\"s1\">&#39;django.contrib.auth.context_processors.auth&#39;</span><span
  class=\"p\">,</span>\n                <span class=\"s1\">&#39;django.contrib.messages.context_processors.messages&#39;</span><span
  class=\"p\">,</span>\n            <span class=\"p\">],</span>\n        <span class=\"p\">},</span>\n
  \   <span class=\"p\">},</span>\n<span class=\"p\">]</span>\n\n<span class=\"n\">STATIC_URL</span>
  <span class=\"o\">=</span> <span class=\"s1\">&#39;static/&#39;</span>\n<span class=\"n\">STATICFILES_DIRS</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span><span class=\"nb\">str</span><span
  class=\"p\">(</span><span class=\"n\">BASE_DIR</span><span class=\"o\">/</span>
  <span class=\"s2\">&quot;static&quot;</span><span class=\"p\">)]</span>\n<span class=\"n\">STATIC_ROOT</span>
  <span class=\"o\">=</span> <span class=\"n\">BASE_DIR</span> <span class=\"o\">/</span>
  <span class=\"s2\">&quot;staticfiles&quot;</span>\n</pre></div>\n\n</pre>\n\n<h3>Initial
  migration</h3>\n<p>Run migration command for the user model and default model in
  the django project.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  manage.py migrate\n</pre></div>\n\n</pre>\n\n<p>So, this project will also include
  authentication simple registration, and login/logout routes. We will be using the
  default Django User model by creating an abstract user just in case we require any
  additional attributes.</p>\n<h2>Setup HTMX</h2>\n<p>We don't have to configure much
  for using HTMX as it is a javascript library, we can call it via a CDN or manually
  install it and link up the static javascript files. Either way, both are equally
  good, you may like the one I might like the other.</p>\n<p>If you already have a
  base template, you can simply put the below script inside the head tag of the template.
  This will make us the htmx attributes available.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">&lt;</span><span class=\"nt\">script</span> <span class=\"na\">src</span><span
  class=\"o\">=</span><span class=\"s\">&quot;https://unpkg.com/htmx.org@1.8.0&quot;</span><span
  class=\"p\">&gt;&lt;/</span><span class=\"nt\">script</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>If
  you don't have a base template, you can create one by creating an HTML file inside
  the <code>templates</code> directory. The name can be anything but be careful for
  following up as it might be different for me. I will choose <code>base.html</code>
  as the template for this project. It will look something like as follows:</p>\n<pre
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cm\">&lt;!-- tempaltes/base.html --&gt;</span>\n\n<span class=\"cp\">&lt;!DOCTYPE
  html&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">html</span> <span
  class=\"na\">lang</span><span class=\"o\">=</span><span class=\"s\">&quot;en&quot;</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">meta</span>
  <span class=\"na\">charset</span><span class=\"o\">=</span><span class=\"s\">&quot;UTF-8&quot;</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>HTMX Blog<span class=\"p\">&lt;/</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>\n    {% load static %}\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">link</span> <span class=\"na\">rel</span><span class=\"o\">=</span><span
  class=\"s\">&quot;stylesheet&quot;</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css&quot;</span>
  <span class=\"na\">integrity</span><span class=\"o\">=</span><span class=\"s\">&quot;sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm&quot;</span>
  <span class=\"na\">crossorigin</span><span class=\"o\">=</span><span class=\"s\">&quot;anonymous&quot;</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">script</span>
  <span class=\"na\">src</span><span class=\"o\">=</span><span class=\"s\">&quot;https://unpkg.com/htmx.org@1.8.0&quot;</span><span
  class=\"p\">&gt;&lt;/</span><span class=\"nt\">script</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;/</span><span class=\"nt\">head</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;</span><span class=\"nt\">body</span> <span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;</span><span class=\"nt\">nav</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>HTMX
  Blog<span class=\"p\">&lt;/</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;</span><span class=\"nt\">div</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;navbar&quot;</span><span class=\"p\">&gt;</span>\n
  \         {% if user.is_authenticated %}\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;nav-item nav-link&quot;</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% url &#39;logout&#39; %}&quot;</span><span
  class=\"p\">&gt;&lt;</span><span class=\"nt\">button</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;btn btn-link&quot;</span><span class=\"p\">&gt;</span>Logout<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n          {% else %}\n            <span
  class=\"p\">&lt;</span><span class=\"nt\">a</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;nav-item nav-link&quot;</span> <span
  class=\"na\">href</span><span class=\"o\">=</span><span class=\"s\">&quot;{% url
  &#39;login&#39; %}&quot;</span><span class=\"p\">&gt;&lt;</span><span class=\"nt\">button</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;btn
  btn-link&quot;</span><span class=\"p\">&gt;</span>Login<span class=\"p\">&lt;/</span><span
  class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span class=\"nt\">a</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;nav-item
  nav-link&quot;</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% url &#39;register&#39; %}&quot;</span><span class=\"p\">&gt;&lt;</span><span
  class=\"nt\">button</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;btn btn-link&quot;</span><span class=\"p\">&gt;</span>Register<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n          {% endif %}\n        <span
  class=\"p\">&lt;/</span><span class=\"nt\">div</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;/</span><span class=\"nt\">nav</span><span class=\"p\">&gt;</span>\n\n
  \   {% block body %}\n    {% endblock %}\n<span class=\"p\">&lt;/</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>I have a nav bar with my user
  authentication views, simply a login or signup button if the user is not logged
  in and a log-out button if the user is authenticated. We have added the htmx script
  file from the CDN just before the end of the head tag. We also have included the
  bootstrap CSS file for a decent UI which we will be creating in this post.</p>\n<p>That
  is one of the ways, htmx can be injected into an HTML template, you can even download
  the javascript file from the <a href=\"https://unpkg.com/browse/htmx.org/dist/\">htmx
  cdn</a>. Further, this can be downloaded or pasted into your local folder and served
  as a static file or embedded directly into an HTML template.</p>\n<h2>Defining Models</h2>\n<p>We
  will start the tutorial by defining the model of the application we are creating.
  Here, we will create a simple Article model with a few parameters like <code>title</code>,
  <code>content</code>, <code>author</code>, etc.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">from</span> <span class=\"nn\">django.db</span> <span class=\"kn\">import</span>
  <span class=\"n\">models</span>\n<span class=\"kn\">from</span> <span class=\"nn\">user.models</span>
  <span class=\"kn\">import</span> <span class=\"n\">Profile</span>\n\n<span class=\"k\">class</span>
  <span class=\"nc\">Article</span><span class=\"p\">(</span><span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">Model</span><span class=\"p\">):</span>\n
  \   <span class=\"n\">Article_Status</span> <span class=\"o\">=</span> <span class=\"p\">(</span>\n
  \       <span class=\"p\">(</span><span class=\"s2\">&quot;DRAFT&quot;</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;Draft&quot;</span><span class=\"p\">),</span>\n
  \       <span class=\"p\">(</span><span class=\"s2\">&quot;PUBLISHED&quot;</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;Published&quot;</span><span class=\"p\">),</span>\n
  \   <span class=\"p\">)</span>\n    <span class=\"n\">title</span> <span class=\"o\">=</span>
  <span class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">CharField</span><span
  class=\"p\">(</span><span class=\"n\">max_length</span><span class=\"o\">=</span><span
  class=\"mi\">128</span><span class=\"p\">,</span> <span class=\"n\">unique</span><span
  class=\"o\">=</span><span class=\"kc\">True</span><span class=\"p\">)</span>\n    <span
  class=\"n\">content</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">TextField</span><span class=\"p\">()</span>\n
  \   <span class=\"n\">author</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">ForeignKey</span><span class=\"p\">(</span><span
  class=\"n\">Profile</span><span class=\"p\">,</span> <span class=\"n\">on_delete</span><span
  class=\"o\">=</span><span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">CASCADE</span><span class=\"p\">)</span>\n    <span class=\"n\">status</span>
  <span class=\"o\">=</span> <span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">CharField</span><span class=\"p\">(</span>\n        <span class=\"n\">max_length</span><span
  class=\"o\">=</span><span class=\"mi\">16</span><span class=\"p\">,</span>\n        <span
  class=\"n\">choices</span><span class=\"o\">=</span><span class=\"n\">Article_Status</span><span
  class=\"p\">,</span>\n        <span class=\"n\">default</span><span class=\"o\">=</span><span
  class=\"n\">Article_Status</span><span class=\"p\">[</span><span class=\"mi\">0</span><span
  class=\"p\">],</span>\n    <span class=\"p\">)</span>\n\n    <span class=\"k\">def</span>
  <span class=\"fm\">__str__</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"p\">):</span>\n        <span class=\"k\">return</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">title</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above model <code>Article</code>, we have a few fields like <code>title</code>
  simple Character Field, <code>content</code> as a text field as it will be a large
  text as the post body, <code>author</code> which is a ForeignKey to the User Model.
  We also have the status, which is defined as a character field but with a few choices
  like <code>draft</code> or <code>published</code>, we can further modify this status
  as public or private. But just keeping it simple and easy to understand.</p>\n<p>The
  object reference name for this model is the title as we have defined in the dunder
  string method. So, that is a simple model created, we can now migrate the changes
  into the database for adding the tables and attributes.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>python manage.py makemigrations\npython
  manage.py migrate\n</pre></div>\n\n</pre>\n\n<p>This will make migrations to the
  database i.e. convert the python model class into database tables and attributes.
  So, once the migration process is completed successfully, we can move into the crust
  of this article which is to actually design the views. In the next section, we will
  be utilizing the models in our views for representing the data on the templates.</p>\n<h2>Creating
  Article Form</h2>\n<p>Before diving into the views section, we need a few things
  like the Article Form, which will be a Django Model-based form. It will help us
  a lot in creating or updating the fields for the article model. We can define a
  form in a python file called <code>forms.py</code>, it's not necessary to keep your
  forms in the <code>forms.py</code> but if you have a lot of forms and models, it
  becomes a good practice to organize the components of our app. So, I'll be creating
  a new file inside of the <code>article</code> app called <code>forms.py</code> and
  defining the <code>ArticleForm</code>.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># article/forms.py</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">django</span>
  <span class=\"kn\">import</span> <span class=\"n\">forms</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">.models</span> <span class=\"kn\">import</span> <span class=\"n\">Article</span>\n\n\n<span
  class=\"k\">class</span> <span class=\"nc\">ArticleForm</span><span class=\"p\">(</span><span
  class=\"n\">forms</span><span class=\"o\">.</span><span class=\"n\">ModelForm</span><span
  class=\"p\">):</span>\n    <span class=\"k\">class</span> <span class=\"nc\">Meta</span><span
  class=\"p\">:</span>\n        <span class=\"n\">model</span> <span class=\"o\">=</span>
  <span class=\"n\">Article</span>\n        <span class=\"n\">exclude</span> <span
  class=\"o\">=</span> <span class=\"p\">(</span>\n            <span class=\"s2\">&quot;created&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;updated&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;author&quot;</span><span
  class=\"p\">,</span>\n        <span class=\"p\">)</span>\n        <span class=\"n\">widgets</span>
  <span class=\"o\">=</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;title&quot;</span><span
  class=\"p\">:</span> <span class=\"n\">forms</span><span class=\"o\">.</span><span
  class=\"n\">TextInput</span><span class=\"p\">(</span>\n                <span class=\"n\">attrs</span><span
  class=\"o\">=</span><span class=\"p\">{</span>\n                    <span class=\"s2\">&quot;class&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;form-control&quot;</span><span class=\"p\">,</span>\n
  \                   <span class=\"s2\">&quot;style&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;max-width: 450px; align: center;&quot;</span><span class=\"p\">,</span>\n
  \                   <span class=\"s2\">&quot;placeholder&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;Title&quot;</span><span class=\"p\">,</span>\n                <span
  class=\"p\">}</span>\n            <span class=\"p\">),</span>\n            <span
  class=\"s2\">&quot;content&quot;</span><span class=\"p\">:</span> <span class=\"n\">forms</span><span
  class=\"o\">.</span><span class=\"n\">Textarea</span><span class=\"p\">(</span>\n
  \               <span class=\"n\">attrs</span><span class=\"o\">=</span><span class=\"p\">{</span>\n
  \                   <span class=\"s2\">&quot;class&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;form-control&quot;</span><span class=\"p\">,</span>\n                    <span
  class=\"s2\">&quot;style&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;max-width:
  900px;&quot;</span><span class=\"p\">,</span>\n                    <span class=\"s2\">&quot;placeholder&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Content&quot;</span><span class=\"p\">,</span>\n
  \               <span class=\"p\">}</span>\n            <span class=\"p\">),</span>\n
  \       <span class=\"p\">}</span>\n</pre></div>\n\n</pre>\n\n<p>So, the forms are
  inherited from the [ModelForm] which allows us to create forms based on our model.
  So, we specify the model name which in our case is <code>Article</code> and further
  we can have <code>exclude</code> or <code>fields</code> tuples. To exclude certain
  fields in the actual form, just parse the tuple of those attributes and if you want
  to only select a few attributes, you can specify the <code>fields</code> tuple and
  mention the required fields for the form.</p>\n<p>So, if we have a lot of things
  to be included in the form, we can specify only the attributes to be excluded with
  the <code>exclude</code> tuple. And if we have a lot of fields to be excluded, we
  can use the <code>fields</code> tuple to specify which attributes to use in the
  form.</p>\n<p>Let's take an example: For the above ArticleForm, if we wanted to
  specify the required fields to be included in the form, then we might use the <code>fields</code>
  tuple like below the rest will be not rendered in the form fields.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>class ArticleForm(forms.ModelForm):\n
  \   class Meta:\n        model = Article\n        fields = (\n            &quot;title&quot;,\n
  \           &quot;content&quot;,\n            &quot;status&quot;,\n        )\n</pre></div>\n\n</pre>\n\n<p>Both
  of them can be used, it just depends on how many fields you have to exclude or include
  in the rendered form.</p>\n<p>We have also specified the <code>widgets</code> attribute
  which gives a bit more control on how we need to display the form in a template.
  So I have specified the type of input it needs to render like a simple text input
  for the title, text area for content, etc. The cool thing about this is it can automatically
  set these by knowing the type of field in the model, but sometimes it can be a bit
  undesired mostly with complex relationships and attributes.</p>\n<h2>Creating Views</h2>\n<p>Let's
  start creating views for creating, reading, updating, and deleting articles from
  the database. I will be using function-based views just because we are understanding
  the flow of how HTMX and Django can be integrated so we need to dive in deeper and
  understand the actual flow of the process.</p>\n<h3>Create View</h3>\n<p>So, creating
  articles seems like a good way to start off. We can create a simple function-based
  view which will initially load in an empty <code>ArticleForm</code> and if the request
  is <code>GET</code> we will render the form in the <code>create.html</code> template.
  If the request is <code>POST</code> which will be after we submit the form, we will
  validate the form and attach the current user as the author of the article and save
  the for instance which will create an article record and this object will be rendered
  to the detail template.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"kn\">from</span> <span class=\"nn\">django.shortcuts</span> <span class=\"kn\">import</span>
  <span class=\"n\">render</span>\n<span class=\"kn\">from</span> <span class=\"nn\">.models</span>
  <span class=\"kn\">import</span> <span class=\"n\">Article</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">.forms</span> <span class=\"kn\">import</span> <span class=\"n\">ArticleForm</span>\n\n<span
  class=\"k\">def</span> <span class=\"nf\">createArticle</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">):</span>\n    <span class=\"n\">form</span>
  <span class=\"o\">=</span> <span class=\"n\">ArticleForm</span><span class=\"p\">()</span>\n
  \   <span class=\"n\">context</span> <span class=\"o\">=</span> <span class=\"p\">{</span>\n
  \       <span class=\"s1\">&#39;form&#39;</span><span class=\"p\">:</span> <span
  class=\"n\">form</span><span class=\"p\">,</span>\n    <span class=\"p\">}</span>\n
  \   <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;articles/create.html&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<h4>Rendering
  the Form</h4>\n<p>We are creating an empty instance of <code>ArticleForm</code>
  and rendering it in the template. So, this will render the empty form in the <code>create.html</code>
  template.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cm\">&lt;!-- templates/articles/create.html --&gt;</span>\n\n{% extends
  &#39;base.html&#39; %}\n\n{% block body %}\n<span class=\"p\">&lt;</span><span class=\"nt\">div</span>
  <span class=\"na\">hx-target</span><span class=\"o\">=</span><span class=\"s\">&quot;this&quot;</span>
  <span class=\"na\">hx-swap</span><span class=\"o\">=</span><span class=\"s\">&quot;outerHTML&quot;</span><span
  class=\"p\">&gt;</span>\n  <span class=\"p\">&lt;</span><span class=\"nt\">form</span><span
  class=\"p\">&gt;</span>\n    {% csrf_token %}\n    {{ form.as_p }}\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">button</span> <span class=\"na\">hx-post</span><span class=\"o\">=</span><span
  class=\"s\">&quot;.&quot;</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;btn btn-success&quot;</span>\n      <span class=\"na\">type</span><span
  class=\"o\">=</span><span class=\"s\">&quot;submit&quot;</span><span class=\"p\">&gt;</span>Save<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;</span>\n
  \ <span class=\"p\">&lt;/</span><span class=\"nt\">form</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;/</span><span class=\"nt\">div</span><span class=\"p\">&gt;</span>\n{%
  endblock %}\n</pre></div>\n\n</pre>\n\n<p>Now, here we are inheriting from the base
  template and creating a form tag in HTML with the <code>{{ form}}</code> for rendering
  the form fields, we finally have the <code>button</code> element for submitting
  the form. We have used the <code>hx-post</code> attribute. More on this in just
  a minute. So, this is we create a template for rendering the article form.</p>\n<p>We
  have used the <code>hx-post</code> attribute here, which will send a <code>POST</code>
  request to the current <code>URL</code> represented by <code>hx-post=&quot;.&quot;</code>.
  You might have noticed the <code>div</code> attributes, the <code>hx-target</code>
  and <code>hx-swap</code>, so these are some of the many attributes provided by the
  htmx library for controlling the reactivity of the requests made. The <code>hx-target</code>
  allow us to specify the element or tag to which the data will be rendered. The <code>hx-swap</code>
  goes hand-in-hand for specifying the target DOM like <code>innerHTML</code>, <code>outerHTML</code>,
  etc. You can see the various options on the <a href=\"https://htmx.org/docs/#swapping\">htmx
  docs</a>. By specifying the <code>hx-swap</code> as  <code>outerHTML</code>, we
  are saying to replace the entire element with the incoming content from the request
  which we will send with nearby request triggers.</p>\n<p>We need to map the view
  to a URL in order to get a good idea about the request and parsed content.</p>\n<p>We'll
  create a <code>create/</code> route and bind it to the <code>createArticle</code>
  view with the name <code>article-create</code>.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># article/urls.py</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">django.urls</span>
  <span class=\"kn\">import</span> <span class=\"n\">path</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">.</span> <span class=\"kn\">import</span> <span class=\"n\">views</span>\n\n<span
  class=\"n\">urlpatterns</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;create/&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">createArticle</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;article-create&#39;</span><span class=\"p\">),</span>
  \n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>This URL will be mapped
  to the global URL in the project, here we can simply specify the prefix for the
  URLs in the <code>article</code> app and include those URLs.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># htmx_blog/urls.py</span>\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.contrib</span> <span class=\"kn\">import</span> <span class=\"n\">admin</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.urls</span> <span class=\"kn\">import</span>
  <span class=\"n\">path</span><span class=\"p\">,</span> <span class=\"n\">include</span>\n\n<span
  class=\"n\">urlpatterns</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;admin/&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">admin</span><span class=\"o\">.</span><span
  class=\"n\">site</span><span class=\"o\">.</span><span class=\"n\">urls</span><span
  class=\"p\">),</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;user/&#39;</span><span class=\"p\">,</span> <span class=\"n\">include</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;user.urls&#39;</span><span class=\"p\">),</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;auth&#39;</span><span
  class=\"p\">),</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span> <span class=\"n\">include</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;article.urls&#39;</span><span class=\"p\">),</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;home&#39;</span><span
  class=\"p\">),</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>Feel,
  free to add any other URL pattern like for instance, the article app is at <code>/</code>
  i.e. <code>127.0.01.:8000/</code>, you can add any other name like <code>127.0.0.1:8000/article/</code>
  by adding <code>path('article/', include('article.urls'))</code>.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659252089/blog-media/django-htmx-create-view.png\"
  alt=\"Django HTMX Create view Form Template\" /></p>\n<p>So, finally, we are sending
  a <code>GET</code> request to the <code>127.0.0.1:8000/create/</code> and this will
  output the form. As we have a <code>POST</code> request embedded in the button inside
  the form, we will send the <code>POST</code> request to the same URL -&gt; <code>127.0.0.1:8000/create/</code>.</p>\n<h4>Submitting
  the Form</h4>\n<p>Let's handle the <code>POST</code> request in the create view.</p>\n<pre
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
  class=\"kn\">from</span> <span class=\"nn\">django.shortcuts</span> <span class=\"kn\">import</span>
  <span class=\"n\">render</span>\n<span class=\"kn\">from</span> <span class=\"nn\">.models</span>
  <span class=\"kn\">import</span> <span class=\"n\">Article</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">.forms</span> <span class=\"kn\">import</span> <span class=\"n\">ArticleForm</span>\n\n<span
  class=\"k\">def</span> <span class=\"nf\">createArticle</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">):</span>\n    <span class=\"n\">form</span>
  <span class=\"o\">=</span> <span class=\"n\">ArticleForm</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"o\">.</span><span class=\"n\">POST</span>
  <span class=\"ow\">or</span> <span class=\"kc\">None</span><span class=\"p\">)</span>\n
  \   <span class=\"k\">if</span> <span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">method</span> <span class=\"o\">==</span> <span class=\"s1\">&#39;POST&#39;</span><span
  class=\"p\">:</span>\n        <span class=\"k\">if</span> <span class=\"n\">form</span><span
  class=\"o\">.</span><span class=\"n\">is_valid</span><span class=\"p\">():</span>\n
  \           <span class=\"n\">form</span><span class=\"o\">.</span><span class=\"n\">instance</span><span
  class=\"o\">.</span><span class=\"n\">author</span> <span class=\"o\">=</span> <span
  class=\"n\">request</span><span class=\"o\">.</span><span class=\"n\">user</span>\n
  \           <span class=\"n\">article</span> <span class=\"o\">=</span> <span class=\"n\">form</span><span
  class=\"o\">.</span><span class=\"n\">save</span><span class=\"p\">()</span>\n            <span
  class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;articles/detail.html&#39;</span><span
  class=\"p\">,</span> <span class=\"p\">{</span><span class=\"s1\">&#39;article&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">article</span><span class=\"p\">})</span>\n
  \   <span class=\"n\">context</span> <span class=\"o\">=</span> <span class=\"p\">{</span>\n
  \       <span class=\"s1\">&#39;form&#39;</span><span class=\"p\">:</span> <span
  class=\"n\">form</span><span class=\"p\">,</span>\n    <span class=\"p\">}</span>\n
  \   <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;articles/create.html&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p><strong>Simple
  explanation</strong></p>\n<ul>\n<li>Create a form instance of ArticleForm with the
  request data or empty -&gt; <code>ArticleForm(request.POST or None)</code></li>\n<li>If
  it's a POST request, validate and create the article, render the article object
  in <code>detail.html</code> template.</li>\n<li>If it's a GET request, render the
  empty form in <code>create.html</code></li>\n</ul>\n<p>There are a few changes in
  the view, instead of initializing the form to empty i.e. <code>ArticleForm()</code>,
  we are initializing with <code>ArticleForm(request.POST or None)</code>. This basically
  means that if we are having something in the <code>request.POST</code> dict, we
  will initialize the Form with that data or else an empty form instance.</p>\n<p>Next,
  we check if the request if <code>POST</code>, if it is then we check if the form
  is valid i.e. the form fields are not empty or if any other constraint on the model
  attributes is satisfied or not. If the form data is valid, we attach the author
  as the currently logged-in User/user who sent the request. Finally, we save the
  form which in turn creates the article record in the database. We then render the
  created article in the <code>detail.html</code> template which is not yet created.</p>\n<p>So,
  the <code>htmx-post</code> attribute has worked and it will send a post request
  to the same URL i.e. <code>127.0.0.1:8000/create</code> and this will again trigger
  the view <code>createArticle</code> this time we will have <code>request.POST</code>
  data. So, we will validate and save the form.</p>\n<h3>Detail View</h3>\n<p>The
  detail view is used for viewing the details of an article. This will be rendered
  after the article has been created or updated. This is quite simple, we need an
  <code>id</code> or <code>primary key(pk)</code> of an article and render the <code>title</code>
  and <code>content</code> of the article in the template.</p>\n<p>We pass in a primary
  key along with the request as a parameter to the view, the <code>pk</code> will
  be passed via the URL. We fetch the Article object with the id as the parsed <code>pk</code>
  and finally render the <code>detail.html</code> template with the article object.
  The <code>context['article']</code> can be accessed from the template to render
  the specific attributes like <code>title</code>, <code>content</code>, etc.</p>\n<pre
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
  class=\"c1\"># article/views.py</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">detailArticle</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"n\">pk</span><span class=\"p\">):</span>\n    <span class=\"n\">article</span>
  <span class=\"o\">=</span> <span class=\"n\">Article</span><span class=\"o\">.</span><span
  class=\"n\">objects</span><span class=\"o\">.</span><span class=\"n\">get</span><span
  class=\"p\">(</span><span class=\"nb\">id</span><span class=\"o\">=</span><span
  class=\"n\">pk</span><span class=\"p\">)</span>\n    <span class=\"n\">context</span>
  <span class=\"o\">=</span> <span class=\"p\">{</span><span class=\"s1\">&#39;article&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">article</span><span class=\"p\">}</span>\n
  \   <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;articles/detail.html&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>We
  can now bind the view to a URL and parse the required parameter <code>pk</code>
  to the view.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"kn\">from</span> <span class=\"nn\">django.urls</span> <span class=\"kn\">import</span>
  <span class=\"n\">path</span>\n<span class=\"kn\">from</span> <span class=\"nn\">.</span>
  <span class=\"kn\">import</span> <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;create/&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">createArticle</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s1\">&#39;article-create&#39;</span><span class=\"p\">),</span> \n    <span
  class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;&lt;int:pk&gt;&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">detailArticle</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;article-detail&#39;</span><span class=\"p\">),</span>
  \n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>We have parsed the <code>pk</code>
  as <code>int</code> to the URL parameter, so for an article with id=4, the URL will
  be, <code>127.0.0.1:8000/4/</code>.</p>\n<p>We need to create the template for rendering
  the context from the <code>detailArticle</code> view. So, we create the <code>detail.html</code>
  in the <code>templates/articles</code> folder. We inherit the base template and
  render the <code>article.title</code> and the <code>article.content</code> with
  a linebreaks template filter so as to display the content properly.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cm\">&lt;!-- templates/articles/detail.html --&gt;</span>\n\n\n{% extends
  &#39;base.html&#39; %}\n{% block body %}\n<span class=\"p\">&lt;</span><span class=\"nt\">div</span>
  <span class=\"na\">id</span><span class=\"o\">=</span><span class=\"s\">&quot;article-card&quot;</span><span
  class=\"p\">&gt;</span>\n  <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span
  class=\"p\">&gt;</span>{{ article.title }}\n  <span class=\"p\">&lt;</span><span
  class=\"nt\">p</span><span class=\"p\">&gt;</span>{{ article.content|linebreaks|safe
  }}<span class=\"p\">&lt;/</span><span class=\"nt\">p</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;</span><span class=\"nt\">div</span><span class=\"p\">&gt;</span>\n{%
  endblock %}\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659252227/blog-media/django-htmx-detail-view.png\"
  alt=\"Detail View Template\" /></p>\n<p>So, we can now use <code>createArticle</code>
  view as well as <code>detailArticle</code> view, this both are configured properly,
  so (CR) or CRUD is completed. We can add <code>listArticle</code> for listing out
  all the author's(logged-in user) articles.</p>\n<h3>List View</h3>\n<p>Listview
  of the articles is much similar to the detail view as it will return a list of articles
  rather than a single article.</p>\n<p>So in the <code>listArticle</code> view, we
  will return all the articles with the author as the user who sent the request/logged-in
  user. We will parse this object list into the template as <code>base.html</code>
  or <code>articles/list.html</code>.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># article/views.py</span>\n\n\n<span class=\"k\">def</span> <span class=\"nf\">listArticle</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"n\">articles</span> <span class=\"o\">=</span> <span class=\"n\">Article</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">filter</span><span class=\"p\">(</span><span class=\"n\">author</span><span
  class=\"o\">=</span><span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">user</span><span class=\"o\">.</span><span class=\"n\">id</span><span
  class=\"p\">)</span>\n    <span class=\"n\">context</span> <span class=\"o\">=</span>
  <span class=\"p\">{</span>\n        <span class=\"s1\">&#39;articles&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">articles</span><span class=\"p\">,</span>\n
  \   <span class=\"p\">}</span>\n    <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;base.html&#39;</span><span class=\"p\">,</span> <span class=\"n\">context</span><span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>We will add the URL route for
  this as the <code>/</code> route that is on <code>127.0.0.1:8000/</code> this is
  the base URL for the article app and is the route for the <code>listArticle</code>
  view. So, we will display the list of articles on the homepage.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># article/urls.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.urls</span> <span class=\"kn\">import</span> <span class=\"n\">path</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">.</span> <span class=\"kn\">import</span>
  <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span> <span class=\"o\">=</span>
  <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&lt;int:pk&gt;&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">detailArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-detail&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;create/&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">createArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-create&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">listArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-list&#39;</span><span
  class=\"p\">),</span> \n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>Let's
  create the template for the list view which will iterate over the articles and display
  the relevant data like the title and link to the article.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cm\">&lt;!-- templates/articles/list.html --&gt;</span>\n\n<span class=\"p\">&lt;</span><span
  class=\"nt\">ul</span> <span class=\"na\">id</span><span class=\"o\">=</span><span
  class=\"s\">&quot;article-list&quot;</span><span class=\"p\">&gt;</span>\n  {% for
  article in articles %}\n  <span class=\"p\">&lt;</span><span class=\"nt\">li</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">div</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;card&quot;</span>
  <span class=\"na\">style</span><span class=\"o\">=</span><span class=\"s\">&quot;width:
  18rem;&quot;</span><span class=\"p\">&gt;</span>\n      <span class=\"p\">&lt;</span><span
  class=\"nt\">div</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;card-body&quot;</span><span class=\"p\">&gt;</span>\n        <span
  class=\"p\">&lt;</span><span class=\"nt\">h5</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;card-title&quot;</span><span class=\"p\">&gt;</span>{{
  article.title }}<span class=\"p\">&lt;/</span><span class=\"nt\">h5</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">p</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;card-text&quot;</span><span
  class=\"p\">&gt;</span>{{ article.content|truncatewords:5  }}<span class=\"p\">&lt;/</span><span
  class=\"nt\">p</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% url &#39;article-detail&#39; article.id %}&quot;</span> <span
  class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;card-link&quot;</span><span
  class=\"p\">&gt;</span>Read more<span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n      <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n  <span class=\"p\">&lt;/</span><span class=\"nt\">li</span><span
  class=\"p\">&gt;</span>\n  {% endfor %}\n<span class=\"p\">&lt;/</span><span class=\"nt\">ul</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>We have used the <code>truncatewords:5</code>
  template filter for only displaying the content of the articles till the first 5
  words as it is just a list view, we don't want to display every detail of the article
  here.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659252293/blog-media/django-htmx-list-view.png\"
  alt=\"List view Template\" /></p>\n<p>We can use this template to render in the
  <code>base.html</code> file.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"cp\">&lt;!DOCTYPE html&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">html</span>
  <span class=\"na\">lang</span><span class=\"o\">=</span><span class=\"s\">&quot;en&quot;</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">meta</span>
  <span class=\"na\">charset</span><span class=\"o\">=</span><span class=\"s\">&quot;UTF-8&quot;</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>HTMX Blog<span class=\"p\">&lt;/</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>\n    {% load static %}\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">link</span> <span class=\"na\">rel</span><span class=\"o\">=</span><span
  class=\"s\">&quot;stylesheet&quot;</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css&quot;</span>
  <span class=\"na\">integrity</span><span class=\"o\">=</span><span class=\"s\">&quot;sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm&quot;</span>
  <span class=\"na\">crossorigin</span><span class=\"o\">=</span><span class=\"s\">&quot;anonymous&quot;</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">script</span>
  <span class=\"na\">src</span><span class=\"o\">=</span><span class=\"s\">&quot;https://unpkg.com/htmx.org@1.8.0&quot;</span><span
  class=\"p\">&gt;&lt;/</span><span class=\"nt\">script</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;/</span><span class=\"nt\">head</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;</span><span class=\"nt\">body</span> <span class=\"na\">hx-target</span><span
  class=\"o\">=</span><span class=\"s\">&quot;this&quot;</span> <span class=\"na\">hx-swap</span><span
  class=\"o\">=</span><span class=\"s\">&quot;outerHTML&quot;</span> <span class=\"na\">hx-headers</span><span
  class=\"o\">=</span><span class=\"s\">&#39;{&quot;X-CSRFToken&quot;: &quot;{{ csrf_token
  }}&quot;}&#39;</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">nav</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>HTMX Blog<span class=\"p\">&lt;/</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">div</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;navbar&quot;</span><span class=\"p\">&gt;</span>\n          {%
  if user.is_authenticated %}\n            <span class=\"p\">&lt;</span><span class=\"nt\">a</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;nav-item
  nav-link&quot;</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% url &#39;article-list&#39; %}&quot;</span><span class=\"p\">&gt;&lt;</span><span
  class=\"nt\">button</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;btn btn-link&quot;</span><span class=\"p\">&gt;</span>Home<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;nav-item nav-link&quot;</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% url &#39;logout&#39; %}&quot;</span><span
  class=\"p\">&gt;&lt;</span><span class=\"nt\">button</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;btn btn-link&quot;</span><span class=\"p\">&gt;</span>Logout<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n          {% else %}\n            <span
  class=\"p\">&lt;</span><span class=\"nt\">a</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;nav-item nav-link&quot;</span> <span
  class=\"na\">href</span><span class=\"o\">=</span><span class=\"s\">&quot;{% url
  &#39;login&#39; %}&quot;</span><span class=\"p\">&gt;&lt;</span><span class=\"nt\">button</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;btn
  btn-link&quot;</span><span class=\"p\">&gt;</span>Login<span class=\"p\">&lt;/</span><span
  class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span class=\"nt\">a</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;nav-item
  nav-link&quot;</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% url &#39;register&#39; %}&quot;</span><span class=\"p\">&gt;&lt;</span><span
  class=\"nt\">button</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;btn btn-link&quot;</span><span class=\"p\">&gt;</span>Register<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n          {% endif %}\n        <span
  class=\"p\">&lt;/</span><span class=\"nt\">div</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;/</span><span class=\"nt\">nav</span><span class=\"p\">&gt;</span>\n\n
  \   {% block body %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">a</span>
  <span class=\"na\">href</span><span class=\"o\">=</span><span class=\"s\">&quot;{%
  url &#39;article-create&#39; %}&quot;</span><span class=\"p\">&gt;&lt;</span><span
  class=\"nt\">button</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;btn btn-success&quot;</span> <span class=\"p\">&gt;</span>Create<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n    {% include &#39;articles/list.html&#39;
  %}\n    {% endblock %}\n<span class=\"p\">&lt;/</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>We have now included the <code>list.html</code>
  template on the homepage and also added the <code>create</code> button as the link
  to the <code>article-create</code> URL.</p>\n<h3>Delete View</h3>\n<p>For deleting
  an article, we will simply rely on htmx for sending the request and on that request,
  we will delete the current article and render the updated list of articles.</p>\n<p>With
  the <code>deleteArticle</code> view, we will take in two parameters the request
  which is by default for a Django function-based view, and the primary key as <code>pk</code>.
  Again we will parse the <code>pk</code> from the URL. We will delete the article
  object and get the latest list of articles. Finally, render the updated list of
  articles in the base template which is our list view.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># article/views.py</span>\n\n\n<span class=\"k\">def</span> <span class=\"nf\">deleteArticle</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"n\">pk</span><span class=\"p\">):</span>\n    <span class=\"n\">Article</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">get</span><span class=\"p\">(</span><span class=\"nb\">id</span><span
  class=\"o\">=</span><span class=\"n\">pk</span><span class=\"p\">)</span><span class=\"o\">.</span><span
  class=\"n\">delete</span><span class=\"p\">()</span>\n    <span class=\"n\">articles</span>
  <span class=\"o\">=</span> <span class=\"n\">Article</span><span class=\"o\">.</span><span
  class=\"n\">objects</span><span class=\"o\">.</span><span class=\"n\">filter</span><span
  class=\"p\">(</span><span class=\"n\">author</span><span class=\"o\">=</span><span
  class=\"n\">request</span><span class=\"o\">.</span><span class=\"n\">user</span><span
  class=\"p\">)</span>\n    <span class=\"n\">context</span> <span class=\"o\">=</span>
  <span class=\"p\">{</span><span class=\"s1\">&#39;article&#39;</span><span class=\"p\">:</span>
  <span class=\"n\">articles</span><span class=\"p\">}</span>\n    <span class=\"k\">return</span>
  <span class=\"n\">render</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;base.html&quot;</span><span class=\"p\">,</span>
  <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>We
  will add the <code>deleteArticle</code> into the URL patterns and call it <code>article-delete</code>
  with the URL of <code>delete/&lt;int:pk&gt;</code>. This will allow us to send a
  request to the URL <code>127.0.0.1:8000/delete/4</code> for deleting the article
  with id <code>4</code>.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># article/urls.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.urls</span> <span class=\"kn\">import</span> <span class=\"n\">path</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">.</span> <span class=\"kn\">import</span>
  <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span> <span class=\"o\">=</span>
  <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">listArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-list&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&lt;int:pk&gt;&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">detailArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-detail&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;create/&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">createArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-create&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;delete/&lt;int:pk&gt;&#39;</span><span class=\"p\">,</span> <span
  class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">deleteArticle</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s1\">&#39;article-delete&#39;</span><span class=\"p\">),</span> \n<span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>In the delete view, the template
  is important as we want to send a request appropriately to the defined URL. To do
  that, we will have a form but it won't have any inputs as such just a button that
  indicates to delete the current article. We will add the <code>hx-delete</code>
  attribute as the URL to the <code>deleteArticle</code> view. with the id of the
  article. This will send a request to the <code>article-delete</code> URL which will,
  in turn, trigger the view with the given id and delete the article.</p>\n<p>We have
  added the <code>hx-confirm</code> attribute for showing a pop-up of confirmation
  of deleting the article. As you can see we have added a little script for adding
  <code>csrf_token</code> into the HTML, this is important in order to submit a form
  with a valid <code>CSRFToken</code>.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"cm\">&lt;!-- templates/article/delete.html --&gt;</span>\n\n<span class=\"p\">&lt;</span><span
  class=\"nt\">script</span><span class=\"p\">&gt;</span>\n  <span class=\"nb\">document</span><span
  class=\"p\">.</span><span class=\"nx\">body</span><span class=\"p\">.</span><span
  class=\"nx\">addEventListener</span><span class=\"p\">(</span><span class=\"s1\">&#39;htmx:configRequest&#39;</span><span
  class=\"p\">,</span> <span class=\"p\">(</span><span class=\"nx\">event</span><span
  class=\"p\">)</span> <span class=\"p\">=&gt;</span> <span class=\"p\">{</span>\n
  \   <span class=\"nx\">event</span><span class=\"p\">.</span><span class=\"nx\">detail</span><span
  class=\"p\">.</span><span class=\"nx\">headers</span><span class=\"p\">[</span><span
  class=\"s1\">&#39;X-CSRFToken&#39;</span><span class=\"p\">]</span> <span class=\"o\">=</span>
  <span class=\"s1\">&#39;{{ csrf_token }}&#39;</span><span class=\"p\">;</span>\n
  \ <span class=\"p\">})</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">script</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">div</span>
  <span class=\"p\">&gt;</span>\n  <span class=\"p\">&lt;</span><span class=\"nt\">form</span>
  <span class=\"na\">method</span><span class=\"o\">=</span><span class=\"s\">&quot;post&quot;</span>
  <span class=\"p\">&gt;</span>\n  {% csrf_token %}\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">button</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;btn btn-danger&quot;</span>\n      <span class=\"na\">hx-delete</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% url &#39;article-delete&#39; article.id
  %}&quot;</span>\n      <span class=\"na\">hx-confirm</span><span class=\"o\">=</span><span
  class=\"s\">&quot;Are you sure, You want to delete this article?&quot;</span>\n
  \     <span class=\"na\">type</span><span class=\"o\">=</span><span class=\"s\">&quot;submit&quot;</span><span
  class=\"p\">&gt;</span>\n      Delete\n    <span class=\"p\">&lt;/</span><span class=\"nt\">button</span><span
  class=\"p\">&gt;</span>\n  <span class=\"p\">&lt;/</span><span class=\"nt\">form</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>Do you have a question like
  how do we access the <code>article.id</code>? we are not rendering the <code>delete.html</code>
  template from the view, so there is no context to pass. We will include this snippet
  into the detail view template, so as to have the option of deleting the current
  article.</p>\n<p>We will modify the <code>articles/detail.html</code> template and
  include the <code>delete.html</code> template. This includes simply adding an HTML
  template in the specified location. So, we will basically inject the delete form
  into the detail template.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>{%
  extends &#39;base.html&#39; %}\n{% block body %}\n<span class=\"p\">&lt;</span><span
  class=\"nt\">div</span> <span class=\"na\">hx-target</span><span class=\"o\">=</span><span
  class=\"s\">&quot;this&quot;</span> <span class=\"na\">hx-swap</span><span class=\"o\">=</span><span
  class=\"s\">&quot;outerHTML&quot;</span><span class=\"p\">&gt;</span>\n  <span class=\"p\">&lt;</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>{{ article.title }}\n  {% include
  &#39;articles/delete.html&#39; %}\n  <span class=\"p\">&lt;</span><span class=\"nt\">p</span><span
  class=\"p\">&gt;</span>{{ article.content|linebreaks|safe }}<span class=\"p\">&lt;/</span><span
  class=\"nt\">p</span><span class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span
  class=\"nt\">div</span><span class=\"p\">&gt;</span>\n{% endblock %}\n</pre></div>\n\n</pre>\n\n<p>Hence,
  we will have a nice option to delete the article in the detail section, this can
  be placed anywhere but remember, we need to add the <code>hx-target=&quot;this&quot;</code>
  and <code>hx-swap=&quot;outerHTML&quot;</code> in the div so as to correctly swap
  the HTML content after the request has been made.</p>\n<h3>Update View</h3>\n<p>We
  can now move into the final piece of the CRUD i.e. <code>Update</code>. This will
  be similar to the <code>createArticle</code> with a couple of changes. We will parse
  parameters like <code>pk</code> to this view as well because we want to update a
  specific article. So, we will have to get the primary key of the article from the
  URL slug.</p>\n<p>Inside the <code>updateArticle</code> view, we will first grab
  the article object from the parsed primary key. We will have two kinds of requests
  here, one will be for fetching the <code>form</code> with the current article data,
  and the next request will be the <code>PUT</code> request for actually saving the
  changes in the article.</p>\n<p>The first request is simple as we need to parse
  the form data with the instance of the article object. We will call the <code>ArticleForm</code>
  with the instance of <code>article</code> this will load the data of the article
  into the form ready to render into the template. So once the <code>GET</code> request
  has been sent, we will render the template with the form pre-filled with the values
  of the article attributes.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># article/views.py</span>\n\n\n<span class=\"k\">def</span> <span class=\"nf\">updateArticle</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"n\">pk</span><span class=\"p\">):</span>\n   <span class=\"n\">article</span>
  <span class=\"o\">=</span> <span class=\"n\">Article</span><span class=\"o\">.</span><span
  class=\"n\">objects</span><span class=\"o\">.</span><span class=\"n\">get</span><span
  class=\"p\">(</span><span class=\"nb\">id</span><span class=\"o\">=</span><span
  class=\"n\">pk</span><span class=\"p\">)</span>\n   <span class=\"n\">form</span>
  <span class=\"o\">=</span> <span class=\"n\">ArticleForm</span><span class=\"p\">(</span><span
  class=\"n\">instance</span><span class=\"o\">=</span><span class=\"n\">article</span><span
  class=\"p\">)</span>\n   <span class=\"n\">context</span> <span class=\"o\">=</span>
  <span class=\"p\">{</span>\n       <span class=\"s1\">&#39;form&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">form</span><span class=\"p\">,</span>\n       <span
  class=\"s1\">&#39;article&#39;</span><span class=\"p\">:</span> <span class=\"n\">article</span><span
  class=\"p\">,</span>\n   <span class=\"p\">}</span>\n   <span class=\"k\">return</span>
  <span class=\"n\">render</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;articles/update.html&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>We
  will create a template in the <code>templates/articles/</code> folder as  <code>update.html</code>
  which will have a simple form for rendering the form fields and a button for sending
  a <code>PUT</code> request. We will render the <code>form</code> and then add a
  button element with the attribute <code>hx-put</code> for sending the <code>PUT</code>
  request to save changes to the article record. We will parse in the <code>article.id</code>
  for the primary key parameter to the view.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"cm\">&lt;!-- templates/articles/update.html --&gt;</span>\n\n<span class=\"p\">&lt;</span><span
  class=\"nt\">div</span> <span class=\"na\">hx-target</span><span class=\"o\">=</span><span
  class=\"s\">&quot;this&quot;</span> <span class=\"na\">hx-swap</span><span class=\"o\">=</span><span
  class=\"s\">&quot;outerHTML&quot;</span><span class=\"p\">&gt;</span>\n  <span class=\"p\">&lt;</span><span
  class=\"nt\">form</span><span class=\"p\">&gt;</span>\n    {% csrf_token %}\n    {{
  form.as_p }}\n    <span class=\"p\">&lt;</span><span class=\"nt\">button</span>
  <span class=\"na\">hx-put</span><span class=\"o\">=</span><span class=\"s\">&quot;{%
  url &#39;article-update&#39; article.id %}&quot;</span>\n      <span class=\"na\">type</span><span
  class=\"o\">=</span><span class=\"s\">&quot;submit&quot;</span><span class=\"p\">&gt;</span>Update<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;</span>\n
  \ <span class=\"p\">&lt;/</span><span class=\"nt\">form</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;/</span><span class=\"nt\">div</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>We
  are yet to link the <code>updateArticle</code> into the URLs. We will add the view
  <code>updateArticle</code> into the URLs with the name as <code>article-update</code>
  and <code>update/&lt;int:pk</code> as the slug pattern. This URL pattern will trigger
  the <code>updateArticle</code> when we send an HTTP request to the <code>127.0.0.1:8000/update/4</code>
  for updating the article with id as <code>4</code>.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># article/urls.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.urls</span> <span class=\"kn\">import</span> <span class=\"n\">path</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">.</span> <span class=\"kn\">import</span>
  <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span> <span class=\"o\">=</span>
  <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">listArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-list&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&lt;int:pk&gt;&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">detailArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-detail&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;create/&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">createArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-create&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;delete/&lt;int:pk&gt;&#39;</span><span class=\"p\">,</span> <span
  class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">deleteArticle</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s1\">&#39;article-delete&#39;</span><span class=\"p\">),</span> \n    <span
  class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;update/&lt;int:pk&gt;&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">updateArticle</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;article-update&#39;</span><span class=\"p\">),</span>
  \n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>This is not done yet,
  we will need to handle the <code>PUT</code> request as well i.e. when the form details
  have been modified and we are about to save changes to the form data. So, we will
  check for the request method's type. If it is a <code>PUT</code> request, we will
  have to process a few things.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># article/views.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.http</span> <span class=\"kn\">import</span> <span class=\"n\">QueryDict</span>\n\n<span
  class=\"k\">def</span> <span class=\"nf\">updateArticle</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"n\">pk</span><span
  class=\"p\">):</span>\n    <span class=\"n\">article</span> <span class=\"o\">=</span>
  <span class=\"n\">Article</span><span class=\"o\">.</span><span class=\"n\">objects</span><span
  class=\"o\">.</span><span class=\"n\">get</span><span class=\"p\">(</span><span
  class=\"nb\">id</span><span class=\"o\">=</span><span class=\"n\">pk</span><span
  class=\"p\">)</span>\n    <span class=\"k\">if</span> <span class=\"n\">request</span><span
  class=\"o\">.</span><span class=\"n\">method</span> <span class=\"o\">==</span>
  <span class=\"s1\">&#39;PUT&#39;</span><span class=\"p\">:</span>\n        <span
  class=\"n\">qd</span> <span class=\"o\">=</span> <span class=\"n\">QueryDict</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">body</span><span class=\"p\">)</span>\n        <span class=\"n\">form</span>
  <span class=\"o\">=</span> <span class=\"n\">ArticleForm</span><span class=\"p\">(</span><span
  class=\"n\">instance</span><span class=\"o\">=</span><span class=\"n\">article</span><span
  class=\"p\">,</span> <span class=\"n\">data</span><span class=\"o\">=</span><span
  class=\"n\">qd</span><span class=\"p\">)</span>\n        <span class=\"k\">if</span>
  <span class=\"n\">form</span><span class=\"o\">.</span><span class=\"n\">is_valid</span><span
  class=\"p\">():</span>\n            <span class=\"n\">article</span> <span class=\"o\">=</span>
  <span class=\"n\">form</span><span class=\"o\">.</span><span class=\"n\">save</span><span
  class=\"p\">()</span>\n            <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;articles/detail.html&#39;</span><span class=\"p\">,</span> <span
  class=\"p\">{</span><span class=\"s1\">&#39;article&#39;</span><span class=\"p\">:</span>
  <span class=\"n\">article</span><span class=\"p\">})</span>\n    <span class=\"n\">form</span>
  <span class=\"o\">=</span> <span class=\"n\">ArticleForm</span><span class=\"p\">(</span><span
  class=\"n\">instance</span><span class=\"o\">=</span><span class=\"n\">article</span><span
  class=\"p\">)</span>\n    <span class=\"n\">context</span> <span class=\"o\">=</span>
  <span class=\"p\">{</span>\n        <span class=\"s1\">&#39;form&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">form</span><span class=\"p\">,</span>\n        <span
  class=\"s1\">&#39;article&#39;</span><span class=\"p\">:</span> <span class=\"n\">article</span><span
  class=\"p\">,</span>\n    <span class=\"p\">}</span>\n    <span class=\"k\">return</span>
  <span class=\"n\">render</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;articles/update.html&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above <code>updateArticle</code> view, we have to check for a <code>PUT</code>
  request, if we are sending a <code>PUT</code> request, the form instance needs to
  be loaded from the request object. We use the <code>request.body</code> to access
  the data sent in the request. The incoming data received from the <code>request.body</code>
  object is not a valid format to parse it to the form instance, so we will parse
  it using <code>QueryDict</code>. This will allow us to modify the <code>request.body</code>
  object into valid python serializable data.</p>\n<p>So, we import the <code>QueryDict</code>
  from <code>django.http</code> module. We parse the data as the parameter to <code>QueryDict</code>
  and store it in a variable. We then have to get the <code>ArticleForm</code> for
  fetching the data as per the form details, so we parse the instance and also the
  <code>data</code> parameter. The instance is the article object and the data is
  the received form data which we have stored in <code>qd</code> as <code>QueryDict(request.body)</code>.
  This will load the new form data and then we can validate it the form.</p>\n<p>After
  we have verified the form details, we can save the form and this will update the
  article record. Thereby we can render the updated article in the <code>detail</code>
  view with the updated <code>article</code> object as the context.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659252091/blog-media/django-htmx-update-view.png\"
  alt=\"Update View Form Template\" /></p>\n<p>So, this will set up the update view
  as well, we can now create, read, update, and delete an article instance with HTMX
  in templates and Django function-based views without writing any javascript.</p>\n<h2>Summary</h2>\n<p>We
  were able to create a basic CRUD application in Django with HTMX. We used simple
  function-based views to demonstrate the inner details of how we can work with HTMX
  and handle requests from the templates. By creating simple standalone templates,
  we can connect those together to make a fully functional and responsive webpage.
  The UI is not great but the purpose of this tutorial was to make a barebone CRUD
  app to work with the backend using HTMX, so hopefully, you would have got a good
  overview of how HTMX can be integrated into a Django application.</p>\n<p>Overall
  HTMX is a great library that can be used to enhance or even create a new web application
  for making the site responsive and without writing any javascript.</p>\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659252296/blog-media/django-htmx-demo.gif\"
  alt=\"Django HTMX CRUD Application Demo GIF\" /></p>\n<p>You can check out the source
  code for this project and blog on the <a href=\"https://github.com/Mr-Destructive/htmx-blog-django\">htmx-blog
  GitHub</a> repository.</p>\n<h2>Conclusion</h2>\n<p>From this post, we were able
  to understand the basics of HTMX and how we can integrate it into a Django application.
  Hopefully, you enjoyed the post, if you have any queries or feedback, please let
  me know in the comments or on my social handles. Thank you for reading. Happy Coding
  :)</p>\n"
cover: ''
date: 2022-07-31
datetime: 2022-07-31 00:00:00+00:00
description: Creating a basic CRUD application with Django and HTMX
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2022-07-31-Django-HTMX-CRUD-App.md
html: "<h2>Introduction</h2>\n<p>Gone are the days of writing Ajax requests with javascript,
  just add a few parameters to the HTML content tags and you will be ready for sending
  requests to your backend. So, we are going back in time and correcting the way we
  think of APIs and client/server-side rendering. We are referring to the <a href=\"https://en.wikipedia.org/wiki/Hypermedia\">Hypermedia
  model</a> for levering the server-side processing of data. Let's get our feets wet
  with this ancient but revolutionary methodology of development with <a href=\"https://htmx.org/\">HTMX</a>.</p>\n<p>Yes,
  HTMX can be used for the API/server-side calls directly in the HTML. We will be
  exploring the basis of HTMX by creating a basic CRUD application.</p>\n<h2>What
  is HTMX?</h2>\n<p>The first question that might come up is what and why HTMX? Htmx
  is a great library, it's a javascript library but wait. It is a javascript library
  designed to allow us to write less or no javascript at all. It acts as a way to
  send AJAX requests without you writing any javascript. It uses native browser features
  directly from HTML.</p>\n<p>So, we can use HTMX to create interactive templates
  in our Django application. We can dynamically call and fetch data from the server
  by using simple HTML attributes like <code>hx-get</code>, <code>hx-post</code>,
  etc. We'll cover those in this article.</p>\n<p>You can check the source code used
  in this article on this <a href=\"https://github.com/Mr-Destructive/htmx-blog-django\">GitHub
  repository</a>.</p>\n<h2>Setup Django Project</h2>\n<p>We'll be creating a Django
  project from scratch and designing a basic blog kind of app. We will be creating
  a quite simple project with a couple of apps like <code>user</code> for authentication
  and <code>article</code> for the CRUD part of our blog application.</p>\n<p>To set
  up a django project, we can run the following commands to quickly get up and running
  with a base django project.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>mkdir htmx_blog\npython3
  -m venv .venv\nsource .venv/bin/activate\npip install django\ndjango-admin startproject
  htmx_blog .\n</pre></div>\n\n</pre>\n\n<p>I have a base user model that I use for
  a simple authentication system in some basic django projects, you can define your
  own user app or get the app from <a href=\"https://github.com/Mr-Destructive/django-todo/tree/master/user\">here</a>.</p>\n<p>So,
  that being said, we will be using the user model for the article model which we
  will be defined next. By creating a basic signup functionality, you are good to
  go!</p>\n<h3>Create the Article app</h3>\n<p>We will need at least an app to work
  with htmx as we will define models, views, and URLs later as we configure the htmx.</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>django-admin startapp article\n</pre></div>\n\n</pre>\n\n<p>After
  the app has been created, you can add those app labels into the <code>INSTALLED_APPS</code>
  config in the <code>settings.py</code> file. The <code>user</code> app and the <code>article</code>
  app need to be added to the installed apps for the django to pick those up for various
  contexts related to the project.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span># htmx_blog/settings.py\n\nINSTALLED_APPS
  = [\n    ...\n    ...\n    ...\n\n    &#39;article&#39;,  \n    &#39;user&#39;,\n]\n</pre></div>\n\n</pre>\n\n<p>We
  are sone with the base setup, we also would require a few more configs for the proper
  working of the project.</p>\n<h3>Setup Templates and Static files</h3>\n<p>Templates
  will play an important role in the htmx part, so it is equally important to configure
  them properly before dabbling into the htmx and client-side rendering of data.</p>\n<p>I
  like to keep all the templates in a single folder in the <code>BASE_DIR</code> with
  separate sub-folders for specific apps. Also a single <code>static</code> folder
  with <code>css</code>, <code>js</code>, and <code>images</code> as the sub-folfers
  for a larger project.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>mkdir templates static\n</pre></div>\n\n</pre>\n\n<p>Further,
  configure the created static and templates in the settings.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"n\">TEMPLATES</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"p\">{</span>\n        <span class=\"s1\">&#39;BACKEND&#39;</span><span
  class=\"p\">:</span> <span class=\"s1\">&#39;django.template.backends.django.DjangoTemplates&#39;</span><span
  class=\"p\">,</span>\n        <span class=\"s1\">&#39;DIRS&#39;</span><span class=\"p\">:</span>
  <span class=\"p\">[</span><span class=\"n\">os</span><span class=\"o\">.</span><span
  class=\"n\">path</span><span class=\"o\">.</span><span class=\"n\">join</span><span
  class=\"p\">(</span><span class=\"n\">BASE_DIR</span><span class=\"p\">,</span>
  <span class=\"s2\">&quot;templates&quot;</span><span class=\"p\">)],</span>\n        <span
  class=\"s1\">&#39;APP_DIRS&#39;</span><span class=\"p\">:</span> <span class=\"kc\">True</span><span
  class=\"p\">,</span>\n        <span class=\"s1\">&#39;OPTIONS&#39;</span><span class=\"p\">:</span>
  <span class=\"p\">{</span>\n            <span class=\"s1\">&#39;context_processors&#39;</span><span
  class=\"p\">:</span> <span class=\"p\">[</span>\n                <span class=\"s1\">&#39;django.template.context_processors.debug&#39;</span><span
  class=\"p\">,</span>\n                <span class=\"s1\">&#39;django.template.context_processors.request&#39;</span><span
  class=\"p\">,</span>\n                <span class=\"s1\">&#39;django.contrib.auth.context_processors.auth&#39;</span><span
  class=\"p\">,</span>\n                <span class=\"s1\">&#39;django.contrib.messages.context_processors.messages&#39;</span><span
  class=\"p\">,</span>\n            <span class=\"p\">],</span>\n        <span class=\"p\">},</span>\n
  \   <span class=\"p\">},</span>\n<span class=\"p\">]</span>\n\n<span class=\"n\">STATIC_URL</span>
  <span class=\"o\">=</span> <span class=\"s1\">&#39;static/&#39;</span>\n<span class=\"n\">STATICFILES_DIRS</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span><span class=\"nb\">str</span><span
  class=\"p\">(</span><span class=\"n\">BASE_DIR</span><span class=\"o\">/</span>
  <span class=\"s2\">&quot;static&quot;</span><span class=\"p\">)]</span>\n<span class=\"n\">STATIC_ROOT</span>
  <span class=\"o\">=</span> <span class=\"n\">BASE_DIR</span> <span class=\"o\">/</span>
  <span class=\"s2\">&quot;staticfiles&quot;</span>\n</pre></div>\n\n</pre>\n\n<h3>Initial
  migration</h3>\n<p>Run migration command for the user model and default model in
  the django project.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  manage.py migrate\n</pre></div>\n\n</pre>\n\n<p>So, this project will also include
  authentication simple registration, and login/logout routes. We will be using the
  default Django User model by creating an abstract user just in case we require any
  additional attributes.</p>\n<h2>Setup HTMX</h2>\n<p>We don't have to configure much
  for using HTMX as it is a javascript library, we can call it via a CDN or manually
  install it and link up the static javascript files. Either way, both are equally
  good, you may like the one I might like the other.</p>\n<p>If you already have a
  base template, you can simply put the below script inside the head tag of the template.
  This will make us the htmx attributes available.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">&lt;</span><span class=\"nt\">script</span> <span class=\"na\">src</span><span
  class=\"o\">=</span><span class=\"s\">&quot;https://unpkg.com/htmx.org@1.8.0&quot;</span><span
  class=\"p\">&gt;&lt;/</span><span class=\"nt\">script</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>If
  you don't have a base template, you can create one by creating an HTML file inside
  the <code>templates</code> directory. The name can be anything but be careful for
  following up as it might be different for me. I will choose <code>base.html</code>
  as the template for this project. It will look something like as follows:</p>\n<pre
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cm\">&lt;!-- tempaltes/base.html --&gt;</span>\n\n<span class=\"cp\">&lt;!DOCTYPE
  html&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">html</span> <span
  class=\"na\">lang</span><span class=\"o\">=</span><span class=\"s\">&quot;en&quot;</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">meta</span>
  <span class=\"na\">charset</span><span class=\"o\">=</span><span class=\"s\">&quot;UTF-8&quot;</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>HTMX Blog<span class=\"p\">&lt;/</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>\n    {% load static %}\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">link</span> <span class=\"na\">rel</span><span class=\"o\">=</span><span
  class=\"s\">&quot;stylesheet&quot;</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css&quot;</span>
  <span class=\"na\">integrity</span><span class=\"o\">=</span><span class=\"s\">&quot;sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm&quot;</span>
  <span class=\"na\">crossorigin</span><span class=\"o\">=</span><span class=\"s\">&quot;anonymous&quot;</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">script</span>
  <span class=\"na\">src</span><span class=\"o\">=</span><span class=\"s\">&quot;https://unpkg.com/htmx.org@1.8.0&quot;</span><span
  class=\"p\">&gt;&lt;/</span><span class=\"nt\">script</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;/</span><span class=\"nt\">head</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;</span><span class=\"nt\">body</span> <span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;</span><span class=\"nt\">nav</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>HTMX
  Blog<span class=\"p\">&lt;/</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;</span><span class=\"nt\">div</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;navbar&quot;</span><span class=\"p\">&gt;</span>\n
  \         {% if user.is_authenticated %}\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;nav-item nav-link&quot;</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% url &#39;logout&#39; %}&quot;</span><span
  class=\"p\">&gt;&lt;</span><span class=\"nt\">button</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;btn btn-link&quot;</span><span class=\"p\">&gt;</span>Logout<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n          {% else %}\n            <span
  class=\"p\">&lt;</span><span class=\"nt\">a</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;nav-item nav-link&quot;</span> <span
  class=\"na\">href</span><span class=\"o\">=</span><span class=\"s\">&quot;{% url
  &#39;login&#39; %}&quot;</span><span class=\"p\">&gt;&lt;</span><span class=\"nt\">button</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;btn
  btn-link&quot;</span><span class=\"p\">&gt;</span>Login<span class=\"p\">&lt;/</span><span
  class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span class=\"nt\">a</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;nav-item
  nav-link&quot;</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% url &#39;register&#39; %}&quot;</span><span class=\"p\">&gt;&lt;</span><span
  class=\"nt\">button</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;btn btn-link&quot;</span><span class=\"p\">&gt;</span>Register<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n          {% endif %}\n        <span
  class=\"p\">&lt;/</span><span class=\"nt\">div</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;/</span><span class=\"nt\">nav</span><span class=\"p\">&gt;</span>\n\n
  \   {% block body %}\n    {% endblock %}\n<span class=\"p\">&lt;/</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>I have a nav bar with my user
  authentication views, simply a login or signup button if the user is not logged
  in and a log-out button if the user is authenticated. We have added the htmx script
  file from the CDN just before the end of the head tag. We also have included the
  bootstrap CSS file for a decent UI which we will be creating in this post.</p>\n<p>That
  is one of the ways, htmx can be injected into an HTML template, you can even download
  the javascript file from the <a href=\"https://unpkg.com/browse/htmx.org/dist/\">htmx
  cdn</a>. Further, this can be downloaded or pasted into your local folder and served
  as a static file or embedded directly into an HTML template.</p>\n<h2>Defining Models</h2>\n<p>We
  will start the tutorial by defining the model of the application we are creating.
  Here, we will create a simple Article model with a few parameters like <code>title</code>,
  <code>content</code>, <code>author</code>, etc.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">from</span> <span class=\"nn\">django.db</span> <span class=\"kn\">import</span>
  <span class=\"n\">models</span>\n<span class=\"kn\">from</span> <span class=\"nn\">user.models</span>
  <span class=\"kn\">import</span> <span class=\"n\">Profile</span>\n\n<span class=\"k\">class</span>
  <span class=\"nc\">Article</span><span class=\"p\">(</span><span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">Model</span><span class=\"p\">):</span>\n
  \   <span class=\"n\">Article_Status</span> <span class=\"o\">=</span> <span class=\"p\">(</span>\n
  \       <span class=\"p\">(</span><span class=\"s2\">&quot;DRAFT&quot;</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;Draft&quot;</span><span class=\"p\">),</span>\n
  \       <span class=\"p\">(</span><span class=\"s2\">&quot;PUBLISHED&quot;</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;Published&quot;</span><span class=\"p\">),</span>\n
  \   <span class=\"p\">)</span>\n    <span class=\"n\">title</span> <span class=\"o\">=</span>
  <span class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">CharField</span><span
  class=\"p\">(</span><span class=\"n\">max_length</span><span class=\"o\">=</span><span
  class=\"mi\">128</span><span class=\"p\">,</span> <span class=\"n\">unique</span><span
  class=\"o\">=</span><span class=\"kc\">True</span><span class=\"p\">)</span>\n    <span
  class=\"n\">content</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">TextField</span><span class=\"p\">()</span>\n
  \   <span class=\"n\">author</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">ForeignKey</span><span class=\"p\">(</span><span
  class=\"n\">Profile</span><span class=\"p\">,</span> <span class=\"n\">on_delete</span><span
  class=\"o\">=</span><span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">CASCADE</span><span class=\"p\">)</span>\n    <span class=\"n\">status</span>
  <span class=\"o\">=</span> <span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">CharField</span><span class=\"p\">(</span>\n        <span class=\"n\">max_length</span><span
  class=\"o\">=</span><span class=\"mi\">16</span><span class=\"p\">,</span>\n        <span
  class=\"n\">choices</span><span class=\"o\">=</span><span class=\"n\">Article_Status</span><span
  class=\"p\">,</span>\n        <span class=\"n\">default</span><span class=\"o\">=</span><span
  class=\"n\">Article_Status</span><span class=\"p\">[</span><span class=\"mi\">0</span><span
  class=\"p\">],</span>\n    <span class=\"p\">)</span>\n\n    <span class=\"k\">def</span>
  <span class=\"fm\">__str__</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"p\">):</span>\n        <span class=\"k\">return</span> <span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">title</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above model <code>Article</code>, we have a few fields like <code>title</code>
  simple Character Field, <code>content</code> as a text field as it will be a large
  text as the post body, <code>author</code> which is a ForeignKey to the User Model.
  We also have the status, which is defined as a character field but with a few choices
  like <code>draft</code> or <code>published</code>, we can further modify this status
  as public or private. But just keeping it simple and easy to understand.</p>\n<p>The
  object reference name for this model is the title as we have defined in the dunder
  string method. So, that is a simple model created, we can now migrate the changes
  into the database for adding the tables and attributes.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>python manage.py makemigrations\npython
  manage.py migrate\n</pre></div>\n\n</pre>\n\n<p>This will make migrations to the
  database i.e. convert the python model class into database tables and attributes.
  So, once the migration process is completed successfully, we can move into the crust
  of this article which is to actually design the views. In the next section, we will
  be utilizing the models in our views for representing the data on the templates.</p>\n<h2>Creating
  Article Form</h2>\n<p>Before diving into the views section, we need a few things
  like the Article Form, which will be a Django Model-based form. It will help us
  a lot in creating or updating the fields for the article model. We can define a
  form in a python file called <code>forms.py</code>, it's not necessary to keep your
  forms in the <code>forms.py</code> but if you have a lot of forms and models, it
  becomes a good practice to organize the components of our app. So, I'll be creating
  a new file inside of the <code>article</code> app called <code>forms.py</code> and
  defining the <code>ArticleForm</code>.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># article/forms.py</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">django</span>
  <span class=\"kn\">import</span> <span class=\"n\">forms</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">.models</span> <span class=\"kn\">import</span> <span class=\"n\">Article</span>\n\n\n<span
  class=\"k\">class</span> <span class=\"nc\">ArticleForm</span><span class=\"p\">(</span><span
  class=\"n\">forms</span><span class=\"o\">.</span><span class=\"n\">ModelForm</span><span
  class=\"p\">):</span>\n    <span class=\"k\">class</span> <span class=\"nc\">Meta</span><span
  class=\"p\">:</span>\n        <span class=\"n\">model</span> <span class=\"o\">=</span>
  <span class=\"n\">Article</span>\n        <span class=\"n\">exclude</span> <span
  class=\"o\">=</span> <span class=\"p\">(</span>\n            <span class=\"s2\">&quot;created&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;updated&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"s2\">&quot;author&quot;</span><span
  class=\"p\">,</span>\n        <span class=\"p\">)</span>\n        <span class=\"n\">widgets</span>
  <span class=\"o\">=</span> <span class=\"p\">{</span>\n            <span class=\"s2\">&quot;title&quot;</span><span
  class=\"p\">:</span> <span class=\"n\">forms</span><span class=\"o\">.</span><span
  class=\"n\">TextInput</span><span class=\"p\">(</span>\n                <span class=\"n\">attrs</span><span
  class=\"o\">=</span><span class=\"p\">{</span>\n                    <span class=\"s2\">&quot;class&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;form-control&quot;</span><span class=\"p\">,</span>\n
  \                   <span class=\"s2\">&quot;style&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;max-width: 450px; align: center;&quot;</span><span class=\"p\">,</span>\n
  \                   <span class=\"s2\">&quot;placeholder&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;Title&quot;</span><span class=\"p\">,</span>\n                <span
  class=\"p\">}</span>\n            <span class=\"p\">),</span>\n            <span
  class=\"s2\">&quot;content&quot;</span><span class=\"p\">:</span> <span class=\"n\">forms</span><span
  class=\"o\">.</span><span class=\"n\">Textarea</span><span class=\"p\">(</span>\n
  \               <span class=\"n\">attrs</span><span class=\"o\">=</span><span class=\"p\">{</span>\n
  \                   <span class=\"s2\">&quot;class&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;form-control&quot;</span><span class=\"p\">,</span>\n                    <span
  class=\"s2\">&quot;style&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;max-width:
  900px;&quot;</span><span class=\"p\">,</span>\n                    <span class=\"s2\">&quot;placeholder&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Content&quot;</span><span class=\"p\">,</span>\n
  \               <span class=\"p\">}</span>\n            <span class=\"p\">),</span>\n
  \       <span class=\"p\">}</span>\n</pre></div>\n\n</pre>\n\n<p>So, the forms are
  inherited from the [ModelForm] which allows us to create forms based on our model.
  So, we specify the model name which in our case is <code>Article</code> and further
  we can have <code>exclude</code> or <code>fields</code> tuples. To exclude certain
  fields in the actual form, just parse the tuple of those attributes and if you want
  to only select a few attributes, you can specify the <code>fields</code> tuple and
  mention the required fields for the form.</p>\n<p>So, if we have a lot of things
  to be included in the form, we can specify only the attributes to be excluded with
  the <code>exclude</code> tuple. And if we have a lot of fields to be excluded, we
  can use the <code>fields</code> tuple to specify which attributes to use in the
  form.</p>\n<p>Let's take an example: For the above ArticleForm, if we wanted to
  specify the required fields to be included in the form, then we might use the <code>fields</code>
  tuple like below the rest will be not rendered in the form fields.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>class ArticleForm(forms.ModelForm):\n
  \   class Meta:\n        model = Article\n        fields = (\n            &quot;title&quot;,\n
  \           &quot;content&quot;,\n            &quot;status&quot;,\n        )\n</pre></div>\n\n</pre>\n\n<p>Both
  of them can be used, it just depends on how many fields you have to exclude or include
  in the rendered form.</p>\n<p>We have also specified the <code>widgets</code> attribute
  which gives a bit more control on how we need to display the form in a template.
  So I have specified the type of input it needs to render like a simple text input
  for the title, text area for content, etc. The cool thing about this is it can automatically
  set these by knowing the type of field in the model, but sometimes it can be a bit
  undesired mostly with complex relationships and attributes.</p>\n<h2>Creating Views</h2>\n<p>Let's
  start creating views for creating, reading, updating, and deleting articles from
  the database. I will be using function-based views just because we are understanding
  the flow of how HTMX and Django can be integrated so we need to dive in deeper and
  understand the actual flow of the process.</p>\n<h3>Create View</h3>\n<p>So, creating
  articles seems like a good way to start off. We can create a simple function-based
  view which will initially load in an empty <code>ArticleForm</code> and if the request
  is <code>GET</code> we will render the form in the <code>create.html</code> template.
  If the request is <code>POST</code> which will be after we submit the form, we will
  validate the form and attach the current user as the author of the article and save
  the for instance which will create an article record and this object will be rendered
  to the detail template.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"kn\">from</span> <span class=\"nn\">django.shortcuts</span> <span class=\"kn\">import</span>
  <span class=\"n\">render</span>\n<span class=\"kn\">from</span> <span class=\"nn\">.models</span>
  <span class=\"kn\">import</span> <span class=\"n\">Article</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">.forms</span> <span class=\"kn\">import</span> <span class=\"n\">ArticleForm</span>\n\n<span
  class=\"k\">def</span> <span class=\"nf\">createArticle</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">):</span>\n    <span class=\"n\">form</span>
  <span class=\"o\">=</span> <span class=\"n\">ArticleForm</span><span class=\"p\">()</span>\n
  \   <span class=\"n\">context</span> <span class=\"o\">=</span> <span class=\"p\">{</span>\n
  \       <span class=\"s1\">&#39;form&#39;</span><span class=\"p\">:</span> <span
  class=\"n\">form</span><span class=\"p\">,</span>\n    <span class=\"p\">}</span>\n
  \   <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;articles/create.html&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<h4>Rendering
  the Form</h4>\n<p>We are creating an empty instance of <code>ArticleForm</code>
  and rendering it in the template. So, this will render the empty form in the <code>create.html</code>
  template.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cm\">&lt;!-- templates/articles/create.html --&gt;</span>\n\n{% extends
  &#39;base.html&#39; %}\n\n{% block body %}\n<span class=\"p\">&lt;</span><span class=\"nt\">div</span>
  <span class=\"na\">hx-target</span><span class=\"o\">=</span><span class=\"s\">&quot;this&quot;</span>
  <span class=\"na\">hx-swap</span><span class=\"o\">=</span><span class=\"s\">&quot;outerHTML&quot;</span><span
  class=\"p\">&gt;</span>\n  <span class=\"p\">&lt;</span><span class=\"nt\">form</span><span
  class=\"p\">&gt;</span>\n    {% csrf_token %}\n    {{ form.as_p }}\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">button</span> <span class=\"na\">hx-post</span><span class=\"o\">=</span><span
  class=\"s\">&quot;.&quot;</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;btn btn-success&quot;</span>\n      <span class=\"na\">type</span><span
  class=\"o\">=</span><span class=\"s\">&quot;submit&quot;</span><span class=\"p\">&gt;</span>Save<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;</span>\n
  \ <span class=\"p\">&lt;/</span><span class=\"nt\">form</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;/</span><span class=\"nt\">div</span><span class=\"p\">&gt;</span>\n{%
  endblock %}\n</pre></div>\n\n</pre>\n\n<p>Now, here we are inheriting from the base
  template and creating a form tag in HTML with the <code>{{ form}}</code> for rendering
  the form fields, we finally have the <code>button</code> element for submitting
  the form. We have used the <code>hx-post</code> attribute. More on this in just
  a minute. So, this is we create a template for rendering the article form.</p>\n<p>We
  have used the <code>hx-post</code> attribute here, which will send a <code>POST</code>
  request to the current <code>URL</code> represented by <code>hx-post=&quot;.&quot;</code>.
  You might have noticed the <code>div</code> attributes, the <code>hx-target</code>
  and <code>hx-swap</code>, so these are some of the many attributes provided by the
  htmx library for controlling the reactivity of the requests made. The <code>hx-target</code>
  allow us to specify the element or tag to which the data will be rendered. The <code>hx-swap</code>
  goes hand-in-hand for specifying the target DOM like <code>innerHTML</code>, <code>outerHTML</code>,
  etc. You can see the various options on the <a href=\"https://htmx.org/docs/#swapping\">htmx
  docs</a>. By specifying the <code>hx-swap</code> as  <code>outerHTML</code>, we
  are saying to replace the entire element with the incoming content from the request
  which we will send with nearby request triggers.</p>\n<p>We need to map the view
  to a URL in order to get a good idea about the request and parsed content.</p>\n<p>We'll
  create a <code>create/</code> route and bind it to the <code>createArticle</code>
  view with the name <code>article-create</code>.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># article/urls.py</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">django.urls</span>
  <span class=\"kn\">import</span> <span class=\"n\">path</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">.</span> <span class=\"kn\">import</span> <span class=\"n\">views</span>\n\n<span
  class=\"n\">urlpatterns</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;create/&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">createArticle</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;article-create&#39;</span><span class=\"p\">),</span>
  \n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>This URL will be mapped
  to the global URL in the project, here we can simply specify the prefix for the
  URLs in the <code>article</code> app and include those URLs.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># htmx_blog/urls.py</span>\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.contrib</span> <span class=\"kn\">import</span> <span class=\"n\">admin</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">django.urls</span> <span class=\"kn\">import</span>
  <span class=\"n\">path</span><span class=\"p\">,</span> <span class=\"n\">include</span>\n\n<span
  class=\"n\">urlpatterns</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;admin/&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">admin</span><span class=\"o\">.</span><span
  class=\"n\">site</span><span class=\"o\">.</span><span class=\"n\">urls</span><span
  class=\"p\">),</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;user/&#39;</span><span class=\"p\">,</span> <span class=\"n\">include</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;user.urls&#39;</span><span class=\"p\">),</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;auth&#39;</span><span
  class=\"p\">),</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span> <span class=\"n\">include</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;article.urls&#39;</span><span class=\"p\">),</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;home&#39;</span><span
  class=\"p\">),</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>Feel,
  free to add any other URL pattern like for instance, the article app is at <code>/</code>
  i.e. <code>127.0.01.:8000/</code>, you can add any other name like <code>127.0.0.1:8000/article/</code>
  by adding <code>path('article/', include('article.urls'))</code>.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659252089/blog-media/django-htmx-create-view.png\"
  alt=\"Django HTMX Create view Form Template\" /></p>\n<p>So, finally, we are sending
  a <code>GET</code> request to the <code>127.0.0.1:8000/create/</code> and this will
  output the form. As we have a <code>POST</code> request embedded in the button inside
  the form, we will send the <code>POST</code> request to the same URL -&gt; <code>127.0.0.1:8000/create/</code>.</p>\n<h4>Submitting
  the Form</h4>\n<p>Let's handle the <code>POST</code> request in the create view.</p>\n<pre
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
  class=\"kn\">from</span> <span class=\"nn\">django.shortcuts</span> <span class=\"kn\">import</span>
  <span class=\"n\">render</span>\n<span class=\"kn\">from</span> <span class=\"nn\">.models</span>
  <span class=\"kn\">import</span> <span class=\"n\">Article</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">.forms</span> <span class=\"kn\">import</span> <span class=\"n\">ArticleForm</span>\n\n<span
  class=\"k\">def</span> <span class=\"nf\">createArticle</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">):</span>\n    <span class=\"n\">form</span>
  <span class=\"o\">=</span> <span class=\"n\">ArticleForm</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"o\">.</span><span class=\"n\">POST</span>
  <span class=\"ow\">or</span> <span class=\"kc\">None</span><span class=\"p\">)</span>\n
  \   <span class=\"k\">if</span> <span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">method</span> <span class=\"o\">==</span> <span class=\"s1\">&#39;POST&#39;</span><span
  class=\"p\">:</span>\n        <span class=\"k\">if</span> <span class=\"n\">form</span><span
  class=\"o\">.</span><span class=\"n\">is_valid</span><span class=\"p\">():</span>\n
  \           <span class=\"n\">form</span><span class=\"o\">.</span><span class=\"n\">instance</span><span
  class=\"o\">.</span><span class=\"n\">author</span> <span class=\"o\">=</span> <span
  class=\"n\">request</span><span class=\"o\">.</span><span class=\"n\">user</span>\n
  \           <span class=\"n\">article</span> <span class=\"o\">=</span> <span class=\"n\">form</span><span
  class=\"o\">.</span><span class=\"n\">save</span><span class=\"p\">()</span>\n            <span
  class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;articles/detail.html&#39;</span><span
  class=\"p\">,</span> <span class=\"p\">{</span><span class=\"s1\">&#39;article&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">article</span><span class=\"p\">})</span>\n
  \   <span class=\"n\">context</span> <span class=\"o\">=</span> <span class=\"p\">{</span>\n
  \       <span class=\"s1\">&#39;form&#39;</span><span class=\"p\">:</span> <span
  class=\"n\">form</span><span class=\"p\">,</span>\n    <span class=\"p\">}</span>\n
  \   <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;articles/create.html&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p><strong>Simple
  explanation</strong></p>\n<ul>\n<li>Create a form instance of ArticleForm with the
  request data or empty -&gt; <code>ArticleForm(request.POST or None)</code></li>\n<li>If
  it's a POST request, validate and create the article, render the article object
  in <code>detail.html</code> template.</li>\n<li>If it's a GET request, render the
  empty form in <code>create.html</code></li>\n</ul>\n<p>There are a few changes in
  the view, instead of initializing the form to empty i.e. <code>ArticleForm()</code>,
  we are initializing with <code>ArticleForm(request.POST or None)</code>. This basically
  means that if we are having something in the <code>request.POST</code> dict, we
  will initialize the Form with that data or else an empty form instance.</p>\n<p>Next,
  we check if the request if <code>POST</code>, if it is then we check if the form
  is valid i.e. the form fields are not empty or if any other constraint on the model
  attributes is satisfied or not. If the form data is valid, we attach the author
  as the currently logged-in User/user who sent the request. Finally, we save the
  form which in turn creates the article record in the database. We then render the
  created article in the <code>detail.html</code> template which is not yet created.</p>\n<p>So,
  the <code>htmx-post</code> attribute has worked and it will send a post request
  to the same URL i.e. <code>127.0.0.1:8000/create</code> and this will again trigger
  the view <code>createArticle</code> this time we will have <code>request.POST</code>
  data. So, we will validate and save the form.</p>\n<h3>Detail View</h3>\n<p>The
  detail view is used for viewing the details of an article. This will be rendered
  after the article has been created or updated. This is quite simple, we need an
  <code>id</code> or <code>primary key(pk)</code> of an article and render the <code>title</code>
  and <code>content</code> of the article in the template.</p>\n<p>We pass in a primary
  key along with the request as a parameter to the view, the <code>pk</code> will
  be passed via the URL. We fetch the Article object with the id as the parsed <code>pk</code>
  and finally render the <code>detail.html</code> template with the article object.
  The <code>context['article']</code> can be accessed from the template to render
  the specific attributes like <code>title</code>, <code>content</code>, etc.</p>\n<pre
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
  class=\"c1\"># article/views.py</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">detailArticle</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"n\">pk</span><span class=\"p\">):</span>\n    <span class=\"n\">article</span>
  <span class=\"o\">=</span> <span class=\"n\">Article</span><span class=\"o\">.</span><span
  class=\"n\">objects</span><span class=\"o\">.</span><span class=\"n\">get</span><span
  class=\"p\">(</span><span class=\"nb\">id</span><span class=\"o\">=</span><span
  class=\"n\">pk</span><span class=\"p\">)</span>\n    <span class=\"n\">context</span>
  <span class=\"o\">=</span> <span class=\"p\">{</span><span class=\"s1\">&#39;article&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">article</span><span class=\"p\">}</span>\n
  \   <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;articles/detail.html&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>We
  can now bind the view to a URL and parse the required parameter <code>pk</code>
  to the view.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"kn\">from</span> <span class=\"nn\">django.urls</span> <span class=\"kn\">import</span>
  <span class=\"n\">path</span>\n<span class=\"kn\">from</span> <span class=\"nn\">.</span>
  <span class=\"kn\">import</span> <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;create/&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">createArticle</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s1\">&#39;article-create&#39;</span><span class=\"p\">),</span> \n    <span
  class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;&lt;int:pk&gt;&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">detailArticle</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;article-detail&#39;</span><span class=\"p\">),</span>
  \n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>We have parsed the <code>pk</code>
  as <code>int</code> to the URL parameter, so for an article with id=4, the URL will
  be, <code>127.0.0.1:8000/4/</code>.</p>\n<p>We need to create the template for rendering
  the context from the <code>detailArticle</code> view. So, we create the <code>detail.html</code>
  in the <code>templates/articles</code> folder. We inherit the base template and
  render the <code>article.title</code> and the <code>article.content</code> with
  a linebreaks template filter so as to display the content properly.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cm\">&lt;!-- templates/articles/detail.html --&gt;</span>\n\n\n{% extends
  &#39;base.html&#39; %}\n{% block body %}\n<span class=\"p\">&lt;</span><span class=\"nt\">div</span>
  <span class=\"na\">id</span><span class=\"o\">=</span><span class=\"s\">&quot;article-card&quot;</span><span
  class=\"p\">&gt;</span>\n  <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span
  class=\"p\">&gt;</span>{{ article.title }}\n  <span class=\"p\">&lt;</span><span
  class=\"nt\">p</span><span class=\"p\">&gt;</span>{{ article.content|linebreaks|safe
  }}<span class=\"p\">&lt;/</span><span class=\"nt\">p</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;</span><span class=\"nt\">div</span><span class=\"p\">&gt;</span>\n{%
  endblock %}\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659252227/blog-media/django-htmx-detail-view.png\"
  alt=\"Detail View Template\" /></p>\n<p>So, we can now use <code>createArticle</code>
  view as well as <code>detailArticle</code> view, this both are configured properly,
  so (CR) or CRUD is completed. We can add <code>listArticle</code> for listing out
  all the author's(logged-in user) articles.</p>\n<h3>List View</h3>\n<p>Listview
  of the articles is much similar to the detail view as it will return a list of articles
  rather than a single article.</p>\n<p>So in the <code>listArticle</code> view, we
  will return all the articles with the author as the user who sent the request/logged-in
  user. We will parse this object list into the template as <code>base.html</code>
  or <code>articles/list.html</code>.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># article/views.py</span>\n\n\n<span class=\"k\">def</span> <span class=\"nf\">listArticle</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"n\">articles</span> <span class=\"o\">=</span> <span class=\"n\">Article</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">filter</span><span class=\"p\">(</span><span class=\"n\">author</span><span
  class=\"o\">=</span><span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">user</span><span class=\"o\">.</span><span class=\"n\">id</span><span
  class=\"p\">)</span>\n    <span class=\"n\">context</span> <span class=\"o\">=</span>
  <span class=\"p\">{</span>\n        <span class=\"s1\">&#39;articles&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">articles</span><span class=\"p\">,</span>\n
  \   <span class=\"p\">}</span>\n    <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;base.html&#39;</span><span class=\"p\">,</span> <span class=\"n\">context</span><span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>We will add the URL route for
  this as the <code>/</code> route that is on <code>127.0.0.1:8000/</code> this is
  the base URL for the article app and is the route for the <code>listArticle</code>
  view. So, we will display the list of articles on the homepage.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># article/urls.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.urls</span> <span class=\"kn\">import</span> <span class=\"n\">path</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">.</span> <span class=\"kn\">import</span>
  <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span> <span class=\"o\">=</span>
  <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&lt;int:pk&gt;&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">detailArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-detail&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;create/&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">createArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-create&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">listArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-list&#39;</span><span
  class=\"p\">),</span> \n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>Let's
  create the template for the list view which will iterate over the articles and display
  the relevant data like the title and link to the article.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cm\">&lt;!-- templates/articles/list.html --&gt;</span>\n\n<span class=\"p\">&lt;</span><span
  class=\"nt\">ul</span> <span class=\"na\">id</span><span class=\"o\">=</span><span
  class=\"s\">&quot;article-list&quot;</span><span class=\"p\">&gt;</span>\n  {% for
  article in articles %}\n  <span class=\"p\">&lt;</span><span class=\"nt\">li</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">div</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;card&quot;</span>
  <span class=\"na\">style</span><span class=\"o\">=</span><span class=\"s\">&quot;width:
  18rem;&quot;</span><span class=\"p\">&gt;</span>\n      <span class=\"p\">&lt;</span><span
  class=\"nt\">div</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;card-body&quot;</span><span class=\"p\">&gt;</span>\n        <span
  class=\"p\">&lt;</span><span class=\"nt\">h5</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;card-title&quot;</span><span class=\"p\">&gt;</span>{{
  article.title }}<span class=\"p\">&lt;/</span><span class=\"nt\">h5</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">p</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;card-text&quot;</span><span
  class=\"p\">&gt;</span>{{ article.content|truncatewords:5  }}<span class=\"p\">&lt;/</span><span
  class=\"nt\">p</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% url &#39;article-detail&#39; article.id %}&quot;</span> <span
  class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;card-link&quot;</span><span
  class=\"p\">&gt;</span>Read more<span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n      <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n  <span class=\"p\">&lt;/</span><span class=\"nt\">li</span><span
  class=\"p\">&gt;</span>\n  {% endfor %}\n<span class=\"p\">&lt;/</span><span class=\"nt\">ul</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>We have used the <code>truncatewords:5</code>
  template filter for only displaying the content of the articles till the first 5
  words as it is just a list view, we don't want to display every detail of the article
  here.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659252293/blog-media/django-htmx-list-view.png\"
  alt=\"List view Template\" /></p>\n<p>We can use this template to render in the
  <code>base.html</code> file.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"cp\">&lt;!DOCTYPE html&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">html</span>
  <span class=\"na\">lang</span><span class=\"o\">=</span><span class=\"s\">&quot;en&quot;</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">meta</span>
  <span class=\"na\">charset</span><span class=\"o\">=</span><span class=\"s\">&quot;UTF-8&quot;</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>HTMX Blog<span class=\"p\">&lt;/</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>\n    {% load static %}\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">link</span> <span class=\"na\">rel</span><span class=\"o\">=</span><span
  class=\"s\">&quot;stylesheet&quot;</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css&quot;</span>
  <span class=\"na\">integrity</span><span class=\"o\">=</span><span class=\"s\">&quot;sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm&quot;</span>
  <span class=\"na\">crossorigin</span><span class=\"o\">=</span><span class=\"s\">&quot;anonymous&quot;</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">script</span>
  <span class=\"na\">src</span><span class=\"o\">=</span><span class=\"s\">&quot;https://unpkg.com/htmx.org@1.8.0&quot;</span><span
  class=\"p\">&gt;&lt;/</span><span class=\"nt\">script</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;/</span><span class=\"nt\">head</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;</span><span class=\"nt\">body</span> <span class=\"na\">hx-target</span><span
  class=\"o\">=</span><span class=\"s\">&quot;this&quot;</span> <span class=\"na\">hx-swap</span><span
  class=\"o\">=</span><span class=\"s\">&quot;outerHTML&quot;</span> <span class=\"na\">hx-headers</span><span
  class=\"o\">=</span><span class=\"s\">&#39;{&quot;X-CSRFToken&quot;: &quot;{{ csrf_token
  }}&quot;}&#39;</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">nav</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>HTMX Blog<span class=\"p\">&lt;/</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">div</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;navbar&quot;</span><span class=\"p\">&gt;</span>\n          {%
  if user.is_authenticated %}\n            <span class=\"p\">&lt;</span><span class=\"nt\">a</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;nav-item
  nav-link&quot;</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% url &#39;article-list&#39; %}&quot;</span><span class=\"p\">&gt;&lt;</span><span
  class=\"nt\">button</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;btn btn-link&quot;</span><span class=\"p\">&gt;</span>Home<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;nav-item nav-link&quot;</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% url &#39;logout&#39; %}&quot;</span><span
  class=\"p\">&gt;&lt;</span><span class=\"nt\">button</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;btn btn-link&quot;</span><span class=\"p\">&gt;</span>Logout<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n          {% else %}\n            <span
  class=\"p\">&lt;</span><span class=\"nt\">a</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;nav-item nav-link&quot;</span> <span
  class=\"na\">href</span><span class=\"o\">=</span><span class=\"s\">&quot;{% url
  &#39;login&#39; %}&quot;</span><span class=\"p\">&gt;&lt;</span><span class=\"nt\">button</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;btn
  btn-link&quot;</span><span class=\"p\">&gt;</span>Login<span class=\"p\">&lt;/</span><span
  class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n            <span class=\"p\">&lt;</span><span class=\"nt\">a</span>
  <span class=\"na\">class</span><span class=\"o\">=</span><span class=\"s\">&quot;nav-item
  nav-link&quot;</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% url &#39;register&#39; %}&quot;</span><span class=\"p\">&gt;&lt;</span><span
  class=\"nt\">button</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;btn btn-link&quot;</span><span class=\"p\">&gt;</span>Register<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n          {% endif %}\n        <span
  class=\"p\">&lt;/</span><span class=\"nt\">div</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;/</span><span class=\"nt\">nav</span><span class=\"p\">&gt;</span>\n\n
  \   {% block body %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">a</span>
  <span class=\"na\">href</span><span class=\"o\">=</span><span class=\"s\">&quot;{%
  url &#39;article-create&#39; %}&quot;</span><span class=\"p\">&gt;&lt;</span><span
  class=\"nt\">button</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;btn btn-success&quot;</span> <span class=\"p\">&gt;</span>Create<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;&lt;/</span><span
  class=\"nt\">a</span><span class=\"p\">&gt;</span>\n    {% include &#39;articles/list.html&#39;
  %}\n    {% endblock %}\n<span class=\"p\">&lt;/</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>We have now included the <code>list.html</code>
  template on the homepage and also added the <code>create</code> button as the link
  to the <code>article-create</code> URL.</p>\n<h3>Delete View</h3>\n<p>For deleting
  an article, we will simply rely on htmx for sending the request and on that request,
  we will delete the current article and render the updated list of articles.</p>\n<p>With
  the <code>deleteArticle</code> view, we will take in two parameters the request
  which is by default for a Django function-based view, and the primary key as <code>pk</code>.
  Again we will parse the <code>pk</code> from the URL. We will delete the article
  object and get the latest list of articles. Finally, render the updated list of
  articles in the base template which is our list view.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># article/views.py</span>\n\n\n<span class=\"k\">def</span> <span class=\"nf\">deleteArticle</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"n\">pk</span><span class=\"p\">):</span>\n    <span class=\"n\">Article</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">get</span><span class=\"p\">(</span><span class=\"nb\">id</span><span
  class=\"o\">=</span><span class=\"n\">pk</span><span class=\"p\">)</span><span class=\"o\">.</span><span
  class=\"n\">delete</span><span class=\"p\">()</span>\n    <span class=\"n\">articles</span>
  <span class=\"o\">=</span> <span class=\"n\">Article</span><span class=\"o\">.</span><span
  class=\"n\">objects</span><span class=\"o\">.</span><span class=\"n\">filter</span><span
  class=\"p\">(</span><span class=\"n\">author</span><span class=\"o\">=</span><span
  class=\"n\">request</span><span class=\"o\">.</span><span class=\"n\">user</span><span
  class=\"p\">)</span>\n    <span class=\"n\">context</span> <span class=\"o\">=</span>
  <span class=\"p\">{</span><span class=\"s1\">&#39;article&#39;</span><span class=\"p\">:</span>
  <span class=\"n\">articles</span><span class=\"p\">}</span>\n    <span class=\"k\">return</span>
  <span class=\"n\">render</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;base.html&quot;</span><span class=\"p\">,</span>
  <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>We
  will add the <code>deleteArticle</code> into the URL patterns and call it <code>article-delete</code>
  with the URL of <code>delete/&lt;int:pk&gt;</code>. This will allow us to send a
  request to the URL <code>127.0.0.1:8000/delete/4</code> for deleting the article
  with id <code>4</code>.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># article/urls.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.urls</span> <span class=\"kn\">import</span> <span class=\"n\">path</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">.</span> <span class=\"kn\">import</span>
  <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span> <span class=\"o\">=</span>
  <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">listArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-list&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&lt;int:pk&gt;&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">detailArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-detail&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;create/&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">createArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-create&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;delete/&lt;int:pk&gt;&#39;</span><span class=\"p\">,</span> <span
  class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">deleteArticle</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s1\">&#39;article-delete&#39;</span><span class=\"p\">),</span> \n<span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>In the delete view, the template
  is important as we want to send a request appropriately to the defined URL. To do
  that, we will have a form but it won't have any inputs as such just a button that
  indicates to delete the current article. We will add the <code>hx-delete</code>
  attribute as the URL to the <code>deleteArticle</code> view. with the id of the
  article. This will send a request to the <code>article-delete</code> URL which will,
  in turn, trigger the view with the given id and delete the article.</p>\n<p>We have
  added the <code>hx-confirm</code> attribute for showing a pop-up of confirmation
  of deleting the article. As you can see we have added a little script for adding
  <code>csrf_token</code> into the HTML, this is important in order to submit a form
  with a valid <code>CSRFToken</code>.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"cm\">&lt;!-- templates/article/delete.html --&gt;</span>\n\n<span class=\"p\">&lt;</span><span
  class=\"nt\">script</span><span class=\"p\">&gt;</span>\n  <span class=\"nb\">document</span><span
  class=\"p\">.</span><span class=\"nx\">body</span><span class=\"p\">.</span><span
  class=\"nx\">addEventListener</span><span class=\"p\">(</span><span class=\"s1\">&#39;htmx:configRequest&#39;</span><span
  class=\"p\">,</span> <span class=\"p\">(</span><span class=\"nx\">event</span><span
  class=\"p\">)</span> <span class=\"p\">=&gt;</span> <span class=\"p\">{</span>\n
  \   <span class=\"nx\">event</span><span class=\"p\">.</span><span class=\"nx\">detail</span><span
  class=\"p\">.</span><span class=\"nx\">headers</span><span class=\"p\">[</span><span
  class=\"s1\">&#39;X-CSRFToken&#39;</span><span class=\"p\">]</span> <span class=\"o\">=</span>
  <span class=\"s1\">&#39;{{ csrf_token }}&#39;</span><span class=\"p\">;</span>\n
  \ <span class=\"p\">})</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">script</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">div</span>
  <span class=\"p\">&gt;</span>\n  <span class=\"p\">&lt;</span><span class=\"nt\">form</span>
  <span class=\"na\">method</span><span class=\"o\">=</span><span class=\"s\">&quot;post&quot;</span>
  <span class=\"p\">&gt;</span>\n  {% csrf_token %}\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">button</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;btn btn-danger&quot;</span>\n      <span class=\"na\">hx-delete</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% url &#39;article-delete&#39; article.id
  %}&quot;</span>\n      <span class=\"na\">hx-confirm</span><span class=\"o\">=</span><span
  class=\"s\">&quot;Are you sure, You want to delete this article?&quot;</span>\n
  \     <span class=\"na\">type</span><span class=\"o\">=</span><span class=\"s\">&quot;submit&quot;</span><span
  class=\"p\">&gt;</span>\n      Delete\n    <span class=\"p\">&lt;/</span><span class=\"nt\">button</span><span
  class=\"p\">&gt;</span>\n  <span class=\"p\">&lt;/</span><span class=\"nt\">form</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>Do you have a question like
  how do we access the <code>article.id</code>? we are not rendering the <code>delete.html</code>
  template from the view, so there is no context to pass. We will include this snippet
  into the detail view template, so as to have the option of deleting the current
  article.</p>\n<p>We will modify the <code>articles/detail.html</code> template and
  include the <code>delete.html</code> template. This includes simply adding an HTML
  template in the specified location. So, we will basically inject the delete form
  into the detail template.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>{%
  extends &#39;base.html&#39; %}\n{% block body %}\n<span class=\"p\">&lt;</span><span
  class=\"nt\">div</span> <span class=\"na\">hx-target</span><span class=\"o\">=</span><span
  class=\"s\">&quot;this&quot;</span> <span class=\"na\">hx-swap</span><span class=\"o\">=</span><span
  class=\"s\">&quot;outerHTML&quot;</span><span class=\"p\">&gt;</span>\n  <span class=\"p\">&lt;</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>{{ article.title }}\n  {% include
  &#39;articles/delete.html&#39; %}\n  <span class=\"p\">&lt;</span><span class=\"nt\">p</span><span
  class=\"p\">&gt;</span>{{ article.content|linebreaks|safe }}<span class=\"p\">&lt;/</span><span
  class=\"nt\">p</span><span class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span
  class=\"nt\">div</span><span class=\"p\">&gt;</span>\n{% endblock %}\n</pre></div>\n\n</pre>\n\n<p>Hence,
  we will have a nice option to delete the article in the detail section, this can
  be placed anywhere but remember, we need to add the <code>hx-target=&quot;this&quot;</code>
  and <code>hx-swap=&quot;outerHTML&quot;</code> in the div so as to correctly swap
  the HTML content after the request has been made.</p>\n<h3>Update View</h3>\n<p>We
  can now move into the final piece of the CRUD i.e. <code>Update</code>. This will
  be similar to the <code>createArticle</code> with a couple of changes. We will parse
  parameters like <code>pk</code> to this view as well because we want to update a
  specific article. So, we will have to get the primary key of the article from the
  URL slug.</p>\n<p>Inside the <code>updateArticle</code> view, we will first grab
  the article object from the parsed primary key. We will have two kinds of requests
  here, one will be for fetching the <code>form</code> with the current article data,
  and the next request will be the <code>PUT</code> request for actually saving the
  changes in the article.</p>\n<p>The first request is simple as we need to parse
  the form data with the instance of the article object. We will call the <code>ArticleForm</code>
  with the instance of <code>article</code> this will load the data of the article
  into the form ready to render into the template. So once the <code>GET</code> request
  has been sent, we will render the template with the form pre-filled with the values
  of the article attributes.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># article/views.py</span>\n\n\n<span class=\"k\">def</span> <span class=\"nf\">updateArticle</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"n\">pk</span><span class=\"p\">):</span>\n   <span class=\"n\">article</span>
  <span class=\"o\">=</span> <span class=\"n\">Article</span><span class=\"o\">.</span><span
  class=\"n\">objects</span><span class=\"o\">.</span><span class=\"n\">get</span><span
  class=\"p\">(</span><span class=\"nb\">id</span><span class=\"o\">=</span><span
  class=\"n\">pk</span><span class=\"p\">)</span>\n   <span class=\"n\">form</span>
  <span class=\"o\">=</span> <span class=\"n\">ArticleForm</span><span class=\"p\">(</span><span
  class=\"n\">instance</span><span class=\"o\">=</span><span class=\"n\">article</span><span
  class=\"p\">)</span>\n   <span class=\"n\">context</span> <span class=\"o\">=</span>
  <span class=\"p\">{</span>\n       <span class=\"s1\">&#39;form&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">form</span><span class=\"p\">,</span>\n       <span
  class=\"s1\">&#39;article&#39;</span><span class=\"p\">:</span> <span class=\"n\">article</span><span
  class=\"p\">,</span>\n   <span class=\"p\">}</span>\n   <span class=\"k\">return</span>
  <span class=\"n\">render</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;articles/update.html&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>We
  will create a template in the <code>templates/articles/</code> folder as  <code>update.html</code>
  which will have a simple form for rendering the form fields and a button for sending
  a <code>PUT</code> request. We will render the <code>form</code> and then add a
  button element with the attribute <code>hx-put</code> for sending the <code>PUT</code>
  request to save changes to the article record. We will parse in the <code>article.id</code>
  for the primary key parameter to the view.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"cm\">&lt;!-- templates/articles/update.html --&gt;</span>\n\n<span class=\"p\">&lt;</span><span
  class=\"nt\">div</span> <span class=\"na\">hx-target</span><span class=\"o\">=</span><span
  class=\"s\">&quot;this&quot;</span> <span class=\"na\">hx-swap</span><span class=\"o\">=</span><span
  class=\"s\">&quot;outerHTML&quot;</span><span class=\"p\">&gt;</span>\n  <span class=\"p\">&lt;</span><span
  class=\"nt\">form</span><span class=\"p\">&gt;</span>\n    {% csrf_token %}\n    {{
  form.as_p }}\n    <span class=\"p\">&lt;</span><span class=\"nt\">button</span>
  <span class=\"na\">hx-put</span><span class=\"o\">=</span><span class=\"s\">&quot;{%
  url &#39;article-update&#39; article.id %}&quot;</span>\n      <span class=\"na\">type</span><span
  class=\"o\">=</span><span class=\"s\">&quot;submit&quot;</span><span class=\"p\">&gt;</span>Update<span
  class=\"p\">&lt;/</span><span class=\"nt\">button</span><span class=\"p\">&gt;</span>\n
  \ <span class=\"p\">&lt;/</span><span class=\"nt\">form</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;/</span><span class=\"nt\">div</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>We
  are yet to link the <code>updateArticle</code> into the URLs. We will add the view
  <code>updateArticle</code> into the URLs with the name as <code>article-update</code>
  and <code>update/&lt;int:pk</code> as the slug pattern. This URL pattern will trigger
  the <code>updateArticle</code> when we send an HTTP request to the <code>127.0.0.1:8000/update/4</code>
  for updating the article with id as <code>4</code>.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># article/urls.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.urls</span> <span class=\"kn\">import</span> <span class=\"n\">path</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">.</span> <span class=\"kn\">import</span>
  <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span> <span class=\"o\">=</span>
  <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">listArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-list&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&lt;int:pk&gt;&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">detailArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-detail&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;create/&#39;</span><span class=\"p\">,</span> <span class=\"n\">views</span><span
  class=\"o\">.</span><span class=\"n\">createArticle</span><span class=\"p\">,</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s1\">&#39;article-create&#39;</span><span
  class=\"p\">),</span> \n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;delete/&lt;int:pk&gt;&#39;</span><span class=\"p\">,</span> <span
  class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">deleteArticle</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s1\">&#39;article-delete&#39;</span><span class=\"p\">),</span> \n    <span
  class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;update/&lt;int:pk&gt;&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">updateArticle</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;article-update&#39;</span><span class=\"p\">),</span>
  \n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>This is not done yet,
  we will need to handle the <code>PUT</code> request as well i.e. when the form details
  have been modified and we are about to save changes to the form data. So, we will
  check for the request method's type. If it is a <code>PUT</code> request, we will
  have to process a few things.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># article/views.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django.http</span> <span class=\"kn\">import</span> <span class=\"n\">QueryDict</span>\n\n<span
  class=\"k\">def</span> <span class=\"nf\">updateArticle</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"n\">pk</span><span
  class=\"p\">):</span>\n    <span class=\"n\">article</span> <span class=\"o\">=</span>
  <span class=\"n\">Article</span><span class=\"o\">.</span><span class=\"n\">objects</span><span
  class=\"o\">.</span><span class=\"n\">get</span><span class=\"p\">(</span><span
  class=\"nb\">id</span><span class=\"o\">=</span><span class=\"n\">pk</span><span
  class=\"p\">)</span>\n    <span class=\"k\">if</span> <span class=\"n\">request</span><span
  class=\"o\">.</span><span class=\"n\">method</span> <span class=\"o\">==</span>
  <span class=\"s1\">&#39;PUT&#39;</span><span class=\"p\">:</span>\n        <span
  class=\"n\">qd</span> <span class=\"o\">=</span> <span class=\"n\">QueryDict</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">body</span><span class=\"p\">)</span>\n        <span class=\"n\">form</span>
  <span class=\"o\">=</span> <span class=\"n\">ArticleForm</span><span class=\"p\">(</span><span
  class=\"n\">instance</span><span class=\"o\">=</span><span class=\"n\">article</span><span
  class=\"p\">,</span> <span class=\"n\">data</span><span class=\"o\">=</span><span
  class=\"n\">qd</span><span class=\"p\">)</span>\n        <span class=\"k\">if</span>
  <span class=\"n\">form</span><span class=\"o\">.</span><span class=\"n\">is_valid</span><span
  class=\"p\">():</span>\n            <span class=\"n\">article</span> <span class=\"o\">=</span>
  <span class=\"n\">form</span><span class=\"o\">.</span><span class=\"n\">save</span><span
  class=\"p\">()</span>\n            <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;articles/detail.html&#39;</span><span class=\"p\">,</span> <span
  class=\"p\">{</span><span class=\"s1\">&#39;article&#39;</span><span class=\"p\">:</span>
  <span class=\"n\">article</span><span class=\"p\">})</span>\n    <span class=\"n\">form</span>
  <span class=\"o\">=</span> <span class=\"n\">ArticleForm</span><span class=\"p\">(</span><span
  class=\"n\">instance</span><span class=\"o\">=</span><span class=\"n\">article</span><span
  class=\"p\">)</span>\n    <span class=\"n\">context</span> <span class=\"o\">=</span>
  <span class=\"p\">{</span>\n        <span class=\"s1\">&#39;form&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">form</span><span class=\"p\">,</span>\n        <span
  class=\"s1\">&#39;article&#39;</span><span class=\"p\">:</span> <span class=\"n\">article</span><span
  class=\"p\">,</span>\n    <span class=\"p\">}</span>\n    <span class=\"k\">return</span>
  <span class=\"n\">render</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;articles/update.html&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above <code>updateArticle</code> view, we have to check for a <code>PUT</code>
  request, if we are sending a <code>PUT</code> request, the form instance needs to
  be loaded from the request object. We use the <code>request.body</code> to access
  the data sent in the request. The incoming data received from the <code>request.body</code>
  object is not a valid format to parse it to the form instance, so we will parse
  it using <code>QueryDict</code>. This will allow us to modify the <code>request.body</code>
  object into valid python serializable data.</p>\n<p>So, we import the <code>QueryDict</code>
  from <code>django.http</code> module. We parse the data as the parameter to <code>QueryDict</code>
  and store it in a variable. We then have to get the <code>ArticleForm</code> for
  fetching the data as per the form details, so we parse the instance and also the
  <code>data</code> parameter. The instance is the article object and the data is
  the received form data which we have stored in <code>qd</code> as <code>QueryDict(request.body)</code>.
  This will load the new form data and then we can validate it the form.</p>\n<p>After
  we have verified the form details, we can save the form and this will update the
  article record. Thereby we can render the updated article in the <code>detail</code>
  view with the updated <code>article</code> object as the context.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659252091/blog-media/django-htmx-update-view.png\"
  alt=\"Update View Form Template\" /></p>\n<p>So, this will set up the update view
  as well, we can now create, read, update, and delete an article instance with HTMX
  in templates and Django function-based views without writing any javascript.</p>\n<h2>Summary</h2>\n<p>We
  were able to create a basic CRUD application in Django with HTMX. We used simple
  function-based views to demonstrate the inner details of how we can work with HTMX
  and handle requests from the templates. By creating simple standalone templates,
  we can connect those together to make a fully functional and responsive webpage.
  The UI is not great but the purpose of this tutorial was to make a barebone CRUD
  app to work with the backend using HTMX, so hopefully, you would have got a good
  overview of how HTMX can be integrated into a Django application.</p>\n<p>Overall
  HTMX is a great library that can be used to enhance or even create a new web application
  for making the site responsive and without writing any javascript.</p>\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1659252296/blog-media/django-htmx-demo.gif\"
  alt=\"Django HTMX CRUD Application Demo GIF\" /></p>\n<p>You can check out the source
  code for this project and blog on the <a href=\"https://github.com/Mr-Destructive/htmx-blog-django\">htmx-blog
  GitHub</a> repository.</p>\n<h2>Conclusion</h2>\n<p>From this post, we were able
  to understand the basics of HTMX and how we can integrate it into a Django application.
  Hopefully, you enjoyed the post, if you have any queries or feedback, please let
  me know in the comments or on my social handles. Thank you for reading. Happy Coding
  :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/django-htmx-crud.png
long_description: Gone are the days of writing Ajax requests with javascript, just
  add a few parameters to the HTML content tags and you will be ready for sending
  requests to your backend. So, we are going back in time and correcting the way we
  think of APIs and clien
now: 2025-05-01 18:11:33.311586
path: blog/posts/2022-07-31-Django-HTMX-CRUD-App.md
prevnext: null
series:
- Django-Series
slug: django-htmx-crud
status: published
tags:
- django
- htmx
- python
templateKey: blog-post
title: Django + HTMX CRUD application
today: 2025-05-01
---

## Introduction

Gone are the days of writing Ajax requests with javascript, just add a few parameters to the HTML content tags and you will be ready for sending requests to your backend. So, we are going back in time and correcting the way we think of APIs and client/server-side rendering. We are referring to the [Hypermedia model](https://en.wikipedia.org/wiki/Hypermedia) for levering the server-side processing of data. Let's get our feets wet with this ancient but revolutionary methodology of development with [HTMX](https://htmx.org/).

Yes, HTMX can be used for the API/server-side calls directly in the HTML. We will be exploring the basis of HTMX by creating a basic CRUD application.


## What is HTMX?

The first question that might come up is what and why HTMX? Htmx is a great library, it's a javascript library but wait. It is a javascript library designed to allow us to write less or no javascript at all. It acts as a way to send AJAX requests without you writing any javascript. It uses native browser features directly from HTML.

So, we can use HTMX to create interactive templates in our Django application. We can dynamically call and fetch data from the server by using simple HTML attributes like `hx-get`, `hx-post`, etc. We'll cover those in this article.

You can check the source code used in this article on this [GitHub repository](https://github.com/Mr-Destructive/htmx-blog-django).

## Setup Django Project

We'll be creating a Django project from scratch and designing a basic blog kind of app. We will be creating a quite simple project with a couple of apps like `user` for authentication and `article` for the CRUD part of our blog application.

To set up a django project, we can run the following commands to quickly get up and running with a base django project.

```
mkdir htmx_blog
python3 -m venv .venv
source .venv/bin/activate
pip install django
django-admin startproject htmx_blog .
```

I have a base user model that I use for a simple authentication system in some basic django projects, you can define your own user app or get the app from [here](https://github.com/Mr-Destructive/django-todo/tree/master/user).

So, that being said, we will be using the user model for the article model which we will be defined next. By creating a basic signup functionality, you are good to go!

### Create the Article app

We will need at least an app to work with htmx as we will define models, views, and URLs later as we configure the htmx.

```
django-admin startapp article
```

After the app has been created, you can add those app labels into the `INSTALLED_APPS` config in the `settings.py` file. The `user` app and the `article` app need to be added to the installed apps for the django to pick those up for various contexts related to the project.

```
# htmx_blog/settings.py

INSTALLED_APPS = [
    ...
    ...
    ...

    'article',  
    'user',
]
```

We are sone with the base setup, we also would require a few more configs for the proper working of the project.

### Setup Templates and Static files

Templates will play an important role in the htmx part, so it is equally important to configure them properly before dabbling into the htmx and client-side rendering of data.

I like to keep all the templates in a single folder in the `BASE_DIR` with separate sub-folders for specific apps. Also a single `static` folder with `css`, `js`, and `images` as the sub-folfers for a larger project.

```
mkdir templates static
```

Further, configure the created static and templates in the settings.

```python

TEMPLATES = [
    {
        'BACKEND': 'django.template.backends.django.DjangoTemplates',
        'DIRS': [os.path.join(BASE_DIR, "templates")],
        'APP_DIRS': True,
        'OPTIONS': {
            'context_processors': [
                'django.template.context_processors.debug',
                'django.template.context_processors.request',
                'django.contrib.auth.context_processors.auth',
                'django.contrib.messages.context_processors.messages',
            ],
        },
    },
]

STATIC_URL = 'static/'
STATICFILES_DIRS = [str(BASE_DIR/ "static")]
STATIC_ROOT = BASE_DIR / "staticfiles"
```

### Initial migration

Run migration command for the user model and default model in the django project.

```
python manage.py makemigrations
python manage.py migrate
```

So, this project will also include authentication simple registration, and login/logout routes. We will be using the default Django User model by creating an abstract user just in case we require any additional attributes. 

## Setup HTMX

We don't have to configure much for using HTMX as it is a javascript library, we can call it via a CDN or manually install it and link up the static javascript files. Either way, both are equally good, you may like the one I might like the other. 

If you already have a base template, you can simply put the below script inside the head tag of the template. This will make us the htmx attributes available.

```html
<script src="https://unpkg.com/htmx.org@1.8.0"></script>
```

If you don't have a base template, you can create one by creating an HTML file inside the `templates` directory. The name can be anything but be careful for following up as it might be different for me. I will choose `base.html` as the template for this project. It will look something like as follows:

```html
<!-- tempaltes/base.html -->

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>HTMX Blog</title>
    {% load static %}
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
    <script src="https://unpkg.com/htmx.org@1.8.0"></script>
</head>
<body >
        <nav>
        <h2>HTMX Blog</h2>
        <div class="navbar">
          {% if user.is_authenticated %}
            <a class="nav-item nav-link" href="{% url 'logout' %}"><button class="btn btn-link">Logout</button></a>
          {% else %}
            <a class="nav-item nav-link" href="{% url 'login' %}"><button class="btn btn-link">Login</button></a>
            <a class="nav-item nav-link" href="{% url 'register' %}"><button class="btn btn-link">Register</button></a>
          {% endif %}
        </div>
        </nav>

    {% block body %}
    {% endblock %}
</body>
</html>
```

I have a nav bar with my user authentication views, simply a login or signup button if the user is not logged in and a log-out button if the user is authenticated. We have added the htmx script file from the CDN just before the end of the head tag. We also have included the bootstrap CSS file for a decent UI which we will be creating in this post.

That is one of the ways, htmx can be injected into an HTML template, you can even download the javascript file from the [htmx cdn](https://unpkg.com/browse/htmx.org/dist/). Further, this can be downloaded or pasted into your local folder and served as a static file or embedded directly into an HTML template.

## Defining Models

We will start the tutorial by defining the model of the application we are creating. Here, we will create a simple Article model with a few parameters like `title`, `content`, `author`, etc. 

```python
from django.db import models
from user.models import Profile

class Article(models.Model):
    Article_Status = (
        ("DRAFT", "Draft"),
        ("PUBLISHED", "Published"),
    )
    title = models.CharField(max_length=128, unique=True)
    content = models.TextField()
    author = models.ForeignKey(Profile, on_delete=models.CASCADE)
    status = models.CharField(
        max_length=16,
        choices=Article_Status,
        default=Article_Status[0],
    )

    def __str__(self):
        return self.title
```

In the above model `Article`, we have a few fields like `title` simple Character Field, `content` as a text field as it will be a large text as the post body, `author` which is a ForeignKey to the User Model. We also have the status, which is defined as a character field but with a few choices like `draft` or `published`, we can further modify this status as public or private. But just keeping it simple and easy to understand.

The object reference name for this model is the title as we have defined in the dunder string method. So, that is a simple model created, we can now migrate the changes into the database for adding the tables and attributes.

```
python manage.py makemigrations
python manage.py migrate
```

This will make migrations to the database i.e. convert the python model class into database tables and attributes. So, once the migration process is completed successfully, we can move into the crust of this article which is to actually design the views. In the next section, we will be utilizing the models in our views for representing the data on the templates.


## Creating Article Form

Before diving into the views section, we need a few things like the Article Form, which will be a Django Model-based form. It will help us a lot in creating or updating the fields for the article model. We can define a form in a python file called `forms.py`, it's not necessary to keep your forms in the `forms.py` but if you have a lot of forms and models, it becomes a good practice to organize the components of our app. So, I'll be creating a new file inside of the `article` app called `forms.py` and defining the `ArticleForm`.

```python
# article/forms.py

from django import forms
from .models import Article


class ArticleForm(forms.ModelForm):
    class Meta:
        model = Article
        exclude = (
            "created",
            "updated",
            "author",
        )
        widgets = {
            "title": forms.TextInput(
                attrs={
                    "class": "form-control",
                    "style": "max-width: 450px; align: center;",
                    "placeholder": "Title",
                }
            ),
            "content": forms.Textarea(
                attrs={
                    "class": "form-control",
                    "style": "max-width: 900px;",
                    "placeholder": "Content",
                }
            ),
        }
```

So, the forms are inherited from the [ModelForm] which allows us to create forms based on our model. So, we specify the model name which in our case is `Article` and further we can have `exclude` or `fields` tuples. To exclude certain fields in the actual form, just parse the tuple of those attributes and if you want to only select a few attributes, you can specify the `fields` tuple and mention the required fields for the form.

So, if we have a lot of things to be included in the form, we can specify only the attributes to be excluded with the `exclude` tuple. And if we have a lot of fields to be excluded, we can use the `fields` tuple to specify which attributes to use in the form.

Let's take an example: For the above ArticleForm, if we wanted to specify the required fields to be included in the form, then we might use the `fields` tuple like below the rest will be not rendered in the form fields.

```
class ArticleForm(forms.ModelForm):
    class Meta:
        model = Article
        fields = (
            "title",
            "content",
            "status",
        )
```

Both of them can be used, it just depends on how many fields you have to exclude or include in the rendered form.

We have also specified the `widgets` attribute which gives a bit more control on how we need to display the form in a template. So I have specified the type of input it needs to render like a simple text input for the title, text area for content, etc. The cool thing about this is it can automatically set these by knowing the type of field in the model, but sometimes it can be a bit undesired mostly with complex relationships and attributes.

## Creating Views

Let's start creating views for creating, reading, updating, and deleting articles from the database. I will be using function-based views just because we are understanding the flow of how HTMX and Django can be integrated so we need to dive in deeper and understand the actual flow of the process.

### Create View

So, creating articles seems like a good way to start off. We can create a simple function-based view which will initially load in an empty `ArticleForm` and if the request is `GET` we will render the form in the `create.html` template. If the request is `POST` which will be after we submit the form, we will validate the form and attach the current user as the author of the article and save the for instance which will create an article record and this object will be rendered to the detail template.

```python
from django.shortcuts import render
from .models import Article
from .forms import ArticleForm

def createArticle(request):
    form = ArticleForm()
    context = {
        'form': form,
    }
    return render(request, 'articles/create.html', context)
```

#### Rendering the Form

We are creating an empty instance of `ArticleForm` and rendering it in the template. So, this will render the empty form in the `create.html` template.

```html
<!-- templates/articles/create.html -->

{% extends 'base.html' %}

{% block body %}
<div hx-target="this" hx-swap="outerHTML">
  <form>
    {% csrf_token %}
    {{ form.as_p }}
    <button hx-post="." class="btn btn-success"
      type="submit">Save</button>
  </form>
</div>
{% endblock %}
```

Now, here we are inheriting from the base template and creating a form tag in HTML with the `{{ form}}` for rendering the form fields, we finally have the `button` element for submitting the form. We have used the `hx-post` attribute. More on this in just a minute. So, this is we create a template for rendering the article form.

We have used the `hx-post` attribute here, which will send a `POST` request to the current `URL` represented by `hx-post="."`. You might have noticed the `div` attributes, the `hx-target` and `hx-swap`, so these are some of the many attributes provided by the htmx library for controlling the reactivity of the requests made. The `hx-target` allow us to specify the element or tag to which the data will be rendered. The `hx-swap` goes hand-in-hand for specifying the target DOM like `innerHTML`, `outerHTML`, etc. You can see the various options on the [htmx docs](https://htmx.org/docs/#swapping). By specifying the `hx-swap` as  `outerHTML`, we are saying to replace the entire element with the incoming content from the request which we will send with nearby request triggers.

We need to map the view to a URL in order to get a good idea about the request and parsed content.

We'll create a `create/` route and bind it to the `createArticle` view with the name `article-create`.

```python
# article/urls.py

from django.urls import path
from . import views

urlpatterns = [
    path('create/', views.createArticle, name='article-create'), 
]
```

This URL will be mapped to the global URL in the project, here we can simply specify the prefix for the URLs in the `article` app and include those URLs.

```python
# htmx_blog/urls.py

from django.contrib import admin
from django.urls import path, include

urlpatterns = [
    path('admin/', admin.site.urls),
    path('user/', include('user.urls'), name='auth'),
    path('', include('article.urls'), name='home'),
]
```
Feel, free to add any other URL pattern like for instance, the article app is at `/` i.e. `127.0.01.:8000/`, you can add any other name like `127.0.0.1:8000/article/` by adding `path('article/', include('article.urls'))`.

![Django HTMX Create view Form Template](https://res.cloudinary.com/techstructive-blog/image/upload/v1659252089/blog-media/django-htmx-create-view.png)

So, finally, we are sending a `GET` request to the `127.0.0.1:8000/create/` and this will output the form. As we have a `POST` request embedded in the button inside the form, we will send the `POST` request to the same URL -> `127.0.0.1:8000/create/`.

#### Submitting the Form

Let's handle the `POST` request in the create view.

```python
from django.shortcuts import render
from .models import Article
from .forms import ArticleForm

def createArticle(request):
    form = ArticleForm(request.POST or None)
    if request.method == 'POST':
        if form.is_valid():
            form.instance.author = request.user
            article = form.save()
            return render(request, 'articles/detail.html', {'article': article})
    context = {
        'form': form,
    }
    return render(request, 'articles/create.html', context)
```

**Simple explanation**

- Create a form instance of ArticleForm with the request data or empty -> `ArticleForm(request.POST or None)`
- If it's a POST request, validate and create the article, render the article object in `detail.html` template.
- If it's a GET request, render the empty form in `create.html`


There are a few changes in the view, instead of initializing the form to empty i.e. `ArticleForm()`, we are initializing with `ArticleForm(request.POST or None)`. This basically means that if we are having something in the `request.POST` dict, we will initialize the Form with that data or else an empty form instance.

Next, we check if the request if `POST`, if it is then we check if the form is valid i.e. the form fields are not empty or if any other constraint on the model attributes is satisfied or not. If the form data is valid, we attach the author as the currently logged-in User/user who sent the request. Finally, we save the form which in turn creates the article record in the database. We then render the created article in the `detail.html` template which is not yet created.

So, the `htmx-post` attribute has worked and it will send a post request to the same URL i.e. `127.0.0.1:8000/create` and this will again trigger the view `createArticle` this time we will have `request.POST` data. So, we will validate and save the form.

### Detail View

The detail view is used for viewing the details of an article. This will be rendered after the article has been created or updated. This is quite simple, we need an `id` or `primary key(pk)` of an article and render the `title` and `content` of the article in the template.

We pass in a primary key along with the request as a parameter to the view, the `pk` will be passed via the URL. We fetch the Article object with the id as the parsed `pk` and finally render the `detail.html` template with the article object. The `context['article']` can be accessed from the template to render the specific attributes like `title`, `content`, etc.

```python
# article/views.py

def detailArticle(request, pk):
    article = Article.objects.get(id=pk)
    context = {'article': article}
    return render(request, 'articles/detail.html', context)

```

We can now bind the view to a URL and parse the required parameter `pk` to the view. 

```python
from django.urls import path
from . import views

urlpatterns = [
    path('create/', views.createArticle, name='article-create'), 
    path('<int:pk>', views.detailArticle, name='article-detail'), 
]
```

We have parsed the `pk` as `int` to the URL parameter, so for an article with id=4, the URL will be, `127.0.0.1:8000/4/`.

We need to create the template for rendering the context from the `detailArticle` view. So, we create the `detail.html` in the `templates/articles` folder. We inherit the base template and render the `article.title` and the `article.content` with a linebreaks template filter so as to display the content properly.

```html
<!-- templates/articles/detail.html -->


{% extends 'base.html' %}
{% block body %}
<div id="article-card">
  <h2>{{ article.title }}
  <p>{{ article.content|linebreaks|safe }}</p>
<div>
{% endblock %}

```

![Detail View Template](https://res.cloudinary.com/techstructive-blog/image/upload/v1659252227/blog-media/django-htmx-detail-view.png)

So, we can now use `createArticle` view as well as `detailArticle` view, this both are configured properly, so (CR) or CRUD is completed. We can add `listArticle` for listing out all the author's(logged-in user) articles.

### List View

Listview of the articles is much similar to the detail view as it will return a list of articles rather than a single article.

So in the `listArticle` view, we will return all the articles with the author as the user who sent the request/logged-in user. We will parse this object list into the template as `base.html` or `articles/list.html`.

```python
# article/views.py


def listArticle(request):
    articles = Article.objects.filter(author=request.user.id)
    context = {
        'articles': articles,
    }
    return render(request, 'base.html', context)
```

We will add the URL route for this as the `/` route that is on `127.0.0.1:8000/` this is the base URL for the article app and is the route for the `listArticle` view. So, we will display the list of articles on the homepage.

```python
# article/urls.py


from django.urls import path
from . import views

urlpatterns = [
    path('<int:pk>', views.detailArticle, name='article-detail'), 
    path('create/', views.createArticle, name='article-create'), 
    path('', views.listArticle, name='article-list'), 
]
```

Let's create the template for the list view which will iterate over the articles and display the relevant data like the title and link to the article.

```html
<!-- templates/articles/list.html -->

<ul id="article-list">
  {% for article in articles %}
  <li>
    <div class="card" style="width: 18rem;">
      <div class="card-body">
        <h5 class="card-title">{{ article.title }}</h5>
        <p class="card-text">{{ article.content|truncatewords:5  }}</p>
        <a href="{% url 'article-detail' article.id %}" class="card-link">Read more</a>
      </div>
    </div>
  </li>
  {% endfor %}
</ul>
```

We have used the `truncatewords:5` template filter for only displaying the content of the articles till the first 5 words as it is just a list view, we don't want to display every detail of the article here.

![List view Template](https://res.cloudinary.com/techstructive-blog/image/upload/v1659252293/blog-media/django-htmx-list-view.png)

We can use this template to render in the `base.html` file.

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>HTMX Blog</title>
    {% load static %}
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
    <script src="https://unpkg.com/htmx.org@1.8.0"></script>
</head>
<body hx-target="this" hx-swap="outerHTML" hx-headers='{"X-CSRFToken": "{{ csrf_token }}"}'>
        <nav>
        <h2>HTMX Blog</h2>
        <div class="navbar">
          {% if user.is_authenticated %}
            <a class="nav-item nav-link" href="{% url 'article-list' %}"><button class="btn btn-link">Home</button></a>
            <a class="nav-item nav-link" href="{% url 'logout' %}"><button class="btn btn-link">Logout</button></a>
          {% else %}
            <a class="nav-item nav-link" href="{% url 'login' %}"><button class="btn btn-link">Login</button></a>
            <a class="nav-item nav-link" href="{% url 'register' %}"><button class="btn btn-link">Register</button></a>
          {% endif %}
        </div>
        </nav>

    {% block body %}
    <a href="{% url 'article-create' %}"><button class="btn btn-success" >Create</button></a>
    {% include 'articles/list.html' %}
    {% endblock %}
</body>
</html>
```

We have now included the `list.html` template on the homepage and also added the `create` button as the link to the `article-create` URL.

### Delete View

For deleting an article, we will simply rely on htmx for sending the request and on that request, we will delete the current article and render the updated list of articles.

With the `deleteArticle` view, we will take in two parameters the request which is by default for a Django function-based view, and the primary key as `pk`. Again we will parse the `pk` from the URL. We will delete the article object and get the latest list of articles. Finally, render the updated list of articles in the base template which is our list view.

```python
# article/views.py


def deleteArticle(request, pk):
    Article.objects.get(id=pk).delete()
    articles = Article.objects.filter(author=request.user)
    context = {'article': articles}
    return render(request, "base.html", context)

```

We will add the `deleteArticle` into the URL patterns and call it `article-delete` with the URL of `delete/<int:pk>`. This will allow us to send a request to the URL `127.0.0.1:8000/delete/4` for deleting the article with id `4`.

```python
# article/urls.py


from django.urls import path
from . import views

urlpatterns = [
    path('', views.listArticle, name='article-list'), 
    path('<int:pk>', views.detailArticle, name='article-detail'), 
    path('create/', views.createArticle, name='article-create'), 
    path('delete/<int:pk>', views.deleteArticle, name='article-delete'), 
]
```

In the delete view, the template is important as we want to send a request appropriately to the defined URL. To do that, we will have a form but it won't have any inputs as such just a button that indicates to delete the current article. We will add the `hx-delete` attribute as the URL to the `deleteArticle` view. with the id of the article. This will send a request to the `article-delete` URL which will, in turn, trigger the view with the given id and delete the article.

We have added the `hx-confirm` attribute for showing a pop-up of confirmation of deleting the article. As you can see we have added a little script for adding `csrf_token` into the HTML, this is important in order to submit a form with a valid `CSRFToken`.

```html
<!-- templates/article/delete.html -->

<script>
  document.body.addEventListener('htmx:configRequest', (event) => {
    event.detail.headers['X-CSRFToken'] = '{{ csrf_token }}';
  })
</script>
<div >
  <form method="post" >
  {% csrf_token %}
    <button class="btn btn-danger"
      hx-delete="{% url 'article-delete' article.id %}"
      hx-confirm="Are you sure, You want to delete this article?"
      type="submit">
      Delete
    </button>
  </form>
</div>
```

Do you have a question like how do we access the `article.id`? we are not rendering the `delete.html` template from the view, so there is no context to pass. We will include this snippet into the detail view template, so as to have the option of deleting the current article.

We will modify the `articles/detail.html` template and include the `delete.html` template. This includes simply adding an HTML template in the specified location. So, we will basically inject the delete form into the detail template.

```html
{% extends 'base.html' %}
{% block body %}
<div hx-target="this" hx-swap="outerHTML">
  <h2>{{ article.title }}
  {% include 'articles/delete.html' %}
  <p>{{ article.content|linebreaks|safe }}</p>
<div>
{% endblock %}
```

Hence, we will have a nice option to delete the article in the detail section, this can be placed anywhere but remember, we need to add the `hx-target="this"` and `hx-swap="outerHTML"` in the div so as to correctly swap the HTML content after the request has been made.

### Update View

We can now move into the final piece of the CRUD i.e. `Update`. This will be similar to the `createArticle` with a couple of changes. We will parse parameters like `pk` to this view as well because we want to update a specific article. So, we will have to get the primary key of the article from the URL slug.

Inside the `updateArticle` view, we will first grab the article object from the parsed primary key. We will have two kinds of requests here, one will be for fetching the `form` with the current article data, and the next request will be the `PUT` request for actually saving the changes in the article.

The first request is simple as we need to parse the form data with the instance of the article object. We will call the `ArticleForm` with the instance of `article` this will load the data of the article into the form ready to render into the template. So once the `GET` request has been sent, we will render the template with the form pre-filled with the values of the article attributes.
 
 ```python
# article/views.py


def updateArticle(request, pk):
    article = Article.objects.get(id=pk)
    form = ArticleForm(instance=article)
    context = {
        'form': form,
        'article': article,
    }
    return render(request, 'articles/update.html', context)
```

We will create a template in the `templates/articles/` folder as  `update.html` which will have a simple form for rendering the form fields and a button for sending a `PUT` request. We will render the `form` and then add a button element with the attribute `hx-put` for sending the `PUT` request to save changes to the article record. We will parse in the `article.id` for the primary key parameter to the view.

```html
<!-- templates/articles/update.html -->

<div hx-target="this" hx-swap="outerHTML">
  <form>
    {% csrf_token %}
    {{ form.as_p }}
    <button hx-put="{% url 'article-update' article.id %}"
      type="submit">Update</button>
  </form>
</div>
```

We are yet to link the `updateArticle` into the URLs. We will add the view `updateArticle` into the URLs with the name as `article-update` and `update/<int:pk` as the slug pattern. This URL pattern will trigger the `updateArticle` when we send an HTTP request to the `127.0.0.1:8000/update/4` for updating the article with id as `4`.

```python
# article/urls.py


from django.urls import path
from . import views

urlpatterns = [
    path('', views.listArticle, name='article-list'), 
    path('<int:pk>', views.detailArticle, name='article-detail'), 
    path('create/', views.createArticle, name='article-create'), 
    path('delete/<int:pk>', views.deleteArticle, name='article-delete'), 
    path('update/<int:pk>', views.updateArticle, name='article-update'), 
]
```

This is not done yet, we will need to handle the `PUT` request as well i.e. when the form details have been modified and we are about to save changes to the form data. So, we will check for the request method's type. If it is a `PUT` request, we will have to process a few things.

```python
# article/views.py


from django.http import QueryDict

def updateArticle(request, pk):
    article = Article.objects.get(id=pk)
    if request.method == 'PUT':
        qd = QueryDict(request.body)
        form = ArticleForm(instance=article, data=qd)
        if form.is_valid():
            article = form.save()
            return render(request, 'articles/detail.html', {'article': article})
    form = ArticleForm(instance=article)
    context = {
        'form': form,
        'article': article,
    }
    return render(request, 'articles/update.html', context)
```

In the above `updateArticle` view, we have to check for a `PUT` request, if we are sending a `PUT` request, the form instance needs to be loaded from the request object. We use the `request.body` to access the data sent in the request. The incoming data received from the `request.body` object is not a valid format to parse it to the form instance, so we will parse it using `QueryDict`. This will allow us to modify the `request.body` object into valid python serializable data.

So, we import the `QueryDict` from `django.http` module. We parse the data as the parameter to `QueryDict` and store it in a variable. We then have to get the `ArticleForm` for fetching the data as per the form details, so we parse the instance and also the `data` parameter. The instance is the article object and the data is the received form data which we have stored in `qd` as `QueryDict(request.body)`. This will load the new form data and then we can validate it the form.

After we have verified the form details, we can save the form and this will update the article record. Thereby we can render the updated article in the `detail` view with the updated `article` object as the context.

![Update View Form Template](https://res.cloudinary.com/techstructive-blog/image/upload/v1659252091/blog-media/django-htmx-update-view.png)

So, this will set up the update view as well, we can now create, read, update, and delete an article instance with HTMX in templates and Django function-based views without writing any javascript.

## Summary

We were able to create a basic CRUD application in Django with HTMX. We used simple function-based views to demonstrate the inner details of how we can work with HTMX and handle requests from the templates. By creating simple standalone templates, we can connect those together to make a fully functional and responsive webpage. The UI is not great but the purpose of this tutorial was to make a barebone CRUD app to work with the backend using HTMX, so hopefully, you would have got a good overview of how HTMX can be integrated into a Django application.

Overall HTMX is a great library that can be used to enhance or even create a new web application for making the site responsive and without writing any javascript.

![Django HTMX CRUD Application Demo GIF](https://res.cloudinary.com/techstructive-blog/image/upload/v1659252296/blog-media/django-htmx-demo.gif)

You can check out the source code for this project and blog on the [htmx-blog GitHub](https://github.com/Mr-Destructive/htmx-blog-django) repository.

## Conclusion

From this post, we were able to understand the basics of HTMX and how we can integrate it into a Django application. Hopefully, you enjoyed the post, if you have any queries or feedback, please let me know in the comments or on my social handles. Thank you for reading. Happy Coding :)