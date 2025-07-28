{
  "title": "Django Basics: CRUD API",
  "status": "published",
  "slug": "django-basics-crud-api",
  "date": "2025-04-05T12:33:48Z"
}

<h2>Introduction</h2>
<p>After understanding the admin section and setting up a database, we can finally move on to the CRUD API in Django. We will create a API in pure Django with Model Form, Class Based Views and templates. We will go in two parts in this section, first revising the workflow of the app creation and setup, the next part all about CRUD API. We'll dive into a lot of familiar topics and concepts from previous parts but there are is still a lot to learn beyond the basics.</p>
<p>The following is the demonstration of the CRUD API we will be making in this section:</p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643985336/blogmedia/hoxf3u9a872uvcbyehov.gif" alt="CRUD API- demonstration"></p>
<h2>Quickly Setting up an app</h2>
<p>We have seen how to create an app in the <a href="https://mr-destructive.github.io/techstructive-blog/django/python/web-development/2021/11/24/Django-Basics-P3.html">folder structure part</a> of the series, after that we have created all the components namely the views, urls, templates, etc in distinct parts. In this article, we will create an app together in a single part. We'll discuss all the process in short because its important to get the workflow of Django to ease the actual coding process. Let's dive in the part one of this section of creating a CRUD API i.e. to set up an app.</p>
<h3>Creating App</h3>
<p>To quickly set up an application, we need to execute an command so that python under the hood creates a folder for an app in the current Django project.</p>
<pre><code>python manage.py startapp api 
</code></pre>
<p>This will set up an folder <code>api</code> in the base directory of current django project. Now, we'll need to configure the Django settings for the project to pick up this app while running the server, making any migrations or any other project level process.</p>
<h3>Registering the App-name in settings</h3>
<p>Simply append the name of the app in a string in the <code>INSTALLED_APPS</code> list inside the <code>settings.py</code> file.</p>
<pre><code class="language-python"># project_name / settings.py

# Application definition

INSTALLED_APPS = [
    'django.contrib.admin',
    'django.contrib.auth',
    'django.contrib.contenttypes',
    'django.contrib.sessions',
    'django.contrib.messages',
    'django.contrib.staticfiles',

    'api',
]
</code></pre>
<h3>Setting up the URLS</h3>
<p>We need to create a <code>urls.py</code> file for the <code>api</code> app and link it in the main project's URLs.</p>
<pre><code class="language-python"># app-name / urls.py

from django.urls import path

urlpatterns = [

]
</code></pre>
<p>We'll keep the url patterns empty but after setting and configuring the project and the application, we'll map the routes to the views in this app.</p>
<p>Next, we need to update the <code>urls.py</code> file in the project folder to include the <code>api</code> routes/urls. I have kept it '' or base route, it could be anything as per your application design.</p>
<pre><code class="language-python"># project_name / urls.py

from django.contrib import admin
from django.urls import path, include

urlpatterns = [
    path('admin/', admin.site.urls),
    path('api/', include('app_name.urls')),
]
</code></pre>
<p>After configuring the URLs we need to set up the templates and static files.</p>
<h3>Setting up Templates and Static files</h3>
<p>To set up the templates, we need to configure the <code>settings.py</code> file to look for the templates in a specific directory. Below is the snippet to change the default configuration.</p>
<pre><code class="language-python"># project_name / settings.py

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
</code></pre>
<p>So, you can now create your templates (HTML documents) inside the <code>templates</code> folder after creating the folder in the base project directory.</p>
<p>TO configure static files, we need to also make modifications to the <code>settings.py</code> file as follows:</p>
<pre><code class="language-python"># project_name / settings.py
# import os

