---
templateKey: blog-post
title : "Django Basics: Folder Structure"
subtitle: "Understanding the folder structure of Django Projects and Applications"
date: 2021-11-24 15:45:00 +0530
status: published
tags: ['django', 'python', 'web-development']
slug: django-basics-folder-struct
image_url: https://res.cloudinary.com/dgpxbrwoz/image/upload/v1637745125/blogmedia/4_gnddxj.png
series: ['Django-Basics','Django-Series']
---

### Introduction

After setting up the development for the Django framework, we will explore the project structure. In this part, we understand the structure along with the various components in the Project as well as individual apps. We will understand the objective of each file and folder in a brief and hopefully by the end of this part, you'll be aware of how the Django project is structured and get a good overview of the flow of development in the Django project.

## Project Structure

We will create a Django project from scratch and understand it from the ground up. As in the previous part, I've shown you how to create a project. In this section, we'll create a project `Blog`. TO do that, we'll create a folder called `Blog`, install and set up the virtual environment as discussed and explained in the previous part.

After the virtual environment is created and activated, we'll create the project.

```shell
django-admin startproject Blog .
```
After this command, if we see the directory structure, it should look something like this:

![Basic Django folder structure](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1637661221/blogmedia/tree_ak3mgm.png)

As we can see there are 6 files and a folder. The base folder is for the configuration at a project level. I have actually not shown the `venv` ( using `-I venv` option on tree command) as it is out of the scope of this series. The `venv` folder contains modules and scripts which are installed in the virtual environment. 

So, lets break the folder structure down into understandable files.
![](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1637745146/blogmedia/16_qenomh.png)
### manage.py

Our project consists of a `manage.py` file which is to execute several commands at a project level. We do not have to edit any of the contents of this file (never). It is the file that allows us to run the server, apply migrations, create an Admin account, create apps, and do a lot of crucial things with the help of python. 

So, it's just the command-line utility that helps us interact with the Django project and applications for configurations. 

### Project Folder

Now, this is the folder where the project-configuration files are located. **The name of the folder is the same as that of the project**. This makes the folder unique and hence creates a standard way to store files in a structured way. 

The folder is a python package which is indicated by the `__init__.py` file. The purpose of the `__init__.py` file is to tell the Python environment that the current folder is a Python package. 

The folder consist of several files(5 files):

### settings.py

This is a really important file from the project's point of view. This contains certain configurations that can be applied to the rest of the project (or all the apps). 

In the `settings.py` file, we can do some of the following operations :

- List of `applications` that might be pre-installed or user-defined.
- Configure the Middleware. 
- Configure and connect the Database.
- Configure Templates/Static/Media files.
- Custom Configuration for Time-Zones, Emails, Authentication, CORS, etc.

Besides the above-mentioned options, there is a lot of project-specific configurations or application-specific settings as well. 

Here, you'll have a question,

### WHAT IS AN APPLICATION?

An application is a component of a project. There are also Python packages that are made to be used as a Django app that allows reusing the components. But when we are developing the project, we can break a complex process/project into individual apps. 

For Example, a project of `Blogging Platform` might have an application for `posts`, `users`, `api`, `homepage`, etc. So the project `Blogging Platform` might have separated the components like its API, Homepage, Post, Users, and so on to keep the development independent and well organized.

So, we can understand apps as separate components of a large project. We can also understand apps as reusable components, you can use the `posts` app in another project or in a particular app of the same project making it easier and faster to create the project.

### urls.py 

This is a file for managing the `URL` routes of the project. We'll discuss URLs and Views in their own part in the series. This file basically has a list of URLs that should be paired with a `view` or any other function. In the project folder, the URL patterns mostly link a baseurl to the URL file of the particular application. Don't worry if you can't get some of the terms, you'll clearly understand when we see them in the future parts of this series.

### wsgi.py

WSGI or Web Server Gateway Interface is a file that is used to configure the project for production or deployment. This takes care of the server when we deploy into production. It is a Synchronous Web Server i.e. it listens to only one request and responds to that at a time.

