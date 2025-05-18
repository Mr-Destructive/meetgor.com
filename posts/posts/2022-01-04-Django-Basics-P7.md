{
  "type": "posts",
  "title": "Django Basics: Creating Models",
  "subtitle": "Understanding and Structuring a Database Schema with Models in Django",
  "date": "2022-01-04 22:15:00 +0530",
  "status": "published",
  "tags": [
    "django",
    "python",
    "web-development"
  ],
  "slug": "django-basics-models",
  "image_url": "https://res.cloudinary.com/dgpxbrwoz/image/upload/v1641315473/blogmedia/dj-7_ixfkka.png",
  "series": [
    "Django-Basics",
    "Django-Series"
  ]
}

## Introduction

We have seen the basics of Django templating in the previous parts of the series. Now, we can move on to the more backend stuff in Django which deals with the Databases, queries, admin section, and so on. In this particular part, we'll cover the fundamental part of any application in Django i.e the `Model`. We'll understand what the model is, how to structure one, how to create relationships and add constraints on the fields, etc. 

## What ate Models?

A model is a Django-way(Pythonic) to structure a database for a given application. It is technically a class that can act as a table in a database generally and inside of the class, the properties of it act as the attributes of that database. It's that simple. Just a blueprint to create a table in a database, don't worry about what and where is our database. We will explore the database and its configuration in the next part. 

By creating a model, you don't have to write all the basic SQL queries like 

```sql
CREATE TABLE NAME(
attrb1_name type,
attrb2_name type,
.
.
.
);
```

If your application is quite big or is complex in terms of the relations among the entities, writing SQL queries manually is a daunting task and also quite repetitive at times. So Django handles all the SQL crap out of the way for the programmer. So Models are just a Pythonic way to create a table for the project/application's database.

## How to create a Model?

Creating a model for an application is as easy as creating a class in python. But hey! It's more than that as there are other questions to address while designing the class. You need to design the database before defining the fields in the model.

