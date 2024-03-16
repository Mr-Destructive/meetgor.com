---
templateKey: blog-post
title: "Golang Web: URL Parsing"
description: "Understanding the fundamentals of web development with URL parsing in Golang. Intro to the net package in golang."
date: 2023-09-05 21:30:00
status: published
slug: golang-web-url-parsing
tags: ['go',]
series: ['100-days-of-golang',]
image_url: https://meetgor-cdn.pages.dev/100-days-of-golang/golang-032-url-parsing.png
---

## Introduction

We have done around 32 posts on the fundamental concepts in golang, With that basic foundation, I'd like to start with the new section of this series which will be a major one as `web-development`. This section will have nearly 40-50 posts, this will cover the fundamental concepts for web development like APIs, Database integrations, Authentication and Authorizations, Web applications, static sites, etc.

## What is a URL?

A URL is a Uniform Resource Locator. It is a string of characters that identifies a resource on the Internet. URLs are the building blocks of the web, allowing us to access websites, documents, and data with just a click. URLs are all over the place, if we want to build a strong foundation in web development, it's quite important to understand what URLs actually mean and what can they store.

A URL looks something like this:

```
[scheme:][//[userinfo@]host][/]path[?query][#fragment]
```

Not all the URLs are like this, majority of the URLs that the common user sees are simply the ones with the `scheme`, `host`, and `paths`. However other components are equally important and are vital in the exchanging of information over the network. 

- The `scheme` is the protocol used for accessing the resource like `http`, `https`, `ftp`, etc.
- The `userinfo` is the username and password used to access the resource.
- The `host` is the domain name of the resource.
- The `path` is the path or folder to the resource.
- The `query` is the query string of the resource. It is usually a key-value pair as a paramter to access resources.
- The `fragment` is used as a reference within the resource.

We will see the use cases of most of them throughout this series for example, the `userinfo` is commonly used in accessing databases over the internet/cloud. The query parameters will be used in making dynamic API calls, etc.

## Basic URL Parsing

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	// simple url
	urlString := "http://www.google.com"
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedUrl)
}
```

```
$ go run main.go

http://www.google.com
```


So, what is getting parsed here, we gave the URL as a string we get the URL back, the only difference is that instead of the URL being a string, it is now a structure of components. For instance, we want the protocol the host name, the port, etc. from the URL.


```
fmt.Printf("%T\n", parsedUrl)
// *url.URL
```

The `parsedUrl` is a pointer to a [url.URL](https://pkg.go.dev/net/url#URL) structure. The structure `url.URL` has a lot of components to it like `Scheme`, `Host`, `User`, `Path`, `RawQuery`, etc. We will dive into each of these soon.

We could get those specific components ourselves, but that would be a bit tedious and might be even prone to bugs.

Let's try to get those components from the URL without any functions, just string manipulation.

```go
package main

import (
    "fmt"
)

func main() {
    urlString := "http://www.google.com"
    protocol := urlString.split(":")[0]
    hostName := urlString.split("//")[1]
    fmt.Println(protocol)
    fmt.Println(hostName)
}
```

This might work for a simple URL, but what if we have more complex URLs which have paths, query parameters, fragments, username, port, etc? This could mess up quickly if we tried to get the parts of the URL ourselves. So, golang has a package called [net/url](https://pkg.go.dev/net/url) explicitly for parsing URLs.

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	urlString := []string{"http://www.google.com",
		"http://www.google.com/about/",
		"http://www.google.com/about?q=hello",
		"http://www.google.com:8080/about?q=hello",
		"http://user:password@example.com:8080/path/to/resource?query=value#fragment",
	}
	for _, url := range urlString {
		hostStr := strings.Split(url, "//")[1]
        hostName := strings.Split(hostStr, "/")[0]
		fmt.Println(hostName)
	}
}
```

```
$ go run main.go

www.google.com
www.google.com
www.google.com
www.google.com:8080
user:password@example.com:8080
```

