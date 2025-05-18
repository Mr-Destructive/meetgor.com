{
  "type": "posts",
  "title": "Golang: Error Handling",
  "description": "Understanding the concept of errors in golang. How to handle errors in funciton calls, create custom error class with interfaces",
  "date": "2022-09-17 20:15:00",
  "status": "published",
  "slug": "golang-error-handling",
  "tags": [
    "go"
  ],
  "series": [
    "100-days-of-golang"
  ],
  "image_url": "https://meetgor-cdn.pages.dev/golang-020-error-handling.png"
}

## Introduction

Error handling is quite an important feature of any programming language to improve the quality and transparency between the user and the application. By raising appropriate error messages, the user can get a clear idea about the things happening in the interface as well as the application can handle the errors with appropriate actions.

In the 20th post of the series, we will be exploring the concept of error handling in golang. From this article, we will be able to learn the fundamentals of error or exception handling in golang, create custom error classes, raise and ignore error messages, and exit or redirect the flow of state of the application when an error is raised.

## Why we need Error Type

We need [error handling](https://go.dev/blog/error-handling-and-go) and catching in order to stop or divert the flow of the application which will restrict the dubious or unintentional execution of the code. Let's say, for example, we have a string as an input, and the user skipped the input and the string is returned as empty, we don't want to execute the further program as the execution might depend on the value of the string. So, in order to catch these kinds of events, we might use errors and log the results for better transparency.

```go
package main

import (
	"fmt"
)

func main() {

	var s string
	n, err := fmt.Scanf("%s", &s)

	if err != nil {
		fmt.Println(err)
		// panic(err)
        // OR
        // return
	} else {
		fmt.Println(n)
		if s[0] == 'a' {
			fmt.Println(s)
		}
	}
}
```

```
$ go run main.go
asdf
1
asdf


$ go run main.go

unexpected newline


$ go run main.go
wsd
1
```

In the above example, we have a simple string `s` input, we will input the string using the `Scanf` function that will return an integer as the number of variables it has scanned and error if any. Here, as the function might return two values, we need two variables to call the function. The `n` variable stores the number of variables successfully scanned and the `err` as the err from the function generated. If there is an error, that is the value stored in `err` is not empty, we will log the error. And move into the rest of the program.

This might be looking cool, but it doesn't break out of the program if there is an error. We want it to log the error as well as exit from the program. We can do that using the panic function.

## Catching Errors

We'll see a few examples, where we will catch errors in some of the regularly used functions. These error messages can be used in deciding the next procedure to be run.

### Comma OK/Error syntax

We use the [comma ok, error](https://go.dev/doc/effective_go#:~:text=is%20called%20the%20%E2%80%9C-,comma%20ok,-%E2%80%9D%20idiom.%20In%20this) syntax where we want multiple return values from a function. It is a narrowed syntax for a two-type return value function, we either return a value that we are expecting as `OK`, or we get an `error` from the function call.

```go

ok, err := function()
if err != nil {
    // handle error
    panic(err)
} else {
    // work with the ok object
    fmt.Println(ok)
}

```

In the above code, we have used the comma-ok, error syntax, the function call will return two objects either an expected object or an error object if there were errors in the processing. We handle the error if the error object is not empty i.e. it contains something, else we can do the rest of the required processing of the program.

We can even ignore the `err` or the `ok` object using the `_` i.e. to say a don't care variable. Remember you can ignore either of the values and not both. It is not recommended to ignore errors but if you know the obvious thing to process, you might as well sometimes.

A more compressed code might look something like below:

```go

if ok, err := function(); err != nil {
    // handle error
    panic(err)
} else {
    // work with the ok object
    fmt.Println(ok)
}
```

The above code wraps the initialization of `ok, err` objects by calling the function inside the if statement and checking for the error.

### Making HTTP requests

Let's say we have a URL, and we want to check if the site exists or not, we can run the [http.Get](https://pkg.go.dev/net/http@go1.19.1#Client.Get) function from the [net/http](https://pkg.go.dev/net/http) package. We will parse a URL in the function and this function also returns a [Response](https://pkg.go.dev/net/http#Response) type object and an error object if there are any errors generated during the function call. If there is an error while we call the function, we simply log the error and panic out of the program. Else we can log the status code and do further processing.

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://meetgor.com/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println(resp.StatusCode)
	}
}
```

```
$ go run web.go
URL: https://meetgor.com/
200



