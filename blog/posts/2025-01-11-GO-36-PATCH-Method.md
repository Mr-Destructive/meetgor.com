---
templateKey: blog-post
title: "Golang Web: PATCH Method"
description: "Exploring the fundamentals of a PATCH method in golang. How to request a resource, parse body, headers, etc. in a HTTP PATCH request."
date: 2025-01-11 23:30:00
status: published
slug: golang-web-patch-method
tags: ['go',]
series: ['100-days-of-golang',]
image_url: https://meetgor-cdn.pages.dev/100-days-of-golang/golang-036-patch-method.png
---


## Introduction

In previous sections of this series, we've covered the GET, POST, and PUT methods. Now, we will explore the PATCH method, which differs from the others in several key ways. The PATCH method is somewhat more flexible and depends on how the server or API you're working with is designed.

In this section, we'll focus on understanding what the PATCH method is and how to use it. While we will dive deeper into building and structuring a full CRUD API later in the series, the focus here will be on the what and why of the PATCH method, not the how.

## What is the PATCH Method?

The PATCH method is often compared to the PUT method, but with one important distinction: PATCH is used to perform partial updates on a resource. Unlike PUT, which typically requires you to send the entire resource to update it, PATCH allows you to send only the fields that need to be updated. This makes it a more efficient option when updating a subset of a resource.

In a PATCH request, the body usually contains instructions in a format like JSON, which specifies the fields to update. These instructions define the changes to be applied to the resource. For example, you may only want to change one field of a user's profile, such as their email address, while leaving the rest of the data untouched.

## PATCH vs. PUT

Key Differences While both PATCH and PUT are used to modify resources, there are significant differences in their behavior:

* PUT replaces the entire resource. When you send a PUT request, you must include the full representation of the resource, even if you're only changing a small part of it.
    
* PATCH, on the other hand, is for partial updates. You only need to include the fields that are changing, not the entire resource.
    

If the update involves more fields than just the ones you're changing, PUT may be the better choice. However, the scope of this article is to focus solely on the PATCH method.

## How Does a PATCH Request Work?

In the simplest terms, a PATCH request allows you to perform a partial update on a resource. It is similar to a PUT request, but more specific in how it updates the resource. According to the HTTP specification, an ideal PATCH request should:

* Accept a "patch document" in the request body, which contains the list of operations to apply (e.g., "replace", "add", "remove").
    
* Apply these updates to the target resource.
    
* If the update cannot be applied correctly, the operation should fail without applying any of the changes.
    

This ensures that no partial or inconsistent updates are left behind

For example, if you're updating a user's email address and something goes wrong in the middle of the operation, the PATCH request should ensure that the email isn't updated partially. If there’s an error, none of the updates should be applied, ensuring data consistency.

Also, the patch method is not idempotent, meaning that if you send the same input/request, it need not necessarily return the same output. Because we are not sending the actual original entity, we are sending the partial data fields that need to be updated, so it might update the original entity on subsequent requests since there is no original request sent in the request; it only identifies the resource from the URI and fields to update in the request body.

Now, let’s sum up the patch request in a few words

* Updates specific fields mentioned in the patch document
    
* Can be partial (only the fields that need to be updated are sent, unlike PUT, which typically replaces the entire resource)
    
* Not necessarily idempotent (depends on the implementation)
    
* Not Safe (since resources will be updated on the server side)
    

## Basic PATCH request

