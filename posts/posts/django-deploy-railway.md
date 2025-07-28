{
  "title": "Django Project with PostgreSQL Deployment on Railway App",
  "status": "published",
  "slug": "django-deploy-railway",
  "date": "2025-04-05T12:33:41Z"
}

<h2>Introduction</h2>
<p>We have already seen the process of deployment of a Django app on Heroku, now we can move to another platform which is <a href="https://railway.app">Railway App</a>. This allows a faster and more intuitive way to deploy a Django app(or any sort of web app). This platform has several <a href="https://railway.app/starters">Starter Apps</a> for shipping your app in several clicks, definitely, it would require some configuration and tweaking a bit.</p>
<p>We will be following similar steps from the <a href="https://mr-destructive.github.io/techstructive-blog/series/django-deployment/">Django-Heroku Deployment Guide</a> for configuring and setting up our Django project for the deployment on the Railway app. We explore a few different ways to create and deploy a Django project on Railway.</p>
<h2>Creating a Project on Railway App</h2>
<p>Once we have our Django project setup, we can move ahead and create a Railway Project. A Railway project is a collection of environments, plugins, services, and deployment variables. By heading on the <a href="https://railway.app/new">New Project</a> webpage, you can log in or sign up for a free account and create a simple project.</p>
<p>Here we can have a few options:</p>
<ol>
<li>Create a Django App Template (Djangopy Starter)</li>
<li>Pull a Project from a GitHub Repository (Deploy from GitHub Repo)</li>
</ol>
<p>We'll look into each of them, the first one is by allowing the <code>Railway</code> bot to create a prebuilt Django template on our GitHub account. The second option is for fetching and deploying an existing Django project repository on GitHub.</p>
<h2>Railway Django Project Template</h2>
<p>Railway provides a simple Django project template, it consists of a single app and a simple view that displays a <code>HttpResponse</code>. The Django project template provided by Railway is open source and available on <a href="https://github.com/railwayapp/starters/tree/master/examples/djangopy">GitHub</a>. With this method, you won't require any django project to create by yourself, it would be a template of django project created by the railway bot. Though for extending and building a meaningful django project, you will have to clone the repository and make necessary changes to it.</p>
<p>Create a Project by heading on to the <a href="https://railway.app/new">New Project</a> Tab and search for <code>Django</code>, you should see a <code>django</code> project as a starter pack.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1652945887/blog-media/django-deployment/railway-django-new-project.png" alt="New Railway Django Project"></p>
<p>Once you select the Django project, we have an initial configuration tab opened up for us. It is mandatory to fill the <code>SECRET_KEY</code> environment variable as it makes your django project more secure. Also, do change the name of the <code>repository</code> as it will be the name of your repository name in your GitHub account.</p>
<p>To create a <code>SECRET_KEY</code> key, you can move into your terminal and do some python wizardry.</p>
<p>Open a Python REPL, by entering the command <code>python</code> or <code>python3</code>. Import the <code>secrets</code> module and run the function <code>token_hex(24)</code> which will generate a key of length 24.</p>
<pre><code>python

