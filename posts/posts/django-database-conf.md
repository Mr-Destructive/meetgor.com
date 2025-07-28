{
  "title": "Django Basics: Database Configurations",
  "status": "published",
  "slug": "django-database-conf",
  "date": "2025-04-05T12:33:49Z"
}

<h2>Introduction</h2>
<p>In this part of the series, we will create an external database and configure the settings for that database. We also look into the process of migrations which is to convert a model which we created in the previous part into the actual structure/schema for the database. I'll be using <code>PostgreSQL</code> for most of the demonstrations but it should be similar for other database management tools. This part will make your Django project link to a local database.</p>
<h2>Selecting a Database</h2>
<p>If you have read my previous article about models, you would have got a glimpse of the tools to manage a database, but not quite to choose a database for your project. This section covers that specifically. We will see what options you have while selecting a database.</p>
<ul>
<li>SQL databases
<ul>
<li>sqlite</li>
<li>PostgreSQL</li>
<li>MySQL</li>
<li>MariaDB</li>
<li>Oracle</li>
</ul>
</li>
</ul>
<p>Selecting a database hugely depends on the type of application you are going to make, but most of the time it's SQL as a choice just because it has been dominated for a variety of application types over 4 decades. Still, NoSQL is growing in popularity and has some advantages over SQL in many modern applications. You need to analyze your project a bit deeper and understand the flow of data in a better way to make a decision about SQL and No-SQL, but most of the time it's gonna be SQL.</p>
<p><strong>Also Django doesn't officially support NoSQL, so you'll have to turn up some third-party libraries to integrate and manage a database.</strong></p>
<p>After deciding the type of database, you have one more decision to make here. It's picking up a DBMS tool. There are a lot of databases like PostgreSQL, MySQL, MariaDB, Oracle, etc. you need to pick whichever you feel comfortable and the one which suits your project architecture and requirements more closely. Though there might be very few differences in all the SQL Database tools there are a few things that distinguishes one from the other.</p>
<h2>Creating a Database</h2>
<p>To create a database, you need to go to the Admin app of the DBMS tool you are using, for Postgres it's pgAdmin, for MySQL it's MySQL Administrator or PHPMyAdmin. You need to do the research for setting up a database locally for your project. But Django already is paired with <code>SQLite</code> which is a relational database but with a few quirks. It is really great to get started with a project without creating or managing an entire database system. The SQLite database is all contained in a file called <code>db.sqlite3</code>.</p>
<p>If you want to work on a particular database like PostgreSQL, MySQL, etc. you need to create the database using the management tool and keep the configuration data like <code>name</code>, <code>host</code>, <code>password</code>, etc. after creating the database.</p>
<p>I'll give a demo of creating a simple database in PostgreSQL but mostly it is a bit different in other DBMS tools as each of them have their own GUI applications. Install <a href="https://www.postgresql.org/download/">Postgres</a> and <a href="https://www.pgadmin.org/download/">pgAdmin</a>.</p>
<p>This is a demonstration of creating a database in pgAdmin -3</p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1642325269/blogmedia/postgcreate_wnmyre.gif" alt="postgres - db creation"></p>
<p>This is how you create a database in pgAdmin-3, it should be quite straightforward and simple to follow in other DBMS tools as well. You will also require a password to open the Admin interface for these tools, so keep that password handy we'll require that later.</p>
<p>The following is the process to create a database/schema in MySQL
<img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1642327473/blogmedia/mysqlcreate_bnkqxg.gif" alt="mysql - db creation"></p>
<p>This will create an empty database in your local machine. We'll want our Django project to link that particular database and use it inside its applications. In the next part, we'll configure the <code>settings.py</code> file to access the database from the local machine.</p>
<h2>Configurations for Database</h2>
<p>We need to configure the <code>settings.py</code> file for integrating the database in our project. Django has a dedicated section for the database in that file. By default, the database configuration is created for the SQLite database which as I said earlier it is the default database that Django uses if not specified and configured.</p>
<p>So, when you first see the configuration for the database in the <code>settings.py</code> file, you will see configuration like the following:</p>
<pre><code class="language-python">DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.sqlite3',
        'NAME': BASE_DIR / 'db.sqlite3',
    }
}
</code></pre>
<p>For PostgreSQL or any other database, we will require more things than the <code>sqlite</code> configuration. Namely, we will need the following:</p>
<ul>
<li><code>name</code> of the database</li>
<li><code>user</code> of the database</li>
<li><code>password</code> for that user</li>
<li><code>host</code> and <code>port</code> for the database.</li>
</ul>
<p>The port is optional as most of the database tools pick up the default port for their purpose. For PostgreSQL, the default port is <code>5432</code> and for MySQL is <code>3306</code>. Leave it blank like <code>'PORT': '',</code> if you are not sure what is the port for that database. We also need to specify the <code>ENGINE</code> as it is the database backend to be used for the project. There are options for specific databases as mentioned in the <a href="https://docs.djangoproject.com/en/4.0/ref/settings/#engine">documentation</a>.</p>
<p><strong>PostgreSQL</strong>:</p>
<p>In PostgreSQL, the default user is <code>postgres</code>, it might depend on your configuration and setup though. The pgAdmin will prompt you for a password when first installing the Postgres on your machine. The password to be used is for the default user of the database. The <code>name</code> is the name that you gave while creating the database in the Postgres admin section. Finally, the host is generally <code>localhost</code> as we are using the local database which is our system, and <code>port</code> as said earlier is <code>5432</code> by default for PostgreSQL.</p>
<pre><code class="language-python">DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.postgresql',
        'NAME': 'blogapp',
        'USER': 'postgres',
        'PASSWORD': '12345678',
        'HOST': 'localhost',
        'PORT': '5432',
    }
}
</code></pre>
<p><strong>MySQL</strong>:</p>
<p>For MySQL, the default <code>user</code> is <code>root</code> and the <code>port</code> is <code>3306</code>. The password is the default password you use to access the MySQL Administrator application.</p>
<pre><code class="language-python">DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.mysql',
        'NAME': 'blogapp',
        'USER': 'root',
        'PASSWORD': '',
        'HOST': 'localhost',
        'PORT': '3306',
    }
}
</code></pre>
<p>For further details on how to configure a Database in a Django project you can check out these articles:</p>
<ul>
<li><a href="https://dev.to/mungaigikure/how-to-set-up-postgres-in-your-django-project-575i">PostgreSQL setup in Django</a></li>
<li><a href="https://medium.com/@omaraamir19966/connect-django-with-mysql-database-f946d0f6f9e3">MySQL setup in Django</a></li>
<li><a href="https://medium.com/code-zen/django-mariadb-85cc9daeeef8">Maria DB setup in Django</a></li>
<li><a href="https://www.mongodb.com/compatibility/mongodb-and-django">Mongo DB setup in Django</a></li>
</ul>
<p>MongoDB is a NoSQL database so, it will be quite different to set up and configure the database at least in the longer run. I've no experience with NoSQL so please forgive me in this part. There are very rare instances you will need a NoSQL database with Django.</p>
<p>There are other configurations as well which might be specific for the purpose and can be explored in the <a href="https://docs.djangoproject.com/en/4.0/ref/databases/">django documentation</a>.</p>
<h3>Verify the database connection</h3>
<p>To check if the database was actually linked in the Django project, I'll introduce you to a great tool in Django: <code>python manage.py shell</code> (make sure to be in the virtual environment).
This command will open a python interpreter in the shell. It is an interactive console so that we can test some aspects in our project. For instance to check if the database is connected or not:</p>
<p>Execute the code after running the command <code>python manage.py shell</code> from a virtual environment.</p>
<pre><code class="language-python">import django    
print(django.db.connection.ensure_connection())
</code></pre>
<p>If this returns <code>None</code> you are good to go. And if the result is tons of error messages, you have something wrong in the configuration or the database itself.</p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1642342286/blogmedia/djb-8-db-connect_l4xqqr.png" alt="db connection test"></p>
<h3>Installing database adapter</h3>
<p>Before we can do anything with the database, we need one thing which is a <strong>database adapter</strong>. Now, this is dependent on the database you are using. The list is as follows:</p>
<ul>
<li><a href="https://pypi.org/project/psycopg2/">psycopg2</a> for PostgreSQL</li>
<li><a href="https://pypi.org/project/mysqlclient/">mysqlclient</a> for MySQL and MariaDB</li>
<li><a href="https://pypi.org/project/cx-Oracle/">cx-Oracle</a> for Oracle</li>
</ul>
<p>SQLite does not require an adapter as the database is a file stored in the base directory as <code>db.sqlite3</code></p>
<p>The above list is nothing but Python packages that allow the Django ORM (which is under the hood python) to operate the database. To install them you can simply install with <code>pip</code> as <code>pip install psycopg2</code> , <code>pip install mysqlclient</code>, and so on. Make sure you are in a python virtual environment.</p>
<p>After installing the particular package, we can now move on to the migration process.</p>
<h2>Migrating the models</h2>
<p>Now, we have a clean and fresh instance of a database created. What next? We'll now use the logic in the <a href="https://mr-destructive.github.io/techstructive-blog/django/python/web-development/2022/01/04/Django-Basics-P7.html">previous part</a>, where we created and designed the logic for the database i.e. <code>Models</code>. Now, we can combine our physical database i.e. the database we just created on a DBMS tool locally, and the logical model to populate and give it a structure.</p>
<p>We'll perform <strong>migration</strong> on our database.</p>
<p>This is where the magic happens and probably the step which should be carefully be executed if working with a real-time (production-level) database. Because the commands you'll run are gonna directly affect the database.</p>
<p>Making a migration is a two-step process. But what is migration?</p>
<p>The Django documentation states it as the version control for the database schema and takes their word for that. It is basically a folder(hidden) that stores the state of your database structure just like commits in git.</p>
<p>Let's see it practically</p>
<h3>Makemigrations</h3>
<p>For actually creating tables, relations attributes in a database from a model. We use the command <code>makemigrations</code>. This command creates a file inside the <code>migrations</code> folder inside the application folder whose model has been recently created or updated. It doesn't affect the database but it creates a file which in turn after a <code>migrate</code> command will be parsed to the actual database using the ORM.
So, for any attribute or logical change inside the model we use the <code>makemigrations</code> command as below:</p>
<pre><code>python manage.py makemigrations
</code></pre>
<p>We do not use the above command if any functional change i.e. operations involving querying the database and other operations that don't affect how the database is structured or stored. Though we have to use the <code>makemigrations</code> command when the fields in the model are changed even slightly.</p>
<h3>Migrate</h3>
<p>To see the result or create the actual tables, attributes, and relations among the tables in the database, we need to run the command <code>migrate</code> which will see the latest file in the migration folder and execute the queries to change the schema of the database. So, this is a very powerful command that can perform SQL queries under the hood with python. The below demonstration shows the output of these two commands in a Postgres database.</p>
<pre><code>python manage.py migrate
</code></pre>
<h3>Demonstration</h3>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1642334950/blogmedia/migration_zi6d2c.gif" alt="db migration"></p>
<p>Let's follow the GIF/video in sequence:</p>
<ol>
<li>Create the database in the DBMS tool</li>
<li>Configure <code>settings.py</code> for the database connection</li>
<li>Install <code>psycopg2</code> or DB-specific adapter.</li>
<li>Create the model.</li>
<li>Run <code>python manage.py makemigrations</code></li>
<li>Run <code>python manage.py migrate</code> (if the above command is a success)</li>
</ol>
<p>Here's the model from the demo:</p>
<pre><code class="language-python">class Article(models.Model):
    title = models.CharField(max_length=255)
    post = models.TextField()
    created = models.DateTimeField(auto_now_add=True)
    updated = models.DateTimeField(auto_now=True)
