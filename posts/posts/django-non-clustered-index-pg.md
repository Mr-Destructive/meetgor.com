{
  "title": "Create a Non-Clustered Index in Django with Postgres as DB",
  "status": "published",
  "slug": "django-non-clustered-index-pg",
  "date": "2025-04-05T12:33:34Z"
}

<h2>What is a non-clustered index?</h2>
<p>A non-clustered index is a seperate structure than an actual table in the database, it stores the non-clustered index key(the column which we want to sort in the table), and a pointer to the actual values based on the index key. So, non-clustered indexes do not change the physical order of the table records, instead it holds a structure that can provide a easier and distinct way to fetch objects based on a particular column as the primary key in the structure.</p>
<h2>How to create a non-clustered index in django</h2>
<p>In django, we can use the <a href="https://docs.djangoproject.com/en/4.1/ref/models/indexes/">db_index</a> property on a field(s) to create a index on the table/model.</p>
<h3>Add the property to the field in the model</h3>
<p>Chose a field in which, you want to add a index. It can be a foreign key or any other normal field defined in your model.</p>
<p>We have used the typical blog model, so used in the some of my <a href="https://www.meetgor.com/tils/">TILS</a> in django, it is just convenient to explain and understand as well. We have a django project named <code>core</code> and it has a app <code>blog</code> with a model defined below. The model <code>Article</code> has a few attributes like <code>title</code>, <code>description</code>, <code>content</code> and <code>status</code>.</p>
<pre><code class="language-python">from django.db import models

ARTICLE_STATUS = [
    (&quot;PUBLISHED&quot;, &quot;Published&quot;),
    (&quot;DRAFT&quot;, &quot;Draft&quot;),
]

class Article(models.Model):
    title = models.CharField(max_length=128, db_index=True)
    description = models.CharField(max_length=512)
    content = models.TextField()
    status = models.CharField(max_length=16, choices=ARTICLE_STATUS, default=&quot;DRAFT&quot;)

    def __str__(self):
        return self.title
</code></pre>
<p>So, we have added a <code>db_index</code> to the title column in the model as a property. This will be equivalent to creating a index in <code>SQL</code> as follows:</p>
<pre><code>$ python manage.py makemigrations

Migrations for 'blog':
  blog/migrations/0002_alter_article_title.py
    - Alter field title on article
</code></pre>
<pre><code>$ python manage.py migrate

Operations to perform:
  Apply all migrations: admin, auth, blog, contenttypes, sessions
Running migrations:
  Applying blog.0002_alter_article_title... OK

</code></pre>
<p>Indexes are not standard as in SQL, but each vendor(sqlite, postgres, mysql) have their own flavour of syntax and naunces.</p>
<pre><code class="language-sql">CREATE INDEX &quot;blog_article_title_3c514952&quot; ON &quot;blog_article&quot; (&quot;title&quot;);

CREATE INDEX &quot;blog_article_title_3c514952_like&quot; ON &quot;blog_article&quot; (&quot;title&quot; varchar_pattern_ops);
</code></pre>
<p>The above index commands are specific to the field, as the title field is a varchar, it has two types of index, it can generate one with simple match and other for <code>LIKE</code> comparisons because of string comparison behaviour.</p>
<p>So, we just created a simple index and now if we query the db for a particular <code>title</code> which now has its own index for the table <code>blog_article</code>. This means, we will be able to fetch queries quickly if we are specifically filtering for <code>title</code>.</p>
<h3>Adding some data records</h3>
<p>We can add a few data records to test the query from the databse, you can ignore this part as it would be just setting up a django project and adding a few records to the databse. This part won't make sense for people reading to get the actual stuff done, move to the next part please.</p>
<pre><code>python manage.py createsuperuser
# Create a super user and run the server

python manage.py runserver
# Locate to http://127.0.0.1:8000/admin
# Create some records in the artilce model
</code></pre>
<p>So, after creating some records, you should have a simple database and a working django application.</p>
<pre><code class="language-sql">SELECT * FROM blog_article;
</code></pre>
<pre><code>blog_test=# SELECT * FROM blog_article;

 id |  title   | description |          content          |  status   
----+----------+-------------+---------------------------+-----------
  1 | test     | test 1      | test content              | DRAFT
  2 | testpost | test 2      | test content more content | DRAFT
  3 | newpost  | test 3      | test nothing              | PUBLISHED
(3 rows)
</code></pre>
<h2>Testing Queries</h2>
<p>We can now use SQL queries or django filters to check if we get results by a sequential or an index scan. If we have a filter of <code>title</code> we will get the results after performing an <code>Index Scan</code> which means, it will look up in the index columns rather than scanning the entire table of records. This is a way <strong>we can test the indexes are working, efficiency is a differnet topic.</strong> We can't get a idea of performance with this little data and just one connection. A real time database and having multiple conncurrent requests and connections is a good environment to test(don't do it in a production db :)</p>
<pre><code class="language-sql">EXPLAIN SELECT * FROM blog_article WHERE description LIKE 'test 2';
</code></pre>
<pre><code>blog_test=# EXPLAIN ANALYSE SELECT * FROM blog_article WHERE description LIKE 'test';
---------------------------------------------------------------------------------------------------------

 Seq Scan on blog_article  (cost=0.00..11.00 rows=1 width=880) (actual time=0.180..0.181 rows=0 loops=1)
   Filter: ((description)::text ~~ 'test'::text)
   Rows Removed by Filter: 3
 Planning Time: 0.189 ms
 Execution Time: 0.217 ms