import secrets
secrets.token_hex(24)
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1652946842/blog-media/django-deployment/python-secret-key-gen.png" alt="Python Secret Key Generation"></p>
<p>Now, copy the <code>SECRET_KEY</code> without quotes into the prompt and this will create a repository on your GitHub with the provided name. The <code>Railway</code> Bot will create a django project with the name <code>djangopy</code> in that repository with some pre-configured settings.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1652947362/blog-media/django-deployment/create-project-django.png" alt="Django Proejct Create"></p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1652947344/blog-media/django-deployment/railway-bot-djangopy.png" alt="Railway Bot creating Djangopy"></p>
<p>This will create a simple repository on your GitHub but also a django application deployed on Railway along with PostgreSQL Database attached to it.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1652947689/blog-media/django-deployment/deployed-project-dashboard.png" alt="Railway Django Project Dashboard"></p>
<p>So, this has deployed the project on Railway with this <a href="https://djangopy-production-43cb.up.railway.app/">https://djangopy-production-43cb.up.railway.app/</a> URL Link. The name of the link can be configured from the Settings Tab in Dashboard Section and editing the Service Domains and even adding a Custom domain.</p>
<p>So this is how we deploy a basic django application on railway app. Further, you can modify the contents of the Github repository and push the code by committing the changes and it will pick it from there and also deploy it automatically, thereby creating a sort of CI-CD.</p>
<p>We'll be demonstrating this in the next section which is a method to deploy the project from a custom GitHub repository i.e. by setting up everything ourselves and then attaching the Railway app for deployment. The benefit of this method is that we can configure some already existing Django applications by tinkering a bit.</p>
<h2>Deploying from GitHub repository to Railway</h2>
<p>This is the second method for deploying the Railway project. For this we need a proper Django Project, we will set up a django application from scratch, I have already created a simple CRUD application in Django for a Blog on <a href="https://github.com/Mr-Destructive/django-blog">GitHub</a>. This won't be a Django guide for setting up views and creating models, I've explained all the setup of the django-blog in my <a href="https://mr-destructive.github.io/techstructive-blog/series/django-basics/">Django Basics series</a>.</p>
<h3>Creating a Django Application</h3>
<p>For deploying an app, we definitely need an app, we need to create a basic Django application to deploy on the web. We'll be creating a simple blog application with a couple of views and a simple model structure. As for the database, we'll be using Postgres as Railway has an database service for it and it is pretty easy to configure.</p>
<h4>Set up a virtual environment</h4>
<p>We need to set up a virtual environment in order to keep the Django project neat and tidy by managing the project-specific dependencies and packages. We can use the <a href="https://pypi.org/project/virtualenv/">virtualenv</a> package to isolate a python project from the rest of the system.</p>
<pre><code># install the virtualenv package
pip install virtualenv

# create a virtual env for the project
virtualenv .venv

# activate the virtualenv
Windows:
.venv\Scriptsctivate

Linux/macOS:
source .venv/bin/activate
</code></pre>
<p>This will set up the project nicely for a Django project, you now install the core Django package and get started with creating a Django application.</p>
<pre><code class="language-bash"># install django
pip install django

# start a django project
django-admin startproject blog .

cd blog

# create a application in django project
python manage.py createapp api

# Create some models, views, URLs, templates

# run the server
python manage.py runserver
</code></pre>
<p>We assume you already have a Django project configured with some basic URLs, views, and templates or static files as per your project and requirements, for this tutorial I will be using the simple blog application from my previous Django tutorials as a reference. As said earlier, you can follow along with my <a href="https://techstructiveblog.hashnode.dev/series/django-basics">Django Basics</a> series and refer to the Blog Application project on <a href="https://github.com/Mr-Destructive/django-blog">GitHub</a>.</p>
<h3>Configuring the Django Application</h3>
<p>Make sure to create and activate the virtual environment for this django project. This should be done to manage the dependencies and packages used in the project. If you are not aware of the virtual environment and django setup, you can follow up with this <a href="https://mr-destructive.github.io/techstructive-blog/django-setup-script/">post</a>.</p>
<h4>Creating a runtime.txt file</h4>
<p>Now, Firstly we need to specify which type and version of language we are using. Since Django is a Python-based web framework, we need to select the python version in a text file.</p>
<p><strong>runtime.txt</strong></p>
<pre><code>python-3.9.5
</code></pre>
<p>Here, the version can be anything as per your project and the packages installed.</p>
<h4>Creating requirements.txt file</h4>
<p>We'll first create a <code>requirements.txt</code> file for storing all the dependencies and packages installed in the application. This will help in installing dependencies while deploying the application. We can either use a <code>requirements.txt</code> file using <code>virtualenv</code> or a <code>Pipfile</code> using Pipenv. Both serve the same purpose but a bit differently.</p>
<p>Assuming you are in an isolated virtual environment for this Django project, you can create a requirements.txt file using the below command:</p>
<p>Make sure the virtualenv is activated before running the command:</p>
<pre><code>pip freeze &gt; requirements.txt
</code></pre>
<p>This will create a simple text file that contains the package names along with the versions used in the current virtual environment. A simple Django requirements file would look something like this:</p>
<pre><code>asgiref==3.4.1
Django==3.2.11
pytz==2021.3
sqlparse==0.4.2
typing_extensions==4.0.1
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1652956558/blog-media/django-deployment/railway_requirements_freeze.png" alt="pip freeze command"></p>
<p>This is vanilla Django without any additional dependencies but if you have integrated other functionalities like Django Rest Framework, PostgreSQL, Crispy Forms, Schedulers, etc. there might be additional dependencies that become quite crucial for the smooth working of the project.</p>
<p>If you are using pipenv, you don't need to make any efforts to manually activate and manage the virtual environment, the dependencies are installed and taken care of by the pipenv installer. You just need to make sure to install any package with <code>pipenv install</code> and not <code>pip install</code> for better and improved package tracking.</p>
<p>So, that's all we need to take care of packages and keep a list of these integrated packages for the project. You need to update the requirements.txt file every time you install any new package or modify the dependencies for a project. Simply use the command <code>pip freeze &gt; requirements.txt</code> in the activated virtual environment.</p>
<h4>Creating a Procfile</h4>
<p>Next up, we have the <code>Procfile</code>, a procfile is a special file that holds information about the processes to be run to start or activate the project. In our case, for django we need a web process that can manage the server.</p>
<p>A Procfile is a simple file without any extension, make sure to write <code>Procfile</code> as it is the name of the file in the root folder of the project. Inside the file add the following contents:</p>
<p><strong>Procfile</strong></p>
<pre><code class="language-Procfile">web: gunicorn &lt;project_name&gt;.wsgi
</code></pre>
<p>For the Procfile, Railway has a built <a href="https://github.com/railwayapp/starters/blob/master/examples/djangopy/Procfile">Procfile</a> in the Django Template, you can refer to it and create it as follows:</p>
<pre><code class="language-Procfile">web: python manage.py migrate &amp;&amp; gunicorn &lt;project_name&gt;.wsgi
</code></pre>
<p>As we can see we have defined the <code>web</code> process using <code>gunicorn</code>, <a href="https://pypi.org/project/gunicorn/">Gunicorn</a> is a python package that helps in creating WSGI HTTP Server for the UNIX operating systems. So, we need to install the package and update the package dependency list.</p>
<pre><code>pip install gunicorn

