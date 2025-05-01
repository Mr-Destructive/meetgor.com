---
article_html: "<h2>Introduction</h2>\n<p>Hello Developers! Want to listen to programming
  podcasts from a single place? Podevcast is the place you should be searching for.</p>\n<p>I
  am Meet Gor and I present this project as a submission to the Netlify x Hashnode
  Hackathon. Podevcast is a webpage(static) for listening to podcasts centered around
  developers and programming. Just pick your favorite one and start listening straight
  away. Let's dive into the making of Podevcast. Head on to <a href=\"https://podevcast.netlify.app/\">https://podevcast.netlify.app/</a>
  to check out the live app.</p>\n<h2>What is Podevcast</h2>\n<p>Podevcast is a web
  application or a static site that renders the top programming/development podcasts.
  You can listen to the top podcasts around the developer community from a single
  source.</p>\n<blockquote>\n<p>Listen to your favorite developer podcasts with Podevcast</p>\n</blockquote>\n<p>Podevcast
  is a static site generated using a script. There is a static site generator that
  is heavily done in Python and deployed to Netlify. You can simply listen to the
  podcasts on the web page or go to the canonical page of the podcast episode. From
  the canonical page, you can choose to hop to your chosen music player, but the default
  music player should be fine for casual listening. The core idea is to keep things
  in a single place for developer podcasts.</p>\n<p><a href=\"https://podevcast.netlify.app/\">Podevcast</a></p>\n<p><a
  href=\"https://github.com/Mr-Destructive/podevcast\">Source Code</a></p>\n<h2>Preview</h2>\n<p>Podevcast
  has popular developer podcasts like <code>Command Line Heroes</code>, <code>The
  Python Podcast</code>, <code>The freeCodeCamp Podcast</code>, and many others to
  choose from. You can go into categories for looking at a specific podcast.</p>\n<h3>Application
  Demonstration</h3>\n<p>Here's a small demonstration of the Podevcast application.</p>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1645200224921/GC8gmxUzX.gif\"
  alt=\"Podevcast Preview gif\" /></p>\n<p>Podevcast has multiple pages like:</p>\n<ol>\n<li><a
  href=\"https://podevcast.netlify.app/\">Home page</a></li>\n<li><a href=\"https://podevcast.netlify.app/list\">Podcast
  page</a></li>\n<li><a href=\"https://podevcast.netlify.app/the_real_python_podcast/ep/1/\">Episode
  page</a></li>\n<li><a href=\"https://podevcast.netlify.app/command_line_heroes/\">Podcast
  List page</a></li>\n<li><a href=\"https://podevcast.netlify.app/category/\">Categories
  page</a></li>\n</ol>\n<p>The Home page has the latest episode of all the podcasts.
  It also has an audio player to play on the go.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1645113477/blog-media/iafi8nthhj0vvvrcbhka.png\"
  alt=\"Podevcast home page\" /></p>\n<p>The Podcast List page has the list of all
  the Podcasts available in the project. It has the name of the podcast with the link
  to the podcast page that has the list of all the episodes of that podcast.</p>\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1645113598/blog-media/cnprgufs3lrouvgdl8jn.png\"
  alt=\"Podevcast Podcast list\" /></p>\n<p>The categories page has a list of categories
  of the podcasts like Web-development, backend, frontend, data science, DevOps, and
  so on. More categories will be added soon.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1645113626/blog-media/uloq4xi1d4zfo8sfl7bm.png\"
  alt=\"Podevcast Categories\" /></p>\n<p>The Episode page has the audio player, the
  summary of the episode, canonical episode, and podcast page.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1645113654/blog-media/omqks44p8b3u7jclkhgz.png\"
  alt=\"Podevcast Episode page\" /></p>\n<h2>Why Podevcast?</h2>\n<p>Listening to
  music is one thing and listening to podcasts is different. I wanted a place from
  where developers can listen to developer-specific podcasts from a single source
  not just give out the article <strong>&quot;Top 10 podcast you should be listening
  to as a developer&quot;</strong>. Having played around with python and some libraries
  like feedparser and jinga previously I saw this Hackathon as an opportunity to convert
  the idea into a project. It fits the JAMStack area well from the Hackathon and project
  perspective.</p>\n<h2>Tech Stack</h2>\n<ul>\n<li>Python\n<ul>\n<li><a href=\"https://pypi.org/project/feedparser/\">feedparser</a></li>\n<li><a
  href=\"https://pypi.org/project/Jinja2/\">jinga2</a></li>\n</ul>\n</li>\n<li>GitHub
  Actions</li>\n<li>HTML / CSS</li>\n</ul>\n<p>The data is extracted from various
  RSS Feeds using the feedparser library in Python.</p>\n<p>Using GitHub Actions,
  the feed is refreshed every 24 hours to fetch the latest episodes from the respective
  podcast feeds. Basically, the GitHub action triggers a Netlify deployment that in
  turn generates the static site by running the script.</p>\n<p>The command for running
  the script on Netlify and generating the <code>Podevcast</code> webpage is :</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>pip install -r rquirements.txt
  &amp;&amp; python script.py\n</pre></div>\n\n</pre>\n\n<p>And the directory for
  deployed web pages (published directory) is <code>site</code> which contains all
  the <code>HTML</code> files that can be rendered as the website itself.</p>\n<h3>Source
  Code</h3>\n<p>The project is available on <a href=\"https://github.com/Mr-Destructive/podevcast\">GitHub</a>.
  Feel free to open a PR to add a Podcast or a Category. The project only has a few
  python files, the main script is <code>script.py</code> which actually creates the
  home and the podcast list pages along with the episode pages. The <code>src</code>
  folder contains some extra bits of scripts like creating the categories and category
  podcast list pages. Also, it has certain config files like <code>runtime.txt</code>,
  <code>requirements.txt</code>, and so on. Finally, there is the <code>podlist.json</code>
  for the list of podcasts and <code>categorylist.json</code> for the categories of
  podcasts.</p>\n<h3>Core Script Snippet</h3>\n<p>The python script looks a lot bigger
  than the below snippet but it is doing the same process multiple times for different
  pages. There is also some type checking and tiny details that are added as per the
  requirement of the templates.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"kn\">import</span> <span class=\"nn\">feedparser</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">jinja2</span> <span class=\"kn\">import</span> <span class=\"n\">Environment</span><span
  class=\"p\">,</span> <span class=\"n\">FileSystemLoader</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">pathlib</span> <span class=\"kn\">import</span> <span class=\"n\">Path</span>\n\n<span
  class=\"n\">template_env</span> <span class=\"o\">=</span> <span class=\"n\">Environment</span><span
  class=\"p\">(</span><span class=\"n\">loader</span><span class=\"o\">=</span><span
  class=\"n\">FileSystemLoader</span><span class=\"p\">(</span><span class=\"n\">searchpath</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;./layouts/&#39;</span><span class=\"p\">))</span>\n<span
  class=\"n\">index_template</span> <span class=\"o\">=</span> <span class=\"n\">template_env</span><span
  class=\"o\">.</span><span class=\"n\">get_template</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;index.html&#39;</span><span class=\"p\">)</span>\n<span class=\"n\">episode_template</span>
  <span class=\"o\">=</span> <span class=\"n\">template_env</span><span class=\"o\">.</span><span
  class=\"n\">get_template</span><span class=\"p\">(</span><span class=\"s1\">&#39;episode.html&#39;</span><span
  class=\"p\">)</span>\n\n<span class=\"n\">feed</span> <span class=\"o\">=</span>
  <span class=\"n\">feedparser</span><span class=\"o\">.</span><span class=\"n\">parse</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;https://freecodecamp.libsyn.com/rss&quot;</span><span
  class=\"p\">)</span>\n\n<span class=\"n\">pod_name</span> <span class=\"o\">=</span>
  <span class=\"n\">feed</span><span class=\"p\">[</span><span class=\"s1\">&#39;feed&#39;</span><span
  class=\"p\">][</span><span class=\"s1\">&#39;title&#39;</span><span class=\"p\">]</span>\n\n<span
  class=\"k\">for</span> <span class=\"n\">i</span> <span class=\"ow\">in</span> <span
  class=\"nb\">range</span><span class=\"p\">(</span><span class=\"mi\">0</span><span
  class=\"p\">,</span> <span class=\"nb\">len</span><span class=\"p\">(</span><span
  class=\"n\">feed</span><span class=\"p\">[</span><span class=\"s1\">&#39;entries&#39;</span><span
  class=\"p\">]):</span>\n    \n    <span class=\"n\">ep_title</span> <span class=\"o\">=</span>
  <span class=\"n\">feed</span><span class=\"p\">[</span><span class=\"s1\">&#39;entries&#39;</span><span
  class=\"p\">][</span><span class=\"n\">i</span><span class=\"p\">][</span><span
  class=\"s1\">&#39;title&#39;</span><span class=\"p\">]</span>\n    <span class=\"n\">audio</span>
  <span class=\"o\">=</span> <span class=\"n\">feed</span><span class=\"p\">[</span><span
  class=\"s1\">&#39;entries&#39;</span><span class=\"p\">][</span><span class=\"n\">i</span><span
  class=\"p\">][</span><span class=\"s1\">&#39;links&#39;</span><span class=\"p\">][</span><span
  class=\"mi\">1</span><span class=\"p\">][</span><span class=\"s1\">&#39;href&#39;</span><span
  class=\"p\">]</span>\n    <span class=\"n\">cover_image</span> <span class=\"o\">=</span>
  <span class=\"n\">feed</span><span class=\"p\">[</span><span class=\"s1\">&#39;entries&#39;</span><span
  class=\"p\">][</span><span class=\"n\">i</span><span class=\"p\">][</span><span
  class=\"s1\">&#39;image&#39;</span><span class=\"p\">][</span><span class=\"s1\">&#39;href&#39;</span><span
  class=\"p\">]</span>\n    <span class=\"n\">og_link</span> <span class=\"o\">=</span>
  <span class=\"n\">feed</span><span class=\"p\">[</span><span class=\"s1\">&#39;entries&#39;</span><span
  class=\"p\">][</span><span class=\"n\">i</span><span class=\"p\">][</span><span
  class=\"s1\">&#39;links&#39;</span><span class=\"p\">][</span><span class=\"mi\">0</span><span
  class=\"p\">][</span><span class=\"s1\">&#39;href&#39;</span><span class=\"p\">]</span>\n\n
  \   <span class=\"n\">episode_obj</span> <span class=\"o\">=</span> <span class=\"p\">{}</span>\n
  \   <span class=\"n\">episode_obj</span><span class=\"p\">[</span><span class=\"s1\">&#39;title&#39;</span><span
  class=\"p\">]</span> <span class=\"o\">=</span> <span class=\"n\">ep_title</span>\n
  \   <span class=\"n\">episode_obj</span><span class=\"p\">[</span><span class=\"s1\">&#39;audiolink&#39;</span><span
  class=\"p\">]</span> <span class=\"o\">=</span> <span class=\"n\">audio</span>\n
  \   <span class=\"n\">episode_obj</span><span class=\"p\">[</span><span class=\"s1\">&#39;cover&#39;</span><span
  class=\"p\">]</span> <span class=\"o\">=</span> <span class=\"n\">cover_image</span>\n
  \   <span class=\"n\">episode_obj</span><span class=\"p\">[</span><span class=\"s1\">&#39;link&#39;</span><span
  class=\"p\">]</span> <span class=\"o\">=</span> <span class=\"n\">og_link</span>\n\n
  \   <span class=\"k\">with</span> <span class=\"nb\">open</span><span class=\"p\">(</span><span
  class=\"n\">os</span><span class=\"o\">.</span><span class=\"n\">path</span><span
  class=\"o\">.</span><span class=\"n\">join</span><span class=\"p\">(</span><span
  class=\"sa\">f</span><span class=\"s2\">&quot;site/</span><span class=\"si\">{</span><span
  class=\"n\">pod_name</span><span class=\"si\">}</span><span class=\"s2\">/ep/</span><span
  class=\"si\">{</span><span class=\"n\">i</span><span class=\"si\">}</span><span
  class=\"s2\">/index.html&quot;</span><span class=\"p\">),</span> <span class=\"s1\">&#39;w&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">encoding</span><span class=\"o\">=</span><span
  class=\"s1\">&#39;utf-8&#39;</span><span class=\"p\">)</span> <span class=\"k\">as</span>
  <span class=\"n\">ep_file</span><span class=\"p\">:</span>\n         <span class=\"n\">ep_file</span><span
  class=\"o\">.</span><span class=\"n\">write</span><span class=\"p\">(</span>\n            <span
  class=\"n\">episode_template</span><span class=\"o\">.</span><span class=\"n\">render</span><span
  class=\"p\">(</span>\n            <span class=\"n\">episode</span> <span class=\"o\">=</span>
  <span class=\"n\">episode_obj</span>\n            <span class=\"p\">)</span>\n         <span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>Above is a simple snippet of
  the core functionality of the script. It basically takes the RSS Feed <code>URL</code>
  of the podcast and using <code>feedparser</code> the data is retrieved in the form
  of a dictionary in Python.</p>\n<ul>\n<li>Iterate over the <code>feed['entries']</code>
  which is a list of lengths same as the number of episodes of that podcast, we then
  assign a set of values like <code>episode title</code>, <code>audio link</code>,
  <code>cover image</code>, <code>canonical link for the episode</code>, <code>date</code>
  and so on.</li>\n<li>Create a dictionary and store the key value as the mentioned
  data to access from the template.</li>\n<li>Open a file in the structured file format
  and then parse the <code>episode_obj</code> which is a dictionary to the episode
  template.</li>\n<li>Access the dictionary using jinga2 templating tags.</li>\n</ul>\n<pre
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
  class=\"p\">&lt;</span><span class=\"nt\">html</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;</span><span class=\"nt\">head</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;</span><span class=\"nt\">title</span><span class=\"p\">&gt;</span>Podevcast<span
  class=\"p\">&lt;/</span><span class=\"nt\">title</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;/</span><span class=\"nt\">head</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;</span><span class=\"nt\">body</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;</span><span class=\"nt\">h3</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;ep-title&quot;</span><span class=\"p\">&gt;</span>{{
  episode.title }}<span class=\"p\">&lt;/</span><span class=\"nt\">h3</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">img</span>
  <span class=\"na\">src</span><span class=\"o\">=</span><span class=\"s\">&quot;{{
  episode.cover }}&quot;</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;ep-link&quot;</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{{ episode.link }}&quot;</span><span class=\"p\">&gt;</span>Episode
  <span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span class=\"p\">&gt;</span>
  \n        <span class=\"p\">&lt;</span><span class=\"nt\">audio</span> <span class=\"na\">controls</span><span
  class=\"o\">=</span><span class=\"s\">&quot;enabled&quot;</span> <span class=\"na\">preload</span><span
  class=\"o\">=</span><span class=\"s\">&quot;none&quot;</span><span class=\"p\">&gt;</span>\n
  \           <span class=\"p\">&lt;</span><span class=\"nt\">source</span> <span
  class=\"na\">src</span><span class=\"o\">=</span><span class=\"s\">&quot;{{ episode.audiolink
  }}&quot;</span> <span class=\"na\">type</span><span class=\"o\">=</span><span class=\"s\">&quot;audio/mpeg&quot;</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;/</span><span class=\"nt\">audio</span><span
  class=\"p\">&gt;</span>   \n    <span class=\"p\">&lt;/</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>We can use <code>{{  }}</code>
  to access any value parsed to the template via the script. Also, we can make use
  of <code>{% %}</code> to run loops, conditionals, blocks, and other tags in the
  template.</p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1645110268/blogmedia/uwdzcwn07oxhppiptem9.png\"
  alt=\"Feedparser Illustration\" /></p>\n<p>So, we can see the feed is basically
  a dictionary that has a key-value pair and further, it can be a nested dictionary
  or a list as a value of a key. As in the case of <code>feed['entries']</code> is
  a list with the length of the number of episodes of a podcast. And in the script,
  we use various keys to access various components, obviously, this requires a bit
  of exploration of the dictionary initially but it becomes easy thereafter to automate
  using Python.</p>\n<h3>Episode List</h3>\n<p>Currently, the episodes are added using
  the JSON file. It is not that user-friendly but still not a big task to simply add
  a link in a file. This is a #TODO that will require some external tooling to integrate
  into the webpage to ask for a form to submit a new Podcast.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;Command Line Heroes&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;https://feeds.pacific-content.com/commandlineheroes&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;Python Podcast__init__&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;https://www.pythonpodcast.com/feed/mp3/&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;Real Python Podcast&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;https://realpython.com/podcasts/rpp/feed&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;The freeCodeCamp Podcast&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;https://freecodecamp.libsyn.com/rss&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;CodeNewbie&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;http://feeds.codenewbie.org/cnpodcast.xml&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;Linux For Everyone&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;https://feeds.fireside.fm/linuxforeveryone/rss&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;JavaScript Jabber&quot;</span><span class=\"w\"> </span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;https://feeds.fireside.fm/javascriptjabber/rss&quot;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  process requires a manual test to validate a given RSS Feed as not all feeds are
  generated the same way and thus there are a few exceptions that need to be sorted
  out manually. For example, the Python Podcast doesn't have a cover image parsed
  into the RSS Feed, so there needs to be a check for it in the script and also in
  the template to restrict parsing and displaying the cover image link.</p>\n<h3>Episode
  Categories</h3>\n<p>This is also a JSON file that holds the keys as the category
  and the value as a list of episode names (strictly the name from <code>feed['feed']['title']</code>).
  There needs to be a human decision to be taken to add the podcast into a specific
  category.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">   </span><span
  class=\"nt\">&quot;Python&quot;</span><span class=\"p\">:[</span><span class=\"w\"></span>\n<span
  class=\"w\">      </span><span class=\"s2\">&quot;Talk Python To Me&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">      </span><span
  class=\"s2\">&quot;The Python Podcast.__init__&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">      </span><span class=\"s2\">&quot;The
  Real Python Podcast&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">      </span><span class=\"s2\">&quot;Python Bytes&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">   </span><span class=\"p\">],</span><span
  class=\"w\"></span>\n<span class=\"w\">   </span><span class=\"nt\">&quot;Javascript&quot;</span><span
  class=\"p\">:[</span><span class=\"w\"></span>\n<span class=\"w\">      </span><span
  class=\"s2\">&quot;Full Stack Radio&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">      </span><span class=\"s2\">&quot;JavaScript
  Jabber&quot;</span><span class=\"w\"></span>\n<span class=\"w\">   </span><span
  class=\"p\">],</span><span class=\"w\"></span>\n<span class=\"w\">   </span><span
  class=\"nt\">&quot;Linux&quot;</span><span class=\"p\">:[</span><span class=\"w\"></span>\n<span
  class=\"w\">      </span><span class=\"s2\">&quot;Command Line Heroes&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">      </span><span
  class=\"s2\">&quot;LINUX Unplugged&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">      </span><span class=\"s2\">&quot;The Linux Cast&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">      </span><span
  class=\"s2\">&quot;Linux For Everyone&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">   </span><span class=\"p\">],</span><span class=\"w\"></span>\n<span
  class=\"w\">   </span><span class=\"nt\">&quot;Data Science&quot;</span><span class=\"p\">:[</span><span
  class=\"w\"></span>\n<span class=\"w\">      </span><span class=\"s2\">&quot;DataFramed&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">      </span><span
  class=\"s2\">&quot;Data Skeptic&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">      </span><span class=\"s2\">&quot;The Banana Data Podcast&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">   </span><span class=\"p\">],</span><span
  class=\"w\"></span>\n<span class=\"w\">   </span><span class=\"nt\">&quot;Dev Ops&quot;</span><span
  class=\"p\">:[</span><span class=\"w\"></span>\n<span class=\"w\">      </span><span
  class=\"s2\">&quot;DevOps Cafe Podcast&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">      </span><span class=\"s2\">&quot;Arrested
  DevOps&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">
  \     </span><span class=\"s2\">&quot;Pulling the Strings&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">      </span><span class=\"s2\">&quot;Azure
  DevOps Podcast&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">      </span><span class=\"s2\">&quot;DevOps and Docker Talk&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">   </span><span class=\"p\">]</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Though
  the JSON file is managed manually the generation of the categories is automated.
  Please feel to add other categories of your choice.</p>\n<h2>What's Coming?</h2>\n<p>Certain
  features like adding podcast using a form, adding more podcasts, and categories
  for sure. Though what looks a bit cloudy in my opinion is adding accessibility links
  to music players because the RSS feed doesn't contain direct links to them. Though
  I still to explore and find out if it can be obtained from the feed itself.</p>\n<ul>\n<li>Search
  box for searching podcasts</li>\n<li>Accessible Links to other platforms (Spotify,
  Itunes, etc)</li>\n<li>More depth in categories (Languages/Frameworks/Niche-specific
  podcasts)</li>\n</ul>\n<p>I'll add these features after checking the feasibility
  of the ideas and the response from the community after releasing them.</p>\n<h2>Final
  Words</h2>\n<p>This project wouldn't have existed without this Hackathon as it gives
  a deadline to finish and hope to win something. Specially thanks to Hashnode and
  Netlify for organizing such a great opportunity in the form of a hackathon. Also,
  the maintainers of Python libraries like feedparser and jinja. The project would
  have been impossible without them.</p>\n<p>If you like the project please give it
  a star on <a href=\"https://github.com/Mr-Destructive/podevcast\">GitHub</a>. Have
  any feedback? Please let me know in the comments or on <a href=\"https://twitter.com/MeetGor21\">Twitter</a>.
  \ Thank you for reading, Hope you have a good time using Podevcast. Happy Coding
  :)</p>\n"
