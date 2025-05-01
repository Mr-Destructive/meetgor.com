---
article_html: "<h2>Introduction</h2>\n<p>If you have been writing articles you know
  the pain to get some attention, if you have already been cross-posting your articles
  it usually takes some time to do that. This task can be automated with a shellscript.
  If you have been cross-posting articles on <code>medium.com</code>, <code>dev.to</code>
  and at <code>hashnode.com</code>, then I have a treat for you.</p>\n<p>Introducing
  <strong><a href=\"http://crossposter.sh\">crossposter.sh</a></strong>!!</p>\n<h2>What
  is <a href=\"http://Crossposter.sh\">Crossposter.sh</a>?</h2>\n<h3>Crosspost to
  <a href=\"http://dev.to/hahsnode/medium\">dev.to/hahsnode/medium</a> from the command
  line.</h3>\n<p><a href=\"http://Crossposter.sh\">Crossposter.sh</a> is a shellscript(BASH)
  to automate crossposting to platforms like <a href=\"http://dev.to\">dev.to</a>,
  <a href=\"http://medium.com\">medium.com</a> and <a href=\"http://hashnode.com\">hashnode.com</a>.
  The script takes in markdown version of your post with a few inputs from you and
  posts it to those platforms. You would require a token/key for each of those platforms
  to post it from the command line. You can check out the official repository of <a
  href=\"https://github.com/Mr-Destructive/crossposter\">Crossposter</a>.</p>\n<p>The
  actual script is still not perfect (has a few bugs). Though it posts on <code>dev.to</code>
  and <code>medium.com</code> easily, the <code>hashnode.com</code> is buggy as it
  parses the raw markdown into the post and doesn't render as desired. So, <strong>its
  a under-development script</strong>, fell free to raise any issues or PRs on the
  official GitHub repo.</p>\n<p>Run the script on a bash interpreter with the command:</p>\n<p><code>bash
  crosspost.sh</code></p>\n<p>For posting the article you need to provide the following
  details:</p>\n<h2>Front-Matter</h2>\n<h3>Meta data about the post</h3>\n<ul>\n<li>Title
  of Post</li>\n<li>Subtitle of Post</li>\n<li>Publish status of post(<code>true</code>
  or <code>false</code>)</li>\n<li>Tags for the post (comma separated values)</li>\n<li>Canonical
  Url (original url of the post)</li>\n<li>Cover Image (URL of the post's image/thumbnail)</li>\n</ul>\n<p>This
  information is a must for <code>dev.to</code> especially the <code>title</code>.
  This should be provide in the same order as given below:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>yaml</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nn\">---</span><span class=\"w\"></span>\n<span class=\"nt\">title</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"l l-Scalar l-Scalar-Plain\">The
  title of the post</span><span class=\"w\"></span>\n<span class=\"nt\">subtitle</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"l l-Scalar l-Scalar-Plain\">The
  description of your article</span><span class=\"w\"></span>\n<span class=\"nt\">published</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"l l-Scalar l-Scalar-Plain\">true</span><span
  class=\"w\"></span>\n<span class=\"nt\">tags</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"l l-Scalar l-Scalar-Plain\">programming, anythingelse</span><span
  class=\"w\"></span>\n<span class=\"nt\">canonical url</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"l l-Scalar l-Scalar-Plain\">url of your original
  blog</span><span class=\"w\"></span>\n<span class=\"nt\">cover_image</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"l l-Scalar l-Scalar-Plain\">coverimage_url</span><span
  class=\"w\"></span>\n<span class=\"nn\">---</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>There
  is no need to enclose any of them with quotation marks. <code>Published</code> argument
  will be <code>true</code> if you want to publish it and <code>false</code> if you
  want to keep it in your Drafts.</p>\n<p>In the demonstrations, we just need to enter
  the tokens once. The tokens will be stored locally in the <code>keys.txt</code>
  file and retrieved later within the script.</p>\n<h2>Posting on <strong><a href=\"http://dev.to\">dev.to</a></strong>:</h2>\n<p>Posting
  on <a href=\"http://dev.to\">dev.to</a> requires their <code>API key</code> which
  can be generated by going on the <a href=\"https://dev.to/settings/account/\">Dev
  Community API Keys</a>. From there you can generate a new key with any name you
  like. You just need to enter the key to CLI once or manually enter in the <code>keys.txt</code>
  file with the format <code>dev.to:key</code> on the first line. This will be used
  for the future cross-posting whenever you execute the shell script(<code>bash crosspost.sh</code>)</p>\n<p>You
  can provide the <a href=\"#front-matter\">front matter</a> manually in your markdown
  file or you will be prompted for the input. So, that is all you will require for
  posting on <a href=\"http://dev.to\">dev.to</a> from the Command line.</p>\n<p>Lets
  see the script in action</p>\n<p><img src=\"https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/devto.gif\"
  alt=\"dev.to\" /></p>\n<p>If you want to add in more stuff to the post, you can
  check out the <a href=\"https://developers.forem.com/api#operation/createArticle\">DEV.to
  API docs</a> which is powered by <a href=\"https://www.forem.com/\">Forem</a>, there
  a ton of options you can hook to the front-matter in the shellscript.</p>\n<p><strong>NOTE:
  There is a limit of 10 requests per 30 seconds, so keep in mind while testing the
  script and don't try to spam</strong></p>\n<h2>Posting on <strong><a href=\"http://hashnode.com\">hashnode.com</a></strong>:</h2>\n<p>This
  part is still under development as it only displays the raw markdown in the post,
  also the <code>tags</code> are too heavy to implement from the API as <code>id</code>
  of every tag is required along with the <code>slug</code> and <code>name</code>.
  Still it serves some purpose at least. For posting on <code>hashnode.com</code>,
  we need <code>Personal Access Token</code>. This can be generated by going to <a
  href=\"https://hashnode.com/settings/developer\">Developer Settings</a>. You will
  also require the user-id of your <code>hashnode</code> account. You can get your
  user-id/username from the <a href=\"https://hashnode.com/settings\">settings</a>
  tab under profile information. We require Username for posting to the Publication
  Blog if any. As usual, the <code>Personal Access Token</code> for interacting with
  the <a href=\"https://api.hashnode.com/\">Hashnodes' GraphQL API</a>. The API is
  quite user friendly and provides everything in one place. There are docs for running
  each and every <code>query</code> and <code>mutations</code> present in the API.</p>\n<p>You
  can paste the token when prompted from the script or manually type in the <code>keys.txt</code>
  text file as <code>hashnode:token</code> on the 4th line. Yes, that should be on
  the <code>4th</code> line, thats make retrieving much more easier and safe. Next
  also input in the <code>username</code> when the script asks for the input or again
  type in on the <code>5th</code> line, <code>hashnode_id:username</code> in the text
  file <code>keys.txt</code>. Please enter the credentials from the script prompt
  so as to avoid errors and misconfigurations when doing manually</p>\n<p>This will
  create the Post on hashnode with the title, subtitle, cover image correctly but
  will mess up the content. I tried hard but its just not happening. There needs to
  be some character for newline as the API rejects the <code>rn</code> characters
  passed in, so I have substited them with <code>br</code> and the result is raw markdown.
  <strong>As the Hashnode API is still under development and they are bringing changes
  and new features in, the API should improve in its core functionality and make it
  much easier for creating some common queries</strong>. So, I'll create issue on
  GitHub for posting the actual content via the script.</p>\n<p>So, this is the demonstration
  of posting on hashnode.</p>\n<p><img src=\"https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/hashnode.gif\"
  alt=\"hashnode\" /></p>\n<h2>Posting on <strong><a href=\"http://medium.com\">medium.com</a></strong>:</h2>\n<p>Medium
  API is much more versatile and markdown friendly, though it has some limitation
  on the number of posts you can make in a day. For posting on <code>Medium.com</code>,
  we will require the <code>Integration Token</code> which can be generated on the
  <a href=\"https://medium.com/me/settings\">settings tab</a>. As similar to <code>hashnode</code>,
  you can name the token whatever you like and then get the token. Paste the token
  when prompted from the script or manually type in the <code>keys.txt</code> text
  file as <code>medium:token</code> on the <code>2nd</code> line. We also require
  the Medium_id, but we can get that from the token itself, so inside the script once
  the token is obtained, the curl command is executed to fetch in the <code>id</code>
  and it is stored on the next(<code>3rd</code>) line in the <code>keys.txt</code>
  file for actually posting on <code>medium.com</code>. So that is all the configuration
  you need for posting on <code>medium.com</code>.</p>\n<p>There is some documentation
  on <a href=\"https://github.com/Medium/medium-api-docs\">Medium API</a>, we can
  even post to a Publication, that shall be created in future. Also the cover images
  can be posted on medium, it is not currently done but that can again be a #TODO.
  <strong>The tags are not rendered on Medium yet with the script.</strong> The way
  we can parse  strings is limited in BASH, so this might still be a doable thing
  later. Most of the checkboxes are ticked like title, subtitle, cover-image, canonical
  url, and importantly the content.</p>\n<p>Let's look at post on medium from the
  script.</p>\n<p><img src=\"https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/medium.gif\"
  alt=\"medium\" /></p>\n<h2>All platforms:</h2>\n<p>Now, once you have configured
  every thing, you can opt for the <code>4</code> choice that is post on all platforms(<a
  href=\"http://dev.to\">dev.to</a>, hashnode and medium), but as hashnode is not
  looking a good option right now, so there is the <code>5</code> option for only
  <code>dev.to</code> and <code>medium</code>.</p>\n<p><img src=\"https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/crossposter.gif\"
  alt=\"allplatforms\" /></p>\n<h2>Why use <a href=\"http://Crossposter.sh\">Crossposter.sh</a>?</h2>\n<p>This
  can be not so big of an issue for most of the people but it was a good side project
  to work and learn more about how APIs work and get some ideas on the design of the
  platform. Though it is quite time saving to cross post on 3 different platforms
  within a minute or two. You can tailor your own script as per your specifications
  and desire.</p>\n<p>So, if you are an author on all of the mentioned platforms,
  please give it a try. Other Platforms are welcome for contributions. If you found
  any unexpected things, please hit them in the <code>issues</code> tab.</p>\n<h2>Script</h2>\n<p>The
  script mostly leverages <code>curl</code>, <code>sed</code>, <code>cat</code> and
  some other basic utilities in BASH.</p>\n<h3>Using <code>curl</code> for posting
  the article from APIs</h3>\n<p>Curl is a life saver command for this project, without
  this tool, the project might not be as great and efficient. Let's see some quick
  commands used in the script.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>curl
  -H <span class=\"s2\">&quot;Content-Type: application/json&quot;</span> -H <span
  class=\"s2\">&quot;api-key&quot;</span>: <span class=\"se\">\\&quot;</span><span
  class=\"s1\">&#39;&quot;$key&quot;&#39;</span><span class=\"se\">\\&quot;</span>
  -d <span class=\"s1\">&#39;{&quot;content&quot;:\\&quot;&#39;</span><span class=\"s2\">&quot;</span><span
  class=\"nv\">$body</span><span class=\"s2\">&quot;</span><span class=\"s1\">&#39;\\&quot;}&#39;</span>
  <span class=\"s2\">&quot;</span><span class=\"nv\">$url</span><span class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  the above command is quite basic, some more additions are also added as per the
  specifications of the Platform. But, let us understand the structure of the command
  we are sending to the APIs. The first part is the Header (<code>-H</code>), in here
  we specify the content that is going to get parsed and the api-keys to access the
  API. Next, we have the body or the data (<code>-d</code>), here we parse in the
  actual contents, it might be the front matter along with the markdown content. Finally
  we have the <code>url</code> where we send the <code>POST</code> request i.e. the
  <code>API endpoint</code>. The is the escape character that is used to preserve
  the literal value of the next character and in short we can shorten the command
  to fit in the next line.</p>\n<p>The wired <code>$body</code> is used to parse the
  value of the variable <code>body</code> inside of single quotes as in BASH, we can
  only access the variables' value in double quotes. We are using single quotes as
  we have to pass the <code>json</code> object and which has already double quotes
  in it.</p>\n<h3>Using <code>sed</code> for editing text</h3>\n<p>Sed is a super-powerful
  stream editor, its somewhat similar to Vim without an interface, only commands.
  We use this tool to manipulate the front-matter for posting on the platforms by
  parsing them to variables in BASH. We also use to enter the api keys inputted by
  user from variables into the file at a specific position to retrieve later.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sed
  -i <span class=\"s2\">&quot;1a title: </span><span class=\"nv\">$title</span><span
  class=\"s2\">&quot;</span> file.txt\n</pre></div>\n\n</pre>\n\n<p>Here, we are appending(<code>a</code>)
  to the 1st line, text <code>title: $title</code>, here <code>$title</code> is the
  variable, so we are technically parsing the value of the variable <code>title</code>.
  We are editing the file <code>file.txt</code> in-place <code>-i</code> i.e. we are
  editing it live without creating any temp or backup files.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sed
  -n -e <span class=\"s2\">&quot;s/dev.to://p&#39; keys.txt&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>Here
  we are essentially getting the text after a particular pattern. In this case we
  are searching in <code>keys.txt</code> file for the string <code>dev.to:</code>
  and anything after that till the end of line is returned, we can further store it
  in the variable and do all sorts of operation.</p>\n<h3>Using <code>awk</code> for
  programmatic editing</h3>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>awk
  <span class=\"s1\">&#39;{print $0&quot;\\r\\n&quot;}&#39;</span> temp.txt &gt;file.txt\n</pre></div>\n\n</pre>\n\n<p>AWK
  is a command-line utility for manipulating or writing certain operations/patterns
  programmatically. We use this tool so as to add <code>4r4n</code> to the end of
  each line, the APIs can't parse the file contents directly so we have to add certain
  characters before the end of line and do further operations.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>cat
  temp.md <span class=\"p\">|</span> tr -d <span class=\"s1\">&#39;\\r\\n&#39;</span>
  &gt; temp.txt\n</pre></div>\n\n</pre>\n\n<p>After we have added the <code>\\r\\n</code>
  characters to the end of the file, we simply can use <code>cat</code> and <code>tr</code>
  to merge all the lines into a single line. This is how we parse the contents to
  the API more safely and concisely, of course we need to parse them in a variable
  by reading the file.</p>\n<p>OK, I won't bore anyone with more BASH but that were
  some of the quite important commands in the script that form the backbone of the
  cross-posting and handling text with the APIs.</p>\n<h2>Conclusion</h2>\n<p>So,
  we can see <code>crosspost.sh</code> is a BASH script that cross-posts markdown
  articles with a bit of inputs to 3 different platforms within a couple of minutes.
  This article was basically to demonstrate the project and its capabilities also
  highlighting the issues present. I hope you liked the project, please do try it
  and comment the feedback please. Thank you for reading, Until next time, Happy Coding
  :)</p>\n"
cover: ''
date: 2021-10-31
datetime: 2021-10-31 00:00:00+00:00
description: If you have been writing articles you know the pain to get some attention,
  if you have already been cross-posting your articles it usually takes some time
  to do
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-10-31-Crossposter.sh.md
html: "<h2>Introduction</h2>\n<p>If you have been writing articles you know the pain
  to get some attention, if you have already been cross-posting your articles it usually
  takes some time to do that. This task can be automated with a shellscript. If you
  have been cross-posting articles on <code>medium.com</code>, <code>dev.to</code>
  and at <code>hashnode.com</code>, then I have a treat for you.</p>\n<p>Introducing
  <strong><a href=\"http://crossposter.sh\">crossposter.sh</a></strong>!!</p>\n<h2>What
  is <a href=\"http://Crossposter.sh\">Crossposter.sh</a>?</h2>\n<h3>Crosspost to
  <a href=\"http://dev.to/hahsnode/medium\">dev.to/hahsnode/medium</a> from the command
  line.</h3>\n<p><a href=\"http://Crossposter.sh\">Crossposter.sh</a> is a shellscript(BASH)
  to automate crossposting to platforms like <a href=\"http://dev.to\">dev.to</a>,
  <a href=\"http://medium.com\">medium.com</a> and <a href=\"http://hashnode.com\">hashnode.com</a>.
  The script takes in markdown version of your post with a few inputs from you and
  posts it to those platforms. You would require a token/key for each of those platforms
  to post it from the command line. You can check out the official repository of <a
  href=\"https://github.com/Mr-Destructive/crossposter\">Crossposter</a>.</p>\n<p>The
  actual script is still not perfect (has a few bugs). Though it posts on <code>dev.to</code>
  and <code>medium.com</code> easily, the <code>hashnode.com</code> is buggy as it
  parses the raw markdown into the post and doesn't render as desired. So, <strong>its
  a under-development script</strong>, fell free to raise any issues or PRs on the
  official GitHub repo.</p>\n<p>Run the script on a bash interpreter with the command:</p>\n<p><code>bash
  crosspost.sh</code></p>\n<p>For posting the article you need to provide the following
  details:</p>\n<h2>Front-Matter</h2>\n<h3>Meta data about the post</h3>\n<ul>\n<li>Title
  of Post</li>\n<li>Subtitle of Post</li>\n<li>Publish status of post(<code>true</code>
  or <code>false</code>)</li>\n<li>Tags for the post (comma separated values)</li>\n<li>Canonical
  Url (original url of the post)</li>\n<li>Cover Image (URL of the post's image/thumbnail)</li>\n</ul>\n<p>This
  information is a must for <code>dev.to</code> especially the <code>title</code>.
  This should be provide in the same order as given below:</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>yaml</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nn\">---</span><span class=\"w\"></span>\n<span class=\"nt\">title</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"l l-Scalar l-Scalar-Plain\">The
  title of the post</span><span class=\"w\"></span>\n<span class=\"nt\">subtitle</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"l l-Scalar l-Scalar-Plain\">The
  description of your article</span><span class=\"w\"></span>\n<span class=\"nt\">published</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"l l-Scalar l-Scalar-Plain\">true</span><span
  class=\"w\"></span>\n<span class=\"nt\">tags</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"l l-Scalar l-Scalar-Plain\">programming, anythingelse</span><span
  class=\"w\"></span>\n<span class=\"nt\">canonical url</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"l l-Scalar l-Scalar-Plain\">url of your original
  blog</span><span class=\"w\"></span>\n<span class=\"nt\">cover_image</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"l l-Scalar l-Scalar-Plain\">coverimage_url</span><span
  class=\"w\"></span>\n<span class=\"nn\">---</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>There
  is no need to enclose any of them with quotation marks. <code>Published</code> argument
  will be <code>true</code> if you want to publish it and <code>false</code> if you
  want to keep it in your Drafts.</p>\n<p>In the demonstrations, we just need to enter
  the tokens once. The tokens will be stored locally in the <code>keys.txt</code>
  file and retrieved later within the script.</p>\n<h2>Posting on <strong><a href=\"http://dev.to\">dev.to</a></strong>:</h2>\n<p>Posting
  on <a href=\"http://dev.to\">dev.to</a> requires their <code>API key</code> which
  can be generated by going on the <a href=\"https://dev.to/settings/account/\">Dev
  Community API Keys</a>. From there you can generate a new key with any name you
  like. You just need to enter the key to CLI once or manually enter in the <code>keys.txt</code>
  file with the format <code>dev.to:key</code> on the first line. This will be used
  for the future cross-posting whenever you execute the shell script(<code>bash crosspost.sh</code>)</p>\n<p>You
  can provide the <a href=\"#front-matter\">front matter</a> manually in your markdown
  file or you will be prompted for the input. So, that is all you will require for
  posting on <a href=\"http://dev.to\">dev.to</a> from the Command line.</p>\n<p>Lets
  see the script in action</p>\n<p><img src=\"https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/devto.gif\"
  alt=\"dev.to\" /></p>\n<p>If you want to add in more stuff to the post, you can
  check out the <a href=\"https://developers.forem.com/api#operation/createArticle\">DEV.to
  API docs</a> which is powered by <a href=\"https://www.forem.com/\">Forem</a>, there
  a ton of options you can hook to the front-matter in the shellscript.</p>\n<p><strong>NOTE:
  There is a limit of 10 requests per 30 seconds, so keep in mind while testing the
  script and don't try to spam</strong></p>\n<h2>Posting on <strong><a href=\"http://hashnode.com\">hashnode.com</a></strong>:</h2>\n<p>This
  part is still under development as it only displays the raw markdown in the post,
  also the <code>tags</code> are too heavy to implement from the API as <code>id</code>
  of every tag is required along with the <code>slug</code> and <code>name</code>.
  Still it serves some purpose at least. For posting on <code>hashnode.com</code>,
  we need <code>Personal Access Token</code>. This can be generated by going to <a
  href=\"https://hashnode.com/settings/developer\">Developer Settings</a>. You will
  also require the user-id of your <code>hashnode</code> account. You can get your
  user-id/username from the <a href=\"https://hashnode.com/settings\">settings</a>
  tab under profile information. We require Username for posting to the Publication
  Blog if any. As usual, the <code>Personal Access Token</code> for interacting with
  the <a href=\"https://api.hashnode.com/\">Hashnodes' GraphQL API</a>. The API is
  quite user friendly and provides everything in one place. There are docs for running
  each and every <code>query</code> and <code>mutations</code> present in the API.</p>\n<p>You
  can paste the token when prompted from the script or manually type in the <code>keys.txt</code>
  text file as <code>hashnode:token</code> on the 4th line. Yes, that should be on
  the <code>4th</code> line, thats make retrieving much more easier and safe. Next
  also input in the <code>username</code> when the script asks for the input or again
  type in on the <code>5th</code> line, <code>hashnode_id:username</code> in the text
  file <code>keys.txt</code>. Please enter the credentials from the script prompt
  so as to avoid errors and misconfigurations when doing manually</p>\n<p>This will
  create the Post on hashnode with the title, subtitle, cover image correctly but
  will mess up the content. I tried hard but its just not happening. There needs to
  be some character for newline as the API rejects the <code>rn</code> characters
  passed in, so I have substited them with <code>br</code> and the result is raw markdown.
  <strong>As the Hashnode API is still under development and they are bringing changes
  and new features in, the API should improve in its core functionality and make it
  much easier for creating some common queries</strong>. So, I'll create issue on
  GitHub for posting the actual content via the script.</p>\n<p>So, this is the demonstration
  of posting on hashnode.</p>\n<p><img src=\"https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/hashnode.gif\"
  alt=\"hashnode\" /></p>\n<h2>Posting on <strong><a href=\"http://medium.com\">medium.com</a></strong>:</h2>\n<p>Medium
  API is much more versatile and markdown friendly, though it has some limitation
  on the number of posts you can make in a day. For posting on <code>Medium.com</code>,
  we will require the <code>Integration Token</code> which can be generated on the
  <a href=\"https://medium.com/me/settings\">settings tab</a>. As similar to <code>hashnode</code>,
  you can name the token whatever you like and then get the token. Paste the token
  when prompted from the script or manually type in the <code>keys.txt</code> text
  file as <code>medium:token</code> on the <code>2nd</code> line. We also require
  the Medium_id, but we can get that from the token itself, so inside the script once
  the token is obtained, the curl command is executed to fetch in the <code>id</code>
  and it is stored on the next(<code>3rd</code>) line in the <code>keys.txt</code>
  file for actually posting on <code>medium.com</code>. So that is all the configuration
  you need for posting on <code>medium.com</code>.</p>\n<p>There is some documentation
  on <a href=\"https://github.com/Medium/medium-api-docs\">Medium API</a>, we can
  even post to a Publication, that shall be created in future. Also the cover images
  can be posted on medium, it is not currently done but that can again be a #TODO.
  <strong>The tags are not rendered on Medium yet with the script.</strong> The way
  we can parse  strings is limited in BASH, so this might still be a doable thing
  later. Most of the checkboxes are ticked like title, subtitle, cover-image, canonical
  url, and importantly the content.</p>\n<p>Let's look at post on medium from the
  script.</p>\n<p><img src=\"https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/medium.gif\"
  alt=\"medium\" /></p>\n<h2>All platforms:</h2>\n<p>Now, once you have configured
  every thing, you can opt for the <code>4</code> choice that is post on all platforms(<a
  href=\"http://dev.to\">dev.to</a>, hashnode and medium), but as hashnode is not
  looking a good option right now, so there is the <code>5</code> option for only
  <code>dev.to</code> and <code>medium</code>.</p>\n<p><img src=\"https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/crossposter.gif\"
  alt=\"allplatforms\" /></p>\n<h2>Why use <a href=\"http://Crossposter.sh\">Crossposter.sh</a>?</h2>\n<p>This
  can be not so big of an issue for most of the people but it was a good side project
  to work and learn more about how APIs work and get some ideas on the design of the
  platform. Though it is quite time saving to cross post on 3 different platforms
  within a minute or two. You can tailor your own script as per your specifications
  and desire.</p>\n<p>So, if you are an author on all of the mentioned platforms,
  please give it a try. Other Platforms are welcome for contributions. If you found
  any unexpected things, please hit them in the <code>issues</code> tab.</p>\n<h2>Script</h2>\n<p>The
  script mostly leverages <code>curl</code>, <code>sed</code>, <code>cat</code> and
  some other basic utilities in BASH.</p>\n<h3>Using <code>curl</code> for posting
  the article from APIs</h3>\n<p>Curl is a life saver command for this project, without
  this tool, the project might not be as great and efficient. Let's see some quick
  commands used in the script.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>curl
  -H <span class=\"s2\">&quot;Content-Type: application/json&quot;</span> -H <span
  class=\"s2\">&quot;api-key&quot;</span>: <span class=\"se\">\\&quot;</span><span
  class=\"s1\">&#39;&quot;$key&quot;&#39;</span><span class=\"se\">\\&quot;</span>
  -d <span class=\"s1\">&#39;{&quot;content&quot;:\\&quot;&#39;</span><span class=\"s2\">&quot;</span><span
  class=\"nv\">$body</span><span class=\"s2\">&quot;</span><span class=\"s1\">&#39;\\&quot;}&#39;</span>
  <span class=\"s2\">&quot;</span><span class=\"nv\">$url</span><span class=\"s2\">&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>So,
  the above command is quite basic, some more additions are also added as per the
  specifications of the Platform. But, let us understand the structure of the command
  we are sending to the APIs. The first part is the Header (<code>-H</code>), in here
  we specify the content that is going to get parsed and the api-keys to access the
  API. Next, we have the body or the data (<code>-d</code>), here we parse in the
  actual contents, it might be the front matter along with the markdown content. Finally
  we have the <code>url</code> where we send the <code>POST</code> request i.e. the
  <code>API endpoint</code>. The is the escape character that is used to preserve
  the literal value of the next character and in short we can shorten the command
  to fit in the next line.</p>\n<p>The wired <code>$body</code> is used to parse the
  value of the variable <code>body</code> inside of single quotes as in BASH, we can
  only access the variables' value in double quotes. We are using single quotes as
  we have to pass the <code>json</code> object and which has already double quotes
  in it.</p>\n<h3>Using <code>sed</code> for editing text</h3>\n<p>Sed is a super-powerful
  stream editor, its somewhat similar to Vim without an interface, only commands.
  We use this tool to manipulate the front-matter for posting on the platforms by
  parsing them to variables in BASH. We also use to enter the api keys inputted by
  user from variables into the file at a specific position to retrieve later.</p>\n<pre
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sed
  -i <span class=\"s2\">&quot;1a title: </span><span class=\"nv\">$title</span><span
  class=\"s2\">&quot;</span> file.txt\n</pre></div>\n\n</pre>\n\n<p>Here, we are appending(<code>a</code>)
  to the 1st line, text <code>title: $title</code>, here <code>$title</code> is the
  variable, so we are technically parsing the value of the variable <code>title</code>.
  We are editing the file <code>file.txt</code> in-place <code>-i</code> i.e. we are
  editing it live without creating any temp or backup files.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>sed
  -n -e <span class=\"s2\">&quot;s/dev.to://p&#39; keys.txt&quot;</span>\n</pre></div>\n\n</pre>\n\n<p>Here
  we are essentially getting the text after a particular pattern. In this case we
  are searching in <code>keys.txt</code> file for the string <code>dev.to:</code>
  and anything after that till the end of line is returned, we can further store it
  in the variable and do all sorts of operation.</p>\n<h3>Using <code>awk</code> for
  programmatic editing</h3>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>awk
  <span class=\"s1\">&#39;{print $0&quot;\\r\\n&quot;}&#39;</span> temp.txt &gt;file.txt\n</pre></div>\n\n</pre>\n\n<p>AWK
  is a command-line utility for manipulating or writing certain operations/patterns
  programmatically. We use this tool so as to add <code>4r4n</code> to the end of
  each line, the APIs can't parse the file contents directly so we have to add certain
  characters before the end of line and do further operations.</p>\n<pre class='wrapper'>\n\n<div
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
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span>cat
  temp.md <span class=\"p\">|</span> tr -d <span class=\"s1\">&#39;\\r\\n&#39;</span>
  &gt; temp.txt\n</pre></div>\n\n</pre>\n\n<p>After we have added the <code>\\r\\n</code>
  characters to the end of the file, we simply can use <code>cat</code> and <code>tr</code>
  to merge all the lines into a single line. This is how we parse the contents to
  the API more safely and concisely, of course we need to parse them in a variable
  by reading the file.</p>\n<p>OK, I won't bore anyone with more BASH but that were
  some of the quite important commands in the script that form the backbone of the
  cross-posting and handling text with the APIs.</p>\n<h2>Conclusion</h2>\n<p>So,
  we can see <code>crosspost.sh</code> is a BASH script that cross-posts markdown
  articles with a bit of inputs to 3 different platforms within a couple of minutes.
  This article was basically to demonstrate the project and its capabilities also
  highlighting the issues present. I hope you liked the project, please do try it
  and comment the feedback please. Thank you for reading, Until next time, Happy Coding
  :)</p>\n"
image_url: https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643288989/blogmedia/trssl38erkdbcqlnjdvp.png
long_description: If you have been writing articles you know the pain to get some
  attention, if you have already been cross-posting your articles it usually takes
  some time to do that. This task can be automated with a shellscript. If you have
  been cross-posting artic
now: 2025-05-01 18:11:33.312008
path: blog/posts/2021-10-31-Crossposter.sh.md
prevnext: null
slug: crossposter-shellscript
status: published
subtitle: Crosspost on dev.to, medium.com and hashnode.com with a single click
tags:
- bash
templateKey: blog-post
title: 'Crossposting with a single script: Crossposter.sh'
today: 2025-05-01
---

## Introduction

If you have been writing articles you know the pain to get some attention, if you have already been cross-posting your articles it usually takes some time to do that. This task can be automated with a shellscript. If you have been cross-posting articles on `medium.com`, `dev.to` and at `hashnode.com`, then I have a treat for you. 

Introducing **crossposter.sh**!!

## What is Crossposter.sh?

### Crosspost to dev.to/hahsnode/medium from the command line.

Crossposter.sh is a shellscript(BASH) to automate crossposting to platforms like dev.to, medium.com and hashnode.com. The script takes in markdown version of your post with a few inputs from you and posts it to those platforms. You would require a token/key for each of those platforms to post it from the command line. You can check out the official repository of [Crossposter](https://github.com/Mr-Destructive/crossposter).

The actual script is still not perfect (has a few bugs). Though it posts on `dev.to` and `medium.com` easily, the `hashnode.com` is buggy as it parses the raw markdown into the post and doesn't render as desired. So, **its a under-development script**, fell free to raise any issues or PRs on the official GitHub repo.   

Run the script on a bash interpreter with the command:

`bash crosspost.sh`

For posting the article you need to provide the following details:

## Front-Matter

### Meta data about the post

- Title of Post
- Subtitle of Post
- Publish status of post(`true` or `false`)
- Tags for the post (comma separated values)
- Canonical Url (original url of the post)
- Cover Image (URL of the post's image/thumbnail)

This information is a must for `dev.to` especially the `title`. This should be provide in the same order as given below:

```yaml

---
title: The title of the post
subtitle: The description of your article
published: true
tags: programming, anythingelse
canonical url: url of your original blog
cover_image: coverimage_url
---
```

There is no need to enclose any of them with quotation marks. `Published` argument will be `true` if you want to publish it and `false` if you want to keep it in your Drafts.

In the demonstrations, we just need to enter the tokens once. The tokens will be stored locally in the `keys.txt` file and retrieved later within the script.

## Posting on **dev.to**:

Posting on dev.to requires their `API key` which can be generated by going on the [Dev Community API Keys](https://dev.to/settings/account/). From there you can generate a new key with any name you like. You just need to enter the key to CLI once or manually enter in the `keys.txt` file with the format `dev.to:key` on the first line. This will be used for the future cross-posting whenever you execute the shell script(`bash crosspost.sh`)

You can provide the [front matter](#front-matter) manually in your markdown file or you will be prompted for the input. So, that is all you will require for posting on dev.to from the Command line.  

Lets see the script in action 

![dev.to](https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/devto.gif)
   
If you want to add in more stuff to the post, you can check out the [DEV.to API docs](https://developers.forem.com/api#operation/createArticle) which is powered by [Forem](https://www.forem.com/), there a ton of options you can hook to the front-matter in the shellscript.

**NOTE: There is a limit of 10 requests per 30 seconds, so keep in mind while testing the script and don't try to spam**

## Posting on **hashnode.com**:

This part is still under development as it only displays the raw markdown in the post, also the `tags` are too heavy to implement from the API as `id` of every tag is required along with the `slug` and `name`. Still it serves some purpose at least. For posting on `hashnode.com`, we need `Personal Access Token`. This can be generated by going to [Developer Settings](https://hashnode.com/settings/developer). You will also require the user-id of your `hashnode` account. You can get your user-id/username from the [settings](https://hashnode.com/settings) tab under profile information. We require Username for posting to the Publication Blog if any. As usual, the `Personal Access Token` for interacting with the [Hashnodes' GraphQL API](https://api.hashnode.com/). The API is quite user friendly and provides everything in one place. There are docs for running each and every `query` and `mutations` present in the API. 

You can paste the token when prompted from the script or manually type in the `keys.txt` text file as `hashnode:token` on the 4th line. Yes, that should be on the `4th` line, thats make retrieving much more easier and safe. Next also input in the `username` when the script asks for the input or again type in on the `5th` line, `hashnode_id:username` in the text file `keys.txt`. Please enter the credentials from the script prompt so as to avoid errors and misconfigurations when doing manually
 
This will create the Post on hashnode with the title, subtitle, cover image correctly but will mess up the content. I tried hard but its just not happening. There needs to be some character for newline as the API rejects the `rn` characters passed in, so I have substited them with `br` and the result is raw markdown. **As the Hashnode API is still under development and they are bringing changes and new features in, the API should improve in its core functionality and make it much easier for creating some common queries**. So, I'll create issue on GitHub for posting the actual content via the script.

So, this is the demonstration of posting on hashnode.

![hashnode](https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/hashnode.gif)


## Posting on **medium.com**:

Medium API is much more versatile and markdown friendly, though it has some limitation on the number of posts you can make in a day. For posting on `Medium.com`, we will require the `Integration Token` which can be generated on the [settings tab](https://medium.com/me/settings). As similar to `hashnode`, you can name the token whatever you like and then get the token. Paste the token when prompted from the script or manually type in the `keys.txt` text file as `medium:token` on the `2nd` line. We also require the Medium_id, but we can get that from the token itself, so inside the script once the token is obtained, the curl command is executed to fetch in the `id` and it is stored on the next(`3rd`) line in the `keys.txt` file for actually posting on `medium.com`. So that is all the configuration you need for posting on `medium.com`.

There is some documentation on [Medium API](https://github.com/Medium/medium-api-docs), we can even post to a Publication, that shall be created in future. Also the cover images can be posted on medium, it is not currently done but that can again be a #TODO. **The tags are not rendered on Medium yet with the script.** The way we can parse  strings is limited in BASH, so this might still be a doable thing later. Most of the checkboxes are ticked like title, subtitle, cover-image, canonical url, and importantly the content.

Let's look at post on medium from the script.

![medium](https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/medium.gif)


## All platforms: 

Now, once you have configured every thing, you can opt for the `4` choice that is post on all platforms(dev.to, hashnode and medium), but as hashnode is not looking a good option right now, so there is the `5` option for only `dev.to` and `medium`. 

![allplatforms](https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/crossposter.gif)

## Why use Crossposter.sh?

This can be not so big of an issue for most of the people but it was a good side project to work and learn more about how APIs work and get some ideas on the design of the platform. Though it is quite time saving to cross post on 3 different platforms within a minute or two. You can tailor your own script as per your specifications and desire.   

So, if you are an author on all of the mentioned platforms, please give it a try. Other Platforms are welcome for contributions. If you found any unexpected things, please hit them in the `issues` tab. 

## Script

The script mostly leverages `curl`, `sed`, `cat` and some other basic utilities in BASH. 

### Using `curl` for posting the article from APIs

Curl is a life saver command for this project, without this tool, the project might not be as great and efficient. Let's see some quick commands used in the script.


```bash
curl -H "Content-Type: application/json" -H "api-key": \"'"$key"'\" -d '{"content":\"'"$body"'\"}' "$url"
```

So, the above command is quite basic, some more additions are also added as per the specifications of the Platform. But, let us understand the structure of the command we are sending to the APIs. The first part is the Header (`-H`), in here we specify the content that is going to get parsed and the api-keys to access the API. Next, we have the body or the data (`-d`), here we parse in the actual contents, it might be the front matter along with the markdown content. Finally we have the `url` where we send the `POST` request i.e. the `API endpoint`. The is the escape character that is used to preserve the literal value of the next character and in short we can shorten the command to fit in the next line.

The wired `$body` is used to parse the value of the variable `body` inside of single quotes as in BASH, we can only access the variables' value in double quotes. We are using single quotes as we have to pass the `json` object and which has already double quotes in it.  

### Using `sed` for editing text

Sed is a super-powerful stream editor, its somewhat similar to Vim without an interface, only commands. We use this tool to manipulate the front-matter for posting on the platforms by parsing them to variables in BASH. We also use to enter the api keys inputted by user from variables into the file at a specific position to retrieve later. 
   
```bash
sed -i "1a title: $title" file.txt
```   


Here, we are appending(`a`) to the 1st line, text `title: $title`, here `$title` is the variable, so we are technically parsing the value of the variable `title`. We are editing the file `file.txt` in-place `-i` i.e. we are editing it live without creating any temp or backup files.       

```bash
sed -n -e "s/dev.to://p' keys.txt"
```

Here we are essentially getting the text after a particular pattern. In this case we are searching in `keys.txt` file for the string `dev.to:` and anything after that till the end of line is returned, we can further store it in the variable and do all sorts of operation.
      
### Using `awk` for programmatic editing 

```bash
awk '{print $0"\r\n"}' temp.txt >file.txt
```

AWK is a command-line utility for manipulating or writing certain operations/patterns programmatically. We use this tool so as to add `4r4n` to the end of each line, the APIs can't parse the file contents directly so we have to add certain characters before the end of line and do further operations.

```bash
cat temp.md | tr -d '\r\n' > temp.txt
```   

After we have added the `\r\n` characters to the end of the file, we simply can use `cat` and `tr` to merge all the lines into a single line. This is how we parse the contents to the API more safely and concisely, of course we need to parse them in a variable by reading the file.

OK, I won't bore anyone with more BASH but that were some of the quite important commands in the script that form the backbone of the cross-posting and handling text with the APIs.

## Conclusion

So, we can see `crosspost.sh` is a BASH script that cross-posts markdown articles with a bit of inputs to 3 different platforms within a couple of minutes. This article was basically to demonstrate the project and its capabilities also highlighting the issues present. I hope you liked the project, please do try it and comment the feedback please. Thank you for reading, Until next time, Happy Coding :)