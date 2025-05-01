---
article_html: "<h2>Introduction</h2>\n<p>After learning the basics of views and URLs,
  we can now move on to the next concept i.e. Templates. In Django, Templates are
  quite an important component for the application as it acts as the <code>frontend</code>
  for the web application. With the help of templates and some features provided by
  Django, it becomes very intuitive and simple to make dynamic web content.\nIn this
  part, we understand what are templates and what is the way to render them.</p>\n<h2>What
  are Templates</h2>\n<p>Templates are simply a <code>html</code> document or kind
  of a wireframe for content to be displayed for the web app. Templates allow us to
  render some more relevant pieces of data rather than simple text HTTP responses
  as we did earlier. We can even re-use certain components of a template in other
  using the Django Templating Language (more on this later).</p>\n<p>So, using HTML
  templates, we can write a complete Webpage. If you are unfamiliar with HTML, you
  can check out the basics of HTML with this <a href=\"https://www.youtube.com/playlist?list=PL081AC329706B2953\">playlist</a>.</p>\n<p>Even
  If you are not familiar with HTML, this tutorial might be quite basic and not overwhelm
  you with all the tags.</p>\n<h2>Creating Templates</h2>\n<p>To create a Template,
  we can write a simple HTML document like the below:</p>\n<p>Create a folder <code>templates</code>
  in the base folder, inside the templates folder, create a file <code>index.html</code></p>\n<p><strong>templates\\index.html</strong></p>\n<pre
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
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>Hello, World!<span class=\"p\">&lt;/</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>This is a simple HTML template,
  with the <code>&lt;h1&gt;</code> tags. As Django is a framework, there is a standard
  for storing all the templates for the project and application. There are a couple
  of standard of options:</p>\n<ul>\n<li>One of which is creating a <code>templates</code>
  folder in the root folder as discussed earlier, also we need to modify the <code>project_name/settings.py</code>
  file.</li>\n</ul>\n<p>Inside the <code>settings.py</code> file, we need to locate
  the <code>TEMPLATES</code> section and modify as below:</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">import</span> <span class=\"nn\">os</span>\n\n<span class=\"n\">TEMPLATES</span>
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
  \   <span class=\"p\">},</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>In
  this snippet, we have changed the <code>DIRS</code> option to search the templates
  in the folder <code>templates</code> in the root directory.</p>\n<ul>\n<li>The other
  standard is creating a templates folder in each application.</li>\n</ul>\n<p>We
  can create the templates folder in each application instead of a single folder.</p>\n<h2>Rendering
  Templates</h2>\n<p>After creating a template and making the required settings to
  make sure Django is able to pick up those templates, we need to work with views
  and URLs to actually render those templates.</p>\n<p>There are a couple of ways
  to render templates in Django and some of them are discussed below:</p>\n<h3>Using
  TemplateView</h3>\n<p><a href=\"https://docs.djangoproject.com/en/3.2/ref/class-based-views/base/#django.views.generic.base.TemplateView\">TemplateView</a>
  is a class which is comes with <code>django.views.generic</code> library. This class
  allows us to render a template by providing in the name of the template, arguments
  or variables to be parsed, and so on.</p>\n<p>The simplest way to render a template
  is by the following way:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"kn\">from</span> <span class=\"nn\">django.contrib</span> <span class=\"kn\">import</span>
  <span class=\"n\">admin</span>\n<span class=\"kn\">from</span> <span class=\"nn\">django.urls</span>
  <span class=\"kn\">import</span> <span class=\"n\">path</span><span class=\"p\">,</span>
  <span class=\"n\">include</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">django.views.generic</span>
  <span class=\"kn\">import</span> <span class=\"n\">TemplateView</span>\n\n<span
  class=\"n\">urlpatterns</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;admin/&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">admin</span><span class=\"o\">.</span><span
  class=\"n\">site</span><span class=\"o\">.</span><span class=\"n\">urls</span><span
  class=\"p\">),</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span> <span class=\"n\">TemplateView</span><span
  class=\"o\">.</span><span class=\"n\">as_view</span><span class=\"p\">(</span><span
  class=\"n\">template_name</span><span class=\"o\">=</span><span class=\"s2\">&quot;index.html&quot;</span><span
  class=\"p\">),</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;index&quot;</span><span class=\"p\">),</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>We
  need to import the <code>TemplateView</code> from the <code>django.core.generic</code>
  so as to use the class for rendering the template.</p>\n<p>The <code>TemplateView</code>
  class takes in a couple of arguments, we'll use the <code>template_name</code> as
  an argument that takes in the name of the template. Here, we use the <code>index.html</code>
  as the template which we created earlier. We don't need to specify the entire path
  to the template as we make modifications in the <code>settings.py</code> file to
  pick the template from the mentioned directory. We use <code>as_view</code> function
  to load the class as a function/view.</p>\n<p>Activate the virtual environment for
  the proper functioning of the project.</p>\n<p>After activating the virtual environment
  we can run the server as follows:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>terminal</div>\n<div class=\"highlight\"><pre><span></span>python
  manage.py runserver\n</pre></div>\n\n</pre>\n\n<p>We can now see the following output
  and thus, we are now rendering a simple HTML template in Django.</p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639384994/blogmedia/templ1_vbwp5d.png\"
  alt=\"\" /></p>\n<h3>Using render</h3>\n<p>We can also use the <a href=\"https://docs.djangoproject.com/en/4.0/topics/http/shortcuts/#render\">render
  function</a> from <code>django.shortcuts</code> to simply render a template. But
  we will create a Python function or a View to actually render the template. So,
  we'll create a View-URL map as we created in the <a href=\"https://mr-destructive.github.io/techstructive-blog/django/python/web-development/2021/11/30/Django-Basics-P4.html\">previous
  part</a>.</p>\n<p>Firstly, let's create a view function in the <code>post/views.py</code>
  file, more generally (<code>app_name/views.py</code> file). Firstly, we need to
  import the render function from <code>django.shortcuts</code> and then return the
  function call of render.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  <span class=\"n\">render</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">home</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span><span class=\"s1\">&#39;index.html&#39;</span><span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>And in the URLs, we'll create
  a different pattern like for e.g. 'home/'</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  <span class=\"n\">path</span>\n<span class=\"kn\">from</span> <span class=\"nn\">post</span>
  <span class=\"kn\">import</span> <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span><span
  class=\"o\">=</span><span class=\"p\">[</span>\n        <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;home/&#39;</span><span class=\"p\">,</span><span
  class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">home</span><span
  class=\"p\">,</span><span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;home&quot;</span><span class=\"p\">),</span>\n        <span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  after creating the View-URL map and making sure the URL of the app is loaded in
  the project URLs, we can see the result as a simple HTML template.</p>\n<p><img
  src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639386932/blogmedia/templ2_rgoppj.png\"
  alt=\"\" /></p>\n<h2>Django Templating Language</h2>\n<p>The <a href=\"https://docs.djangoproject.com/en/3.2/ref/templates/language/\">Django
  Templating Language</a> is Django's way of making templates more dynamic and easy
  to write dynamic web applications. We'll take a brief look at what we can do with
  this type of Templating Language in Django.</p>\n<h3>Variables</h3>\n<p>This is
  the most common use case for the Django Templating Language/Engine as we can use
  the <a href=\"https://docs.djangoproject.com/en/3.2/ref/templates/language/#variables\">variables</a>
  from the Backend and inject it in the template. We can parse the variable into the
  template by the syntax : <code>{{ variable_name &quot; }}}}</code></p>\n<p>To show
  its use cases, we can declare a variable in a view and then parse it in the Template.
  Though it is not dynamic right now we can later on fetch values from the database
  and store them in the form of variables in our views.</p>\n<p><strong>templates/home.html</strong></p>\n<pre
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
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>Hello, {{ name }}<span class=\"p\">&lt;/</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p><strong>post/views.py</strong></p>\n<pre
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
  <span class=\"n\">render</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">variable_demo</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"n\">name</span> <span class=\"o\">=</span> <span class=\"s2\">&quot;Kevin&quot;</span>\n
  \   <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;home.html&#39;</span><span
  class=\"p\">,</span> <span class=\"p\">{</span><span class=\"s1\">&#39;name&#39;</span><span
  class=\"p\">:</span><span class=\"n\">name</span><span class=\"p\">})</span>\n    <span
  class=\"c1\">#The name can be anything, like a database query object, form detail,
  etc</span>\n</pre></div>\n\n</pre>\n\n<p>As we can see the variable in views is
  passed as a dictionary in python. The reference key along with a value of the variable
  as the name of the variable. We will use the key in the templates to parse the value
  of the variable.</p>\n<p><strong>post/urls.py</strong></p>\n<pre class='wrapper'>\n\n<div
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
  <span class=\"kn\">import</span> <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span><span
  class=\"o\">=</span><span class=\"p\">[</span>\n        <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;home/&#39;</span><span class=\"p\">,</span><span
  class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">home</span><span
  class=\"p\">,</span><span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;home&quot;</span><span class=\"p\">),</span>\n        <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;vardemo/&#39;</span><span class=\"p\">,</span><span
  class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">variable_demo</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;var&quot;</span><span class=\"p\">),</span>\n        <span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639389288/blogmedia/templ3_wvhagw.png\"
  alt=\"variable demo\" /></p>\n<p>As we can see, we were able to load the variable
  into the template using the Django Templating Engine.</p>\n<h3>Conditional statement</h3>\n<p>We
  can even use the conditional statement in the Template using a very simple syntax.
  We can use <code>{% if condition&quot;  }} %}</code> to use certain special kinds
  of blocks in the Template. We need to end those blocks as well using the syntax
  <code>{% endif  %}</code>, here <code>if</code> can be other blocks which we'll
  explore ahead.</p>\n<p>To create a basic if condition in the template, we can understand
  with the following example.</p>\n<p><strong>app_name/views.py</strong></p>\n<pre
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
  <span class=\"n\">render</span>\n<span class=\"kn\">from</span> <span class=\"nn\">random</span>
  <span class=\"kn\">import</span> <span class=\"n\">randint</span>\n\n<span class=\"k\">def</span>
  <span class=\"nf\">if_demo</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">):</span>\n   <span class=\"n\">number</span> <span class=\"o\">=</span>
  <span class=\"n\">randint</span><span class=\"p\">(</span><span class=\"mi\">1</span><span
  class=\"p\">,</span><span class=\"mi\">10</span><span class=\"p\">)</span>\n   <span
  class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;if_else.html&#39;</span><span
  class=\"p\">,</span> <span class=\"p\">{</span><span class=\"s1\">&#39;num&#39;</span><span
  class=\"p\">:</span><span class=\"n\">number</span><span class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p>Here,
  we have used the key name as <code>num</code> indicating we can give different names
  to the key which needs to be used in the template to render the values.</p>\n<p><strong>app_name/urls.py</strong></p>\n<pre
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
  class=\"p\">(</span><span class=\"s1\">&#39;if/&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">if_demo</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;ifdemo&quot;</span><span class=\"p\">),</span>\n        <span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p><strong>templates/if_else.html</strong></p>\n<pre
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
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n    {{ num }}\n    {% if num &gt; 5 %}\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>It&#39;s Greater then 5<span
  class=\"p\">&lt;/</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n
  \   {% elif num == 5 %}\n        <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span
  class=\"p\">&gt;</span>It&#39;s five!<span class=\"p\">&lt;/</span><span class=\"nt\">h2</span><span
  class=\"p\">&gt;</span>\n    {% else %}\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>It&#39;s less than 5<span class=\"p\">&lt;/</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n    {% endif %}\n<span class=\"p\">&lt;/</span><span
  class=\"nt\">body</span><span class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span
  class=\"nt\">html</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639411425/blogmedia/templ3_exj0fv.png\"
  alt=\"if-else demo\" /></p>\n<p>So, as we can see that, we can use the if-else conditions
  in the template and that is already powerful. This can be a bit messy as to handle
  mathematical operations or conditions into a single condition. This can really be
  used for really large datasets that can be shimmed down to really less coding and
  also improve readability.</p>\n<h3>For loop</h3>\n<p>Now, the most crucial component
  of the Django templating language is the loops. We can actually iterate over objects/lists
  in the template. This becomes a huge concept for actually making a dynamic web application.
  We n\\might want to iterate over all the entries in a database, or any other form
  of data which can make the app a lot dynamic and feel real-time.</p>\n<p>The syntax
  of for loop is almost similar to the if-else condition. We just replace the condition
  with the iterator and the list/object from the view context. <code>{% for i in list
  %}</code>, also end the for loop like <code>{% endfor %}</code>.</p>\n<p><strong>app_name/views.py</strong></p>\n<pre
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
  class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p>We have created a simple Python
  list called <code>sports</code> and we parse them to the template using a dictionary
  object, <code>sport_list</code> as the key for storing the value of the <code>sports</code>
  list.</p>\n<p><strong>app_name/urls.py</strong></p>\n<pre class='wrapper'>\n\n<div
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
  class=\"p\">(</span><span class=\"s1\">&#39;for/&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">for_demo</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;fordemo&quot;</span><span class=\"p\">),</span>\n        <span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p><strong>templates/for.html</strong></p>\n<pre
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
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">ul</span><span
  class=\"p\">&gt;</span>\n        {% for sport in sport_list %}\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">li</span><span class=\"p\">&gt;</span>{{ sport }}<span class=\"p\">&lt;/</span><span
  class=\"nt\">li</span><span class=\"p\">&gt;</span>\n        {% endfor %}\n    <span
  class=\"p\">&lt;/</span><span class=\"nt\">ul</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;/</span><span class=\"nt\">body</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;/</span><span class=\"nt\">html</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639475328/blogmedia/templ3_q8z8fr.png\"
  alt=\"for loop demo\" /></p>\n<p>We have used simple for loop in Pythonic syntax,
  we use an iterator in this case, <code>sport</code> acts as an iterator. We use
  this to store values one by one from the list <code>sport_list</code> which was
  earlier passed in the views as a key in the dictionary.</p>\n<p>Hence, this is quite
  scalable and used to fetch the objects/entries in the database and hence making
  it a lot easier to make a dynamic web application faster.</p>\n<h2>Template Inheritance</h2>\n<p>So,
  far we have seen that we need to create the base template again and again like all
  the basic HTML elements, title, and all the basic structure. But what if, we can
  reuse a specific template in another and extend the functionality of that template
  into a new one. This avoids the redundancy of writing the entire basic template
  or the layout of a web app over and over again.</p>\n<p>To do that, Django has the
  Template inheritance. We can use a template as its basic layout or a specific component
  in the web application. Again, similar to the for, if-else blocks the syntax for
  inheriting a template is quite similar.</p>\n<p>Take, for example, the home.html
  which consisted of only a <code>&lt;h1&gt;</code> tag in it. We can use this kind
  of template in other templates to really make it the home page. For that, we first
  need to enclose the template in a <code>block</code>, which is what allows us to
  use it in other templates.\nTo create a <code>block</code>, we simply need to write
  the following syntax before the component which we do not want in other templates:</p>\n<p><strong>templates/home.html</strong></p>\n<pre
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
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>Hello, World!<span class=\"p\">&lt;/</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>\n    {% block body %}\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">p</span><span class=\"p\">&gt;</span>This is not going to get inherited
  <span class=\"p\">&lt;/</span><span class=\"nt\">p</span><span class=\"p\">&gt;</span>\n
  \   {% endblock %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">p</span><span
  class=\"p\">&gt;</span>This will be inherited<span class=\"p\">&lt;/</span><span
  class=\"nt\">p</span><span class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span
  class=\"nt\">body</span><span class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span
  class=\"nt\">html</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>In
  this we have used the <code>blocks</code> with a name like <code>body</code> as
  <code>{% block body %}</code> this can be anything you like. We end the block with
  the similar syntax as the for/if blocks as <code>{% endblock %}</code>. Anything
  in between the blocks i.e <code>block block_name</code> and <code>endblock</code>
  is not inherited i.e it is unique to this template.</p>\n<p>We will see how we can
  use this template in other templates. We will actually extend this template and
  use the blocks to render the content in the template.</p>\n<p><strong>templates/if_else.html</strong></p>\n<pre
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>{%
  extends &#39;home.html&#39; %}\n{% block body %}\n    {{ num }}\n    {% if num &gt;
  5 %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>It&#39;s
  Greater then 5<span class=\"p\">&lt;/</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n
  \   {% elif num == 5 %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span
  class=\"p\">&gt;</span>It&#39;s five!<span class=\"p\">&lt;/</span><span class=\"nt\">h2</span><span
  class=\"p\">&gt;</span>\n    {% else %}\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>It&#39;s less than 5<span class=\"p\">&lt;/</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n    {% endif %}\n{% endblock
  %}\n</pre></div>\n\n</pre>\n\n<p>So, we first say to Django to extend the <code>home</code>
  template i.e. the Django will load the blocks from this template only, remember
  it will just load and not use the blocks until we explicitly tell it to.</p>\n<p>To
  use the blocks or kind of plug in the template content in the <code>if_else.html</code>
  or any other template, we need to again call the <code>blocks</code>. Here, we need
  to write the content inside the <code>blocks</code> to properly parse the blocks
  as this is an HTML template. The order of opening and closing elements do matter.\nSo,
  when we say <code>endblock</code> the last part of the base template is loaded i.e.
  the closing <code>body</code> and <code>html</code> tags. This is like plugging
  the template as it is before and after the block body.</p>\n<p><strong>app_name/views.py</strong></p>\n<pre
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
  class=\"p\">)</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">random</span>
  <span class=\"kn\">import</span> <span class=\"n\">randint</span>\n\n<span class=\"k\">def</span>
  <span class=\"nf\">if_demo</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">):</span>\n   <span class=\"n\">number</span> <span class=\"o\">=</span>
  <span class=\"n\">randint</span><span class=\"p\">(</span><span class=\"mi\">1</span><span
  class=\"p\">,</span><span class=\"mi\">10</span><span class=\"p\">)</span>\n   <span
  class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;if_else.html&#39;</span><span
  class=\"p\">,</span> <span class=\"p\">{</span><span class=\"s1\">&#39;num&#39;</span><span
  class=\"p\">:</span><span class=\"n\">number</span><span class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p><strong>app_name/urls.py</strong></p>\n<pre
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
  class=\"s2\">&quot;home&quot;</span><span class=\"p\">),</span>\n        <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;if/&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">if_demo</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;ifdemo&quot;</span><span class=\"p\">),</span>\n        <span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639479714/blogmedia/tempinher2_enisls.png\"
  alt=\"\" />\n<img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639477721/blogmedia/tempinher_lk0op0.png\"
  alt=\"template inheritance demo\" /></p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639479954/blogmedia/template-inh_lc8szo.gif\"
  alt=\"\" /></p>\n<p>The above gif illustrates the example in a neat way. The block
  is loaded from the given template as the extended template and hence it plugs the
  block into the frame of the template.</p>\n<h2>Conclusion</h2>\n<p>So, from this
  part of the series, we were able to understand the concept of Templates in Django,
  we were able to use variables, loops, conditional statements, and template inheriting
  in a Django application. In the next part, we'll try to touch up with the static
  files and see how to properly structure and configure them.</p>\n<p>Thank you for
  reading, if you didn't understand any of the examples, please let me know, I'll
  be happy to share the code. Happy Coding :)</p>\n"