cover: ''
date: 2022-02-18
datetime: 2022-02-18 00:00:00+00:00
description: Hello Developers I am Meet Gor and I present this project as a submission
  to the Netlify x Hashnode Hackathon. Podevcast is a webpage(static) for listening
  to p
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2022-02-18-Podevcast.md
html: "<h2>Introduction</h2>\n<p>Hello Developers! Want to listen to programming podcasts
  from a single place? Podevcast is the place you should be searching for.</p>\n<p>I
  am Meet Gor and I present this project as a submission to the Netlify x Hashnode
  Hackathon. Podevcast is a webpage(static) for listening to podcasts centered around
  developers and programming. Just pick your favorite one and start listening straight
  away. Let's dive into the making of Podevcast. Head on to <a href=\"https://podevcast.netlify.app/\">https://podevcast.netlify.app/</a>
  to check out the live app.</p>\n<h2>What is Podevcast</h2>\n<p>Podevcast is a web
  application or a static site that renders the top programming/development podcasts.
  You can listen to the top podcasts around the developer community from a single
  source.</p>\n<blockquote>\n<p>Listen to your favorite developer podcasts with Podevcast</p>\n</blockquote>\n<p>Podevcast
  is a static site generated using a script. There is a static site generator that
  is heavily done in Python and deployed to Netlify. You can simply listen to the
  podcasts on the web page or go to the canonical page of the podcast episode. From
  the canonical page, you can choose to hop to your chosen music player, but the default
  music player should be fine for casual listening. The core idea is to keep things
  in a single place for developer podcasts.</p>\n<p><a href=\"https://podevcast.netlify.app/\">Podevcast</a></p>\n<p><a
  href=\"https://github.com/Mr-Destructive/podevcast\">Source Code</a></p>\n<h2>Preview</h2>\n<p>Podevcast
  has popular developer podcasts like <code>Command Line Heroes</code>, <code>The
  Python Podcast</code>, <code>The freeCodeCamp Podcast</code>, and many others to
  choose from. You can go into categories for looking at a specific podcast.</p>\n<h3>Application
  Demonstration</h3>\n<p>Here's a small demonstration of the Podevcast application.</p>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1645200224921/GC8gmxUzX.gif\"
  alt=\"Podevcast Preview gif\" /></p>\n<p>Podevcast has multiple pages like:</p>\n<ol>\n<li><a
  href=\"https://podevcast.netlify.app/\">Home page</a></li>\n<li><a href=\"https://podevcast.netlify.app/list\">Podcast
  page</a></li>\n<li><a href=\"https://podevcast.netlify.app/the_real_python_podcast/ep/1/\">Episode
  page</a></li>\n<li><a href=\"https://podevcast.netlify.app/command_line_heroes/\">Podcast
  List page</a></li>\n<li><a href=\"https://podevcast.netlify.app/category/\">Categories
  page</a></li>\n</ol>\n<p>The Home page has the latest episode of all the podcasts.
  It also has an audio player to play on the go.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1645113477/blog-media/iafi8nthhj0vvvrcbhka.png\"
  alt=\"Podevcast home page\" /></p>\n<p>The Podcast List page has the list of all
  the Podcasts available in the project. It has the name of the podcast with the link
  to the podcast page that has the list of all the episodes of that podcast.</p>\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1645113598/blog-media/cnprgufs3lrouvgdl8jn.png\"
  alt=\"Podevcast Podcast list\" /></p>\n<p>The categories page has a list of categories
  of the podcasts like Web-development, backend, frontend, data science, DevOps, and
  so on. More categories will be added soon.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1645113626/blog-media/uloq4xi1d4zfo8sfl7bm.png\"
  alt=\"Podevcast Categories\" /></p>\n<p>The Episode page has the audio player, the
  summary of the episode, canonical episode, and podcast page.</p>\n<p><img src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1645113654/blog-media/omqks44p8b3u7jclkhgz.png\"
  alt=\"Podevcast Episode page\" /></p>\n<h2>Why Podevcast?</h2>\n<p>Listening to
  music is one thing and listening to podcasts is different. I wanted a place from
  where developers can listen to developer-specific podcasts from a single source
  not just give out the article <strong>&quot;Top 10 podcast you should be listening
  to as a developer&quot;</strong>. Having played around with python and some libraries
  like feedparser and jinga previously I saw this Hackathon as an opportunity to convert
  the idea into a project. It fits the JAMStack area well from the Hackathon and project
  perspective.</p>\n<h2>Tech Stack</h2>\n<ul>\n<li>Python\n<ul>\n<li><a href=\"https://pypi.org/project/feedparser/\">feedparser</a></li>\n<li><a
  href=\"https://pypi.org/project/Jinja2/\">jinga2</a></li>\n</ul>\n</li>\n<li>GitHub
  Actions</li>\n<li>HTML / CSS</li>\n</ul>\n<p>The data is extracted from various
  RSS Feeds using the feedparser library in Python.</p>\n<p>Using GitHub Actions,
  the feed is refreshed every 24 hours to fetch the latest episodes from the respective
  podcast feeds. Basically, the GitHub action triggers a Netlify deployment that in
  turn generates the static site by running the script.</p>\n<p>The command for running
  the script on Netlify and generating the <code>Podevcast</code> webpage is :</p>\n<pre
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
  \       \n\n<div class=\"highlight\"><pre><span></span>pip install -r rquirements.txt
  &amp;&amp; python script.py\n</pre></div>\n\n</pre>\n\n<p>And the directory for
  deployed web pages (published directory) is <code>site</code> which contains all
  the <code>HTML</code> files that can be rendered as the website itself.</p>\n<h3>Source
  Code</h3>\n<p>The project is available on <a href=\"https://github.com/Mr-Destructive/podevcast\">GitHub</a>.
  Feel free to open a PR to add a Podcast or a Category. The project only has a few
  python files, the main script is <code>script.py</code> which actually creates the
  home and the podcast list pages along with the episode pages. The <code>src</code>
  folder contains some extra bits of scripts like creating the categories and category
  podcast list pages. Also, it has certain config files like <code>runtime.txt</code>,
  <code>requirements.txt</code>, and so on. Finally, there is the <code>podlist.json</code>
  for the list of podcasts and <code>categorylist.json</code> for the categories of
  podcasts.</p>\n<h3>Core Script Snippet</h3>\n<p>The python script looks a lot bigger
  than the below snippet but it is doing the same process multiple times for different
  pages. There is also some type checking and tiny details that are added as per the
  requirement of the templates.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"kn\">import</span> <span class=\"nn\">feedparser</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">jinja2</span> <span class=\"kn\">import</span> <span class=\"n\">Environment</span><span
  class=\"p\">,</span> <span class=\"n\">FileSystemLoader</span>\n<span class=\"kn\">from</span>
  <span class=\"nn\">pathlib</span> <span class=\"kn\">import</span> <span class=\"n\">Path</span>\n\n<span
  class=\"n\">template_env</span> <span class=\"o\">=</span> <span class=\"n\">Environment</span><span
  class=\"p\">(</span><span class=\"n\">loader</span><span class=\"o\">=</span><span
  class=\"n\">FileSystemLoader</span><span class=\"p\">(</span><span class=\"n\">searchpath</span><span
  class=\"o\">=</span><span class=\"s1\">&#39;./layouts/&#39;</span><span class=\"p\">))</span>\n<span
  class=\"n\">index_template</span> <span class=\"o\">=</span> <span class=\"n\">template_env</span><span
  class=\"o\">.</span><span class=\"n\">get_template</span><span class=\"p\">(</span><span
  class=\"s1\">&#39;index.html&#39;</span><span class=\"p\">)</span>\n<span class=\"n\">episode_template</span>
  <span class=\"o\">=</span> <span class=\"n\">template_env</span><span class=\"o\">.</span><span
  class=\"n\">get_template</span><span class=\"p\">(</span><span class=\"s1\">&#39;episode.html&#39;</span><span
  class=\"p\">)</span>\n\n<span class=\"n\">feed</span> <span class=\"o\">=</span>
  <span class=\"n\">feedparser</span><span class=\"o\">.</span><span class=\"n\">parse</span><span
  class=\"p\">(</span><span class=\"s2\">&quot;https://freecodecamp.libsyn.com/rss&quot;</span><span
  class=\"p\">)</span>\n\n<span class=\"n\">pod_name</span> <span class=\"o\">=</span>
  <span class=\"n\">feed</span><span class=\"p\">[</span><span class=\"s1\">&#39;feed&#39;</span><span
  class=\"p\">][</span><span class=\"s1\">&#39;title&#39;</span><span class=\"p\">]</span>\n\n<span
  class=\"k\">for</span> <span class=\"n\">i</span> <span class=\"ow\">in</span> <span
  class=\"nb\">range</span><span class=\"p\">(</span><span class=\"mi\">0</span><span
  class=\"p\">,</span> <span class=\"nb\">len</span><span class=\"p\">(</span><span
  class=\"n\">feed</span><span class=\"p\">[</span><span class=\"s1\">&#39;entries&#39;</span><span
  class=\"p\">]):</span>\n    \n    <span class=\"n\">ep_title</span> <span class=\"o\">=</span>
  <span class=\"n\">feed</span><span class=\"p\">[</span><span class=\"s1\">&#39;entries&#39;</span><span
  class=\"p\">][</span><span class=\"n\">i</span><span class=\"p\">][</span><span
  class=\"s1\">&#39;title&#39;</span><span class=\"p\">]</span>\n    <span class=\"n\">audio</span>
  <span class=\"o\">=</span> <span class=\"n\">feed</span><span class=\"p\">[</span><span
  class=\"s1\">&#39;entries&#39;</span><span class=\"p\">][</span><span class=\"n\">i</span><span
  class=\"p\">][</span><span class=\"s1\">&#39;links&#39;</span><span class=\"p\">][</span><span
  class=\"mi\">1</span><span class=\"p\">][</span><span class=\"s1\">&#39;href&#39;</span><span
  class=\"p\">]</span>\n    <span class=\"n\">cover_image</span> <span class=\"o\">=</span>
  <span class=\"n\">feed</span><span class=\"p\">[</span><span class=\"s1\">&#39;entries&#39;</span><span
  class=\"p\">][</span><span class=\"n\">i</span><span class=\"p\">][</span><span
  class=\"s1\">&#39;image&#39;</span><span class=\"p\">][</span><span class=\"s1\">&#39;href&#39;</span><span
  class=\"p\">]</span>\n    <span class=\"n\">og_link</span> <span class=\"o\">=</span>
  <span class=\"n\">feed</span><span class=\"p\">[</span><span class=\"s1\">&#39;entries&#39;</span><span
  class=\"p\">][</span><span class=\"n\">i</span><span class=\"p\">][</span><span
  class=\"s1\">&#39;links&#39;</span><span class=\"p\">][</span><span class=\"mi\">0</span><span
  class=\"p\">][</span><span class=\"s1\">&#39;href&#39;</span><span class=\"p\">]</span>\n\n
  \   <span class=\"n\">episode_obj</span> <span class=\"o\">=</span> <span class=\"p\">{}</span>\n
  \   <span class=\"n\">episode_obj</span><span class=\"p\">[</span><span class=\"s1\">&#39;title&#39;</span><span
  class=\"p\">]</span> <span class=\"o\">=</span> <span class=\"n\">ep_title</span>\n
  \   <span class=\"n\">episode_obj</span><span class=\"p\">[</span><span class=\"s1\">&#39;audiolink&#39;</span><span
  class=\"p\">]</span> <span class=\"o\">=</span> <span class=\"n\">audio</span>\n
  \   <span class=\"n\">episode_obj</span><span class=\"p\">[</span><span class=\"s1\">&#39;cover&#39;</span><span
  class=\"p\">]</span> <span class=\"o\">=</span> <span class=\"n\">cover_image</span>\n
  \   <span class=\"n\">episode_obj</span><span class=\"p\">[</span><span class=\"s1\">&#39;link&#39;</span><span
  class=\"p\">]</span> <span class=\"o\">=</span> <span class=\"n\">og_link</span>\n\n
  \   <span class=\"k\">with</span> <span class=\"nb\">open</span><span class=\"p\">(</span><span
  class=\"n\">os</span><span class=\"o\">.</span><span class=\"n\">path</span><span
  class=\"o\">.</span><span class=\"n\">join</span><span class=\"p\">(</span><span
  class=\"sa\">f</span><span class=\"s2\">&quot;site/</span><span class=\"si\">{</span><span
  class=\"n\">pod_name</span><span class=\"si\">}</span><span class=\"s2\">/ep/</span><span
  class=\"si\">{</span><span class=\"n\">i</span><span class=\"si\">}</span><span
  class=\"s2\">/index.html&quot;</span><span class=\"p\">),</span> <span class=\"s1\">&#39;w&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">encoding</span><span class=\"o\">=</span><span
  class=\"s1\">&#39;utf-8&#39;</span><span class=\"p\">)</span> <span class=\"k\">as</span>
  <span class=\"n\">ep_file</span><span class=\"p\">:</span>\n         <span class=\"n\">ep_file</span><span
  class=\"o\">.</span><span class=\"n\">write</span><span class=\"p\">(</span>\n            <span
  class=\"n\">episode_template</span><span class=\"o\">.</span><span class=\"n\">render</span><span
  class=\"p\">(</span>\n            <span class=\"n\">episode</span> <span class=\"o\">=</span>
  <span class=\"n\">episode_obj</span>\n            <span class=\"p\">)</span>\n         <span
  class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p>Above is a simple snippet of
  the core functionality of the script. It basically takes the RSS Feed <code>URL</code>
  of the podcast and using <code>feedparser</code> the data is retrieved in the form
  of a dictionary in Python.</p>\n<ul>\n<li>Iterate over the <code>feed['entries']</code>
  which is a list of lengths same as the number of episodes of that podcast, we then
  assign a set of values like <code>episode title</code>, <code>audio link</code>,
  <code>cover image</code>, <code>canonical link for the episode</code>, <code>date</code>
  and so on.</li>\n<li>Create a dictionary and store the key value as the mentioned
  data to access from the template.</li>\n<li>Open a file in the structured file format
  and then parse the <code>episode_obj</code> which is a dictionary to the episode
  template.</li>\n<li>Access the dictionary using jinga2 templating tags.</li>\n</ul>\n<pre
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
  class=\"p\">&lt;</span><span class=\"nt\">html</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;</span><span class=\"nt\">head</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;</span><span class=\"nt\">title</span><span class=\"p\">&gt;</span>Podevcast<span
  class=\"p\">&lt;/</span><span class=\"nt\">title</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;/</span><span class=\"nt\">head</span><span class=\"p\">&gt;</span>\n
  \   <span class=\"p\">&lt;</span><span class=\"nt\">body</span><span class=\"p\">&gt;</span>\n
  \       <span class=\"p\">&lt;</span><span class=\"nt\">h3</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;ep-title&quot;</span><span class=\"p\">&gt;</span>{{
  episode.title }}<span class=\"p\">&lt;/</span><span class=\"nt\">h3</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span class=\"nt\">img</span>
  <span class=\"na\">src</span><span class=\"o\">=</span><span class=\"s\">&quot;{{
  episode.cover }}&quot;</span><span class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;ep-link&quot;</span> <span class=\"na\">href</span><span class=\"o\">=</span><span
  class=\"s\">&quot;{{ episode.link }}&quot;</span><span class=\"p\">&gt;</span>Episode
  <span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span class=\"p\">&gt;</span>
  \n        <span class=\"p\">&lt;</span><span class=\"nt\">audio</span> <span class=\"na\">controls</span><span
  class=\"o\">=</span><span class=\"s\">&quot;enabled&quot;</span> <span class=\"na\">preload</span><span
  class=\"o\">=</span><span class=\"s\">&quot;none&quot;</span><span class=\"p\">&gt;</span>\n
  \           <span class=\"p\">&lt;</span><span class=\"nt\">source</span> <span
  class=\"na\">src</span><span class=\"o\">=</span><span class=\"s\">&quot;{{ episode.audiolink
  }}&quot;</span> <span class=\"na\">type</span><span class=\"o\">=</span><span class=\"s\">&quot;audio/mpeg&quot;</span><span
  class=\"p\">&gt;</span>\n        <span class=\"p\">&lt;/</span><span class=\"nt\">audio</span><span
  class=\"p\">&gt;</span>   \n    <span class=\"p\">&lt;/</span><span class=\"nt\">body</span><span
  class=\"p\">&gt;</span>\n<span class=\"p\">&lt;/</span><span class=\"nt\">html</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>We can use <code>{{  }}</code>
  to access any value parsed to the template via the script. Also, we can make use
  of <code>{% %}</code> to run loops, conditionals, blocks, and other tags in the
  template.</p>\n<p><img src=\"https://res.cloudinary.com/dgpxbrwoz/image/upload/v1645110268/blogmedia/uwdzcwn07oxhppiptem9.png\"
  alt=\"Feedparser Illustration\" /></p>\n<p>So, we can see the feed is basically
  a dictionary that has a key-value pair and further, it can be a nested dictionary
  or a list as a value of a key. As in the case of <code>feed['entries']</code> is
  a list with the length of the number of episodes of a podcast. And in the script,
  we use various keys to access various components, obviously, this requires a bit
  of exploration of the dictionary initially but it becomes easy thereafter to automate
  using Python.</p>\n<h3>Episode List</h3>\n<p>Currently, the episodes are added using
  the JSON file. It is not that user-friendly but still not a big task to simply add
  a link in a file. This is a #TODO that will require some external tooling to integrate
  into the webpage to ask for a form to submit a new Podcast.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;Command Line Heroes&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;https://feeds.pacific-content.com/commandlineheroes&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;Python Podcast__init__&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;https://www.pythonpodcast.com/feed/mp3/&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;Real Python Podcast&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;https://realpython.com/podcasts/rpp/feed&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;The freeCodeCamp Podcast&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;https://freecodecamp.libsyn.com/rss&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;CodeNewbie&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;http://feeds.codenewbie.org/cnpodcast.xml&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;Linux For Everyone&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;https://feeds.fireside.fm/linuxforeveryone/rss&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;JavaScript Jabber&quot;</span><span class=\"w\"> </span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;https://feeds.fireside.fm/javascriptjabber/rss&quot;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  process requires a manual test to validate a given RSS Feed as not all feeds are
  generated the same way and thus there are a few exceptions that need to be sorted
  out manually. For example, the Python Podcast doesn't have a cover image parsed
  into the RSS Feed, so there needs to be a check for it in the script and also in
  the template to restrict parsing and displaying the cover image link.</p>\n<h3>Episode
  Categories</h3>\n<p>This is also a JSON file that holds the keys as the category
  and the value as a list of episode names (strictly the name from <code>feed['feed']['title']</code>).
  There needs to be a human decision to be taken to add the podcast into a specific
  category.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">   </span><span
  class=\"nt\">&quot;Python&quot;</span><span class=\"p\">:[</span><span class=\"w\"></span>\n<span
  class=\"w\">      </span><span class=\"s2\">&quot;Talk Python To Me&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">      </span><span
  class=\"s2\">&quot;The Python Podcast.__init__&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">      </span><span class=\"s2\">&quot;The
  Real Python Podcast&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">      </span><span class=\"s2\">&quot;Python Bytes&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">   </span><span class=\"p\">],</span><span
  class=\"w\"></span>\n<span class=\"w\">   </span><span class=\"nt\">&quot;Javascript&quot;</span><span
  class=\"p\">:[</span><span class=\"w\"></span>\n<span class=\"w\">      </span><span
  class=\"s2\">&quot;Full Stack Radio&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">      </span><span class=\"s2\">&quot;JavaScript
  Jabber&quot;</span><span class=\"w\"></span>\n<span class=\"w\">   </span><span
  class=\"p\">],</span><span class=\"w\"></span>\n<span class=\"w\">   </span><span
  class=\"nt\">&quot;Linux&quot;</span><span class=\"p\">:[</span><span class=\"w\"></span>\n<span
  class=\"w\">      </span><span class=\"s2\">&quot;Command Line Heroes&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">      </span><span
  class=\"s2\">&quot;LINUX Unplugged&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">      </span><span class=\"s2\">&quot;The Linux Cast&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">      </span><span
  class=\"s2\">&quot;Linux For Everyone&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">   </span><span class=\"p\">],</span><span class=\"w\"></span>\n<span
  class=\"w\">   </span><span class=\"nt\">&quot;Data Science&quot;</span><span class=\"p\">:[</span><span
  class=\"w\"></span>\n<span class=\"w\">      </span><span class=\"s2\">&quot;DataFramed&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">      </span><span
  class=\"s2\">&quot;Data Skeptic&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">      </span><span class=\"s2\">&quot;The Banana Data Podcast&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">   </span><span class=\"p\">],</span><span
  class=\"w\"></span>\n<span class=\"w\">   </span><span class=\"nt\">&quot;Dev Ops&quot;</span><span
  class=\"p\">:[</span><span class=\"w\"></span>\n<span class=\"w\">      </span><span
  class=\"s2\">&quot;DevOps Cafe Podcast&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">      </span><span class=\"s2\">&quot;Arrested
  DevOps&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">
  \     </span><span class=\"s2\">&quot;Pulling the Strings&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">      </span><span class=\"s2\">&quot;Azure
  DevOps Podcast&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">      </span><span class=\"s2\">&quot;DevOps and Docker Talk&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">   </span><span class=\"p\">]</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Though
  the JSON file is managed manually the generation of the categories is automated.
  Please feel to add other categories of your choice.</p>\n<h2>What's Coming?</h2>\n<p>Certain
  features like adding podcast using a form, adding more podcasts, and categories
  for sure. Though what looks a bit cloudy in my opinion is adding accessibility links
  to music players because the RSS feed doesn't contain direct links to them. Though
  I still to explore and find out if it can be obtained from the feed itself.</p>\n<ul>\n<li>Search
  box for searching podcasts</li>\n<li>Accessible Links to other platforms (Spotify,
  Itunes, etc)</li>\n<li>More depth in categories (Languages/Frameworks/Niche-specific
  podcasts)</li>\n</ul>\n<p>I'll add these features after checking the feasibility
  of the ideas and the response from the community after releasing them.</p>\n<h2>Final
  Words</h2>\n<p>This project wouldn't have existed without this Hackathon as it gives
  a deadline to finish and hope to win something. Specially thanks to Hashnode and
  Netlify for organizing such a great opportunity in the form of a hackathon. Also,
  the maintainers of Python libraries like feedparser and jinja. The project would
  have been impossible without them.</p>\n<p>If you like the project please give it
  a star on <a href=\"https://github.com/Mr-Destructive/podevcast\">GitHub</a>. Have
  any feedback? Please let me know in the comments or on <a href=\"https://twitter.com/MeetGor21\">Twitter</a>.
  \ Thank you for reading, Hope you have a good time using Podevcast. Happy Coding
  :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/podevcast-python.png
