---
article_html: "<h2>Introduction</h2>\n<p>Web Scraping is quite an interesting and
  powerful tool or skill to have in a Programmer's toolkit.  It helps in analyzing
  data and getting some information in various formats. Web Scraping is a process
  in which a user fetches a website's content using some pattern in those HTML tags
  and the desired content to be fetched or scraped.</p>\n<p>For this article, we aim
  to fetch the meaning of a word entered by the user from the <a href=\"http://dictionary.com\">dictionary.com</a>
  website. We need to print just the meaning of the word from the HTML tags in it.
  We must have a good understanding of HTML and some basic Linux tools such as cURL,
  grep, sed, and others for doing all of these.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625737499658/FGLusWSII.png\"
  alt=\"Inspecting the Target Website\" /></p>\n<h2>Inspecting the Target Website</h2>\n<p>To
  begin with, scrapping the website, first, it is absolutely important to inspect
  the website and view its source code. For that, we can make use of Inspect tool
  in our Browsers. Just Right-click on the website you are viewing or the website
  for scraping, a list of options appears in front of you. You have to select Inspect
  option( also Shift + Ctrl + I), this will open a side window with a plethora of
  options. You simply have to select Elements from the top of the menus. The code
  that you will see is the source code of the website. No, don't think you can change
  the content of the website from here :)</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625737510444/KonUrEpcq-.png\"
  alt=\"image.png\" />\nInspect Tool in the browser.</p>\n<p>Now we have to analyze
  the website with the content which we want to scrape. You can go on for clicking
  the <code>select the element in the page to inspect it</code> option or icon in
  the top left-hand side corner. This will allow you to inspect the particular element
  that you selected on the webpage. You can now see the element tag, id, class, and
  other attributes required to fetch the element's content.</p>\n<h2>Selecting the
  particular element from the website to view the source code.</h2>\n<h3>Accessing
  the website from the Command line/terminal</h3>\n<p>Now the website structure is
  being understood we can actually move to scrap it. For that, we need to have the
  web site's content on our local machine. First of all, we need to access the website
  from elsewhere not from the browser, because you cannot copy-paste content from
  there. So let's use Command Line here. We have a popular tool known as <code>cURL</code>,
  which stands for client URL. The tool fetches the contents of the provided URL.
  It also has several parameters or arguments that can be used to modify its output.
  We can use the command</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>curl -o meaning.txt https://www.dictionary.com/browse/computer#\n</pre></div>\n\n</pre>\n\n<p>The
  above command fetches the HTML page for the word Computer, it could be any word
  you might be searching for.</p>\n<h3>Understanding the Website Structure.</h3>\n<p>Here
  comes the time to explain the structure of <a href=\"http://dictionary.com\">dictionary.com</a>.
  When you search a word on the website(<a href=\"http://dictionary.com\">dictionary.com</a>),
  you are routed to <code>/browse</code> which then fetches the word for you and defaults
  you to the <code>/browse/word#</code> (the word can be any word you searched). The
  curl command dumps the output in the <code>meaning.txt</code> or any specified file.
  If you see the contents of the file, it is the same as on the web.  So we are going
  to store the meaning of the searched word in the meaning.txt file, you can customize
  the name and command however you like.</p>\n<p>Voila! you successfully scraped a
  webpage. Now the next target is to filter the webpage content.</p>\n<h3>Filtering
  Content from Website local file</h3>\n<p>Now we have the content of the webpage
  on our local machine, we need to search or filter out the useful content and remove
  the unwanted tags and elements. For that, we can use commands such as <code>grep</code>
  and <code>sed</code>.</p>\n<h3>Finding Tags to Extract content.</h3>\n<p>We need
  to find patterns and similarities in the tags that contain the text of the meaning
  of the specified word. From the analysis of the webpage, we see that the element
  <code>&lt;span class=&quot;one-click-content css-nnyc96 e1q3nk1v1&quot;&gt;</code>
  contains the actual meaning. We just need the basic meaning, we may not need examples
  and long lengthy definitions on our Terminal, So we will go with filtering out the
  span tag with a class called <code>one-click-content css-nnyc96 e1q3nk1v1</code>.
  To do that we can use the grep command, which can print the text or line matching
  the specified expression or text. Here we need the span element with the particular
  class name so we will use regular expressions to find it more effectively.</p>\n<pre
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>grep
  -oP <span class=\"s1\">&#39;(?&lt;=&lt;span class=&quot;one-click-content css-nnyc96
  e1q3nk1v1&quot;&gt;).*?(?=&lt;/span&gt;)&#39;</span> meaning.txt &gt;temp.txt \n</pre></div>\n\n</pre>\n\n<h3>Using
  GREP command to filter.</h3>\n<p>The above command will search and return only lines
  that are contained in the span tags with that particular class name from the meaning.txt
  file which we appended to fill the webpage's source code. The <code>-oP</code> are
  the arguments that return Only the matching cases and <code>-P</code> the coming
  expression is a Perl Regex. The command will return everything in between those
  tags. Finally, we are storing the result or output in <code>temp.txt</code>.</p>\n<p>Now,
  if you think we are done, then it's wrong, the webpage can have internal or external
  links embedded inside of the elements as well, so we need to again filter out the
  HTML tags from the <code>temp.txt</code> file. For that, we will introduce another
  tool to filter text called <code>sed</code> or Stream editor. This tool allows us
  to filter the stream field and print or store the outcome. The following code will
  remove the HTML tags from the scrapped text.</p>\n<h3>Using SED command to remove
  embedded</h3>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>
  sed -i <span class=\"s1\">&#39;s/&lt;[^&gt;]*&gt;//g&#39;</span> temp.txt &gt;meaning.txt\n</pre></div>\n\n</pre>\n\n<p>The
  above command filters the text and removes the HTML tags from the <code>temp.txt
  </code>file using regular expressions. The <code>-i</code> command allows us to
  store the output in a file <code>meaning.txt</code>.  We have used Regex to remove
  <code>&lt;&gt;</code> tags from the file and hence anything in between these is
  also removed and we get the only pure text but it may also contain special characters
  and symbols. To remove that we'll again use <code>grep</code> and filter the fine
  meaning in our file.</p>\n<h3>Removing Special Characters from the Content using
  GREP commands.</h3>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>
  grep -v <span class=\"s1\">&#39;^\\s*$\\|^\\s*\\#&#39;</span> temp.txt &gt;meaning.txt\n</pre></div>\n\n</pre>\n\n<p>Now
  from the above command removes the special characters such as <code>$,#</code>,
  and others from the temp.txt file. We finally store everything filtered in the meaning.txt
  file. If you understood till here, the next concrete step will be super easy for
  you, as we will assemble everything here in a shell script.</p>\n<h2>Making the
  Shell Script</h2>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  the word to find meaning : &quot;</span> word\n<span class=\"nv\">output</span><span
  class=\"o\">=</span><span class=\"s2\">&quot;meaning.txt&quot;</span>\n<span class=\"nv\">url</span><span
  class=\"o\">=</span><span class=\"s2\">&quot;https://www.dictionary.com/browse/</span><span
  class=\"nv\">$word</span><span class=\"s2\">#&quot;</span>\n\ncurl -o <span class=\"nv\">$output</span>
  <span class=\"nv\">$url</span> \nclear\ngrep -oP <span class=\"s1\">&#39;(?&lt;=&lt;span
  class=&quot;one-click-content css-nnyc96 e1q3nk1v1&quot;&gt;).*?(?=&lt;/span&gt;)&#39;</span>
  <span class=\"nv\">$output</span> &gt;temp.txt \n\nsed -i <span class=\"s1\">&#39;s/&lt;[^&gt;]*&gt;//g&#39;</span>
  temp.txt &gt;<span class=\"nv\">$output</span>\ngrep -v <span class=\"s1\">&#39;^\\s*$\\|^\\s*\\#&#39;</span>
  temp.txt &gt;<span class=\"nv\">$output</span>\n<span class=\"nb\">echo</span> <span
  class=\"s2\">&quot;</span><span class=\"nv\">$word</span><span class=\"s2\">&quot;</span>\n<span
  class=\"k\">while</span> <span class=\"nb\">read</span> meaning \n<span class=\"k\">do</span>\n\t<span
  class=\"nb\">echo</span> <span class=\"nv\">$meaning</span>\n<span class=\"k\">done</span>
  &lt; <span class=\"nv\">$output</span>\n</pre></div>\n\n</pre>\n\n<p>We can clearly
  see most of the commands are the same, but some have been modified to avoid repetition
  and automation. Firstly, I have taken user input of word from the user and stored
  it in with an appropriate variable name.  Next, I have created another variable
  to store the file name in which we are going to store the meaning of the word, Also
  a variable for the URL of the website we are searching for. We have used a variable
  to access the required URL. Then we invoke <code>cURL</code> to the file which we
  want to store using the variable we created and the URL variable So creating variables
  makes our script quite easy to manage and also it improves the readability of the
  script.</p>\n<h2>Updating cURL command</h2>\n<p>We can also update the curl command
  by adding <code>&quot;&amp;&gt; /dev/null&quot;</code> this will dump the curl output
  of network analysis. So we will only get the output of the meaning.txt file.  It
  is optional to add the following into your code as it depends on the operating system
  so we can optionally use clear command to wipe out the curl output.</p>\n<h2>Printing
  the output file line by line.</h2>\n<p>To print the meaning in the output file,
  we need to print each line separately as the meanings are distinct. Therefore, we
  will use a while loop with the output file and echo the line variable we have used
  as the loop iterator.</p>\n<h2>Script Screenshots:</h2>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1627366344193/We_heehuL.gif\"
  alt=\"dict.gif\" /></p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1627365131696/YH8Vaqoh_.png\"
  alt=\"image.png\" /></p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1627365274090/D9IETfRAh.png\"
  alt=\"image.png\" /></p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1627365304653/A9AXuHDH8.png\"
  alt=\"image.png\" /></p>\n<h2>Output Conclusion</h2>\n<p>From the above output,
  we have scrapped the meaning of the word <code>Mathematics</code>, <code>code</code>,
  and <code>python</code>.  It works only for the words which are on the <a href=\"http://dictionary.com\">dictionary.com</a>
  website. We have successfully made a scrapper that scraps the meaning of the input
  word from the <a href=\"http://dictionary.com\">dictionary.com</a> website,</p>\n<h2>Appropriate
  use of Web-Scrapping.</h2>\n<p>We must be careful and not scrape any website without
  reading its privacy policy. If they allow scraping from their website, then only
  you should scrape the content and not use it for any monetization of the content.
  This was just used for demonstrating some idea about web scrapping using BASH and
  just meant for teaching purposes.</p>\n<p>Therefore, it is quite easy to scrape
  the website's content especially if you find any patterns in the code structure.
  We were able to make a script that can print the meaning of the input word from
  the base of the website <a href=\"http://dictionary.com\">dictionary.com</a>.</p>\n<p>We
  can see how Bash can be powerful in terms of web scrapping. I hope you found this
  interesting and inspiring. Thank you for reading. Happy Coding :)</p>\n"
