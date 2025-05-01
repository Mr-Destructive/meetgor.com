---
article_html: "<h2>Introduction</h2>\n<blockquote>\n<p>No one can understand the joy
  in finishing a personal project, except the dreamer</p>\n</blockquote>\n<p>It was
  a while,since I have posted an article here, as I was busy on a project or a hackathon.</p>\n<p>Hello,
  world! I am Meet a student and a self-taught web developer. I like to make and break
  stuff, especially when it comes to programming and Linux. I like shell scripting
  and learning different languages at once, love to learn about Vim and Linux everyday.</p>\n<p>Every
  time I start a project something else comes and distracts me let that be any other
  programming language or technology. That leads to creating new projects and leaving
  the one behind unfinished, I know most of the developers face this.  But this time,
  thanks to Auth0 X Hashnode Hackathon, I was able to create an almost finished project
  within almost 10 days. Having a deadline and competition creates a mindset to finish
  a project on time, that's my first takeaway from this Hackathon. OH! this is my
  first Hackathon by the way, and it has been amazing so far.</p>\n<p>** Applying
  a framework to do something you desire and then everything working smoothly (after
  fixing 100s of bugs) is such a great feeling that no one can understand except for
  the person who just dreamt of it. **</p>\n<p>I'll like to share my project which
  is a web application for the Auth0 x Hashnode Hackathon. Here it goes.</p>\n<h2>What
  is Dev Quotes?</h2>\n<p>Dev quotes is a web app designed for publishing and viewing
  quotes related to programming, developer mindset, and all the technicalities involved
  in a developer's life. It's basically a medium to express the life of developers
  and get inspired by others.  Here it is <a href=\"https://devquotess.herokuapp.com/\">devquotes</a></p>\n<h4>Dark
  Mode:</h4>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1630074051548/TQz9Koh7l.png\"
  alt=\"image.png\" /></p>\n<h4>Light Mode:</h4>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1630078314355/VhfLrcjJa.png\"
  alt=\"image.png\" /></p>\n<h2>Why Dev Quotes?</h2>\n<blockquote>\n<p>Developers
  are not the people who only understand how to write code but they're also the people
  who can make the code understandable</p>\n</blockquote>\n<p>As a developer, there
  are often times where you have no motivation left inside, but you never know you
  might be just a few lines of code away from making another project or fixing a bug.
  For that, we require some inspiration or a push to break the barrier of.  I am not
  saying it's just for developers, it's designed for developers but everyone is open
  to understanding the developers' lives and their struggles.</p>\n<p>I also felt
  the need to give back some love-crafted web app to the ever-wonderful and supportive
  dev community. It's a small application but still, I would like to give in something
  instead of nothing at all. Start small grow big, hopefully :)</p>\n<h2>Features</h2>\n<p>Some
  of the main features of the web application are as follows:</p>\n<ul>\n<li>\n<p><strong>Write\\Edit\\Delete
  Quotes if logged in.</strong></p>\n</li>\n<li>\n<p><strong>Like / Unlike a Quote.</strong></p>\n</li>\n<li>\n<p><strong>See
  all of your quotes.</strong></p>\n</li>\n<li>\n<p><strong>Randomized Quotes on Homepage.</strong></p>\n</li>\n<li>\n<p><strong>Dark/Light
  theme based on Browser's Preference and local storage.</strong></p>\n</li>\n<li>\n<p><strong>The
  app is mobile responsive as well, though the navbar is a bit wonky with the light/dark
  mode switch toggle, which can be taken care of soon.</strong></p>\n</li>\n</ul>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1630084573444/YEj38dUTD_.png\"
  alt=\"dqmob.png\" /></p>\n<h2>In the making</h2>\n<blockquote>\n<p>Have the curiosity
  to learn, rest is automated</p>\n</blockquote>\n<p>The project was made by using
  various inspirational articles and videos about making a web application. But the
  idea stuck in my mind when I was thinking about the people who don't get inspired
  as a developer. Like there is no way you can remain sad about being a developer
  and keep on dealing with imposter syndrome. Every developer has a perspective of
  programming but there is an infinite number of opportunities if you are curious
  enough. Just started making the project and got so much into it that I was literally
  dreaming about it like I saw parts of the webpage. In my dream and I am making it
  that was genuinely a thing that powered me to complete it.</p>\n<p>The project roughly
  started on 19th August and almost ended on 26th August, like the actual webpage
  and its core functionalities. Though on 27th were some styling and extra additions
  such as the About section and Footer. That was like the most productive week I ever
  had in my programming journey. That was fun as heck.</p>\n<h2>Tech Stack</h2>\n<p>The
  Tech Stack involved with this app is :</p>\n<ul>\n<li><code>Django</code></li>\n<li><code>PostgreSQL</code></li>\n<li><code>HTML/CSS/JS</code></li>\n<li><code>
  Bootstrap</code></li>\n</ul>\n<p>I have not used any front-end end frameworks just
  because I never found the need to learn them.  I had experience with Django for
  just 2 months and I am surprised I was able to make it. As obvious I have used Auth0
  for authentication in my web application.</p>\n<h3>Auth0 integration for Authentication</h3>\n<p>I
  must tell you using Auth0 was just flawless addition to my app as I have to do almost
  nothing, just drop some credentials of the Auth0 application with my Django project
  using a  <a href=\"https://auth0.com/docs/quickstarts\">well-documented guide</a>
  \ for every type of framework. Simply straight-forward was the name for integrating
  authentication in my app.</p>\n<h4>How I used Auth0 with Django</h4>\n<p>I've used
  Template tags such as if blocks to verify if the user is authenticated or not.</p>\n<pre
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
  class=\"p\">&lt;</span><span class=\"nt\">div</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;login-box auth0-box before&quot;</span><span
  class=\"p\">&gt;</span>\n\t\t{{ &quot;{% if user.is_authenticated &quot;}} %}\n\t\t
  \   <span class=\"p\">&lt;</span><span class=\"nt\">a</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;btn btn-primary btn-sm tn-logout &quot;</span>
  <span class=\"na\">href</span><span class=\"o\">=</span><span class=\"s\">&quot;/logout&quot;</span><span
  class=\"p\">&gt;</span>Log Out<span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n\t\t{{ &quot;{% else &quot;}} %}\n\t\t    <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;btn btn-primary btn-sm tn-login &quot;</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;/login/auth0&quot;</span><span class=\"p\">&gt;</span>Log
  In<span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span class=\"p\">&gt;</span>\n
  \   {{ &quot;{% endif &quot;}} %}\n<span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>This was just readily available
  on their documentation though there were some adjustments as per the project requirements
  in this code to fit in the place.</p>\n<p>I must say, integrating Auth0 is even
  easier than using Django User Model in some sense as most of the stuff is handled
  by the Auth0, on our side, we simply have to create the Auth0 specific app with
  the credentials from the dashboard rest just works flawlessly till now. How sweet
  and</p>\n<h3>Specifications</h3>\n<p>I won't go in-depth about the technicalities
  of the project but would like to address certain things. Firstly I have mostly used
  Class-based views for the major part, certain areas are still function-based just
  for the simplicity of the application and a few of them are handled and documented
  by Auth0 so just preferred that.</p>\n<p>Another thing is about Models, I just have
  a simple single model called <code>Quote</code> which has an Author as a Foreign
  Key from the Django User Model. I would have also created multiple emojis for the
  like system but I was too excited and in a rush to see the actual app, So just kept
  it simple. XD\n<img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1630060555499/246ZKnypP.png\"
  alt=\"image.png\" /></p>\n<p>The rest of the stuff like <code>URLs</code>, <code>Templates</code>,
  and <code>static files</code> are handled in a neatly organized way as depicted
  in the below diagram.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1630060426600/JHMlrfOKZ.png\"
  alt=\"image.png\" /></p>\n<p>As it might not be clear from that, there are 3 apps
  -&gt; <code>auth0login</code>, <code>quotes</code>, and <code>user</code>, here
  <code>quotes</code> is kind of the most important app as it has the models, forms,
  URLs, and the views linked to them here.</p>\n<h3>Hosting</h3>\n<p>Hosting as you
  can guess, it's on  <a href=\"https://www.heroku.com/\">Heroku</a> , quite beginner-friendly
  and feature-rich. I also have a free addon for PostgreSQL Database here.  It's limited
  to 10K rows but that's sufficient for a starter app like this in my opinion. Also,
  it has 500 free hours of dyno, which is also sufficient for a small low-traffic
  app like this.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1630076036133/9ofxnM5VN.png\"
  alt=\"image.png\" /></p>\n<h2>Bugs Encountered</h2>\n<blockquote>\n<p>Love Bugs,
  they'll give you experience</p>\n</blockquote>\n<p>This is the most exciting and
  important part if you consider the hackathon because this determines the level of
  experience that a developer shoes in achieving certain things or features if you
  will. Faced some typical Django errors as usual but also some unexpected things
  like Dark mode and light mode clashing together due to poorly written media queries
  in CSS.</p>\n<p>As usual, the start is often hard, there is no motivation whatsoever
  in using the admin section to test the database queries and gibberish basic HTML
  page. In that process, I faced some primary key access issues and was able to understand
  the concept more clearly by fixing it.</p>\n<p>Another instance was with handling
  post requests in Django which I've failed to do before. I used forms and a hybrid
  of CSS and bootstrap to style those forms which just works brilliantly. That took
  some time to figure out the exact working but after a while, it was working charms.</p>\n<h2>Future
  Updates</h2>\n<p>As said, I would like to add more like buttons expressing different
  emotions. Some other features to add are:</p>\n<ul>\n<li>To add more emojis like
  claps, cheers, and others.</li>\n<li>To add a profile page in the app that would
  display all the quotes of the particular author and the details related to him/her.</li>\n<li>Adding
  some tags to filter out particular types of quotes.</li>\n<li>Improve UI-UX a bit
  more to make it professional and pleasant.</li>\n</ul>\n<h2>Source Code</h2>\n<blockquote>\n<p>Talk
  is cheap, show me the code - Linus Torvalds</p>\n</blockquote>\n<p>The source code
  is available at GitHub on this  <a href=\"https://github.com/Mr-Destructive/devquotes\">Link</a>.\nIt's
  freely open for any contribution after the hackathon(mid-September).  Some of the
  files such as the environment variables, virtual environments, cached data are not
  uploaded for security and obvious reasons.</p>\n<p>Enough of technical talks, let's
  see the DEMO,</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1630073466386/3wgnST5hc.gif\"
  alt=\"dqss.gif\" /></p>\n<p>Silent claps.......</p>\n<p>It's not a great UI-UX but
  works and is not too shabby in my opinion considering it only has base HTML and
  CSS with a little Bootstrap and Javascript. But ya, a fully functional Backend that's
  what I was looking for a full-stack app with some decent features. Hope it helps
  some developers stay motivated and hack into some hackathons like this.</p>\n<h3>References
  used while creating the app:</h3>\n<ul>\n<li><a href=\"https://www.youtube.com/watch?v=B40bteAMM_M&amp;list=PLCC34OHNcOtr025c1kHSPrnP18YPB-NFi\">Codemy
  -John Elder Django tutorial</a></li>\n<li><a href=\"https://docs.djangoproject.com/en/3.2/topics/forms/\">Django
  - Documentation for Forms</a></li>\n<li><a href=\"https://stackoverflow.com/questions/28837511/django-template-how-to-randomize-order-when-populating-page-with-objects\">Django
  template randomizer shuffle</a></li>\n<li><a href=\"https://www.youtube.com/watch?v=kzN_VCFG9NM\">Auth0
  app Django integration</a></li>\n</ul>\n<h2>Closing Words</h2>\n<blockquote>\n<p>Why
  developers find solutions to bugs at the stroke of sleeping, that's multithreading
  in our brains</p>\n</blockquote>\n<p>Hope you liked the project and hopefully will
  inspire developers to stay motivated and can focus on their goals more than dealing
  with imposter syndrome and whatnot.</p>\n<p>Thank you for reading and using the
  app, for any feedbacks, Twitter handles, comment section, GitHub issues, LinkedIn
  messages are all freely open. Thanks. Happy Coding :)</p>\n"
