---
article_html: "<h2>Introduction</h2>\n<p>After getting familiar with the folder structure
  of the Django framework, we'll create our first view in an app. The basics of creating
  and mapping a view with a URL will be cleared by the end of this part.</p>\n<h2>Creating
  Views</h2>\n<blockquote>\n<p>Views are the functions written in python as a logic
  control unit of the webserver</p>\n</blockquote>\n<p>To create a view or typically-like
  function, we need to write a function in the <code>views.py</code> file inside of
  the application folder. The function name can be anything but should be a sensible
  name as far as its usability is concerned. Let's take a basic example of sending
  an HTTP response of &quot;Hello World&quot;.</p>\n<h4>project_name/app_name/views.py</h4>\n<pre
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
  class=\"kn\">from</span> <span class=\"nn\">django.http</span> <span class=\"kn\">import</span>
  <span class=\"n\">HttpResponse</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">index</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \ <span class=\"k\">return</span> <span class=\"n\">HttpResponse</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;Hello World&quot;</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>Yes,
  we are simply returning an HTTP Response right now, but rendering Templates/HTML
  Documents is quite similar and easy to grasp in Django. So, this is a view or a
  piece of logic but there is a piece missing in this. Where should this function
  be used? Of course a URL i.e a path to a web server.</p>\n<p>We'll see how to map
  the views to an URL in Django in the next section</p>\n<h2>Mapping the Views to
  a URL</h2>\n<p>We need to first create a <code>urls.py</code> file in the application
  folder to create a map of the URL to be mapped with the view. After creating the
  file in the same app folder as the <code>views.py</code>, import the function in
  the view into the file.</p>\n<h4>project_name/app_name/urls.py</h4>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">from</span> <span class=\"nn\">.views</span> <span class=\"kn\">import</span>
  <span class=\"n\">index</span>\n<span class=\"kn\">from</span> <span class=\"nn\">django.urls</span>
  <span class=\"kn\">import</span> <span class=\"n\">path</span> \n\n<span class=\"n\">urlpatterns</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">index</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s2\">&quot;index&quot;</span><span class=\"p\">),</span>\n<span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>The path can be anything you
  like but for simplicity, we'll keep it blank('') for now.</p>\n<p>Now, you have
  the path for your view to work but it's not linked to the main project. We need
  to link the app urls to the project urls.</p>\n<p>To link the urls of your app to
  the main project folder, you need to just add a single line of code in the <code>urls.py</code>
  file of the project folder.</p>\n<p>In projectname folder -&gt; <a href=\"http://urls.py\">urls.py</a></p>\n<h4>project_name/urls.py</h4>\n<pre
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
  <span class=\"n\">include</span>\n\n<span class=\"n\">urlpatterns</span> <span class=\"o\">=</span>
  <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;admin/&#39;</span><span class=\"p\">,</span> <span class=\"n\">admin</span><span
  class=\"o\">.</span><span class=\"n\">site</span><span class=\"o\">.</span><span
  class=\"n\">urls</span><span class=\"p\">),</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">include</span><span class=\"p\">(</span><span class=\"s1\">&#39;post.urls&#39;</span><span
  class=\"p\">)),</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>You
  need to add the line <code>path('', include('post.urls')),</code> and also import
  the <code>include</code> function from <code>django.urls</code>. This additional
  statement includes the urls or all the <code>urlpatterns</code> in the <code>post</code>
  app from the <code>urls.py</code> file into the project's url-routes.</p>\n<p>Here,
  the URL path can be anything like <code>'home/'</code>, <code>'about/'</code>, <code>'posts/'</code>,
  etc. but since we are just understanding the basics, we'll keep it <code>''</code>
  i.e. the root URL.</p>\n<p>You can also see that there is another route in our project
  <code>'admin/'</code> which is a path to the admin section. We'll explore this path
  and the entire Admin Section in some other part of this series.</p>\n<p>Now if you
  start the server and visit the default URL i.e <code>http://127.0.0.1:8000</code>,
  you will see a simple HTTP message <code>Hello World</code>.</p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1638194390/blogmedia/uv1_xf4byq.png\"
  alt=\"Hello World view\" /></p>\n<h2>Breaking the <code>path</code> function in
  urlpatterns</h2>\n<p>The path function in the urlpatterns takes in at least 2 parameters,
  i.e. the URL pattern and the view of any other function that can be related to the
  webserver.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n\n<div class=\"highlight\"><pre><span></span>path( &#39; &#39;,   view,
  \   name )\n       ^       ^        ^ \n       |       |        |\n       |       |
  \    url_name\n       |   function_name\n   url_path    \n</pre></div>\n\n</pre>\n\n<h3>URL
  path</h3>\n<p>The URL Path is the pattern or literally the path which you use in
  the Browser's search bar. This can be static i.e. some hard-coded text like <code>home/</code>,
  <code>user/</code>, <code>post/home/</code>, etc. and we can also have dynamic URLs
  like <code>post/&lt;pk:id&gt;/</code>, <code>user/&lt;str:name&gt;/</code>, etc.
  here the characters <code>&lt;pk:id&gt;</code> and <code>&lt;str:name&gt;</code>
  will be replaced by the actual id(integer/primary key) or the name(String) itself.</p>\n<p>This
  is used in an actual web application, where there might be a user profile that needs
  the unique user-id to render it specifically for that user. The User-Profile is
  just an example, it can anything like posts, emails, products, any form of a content-driven
  application.</p>\n<h3>View</h3>\n<p>The view or the function is the name of the
  function that will be attached to that URL path. That means once the user visits
  that URL, the function is literally called. <strong>View is just a fancy word for
  a function(or any logic basically).</strong> There is a lot to be covered when it
  comes to <code>View</code> as there are a lot of ways to create it, there are two
  types of views, how to use them for various use-cases that can be learned along
  the way because it is a topic where the crust of Django exists.</p>\n<p>We'll learn
  to create different implementations and structure our views, for time-being just
  consider them as the unit where every operation on the web can be performed. We
  can create other standalone functions in python to work with the views to make it
  a bit structured and readable.</p>\n<h3>URL Name</h3>\n<p>This is an optional parameter
  to the path function as we do not mandatorily need to give the URL map a name. This
  can be really useful in multi-page application websites where you need to link one
  page to another and that becomes a lot easier with the URL name. We do not need
  this right now, we'll touch it when we'll see the Django Templating Language.</p>\n<h2>Example
  Views</h2>\n<p>Let's create some examples to understand the working of Views and
  URLs. We'll create a dynamic URL and integrate the Python module in the views to
  get familiarized with the concept.</p>\n<h3>Dynamic URLs</h3>\n<p>We can use the
  dynamic URLs or placeholder variables to render out the content dynamically. Let's
  create another set of View and URL map.</p>\n<h4>project_name/app_name/views.py</h4>\n<pre
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
  class=\"k\">def</span> <span class=\"nf\">greet</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"p\">):</span>\n    <span class=\"k\">return</span> <span class=\"n\">HttpResponse</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;Welcome, &quot;</span> <span class=\"o\">+</span>
  <span class=\"n\">name</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>This
  view or function takes an additional argument called <code>name</code> and in response,
  it just says <code>Welcome, name</code> where the name can be any string. Now after
  creating the view, we need to map the view to a URL pattern, We'll add a path for
  this greet function.</p>\n<h4>project_name/app_name/urls.py</h4>\n<pre class='wrapper'>\n\n<div
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
  class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;greet/&lt;str:name&gt;/&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">greet</span><span class=\"p\">,</span> <span
  class=\"n\">name</span><span class=\"o\">=</span><span class=\"s2\">&quot;greet&quot;</span><span
  class=\"p\">),</span>\n</pre></div>\n\n</pre>\n\n<p>You can see how we have created
  the url-pattern here. The greet part is static but the <code>&lt;str:name&gt;</code>
  is a variable or just a URL parameter to be passed to the view as the value of the
  variable <code>name</code>. We have also given the URL map a name called greet,
  just for demonstration of its creation.</p>\n<p>You'll get an error, 100% if you
  are blindly following me! Didn't you forget something?</p>\n<p>Import the greet
  function from the views like so:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"kn\">from</span> <span class=\"nn\">.views</span> <span class=\"kn\">import</span>
  <span class=\"n\">index</span><span class=\"p\">,</span> <span class=\"n\">greet</span>
  \ \n</pre></div>\n\n</pre>\n\n<p>So, after we visit the URL <code>https://127.0.0.1:8000/greet/harry</code>,
  you should see a response <code>Welcome, harry</code> as simple as that.</p>\n<p><img
  src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1638252762/blogmedia/uv-greet_e2wg5o.gif\"
  alt=\"Greet URL Demo\" /></p>\n<p>Now, how is this working? We see the view first.
  The function takes two parameters one is most common the request which stores the
  meta-data about the request, the other parameter is the name that we will be use
  to respond to the server dynamically. The name variable is used in the string with
  the HttpResponse function to return a simple string.</p>\n<p>Then, in the URLs,
  we need to find a way to pass the variable name to the view, for that we use the
  <code>&lt;string:name&gt;</code> which is like a URL parameter to the view. The
  path function automatically parses the name to the appropriate view and hence we
  call the greet function with the name variable from the URL.</p>\n<h3>Using Pythonic
  things</h3>\n<p>We'll use some Python libraries or functions in the Django App.
  In this way, we'll see it's nearly no-brainer to use Python functions or libraries
  in the Django framework as indeed all files which we are working with are Python
  files.</p>\n<h4>project_name/app_name/views.py</h4>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">from</span> <span class=\"nn\">random</span> <span class=\"kn\">import</span>
  <span class=\"n\">randint</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">dice</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"n\">number</span> <span class=\"o\">=</span> <span class=\"n\">randint</span><span
  class=\"p\">(</span><span class=\"mi\">1</span><span class=\"p\">,</span><span class=\"mi\">6</span><span
  class=\"p\">)</span>\n    <span class=\"k\">return</span> <span class=\"n\">HttpResponse</span><span
  class=\"p\">(</span><span class=\"sa\">f</span><span class=\"s2\">&quot;It&#39;s
  </span><span class=\"si\">{</span><span class=\"n\">number</span><span class=\"si\">}</span><span
  class=\"s2\">&quot;</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>This
  view is using the random module, you can pretty much use other web-compatible modules
  or libraries. We have used the <code>random.randint</code> function to generate
  the pseudo-random number between 1 and 6. We have used the f-string (<code>f&quot;{variable}&quot;</code>)styled
  Response string as int is not compatible with the response concatenation. So this
  is the logic of our map, now we'll need to link it to a URL-path.</p>\n<h4>project_name/app_name/urls.py</h4>\n<pre
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
  class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;throw/&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">dice</span><span class=\"p\">,</span> <span
  class=\"n\">name</span><span class=\"o\">=</span><span class=\"s2\">&quot;dice&quot;</span><span
  class=\"p\">),</span>\n</pre></div>\n\n</pre>\n\n<p>Also, import the view name from
  the views as <code>from .views import dice</code> also add other views if present.
  Now if we go to the URL <code>https://127.0.0.1:8000/throw/</code>, we shall see
  a random number in the response. This is how we used Python to make the logic of
  our view.</p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1638252951/blogmedia/uv-dice_bsodzq.gif\"
  alt=\"Dice URL Demo\" /></p>\n<p>So, that was the basics of creating and mapping
  views and urls. It is the most fundamental of the workflow in Django project development.
  You need to get familiar with the process of mapping Views and urls before diving
  into Templates, Models, and other complex stuff.</p>\n<h2>Conclusion</h2>\n<p>From
  this part of the series, we touched upon the basics of views and URLs. The concept
  of mapping URLs and views might have been much cleared and it will be even gripping
  after we explore the Template handling and Static files in the next part. If you
  have any queries or mistakes I might have made please let me know. Thanks for reading
  and Happy Coding :)</p>\n"
