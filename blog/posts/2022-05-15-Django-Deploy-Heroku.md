---
templateKey: blog-post
title: "Django + PostgreSQL Deployment on Heroku"
description: ""
date: 2022-05-15 20:00:00
status: published
slug: django-deploy-heroku
tags: ['django', 'web-development', 'python',]
image_url: https://meetgor-cdn.pages.dev/django-deploy-heroku.png
series: ['Django-Deployment','Django-Series']
series_description: "Django Deployment is a series for understanding various platforms for Django project deployment with database attachment. Platforms to be covered: Heroku, Railway, Qovery, Python Anywhere, etc."
---

## Introduction

Django projects are quite easy to build and simple to understand, you might have created a Django application and wanted to show it to the world? You can deploy a basic Django application with a database(PostgreSQL) with Heroku. It provides a decent free tier with some great features and add-ons. A free tier Heroku account has a limitation of 5 apps, limited data in the database, limited connections to the server per month, and so on.  

Though the free tier is not a great option for bigger applications, it suits really well for smaller scale and ide projects, so we will be looking into the details of how to deploy a Django application to [Heroku](https://heroku.com/) which is a Platform as Service (PaS). 

This series will be an extension of the series [Django basics](https://techstructiveblog.hashnode.dev/series/django-basics) which covered the basics of the Django framework, we covered from basic Django fundamentals to building a CRUD API. In this series, we will be exploring some platforms for giving a better understanding of how things work and understanding a few components that were left as default while understanding the basics of Django. Let's get started with [Django Deployment](https://techstructiveblog.hashnode.dev/series/django-deployment)!

## Creating a Django Application

For deploying an app, we definitely need an app, we need to create a basic Django application to deploy on the web. We'll be creating a simple blog application with a couple of views and a simple model structure. As for the database, we'll be using Postgres as Heroku has an add-on for it and it is pretty easy to configure. 

### Set up a virtual environment

We need to set up a virtual environment in order to keep the Django project neat and tidy by managing the project-specific dependencies and packages. We can use the [virtualenv](https://pypi.org/project/virtualenv/) package to isolate a python project from the rest of the system.

```
# install the virtualenv package
pip install virtualenv

# create a virtual env for the project
virtualenv .venv

# activate the virtualenv
Windows:
.venv\Scripts\activate

Linux/macOS:
source .venv/bin/activate
```

This will set up the project nicely for a Django project, you now install the core Django package and get started with creating a Django application.


```bash
# install django
pip install django

# start a django project
django-admin startproject blog .

cd blog

# create a application in django project
python manage.py createapp api

# Create some models, views, URLs, templates

# run the server
python manag.py runserver
```

We assume you already have a Django project configured with some basic URLs, views and templates or static files as per your project and requirements, for this tutorial I will be using the simple blog application from my previous Django tutorials as a reference. You can follow along with my [Django Basics](https://techstructiveblog.hashnode.dev/series/django-basics) series and refer to the Blog Application project on [GitHub](https://github.com/Mr-Destructive/django-blog).

## Configuring the Django Application

Make sure to create and activate the virtual environment for this django project. This should be done to manage the dependencies and packages used in the project. If you are not aware of the virtual environment and django setup, you can follow up with this [post](https://mr-destructive.github.io/techstructive-blog/django-setup-script/).

### Creating a runtime.txt file

Now, Firstly we need to specify which type and version of language we are using. Since Django is a Python-based web framework, we need to select the python version in a text file.

**runtime.txt**
```
python-3.9.5
```
 
Here, the version can be anything as per your project and the packages installed.  

### Creating requirements.txt file

We'll first create a `requirements.txt` file for storing all the dependencies and packages installed in the application. This will help in installing dependencies while deploying the application. We can either use a `requirements.txt` file using `virtualenv` or a `Pipfile` using Pipenv. Both serve the same purpose but a bit differently. 

Assuming you are in an isolated virtual environment for this Django project, you can create a requirements.txt file using the below command:

Make sure the virtualenv is activated before running the command:

```
pip freeze > requirements.txt
```

This will create a simple text file that contains the package names along with the versions used in the current virtual environment. A simple Django requirements file would look something like this:

```
asgiref==3.4.1
Django==3.2.11
pytz==2021.3
sqlparse==0.4.2
typing_extensions==4.0.1
```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652614060461/kPTZ9R8Xp.png)

This is vanilla Django without any additional dependencies but if you have integrated other functionalities like Django Rest Framework, PostgreSQL, Crispy Forms, Schedulers, etc. there might be additional dependencies that become quite crucial for the smooth working of the project.  

If you are using pipenv, you don't need to make any efforts to manually activate and manage virtual environment, the dependencies are installed and taken care of by the pipenv installer. You just need to make sure to install any package with `pipenv install` and not `pip install` for better and improved package tracking.

So, that's all we need to take care of packages and keep a list of these integrated packages for the project. You need to update the requirements.txt file every time you install any new package or modify the dependencies for a project. Simply use the command `pip freeze > requirements.txt` in the activated virtual environment.  

### Creating a Procfile

Next up, we have the `Procfile`, a procfile is a special file that holds information about the processes to be run to start or activate the project. In our case, for django we need a web process that can manage the server.

A Procfile is a simple file without any extension, make sure to write `Procfile` as it is as the name of the file in the root folder of the project. Inside the file add the following contents:

**Procfile**
```Procfile
web: gunicorn <project_name>.wsgi
```

As we can see we have defined the `web` process using `gunicorn`, [Gunicorn](https://pypi.org/project/gunicorn/) is a python package that helps in creating WSGI HTTP Server for the UNIX operating systems. So, we need to install the package and update the package dependency list. 

```
pip install gunicorn

pip freeze > requirements.txt
``` 

That would be good to go for creating and serving up the project while deploying the project on Heroku.

## Creating a Heroku App

A Heroku App is basically like your Django Project, you can create apps for deploying your django projects on Heroku. You are limited to 5 apps on the Free tier but that can be expanded on the paid plans.  

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652456732519/cyOQZ3UZK.png)

The name of your heroku app should be unique globally, you need to try a few combinations before a good one fits your project. This name has no significance on your django project code, though it would be the name from which you would access the web application as a name `<app-name>.herokuapp.com`.   

So, choose it wisely if you are not attaching a custom domain. You can attach a custom domain, you can navigate to the `domain` section in the settings tab. 


## Setting up the database 

To set up and configure a database in django on Heroku, we need to manually acquire and attach a PostgreSQL add-on to the heroku app.

- Firstly locate to the Resources Tab in your Heroku app.
- Search `postgres` in the Add-ons Search bar
- Click on the `Heroku Postgres` Add-on
- Submit the Order Form and you have the add-on enabled in the app.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652456842273/ijeWsVdOf.png)

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652456877447/dLG30ac_m.png)

Alternately, you can use the Heroku CLI,

### Heroku CLI Setup

You can use the Heroku CLI which is a command-line interface for managing and creating Heroku applications.  You need to first install the CLI by heading towards the [heroku documentation](https://devcenter.heroku.com/articles/heroku-command-line). Once the CLI is installed, you need to login into your Heroku account by entering the following command:

```
heroku login

``` 

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652605604920/HnTr2KbTi.png)

This will allow us to work with heroku commands and manage our heroku application from the command line itself. The below command will create a add-on for `heroku-postgres` for the application provided as the parameter options 

```
heroku addons:create heroku-postgresql:hobby-dev --app <app_name>
```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652507166978/i1IJ5EGjJ.png)

This should hopefully add a fresh instance of a postgres database for your heroku app. You can now configure the database for your project, we simply need the Database URL from the heroku app dashboard. We'll see how to configure the environment variables in Django for Heroku to keep your secrets like the `SECRET_KEY`, `DATABSE_URL`, etc. 

If you want MySQL as a database, you can check out the [ClearDB](https://devcenter.heroku.com/articles/cleardb) Add-On for Heroku to simply attach it like Postgres on your Heroku application. Also, if you wish to add [MongoDB](https://www.mongodb.com/compatibility/mongodb-and-django) into your Django application on Heroku, you can use [Object Rocket](https://devcenter.heroku.com/articles/ormongo)(OR Mongo). It is not available in the free tier though, unlike PostgreSQL and MySQL.

## Configuring Environment Variables

We need to keep our secrets for the django project out of the deployed code and in a safe place, we can create environment variables and keep them in a `.env` file which will be git-ignored. We do not want this `.env` file to be open source and thus should not be committed.  



We first need to create a new secret key if you already have a GitHub repository, chances are you would have committed the default secret key open for the world to see, it is an insecure way of deploying django apps in production.

So, open up a terminal and launch a python REPL:

```
python

import secrets
secrets.token_hex(24)
```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652512239319/5AACaTGOD.png)

This should generate a secret key that only you know now. So, just copy the key without the quotes and create a file `.env` in the root project folder. 

**.env**
```
SECRET_KEY=76419fd6885a677f802fd1d2b5acd0188e23e001042b05a8
```

The `.env` file should also be added to the `.gitignore` file, so simply append the following in the `.gitignore` file

```
.env
```
This file is only created to test the project locally, so we need to also make this key available on heroku. For doing that we need to add Config Variables to the heroku app. 

- Locate to the Settings Tab in your Heroku Application Dashboard
- We have the `Config Vars` section in the located tab
= We need to reveal those variables and we will be able to see all the variables.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652456988713/5VM6E29_o.png)

