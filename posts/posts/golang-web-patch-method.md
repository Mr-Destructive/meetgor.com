{
  "title": "Golang Web: PATCH Method",
  "status": "published",
  "slug": "golang-web-patch-method",
  "date": "2025-04-05T12:33:17Z"
}

<h2>Introduction</h2>
<p>In previous sections of this series, we've covered the GET, POST, and PUT methods. Now, we will explore the PATCH method, which differs from the others in several key ways. The PATCH method is somewhat more flexible and depends on how the server or API you're working with is designed.</p>
<p>In this section, we'll focus on understanding what the PATCH method is and how to use it. While we will dive deeper into building and structuring a full CRUD API later in the series, the focus here will be on the what and why of the PATCH method, not the how.</p>
<h2>What is the PATCH Method?</h2>
<p>The PATCH method is often compared to the PUT method, but with one important distinction: PATCH is used to perform partial updates on a resource. Unlike PUT, which typically requires you to send the entire resource to update it, PATCH allows you to send only the fields that need to be updated. This makes it a more efficient option when updating a subset of a resource.</p>
<p>In a PATCH request, the body usually contains instructions in a format like JSON, which specifies the fields to update. These instructions define the changes to be applied to the resource. For example, you may only want to change one field of a user's profile, such as their email address, while leaving the rest of the data untouched.</p>
<h2>PATCH vs. PUT</h2>
<p>Key Differences While both PATCH and PUT are used to modify resources, there are significant differences in their behavior:</p>
<ul>
<li>
<p>PUT replaces the entire resource. When you send a PUT request, you must include the full representation of the resource, even if you're only changing a small part of it.</p>
</li>
<li>
<p>PATCH, on the other hand, is for partial updates. You only need to include the fields that are changing, not the entire resource.</p>
</li>
</ul>
<p>If the update involves more fields than just the ones you're changing, PUT may be the better choice. However, the scope of this article is to focus solely on the PATCH method.</p>
<h2>How Does a PATCH Request Work?</h2>
<p>In the simplest terms, a PATCH request allows you to perform a partial update on a resource. It is similar to a PUT request, but more specific in how it updates the resource. According to the HTTP specification, an ideal PATCH request should:</p>
<ul>
<li>
<p>Accept a &quot;patch document&quot; in the request body, which contains the list of operations to apply (e.g., &quot;replace&quot;, &quot;add&quot;, &quot;remove&quot;).</p>
</li>
<li>
<p>Apply these updates to the target resource.</p>
</li>
<li>
<p>If the update cannot be applied correctly, the operation should fail without applying any of the changes.</p>
</li>
</ul>
<p>This ensures that no partial or inconsistent updates are left behind</p>
<p>For example, if you're updating a user's email address and something goes wrong in the middle of the operation, the PATCH request should ensure that the email isn't updated partially. If there’s an error, none of the updates should be applied, ensuring data consistency.</p>
<p>Also, the patch method is not idempotent, meaning that if you send the same input/request, it need not necessarily return the same output. Because we are not sending the actual original entity, we are sending the partial data fields that need to be updated, so it might update the original entity on subsequent requests since there is no original request sent in the request; it only identifies the resource from the URI and fields to update in the request body.</p>
<p>Now, let’s sum up the patch request in a few words</p>
<ul>
<li>
<p>Updates specific fields mentioned in the patch document</p>
</li>
<li>
<p>Can be partial (only the fields that need to be updated are sent, unlike PUT, which typically replaces the entire resource)</p>
</li>
<li>
<p>Not necessarily idempotent (depends on the implementation)</p>
</li>
<li>
<p>Not Safe (since resources will be updated on the server side)</p>
</li>
</ul>
<h2>Basic PATCH request</h2>
<p>Let’s start with the basic PATCH request that we can create in Golang. The [net/http]() package will be used to construct the request, and will be using <code>encoding/json</code> and some other utilities for string and byte parsing.</p>
<p>So, first we will construct a HTTP request using the [http.NewRequest]() with the parameters like the http method to use, the URL to hit, and the request body if any. We will then need to send the json body which would consist of the fields to be updated.</p>
<h3>Defining the API/Server Endpoint URL</h3>
<p>We would need to define the endpoint that we are hitting, we can directly use the API URL or we can construct the API URL on the fly depending on the id, and parameter that will be dynamic. As PATCH request, usually modify a particular entity, we would generally have some form of identifier for that entity/object on the database, etc. So in this case, it is <code>id</code> of the post, so, we can pass the post in the URL.</p>
<pre><code class="language-go">// define URL to hit the API
apiURL := &quot;https://jsonplaceholder.typicode.com/posts/4&quot;

