---
type: til
title: "Django Bulk Update QuerySet objects"
description: "Using bulk_update to update multiple objects in one go."
status: published
slug: django-bulk-update-queryset
tags: ["django", "python"]
date: 2022-10-31 18:00:00
---


Let's say, I have a lots of objects which I want to update with a particular field or fields. We can use the [bulk_update](https://docs.djangoproject.com/en/4.1/ref/models/querysets/#bulk-update) method with the model name.

```python
# blog/models.py

from django.db import models

ARTICLE_STATUS = [
    ("PUBLISHED", "Published"),
    ("DRAFT", "Draft"),
]

class Article(models.Model):
    title = models.CharField(max_length=128)
    description = models.CharField(max_length=512)
    content = models.TextField()
    status = models.CharField(max_length=16, choices=ARTICLE_STATUS, default="DRAFT")

    def __str__(self):
        return self.title

```

Let's say we have a simple model `Article` with a few typical attributes like `title`, `description`, `content`, and `status`. We have the status as a choice field from two options as `Draft` and `Published`. It could be a boolean field, but that looks too gross for a article status.


```python

from blog.models import Article

articles = Article.objects.filter(status="draft")

for i in range(len(articles)):
    articles[i].status = "published"

Article.objects.bulk_update(articles, ["status",])

```


In the above code, the `Articles` model is filtered by the status of `draft`. We iterate over the `QuerySet` which will contain the objects of the articles, by setting the object's properties to the value we want to set. We are jsut setting the value of the property of the object for each object.

This just makes a changes to the `QuerySet`, by using the `bulk_update` method, the two parameters required are the `QuerySet` and the list of `fields` which are to be updated. The function returns the number of objects/records updated.

```python
>>> from blog.models import Article
>>> articles = Article.objects.filter(status="DRAFT")
>>> articles
<QuerySet [<Article: test 1>, <Article: test 3>]>

>>> for i in range(len(articles)):
...     articles[i].status = "PUBLISHED"
...
>>> articles
<QuerySet [<Article: test 1>, <Article: test 3>]>
>>>

>>> Article.objects.bulk_update(articles, ['status',])
2

>>> Article.obejcts.get(title="test 1").status
'PUBLISHED'

>>> Article.objects.filter(status="DRAFT")
<QuerySet []>
>>>
```

As, we can see here there were two obejcts `test 1` and `test 2` objects with the status as `Draft`. By iterating over the queryset and assigning the status of the object to published, the query set was changed and modified locally.
By using the `bulk_update` method, we parsed the queryset and the list of attributes to be updated into the function. This gives us the number of objects which were updated, in this case `2`. We then look into the article actual record in the database and it has indeed updated to the value we set in this operation.

