---
article_html: "<h1>Introduction</h1>\n<p>In this part, topics such as switch cases,
  mathematical expression, arrays, and strings. This won't be an in-depth guide to
  understand each and every concept, but to make users aware of the things and features
  in Bash. This also would not be an absolute basic guide, I expect to have some basic
  programming knowledge such as binary systems, logical and mathematical concepts.
  Don't worry, you won't be bombarded with commands, I'll just explain with easy examples
  to get started.</p>\n<p>Topics to be covered in this part are as follows:</p>\n<ul>\n<li>\n<p>User
  Input</p>\n<ul>\n<li>User Prompt</li>\n<li>Changing the Delimiter</li>\n<li>Password
  as Input</li>\n<li>Limiting the length of Input</li>\n</ul>\n</li>\n<li>\n<p>Cases</p>\n</li>\n<li>\n<p>Arrays</p>\n<ul>\n<li>Declaring
  and Printing Arrays</li>\n<li>Number of elements in an array</li>\n<li>Splicing
  the array</li>\n<li>Inserting and Deleting elements</li>\n</ul>\n</li>\n<li>\n<p>Strings</p>\n<ul>\n<li>Making
  Substrings</li>\n<li>String Concatenation</li>\n<li>String Comparison</li>\n</ul>\n</li>\n<li>\n<p>Arithmetic
  in Bash</p>\n<ul>\n<li>Integer Arithmetic</li>\n<li>Floating-Point Arithmetic</li>\n</ul>\n</li>\n</ul>\n<h1>User
  Input</h1>\n<p>Taking user input in Bash is quite straightforward and quite readable
  as well. We can make use of <code>read</code> command to take in input from the
  user. We just specify the variable in which we want to store the input.<code> read
  x</code> Here, the input will be stored in x. We can also pass in certain arguments
  to the read command such as -p (prompt with string), -r ( delimiter variation),
  -a(pass to the array), and others as well. Each of them will make the foundation
  of various complicated tasks to perform in logical operations.</p>\n<h3>User prompt
  argument</h3>\n<p>The -p argument will prompt the user with a string before they
  input anything. It makes quite informative and useful user input. This becomes quite
  a useful argument/parameter to make it quite readable and understand what to do
  directly with much hassle. The below is the general syntax of passing the argument
  to the read function.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nb\">read</span> -p <span class=\"s2\">&quot;Enter the number &quot;</span>
  n\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">read</span> -p <span class=\"s2\">&quot;Enter
  the number &quot;</span> n\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;The
  inputted number was </span><span class=\"nv\">$n</span><span class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625118915300/NRF7Ci2rK.png\"
  alt=\"bashs2.png\" /></p>\n<p>In this example, we can prompt the user with the string
  <strong>Enter the number</strong>, and it gives certain information to the user
  about what to input.</p>\n<h3>Changing the delimiter</h3>\n<p>Next, we can make
  use of -r which depending on the use case, we can change the delimiter while taking
  the input.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nv\">IFS</span><span class=\"o\">=</span><span
  class=\"s1\">&#39;/&#39;</span> <span class=\"nb\">read</span> -p <span class=\"s2\">&quot;Enter
  the file path : &quot;</span> user project app \n<span class=\"nb\">echo</span>
  <span class=\"nv\">$user</span> <span class=\"nv\">$project</span> <span class=\"nv\">$app</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625202319183/R9Eo3lg0oU.png\"
  alt=\"bashs2.png\" /></p>\n<p>In the above script, we separated the path of a directory(user-entered)
  into separate components such as the username, project name, and the app name, this
  can get pretty important and a great tool for automation of making project and application
  structures. At the beginning of the command, we use IFS which stands for Internal
  Field Separator, which does the separation of variables based on the field provided,
  in this case it was <code>//</code>, you can use any other field characters appropriate
  to your needs.</p>\n<p>This command will change the delimiter, by default it uses
  spaces or tabs etc to identify distinct input variables but we change it to other
  internal field separators such as / , \\ ,- , |, etc. This can make the user input
  more customizable and dynamic.</p>\n<h3>Password Typing</h3>\n<p>We can hide the
  input from the screen so as to keep it confidential and keep sensitive information
  such as passwords and keys private and protected.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">read</span> -sp <span class=\"s2\">&quot;Password:
  &quot;</span> pswd\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;the
  password was </span><span class=\"nv\">$pswd</span><span class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625161571993/MkCadbyKW.png\"
  alt=\"bashs2.png\" /></p>\n<p>We used the -s to keep the input hidden, the screen
  doesn't reflect what the user is typing, and -p for the prompt to offer the user
  some information on the text.</p>\n<h3>Limiting Length of Input</h3>\n<p>We can
  limit the user to only a certain number of characters as input. It becomes quite
  useful in constrained environments such as usernames and passwords to be restricted
  with a certain limit.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">read</span> -n <span class=\"m\">6</span>
  -p <span class=\"s2\">&quot;Enter the name: &quot;</span> n\n<span class=\"nb\">echo</span>
  <span class=\"nv\">$n</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625161752440/8xI5Lovxo.png\"
  alt=\"bashs2.png\" /></p>\n<p>In the above script, we demonstrate the limit of characters
  of 6 in the variable n. This restricts the user with only the first 6 characters,
  it just doesn't exceed ahead, directly to the next command.</p>\n<h3>Passing to
  the array</h3>\n<p>Another important argument to be passed after read command is
  -a which inserts the value to the array elements.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">read</span> -a nums -p <span
  class=\"s2\">&quot;Enter the elements : &quot;</span> \n<span class=\"k\">for</span>
  i <span class=\"k\">in</span> <span class=\"si\">${</span><span class=\"nv\">nums</span><span
  class=\"p\">[*]</span><span class=\"si\">}</span><span class=\"p\">;</span>\n<span
  class=\"k\">do</span> \n\t<span class=\"nb\">echo</span> -e <span class=\"s2\">&quot;</span><span
  class=\"nv\">$i</span><span class=\"s2\">\\n&quot;</span>\n<span class=\"k\">done</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above script, we have used array, don't worry, I'll explain it in the coming
  sections of this part. We have assigned an empty array and made the user enter those
  arrays, they are space-separated values. We have used the -a operator to insert
  the input to the elements of the array. The for loop is range-based which means
  it will do certain commands until there are no elements left in nums. The command
  <code>${nums[@]}</code> indicates every element in the array nums.</p>\n<h1>Cases</h1>\n<p>Cases
  are quite a good way of replacing nested if-else statements to make them nice and
  readable in the script.  Cases in Bash are quite powerful and easy to use compared
  with C/ C++ styled switch cases.</p>\n<p>The general structure of using a case in
  Bash is as follows:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"k\">case</span> variable <span class=\"k\">in</span>\n    pattern <span
  class=\"m\">1</span><span class=\"o\">)</span>\n        statements\n        <span
  class=\"p\">;;</span>\n    pattern <span class=\"m\">2</span><span class=\"o\">)</span>\n
  \       statements\n        <span class=\"p\">;;</span>\n    pattern <span class=\"m\">3</span><span
  class=\"o\">)</span>\n        statements\n        <span class=\"p\">;;</span>\n
  \   pattern <span class=\"m\">4</span><span class=\"o\">)</span>\n        statements\n
  \       <span class=\"p\">;;</span> \n    *<span class=\"o\">)</span>\n        statements\n
  \       <span class=\"p\">;;</span>\n<span class=\"k\">esac</span>\n</pre></div>\n\n</pre>\n\n<p>It
  follows a particular pattern if it matches it stops the search and executes the
  statements, finally comes out of the block. If it doesn't find any match it redirects
  to a default condition if any.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash </span>\n\n<span class=\"nb\">read</span> -p <span class=\"s2\">&quot;Enter
  a name : &quot;</span> n\n<span class=\"k\">case</span> <span class=\"nv\">$n</span>
  <span class=\"k\">in</span> \n\tadmin<span class=\"o\">)</span>\n\t\t<span class=\"nb\">echo</span>
  <span class=\"s2\">&quot;You are logged as root&quot;</span>\n\t\t<span class=\"p\">;;</span>\n\tunknown<span
  class=\"o\">)</span>\n\t\t<span class=\"nb\">echo</span> <span class=\"s2\">&quot;A
  hacker probably&quot;</span>\n\t\t<span class=\"p\">;;</span>\n\tmanager<span class=\"o\">)</span>\n\t\t<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;Weolcome Manager!&quot;</span>\n\t\t<span
  class=\"p\">;;</span>\n\t*<span class=\"o\">)</span>\n\t\t<span class=\"nb\">echo</span>
  <span class=\"s2\">&quot;A normal person&quot;</span>\n\t\t<span class=\"p\">;;</span>\n<span
  class=\"k\">esac</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625160454707/YDYGkU34d.png\"
  alt=\"bashs2.png\" /></p>\n<p>Seeing the above example, it is quite clear, that
  it looks quite structured and readable than the nested ladder of If-else statements.
  Cases are rendered based on the variable after the <code>case</code> keyword. We
  use the patterns before <code>)</code> as making the match in the variable provided
  before. Once the interpreter finds a match it returns to <code>esac</code> command
  which is <code>case</code> spelled in reverse just like <code>fi</code> for <code>if</code>
  and <code>done</code> for <code>do</code> in loops :) If it doesn't match any pattern,
  we have a default case represented by <code>*)</code> and it executes for any non-matching
  expression.</p>\n<h2>Arrays</h2>\n<p>Arrays or a way to store a list of numbers
  is implemented in Bash identical to most of the general programming languages.</p>\n<h3>Declaring
  and Printing Arrays</h3>\n<p>We declare an array similar to a variable but we mention
  the index of the element in the array(0 based index).  We can also simply declare
  an empty array using the declare command <code>declare -A nums</code></p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/bin/bash </span>\n\nnums<span class=\"o\">[</span><span class=\"m\">0</span><span
  class=\"o\">]=</span><span class=\"m\">7</span>\nnums<span class=\"o\">[</span><span
  class=\"m\">1</span><span class=\"o\">]=</span><span class=\"m\">5</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">2</span><span class=\"o\">]=</span><span class=\"m\">8</span>
  \n<span class=\"k\">for</span> i <span class=\"k\">in</span> <span class=\"si\">${</span><span
  class=\"nv\">nums</span><span class=\"p\">[@]</span><span class=\"si\">}</span>\n<span
  class=\"k\">do</span>\n<span class=\"nb\">echo</span> -e <span class=\"s2\">&quot;</span><span
  class=\"nv\">$i</span><span class=\"s2\"> \\n&quot;</span>\n<span class=\"k\">done</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625124595563/TzBEaH1E4.png\"
  alt=\"bashs2.png\" /></p>\n<p>The above script initializes an array with some hard-coded
  elements, surely you can input those from the user. For printing and accessing those
  elements in the array, We can use a loop, here we have used a range-based for loop.
  You can use any other loop you prefer. The iterator is &quot; i &quot; and we use
  $ to access the values from the array, we use <code>{}</code> as we have nested
  expression for indexing the element and <code>*</code> for every element in the
  array ( <code>@</code> will also work fine ), that's why range-based for loops make
  it quite simple to use. And we have simply printed &quot; i &quot; as it holds a
  particular element based on the iteration.</p>\n<p>OR</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">declare</span> -A <span class=\"nv\">nums</span><span
  class=\"o\">=(</span>\n<span class=\"o\">[</span><span class=\"m\">0</span><span
  class=\"o\">]=</span><span class=\"m\">44</span>\n<span class=\"o\">[</span><span
  class=\"m\">1</span><span class=\"o\">]=</span><span class=\"m\">45</span>\n<span
  class=\"o\">[</span><span class=\"m\">2</span><span class=\"o\">]=</span><span class=\"m\">46</span>\n<span
  class=\"o\">)</span>\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span
  class=\"si\">${</span><span class=\"nv\">nums</span><span class=\"p\">[@]</span><span
  class=\"si\">}</span><span class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625128343654/NCKUzurXe.png\"
  alt=\"bashs2.png\" /></p>\n<p>The above script uses declare an array, it can be
  empty as well after the name declaration. We used the <code>()</code> to include
  the values in the array, using indices in the array we assigned the values to the
  particular index.</p>\n<p>If you just want to print the elements, we can use <code>${nums[@]}</code>
  or <code>${nums[*]}</code>, this will print every element without using any iteration
  loops.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"ch\">#!/bin/bash</span>\n\nnums<span class=\"o\">[</span><span class=\"m\">0</span><span
  class=\"o\">]=</span><span class=\"m\">7</span>\nnums<span class=\"o\">[</span><span
  class=\"m\">1</span><span class=\"o\">]=</span><span class=\"m\">5</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">2</span><span class=\"o\">]=</span><span class=\"m\">8</span>
  \n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"si\">${</span><span
  class=\"nv\">nums</span><span class=\"p\">[@]</span><span class=\"si\">}</span><span
  class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625125456166/85zXjapQ_.png\"
  alt=\"bashs2.png\" /></p>\n<h3>Number of Elements in the array</h3>\n<p>To get the
  length of the array, we can use # in the expression <code>${nums[@]}</code>, like
  <code>${#nums[@]}</code> to get the number of elements from the array.</p>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625125770005/VzIr8CD7-.png\"
  alt=\"bashs2.png\" /></p>\n<p>Since we had 4 elements in the array, it accurately
  printed 4.</p>\n<h3>Inserting and Deleting elements from Array</h3>\n<p>We can push
  elements to the array using the assignment operator.</p>\n<p><code>nums=(${nums[@]}
  76) </code></p>\n<p>This will push 76 into the array, i.e. in the last index( length
  -1 index).</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"ch\">#!/bin/bash</span>\n\nnums<span class=\"o\">[</span><span class=\"m\">0</span><span
  class=\"o\">]=</span><span class=\"m\">7</span>\nnums<span class=\"o\">[</span><span
  class=\"m\">1</span><span class=\"o\">]=</span><span class=\"m\">5</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">2</span><span class=\"o\">]=</span><span class=\"m\">8</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">3</span><span class=\"o\">]=</span><span class=\"m\">19</span>\n<span
  class=\"nv\">nums</span><span class=\"o\">=(</span><span class=\"si\">${</span><span
  class=\"nv\">nums</span><span class=\"p\">[@]</span><span class=\"si\">}</span>
  <span class=\"m\">76</span><span class=\"o\">)</span>\n<span class=\"nb\">echo</span>
  <span class=\"s2\">&quot;</span><span class=\"si\">${</span><span class=\"nv\">nums</span><span
  class=\"p\">[@]</span><span class=\"si\">}</span><span class=\"s2\">&quot;</span>\n<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;Length of nums = </span><span
  class=\"si\">${#</span><span class=\"nv\">nums</span><span class=\"p\">[@]</span><span
  class=\"si\">}</span><span class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625126198675/A8IAE-2FR.png\"
  alt=\"bashs2.png\" /></p>\n<p>As you can see the element was added at the end of
  the array, you can also specify the index you want to insert. We can use <code>unset
  nums[3] </code> to delete the element at the particular location or we can pop back
  (delete from the end) an element from the array using the index <code>-1</code>
  from the array using the following command.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nb\">unset</span> nums<span class=\"o\">[</span>-1<span class=\"o\">]</span>\n</pre></div>\n\n</pre>\n\n<p>Any
  index provided will delete the element at that location by using unset. By using
  -1, we intend to refer to the last element. This can be quite handy and important
  as well in certain cases.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\nnums<span class=\"o\">[</span><span class=\"m\">0</span><span
  class=\"o\">]=</span><span class=\"m\">7</span>\nnums<span class=\"o\">[</span><span
  class=\"m\">1</span><span class=\"o\">]=</span><span class=\"m\">5</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">2</span><span class=\"o\">]=</span><span class=\"m\">8</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">3</span><span class=\"o\">]=</span><span class=\"m\">19</span>\n<span
  class=\"nb\">unset</span> nums<span class=\"o\">[</span>-1<span class=\"o\">]</span>\n<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"si\">${</span><span
  class=\"nv\">nums</span><span class=\"p\">[@]</span><span class=\"si\">}</span><span
  class=\"s2\">&quot;</span>\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;Length
  of nums = </span><span class=\"si\">${#</span><span class=\"nv\">nums</span><span
  class=\"p\">[@]</span><span class=\"si\">}</span><span class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625126770211/LYK2Q0Rp0.png\"
  alt=\"bashs2.png\" /></p>\n<p>There you can see we removed the element using unset.</p>\n<h3>Splicing
  an Array</h3>\n<p>We can splice the array to print/ copy a portion of the array
  to another one.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"si\">${</span><span
  class=\"nv\">nums</span><span class=\"p\">[@]:</span><span class=\"nv\">1</span><span
  class=\"p\">:</span><span class=\"nv\">3</span><span class=\"si\">}</span><span
  class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>Using two colons and numbers
  in between them, we can print in this case certain elements in the array from a
  particular range. Here the first number after the colon is the starting index to
  print from(inclusive) and the next number after the second colon is the length to
  which we would like to print the element, it is not the index but the number of
  elements after the start index to be spliced</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"ch\">#!/bin/bash</span>\n\nnums<span class=\"o\">[</span><span class=\"m\">0</span><span
  class=\"o\">]=</span><span class=\"m\">7</span>\nnums<span class=\"o\">[</span><span
  class=\"m\">1</span><span class=\"o\">]=</span><span class=\"m\">5</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">2</span><span class=\"o\">]=</span><span class=\"m\">8</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">3</span><span class=\"o\">]=</span><span class=\"m\">19</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">4</span><span class=\"o\">]=</span><span class=\"m\">76</span>\n<span
  class=\"nv\">newarr</span><span class=\"o\">=</span><span class=\"si\">${</span><span
  class=\"nv\">nums</span><span class=\"p\">[@]:</span><span class=\"nv\">1</span><span
  class=\"p\">:</span><span class=\"nv\">3</span><span class=\"si\">}</span>\n<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"si\">${</span><span
  class=\"nv\">newarr</span><span class=\"p\">[@]</span><span class=\"si\">}</span><span
  class=\"s2\">&quot;</span>\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span
  class=\"si\">${</span><span class=\"nv\">nums</span><span class=\"p\">[@]</span><span
  class=\"si\">}</span><span class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625127387739/IH2Fc6ghk.png\"
  alt=\"bashs2.png\" /></p>\n<p>In this case, we have copied the slice of an array
  to another new array using the double colon operator. We added the elements from
  the 1st index till <code>1+3</code> indices i.e till 4. 3 is not the index but the
  length till you would like to copy or print.</p>\n<p>This was a basic introduction
  to arrays, definitely, there will be much more stuff I didn't cover. Just to give
  an overview of how an array looks like in BASH scripting. Next, we move on to strings.</p>\n<h1>Strings</h1>\n<p>Strings
  are quite important as it forms the core of any script to deal with filenames, user
  information, etc all contain strings or array of characters. Let's take a closer
  look at how strings are declared, handled, and manipulated in Bash scripting.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nv\">s</span><span class=\"o\">=</span><span class=\"s2\">&quot;World&quot;</span>\n<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"nv\">$s</span><span
  class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625129318948/02V3bpP1I.png\"
  alt=\"bashs2.png\" /></p>\n<p>Strings are again declared as normal variables but
  are enclosed in double quotation marks.  And we access them in the exact same way
  as we do with variables. If you were to use single quotes instead of double quotes
  Bash would not interpret the variable name as a variable, it would print the name
  literally and not the value of the variable, So prefer using double quotes in echo
  and other commands to make variables accessible.</p>\n<h3>Making Substrings</h3>\n<p>We
  can even splice the string as we did with the arrays, in strings we can call it
  substrings. The syntax is almost identical as we just have to get the variable name.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nv\">s</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;Hello World&quot;</span>\n<span class=\"nv\">p</span><span class=\"o\">=</span><span
  class=\"si\">${</span><span class=\"nv\">s</span><span class=\"p\">:</span><span
  class=\"nv\">6</span><span class=\"si\">}</span>\n<span class=\"nb\">echo</span>
  <span class=\"nv\">$p</span>\n<span class=\"nv\">q</span><span class=\"o\">=</span><span
  class=\"si\">${</span><span class=\"nv\">s</span><span class=\"p\">::</span><span
  class=\"nv\">5</span><span class=\"si\">}</span>\n<span class=\"nb\">echo</span>
  <span class=\"nv\">$q</span>\n<span class=\"nv\">t</span><span class=\"o\">=</span><span
  class=\"si\">${</span><span class=\"nv\">s</span><span class=\"p\">:</span><span
  class=\"nv\">4</span><span class=\"p\">:</span><span class=\"nv\">3</span><span
  class=\"si\">}</span>\n<span class=\"nb\">echo</span> <span class=\"nv\">$t</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above script, we had a base string 's' which was then sliced from the 6th index
  to the end, If we do not pass the second number and colon, it interprets as the
  end of the string and if we do not pass the first number, it will interpret as the
  first character in the string. We sliced s from the 6th index till the end of the
  string and copied it in the string 'p''. In the 'q' string, we sliced the first
  5 characters from the string 's'. In the 't' string we sliced starting from the
  4th index and 3 characters in length i.e till  7th index, not the 7th index.</p>\n<p>We
  can use the # before the variable name to get the length of the string as we saw
  in the array section. So we can use the <code>echo ${#s}</code> command to print
  the length of the string where s is the string variable name.</p>\n<h3>String Concatenation</h3>\n<p>String
  concatenation on Bash is quite straightforward as it is just the matter of adding
  strings in a very simple way.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nv\">s</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;Hello&quot;</span>\n<span class=\"nv\">p</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;World&quot;</span>\n<span class=\"nv\">q</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;</span><span class=\"nv\">$s</span><span class=\"s2\"> </span><span
  class=\"nv\">$p</span><span class=\"s2\">&quot;</span>\n<span class=\"nb\">echo</span>
  <span class=\"nv\">$q</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625135997314/_n05RIoTTM.png\"
  alt=\"bashs2.png\" /></p>\n<p>The space between the two variables is quite literal,
  anything you place between this space or the double quotes will get stored in the
  variable or get printed.</p>\n<h3>String Comparison</h3>\n<p>Moving on to the string
  comparison in Bash. String comparison is quite complex in certain programming languages
  but it's quite straightforward in some languages such as Bash. We can compare strings
  quite easily in Bash, either they are equal or they are not, it's just comparison
  operators to perform the heavy-lifting for us.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nv\">s</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;hello&quot;</span>\n<span class=\"nv\">p</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;Hello&quot;</span>\n<span class=\"k\">if</span> <span class=\"o\">[</span>
  <span class=\"nv\">$s</span> <span class=\"o\">=</span> <span class=\"nv\">$p</span>
  <span class=\"o\">]</span><span class=\"p\">;</span>\n<span class=\"k\">then</span>\n
  \   <span class=\"nb\">echo</span> <span class=\"s2\">&quot;Equal&quot;</span>\n<span
  class=\"k\">else</span> \n    <span class=\"nb\">echo</span> <span class=\"s2\">&quot;Not
  equal&quot;</span>\n<span class=\"k\">fi</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625138020711/UWlRN8aPq.png\"
  alt=\"bashs2.png\" /></p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625137884154/7bNPkpGd3.png\"
  alt=\"bashs2.png\" /></p>\n<p>From the above code, it is quite clear that the strings
  as not equal and we compared them with the &quot;equality&quot; operator (=) and
  checked if that condition was true, and perform commands accordingly. We can also
  check if the strings are not equal using <code>!=</code> operator and we can perform
  commands based on the desired logic. We also have operators to compare the length
  of the strings. We can use <code>\\&lt;</code> operator to check if the first string
  is less than the second string(compared characters in ASCII).  And check if the
  first string is greater than the second string using <code>\\&gt;</code> operator.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nv\">s</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;hello&quot;</span>\n<span class=\"nv\">p</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;Hello&quot;</span>\n<span class=\"k\">if</span> <span class=\"o\">[</span>
  <span class=\"nv\">$s</span> <span class=\"se\">\\&lt;</span> <span class=\"nv\">$p</span>
  <span class=\"o\">]</span><span class=\"p\">;</span>\n<span class=\"k\">then</span>\n\n\t<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"nv\">$s</span><span
  class=\"s2\"> is Less than </span><span class=\"nv\">$p</span><span class=\"s2\">&quot;</span>\n<span
  class=\"k\">elif</span> <span class=\"o\">[</span> <span class=\"nv\">$s</span>
  <span class=\"se\">\\&gt;</span> <span class=\"nv\">$p</span> <span class=\"o\">]</span><span
  class=\"p\">;</span>\n<span class=\"k\">then</span>\n\t<span class=\"nb\">echo</span>
  <span class=\"s2\">&quot;</span><span class=\"nv\">$s</span><span class=\"s2\">
  is greater than </span><span class=\"nv\">$p</span><span class=\"s2\">&quot;</span>\n<span
  class=\"k\">else</span>\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;Equal&quot;</span>\n<span
  class=\"k\">fi</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625137393683/u3WbgDIrN.png\"
  alt=\"bashs2.png\" /></p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625137467448/NP1UXZAbv.png\"
  alt=\"bashs2.png\" /></p>\n<p>Here, we are using the ASCII equivalent of strings
  to compare them as it gives an idea in terms of the value of the strings. We see
  that 'h'( 104)has a greater ASCII value than 'H&quot; (72) which is why we see the
  shown outcome.</p>\n<p>We also have operators to check for the string being empty
  or not using the -z operator. Also, we have arguments to pass to the string comparison
  to check for non-empty strings as well, specifically for input validation and some
  error handling.</p>\n<p>We can quite easily use -n to check for non-empty string
  and -z for the length of the string being zero.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">read</span> -p <span class=\"s2\">&quot;Enter
  a string : &quot;</span> s\n<span class=\"k\">if</span> <span class=\"o\">[</span>
  -s <span class=\"nv\">$s</span> <span class=\"o\">]</span><span class=\"p\">;</span>\n<span
  class=\"k\">then</span> \n    <span class=\"nb\">echo</span> <span class=\"s2\">&quot;Empty
  Input&quot;</span>\n<span class=\"k\">else</span>\n   <span class=\"nb\">echo</span>
  <span class=\"s2\">&quot;Valid input&quot;</span>\n<span class=\"k\">fi</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625138907051/tbjRDda0U1.png\"
  alt=\"bashs2.png\" /></p>\n<p>So the string topic is quite straightforward and self-explanatory
  as it doesn't how that much complexity but is still powerful and convenient to use.</p>\n<h1>Arithmetic
  in Bash</h1>\n<p>Performing any Arithmetic operations is the core for scripting.
  Without arithmetic, it feels incomplete to programmatically create something, it
  would be quite menial to write commands by hand without having the ability to perform
  arithmetic operations.</p>\n<h3>Integer Arithmetic</h3>\n<p>Firstly we quite commonly
  use operations on variables, so let us see how to perform an arithmetic operation
  on variables in Bash. We use double curly braces to evaluate certain results of
  the operations performed on variables.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nv\">x</span><span class=\"o\">=</span><span
  class=\"m\">4</span>\n<span class=\"nv\">y</span><span class=\"o\">=</span><span
  class=\"m\">9</span>\n<span class=\"nv\">z</span><span class=\"o\">=</span><span
  class=\"k\">$((</span><span class=\"nv\">$x</span> <span class=\"o\">*</span> <span
  class=\"nv\">$y</span><span class=\"k\">))</span>\n<span class=\"nb\">echo</span>
  <span class=\"nv\">$z</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625139582135/Sb4hdE990.png\"
  alt=\"bashs2.png\" /></p>\n<p>We use double curly braces in order to evaluate the
  operations performed on the variables inside them. We definitely have to use the
  $ symbol to extract the value of the variable.</p>\n<p>We can surely use operators
  such as addition(<code>+</code>), subtraction(<code>-</code>), multiplication(<code>*</code>),
  division(<code>/</code>), and modulus(<code>%</code>, it stores the remainder of
  the division,17%3 gets you 2) in statements. We can also perform operations such
  as <code>&lt;&lt;</code> to do left bitwise shift and <code>&gt;&gt;</code> right
  bitwise shift to shift the binary digits in left tor right respectively in a variable.
  There are also logical operations such as Bitwise and logical AND(<code>&amp;</code>),
  OR(<code>|</code>), EX-OR(<code>^</code>), and ternary expressions.</p>\n<p>Alternative
  to double curly braces is <code>expr</code>, expr allows you to freely wherever
  you need to evaluate an arithmetic operation. Though this is not native in-built
  in shells, it uses a binary process to evaluate the arithmetic operations. It can
  also defer depending on the implementation of such commands in various environments.</p>\n<p>We
  can also use the <code>let</code> command to initialize a variable and perform expressions
  in the initialization itself.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">let</span> <span class=\"nv\">a</span><span
  class=\"o\">=</span><span class=\"m\">4</span>\n<span class=\"nb\">let</span> <span
  class=\"nv\">b</span><span class=\"o\">=</span>a*a\n<span class=\"nb\">let</span>
  <span class=\"nv\">c</span><span class=\"o\">=</span><span class=\"s2\">&quot;b/(a*2)&quot;</span>\n<span
  class=\"nb\">echo</span> <span class=\"nv\">$b</span>\n</pre></div>\n\n</pre>\n\n<p>We
  can perform quite complex operations using simple implementation using <code>let</code>,
  this allows much readable and bug-free scripts.  If you would like to perform operations
  with brackets and other operations you can enclose the expression in double quotation
  marks.</p>\n<h3>Floating-Point Arithmetic</h3>\n<p>Performing floating-point arithmetic
  in Bash is not quite well though. We won't get 100% accurate answers in the expressions
  this is because it is <strong>not designed</strong> for such things. Doing <strong>things
  related to floating-point is a bad idea</strong>, Still, you can improve the precision
  to a little extent to do some basic things. I <strong>don't recommend this</strong>
  only do this if there are no other options.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nb\">printf</span> %.9f <span class=\"s2\">&quot;</span><span class=\"k\">$((</span><span
  class=\"m\">10</span><span class=\"o\">/</span><span class=\"m\">3</span><span class=\"k\">))</span><span
  class=\"s2\"></span>\n</pre></div>\n\n</pre>\n\n<p>The result of this is 3.0000000..064
  roughly, which is pretty bad. Bash at its core doesn't support floating-point calculations.
  But there is good news, we have  <a href=\"https://en.wikipedia.org/wiki/AWK\">awk</a>
  \ and other tools such as  <a href=\"https://en.wikipedia.org/wiki/Bc_(programming_language)\">bc</a>
  \ and others which is planned for the next part in the series. I'll explain awk
  just for floating-point here, in the next part, I'll cover it in depth.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/bin/bash</span>\n<span class=\"nv\">a</span><span class=\"o\">=</span><span
  class=\"m\">10</span>\n<span class=\"nv\">b</span><span class=\"o\">=</span><span
  class=\"k\">$(</span><span class=\"nb\">echo</span> <span class=\"p\">|</span> awk
  -v <span class=\"nv\">a</span><span class=\"o\">=</span><span class=\"s2\">&quot;</span><span
  class=\"nv\">$a</span><span class=\"s2\">&quot;</span> <span class=\"s1\">&#39;{print
  a/3}&#39;</span><span class=\"k\">)</span>\n<span class=\"nb\">echo</span> <span
  class=\"nv\">$b</span> \n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625157391350/gHudsNntM4.png\"
  alt=\"bashs2.png\" /></p>\n<p>WOW! That is to the point, but that was a lot of hassle
  using echo but printing nothing! HH? OK, you see certain things can get really annoying
  when things aren't supported natively. Firstly, we use  | to pipe echo command with
  awk, the echo command doesn't do anything just a way to use awk command in assigning
  variables here. Then the general syntax of the awk command is <code> awk -options
  -commands</code>. In this case, we are using -v as an argument and passing in an
  as a variable which is equal to a, which is stupid and silly but that is what it
  is, you can name any variable name you want. Then we simply have to use the variable
  in the print function which generally evaluates the expressions or other operations
  and returns to the interpreter. And that is how we print the expression, Phew! That
  took a while to do some silly things, But hey! That's possible though.</p>\n<p>That
  is the basic overview of Arithmetic in Bash, you can also perform logical operations
  in it which is very easy and can be understood on a quick run-through in the  <a
  href=\"https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html#Arithmetic-Expansion\">documentation</a>.</p>\n<p>I
  hope you understood the mentioned topics and what are their use cases depending
  on the requirements. Some topics such as positional parameters, tools and utilities,
  dictionaries, and some other important aspects of Bash scripting will be covered
  in the next part. Happy Coding.</p>\n"
cover: ''
date: 2021-07-02
datetime: 2021-07-02 00:00:00+00:00
description: 'In this part, topics such as switch cases, mathematical expression,
  arrays, and strings. This won Topics to be covered in this part are as follows:
  User Input U'
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-07-02-BASH-scripting-P2.md
html: "<h1>Introduction</h1>\n<p>In this part, topics such as switch cases, mathematical
  expression, arrays, and strings. This won't be an in-depth guide to understand each
  and every concept, but to make users aware of the things and features in Bash. This
  also would not be an absolute basic guide, I expect to have some basic programming
  knowledge such as binary systems, logical and mathematical concepts. Don't worry,
  you won't be bombarded with commands, I'll just explain with easy examples to get
  started.</p>\n<p>Topics to be covered in this part are as follows:</p>\n<ul>\n<li>\n<p>User
  Input</p>\n<ul>\n<li>User Prompt</li>\n<li>Changing the Delimiter</li>\n<li>Password
  as Input</li>\n<li>Limiting the length of Input</li>\n</ul>\n</li>\n<li>\n<p>Cases</p>\n</li>\n<li>\n<p>Arrays</p>\n<ul>\n<li>Declaring
  and Printing Arrays</li>\n<li>Number of elements in an array</li>\n<li>Splicing
  the array</li>\n<li>Inserting and Deleting elements</li>\n</ul>\n</li>\n<li>\n<p>Strings</p>\n<ul>\n<li>Making
  Substrings</li>\n<li>String Concatenation</li>\n<li>String Comparison</li>\n</ul>\n</li>\n<li>\n<p>Arithmetic
  in Bash</p>\n<ul>\n<li>Integer Arithmetic</li>\n<li>Floating-Point Arithmetic</li>\n</ul>\n</li>\n</ul>\n<h1>User
  Input</h1>\n<p>Taking user input in Bash is quite straightforward and quite readable
  as well. We can make use of <code>read</code> command to take in input from the
  user. We just specify the variable in which we want to store the input.<code> read
  x</code> Here, the input will be stored in x. We can also pass in certain arguments
  to the read command such as -p (prompt with string), -r ( delimiter variation),
  -a(pass to the array), and others as well. Each of them will make the foundation
  of various complicated tasks to perform in logical operations.</p>\n<h3>User prompt
  argument</h3>\n<p>The -p argument will prompt the user with a string before they
  input anything. It makes quite informative and useful user input. This becomes quite
  a useful argument/parameter to make it quite readable and understand what to do
  directly with much hassle. The below is the general syntax of passing the argument
  to the read function.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nb\">read</span> -p <span class=\"s2\">&quot;Enter the number &quot;</span>
  n\n</pre></div>\n\n</pre>\n\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">read</span> -p <span class=\"s2\">&quot;Enter
  the number &quot;</span> n\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;The
  inputted number was </span><span class=\"nv\">$n</span><span class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625118915300/NRF7Ci2rK.png\"
  alt=\"bashs2.png\" /></p>\n<p>In this example, we can prompt the user with the string
  <strong>Enter the number</strong>, and it gives certain information to the user
  about what to input.</p>\n<h3>Changing the delimiter</h3>\n<p>Next, we can make
  use of -r which depending on the use case, we can change the delimiter while taking
  the input.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nv\">IFS</span><span class=\"o\">=</span><span
  class=\"s1\">&#39;/&#39;</span> <span class=\"nb\">read</span> -p <span class=\"s2\">&quot;Enter
  the file path : &quot;</span> user project app \n<span class=\"nb\">echo</span>
  <span class=\"nv\">$user</span> <span class=\"nv\">$project</span> <span class=\"nv\">$app</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625202319183/R9Eo3lg0oU.png\"
  alt=\"bashs2.png\" /></p>\n<p>In the above script, we separated the path of a directory(user-entered)
  into separate components such as the username, project name, and the app name, this
  can get pretty important and a great tool for automation of making project and application
  structures. At the beginning of the command, we use IFS which stands for Internal
  Field Separator, which does the separation of variables based on the field provided,
  in this case it was <code>//</code>, you can use any other field characters appropriate
  to your needs.</p>\n<p>This command will change the delimiter, by default it uses
  spaces or tabs etc to identify distinct input variables but we change it to other
  internal field separators such as / , \\ ,- , |, etc. This can make the user input
  more customizable and dynamic.</p>\n<h3>Password Typing</h3>\n<p>We can hide the
  input from the screen so as to keep it confidential and keep sensitive information
  such as passwords and keys private and protected.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">read</span> -sp <span class=\"s2\">&quot;Password:
  &quot;</span> pswd\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;the
  password was </span><span class=\"nv\">$pswd</span><span class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625161571993/MkCadbyKW.png\"
  alt=\"bashs2.png\" /></p>\n<p>We used the -s to keep the input hidden, the screen
  doesn't reflect what the user is typing, and -p for the prompt to offer the user
  some information on the text.</p>\n<h3>Limiting Length of Input</h3>\n<p>We can
  limit the user to only a certain number of characters as input. It becomes quite
  useful in constrained environments such as usernames and passwords to be restricted
  with a certain limit.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">read</span> -n <span class=\"m\">6</span>
  -p <span class=\"s2\">&quot;Enter the name: &quot;</span> n\n<span class=\"nb\">echo</span>
  <span class=\"nv\">$n</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625161752440/8xI5Lovxo.png\"
  alt=\"bashs2.png\" /></p>\n<p>In the above script, we demonstrate the limit of characters
  of 6 in the variable n. This restricts the user with only the first 6 characters,
  it just doesn't exceed ahead, directly to the next command.</p>\n<h3>Passing to
  the array</h3>\n<p>Another important argument to be passed after read command is
  -a which inserts the value to the array elements.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">read</span> -a nums -p <span
  class=\"s2\">&quot;Enter the elements : &quot;</span> \n<span class=\"k\">for</span>
  i <span class=\"k\">in</span> <span class=\"si\">${</span><span class=\"nv\">nums</span><span
  class=\"p\">[*]</span><span class=\"si\">}</span><span class=\"p\">;</span>\n<span
  class=\"k\">do</span> \n\t<span class=\"nb\">echo</span> -e <span class=\"s2\">&quot;</span><span
  class=\"nv\">$i</span><span class=\"s2\">\\n&quot;</span>\n<span class=\"k\">done</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above script, we have used array, don't worry, I'll explain it in the coming
  sections of this part. We have assigned an empty array and made the user enter those
  arrays, they are space-separated values. We have used the -a operator to insert
  the input to the elements of the array. The for loop is range-based which means
  it will do certain commands until there are no elements left in nums. The command
  <code>${nums[@]}</code> indicates every element in the array nums.</p>\n<h1>Cases</h1>\n<p>Cases
  are quite a good way of replacing nested if-else statements to make them nice and
  readable in the script.  Cases in Bash are quite powerful and easy to use compared
  with C/ C++ styled switch cases.</p>\n<p>The general structure of using a case in
  Bash is as follows:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"k\">case</span> variable <span class=\"k\">in</span>\n    pattern <span
  class=\"m\">1</span><span class=\"o\">)</span>\n        statements\n        <span
  class=\"p\">;;</span>\n    pattern <span class=\"m\">2</span><span class=\"o\">)</span>\n
  \       statements\n        <span class=\"p\">;;</span>\n    pattern <span class=\"m\">3</span><span
  class=\"o\">)</span>\n        statements\n        <span class=\"p\">;;</span>\n
  \   pattern <span class=\"m\">4</span><span class=\"o\">)</span>\n        statements\n
  \       <span class=\"p\">;;</span> \n    *<span class=\"o\">)</span>\n        statements\n
  \       <span class=\"p\">;;</span>\n<span class=\"k\">esac</span>\n</pre></div>\n\n</pre>\n\n<p>It
  follows a particular pattern if it matches it stops the search and executes the
  statements, finally comes out of the block. If it doesn't find any match it redirects
  to a default condition if any.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash </span>\n\n<span class=\"nb\">read</span> -p <span class=\"s2\">&quot;Enter
  a name : &quot;</span> n\n<span class=\"k\">case</span> <span class=\"nv\">$n</span>
  <span class=\"k\">in</span> \n\tadmin<span class=\"o\">)</span>\n\t\t<span class=\"nb\">echo</span>
  <span class=\"s2\">&quot;You are logged as root&quot;</span>\n\t\t<span class=\"p\">;;</span>\n\tunknown<span
  class=\"o\">)</span>\n\t\t<span class=\"nb\">echo</span> <span class=\"s2\">&quot;A
  hacker probably&quot;</span>\n\t\t<span class=\"p\">;;</span>\n\tmanager<span class=\"o\">)</span>\n\t\t<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;Weolcome Manager!&quot;</span>\n\t\t<span
  class=\"p\">;;</span>\n\t*<span class=\"o\">)</span>\n\t\t<span class=\"nb\">echo</span>
  <span class=\"s2\">&quot;A normal person&quot;</span>\n\t\t<span class=\"p\">;;</span>\n<span
  class=\"k\">esac</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625160454707/YDYGkU34d.png\"
  alt=\"bashs2.png\" /></p>\n<p>Seeing the above example, it is quite clear, that
  it looks quite structured and readable than the nested ladder of If-else statements.
  Cases are rendered based on the variable after the <code>case</code> keyword. We
  use the patterns before <code>)</code> as making the match in the variable provided
  before. Once the interpreter finds a match it returns to <code>esac</code> command
  which is <code>case</code> spelled in reverse just like <code>fi</code> for <code>if</code>
  and <code>done</code> for <code>do</code> in loops :) If it doesn't match any pattern,
  we have a default case represented by <code>*)</code> and it executes for any non-matching
  expression.</p>\n<h2>Arrays</h2>\n<p>Arrays or a way to store a list of numbers
  is implemented in Bash identical to most of the general programming languages.</p>\n<h3>Declaring
  and Printing Arrays</h3>\n<p>We declare an array similar to a variable but we mention
  the index of the element in the array(0 based index).  We can also simply declare
  an empty array using the declare command <code>declare -A nums</code></p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/bin/bash </span>\n\nnums<span class=\"o\">[</span><span class=\"m\">0</span><span
  class=\"o\">]=</span><span class=\"m\">7</span>\nnums<span class=\"o\">[</span><span
  class=\"m\">1</span><span class=\"o\">]=</span><span class=\"m\">5</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">2</span><span class=\"o\">]=</span><span class=\"m\">8</span>
  \n<span class=\"k\">for</span> i <span class=\"k\">in</span> <span class=\"si\">${</span><span
  class=\"nv\">nums</span><span class=\"p\">[@]</span><span class=\"si\">}</span>\n<span
  class=\"k\">do</span>\n<span class=\"nb\">echo</span> -e <span class=\"s2\">&quot;</span><span
  class=\"nv\">$i</span><span class=\"s2\"> \\n&quot;</span>\n<span class=\"k\">done</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625124595563/TzBEaH1E4.png\"
  alt=\"bashs2.png\" /></p>\n<p>The above script initializes an array with some hard-coded
  elements, surely you can input those from the user. For printing and accessing those
  elements in the array, We can use a loop, here we have used a range-based for loop.
  You can use any other loop you prefer. The iterator is &quot; i &quot; and we use
  $ to access the values from the array, we use <code>{}</code> as we have nested
  expression for indexing the element and <code>*</code> for every element in the
  array ( <code>@</code> will also work fine ), that's why range-based for loops make
  it quite simple to use. And we have simply printed &quot; i &quot; as it holds a
  particular element based on the iteration.</p>\n<p>OR</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">declare</span> -A <span class=\"nv\">nums</span><span
  class=\"o\">=(</span>\n<span class=\"o\">[</span><span class=\"m\">0</span><span
  class=\"o\">]=</span><span class=\"m\">44</span>\n<span class=\"o\">[</span><span
  class=\"m\">1</span><span class=\"o\">]=</span><span class=\"m\">45</span>\n<span
  class=\"o\">[</span><span class=\"m\">2</span><span class=\"o\">]=</span><span class=\"m\">46</span>\n<span
  class=\"o\">)</span>\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span
  class=\"si\">${</span><span class=\"nv\">nums</span><span class=\"p\">[@]</span><span
  class=\"si\">}</span><span class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625128343654/NCKUzurXe.png\"
  alt=\"bashs2.png\" /></p>\n<p>The above script uses declare an array, it can be
  empty as well after the name declaration. We used the <code>()</code> to include
  the values in the array, using indices in the array we assigned the values to the
  particular index.</p>\n<p>If you just want to print the elements, we can use <code>${nums[@]}</code>
  or <code>${nums[*]}</code>, this will print every element without using any iteration
  loops.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"ch\">#!/bin/bash</span>\n\nnums<span class=\"o\">[</span><span class=\"m\">0</span><span
  class=\"o\">]=</span><span class=\"m\">7</span>\nnums<span class=\"o\">[</span><span
  class=\"m\">1</span><span class=\"o\">]=</span><span class=\"m\">5</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">2</span><span class=\"o\">]=</span><span class=\"m\">8</span>
  \n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"si\">${</span><span
  class=\"nv\">nums</span><span class=\"p\">[@]</span><span class=\"si\">}</span><span
  class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625125456166/85zXjapQ_.png\"
  alt=\"bashs2.png\" /></p>\n<h3>Number of Elements in the array</h3>\n<p>To get the
  length of the array, we can use # in the expression <code>${nums[@]}</code>, like
  <code>${#nums[@]}</code> to get the number of elements from the array.</p>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625125770005/VzIr8CD7-.png\"
  alt=\"bashs2.png\" /></p>\n<p>Since we had 4 elements in the array, it accurately
  printed 4.</p>\n<h3>Inserting and Deleting elements from Array</h3>\n<p>We can push
  elements to the array using the assignment operator.</p>\n<p><code>nums=(${nums[@]}
  76) </code></p>\n<p>This will push 76 into the array, i.e. in the last index( length
  -1 index).</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  class=\"ch\">#!/bin/bash</span>\n\nnums<span class=\"o\">[</span><span class=\"m\">0</span><span
  class=\"o\">]=</span><span class=\"m\">7</span>\nnums<span class=\"o\">[</span><span
  class=\"m\">1</span><span class=\"o\">]=</span><span class=\"m\">5</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">2</span><span class=\"o\">]=</span><span class=\"m\">8</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">3</span><span class=\"o\">]=</span><span class=\"m\">19</span>\n<span
  class=\"nv\">nums</span><span class=\"o\">=(</span><span class=\"si\">${</span><span
  class=\"nv\">nums</span><span class=\"p\">[@]</span><span class=\"si\">}</span>
  <span class=\"m\">76</span><span class=\"o\">)</span>\n<span class=\"nb\">echo</span>
  <span class=\"s2\">&quot;</span><span class=\"si\">${</span><span class=\"nv\">nums</span><span
  class=\"p\">[@]</span><span class=\"si\">}</span><span class=\"s2\">&quot;</span>\n<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;Length of nums = </span><span
  class=\"si\">${#</span><span class=\"nv\">nums</span><span class=\"p\">[@]</span><span
  class=\"si\">}</span><span class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625126198675/A8IAE-2FR.png\"
  alt=\"bashs2.png\" /></p>\n<p>As you can see the element was added at the end of
  the array, you can also specify the index you want to insert. We can use <code>unset
  nums[3] </code> to delete the element at the particular location or we can pop back
  (delete from the end) an element from the array using the index <code>-1</code>
  from the array using the following command.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nb\">unset</span> nums<span class=\"o\">[</span>-1<span class=\"o\">]</span>\n</pre></div>\n\n</pre>\n\n<p>Any
  index provided will delete the element at that location by using unset. By using
  -1, we intend to refer to the last element. This can be quite handy and important
  as well in certain cases.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\nnums<span class=\"o\">[</span><span class=\"m\">0</span><span
  class=\"o\">]=</span><span class=\"m\">7</span>\nnums<span class=\"o\">[</span><span
  class=\"m\">1</span><span class=\"o\">]=</span><span class=\"m\">5</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">2</span><span class=\"o\">]=</span><span class=\"m\">8</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">3</span><span class=\"o\">]=</span><span class=\"m\">19</span>\n<span
  class=\"nb\">unset</span> nums<span class=\"o\">[</span>-1<span class=\"o\">]</span>\n<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"si\">${</span><span
  class=\"nv\">nums</span><span class=\"p\">[@]</span><span class=\"si\">}</span><span
  class=\"s2\">&quot;</span>\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;Length
  of nums = </span><span class=\"si\">${#</span><span class=\"nv\">nums</span><span
  class=\"p\">[@]</span><span class=\"si\">}</span><span class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625126770211/LYK2Q0Rp0.png\"
  alt=\"bashs2.png\" /></p>\n<p>There you can see we removed the element using unset.</p>\n<h3>Splicing
  an Array</h3>\n<p>We can splice the array to print/ copy a portion of the array
  to another one.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"si\">${</span><span
  class=\"nv\">nums</span><span class=\"p\">[@]:</span><span class=\"nv\">1</span><span
  class=\"p\">:</span><span class=\"nv\">3</span><span class=\"si\">}</span><span
  class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>Using two colons and numbers
  in between them, we can print in this case certain elements in the array from a
  particular range. Here the first number after the colon is the starting index to
  print from(inclusive) and the next number after the second colon is the length to
  which we would like to print the element, it is not the index but the number of
  elements after the start index to be spliced</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"ch\">#!/bin/bash</span>\n\nnums<span class=\"o\">[</span><span class=\"m\">0</span><span
  class=\"o\">]=</span><span class=\"m\">7</span>\nnums<span class=\"o\">[</span><span
  class=\"m\">1</span><span class=\"o\">]=</span><span class=\"m\">5</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">2</span><span class=\"o\">]=</span><span class=\"m\">8</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">3</span><span class=\"o\">]=</span><span class=\"m\">19</span>\nnums<span
  class=\"o\">[</span><span class=\"m\">4</span><span class=\"o\">]=</span><span class=\"m\">76</span>\n<span
  class=\"nv\">newarr</span><span class=\"o\">=</span><span class=\"si\">${</span><span
  class=\"nv\">nums</span><span class=\"p\">[@]:</span><span class=\"nv\">1</span><span
  class=\"p\">:</span><span class=\"nv\">3</span><span class=\"si\">}</span>\n<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"si\">${</span><span
  class=\"nv\">newarr</span><span class=\"p\">[@]</span><span class=\"si\">}</span><span
  class=\"s2\">&quot;</span>\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span
  class=\"si\">${</span><span class=\"nv\">nums</span><span class=\"p\">[@]</span><span
  class=\"si\">}</span><span class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625127387739/IH2Fc6ghk.png\"
  alt=\"bashs2.png\" /></p>\n<p>In this case, we have copied the slice of an array
  to another new array using the double colon operator. We added the elements from
  the 1st index till <code>1+3</code> indices i.e till 4. 3 is not the index but the
  length till you would like to copy or print.</p>\n<p>This was a basic introduction
  to arrays, definitely, there will be much more stuff I didn't cover. Just to give
  an overview of how an array looks like in BASH scripting. Next, we move on to strings.</p>\n<h1>Strings</h1>\n<p>Strings
  are quite important as it forms the core of any script to deal with filenames, user
  information, etc all contain strings or array of characters. Let's take a closer
  look at how strings are declared, handled, and manipulated in Bash scripting.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nv\">s</span><span class=\"o\">=</span><span class=\"s2\">&quot;World&quot;</span>\n<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"nv\">$s</span><span
  class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625129318948/02V3bpP1I.png\"
  alt=\"bashs2.png\" /></p>\n<p>Strings are again declared as normal variables but
  are enclosed in double quotation marks.  And we access them in the exact same way
  as we do with variables. If you were to use single quotes instead of double quotes
  Bash would not interpret the variable name as a variable, it would print the name
  literally and not the value of the variable, So prefer using double quotes in echo
  and other commands to make variables accessible.</p>\n<h3>Making Substrings</h3>\n<p>We
  can even splice the string as we did with the arrays, in strings we can call it
  substrings. The syntax is almost identical as we just have to get the variable name.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nv\">s</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;Hello World&quot;</span>\n<span class=\"nv\">p</span><span class=\"o\">=</span><span
  class=\"si\">${</span><span class=\"nv\">s</span><span class=\"p\">:</span><span
  class=\"nv\">6</span><span class=\"si\">}</span>\n<span class=\"nb\">echo</span>
  <span class=\"nv\">$p</span>\n<span class=\"nv\">q</span><span class=\"o\">=</span><span
  class=\"si\">${</span><span class=\"nv\">s</span><span class=\"p\">::</span><span
  class=\"nv\">5</span><span class=\"si\">}</span>\n<span class=\"nb\">echo</span>
  <span class=\"nv\">$q</span>\n<span class=\"nv\">t</span><span class=\"o\">=</span><span
  class=\"si\">${</span><span class=\"nv\">s</span><span class=\"p\">:</span><span
  class=\"nv\">4</span><span class=\"p\">:</span><span class=\"nv\">3</span><span
  class=\"si\">}</span>\n<span class=\"nb\">echo</span> <span class=\"nv\">$t</span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above script, we had a base string 's' which was then sliced from the 6th index
  to the end, If we do not pass the second number and colon, it interprets as the
  end of the string and if we do not pass the first number, it will interpret as the
  first character in the string. We sliced s from the 6th index till the end of the
  string and copied it in the string 'p''. In the 'q' string, we sliced the first
  5 characters from the string 's'. In the 't' string we sliced starting from the
  4th index and 3 characters in length i.e till  7th index, not the 7th index.</p>\n<p>We
  can use the # before the variable name to get the length of the string as we saw
  in the array section. So we can use the <code>echo ${#s}</code> command to print
  the length of the string where s is the string variable name.</p>\n<h3>String Concatenation</h3>\n<p>String
  concatenation on Bash is quite straightforward as it is just the matter of adding
  strings in a very simple way.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nv\">s</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;Hello&quot;</span>\n<span class=\"nv\">p</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;World&quot;</span>\n<span class=\"nv\">q</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;</span><span class=\"nv\">$s</span><span class=\"s2\"> </span><span
  class=\"nv\">$p</span><span class=\"s2\">&quot;</span>\n<span class=\"nb\">echo</span>
  <span class=\"nv\">$q</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625135997314/_n05RIoTTM.png\"
  alt=\"bashs2.png\" /></p>\n<p>The space between the two variables is quite literal,
  anything you place between this space or the double quotes will get stored in the
  variable or get printed.</p>\n<h3>String Comparison</h3>\n<p>Moving on to the string
  comparison in Bash. String comparison is quite complex in certain programming languages
  but it's quite straightforward in some languages such as Bash. We can compare strings
  quite easily in Bash, either they are equal or they are not, it's just comparison
  operators to perform the heavy-lifting for us.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nv\">s</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;hello&quot;</span>\n<span class=\"nv\">p</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;Hello&quot;</span>\n<span class=\"k\">if</span> <span class=\"o\">[</span>
  <span class=\"nv\">$s</span> <span class=\"o\">=</span> <span class=\"nv\">$p</span>
  <span class=\"o\">]</span><span class=\"p\">;</span>\n<span class=\"k\">then</span>\n
  \   <span class=\"nb\">echo</span> <span class=\"s2\">&quot;Equal&quot;</span>\n<span
  class=\"k\">else</span> \n    <span class=\"nb\">echo</span> <span class=\"s2\">&quot;Not
  equal&quot;</span>\n<span class=\"k\">fi</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625138020711/UWlRN8aPq.png\"
  alt=\"bashs2.png\" /></p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625137884154/7bNPkpGd3.png\"
  alt=\"bashs2.png\" /></p>\n<p>From the above code, it is quite clear that the strings
  as not equal and we compared them with the &quot;equality&quot; operator (=) and
  checked if that condition was true, and perform commands accordingly. We can also
  check if the strings are not equal using <code>!=</code> operator and we can perform
  commands based on the desired logic. We also have operators to compare the length
  of the strings. We can use <code>\\&lt;</code> operator to check if the first string
  is less than the second string(compared characters in ASCII).  And check if the
  first string is greater than the second string using <code>\\&gt;</code> operator.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nv\">s</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;hello&quot;</span>\n<span class=\"nv\">p</span><span class=\"o\">=</span><span
  class=\"s2\">&quot;Hello&quot;</span>\n<span class=\"k\">if</span> <span class=\"o\">[</span>
  <span class=\"nv\">$s</span> <span class=\"se\">\\&lt;</span> <span class=\"nv\">$p</span>
  <span class=\"o\">]</span><span class=\"p\">;</span>\n<span class=\"k\">then</span>\n\n\t<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"nv\">$s</span><span
  class=\"s2\"> is Less than </span><span class=\"nv\">$p</span><span class=\"s2\">&quot;</span>\n<span
  class=\"k\">elif</span> <span class=\"o\">[</span> <span class=\"nv\">$s</span>
  <span class=\"se\">\\&gt;</span> <span class=\"nv\">$p</span> <span class=\"o\">]</span><span
  class=\"p\">;</span>\n<span class=\"k\">then</span>\n\t<span class=\"nb\">echo</span>
  <span class=\"s2\">&quot;</span><span class=\"nv\">$s</span><span class=\"s2\">
  is greater than </span><span class=\"nv\">$p</span><span class=\"s2\">&quot;</span>\n<span
  class=\"k\">else</span>\n<span class=\"nb\">echo</span> <span class=\"s2\">&quot;Equal&quot;</span>\n<span
  class=\"k\">fi</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625137393683/u3WbgDIrN.png\"
  alt=\"bashs2.png\" /></p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625137467448/NP1UXZAbv.png\"
  alt=\"bashs2.png\" /></p>\n<p>Here, we are using the ASCII equivalent of strings
  to compare them as it gives an idea in terms of the value of the strings. We see
  that 'h'( 104)has a greater ASCII value than 'H&quot; (72) which is why we see the
  shown outcome.</p>\n<p>We also have operators to check for the string being empty
  or not using the -z operator. Also, we have arguments to pass to the string comparison
  to check for non-empty strings as well, specifically for input validation and some
  error handling.</p>\n<p>We can quite easily use -n to check for non-empty string
  and -z for the length of the string being zero.</p>\n<pre class='wrapper'>\n\n<div
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">read</span> -p <span class=\"s2\">&quot;Enter
  a string : &quot;</span> s\n<span class=\"k\">if</span> <span class=\"o\">[</span>
  -s <span class=\"nv\">$s</span> <span class=\"o\">]</span><span class=\"p\">;</span>\n<span
  class=\"k\">then</span> \n    <span class=\"nb\">echo</span> <span class=\"s2\">&quot;Empty
  Input&quot;</span>\n<span class=\"k\">else</span>\n   <span class=\"nb\">echo</span>
  <span class=\"s2\">&quot;Valid input&quot;</span>\n<span class=\"k\">fi</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625138907051/tbjRDda0U1.png\"
  alt=\"bashs2.png\" /></p>\n<p>So the string topic is quite straightforward and self-explanatory
  as it doesn't how that much complexity but is still powerful and convenient to use.</p>\n<h1>Arithmetic
  in Bash</h1>\n<p>Performing any Arithmetic operations is the core for scripting.
  Without arithmetic, it feels incomplete to programmatically create something, it
  would be quite menial to write commands by hand without having the ability to perform
  arithmetic operations.</p>\n<h3>Integer Arithmetic</h3>\n<p>Firstly we quite commonly
  use operations on variables, so let us see how to perform an arithmetic operation
  on variables in Bash. We use double curly braces to evaluate certain results of
  the operations performed on variables.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nv\">x</span><span class=\"o\">=</span><span
  class=\"m\">4</span>\n<span class=\"nv\">y</span><span class=\"o\">=</span><span
  class=\"m\">9</span>\n<span class=\"nv\">z</span><span class=\"o\">=</span><span
  class=\"k\">$((</span><span class=\"nv\">$x</span> <span class=\"o\">*</span> <span
  class=\"nv\">$y</span><span class=\"k\">))</span>\n<span class=\"nb\">echo</span>
  <span class=\"nv\">$z</span>\n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625139582135/Sb4hdE990.png\"
  alt=\"bashs2.png\" /></p>\n<p>We use double curly braces in order to evaluate the
  operations performed on the variables inside them. We definitely have to use the
  $ symbol to extract the value of the variable.</p>\n<p>We can surely use operators
  such as addition(<code>+</code>), subtraction(<code>-</code>), multiplication(<code>*</code>),
  division(<code>/</code>), and modulus(<code>%</code>, it stores the remainder of
  the division,17%3 gets you 2) in statements. We can also perform operations such
  as <code>&lt;&lt;</code> to do left bitwise shift and <code>&gt;&gt;</code> right
  bitwise shift to shift the binary digits in left tor right respectively in a variable.
  There are also logical operations such as Bitwise and logical AND(<code>&amp;</code>),
  OR(<code>|</code>), EX-OR(<code>^</code>), and ternary expressions.</p>\n<p>Alternative
  to double curly braces is <code>expr</code>, expr allows you to freely wherever
  you need to evaluate an arithmetic operation. Though this is not native in-built
  in shells, it uses a binary process to evaluate the arithmetic operations. It can
  also defer depending on the implementation of such commands in various environments.</p>\n<p>We
  can also use the <code>let</code> command to initialize a variable and perform expressions
  in the initialization itself.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">let</span> <span class=\"nv\">a</span><span
  class=\"o\">=</span><span class=\"m\">4</span>\n<span class=\"nb\">let</span> <span
  class=\"nv\">b</span><span class=\"o\">=</span>a*a\n<span class=\"nb\">let</span>
  <span class=\"nv\">c</span><span class=\"o\">=</span><span class=\"s2\">&quot;b/(a*2)&quot;</span>\n<span
  class=\"nb\">echo</span> <span class=\"nv\">$b</span>\n</pre></div>\n\n</pre>\n\n<p>We
  can perform quite complex operations using simple implementation using <code>let</code>,
  this allows much readable and bug-free scripts.  If you would like to perform operations
  with brackets and other operations you can enclose the expression in double quotation
  marks.</p>\n<h3>Floating-Point Arithmetic</h3>\n<p>Performing floating-point arithmetic
  in Bash is not quite well though. We won't get 100% accurate answers in the expressions
  this is because it is <strong>not designed</strong> for such things. Doing <strong>things
  related to floating-point is a bad idea</strong>, Still, you can improve the precision
  to a little extent to do some basic things. I <strong>don't recommend this</strong>
  only do this if there are no other options.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  class=\"nb\">printf</span> %.9f <span class=\"s2\">&quot;</span><span class=\"k\">$((</span><span
  class=\"m\">10</span><span class=\"o\">/</span><span class=\"m\">3</span><span class=\"k\">))</span><span
  class=\"s2\"></span>\n</pre></div>\n\n</pre>\n\n<p>The result of this is 3.0000000..064
  roughly, which is pretty bad. Bash at its core doesn't support floating-point calculations.
  But there is good news, we have  <a href=\"https://en.wikipedia.org/wiki/AWK\">awk</a>
  \ and other tools such as  <a href=\"https://en.wikipedia.org/wiki/Bc_(programming_language)\">bc</a>
  \ and others which is planned for the next part in the series. I'll explain awk
  just for floating-point here, in the next part, I'll cover it in depth.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/bin/bash</span>\n<span class=\"nv\">a</span><span class=\"o\">=</span><span
  class=\"m\">10</span>\n<span class=\"nv\">b</span><span class=\"o\">=</span><span
  class=\"k\">$(</span><span class=\"nb\">echo</span> <span class=\"p\">|</span> awk
  -v <span class=\"nv\">a</span><span class=\"o\">=</span><span class=\"s2\">&quot;</span><span
  class=\"nv\">$a</span><span class=\"s2\">&quot;</span> <span class=\"s1\">&#39;{print
  a/3}&#39;</span><span class=\"k\">)</span>\n<span class=\"nb\">echo</span> <span
  class=\"nv\">$b</span> \n</pre></div>\n\n</pre>\n\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625157391350/gHudsNntM4.png\"
  alt=\"bashs2.png\" /></p>\n<p>WOW! That is to the point, but that was a lot of hassle
  using echo but printing nothing! HH? OK, you see certain things can get really annoying
  when things aren't supported natively. Firstly, we use  | to pipe echo command with
  awk, the echo command doesn't do anything just a way to use awk command in assigning
  variables here. Then the general syntax of the awk command is <code> awk -options
  -commands</code>. In this case, we are using -v as an argument and passing in an
  as a variable which is equal to a, which is stupid and silly but that is what it
  is, you can name any variable name you want. Then we simply have to use the variable
  in the print function which generally evaluates the expressions or other operations
  and returns to the interpreter. And that is how we print the expression, Phew! That
  took a while to do some silly things, But hey! That's possible though.</p>\n<p>That
  is the basic overview of Arithmetic in Bash, you can also perform logical operations
  in it which is very easy and can be understood on a quick run-through in the  <a
  href=\"https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html#Arithmetic-Expansion\">documentation</a>.</p>\n<p>I
  hope you understood the mentioned topics and what are their use cases depending
  on the requirements. Some topics such as positional parameters, tools and utilities,
  dictionaries, and some other important aspects of Bash scripting will be covered
  in the next part. Happy Coding.</p>\n"
image_url: https://meetgor-cdn.pages.dev/bash-scripting-guide-p2.webp
long_description: 'In this part, topics such as switch cases, mathematical expression,
  arrays, and strings. This won Topics to be covered in this part are as follows:
  User Input User Prompt Changing the Delimiter Password as Input Limiting the length
  of Input Cases Arr'
now: 2025-05-01 18:11:33.314469
path: blog/posts/2021-07-02-BASH-scripting-P2.md
prevnext: null
series:
- BASH Scripting Guide
slug: bash-guide-p2
status: published
subtitle: Parsing and Managing data using Bash
tags:
- bash
templateKey: blog-post
title: BASH Scripting Guide - PART - 2
today: 2025-05-01
---

# Introduction

In this part, topics such as switch cases, mathematical expression, arrays, and strings. This won't be an in-depth guide to understand each and every concept, but to make users aware of the things and features in Bash. This also would not be an absolute basic guide, I expect to have some basic programming knowledge such as binary systems, logical and mathematical concepts. Don't worry, you won't be bombarded with commands, I'll just explain with easy examples to get started.

Topics to be covered in this part are as follows:
- User Input

    - User Prompt
    - Changing the Delimiter
    - Password as Input
    - Limiting the length of Input

- Cases

- Arrays
    - Declaring and Printing Arrays
    - Number of elements in an array
    - Splicing the array
    - Inserting and Deleting elements

- Strings
    - Making Substrings
    - String Concatenation
    - String Comparison
- Arithmetic in Bash
    - Integer Arithmetic 
    - Floating-Point Arithmetic

# User Input

Taking user input in Bash is quite straightforward and quite readable as well. We can make use of `read` command to take in input from the user. We just specify the variable in which we want to store the input.` read x` Here, the input will be stored in x. We can also pass in certain arguments to the read command such as -p (prompt with string), -r ( delimiter variation), -a(pass to the array), and others as well. Each of them will make the foundation of various complicated tasks to perform in logical operations. 

### User prompt argument
The -p argument will prompt the user with a string before they input anything. It makes quite informative and useful user input. This becomes quite a useful argument/parameter to make it quite readable and understand what to do directly with much hassle. The below is the general syntax of passing the argument to the read function.

```bash
read -p "Enter the number " n
```

```bash
#!/bin/bash

read -p "Enter the number " n
echo "The inputted number was $n"
```
![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625118915300/NRF7Ci2rK.png)

In this example, we can prompt the user with the string **Enter the number**, and it gives certain information to the user about what to input. 


### Changing the delimiter
Next, we can make use of -r which depending on the use case, we can change the delimiter while taking the input.

```bash
#!/bin/bash

IFS='/' read -p "Enter the file path : " user project app 
echo $user $project $app

```

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625202319183/R9Eo3lg0oU.png)