// OR
// baseURL := &quot;https://jsonplaceholder.typicode.com&quot;
// postId := 4
// postURL := fmt.Sprintf(&quot;%s/posts/%d&quot;, baseURL, postId)
</code></pre>
<p>We can either directly define the URL or dynamically construct the URL, that is quite straightforward. The later one is the one we usually use and design.</p>
<h3>Constructing the JSON Body</h3>
<p>This section is a little dependent on the context as you might have a direct json string that you can directly pass to the API or you might have a golang object that you need to Marshal in order to convert that object into string/bytes.</p>
<ol>
<li>
<p>Direct JSON String</p>
<p>So, there is nothing to do here, since the object is already in the form of a json string.</p>
<pre><code class="language-go">reqBody := `{&quot;body&quot;: &quot;new body&quot;}`
</code></pre>
<p>However, if you have certain fields that you need to exclude or omit, you have to construct a struct and then marshal it</p>
</li>
<li>
<p>Marshalling (converting object into bytes/string)</p>
<p>We need to convert the Golang native object into some form of a json string or bytes that could be sent over the network. That process is called <a href="https://en.wikipedia.org/wiki/Marshalling_(computer_science)">marshalling</a> or serialisation.</p>
</li>
</ol>
<pre><code class="language-go">type Post struct {
	ID     int    `json:&quot;id,omitempty&quot;`
	Title  string `json:&quot;title,omitempty&quot;`
	Body   string `json:&quot;body,omitempty&quot;`
	UserId int    `json:&quot;userId,omitempty&quot;`
}

userObj := Post{
    Body: &quot;New Body&quot;,
}

var reqBody []byte
reqBody, err := json.Marshal(userObj)
if err != nil {
	log.Fatal(err)
}

log.Println(&quot;New body:&quot;, string(reqBody))
// New body: {&quot;body&quot;:&quot;New Body&quot;}
</code></pre>
<p>In the above snippet, we have defined a <code>Post</code> struct with the fields like <code>ID</code>, <code>Title</code>, <code>Body</code>, <code>UserID</code> , and those have <code>omitempty</code> tag along with the json fields that we want to marshal into. The omitempty will omit or ignore the fields if they are empty or not present in the object/instance of this structure. So in the example, <code>userObj</code> is the instance of the <code>Post</code> struct and it only has the <code>Body</code> populated, so the reqBody will only contain one field <code>body</code> in the json representation. The [json.Marshal]() is the function that we use to convert the object/instance into a byte form.</p>
<p>This <code>reqBody</code> will serve as the request body for the request that will be a <code>PATCH</code> method to the mentioned endpoint / API URL.</p>
<h3>Constructing the HTTP PATCH Request</h3>
<p>Now, we have the parts that we need to construct the request, we can combine the parts and hit the endpoint. However, it is a bit different compared to <code>GET</code> and <code>POST</code> request that we do in Golang, the HTTP package has built in methods for the <code>GET</code> and <code>POST</code> methods, however for methods like <code>PUT</code>, <code>PATCH</code>, <code>DELETE</code> and others, we need to construct a <a href="https://pkg.go.dev/net/http#Request">Request</a> object and then send that request.</p>
<pre><code class="language-go">req, err := http.NewRequest(&quot;PATCH&quot;, postURL, strings.NewReader(reqBody))
if err != nil {
	log.Fatal(err)
}
req.Header.Set(&quot;Content-Type&quot;, &quot;application/json&quot;)