STATIC_URL = '/static/'
STATICFILES_DIRS = [os.path.join(BASE_DIR, &quot;static&quot;)]
STATIC_ROOT  = os.path.join(BASE_DIR, 'staticfiles')
</code></pre>
<p>Here, we are configuring the static files ( CSS, Js, Assets) in the <code>static</code> folder. This can be anything you like but the folder name should be then changed accordingly. Similar to the Templates folder, the static folder is also located in the root directory of the project.</p>
<p>This is all the necessary configuration for simple full stack application you can make, still we have to configure the backend and write the actual logic for the application. We'll move on to the second part of this section i.e. creating the CRUD API.</p>
<h2>Making an CRUD API</h2>
<p>There are certain things to be planned before we move ahead like the selection of database, the schema of the database and basically the structure of the backend. For now, we'll go with SQL database, a REST API with PostgreSQL hosted locally on the system.</p>
<h3>Configure the database</h3>
<p>To configure the database, we simply need to first create the database. This can be done by using a CLI or the Admin interface of the particular database we are working with in our case it's PostgreSQL. Postgres comes with pgAdmin to create and manage the databases and the server locally. The detailed explanation of the <a href="https://mr-destructive.github.io/techstructive-blog/django/python/web-development/2022/01/16/Django-Basics-P8.html">database creation</a> is explained in the previous parts of the series.</p>
<p>After creating the database locally, you need to tweak the <code>settings.py</code> file database configuration object as per your credentials of the database instance.</p>
<pre><code class="language-python"># project_name / settings.py

DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.postgresql',
        'NAME': 'blogapp',
        'USER': 'postgres',
        'PASSWORD': '@1234567',
        'HOST': 'localhost',
        'PORT': '5432',
    }
}
</code></pre>
<p>If you have already created the databse, there is no problem in reusing the same one but for production level applications make sure to keep things separate. Also the model name has to be distinct in each app to create tables in the database.</p>
<h3>Creating Model</h3>
<p>We can now finally move on to the logic of the application, this is the part which acts as the backbone of backend development in django. You need to design the model carefully and programmatically. Make sure to include most of the logic as you can inside the models as it is a standard and a good practise to develop in professional projects.</p>
<p>For this app, we simply are going to create a blog post model. We are going to have some simple attributes and structure.</p>
<pre><code class="language-python"># app_name / models.py

from django.db import models
from django.contrib.auth.models import User

class Article(models.Model):
    title = models.CharField(max_length=127, verbose_name=&quot;headling&quot;)
    post = models.TextField(verbose_name='content')
    author = models.ForeignKey(User, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now=True)

    def __str__(self):
        return self.title

    class Meta:
        verbose_name_plural = 'Articles'
</code></pre>
<p>This is again a simple model for demonstration of a lot of things like the <code>str</code> function, <code>Meta</code> class, <code>verbose_names</code> and so on. You can design models as per your requirements and proper planning. It really determines the structure of an application.</p>
<p>Now, after we have created the models, we need to migrate the schema into the database. Remember it acts as a checkpoint in the history of changes to the database.</p>
<pre><code>python manage.py makemigrations
</code></pre>
<p>This will create a checkpoint in the <code>migrations</code> folder in the app.</p>
<p>To make changes to the database, we use the migrate command.</p>
<pre><code>python manage.py migrate
</code></pre>
<p>Here's a quick demonstration of everything covered so far.</p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643648979/blogmedia/xnqabmxtenajntqepqo3.gif" alt="app setup and config"></p>
<h3>Creating Form</h3>
<p>Django has a functionality to convert the fields in a model to a form which can be used for input in the frontend side. This allows us to just configure the Form and map the form to a view for the specific operation and simply add a tiny snippet of HTML to make a basic API and thus handle the entire heavy lifting.</p>
<p><a href="https://docs.djangoproject.com/en/4.0/topics/forms/modelforms/#modelform">Model Forms</a> are the type of forms in django that does the above mentioned functionality. We need to create a form class in a new file inside the app called <code>forms.py</code> and add the following content:</p>
<pre><code class="language-python"># app_name / forms.py

from django import forms
from .models import Article

class ArticleForm(forms.ModelForm):
    class Meta:
        model = Article
        fields = [
            'title',
            'post',
        ]