cover: ''
date: 2021-07-27
datetime: 2021-07-27 00:00:00+00:00
description: Web Scraping is quite an interesting and powerful tool or skill to have
  in a Programmer For this article, we aim to fetch the meaning of a word entered
  by the u
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-07-27-BASH-script-dictionary-scrap.md
html: "<h2>Introduction</h2>\n<p>Web Scraping is quite an interesting and powerful
  tool or skill to have in a Programmer's toolkit.  It helps in analyzing data and
  getting some information in various formats. Web Scraping is a process in which
  a user fetches a website's content using some pattern in those HTML tags and the
  desired content to be fetched or scraped.</p>\n<p>For this article, we aim to fetch
  the meaning of a word entered by the user from the <a href=\"http://dictionary.com\">dictionary.com</a>
  website. We need to print just the meaning of the word from the HTML tags in it.
  We must have a good understanding of HTML and some basic Linux tools such as cURL,
  grep, sed, and others for doing all of these.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625737499658/FGLusWSII.png\"
  alt=\"Inspecting the Target Website\" /></p>\n<h2>Inspecting the Target Website</h2>\n<p>To
  begin with, scrapping the website, first, it is absolutely important to inspect
  the website and view its source code. For that, we can make use of Inspect tool
  in our Browsers. Just Right-click on the website you are viewing or the website
  for scraping, a list of options appears in front of you. You have to select Inspect
  option( also Shift + Ctrl + I), this will open a side window with a plethora of
  options. You simply have to select Elements from the top of the menus. The code
  that you will see is the source code of the website. No, don't think you can change
  the content of the website from here :)</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625737510444/KonUrEpcq-.png\"
  alt=\"image.png\" />\nInspect Tool in the browser.</p>\n<p>Now we have to analyze
  the website with the content which we want to scrape. You can go on for clicking
  the <code>select the element in the page to inspect it</code> option or icon in
  the top left-hand side corner. This will allow you to inspect the particular element
  that you selected on the webpage. You can now see the element tag, id, class, and
  other attributes required to fetch the element's content.</p>\n<h2>Selecting the
  particular element from the website to view the source code.</h2>\n<h3>Accessing
  the website from the Command line/terminal</h3>\n<p>Now the website structure is
  being understood we can actually move to scrap it. For that, we need to have the
  web site's content on our local machine. First of all, we need to access the website
  from elsewhere not from the browser, because you cannot copy-paste content from
  there. So let's use Command Line here. We have a popular tool known as <code>cURL</code>,
  which stands for client URL. The tool fetches the contents of the provided URL.
  It also has several parameters or arguments that can be used to modify its output.
  We can use the command</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n\n<div class=\"highlight\"><pre><span></span>curl -o meaning.txt https://www.dictionary.com/browse/computer#\n</pre></div>\n\n</pre>\n\n<p>The
  above command fetches the HTML page for the word Computer, it could be any word
  you might be searching for.</p>\n<h3>Understanding the Website Structure.</h3>\n<p>Here
  comes the time to explain the structure of <a href=\"http://dictionary.com\">dictionary.com</a>.
  When you search a word on the website(<a href=\"http://dictionary.com\">dictionary.com</a>),
  you are routed to <code>/browse</code> which then fetches the word for you and defaults
  you to the <code>/browse/word#</code> (the word can be any word you searched). The
  curl command dumps the output in the <code>meaning.txt</code> or any specified file.
  If you see the contents of the file, it is the same as on the web.  So we are going
  to store the meaning of the searched word in the meaning.txt file, you can customize
  the name and command however you like.</p>\n<p>Voila! you successfully scraped a
  webpage. Now the next target is to filter the webpage content.</p>\n<h3>Filtering
  Content from Website local file</h3>\n<p>Now we have the content of the webpage
  on our local machine, we need to search or filter out the useful content and remove
  the unwanted tags and elements. For that, we can use commands such as <code>grep</code>
  and <code>sed</code>.</p>\n<h3>Finding Tags to Extract content.</h3>\n<p>We need
  to find patterns and similarities in the tags that contain the text of the meaning
  of the specified word. From the analysis of the webpage, we see that the element
  <code>&lt;span class=&quot;one-click-content css-nnyc96 e1q3nk1v1&quot;&gt;</code>
  contains the actual meaning. We just need the basic meaning, we may not need examples
  and long lengthy definitions on our Terminal, So we will go with filtering out the
  span tag with a class called <code>one-click-content css-nnyc96 e1q3nk1v1</code>.
  To do that we can use the grep command, which can print the text or line matching
  the specified expression or text. Here we need the span element with the particular
  class name so we will use regular expressions to find it more effectively.</p>\n<pre
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>grep
  -oP <span class=\"s1\">&#39;(?&lt;=&lt;span class=&quot;one-click-content css-nnyc96
  e1q3nk1v1&quot;&gt;).*?(?=&lt;/span&gt;)&#39;</span> meaning.txt &gt;temp.txt \n</pre></div>\n\n</pre>\n\n<h3>Using
  GREP command to filter.</h3>\n<p>The above command will search and return only lines
  that are contained in the span tags with that particular class name from the meaning.txt
  file which we appended to fill the webpage's source code. The <code>-oP</code> are
  the arguments that return Only the matching cases and <code>-P</code> the coming
  expression is a Perl Regex. The command will return everything in between those
  tags. Finally, we are storing the result or output in <code>temp.txt</code>.</p>\n<p>Now,
  if you think we are done, then it's wrong, the webpage can have internal or external
  links embedded inside of the elements as well, so we need to again filter out the
  HTML tags from the <code>temp.txt</code> file. For that, we will introduce another
  tool to filter text called <code>sed</code> or Stream editor. This tool allows us
  to filter the stream field and print or store the outcome. The following code will
  remove the HTML tags from the scrapped text.</p>\n<h3>Using SED command to remove
  embedded</h3>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>
  sed -i <span class=\"s1\">&#39;s/&lt;[^&gt;]*&gt;//g&#39;</span> temp.txt &gt;meaning.txt\n</pre></div>\n\n</pre>\n\n<p>The
  above command filters the text and removes the HTML tags from the <code>temp.txt
  </code>file using regular expressions. The <code>-i</code> command allows us to
  store the output in a file <code>meaning.txt</code>.  We have used Regex to remove
  <code>&lt;&gt;</code> tags from the file and hence anything in between these is
  also removed and we get the only pure text but it may also contain special characters
  and symbols. To remove that we'll again use <code>grep</code> and filter the fine
  meaning in our file.</p>\n<h3>Removing Special Characters from the Content using
  GREP commands.</h3>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>shell</div>\n<div class=\"highlight\"><pre><span></span>
  grep -v <span class=\"s1\">&#39;^\\s*$\\|^\\s*\\#&#39;</span> temp.txt &gt;meaning.txt\n</pre></div>\n\n</pre>\n\n<p>Now
  from the above command removes the special characters such as <code>$,#</code>,
  and others from the temp.txt file. We finally store everything filtered in the meaning.txt
  file. If you understood till here, the next concrete step will be super easy for
  you, as we will assemble everything here in a shell script.</p>\n<h2>Making the
  Shell Script</h2>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  the word to find meaning : &quot;</span> word\n<span class=\"nv\">output</span><span
  class=\"o\">=</span><span class=\"s2\">&quot;meaning.txt&quot;</span>\n<span class=\"nv\">url</span><span
  class=\"o\">=</span><span class=\"s2\">&quot;https://www.dictionary.com/browse/</span><span
  class=\"nv\">$word</span><span class=\"s2\">#&quot;</span>\n\ncurl -o <span class=\"nv\">$output</span>
  <span class=\"nv\">$url</span> \nclear\ngrep -oP <span class=\"s1\">&#39;(?&lt;=&lt;span
  class=&quot;one-click-content css-nnyc96 e1q3nk1v1&quot;&gt;).*?(?=&lt;/span&gt;)&#39;</span>
  <span class=\"nv\">$output</span> &gt;temp.txt \n\nsed -i <span class=\"s1\">&#39;s/&lt;[^&gt;]*&gt;//g&#39;</span>
  temp.txt &gt;<span class=\"nv\">$output</span>\ngrep -v <span class=\"s1\">&#39;^\\s*$\\|^\\s*\\#&#39;</span>
  temp.txt &gt;<span class=\"nv\">$output</span>\n<span class=\"nb\">echo</span> <span
  class=\"s2\">&quot;</span><span class=\"nv\">$word</span><span class=\"s2\">&quot;</span>\n<span
  class=\"k\">while</span> <span class=\"nb\">read</span> meaning \n<span class=\"k\">do</span>\n\t<span
  class=\"nb\">echo</span> <span class=\"nv\">$meaning</span>\n<span class=\"k\">done</span>
  &lt; <span class=\"nv\">$output</span>\n</pre></div>\n\n</pre>\n\n<p>We can clearly
  see most of the commands are the same, but some have been modified to avoid repetition
  and automation. Firstly, I have taken user input of word from the user and stored
  it in with an appropriate variable name.  Next, I have created another variable
  to store the file name in which we are going to store the meaning of the word, Also
  a variable for the URL of the website we are searching for. We have used a variable
  to access the required URL. Then we invoke <code>cURL</code> to the file which we
  want to store using the variable we created and the URL variable So creating variables
  makes our script quite easy to manage and also it improves the readability of the
  script.</p>\n<h2>Updating cURL command</h2>\n<p>We can also update the curl command
  by adding <code>&quot;&amp;&gt; /dev/null&quot;</code> this will dump the curl output
  of network analysis. So we will only get the output of the meaning.txt file.  It
  is optional to add the following into your code as it depends on the operating system
  so we can optionally use clear command to wipe out the curl output.</p>\n<h2>Printing
  the output file line by line.</h2>\n<p>To print the meaning in the output file,
  we need to print each line separately as the meanings are distinct. Therefore, we
  will use a while loop with the output file and echo the line variable we have used
  as the loop iterator.</p>\n<h2>Script Screenshots:</h2>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1627366344193/We_heehuL.gif\"
  alt=\"dict.gif\" /></p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1627365131696/YH8Vaqoh_.png\"
  alt=\"image.png\" /></p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1627365274090/D9IETfRAh.png\"
  alt=\"image.png\" /></p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1627365304653/A9AXuHDH8.png\"
  alt=\"image.png\" /></p>\n<h2>Output Conclusion</h2>\n<p>From the above output,
  we have scrapped the meaning of the word <code>Mathematics</code>, <code>code</code>,
  and <code>python</code>.  It works only for the words which are on the <a href=\"http://dictionary.com\">dictionary.com</a>
  website. We have successfully made a scrapper that scraps the meaning of the input
  word from the <a href=\"http://dictionary.com\">dictionary.com</a> website,</p>\n<h2>Appropriate
  use of Web-Scrapping.</h2>\n<p>We must be careful and not scrape any website without
  reading its privacy policy. If they allow scraping from their website, then only
  you should scrape the content and not use it for any monetization of the content.
  This was just used for demonstrating some idea about web scrapping using BASH and
  just meant for teaching purposes.</p>\n<p>Therefore, it is quite easy to scrape
  the website's content especially if you find any patterns in the code structure.
  We were able to make a script that can print the meaning of the input word from
  the base of the website <a href=\"http://dictionary.com\">dictionary.com</a>.</p>\n<p>We
  can see how Bash can be powerful in terms of web scrapping. I hope you found this
  interesting and inspiring. Thank you for reading. Happy Coding :)</p>\n"