pip freeze &gt; requirements.txt
</code></pre>
<h4>Configuring Environment Variables</h4>
<p>We need to keep our secrets for the django project out of the deployed code and in a safe place, we can create environment variables and keep them in a <code>.env</code> file which will be git-ignored. We do not want this <code>.env</code> file to be open source and thus should not be committed.</p>
<p>We first need to create a new secret key if you already have a GitHub repository, chances are you would have committed the default secret key open for the world to see, it is an insecure way of deploying django apps in production.</p>
<p>This should generate a secret key that only you know now. So, just copy the key without the quotes and create a file <code>.env</code> in the root project folder.</p>
<p><strong>.env</strong></p>
<pre><code>SECRET_KEY=76419fd6885a677f802fd1d2b5acd0188e23e001042b05a8
</code></pre>
<p>The <code>.env</code> file should also be added to the <code>.gitignore</code> file, so simply append the following in the <code>.gitignore</code> file</p>
<pre><code>.env
</code></pre>
<p>This is just one of the environment variables in our django project, further, we will also be adding a few other variables like database credentials, debug status, etc.</p>
<p>We have now created environment variables for our django application, we now need a way to parse these environment variables into the Django project.</p>
<h4>Parsing Environment variables using python-dotenv</h4>
<p>We will use <a href="https://pypi.org/project/python-dotenv/">python-dotenv</a> to parse variables into the django settings configurations like <code>SECRET_KEY</code> and <code>DATABASES</code>.</p>
<ul>
<li>Install <code>python-dotenv</code> with pip with the command :</li>
</ul>
<pre><code>pip install python-dotenv
</code></pre>
<p>We need to then modify the default variables in the <code>settings.py</code> file. Firstly, we will load in the <code>.env</code> file for accessing the environment variables for the configuration of the project.</p>
<p>Append the following code, to the top of the <code>settings.py</code> file, make sure don't override the configuration so remove unnecessary imports and configurations.</p>
<pre><code class="language-python"># &lt;project_name&gt;/settings.py

import os
from dotenv import load_dotenv

BASE_DIR = Path(__file__).resolve().parent.parent

load_dotenv(os.path.join(BASE_DIR, &quot;.env&quot;))