cover: ''
date: 2021-11-30
datetime: 2021-11-30 00:00:00+00:00
description: 'After getting familiar with the folder structure of the Django framework,
  we Views are the functions written in python as a logic control unit of the webserver '
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-11-30-Django-Basics-P4.md
html: "<h2>Introduction</h2>\n<p>After getting familiar with the folder structure
  of the Django framework, we'll create our first view in an app. The basics of creating
  and mapping a view with a URL will be cleared by the end of this part.</p>\n<h2>Creating
  Views</h2>\n<blockquote>\n<p>Views are the functions written in python as a logic
  control unit of the webserver</p>\n</blockquote>\n<p>To create a view or typically-like
  function, we need to write a function in the <code>views.py</code> file inside of
  the application folder. The function name can be anything but should be a sensible
  name as far as its usability is concerned. Let's take a basic example of sending
  an HTTP response of &quot;Hello World&quot;.</p>\n<h4>project_name/app_name/views.py</h4>\n<pre
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
  class=\"kn\">from</span> <span class=\"nn\">django.http</span> <span class=\"kn\">import</span>
  <span class=\"n\">HttpResponse</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">index</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \ <span class=\"k\">return</span> <span class=\"n\">HttpResponse</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;Hello World&quot;</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>Yes,
  we are simply returning an HTTP Response right now, but rendering Templates/HTML
  Documents is quite similar and easy to grasp in Django. So, this is a view or a
  piece of logic but there is a piece missing in this. Where should this function
  be used? Of course a URL i.e a path to a web server.</p>\n<p>We'll see how to map
  the views to an URL in Django in the next section</p>\n<h2>Mapping the Views to
  a URL</h2>\n<p>We need to first create a <code>urls.py</code> file in the application
  folder to create a map of the URL to be mapped with the view. After creating the
  file in the same app folder as the <code>views.py</code>, import the function in
  the view into the file.</p>\n<h4>project_name/app_name/urls.py</h4>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">from</span> <span class=\"nn\">.views</span> <span class=\"kn\">import</span>
  <span class=\"n\">index</span>\n<span class=\"kn\">from</span> <span class=\"nn\">django.urls</span>
  <span class=\"kn\">import</span> <span class=\"n\">path</span> \n\n<span class=\"n\">urlpatterns</span>
  <span class=\"o\">=</span> <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">index</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"o\">=</span><span class=\"s2\">&quot;index&quot;</span><span class=\"p\">),</span>\n<span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>The path can be anything you
  like but for simplicity, we'll keep it blank('') for now.</p>\n<p>Now, you have
  the path for your view to work but it's not linked to the main project. We need
  to link the app urls to the project urls.</p>\n<p>To link the urls of your app to
  the main project folder, you need to just add a single line of code in the <code>urls.py</code>
  file of the project folder.</p>\n<p>In projectname folder -&gt; <a href=\"http://urls.py\">urls.py</a></p>\n<h4>project_name/urls.py</h4>\n<pre
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
  <span class=\"n\">include</span>\n\n<span class=\"n\">urlpatterns</span> <span class=\"o\">=</span>
  <span class=\"p\">[</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;admin/&#39;</span><span class=\"p\">,</span> <span class=\"n\">admin</span><span
  class=\"o\">.</span><span class=\"n\">site</span><span class=\"o\">.</span><span
  class=\"n\">urls</span><span class=\"p\">),</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">include</span><span class=\"p\">(</span><span class=\"s1\">&#39;post.urls&#39;</span><span
  class=\"p\">)),</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>You
  need to add the line <code>path('', include('post.urls')),</code> and also import
  the <code>include</code> function from <code>django.urls</code>. This additional
  statement includes the urls or all the <code>urlpatterns</code> in the <code>post</code>
  app from the <code>urls.py</code> file into the project's url-routes.</p>\n<p>Here,
  the URL path can be anything like <code>'home/'</code>, <code>'about/'</code>, <code>'posts/'</code>,
  etc. but since we are just understanding the basics, we'll keep it <code>''</code>
  i.e. the root URL.</p>\n<p>You can also see that there is another route in our project
  <code>'admin/'</code> which is a path to the admin section. We'll explore this path
  and the entire Admin Section in some other part of this series.</p>\n<p>Now if you
  start the server and visit the default URL i.e <code>http://127.0.0.1:8000</code>,
  you will see a simple HTTP message <code>Hello World</code>.</p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1638194390/blogmedia/uv1_xf4byq.png\"
  alt=\"Hello World view\" /></p>\n<h2>Breaking the <code>path</code> function in
  urlpatterns</h2>\n<p>The path function in the urlpatterns takes in at least 2 parameters,
  i.e. the URL pattern and the view of any other function that can be related to the
  webserver.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n\n<div class=\"highlight\"><pre><span></span>path( &#39; &#39;,   view,
  \   name )\n       ^       ^        ^ \n       |       |        |\n       |       |
  \    url_name\n       |   function_name\n   url_path    \n</pre></div>\n\n</pre>\n\n<h3>URL
  path</h3>\n<p>The URL Path is the pattern or literally the path which you use in
  the Browser's search bar. This can be static i.e. some hard-coded text like <code>home/</code>,
  <code>user/</code>, <code>post/home/</code>, etc. and we can also have dynamic URLs
  like <code>post/&lt;pk:id&gt;/</code>, <code>user/&lt;str:name&gt;/</code>, etc.
  here the characters <code>&lt;pk:id&gt;</code> and <code>&lt;str:name&gt;</code>
  will be replaced by the actual id(integer/primary key) or the name(String) itself.</p>\n<p>This
  is used in an actual web application, where there might be a user profile that needs
  the unique user-id to render it specifically for that user. The User-Profile is
  just an example, it can anything like posts, emails, products, any form of a content-driven
  application.</p>\n<h3>View</h3>\n<p>The view or the function is the name of the
  function that will be attached to that URL path. That means once the user visits
  that URL, the function is literally called. <strong>View is just a fancy word for
  a function(or any logic basically).</strong> There is a lot to be covered when it
  comes to <code>View</code> as there are a lot of ways to create it, there are two
  types of views, how to use them for various use-cases that can be learned along
  the way because it is a topic where the crust of Django exists.</p>\n<p>We'll learn
  to create different implementations and structure our views, for time-being just
  consider them as the unit where every operation on the web can be performed. We
  can create other standalone functions in python to work with the views to make it
  a bit structured and readable.</p>\n<h3>URL Name</h3>\n<p>This is an optional parameter
  to the path function as we do not mandatorily need to give the URL map a name. This
  can be really useful in multi-page application websites where you need to link one
  page to another and that becomes a lot easier with the URL name. We do not need
  this right now, we'll touch it when we'll see the Django Templating Language.</p>\n<h2>Example
  Views</h2>\n<p>Let's create some examples to understand the working of Views and
  URLs. We'll create a dynamic URL and integrate the Python module in the views to
  get familiarized with the concept.</p>\n<h3>Dynamic URLs</h3>\n<p>We can use the
  dynamic URLs or placeholder variables to render out the content dynamically. Let's
  create another set of View and URL map.</p>\n<h4>project_name/app_name/views.py</h4>\n<pre
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
  class=\"k\">def</span> <span class=\"nf\">greet</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"n\">name</span><span
  class=\"p\">):</span>\n    <span class=\"k\">return</span> <span class=\"n\">HttpResponse</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;Welcome, &quot;</span> <span class=\"o\">+</span>
  <span class=\"n\">name</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>This
  view or function takes an additional argument called <code>name</code> and in response,
  it just says <code>Welcome, name</code> where the name can be any string. Now after
  creating the view, we need to map the view to a URL pattern, We'll add a path for
  this greet function.</p>\n<h4>project_name/app_name/urls.py</h4>\n<pre class='wrapper'>\n\n<div
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
  class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;greet/&lt;str:name&gt;/&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">greet</span><span class=\"p\">,</span> <span
  class=\"n\">name</span><span class=\"o\">=</span><span class=\"s2\">&quot;greet&quot;</span><span
  class=\"p\">),</span>\n</pre></div>\n\n</pre>\n\n<p>You can see how we have created
  the url-pattern here. The greet part is static but the <code>&lt;str:name&gt;</code>
  is a variable or just a URL parameter to be passed to the view as the value of the
  variable <code>name</code>. We have also given the URL map a name called greet,
  just for demonstration of its creation.</p>\n<p>You'll get an error, 100% if you
  are blindly following me! Didn't you forget something?</p>\n<p>Import the greet
  function from the views like so:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"kn\">from</span> <span class=\"nn\">.views</span> <span class=\"kn\">import</span>
  <span class=\"n\">index</span><span class=\"p\">,</span> <span class=\"n\">greet</span>
  \ \n</pre></div>\n\n</pre>\n\n<p>So, after we visit the URL <code>https://127.0.0.1:8000/greet/harry</code>,
  you should see a response <code>Welcome, harry</code> as simple as that.</p>\n<p><img
  src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1638252762/blogmedia/uv-greet_e2wg5o.gif\"
  alt=\"Greet URL Demo\" /></p>\n<p>Now, how is this working? We see the view first.
  The function takes two parameters one is most common the request which stores the
  meta-data about the request, the other parameter is the name that we will be use
  to respond to the server dynamically. The name variable is used in the string with
  the HttpResponse function to return a simple string.</p>\n<p>Then, in the URLs,
  we need to find a way to pass the variable name to the view, for that we use the
  <code>&lt;string:name&gt;</code> which is like a URL parameter to the view. The
  path function automatically parses the name to the appropriate view and hence we
  call the greet function with the name variable from the URL.</p>\n<h3>Using Pythonic
  things</h3>\n<p>We'll use some Python libraries or functions in the Django App.
  In this way, we'll see it's nearly no-brainer to use Python functions or libraries
  in the Django framework as indeed all files which we are working with are Python
  files.</p>\n<h4>project_name/app_name/views.py</h4>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">from</span> <span class=\"nn\">random</span> <span class=\"kn\">import</span>
  <span class=\"n\">randint</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">dice</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"n\">number</span> <span class=\"o\">=</span> <span class=\"n\">randint</span><span
  class=\"p\">(</span><span class=\"mi\">1</span><span class=\"p\">,</span><span class=\"mi\">6</span><span
  class=\"p\">)</span>\n    <span class=\"k\">return</span> <span class=\"n\">HttpResponse</span><span
  class=\"p\">(</span><span class=\"sa\">f</span><span class=\"s2\">&quot;It&#39;s
  </span><span class=\"si\">{</span><span class=\"n\">number</span><span class=\"si\">}</span><span
  class=\"s2\">&quot;</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>This
  view is using the random module, you can pretty much use other web-compatible modules
  or libraries. We have used the <code>random.randint</code> function to generate
  the pseudo-random number between 1 and 6. We have used the f-string (<code>f&quot;{variable}&quot;</code>)styled
  Response string as int is not compatible with the response concatenation. So this
  is the logic of our map, now we'll need to link it to a URL-path.</p>\n<h4>project_name/app_name/urls.py</h4>\n<pre
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
  class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;throw/&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">dice</span><span class=\"p\">,</span> <span
  class=\"n\">name</span><span class=\"o\">=</span><span class=\"s2\">&quot;dice&quot;</span><span
  class=\"p\">),</span>\n</pre></div>\n\n</pre>\n\n<p>Also, import the view name from
  the views as <code>from .views import dice</code> also add other views if present.
  Now if we go to the URL <code>https://127.0.0.1:8000/throw/</code>, we shall see
  a random number in the response. This is how we used Python to make the logic of
  our view.</p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1638252951/blogmedia/uv-dice_bsodzq.gif\"
  alt=\"Dice URL Demo\" /></p>\n<p>So, that was the basics of creating and mapping
  views and urls. It is the most fundamental of the workflow in Django project development.
  You need to get familiar with the process of mapping Views and urls before diving
  into Templates, Models, and other complex stuff.</p>\n<h2>Conclusion</h2>\n<p>From
  this part of the series, we touched upon the basics of views and URLs. The concept
  of mapping URLs and views might have been much cleared and it will be even gripping
  after we explore the Template handling and Static files in the next part. If you
  have any queries or mistakes I might have made please let me know. Thanks for reading
  and Happy Coding :)</p>\n"
