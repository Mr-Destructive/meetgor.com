{
  "title": "Django + Auth0 Quick Setup",
  "status": "published",
  "slug": "djagno-auth0-script",
  "date": "2025-04-05T12:33:52Z"
}

<h2>Introduction</h2>
<p>This is a guide and a walkthrough of how to quickly set up a base Django project with Auth0 as integration for authentication and authorization. I will walk you through the Django setup and how to use and integrate the functionalities of the Auth0.  I will also discuss how why you should be using Auth0 and why I love it.</p>
<p>The script takes <code>2:44</code> minutes time to do everything from scratch. From installing virtualenv in python to integrating the Auth0 application.</p>
<p>Here's how the script works:</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1632830813802/MOvedPYdt.gif" alt="authodj.gif"></p>
<h3>Contents</h3>
<ul>
<li><a href="#what-is-auth0">What is Auth0</a></li>
<li><a href="#why-i-love-auth0">Why I love Auth0</a></li>
<li><a href="#set-up-a-django-project">Set up a Django Project</a></li>
<li><a href="#integrate-auth0-to-a-django-project">Integrate Auth0 to a Django project</a></li>
<li><a href="#creating-a-bash-script-for-integrating-auth0">Creating a BASH Script for integrating Auth0</a>
<ul>
<li><a href="#appending-to-a-file">Appending to a file</a></li>
<li><a href="#adding-text-before-a-particular-line-using-sed">Adding text before a particular line using <code>sed</code> </a></li>
<li><a href="#appending-to-a-line-using-sed">Appending to a line using <code>sed</code></a></li>
</ul>
</li>
<li><a href="#complete-bash-script">Complete BASH Script</a></li>
<li><a href="#conclusion">Conclusion</a></li>
</ul>
<h2>What is Auth0</h2>
<p>Auth0 (<code>Auth zero</code>) is a platform that provides easy authentication and authorization for a number of platforms in various programming languages and frameworks. The easy-to-follow documentation, availability for almost all web frameworks across platforms make it a big bonus for developers. They actually make the Developer experience flawless and beginner-friendly.</p>
<p>According to Auth0,</p>
<blockquote>
<p>They make your login box awesome</p>
</blockquote>
<p>And how true is that they make things pretty convenient and wicked fast to integrate a smooth functional backend for authentication and authorization. Of course, there are more things they offer than just making authentication systems but it is by far what the world knows them for.</p>
<h2>Why I love Auth0</h2>
<p>Auth0 is a generous company that provides a free tier for a limited capacity of authentication and that might be more than enough for a developer getting its feet wet in the web development (backend).</p>
<p>They even provide a nice user interface out of the box for login/signup and even a dashboard ready-made, which is quite a lot of heavy lifting already done for you. Also, there is a dashboard for analyzing the number of sign-in/logins into the particular app. This provides the admin/developer of the app to get a closer look at the user registered in a day/week/months, number of active users, and so on.</p>
<p>So, who would not love it? I am willing to write and use their service for some of my projects. I already have used one for the Hashnode x Auth0 Hackathon, I made <a href="https://github.com/Mr-Destructive/devquotes">devquotes</a> using the authentication of Auth0 in my Django application.</p>
<h2>Set up a Django Project</h2>
<p>If you are reading this you already know how to set up a Django project, I assume. But nevertheless, I can just include a quick introduction on how to do it. I have a script to do this.</p>
<pre><code class="language-bash">#!/usr/bin/env bash

mkdir $1
cd $1
pip install virtualenv
virtualenv env
source envinctivate