</code></pre>
<p>In this snippet, we are creating a Form called <code>ArticleForm</code> it can be anything but (Model_NameForm) is a standard way to name a Model Form. It is a child class of the ModelForm, i.e. it inherits the parent's attributes/properties and methods. So we have the access to it's Meta class which defines some of the attributes like model, fields, etc.</p>
<p>Here, we are only adding <code>title</code> and <code>post</code>, because <code>created</code> and <code>updated</code> are automatically added. So what about <code>author</code>? Well, we can add it automatically by processing the request that will be sent when the form is submitted (we'll discus while creating views).</p>
<p>In a model form, we also have other attributes like <code>exclude</code> which is the opposite of <code>fields</code>, say you have a lot of attributes in a model and you want to skip certain fields then you don't use fields and use exclude to specify only the attributes to  be excluded. Also <code>widgets</code> which can be used to style and properly process the fields in the form to have more control on how the form should be validated and presented.</p>
<h3>Updating Models</h3>
<p>Now, when we have chosen to use Model Form, we need to update the model logic (not the structure so no migrations).</p>
<p>We'll have to add some methods and redefine certain default parameters in able to use the Class based views and Model forms to the fullest strength.</p>
<pre><code class="language-python"># app_name / models.py

from django.db import models
from django.contrib.auth.models import User
from  django.core.serializers import serialize
import json

class ArticleQuerySet(models.QuerySet):
    def serialize(self):
        list_value = list(self.values(&quot;id&quot;,&quot;author&quot;,&quot;title&quot;,&quot;post&quot;))
        return json.dumps(list_value)

class ArticleManager(models.Manager):
    def get_queryset(self):
        return ArticleQuerySet(self.model,using=self._db)

class Article(models.Model):
    title = models.CharField(max_length=127, verbose_name=&quot;heading&quot;)
    post = models.TextField(verbose_name='content')
    author = models.ForeignKey(User, on_delete=models.CASCADE)
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now=True)

    objects = ArticleManager() 

    def __str__(self):
        return self.title

    def serialize(self):
        data = {
            &quot;id&quot;: self.id,
            &quot;title&quot;: self.title,
            &quot;author&quot;: self.author.id,
            &quot;post&quot;: self.post,
        }
        data = json.dumps(data)
        return data

    class Meta:
        verbose_name_plural = 'Articles'
</code></pre>
<p>Let's break down what was added:</p>
<ul>
<li>
<p>Created two new classes (<code>ArticleManager</code> and <code>ArticleQuerySet</code>)
This was done to manage and serialize the model attributes. We need a standard like JSON to parse and return the data from the client to the server and vice versa. The <code>ArticleQuerySet</code> actually converts the <code>QuerySet</code> into list and dumps into a json object. <strong>A QuerySet in Django is collection of SQL queries.</strong></p>
</li>
<li>
<p>Function serialize
The serialize function actually converts the <code>QuerySet</code> into a JSON object which in turn is used as a utility function for the model. <strong>It returns the JSON object hence a serialized object from a queryset.</strong></p>
</li>
<li>
<p>Editing the object for the model
The object for the model is re initialized as a object of the <code>ArticleManager</code> class which in turn manages the and process the query set and returns the serialized object.</p>
<p>That's all done from the models for the app and now we finally move to the views.</p>
</li>
</ul>
<h3>Creating Views</h3>
<p>Now, we can start writing the views to actually add the CRUD functionality ourselves and slowly develop the frontend as well.</p>
<p>To start creating views, we need to import some built-in classes used for making APIs like the Django generic views. Some of the core views(classes based) for making a CRUD API include:</p>
<ol>
<li><a href="https://docs.djangoproject.com/en/4.0/ref/class-based-views/generic-editing/#createview">CreateView</a></li>
<li><a href="https://docs.djangoproject.com/en/4.0/ref/class-based-views/generic-editing/#updateview">UpdateView</a></li>
<li><a href="https://docs.djangoproject.com/en/4.0/ref/class-based-views/generic-editing/#deleteview">DeleteView</a></li>
<li><a href="https://docs.djangoproject.com/en/4.0/ref/class-based-views/generic-display/#detailview">DetailView</a></li>
<li><a href="https://docs.djangoproject.com/en/4.0/ref/class-based-views/generic-display/#listview">ListView</a></li>
</ol>
<p>These views help in making the API in Django easily. We simply need to add the template name to add a customize the layout, make forms, define the sets of fields to input from the client and which fields to process by the server side.</p>
<p>Let's create them one by one and understand the structure of class based views.</p>
<h4>Create View</h4>
<p>This view as the name suggests is used for creating a entry in a model(database) or we can also say that it will act as the <code>Create</code>(POST request) in the CRUD API.</p>
<pre><code class="language-python"># app_name / views.py

