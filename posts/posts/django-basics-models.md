{
  "title": "Django Basics: Creating Models",
  "status": "published",
  "slug": "django-basics-models",
  "date": "2025-04-05T12:33:49Z"
}

<h2>Introduction</h2>
<p>We have seen the basics of Django templating in the previous parts of the series. Now, we can move on to the more backend stuff in Django which deals with the Databases, queries, admin section, and so on. In this particular part, we'll cover the fundamental part of any application in Django i.e the <code>Model</code>. We'll understand what the model is, how to structure one, how to create relationships and add constraints on the fields, etc.</p>
<h2>What ate Models?</h2>
<p>A model is a Django-way(Pythonic) to structure a database for a given application. It is technically a class that can act as a table in a database generally and inside of the class, the properties of it act as the attributes of that database. It's that simple. Just a blueprint to create a table in a database, don't worry about what and where is our database. We will explore the database and its configuration in the next part.</p>
<p>By creating a model, you don't have to write all the basic SQL queries like</p>
<pre><code class="language-sql">CREATE TABLE NAME(
attrb1_name type,
attrb2_name type,
.
.
.
);
</code></pre>
<p>If your application is quite big or is complex in terms of the relations among the entities, writing SQL queries manually is a daunting task and also quite repetitive at times. So Django handles all the SQL crap out of the way for the programmer. So Models are just a Pythonic way to create a table for the project/application's database.</p>
<h2>How to create a Model?</h2>
<p>Creating a model for an application is as easy as creating a class in python. But hey! It's more than that as there are other questions to address while designing the class. You need to design the database before defining the fields in the model.</p>
<p>OK, we'll it's not straightforward as it seems to but still for creating simple and dummy projects to start with. You can use certain tools like <a href="https://www.lucidchart.com/pages/database-diagram/database-design-tool">lucidchart</a>, <a href="https://dbdiagram.io/home">dbdiagrams.io</a>, and other tools you are comfortable with. It's important to visualize the database schema or the structure of the application before tinkering with the actual database inside the project. Let's not go too crazy and design a simple model to understand the process.</p>
<p>Here's a basic model for a Blog:</p>
<pre><code class="language-python">#from django.db import models
from django.contrib.auth.models import User

class Article(models.Model):
    title = models.CharField(max_length=255)
    post = models.TextField()
    author = models.ForeignKey(User, on_delete=models.CASCADE, related_name='Article')
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now=True)
</code></pre>
<p>Ignore the <code>from django.db import models</code> as it is already in the file created by Django. If not, please uncomment the line and that should be good to go.
This is a basic model you might wanna play with but don't dump it anywhere.</p>
<p>We define or create our models in the application inside the project. Inside the application there is already a file called <code>models.py</code> just <strong>append</strong> the above code into it. The application can be any application which makes the most sense to you or better create a app if not already created and name it as <code>article</code> or <code>post</code> or anything you like.</p>
<p>If you are familiar with Python OOP(object-oriented programming), we have basically inherited the <code>models.Model</code> class from the <code>django.db</code> module into our model.</p>
<p>If you want more such examples, let's see more such models :</p>
<p>An E-Mail application core model. Attributes like <code>sender</code>, <code>subject</code> of the mail, <code>body</code> of the mail, <code>recipients_list</code> i.e. the <code>To:</code> section in a mail system and the <code>attachment_file</code> for a file attachment to a mail if any.</p>
<pre><code class="language-python">#from django.db import models
from user import EmailUser

class EMail(models.Model):
    sender = models.EmailField(max_length = 255) 
    subject = models.CharField(max_length = 78)
    body = models.CharField(max_length = 40000)
    recipients_list = models.ManyToManyField(EmailUser, related_name = 'mail_list')
    attachment_file = models.FileField(blank=True)