// in case of wired utf-8 characters appear in the body
//req.Header.Set(&quot;Content-Type&quot;, &quot;application/json; charset=utf-8&quot;)
</code></pre>
<p>To do that, we call the <a href="https://pkg.go.dev/net/http#NewRequest">NewRequest</a> method with the parameters like the HTTP method, the URL, and the request Body all of which we have at the moment.</p>
<ul>
<li>
<p>The method is <code>PATCH</code></p>
</li>
<li>
<p>The URL is <code>postURL</code></p>
</li>
<li>
<p>The body is <code>strings.NewReader(reqBody)</code> as we need a <code>io.Reader</code> object instead of string or byte slice</p>
</li>
</ul>
<p>So, once we have that, we would also set the <code>Header</code> with the field of <code>Content-Type</code> and the value as <code>application/json</code> since the request body has json representation of the patch document that will be sent.</p>
<h3>Sending the Request</h3>
<p>Once, the <code>req</code> object is created, we also need a <a href="https://pkg.go.dev/net/http#Client">Client</a> to send the request, so we create the client as default http.Client object with defaults and call the <a href="https://pkg.go.dev/net/http#Client.Do">Do</a> method with the <code>req</code> as the request parameter in order to send the request with the default client.</p>
<p>This method returns the response object, and an error if any.</p>
<p>We also add the <code>defer resp.Body.Close()</code> in order to avoid leaks and safely access the response body.</p>
<pre><code class="language-go">client := &amp;http.Client{}
resp, err := client.Do(req)
if err != nil {
	log.Fatal(err)
}
defer resp.Body.Close()
</code></pre>
<p>At this point, we can now start consuming the response and use it for further processing as per the needs.</p>
<h3>Unmarshalling the Response</h3>
<p>We first read the response into a string or byte representation using the io.ReadAll method and then use the json.Unmarshal to convert the bytes into golang object/instance.</p>
<pre><code class="language-go">var updatedPost Post

respBody, err := io.ReadAll(resp.Body)
if err != nil {
	log.Fatal(err)
}
    
// convert the response json bytes to Post object in golang
err = json.Unmarshal(respBody, &amp;updatedPost)
if err != nil {
	log.Fatal(err)
}

fmt.Println(updatedPost)
fmt.Println(updatedPost.Title)
</code></pre>
<p>In the above example, we have read the response Body which can be accessed as the <code>Body</code> field in the <a href="https://pkg.go.dev/net/http#Response">Response</a> object via the <code>resp</code> variable. So, the function will return the <code>respBody</code> as a string or an error if any. Then using this string object, we can use the json.Unmarshal function to send this string and create this <code>updatedPost</code> object of Post struct. This method will mutate this object as we have passed it by reference indicated by <code>&amp;updatedPost</code> . So, this will do two things, one update / mutate the <code>updatedPost</code> instance from the <code>respBody</code> and give any error if any arrises during the <a href="https://developer.mozilla.org/en-US/docs/Glossary/Deserialization">deserialsation</a> of the response .</p>
<p>Now, we have the object in golang from the response bytes, we can use it as per requirements.</p>
<p>So, that is the example in the entirety.</p>
<p>Let’s simplify the steps which are similar to the POST/PUT method as well</p>
<ul>
<li>
<p>Define/construct URL</p>
</li>
<li>
<p>Marshal the object into JSON string as the request body</p>
</li>
<li>
<p>Construct the request object (method, URL and the body)</p>
</li>
<li>
<p>Send the request with the default client</p>
</li>
<li>
<p>Read the response and deserialise/unmarshall</p>
</li>
<li>
<p>Access the object in golang</p>
</li>
</ul>
<pre><code class="language-go">package main

import (
	&quot;encoding/json&quot;
	&quot;fmt&quot;
	&quot;io&quot;
	&quot;log&quot;
	&quot;net/http&quot;
	&quot;strings&quot;
)

type Post struct {
	ID     int    `json:&quot;id,omitempty&quot;`
	Title  string `json:&quot;title,omitempty&quot;`
	Body   string `json:&quot;body,omitempty&quot;`
	UserId int    `json:&quot;userId,omitempty&quot;`
}

