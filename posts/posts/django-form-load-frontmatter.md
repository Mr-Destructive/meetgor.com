{
  "title": "Django Blog DevLog: Load Frontmatter data into Template/Model Form Fields",
  "status": "published",
  "slug": "django-form-load-frontmatter",
  "date": "2025-04-05T12:33:37Z"
}

<h2>Introduction</h2>
<p>I have an Article Form where I load my post into it directly, it might have frontmatter. So what I wish to achieve is when I paste in the content, the frontmatter should be picked up and it should render the form fields like <code>title</code>, <code>description</code>, and then also remove the frontmatter from the content.</p>
<p>To do that, we will require a model to work with and a form based on that model. We will exclude a few fields from that model so as to process these attributes on the server side. I am working on my Blog project which is a simple Django application.  You can get the source code for the project on the <a href="https://github.com/mr-destructive/techstructive-blog/">GitHub repository</a>.</p>
<h2>Django Project Context</h2>
<p>The techstructive-blog is a django project, which has a couple of applications currently, not in a good situation. There are apps like <code>article</code>, <code>blog</code>, and <code>user</code>. This project has templates and static folder in the base directory. The project is deployed on <a href="https://www.railway.app">railway</a> this is an always under development project, you can check out the <a href="https://django-blog.up.railway.app">techstructive-blog</a>. We can get the bits and pieces of the project details required for understanding what I want to do with the following sections.</p>
<h3>Article Model</h3>
<p>We have an <code>Article</code> model with attributes like <code>title</code>, <code>description</code>,  <code>author</code> as a Foreign Key to the user model, and a few other attributes which is not related to what we are trying to achieve right now or at least don't require an explanation. We have a model method called <code>get_absolute_url</code> for getting a URL in order to redirect the client when the model instance is created or updated from the backend. You can definitely check out the details of these small components or templates in the project repository.</p>
<pre><code class="language-python"># articles/models.py


class Article(TimeStampedModel):
    class Article_Status(models.TextChoices):
        DRAFT = (
            &quot;DRAFT&quot;,
            _(&quot;Draft&quot;),
        )
        PUBLISHED = (
            &quot;PUBLISHED&quot;,
            _(&quot;Published&quot;),
        )

    title = models.CharField(max_length=128)
    description = models.CharField(max_length=256)
    content = models.TextField(default=&quot;&quot;, null=False, blank=False)
    status = models.CharField(
        max_length=16,
        choices=Article_Status.choices,
        default=Article_Status.DRAFT,
    )
    blog = models.ForeignKey(Blog, on_delete=models.CASCADE, null=True, blank=True)
    author = models.ForeignKey(User, on_delete=models.CASCADE, related_name=&quot;author&quot;)

    def __str__(self):
        return self.title

    def get_absolute_url(self):      
        return reverse('articles:article-detail', args=[str(self.id)])
</code></pre>
<p>In the below snippet, we have the forms defined in the article application for creating or updating of article instance.  We will be using model forms as our form data should contain fields related to a model in this case the <code>Article</code> model. So when we inherit the <code>forms.ModelForm</code> in our custom <code>ArticleForm</code> we simply need to specify the model and it will take in all the attributes of that model by default, but if we specify the <code>fields</code> or <code>exclude</code> tuples, it will include only or exclude only the provided list of attributes from the model.</p>
<p>We have also added the widgets for the form fields which will allow us to customize the way the fields in the template/form will render. We can specify the HTML attributes like <code>width</code>, <code>height</code>, <code>style</code>, etc.</p>
<h3>Article Form</h3>
<pre><code class="language-python"># article/forms.py


from django import forms
from .models import Article