image_url: https://cdn.hashnode.com/res/hashnode/image/upload/v1627367329063/dabJLKcD-.png?w=1600&h=840&fit=crop&crop=entropy&auto=compress
long_description: Web Scraping is quite an interesting and powerful tool or skill
  to have in a Programmer For this article, we aim to fetch the meaning of a word
  entered by the user from the dictionary.com website. We need to print just the meaning
  of the word from th
now: 2025-05-01 18:11:33.313622
path: blog/posts/2021-07-27-BASH-script-dictionary-scrap.md
prevnext: null
slug: bash-dictionary-scrapper
status: published
subtitle: Using some shell tools and utilities to understand web scrapping and making
  a script to scrap a webpage.
tags:
- bash
templateKey: blog-post
title: Scrapping the meaning of a word from dictionary.com using BASH script.
today: 2025-05-01
---

## Introduction

Web Scraping is quite an interesting and powerful tool or skill to have in a Programmer's toolkit.  It helps in analyzing data and getting some information in various formats. Web Scraping is a process in which a user fetches a website's content using some pattern in those HTML tags and the desired content to be fetched or scraped.

For this article, we aim to fetch the meaning of a word entered by the user from the dictionary.com website. We need to print just the meaning of the word from the HTML tags in it. We must have a good understanding of HTML and some basic Linux tools such as cURL, grep, sed, and others for doing all of these. 

