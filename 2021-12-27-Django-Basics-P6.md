---
article_html: "<h2>Introduction</h2>\n<p>After creating templates, it should be rather
  tempting to add some styles and logic to them. Well yes, we'll see how to add static
  files in a web application using django. Static files are not only CSS, but also
  media/images and Javascript files as well. In this part of the series, we'll cover
  the basics of working with static files in django including the configuration, rendering
  and storing of the static files.</p>\n<h2>What are Static Files?</h2>\n<p>Static
  files as the name suggests are the files that don't change, your style sheets(css/scss)
  are not gonna change for every request from the client side, though the template
  might be dynamic. Also your logo, images in the design will not change unless you
  re-design it XD So these are the static files that needs to be rendered along with
  the templates.</p>\n<p>We have basically 3 types of static files, CSS, Javascript
  files and media files/static templates,etc. They are all rendered in the same way
  but as per their conventions and usage.</p>\n<p>You can learn about the theoretical
  information on <a href=\"https://docs.djangoproject.com/en/4.0/howto/static-files/\">static
  files</a> from the django documentation.</p>\n<h2>How to configure Static Files</h2>\n<p>Firstly
  you can create a folder for all the static files in the root folder. Usually the
  convention is <code>static</code> as the name of the folder. So, if you have created
  the template folder in the root directory, similar to that static folder can be
  created in that path.</p>\n<p>Next after creating the static folder in the project
  root folder, we need to configure the <code>settings.py</code> file to actually
  tell Django web server to look for all our static files in that folder. To do that,
  go to the <code>settings.py</code> file, now by this time you would have known where
  the <code>settings.py</code> file is (inside the project-named folder). Add the
  following at the end of the <code>settings.py</code> file.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># import os</span>\n<span class=\"c1\"># STATIC_URL = &#39;/static/&#39;</span>\n\n<span
  class=\"n\">STATICFILES_DIRS</span> <span class=\"o\">=</span> <span class=\"p\">(</span>\n
  \   <span class=\"n\">os</span><span class=\"o\">.</span><span class=\"n\">path</span><span
  class=\"o\">.</span><span class=\"n\">join</span><span class=\"p\">(</span><span
  class=\"n\">BASE_DIR</span><span class=\"p\">,</span> <span class=\"s2\">&quot;static/&quot;</span><span
  class=\"p\">),</span>\n<span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>Ignore
  the <code>import os</code> if you already have imported and the <code>STATIC_URL</code>
  if already there in the file. The <code>STATICFILES_DIRS</code> is the configuration
  that we tell the django environment to look for all our static files in the base/root
  directory of the project where the <code>static/</code> folder is. The <code>os.path.join()</code>
  actually gets the path of the directory in our operating system to the folder specified
  in the case of our project the <code>BASE_DIR</code> is the path of the project
  and we add in the static folder to actually the project path. The final piece and
  the crucial one is the <code>&quot;static/&quot;</code> path, this can be other
  location where you have created your static folder within the project.</p>\n<p>That's
  it! Yes, it's that simple. We can now create static files and render them in our
  templates.</p>\n<h2>Creating and Storing Static files</h2>\n<p>Now this part is
  customizable and it depends on your preference, how you want to organize the static
  folder. The convention that I follow is creating separate folders namely for <code>css</code>,
  <code>js</code> and <code>assets</code>(or <code>img</code>) mostly. And inside
  of this folders you can store the respective static files. This also creates the
  project more scalable in terms of it's maintenance.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>static\\\n  |__css\\\n  |__js\\\n
  \ |__assets\\\n</pre></div>\n\n</pre>\n\n<p>Let's create a static file and an image
  to demonstrate the concept of static files in django.</p>\n<ul>\n<li>css/style.css</li>\n</ul>\n<pre
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
  \       \n<div class='language-bar'>css</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nt\">body</span><span class=\"w\"> </span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">background-color</span><span
  class=\"p\">:</span><span class=\"mh\">#1d1dff</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">color</span><span
  class=\"p\">:</span><span class=\"kc\">white</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nt\">h1</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">text-align</span><span
  class=\"p\">:</span><span class=\"kc\">center</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">font-family</span><span class=\"o\">:</span><span
  class=\"w\"> </span><span class=\"kc\">monospace</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nt\">p</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">color</span><span
  class=\"p\">:</span><span class=\"mh\">#ff6600</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">font-weight</span><span
  class=\"p\">:</span><span class=\"mi\">500</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nt\">ul</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">list-style-type</span><span
  class=\"p\">:</span><span class=\"kc\">square</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<ul>\n<li>assets/tbicon.png</li>\n</ul>\n<p>Demo
  Image (that's my blog icon)</p>\n<p><img src=\"https://github.com/Mr-Destructive/techstructive-blog/blob/gh-pages/assets/img/tbicon.png?raw=true\"
  alt=\"Demo image\" /></p>\n<h2>Rendering Static Files from Templates</h2>\n<p>So,
  after configuring and creating the static files, we now can inject them into our
  templates. If you try to do the traditional way i.e. linking stylesheets/images/script
  files with HTML, it just won't work as you expect to and there's no point in using
  traditional way while creating a web application with a framework. So, there is
  a framework specific way to do things which make it easier and efficient for the
  project.</p>\n<p>To render any static file, we need to load the static tag which
  allows us to embed links for the static files into the templates. This means if
  the static files are not loaded directly instead in production(deploying our application)
  the static files are stored in a folder <code>STATIC_ROOT</code> which the server
  then loads, we'll see how that internally works when we get to deployment techniques
  for Django project.</p>\n<p>To load the static files from our configuration, we
  can simpy include the tag on top of the template.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>{% load static %}\n</pre></div>\n\n</pre>\n\n<p>The
  above templating tag will load the <code>static</code> tag which allows us to embed
  the links to the static files as explained earlier.</p>\n<p>Now, we can actually
  access any file with the static folder in our templates with a particular syntax
  as below:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"p\">&lt;</span><span class=\"nt\">link</span> <span class=\"na\">rel</span><span
  class=\"o\">=</span><span class=\"s\">&quot;stylesheet&quot;</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% static &#39;css/style.css&#39; %}&quot;</span><span
  class=\"p\">&gt;</span>  \n</pre></div>\n\n</pre>\n\n<p>Its just a example how to
  load the file, we are calling the static tag which we have loaded in previously
  and from there we are referencing the css file. The compact syntax would be : <code>{%
  static  'path-to-file'  %}</code></p>\n<p><strong>NOTE: The path to the static file
  is relative from the Static folder, i.e. enter the path of the file considering
  the static folder as the base directory.</strong></p>\n<h3>Demonstration of the
  static file</h3>\n<p>Let's render the static file which we created earlier i.e.
  the css file and the image into a template.</p>\n<p>Assuming you have a app called
  <code>post</code> in your django project, you can render static files as below:</p>\n<h1>templates/home.html</h1>\n<pre
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
  class=\"cp\">&lt;!DOCTYPE html&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">html</span>
  <span class=\"na\">lang</span><span class=\"o\">=</span><span class=\"s\">&quot;en&quot;</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">meta</span>
  <span class=\"na\">charset</span><span class=\"o\">=</span><span class=\"s\">&quot;UTF-8&quot;</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>Django Blog<span class=\"p\">&lt;/</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>\n    {% load static %}\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">link</span> <span class=\"na\">rel</span><span class=\"o\">=</span><span
  class=\"s\">&quot;stylesheet&quot;</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% static &#39;css/style.css&#39; %}&quot;</span><span class=\"p\">&gt;</span>
  \ \n<span class=\"p\">&lt;/</span><span class=\"nt\">head</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;</span><span class=\"nt\">body</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;</span><span class=\"nt\">h1</span><span class=\"p\">&gt;</span>Hello,
  World!<span class=\"p\">&lt;/</span><span class=\"nt\">h1</span><span class=\"p\">&gt;</span>\n
  \   {% block body %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">p</span><span
  class=\"p\">&gt;</span>This is not going to get inherited <span class=\"p\">&lt;/</span><span
  class=\"nt\">p</span><span class=\"p\">&gt;</span>\n    {% endblock %}\n    <span
  class=\"p\">&lt;</span><span class=\"nt\">p</span><span class=\"p\">&gt;</span>This
  will be inherited<span class=\"p\">&lt;/</span><span class=\"nt\">p</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>We are loading the static
  tag and then loading the css file using the tag syntax as explained above.</p>\n<h1>static/css/style.css</h1>\n<pre
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
  \       \n<div class='language-bar'>css</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nt\">body</span><span class=\"w\"> </span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">background-color</span><span
  class=\"p\">:</span><span class=\"mh\">#1d1dff</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">color</span><span
  class=\"p\">:</span><span class=\"kc\">white</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nt\">h1</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">text-align</span><span
  class=\"p\">:</span><span class=\"kc\">center</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">font-family</span><span class=\"o\">:</span><span
  class=\"w\"> </span><span class=\"kc\">monospace</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nt\">p</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">color</span><span
  class=\"p\">:</span><span class=\"mh\">#ff6600</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">font-weight</span><span
  class=\"p\">:</span><span class=\"mi\">500</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nt\">ul</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">list-style-type</span><span
  class=\"p\">:</span><span class=\"kc\">square</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>This
  is the static file,<code>style.css</code> stored inside the css folder of the static
  folder. This contains basic (very lame) CSS styling as we can understand.</p>\n<h1>post/views.py</h1>\n<pre
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
  <span class=\"n\">render</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">home</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;home.html&#39;</span><span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>The <code>views.py</code> file
  has the function that renders the template <code>home.html</code> from the templates
  folder inside the application specific folder.</p>\n<h1>post/urls.py</h1>\n<pre
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
  class=\"kn\">from</span> <span class=\"nn\">django.urls</span> <span class=\"kn\">import</span>
  <span class=\"n\">path</span>\n<span class=\"kn\">from</span> <span class=\"nn\">post</span>
  <span class=\"kn\">import</span> <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n        <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">home</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;home&quot;</span><span class=\"p\">),</span>\n        <span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>This
  is the application level configuration for the url routes to the views linking the
  views(functions) from the <code>views.py</code> file. The url in this file(code-snippet)
  is linking the root url('') to the home view in the <code>views.py</code> file.</p>\n<h1>Blog/urls.py</h1>\n<pre
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
  class=\"kn\">from</span> <span class=\"nn\">django.contrib</span> <span class=\"kn\">import</span>
  <span class=\"n\">admin</span>\n<span class=\"kn\">from</span> <span class=\"nn\">django.urls</span>
  <span class=\"kn\">import</span> <span class=\"n\">path</span><span class=\"p\">,</span>
  <span class=\"n\">include</span>\n<span class=\"kn\">from</span> <span class=\"nn\">django.views.generic</span>
  <span class=\"kn\">import</span> <span class=\"n\">TemplateView</span>\n\n<span
  class=\"n\">urlpatterns</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;admin/&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">admin</span><span class=\"o\">.</span><span
  class=\"n\">site</span><span class=\"o\">.</span><span class=\"n\">urls</span><span
  class=\"p\">),</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span> <span class=\"n\">include</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;post.urls&#39;</span><span class=\"p\">)),</span>\n<span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>The urls file in the project
  folder is the core configuration for project level url routes to individual applications
  within the project.</p>\n<p>Append the following if your templates and static files
  are not configured properly.</p>\n<h1>Blog/settings.py</h1>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">import</span> <span class=\"nn\">os</span>   \n\n<span class=\"n\">TEMPLATES</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"p\">{</span>\n
  \       <span class=\"s1\">&#39;BACKEND&#39;</span><span class=\"p\">:</span> <span
  class=\"s1\">&#39;django.template.backends.django.DjangoTemplates&#39;</span><span
  class=\"p\">,</span>\n        <span class=\"s1\">&#39;DIRS&#39;</span><span class=\"p\">:</span>
  <span class=\"p\">[</span><span class=\"n\">os</span><span class=\"o\">.</span><span
  class=\"n\">path</span><span class=\"o\">.</span><span class=\"n\">join</span><span
  class=\"p\">(</span><span class=\"n\">BASE_DIR</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;templates&#39;</span><span class=\"p\">),],</span>\n        <span
  class=\"s1\">&#39;APP_DIRS&#39;</span><span class=\"p\">:</span> <span class=\"kc\">True</span><span
  class=\"p\">,</span>\n        <span class=\"s1\">&#39;OPTIONS&#39;</span><span class=\"p\">:</span>
  <span class=\"p\">{</span>\n            <span class=\"s1\">&#39;context_processors&#39;</span><span
  class=\"p\">:</span> <span class=\"p\">[</span>\n                <span class=\"s1\">&#39;django.template.context_processors.debug&#39;</span><span
  class=\"p\">,</span>\n                <span class=\"s1\">&#39;django.template.context_processors.request&#39;</span><span
  class=\"p\">,</span>\n                <span class=\"s1\">&#39;django.contrib.auth.context_processors.auth&#39;</span><span
  class=\"p\">,</span>\n                <span class=\"s1\">&#39;django.contrib.messages.context_processors.messages&#39;</span><span
  class=\"p\">,</span>\n            <span class=\"p\">],</span>\n        <span class=\"p\">},</span>\n
  \   <span class=\"p\">},</span>\n<span class=\"p\">]</span>\n<span class=\"n\">STATIC_URL</span>
  <span class=\"o\">=</span> <span class=\"s1\">&#39;/static/&#39;</span>\n<span class=\"n\">STATICFILES_DIRS</span>
  <span class=\"o\">=</span> <span class=\"p\">(</span>\n    <span class=\"n\">os</span><span
  class=\"o\">.</span><span class=\"n\">path</span><span class=\"o\">.</span><span
  class=\"n\">join</span><span class=\"p\">(</span><span class=\"n\">BASE_DIR</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;static/&quot;</span><span class=\"p\">),</span>\n<span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>SO, the result of the above code
  is as simple template as shown in the picture below:</p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640621276/blogmedia/static-1_vu41gf.png\"
  alt=\"Static file demo\" /></p>\n<p>This will also work if you do it with traditional
  HTML syntax, but I'd explained why it's not recommended to do it while using frameworks.</p>\n<p>Let's
  see how static files are rendered in inherited templates. We'll tinker with the
  <code>for.html</code> template created in the <a href=\"https://mr-destructive.github.io/techstructive-blog/django/python/web-development/2021/12/14/Django-Basics-P5.html\">previous
  part</a>.</p>\n<h1>template/for.html</h1>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>django</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cp\">{%</span> <span class=\"k\">extends</span> <span class=\"s1\">&#39;home.html&#39;</span>
  <span class=\"cp\">%}</span><span class=\"x\"></span>\n<span class=\"cp\">{%</span>
  <span class=\"k\">load</span> <span class=\"nv\">static</span> <span class=\"cp\">%}</span><span
  class=\"x\"></span>\n\n<span class=\"cp\">{%</span> <span class=\"k\">block</span>
  <span class=\"nv\">body</span> <span class=\"cp\">%}</span><span class=\"x\"></span>\n<span
  class=\"x\">    &lt;img src=&quot;</span><span class=\"cp\">{%</span> <span class=\"k\">static</span>
  <span class=\"s1\">&#39;assets/tbicon.png&#39;</span> <span class=\"cp\">%}</span><span
  class=\"x\">&quot; height=&quot;50px&quot; width=&quot;50px&quot; /&gt;</span>\n<span
  class=\"x\">    &lt;ul&gt;</span>\n<span class=\"x\">        </span><span class=\"cp\">{%</span>
  <span class=\"k\">for</span> <span class=\"nv\">sport</span> <span class=\"k\">in</span>
  <span class=\"nv\">sport_list</span> <span class=\"cp\">%}</span><span class=\"x\"></span>\n<span
  class=\"x\">        &lt;li&gt;</span><span class=\"cp\">{{</span> <span class=\"nv\">sport</span>
  <span class=\"cp\">}}</span><span class=\"x\">&lt;/li&gt;</span>\n<span class=\"x\">
  \       </span><span class=\"cp\">{%</span> <span class=\"k\">endfor</span> <span
  class=\"cp\">%}</span><span class=\"x\"></span>\n<span class=\"x\">    &lt;/ul&gt;</span>\n<span
  class=\"cp\">{%</span> <span class=\"k\">endblock</span> <span class=\"cp\">%}</span><span
  class=\"x\"></span>\n</pre></div>\n\n</pre>\n\n<p>We will have re-load the static
  tag for each template only if we need to include a new static file in the template.
  So we use the <code>{% load static %}</code> again as we are loading the static
  file (image) in this template.</p>\n<h1>post/views.py</h1>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">from</span> <span class=\"nn\">django.shortcuts</span> <span class=\"kn\">import</span>
  <span class=\"n\">render</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">for_demo</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"n\">sports</span> <span class=\"o\">=</span> <span class=\"p\">(</span><span
  class=\"s1\">&#39;football&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;cricket&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;volleyball&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;hockey&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;basketball&#39;</span><span
  class=\"p\">)</span>\n    <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;for.html&#39;</span><span class=\"p\">,</span> <span class=\"p\">{</span><span
  class=\"s1\">&#39;sport_list&#39;</span><span class=\"p\">:</span> <span class=\"n\">sports</span><span
  class=\"p\">})</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">home</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;home.html&#39;</span><span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<h1>post/urls.py</h1>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">from</span> <span class=\"nn\">django.urls</span> <span class=\"kn\">import</span>
  <span class=\"n\">path</span>\n<span class=\"kn\">from</span> <span class=\"nn\">post</span>
  <span class=\"kn\">import</span> <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n        <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">home</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;home&quot;</span><span class=\"p\">),</span>\n        <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;for/&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">for_demo</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;fordemo&quot;</span><span class=\"p\">),</span>\n        <span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>So, that's the url and view map
  created, we can now be able to see the result in the <code>127.0.0.1:8000/for/</code>
  url to see the below result:</p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640622976/blogmedia/static-tempinh_peyjrg.png\"
  alt=\"Static demo for inheritance of tempaltes\" /></p>\n<p>The list style has been
  changed and thus we can see that the CSS from the parent template is also being
  inherited.</p>\n<p>Here is the django project structure which I have created with
  this series so far:</p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640624705/blogmedia/trr-static_bgt9du.png\"
  alt=\"Folder tree structure\" /></p>\n<p>So that has been it for the Static files
  in Django. Though there are lot of depth for rendering and loading the static files,
  we'll explore as we get our grasp in the django and web development terminologies.</p>\n<h2>Conclusion</h2>\n<p>So,
  from this article, we were able to configure and render static files like CSS/Images
  and optionally Javascript into the Django application. We covered from ground how
  to configure, load and structure the folder for storing all the static files at
  the project level.</p>\n<p>Hope you found it helpful and if you have any queries
  please let me know. We'll start with the databases probably from the next part in
  the Django Basics Series. Until then have a great week and as always Happy Coding
  :)</p>\n"
cover: ''
date: 2021-12-27
datetime: 2021-12-27 00:00:00+00:00
description: After creating templates, it should be rather tempting to add some styles
  and logic to them. Well yes, we Static files as the name suggests are the files
  that d
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-12-27-Django-Basics-P6.md
html: "<h2>Introduction</h2>\n<p>After creating templates, it should be rather tempting
  to add some styles and logic to them. Well yes, we'll see how to add static files
  in a web application using django. Static files are not only CSS, but also media/images
  and Javascript files as well. In this part of the series, we'll cover the basics
  of working with static files in django including the configuration, rendering and
  storing of the static files.</p>\n<h2>What are Static Files?</h2>\n<p>Static files
  as the name suggests are the files that don't change, your style sheets(css/scss)
  are not gonna change for every request from the client side, though the template
  might be dynamic. Also your logo, images in the design will not change unless you
  re-design it XD So these are the static files that needs to be rendered along with
  the templates.</p>\n<p>We have basically 3 types of static files, CSS, Javascript
  files and media files/static templates,etc. They are all rendered in the same way
  but as per their conventions and usage.</p>\n<p>You can learn about the theoretical
  information on <a href=\"https://docs.djangoproject.com/en/4.0/howto/static-files/\">static
  files</a> from the django documentation.</p>\n<h2>How to configure Static Files</h2>\n<p>Firstly
  you can create a folder for all the static files in the root folder. Usually the
  convention is <code>static</code> as the name of the folder. So, if you have created
  the template folder in the root directory, similar to that static folder can be
  created in that path.</p>\n<p>Next after creating the static folder in the project
  root folder, we need to configure the <code>settings.py</code> file to actually
  tell Django web server to look for all our static files in that folder. To do that,
  go to the <code>settings.py</code> file, now by this time you would have known where
  the <code>settings.py</code> file is (inside the project-named folder). Add the
  following at the end of the <code>settings.py</code> file.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># import os</span>\n<span class=\"c1\"># STATIC_URL = &#39;/static/&#39;</span>\n\n<span
  class=\"n\">STATICFILES_DIRS</span> <span class=\"o\">=</span> <span class=\"p\">(</span>\n
  \   <span class=\"n\">os</span><span class=\"o\">.</span><span class=\"n\">path</span><span
  class=\"o\">.</span><span class=\"n\">join</span><span class=\"p\">(</span><span
  class=\"n\">BASE_DIR</span><span class=\"p\">,</span> <span class=\"s2\">&quot;static/&quot;</span><span
  class=\"p\">),</span>\n<span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>Ignore
  the <code>import os</code> if you already have imported and the <code>STATIC_URL</code>
  if already there in the file. The <code>STATICFILES_DIRS</code> is the configuration
  that we tell the django environment to look for all our static files in the base/root
  directory of the project where the <code>static/</code> folder is. The <code>os.path.join()</code>
  actually gets the path of the directory in our operating system to the folder specified
  in the case of our project the <code>BASE_DIR</code> is the path of the project
  and we add in the static folder to actually the project path. The final piece and
  the crucial one is the <code>&quot;static/&quot;</code> path, this can be other
  location where you have created your static folder within the project.</p>\n<p>That's
  it! Yes, it's that simple. We can now create static files and render them in our
  templates.</p>\n<h2>Creating and Storing Static files</h2>\n<p>Now this part is
  customizable and it depends on your preference, how you want to organize the static
  folder. The convention that I follow is creating separate folders namely for <code>css</code>,
  <code>js</code> and <code>assets</code>(or <code>img</code>) mostly. And inside
  of this folders you can store the respective static files. This also creates the
  project more scalable in terms of it's maintenance.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>static\\\n  |__css\\\n  |__js\\\n
  \ |__assets\\\n</pre></div>\n\n</pre>\n\n<p>Let's create a static file and an image
  to demonstrate the concept of static files in django.</p>\n<ul>\n<li>css/style.css</li>\n</ul>\n<pre
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
  \       \n<div class='language-bar'>css</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nt\">body</span><span class=\"w\"> </span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">background-color</span><span
  class=\"p\">:</span><span class=\"mh\">#1d1dff</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">color</span><span
  class=\"p\">:</span><span class=\"kc\">white</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nt\">h1</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">text-align</span><span
  class=\"p\">:</span><span class=\"kc\">center</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">font-family</span><span class=\"o\">:</span><span
  class=\"w\"> </span><span class=\"kc\">monospace</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nt\">p</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">color</span><span
  class=\"p\">:</span><span class=\"mh\">#ff6600</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">font-weight</span><span
  class=\"p\">:</span><span class=\"mi\">500</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nt\">ul</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">list-style-type</span><span
  class=\"p\">:</span><span class=\"kc\">square</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<ul>\n<li>assets/tbicon.png</li>\n</ul>\n<p>Demo
  Image (that's my blog icon)</p>\n<p><img src=\"https://github.com/Mr-Destructive/techstructive-blog/blob/gh-pages/assets/img/tbicon.png?raw=true\"
  alt=\"Demo image\" /></p>\n<h2>Rendering Static Files from Templates</h2>\n<p>So,
  after configuring and creating the static files, we now can inject them into our
  templates. If you try to do the traditional way i.e. linking stylesheets/images/script
  files with HTML, it just won't work as you expect to and there's no point in using
  traditional way while creating a web application with a framework. So, there is
  a framework specific way to do things which make it easier and efficient for the
  project.</p>\n<p>To render any static file, we need to load the static tag which
  allows us to embed links for the static files into the templates. This means if
  the static files are not loaded directly instead in production(deploying our application)
  the static files are stored in a folder <code>STATIC_ROOT</code> which the server
  then loads, we'll see how that internally works when we get to deployment techniques
  for Django project.</p>\n<p>To load the static files from our configuration, we
  can simpy include the tag on top of the template.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>{% load static %}\n</pre></div>\n\n</pre>\n\n<p>The
  above templating tag will load the <code>static</code> tag which allows us to embed
  the links to the static files as explained earlier.</p>\n<p>Now, we can actually
  access any file with the static folder in our templates with a particular syntax
  as below:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"p\">&lt;</span><span class=\"nt\">link</span> <span class=\"na\">rel</span><span
  class=\"o\">=</span><span class=\"s\">&quot;stylesheet&quot;</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;{% static &#39;css/style.css&#39; %}&quot;</span><span
  class=\"p\">&gt;</span>  \n</pre></div>\n\n</pre>\n\n<p>Its just a example how to
  load the file, we are calling the static tag which we have loaded in previously
  and from there we are referencing the css file. The compact syntax would be : <code>{%
  static  'path-to-file'  %}</code></p>\n<p><strong>NOTE: The path to the static file
  is relative from the Static folder, i.e. enter the path of the file considering
  the static folder as the base directory.</strong></p>\n<h3>Demonstration of the
  static file</h3>\n<p>Let's render the static file which we created earlier i.e.
  the css file and the image into a template.</p>\n<p>Assuming you have a app called
  <code>post</code> in your django project, you can render static files as below:</p>\n<h1>templates/home.html</h1>\n<pre
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
  class=\"cp\">&lt;!DOCTYPE html&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">html</span>
  <span class=\"na\">lang</span><span class=\"o\">=</span><span class=\"s\">&quot;en&quot;</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">meta</span>
  <span class=\"na\">charset</span><span class=\"o\">=</span><span class=\"s\">&quot;UTF-8&quot;</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>Django Blog<span class=\"p\">&lt;/</span><span class=\"nt\">title</span><span
  class=\"p\">&gt;</span>\n    {% load static %}\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">link</span> <span class=\"na\">rel</span><span class=\"o\">=</span><span
  class=\"s\">&quot;stylesheet&quot;</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{% static &#39;css/style.css&#39; %}&quot;</span><span class=\"p\">&gt;</span>
  \ \n<span class=\"p\">&lt;/</span><span class=\"nt\">head</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;</span><span class=\"nt\">body</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;</span><span class=\"nt\">h1</span><span class=\"p\">&gt;</span>Hello,
  World!<span class=\"p\">&lt;/</span><span class=\"nt\">h1</span><span class=\"p\">&gt;</span>\n
  \   {% block body %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">p</span><span
  class=\"p\">&gt;</span>This is not going to get inherited <span class=\"p\">&lt;/</span><span
  class=\"nt\">p</span><span class=\"p\">&gt;</span>\n    {% endblock %}\n    <span
  class=\"p\">&lt;</span><span class=\"nt\">p</span><span class=\"p\">&gt;</span>This
  will be inherited<span class=\"p\">&lt;/</span><span class=\"nt\">p</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>We are loading the static
  tag and then loading the css file using the tag syntax as explained above.</p>\n<h1>static/css/style.css</h1>\n<pre
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
  \       \n<div class='language-bar'>css</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nt\">body</span><span class=\"w\"> </span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">background-color</span><span
  class=\"p\">:</span><span class=\"mh\">#1d1dff</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">color</span><span
  class=\"p\">:</span><span class=\"kc\">white</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nt\">h1</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">text-align</span><span
  class=\"p\">:</span><span class=\"kc\">center</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"k\">font-family</span><span class=\"o\">:</span><span
  class=\"w\"> </span><span class=\"kc\">monospace</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nt\">p</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">color</span><span
  class=\"p\">:</span><span class=\"mh\">#ff6600</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">font-weight</span><span
  class=\"p\">:</span><span class=\"mi\">500</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nt\">ul</span><span class=\"w\"></span>\n<span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"k\">list-style-type</span><span
  class=\"p\">:</span><span class=\"kc\">square</span><span class=\"p\">;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>This
  is the static file,<code>style.css</code> stored inside the css folder of the static
  folder. This contains basic (very lame) CSS styling as we can understand.</p>\n<h1>post/views.py</h1>\n<pre
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
  <span class=\"n\">render</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">home</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;home.html&#39;</span><span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>The <code>views.py</code> file
  has the function that renders the template <code>home.html</code> from the templates
  folder inside the application specific folder.</p>\n<h1>post/urls.py</h1>\n<pre
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
  class=\"kn\">from</span> <span class=\"nn\">django.urls</span> <span class=\"kn\">import</span>
  <span class=\"n\">path</span>\n<span class=\"kn\">from</span> <span class=\"nn\">post</span>
  <span class=\"kn\">import</span> <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n        <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">home</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;home&quot;</span><span class=\"p\">),</span>\n        <span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>This
  is the application level configuration for the url routes to the views linking the
  views(functions) from the <code>views.py</code> file. The url in this file(code-snippet)
  is linking the root url('') to the home view in the <code>views.py</code> file.</p>\n<h1>Blog/urls.py</h1>\n<pre
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
  class=\"kn\">from</span> <span class=\"nn\">django.contrib</span> <span class=\"kn\">import</span>
  <span class=\"n\">admin</span>\n<span class=\"kn\">from</span> <span class=\"nn\">django.urls</span>
  <span class=\"kn\">import</span> <span class=\"n\">path</span><span class=\"p\">,</span>
  <span class=\"n\">include</span>\n<span class=\"kn\">from</span> <span class=\"nn\">django.views.generic</span>
  <span class=\"kn\">import</span> <span class=\"n\">TemplateView</span>\n\n<span
  class=\"n\">urlpatterns</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;admin/&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">admin</span><span class=\"o\">.</span><span
  class=\"n\">site</span><span class=\"o\">.</span><span class=\"n\">urls</span><span
  class=\"p\">),</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span> <span class=\"n\">include</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;post.urls&#39;</span><span class=\"p\">)),</span>\n<span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>The urls file in the project
  folder is the core configuration for project level url routes to individual applications
  within the project.</p>\n<p>Append the following if your templates and static files
  are not configured properly.</p>\n<h1>Blog/settings.py</h1>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">import</span> <span class=\"nn\">os</span>   \n\n<span class=\"n\">TEMPLATES</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"p\">{</span>\n
  \       <span class=\"s1\">&#39;BACKEND&#39;</span><span class=\"p\">:</span> <span
  class=\"s1\">&#39;django.template.backends.django.DjangoTemplates&#39;</span><span
  class=\"p\">,</span>\n        <span class=\"s1\">&#39;DIRS&#39;</span><span class=\"p\">:</span>
  <span class=\"p\">[</span><span class=\"n\">os</span><span class=\"o\">.</span><span
  class=\"n\">path</span><span class=\"o\">.</span><span class=\"n\">join</span><span
  class=\"p\">(</span><span class=\"n\">BASE_DIR</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;templates&#39;</span><span class=\"p\">),],</span>\n        <span
  class=\"s1\">&#39;APP_DIRS&#39;</span><span class=\"p\">:</span> <span class=\"kc\">True</span><span
  class=\"p\">,</span>\n        <span class=\"s1\">&#39;OPTIONS&#39;</span><span class=\"p\">:</span>
  <span class=\"p\">{</span>\n            <span class=\"s1\">&#39;context_processors&#39;</span><span
  class=\"p\">:</span> <span class=\"p\">[</span>\n                <span class=\"s1\">&#39;django.template.context_processors.debug&#39;</span><span
  class=\"p\">,</span>\n                <span class=\"s1\">&#39;django.template.context_processors.request&#39;</span><span
  class=\"p\">,</span>\n                <span class=\"s1\">&#39;django.contrib.auth.context_processors.auth&#39;</span><span
  class=\"p\">,</span>\n                <span class=\"s1\">&#39;django.contrib.messages.context_processors.messages&#39;</span><span
  class=\"p\">,</span>\n            <span class=\"p\">],</span>\n        <span class=\"p\">},</span>\n
  \   <span class=\"p\">},</span>\n<span class=\"p\">]</span>\n<span class=\"n\">STATIC_URL</span>
  <span class=\"o\">=</span> <span class=\"s1\">&#39;/static/&#39;</span>\n<span class=\"n\">STATICFILES_DIRS</span>
  <span class=\"o\">=</span> <span class=\"p\">(</span>\n    <span class=\"n\">os</span><span
  class=\"o\">.</span><span class=\"n\">path</span><span class=\"o\">.</span><span
  class=\"n\">join</span><span class=\"p\">(</span><span class=\"n\">BASE_DIR</span><span
  class=\"p\">,</span> <span class=\"s2\">&quot;static/&quot;</span><span class=\"p\">),</span>\n<span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>SO, the result of the above code
  is as simple template as shown in the picture below:</p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640621276/blogmedia/static-1_vu41gf.png\"
  alt=\"Static file demo\" /></p>\n<p>This will also work if you do it with traditional
  HTML syntax, but I'd explained why it's not recommended to do it while using frameworks.</p>\n<p>Let's
  see how static files are rendered in inherited templates. We'll tinker with the
  <code>for.html</code> template created in the <a href=\"https://mr-destructive.github.io/techstructive-blog/django/python/web-development/2021/12/14/Django-Basics-P5.html\">previous
  part</a>.</p>\n<h1>template/for.html</h1>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>django</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"cp\">{%</span> <span class=\"k\">extends</span> <span class=\"s1\">&#39;home.html&#39;</span>
  <span class=\"cp\">%}</span><span class=\"x\"></span>\n<span class=\"cp\">{%</span>
  <span class=\"k\">load</span> <span class=\"nv\">static</span> <span class=\"cp\">%}</span><span
  class=\"x\"></span>\n\n<span class=\"cp\">{%</span> <span class=\"k\">block</span>
  <span class=\"nv\">body</span> <span class=\"cp\">%}</span><span class=\"x\"></span>\n<span
  class=\"x\">    &lt;img src=&quot;</span><span class=\"cp\">{%</span> <span class=\"k\">static</span>
  <span class=\"s1\">&#39;assets/tbicon.png&#39;</span> <span class=\"cp\">%}</span><span
  class=\"x\">&quot; height=&quot;50px&quot; width=&quot;50px&quot; /&gt;</span>\n<span
  class=\"x\">    &lt;ul&gt;</span>\n<span class=\"x\">        </span><span class=\"cp\">{%</span>
  <span class=\"k\">for</span> <span class=\"nv\">sport</span> <span class=\"k\">in</span>
  <span class=\"nv\">sport_list</span> <span class=\"cp\">%}</span><span class=\"x\"></span>\n<span
  class=\"x\">        &lt;li&gt;</span><span class=\"cp\">{{</span> <span class=\"nv\">sport</span>
  <span class=\"cp\">}}</span><span class=\"x\">&lt;/li&gt;</span>\n<span class=\"x\">
  \       </span><span class=\"cp\">{%</span> <span class=\"k\">endfor</span> <span
  class=\"cp\">%}</span><span class=\"x\"></span>\n<span class=\"x\">    &lt;/ul&gt;</span>\n<span
  class=\"cp\">{%</span> <span class=\"k\">endblock</span> <span class=\"cp\">%}</span><span
  class=\"x\"></span>\n</pre></div>\n\n</pre>\n\n<p>We will have re-load the static
  tag for each template only if we need to include a new static file in the template.
  So we use the <code>{% load static %}</code> again as we are loading the static
  file (image) in this template.</p>\n<h1>post/views.py</h1>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">from</span> <span class=\"nn\">django.shortcuts</span> <span class=\"kn\">import</span>
  <span class=\"n\">render</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">for_demo</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"n\">sports</span> <span class=\"o\">=</span> <span class=\"p\">(</span><span
  class=\"s1\">&#39;football&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;cricket&#39;</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;volleyball&#39;</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;hockey&#39;</span><span class=\"p\">,</span> <span class=\"s1\">&#39;basketball&#39;</span><span
  class=\"p\">)</span>\n    <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;for.html&#39;</span><span class=\"p\">,</span> <span class=\"p\">{</span><span
  class=\"s1\">&#39;sport_list&#39;</span><span class=\"p\">:</span> <span class=\"n\">sports</span><span
  class=\"p\">})</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">home</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;home.html&#39;</span><span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<h1>post/urls.py</h1>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">from</span> <span class=\"nn\">django.urls</span> <span class=\"kn\">import</span>
  <span class=\"n\">path</span>\n<span class=\"kn\">from</span> <span class=\"nn\">post</span>
  <span class=\"kn\">import</span> <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n        <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">home</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;home&quot;</span><span class=\"p\">),</span>\n        <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;for/&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">for_demo</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;fordemo&quot;</span><span class=\"p\">),</span>\n        <span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>So, that's the url and view map
  created, we can now be able to see the result in the <code>127.0.0.1:8000/for/</code>
  url to see the below result:</p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640622976/blogmedia/static-tempinh_peyjrg.png\"
  alt=\"Static demo for inheritance of tempaltes\" /></p>\n<p>The list style has been
  changed and thus we can see that the CSS from the parent template is also being
  inherited.</p>\n<p>Here is the django project structure which I have created with
  this series so far:</p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640624705/blogmedia/trr-static_bgt9du.png\"
  alt=\"Folder tree structure\" /></p>\n<p>So that has been it for the Static files
  in Django. Though there are lot of depth for rendering and loading the static files,
  we'll explore as we get our grasp in the django and web development terminologies.</p>\n<h2>Conclusion</h2>\n<p>So,
  from this article, we were able to configure and render static files like CSS/Images
  and optionally Javascript into the Django application. We covered from ground how
  to configure, load and structure the folder for storing all the static files at
  the project level.</p>\n<p>Hope you found it helpful and if you have any queries
  please let me know. We'll start with the databases probably from the next part in
  the Django Basics Series. Until then have a great week and as always Happy Coding
  :)</p>\n"
image_url: https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640610003/blogmedia/dj-static-6_pjipoj.png
long_description: After creating templates, it should be rather tempting to add some
  styles and logic to them. Well yes, we Static files as the name suggests are the
  files that don We have basically 3 types of static files, CSS, Javascript files
  and media files/static
now: 2025-05-01 18:11:33.314913
path: blog/posts/2021-12-27-Django-Basics-P6.md
prevnext: null
series:
- Django-Basics
- Django-Series
slug: django-basics-static-files
status: published
subtitle: Loading Static Files in a Django project/app
tags:
- django
- python
- web-development
templateKey: blog-post
title: 'Django Basics: Static Files'
today: 2025-05-01
---

## Introduction

After creating templates, it should be rather tempting to add some styles and logic to them. Well yes, we'll see how to add static files in a web application using django. Static files are not only CSS, but also media/images and Javascript files as well. In this part of the series, we'll cover the basics of working with static files in django including the configuration, rendering and storing of the static files. 

## What are Static Files?

Static files as the name suggests are the files that don't change, your style sheets(css/scss) are not gonna change for every request from the client side, though the template might be dynamic. Also your logo, images in the design will not change unless you re-design it XD So these are the static files that needs to be rendered along with the templates.

We have basically 3 types of static files, CSS, Javascript files and media files/static templates,etc. They are all rendered in the same way but as per their conventions and usage. 

You can learn about the theoretical information on [static files](https://docs.djangoproject.com/en/4.0/howto/static-files/) from the django documentation.

## How to configure Static Files

Firstly you can create a folder for all the static files in the root folder. Usually the convention is `static` as the name of the folder. So, if you have created the template folder in the root directory, similar to that static folder can be created in that path. 

Next after creating the static folder in the project root folder, we need to configure the `settings.py` file to actually tell Django web server to look for all our static files in that folder. To do that, go to the `settings.py` file, now by this time you would have known where the `settings.py` file is (inside the project-named folder). Add the following at the end of the `settings.py` file.

```python
# import os
# STATIC_URL = '/static/'

STATICFILES_DIRS = (
    os.path.join(BASE_DIR, "static/"),
)
```   

Ignore the `import os` if you already have imported and the `STATIC_URL` if already there in the file. The `STATICFILES_DIRS` is the configuration that we tell the django environment to look for all our static files in the base/root directory of the project where the `static/` folder is. The `os.path.join()` actually gets the path of the directory in our operating system to the folder specified in the case of our project the `BASE_DIR` is the path of the project and we add in the static folder to actually the project path. The final piece and the crucial one is the `"static/"` path, this can be other location where you have created your static folder within the project.

That's it! Yes, it's that simple. We can now create static files and render them in our templates. 

## Creating and Storing Static files

Now this part is customizable and it depends on your preference, how you want to organize the static folder. The convention that I follow is creating separate folders namely for `css`, `js` and `assets`(or `img`) mostly. And inside of this folders you can store the respective static files. This also creates the project more scalable in terms of it's maintenance. 

```
static\
  |__css\
  |__js\
  |__assets\
```

Let's create a static file and an image to demonstrate the concept of static files in django. 

- css/style.css

```css
body 
{
    background-color:#1d1dff;
    color:white;
}

h1
{
    text-align:center
    font-family: monospace;
}

p
{
    color:#ff6600;
    font-weight:500;
}

ul
{
    list-style-type:square;
}
```

- assets/tbicon.png 

Demo Image (that's my blog icon)   

![Demo image](https://github.com/Mr-Destructive/techstructive-blog/blob/gh-pages/assets/img/tbicon.png?raw=true)

## Rendering Static Files from Templates

So, after configuring and creating the static files, we now can inject them into our templates. If you try to do the traditional way i.e. linking stylesheets/images/script files with HTML, it just won't work as you expect to and there's no point in using traditional way while creating a web application with a framework. So, there is a framework specific way to do things which make it easier and efficient for the project. 

To render any static file, we need to load the static tag which allows us to embed links for the static files into the templates. This means if the static files are not loaded directly instead in production(deploying our application) the static files are stored in a folder `STATIC_ROOT` which the server then loads, we'll see how that internally works when we get to deployment techniques for Django project. 

To load the static files from our configuration, we can simpy include the tag on top of the template.

```
{% load static %}
```

The above templating tag will load the `static` tag which allows us to embed the links to the static files as explained earlier. 

Now, we can actually access any file with the static folder in our templates with a particular syntax as below:

```html
<link rel="stylesheet" href="{% static 'css/style.css' %}">  
```   
Its just a example how to load the file, we are calling the static tag which we have loaded in previously and from there we are referencing the css file. The compact syntax would be : `{% static  'path-to-file'  %}`      

**NOTE: The path to the static file is relative from the Static folder, i.e. enter the path of the file considering the static folder as the base directory.** 

### Demonstration of the static file

Let's render the static file which we created earlier i.e. the css file and the image into a template. 

Assuming you have a app called `post` in your django project, you can render static files as below:

# templates/home.html 
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Django Blog</title>
    {% load static %}
    <link rel="stylesheet" href="{% static 'css/style.css' %}">  
</head>
<body>
    <h1>Hello, World!</h1>
    {% block body %}
    <p>This is not going to get inherited </p>
    {% endblock %}
    <p>This will be inherited</p>
</body>
</html>
```   
We are loading the static tag and then loading the css file using the tag syntax as explained above.       

# static/css/style.css      
```css
body 
{
    background-color:#1d1dff;
    color:white;
}

h1
{
    text-align:center
    font-family: monospace;
}

p
{
    color:#ff6600;
    font-weight:500;
}

ul
{
    list-style-type:square;
}
```   
This is the static file,`style.css` stored inside the css folder of the static folder. This contains basic (very lame) CSS styling as we can understand.     

# post/views.py 

```python
from django.shortcuts import render

def home(request):
    return render(request, 'home.html')
```
The `views.py` file has the function that renders the template `home.html` from the templates folder inside the application specific folder.   

# post/urls.py   
```python
from django.urls import path
from post import views

urlpatterns = [
        path('', views.home, name="home"),
        ]
```   
This is the application level configuration for the url routes to the views linking the views(functions) from the `views.py` file. The url in this file(code-snippet) is linking the root url('') to the home view in the `views.py` file.

# Blog/urls.py
```python
from django.contrib import admin
from django.urls import path, include
from django.views.generic import TemplateView

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', include('post.urls')),
]
```
The urls file in the project folder is the core configuration for project level url routes to individual applications within the project.

Append the following if your templates and static files are not configured properly.

# Blog/settings.py
```python
import os   

TEMPLATES = [
    {
        'BACKEND': 'django.template.backends.django.DjangoTemplates',
        'DIRS': [os.path.join(BASE_DIR, 'templates'),],
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
STATIC_URL = '/static/'
STATICFILES_DIRS = (
    os.path.join(BASE_DIR, "static/"),
)
```
SO, the result of the above code is as simple template as shown in the picture below:

![Static file demo](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640621276/blogmedia/static-1_vu41gf.png)

This will also work if you do it with traditional HTML syntax, but I'd explained why it's not recommended to do it while using frameworks.

Let's see how static files are rendered in inherited templates. We'll tinker with the `for.html` template created in the [previous part](https://mr-destructive.github.io/techstructive-blog/django/python/web-development/2021/12/14/Django-Basics-P5.html).  

# template/for.html
```django
{% extends 'home.html' %}
{% load static %}

{% block body %}
    <img src="{% static 'assets/tbicon.png' %}" height="50px" width="50px" />
    <ul>
        {% for sport in sport_list %}
        <li>{{ sport }}</li>
        {% endfor %}
    </ul>
{% endblock %}
```
We will have re-load the static tag for each template only if we need to include a new static file in the template. So we use the `{% load static %}` again as we are loading the static file (image) in this template.

# post/views.py
```python
from django.shortcuts import render

def for_demo(request):
    sports = ('football', 'cricket', 'volleyball', 'hockey', 'basketball')
    return render(request, 'for.html', {'sport_list': sports})

def home(request):
    return render(request, 'home.html')
```

# post/urls.py
```python
from django.urls import path
from post import views

urlpatterns = [
        path('', views.home, name="home"),
        path('for/', views.for_demo, name="fordemo"),
        ]
```

So, that's the url and view map created, we can now be able to see the result in the `127.0.0.1:8000/for/` url to see the below result:

![Static demo for inheritance of tempaltes](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640622976/blogmedia/static-tempinh_peyjrg.png)

The list style has been changed and thus we can see that the CSS from the parent template is also being inherited. 

Here is the django project structure which I have created with this series so far:

![Folder tree structure](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640624705/blogmedia/trr-static_bgt9du.png)

So that has been it for the Static files in Django. Though there are lot of depth for rendering and loading the static files, we'll explore as we get our grasp in the django and web development terminologies. 

## Conclusion

So, from this article, we were able to configure and render static files like CSS/Images and optionally Javascript into the Django application. We covered from ground how to configure, load and structure the folder for storing all the static files at the project level. 

Hope you found it helpful and if you have any queries please let me know. We'll start with the databases probably from the next part in the Django Basics Series. Until then have a great week and as always Happy Coding :)