Let’s start with the basic PATCH request that we can create in Golang. The [net/http](https://pkg.go.dev/net/http) package will be used to construct the request and will be using `encoding/json` and some other utilities for string and byte parsing.

So, first, we will construct an HTTP request using the [http.NewRequest](https://pkg.go.dev/net/http#NewRequest) with the parameters like the http method to use, the URL to hit, and the request body if any. We will then need to send the json body which would consist of the fields to be updated.

### Defining the API/Server Endpoint URL

We would need to define the endpoint that we are hitting, we can directly use the API URL or we can construct the API URL on the fly depending on the ID, and parameter that will be dynamic. As PATCH requests, usually modify a particular entity, we would generally have some form of identifier for that entity/object on the database, etc. So in this case, it is `id` of the post, so, we can pass the post in the URL.

```go
// define URL to hit the API
apiURL := "https://jsonplaceholder.typicode.com/posts/4"

// OR
// baseURL := "https://jsonplaceholder.typicode.com"
// postId := 4
// postURL := fmt.Sprintf("%s/posts/%d", baseURL, postId)
```

We can either directly define the URL or dynamically construct the URL, that is quite straightforward. The latter one is the one we usually use and design.

### Constructing the JSON Body

This section is a little dependent on the context as you might have a direct json string that you can directly pass to the API or you might have a golang object that you need to Marshal in order to convert that object into string/bytes.

1. Direct JSON String
    
    So, there is nothing to do here, since the object is already in the form of a json string.
    
    ```go
    reqBody := `{"body": "new body"}`
    ```
    
    However, if you have certain fields that you need to exclude or omit, you have to construct a struct and then marshal it
    
2. Marshaling (converting object into bytes/string)
    
    We need to convert the Golang native object into some form of a json string or bytes that could be sent over the network. That process is called [marshaling](https://en.wikipedia.org/wiki/Marshalling_\(computer_science\)) or serialization.
    

```go
type Post struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
	UserId int    `json:"userId,omitempty"`
}

userObj := Post{
    Body: "New Body",
}

var reqBody []byte
reqBody, err := json.Marshal(userObj)
if err != nil {
	log.Fatal(err)
}

log.Println("New body:", string(reqBody))
// New body: {"body":"New Body"}
```

In the above snippet, we have defined a `Post` struct with the fields like `ID`, `Title`, `Body`, `UserID` , and those who have `omitempty` tag along with the json fields that we want to marshal into. The omit empty will omit or ignore the fields if they are empty or not present in the object/instance of this structure. So in the example, `userObj` is the instance of the `Post` struct and it only has the `Body` populated, so the reqBody will only contain one field `body` in the json representation. The [json.Marshal](https://pkg.go.dev/encoding/json#Marshal) is the function that we use to convert the object/instance into a byte form.

This `reqBody` will serve as the request body for the request that will be a `PATCH` method to the mentioned endpoint / API URL.

### Constructing the HTTP PATCH Request

Now, we have the parts that we need to construct the request, we can combine the parts and hit the endpoint. However, it is a bit different compared to `GET` and `POST` request that we do in Golang, the HTTP package has built-in methods for the `GET` and `POST` methods, however for methods like `PUT`, `PATCH`, `DELETE` and others, we need to construct a [Request](https://pkg.go.dev/net/http#Request) object and then send that request.

```go
req, err := http.NewRequest("PATCH", postURL, strings.NewReader(reqBody))
if err != nil {
	log.Fatal(err)
}
req.Header.Set("Content-Type", "application/json")

// in case of wired utf-8 characters appear in the body
//req.Header.Set("Content-Type", "application/json; charset=utf-8")
```

To do that, we call the [NewRequest](https://pkg.go.dev/net/http#NewRequest) method with the parameters like the HTTP method, the URL, and the request Body all of which we have at the moment.

* The method is `PATCH`
    
* The URL is `postURL`
    
* The body is `strings.NewReader(reqBody)` as we need a `io.Reader` object instead of string or byte slice
    

So, once we have that, we will also set the `Header` In the field of `Content-Type` and the value as `application/json` since the request body has json representation of the patch document that will be sent.

### Sending the Request

Once, the `req` the object is created, we also need a [Client](https://pkg.go.dev/net/http#Client) to send the request, so we create the client as default http.Client object with defaults and call the [Do](https://pkg.go.dev/net/http#Client.Do) method with the `req` as the request parameter in order to send the request with the default client.

This method returns the response object and an error if any.

We also add the `defer resp.Body.Close()` in order to avoid leaks and safely access the response body.

```go
client := &http.Client{}
resp, err := client.Do(req)
if err != nil {
	log.Fatal(err)
}
defer resp.Body.Close()
```

At this point, we can now start consuming the response and use it for further processing as per the needs.

### Unmarshalling the Response

We first read the response into a string or byte representation using the io.ReadAll method and then use the json.Unmarshal to convert the bytes into golang object/instance.

```go
var updatedPost Post

respBody, err := io.ReadAll(resp.Body)
if err != nil {
	log.Fatal(err)
}
    
// convert the response json bytes to Post object in golang
err = json.Unmarshal(respBody, &updatedPost)
if err != nil {
	log.Fatal(err)
}

fmt.Println(updatedPost)
fmt.Println(updatedPost.Title)
```

In the above example, we have read the response Body which can be accessed as the `Body` field in the [Response](https://pkg.go.dev/net/http#Response) object via the `resp` variable. So, the function will return the `respBody` as a string or an error if any. Then using this string object, we can use the [json.Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) function to send this string and create this `updatedPost` object of Post struct. This method will mutate this object as we have passed it by reference indicated by `&updatedPost` . So, this will do two things, one update / mutate the `updatedPost` instance from the `respBody` and give any error if any arises during the [deserialization](https://developer.mozilla.org/en-US/docs/Glossary/Deserialization) of the response.

Now, we have the object in golang from the response bytes, we can use it as per requirements.

So, that is the example in the entirety.

Let’s simplify the steps which are similar to the POST/PUT method as well

* Define/construct URL
    
* Marshal the object into JSON string as the request body
    
* Construct the request object (method, URL, and the body)
    
* Send the request with the default client
    
* Read the response and deserialize/unmarshal
    
* Access the object in golang
    

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Post struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
	UserId int    `json:"userId,omitempty"`
}

func main() {

	// define URL to hit the API
    baseURL := "https://jsonplaceholder.typicode.com"
    postId := 4
    postURL := fmt.Sprintf("%s/posts/%d", baseURL, postId)

    // define the body -> with the field to update
	reqBody := `{"body": "new body"}`
	fmt.Println("New body:", reqBody)

    // send a new request, with the `PATCH` method, url and the body
	req, err := http.NewRequest("PATCH", postURL, strings.NewReader(reqBody))
	if err != nil {
		log.Fatal(err)
	}
    // set the header content type to json
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status code:", resp.StatusCode)
	fmt.Println("Response Status:", resp.Status)
	
    var updatedPost Post

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
    
    // convert the response json bytes to Post object in golang
	err = json.Unmarshal(respBody, &updatedPost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updatedPost)
	fmt.Println(updatedPost.Title)

}
```

```plaintext
New body: {"body": "new body"}
Response status code: 200
Response Status: 200 OK
{4 eum et est occaecati new body 1}

title: eum et est occaecati                                                                         
body: new body
id: 4
user id: 1
```

As you can see, it has only updated the `body` and has not updated the other fields.

If you would have sent a similar body with a `PUT` method, the results would have been different. That would have been dependent on the implementation of the API of course, but if there is only a few fields in the request body for a PUT method, it would have replaced the values with the empry values which are not present in the request body.

That is the difference between a `PUT` and a `PATCH` method, the `PATCH` method, ideally should only update the fields of the entity that are mentioned in the request body, whereas the `PUT` method has to update the entire resource whether the fields are provided or not. Again, the implementation of these API on the server plays a vital role in how the behavior defers and the method in itself would perform.

This is also called as `JSON Merge Patch`

## JSON Merge PATCH

The above API is implementing a [Merge PATCH](https://datatracker.ietf.org/doc/html/rfc7386) which is to say, merge the changes in the actual entity.

Let’s say there is a Blog post Entity on a Server, you have a post that you are writing as an author. The post has an id of `4` let’s say and you are constantly changing the body of the post.

So, you don’t want to send the `title` or `author_id` or any field that is not changing from the client again and again while saving, so the `MERGE PATCH` endpoint will be helpful in that case, where the client only sends the required fields to be updated.

In this example, the user would only send the `body` of the post to the API every time it makes changes or saves the draft. In some cases, it might also want to change the title, so it will include the title, but not all the fields. The API will know it is a `PATCH` request and the content type is `json` so it will only change or update the fields that are provided in the request body to the actual entity in the database or wherever it is stored on the server.

So, that is what is the JSON Merge PATCH or Merge PATCH in general. JSON Merge PATCH is specific to the JSON-based document APIs.

Below is the example, the same steps but a different endpoint. A user API that I have specifically created to demonstrate the difference in PUT vs Merge PATCH vs JSON PATCH requests.

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type User struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Roles string `json:"roles,omitempty"`
}

func main() {
    baseURL := "https://dummy-json-patch.netlify.app/.netlify/functions"
    userID := 2
	apiURL := fmt.Sprintf("%s/users/?id=%d", baseURL, userID)
    
    userObj := User{
		Name:  "dummy name",
		Roles: "dummy role",
	}

	var jsonPatchBody []byte
	jsonPatchBody, err := json.Marshal(userObj)
	if err != nil {
		panic(err)
	}
    
    // OR directly define the json as string
	//jsonMergePatchBody := `{
    //    "name": "new dummy name",
    //    "roles": "new dummy role"
    //}`

	req, err := http.NewRequest("PATCH", apiURL, strings.NewReader(jsonMergePatchBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var user User

	resBody, err := io.ReadAll(resp.Body)
	fmt.Println(string(resBody))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated/Patched User", user)
	fmt.Println("Id:", user.ID)
	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)
    fmt.Println("Bio:", user.Bio)
	fmt.Println("Roles:", user.Roles)
}
```

Original User with ID 2

```plaintext
{"id":2,"name":"dummy","email":"dummyyummy@user.com","bio":"empty bio","roles":"read"}

Id: 2
Name: dummy
Email: dummyyummy@user.com
Bio: empty bio
Roles: read
```

Output of the program

```plaintext
Request Body: {"name":"dummy name","roles":"dummy role"}

{"id":2,"name":"dummy name","email":"dummyyummy@user.com","bio":"empty bio","roles":"dummy role"}

Updated/Patched User {2 dummy name dummyyummy@user.com empty bio dummy role}

Id: 2
Name: dummy name
Email: dummyyummy@user.com
Bio: empty bio
Roles: dummy role
```

In the above example, the only fields that will be updated are `name` and `roles` , since the API is implemented to only update the fields provided in the json merge patch document (request body).

As you can see that, only the `name` and `roles` are changed. The name was `dummy` that changed to `dummy name` and role changed from `read` to `dummy role` .

Now, let’s see the same request but with the PUT method on it.

Before we hit this API however, let’s note what the user with ID 2 is

```json
{
  "id": 2,
  "name": "dummy name",
  "email": "dummyyummy@user.com",
  "bio": "empty bio",
  "roles": "dummy role"
}
```

This is the result of our recent patch request. Now, we will send a PUT request to the same user with a different body.

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type User struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Bio   string `json:"bio,omitempty"`
	Roles string `json:"roles,omitempt"`
}

func main() {
    baseURL := "https://dummy-json-patch.netlify.app/.netlify/functions"
    userID := 2
	apiURL := fmt.Sprintf("%s/users/?id=%d", baseURL, userID)
    
	userObj := User{
		Name:  "not a dummy name",
		Roles: "not a dummy role",
	}

	var jsonPatchBody []byte
	jsonPatchBody, err := json.Marshal(userObj)
	if err != nil {
		panic(err)
	}
	fmt.Println("Request Body:", string(jsonPatchBody))

	//jsonPatchBody := `{
	//    "name": "dummy",
	//    "roles": "new dummy role"
	//}`

	req, err := http.NewRequest("PUT", apiURL, strings.NewReader(string(jsonPatchBody)))
	req.Header.Set("Content-Type", "application/merge-patch+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var user User

	resBody, err := io.ReadAll(resp.Body)
	fmt.Println(string(resBody))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated/Patched User", user)
	fmt.Println("Id:", user.ID)
	fmt.Println("Name:", user.Name)
	fmt.Println("Bio:", user.Bio)
	fmt.Println("Email:", user.Email)
	fmt.Println("Roles:", user.Roles)
}
```

Output:

```plaintext
Request Body: {"name":"not a dummy name","roles":"not a dummy role"}

{"id":2,"name":"not a dummy name","email":"","bio":"","roles":"not a dummy role"}

Updated/Patched User {2 not a dummy name   not a dummy role}

Id: 2
Name: not a dummy name
Bio:
Email:
Roles: not a dummy role
```

As you can see the `name` and `roles` are updated, however, the `bio` and `email` fields are empty. Since we only said to update the `name` and `roles` fields, but it was a `PUT` request, it expects all the fields, and if any of the fields are missing, it will consider them as empty and update those as well.

So, the difference might be crystal clear now. When to use `PATCH` and when to avoid `PUT`.

* When you have a large set of updates, preference could be given to PUT
    
* If you have very specific fields to update and minimal fields PATCH is recommended
    

There is another type of PATCH specifically designed for JSON APIs or JSON Documents APIs.

## JSON PATCH

The [JSON PATCH](https://datatracker.ietf.org/doc/html/rfc6902/) is a specification in which we can define what operations to perform on which fields, or path of the fields to replace, move, or copy to.

> A JSON Patch document is a JSON document that represents an array of objects. Each object represents a single operation to be applied to the target JSON document.

As it takes these operations, it applies them sequentially and hence it won’t replace all the fields for the entity, as that is the expected behavior of the PATCH method. In other words, it would only apply changes to the fields and the related fields provided in the json patch document (request body).

There are a few operations that you can perform with this json patch method, and provide the instructions accordingly for individual operations in the JSON PATCH document.

Operations

1. add
    
2. remove
    
3. replace
    
4. move
    
5. copy
    
6. test
    

So, for each of the operations, a high-level definition can be considered as:

* To add a field you can specify the operation as `add` , the path as the field to be added, and the value as the actual value to be added
    
* To remove a field, you can specify the operation as `remove` , and the path as the field to remove
    
* To replace a field, you can specify the operation as `replace`, the path as the field to be updated/replaced, and the value of the actual value to be added
    
* To move a field, you can specify the operation as `move`, the **form** as the field to be updated/moved from and the path to the field the from value should be moved to.
    
* To copy a field, you can specify the operation as `copy`, the form as the field to update/copied from, and the path to the field to which the value should be copied to.
    
* The test operation is a bit different as it is used for comparison of a `path` value to the actual value specified in the object. It might return true or false, but not actually return it might be used as a checkpoint for continuing with the operation in the document.
    

In this example, we are creating a similar patch request, but using this json patch document kind of structure.

### Construct the json-patch document

We need to construct the JSON PATCH document, which is a list of operations (instructions). Each instruction has an operation `op` field that defines what operation need to be performed among the available 6 operations. The `path` field is the field that we need to update, it is called path because, we need to provide the path to the json field, in case of nested json fields. The `value` and `from` fields are a bit dependent on the operation as `value` is not required in `remove` operation and `from` is only used in `move` and `copy` operations.

```go
jsonPatchBody := `[
    {
        "op": "replace",
        "path": "/name",
        "value": "new dummy name"
    },
    {
        "op": "replace",
        "path": "/roles",
        "value": "new dummy role"
    },
]`


// OR
// construct a go type and then marshal it

type JSONPatch struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	From  string `json:"from,omitempty"`
	Value string `json:"value,omitempty"`
}

jsonPatchData := []JSONPatch{
	{
		Op:    "replace",
		Path:  "/name",
		Value: "new dummy name",
	},
	{
		Op:    "replace",
		Path:  "/roles",
		Value: "new dummy role",
	},
}

jsonPatchBytes, err := json.Marshal(jsonPatchData)
if err != nil {
	fmt.Println("Error marshalling JSON:", err)
	return
}

jsonPatchBody := string(jsonPatchBytes)
```

In the above example, we have defined the two ways to create the body for a json patch request. One is directly writing json string (not recommended). Also, the other way is to define a struct as JSONPatch and construct the object for the operations in the json patch request. The JSONPatch object has the 4 fields as mentioned earlier. In this case, we are only using the `replace` operation, which will update the value in the provided path field to the value provided in the value field.

In this specific example, we have 2 operations to perform, first to update the `name`field and the second to update the `roles` field. We simply use the `json.Marshal` method to convert the object into bytes. We then cast the slice of bytes into string and that becomes a final `jsonPatchBody` for the request that we will send.

**<mark>NOTE: This API </mark>** `https://dummy-json-patch.netlify.app/.netlify/functions/users/` **<mark> is created by me, it only implements the </mark>** `replace` **<mark> operation, as it makes sense to only have an update operation on the user table endpoint, If the model is more like a json document-like structure, the other operations might make sense. I might implement those in a separate article, while we explore the implementation of the PATCH Request in the CRUD API Section of this series soon.</mark>**

### Constructing and Sending the Request

We do the same stuff as the normal merge patch request, that is provide the `PATCH` as the request method, the `apiURL` and the `strings.NewReader(jsonPatchBody)` into the body of the request.

However, the only difference is that we need to set the `Content-Type` header as `application/json-patch+json` . This is defined in the [specification](https://datatracker.ietf.org/doc/html/rfc6902/#section-1), it would make the client as well as the server understand the request body. This can also be used by the server to differentiate between a standard patch (which can be referred to as a "Merge PATCH" for JSON APIs) and a JSON PATCH request, as shown here.

```go
req, err := http.NewRequest("PATCH", apiURL, strings.NewReader(jsonPatchBody))
req.Header.Set("Content-Type", "application/json-patch+json")
```

Once, we have created the request object we then simply can create the client and send the request.

```go
client := &http.Client{}
resp, err := client.Do(req)
if err != nil {
	panic(err)
}
defer resp.Body.Close()

var updatedUser User

resBody, err := io.ReadAll(resp.Body)
fmt.Println(string(resBody))
if err != nil {
	panic(err)
}

err = json.Unmarshal(resBody, &updatedUser)
if err != nil {
	panic(err)
}

fmt.Println("Updated/Patched User", updatedUser)
```

Once, we have the response from the request, we just do the normal response processing and unmarshal the response string into a golang object type (passing by reference) which is provided as `&updatedUser` in the parameter.

This will give us the expected result back in the updated User object.

The below code is the entire JSON PATCH request:

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
    Bio   string `json:"bio"`
	Roles string `json:"roles"`
}

func main() {
	baseURL := "https://dummy-json-patch.netlify.app/.netlify/functions"
    userID := 2
	apiURL := fmt.Sprintf("%s/users/?id=%d", baseURL, userID)

	jsonPatchBody := `[
        {
            "op": "replace",
            "path": "/name",
            "value": "new dummy name"
        },
        {
            "op": "replace",
            "path": "/roles",
            "value": "new dummy role"
        },
    ]`

	req, err := http.NewRequest("PATCH", apiURL, strings.NewReader(jsonPatchBody))
	req.Header.Set("Content-Type", "application/json-patch+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var user User

	resBody, err := io.ReadAll(resp.Body)
	fmt.Println(string(resBody))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated/Patched User", user)
	fmt.Println("Id:", user.ID)
	fmt.Println("Name:", user.Name)
	fmt.Println("Bio:", user.Bio)
	fmt.Println("Email:", user.Email)
	fmt.Println("Roles:", user.Roles)
}
```

Original User (id=2) before the request

```json
{
  "id": 2,
  "name": "dummy name",
  "email": "dummyyummy@user.com",
  "bio": "empty bio",
  "roles": "dummy role"
}
```

Output of the Program (JSON PATCH Request)

```plaintext
{"id":2,"name":"new dummy name","email":"dummyyummy@user.com","bio": "empty bio", "roles":"new dummy role"}     

Updated/Patched User {2 new dummy name dummyyummy@user.com empty bio new dummy role} 

Id: 2
Name: new dummy name
Email: dummyyummy@user.com
Bio: empty bio
Roles: read
```

So, as expected, it only updates the `name` and the `roles` field in the user object with id = 2. The patch document has the replace operation that updates the name and the roles fields from the existing field.

To add more examples for `add` , `move`,`copy`, `remove` and `test` operations.

PS: I added an endpoint for documents, which will include all the crud operations with all the JSON PATCH operations added as well.

Let’s take a simple example for each operation

### Add operation

Let’s take an example for a document table at the `https://dummy-json-patch.netlify.app/.netlify/functions/documents` endpoint, the document is a table with the columns like data and the id as the index.

```go
{
  "data": {},
  "id": 4
}
```

The initial data is empty for a document with an ID is 4.

We simply use the operation of `add` that will add the value of the `name` in the `data` object with the field name as `name` and the value as `dummy` .

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Document struct {
	ID   int         `json:"id"`
	Data interface{} `json:"data"`
}

func main() {
	apiURL := "https://dummy-json-patch.netlify.app/.netlify/functions/documents/?id=4"

	jsonPatchBody := `[
        {
            "op": "add",
            "path": "/name",
            "value": "dummy"
        }
    ]`

	req, err := http.NewRequest("PATCH", apiURL, strings.NewReader(jsonPatchBody))
	req.Header.Set("Content-Type", "application/json-patch+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var updatedDoc Document

	resBody, err := io.ReadAll(resp.Body)
	fmt.Println(string(resBody))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(resBody, &updatedDoc)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated/Patched Document", updatedDoc)
	fmt.Println("id:", updatedDoc.ID)
	fmt.Println("document:", updatedDoc.Data)

}
```

This is a simple add operation, as you would see that the data now has the name field populated after the request was sent.

```go
{"data":{"name":"dummy"},"id":4}
Updated/Patched Document {4 map[name:dummy]}
id: 4
document: map[name:dummy] 
```

### Remove Operation

We can take an example for remove, let’s remove the `name` in the `data` field that we added previously.

The data before the request looks like this:

```go
{
    "data": {
        "name": "dummy"
    },
    "id": 4
}
```

Now, we will construct the request, we can do it with the json marshaling by constructing the object, but I am keeping it simple and shorter by directly constructing the json body.

For the `remove` operation, we only need the `path` field to be added, in order to remove the field from the object.

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Document struct {
	ID   int         `json:"id"`
	Data interface{} `json:"data"`
}

func main() {
	apiURL := "https://dummy-json-patch.netlify.app/.netlify/functions/documents/?id=4"

	jsonPatchBody := `[
        {
            "op": "remove",
            "path": "/name"
        }
    ]`

	req, err := http.NewRequest("PATCH", apiURL, strings.NewReader(jsonPatchBody))
	req.Header.Set("Content-Type", "application/json-patch+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var updatedDoc Document

	resBody, err := io.ReadAll(resp.Body)
	fmt.Println(string(resBody))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(resBody, &updatedDoc)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated/Patched Document", updatedDoc)
	fmt.Println("id:", updatedDoc.ID)
	fmt.Println("document:", updatedDoc.Data)

}
```

```plaintext
{"data":{},"id":4}                

Updated/Patched Document {4 map[]}
id: 4                             
document: map[]                   
```

So, as you can see the updated data has no fields.

### Copy Operation

Now, we can also perform the copy operation, which can copy the field to another field, this instruction requires the `from` and the `path` . The `from` is the path to the field to copy it from, and the `path` is the path to the field to copy the field value.

```json
[
    {
        "op": "copy",
        "from": "/name",
        "path": "/new_name"
    }
]
```

The above json path operation is for copying the value of the field `name` to the field `new_name` i.e. if the original document looks likes this:

```json
{
    "data": {
        "name": "dummy",
        "new_name": "hello"
    },
    "id": 4
}
```

So, after this the value of `new_name` will be `dummy` after the operation

```json
 {
    "data": {
        "name": "dummy",
        "new_name": "dummy"
    },
    "id": 4
}
```

So, let’s take a look at the code where we will perform this operation with this API endpoint and constructing the json body.

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Document struct {
	ID   int         `json:"id"`
	Data interface{} `json:"data"`
}

func main() {
	apiURL := "https://dummy-json-patch.netlify.app/.netlify/functions/documents/?id=4"

	jsonPatchBody := `[
        {
            "op": "copy",
            "from": "/name",
            "path": "/new_name"
        }
    ]`

	req, err := http.NewRequest("PATCH", apiURL, strings.NewReader(jsonPatchBody))
	req.Header.Set("Content-Type", "application/json-patch+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var updatedDoc Document

	resBody, err := io.ReadAll(resp.Body)
	fmt.Println(string(resBody))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(resBody, &updatedDoc)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated/Patched Document", updatedDoc)
	fmt.Println("id:", updatedDoc.ID)
	fmt.Println("document:", updatedDoc.Data)

}
```

In the above example, we are creating the json payload for copying the value of `name` to the `new_name` field.

```plaintext
{"data":{"name":"dummy","new_name":"dummy"},"id":4}
Updated/Patched Document {4 map[name:dummy new_name:dummy]}
id: 4
document: map[name:dummy new_name:dummy]
```

### Move Operation

Now, we can also perform the move operation, which can move the field to another field, this instruction requires the `from` and the `path` . The `from` is the path to the field to move it from, and the `path` is the path to the field to move.

```json
[
    {
        "op": "move",
        "from": "/name",
        "path": "/user/description"
    }
]
```

The above json path operation is for moving the value of the field `name` to the field `/user/description` i.e. if the original document looks like this:

```json
 {
    "name": "dummy",
    "new_name": "dummy",
    "user": {
        "description": "hello"
    }
}
```

So, after the operation, the value of `name` will be `dummy`. The name field is moved not replaced, that is the difference of this operation.

```json
 {
    "new_name": "dummy",
    "user": {
        "description": "dummy"
    }
}
```

So, let’s take a look at the code where we will construct the payload body and perform the json patch request with the move operations.

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Document struct {
	ID   int         `json:"id"`
	Data interface{} `json:"data"`
}

func main() {
	apiURL := "https://dummy-json-patch.netlify.app/.netlify/functions/documents/?id=4"

	jsonPatchBody := `[
        {
            "op": "move",
            "from": "/name",
            "path": "/user/description"
        }
    ]`

	req, err := http.NewRequest("PATCH", apiURL, strings.NewReader(jsonPatchBody))
	req.Header.Set("Content-Type", "application/json-patch+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var updatedDoc Document

	resBody, err := io.ReadAll(resp.Body)
	fmt.Println(string(resBody))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(resBody, &updatedDoc)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated/Patched Document", updatedDoc)
	fmt.Println("id:", updatedDoc.ID)
	fmt.Println("document:", updatedDoc.Data)

}
```

The above code simply takes in the json string as a payload for the API with the move operation. Hence the name field’s value is swapped to the value of `dummy` of the `/user/description` field.

```go
{"data":{"new_name":"dummy","user":{"description":"dummy"}},"id":4}
Updated/Patched Document {4 map[new_name:dummy user:map[description:dummy]]}
id: 4
document: map[new_name:dummy user:map[description:dummy]]
```

### Test Operation

The test operation requires the `path` field with the value of the field to test, and the `value` field as the value to compare with. This is a test, so we can use it as a flag to proceed ahead with the rest of the operations or abort.

```json
[
    {
        "op": "test",
        "path": "/new_name",
        "value": "testing value"
    },
    {
        "op": "add",
        "path": "/user/id",
        "value": 123
    }
]
```

if the original document looks like this:

```json
 {
    "new_name": "dummy",
    "user": {
        "description": "dummy"
    }
}
```

Then after applying the test operation, we can find that the `new_name` has the value is `dummy` and not `testing value` so it will abort with an error.

```json
{
    "error":"test failed: values do not match at path new_name"
}
```

However, if we keep the operation like this:

So, we compare the value of `new_name` field is `dummy` which it is then we will perform the rest of the operations in this case, the `add` operations will be performed to add the `id` field to the `user` object with the value `123`.

```json
[
    {
        "op": "test",
        "path": "/new_name",
        "value": "dummy"
    },
    {
        "op": "add",
        "path": "/user/id",
        "value": 123
    }
]
```

if the original document looks like this:

```json
{
    "new_name": "dummy",
    "user": {
        "description": "dummy"
    }
}
```

Then after the operation of test and add, the document becomes:

```json
{
    "new_name": "dummy",
    "user": {
        "description": "dummy",
        "id": 123
    }
}
```

You can observe the `id` field is added to the `user` field, as the test was successful. Hence, the operation `add` was performed and the request was completed.

Let’s take a look at the code for doing this operation:

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Document struct {
	ID   int         `json:"id"`
	Data interface{} `json:"data"`
}

func main() {
	apiURL := "https://dummy-json-patch.netlify.app/.netlify/functions/documents/?id=4"

    /*
    jsonPatchBody := `[
        {
            "op": "test",
            "path": "/new_name",
            "value": "testing value"
        },
        {
            "op": "add",
            "path": "/user/id",
            "value": 123
        }
    ]`
    */
	jsonPatchBody := `[
        {
            "op": "test",
            "path": "/new_name",
            "value": "dummy"
        },
        {
            "op": "add",
            "path": "/user/id",
            "value": 123
        }
    ]`

	req, err := http.NewRequest("PATCH", apiURL, strings.NewReader(jsonPatchBody))
	req.Header.Set("Content-Type", "application/json-patch+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var updatedDoc Document

	resBody, err := io.ReadAll(resp.Body)
	fmt.Println(string(resBody))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(resBody, &updatedDoc)
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated/Patched Document", updatedDoc)
	fmt.Println("id:", updatedDoc.ID)
	fmt.Println("document:", updatedDoc.Data)

}
```

So, the above code is an example of performing the test operation.

```json
{"data":{"new_name":"dummy","user":{"description":"dummy","id":123}},"id":4}                                                                    
Updated/Patched Document {4 map[new_name:dummy user:map[description:dummy id:123]]}                                                             
id: 4                                                                                                                                           
document: map[new_name:dummy user:map[description:dummy id:123]]
```

So, as you would see, the id is indeed added to the document.

Here’s an overview of PUT, PATCH, MERGE PATCH, and JSON PATCH requests.

|  | PUT | PATCH (General) | Merge Patch (JSON) | JSON Patch |
| --- | --- | --- | --- | --- |
| **Purpose** | Replace the entire resource. | Partially modify a resource. | Partially modify a JSON resource by merging changes. | Partially modify a JSON resource using explicit operations. |
| **Request Body** | Full representation of the resource. | Instructions for modification. | JSON object containing fields to be updated. | JSON array of operation objects (e.g., `add`, `remove`, `replace`). |
| **Idempotency** | Idempotent (repeated requests have the same effect). | Not necessarily idempotent (depends on implementation). | Idempotent (repeated requests have the same effect). | Idempotent (if operations are idempotent, like `replace`). |
| Safe | No | No | No | No |
| **Data Sent** | All resource fields. | Only the fields/instructions to be modified. | Only the fields to be updated. | Instructions specifying how to modify the resource (path, operation, value). |
| **Server Behavior** | Replaces the entire resource. | Applies the provided modifications. | Merges the provided fields into the existing resource. | Applies the operations in the order they appear in the array. |
| **Content-Type** | `application/json` (usually) | Varies depending on the patch format. | `application/merge-patch+json` | `application/json-patch+json` |

So, that is all from the PATCH method in GOlang, I might have explained a bit too much for this, but there is a lot more to learn, which we will explore while creating a CRUD API in Golang later in this series.

## Conclusion

I have also included some more examples of PATCH requests [**here**](https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/patch/).

That's it from the 36th part of the series, all the source code for the examples are linked in the GitHub on the [**100 days of Golang**](https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/patch/) repository.

[**100-days-of-golang**](https://github.com/Mr-Destructive/100-days-of-golang)

### References:

* [HTTP Request Methods](https://en.wikipedia.org/wiki/HTTP#Request_methods)
    
* [Ruby on Rails Patch method](https://rubyonrails.org/2012/2/26/edge-rails-patch-is-the-new-primary-http-method-for-updates)
    
* [HTTP PATCH RFC](https://datatracker.ietf.org/doc/html/rfc5789)
    
* [HTTP JSON PATCH RFC](https://datatracker.ietf.org/doc/html/rfc6902)
    
* [HTTP JSON MERGE PATCH](https://datatracker.ietf.org/doc/html/rfc7386)
