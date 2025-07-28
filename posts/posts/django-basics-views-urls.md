{
  "title": "Django Basics: Views and URLS",
  "status": "published",
  "slug": "django-basics-views-urls",
  "date": "2025-04-05T12:33:50Z"
}

<h2>Introduction</h2>
<p>After getting familiar with the folder structure of the Django framework, we'll create our first view in an app. The basics of creating and mapping a view with a URL will be cleared by the end of this part.</p>
<h2>Creating Views</h2>
<blockquote>
<p>Views are the functions written in python as a logic control unit of the webserver</p>
</blockquote>
<p>To create a view or typically-like function, we need to write a function in the <code>views.py</code> file inside of the application folder. The function name can be anything but should be a sensible name as far as its usability is concerned. Let's take a basic example of sending an HTTP response of &quot;Hello World&quot;.</p>
<h4>project_name/app_name/views.py</h4>
<pre><code class="language-python">from django.http import HttpResponse

def index(request):
  return HttpResponse(&quot;Hello World&quot;)
</code></pre>
<p>Yes, we are simply returning an HTTP Response right now, but rendering Templates/HTML Documents is quite similar and easy to grasp in Django. So, this is a view or a piece of logic but there is a piece missing in this. Where should this function be used? Of course a URL i.e a path to a web server.</p>
<p>We'll see how to map the views to an URL in Django in the next section</p>
<h2>Mapping the Views to a URL</h2>
<p>We need to first create a <code>urls.py</code> file in the application folder to create a map of the URL to be mapped with the view. After creating the file in the same app folder as the <code>views.py</code>, import the function in the view into the file.</p>
<h4>project_name/app_name/urls.py</h4>
<pre><code class="language-python">from .views import index
from django.urls import path 

urlpatterns = [
    path('', index, name=&quot;index&quot;),
]
</code></pre>
<p>The path can be anything you like but for simplicity, we'll keep it blank('') for now.</p>
<p>Now, you have the path for your view to work but it's not linked to the main project. We need to link the app urls to the project urls.</p>
<p>To link the urls of your app to the main project folder, you need to just add a single line of code in the <code>urls.py</code> file of the project folder.</p>
<p>In projectname folder -&gt; urls.py</p>
<h4>project_name/urls.py</h4>
<pre><code class="language-python">from django.contrib import admin
from django.urls import path, include

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', include('post.urls')),
]
</code></pre>
<p>You need to add the line <code>path('', include('post.urls')),</code> and also import the <code>include</code> function from <code>django.urls</code>. This additional statement includes the urls or all the <code>urlpatterns</code> in the <code>post</code> app from the <code>urls.py</code> file into the project's url-routes.</p>
<p>Here, the URL path can be anything like <code>'home/'</code>, <code>'about/'</code>, <code>'posts/'</code>, etc. but since we are just understanding the basics, we'll keep it <code>''</code> i.e. the root URL.</p>
<p>You can also see that there is another route in our project <code>'admin/'</code> which is a path to the admin section. We'll explore this path and the entire Admin Section in some other part of this series.</p>
<p>Now if you start the server and visit the default URL i.e <code>http://127.0.0.1:8000</code>, you will see a simple HTTP message <code>Hello World</code>.</p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1638194390/blogmedia/uv1_xf4byq.png" alt="Hello World view"></p>
<h2>Breaking the <code>path</code> function in urlpatterns</h2>
<p>The path function in the urlpatterns takes in at least 2 parameters, i.e. the URL pattern and the view of any other function that can be related to the webserver.</p>
<pre><code>path( ' ',   view,    name )
       ^       ^        ^ 
       |       |        |
       |       |     url_name
       |   function_name
   url_path    