![Inspecting the Target Website](https://cdn.hashnode.com/res/hashnode/image/upload/v1625737499658/FGLusWSII.png)

## Inspecting the Target Website

To begin with, scrapping the website, first, it is absolutely important to inspect the website and view its source code. For that, we can make use of Inspect tool in our Browsers. Just Right-click on the website you are viewing or the website for scraping, a list of options appears in front of you. You have to select Inspect option( also Shift + Ctrl + I), this will open a side window with a plethora of options. You simply have to select Elements from the top of the menus. The code that you will see is the source code of the website. No, don't think you can change the content of the website from here :)

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625737510444/KonUrEpcq-.png)
Inspect Tool in the browser.

Now we have to analyze the website with the content which we want to scrape. You can go on for clicking the `select the element in the page to inspect it` option or icon in the top left-hand side corner. This will allow you to inspect the particular element that you selected on the webpage. You can now see the element tag, id, class, and other attributes required to fetch the element's content.


## Selecting the particular element from the website to view the source code.


### Accessing the website from the Command line/terminal

Now the website structure is being understood we can actually move to scrap it. For that, we need to have the web site's content on our local machine. First of all, we need to access the website from elsewhere not from the browser, because you cannot copy-paste content from there. So let's use Command Line here. We have a popular tool known as `cURL`, which stands for client URL. The tool fetches the contents of the provided URL. It also has several parameters or arguments that can be used to modify its output. We can use the command

