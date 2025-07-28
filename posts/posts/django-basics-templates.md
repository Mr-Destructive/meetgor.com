{
  "title": "Django Basics: Templates",
  "status": "published",
  "slug": "django-basics-templates",
  "date": "2025-04-05T12:33:49Z"
}

<h2>Introduction</h2>
<p>After learning the basics of views and URLs, we can now move on to the next concept i.e. Templates. In Django, Templates are quite an important component for the application as it acts as the <code>frontend</code> for the web application. With the help of templates and some features provided by Django, it becomes very intuitive and simple to make dynamic web content.
In this part, we understand what are templates and what is the way to render them.</p>
<h2>What are Templates</h2>
<p>Templates are simply a <code>html</code> document or kind of a wireframe for content to be displayed for the web app. Templates allow us to render some more relevant pieces of data rather than simple text HTTP responses as we did earlier. We can even re-use certain components of a template in other using the Django Templating Language (more on this later).</p>
<p>So, using HTML templates, we can write a complete Webpage. If you are unfamiliar with HTML, you can check out the basics of HTML with this <a href="https://www.youtube.com/playlist?list=PL081AC329706B2953">playlist</a>.</p>
<p>Even If you are not familiar with HTML, this tutorial might be quite basic and not overwhelm you with all the tags.</p>
<h2>Creating Templates</h2>
<p>To create a Template, we can write a simple HTML document like the below:</p>
<p>Create a folder <code>templates</code> in the base folder, inside the templates folder, create a file <code>index.html</code></p>
<p><strong>templates\index.html</strong></p>
<pre><code class="language-html">&lt;!DOCTYPE html&gt;
&lt;html lang=&quot;en&quot;&gt;
&lt;head&gt;
    &lt;meta charset=&quot;UTF-8&quot;&gt;
    &lt;title&gt;Django Blog&lt;/title&gt;
&lt;/head&gt;
&lt;body&gt;
    &lt;h1&gt;Hello, World!&lt;/h1&gt;
&lt;/body&gt;
&lt;/html&gt;
</code></pre>
<p>This is a simple HTML template, with the <code>&lt;h1&gt;</code> tags. As Django is a framework, there is a standard for storing all the templates for the project and application. There are a couple of standard of options:</p>
<ul>
<li>One of which is creating a <code>templates</code> folder in the root folder as discussed earlier, also we need to modify the <code>project_name/settings.py</code> file.</li>
</ul>
<p>Inside the <code>settings.py</code> file, we need to locate the <code>TEMPLATES</code> section and modify as below:</p>
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
</code></pre>
<p>In this snippet, we have changed the <code>DIRS</code> option to search the templates in the folder <code>templates</code> in the root directory.</p>
<ul>
<li>The other standard is creating a templates folder in each application.</li>
</ul>
<p>We can create the templates folder in each application instead of a single folder.</p>
<h2>Rendering Templates</h2>
<p>After creating a template and making the required settings to make sure Django is able to pick up those templates, we need to work with views and URLs to actually render those templates.</p>
<p>There are a couple of ways to render templates in Django and some of them are discussed below:</p>
<h3>Using TemplateView</h3>
<p><a href="https://docs.djangoproject.com/en/3.2/ref/class-based-views/base/#django.views.generic.base.TemplateView">TemplateView</a> is a class which is comes with <code>django.views.generic</code> library. This class allows us to render a template by providing in the name of the template, arguments or variables to be parsed, and so on.</p>
<p>The simplest way to render a template is by the following way:</p>
<pre><code class="language-python">from django.contrib import admin
from django.urls import path, include

