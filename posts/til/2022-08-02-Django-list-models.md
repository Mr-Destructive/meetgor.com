---
type: til 
title: "Django: Get list of all models and associated fields in a django project"
description: "Get the list of all the models and associated fields/attributes in a django project or an application"
date: 2022-08-02 15:30:00
status: published
slug: django-list-models
tags: ['django', 'python']
---

## Context

Let's say we want the list of all the models and associated attributes in all the applications of a django project, we can do that using the [django.apps](https://docs.djangoproject.com/en/4.0/ref/applications/) with apps method. 

## Get all the models in a project

To fetch all the models, we can use the [get_models](https://docs.djangoproject.com/en/4.0/ref/applications/#django.apps.AppConfig.get_models) methods, it will return a list of model classes in all the entire project(all applications). We can import all the models in the django project with the command:
 
```python
from django.apps import apps
models = apps.get_models()
```

```
[<class 'django.contrib.admin.models.LogEntry'>, <class 'django.contrib.auth.models.Permission'>, 
<class 'django.contrib.auth.models.Group'>, <class 'django.contrib.contenttypes.models.ContentType'>,
 <class 'django.contrib.sessions.models.Session'>, <class 'allauth.account.models.EmailAddress'>, 
<class 'allauth.account.models.EmailConfirmation'>, <class 'allauth.socialaccount.models.SocialApp'>, 
<class 'allauth.socialaccount.models.SocialAccount'>, <class 'allauth.socialaccount.models.SocialToken'>, 
<class 'user.models.TimeStampedModel'>, <class 'user.models.User'>, <class 'articles.models.Tags'>,
 <class 'articles.models.Series'>, <class 'articles.models.Article'>, <class 'blog.models.Blog'>]
```

We are importing the apps and creating a list of the models in our django project. The Django app command will load all the applications in the project, and the [get_models]() method will fetch the associated models. This has resulted in a list of model class objects, we can iterate over them and fetch the required details, we want.

For instance, If I am interested in the name of these models, I can use the `__name__` property to fetch the model's name. 

```python
from django.apps import apps
model_list = apps.get_models()

for model in model_list:
    print(model.__name__)
```

```
LogEntry
Permission
Group
ContentType
Session
EmailAddress
EmailConfirmation
SocialApp
SocialAccount
SocialToken
TimeStampedModel
User
Tags
Series
Article
Blog

```

So, from the above example, we can see we have accessed all the model names in our entire django project. 

## Access Application name associated with a model

For accessing the name of the application from the model class, we can use the `_meta` attribute followed by the `app_label` property to get the `app_name` associated with the model.

```python
from django.apps import apps
model_list = apps.get_models()

for model in model_list:
    print(f"{model._meta.app_label}  -> {model.__name__}")
```

```
admin  -> LogEntry
auth  -> Permission
auth  -> Group
contenttypes  -> ContentType
sessions  -> Session
account  -> EmailAddress
account  -> EmailConfirmation
socialaccount  -> SocialApp
socialaccount  -> SocialAccount
socialaccount  -> SocialToken
user  -> TimeStampedModel
user  -> User
articles  -> Tags
articles  -> Series
articles  -> Article
blog  -> Blog
```

In the above example, we can see we have printed all the models with their associated application names. 

## Accessing all the attributes associated with a model
 
To access all the fields/property/attributes associated with a model, we can again use the `_meta` attribute followed by the `get_fields` method.  This method will return a list of field objects. For accessing the name of those attributes/fields, we have to iterate over the list and then further use `name` property.

```python
from django.apps import apps
model_list = apps.get_models()
for model in model_list:
    print(model.__name__)
    field_list = model._meta.get_fields()
    for field in field_list:
        print(field.name)
```

```
LogEntry
id
action_time
user
content_type
object_id
object_repr
action_flag
change_message
Permission
group
user
id
...
...
Blog
article
id
name
description
authors
```
So, that is how we get all the associated field names in the associated models in our django projects. Also, there are a lot of attributes, we can access with the apps property. The `__dict__.keys()` can be used to get the list of all associated properties or other methods in a class instance.

```
>>> m[14]._meta.get_fields()[4].__dict__.keys()

dict_keys(['name', 'verbose_name', '_verbose_name', 'primary_key', 'max_length', '_unique', 'blank', 'null', 'remote_field',
 'is_relation', 'default', 'editable', 'serialize', 'unique_for_date', 'unique_for_month', 'unique_for_year', 'choices', 
'help_text', 'db_index', 'db_column', '_db_tablespace', 'auto_created', 'creation_counter', '_validators', '_error_messages', 
'error_messages', 'db_collation', 'validators', 'attname', 'column', 'concrete', 'model'])
```
In the above example, I am using a list of models and getting the list of all the attributes associated with a field of a model. This can be applied and other properties can be accessed. 

## Get Models with a specific app

Let's say we want all the models associated with a particular application in the project, we can do that by specifying the name of the application.

```python
from django.apps import apps

app_info = apps.get_app_config('articles')

print(app_info.__dict__.keys()

print(app_info.verbose_name)

print(app_info.models)

print(app_info.models['article'].__dict__.keys())
```

```
dict_keys(['name', 'module', 'apps', 'label', 'verbose_name', 'path', 'models_module', 'models'])

'Article'

{'tags': <class 'articles.models.Tags'>, 'series': <class 'articles.models.Series'>, 'article': <class 'articles.models.Article'>}

dict_keys(['__module__', 'Article_Status', '__str__', 'get_absolute_url', '__doc__', '_meta', 'DoesNotExist', 'MultipleObjectsReturned', 'title', 'description', 'content', 'status', 'get_status_display', 'blog_id', 'blog', 'author_id', 'author', 'timestampedmodel_ptr_id', 'timestampedmodel_ptr'])
```

So, we can see that we have got the information about the app `articles` in the proejct where we can get the `verbose_name` property to fetch the human-readable format of the article model. Further, we can get all the models associated with the `articles` application. We get back a dict with the model name as the key and the class reference as the value.

We have accessed the `article` model in the `articles` application and fetched all the associated properties or methods in the model.

For further references, you can visit the [django apps documentation](https://docs.djangoproject.com/en/4.0/ref/applications/) to get more relevant methods and properties.