cover: ''
date: 2021-08-27
datetime: 2021-08-27 00:00:00+00:00
description: No one can understand the joy in finishing a personal project, except
  the dreamer It was a while,since I have posted an article here, as I was busy on
  a project
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-08-27-DevQuotes-Auth0-x-Hashnode.md
html: "<h2>Introduction</h2>\n<blockquote>\n<p>No one can understand the joy in finishing
  a personal project, except the dreamer</p>\n</blockquote>\n<p>It was a while,since
  I have posted an article here, as I was busy on a project or a hackathon.</p>\n<p>Hello,
  world! I am Meet a student and a self-taught web developer. I like to make and break
  stuff, especially when it comes to programming and Linux. I like shell scripting
  and learning different languages at once, love to learn about Vim and Linux everyday.</p>\n<p>Every
  time I start a project something else comes and distracts me let that be any other
  programming language or technology. That leads to creating new projects and leaving
  the one behind unfinished, I know most of the developers face this.  But this time,
  thanks to Auth0 X Hashnode Hackathon, I was able to create an almost finished project
  within almost 10 days. Having a deadline and competition creates a mindset to finish
  a project on time, that's my first takeaway from this Hackathon. OH! this is my
  first Hackathon by the way, and it has been amazing so far.</p>\n<p>** Applying
  a framework to do something you desire and then everything working smoothly (after
  fixing 100s of bugs) is such a great feeling that no one can understand except for
  the person who just dreamt of it. **</p>\n<p>I'll like to share my project which
  is a web application for the Auth0 x Hashnode Hackathon. Here it goes.</p>\n<h2>What
  is Dev Quotes?</h2>\n<p>Dev quotes is a web app designed for publishing and viewing
  quotes related to programming, developer mindset, and all the technicalities involved
  in a developer's life. It's basically a medium to express the life of developers
  and get inspired by others.  Here it is <a href=\"https://devquotess.herokuapp.com/\">devquotes</a></p>\n<h4>Dark
  Mode:</h4>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1630074051548/TQz9Koh7l.png\"
  alt=\"image.png\" /></p>\n<h4>Light Mode:</h4>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1630078314355/VhfLrcjJa.png\"
  alt=\"image.png\" /></p>\n<h2>Why Dev Quotes?</h2>\n<blockquote>\n<p>Developers
  are not the people who only understand how to write code but they're also the people
  who can make the code understandable</p>\n</blockquote>\n<p>As a developer, there
  are often times where you have no motivation left inside, but you never know you
  might be just a few lines of code away from making another project or fixing a bug.
  For that, we require some inspiration or a push to break the barrier of.  I am not
  saying it's just for developers, it's designed for developers but everyone is open
  to understanding the developers' lives and their struggles.</p>\n<p>I also felt
  the need to give back some love-crafted web app to the ever-wonderful and supportive
  dev community. It's a small application but still, I would like to give in something
  instead of nothing at all. Start small grow big, hopefully :)</p>\n<h2>Features</h2>\n<p>Some
  of the main features of the web application are as follows:</p>\n<ul>\n<li>\n<p><strong>Write\\Edit\\Delete
  Quotes if logged in.</strong></p>\n</li>\n<li>\n<p><strong>Like / Unlike a Quote.</strong></p>\n</li>\n<li>\n<p><strong>See
  all of your quotes.</strong></p>\n</li>\n<li>\n<p><strong>Randomized Quotes on Homepage.</strong></p>\n</li>\n<li>\n<p><strong>Dark/Light
  theme based on Browser's Preference and local storage.</strong></p>\n</li>\n<li>\n<p><strong>The
  app is mobile responsive as well, though the navbar is a bit wonky with the light/dark
  mode switch toggle, which can be taken care of soon.</strong></p>\n</li>\n</ul>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1630084573444/YEj38dUTD_.png\"
  alt=\"dqmob.png\" /></p>\n<h2>In the making</h2>\n<blockquote>\n<p>Have the curiosity
  to learn, rest is automated</p>\n</blockquote>\n<p>The project was made by using
  various inspirational articles and videos about making a web application. But the
  idea stuck in my mind when I was thinking about the people who don't get inspired
  as a developer. Like there is no way you can remain sad about being a developer
  and keep on dealing with imposter syndrome. Every developer has a perspective of
  programming but there is an infinite number of opportunities if you are curious
  enough. Just started making the project and got so much into it that I was literally
  dreaming about it like I saw parts of the webpage. In my dream and I am making it
  that was genuinely a thing that powered me to complete it.</p>\n<p>The project roughly
  started on 19th August and almost ended on 26th August, like the actual webpage
  and its core functionalities. Though on 27th were some styling and extra additions
  such as the About section and Footer. That was like the most productive week I ever
  had in my programming journey. That was fun as heck.</p>\n<h2>Tech Stack</h2>\n<p>The
  Tech Stack involved with this app is :</p>\n<ul>\n<li><code>Django</code></li>\n<li><code>PostgreSQL</code></li>\n<li><code>HTML/CSS/JS</code></li>\n<li><code>
  Bootstrap</code></li>\n</ul>\n<p>I have not used any front-end end frameworks just
  because I never found the need to learn them.  I had experience with Django for
  just 2 months and I am surprised I was able to make it. As obvious I have used Auth0
  for authentication in my web application.</p>\n<h3>Auth0 integration for Authentication</h3>\n<p>I
  must tell you using Auth0 was just flawless addition to my app as I have to do almost
  nothing, just drop some credentials of the Auth0 application with my Django project
  using a  <a href=\"https://auth0.com/docs/quickstarts\">well-documented guide</a>
  \ for every type of framework. Simply straight-forward was the name for integrating
  authentication in my app.</p>\n<h4>How I used Auth0 with Django</h4>\n<p>I've used
  Template tags such as if blocks to verify if the user is authenticated or not.</p>\n<pre
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
  class=\"p\">&lt;</span><span class=\"nt\">div</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;login-box auth0-box before&quot;</span><span
  class=\"p\">&gt;</span>\n\t\t{{ &quot;{% if user.is_authenticated &quot;}} %}\n\t\t
  \   <span class=\"p\">&lt;</span><span class=\"nt\">a</span> <span class=\"na\">class</span><span
  class=\"o\">=</span><span class=\"s\">&quot;btn btn-primary btn-sm tn-logout &quot;</span>
  <span class=\"na\">href</span><span class=\"o\">=</span><span class=\"s\">&quot;/logout&quot;</span><span
  class=\"p\">&gt;</span>Log Out<span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span
  class=\"p\">&gt;</span>\n\t\t{{ &quot;{% else &quot;}} %}\n\t\t    <span class=\"p\">&lt;</span><span
  class=\"nt\">a</span> <span class=\"na\">class</span><span class=\"o\">=</span><span
  class=\"s\">&quot;btn btn-primary btn-sm tn-login &quot;</span> <span class=\"na\">href</span><span
  class=\"o\">=</span><span class=\"s\">&quot;/login/auth0&quot;</span><span class=\"p\">&gt;</span>Log
  In<span class=\"p\">&lt;/</span><span class=\"nt\">a</span><span class=\"p\">&gt;</span>\n
  \   {{ &quot;{% endif &quot;}} %}\n<span class=\"p\">&lt;/</span><span class=\"nt\">div</span><span
  class=\"p\">&gt;</span>\n</pre></div>\n\n</pre>\n\n<p>This was just readily available
  on their documentation though there were some adjustments as per the project requirements
  in this code to fit in the place.</p>\n<p>I must say, integrating Auth0 is even
  easier than using Django User Model in some sense as most of the stuff is handled
  by the Auth0, on our side, we simply have to create the Auth0 specific app with
  the credentials from the dashboard rest just works flawlessly till now. How sweet
  and</p>\n<h3>Specifications</h3>\n<p>I won't go in-depth about the technicalities
  of the project but would like to address certain things. Firstly I have mostly used
  Class-based views for the major part, certain areas are still function-based just
  for the simplicity of the application and a few of them are handled and documented
  by Auth0 so just preferred that.</p>\n<p>Another thing is about Models, I just have
  a simple single model called <code>Quote</code> which has an Author as a Foreign
  Key from the Django User Model. I would have also created multiple emojis for the
  like system but I was too excited and in a rush to see the actual app, So just kept
  it simple. XD\n<img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1630060555499/246ZKnypP.png\"
  alt=\"image.png\" /></p>\n<p>The rest of the stuff like <code>URLs</code>, <code>Templates</code>,
  and <code>static files</code> are handled in a neatly organized way as depicted
  in the below diagram.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1630060426600/JHMlrfOKZ.png\"
  alt=\"image.png\" /></p>\n<p>As it might not be clear from that, there are 3 apps
  -&gt; <code>auth0login</code>, <code>quotes</code>, and <code>user</code>, here
  <code>quotes</code> is kind of the most important app as it has the models, forms,
  URLs, and the views linked to them here.</p>\n<h3>Hosting</h3>\n<p>Hosting as you
  can guess, it's on  <a href=\"https://www.heroku.com/\">Heroku</a> , quite beginner-friendly
  and feature-rich. I also have a free addon for PostgreSQL Database here.  It's limited
  to 10K rows but that's sufficient for a starter app like this in my opinion. Also,
  it has 500 free hours of dyno, which is also sufficient for a small low-traffic
  app like this.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1630076036133/9ofxnM5VN.png\"
  alt=\"image.png\" /></p>\n<h2>Bugs Encountered</h2>\n<blockquote>\n<p>Love Bugs,
  they'll give you experience</p>\n</blockquote>\n<p>This is the most exciting and
  important part if you consider the hackathon because this determines the level of
  experience that a developer shoes in achieving certain things or features if you
  will. Faced some typical Django errors as usual but also some unexpected things
  like Dark mode and light mode clashing together due to poorly written media queries
  in CSS.</p>\n<p>As usual, the start is often hard, there is no motivation whatsoever
  in using the admin section to test the database queries and gibberish basic HTML
  page. In that process, I faced some primary key access issues and was able to understand
  the concept more clearly by fixing it.</p>\n<p>Another instance was with handling
  post requests in Django which I've failed to do before. I used forms and a hybrid
  of CSS and bootstrap to style those forms which just works brilliantly. That took
  some time to figure out the exact working but after a while, it was working charms.</p>\n<h2>Future
  Updates</h2>\n<p>As said, I would like to add more like buttons expressing different
  emotions. Some other features to add are:</p>\n<ul>\n<li>To add more emojis like
  claps, cheers, and others.</li>\n<li>To add a profile page in the app that would
  display all the quotes of the particular author and the details related to him/her.</li>\n<li>Adding
  some tags to filter out particular types of quotes.</li>\n<li>Improve UI-UX a bit
  more to make it professional and pleasant.</li>\n</ul>\n<h2>Source Code</h2>\n<blockquote>\n<p>Talk
  is cheap, show me the code - Linus Torvalds</p>\n</blockquote>\n<p>The source code
  is available at GitHub on this  <a href=\"https://github.com/Mr-Destructive/devquotes\">Link</a>.\nIt's
  freely open for any contribution after the hackathon(mid-September).  Some of the
  files such as the environment variables, virtual environments, cached data are not
  uploaded for security and obvious reasons.</p>\n<p>Enough of technical talks, let's
  see the DEMO,</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1630073466386/3wgnST5hc.gif\"
  alt=\"dqss.gif\" /></p>\n<p>Silent claps.......</p>\n<p>It's not a great UI-UX but
  works and is not too shabby in my opinion considering it only has base HTML and
  CSS with a little Bootstrap and Javascript. But ya, a fully functional Backend that's
  what I was looking for a full-stack app with some decent features. Hope it helps
  some developers stay motivated and hack into some hackathons like this.</p>\n<h3>References
  used while creating the app:</h3>\n<ul>\n<li><a href=\"https://www.youtube.com/watch?v=B40bteAMM_M&amp;list=PLCC34OHNcOtr025c1kHSPrnP18YPB-NFi\">Codemy
  -John Elder Django tutorial</a></li>\n<li><a href=\"https://docs.djangoproject.com/en/3.2/topics/forms/\">Django
  - Documentation for Forms</a></li>\n<li><a href=\"https://stackoverflow.com/questions/28837511/django-template-how-to-randomize-order-when-populating-page-with-objects\">Django
  template randomizer shuffle</a></li>\n<li><a href=\"https://www.youtube.com/watch?v=kzN_VCFG9NM\">Auth0
  app Django integration</a></li>\n</ul>\n<h2>Closing Words</h2>\n<blockquote>\n<p>Why
  developers find solutions to bugs at the stroke of sleeping, that's multithreading
  in our brains</p>\n</blockquote>\n<p>Hope you liked the project and hopefully will
  inspire developers to stay motivated and can focus on their goals more than dealing
  with imposter syndrome and whatnot.</p>\n<p>Thank you for reading and using the
  app, for any feedbacks, Twitter handles, comment section, GitHub issues, LinkedIn
  messages are all freely open. Thanks. Happy Coding :)</p>\n"