from django.views.generic.edit import ( 
    CreateView, 
    UpdateView, 
    DeleteView )

from .models import Article
from .forms import ArticleForm

class ArticleCreateView(CreateView):
    model = Article
    form_class = ArticleForm
    success_url = '/api/'

    def form_valid(self, form):
        form.instance.author = self.request.user
        return super(ArticleCreateView, self).form_valid(form)
</code></pre>
<p>Here, we have created a Class based view called <code>ArticleCreateView</code> which inherits the <code>CreateView</code> class from the <code>django.views.generic.edit</code> module. Here as similar to the <code>ArticleForm</code> class, we have certain attributes to pass like the model, form class and the success url.</p>
<ul>
<li>The <code>model</code> simply indicates to which table or model we are creating the view for.</li>
<li>The <code>form_class</code> denotes the ModelForm class we are using for the Create view.</li>
<li>The <code>success_url</code> is the url route to which to redirect after a successful POST request/form submission.</li>
</ul>
<p>Remember we discussed about adding <code>author</code> field automatically from the request from the form. This process is carried out in the <code>form_valid</code> function. <a href="https://docs.djangoproject.com/en/4.0/ref/class-based-views/mixins-editing/#django.views.generic.edit.FormMixin.form_valid">Form_valid</a> is a helper built-in function to redirect to the success_url when the form data is being posted. Before actually doing that, we add the author field by setting it as the User by accessing the <code>self.request</code> object. The <code>self.request</code> object actually holds the meta-data about the request that is being sent to the API so we can access the User who is currently logged in.</p>
<p>Though we don't have a login system, we can assume the user is logged in with some user account. Now to handle certain exceptions we can add <a href="https://docs.djangoproject.com/en/4.0/topics/auth/default/#django.contrib.auth.mixins.AccessMixin">Mixins</a>. We can use <a href="https://docs.djangoproject.com/en/4.0/topics/auth/default/#django.contrib.auth.mixins.LoginRequiredMixin">LoginRequiredMixin</a> to only allow the form submission for logged in users and so on.</p>
<h4>Update View</h4>
<pre><code class="language-python"># app_name / views.py

class ArticleUpdateView(UpdateView):
    model = Article
    form_class = ArticleForm
    success_url = '/api/'

    def form_valid(self, form):
        form.instance.author = self.request.user
        return super(ArticleUpdateView, self).form_valid(form)
</code></pre>
<p>This is quite similar to the <code>CreateView</code> except we have to use <code>UpdateView</code> as the base model of the <code>ArticleUpdateView</code> the rest of the attributes remain the same and functioning of the form like pre-rendering the fields to edit are all managed by the <code>UpdateView</code> out of the box.</p>
<h4>Delete View</h4>
<pre><code class="language-python"># app_name / views.py

class ArticleDeleteView(DeleteView):
    model = Article
    form_class = ArticleForm
    success_url = '/api/'

    def form_valid(self, form):
        form.instance.author = self.request.user
        return super(ArticleDeleteView, self).form_valid(form)
</code></pre>
<p>Again, here we do not have to change anything as for attributes and the <code>valid_form</code> function. We'll see the details when we look at the templates.</p>
<h4>List View</h4>
<pre><code class="language-python">#app_name / views.py

class ArticleView(ListView):
    model = Article
    template_name = 'api/list.html'

    def get_context_data(self, **kwargs):
        context = super().get_context_data(**kwargs)
        context['articles'] = Article.objects.filter(author=self.request.user)
        return context
</code></pre>
<p>This view is for listing all the objects/articles for the current logged in author. If we want something like a homepage for all the articles, the thing is as simple as replacing</p>
<pre><code class="language-python">context['articles'] = Article.objects.filter(author=self.request.user)
</code></pre>
<p>by</p>
<pre><code class="language-python">context['articles'] = Article.objects.all()
</code></pre>
<p>You can even apply filters for fetching the latest post by <code>Article.objects.filter().order_by(&quot;-created&quot;)</code>. You get the idea. It boils down to simple python and library function.
Also, we have a new attribute <code>template_name</code> which allows us to use the data in our custom template. In this case we will create the template in the <code>templtes/api/</code> folder called <code>list.html</code>. The function <code>get_context_data</code> is used to fetch the objects from the database and return it as a special dictionary(JSON object) called <code>context</code> which can be rendered in the templates by the key in this case, the <code>articles</code> key will give us all the articles by the current user.</p>
<h4>Detail View</h4>
<pre><code class="language-python"># app_name / views.py

