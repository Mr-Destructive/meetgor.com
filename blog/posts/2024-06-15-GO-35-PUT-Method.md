---
templateKey: blog-post
title: "Golang Web: PUT Method"
description: "Exploring the fundamentals of a PUT method in golang. How to request a resource, parse body, headers, etc. in a HTTP PUT request."
date: 2024-06-15 18:30:00
status: published
slug: golang-web-put-method
tags: ['go',]
series: ['100-days-of-golang',]
image_url: https://meetgor-cdn.pages.dev/100-days-of-golang/golang-035-put-method.png
---

## Introduction

In this section of the series, we will be exploring how to send a `PUT` HTTP request in golang. We will understand how to send a basic PUT request, create an HTTP request, update a resource on a server, parsing the content from struct to json, headers, etc in the following section of this post.

## What is a PUT Method

A PUT method is a type of request that is used to update or modify an entire resource on a server/database.

Imagine you have ordered a pizza at a restaurant and realized you want to change the toppings after you've already placed the order. With a PUT request, it's like informing the waiter about the changes you want to make to your existing order. You specify the updated toppings or any modifications (the data you send). The waiter then takes this updated information (PUT request) back to the kitchen (the server) to apply the changes to your order.

Let's say you created a order.

```nginx
PUT /api/order/456 HTTP/1.1
Host: example.com
Content-Type: application/json
Content-Length: 123

{
    "userID": 123,
    "orderID": 456,
    "items": [
        {
            "itemID": 789,
            "name": "Pizza",
            "quantity": 2,
            "toppings": ["Mushrooms"]
        }
    ]
}
```

In the context of web development, PUT requests are often used for actions such as:

* Updating existing records or resources
    
* Modifying specific parts of an existing resource
    
* Replacing an entire resource with updated data
    

Here's an example of what the PUT request might look like in this scenario:

```nginx
PUT /api/order/456 HTTP/1.1
Host: example.com
Content-Type: application/json
Content-Length: 155

{
    "userID": 123,
    "orderID": 456,
    "items": [
        {
            "itemID": 789,
            "name": "Pizza",
            "quantity": 2,
            "toppings": ["Mushrooms", "Olives"]
        }
    ]
}
```

In this example:

* The PUT method is used to update the resource identified by `/api/order/456`.
    
* The application/json is the content type of the request.
    
* The 155 is the content length of the request.
    
* The body contains the updated details of the order, including the addition of toppings to the pizza.
    

PUT requests are crucial for maintaining and updating data in applications where accuracy and consistency are paramount, ensuring that resources are kept current and reflect the latest changes made by users or systems

## Why the need of PUT Method

In the world of HTTP requests, we use the PUT method to update or modify an entire resource on a server or database. This method is crucial because the POST method, while convenient for creating new data, is not intended for updating existing resources according to standard conventions. While it's possible to misuse the POST method for updates internally, doing so can lead to confusion and inconsistencies in how requests are understood and processed.

## How PUT Method request works

