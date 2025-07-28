{
  "title": "Django Basics: Static Files",
  "status": "published",
  "slug": "django-basics-static-files",
  "date": "2025-04-05T12:33:49Z"
}

<h2>Introduction</h2>
<p>After creating templates, it should be rather tempting to add some styles and logic to them. Well yes, we'll see how to add static files in a web application using django. Static files are not only CSS, but also media/images and Javascript files as well. In this part of the series, we'll cover the basics of working with static files in django including the configuration, rendering and storing of the static files.</p>
<h2>What are Static Files?</h2>
<p>Static files as the name suggests are the files that don't change, your style sheets(css/scss) are not gonna change for every request from the client side, though the template might be dynamic. Also your logo, images in the design will not change unless you re-design it XD So these are the static files that needs to be rendered along with the templates.</p>
<p>We have basically 3 types of static files, CSS, Javascript files and media files/static templates,etc. They are all rendered in the same way but as per their conventions and usage.</p>
<p>You can learn about the theoretical information on <a href="https://docs.djangoproject.com/en/4.0/howto/static-files/">static files</a> from the django documentation.</p>
<h2>How to configure Static Files</h2>
<p>Firstly you can create a folder for all the static files in the root folder. Usually the convention is <code>static</code> as the name of the folder. So, if you have created the template folder in the root directory, similar to that static folder can be created in that path.</p>
<p>Next after creating the static folder in the project root folder, we need to configure the <code>settings.py</code> file to actually tell Django web server to look for all our static files in that folder. To do that, go to the <code>settings.py</code> file, now by this time you would have known where the <code>settings.py</code> file is (inside the project-named folder). Add the following at the end of the <code>settings.py</code> file.</p>
<pre><code class="language-python"># import os
# STATIC_URL = '/static/'

STATICFILES_DIRS = (
    os.path.join(BASE_DIR, &quot;static/&quot;),
)
</code></pre>
<p>Ignore the <code>import os</code> if you already have imported and the <code>STATIC_URL</code> if already there in the file. The <code>STATICFILES_DIRS</code> is the configuration that we tell the django environment to look for all our static files in the base/root directory of the project where the <code>static/</code> folder is. The <code>os.path.join()</code> actually gets the path of the directory in our operating system to the folder specified in the case of our project the <code>BASE_DIR</code> is the path of the project and we add in the static folder to actually the project path. The final piece and the crucial one is the <code>&quot;static/&quot;</code> path, this can be other location where you have created your static folder within the project.</p>
<p>That's it! Yes, it's that simple. We can now create static files and render them in our templates.</p>
<h2>Creating and Storing Static files</h2>
<p>Now this part is customizable and it depends on your preference, how you want to organize the static folder. The convention that I follow is creating separate folders namely for <code>css</code>, <code>js</code> and <code>assets</code>(or <code>img</code>) mostly. And inside of this folders you can store the respective static files. This also creates the project more scalable in terms of it's maintenance.</p>
<pre><code>static\
  |__css\
  |__js\
  |__assets\