image_url: https://res.cloudinary.com/dgpxbrwoz/image/upload/v1638253939/blogmedia/dj-uv_esbld2.png
long_description: After getting familiar with the folder structure of the Django framework,
  we Views are the functions written in python as a logic control unit of the webserver
  To create a view or typically-like function, we need to write a function in the  Yes,
  we a
now: 2025-05-01 18:11:33.314382
path: blog/posts/2021-11-30-Django-Basics-P4.md
prevnext: null
series:
- Django-Basics
- Django-Series
slug: django-basics-views-urls
status: published
subtitle: Creating and Understanding working of views and urls in Django
tags:
- django
- python
- web-development
templateKey: blog-post
title: 'Django Basics: Views and URLS'
today: 2025-05-01
---

## Introduction

After getting familiar with the folder structure of the Django framework, we'll create our first view in an app. The basics of creating and mapping a view with a URL will be cleared by the end of this part.

## Creating Views

> Views are the functions written in python as a logic control unit of the webserver

To create a view or typically-like function, we need to write a function in the `views.py` file inside of the application folder. The function name can be anything but should be a sensible name as far as its usability is concerned. Let's take a basic example of sending an HTTP response of "Hello World".

#### project_name/app_name/views.py
```python
from django.http import HttpResponse

def index(request):
  return HttpResponse("Hello World")
```  