func main() {

	// define URL to hit the API
    baseURL := &quot;https://jsonplaceholder.typicode.com&quot;
    postId := 4
    postURL := fmt.Sprintf(&quot;%s/posts/%d&quot;, baseURL, postId)

    // define the body -&gt; with the field to update
	reqBody := `{&quot;body&quot;: &quot;new body&quot;}`
	fmt.Println(&quot;New body:&quot;, reqBody)

    // send a new request, with the `PATCH` method, url and the body
	req, err := http.NewRequest(&quot;PATCH&quot;, postURL, strings.NewReader(reqBody))
	if err != nil {
		log.Fatal(err)
	}
    // set the header content type to json
	req.Header.Set(&quot;Content-Type&quot;, &quot;application/json&quot;)

	client := &amp;http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(&quot;Response status code:&quot;, resp.StatusCode)
	fmt.Println(&quot;Response Status:&quot;, resp.Status)
	
    var updatedPost Post

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
    
    // convert the response json bytes to Post object in golang
	err = json.Unmarshal(respBody, &amp;updatedPost)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updatedPost)
	fmt.Println(updatedPost.Title)

}
</code></pre>
<pre><code class="language-plaintext">New body: {&quot;body&quot;: &quot;new body&quot;}
Response status code: 200
Response Status: 200 OK
{4 eum et est occaecati new body 1}

title: eum et est occaecati                                                                         
body: new body
id: 4
user id: 1
</code></pre>
<p>As you would see, it has only updated the <code>body</code> and has not updated the other fields.</p>
<p>If you would have sent a similar body with a <code>PUT</code> method, the results would have been different. That would have been dependent on the implementation of the API of course, but if there is only a few fields in the request body for a PUT method, it would have replaced the values with the empry values which are not present in the request body.</p>
<p>That is the difference between a <code>PUT</code> and a <code>PATCH</code> method, the <code>PATCH</code> method, ideally should only update the fields of the entity which are mentioned in the request body, whereas the <code>PUT</code> method has to update the entire resource whether the fields are provided or not. Again, the implementation of these API on the server plays a vital role in how the behaviour defers and the method in itself would perform.</p>
<p>This is also called as <code>JSON Merge Patch</code></p>
<h2>JSON Merge PATCH</h2>
<p>The above API is implementing a <a href="https://datatracker.ietf.org/doc/html/rfc7386">Merge PATCH</a> which is to say, merge the changes in the actual entity.</p>
<p>Let’s say there is a Blog post Entity on a Server, you have a post that you are writing as an author. The post has a id of <code>4</code> let’s say and you are constantly changing the body of the post.</p>
<p>So, you don’t want to send the <code>title</code> or <code>author_id</code> or any field that is not changing from the client again and again while saving, so the <code>MERGE PATCH</code> endpoint will be helpful in that case, where the client only sends the required fields to be updated.</p>
<p>In this example, the user would only send the <code>body</code> of the post to the API every time it makes changes or saves the draft. In some cases, it might also want to change the title, so it will include the title, but not all the fields. The API will know it is a <code>PATCH</code> request and the content type is <code>json</code> so it will only change or update the fields that are provided in the request body to the actual entity in the database or wherever it is stored on the server.</p>
<p>So, that is what is the JSON Merge PATCH or Merge PATCH in general. JSON Merge PATCH is specific to the JSON based document APIs.</p>
<p>Below is the example, the same steps but a different endpoint. A user API that I have specifically created for demonstrating the difference in a PUT vs Merge PATCH vs JSON PATCH requests.</p>
<pre><code class="language-go">package main

import (
	&quot;encoding/json&quot;
	&quot;fmt&quot;
	&quot;io&quot;
	&quot;net/http&quot;
	&quot;strings&quot;
)

type User struct {
	ID    int    `json:&quot;id,omitempty&quot;`
	Name  string `json:&quot;name,omitempty&quot;`
	Email string `json:&quot;email,omitempty&quot;`
	Roles string `json:&quot;roles,omitempty&quot;`
}