class ArticleForm(forms.ModelForm):
    class Meta:
        model = Article
        exclude = (
            &quot;created&quot;,
            &quot;updated&quot;,
            &quot;author&quot;,
        )
        widgets = {
            &quot;title&quot;: forms.TextInput(
                attrs={
                    &quot;class&quot;: &quot;form-control&quot;,
                    &quot;style&quot;: &quot;max-width: 450px; align: center;&quot;,
                    &quot;placeholder&quot;: &quot;Title&quot;,
                }
            ),
            &quot;description&quot;: forms.TextInput(
                attrs={
                    &quot;class&quot;: &quot;form-control&quot;,
                    &quot;style&quot;: &quot;max-width: 900px; &quot;,
                    &quot;placeholder&quot;: &quot;Description&quot;,
                }
            ),
            &quot;content&quot;: forms.Textarea(
                attrs={
                    &quot;class&quot;: &quot;form-control post-body&quot;,
                    &quot;id&quot;: &quot;text-content&quot;,
                    &quot;style&quot;: &quot;max-width:900px;&quot;,
                    &quot;hx-post&quot;: &quot;/article/meta/&quot;,
                    &quot;placeholder&quot;: &quot;Content&quot;,
                }
            ),
            &quot;blog&quot;: forms.Select(
                attrs={
                    &quot;class&quot;: &quot;form-control&quot;,
                    &quot;placeholder&quot;: &quot;Blog Publication&quot;,
                }
            ),
        }

</code></pre>
<p>So, these are my models and form files in the article app. Using htmx and without any javascript I want to update the form so that it picks up the front matter in the content field which is a text area and fills the title, description other attributes automatically for me.</p>
<p>This can be done in a lot of ways, but I will be sharing one of the ways that I recently used in my blog project. This process involves creating a class-based view and adding a <code>POST</code> method that won't post any data to the backend but will send necessary data to the view.</p>
<p>Let's see the process before diving into any of the code:</p>
<h2>Gist of the Process</h2>
<ul>
<li>Attach a <code>hx-post</code> attribute to the form field for sending the request to a view</li>
<li>When the request is sent, the data is loaded with <code>request.POST</code>, it is cleaned and converted in python-readable format with json.</li>
<li>Once we have the data, we try to use the <code>frontmatter.loads</code> function that will load the content and if we have a frontmatter in the text, it will load it as a <code>frontmatter.POST</code> object.</li>
<li>We will extract <code>title</code>, <code>description</code>, and other data fields from the object.</li>
<li>We will initialize a Form instance of Article, with the initial data values as the extracted data from the frontmatter.</li>
<li>Now we have two options:
<ul>
<li>If the article instance already exists i.e. we are updating the article</li>
<li>Else we are creating a new article</li>
</ul>
</li>
</ul>
<p>Accordingly, we will load the form in the respective templates i.e. <code>update.html</code> for updating the existing articles and <code>article-form.html</code> for a new article.</p>
<h2>Adding HTMX Magic</h2>
<p>If you'd have seen we have a <code>hx-post</code> attribute in the <code>article/forms.py</code> file, the <code>content</code> widget has a <code>hx-post</code> attribute which sends a post request to the <code>article/meta/</code> URL route. This route we will bind to the <code>ArticleMetaView</code> which we will define in a few moments. This is usually sent once we change a certain text in the content field, so we can modify it as per your requirement with <code>hx-trigger</code> that can enable us to specify the trigger event or the type of trigger we want. Further, you can read from the <a href="https://htmx.org/docs/#trigger-modifiers">htmx docs</a> about these triggers and other attributes.</p>
<pre><code class="language-python"># article/urls.py

from django.urls import path
from . import views

app_name = &quot;articles&quot;