Yes, we are simply returning an HTTP Response right now, but rendering Templates/HTML Documents is quite similar and easy to grasp in Django. So, this is a view or a piece of logic but there is a piece missing in this. Where should this function be used? Of course a URL i.e a path to a web server.

We'll see how to map the views to an URL in Django in the next section

## Mapping the Views to a URL

We need to first create a `urls.py` file in the application folder to create a map of the URL to be mapped with the view. After creating the file in the same app folder as the `views.py`, import the function in the view into the file.

#### project_name/app_name/urls.py
```python
from .views import index
from django.urls import path 

urlpatterns = [
    path('', index, name="index"),
]
```
The path can be anything you like but for simplicity, we'll keep it blank('') for now.   

Now, you have the path for your view to work but it's not linked to the main project. We need to link the app urls to the project urls. 

To link the urls of your app to the main project folder, you need to just add a single line of code in the `urls.py` file of the project folder.

In projectname folder -> urls.py

#### project_name/urls.py
```python
from django.contrib import admin
from django.urls import path, include

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', include('post.urls')),
]
```

You need to add the line `path('', include('post.urls')),` and also import the `include` function from `django.urls`. This additional statement includes the urls or all the `urlpatterns` in the `post` app from the `urls.py` file into the project's url-routes. 

Here, the URL path can be anything like `'home/'`, `'about/'`, `'posts/'`, etc. but since we are just understanding the basics, we'll keep it `''` i.e. the root URL. 

