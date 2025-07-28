{
  "title": "Creating a Chat Application with Django and HTMX",
  "status": "published",
  "slug": "django-htmx-chat-app",
  "date": "2025-04-05T12:33:33Z"
}

<h1>Django + HTMX Chat application</h1>
<h2>Introduction</h2>
<p>In this article, we will be creating a Django project, which will be a chat-room kind of application. The user needs to authenticate to the app and then and there he/she can create or join rooms, every room will have some name and URL associated with it. So, the user simply needs to enter the name of the room, which will be unique. The user can then simply enter the messages in the chat room. This is a core chat application that uses web sockets.</p>
<p>The unique thing about this app will be that, we don't have to write a javascript client. It will all be handled by some HTMX magic. The web socket in the backend will definitely handle using Django channels.</p>
<p>Demo:</p>
<p><img src="https://meetgor-cdn.pages.dev/django-htmx/chat-app-demo.webm" alt="Demonstration of the Chat App"></p>
<p><a href="https://github.com/Mr-Destructive/django-htmx-chat">GitHub Repository</a></p>
<h3>Requirements:</h3>
<ul>
<li>
<p>Django</p>
</li>
<li>
<p>Django-channels</p>
</li>
<li>
<p>daphne</p>
</li>
<li>
<p>HTMX</p>
</li>
<li>
<p>SQLite or any relational database</p>
</li>
</ul>
<p>Also if we want to use the application on a large and production scale:</p>
<ul>
<li>
<p>Redis</p>
</li>
<li>
<p>channels_redis</p>
</li>
</ul>
<p>The code for this chat app is provided in the <a href="https://github.com/Mr-Destructive/django-htmx-chat">GitHub repository</a>.</p>
<h2>Setup for Django project</h2>
<p>We will create a simple Django project to start with. The project can have two apps, one for auth and the other for the chat. You can customize the way you want your existing project accordingly. This project is just a demonstration of the use of the chat application with websockets and Django channels.</p>
<p>I'll call the project <code>backchat</code>, you can call it whatever you want. We will create a virtual environment and install Django in that virtual environment</p>
<pre><code class="language-bash">virtualenv .venv

For Linux/macOS:
source .venv/bin/activate

For Windows:
.venv\scriptsctivate

pip install django
django-admin startproject backchat
cd backchat
</code></pre>
<p>This will set up a base Django project. We can now work on the actual implementation of the Django project. Firstly, we will start with authentication.</p>
<h2>Adding basic Authentication and Authorization</h2>
<h3>Creating the accounts app</h3>
<p>We can separate the authentication of the user from the rest of the project, by creating a separate app called <code>user</code> or <code>accounts</code> app.</p>
<pre><code class="language-bash">python manage.py startapp accounts
</code></pre>
<h3>Creating a base user model</h3>
<p>We'll start by inheriting the <a href="https://docs.djangoproject.com/en/4.1/topics/auth/customizing/#using-a-custom-user-model-when-starting-a-project">AbstractUser</a> the model provided in the <code>django.contrib.auth.models</code> module. The model has base fields like <code>username</code> and <code>password</code> which are required fields, and <code>email</code>, <code>first_name</code>, <code>last_name</code>, etc. which are not mandatory. It is better to create a custom model by inheriting the <code>AbstractUser</code> because if in the longer run, we need to make custom fields in the user model, it becomes a breeze.</p>
<pre><code class="language-python"># accounts/models.py


from djnago.contrib.auth.models import AbstractUser


class User(AbstractUser):
    pass
</code></pre>
<p>This creates a basic custom user rather than using the Django built-in user. Next, we need to make sure, Django understands the default user is the one we defined in the <code>accounts</code> app and not the default <code>User</code>. So, we just need to add a field called <code>AUTH_USER_MODEL</code> in the <code>settings.py</code> file. The value of this field will be the app name or the module name and the model in that python module that we need our custom user model to be set as the default user model.</p>
<pre><code class="language-python"># backchat/settings.py


INSALLED_APPS = [
    ...
    ...
    &quot;accounts&quot;,
]

# Append to the end of file
AUTH_USER_MODEL = 'accounts.User'
</code></pre>
<h3>Initial migrations for the Django project</h3>
<p>Now, this will get picked up as the default user model while referring to any processing related to the user. We can move into migrating the changes for the basic Django project and the user model.</p>
<pre><code class="language-bash">python manage.py makemigrations
python manage.py migrate
</code></pre>
<p><img src="https://meetgor-cdn.pages.dev/django-htmx/chat-accounts-migrations.png" alt="initial migration for base django and user model"></p>
<h3>Creating register view</h3>
<p>Further, we can add the views like register and profile for the accounts app that can be used for the basic authentication. The Login and Logout views are provided in the <code>django.contrib.auth.views</code> module, we only have to write our own register view. I will be using function-based views to keep it simple to understand but it can simply be a class-based view as well.</p>
<p>To define a view first, we need to have form, a user registration form. The form will inherit from the <a href="https://docs.djangoproject.com/en/4.1/topics/auth/default/#django.contrib.auth.forms.UserCreationForm">UserCreationForm</a> form which will make the bulk of the task for us at the user registration. We can simply then override the Meta class with the fields that we want to display, so hence we just keep the <code>username</code> and the <code>password</code> fields. The form can be customized by adding in the widget attribute and specifying the classes used in them.</p>
<pre><code class="language-python"># accounts/forms.py


from accounts.models import User
from django.contrib.auth.forms import UserCreationForm

class UserRegisterForm(UserCreationForm):

    class Meta:
        model= User
        fields = ['username', 'password1', 'password2']
</code></pre>
<p>This will give us the <code>UserRegisterForm</code> form that can be used for displaying in the register view that we will create in the next step.</p>
<p>We will have to create the register view that will render the form for user registration and will also process the form submission.</p>
<pre><code class="language-python"># accounts/views.py


from django.contrib import messages
from django.shortcuts import redirect, render
from accounts.forms import UserRegisterForm

def register(request):
    if request.method == &quot;POST&quot;:
        form = UserRegisterForm(request.POST)
        if form.is_valid():
            form.save()
            username = form.cleaned_data.get(&quot;username&quot;)
            messages.success(
                request, f&quot;Your account has been created! You are now able to log in&quot;
            )
            return redirect(&quot;login&quot;)
    else:
        form = UserRegisterForm()
        return render(request, &quot;accounts/register.html&quot;, {&quot;form&quot;: form})
