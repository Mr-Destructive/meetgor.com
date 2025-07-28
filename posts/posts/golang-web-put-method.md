{
  "title": "Golang Web: PUT Method",
  "status": "published",
  "slug": "golang-web-put-method",
  "date": "2025-04-05T12:33:26Z"
}

<h2>Introduction</h2>
<p>In this section of the series, we will be exploring how to send a <code>PUT</code> HTTP request in golang. We will understand how to send a basic PUT request, create an HTTP request, update a resource on a server, parsing the content from struct to json, headers, etc in the following section of this post.</p>
<h2>What is a PUT Method</h2>
<p>A PUT method is a type of request that is used to update or modify an entire resource on a server/database.</p>
<p>Imagine you have ordered a pizza at a restaurant and realized you want to change the toppings after you've already placed the order. With a PUT request, it's like informing the waiter about the changes you want to make to your existing order. You specify the updated toppings or any modifications (the data you send). The waiter then takes this updated information (PUT request) back to the kitchen (the server) to apply the changes to your order.</p>
<p>Let's say you created a order.</p>
<pre><code class="language-nginx">PUT /api/order/456 HTTP/1.1
Host: example.com
Content-Type: application/json
Content-Length: 123

{
    &quot;userID&quot;: 123,
    &quot;orderID&quot;: 456,
    &quot;items&quot;: [
        {
            &quot;itemID&quot;: 789,
            &quot;name&quot;: &quot;Pizza&quot;,
            &quot;quantity&quot;: 2,
            &quot;toppings&quot;: [&quot;Mushrooms&quot;]
        }
    ]
}
</code></pre>
<p>In the context of web development, PUT requests are often used for actions such as:</p>
<ul>
<li>
<p>Updating existing records or resources</p>
</li>
<li>
<p>Modifying specific parts of an existing resource</p>
</li>
<li>
<p>Replacing an entire resource with updated data</p>
</li>
</ul>
<p>Here's an example of what the PUT request might look like in this scenario:</p>
<pre><code class="language-nginx">PUT /api/order/456 HTTP/1.1
Host: example.com
Content-Type: application/json
Content-Length: 155

{
    &quot;userID&quot;: 123,
    &quot;orderID&quot;: 456,
    &quot;items&quot;: [
        {
            &quot;itemID&quot;: 789,
            &quot;name&quot;: &quot;Pizza&quot;,
            &quot;quantity&quot;: 2,
            &quot;toppings&quot;: [&quot;Mushrooms&quot;, &quot;Olives&quot;]
        }
    ]
}
</code></pre>
<p>In this example:</p>
<ul>
<li>
<p>The PUT method is used to update the resource identified by <code>/api/order/456</code>.</p>
</li>
<li>
<p>The application/json is the content type of the request.</p>
</li>
<li>
<p>The 155 is the content length of the request.</p>
</li>
<li>
<p>The body contains the updated details of the order, including the addition of toppings to the pizza.</p>
</li>
</ul>
<p>PUT requests are crucial for maintaining and updating data in applications where accuracy and consistency are paramount, ensuring that resources are kept current and reflect the latest changes made by users or systems</p>
<h2>Why the need of PUT Method</h2>
<p>In the world of HTTP requests, we use the PUT method to update or modify an entire resource on a server or database. This method is crucial because the POST method, while convenient for creating new data, is not intended for updating existing resources according to standard conventions. While it's possible to misuse the POST method for updates internally, doing so can lead to confusion and inconsistencies in how requests are understood and processed.</p>
<h2>How PUT Method request works</h2>
<p>A <a href="https://www.rfc-editor.org/rfc/rfc9110#PUT">PUT</a> request is utilized to send data to a server for the purpose of updating a resource. When a client (such as a browser or other APIs) sends a PUT request to the server's API endpoint, it includes data in the request body, typically formatted as JSON, XML, or form data.</p>
<p>The server processes the PUT request by first identifying the resource to be updated, either through the URL or data provided in the request body. It then validates, parses, and applies the data from the request body to make modifications to the resource. Following this, the server returns a response that includes a status code indicating the success or failure of the operation. Optionally, the response may also include the updated resource in the response body.</p>
<p>Unlike the POST method, which is primarily used for creating new resources, PUT is specifically designed for updating existing resources on the server. The request body of a PUT contains the data necessary for the update, while the URL identifies the specific resource to be updated.</p>
<p>In summary, PUT requests facilitate the transfer of data to the server specifically for updating resources, ensuring that changes to existing data are accurately processed and reflected.</p>
<h2>Basic PUT Method</h2>
<p>To send a <code>PUT</code> request to an API in golang, we need to create a <code>http.Request</code> object. For <code>POST</code> method, the <code>http</code> package had the <code>Post</code> function defined, however for <code>PUT</code> method, there is no separate function. The Go philosophy is right now against adding all the method functions. There have been a couple of discussions on this on <a href="https://github.com/golang/go/issues/22841">GitHub</a>, but it is not been adopted as of 2024.</p>
<p>So, we need to create a <code>http.Request</code> object for <code>PUT</code> method.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
    &quot;net/http&quot;
)