</code></pre>
<p>If we have multiple apps and we only want Django to migrate a model from a specific app, we can simply add the name of the app to the command. By default, Django will migrate all the models from the <code>INSTALLED_APPS</code> list so we need to specify which app to migrate explicitly.</p>
<pre><code>python manage.py makemigrations app_name
</code></pre>
<p>This also applies to the <code>migrate</code> command.</p>
<p>To understand the migration process more deeply, let us see another demonstration of changing the model and then applying migrations. We'll keep an eye on the <code>migrations</code> folder inside of the app in the project.</p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1642338399/blogmedia/migratefolder_e7fm7n.gif" alt="migration folder demo"></p>
<p>We can see that initially when we applied the migrations in the previous demo, there was a single migration file called <code>0001_initial.py</code> but after we altered/changed the title's <code>max_length</code> from 255 to 127 and executed the <code>makemigrations</code> command another file called <code>0002_alter_article_title.py</code> was created that only contained the changed field. This is very similar to git diff in which we are about to commit to the actual database.</p>
<p>So, that was all about migrations. There is definitely a lot to be covered about migrations and databases, but I'll leave you here for now. There are a lot of things to understand and learn before we can talk more about the databases. We now have a base for other things to get our heads around. We now have a database linked to our project, still, we don't know how to use fetch or query them. That is a topic for another part.</p>
<p>We didn't use SQLite database which is by default provided by Django because in production it can just listen to one request/query at a time. That is not ideal for most web applications. Still, there are use cases of this database as it is suitable for embedded database systems for storage applications in desktop/android applications.</p>
<h2>Conclusion</h2>
<p>So, from this part of the series, we were able to configure a database more specifically an external database (not the default SQLite DB) in a Django project. The concept of migrations was understood and demonstrated with live examples. Also, the process of creating and designing models was used from the previous part to create the structure in an actual database. In the next part, we shall dive into the Admin Section that Django provides to view and tinker with the local database. If you have any queries(not database queries) then please let me know, thank you for reading, and until then Happy Coding :)</p>