class ArticleDetailView(DetailView):
    model = Article
    template_name = 'api/post.html'
</code></pre>
<p>As opposite to the ListView, we have the DetailView that renders the details of the particular article. Here we don't have to write the <code>get_context_data</code> function as the default object for accessing the model data is <code>object</code>. So, we simply need to specify the <code>template_name</code> and the <code>model</code>. Here, the template is <code>api/post.html</code> in the templates folder.</p>
<h3>Mapping the URLS</h3>
<p>After completing the views, we can map those into a urls to access them as endpoints in the api app. Let's build upon the <code>urls.py</code> file that we created with no url paths.</p>
<pre><code class="language-python"># app_name / urls.py

from django.urls import path
from .views import ( 
        ArticleView, ArticleDetailView, ArticleCreateView, 
        ArticleUpdateView, ArticleDeleteView
        )

urlpatterns = [
        path('', ArticleView.as_view(), name=&quot;listpost&quot;),
        path('article/&lt;pk&gt;/', ArticleDetailView.as_view(), name=&quot;detailpost&quot;),
        path('create/', ArticleCreateView.as_view(), name=&quot;createpost&quot;),
        path('update/&lt;pk&gt;/', ArticleUpdateView.as_view(), name=&quot;updatepost&quot;),
        path('delete/&lt;pk&gt;/', ArticleDeleteView.as_view(), name=&quot;deletepost&quot;),
]
</code></pre>
<p>So, we can see the views are imported from the .views and are mapped to a particular route. We use <code>.as_view</code> function to take a request from a callable class based view and returns the processed response. The <code>&lt;pk&gt;</code> is the id for accessing a particular object. We use the id for detail view, update and delete views.</p>
<h3>Creating Templates and Static files</h3>
<p>We now, have to join the final piece of the puzzle i.e. to create templates. It's quite simple but requires a bit of logic to connect things together. The backend is handled flawlessly by Model Form and Generic views, we simply need to put simple HTML and Django templating language to its correct use.</p>
<pre><code class="language-html"># templates / index.html

&lt;!DOCTYPE html&gt;
&lt;html lang=&quot;en&quot;&gt;
&lt;head&gt;
    &lt;meta charset=&quot;UTF-8&quot;&gt;
    &lt;title&gt;Django Blog&lt;/title&gt;
    {% load static %}
    &lt;link rel=&quot;stylesheet&quot; href=&quot;{% url 'css/style.css' %}&quot;&gt;
&lt;/head&gt;
&lt;body&gt;
    &lt;h1&gt;Articles&lt;/h1&gt;
   {% block body %}
   {% endblock %}
