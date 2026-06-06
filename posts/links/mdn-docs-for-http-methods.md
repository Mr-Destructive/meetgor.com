---
title: "MDN Docs for HTTP Methods"
date: 2026-06-05
link: "https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Methods"
status: published
image_url: ""
source: newsletter
newsletter: techstructive-weekly-97
type: links
slug: mdn-docs-for-http-methods
tags:
description: "HTTP defines a set of request methods to indicate the purpose of the request and what is expected if the request is successful.
Although they can also be nouns, these request methods are sometimes referred to as HTTP verbs.
Each request method has its own semantics, but some characteristics are shared across multiple methods, specifically request methods can be safe, idempotent, or cacheable."
hash: a318e1e32fdd385237ff11f09f9f7eca0f55f7ac3c10d082e74721a849cf6b7e
---
My thoughts on [MDN Docs for HTTP Methods](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Methods): MDN Docs for HTTP Methods

## Commentary

- MDN Docs for HTTP Methods: I didn’t paid closer attention to these HEAD and OPTIONS method, but now all of them make sense. These are utility methods, nothing for data, but for communication meta information.
- OPTIONSThis is the request that asks for the available options in a API URI endpoint. It receives the response in response headers only, there is no response for a OPTIONS request. It usually has a 204 No content response status and rightly so, since its purpose is to make the caller(client) aware which methods are available on the API.It gets the headers like this `Access-Control-Allow-Methods: GET, POST, PUT, DELETE
- HEADThe header is powerful to get only the headers and not the actual response.It is usually used to check if a resource exists, save bandwidth. It might still fetch the resource till the request leaves the server.Pretty handy thing to know that this exists and can even use and implement one.