</code></pre>
<p>The above register view has two cases, one for the user requesting the registration form and the second request when the user submits the form. So, when the user makes a get request, we load an empty form <code>UserRegisterForm</code> and render the <code>register</code> template with the form. We will make the templates later.</p>
<p>So, the template is just rendered when the user wants to register and when the user submits the form(sends a post request) we get the details from the post request and make it an instance of <code>UserRegisterForm</code> and save the form if it is valid. We then redirect the user to the login view (will use the default one in the next section). We parse the message as the indication that the user was created.</p>
<h3>Adding URLs for Authentication and Authorization</h3>
<p>Once, we have the register view setup, we can also add login and logout views in the app. But, we don't have to write it ourselves, we can override them if needed, but we'll keep the default ones. We will use the <a href="https://docs.djangoproject.com/en/4.1/topics/auth/default/#django.contrib.auth.views.LoginView">LoginView</a> and the <a href="https://docs.djangoproject.com/en/4.1/topics/auth/default/#django.contrib.auth.views.LogoutView">LogoutView</a> view which are class-based views provided in the <code>django.contrib.auth.views</code> module. We will provide the respective templates for each of these views.</p>
<pre><code class="language-python"># accounts/urls.py


from django.urls import path
from django.contrib.auth import views as auth_views
import user.views as user_views

urlpatterns = [
    path(&quot;register/&quot;, user_views.register, name=&quot;register&quot;),
    path(
        &quot;login/&quot;,
        auth_views.LoginView.as_view(template_name=&quot;accounts/login.html&quot;),
        name=&quot;login&quot;,
    ),
    path(
        &quot;logout/&quot;,
        auth_views.LogoutView.as_view(template_name=&quot;accounts/logout.html&quot;),
        name=&quot;logout&quot;,
    ),
]
</code></pre>
<p>We have named the URLs as <code>register</code>, <code>login</code>, and <code>logout</code> so that we can use them while rendering the links for them in the templates. Now, we also need to include the URLs from the accounts app in the project URLs. We can do that by using the <code>include</code> method and specifying the app name with the module where the urlpatterns are located.</p>
<pre><code class="language-python"># backchat/urls.py


from django.contrib import admin
from django.urls import include, path

urlpatterns = [
    path(&quot;admin/&quot;, admin.site.urls),
    path(&quot;auth/&quot;, include(&quot;accounts.urls&quot;)),
]
</code></pre>
<p>So, we have routed the <code>/auth</code> path to include all the URLs in the accounts app. So, the login view will be at the URL <code>/auth/login/</code> , and so on.</p>
<p>Also, we need to add the <code>LOGIN_REDIRECT_URL</code> and <code>LOGIN_URL</code> for specifying the url name which will be redirected once the user is logged in and the default login url name respectively.</p>
<pre><code class="language-python"># backchat / settings.py


LOGIN_REDIRECT_URL = &quot;index&quot;
LOGIN_URL = &quot;login&quot;
</code></pre>
<p>We are now almost done with the view and routing part of the accounts app and can move into the creation of templates.</p>
<h3>Adding Templates for authentication views</h3>
<p>We need a few templates that we have been referencing in the views and the URLs of the accounts app in the project. So there are a couple of ways to use templates in a Django project. I prefer to have a single templates folder in the root of the project and have subfolders as the app which will have the templates specific to those apps.</p>
<p>I usually create a <code>base.html</code> file in the templates folder and use that for inheriting other templates. So, we will have to change one setting in the project to make sure it looks at <code>templates/</code> from the root of the project.</p>
<pre><code class="language-bash"># backchat/settings.py

import os

...
...

TEMPLATES = [
    {
        &quot;BACKEND&quot;: &quot;django.template.backends.django.DjangoTemplates&quot;,
        &quot;DIRS&quot;: [ os.path.join(BASE_DIR, &quot;templates&quot;), ],
        &quot;APP_DIRS&quot;: True,
        &quot;OPTIONS&quot;: {
            &quot;context_processors&quot;: [
                &quot;django.template.context_processors.debug&quot;,
                &quot;django.template.context_processors.request&quot;,
                &quot;django.contrib.auth.context_processors.auth&quot;,
                &quot;django.contrib.messages.context_processors.messages&quot;,
            ],
        },
    },
]
</code></pre>
<p>Then create the folder in the same path as you see your <code>manage.py</code> file.</p>
<pre><code class="language-bash">mkdir templates
</code></pre>
<p><img src="https://meetgor-cdn.pages.dev/django-htmx/chat-mkdir-templates.png" alt="Template Set Up"></p>
<h4>Creating the base template</h4>
<p>The below will be the base template used for the chat application, you can custmize it as per your needs.</p>
<pre><code class="language-html">&lt;!DOCTYPE html&gt;
&lt;html&gt;
    &lt;head&gt;
        &lt;meta charset=&quot;utf-8&quot; /&gt;
        &lt;title&gt;Chat App&lt;/title&gt;
        {% load static %}
        &lt;script src=&quot;https://unpkg.com/htmx.org@1.8.5&quot;&gt;&lt;/script&gt;
    &lt;/head&gt;
    &lt;body&gt;
        {% if user.is_authenticated %}
            &lt;a href=&quot;{% url 'logout' %}&quot;&gt;Logout&lt;/a&gt;
        {% else %}
            &lt;a href=&quot;{% url 'login' %}&quot;&gt;Login&lt;/a&gt;
        {% endif %}
        &lt;h1&gt;Back Chat&lt;/h1&gt;
        {% block base %}
        {% endblock %}
    &lt;/body&gt;
&lt;/html&gt;
</code></pre>
<p>We have included the <a href="https://htmx.org/docs/#installing">htmx</a> package as the script into this template as we will be using it in the actual part of chat application.</p>
<h4>Creating the Register Template</h4>
<pre><code class="language-html"># templates / accounts / register.html


