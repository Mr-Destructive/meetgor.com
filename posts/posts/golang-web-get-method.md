{
  "title": "Golang Web: GET Method",
  "status": "published",
  "slug": "golang-web-get-method",
  "date": "2025-04-05T12:33:29Z"
}

<h2>Introduction</h2>
<p>In this section of the series, we will be exploring how to send a <code>GET</code> HTTP request in golang. We will be understanding how to send a basic GET request, create an HTTP request and customize the client, add headers, read the response body, etc in the following sections of this post.</p>
<h2>What is a GET method?</h2>
<p>A <a href="https://en.wikipedia.org/wiki/HTTP#Request_methods">GET</a> method in the context of an HTTP request is an action that is used to obtain data/resources. The <code>GET</code> method is used in a web application to get a resource like an HTML page, image, video, media, etc.</p>
<p>Some common usecases of the <code>GET</code> method are:</p>
<ul>
<li>Loading a webpage</li>
<li>Getting an image, file or other resource</li>
<li>API requests to retrieve data</li>
<li>Search queries sending filters and parameters</li>
</ul>
<h2>Basic GET Request</h2>
<p>To use the HTTP method <code>GET</code> in golang, the <a href="https://pkg.go.dev/net/http">net/http</a> package has a <a href="https://pkg.go.dev/net/http#Get">Get</a> method. This method simply takes in a URL as a string and returns the <a href="https://pkg.go.dev/net/http#Response">response</a> or an error. Let's look at how to send a basic HTTP GET request in golang.</p>
<pre><code class="language-go">// web/methods/get/main.go


package main

import (
	&quot;fmt&quot;
	&quot;net/http&quot;
)

func main() {
	reqURL := &quot;https://www.google.com&quot;
	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
    fmt.Println(resp.Status)
    fmt.Println(&quot;Status Code:&quot;, resp.StatusCode)
}
</code></pre>
<pre><code>$ go run main.go