</code></pre>
<p>A sample model for a note-taking app, consisting of a Note and a Book. A book might be a collection of multiple notes i.e. a single book can have multiple notes so we are using a <code>ManyToManyField</code>, what is that? We'll see that shortly.</p>
<pre><code class="language-python">from django.db import models
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
</code></pre>
<p>These are just dummies and are not recommended to use anywhere especially in a serious project.
So, we have seen a model, but what are these fields and the constraints like <code>on_delete</code>, <code>max_length</code>, and others in the upcoming section on fields.</p>
<h2>Fields in Django</h2>
<p>Fields are technically the attributes of the class which here is the model, but they are further treated as a attribute in a table of a database. So the model becomes a list of attributes which will be then parsed into an actual database.</p>
<p>By creating attributes inside a class we are defining the structure for a table. We have several types of fields defined already by django for the ease of validating and making a constrained setup for the database schema.</p>
<p>Let's look at some of the types of fields in Django Models.</p>
<h3>Types of Fields</h3>
<p>Django has a lot of fields defined in the models class. If you want to go through all the fields, you read through the django docs <a href="https://docs.djangoproject.com/en/4.0/ref/models/fields/#model-field-types">field references</a>. We can access the fields from the <code>models</code> module like <code>name = models.CharField(max_length=10)</code>, this is a example of defining a attributes <code>name</code> which is a CharField. We can set the max_length which acts a constraint to the attribute as we do not want the name field to be greater than 10 and hence parsing the parameter <code>max_length</code> to 10.</p>
<p>We have other field types like:</p>
<ul>
<li><code>IntegerField</code> -&gt; for an integer value.</li>
<li><code>TextField</code> -&gt; for long input of text (like text area in html).</li>
<li><code>EmailField</code> -&gt; for an single valid email field.</li>
<li><code>DateField</code> -&gt; for inputting in a date format.</li>
<li><code>URLField</code> -&gt; for input a URL field.</li>
<li><code>BooleanField</code> -&gt; for a boolean value input.</li>
</ul>
<p>And there are other fields as well which can be used as per requirements.</p>
<p>We also have some other fields which are not directly fields so to speak but are kind of relationship defining fields like:</p>
<ul>
<li><code>ForeignKey</code> -&gt; Define a many-to-one relationship to another model/class.</li>
<li><code>ManyToManyField</code> -&gt; define a many-to-many relationship to another model/class.</li>
<li><code>OneToOneField</code> -&gt; define a one to one relationship between different tables/model/class.</li>
</ul>
<p>So, that's about the field types for just a feel of how to structure or design a database table using a model with some types of attributes. We also need to talk about constraints which needs to added to the fields inside the models.</p>
<h3>Field Options/Arguments</h3>
<p>We can add constraints and pass arguments to the fields in the models. We can add arguments like <code>null</code>, <code>blank</code>, <code>defualt</code>, <code>choices</code>, etc.</p>
<ul>
<li><code>null=True/False</code> -&gt; Set a check for the entry in the table as not null in the database.</li>
<li><code>blank=True/False</code> -&gt; Set a check for the input validation to empty or not.</li>
<li><code>unique=True/False</code> -&gt; Set a constraint to make the entry unique throughout the table.</li>
<li><code>defualt=anyvalue</code> -&gt; Set a default value for the field.</li>
<li><code>choices=list</code> -&gt; Set a list of defined choices to select in the field (a list of two valued tuple).</li>
</ul>
<p>We also have another constraint specific to the fields like <code>max_length</code> for <code>CharField</code>, <code>on_delete</code> for ForeignKey which can be used as a controller for the model when the related model is deleted, <code>verbose_name</code> to set a different name for referencing the entry in the table/model from the admin section compared to the default name of the model, <code>verbose_name_plural</code> similar to the <code>verbose_name</code> but for referencing the entire table/model. Also <code>auto_now_add</code> and <code>auto_now</code> for <code>DateTimeField</code> so as to set the current date-time by default.</p>
<p>More options and arguments that can be passed to the fields in models are given in the django docs <a href="https://docs.djangoproject.com/en/4.0/topics/db/models/#field-options">field options</a></p>
<p>These are some of the options or arguments that we can or need to pass to the fields to set up a constrained schema for our database.</p>
<h3>Meta class</h3>
<p>Meta class is a nested class inside the model class which is most of the times used for ordering the entries(objects) in the table, managing permissions for accessing the model, add constraints to the models related to the attributes/fields inside it, etc.</p>
<p>You can read about the functionalities of the Meta class in the <a href="https://docs.djangoproject.com/en/4.0/ref/models/options/">documentation</a>.</p>
<h2>Model methods</h2>
<p>As a class can have functions, so does a model as it is a Python class after all. We can create kind of a helper methods/functions inside the model. The model class provides a helpful <code>__str__()</code> function which is used to rename an object from the database. We also have other predefined helper functions like <code>get_absolute_url</code> that generates the URL and returns it for further redirection or rendering.</p>
<p>Also, you can define the custom functions that can be used as to help the attributes inside the model class.</p>
<h2>Django ORM</h2>
<p>Django has an Object Relational Mapper is the core concept in Django or the component in Django that allows us to interact with the database without the programmer writing SQL/DB queries. It is like a Pythonic way to write and execute sql queries, it basically abstracts away the layer to manually write SQL queries.</p>
<p>We'll explore the details of how the ORM works under the hood but it's really interesting and fascinating for a Beginner to make web applications without learning SQL(not recommended though personally). For now, its just magical to see Django handling the DB operations for you. You can get the references for learning about the Queryset in ORM from the <a href="https://docs.djangoproject.com/en/4.0/ref/models/querysets/">docs</a></p>
<h2>Example Model</h2>
<p>Let us set up a model from what we have learned so far.</p>
<p>We'll create a model for a Blog Post again but with more robust fields and structure.</p>
<pre><code class="language-python">#from django.db import models
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
      