</code></pre>
<p>We have imported the package <code>dotenv</code> basically the <code>python-dotenv</code> into the <code>settings.py</code> file and the module <code>os</code> for loading the <code>.env</code> file. The <code>load_dotenv</code> function helps in loading the <code>key-value</code> pairs which are the configuration variables that can act as actual environment variables. We provide in a file to the <a href="https://saurabh-kumar.com/python-dotenv/">load_dotenv</a> function which is the <code>.env</code> file in our case, you can pick up any location for the <code>.env</code> file but make sure to change the location while parsing the file into the <code>load_dotenv</code> function.</p>
<p>After loading the variables into the <code>settings.py</code> file, we now need to access those variables and set the appropriate variables the configuration from the variables received from the <code>load_dotenv</code> function. The <code>os.getenv</code> function to access the environment variables. The <code>os.getenv</code> function takes a parameter as the <code>key</code> for the environment variable and returns the value of the environment variable.</p>
<pre><code class="language-python">SECRET_KEY = os.getenv(&quot;SECRET_KEY&quot;)
</code></pre>
<p>We have secretly configured the <code>SECRET_KEY</code> for the django project. If you have any other variables as simple key-value pairs like <code>AUTH</code> passwords, username, etc. you can use this method to get the configuration variables.</p>
<h4>Add gitignore file</h4>
<p>We would need a <code>.gitignore</code> file for only committing the project and pushing it to the remote repository. Here, we will set up a minimal <code>.gitignore</code> file in our root repository.</p>
<p>Here's a sample <code>.gitignore</code> for a minimal django project.</p>
<pre><code class="language-gitignore">.env/
.venv/
env/
venv/
*.env

*.pyc
db.sqlite3
staticfiles/
</code></pre>
<p>If you want a full-fledged <code>.gitignore</code> for a complex django project, you can take the reference from Jose Padilla's <a href="https://github.com/jpadilla/django-project-template/blob/master/.gitignore">gitignore Template</a> for a django project.</p>
<p>That would be good to go for creating and serving up the project while deploying the project on Railway App.</p>
<h3>Git Commit the Django Project</h3>
<p>Now, we can safely commit the code and push it to a GitHub repository. This will make sure you have a basic django proejct on GitHub from which we can build the Railway app into deployment.</p>
<pre><code>git status 

git add .

git commit -m &quot;config for railway deployment&quot;
</code></pre>
<p>Carefully check the status of the git repository before committing and make sure you don't forget anything by mistake, only commit the files which you think are important and ignore the rest.</p>
<pre><code>git remote add rail https://github.com/Mr-Destructive/django-blog/tree/railway

git push rail railway
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1652956515/blog-media/django-deployment/railway_push_github.png" alt="Django Project Push GitHub"></p>
<h3>Creating the Railway project</h3>
<p>Now, since we have a django project nicely configured and setup on GitHub, we can pull out a railway project and fetch the project from the GitHub repository.</p>
<!-- raw HTML omitted -->
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1652970665/blog-media/django-deployment/railway-proejct-github.png" alt="Railway Proejct from GitHub"></p>
<h4>Spinning up the database</h4>
<p>We also need a database that we can attach in the django project in our Railway application. We can integrate a <code>PostgreSQL</code> database as a service running in our Railway project instance. We can do that by locating the <code>+New</code> tab on the project dashboard and attaching a <code>Database</code> in the drop-down menu.</p>
<!-- raw HTML omitted -->
<p>After creating a database service, we need the credentials of the database or the <code>DATABASE_URL</code> of the PostgreSQL in order to integrate it into the django settings. We can locate into the <code>Connect</code> of the PostgreSQL service and grab the URL of the database. This can be stored in the main django application serves as an environment variable.</p>
<pre><code>DATABASE_URL=postgresql://postgres:SE74bEw@containers-51.railway.app:6919/railway
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1652964755/blog-media/django-deployment/postgres_database_config.png" alt="PostgreSQL database variable config"></p>
<p>In the Django Starter Project provided by Railway, we should have a PostgreSQL database integrated as a Service. So, we can consume the database in our django project.</p>
<p>We will add the <code>DATABASE_URL</code> config variable into the <code>.env</code> file and also add it into the main Django project service so that it can communicate to the database. You need to copy the <code>DATABSE_URL</code> into our local setup file(<code>.env</code> file). Copy the Database URL and paste it into the <code>.env</code> file as follows:</p>
<pre><code class="language-env">DATABASE_URL=postgres://sjxgipufegmgsw:78cbb568e@ec2-52-4-104-184.compute-1.amazonaws.com:5432/dbmuget
</code></pre>
<p>The format for the Postgres URL is as follows:</p>
<pre><code>postgresql://[user[:password]@][netloc][:port][/dbname]
</code></pre>
<h4>Loading Database configuration</h4>
<p>Databases are a bit different as compared to other simple configurations in django. We need to make a few adjustments to the default database configuration. We will install another package <code>dj-database-url</code> to configure <code>DATABASE_URL</code>. Since the databse_url has a few components we need a way to extract the details like the <code>hostname</code>, <code>port</code>, <code>database_name</code>, and <code>password</code>. Using the <code>dj-database-url</code> package we have a few functions that can serve the purpose.</p>
<pre><code>pip install dj-database-url
</code></pre>
<p>At the end of your <code>settings.py</code> file, append the following code.</p>
<pre><code class="language-python">import dj_database_url