You can also see that there is another route in our project `'admin/'` which is a path to the admin section. We'll explore this path and the entire Admin Section in some other part of this series.
   
Now if you start the server and visit the default URL i.e `http://127.0.0.1:8000`, you will see a simple HTTP message `Hello World`.

![Hello World view](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1638194390/blogmedia/uv1_xf4byq.png)

## Breaking the `path` function in urlpatterns

The path function in the urlpatterns takes in at least 2 parameters, i.e. the URL pattern and the view of any other function that can be related to the webserver. 

```
path( ' ',   view,    name )
       ^       ^        ^ 
       |       |        |
       |       |     url_name
       |   function_name
   url_path    
```   

### URL path

The URL Path is the pattern or literally the path which you use in the Browser's search bar. This can be static i.e. some hard-coded text like `home/`, `user/`, `post/home/`, etc. and we can also have dynamic URLs like `post/<pk:id>/`, `user/<str:name>/`, etc. here the characters `<pk:id>` and `<str:name>` will be replaced by the actual id(integer/primary key) or the name(String) itself. 

This is used in an actual web application, where there might be a user profile that needs the unique user-id to render it specifically for that user. The User-Profile is just an example, it can anything like posts, emails, products, any form of a content-driven application. 

### View

The view or the function is the name of the function that will be attached to that URL path. That means once the user visits that URL, the function is literally called. **View is just a fancy word for a function(or any logic basically).** There is a lot to be covered when it comes to `View` as there are a lot of ways to create it, there are two types of views, how to use them for various use-cases that can be learned along the way because it is a topic where the crust of Django exists.  

