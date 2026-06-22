---
title: "HTTP QUERY Method"
description: "Exploring the new HTTP Query method with first principles, what is the need for new method, and where it sits. We explore the limitations of GET and the unintuitive use of POST methods for filtering"
date: 2026-06-22
slug: "http-query-method"
tags: ["http","web","api"]
type: posts
hash: 2364bc96fdf2bf58aeecc42ac7f989bc062e7c2408052753a5f054663730437e
---
## Why a new method?
-----------------

We all know and love the 5 HTTP Methods right? GET, POST, PATCH, PUT, DELETE. Some of us even know a couple more with OPTIONS and HEAD. Then you might ask, why is there a need for another method like QUERY. Well, let me explain with an example and why it makes sense for HTTP to have a another method.

### Let's create an API

Of course, its HTTP, so we understand that we need to interact with the resources on the server. The best understandable example is creating an API. Let's create a simple API for a Blog. I know its a simple example, but developer (me) have the knack to make simple problems complex, or I should say users force into making the requirements complex. Whatever you think it is. I would start by creating a simple blog API that does CRUD for posts.

So, here's the simple spec:

```
POST   /posts
GET    /posts
GET    /posts/{id}
PATCH  /posts/{id}
DELETE /posts/{id}
```

Everything looks fine so far.

But we haven't defined what a Post actually looks like so, let's define one

```json
{
  "id": 1,
  "title": "Understanding HTTP Methods",
  "content": "The core HTTP methods are GET.....",
  "author": "Meet Gor",
  "tags": ["http", "web"],
  "status": "published",
  "published_at": "2026-06-22T10:00:00Z"
}
```

But users rarely stop at "give me all posts."

They want, I want your "TIL" posts, oh no maybe just "articles", or just "newsletter". Well. Now we need to add "type" as the field to sort out different types of articles.

```json
{
  "id": 1,
  "title": "Understanding HTTP Methods",
  "content": "The core HTTP methods are GET.....",
  "author": "Meet Gor",
  "type": "article",
  "tags": ["http", "web"],
  "status": "published",
  "created_at": "2026-06-22T10:00:00Z",
  "published_at": "2026-06-22T10:00:00Z"
}
```

Now you can say, it looks better and more extendable right? for now at least it does maybe. The type could be an enum in the api spec like

-   Article
-   TIL (today I learnt)
-   Newsletter
-   Thoughts
-   Link Post
-   Dev blog
-   Review
Etc....

But the point is. It could be one of some of the values, and the other end of it is tags, they could be many of some values.

Now, we can support filtering of those on the `GET /posts` method since that is where we list out all the posts.

Lets get all the published newsletter type posts

```http
GET /posts?status=published&type=newsletter
```

That is ideal. But this is a blog api. People can go crazy here:

```http
GET /posts?status=published&type=newsletter&tags=python&sort=-published_at
```

We'll have to support sorting, custom filters for tags and things like that. It can get messy pretty soon.

So. Let's create a separate endpoint for filtering post, it makes it isolated and clean. At least that's what I would think, if searching and filtering is very specific and needed all the time. To create a separate filter/query endpoint for that resource(s)

### Let's add an endpoint to filter things

That should be straight forward as :

```http
GET /posts/search?
```

This will make it explicit that we are searching for certain resource in this case "posts"

Now.

We simply do what we did before, we allow query parameters to filter the fields like so:

```http
GET /posts/search?status=published&type=newsletter&tags=python&sort=-published_at
```

This looks ok, but we can see the query in the URI getting too long. It might create issues right? There is a limit of the number of characters in the URL we can enter, plus it looks wired, too long.

### Limitation of GET

You can see, this is a bit of problem. There is a limitation of 2k characters in the query. I am not saying that this blog post search api will require 2k characters in the uri for filtering post, but think of similar use cases which might require that. Its not a good design but sometimes it becomes inevitable.

So, what is the other option?

Request body.

We can define something like this right?

```http
GET /posts/search/
Content-Type: application/json

{
  "type": "newsletter",
  "status": "published",
  "tags": "python",
  "sort": "-published_at"
}
```

The issue is that if we want to express this into a request body if the request grows, we hit a roadblock. There is no one stopping us from adding a request body in GET but its not defined in the specification. Its not a practical way and a documented way to follow a convention while sending request body in GET method. HTTP does not define semantics for a GET request body. Some clients and servers may support it, but relying on it makes interoperability unpredictable.

So, what is the better way?

### Let's use POST maybe

Well, this looks inevitable at this point.

Request Query Parameters are too confusing for readers and not practical for complex queries.

If we want to use request body the only option is to use POST method.

```http
POST /posts/search/
Content-Type: application/json

{
  "type": "newsletter",
  "status": "published",
  "tags": "python",
  "sort": "-published_at"
}
```

yes, this looks good right? For the problem development, yes, as a api developer, nope!

POST is not idempotent by definition. Repeating the same POST request may create multiple resources or trigger multiple side effects depending on the server implementation. For instance we send a request first, the resource is created, but the subsequent same request body looks different from the first right? since it depends on the server if it handles duplicates or allows creation of multiple resources.