cover: ''
date: 2021-12-14
datetime: 2021-12-14 00:00:00+00:00
description: After learning the basics of views and URLs, we can now move on to the
  next concept i.e. Templates. In Django, Templates are quite an important component
  for th
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-12-14-Django-Basics-P5.md
html: "<h2>Introduction</h2>\n<p>After learning the basics of views and URLs, we can
  now move on to the next concept i.e. Templates. In Django, Templates are quite an
  important component for the application as it acts as the <code>frontend</code>
  for the web application. With the help of templates and some features provided by
  Django, it becomes very intuitive and simple to make dynamic web content.\nIn this
  part, we understand what are templates and what is the way to render them.</p>\n<h2>What
  are Templates</h2>\n<p>Templates are simply a <code>html</code> document or kind
  of a wireframe for content to be displayed for the web app. Templates allow us to
  render some more relevant pieces of data rather than simple text HTTP responses
  as we did earlier. We can even re-use certain components of a template in other
  using the Django Templating Language (more on this later).</p>\n<p>So, using HTML
  templates, we can write a complete Webpage. If you are unfamiliar with HTML, you
  can check out the basics of HTML with this <a href=\"https://www.youtube.com/playlist?list=PL081AC329706B2953\">playlist</a>.</p>\n<p>Even
  If you are not familiar with HTML, this tutorial might be quite basic and not overwhelm
  you with all the tags.</p>\n<h2>Creating Templates</h2>\n<p>To create a Template,
  we can write a simple HTML document like the below:</p>\n<p>Create a folder <code>templates</code>
  in the base folder, inside the templates folder, create a file <code>index.html</code></p>\n<p><strong>templates\\index.html</strong></p>\n<pre
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
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>Hello, World!<span class=\"p\">&lt;/</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>This is a simple HTML template,
  with the <code>&lt;h1&gt;</code> tags. As Django is a framework, there is a standard
  for storing all the templates for the project and application. There are a couple
  of standard of options:</p>\n<ul>\n<li>One of which is creating a <code>templates</code>
  folder in the root folder as discussed earlier, also we need to modify the <code>project_name/settings.py</code>
  file.</li>\n</ul>\n<p>Inside the <code>settings.py</code> file, we need to locate
  the <code>TEMPLATES</code> section and modify as below:</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"kn\">import</span> <span class=\"nn\">os</span>\n\n<span class=\"n\">TEMPLATES</span>
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
  \   <span class=\"p\">},</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>In
  this snippet, we have changed the <code>DIRS</code> option to search the templates
  in the folder <code>templates</code> in the root directory.</p>\n<ul>\n<li>The other
  standard is creating a templates folder in each application.</li>\n</ul>\n<p>We
  can create the templates folder in each application instead of a single folder.</p>\n<h2>Rendering
  Templates</h2>\n<p>After creating a template and making the required settings to
  make sure Django is able to pick up those templates, we need to work with views
  and URLs to actually render those templates.</p>\n<p>There are a couple of ways
  to render templates in Django and some of them are discussed below:</p>\n<h3>Using
  TemplateView</h3>\n<p><a href=\"https://docs.djangoproject.com/en/3.2/ref/class-based-views/base/#django.views.generic.base.TemplateView\">TemplateView</a>
  is a class which is comes with <code>django.views.generic</code> library. This class
  allows us to render a template by providing in the name of the template, arguments
  or variables to be parsed, and so on.</p>\n<p>The simplest way to render a template
  is by the following way:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"kn\">from</span> <span class=\"nn\">django.contrib</span> <span class=\"kn\">import</span>
  <span class=\"n\">admin</span>\n<span class=\"kn\">from</span> <span class=\"nn\">django.urls</span>
  <span class=\"kn\">import</span> <span class=\"n\">path</span><span class=\"p\">,</span>
  <span class=\"n\">include</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">django.views.generic</span>
  <span class=\"kn\">import</span> <span class=\"n\">TemplateView</span>\n\n<span
  class=\"n\">urlpatterns</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s1\">&#39;admin/&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">admin</span><span class=\"o\">.</span><span
  class=\"n\">site</span><span class=\"o\">.</span><span class=\"n\">urls</span><span
  class=\"p\">),</span>\n    <span class=\"n\">path</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;&#39;</span><span class=\"p\">,</span> <span class=\"n\">TemplateView</span><span
  class=\"o\">.</span><span class=\"n\">as_view</span><span class=\"p\">(</span><span
  class=\"n\">template_name</span><span class=\"o\">=</span><span class=\"s2\">&quot;index.html&quot;</span><span
  class=\"p\">),</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;index&quot;</span><span class=\"p\">),</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>We
  need to import the <code>TemplateView</code> from the <code>django.core.generic</code>
  so as to use the class for rendering the template.</p>\n<p>The <code>TemplateView</code>
  class takes in a couple of arguments, we'll use the <code>template_name</code> as
  an argument that takes in the name of the template. Here, we use the <code>index.html</code>
  as the template which we created earlier. We don't need to specify the entire path
  to the template as we make modifications in the <code>settings.py</code> file to
  pick the template from the mentioned directory. We use <code>as_view</code> function
  to load the class as a function/view.</p>\n<p>Activate the virtual environment for
  the proper functioning of the project.</p>\n<p>After activating the virtual environment
  we can run the server as follows:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>terminal</div>\n<div class=\"highlight\"><pre><span></span>python
  manage.py runserver\n</pre></div>\n\n</pre>\n\n<p>We can now see the following output
  and thus, we are now rendering a simple HTML template in Django.</p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639384994/blogmedia/templ1_vbwp5d.png\"
  alt=\"\" /></p>\n<h3>Using render</h3>\n<p>We can also use the <a href=\"https://docs.djangoproject.com/en/4.0/topics/http/shortcuts/#render\">render
  function</a> from <code>django.shortcuts</code> to simply render a template. But
  we will create a Python function or a View to actually render the template. So,
  we'll create a View-URL map as we created in the <a href=\"https://mr-destructive.github.io/techstructive-blog/django/python/web-development/2021/11/30/Django-Basics-P4.html\">previous
  part</a>.</p>\n<p>Firstly, let's create a view function in the <code>post/views.py</code>
  file, more generally (<code>app_name/views.py</code> file). Firstly, we need to
  import the render function from <code>django.shortcuts</code> and then return the
  function call of render.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  <span class=\"n\">render</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">home</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span><span class=\"s1\">&#39;index.html&#39;</span><span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>And in the URLs, we'll create
  a different pattern like for e.g. 'home/'</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  <span class=\"n\">path</span>\n<span class=\"kn\">from</span> <span class=\"nn\">post</span>
  <span class=\"kn\">import</span> <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span><span
  class=\"o\">=</span><span class=\"p\">[</span>\n        <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;home/&#39;</span><span class=\"p\">,</span><span
  class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">home</span><span
  class=\"p\">,</span><span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;home&quot;</span><span class=\"p\">),</span>\n        <span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  after creating the View-URL map and making sure the URL of the app is loaded in
  the project URLs, we can see the result as a simple HTML template.</p>\n<p><img
  src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639386932/blogmedia/templ2_rgoppj.png\"
  alt=\"\" /></p>\n<h2>Django Templating Language</h2>\n<p>The <a href=\"https://docs.djangoproject.com/en/3.2/ref/templates/language/\">Django
  Templating Language</a> is Django's way of making templates more dynamic and easy
  to write dynamic web applications. We'll take a brief look at what we can do with
  this type of Templating Language in Django.</p>\n<h3>Variables</h3>\n<p>This is
  the most common use case for the Django Templating Language/Engine as we can use
  the <a href=\"https://docs.djangoproject.com/en/3.2/ref/templates/language/#variables\">variables</a>
  from the Backend and inject it in the template. We can parse the variable into the
  template by the syntax : <code>{{ variable_name &quot; }}}}</code></p>\n<p>To show
  its use cases, we can declare a variable in a view and then parse it in the Template.
  Though it is not dynamic right now we can later on fetch values from the database
  and store them in the form of variables in our views.</p>\n<p><strong>templates/home.html</strong></p>\n<pre
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
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>Hello, {{ name }}<span class=\"p\">&lt;/</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p><strong>post/views.py</strong></p>\n<pre
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
  <span class=\"n\">render</span>\n\n<span class=\"k\">def</span> <span class=\"nf\">variable_demo</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">):</span>\n
  \   <span class=\"n\">name</span> <span class=\"o\">=</span> <span class=\"s2\">&quot;Kevin&quot;</span>\n
  \   <span class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;home.html&#39;</span><span
  class=\"p\">,</span> <span class=\"p\">{</span><span class=\"s1\">&#39;name&#39;</span><span
  class=\"p\">:</span><span class=\"n\">name</span><span class=\"p\">})</span>\n    <span
  class=\"c1\">#The name can be anything, like a database query object, form detail,
  etc</span>\n</pre></div>\n\n</pre>\n\n<p>As we can see the variable in views is
  passed as a dictionary in python. The reference key along with a value of the variable
  as the name of the variable. We will use the key in the templates to parse the value
  of the variable.</p>\n<p><strong>post/urls.py</strong></p>\n<pre class='wrapper'>\n\n<div
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
  <span class=\"kn\">import</span> <span class=\"n\">views</span>\n\n<span class=\"n\">urlpatterns</span><span
  class=\"o\">=</span><span class=\"p\">[</span>\n        <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;home/&#39;</span><span class=\"p\">,</span><span
  class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">home</span><span
  class=\"p\">,</span><span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;home&quot;</span><span class=\"p\">),</span>\n        <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;vardemo/&#39;</span><span class=\"p\">,</span><span
  class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">variable_demo</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;var&quot;</span><span class=\"p\">),</span>\n        <span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639389288/blogmedia/templ3_wvhagw.png\"
  alt=\"variable demo\" /></p>\n<p>As we can see, we were able to load the variable
  into the template using the Django Templating Engine.</p>\n<h3>Conditional statement</h3>\n<p>We
  can even use the conditional statement in the Template using a very simple syntax.
  We can use <code>{% if condition&quot;  }} %}</code> to use certain special kinds
  of blocks in the Template. We need to end those blocks as well using the syntax
  <code>{% endif  %}</code>, here <code>if</code> can be other blocks which we'll
  explore ahead.</p>\n<p>To create a basic if condition in the template, we can understand
  with the following example.</p>\n<p><strong>app_name/views.py</strong></p>\n<pre
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
  <span class=\"n\">render</span>\n<span class=\"kn\">from</span> <span class=\"nn\">random</span>
  <span class=\"kn\">import</span> <span class=\"n\">randint</span>\n\n<span class=\"k\">def</span>
  <span class=\"nf\">if_demo</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">):</span>\n   <span class=\"n\">number</span> <span class=\"o\">=</span>
  <span class=\"n\">randint</span><span class=\"p\">(</span><span class=\"mi\">1</span><span
  class=\"p\">,</span><span class=\"mi\">10</span><span class=\"p\">)</span>\n   <span
  class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;if_else.html&#39;</span><span
  class=\"p\">,</span> <span class=\"p\">{</span><span class=\"s1\">&#39;num&#39;</span><span
  class=\"p\">:</span><span class=\"n\">number</span><span class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p>Here,
  we have used the key name as <code>num</code> indicating we can give different names
  to the key which needs to be used in the template to render the values.</p>\n<p><strong>app_name/urls.py</strong></p>\n<pre
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
  class=\"p\">(</span><span class=\"s1\">&#39;if/&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">if_demo</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;ifdemo&quot;</span><span class=\"p\">),</span>\n        <span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p><strong>templates/if_else.html</strong></p>\n<pre
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
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n    {{ num }}\n    {% if num &gt; 5 %}\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>It&#39;s Greater then 5<span
  class=\"p\">&lt;/</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n
  \   {% elif num == 5 %}\n        <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span
  class=\"p\">&gt;</span>It&#39;s five!<span class=\"p\">&lt;/</span><span class=\"nt\">h2</span><span
  class=\"p\">&gt;</span>\n    {% else %}\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>It&#39;s less than 5<span class=\"p\">&lt;/</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n    {% endif %}\n<span class=\"p\">&lt;/</span><span
  class=\"nt\">body</span><span class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span
  class=\"nt\">html</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639411425/blogmedia/templ3_exj0fv.png\"
  alt=\"if-else demo\" /></p>\n<p>So, as we can see that, we can use the if-else conditions
  in the template and that is already powerful. This can be a bit messy as to handle
  mathematical operations or conditions into a single condition. This can really be
  used for really large datasets that can be shimmed down to really less coding and
  also improve readability.</p>\n<h3>For loop</h3>\n<p>Now, the most crucial component
  of the Django templating language is the loops. We can actually iterate over objects/lists
  in the template. This becomes a huge concept for actually making a dynamic web application.
  We n\\might want to iterate over all the entries in a database, or any other form
  of data which can make the app a lot dynamic and feel real-time.</p>\n<p>The syntax
  of for loop is almost similar to the if-else condition. We just replace the condition
  with the iterator and the list/object from the view context. <code>{% for i in list
  %}</code>, also end the for loop like <code>{% endfor %}</code>.</p>\n<p><strong>app_name/views.py</strong></p>\n<pre
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
  class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p>We have created a simple Python
  list called <code>sports</code> and we parse them to the template using a dictionary
  object, <code>sport_list</code> as the key for storing the value of the <code>sports</code>
  list.</p>\n<p><strong>app_name/urls.py</strong></p>\n<pre class='wrapper'>\n\n<div
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
  class=\"p\">(</span><span class=\"s1\">&#39;for/&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">for_demo</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;fordemo&quot;</span><span class=\"p\">),</span>\n        <span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p><strong>templates/for.html</strong></p>\n<pre
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
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">ul</span><span
  class=\"p\">&gt;</span>\n        {% for sport in sport_list %}\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">li</span><span class=\"p\">&gt;</span>{{ sport }}<span class=\"p\">&lt;/</span><span
  class=\"nt\">li</span><span class=\"p\">&gt;</span>\n        {% endfor %}\n    <span
  class=\"p\">&lt;/</span><span class=\"nt\">ul</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;/</span><span class=\"nt\">body</span><span class=\"p\">&gt;</span>\n<span
  class=\"p\">&lt;/</span><span class=\"nt\">html</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639475328/blogmedia/templ3_q8z8fr.png\"
  alt=\"for loop demo\" /></p>\n<p>We have used simple for loop in Pythonic syntax,
  we use an iterator in this case, <code>sport</code> acts as an iterator. We use
  this to store values one by one from the list <code>sport_list</code> which was
  earlier passed in the views as a key in the dictionary.</p>\n<p>Hence, this is quite
  scalable and used to fetch the objects/entries in the database and hence making
  it a lot easier to make a dynamic web application faster.</p>\n<h2>Template Inheritance</h2>\n<p>So,
  far we have seen that we need to create the base template again and again like all
  the basic HTML elements, title, and all the basic structure. But what if, we can
  reuse a specific template in another and extend the functionality of that template
  into a new one. This avoids the redundancy of writing the entire basic template
  or the layout of a web app over and over again.</p>\n<p>To do that, Django has the
  Template inheritance. We can use a template as its basic layout or a specific component
  in the web application. Again, similar to the for, if-else blocks the syntax for
  inheriting a template is quite similar.</p>\n<p>Take, for example, the home.html
  which consisted of only a <code>&lt;h1&gt;</code> tag in it. We can use this kind
  of template in other templates to really make it the home page. For that, we first
  need to enclose the template in a <code>block</code>, which is what allows us to
  use it in other templates.\nTo create a <code>block</code>, we simply need to write
  the following syntax before the component which we do not want in other templates:</p>\n<p><strong>templates/home.html</strong></p>\n<pre
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
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">head</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n    <span class=\"p\">&lt;</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>Hello, World!<span class=\"p\">&lt;/</span><span class=\"nt\">h1</span><span
  class=\"p\">&gt;</span>\n    {% block body %}\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">p</span><span class=\"p\">&gt;</span>This is not going to get inherited
  <span class=\"p\">&lt;/</span><span class=\"nt\">p</span><span class=\"p\">&gt;</span>\n
  \   {% endblock %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">p</span><span
  class=\"p\">&gt;</span>This will be inherited<span class=\"p\">&lt;/</span><span
  class=\"nt\">p</span><span class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span
  class=\"nt\">body</span><span class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span
  class=\"nt\">html</span><span class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>In
  this we have used the <code>blocks</code> with a name like <code>body</code> as
  <code>{% block body %}</code> this can be anything you like. We end the block with
  the similar syntax as the for/if blocks as <code>{% endblock %}</code>. Anything
  in between the blocks i.e <code>block block_name</code> and <code>endblock</code>
  is not inherited i.e it is unique to this template.</p>\n<p>We will see how we can
  use this template in other templates. We will actually extend this template and
  use the blocks to render the content in the template.</p>\n<p><strong>templates/if_else.html</strong></p>\n<pre
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
  \       \n<div class='language-bar'>html</div>\n<div class=\"highlight\"><pre><span></span>{%
  extends &#39;home.html&#39; %}\n{% block body %}\n    {{ num }}\n    {% if num &gt;
  5 %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>It&#39;s
  Greater then 5<span class=\"p\">&lt;/</span><span class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n
  \   {% elif num == 5 %}\n    <span class=\"p\">&lt;</span><span class=\"nt\">h2</span><span
  class=\"p\">&gt;</span>It&#39;s five!<span class=\"p\">&lt;/</span><span class=\"nt\">h2</span><span
  class=\"p\">&gt;</span>\n    {% else %}\n    <span class=\"p\">&lt;</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>It&#39;s less than 5<span class=\"p\">&lt;/</span><span
  class=\"nt\">h2</span><span class=\"p\">&gt;</span>\n    {% endif %}\n{% endblock
  %}\n</pre></div>\n\n</pre>\n\n<p>So, we first say to Django to extend the <code>home</code>
  template i.e. the Django will load the blocks from this template only, remember
  it will just load and not use the blocks until we explicitly tell it to.</p>\n<p>To
  use the blocks or kind of plug in the template content in the <code>if_else.html</code>
  or any other template, we need to again call the <code>blocks</code>. Here, we need
  to write the content inside the <code>blocks</code> to properly parse the blocks
  as this is an HTML template. The order of opening and closing elements do matter.\nSo,
  when we say <code>endblock</code> the last part of the base template is loaded i.e.
  the closing <code>body</code> and <code>html</code> tags. This is like plugging
  the template as it is before and after the block body.</p>\n<p><strong>app_name/views.py</strong></p>\n<pre
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
  class=\"p\">)</span>\n\n<span class=\"kn\">from</span> <span class=\"nn\">random</span>
  <span class=\"kn\">import</span> <span class=\"n\">randint</span>\n\n<span class=\"k\">def</span>
  <span class=\"nf\">if_demo</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">):</span>\n   <span class=\"n\">number</span> <span class=\"o\">=</span>
  <span class=\"n\">randint</span><span class=\"p\">(</span><span class=\"mi\">1</span><span
  class=\"p\">,</span><span class=\"mi\">10</span><span class=\"p\">)</span>\n   <span
  class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;if_else.html&#39;</span><span
  class=\"p\">,</span> <span class=\"p\">{</span><span class=\"s1\">&#39;num&#39;</span><span
  class=\"p\">:</span><span class=\"n\">number</span><span class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p><strong>app_name/urls.py</strong></p>\n<pre
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
  class=\"s2\">&quot;home&quot;</span><span class=\"p\">),</span>\n        <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s1\">&#39;if/&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">if_demo</span><span
  class=\"p\">,</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;ifdemo&quot;</span><span class=\"p\">),</span>\n        <span
  class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639479714/blogmedia/tempinher2_enisls.png\"
  alt=\"\" />\n<img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639477721/blogmedia/tempinher_lk0op0.png\"
  alt=\"template inheritance demo\" /></p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639479954/blogmedia/template-inh_lc8szo.gif\"
  alt=\"\" /></p>\n<p>The above gif illustrates the example in a neat way. The block
  is loaded from the given template as the extended template and hence it plugs the
  block into the frame of the template.</p>\n<h2>Conclusion</h2>\n<p>So, from this
  part of the series, we were able to understand the concept of Templates in Django,
  we were able to use variables, loops, conditional statements, and template inheriting
  in a Django application. In the next part, we'll try to touch up with the static
  files and see how to properly structure and configure them.</p>\n<p>Thank you for
  reading, if you didn't understand any of the examples, please let me know, I'll
  be happy to share the code. Happy Coding :)</p>\n"