func main() {
	apiURL := &quot;https://reqres.in/api/users/5&quot;

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
</code></pre>
<pre><code class="language-bash">$ go run main.go

200
</code></pre>
<p>The above code sends a <code>PUT</code> request to the <a href="https://reqres.in/api/users/5"><code>https://reqres.in/api/users/5</code></a> endpoint. The resource we are trying to update is fetched with the identifier <code>5</code> which could probably be <code>id</code> for the user in the database of the server.</p>
<h2>PUT Method with JSON</h2>
<p>Marshaling and encoding are essential in Go for preparing structured data, such as from a struct, into JSON format suitable for HTTP requests like PUT. This conversion ensures data integrity and compatibility between Go types and JSON representations. It's crucial when updating resources on servers, as APIs often require specific data formats for processing updates correctly. Marshaling transforms Go structs into JSON bytes, while encoding further prepares them as request bodies, facilitating seamless communication with web services. This process ensures data consistency and adherence to API specifications, maintaining robust communication in distributed systems.</p>
<pre><code class="language-go">package main

import (
	&quot;bytes&quot;
	&quot;encoding/json&quot;
	&quot;fmt&quot;
	&quot;io&quot;
	&quot;net/http&quot;
)

type User struct {
	Name   string `json:&quot;name&quot;`
	Salary int    `json:&quot;salary&quot;`
	Age    string `json:&quot;age&quot;`
	ID     int    `json:&quot;id,omitempty&quot;`
}

type UserResponse struct {
	Status  string `json:&quot;status&quot;`
	Message string `json:&quot;message&quot;`
}

func main() {
	user := User{
		Name:   &quot;Alice&quot;,
		Salary: 50000,
		Age:    &quot;25&quot;,
	}
	apiURL := &quot;https://dummy.restapiexample.com/api/v1/update/11&quot;

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
		fmt.Println(&quot;too many requests&quot;)
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
	if err := json.Unmarshal(respBody, &amp;userResponse); err != nil {
		panic(err)
	}
	fmt.Println(userResponse)
}
</code></pre>
<pre><code class="language-bash">$ go run json.go                                                                 
{&quot;name&quot;:&quot;Alice&quot;,&quot;salary&quot;:50000,&quot;age&quot;:&quot;25&quot;}
200                                                                              
{&quot;status&quot;:&quot;success&quot;,&quot;data&quot;:[],&quot;message&quot;:&quot;Successfully! Record has been updated.&quot;}
{success Successfully! Record has been updated.}
</code></pre>
<p>In the provided Go code example, the <code>json.Marshal</code> function is used to convert a Go struct (<code>User</code>) into a JSON formatted byte slice (<code>[]byte</code>). Here's a breakdown of the steps involved:</p>
<ul>
<li>
<p>Struct Definition: Define a Go struct with json tags.</p>
</li>
<li>
<p>Marshalling: Use json.Marshal to convert the struct into JSON byte slice.</p>
</li>
<li>
<p>Buffer Creation: Wrap the JSON byte slice into an in-memory buffer (bytes.Buffer).</p>
</li>
<li>
<p>Request Sending: Send a PUT request with the buffer as the request body and set appropriate headers.</p>
</li>
</ul>
<p>Let's explore it step by step in detail:</p>
<p>When using the PUT method in Go to update a resource on a server, you often need to send data in JSON format as the request body. Here's how you can achieve this using marshaling and encoding:</p>
<ol>
<li>Define the Struct</li>
</ol>
<pre><code class="language-go">type User struct {
    Name   string `json:&quot;name&quot;`
    Salary int    `json:&quot;salary&quot;`
    Age    string `json:&quot;age&quot;`
    ID     int    `json:&quot;id,omitempty&quot;`
}
</code></pre>
<p>Define a Go struct (<code>User</code>) that represents the data structure you want to send in JSON format. The json tags specify how each field should be serialized into JSON.</p>
<ol start="2">
<li>Create an Object</li>
</ol>
<pre><code class="language-go">user := User{
    Name:   &quot;Alice&quot;,
    Salary: 50000,
    Age:    &quot;25&quot;,
}
</code></pre>
<p>Create an instance of the User struct (user) with sample data. This data will be marshaled into JSON format to send in the <code>PUT</code> request body.</p>
<ol start="3">
<li>Marshal the Struct</li>
</ol>
<pre><code class="language-go">bodyBytes, err := json.Marshal(user)
if err != nil {
    panic(err)
}
</code></pre>
<p>Use json.Marshal(user) to convert the user struct into a JSON byte slice (bodyBytes). This byte slice contains the serialized JSON representation of the User struct.</p>
<ol start="4">
<li>Create a Buffer</li>
</ol>
<pre><code class="language-go">body := bytes.NewBuffer(bodyBytes)
</code></pre>
<p>Use <code>bytes.NewBuffer(bodyBytes)</code> to create an in-memory buffer (<code>body</code>) containing the JSON byte slice (<code>bodyBytes</code>). The buffer implements the <code>io.Reader</code> interface needed for the PUT request body.</p>
<ol start="5">
<li>Create a PUT Request</li>
</ol>
<pre><code class="language-go">req, err := http.NewRequest(http.MethodPut, apiURL, body)
if err != nil {
    panic(err)
}
</code></pre>
<p>Use http.NewRequest to create a new PUT request to the specified URL with the JSON buffer (<code>body</code>) as the request body. Set appropriate headers if needed (e.g., Content-Type as application/json).</p>
<ol start="6">
<li>Send the Request</li>
</ol>
<pre><code class="language-go">resp, err := http.DefaultClient.Do(req)
if err != nil {
    panic(err)
}
</code></pre>
<p>Use <a href="http://http.DefaultClient.Do"><code>http.DefaultClient.Do</code></a><code>(req)</code> to execute the PUT request and obtain the response. Handle any errors that may occur during the request execution.</p>
<ol start="7">
<li>Process the Response</li>
</ol>
<pre><code class="language-go">respBody, err := io.ReadAll(resp.Body)
if err != nil {
    panic(err)
}
</code></pre>
<p>Use <code>io.ReadAll(resp.Body)</code> to read and store the response body from the server. Handle any errors encountered during the reading process.</p>
<ol start="8">
<li>Unmarshal the Response</li>
</ol>
<pre><code class="language-go">var userResponse UserResponse
if err := json.Unmarshal(respBody, &amp;userResponse); err != nil {
    panic(err)
}
</code></pre>
<p>Use <code>json.Unmarshal(respBody, &amp;userResponse)</code> to deserialize the JSON response body into a Go struct. This allows you to work with the response data in a structured manner.</p>
<p>The parsing of files and form data is also possible with <code>PUT</code> requests, however, that has been covered in the <a href="https://meetgor.com/golang-web-post-method">POST Method</a>. Those snippets would be handy in these request method as well.</p>
<p>I have also included some more examples of PUT requests <a href="https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/put/">here</a>.</p>
<p>That's it from the 35th part of the series, all the source code for the examples are linked in the GitHub on the <a href="https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/put/">100 days of Golang</a> repository.</p>
<p><a href="https://github.com/Mr-Destructive/100-days-of-golang">100-days-of-golang</a></p>
<h2>Conclusion</h2>
<p>That's it from this post of the series, a post on the PUT method in golang :)</p>
<p>We have covered topics like creating basic PUT requests and marshaling golang types into JSON format. Hope you found this article helpful. If you have any queries, questions, or feedback, please let me know in the comments or on my social handles. Thank you for reading.</p>
<p>Happy Coding :)</p>
