{
  "title": "Django Bulk Update QuerySet objects",
  "status": "published",
  "slug": "django-bulk-update-queryset",
  "date": "2025-04-05T12:33:35Z"
}

<p>Let's say, I have a lots of objects which I want to update with a particular field or fields. We can use the <a href="https://docs.djangoproject.com/en/4.1/ref/models/querysets/#bulk-update">bulk_update</a> method with the model name.</p>
<pre><code class="language-python"># blog/models.py

from django.db import models

ARTICLE_STATUS = [
    (&quot;PUBLISHED&quot;, &quot;Published&quot;),
    (&quot;DRAFT&quot;, &quot;Draft&quot;),
]

class Article(models.Model):
    title = models.CharField(max_length=128)
    description = models.CharField(max_length=512)
    content = models.TextField()
    status = models.CharField(max_length=16, choices=ARTICLE_STATUS, default=&quot;DRAFT&quot;)

    def __str__(self):
        return self.title

</code></pre>
<p>Let's say we have a simple model <code>Article</code> with a few typical attributes like <code>title</code>, <code>description</code>, <code>content</code>, and <code>status</code>. We have the status as a choice field from two options as <code>Draft</code> and <code>Published</code>. It could be a boolean field, but that looks too gross for a article status.</p>
<pre><code class="language-python">
from blog.models import Article

articles = Article.objects.filter(status=&quot;draft&quot;)

for i in range(len(articles)):
    articles[i].status = &quot;published&quot;

Article.objects.bulk_update(articles, [&quot;status&quot;,])

</code></pre>
<p>In the above code, the <code>Articles</code> model is filtered by the status of <code>draft</code>. We iterate over the <code>QuerySet</code> which will contain the objects of the articles, by setting the object's properties to the value we want to set. We are jsut setting the value of the property of the object for each object.</p>
<p>This just makes a changes to the <code>QuerySet</code>, by using the <code>bulk_update</code> method, the two parameters required are the <code>QuerySet</code> and the list of <code>fields</code> which are to be updated. The function returns the number of objects/records updated.</p>
<pre><code class="language-python">&gt;&gt;&gt; from blog.models import Article
&gt;&gt;&gt; articles = Article.objects.filter(status=&quot;DRAFT&quot;)
&gt;&gt;&gt; articles
&lt;QuerySet [&lt;Article: test 1&gt;, &lt;Article: test 3&gt;]&gt;

&gt;&gt;&gt; for i in range(len(articles)):
...     articles[i].status = &quot;PUBLISHED&quot;
...
&gt;&gt;&gt; articles
&lt;QuerySet [&lt;Article: test 1&gt;, &lt;Article: test 3&gt;]&gt;
&gt;&gt;&gt;

&gt;&gt;&gt; Article.objects.bulk_update(articles, ['status',])
2

&gt;&gt;&gt; Article.obejcts.get(title=&quot;test 1&quot;).status
'PUBLISHED'

&gt;&gt;&gt; Article.objects.filter(status=&quot;DRAFT&quot;)
&lt;QuerySet []&gt;
&gt;&gt;&gt;
</code></pre>
<p>As, we can see here there were two obejcts <code>test 1</code> and <code>test 2</code> objects with the status as <code>Draft</code>. By iterating over the queryset and assigning the status of the object to published, the query set was changed and modified locally.
By using the <code>bulk_update</code> method, we parsed the queryset and the list of attributes to be updated into the function. This gives us the number of objects which were updated, in this case <code>2</code>. We then look into the article actual record in the database and it has indeed updated to the value we set in this operation.</p>