image_url: https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639387566/blogmedia/dj5temp_zukvj7.png
long_description: 'After learning the basics of views and URLs, we can now move on
  to the next concept i.e. Templates. In Django, Templates are quite an important
  component for the application as it acts as the  Templates are simply a  So, using
  HTML templates, we can '
now: 2025-05-01 18:11:33.311060
path: blog/posts/2021-12-14-Django-Basics-P5.md
prevnext: null
series:
- Django-Basics
- Django-Series
slug: django-basics-templates
status: published
subtitle: Rendering templates/html docs in the Django project/app
tags:
- django
- python
- web-development
templateKey: blog-post
title: 'Django Basics: Templates'
today: 2025-05-01
---

## Introduction

After learning the basics of views and URLs, we can now move on to the next concept i.e. Templates. In Django, Templates are quite an important component for the application as it acts as the `frontend` for the web application. With the help of templates and some features provided by Django, it becomes very intuitive and simple to make dynamic web content.
In this part, we understand what are templates and what is the way to render them. 

## What are Templates

Templates are simply a `html` document or kind of a wireframe for content to be displayed for the web app. Templates allow us to render some more relevant pieces of data rather than simple text HTTP responses as we did earlier. We can even re-use certain components of a template in other using the Django Templating Language (more on this later). 