In the above script, we separated the path of a directory(user-entered) into separate components such as the username, project name, and the app name, this can get pretty important and a great tool for automation of making project and application structures. At the beginning of the command, we use IFS which stands for Internal Field Separator, which does the separation of variables based on the field provided, in this case it was `//`, you can use any other field characters appropriate to your needs.

This command will change the delimiter, by default it uses spaces or tabs etc to identify distinct input variables but we change it to other internal field separators such as / , \ ,- , \|, etc. This can make the user input more customizable and dynamic. 

### Password Typing
We can hide the input from the screen so as to keep it confidential and keep sensitive information such as passwords and keys private and protected. 
```bash
#!/bin/bash

read -sp "Password: " pswd
echo "the password was $pswd"

```

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625161571993/MkCadbyKW.png)


We used the -s to keep the input hidden, the screen doesn't reflect what the user is typing, and -p for the prompt to offer the user some information on the text. 

### Limiting Length of Input

We can limit the user to only a certain number of characters as input. It becomes quite useful in constrained environments such as usernames and passwords to be restricted with a certain limit. 

```bash
#!/bin/bash

read -n 6 -p "Enter the name: " n
echo $n
```

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625161752440/8xI5Lovxo.png)

In the above script, we demonstrate the limit of characters of 6 in the variable n. This restricts the user with only the first 6 characters, it just doesn't exceed ahead, directly to the next command.

