{
  "title": "Golang Web: POST Method",
  "status": "published",
  "slug": "golang-web-post-method",
  "date": "2025-04-05T12:33:27Z"
}

<h2>Introduction</h2>
<p>In this section of the series, we will be exploring how to send a <code>POST</code> HTTP request in golang. We will understand how to send a basic POST request, create an HTTP request, and parse json, structs into the request body, add headers, etc in the following sections of this post. We will understand how to marshal the golang struct/types into JSON format, send files in the request, and handle form data with examples of each in this article. Let's answer a few questions first.</p>
<h2>What is a POST request?</h2>
<p>POST method is a type of request that is used to send data to a server(a machine on the internet).</p>
<p>Imagine you are placing an order at a restaurant. With a GET request, it would be like asking the waiter, &quot;What kind of pizza do you have?&quot; The waiter would respond by telling you the menu options (the information retrieved from the server).</p>
<p>However, a POST request is more like giving your completed order to the waiter. You tell them the specific pizza you want, its size, and any additional toppings (the data you send). The waiter then takes this information (POST request) back to the kitchen (the server) to process it (fulfill your order).</p>
<p>In the world of web development, POST requests are often used for things like:</p>
<p>Submitting forms (e.g., contact forms, login forms) Uploading files (e.g., photos, videos) Creating new accounts Sending data to be processed (e.g., online purchases)</p>
<p>Here's an example of what the POST request might look like in this scenario:</p>
<pre><code class="language-nginx">POST /api/order HTTP/1.1
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
            &quot;quantity&quot;: 2
        },
        {
            &quot;itemID&quot;: 999,
            &quot;name&quot;: &quot;Burger&quot;,
            &quot;quantity&quot;: 1
        }
    ]
}
</code></pre>
<p>In this example:</p>
<ul>
<li>The <code>POST</code> method is used to send data to the server.</li>
<li>The <code>/api/order</code> is the endpoint of the server.</li>
<li>The <code>application/json</code> is the content type of the request.</li>
<li>The <code>123</code> is the content length of the request.</li>
<li>The <code>{&quot;userID&quot;: 123, &quot;orderID&quot;: 456, &quot;items&quot;: [{&quot;itemID&quot;: 789, &quot;name&quot;: &quot;Pizza&quot;, &quot;quantity&quot;: 2}, {&quot;itemID&quot;: 999, &quot;name&quot;: &quot;Burger&quot;, &quot;quantity&quot;: 1}]}</code> is the body of the request.</li>
</ul>
<h2>Why the need for a POST request?</h2>
<p>In the world of HTTP requests, we use the POST method to securely send data from a client (like a user's browser) to a server. This is crucial because the GET method, while convenient for retrieving data, has limitations:</p>
<p>Imagine you are in registering for an event via Google form, you type in your details on the webpage like name, email, address, phone number, and other personal details. If the website/app was using the <code>GET</code> method to send the request to register or do any other authentication/privacy-related requests, it could expose the data in the URL itself. It would be something along the lines <a href="https://form.google.com/register/%3Cform-name%3E-%3Cid%3E/?name=John&amp;phone_number=1234567890"><code>https://form.google.com/register/&lt;form-name&gt;-&lt;id&gt;/?name=John&amp;phone_number=1234567890</code></a>, if a user maliciously sniffs into your network and inspects the URL, your data will be exposed. That is the reason we need <code>POST</code> a method.</p>
<h2>How a POST method works?</h2>
<p>A <a href="https://www.rfc-editor.org/rfc/rfc9110#POST">POST</a> request is used to send data to a server to create or update(there is a separate method for updating) a resource. The client(browser/other APIs) sends a POST request to the server's API endpoint with the data in the request body. This data can be in formats like JSON, XML, or form data. The server processes the POST request, validates and parses the data in the request body, makes any changes or creates resources based on that data, and returns a response. The response would contain a status code indicating the success or failure of the operation and may contain the newly created or updated resource in the response body. The client must check the response status code to verify the outcome and process the response accordingly. Unlike GET, POST can create new resources on the server. The body of a POST contains the data for creation while the URL identifies the resource to be created. Overall, POST transfers data to the server for processing, creation or updating of resources.</p>
<p>The status code is usually <code>201</code> indicating the resource is successfully created or <code>200</code> for just indicating success.</p>
<p>Some common steps for creating and sending a POST request as a developer include:</p>
<ul>
<li>
<p>Defining the API endpoint</p>
</li>
<li>
<p>Clarifying the data format (json, language native objects, xml , text, form-data, etc)</p>
</li>
<li>
<p>Converting / Marshalling the data</p>
</li>
<li>
<p>Attaching header for <code>Content-Type</code> as key and value as the format of the data type (e.g. <code>application/json</code> for json)</p>
</li>
<li>
<p>Sending the request</p>
</li>
</ul>
<p>The above steps are general for creating and sending a POST request, they are not specific to Golang. For golang specific steps, we need to dive a bit deeper, let's get started.</p>
<h2>Basic POST method in Golang</h2>
<p>To send a POST request in golang, we need to use the <code>http</code> package. The <code>http</code> package has the <code>Post</code> method, which takes in 3 parameters, namely the URL, the Content-Type, and the Body. The body can be <code>nil</code> if the URL endpoint doesn't necessarily require a body. The <code>Content-Type</code> is the string, since we are just touching on how the Post request is constructed, we will see what the <code>Content-Type</code> string value should be in the later sections.</p>
<blockquote>
<p><a href="http://http.Post">http.Post</a>(URL, Content-Type, Body)</p>
</blockquote>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/http&quot;
)

