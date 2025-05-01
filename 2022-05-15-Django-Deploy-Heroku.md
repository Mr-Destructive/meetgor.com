---
article_html: "<h2>Introduction</h2>\n<p>Django projects are quite easy to build and
  simple to understand, you might have created a Django application and wanted to
  show it to the world? You can deploy a basic Django application with a database(PostgreSQL)
  with Heroku. It provides a decent free tier with some great features and add-ons.
  A free tier Heroku account has a limitation of 5 apps, limited data in the database,
  limited connections to the server per month, and so on.</p>\n<p>Though the free
  tier is not a great option for bigger applications, it suits really well for smaller
  scale and ide projects, so we will be looking into the details of how to deploy
  a Django application to <a href=\"https://heroku.com/\">Heroku</a> which is a Platform
  as Service (PaS).</p>\n<p>This series will be an extension of the series <a href=\"https://techstructiveblog.hashnode.dev/series/django-basics\">Django
  basics</a> which covered the basics of the Django framework, we covered from basic
  Django fundamentals to building a CRUD API. In this series, we will be exploring
  some platforms for giving a better understanding of how things work and understanding
  a few components that were left as default while understanding the basics of Django.
  Let's get started with <a href=\"https://techstructiveblog.hashnode.dev/series/django-deployment\">Django
  Deployment</a>!</p>\n<h2>Creating a Django Application</h2>\n<p>For deploying an
  app, we definitely need an app, we need to create a basic Django application to
  deploy on the web. We'll be creating a simple blog application with a couple of
  views and a simple model structure. As for the database, we'll be using Postgres
  as Heroku has an add-on for it and it is pretty easy to configure.</p>\n<h3>Set
  up a virtual environment</h3>\n<p>We need to set up a virtual environment in order
  to keep the Django project neat and tidy by managing the project-specific dependencies
  and packages. We can use the <a href=\"https://pypi.org/project/virtualenv/\">virtualenv</a>
  package to isolate a python project from the rest of the system.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span># install the virtualenv
  package\npip install virtualenv\n\n# create a virtual env for the project\nvirtualenv
  .venv\n\n# activate the virtualenv\nWindows:\n.venv\\Scripts\\activate\n\nLinux/macOS:\nsource
  .venv/bin/activate\n</pre></div>\n\n</pre>\n\n<p>This will set up the project nicely
  for a Django project, you now install the core Django package and get started with
  creating a Django application.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># install django</span>\npip install django\n\n<span class=\"c1\">#
  start a django project</span>\ndjango-admin startproject blog .\n\n<span class=\"nb\">cd</span>
  blog\n\n<span class=\"c1\"># create a application in django project</span>\npython
  manage.py createapp api\n\n<span class=\"c1\"># Create some models, views, URLs,
  templates</span>\n\n<span class=\"c1\"># run the server</span>\npython manag.py
  runserver\n</pre></div>\n\n</pre>\n\n<p>We assume you already have a Django project
  configured with some basic URLs, views and templates or static files as per your
  project and requirements, for this tutorial I will be using the simple blog application
  from my previous Django tutorials as a reference. You can follow along with my <a
  href=\"https://techstructiveblog.hashnode.dev/series/django-basics\">Django Basics</a>
  series and refer to the Blog Application project on <a href=\"https://github.com/Mr-Destructive/django-blog\">GitHub</a>.</p>\n<h2>Configuring
  the Django Application</h2>\n<p>Make sure to create and activate the virtual environment
  for this django project. This should be done to manage the dependencies and packages
  used in the project. If you are not aware of the virtual environment and django
  setup, you can follow up with this <a href=\"https://mr-destructive.github.io/techstructive-blog/django-setup-script/\">post</a>.</p>\n<h3>Creating
  a runtime.txt file</h3>\n<p>Now, Firstly we need to specify which type and version
  of language we are using. Since Django is a Python-based web framework, we need
  to select the python version in a text file.</p>\n<p><strong>runtime.txt</strong></p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>python-3.9.5\n</pre></div>\n\n</pre>\n\n<p>Here,
  the version can be anything as per your project and the packages installed.</p>\n<h3>Creating
  requirements.txt file</h3>\n<p>We'll first create a <code>requirements.txt</code>
  file for storing all the dependencies and packages installed in the application.
  This will help in installing dependencies while deploying the application. We can
  either use a <code>requirements.txt</code> file using <code>virtualenv</code> or
  a <code>Pipfile</code> using Pipenv. Both serve the same purpose but a bit differently.</p>\n<p>Assuming
  you are in an isolated virtual environment for this Django project, you can create
  a requirements.txt file using the below command:</p>\n<p>Make sure the virtualenv
  is activated before running the command:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>pip freeze &gt; requirements.txt\n</pre></div>\n\n</pre>\n\n<p>This
  will create a simple text file that contains the package names along with the versions
  used in the current virtual environment. A simple Django requirements file would
  look something like this:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>asgiref==3.4.1\nDjango==3.2.11\npytz==2021.3\nsqlparse==0.4.2\ntyping_extensions==4.0.1\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652614060461/kPTZ9R8Xp.png\"
  alt=\"image.png\" /></p>\n<p>This is vanilla Django without any additional dependencies
  but if you have integrated other functionalities like Django Rest Framework, PostgreSQL,
  Crispy Forms, Schedulers, etc. there might be additional dependencies that become
  quite crucial for the smooth working of the project.</p>\n<p>If you are using pipenv,
  you don't need to make any efforts to manually activate and manage virtual environment,
  the dependencies are installed and taken care of by the pipenv installer. You just
  need to make sure to install any package with <code>pipenv install</code> and not
  <code>pip install</code> for better and improved package tracking.</p>\n<p>So, that's
  all we need to take care of packages and keep a list of these integrated packages
  for the project. You need to update the requirements.txt file every time you install
  any new package or modify the dependencies for a project. Simply use the command
  <code>pip freeze &gt; requirements.txt</code> in the activated virtual environment.</p>\n<h3>Creating
  a Procfile</h3>\n<p>Next up, we have the <code>Procfile</code>, a procfile is a
  special file that holds information about the processes to be run to start or activate
  the project. In our case, for django we need a web process that can manage the server.</p>\n<p>A
  Procfile is a simple file without any extension, make sure to write <code>Procfile</code>
  as it is as the name of the file in the root folder of the project. Inside the file
  add the following contents:</p>\n<p><strong>Procfile</strong></p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>Procfile</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nl\">web</span><span class=\"p\">:</span><span class=\"w\"> </span>gunicorn<span
  class=\"w\"> </span><span class=\"err\">&lt;</span>project_name<span class=\"err\">&gt;</span>.wsgi<span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>As we can see we have defined
  the <code>web</code> process using <code>gunicorn</code>, <a href=\"https://pypi.org/project/gunicorn/\">Gunicorn</a>
  is a python package that helps in creating WSGI HTTP Server for the UNIX operating
  systems. So, we need to install the package and update the package dependency list.</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>pip install gunicorn\n\npip
  freeze &gt; requirements.txt\n</pre></div>\n\n</pre>\n\n<p>That would be good to
  go for creating and serving up the project while deploying the project on Heroku.</p>\n<h2>Creating
  a Heroku App</h2>\n<p>A Heroku App is basically like your Django Project, you can
  create apps for deploying your django projects on Heroku. You are limited to 5 apps
  on the Free tier but that can be expanded on the paid plans.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652456732519/cyOQZ3UZK.png\"
  alt=\"image.png\" /></p>\n<p>The name of your heroku app should be unique globally,
  you need to try a few combinations before a good one fits your project. This name
  has no significance on your django project code, though it would be the name from
  which you would access the web application as a name <code>&lt;app-name&gt;.herokuapp.com</code>.</p>\n<p>So,
  choose it wisely if you are not attaching a custom domain. You can attach a custom
  domain, you can navigate to the <code>domain</code> section in the settings tab.</p>\n<h2>Setting
  up the database</h2>\n<p>To set up and configure a database in django on Heroku,
  we need to manually acquire and attach a PostgreSQL add-on to the heroku app.</p>\n<ul>\n<li>Firstly
  locate to the Resources Tab in your Heroku app.</li>\n<li>Search <code>postgres</code>
  in the Add-ons Search bar</li>\n<li>Click on the <code>Heroku Postgres</code> Add-on</li>\n<li>Submit
  the Order Form and you have the add-on enabled in the app.</li>\n</ul>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652456842273/ijeWsVdOf.png\"
  alt=\"image.png\" /></p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652456877447/dLG30ac_m.png\"
  alt=\"image.png\" /></p>\n<p>Alternately, you can use the Heroku CLI,</p>\n<h3>Heroku
  CLI Setup</h3>\n<p>You can use the Heroku CLI which is a command-line interface
  for managing and creating Heroku applications.  You need to first install the CLI
  by heading towards the <a href=\"https://devcenter.heroku.com/articles/heroku-command-line\">heroku
  documentation</a>. Once the CLI is installed, you need to login into your Heroku
  account by entering the following command:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>heroku login\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652605604920/HnTr2KbTi.png\"
  alt=\"image.png\" /></p>\n<p>This will allow us to work with heroku commands and
  manage our heroku application from the command line itself. The below command will
  create a add-on for <code>heroku-postgres</code> for the application provided as
  the parameter options</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>heroku addons:create heroku-postgresql:hobby-dev
  --app &lt;app_name&gt;\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652507166978/i1IJ5EGjJ.png\"
  alt=\"image.png\" /></p>\n<p>This should hopefully add a fresh instance of a postgres
  database for your heroku app. You can now configure the database for your project,
  we simply need the Database URL from the heroku app dashboard. We'll see how to
  configure the environment variables in Django for Heroku to keep your secrets like
  the <code>SECRET_KEY</code>, <code>DATABSE_URL</code>, etc.</p>\n<p>If you want
  MySQL as a database, you can check out the <a href=\"https://devcenter.heroku.com/articles/cleardb\">ClearDB</a>
  Add-On for Heroku to simply attach it like Postgres on your Heroku application.
  Also, if you wish to add <a href=\"https://www.mongodb.com/compatibility/mongodb-and-django\">MongoDB</a>
  into your Django application on Heroku, you can use <a href=\"https://devcenter.heroku.com/articles/ormongo\">Object
  Rocket</a>(OR Mongo). It is not available in the free tier though, unlike PostgreSQL
  and MySQL.</p>\n<h2>Configuring Environment Variables</h2>\n<p>We need to keep our
  secrets for the django project out of the deployed code and in a safe place, we
  can create environment variables and keep them in a <code>.env</code> file which
  will be git-ignored. We do not want this <code>.env</code> file to be open source
  and thus should not be committed.</p>\n<p>We first need to create a new secret key
  if you already have a GitHub repository, chances are you would have committed the
  default secret key open for the world to see, it is an insecure way of deploying
  django apps in production.</p>\n<p>So, open up a terminal and launch a python REPL:</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>python\n\nimport secrets\nsecrets.token_hex(24)\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652512239319/5AACaTGOD.png\"
  alt=\"image.png\" /></p>\n<p>This should generate a secret key that only you know
  now. So, just copy the key without the quotes and create a file <code>.env</code>
  in the root project folder.</p>\n<p><strong>.env</strong></p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>SECRET_KEY=76419fd6885a677f802fd1d2b5acd0188e23e001042b05a8\n</pre></div>\n\n</pre>\n\n<p>The
  <code>.env</code> file should also be added to the <code>.gitignore</code> file,
  so simply append the following in the <code>.gitignore</code> file</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>.env\n</pre></div>\n\n</pre>\n\n<p>This
  file is only created to test the project locally, so we need to also make this key
  available on heroku. For doing that we need to add Config Variables to the heroku
  app.</p>\n<ul>\n<li>Locate to the Settings Tab in your Heroku Application Dashboard</li>\n<li>We
  have the <code>Config Vars</code> section in the located tab\n= We need to reveal
  those variables and we will be able to see all the variables.</li>\n</ul>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652456988713/5VM6E29_o.png\"
  alt=\"image.png\" /></p>\n<p>We already have a <code>DATABASE_URL</code> variable
  declared when we attached the <code>django-postgres</code> database to our application.
  We will now add one more variable to the Application, i.e. the <code>SECRET_KEY</code>.
  Simply, enter the name of the key and also enter the value of the key, so basically
  a key-value pair. Simply click on the <code>Add</code> button and this will add
  the variable to your application.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652515244870/LRyPzJr41.png\"
  alt=\"image.png\" /></p>\n<p>You also need to copy the <code>DATABSE_URL</code>
  into our local setup file(<code>.env</code> file). Copy the Database URL and paste
  it into the <code>.env</code> file as follows:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>env</div>\n<div class=\"highlight\"><pre><span></span>DATABASE_URL=postgres://sjxgipufegmgsw:78cbb568e@ec2-52-4-104-184.compute-1.amazonaws.com:5432/dbmuget\n</pre></div>\n\n</pre>\n\n<p>The
  format for the postgres URL is as follows:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>postgresql://[user[:password]@][netloc][:port][/dbname]\n</pre></div>\n\n</pre>\n\n<p>We
  have now created environment variables for our django application and also added
  config vars in the heroku app, we now need a way to parse these environment variables
  into the Django project.</p>\n<h3>Parsing Environment variables using python-dotenv</h3>\n<p>We
  will use <a href=\"https://pypi.org/project/python-dotenv/\">python-dotenv</a> to
  parse variables into the django settings configurations like <code>SECRET_KEY</code>
  and <code>DATABASES</code>.</p>\n<ul>\n<li>Install <code>python-dotenv</code> with
  pip with the command :</li>\n</ul>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>pip install python-dotenv\n</pre></div>\n\n</pre>\n\n<p>We
  need to then modify the default variables in the <code>settings.py</code> file.
  Firstly, we will load in the <code>.env</code> file for accessing the environment
  variables for the configuration of the project.</p>\n<p>Append the following code,
  to the top of the <code>settings.py</code> file, make sure don't override the configuration
  so remove unnecessary imports and configurations.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># &lt;project_name&gt;/settings.py</span>\n\n<span class=\"kn\">import</span>
  <span class=\"nn\">os</span>\n<span class=\"kn\">from</span> <span class=\"nn\">dotenv</span>
  <span class=\"kn\">import</span> <span class=\"n\">load_dotenv</span>\n\n<span class=\"n\">BASE_DIR</span>
  <span class=\"o\">=</span> <span class=\"n\">Path</span><span class=\"p\">(</span><span
  class=\"vm\">__file__</span><span class=\"p\">)</span><span class=\"o\">.</span><span
  class=\"n\">resolve</span><span class=\"p\">()</span><span class=\"o\">.</span><span
  class=\"n\">parent</span><span class=\"o\">.</span><span class=\"n\">parent</span>\n\n<span
  class=\"n\">load_dotenv</span><span class=\"p\">(</span><span class=\"n\">os</span><span
  class=\"o\">.</span><span class=\"n\">path</span><span class=\"o\">.</span><span
  class=\"n\">join</span><span class=\"p\">(</span><span class=\"n\">BASE_DIR</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;.env&quot;</span><span class=\"p\">))</span>\n</pre></div>\n\n</pre>\n\n<p>We
  have imported the package <code>dotenv</code> basically the <code>python-dotenv</code>
  into the <code>settings.py</code> file and the module <code>os</code> for loading
  the <code>.env</code> file. The <code>load_dotenv</code> function helps in loading
  the <code>key-value</code> pairs which are the configuration variables that can
  act as actual environment variables. We provide in a file to the <a href=\"https://saurabh-kumar.com/python-dotenv/\">load_dotenv</a>
  function which is the <code>.env</code> file in our case, you can pick up any location
  for the <code>.env</code> file but make sure to change the location while parsing
  the file into the <code>load_dotenv</code> function.</p>\n<p>After loading the variables
  into the <code>settings.py</code> file, we now need to access those variables and
  set the appropriate variables the configuration from the variables received from
  the <code>load_dotenv</code> function. The <code>os.getenv</code> function to access
  the environment variables. The <code>os.getenv</code> function takes a parameter
  as the <code>key</code> for the environment variable and returns the value of the
  environment variable.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"n\">SECRET_KEY</span> <span class=\"o\">=</span> <span class=\"n\">os</span><span
  class=\"o\">.</span><span class=\"n\">getenv</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;SECRET_KEY&quot;</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>We
  have secretly configured the <code>SECRET_KEY</code> for the django project. If
  you have any other variables as simple key-value pairs like <code>AUTH</code> passwords,
  username, etc. you can use this method to get the configuration variables.</p>\n<h3>Loading
  Database configuration</h3>\n<p>Databases are a bit different as compared to other
  simple configurations in django. We need to make a few adjustments to the default
  database configuration. We will install another package <code>dj-database-url</code>
  to configure <code>DATABASE_URL</code>. Since the databse_url has a few components
  we need a way to extract the details like the <code>hostname</code>, <code>port</code>,
  <code>database_name</code>, and <code>password</code>. Using the <code>dj-database-url</code>
  package we have a few functions that can serve the purpose.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>pip install dj-database-url\n</pre></div>\n\n</pre>\n\n<p>At
  the end of your <code>settings.py</code> file, append the following code.</p>\n<pre
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
  class=\"kn\">import</span> <span class=\"nn\">dj_database_url</span>\n\n<span class=\"n\">DATABASE_URL</span>
  <span class=\"o\">=</span> <span class=\"n\">os</span><span class=\"o\">.</span><span
  class=\"n\">getenv</span><span class=\"p\">(</span><span class=\"s2\">&quot;DATABASE_URL&quot;</span><span
  class=\"p\">)</span>\n\n<span class=\"n\">DATABASES</span> <span class=\"o\">=</span>
  <span class=\"p\">{</span>\n    <span class=\"s2\">&quot;default&quot;</span><span
  class=\"p\">:</span> <span class=\"n\">dj_database_url</span><span class=\"o\">.</span><span
  class=\"n\">config</span><span class=\"p\">(</span><span class=\"n\">default</span><span
  class=\"o\">=</span><span class=\"n\">DATABASE_URL</span><span class=\"p\">,</span>
  <span class=\"n\">conn_max_age</span><span class=\"o\">=</span><span class=\"mi\">1800</span><span
  class=\"p\">),</span>\n<span class=\"p\">}</span>\n</pre></div>\n\n</pre>\n\n<p>We
  would need an adapter for making migrations to the <code>PostgreSQL</code> database
  i.e. the <code>psycopg2</code> package. This is a mandatory step if you are working
  with <code>postgres</code> database. This can be installed with the simple pip install:</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>pip install psycopg2\n\n#
  If it does not work try\npip install psycopg2-binary\n\n\n# if still error persists
  try installing setuptools\npip install -U setuptools\npip install psycopg2\n</pre></div>\n\n</pre>\n\n<p>Now,
  that we have configured the database, we can now apply migrations to the fresh database
  of postgres provided by heroku. We will simply run the migrate command and in the
  heroku app the PostgreSQL database would have been modified and an appropriate schema
  should be applied.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>python manage.py migrate\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652602284553/oTtGev28-.png\"
  alt=\"image.png\" /></p>\n<p>Make sure to update the <code>requirements.txt</code>
  file before pushing the project to Heroku for making sure everything works as expected.
  Since we have installed a few additional packages that are directly used in the
  <code>settings.py</code> file, we need to run the <code>pip freeze</code> command
  to update the <code>requiremnets.txt</code> file.</p>\n<h2>Serving Static Files</h2>\n<p>Now,
  if you have some static files like <code>CSS</code>, <code>Javascript</code>, or
  <code>images</code>, you need to configure the staticfiles in order to serve them
  from the heroku server. We will require another config variable for collecting the
  static files from the selected repository.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"n\">STATIC_URL</span> <span class=\"o\">=</span> <span class=\"s2\">&quot;static/&quot;</span>\n<span
  class=\"n\">STATICFILES_DIRS</span> <span class=\"o\">=</span> <span class=\"p\">[</span><span
  class=\"n\">os</span><span class=\"o\">.</span><span class=\"n\">path</span><span
  class=\"o\">.</span><span class=\"n\">join</span><span class=\"p\">(</span><span
  class=\"n\">BASE_DIR</span><span class=\"p\">,</span> <span class=\"s2\">&quot;static&quot;</span><span
  class=\"p\">)]</span>\n<span class=\"n\">STATIC_ROOT</span> <span class=\"o\">=</span>
  <span class=\"n\">os</span><span class=\"o\">.</span><span class=\"n\">path</span><span
  class=\"o\">.</span><span class=\"n\">join</span><span class=\"p\">(</span><span
  class=\"n\">BASE_DIR</span><span class=\"p\">,</span> <span class=\"s2\">&quot;staticfiles&quot;</span><span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>Here, if you have served your
  static files from the <code>static</code> folder in the root directory of your django
  project, you can add the above code in the <a href=\"http://settings.py\">settings.py</a>
  file. We will collect all static files in the folder along with the default static
  files provided by django in the <code>staticfiles</code> directory. Run the following
  command if you want to test whether the static files are properly collected and
  served.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n\n<div class=\"highlight\"><pre><span></span>python manage.py collectstatic
  \n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652600828657/JgL4nLuiL.png\"
  alt=\"image.png\" /></p>\n<p>So, this command will collect all the static files
  and store them in a single place. We see that the files from the admin section are
  also copied as well as the custom static files from the project configuration. Now,
  we can move on to set the config variable for the heroku app.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>DISABLE_COLLECTSTATIC = 0\n</pre></div>\n\n</pre>\n\n<p>We
  can set the <code>DISABLE_COLLECTSTATIC</code> variable as <code>0</code> or <code>1</code>
  indicating whether to disable it or not. We have currently enabled the static file
  collection while deploying the app but you can set it to <code>1</code> to disable
  the collection of static files.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652613798420/mbqzf1Kqd.png\"
  alt=\"image.png\" /></p>\n<p>Since I first tested the application on heroku, the
  static files don't work as expected, we need another package to make sure the staticfiles
  are served property. We will be installing the <code>whitenoise</code> package which
  serves as the middleware for serving the static files.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>pip install whitenoise\n</pre></div>\n\n</pre>\n\n<p>Add
  the whitenoise middleware <code>whitenoise.middleware.WhiteNoiseMiddleware</code>
  to the <code>MIDDLEWARE</code> list in the <code>settings.py</code> file.</p>\n<pre
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
  class=\"n\">MIDDLEWARE</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n<span
  class=\"o\">...</span>\n<span class=\"o\">...</span>\n<span class=\"o\">...</span>\n
  \   <span class=\"s1\">&#39;whitenoise.middleware.WhiteNoiseMiddleware&#39;</span><span
  class=\"p\">,</span>\n<span class=\"p\">]</span>\n\n<span class=\"err\">```</span>\n\n<span
  class=\"n\">That</span> <span class=\"n\">should</span> <span class=\"n\">be</span>
  <span class=\"n\">enough</span> <span class=\"n\">to</span> <span class=\"n\">make</span>
  <span class=\"n\">the</span> <span class=\"n\">most</span> <span class=\"n\">of</span>
  <span class=\"n\">the</span> <span class=\"n\">deployment</span> <span class=\"n\">on</span>
  <span class=\"n\">heroku</span><span class=\"o\">.</span> <span class=\"n\">You</span>
  <span class=\"n\">will</span> <span class=\"n\">have</span> <span class=\"n\">to</span>
  <span class=\"n\">make</span> <span class=\"n\">a</span> <span class=\"n\">few</span>
  <span class=\"n\">adjustments</span> <span class=\"k\">as</span> <span class=\"n\">per</span>
  <span class=\"n\">your</span> <span class=\"n\">requirements</span> <span class=\"ow\">and</span>
  <span class=\"n\">project</span><span class=\"o\">.</span>\n\n<span class=\"c1\">##
  Deploy from GitHub</span>\n\n<span class=\"n\">We</span> <span class=\"n\">are</span>
  <span class=\"n\">now</span> <span class=\"nb\">all</span> <span class=\"nb\">set</span>
  <span class=\"n\">to</span> <span class=\"n\">deploy</span> <span class=\"n\">the</span>
  <span class=\"n\">application</span> <span class=\"n\">on</span> <span class=\"n\">Heroku</span><span
  class=\"p\">,</span> <span class=\"n\">we</span> <span class=\"n\">can</span> <span
  class=\"n\">use</span> <span class=\"n\">the</span> <span class=\"err\">`</span><span
  class=\"n\">Connect</span> <span class=\"n\">to</span> <span class=\"n\">GitHub</span><span
  class=\"err\">`</span> <span class=\"ow\">or</span> <span class=\"err\">`</span><span
  class=\"n\">Heroku</span> <span class=\"n\">CLI</span><span class=\"err\">`</span>
  <span class=\"n\">to</span> <span class=\"n\">push</span> <span class=\"n\">the</span>
  <span class=\"n\">code</span> <span class=\"n\">into</span> <span class=\"n\">production</span><span
  class=\"o\">.</span> <span class=\"n\">Heroku</span> <span class=\"n\">CLI</span>
  <span class=\"ow\">is</span> <span class=\"n\">quite</span> <span class=\"n\">easy</span>
  <span class=\"k\">with</span> <span class=\"n\">a</span> <span class=\"n\">few</span>
  <span class=\"n\">sets</span> <span class=\"n\">of</span> <span class=\"n\">commands</span>
  <span class=\"n\">but</span> <span class=\"k\">if</span> <span class=\"n\">your</span>
  <span class=\"n\">project</span> <span class=\"ow\">is</span> <span class=\"n\">deployed</span>
  <span class=\"n\">on</span> <span class=\"n\">GitHub</span><span class=\"p\">,</span>
  <span class=\"n\">you</span> <span class=\"n\">can</span> <span class=\"n\">straightaway</span>
  <span class=\"n\">let</span> <span class=\"n\">the</span> <span class=\"n\">deployment</span>
  <span class=\"n\">start</span> <span class=\"n\">the</span> <span class=\"n\">build</span>
  <span class=\"n\">on</span> <span class=\"n\">a</span> <span class=\"n\">push</span>
  <span class=\"n\">to</span> <span class=\"n\">a</span> <span class=\"n\">specific</span>
  <span class=\"n\">branch</span><span class=\"o\">.</span> <span class=\"n\">This</span>
  <span class=\"n\">becomes</span> <span class=\"n\">quite</span> <span class=\"n\">automotive</span>
  <span class=\"ow\">and</span> <span class=\"n\">easy</span> <span class=\"n\">to</span>
  <span class=\"n\">scale</span> <span class=\"k\">while</span> <span class=\"n\">deploying</span>
  <span class=\"n\">a</span> <span class=\"n\">large</span><span class=\"o\">-</span><span
  class=\"n\">scale</span> <span class=\"n\">application</span><span class=\"o\">.</span>
  \n\n<span class=\"err\">```</span>\n<span class=\"n\">pip</span> <span class=\"n\">freeze</span>
  <span class=\"o\">&gt;</span> <span class=\"n\">requirements</span><span class=\"o\">.</span><span
  class=\"n\">txt</span>\n<span class=\"err\">```</span>\n\n<span class=\"n\">This</span>
  <span class=\"n\">step</span> <span class=\"ow\">is</span> <span class=\"n\">quite</span>
  <span class=\"n\">important</span> <span class=\"n\">because</span> <span class=\"n\">you</span>
  <span class=\"n\">need</span> <span class=\"n\">to</span> <span class=\"n\">make</span>
  <span class=\"n\">sure</span> <span class=\"n\">that</span> <span class=\"nb\">all</span>
  <span class=\"n\">the</span> <span class=\"n\">packages</span> <span class=\"n\">are</span>
  <span class=\"n\">listed</span> <span class=\"ow\">in</span> <span class=\"n\">the</span>
  <span class=\"err\">`</span><span class=\"n\">requirements</span><span class=\"o\">.</span><span
  class=\"n\">txt</span><span class=\"err\">`</span> <span class=\"n\">file</span>
  <span class=\"k\">else</span> <span class=\"n\">you</span> <span class=\"n\">will</span>
  <span class=\"n\">have</span> <span class=\"n\">to</span> <span class=\"n\">wait</span>
  <span class=\"k\">for</span> <span class=\"n\">the</span> <span class=\"n\">build</span>
  <span class=\"n\">to</span> <span class=\"n\">fail</span> <span class=\"ow\">and</span>
  <span class=\"n\">redeploy</span><span class=\"o\">.</span>\n\n<span class=\"n\">Make</span>
  <span class=\"n\">sure</span> <span class=\"n\">the</span> <span class=\"n\">server</span>
  <span class=\"ow\">is</span> <span class=\"n\">running</span> <span class=\"n\">first</span>
  <span class=\"n\">on</span> <span class=\"n\">your</span> <span class=\"n\">local</span>
  <span class=\"n\">machine</span><span class=\"p\">,</span> <span class=\"n\">remember</span>
  <span class=\"n\">the</span> <span class=\"n\">server</span> <span class=\"n\">will</span>
  <span class=\"n\">be</span> <span class=\"nb\">set</span> <span class=\"n\">up</span>
  <span class=\"kn\">from</span> <span class=\"nn\">scratch</span> <span class=\"n\">but</span>
  <span class=\"n\">the</span> <span class=\"n\">database</span> <span class=\"n\">will</span>
  <span class=\"n\">already</span> <span class=\"n\">have</span> <span class=\"n\">applied</span>
  <span class=\"n\">migrations</span> <span class=\"k\">if</span> <span class=\"n\">you</span>
  <span class=\"n\">have</span> <span class=\"n\">applied</span> <span class=\"n\">migrations</span>
  <span class=\"n\">before</span> <span class=\"n\">after</span> <span class=\"n\">connecting</span>
  <span class=\"n\">the</span> <span class=\"n\">Heroku</span> <span class=\"n\">Postgres</span>
  <span class=\"n\">database</span><span class=\"o\">.</span>\n \n<span class=\"err\">```</span>\n<span
  class=\"n\">python</span> <span class=\"n\">manage</span><span class=\"o\">.</span><span
  class=\"n\">py</span> <span class=\"n\">collectstatic</span>\n\n<span class=\"n\">python</span>
  <span class=\"n\">manage</span><span class=\"o\">.</span><span class=\"n\">py</span>
  <span class=\"n\">runserver</span>\n<span class=\"err\">```</span>\n\n<span class=\"n\">This</span>
  <span class=\"n\">will</span> <span class=\"nb\">set</span> <span class=\"n\">up</span>
  <span class=\"n\">the</span> <span class=\"n\">origin</span> <span class=\"n\">of</span>
  <span class=\"n\">the</span> <span class=\"n\">remote</span> <span class=\"n\">repository</span>
  <span class=\"n\">that</span> <span class=\"n\">will</span> <span class=\"n\">be</span>
  <span class=\"n\">pushing</span> <span class=\"n\">the</span> <span class=\"n\">project</span>
  <span class=\"n\">code</span><span class=\"o\">.</span> <span class=\"n\">Next</span><span
  class=\"p\">,</span> <span class=\"n\">make</span> <span class=\"n\">sure</span>
  <span class=\"n\">to</span> <span class=\"n\">commit</span> <span class=\"n\">the</span>
  <span class=\"n\">code</span> <span class=\"n\">which</span> <span class=\"n\">will</span>
  <span class=\"n\">contain</span> <span class=\"nb\">all</span> <span class=\"n\">the</span>
  <span class=\"n\">required</span> <span class=\"n\">stuff</span> <span class=\"k\">for</span>
  <span class=\"n\">deploying</span> <span class=\"n\">the</span> <span class=\"n\">code</span><span
  class=\"o\">.</span>\n\n<span class=\"n\">Checklist</span> <span class=\"k\">for</span>
  <span class=\"n\">deploying</span> <span class=\"n\">the</span> <span class=\"n\">code</span>\n\n<span
  class=\"err\">```</span>\n<span class=\"o\">-</span> <span class=\"n\">requirements</span><span
  class=\"o\">.</span><span class=\"n\">txt</span>\n<span class=\"o\">-</span> <span
  class=\"n\">Procfile</span>\n<span class=\"o\">-</span> <span class=\"n\">runtime</span><span
  class=\"o\">.</span><span class=\"n\">txt</span>\n<span class=\"o\">-</span> <span
  class=\"n\">django</span><span class=\"o\">-</span><span class=\"n\">project</span>\n<span
  class=\"o\">-</span> <span class=\"n\">environment</span> <span class=\"n\">variables</span>
  <span class=\"o\">/</span> <span class=\"n\">config</span> <span class=\"n\">variables</span>
  \n<span class=\"o\">-</span> <span class=\"n\">static</span> <span class=\"n\">file</span>
  <span class=\"n\">configuration</span>\n<span class=\"o\">-</span> <span class=\"n\">database</span>
  <span class=\"n\">configuration</span>\n<span class=\"o\">-</span> <span class=\"n\">migrate</span>
  <span class=\"n\">schema</span> <span class=\"n\">of</span> <span class=\"n\">database</span>
  \n<span class=\"o\">-</span> <span class=\"n\">gitignore</span> <span class=\"n\">file</span>
  <span class=\"k\">for</span> <span class=\"n\">ignoring</span> <span class=\"n\">virtualenvs</span><span
  class=\"p\">,</span> <span class=\"o\">.</span><span class=\"n\">env</span> <span
  class=\"n\">file</span><span class=\"p\">,</span> <span class=\"n\">staticfiles</span><span
  class=\"p\">,</span> <span class=\"n\">etc</span>\n<span class=\"err\">```</span>\n\n<span
  class=\"n\">here</span><span class=\"s1\">&#39;s a sample `.gitignore` for a minimal
  django project.</span>\n\n<span class=\"err\">```</span><span class=\"n\">gitignore</span>\n<span
  class=\"o\">.</span><span class=\"n\">env</span><span class=\"o\">/</span>\n<span
  class=\"o\">.</span><span class=\"n\">venv</span><span class=\"o\">/</span>\n<span
  class=\"n\">env</span><span class=\"o\">/</span>\n<span class=\"n\">venv</span><span
  class=\"o\">/</span>\n<span class=\"o\">*.</span><span class=\"n\">env</span>\n\n<span
  class=\"o\">*.</span><span class=\"n\">pyc</span>\n<span class=\"n\">db</span><span
  class=\"o\">.</span><span class=\"n\">sqlite3</span>\n<span class=\"n\">staticfiles</span><span
  class=\"o\">/</span>\n<span class=\"err\">```</span>\n\n<span class=\"n\">If</span>
  <span class=\"n\">you</span> <span class=\"n\">want</span> <span class=\"n\">a</span>
  <span class=\"n\">full</span><span class=\"o\">-</span><span class=\"n\">fledged</span>
  <span class=\"err\">`</span><span class=\"o\">.</span><span class=\"n\">gitignore</span><span
  class=\"err\">`</span> <span class=\"k\">for</span> <span class=\"n\">a</span> <span
  class=\"nb\">complex</span> <span class=\"n\">django</span> <span class=\"n\">project</span><span
  class=\"p\">,</span> <span class=\"n\">you</span> <span class=\"n\">can</span> <span
  class=\"n\">take</span> <span class=\"n\">the</span> <span class=\"n\">reference</span>
  <span class=\"kn\">from</span> <span class=\"nn\">Jose</span> <span class=\"n\">Padilla</span><span
  class=\"s1\">&#39;s [gitignore Template](https://github.com/jpadilla/django-project-template/blob/master/.gitignore)
  for a django project.  </span>\n\n<span class=\"c1\">### Git Commit the Django Project</span>\n<span
  class=\"err\">```</span>\n<span class=\"n\">git</span> <span class=\"n\">status</span>
  \n\n<span class=\"n\">git</span> <span class=\"n\">add</span> <span class=\"o\">.</span>\n\n<span
  class=\"n\">git</span> <span class=\"n\">commit</span> <span class=\"o\">-</span><span
  class=\"n\">m</span> <span class=\"s2\">&quot;config for heroku deployment&quot;</span>\n<span
  class=\"err\">```</span>\n<span class=\"n\">Carefully</span> <span class=\"n\">check</span>
  <span class=\"n\">the</span> <span class=\"n\">status</span> <span class=\"n\">of</span>
  <span class=\"n\">the</span> <span class=\"n\">git</span> <span class=\"n\">repository</span>
  <span class=\"n\">before</span> <span class=\"n\">committing</span> <span class=\"ow\">and</span>
  <span class=\"n\">make</span> <span class=\"n\">sure</span> <span class=\"n\">you</span>
  <span class=\"n\">don</span><span class=\"s1\">&#39;t forget anything by mistake,
  it won&#39;</span><span class=\"n\">t</span> <span class=\"n\">a</span> <span class=\"n\">big</span>
  <span class=\"n\">problem</span> <span class=\"n\">but</span> <span class=\"n\">it</span>
  <span class=\"n\">would</span> <span class=\"n\">mess</span> <span class=\"n\">up</span>
  <span class=\"n\">the</span> <span class=\"n\">build</span> <span class=\"n\">process</span><span
  class=\"o\">.</span>\n\n\n<span class=\"err\">!</span><span class=\"p\">[</span><span
  class=\"n\">image</span><span class=\"o\">.</span><span class=\"n\">png</span><span
  class=\"p\">](</span><span class=\"n\">https</span><span class=\"p\">:</span><span
  class=\"o\">//</span><span class=\"n\">cdn</span><span class=\"o\">.</span><span
  class=\"n\">hashnode</span><span class=\"o\">.</span><span class=\"n\">com</span><span
  class=\"o\">/</span><span class=\"n\">res</span><span class=\"o\">/</span><span
  class=\"n\">hashnode</span><span class=\"o\">/</span><span class=\"n\">image</span><span
  class=\"o\">/</span><span class=\"n\">upload</span><span class=\"o\">/</span><span
  class=\"n\">v1652613991525</span><span class=\"o\">/</span><span class=\"n\">hxQgtGOoM</span><span
  class=\"o\">.</span><span class=\"n\">png</span><span class=\"p\">)</span>\n\n<span
  class=\"n\">After</span> <span class=\"n\">committing</span> <span class=\"n\">the</span>
  <span class=\"n\">code</span><span class=\"p\">,</span> <span class=\"n\">we</span>
  <span class=\"n\">can</span> <span class=\"n\">now</span> <span class=\"n\">push</span>
  <span class=\"n\">the</span> <span class=\"n\">code</span> <span class=\"n\">to</span>
  <span class=\"n\">GitHub</span><span class=\"o\">.</span> <span class=\"n\">We</span>
  <span class=\"n\">first</span> <span class=\"n\">need</span> <span class=\"n\">to</span>
  <span class=\"nb\">set</span> <span class=\"n\">the</span> <span class=\"n\">remote</span>
  <span class=\"n\">repository</span> <span class=\"n\">reference</span> <span class=\"n\">to</span>
  <span class=\"n\">be</span> <span class=\"n\">able</span> <span class=\"n\">to</span>
  <span class=\"n\">push</span> <span class=\"n\">the</span> <span class=\"n\">code</span>
  <span class=\"n\">to</span> <span class=\"n\">it</span><span class=\"o\">.</span>
  \n\n<span class=\"err\">```</span>\n<span class=\"n\">git</span> <span class=\"n\">remote</span>
  <span class=\"n\">add</span> <span class=\"n\">origin</span> <span class=\"n\">https</span><span
  class=\"p\">:</span><span class=\"o\">//</span><span class=\"n\">github</span><span
  class=\"o\">.</span><span class=\"n\">com</span><span class=\"o\">/&lt;</span><span
  class=\"n\">username</span><span class=\"o\">&gt;/&lt;</span><span class=\"n\">repo_name</span><span
  class=\"o\">&gt;</span>\n<span class=\"err\">```</span>\n<span class=\"n\">This</span>
  <span class=\"n\">will</span> <span class=\"nb\">set</span> <span class=\"n\">up</span>
  <span class=\"n\">the</span> <span class=\"err\">`</span><span class=\"n\">origin</span><span
  class=\"err\">`</span> <span class=\"k\">as</span> <span class=\"n\">the</span>
  <span class=\"n\">remote</span> <span class=\"n\">repository</span> <span class=\"n\">on</span>
  <span class=\"n\">GitHub</span><span class=\"o\">.</span> <span class=\"n\">Once</span>
  <span class=\"n\">the</span> <span class=\"n\">remote</span> <span class=\"n\">repository</span>
  <span class=\"ow\">is</span> <span class=\"n\">created</span><span class=\"p\">,</span>
  <span class=\"n\">we</span> <span class=\"n\">can</span> <span class=\"n\">move</span>
  <span class=\"n\">into</span> <span class=\"n\">the</span> <span class=\"n\">push</span>
  <span class=\"n\">of</span> <span class=\"n\">the</span> <span class=\"n\">git</span>
  <span class=\"n\">repository</span> <span class=\"n\">that</span> <span class=\"n\">will</span>
  <span class=\"n\">trigger</span> <span class=\"n\">the</span> <span class=\"n\">build</span><span
  class=\"o\">.</span> <span class=\"n\">First</span><span class=\"p\">,</span> <span
  class=\"n\">navigate</span> <span class=\"n\">to</span> <span class=\"n\">the</span>
  <span class=\"err\">`</span><span class=\"n\">Deploy</span><span class=\"err\">`</span>
  <span class=\"n\">section</span> <span class=\"ow\">in</span> <span class=\"n\">the</span>
  <span class=\"n\">heroku</span> <span class=\"n\">app</span><span class=\"s1\">&#39;s
  dashboard where we want to connect the `GitHub` repository and allow the automatic
  deploy from a branch in this case we have chosen the `main` branch.</span>\n\n<span
  class=\"n\">Due</span> <span class=\"n\">to</span> <span class=\"n\">some</span>
  <span class=\"err\">`</span><span class=\"n\">Heroku</span><span class=\"err\">`</span>
  <span class=\"n\">Internal</span> <span class=\"n\">Server</span> <span class=\"n\">Issues</span><span
  class=\"p\">,</span> <span class=\"n\">the</span> <span class=\"n\">GitHub</span>
  <span class=\"n\">integration</span> <span class=\"n\">seems</span> <span class=\"n\">to</span>
  <span class=\"n\">have</span> <span class=\"n\">broken</span> <span class=\"ow\">and</span>
  <span class=\"n\">isn</span><span class=\"s1\">&#39;t working as of May 2022, but
  it might work later. </span>\n\n<span class=\"n\">Below</span> <span class=\"ow\">is</span>
  <span class=\"n\">a</span> <span class=\"n\">screenshot</span> <span class=\"n\">of</span>
  <span class=\"n\">my</span> <span class=\"n\">previous</span> <span class=\"n\">project</span>
  <span class=\"n\">deployed</span> <span class=\"n\">to</span> <span class=\"n\">Heroku</span>
  <span class=\"n\">using</span> <span class=\"n\">a</span> <span class=\"n\">GitHub</span>
  <span class=\"n\">repository</span><span class=\"o\">.</span>\n\n<span class=\"err\">!</span><span
  class=\"p\">[</span><span class=\"n\">image</span><span class=\"o\">.</span><span
  class=\"n\">png</span><span class=\"p\">](</span><span class=\"n\">https</span><span
  class=\"p\">:</span><span class=\"o\">//</span><span class=\"n\">cdn</span><span
  class=\"o\">.</span><span class=\"n\">hashnode</span><span class=\"o\">.</span><span
  class=\"n\">com</span><span class=\"o\">/</span><span class=\"n\">res</span><span
  class=\"o\">/</span><span class=\"n\">hashnode</span><span class=\"o\">/</span><span
  class=\"n\">image</span><span class=\"o\">/</span><span class=\"n\">upload</span><span
  class=\"o\">/</span><span class=\"n\">v1652605497382</span><span class=\"o\">/</span><span
  class=\"mi\">5</span><span class=\"n\">VuQUQ0t0</span><span class=\"o\">.</span><span
  class=\"n\">png</span><span class=\"p\">)</span>\n\n<span class=\"err\">```</span>\n<span
  class=\"n\">git</span> <span class=\"n\">push</span> <span class=\"n\">origin</span>
  <span class=\"n\">main</span>\n<span class=\"err\">```</span>\n\n<span class=\"n\">This</span>
  <span class=\"n\">will</span> <span class=\"n\">prompt</span> <span class=\"n\">you</span>
  <span class=\"k\">for</span> <span class=\"n\">your</span> <span class=\"n\">GitHub</span>
  <span class=\"n\">credentials</span> <span class=\"ow\">and</span> <span class=\"n\">will</span>
  <span class=\"n\">deploy</span> <span class=\"n\">the</span> <span class=\"n\">commits</span>
  <span class=\"n\">to</span> <span class=\"n\">the</span> <span class=\"n\">remote</span>
  <span class=\"n\">repository</span> <span class=\"n\">on</span> <span class=\"n\">GitHub</span><span
  class=\"o\">.</span> <span class=\"n\">This</span> <span class=\"n\">push</span>
  <span class=\"n\">on</span> <span class=\"n\">the</span> <span class=\"n\">main</span>
  <span class=\"n\">branch</span> <span class=\"n\">should</span> <span class=\"n\">also</span>
  <span class=\"n\">trigger</span> <span class=\"n\">the</span> <span class=\"n\">build</span>
  <span class=\"n\">process</span> <span class=\"n\">of</span> <span class=\"n\">the</span>
  <span class=\"n\">heroku</span> <span class=\"n\">app</span> <span class=\"k\">for</span>
  <span class=\"n\">this</span> <span class=\"n\">django</span> <span class=\"n\">project</span><span
  class=\"o\">.</span> <span class=\"n\">You</span> <span class=\"n\">can</span> <span
  class=\"n\">navigate</span> <span class=\"n\">to</span> <span class=\"n\">the</span>
  <span class=\"n\">Activity</span> <span class=\"n\">section</span> <span class=\"k\">for</span>
  <span class=\"n\">the</span> <span class=\"n\">Build</span> <span class=\"n\">logs</span><span
  class=\"o\">.</span> \n\n<span class=\"n\">If</span> <span class=\"n\">you</span>
  <span class=\"n\">have</span> <span class=\"n\">followed</span> <span class=\"n\">the</span>
  <span class=\"n\">article</span> <span class=\"n\">well</span><span class=\"p\">,</span>
  <span class=\"ow\">and</span> <span class=\"n\">your</span> <span class=\"n\">repository</span>
  <span class=\"n\">has</span> <span class=\"nb\">all</span> <span class=\"n\">the</span>
  <span class=\"n\">correct</span> <span class=\"n\">configurations</span><span class=\"p\">,</span>
  <span class=\"n\">the</span> <span class=\"n\">build</span> <span class=\"n\">will</span>
  <span class=\"n\">succeed</span><span class=\"p\">,</span> <span class=\"k\">else</span>
  <span class=\"n\">chances</span> <span class=\"n\">are</span> <span class=\"n\">you</span>
  <span class=\"n\">might</span> <span class=\"n\">have</span> <span class=\"n\">missed</span>
  <span class=\"n\">a</span> <span class=\"n\">few</span> <span class=\"n\">things</span>
  <span class=\"ow\">and</span> <span class=\"n\">the</span> <span class=\"n\">app</span>
  <span class=\"n\">might</span> <span class=\"n\">have</span> <span class=\"n\">crashed</span><span
  class=\"o\">.</span> <span class=\"n\">You</span> <span class=\"n\">can</span> <span
  class=\"n\">debug</span> <span class=\"n\">your</span> <span class=\"n\">application</span>
  <span class=\"n\">build</span> <span class=\"k\">with</span> <span class=\"n\">the</span>
  <span class=\"n\">simple</span> <span class=\"n\">heroku</span> <span class=\"n\">CLI</span>
  <span class=\"n\">command</span><span class=\"p\">:</span>\n\n<span class=\"err\">```</span>\n<span
  class=\"n\">heroku</span> <span class=\"n\">logs</span> <span class=\"o\">--</span><span
  class=\"n\">tail</span> <span class=\"o\">-</span><span class=\"n\">a</span> <span
  class=\"o\">&lt;</span><span class=\"n\">app_name</span><span class=\"o\">&gt;</span>\n<span
  class=\"err\">```</span>\n\n<span class=\"n\">This</span> <span class=\"n\">can</span>
  <span class=\"n\">be</span> <span class=\"n\">quite</span> <span class=\"n\">handy</span>
  <span class=\"ow\">and</span> <span class=\"n\">saves</span> <span class=\"n\">a</span>
  <span class=\"n\">lot</span> <span class=\"n\">of</span> <span class=\"n\">time</span>
  <span class=\"ow\">in</span> <span class=\"n\">understanding</span> <span class=\"n\">what</span>
  <span class=\"n\">went</span> <span class=\"n\">wrong</span> <span class=\"ow\">in</span>
  <span class=\"n\">the</span> <span class=\"n\">build</span><span class=\"o\">.</span>
  <span class=\"n\">It</span> <span class=\"n\">might</span> <span class=\"n\">be</span>
  <span class=\"n\">related</span> <span class=\"n\">to</span> <span class=\"n\">database</span>
  <span class=\"n\">migration</span><span class=\"p\">,</span> <span class=\"n\">static</span>
  <span class=\"n\">files</span><span class=\"p\">,</span> <span class=\"n\">python</span>
  <span class=\"n\">package</span> <span class=\"ow\">not</span> <span class=\"n\">found</span><span
  class=\"p\">,</span> <span class=\"ow\">and</span> <span class=\"n\">some</span>
  <span class=\"n\">silly</span> <span class=\"n\">mistakes</span> <span class=\"ow\">and</span>
  <span class=\"n\">errors</span> <span class=\"n\">that</span> <span class=\"n\">can</span>
  <span class=\"n\">be</span> <span class=\"n\">fixed</span> <span class=\"n\">after</span>
  <span class=\"n\">committing</span> <span class=\"n\">the</span> <span class=\"n\">code</span>
  <span class=\"ow\">and</span> <span class=\"n\">pushing</span> <span class=\"n\">it</span>
  <span class=\"n\">to</span> <span class=\"n\">GitHub</span> <span class=\"n\">again</span><span
  class=\"o\">.</span>\n\n<span class=\"n\">If</span> <span class=\"n\">you</span>
  <span class=\"n\">do</span> <span class=\"ow\">not</span> <span class=\"n\">want</span>
  <span class=\"n\">a</span> <span class=\"n\">GitHub</span> <span class=\"n\">repository</span><span
  class=\"p\">,</span> <span class=\"n\">you</span> <span class=\"n\">can</span> <span
  class=\"n\">directly</span> <span class=\"n\">push</span> <span class=\"n\">the</span>
  <span class=\"n\">code</span> <span class=\"kn\">from</span> <span class=\"nn\">the</span>
  <span class=\"n\">local</span> <span class=\"n\">git</span> <span class=\"n\">repository</span>
  <span class=\"n\">to</span> <span class=\"n\">the</span> <span class=\"n\">remote</span>
  <span class=\"n\">heroku</span> <span class=\"n\">app</span> <span class=\"n\">center</span><span
  class=\"o\">.</span> <span class=\"n\">This</span> <span class=\"n\">will</span>
  <span class=\"n\">require</span> <span class=\"n\">us</span> <span class=\"n\">the</span>
  <span class=\"n\">Heroku</span> <span class=\"n\">CLI</span><span class=\"o\">.</span>\n\n<span
  class=\"c1\">## Heroku CLI</span>\n\n<span class=\"n\">We</span> <span class=\"n\">can</span>
  <span class=\"n\">use</span> <span class=\"n\">the</span> <span class=\"n\">heroku</span>
  <span class=\"n\">CLI</span> <span class=\"k\">for</span> <span class=\"n\">pushing</span>
  <span class=\"n\">the</span> <span class=\"n\">code</span> <span class=\"n\">via</span>
  <span class=\"n\">the</span> <span class=\"n\">simple</span> <span class=\"n\">git</span>
  <span class=\"n\">repository</span><span class=\"o\">.</span> <span class=\"n\">We</span>
  <span class=\"n\">can</span> <span class=\"n\">push</span> <span class=\"n\">the</span>
  <span class=\"n\">references</span> <span class=\"n\">via</span> <span class=\"n\">the</span>
  <span class=\"n\">branch</span> <span class=\"ow\">and</span> <span class=\"n\">a</span>
  <span class=\"n\">remote</span> <span class=\"n\">repository</span> <span class=\"n\">on</span>
  <span class=\"n\">heroku</span> <span class=\"n\">to</span> <span class=\"n\">build</span>
  <span class=\"n\">our</span> <span class=\"n\">app</span><span class=\"o\">.</span>
  \ <span class=\"n\">For</span> <span class=\"n\">this</span><span class=\"p\">,</span>
  <span class=\"n\">we</span> <span class=\"n\">assume</span> <span class=\"n\">you</span>
  <span class=\"n\">have</span> <span class=\"n\">heroku</span> <span class=\"n\">installed</span>
  <span class=\"ow\">and</span> <span class=\"n\">logged</span> <span class=\"ow\">in</span><span
  class=\"o\">.</span> <span class=\"n\">We</span> <span class=\"n\">will</span> <span
  class=\"n\">require</span> <span class=\"n\">the</span> <span class=\"n\">django</span>
  <span class=\"n\">project</span> <span class=\"n\">code</span> <span class=\"ow\">and</span>
  <span class=\"n\">heroku</span> <span class=\"n\">cli</span> <span class=\"n\">to</span>
  <span class=\"n\">build</span> <span class=\"n\">the</span> <span class=\"n\">django</span>
  <span class=\"n\">web</span> <span class=\"n\">application</span><span class=\"o\">.</span>\n\n<span
  class=\"err\">```</span><span class=\"n\">bash</span>\n<span class=\"n\">heroku</span>
  <span class=\"n\">git</span><span class=\"p\">:</span><span class=\"n\">remote</span>
  <span class=\"o\">-</span><span class=\"n\">a</span> <span class=\"o\">&lt;</span><span
  class=\"n\">heroku_app_name</span><span class=\"o\">&gt;</span>\n\n<span class=\"c1\">#
  for my case</span>\n<span class=\"n\">heroku</span> <span class=\"n\">git</span><span
  class=\"p\">:</span><span class=\"n\">remote</span> <span class=\"o\">-</span><span
  class=\"n\">a</span> <span class=\"n\">blog</span><span class=\"o\">-</span><span
  class=\"n\">django</span><span class=\"o\">-</span><span class=\"n\">dep</span>\n<span
  class=\"err\">```</span>\n\n<span class=\"err\">!</span><span class=\"p\">[</span><span
  class=\"n\">image</span><span class=\"o\">.</span><span class=\"n\">png</span><span
  class=\"p\">](</span><span class=\"n\">https</span><span class=\"p\">:</span><span
  class=\"o\">//</span><span class=\"n\">cdn</span><span class=\"o\">.</span><span
  class=\"n\">hashnode</span><span class=\"o\">.</span><span class=\"n\">com</span><span
  class=\"o\">/</span><span class=\"n\">res</span><span class=\"o\">/</span><span
  class=\"n\">hashnode</span><span class=\"o\">/</span><span class=\"n\">image</span><span
  class=\"o\">/</span><span class=\"n\">upload</span><span class=\"o\">/</span><span
  class=\"n\">v1652614221069</span><span class=\"o\">/</span><span class=\"n\">vCAKD0zsz</span><span
  class=\"o\">.</span><span class=\"n\">png</span><span class=\"p\">)</span>\n\n<span
  class=\"n\">After</span> <span class=\"n\">this</span><span class=\"p\">,</span>
  <span class=\"n\">you</span> <span class=\"n\">can</span> <span class=\"n\">commit</span>
  <span class=\"n\">your</span> <span class=\"n\">code</span> <span class=\"ow\">and</span>
  <span class=\"n\">the</span> <span class=\"n\">project</span> <span class=\"k\">as</span>
  <span class=\"n\">git</span> <span class=\"n\">repository</span><span class=\"o\">.</span>
  <span class=\"n\">We</span> <span class=\"n\">have</span> <span class=\"n\">added</span>
  <span class=\"n\">the</span> <span class=\"n\">remote</span> <span class=\"n\">repository</span>
  <span class=\"n\">location</span> <span class=\"n\">on</span> <span class=\"n\">heroku</span><span
  class=\"p\">,</span> <span class=\"n\">we</span> <span class=\"n\">can</span> <span
  class=\"n\">now</span> <span class=\"n\">simply</span> <span class=\"n\">push</span>
  <span class=\"n\">the</span> <span class=\"n\">code</span> <span class=\"n\">to</span>
  <span class=\"n\">the</span> <span class=\"n\">remote</span> <span class=\"n\">repository</span><span
  class=\"o\">.</span>\n\n<span class=\"err\">```</span>\n<span class=\"n\">git</span>
  <span class=\"n\">push</span> <span class=\"n\">heroku</span> <span class=\"n\">main</span>\n<span
  class=\"err\">```</span>\n\n<span class=\"err\">!</span><span class=\"p\">[</span><span
  class=\"n\">image</span><span class=\"o\">.</span><span class=\"n\">png</span><span
  class=\"p\">](</span><span class=\"n\">https</span><span class=\"p\">:</span><span
  class=\"o\">//</span><span class=\"n\">cdn</span><span class=\"o\">.</span><span
  class=\"n\">hashnode</span><span class=\"o\">.</span><span class=\"n\">com</span><span
  class=\"o\">/</span><span class=\"n\">res</span><span class=\"o\">/</span><span
  class=\"n\">hashnode</span><span class=\"o\">/</span><span class=\"n\">image</span><span
  class=\"o\">/</span><span class=\"n\">upload</span><span class=\"o\">/</span><span
  class=\"n\">v1652614125785</span><span class=\"o\">/</span><span class=\"n\">uEzFQ9VvQ</span><span
  class=\"o\">.</span><span class=\"n\">png</span><span class=\"p\">)</span>\n\n<span
  class=\"n\">Here</span><span class=\"p\">,</span> <span class=\"err\">`</span><span
  class=\"n\">heroku</span><span class=\"err\">`</span> <span class=\"ow\">is</span>
  <span class=\"n\">the</span> <span class=\"n\">remote</span> <span class=\"n\">repository</span>
  <span class=\"n\">location</span> <span class=\"ow\">and</span> <span class=\"err\">`</span><span
  class=\"n\">main</span><span class=\"err\">`</span> <span class=\"ow\">is</span>
  <span class=\"n\">the</span> <span class=\"n\">branch</span> <span class=\"n\">of</span>
  <span class=\"n\">the</span> <span class=\"n\">repository</span><span class=\"o\">.</span>
  <span class=\"n\">This</span> <span class=\"n\">will</span> <span class=\"n\">push</span>
  <span class=\"n\">the</span> <span class=\"n\">code</span> <span class=\"n\">to</span>
  <span class=\"n\">the</span> <span class=\"n\">repository</span> <span class=\"ow\">and</span>
  <span class=\"n\">thereafter</span> <span class=\"n\">create</span> <span class=\"n\">a</span>
  <span class=\"n\">build</span> <span class=\"n\">to</span> <span class=\"n\">deploy</span>
  <span class=\"n\">the</span> <span class=\"n\">Django</span> <span class=\"n\">project</span>
  <span class=\"k\">as</span> <span class=\"n\">a</span> <span class=\"n\">Heroku</span>
  <span class=\"n\">application</span><span class=\"o\">.</span>\n\n<span class=\"err\">!</span><span
  class=\"p\">[</span><span class=\"n\">image</span><span class=\"o\">.</span><span
  class=\"n\">png</span><span class=\"p\">](</span><span class=\"n\">https</span><span
  class=\"p\">:</span><span class=\"o\">//</span><span class=\"n\">cdn</span><span
  class=\"o\">.</span><span class=\"n\">hashnode</span><span class=\"o\">.</span><span
  class=\"n\">com</span><span class=\"o\">/</span><span class=\"n\">res</span><span
  class=\"o\">/</span><span class=\"n\">hashnode</span><span class=\"o\">/</span><span
  class=\"n\">image</span><span class=\"o\">/</span><span class=\"n\">upload</span><span
  class=\"o\">/</span><span class=\"n\">v1652614381808</span><span class=\"o\">/</span><span
  class=\"n\">kYTmB3EO2p</span><span class=\"o\">.</span><span class=\"n\">png</span><span
  class=\"p\">)</span>\n\n<span class=\"n\">You</span> <span class=\"n\">can</span>
  <span class=\"n\">hit</span> <span class=\"n\">the</span> <span class=\"err\">`</span><span
  class=\"n\">Open</span> <span class=\"n\">App</span><span class=\"err\">`</span>
  <span class=\"n\">button</span> <span class=\"n\">on</span> <span class=\"n\">the</span>
  <span class=\"n\">top</span> <span class=\"n\">right</span> <span class=\"n\">corner</span>
  <span class=\"ow\">and</span> <span class=\"n\">there</span> <span class=\"n\">should</span>
  <span class=\"n\">be</span> <span class=\"n\">your</span> <span class=\"n\">deployed</span>
  <span class=\"n\">Django</span> <span class=\"n\">application</span><span class=\"o\">.</span>
  \n\n<span class=\"err\">!</span><span class=\"p\">[</span><span class=\"n\">image</span><span
  class=\"o\">.</span><span class=\"n\">png</span><span class=\"p\">](</span><span
  class=\"n\">https</span><span class=\"p\">:</span><span class=\"o\">//</span><span
  class=\"n\">cdn</span><span class=\"o\">.</span><span class=\"n\">hashnode</span><span
  class=\"o\">.</span><span class=\"n\">com</span><span class=\"o\">/</span><span
  class=\"n\">res</span><span class=\"o\">/</span><span class=\"n\">hashnode</span><span
  class=\"o\">/</span><span class=\"n\">image</span><span class=\"o\">/</span><span
  class=\"n\">upload</span><span class=\"o\">/</span><span class=\"n\">v1652610395538</span><span
  class=\"o\">/</span><span class=\"n\">xjUiODhoK</span><span class=\"o\">.</span><span
  class=\"n\">png</span><span class=\"p\">)</span>\n\n\n<span class=\"c1\">## A Few
  Tricks and Gotchas</span>\n\n<span class=\"n\">There</span> <span class=\"n\">are</span>
  <span class=\"n\">a</span> <span class=\"n\">few</span> <span class=\"n\">tricks</span>
  <span class=\"ow\">and</span> <span class=\"n\">issues</span> <span class=\"n\">that</span>
  <span class=\"n\">you</span> <span class=\"n\">might</span> <span class=\"n\">pop</span>
  <span class=\"n\">into</span> <span class=\"k\">while</span> <span class=\"n\">deploying</span>
  <span class=\"n\">a</span> <span class=\"n\">django</span> <span class=\"n\">project</span>
  <span class=\"n\">on</span> <span class=\"n\">heroku</span><span class=\"p\">,</span>
  <span class=\"n\">especially</span> <span class=\"n\">the</span> <span class=\"n\">build</span>
  <span class=\"n\">process</span><span class=\"o\">.</span> <span class=\"n\">It</span>
  <span class=\"n\">requires</span> <span class=\"n\">a</span> <span class=\"n\">few</span>
  <span class=\"n\">iterations</span> <span class=\"n\">to</span> <span class=\"n\">get</span>
  <span class=\"n\">the</span> <span class=\"n\">complete</span> <span class=\"n\">app</span>
  <span class=\"n\">setup</span><span class=\"o\">.</span>\n\n<span class=\"c1\">###
  Run command from the Dashboard console</span>\n\n<span class=\"n\">If</span> <span
  class=\"n\">you</span> <span class=\"n\">don</span><span class=\"s1\">&#39;t have
  heroku CLI set up and want to fix a few things on the pushed code to the heroku
  app, you can use the `Run Console` option from the `More` Tab on the top right corner
  of theApplication dashboard. This is a great way to run migrations, configure static
  files or tweak a few things here and there without messing up the code on GitHub
  or the remote git repositories. </span>\n\n<span class=\"err\">!</span><span class=\"p\">[</span><span
  class=\"n\">image</span><span class=\"o\">.</span><span class=\"n\">png</span><span
  class=\"p\">](</span><span class=\"n\">https</span><span class=\"p\">:</span><span
  class=\"o\">//</span><span class=\"n\">cdn</span><span class=\"o\">.</span><span
  class=\"n\">hashnode</span><span class=\"o\">.</span><span class=\"n\">com</span><span
  class=\"o\">/</span><span class=\"n\">res</span><span class=\"o\">/</span><span
  class=\"n\">hashnode</span><span class=\"o\">/</span><span class=\"n\">image</span><span
  class=\"o\">/</span><span class=\"n\">upload</span><span class=\"o\">/</span><span
  class=\"n\">v1652614775294</span><span class=\"o\">/</span><span class=\"n\">lgDPwr2yr</span><span
  class=\"o\">.</span><span class=\"n\">png</span><span class=\"p\">)</span>\n\n\n<span
  class=\"err\">!</span><span class=\"p\">[</span><span class=\"n\">image</span><span
  class=\"o\">.</span><span class=\"n\">png</span><span class=\"p\">](</span><span
  class=\"n\">https</span><span class=\"p\">:</span><span class=\"o\">//</span><span
  class=\"n\">cdn</span><span class=\"o\">.</span><span class=\"n\">hashnode</span><span
  class=\"o\">.</span><span class=\"n\">com</span><span class=\"o\">/</span><span
  class=\"n\">res</span><span class=\"o\">/</span><span class=\"n\">hashnode</span><span
  class=\"o\">/</span><span class=\"n\">image</span><span class=\"o\">/</span><span
  class=\"n\">upload</span><span class=\"o\">/</span><span class=\"n\">v1652614821950</span><span
  class=\"o\">/</span><span class=\"n\">uTzQVB8sC</span><span class=\"o\">.</span><span
  class=\"n\">png</span><span class=\"p\">)</span>\n\n\n<span class=\"err\">!</span><span
  class=\"p\">[</span><span class=\"n\">image</span><span class=\"o\">.</span><span
  class=\"n\">png</span><span class=\"p\">](</span><span class=\"n\">https</span><span
  class=\"p\">:</span><span class=\"o\">//</span><span class=\"n\">cdn</span><span
  class=\"o\">.</span><span class=\"n\">hashnode</span><span class=\"o\">.</span><span
  class=\"n\">com</span><span class=\"o\">/</span><span class=\"n\">res</span><span
  class=\"o\">/</span><span class=\"n\">hashnode</span><span class=\"o\">/</span><span
  class=\"n\">image</span><span class=\"o\">/</span><span class=\"n\">upload</span><span
  class=\"o\">/</span><span class=\"n\">v1652614845269</span><span class=\"o\">/</span><span
  class=\"n\">BkZhu3SGH</span><span class=\"o\">.</span><span class=\"n\">png</span><span
  class=\"p\">)</span>\n\n<span class=\"c1\">### Deploy with Docker </span>\n\n<span
  class=\"n\">You</span> <span class=\"n\">can</span> <span class=\"n\">even</span>
  <span class=\"n\">use</span> <span class=\"n\">the</span> <span class=\"n\">docker</span>
  <span class=\"n\">container</span> <span class=\"n\">to</span> <span class=\"n\">deploy</span>
  <span class=\"n\">a</span> <span class=\"n\">Django</span> <span class=\"n\">application</span>
  <span class=\"n\">on</span> <span class=\"n\">Heroku</span><span class=\"o\">.</span>
  <span class=\"n\">It</span> <span class=\"ow\">is</span> <span class=\"n\">a</span>
  <span class=\"n\">great</span> <span class=\"n\">way</span> <span class=\"n\">of</span>
  <span class=\"n\">learning</span> <span class=\"n\">a</span> <span class=\"n\">lot</span>
  <span class=\"n\">of</span> <span class=\"n\">deployment</span> <span class=\"n\">strategies</span>
  <span class=\"ow\">and</span> <span class=\"n\">techniques</span> <span class=\"n\">using</span>
  <span class=\"n\">Docker</span><span class=\"o\">.</span> <span class=\"n\">You</span><span
  class=\"s1\">&#39;ll get familiar with interesting concepts like virtualization,
  and containerization, and also learn Docker on the way. You can follow this tutorial
  on [Deploying Django applications with Docker on Heroku](https://testdriven.io/blog/deploying-django-to-heroku-with-docker/).
  Also, you can check out the official Heroku documentation for [deploying python
  applications](https://devcenter.heroku.com/articles/deploying-python).</span>\n\n<span
  class=\"c1\">### What are Dynos?</span>\n\n<span class=\"n\">Dynos</span> <span
  class=\"n\">are</span> <span class=\"n\">simply</span> <span class=\"n\">web</span>
  <span class=\"n\">processes</span> <span class=\"ow\">or</span> <span class=\"n\">workers</span>
  <span class=\"n\">that</span> <span class=\"n\">serve</span> <span class=\"n\">your</span>
  <span class=\"n\">web</span> <span class=\"n\">application</span><span class=\"o\">.</span>
  <span class=\"n\">Dynos</span> <span class=\"ow\">in</span> <span class=\"n\">Heroku</span>
  <span class=\"n\">are</span> <span class=\"n\">allocated</span> <span class=\"n\">based</span>
  <span class=\"n\">on</span> <span class=\"n\">the</span> <span class=\"n\">build</span>
  <span class=\"n\">process</span><span class=\"p\">,</span> <span class=\"n\">once</span>
  <span class=\"n\">the</span> <span class=\"n\">slug</span> <span class=\"ow\">is</span>
  <span class=\"n\">created</span> <span class=\"n\">a</span> <span class=\"n\">dyno</span>
  <span class=\"ow\">is</span> <span class=\"n\">created</span> <span class=\"k\">as</span>
  <span class=\"n\">it</span> <span class=\"n\">runs</span> <span class=\"n\">on</span>
  <span class=\"n\">a</span> <span class=\"n\">VM</span> <span class=\"n\">container</span><span
  class=\"o\">.</span> <span class=\"n\">This</span> <span class=\"n\">simply</span>
  <span class=\"n\">means</span> <span class=\"n\">there</span> <span class=\"n\">are</span>
  <span class=\"n\">limitations</span> <span class=\"n\">on</span> <span class=\"n\">how</span>
  <span class=\"n\">to</span> <span class=\"n\">use</span> <span class=\"n\">the</span>
  <span class=\"n\">web</span> <span class=\"n\">application</span> <span class=\"ow\">and</span>
  <span class=\"n\">its</span> <span class=\"n\">sleep</span> <span class=\"n\">process</span><span
  class=\"o\">.</span> <span class=\"n\">The</span> <span class=\"n\">hobby</span>
  <span class=\"n\">tier</span> <span class=\"ow\">is</span> <span class=\"n\">sufficient</span>
  <span class=\"k\">for</span> <span class=\"n\">normal</span> <span class=\"n\">testing</span>
  <span class=\"n\">projects</span> <span class=\"ow\">and</span> <span class=\"n\">side</span>
  <span class=\"n\">projects</span> <span class=\"n\">though</span> <span class=\"n\">you</span>
  <span class=\"n\">will</span> <span class=\"n\">have</span> <span class=\"n\">to</span>
  <span class=\"n\">pay</span> <span class=\"ow\">and</span> <span class=\"n\">move</span>
  <span class=\"n\">into</span> <span class=\"n\">advance</span> <span class=\"n\">tiers</span>
  <span class=\"n\">to</span> <span class=\"n\">increase</span> <span class=\"n\">the</span>
  <span class=\"n\">dyno</span> <span class=\"n\">allocations</span> <span class=\"ow\">and</span>
  <span class=\"n\">scaling</span> <span class=\"n\">of</span> <span class=\"n\">those</span>
  <span class=\"n\">web</span> <span class=\"n\">processes</span><span class=\"o\">.</span>
  \n\n<span class=\"n\">It</span><span class=\"s1\">&#39;s not a simple thing to understand
  but to keep it simple, it might be a container to process the client&#39;</span><span
  class=\"n\">s</span> <span class=\"n\">request</span> <span class=\"ow\">and</span>
  <span class=\"n\">serve</span> <span class=\"n\">the</span> <span class=\"n\">page</span>
  <span class=\"k\">for</span> <span class=\"n\">a</span> <span class=\"n\">finite</span>
  <span class=\"n\">duration</span> <span class=\"n\">of</span> <span class=\"n\">the</span>
  <span class=\"n\">interaction</span><span class=\"o\">.</span> <span class=\"n\">Also</span><span
  class=\"p\">,</span> <span class=\"n\">your</span> <span class=\"n\">application</span>
  <span class=\"n\">will</span> <span class=\"n\">sleep</span> <span class=\"n\">after</span>
  <span class=\"n\">half</span> <span class=\"n\">an</span> <span class=\"n\">hour</span><span
  class=\"p\">,</span> <span class=\"k\">if</span> <span class=\"n\">you</span> <span
  class=\"k\">try</span> <span class=\"n\">to</span> <span class=\"n\">reload</span>
  <span class=\"n\">the</span> <span class=\"n\">application</span> <span class=\"n\">every</span>
  <span class=\"n\">half</span> <span class=\"n\">an</span> <span class=\"n\">hour</span>
  <span class=\"n\">it</span> <span class=\"n\">will</span> <span class=\"n\">consume</span>
  <span class=\"n\">your</span> <span class=\"n\">resource</span> <span class=\"n\">allocation</span>
  <span class=\"k\">for</span> <span class=\"n\">the</span> <span class=\"n\">month</span>
  <span class=\"ow\">and</span> <span class=\"n\">this</span> <span class=\"ow\">is</span>
  <span class=\"n\">how</span> <span class=\"n\">the</span> <span class=\"n\">tiers</span>
  <span class=\"ow\">and</span> <span class=\"n\">divided</span> <span class=\"k\">for</span>
  <span class=\"n\">paid</span> <span class=\"n\">services</span> <span class=\"n\">on</span>
  <span class=\"n\">Heroku</span><span class=\"o\">.</span> <span class=\"n\">You</span>
  <span class=\"n\">can</span> <span class=\"n\">check</span> <span class=\"n\">out</span>
  <span class=\"n\">the</span> <span class=\"n\">detail</span> <span class=\"n\">over</span>
  <span class=\"p\">[</span><span class=\"n\">here</span><span class=\"p\">](</span><span
  class=\"n\">https</span><span class=\"p\">:</span><span class=\"o\">//</span><span
  class=\"n\">www</span><span class=\"o\">.</span><span class=\"n\">heroku</span><span
  class=\"o\">.</span><span class=\"n\">com</span><span class=\"o\">/</span><span
  class=\"n\">pricing</span><span class=\"c1\">#containers).</span>\n\n<span class=\"c1\">##
  Conclusion</span>\n\n<span class=\"n\">So</span><span class=\"p\">,</span> <span
  class=\"n\">that</span> <span class=\"ow\">is</span> <span class=\"n\">one</span>
  <span class=\"n\">of</span> <span class=\"n\">the</span> <span class=\"n\">ways</span>
  <span class=\"n\">we</span> <span class=\"n\">can</span> <span class=\"n\">deploy</span>
  <span class=\"n\">a</span> <span class=\"n\">Django</span> <span class=\"n\">application</span>
  <span class=\"n\">on</span> <span class=\"n\">Heroku</span> <span class=\"k\">with</span>
  <span class=\"n\">the</span> <span class=\"n\">PostgreSQL</span> <span class=\"n\">database</span><span
  class=\"o\">.</span> <span class=\"n\">You</span> <span class=\"n\">can</span> <span
  class=\"n\">find</span> <span class=\"n\">the</span> <span class=\"p\">[</span><span
  class=\"n\">django</span><span class=\"o\">-</span><span class=\"n\">blog</span>
  <span class=\"n\">project</span><span class=\"p\">]</span> <span class=\"n\">on</span>
  <span class=\"p\">[</span><span class=\"n\">GitHub</span><span class=\"p\">]</span>
  <span class=\"k\">for</span> <span class=\"n\">following</span> <span class=\"n\">along</span>
  <span class=\"k\">with</span> <span class=\"n\">the</span> <span class=\"n\">deployment</span>
  <span class=\"n\">process</span><span class=\"o\">.</span>  <span class=\"n\">In</span>
  <span class=\"n\">the</span> <span class=\"nb\">next</span> <span class=\"n\">few</span>
  <span class=\"n\">parts</span> <span class=\"n\">of</span> <span class=\"n\">the</span>
  <span class=\"n\">series</span><span class=\"p\">,</span> <span class=\"n\">we</span>
  <span class=\"n\">will</span> <span class=\"n\">be</span> <span class=\"n\">hopefully</span>
  <span class=\"n\">covering</span> <span class=\"n\">other</span> <span class=\"n\">platforms</span>
  <span class=\"n\">where</span> <span class=\"n\">you</span> <span class=\"n\">can</span>
  <span class=\"n\">deploy</span> <span class=\"n\">a</span> <span class=\"n\">Django</span>
  <span class=\"n\">o</span> <span class=\"n\">application</span><span class=\"o\">.</span>\n\n<span
  class=\"n\">Hopefully</span><span class=\"p\">,</span> <span class=\"n\">you</span>
  <span class=\"n\">liked</span> <span class=\"n\">the</span> <span class=\"n\">above</span>
  <span class=\"n\">tutorial</span><span class=\"p\">,</span> <span class=\"k\">if</span>
  <span class=\"n\">you</span> <span class=\"n\">have</span> <span class=\"nb\">any</span>
  <span class=\"n\">questions</span><span class=\"o\">.</span> <span class=\"n\">feedback</span><span
  class=\"p\">,</span> <span class=\"ow\">or</span> <span class=\"n\">queries</span><span
  class=\"p\">,</span> <span class=\"n\">you</span> <span class=\"n\">can</span> <span
  class=\"n\">contact</span> <span class=\"n\">me</span> <span class=\"n\">on</span>
  <span class=\"n\">the</span> <span class=\"n\">Social</span> <span class=\"n\">handles</span>
  <span class=\"n\">provided</span> <span class=\"n\">below</span><span class=\"o\">.</span>
  <span class=\"n\">Thank</span> <span class=\"n\">you</span> <span class=\"k\">for</span>
  <span class=\"n\">reading</span> <span class=\"ow\">and</span> <span class=\"n\">till</span>
  <span class=\"n\">the</span> <span class=\"nb\">next</span> <span class=\"n\">post</span>
  <span class=\"n\">Happy</span> <span class=\"n\">Coding</span> <span class=\"p\">:)</span>\n</pre></div>\n\n</pre>\n\n"
cover: ''
date: 2022-05-15
datetime: 2022-05-15 00:00:00+00:00
description: Django projects are quite easy to build and simple to understand, you
  might have created a Django application and wanted to show it to the world? You
  can deploy
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2022-05-15-Django-Deploy-Heroku.md
html: "<h2>Introduction</h2>\n<p>Django projects are quite easy to build and simple
  to understand, you might have created a Django application and wanted to show it
  to the world? You can deploy a basic Django application with a database(PostgreSQL)
  with Heroku. It provides a decent free tier with some great features and add-ons.
  A free tier Heroku account has a limitation of 5 apps, limited data in the database,
  limited connections to the server per month, and so on.</p>\n<p>Though the free
  tier is not a great option for bigger applications, it suits really well for smaller
  scale and ide projects, so we will be looking into the details of how to deploy
  a Django application to <a href=\"https://heroku.com/\">Heroku</a> which is a Platform
  as Service (PaS).</p>\n<p>This series will be an extension of the series <a href=\"https://techstructiveblog.hashnode.dev/series/django-basics\">Django
  basics</a> which covered the basics of the Django framework, we covered from basic
  Django fundamentals to building a CRUD API. In this series, we will be exploring
  some platforms for giving a better understanding of how things work and understanding
  a few components that were left as default while understanding the basics of Django.
  Let's get started with <a href=\"https://techstructiveblog.hashnode.dev/series/django-deployment\">Django
  Deployment</a>!</p>\n<h2>Creating a Django Application</h2>\n<p>For deploying an
  app, we definitely need an app, we need to create a basic Django application to
  deploy on the web. We'll be creating a simple blog application with a couple of
  views and a simple model structure. As for the database, we'll be using Postgres
  as Heroku has an add-on for it and it is pretty easy to configure.</p>\n<h3>Set
  up a virtual environment</h3>\n<p>We need to set up a virtual environment in order
  to keep the Django project neat and tidy by managing the project-specific dependencies
  and packages. We can use the <a href=\"https://pypi.org/project/virtualenv/\">virtualenv</a>
  package to isolate a python project from the rest of the system.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span># install the virtualenv
  package\npip install virtualenv\n\n# create a virtual env for the project\nvirtualenv
  .venv\n\n# activate the virtualenv\nWindows:\n.venv\\Scripts\\activate\n\nLinux/macOS:\nsource
  .venv/bin/activate\n</pre></div>\n\n</pre>\n\n<p>This will set up the project nicely
  for a Django project, you now install the core Django package and get started with
  creating a Django application.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># install django</span>\npip install django\n\n<span class=\"c1\">#
  start a django project</span>\ndjango-admin startproject blog .\n\n<span class=\"nb\">cd</span>
  blog\n\n<span class=\"c1\"># create a application in django project</span>\npython
  manage.py createapp api\n\n<span class=\"c1\"># Create some models, views, URLs,
  templates</span>\n\n<span class=\"c1\"># run the server</span>\npython manag.py
  runserver\n</pre></div>\n\n</pre>\n\n<p>We assume you already have a Django project
  configured with some basic URLs, views and templates or static files as per your
  project and requirements, for this tutorial I will be using the simple blog application
  from my previous Django tutorials as a reference. You can follow along with my <a
  href=\"https://techstructiveblog.hashnode.dev/series/django-basics\">Django Basics</a>
  series and refer to the Blog Application project on <a href=\"https://github.com/Mr-Destructive/django-blog\">GitHub</a>.</p>\n<h2>Configuring
  the Django Application</h2>\n<p>Make sure to create and activate the virtual environment
  for this django project. This should be done to manage the dependencies and packages
  used in the project. If you are not aware of the virtual environment and django
  setup, you can follow up with this <a href=\"https://mr-destructive.github.io/techstructive-blog/django-setup-script/\">post</a>.</p>\n<h3>Creating
  a runtime.txt file</h3>\n<p>Now, Firstly we need to specify which type and version
  of language we are using. Since Django is a Python-based web framework, we need
  to select the python version in a text file.</p>\n<p><strong>runtime.txt</strong></p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>python-3.9.5\n</pre></div>\n\n</pre>\n\n<p>Here,
  the version can be anything as per your project and the packages installed.</p>\n<h3>Creating
  requirements.txt file</h3>\n<p>We'll first create a <code>requirements.txt</code>
  file for storing all the dependencies and packages installed in the application.
  This will help in installing dependencies while deploying the application. We can
  either use a <code>requirements.txt</code> file using <code>virtualenv</code> or
  a <code>Pipfile</code> using Pipenv. Both serve the same purpose but a bit differently.</p>\n<p>Assuming
  you are in an isolated virtual environment for this Django project, you can create
  a requirements.txt file using the below command:</p>\n<p>Make sure the virtualenv
  is activated before running the command:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>pip freeze &gt; requirements.txt\n</pre></div>\n\n</pre>\n\n<p>This
  will create a simple text file that contains the package names along with the versions
  used in the current virtual environment. A simple Django requirements file would
  look something like this:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>asgiref==3.4.1\nDjango==3.2.11\npytz==2021.3\nsqlparse==0.4.2\ntyping_extensions==4.0.1\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652614060461/kPTZ9R8Xp.png\"
  alt=\"image.png\" /></p>\n<p>This is vanilla Django without any additional dependencies
  but if you have integrated other functionalities like Django Rest Framework, PostgreSQL,
  Crispy Forms, Schedulers, etc. there might be additional dependencies that become
  quite crucial for the smooth working of the project.</p>\n<p>If you are using pipenv,
  you don't need to make any efforts to manually activate and manage virtual environment,
  the dependencies are installed and taken care of by the pipenv installer. You just
  need to make sure to install any package with <code>pipenv install</code> and not
  <code>pip install</code> for better and improved package tracking.</p>\n<p>So, that's
  all we need to take care of packages and keep a list of these integrated packages
  for the project. You need to update the requirements.txt file every time you install
  any new package or modify the dependencies for a project. Simply use the command
  <code>pip freeze &gt; requirements.txt</code> in the activated virtual environment.</p>\n<h3>Creating
  a Procfile</h3>\n<p>Next up, we have the <code>Procfile</code>, a procfile is a
  special file that holds information about the processes to be run to start or activate
  the project. In our case, for django we need a web process that can manage the server.</p>\n<p>A
  Procfile is a simple file without any extension, make sure to write <code>Procfile</code>
  as it is as the name of the file in the root folder of the project. Inside the file
  add the following contents:</p>\n<p><strong>Procfile</strong></p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>Procfile</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nl\">web</span><span class=\"p\">:</span><span class=\"w\"> </span>gunicorn<span
  class=\"w\"> </span><span class=\"err\">&lt;</span>project_name<span class=\"err\">&gt;</span>.wsgi<span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>As we can see we have defined
  the <code>web</code> process using <code>gunicorn</code>, <a href=\"https://pypi.org/project/gunicorn/\">Gunicorn</a>
  is a python package that helps in creating WSGI HTTP Server for the UNIX operating
  systems. So, we need to install the package and update the package dependency list.</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>pip install gunicorn\n\npip
  freeze &gt; requirements.txt\n</pre></div>\n\n</pre>\n\n<p>That would be good to
  go for creating and serving up the project while deploying the project on Heroku.</p>\n<h2>Creating
  a Heroku App</h2>\n<p>A Heroku App is basically like your Django Project, you can
  create apps for deploying your django projects on Heroku. You are limited to 5 apps
  on the Free tier but that can be expanded on the paid plans.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652456732519/cyOQZ3UZK.png\"
  alt=\"image.png\" /></p>\n<p>The name of your heroku app should be unique globally,
  you need to try a few combinations before a good one fits your project. This name
  has no significance on your django project code, though it would be the name from
  which you would access the web application as a name <code>&lt;app-name&gt;.herokuapp.com</code>.</p>\n<p>So,
  choose it wisely if you are not attaching a custom domain. You can attach a custom
  domain, you can navigate to the <code>domain</code> section in the settings tab.</p>\n<h2>Setting
  up the database</h2>\n<p>To set up and configure a database in django on Heroku,
  we need to manually acquire and attach a PostgreSQL add-on to the heroku app.</p>\n<ul>\n<li>Firstly
  locate to the Resources Tab in your Heroku app.</li>\n<li>Search <code>postgres</code>
  in the Add-ons Search bar</li>\n<li>Click on the <code>Heroku Postgres</code> Add-on</li>\n<li>Submit
  the Order Form and you have the add-on enabled in the app.</li>\n</ul>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652456842273/ijeWsVdOf.png\"
  alt=\"image.png\" /></p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652456877447/dLG30ac_m.png\"
  alt=\"image.png\" /></p>\n<p>Alternately, you can use the Heroku CLI,</p>\n<h3>Heroku
  CLI Setup</h3>\n<p>You can use the Heroku CLI which is a command-line interface
  for managing and creating Heroku applications.  You need to first install the CLI
  by heading towards the <a href=\"https://devcenter.heroku.com/articles/heroku-command-line\">heroku
  documentation</a>. Once the CLI is installed, you need to login into your Heroku
  account by entering the following command:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>heroku login\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652605604920/HnTr2KbTi.png\"
  alt=\"image.png\" /></p>\n<p>This will allow us to work with heroku commands and
  manage our heroku application from the command line itself. The below command will
  create a add-on for <code>heroku-postgres</code> for the application provided as
  the parameter options</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>heroku addons:create heroku-postgresql:hobby-dev
  --app &lt;app_name&gt;\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652507166978/i1IJ5EGjJ.png\"
  alt=\"image.png\" /></p>\n<p>This should hopefully add a fresh instance of a postgres
  database for your heroku app. You can now configure the database for your project,
  we simply need the Database URL from the heroku app dashboard. We'll see how to
  configure the environment variables in Django for Heroku to keep your secrets like
  the <code>SECRET_KEY</code>, <code>DATABSE_URL</code>, etc.</p>\n<p>If you want
  MySQL as a database, you can check out the <a href=\"https://devcenter.heroku.com/articles/cleardb\">ClearDB</a>
  Add-On for Heroku to simply attach it like Postgres on your Heroku application.
  Also, if you wish to add <a href=\"https://www.mongodb.com/compatibility/mongodb-and-django\">MongoDB</a>
  into your Django application on Heroku, you can use <a href=\"https://devcenter.heroku.com/articles/ormongo\">Object
  Rocket</a>(OR Mongo). It is not available in the free tier though, unlike PostgreSQL
  and MySQL.</p>\n<h2>Configuring Environment Variables</h2>\n<p>We need to keep our
  secrets for the django project out of the deployed code and in a safe place, we
  can create environment variables and keep them in a <code>.env</code> file which
  will be git-ignored. We do not want this <code>.env</code> file to be open source
  and thus should not be committed.</p>\n<p>We first need to create a new secret key
  if you already have a GitHub repository, chances are you would have committed the
  default secret key open for the world to see, it is an insecure way of deploying
  django apps in production.</p>\n<p>So, open up a terminal and launch a python REPL:</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>python\n\nimport secrets\nsecrets.token_hex(24)\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652512239319/5AACaTGOD.png\"
  alt=\"image.png\" /></p>\n<p>This should generate a secret key that only you know
  now. So, just copy the key without the quotes and create a file <code>.env</code>
  in the root project folder.</p>\n<p><strong>.env</strong></p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>SECRET_KEY=76419fd6885a677f802fd1d2b5acd0188e23e001042b05a8\n</pre></div>\n\n</pre>\n\n<p>The
  <code>.env</code> file should also be added to the <code>.gitignore</code> file,
  so simply append the following in the <code>.gitignore</code> file</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>.env\n</pre></div>\n\n</pre>\n\n<p>This
  file is only created to test the project locally, so we need to also make this key
  available on heroku. For doing that we need to add Config Variables to the heroku
  app.</p>\n<ul>\n<li>Locate to the Settings Tab in your Heroku Application Dashboard</li>\n<li>We
  have the <code>Config Vars</code> section in the located tab\n= We need to reveal
  those variables and we will be able to see all the variables.</li>\n</ul>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652456988713/5VM6E29_o.png\"
  alt=\"image.png\" /></p>\n<p>We already have a <code>DATABASE_URL</code> variable
  declared when we attached the <code>django-postgres</code> database to our application.
  We will now add one more variable to the Application, i.e. the <code>SECRET_KEY</code>.
  Simply, enter the name of the key and also enter the value of the key, so basically
  a key-value pair. Simply click on the <code>Add</code> button and this will add
  the variable to your application.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652515244870/LRyPzJr41.png\"
  alt=\"image.png\" /></p>\n<p>You also need to copy the <code>DATABSE_URL</code>
  into our local setup file(<code>.env</code> file). Copy the Database URL and paste
  it into the <code>.env</code> file as follows:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>env</div>\n<div class=\"highlight\"><pre><span></span>DATABASE_URL=postgres://sjxgipufegmgsw:78cbb568e@ec2-52-4-104-184.compute-1.amazonaws.com:5432/dbmuget\n</pre></div>\n\n</pre>\n\n<p>The
  format for the postgres URL is as follows:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>postgresql://[user[:password]@][netloc][:port][/dbname]\n</pre></div>\n\n</pre>\n\n<p>We
  have now created environment variables for our django application and also added
  config vars in the heroku app, we now need a way to parse these environment variables
  into the Django project.</p>\n<h3>Parsing Environment variables using python-dotenv</h3>\n<p>We
  will use <a href=\"https://pypi.org/project/python-dotenv/\">python-dotenv</a> to
  parse variables into the django settings configurations like <code>SECRET_KEY</code>
  and <code>DATABASES</code>.</p>\n<ul>\n<li>Install <code>python-dotenv</code> with
  pip with the command :</li>\n</ul>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>pip install python-dotenv\n</pre></div>\n\n</pre>\n\n<p>We
  need to then modify the default variables in the <code>settings.py</code> file.
  Firstly, we will load in the <code>.env</code> file for accessing the environment
  variables for the configuration of the project.</p>\n<p>Append the following code,
  to the top of the <code>settings.py</code> file, make sure don't override the configuration
  so remove unnecessary imports and configurations.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># &lt;project_name&gt;/settings.py</span>\n\n<span class=\"kn\">import</span>
  <span class=\"nn\">os</span>\n<span class=\"kn\">from</span> <span class=\"nn\">dotenv</span>
  <span class=\"kn\">import</span> <span class=\"n\">load_dotenv</span>\n\n<span class=\"n\">BASE_DIR</span>
  <span class=\"o\">=</span> <span class=\"n\">Path</span><span class=\"p\">(</span><span
  class=\"vm\">__file__</span><span class=\"p\">)</span><span class=\"o\">.</span><span
  class=\"n\">resolve</span><span class=\"p\">()</span><span class=\"o\">.</span><span
  class=\"n\">parent</span><span class=\"o\">.</span><span class=\"n\">parent</span>\n\n<span
  class=\"n\">load_dotenv</span><span class=\"p\">(</span><span class=\"n\">os</span><span
  class=\"o\">.</span><span class=\"n\">path</span><span class=\"o\">.</span><span
  class=\"n\">join</span><span class=\"p\">(</span><span class=\"n\">BASE_DIR</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;.env&quot;</span><span class=\"p\">))</span>\n</pre></div>\n\n</pre>\n\n<p>We
  have imported the package <code>dotenv</code> basically the <code>python-dotenv</code>
  into the <code>settings.py</code> file and the module <code>os</code> for loading
  the <code>.env</code> file. The <code>load_dotenv</code> function helps in loading
  the <code>key-value</code> pairs which are the configuration variables that can
  act as actual environment variables. We provide in a file to the <a href=\"https://saurabh-kumar.com/python-dotenv/\">load_dotenv</a>
  function which is the <code>.env</code> file in our case, you can pick up any location
  for the <code>.env</code> file but make sure to change the location while parsing
  the file into the <code>load_dotenv</code> function.</p>\n<p>After loading the variables
  into the <code>settings.py</code> file, we now need to access those variables and
  set the appropriate variables the configuration from the variables received from
  the <code>load_dotenv</code> function. The <code>os.getenv</code> function to access
  the environment variables. The <code>os.getenv</code> function takes a parameter
  as the <code>key</code> for the environment variable and returns the value of the
  environment variable.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"n\">SECRET_KEY</span> <span class=\"o\">=</span> <span class=\"n\">os</span><span
  class=\"o\">.</span><span class=\"n\">getenv</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;SECRET_KEY&quot;</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>We
  have secretly configured the <code>SECRET_KEY</code> for the django project. If
  you have any other variables as simple key-value pairs like <code>AUTH</code> passwords,
  username, etc. you can use this method to get the configuration variables.</p>\n<h3>Loading
  Database configuration</h3>\n<p>Databases are a bit different as compared to other
  simple configurations in django. We need to make a few adjustments to the default
  database configuration. We will install another package <code>dj-database-url</code>
  to configure <code>DATABASE_URL</code>. Since the databse_url has a few components
  we need a way to extract the details like the <code>hostname</code>, <code>port</code>,
  <code>database_name</code>, and <code>password</code>. Using the <code>dj-database-url</code>
  package we have a few functions that can serve the purpose.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>pip install dj-database-url\n</pre></div>\n\n</pre>\n\n<p>At
  the end of your <code>settings.py</code> file, append the following code.</p>\n<pre
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
  class=\"kn\">import</span> <span class=\"nn\">dj_database_url</span>\n\n<span class=\"n\">DATABASE_URL</span>
  <span class=\"o\">=</span> <span class=\"n\">os</span><span class=\"o\">.</span><span
  class=\"n\">getenv</span><span class=\"p\">(</span><span class=\"s2\">&quot;DATABASE_URL&quot;</span><span
  class=\"p\">)</span>\n\n<span class=\"n\">DATABASES</span> <span class=\"o\">=</span>
  <span class=\"p\">{</span>\n    <span class=\"s2\">&quot;default&quot;</span><span
  class=\"p\">:</span> <span class=\"n\">dj_database_url</span><span class=\"o\">.</span><span
  class=\"n\">config</span><span class=\"p\">(</span><span class=\"n\">default</span><span
  class=\"o\">=</span><span class=\"n\">DATABASE_URL</span><span class=\"p\">,</span>
  <span class=\"n\">conn_max_age</span><span class=\"o\">=</span><span class=\"mi\">1800</span><span
  class=\"p\">),</span>\n<span class=\"p\">}</span>\n</pre></div>\n\n</pre>\n\n<p>We
  would need an adapter for making migrations to the <code>PostgreSQL</code> database
  i.e. the <code>psycopg2</code> package. This is a mandatory step if you are working
  with <code>postgres</code> database. This can be installed with the simple pip install:</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>pip install psycopg2\n\n#
  If it does not work try\npip install psycopg2-binary\n\n\n# if still error persists
  try installing setuptools\npip install -U setuptools\npip install psycopg2\n</pre></div>\n\n</pre>\n\n<p>Now,
  that we have configured the database, we can now apply migrations to the fresh database
  of postgres provided by heroku. We will simply run the migrate command and in the
  heroku app the PostgreSQL database would have been modified and an appropriate schema
  should be applied.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>python manage.py migrate\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652602284553/oTtGev28-.png\"
  alt=\"image.png\" /></p>\n<p>Make sure to update the <code>requirements.txt</code>
  file before pushing the project to Heroku for making sure everything works as expected.
  Since we have installed a few additional packages that are directly used in the
  <code>settings.py</code> file, we need to run the <code>pip freeze</code> command
  to update the <code>requiremnets.txt</code> file.</p>\n<h2>Serving Static Files</h2>\n<p>Now,
  if you have some static files like <code>CSS</code>, <code>Javascript</code>, or
  <code>images</code>, you need to configure the staticfiles in order to serve them
  from the heroku server. We will require another config variable for collecting the
  static files from the selected repository.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"n\">STATIC_URL</span> <span class=\"o\">=</span> <span class=\"s2\">&quot;static/&quot;</span>\n<span
  class=\"n\">STATICFILES_DIRS</span> <span class=\"o\">=</span> <span class=\"p\">[</span><span
  class=\"n\">os</span><span class=\"o\">.</span><span class=\"n\">path</span><span
  class=\"o\">.</span><span class=\"n\">join</span><span class=\"p\">(</span><span
  class=\"n\">BASE_DIR</span><span class=\"p\">,</span> <span class=\"s2\">&quot;static&quot;</span><span
  class=\"p\">)]</span>\n<span class=\"n\">STATIC_ROOT</span> <span class=\"o\">=</span>
  <span class=\"n\">os</span><span class=\"o\">.</span><span class=\"n\">path</span><span
  class=\"o\">.</span><span class=\"n\">join</span><span class=\"p\">(</span><span
  class=\"n\">BASE_DIR</span><span class=\"p\">,</span> <span class=\"s2\">&quot;staticfiles&quot;</span><span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>Here, if you have served your
  static files from the <code>static</code> folder in the root directory of your django
  project, you can add the above code in the <a href=\"http://settings.py\">settings.py</a>
  file. We will collect all static files in the folder along with the default static
  files provided by django in the <code>staticfiles</code> directory. Run the following
  command if you want to test whether the static files are properly collected and
  served.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n\n<div class=\"highlight\"><pre><span></span>python manage.py collectstatic
  \n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652600828657/JgL4nLuiL.png\"
  alt=\"image.png\" /></p>\n<p>So, this command will collect all the static files
  and store them in a single place. We see that the files from the admin section are
  also copied as well as the custom static files from the project configuration. Now,
  we can move on to set the config variable for the heroku app.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>DISABLE_COLLECTSTATIC = 0\n</pre></div>\n\n</pre>\n\n<p>We
  can set the <code>DISABLE_COLLECTSTATIC</code> variable as <code>0</code> or <code>1</code>
  indicating whether to disable it or not. We have currently enabled the static file
  collection while deploying the app but you can set it to <code>1</code> to disable
  the collection of static files.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1652613798420/mbqzf1Kqd.png\"
  alt=\"image.png\" /></p>\n<p>Since I first tested the application on heroku, the
  static files don't work as expected, we need another package to make sure the staticfiles
  are served property. We will be installing the <code>whitenoise</code> package which
  serves as the middleware for serving the static files.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>pip install whitenoise\n</pre></div>\n\n</pre>\n\n<p>Add
  the whitenoise middleware <code>whitenoise.middleware.WhiteNoiseMiddleware</code>
  to the <code>MIDDLEWARE</code> list in the <code>settings.py</code> file.</p>\n<pre
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
  class=\"n\">MIDDLEWARE</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n<span
  class=\"o\">...</span>\n<span class=\"o\">...</span>\n<span class=\"o\">...</span>\n
  \   <span class=\"s1\">&#39;whitenoise.middleware.WhiteNoiseMiddleware&#39;</span><span
  class=\"p\">,</span>\n<span class=\"p\">]</span>\n\n<span class=\"err\">```</span>\n\n<span
  class=\"n\">That</span> <span class=\"n\">should</span> <span class=\"n\">be</span>
  <span class=\"n\">enough</span> <span class=\"n\">to</span> <span class=\"n\">make</span>
  <span class=\"n\">the</span> <span class=\"n\">most</span> <span class=\"n\">of</span>
  <span class=\"n\">the</span> <span class=\"n\">deployment</span> <span class=\"n\">on</span>
  <span class=\"n\">heroku</span><span class=\"o\">.</span> <span class=\"n\">You</span>
  <span class=\"n\">will</span> <span class=\"n\">have</span> <span class=\"n\">to</span>
  <span class=\"n\">make</span> <span class=\"n\">a</span> <span class=\"n\">few</span>
  <span class=\"n\">adjustments</span> <span class=\"k\">as</span> <span class=\"n\">per</span>
  <span class=\"n\">your</span> <span class=\"n\">requirements</span> <span class=\"ow\">and</span>
  <span class=\"n\">project</span><span class=\"o\">.</span>\n\n<span class=\"c1\">##
  Deploy from GitHub</span>\n\n<span class=\"n\">We</span> <span class=\"n\">are</span>
  <span class=\"n\">now</span> <span class=\"nb\">all</span> <span class=\"nb\">set</span>
  <span class=\"n\">to</span> <span class=\"n\">deploy</span> <span class=\"n\">the</span>
  <span class=\"n\">application</span> <span class=\"n\">on</span> <span class=\"n\">Heroku</span><span
  class=\"p\">,</span> <span class=\"n\">we</span> <span class=\"n\">can</span> <span
  class=\"n\">use</span> <span class=\"n\">the</span> <span class=\"err\">`</span><span
  class=\"n\">Connect</span> <span class=\"n\">to</span> <span class=\"n\">GitHub</span><span
  class=\"err\">`</span> <span class=\"ow\">or</span> <span class=\"err\">`</span><span
  class=\"n\">Heroku</span> <span class=\"n\">CLI</span><span class=\"err\">`</span>
  <span class=\"n\">to</span> <span class=\"n\">push</span> <span class=\"n\">the</span>
  <span class=\"n\">code</span> <span class=\"n\">into</span> <span class=\"n\">production</span><span
  class=\"o\">.</span> <span class=\"n\">Heroku</span> <span class=\"n\">CLI</span>
  <span class=\"ow\">is</span> <span class=\"n\">quite</span> <span class=\"n\">easy</span>
  <span class=\"k\">with</span> <span class=\"n\">a</span> <span class=\"n\">few</span>
  <span class=\"n\">sets</span> <span class=\"n\">of</span> <span class=\"n\">commands</span>
  <span class=\"n\">but</span> <span class=\"k\">if</span> <span class=\"n\">your</span>
  <span class=\"n\">project</span> <span class=\"ow\">is</span> <span class=\"n\">deployed</span>
  <span class=\"n\">on</span> <span class=\"n\">GitHub</span><span class=\"p\">,</span>
  <span class=\"n\">you</span> <span class=\"n\">can</span> <span class=\"n\">straightaway</span>
  <span class=\"n\">let</span> <span class=\"n\">the</span> <span class=\"n\">deployment</span>
  <span class=\"n\">start</span> <span class=\"n\">the</span> <span class=\"n\">build</span>
  <span class=\"n\">on</span> <span class=\"n\">a</span> <span class=\"n\">push</span>
  <span class=\"n\">to</span> <span class=\"n\">a</span> <span class=\"n\">specific</span>
  <span class=\"n\">branch</span><span class=\"o\">.</span> <span class=\"n\">This</span>
  <span class=\"n\">becomes</span> <span class=\"n\">quite</span> <span class=\"n\">automotive</span>
  <span class=\"ow\">and</span> <span class=\"n\">easy</span> <span class=\"n\">to</span>
  <span class=\"n\">scale</span> <span class=\"k\">while</span> <span class=\"n\">deploying</span>
  <span class=\"n\">a</span> <span class=\"n\">large</span><span class=\"o\">-</span><span
  class=\"n\">scale</span> <span class=\"n\">application</span><span class=\"o\">.</span>
  \n\n<span class=\"err\">```</span>\n<span class=\"n\">pip</span> <span class=\"n\">freeze</span>
  <span class=\"o\">&gt;</span> <span class=\"n\">requirements</span><span class=\"o\">.</span><span
  class=\"n\">txt</span>\n<span class=\"err\">```</span>\n\n<span class=\"n\">This</span>
  <span class=\"n\">step</span> <span class=\"ow\">is</span> <span class=\"n\">quite</span>
  <span class=\"n\">important</span> <span class=\"n\">because</span> <span class=\"n\">you</span>
  <span class=\"n\">need</span> <span class=\"n\">to</span> <span class=\"n\">make</span>
  <span class=\"n\">sure</span> <span class=\"n\">that</span> <span class=\"nb\">all</span>
  <span class=\"n\">the</span> <span class=\"n\">packages</span> <span class=\"n\">are</span>
  <span class=\"n\">listed</span> <span class=\"ow\">in</span> <span class=\"n\">the</span>
  <span class=\"err\">`</span><span class=\"n\">requirements</span><span class=\"o\">.</span><span
  class=\"n\">txt</span><span class=\"err\">`</span> <span class=\"n\">file</span>
  <span class=\"k\">else</span> <span class=\"n\">you</span> <span class=\"n\">will</span>
  <span class=\"n\">have</span> <span class=\"n\">to</span> <span class=\"n\">wait</span>
  <span class=\"k\">for</span> <span class=\"n\">the</span> <span class=\"n\">build</span>
  <span class=\"n\">to</span> <span class=\"n\">fail</span> <span class=\"ow\">and</span>
  <span class=\"n\">redeploy</span><span class=\"o\">.</span>\n\n<span class=\"n\">Make</span>
  <span class=\"n\">sure</span> <span class=\"n\">the</span> <span class=\"n\">server</span>
  <span class=\"ow\">is</span> <span class=\"n\">running</span> <span class=\"n\">first</span>
  <span class=\"n\">on</span> <span class=\"n\">your</span> <span class=\"n\">local</span>
  <span class=\"n\">machine</span><span class=\"p\">,</span> <span class=\"n\">remember</span>
  <span class=\"n\">the</span> <span class=\"n\">server</span> <span class=\"n\">will</span>
  <span class=\"n\">be</span> <span class=\"nb\">set</span> <span class=\"n\">up</span>
  <span class=\"kn\">from</span> <span class=\"nn\">scratch</span> <span class=\"n\">but</span>
  <span class=\"n\">the</span> <span class=\"n\">database</span> <span class=\"n\">will</span>
  <span class=\"n\">already</span> <span class=\"n\">have</span> <span class=\"n\">applied</span>
  <span class=\"n\">migrations</span> <span class=\"k\">if</span> <span class=\"n\">you</span>
  <span class=\"n\">have</span> <span class=\"n\">applied</span> <span class=\"n\">migrations</span>
  <span class=\"n\">before</span> <span class=\"n\">after</span> <span class=\"n\">connecting</span>
  <span class=\"n\">the</span> <span class=\"n\">Heroku</span> <span class=\"n\">Postgres</span>
  <span class=\"n\">database</span><span class=\"o\">.</span>\n \n<span class=\"err\">```</span>\n<span
  class=\"n\">python</span> <span class=\"n\">manage</span><span class=\"o\">.</span><span
  class=\"n\">py</span> <span class=\"n\">collectstatic</span>\n\n<span class=\"n\">python</span>
  <span class=\"n\">manage</span><span class=\"o\">.</span><span class=\"n\">py</span>
  <span class=\"n\">runserver</span>\n<span class=\"err\">```</span>\n\n<span class=\"n\">This</span>
  <span class=\"n\">will</span> <span class=\"nb\">set</span> <span class=\"n\">up</span>
  <span class=\"n\">the</span> <span class=\"n\">origin</span> <span class=\"n\">of</span>
  <span class=\"n\">the</span> <span class=\"n\">remote</span> <span class=\"n\">repository</span>
  <span class=\"n\">that</span> <span class=\"n\">will</span> <span class=\"n\">be</span>
  <span class=\"n\">pushing</span> <span class=\"n\">the</span> <span class=\"n\">project</span>
  <span class=\"n\">code</span><span class=\"o\">.</span> <span class=\"n\">Next</span><span
  class=\"p\">,</span> <span class=\"n\">make</span> <span class=\"n\">sure</span>
  <span class=\"n\">to</span> <span class=\"n\">commit</span> <span class=\"n\">the</span>
  <span class=\"n\">code</span> <span class=\"n\">which</span> <span class=\"n\">will</span>
  <span class=\"n\">contain</span> <span class=\"nb\">all</span> <span class=\"n\">the</span>
  <span class=\"n\">required</span> <span class=\"n\">stuff</span> <span class=\"k\">for</span>
  <span class=\"n\">deploying</span> <span class=\"n\">the</span> <span class=\"n\">code</span><span
  class=\"o\">.</span>\n\n<span class=\"n\">Checklist</span> <span class=\"k\">for</span>
  <span class=\"n\">deploying</span> <span class=\"n\">the</span> <span class=\"n\">code</span>\n\n<span
  class=\"err\">```</span>\n<span class=\"o\">-</span> <span class=\"n\">requirements</span><span
  class=\"o\">.</span><span class=\"n\">txt</span>\n<span class=\"o\">-</span> <span
  class=\"n\">Procfile</span>\n<span class=\"o\">-</span> <span class=\"n\">runtime</span><span
  class=\"o\">.</span><span class=\"n\">txt</span>\n<span class=\"o\">-</span> <span
  class=\"n\">django</span><span class=\"o\">-</span><span class=\"n\">project</span>\n<span
  class=\"o\">-</span> <span class=\"n\">environment</span> <span class=\"n\">variables</span>
  <span class=\"o\">/</span> <span class=\"n\">config</span> <span class=\"n\">variables</span>
  \n<span class=\"o\">-</span> <span class=\"n\">static</span> <span class=\"n\">file</span>
  <span class=\"n\">configuration</span>\n<span class=\"o\">-</span> <span class=\"n\">database</span>
  <span class=\"n\">configuration</span>\n<span class=\"o\">-</span> <span class=\"n\">migrate</span>
  <span class=\"n\">schema</span> <span class=\"n\">of</span> <span class=\"n\">database</span>
  \n<span class=\"o\">-</span> <span class=\"n\">gitignore</span> <span class=\"n\">file</span>
  <span class=\"k\">for</span> <span class=\"n\">ignoring</span> <span class=\"n\">virtualenvs</span><span
  class=\"p\">,</span> <span class=\"o\">.</span><span class=\"n\">env</span> <span
  class=\"n\">file</span><span class=\"p\">,</span> <span class=\"n\">staticfiles</span><span
  class=\"p\">,</span> <span class=\"n\">etc</span>\n<span class=\"err\">```</span>\n\n<span
  class=\"n\">here</span><span class=\"s1\">&#39;s a sample `.gitignore` for a minimal
  django project.</span>\n\n<span class=\"err\">```</span><span class=\"n\">gitignore</span>\n<span
  class=\"o\">.</span><span class=\"n\">env</span><span class=\"o\">/</span>\n<span
  class=\"o\">.</span><span class=\"n\">venv</span><span class=\"o\">/</span>\n<span
  class=\"n\">env</span><span class=\"o\">/</span>\n<span class=\"n\">venv</span><span
  class=\"o\">/</span>\n<span class=\"o\">*.</span><span class=\"n\">env</span>\n\n<span
  class=\"o\">*.</span><span class=\"n\">pyc</span>\n<span class=\"n\">db</span><span
  class=\"o\">.</span><span class=\"n\">sqlite3</span>\n<span class=\"n\">staticfiles</span><span
  class=\"o\">/</span>\n<span class=\"err\">```</span>\n\n<span class=\"n\">If</span>
  <span class=\"n\">you</span> <span class=\"n\">want</span> <span class=\"n\">a</span>
  <span class=\"n\">full</span><span class=\"o\">-</span><span class=\"n\">fledged</span>
  <span class=\"err\">`</span><span class=\"o\">.</span><span class=\"n\">gitignore</span><span
  class=\"err\">`</span> <span class=\"k\">for</span> <span class=\"n\">a</span> <span
  class=\"nb\">complex</span> <span class=\"n\">django</span> <span class=\"n\">project</span><span
  class=\"p\">,</span> <span class=\"n\">you</span> <span class=\"n\">can</span> <span
  class=\"n\">take</span> <span class=\"n\">the</span> <span class=\"n\">reference</span>
  <span class=\"kn\">from</span> <span class=\"nn\">Jose</span> <span class=\"n\">Padilla</span><span
  class=\"s1\">&#39;s [gitignore Template](https://github.com/jpadilla/django-project-template/blob/master/.gitignore)
  for a django project.  </span>\n\n<span class=\"c1\">### Git Commit the Django Project</span>\n<span
  class=\"err\">```</span>\n<span class=\"n\">git</span> <span class=\"n\">status</span>
  \n\n<span class=\"n\">git</span> <span class=\"n\">add</span> <span class=\"o\">.</span>\n\n<span
  class=\"n\">git</span> <span class=\"n\">commit</span> <span class=\"o\">-</span><span
  class=\"n\">m</span> <span class=\"s2\">&quot;config for heroku deployment&quot;</span>\n<span
  class=\"err\">```</span>\n<span class=\"n\">Carefully</span> <span class=\"n\">check</span>
  <span class=\"n\">the</span> <span class=\"n\">status</span> <span class=\"n\">of</span>
  <span class=\"n\">the</span> <span class=\"n\">git</span> <span class=\"n\">repository</span>
  <span class=\"n\">before</span> <span class=\"n\">committing</span> <span class=\"ow\">and</span>
  <span class=\"n\">make</span> <span class=\"n\">sure</span> <span class=\"n\">you</span>
  <span class=\"n\">don</span><span class=\"s1\">&#39;t forget anything by mistake,
  it won&#39;</span><span class=\"n\">t</span> <span class=\"n\">a</span> <span class=\"n\">big</span>
  <span class=\"n\">problem</span> <span class=\"n\">but</span> <span class=\"n\">it</span>
  <span class=\"n\">would</span> <span class=\"n\">mess</span> <span class=\"n\">up</span>
  <span class=\"n\">the</span> <span class=\"n\">build</span> <span class=\"n\">process</span><span
  class=\"o\">.</span>\n\n\n<span class=\"err\">!</span><span class=\"p\">[</span><span
  class=\"n\">image</span><span class=\"o\">.</span><span class=\"n\">png</span><span
  class=\"p\">](</span><span class=\"n\">https</span><span class=\"p\">:</span><span
  class=\"o\">//</span><span class=\"n\">cdn</span><span class=\"o\">.</span><span
  class=\"n\">hashnode</span><span class=\"o\">.</span><span class=\"n\">com</span><span
  class=\"o\">/</span><span class=\"n\">res</span><span class=\"o\">/</span><span
  class=\"n\">hashnode</span><span class=\"o\">/</span><span class=\"n\">image</span><span
  class=\"o\">/</span><span class=\"n\">upload</span><span class=\"o\">/</span><span
  class=\"n\">v1652613991525</span><span class=\"o\">/</span><span class=\"n\">hxQgtGOoM</span><span
  class=\"o\">.</span><span class=\"n\">png</span><span class=\"p\">)</span>\n\n<span
  class=\"n\">After</span> <span class=\"n\">committing</span> <span class=\"n\">the</span>
  <span class=\"n\">code</span><span class=\"p\">,</span> <span class=\"n\">we</span>
  <span class=\"n\">can</span> <span class=\"n\">now</span> <span class=\"n\">push</span>
  <span class=\"n\">the</span> <span class=\"n\">code</span> <span class=\"n\">to</span>
  <span class=\"n\">GitHub</span><span class=\"o\">.</span> <span class=\"n\">We</span>
  <span class=\"n\">first</span> <span class=\"n\">need</span> <span class=\"n\">to</span>
  <span class=\"nb\">set</span> <span class=\"n\">the</span> <span class=\"n\">remote</span>
  <span class=\"n\">repository</span> <span class=\"n\">reference</span> <span class=\"n\">to</span>
  <span class=\"n\">be</span> <span class=\"n\">able</span> <span class=\"n\">to</span>
  <span class=\"n\">push</span> <span class=\"n\">the</span> <span class=\"n\">code</span>
  <span class=\"n\">to</span> <span class=\"n\">it</span><span class=\"o\">.</span>
  \n\n<span class=\"err\">```</span>\n<span class=\"n\">git</span> <span class=\"n\">remote</span>
  <span class=\"n\">add</span> <span class=\"n\">origin</span> <span class=\"n\">https</span><span
  class=\"p\">:</span><span class=\"o\">//</span><span class=\"n\">github</span><span
  class=\"o\">.</span><span class=\"n\">com</span><span class=\"o\">/&lt;</span><span
  class=\"n\">username</span><span class=\"o\">&gt;/&lt;</span><span class=\"n\">repo_name</span><span
  class=\"o\">&gt;</span>\n<span class=\"err\">```</span>\n<span class=\"n\">This</span>
  <span class=\"n\">will</span> <span class=\"nb\">set</span> <span class=\"n\">up</span>
  <span class=\"n\">the</span> <span class=\"err\">`</span><span class=\"n\">origin</span><span
  class=\"err\">`</span> <span class=\"k\">as</span> <span class=\"n\">the</span>
  <span class=\"n\">remote</span> <span class=\"n\">repository</span> <span class=\"n\">on</span>
  <span class=\"n\">GitHub</span><span class=\"o\">.</span> <span class=\"n\">Once</span>
  <span class=\"n\">the</span> <span class=\"n\">remote</span> <span class=\"n\">repository</span>
  <span class=\"ow\">is</span> <span class=\"n\">created</span><span class=\"p\">,</span>
  <span class=\"n\">we</span> <span class=\"n\">can</span> <span class=\"n\">move</span>
  <span class=\"n\">into</span> <span class=\"n\">the</span> <span class=\"n\">push</span>
  <span class=\"n\">of</span> <span class=\"n\">the</span> <span class=\"n\">git</span>
  <span class=\"n\">repository</span> <span class=\"n\">that</span> <span class=\"n\">will</span>
  <span class=\"n\">trigger</span> <span class=\"n\">the</span> <span class=\"n\">build</span><span
  class=\"o\">.</span> <span class=\"n\">First</span><span class=\"p\">,</span> <span
  class=\"n\">navigate</span> <span class=\"n\">to</span> <span class=\"n\">the</span>
  <span class=\"err\">`</span><span class=\"n\">Deploy</span><span class=\"err\">`</span>
  <span class=\"n\">section</span> <span class=\"ow\">in</span> <span class=\"n\">the</span>
  <span class=\"n\">heroku</span> <span class=\"n\">app</span><span class=\"s1\">&#39;s
  dashboard where we want to connect the `GitHub` repository and allow the automatic
  deploy from a branch in this case we have chosen the `main` branch.</span>\n\n<span
  class=\"n\">Due</span> <span class=\"n\">to</span> <span class=\"n\">some</span>
  <span class=\"err\">`</span><span class=\"n\">Heroku</span><span class=\"err\">`</span>
  <span class=\"n\">Internal</span> <span class=\"n\">Server</span> <span class=\"n\">Issues</span><span
  class=\"p\">,</span> <span class=\"n\">the</span> <span class=\"n\">GitHub</span>
  <span class=\"n\">integration</span> <span class=\"n\">seems</span> <span class=\"n\">to</span>
  <span class=\"n\">have</span> <span class=\"n\">broken</span> <span class=\"ow\">and</span>
  <span class=\"n\">isn</span><span class=\"s1\">&#39;t working as of May 2022, but
  it might work later. </span>\n\n<span class=\"n\">Below</span> <span class=\"ow\">is</span>
  <span class=\"n\">a</span> <span class=\"n\">screenshot</span> <span class=\"n\">of</span>
  <span class=\"n\">my</span> <span class=\"n\">previous</span> <span class=\"n\">project</span>
  <span class=\"n\">deployed</span> <span class=\"n\">to</span> <span class=\"n\">Heroku</span>
  <span class=\"n\">using</span> <span class=\"n\">a</span> <span class=\"n\">GitHub</span>
  <span class=\"n\">repository</span><span class=\"o\">.</span>\n\n<span class=\"err\">!</span><span
  class=\"p\">[</span><span class=\"n\">image</span><span class=\"o\">.</span><span
  class=\"n\">png</span><span class=\"p\">](</span><span class=\"n\">https</span><span
  class=\"p\">:</span><span class=\"o\">//</span><span class=\"n\">cdn</span><span
  class=\"o\">.</span><span class=\"n\">hashnode</span><span class=\"o\">.</span><span
  class=\"n\">com</span><span class=\"o\">/</span><span class=\"n\">res</span><span
  class=\"o\">/</span><span class=\"n\">hashnode</span><span class=\"o\">/</span><span
  class=\"n\">image</span><span class=\"o\">/</span><span class=\"n\">upload</span><span
  class=\"o\">/</span><span class=\"n\">v1652605497382</span><span class=\"o\">/</span><span
  class=\"mi\">5</span><span class=\"n\">VuQUQ0t0</span><span class=\"o\">.</span><span
  class=\"n\">png</span><span class=\"p\">)</span>\n\n<span class=\"err\">```</span>\n<span
  class=\"n\">git</span> <span class=\"n\">push</span> <span class=\"n\">origin</span>
  <span class=\"n\">main</span>\n<span class=\"err\">```</span>\n\n<span class=\"n\">This</span>
  <span class=\"n\">will</span> <span class=\"n\">prompt</span> <span class=\"n\">you</span>
  <span class=\"k\">for</span> <span class=\"n\">your</span> <span class=\"n\">GitHub</span>
  <span class=\"n\">credentials</span> <span class=\"ow\">and</span> <span class=\"n\">will</span>
  <span class=\"n\">deploy</span> <span class=\"n\">the</span> <span class=\"n\">commits</span>
  <span class=\"n\">to</span> <span class=\"n\">the</span> <span class=\"n\">remote</span>
  <span class=\"n\">repository</span> <span class=\"n\">on</span> <span class=\"n\">GitHub</span><span
  class=\"o\">.</span> <span class=\"n\">This</span> <span class=\"n\">push</span>
  <span class=\"n\">on</span> <span class=\"n\">the</span> <span class=\"n\">main</span>
  <span class=\"n\">branch</span> <span class=\"n\">should</span> <span class=\"n\">also</span>
  <span class=\"n\">trigger</span> <span class=\"n\">the</span> <span class=\"n\">build</span>
  <span class=\"n\">process</span> <span class=\"n\">of</span> <span class=\"n\">the</span>
  <span class=\"n\">heroku</span> <span class=\"n\">app</span> <span class=\"k\">for</span>
  <span class=\"n\">this</span> <span class=\"n\">django</span> <span class=\"n\">project</span><span
  class=\"o\">.</span> <span class=\"n\">You</span> <span class=\"n\">can</span> <span
  class=\"n\">navigate</span> <span class=\"n\">to</span> <span class=\"n\">the</span>
  <span class=\"n\">Activity</span> <span class=\"n\">section</span> <span class=\"k\">for</span>
  <span class=\"n\">the</span> <span class=\"n\">Build</span> <span class=\"n\">logs</span><span
  class=\"o\">.</span> \n\n<span class=\"n\">If</span> <span class=\"n\">you</span>
  <span class=\"n\">have</span> <span class=\"n\">followed</span> <span class=\"n\">the</span>
  <span class=\"n\">article</span> <span class=\"n\">well</span><span class=\"p\">,</span>
  <span class=\"ow\">and</span> <span class=\"n\">your</span> <span class=\"n\">repository</span>
  <span class=\"n\">has</span> <span class=\"nb\">all</span> <span class=\"n\">the</span>
  <span class=\"n\">correct</span> <span class=\"n\">configurations</span><span class=\"p\">,</span>
  <span class=\"n\">the</span> <span class=\"n\">build</span> <span class=\"n\">will</span>
  <span class=\"n\">succeed</span><span class=\"p\">,</span> <span class=\"k\">else</span>
  <span class=\"n\">chances</span> <span class=\"n\">are</span> <span class=\"n\">you</span>
  <span class=\"n\">might</span> <span class=\"n\">have</span> <span class=\"n\">missed</span>
  <span class=\"n\">a</span> <span class=\"n\">few</span> <span class=\"n\">things</span>
  <span class=\"ow\">and</span> <span class=\"n\">the</span> <span class=\"n\">app</span>
  <span class=\"n\">might</span> <span class=\"n\">have</span> <span class=\"n\">crashed</span><span
  class=\"o\">.</span> <span class=\"n\">You</span> <span class=\"n\">can</span> <span
  class=\"n\">debug</span> <span class=\"n\">your</span> <span class=\"n\">application</span>
  <span class=\"n\">build</span> <span class=\"k\">with</span> <span class=\"n\">the</span>
  <span class=\"n\">simple</span> <span class=\"n\">heroku</span> <span class=\"n\">CLI</span>
  <span class=\"n\">command</span><span class=\"p\">:</span>\n\n<span class=\"err\">```</span>\n<span
  class=\"n\">heroku</span> <span class=\"n\">logs</span> <span class=\"o\">--</span><span
  class=\"n\">tail</span> <span class=\"o\">-</span><span class=\"n\">a</span> <span
  class=\"o\">&lt;</span><span class=\"n\">app_name</span><span class=\"o\">&gt;</span>\n<span
  class=\"err\">```</span>\n\n<span class=\"n\">This</span> <span class=\"n\">can</span>
  <span class=\"n\">be</span> <span class=\"n\">quite</span> <span class=\"n\">handy</span>
  <span class=\"ow\">and</span> <span class=\"n\">saves</span> <span class=\"n\">a</span>
  <span class=\"n\">lot</span> <span class=\"n\">of</span> <span class=\"n\">time</span>
  <span class=\"ow\">in</span> <span class=\"n\">understanding</span> <span class=\"n\">what</span>
  <span class=\"n\">went</span> <span class=\"n\">wrong</span> <span class=\"ow\">in</span>
  <span class=\"n\">the</span> <span class=\"n\">build</span><span class=\"o\">.</span>
  <span class=\"n\">It</span> <span class=\"n\">might</span> <span class=\"n\">be</span>
  <span class=\"n\">related</span> <span class=\"n\">to</span> <span class=\"n\">database</span>
  <span class=\"n\">migration</span><span class=\"p\">,</span> <span class=\"n\">static</span>
  <span class=\"n\">files</span><span class=\"p\">,</span> <span class=\"n\">python</span>
  <span class=\"n\">package</span> <span class=\"ow\">not</span> <span class=\"n\">found</span><span
  class=\"p\">,</span> <span class=\"ow\">and</span> <span class=\"n\">some</span>
  <span class=\"n\">silly</span> <span class=\"n\">mistakes</span> <span class=\"ow\">and</span>
  <span class=\"n\">errors</span> <span class=\"n\">that</span> <span class=\"n\">can</span>
  <span class=\"n\">be</span> <span class=\"n\">fixed</span> <span class=\"n\">after</span>
  <span class=\"n\">committing</span> <span class=\"n\">the</span> <span class=\"n\">code</span>
  <span class=\"ow\">and</span> <span class=\"n\">pushing</span> <span class=\"n\">it</span>
  <span class=\"n\">to</span> <span class=\"n\">GitHub</span> <span class=\"n\">again</span><span
  class=\"o\">.</span>\n\n<span class=\"n\">If</span> <span class=\"n\">you</span>
  <span class=\"n\">do</span> <span class=\"ow\">not</span> <span class=\"n\">want</span>
  <span class=\"n\">a</span> <span class=\"n\">GitHub</span> <span class=\"n\">repository</span><span
  class=\"p\">,</span> <span class=\"n\">you</span> <span class=\"n\">can</span> <span
  class=\"n\">directly</span> <span class=\"n\">push</span> <span class=\"n\">the</span>
  <span class=\"n\">code</span> <span class=\"kn\">from</span> <span class=\"nn\">the</span>
  <span class=\"n\">local</span> <span class=\"n\">git</span> <span class=\"n\">repository</span>
  <span class=\"n\">to</span> <span class=\"n\">the</span> <span class=\"n\">remote</span>
  <span class=\"n\">heroku</span> <span class=\"n\">app</span> <span class=\"n\">center</span><span
  class=\"o\">.</span> <span class=\"n\">This</span> <span class=\"n\">will</span>
  <span class=\"n\">require</span> <span class=\"n\">us</span> <span class=\"n\">the</span>
  <span class=\"n\">Heroku</span> <span class=\"n\">CLI</span><span class=\"o\">.</span>\n\n<span
  class=\"c1\">## Heroku CLI</span>\n\n<span class=\"n\">We</span> <span class=\"n\">can</span>
  <span class=\"n\">use</span> <span class=\"n\">the</span> <span class=\"n\">heroku</span>
  <span class=\"n\">CLI</span> <span class=\"k\">for</span> <span class=\"n\">pushing</span>
  <span class=\"n\">the</span> <span class=\"n\">code</span> <span class=\"n\">via</span>
  <span class=\"n\">the</span> <span class=\"n\">simple</span> <span class=\"n\">git</span>
  <span class=\"n\">repository</span><span class=\"o\">.</span> <span class=\"n\">We</span>
  <span class=\"n\">can</span> <span class=\"n\">push</span> <span class=\"n\">the</span>
  <span class=\"n\">references</span> <span class=\"n\">via</span> <span class=\"n\">the</span>
  <span class=\"n\">branch</span> <span class=\"ow\">and</span> <span class=\"n\">a</span>
  <span class=\"n\">remote</span> <span class=\"n\">repository</span> <span class=\"n\">on</span>
  <span class=\"n\">heroku</span> <span class=\"n\">to</span> <span class=\"n\">build</span>
  <span class=\"n\">our</span> <span class=\"n\">app</span><span class=\"o\">.</span>
  \ <span class=\"n\">For</span> <span class=\"n\">this</span><span class=\"p\">,</span>
  <span class=\"n\">we</span> <span class=\"n\">assume</span> <span class=\"n\">you</span>
  <span class=\"n\">have</span> <span class=\"n\">heroku</span> <span class=\"n\">installed</span>
  <span class=\"ow\">and</span> <span class=\"n\">logged</span> <span class=\"ow\">in</span><span
  class=\"o\">.</span> <span class=\"n\">We</span> <span class=\"n\">will</span> <span
  class=\"n\">require</span> <span class=\"n\">the</span> <span class=\"n\">django</span>
  <span class=\"n\">project</span> <span class=\"n\">code</span> <span class=\"ow\">and</span>
  <span class=\"n\">heroku</span> <span class=\"n\">cli</span> <span class=\"n\">to</span>
  <span class=\"n\">build</span> <span class=\"n\">the</span> <span class=\"n\">django</span>
  <span class=\"n\">web</span> <span class=\"n\">application</span><span class=\"o\">.</span>\n\n<span
  class=\"err\">```</span><span class=\"n\">bash</span>\n<span class=\"n\">heroku</span>
  <span class=\"n\">git</span><span class=\"p\">:</span><span class=\"n\">remote</span>
  <span class=\"o\">-</span><span class=\"n\">a</span> <span class=\"o\">&lt;</span><span
  class=\"n\">heroku_app_name</span><span class=\"o\">&gt;</span>\n\n<span class=\"c1\">#
  for my case</span>\n<span class=\"n\">heroku</span> <span class=\"n\">git</span><span
  class=\"p\">:</span><span class=\"n\">remote</span> <span class=\"o\">-</span><span
  class=\"n\">a</span> <span class=\"n\">blog</span><span class=\"o\">-</span><span
  class=\"n\">django</span><span class=\"o\">-</span><span class=\"n\">dep</span>\n<span
  class=\"err\">```</span>\n\n<span class=\"err\">!</span><span class=\"p\">[</span><span
  class=\"n\">image</span><span class=\"o\">.</span><span class=\"n\">png</span><span
  class=\"p\">](</span><span class=\"n\">https</span><span class=\"p\">:</span><span
  class=\"o\">//</span><span class=\"n\">cdn</span><span class=\"o\">.</span><span
  class=\"n\">hashnode</span><span class=\"o\">.</span><span class=\"n\">com</span><span
  class=\"o\">/</span><span class=\"n\">res</span><span class=\"o\">/</span><span
  class=\"n\">hashnode</span><span class=\"o\">/</span><span class=\"n\">image</span><span
  class=\"o\">/</span><span class=\"n\">upload</span><span class=\"o\">/</span><span
  class=\"n\">v1652614221069</span><span class=\"o\">/</span><span class=\"n\">vCAKD0zsz</span><span
  class=\"o\">.</span><span class=\"n\">png</span><span class=\"p\">)</span>\n\n<span
  class=\"n\">After</span> <span class=\"n\">this</span><span class=\"p\">,</span>
  <span class=\"n\">you</span> <span class=\"n\">can</span> <span class=\"n\">commit</span>
  <span class=\"n\">your</span> <span class=\"n\">code</span> <span class=\"ow\">and</span>
  <span class=\"n\">the</span> <span class=\"n\">project</span> <span class=\"k\">as</span>
  <span class=\"n\">git</span> <span class=\"n\">repository</span><span class=\"o\">.</span>
  <span class=\"n\">We</span> <span class=\"n\">have</span> <span class=\"n\">added</span>
  <span class=\"n\">the</span> <span class=\"n\">remote</span> <span class=\"n\">repository</span>
  <span class=\"n\">location</span> <span class=\"n\">on</span> <span class=\"n\">heroku</span><span
  class=\"p\">,</span> <span class=\"n\">we</span> <span class=\"n\">can</span> <span
  class=\"n\">now</span> <span class=\"n\">simply</span> <span class=\"n\">push</span>
  <span class=\"n\">the</span> <span class=\"n\">code</span> <span class=\"n\">to</span>
  <span class=\"n\">the</span> <span class=\"n\">remote</span> <span class=\"n\">repository</span><span
  class=\"o\">.</span>\n\n<span class=\"err\">```</span>\n<span class=\"n\">git</span>
  <span class=\"n\">push</span> <span class=\"n\">heroku</span> <span class=\"n\">main</span>\n<span
  class=\"err\">```</span>\n\n<span class=\"err\">!</span><span class=\"p\">[</span><span
  class=\"n\">image</span><span class=\"o\">.</span><span class=\"n\">png</span><span
  class=\"p\">](</span><span class=\"n\">https</span><span class=\"p\">:</span><span
  class=\"o\">//</span><span class=\"n\">cdn</span><span class=\"o\">.</span><span
  class=\"n\">hashnode</span><span class=\"o\">.</span><span class=\"n\">com</span><span
  class=\"o\">/</span><span class=\"n\">res</span><span class=\"o\">/</span><span
  class=\"n\">hashnode</span><span class=\"o\">/</span><span class=\"n\">image</span><span
  class=\"o\">/</span><span class=\"n\">upload</span><span class=\"o\">/</span><span
  class=\"n\">v1652614125785</span><span class=\"o\">/</span><span class=\"n\">uEzFQ9VvQ</span><span
  class=\"o\">.</span><span class=\"n\">png</span><span class=\"p\">)</span>\n\n<span
  class=\"n\">Here</span><span class=\"p\">,</span> <span class=\"err\">`</span><span
  class=\"n\">heroku</span><span class=\"err\">`</span> <span class=\"ow\">is</span>
  <span class=\"n\">the</span> <span class=\"n\">remote</span> <span class=\"n\">repository</span>
  <span class=\"n\">location</span> <span class=\"ow\">and</span> <span class=\"err\">`</span><span
  class=\"n\">main</span><span class=\"err\">`</span> <span class=\"ow\">is</span>
  <span class=\"n\">the</span> <span class=\"n\">branch</span> <span class=\"n\">of</span>
  <span class=\"n\">the</span> <span class=\"n\">repository</span><span class=\"o\">.</span>
  <span class=\"n\">This</span> <span class=\"n\">will</span> <span class=\"n\">push</span>
  <span class=\"n\">the</span> <span class=\"n\">code</span> <span class=\"n\">to</span>
  <span class=\"n\">the</span> <span class=\"n\">repository</span> <span class=\"ow\">and</span>
  <span class=\"n\">thereafter</span> <span class=\"n\">create</span> <span class=\"n\">a</span>
  <span class=\"n\">build</span> <span class=\"n\">to</span> <span class=\"n\">deploy</span>
  <span class=\"n\">the</span> <span class=\"n\">Django</span> <span class=\"n\">project</span>
  <span class=\"k\">as</span> <span class=\"n\">a</span> <span class=\"n\">Heroku</span>
  <span class=\"n\">application</span><span class=\"o\">.</span>\n\n<span class=\"err\">!</span><span
  class=\"p\">[</span><span class=\"n\">image</span><span class=\"o\">.</span><span
  class=\"n\">png</span><span class=\"p\">](</span><span class=\"n\">https</span><span
  class=\"p\">:</span><span class=\"o\">//</span><span class=\"n\">cdn</span><span
  class=\"o\">.</span><span class=\"n\">hashnode</span><span class=\"o\">.</span><span
  class=\"n\">com</span><span class=\"o\">/</span><span class=\"n\">res</span><span
  class=\"o\">/</span><span class=\"n\">hashnode</span><span class=\"o\">/</span><span
  class=\"n\">image</span><span class=\"o\">/</span><span class=\"n\">upload</span><span
  class=\"o\">/</span><span class=\"n\">v1652614381808</span><span class=\"o\">/</span><span
  class=\"n\">kYTmB3EO2p</span><span class=\"o\">.</span><span class=\"n\">png</span><span
  class=\"p\">)</span>\n\n<span class=\"n\">You</span> <span class=\"n\">can</span>
  <span class=\"n\">hit</span> <span class=\"n\">the</span> <span class=\"err\">`</span><span
  class=\"n\">Open</span> <span class=\"n\">App</span><span class=\"err\">`</span>
  <span class=\"n\">button</span> <span class=\"n\">on</span> <span class=\"n\">the</span>
  <span class=\"n\">top</span> <span class=\"n\">right</span> <span class=\"n\">corner</span>
  <span class=\"ow\">and</span> <span class=\"n\">there</span> <span class=\"n\">should</span>
  <span class=\"n\">be</span> <span class=\"n\">your</span> <span class=\"n\">deployed</span>
  <span class=\"n\">Django</span> <span class=\"n\">application</span><span class=\"o\">.</span>
  \n\n<span class=\"err\">!</span><span class=\"p\">[</span><span class=\"n\">image</span><span
  class=\"o\">.</span><span class=\"n\">png</span><span class=\"p\">](</span><span
  class=\"n\">https</span><span class=\"p\">:</span><span class=\"o\">//</span><span
  class=\"n\">cdn</span><span class=\"o\">.</span><span class=\"n\">hashnode</span><span
  class=\"o\">.</span><span class=\"n\">com</span><span class=\"o\">/</span><span
  class=\"n\">res</span><span class=\"o\">/</span><span class=\"n\">hashnode</span><span
  class=\"o\">/</span><span class=\"n\">image</span><span class=\"o\">/</span><span
  class=\"n\">upload</span><span class=\"o\">/</span><span class=\"n\">v1652610395538</span><span
  class=\"o\">/</span><span class=\"n\">xjUiODhoK</span><span class=\"o\">.</span><span
  class=\"n\">png</span><span class=\"p\">)</span>\n\n\n<span class=\"c1\">## A Few
  Tricks and Gotchas</span>\n\n<span class=\"n\">There</span> <span class=\"n\">are</span>
  <span class=\"n\">a</span> <span class=\"n\">few</span> <span class=\"n\">tricks</span>
  <span class=\"ow\">and</span> <span class=\"n\">issues</span> <span class=\"n\">that</span>
  <span class=\"n\">you</span> <span class=\"n\">might</span> <span class=\"n\">pop</span>
  <span class=\"n\">into</span> <span class=\"k\">while</span> <span class=\"n\">deploying</span>
  <span class=\"n\">a</span> <span class=\"n\">django</span> <span class=\"n\">project</span>
  <span class=\"n\">on</span> <span class=\"n\">heroku</span><span class=\"p\">,</span>
  <span class=\"n\">especially</span> <span class=\"n\">the</span> <span class=\"n\">build</span>
  <span class=\"n\">process</span><span class=\"o\">.</span> <span class=\"n\">It</span>
  <span class=\"n\">requires</span> <span class=\"n\">a</span> <span class=\"n\">few</span>
  <span class=\"n\">iterations</span> <span class=\"n\">to</span> <span class=\"n\">get</span>
  <span class=\"n\">the</span> <span class=\"n\">complete</span> <span class=\"n\">app</span>
  <span class=\"n\">setup</span><span class=\"o\">.</span>\n\n<span class=\"c1\">###
  Run command from the Dashboard console</span>\n\n<span class=\"n\">If</span> <span
  class=\"n\">you</span> <span class=\"n\">don</span><span class=\"s1\">&#39;t have
  heroku CLI set up and want to fix a few things on the pushed code to the heroku
  app, you can use the `Run Console` option from the `More` Tab on the top right corner
  of theApplication dashboard. This is a great way to run migrations, configure static
  files or tweak a few things here and there without messing up the code on GitHub
  or the remote git repositories. </span>\n\n<span class=\"err\">!</span><span class=\"p\">[</span><span
  class=\"n\">image</span><span class=\"o\">.</span><span class=\"n\">png</span><span
  class=\"p\">](</span><span class=\"n\">https</span><span class=\"p\">:</span><span
  class=\"o\">//</span><span class=\"n\">cdn</span><span class=\"o\">.</span><span
  class=\"n\">hashnode</span><span class=\"o\">.</span><span class=\"n\">com</span><span
  class=\"o\">/</span><span class=\"n\">res</span><span class=\"o\">/</span><span
  class=\"n\">hashnode</span><span class=\"o\">/</span><span class=\"n\">image</span><span
  class=\"o\">/</span><span class=\"n\">upload</span><span class=\"o\">/</span><span
  class=\"n\">v1652614775294</span><span class=\"o\">/</span><span class=\"n\">lgDPwr2yr</span><span
  class=\"o\">.</span><span class=\"n\">png</span><span class=\"p\">)</span>\n\n\n<span
  class=\"err\">!</span><span class=\"p\">[</span><span class=\"n\">image</span><span
  class=\"o\">.</span><span class=\"n\">png</span><span class=\"p\">](</span><span
  class=\"n\">https</span><span class=\"p\">:</span><span class=\"o\">//</span><span
  class=\"n\">cdn</span><span class=\"o\">.</span><span class=\"n\">hashnode</span><span
  class=\"o\">.</span><span class=\"n\">com</span><span class=\"o\">/</span><span
  class=\"n\">res</span><span class=\"o\">/</span><span class=\"n\">hashnode</span><span
  class=\"o\">/</span><span class=\"n\">image</span><span class=\"o\">/</span><span
  class=\"n\">upload</span><span class=\"o\">/</span><span class=\"n\">v1652614821950</span><span
  class=\"o\">/</span><span class=\"n\">uTzQVB8sC</span><span class=\"o\">.</span><span
  class=\"n\">png</span><span class=\"p\">)</span>\n\n\n<span class=\"err\">!</span><span
  class=\"p\">[</span><span class=\"n\">image</span><span class=\"o\">.</span><span
  class=\"n\">png</span><span class=\"p\">](</span><span class=\"n\">https</span><span
  class=\"p\">:</span><span class=\"o\">//</span><span class=\"n\">cdn</span><span
  class=\"o\">.</span><span class=\"n\">hashnode</span><span class=\"o\">.</span><span
  class=\"n\">com</span><span class=\"o\">/</span><span class=\"n\">res</span><span
  class=\"o\">/</span><span class=\"n\">hashnode</span><span class=\"o\">/</span><span
  class=\"n\">image</span><span class=\"o\">/</span><span class=\"n\">upload</span><span
  class=\"o\">/</span><span class=\"n\">v1652614845269</span><span class=\"o\">/</span><span
  class=\"n\">BkZhu3SGH</span><span class=\"o\">.</span><span class=\"n\">png</span><span
  class=\"p\">)</span>\n\n<span class=\"c1\">### Deploy with Docker </span>\n\n<span
  class=\"n\">You</span> <span class=\"n\">can</span> <span class=\"n\">even</span>
  <span class=\"n\">use</span> <span class=\"n\">the</span> <span class=\"n\">docker</span>
  <span class=\"n\">container</span> <span class=\"n\">to</span> <span class=\"n\">deploy</span>
  <span class=\"n\">a</span> <span class=\"n\">Django</span> <span class=\"n\">application</span>
  <span class=\"n\">on</span> <span class=\"n\">Heroku</span><span class=\"o\">.</span>
  <span class=\"n\">It</span> <span class=\"ow\">is</span> <span class=\"n\">a</span>
  <span class=\"n\">great</span> <span class=\"n\">way</span> <span class=\"n\">of</span>
  <span class=\"n\">learning</span> <span class=\"n\">a</span> <span class=\"n\">lot</span>
  <span class=\"n\">of</span> <span class=\"n\">deployment</span> <span class=\"n\">strategies</span>
  <span class=\"ow\">and</span> <span class=\"n\">techniques</span> <span class=\"n\">using</span>
  <span class=\"n\">Docker</span><span class=\"o\">.</span> <span class=\"n\">You</span><span
  class=\"s1\">&#39;ll get familiar with interesting concepts like virtualization,
  and containerization, and also learn Docker on the way. You can follow this tutorial
  on [Deploying Django applications with Docker on Heroku](https://testdriven.io/blog/deploying-django-to-heroku-with-docker/).
  Also, you can check out the official Heroku documentation for [deploying python
  applications](https://devcenter.heroku.com/articles/deploying-python).</span>\n\n<span
  class=\"c1\">### What are Dynos?</span>\n\n<span class=\"n\">Dynos</span> <span
  class=\"n\">are</span> <span class=\"n\">simply</span> <span class=\"n\">web</span>
  <span class=\"n\">processes</span> <span class=\"ow\">or</span> <span class=\"n\">workers</span>
  <span class=\"n\">that</span> <span class=\"n\">serve</span> <span class=\"n\">your</span>
  <span class=\"n\">web</span> <span class=\"n\">application</span><span class=\"o\">.</span>
  <span class=\"n\">Dynos</span> <span class=\"ow\">in</span> <span class=\"n\">Heroku</span>
  <span class=\"n\">are</span> <span class=\"n\">allocated</span> <span class=\"n\">based</span>
  <span class=\"n\">on</span> <span class=\"n\">the</span> <span class=\"n\">build</span>
  <span class=\"n\">process</span><span class=\"p\">,</span> <span class=\"n\">once</span>
  <span class=\"n\">the</span> <span class=\"n\">slug</span> <span class=\"ow\">is</span>
  <span class=\"n\">created</span> <span class=\"n\">a</span> <span class=\"n\">dyno</span>
  <span class=\"ow\">is</span> <span class=\"n\">created</span> <span class=\"k\">as</span>
  <span class=\"n\">it</span> <span class=\"n\">runs</span> <span class=\"n\">on</span>
  <span class=\"n\">a</span> <span class=\"n\">VM</span> <span class=\"n\">container</span><span
  class=\"o\">.</span> <span class=\"n\">This</span> <span class=\"n\">simply</span>
  <span class=\"n\">means</span> <span class=\"n\">there</span> <span class=\"n\">are</span>
  <span class=\"n\">limitations</span> <span class=\"n\">on</span> <span class=\"n\">how</span>
  <span class=\"n\">to</span> <span class=\"n\">use</span> <span class=\"n\">the</span>
  <span class=\"n\">web</span> <span class=\"n\">application</span> <span class=\"ow\">and</span>
  <span class=\"n\">its</span> <span class=\"n\">sleep</span> <span class=\"n\">process</span><span
  class=\"o\">.</span> <span class=\"n\">The</span> <span class=\"n\">hobby</span>
  <span class=\"n\">tier</span> <span class=\"ow\">is</span> <span class=\"n\">sufficient</span>
  <span class=\"k\">for</span> <span class=\"n\">normal</span> <span class=\"n\">testing</span>
  <span class=\"n\">projects</span> <span class=\"ow\">and</span> <span class=\"n\">side</span>
  <span class=\"n\">projects</span> <span class=\"n\">though</span> <span class=\"n\">you</span>
  <span class=\"n\">will</span> <span class=\"n\">have</span> <span class=\"n\">to</span>
  <span class=\"n\">pay</span> <span class=\"ow\">and</span> <span class=\"n\">move</span>
  <span class=\"n\">into</span> <span class=\"n\">advance</span> <span class=\"n\">tiers</span>
  <span class=\"n\">to</span> <span class=\"n\">increase</span> <span class=\"n\">the</span>
  <span class=\"n\">dyno</span> <span class=\"n\">allocations</span> <span class=\"ow\">and</span>
  <span class=\"n\">scaling</span> <span class=\"n\">of</span> <span class=\"n\">those</span>
  <span class=\"n\">web</span> <span class=\"n\">processes</span><span class=\"o\">.</span>
  \n\n<span class=\"n\">It</span><span class=\"s1\">&#39;s not a simple thing to understand
  but to keep it simple, it might be a container to process the client&#39;</span><span
  class=\"n\">s</span> <span class=\"n\">request</span> <span class=\"ow\">and</span>
  <span class=\"n\">serve</span> <span class=\"n\">the</span> <span class=\"n\">page</span>
  <span class=\"k\">for</span> <span class=\"n\">a</span> <span class=\"n\">finite</span>
  <span class=\"n\">duration</span> <span class=\"n\">of</span> <span class=\"n\">the</span>
  <span class=\"n\">interaction</span><span class=\"o\">.</span> <span class=\"n\">Also</span><span
  class=\"p\">,</span> <span class=\"n\">your</span> <span class=\"n\">application</span>
  <span class=\"n\">will</span> <span class=\"n\">sleep</span> <span class=\"n\">after</span>
  <span class=\"n\">half</span> <span class=\"n\">an</span> <span class=\"n\">hour</span><span
  class=\"p\">,</span> <span class=\"k\">if</span> <span class=\"n\">you</span> <span
  class=\"k\">try</span> <span class=\"n\">to</span> <span class=\"n\">reload</span>
  <span class=\"n\">the</span> <span class=\"n\">application</span> <span class=\"n\">every</span>
  <span class=\"n\">half</span> <span class=\"n\">an</span> <span class=\"n\">hour</span>
  <span class=\"n\">it</span> <span class=\"n\">will</span> <span class=\"n\">consume</span>
  <span class=\"n\">your</span> <span class=\"n\">resource</span> <span class=\"n\">allocation</span>
  <span class=\"k\">for</span> <span class=\"n\">the</span> <span class=\"n\">month</span>
  <span class=\"ow\">and</span> <span class=\"n\">this</span> <span class=\"ow\">is</span>
  <span class=\"n\">how</span> <span class=\"n\">the</span> <span class=\"n\">tiers</span>
  <span class=\"ow\">and</span> <span class=\"n\">divided</span> <span class=\"k\">for</span>
  <span class=\"n\">paid</span> <span class=\"n\">services</span> <span class=\"n\">on</span>
  <span class=\"n\">Heroku</span><span class=\"o\">.</span> <span class=\"n\">You</span>
  <span class=\"n\">can</span> <span class=\"n\">check</span> <span class=\"n\">out</span>
  <span class=\"n\">the</span> <span class=\"n\">detail</span> <span class=\"n\">over</span>
  <span class=\"p\">[</span><span class=\"n\">here</span><span class=\"p\">](</span><span
  class=\"n\">https</span><span class=\"p\">:</span><span class=\"o\">//</span><span
  class=\"n\">www</span><span class=\"o\">.</span><span class=\"n\">heroku</span><span
  class=\"o\">.</span><span class=\"n\">com</span><span class=\"o\">/</span><span
  class=\"n\">pricing</span><span class=\"c1\">#containers).</span>\n\n<span class=\"c1\">##
  Conclusion</span>\n\n<span class=\"n\">So</span><span class=\"p\">,</span> <span
  class=\"n\">that</span> <span class=\"ow\">is</span> <span class=\"n\">one</span>
  <span class=\"n\">of</span> <span class=\"n\">the</span> <span class=\"n\">ways</span>
  <span class=\"n\">we</span> <span class=\"n\">can</span> <span class=\"n\">deploy</span>
  <span class=\"n\">a</span> <span class=\"n\">Django</span> <span class=\"n\">application</span>
  <span class=\"n\">on</span> <span class=\"n\">Heroku</span> <span class=\"k\">with</span>
  <span class=\"n\">the</span> <span class=\"n\">PostgreSQL</span> <span class=\"n\">database</span><span
  class=\"o\">.</span> <span class=\"n\">You</span> <span class=\"n\">can</span> <span
  class=\"n\">find</span> <span class=\"n\">the</span> <span class=\"p\">[</span><span
  class=\"n\">django</span><span class=\"o\">-</span><span class=\"n\">blog</span>
  <span class=\"n\">project</span><span class=\"p\">]</span> <span class=\"n\">on</span>
  <span class=\"p\">[</span><span class=\"n\">GitHub</span><span class=\"p\">]</span>
  <span class=\"k\">for</span> <span class=\"n\">following</span> <span class=\"n\">along</span>
  <span class=\"k\">with</span> <span class=\"n\">the</span> <span class=\"n\">deployment</span>
  <span class=\"n\">process</span><span class=\"o\">.</span>  <span class=\"n\">In</span>
  <span class=\"n\">the</span> <span class=\"nb\">next</span> <span class=\"n\">few</span>
  <span class=\"n\">parts</span> <span class=\"n\">of</span> <span class=\"n\">the</span>
  <span class=\"n\">series</span><span class=\"p\">,</span> <span class=\"n\">we</span>
  <span class=\"n\">will</span> <span class=\"n\">be</span> <span class=\"n\">hopefully</span>
  <span class=\"n\">covering</span> <span class=\"n\">other</span> <span class=\"n\">platforms</span>
  <span class=\"n\">where</span> <span class=\"n\">you</span> <span class=\"n\">can</span>
  <span class=\"n\">deploy</span> <span class=\"n\">a</span> <span class=\"n\">Django</span>
  <span class=\"n\">o</span> <span class=\"n\">application</span><span class=\"o\">.</span>\n\n<span
  class=\"n\">Hopefully</span><span class=\"p\">,</span> <span class=\"n\">you</span>
  <span class=\"n\">liked</span> <span class=\"n\">the</span> <span class=\"n\">above</span>
  <span class=\"n\">tutorial</span><span class=\"p\">,</span> <span class=\"k\">if</span>
  <span class=\"n\">you</span> <span class=\"n\">have</span> <span class=\"nb\">any</span>
  <span class=\"n\">questions</span><span class=\"o\">.</span> <span class=\"n\">feedback</span><span
  class=\"p\">,</span> <span class=\"ow\">or</span> <span class=\"n\">queries</span><span
  class=\"p\">,</span> <span class=\"n\">you</span> <span class=\"n\">can</span> <span
  class=\"n\">contact</span> <span class=\"n\">me</span> <span class=\"n\">on</span>
  <span class=\"n\">the</span> <span class=\"n\">Social</span> <span class=\"n\">handles</span>
  <span class=\"n\">provided</span> <span class=\"n\">below</span><span class=\"o\">.</span>
  <span class=\"n\">Thank</span> <span class=\"n\">you</span> <span class=\"k\">for</span>
  <span class=\"n\">reading</span> <span class=\"ow\">and</span> <span class=\"n\">till</span>
  <span class=\"n\">the</span> <span class=\"nb\">next</span> <span class=\"n\">post</span>
  <span class=\"n\">Happy</span> <span class=\"n\">Coding</span> <span class=\"p\">:)</span>\n</pre></div>\n\n</pre>\n\n"
image_url: https://meetgor-cdn.pages.dev/django-deploy-heroku.png
long_description: 'Django projects are quite easy to build and simple to understand,
  you might have created a Django application and wanted to show it to the world?
  You can deploy a basic Django application with a database(PostgreSQL) with Heroku.
  It provides a decent '
now: 2025-05-01 18:11:33.314481
path: blog/posts/2022-05-15-Django-Deploy-Heroku.md
prevnext: null
series:
- Django-Deployment
- Django-Series
series_description: 'Django Deployment is a series for understanding various platforms
  for Django project deployment with database attachment. Platforms to be covered:
  Heroku, Railway, Qovery, Python Anywhere, etc.'
slug: django-deploy-heroku
status: published
tags:
- django
- web-development
- python
templateKey: blog-post
title: Django + PostgreSQL Deployment on Heroku
today: 2025-05-01
---

## Introduction

Django projects are quite easy to build and simple to understand, you might have created a Django application and wanted to show it to the world? You can deploy a basic Django application with a database(PostgreSQL) with Heroku. It provides a decent free tier with some great features and add-ons. A free tier Heroku account has a limitation of 5 apps, limited data in the database, limited connections to the server per month, and so on.  

Though the free tier is not a great option for bigger applications, it suits really well for smaller scale and ide projects, so we will be looking into the details of how to deploy a Django application to [Heroku](https://heroku.com/) which is a Platform as Service (PaS). 

This series will be an extension of the series [Django basics](https://techstructiveblog.hashnode.dev/series/django-basics) which covered the basics of the Django framework, we covered from basic Django fundamentals to building a CRUD API. In this series, we will be exploring some platforms for giving a better understanding of how things work and understanding a few components that were left as default while understanding the basics of Django. Let's get started with [Django Deployment](https://techstructiveblog.hashnode.dev/series/django-deployment)!

## Creating a Django Application

For deploying an app, we definitely need an app, we need to create a basic Django application to deploy on the web. We'll be creating a simple blog application with a couple of views and a simple model structure. As for the database, we'll be using Postgres as Heroku has an add-on for it and it is pretty easy to configure. 

### Set up a virtual environment

We need to set up a virtual environment in order to keep the Django project neat and tidy by managing the project-specific dependencies and packages. We can use the [virtualenv](https://pypi.org/project/virtualenv/) package to isolate a python project from the rest of the system.

```
# install the virtualenv package
pip install virtualenv

# create a virtual env for the project
virtualenv .venv

# activate the virtualenv
Windows:
.venv\Scripts\activate

Linux/macOS:
source .venv/bin/activate
```

This will set up the project nicely for a Django project, you now install the core Django package and get started with creating a Django application.


```bash
# install django
pip install django

# start a django project
django-admin startproject blog .

cd blog

# create a application in django project
python manage.py createapp api

# Create some models, views, URLs, templates

# run the server
python manag.py runserver
```

We assume you already have a Django project configured with some basic URLs, views and templates or static files as per your project and requirements, for this tutorial I will be using the simple blog application from my previous Django tutorials as a reference. You can follow along with my [Django Basics](https://techstructiveblog.hashnode.dev/series/django-basics) series and refer to the Blog Application project on [GitHub](https://github.com/Mr-Destructive/django-blog).

## Configuring the Django Application

Make sure to create and activate the virtual environment for this django project. This should be done to manage the dependencies and packages used in the project. If you are not aware of the virtual environment and django setup, you can follow up with this [post](https://mr-destructive.github.io/techstructive-blog/django-setup-script/).

### Creating a runtime.txt file

Now, Firstly we need to specify which type and version of language we are using. Since Django is a Python-based web framework, we need to select the python version in a text file.

**runtime.txt**
```
python-3.9.5
```
 
Here, the version can be anything as per your project and the packages installed.  

### Creating requirements.txt file

We'll first create a `requirements.txt` file for storing all the dependencies and packages installed in the application. This will help in installing dependencies while deploying the application. We can either use a `requirements.txt` file using `virtualenv` or a `Pipfile` using Pipenv. Both serve the same purpose but a bit differently. 

Assuming you are in an isolated virtual environment for this Django project, you can create a requirements.txt file using the below command:

Make sure the virtualenv is activated before running the command:

```
pip freeze > requirements.txt
```

This will create a simple text file that contains the package names along with the versions used in the current virtual environment. A simple Django requirements file would look something like this:

```
asgiref==3.4.1
Django==3.2.11
pytz==2021.3
sqlparse==0.4.2
typing_extensions==4.0.1
```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652614060461/kPTZ9R8Xp.png)

This is vanilla Django without any additional dependencies but if you have integrated other functionalities like Django Rest Framework, PostgreSQL, Crispy Forms, Schedulers, etc. there might be additional dependencies that become quite crucial for the smooth working of the project.  

If you are using pipenv, you don't need to make any efforts to manually activate and manage virtual environment, the dependencies are installed and taken care of by the pipenv installer. You just need to make sure to install any package with `pipenv install` and not `pip install` for better and improved package tracking.

So, that's all we need to take care of packages and keep a list of these integrated packages for the project. You need to update the requirements.txt file every time you install any new package or modify the dependencies for a project. Simply use the command `pip freeze > requirements.txt` in the activated virtual environment.  

### Creating a Procfile

Next up, we have the `Procfile`, a procfile is a special file that holds information about the processes to be run to start or activate the project. In our case, for django we need a web process that can manage the server.

A Procfile is a simple file without any extension, make sure to write `Procfile` as it is as the name of the file in the root folder of the project. Inside the file add the following contents:

**Procfile**
```Procfile
web: gunicorn <project_name>.wsgi
```

As we can see we have defined the `web` process using `gunicorn`, [Gunicorn](https://pypi.org/project/gunicorn/) is a python package that helps in creating WSGI HTTP Server for the UNIX operating systems. So, we need to install the package and update the package dependency list. 

```
pip install gunicorn

pip freeze > requirements.txt
``` 

That would be good to go for creating and serving up the project while deploying the project on Heroku.

## Creating a Heroku App

A Heroku App is basically like your Django Project, you can create apps for deploying your django projects on Heroku. You are limited to 5 apps on the Free tier but that can be expanded on the paid plans.  

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652456732519/cyOQZ3UZK.png)

The name of your heroku app should be unique globally, you need to try a few combinations before a good one fits your project. This name has no significance on your django project code, though it would be the name from which you would access the web application as a name `<app-name>.herokuapp.com`.   

So, choose it wisely if you are not attaching a custom domain. You can attach a custom domain, you can navigate to the `domain` section in the settings tab. 


## Setting up the database 

To set up and configure a database in django on Heroku, we need to manually acquire and attach a PostgreSQL add-on to the heroku app.

- Firstly locate to the Resources Tab in your Heroku app.
- Search `postgres` in the Add-ons Search bar
- Click on the `Heroku Postgres` Add-on
- Submit the Order Form and you have the add-on enabled in the app.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652456842273/ijeWsVdOf.png)

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652456877447/dLG30ac_m.png)

Alternately, you can use the Heroku CLI,

### Heroku CLI Setup

You can use the Heroku CLI which is a command-line interface for managing and creating Heroku applications.  You need to first install the CLI by heading towards the [heroku documentation](https://devcenter.heroku.com/articles/heroku-command-line). Once the CLI is installed, you need to login into your Heroku account by entering the following command:

```
heroku login

``` 

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652605604920/HnTr2KbTi.png)

This will allow us to work with heroku commands and manage our heroku application from the command line itself. The below command will create a add-on for `heroku-postgres` for the application provided as the parameter options 

```
heroku addons:create heroku-postgresql:hobby-dev --app <app_name>
```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652507166978/i1IJ5EGjJ.png)

This should hopefully add a fresh instance of a postgres database for your heroku app. You can now configure the database for your project, we simply need the Database URL from the heroku app dashboard. We'll see how to configure the environment variables in Django for Heroku to keep your secrets like the `SECRET_KEY`, `DATABSE_URL`, etc. 

If you want MySQL as a database, you can check out the [ClearDB](https://devcenter.heroku.com/articles/cleardb) Add-On for Heroku to simply attach it like Postgres on your Heroku application. Also, if you wish to add [MongoDB](https://www.mongodb.com/compatibility/mongodb-and-django) into your Django application on Heroku, you can use [Object Rocket](https://devcenter.heroku.com/articles/ormongo)(OR Mongo). It is not available in the free tier though, unlike PostgreSQL and MySQL.

## Configuring Environment Variables

We need to keep our secrets for the django project out of the deployed code and in a safe place, we can create environment variables and keep them in a `.env` file which will be git-ignored. We do not want this `.env` file to be open source and thus should not be committed.  



We first need to create a new secret key if you already have a GitHub repository, chances are you would have committed the default secret key open for the world to see, it is an insecure way of deploying django apps in production.

So, open up a terminal and launch a python REPL:

```
python

import secrets
secrets.token_hex(24)
```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652512239319/5AACaTGOD.png)

This should generate a secret key that only you know now. So, just copy the key without the quotes and create a file `.env` in the root project folder. 

**.env**
```
SECRET_KEY=76419fd6885a677f802fd1d2b5acd0188e23e001042b05a8
```

The `.env` file should also be added to the `.gitignore` file, so simply append the following in the `.gitignore` file

```
.env
```
This file is only created to test the project locally, so we need to also make this key available on heroku. For doing that we need to add Config Variables to the heroku app. 

- Locate to the Settings Tab in your Heroku Application Dashboard
- We have the `Config Vars` section in the located tab
= We need to reveal those variables and we will be able to see all the variables.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652456988713/5VM6E29_o.png)

We already have a `DATABASE_URL` variable declared when we attached the `django-postgres` database to our application. We will now add one more variable to the Application, i.e. the `SECRET_KEY`. Simply, enter the name of the key and also enter the value of the key, so basically a key-value pair. Simply click on the `Add` button and this will add the variable to your application.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652515244870/LRyPzJr41.png)

You also need to copy the `DATABSE_URL` into our local setup file(`.env` file). Copy the Database URL and paste it into the `.env` file as follows:

```env
DATABASE_URL=postgres://sjxgipufegmgsw:78cbb568e@ec2-52-4-104-184.compute-1.amazonaws.com:5432/dbmuget
```
The format for the postgres URL is as follows:

```
postgresql://[user[:password]@][netloc][:port][/dbname]
```

We have now created environment variables for our django application and also added config vars in the heroku app, we now need a way to parse these environment variables into the Django project.  

### Parsing Environment variables using python-dotenv

We will use [python-dotenv](https://pypi.org/project/python-dotenv/) to parse variables into the django settings configurations like `SECRET_KEY` and `DATABASES`. 

- Install `python-dotenv` with pip with the command :  
```
pip install python-dotenv
```
We need to then modify the default variables in the `settings.py` file. Firstly, we will load in the `.env` file for accessing the environment variables for the configuration of the project.

Append the following code, to the top of the `settings.py` file, make sure don't override the configuration so remove unnecessary imports and configurations.

``` python
# <project_name>/settings.py

import os
from dotenv import load_dotenv

BASE_DIR = Path(__file__).resolve().parent.parent

load_dotenv(os.path.join(BASE_DIR, ".env"))

```

We have imported the package `dotenv` basically the `python-dotenv` into the `settings.py` file and the module `os` for loading the `.env` file. The `load_dotenv` function helps in loading the `key-value` pairs which are the configuration variables that can act as actual environment variables. We provide in a file to the [load_dotenv](https://saurabh-kumar.com/python-dotenv/) function which is the `.env` file in our case, you can pick up any location for the `.env` file but make sure to change the location while parsing the file into the `load_dotenv` function. 

After loading the variables into the `settings.py` file, we now need to access those variables and set the appropriate variables the configuration from the variables received from the `load_dotenv` function. The `os.getenv` function to access the environment variables. The `os.getenv` function takes a parameter as the `key` for the environment variable and returns the value of the environment variable.

``` python
SECRET_KEY = os.getenv("SECRET_KEY")
```

We have secretly configured the `SECRET_KEY` for the django project. If you have any other variables as simple key-value pairs like `AUTH` passwords, username, etc. you can use this method to get the configuration variables. 

### Loading Database configuration

Databases are a bit different as compared to other simple configurations in django. We need to make a few adjustments to the default database configuration. We will install another package `dj-database-url` to configure `DATABASE_URL`. Since the databse_url has a few components we need a way to extract the details like the `hostname`, `port`, `database_name`, and `password`. Using the `dj-database-url` package we have a few functions that can serve the purpose.

```
pip install dj-database-url
```

At the end of your `settings.py` file, append the following code. 

``` python
import dj_database_url

DATABASE_URL = os.getenv("DATABASE_URL")

DATABASES = {
    "default": dj_database_url.config(default=DATABASE_URL, conn_max_age=1800),
}
```

We would need an adapter for making migrations to the `PostgreSQL` database i.e. the `psycopg2` package. This is a mandatory step if you are working with `postgres` database. This can be installed with the simple pip install:

```
pip install psycopg2

# If it does not work try
pip install psycopg2-binary


# if still error persists try installing setuptools
pip install -U setuptools
pip install psycopg2
```

Now, that we have configured the database, we can now apply migrations to the fresh database of postgres provided by heroku. We will simply run the migrate command and in the heroku app the PostgreSQL database would have been modified and an appropriate schema should be applied.

```
python manage.py migrate
```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652602284553/oTtGev28-.png)

Make sure to update the `requirements.txt` file before pushing the project to Heroku for making sure everything works as expected. Since we have installed a few additional packages that are directly used in the `settings.py` file, we need to run the `pip freeze` command to update the `requiremnets.txt` file.

## Serving Static Files

Now, if you have some static files like `CSS`, `Javascript`, or `images`, you need to configure the staticfiles in order to serve them from the heroku server. We will require another config variable for collecting the static files from the selected repository. 

```python

STATIC_URL = "static/"
STATICFILES_DIRS = [os.path.join(BASE_DIR, "static")]
STATIC_ROOT = os.path.join(BASE_DIR, "staticfiles")

``` 

Here, if you have served your static files from the `static` folder in the root directory of your django project, you can add the above code in the settings.py file. We will collect all static files in the folder along with the default static files provided by django in the `staticfiles` directory. Run the following command if you want to test whether the static files are properly collected and served.

```
python manage.py collectstatic 
```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652600828657/JgL4nLuiL.png)

So, this command will collect all the static files and store them in a single place. We see that the files from the admin section are also copied as well as the custom static files from the project configuration. Now, we can move on to set the config variable for the heroku app.

```
DISABLE_COLLECTSTATIC = 0
```

We can set the `DISABLE_COLLECTSTATIC` variable as `0` or `1` indicating whether to disable it or not. We have currently enabled the static file collection while deploying the app but you can set it to `1` to disable the collection of static files.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652613798420/mbqzf1Kqd.png)

Since I first tested the application on heroku, the static files don't work as expected, we need another package to make sure the staticfiles are served property. We will be installing the `whitenoise` package which serves as the middleware for serving the static files.

```
pip install whitenoise
```

Add the whitenoise middleware `whitenoise.middleware.WhiteNoiseMiddleware` to the `MIDDLEWARE` list in the `settings.py` file.

````python
MIDDLEWARE = [
...
...
...
    'whitenoise.middleware.WhiteNoiseMiddleware',
]

```

That should be enough to make the most of the deployment on heroku. You will have to make a few adjustments as per your requirements and project.

## Deploy from GitHub

We are now all set to deploy the application on Heroku, we can use the `Connect to GitHub` or `Heroku CLI` to push the code into production. Heroku CLI is quite easy with a few sets of commands but if your project is deployed on GitHub, you can straightaway let the deployment start the build on a push to a specific branch. This becomes quite automotive and easy to scale while deploying a large-scale application. 

```
pip freeze > requirements.txt
```

This step is quite important because you need to make sure that all the packages are listed in the `requirements.txt` file else you will have to wait for the build to fail and redeploy.

Make sure the server is running first on your local machine, remember the server will be set up from scratch but the database will already have applied migrations if you have applied migrations before after connecting the Heroku Postgres database.
 
```
python manage.py collectstatic

python manage.py runserver
```

This will set up the origin of the remote repository that will be pushing the project code. Next, make sure to commit the code which will contain all the required stuff for deploying the code.

Checklist for deploying the code

```
- requirements.txt
- Procfile
- runtime.txt
- django-project
- environment variables / config variables 
- static file configuration
- database configuration
- migrate schema of database 
- gitignore file for ignoring virtualenvs, .env file, staticfiles, etc
```

here's a sample `.gitignore` for a minimal django project.

```gitignore
.env/
.venv/
env/
venv/
*.env

*.pyc
db.sqlite3
staticfiles/
```

If you want a full-fledged `.gitignore` for a complex django project, you can take the reference from Jose Padilla's [gitignore Template](https://github.com/jpadilla/django-project-template/blob/master/.gitignore) for a django project.  

### Git Commit the Django Project
```
git status 

git add .

git commit -m "config for heroku deployment"
```
Carefully check the status of the git repository before committing and make sure you don't forget anything by mistake, it won't a big problem but it would mess up the build process.


![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652613991525/hxQgtGOoM.png)

After committing the code, we can now push the code to GitHub. We first need to set the remote repository reference to be able to push the code to it. 

```
git remote add origin https://github.com/<username>/<repo_name>
```
This will set up the `origin` as the remote repository on GitHub. Once the remote repository is created, we can move into the push of the git repository that will trigger the build. First, navigate to the `Deploy` section in the heroku app's dashboard where we want to connect the `GitHub` repository and allow the automatic deploy from a branch in this case we have chosen the `main` branch.

Due to some `Heroku` Internal Server Issues, the GitHub integration seems to have broken and isn't working as of May 2022, but it might work later. 

Below is a screenshot of my previous project deployed to Heroku using a GitHub repository.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652605497382/5VuQUQ0t0.png)

```
git push origin main
```

This will prompt you for your GitHub credentials and will deploy the commits to the remote repository on GitHub. This push on the main branch should also trigger the build process of the heroku app for this django project. You can navigate to the Activity section for the Build logs. 

If you have followed the article well, and your repository has all the correct configurations, the build will succeed, else chances are you might have missed a few things and the app might have crashed. You can debug your application build with the simple heroku CLI command:

```
heroku logs --tail -a <app_name>
```

This can be quite handy and saves a lot of time in understanding what went wrong in the build. It might be related to database migration, static files, python package not found, and some silly mistakes and errors that can be fixed after committing the code and pushing it to GitHub again.

If you do not want a GitHub repository, you can directly push the code from the local git repository to the remote heroku app center. This will require us the Heroku CLI.

## Heroku CLI

We can use the heroku CLI for pushing the code via the simple git repository. We can push the references via the branch and a remote repository on heroku to build our app.  For this, we assume you have heroku installed and logged in. We will require the django project code and heroku cli to build the django web application.

```bash
heroku git:remote -a <heroku_app_name>

# for my case
heroku git:remote -a blog-django-dep
```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652614221069/vCAKD0zsz.png)

After this, you can commit your code and the project as git repository. We have added the remote repository location on heroku, we can now simply push the code to the remote repository.

```
git push heroku main
```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652614125785/uEzFQ9VvQ.png)

Here, `heroku` is the remote repository location and `main` is the branch of the repository. This will push the code to the repository and thereafter create a build to deploy the Django project as a Heroku application.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652614381808/kYTmB3EO2p.png)

You can hit the `Open App` button on the top right corner and there should be your deployed Django application. 

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652610395538/xjUiODhoK.png)


## A Few Tricks and Gotchas

There are a few tricks and issues that you might pop into while deploying a django project on heroku, especially the build process. It requires a few iterations to get the complete app setup.

### Run command from the Dashboard console

If you don't have heroku CLI set up and want to fix a few things on the pushed code to the heroku app, you can use the `Run Console` option from the `More` Tab on the top right corner of theApplication dashboard. This is a great way to run migrations, configure static files or tweak a few things here and there without messing up the code on GitHub or the remote git repositories. 

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652614775294/lgDPwr2yr.png)


![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652614821950/uTzQVB8sC.png)


![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652614845269/BkZhu3SGH.png)

### Deploy with Docker 

You can even use the docker container to deploy a Django application on Heroku. It is a great way of learning a lot of deployment strategies and techniques using Docker. You'll get familiar with interesting concepts like virtualization, and containerization, and also learn Docker on the way. You can follow this tutorial on [Deploying Django applications with Docker on Heroku](https://testdriven.io/blog/deploying-django-to-heroku-with-docker/). Also, you can check out the official Heroku documentation for [deploying python applications](https://devcenter.heroku.com/articles/deploying-python).

### What are Dynos?

Dynos are simply web processes or workers that serve your web application. Dynos in Heroku are allocated based on the build process, once the slug is created a dyno is created as it runs on a VM container. This simply means there are limitations on how to use the web application and its sleep process. The hobby tier is sufficient for normal testing projects and side projects though you will have to pay and move into advance tiers to increase the dyno allocations and scaling of those web processes. 

It's not a simple thing to understand but to keep it simple, it might be a container to process the client's request and serve the page for a finite duration of the interaction. Also, your application will sleep after half an hour, if you try to reload the application every half an hour it will consume your resource allocation for the month and this is how the tiers and divided for paid services on Heroku. You can check out the detail over [here](https://www.heroku.com/pricing#containers).

## Conclusion

So, that is one of the ways we can deploy a Django application on Heroku with the PostgreSQL database. You can find the [django-blog project] on [GitHub] for following along with the deployment process.  In the next few parts of the series, we will be hopefully covering other platforms where you can deploy a Django o application.

Hopefully, you liked the above tutorial, if you have any questions. feedback, or queries, you can contact me on the Social handles provided below. Thank you for reading and till the next post Happy Coding :)