### Passing to the array

Another important argument to be passed after read command is -a which inserts the value to the array elements.
 
```bash
#!/bin/bash

read -a nums -p "Enter the elements : " 
for i in ${nums[*]};
do 
	echo -e "$i\n"
done

```
In the above script, we have used array, don't worry, I'll explain it in the coming sections of this part. We have assigned an empty array and made the user enter those arrays, they are space-separated values. We have used the -a operator to insert the input to the elements of the array. The for loop is range-based which means it will do certain commands until there are no elements left in nums. The command `${nums[@]}` indicates every element in the array nums. 
  
# Cases

Cases are quite a good way of replacing nested if-else statements to make them nice and readable in the script.  Cases in Bash are quite powerful and easy to use compared with C/ C++ styled switch cases. 

The general structure of using a case in Bash is as follows:

```bash
case variable in
    pattern 1)
        statements
        ;;
    pattern 2)
        statements
        ;;
    pattern 3)
        statements
        ;;
    pattern 4)
        statements
        ;; 
    *)
        statements
        ;;
esac
```

It follows a particular pattern if it matches it stops the search and executes the statements, finally comes out of the block. If it doesn't find any match it redirects to a default condition if any. 

```bash
#!/bin/bash 

read -p "Enter a name : " n
case $n in 
	admin)
		echo "You are logged as root"
		;;
	unknown)
		echo "A hacker probably"
		;;
	manager)
		echo "Weolcome Manager!"
		;;
	*)
		echo "A normal person"
		;;
esac

```

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625160454707/YDYGkU34d.png)