{% extends 'base.html' %}
{% block base %}
    &lt;div class=&quot;content-section&quot;&gt;
        &lt;form method=&quot;POST&quot;&gt;
            {% csrf_token %}
            &lt;fieldset class=&quot;form-group&quot;&gt;
                &lt;legend class=&quot;border-bottom mb-4&quot;&gt;Register Now&lt;/legend&gt;
                {{ form.as_p }}
            &lt;/fieldset&gt;
            &lt;div class=&quot;form-group&quot;&gt;
                &lt;button class=&quot;btn btn-outline-info&quot; type=&quot;submit&quot;&gt;Sign Up&lt;/button&gt;
            &lt;/div&gt;
        &lt;/form&gt;
        &lt;div class=&quot;border-top pt-3&quot;&gt;
            &lt;small class=&quot;text-muted&quot;&gt;
		    Already Have An Account? &lt;a class=&quot;ml-2&quot; href=&quot;{% url 'login' %}&quot;&gt;Log In&lt;/a&gt;
            &lt;/small&gt;
        &lt;/div&gt;
    &lt;/div&gt;
{% endblock %}
</code></pre>
<p><img src="https://meetgor-cdn.pages.dev/django-htmx/chat-register-page.png" alt="User Registration Page"></p>
<h4>Creating the Login Template</h4>
<pre><code class="language-html"># templates / accounts / login.html    


{% extends 'base.html' %}
{% block base %}
    &lt;div class=&quot;content-section&quot; id=&quot;login&quot;&gt;
        &lt;form method=&quot;POST&quot; &gt;
            {% csrf_token %}
            &lt;fieldset class=&quot;form-group&quot;&gt;
                &lt;legend class=&quot;border-bottom mb-4&quot;&gt;LOG IN&lt;/legend&gt;
                {{ form.as_p }}
            &lt;/fieldset&gt;
            &lt;div class=&quot;form-group&quot;&gt;
                &lt;button class=&quot;btn btn-outline-info&quot; type=&quot;submit&quot;&gt;Log In&lt;/button&gt;
            &lt;/div&gt;
        &lt;/form&gt;
        &lt;div class=&quot;border-top pt-3&quot;&gt;
            &lt;small class=&quot;text-muted&quot;&gt;
                Register Here &lt;a class=&quot;ml-2&quot; href=&quot;{% url 'register' %}&quot;&gt;Sign Up&lt;/a&gt;
            &lt;/small&gt;
        &lt;/div&gt;
    &lt;/div&gt;
{% endblock %}
</code></pre>
<p><img src="https://meetgor-cdn.pages.dev/django-htmx/chat-login-page.png" alt="User Login Page"></p>
<p>Creating the Logout Template</p>
<pre><code class="language-html"># templates / accounts / logout.html    


{% extends 'base.html' %}
{% block base %}
    &lt;h2&gt;You have been logged out&lt;/h2&gt;
    &lt;div class=&quot;border-top pt-3&quot;&gt;
        &lt;small class=&quot;text-muted&quot;&gt;
            &lt;a href=&quot;{% url 'login' %}&quot;&gt;Log In Again&lt;/a&gt;
        &lt;/small&gt;
    &lt;/div&gt;
{% endblock %}
</code></pre>
<h2>Install and setup channels</h2>
<p>We will be using channels to create long-running connections, it is a wrapper around Django's asynchronous components and allows us to incorporate other protocols like web sockets and other asynchronous protocols.</p>
<p>So, we will be using the Django channels package that will allow us to use the WebSocket protocol in the chat application. <a href="https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API">WebSocket</a> is a communication protocol(set of rules and standards to be followed) that allows both the client and server to send and receive messages(communication).</p>
<p>To install Django channels, we can use pip and install channels with daphne which will be used as the web server gateway interface for asynchronous web applications.</p>
<pre><code class="language-bash">pip install -U channels[&quot;daphne&quot;]
</code></pre>
<p>So this will install the latest version of channels and daphne. We now have <a href="https://channels.readthedocs.io/en/stable/installation.html">channels</a> in our Django project, we simply need to configure the <code>CHANNEL_LAYER</code> config for specifying the backend used that can be <code>Redis</code>, <code>In-Memory</code>, or others. We need to add <code>channels</code> , <code>daphne</code> to the <code>INSALLED_APPS</code> config of the project. Make sure the <code>daphne</code> app is on top of the applications list.</p>
<pre><code class="language-bash"># backchat/settings.py

...
...

INSALLED_APPS = [
    &quot;daphne&quot;,
    ...
    ...
    &quot;channels&quot;,
]


ASGI_APPLICATION = &quot;backchat.asgi.application&quot;

...
...

# For InMemory channels

CHANNEL_LAYERS = {
    'default': {
        &quot;BACKEND&quot;: &quot;channels.layers.InMemoryChannelLayer&quot;,
    }
}


# For Redis

CHANNEL_LAYERS = {
    &quot;default&quot;: {
        &quot;BACKEND&quot;: &quot;asgi_redis.RedisChannelLayer&quot;,
        &quot;CONFIG&quot;: {
            &quot;hosts&quot;: [(&quot;redis-host-url&quot;, 6379)],
        },
    },
}
</code></pre>
<h3>Redis Configuration (Optional)</h3>
<p>You can either use the <a href="https://channels.readthedocs.io/en/latest/topics/channel_layers.html">InMemoryChannelLayer</a> or you can use them <code>RedisChannelLayer</code> for the backend of the chat app. There are other types of backends like <code>Amazon SQS</code> services, <code>RabbitMQ</code>, <code>Kafka</code>, <code>Google Cloud Pub/Sub</code>, etc. I will be creating the app with only the <code>InMemoryChannelLayer</code> but will provide a guide for redis as well, both are quite similar and only have a few nuances.</p>
<p>We need to install <a href="https://github.com/django/channels_redis/">channels_redis</a> it for integrating redis in the Django project with channels.</p>
<pre><code class="language-bash">pip install channels_redis
</code></pre>
<p>So, this will make the <code>channels_redis</code> package available in the project, we use this package for real-time storage, in the case of the chat app, we might use it for storing messages or room details, etc.</p>
<h2>Creating the Chat App</h2>
<p>Further, we now can create another app for handling the rooms and chat application logic. This app will have its own models, views, and URLs. Also, we will define consumers and routers just like views and URLs but for asynchronous connections. More on that soon.</p>
<p>So, let's create the <code>chat</code> app.</p>
<pre><code class="language-bash">python manage.py startapp chat
</code></pre>
<p>Then we will add the chat app to the <code>INSALLED_APPS</code> config.</p>
<pre><code class="language-python"># backchat/settings.py

