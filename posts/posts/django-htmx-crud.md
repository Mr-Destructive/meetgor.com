{
  "title": "Django + HTMX CRUD application",
  "status": "published",
  "slug": "django-htmx-crud",
  "date": "2025-04-05T12:33:38Z"
}

<h2>Introduction</h2>
<p>Gone are the days of writing Ajax requests with javascript, just add a few parameters to the HTML content tags and you will be ready for sending requests to your backend. So, we are going back in time and correcting the way we think of APIs and client/server-side rendering. We are referring to the <a href="https://en.wikipedia.org/wiki/Hypermedia">Hypermedia model</a> for levering the server-side processing of data. Let's get our feets wet with this ancient but revolutionary methodology of development with <a href="https://htmx.org/">HTMX</a>.</p>
<p>Yes, HTMX can be used for the API/server-side calls directly in the HTML. We will be exploring the basis of HTMX by creating a basic CRUD application.</p>
<h2>What is HTMX?</h2>
<p>The first question that might come up is what and why HTMX? Htmx is a great library, it's a javascript library but wait. It is a javascript library designed to allow us to write less or no javascript at all. It acts as a way to send AJAX requests without you writing any javascript. It uses native browser features directly from HTML.</p>
<p>So, we can use HTMX to create interactive templates in our Django application. We can dynamically call and fetch data from the server by using simple HTML attributes like <code>hx-get</code>, <code>hx-post</code>, etc. We'll cover those in this article.</p>
<p>You can check the source code used in this article on this <a href="https://github.com/Mr-Destructive/htmx-blog-django">GitHub repository</a>.</p>
<h2>Setup Django Project</h2>
<p>We'll be creating a Django project from scratch and designing a basic blog kind of app. We will be creating a quite simple project with a couple of apps like <code>user</code> for authentication and <code>article</code> for the CRUD part of our blog application.</p>
<p>To set up a django project, we can run the following commands to quickly get up and running with a base django project.</p>
<pre><code>mkdir htmx_blog
python3 -m venv .venv
source .venv/bin/activate
pip install django
django-admin startproject htmx_blog .
</code></pre>
<p>I have a base user model that I use for a simple authentication system in some basic django projects, you can define your own user app or get the app from <a href="https://github.com/Mr-Destructive/django-todo/tree/master/user">here</a>.</p>
<p>So, that being said, we will be using the user model for the article model which we will be defined next. By creating a basic signup functionality, you are good to go!</p>
<h3>Create the Article app</h3>
<p>We will need at least an app to work with htmx as we will define models, views, and URLs later as we configure the htmx.</p>
<pre><code>django-admin startapp article
</code></pre>
<p>After the app has been created, you can add those app labels into the <code>INSTALLED_APPS</code> config in the <code>settings.py</code> file. The <code>user</code> app and the <code>article</code> app need to be added to the installed apps for the django to pick those up for various contexts related to the project.</p>
<pre><code># htmx_blog/settings.py