pip install django
django-admin startproject $1 .
clear
</code></pre>
<p>You can check out  <a href="https://techstructiveblog.hashnode.dev/django-quick-setup-script">Django Quick Setup Script</a>  for the details of this script and also a more in-depth guide of Django project setup.</p>
<p>But if you want to understand the basics of the Django project setup here is a little guide about it:</p>
<p>Firstly, create a virtual environment, it's not mandatory but it keeps things simple and easy for your project in correspondence to the entire OS. So in python, we have a module to create the virtual environment pretty easily,</p>
<pre><code class="language-shell">pip install virtualenv
</code></pre>
<p>You can use <code>pip3</code> or <code>pip -m</code>, or however you install normal python modules. This just installs the python virtual environment, we need to create one in the current folder, so for that navigate to the folder where you want to create the project and enter the following command:</p>
<pre><code class="language-shell">virtualenv venv
</code></pre>
<p>Here, <code>venv</code> can be anything like <code>env</code> just for your understanding and simplicity it's a standard name kept for the same. After this, you will see a folder of the same name i.e. <code>venv</code> or any other name you have used. This is the folder where python will keep every installation private to the local folder itself. Now, we need to activate the virtual environment, for that we can use the command :</p>
<pre><code class="language-bash"># for Linux/macOS :
source venv/bin/activate
</code></pre>
<pre><code class="language-batch"># for Windows:
venv\Scriptsctivate
</code></pre>
<p>After this, your command prompt will have a <code>(venv)</code> attached to its start. This indicates you are in a virtual environment, things you do here, may it be module installation or any configuration related to python will stay in the local folder itself.</p>
<p>After the virtual environment is set up and activated, you can install Django and get started with it. Firstly, install Django using pip:</p>
<pre><code class="language-shell">pip install django
</code></pre>
<p>After the installation is completed, you can start a Django project in the current folder using the command:</p>
<pre><code class="language-shell">django-admin startproject name
</code></pre>
<p>Here name can be your project name. After this, you will see one new folder and one file pop up.
Namely, the <code>project named</code> folder and <code>manage.py</code> file. So you don't have to touch the <code>manage.py</code> file but we use it in most of the commands to use the Django functionalities.</p>
<p>You can now run your basic server using the command :</p>
<pre><code class="language-shell">python manage.py runserver
</code></pre>
<p>There is a base installed/setup of the Django project. Moving on in integrating the Auth0 login functionality in our webpage.</p>
<h2>Integrate the Auth0 app in your project</h2>
<p>So, for integrating the Auth0 app for your web application, you need to have an Auth0 account, you can signup here. After this you can create an Auth0 application for any type of application, we have a couple of options:</p>
<ul>
<li>Native Application</li>
<li>Single Page Application</li>
<li>Regular Web Application</li>
<li>Machine to Machine Application</li>
</ul>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1632748408741/sUeS8AgrX.png" alt="image.png"></p>
<p>You can choose as per your needs, but mostly it would be a Regular Web application if you have a backend like Django, Nodejs, PHP, or other types of frameworks and languages. So, moving ahead we have created an application for the Django framework as a backend. Now, we have a <code>Settings</code> tab in the application dashboard, where we have different credentials for the Auth0 app to talk to our application.</p>
<p>The credentials needed to be stored safely are:</p>
<ul>
<li>domain</li>
<li>Client ID (Key)</li>
<li>Client Secret</li>
</ul>
<p>This has to be secured for our local application which will go into production when ready. You can use several options like dotenv, environment variables, and so on when the application is being deployed but for now, let's hardcode them in our Django project.</p>
<p>Now, you can follow the simple straightforward procedure to copy-paste your credentials from the  <a href="https://auth0.com/docs/quickstart/webapp/django/01-login#logout">Auth0 official documentation</a>. It's quite straightforward to follow the steps even for a beginner.</p>
<p>After the Auth0 app has been configured following the procedure in the documentation, you need to integrate several files like dashboard and index templates into your custom templates.</p>
<p>Following additional changes are also to be made if you have a user-defined app for your Django project.</p>
<p>In the <code>auth0login</code> app, <code>view.py</code> file:</p>
<ol>
<li>The <code>index</code> function renders the base file for your project if the user is logged in.</li>
<li>The <code>dashboard</code> function renders the baked version of your profile/dashboard of users on your app.</li>
</ol>
<p>You would also need to add the root URIs of your app that you will be using for testing or in production. For example, we can add <code>http://127.0.0.1:8000</code> to allow and use Auth0 in our development environment locally.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1632828981455/2gc4h7tTh.png" alt="image.png"></p>
<p>You also need to specify the callback URLs for your application which is <code>/complete/auth0</code> by default.</p>
<h2>Creating a BASH Script for integrating Auth0</h2>
<p>So, we can now dive into creating the BASH Script to set up the Django x Auth0 application in minutes. The script is quite large, like 200 lines but don't worry! Its automation reduces the pain of integrating a User Authorization flawlessly. I am also thinking of adding the <code>cURL</code> command and parsing in the Client ids, keys, and secret keys, etc.</p>
<h3>Appending to a file</h3>
<p>We can use the <code>cat</code> command to append text to a file, using the syntax as below:</p>
<pre><code class="language-shell">cat &lt;&lt; EOF &gt;&gt; filename
text
more text
EOF
</code></pre>
<p>Remember here EOF is just a label to stop the command and save it to the file.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1632826339707/_g_RXP3NE.png" alt="image.png"></p>
<p>So, we can see that we were able to append to a file, multiple lines using the cat command.</p>
<p>We have used this concept in adding configuration and credentials to the <code>settings.py</code> or the <code>urls.py</code> files.</p>
<h3>Adding text before a particular line using <code>sed</code></h3>
<p><code>sed</code> is a great command, and there is nothing you can't do with it, OK there might be exceptions. We can get write to a file directly (don't display the output) and specify the line number before which we want to append the text. We can then add the text we want and followed by the filename.</p>
<pre><code class="language-shell">sed -i '33 i sometext here' filename
</code></pre>
<p>Here, <code>33</code> is the line number in the file which we want to insert before. We have used <code>'&quot;'</code> to add a <code>'</code> inside a <code>'</code>, this might feel a bit wired but that is how it is in BASH.</p>
<p>Let's say you want to add <code>print('Hello, World!')</code> to a particular line, we have to enclose <code>'</code> with these <code>&quot;'</code>( double and single quotes),</p>
<pre><code class="language-shell">sed -i '2i print('&quot;'Hello, World'&quot;')' hello.py
</code></pre>
<p>This will add the line <code>print('Hello World')</code> to the file <code>hello.py</code></p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1632824742835/Uj8AF07UG.png" alt="image.png"></p>
<h3>Appending to a line using sed</h3>
<p>We can even append text to a particular line using sed, we can use some escape characters and regex to add the text from the end of the line.</p>
<pre><code class="language-shell">sed -i '2i s/$/ textgoes here /' filename
</code></pre>
<p>Here 2 is any number of line you want to add text to, next <code>i</code> a prompt for inserting text and then we have regex like <code>s/$/ /</code>, this will put the text enclosed in <code>/ /</code> to the end of the line as indicated by <code>$</code>.  We have the filename at its usual place as before.</p>
<p>So, lets say, I want to add a comment to the second line in the previous example, I can use the following command to do it:</p>
<pre><code class="language-shell">sed -i '2 s/$/ # another comment/' hello.py

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1632825067925/0eU2mkCDI.png" alt="image.png"></p>
<p>We have used these commands to add the <code>include</code> function in the <code>urls.py</code> in the project folder.</p>
<p>So those were all the operations we used for doing some automated editing for the Auth0 app integration to our Django project.</p>
<p>Below is the entire script and is also uploaded on <a href="https://github.com/Mr-Destructive/django-auth0-quick-setup">GitHub</a>.</p>
<h2>Complete BASH Script</h2>
<p>You can run the file by parsing the name of your project.</p>
<pre><code class="language-shell">bash script.sh mywebsite
</code></pre>
<p>Wait for some 2-3 minutes, and the script will produce the Django application with the Auth0 app integrated. You will have to enter the credentials manually wherever applicable.</p>
<pre><code class="language-bash">#!/usr/bin/env bash