from django.views.generic import TemplateView

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', TemplateView.as_view(template_name=&quot;index.html&quot;), name=&quot;index&quot;),
]
</code></pre>
<p>We need to import the <code>TemplateView</code> from the <code>django.core.generic</code> so as to use the class for rendering the template.</p>
<p>The <code>TemplateView</code> class takes in a couple of arguments, we'll use the <code>template_name</code> as an argument that takes in the name of the template. Here, we use the <code>index.html</code> as the template which we created earlier. We don't need to specify the entire path to the template as we make modifications in the <code>settings.py</code> file to pick the template from the mentioned directory. We use <code>as_view</code> function to load the class as a function/view.</p>
<p>Activate the virtual environment for the proper functioning of the project.</p>
<p>After activating the virtual environment we can run the server as follows:</p>
<pre><code class="language-terminal">python manage.py runserver
</code></pre>
<p>We can now see the following output and thus, we are now rendering a simple HTML template in Django.</p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639384994/blogmedia/templ1_vbwp5d.png" alt=""></p>
<h3>Using render</h3>
<p>We can also use the <a href="https://docs.djangoproject.com/en/4.0/topics/http/shortcuts/#render">render function</a> from <code>django.shortcuts</code> to simply render a template. But we will create a Python function or a View to actually render the template. So, we'll create a View-URL map as we created in the <a href="https://mr-destructive.github.io/techstructive-blog/django/python/web-development/2021/11/30/Django-Basics-P4.html">previous part</a>.</p>
<p>Firstly, let's create a view function in the <code>post/views.py</code> file, more generally (<code>app_name/views.py</code> file). Firstly, we need to import the render function from <code>django.shortcuts</code> and then return the function call of render.</p>
<pre><code class="language-python">from django.shortcuts import render

def home(request):
    return render(request,'index.html')
</code></pre>
<p>And in the URLs, we'll create a different pattern like for e.g. 'home/'</p>
<pre><code class="language-python">from django.urls import path
from post import views

urlpatterns=[
        path('home/',views.home,name=&quot;home&quot;),
        ]
</code></pre>
<p>So, after creating the View-URL map and making sure the URL of the app is loaded in the project URLs, we can see the result as a simple HTML template.</p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639386932/blogmedia/templ2_rgoppj.png" alt=""></p>
<h2>Django Templating Language</h2>
<p>The <a href="https://docs.djangoproject.com/en/3.2/ref/templates/language/">Django Templating Language</a> is Django's way of making templates more dynamic and easy to write dynamic web applications. We'll take a brief look at what we can do with this type of Templating Language in Django.</p>
<h3>Variables</h3>
<p>This is the most common use case for the Django Templating Language/Engine as we can use the <a href="https://docs.djangoproject.com/en/3.2/ref/templates/language/#variables">variables</a> from the Backend and inject it in the template. We can parse the variable into the template by the syntax : <code>{{ variable_name &quot; }}}}</code></p>
<p>To show its use cases, we can declare a variable in a view and then parse it in the Template. Though it is not dynamic right now we can later on fetch values from the database and store them in the form of variables in our views.</p>
<p><strong>templates/home.html</strong></p>
<pre><code class="language-html">&lt;!DOCTYPE html&gt;
&lt;html lang=&quot;en&quot;&gt;
&lt;head&gt;
    &lt;meta charset=&quot;UTF-8&quot;&gt;
    &lt;title&gt;Django Blog&lt;/title&gt;
&lt;/head&gt;
&lt;body&gt;
    &lt;h1&gt;Hello, {{ name }}&lt;/h1&gt;
&lt;/body&gt;
&lt;/html&gt;
</code></pre>
<p><strong>post/views.py</strong></p>
<pre><code class="language-python">from django.shortcuts import render

def variable_demo(request):
    name = &quot;Kevin&quot;
    return render(request, 'home.html', {'name':name})
    #The name can be anything, like a database query object, form detail, etc

</code></pre>
<p>As we can see the variable in views is passed as a dictionary in python. The reference key along with a value of the variable as the name of the variable. We will use the key in the templates to parse the value of the variable.</p>
<p><strong>post/urls.py</strong></p>
<pre><code class="language-python">from django.urls import path
from post import views

urlpatterns=[
        path('home/',views.home,name=&quot;home&quot;),
        path('vardemo/',views.variable_demo, name=&quot;var&quot;),
        ]
</code></pre>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639389288/blogmedia/templ3_wvhagw.png" alt="variable demo"></p>
<p>As we can see, we were able to load the variable into the template using the Django Templating Engine.</p>
<h3>Conditional statement</h3>
<p>We can even use the conditional statement in the Template using a very simple syntax. We can use <code>{% if condition&quot;  }} %}</code> to use certain special kinds of blocks in the Template. We need to end those blocks as well using the syntax <code>{% endif  %}</code>, here <code>if</code> can be other blocks which we'll explore ahead.</p>
<p>To create a basic if condition in the template, we can understand with the following example.</p>
<p><strong>app_name/views.py</strong></p>
<pre><code class="language-python">from django.shortcuts import render
from random import randint