long_description: 'Hello Developers I am Meet Gor and I present this project as a
  submission to the Netlify x Hashnode Hackathon. Podevcast is a webpage(static) for
  listening to podcasts centered around developers and programming. Just pick your
  favorite one and start '
now: 2025-05-01 18:11:33.315116
path: blog/posts/2022-02-18-Podevcast.md
prevnext: null
slug: podevcast-project
status: published
subtitle: 'Developer podcast from a single place, a podcast player static site generated
  using Python : Netlify x Hashnode Hackathon'
tags:
- python
- hashnode
templateKey: blog-post
title: 'Podevcast: A single source for developer podcasts'
today: 2025-05-01
---

## Introduction

Hello Developers! Want to listen to programming podcasts from a single place? Podevcast is the place you should be searching for. 

I am Meet Gor and I present this project as a submission to the Netlify x Hashnode Hackathon. Podevcast is a webpage(static) for listening to podcasts centered around developers and programming. Just pick your favorite one and start listening straight away. Let's dive into the making of Podevcast. Head on to https://podevcast.netlify.app/ to check out the live app.

## What is Podevcast

Podevcast is a web application or a static site that renders the top programming/development podcasts. You can listen to the top podcasts around the developer community from a single source. 

> Listen to your favorite developer podcasts with Podevcast