Also the bigger case for this problem is this

### Using POST to query is unintuitive

POST to query or filter resources doesn't make sense right? We aren't creating any resources. So sending a POST request to filter a resource feels a bit unintuitive to me. Since we aren't creating any resources we should not be sending post requests to filter resources.

Just take a look at this query, if i didn't add search, how does that distinguish from posts to create and post to search, just adding /`search` to the path makes it easily identifiable.

```http
POST /posts/
POST /posts/search/

```

HTTP methods are not just verbs. They communicate intent. That's why we have PUT and PATCH differently separated and defined. The same is for GET and POST. For people or code generators (LLMs now) need to understand the documentation to make it clear that the POST /posts/search doesn't create a resource.

But there is no option to send complex data for filtering resources unless we specify a request body.

Until the RFC 10008!

HTTP now has added more method, the QUERY Method. That solves that exact problem of GET being too simple and POST being non-sensicle for filtering or querying data with complex filters.

## QUERY Method


The query method bridges the gap between the two methods. GET and POST.

The simplicity of GET request is too limiting.

The usage of POST request is not ideal for filtering.

```http
QUERY /posts/search/

Content-Type: application/json

{
  "type": "newsletter",
  "status": "published",
  "tags": "python",
  "sort": "-published_at"
}
```

HTTP has an method verb for that operation `QUERY` to filter/search or literal query the resource.

We can use the body as defined in the specification that can follow a standard like content-type which should be something that the server supports. let's get into the details of the specification.

## RFC 10008

The request for comments was first started in [2021](https://datatracker.ietf.org/doc/draft-ietf-httpbis-safe-method-w-body/00/) as SEARCH method. But it evolved into QUERY in 4-5 years to become [RFC 10008](https://datatracker.ietf.org/doc/html/rfc10008) in June 2026

The specification states that the server and the client needs to understand each other, the query that the client will send, the server needs to understand that medium right?

So, in `POST`, `PUT` or `PATCH` we use the `Accept header` from the server to understand that the server can understand a few set of formats like `application/json` or `application/xml` etc.

But think for a moment.

What can happen if we have this

```http
POST   /posts
GET    /posts
GET    /posts/{id}
PATCH  /posts/{id}
DELETE /posts/{id}
QUERY  /posts/
```

We need to know what methods the api supports right? before sending a `QUERY` request.

That's when we use `OPTIONS`. The request that sends a header as `Allow: OPTIONS, GET, HEAD, POST` . This means, the API endpoint supports these many request methods. For our case, we need to add the OPTIONS method

```http
OPTIONS /posts
POST    /posts
GET     /posts
GET     /posts/{id}
PATCH   /posts/{id}
DELETE  /posts/{id}
QUERY   /posts/
```

So, when we send

```http
OPTIONS /posts/ HTTP/1.1
```

The server should send something like this

```http
HTTP/1.1 204 No Content
Allow: OPTIONS, GET, HEAD, POST, PUT, PATCH, DELETE, QUERY
Date: Thu, 22 Jun 2022 11:45:00 GMT

```

From this Header of `ALLOW` we know that `QUERY` is supported on /posts/ endpoint.

We can now procee...

Wait, we need to know what content type is supported still.

Now, we also need

Server sends the Accept in the Response header of the OPTIONS request

There are a few ifs-elses here.

Why?

Let's see the problem first.

Let's say you send :

```
curl -X OPTIONS https://apis.meetgor.com/posts/
```

This is equivalent of:

```http
OPTIONS /posts/ HTTP/1.1
Host: apis.meetgor.com
```

This means, for the request of /posts/ what headers are given path on a endpoint. This results into something like this:

```http
HTTP/1.1 200 OK
Allow: OPTIONS, GET, HEAD, POST, PUT, PATCH, DELETE, QUERY
Accept: application/json, application/yaml

```

This means, the server accepts json and yaml in the request body.

But but..

There is where we need to make a distinction here.

The `Accept` can be for `POST` `PUT` or `PATCH` too right? How can we distinguish those from `QUERY`. Its likely we support different format for filtering/searching than for creating, listing or updating resources right?

That's the specification provides the semantics. It specifies `Accept-Query` as the header specifically for the `QUERY` method.

Since query itself has a representation for itself, it needs a accept header to identify what format the server can accept as a query.

So, once that is know, its a matter of time, to send the request body in that format in the `QUERY` method.

```http
QUERY /posts
Content-Type: application/json

{
  "status": "published",
  "type": "newsletter"
}
```

Done.

The query representation is something not a discussion for this post. Its quite taste and usecase specific thing. So we leave it to the developers in making and designing that format.

But the exchange and protocol we follow is this:

-   Check if `QUERY` method is supported with `OPTIONS` method
-   Also get the `Accept-Query` header to know what the server api supports for representing the query format.
-   Send the `QUERY` Method with the request body as the supported format for the query and get the response.

Simple.

That's the basics of HTTP QUERY method.

I will be making a API that supports QUERY method on my apis soon. Will be writing a write up like this for the same. So, stay tuned for that.