image_url: https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643286845/blogmedia/bhajgzmqaknzpk1z9qdn.png
long_description: No one can understand the joy in finishing a personal project, except
  the dreamer It was a while,since I have posted an article here, as I was busy on
  a project or a hackathon. Hello, world Every time I start a project something else
  comes and distra
now: 2025-05-01 18:11:33.313781
path: blog/posts/2021-08-27-DevQuotes-Auth0-x-Hashnode.md
prevnext: null
slug: devquotes-platform
status: published
subtitle: A platform created in a Hackathon for developer to get motivated by some
  fun and inspireing quotes.
tags:
- hashnode
- auth0
- django
- web-development
- python
templateKey: blog-post
title: 'Dev Quotes: A platform for developers to quote and get inspired - Auth0 x
  Hashnode Hackathon'
today: 2025-05-01
---

## Introduction

> No one can understand the joy in finishing a personal project, except the dreamer 

It was a while,since I have posted an article here, as I was busy on a project or a hackathon.

Hello, world! I am Meet a student and a self-taught web developer. I like to make and break stuff, especially when it comes to programming and Linux. I like shell scripting and learning different languages at once, love to learn about Vim and Linux everyday.

Every time I start a project something else comes and distracts me let that be any other programming language or technology. That leads to creating new projects and leaving the one behind unfinished, I know most of the developers face this.  But this time, thanks to Auth0 X Hashnode Hackathon, I was able to create an almost finished project within almost 10 days. Having a deadline and competition creates a mindset to finish a project on time, that's my first takeaway from this Hackathon. OH! this is my first Hackathon by the way, and it has been amazing so far.  