OK, we'll it's not straightforward as it seems to but still for creating simple and dummy projects to start with. You can use certain tools like [lucidchart](https://www.lucidchart.com/pages/database-diagram/database-design-tool), [dbdiagrams.io](https://dbdiagram.io/home), and other tools you are comfortable with. It's important to visualize the database schema or the structure of the application before tinkering with the actual database inside the project. Let's not go too crazy and design a simple model to understand the process.

Here's a basic model for a Blog:

```python
#from django.db import models
from django.contrib.auth.models import User

class Article(models.Model):
    title = models.CharField(max_length=255)
    post = models.TextField()
    author = models.ForeignKey(User, on_delete=models.CASCADE, related_name='Article')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now=True)
```   

Ignore the `from django.db import models` as it is already in the file created by Django. If not, please uncomment the line and that should be good to go.
This is a basic model you might wanna play with but don't dump it anywhere. 

We define or create our models in the application inside the project. Inside the application there is already a file called `models.py` just **append** the above code into it. The application can be any application which makes the most sense to you or better create a app if not already created and name it as `article` or `post` or anything you like.

If you are familiar with Python OOP(object-oriented programming), we have basically inherited the `models.Model` class from the `django.db` module into our model.

If you want more such examples, let's see more such models :

An E-Mail application core model. Attributes like `sender`, `subject` of the mail, `body` of the mail, `recipients_list` i.e. the `To:` section in a mail system and the `attachment_file` for a file attachment to a mail if any.

```python
#from django.db import models
from user import EmailUser

class EMail(models.Model):
    sender = models.EmailField(max_length = 255) 
    subject = models.CharField(max_length = 78)
    body = models.CharField(max_length = 40000)
    recipients_list = models.ManyToManyField(EmailUser, related_name = 'mail_list')
    attachment_file = models.FileField(blank=True)
```

A sample model for a note-taking app, consisting of a Note and a Book. A book might be a collection of multiple notes i.e. a single book can have multiple notes so we are using a `ManyToManyField`, what is that? We'll see that shortly.

```python
from django.db import models
from user.models import User

class Notes(models.Model):
    author = models.ForeignKey(User, on_delete=models.CASCADE)
    title = models.CharField(max_length = 1024)
    content = models.Textfield()
    created = models.DateTimeField(auto_now_add = True)
    modified = models.DateTimeField(auto_now = True)
    book = models.ManyToManyField(Book, related_name = 'book')

class Book():
    name = models.CharField(max_length = 1024)
```   

   These are just dummies and are not recommended to use anywhere especially in a serious project. 
So, we have seen a model, but what are these fields and the constraints like `on_delete`, `max_length`, and others in the upcoming section on fields.

## Fields in Django

Fields are technically the attributes of the class which here is the model, but they are further treated as a attribute in a table of a database. So the model becomes a list of attributes which will be then parsed into an actual database. 

By creating attributes inside a class we are defining the structure for a table. We have several types of fields defined already by django for the ease of validating and making a constrained setup for the database schema.

Let's look at some of the types of fields in Django Models.

### Types of Fields

Django has a lot of fields defined in the models class. If you want to go through all the fields, you read through the django docs [field references](https://docs.djangoproject.com/en/4.0/ref/models/fields/#model-field-types). We can access the fields from the `models` module like `name = models.CharField(max_length=10)`, this is a example of defining a attributes `name` which is a CharField. We can set the max_length which acts a constraint to the attribute as we do not want the name field to be greater than 10 and hence parsing the parameter `max_length` to 10. 

We have other field types like:

- `IntegerField` -> for an integer value.
- `TextField` -> for long input of text (like text area in html).
- `EmailField` -> for an single valid email field.
- `DateField` -> for inputting in a date format. 
- `URLField` -> for input a URL field.
- `BooleanField` -> for a boolean value input.

And there are other fields as well which can be used as per requirements.

We also have some other fields which are not directly fields so to speak but are kind of relationship defining fields like:

- `ForeignKey` -> Define a many-to-one relationship to another model/class. 
- `ManyToManyField` -> define a many-to-many relationship to another model/class.
- `OneToOneField` -> define a one to one relationship between different tables/model/class.

So, that's about the field types for just a feel of how to structure or design a database table using a model with some types of attributes. We also need to talk about constraints which needs to added to the fields inside the models. 

### Field Options/Arguments

We can add constraints and pass arguments to the fields in the models. We can add arguments like `null`, `blank`, `defualt`, `choices`, etc. 

- `null=True/False` -> Set a check for the entry in the table as not null in the database.
- `blank=True/False` -> Set a check for the input validation to empty or not.
- `unique=True/False` -> Set a constraint to make the entry unique throughout the table.
- `defualt=anyvalue` -> Set a default value for the field.
- `choices=list` -> Set a list of defined choices to select in the field (a list of two valued tuple).

We also have another constraint specific to the fields like `max_length` for `CharField`, `on_delete` for ForeignKey which can be used as a controller for the model when the related model is deleted, `verbose_name` to set a different name for referencing the entry in the table/model from the admin section compared to the default name of the model, `verbose_name_plural` similar to the `verbose_name` but for referencing the entire table/model. Also `auto_now_add` and `auto_now` for `DateTimeField` so as to set the current date-time by default.

More options and arguments that can be passed to the fields in models are given in the django docs [field options](https://docs.djangoproject.com/en/4.0/topics/db/models/#field-options)

These are some of the options or arguments that we can or need to pass to the fields to set up a constrained schema for our database. 

### Meta class

Meta class is a nested class inside the model class which is most of the times used for ordering the entries(objects) in the table, managing permissions for accessing the model, add constraints to the models related to the attributes/fields inside it, etc.

You can read about the functionalities of the Meta class in the [documentation](https://docs.djangoproject.com/en/4.0/ref/models/options/).

## Model methods

As a class can have functions, so does a model as it is a Python class after all. We can create kind of a helper methods/functions inside the model. The model class provides a helpful `__str__()` function which is used to rename an object from the database. We also have other predefined helper functions like `get_absolute_url` that generates the URL and returns it for further redirection or rendering.

Also, you can define the custom functions that can be used as to help the attributes inside the model class.

## Django ORM

Django has an Object Relational Mapper is the core concept in Django or the component in Django that allows us to interact with the database without the programmer writing SQL/DB queries. It is like a Pythonic way to write and execute sql queries, it basically abstracts away the layer to manually write SQL queries. 

We'll explore the details of how the ORM works under the hood but it's really interesting and fascinating for a Beginner to make web applications without learning SQL(not recommended though personally). For now, its just magical to see Django handling the DB operations for you. You can get the references for learning about the Queryset in ORM from the [docs](https://docs.djangoproject.com/en/4.0/ref/models/querysets/)

## Example Model
Let us set up a model from what we have learned so far. 

We'll create a model for a Blog Post again but with more robust fields and structure. 

```python
#from django.db import models
from django.contrib.auth.models import User

class Article(models.Model):

    options = (
    ('draft', 'Draft'),
    ('published', 'Published'),
    )

    title = models.CharField(max_length=255, unique=True)
    slug = models.SlugField(max_length=255, unique_for_date='publish')
    post = models.TextField()
    author = models.ForeignKey(User, on_delete=models.CASCADE, related_name='Posts')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now=True)
    status = models.CharField(max_length=16, choices=option, default='draft')
    
    def __str__()
        return self.title

    class Meta:
        ordering = ('-publish',)
      
```

   We can see in the above model that we have defined the Meta class which is optional and is generally written to modify how to entries inside the table appear or order with other functionalities as well. We have also added the choices option in the status field which has two choices `Draft` and `Publish` one which is seen by the django interface and the other to the end-users. We have also added certain fields like slug that will create the URL for the blog post, also certain options like `unique` has been set to restrict duplicate entries being posted to the database. The `related_name` in the `ForeignKey` refers to the name given to the relation from the Article model to the User model in this case. 

   So, we can see that Django allows us to structure the schema of a database. Though nothing is seen as an end result, when we configure and migrate the model to our database we will see the results of the hard work spent in creating and designing the model. 

## Database Specific fields

By this time, you will have gotten a feel of what a database might be. Most of the projects are designed around SQL databases but No-SQL databases and others are also used in cases which suite them the most. We have tools to manage this database in SQL we call it the Database Management System (DBMS). It's just a tool to manage data, but there is not just a single Database management tool out there, there are gazillions and bazillions of them. Most  popular include `MySQL`, `PostgreSQL`, `SQLite`, `Oracle`, `Microsoft Access`, `Maria DB`, and tons of others. 

Well, these different DBMS tools are almost similar with a few hiccups here and there. So, different Database tools might have different fields they provide. For Example, in Database `PostgreSQL` provides the ListField which `SQLite` doesn't that can be the decision to be taken before creating any project. There might be some fields that some DBMS provide and other doesn't.

## Conclusion

We understood the basics of creating a model. We didn't touch on the database yet but the next part is all about configuration and migration so we'll get hands-on with the databases. We covered how to structure our database, how to write fields in the model, add constraints and logic to them and explore the terminologies in Django like ORM, Database Types, etc. 

Thank you for reading the article, if you have any feedback kindly let me know, and until then Happy Coding :)