A [PUT](https://www.rfc-editor.org/rfc/rfc9110#PUT) request is utilized to send data to a server for the purpose of updating a resource. When a client (such as a browser or other APIs) sends a PUT request to the server's API endpoint, it includes data in the request body, typically formatted as JSON, XML, or form data.

The server processes the PUT request by first identifying the resource to be updated, either through the URL or data provided in the request body. It then validates, parses, and applies the data from the request body to make modifications to the resource. Following this, the server returns a response that includes a status code indicating the success or failure of the operation. Optionally, the response may also include the updated resource in the response body.

Unlike the POST method, which is primarily used for creating new resources, PUT is specifically designed for updating existing resources on the server. The request body of a PUT contains the data necessary for the update, while the URL identifies the specific resource to be updated.

In summary, PUT requests facilitate the transfer of data to the server specifically for updating resources, ensuring that changes to existing data are accurately processed and reflected.

## Basic PUT Method

To send a `PUT` request to an API in golang, we need to create a `http.Request` object. For `POST` method, the `http` package had the `Post` function defined, however for `PUT` method, there is no separate function. The Go philosophy is right now against adding all the method functions. There have been a couple of discussions on this on [GitHub](https://github.com/golang/go/issues/22841), but it is not been adopted as of 2024.

So, we need to create a `http.Request` object for `PUT` method.

```go
package main

import (
    "fmt"
    "net/http"
)

func main() {
	apiURL := "https://reqres.in/api/users/5"

	req, err := http.NewRequest(http.MethodPut, apiURL, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)

	defer resp.Body.Close()
}
```

```bash
$ go run main.go

200
```

The above code sends a `PUT` request to the [`https://reqres.in/api/users/5`](https://reqres.in/api/users/5) endpoint. The resource we are trying to update is fetched with the identifier `5` which could probably be `id` for the user in the database of the server.

## PUT Method with JSON

Marshaling and encoding are essential in Go for preparing structured data, such as from a struct, into JSON format suitable for HTTP requests like PUT. This conversion ensures data integrity and compatibility between Go types and JSON representations. It's crucial when updating resources on servers, as APIs often require specific data formats for processing updates correctly. Marshaling transforms Go structs into JSON bytes, while encoding further prepares them as request bodies, facilitating seamless communication with web services. This process ensures data consistency and adherence to API specifications, maintaining robust communication in distributed systems.

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Age    string `json:"age"`
	ID     int    `json:"id,omitempty"`
}

type UserResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	user := User{
		Name:   "Alice",
		Salary: 50000,
		Age:    "25",
	}
	apiURL := "https://dummy.restapiexample.com/api/v1/update/11"

	// marshalling process
	// converting Go specific data structure/types to JSON
	bodyBytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyBytes))

	// reading json into a buffer/in-memory
	body := bytes.NewBuffer(bodyBytes)

	// post request
	req, err := http.NewRequest(http.MethodPut, apiURL, body)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 429 {
		fmt.Println("too many requests")
		return
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(respBody))
	defer resp.Body.Close()

	// unmarshalling process
	// converting JSON to Go specific data structure/types
	var userResponse UserResponse
	if err := json.Unmarshal(respBody, &userResponse); err != nil {
		panic(err)
	}
	fmt.Println(userResponse)
}
```

```bash
$ go run json.go                                                                 
{"name":"Alice","salary":50000,"age":"25"}
200                                                                              
{"status":"success","data":[],"message":"Successfully! Record has been updated."}
{success Successfully! Record has been updated.}
```

In the provided Go code example, the `json.Marshal` function is used to convert a Go struct (`User`) into a JSON formatted byte slice (`[]byte`). Here's a breakdown of the steps involved:

* Struct Definition: Define a Go struct with json tags.
    
* Marshalling: Use json.Marshal to convert the struct into JSON byte slice.
    
* Buffer Creation: Wrap the JSON byte slice into an in-memory buffer (bytes.Buffer).
    
* Request Sending: Send a PUT request with the buffer as the request body and set appropriate headers.
    

Let's explore it step by step in detail:

When using the PUT method in Go to update a resource on a server, you often need to send data in JSON format as the request body. Here's how you can achieve this using marshaling and encoding:

1. Define the Struct
    

```go
type User struct {
    Name   string `json:"name"`
    Salary int    `json:"salary"`
    Age    string `json:"age"`
    ID     int    `json:"id,omitempty"`
}
```

Define a Go struct (`User`) that represents the data structure you want to send in JSON format. The json tags specify how each field should be serialized into JSON.

2. Create an Object
    

```go
user := User{
    Name:   "Alice",
    Salary: 50000,
    Age:    "25",
}
```

Create an instance of the User struct (user) with sample data. This data will be marshaled into JSON format to send in the `PUT` request body.

3. Marshal the Struct
    

```go
bodyBytes, err := json.Marshal(user)
if err != nil {
    panic(err)
}
```

Use json.Marshal(user) to convert the user struct into a JSON byte slice (bodyBytes). This byte slice contains the serialized JSON representation of the User struct.

4. Create a Buffer
    

```go
body := bytes.NewBuffer(bodyBytes)
```

Use `bytes.NewBuffer(bodyBytes)` to create an in-memory buffer (`body`) containing the JSON byte slice (`bodyBytes`). The buffer implements the `io.Reader` interface needed for the PUT request body.

5. Create a PUT Request
    

```go
req, err := http.NewRequest(http.MethodPut, apiURL, body)
if err != nil {
    panic(err)
}
```

Use http.NewRequest to create a new PUT request to the specified URL with the JSON buffer (`body`) as the request body. Set appropriate headers if needed (e.g., Content-Type as application/json).

6. Send the Request
    

```go
resp, err := http.DefaultClient.Do(req)
if err != nil {
    panic(err)
}
```

Use [`http.DefaultClient.Do`](http://http.DefaultClient.Do)`(req)` to execute the PUT request and obtain the response. Handle any errors that may occur during the request execution.

7. Process the Response
    

```go
respBody, err := io.ReadAll(resp.Body)
if err != nil {
    panic(err)
}
```

Use `io.ReadAll(resp.Body)` to read and store the response body from the server. Handle any errors encountered during the reading process.

8. Unmarshal the Response
    

```go
var userResponse UserResponse
if err := json.Unmarshal(respBody, &userResponse); err != nil {
    panic(err)
}
```

Use `json.Unmarshal(respBody, &userResponse)` to deserialize the JSON response body into a Go struct. This allows you to work with the response data in a structured manner.

The parsing of files and form data is also possible with `PUT` requests, however, that has been covered in the [POST Method](https://meetgor.com/golang-web-post-method). Those snippets would be handy in these request method as well.

I have also included some more examples of PUT requests [here](https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/put/).

That's it from the 35th part of the series, all the source code for the examples are linked in the GitHub on the [100 days of Golang](https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/put/) repository.

[100-days-of-golang](https://github.com/Mr-Destructive/100-days-of-golang)

## Conclusion

That's it from this post of the series, a post on the PUT method in golang :)

We have covered topics like creating basic PUT requests and marshaling golang types into JSON format. Hope you found this article helpful. If you have any queries, questions, or feedback, please let me know in the comments or on my social handles. Thank you for reading.

Happy Coding :)