</code></pre>
<p>Let's create a static file and an image to demonstrate the concept of static files in django.</p>
<ul>
<li>css/style.css</li>
</ul>
<pre><code class="language-css">body 
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
</code></pre>
<ul>
<li>assets/tbicon.png</li>
</ul>
<p>Demo Image (that's my blog icon)</p>
<p><img src="https://github.com/Mr-Destructive/techstructive-blog/blob/gh-pages/assets/img/tbicon.png?raw=true" alt="Demo image"></p>
<h2>Rendering Static Files from Templates</h2>
<p>So, after configuring and creating the static files, we now can inject them into our templates. If you try to do the traditional way i.e. linking stylesheets/images/script files with HTML, it just won't work as you expect to and there's no point in using traditional way while creating a web application with a framework. So, there is a framework specific way to do things which make it easier and efficient for the project.</p>
<p>To render any static file, we need to load the static tag which allows us to embed links for the static files into the templates. This means if the static files are not loaded directly instead in production(deploying our application) the static files are stored in a folder <code>STATIC_ROOT</code> which the server then loads, we'll see how that internally works when we get to deployment techniques for Django project.</p>
<p>To load the static files from our configuration, we can simpy include the tag on top of the template.</p>
<pre><code>{% load static %}
</code></pre>
<p>The above templating tag will load the <code>static</code> tag which allows us to embed the links to the static files as explained earlier.</p>
<p>Now, we can actually access any file with the static folder in our templates with a particular syntax as below:</p>
<pre><code class="language-html">&lt;link rel=&quot;stylesheet&quot; href=&quot;{% static 'css/style.css' %}&quot;&gt;  
</code></pre>
<p>Its just a example how to load the file, we are calling the static tag which we have loaded in previously and from there we are referencing the css file. The compact syntax would be : <code>{% static  'path-to-file'  %}</code></p>
<p><strong>NOTE: The path to the static file is relative from the Static folder, i.e. enter the path of the file considering the static folder as the base directory.</strong></p>
<h3>Demonstration of the static file</h3>
<p>Let's render the static file which we created earlier i.e. the css file and the image into a template.</p>
<p>Assuming you have a app called <code>post</code> in your django project, you can render static files as below:</p>
<h1>templates/home.html</h1>
<pre><code class="language-html">&lt;!DOCTYPE html&gt;
&lt;html lang=&quot;en&quot;&gt;
&lt;head&gt;
    &lt;meta charset=&quot;UTF-8&quot;&gt;
    &lt;title&gt;Django Blog&lt;/title&gt;
    {% load static %}
    &lt;link rel=&quot;stylesheet&quot; href=&quot;{% static 'css/style.css' %}&quot;&gt;  
&lt;/head&gt;
&lt;body&gt;
    &lt;h1&gt;Hello, World!&lt;/h1&gt;
    {% block body %}
    &lt;p&gt;This is not going to get inherited &lt;/p&gt;
    {% endblock %}
    &lt;p&gt;This will be inherited&lt;/p&gt;
&lt;/body&gt;
&lt;/html&gt;
</code></pre>
<p>We are loading the static tag and then loading the css file using the tag syntax as explained above.</p>
<h1>static/css/style.css</h1>
<pre><code class="language-css">body 
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
</code></pre>
<p>This is the static file,<code>style.css</code> stored inside the css folder of the static folder. This contains basic (very lame) CSS styling as we can understand.</p>
<h1>post/views.py</h1>
<pre><code class="language-python">from django.shortcuts import render

def home(request):
    return render(request, 'home.html')
</code></pre>
<p>The <code>views.py</code> file has the function that renders the template <code>home.html</code> from the templates folder inside the application specific folder.</p>
<h1>post/urls.py</h1>
<pre><code class="language-python">from django.urls import path
from post import views

urlpatterns = [
        path('', views.home, name=&quot;home&quot;),
        ]
</code></pre>
<p>This is the application level configuration for the url routes to the views linking the views(functions) from the <code>views.py</code> file. The url in this file(code-snippet) is linking the root url('') to the home view in the <code>views.py</code> file.</p>
<h1>Blog/urls.py</h1>
<pre><code class="language-python">from django.contrib import admin
from django.urls import path, include
from django.views.generic import TemplateView

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', include('post.urls')),
]
</code></pre>
<p>The urls file in the project folder is the core configuration for project level url routes to individual applications within the project.</p>
<p>Append the following if your templates and static files are not configured properly.</p>
<h1>Blog/settings.py</h1>
<pre><code class="language-python">import os   

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
    os.path.join(BASE_DIR, &quot;static/&quot;),
)
</code></pre>
<p>SO, the result of the above code is as simple template as shown in the picture below:</p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640621276/blogmedia/static-1_vu41gf.png" alt="Static file demo"></p>
<p>This will also work if you do it with traditional HTML syntax, but I'd explained why it's not recommended to do it while using frameworks.</p>
<p>Let's see how static files are rendered in inherited templates. We'll tinker with the <code>for.html</code> template created in the <a href="https://mr-destructive.github.io/techstructive-blog/django/python/web-development/2021/12/14/Django-Basics-P5.html">previous part</a>.</p>
<h1>template/for.html</h1>
<pre><code class="language-django">{% extends 'home.html' %}
{% load static %}

{% block body %}
    &lt;img src=&quot;{% static 'assets/tbicon.png' %}&quot; height=&quot;50px&quot; width=&quot;50px&quot; /&gt;
    &lt;ul&gt;
        {% for sport in sport_list %}
        &lt;li&gt;{{ sport }}&lt;/li&gt;
        {% endfor %}
    &lt;/ul&gt;
{% endblock %}
</code></pre>
<p>We will have re-load the static tag for each template only if we need to include a new static file in the template. So we use the <code>{% load static %}</code> again as we are loading the static file (image) in this template.</p>
<h1>post/views.py</h1>
<pre><code class="language-python">from django.shortcuts import render

def for_demo(request):
    sports = ('football', 'cricket', 'volleyball', 'hockey', 'basketball')
    return render(request, 'for.html', {'sport_list': sports})

def home(request):
    return render(request, 'home.html')
</code></pre>
<h1>post/urls.py</h1>
<pre><code class="language-python">from django.urls import path
from post import views

urlpatterns = [
        path('', views.home, name=&quot;home&quot;),
        path('for/', views.for_demo, name=&quot;fordemo&quot;),
        ]
</code></pre>
<p>So, that's the url and view map created, we can now be able to see the result in the <code>127.0.0.1:8000/for/</code> url to see the below result:</p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640622976/blogmedia/static-tempinh_peyjrg.png" alt="Static demo for inheritance of tempaltes"></p>
<p>The list style has been changed and thus we can see that the CSS from the parent template is also being inherited.</p>
<p>Here is the django project structure which I have created with this series so far:</p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1640624705/blogmedia/trr-static_bgt9du.png" alt="Folder tree structure"></p>
<p>So that has been it for the Static files in Django. Though there are lot of depth for rendering and loading the static files, we'll explore as we get our grasp in the django and web development terminologies.</p>
<h2>Conclusion</h2>
<p>So, from this article, we were able to configure and render static files like CSS/Images and optionally Javascript into the Django application. We covered from ground how to configure, load and structure the folder for storing all the static files at the project level.</p>
<p>Hope you found it helpful and if you have any queries please let me know. We'll start with the databases probably from the next part in the Django Basics Series. Until then have a great week and as always Happy Coding :)</p>
