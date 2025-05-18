{
    "type": "posts",
    "title": "Golang Web: DELETE Method",
    "description": "Exploring the fundamentals of a DELETE method in golang. How to request a resource, parse body, headers, etc. in a HTTP DELETE request.",
    "date": "2025-01-12 20:30:00",
    "status": "published",
    "slug": "golang-web-delete-method",
    "tags": ["go"],
    "series": ["100-days-of-golang"],
    "image_url": "https://meetgor-cdn.pages.dev/100-days-of-golang/golang-037-delete-method.png"
}

## Introduction

We have explored GET, POST, PUT, and PATCH methods in our previous entries for this series. It is the final entry for all the HTTP methods which is going to be the `DELETE` method. In this entry of the series, we will take a view on how to construct and request an HTTP DELETE Method to an API Endpoint.

The DELETE Method is quite simple. You just provide the URI of the resource. Most often, the request body is not needed. That request will simply delete the request entity from the server (in the database or wherever the resource is stored).

Let’s understand it in a more detailed way.

## What is a DELETE Method

The DELETE method requests that the server remove the association between the target resource and its URI (Uniform Resource Identifier). This doesn't necessarily mean the underlying data is physically deleted; it means the resource is no longer accessible through that specific URL. DELETE can also be used to remove relationships between resources, effectively "delinking" them.

According to the RFC:

> The DELETE method requests that the origin server remove the association between the target resource and its current functionality.

Examples:

* **Social Media** (Deleting a Tweet): When you delete a tweet, you're sending a DELETE request to the server. This removes the tweet from your timeline and makes it inaccessible via its URL. While the data might be archived or retained for a period, the key action is removing the public association between the tweet and its online presence. This is closer to a true deletion than the cart example.
    
* **E-Commerce** (Removing an Item from a Cart): When you remove an item from your online shopping cart, you're sending a request (often a DELETE) to remove the item from your cart. The actual product remains available in the store's inventory. This is a clear example of delinking. You're deleting the link between your cart and the product, not the product itself.
    

Let’s start constructing a simple DELETE Request in Golang.

## A Simple DELETE Request

We don’t have a specific method for `DELETE` in `net/http` as we have for `GET` and `POST`, so we need to create a request and use a client to send the request.

### Constructing the URL

We would need to define the endpoint that we are hitting. We can directly use the API URL or construct the API URL on the fly, depending on the ID and dynamic parameters. DELETE requests usually delete a particular entity. We would generally have some form of identifier for that entity/object on the database, etc. So, in this case, it is the user's ID, so we can pass the post to the URL.

```go
// define URL to hit the API
apiURL := "https://reqres.in/api/users/4"

// OR

baseURL := "https://reqres.in"
userID := 4
apiURL := fmt.Sprintf("%s/api/users/%d", baseURL, userID)
```

We can either directly define the URL or dynamically construct the URL, that is quite straightforward. The latter one is the one we usually use and design.

The DELETE Request doesn’t usually require a request body, however, if your server requires some specifications, you can construct the body as we did with the previous examples in POST, PUT, or PATCH method requests.

### Constructing and sending the DELETE Request

We can simply construct the request by specifying the http.MethodDelete as the request method, the URL to hit, and a body(optional) just like a `GET` request. Once we have the request, we can create the default client and send the request.

```go
// create a DELETE request
req, err := http.NewRequest(http.MethodDelete, apiURL, nil)
if err != nil {
	log.Fatal(err)
}

// construct the default http client and send the request
client := &http.Client{}
resp, err := client.Do(req)
if err != nil {
	log.Fatal(err)
}
```

This is the normal code used for constructing an HTTP request in Golang, we create a request using the NewRequest function that takes in the method type, the URL to send the request, and the body if any. Then we need to create a http.Client for sending the request, we usually create a client with default values and send the request using the Do method on the created client using the request that we constructed earlier.

### Fetching the Response

Once the request is sent, we can fetch the response and read the body as bytes, and check the status if that succeeded or failed.

```go
fmt.Println("Response Status:", resp.Status)
respBody, err := io.ReadAll(resp.Body)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Response Body:", string(respBody))
```

We can grab the Status field for checking the status code and message for the request. Usually, the body would be empty since there is no resource we are expecting after deletion of the object. However, if the server is implemented in a way to return the deleted object, you can read the bytes of the body and unmarshal it to the desired struct.

So, that is the entire code to create a simple Delete request with Go, simply construct the URL with the identifier for the resource to be deleted, create the request, and send the request, and if the status code is 204 (usually) then we can assume it succeeded.

```go
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	baseURL := "https://reqres.in"
	userID := 2
	apiURL := fmt.Sprintf("%s/api/users/%d", baseURL, userID)

	req, err := http.NewRequest(http.MethodDelete, apiURL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Response Status:", resp.Status)
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Response Body:", string(respBody))
}
```

## Facts about the DELETE Method

* DELETE Method is idempotent: Similar requests will result in identical behavior or response since once the resource is deleted, the resource won’t exist and hence the behavior would not hinder any other parts.
    
* DELETE Method is not safe: The operation is quite dangerous as it could literally remove a resource from a database/storage. Hence it is called not safe as it is making changes on the server.
    

I have also included some more examples of DELETE requests [**here**](https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/delete/).

That's it from the 36th part of the series, all the source code for the examples are linked in the GitHub on the [**100 days of Golang**](https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/delete/) repository.

[**100-days-of-golang**](https://github.com/Mr-Destructive/100-days-of-golang)

## Conclusion

That would be it from the DELETE Method in Golang. We can use this method just like a normal `GET` request however a bit more carefully.

Hope you found this article, helpful, leave some feedback or any suggestions if you have any. Thank you for reading.

Happy Coding :)