INSALLED_APPS = [
    ...
    ...,
    &quot;chat&quot;,
]
</code></pre>
<p>So, this will make sure to load the chat app in the project. Whenever we run any commands like migrations or running the server, it keeps the app in the <code>INSALLED_APPS</code> checked up.</p>
<h3>Defining models</h3>
<p>This is optional, but we'll do it for since we are making a Django app. We already have an auth system configured, adding rooms and authorizing the users will become easier then.</p>
<p>So, let's create the models for the chat app which will be the <code>Room</code>.</p>
<pre><code class="language-python"># chat/models.py


from django.db import models
from accounts.models import User

class Room(models.Model):
    name = models.CharField(max_length=128)
    slug = models.SlugField(unique=True)
    users = models.ManyToManyField(User)

    def __str__(self):
        return self.name


class Message(models.Model):
    room = models.ForeignKey(Room, on_delete=models.CASCADE)
    user = models.ForeignKey(User, on_delete=models.CASCADE)
    message = models.TextField()
    created_at = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return (
            self.room.name + &quot; - &quot; +
            str(self.user.username) + &quot; : &quot; +
            str(self.message)
        )
</code></pre>
<p>So, we simply have the name which will be taken from the user, and the slug which will be served as the URL to the room. The users are set as <a href="https://docs.djangoproject.com/en/4.1/ref/models/fields/#django.db.models.ManyToManyField">ManyToManyField</a> since one room can have multiple users and a user can be in multiple rooms. Also we are creating the model <code>Message</code> that will be storing the room and the user as the foreign key and the actual text as the message, we can improve the security by encrypting the text but it's not the point of this article.</p>
<p>We have set the <code>created_at</code> the field which will set the time when the object was created. Finally, the dunder string method is used for representing the message object as a price of the concatenation of strings of room name, username, and the message. This is useful for admin stuff as it makes it easier to read the object rather than the default id.</p>
<p>Now, once the models are designed we can migrate the models into the database.</p>
<pre><code>python manage.py makemigrations
python manage.py migrate
</code></pre>
<p><img src="https://meetgor-cdn.pages.dev/django-htmx/chat-migrations.png" alt="Chat app migrations"></p>
<p>We now have a decent model structure for the chat application. We can now start the crux of the application i.e. consumers and routing with channels.</p>
<h3>Writing consumers and routers for WebSockets</h3>
<p>So, we start with the bare bones provided in the tutorial on the channel <a href="https://channels.readthedocs.io/en/stable/tutorial/part_3.html">documentation</a>. We create a class (consumer) called <code>ChatConsumer</code> which inherits from the <code>AsyncWebsocketConsumer</code> provided by the <code>channels.generic.websocket</code> module. This has a few methods like <code>connect</code>, <code>disconnect</code>, and <code>receive</code>. These are the building blocks of a consumer. We define these methods as they will be used for communication via the WebSocket protocol through the channels interface.</p>
<p>In the following block of code, we are essentially doing the following:</p>
<ul>
<li>
<p>Accepting connection on the requested room name</p>
</li>
<li>
<p>Sending and Receiving messages on the room/group</p>
</li>
<li>
<p>Closing the WebSocket connection and removing the client from the room/group</p>
</li>
</ul>
<pre><code class="language-python"># chat/consumers.py


import json

from asgiref.sync import sync_to_async
from channels.generic.websocket import AsyncWebsocketConsumer

from chat.models import Room, Message


class ChatConsumer(AsyncWebsocketConsumer):
    async def connect(self):
        self.room_name = self.scope[&quot;url_route&quot;][&quot;kwargs&quot;][&quot;room_slug&quot;]
        self.room_group_name = f&quot;chat_{self.room_name}&quot;
        self.user = self.scope[&quot;user&quot;]

        await self.channel_layer.group_add(
            self.room_group_name, self.channel_name
        )

        await self.accept()

    async def disconnect(self, close_code):
        await self.channel_layer.group_discard(
            self.room_group_name, self.channel_name
        )

    async def receive(self, text_data):
        text_data_json = json.loads(text_data)
        message = text_data_json[&quot;message&quot;]
        username = self.user.username
        
        await self.channel_layer.group_send(
            self.room_group_name, 
            {
                &quot;type&quot;: &quot;chat_message&quot;,
                &quot;message&quot;: message,
                &quot;username&quot;: username,
            }
        )

    async def chat_message(self, event):
        message = event[&quot;message&quot;]
        username = event[&quot;username&quot;]

        await self.send(
            text_data=json.dumps(
                {
                    &quot;message&quot;: message,
                    &quot;username&quot;: username
                }
            )
        )