We'll learn to create different implementations and structure our views, for time-being just consider them as the unit where every operation on the web can be performed. We can create other standalone functions in python to work with the views to make it a bit structured and readable.

### URL Name

This is an optional parameter to the path function as we do not mandatorily need to give the URL map a name. This can be really useful in multi-page application websites where you need to link one page to another and that becomes a lot easier with the URL name. We do not need this right now, we'll touch it when we'll see the Django Templating Language. 

## Example Views

Let's create some examples to understand the working of Views and URLs. We'll create a dynamic URL and integrate the Python module in the views to get familiarized with the concept.

### Dynamic URLs 

We can use the dynamic URLs or placeholder variables to render out the content dynamically. Let's create another set of View and URL map.

#### project_name/app_name/views.py
```python
def greet(request, name):
    return HttpResponse("Welcome, " + name)
```

This view or function takes an additional argument called `name` and in response, it just says `Welcome, name` where the name can be any string. Now after creating the view, we need to map the view to a URL pattern, We'll add a path for this greet function. 

#### project_name/app_name/urls.py
```python
path('greet/<str:name>/', greet, name="greet"),
```

You can see how we have created the url-pattern here. The greet part is static but the `<str:name>` is a variable or just a URL parameter to be passed to the view as the value of the variable `name`. We have also given the URL map a name called greet, just for demonstration of its creation. 