def if_demo(request):
   number = randint(1,10)
   return render(request, 'if_else.html', {'num':number})
</code></pre>
<p>Here, we have used the key name as <code>num</code> indicating we can give different names to the key which needs to be used in the template to render the values.</p>
<p><strong>app_name/urls.py</strong></p>
<pre><code class="language-python">from django.urls import path
from post import views

urlpatterns = [
        path('if/', views.if_demo, name=&quot;ifdemo&quot;),
        ]
</code></pre>
<p><strong>templates/if_else.html</strong></p>
<pre><code class="language-html">&lt;!DOCTYPE html&gt;
&lt;html lang=&quot;en&quot;&gt;
&lt;head&gt;
    &lt;meta charset=&quot;UTF-8&quot;&gt;
    &lt;title&gt;Django Blog&lt;/title&gt;
&lt;/head&gt;
&lt;body&gt;
    {{ num }}
    {% if num &gt; 5 %}
        &lt;h2&gt;It's Greater then 5&lt;/h2&gt;
    {% elif num == 5 %}
        &lt;h2&gt;It's five!&lt;/h2&gt;
    {% else %}
        &lt;h2&gt;It's less than 5&lt;/h2&gt;
    {% endif %}
&lt;/body&gt;
&lt;/html&gt;
</code></pre>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639411425/blogmedia/templ3_exj0fv.png" alt="if-else demo"></p>
<p>So, as we can see that, we can use the if-else conditions in the template and that is already powerful. This can be a bit messy as to handle mathematical operations or conditions into a single condition. This can really be used for really large datasets that can be shimmed down to really less coding and also improve readability.</p>
<h3>For loop</h3>
<p>Now, the most crucial component of the Django templating language is the loops. We can actually iterate over objects/lists in the template. This becomes a huge concept for actually making a dynamic web application. We n\might want to iterate over all the entries in a database, or any other form of data which can make the app a lot dynamic and feel real-time.</p>
<p>The syntax of for loop is almost similar to the if-else condition. We just replace the condition with the iterator and the list/object from the view context. <code>{% for i in list %}</code>, also end the for loop like <code>{% endfor %}</code>.</p>
<p><strong>app_name/views.py</strong></p>
<pre><code class="language-python">from django.shortcuts import render

def for_demo(request):
    sports = ('football', 'cricket', 'volleyball', 'hockey', 'basketball')
    return render(request, 'for.html', {'sport_list': sports})

</code></pre>
<p>We have created a simple Python list called <code>sports</code> and we parse them to the template using a dictionary object, <code>sport_list</code> as the key for storing the value of the <code>sports</code> list.</p>
<p><strong>app_name/urls.py</strong></p>
<pre><code class="language-python">from django.urls import path
from post import views

urlpatterns = [
        path('for/', views.for_demo, name=&quot;fordemo&quot;),
        ]
</code></pre>
<p><strong>templates/for.html</strong></p>
<pre><code class="language-html">&lt;!DOCTYPE html&gt;
&lt;html lang=&quot;en&quot;&gt;
&lt;head&gt;
    &lt;meta charset=&quot;UTF-8&quot;&gt;
    &lt;title&gt;Django Blog&lt;/title&gt;
&lt;/head&gt;
&lt;body&gt;
    &lt;ul&gt;
        {% for sport in sport_list %}
        &lt;li&gt;{{ sport }}&lt;/li&gt;
        {% endfor %}
    &lt;/ul&gt;