We already have a `DATABASE_URL` variable declared when we attached the `django-postgres` database to our application. We will now add one more variable to the Application, i.e. the `SECRET_KEY`. Simply, enter the name of the key and also enter the value of the key, so basically a key-value pair. Simply click on the `Add` button and this will add the variable to your application.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652515244870/LRyPzJr41.png)

You also need to copy the `DATABSE_URL` into our local setup file(`.env` file). Copy the Database URL and paste it into the `.env` file as follows:

```env
DATABASE_URL=postgres://sjxgipufegmgsw:78cbb568e@ec2-52-4-104-184.compute-1.amazonaws.com:5432/dbmuget
```
The format for the postgres URL is as follows:

```
postgresql://[user[:password]@][netloc][:port][/dbname]
```

We have now created environment variables for our django application and also added config vars in the heroku app, we now need a way to parse these environment variables into the Django project.  

### Parsing Environment variables using python-dotenv

We will use [python-dotenv](https://pypi.org/project/python-dotenv/) to parse variables into the django settings configurations like `SECRET_KEY` and `DATABASES`. 

- Install `python-dotenv` with pip with the command :  
```
pip install python-dotenv
```
We need to then modify the default variables in the `settings.py` file. Firstly, we will load in the `.env` file for accessing the environment variables for the configuration of the project.

Append the following code, to the top of the `settings.py` file, make sure don't override the configuration so remove unnecessary imports and configurations.

``` python
# <project_name>/settings.py

import os
from dotenv import load_dotenv

BASE_DIR = Path(__file__).resolve().parent.parent

load_dotenv(os.path.join(BASE_DIR, ".env"))

```

We have imported the package `dotenv` basically the `python-dotenv` into the `settings.py` file and the module `os` for loading the `.env` file. The `load_dotenv` function helps in loading the `key-value` pairs which are the configuration variables that can act as actual environment variables. We provide in a file to the [load_dotenv](https://saurabh-kumar.com/python-dotenv/) function which is the `.env` file in our case, you can pick up any location for the `.env` file but make sure to change the location while parsing the file into the `load_dotenv` function. 

After loading the variables into the `settings.py` file, we now need to access those variables and set the appropriate variables the configuration from the variables received from the `load_dotenv` function. The `os.getenv` function to access the environment variables. The `os.getenv` function takes a parameter as the `key` for the environment variable and returns the value of the environment variable.

``` python
SECRET_KEY = os.getenv("SECRET_KEY")
```

We have secretly configured the `SECRET_KEY` for the django project. If you have any other variables as simple key-value pairs like `AUTH` passwords, username, etc. you can use this method to get the configuration variables. 

### Loading Database configuration

Databases are a bit different as compared to other simple configurations in django. We need to make a few adjustments to the default database configuration. We will install another package `dj-database-url` to configure `DATABASE_URL`. Since the databse_url has a few components we need a way to extract the details like the `hostname`, `port`, `database_name`, and `password`. Using the `dj-database-url` package we have a few functions that can serve the purpose.

```
pip install dj-database-url
```

At the end of your `settings.py` file, append the following code. 

``` python
import dj_database_url

DATABASE_URL = os.getenv("DATABASE_URL")

DATABASES = {
    "default": dj_database_url.config(default=DATABASE_URL, conn_max_age=1800),
}
```

We would need an adapter for making migrations to the `PostgreSQL` database i.e. the `psycopg2` package. This is a mandatory step if you are working with `postgres` database. This can be installed with the simple pip install:

```
pip install psycopg2

# If it does not work try
pip install psycopg2-binary


# if still error persists try installing setuptools
pip install -U setuptools
pip install psycopg2
```

Now, that we have configured the database, we can now apply migrations to the fresh database of postgres provided by heroku. We will simply run the migrate command and in the heroku app the PostgreSQL database would have been modified and an appropriate schema should be applied.

```
python manage.py migrate
```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652602284553/oTtGev28-.png)

Make sure to update the `requirements.txt` file before pushing the project to Heroku for making sure everything works as expected. Since we have installed a few additional packages that are directly used in the `settings.py` file, we need to run the `pip freeze` command to update the `requiremnets.txt` file.

## Serving Static Files

Now, if you have some static files like `CSS`, `Javascript`, or `images`, you need to configure the staticfiles in order to serve them from the heroku server. We will require another config variable for collecting the static files from the selected repository. 

```python

STATIC_URL = "static/"
STATICFILES_DIRS = [os.path.join(BASE_DIR, "static")]
STATIC_ROOT = os.path.join(BASE_DIR, "staticfiles")

``` 

Here, if you have served your static files from the `static` folder in the root directory of your django project, you can add the above code in the settings.py file. We will collect all static files in the folder along with the default static files provided by django in the `staticfiles` directory. Run the following command if you want to test whether the static files are properly collected and served.

```
python manage.py collectstatic 
```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652600828657/JgL4nLuiL.png)

