---
cover: ''
date: 2021-10-24
datetime: 2021-10-24 00:00:00+00:00
description: We all have seen the buzz around web frameworks like Django, Flask, Node.js,
  etc but have you taken time to learn all of them? No, and you shouldn Before we
  get
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-10-24-Flask-Django.md
image_url: https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643288439/blogmedia/p8smflzd3rjb6scowhmc.png
long_description: We all have seen the buzz around web frameworks like Django, Flask,
  Node.js, etc but have you taken time to learn all of them? No, and you shouldn Before
  we get into frameworks let us understand the key components of the web application
  Every web fra
now: 2022-07-24 13:37:38.456123
path: blog/posts/2021-10-24-Flask-Django.md
slug: flask-django-frameworks
status: published
subtitle: Understanding the concept of backend web frameworks like Django and Flask.
tags:
- django
- flask
- python
- web-development
templateKey: blog-post
title: Flask and Django - the backend web frameworks
today: 2022-07-24
---

## Introduction

We all have seen the buzz around web frameworks like Django, Flask, Node.js, etc but have you taken time to learn all of them? No, and you shouldn't! Because many web frameworks share the same principle and workflow with a bit of difference. It's just like learning one programming language and applying the same concepts in a different syntax and mechanism. In the world of web frameworks, this is the case as well, but most of them will disagree with it as every web framework is unique in its design and that's true, don't get me wrong.

Before we get into frameworks let us understand the key components of the web application
- **Database** - It holds the data for our application.
- **Server** - Used to fetch/store/manage requests from the client. 
- **API** - Used as an interface between the client and the Database. 
- **Client** - The browser or any client that requests for resources.

![djflask-webapp.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1635081505223/rLnSyA_7Y.png)

Every web framework will serve the same purpose with different design, architecture, language but it will have a similar pattern in developing the application. Let's clear the concepts in this article.

## What is a back-end Web framework?

A web framework is a tool/application meant for designing, creating, testing web applications a lot quicker with a lot of ease. Without web frameworks, you will have been writing some code that will tire you very quickly. 

It even seems impossible to manually write markups for each piece of data in the application, which is taken care of by dynamic templating in Python-based frameworks like Django, Flask. The database queries are managed by the web frameworks as well, otherwise, you will have been writing SQL queries manually! How painful and frustrating that would look, of course, you can create scripts for querying to the database but you are then creating a component of a framework. `Don't waste time` that's a takeaway from the philosophy of all the web frameworks.

Another thing that back-end web frameworks do is create homogeneity in development across different environments and applications. It also creates a developer-friendly environment. We must not forget how easy and quick applications can be built using the back-end web frameworks. 



### A back-end Web framework provides some of the features like:

- Handle web requests
- Manage DB by just using some simple scripts
- Render Dynamic Templates
- Provide a lot of native-language libraries integration
- Organize a project much easily and effectively
- Options to scale the application at any level
- Provide some standard and secure way to run the server(production)
- Design APIs much easily


Let us look at two of the most popular frameworks in the Python community.

### 1. Flask
### 2. Django

![backend framework ranking](https://cdn.hashnode.com/res/hashnode/image/upload/v1635070666410/JbMc7NKP0.png)

We are seeing that Django and Flask are among the top 3 back-end web frameworks in 2021. So there is no double thought on why you should learn these technologies.

### What is Flask

Flask is the bare-bones framework that provides a lot of customizability with a lot less boilerplate code. It is a framework that provides a lot of third-party libraries to add functionalities to our application. 

> Flask is a micro web framework

Flask as per the official documentation is a `micro` framework indicating it has a very minimal setup. It is a back-end web framework that can be structured as per needs with a very little configuration overhead. That being said, it can get a bit limited in structuring and functionalities as it needs to taken care of manually.

**Flask is the easiest back-end web framework to get started and learn the fundamentals of server-side**. Flask is quite flexible in terms of scalability and maintenance of decent-sized applications as well. Though the community is not that big and absence of standardization in Flask, it is a go-to back-end web framework for beginners as well as experts due to its simplicity and flawless integration with Python libraries. 

The main concepts in Flask might be:

- Virtual Environment 
- WSGI as a web server
- App routing
- Jinga2 as a templating language
- Creating Database connections

So, **Flask is kind of a DIY back-end web framework with rich sets of libraries and customizability out of the box**. This can easily be a beginner's choice and a right one too.

### What is Django

Django is also a back-end web framework based on Python programming language but it is more standardized and high-level. Django encourages a defined pattern for development but with customization and freedom in mind.

Django also modularizes the components into so-called `apps` to provide a scalable experience. It has a lot of boilerplate code to get up and running quite easily, it also has a `Admin section` pre-built with all the functionalities. Similar to `Flask`, it also provides flawless integration with all the Python libraries. It provides a much easier Database integration and pre-built `User` authentication along with its model ready to plug in and use. 

> Django is a Batteries included Framework

That means it has baked in functionalities like User-Authentication, Admin Section, Database Integration, RSS/Atom syndication feeds, etc. 

![djflask-dj.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1635079576954/WcjcokoiX.png)

The main concepts in Django include:

- Virtual Environment
- WSGI/ASGI as web servers
- Project structure
- `Model View Controller` Architecture in developing apps
- Django Templating Language for rendering Dynamic Templates 
- `Object-Relational Mapping` in creating the applications

Unlike Flask, Django is already baked in with a lot of functionalities and integration with a ton of features. It should be good for beginners but many things are already taken care of that can be a huddle in **actual learning process**, that being said it is a much scalable and production-ready web framework (not only back-end).

### What are the similarities between them?

Well, if you learn one the other will be quite easy enough to pick up. The overall development is almost similar but unique in its own way. 

- Pythonic syntax and libraries
- Project Structure is quite similar to `blueprints` in Flask and `apps` in Django
- Templating Language is almost similar
- Static Files are handled similarly with a different syntax 
- URL Routing is the same as it binds the view(functions) with a pattern
- Ease in Deployment with minimal configuration

## What should you learn?

That question is dependent on the type of application you are trying to make but for a beginner trying to get hands dirty on the server-side, I would recommend `Flask` as it is quite minimal and helps in constructing the base for the concepts like APIs, Databases, Requests, Admin section, etc. 

This might not be that difficult for people trying to learn back-end from scratch but for people with a bit of programming and server-side experience, `Django` should be a go-to framework for all their needs. 

At the end of the day, it hardly matters what you do with which framework, what people see is the end result.

## Conclusion

Thus, from this article, you might have got a bit understanding of why are the frameworks used in making applications and also the similarities and differences in the Python-based back-end web frameworks like Django and Flask. If you have any thoughts please let me know in the comments or on my social handles, any kind of feedback is much appreciated. 

Thank you for reading till here, until then as always Happy Coding :)
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
    
    <a class='prev' href='/golang-input'>
    
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M13.5 8.25L9.75 12L13.5 15.75" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"> </path>
        </svg>
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>prev</p>
            <p class='prevnext-title'>Golang: Input</p>
        </div>
    </a>
    
    <a class='next' href='/golang-conditionals-loops'>
    
        <div class='prevnext-text'>
            <p class='prevnext-subtitle'>next</p>
            <p class='prevnext-title'>Golang: Conditionals and Loops</p>
        </div>
        <svg width="50px" height="50px" viewbox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M10.5 15.75L14.25 12L10.5 8.25" stroke="var(--prevnext-color-angle)" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"></path>
        </svg>
    </a>
  </div>