Seeing the above example, it is quite clear, that it looks quite structured and readable than the nested ladder of If-else statements. Cases are rendered based on the variable after the `case` keyword. We use the patterns before `)` as making the match in the variable provided before. Once the interpreter finds a match it returns to `esac` command which is `case` spelled in reverse just like `fi` for `if` and `done` for `do` in loops :) If it doesn't match any pattern, we have a default case represented by `*)` and it executes for any non-matching expression. 


## Arrays

Arrays or a way to store a list of numbers is implemented in Bash identical to most of the general programming languages. 

### Declaring and Printing Arrays 

We declare an array similar to a variable but we mention the index of the element in the array(0 based index).  We can also simply declare an empty array using the declare command `declare -A nums`

```bash
#!/bin/bash 

nums[0]=7
nums[1]=5
nums[2]=8 
for i in ${nums[@]}
do
echo -e "$i \n"
done
```

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625124595563/TzBEaH1E4.png)

The above script initializes an array with some hard-coded elements, surely you can input those from the user. For printing and accessing those elements in the array, We can use a loop, here we have used a range-based for loop. You can use any other loop you prefer. The iterator is " i " and we use $ to access the values from the array, we use `{}` as we have nested expression for indexing the element and `*` for every element in the array ( `@` will also work fine ), that's why range-based for loops make it quite simple to use. And we have simply printed " i " as it holds a particular element based on the iteration. 