&lt;/body&gt;
&lt;/html&gt;
</code></pre>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639475328/blogmedia/templ3_q8z8fr.png" alt="for loop demo"></p>
<p>We have used simple for loop in Pythonic syntax, we use an iterator in this case, <code>sport</code> acts as an iterator. We use this to store values one by one from the list <code>sport_list</code> which was earlier passed in the views as a key in the dictionary.</p>
<p>Hence, this is quite scalable and used to fetch the objects/entries in the database and hence making it a lot easier to make a dynamic web application faster.</p>
<h2>Template Inheritance</h2>
<p>So, far we have seen that we need to create the base template again and again like all the basic HTML elements, title, and all the basic structure. But what if, we can reuse a specific template in another and extend the functionality of that template into a new one. This avoids the redundancy of writing the entire basic template or the layout of a web app over and over again.</p>
<p>To do that, Django has the Template inheritance. We can use a template as its basic layout or a specific component in the web application. Again, similar to the for, if-else blocks the syntax for inheriting a template is quite similar.</p>
<p>Take, for example, the home.html which consisted of only a <code>&lt;h1&gt;</code> tag in it. We can use this kind of template in other templates to really make it the home page. For that, we first need to enclose the template in a <code>block</code>, which is what allows us to use it in other templates.
To create a <code>block</code>, we simply need to write the following syntax before the component which we do not want in other templates:</p>
<p><strong>templates/home.html</strong></p>
<pre><code class="language-html">&lt;!DOCTYPE html&gt;
&lt;html lang=&quot;en&quot;&gt;
&lt;head&gt;
    &lt;meta charset=&quot;UTF-8&quot;&gt;
    &lt;title&gt;Django Blog&lt;/title&gt;
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
<p>In this we have used the <code>blocks</code> with a name like <code>body</code> as <code>{% block body %}</code> this can be anything you like. We end the block with the similar syntax as the for/if blocks as <code>{% endblock %}</code>. Anything in between the blocks i.e <code>block block_name</code> and <code>endblock</code> is not inherited i.e it is unique to this template.</p>
<p>We will see how we can use this template in other templates. We will actually extend this template and use the blocks to render the content in the template.</p>
<p><strong>templates/if_else.html</strong></p>
<pre><code class="language-html">{% extends 'home.html' %}
{% block body %}
    {{ num }}
    {% if num &gt; 5 %}
    &lt;h2&gt;It's Greater then 5&lt;/h2&gt;
    {% elif num == 5 %}
    &lt;h2&gt;It's five!&lt;/h2&gt;
    {% else %}
    &lt;h2&gt;It's less than 5&lt;/h2&gt;
    {% endif %}
{% endblock %}
</code></pre>
<p>So, we first say to Django to extend the <code>home</code> template i.e. the Django will load the blocks from this template only, remember it will just load and not use the blocks until we explicitly tell it to.</p>
<p>To use the blocks or kind of plug in the template content in the <code>if_else.html</code> or any other template, we need to again call the <code>blocks</code>. Here, we need to write the content inside the <code>blocks</code> to properly parse the blocks as this is an HTML template. The order of opening and closing elements do matter.
So, when we say <code>endblock</code> the last part of the base template is loaded i.e. the closing <code>body</code> and <code>html</code> tags. This is like plugging the template as it is before and after the block body.</p>
<p><strong>app_name/views.py</strong></p>
<pre><code class="language-python">from django.shortcuts import render

def home(request):
    return render(request, 'home.html')

from random import randint

def if_demo(request):
   number = randint(1,10)
   return render(request, 'if_else.html', {'num':number})
</code></pre>
<p><strong>app_name/urls.py</strong></p>
<pre><code class="language-python">from django.urls import path
from post import views

urlpatterns = [
        path('', views.home, name=&quot;home&quot;),
        path('if/', views.if_demo, name=&quot;ifdemo&quot;),
        ]
</code></pre>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639479714/blogmedia/tempinher2_enisls.png" alt="">
<img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639477721/blogmedia/tempinher_lk0op0.png" alt="template inheritance demo"></p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1639479954/blogmedia/template-inh_lc8szo.gif" alt=""></p>
<p>The above gif illustrates the example in a neat way. The block is loaded from the given template as the extended template and hence it plugs the block into the frame of the template.</p>
<h2>Conclusion</h2>
<p>So, from this part of the series, we were able to understand the concept of Templates in Django, we were able to use variables, loops, conditional statements, and template inheriting in a Django application. In the next part, we'll try to touch up with the static files and see how to properly structure and configure them.</p>
<p>Thank you for reading, if you didn't understand any of the examples, please let me know, I'll be happy to share the code. Happy Coding :)</p>