** Applying a framework to do something you desire and then everything working smoothly (after fixing 100s of bugs) is such a great feeling that no one can understand except for the person who just dreamt of it. **

I'll like to share my project which is a web application for the Auth0 x Hashnode Hackathon. Here it goes.

## What is Dev Quotes?

Dev quotes is a web app designed for publishing and viewing quotes related to programming, developer mindset, and all the technicalities involved in a developer's life. It's basically a medium to express the life of developers and get inspired by others.  Here it is [devquotes](https://devquotess.herokuapp.com/)

#### Dark Mode:
![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1630074051548/TQz9Koh7l.png)

#### Light Mode:
![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1630078314355/VhfLrcjJa.png)


## Why Dev Quotes?
> Developers are not the people who only understand how to write code but they're also the people who can make the code understandable

As a developer, there are often times where you have no motivation left inside, but you never know you might be just a few lines of code away from making another project or fixing a bug. For that, we require some inspiration or a push to break the barrier of.  I am not saying it's just for developers, it's designed for developers but everyone is open to understanding the developers' lives and their struggles. 

I also felt the need to give back some love-crafted web app to the ever-wonderful and supportive dev community. It's a small application but still, I would like to give in something instead of nothing at all. Start small grow big, hopefully :)

## Features

Some of the main features of the web application are as follows:

- **Write\Edit\Delete Quotes if logged in.**

- **Like / Unlike a Quote.**

- **See all of your quotes.**

- **Randomized Quotes on Homepage.**

- **Dark/Light theme based on Browser's Preference and local storage.**
 
- **The app is mobile responsive as well, though the navbar is a bit wonky with the light/dark mode switch toggle, which can be taken care of soon.**



![dqmob.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1630084573444/YEj38dUTD_.png)



## In the making

> Have the curiosity to learn, rest is automated

The project was made by using various inspirational articles and videos about making a web application. But the idea stuck in my mind when I was thinking about the people who don't get inspired as a developer. Like there is no way you can remain sad about being a developer and keep on dealing with imposter syndrome. Every developer has a perspective of programming but there is an infinite number of opportunities if you are curious enough. Just started making the project and got so much into it that I was literally dreaming about it like I saw parts of the webpage. In my dream and I am making it that was genuinely a thing that powered me to complete it. 

The project roughly started on 19th August and almost ended on 26th August, like the actual webpage and its core functionalities. Though on 27th were some styling and extra additions such as the About section and Footer. That was like the most productive week I ever had in my programming journey. That was fun as heck.

## Tech Stack

The Tech Stack involved with this app is :
- `Django`
- `PostgreSQL` 
- `HTML/CSS/JS`
- ` Bootstrap`

 I have not used any front-end end frameworks just because I never found the need to learn them.  I had experience with Django for just 2 months and I am surprised I was able to make it. As obvious I have used Auth0 for authentication in my web application.

### Auth0 integration for Authentication

I must tell you using Auth0 was just flawless addition to my app as I have to do almost nothing, just drop some credentials of the Auth0 application with my Django project using a  [well-documented guide](https://auth0.com/docs/quickstarts)  for every type of framework. Simply straight-forward was the name for integrating authentication in my app.

#### How I used Auth0 with Django

I've used Template tags such as if blocks to verify if the user is authenticated or not. 
```html
<div class="login-box auth0-box before">
		{{ "{% if user.is_authenticated "}} %}
		    <a class="btn btn-primary btn-sm tn-logout " href="/logout">Log Out</a>
		{{ "{% else "}} %}
		    <a class="btn btn-primary btn-sm tn-login " href="/login/auth0">Log In</a>
    {{ "{% endif "}} %}
</div>
```

This was just readily available on their documentation though there were some adjustments as per the project requirements in this code to fit in the place.

I must say, integrating Auth0 is even easier than using Django User Model in some sense as most of the stuff is handled by the Auth0, on our side, we simply have to create the Auth0 specific app with the credentials from the dashboard rest just works flawlessly till now. How sweet and 
 
### Specifications

I won't go in-depth about the technicalities of the project but would like to address certain things. Firstly I have mostly used Class-based views for the major part, certain areas are still function-based just for the simplicity of the application and a few of them are handled and documented by Auth0 so just preferred that. 

Another thing is about Models, I just have a simple single model called `Quote` which has an Author as a Foreign Key from the Django User Model. I would have also created multiple emojis for the like system but I was too excited and in a rush to see the actual app, So just kept it simple. XD
![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1630060555499/246ZKnypP.png) 

The rest of the stuff like `URLs`, `Templates`, and `static files` are handled in a neatly organized way as depicted in the below diagram.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1630060426600/JHMlrfOKZ.png)

As it might not be clear from that, there are 3 apps -> `auth0login`, `quotes`, and `user`, here `quotes` is kind of the most important app as it has the models, forms, URLs, and the views linked to them here. 

### Hosting

Hosting as you can guess, it's on  [Heroku](https://www.heroku.com/) , quite beginner-friendly and feature-rich. I also have a free addon for PostgreSQL Database here.  It's limited to 10K rows but that's sufficient for a starter app like this in my opinion. Also, it has 500 free hours of dyno, which is also sufficient for a small low-traffic app like this.  

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1630076036133/9ofxnM5VN.png)