So, this command will collect all the static files and store them in a single place. We see that the files from the admin section are also copied as well as the custom static files from the project configuration. Now, we can move on to set the config variable for the heroku app.

```
DISABLE_COLLECTSTATIC = 0
```

We can set the `DISABLE_COLLECTSTATIC` variable as `0` or `1` indicating whether to disable it or not. We have currently enabled the static file collection while deploying the app but you can set it to `1` to disable the collection of static files.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652613798420/mbqzf1Kqd.png)

Since I first tested the application on heroku, the static files don't work as expected, we need another package to make sure the staticfiles are served property. We will be installing the `whitenoise` package which serves as the middleware for serving the static files.

```
pip install whitenoise
```

Add the whitenoise middleware `whitenoise.middleware.WhiteNoiseMiddleware` to the `MIDDLEWARE` list in the `settings.py` file.

````python
MIDDLEWARE = [
...
...
...
    'whitenoise.middleware.WhiteNoiseMiddleware',
]

```

That should be enough to make the most of the deployment on heroku. You will have to make a few adjustments as per your requirements and project.

## Deploy from GitHub

We are now all set to deploy the application on Heroku, we can use the `Connect to GitHub` or `Heroku CLI` to push the code into production. Heroku CLI is quite easy with a few sets of commands but if your project is deployed on GitHub, you can straightaway let the deployment start the build on a push to a specific branch. This becomes quite automotive and easy to scale while deploying a large-scale application. 

```
pip freeze > requirements.txt
```

This step is quite important because you need to make sure that all the packages are listed in the `requirements.txt` file else you will have to wait for the build to fail and redeploy.

Make sure the server is running first on your local machine, remember the server will be set up from scratch but the database will already have applied migrations if you have applied migrations before after connecting the Heroku Postgres database.
 
```
python manage.py collectstatic