So, using HTML templates, we can write a complete Webpage. If you are unfamiliar with HTML, you can check out the basics of HTML with this [playlist](https://www.youtube.com/playlist?list=PL081AC329706B2953). 

Even If you are not familiar with HTML, this tutorial might be quite basic and not overwhelm you with all the tags.

## Creating Templates

To create a Template, we can write a simple HTML document like the below:

Create a folder `templates` in the base folder, inside the templates folder, create a file `index.html`

**templates\index.html**

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Django Blog</title>
</head>
<body>
    <h1>Hello, World!</h1>
</body>
</html>
```

This is a simple HTML template, with the `<h1>` tags. As Django is a framework, there is a standard for storing all the templates for the project and application. There are a couple of standard of options:
- One of which is creating a `templates` folder in the root folder as discussed earlier, also we need to modify the `project_name/settings.py` file. 

Inside the `settings.py` file, we need to locate the `TEMPLATES` section and modify as below:

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
```   

In this snippet, we have changed the `DIRS` option to search the templates in the folder `templates` in the root directory. 

- The other standard is creating a templates folder in each application.

We can create the templates folder in each application instead of a single folder. 

## Rendering Templates

After creating a template and making the required settings to make sure Django is able to pick up those templates, we need to work with views and URLs to actually render those templates. 

There are a couple of ways to render templates in Django and some of them are discussed below:

### Using TemplateView

[TemplateView](https://docs.djangoproject.com/en/3.2/ref/class-based-views/base/#django.views.generic.base.TemplateView) is a class which is comes with `django.views.generic` library. This class allows us to render a template by providing in the name of the template, arguments or variables to be parsed, and so on. 

The simplest way to render a template is by the following way:

```python
from django.contrib import admin
from django.urls import path, include

from django.views.generic import TemplateView

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', TemplateView.as_view(template_name="index.html"), name="index"),
]
```

We need to import the `TemplateView` from the `django.core.generic` so as to use the class for rendering the template. 

The `TemplateView` class takes in a couple of arguments, we'll use the `template_name` as an argument that takes in the name of the template. Here, we use the `index.html` as the template which we created earlier. We don't need to specify the entire path to the template as we make modifications in the `settings.py` file to pick the template from the mentioned directory. We use `as_view` function to load the class as a function/view.  

Activate the virtual environment for the proper functioning of the project.

After activating the virtual environment we can run the server as follows:
```terminal
python manage.py runserver
```
   We can now see the following output and thus, we are now rendering a simple HTML template in Django.

![](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639384994/blogmedia/templ1_vbwp5d.png)

### Using render

We can also use the [render function](https://docs.djangoproject.com/en/4.0/topics/http/shortcuts/#render) from `django.shortcuts` to simply render a template. But we will create a Python function or a View to actually render the template. So, we'll create a View-URL map as we created in the [previous part](https://mr-destructive.github.io/techstructive-blog/django/python/web-development/2021/11/30/Django-Basics-P4.html).

Firstly, let's create a view function in the `post/views.py` file, more generally (`app_name/views.py` file). Firstly, we need to import the render function from `django.shortcuts` and then return the function call of render.

```python
from django.shortcuts import render

def home(request):
    return render(request,'index.html')
```   

And in the URLs, we'll create a different pattern like for e.g. 'home/'

```python
from django.urls import path
from post import views

urlpatterns=[
        path('home/',views.home,name="home"),
        ]
```

So, after creating the View-URL map and making sure the URL of the app is loaded in the project URLs, we can see the result as a simple HTML template.   

![](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639386932/blogmedia/templ2_rgoppj.png)
   
## Django Templating Language 

The [Django Templating Language](https://docs.djangoproject.com/en/3.2/ref/templates/language/) is Django's way of making templates more dynamic and easy to write dynamic web applications. We'll take a brief look at what we can do with this type of Templating Language in Django.

### Variables

This is the most common use case for the Django Templating Language/Engine as we can use the [variables](https://docs.djangoproject.com/en/3.2/ref/templates/language/#variables) from the Backend and inject it in the template. We can parse the variable into the template by the syntax : `{{ variable_name " }}}}`

To show its use cases, we can declare a variable in a view and then parse it in the Template. Though it is not dynamic right now we can later on fetch values from the database and store them in the form of variables in our views. 

**templates/home.html**
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Django Blog</title>
</head>
<body>
    <h1>Hello, {{ name }}</h1>
</body>
</html>
```


**post/views.py**
```python
from django.shortcuts import render

def variable_demo(request):
    name = "Kevin"
    return render(request, 'home.html', {'name':name})
    #The name can be anything, like a database query object, form detail, etc

```

As we can see the variable in views is passed as a dictionary in python. The reference key along with a value of the variable as the name of the variable. We will use the key in the templates to parse the value of the variable.

**post/urls.py**
```python
from django.urls import path
from post import views

urlpatterns=[
        path('home/',views.home,name="home"),
        path('vardemo/',views.variable_demo, name="var"),
        ]
```

![variable demo](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639389288/blogmedia/templ3_wvhagw.png)

As we can see, we were able to load the variable into the template using the Django Templating Engine. 

### Conditional statement

We can even use the conditional statement in the Template using a very simple syntax. We can use `{% if condition"  }} %}` to use certain special kinds of blocks in the Template. We need to end those blocks as well using the syntax `{% endif  %}`, here `if` can be other blocks which we'll explore ahead.

To create a basic if condition in the template, we can understand with the following example.

**app_name/views.py**
```python
from django.shortcuts import render
from random import randint

def if_demo(request):
   number = randint(1,10)
   return render(request, 'if_else.html', {'num':number})
```

Here, we have used the key name as `num` indicating we can give different names to the key which needs to be used in the template to render the values.

**app_name/urls.py**
```python
from django.urls import path
from post import views

urlpatterns = [
        path('if/', views.if_demo, name="ifdemo"),
        ]
```

**templates/if_else.html**
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Django Blog</title>
</head>
<body>
    {{ num }}
    {% if num > 5 %}
        <h2>It's Greater then 5</h2>
    {% elif num == 5 %}
        <h2>It's five!</h2>
    {% else %}
        <h2>It's less than 5</h2>
    {% endif %}
</body>
</html>
```
![if-else demo](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639411425/blogmedia/templ3_exj0fv.png)

So, as we can see that, we can use the if-else conditions in the template and that is already powerful. This can be a bit messy as to handle mathematical operations or conditions into a single condition. This can really be used for really large datasets that can be shimmed down to really less coding and also improve readability.


### For loop

Now, the most crucial component of the Django templating language is the loops. We can actually iterate over objects/lists in the template. This becomes a huge concept for actually making a dynamic web application. We n\might want to iterate over all the entries in a database, or any other form of data which can make the app a lot dynamic and feel real-time. 

The syntax of for loop is almost similar to the if-else condition. We just replace the condition with the iterator and the list/object from the view context. `{% for i in list %}`, also end the for loop like `{% endfor %}`.

**app_name/views.py**
```python
from django.shortcuts import render

def for_demo(request):
    sports = ('football', 'cricket', 'volleyball', 'hockey', 'basketball')
    return render(request, 'for.html', {'sport_list': sports})

```

We have created a simple Python list called `sports` and we parse them to the template using a dictionary object, `sport_list` as the key for storing the value of the `sports` list.

**app_name/urls.py**
```python
from django.urls import path
from post import views

urlpatterns = [
        path('for/', views.for_demo, name="fordemo"),
        ]
```

**templates/for.html**
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Django Blog</title>
</head>
<body>
    <ul>
        {% for sport in sport_list %}
        <li>{{ sport }}</li>
        {% endfor %}
    </ul>
</body>
</html>
```

![for loop demo](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639475328/blogmedia/templ3_q8z8fr.png)

We have used simple for loop in Pythonic syntax, we use an iterator in this case, `sport` acts as an iterator. We use this to store values one by one from the list `sport_list` which was earlier passed in the views as a key in the dictionary. 

Hence, this is quite scalable and used to fetch the objects/entries in the database and hence making it a lot easier to make a dynamic web application faster.

## Template Inheritance

So, far we have seen that we need to create the base template again and again like all the basic HTML elements, title, and all the basic structure. But what if, we can reuse a specific template in another and extend the functionality of that template into a new one. This avoids the redundancy of writing the entire basic template or the layout of a web app over and over again.

To do that, Django has the Template inheritance. We can use a template as its basic layout or a specific component in the web application. Again, similar to the for, if-else blocks the syntax for inheriting a template is quite similar. 

Take, for example, the home.html which consisted of only a `<h1>` tag in it. We can use this kind of template in other templates to really make it the home page. For that, we first need to enclose the template in a `block`, which is what allows us to use it in other templates.
To create a `block`, we simply need to write the following syntax before the component which we do not want in other templates:

**templates/home.html**
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Django Blog</title>
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
In this we have used the `blocks` with a name like `body` as `{% block body %}` this can be anything you like. We end the block with the similar syntax as the for/if blocks as `{% endblock %}`. Anything in between the blocks i.e `block block_name` and `endblock` is not inherited i.e it is unique to this template.

We will see how we can use this template in other templates. We will actually extend this template and use the blocks to render the content in the template.
   

**templates/if_else.html**
```html
{% extends 'home.html' %}
{% block body %}
    {{ num }}
    {% if num > 5 %}
    <h2>It's Greater then 5</h2>
    {% elif num == 5 %}
    <h2>It's five!</h2>
    {% else %}
    <h2>It's less than 5</h2>
    {% endif %}
{% endblock %}
```
   So, we first say to Django to extend the `home` template i.e. the Django will load the blocks from this template only, remember it will just load and not use the blocks until we explicitly tell it to. 

To use the blocks or kind of plug in the template content in the `if_else.html` or any other template, we need to again call the `blocks`. Here, we need to write the content inside the `blocks` to properly parse the blocks as this is an HTML template. The order of opening and closing elements do matter. 
So, when we say `endblock` the last part of the base template is loaded i.e. the closing `body` and `html` tags. This is like plugging the template as it is before and after the block body. 

**app_name/views.py**
```python
from django.shortcuts import render

def home(request):
    return render(request, 'home.html')

from random import randint

def if_demo(request):
   number = randint(1,10)
   return render(request, 'if_else.html', {'num':number})
```   

**app_name/urls.py**
```python
from django.urls import path
from post import views

urlpatterns = [
        path('', views.home, name="home"),
        path('if/', views.if_demo, name="ifdemo"),
        ]
```   
![](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639479714/blogmedia/tempinher2_enisls.png)
![template inheritance demo](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639477721/blogmedia/tempinher_lk0op0.png)      

![](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639479954/blogmedia/template-inh_lc8szo.gif)

The above gif illustrates the example in a neat way. The block is loaded from the given template as the extended template and hence it plugs the block into the frame of the template. 

## Conclusion

So, from this part of the series, we were able to understand the concept of Templates in Django, we were able to use variables, loops, conditional statements, and template inheriting in a Django application. In the next part, we'll try to touch up with the static files and see how to properly structure and configure them.

Thank you for reading, if you didn't understand any of the examples, please let me know, I'll be happy to share the code. Happy Coding :)