</code></pre>
<p>We can see in the above model that we have defined the Meta class which is optional and is generally written to modify how to entries inside the table appear or order with other functionalities as well. We have also added the choices option in the status field which has two choices <code>Draft</code> and <code>Publish</code> one which is seen by the django interface and the other to the end-users. We have also added certain fields like slug that will create the URL for the blog post, also certain options like <code>unique</code> has been set to restrict duplicate entries being posted to the database. The <code>related_name</code> in the <code>ForeignKey</code> refers to the name given to the relation from the Article model to the User model in this case.</p>
<p>So, we can see that Django allows us to structure the schema of a database. Though nothing is seen as an end result, when we configure and migrate the model to our database we will see the results of the hard work spent in creating and designing the model.</p>
<h2>Database Specific fields</h2>
<p>By this time, you will have gotten a feel of what a database might be. Most of the projects are designed around SQL databases but No-SQL databases and others are also used in cases which suite them the most. We have tools to manage this database in SQL we call it the Database Management System (DBMS). It's just a tool to manage data, but there is not just a single Database management tool out there, there are gazillions and bazillions of them. Most  popular include <code>MySQL</code>, <code>PostgreSQL</code>, <code>SQLite</code>, <code>Oracle</code>, <code>Microsoft Access</code>, <code>Maria DB</code>, and tons of others.</p>
<p>Well, these different DBMS tools are almost similar with a few hiccups here and there. So, different Database tools might have different fields they provide. For Example, in Database <code>PostgreSQL</code> provides the ListField which <code>SQLite</code> doesn't that can be the decision to be taken before creating any project. There might be some fields that some DBMS provide and other doesn't.</p>
<h2>Conclusion</h2>
<p>We understood the basics of creating a model. We didn't touch on the database yet but the next part is all about configuration and migration so we'll get hands-on with the databases. We covered how to structure our database, how to write fields in the model, add constraints and logic to them and explore the terminologies in Django like ORM, Database Types, etc.</p>
<p>Thank you for reading the article, if you have any feedback kindly let me know, and until then Happy Coding :)</p>