python manage.py runserver
```

This will set up the origin of the remote repository that will be pushing the project code. Next, make sure to commit the code which will contain all the required stuff for deploying the code.

Checklist for deploying the code

```
- requirements.txt
- Procfile
- runtime.txt
- django-project
- environment variables / config variables 
- static file configuration
- database configuration
- migrate schema of database 
- gitignore file for ignoring virtualenvs, .env file, staticfiles, etc
```

here's a sample `.gitignore` for a minimal django project.

```gitignore
.env/
.venv/
env/
venv/
*.env

*.pyc
db.sqlite3
staticfiles/
```

If you want a full-fledged `.gitignore` for a complex django project, you can take the reference from Jose Padilla's [gitignore Template](https://github.com/jpadilla/django-project-template/blob/master/.gitignore) for a django project.  

### Git Commit the Django Project
```
git status 

git add .

git commit -m "config for heroku deployment"
```
Carefully check the status of the git repository before committing and make sure you don't forget anything by mistake, it won't a big problem but it would mess up the build process.


![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652613991525/hxQgtGOoM.png)

After committing the code, we can now push the code to GitHub. We first need to set the remote repository reference to be able to push the code to it. 

```
git remote add origin https://github.com/<username>/<repo_name>
```
This will set up the `origin` as the remote repository on GitHub. Once the remote repository is created, we can move into the push of the git repository that will trigger the build. First, navigate to the `Deploy` section in the heroku app's dashboard where we want to connect the `GitHub` repository and allow the automatic deploy from a branch in this case we have chosen the `main` branch.

Due to some `Heroku` Internal Server Issues, the GitHub integration seems to have broken and isn't working as of May 2022, but it might work later. 

Below is a screenshot of my previous project deployed to Heroku using a GitHub repository.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652605497382/5VuQUQ0t0.png)

```
git push origin main
```

This will prompt you for your GitHub credentials and will deploy the commits to the remote repository on GitHub. This push on the main branch should also trigger the build process of the heroku app for this django project. You can navigate to the Activity section for the Build logs. 

If you have followed the article well, and your repository has all the correct configurations, the build will succeed, else chances are you might have missed a few things and the app might have crashed. You can debug your application build with the simple heroku CLI command:

```
heroku logs --tail -a <app_name>
```

This can be quite handy and saves a lot of time in understanding what went wrong in the build. It might be related to database migration, static files, python package not found, and some silly mistakes and errors that can be fixed after committing the code and pushing it to GitHub again.

If you do not want a GitHub repository, you can directly push the code from the local git repository to the remote heroku app center. This will require us the Heroku CLI.

## Heroku CLI

We can use the heroku CLI for pushing the code via the simple git repository. We can push the references via the branch and a remote repository on heroku to build our app.  For this, we assume you have heroku installed and logged in. We will require the django project code and heroku cli to build the django web application.

```bash
heroku git:remote -a <heroku_app_name>