</code></pre>
<h4>Accept the WebSocket connection</h4>
<p>Here, room and group more or less mean the same thing but are different in different contexts. For instance, the group refers to the collection of clients which are connected to a channel(communication between consumer and web socket) and the Room is referring to the collection of clients connected to a single connection stream like a literal room. So we can say, the group is a technical term and the room is more of a layman's term for the same thing.</p>
<p>The method <code>connect</code> is called when the client establishes a websocket connection. When that happens, the function gets the room slug from the URL of the client and stores <code>room_name</code> which is a string. It creates a separate variable called <code>room_group_name</code> by prepending the <code>chat_</code> string to the <code>room_name</code>, it also gets the currently logged-in user from the request. It then adds the <code>channel_name</code> to the group named <code>room_group_name</code>. The <code>channel_name</code> is a unique identifier to the connection/consumer in the channel. By adding the <code>channel_name</code>, the consumer then can broadcast the message to all the channels within the group. Finally, the function accepts the connection, and a <strong>webcoket connection is established from both ends, connection is sent from the client and is now accepted from the backend.</strong></p>
<h4>Disconnect from the WebSocket connection</h4>
<p>When the client sends a close connection request to the server, the <code>disconnect</code> method is triggered and it basically removes the <code>channel_name</code> from the group i.e. the group name <code>room_group_name</code> whatever the string it happens to be. So, it basically removes the client from the broadcast group and hence it can't receive or send messages through the websocket since it has been closed from both ends.</p>
<p>If you would have paid attention to the <code>close_code</code> in-method parameter, it is not being used currently. However, we can use it for checking why the connection was closed, as the <code>close_code</code> is a numeric value just like the status code in the web request for letting the server know the reason for disconnecting from the client.</p>
<h4>Receive a message from the WebSocket connection</h4>
<p>The <code>recieve</code> method is the core of the consumer as it is responsible for all the logic and parsing of the message and broadcasting of the messages from the clients to the group channels. The function takes in a parameter called <code>text_data</code> and it is sent from the client through websocket, so it is JSON content. We need to get the actual message from the JSON object or any other piece of content from the client. So, we deserialize (convert the JSON object to python objects) the received payload, and get the value from the key <code>message</code>. The key is the input name or id from the client sending the request through the web socket, so it can be different depending on the frontend template(we'll see the front end soon as well).</p>
<p>We get the user from the scope of the consumer as we previously initialized it in the connect method. This can be used for understanding which user has sent the message, it will be used later on as we send/broadcast the message to the group.</p>
<p>Now, the final piece in the receive method is the <code>channel_layer.group_send</code> method, this method as the name suggests is used to send or broadcast the received message to the entire group. The method has two parameters:</p>
<ol>
<li>
<p>The name of the group</p>
</li>
<li>
<p>The JSON body containing the message and other details</p>
</li>
</ol>
<p>The method is not directly sending the message but it has a type key in the JSON body which will be the function name to call. The function will simply call the other funciton mentioned in the type key in the dict. The following keys in the dict will be the parameters of that funciton. In this case, the funciton specified in the <code>type</code> key is <code>chat_message</code> which takes in the <code>event</code> as the parameter. This event will have all the parameters from the <code>group_send</code> method.</p>
<p>So, the <code>chat_message</code> will load in this message, username, and the room name and then call the <code>send</code> method which actually sends the message as a JSON payload to the WebSocket connection which will be received by all the clients in the same group as provided in the <code>room_group_name</code> string.</p>
<h3>Adding Routers for WebSocket connections</h3>
<p>So, till this point have consumers, which are just like views in terms of channels. Now, we need some URL routes to map these consumers to a path. So, we will create a file/module called <code>routing.py</code> which will look quite similar to the <code>urls.py</code> file. It will have a list called <code>websocket_urlpatterns</code> just like <code>urlpatterns</code> with the list of <code>path</code>. These paths however are not <code>http</code> routes but will serve for the <code>WebSocket</code> path.</p>
<pre><code class="language-python"># chat / routing.py


from django.urls import path

from chat import consumers

websocket_urlpatterns = [
    path('chat/&lt;str:room_slug&gt;/', consumers.ChatConsumer.as_asgi()),
]
</code></pre>
<p>In the above code block, we have defined a URL for the web socket with the path as <code>/chat/&lt;room_slug&gt;</code> where room_name will be the <code>slug</code> for the room. The path is bound with the consumer-defined in the <code>consumers.py</code> module as <code>ChatConsumer</code>. The <code>as_asgi</code> method is used for converting a view into an ASGI-compatible view for the WebSocket interface.</p>
<h3>Setting up ASGI Application</h3>
<p>We are using the ASGI application config rather than a typical WSGI application which only works one request at a time. We want the chat application to be asynchronous because multiple clients might send and receive the messages at the same time, we don't want to make a client wait before the server process a message from another client, that's just the reason we are using WebSocket protocol.</p>
<p>So, we need to also make sure, it makes the http request and also add our websocket config from the chat app we created in the previous sections. So, inside the <code>asgi.py</code> file in the project config module, we need to make some changes to include the chat application configurations.</p>
<pre><code class="language-python"># backchat / asgi.py


import os
from django.core.asgi import get_asgi_application
from channels.auth import AuthMiddlewareStack
from channels.routing import ProtocolTypeRouter, URLRouter

os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'backchat.settings')

from chat import routing


application = ProtocolTypeRouter({
    'http': get_asgi_application(),
    &quot;websocket&quot;:AuthMiddlewareStack(
        URLRouter(
            routing.websocket_urlpatterns
        )
    )
})
</code></pre>
<p>We will override the <code>application</code> config which is a component used for routing different types of protocols for the <code>ASGI</code> application. We have set the two keys, <code>http</code> and <code>websocket</code> in our application. The <code>http</code> type of requests will be served with the <code>get_asgi_application</code> application which is used for running the application in an ASGI environment.</p>
<p>For <code>websocket</code> type of requests, we are setting the <a href="https://channels.readthedocs.io/en/latest/topics/authentication.html">AuthMiddlewareStack</a> which helps in authenticating the users requesting the WebSocket connection and allows only authorized users to make a connection to the application. The <a href="https://channels.readthedocs.io/en/stable/topics/routing.html">URLRouter</a> is used for mapping the list of URL patterns with the incoming request. So, this basically serves the request URL with the appropriate consumer in the application. We are parsing in the <code>websocket_urlpatterns</code> as the list of URLs that will be used for the WebSocket connections.</p>
<p>Now, we run the server, we should be seeing the <code>ASGI</code> server serving our application rather than the plain WSGI application.</p>
<pre><code>$ python manage.py runserver

Watching for file changes with StatReloader
Performing system checks...

System check identified no issues (0 silenced).
February 05, 2023 - 09:22:45
Django version 4.1.5, using settings 'backchat.settings'
Starting ASGI/Daphne version 4.0.0 development server at http://127.0.0.1:8000/
Quit the server with CONTROL-C.
</code></pre>
<p>The application is not complete yet, it might not have most components working functional yet. So, we'll now get into making the user interfaces for the application, to create, join, and view rooms in the application.</p>
<h3>Adding Views for Chat Rooms</h3>
<p>We will have a couple of views like create room page, the join room page, and the chat room page. We will define each view as a distinct view and all of them will require authenticated users.</p>
<pre><code class="language-python"># chat / views.py


import string
import random
from django.contrib.auth.decorators import login_required
from django.shortcuts import render, reverse, redirect
from django.utils.text import slugify
from chat.models import Room


@login_required
def index(request, slug):
    room = Room.objects.get(slug=slug)
    return render(request, 'chat/room.html', {'name': room.name, 'slug': room.slug})