Some of the common WSGI servers are [Gunicorn](https://gunicorn.org/), [Apache](https://docs.djangoproject.com/en/3.2/howto/deployment/wsgi/modwsgi/), [uWSGI](https://docs.djangoproject.com/en/3.2/howto/deployment/wsgi/uwsgi/), [cherrypy](https://docs.cherrypy.dev/), [Aspen](https://github.com/buchuki/aspen/blob/master/aspen/wsgi.py), etc.

### asgi.py

ASGI or Asynchronous Server Gateway Interface is also similar to the WSGI file but it serves as an asynchronous web server. This file handles the requests which might be asynchronous i.e. the web server can respond to multiple requests and respond to them at a time. We can even send tasks to the background using this type of server configuration. 

Some of the common ASGI servers are [Uvicorn](https://www.uvicorn.org/), [Daphne](https://docs.djangoproject.com/en/3.2/howto/deployment/asgi/daphne/), [Hypercorn](https://docs.djangoproject.com/en/3.2/howto/deployment/asgi/hypercorn/), etc.

## Creating a Django Project Application 

So, let's create an application to see the structure of the basic app in Django. To create an app, we can use the `startapp` option with the `python manage.py` command followed by the name of the app like:

```shell
python manage.py startapp name
```

Here, `name` can be any app name you'd like to give. 

## Application Structure

After creating an app, let the name be anything it should have a similar structure as :

![](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1637731715/blogmedia/apptree_zr12s0.png)

As we can see there are a couple of things to be discussed here. The main components that we are going to work on within developing the application in the project are: `models.py`, `views.py`, `test.py`. There are other files that we will create manually like the `urls.py`, `serializers.py`, etc.

You also need to add the name of the app in quotes in the `INSTALLED_APPS` list in the `settings.py` file. Something like this:

![](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1637748258/blogmedia/installed_apps_ozir1p.png)

The application files can be summarized as :

![App structure summary](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1637745151/blogmedia/17_xyefpq.png)

Let us look at each of the files and folders in the application:

### models.py

As the same suggests, we need to define the model of a database here. The actual structure and the relationship are created with the help of python and Django in this file. This is the place where the crust of the web application might be defined. 

There are various aspects in creating a model like `Fields`, `Relationship`, `Meta-data`, `methods`, etc. These are defined with the help of python along with the Django Models. In most cases, a model is like a single `table` in an actual database. 

The file is quite important and interesting as it abstracts away the manual work of typing `SQL` queries to create the database. 

### migrations

This migrations folder is a way for Django to keep track of the changes in the database. At every `migration` or actual query that runs to create the table or the database structure. There might be multiple steps or iteration of the database, this folder stores those pieces of information.

To make an analogy, it is like a `.git` folder but for keeping track of the migrations or changes to the database. 

### admin.py

This is the file for performing the operations at the admin level. We generally use this file to register the models into the `Admin section` without touching any frontend part. It provides a built-in `CRUD`(Create Read Update Delete) functionality to the model. This is really good for testing up the model manually before putting effort into the frontend part. 

Other than this, we can customize the admin section with this file. We will see the details in the part of the `Admin section` in this series.

### views.py

This is a file, that acts as a controller/server logic for the Django framework. We can define functions and classes as a response to the incoming requests from the server via the `urls.py` file. There are a couple of approaches when it comes to writing the format of the functions like `class-based views`, `function-based views`, and others depending on the type of operation is been done.

As said earlier, it is the `V`(View) in the `MVT` architecture in Django Framework. This is the place where we write the logic from the server-side to let's say render HTML pages(templates), query to the database with CRUD operations, return an HTTP response, etc.  

### urls.py

This is the file in which a list of URL patterns is mapped to the particular view function. This `urls.py` is specific to the app and it might be prefixed with the URL route mentioned in the project folder's `urls.py` file.  

So, not going much deeper, simply to put it's a map of a particular URL path with a function associated to it which gets triggered(called) when the user visits the URL. 

### tests.py

This is a file where we can perform automated tests on the application. This might be in integration with models, other applications, project settings, etc. This is a component that Django makes it to have easy and quick unit testing along with the Python modules for advanced testing. It is quite easier to integrate python modules and libraries into almost anything in the Django Project.

### apps.py

This is the file for app-level configuration. We can change the default fields, app name, email settings, other module-specific settings that can be used in the models, views, or in another place that can be defined here. 

## Other Folders/files

Apart from the app folder and the project folder, we may have other folders like the `templates`, `static`, `media`, etc. There are also python package-specific folders for which you may need to create folders.

### Templates

There are a couple of standard ways you can set up your Templates folder. Either in the root project or inside individual apps. The choice is yours, however, you feel comfortable. I personally use only one template folder in the root directory, but you can keep it wherever you want, but these two are the standard ones for ease of reading and maintaining the projects.


### Static 

The Static folder is the folder in which you store your `css`, `javascript`, and `images`(images or media files that are used in the templates). This is a good way to improve the performance as in the production the webserver collects all the static files and stores them in a single place for responding to the requests. 
The template folder if present in the root folder, has a sub-folder as the application names and inside the `app-name`, we put in all the `.html` or other template files. 

As similar to the `template` folder, the location can be modified or set as a configuration from the settings.py file. Usually, the static files(`.css`, `js`, etc) are stored in the root folder with app names as subfolders. 

### Media

The media folder is where you can store the media-specific to the user or the application processed data. For example, we can store the profile pictures of users, email attachments if it's an email application, thumbnails of the posts for a blogging platform, etc. 

The configuration of the Media folder is quite similar to the Static folder but it has certain additional configurations. We'll look at them in their sections in this series.

Phew! That was all the folder structure you need to get started with Django. There might be other folders and files specific for project, application, python modules but it won't be much hard to understand those as well.

## Conclusion

From this part, we were able to understand the folder structure of the Django framework. We explored the various files and folders with their use cases and their purpose. So, by reading the above description of the files and folders you might have got a rough idea about the flow of the development cycle in Django.

In the next part, we'll start with actually getting hands dirty in the code and making our first view. Thank you for reading and Happy Coding :)