(5 rows)

</code></pre>
<p>The above query selects the records whose <code>description</code> is like <code>test 2</code>, this performs a <code>Sequenitial Scan</code> in the database i.e. iterating over the records one by one of the order of the primary key / id of the records in the table.</p>
<pre><code class="language-sql">EXPLAIN SELECT * FROM blog_article WHERE title LIKE 'test 2';
</code></pre>
<pre><code>blog_test=# EXPLAIN ANALYSE SELECT * FROM blog_article WHERE title LIKE 'test';
---------------------------------------------------------------------------------------------------------

Index Scan using blog_article_title_3c514952_like on blog_article  (cost=0.14..8.16 rows=1 width=880) (actual time=0.043..0.048 rows=1 loops=1)
   Index Cond: ((title)::text = 'test'::text)
   Filter: ((title)::text ~~ 'test'::text)
 Planning Time: 0.208 ms
 Execution Time: 0.093 ms
(5 rows)
</code></pre>
<p>In the above query, the select statement has a filter with the title being like <code>test 2</code>, and since we have a index for looking for like of title column, the database performs a index scan on that table and fetches the result.</p>
<p>Here are some tradeoffs, the planning is more and the execution time is less, this is quite logical as it would take time to make decision because the database has more options than before creating indexes.</p>
<p>In the query where we filtered the description, the planning time was less as it makes sense there was just one option to go for sequential scan, but it took time to perform the operation as it would scan the entire table one by one.</p>
<h2>Using Django to test queries</h2>
<p>We can even use django to filter out the objects in the table. We simply use the <code>filter</code> method to check with a particular value.</p>
<p>We can use the shell, to perform some queries. You can use this in your views or viewsets as per your requirements and constraints.</p>
<p>We can even use <code>explain</code> to see what the underlying <code>sql</code> got executed out from the ORM. The <a href="https://docs.djangoproject.com/en/3.1/ref/models/querysets/#explain">explain</a> function is similar to the <code>EXPLAIN ANALYSE</code> command in the <code>sql</code> queries. It gives a bit of context on how the query was executed.</p>
<pre><code>$ python manage.py shell
</code></pre>
<pre><code class="language-python">&gt;&gt;&gt; from blog.models import Article                                                                
&gt;&gt;&gt; Article.objects.filter(description='test 1')                                                   

&lt;QuerySet [&lt;Article: test&gt;]&gt;                                                                       


&gt;&gt;&gt; Article.objects.filter(description='test 1').explain()                                         

&quot;Seq Scan on blog_article  (cost=0.00..11.00 rows=1 width=880)\n  Filter: ((description)::text = 't
est 1'::text)&quot;                                                                                     


&gt;&gt;&gt; Article.objects.filter(title='test')                                                           

&lt;QuerySet [&lt;Article: test&gt;]&gt;                                                                       


&gt;&gt;&gt; Article.objects.filter(title='test').explain()                                                 

&quot;Index Scan using blog_article_title_3c514952_like on blog_article  (cost=0.14..8.16 rows=1 width=8
80)\n  Index Cond: ((title)::text = 'test'::text)&quot;                                                 

</code></pre>
<p>We can use <code>__contains</code> for replicating the behaviour of <code>LIKE</code> in python/django from SQL. The below example will check if the title has a word <code>test</code> in any records of the database.</p>
<pre><code>&gt;&gt;&gt; Article.objects.filter(title__contains='test')

&lt;QuerySet [&lt;Article: test&gt;, &lt;Article: testpost&gt;]&gt; 
</code></pre>
<p>BONUS: We can even get the underlying SQL with the <code>.query.__str__()</code> method.</p>
<pre><code>articles = Article.objects.filter(title__contains='test')

articles.query.__str__()
</code></pre>
<pre><code>'SELECT &quot;blog_article&quot;.&quot;id&quot;, &quot;blog_article&quot;.&quot;title&quot;, &quot;blog_article&quot;.&quot;description&quot;, &quot;blog_article&quot;.&quot;
content&quot;, &quot;blog_article&quot;.&quot;status&quot; FROM &quot;blog_article&quot; WHERE &quot;blog_article&quot;.&quot;title&quot;::text LIKE %test
%'
</code></pre>
<p>Here, we are able to see that clearly, that the django orm used the <code>LIKE</code> clause for comparing the title.</p>
<p>Further readings and references:</p>
<ul>
<li><a href="https://medium.com/geekculture/indexing-in-postgres-db-4cf502ce1b4e">Indexing in Postgres</a></li>
<li><a href="https://docs.djangoproject.com/en/4.1/ref/models/indexes/">Indexing refernece Django docs</a></li>
<li><a href="https://gudevsoc.com/what-is-non-clustered-index-in-sql-with-example/">Non-Clustered indexing</a></li>
</ul>