</code></pre>
<h3>URL path</h3>
<p>The URL Path is the pattern or literally the path which you use in the Browser's search bar. This can be static i.e. some hard-coded text like <code>home/</code>, <code>user/</code>, <code>post/home/</code>, etc. and we can also have dynamic URLs like <code>post/&lt;pk:id&gt;/</code>, <code>user/&lt;str:name&gt;/</code>, etc. here the characters <code>&lt;pk:id&gt;</code> and <code>&lt;str:name&gt;</code> will be replaced by the actual id(integer/primary key) or the name(String) itself.</p>
<p>This is used in an actual web application, where there might be a user profile that needs the unique user-id to render it specifically for that user. The User-Profile is just an example, it can anything like posts, emails, products, any form of a content-driven application.</p>
<h3>View</h3>
<p>The view or the function is the name of the function that will be attached to that URL path. That means once the user visits that URL, the function is literally called. <strong>View is just a fancy word for a function(or any logic basically).</strong> There is a lot to be covered when it comes to <code>View</code> as there are a lot of ways to create it, there are two types of views, how to use them for various use-cases that can be learned along the way because it is a topic where the crust of Django exists.</p>
<p>We'll learn to create different implementations and structure our views, for time-being just consider them as the unit where every operation on the web can be performed. We can create other standalone functions in python to work with the views to make it a bit structured and readable.</p>
<h3>URL Name</h3>
<p>This is an optional parameter to the path function as we do not mandatorily need to give the URL map a name. This can be really useful in multi-page application websites where you need to link one page to another and that becomes a lot easier with the URL name. We do not need this right now, we'll touch it when we'll see the Django Templating Language.</p>
<h2>Example Views</h2>
<p>Let's create some examples to understand the working of Views and URLs. We'll create a dynamic URL and integrate the Python module in the views to get familiarized with the concept.</p>
<h3>Dynamic URLs</h3>
<p>We can use the dynamic URLs or placeholder variables to render out the content dynamically. Let's create another set of View and URL map.</p>
<h4>project_name/app_name/views.py</h4>
<pre><code class="language-python">def greet(request, name):
    return HttpResponse(&quot;Welcome, &quot; + name)
</code></pre>
<p>This view or function takes an additional argument called <code>name</code> and in response, it just says <code>Welcome, name</code> where the name can be any string. Now after creating the view, we need to map the view to a URL pattern, We'll add a path for this greet function.</p>
<h4>project_name/app_name/urls.py</h4>
<pre><code class="language-python">path('greet/&lt;str:name&gt;/', greet, name=&quot;greet&quot;),
</code></pre>
<p>You can see how we have created the url-pattern here. The greet part is static but the <code>&lt;str:name&gt;</code> is a variable or just a URL parameter to be passed to the view as the value of the variable <code>name</code>. We have also given the URL map a name called greet, just for demonstration of its creation.</p>
<p>You'll get an error, 100% if you are blindly following me! Didn't you forget something?</p>
<p>Import the greet function from the views like so:</p>
<pre><code class="language-python">from .views import index, greet  
</code></pre>
<p>So, after we visit the URL <code>https://127.0.0.1:8000/greet/harry</code>, you should see a response <code>Welcome, harry</code> as simple as that.</p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1638252762/blogmedia/uv-greet_e2wg5o.gif" alt="Greet URL Demo"></p>
<p>Now, how is this working? We see the view first. The function takes two parameters one is most common the request which stores the meta-data about the request, the other parameter is the name that we will be use to respond to the server dynamically. The name variable is used in the string with the HttpResponse function to return a simple string.</p>
<p>Then, in the URLs, we need to find a way to pass the variable name to the view, for that we use the <code>&lt;string:name&gt;</code> which is like a URL parameter to the view. The path function automatically parses the name to the appropriate view and hence we call the greet function with the name variable from the URL.</p>
<h3>Using Pythonic things</h3>
<p>We'll use some Python libraries or functions in the Django App. In this way, we'll see it's nearly no-brainer to use Python functions or libraries in the Django framework as indeed all files which we are working with are Python files.</p>
<h4>project_name/app_name/views.py</h4>
<pre><code class="language-python">from random import randint

def dice(request):
    number = randint(1,6)
    return HttpResponse(f&quot;It's {number}&quot;)
</code></pre>
<p>This view is using the random module, you can pretty much use other web-compatible modules or libraries. We have used the <code>random.randint</code> function to generate the pseudo-random number between 1 and 6. We have used the f-string (<code>f&quot;{variable}&quot;</code>)styled Response string as int is not compatible with the response concatenation. So this is the logic of our map, now we'll need to link it to a URL-path.</p>
<h4>project_name/app_name/urls.py</h4>
<pre><code class="language-python">path('throw/', dice, name=&quot;dice&quot;),
</code></pre>
<p>Also, import the view name from the views as <code>from .views import dice</code> also add other views if present. Now if we go to the URL <code>https://127.0.0.1:8000/throw/</code>, we shall see a random number in the response. This is how we used Python to make the logic of our view.</p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1638252951/blogmedia/uv-dice_bsodzq.gif" alt="Dice URL Demo"></p>
<p>So, that was the basics of creating and mapping views and urls. It is the most fundamental of the workflow in Django project development. You need to get familiar with the process of mapping Views and urls before diving into Templates, Models, and other complex stuff.</p>
<h2>Conclusion</h2>
<p>From this part of the series, we touched upon the basics of views and URLs. The concept of mapping URLs and views might have been much cleared and it will be even gripping after we explore the Template handling and Static files in the next part. If you have any queries or mistakes I might have made please let me know. Thanks for reading and Happy Coding :)</p>
