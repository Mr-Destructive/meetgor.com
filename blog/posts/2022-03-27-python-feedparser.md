---
templateKey: blog-post
title: "Feedparser: Python package for reading RSS feeds"
description: "Read and extract content from RSS feeds in python using feedparser"
date: 2022-03-26 15:45:00
status: published
slug: python-feedparser
tags: ['python',]
image_url: https://meetgor-cdn.pages.dev/python-pkg-feedparser.png
---

## Introduction

[Feedparser](https://pypi.org/project/feedparser/) is a simple but powerful python package that can be used to extract information about a specific webpage or a publication with its RSS feed(not only RSS). By providing the RSS feed link, we can get structured information in the form of python lists and dictionaries. It can be basically used in a pythonic way to read RSS feeds, it is really simple to use and it even normalizes different types of feeds.

Today, we will be taking a look at the feedparser package in python and how to extract information from a given RSS feed.

## What is feedparser

Feedparser is a python package for parsing feeds of almost any type such as RSS, Atom, RDF, etc. It is a package that allows us to parse or extract information using python semantics. For example, all the latest posts from a given blog can be accessed on a list in python, further different attributes like links, images, titles, descriptions, can be accessed within a dictionary as key-value pairs. 

## Installing feedparser

As feedparser is a python package you can install it with pip very easily.

```
pip install feedparser
```

This will install feedparser in your respective python environment, it can be a virtual environment or a global environment. 


## Using feedparser

To test out feedparser, you can open up a python repl, in the environment where you installed the Feedparser package.

```
python
```

Firstly import the package.

```python
import feedparser
```

Now, we can use the module in our application to get all of the functions or methods from the package.

## Parse an RSS feed URL

To parse an RSS feed link, we can simply use the `parse` function from the feedparser package. The [parse](https://feedparser.readthedocs.io/en/latest/introduction.html) function takes in a string that can be a URL or a file path. Generally, the URL seems to be more useful. So, we can look up any RSS feed on the internet like your blog's feed, publications feeds, and so on. 

```python
feedparser.parse("url_of_the_rss_feed")
```

The parse function basically fetches the feed from the provided URL or the file. It extracts the feed in a systematic way storing each piece of information in a structured format. At the high level, it returns a dictionary with a few key-value pairs. Further, each key might have a list or nested dictionaries in it. The key identifiers are named in a uniform manner for any feed you parse in the function. Though there might be a few cases where there might be additional information to be parsed, it can even add more information ad shape the structure accordingly.

This will give you a dictionary in python, that can have more or less similar keys. The most common keys that can be used in extracting information are `entries` and `feed`. We can get all the keys associated with a feed that is parsed using the `keys` function.

```python
feedparser.parse("url_of_the_rss_feed").keys()
```

![Feedparser Keys](https://res.cloudinary.com/techstructive-blog/image/upload/v1648370871/blog-media/ph6bsxobyifqmusumirx.png)

The keys function basically gets all the keys in the dictionary in python.

```
>>> feedparser.parse("https://dev.to/feed/").keys()
dict_keys(['bozo', 'entries', 'feed', 'headers', 'etag', 'href', 'status', 'encoding', 'version', 'namespaces'])
```   

This will give out a list of all the keys in the feed which we have parsed from the RSS feed previously. From this list of keys, we can extract the required information from the feed.

Before we extract content from the feed, we can store the dictionary that we get from calling the parse function. We can assign it to a variable and store the dictionary for later use.

```python
feed = feedparser.parse("url_of_the_rss_feed")
```


## Extract the contents from the feed

Now, we have the dictionary of the feed, we can easily access the values from the listed keys. We can get the list of all the posts/podcasts/entries or any other form of content the feed is serving for from the `entries` key in the dictionary. 

To get more information and the most possible keys in the returned dictionary, you can refer to the feedparser [reference list](https://feedparser.readthedocs.io/en/latest/reference.html)

### Access Articles from Feed

To access the articles from the feed, we can access those as a list of the articles. Using the `entries` key in the dictonary as follows:

```python
feedparser.parse("url_of_the_rss_feed")["entries"]

OR

feedparser.parse("url_of_the_rss_feed").entries
```

If you have already defined a variable set to the parse function, you can use that for more efficient extraction.

```python
feed = feedparser.parse("url_of_the_rss_feed")

feed['entries']

OR 

feed.entries
```

### Get Number of Articles/Entries from Feed

To get the number of entries in the list, we can simply use the len function in python.

```python
len(feed.entries)

OR 

len(feedparser.parse("url_of_the_rss_feed").entries)
```

![Feedparser Number of Entries](https://res.cloudinary.com/techstructive-blog/image/upload/v1648371042/blog-media/didijxcvsgvl4scrnhje.png)

This gives us the number of entries in the provided feed. This is basically the list that stores all the content from the publication/website. So, we can iterate over the list and find all the different attributes we can extract.

### Get details of the entries from the feed

To get detail information about a particular article/entry in the feed, we can iterate over the `feed.entries` list and access what we require. 

So, we will iterate over the entries and simply print those entries one by one to inspect what and how we can extract them. 

```python
for entry in feed.entries:
  print(entry)
```

It turns out that every entry in the list is a dictionary again containing a few key-value pairs like `title`, `summary`, `link`, etc. To get a clear idea of those keys we can again use the keys function in python.

```python
feed = feedparser.parse("url_of_the_rss_feed")
print(feed.entries[0].keys())
```

![Feedparser Entries Keys](https://res.cloudinary.com/techstructive-blog/image/upload/v1648371221/blog-media/c8uog85goe9jzrzl1pq1.png)

```python
>>> feed.entries[0].keys()
dict_keys(['title', 'title_detail', 'authors', 'author', 'author_detail', 'published', 'published_parsed', 'links', 'link', 'id', 'guidislink', 'summary', 'summary_detail', 'tags'])
```

Now, we have all the keys associated with the entries we can now extract the specific details like the content, like `title`, `author`, `summary_detail`(actual content in this case).  

Though this might not be the same for all RSS feeds, it might be very similar and a matter of using the right keyword for the associated keys in the list of dictionaries.  

Let's say, we want to print out the titles of all the entries in the feed, we can do that by iterating over the entries list and fetching the title from the iterator as `entry.title` if `entry` is the iterator.

```python
for entry in feed.entries:
  print(entry.title)
```

![Feedparser List of Entries](https://res.cloudinary.com/techstructive-blog/image/upload/v1648372532/blog-media/lhofdzmr3ks0fuut7pxm.png)

Similarly, we will get the links of the entries using the link key in the dictionary.

```python
for entry in feed.entries:
  print(entry.link)
```

### Get information about the Website/Publication

To get the metadata about the information you are extracting from i.e. the website or any publication, we can use the key `feed`. This key stores another dictionary as its value which might have information like `title`, `description` or `subtitle`, `canonical_url`, or any other data related to the website company.

```python
feed.feed

or

feedparser.parse("url_of_the_rss_feed").feed
```

![Feedparser Feed](https://res.cloudinary.com/techstructive-blog/image/upload/v1648373487/blog-media/r7hiojfdrtrjqfhkjbdt.png)

From this dictionary, we can now simply extract the specific information from the keys. But first, as in the previous examples, we need a clear idea of what are the keys in the dictionary where we can extract the specific value.

```python
feed.feed.keys()

or

feedparser.parse("url_of_the_rss_feed").feed.keys()
```

Using the keys like `title`, `links`, `subtitle`, we can get the information on the website/company level and not related to the specific post in the entries list. 

```python
# get the title of the webpage/publication
feed.feed.title

# get the links associated with the webpage
feed.feed.links

# get the cover-image for the webpage
feed.feed.image
``` 

You can further get information specific to your feed. 

## Checking for keys existence in the dictionary of feed

We also need to check for the existence of a key in a dictionary in the provided feed, this can be a good problem if we are parsing multiple RSS feeds which might have a different structure. This problem occurred to me in the making of [podevcast](https://podevcast.netlify.app) where I had to parse multiple RSS feeds from different RSS generators. Some of the feeds had the cover image but most of them didn't. So, we need to make sure we have a check over those missing keys.

```python
feedlist = ['https://freecodecamp.libsyn.com/rss', 'https://feeds.devpods.dev/devdiscuss_podcast.xml']

for feed in feedlist:
    feed = feedparser.parse(feed)

    print(feed.entries[0].keys())
    for entry in feed.entries:
        if 'image' in entry:
            image_url = entry.image
        else:
            image_url = feed.feed.image
        
        #print(image_url)
```

```python
>>> feedlist = ['https://freecodecamp.libsyn.com/rss', 'https://feeds.devpods.dev/devdiscuss_podcast.xml']
>>> for feed in feedlist:
...     feed = feedparser.parse(feed)
...     for entry in feed.entries:
...             if 'image' in entry:
...                     image_url = entry.image
...             else:
...                     image_url = feed.feed.image
...     print(feed.entries[0].keys())
...

dict_keys(['title', 'title_detail', 'itunes_title', 'published', 'published_parsed', 'id', 'guidislink', 'links', 'link', 'image', 'summary', 'summary_detail', 'content', 'itunes_duration', 'itunes_explicit', 'subtitle', 'subtitle_detail', 'itunes_episode', 'itunes_episodetype', 'authors', 'author', 'author_detail'])

dict_keys(['title', 'title_detail', 'links', 'link', 'published', 'published_parsed', 'id', 'guidislink', 'tags', 'summary', 'summary_detail', 'content', 'subtitle', 'subtitle_detail', 'authors', 'author', 'author_detail', 'itunes_explicit', 'itunes_duration'])
```

As we can see we do not have an image key in the second RSS feed which means each entry doesn't have a unique cover image, so we have to fetch the image from the `feed` key then the `image` key in the entries list.

![Feedparser Cover Image Demo](https://res.cloudinary.com/techstructive-blog/image/upload/v1648373275/blog-media/fzdqie5dubigxzfhtv2x.png)


As we can see here, the image_url will pick up the `image` key in the dictionary if it is present else we will assign it to another URL which is the website/podcast cover image. This is how we handle exceptions in providing the keys when there are multiple feeds to be extracted though they are quite similar, they will have subtle changes like this that need to be handled and taken care of.


## Conclusion

From this little article, we were able to understand and use the feedparser Python package which can be used to extract information from different feeds. We saw how to extract contents for the entries, a number of entries in the feed, check for keys in the dictionary, and so on. Using Python's Feedparser package, some of the projects I have created include:

- [podevcast](https://podevcast.netlify.app)
- [dailydotdev-bookmark-cli](https://pypi.org/project/dailydotdev-bookmark-cli/)
- [Django Newsletter](https://github.com/Mr-Destructive/newsletter)

For further reading, you can specifically target a feed type by getting the appropriate methods from the feedparser [documentation](https://feedparser.readthedocs.io/en/latest/)

Thank you for reading, if you have any suggestions, additions, feedback, please let me know in the comments or my social handles below. Hope you enjoyed reading. Happy Coding :)
