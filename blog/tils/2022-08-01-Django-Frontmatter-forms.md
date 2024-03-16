---
templateKey: til 
title: "Django Blog DevLog: Load Frontmatter data into Template/Model Form Fields"
description: "Rendering frontatter from content field into the Template Form field using HTMX and frontmatter libraries"
date: 2022-08-01 20:30:00
status: published-til
slug: django-form-load-frontmatter
tags: ['django', 'htmx','python']
---

## Introduction

I have an Article Form where I load my post into it directly, it might have frontmatter. So what I wish to achieve is when I paste in the content, the frontmatter should be picked up and it should render the form fields like `title`, `description`, and then also remove the frontmatter from the content.

To do that, we will require a model to work with and a form based on that model. We will exclude a few fields from that model so as to process these attributes on the server side. I am working on my Blog project which is a simple Django application.  You can get the source code for the project on the [GitHub repository](https://github.com/mr-destructive/techstructive-blog/).

## Django Project Context

The techstructive-blog is a django project, which has a couple of applications currently, not in a good situation. There are apps like `article`, `blog`, and `user`. This project has templates and static folder in the base directory. The project is deployed on [railway](https://www.railway.app) this is an always under development project, you can check out the [techstructive-blog](https://django-blog.up.railway.app). We can get the bits and pieces of the project details required for understanding what I want to do with the following sections.

### Article Model

We have an `Article` model with attributes like `title`, `description`,  `author` as a Foreign Key to the user model, and a few other attributes which is not related to what we are trying to achieve right now or at least don't require an explanation. We have a model method called `get_absolute_url` for getting a URL in order to redirect the client when the model instance is created or updated from the backend. You can definitely check out the details of these small components or templates in the project repository. 

```python
# articles/models.py


class Article(TimeStampedModel):
    class Article_Status(models.TextChoices):
        DRAFT = (
            "DRAFT",
            _("Draft"),
        )
        PUBLISHED = (
            "PUBLISHED",
            _("Published"),
        )

    title = models.CharField(max_length=128)
    description = models.CharField(max_length=256)
    content = models.TextField(default="", null=False, blank=False)
    status = models.CharField(
        max_length=16,
        choices=Article_Status.choices,
        default=Article_Status.DRAFT,
    )
    blog = models.ForeignKey(Blog, on_delete=models.CASCADE, null=True, blank=True)
    author = models.ForeignKey(User, on_delete=models.CASCADE, related_name="author")

    def __str__(self):
        return self.title

    def get_absolute_url(self):      
        return reverse('articles:article-detail', args=[str(self.id)])
```

In the below snippet, we have the forms defined in the article application for creating or updating of article instance.  We will be using model forms as our form data should contain fields related to a model in this case the `Article` model. So when we inherit the `forms.ModelForm` in our custom `ArticleForm` we simply need to specify the model and it will take in all the attributes of that model by default, but if we specify the `fields` or `exclude` tuples, it will include only or exclude only the provided list of attributes from the model. 

We have also added the widgets for the form fields which will allow us to customize the way the fields in the template/form will render. We can specify the HTML attributes like `width`, `height`, `style`, etc.  

### Article Form

```python
# article/forms.py


from django import forms
from .models import Article


class ArticleForm(forms.ModelForm):
    class Meta:
        model = Article
        exclude = (
            "created",
            "updated",
            "author",
        )
        widgets = {
            "title": forms.TextInput(
                attrs={
                    "class": "form-control",
                    "style": "max-width: 450px; align: center;",
                    "placeholder": "Title",
                }
            ),
            "description": forms.TextInput(
                attrs={
                    "class": "form-control",
                    "style": "max-width: 900px; ",
                    "placeholder": "Description",
                }
            ),
            "content": forms.Textarea(
                attrs={
                    "class": "form-control post-body",
                    "id": "text-content",
                    "style": "max-width:900px;",
                    "hx-post": "/article/meta/",
                    "placeholder": "Content",
                }
            ),
            "blog": forms.Select(
                attrs={
                    "class": "form-control",
                    "placeholder": "Blog Publication",
                }
            ),
        }

```

So, these are my models and form files in the article app. Using htmx and without any javascript I want to update the form so that it picks up the front matter in the content field which is a text area and fills the title, description other attributes automatically for me. 

This can be done in a lot of ways, but I will be sharing one of the ways that I recently used in my blog project. This process involves creating a class-based view and adding a `POST` method that won't post any data to the backend but will send necessary data to the view.


Let's see the process before diving into any of the code:

## Gist of the Process

- Attach a `hx-post` attribute to the form field for sending the request to a view
- When the request is sent, the data is loaded with `request.POST`, it is cleaned and converted in python-readable format with json.
- Once we have the data, we try to use the `frontmatter.loads` function that will load the content and if we have a frontmatter in the text, it will load it as a `frontmatter.POST` object.
- We will extract `title`, `description`, and other data fields from the object.
- We will initialize a Form instance of Article, with the initial data values as the extracted data from the frontmatter.
- Now we have two options:
    - If the article instance already exists i.e. we are updating the article
   - Else we are creating a new article

Accordingly, we will load the form in the respective templates i.e. `update.html` for updating the existing articles and `article-form.html` for a new article.

## Adding HTMX Magic

If you'd have seen we have a `hx-post` attribute in the `article/forms.py` file, the `content` widget has a `hx-post` attribute which sends a post request to the `article/meta/` URL route. This route we will bind to the `ArticleMetaView` which we will define in a few moments. This is usually sent once we change a certain text in the content field, so we can modify it as per your requirement with `hx-trigger` that can enable us to specify the trigger event or the type of trigger we want. Further, you can read from the [htmx docs](https://htmx.org/docs/#trigger-modifiers) about these triggers and other attributes. 

```python
# article/urls.py

from django.urls import path
from . import views

app_name = "articles"

urlpatterns = [
    path("", views.ArticleCreateView.as_view(), name="article-create"),
    path("<int:pk>/", views.ArticleDetailView.as_view(), name="article-detail"),
    path("delete/<int:pk>/", views.ArticleDetailView.as_view(), name="article-delete"),
    path("edit/<int:pk>", views.ArticleDetailView.as_view(), name="article-update"),

    # the new view that we will create
    path("meta/", views.ArticleMetaView.as_view(), name="article-meta"),
]
```

## Capture Frontmatter Meta-data View 

Along with the Create, Detail/List, Update, Delete View, I will create a separate class called `ArticleMetaView` that will fetch the form fields and render the templates again but this time it will fill in the frontmatter meta-data in the fields if the content is parsed with the relvant frontmatter.

```python
# articles/view.py

class ArticleMetaView(View):
    model = Article

    def post(self, request, *args, **kwargs):
        
        data = json.loads(json.dumps(dict(request.POST)))
        loaded_frontmatter = frontmatter.loads(data['content'][0])

       # frontmatter has keys i.e. attributes like title, description, etc.
        if dict(loaded_frontmatter):
            article_title = loaded_frontmatter['title']
            article_description = loaded_frontmatter['description']
            form = ArticleForm(initial={'title': article_title, 
            'description': article_description, 'content': loaded_frontmatter.content})
            context = {'form': form}
            article_list = Article.objects.filter(title=article_title)
            if article_list:
                article = article_list.last()
                context['article'] = article
                return render(request, 'articles/edit_article.html', context)
            return render(request, 'articles/article_form.html', context)

        article_list = Article.objects.filter(title=data['title'][0])
       
       # if the article title has been already taken i.e. we are updating an article

        if article_list:
            article = article_list.last()
            form = ArticleForm(data=request.POST)
            context = {'form': form}
            context['article'] = article
            return render(request, 'articles/edit_article.html', context)

        form = ArticleForm(data=request.POST)
        context = {'form': form}
        return render(request, 'articles/article_form.html', context)

```

In the above `ArticleMetaView` we have created a `post` method as we want to get hold of the content from the form. So, we start by extracting and converting the `request.body` data into an appropriate type for easily working with python. So, the `request.body` will contain the data like `csrf_token`, `form_data`, etc. received from the frontend template. We store the received data as `data` and now from this data, we can load the content field which will have the content information.

Firstly we will extract the `request.body` which will contain the data from the form as we have made a `POST` request to this endpoint. For doing that we need to parse the content in a apropriate format such that it is python friendly. So we wrap the `request.body` into json format and then decode it back into the json string. This will give us the dict of the request data.

```python
data = json.loads(json.dumps(dict(request.POST)))
```

```
{'csrfmiddlewaretoken': ['bSYJxD39XH509tD1tZGd0WU21PUaKaLeqjjGbyzRvLXF4P8iIxb5l0fmTWVFjELQ'], 'title': ['test2'], 'description': ['test'], 'content': ['test something'], 'status': ['DRAFT'], 'blog': ['']}
```

So, this will grab the request data as a dict, we can then extract the data from this as it has data from the Form fields. We are interested in the content field in the Form, so we can get it by specifying the key `content` from the extracted data. But as we can see the data doesn't contain the actual data instead it is wrapped in a list i.e. `['test something']`, so we will have to index it and then fetch it.

```python
content_string = data['content'][0]
```

This will give us the exact content field as a string. So, we can now move into extracting the frontmatter from the fields. 

Now, we can use the [frontmatter](https://python-frontmatter.readthedocs.io/en/latest/index.html) library to parse the content into the [loads](https://python-frontmatter.readthedocs.io/en/latest/api.html#frontmatter.loads) funciton and extract the frontmatter if it is present in the content field. The frontmatter library has a `loads` function which takes in a string and can give out a [frontmatter.Post](https://python-frontmatter.readthedocs.io/en/latest/api.html#post-objects) object. The loads function is differnet from the [load](https://python-frontmatter.readthedocs.io/en/latest/api.html#frontmatter.load) function as the load frunciton is for reading data from a stream of bytes i.e. a file or othe related byte object. The differnece is subtle but it took a read at the [documentation](https://python-frontmatter.readthedocs.io/en/latest/api.html#module-frontmatter).

```python
post = data['content'][0]
loaded_frontmatter = frontmatter.loads(post)
```

This wil load the content and give us a `frontmatter.Post` as said earlier. This will contain a dict with all the frontmatter if it has any and will by default parse the non-frontmatter data i.e. the remaining text into the `content` key. We need a chack if the Form field had any fronmatter this can be checked by the `dict(loaded_frontmatter)` which will return None if it cannot load the frontmatter.

```python
loaded_frontmatter = frontmatter.loads(data['content'][0])
if dict(loaded_frontmatter):
  print(loaded_frontmatter.keys())
```

```
dict_keys(['templateKey', 'title', 'description', 'date', 'status'])
```

So once we have the frontmatter loaded we can get specific keys from it and initialize the form vaues to them. But we have made clear distictions that we want to perform a specific task if we have frontmatter keys in the content field of the Form else we can do something else.

First let's handle the loading of the frontmatter into the form. For doing that we will get all the required attributes from the frontmatter like `title`, `description`, `content`, etc which can be accessed normally as we extract the value from a key in a dict.

Once we have got those keys, we can start filling in the Form data with initial values. The [Django Model form](https://docs.djangoproject.com/en/4.0/topics/forms/modelforms/) takes in a parameter like [initial](https://docs.djangoproject.com/en/4.0/topics/forms/modelforms/#providing-initial-values) which can be a dict of the fiields along with the value that can be used for rendering the form initially when we load the template.

```python
article_title = loaded_frontmatter['title']
article_description = loaded_frontmatter['description']

form = ArticleForm(initial={'title': article_title, 'description': article_description, 'content': loaded_frontmatter.content})
```

This will take in a `ArticleForm` and fill the initial values like `title`, `description`, etc which we have provided in the dict with the values. Now, we need to parse this form in the current template or re-render the template. But before that, we need to parse this context into the template. We will create a dict with `form` as the key which can be used to render in the template.

Also, we have a two ways here, either the user is creating a new article or it is updating a existing article. We need to make sure that we preserve the initial fields in the form as we are updating the existing article. So, we can filter the article objects as per the title of the current title and then if we find a article with that title, we will parse the context with that article object.

```python
article_list = Article.objects.filter(title=article_title)
if article_list:
    article = article_list.last()
    context{
      'form': form,
      'article': article
    }
    return render(request, 'articles/edit_article.html', context)
context = {'form': form}
return render(request, 'articles/article_form.html', context)
```

Now, we have form data along with the article instance used for rendering the form with appropriate content. So, this will work for editing an already existing article. For a new article, we have to simply parse the form to the template and it will render the title picked from the fotnmatter or leave it empty.

Similarly, for the article with no frontmatter we will iterate over the article and if the article's title already exist, we will render the article data with the form else render the form with the parsed title and other meta-data in the form.


<video width="800" height="450" controls>
  <source src="https://res.cloudinary.com/techstructive-blog/video/upload/v1659370006/blog-media/frontmatter-load-htmx.mp4" type="video/mp4">
</video>

So that is how we render the form data with frontmatter into appropriate meta-data in the form. We have used Django forms and make use of HTMX for the dynamic updation of form.
