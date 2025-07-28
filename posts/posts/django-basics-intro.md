{
  "title": "Django Basics: What is it?",
  "status": "published",
  "slug": "django-basics-intro",
  "date": "2025-04-05T12:33:50Z"
}

<h2>Introduction</h2>
<p>Welcome to Django Basics series, in this series we'll explore the basics of the Django web framework. In this part, we'll understand what this web framework provides and what actually the back-end development consists of. We'll discuss where Django is used and why it is a great choice for beginners as well as experienced developers.</p>
<h2>What is Django?</h2>
<p><a href="https://developer.mozilla.org/en-US/docs/Learn/Server-side/Django/Introduction">Django</a> is a back-end web framework. It is based on python which means you have to write most of the project's code in Python. But Django comes with a lot of boilerplate code and thus it becomes quite quick in the development.</p>
<p>Django is an open-source framework, it is maintained by the Django Software Foundation Organization. You can view the source code at <a href="https://github.com/django/django">GitHub</a>.</p>
<h3>BACKEND ?</h3>
<p>The term <code>backend</code> refers to the section or an essential component in Web development, it consists of a <code>database</code>, <code>API</code>, and the <code>web server</code> itself which allows the components to connect together. There might be other components like <code>load-balancers</code>, <code>middleware</code>, etc. But the core of web applications revolves around <strong>Databases</strong> and <strong>API</strong>.</p>
<h4>Database</h4>
<p>A database is a technology or tool that lets you store the data which might be used for serving the actual application, that might be a frontend app, standalone API, etc. The data you want to store might be generally the User Accounts, Content of the App, basically any some form of data(there are exceptions here, you can't directly store media files in DB). The Database allows to make content management and the application dynamic and can be personalized. We have certain types of databases like SQL(relational), NO-SQL, Cloud, Network, etc. The tools of these database management are PostgreSQL, MySQL, MongoDB, HarperDB,etc. These tools allow you to manage your database in a convenient way.</p>
<h4>API</h4>
<p>An API or Application Programming Interface is a way for any frontend app, outside the system to access the database. API allows you to query to the database with GET, POST, DELETE, PUT, etc kinds of operation/requests to the database via the webserver. In API, we have endpoints or (URL routes) at which a particular designated operation can be performed. In APIs, we currently have four primary architectures namely RESTful (quite famous and well established), SOAP, gRPC, and GRAPHQL (new and increasing in popularity).</p>
<h3>Framework?</h3>
<p>A framework is a tool to do a certain task efficiently and avoid some repetitive patterns by abstracting many layers in developing it. Django is a high-level framework which means it abstracts certain processes in making the application. It is ideal for beginners to get up and running with a professional full-stack web application(though it requires some learning).</p>
<p>Django makes the project ideal for experienced as well as beginner web developers. The community and the ecosystem of Python are quite amazing as well as there are a ton of resources to get you through your projects.</p>
<p><img src="https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/dj-1.png" alt=""></p>
<p>The above is a high-level view of how Django project development works, the application might be not only one but several other standalone applications working together to make one project in Django. There is a lot of abstraction in Django like the Middleware, Session Management, Security, etc. This should be a good overview of the development map in Django.</p>
<p>Django follows an MVT architecture. Architecture is a standard in developing an application/project for the ease of the workflow and making it an even experience.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1635079576954/WcjcokoiX.png" alt=""></p>
<p>The above diagram depicts the architecture in Django, the components in the Django server include the Model, View, and Template.</p>
<h3>Model</h3>
<p><code>Model</code> refers to the design of the database or a blueprint of the data that is bound with the application in the project.</p>
<h3>View</h3>
<p>The <code>View</code> is the part to control the way the data should be presented or the how response should be given back to a request from the server(client)</p>
<h3>Template</h3>
<p>The <code>Template</code> is the markup or the form of document that is to be rendered on the client-side and these are controlled by the views and parsed with the data from the models.</p>
<h2>Why would you need it?</h2>
<p>As a developer you would find a need to host your projects on the internet, for that learning and deploying a web server from the ground up might be quite complex and time-consuming, Django solves this problem quite well. Not only it is easy but even scalable at a production level, making it quite a robust choice for anyone. And as a bonus thing, it is based on Python, which makes it even easier to write code for people staying at an abstracted perspective in programming. Python has by far the richest sets of libraries and utilities for any domain, this integration with Django is a deadly combination.</p>
<h4>Batteries included?</h4>
<p>Django solves many problems by abstracting away many things like managing the database, rendering dynamic templates(HTML), properly structuring and serving static and media files, well-organized project structure, and many other things. You just have to get the actual thing done i.e. the server logic(or how to design the API/Database models). On top of that, Django has a built-in fully fledged Admin section and a User model. An Admin section is where you can manage the project in a better way without touching the code. It also has certain applications/libraries to make the development of APIs, integrating various databases, forms for posting data, support for Bootstrap a lot easier. It's like a <code>plug and play</code> kind of thing for the development of web applications.</p>
<p>Hence, it is rightly called the <code>Batteries Included</code> web framework.</p>
<h3>Key features of Django</h3>
<ul>
<li>Ease in integrating a database</li>
<li>Flawless Django Template Engine</li>
<li>Easy to scale up/down</li>
<li>Python libraries support out of the box</li>
<li>Amazing Documentation / Helpful community</li>
<li>Developing Production-ready projects quickly</li>
<li>Baked in support for testing, APIs, cookies, sessions, etc</li>
<li>Optimized for security, SEO, and DRY(don't repeat yourself) principles</li>
</ul>
<h2>Applications built with Django</h2>
<p>Django is used in quite a famous application that you might be using daily.</p>
<p>Django along with Python powers the top applications on the internet like:</p>
<ol>
<li>YouTube</li>
<li>Instagram</li>
<li>Spotify</li>
<li>Disqus</li>
<li>Dropbox</li>
<li>Pinterest</li>
<li>National Geographic</li>
<li>Mozilla</li>
<li>BitBucket</li>
<li>Discovery Network</li>
</ol>
<p>You have to say, it is powerful and has firm grounds in the tech industry. It's highly unlikely that Django will be overtaken by another framework at least some years from now.</p>
<blockquote>
<p>Django is a tool to build web applications fast and in a scalable and Pythonic way</p>
</blockquote>
<h2>What will this series cover?</h2>
<p>Learning Django from the ground up. We will learn the setup, folder structure, architecture of Django, What are apps, views, URLs, models, serializers, static and template files, and there is a ton of stuff to be covered.</p>
<h3>Resources to learn Django</h3>
<ul>
<li><a href="https://www.djangoproject.com/start/">Django Official Docs</a></li>
<li><a href="https://www.youtube.com/c/veryacademy/playlists?view=50&amp;sort=dd&amp;shelf_id=2">Very Academy - Django Playlist</a></li>
<li><a href="https://www.youtube.com/watch?v=HHx3tTQWUx0&amp;list=PLCC34OHNcOtqW9BJmgQPPzUpJ8hl49AGy">Codemy.com - Django</a></li>
<li><a href="https://www.youtube.com/watch?v=UmljXZIypDc&amp;list=PL-osiE80TeTtoQCKZ03TU5fNfx2UY6U4p">Corey Schafer</a></li>
<li><a href="https://www.youtube.com/watch?v=SIyxjRJ8VNY&amp;list=PLsyeobzWxl7r2ukVgTqIQcl-1T0C2mzau">Telusko</a></li>
</ul>
<h2>Conclusion</h2>
<p>From this article, we were able to understand the Django framework, what is it, and why it should be used on a high level. Further, we explored the web application(backend) components which are targeted by Django for ease of developing applications. We also saw the baseline architecture that Django uses to make projects.</p>
<p>In the next section, we'll start the actual coding in Django, firstly how to set up the environment and understanding the folder structure, and so on. So I hoped you enjoyed the article. Thank you for reading. Happy Coding :)</p>