```
curl -o meaning.txt https://www.dictionary.com/browse/computer#
```

The above command fetches the HTML page for the word Computer, it could be any word you might be searching for.

### Understanding the Website Structure.

Here comes the time to explain the structure of dictionary.com. When you search a word on the website(dictionary.com), you are routed to `/browse` which then fetches the word for you and defaults you to the `/browse/word#` (the word can be any word you searched). The curl command dumps the output in the `meaning.txt` or any specified file. If you see the contents of the file, it is the same as on the web.  So we are going to store the meaning of the searched word in the meaning.txt file, you can customize the name and command however you like.

Voila! you successfully scraped a webpage. Now the next target is to filter the webpage content.

### Filtering Content from Website local file

Now we have the content of the webpage on our local machine, we need to search or filter out the useful content and remove the unwanted tags and elements. For that, we can use commands such as `grep` and `sed`. 

### Finding Tags to Extract content.

We need to find patterns and similarities in the tags that contain the text of the meaning of the specified word. From the analysis of the webpage, we see that the element `<span class="one-click-content css-nnyc96 e1q3nk1v1">` contains the actual meaning. We just need the basic meaning, we may not need examples and long lengthy definitions on our Terminal, So we will go with filtering out the span tag with a class called `one-click-content css-nnyc96 e1q3nk1v1`. To do that we can use the grep command, which can print the text or line matching the specified expression or text. Here we need the span element with the particular class name so we will use regular expressions to find it more effectively.