OR

```bash
#!/bin/bash

declare -A nums=(
[0]=44
[1]=45
[2]=46
)
echo "${nums[@]}"

```

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625128343654/NCKUzurXe.png)

The above script uses declare an array, it can be empty as well after the name declaration. We used the `()` to include the values in the array, using indices in the array we assigned the values to the particular index.

If you just want to print the elements, we can use `${nums[@]}` or `${nums[*]}`, this will print every element without using any iteration loops.

```bash
#!/bin/bash

nums[0]=7
nums[1]=5
nums[2]=8 
echo "${nums[@]}"

```

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625125456166/85zXjapQ_.png)

### Number of Elements in the array

To get the length of the array, we can use # in the expression `${nums[@]}`, like `${#nums[@]}` to get the number of elements from the array.

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625125770005/VzIr8CD7-.png)

Since we had 4 elements in the array, it accurately printed 4. 

### Inserting and Deleting elements from Array 

We can push elements to the array using the assignment operator. 

`nums=(${nums[@]} 76) `

This will push 76 into the array, i.e. in the last index( length -1 index). 

```bash
#!/bin/bash

nums[0]=7
nums[1]=5
nums[2]=8
nums[3]=19
nums=(${nums[@]} 76)
echo "${nums[@]}"
echo "Length of nums = ${#nums[@]}"

```

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625126198675/A8IAE-2FR.png)

