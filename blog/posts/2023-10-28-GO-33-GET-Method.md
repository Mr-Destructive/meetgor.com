---
templateKey: blog-post
title: "Golang Web: GET Method"
description: "Exploring the fundamentals of a GET method in golang. How to request a resource, parse body, headers, etc. in a HTTP GET request."
date: 2023-10-28 20:15:00
status: published
slug: golang-web-get-method
tags: ['go',]
series: ['100-days-of-golang',]
image_url: https://meetgor-cdn.pages.dev/100-days-of-golang/golang-033-get-method.png
---

## Introduction

In this section of the series, we will be exploring how to send a `GET` HTTP request in golang. We will be understanding how to send a basic GET request, create an HTTP request and customize the client, add headers, read the response body, etc in the following sections of this post.

## What is a GET method?

A [GET](https://en.wikipedia.org/wiki/HTTP#Request_methods) method in the context of an HTTP request is an action that is used to obtain data/resources. The `GET` method is used in a web application to get a resource like an HTML page, image, video, media, etc.

Some common usecases of the `GET` method are:

- Loading a webpage
- Getting an image, file or other resource
- API requests to retrieve data
- Search queries sending filters and parameters

## Basic GET Request

To use the HTTP method `GET` in golang, the [net/http](https://pkg.go.dev/net/http) package has a [Get](https://pkg.go.dev/net/http#Get) method. This method simply takes in a URL as a string and returns the [response](https://pkg.go.dev/net/http#Response) or an error. Let's look at how to send a basic HTTP GET request in golang.

```go
// web/methods/get/main.go


package main

import (
	"fmt"
	"net/http"
)

func main() {
	reqURL := "https://www.google.com"
	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
    fmt.Println(resp.Status)
    fmt.Println("Status Code:", resp.StatusCode)
}
```

```
$ go run main.go

&{200 OK 200 HTTP/2.0 2 0 map[Alt-Svc:[h3=":443"; ma=2592000,h3-29=":443"; ma=2592000] Cache-Control:[private, max-age=0] Content-Security-Policy-Report-Only:[object-src 'none';base-uri 'self';script-src 'nonce-pdz_s8Gr0owwMbX8I9qNEQ' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp] Content-Type:[text/html; charset=ISO-8859-1] Date:[Fri, 27 Oct 2023 09:37:04 GMT] Expires:[-1] P3p:[CP="This is not a P3P policy! See g.co/p3phelp for more info."] Server:[gws] Set-Cookie:[1P_JAR=2023-10-27-09; expires=Sun, 26-Nov-2023 09:37:04 GMT; path=/; domain=.google.com; Secure AEC=Ackid1Q5FARA_9d7f7znggUdw6DoJA1DBpI17Z0SWxN519Dc64EqmYVHlFg; expires=Wed, 24-Apr-2024 09:37:04 GMT; path=/; domain=.google.com; Secure; HttpOnly; SameSite=lax NID=511=EToBPqckCVRE7Paehug1PgNBKqe7lFLI9d12xJrGbvP9r8tkFIRWciry3gsy8FZ8OUIK4gE4PD-irgNzg4Y1fVePLdyu0AJdY_HcqL6zQYok-Adn-y5TDPmMCNuDnrouBfoxtqVjYY_4RFOe8jalkYto5fQAwzWnNJyw8K0avsw; expires=Sat, 27-Apr-2024 09:37:04 GMT; path=/; domain=.google.com; HttpOnly] X-Frame-Options:[SAMEORIGIN] X-Xss-Protection:[0]] 0xc000197920 -1 [] false true map[] 0xc0000ee000 0xc0000d8420}

200 OK

Status Code: 200
```

In the above code, we have defined a URL string as `reqURL` and used the [Get](https://pkg.go.dev/net/http#Get) method to send a GET request. The `Get` is parsed with the `reqURL` string and the return values are stored as `resp` and `err`. We have added an error check after calling the `Get` method in order to avoid errors later in the code.

The `Get` method as seen from the output has returned a `*http.Response` object, we can use the `Status` and `StatusCode` properties to get the status code of the response. In this case, the status code of the response was `200`. The response `resp` is an object of type `http.Response` i.e. it has fields like `Body`, `StatusCode`, `Headers`, `Proto`, etc. We can get each individual field from the `resp` object. We will later look into how to read the `Body` field from the response, it is not directly read as a string nor it is stored in other forms, rather it is streamed from the requested URL.

## Creating a GET request

We can even construct a GET request using the [NewRequest](https://pkg.go.dev/net/http#NewRequest) method. This is a low-level way of creating a `GET` request. We mention the `method`, `URL`, and the `body`, in the case of `GET` request, there is nobody. So, the `NewRequest` is a general way of creating a `http` request.

```go
// web/methods/get/newreq.go

package main

import (
	"fmt"
	"net/http"
)

func main() {
	reqURL := "https://www.google.com"
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
```

As we can see, we construct a `GET` request using the `NewRequest` method and then use the [Do](https://pkg.go.dev/net/http#Client.Do) method to send the request to the server. The [http.DefaultClient](https://pkg.go.dev/net/http#DefaultClient) is used as a client to send the request, if we want to customize this we can create a new instance object of [http.Client](https://pkg.go.dev/net/http#Client) and use it to send requests. We will be taking a look at clients in another part of this series when we want to persist a connection or avoid connecting multiple times to the same application/URL.

For now, we will go ahead with the DefaultClient. This will trigger the request, in this case, a `GET` request to the specified URL in the `reqURL` string. The `Do` method returns either a `http.Response` or an `error` just like the `Get` method did.

## Reading the Response Body

We saw some different ways to send a `GET` request, now the below example will demonstrate how to read the body of the response. The response body is read from a buffer rather than loading the entire response into memory. It makes it flexible to parse the data efficiently and as per the needs. We will see how we use the [io](https://pkg.go.dev/io) package's [ReadAll](https://pkg.go.dev/io#ReadAll) method can be used to read from the response buffer.

```go
// web/methods/get/body.go

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	reqURL := "https://httpbin.org/html"
	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}
	// close the body object before returning the function
	// this is to avoid the memory leak
	defer resp.Body.Close()

	// stream the data from the response body only once
	// it is not buffered in the memory
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
```

In the above example, we are trying to get the body from the response to the request sent at [`https://httpbin.org/html`](https://httpbin.org/html). We have used the simple `Get` method instead of `NewRequest` and `Do` for simplicity. The response is stored in `resp`, we also have added `defer resp.Body.Close()` which is to say that we will close the body reader object when the function is returned/closed. So, this means that the `Body` is not readily available data, we need to obtain/stream the data from the server. We have to receive the body in bytes as a tcp request, the body is streamed in a buffer.

The response body is streamed from the server, which means that it's not immediately available as a complete data set. We read the body in bytes as it arrives over the network, and it's stored in a buffer, which allows us to process the data efficiently.

### Reading Body in bytes

We can even read the body in bytes i.e. by reading a chunk of the buffer at a time. We can use the [bytes.Buffer](https://pkg.go.dev/bytes#Buffer) container object to store the body. Then we can create a slice of bytes as `[]bytes` of a certain size and read the body into the chunk. By writing the chunk into the buffer, we get the entire body from the response.

```go
// web/methods/get/body.go


package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	reqURL := "https://httpbin.org/html"
	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

    // create a empty buffer
	buf := new(bytes.Buffer)

    // create a chunk buffer of a fixed size
	chunk := make([]byte, 1024)

	for {
		// Read into buffer
		n, err := resp.Body.Read(chunk)
		if err != nil {
			break
		}
        // append the chunk to the buffer
		buf.Write(chunk[:n])
		fmt.Printf("%s\n", chunk[:n])
	}

    // entire body stored in bytes
	fmt.Println(buf.String())
}
```

In the above example, the body is read chunkwise buffers and obtained as a slice of bytes. We define the request as a `GET` request to the [`httpbin.org/html`](http://httpbin.org/html). We create a new Buffer as a slice of bytes with [bytes.Buffer](https://pkg.go.dev/bytes#Buffer), then we define chunk as a container to stream the response body with a particular size. We have taken `1024` bytes as the size of the chunk. Then inside an infinite for loop, we read the body as `n, err :=` [`resp.Body.Read`](http://resp.Body.Read)`(chunk)`. The code will read the body into the chunk(slice of byte) and the return value will be the size of the bytes read or an error. Then we check if there is no error, and if there is an error, we break the loop indicating we have completed reading the entire body or something went wrong. Then we write the chunk into the buffer that we allocated earlier as `buf`. This is a slice of bytes, we are basically populating the buffer with more slices of bytes.

The entire body is then stored in the buffer as a slice of bytes. So, we have to cast it into a string to see the contents. So, this is how we can read the contents of a body in a response in chunks.

## Adding Headers to a GET Request

We can even add headers before sending a `GET` request to a URL. By creating a `Request` object with the `NewRequest` method and adding a [Header](https://pkg.go.dev/net/http#Header) with the [Add](https://pkg.go.dev/net/http#Header.Add) method. The `Add` method will take in two parameters i.e. the key of the header, and the value of the header both as strings.

```go
// web/methods/get/header.go

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/users/mr-destructive", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "token YOUR_TOKEN")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
```

```
$ go run web/methods/get/header.go

{"message":"Bad credentials","documentation_url":"https://docs.github.com/rest"}
```

In the above example, we have created a `GET` request to the [`https://api.github.com/users/mr-destructive`](https://api.github.com/users/mr-destructive) the last portion is the username, it could be any valid username. The request is to the GitHub API, so it might require API Key/Tokens in the headers, however, if there are certain endpoints that do not require Authorization headers might work just fine.

So, the above code will give `401` error indicating the request has wrong or invalid credentials, if we remove the header, the request will work fine. This is just an example, but headers are useful in working with APIs.

Without adding the header:

```
$ go run web/methods/get/header.go

{"login":"Mr-Destructive","id":40317114,"node_id":"MDQ6VXNlcjQwMzE3MTE0","avatar_url":"https://avatars.githubusercontent.com/u/40317114?v=4","gravatar_id":"","url":"https://api.github.com/users/Mr-Destructive",
... 
"updated_at":"2023-10-10T17:57:22Z"}
```

That's it from the 33rd part of the series, all the source code for the examples are linked in the GitHub on the [100 days of Golang](https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/get/) repository.

[100-days-of-golang](https://github.com/Mr-Destructive/100-days-of-golang)

## Conclusion

From this article, we explored the `GET` HTTP method in golang. By using a few examples for creating a get request, adding headers, reading response body, the basic use cases were demonstrated.

Hopefully, you found this section helpful, if you have any comments or feedback, please let me know in the comments section or on my social handles. Thank you for reading. Happy Coding :)