@login_required
def room_create(request):
    if request.method == &quot;POST&quot;:
        room_name = request.POST[&quot;room_name&quot;]
        uid = str(''.join(random.choices(string.ascii_letters + string.digits, k=4)))
        room_slug = slugify(room_name + &quot;_&quot; + uid)
        room = Room.objects.create(name=room_name, slug=room_slug)
        return redirect(reverse('chat', kwargs={'slug': room.slug}))
    else:
        return render(request, 'chat/create.html')

@login_required
def room_join(request):
    if request.method == &quot;POST&quot;:
        room_slug = request.POST[&quot;room_slug&quot;]
        room = Room.objects.get(slug=room_slug)
        return redirect(reverse('chat', kwargs={'slug': room.slug}))
    else:
        return render(request, 'chat/join.html')
</code></pre>
<p>In the above views module, we have added 3 views namely <code>index</code> as the room page, <code>room_create</code> for the room creation page, and the <code>room_join</code> for the room join page. The index view is a simple get request to the provided slug of the room, it gets the slug from the URL from the request and fetches an object of the room associated with that slug. Then it renders the room template with the context variables like the name of the room and the slug associated with that room.</p>
<p>The <code>room_create</code> view is a simple two-case view that either can render the room creation page or process the submitted form and create the room. Just like we used in the <code>register</code> view in the accounts app. When the user will send a <code>GET</code> request to the URL which we will map to <code>/create/</code> shortly after this, the user will be given a form. So, we will render the <code>create.html</code> template. We will create the html template shortly.
If the user has sent a <code>POST</code> request to the view via the <code>/create</code> URL, we will fetch the name field in the sent request and create a unique identifier with the name of the room. We will slugify the concatenation of the name with the uid and will set it as the slug of the room. We will then simply create the room and redirect the user to the <code>room</code> page.</p>
<p>The <code>room_join</code> view also is a two-case view, where the user can either request the join room form or send a slug with the form submission. If the user is requesting a form, we will render the <code>join.html</code> template. If the user is submitting the form, we will fetch the room based on the slug provided and redirect the user to the <code>room</code> page.</p>
<p>So, the <code>room_join</code> and <code>room_create</code> views are quite similar, we are fetching an already existing room in the case of the join view and creating a new instance of room in the create view. Now, we will connect the views to the URLs and finally get to the templates.</p>
<h3>Connecting Views and URLs</h3>
<p>We have three views to route to a URL. But we will also have one additional URL which will be the home page for the application, on that page the user can choose to either join or create a room. We have the room creation, join the room and the room view to be mapped in this URL routing of the app.</p>
<pre><code class="language-python"># chat / urls.py


from django.urls import path
from django.views.generic import TemplateView
from chat import views


urlpatterns = [
    path(&quot;&quot;, TemplateView.as_view(template_name=&quot;base.html&quot;), name='index'),
    path(&quot;room/&lt;str:slug&gt;/&quot;, views.index, name='chat'),
    path(&quot;create/&quot;, views.room_create, name='room-create'),
    path(&quot;join/&quot;, views.room_join, name='room-join'),
]
</code></pre>
<p>So, the first route is the home page view called <code>index</code>, we have used the <a href="https://docs.djangoproject.com/en/4.1/ref/class-based-views/base/#templateview">TemplateView</a> which will simply render the template provided. We don't have to create a separate view for just rendering a template. We already have defined the <code>base.html</code> template while setting up the <code>accounts</code> app. This will be the same template, we will add some more content to the template later on. The URL mapped is the <code>/</code> since we will add the URLs of this app to the project URLs with an empty <code>&quot;&quot;</code> path.</p>
<p>The second route is used for the room index page, i.e. where the user will be able to send and receive messages. The path is assigned as <code>/room/&lt;str:slug&gt;/</code> indicating a parameter called slug of type string will be used in accessing a particular room. The URL will be bound to the index view that we created in the views file, which fetches the room based on the slug, so here is where the slug will be coming from. The name of the URL-View route will be <code>chat</code> but you can keep it as per your requirements. The URL name is really handy for use in the templates.</p>
<p>The third route is for the room creation page. The <code>/create/</code> URL will be bound to the <code>room_create</code> view, as we discussed, it will serve two purposes, one to render the form for creating the room, and the other for sending a <code>POST</code> request to the same path for the creating the Room with the name provided. The name is not required but helps in identifying and making it user-friendly. The name of this URL is set as <code>room-create</code>.</p>
<p>The final route is for joining the room, the <code>/join/</code> URL will be triggering the <code>room_join</code> view. Similar to the <code>room-create</code> URL, the URL will render the join room form on a <code>GET</code> request, fetch the room with the provided slug and redirect to the room page. Here, the slug field in the form will be required for actually finding the matching room. The name of the URL route is set as <code>room-join</code>.</p>
<p>We now add the URLs from the chat app to the project URLs. This will make the <code>/</code> as the entry point for the chat application URLs.</p>
<pre><code class="language-python"># backchat / urls.py


from django.contrib import admin
from django.urls import path, include

urlpatterns = [
    path(&quot;admin/&quot;, admin.site.urls),
    path(&quot;auth/&quot;, include('accounts.urls')),
    path(&quot;&quot;, include('chat.urls')),
]
</code></pre>
<p>Hence the process is completed for the backend to process the message, it then is dependent on the client to process and render the message.</p>
<p>Till HTMX was not a thing!</p>
<p>We won't have to write a single line of javascript to receive and handle the WebSocket connection!</p>
<h3>Creating Templates and adding htmx</h3>
<p>We now move into the actual frontend or creating the template for actually working with the rooms and user interaction. We will have three pieces of templates, a room creates the page, a room join page, and a room chat page. As these template names suggest, they will be used for creating a room with the name, joining the room with the slug, and the room chat page where the user will send and receive messages.</p>
<p>Let/s modify the base template first.</p>
<pre><code class="language-html">&lt;!DOCTYPE html&gt;
&lt;html&gt;
    &lt;head&gt;
        &lt;meta charset=&quot;utf-8&quot; /&gt;
        &lt;title&gt;Chat App&lt;/title&gt;
        {% load static %}
        &lt;script src=&quot;https://unpkg.com/htmx.org@1.8.5&quot;&gt;&lt;/script&gt;
    &lt;/head&gt;
    &lt;a href=&quot;{% url 'index' %}&quot;&gt;Home&lt;/a&gt;
        {% if user.is_authenticated %}
            &lt;a href=&quot;{% url 'logout' %}&quot;&gt;Logout&lt;/a&gt;
        {% else %}
            &lt;a href=&quot;{% url 'login' %}&quot;&gt;Login&lt;/a&gt;
        {% endif %}
    &lt;body hx-ext=&quot;ws&quot;&gt;
        &lt;h1&gt;Back Chat&lt;/h1&gt;
        {% block base %}
            &lt;a href=&quot;{% url 'room-join' %}&quot;&gt;Join Room&lt;/a&gt;
            &lt;a href=&quot;{% url 'room-create' %}&quot;&gt;Create Room&lt;/a&gt;
        {% endblock %}
    &lt;/body&gt;