$ go run web.go

URL: htts://meetgor.com/
Get "htts://meetgor.com/": unsupported protocol scheme "htts"
panic: Get "htts://meetgor.com/": unsupported protocol scheme "htts"

goroutine 1 [running]:
main.main()
        /home/meet/code/100-days-of-golang/scripts/errors/https.go:14 +0x170
exit status 2
```


This is how we can validate a URL handling the error if the parsed URL is invalid or does not exist.

There is one more variation of the above code style, it is a bit compressed and might be just a syntactic change.

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://meetgor.com/"
	if resp, err := http.Get(url); err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println(resp.StatusCode)
	}
}
```

This can be used wherever you are using the `ok, err` kind of syntax, but I prefer the clean syntax so I won't move ahead with this.

### Opening or Handling of File

We can even use error handling while dealing with Files or Folders. We can use the [os](https://pkg.go.dev/os) package to read the file in golang. The [Open](https://pkg.go.dev/os#Open) function will read the file if it exists or else it will return an error. We can catch the error from the comma `ok,error` syntax and do the required processing in the program.


```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file_name := "hello.txt"
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println("Error: ", err)

        // Create a File 

		// _, err := os.Create(file_name)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println("Created File", file_name)
		// file, _ = os.Open(file_name)
	}
	fmt.Println(file.Stat())
}
```

![File handling Error](https://res.cloudinary.com/techstructive-blog/image/upload/v1663425619/blog-media/golang-file-error.png)

We can use the error as a hint that the file doesn't exist and create a file and then move toward the actual execution of the program. We can also ignore the file object while creating the file, as we are interested in only knowing that the file is just created without any errors, we use the `_` to ignore the variable in the assignment in the function call.

Inside the commented code, we use the [Create](https://pkg.go.dev/os#Create) function to create a file and check for any errors in the process. We finally create the file and `Open` the newly created file.

## Custom Error

We can create custom error types in golang with the help of interfaces and structs. An Error struct will simply consist of a string `message`, that string will display the error. By overriding or creating the `Error` method as an interface for the struct we can construct custom errors.

```go
type Invalid_URL_Error struct {
	message string
}

func (e Invalid_URL_Error) Error() string {
	return "Invalid URL"
}

```

Here, we have the `Invalid_URL_Error` as the custom struct name and the `Error()` method as an interface that will print the error log. This Error method will be used while raising errors in the program. It might be called from another function while doing the actual processing of the URL while sending a GET request.

Further, we can call this custom error method when we wish, by using the package functions, we can override the function call with the custom method.

```go
package main

import (
    "fmt"
    "net/http"
)

type Invalid_URL_Error struct {
	message string
}

func (e Invalid_URL_Error) Error() string {
	return "Invalid URL"
}

func main() {
	url := "htt://meetgor.com"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(Invalid_URL_Error{})
		fmt.Println(err)
	} else {
		fmt.Println(response)
		defer response.Body.Close()
	}
}
```

```
$ go run custom_error.go
Invalid URL 
Get "htt://meetgor.com": unsupported protocol scheme "htt"

```

In the above code, we are basically calling the function `http.Get` that will return a `Response` or an `err` object. We can even call the custom error method with an empty `Invalid_URL_Error` object, this will call the function `Error` from that interface. The function will print the custom error message and thereby we are able to log the custom error message for the invalid URL example.

Also, we can parse the default error method to the custom error method which will get us additional information inside the error interface method.

```go
package main

import (
	"fmt"
	"net/http"
)

type Invalid_URL_Error struct {
	message string
}

func (e Invalid_URL_Error) Error() string {
	return "Invalid URL : " + e.message
}

func main() {
	url := "htt://meetgor.com"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(Invalid_URL_Error{err.Error()})
	} else {
		fmt.Println(response)
		defer response.Body.Close()
	}
}
```

```
$ go run custom_error.go
Invalid URL : Get "htt://meetgor.com": unsupported protocol scheme "htt"

```
When the URL is invalid, we will call the custom error interface by parsing the default `err.Error` method. This will get us the error object from the main method to our custom interface. That is how we will be able to fetch additional information about the error from the interface with the `.` operator as `e.message`. So, the syntax is `Invalid_URL_Error{err.Error()}`, i.e. an object of type `Invalid_URL_Error` with the message set as the value returned from the default `Error()` function. hence we can implement the custom error message.

We also need to look for the response object and close the Response Body as it is mandatory to do so and the responsibility of the caller.

### Creating a function that returns two values (ok,error)

We can even nest the calling of this error method inside another function. This will give us a good overview of how to deal with errors more thoroughly. We will construct a function with two return values one can be any normal desirable object (which we want from the function call) and the other as an error. This will check for any cases that we can call the custom error and return that error interface and the object which was to be returned will be nil if there is an error. If there are no errors, we will return the object and set the error as nil. This way, we can use the `ok, error` syntax while calling this function and thereby make it a lot easier in case of complex programs or multiple types of errors.


```go
package main

import (
	"fmt"
	"net/http"
)

type Invalid_URL_Error struct {
	message string
}

func (e Invalid_URL_Error) Error() string {
	return "Invalid URL"
}

func get_resp(url_link string) (http.Response, error) {

	resp, err := http.Get(url_link)

	if err != nil {
		return http.Response{}, &Invalid_URL_Error{}
	} else {
		defer resp.Body.Close()
		return *resp, nil
	}

}

func main() {

	url := "htts://meetgor.com"
	resp, err := get_resp(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}

}
```

```
$ go run errors.go
Invalid URL
```

This is the simple function `get_resp` which will either fetch the `Response` or `error`. We call the `http.Get` method internally in the function and if something is invalid, it will return an empty Response but the error will be an `Invalid_URL_Error` object which is the custom error class. This means, that if we have an error, we will return a string object from the `Error` method in the interface and if there is no error, we will return the `Response` object and the error will be set as `nil`. Hence, you can now realize, why we check for `err != nil`, it is used for checking if an error object has returned a string or not. As said earlier, we also need to close the request Body.

Further, we can pass the default error method to the custom error class as `Invalid_URL_Error{err.Error()}`. This will ensure, we get additional context from the custom error interface.

```go
package main

import (
	"fmt"
	"net/http"
)

type Invalid_URL_Error struct {
	message string
}

func (e Invalid_URL_Error) Error() string {
	return "Invalid URL : " + e.message
}

func get_resp(url_link string) (http.Response, error) {

	resp, err := http.Get(url_link)

	if err != nil {
		return http.Response{}, &Invalid_URL_Error{err.Error()}
	} else {
		defer resp.Body.Close()
		return *resp, nil
	}

}

func main() {

	url := "htts://meetgor.com"
	resp, err := get_resp(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}

}
```

```
$ go run custom_error
Invalid URL : Get "htts://meetgor.com": unsupported protocol scheme "htts"
```

That's it from this part. Reference for all the code examples and commands can be found in the [100 days of Golang](https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/error-handling/main.go) GitHub repository.

## Conclusion

From this article, we were able to understand the basics of error handling in golang. We can now work with handling errors in function calls and create custom error interfaces. Thank you for reading, if you have any queries, feedback, or questions, you can freely ask me on my social handles. Happy Coding :)