## Bugs Encountered 
> Love Bugs, they'll give you experience

This is the most exciting and important part if you consider the hackathon because this determines the level of experience that a developer shoes in achieving certain things or features if you will. Faced some typical Django errors as usual but also some unexpected things like Dark mode and light mode clashing together due to poorly written media queries in CSS.

 As usual, the start is often hard, there is no motivation whatsoever in using the admin section to test the database queries and gibberish basic HTML page. In that process, I faced some primary key access issues and was able to understand the concept more clearly by fixing it.

Another instance was with handling post requests in Django which I've failed to do before. I used forms and a hybrid of CSS and bootstrap to style those forms which just works brilliantly. That took some time to figure out the exact working but after a while, it was working charms. 


## Future Updates

As said, I would like to add more like buttons expressing different emotions. Some other features to add are:

- To add more emojis like claps, cheers, and others.
- To add a profile page in the app that would display all the quotes of the particular author and the details related to him/her.  
- Adding some tags to filter out particular types of quotes.
- Improve UI-UX a bit more to make it professional and pleasant.

## Source Code

>Talk is cheap, show me the code - Linus Torvalds

The source code is available at GitHub on this  [Link](https://github.com/Mr-Destructive/devquotes). 
It's freely open for any contribution after the hackathon(mid-September).  Some of the files such as the environment variables, virtual environments, cached data are not uploaded for security and obvious reasons.

Enough of technical talks, let's see the DEMO,


![dqss.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1630073466386/3wgnST5hc.gif)

Silent claps.......

It's not a great UI-UX but works and is not too shabby in my opinion considering it only has base HTML and CSS with a little Bootstrap and Javascript. But ya, a fully functional Backend that's what I was looking for a full-stack app with some decent features. Hope it helps some developers stay motivated and hack into some hackathons like this.

### References used while creating the app:
- [Codemy -John Elder Django tutorial]( https://www.youtube.com/watch?v=B40bteAMM_M&list=PLCC34OHNcOtr025c1kHSPrnP18YPB-NFi)
- [Django - Documentation for Forms](https://docs.djangoproject.com/en/3.2/topics/forms/)
- [Django template randomizer shuffle](https://stackoverflow.com/questions/28837511/django-template-how-to-randomize-order-when-populating-page-with-objects)
- [Auth0 app Django integration](https://www.youtube.com/watch?v=kzN_VCFG9NM)


## Closing Words

> Why developers find solutions to bugs at the stroke of sleeping, that's multithreading in our brains 

Hope you liked the project and hopefully will inspire developers to stay motivated and can focus on their goals more than dealing with imposter syndrome and whatnot. 

Thank you for reading and using the app, for any feedbacks, Twitter handles, comment section, GitHub issues, LinkedIn messages are all freely open. Thanks. Happy Coding :)