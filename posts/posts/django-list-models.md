{
  "title": "Django: Get list of all models and associated fields in a django project",
  "status": "published",
  "slug": "django-list-models",
  "date": "2025-04-05T12:33:37Z"
}

<h2>Context</h2>
<p>Let's say we want the list of all the models and associated attributes in all the applications of a django project, we can do that using the <a href="https://docs.djangoproject.com/en/4.0/ref/applications/">django.apps</a> with apps method.</p>
<h2>Get all the models in a project</h2>
<p>To fetch all the models, we can use the <a href="https://docs.djangoproject.com/en/4.0/ref/applications/#django.apps.AppConfig.get_models">get_models</a> methods, it will return a list of model classes in all the entire project(all applications). We can import all the models in the django project with the command:</p>
<pre><code class="language-python">from django.apps import apps
models = apps.get_models()
</code></pre>
<pre><code>[&lt;class 'django.contrib.admin.models.LogEntry'&gt;, &lt;class 'django.contrib.auth.models.Permission'&gt;, 
&lt;class 'django.contrib.auth.models.Group'&gt;, &lt;class 'django.contrib.contenttypes.models.ContentType'&gt;,
 &lt;class 'django.contrib.sessions.models.Session'&gt;, &lt;class 'allauth.account.models.EmailAddress'&gt;, 
&lt;class 'allauth.account.models.EmailConfirmation'&gt;, &lt;class 'allauth.socialaccount.models.SocialApp'&gt;, 
&lt;class 'allauth.socialaccount.models.SocialAccount'&gt;, &lt;class 'allauth.socialaccount.models.SocialToken'&gt;, 
&lt;class 'user.models.TimeStampedModel'&gt;, &lt;class 'user.models.User'&gt;, &lt;class 'articles.models.Tags'&gt;,
 &lt;class 'articles.models.Series'&gt;, &lt;class 'articles.models.Article'&gt;, &lt;class 'blog.models.Blog'&gt;]
</code></pre>
<p>We are importing the apps and creating a list of the models in our django project. The Django app command will load all the applications in the project, and the <a href="">get_models</a> method will fetch the associated models. This has resulted in a list of model class objects, we can iterate over them and fetch the required details, we want.</p>
<p>For instance, If I am interested in the name of these models, I can use the <code>__name__</code> property to fetch the model's name.</p>
<pre><code class="language-python">from django.apps import apps
model_list = apps.get_models()

for model in model_list:
    print(model.__name__)
</code></pre>
<pre><code>LogEntry
Permission
Group
ContentType
Session
EmailAddress
EmailConfirmation
SocialApp
SocialAccount
SocialToken
TimeStampedModel
User
Tags
Series
Article
Blog

</code></pre>
<p>So, from the above example, we can see we have accessed all the model names in our entire django project.</p>
<h2>Access Application name associated with a model</h2>
<p>For accessing the name of the application from the model class, we can use the <code>_meta</code> attribute followed by the <code>app_label</code> property to get the <code>app_name</code> associated with the model.</p>
<pre><code class="language-python">from django.apps import apps
model_list = apps.get_models()

for model in model_list:
    print(f&quot;{model._meta.app_label}  -&gt; {model.__name__}&quot;)
</code></pre>
<pre><code>admin  -&gt; LogEntry
auth  -&gt; Permission
auth  -&gt; Group
contenttypes  -&gt; ContentType
sessions  -&gt; Session
account  -&gt; EmailAddress
account  -&gt; EmailConfirmation
socialaccount  -&gt; SocialApp
socialaccount  -&gt; SocialAccount
socialaccount  -&gt; SocialToken
user  -&gt; TimeStampedModel
user  -&gt; User
articles  -&gt; Tags
articles  -&gt; Series
articles  -&gt; Article
blog  -&gt; Blog
</code></pre>
<p>In the above example, we can see we have printed all the models with their associated application names.</p>
<h2>Accessing all the attributes associated with a model</h2>
<p>To access all the fields/property/attributes associated with a model, we can again use the <code>_meta</code> attribute followed by the <code>get_fields</code> method.  This method will return a list of field objects. For accessing the name of those attributes/fields, we have to iterate over the list and then further use <code>name</code> property.</p>
<pre><code class="language-python">from django.apps import apps
model_list = apps.get_models()
for model in model_list:
    print(model.__name__)
    field_list = model._meta.get_fields()
    for field in field_list:
        print(field.name)
</code></pre>
<pre><code>LogEntry
id
action_time
user
content_type
object_id
object_repr
action_flag
change_message
Permission
group
user
id
...
...
Blog
article
id
name
description
authors
</code></pre>
<p>So, that is how we get all the associated field names in the associated models in our django projects. Also, there are a lot of attributes, we can access with the apps property. The <code>__dict__.keys()</code> can be used to get the list of all associated properties or other methods in a class instance.</p>
<pre><code>&gt;&gt;&gt; m[14]._meta.get_fields()[4].__dict__.keys()

dict_keys(['name', 'verbose_name', '_verbose_name', 'primary_key', 'max_length', '_unique', 'blank', 'null', 'remote_field',
 'is_relation', 'default', 'editable', 'serialize', 'unique_for_date', 'unique_for_month', 'unique_for_year', 'choices', 
'help_text', 'db_index', 'db_column', '_db_tablespace', 'auto_created', 'creation_counter', '_validators', '_error_messages', 
'error_messages', 'db_collation', 'validators', 'attname', 'column', 'concrete', 'model'])
</code></pre>
<p>In the above example, I am using a list of models and getting the list of all the attributes associated with a field of a model. This can be applied and other properties can be accessed.</p>
<h2>Get Models with a specific app</h2>
<p>Let's say we want all the models associated with a particular application in the project, we can do that by specifying the name of the application.</p>
<pre><code class="language-python">from django.apps import apps

app_info = apps.get_app_config('articles')

print(app_info.__dict__.keys()

print(app_info.verbose_name)

print(app_info.models)

print(app_info.models['article'].__dict__.keys())
</code></pre>
<pre><code>dict_keys(['name', 'module', 'apps', 'label', 'verbose_name', 'path', 'models_module', 'models'])

'Article'

{'tags': &lt;class 'articles.models.Tags'&gt;, 'series': &lt;class 'articles.models.Series'&gt;, 'article': &lt;class 'articles.models.Article'&gt;}

dict_keys(['__module__', 'Article_Status', '__str__', 'get_absolute_url', '__doc__', '_meta', 'DoesNotExist', 'MultipleObjectsReturned', 'title', 'description', 'content', 'status', 'get_status_display', 'blog_id', 'blog', 'author_id', 'author', 'timestampedmodel_ptr_id', 'timestampedmodel_ptr'])
</code></pre>
<p>So, we can see that we have got the information about the app <code>articles</code> in the proejct where we can get the <code>verbose_name</code> property to fetch the human-readable format of the article model. Further, we can get all the models associated with the <code>articles</code> application. We get back a dict with the model name as the key and the class reference as the value.</p>
<p>We have accessed the <code>article</code> model in the <code>articles</code> application and fetched all the associated properties or methods in the model.</p>
<p>For further references, you can visit the <a href="https://docs.djangoproject.com/en/4.0/ref/applications/">django apps documentation</a> to get more relevant methods and properties.</p>