mkdir $1
cd $1
pip install virtualenv
virtualenv venv
source venv/Scripts/activate

pip install django
django-admin startproject $1 .

cat &lt;&lt; EOF &gt;&gt; requirements.txt
social-auth-app-django~=3.1 
python-jose~=3.0 
python-dotenv~=0.9
EOF

pip install -r requirements.txt

pip freeze &gt; requirements.txt

python manage.py startapp auth0login

touch auth0login/urls.py
mkdir auth0login/templates
touch auth0login/templates/index.html
touch auth0login/templates/dashboard.html

sed -i '40 i \    '&quot;'&quot;'social_django'&quot;'&quot;',' $1/settings.py
sed -i '41 i \    '&quot;'&quot;'auth0login'&quot;'&quot;',' $1/settings.py
sed -i '21 i \    path('&quot;''&quot;', include('&quot;'auth0login.urls'&quot;')),' $1/urls.py
sed -i '17 s/$/, include/' $1/urls.py 

cat &lt;&lt; EOF &gt;&gt; $1/settings.py
SOCIAL_AUTH_TRAILING_SLASH = False  # Remove trailing slash from routes
SOCIAL_AUTH_AUTH0_DOMAIN = 'YOUR_DOMAIN'
SOCIAL_AUTH_AUTH0_KEY = 'YOUR_CLIENT_ID'
SOCIAL_AUTH_AUTH0_SECRET = 'YOUR_CLIENT_SECRET'
EOF

cat &lt;&lt; EOF &gt;&gt; $1/settings.py 
SOCIAL_AUTH_AUTH0_SCOPE = [
    'openid',
    'profile',
    'email'
]
EOF

python manage.py migrate

cat &lt;&lt; EOF &gt;&gt;auth0login/auth0backend.py

from urllib import request
from jose import jwt
from social_core.backends.oauth import BaseOAuth2