DATABASE_URL = os.getenv(&quot;DATABASE_URL&quot;)

DATABASES = {
    &quot;default&quot;: dj_database_url.config(default=DATABASE_URL, conn_max_age=1800),
}
</code></pre>
<p>We would need an adapter for making migrations to the <code>PostgreSQL</code> database i.e. the <code>psycopg2</code> package. This is a mandatory step if you are working with <code>postgres</code> database. This can be installed with the simple pip install:</p>
<pre><code>pip install psycopg2

# If it does not work try
pip install psycopg2-binary


# if still error persists try installing setuptools
pip install -U setuptools
pip install psycopg2
</code></pre>
<p>Now, that we have configured the database, we can now apply migrations to the new database of Postgres provided by Railway. We will simply run the migrate command and in the Railway Project the PostgreSQL database would have been modified and an appropriate schema should be applied.</p>
<pre><code>python manage.py migrate
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1652965335/blog-media/django-deployment/railway_database_migrate.png" alt="Railway PostgreSQL db migrate"></p>
<p>Make sure to update the <code>requirements.txt</code> file before pushing the project to Railway app for making sure everything works as expected. Since we have installed a few additional packages that are directly used in the <code>settings.py</code> file, we need to run the <code>pip freeze</code> command to update the <code>requiremnets.txt</code> file.</p>
<h4>Serving Static Files</h4>
<p>Now, if you have some static files like <code>CSS</code>, <code>Javascript</code>, or <code>images</code>, you need to configure the static files in order to serve them from the Railway app server. We will require another config variable for collecting the static files from the selected repository.</p>
<pre><code class="language-python">
STATIC_URL = &quot;static/&quot;
STATICFILES_DIRS = [os.path.join(BASE_DIR, &quot;static&quot;)]
STATIC_ROOT = os.path.join(BASE_DIR, &quot;staticfiles&quot;)

</code></pre>
<p>Here, if you have served your static files from the <code>static</code> folder in the root directory of your django project, you can add the above code in the settings.py file. We will collect all static files in the folder along with the default static files provided by django in the <code>staticfiles</code> directory. Run the following command if you want to test whether the static files are properly collected and served.</p>
<pre><code>python manage.py collectstatic 
</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1652600828657/JgL4nLuiL.png" alt="image.png"></p>
<p>So, this command will collect all the static files and store them in a single place. We see that the files from the admin section are also copied as well as the custom static files from the project configuration. Now, we can move on to set the config variable for the Railway app in the Variables Tab.</p>
<pre><code>DISABLE_COLLECTSTATIC = 0
</code></pre>
<p>We can set the <code>DISABLE_COLLECTSTATIC</code> variable as <code>0</code> or <code>1</code> indicating whether to disable it or not. We have currently enabled the static file collection while deploying the app but you can set it to <code>1</code> to disable the collection of static files.</p>
<p>Since I first tested the application on Railway, the static files don't work as expected, we need another package to make sure the staticfiles are served property. We will be installing the <code>whitenoise</code> package which serves as the middleware for serving the static files.</p>
<pre><code>pip install whitenoise
</code></pre>
<p>Add the whitenoise middleware <code>whitenoise.middleware.WhiteNoiseMiddleware</code> to the <code>MIDDLEWARE</code> list in the <code>settings.py</code> file.</p>
<pre><code class="language-python">MIDDLEWARE = [
...
...
...
    'whitenoise.middleware.WhiteNoiseMiddleware',
]