As you can see the element was added at the end of the array, you can also specify the index you want to insert. We can use `unset nums[3] ` to delete the element at the particular location or we can pop back (delete from the end) an element from the array using the index `-1` from the array using the following command.

```bash
unset nums[-1]
```

Any index provided will delete the element at that location by using unset. By using -1, we intend to refer to the last element. This can be quite handy and important as well in certain cases.

```bash
#!/bin/bash

nums[0]=7
nums[1]=5
nums[2]=8
nums[3]=19
unset nums[-1]
echo "${nums[@]}"
echo "Length of nums = ${#nums[@]}"

```

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625126770211/LYK2Q0Rp0.png)

There you can see we removed the element using unset. 

### Splicing an Array

We can splice the array to print/ copy a portion of the array to another one. 

```bash
echo "${nums[@]:1:3}"
```

Using two colons and numbers in between them, we can print in this case certain elements in the array from a particular range. Here the first number after the colon is the starting index to print from(inclusive) and the next number after the second colon is the length to which we would like to print the element, it is not the index but the number of elements after the start index to be spliced

```bash
#!/bin/bash

nums[0]=7
nums[1]=5
nums[2]=8
nums[3]=19
nums[4]=76
newarr=${nums[@]:1:3}
echo "${newarr[@]}"
echo "${nums[@]}"

```

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625127387739/IH2Fc6ghk.png)