func main() {
	apiURL := &quot;https://reqres.in/api/users&quot;

	// POST request
	resp, err := http.Post(apiURL, &quot;&quot;, nil)
    // ideally the Content-Type header should be set to the relevant format
	// resp, err := http.Post(apiURL, &quot;application/json&quot;, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
    fmt.Println(resp)
	defer resp.Body.Close()
}
</code></pre>
<pre><code class="language-bash">
$ go run main.go

201
&amp;{
    201 Created
    201
    HTTP/2.0
    2
    0
    map[
        Access-Control-Allow-Origin:[*]
        Cf-Cache-Status:[DYNAMIC]
        Cf-Ray:[861cd9aec8223e4b-BOM]
        Content-Length:[50]
        Content-Type:[application/json; charset=utf-8]
        Date:[Sat, 09 Mar 2024 17:40:28 GMT]
        Server:[cloudflare]
        ...
        ...
        ...
        X-Powered-By:[Express]
    ]
    {0xc00017c180}
    50
    []
    false
    false
    map[]
    0xc000156000
    0xc00012a420
}
</code></pre>
<p>The above code is sending the <code>POST</code> request to the <a href="https://reqres.in/api/users"><code>https://reqres.in/api/users</code></a> endpoint with an empty body and no specific format for <code>Content-Type</code> header. The response is according to the <a href="https://pkg.go.dev/net/http#Response">Response</a> structure. We can see we got <code>201</code> status, which indicates the server received the POST request successfully, the API is a dummy api, so we don't care about the data we are processing, we are just using the API as a placeholder for sending the POST request.</p>
<p>We can use <code>map[string]interface{}</code> it to pass the data in the request body. The <code>json.Marshal</code> method is used to convert the map into JSON format. We will look into the details shortly in the next few examples.</p>
<pre><code class="language-go">package main

import (
	&quot;bytes&quot;
	&quot;encoding/json&quot;
	&quot;fmt&quot;
	&quot;net/http&quot;
)

func main() {
    apiURL := &quot;https://reqres.in/api/users&quot;
    bodyMap := map[string]interface{}{
        &quot;name&quot;: &quot;morpheus&quot;,
        &quot;job&quot;: &quot;leader&quot;,
    }

    requestBody, err := json.Marshal(bodyMap)
    if err != nil {
        panic(err)
    }
    body := bytes.NewBuffer(requestBody)

    resp, err := http.Post(apiURL, &quot;application/json&quot;, body)
    if err != nil {
        panic(err)
    }
    fmt.Println(resp.StatusCode)
    defer resp.Body.Close()
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

201
</code></pre>
<p>The above code sends the <code>POST</code> request to the <a href="https://reqres.in/api/users"><code>https://reqres.in/api/users</code></a> endpoint with the data in the request body in JSON format.</p>
<h2>Creating a POST request in Golang</h2>
<p>We can construct the POST request with the <a href="https://pkg.go.dev/net/http#NewRequest">NewRequest</a> method. The method takes in 3 parameters, namely the <code>method</code> (e.g. <code>POST</code>, <code>GET</code>), the <code>URL</code> and the <code>body</code> (if there is any). We can then add extra information to the headers or the Request object after constructing the basic HTTP <a href="https://pkg.go.dev/net/http#Request">Request</a> object.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/http&quot;
)