</code></pre>
<p>That should be enough to make the most of the deployment on Railway app. You will have to make a few adjustments as per your requirements and project.</p>
<h2>Deploy to GitHub</h2>
<p>Finally, we will have all the pieces connected, only we need to push the code to the GitHub repository in order to trigger a build on the railway app. So, we make sure we commit every critical thing that are in our django project and include every sensitiv information in the gitignore file.</p>
<pre><code>pip freeze &gt; requirements.txt
</code></pre>
<p>This step is quite important because you need to make sure that all the packages are listed in the <code>requirements.txt</code> file else you will have to wait for the build to fail and redeploy.</p>
<p>Make sure the server is running first on your local machine, remember the server will be set up from scratch but the database will already have applied migrations if you have applied migrations before after connecting the Railway database service.</p>
<pre><code>python manage.py collectstatic

python manage.py runserver
</code></pre>
<p>This will set up the origin of the remote repository that will be pushing the project code. Next, make sure to commit the code which will contain all the required stuff for deploying the code.</p>
<p>Checklist for deploying the code</p>
<pre><code>- requirements.txt
- Procfile
- runtime.txt
- django-project
- environment variables / config variables 
- static file configuration
- database configuration
- migrate schema of database 
- gitignore file for ignoring virtualenvs, .env file, staticfiles, etc
</code></pre>
<pre><code>git push origin main

origin -&gt; remote repository URL
main   -&gt; branch name
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1652966755/blog-media/django-deployment/railway-production-ss.png" alt="Deployed Railway Project"></p>
<h2>Railway CLI</h2>
<p>Railway also provides a cool CLI, it has some quite handy features like creating and managing services, local development environment, etc. We'll just dive into a few nice features of the CLI tool.</p>
<p>We'll first install the CLI on our local system, for that the guide is quite limited in a way for a few options to choose from like <code>npm</code>, <code>shell</code>, and <code>scoop</code>. For me, the shell was the way to go, but it had a few issues with permission, so I made a few changes in the <a href="https://github.com/railwayapp/cli/blob/master/install.sh">install.sh</a> script ran on my machine and it worked fine.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1652944836/blog-media/django-deployment/railway-install.png" alt="Install Railway CLI"></p>
<p>Now, that we have the <code>Railway CLi</code> set up we can run a few commands like:</p>
<pre><code>railway login

OR

# if it doesn't work
railway login --browserless

OR

# with node
npx railway login
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1652967978/blog-media/django-deployment/railway-cli-login.png" alt="Railway CLI Login"></p>
<p>This will sign you in with your railway account.</p>
<p>We now need to link our project and execute and work around the command in that project from the railway app. To link a project from a railway account, you have to use the <code>link</code> command. The link command takes in a parameter as the project id which can be obtained from the project dashboard settings tab.</p>
<pre><code>railway link &lt;proejct_id&gt;
</code></pre>
<!-- raw HTML omitted -->
<p>Now, we can explore some more commands like <code>run</code>, <code>up</code>, and so on. The <code>run</code> command is pretty solid and provides an interface to run commands for your project for instance in Django, create a superuser, manage dependencies, collectstatic files, etc. This allows us to set up and run commands into the actual production environment just from the CLI.</p>
<pre><code># railway run (command to be executed)
railway run python manage.py createsuperuser

# deploy the project
railway up
</code></pre>
<!-- raw HTML omitted -->
<p>So, this was all about creating and deploying a django application on Railway. Here are some reference links:</p>
<ul>
<li><a href="https://github.com/Mr-Destructive/django-blog/tree/railway">GitHub Repository</a></li>
<li><a href="https://django-blog-production.up.railway.app/">Live Webpage</a></li>
</ul>
<h2>Conclusion</h2>
<p>So, that's how we deploy a Django project on the Railway app. We have seen two of the many ways to create and deploy a Django application on Railway. We also integrated a PostgreSQL service provided by Railway in our Django Project. Using the Railway CLI, we were able to manage and create Projects from the terminal, also interaction of the railway project was achieved along with the production build from CLI.</p>
<p>Hopefully, you liked the above tutorial, if you have any questions. feedback, or queries, you can contact me on the Social handles provided below. Thank you for reading and till the next post Happy Coding :)</p>
