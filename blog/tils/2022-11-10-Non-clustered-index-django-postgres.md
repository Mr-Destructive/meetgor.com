---
templateKey: til
title: "Create a Non-Clustered Index in Django with Postgres as DB"
description: "Understanding how to add a non-clustered index in a postgres database in a django project."
status: published
slug: django-non-clustered-index-pg
tags: ["django", "python", "sql", "postgres"]
date: 2022-11-10 22:10:00
---

## What is a non-clustered index?

A non-clustered index is a seperate structure than an actual table in the database, it stores the non-clustered index key(the column which we want to sort in the table), and a pointer to the actual values based on the index key. So, non-clustered indexes do not change the physical order of the table records, instead it holds a structure that can provide a easier and distinct way to fetch objects based on a particular column as the primary key in the structure.

## How to create a non-clustered index in django

In django, we can use the [db_index](https://docs.djangoproject.com/en/4.1/ref/models/indexes/) property on a field(s) to create a index on the table/model. 

### Add the property to the field in the model

Chose a field in which, you want to add a index. It can be a foreign key or any other normal field defined in your model.

We have used the typical blog model, so used in the some of my [TILS](https://www.meetgor.com/tils/) in django, it is just convenient to explain and understand as well. We have a django project named `core` and it has a app `blog` with a model defined below. The model `Article` has a few attributes like `title`, `description`, `content` and `status`.

```python
from django.db import models

ARTICLE_STATUS = [
    ("PUBLISHED", "Published"),
    ("DRAFT", "Draft"),
]

class Article(models.Model):
    title = models.CharField(max_length=128, db_index=True)
    description = models.CharField(max_length=512)
    content = models.TextField()
    status = models.CharField(max_length=16, choices=ARTICLE_STATUS, default="DRAFT")

    def __str__(self):
        return self.title
```

So, we have added a `db_index` to the title column in the model as a property. This will be equivalent to creating a index in `SQL` as follows:

```
$ python manage.py makemigrations

Migrations for 'blog':
  blog/migrations/0002_alter_article_title.py
    - Alter field title on article
```

```
$ python manage.py migrate

Operations to perform:
  Apply all migrations: admin, auth, blog, contenttypes, sessions
Running migrations:
  Applying blog.0002_alter_article_title... OK

```

Indexes are not standard as in SQL, but each vendor(sqlite, postgres, mysql) have their own flavour of syntax and naunces.

```sql
CREATE INDEX "blog_article_title_3c514952" ON "blog_article" ("title");

CREATE INDEX "blog_article_title_3c514952_like" ON "blog_article" ("title" varchar_pattern_ops);
```

The above index commands are specific to the field, as the title field is a varchar, it has two types of index, it can generate one with simple match and other for `LIKE` comparisons because of string comparison behaviour.

So, we just created a simple index and now if we query the db for a particular `title` which now has its own index for the table `blog_article`. This means, we will be able to fetch queries quickly if we are specifically filtering for `title`.

### Adding some data records

We can add a few data records to test the query from the databse, you can ignore this part as it would be just setting up a django project and adding a few records to the databse. This part won't make sense for people reading to get the actual stuff done, move to the next part please.

```
python manage.py createsuperuser
# Create a super user and run the server

python manage.py runserver
# Locate to http://127.0.0.1:8000/admin
# Create some records in the artilce model
```

So, after creating some records, you should have a simple database and a working django application.

```sql
SELECT * FROM blog_article;
```
```
blog_test=# SELECT * FROM blog_article;

 id |  title   | description |          content          |  status   
----+----------+-------------+---------------------------+-----------
  1 | test     | test 1      | test content              | DRAFT
  2 | testpost | test 2      | test content more content | DRAFT
  3 | newpost  | test 3      | test nothing              | PUBLISHED
(3 rows)
```

## Testing Queries

We can now use SQL queries or django filters to check if we get results by a sequential or an index scan. If we have a filter of `title` we will get the results after performing an `Index Scan` which means, it will look up in the index columns rather than scanning the entire table of records. This is a way **we can test the indexes are working, efficiency is a differnet topic.** We can't get a idea of performance with this little data and just one connection. A real time database and having multiple conncurrent requests and connections is a good environment to test(don't do it in a production db :)


```sql
EXPLAIN SELECT * FROM blog_article WHERE description LIKE 'test 2';
```

```
blog_test=# EXPLAIN ANALYSE SELECT * FROM blog_article WHERE description LIKE 'test';
---------------------------------------------------------------------------------------------------------

 Seq Scan on blog_article  (cost=0.00..11.00 rows=1 width=880) (actual time=0.180..0.181 rows=0 loops=1)
   Filter: ((description)::text ~~ 'test'::text)
   Rows Removed by Filter: 3
 Planning Time: 0.189 ms
 Execution Time: 0.217 ms
(5 rows)

```

The above query selects the records whose `description` is like `test 2`, this performs a `Sequenitial Scan` in the database i.e. iterating over the records one by one of the order of the primary key / id of the records in the table. 

```sql
EXPLAIN SELECT * FROM blog_article WHERE title LIKE 'test 2';
```

```
blog_test=# EXPLAIN ANALYSE SELECT * FROM blog_article WHERE title LIKE 'test';
---------------------------------------------------------------------------------------------------------

Index Scan using blog_article_title_3c514952_like on blog_article  (cost=0.14..8.16 rows=1 width=880) (actual time=0.043..0.048 rows=1 loops=1)
   Index Cond: ((title)::text = 'test'::text)
   Filter: ((title)::text ~~ 'test'::text)
 Planning Time: 0.208 ms
 Execution Time: 0.093 ms
(5 rows)
```

In the above query, the select statement has a filter with the title being like `test 2`, and since we have a index for looking for like of title column, the database performs a index scan on that table and fetches the result.

Here are some tradeoffs, the planning is more and the execution time is less, this is quite logical as it would take time to make decision because the database has more options than before creating indexes.

In the query where we filtered the description, the planning time was less as it makes sense there was just one option to go for sequential scan, but it took time to perform the operation as it would scan the entire table one by one.

## Using Django to test queries

We can even use django to filter out the objects in the table. We simply use the `filter` method to check with a particular value.

We can use the shell, to perform some queries. You can use this in your views or viewsets as per your requirements and constraints.

We can even use `explain` to see what the underlying `sql` got executed out from the ORM. The [explain](https://docs.djangoproject.com/en/3.1/ref/models/querysets/#explain) function is similar to the `EXPLAIN ANALYSE` command in the `sql` queries. It gives a bit of context on how the query was executed.

```
$ python manage.py shell
```

```python
>>> from blog.models import Article                                                                
>>> Article.objects.filter(description='test 1')                                                   

<QuerySet [<Article: test>]>                                                                       


>>> Article.objects.filter(description='test 1').explain()                                         

"Seq Scan on blog_article  (cost=0.00..11.00 rows=1 width=880)\n  Filter: ((description)::text = 't
est 1'::text)"                                                                                     


>>> Article.objects.filter(title='test')                                                           

<QuerySet [<Article: test>]>                                                                       


>>> Article.objects.filter(title='test').explain()                                                 

"Index Scan using blog_article_title_3c514952_like on blog_article  (cost=0.14..8.16 rows=1 width=8
80)\n  Index Cond: ((title)::text = 'test'::text)"                                                 

```

We can use `__contains` for replicating the behaviour of `LIKE` in python/django from SQL. The below example will check if the title has a word `test` in any records of the database.

```
>>> Article.objects.filter(title__contains='test')

<QuerySet [<Article: test>, <Article: testpost>]> 
```

BONUS: We can even get the underlying SQL with the `.query.__str__()` method. 

```
articles = Article.objects.filter(title__contains='test')

articles.query.__str__()
```

```
'SELECT "blog_article"."id", "blog_article"."title", "blog_article"."description", "blog_article"."
content", "blog_article"."status" FROM "blog_article" WHERE "blog_article"."title"::text LIKE %test
%'
```

Here, we are able to see that clearly, that the django orm used the `LIKE` clause for comparing the title.

Further readings and references: 

- [Indexing in Postgres](https://medium.com/geekculture/indexing-in-postgres-db-4cf502ce1b4e)
- [Indexing refernece Django docs](https://docs.djangoproject.com/en/4.1/ref/models/indexes/)
- [Non-Clustered indexing](https://gudevsoc.com/what-is-non-clustered-index-in-sql-with-example/)