The above code might work for most of the URLs, but what if we have a more complex URL like the one with `port` or `user`, it doesn't give the thing we want exactly. In the above example, we have created a list of URLs as strings and simply iterated over each of the `url` in the `urlString`. Thereafter, we split the `url` on `//` so we get `https:` and `www.google.com`, if we want the host/domain name, we could simply get the `1` index in the slice since the [strings.Split](https://pkg.go.dev/strings#Split) method returns a slice after splitting the string with the provided separator. The `hostName` could be fetched from the `1` index. However, this time for the 2nd element in the list, we have `https://www.google.com/about`, which would return `www.google.com/about` as the hostname which is not ideal, so we will again have to split this string with `/` and grab the first part i.e. 0th index.

The above code would work for `paths` and `query` parameters but if we had ports, and username, and password, it would not work as expected as evident from the last 2 examples in the list.

So, now we know the downsides of manually parsing the URLs, we can use the [net/url](https://pkg.go.dev/net/url) package to do it for us.

## Parsing DB URLs

Databases have a connection URL or connection string which provides a standard way to connect to a database/database server. The format of the URL is just the `URL` with all the components from the `scheme` to the `path`. The common examples of some database connection URLs might include:


```
# PostgreSQL DB Connection URL/string
postgresql://username:password@hostname:port/database_name

# MongoDB Connection URL/string
mongodb://username:password@hostname:port/database_name
```

The above are examples of the Postgres and MongoDB connection URLs, they have a `protocol` which usually for databases is their short name, the user credentials i.e. `username` and `password`, the `hostname` i.e. the server IP address, the `port` on which the database is running, and finally the path as the `database name`.

We can construct a simple snippet in golang to grab all the details from the simple connection URL string with the `net/url` package.

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	//postgres db url
	dbUrl, err := url.Parse("postgres://admin:pass1234@localhost:5432/mydb")
	if err != nil {
		panic(err)
	}
	fmt.Println(dbUrl)
	fmt.Println("Scheme/Protocol = ", dbUrl.Scheme)
	fmt.Println("User = ", dbUrl.User)
	//fmt.Println("User = ", dbUrl.User.String())
	fmt.Println("Username = ", dbUrl.User.Username())
	password, _ := dbUrl.User.Password()
	fmt.Println("Password = ", password)
	fmt.Println("Host = ", dbUrl.Host)
	fmt.Println("HostName = ", dbUrl.Hostname())
	fmt.Println("Port = ", dbUrl.Port())
	fmt.Println("DB Name = ", dbUrl.Path)
}
```

```
$ go run main.go

postgres://postgres:pass1234@localhost:5432/mydb
Scheme/Protocol =  postgres
User = admin:pass1234
Username =  admin
Password =  pass1234
Host =  localhost:5432
HostName =  localhost
Port =  5432
DB Name =  mydb
```

In the above code, we have given the string `postgres://admin:pass1234@localhost:5432/mydb`, and we have parsed the URL using the `net/url` package. The result is we have a `parsedUrl` object which has all the components that can be accessed as either fields or methods. Let's break down each field/method we used in the above example:

- The `Scheme` is simply a string representing the protocol of the resource(URL).
- The `User` is the [UserInfo](https://pkg.go.dev/net/url#Userinfo) object having immutable username and password fields.
- The `Username` is the method on [UserInfo](https://pkg.go.dev/net/url#Userinfo.Username) that returns the string representing the username of the URL.
- The `Password` is the method on [UserInfo](https://pkg.go.dev/net/url#Userinfo.Password) that returns the string representing the password of the URL.
- The `Host` is the field on `URL` as a string representing the host:port of the URL.
- The `Hostname` is the method on [URL](https://pkg.go.dev/net/url#URL.Hostname) that returns the string representing the hostname of the URL.
- The `Port` is the method on [URL](https://pkg.go.dev/net/url#URL.Port) that returns the string representing the port of the URL.
- The `Path` is the field as the string representing the path of the URL.

So, this is how we can get almost every detail like `db protocol`, `username`, `password`, `hostname`, `port`, and the `database name` from the database connection string URL.


## Parsing Query Parameters

We can even get the query parameters of a URL using the [Query](https://pkg.go.dev/net/url#URL.Query) method on the `URL` object. The `Query` method returns a `map[string][]string` which is to say a map with the key as `string` and the value as a `[]string` slice of string. For example, if the URL is something like `https://google.com/?q=hello`, the `Query` method will return `map[q:[hello]]` which means the key is `q` and the value is a list of strings of which the only element is `hello`.

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	// a complex url with query params
	urlStr := "http://www.google.com/?q=hello&q=world&lang=en"
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedUrl)
	fmt.Println(parsedUrl.Query())
	for k, v := range parsedUrl.Query() {
		fmt.Println("KEY:", k)
		for _, vv := range v {
			fmt.Println(vv)
		}
	}
}
```

```
$ go run main.go

http://www.google.com/?q=hello+world&lang=en&q=gopher
map[lang:[en] q:[hello world gopher]]
KEY: q
hello world
gopher
KEY: lang
en
```

We have taken a bit of a complex example that might cover many use cases of the `Query` method. We have a URL as `http://www.google.com/?q=hello&q=world&lang=en`, and the `Query` method returns `map[lang:[en] q:[hello world]]` which means the `q` key has a slice of a string with elements `hello world` and `gopher` and the `lang` key has a value of `en`. Here, the first paramter, `q=hello+world` is basically `hello world` or `hello%20world`, which is to say escaping the space in the URL. We can have multiple values for the same key, which is evident as we have added `q=gopher` at the end of the `URL`, the key `q` has two elements in the slice as `hello world` and `gopher`. The `lang=en` is simply a key as `lang` with the only element as `en` in the slice. We use `&` to separate different query parameters in the URL.

### Checking Values in Query Parameters

We can even check the values in the query parameters without requiring the construction of for loops to find a particular value in a key. The [Values](https://pkg.go.dev/net/url#Values) is a type that stores the map as a return value from the `Query` method. It has a few handy methods like:
- [Has](https://pkg.go.dev/net/url#Values.Has) to check if the key exists in the map (paramter as key `string` and returns a `bool`).
- [Get](https://pkg.go.dev/net/url#Values.Get) to fetch the first value of the given key as a string or if not present then returns an empty string (paramter as key `string` and returns `string`).
- [Add](https://pkg.go.dev/net/url#Values.Add) method is used to append the value for a given key (paramter as key `string` and value to be added as `string`).
- [Set](https://pkg.go.dev/net/url#Values.Set) method is used to replace the value for a given key if already exists (paramter as key `string` and value as `string`).
- [Del](https://pkg.go.dev/net/url#Values.Del) method is used to delete the value for a given key (paramter as key `string`).

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	// a complex url with query params
	urlStr := "http://www.google.com/?q=hello+world&lang=en&q=gopher"
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedUrl)
	fmt.Println(parsedUrl.Query())

	queryParams := parsedUrl.Query()

	fmt.Println(queryParams.Get("q"))

	fmt.Println(queryParams.Has("q"))

	if queryParams.Has("lang") {
		fmt.Println(queryParams.Get("lang"))
	}

	queryParams.Add("q", "ferris")
	fmt.Println(queryParams)

	queryParams.Set("q", "books")
	fmt.Println(queryParams)

	queryParams.Del("q")
	fmt.Println(queryParams)
}
```

```
$ go run main.go

http://www.google.com/?q=hello+world&lang=en&q=gopher
hello world
true
en
map[lang:[en] q:[hello world gopher ferris]]
map[lang:[en] q:[books]]
map[lang:[en]]
```

The above code example demonstrates almost all the methods available on the `Values` type. The `Get` method is used to fetch the first value for a given key, so we parse the key as a `string` to the method and it would return a `string`. We checked for the `q` as the key and it returned the first element in the `queryParams` for key `q` as `hello world` from the list `[hello world, gopher]`. The `Has` method takes in the paramter as key as `string` and returns if the key exists in the `queryParams` or not as a bool. The `Add` method, we have used to `Add` a key with a particular value, we added the value `ferris` to the key `q` hence it appended to the list and the list of `queryParams[q]` became `[hello world, gopher, ferris]`. The `Set` method is used to override the existing key with a particular value, here we have set the value `books` to the key `q` and hence the list of `queryParams[q]` becomes `[books]`. We can use the `Del` method to remove the key from the `queryParams`, so we delete `q` from `queryParams`, then the `queryParams` simply has no key as `q` in it.

### Parsing Query Parameters to String

Now, that you have manipulated the query parameters, let's say you want to construct back that string representation of the query particular or the URL for it. The [Encode]() method is used to grab the `queryParams` i.e. `Values` object and convert it into the `string` representation of the encoded URL.

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	// a complex url with query params
	urlStr := "http://www.google.com/?q=hello+world&lang=en&q=gopher"
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
    queryParams := parsedUrl.Query()
    queryParams.Add("name", "ferris")

    q := queryParams.Encode()
    fmt.Println(q)
}
```

```
$ go run main.go

lang=en&name=ferris&q=hello+world&q=gopher
```

So, we can see the `Encode` method has given a query parameter in the form of a URL encoded string. We first grab the query parameters from the `parsedUrl` which is a `URL` object via the `Query` method, we then `Add` the key `name` with a value of `ferris` to the `queryParams`. This is then used to `Encode` the object back to a string representation. This could be useful to construct a query paramter for requesting other websites/APIs.

## Parsing URL object back to String

We can even get the `URL` object back to a string representation using the [String](https://pkg.go.dev/net/url#URL.String) method on the `URL` object. The `String` method returns a `string` representation of the URL object.

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlStr := "http://www.google.com/?q=hello+world&lang=en&q=gopher"
	fmt.Println("URL:", urlStr)
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	queryParams := parsedUrl.Query()
	queryParams.Add("name", "ferris")

	q := queryParams.Encode()
	fmt.Println(q)
	parsedUrl.RawQuery = q
	newUrl := parsedUrl.String()
	fmt.Println("New URL:", newUrl)
}
```

```
$ go run main.go

URL: http://www.google.com/?q=hello+world&lang=en&q=gopher
lang=en&name=ferris&q=hello+world&q=gopher
New URL: http://www.google.com/?lang=en&name=ferris&q=hello+world&q=gopher
```

In the example above, we parse a URL string into a `URL` object as `parsedUrl`, then we `Add` the key `name` with a value of `ferris` to the `queryParams`. We then `Encode` the `URL` object back to a string representation. But this won't change the `parsedUrl` object we want to change the entire `URL` object. For that, we have overwritten the `RawQuery` field of the `URL` object with the query parameter encoded string as `q`. The `String` method returns a `string` representation of the `URL` object.

## Parsing Fragments

The fragment in a URL is usually present in a static website like `#about`, `#contact`, `#blog`, etc. The `Fragment` is a string that is usually a reference to a specific section or anchor point within a web page or resource. When a URL with a fragment is accessed, the web browser or user agent will scroll the page to display the section identified by the fragment.

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	// url with fragments
	urlStr := "https://pkg.go.dev/math/rand#Norm Float64"
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedUrl)
	fmt.Println(parsedUrl.Fragment)
	fmt.Println(parsedUrl.RawFragment)
	fmt.Println(parsedUrl.EscapedFragment())
}
```

```
$ go run main.go

https://pkg.go.dev/math/rand#Norm Float64

Norm Float64
Norm Float64
Norm%20Float64
```

The above code is used to fetch the `#Norm Float64` fragment from the URL `https://pkg.go.dev/math/rand#NormFloat64`. We can use the [Fragment](https://pkg.go.dev/net/url#URL) field in the `URL` object to get the fragment text. There is [RawFragment](https://pkg.go.dev/net/url#URL) field that is used to parse the fragment text as it is, without trying to escape any special characters in the URL. The [EscapedFragment](https://pkg.go.dev/net/url#URL.EscapedFragment) is used to parse the fragment text by escaping the characters in the URL.

That's it from the 32nd part of the series, all the source code for the examples are linked in the GitHub on the [100 days of Golang](https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/url-parsing) repository.

[100-days-of-golang](https://github.com/Mr-Destructive/100-days-of-golang)

## Conclusion

From the first post of the web development section, we covered the fundamentals of URL parsing and got a bit introduced to the `net` package, which will be heavily used for most of the core language's features for working with the web. We covered the concepts for parsing URLs, getting components of URLs from the parsed object, Database connection URL resolving, parsing query parameters, and some other URL-related concepts.

Hopefully, you found this section helpful, if you have any comments or feedback, please let me know in the comments section or on my social handles. Thank you for reading. Happy Coding :)