&lt;/body&gt;
&lt;/html&gt;
</code></pre>
<p>A simple HTML doc with link to a stylesheet. We also have a block to further inherit this as the base template.</p>
<pre><code class="language-css"># static / css / style.css

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
a{
    text-decoration-color: #00ffff;
    color: #ff6600;
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
</code></pre>
<p>The below template is a detail view page, that renders the details of the article. We are using the <code>object</code> key from the <code>context</code> dictionary provided by default as mentioned <a href="">here</a>. We are also embedding the <code>Update</code> and <code>Delete</code> buttons to manage the Article.</p>
<p>This is the time to explain about the dynamic urls in Django. We have used the <code>{% url 'updatepost' object.id %}</code> to create a dynamic url that will redirect to the <code>updatepost</code> url and parse with it the id of the object. This is important as the url itself takes the <code>&lt;pk&gt;</code> value to update the post. So we just pass the id like a parameter to the URL. This is the way we create dynamic urls in Django templates which is quite similar to static file urls.</p>
<pre><code class="language-html"># templates / app_name / post.html

{% extends 'index.html' %}

{% block body %}
    &lt;h2&gt;{{ object.title&quot; }}&lt;/h2&gt;
    &lt;p&gt;{{ object.post&quot; &lt;/p&gt;

    &lt;button type=&quot;submit&quot; onclick=&quot;window.location.href='{% url 'updatepost' object.id %}'&quot;&gt;
    Update
    &lt;/button&gt;

    &lt;button type=&quot;submit&quot; onclick=&quot;window.location.href='{% url 'deletepost' object.id %}'&quot;&gt;
    Delete
    &lt;/button&gt;

{% endblock %}
</code></pre>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643911468/blogmedia/n6z33yu4nq9tjsr6nba2.png" alt="CRUD API - Detail View"></p>
<p>We have used the dynamic URLs in Update and Delete View buttons.</p>
<p>The below template is for creating a home page like structure but for the current logged in user. We are displaying all the articles of the logged in the form of a list with a for loop as discussed in the  templating tutorial part. We have access to the key <code>articles</code> from the dictionary <code>context</code> and hence we iterate over the loop from that object and display the title with the url for the post detail view.</p>
<p>Here as well we are using the dynamic url by parsing the <code>article.id</code> as the <code>&lt;pk&gt;</code> primary key to the url.</p>
<pre><code class="language-html"># templates / app_name / list.html

{% extends 'index.html' %}

{% block body %}
    &lt;ul&gt;
        {% for article in articles %}
        &lt;li&gt;&lt;a href=&quot;{% url 'detailpost' article.id %}&quot;&gt;{{ article.title }}&lt;/a&gt;&lt;/li&gt;
        {% endfor %}
    &lt;/ul&gt;
{% endblock %}
</code></pre>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643911630/blogmedia/lakpdkldqwopimhixxya.png" alt="CRUD API - Home Page/List View"></p>
<pre><code class="language-html"># templates / app_name / article_confirm_delete.html

{% extends 'index.html' %}

{% block body %}
    &lt;form method = &quot;post&quot;&gt;
        {% csrf_token %}
        &lt;p&gt; Are you sure to delete Post:&lt;b&gt; {{ object.title }}&lt;/b&gt; ? &lt;/p&gt;
        &lt;input type=&quot;submit&quot; value=&quot;Delete&quot;/&gt;
{% endblock %}
</code></pre>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643911775/blogmedia/qvilf8p3qpj9qfdkpkn4.png" alt="CRUD API - Delete View"></p>
<pre><code class="language-html"># templates/ app_name / article_form.html

&lt;form method=&quot;POST&quot; enctype=&quot;multipart/form-data&quot;&gt;
    {% csrf_token %}
    {{ form.as_p&quot; }}
    &lt;input type=&quot;submit&quot; /&gt;
&lt;/form&gt;
</code></pre>
<p>This is the template that forms the basis for all the forms like Create and Update. We use the template variable <code>{{ form }}</code> to render the form in the template. Additionally we use <code>.as_p</code> to make the form fields as a paragraph tag in HTML. The <code>{% csrf_token %}</code> is the Cross site forgery token allowing secure posting of a form. Don't bother about it much it is important for identifying the user posting the data from the site. The final piece of element is the submit button used to submit the form.</p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643911374/blogmedia/xbwxtg3ry2x4oegtzxmi.png" alt="CRUD API - CREATE View"></p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643911726/blogmedia/zeapgshsk9x9agsw4gsu.png" alt="CRUD API - UPDATE View"></p>
<p>For any references, you can take a look at the <a href="https://github.com/Mr-Destructive/django-blog">GitHub</a> repository for the CRUD API in pure Django.</p>
<h2>Conclusion</h2>
<p>So, from this section, we were able to create a CRUD API in pure Django. This should be the end of the core series of Django Basics, but there are more parts coming up still which are not basics but still worth learning as a beginner. We'll further explore DRF, Forms, CORS, etc. We'll see some other libraries specific for Django. Hopefully from this tutorial series, you were able to learn something and if you have any queries or issues please let me know.</p>
<p>Thank you for reading, please provide feedback on how you felt about the series. Though the series is not over, it's sufficient for a beginner to get pace in Django. Hope you had a great time. Happy Coding :)</p>