func main() {
    baseURL := &quot;https://dummy-json-patch.netlify.app/.netlify/functions&quot;
    userID := 2
	apiURL := fmt.Sprintf(&quot;%s/users/?id=%d&quot;, baseURL, userID)
    
    userObj := User{
		Name:  &quot;dummy name&quot;,
		Roles: &quot;dummy role&quot;,
	}

	var jsonPatchBody []byte
	jsonPatchBody, err := json.Marshal(userObj)
	if err != nil {
		panic(err)
	}
    
    // OR directly define the json as string
	//jsonMergePatchBody := `{
    //    &quot;name&quot;: &quot;new dummy name&quot;,
    //    &quot;roles&quot;: &quot;new dummy role&quot;
    //}`

	req, err := http.NewRequest(&quot;PATCH&quot;, apiURL, strings.NewReader(jsonMergePatchBody))
	req.Header.Set(&quot;Content-Type&quot;, &quot;application/json&quot;)

	client := &amp;http.Client{}
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
	err = json.Unmarshal(resBody, &amp;user)
	if err != nil {
		panic(err)
	}

	fmt.Println(&quot;Updated/Patched User&quot;, user)
	fmt.Println(&quot;Id:&quot;, user.ID)
	fmt.Println(&quot;Name:&quot;, user.Name)
	fmt.Println(&quot;Email:&quot;, user.Email)
    fmt.Println(&quot;Bio:&quot;, user.Bio)
	fmt.Println(&quot;Roles:&quot;, user.Roles)
}
</code></pre>
<p>Original User with id 2</p>
<pre><code class="language-plaintext">{&quot;id&quot;:2,&quot;name&quot;:&quot;dummy&quot;,&quot;email&quot;:&quot;dummyyummy@user.com&quot;,&quot;bio&quot;:&quot;empty bio&quot;,&quot;roles&quot;:&quot;read&quot;}

Id: 2
Name: dummy
Email: dummyyummy@user.com
Bio: empty bio
Roles: read
</code></pre>
<p>Output of the program</p>
<pre><code class="language-plaintext">Request Body: {&quot;name&quot;:&quot;dummy name&quot;,&quot;roles&quot;:&quot;dummy role&quot;}

{&quot;id&quot;:2,&quot;name&quot;:&quot;dummy name&quot;,&quot;email&quot;:&quot;dummyyummy@user.com&quot;,&quot;bio&quot;:&quot;empty bio&quot;,&quot;roles&quot;:&quot;dummy role&quot;}

Updated/Patched User {2 dummy name dummyyummy@user.com empty bio dummy role}

Id: 2
Name: dummy name
Email: dummyyummy@user.com
Bio: empty bio
Roles: dummy role 
</code></pre>
<p>In the above example, the only fields that will be updated are <code>name</code> and <code>roles</code> , since the API is implemented to only update the fields provided in the json merge patch document (request body).</p>
<p>As you can see that, only the <code>name</code> and <code>roles</code> are changed. The name was <code>dummy</code> that changed to <code>dummy name</code> and role changed from <code>read</code> to <code>dummy role</code> .</p>
<p>Now, let’s see the same request but with PUT method on it.</p>
<p>Before we hit this API however, let’s note what the user with id 2 is</p>
<pre><code class="language-json">{
  &quot;id&quot;: 2,
  &quot;name&quot;: &quot;dummy name&quot;,
  &quot;email&quot;: &quot;dummyyummy@user.com&quot;,
  &quot;bio&quot;: &quot;empty bio&quot;,
  &quot;roles&quot;: &quot;dummy role&quot;
}
</code></pre>
<p>This is the result of our recent patch request. Now, we will send a PUT request to the same user with a different body.</p>
<pre><code class="language-go">package main

import (
	&quot;encoding/json&quot;
	&quot;fmt&quot;
	&quot;io&quot;
	&quot;net/http&quot;
	&quot;strings&quot;
)

type User struct {
	ID    int    `json:&quot;id,omitempty&quot;`
	Name  string `json:&quot;name,omitempty&quot;`
	Email string `json:&quot;email,omitempty&quot;`
	Bio   string `json:&quot;bio,omitempty&quot;`
	Roles string `json:&quot;roles,omitempt&quot;`
}