Podevcast is a static site generated using a script. There is a static site generator that is heavily done in Python and deployed to Netlify. You can simply listen to the podcasts on the web page or go to the canonical page of the podcast episode. From the canonical page, you can choose to hop to your chosen music player, but the default music player should be fine for casual listening. The core idea is to keep things in a single place for developer podcasts.

[Podevcast](https://podevcast.netlify.app/)

[Source Code](https://github.com/Mr-Destructive/podevcast)

## Preview

Podevcast has popular developer podcasts like `Command Line Heroes`, `The Python Podcast`, `The freeCodeCamp Podcast`, and many others to choose from. You can go into categories for looking at a specific podcast. 

### Application Demonstration

Here's a small demonstration of the Podevcast application.

![Podevcast Preview gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1645200224921/GC8gmxUzX.gif)

Podevcast has multiple pages like:

1. [Home page](https://podevcast.netlify.app/)
2. [Podcast page](https://podevcast.netlify.app/list)
3. [Episode page](https://podevcast.netlify.app/the_real_python_podcast/ep/1/)
4. [Podcast List page](https://podevcast.netlify.app/command_line_heroes/)
5. [Categories page](https://podevcast.netlify.app/category/)

The Home page has the latest episode of all the podcasts. It also has an audio player to play on the go.

![Podevcast home page](https://res.cloudinary.com/techstructive-blog/image/upload/v1645113477/blog-media/iafi8nthhj0vvvrcbhka.png)

The Podcast List page has the list of all the Podcasts available in the project. It has the name of the podcast with the link to the podcast page that has the list of all the episodes of that podcast.

![Podevcast Podcast list](https://res.cloudinary.com/techstructive-blog/image/upload/v1645113598/blog-media/cnprgufs3lrouvgdl8jn.png)

The categories page has a list of categories of the podcasts like Web-development, backend, frontend, data science, DevOps, and so on. More categories will be added soon.

![Podevcast Categories](https://res.cloudinary.com/techstructive-blog/image/upload/v1645113626/blog-media/uloq4xi1d4zfo8sfl7bm.png)

The Episode page has the audio player, the summary of the episode, canonical episode, and podcast page. 

![Podevcast Episode page](https://res.cloudinary.com/techstructive-blog/image/upload/v1645113654/blog-media/omqks44p8b3u7jclkhgz.png)

## Why Podevcast?

Listening to music is one thing and listening to podcasts is different. I wanted a place from where developers can listen to developer-specific podcasts from a single source not just give out the article **"Top 10 podcast you should be listening to as a developer"**. Having played around with python and some libraries like feedparser and jinga previously I saw this Hackathon as an opportunity to convert the idea into a project. It fits the JAMStack area well from the Hackathon and project perspective.  

## Tech Stack

- Python
  - [feedparser](https://pypi.org/project/feedparser/)
  - [jinga2](https://pypi.org/project/Jinja2/)
- GitHub Actions
- HTML / CSS

The data is extracted from various RSS Feeds using the feedparser library in Python. 

Using GitHub Actions, the feed is refreshed every 24 hours to fetch the latest episodes from the respective podcast feeds. Basically, the GitHub action triggers a Netlify deployment that in turn generates the static site by running the script.

The command for running the script on Netlify and generating the `Podevcast` webpage is :

```
pip install -r rquirements.txt && python script.py
```

And the directory for deployed web pages (published directory) is `site` which contains all the `HTML` files that can be rendered as the website itself. 

### Source Code

The project is available on [GitHub](https://github.com/Mr-Destructive/podevcast). Feel free to open a PR to add a Podcast or a Category. The project only has a few python files, the main script is `script.py` which actually creates the home and the podcast list pages along with the episode pages. The `src` folder contains some extra bits of scripts like creating the categories and category podcast list pages. Also, it has certain config files like `runtime.txt`, `requirements.txt`, and so on. Finally, there is the `podlist.json` for the list of podcasts and `categorylist.json` for the categories of podcasts. 

### Core Script Snippet 

The python script looks a lot bigger than the below snippet but it is doing the same process multiple times for different pages. There is also some type checking and tiny details that are added as per the requirement of the templates. 

```python
import feedparser
from jinja2 import Environment, FileSystemLoader
from pathlib import Path

template_env = Environment(loader=FileSystemLoader(searchpath='./layouts/'))
index_template = template_env.get_template('index.html')
episode_template = template_env.get_template('episode.html')

feed = feedparser.parse("https://freecodecamp.libsyn.com/rss")

pod_name = feed['feed']['title']

for i in range(0, len(feed['entries']):
    
    ep_title = feed['entries'][i]['title']
    audio = feed['entries'][i]['links'][1]['href']
    cover_image = feed['entries'][i]['image']['href']
    og_link = feed['entries'][i]['links'][0]['href']

    episode_obj = {}
    episode_obj['title'] = ep_title
    episode_obj['audiolink'] = audio
    episode_obj['cover'] = cover_image
    episode_obj['link'] = og_link

    with open(os.path.join(f"site/{pod_name}/ep/{i}/index.html"), 'w', encoding='utf-8') as ep_file:
         ep_file.write(
            episode_template.render(
            episode = episode_obj
            )
         )
```

   Above is a simple snippet of the core functionality of the script. It basically takes the RSS Feed `URL` of the podcast and using `feedparser` the data is retrieved in the form of a dictionary in Python. 

   - Iterate over the `feed['entries']` which is a list of lengths same as the number of episodes of that podcast, we then assign a set of values like `episode title`, `audio link`, `cover image`, `canonical link for the episode`, `date` and so on. 
   - Create a dictionary and store the key value as the mentioned data to access from the template. 
   - Open a file in the structured file format and then parse the `episode_obj` which is a dictionary to the episode template. 
   - Access the dictionary using jinga2 templating tags. 

```html
<html>
    <head>
        <title>Podevcast</title>
    </head>
    <body>
        <h3 class="ep-title">{{ episode.title }}</h3>
        <img src="{{ episode.cover }}">
        <a class="ep-link" href="{{ episode.link }}">Episode </a> 
        <audio controls="enabled" preload="none">
            <source src="{{ episode.audiolink }}" type="audio/mpeg">
        </audio>   
    </body>
</html>
```
We can use `{{  }}` to access any value parsed to the template via the script. Also, we can make use of `{% %}` to run loops, conditionals, blocks, and other tags in the template. 

![Feedparser Illustration](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1645110268/blogmedia/uwdzcwn07oxhppiptem9.png)

So, we can see the feed is basically a dictionary that has a key-value pair and further, it can be a nested dictionary or a list as a value of a key. As in the case of `feed['entries']` is a list with the length of the number of episodes of a podcast. And in the script, we use various keys to access various components, obviously, this requires a bit of exploration of the dictionary initially but it becomes easy thereafter to automate using Python. 

### Episode List

Currently, the episodes are added using the JSON file. It is not that user-friendly but still not a big task to simply add a link in a file. This is a #TODO that will require some external tooling to integrate into the webpage to ask for a form to submit a new Podcast. 

```json
{
    "Command Line Heroes": "https://feeds.pacific-content.com/commandlineheroes",
    "Python Podcast__init__": "https://www.pythonpodcast.com/feed/mp3/",
    "Real Python Podcast": "https://realpython.com/podcasts/rpp/feed",
    "The freeCodeCamp Podcast": "https://freecodecamp.libsyn.com/rss",
    "CodeNewbie": "http://feeds.codenewbie.org/cnpodcast.xml",
    "Linux For Everyone": "https://feeds.fireside.fm/linuxforeveryone/rss",
    "JavaScript Jabber" : "https://feeds.fireside.fm/javascriptjabber/rss"
}
```

The process requires a manual test to validate a given RSS Feed as not all feeds are generated the same way and thus there are a few exceptions that need to be sorted out manually. For example, the Python Podcast doesn't have a cover image parsed into the RSS Feed, so there needs to be a check for it in the script and also in the template to restrict parsing and displaying the cover image link. 

### Episode Categories

This is also a JSON file that holds the keys as the category and the value as a list of episode names (strictly the name from `feed['feed']['title']`). There needs to be a human decision to be taken to add the podcast into a specific category. 

```json
{
   "Python":[
      "Talk Python To Me",
      "The Python Podcast.__init__",
      "The Real Python Podcast",
      "Python Bytes"
   ],
   "Javascript":[
      "Full Stack Radio",
      "JavaScript Jabber"
   ],
   "Linux":[
      "Command Line Heroes",
      "LINUX Unplugged",
      "The Linux Cast",
      "Linux For Everyone"
   ],
   "Data Science":[
      "DataFramed",
      "Data Skeptic",
      "The Banana Data Podcast"
   ],
   "Dev Ops":[
      "DevOps Cafe Podcast",
      "Arrested DevOps",
      "Pulling the Strings",
      "Azure DevOps Podcast",
      "DevOps and Docker Talk"
   ]
}
```

   Though the JSON file is managed manually the generation of the categories is automated. Please feel to add other categories of your choice. 

## What's Coming?

Certain features like adding podcast using a form, adding more podcasts, and categories for sure. Though what looks a bit cloudy in my opinion is adding accessibility links to music players because the RSS feed doesn't contain direct links to them. Though I still to explore and find out if it can be obtained from the feed itself. 

- Search box for searching podcasts
- Accessible Links to other platforms (Spotify, Itunes, etc)
- More depth in categories (Languages/Frameworks/Niche-specific podcasts)

I'll add these features after checking the feasibility of the ideas and the response from the community after releasing them.

## Final Words

This project wouldn't have existed without this Hackathon as it gives a deadline to finish and hope to win something. Specially thanks to Hashnode and Netlify for organizing such a great opportunity in the form of a hackathon. Also, the maintainers of Python libraries like feedparser and jinja. The project would have been impossible without them. 

If you like the project please give it a star on [GitHub](https://github.com/Mr-Destructive/podevcast). Have any feedback? Please let me know in the comments or on [Twitter](https://twitter.com/MeetGor21).  Thank you for reading, Hope you have a good time using Podevcast. Happy Coding :)