# for my case
heroku git:remote -a blog-django-dep
```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652614221069/vCAKD0zsz.png)

After this, you can commit your code and the project as git repository. We have added the remote repository location on heroku, we can now simply push the code to the remote repository.

```
git push heroku main
```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652614125785/uEzFQ9VvQ.png)

Here, `heroku` is the remote repository location and `main` is the branch of the repository. This will push the code to the repository and thereafter create a build to deploy the Django project as a Heroku application.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652614381808/kYTmB3EO2p.png)

You can hit the `Open App` button on the top right corner and there should be your deployed Django application. 

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652610395538/xjUiODhoK.png)


## A Few Tricks and Gotchas

There are a few tricks and issues that you might pop into while deploying a django project on heroku, especially the build process. It requires a few iterations to get the complete app setup.

### Run command from the Dashboard console

If you don't have heroku CLI set up and want to fix a few things on the pushed code to the heroku app, you can use the `Run Console` option from the `More` Tab on the top right corner of theApplication dashboard. This is a great way to run migrations, configure static files or tweak a few things here and there without messing up the code on GitHub or the remote git repositories. 

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652614775294/lgDPwr2yr.png)


![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652614821950/uTzQVB8sC.png)


![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1652614845269/BkZhu3SGH.png)

### Deploy with Docker 

You can even use the docker container to deploy a Django application on Heroku. It is a great way of learning a lot of deployment strategies and techniques using Docker. You'll get familiar with interesting concepts like virtualization, and containerization, and also learn Docker on the way. You can follow this tutorial on [Deploying Django applications with Docker on Heroku](https://testdriven.io/blog/deploying-django-to-heroku-with-docker/). Also, you can check out the official Heroku documentation for [deploying python applications](https://devcenter.heroku.com/articles/deploying-python).

### What are Dynos?

Dynos are simply web processes or workers that serve your web application. Dynos in Heroku are allocated based on the build process, once the slug is created a dyno is created as it runs on a VM container. This simply means there are limitations on how to use the web application and its sleep process. The hobby tier is sufficient for normal testing projects and side projects though you will have to pay and move into advance tiers to increase the dyno allocations and scaling of those web processes. 

It's not a simple thing to understand but to keep it simple, it might be a container to process the client's request and serve the page for a finite duration of the interaction. Also, your application will sleep after half an hour, if you try to reload the application every half an hour it will consume your resource allocation for the month and this is how the tiers and divided for paid services on Heroku. You can check out the detail over [here](https://www.heroku.com/pricing#containers).

## Conclusion

So, that is one of the ways we can deploy a Django application on Heroku with the PostgreSQL database. You can find the [django-blog project] on [GitHub] for following along with the deployment process.  In the next few parts of the series, we will be hopefully covering other platforms where you can deploy a Django o application.

Hopefully, you liked the above tutorial, if you have any questions. feedback, or queries, you can contact me on the Social handles provided below. Thank you for reading and till the next post Happy Coding :) 