&lt;/html&gt;
</code></pre>
<p><img src="https://meetgor-cdn.pages.dev/django-htmx/chat-home-page.png" alt="Chat App Home Page"></p>
<h4>Create Room Template</h4>
<p>We will have to render the form with a field like <code>name</code> for setting it as the name of the room, it is not required but again, it makes it easier for the user to find the name of the room a bit more friendly than random characters.</p>
<pre><code class="language-html"># templates / chat / create.html


{% extends 'base.html' %}

{% block base %}
    &lt;form method='post' action='{% url 'room-create' %}'&gt;
        {% csrf_token %}
        &lt;input name='room_name' id='room_name' placeholder='Room Name'&gt;
        &lt;input type='submit' id=&quot;submit&quot;&gt;
    &lt;/form&gt;
{% endblock %}
</code></pre>
<p><img src="https://meetgor-cdn.pages.dev/django-htmx/chat-create-room-page.png" alt="Chat Room Create Page"></p>
<p>The template is inherited from the <code>base.html</code> template and we render a form with the <code>room_name</code> input. The form will be submitted to the URL named <code>room-create</code> hence it will send a <code>POST</code> request to the server where it will create the room and further process the request.</p>
<h4>Join Room Template</h4>
<p>The join room template is similar to the create room template except it gets the slug of the room from the user rather than the name which is not a unique one to join the room.</p>
<pre><code class="language-html"># templates / chat / join.html


{% extends 'base.html' %}

{% block base %}
    &lt;form method='post' action='{% url 'room-join' %}'&gt;
        {% csrf_token %}
        &lt;input name='room_slug' id='room_slug' required='true' placeholder='Room Code'&gt;
        &lt;input type='submit' id=&quot;submit&quot;&gt;
    &lt;/form&gt;
{% endblock %}
</code></pre>
<p><img src="https://meetgor-cdn.pages.dev/django-htmx/chat-join-room-page.png" alt="Chat Room Join Page"></p>
<p>The form will be submitted to the URL named <code>room-join</code> hence it will send a <code>POST</code> request to the server where it will fetch the room and further process the request.</p>
<h3>Room Template (HTMX code)</h3>
<p>Now, time for the actual ingredient in the application, some HTMX magic!</p>
<p>This template, as the two templates above inherit from the base template, that's the same thing. But it has a special <code>div</code> with the attribute <a href="https://htmx.org/attributes/hx-ws/">hx-ws</a> which is used for using attributes related to the web socket in the htmx library. The <code>connect</code> value is used for connecting to a WebSocket. The value of the attribute must be set to the URL of the WebSocket to be connected. In our case, it is the URL path from the <code>routing</code> app as <code>/chat/&lt;room_slug&gt;/</code>. This simply will connect the client to the WebSocket from the backend. The other important attribute is the <code>send</code> which is used for sending a message to the connected web socket.</p>
<pre><code class="language-html"># templates / chat / room.html


{% extends 'base.html' %}

{% block base %}
    &lt;h2&gt;{{ name }}&lt;/h2&gt;
    &lt;div hx-ws=&quot;connect:/chat/{{ slug }}/&quot;&gt;
        &lt;form hx-ws=&quot;send:submit&quot;&gt;
            &lt;input name=&quot;message&quot;&gt;
            &lt;input type=&quot;submit&quot;&gt;
        &lt;/form&gt;
     &lt;/div&gt;
     &lt;div id='messages'&gt;&lt;/div&gt;
{% endblock %}
</code></pre>
<p><img src="https://meetgor-cdn.pages.dev/django-htmx/chat-room-page.png" alt="Chat Room Page"></p>
<p>NOTE: The template has a div with the id <code>messages</code> which will be very important for sending the messages from the WebSocket to the client, so more on that when we use the HTMX part.</p>
<p>For testing this template, you can create a room, and that will redirect you to the room template as we have seen in the views for the room creation. If you see something like <code>WebSocket CONNECT</code> it means, that the application has been able to establish a WebSocket connection to the backend, and we can be ready to accept messages and other stuff.</p>
<pre><code>HTTP GET /chat/room/def_teas/ 200 [0.03, 127.0.0.1:38660]
WebSocket HANDSHAKING /chat/def_teas/ [127.0.0.1:38666]
WebSocket CONNECT /chat/def_teas/ [127.0.0.1:38666]
</code></pre>
<p><img src="https://meetgor-cdn.pages.dev/django-htmx/chat-asgi-server.png" alt="Django ASGI server websocket connection"></p>
<p>Till this point, we should have a running and almost complete application, though we just have a minor part missing that will be the most important part.</p>
<h3>Sending HTML response from backend for htmx</h3>
<p>We will be sending a fragment of HTML from the backend when the user sends a message, to broadcast it to the group. Let's make some changes to the application, especially to the receive method in the <code>ChatConsumer</code> of the chat application.</p>
<pre><code class="language-python"># chat / consumers.py
    

    ...
    ...

    async def receive(self, text_data):
        text_data_json = json.loads(text_data)
        message = text_data_json[&quot;message&quot;]
        user = self.user
        username = user.username

        await self.channel_layer.group_send(
            self.room_group_name, 
            {
                &quot;type&quot;: &quot;chat_message&quot;,
                &quot;message&quot;: message,
                &quot;username&quot;: username,
            }
        )

    async def chat_message(self, event):
        message = event[&quot;message&quot;]
        username = event[&quot;username&quot;]

        # This is the crucial part of the application
        message_html = f&quot;&lt;div hx-swap-oob='beforeend:#messages'&gt;&lt;p&gt;&lt;b&gt;{username}&lt;/b&gt;: {message}&lt;/p&gt;&lt;/div&gt;&quot;
        await self.send(
            text_data=json.dumps(
                {
                    &quot;message&quot;: message_html,
                    &quot;username&quot;: username
                }
            )
        )