func main() {
    baseURL := &quot;https://dummy-json-patch.netlify.app/.netlify/functions&quot;
    userID := 2
	apiURL := fmt.Sprintf(&quot;%s/users/?id=%d&quot;, baseURL, userID)
    
	userObj := User{
		Name:  &quot;not a dummy name&quot;,
		Roles: &quot;not a dummy role&quot;,
	}

	var jsonPatchBody []byte
	jsonPatchBody, err := json.Marshal(userObj)
	if err != nil {
		panic(err)
	}
	fmt.Println(&quot;Request Body:&quot;, string(jsonPatchBody))

	//jsonPatchBody := `{
	//    &quot;name&quot;: &quot;dummy&quot;,
	//    &quot;roles&quot;: &quot;new dummy role&quot;
	//}`

	req, err := http.NewRequest(&quot;PUT&quot;, apiURL, strings.NewReader(string(jsonPatchBody)))
	req.Header.Set(&quot;Content-Type&quot;, &quot;application/merge-patch+json&quot;)

	client := &amp;http.Client{}
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
	err = json.Unmarshal(resBody, &amp;user)
	if err != nil {
		panic(err)
	}

	fmt.Println(&quot;Updated/Patched User&quot;, user)
	fmt.Println(&quot;Id:&quot;, user.ID)
	fmt.Println(&quot;Name:&quot;, user.Name)
	fmt.Println(&quot;Bio:&quot;, user.Bio)
	fmt.Println(&quot;Email:&quot;, user.Email)
	fmt.Println(&quot;Roles:&quot;, user.Roles)
}
</code></pre>
<p>Output:</p>
<pre><code class="language-plaintext">Request Body: {&quot;name&quot;:&quot;not a dummy name&quot;,&quot;roles&quot;:&quot;not a dummy role&quot;}

{&quot;id&quot;:2,&quot;name&quot;:&quot;not a dummy name&quot;,&quot;email&quot;:&quot;&quot;,&quot;bio&quot;:&quot;&quot;,&quot;roles&quot;:&quot;not a dummy role&quot;}

Updated/Patched User {2 not a dummy name   not a dummy role}

Id: 2
Name: not a dummy name
Bio:
Email:
Roles: not a dummy role
</code></pre>
<p>As you can see the <code>name</code> and <code>roles</code> are updated, however the <code>bio</code> and <code>email</code> fields are empty. Since we only said to update the <code>name</code> and <code>roles</code> fields, but it was a <code>PUT</code> request, it expects all the fields, and if any of the fields are missing, it will consider them as empty and update those as well.</p>
<p>So, the difference might be crystal clear now. When to use <code>PATCH</code> and when to avoid <code>PUT</code>.</p>
<ul>
<li>
<p>When you have a large set of updates, preference could be given to PUT</p>
</li>
<li>
<p>If you have very specific fields to update and a very limited fields PATCH is recommended</p>
</li>
</ul>
<p>There is other type of PATCH specifically designed for JSON APIs, or I should say JSON Documents APIs.</p>
<h2>JSON PATCH</h2>
<p>The <a href="https://datatracker.ietf.org/doc/html/rfc6902/">JSON PATCH</a> is a specification in which we can define what operations to perform on which fields, or path of the fields to replace, move or copy to.</p>
<blockquote>
<p>A JSON Patch document is a JSON document that represents an array of objects. Each object represents a single operation to be applied to the target JSON document.</p>
</blockquote>
<p>As it takes this operations, it applies them sequentially and hence it won’t replace all the fields for the entity, as that is the expected behavior of the PATCH method. In other words, it would only apply changes to the fields and the related fields provided in the json patch document (request body).</p>
<p>There are a few operations that you can perform with this json patch method, and provide the instructions accordingly for individual operations in the JSON PATCH document.</p>
<p>Operations</p>
<ul>
<li>
<p>add</p>
</li>
<li>
<p>remove</p>
</li>
<li>
<p>replace</p>
</li>
<li>
<p>move</p>
</li>
<li>
<p>copy</p>
</li>
<li>
<p>test</p>
</li>
</ul>
<p>So, for each of the operations, a high level definition can be considered as:</p>
<ul>
<li>
<p>To add a field you can specify the operation as <code>add</code> , the path as the field to be added, and the value as the actual value to be added</p>
</li>
<li>
<p>To remove a field, you can specify the operation as <code>remove</code> , and the path as the field to remove</p>
</li>
<li>
<p>To replace a field, you can specify the operation as <code>replace</code>, the path as the field to be updated/replaced, and the value of the actual value to be added</p>
</li>
<li>
<p>To move a field, you can specify the operation as <code>move</code>, the <strong>from</strong> as the field to be updated/moved from and the path to the field the from value should be moved to.</p>
</li>
<li>
<p>To copy a field, you can specify the operation as <code>copy</code>, the from as the field to updated/copied from and the path to the field to which the value should be copied to.</p>
</li>
<li>
<p>The test operation is a bit different as it is used for comparison of a <code>path</code> value to the actual value specified in the object. It might return true or false, but not actually return it might be used as a checkpoint for continuing with the operation in the document.</p>
</li>
</ul>
<p>In this example, we are creating a similar patch request, but using this json patch document kind of structure.</p>
<h3>Construct the json-patch document</h3>
<h3>Send the request</h3>
<pre><code class="language-go">package main

