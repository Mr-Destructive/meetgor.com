---
templateKey: blog-post
title : "Django Basics: Static Files"
subtitle: "Loading Static Files in a Django project/app"
date: 2021-12-27 22:45:00 +0530
status: published
tags: ['django', 'python', 'web-development',]
slug: django-basics-static-files
image_url: https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640610003/blogmedia/dj-static-6_pjipoj.png
series: ['Django-Basics','Django-Series']
---

## Introduction

After creating templates, it should be rather tempting to add some styles and logic to them. Well yes, we'll see how to add static files in a web application using django. Static files are not only CSS, but also media/images and Javascript files as well. In this part of the series, we'll cover the basics of working with static files in django including the configuration, rendering and storing of the static files. 

## What are Static Files?

Static files as the name suggests are the files that don't change, your style sheets(css/scss) are not gonna change for every request from the client side, though the template might be dynamic. Also your logo, images in the design will not change unless you re-design it XD So these are the static files that needs to be rendered along with the templates.

We have basically 3 types of static files, CSS, Javascript files and media files/static templates,etc. They are all rendered in the same way but as per their conventions and usage. 

You can learn about the theoretical information on [static files](https://docs.djangoproject.com/en/4.0/howto/static-files/) from the django documentation.

## How to configure Static Files

Firstly you can create a folder for all the static files in the root folder. Usually the convention is `static` as the name of the folder. So, if you have created the template folder in the root directory, similar to that static folder can be created in that path. 

Next after creating the static folder in the project root folder, we need to configure the `settings.py` file to actually tell Django web server to look for all our static files in that folder. To do that, go to the `settings.py` file, now by this time you would have known where the `settings.py` file is (inside the project-named folder). Add the following at the end of the `settings.py` file.

```python
# import os
# STATIC_URL = '/static/'

STATICFILES_DIRS = (
    os.path.join(BASE_DIR, "static/"),
)
```   

Ignore the `import os` if you already have imported and the `STATIC_URL` if already there in the file. The `STATICFILES_DIRS` is the configuration that we tell the django environment to look for all our static files in the base/root directory of the project where the `static/` folder is. The `os.path.join()` actually gets the path of the directory in our operating system to the folder specified in the case of our project the `BASE_DIR` is the path of the project and we add in the static folder to actually the project path. The final piece and the crucial one is the `"static/"` path, this can be other location where you have created your static folder within the project.

That's it! Yes, it's that simple. We can now create static files and render them in our templates. 

## Creating and Storing Static files

Now this part is customizable and it depends on your preference, how you want to organize the static folder. The convention that I follow is creating separate folders namely for `css`, `js` and `assets`(or `img`) mostly. And inside of this folders you can store the respective static files. This also creates the project more scalable in terms of it's maintenance. 

```
static\
  |__css\
  |__js\
  |__assets\
```

Let's create a static file and an image to demonstrate the concept of static files in django. 

- css/style.css

```css
body 
{
    background-color:#1d1dff;
    color:white;
}

h1
{
    text-align:center
    font-family: monospace;
}

p
{
    color:#ff6600;
    font-weight:500;
}

ul
{
    list-style-type:square;
}
```

- assets/tbicon.png 

Demo Image (that's my blog icon)   

![Demo image](https://github.com/Mr-Destructive/techstructive-blog/blob/gh-pages/assets/img/tbicon.png?raw=true)

## Rendering Static Files from Templates

So, after configuring and creating the static files, we now can inject them into our templates. If you try to do the traditional way i.e. linking stylesheets/images/script files with HTML, it just won't work as you expect to and there's no point in using traditional way while creating a web application with a framework. So, there is a framework specific way to do things which make it easier and efficient for the project. 

To render any static file, we need to load the static tag which allows us to embed links for the static files into the templates. This means if the static files are not loaded directly instead in production(deploying our application) the static files are stored in a folder `STATIC_ROOT` which the server then loads, we'll see how that internally works when we get to deployment techniques for Django project. 

To load the static files from our configuration, we can simpy include the tag on top of the template.

```
{% load static %}
```

The above templating tag will load the `static` tag which allows us to embed the links to the static files as explained earlier. 

Now, we can actually access any file with the static folder in our templates with a particular syntax as below:

```html
<link rel="stylesheet" href="{% static 'css/style.css' %}">  
```   
Its just a example how to load the file, we are calling the static tag which we have loaded in previously and from there we are referencing the css file. The compact syntax would be : `{% static  'path-to-file'  %}`      

**NOTE: The path to the static file is relative from the Static folder, i.e. enter the path of the file considering the static folder as the base directory.** 

### Demonstration of the static file

Let's render the static file which we created earlier i.e. the css file and the image into a template. 

Assuming you have a app called `post` in your django project, you can render static files as below:

# templates/home.html 
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Django Blog</title>
    {% load static %}
    <link rel="stylesheet" href="{% static 'css/style.css' %}">  
</head>
<body>
    <h1>Hello, World!</h1>
    {% block body %}
    <p>This is not going to get inherited </p>
    {% endblock %}
    <p>This will be inherited</p>
</body>
</html>
```   
We are loading the static tag and then loading the css file using the tag syntax as explained above.       

# static/css/style.css      
```css
body 
{
    background-color:#1d1dff;
    color:white;
}

h1
{
    text-align:center
    font-family: monospace;
}

p
{
    color:#ff6600;
    font-weight:500;
}

ul
{
    list-style-type:square;
}
```   
This is the static file,`style.css` stored inside the css folder of the static folder. This contains basic (very lame) CSS styling as we can understand.     

# post/views.py 

```python
from django.shortcuts import render

def home(request):
    return render(request, 'home.html')
```
The `views.py` file has the function that renders the template `home.html` from the templates folder inside the application specific folder.   

# post/urls.py   
```python
from django.urls import path
from post import views

urlpatterns = [
        path('', views.home, name="home"),
        ]
```   
This is the application level configuration for the url routes to the views linking the views(functions) from the `views.py` file. The url in this file(code-snippet) is linking the root url('') to the home view in the `views.py` file.

# Blog/urls.py
```python
from django.contrib import admin
from django.urls import path, include
from django.views.generic import TemplateView

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', include('post.urls')),
]
```
The urls file in the project folder is the core configuration for project level url routes to individual applications within the project.

Append the following if your templates and static files are not configured properly.

# Blog/settings.py
```python
import os   

TEMPLATES = [
    {
        'BACKEND': 'django.template.backends.django.DjangoTemplates',
        'DIRS': [os.path.join(BASE_DIR, 'templates'),],
        'APP_DIRS': True,
        'OPTIONS': {
            'context_processors': [
                'django.template.context_processors.debug',
                'django.template.context_processors.request',
                'django.contrib.auth.context_processors.auth',
                'django.contrib.messages.context_processors.messages',
            ],
        },
    },
]
STATIC_URL = '/static/'
STATICFILES_DIRS = (
    os.path.join(BASE_DIR, "static/"),
)
```
SO, the result of the above code is as simple template as shown in the picture below:

![Static file demo](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640621276/blogmedia/static-1_vu41gf.png)

This will also work if you do it with traditional HTML syntax, but I'd explained why it's not recommended to do it while using frameworks.

Let's see how static files are rendered in inherited templates. We'll tinker with the `for.html` template created in the [previous part](https://mr-destructive.github.io/techstructive-blog/django/python/web-development/2021/12/14/Django-Basics-P5.html).  

# template/for.html
```django
{% extends 'home.html' %}
{% load static %}

{% block body %}
    <img src="{% static 'assets/tbicon.png' %}" height="50px" width="50px" />
    <ul>
        {% for sport in sport_list %}
        <li>{{ sport }}</li>
        {% endfor %}
    </ul>
{% endblock %}
```
We will have re-load the static tag for each template only if we need to include a new static file in the template. So we use the `{% load static %}` again as we are loading the static file (image) in this template.

# post/views.py
```python
from django.shortcuts import render

def for_demo(request):
    sports = ('football', 'cricket', 'volleyball', 'hockey', 'basketball')
    return render(request, 'for.html', {'sport_list': sports})

def home(request):
    return render(request, 'home.html')
```

# post/urls.py
```python
from django.urls import path
from post import views

urlpatterns = [
        path('', views.home, name="home"),
        path('for/', views.for_demo, name="fordemo"),
        ]
```

So, that's the url and view map created, we can now be able to see the result in the `127.0.0.1:8000/for/` url to see the below result:

![Static demo for inheritance of tempaltes](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640622976/blogmedia/static-tempinh_peyjrg.png)

The list style has been changed and thus we can see that the CSS from the parent template is also being inherited. 

Here is the django project structure which I have created with this series so far:

![Folder tree structure](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640624705/blogmedia/trr-static_bgt9du.png)

So that has been it for the Static files in Django. Though there are lot of depth for rendering and loading the static files, we'll explore as we get our grasp in the django and web development terminologies. 

## Conclusion

So, from this article, we were able to configure and render static files like CSS/Images and optionally Javascript into the Django application. We covered from ground how to configure, load and structure the folder for storing all the static files at the project level. 

Hope you found it helpful and if you have any queries please let me know. We'll start with the databases probably from the next part in the Django Basics Series. Until then have a great week and as always Happy Coding :)