In this case, we have copied the slice of an array to another new array using the double colon operator. We added the elements from the 1st index till `1+3` indices i.e till 4. 3 is not the index but the length till you would like to copy or print. 

This was a basic introduction to arrays, definitely, there will be much more stuff I didn't cover. Just to give an overview of how an array looks like in BASH scripting. Next, we move on to strings.

# Strings

Strings are quite important as it forms the core of any script to deal with filenames, user information, etc all contain strings or array of characters. Let's take a closer look at how strings are declared, handled, and manipulated in Bash scripting. 

```bash
s="World"
echo "$s"
```

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625129318948/02V3bpP1I.png)

Strings are again declared as normal variables but are enclosed in double quotation marks.  And we access them in the exact same way as we do with variables. If you were to use single quotes instead of double quotes Bash would not interpret the variable name as a variable, it would print the name literally and not the value of the variable, So prefer using double quotes in echo and other commands to make variables accessible.

### Making Substrings

We can even splice the string as we did with the arrays, in strings we can call it substrings. The syntax is almost identical as we just have to get the variable name.

```bash
#!/bin/bash

s="Hello World"
p=${s:6}
echo $p
q=${s::5}
echo $q
t=${s:4:3}
echo $t

```

In the above script, we had a base string 's' which was then sliced from the 6th index to the end, If we do not pass the second number and colon, it interprets as the end of the string and if we do not pass the first number, it will interpret as the first character in the string. We sliced s from the 6th index till the end of the string and copied it in the string 'p''. In the 'q' string, we sliced the first 5 characters from the string 's'. In the 't' string we sliced starting from the 4th index and 3 characters in length i.e till  7th index, not the 7th index.