```shell
grep -oP '(?<=<span class="one-click-content css-nnyc96 e1q3nk1v1">).*?(?=</span>)' meaning.txt >temp.txt 
```

### Using GREP command to filter.

The above command will search and return only lines that are contained in the span tags with that particular class name from the meaning.txt file which we appended to fill the webpage's source code. The `-oP` are the arguments that return Only the matching cases and `-P` the coming expression is a Perl Regex. The command will return everything in between those tags. Finally, we are storing the result or output in `temp.txt`. 

Now, if you think we are done, then it's wrong, the webpage can have internal or external links embedded inside of the elements as well, so we need to again filter out the HTML tags from the `temp.txt` file. For that, we will introduce another tool to filter text called `sed` or Stream editor. This tool allows us to filter the stream field and print or store the outcome. The following code will remove the HTML tags from the scrapped text.

### Using SED command to remove embedded 

```shell
 sed -i 's/<[^>]*>//g' temp.txt >meaning.txt
```
The above command filters the text and removes the HTML tags from the `temp.txt `file using regular expressions. The `-i` command allows us to store the output in a file `meaning.txt`.  We have used Regex to remove `<>` tags from the file and hence anything in between these is also removed and we get the only pure text but it may also contain special characters and symbols. To remove that we'll again use `grep` and filter the fine meaning in our file.