import (
	&quot;encoding/json&quot;
	&quot;fmt&quot;
	&quot;io&quot;
	&quot;net/http&quot;
	&quot;strings&quot;
)

type User struct {
	ID    int    `json:&quot;id&quot;`
	Name  string `json:&quot;name&quot;`
	Email string `json:&quot;email&quot;`
    Bio   string `json:&quot;bio&quot;`
	Roles string `json:&quot;roles&quot;`
}

func main() {
	baseURL := &quot;https://dummy-json-patch.netlify.app/.netlify/functions&quot;
    userID := 2
	apiURL := fmt.Sprintf(&quot;%s/users/?id=%d&quot;, baseURL, userID)

	jsonPatchBody := `[
        {
            &quot;op&quot;: &quot;replace&quot;,
            &quot;path&quot;: &quot;/name&quot;,
            &quot;value&quot;: &quot;new dummy name&quot;
        },
        {
            &quot;op&quot;: &quot;replace&quot;,
            &quot;path&quot;: &quot;/roles&quot;,
            &quot;value&quot;: &quot;new dummy role&quot;
        },
    ]`

	req, err := http.NewRequest(&quot;PATCH&quot;, apiURL, strings.NewReader(jsonPatchBody))
	req.Header.Set(&quot;Content-Type&quot;, &quot;application/json-patch+json&quot;)

	client := &amp;http.Client{}
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
	err = json.Unmarshal(resBody, &amp;user)
	if err != nil {
		panic(err)
	}

	fmt.Println(&quot;Updated/Patched User&quot;, user)
	fmt.Println(&quot;Id:&quot;, user.ID)
	fmt.Println(&quot;Name:&quot;, user.Name)
	fmt.Println(&quot;Bio:&quot;, user.Bio)
	fmt.Println(&quot;Email:&quot;, user.Email)
	fmt.Println(&quot;Roles:&quot;, user.Roles)
}
</code></pre>
<p>Original User (id=2) before the request</p>
<pre><code class="language-json">{
  &quot;id&quot;: 2,
  &quot;name&quot;: &quot;dummy name&quot;,
  &quot;email&quot;: &quot;dummyyummy@user.com&quot;,
  &quot;bio&quot;: &quot;empty bio&quot;,
  &quot;roles&quot;: &quot;dummy role&quot;
}
</code></pre>
<p>Output of the Program (JSON PATCH Request)</p>
<pre><code class="language-plaintext">{&quot;id&quot;:2,&quot;name&quot;:&quot;new dummy name&quot;,&quot;email&quot;:&quot;dummyyummy@user.com&quot;,&quot;bio&quot;: &quot;empty bio&quot;, &quot;roles&quot;:&quot;new dummy role&quot;}     

Updated/Patched User {2 new dummy name dummyyummy@user.com empty bio new dummy role} 

Id: 2
Name: new dummy name
Email: dummyyummy@user.com
Bio: empty bio
Roles: read
</code></pre>
<h3>References:</h3>
<ul>
<li>
<p><a href="https://en.wikipedia.org/wiki/HTTP#Request_methods">Wikipedia: HTTP Request Methods</a></p>
</li>
<li>
<p><a href="https://rubyonrails.org/2012/2/26/edge-rails-patch-is-the-new-primary-http-method-for-updates">Ruby on Rails: Patch is the new primary HTTP method for updates</a></p>
</li>
<li>
<p><a href="https://datatracker.ietf.org/doc/html/rfc5789">RFC 5789</a></p>
</li>
</ul>
