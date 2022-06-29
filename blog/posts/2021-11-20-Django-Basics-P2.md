---
templateKey: blog-post
title : "Django Basics: Setup and Installation"
subtitle: "Setting up environment and installing Django framework"
date: 2021-11-20 16:30:00 +0530
status: published
tags: ['django', 'python', 'web-development',]
slug: django-basics-setup
image_url: https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643290071/blogmedia/s8ahlep1e8lmgiboyjhz.png
series: ['Django-Basics', 'Django-Series']
---

## Introduction

The crucial aspect of starting to learn any framework is the ease to set it up and Django by far is the easiest of the options out there. There is just a few lines of code to install django if you already have python installed in your system. In this article, we see how to setup a django project along with a virtual environment. 

If you already have python and pip installed, you can move on to the [virtual environment setup](#setting-up-virtual-environment-in-python).

## Installing Python and PIP

Django is a python based framework so that makes sense to have Python installed along with its package manager to use Django.  

To install Python, you can visit the official [Python](https://www.python.org/downloads/) website to download any relevant version for your system (recommended 3.7 and above). 

Mostly the Python installation comes with the option to install `pip`(python's package manager) but if you missed that, that's fine, you can install the [get-pip.py](https://bootstrap.pypa.io/get-pip.py) file into your system and run the below code:

```
python get-pip.py   
```

Make sure the include the relative path to the file if you are not in the same folder as the file.

So, that should be python setup in your local machine. To check that python was installed correctly, type in `python --version` and `pip --version` to check if they return any version number. IF they do, Congratulations !! You installed Python successfully and if not, don't worry there might be some simple issues that can be googled out and resolved easily. 
   
---

Let's move on to the actual setting of the Django project set up.   

## Setting up Virtual Environment in python

Virtual Environment is a software which isolates the installation of dependencies and libraries for a specific project, making it a clean and safe environment for deployment as well as maintenance. 

In Python, we have a virtual environment package known as `virtualenv` that does this thing. It is for installing the Python related packages into a isolated folder. So, we can install the `virtualenv` package in python by following the following steps:

### Installing Virtualenv
 
Firstly, install the virtual environment package, it's not mandatory but it keeps things simple and easy for your project in correspondence to the entire OS. So in python, we have a module to create the virtual environment pretty easily,

```
pip install virtualenv
```

You can use `pip3` or `pip -m`, or however you install normal python modules. This just installs the python virtual environment, we need to create a virtual environment in the current folder.

### Creating a virtual environment

We need to create the environment so as to give the Python interpreter an indication to consider the current folder as an isolated Python environment. We need to create a virtual environment in the current folder, so for that navigate to the folder where you want to create the project and enter the following command: 

```
virtualenv venv
``` 

Here, `venv` can be anything like `env` just for your understanding and simplicity it's a standard name kept for the same. After this, you will see a folder of the same name i.e. `venv` or any other name you have used. This is the folder where python will keep every installation private to the local folder itself. 

### Activating Virtual environment 

Now, we need to activate the virtual environment, this means that any thing installed in the prompt with the virtualenv activated will be isolated from the entire system and will be installed on in the virtual environment. To activate the environment, we can use the command :

#### for Linux/macOS :

```
source venv/bin/activate
```

#### for Windows:

```
venv\Scripts\activate
```

After this, your command prompt will have a `(venv)` attached in the beginning. This indicates you are in a virtual environment, things you do here, may it be module installation or any configuration related to python will stay in the local folder itself.


## Installing Django

After the virtual environment is set up and activated, you can install Django and get started with it. Django is a python module or package, which can be easily installed using its package manager `pip`. 

Install Django using pip:

```
pip install django
```

## Create a Django Project

After the installation is completed, you can start a Django project in the current folder from the django package we installed. There are several commands available in the django module which you can execute in the command line that we'll discuss later.
For now, we will use the command `startproject` this is one of the [management commands](https://github.com/django/django/tree/main/django/core/management/commands) in Django. The [django-admin](https://docs.djangoproject.com/en/3.2/ref/django-admin/) is a command line utility for doing the administrative tasks related to Django.

```
django-admin startproject myproject
```

Here `myproject` can be your project name. After this, you will see one new folder and one file pop up.

Namely, the `<project-name>` folder and `manage.py` file. We don't have to touch the `manage.py` file but we use it in most of the commands to use the Django functionalities, it is quite similar to the `django-admin` command. 

You can now run your basic server using the command : 

```
python manage.py runserver
```

OR

You can use `djagno-admin` command, but you need to set certain environment variables and modify the settings.py file as per the project-name. You can use the `django-admin` as the steps given in the django [documentation](https://docs.djangoproject.com/en/3.2/ref/django-admin/#cmdoption-settings).

The output of the command `python manage.py runserver` should be visible in the browser at `https://127.0.0.1:8000` as below :

![Django-Base-Project](https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/screenshotr_2021-11-20T15-40-50.png)

That's it the base django project is installed in your system. To stop the server simply press `Ctrl+C`. 

Follow the below GIF for a clear understanding of those instructions:

![Django-basics-part2-setup](https://gitlab.com/MR_DESTRUCTIVE/tblog-img/-/raw/main/djp2.gif)

---

## Quick-Setup-Script

You can avoid manually typing the commands once you get the idea of the process in setting up a django project by executing a simple shell script (for Linux/macOS) or a batch script (for Windows). The script looks something like:

For Linux/macOS:

```shell
#!/usr/bin/env bash

mkdir $1
cd $1
pip install virtualenv
virtualenv venv
source venv/bin/activate

pip install django
django-admin startproject $1 .
clear
```

Save as commands.sh file

For Windows: 

```batch
mkdir %1 
cd %1
pip install virtualenv
virtualenv env
call env\Scripts\activate

pip install django
django-admin startproject %1 .
cls

```   
save as commands.bat file

For further instructions you can checkout the [GitHub repository](https://github.com/Mr-Destructive/django-quick-setup-script) or a detailed [article](https://mr-destructive.github.io/techstructive-blog/django/web-development/python/2021/08/15/Django-Quick-Setup.html) about it.

## Conclusion

From this section, we were able to setup the Django project in our local system. In the next part, we will cover the `folder structure` of the Django project. We won't directly go into the code part because that is very easy once you understand the flow of the framework and its internal working. So, thanks for reading and Happy Coding :)
