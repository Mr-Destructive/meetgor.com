{
  "type": "posts",
  "title": "Django Basics: What is it?",
  "subtitle": "Understanding what and why of the Django framework",
  "date": "2021-11-16 20:45:00 +0530",
  "status": "published",
  "tags": [
    "django",
    "python",
    "web-development"
  ],
  "slug": "django-basics-intro",
  "image_url": "https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643289206/blogmedia/gbq3rlfat3erbxocn7yn.png",
  "series": [
    "Django-Basics",
    "Django-Series"
  ],
  "series_description": "Understand the fundamentals of backend web development with Django web framework."
}


## Introduction

Welcome to Django Basics series, in this series we'll explore the basics of the Django web framework. In this part, we'll understand what this web framework provides and what actually the back-end development consists of. We'll discuss where Django is used and why it is a great choice for beginners as well as experienced developers. 

## What is Django?

[Django](https://developer.mozilla.org/en-US/docs/Learn/Server-side/Django/Introduction) is a back-end web framework. It is based on python which means you have to write most of the project's code in Python. But Django comes with a lot of boilerplate code and thus it becomes quite quick in the development. 

Django is an open-source framework, it is maintained by the Django Software Foundation Organization. You can view the source code at [GitHub](https://github.com/django/django).

### BACKEND ?

The term `backend` refers to the section or an essential component in Web development, it consists of a `database`, `API`, and the `web server` itself which allows the components to connect together. There might be other components like `load-balancers`, `middleware`, etc. But the core of web applications revolves around **Databases** and **API**. 

#### Database

A database is a technology or tool that lets you store the data which might be used for serving the actual application, that might be a frontend app, standalone API, etc. The data you want to store might be generally the User Accounts, Content of the App, basically any some form of data(there are exceptions here, you can't directly store media files in DB). The Database allows to make content management and the application dynamic and can be personalized. We have certain types of databases like SQL(relational), NO-SQL, Cloud, Network, etc. The tools of these database management are PostgreSQL, MySQL, MongoDB, HarperDB,etc. These tools allow you to manage your database in a convenient way.  

#### API

An API or Application Programming Interface is a way for any frontend app, outside the system to access the database. API allows you to query to the database with GET, POST, DELETE, PUT, etc kinds of operation/requests to the database via the webserver. In API, we have endpoints or (URL routes) at which a particular designated operation can be performed. In APIs, we currently have four primary architectures namely RESTful (quite famous and well established), SOAP, gRPC, and GRAPHQL (new and increasing in popularity). 

### Framework?

A framework is a tool to do a certain task efficiently and avoid some repetitive patterns by abstracting many layers in developing it. Django is a high-level framework which means it abstracts certain processes in making the application. It is ideal for beginners to get up and running with a professional full-stack web application(though it requires some learning).

Django makes the project ideal for experienced as well as beginner web developers. The community and the ecosystem of Python are quite amazing as well as there are a ton of resources to get you through your projects.  

![](https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/dj-1.png)

The above is a high-level view of how Django project development works, the application might be not only one but several other standalone applications working together to make one project in Django. There is a lot of abstraction in Django like the Middleware, Session Management, Security, etc. This should be a good overview of the development map in Django.

Django follows an MVT architecture. Architecture is a standard in developing an application/project for the ease of the workflow and making it an even experience. 

![](https://cdn.hashnode.com/res/hashnode/image/upload/v1635079576954/WcjcokoiX.png)

The above diagram depicts the architecture in Django, the components in the Django server include the Model, View, and Template.

### Model

`Model` refers to the design of the database or a blueprint of the data that is bound with the application in the project. 

### View

The `View` is the part to control the way the data should be presented or the how response should be given back to a request from the server(client)

### Template

The `Template` is the markup or the form of document that is to be rendered on the client-side and these are controlled by the views and parsed with the data from the models.

## Why would you need it?

As a developer you would find a need to host your projects on the internet, for that learning and deploying a web server from the ground up might be quite complex and time-consuming, Django solves this problem quite well. Not only it is easy but even scalable at a production level, making it quite a robust choice for anyone. And as a bonus thing, it is based on Python, which makes it even easier to write code for people staying at an abstracted perspective in programming. Python has by far the richest sets of libraries and utilities for any domain, this integration with Django is a deadly combination. 

#### Batteries included?

Django solves many problems by abstracting away many things like managing the database, rendering dynamic templates(HTML), properly structuring and serving static and media files, well-organized project structure, and many other things. You just have to get the actual thing done i.e. the server logic(or how to design the API/Database models). On top of that, Django has a built-in fully fledged Admin section and a User model. An Admin section is where you can manage the project in a better way without touching the code. It also has certain applications/libraries to make the development of APIs, integrating various databases, forms for posting data, support for Bootstrap a lot easier. It's like a `plug and play` kind of thing for the development of web applications. 

Hence, it is rightly called the `Batteries Included` web framework.

### Key features of Django

- Ease in integrating a database
- Flawless Django Template Engine
- Easy to scale up/down
- Python libraries support out of the box
- Amazing Documentation / Helpful community
- Developing Production-ready projects quickly
- Baked in support for testing, APIs, cookies, sessions, etc
- Optimized for security, SEO, and DRY(don't repeat yourself) principles

## Applications built with Django

Django is used in quite a famous application that you might be using daily. 

Django along with Python powers the top applications on the internet like:

1. YouTube
2. Instagram
3. Spotify
4. Disqus
5. Dropbox
6. Pinterest
7. National Geographic
8. Mozilla
9. BitBucket
10. Discovery Network

You have to say, it is powerful and has firm grounds in the tech industry. It's highly unlikely that Django will be overtaken by another framework at least some years from now.  

> Django is a tool to build web applications fast and in a scalable and Pythonic way

## What will this series cover?

Learning Django from the ground up. We will learn the setup, folder structure, architecture of Django, What are apps, views, URLs, models, serializers, static and template files, and there is a ton of stuff to be covered. 

### Resources to learn Django

- [Django Official Docs](https://www.djangoproject.com/start/)
- [Very Academy - Django Playlist](https://www.youtube.com/c/veryacademy/playlists?view=50&sort=dd&shelf_id=2)
- [Codemy.com - Django](https://www.youtube.com/watch?v=HHx3tTQWUx0&list=PLCC34OHNcOtqW9BJmgQPPzUpJ8hl49AGy)
- [Corey Schafer](https://www.youtube.com/watch?v=UmljXZIypDc&list=PL-osiE80TeTtoQCKZ03TU5fNfx2UY6U4p)
- [Telusko](https://www.youtube.com/watch?v=SIyxjRJ8VNY&list=PLsyeobzWxl7r2ukVgTqIQcl-1T0C2mzau)

## Conclusion

From this article, we were able to understand the Django framework, what is it, and why it should be used on a high level. Further, we explored the web application(backend) components which are targeted by Django for ease of developing applications. We also saw the baseline architecture that Django uses to make projects. 

In the next section, we'll start the actual coding in Django, firstly how to set up the environment and understanding the folder structure, and so on. So I hoped you enjoyed the article. Thank you for reading. Happy Coding :)