You'll get an error, 100% if you are blindly following me! Didn't you forget something?

Import the greet function from the views like so:

```python
from .views import index, greet  
```

So, after we visit the URL `https://127.0.0.1:8000/greet/harry`, you should see a response `Welcome, harry` as simple as that. 

![Greet URL Demo](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1638252762/blogmedia/uv-greet_e2wg5o.gif)

Now, how is this working? We see the view first. The function takes two parameters one is most common the request which stores the meta-data about the request, the other parameter is the name that we will be use to respond to the server dynamically. The name variable is used in the string with the HttpResponse function to return a simple string.

Then, in the URLs, we need to find a way to pass the variable name to the view, for that we use the `<string:name>` which is like a URL parameter to the view. The path function automatically parses the name to the appropriate view and hence we call the greet function with the name variable from the URL.

### Using Pythonic things

We'll use some Python libraries or functions in the Django App. In this way, we'll see it's nearly no-brainer to use Python functions or libraries in the Django framework as indeed all files which we are working with are Python files.

#### project_name/app_name/views.py
```python
from random import randint

def dice(request):
    number = randint(1,6)
    return HttpResponse(f"It's {number}")
```

This view is using the random module, you can pretty much use other web-compatible modules or libraries. We have used the `random.randint` function to generate the pseudo-random number between 1 and 6. We have used the f-string (`f"{variable}"`)styled Response string as int is not compatible with the response concatenation. So this is the logic of our map, now we'll need to link it to a URL-path. 

#### project_name/app_name/urls.py
```python
path('throw/', dice, name="dice"),
```

Also, import the view name from the views as `from .views import dice` also add other views if present. Now if we go to the URL `https://127.0.0.1:8000/throw/`, we shall see a random number in the response. This is how we used Python to make the logic of our view.

![Dice URL Demo](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1638252951/blogmedia/uv-dice_bsodzq.gif)

So, that was the basics of creating and mapping views and urls. It is the most fundamental of the workflow in Django project development. You need to get familiar with the process of mapping Views and urls before diving into Templates, Models, and other complex stuff. 

## Conclusion

From this part of the series, we touched upon the basics of views and URLs. The concept of mapping URLs and views might have been much cleared and it will be even gripping after we explore the Template handling and Static files in the next part. If you have any queries or mistakes I might have made please let me know. Thanks for reading and Happy Coding :)