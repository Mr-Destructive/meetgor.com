---
article_html: "<h2>Introduction</h2>\n<p>I have an Article Form where I load my post
  into it directly, it might have frontmatter. So what I wish to achieve is when I
  paste in the content, the frontmatter should be picked up and it should render the
  form fields like <code>title</code>, <code>description</code>, and then also remove
  the frontmatter from the content.</p>\n<p>To do that, we will require a model to
  work with and a form based on that model. We will exclude a few fields from that
  model so as to process these attributes on the server side. I am working on my Blog
  project which is a simple Django application.  You can get the source code for the
  project on the <a href=\"https://github.com/mr-destructive/techstructive-blog/\">GitHub
  repository</a>.</p>\n<h2>Django Project Context</h2>\n<p>The techstructive-blog
  is a django project, which has a couple of applications currently, not in a good
  situation. There are apps like <code>article</code>, <code>blog</code>, and <code>user</code>.
  This project has templates and static folder in the base directory. The project
  is deployed on <a href=\"https://www.railway.app\">railway</a> this is an always
  under development project, you can check out the <a href=\"https://django-blog.up.railway.app\">techstructive-blog</a>.
  We can get the bits and pieces of the project details required for understanding
  what I want to do with the following sections.</p>\n<h3>Article Model</h3>\n<p>We
  have an <code>Article</code> model with attributes like <code>title</code>, <code>description</code>,
  \ <code>author</code> as a Foreign Key to the user model, and a few other attributes
  which is not related to what we are trying to achieve right now or at least don't
  require an explanation. We have a model method called <code>get_absolute_url</code>
  for getting a URL in order to redirect the client when the model instance is created
  or updated from the backend. You can definitely check out the details of these small
  components or templates in the project repository.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># articles/models.py</span>\n\n\n<span class=\"k\">class</span> <span
  class=\"nc\">Article</span><span class=\"p\">(</span><span class=\"n\">TimeStampedModel</span><span
  class=\"p\">):</span>\n    <span class=\"k\">class</span> <span class=\"nc\">Article_Status</span><span
  class=\"p\">(</span><span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">TextChoices</span><span class=\"p\">):</span>\n        <span class=\"n\">DRAFT</span>
  <span class=\"o\">=</span> <span class=\"p\">(</span>\n            <span class=\"s2\">&quot;DRAFT&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"n\">_</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;Draft&quot;</span><span class=\"p\">),</span>\n        <span
  class=\"p\">)</span>\n        <span class=\"n\">PUBLISHED</span> <span class=\"o\">=</span>
  <span class=\"p\">(</span>\n            <span class=\"s2\">&quot;PUBLISHED&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"n\">_</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;Published&quot;</span><span class=\"p\">),</span>\n        <span
  class=\"p\">)</span>\n\n    <span class=\"n\">title</span> <span class=\"o\">=</span>
  <span class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">CharField</span><span
  class=\"p\">(</span><span class=\"n\">max_length</span><span class=\"o\">=</span><span
  class=\"mi\">128</span><span class=\"p\">)</span>\n    <span class=\"n\">description</span>
  <span class=\"o\">=</span> <span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">CharField</span><span class=\"p\">(</span><span class=\"n\">max_length</span><span
  class=\"o\">=</span><span class=\"mi\">256</span><span class=\"p\">)</span>\n    <span
  class=\"n\">content</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">TextField</span><span class=\"p\">(</span><span
  class=\"n\">default</span><span class=\"o\">=</span><span class=\"s2\">&quot;&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">null</span><span class=\"o\">=</span><span
  class=\"kc\">False</span><span class=\"p\">,</span> <span class=\"n\">blank</span><span
  class=\"o\">=</span><span class=\"kc\">False</span><span class=\"p\">)</span>\n
  \   <span class=\"n\">status</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">CharField</span><span class=\"p\">(</span>\n
  \       <span class=\"n\">max_length</span><span class=\"o\">=</span><span class=\"mi\">16</span><span
  class=\"p\">,</span>\n        <span class=\"n\">choices</span><span class=\"o\">=</span><span
  class=\"n\">Article_Status</span><span class=\"o\">.</span><span class=\"n\">choices</span><span
  class=\"p\">,</span>\n        <span class=\"n\">default</span><span class=\"o\">=</span><span
  class=\"n\">Article_Status</span><span class=\"o\">.</span><span class=\"n\">DRAFT</span><span
  class=\"p\">,</span>\n    <span class=\"p\">)</span>\n    <span class=\"n\">blog</span>
  <span class=\"o\">=</span> <span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">ForeignKey</span><span class=\"p\">(</span><span class=\"n\">Blog</span><span
  class=\"p\">,</span> <span class=\"n\">on_delete</span><span class=\"o\">=</span><span
  class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">CASCADE</span><span
  class=\"p\">,</span> <span class=\"n\">null</span><span class=\"o\">=</span><span
  class=\"kc\">True</span><span class=\"p\">,</span> <span class=\"n\">blank</span><span
  class=\"o\">=</span><span class=\"kc\">True</span><span class=\"p\">)</span>\n    <span
  class=\"n\">author</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">ForeignKey</span><span class=\"p\">(</span><span
  class=\"n\">User</span><span class=\"p\">,</span> <span class=\"n\">on_delete</span><span
  class=\"o\">=</span><span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">CASCADE</span><span class=\"p\">,</span> <span class=\"n\">related_name</span><span
  class=\"o\">=</span><span class=\"s2\">&quot;author&quot;</span><span class=\"p\">)</span>\n\n
  \   <span class=\"k\">def</span> <span class=\"fm\">__str__</span><span class=\"p\">(</span><span
  class=\"bp\">self</span><span class=\"p\">):</span>\n        <span class=\"k\">return</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">title</span>\n\n
  \   <span class=\"k\">def</span> <span class=\"nf\">get_absolute_url</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"p\">):</span>      \n
  \       <span class=\"k\">return</span> <span class=\"n\">reverse</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;articles:article-detail&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">args</span><span class=\"o\">=</span><span class=\"p\">[</span><span
  class=\"nb\">str</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">id</span><span class=\"p\">)])</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the below snippet, we have the forms defined in the article application for creating
  or updating of article instance.  We will be using model forms as our form data
  should contain fields related to a model in this case the <code>Article</code> model.
  So when we inherit the <code>forms.ModelForm</code> in our custom <code>ArticleForm</code>
  we simply need to specify the model and it will take in all the attributes of that
  model by default, but if we specify the <code>fields</code> or <code>exclude</code>
  tuples, it will include only or exclude only the provided list of attributes from
  the model.</p>\n<p>We have also added the widgets for the form fields which will
  allow us to customize the way the fields in the template/form will render. We can
  specify the HTML attributes like <code>width</code>, <code>height</code>, <code>style</code>,
  etc.</p>\n<h3>Article Form</h3>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># article/forms.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django</span> <span class=\"kn\">import</span> <span class=\"n\">forms</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">.models</span> <span class=\"kn\">import</span>
  <span class=\"n\">Article</span>\n\n\n<span class=\"k\">class</span> <span class=\"nc\">ArticleForm</span><span
  class=\"p\">(</span><span class=\"n\">forms</span><span class=\"o\">.</span><span
  class=\"n\">ModelForm</span><span class=\"p\">):</span>\n    <span class=\"k\">class</span>
  <span class=\"nc\">Meta</span><span class=\"p\">:</span>\n        <span class=\"n\">model</span>
  <span class=\"o\">=</span> <span class=\"n\">Article</span>\n        <span class=\"n\">exclude</span>
  <span class=\"o\">=</span> <span class=\"p\">(</span>\n            <span class=\"s2\">&quot;created&quot;</span><span
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
  class=\"s2\">&quot;description&quot;</span><span class=\"p\">:</span> <span class=\"n\">forms</span><span
  class=\"o\">.</span><span class=\"n\">TextInput</span><span class=\"p\">(</span>\n
  \               <span class=\"n\">attrs</span><span class=\"o\">=</span><span class=\"p\">{</span>\n
  \                   <span class=\"s2\">&quot;class&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;form-control&quot;</span><span class=\"p\">,</span>\n                    <span
  class=\"s2\">&quot;style&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;max-width:
  900px; &quot;</span><span class=\"p\">,</span>\n                    <span class=\"s2\">&quot;placeholder&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Description&quot;</span><span class=\"p\">,</span>\n
  \               <span class=\"p\">}</span>\n            <span class=\"p\">),</span>\n
  \           <span class=\"s2\">&quot;content&quot;</span><span class=\"p\">:</span>
  <span class=\"n\">forms</span><span class=\"o\">.</span><span class=\"n\">Textarea</span><span
  class=\"p\">(</span>\n                <span class=\"n\">attrs</span><span class=\"o\">=</span><span
  class=\"p\">{</span>\n                    <span class=\"s2\">&quot;class&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;form-control post-body&quot;</span><span
  class=\"p\">,</span>\n                    <span class=\"s2\">&quot;id&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;text-content&quot;</span><span class=\"p\">,</span>\n
  \                   <span class=\"s2\">&quot;style&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;max-width:900px;&quot;</span><span class=\"p\">,</span>\n
  \                   <span class=\"s2\">&quot;hx-post&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;/article/meta/&quot;</span><span class=\"p\">,</span>\n
  \                   <span class=\"s2\">&quot;placeholder&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;Content&quot;</span><span class=\"p\">,</span>\n                <span
  class=\"p\">}</span>\n            <span class=\"p\">),</span>\n            <span
  class=\"s2\">&quot;blog&quot;</span><span class=\"p\">:</span> <span class=\"n\">forms</span><span
  class=\"o\">.</span><span class=\"n\">Select</span><span class=\"p\">(</span>\n
  \               <span class=\"n\">attrs</span><span class=\"o\">=</span><span class=\"p\">{</span>\n
  \                   <span class=\"s2\">&quot;class&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;form-control&quot;</span><span class=\"p\">,</span>\n                    <span
  class=\"s2\">&quot;placeholder&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;Blog
  Publication&quot;</span><span class=\"p\">,</span>\n                <span class=\"p\">}</span>\n
  \           <span class=\"p\">),</span>\n        <span class=\"p\">}</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  these are my models and form files in the article app. Using htmx and without any
  javascript I want to update the form so that it picks up the front matter in the
  content field which is a text area and fills the title, description other attributes
  automatically for me.</p>\n<p>This can be done in a lot of ways, but I will be sharing
  one of the ways that I recently used in my blog project. This process involves creating
  a class-based view and adding a <code>POST</code> method that won't post any data
  to the backend but will send necessary data to the view.</p>\n<p>Let's see the process
  before diving into any of the code:</p>\n<h2>Gist of the Process</h2>\n<ul>\n<li>Attach
  a <code>hx-post</code> attribute to the form field for sending the request to a
  view</li>\n<li>When the request is sent, the data is loaded with <code>request.POST</code>,
  it is cleaned and converted in python-readable format with json.</li>\n<li>Once
  we have the data, we try to use the <code>frontmatter.loads</code> function that
  will load the content and if we have a frontmatter in the text, it will load it
  as a <code>frontmatter.POST</code> object.</li>\n<li>We will extract <code>title</code>,
  <code>description</code>, and other data fields from the object.</li>\n<li>We will
  initialize a Form instance of Article, with the initial data values as the extracted
  data from the frontmatter.</li>\n<li>Now we have two options:\n<ul>\n<li>If the
  article instance already exists i.e. we are updating the article</li>\n<li>Else
  we are creating a new article</li>\n</ul>\n</li>\n</ul>\n<p>Accordingly, we will
  load the form in the respective templates i.e. <code>update.html</code> for updating
  the existing articles and <code>article-form.html</code> for a new article.</p>\n<h2>Adding
  HTMX Magic</h2>\n<p>If you'd have seen we have a <code>hx-post</code> attribute
  in the <code>article/forms.py</code> file, the <code>content</code> widget has a
  <code>hx-post</code> attribute which sends a post request to the <code>article/meta/</code>
  URL route. This route we will bind to the <code>ArticleMetaView</code> which we
  will define in a few moments. This is usually sent once we change a certain text
  in the content field, so we can modify it as per your requirement with <code>hx-trigger</code>
  that can enable us to specify the trigger event or the type of trigger we want.
  Further, you can read from the <a href=\"https://htmx.org/docs/#trigger-modifiers\">htmx
  docs</a> about these triggers and other attributes.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"n\">app_name</span> <span class=\"o\">=</span> <span class=\"s2\">&quot;articles&quot;</span>\n\n<span
  class=\"n\">urlpatterns</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">ArticleCreateView</span><span class=\"o\">.</span><span class=\"n\">as_view</span><span
  class=\"p\">(),</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;article-create&quot;</span><span class=\"p\">),</span>\n    <span
  class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;&lt;int:pk&gt;/&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">ArticleDetailView</span><span class=\"o\">.</span><span class=\"n\">as_view</span><span
  class=\"p\">(),</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;article-detail&quot;</span><span class=\"p\">),</span>\n    <span
  class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;delete/&lt;int:pk&gt;/&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">ArticleDetailView</span><span class=\"o\">.</span><span class=\"n\">as_view</span><span
  class=\"p\">(),</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;article-delete&quot;</span><span class=\"p\">),</span>\n    <span
  class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;edit/&lt;int:pk&gt;&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">ArticleDetailView</span><span class=\"o\">.</span><span class=\"n\">as_view</span><span
  class=\"p\">(),</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;article-update&quot;</span><span class=\"p\">),</span>\n\n    <span
  class=\"c1\"># the new view that we will create</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;meta/&quot;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">ArticleMetaView</span><span
  class=\"o\">.</span><span class=\"n\">as_view</span><span class=\"p\">(),</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s2\">&quot;article-meta&quot;</span><span
  class=\"p\">),</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<h2>Capture
  Frontmatter Meta-data View</h2>\n<p>Along with the Create, Detail/List, Update,
  Delete View, I will create a separate class called <code>ArticleMetaView</code>
  that will fetch the form fields and render the templates again but this time it
  will fill in the frontmatter meta-data in the fields if the content is parsed with
  the relvant frontmatter.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># articles/view.py</span>\n\n<span class=\"k\">class</span> <span class=\"nc\">ArticleMetaView</span><span
  class=\"p\">(</span><span class=\"n\">View</span><span class=\"p\">):</span>\n    <span
  class=\"n\">model</span> <span class=\"o\">=</span> <span class=\"n\">Article</span>\n\n
  \   <span class=\"k\">def</span> <span class=\"nf\">post</span><span class=\"p\">(</span><span
  class=\"bp\">self</span><span class=\"p\">,</span> <span class=\"n\">request</span><span
  class=\"p\">,</span> <span class=\"o\">*</span><span class=\"n\">args</span><span
  class=\"p\">,</span> <span class=\"o\">**</span><span class=\"n\">kwargs</span><span
  class=\"p\">):</span>\n        \n        <span class=\"n\">data</span> <span class=\"o\">=</span>
  <span class=\"n\">json</span><span class=\"o\">.</span><span class=\"n\">loads</span><span
  class=\"p\">(</span><span class=\"n\">json</span><span class=\"o\">.</span><span
  class=\"n\">dumps</span><span class=\"p\">(</span><span class=\"nb\">dict</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">POST</span><span class=\"p\">)))</span>\n        <span class=\"n\">loaded_frontmatter</span>
  <span class=\"o\">=</span> <span class=\"n\">frontmatter</span><span class=\"o\">.</span><span
  class=\"n\">loads</span><span class=\"p\">(</span><span class=\"n\">data</span><span
  class=\"p\">[</span><span class=\"s1\">&#39;content&#39;</span><span class=\"p\">][</span><span
  class=\"mi\">0</span><span class=\"p\">])</span>\n\n       <span class=\"c1\">#
  frontmatter has keys i.e. attributes like title, description, etc.</span>\n        <span
  class=\"k\">if</span> <span class=\"nb\">dict</span><span class=\"p\">(</span><span
  class=\"n\">loaded_frontmatter</span><span class=\"p\">):</span>\n            <span
  class=\"n\">article_title</span> <span class=\"o\">=</span> <span class=\"n\">loaded_frontmatter</span><span
  class=\"p\">[</span><span class=\"s1\">&#39;title&#39;</span><span class=\"p\">]</span>\n
  \           <span class=\"n\">article_description</span> <span class=\"o\">=</span>
  <span class=\"n\">loaded_frontmatter</span><span class=\"p\">[</span><span class=\"s1\">&#39;description&#39;</span><span
  class=\"p\">]</span>\n            <span class=\"n\">form</span> <span class=\"o\">=</span>
  <span class=\"n\">ArticleForm</span><span class=\"p\">(</span><span class=\"n\">initial</span><span
  class=\"o\">=</span><span class=\"p\">{</span><span class=\"s1\">&#39;title&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">article_title</span><span class=\"p\">,</span>
  \n            <span class=\"s1\">&#39;description&#39;</span><span class=\"p\">:</span>
  <span class=\"n\">article_description</span><span class=\"p\">,</span> <span class=\"s1\">&#39;content&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">loaded_frontmatter</span><span class=\"o\">.</span><span
  class=\"n\">content</span><span class=\"p\">})</span>\n            <span class=\"n\">context</span>
  <span class=\"o\">=</span> <span class=\"p\">{</span><span class=\"s1\">&#39;form&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">form</span><span class=\"p\">}</span>\n            <span
  class=\"n\">article_list</span> <span class=\"o\">=</span> <span class=\"n\">Article</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">filter</span><span class=\"p\">(</span><span class=\"n\">title</span><span
  class=\"o\">=</span><span class=\"n\">article_title</span><span class=\"p\">)</span>\n
  \           <span class=\"k\">if</span> <span class=\"n\">article_list</span><span
  class=\"p\">:</span>\n                <span class=\"n\">article</span> <span class=\"o\">=</span>
  <span class=\"n\">article_list</span><span class=\"o\">.</span><span class=\"n\">last</span><span
  class=\"p\">()</span>\n                <span class=\"n\">context</span><span class=\"p\">[</span><span
  class=\"s1\">&#39;article&#39;</span><span class=\"p\">]</span> <span class=\"o\">=</span>
  <span class=\"n\">article</span>\n                <span class=\"k\">return</span>
  <span class=\"n\">render</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;articles/edit_article.html&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">context</span><span class=\"p\">)</span>\n
  \           <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;articles/article_form.html&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">context</span><span class=\"p\">)</span>\n\n        <span class=\"n\">article_list</span>
  <span class=\"o\">=</span> <span class=\"n\">Article</span><span class=\"o\">.</span><span
  class=\"n\">objects</span><span class=\"o\">.</span><span class=\"n\">filter</span><span
  class=\"p\">(</span><span class=\"n\">title</span><span class=\"o\">=</span><span
  class=\"n\">data</span><span class=\"p\">[</span><span class=\"s1\">&#39;title&#39;</span><span
  class=\"p\">][</span><span class=\"mi\">0</span><span class=\"p\">])</span>\n       \n
  \      <span class=\"c1\"># if the article title has been already taken i.e. we
  are updating an article</span>\n\n        <span class=\"k\">if</span> <span class=\"n\">article_list</span><span
  class=\"p\">:</span>\n            <span class=\"n\">article</span> <span class=\"o\">=</span>
  <span class=\"n\">article_list</span><span class=\"o\">.</span><span class=\"n\">last</span><span
  class=\"p\">()</span>\n            <span class=\"n\">form</span> <span class=\"o\">=</span>
  <span class=\"n\">ArticleForm</span><span class=\"p\">(</span><span class=\"n\">data</span><span
  class=\"o\">=</span><span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">POST</span><span class=\"p\">)</span>\n            <span class=\"n\">context</span>
  <span class=\"o\">=</span> <span class=\"p\">{</span><span class=\"s1\">&#39;form&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">form</span><span class=\"p\">}</span>\n            <span
  class=\"n\">context</span><span class=\"p\">[</span><span class=\"s1\">&#39;article&#39;</span><span
  class=\"p\">]</span> <span class=\"o\">=</span> <span class=\"n\">article</span>\n
  \           <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;articles/edit_article.html&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">context</span><span class=\"p\">)</span>\n\n        <span class=\"n\">form</span>
  <span class=\"o\">=</span> <span class=\"n\">ArticleForm</span><span class=\"p\">(</span><span
  class=\"n\">data</span><span class=\"o\">=</span><span class=\"n\">request</span><span
  class=\"o\">.</span><span class=\"n\">POST</span><span class=\"p\">)</span>\n        <span
  class=\"n\">context</span> <span class=\"o\">=</span> <span class=\"p\">{</span><span
  class=\"s1\">&#39;form&#39;</span><span class=\"p\">:</span> <span class=\"n\">form</span><span
  class=\"p\">}</span>\n        <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;articles/article_form.html&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above <code>ArticleMetaView</code> we have created a <code>post</code> method
  as we want to get hold of the content from the form. So, we start by extracting
  and converting the <code>request.body</code> data into an appropriate type for easily
  working with python. So, the <code>request.body</code> will contain the data like
  <code>csrf_token</code>, <code>form_data</code>, etc. received from the frontend
  template. We store the received data as <code>data</code> and now from this data,
  we can load the content field which will have the content information.</p>\n<p>Firstly
  we will extract the <code>request.body</code> which will contain the data from the
  form as we have made a <code>POST</code> request to this endpoint. For doing that
  we need to parse the content in a apropriate format such that it is python friendly.
  So we wrap the <code>request.body</code> into json format and then decode it back
  into the json string. This will give us the dict of the request data.</p>\n<pre
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
  class=\"n\">data</span> <span class=\"o\">=</span> <span class=\"n\">json</span><span
  class=\"o\">.</span><span class=\"n\">loads</span><span class=\"p\">(</span><span
  class=\"n\">json</span><span class=\"o\">.</span><span class=\"n\">dumps</span><span
  class=\"p\">(</span><span class=\"nb\">dict</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"o\">.</span><span class=\"n\">POST</span><span
  class=\"p\">)))</span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>{&#39;csrfmiddlewaretoken&#39;:
  [&#39;bSYJxD39XH509tD1tZGd0WU21PUaKaLeqjjGbyzRvLXF4P8iIxb5l0fmTWVFjELQ&#39;], &#39;title&#39;:
  [&#39;test2&#39;], &#39;description&#39;: [&#39;test&#39;], &#39;content&#39;: [&#39;test
  something&#39;], &#39;status&#39;: [&#39;DRAFT&#39;], &#39;blog&#39;: [&#39;&#39;]}\n</pre></div>\n\n</pre>\n\n<p>So,
  this will grab the request data as a dict, we can then extract the data from this
  as it has data from the Form fields. We are interested in the content field in the
  Form, so we can get it by specifying the key <code>content</code> from the extracted
  data. But as we can see the data doesn't contain the actual data instead it is wrapped
  in a list i.e. <code>['test something']</code>, so we will have to index it and
  then fetch it.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"n\">content_string</span> <span class=\"o\">=</span> <span class=\"n\">data</span><span
  class=\"p\">[</span><span class=\"s1\">&#39;content&#39;</span><span class=\"p\">][</span><span
  class=\"mi\">0</span><span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>This
  will give us the exact content field as a string. So, we can now move into extracting
  the frontmatter from the fields.</p>\n<p>Now, we can use the <a href=\"https://python-frontmatter.readthedocs.io/en/latest/index.html\">frontmatter</a>
  library to parse the content into the <a href=\"https://python-frontmatter.readthedocs.io/en/latest/api.html#frontmatter.loads\">loads</a>
  funciton and extract the frontmatter if it is present in the content field. The
  frontmatter library has a <code>loads</code> function which takes in a string and
  can give out a <a href=\"https://python-frontmatter.readthedocs.io/en/latest/api.html#post-objects\">frontmatter.Post</a>
  object. The loads function is differnet from the <a href=\"https://python-frontmatter.readthedocs.io/en/latest/api.html#frontmatter.load\">load</a>
  function as the load frunciton is for reading data from a stream of bytes i.e. a
  file or othe related byte object. The differnece is subtle but it took a read at
  the <a href=\"https://python-frontmatter.readthedocs.io/en/latest/api.html#module-frontmatter\">documentation</a>.</p>\n<pre
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
  class=\"n\">post</span> <span class=\"o\">=</span> <span class=\"n\">data</span><span
  class=\"p\">[</span><span class=\"s1\">&#39;content&#39;</span><span class=\"p\">][</span><span
  class=\"mi\">0</span><span class=\"p\">]</span>\n<span class=\"n\">loaded_frontmatter</span>
  <span class=\"o\">=</span> <span class=\"n\">frontmatter</span><span class=\"o\">.</span><span
  class=\"n\">loads</span><span class=\"p\">(</span><span class=\"n\">post</span><span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>This wil load the content and
  give us a <code>frontmatter.Post</code> as said earlier. This will contain a dict
  with all the frontmatter if it has any and will by default parse the non-frontmatter
  data i.e. the remaining text into the <code>content</code> key. We need a chack
  if the Form field had any fronmatter this can be checked by the <code>dict(loaded_frontmatter)</code>
  which will return None if it cannot load the frontmatter.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"n\">loaded_frontmatter</span> <span class=\"o\">=</span> <span class=\"n\">frontmatter</span><span
  class=\"o\">.</span><span class=\"n\">loads</span><span class=\"p\">(</span><span
  class=\"n\">data</span><span class=\"p\">[</span><span class=\"s1\">&#39;content&#39;</span><span
  class=\"p\">][</span><span class=\"mi\">0</span><span class=\"p\">])</span>\n<span
  class=\"k\">if</span> <span class=\"nb\">dict</span><span class=\"p\">(</span><span
  class=\"n\">loaded_frontmatter</span><span class=\"p\">):</span>\n  <span class=\"nb\">print</span><span
  class=\"p\">(</span><span class=\"n\">loaded_frontmatter</span><span class=\"o\">.</span><span
  class=\"n\">keys</span><span class=\"p\">())</span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>dict_keys([&#39;templateKey&#39;,
  &#39;title&#39;, &#39;description&#39;, &#39;date&#39;, &#39;status&#39;])\n</pre></div>\n\n</pre>\n\n<p>So
  once we have the frontmatter loaded we can get specific keys from it and initialize
  the form vaues to them. But we have made clear distictions that we want to perform
  a specific task if we have frontmatter keys in the content field of the Form else
  we can do something else.</p>\n<p>First let's handle the loading of the frontmatter
  into the form. For doing that we will get all the required attributes from the frontmatter
  like <code>title</code>, <code>description</code>, <code>content</code>, etc which
  can be accessed normally as we extract the value from a key in a dict.</p>\n<p>Once
  we have got those keys, we can start filling in the Form data with initial values.
  The <a href=\"https://docs.djangoproject.com/en/4.0/topics/forms/modelforms/\">Django
  Model form</a> takes in a parameter like <a href=\"https://docs.djangoproject.com/en/4.0/topics/forms/modelforms/#providing-initial-values\">initial</a>
  which can be a dict of the fiields along with the value that can be used for rendering
  the form initially when we load the template.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"n\">article_title</span> <span class=\"o\">=</span> <span class=\"n\">loaded_frontmatter</span><span
  class=\"p\">[</span><span class=\"s1\">&#39;title&#39;</span><span class=\"p\">]</span>\n<span
  class=\"n\">article_description</span> <span class=\"o\">=</span> <span class=\"n\">loaded_frontmatter</span><span
  class=\"p\">[</span><span class=\"s1\">&#39;description&#39;</span><span class=\"p\">]</span>\n\n<span
  class=\"n\">form</span> <span class=\"o\">=</span> <span class=\"n\">ArticleForm</span><span
  class=\"p\">(</span><span class=\"n\">initial</span><span class=\"o\">=</span><span
  class=\"p\">{</span><span class=\"s1\">&#39;title&#39;</span><span class=\"p\">:</span>
  <span class=\"n\">article_title</span><span class=\"p\">,</span> <span class=\"s1\">&#39;description&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">article_description</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;content&#39;</span><span class=\"p\">:</span> <span class=\"n\">loaded_frontmatter</span><span
  class=\"o\">.</span><span class=\"n\">content</span><span class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p>This
  will take in a <code>ArticleForm</code> and fill the initial values like <code>title</code>,
  <code>description</code>, etc which we have provided in the dict with the values.
  Now, we need to parse this form in the current template or re-render the template.
  But before that, we need to parse this context into the template. We will create
  a dict with <code>form</code> as the key which can be used to render in the template.</p>\n<p>Also,
  we have a two ways here, either the user is creating a new article or it is updating
  a existing article. We need to make sure that we preserve the initial fields in
  the form as we are updating the existing article. So, we can filter the article
  objects as per the title of the current title and then if we find a article with
  that title, we will parse the context with that article object.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"n\">article_list</span> <span class=\"o\">=</span> <span class=\"n\">Article</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">filter</span><span class=\"p\">(</span><span class=\"n\">title</span><span
  class=\"o\">=</span><span class=\"n\">article_title</span><span class=\"p\">)</span>\n<span
  class=\"k\">if</span> <span class=\"n\">article_list</span><span class=\"p\">:</span>\n
  \   <span class=\"n\">article</span> <span class=\"o\">=</span> <span class=\"n\">article_list</span><span
  class=\"o\">.</span><span class=\"n\">last</span><span class=\"p\">()</span>\n    <span
  class=\"n\">context</span><span class=\"p\">{</span>\n      <span class=\"s1\">&#39;form&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">form</span><span class=\"p\">,</span>\n      <span
  class=\"s1\">&#39;article&#39;</span><span class=\"p\">:</span> <span class=\"n\">article</span>\n
  \   <span class=\"p\">}</span>\n    <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;articles/edit_article.html&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">context</span><span class=\"p\">)</span>\n<span class=\"n\">context</span>
  <span class=\"o\">=</span> <span class=\"p\">{</span><span class=\"s1\">&#39;form&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">form</span><span class=\"p\">}</span>\n<span
  class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;articles/article_form.html&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>Now,
  we have form data along with the article instance used for rendering the form with
  appropriate content. So, this will work for editing an already existing article.
  For a new article, we have to simply parse the form to the template and it will
  render the title picked from the fotnmatter or leave it empty.</p>\n<p>Similarly,
  for the article with no frontmatter we will iterate over the article and if the
  article's title already exist, we will render the article data with the form else
  render the form with the parsed title and other meta-data in the form.</p>\n<video
  width=\"800\" height=\"450\" controls>\n  <source src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1659370006/blog-media/frontmatter-load-htmx.mp4\"
  type=\"video/mp4\">\n</video>\n<p>So that is how we render the form data with frontmatter
  into appropriate meta-data in the form. We have used Django forms and make use of
  HTMX for the dynamic updation of form.</p>\n<div class='prevnext'>\n    <style type='text/css'>\n
  \   :root {\n      --prevnext-color-text: #eefbfe;\n      --prevnext-color-angle:
  #ff6600;\n      --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"light\"]
  {\n      --prevnext-color-text: #1f2022;\n      --prevnext-color-angle: #ffeb00;\n
  \     --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"dark\"] {\n      --prevnext-color-text:
  #eefbfe;\n      --prevnext-color-angle: #ff6600;\n      --prevnext-subtitle-brightness:
  3;\n    }\n    .prevnext {\n      display: flex;\n      flex-direction: row;\n      justify-content:
  space-around;\n      align-items: flex-start;\n    }\n    .prevnext a {\n      display:
  flex;\n      align-items: center;\n      width: 100%;\n      text-decoration: none;\n
  \   }\n    a.next {\n      justify-content: flex-end;\n    }\n    .prevnext a:hover
  {\n      background: #00000006;\n    }\n    .prevnext-subtitle {\n      color: var(--prevnext-color-text);\n
  \     filter: brightness(var(--prevnext-subtitle-brightness));\n      font-size:
  .8rem;\n    }\n    .prevnext-title {\n      color: var(--prevnext-color-text);\n
  \     font-size: 1rem;\n    }\n    .prevnext-text {\n      max-width: 30vw;\n    }\n
  \   </style>\n<pre><code>&lt;a class='prev' href='/libsql-query-remote-db'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;LibSQL:
  Query a remote Turso database with cURL&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a
  class='next' href='/python-chain-assignment-gotcha'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Gotcha
  with Chained Assignment in Python&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot;
  height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
cover: ''
date: 2022-08-01
datetime: 2022-08-01 00:00:00+00:00
description: Rendering frontatter from content field into the Template Form field
  using HTMX and frontmatter libraries
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/tils/2022-08-01-Django-Frontmatter-forms.md
html: "<h2>Introduction</h2>\n<p>I have an Article Form where I load my post into
  it directly, it might have frontmatter. So what I wish to achieve is when I paste
  in the content, the frontmatter should be picked up and it should render the form
  fields like <code>title</code>, <code>description</code>, and then also remove the
  frontmatter from the content.</p>\n<p>To do that, we will require a model to work
  with and a form based on that model. We will exclude a few fields from that model
  so as to process these attributes on the server side. I am working on my Blog project
  which is a simple Django application.  You can get the source code for the project
  on the <a href=\"https://github.com/mr-destructive/techstructive-blog/\">GitHub
  repository</a>.</p>\n<h2>Django Project Context</h2>\n<p>The techstructive-blog
  is a django project, which has a couple of applications currently, not in a good
  situation. There are apps like <code>article</code>, <code>blog</code>, and <code>user</code>.
  This project has templates and static folder in the base directory. The project
  is deployed on <a href=\"https://www.railway.app\">railway</a> this is an always
  under development project, you can check out the <a href=\"https://django-blog.up.railway.app\">techstructive-blog</a>.
  We can get the bits and pieces of the project details required for understanding
  what I want to do with the following sections.</p>\n<h3>Article Model</h3>\n<p>We
  have an <code>Article</code> model with attributes like <code>title</code>, <code>description</code>,
  \ <code>author</code> as a Foreign Key to the user model, and a few other attributes
  which is not related to what we are trying to achieve right now or at least don't
  require an explanation. We have a model method called <code>get_absolute_url</code>
  for getting a URL in order to redirect the client when the model instance is created
  or updated from the backend. You can definitely check out the details of these small
  components or templates in the project repository.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"c1\"># articles/models.py</span>\n\n\n<span class=\"k\">class</span> <span
  class=\"nc\">Article</span><span class=\"p\">(</span><span class=\"n\">TimeStampedModel</span><span
  class=\"p\">):</span>\n    <span class=\"k\">class</span> <span class=\"nc\">Article_Status</span><span
  class=\"p\">(</span><span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">TextChoices</span><span class=\"p\">):</span>\n        <span class=\"n\">DRAFT</span>
  <span class=\"o\">=</span> <span class=\"p\">(</span>\n            <span class=\"s2\">&quot;DRAFT&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"n\">_</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;Draft&quot;</span><span class=\"p\">),</span>\n        <span
  class=\"p\">)</span>\n        <span class=\"n\">PUBLISHED</span> <span class=\"o\">=</span>
  <span class=\"p\">(</span>\n            <span class=\"s2\">&quot;PUBLISHED&quot;</span><span
  class=\"p\">,</span>\n            <span class=\"n\">_</span><span class=\"p\">(</span><span
  class=\"s2\">&quot;Published&quot;</span><span class=\"p\">),</span>\n        <span
  class=\"p\">)</span>\n\n    <span class=\"n\">title</span> <span class=\"o\">=</span>
  <span class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">CharField</span><span
  class=\"p\">(</span><span class=\"n\">max_length</span><span class=\"o\">=</span><span
  class=\"mi\">128</span><span class=\"p\">)</span>\n    <span class=\"n\">description</span>
  <span class=\"o\">=</span> <span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">CharField</span><span class=\"p\">(</span><span class=\"n\">max_length</span><span
  class=\"o\">=</span><span class=\"mi\">256</span><span class=\"p\">)</span>\n    <span
  class=\"n\">content</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">TextField</span><span class=\"p\">(</span><span
  class=\"n\">default</span><span class=\"o\">=</span><span class=\"s2\">&quot;&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">null</span><span class=\"o\">=</span><span
  class=\"kc\">False</span><span class=\"p\">,</span> <span class=\"n\">blank</span><span
  class=\"o\">=</span><span class=\"kc\">False</span><span class=\"p\">)</span>\n
  \   <span class=\"n\">status</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">CharField</span><span class=\"p\">(</span>\n
  \       <span class=\"n\">max_length</span><span class=\"o\">=</span><span class=\"mi\">16</span><span
  class=\"p\">,</span>\n        <span class=\"n\">choices</span><span class=\"o\">=</span><span
  class=\"n\">Article_Status</span><span class=\"o\">.</span><span class=\"n\">choices</span><span
  class=\"p\">,</span>\n        <span class=\"n\">default</span><span class=\"o\">=</span><span
  class=\"n\">Article_Status</span><span class=\"o\">.</span><span class=\"n\">DRAFT</span><span
  class=\"p\">,</span>\n    <span class=\"p\">)</span>\n    <span class=\"n\">blog</span>
  <span class=\"o\">=</span> <span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">ForeignKey</span><span class=\"p\">(</span><span class=\"n\">Blog</span><span
  class=\"p\">,</span> <span class=\"n\">on_delete</span><span class=\"o\">=</span><span
  class=\"n\">models</span><span class=\"o\">.</span><span class=\"n\">CASCADE</span><span
  class=\"p\">,</span> <span class=\"n\">null</span><span class=\"o\">=</span><span
  class=\"kc\">True</span><span class=\"p\">,</span> <span class=\"n\">blank</span><span
  class=\"o\">=</span><span class=\"kc\">True</span><span class=\"p\">)</span>\n    <span
  class=\"n\">author</span> <span class=\"o\">=</span> <span class=\"n\">models</span><span
  class=\"o\">.</span><span class=\"n\">ForeignKey</span><span class=\"p\">(</span><span
  class=\"n\">User</span><span class=\"p\">,</span> <span class=\"n\">on_delete</span><span
  class=\"o\">=</span><span class=\"n\">models</span><span class=\"o\">.</span><span
  class=\"n\">CASCADE</span><span class=\"p\">,</span> <span class=\"n\">related_name</span><span
  class=\"o\">=</span><span class=\"s2\">&quot;author&quot;</span><span class=\"p\">)</span>\n\n
  \   <span class=\"k\">def</span> <span class=\"fm\">__str__</span><span class=\"p\">(</span><span
  class=\"bp\">self</span><span class=\"p\">):</span>\n        <span class=\"k\">return</span>
  <span class=\"bp\">self</span><span class=\"o\">.</span><span class=\"n\">title</span>\n\n
  \   <span class=\"k\">def</span> <span class=\"nf\">get_absolute_url</span><span
  class=\"p\">(</span><span class=\"bp\">self</span><span class=\"p\">):</span>      \n
  \       <span class=\"k\">return</span> <span class=\"n\">reverse</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;articles:article-detail&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">args</span><span class=\"o\">=</span><span class=\"p\">[</span><span
  class=\"nb\">str</span><span class=\"p\">(</span><span class=\"bp\">self</span><span
  class=\"o\">.</span><span class=\"n\">id</span><span class=\"p\">)])</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the below snippet, we have the forms defined in the article application for creating
  or updating of article instance.  We will be using model forms as our form data
  should contain fields related to a model in this case the <code>Article</code> model.
  So when we inherit the <code>forms.ModelForm</code> in our custom <code>ArticleForm</code>
  we simply need to specify the model and it will take in all the attributes of that
  model by default, but if we specify the <code>fields</code> or <code>exclude</code>
  tuples, it will include only or exclude only the provided list of attributes from
  the model.</p>\n<p>We have also added the widgets for the form fields which will
  allow us to customize the way the fields in the template/form will render. We can
  specify the HTML attributes like <code>width</code>, <code>height</code>, <code>style</code>,
  etc.</p>\n<h3>Article Form</h3>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># article/forms.py</span>\n\n\n<span class=\"kn\">from</span> <span
  class=\"nn\">django</span> <span class=\"kn\">import</span> <span class=\"n\">forms</span>\n<span
  class=\"kn\">from</span> <span class=\"nn\">.models</span> <span class=\"kn\">import</span>
  <span class=\"n\">Article</span>\n\n\n<span class=\"k\">class</span> <span class=\"nc\">ArticleForm</span><span
  class=\"p\">(</span><span class=\"n\">forms</span><span class=\"o\">.</span><span
  class=\"n\">ModelForm</span><span class=\"p\">):</span>\n    <span class=\"k\">class</span>
  <span class=\"nc\">Meta</span><span class=\"p\">:</span>\n        <span class=\"n\">model</span>
  <span class=\"o\">=</span> <span class=\"n\">Article</span>\n        <span class=\"n\">exclude</span>
  <span class=\"o\">=</span> <span class=\"p\">(</span>\n            <span class=\"s2\">&quot;created&quot;</span><span
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
  class=\"s2\">&quot;description&quot;</span><span class=\"p\">:</span> <span class=\"n\">forms</span><span
  class=\"o\">.</span><span class=\"n\">TextInput</span><span class=\"p\">(</span>\n
  \               <span class=\"n\">attrs</span><span class=\"o\">=</span><span class=\"p\">{</span>\n
  \                   <span class=\"s2\">&quot;class&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;form-control&quot;</span><span class=\"p\">,</span>\n                    <span
  class=\"s2\">&quot;style&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;max-width:
  900px; &quot;</span><span class=\"p\">,</span>\n                    <span class=\"s2\">&quot;placeholder&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;Description&quot;</span><span class=\"p\">,</span>\n
  \               <span class=\"p\">}</span>\n            <span class=\"p\">),</span>\n
  \           <span class=\"s2\">&quot;content&quot;</span><span class=\"p\">:</span>
  <span class=\"n\">forms</span><span class=\"o\">.</span><span class=\"n\">Textarea</span><span
  class=\"p\">(</span>\n                <span class=\"n\">attrs</span><span class=\"o\">=</span><span
  class=\"p\">{</span>\n                    <span class=\"s2\">&quot;class&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;form-control post-body&quot;</span><span
  class=\"p\">,</span>\n                    <span class=\"s2\">&quot;id&quot;</span><span
  class=\"p\">:</span> <span class=\"s2\">&quot;text-content&quot;</span><span class=\"p\">,</span>\n
  \                   <span class=\"s2\">&quot;style&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;max-width:900px;&quot;</span><span class=\"p\">,</span>\n
  \                   <span class=\"s2\">&quot;hx-post&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;/article/meta/&quot;</span><span class=\"p\">,</span>\n
  \                   <span class=\"s2\">&quot;placeholder&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;Content&quot;</span><span class=\"p\">,</span>\n                <span
  class=\"p\">}</span>\n            <span class=\"p\">),</span>\n            <span
  class=\"s2\">&quot;blog&quot;</span><span class=\"p\">:</span> <span class=\"n\">forms</span><span
  class=\"o\">.</span><span class=\"n\">Select</span><span class=\"p\">(</span>\n
  \               <span class=\"n\">attrs</span><span class=\"o\">=</span><span class=\"p\">{</span>\n
  \                   <span class=\"s2\">&quot;class&quot;</span><span class=\"p\">:</span>
  <span class=\"s2\">&quot;form-control&quot;</span><span class=\"p\">,</span>\n                    <span
  class=\"s2\">&quot;placeholder&quot;</span><span class=\"p\">:</span> <span class=\"s2\">&quot;Blog
  Publication&quot;</span><span class=\"p\">,</span>\n                <span class=\"p\">}</span>\n
  \           <span class=\"p\">),</span>\n        <span class=\"p\">}</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  these are my models and form files in the article app. Using htmx and without any
  javascript I want to update the form so that it picks up the front matter in the
  content field which is a text area and fills the title, description other attributes
  automatically for me.</p>\n<p>This can be done in a lot of ways, but I will be sharing
  one of the ways that I recently used in my blog project. This process involves creating
  a class-based view and adding a <code>POST</code> method that won't post any data
  to the backend but will send necessary data to the view.</p>\n<p>Let's see the process
  before diving into any of the code:</p>\n<h2>Gist of the Process</h2>\n<ul>\n<li>Attach
  a <code>hx-post</code> attribute to the form field for sending the request to a
  view</li>\n<li>When the request is sent, the data is loaded with <code>request.POST</code>,
  it is cleaned and converted in python-readable format with json.</li>\n<li>Once
  we have the data, we try to use the <code>frontmatter.loads</code> function that
  will load the content and if we have a frontmatter in the text, it will load it
  as a <code>frontmatter.POST</code> object.</li>\n<li>We will extract <code>title</code>,
  <code>description</code>, and other data fields from the object.</li>\n<li>We will
  initialize a Form instance of Article, with the initial data values as the extracted
  data from the frontmatter.</li>\n<li>Now we have two options:\n<ul>\n<li>If the
  article instance already exists i.e. we are updating the article</li>\n<li>Else
  we are creating a new article</li>\n</ul>\n</li>\n</ul>\n<p>Accordingly, we will
  load the form in the respective templates i.e. <code>update.html</code> for updating
  the existing articles and <code>article-form.html</code> for a new article.</p>\n<h2>Adding
  HTMX Magic</h2>\n<p>If you'd have seen we have a <code>hx-post</code> attribute
  in the <code>article/forms.py</code> file, the <code>content</code> widget has a
  <code>hx-post</code> attribute which sends a post request to the <code>article/meta/</code>
  URL route. This route we will bind to the <code>ArticleMetaView</code> which we
  will define in a few moments. This is usually sent once we change a certain text
  in the content field, so we can modify it as per your requirement with <code>hx-trigger</code>
  that can enable us to specify the trigger event or the type of trigger we want.
  Further, you can read from the <a href=\"https://htmx.org/docs/#trigger-modifiers\">htmx
  docs</a> about these triggers and other attributes.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"n\">app_name</span> <span class=\"o\">=</span> <span class=\"s2\">&quot;articles&quot;</span>\n\n<span
  class=\"n\">urlpatterns</span> <span class=\"o\">=</span> <span class=\"p\">[</span>\n
  \   <span class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">ArticleCreateView</span><span class=\"o\">.</span><span class=\"n\">as_view</span><span
  class=\"p\">(),</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;article-create&quot;</span><span class=\"p\">),</span>\n    <span
  class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;&lt;int:pk&gt;/&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">ArticleDetailView</span><span class=\"o\">.</span><span class=\"n\">as_view</span><span
  class=\"p\">(),</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;article-detail&quot;</span><span class=\"p\">),</span>\n    <span
  class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;delete/&lt;int:pk&gt;/&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">ArticleDetailView</span><span class=\"o\">.</span><span class=\"n\">as_view</span><span
  class=\"p\">(),</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;article-delete&quot;</span><span class=\"p\">),</span>\n    <span
  class=\"n\">path</span><span class=\"p\">(</span><span class=\"s2\">&quot;edit/&lt;int:pk&gt;&quot;</span><span
  class=\"p\">,</span> <span class=\"n\">views</span><span class=\"o\">.</span><span
  class=\"n\">ArticleDetailView</span><span class=\"o\">.</span><span class=\"n\">as_view</span><span
  class=\"p\">(),</span> <span class=\"n\">name</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;article-update&quot;</span><span class=\"p\">),</span>\n\n    <span
  class=\"c1\"># the new view that we will create</span>\n    <span class=\"n\">path</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;meta/&quot;</span><span class=\"p\">,</span>
  <span class=\"n\">views</span><span class=\"o\">.</span><span class=\"n\">ArticleMetaView</span><span
  class=\"o\">.</span><span class=\"n\">as_view</span><span class=\"p\">(),</span>
  <span class=\"n\">name</span><span class=\"o\">=</span><span class=\"s2\">&quot;article-meta&quot;</span><span
  class=\"p\">),</span>\n<span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<h2>Capture
  Frontmatter Meta-data View</h2>\n<p>Along with the Create, Detail/List, Update,
  Delete View, I will create a separate class called <code>ArticleMetaView</code>
  that will fetch the form fields and render the templates again but this time it
  will fill in the frontmatter meta-data in the fields if the content is parsed with
  the relvant frontmatter.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"c1\"># articles/view.py</span>\n\n<span class=\"k\">class</span> <span class=\"nc\">ArticleMetaView</span><span
  class=\"p\">(</span><span class=\"n\">View</span><span class=\"p\">):</span>\n    <span
  class=\"n\">model</span> <span class=\"o\">=</span> <span class=\"n\">Article</span>\n\n
  \   <span class=\"k\">def</span> <span class=\"nf\">post</span><span class=\"p\">(</span><span
  class=\"bp\">self</span><span class=\"p\">,</span> <span class=\"n\">request</span><span
  class=\"p\">,</span> <span class=\"o\">*</span><span class=\"n\">args</span><span
  class=\"p\">,</span> <span class=\"o\">**</span><span class=\"n\">kwargs</span><span
  class=\"p\">):</span>\n        \n        <span class=\"n\">data</span> <span class=\"o\">=</span>
  <span class=\"n\">json</span><span class=\"o\">.</span><span class=\"n\">loads</span><span
  class=\"p\">(</span><span class=\"n\">json</span><span class=\"o\">.</span><span
  class=\"n\">dumps</span><span class=\"p\">(</span><span class=\"nb\">dict</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">POST</span><span class=\"p\">)))</span>\n        <span class=\"n\">loaded_frontmatter</span>
  <span class=\"o\">=</span> <span class=\"n\">frontmatter</span><span class=\"o\">.</span><span
  class=\"n\">loads</span><span class=\"p\">(</span><span class=\"n\">data</span><span
  class=\"p\">[</span><span class=\"s1\">&#39;content&#39;</span><span class=\"p\">][</span><span
  class=\"mi\">0</span><span class=\"p\">])</span>\n\n       <span class=\"c1\">#
  frontmatter has keys i.e. attributes like title, description, etc.</span>\n        <span
  class=\"k\">if</span> <span class=\"nb\">dict</span><span class=\"p\">(</span><span
  class=\"n\">loaded_frontmatter</span><span class=\"p\">):</span>\n            <span
  class=\"n\">article_title</span> <span class=\"o\">=</span> <span class=\"n\">loaded_frontmatter</span><span
  class=\"p\">[</span><span class=\"s1\">&#39;title&#39;</span><span class=\"p\">]</span>\n
  \           <span class=\"n\">article_description</span> <span class=\"o\">=</span>
  <span class=\"n\">loaded_frontmatter</span><span class=\"p\">[</span><span class=\"s1\">&#39;description&#39;</span><span
  class=\"p\">]</span>\n            <span class=\"n\">form</span> <span class=\"o\">=</span>
  <span class=\"n\">ArticleForm</span><span class=\"p\">(</span><span class=\"n\">initial</span><span
  class=\"o\">=</span><span class=\"p\">{</span><span class=\"s1\">&#39;title&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">article_title</span><span class=\"p\">,</span>
  \n            <span class=\"s1\">&#39;description&#39;</span><span class=\"p\">:</span>
  <span class=\"n\">article_description</span><span class=\"p\">,</span> <span class=\"s1\">&#39;content&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">loaded_frontmatter</span><span class=\"o\">.</span><span
  class=\"n\">content</span><span class=\"p\">})</span>\n            <span class=\"n\">context</span>
  <span class=\"o\">=</span> <span class=\"p\">{</span><span class=\"s1\">&#39;form&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">form</span><span class=\"p\">}</span>\n            <span
  class=\"n\">article_list</span> <span class=\"o\">=</span> <span class=\"n\">Article</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">filter</span><span class=\"p\">(</span><span class=\"n\">title</span><span
  class=\"o\">=</span><span class=\"n\">article_title</span><span class=\"p\">)</span>\n
  \           <span class=\"k\">if</span> <span class=\"n\">article_list</span><span
  class=\"p\">:</span>\n                <span class=\"n\">article</span> <span class=\"o\">=</span>
  <span class=\"n\">article_list</span><span class=\"o\">.</span><span class=\"n\">last</span><span
  class=\"p\">()</span>\n                <span class=\"n\">context</span><span class=\"p\">[</span><span
  class=\"s1\">&#39;article&#39;</span><span class=\"p\">]</span> <span class=\"o\">=</span>
  <span class=\"n\">article</span>\n                <span class=\"k\">return</span>
  <span class=\"n\">render</span><span class=\"p\">(</span><span class=\"n\">request</span><span
  class=\"p\">,</span> <span class=\"s1\">&#39;articles/edit_article.html&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">context</span><span class=\"p\">)</span>\n
  \           <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;articles/article_form.html&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">context</span><span class=\"p\">)</span>\n\n        <span class=\"n\">article_list</span>
  <span class=\"o\">=</span> <span class=\"n\">Article</span><span class=\"o\">.</span><span
  class=\"n\">objects</span><span class=\"o\">.</span><span class=\"n\">filter</span><span
  class=\"p\">(</span><span class=\"n\">title</span><span class=\"o\">=</span><span
  class=\"n\">data</span><span class=\"p\">[</span><span class=\"s1\">&#39;title&#39;</span><span
  class=\"p\">][</span><span class=\"mi\">0</span><span class=\"p\">])</span>\n       \n
  \      <span class=\"c1\"># if the article title has been already taken i.e. we
  are updating an article</span>\n\n        <span class=\"k\">if</span> <span class=\"n\">article_list</span><span
  class=\"p\">:</span>\n            <span class=\"n\">article</span> <span class=\"o\">=</span>
  <span class=\"n\">article_list</span><span class=\"o\">.</span><span class=\"n\">last</span><span
  class=\"p\">()</span>\n            <span class=\"n\">form</span> <span class=\"o\">=</span>
  <span class=\"n\">ArticleForm</span><span class=\"p\">(</span><span class=\"n\">data</span><span
  class=\"o\">=</span><span class=\"n\">request</span><span class=\"o\">.</span><span
  class=\"n\">POST</span><span class=\"p\">)</span>\n            <span class=\"n\">context</span>
  <span class=\"o\">=</span> <span class=\"p\">{</span><span class=\"s1\">&#39;form&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">form</span><span class=\"p\">}</span>\n            <span
  class=\"n\">context</span><span class=\"p\">[</span><span class=\"s1\">&#39;article&#39;</span><span
  class=\"p\">]</span> <span class=\"o\">=</span> <span class=\"n\">article</span>\n
  \           <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;articles/edit_article.html&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">context</span><span class=\"p\">)</span>\n\n        <span class=\"n\">form</span>
  <span class=\"o\">=</span> <span class=\"n\">ArticleForm</span><span class=\"p\">(</span><span
  class=\"n\">data</span><span class=\"o\">=</span><span class=\"n\">request</span><span
  class=\"o\">.</span><span class=\"n\">POST</span><span class=\"p\">)</span>\n        <span
  class=\"n\">context</span> <span class=\"o\">=</span> <span class=\"p\">{</span><span
  class=\"s1\">&#39;form&#39;</span><span class=\"p\">:</span> <span class=\"n\">form</span><span
  class=\"p\">}</span>\n        <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;articles/article_form.html&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above <code>ArticleMetaView</code> we have created a <code>post</code> method
  as we want to get hold of the content from the form. So, we start by extracting
  and converting the <code>request.body</code> data into an appropriate type for easily
  working with python. So, the <code>request.body</code> will contain the data like
  <code>csrf_token</code>, <code>form_data</code>, etc. received from the frontend
  template. We store the received data as <code>data</code> and now from this data,
  we can load the content field which will have the content information.</p>\n<p>Firstly
  we will extract the <code>request.body</code> which will contain the data from the
  form as we have made a <code>POST</code> request to this endpoint. For doing that
  we need to parse the content in a apropriate format such that it is python friendly.
  So we wrap the <code>request.body</code> into json format and then decode it back
  into the json string. This will give us the dict of the request data.</p>\n<pre
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
  class=\"n\">data</span> <span class=\"o\">=</span> <span class=\"n\">json</span><span
  class=\"o\">.</span><span class=\"n\">loads</span><span class=\"p\">(</span><span
  class=\"n\">json</span><span class=\"o\">.</span><span class=\"n\">dumps</span><span
  class=\"p\">(</span><span class=\"nb\">dict</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"o\">.</span><span class=\"n\">POST</span><span
  class=\"p\">)))</span>\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div
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
  \       \n\n<div class=\"highlight\"><pre><span></span>{&#39;csrfmiddlewaretoken&#39;:
  [&#39;bSYJxD39XH509tD1tZGd0WU21PUaKaLeqjjGbyzRvLXF4P8iIxb5l0fmTWVFjELQ&#39;], &#39;title&#39;:
  [&#39;test2&#39;], &#39;description&#39;: [&#39;test&#39;], &#39;content&#39;: [&#39;test
  something&#39;], &#39;status&#39;: [&#39;DRAFT&#39;], &#39;blog&#39;: [&#39;&#39;]}\n</pre></div>\n\n</pre>\n\n<p>So,
  this will grab the request data as a dict, we can then extract the data from this
  as it has data from the Form fields. We are interested in the content field in the
  Form, so we can get it by specifying the key <code>content</code> from the extracted
  data. But as we can see the data doesn't contain the actual data instead it is wrapped
  in a list i.e. <code>['test something']</code>, so we will have to index it and
  then fetch it.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"n\">content_string</span> <span class=\"o\">=</span> <span class=\"n\">data</span><span
  class=\"p\">[</span><span class=\"s1\">&#39;content&#39;</span><span class=\"p\">][</span><span
  class=\"mi\">0</span><span class=\"p\">]</span>\n</pre></div>\n\n</pre>\n\n<p>This
  will give us the exact content field as a string. So, we can now move into extracting
  the frontmatter from the fields.</p>\n<p>Now, we can use the <a href=\"https://python-frontmatter.readthedocs.io/en/latest/index.html\">frontmatter</a>
  library to parse the content into the <a href=\"https://python-frontmatter.readthedocs.io/en/latest/api.html#frontmatter.loads\">loads</a>
  funciton and extract the frontmatter if it is present in the content field. The
  frontmatter library has a <code>loads</code> function which takes in a string and
  can give out a <a href=\"https://python-frontmatter.readthedocs.io/en/latest/api.html#post-objects\">frontmatter.Post</a>
  object. The loads function is differnet from the <a href=\"https://python-frontmatter.readthedocs.io/en/latest/api.html#frontmatter.load\">load</a>
  function as the load frunciton is for reading data from a stream of bytes i.e. a
  file or othe related byte object. The differnece is subtle but it took a read at
  the <a href=\"https://python-frontmatter.readthedocs.io/en/latest/api.html#module-frontmatter\">documentation</a>.</p>\n<pre
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
  class=\"n\">post</span> <span class=\"o\">=</span> <span class=\"n\">data</span><span
  class=\"p\">[</span><span class=\"s1\">&#39;content&#39;</span><span class=\"p\">][</span><span
  class=\"mi\">0</span><span class=\"p\">]</span>\n<span class=\"n\">loaded_frontmatter</span>
  <span class=\"o\">=</span> <span class=\"n\">frontmatter</span><span class=\"o\">.</span><span
  class=\"n\">loads</span><span class=\"p\">(</span><span class=\"n\">post</span><span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>This wil load the content and
  give us a <code>frontmatter.Post</code> as said earlier. This will contain a dict
  with all the frontmatter if it has any and will by default parse the non-frontmatter
  data i.e. the remaining text into the <code>content</code> key. We need a chack
  if the Form field had any fronmatter this can be checked by the <code>dict(loaded_frontmatter)</code>
  which will return None if it cannot load the frontmatter.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"n\">loaded_frontmatter</span> <span class=\"o\">=</span> <span class=\"n\">frontmatter</span><span
  class=\"o\">.</span><span class=\"n\">loads</span><span class=\"p\">(</span><span
  class=\"n\">data</span><span class=\"p\">[</span><span class=\"s1\">&#39;content&#39;</span><span
  class=\"p\">][</span><span class=\"mi\">0</span><span class=\"p\">])</span>\n<span
  class=\"k\">if</span> <span class=\"nb\">dict</span><span class=\"p\">(</span><span
  class=\"n\">loaded_frontmatter</span><span class=\"p\">):</span>\n  <span class=\"nb\">print</span><span
  class=\"p\">(</span><span class=\"n\">loaded_frontmatter</span><span class=\"o\">.</span><span
  class=\"n\">keys</span><span class=\"p\">())</span>\n</pre></div>\n\n</pre>\n\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>dict_keys([&#39;templateKey&#39;,
  &#39;title&#39;, &#39;description&#39;, &#39;date&#39;, &#39;status&#39;])\n</pre></div>\n\n</pre>\n\n<p>So
  once we have the frontmatter loaded we can get specific keys from it and initialize
  the form vaues to them. But we have made clear distictions that we want to perform
  a specific task if we have frontmatter keys in the content field of the Form else
  we can do something else.</p>\n<p>First let's handle the loading of the frontmatter
  into the form. For doing that we will get all the required attributes from the frontmatter
  like <code>title</code>, <code>description</code>, <code>content</code>, etc which
  can be accessed normally as we extract the value from a key in a dict.</p>\n<p>Once
  we have got those keys, we can start filling in the Form data with initial values.
  The <a href=\"https://docs.djangoproject.com/en/4.0/topics/forms/modelforms/\">Django
  Model form</a> takes in a parameter like <a href=\"https://docs.djangoproject.com/en/4.0/topics/forms/modelforms/#providing-initial-values\">initial</a>
  which can be a dict of the fiields along with the value that can be used for rendering
  the form initially when we load the template.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"n\">article_title</span> <span class=\"o\">=</span> <span class=\"n\">loaded_frontmatter</span><span
  class=\"p\">[</span><span class=\"s1\">&#39;title&#39;</span><span class=\"p\">]</span>\n<span
  class=\"n\">article_description</span> <span class=\"o\">=</span> <span class=\"n\">loaded_frontmatter</span><span
  class=\"p\">[</span><span class=\"s1\">&#39;description&#39;</span><span class=\"p\">]</span>\n\n<span
  class=\"n\">form</span> <span class=\"o\">=</span> <span class=\"n\">ArticleForm</span><span
  class=\"p\">(</span><span class=\"n\">initial</span><span class=\"o\">=</span><span
  class=\"p\">{</span><span class=\"s1\">&#39;title&#39;</span><span class=\"p\">:</span>
  <span class=\"n\">article_title</span><span class=\"p\">,</span> <span class=\"s1\">&#39;description&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">article_description</span><span class=\"p\">,</span>
  <span class=\"s1\">&#39;content&#39;</span><span class=\"p\">:</span> <span class=\"n\">loaded_frontmatter</span><span
  class=\"o\">.</span><span class=\"n\">content</span><span class=\"p\">})</span>\n</pre></div>\n\n</pre>\n\n<p>This
  will take in a <code>ArticleForm</code> and fill the initial values like <code>title</code>,
  <code>description</code>, etc which we have provided in the dict with the values.
  Now, we need to parse this form in the current template or re-render the template.
  But before that, we need to parse this context into the template. We will create
  a dict with <code>form</code> as the key which can be used to render in the template.</p>\n<p>Also,
  we have a two ways here, either the user is creating a new article or it is updating
  a existing article. We need to make sure that we preserve the initial fields in
  the form as we are updating the existing article. So, we can filter the article
  objects as per the title of the current title and then if we find a article with
  that title, we will parse the context with that article object.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"n\">article_list</span> <span class=\"o\">=</span> <span class=\"n\">Article</span><span
  class=\"o\">.</span><span class=\"n\">objects</span><span class=\"o\">.</span><span
  class=\"n\">filter</span><span class=\"p\">(</span><span class=\"n\">title</span><span
  class=\"o\">=</span><span class=\"n\">article_title</span><span class=\"p\">)</span>\n<span
  class=\"k\">if</span> <span class=\"n\">article_list</span><span class=\"p\">:</span>\n
  \   <span class=\"n\">article</span> <span class=\"o\">=</span> <span class=\"n\">article_list</span><span
  class=\"o\">.</span><span class=\"n\">last</span><span class=\"p\">()</span>\n    <span
  class=\"n\">context</span><span class=\"p\">{</span>\n      <span class=\"s1\">&#39;form&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">form</span><span class=\"p\">,</span>\n      <span
  class=\"s1\">&#39;article&#39;</span><span class=\"p\">:</span> <span class=\"n\">article</span>\n
  \   <span class=\"p\">}</span>\n    <span class=\"k\">return</span> <span class=\"n\">render</span><span
  class=\"p\">(</span><span class=\"n\">request</span><span class=\"p\">,</span> <span
  class=\"s1\">&#39;articles/edit_article.html&#39;</span><span class=\"p\">,</span>
  <span class=\"n\">context</span><span class=\"p\">)</span>\n<span class=\"n\">context</span>
  <span class=\"o\">=</span> <span class=\"p\">{</span><span class=\"s1\">&#39;form&#39;</span><span
  class=\"p\">:</span> <span class=\"n\">form</span><span class=\"p\">}</span>\n<span
  class=\"k\">return</span> <span class=\"n\">render</span><span class=\"p\">(</span><span
  class=\"n\">request</span><span class=\"p\">,</span> <span class=\"s1\">&#39;articles/article_form.html&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">context</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>Now,
  we have form data along with the article instance used for rendering the form with
  appropriate content. So, this will work for editing an already existing article.
  For a new article, we have to simply parse the form to the template and it will
  render the title picked from the fotnmatter or leave it empty.</p>\n<p>Similarly,
  for the article with no frontmatter we will iterate over the article and if the
  article's title already exist, we will render the article data with the form else
  render the form with the parsed title and other meta-data in the form.</p>\n<video
  width=\"800\" height=\"450\" controls>\n  <source src=\"https://res.cloudinary.com/techstructive-blog/video/upload/v1659370006/blog-media/frontmatter-load-htmx.mp4\"
  type=\"video/mp4\">\n</video>\n<p>So that is how we render the form data with frontmatter
  into appropriate meta-data in the form. We have used Django forms and make use of
  HTMX for the dynamic updation of form.</p>\n<div class='prevnext'>\n    <style type='text/css'>\n
  \   :root {\n      --prevnext-color-text: #eefbfe;\n      --prevnext-color-angle:
  #ff6600;\n      --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"light\"]
  {\n      --prevnext-color-text: #1f2022;\n      --prevnext-color-angle: #ffeb00;\n
  \     --prevnext-subtitle-brightness: 3;\n    }\n    [data-theme=\"dark\"] {\n      --prevnext-color-text:
  #eefbfe;\n      --prevnext-color-angle: #ff6600;\n      --prevnext-subtitle-brightness:
  3;\n    }\n    .prevnext {\n      display: flex;\n      flex-direction: row;\n      justify-content:
  space-around;\n      align-items: flex-start;\n    }\n    .prevnext a {\n      display:
  flex;\n      align-items: center;\n      width: 100%;\n      text-decoration: none;\n
  \   }\n    a.next {\n      justify-content: flex-end;\n    }\n    .prevnext a:hover
  {\n      background: #00000006;\n    }\n    .prevnext-subtitle {\n      color: var(--prevnext-color-text);\n
  \     filter: brightness(var(--prevnext-subtitle-brightness));\n      font-size:
  .8rem;\n    }\n    .prevnext-title {\n      color: var(--prevnext-color-text);\n
  \     font-size: 1rem;\n    }\n    .prevnext-text {\n      max-width: 30vw;\n    }\n
  \   </style>\n<pre><code>&lt;a class='prev' href='/libsql-query-remote-db'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;LibSQL:
  Query a remote Turso database with cURL&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a
  class='next' href='/python-chain-assignment-gotcha'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Gotcha
  with Chained Assignment in Python&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot;
  height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
long_description: I have an Article Form where I load my post into it directly, it
  might have frontmatter. So what I wish to achieve is when I paste in the content,
  the frontmatter should be picked up and it should render the form fields like  To
  do that, we will requ
now: 2025-05-01 18:11:33.315189
path: blog/tils/2022-08-01-Django-Frontmatter-forms.md
slug: django-form-load-frontmatter
status: published-til
tags:
- django
- htmx
- python
templateKey: til
title: 'Django Blog DevLog: Load Frontmatter data into Template/Model Form Fields'
today: 2025-05-01
---

## Introduction

I have an Article Form where I load my post into it directly, it might have frontmatter. So what I wish to achieve is when I paste in the content, the frontmatter should be picked up and it should render the form fields like `title`, `description`, and then also remove the frontmatter from the content.

To do that, we will require a model to work with and a form based on that model. We will exclude a few fields from that model so as to process these attributes on the server side. I am working on my Blog project which is a simple Django application.  You can get the source code for the project on the [GitHub repository](https://github.com/mr-destructive/techstructive-blog/).

## Django Project Context

The techstructive-blog is a django project, which has a couple of applications currently, not in a good situation. There are apps like `article`, `blog`, and `user`. This project has templates and static folder in the base directory. The project is deployed on [railway](https://www.railway.app) this is an always under development project, you can check out the [techstructive-blog](https://django-blog.up.railway.app). We can get the bits and pieces of the project details required for understanding what I want to do with the following sections.

### Article Model

We have an `Article` model with attributes like `title`, `description`,  `author` as a Foreign Key to the user model, and a few other attributes which is not related to what we are trying to achieve right now or at least don't require an explanation. We have a model method called `get_absolute_url` for getting a URL in order to redirect the client when the model instance is created or updated from the backend. You can definitely check out the details of these small components or templates in the project repository. 

```python
# articles/models.py


class Article(TimeStampedModel):
    class Article_Status(models.TextChoices):
        DRAFT = (
            "DRAFT",
            _("Draft"),
        )
        PUBLISHED = (
            "PUBLISHED",
            _("Published"),
        )

    title = models.CharField(max_length=128)
    description = models.CharField(max_length=256)
    content = models.TextField(default="", null=False, blank=False)
    status = models.CharField(
        max_length=16,
        choices=Article_Status.choices,
        default=Article_Status.DRAFT,
    )
    blog = models.ForeignKey(Blog, on_delete=models.CASCADE, null=True, blank=True)
    author = models.ForeignKey(User, on_delete=models.CASCADE, related_name="author")

    def __str__(self):
        return self.title

    def get_absolute_url(self):      
        return reverse('articles:article-detail', args=[str(self.id)])
```

In the below snippet, we have the forms defined in the article application for creating or updating of article instance.  We will be using model forms as our form data should contain fields related to a model in this case the `Article` model. So when we inherit the `forms.ModelForm` in our custom `ArticleForm` we simply need to specify the model and it will take in all the attributes of that model by default, but if we specify the `fields` or `exclude` tuples, it will include only or exclude only the provided list of attributes from the model. 

We have also added the widgets for the form fields which will allow us to customize the way the fields in the template/form will render. We can specify the HTML attributes like `width`, `height`, `style`, etc.  

### Article Form

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
            "description": forms.TextInput(
                attrs={
                    "class": "form-control",
                    "style": "max-width: 900px; ",
                    "placeholder": "Description",
                }
            ),
            "content": forms.Textarea(
                attrs={
                    "class": "form-control post-body",
                    "id": "text-content",
                    "style": "max-width:900px;",
                    "hx-post": "/article/meta/",
                    "placeholder": "Content",
                }
            ),
            "blog": forms.Select(
                attrs={
                    "class": "form-control",
                    "placeholder": "Blog Publication",
                }
            ),
        }

```

So, these are my models and form files in the article app. Using htmx and without any javascript I want to update the form so that it picks up the front matter in the content field which is a text area and fills the title, description other attributes automatically for me. 

This can be done in a lot of ways, but I will be sharing one of the ways that I recently used in my blog project. This process involves creating a class-based view and adding a `POST` method that won't post any data to the backend but will send necessary data to the view.


Let's see the process before diving into any of the code:

## Gist of the Process

- Attach a `hx-post` attribute to the form field for sending the request to a view
- When the request is sent, the data is loaded with `request.POST`, it is cleaned and converted in python-readable format with json.
- Once we have the data, we try to use the `frontmatter.loads` function that will load the content and if we have a frontmatter in the text, it will load it as a `frontmatter.POST` object.
- We will extract `title`, `description`, and other data fields from the object.
- We will initialize a Form instance of Article, with the initial data values as the extracted data from the frontmatter.
- Now we have two options:
    - If the article instance already exists i.e. we are updating the article
   - Else we are creating a new article

Accordingly, we will load the form in the respective templates i.e. `update.html` for updating the existing articles and `article-form.html` for a new article.

## Adding HTMX Magic

If you'd have seen we have a `hx-post` attribute in the `article/forms.py` file, the `content` widget has a `hx-post` attribute which sends a post request to the `article/meta/` URL route. This route we will bind to the `ArticleMetaView` which we will define in a few moments. This is usually sent once we change a certain text in the content field, so we can modify it as per your requirement with `hx-trigger` that can enable us to specify the trigger event or the type of trigger we want. Further, you can read from the [htmx docs](https://htmx.org/docs/#trigger-modifiers) about these triggers and other attributes. 

```python
# article/urls.py

from django.urls import path
from . import views

app_name = "articles"

urlpatterns = [
    path("", views.ArticleCreateView.as_view(), name="article-create"),
    path("<int:pk>/", views.ArticleDetailView.as_view(), name="article-detail"),
    path("delete/<int:pk>/", views.ArticleDetailView.as_view(), name="article-delete"),
    path("edit/<int:pk>", views.ArticleDetailView.as_view(), name="article-update"),

    # the new view that we will create
    path("meta/", views.ArticleMetaView.as_view(), name="article-meta"),
]
```

## Capture Frontmatter Meta-data View 

Along with the Create, Detail/List, Update, Delete View, I will create a separate class called `ArticleMetaView` that will fetch the form fields and render the templates again but this time it will fill in the frontmatter meta-data in the fields if the content is parsed with the relvant frontmatter.

```python
# articles/view.py

class ArticleMetaView(View):
    model = Article

    def post(self, request, *args, **kwargs):
        
        data = json.loads(json.dumps(dict(request.POST)))
        loaded_frontmatter = frontmatter.loads(data['content'][0])

       # frontmatter has keys i.e. attributes like title, description, etc.
        if dict(loaded_frontmatter):
            article_title = loaded_frontmatter['title']
            article_description = loaded_frontmatter['description']
            form = ArticleForm(initial={'title': article_title, 
            'description': article_description, 'content': loaded_frontmatter.content})
            context = {'form': form}
            article_list = Article.objects.filter(title=article_title)
            if article_list:
                article = article_list.last()
                context['article'] = article
                return render(request, 'articles/edit_article.html', context)
            return render(request, 'articles/article_form.html', context)

        article_list = Article.objects.filter(title=data['title'][0])
       
       # if the article title has been already taken i.e. we are updating an article

        if article_list:
            article = article_list.last()
            form = ArticleForm(data=request.POST)
            context = {'form': form}
            context['article'] = article
            return render(request, 'articles/edit_article.html', context)

        form = ArticleForm(data=request.POST)
        context = {'form': form}
        return render(request, 'articles/article_form.html', context)

```

In the above `ArticleMetaView` we have created a `post` method as we want to get hold of the content from the form. So, we start by extracting and converting the `request.body` data into an appropriate type for easily working with python. So, the `request.body` will contain the data like `csrf_token`, `form_data`, etc. received from the frontend template. We store the received data as `data` and now from this data, we can load the content field which will have the content information.

Firstly we will extract the `request.body` which will contain the data from the form as we have made a `POST` request to this endpoint. For doing that we need to parse the content in a apropriate format such that it is python friendly. So we wrap the `request.body` into json format and then decode it back into the json string. This will give us the dict of the request data.

```python
data = json.loads(json.dumps(dict(request.POST)))
```

```
{'csrfmiddlewaretoken': ['bSYJxD39XH509tD1tZGd0WU21PUaKaLeqjjGbyzRvLXF4P8iIxb5l0fmTWVFjELQ'], 'title': ['test2'], 'description': ['test'], 'content': ['test something'], 'status': ['DRAFT'], 'blog': ['']}
```

So, this will grab the request data as a dict, we can then extract the data from this as it has data from the Form fields. We are interested in the content field in the Form, so we can get it by specifying the key `content` from the extracted data. But as we can see the data doesn't contain the actual data instead it is wrapped in a list i.e. `['test something']`, so we will have to index it and then fetch it.

```python
content_string = data['content'][0]
```

This will give us the exact content field as a string. So, we can now move into extracting the frontmatter from the fields. 

Now, we can use the [frontmatter](https://python-frontmatter.readthedocs.io/en/latest/index.html) library to parse the content into the [loads](https://python-frontmatter.readthedocs.io/en/latest/api.html#frontmatter.loads) funciton and extract the frontmatter if it is present in the content field. The frontmatter library has a `loads` function which takes in a string and can give out a [frontmatter.Post](https://python-frontmatter.readthedocs.io/en/latest/api.html#post-objects) object. The loads function is differnet from the [load](https://python-frontmatter.readthedocs.io/en/latest/api.html#frontmatter.load) function as the load frunciton is for reading data from a stream of bytes i.e. a file or othe related byte object. The differnece is subtle but it took a read at the [documentation](https://python-frontmatter.readthedocs.io/en/latest/api.html#module-frontmatter).

```python
post = data['content'][0]
loaded_frontmatter = frontmatter.loads(post)
```

This wil load the content and give us a `frontmatter.Post` as said earlier. This will contain a dict with all the frontmatter if it has any and will by default parse the non-frontmatter data i.e. the remaining text into the `content` key. We need a chack if the Form field had any fronmatter this can be checked by the `dict(loaded_frontmatter)` which will return None if it cannot load the frontmatter.

```python
loaded_frontmatter = frontmatter.loads(data['content'][0])
if dict(loaded_frontmatter):
  print(loaded_frontmatter.keys())
```

```
dict_keys(['templateKey', 'title', 'description', 'date', 'status'])
```

So once we have the frontmatter loaded we can get specific keys from it and initialize the form vaues to them. But we have made clear distictions that we want to perform a specific task if we have frontmatter keys in the content field of the Form else we can do something else.

First let's handle the loading of the frontmatter into the form. For doing that we will get all the required attributes from the frontmatter like `title`, `description`, `content`, etc which can be accessed normally as we extract the value from a key in a dict.

Once we have got those keys, we can start filling in the Form data with initial values. The [Django Model form](https://docs.djangoproject.com/en/4.0/topics/forms/modelforms/) takes in a parameter like [initial](https://docs.djangoproject.com/en/4.0/topics/forms/modelforms/#providing-initial-values) which can be a dict of the fiields along with the value that can be used for rendering the form initially when we load the template.

```python
article_title = loaded_frontmatter['title']
article_description = loaded_frontmatter['description']

form = ArticleForm(initial={'title': article_title, 'description': article_description, 'content': loaded_frontmatter.content})
```

This will take in a `ArticleForm` and fill the initial values like `title`, `description`, etc which we have provided in the dict with the values. Now, we need to parse this form in the current template or re-render the template. But before that, we need to parse this context into the template. We will create a dict with `form` as the key which can be used to render in the template.

Also, we have a two ways here, either the user is creating a new article or it is updating a existing article. We need to make sure that we preserve the initial fields in the form as we are updating the existing article. So, we can filter the article objects as per the title of the current title and then if we find a article with that title, we will parse the context with that article object.

```python
article_list = Article.objects.filter(title=article_title)
if article_list:
    article = article_list.last()
    context{
      'form': form,
      'article': article
    }
    return render(request, 'articles/edit_article.html', context)
context = {'form': form}
return render(request, 'articles/article_form.html', context)
```

Now, we have form data along with the article instance used for rendering the form with appropriate content. So, this will work for editing an already existing article. For a new article, we have to simply parse the form to the template and it will render the title picked from the fotnmatter or leave it empty.

Similarly, for the article with no frontmatter we will iterate over the article and if the article's title already exist, we will render the article data with the form else render the form with the parsed title and other meta-data in the form.


<video width="800" height="450" controls>
  <source src="https://res.cloudinary.com/techstructive-blog/video/upload/v1659370006/blog-media/frontmatter-load-htmx.mp4" type="video/mp4">
</video>

So that is how we render the form data with frontmatter into appropriate meta-data in the form. We have used Django forms and make use of HTMX for the dynamic updation of form.
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
    
    <a class='prev' href='/libsql-query-remote-db'>
    
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M13.5 8.25L9.75 12L13.5 15.75" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"> </path>
        </svg>
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>prev</p>
            <p class='prevnext-title'>LibSQL: Query a remote Turso database with cURL</p>
        </div>
    </a>
    
    <a class='next' href='/python-chain-assignment-gotcha'>
    
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>next</p>
            <p class='prevnext-title'>Gotcha with Chained Assignment in Python</p>
        </div>
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.5 15.75L14.25 12L10.5 8.25" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path>
        </svg>
    </a>
  </div>