We can use the # before the variable name to get the length of the string as we saw in the array section. So we can use the `echo ${#s}` command to print the length of the string where s is the string variable name. 

### String Concatenation

String concatenation on Bash is quite straightforward as it is just the matter of adding strings in a very simple way. 

```bash
#!/bin/bash

s="Hello"
p="World"
q="$s $p"
echo $q
```  

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625135997314/_n05RIoTTM.png)

The space between the two variables is quite literal, anything you place between this space or the double quotes will get stored in the variable or get printed.

### String Comparison

Moving on to the string comparison in Bash. String comparison is quite complex in certain programming languages but it's quite straightforward in some languages such as Bash. We can compare strings quite easily in Bash, either they are equal or they are not, it's just comparison operators to perform the heavy-lifting for us. 

```bash
#!/bin/bash

s="hello"
p="Hello"
if [ $s = $p ];
then
    echo "Equal"
else 
    echo "Not equal"
fi

```

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625138020711/UWlRN8aPq.png)

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625137884154/7bNPkpGd3.png)

From the above code, it is quite clear that the strings as not equal and we compared them with the "equality" operator (=) and checked if that condition was true, and perform commands accordingly. We can also check if the strings are not equal using `!=` operator and we can perform commands based on the desired logic. We also have operators to compare the length of the strings. We can use `\<` operator to check if the first string is less than the second string(compared characters in ASCII).  And check if the first string is greater than the second string using `\>` operator. 

```bash
#!/bin/bash

s="hello"
p="Hello"
if [ $s \< $p ];
then

	echo "$s is Less than $p"
elif [ $s \> $p ];
then
	echo "$s is greater than $p"
else
echo "Equal"
fi

```

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625137393683/u3WbgDIrN.png)


![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625137467448/NP1UXZAbv.png)

Here, we are using the ASCII equivalent of strings to compare them as it gives an idea in terms of the value of the strings. We see that 'h'( 104)has a greater ASCII value than 'H" (72) which is why we see the shown outcome. 

We also have operators to check for the string being empty or not using the -z operator. Also, we have arguments to pass to the string comparison to check for non-empty strings as well, specifically for input validation and some error handling. 

We can quite easily use -n to check for non-empty string and -z for the length of the string being zero.

```bash
#!/bin/bash

read -p "Enter a string : " s
if [ -s $s ];
then 
    echo "Empty Input"
else
   echo "Valid input"
fi

```  

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625138907051/tbjRDda0U1.png)

So the string topic is quite straightforward and self-explanatory as it doesn't how that much complexity but is still powerful and convenient to use.

# Arithmetic in Bash

Performing any Arithmetic operations is the core for scripting. Without arithmetic, it feels incomplete to programmatically create something, it would be quite menial to write commands by hand without having the ability to perform arithmetic operations. 

### Integer Arithmetic

Firstly we quite commonly use operations on variables, so let us see how to perform an arithmetic operation on variables in Bash. We use double curly braces to evaluate certain results of the operations performed on variables. 

```bash
#!/bin/bash

x=4
y=9
z=$(($x * $y))
echo $z

```

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625139582135/Sb4hdE990.png)

We use double curly braces in order to evaluate the operations performed on the variables inside them. We definitely have to use the $ symbol to extract the value of the variable. 

We can surely use operators such as addition(`+`), subtraction(`-`), multiplication(`*`), division(`/`), and modulus(`%`, it stores the remainder of the division,17%3 gets you 2) in statements. We can also perform operations such as `<<` to do left bitwise shift and `>>` right bitwise shift to shift the binary digits in left tor right respectively in a variable. There are also logical operations such as Bitwise and logical AND(`&`), OR(`|`), EX-OR(`^`), and ternary expressions.  

Alternative to double curly braces is `expr`, expr allows you to freely wherever you need to evaluate an arithmetic operation. Though this is not native in-built in shells, it uses a binary process to evaluate the arithmetic operations. It can also defer depending on the implementation of such commands in various environments. 

We can also use the `let` command to initialize a variable and perform expressions in the initialization itself. 

```bash
#!/bin/bash

let a=4
let b=a*a
let c="b/(a*2)"
echo $b
```

We can perform quite complex operations using simple implementation using `let`, this allows much readable and bug-free scripts.  If you would like to perform operations with brackets and other operations you can enclose the expression in double quotation marks. 

### Floating-Point Arithmetic

Performing floating-point arithmetic in Bash is not quite well though. We won't get 100% accurate answers in the expressions this is because it is **not designed** for such things. Doing **things related to floating-point is a bad idea**, Still, you can improve the precision to a little extent to do some basic things. I **don't recommend this** only do this if there are no other options. 

```bash
printf %.9f "$((10/3))
```

The result of this is 3.0000000..064 roughly, which is pretty bad. Bash at its core doesn't support floating-point calculations. But there is good news, we have  [awk](https://en.wikipedia.org/wiki/AWK)  and other tools such as  [bc](https://en.wikipedia.org/wiki/Bc_(programming_language))  and others which is planned for the next part in the series. I'll explain awk just for floating-point here, in the next part, I'll cover it in depth. 

```bash
#!/bin/bash
a=10
b=$(echo | awk -v a="$a" '{print a/3}')
echo $b 

```

![bashs2.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625157391350/gHudsNntM4.png)

WOW! That is to the point, but that was a lot of hassle using echo but printing nothing! HH? OK, you see certain things can get really annoying when things aren't supported natively. Firstly, we use  | to pipe echo command with awk, the echo command doesn't do anything just a way to use awk command in assigning variables here. Then the general syntax of the awk command is ` awk -options -commands`. In this case, we are using -v as an argument and passing in an as a variable which is equal to a, which is stupid and silly but that is what it is, you can name any variable name you want. Then we simply have to use the variable in the print function which generally evaluates the expressions or other operations and returns to the interpreter. And that is how we print the expression, Phew! That took a while to do some silly things, But hey! That's possible though. 

That is the basic overview of Arithmetic in Bash, you can also perform logical operations in it which is very easy and can be understood on a quick run-through in the  [documentation](https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html#Arithmetic-Expansion).

I hope you understood the mentioned topics and what are their use cases depending on the requirements. Some topics such as positional parameters, tools and utilities, dictionaries, and some other important aspects of Bash scripting will be covered in the next part. Happy Coding.