</code></pre>
<p><img src="https://meetgor-cdn.pages.dev/django-htmx/chat-room-msg-1.png" alt="Chat Room Message">
<img src="https://meetgor-cdn.pages.dev/django-htmx/chat-room-msg-2.png" alt="Chat Room Message 2 Users">
<img src="https://meetgor-cdn.pages.dev/django-htmx/chat-room-msg-3.png" alt="Chat Room Message"></p>
<p>In the above snippet, we are just changing the final message object to include some HTML just simple. The HTML however has home htmx attributes like <a href="https://htmx.org/attributes/hx-swap-oob/">hx-swap-oob</a> which will just update the specified DOM element to the content in the div. In this case, the DOM element is <code>#message</code> which is the id message present in the room template. We basically add the username and the message into the same id by appending it to the element. That's it, it would work and it would start showing the messages from the connected clients and broadcast them as well.</p>
<p>There are some things to keep in mind while using htmx in the long run especially when the htmx 2.0 is released, it will have <code>ws</code> as a separate extension. It will have a bit of a different syntax than above. I have tried the latest version but doesn't seem to work. I'll just leave a few snippets for your understanding of the problem.</p>
<pre><code class="language-html"># templates / chat / room.html


{% extends 'base.html' %}

{% block base %}
    &lt;h2&gt;{{ name }}&lt;/h2&gt;
    &lt;div hx-ext=&quot;ws&quot; ws-connect=&quot;/chat/{{ slug }}/&quot;&gt;
        &lt;form ws-send&gt;
            &lt;input name=&quot;message&quot;&gt;
        &lt;/form&gt;
    &lt;/div&gt;
    &lt;div id='messages'&gt;&lt;/div&gt;
{% endblock %}
</code></pre>
<p>I have added, the <code>hx-ext</code> as <code>ws</code> which is a htmx <a href="https://htmx.org/extensions/web-sockets/">extension for websocket</a>. This extension has websocket-specific attributes like <code>ws-connect</code> and <code>ws-send</code>. I have tried a combination like changing the htmx versions, adding submit value as the <code>ws-send</code> attribute, etc, but no results yet. I have opened a <a href="https://github.com/bigskysoftware/htmx/discussions/1231">discussion</a> on GitHub for this issue, you can express your solution or views there.</p>
<h3>Adding some utility features for the chat app</h3>
<p>We can save messages, add and remove the users from the room according to the connection, and other stuff that can make this a fully-fledged app. So, I have made a few changes to the chat consumers for saving the messages and also updating the room with the users in the room.</p>
<pre><code class="language-python"># chat / consumers.py


import json

from asgiref.sync import sync_to_async
from channels.generic.websocket import AsyncWebsocketConsumer

from chat.models import Room, Message


class ChatConsumer(AsyncWebsocketConsumer):
    async def connect(self):
        self.room_name = self.scope[&quot;url_route&quot;][&quot;kwargs&quot;][&quot;room_slug&quot;]
        self.room_group_name = &quot;chat_%s&quot; % self.room_name
        self.user = self.scope[&quot;user&quot;]

        await self.channel_layer.group_add(
            self.room_group_name, self.channel_name
        )

        # Add the user when the client connects
        await self.add_user(self.room_name, self.user)

        await self.accept()

    async def disconnect(self, close_code):

        # Remove the user when the client disconnects
        await self.remove_user(self.room_name, self.user)

        await self.channel_layer.group_discard(
            self.room_group_name, self.channel_name
        )

    async def receive(self, text_data):
        text_data_json = json.loads(text_data)
        message = text_data_json[&quot;message&quot;]
        user = self.user
        username = user.username
        room = self.room_name

        # Save the message on recieving
        await self.save_message(room, user, message)

        await self.channel_layer.group_send(
            self.room_group_name, 
            {
                &quot;type&quot;: &quot;chat_message&quot;,
                &quot;message&quot;: message,
                &quot;username&quot;: username,
            }
        )

    async def chat_message(self, event):
        message = event[&quot;message&quot;]
        username = event[&quot;username&quot;]


        message_html = f&quot;&lt;div hx-swap-oob='beforeend:#messages'&gt;&lt;p&gt;&lt;b&gt;{username}&lt;/b&gt;: {message}&lt;/p&gt;&lt;/div&gt;&quot;
        await self.send(
            text_data=json.dumps(
                {
                    &quot;message&quot;: message_html,
                    &quot;username&quot;: username
                }
            )
        )

    @sync_to_async
    def save_message(self, room, user, message):
        room = Room.objects.get(slug=room)
        Message.objects.create(room=room, user=user, message=message)

    @sync_to_async
    def add_user(self, room, user):
        room = Room.objects.get(slug=room)
        if user not in room.users.all():
            room.users.add(user)
            room.save()

    @sync_to_async
    def remove_user(self, room, user):
        room = Room.objects.get(slug=room)
        if user in room.users.all():
            room.users.remove(user)
            room.save()
</code></pre>
<p>So, we have created a few methods like <code>save_message</code>, <code>add_user</code>, and <code>remove_user</code> which all are <code>synchronous</code> methods but we are using an asynchronous web server, so we add in the <code>sync_to_async</code> decorator which wraps a synchronous method to an asynchronous method. Inside the methods, we simply perform the database operations like creating a message object, and adding or removing the user from the room.</p>
<p>That's only a few features that I have added, you can add to this application and customize them as per your needs.</p>
<p>The code for this chat app is provided in the <a href="https://github.com/Mr-Destructive/django-htmx-chat">GitHub repository</a>.</p>
<h2>Conclusion</h2>
<p>So, from this post, we were able to create a simple chat app (frontendless) with Django and htmx. We used Django channels and HTMX to make a chat application without the need to write javascript for the client-side connection. Hope you found this tutorial helpful, do give your feedback and thoughts on it, I'll be eager to improve this post. Thank you for your patient listening. Happy Coding :)</p>