&amp;{200 OK 200 HTTP/2.0 2 0 map[Alt-Svc:[h3=&quot;:443&quot;; ma=2592000,h3-29=&quot;:443&quot;; ma=2592000] Cache-Control:[private, max-age=0] Content-Security-Policy-Report-Only:[object-src 'none';base-uri 'self';script-src 'nonce-pdz_s8Gr0owwMbX8I9qNEQ' 'strict-dynamic' 'report-sample' 'unsafe-eval' 'unsafe-inline' https: http:;report-uri https://csp.withgoogle.com/csp/gws/other-hp] Content-Type:[text/html; charset=ISO-8859-1] Date:[Fri, 27 Oct 2023 09:37:04 GMT] Expires:[-1] P3p:[CP=&quot;This is not a P3P policy! See g.co/p3phelp for more info.&quot;] Server:[gws] Set-Cookie:[1P_JAR=2023-10-27-09; expires=Sun, 26-Nov-2023 09:37:04 GMT; path=/; domain=.google.com; Secure AEC=Ackid1Q5FARA_9d7f7znggUdw6DoJA1DBpI17Z0SWxN519Dc64EqmYVHlFg; expires=Wed, 24-Apr-2024 09:37:04 GMT; path=/; domain=.google.com; Secure; HttpOnly; SameSite=lax NID=511=EToBPqckCVRE7Paehug1PgNBKqe7lFLI9d12xJrGbvP9r8tkFIRWciry3gsy8FZ8OUIK4gE4PD-irgNzg4Y1fVePLdyu0AJdY_HcqL6zQYok-Adn-y5TDPmMCNuDnrouBfoxtqVjYY_4RFOe8jalkYto5fQAwzWnNJyw8K0avsw; expires=Sat, 27-Apr-2024 09:37:04 GMT; path=/; domain=.google.com; HttpOnly] X-Frame-Options:[SAMEORIGIN] X-Xss-Protection:[0]] 0xc000197920 -1 [] false true map[] 0xc0000ee000 0xc0000d8420}

200 OK

Status Code: 200
</code></pre>
<p>In the above code, we have defined a URL string as <code>reqURL</code> and used the <a href="https://pkg.go.dev/net/http#Get">Get</a> method to send a GET request. The <code>Get</code> is parsed with the <code>reqURL</code> string and the return values are stored as <code>resp</code> and <code>err</code>. We have added an error check after calling the <code>Get</code> method in order to avoid errors later in the code.</p>
<p>The <code>Get</code> method as seen from the output has returned a <code>*http.Response</code> object, we can use the <code>Status</code> and <code>StatusCode</code> properties to get the status code of the response. In this case, the status code of the response was <code>200</code>. The response <code>resp</code> is an object of type <code>http.Response</code> i.e. it has fields like <code>Body</code>, <code>StatusCode</code>, <code>Headers</code>, <code>Proto</code>, etc. We can get each individual field from the <code>resp</code> object. We will later look into how to read the <code>Body</code> field from the response, it is not directly read as a string nor it is stored in other forms, rather it is streamed from the requested URL.</p>
<h2>Creating a GET request</h2>
<p>We can even construct a GET request using the <a href="https://pkg.go.dev/net/http#NewRequest">NewRequest</a> method. This is a low-level way of creating a <code>GET</code> request. We mention the <code>method</code>, <code>URL</code>, and the <code>body</code>, in the case of <code>GET</code> request, there is nobody. So, the <code>NewRequest</code> is a general way of creating a <code>http</code> request.</p>
<pre><code class="language-go">// web/methods/get/newreq.go

package main

import (
	&quot;fmt&quot;
	&quot;net/http&quot;
)

func main() {
	reqURL := &quot;https://www.google.com&quot;
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
</code></pre>
<p>As we can see, we construct a <code>GET</code> request using the <code>NewRequest</code> method and then use the <a href="https://pkg.go.dev/net/http#Client.Do">Do</a> method to send the request to the server. The <a href="https://pkg.go.dev/net/http#DefaultClient">http.DefaultClient</a> is used as a client to send the request, if we want to customize this we can create a new instance object of <a href="https://pkg.go.dev/net/http#Client">http.Client</a> and use it to send requests. We will be taking a look at clients in another part of this series when we want to persist a connection or avoid connecting multiple times to the same application/URL.</p>
<p>For now, we will go ahead with the DefaultClient. This will trigger the request, in this case, a <code>GET</code> request to the specified URL in the <code>reqURL</code> string. The <code>Do</code> method returns either a <code>http.Response</code> or an <code>error</code> just like the <code>Get</code> method did.</p>
<h2>Reading the Response Body</h2>
<p>We saw some different ways to send a <code>GET</code> request, now the below example will demonstrate how to read the body of the response. The response body is read from a buffer rather than loading the entire response into memory. It makes it flexible to parse the data efficiently and as per the needs. We will see how we use the <a href="https://pkg.go.dev/io">io</a> package's <a href="https://pkg.go.dev/io#ReadAll">ReadAll</a> method can be used to read from the response buffer.</p>
<pre><code class="language-go">// web/methods/get/body.go

package main

import (
	&quot;fmt&quot;
	&quot;io&quot;
	&quot;net/http&quot;
)

func main() {
	reqURL := &quot;https://httpbin.org/html&quot;
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
</code></pre>
<p>In the above example, we are trying to get the body from the response to the request sent at <a href="https://httpbin.org/html"><code>https://httpbin.org/html</code></a>. We have used the simple <code>Get</code> method instead of <code>NewRequest</code> and <code>Do</code> for simplicity. The response is stored in <code>resp</code>, we also have added <code>defer resp.Body.Close()</code> which is to say that we will close the body reader object when the function is returned/closed. So, this means that the <code>Body</code> is not readily available data, we need to obtain/stream the data from the server. We have to receive the body in bytes as a tcp request, the body is streamed in a buffer.</p>
<p>The response body is streamed from the server, which means that it's not immediately available as a complete data set. We read the body in bytes as it arrives over the network, and it's stored in a buffer, which allows us to process the data efficiently.</p>
<h3>Reading Body in bytes</h3>
<p>We can even read the body in bytes i.e. by reading a chunk of the buffer at a time. We can use the <a href="https://pkg.go.dev/bytes#Buffer">bytes.Buffer</a> container object to store the body. Then we can create a slice of bytes as <code>[]bytes</code> of a certain size and read the body into the chunk. By writing the chunk into the buffer, we get the entire body from the response.</p>
<pre><code class="language-go">// web/methods/get/body.go


package main

import (
	&quot;bytes&quot;
	&quot;fmt&quot;
	&quot;net/http&quot;
)

func main() {
	reqURL := &quot;https://httpbin.org/html&quot;
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
		fmt.Printf(&quot;%s
&quot;, chunk[:n])
	}

    // entire body stored in bytes
	fmt.Println(buf.String())
}
</code></pre>
<p>In the above example, the body is read chunkwise buffers and obtained as a slice of bytes. We define the request as a <code>GET</code> request to the <a href="http://httpbin.org/html"><code>httpbin.org/html</code></a>. We create a new Buffer as a slice of bytes with <a href="https://pkg.go.dev/bytes#Buffer">bytes.Buffer</a>, then we define chunk as a container to stream the response body with a particular size. We have taken <code>1024</code> bytes as the size of the chunk. Then inside an infinite for loop, we read the body as <code>n, err :=</code> <a href="http://resp.Body.Read"><code>resp.Body.Read</code></a><code>(chunk)</code>. The code will read the body into the chunk(slice of byte) and the return value will be the size of the bytes read or an error. Then we check if there is no error, and if there is an error, we break the loop indicating we have completed reading the entire body or something went wrong. Then we write the chunk into the buffer that we allocated earlier as <code>buf</code>. This is a slice of bytes, we are basically populating the buffer with more slices of bytes.</p>
<p>The entire body is then stored in the buffer as a slice of bytes. So, we have to cast it into a string to see the contents. So, this is how we can read the contents of a body in a response in chunks.</p>
<h3>Parsing the JSON body with structs</h3>
<p>If we have the structure of the response body already decided, then we can define a struct for the response body and then we can <a href="https://doc.akka.io/docs/akka-http/current/common/unmarshalling.html#unmarshalling:~:text=Unmarshalling,type%20T.">Unmarshal</a> / deserialize/unpickle. This means we can convert the bytes representation of the data into a Golang-specific structure which is called a high-level representation of the data. We can parse the JSON body into a defined struct using <a href="https://pkg.go.dev/encoding/json#Unmarshal">Unmarshal</a> or <a href="https://pkg.go.dev/encoding/json#Decoder.Decode">Decode</a> methods in the <a href="https://pkg.go.dev/encoding/json">json</a> package.</p>
<p>Let's look at both the methods.</p>
<h4>Using Unmarshal</h4>
<p>The <code>Unmarshal</code> method takes in two parameters i.e. the body in bytes and the reference of the object that we want to unmarshal into. The method returns an error if there is a discrepancy in the returned JSON or the structure defined it is unable to deserialize the JSON object into the defined structure.</p>
<pre><code class="language-go">package main

import (
	&quot;encoding/json&quot;
	&quot;fmt&quot;
	&quot;io&quot;
	&quot;net/http&quot;
)

type Product struct {
	ID                 int      `json:&quot;id&quot;`
	Title              string   `json:&quot;title&quot;`
	Description        string   `json:&quot;description&quot;`
	Price              float64  `json:&quot;price&quot;`
	DiscountPercentage float64  `json:&quot;discountPercentage&quot;`
	Rating             float64  `json:&quot;rating&quot;`
	Stock              int      `json:&quot;stock&quot;`
	Brand              string   `json:&quot;brand&quot;`
	Category           string   `json:&quot;category&quot;`
    Thumbnail          string   `json:&quot;thumbnail,omitempty&quot;`
    Images             []string `json:&quot;-&quot;`
}

func main() {
	reqURL := &quot;https://dummyjson.com/products/1&quot;
	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var data Product
	if err := json.Unmarshal(body, &amp;data); err != nil {
		panic(err)
	}

	fmt.Println(data)
	fmt.Println(data.Title)
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

{1 iPhone 9 An apple mobile which is nothing like apple 549 12.96 4.69 94 Apple smartphones https://cdn.dummyjson.com/product-images/1/thumbnail.jpg []}
iPhone 9
</code></pre>
<p>In the above example, we have defined a structure called Product with fields such as <code>ID</code>, <code>Title</code>, <code>Description</code>, etc. We use the JSON tag to specify how each field should be encoded to or decoded from JSON. These tags guide the Golang JSON encoders and decoders to correctly map JSON data to struct fields and vice versa. The <code>omitempty</code> option in a struct tag instructs the encoder to omit the field from the JSON output if the field's value is the zero value for its type (e.g., 0 for integers, &quot;&quot; for strings, nil for pointers, slices, and maps). This is useful for reducing the size of the JSON output by excluding empty or default-valued fields.</p>
<p>Conversely, the <code>-</code> option in a struct tag tells the encoder and decoder to completely ignore the field. It will not be included in encoded JSON, nor will it be populated when decoding JSON into a struct. This is particularly useful for excluding fields that are meant for internal use only and should not be exposed through JSON.</p>
<p>Therefore, <code>omitempty</code> is used to control the inclusion of fields in the JSON output based on their values, while <code>-</code> is used to exclude fields from both encoding and decoding from JSON.</p>
<p>We send the <code>GET</code> request to the api <code>https://dummyjson.com/products/1</code>. The response from the request is read into a slice of bytes with <a href="https://pkg.go.dev/io#ReadAll">io.ReadAll</a> that takes in a <a href="https://pkg.go.dev/io#Reader">io.Reader</a> object in this case it is the <code>resp.Body</code> and it returns a slice of byte and error if any issue arises while reading in the body. Further, we can use the <a href="https://pkg.go.dev/encoding/json#Unmarshal">Unmarshal</a> method to parse the slice of body <code>body</code> into the struct <code>Product</code> with the variable <code>data</code>, the reference to <code>&amp;data</code> indicates that the method will directly mutate/change this variable and populate the object with the fields from the body.</p>
<p>So in a gist, to convert the JSON body into a golang native structure with <code>Unmarshal</code> with the following steps:</p>
<ul>
<li>Read the body into a slice of bytes using <code>io.ReadAll</code></li>
<li>Create an object of the struct</li>
<li>Pass the body as a slice of bytes and the reference of that object (struct instance) into the Unmarshal method</li>
<li>Access the object with the fields in the struct</li>
</ul>
<p>In the output response, we can see the object is populated with the fields from the body. The <code>Title</code> field is accessed using the <code>data.Title</code> as we do with a normal golang struct. The <code>Images</code> field is not populated because we have always ignored/omitted from the json tag with <code>-</code>.</p>
<h4>Using Decoder</h4>
<p>Similar to the <code>Unmarshal</code> we can use the <a href="https://pkg.go.dev/encoding/json#Decoder">Decoder</a> to parse the body into a struct. However, the parameters it takes are a bit different and it is a two-step process. We first create a <a href="https://pkg.go.dev/encoding/json#Decoder">Decoder</a> object using the <a href="https://pkg.go.dev/encoding/json#NewDecoder">NewDecoder</a> method, which takes in a <code>io.Reader</code> object, luckily the body from the response is already in that structure, so we can directly pass that <code>resp.Body</code> into the <code>NewDecoder</code> method. The second step is to decode the data into the object, here as well, we need to create the object of the struct and parse the reference to the object to the <a href="https://pkg.go.dev/encoding/json#Decoder.Decode">Decode</a> method. The <code>Decode</code> method converts the bytes parsed in the <code>resp.Body</code> from the <code>Decoder</code> object and populates the fields of the object provided in the reference struct.</p>
<p>So the steps for deserializing the json object into the struct with the decode method are:</p>
<ul>
<li>Create a decoder with <code>NewDecoder</code> method and pass the <code>resp.Body</code> as the parameter which is an <code>io.Reader</code> object</li>
<li>Create an object of the struct</li>
<li>Decode the body into that object using the <code>decoder.Decode</code> method</li>
<li>Access the object with the fields in the struct</li>
</ul>
<pre><code class="language-go">package main

import (
	&quot;encoding/json&quot;
	&quot;fmt&quot;
	&quot;net/http&quot;
)

type Product struct {
	ID                 int     `json:&quot;id&quot;`
	Title              string  `json:&quot;title&quot;`
	Description        string  `json:&quot;description&quot;`
	Price              float64 `json:&quot;price&quot;`
	DiscountPercentage float64 `json:&quot;discountPercentage&quot;`
	Rating             float64 `json:&quot;rating&quot;`
	Stock              int     `json:&quot;stock&quot;`
	Brand              string  `json:&quot;brand&quot;`
	Category           string  `json:&quot;category&quot;`
    Thumbnail          string   `json:&quot;thumbnail,omitempty&quot;`
    Images             []string `json:&quot;-&quot;`
}

func main() {
	reqURL := &quot;https://dummyjson.com/products/1&quot;
	resp, err := http.Get(reqURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data Product
	decoder := json.NewDecoder(resp.Body)
    err = decoder.Decode(&amp;data)
    if err != nil {
        panic(err)
    }

	fmt.Println(data)
	fmt.Println(data.Title)
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

{1 iPhone 9 An apple mobile which is nothing like apple 549 12.96 4.69 94 Apple smartphones https://cdn.dummyjson.com/product-images/1/thumbnail.jpg []}
iPhone 9
</code></pre>
<p>We have first defined the struct <code>Product</code> with the <code>json:&quot;id&quot;</code> tag. As explained earlier, we have used the json tags to identify the fields from the json data to the structures while encoding and decoding.
In the above example, we have sent a <code>GET</code> request to the api endpoint <code>https://dummyjson.com/products/1</code>, and we have created a new decoder with the <code>NewDecoder</code> method with the <code>resp.Body</code> as the parameter. The <code>data</code> is created as a <code>Product</code> instance. The reference to <code>data</code> is parsed to the <code>Decode</code> method from the <code>decoder</code> instance as <code>&amp;data</code>. This method will either return <code>nil</code> or an <code>error</code>. Thereafter, we can check for errors and then only access the data object with its populated fields from the response body.</p>
<p>There is a certain difference between the <code>Unmarshal</code> and <code>Decode</code> methods. The difference is just a slight performance improvement in the <code>NewDecoder</code> and <code>Decode</code> methods. Though it is not significant, having a little info about it might be handy in your use case. Read here for more info : <a href="https://dev.to/jpoly1219/to-unmarshal-or-to-decode-json-processing-in-go-explained-5870">To Unmarshal or Decode</a></p>
<h2>Adding Headers to a GET Request</h2>
<p>We can even add headers before sending a <code>GET</code> request to a URL. By creating a <code>Request</code> object with the <code>NewRequest</code> method and adding a <a href="https://pkg.go.dev/net/http#Header">Header</a> with the <a href="https://pkg.go.dev/net/http#Header.Add">Add</a> method. The <code>Add</code> method will take in two parameters i.e. the key of the header, and the value of the header both as strings.</p>
<pre><code class="language-go">// web/methods/get/header.go

package main

import (
	&quot;fmt&quot;
	&quot;io&quot;
	&quot;net/http&quot;
)

func main() {
	req, err := http.NewRequest(http.MethodGet, &quot;https://api.github.com/users/mr-destructive&quot;, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add(&quot;Authorization&quot;, &quot;token YOUR_TOKEN&quot;)
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
</code></pre>
<pre><code>$ go run web/methods/get/header.go

{&quot;message&quot;:&quot;Bad credentials&quot;,&quot;documentation_url&quot;:&quot;https://docs.github.com/rest&quot;}
</code></pre>
<p>In the above example, we have created a <code>GET</code> request to the <a href="https://api.github.com/users/mr-destructive"><code>https://api.github.com/users/mr-destructive</code></a> the last portion is the username, it could be any valid username. The request is to the GitHub API, so it might require API Key/Tokens in the headers, however, if there are certain endpoints that do not require Authorization headers might work just fine.</p>
<p>So, the above code will give <code>401</code> error indicating the request has wrong or invalid credentials, if we remove the header, the request will work fine. This is just an example, but headers are useful in working with APIs.</p>
<p>Without adding the header:</p>
<pre><code>$ go run web/methods/get/header.go

{&quot;login&quot;:&quot;Mr-Destructive&quot;,&quot;id&quot;:40317114,&quot;node_id&quot;:&quot;MDQ6VXNlcjQwMzE3MTE0&quot;,&quot;avatar_url&quot;:&quot;https://avatars.githubusercontent.com/u/40317114?v=4&quot;,&quot;gravatar_id&quot;:&quot;&quot;,&quot;url&quot;:&quot;https://api.github.com/users/Mr-Destructive&quot;,
... 
&quot;updated_at&quot;:&quot;2023-10-10T17:57:22Z&quot;}
</code></pre>
<p>That's it from the 33rd part of the series, all the source code for the examples are linked in the GitHub on the <a href="https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/get/">100 days of Golang</a> repository.</p>
<p><a href="https://github.com/Mr-Destructive/100-days-of-golang">100-days-of-golang</a></p>
<h2>References</h2>
<ul>
<li><a href="https://dev.to/jpoly1219/to-unmarshal-or-to-decode-json-processing-in-go-explained-5870">To Unmarshal or Decode</a></li>
<li><a href="https://drstearns.github.io/tutorials/gojson/">Golang JSON tutorial</a></li>
<li><a href="https://www.sohamkamani.com/golang/omitempty/">Golang OmitEmpty</a></li>
</ul>
<h2>Conclusion</h2>
<p>From this article, we explored the <code>GET</code> HTTP method in golang. By using a few examples for creating a get request, adding headers, reading response body, the basic use cases were demonstrated.</p>
<p>Hopefully, you found this section helpful, if you have any comments or feedback, please let me know in the comments section or on my social handles. Thank you for reading. Happy Coding :)</p>