### Removing Special Characters from the Content using GREP commands.
```shell
 grep -v '^\s*$\|^\s*\#' temp.txt >meaning.txt
```

Now from the above command removes the special characters such as `$,#`, and others from the temp.txt file. We finally store everything filtered in the meaning.txt file. If you understood till here, the next concrete step will be super easy for you, as we will assemble everything here in a shell script.

## Making the Shell Script

```bash
#!/bin/bash

read -p "Enter the word to find meaning : " word
output="meaning.txt"
url="https://www.dictionary.com/browse/$word#"

curl -o $output $url 
clear
grep -oP '(?<=<span class="one-click-content css-nnyc96 e1q3nk1v1">).*?(?=</span>)' $output >temp.txt 

sed -i 's/<[^>]*>//g' temp.txt >$output
grep -v '^\s*$\|^\s*\#' temp.txt >$output
echo "$word"
while read meaning 
do
	echo $meaning
done < $output
```

We can clearly see most of the commands are the same, but some have been modified to avoid repetition and automation. Firstly, I have taken user input of word from the user and stored it in with an appropriate variable name.  Next, I have created another variable to store the file name in which we are going to store the meaning of the word, Also a variable for the URL of the website we are searching for. We have used a variable to access the required URL. Then we invoke `cURL` to the file which we want to store using the variable we created and the URL variable So creating variables makes our script quite easy to manage and also it improves the readability of the script. 

