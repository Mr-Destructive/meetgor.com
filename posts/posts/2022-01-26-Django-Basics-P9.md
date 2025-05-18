{
  "type": "posts",
  "title": "Django Basics: Admin Section",
  "subtitle": "Exploring and Interacting with the Django's built-in Admin Section",
  "date": "2022-01-26 22:25:00 +0530",
  "status": "published",
  "tags": [
    "django",
    "python",
    "web-development"
  ],
  "slug": "django-basics-admin",
  "image_url": "https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643202559/blogmedia/k5oebqxixvyhn8o60qpo.png",
  "series": [
    "Django-Basics",
    "Django-Series"
  ]
}

## Introduction

In the previous section, we configured the database in our Django project. We will move ahead and interact with the Admin Section in Django. Django has a built-in Admin Section which we can use to manage our apps and models. We will create a admin account(superuser) and try to render our custom Model in the Admin Interface. We shall further discuss the customization in the Admin section as per our requirements and necessity. 

## What is the Admin Section?

Admin Section is a powerful built-in utility provided by Django. It gives the administrative rights over the web project, the interface is neat and provides out of the box functionality to interact with the models in our project without us manually creating any mapping the views and urls. It is restricted to only superusers or trusted users to use as it is for administrative purpose. 

The Admin section is present by default for any django application. The interface provides the Django User and Group Model by default. Additionally we can have our own custom models to interact with. For every registered model you have the CRUD (create / read / update / delete ) functionality which makes it very easy and convenient to test the working of model before working around with APIs or moving ahead in the project.  

## Setting up an admin account (superuser)

In order to access the Admin section, we need to create a superuser. A superuser as the name suggests is a user who has the supreme authority for performing operations in the project in this case a web application. To create a super user we need to run a command from the command line that takes our name, email and password as input to create the super user. 

```bash
python manage.py createsuperuser
```

This will prompt you for a couple of things like :
- `username` the default is `admin`.
- `email` it's not necessary to put one.
- `password` should be at least eight characters long

The password input will be silent which means you cannot see what you type for security reasons, and the password field will be confirmed once, so you'll have to enter the password once more. But that's all you have to do to create a super user in Django for your web project.