urlpatterns = [
    path(&quot;&quot;, views.ArticleCreateView.as_view(), name=&quot;article-create&quot;),
    path(&quot;&lt;int:pk&gt;/&quot;, views.ArticleDetailView.as_view(), name=&quot;article-detail&quot;),
    path(&quot;delete/&lt;int:pk&gt;/&quot;, views.ArticleDetailView.as_view(), name=&quot;article-delete&quot;),
    path(&quot;edit/&lt;int:pk&gt;&quot;, views.ArticleDetailView.as_view(), name=&quot;article-update&quot;),

    # the new view that we will create
    path(&quot;meta/&quot;, views.ArticleMetaView.as_view(), name=&quot;article-meta&quot;),
]
</code></pre>
<h2>Capture Frontmatter Meta-data View</h2>
<p>Along with the Create, Detail/List, Update, Delete View, I will create a separate class called <code>ArticleMetaView</code> that will fetch the form fields and render the templates again but this time it will fill in the frontmatter meta-data in the fields if the content is parsed with the relvant frontmatter.</p>
<pre><code class="language-python"># articles/view.py

class ArticleMetaView(View):
    model = Article

    def post(self, request, *args, **kwargs):
        
        data = json.loads(json.dumps(dict(request.POST)))
        loaded_frontmatter = frontmatter.loads(data['content'][0])

       # frontmatter has keys i.e. attributes like title, description, etc.
        if dict(loaded_frontmatter):
            article_title = loaded_frontmatter['title']
            article_description = loaded_frontmatter['description']
            form = ArticleForm(initial={'title': article_title, 
            'description': article_description, 'content': loaded_frontmatter.content})
            context = {'form': form}
            article_list = Article.objects.filter(title=article_title)
            if article_list:
                article = article_list.last()
                context['article'] = article
                return render(request, 'articles/edit_article.html', context)
            return render(request, 'articles/article_form.html', context)

        article_list = Article.objects.filter(title=data['title'][0])
       
       # if the article title has been already taken i.e. we are updating an article

        if article_list:
            article = article_list.last()
            form = ArticleForm(data=request.POST)
            context = {'form': form}
            context['article'] = article
            return render(request, 'articles/edit_article.html', context)

        form = ArticleForm(data=request.POST)
        context = {'form': form}
        return render(request, 'articles/article_form.html', context)

</code></pre>
<p>In the above <code>ArticleMetaView</code> we have created a <code>post</code> method as we want to get hold of the content from the form. So, we start by extracting and converting the <code>request.body</code> data into an appropriate type for easily working with python. So, the <code>request.body</code> will contain the data like <code>csrf_token</code>, <code>form_data</code>, etc. received from the frontend template. We store the received data as <code>data</code> and now from this data, we can load the content field which will have the content information.</p>
<p>Firstly we will extract the <code>request.body</code> which will contain the data from the form as we have made a <code>POST</code> request to this endpoint. For doing that we need to parse the content in a apropriate format such that it is python friendly. So we wrap the <code>request.body</code> into json format and then decode it back into the json string. This will give us the dict of the request data.</p>
<pre><code class="language-python">data = json.loads(json.dumps(dict(request.POST)))
</code></pre>
<pre><code>{'csrfmiddlewaretoken': ['bSYJxD39XH509tD1tZGd0WU21PUaKaLeqjjGbyzRvLXF4P8iIxb5l0fmTWVFjELQ'], 'title': ['test2'], 'description': ['test'], 'content': ['test something'], 'status': ['DRAFT'], 'blog': ['']}
</code></pre>
<p>So, this will grab the request data as a dict, we can then extract the data from this as it has data from the Form fields. We are interested in the content field in the Form, so we can get it by specifying the key <code>content</code> from the extracted data. But as we can see the data doesn't contain the actual data instead it is wrapped in a list i.e. <code>['test something']</code>, so we will have to index it and then fetch it.</p>
<pre><code class="language-python">content_string = data['content'][0]
</code></pre>
<p>This will give us the exact content field as a string. So, we can now move into extracting the frontmatter from the fields.</p>
<p>Now, we can use the <a href="https://python-frontmatter.readthedocs.io/en/latest/index.html">frontmatter</a> library to parse the content into the <a href="https://python-frontmatter.readthedocs.io/en/latest/api.html#frontmatter.loads">loads</a> funciton and extract the frontmatter if it is present in the content field. The frontmatter library has a <code>loads</code> function which takes in a string and can give out a <a href="https://python-frontmatter.readthedocs.io/en/latest/api.html#post-objects">frontmatter.Post</a> object. The loads function is differnet from the <a href="https://python-frontmatter.readthedocs.io/en/latest/api.html#frontmatter.load">load</a> function as the load frunciton is for reading data from a stream of bytes i.e. a file or othe related byte object. The differnece is subtle but it took a read at the <a href="https://python-frontmatter.readthedocs.io/en/latest/api.html#module-frontmatter">documentation</a>.</p>
<pre><code class="language-python">post = data['content'][0]
loaded_frontmatter = frontmatter.loads(post)
</code></pre>
<p>This wil load the content and give us a <code>frontmatter.Post</code> as said earlier. This will contain a dict with all the frontmatter if it has any and will by default parse the non-frontmatter data i.e. the remaining text into the <code>content</code> key. We need a chack if the Form field had any fronmatter this can be checked by the <code>dict(loaded_frontmatter)</code> which will return None if it cannot load the frontmatter.</p>
<pre><code class="language-python">loaded_frontmatter = frontmatter.loads(data['content'][0])
if dict(loaded_frontmatter):
  print(loaded_frontmatter.keys())
