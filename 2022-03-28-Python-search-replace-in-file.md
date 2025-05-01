---
article_html: "<h2>Searching and Replacing the text in a File</h2>\n<p>Using simple
  python semantics, we can perform search and replace in a file. Firstly, we will
  define the file name, along with the words to search and replace. After defining
  the sets of variables, we will open the file in <code>r+</code> mode i.e. we can
  perform read as well as write operations in the file.</p>\n<p>We will store the
  entire file contents using the <a href=\"https://docs.python.org/3/tutorial/inputoutput.html#reading-and-writing-files\">read</a>
  function, the contents of file are now stored in the form of a string. We further
  can set the position of the cursor or the current position in the file using the
  <a href=\"https://python-reference.readthedocs.io/en/latest/docs/file/seek.html\">seek</a>
  function. The seek function takes in a optional parameter as the position to set
  for reading/writing of file. Using the <a href=\"https://python-reference.readthedocs.io/en/latest/docs/file/truncate.html\">truncate</a>
  function, we can clear all the contents of the file.</p>\n<p>Finally, we generate
  the content by replacing the words i.e. the old word with the new word from the
  string which we store the contents. After replacing the content in the string, we
  write the string variable into the file and hence the substitution was performed
  in the text file.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"n\">file_name</span> <span class=\"o\">=</span> <span class=\"s1\">&#39;temp.txt&#39;</span>\n<span
  class=\"n\">old_text</span> <span class=\"o\">=</span> <span class=\"s1\">&#39;foo&#39;</span>\n<span
  class=\"n\">new_text</span> <span class=\"o\">=</span> <span class=\"s1\">&#39;python&#39;</span>\n\n<span
  class=\"k\">with</span> <span class=\"nb\">open</span><span class=\"p\">(</span><span
  class=\"n\">file_name</span><span class=\"p\">,</span> <span class=\"s2\">&quot;r+&quot;</span><span
  class=\"p\">)</span> <span class=\"k\">as</span> <span class=\"n\">fname</span><span
  class=\"p\">:</span>\n    <span class=\"n\">lines</span> <span class=\"o\">=</span>
  <span class=\"n\">fname</span><span class=\"o\">.</span><span class=\"n\">read</span><span
  class=\"p\">()</span>\n    <span class=\"n\">fname</span><span class=\"o\">.</span><span
  class=\"n\">seek</span><span class=\"p\">(</span><span class=\"mi\">0</span><span
  class=\"p\">)</span>\n    <span class=\"n\">fname</span><span class=\"o\">.</span><span
  class=\"n\">truncate</span><span class=\"p\">(</span><span class=\"mi\">0</span><span
  class=\"p\">)</span>\n    <span class=\"n\">subs</span> <span class=\"o\">=</span>
  <span class=\"n\">lines</span><span class=\"o\">.</span><span class=\"n\">replace</span><span
  class=\"p\">(</span><span class=\"n\">old_text</span><span class=\"p\">,</span>
  <span class=\"n\">new_text</span><span class=\"p\">)</span>\n    <span class=\"n\">fname</span><span
  class=\"o\">.</span><span class=\"n\">write</span><span class=\"p\">(</span><span
  class=\"n\">subs</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1648479344/blog-media/cstvfdlazyfriwvnilju.png\"
  alt=\"Python Search Replace in File\" /></p>\n<div class='prevnext'>\n    <style
  type='text/css'>\n    :root {\n      --prevnext-color-text: #eefbfe;\n      --prevnext-color-angle:
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
  \   </style>\n<pre><code>&lt;a class='prev' href='/python-dict-to-csv-table'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Turn
  Python dictionary into a neat CSV table&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a
  class='next' href='/dockerize-django-prj'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Dockerize
  a Django project&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot;
  height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
cover: ''
date: 2022-03-28
datetime: 2022-03-28 00:00:00+00:00
description: Perform search-replace operation in a file using python
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/tils/2022-03-28-Python-search-replace-in-file.md
html: "<h2>Searching and Replacing the text in a File</h2>\n<p>Using simple python
  semantics, we can perform search and replace in a file. Firstly, we will define
  the file name, along with the words to search and replace. After defining the sets
  of variables, we will open the file in <code>r+</code> mode i.e. we can perform
  read as well as write operations in the file.</p>\n<p>We will store the entire file
  contents using the <a href=\"https://docs.python.org/3/tutorial/inputoutput.html#reading-and-writing-files\">read</a>
  function, the contents of file are now stored in the form of a string. We further
  can set the position of the cursor or the current position in the file using the
  <a href=\"https://python-reference.readthedocs.io/en/latest/docs/file/seek.html\">seek</a>
  function. The seek function takes in a optional parameter as the position to set
  for reading/writing of file. Using the <a href=\"https://python-reference.readthedocs.io/en/latest/docs/file/truncate.html\">truncate</a>
  function, we can clear all the contents of the file.</p>\n<p>Finally, we generate
  the content by replacing the words i.e. the old word with the new word from the
  string which we store the contents. After replacing the content in the string, we
  write the string variable into the file and hence the substitution was performed
  in the text file.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"n\">file_name</span> <span class=\"o\">=</span> <span class=\"s1\">&#39;temp.txt&#39;</span>\n<span
  class=\"n\">old_text</span> <span class=\"o\">=</span> <span class=\"s1\">&#39;foo&#39;</span>\n<span
  class=\"n\">new_text</span> <span class=\"o\">=</span> <span class=\"s1\">&#39;python&#39;</span>\n\n<span
  class=\"k\">with</span> <span class=\"nb\">open</span><span class=\"p\">(</span><span
  class=\"n\">file_name</span><span class=\"p\">,</span> <span class=\"s2\">&quot;r+&quot;</span><span
  class=\"p\">)</span> <span class=\"k\">as</span> <span class=\"n\">fname</span><span
  class=\"p\">:</span>\n    <span class=\"n\">lines</span> <span class=\"o\">=</span>
  <span class=\"n\">fname</span><span class=\"o\">.</span><span class=\"n\">read</span><span
  class=\"p\">()</span>\n    <span class=\"n\">fname</span><span class=\"o\">.</span><span
  class=\"n\">seek</span><span class=\"p\">(</span><span class=\"mi\">0</span><span
  class=\"p\">)</span>\n    <span class=\"n\">fname</span><span class=\"o\">.</span><span
  class=\"n\">truncate</span><span class=\"p\">(</span><span class=\"mi\">0</span><span
  class=\"p\">)</span>\n    <span class=\"n\">subs</span> <span class=\"o\">=</span>
  <span class=\"n\">lines</span><span class=\"o\">.</span><span class=\"n\">replace</span><span
  class=\"p\">(</span><span class=\"n\">old_text</span><span class=\"p\">,</span>
  <span class=\"n\">new_text</span><span class=\"p\">)</span>\n    <span class=\"n\">fname</span><span
  class=\"o\">.</span><span class=\"n\">write</span><span class=\"p\">(</span><span
  class=\"n\">subs</span><span class=\"p\">)</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://res.cloudinary.com/techstructive-blog/image/upload/v1648479344/blog-media/cstvfdlazyfriwvnilju.png\"
  alt=\"Python Search Replace in File\" /></p>\n<div class='prevnext'>\n    <style
  type='text/css'>\n    :root {\n      --prevnext-color-text: #eefbfe;\n      --prevnext-color-angle:
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
  \   </style>\n<pre><code>&lt;a class='prev' href='/python-dict-to-csv-table'&gt;\n\n
  \   &lt;svg width=&quot;50px&quot; height=&quot;50px&quot; viewbox=&quot;0 0 24
  24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M13.5 8.25L9.75 12L13.5 15.75&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;
  &lt;/path&gt;\n    &lt;/svg&gt;\n    &lt;div class='prevnext-text'&gt;\n        &lt;p
  class='prevnext-subtitle'&gt;prev&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Turn
  Python dictionary into a neat CSV table&lt;/p&gt;\n    &lt;/div&gt;\n&lt;/a&gt;\n\n&lt;a
  class='next' href='/dockerize-django-prj'&gt;\n\n    &lt;div class='prevnext-text'&gt;\n
  \       &lt;p class='prevnext-subtitle'&gt;next&lt;/p&gt;\n        &lt;p class='prevnext-title'&gt;Dockerize
  a Django project&lt;/p&gt;\n    &lt;/div&gt;\n    &lt;svg width=&quot;50px&quot;
  height=&quot;50px&quot; viewbox=&quot;0 0 24 24&quot; fill=&quot;none&quot; xmlns=&quot;http://www.w3.org/2000/svg&quot;&gt;\n
  \       &lt;path d=&quot;M10.5 15.75L14.25 12L10.5 8.25&quot; stroke=&quot;var(--prevnext-color-angle)&quot;
  stroke-width=&quot;1.5&quot; stroke-linecap=&quot;round&quot; stroke-linejoin=&quot;round&quot;&gt;&lt;/path&gt;\n
  \   &lt;/svg&gt;\n&lt;/a&gt;\n</code></pre>\n  </div>"
long_description: Using simple python semantics, we can perform search and replace
  in a file. Firstly, we will define the file name, along with the words to search
  and replace. After defining the sets of variables, we will open the file in  We
  will store the entire fi
now: 2025-05-01 18:11:33.315169
path: blog/tils/2022-03-28-Python-search-replace-in-file.md
slug: python-search-replace-file
status: published-til
tags:
- python
templateKey: til
title: 'Python: Search and Replace in File'
today: 2025-05-01
---

## Searching and Replacing the text in a File

Using simple python semantics, we can perform search and replace in a file. Firstly, we will define the file name, along with the words to search and replace. After defining the sets of variables, we will open the file in `r+` mode i.e. we can perform read as well as write operations in the file.

We will store the entire file contents using the [read](https://docs.python.org/3/tutorial/inputoutput.html#reading-and-writing-files) function, the contents of file are now stored in the form of a string. We further can set the position of the cursor or the current position in the file using the [seek](https://python-reference.readthedocs.io/en/latest/docs/file/seek.html) function. The seek function takes in a optional parameter as the position to set for reading/writing of file. Using the [truncate](https://python-reference.readthedocs.io/en/latest/docs/file/truncate.html) function, we can clear all the contents of the file.

Finally, we generate the content by replacing the words i.e. the old word with the new word from the string which we store the contents. After replacing the content in the string, we write the string variable into the file and hence the substitution was performed in the text file.

```python
file_name = 'temp.txt'
old_text = 'foo'
new_text = 'python'

with open(file_name, "r+") as fname:
    lines = fname.read()
    fname.seek(0)
    fname.truncate(0)
    subs = lines.replace(old_text, new_text)
    fname.write(subs)
```

![Python Search Replace in File](https://res.cloudinary.com/techstructive-blog/image/upload/v1648479344/blog-media/cstvfdlazyfriwvnilju.png)
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
    
    <a class='prev' href='/python-dict-to-csv-table'>
    
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M13.5 8.25L9.75 12L13.5 15.75" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"> </path>
        </svg>
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>prev</p>
            <p class='prevnext-title'>Turn Python dictionary into a neat CSV table</p>
        </div>
    </a>
    
    <a class='next' href='/dockerize-django-prj'>
    
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>next</p>
            <p class='prevnext-title'>Dockerize a Django project</p>
        </div>
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.5 15.75L14.25 12L10.5 8.25" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path>
        </svg>
    </a>
  </div>