class Auth0(BaseOAuth2):
    &quot;&quot;&quot;Auth0 OAuth authentication backend&quot;&quot;&quot;
    name = 'auth0'
    SCOPE_SEPARATOR = ' '
    ACCESS_TOKEN_METHOD = 'POST'
    REDIRECT_STATE = False
    EXTRA_DATA = [
        ('picture', 'picture'),
        ('email', 'email')
    ]

    def authorization_url(self):
        return 'https://' + self.setting('DOMAIN') + '/authorize'

    def access_token_url(self):
        return 'https://' + self.setting('DOMAIN') + '/oauth/token'

    def get_user_id(self, details, response):
        &quot;&quot;&quot;Return current user id.&quot;&quot;&quot;
        return details['user_id']

    def get_user_details(self, response):
        # Obtain JWT and the keys to validate the signature
        id_token = response.get('id_token')
        jwks = request.urlopen('https://' + self.setting('DOMAIN') + '/.well-known/jwks.json')
        issuer = 'https://' + self.setting('DOMAIN') + '/'
        audience = self.setting('KEY')  # CLIENT_ID
        payload = jwt.decode(id_token, jwks.read(), algorithms=['RS256'], audience=audience, issuer=issuer)

        return {'username': payload['nickname'],
                'first_name': payload['name'],
                'picture': payload['picture'],
                'user_id': payload['sub'],
                'email': payload['email']}

EOF

cat &lt;&lt; EOF &gt;&gt; $1/settings.py

AUTHENTICATION_BACKENDS = {
    #'YOUR_DJANGO_APP_NAME.auth0backend.Auth0',
    'django.contrib.auth.backends.ModelBackend'
}

EOF

cat &lt;&lt; EOF &gt;&gt; $1/settings.py

LOGIN_URL = '/login/auth0'
LOGIN_REDIRECT_URL = '/dashboard'
EOF

cat &gt; auth0login/views.py&lt;&lt;EOF

from django.shortcuts import render, redirect
from django.contrib.auth.decorators import login_required
from django.contrib.auth import logout as log_out
from django.conf import settings
from django.http import HttpResponseRedirect
from urllib.parse import urlencode
import json

def index(request):
    user = request.user
    if user.is_authenticated:
        return redirect(dashboard)
    else:
        return render(request, 'index.html')


@login_required
def dashboard(request):
    user = request.user
    auth0user = user.social_auth.get(provider='auth0')
    userdata = {
        'user_id': auth0user.uid,
        'name': user.first_name,
        'picture': auth0user.extra_data['picture'],
        'email': auth0user.extra_data['email'],
    }

    return render(request, 'dashboard.html', {
        'auth0User': auth0user,
        'userdata': json.dumps(userdata, indent=4)
    })

def logout(request):
    log_out(request)
    return_to = urlencode({'returnTo': request.build_absolute_uri('/')})
    logout_url = 'https://%s/v2/logout?client_id=%s&amp;%s' % \
                 (settings.SOCIAL_AUTH_AUTH0_DOMAIN, settings.SOCIAL_AUTH_AUTH0_KEY, return_to)
    return HttpResponseRedirect(logout_url)

EOF

cat &lt;&lt; EOF &gt;&gt; auth0login/templates/index.html

&lt;div class=&quot;login-box auth0-box before&quot;&gt;
    &lt;img src=&quot;https://i.cloudup.com/StzWWrY34s.png&quot; /&gt;
    &lt;h3&gt;Auth0 Example&lt;/h3&gt;
    &lt;p&gt;Zero friction identity infrastructure, built for developers&lt;/p&gt;
    &lt;a class=&quot;btn btn-primary btn-lg btn-login btn-block&quot; href=&quot;/login/auth0&quot;&gt;Log In&lt;/a&gt;
&lt;/div&gt;
EOF

cat &lt;&lt; EOF &gt;&gt; auth0login/templates/dashboard.html

&lt;div class=&quot;logged-in-box auth0-box logged-in&quot;&gt;
    &lt;h1 id=&quot;logo&quot;&gt;&lt;img src=&quot;//cdn.auth0.com/samples/auth0_logo_final_blue_RGB.png&quot; /&gt;&lt;/h1&gt;
    &lt;img class=&quot;avatar&quot; src=&quot;{{ auth0User.extra_data.picture }}&quot;/&gt;
    &lt;h2&gt;Welcome {{ user.username }}&lt;/h2&gt;
    &lt;pre&gt;{{ userdata }}&lt;/pre&gt;
&lt;/div&gt;
EOF

cat &lt;&lt; EOF &gt;&gt; auth0login/urls.py
from django.urls import path, include
from . import views

urlpatterns = [
    path('', views.index),
    path('dashboard', views.dashboard),
    path('logout', views.logout),
    path('', include('django.contrib.auth.urls')),
    path('', include('social_django.urls')),
]

EOF

python manage.py makemigrations
python manage.py migrate

</code></pre>
<h2>Conclusion</h2>
<p>Ok, so this was it, a quite big script but that's how automation can be. We were able to set up a Django base application with a ready app of Auth0 to extend the functionality. This was just a basic script also you can extend the functionalities like adding a curl command to fetch the credentials and make it more automated but that was not the aim of this article.</p>
<p>If you had any issues using the script please let me know, I'll be happy to fix those. Thanks for reading. Happy Coding :)</p>