func main() {
	apiURL := &quot;https://reqres.in/api/users&quot;

	req, err := http.NewRequest(http.MethodPost, apiURL, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	//fmt.Println(resp)
	defer resp.Body.Close()
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

201
</code></pre>
<p>In the above example, we have created an HTTP Request as the <code>POST</code> method, with <a href="https://reqres.in/api/users"><code>https://reqres.in/api/users</code></a> as the URL, and no body. This constructs an HTTP Request object, which can be sent as the parameter to the <a href="http://http.DefaultClient.Do">http.DefaultClient.Do</a> method, which is the default client for the request we sent in the earlier examples as <code>http.Get</code> or <a href="http://http.Post"><code>http.Post</code></a> methods. We can implement a custom client as well, and then apply <code>Do</code> the method with the request parameters. The <code>Do</code> method returns the <code>Request</code> object or the <code>error</code> if any.</p>
<p>More on the customizing Client will be explained in a separate post in the series.</p>
<p>The response is also in the same format as the <a href="https://pkg.go.dev/net/http#Response">Response</a> structure that we have seen earlier. This section of the series aims to construct a post request, and not to parse the response, we have already understood the parsing of the response in the <a href="https://www.meetgor.com/golang-web-get-method/#?:~:text=Parsing%20the%20JSON%20body%20with%20structs">Get method</a> section of the series.</p>
<h3>Parsing objects to JSON for POST method request</h3>
<p>We might have a golang object that we want to send as a body to an API in the POST request, for that we need to convert the golang struct object to JSON. We can do this by using the <a href="https://pkg.go.dev/encoding/json#Marshal">Marshal</a> or the <a href="https://pkg.go.dev/encoding/json#Encoder.Encode">Encode</a> method for serialization of the golang struct object to JSON.</p>
<h4>Using Marshal method</h4>
<p>Marshaling is the process of converting data from a data structure into a format suitable for transmission over a network or for storage. It's commonly used to convert native objects in a programming language into a serialized format, typically a byte stream, that can be transmitted or stored efficiently. You might get a question here, what is the difference between <code>Marshalling</code> and <code>Serialization</code>? Well, Serialization, is a broader term that encompasses marshalling. It refers to the process of converting an object or data structure into a format that can be stored or transmitted and later reconstructed into the original object. Serialization may involve converting data into byte streams, XML, JSON, or other formats. So, in summary, marshaling specifically deals with converting native objects into a format suitable for transmission, while serialization encompasses the broader process of preparing data for storage or transmission.</p>
<p>The <code>json</code> package has the <a href="https://pkg.go.dev/encoding/json#Marshal">Marshal</a> method that converts the golang object into JSON. The <code>Marshal</code> method takes in a parameter as the struct object with type <code>any</code> and returns a byte slice <code>[]byte</code> and error (if any).</p>
<pre><code class="language-go">package main

import (
	&quot;bytes&quot;
	&quot;encoding/json&quot;
	&quot;fmt&quot;
	&quot;net/http&quot;
)

type User struct {
	Name   string `json:&quot;name&quot;`
	Salary int    `json:&quot;salary&quot;`
	Age    int    `json:&quot;age&quot;`
}

func main() {
	user := User{
		Name:   &quot;Alice&quot;,
		Salary: 50000,
		Age:    25,
	}
	apiURL := &quot;https://dummy.restapiexample.com/api/v1/create&quot;

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
	resp, err := http.Post(apiURL, &quot;application/json&quot;, body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

{&quot;name&quot;:&quot;Alice&quot;,&quot;salary&quot;:50000,&quot;age&quot;:25}
200
</code></pre>
<p>In the above example, we have created a struct <code>User</code> with fields <code>Name</code>, <code>Salary</code>, and <code>Age</code>, the json tags will help label each key in JSON with the tag for the respective fields in the struct. We create an object <code>user</code> of a type <code>User</code> with the values as <code>Alice</code>, <code>50000</code>, and <code>25</code> respectively.</p>
<p>We call the <code>json.Marshal</code> method with the parameter <code>user</code> that represents the struct object <code>User</code>, the method returns a slice of bytes, or an error either or both could be nil. If we try to see the stringified representation of the byte slice, we can see something like <code>{&quot;name&quot;:&quot;Alice&quot;,&quot;salary&quot;:50000,&quot;age&quot;:25}</code> which is a JSON string for the user struct. We can't parse the byte slice as the body in the POST request, we need the <code>io.Reader</code> object, so we can load the byte slice <code>bodyBytes</code> into a buffer and parse that as a body for the POST request.</p>
<p>We then send a <code>POST</code> request to the endpoint <a href="https://dummy.restapiexample.com/api/v1/create"><code>https://dummy.restapiexample.com/api/v1/create</code></a> with the content type as <code>application/json</code> and with the body as <code>body</code> which was a <code>io.Reader</code> object as an in-memory buffer.</p>
<p>In brief, we can summarize the marshaling of the golang object into JSON with <code>Marshal</code> function as the following steps:</p>
<ul>
<li>
<p>Defining the structure as per the request body</p>
</li>
<li>
<p>Creating the struct object for parsing the data as body to the request</p>
</li>
<li>
<p>Calling the <code>json.Marshal</code> function to convert the object to JSON (parameter as the struct object <code>any</code> type)</p>
</li>
<li>
<p>Loading the byte slice into a buffer with <code>bytes.NewBuffer()</code></p>
</li>
<li>
<p>Sending the POST request to the endpoint with the body as the <code>io.Reader</code> object and content type as <code>application/json</code></p>
</li>
</ul>
<h4>Using Encode method</h4>
<p>We can even use the <a href="https://pkg.go.dev/encoding/json#Encoder.Encode">Encoder.Encode</a> method to parse the golang struct object to JSON. Firstly, we should have the struct defined as per the request body that the particular API takes, we can make use of the json tags, omitempty, omit(-) options to make the marshaling process work accordingly. We can then create the object of that particular struct with the data we require to be created as a resource with the POST request on that API service.</p>
<p>Thereafter we can create an empty buffer object with <a href="https://pkg.go.dev/bytes#Buffer">bytes.Buffer</a>, this buffer object would be used to initialize the <a href="https://pkg.go.dev/encoding/json#Encoder">Encoder</a> object with the <a href="https://pkg.go.dev/encoding/json#NewEncoder">NewEncoder</a> method. This would give access to the <a href="https://pkg.go.dev/encoding/json#Encoder.Encode">Encode</a> method, which is used to take in the struct object (<code>any</code> type) and this will populate the buffer we initialized with the <code>NewEncoder</code> method.</p>
<p>Later we can access that buffer to parse it to the Post request as the body. Let's understand it better with an example.</p>
<pre><code class="language-go">package main

import (
	&quot;bytes&quot;
	&quot;encoding/json&quot;
	&quot;fmt&quot;
	&quot;net/http&quot;
)

type User struct {
	Name   string
	Salary int
	Age    int
}

func main() {
	user := User{
		Name:   &quot;Alice&quot;,
		Salary: 50000,
		Age:    25,
	}
	apiURL := &quot;https://dummy.restapiexample.com/api/v1/create&quot;

	var bodyBuffer bytes.Buffer
	var encoder = json.NewEncoder(&amp;bodyBuffer)
	err := encoder.Encode(user)
    if err != nil {
        panic(err)
    }

	resp, err := http.Post(apiURL, &quot;application/json&quot;, &amp;bodyBuffer)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp)
	defer resp.Body.Close()
}
</code></pre>
<p>Over here, we have created a struct <code>User</code> with fields <code>Name</code>, <code>Salary</code>, and <code>Age</code>, we initialize the <code>user</code> as the object of the <code>User</code> struct. Then we create a buffer <code>bodyBuffer</code> of type <code>bytes.Buffer</code> this is the actual buffer that we will send as the body. Further, we initialize the <code>Encoder</code> object as <code>encoder</code> with the <code>json.NewEncoder</code> method by parsing the reference of <code>bodyBuffer</code> as the parameter. Since <code>bytes.Buffer</code> implements the <code>io.Writer</code> interface, we can pass the <code>bodyBuffer</code> to the <code>NewEncoder</code> method. This will create the <code>Encoder</code> object which in turn will give us access to the <code>Encode</code> method, where we will parse the struct instance and it will populate the buffer with which we initialized the <code>Encoder</code> object earlier.</p>
<p>Now, we have the <code>encode</code> object, this gives us the access to <code>Encode</code> method, we call the <code>Encode</code> method with the parameter of <code>user</code> which is a User struct instance/object. The Encode method will populate the <code>bodyBuffer</code> object or it will result in an error if anything goes wrong (the data is incorrectly parsed or is not in the required format).</p>
<p>We can call the <code>Post</code> method with the initialized URL, the <code>Content-Type</code> as <code>application/json</code> since we have converted the struct instance to JSON object, and the body as the reference to the buffer as <code>&amp;bodyBuffer</code></p>
<p>So, the steps for parsing struct instances into JSON objects with the <code>Encoder.Encode</code> method is as follows:</p>
<ul>
<li>
<p>Defining the structure as per the request body</p>
</li>
<li>
<p>Creating the struct object for parsing the data as body to the request</p>
</li>
<li>
<p>Creating an empty <code>bytes.Buffer</code> object as an in-memory buffer</p>
</li>
<li>
<p>Initializing the <code>Encoder</code> object with <code>NewEncoder</code> method by parsing the reference of <code>bodyBuffer</code> as the parameter</p>
</li>
<li>
<p>Calling the <code>Encode</code> method with the parameter of struct instance/object</p>
</li>
<li>
<p>Sending the POST request to the endpoint with the content type as <code>application/json</code> and body as the reference to the buffer</p>
</li>
</ul>
<p>The results are the same as the above example just the way we have parsed the struct instance to JSON object is different.</p>
<h3>Parsing JSON to POST request</h3>
<p>We have seen how we can parse golang struct instances to JSON and then send the post request, but what if we had the JSON string already with us, and we want to send the request? Well, that's much easier, right? We already have parsed the JSON string to the Post request by loading the slice of bytes into a buffer, so we just need to convert the string to a slice of bytes which is quite an easy task, and then load that byte slice to the buffer.</p>
<pre><code class="language-go">package main

import (
	&quot;bytes&quot;
	&quot;fmt&quot;
	&quot;net/http&quot;
)

func main() {
	// dummy api
	apiURL := &quot;https://dummy.restapiexample.com/api/v1/create&quot;

	// json data
	data := `{
        &quot;name&quot;: &quot;Alice&quot;,
        &quot;job&quot;: &quot;Teacher&quot;
    }`
	body := bytes.NewBuffer([]byte(data))

	// POST request
	resp, err := http.Post(apiURL, &quot;application/json&quot;, body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp)
	defer resp.Body.Close()
}
</code></pre>
<p>In the example above, we already have a JSON string <code>data</code> with keys as <code>name</code> and <code>job</code> but it is not JSON, it is a stringified JSON. We can convert the stringified JSON to a slice of bytes using the <code>[]byte</code> function. Further, we have used the <code>bytes.NewBuffer</code> method to load the byte slice into an <code>io.Reader</code> object. This object returned by the <code>bytes.NewBuffer</code> will serve as the body for the POST request.</p>
<h3>Parsing JSON to objects in Golang from POST method response</h3>
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
	Status string `json:&quot;status&quot;`
	Data   User   `json:&quot;data&quot;`
}

func main() {
	user := User{
		Name:   &quot;Alice&quot;,
		Salary: 50000,
		Age:    &quot;25&quot;,
	}
	apiURL := &quot;https://dummy.restapiexample.com/api/v1/create&quot;

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
	resp, err := http.Post(apiURL, &quot;application/json&quot;, body)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(resp)
	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// unmarshalling process
	// converting JSON to Go specific data structure/types
	var userResponse UserResponse
	if err := json.Unmarshal(respBody, &amp;userResponse); err != nil {
		panic(err)
	}
	fmt.Println(userResponse)
	fmt.Println(userResponse.Data)
}
</code></pre>
<pre><code class="language-nginx">
{success {Alice 50000 25 3239}}
{Alice 50000 25 577}
</code></pre>
<p>The above example is a POST request with a struct instance being loaded as a JSON string and then sent as a buffer to the API endpoint, it also reads the response body with a specific structure <code>UserResponse</code> and unmarshalled the <code>resp.Body</code> from the <code>io.Reader</code> as <code>respBody</code> and then loads into <code>userResponse</code> object. This example gives an entire process of what we have understood in the JSON data parsing for a POST request.</p>
<h3>Sending Form data in a POST request</h3>
<p>We can also send data to a POST request in the form of a form, the form which we use in the HTML. Golang has a <code>net/url</code> package to parse the form data. The form data is sent in the <code>application/x-www-form-urlencoded</code> format.</p>
<pre><code class="language-go">package main

import (
	&quot;encoding/json&quot;
	&quot;fmt&quot;
	&quot;io&quot;
	&quot;net/http&quot;
	&quot;net/url&quot;
	&quot;strings&quot;
)

type ResponseLogin struct {
	Token string `json:&quot;token&quot;`
}

func main() {
	// dummy api
	apiURL := &quot;https://reqres.in/api/login&quot;

	// Define form data
	formData := url.Values{}
	formData.Set(&quot;email&quot;, &quot;eve.holt@reqres.in&quot;)
	formData.Set(&quot;password&quot;, &quot;cityslicka&quot;)

	// Encode the form data
	fmt.Println(formData.Encode())
	reqBody := strings.NewReader(formData.Encode())
	fmt.Println(reqBody)

	// Make a POST request with form data
	resp, err := http.Post(apiURL, &quot;application/x-www-form-urlencoded&quot;, reqBody)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print response status code
	fmt.Println(&quot;Status Code:&quot;, resp.StatusCode)

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
	token := ResponseLogin{}

	json.Unmarshal(respBody, &amp;token)
	fmt.Println(token)
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

email=eve.holt%40reqres.in&amp;password=cityslicka
&amp;{email=eve.holt%40reqres.in&amp;password=cityslicka 0 -1}
Status Code: 200
{QpwL5tke4Pnpja7X4}
</code></pre>
<p>In the above example, we set a <code>formData</code> with the values of <code>email</code> and <code>password</code> which are <code>url.Values</code> object. The <code>url.Values</code> the object is used to store the key-value pairs of the form data. The <code>formData</code> is encoded with the <code>url.Encode</code> method, We load the encoded string to a <code>buffer</code> with <code>strings.NewReader</code> which implements the <code>io.Reader</code> interface, so that way we can pass that object as the body to the post request.</p>
<p>We send the <code>POST</code> request to the endpoint <a href="https://reqres.in/api/login"><code>https://reqres.in/api/login</code></a> with the content type as <code>application/x-www-form-urlencoded</code> and with the body as <code>reqBody</code> which implements the <code>io.Reader</code> interface as an in-memory buffer. The response from the request is read into the buffer with <code>io.ReadAll</code> method and we can <code>Unmarshal</code> the stream of bytes as a buffer into the <code>ResponseLogin</code> struct object.</p>
<p>The output shows the <code>formData</code> as encoded string <code>email=eve.holt%</code><a href="http://40reqres.in"><code>40reqres.in</code></a><code>&amp;password=cityslicka</code> as <code>@</code> is encoded to <code>%40</code>, then we wrap the <code>formData</code> in a <code>strings.NewReader</code> object which is a buffer that implements <code>io.Reader</code> interface, hence we can see the result as the object. The status code for the request is <code>200</code> indicating the server received the <code>form-data</code> in the body and upon unmarshalling, we get the token as a response to the POST request which was a dummy login API.</p>
<p>This way we have parsed the form-data to the body of a POST request.</p>
<h3>Sending File in a POST request</h3>
<p>We have covered, parsing text, JSON, and form data, and now we need to move into sending files in a POST request. We can use the <code>multipart</code> package to parse files into the request body and set appropriate headers for reading the file from the API services.</p>
<p>We first read the file contents <a href="http://os.Open"><code>os.Open</code></a> which returns a reference to the <code>file</code> object or an error. We create an empty <code>bytes.Buffer</code> object as <code>body</code> which will be populated later. The <a href="https://pkg.go.dev/mime/multipart#NewWriter">multipart.NewWriter</a> method takes in the <code>io.Writer</code> object which will be the <code>body</code> as it is an <code>bytes.Buffer</code> object that implements the <code>io.Writer</code> interface. This will initialize the <a href="https://pkg.go.dev/mime/multipart#Writer">Writer</a> object in the <code>multipart</code> package.</p>
<p>We create a <code>form-field</code> in the <code>Writer</code> object with the <a href="https://pkg.go.dev/mime/multipart#Writer.CreateFormFile">CreateFormFile</a> method, which takes in the <code>fieldName</code> as the name of the field, and the <code>fileName</code> as the name of the file which will be read later in the multipart form. The method returns either the part or the error. The <code>part</code> is an object that implements the <code>io.Writer</code> interface.</p>
<p>Since we have stored the file contents in the <code>file</code> object, we copy the contents into the <code>form-field</code> with the <a href="https://pkg.go.dev/io#Copy">Copy</a> method. Since the <code>part</code> return from the <code>CreateFormFile</code> was implementing the <code>io.Writer</code> interface, we can use it to Copy the contents from source to destination. The source is the <code>io.Reader</code> object and the destination is the <code>io.Writer</code> object, the destination for the <code>Copy</code> method is the first parameter, the source is the second parameter.</p>
<p>This Copy method will populate the buffer initialized earlier in the <code>NewWriter</code> method. This will give us a buffer that has the file contents in it. We can pass this buffer to the POST request with the <code>body</code> parameter. We also need to make sure we close the <code>Writer</code> object after copying the contents of the file. We can extract the type of file which will serve as the <code>Content-Type</code> of the request.</p>
<p>Let's clear the explanation with an example.</p>
<pre><code class="language-go">package main

import (
	&quot;bytes&quot;
	&quot;encoding/json&quot;
	&quot;fmt&quot;
	&quot;io&quot;
	&quot;mime/multipart&quot;
	&quot;net/http&quot;
	&quot;os&quot;
)

type ResponseFile struct {
	Files map[string]string `json:&quot;files&quot;`
}

func main() {
	apiURL := &quot;http://postman-echo.com/post&quot;
	fileName := &quot;sample.csv&quot;

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	body := &amp;bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(&quot;csvFile&quot;, fileName)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

    contentType := writer.FormDataContentType()
    fmt.Println(contentType)

	writer.Close()

	resp, err := http.Post(apiURL, contentType, body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(&quot;Status Code:&quot;, resp.StatusCode)

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	token := ResponseFile{}
	json.Unmarshal(respBody, &amp;token)
	fmt.Println(token)
	fmt.Println(token.Files[fileName])
}
</code></pre>
<pre><code class="language-bash">multipart/form-data; boundary=7e0eacfff890be395eba19c70415c908124b503a56f23ebeec0ab3c665ca


--619671ea2c0aa47ca6664a7cda422169d73f3b8a089c659203f5413d03de
Content-Disposition: form-data; name=&quot;csvFile&quot;; filename=&quot;sample.csv&quot;
Content-Type: application/octet-stream

User,City,Age,Country
Alex Smith,Los Angeles,20,USA
John Doe,New York,30,USA
Jane Smith,Paris,25,France
Bob Johnson,London,40,UK

--619671ea2c0aa47ca6664a7cda422169d73f3b8a089c659203f5413d03de--



Status Code: 200

{map[sample.csv:data:application/octet-stream;base64,VXNlcixDaXR5LEFnZSxDb3VudHJ5CkFsZXggU21pdGgsTG9zIEFuZ2VsZXMsMjAsVVNBCkpvaG4gRG9lLE5ldyBZb3JrLDMwLFVTQQpKYW5lIFNtaXRoLFBhmlzLDI1LEZyYW5jZQpCb2IgSm9obnNvbixMb25kb24sNDAsVUsK]}

data:application/octet-stream;base64,VXNlcixDaXR5LEFnZSxDb3VudHJ5CkFsZXggU21pdGgsTG9zIEFuZ2VsZXMsMjAsVVNBCkpvaG4gRG9lLE5ldyBZb3JrLDMwLFVTQQpKYW5lIFNtaXRoLFBhmlzLDI1LEZyYW5jZQpCb2IgSm9obnNvbixMb25kb24sNDAsVUsK
</code></pre>
<p>In the above example, we first read the file <code>sample.csv</code> into the <code>file</code> object with <a href="http://os.Open"><code>os.Open</code></a> method, this will return a reference to the file object or return an error if any arises while opening the file.</p>
<p>Then we create an empty buffer <code>bytes.Buffer</code> object which will serve as the body of the post request later as it will get populated with the file contents in the form of <code>multipart/form-data</code>.</p>
<p>We initialize the <code>Writer</code> object with <code>multipart.NewWriter</code> method which takes in the empty buffer as the parameter, we parse the <code>body</code> as the parameter. The method will return a reference to the <code>multipart.Writer</code> object.</p>
<p>With the <code>Writer</code> object we access the <code>CreateFormFile</code> method which takes in the <code>fieldName</code> as the name of the field, and the <code>fileName</code> as the name of the file. The method will return either the part or an error. The <code>part</code> in this case, is the reference to the <code>io.Writer</code> object that will be used to write the contents from the uploaded file.</p>
<p>Then, we can use the <code>io.Copy</code> method to copy the contents from the <code>io.Reader</code> object to the <code>io.Writer</code> object. The source is the <code>io.Reader</code> object and the destination is the <code>io.Writer</code> object. The first parameter is however the destination and the second parameter is the source. In the example, we call <code>io.Copy(part, file)</code> which will copy the contents of <code>file</code> to the <code>part</code> buffer.</p>
<p>We get the <code>Content-Type</code> by calling the <a href="https://pkg.go.dev/mime/multipart#Writer.FormDataContentType">Writer.FormDataContentType</a> method. This returns us <code>multipart/form-data; boundary=7e0eacfff890be395eba19c70415c908124b503a56f23ebeec0ab3c665ca</code> which will serve the <code>Content-Type</code> for the Post request.</p>
<p>We need to make sure we close the <code>Writer</code> object with the <code>Close</code> method.</p>
<p>We just print the <code>body.String()</code> to get a look at what the actual body looks like, we can see there is a form for the file as a <code>form-data</code> with keys like <code>Content-Type</code>, <code>Content-Disposition</code>, etc. The file has the <code>Content-Type</code> as <code>application/octet-stream</code> and the actual content is rendered in the output.</p>
<p>The dummy API responds with a 200 status code and also sends the JSON data with the name of the file as the key and the value as the <code>base64</code> encoded value of the file contents. This indicates that we were able to upload the file to the server API using a POST request. Well done!</p>
<p>I have also included some more examples of POST requests with files <a href="https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/post/file_2.go">here</a> which extends the above example by taking the encoded values and decoding to get the actual contents of the file back.</p>
<h2>Best Practices for POST method</h2>
<p>Here are some of the best practices for the POST method which are followed to make sure you consume or create the POST request in the most secure, efficient, and graceful way.</p>
<h3>Always Close the Response Body</h3>
<p>Ensure that you close the response body after reading from it. Use <code>defer response.Body.Close()</code> to automatically close the body when the surrounding function returns. This is crucial for releasing associated resources like network connections or file descriptors. Failure to close the response body can lead to memory leaks, particularly with a large volume of requests. Properly closing the body prevents resource exhaustion and maintains efficient memory usage.</p>
<h3>Client Customization</h3>
<p>Utilize the <a href="https://pkg.go.dev/net/http#Client">Client</a> struct to customize the HTTP client behavior. By using a custom client, you can set timeouts, headers, user agents, and other configurations without modifying the <code>DefaultClient</code> provided by the <code>http</code> package. This approach allows for flexibility and avoids repetitive adjustments to the client configuration for each request.</p>
<h3>Set Content-Type Appropriately</h3>
<p>Ensure that you set the <code>Content-Type</code> header according to the request payload. Correctly specifying the Content-Type is crucial for the server to interpret the request payload correctly. Failing to set the Content-Type header accurately may result in the server rejecting the request. Always verify and match the Content-Type header with the content being sent in the POST request to ensure smooth communication with the server.</p>
<h2>Reference</h2>
<ul>
<li>
<p><a href="https://www.postman.com/postman/workspace/postman-answers/documentation/13455110-00378d5c-5b08-4813-98da-bc47a2e6021d">Postman POST API</a> (For POST request with file upload)</p>
</li>
<li>
<p><a href="https://pkg.go.dev/net/http">Golang net/http Package</a></p>
</li>
</ul>
<h2>Conclusion</h2>
<p>That's it from this post of the series, a post on the POST method in golang :)</p>
<p>We have covered topics like creating basic post requests, Marshalling golang types into JSON format, parsing form data, sending a POST request with files, and best practices for the POST method. Hope you found this article helpful. If you have any queries, questions, or feedback, please let me know in the comments or on my social handles.</p>
<p>Happy Coding :)</p>
