---
templateKey: blog-post
title: "Safely using Maps in Golang: Differences in declaration and initialization"
description: "Walkthrough the differences and pitfalls in declaring and initializing maps in Golang"
date: 2024-08-31 18:30:00
status: published
slug: golang-safely-using-maps
tags: ['go',]
image_url: https://meetgor-cdn.pages.dev/golang-safe-map-usage.png
---

## Introduction

This week, I was working on one of the API wrapper packages for golang, and that dealt with sending post requests with URL encoded values, setting cookies, and all the fun stuff. However, while I was constructing the body, I was using [url.Value](https://pkg.go.dev/net/url#Values) type to construct the body, and use that to add and set key-value pairs. However, I was getting a wired `nil` pointer reference error in some of the parts, I thought it was because of some of the variables I set manually. However, by debugging closer, I found out a common pitfall or bad practice of just declaring a type but initializing it and that causing `nil` reference errors.

In this post, I will cover, what are maps, how to create maps, and especially how to properly declare and initialize them. Create a proper distinction between the declaration and initialization of maps or any similar data type in golang.

## What is a map in Golang?

A [map](https://go.dev/src/runtime/map.go) or a hashmap in golang is a basic data type that allows us to store key-value pairs. Under the hood, it is a header map-like data structure that holds buckets, which are basically pointers to bucket arrays (contiguous memory). It has hash codes that store the actual key-value pairs, and pointers to new buckets if the current overflows with the number of keys. This is a really smart data structure that provides almost constant time access.

## How to create maps in Golang

To create a simple map in golang, you can take an example of a letter frequency counter using a map of string and integer. The map will store the letters as keys and their frequency as values.

```go
package main

import "fmt"

func main() {
    words := "hello how are you"
    letters := map[string]int{}

    for _, word := range words {
        wordCount[word]++
    }

    fmt.Println("Word counts:")
    for word, count := range wordCount {
        fmt.Printf("%s: %d\n", word, count)
    }
}
```

```bash
$ go run main.go

Word counts:
e: 2
 : 3
w: 1
r: 1
y: 1
u: 1
h: 2
l: 2
o: 3
a: 1
```

So, by initializing the map as `map[string]int{}` you will get an empty map. This can be then used to populate the keys and values, we iterate over the string, and for each character (rune) we cast that byte of character into the string and increment the value, the zero value for int is `0`, so by default if the key is not present, it will be zero, it is a bit of double-edged swords though, we never know the key is present in the map with the value `0` or the key is not present but the default value is `0`. For that, you need to check if the `key` exists in the map or not.

For further details, you can check out my [Golang Maps](https://www.meetgor.com/golang-maps/) post in detail.

## Difference between declaration and initialization

There is a difference in declaring and initializing any variable in a programming language and has to do a lot more with the implementation of the underlying type. In the case of primary data types like `int`, `string`, `float`, etc. there are default/zero values, so that is the same as the declaration and initialization of the variables. However, in the case of maps and slices, the declaration is just making sure the variable is available to the scope of the program, however for initialization is setting it to its default/zero value or the actual value that should be assigned.

So, declaration simply makes the variable available within the scope of the program. For maps and slices, declaring a variable without initialization sets it to `nil`, meaning it points to no allocated memory and cannot be used directly.

Whereas the `initialization` allocates memory and sets the variable to a usable state. For maps and slices, you need to explicitly initialize them using syntax like `myMap = make(map[keyType]valueType)` or `slice = []type{}`. Without this initialization, attempting to use the map or slice will lead to runtime errors, such as panics for accessing or modifying a nil map or slice.

Let's looks at the values of a map when it is declared/initialized/not initialized.

Imagine you're building a configuration manager that reads settings from a map. The map will be declared globally but initialized only when the configuration is loaded.

1. Declared but not initialized
    

The below code demonstrates a map access that is not initialized.

```go
package main

import (
	"fmt"
	"log"
)

// Global map to store configuration settings
var configSettings map[string]string // Declared but not initialized

func main() {
	// Attempt to get a configuration setting before initializing the map
	serverPort := getConfigSetting("server_port")
	fmt.Printf("Server port: %s\n", serverPort)
}

func getConfigSetting(key string) string {
	if configSettings == nil {
		log.Fatal("Configuration settings map is not initialized")
	}
	value, exists := configSettings[key]
	if !exists {
		return "Setting not found"
	}
	return value
}
```

```bash
$ go run main.go
Server port: Setting not found
```

2. Declared and Initialized at the same time
    

The below code demonstrates a map access that is initialized at the same time.

```go
package main

import (
	"fmt"
	"log"
)

// Global map to store configuration settings
var configSettings = map[string]string{
	"server_port":  "8080",
	"database_url": "localhost:5432",
}

func main() {
	serverPort := getConfigSetting("server_port")
	fmt.Printf("Server port: %s\n", serverPort)
}

func getConfigSetting(key string) string {
	value, exists := configSettings[key]
	if !exists {
		return "Setting not found"
	}
	return value
}
```

```bash
$ go run main.go
Server port: 8080
```

3. Declared and later initialized
    

The below code demonstrates a map access that is initialized later.

```go
package main

import (
	"fmt"
	"log"
)

// Global map to store configuration settings
var configSettings map[string]string // declared but not initialized

func main() {
	// Initialize configuration settings
	initializeConfigSettings()
    // if the function is not called, the map will be nil

	// Get a configuration setting safely
	serverPort := getConfigSetting("server_port")
	fmt.Printf("Server port: %s\n", serverPort)
}

func initializeConfigSettings() {
	if configSettings == nil {
		configSettings = make(map[string]string) // Properly initialize the map
		configSettings["server_port"] = "8080"
		configSettings["database_url"] = "localhost:5432"
		fmt.Println("Configuration settings initialized")
	}
}

func getConfigSetting(key string) string {
	if configSettings == nil {
		log.Fatal("Configuration settings map is not initialized")
	}
	value, exists := configSettings[key]
	if !exists {
		return "Setting not found"
	}
	return value
}
```

```bash
$ go run main.go
Configuration settings initialized
Server port: 8080
```


In the above code, we declared the global map `configSettings` but didn't initialize it at that point, until we wanted to access the map. We initialize the map in the main function, this main function could be other specific parts of the code, and the global variable `configSettings` a map from another part of the code, by initializing it in the required scope, we prevent it from causing nil pointer access errors. We only initialize the map if it is `nil` i.e. it has not been initialized elsewhere in the code. This prevents overriding the map/flushing out the config set from other parts of the scope.

## Pitfalls in access of un-initialized maps

But since it deals with pointers, it comes with its own pitfalls like nil pointers access when the map is not initialized.

Let's take a look at an example, a real case where this might happen.

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
        var vals url.Values
        vals.Add("foo", "bar")
        fmt.Println(vals)
}
```

This will result in a runtime panic.

```nginx
$ go run main.go
panic: assignment to entry in nil map

goroutine 1 [running]:
net/url.Values.Add(...)
        /usr/local/go/src/net/url/url.go:902
main.main()
        /home/meet/code/playground/go/main.go:10 +0x2d
exit status 2
```

This is because the [url.Values](https://pkg.go.dev/net/url#Values) is a map of string and a list of string values. Since the underlying type is a map for `Values`, and in the example, we only have declared the variable `vals` with the type `url.Values`, it will point to a `nil` reference, hence the message on adding the value to the type. So, it is a good practice to use `make` while declaring or initializing a map data type. If you are not sure the underlying type is `map` then you could use `Type{}` to initialize an empty value of that type.

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
        vals := make(url.Values)
        // OR
        // vals := url.Values{}
        vals.Add("foo", "bar")
        fmt.Println(vals)
}
```

```bash
$ go run urlvals.go
map[foo:[bar]]
foo=bar
```

It is also recommended by the [golang team](https://go.dev/blog/maps) to use the make function while initializing a map. So, either use `make` for maps, slices, and channels, or initialize the empty value variable with `Type{}`. Both of them work similarly, but the latter is more generally applicable to structs as well.

## Conclusion

Understanding the difference between declaring and initializing maps in Golang is essential for any developer, not just in golang, but in general. As we've explored, simply declaring a map variable without initializing it can lead to runtime errors, such as panics when attempting to access or modify a nil map. Initializing a map ensures that it is properly allocated in memory and ready for use, thereby avoiding these pitfalls.

By following best practices—such as using the make function or initializing with Type{}—you can prevent common issues related to uninitialized maps. Always ensure that maps and slices are explicitly initialized before use to safeguard against unexpected nil pointer dereferences

Thank you for reading this post, If you have any questions, feedback, and suggestions, feel free to drop them in the comments.

Happy Coding :)