</code></pre>
<pre><code>dict_keys(['templateKey', 'title', 'description', 'date', 'status'])
</code></pre>
<p>So once we have the frontmatter loaded we can get specific keys from it and initialize the form vaues to them. But we have made clear distictions that we want to perform a specific task if we have frontmatter keys in the content field of the Form else we can do something else.</p>
<p>First let's handle the loading of the frontmatter into the form. For doing that we will get all the required attributes from the frontmatter like <code>title</code>, <code>description</code>, <code>content</code>, etc which can be accessed normally as we extract the value from a key in a dict.</p>
<p>Once we have got those keys, we can start filling in the Form data with initial values. The <a href="https://docs.djangoproject.com/en/4.0/topics/forms/modelforms/">Django Model form</a> takes in a parameter like <a href="https://docs.djangoproject.com/en/4.0/topics/forms/modelforms/#providing-initial-values">initial</a> which can be a dict of the fiields along with the value that can be used for rendering the form initially when we load the template.</p>
<pre><code class="language-python">article_title = loaded_frontmatter['title']
article_description = loaded_frontmatter['description']

form = ArticleForm(initial={'title': article_title, 'description': article_description, 'content': loaded_frontmatter.content})
</code></pre>
<p>This will take in a <code>ArticleForm</code> and fill the initial values like <code>title</code>, <code>description</code>, etc which we have provided in the dict with the values. Now, we need to parse this form in the current template or re-render the template. But before that, we need to parse this context into the template. We will create a dict with <code>form</code> as the key which can be used to render in the template.</p>
<p>Also, we have a two ways here, either the user is creating a new article or it is updating a existing article. We need to make sure that we preserve the initial fields in the form as we are updating the existing article. So, we can filter the article objects as per the title of the current title and then if we find a article with that title, we will parse the context with that article object.</p>
<pre><code class="language-python">article_list = Article.objects.filter(title=article_title)
if article_list:
    article = article_list.last()
    context{
      'form': form,
      'article': article
    }
    return render(request, 'articles/edit_article.html', context)
context = {'form': form}
return render(request, 'articles/article_form.html', context)
</code></pre>
<p>Now, we have form data along with the article instance used for rendering the form with appropriate content. So, this will work for editing an already existing article. For a new article, we have to simply parse the form to the template and it will render the title picked from the fotnmatter or leave it empty.</p>
<p>Similarly, for the article with no frontmatter we will iterate over the article and if the article's title already exist, we will render the article data with the form else render the form with the parsed title and other meta-data in the form.</p>
<!-- raw HTML omitted -->
<p>So that is how we render the form data with frontmatter into appropriate meta-data in the form. We have used Django forms and make use of HTMX for the dynamic updation of form.</p>