## Updating cURL command

We can also update the curl command by adding `"&> /dev/null"` this will dump the curl output of network analysis. So we will only get the output of the meaning.txt file.  It is optional to add the following into your code as it depends on the operating system so we can optionally use clear command to wipe out the curl output.

## Printing the output file line by line.

To print the meaning in the output file, we need to print each line separately as the meanings are distinct. Therefore, we will use a while loop with the output file and echo the line variable we have used as the loop iterator.

## Script Screenshots:


![dict.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1627366344193/We_heehuL.gif)


![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1627365131696/YH8Vaqoh_.png)


![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1627365274090/D9IETfRAh.png)


![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1627365304653/A9AXuHDH8.png)
 

## Output Conclusion

From the above output, we have scrapped the meaning of the word `Mathematics`, `code`, and `python`.  It works only for the words which are on the dictionary.com website. We have successfully made a scrapper that scraps the meaning of the input word from the dictionary.com website, 

## Appropriate use of Web-Scrapping.

We must be careful and not scrape any website without reading its privacy policy. If they allow scraping from their website, then only you should scrape the content and not use it for any monetization of the content. This was just used for demonstrating some idea about web scrapping using BASH and just meant for teaching purposes.

Therefore, it is quite easy to scrape the website's content especially if you find any patterns in the code structure. We were able to make a script that can print the meaning of the input word from the base of the website dictionary.com.

We can see how Bash can be powerful in terms of web scrapping. I hope you found this interesting and inspiring. Thank you for reading. Happy Coding :)