INSTALLED_APPS = [
    ...
    ...
    ...

    'article',  
    'user',
]
</code></pre>
<p>We are sone with the base setup, we also would require a few more configs for the proper working of the project.</p>
<h3>Setup Templates and Static files</h3>
<p>Templates will play an important role in the htmx part, so it is equally important to configure them properly before dabbling into the htmx and client-side rendering of data.</p>
<p>I like to keep all the templates in a single folder in the <code>BASE_DIR</code> with separate sub-folders for specific apps. Also a single <code>static</code> folder with <code>css</code>, <code>js</code>, and <code>images</code> as the sub-folfers for a larger project.</p>
<pre><code>mkdir templates static
</code></pre>
<p>Further, configure the created static and templates in the settings.</p>
<pre><code class="language-python">
TEMPLATES = [
    {
        'BACKEND': 'django.template.backends.django.DjangoTemplates',
        'DIRS': [os.path.join(BASE_DIR, &quot;templates&quot;)],
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

STATIC_URL = 'static/'
STATICFILES_DIRS = [str(BASE_DIR/ &quot;static&quot;)]
STATIC_ROOT = BASE_DIR / &quot;staticfiles&quot;
</code></pre>
<h3>Initial migration</h3>
<p>Run migration command for the user model and default model in the django project.</p>
<pre><code>python manage.py makemigrations
python manage.py migrate
</code></pre>
<p>So, this project will also include authentication simple registration, and login/logout routes. We will be using the default Django User model by creating an abstract user just in case we require any additional attributes.</p>
<h2>Setup HTMX</h2>
<p>We don't have to configure much for using HTMX as it is a javascript library, we can call it via a CDN or manually install it and link up the static javascript files. Either way, both are equally good, you may like the one I might like the other.</p>
<p>If you already have a base template, you can simply put the below script inside the head tag of the template. This will make us the htmx attributes available.</p>
<pre><code class="language-html">&lt;script src=&quot;https://unpkg.com/htmx.org@1.8.0&quot;&gt;&lt;/script&gt;
</code></pre>
<p>If you don't have a base template, you can create one by creating an HTML file inside the <code>templates</code> directory. The name can be anything but be careful for following up as it might be different for me. I will choose <code>base.html</code> as the template for this project. It will look something like as follows:</p>
<pre><code class="language-html">&lt;!-- tempaltes/base.html --&gt;

&lt;!DOCTYPE html&gt;
&lt;html lang=&quot;en&quot;&gt;
&lt;head&gt;
    &lt;meta charset=&quot;UTF-8&quot;&gt;
    &lt;title&gt;HTMX Blog&lt;/title&gt;
    {% load static %}
    &lt;link rel=&quot;stylesheet&quot; href=&quot;https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css&quot; integrity=&quot;sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm&quot; crossorigin=&quot;anonymous&quot;&gt;
    &lt;script src=&quot;https://unpkg.com/htmx.org@1.8.0&quot;&gt;&lt;/script&gt;
&lt;/head&gt;
&lt;body &gt;
        &lt;nav&gt;
        &lt;h2&gt;HTMX Blog&lt;/h2&gt;
        &lt;div class=&quot;navbar&quot;&gt;
          {% if user.is_authenticated %}
            &lt;a class=&quot;nav-item nav-link&quot; href=&quot;{% url 'logout' %}&quot;&gt;&lt;button class=&quot;btn btn-link&quot;&gt;Logout&lt;/button&gt;&lt;/a&gt;
          {% else %}
            &lt;a class=&quot;nav-item nav-link&quot; href=&quot;{% url 'login' %}&quot;&gt;&lt;button class=&quot;btn btn-link&quot;&gt;Login&lt;/button&gt;&lt;/a&gt;
            &lt;a class=&quot;nav-item nav-link&quot; href=&quot;{% url 'register' %}&quot;&gt;&lt;button class=&quot;btn btn-link&quot;&gt;Register&lt;/button&gt;&lt;/a&gt;
          {% endif %}
        &lt;/div&gt;
        &lt;/nav&gt;

    {% block body %}
    {% endblock %}
&lt;/body&gt;
&lt;/html&gt;
</code></pre>
<p>I have a nav bar with my user authentication views, simply a login or signup button if the user is not logged in and a log-out button if the user is authenticated. We have added the htmx script file from the CDN just before the end of the head tag. We also have included the bootstrap CSS file for a decent UI which we will be creating in this post.</p>
<p>That is one of the ways, htmx can be injected into an HTML template, you can even download the javascript file from the <a href="https://unpkg.com/browse/htmx.org/dist/">htmx cdn</a>. Further, this can be downloaded or pasted into your local folder and served as a static file or embedded directly into an HTML template.</p>
<h2>Defining Models</h2>
<p>We will start the tutorial by defining the model of the application we are creating. Here, we will create a simple Article model with a few parameters like <code>title</code>, <code>content</code>, <code>author</code>, etc.</p>
<pre><code class="language-python">from django.db import models
from user.models import Profile

class Article(models.Model):
    Article_Status = (
        (&quot;DRAFT&quot;, &quot;Draft&quot;),
        (&quot;PUBLISHED&quot;, &quot;Published&quot;),
    )
    title = models.CharField(max_length=128, unique=True)
    content = models.TextField()
    author = models.ForeignKey(Profile, on_delete=models.CASCADE)
    status = models.CharField(
        max_length=16,
        choices=Article_Status,
        default=Article_Status[0],
    )

    def __str__(self):
        return self.title
</code></pre>
<p>In the above model <code>Article</code>, we have a few fields like <code>title</code> simple Character Field, <code>content</code> as a text field as it will be a large text as the post body, <code>author</code> which is a ForeignKey to the User Model. We also have the status, which is defined as a character field but with a few choices like <code>draft</code> or <code>published</code>, we can further modify this status as public or private. But just keeping it simple and easy to understand.</p>
<p>The object reference name for this model is the title as we have defined in the dunder string method. So, that is a simple model created, we can now migrate the changes into the database for adding the tables and attributes.</p>
<pre><code>python manage.py makemigrations
python manage.py migrate
</code></pre>
<p>This will make migrations to the database i.e. convert the python model class into database tables and attributes. So, once the migration process is completed successfully, we can move into the crust of this article which is to actually design the views. In the next section, we will be utilizing the models in our views for representing the data on the templates.</p>
<h2>Creating Article Form</h2>
<p>Before diving into the views section, we need a few things like the Article Form, which will be a Django Model-based form. It will help us a lot in creating or updating the fields for the article model. We can define a form in a python file called <code>forms.py</code>, it's not necessary to keep your forms in the <code>forms.py</code> but if you have a lot of forms and models, it becomes a good practice to organize the components of our app. So, I'll be creating a new file inside of the <code>article</code> app called <code>forms.py</code> and defining the <code>ArticleForm</code>.</p>
<pre><code class="language-python"># article/forms.py

from django import forms
from .models import Article


class ArticleForm(forms.ModelForm):
    class Meta:
        model = Article
        exclude = (
            &quot;created&quot;,
            &quot;updated&quot;,
            &quot;author&quot;,
        )
        widgets = {
            &quot;title&quot;: forms.TextInput(
                attrs={
                    &quot;class&quot;: &quot;form-control&quot;,
                    &quot;style&quot;: &quot;max-width: 450px; align: center;&quot;,
                    &quot;placeholder&quot;: &quot;Title&quot;,
                }
            ),
            &quot;content&quot;: forms.Textarea(
                attrs={
                    &quot;class&quot;: &quot;form-control&quot;,
                    &quot;style&quot;: &quot;max-width: 900px;&quot;,
                    &quot;placeholder&quot;: &quot;Content&quot;,
                }
            ),
        }
</code></pre>
<p>So, the forms are inherited from the [ModelForm] which allows us to create forms based on our model. So, we specify the model name which in our case is <code>Article</code> and further we can have <code>exclude</code> or <code>fields</code> tuples. To exclude certain fields in the actual form, just parse the tuple of those attributes and if you want to only select a few attributes, you can specify the <code>fields</code> tuple and mention the required fields for the form.</p>
<p>So, if we have a lot of things to be included in the form, we can specify only the attributes to be excluded with the <code>exclude</code> tuple. And if we have a lot of fields to be excluded, we can use the <code>fields</code> tuple to specify which attributes to use in the form.</p>
<p>Let's take an example: For the above ArticleForm, if we wanted to specify the required fields to be included in the form, then we might use the <code>fields</code> tuple like below the rest will be not rendered in the form fields.</p>
<pre><code>class ArticleForm(forms.ModelForm):
    class Meta:
        model = Article
        fields = (
            &quot;title&quot;,
            &quot;content&quot;,
            &quot;status&quot;,
        )
</code></pre>
<p>Both of them can be used, it just depends on how many fields you have to exclude or include in the rendered form.</p>
<p>We have also specified the <code>widgets</code> attribute which gives a bit more control on how we need to display the form in a template. So I have specified the type of input it needs to render like a simple text input for the title, text area for content, etc. The cool thing about this is it can automatically set these by knowing the type of field in the model, but sometimes it can be a bit undesired mostly with complex relationships and attributes.</p>
<h2>Creating Views</h2>
<p>Let's start creating views for creating, reading, updating, and deleting articles from the database. I will be using function-based views just because we are understanding the flow of how HTMX and Django can be integrated so we need to dive in deeper and understand the actual flow of the process.</p>
<h3>Create View</h3>
<p>So, creating articles seems like a good way to start off. We can create a simple function-based view which will initially load in an empty <code>ArticleForm</code> and if the request is <code>GET</code> we will render the form in the <code>create.html</code> template. If the request is <code>POST</code> which will be after we submit the form, we will validate the form and attach the current user as the author of the article and save the for instance which will create an article record and this object will be rendered to the detail template.</p>
<pre><code class="language-python">from django.shortcuts import render
from .models import Article
from .forms import ArticleForm

def createArticle(request):
    form = ArticleForm()
    context = {
        'form': form,
    }
    return render(request, 'articles/create.html', context)
</code></pre>
<h4>Rendering the Form</h4>
<p>We are creating an empty instance of <code>ArticleForm</code> and rendering it in the template. So, this will render the empty form in the <code>create.html</code> template.</p>
<pre><code class="language-html">&lt;!-- templates/articles/create.html --&gt;

{% extends 'base.html' %}

{% block body %}
&lt;div hx-target=&quot;this&quot; hx-swap=&quot;outerHTML&quot;&gt;
  &lt;form&gt;
    {% csrf_token %}
    {{ form.as_p }}
    &lt;button hx-post=&quot;.&quot; class=&quot;btn btn-success&quot;
      type=&quot;submit&quot;&gt;Save&lt;/button&gt;
  &lt;/form&gt;
&lt;/div&gt;
{% endblock %}
</code></pre>
<p>Now, here we are inheriting from the base template and creating a form tag in HTML with the <code>{{ form}}</code> for rendering the form fields, we finally have the <code>button</code> element for submitting the form. We have used the <code>hx-post</code> attribute. More on this in just a minute. So, this is we create a template for rendering the article form.</p>
<p>We have used the <code>hx-post</code> attribute here, which will send a <code>POST</code> request to the current <code>URL</code> represented by <code>hx-post=&quot;.&quot;</code>. You might have noticed the <code>div</code> attributes, the <code>hx-target</code> and <code>hx-swap</code>, so these are some of the many attributes provided by the htmx library for controlling the reactivity of the requests made. The <code>hx-target</code> allow us to specify the element or tag to which the data will be rendered. The <code>hx-swap</code> goes hand-in-hand for specifying the target DOM like <code>innerHTML</code>, <code>outerHTML</code>, etc. You can see the various options on the <a href="https://htmx.org/docs/#swapping">htmx docs</a>. By specifying the <code>hx-swap</code> as  <code>outerHTML</code>, we are saying to replace the entire element with the incoming content from the request which we will send with nearby request triggers.</p>
<p>We need to map the view to a URL in order to get a good idea about the request and parsed content.</p>
<p>We'll create a <code>create/</code> route and bind it to the <code>createArticle</code> view with the name <code>article-create</code>.</p>
<pre><code class="language-python"># article/urls.py

from django.urls import path
from . import views

urlpatterns = [
    path('create/', views.createArticle, name='article-create'), 
]
</code></pre>
<p>This URL will be mapped to the global URL in the project, here we can simply specify the prefix for the URLs in the <code>article</code> app and include those URLs.</p>
<pre><code class="language-python"># htmx_blog/urls.py

from django.contrib import admin
from django.urls import path, include

urlpatterns = [
    path('admin/', admin.site.urls),
    path('user/', include('user.urls'), name='auth'),
    path('', include('article.urls'), name='home'),
]
</code></pre>
<p>Feel, free to add any other URL pattern like for instance, the article app is at <code>/</code> i.e. <code>127.0.01.:8000/</code>, you can add any other name like <code>127.0.0.1:8000/article/</code> by adding <code>path('article/', include('article.urls'))</code>.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659252089/blog-media/django-htmx-create-view.png" alt="Django HTMX Create view Form Template"></p>
<p>So, finally, we are sending a <code>GET</code> request to the <code>127.0.0.1:8000/create/</code> and this will output the form. As we have a <code>POST</code> request embedded in the button inside the form, we will send the <code>POST</code> request to the same URL -&gt; <code>127.0.0.1:8000/create/</code>.</p>
<h4>Submitting the Form</h4>
<p>Let's handle the <code>POST</code> request in the create view.</p>
<pre><code class="language-python">from django.shortcuts import render
from .models import Article
from .forms import ArticleForm

def createArticle(request):
    form = ArticleForm(request.POST or None)
    if request.method == 'POST':
        if form.is_valid():
            form.instance.author = request.user
            article = form.save()
            return render(request, 'articles/detail.html', {'article': article})
    context = {
        'form': form,
    }
    return render(request, 'articles/create.html', context)
</code></pre>
<p><strong>Simple explanation</strong></p>
<ul>
<li>Create a form instance of ArticleForm with the request data or empty -&gt; <code>ArticleForm(request.POST or None)</code></li>
<li>If it's a POST request, validate and create the article, render the article object in <code>detail.html</code> template.</li>
<li>If it's a GET request, render the empty form in <code>create.html</code></li>
</ul>
<p>There are a few changes in the view, instead of initializing the form to empty i.e. <code>ArticleForm()</code>, we are initializing with <code>ArticleForm(request.POST or None)</code>. This basically means that if we are having something in the <code>request.POST</code> dict, we will initialize the Form with that data or else an empty form instance.</p>
<p>Next, we check if the request if <code>POST</code>, if it is then we check if the form is valid i.e. the form fields are not empty or if any other constraint on the model attributes is satisfied or not. If the form data is valid, we attach the author as the currently logged-in User/user who sent the request. Finally, we save the form which in turn creates the article record in the database. We then render the created article in the <code>detail.html</code> template which is not yet created.</p>
<p>So, the <code>htmx-post</code> attribute has worked and it will send a post request to the same URL i.e. <code>127.0.0.1:8000/create</code> and this will again trigger the view <code>createArticle</code> this time we will have <code>request.POST</code> data. So, we will validate and save the form.</p>
<h3>Detail View</h3>
<p>The detail view is used for viewing the details of an article. This will be rendered after the article has been created or updated. This is quite simple, we need an <code>id</code> or <code>primary key(pk)</code> of an article and render the <code>title</code> and <code>content</code> of the article in the template.</p>
<p>We pass in a primary key along with the request as a parameter to the view, the <code>pk</code> will be passed via the URL. We fetch the Article object with the id as the parsed <code>pk</code> and finally render the <code>detail.html</code> template with the article object. The <code>context['article']</code> can be accessed from the template to render the specific attributes like <code>title</code>, <code>content</code>, etc.</p>
<pre><code class="language-python"># article/views.py

def detailArticle(request, pk):
    article = Article.objects.get(id=pk)
    context = {'article': article}
    return render(request, 'articles/detail.html', context)

</code></pre>
<p>We can now bind the view to a URL and parse the required parameter <code>pk</code> to the view.</p>
<pre><code class="language-python">from django.urls import path
from . import views

urlpatterns = [
    path('create/', views.createArticle, name='article-create'), 
    path('&lt;int:pk&gt;', views.detailArticle, name='article-detail'), 
]
</code></pre>
<p>We have parsed the <code>pk</code> as <code>int</code> to the URL parameter, so for an article with id=4, the URL will be, <code>127.0.0.1:8000/4/</code>.</p>
<p>We need to create the template for rendering the context from the <code>detailArticle</code> view. So, we create the <code>detail.html</code> in the <code>templates/articles</code> folder. We inherit the base template and render the <code>article.title</code> and the <code>article.content</code> with a linebreaks template filter so as to display the content properly.</p>
<pre><code class="language-html">&lt;!-- templates/articles/detail.html --&gt;


{% extends 'base.html' %}
{% block body %}
&lt;div id=&quot;article-card&quot;&gt;
  &lt;h2&gt;{{ article.title }}
  &lt;p&gt;{{ article.content|linebreaks|safe }}&lt;/p&gt;
&lt;div&gt;
{% endblock %}

</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659252227/blog-media/django-htmx-detail-view.png" alt="Detail View Template"></p>
<p>So, we can now use <code>createArticle</code> view as well as <code>detailArticle</code> view, this both are configured properly, so (CR) or CRUD is completed. We can add <code>listArticle</code> for listing out all the author's(logged-in user) articles.</p>
<h3>List View</h3>
<p>Listview of the articles is much similar to the detail view as it will return a list of articles rather than a single article.</p>
<p>So in the <code>listArticle</code> view, we will return all the articles with the author as the user who sent the request/logged-in user. We will parse this object list into the template as <code>base.html</code> or <code>articles/list.html</code>.</p>
<pre><code class="language-python"># article/views.py


def listArticle(request):
    articles = Article.objects.filter(author=request.user.id)
    context = {
        'articles': articles,
    }
    return render(request, 'base.html', context)
</code></pre>
<p>We will add the URL route for this as the <code>/</code> route that is on <code>127.0.0.1:8000/</code> this is the base URL for the article app and is the route for the <code>listArticle</code> view. So, we will display the list of articles on the homepage.</p>
<pre><code class="language-python"># article/urls.py


from django.urls import path
from . import views

urlpatterns = [
    path('&lt;int:pk&gt;', views.detailArticle, name='article-detail'), 
    path('create/', views.createArticle, name='article-create'), 
    path('', views.listArticle, name='article-list'), 
]
</code></pre>
<p>Let's create the template for the list view which will iterate over the articles and display the relevant data like the title and link to the article.</p>
<pre><code class="language-html">&lt;!-- templates/articles/list.html --&gt;

&lt;ul id=&quot;article-list&quot;&gt;
  {% for article in articles %}
  &lt;li&gt;
    &lt;div class=&quot;card&quot; style=&quot;width: 18rem;&quot;&gt;
      &lt;div class=&quot;card-body&quot;&gt;
        &lt;h5 class=&quot;card-title&quot;&gt;{{ article.title }}&lt;/h5&gt;
        &lt;p class=&quot;card-text&quot;&gt;{{ article.content|truncatewords:5  }}&lt;/p&gt;
        &lt;a href=&quot;{% url 'article-detail' article.id %}&quot; class=&quot;card-link&quot;&gt;Read more&lt;/a&gt;
      &lt;/div&gt;
    &lt;/div&gt;
  &lt;/li&gt;
  {% endfor %}
&lt;/ul&gt;
</code></pre>
<p>We have used the <code>truncatewords:5</code> template filter for only displaying the content of the articles till the first 5 words as it is just a list view, we don't want to display every detail of the article here.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659252293/blog-media/django-htmx-list-view.png" alt="List view Template"></p>
<p>We can use this template to render in the <code>base.html</code> file.</p>
<pre><code class="language-html">&lt;!DOCTYPE html&gt;
&lt;html lang=&quot;en&quot;&gt;
&lt;head&gt;
    &lt;meta charset=&quot;UTF-8&quot;&gt;
    &lt;title&gt;HTMX Blog&lt;/title&gt;
    {% load static %}
    &lt;link rel=&quot;stylesheet&quot; href=&quot;https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css&quot; integrity=&quot;sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm&quot; crossorigin=&quot;anonymous&quot;&gt;
    &lt;script src=&quot;https://unpkg.com/htmx.org@1.8.0&quot;&gt;&lt;/script&gt;
&lt;/head&gt;
&lt;body hx-target=&quot;this&quot; hx-swap=&quot;outerHTML&quot; hx-headers='{&quot;X-CSRFToken&quot;: &quot;{{ csrf_token }}&quot;}'&gt;
        &lt;nav&gt;
        &lt;h2&gt;HTMX Blog&lt;/h2&gt;
        &lt;div class=&quot;navbar&quot;&gt;
          {% if user.is_authenticated %}
            &lt;a class=&quot;nav-item nav-link&quot; href=&quot;{% url 'article-list' %}&quot;&gt;&lt;button class=&quot;btn btn-link&quot;&gt;Home&lt;/button&gt;&lt;/a&gt;
            &lt;a class=&quot;nav-item nav-link&quot; href=&quot;{% url 'logout' %}&quot;&gt;&lt;button class=&quot;btn btn-link&quot;&gt;Logout&lt;/button&gt;&lt;/a&gt;
          {% else %}
            &lt;a class=&quot;nav-item nav-link&quot; href=&quot;{% url 'login' %}&quot;&gt;&lt;button class=&quot;btn btn-link&quot;&gt;Login&lt;/button&gt;&lt;/a&gt;
            &lt;a class=&quot;nav-item nav-link&quot; href=&quot;{% url 'register' %}&quot;&gt;&lt;button class=&quot;btn btn-link&quot;&gt;Register&lt;/button&gt;&lt;/a&gt;
          {% endif %}
        &lt;/div&gt;
        &lt;/nav&gt;

    {% block body %}
    &lt;a href=&quot;{% url 'article-create' %}&quot;&gt;&lt;button class=&quot;btn btn-success&quot; &gt;Create&lt;/button&gt;&lt;/a&gt;
    {% include 'articles/list.html' %}
    {% endblock %}
&lt;/body&gt;
&lt;/html&gt;
</code></pre>
<p>We have now included the <code>list.html</code> template on the homepage and also added the <code>create</code> button as the link to the <code>article-create</code> URL.</p>
<h3>Delete View</h3>
<p>For deleting an article, we will simply rely on htmx for sending the request and on that request, we will delete the current article and render the updated list of articles.</p>
<p>With the <code>deleteArticle</code> view, we will take in two parameters the request which is by default for a Django function-based view, and the primary key as <code>pk</code>. Again we will parse the <code>pk</code> from the URL. We will delete the article object and get the latest list of articles. Finally, render the updated list of articles in the base template which is our list view.</p>
<pre><code class="language-python"># article/views.py


def deleteArticle(request, pk):
    Article.objects.get(id=pk).delete()
    articles = Article.objects.filter(author=request.user)
    context = {'article': articles}
    return render(request, &quot;base.html&quot;, context)

</code></pre>
<p>We will add the <code>deleteArticle</code> into the URL patterns and call it <code>article-delete</code> with the URL of <code>delete/&lt;int:pk&gt;</code>. This will allow us to send a request to the URL <code>127.0.0.1:8000/delete/4</code> for deleting the article with id <code>4</code>.</p>
<pre><code class="language-python"># article/urls.py


from django.urls import path
from . import views

urlpatterns = [
    path('', views.listArticle, name='article-list'), 
    path('&lt;int:pk&gt;', views.detailArticle, name='article-detail'), 
    path('create/', views.createArticle, name='article-create'), 
    path('delete/&lt;int:pk&gt;', views.deleteArticle, name='article-delete'), 
]
</code></pre>
<p>In the delete view, the template is important as we want to send a request appropriately to the defined URL. To do that, we will have a form but it won't have any inputs as such just a button that indicates to delete the current article. We will add the <code>hx-delete</code> attribute as the URL to the <code>deleteArticle</code> view. with the id of the article. This will send a request to the <code>article-delete</code> URL which will, in turn, trigger the view with the given id and delete the article.</p>
<p>We have added the <code>hx-confirm</code> attribute for showing a pop-up of confirmation of deleting the article. As you can see we have added a little script for adding <code>csrf_token</code> into the HTML, this is important in order to submit a form with a valid <code>CSRFToken</code>.</p>
<pre><code class="language-html">&lt;!-- templates/article/delete.html --&gt;

&lt;script&gt;
  document.body.addEventListener('htmx:configRequest', (event) =&gt; {
    event.detail.headers['X-CSRFToken'] = '{{ csrf_token }}';
  })
&lt;/script&gt;
&lt;div &gt;
  &lt;form method=&quot;post&quot; &gt;
  {% csrf_token %}
    &lt;button class=&quot;btn btn-danger&quot;
      hx-delete=&quot;{% url 'article-delete' article.id %}&quot;
      hx-confirm=&quot;Are you sure, You want to delete this article?&quot;
      type=&quot;submit&quot;&gt;
      Delete
    &lt;/button&gt;
  &lt;/form&gt;
&lt;/div&gt;
</code></pre>
<p>Do you have a question like how do we access the <code>article.id</code>? we are not rendering the <code>delete.html</code> template from the view, so there is no context to pass. We will include this snippet into the detail view template, so as to have the option of deleting the current article.</p>
<p>We will modify the <code>articles/detail.html</code> template and include the <code>delete.html</code> template. This includes simply adding an HTML template in the specified location. So, we will basically inject the delete form into the detail template.</p>
<pre><code class="language-html">{% extends 'base.html' %}
{% block body %}
&lt;div hx-target=&quot;this&quot; hx-swap=&quot;outerHTML&quot;&gt;
  &lt;h2&gt;{{ article.title }}
  {% include 'articles/delete.html' %}
  &lt;p&gt;{{ article.content|linebreaks|safe }}&lt;/p&gt;
&lt;div&gt;
{% endblock %}
</code></pre>
<p>Hence, we will have a nice option to delete the article in the detail section, this can be placed anywhere but remember, we need to add the <code>hx-target=&quot;this&quot;</code> and <code>hx-swap=&quot;outerHTML&quot;</code> in the div so as to correctly swap the HTML content after the request has been made.</p>
<h3>Update View</h3>
<p>We can now move into the final piece of the CRUD i.e. <code>Update</code>. This will be similar to the <code>createArticle</code> with a couple of changes. We will parse parameters like <code>pk</code> to this view as well because we want to update a specific article. So, we will have to get the primary key of the article from the URL slug.</p>
<p>Inside the <code>updateArticle</code> view, we will first grab the article object from the parsed primary key. We will have two kinds of requests here, one will be for fetching the <code>form</code> with the current article data, and the next request will be the <code>PUT</code> request for actually saving the changes in the article.</p>
<p>The first request is simple as we need to parse the form data with the instance of the article object. We will call the <code>ArticleForm</code> with the instance of <code>article</code> this will load the data of the article into the form ready to render into the template. So once the <code>GET</code> request has been sent, we will render the template with the form pre-filled with the values of the article attributes.</p>
<pre><code class="language-python"># article/views.py


def updateArticle(request, pk):
   article = Article.objects.get(id=pk)
   form = ArticleForm(instance=article)
   context = {
       'form': form,
       'article': article,
   }
   return render(request, 'articles/update.html', context)
</code></pre>
<p>We will create a template in the <code>templates/articles/</code> folder as  <code>update.html</code> which will have a simple form for rendering the form fields and a button for sending a <code>PUT</code> request. We will render the <code>form</code> and then add a button element with the attribute <code>hx-put</code> for sending the <code>PUT</code> request to save changes to the article record. We will parse in the <code>article.id</code> for the primary key parameter to the view.</p>
<pre><code class="language-html">&lt;!-- templates/articles/update.html --&gt;

&lt;div hx-target=&quot;this&quot; hx-swap=&quot;outerHTML&quot;&gt;
  &lt;form&gt;
    {% csrf_token %}
    {{ form.as_p }}
    &lt;button hx-put=&quot;{% url 'article-update' article.id %}&quot;
      type=&quot;submit&quot;&gt;Update&lt;/button&gt;
  &lt;/form&gt;
&lt;/div&gt;
</code></pre>
<p>We are yet to link the <code>updateArticle</code> into the URLs. We will add the view <code>updateArticle</code> into the URLs with the name as <code>article-update</code> and <code>update/&lt;int:pk</code> as the slug pattern. This URL pattern will trigger the <code>updateArticle</code> when we send an HTTP request to the <code>127.0.0.1:8000/update/4</code> for updating the article with id as <code>4</code>.</p>
<pre><code class="language-python"># article/urls.py


from django.urls import path
from . import views

urlpatterns = [
    path('', views.listArticle, name='article-list'), 
    path('&lt;int:pk&gt;', views.detailArticle, name='article-detail'), 
    path('create/', views.createArticle, name='article-create'), 
    path('delete/&lt;int:pk&gt;', views.deleteArticle, name='article-delete'), 
    path('update/&lt;int:pk&gt;', views.updateArticle, name='article-update'), 
]
</code></pre>
<p>This is not done yet, we will need to handle the <code>PUT</code> request as well i.e. when the form details have been modified and we are about to save changes to the form data. So, we will check for the request method's type. If it is a <code>PUT</code> request, we will have to process a few things.</p>
<pre><code class="language-python"># article/views.py


from django.http import QueryDict

def updateArticle(request, pk):
    article = Article.objects.get(id=pk)
    if request.method == 'PUT':
        qd = QueryDict(request.body)
        form = ArticleForm(instance=article, data=qd)
        if form.is_valid():
            article = form.save()
            return render(request, 'articles/detail.html', {'article': article})
    form = ArticleForm(instance=article)
    context = {
        'form': form,
        'article': article,
    }
    return render(request, 'articles/update.html', context)
</code></pre>
<p>In the above <code>updateArticle</code> view, we have to check for a <code>PUT</code> request, if we are sending a <code>PUT</code> request, the form instance needs to be loaded from the request object. We use the <code>request.body</code> to access the data sent in the request. The incoming data received from the <code>request.body</code> object is not a valid format to parse it to the form instance, so we will parse it using <code>QueryDict</code>. This will allow us to modify the <code>request.body</code> object into valid python serializable data.</p>
<p>So, we import the <code>QueryDict</code> from <code>django.http</code> module. We parse the data as the parameter to <code>QueryDict</code> and store it in a variable. We then have to get the <code>ArticleForm</code> for fetching the data as per the form details, so we parse the instance and also the <code>data</code> parameter. The instance is the article object and the data is the received form data which we have stored in <code>qd</code> as <code>QueryDict(request.body)</code>. This will load the new form data and then we can validate it the form.</p>
<p>After we have verified the form details, we can save the form and this will update the article record. Thereby we can render the updated article in the <code>detail</code> view with the updated <code>article</code> object as the context.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659252091/blog-media/django-htmx-update-view.png" alt="Update View Form Template"></p>
<p>So, this will set up the update view as well, we can now create, read, update, and delete an article instance with HTMX in templates and Django function-based views without writing any javascript.</p>
<h2>Summary</h2>
<p>We were able to create a basic CRUD application in Django with HTMX. We used simple function-based views to demonstrate the inner details of how we can work with HTMX and handle requests from the templates. By creating simple standalone templates, we can connect those together to make a fully functional and responsive webpage. The UI is not great but the purpose of this tutorial was to make a barebone CRUD app to work with the backend using HTMX, so hopefully, you would have got a good overview of how HTMX can be integrated into a Django application.</p>
<p>Overall HTMX is a great library that can be used to enhance or even create a new web application for making the site responsive and without writing any javascript.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1659252296/blog-media/django-htmx-demo.gif" alt="Django HTMX CRUD Application Demo GIF"></p>
<p>You can check out the source code for this project and blog on the <a href="https://github.com/Mr-Destructive/htmx-blog-django">htmx-blog GitHub</a> repository.</p>
<h2>Conclusion</h2>
<p>From this post, we were able to understand the basics of HTMX and how we can integrate it into a Django application. Hopefully, you enjoyed the post, if you have any queries or feedback, please let me know in the comments or on my social handles. Thank you for reading. Happy Coding :)</p>