![createsuperuser demo](http://res.cloudinary.com/dgpxbrwoz/image/upload/v1643198415/blogmedia/etubc7efwls47n3cj2bw.gif)

## Navigating the admin section 

After creating a superuser, we can now navigate the admin section from the browser. The admin section is by default located in the `/admin` url-route i.e. you need to navigate to `http://127.0.0.1:8000/admin` here `8000` can be any port as your preferred port number for the django application. 

After vising the Admin route, you will be prompted to a Login Screen. You simply need to add in the username and password which you entered while creating the superuser a while ago and you should be in the Admin Section. The default admin section as of Django `3.2.9` looks like following: 

![Django Admin Section](http://res.cloudinary.com/dgpxbrwoz/image/upload/v1643199349/blogmedia/h0k9jkqxozrtxvwsapkj.png)

## Components of the Admin Section

The Admin interface looks quite simple but is quite powerful and customizable. We have the Default Models in Django which are labelled in the `Authentication and Authorization` section namely the `Users` and `Groups`. You can see we have `+ Add` and the `Change` links to actually create the data associated with those Models. In the Admin Section you can basically play around with your models, it was not designed to act as the frontend for your application that's what the Django [documentation](https://docs.djangoproject.com/en/4.0/ref/contrib/admin/#module-django.contrib.admin) says and is absolutely correct. 

So, we don't have much things to explore when the UI is concerned as it is simple and straight forward to understand. We'll dive into how to register our models into the Admin section and from there on we can explore the UI to perform CRUD operations.

### Built-in Models

Django has two built-in and registered models in the Admin Section as said earlier.

1. Users
2. Groups

Users is basically the Django's User Model which provides the basic Authorization functionalities which further can be added to the we application. The Super User that was created from the `createsuperuser` command was associated with the Django User model. 

We have basic fields in the User model like:

1. Username
2. Email-ID
3. Password

If we go to the route `http://127.0.0.1:8000/admin/auth/user/add/`, we can see a form like UI that allows us to add a User.

![User Creation Form - Admin](http://res.cloudinary.com/dgpxbrwoz/image/upload/v1643208170/blogmedia/m3fdul2qcmgvgegm5r6y.png)

But there is no option for the `Email-ID` that's because for the User model has been modifies from the base [User Model](https://docs.djangoproject.com/en/4.0/ref/contrib/auth/#user-model) and we can see the `superuser` has all the attributes the `User` class has like the email, first name, last name and so on. 

![Super User Model - Admin](http://res.cloudinary.com/dgpxbrwoz/image/upload/v1643208828/blogmedia/wrdfkiqaqmw4wxtpopgn.png)

We can even extend this functionality or modify the existing attributes of the User Model in our custom Model. For example, we can add Age, Phone number, etc in to our Custom User Model. How? We'll look into that later but that is to give an idea about the User Model.

There is a lot more than just the details like Username, email and password. We need a way to manage which user is allowed to access what components. So, this is termed as `Permissions` in the User model, for the super user we might have the access to the admin page but a regular User might not. This is a permission which is by default implemented by Django in the User model. Like wise we can extend this functionality to add more permissions depending on the Model we are working with. 

### Groups 

This is the model which can hold a group of certain Model. The typical example hers is a group of User and its permissions. We can have a group for the developers of a project(this project) and a separate Group for rest of the Users. This creates a well-defined boundary for different types of User in a larger application.  

Currently, we don't have any groups created by default. Since it is a concept to be learnt for a quite large project with thousands of Users.

## Registering Models in Admin Section

Now, what are the default Admin section looks like, we can move on to register our own models in the Admin section. To do that, inside the app folder(a django app) you will see a `admin.py` file. Make sure to be in the app in which you have created a model. We need to register a Model to the admin section. 

```python
# app_name/admin.py

from django.contrib import admin
from .models import Article

admin.site.register(Article)
```
The `admin.site.register` basically adds a Model to the Admin Interface. 
The article Model is defined as follows:

```python
# app_name/models.py

from django.db import models

class Article(models.Model):
    title = models.CharField(max_length=127)
    post = models.TextField()
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now=True)
```

This will add the Model `Article` in the Admin Section. We can now perform CRUD operations in the Admin section. 
   
## Performing actions with Admin Section

So, perform CRUD operations, we can navigate to the Model Article and simply click the `Add Article` Button to add a object of the model Article. We will be presented a form to fill. Now here we, can see the fields which are actually to be inputted by the user. We don't see the fields like `created` and `updated` as they are automatically set as per the current time. 

![Add Article - Admin](http://res.cloudinary.com/dgpxbrwoz/image/upload/v1643215768/blogmedia/bq0gvbxhhxzwiwutgqpi.png)
After filling this form you will see `Article object (1)` which looks like a non-sense thing to look at. This is where the `__str__` [function](https://docs.djangoproject.com/en/4.0/ref/models/instances/) comes handy and saves the day. 

![Article Object](http://res.cloudinary.com/dgpxbrwoz/image/upload/v1643215997/blogmedia/sy7ygnskxfl0psgaj4z4.png)

Now, we can see we also have the ability to Delete and Update the Post. This is the best thing about the Django admin interface. It's quite intuitive for complex model to be tested before moving ahead and making necessary corrections if required.

### Modifying the Model (without migrations)

Now, at this stage if we forgot to add the `__str__` function, we need to add it into our models. But what about the migrations? We do not need to migrate this changes as there is no change in how to model is structures. We are changing how to Admin Interface should present our model. So, we'll make changes to our model but only aesthetically. 

```python
#app_name/models.py

from django.db import models

class Article(models.Model):
    title = models.CharField(max_length=127)
    post = models.TextField()
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now=True)

    def __str__(self):
        return self.title
```

And now if we refresh the Admin page, we can see the title is the object name :)

![Model str function](http://res.cloudinary.com/dgpxbrwoz/image/upload/v1643216721/blogmedia/vwyoccgmhnl4aosqc6qf.png)

This looks a trivial change but makes a big difference for large applications and dataset. 

### Verbose name

We can add a verbose name for an attribute inside an Model. A verbose name is a human readable name for a field. So, let's say we have a attribute/field called `fname`, the person who might have created the model might know it stands for `first name` but someone else might not. So in this case, we can add the `verbose_name` to be used as the name in the Admin Section. 

We also have `verbose_name_plural` which will be handy for a model name. We do not wnat Django to just add `s` before any Model name, it might look good for `Articles`, `Questions`, `Posts` but for `Quizs`, `Categorys`, `Heros` look too funny than `Quizzes`, `Categories` and `Heroes` respectively. 

```python
# app_name/models.py

from django.db import models

class Article(models.Model):
    title = models.CharField(max_length=127, verbose_name="headline")
    post = models.TextField(verbose_name='content')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now=True)

    def __str__(self):
        return self.title

```

![Verbose Name Attribute](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643219495/blogmedia/v5zphhohi27tvwsa3gsp.png)

Again, we do not need to migrate any changes to the database as it is not a logical change in the schema of the database.

```python 
# app_name/models.py

from django.db import models

class Article(models.Model):
    title = models.CharField(max_length=127, verbose_name="title")
    post = models.TextField(verbose_name='content')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now=True)

    def __str__(self):
        return self.title

    class Meta:
        verbose_name_plural = 'Articless'
```

![Verbose Name Plural](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643219521/blogmedia/sf77v52ic9dheyvv0pvi.png)

Though `Articless` sounds weird, it is just made for realizing that Django by default adds `s` to the Model name for representing as a Class in the Admin. 

### Admin Register Class

We can even list not only the title but many things in a particular format, we need to define a class which will be derived by the `admin.ModelAdmin` class. 

```python
# app_name/admin.py

from django.contrib import admin
from .models import Article

@admin.register(Article)
class Article(admin.ModelAdmin):
    list_display = [
            'title',
            'created',
            'updated',
            ]
```

![Admin-List](https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643217326/blogmedia/y32jmboypbmzyypt68l1.png)

We are using the `admin.ModelAdmin` class as the base class to overwrite the `list_display` list which will display the mentioned items in the Admin Section as a list. We are also using the `admin.register` as the class Decorator to actually register the model in the Admin section. 

Now, we can see a lot of customization in the Admin section. This is just a glimpse of what customization is. This can be tailored as per your model and needs. 

## Conclusion

So, in this section we were able to interact with the Admin section in Django. We were able to register our custom Model in the Admin interface and then customize the format of how it is displayed in the interface. Hopefully, from this part we are able to interact with the data and get our application to test how to plan ahead. In the next section we shall cover how to map up all of these together to create something meaningful and finally understand the Django development process. 

Thank you for reading, if you have any questions or improvements to suggest, please let me know in the comments. I'll be grateful if you you provide a feedback. Happy Coding :)
