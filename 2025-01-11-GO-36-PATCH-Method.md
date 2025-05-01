---
article_html: "<h2>Introduction</h2>\n<p>In previous sections of this series, we've
  covered the GET, POST, and PUT methods. Now, we will explore the PATCH method, which
  differs from the others in several key ways. The PATCH method is somewhat more flexible
  and depends on how the server or API you're working with is designed.</p>\n<p>In
  this section, we'll focus on understanding what the PATCH method is and how to use
  it. While we will dive deeper into building and structuring a full CRUD API later
  in the series, the focus here will be on the what and why of the PATCH method, not
  the how.</p>\n<h2>What is the PATCH Method?</h2>\n<p>The PATCH method is often compared
  to the PUT method, but with one important distinction: PATCH is used to perform
  partial updates on a resource. Unlike PUT, which typically requires you to send
  the entire resource to update it, PATCH allows you to send only the fields that
  need to be updated. This makes it a more efficient option when updating a subset
  of a resource.</p>\n<p>In a PATCH request, the body usually contains instructions
  in a format like JSON, which specifies the fields to update. These instructions
  define the changes to be applied to the resource. For example, you may only want
  to change one field of a user's profile, such as their email address, while leaving
  the rest of the data untouched.</p>\n<h2>PATCH vs. PUT</h2>\n<p>Key Differences
  While both PATCH and PUT are used to modify resources, there are significant differences
  in their behavior:</p>\n<ul>\n<li>\n<p>PUT replaces the entire resource. When you
  send a PUT request, you must include the full representation of the resource, even
  if you're only changing a small part of it.</p>\n</li>\n<li>\n<p>PATCH, on the other
  hand, is for partial updates. You only need to include the fields that are changing,
  not the entire resource.</p>\n</li>\n</ul>\n<p>If the update involves more fields
  than just the ones you're changing, PUT may be the better choice. However, the scope
  of this article is to focus solely on the PATCH method.</p>\n<h2>How Does a PATCH
  Request Work?</h2>\n<p>In the simplest terms, a PATCH request allows you to perform
  a partial update on a resource. It is similar to a PUT request, but more specific
  in how it updates the resource. According to the HTTP specification, an ideal PATCH
  request should:</p>\n<ul>\n<li>\n<p>Accept a &quot;patch document&quot; in the request
  body, which contains the list of operations to apply (e.g., &quot;replace&quot;,
  &quot;add&quot;, &quot;remove&quot;).</p>\n</li>\n<li>\n<p>Apply these updates to
  the target resource.</p>\n</li>\n<li>\n<p>If the update cannot be applied correctly,
  the operation should fail without applying any of the changes.</p>\n</li>\n</ul>\n<p>This
  ensures that no partial or inconsistent updates are left behind</p>\n<p>For example,
  if you're updating a user's email address and something goes wrong in the middle
  of the operation, the PATCH request should ensure that the email isn't updated partially.
  If there’s an error, none of the updates should be applied, ensuring data consistency.</p>\n<p>Also,
  the patch method is not idempotent, meaning that if you send the same input/request,
  it need not necessarily return the same output. Because we are not sending the actual
  original entity, we are sending the partial data fields that need to be updated,
  so it might update the original entity on subsequent requests since there is no
  original request sent in the request; it only identifies the resource from the URI
  and fields to update in the request body.</p>\n<p>Now, let’s sum up the patch request
  in a few words</p>\n<ul>\n<li>\n<p>Updates specific fields mentioned in the patch
  document</p>\n</li>\n<li>\n<p>Can be partial (only the fields that need to be updated
  are sent, unlike PUT, which typically replaces the entire resource)</p>\n</li>\n<li>\n<p>Not
  necessarily idempotent (depends on the implementation)</p>\n</li>\n<li>\n<p>Not
  Safe (since resources will be updated on the server side)</p>\n</li>\n</ul>\n<h2>Basic
  PATCH request</h2>\n<p>Let’s start with the basic PATCH request that we can create
  in Golang. The <a href=\"https://pkg.go.dev/net/http\">net/http</a> package will
  be used to construct the request and will be using <code>encoding/json</code> and
  some other utilities for string and byte parsing.</p>\n<p>So, first, we will construct
  an HTTP request using the <a href=\"https://pkg.go.dev/net/http#NewRequest\">http.NewRequest</a>
  with the parameters like the http method to use, the URL to hit, and the request
  body if any. We will then need to send the json body which would consist of the
  fields to be updated.</p>\n<h3>Defining the API/Server Endpoint URL</h3>\n<p>We
  would need to define the endpoint that we are hitting, we can directly use the API
  URL or we can construct the API URL on the fly depending on the ID, and parameter
  that will be dynamic. As PATCH requests, usually modify a particular entity, we
  would generally have some form of identifier for that entity/object on the database,
  etc. So in this case, it is <code>id</code> of the post, so, we can pass the post
  in the URL.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">// define URL to hit the API</span><span class=\"w\"></span>\n<span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://jsonplaceholder.typicode.com/posts/4&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"c1\">// OR</span><span class=\"w\"></span>\n<span
  class=\"c1\">// baseURL := &quot;https://jsonplaceholder.typicode.com&quot;</span><span
  class=\"w\"></span>\n<span class=\"c1\">// postId := 4</span><span class=\"w\"></span>\n<span
  class=\"c1\">// postURL := fmt.Sprintf(&quot;%s/posts/%d&quot;, baseURL, postId)</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>We can either directly define
  the URL or dynamically construct the URL, that is quite straightforward. The latter
  one is the one we usually use and design.</p>\n<h3>Constructing the JSON Body</h3>\n<p>This
  section is a little dependent on the context as you might have a direct json string
  that you can directly pass to the API or you might have a golang object that you
  need to Marshal in order to convert that object into string/bytes.</p>\n<ol>\n<li>\n<p>Direct
  JSON String</p>\n<p>So, there is nothing to do here, since the object is already
  in the form of a json string.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nx\">reqBody</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">`{&quot;body&quot;: &quot;new body&quot;}`</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>However, if you have certain fields
  that you need to exclude or omit, you have to construct a struct and then marshal
  it</p>\n</li>\n<li>\n<p>Marshaling (converting object into bytes/string)</p>\n<p>We
  need to convert the Golang native object into some form of a json string or bytes
  that could be sent over the network. That process is called <a href=\"https://en.wikipedia.org/wiki/Marshalling_(computer_science)\">marshaling</a>
  or serialization.</p>\n</li>\n</ol>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kd\">type</span><span class=\"w\"> </span><span class=\"nx\">Post</span><span
  class=\"w\"> </span><span class=\"kd\">struct</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">     </span><span class=\"kt\">int</span><span
  class=\"w\">    </span><span class=\"s\">`json:&quot;id,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Title</span><span
  class=\"w\">  </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;title,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Body</span><span class=\"w\">   </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;body,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">UserId</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\">    </span><span
  class=\"s\">`json:&quot;userId,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"nx\">userObj</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">Post</span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">Body</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s\">&quot;New Body&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">reqBody</span><span class=\"w\"> </span><span
  class=\"p\">[]</span><span class=\"kt\">byte</span><span class=\"w\"></span>\n<span
  class=\"nx\">reqBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Marshal</span><span class=\"p\">(</span><span class=\"nx\">userObj</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;New body:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">reqBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"c1\">// New body: {&quot;body&quot;:&quot;New Body&quot;}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above snippet, we have defined a <code>Post</code> struct with the fields like
  <code>ID</code>, <code>Title</code>, <code>Body</code>, <code>UserID</code> , and
  those who have <code>omitempty</code> tag along with the json fields that we want
  to marshal into. The omit empty will omit or ignore the fields if they are empty
  or not present in the object/instance of this structure. So in the example, <code>userObj</code>
  is the instance of the <code>Post</code> struct and it only has the <code>Body</code>
  populated, so the reqBody will only contain one field <code>body</code> in the json
  representation. The <a href=\"https://pkg.go.dev/encoding/json#Marshal\">json.Marshal</a>
  is the function that we use to convert the object/instance into a byte form.</p>\n<p>This
  <code>reqBody</code> will serve as the request body for the request that will be
  a <code>PATCH</code> method to the mentioned endpoint / API URL.</p>\n<h3>Constructing
  the HTTP PATCH Request</h3>\n<p>Now, we have the parts that we need to construct
  the request, we can combine the parts and hit the endpoint. However, it is a bit
  different compared to <code>GET</code> and <code>POST</code> request that we do
  in Golang, the HTTP package has built-in methods for the <code>GET</code> and <code>POST</code>
  methods, however for methods like <code>PUT</code>, <code>PATCH</code>, <code>DELETE</code>
  and others, we need to construct a <a href=\"https://pkg.go.dev/net/http#Request\">Request</a>
  object and then send that request.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"s\">&quot;PATCH&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">postURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">strings</span><span
  class=\"p\">.</span><span class=\"nx\">NewReader</span><span class=\"p\">(</span><span
  class=\"nx\">reqBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"nx\">req</span><span class=\"p\">.</span><span class=\"nx\">Header</span><span
  class=\"p\">.</span><span class=\"nx\">Set</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Content-Type&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;application/json&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"c1\">// in case of wired utf-8 characters appear
  in the body</span><span class=\"w\"></span>\n<span class=\"c1\">//req.Header.Set(&quot;Content-Type&quot;,
  &quot;application/json; charset=utf-8&quot;)</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>To
  do that, we call the <a href=\"https://pkg.go.dev/net/http#NewRequest\">NewRequest</a>
  method with the parameters like the HTTP method, the URL, and the request Body all
  of which we have at the moment.</p>\n<ul>\n<li>\n<p>The method is <code>PATCH</code></p>\n</li>\n<li>\n<p>The
  URL is <code>postURL</code></p>\n</li>\n<li>\n<p>The body is <code>strings.NewReader(reqBody)</code>
  as we need a <code>io.Reader</code> object instead of string or byte slice</p>\n</li>\n</ul>\n<p>So,
  once we have that, we will also set the <code>Header</code> In the field of <code>Content-Type</code>
  and the value as <code>application/json</code> since the request body has json representation
  of the patch document that will be sent.</p>\n<h3>Sending the Request</h3>\n<p>Once,
  the <code>req</code> the object is created, we also need a <a href=\"https://pkg.go.dev/net/http#Client\">Client</a>
  to send the request, so we create the client as default http.Client object with
  defaults and call the <a href=\"https://pkg.go.dev/net/http#Client.Do\">Do</a> method
  with the <code>req</code> as the request parameter in order to send the request
  with the default client.</p>\n<p>This method returns the response object and an
  error if any.</p>\n<p>We also add the <code>defer resp.Body.Close()</code> in order
  to avoid leaks and safely access the response body.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"nx\">resp</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>At
  this point, we can now start consuming the response and use it for further processing
  as per the needs.</p>\n<h3>Unmarshalling the Response</h3>\n<p>We first read the
  response into a string or byte representation using the io.ReadAll method and then
  use the json.Unmarshal to convert the bytes into golang object/instance.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">updatedPost</span><span
  class=\"w\"> </span><span class=\"nx\">Post</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">io</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span>\n<span class=\"c1\">// convert the response json bytes to
  Post object in golang</span><span class=\"w\"></span>\n<span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">updatedPost</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">updatedPost</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">updatedPost</span><span
  class=\"p\">.</span><span class=\"nx\">Title</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In the above example, we have
  read the response Body which can be accessed as the <code>Body</code> field in the
  <a href=\"https://pkg.go.dev/net/http#Response\">Response</a> object via the <code>resp</code>
  variable. So, the function will return the <code>respBody</code> as a string or
  an error if any. Then using this string object, we can use the <a href=\"https://pkg.go.dev/encoding/json#Unmarshal\">json.Unmarshal</a>
  function to send this string and create this <code>updatedPost</code> object of
  Post struct. This method will mutate this object as we have passed it by reference
  indicated by <code>&amp;updatedPost</code> . So, this will do two things, one update
  / mutate the <code>updatedPost</code> instance from the <code>respBody</code> and
  give any error if any arises during the <a href=\"https://developer.mozilla.org/en-US/docs/Glossary/Deserialization\">deserialization</a>
  of the response.</p>\n<p>Now, we have the object in golang from the response bytes,
  we can use it as per requirements.</p>\n<p>So, that is the example in the entirety.</p>\n<p>Let’s
  simplify the steps which are similar to the POST/PUT method as well</p>\n<ul>\n<li>\n<p>Define/construct
  URL</p>\n</li>\n<li>\n<p>Marshal the object into JSON string as the request body</p>\n</li>\n<li>\n<p>Construct
  the request object (method, URL, and the body)</p>\n</li>\n<li>\n<p>Send the request
  with the default client</p>\n</li>\n<li>\n<p>Read the response and deserialize/unmarshal</p>\n</li>\n<li>\n<p>Access
  the object in golang</p>\n</li>\n</ul>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Post</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">     </span><span class=\"kt\">int</span><span
  class=\"w\">    </span><span class=\"s\">`json:&quot;id,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Title</span><span
  class=\"w\">  </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;title,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Body</span><span class=\"w\">   </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;body,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">UserId</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\">    </span><span
  class=\"s\">`json:&quot;userId,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">// define URL to hit the API</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">baseURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://jsonplaceholder.typicode.com&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">postId</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">4</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">postURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Sprintf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s/posts/%d&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">baseURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">postId</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">    </span><span class=\"c1\">// define the body -&gt; with the field
  to update</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">reqBody</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`{&quot;body&quot;:
  &quot;new body&quot;}`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;New body:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">reqBody</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"c1\">// send a
  new request, with the `PATCH` method, url and the body</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"s\">&quot;PATCH&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">postURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">strings</span><span class=\"p\">.</span><span class=\"nx\">NewReader</span><span
  class=\"p\">(</span><span class=\"nx\">reqBody</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// set the
  header content type to json</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">.</span><span class=\"nx\">Header</span><span
  class=\"p\">.</span><span class=\"nx\">Set</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Content-Type&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;application/json&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">client</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">Client</span><span
  class=\"p\">{}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">client</span><span class=\"p\">.</span><span
  class=\"nx\">Do</span><span class=\"p\">(</span><span class=\"nx\">req</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">Body</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Response status code:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">StatusCode</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Response Status:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Status</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span>\n<span
  class=\"w\">    </span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">updatedPost</span><span class=\"w\"> </span><span class=\"nx\">Post</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">respBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span>\n<span class=\"w\">    </span><span
  class=\"c1\">// convert the response json bytes to Post object in golang</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">updatedPost</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">updatedPost</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">updatedPost</span><span
  class=\"p\">.</span><span class=\"nx\">Title</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>New
  body: {&quot;body&quot;: &quot;new body&quot;}\nResponse status code: 200\nResponse
  Status: 200 OK\n{4 eum et est occaecati new body 1}\n\ntitle: eum et est occaecati
  \                                                                        \nbody:
  new body\nid: 4\nuser id: 1\n</pre></div>\n\n</pre>\n\n<p>As you can see, it has
  only updated the <code>body</code> and has not updated the other fields.</p>\n<p>If
  you would have sent a similar body with a <code>PUT</code> method, the results would
  have been different. That would have been dependent on the implementation of the
  API of course, but if there is only a few fields in the request body for a PUT method,
  it would have replaced the values with the empry values which are not present in
  the request body.</p>\n<p>That is the difference between a <code>PUT</code> and
  a <code>PATCH</code> method, the <code>PATCH</code> method, ideally should only
  update the fields of the entity that are mentioned in the request body, whereas
  the <code>PUT</code> method has to update the entire resource whether the fields
  are provided or not. Again, the implementation of these API on the server plays
  a vital role in how the behavior defers and the method in itself would perform.</p>\n<p>This
  is also called as <code>JSON Merge Patch</code></p>\n<h2>JSON Merge PATCH</h2>\n<p>The
  above API is implementing a <a href=\"https://datatracker.ietf.org/doc/html/rfc7386\">Merge
  PATCH</a> which is to say, merge the changes in the actual entity.</p>\n<p>Let’s
  say there is a Blog post Entity on a Server, you have a post that you are writing
  as an author. The post has an id of <code>4</code> let’s say and you are constantly
  changing the body of the post.</p>\n<p>So, you don’t want to send the <code>title</code>
  or <code>author_id</code> or any field that is not changing from the client again
  and again while saving, so the <code>MERGE PATCH</code> endpoint will be helpful
  in that case, where the client only sends the required fields to be updated.</p>\n<p>In
  this example, the user would only send the <code>body</code> of the post to the
  API every time it makes changes or saves the draft. In some cases, it might also
  want to change the title, so it will include the title, but not all the fields.
  The API will know it is a <code>PATCH</code> request and the content type is <code>json</code>
  so it will only change or update the fields that are provided in the request body
  to the actual entity in the database or wherever it is stored on the server.</p>\n<p>So,
  that is what is the JSON Merge PATCH or Merge PATCH in general. JSON Merge PATCH
  is specific to the JSON-based document APIs.</p>\n<p>Below is the example, the same
  steps but a different endpoint. A user API that I have specifically created to demonstrate
  the difference in PUT vs Merge PATCH vs JSON PATCH requests.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">User</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">    </span><span class=\"kt\">int</span><span
  class=\"w\">    </span><span class=\"s\">`json:&quot;id,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Name</span><span
  class=\"w\">  </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;name,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Email</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;email,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Roles</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;roles,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nx\">baseURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">userID</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">2</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">apiURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Sprintf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s/users/?id=%d&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">baseURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">userID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span>\n<span class=\"w\">    </span><span class=\"nx\">userObj</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">User</span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">Name</span><span class=\"p\">:</span><span class=\"w\">  </span><span
  class=\"s\">&quot;dummy name&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">Roles</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;dummy role&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">jsonPatchBody</span><span class=\"w\"> </span><span
  class=\"p\">[]</span><span class=\"kt\">byte</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Marshal</span><span class=\"p\">(</span><span
  class=\"nx\">userObj</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">    </span>\n<span
  class=\"w\">    </span><span class=\"c1\">// OR directly define the json as string</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">//jsonMergePatchBody
  := `{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">//
  \   &quot;name&quot;: &quot;new dummy name&quot;,</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"c1\">//    &quot;roles&quot;: &quot;new dummy
  role&quot;</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">//}`</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"s\">&quot;PATCH&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">strings</span><span class=\"p\">.</span><span class=\"nx\">NewReader</span><span
  class=\"p\">(</span><span class=\"nx\">jsonMergePatchBody</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">.</span><span class=\"nx\">Header</span><span class=\"p\">.</span><span
  class=\"nx\">Set</span><span class=\"p\">(</span><span class=\"s\">&quot;Content-Type&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;application/json&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">user</span><span class=\"w\"> </span><span class=\"nx\">User</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">user</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched User&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">user</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Name:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">.</span><span class=\"nx\">Name</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Email:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">user</span><span class=\"p\">.</span><span class=\"nx\">Email</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Bio:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">user</span><span class=\"p\">.</span><span
  class=\"nx\">Bio</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Roles:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">.</span><span class=\"nx\">Roles</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Original
  User with ID 2</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>{&quot;id&quot;:2,&quot;name&quot;:&quot;dummy&quot;,&quot;email&quot;:&quot;dummyyummy@user.com&quot;,&quot;bio&quot;:&quot;empty
  bio&quot;,&quot;roles&quot;:&quot;read&quot;}\n\nId: 2\nName: dummy\nEmail: dummyyummy@user.com\nBio:
  empty bio\nRoles: read\n</pre></div>\n\n</pre>\n\n<p>Output of the program</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>Request
  Body: {&quot;name&quot;:&quot;dummy name&quot;,&quot;roles&quot;:&quot;dummy role&quot;}\n\n{&quot;id&quot;:2,&quot;name&quot;:&quot;dummy
  name&quot;,&quot;email&quot;:&quot;dummyyummy@user.com&quot;,&quot;bio&quot;:&quot;empty
  bio&quot;,&quot;roles&quot;:&quot;dummy role&quot;}\n\nUpdated/Patched User {2 dummy
  name dummyyummy@user.com empty bio dummy role}\n\nId: 2\nName: dummy name\nEmail:
  dummyyummy@user.com\nBio: empty bio\nRoles: dummy role\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, the only fields that will be updated are <code>name</code> and
  <code>roles</code> , since the API is implemented to only update the fields provided
  in the json merge patch document (request body).</p>\n<p>As you can see that, only
  the <code>name</code> and <code>roles</code> are changed. The name was <code>dummy</code>
  that changed to <code>dummy name</code> and role changed from <code>read</code>
  to <code>dummy role</code> .</p>\n<p>Now, let’s see the same request but with the
  PUT method on it.</p>\n<p>Before we hit this API however, let’s note what the user
  with ID 2 is</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nt\">&quot;id&quot;</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">2</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"nt\">&quot;name&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;dummy name&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nt\">&quot;email&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummyyummy@user.com&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nt\">&quot;bio&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;empty bio&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nt\">&quot;roles&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy role&quot;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>This
  is the result of our recent patch request. Now, we will send a PUT request to the
  same user with a different body.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">User</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">    </span><span class=\"kt\">int</span><span
  class=\"w\">    </span><span class=\"s\">`json:&quot;id,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Name</span><span
  class=\"w\">  </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;name,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Email</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;email,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Bio</span><span
  class=\"w\">   </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;bio,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Roles</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;roles,omitempt&quot;`</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">baseURL</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">userID</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">2</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">apiURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Sprintf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s/users/?id=%d&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">baseURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">userID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span>\n<span class=\"w\">\t</span><span class=\"nx\">userObj</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">User</span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">Name</span><span class=\"p\">:</span><span class=\"w\">  </span><span
  class=\"s\">&quot;not a dummy name&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">Roles</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;not a dummy role&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">jsonPatchBody</span><span class=\"w\"> </span><span
  class=\"p\">[]</span><span class=\"kt\">byte</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Marshal</span><span class=\"p\">(</span><span
  class=\"nx\">userObj</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Request Body:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">//jsonPatchBody := `{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"c1\">//    &quot;name&quot;: &quot;dummy&quot;,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">//    &quot;roles&quot;:
  &quot;new dummy role&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">//}`</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"s\">&quot;PUT&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">strings</span><span
  class=\"p\">.</span><span class=\"nx\">NewReader</span><span class=\"p\">(</span><span
  class=\"nb\">string</span><span class=\"p\">(</span><span class=\"nx\">jsonPatchBody</span><span
  class=\"p\">)))</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">.</span><span class=\"nx\">Header</span><span
  class=\"p\">.</span><span class=\"nx\">Set</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Content-Type&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;application/merge-patch+json&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">client</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">Client</span><span
  class=\"p\">{}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">client</span><span class=\"p\">.</span><span
  class=\"nx\">Do</span><span class=\"p\">(</span><span class=\"nx\">req</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">user</span><span class=\"w\"> </span><span class=\"nx\">User</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">user</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched User&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">user</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Name:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">.</span><span class=\"nx\">Name</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Bio:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">user</span><span class=\"p\">.</span><span class=\"nx\">Bio</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Email:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">user</span><span class=\"p\">.</span><span
  class=\"nx\">Email</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Roles:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">.</span><span class=\"nx\">Roles</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Output:</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>Request
  Body: {&quot;name&quot;:&quot;not a dummy name&quot;,&quot;roles&quot;:&quot;not
  a dummy role&quot;}\n\n{&quot;id&quot;:2,&quot;name&quot;:&quot;not a dummy name&quot;,&quot;email&quot;:&quot;&quot;,&quot;bio&quot;:&quot;&quot;,&quot;roles&quot;:&quot;not
  a dummy role&quot;}\n\nUpdated/Patched User {2 not a dummy name   not a dummy role}\n\nId:
  2\nName: not a dummy name\nBio:\nEmail:\nRoles: not a dummy role\n</pre></div>\n\n</pre>\n\n<p>As
  you can see the <code>name</code> and <code>roles</code> are updated, however, the
  <code>bio</code> and <code>email</code> fields are empty. Since we only said to
  update the <code>name</code> and <code>roles</code> fields, but it was a <code>PUT</code>
  request, it expects all the fields, and if any of the fields are missing, it will
  consider them as empty and update those as well.</p>\n<p>So, the difference might
  be crystal clear now. When to use <code>PATCH</code> and when to avoid <code>PUT</code>.</p>\n<ul>\n<li>\n<p>When
  you have a large set of updates, preference could be given to PUT</p>\n</li>\n<li>\n<p>If
  you have very specific fields to update and minimal fields PATCH is recommended</p>\n</li>\n</ul>\n<p>There
  is another type of PATCH specifically designed for JSON APIs or JSON Documents APIs.</p>\n<h2>JSON
  PATCH</h2>\n<p>The <a href=\"https://datatracker.ietf.org/doc/html/rfc6902/\">JSON
  PATCH</a> is a specification in which we can define what operations to perform on
  which fields, or path of the fields to replace, move, or copy to.</p>\n<blockquote>\n<p>A
  JSON Patch document is a JSON document that represents an array of objects. Each
  object represents a single operation to be applied to the target JSON document.</p>\n</blockquote>\n<p>As
  it takes these operations, it applies them sequentially and hence it won’t replace
  all the fields for the entity, as that is the expected behavior of the PATCH method.
  In other words, it would only apply changes to the fields and the related fields
  provided in the json patch document (request body).</p>\n<p>There are a few operations
  that you can perform with this json patch method, and provide the instructions accordingly
  for individual operations in the JSON PATCH document.</p>\n<p>Operations</p>\n<ol>\n<li>\n<p>add</p>\n</li>\n<li>\n<p>remove</p>\n</li>\n<li>\n<p>replace</p>\n</li>\n<li>\n<p>move</p>\n</li>\n<li>\n<p>copy</p>\n</li>\n<li>\n<p>test</p>\n</li>\n</ol>\n<p>So,
  for each of the operations, a high-level definition can be considered as:</p>\n<ul>\n<li>\n<p>To
  add a field you can specify the operation as <code>add</code> , the path as the
  field to be added, and the value as the actual value to be added</p>\n</li>\n<li>\n<p>To
  remove a field, you can specify the operation as <code>remove</code> , and the path
  as the field to remove</p>\n</li>\n<li>\n<p>To replace a field, you can specify
  the operation as <code>replace</code>, the path as the field to be updated/replaced,
  and the value of the actual value to be added</p>\n</li>\n<li>\n<p>To move a field,
  you can specify the operation as <code>move</code>, the <strong>form</strong> as
  the field to be updated/moved from and the path to the field the from value should
  be moved to.</p>\n</li>\n<li>\n<p>To copy a field, you can specify the operation
  as <code>copy</code>, the form as the field to update/copied from, and the path
  to the field to which the value should be copied to.</p>\n</li>\n<li>\n<p>The test
  operation is a bit different as it is used for comparison of a <code>path</code>
  value to the actual value specified in the object. It might return true or false,
  but not actually return it might be used as a checkpoint for continuing with the
  operation in the document.</p>\n</li>\n</ul>\n<p>In this example, we are creating
  a similar patch request, but using this json patch document kind of structure.</p>\n<h3>Construct
  the json-patch document</h3>\n<p>We need to construct the JSON PATCH document, which
  is a list of operations (instructions). Each instruction has an operation <code>op</code>
  field that defines what operation need to be performed among the available 6 operations.
  The <code>path</code> field is the field that we need to update, it is called path
  because, we need to provide the path to the json field, in case of nested json fields.
  The <code>value</code> and <code>from</code> fields are a bit dependent on the operation
  as <code>value</code> is not required in <code>remove</code> operation and <code>from</code>
  is only used in <code>move</code> and <code>copy</code> operations.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nx\">jsonPatchBody</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">`[</span>\n<span class=\"s\">    {</span>\n<span
  class=\"s\">        &quot;op&quot;: &quot;replace&quot;,</span>\n<span class=\"s\">
  \       &quot;path&quot;: &quot;/name&quot;,</span>\n<span class=\"s\">        &quot;value&quot;:
  &quot;new dummy name&quot;</span>\n<span class=\"s\">    },</span>\n<span class=\"s\">
  \   {</span>\n<span class=\"s\">        &quot;op&quot;: &quot;replace&quot;,</span>\n<span
  class=\"s\">        &quot;path&quot;: &quot;/roles&quot;,</span>\n<span class=\"s\">
  \       &quot;value&quot;: &quot;new dummy role&quot;</span>\n<span class=\"s\">
  \   },</span>\n<span class=\"s\">]`</span><span class=\"w\"></span>\n\n\n<span class=\"c1\">//
  OR</span><span class=\"w\"></span>\n<span class=\"c1\">// construct a go type and
  then marshal it</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">JSONPatch</span><span class=\"w\"> </span><span
  class=\"kd\">struct</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Op</span><span
  class=\"w\">    </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;op&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Path</span><span class=\"w\">  </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;path&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">From</span><span class=\"w\">  </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;from,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Value</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;value,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"nx\">jsonPatchData</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"nx\">JSONPatch</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">Op</span><span class=\"p\">:</span><span
  class=\"w\">    </span><span class=\"s\">&quot;replace&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">Path</span><span
  class=\"p\">:</span><span class=\"w\">  </span><span class=\"s\">&quot;/name&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">Value</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s\">&quot;new dummy name&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">},</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">Op</span><span class=\"p\">:</span><span
  class=\"w\">    </span><span class=\"s\">&quot;replace&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">Path</span><span
  class=\"p\">:</span><span class=\"w\">  </span><span class=\"s\">&quot;/roles&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">Value</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s\">&quot;new dummy role&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">},</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"nx\">jsonPatchBytes</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Marshal</span><span class=\"p\">(</span><span
  class=\"nx\">jsonPatchData</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Error marshalling JSON:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">return</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"nx\">jsonPatchBody</span><span class=\"w\">
  </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nb\">string</span><span
  class=\"p\">(</span><span class=\"nx\">jsonPatchBytes</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In the above example, we have
  defined the two ways to create the body for a json patch request. One is directly
  writing json string (not recommended). Also, the other way is to define a struct
  as JSONPatch and construct the object for the operations in the json patch request.
  The JSONPatch object has the 4 fields as mentioned earlier. In this case, we are
  only using the <code>replace</code> operation, which will update the value in the
  provided path field to the value provided in the value field.</p>\n<p>In this specific
  example, we have 2 operations to perform, first to update the <code>name</code>field
  and the second to update the <code>roles</code> field. We simply use the <code>json.Marshal</code>
  method to convert the object into bytes. We then cast the slice of bytes into string
  and that becomes a final <code>jsonPatchBody</code> for the request that we will
  send.</p>\n<p><strong><mark>NOTE: This API </mark></strong> <code>https://dummy-json-patch.netlify.app/.netlify/functions/users/</code>
  <strong><mark> is created by me, it only implements the </mark></strong> <code>replace</code>
  <strong><mark> operation, as it makes sense to only have an update operation on
  the user table endpoint, If the model is more like a json document-like structure,
  the other operations might make sense. I might implement those in a separate article,
  while we explore the implementation of the PATCH Request in the CRUD API Section
  of this series soon.</mark></strong></p>\n<h3>Constructing and Sending the Request</h3>\n<p>We
  do the same stuff as the normal merge patch request, that is provide the <code>PATCH</code>
  as the request method, the <code>apiURL</code> and the <code>strings.NewReader(jsonPatchBody)</code>
  into the body of the request.</p>\n<p>However, the only difference is that we need
  to set the <code>Content-Type</code> header as <code>application/json-patch+json</code>
  . This is defined in the <a href=\"https://datatracker.ietf.org/doc/html/rfc6902/#section-1\">specification</a>,
  it would make the client as well as the server understand the request body. This
  can also be used by the server to differentiate between a standard patch (which
  can be referred to as a &quot;Merge PATCH&quot; for JSON APIs) and a JSON PATCH
  request, as shown here.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"s\">&quot;PATCH&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">strings</span><span
  class=\"p\">.</span><span class=\"nx\">NewReader</span><span class=\"p\">(</span><span
  class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"nx\">req</span><span class=\"p\">.</span><span class=\"nx\">Header</span><span
  class=\"p\">.</span><span class=\"nx\">Set</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Content-Type&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;application/json-patch+json&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Once, we have created the request
  object we then simply can create the client and send the request.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"nx\">resp</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">updatedUser</span><span
  class=\"w\"> </span><span class=\"nx\">User</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">resBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">io</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nb\">string</span><span
  class=\"p\">(</span><span class=\"nx\">resBody</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">updatedUser</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched User&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedUser</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Once,
  we have the response from the request, we just do the normal response processing
  and unmarshal the response string into a golang object type (passing by reference)
  which is provided as <code>&amp;updatedUser</code> in the parameter.</p>\n<p>This
  will give us the expected result back in the updated User object.</p>\n<p>The below
  code is the entire JSON PATCH request:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">User</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">    </span><span class=\"kt\">int</span><span
  class=\"w\">    </span><span class=\"s\">`json:&quot;id&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Name</span><span class=\"w\">  </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;name&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Email</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;email&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nx\">Bio</span><span class=\"w\">   </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;bio&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Roles</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;roles&quot;`</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">baseURL</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">userID</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">2</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">apiURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Sprintf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s/users/?id=%d&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">baseURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">userID</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`[</span>\n<span
  class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;: &quot;replace&quot;,</span>\n<span
  class=\"s\">            &quot;path&quot;: &quot;/name&quot;,</span>\n<span class=\"s\">
  \           &quot;value&quot;: &quot;new dummy name&quot;</span>\n<span class=\"s\">
  \       },</span>\n<span class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;:
  &quot;replace&quot;,</span>\n<span class=\"s\">            &quot;path&quot;: &quot;/roles&quot;,</span>\n<span
  class=\"s\">            &quot;value&quot;: &quot;new dummy role&quot;</span>\n<span
  class=\"s\">        },</span>\n<span class=\"s\">    ]`</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"s\">&quot;PATCH&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">strings</span><span class=\"p\">.</span><span class=\"nx\">NewReader</span><span
  class=\"p\">(</span><span class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">.</span><span class=\"nx\">Header</span><span class=\"p\">.</span><span
  class=\"nx\">Set</span><span class=\"p\">(</span><span class=\"s\">&quot;Content-Type&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;application/json-patch+json&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">user</span><span class=\"w\"> </span><span class=\"nx\">User</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">user</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched User&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">user</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Name:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">.</span><span class=\"nx\">Name</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Bio:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">user</span><span class=\"p\">.</span><span class=\"nx\">Bio</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Email:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">user</span><span class=\"p\">.</span><span
  class=\"nx\">Email</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Roles:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">.</span><span class=\"nx\">Roles</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Original
  User (id=2) before the request</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nt\">&quot;id&quot;</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">2</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"nt\">&quot;name&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;dummy name&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nt\">&quot;email&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummyyummy@user.com&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nt\">&quot;bio&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;empty bio&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nt\">&quot;roles&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy role&quot;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Output
  of the Program (JSON PATCH Request)</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>{&quot;id&quot;:2,&quot;name&quot;:&quot;new
  dummy name&quot;,&quot;email&quot;:&quot;dummyyummy@user.com&quot;,&quot;bio&quot;:
  &quot;empty bio&quot;, &quot;roles&quot;:&quot;new dummy role&quot;}     \n\nUpdated/Patched
  User {2 new dummy name dummyyummy@user.com empty bio new dummy role} \n\nId: 2\nName:
  new dummy name\nEmail: dummyyummy@user.com\nBio: empty bio\nRoles: read\n</pre></div>\n\n</pre>\n\n<p>So,
  as expected, it only updates the <code>name</code> and the <code>roles</code> field
  in the user object with id = 2. The patch document has the replace operation that
  updates the name and the roles fields from the existing field.</p>\n<p>To add more
  examples for <code>add</code> , <code>move</code>,<code>copy</code>, <code>remove</code>
  and <code>test</code> operations.</p>\n<p>PS: I added an endpoint for documents,
  which will include all the crud operations with all the JSON PATCH operations added
  as well.</p>\n<p>Let’s take a simple example for each operation</p>\n<h3>Add operation</h3>\n<p>Let’s
  take an example for a document table at the <code>https://dummy-json-patch.netlify.app/.netlify/functions/documents</code>
  endpoint, the document is a table with the columns like data and the id as the index.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"s\">&quot;data&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"p\">{},</span><span class=\"w\"></span>\n<span class=\"w\">
  \ </span><span class=\"s\">&quot;id&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">4</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  initial data is empty for a document with an ID is 4.</p>\n<p>We simply use the
  operation of <code>add</code> that will add the value of the <code>name</code> in
  the <code>data</code> object with the field name as <code>name</code> and the value
  as <code>dummy</code> .</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Document</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">   </span><span class=\"kt\">int</span><span
  class=\"w\">         </span><span class=\"s\">`json:&quot;id&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Data</span><span
  class=\"w\"> </span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;data&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions/documents/?id=4&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`[</span>\n<span
  class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;: &quot;add&quot;,</span>\n<span
  class=\"s\">            &quot;path&quot;: &quot;/name&quot;,</span>\n<span class=\"s\">
  \           &quot;value&quot;: &quot;dummy&quot;</span>\n<span class=\"s\">        }</span>\n<span
  class=\"s\">    ]`</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"s\">&quot;PATCH&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">strings</span><span
  class=\"p\">.</span><span class=\"nx\">NewReader</span><span class=\"p\">(</span><span
  class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">.</span><span
  class=\"nx\">Header</span><span class=\"p\">.</span><span class=\"nx\">Set</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Content-Type&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;application/json-patch+json&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">updatedDoc</span><span class=\"w\"> </span><span class=\"nx\">Document</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">updatedDoc</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched Document&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;document:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">.</span><span class=\"nx\">Data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>This
  is a simple add operation, as you would see that the data now has the name field
  populated after the request was sent.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"s\">&quot;data&quot;</span><span class=\"p\">:{</span><span
  class=\"s\">&quot;name&quot;</span><span class=\"p\">:</span><span class=\"s\">&quot;dummy&quot;</span><span
  class=\"p\">},</span><span class=\"s\">&quot;id&quot;</span><span class=\"p\">:</span><span
  class=\"mi\">4</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"nx\">Updated</span><span class=\"o\">/</span><span class=\"nx\">Patched</span><span
  class=\"w\"> </span><span class=\"nx\">Document</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"mi\">4</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"nx\">name</span><span class=\"p\">:</span><span
  class=\"nx\">dummy</span><span class=\"p\">]}</span><span class=\"w\"></span>\n<span
  class=\"nx\">id</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">4</span><span class=\"w\"></span>\n<span class=\"nx\">document</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"nx\">name</span><span class=\"p\">:</span><span
  class=\"nx\">dummy</span><span class=\"p\">]</span><span class=\"w\"> </span>\n</pre></div>\n\n</pre>\n\n<h3>Remove
  Operation</h3>\n<p>We can take an example for remove, let’s remove the <code>name</code>
  in the <code>data</code> field that we added previously.</p>\n<p>The data before
  the request looks like this:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;data&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"s\">&quot;name&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s\">&quot;dummy&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"p\">},</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"s\">&quot;id&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">4</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Now,
  we will construct the request, we can do it with the json marshaling by constructing
  the object, but I am keeping it simple and shorter by directly constructing the
  json body.</p>\n<p>For the <code>remove</code> operation, we only need the <code>path</code>
  field to be added, in order to remove the field from the object.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Document</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">   </span><span class=\"kt\">int</span><span
  class=\"w\">         </span><span class=\"s\">`json:&quot;id&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Data</span><span
  class=\"w\"> </span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;data&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions/documents/?id=4&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`[</span>\n<span
  class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;: &quot;remove&quot;,</span>\n<span
  class=\"s\">            &quot;path&quot;: &quot;/name&quot;</span>\n<span class=\"s\">
  \       }</span>\n<span class=\"s\">    ]`</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"s\">&quot;PATCH&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">strings</span><span class=\"p\">.</span><span class=\"nx\">NewReader</span><span
  class=\"p\">(</span><span class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">.</span><span class=\"nx\">Header</span><span class=\"p\">.</span><span
  class=\"nx\">Set</span><span class=\"p\">(</span><span class=\"s\">&quot;Content-Type&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;application/json-patch+json&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">updatedDoc</span><span class=\"w\"> </span><span class=\"nx\">Document</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">updatedDoc</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched Document&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;document:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">.</span><span class=\"nx\">Data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>{&quot;data&quot;:{},&quot;id&quot;:4}
  \               \n\nUpdated/Patched Document {4 map[]}\nid: 4                             \ndocument:
  map[]                   \n</pre></div>\n\n</pre>\n\n<p>So, as you can see the updated
  data has no fields.</p>\n<h3>Copy Operation</h3>\n<p>Now, we can also perform the
  copy operation, which can copy the field to another field, this instruction requires
  the <code>from</code> and the <code>path</code> . The <code>from</code> is the path
  to the field to copy it from, and the <code>path</code> is the path to the field
  to copy the field value.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">[</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;op&quot;</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s2\">&quot;copy&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;from&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;/name&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;path&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;/new_name&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">]</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  above json path operation is for copying the value of the field <code>name</code>
  to the field <code>new_name</code> i.e. if the original document looks likes this:</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;data&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;name&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;dummy&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;new_name&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;hello&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">},</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;id&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"mi\">4</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  after this the value of <code>new_name</code> will be <code>dummy</code> after the
  operation</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nt\">&quot;data&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nt\">&quot;name&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;new_name&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">},</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;id&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"mi\">4</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  let’s take a look at the code where we will perform this operation with this API
  endpoint and constructing the json body.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Document</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">   </span><span class=\"kt\">int</span><span
  class=\"w\">         </span><span class=\"s\">`json:&quot;id&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Data</span><span
  class=\"w\"> </span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;data&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions/documents/?id=4&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`[</span>\n<span
  class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;: &quot;copy&quot;,</span>\n<span
  class=\"s\">            &quot;from&quot;: &quot;/name&quot;,</span>\n<span class=\"s\">
  \           &quot;path&quot;: &quot;/new_name&quot;</span>\n<span class=\"s\">        }</span>\n<span
  class=\"s\">    ]`</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"s\">&quot;PATCH&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">strings</span><span
  class=\"p\">.</span><span class=\"nx\">NewReader</span><span class=\"p\">(</span><span
  class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">.</span><span
  class=\"nx\">Header</span><span class=\"p\">.</span><span class=\"nx\">Set</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Content-Type&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;application/json-patch+json&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">updatedDoc</span><span class=\"w\"> </span><span class=\"nx\">Document</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">updatedDoc</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched Document&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;document:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">.</span><span class=\"nx\">Data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we are creating the json payload for copying the value of <code>name</code>
  to the <code>new_name</code> field.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>{&quot;data&quot;:{&quot;name&quot;:&quot;dummy&quot;,&quot;new_name&quot;:&quot;dummy&quot;},&quot;id&quot;:4}\nUpdated/Patched
  Document {4 map[name:dummy new_name:dummy]}\nid: 4\ndocument: map[name:dummy new_name:dummy]\n</pre></div>\n\n</pre>\n\n<h3>Move
  Operation</h3>\n<p>Now, we can also perform the move operation, which can move the
  field to another field, this instruction requires the <code>from</code> and the
  <code>path</code> . The <code>from</code> is the path to the field to move it from,
  and the <code>path</code> is the path to the field to move.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">[</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;op&quot;</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s2\">&quot;move&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;from&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;/name&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;path&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;/user/description&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">]</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  above json path operation is for moving the value of the field <code>name</code>
  to the field <code>/user/description</code> i.e. if the original document looks
  like this:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nt\">&quot;name&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;new_name&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;user&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;description&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;hello&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  after the operation, the value of <code>name</code> will be <code>dummy</code>.
  The name field is moved not replaced, that is the difference of this operation.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nt\">&quot;new_name&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;user&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;description&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  let’s take a look at the code where we will construct the payload body and perform
  the json patch request with the move operations.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Document</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">   </span><span class=\"kt\">int</span><span
  class=\"w\">         </span><span class=\"s\">`json:&quot;id&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Data</span><span
  class=\"w\"> </span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;data&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions/documents/?id=4&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`[</span>\n<span
  class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;: &quot;move&quot;,</span>\n<span
  class=\"s\">            &quot;from&quot;: &quot;/name&quot;,</span>\n<span class=\"s\">
  \           &quot;path&quot;: &quot;/user/description&quot;</span>\n<span class=\"s\">
  \       }</span>\n<span class=\"s\">    ]`</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"s\">&quot;PATCH&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">strings</span><span class=\"p\">.</span><span class=\"nx\">NewReader</span><span
  class=\"p\">(</span><span class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">.</span><span class=\"nx\">Header</span><span class=\"p\">.</span><span
  class=\"nx\">Set</span><span class=\"p\">(</span><span class=\"s\">&quot;Content-Type&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;application/json-patch+json&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">updatedDoc</span><span class=\"w\"> </span><span class=\"nx\">Document</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">updatedDoc</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched Document&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;document:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">.</span><span class=\"nx\">Data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  above code simply takes in the json string as a payload for the API with the move
  operation. Hence the name field’s value is swapped to the value of <code>dummy</code>
  of the <code>/user/description</code> field.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"s\">&quot;data&quot;</span><span class=\"p\">:{</span><span
  class=\"s\">&quot;new_name&quot;</span><span class=\"p\">:</span><span class=\"s\">&quot;dummy&quot;</span><span
  class=\"p\">,</span><span class=\"s\">&quot;user&quot;</span><span class=\"p\">:{</span><span
  class=\"s\">&quot;description&quot;</span><span class=\"p\">:</span><span class=\"s\">&quot;dummy&quot;</span><span
  class=\"p\">}},</span><span class=\"s\">&quot;id&quot;</span><span class=\"p\">:</span><span
  class=\"mi\">4</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"nx\">Updated</span><span class=\"o\">/</span><span class=\"nx\">Patched</span><span
  class=\"w\"> </span><span class=\"nx\">Document</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"mi\">4</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"nx\">new_name</span><span class=\"p\">:</span><span
  class=\"nx\">dummy</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">:</span><span class=\"kd\">map</span><span class=\"p\">[</span><span
  class=\"nx\">description</span><span class=\"p\">:</span><span class=\"nx\">dummy</span><span
  class=\"p\">]]}</span><span class=\"w\"></span>\n<span class=\"nx\">id</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"mi\">4</span><span class=\"w\"></span>\n<span
  class=\"nx\">document</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"kd\">map</span><span class=\"p\">[</span><span class=\"nx\">new_name</span><span
  class=\"p\">:</span><span class=\"nx\">dummy</span><span class=\"w\"> </span><span
  class=\"nx\">user</span><span class=\"p\">:</span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"nx\">description</span><span class=\"p\">:</span><span
  class=\"nx\">dummy</span><span class=\"p\">]]</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<h3>Test
  Operation</h3>\n<p>The test operation requires the <code>path</code> field with
  the value of the field to test, and the <code>value</code> field as the value to
  compare with. This is a test, so we can use it as a flag to proceed ahead with the
  rest of the operations or abort.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">[</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;op&quot;</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s2\">&quot;test&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;path&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;/new_name&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;value&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;testing value&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">},</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;op&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;add&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;path&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;/user/id&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;value&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"mi\">123</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">]</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>if
  the original document looks like this:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nt\">&quot;new_name&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;user&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;description&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Then
  after applying the test operation, we can find that the <code>new_name</code> has
  the value is <code>dummy</code> and not <code>testing value</code> so it will abort
  with an error.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;error&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;test
  failed: values do not match at path new_name&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>However,
  if we keep the operation like this:</p>\n<p>So, we compare the value of <code>new_name</code>
  field is <code>dummy</code> which it is then we will perform the rest of the operations
  in this case, the <code>add</code> operations will be performed to add the <code>id</code>
  field to the <code>user</code> object with the value <code>123</code>.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">[</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;op&quot;</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s2\">&quot;test&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;path&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;/new_name&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;value&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">},</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;op&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;add&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;path&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;/user/id&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;value&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"mi\">123</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">]</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>if
  the original document looks like this:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;new_name&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;dummy&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;user&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;description&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Then
  after the operation of test and add, the document becomes:</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;new_name&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;dummy&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;user&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;description&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;id&quot;</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">123</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>You
  can observe the <code>id</code> field is added to the <code>user</code> field, as
  the test was successful. Hence, the operation <code>add</code> was performed and
  the request was completed.</p>\n<p>Let’s take a look at the code for doing this
  operation:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Document</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">   </span><span class=\"kt\">int</span><span
  class=\"w\">         </span><span class=\"s\">`json:&quot;id&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Data</span><span
  class=\"w\"> </span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;data&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions/documents/?id=4&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"cm\">/*</span>\n<span
  class=\"cm\">    jsonPatchBody := `[</span>\n<span class=\"cm\">        {</span>\n<span
  class=\"cm\">            &quot;op&quot;: &quot;test&quot;,</span>\n<span class=\"cm\">
  \           &quot;path&quot;: &quot;/new_name&quot;,</span>\n<span class=\"cm\">
  \           &quot;value&quot;: &quot;testing value&quot;</span>\n<span class=\"cm\">
  \       },</span>\n<span class=\"cm\">        {</span>\n<span class=\"cm\">            &quot;op&quot;:
  &quot;add&quot;,</span>\n<span class=\"cm\">            &quot;path&quot;: &quot;/user/id&quot;,</span>\n<span
  class=\"cm\">            &quot;value&quot;: 123</span>\n<span class=\"cm\">        }</span>\n<span
  class=\"cm\">    ]`</span>\n<span class=\"cm\">    */</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`[</span>\n<span
  class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;: &quot;test&quot;,</span>\n<span
  class=\"s\">            &quot;path&quot;: &quot;/new_name&quot;,</span>\n<span class=\"s\">
  \           &quot;value&quot;: &quot;dummy&quot;</span>\n<span class=\"s\">        },</span>\n<span
  class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;: &quot;add&quot;,</span>\n<span
  class=\"s\">            &quot;path&quot;: &quot;/user/id&quot;,</span>\n<span class=\"s\">
  \           &quot;value&quot;: 123</span>\n<span class=\"s\">        }</span>\n<span
  class=\"s\">    ]`</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"s\">&quot;PATCH&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">strings</span><span
  class=\"p\">.</span><span class=\"nx\">NewReader</span><span class=\"p\">(</span><span
  class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">.</span><span
  class=\"nx\">Header</span><span class=\"p\">.</span><span class=\"nx\">Set</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Content-Type&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;application/json-patch+json&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">updatedDoc</span><span class=\"w\"> </span><span class=\"nx\">Document</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">updatedDoc</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched Document&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;document:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">.</span><span class=\"nx\">Data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  the above code is an example of performing the test operation.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"nt\">&quot;data&quot;</span><span class=\"p\">:{</span><span
  class=\"nt\">&quot;new_name&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;user&quot;</span><span class=\"p\">:{</span><span
  class=\"nt\">&quot;description&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;id&quot;</span><span class=\"p\">:</span><span
  class=\"mi\">123</span><span class=\"p\">}},</span><span class=\"nt\">&quot;id&quot;</span><span
  class=\"p\">:</span><span class=\"mi\">4</span><span class=\"p\">}</span><span class=\"w\">
  \                                                                   </span>\n<span
  class=\"err\">Upda</span><span class=\"kc\">te</span><span class=\"err\">d/Pa</span><span
  class=\"kc\">t</span><span class=\"err\">ched</span><span class=\"w\"> </span><span
  class=\"err\">Docume</span><span class=\"kc\">nt</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"mi\">4</span><span class=\"w\"> </span><span class=\"err\">map</span><span
  class=\"p\">[</span><span class=\"kc\">ne</span><span class=\"err\">w_</span><span
  class=\"kc\">na</span><span class=\"err\">me</span><span class=\"p\">:</span><span
  class=\"err\">dummy</span><span class=\"w\"> </span><span class=\"err\">user</span><span
  class=\"p\">:</span><span class=\"err\">map</span><span class=\"p\">[</span><span
  class=\"err\">descrip</span><span class=\"kc\">t</span><span class=\"err\">io</span><span
  class=\"kc\">n</span><span class=\"p\">:</span><span class=\"err\">dummy</span><span
  class=\"w\"> </span><span class=\"err\">id</span><span class=\"p\">:</span><span
  class=\"mi\">123</span><span class=\"p\">]]}</span><span class=\"w\">                                                             </span>\n<span
  class=\"err\">id</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">4</span><span class=\"w\">                                                                                                                                           </span>\n<span
  class=\"err\">docume</span><span class=\"kc\">nt</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"err\">map</span><span class=\"p\">[</span><span
  class=\"kc\">ne</span><span class=\"err\">w_</span><span class=\"kc\">na</span><span
  class=\"err\">me</span><span class=\"p\">:</span><span class=\"err\">dummy</span><span
  class=\"w\"> </span><span class=\"err\">user</span><span class=\"p\">:</span><span
  class=\"err\">map</span><span class=\"p\">[</span><span class=\"err\">descrip</span><span
  class=\"kc\">t</span><span class=\"err\">io</span><span class=\"kc\">n</span><span
  class=\"p\">:</span><span class=\"err\">dummy</span><span class=\"w\"> </span><span
  class=\"err\">id</span><span class=\"p\">:</span><span class=\"mi\">123</span><span
  class=\"p\">]]</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  as you would see, the id is indeed added to the document.</p>\n<p>Here’s an overview
  of PUT, PATCH, MERGE PATCH, and JSON PATCH requests.</p>\n<table>\n<thead>\n<tr>\n<th></th>\n<th>PUT</th>\n<th>PATCH
  (General)</th>\n<th>Merge Patch (JSON)</th>\n<th>JSON Patch</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td><strong>Purpose</strong></td>\n<td>Replace
  the entire resource.</td>\n<td>Partially modify a resource.</td>\n<td>Partially
  modify a JSON resource by merging changes.</td>\n<td>Partially modify a JSON resource
  using explicit operations.</td>\n</tr>\n<tr>\n<td><strong>Request Body</strong></td>\n<td>Full
  representation of the resource.</td>\n<td>Instructions for modification.</td>\n<td>JSON
  object containing fields to be updated.</td>\n<td>JSON array of operation objects
  (e.g., <code>add</code>, <code>remove</code>, <code>replace</code>).</td>\n</tr>\n<tr>\n<td><strong>Idempotency</strong></td>\n<td>Idempotent
  (repeated requests have the same effect).</td>\n<td>Not necessarily idempotent (depends
  on implementation).</td>\n<td>Idempotent (repeated requests have the same effect).</td>\n<td>Idempotent
  (if operations are idempotent, like <code>replace</code>).</td>\n</tr>\n<tr>\n<td>Safe</td>\n<td>No</td>\n<td>No</td>\n<td>No</td>\n<td>No</td>\n</tr>\n<tr>\n<td><strong>Data
  Sent</strong></td>\n<td>All resource fields.</td>\n<td>Only the fields/instructions
  to be modified.</td>\n<td>Only the fields to be updated.</td>\n<td>Instructions
  specifying how to modify the resource (path, operation, value).</td>\n</tr>\n<tr>\n<td><strong>Server
  Behavior</strong></td>\n<td>Replaces the entire resource.</td>\n<td>Applies the
  provided modifications.</td>\n<td>Merges the provided fields into the existing resource.</td>\n<td>Applies
  the operations in the order they appear in the array.</td>\n</tr>\n<tr>\n<td><strong>Content-Type</strong></td>\n<td><code>application/json</code>
  (usually)</td>\n<td>Varies depending on the patch format.</td>\n<td><code>application/merge-patch+json</code></td>\n<td><code>application/json-patch+json</code></td>\n</tr>\n</tbody>\n</table>\n<p>So,
  that is all from the PATCH method in GOlang, I might have explained a bit too much
  for this, but there is a lot more to learn, which we will explore while creating
  a CRUD API in Golang later in this series.</p>\n<h2>Conclusion</h2>\n<p>I have also
  included some more examples of PATCH requests <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/patch/\"><strong>here</strong></a>.</p>\n<p>That's
  it from the 36th part of the series, all the source code for the examples are linked
  in the GitHub on the <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/patch/\"><strong>100
  days of Golang</strong></a> repository.</p>\n<p><a href=\"https://github.com/Mr-Destructive/100-days-of-golang\"><strong>100-days-of-golang</strong></a></p>\n<h3>References:</h3>\n<ul>\n<li>\n<p><a
  href=\"https://en.wikipedia.org/wiki/HTTP#Request_methods\">HTTP Request Methods</a></p>\n</li>\n<li>\n<p><a
  href=\"https://rubyonrails.org/2012/2/26/edge-rails-patch-is-the-new-primary-http-method-for-updates\">Ruby
  on Rails Patch method</a></p>\n</li>\n<li>\n<p><a href=\"https://datatracker.ietf.org/doc/html/rfc5789\">HTTP
  PATCH RFC</a></p>\n</li>\n<li>\n<p><a href=\"https://datatracker.ietf.org/doc/html/rfc6902\">HTTP
  JSON PATCH RFC</a></p>\n</li>\n<li>\n<p><a href=\"https://datatracker.ietf.org/doc/html/rfc7386\">HTTP
  JSON MERGE PATCH</a></p>\n</li>\n</ul>\n"
cover: ''
date: 2025-01-11
datetime: 2025-01-11 00:00:00+00:00
description: Exploring the fundamentals of a PATCH method in golang. How to request
  a resource, parse body, headers, etc. in a HTTP PATCH request.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2025-01-11-GO-36-PATCH-Method.md
html: "<h2>Introduction</h2>\n<p>In previous sections of this series, we've covered
  the GET, POST, and PUT methods. Now, we will explore the PATCH method, which differs
  from the others in several key ways. The PATCH method is somewhat more flexible
  and depends on how the server or API you're working with is designed.</p>\n<p>In
  this section, we'll focus on understanding what the PATCH method is and how to use
  it. While we will dive deeper into building and structuring a full CRUD API later
  in the series, the focus here will be on the what and why of the PATCH method, not
  the how.</p>\n<h2>What is the PATCH Method?</h2>\n<p>The PATCH method is often compared
  to the PUT method, but with one important distinction: PATCH is used to perform
  partial updates on a resource. Unlike PUT, which typically requires you to send
  the entire resource to update it, PATCH allows you to send only the fields that
  need to be updated. This makes it a more efficient option when updating a subset
  of a resource.</p>\n<p>In a PATCH request, the body usually contains instructions
  in a format like JSON, which specifies the fields to update. These instructions
  define the changes to be applied to the resource. For example, you may only want
  to change one field of a user's profile, such as their email address, while leaving
  the rest of the data untouched.</p>\n<h2>PATCH vs. PUT</h2>\n<p>Key Differences
  While both PATCH and PUT are used to modify resources, there are significant differences
  in their behavior:</p>\n<ul>\n<li>\n<p>PUT replaces the entire resource. When you
  send a PUT request, you must include the full representation of the resource, even
  if you're only changing a small part of it.</p>\n</li>\n<li>\n<p>PATCH, on the other
  hand, is for partial updates. You only need to include the fields that are changing,
  not the entire resource.</p>\n</li>\n</ul>\n<p>If the update involves more fields
  than just the ones you're changing, PUT may be the better choice. However, the scope
  of this article is to focus solely on the PATCH method.</p>\n<h2>How Does a PATCH
  Request Work?</h2>\n<p>In the simplest terms, a PATCH request allows you to perform
  a partial update on a resource. It is similar to a PUT request, but more specific
  in how it updates the resource. According to the HTTP specification, an ideal PATCH
  request should:</p>\n<ul>\n<li>\n<p>Accept a &quot;patch document&quot; in the request
  body, which contains the list of operations to apply (e.g., &quot;replace&quot;,
  &quot;add&quot;, &quot;remove&quot;).</p>\n</li>\n<li>\n<p>Apply these updates to
  the target resource.</p>\n</li>\n<li>\n<p>If the update cannot be applied correctly,
  the operation should fail without applying any of the changes.</p>\n</li>\n</ul>\n<p>This
  ensures that no partial or inconsistent updates are left behind</p>\n<p>For example,
  if you're updating a user's email address and something goes wrong in the middle
  of the operation, the PATCH request should ensure that the email isn't updated partially.
  If there’s an error, none of the updates should be applied, ensuring data consistency.</p>\n<p>Also,
  the patch method is not idempotent, meaning that if you send the same input/request,
  it need not necessarily return the same output. Because we are not sending the actual
  original entity, we are sending the partial data fields that need to be updated,
  so it might update the original entity on subsequent requests since there is no
  original request sent in the request; it only identifies the resource from the URI
  and fields to update in the request body.</p>\n<p>Now, let’s sum up the patch request
  in a few words</p>\n<ul>\n<li>\n<p>Updates specific fields mentioned in the patch
  document</p>\n</li>\n<li>\n<p>Can be partial (only the fields that need to be updated
  are sent, unlike PUT, which typically replaces the entire resource)</p>\n</li>\n<li>\n<p>Not
  necessarily idempotent (depends on the implementation)</p>\n</li>\n<li>\n<p>Not
  Safe (since resources will be updated on the server side)</p>\n</li>\n</ul>\n<h2>Basic
  PATCH request</h2>\n<p>Let’s start with the basic PATCH request that we can create
  in Golang. The <a href=\"https://pkg.go.dev/net/http\">net/http</a> package will
  be used to construct the request and will be using <code>encoding/json</code> and
  some other utilities for string and byte parsing.</p>\n<p>So, first, we will construct
  an HTTP request using the <a href=\"https://pkg.go.dev/net/http#NewRequest\">http.NewRequest</a>
  with the parameters like the http method to use, the URL to hit, and the request
  body if any. We will then need to send the json body which would consist of the
  fields to be updated.</p>\n<h3>Defining the API/Server Endpoint URL</h3>\n<p>We
  would need to define the endpoint that we are hitting, we can directly use the API
  URL or we can construct the API URL on the fly depending on the ID, and parameter
  that will be dynamic. As PATCH requests, usually modify a particular entity, we
  would generally have some form of identifier for that entity/object on the database,
  etc. So in this case, it is <code>id</code> of the post, so, we can pass the post
  in the URL.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"c1\">// define URL to hit the API</span><span class=\"w\"></span>\n<span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://jsonplaceholder.typicode.com/posts/4&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"c1\">// OR</span><span class=\"w\"></span>\n<span
  class=\"c1\">// baseURL := &quot;https://jsonplaceholder.typicode.com&quot;</span><span
  class=\"w\"></span>\n<span class=\"c1\">// postId := 4</span><span class=\"w\"></span>\n<span
  class=\"c1\">// postURL := fmt.Sprintf(&quot;%s/posts/%d&quot;, baseURL, postId)</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>We can either directly define
  the URL or dynamically construct the URL, that is quite straightforward. The latter
  one is the one we usually use and design.</p>\n<h3>Constructing the JSON Body</h3>\n<p>This
  section is a little dependent on the context as you might have a direct json string
  that you can directly pass to the API or you might have a golang object that you
  need to Marshal in order to convert that object into string/bytes.</p>\n<ol>\n<li>\n<p>Direct
  JSON String</p>\n<p>So, there is nothing to do here, since the object is already
  in the form of a json string.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nx\">reqBody</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">`{&quot;body&quot;: &quot;new body&quot;}`</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>However, if you have certain fields
  that you need to exclude or omit, you have to construct a struct and then marshal
  it</p>\n</li>\n<li>\n<p>Marshaling (converting object into bytes/string)</p>\n<p>We
  need to convert the Golang native object into some form of a json string or bytes
  that could be sent over the network. That process is called <a href=\"https://en.wikipedia.org/wiki/Marshalling_(computer_science)\">marshaling</a>
  or serialization.</p>\n</li>\n</ol>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kd\">type</span><span class=\"w\"> </span><span class=\"nx\">Post</span><span
  class=\"w\"> </span><span class=\"kd\">struct</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">     </span><span class=\"kt\">int</span><span
  class=\"w\">    </span><span class=\"s\">`json:&quot;id,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Title</span><span
  class=\"w\">  </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;title,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Body</span><span class=\"w\">   </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;body,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">UserId</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\">    </span><span
  class=\"s\">`json:&quot;userId,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"nx\">userObj</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">Post</span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">Body</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s\">&quot;New Body&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">reqBody</span><span class=\"w\"> </span><span
  class=\"p\">[]</span><span class=\"kt\">byte</span><span class=\"w\"></span>\n<span
  class=\"nx\">reqBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Marshal</span><span class=\"p\">(</span><span class=\"nx\">userObj</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;New body:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">reqBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"c1\">// New body: {&quot;body&quot;:&quot;New Body&quot;}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above snippet, we have defined a <code>Post</code> struct with the fields like
  <code>ID</code>, <code>Title</code>, <code>Body</code>, <code>UserID</code> , and
  those who have <code>omitempty</code> tag along with the json fields that we want
  to marshal into. The omit empty will omit or ignore the fields if they are empty
  or not present in the object/instance of this structure. So in the example, <code>userObj</code>
  is the instance of the <code>Post</code> struct and it only has the <code>Body</code>
  populated, so the reqBody will only contain one field <code>body</code> in the json
  representation. The <a href=\"https://pkg.go.dev/encoding/json#Marshal\">json.Marshal</a>
  is the function that we use to convert the object/instance into a byte form.</p>\n<p>This
  <code>reqBody</code> will serve as the request body for the request that will be
  a <code>PATCH</code> method to the mentioned endpoint / API URL.</p>\n<h3>Constructing
  the HTTP PATCH Request</h3>\n<p>Now, we have the parts that we need to construct
  the request, we can combine the parts and hit the endpoint. However, it is a bit
  different compared to <code>GET</code> and <code>POST</code> request that we do
  in Golang, the HTTP package has built-in methods for the <code>GET</code> and <code>POST</code>
  methods, however for methods like <code>PUT</code>, <code>PATCH</code>, <code>DELETE</code>
  and others, we need to construct a <a href=\"https://pkg.go.dev/net/http#Request\">Request</a>
  object and then send that request.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"s\">&quot;PATCH&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">postURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">strings</span><span
  class=\"p\">.</span><span class=\"nx\">NewReader</span><span class=\"p\">(</span><span
  class=\"nx\">reqBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"nx\">req</span><span class=\"p\">.</span><span class=\"nx\">Header</span><span
  class=\"p\">.</span><span class=\"nx\">Set</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Content-Type&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;application/json&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"c1\">// in case of wired utf-8 characters appear
  in the body</span><span class=\"w\"></span>\n<span class=\"c1\">//req.Header.Set(&quot;Content-Type&quot;,
  &quot;application/json; charset=utf-8&quot;)</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>To
  do that, we call the <a href=\"https://pkg.go.dev/net/http#NewRequest\">NewRequest</a>
  method with the parameters like the HTTP method, the URL, and the request Body all
  of which we have at the moment.</p>\n<ul>\n<li>\n<p>The method is <code>PATCH</code></p>\n</li>\n<li>\n<p>The
  URL is <code>postURL</code></p>\n</li>\n<li>\n<p>The body is <code>strings.NewReader(reqBody)</code>
  as we need a <code>io.Reader</code> object instead of string or byte slice</p>\n</li>\n</ul>\n<p>So,
  once we have that, we will also set the <code>Header</code> In the field of <code>Content-Type</code>
  and the value as <code>application/json</code> since the request body has json representation
  of the patch document that will be sent.</p>\n<h3>Sending the Request</h3>\n<p>Once,
  the <code>req</code> the object is created, we also need a <a href=\"https://pkg.go.dev/net/http#Client\">Client</a>
  to send the request, so we create the client as default http.Client object with
  defaults and call the <a href=\"https://pkg.go.dev/net/http#Client.Do\">Do</a> method
  with the <code>req</code> as the request parameter in order to send the request
  with the default client.</p>\n<p>This method returns the response object and an
  error if any.</p>\n<p>We also add the <code>defer resp.Body.Close()</code> in order
  to avoid leaks and safely access the response body.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"nx\">resp</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>At
  this point, we can now start consuming the response and use it for further processing
  as per the needs.</p>\n<h3>Unmarshalling the Response</h3>\n<p>We first read the
  response into a string or byte representation using the io.ReadAll method and then
  use the json.Unmarshal to convert the bytes into golang object/instance.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">updatedPost</span><span
  class=\"w\"> </span><span class=\"nx\">Post</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">io</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span>\n<span class=\"c1\">// convert the response json bytes to
  Post object in golang</span><span class=\"w\"></span>\n<span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">updatedPost</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nx\">updatedPost</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">updatedPost</span><span
  class=\"p\">.</span><span class=\"nx\">Title</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In the above example, we have
  read the response Body which can be accessed as the <code>Body</code> field in the
  <a href=\"https://pkg.go.dev/net/http#Response\">Response</a> object via the <code>resp</code>
  variable. So, the function will return the <code>respBody</code> as a string or
  an error if any. Then using this string object, we can use the <a href=\"https://pkg.go.dev/encoding/json#Unmarshal\">json.Unmarshal</a>
  function to send this string and create this <code>updatedPost</code> object of
  Post struct. This method will mutate this object as we have passed it by reference
  indicated by <code>&amp;updatedPost</code> . So, this will do two things, one update
  / mutate the <code>updatedPost</code> instance from the <code>respBody</code> and
  give any error if any arises during the <a href=\"https://developer.mozilla.org/en-US/docs/Glossary/Deserialization\">deserialization</a>
  of the response.</p>\n<p>Now, we have the object in golang from the response bytes,
  we can use it as per requirements.</p>\n<p>So, that is the example in the entirety.</p>\n<p>Let’s
  simplify the steps which are similar to the POST/PUT method as well</p>\n<ul>\n<li>\n<p>Define/construct
  URL</p>\n</li>\n<li>\n<p>Marshal the object into JSON string as the request body</p>\n</li>\n<li>\n<p>Construct
  the request object (method, URL, and the body)</p>\n</li>\n<li>\n<p>Send the request
  with the default client</p>\n</li>\n<li>\n<p>Read the response and deserialize/unmarshal</p>\n</li>\n<li>\n<p>Access
  the object in golang</p>\n</li>\n</ul>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;log&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Post</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">     </span><span class=\"kt\">int</span><span
  class=\"w\">    </span><span class=\"s\">`json:&quot;id,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Title</span><span
  class=\"w\">  </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;title,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Body</span><span class=\"w\">   </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;body,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">UserId</span><span
  class=\"w\"> </span><span class=\"kt\">int</span><span class=\"w\">    </span><span
  class=\"s\">`json:&quot;userId,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">// define URL to hit the API</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">baseURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://jsonplaceholder.typicode.com&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">postId</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">4</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">postURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Sprintf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s/posts/%d&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">baseURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">postId</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">    </span><span class=\"c1\">// define the body -&gt; with the field
  to update</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">reqBody</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`{&quot;body&quot;:
  &quot;new body&quot;}`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;New body:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">reqBody</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"c1\">// send a
  new request, with the `PATCH` method, url and the body</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"s\">&quot;PATCH&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">postURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">strings</span><span class=\"p\">.</span><span class=\"nx\">NewReader</span><span
  class=\"p\">(</span><span class=\"nx\">reqBody</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">// set the
  header content type to json</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">.</span><span class=\"nx\">Header</span><span
  class=\"p\">.</span><span class=\"nx\">Set</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Content-Type&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;application/json&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">client</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">Client</span><span
  class=\"p\">{}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">client</span><span class=\"p\">.</span><span
  class=\"nx\">Do</span><span class=\"p\">(</span><span class=\"nx\">req</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">defer</span><span
  class=\"w\"> </span><span class=\"nx\">resp</span><span class=\"p\">.</span><span
  class=\"nx\">Body</span><span class=\"p\">.</span><span class=\"nx\">Close</span><span
  class=\"p\">()</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Response status code:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">StatusCode</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Response Status:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Status</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span>\n<span
  class=\"w\">    </span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">updatedPost</span><span class=\"w\"> </span><span class=\"nx\">Post</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">respBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span>\n<span class=\"w\">    </span><span
  class=\"c1\">// convert the response json bytes to Post object in golang</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"p\">=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span
  class=\"nx\">respBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"o\">&amp;</span><span class=\"nx\">updatedPost</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"k\">if</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">log</span><span class=\"p\">.</span><span class=\"nx\">Fatal</span><span
  class=\"p\">(</span><span class=\"nx\">err</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"nx\">updatedPost</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nx\">updatedPost</span><span
  class=\"p\">.</span><span class=\"nx\">Title</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>New
  body: {&quot;body&quot;: &quot;new body&quot;}\nResponse status code: 200\nResponse
  Status: 200 OK\n{4 eum et est occaecati new body 1}\n\ntitle: eum et est occaecati
  \                                                                        \nbody:
  new body\nid: 4\nuser id: 1\n</pre></div>\n\n</pre>\n\n<p>As you can see, it has
  only updated the <code>body</code> and has not updated the other fields.</p>\n<p>If
  you would have sent a similar body with a <code>PUT</code> method, the results would
  have been different. That would have been dependent on the implementation of the
  API of course, but if there is only a few fields in the request body for a PUT method,
  it would have replaced the values with the empry values which are not present in
  the request body.</p>\n<p>That is the difference between a <code>PUT</code> and
  a <code>PATCH</code> method, the <code>PATCH</code> method, ideally should only
  update the fields of the entity that are mentioned in the request body, whereas
  the <code>PUT</code> method has to update the entire resource whether the fields
  are provided or not. Again, the implementation of these API on the server plays
  a vital role in how the behavior defers and the method in itself would perform.</p>\n<p>This
  is also called as <code>JSON Merge Patch</code></p>\n<h2>JSON Merge PATCH</h2>\n<p>The
  above API is implementing a <a href=\"https://datatracker.ietf.org/doc/html/rfc7386\">Merge
  PATCH</a> which is to say, merge the changes in the actual entity.</p>\n<p>Let’s
  say there is a Blog post Entity on a Server, you have a post that you are writing
  as an author. The post has an id of <code>4</code> let’s say and you are constantly
  changing the body of the post.</p>\n<p>So, you don’t want to send the <code>title</code>
  or <code>author_id</code> or any field that is not changing from the client again
  and again while saving, so the <code>MERGE PATCH</code> endpoint will be helpful
  in that case, where the client only sends the required fields to be updated.</p>\n<p>In
  this example, the user would only send the <code>body</code> of the post to the
  API every time it makes changes or saves the draft. In some cases, it might also
  want to change the title, so it will include the title, but not all the fields.
  The API will know it is a <code>PATCH</code> request and the content type is <code>json</code>
  so it will only change or update the fields that are provided in the request body
  to the actual entity in the database or wherever it is stored on the server.</p>\n<p>So,
  that is what is the JSON Merge PATCH or Merge PATCH in general. JSON Merge PATCH
  is specific to the JSON-based document APIs.</p>\n<p>Below is the example, the same
  steps but a different endpoint. A user API that I have specifically created to demonstrate
  the difference in PUT vs Merge PATCH vs JSON PATCH requests.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">User</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">    </span><span class=\"kt\">int</span><span
  class=\"w\">    </span><span class=\"s\">`json:&quot;id,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Name</span><span
  class=\"w\">  </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;name,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Email</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;email,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Roles</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;roles,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nx\">baseURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">userID</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">2</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">apiURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Sprintf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s/users/?id=%d&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">baseURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">userID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span>\n<span class=\"w\">    </span><span class=\"nx\">userObj</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">User</span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">Name</span><span class=\"p\">:</span><span class=\"w\">  </span><span
  class=\"s\">&quot;dummy name&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">Roles</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;dummy role&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">jsonPatchBody</span><span class=\"w\"> </span><span
  class=\"p\">[]</span><span class=\"kt\">byte</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Marshal</span><span class=\"p\">(</span><span
  class=\"nx\">userObj</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">    </span>\n<span
  class=\"w\">    </span><span class=\"c1\">// OR directly define the json as string</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">//jsonMergePatchBody
  := `{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">//
  \   &quot;name&quot;: &quot;new dummy name&quot;,</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"c1\">//    &quot;roles&quot;: &quot;new dummy
  role&quot;</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"c1\">//}`</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"s\">&quot;PATCH&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">strings</span><span class=\"p\">.</span><span class=\"nx\">NewReader</span><span
  class=\"p\">(</span><span class=\"nx\">jsonMergePatchBody</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">.</span><span class=\"nx\">Header</span><span class=\"p\">.</span><span
  class=\"nx\">Set</span><span class=\"p\">(</span><span class=\"s\">&quot;Content-Type&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;application/json&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">user</span><span class=\"w\"> </span><span class=\"nx\">User</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">user</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched User&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">user</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Name:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">.</span><span class=\"nx\">Name</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Email:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">user</span><span class=\"p\">.</span><span class=\"nx\">Email</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Bio:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">user</span><span class=\"p\">.</span><span
  class=\"nx\">Bio</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Roles:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">.</span><span class=\"nx\">Roles</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Original
  User with ID 2</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>{&quot;id&quot;:2,&quot;name&quot;:&quot;dummy&quot;,&quot;email&quot;:&quot;dummyyummy@user.com&quot;,&quot;bio&quot;:&quot;empty
  bio&quot;,&quot;roles&quot;:&quot;read&quot;}\n\nId: 2\nName: dummy\nEmail: dummyyummy@user.com\nBio:
  empty bio\nRoles: read\n</pre></div>\n\n</pre>\n\n<p>Output of the program</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>Request
  Body: {&quot;name&quot;:&quot;dummy name&quot;,&quot;roles&quot;:&quot;dummy role&quot;}\n\n{&quot;id&quot;:2,&quot;name&quot;:&quot;dummy
  name&quot;,&quot;email&quot;:&quot;dummyyummy@user.com&quot;,&quot;bio&quot;:&quot;empty
  bio&quot;,&quot;roles&quot;:&quot;dummy role&quot;}\n\nUpdated/Patched User {2 dummy
  name dummyyummy@user.com empty bio dummy role}\n\nId: 2\nName: dummy name\nEmail:
  dummyyummy@user.com\nBio: empty bio\nRoles: dummy role\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, the only fields that will be updated are <code>name</code> and
  <code>roles</code> , since the API is implemented to only update the fields provided
  in the json merge patch document (request body).</p>\n<p>As you can see that, only
  the <code>name</code> and <code>roles</code> are changed. The name was <code>dummy</code>
  that changed to <code>dummy name</code> and role changed from <code>read</code>
  to <code>dummy role</code> .</p>\n<p>Now, let’s see the same request but with the
  PUT method on it.</p>\n<p>Before we hit this API however, let’s note what the user
  with ID 2 is</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nt\">&quot;id&quot;</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">2</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"nt\">&quot;name&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;dummy name&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nt\">&quot;email&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummyyummy@user.com&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nt\">&quot;bio&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;empty bio&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nt\">&quot;roles&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy role&quot;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>This
  is the result of our recent patch request. Now, we will send a PUT request to the
  same user with a different body.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">User</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">    </span><span class=\"kt\">int</span><span
  class=\"w\">    </span><span class=\"s\">`json:&quot;id,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Name</span><span
  class=\"w\">  </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;name,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Email</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;email,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Bio</span><span
  class=\"w\">   </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;bio,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Roles</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;roles,omitempt&quot;`</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"nx\">baseURL</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">userID</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">2</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">apiURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Sprintf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s/users/?id=%d&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">baseURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">userID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span>\n<span class=\"w\">\t</span><span class=\"nx\">userObj</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">User</span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">Name</span><span class=\"p\">:</span><span class=\"w\">  </span><span
  class=\"s\">&quot;not a dummy name&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">Roles</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s\">&quot;not a dummy role&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"kd\">var</span><span
  class=\"w\"> </span><span class=\"nx\">jsonPatchBody</span><span class=\"w\"> </span><span
  class=\"p\">[]</span><span class=\"kt\">byte</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Marshal</span><span class=\"p\">(</span><span
  class=\"nx\">userObj</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Request Body:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"c1\">//jsonPatchBody := `{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"c1\">//    &quot;name&quot;: &quot;dummy&quot;,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"c1\">//    &quot;roles&quot;:
  &quot;new dummy role&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"c1\">//}`</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"s\">&quot;PUT&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">strings</span><span
  class=\"p\">.</span><span class=\"nx\">NewReader</span><span class=\"p\">(</span><span
  class=\"nb\">string</span><span class=\"p\">(</span><span class=\"nx\">jsonPatchBody</span><span
  class=\"p\">)))</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">.</span><span class=\"nx\">Header</span><span
  class=\"p\">.</span><span class=\"nx\">Set</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Content-Type&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;application/merge-patch+json&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">client</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">http</span><span class=\"p\">.</span><span class=\"nx\">Client</span><span
  class=\"p\">{}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">resp</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">client</span><span class=\"p\">.</span><span
  class=\"nx\">Do</span><span class=\"p\">(</span><span class=\"nx\">req</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">user</span><span class=\"w\"> </span><span class=\"nx\">User</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">user</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched User&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">user</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Name:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">.</span><span class=\"nx\">Name</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Bio:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">user</span><span class=\"p\">.</span><span class=\"nx\">Bio</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Email:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">user</span><span class=\"p\">.</span><span
  class=\"nx\">Email</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Roles:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">.</span><span class=\"nx\">Roles</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Output:</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>Request
  Body: {&quot;name&quot;:&quot;not a dummy name&quot;,&quot;roles&quot;:&quot;not
  a dummy role&quot;}\n\n{&quot;id&quot;:2,&quot;name&quot;:&quot;not a dummy name&quot;,&quot;email&quot;:&quot;&quot;,&quot;bio&quot;:&quot;&quot;,&quot;roles&quot;:&quot;not
  a dummy role&quot;}\n\nUpdated/Patched User {2 not a dummy name   not a dummy role}\n\nId:
  2\nName: not a dummy name\nBio:\nEmail:\nRoles: not a dummy role\n</pre></div>\n\n</pre>\n\n<p>As
  you can see the <code>name</code> and <code>roles</code> are updated, however, the
  <code>bio</code> and <code>email</code> fields are empty. Since we only said to
  update the <code>name</code> and <code>roles</code> fields, but it was a <code>PUT</code>
  request, it expects all the fields, and if any of the fields are missing, it will
  consider them as empty and update those as well.</p>\n<p>So, the difference might
  be crystal clear now. When to use <code>PATCH</code> and when to avoid <code>PUT</code>.</p>\n<ul>\n<li>\n<p>When
  you have a large set of updates, preference could be given to PUT</p>\n</li>\n<li>\n<p>If
  you have very specific fields to update and minimal fields PATCH is recommended</p>\n</li>\n</ul>\n<p>There
  is another type of PATCH specifically designed for JSON APIs or JSON Documents APIs.</p>\n<h2>JSON
  PATCH</h2>\n<p>The <a href=\"https://datatracker.ietf.org/doc/html/rfc6902/\">JSON
  PATCH</a> is a specification in which we can define what operations to perform on
  which fields, or path of the fields to replace, move, or copy to.</p>\n<blockquote>\n<p>A
  JSON Patch document is a JSON document that represents an array of objects. Each
  object represents a single operation to be applied to the target JSON document.</p>\n</blockquote>\n<p>As
  it takes these operations, it applies them sequentially and hence it won’t replace
  all the fields for the entity, as that is the expected behavior of the PATCH method.
  In other words, it would only apply changes to the fields and the related fields
  provided in the json patch document (request body).</p>\n<p>There are a few operations
  that you can perform with this json patch method, and provide the instructions accordingly
  for individual operations in the JSON PATCH document.</p>\n<p>Operations</p>\n<ol>\n<li>\n<p>add</p>\n</li>\n<li>\n<p>remove</p>\n</li>\n<li>\n<p>replace</p>\n</li>\n<li>\n<p>move</p>\n</li>\n<li>\n<p>copy</p>\n</li>\n<li>\n<p>test</p>\n</li>\n</ol>\n<p>So,
  for each of the operations, a high-level definition can be considered as:</p>\n<ul>\n<li>\n<p>To
  add a field you can specify the operation as <code>add</code> , the path as the
  field to be added, and the value as the actual value to be added</p>\n</li>\n<li>\n<p>To
  remove a field, you can specify the operation as <code>remove</code> , and the path
  as the field to remove</p>\n</li>\n<li>\n<p>To replace a field, you can specify
  the operation as <code>replace</code>, the path as the field to be updated/replaced,
  and the value of the actual value to be added</p>\n</li>\n<li>\n<p>To move a field,
  you can specify the operation as <code>move</code>, the <strong>form</strong> as
  the field to be updated/moved from and the path to the field the from value should
  be moved to.</p>\n</li>\n<li>\n<p>To copy a field, you can specify the operation
  as <code>copy</code>, the form as the field to update/copied from, and the path
  to the field to which the value should be copied to.</p>\n</li>\n<li>\n<p>The test
  operation is a bit different as it is used for comparison of a <code>path</code>
  value to the actual value specified in the object. It might return true or false,
  but not actually return it might be used as a checkpoint for continuing with the
  operation in the document.</p>\n</li>\n</ul>\n<p>In this example, we are creating
  a similar patch request, but using this json patch document kind of structure.</p>\n<h3>Construct
  the json-patch document</h3>\n<p>We need to construct the JSON PATCH document, which
  is a list of operations (instructions). Each instruction has an operation <code>op</code>
  field that defines what operation need to be performed among the available 6 operations.
  The <code>path</code> field is the field that we need to update, it is called path
  because, we need to provide the path to the json field, in case of nested json fields.
  The <code>value</code> and <code>from</code> fields are a bit dependent on the operation
  as <code>value</code> is not required in <code>remove</code> operation and <code>from</code>
  is only used in <code>move</code> and <code>copy</code> operations.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nx\">jsonPatchBody</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">`[</span>\n<span class=\"s\">    {</span>\n<span
  class=\"s\">        &quot;op&quot;: &quot;replace&quot;,</span>\n<span class=\"s\">
  \       &quot;path&quot;: &quot;/name&quot;,</span>\n<span class=\"s\">        &quot;value&quot;:
  &quot;new dummy name&quot;</span>\n<span class=\"s\">    },</span>\n<span class=\"s\">
  \   {</span>\n<span class=\"s\">        &quot;op&quot;: &quot;replace&quot;,</span>\n<span
  class=\"s\">        &quot;path&quot;: &quot;/roles&quot;,</span>\n<span class=\"s\">
  \       &quot;value&quot;: &quot;new dummy role&quot;</span>\n<span class=\"s\">
  \   },</span>\n<span class=\"s\">]`</span><span class=\"w\"></span>\n\n\n<span class=\"c1\">//
  OR</span><span class=\"w\"></span>\n<span class=\"c1\">// construct a go type and
  then marshal it</span><span class=\"w\"></span>\n\n<span class=\"kd\">type</span><span
  class=\"w\"> </span><span class=\"nx\">JSONPatch</span><span class=\"w\"> </span><span
  class=\"kd\">struct</span><span class=\"w\"> </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Op</span><span
  class=\"w\">    </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;op&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">Path</span><span class=\"w\">  </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;path&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">From</span><span class=\"w\">  </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;from,omitempty&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Value</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;value,omitempty&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"nx\">jsonPatchData</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"p\">[]</span><span
  class=\"nx\">JSONPatch</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">Op</span><span class=\"p\">:</span><span
  class=\"w\">    </span><span class=\"s\">&quot;replace&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">Path</span><span
  class=\"p\">:</span><span class=\"w\">  </span><span class=\"s\">&quot;/name&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">Value</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s\">&quot;new dummy name&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">},</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t\t</span><span class=\"nx\">Op</span><span class=\"p\">:</span><span
  class=\"w\">    </span><span class=\"s\">&quot;replace&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">\t\t</span><span class=\"nx\">Path</span><span
  class=\"p\">:</span><span class=\"w\">  </span><span class=\"s\">&quot;/roles&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nx\">Value</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s\">&quot;new dummy role&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"p\">},</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"nx\">jsonPatchBytes</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">json</span><span
  class=\"p\">.</span><span class=\"nx\">Marshal</span><span class=\"p\">(</span><span
  class=\"nx\">jsonPatchData</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Error marshalling JSON:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">return</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span
  class=\"w\"></span>\n\n<span class=\"nx\">jsonPatchBody</span><span class=\"w\">
  </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nb\">string</span><span
  class=\"p\">(</span><span class=\"nx\">jsonPatchBytes</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In the above example, we have
  defined the two ways to create the body for a json patch request. One is directly
  writing json string (not recommended). Also, the other way is to define a struct
  as JSONPatch and construct the object for the operations in the json patch request.
  The JSONPatch object has the 4 fields as mentioned earlier. In this case, we are
  only using the <code>replace</code> operation, which will update the value in the
  provided path field to the value provided in the value field.</p>\n<p>In this specific
  example, we have 2 operations to perform, first to update the <code>name</code>field
  and the second to update the <code>roles</code> field. We simply use the <code>json.Marshal</code>
  method to convert the object into bytes. We then cast the slice of bytes into string
  and that becomes a final <code>jsonPatchBody</code> for the request that we will
  send.</p>\n<p><strong><mark>NOTE: This API </mark></strong> <code>https://dummy-json-patch.netlify.app/.netlify/functions/users/</code>
  <strong><mark> is created by me, it only implements the </mark></strong> <code>replace</code>
  <strong><mark> operation, as it makes sense to only have an update operation on
  the user table endpoint, If the model is more like a json document-like structure,
  the other operations might make sense. I might implement those in a separate article,
  while we explore the implementation of the PATCH Request in the CRUD API Section
  of this series soon.</mark></strong></p>\n<h3>Constructing and Sending the Request</h3>\n<p>We
  do the same stuff as the normal merge patch request, that is provide the <code>PATCH</code>
  as the request method, the <code>apiURL</code> and the <code>strings.NewReader(jsonPatchBody)</code>
  into the body of the request.</p>\n<p>However, the only difference is that we need
  to set the <code>Content-Type</code> header as <code>application/json-patch+json</code>
  . This is defined in the <a href=\"https://datatracker.ietf.org/doc/html/rfc6902/#section-1\">specification</a>,
  it would make the client as well as the server understand the request body. This
  can also be used by the server to differentiate between a standard patch (which
  can be referred to as a &quot;Merge PATCH&quot; for JSON APIs) and a JSON PATCH
  request, as shown here.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"s\">&quot;PATCH&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">strings</span><span
  class=\"p\">.</span><span class=\"nx\">NewReader</span><span class=\"p\">(</span><span
  class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"nx\">req</span><span class=\"p\">.</span><span class=\"nx\">Header</span><span
  class=\"p\">.</span><span class=\"nx\">Set</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Content-Type&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"s\">&quot;application/json-patch+json&quot;</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Once, we have created the request
  object we then simply can create the client and send the request.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"nx\">resp</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">var</span><span class=\"w\"> </span><span class=\"nx\">updatedUser</span><span
  class=\"w\"> </span><span class=\"nx\">User</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">resBody</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">io</span><span class=\"p\">.</span><span
  class=\"nx\">ReadAll</span><span class=\"p\">(</span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"nb\">string</span><span
  class=\"p\">(</span><span class=\"nx\">resBody</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">updatedUser</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"k\">if</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">!=</span><span class=\"w\"> </span><span class=\"kc\">nil</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched User&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedUser</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Once,
  we have the response from the request, we just do the normal response processing
  and unmarshal the response string into a golang object type (passing by reference)
  which is provided as <code>&amp;updatedUser</code> in the parameter.</p>\n<p>This
  will give us the expected result back in the updated User object.</p>\n<p>The below
  code is the entire JSON PATCH request:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">User</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">    </span><span class=\"kt\">int</span><span
  class=\"w\">    </span><span class=\"s\">`json:&quot;id&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Name</span><span class=\"w\">  </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;name&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Email</span><span
  class=\"w\"> </span><span class=\"kt\">string</span><span class=\"w\"> </span><span
  class=\"s\">`json:&quot;email&quot;`</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nx\">Bio</span><span class=\"w\">   </span><span class=\"kt\">string</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;bio&quot;`</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">Roles</span><span class=\"w\"> </span><span
  class=\"kt\">string</span><span class=\"w\"> </span><span class=\"s\">`json:&quot;roles&quot;`</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n\n<span
  class=\"kd\">func</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"p\">()</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">baseURL</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nx\">userID</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"mi\">2</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">apiURL</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Sprintf</span><span class=\"p\">(</span><span
  class=\"s\">&quot;%s/users/?id=%d&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">baseURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">userID</span><span class=\"p\">)</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`[</span>\n<span
  class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;: &quot;replace&quot;,</span>\n<span
  class=\"s\">            &quot;path&quot;: &quot;/name&quot;,</span>\n<span class=\"s\">
  \           &quot;value&quot;: &quot;new dummy name&quot;</span>\n<span class=\"s\">
  \       },</span>\n<span class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;:
  &quot;replace&quot;,</span>\n<span class=\"s\">            &quot;path&quot;: &quot;/roles&quot;,</span>\n<span
  class=\"s\">            &quot;value&quot;: &quot;new dummy role&quot;</span>\n<span
  class=\"s\">        },</span>\n<span class=\"s\">    ]`</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"s\">&quot;PATCH&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">strings</span><span class=\"p\">.</span><span class=\"nx\">NewReader</span><span
  class=\"p\">(</span><span class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">.</span><span class=\"nx\">Header</span><span class=\"p\">.</span><span
  class=\"nx\">Set</span><span class=\"p\">(</span><span class=\"s\">&quot;Content-Type&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;application/json-patch+json&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">user</span><span class=\"w\"> </span><span class=\"nx\">User</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">user</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched User&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">user</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Name:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">.</span><span class=\"nx\">Name</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">fmt</span><span
  class=\"p\">.</span><span class=\"nx\">Println</span><span class=\"p\">(</span><span
  class=\"s\">&quot;Bio:&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">user</span><span class=\"p\">.</span><span class=\"nx\">Bio</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Email:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">user</span><span class=\"p\">.</span><span
  class=\"nx\">Email</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;Roles:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">.</span><span class=\"nx\">Roles</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Original
  User (id=2) before the request</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nt\">&quot;id&quot;</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">2</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">  </span><span class=\"nt\">&quot;name&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;dummy name&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nt\">&quot;email&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummyyummy@user.com&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"nt\">&quot;bio&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;empty bio&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">  </span><span class=\"nt\">&quot;roles&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy role&quot;</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Output
  of the Program (JSON PATCH Request)</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>{&quot;id&quot;:2,&quot;name&quot;:&quot;new
  dummy name&quot;,&quot;email&quot;:&quot;dummyyummy@user.com&quot;,&quot;bio&quot;:
  &quot;empty bio&quot;, &quot;roles&quot;:&quot;new dummy role&quot;}     \n\nUpdated/Patched
  User {2 new dummy name dummyyummy@user.com empty bio new dummy role} \n\nId: 2\nName:
  new dummy name\nEmail: dummyyummy@user.com\nBio: empty bio\nRoles: read\n</pre></div>\n\n</pre>\n\n<p>So,
  as expected, it only updates the <code>name</code> and the <code>roles</code> field
  in the user object with id = 2. The patch document has the replace operation that
  updates the name and the roles fields from the existing field.</p>\n<p>To add more
  examples for <code>add</code> , <code>move</code>,<code>copy</code>, <code>remove</code>
  and <code>test</code> operations.</p>\n<p>PS: I added an endpoint for documents,
  which will include all the crud operations with all the JSON PATCH operations added
  as well.</p>\n<p>Let’s take a simple example for each operation</p>\n<h3>Add operation</h3>\n<p>Let’s
  take an example for a document table at the <code>https://dummy-json-patch.netlify.app/.netlify/functions/documents</code>
  endpoint, the document is a table with the columns like data and the id as the index.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">  </span><span
  class=\"s\">&quot;data&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"p\">{},</span><span class=\"w\"></span>\n<span class=\"w\">
  \ </span><span class=\"s\">&quot;id&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">4</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  initial data is empty for a document with an ID is 4.</p>\n<p>We simply use the
  operation of <code>add</code> that will add the value of the <code>name</code> in
  the <code>data</code> object with the field name as <code>name</code> and the value
  as <code>dummy</code> .</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Document</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">   </span><span class=\"kt\">int</span><span
  class=\"w\">         </span><span class=\"s\">`json:&quot;id&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Data</span><span
  class=\"w\"> </span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;data&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions/documents/?id=4&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`[</span>\n<span
  class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;: &quot;add&quot;,</span>\n<span
  class=\"s\">            &quot;path&quot;: &quot;/name&quot;,</span>\n<span class=\"s\">
  \           &quot;value&quot;: &quot;dummy&quot;</span>\n<span class=\"s\">        }</span>\n<span
  class=\"s\">    ]`</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"s\">&quot;PATCH&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">strings</span><span
  class=\"p\">.</span><span class=\"nx\">NewReader</span><span class=\"p\">(</span><span
  class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">.</span><span
  class=\"nx\">Header</span><span class=\"p\">.</span><span class=\"nx\">Set</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Content-Type&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;application/json-patch+json&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">updatedDoc</span><span class=\"w\"> </span><span class=\"nx\">Document</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">updatedDoc</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched Document&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;document:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">.</span><span class=\"nx\">Data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>This
  is a simple add operation, as you would see that the data now has the name field
  populated after the request was sent.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"s\">&quot;data&quot;</span><span class=\"p\">:{</span><span
  class=\"s\">&quot;name&quot;</span><span class=\"p\">:</span><span class=\"s\">&quot;dummy&quot;</span><span
  class=\"p\">},</span><span class=\"s\">&quot;id&quot;</span><span class=\"p\">:</span><span
  class=\"mi\">4</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"nx\">Updated</span><span class=\"o\">/</span><span class=\"nx\">Patched</span><span
  class=\"w\"> </span><span class=\"nx\">Document</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"mi\">4</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"nx\">name</span><span class=\"p\">:</span><span
  class=\"nx\">dummy</span><span class=\"p\">]}</span><span class=\"w\"></span>\n<span
  class=\"nx\">id</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">4</span><span class=\"w\"></span>\n<span class=\"nx\">document</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"nx\">name</span><span class=\"p\">:</span><span
  class=\"nx\">dummy</span><span class=\"p\">]</span><span class=\"w\"> </span>\n</pre></div>\n\n</pre>\n\n<h3>Remove
  Operation</h3>\n<p>We can take an example for remove, let’s remove the <code>name</code>
  in the <code>data</code> field that we added previously.</p>\n<p>The data before
  the request looks like this:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"s\">&quot;data&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"s\">&quot;name&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s\">&quot;dummy&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"p\">},</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"s\">&quot;id&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"mi\">4</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Now,
  we will construct the request, we can do it with the json marshaling by constructing
  the object, but I am keeping it simple and shorter by directly constructing the
  json body.</p>\n<p>For the <code>remove</code> operation, we only need the <code>path</code>
  field to be added, in order to remove the field from the object.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Document</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">   </span><span class=\"kt\">int</span><span
  class=\"w\">         </span><span class=\"s\">`json:&quot;id&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Data</span><span
  class=\"w\"> </span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;data&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions/documents/?id=4&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`[</span>\n<span
  class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;: &quot;remove&quot;,</span>\n<span
  class=\"s\">            &quot;path&quot;: &quot;/name&quot;</span>\n<span class=\"s\">
  \       }</span>\n<span class=\"s\">    ]`</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"s\">&quot;PATCH&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">strings</span><span class=\"p\">.</span><span class=\"nx\">NewReader</span><span
  class=\"p\">(</span><span class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">.</span><span class=\"nx\">Header</span><span class=\"p\">.</span><span
  class=\"nx\">Set</span><span class=\"p\">(</span><span class=\"s\">&quot;Content-Type&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;application/json-patch+json&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">updatedDoc</span><span class=\"w\"> </span><span class=\"nx\">Document</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">updatedDoc</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched Document&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;document:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">.</span><span class=\"nx\">Data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>{&quot;data&quot;:{},&quot;id&quot;:4}
  \               \n\nUpdated/Patched Document {4 map[]}\nid: 4                             \ndocument:
  map[]                   \n</pre></div>\n\n</pre>\n\n<p>So, as you can see the updated
  data has no fields.</p>\n<h3>Copy Operation</h3>\n<p>Now, we can also perform the
  copy operation, which can copy the field to another field, this instruction requires
  the <code>from</code> and the <code>path</code> . The <code>from</code> is the path
  to the field to copy it from, and the <code>path</code> is the path to the field
  to copy the field value.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">[</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;op&quot;</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s2\">&quot;copy&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;from&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;/name&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;path&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;/new_name&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">]</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  above json path operation is for copying the value of the field <code>name</code>
  to the field <code>new_name</code> i.e. if the original document looks likes this:</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;data&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;name&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;dummy&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;new_name&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;hello&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">},</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;id&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"mi\">4</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  after this the value of <code>new_name</code> will be <code>dummy</code> after the
  operation</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nt\">&quot;data&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \       </span><span class=\"nt\">&quot;name&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;new_name&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">},</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;id&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"mi\">4</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  let’s take a look at the code where we will perform this operation with this API
  endpoint and constructing the json body.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Document</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">   </span><span class=\"kt\">int</span><span
  class=\"w\">         </span><span class=\"s\">`json:&quot;id&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Data</span><span
  class=\"w\"> </span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;data&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions/documents/?id=4&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`[</span>\n<span
  class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;: &quot;copy&quot;,</span>\n<span
  class=\"s\">            &quot;from&quot;: &quot;/name&quot;,</span>\n<span class=\"s\">
  \           &quot;path&quot;: &quot;/new_name&quot;</span>\n<span class=\"s\">        }</span>\n<span
  class=\"s\">    ]`</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"s\">&quot;PATCH&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">strings</span><span
  class=\"p\">.</span><span class=\"nx\">NewReader</span><span class=\"p\">(</span><span
  class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">.</span><span
  class=\"nx\">Header</span><span class=\"p\">.</span><span class=\"nx\">Set</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Content-Type&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;application/json-patch+json&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">updatedDoc</span><span class=\"w\"> </span><span class=\"nx\">Document</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">updatedDoc</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched Document&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;document:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">.</span><span class=\"nx\">Data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>In
  the above example, we are creating the json payload for copying the value of <code>name</code>
  to the <code>new_name</code> field.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>{&quot;data&quot;:{&quot;name&quot;:&quot;dummy&quot;,&quot;new_name&quot;:&quot;dummy&quot;},&quot;id&quot;:4}\nUpdated/Patched
  Document {4 map[name:dummy new_name:dummy]}\nid: 4\ndocument: map[name:dummy new_name:dummy]\n</pre></div>\n\n</pre>\n\n<h3>Move
  Operation</h3>\n<p>Now, we can also perform the move operation, which can move the
  field to another field, this instruction requires the <code>from</code> and the
  <code>path</code> . The <code>from</code> is the path to the field to move it from,
  and the <code>path</code> is the path to the field to move.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">[</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;op&quot;</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s2\">&quot;move&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;from&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;/name&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;path&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;/user/description&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">]</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  above json path operation is for moving the value of the field <code>name</code>
  to the field <code>/user/description</code> i.e. if the original document looks
  like this:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nt\">&quot;name&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;new_name&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;user&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;description&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;hello&quot;</span><span class=\"w\"></span>\n<span
  class=\"w\">    </span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  after the operation, the value of <code>name</code> will be <code>dummy</code>.
  The name field is moved not replaced, that is the difference of this operation.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nt\">&quot;new_name&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;user&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;description&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  let’s take a look at the code where we will construct the payload body and perform
  the json patch request with the move operations.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Document</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">   </span><span class=\"kt\">int</span><span
  class=\"w\">         </span><span class=\"s\">`json:&quot;id&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Data</span><span
  class=\"w\"> </span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;data&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions/documents/?id=4&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`[</span>\n<span
  class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;: &quot;move&quot;,</span>\n<span
  class=\"s\">            &quot;from&quot;: &quot;/name&quot;,</span>\n<span class=\"s\">
  \           &quot;path&quot;: &quot;/user/description&quot;</span>\n<span class=\"s\">
  \       }</span>\n<span class=\"s\">    ]`</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">err</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">NewRequest</span><span class=\"p\">(</span><span
  class=\"s\">&quot;PATCH&quot;</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">apiURL</span><span class=\"p\">,</span><span class=\"w\">
  </span><span class=\"nx\">strings</span><span class=\"p\">.</span><span class=\"nx\">NewReader</span><span
  class=\"p\">(</span><span class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">req</span><span
  class=\"p\">.</span><span class=\"nx\">Header</span><span class=\"p\">.</span><span
  class=\"nx\">Set</span><span class=\"p\">(</span><span class=\"s\">&quot;Content-Type&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"s\">&quot;application/json-patch+json&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">updatedDoc</span><span class=\"w\"> </span><span class=\"nx\">Document</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">updatedDoc</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched Document&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;document:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">.</span><span class=\"nx\">Data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>The
  above code simply takes in the json string as a payload for the API with the move
  operation. Hence the name field’s value is swapped to the value of <code>dummy</code>
  of the <code>/user/description</code> field.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"s\">&quot;data&quot;</span><span class=\"p\">:{</span><span
  class=\"s\">&quot;new_name&quot;</span><span class=\"p\">:</span><span class=\"s\">&quot;dummy&quot;</span><span
  class=\"p\">,</span><span class=\"s\">&quot;user&quot;</span><span class=\"p\">:{</span><span
  class=\"s\">&quot;description&quot;</span><span class=\"p\">:</span><span class=\"s\">&quot;dummy&quot;</span><span
  class=\"p\">}},</span><span class=\"s\">&quot;id&quot;</span><span class=\"p\">:</span><span
  class=\"mi\">4</span><span class=\"p\">}</span><span class=\"w\"></span>\n<span
  class=\"nx\">Updated</span><span class=\"o\">/</span><span class=\"nx\">Patched</span><span
  class=\"w\"> </span><span class=\"nx\">Document</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"mi\">4</span><span class=\"w\"> </span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"nx\">new_name</span><span class=\"p\">:</span><span
  class=\"nx\">dummy</span><span class=\"w\"> </span><span class=\"nx\">user</span><span
  class=\"p\">:</span><span class=\"kd\">map</span><span class=\"p\">[</span><span
  class=\"nx\">description</span><span class=\"p\">:</span><span class=\"nx\">dummy</span><span
  class=\"p\">]]}</span><span class=\"w\"></span>\n<span class=\"nx\">id</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"mi\">4</span><span class=\"w\"></span>\n<span
  class=\"nx\">document</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"kd\">map</span><span class=\"p\">[</span><span class=\"nx\">new_name</span><span
  class=\"p\">:</span><span class=\"nx\">dummy</span><span class=\"w\"> </span><span
  class=\"nx\">user</span><span class=\"p\">:</span><span class=\"kd\">map</span><span
  class=\"p\">[</span><span class=\"nx\">description</span><span class=\"p\">:</span><span
  class=\"nx\">dummy</span><span class=\"p\">]]</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<h3>Test
  Operation</h3>\n<p>The test operation requires the <code>path</code> field with
  the value of the field to test, and the <code>value</code> field as the value to
  compare with. This is a test, so we can use it as a flag to proceed ahead with the
  rest of the operations or abort.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">[</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;op&quot;</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s2\">&quot;test&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;path&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;/new_name&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;value&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;testing value&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">},</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;op&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;add&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;path&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;/user/id&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;value&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"mi\">123</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">]</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>if
  the original document looks like this:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">
  \   </span><span class=\"nt\">&quot;new_name&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;user&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;description&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Then
  after applying the test operation, we can find that the <code>new_name</code> has
  the value is <code>dummy</code> and not <code>testing value</code> so it will abort
  with an error.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;error&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;test
  failed: values do not match at path new_name&quot;</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>However,
  if we keep the operation like this:</p>\n<p>So, we compare the value of <code>new_name</code>
  field is <code>dummy</code> which it is then we will perform the rest of the operations
  in this case, the <code>add</code> operations will be performed to add the <code>id</code>
  field to the <code>user</code> object with the value <code>123</code>.</p>\n<pre
  class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy' title='copy
  code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">[</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;op&quot;</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"s2\">&quot;test&quot;</span><span class=\"p\">,</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;path&quot;</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"s2\">&quot;/new_name&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;value&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">},</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">{</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;op&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;add&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;path&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;/user/id&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">        </span><span class=\"nt\">&quot;value&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"mi\">123</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">]</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>if
  the original document looks like this:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;new_name&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;dummy&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;user&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;description&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"p\">}</span><span
  class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>Then
  after the operation of test and add, the document becomes:</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"nt\">&quot;new_name&quot;</span><span class=\"p\">:</span><span class=\"w\">
  </span><span class=\"s2\">&quot;dummy&quot;</span><span class=\"p\">,</span><span
  class=\"w\"></span>\n<span class=\"w\">    </span><span class=\"nt\">&quot;user&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span
  class=\"w\">        </span><span class=\"nt\">&quot;description&quot;</span><span
  class=\"p\">:</span><span class=\"w\"> </span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"p\">,</span><span class=\"w\"></span>\n<span class=\"w\">        </span><span
  class=\"nt\">&quot;id&quot;</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">123</span><span class=\"w\"></span>\n<span class=\"w\">    </span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>You
  can observe the <code>id</code> field is added to the <code>user</code> field, as
  the test was successful. Hence, the operation <code>add</code> was performed and
  the request was completed.</p>\n<p>Let’s take a look at the code for doing this
  operation:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>go</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kn\">package</span><span class=\"w\"> </span><span class=\"nx\">main</span><span
  class=\"w\"></span>\n\n<span class=\"kn\">import</span><span class=\"w\"> </span><span
  class=\"p\">(</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;encoding/json&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;fmt&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;io&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;net/http&quot;</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"s\">&quot;strings&quot;</span><span class=\"w\"></span>\n<span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"kd\">type</span><span class=\"w\"> </span><span
  class=\"nx\">Document</span><span class=\"w\"> </span><span class=\"kd\">struct</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">ID</span><span class=\"w\">   </span><span class=\"kt\">int</span><span
  class=\"w\">         </span><span class=\"s\">`json:&quot;id&quot;`</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">Data</span><span
  class=\"w\"> </span><span class=\"kd\">interface</span><span class=\"p\">{}</span><span
  class=\"w\"> </span><span class=\"s\">`json:&quot;data&quot;`</span><span class=\"w\"></span>\n<span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"kd\">func</span><span
  class=\"w\"> </span><span class=\"nx\">main</span><span class=\"p\">()</span><span
  class=\"w\"> </span><span class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">apiURL</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"s\">&quot;https://dummy-json-patch.netlify.app/.netlify/functions/documents/?id=4&quot;</span><span
  class=\"w\"></span>\n\n<span class=\"w\">    </span><span class=\"cm\">/*</span>\n<span
  class=\"cm\">    jsonPatchBody := `[</span>\n<span class=\"cm\">        {</span>\n<span
  class=\"cm\">            &quot;op&quot;: &quot;test&quot;,</span>\n<span class=\"cm\">
  \           &quot;path&quot;: &quot;/new_name&quot;,</span>\n<span class=\"cm\">
  \           &quot;value&quot;: &quot;testing value&quot;</span>\n<span class=\"cm\">
  \       },</span>\n<span class=\"cm\">        {</span>\n<span class=\"cm\">            &quot;op&quot;:
  &quot;add&quot;,</span>\n<span class=\"cm\">            &quot;path&quot;: &quot;/user/id&quot;,</span>\n<span
  class=\"cm\">            &quot;value&quot;: 123</span>\n<span class=\"cm\">        }</span>\n<span
  class=\"cm\">    ]`</span>\n<span class=\"cm\">    */</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">jsonPatchBody</span><span class=\"w\"> </span><span
  class=\"o\">:=</span><span class=\"w\"> </span><span class=\"s\">`[</span>\n<span
  class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;: &quot;test&quot;,</span>\n<span
  class=\"s\">            &quot;path&quot;: &quot;/new_name&quot;,</span>\n<span class=\"s\">
  \           &quot;value&quot;: &quot;dummy&quot;</span>\n<span class=\"s\">        },</span>\n<span
  class=\"s\">        {</span>\n<span class=\"s\">            &quot;op&quot;: &quot;add&quot;,</span>\n<span
  class=\"s\">            &quot;path&quot;: &quot;/user/id&quot;,</span>\n<span class=\"s\">
  \           &quot;value&quot;: 123</span>\n<span class=\"s\">        }</span>\n<span
  class=\"s\">    ]`</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">req</span><span class=\"p\">,</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"nx\">http</span><span class=\"p\">.</span><span
  class=\"nx\">NewRequest</span><span class=\"p\">(</span><span class=\"s\">&quot;PATCH&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">apiURL</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">strings</span><span
  class=\"p\">.</span><span class=\"nx\">NewReader</span><span class=\"p\">(</span><span
  class=\"nx\">jsonPatchBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">req</span><span class=\"p\">.</span><span
  class=\"nx\">Header</span><span class=\"p\">.</span><span class=\"nx\">Set</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Content-Type&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"s\">&quot;application/json-patch+json&quot;</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">client</span><span class=\"w\"> </span><span class=\"o\">:=</span><span
  class=\"w\"> </span><span class=\"o\">&amp;</span><span class=\"nx\">http</span><span
  class=\"p\">.</span><span class=\"nx\">Client</span><span class=\"p\">{}</span><span
  class=\"w\"></span>\n<span class=\"w\">\t</span><span class=\"nx\">resp</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">client</span><span
  class=\"p\">.</span><span class=\"nx\">Do</span><span class=\"p\">(</span><span
  class=\"nx\">req</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"k\">defer</span><span class=\"w\"> </span><span class=\"nx\">resp</span><span
  class=\"p\">.</span><span class=\"nx\">Body</span><span class=\"p\">.</span><span
  class=\"nx\">Close</span><span class=\"p\">()</span><span class=\"w\"></span>\n\n<span
  class=\"w\">\t</span><span class=\"kd\">var</span><span class=\"w\"> </span><span
  class=\"nx\">updatedDoc</span><span class=\"w\"> </span><span class=\"nx\">Document</span><span
  class=\"w\"></span>\n\n<span class=\"w\">\t</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">err</span><span
  class=\"w\"> </span><span class=\"o\">:=</span><span class=\"w\"> </span><span class=\"nx\">io</span><span
  class=\"p\">.</span><span class=\"nx\">ReadAll</span><span class=\"p\">(</span><span
  class=\"nx\">resp</span><span class=\"p\">.</span><span class=\"nx\">Body</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"nb\">string</span><span class=\"p\">(</span><span
  class=\"nx\">resBody</span><span class=\"p\">))</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"p\">=</span><span
  class=\"w\"> </span><span class=\"nx\">json</span><span class=\"p\">.</span><span
  class=\"nx\">Unmarshal</span><span class=\"p\">(</span><span class=\"nx\">resBody</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"o\">&amp;</span><span
  class=\"nx\">updatedDoc</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"k\">if</span><span class=\"w\"> </span><span
  class=\"nx\">err</span><span class=\"w\"> </span><span class=\"o\">!=</span><span
  class=\"w\"> </span><span class=\"kc\">nil</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"w\"></span>\n<span class=\"w\">\t\t</span><span
  class=\"nb\">panic</span><span class=\"p\">(</span><span class=\"nx\">err</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"p\">}</span><span class=\"w\"></span>\n\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;Updated/Patched Document&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">)</span><span class=\"w\"></span>\n<span class=\"w\">\t</span><span
  class=\"nx\">fmt</span><span class=\"p\">.</span><span class=\"nx\">Println</span><span
  class=\"p\">(</span><span class=\"s\">&quot;id:&quot;</span><span class=\"p\">,</span><span
  class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span class=\"p\">.</span><span
  class=\"nx\">ID</span><span class=\"p\">)</span><span class=\"w\"></span>\n<span
  class=\"w\">\t</span><span class=\"nx\">fmt</span><span class=\"p\">.</span><span
  class=\"nx\">Println</span><span class=\"p\">(</span><span class=\"s\">&quot;document:&quot;</span><span
  class=\"p\">,</span><span class=\"w\"> </span><span class=\"nx\">updatedDoc</span><span
  class=\"p\">.</span><span class=\"nx\">Data</span><span class=\"p\">)</span><span
  class=\"w\"></span>\n\n<span class=\"p\">}</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  the above code is an example of performing the test operation.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>json</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"p\">{</span><span class=\"nt\">&quot;data&quot;</span><span class=\"p\">:{</span><span
  class=\"nt\">&quot;new_name&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;user&quot;</span><span class=\"p\">:{</span><span
  class=\"nt\">&quot;description&quot;</span><span class=\"p\">:</span><span class=\"s2\">&quot;dummy&quot;</span><span
  class=\"p\">,</span><span class=\"nt\">&quot;id&quot;</span><span class=\"p\">:</span><span
  class=\"mi\">123</span><span class=\"p\">}},</span><span class=\"nt\">&quot;id&quot;</span><span
  class=\"p\">:</span><span class=\"mi\">4</span><span class=\"p\">}</span><span class=\"w\">
  \                                                                   </span>\n<span
  class=\"err\">Upda</span><span class=\"kc\">te</span><span class=\"err\">d/Pa</span><span
  class=\"kc\">t</span><span class=\"err\">ched</span><span class=\"w\"> </span><span
  class=\"err\">Docume</span><span class=\"kc\">nt</span><span class=\"w\"> </span><span
  class=\"p\">{</span><span class=\"mi\">4</span><span class=\"w\"> </span><span class=\"err\">map</span><span
  class=\"p\">[</span><span class=\"kc\">ne</span><span class=\"err\">w_</span><span
  class=\"kc\">na</span><span class=\"err\">me</span><span class=\"p\">:</span><span
  class=\"err\">dummy</span><span class=\"w\"> </span><span class=\"err\">user</span><span
  class=\"p\">:</span><span class=\"err\">map</span><span class=\"p\">[</span><span
  class=\"err\">descrip</span><span class=\"kc\">t</span><span class=\"err\">io</span><span
  class=\"kc\">n</span><span class=\"p\">:</span><span class=\"err\">dummy</span><span
  class=\"w\"> </span><span class=\"err\">id</span><span class=\"p\">:</span><span
  class=\"mi\">123</span><span class=\"p\">]]}</span><span class=\"w\">                                                             </span>\n<span
  class=\"err\">id</span><span class=\"p\">:</span><span class=\"w\"> </span><span
  class=\"mi\">4</span><span class=\"w\">                                                                                                                                           </span>\n<span
  class=\"err\">docume</span><span class=\"kc\">nt</span><span class=\"p\">:</span><span
  class=\"w\"> </span><span class=\"err\">map</span><span class=\"p\">[</span><span
  class=\"kc\">ne</span><span class=\"err\">w_</span><span class=\"kc\">na</span><span
  class=\"err\">me</span><span class=\"p\">:</span><span class=\"err\">dummy</span><span
  class=\"w\"> </span><span class=\"err\">user</span><span class=\"p\">:</span><span
  class=\"err\">map</span><span class=\"p\">[</span><span class=\"err\">descrip</span><span
  class=\"kc\">t</span><span class=\"err\">io</span><span class=\"kc\">n</span><span
  class=\"p\">:</span><span class=\"err\">dummy</span><span class=\"w\"> </span><span
  class=\"err\">id</span><span class=\"p\">:</span><span class=\"mi\">123</span><span
  class=\"p\">]]</span><span class=\"w\"></span>\n</pre></div>\n\n</pre>\n\n<p>So,
  as you would see, the id is indeed added to the document.</p>\n<p>Here’s an overview
  of PUT, PATCH, MERGE PATCH, and JSON PATCH requests.</p>\n<table>\n<thead>\n<tr>\n<th></th>\n<th>PUT</th>\n<th>PATCH
  (General)</th>\n<th>Merge Patch (JSON)</th>\n<th>JSON Patch</th>\n</tr>\n</thead>\n<tbody>\n<tr>\n<td><strong>Purpose</strong></td>\n<td>Replace
  the entire resource.</td>\n<td>Partially modify a resource.</td>\n<td>Partially
  modify a JSON resource by merging changes.</td>\n<td>Partially modify a JSON resource
  using explicit operations.</td>\n</tr>\n<tr>\n<td><strong>Request Body</strong></td>\n<td>Full
  representation of the resource.</td>\n<td>Instructions for modification.</td>\n<td>JSON
  object containing fields to be updated.</td>\n<td>JSON array of operation objects
  (e.g., <code>add</code>, <code>remove</code>, <code>replace</code>).</td>\n</tr>\n<tr>\n<td><strong>Idempotency</strong></td>\n<td>Idempotent
  (repeated requests have the same effect).</td>\n<td>Not necessarily idempotent (depends
  on implementation).</td>\n<td>Idempotent (repeated requests have the same effect).</td>\n<td>Idempotent
  (if operations are idempotent, like <code>replace</code>).</td>\n</tr>\n<tr>\n<td>Safe</td>\n<td>No</td>\n<td>No</td>\n<td>No</td>\n<td>No</td>\n</tr>\n<tr>\n<td><strong>Data
  Sent</strong></td>\n<td>All resource fields.</td>\n<td>Only the fields/instructions
  to be modified.</td>\n<td>Only the fields to be updated.</td>\n<td>Instructions
  specifying how to modify the resource (path, operation, value).</td>\n</tr>\n<tr>\n<td><strong>Server
  Behavior</strong></td>\n<td>Replaces the entire resource.</td>\n<td>Applies the
  provided modifications.</td>\n<td>Merges the provided fields into the existing resource.</td>\n<td>Applies
  the operations in the order they appear in the array.</td>\n</tr>\n<tr>\n<td><strong>Content-Type</strong></td>\n<td><code>application/json</code>
  (usually)</td>\n<td>Varies depending on the patch format.</td>\n<td><code>application/merge-patch+json</code></td>\n<td><code>application/json-patch+json</code></td>\n</tr>\n</tbody>\n</table>\n<p>So,
  that is all from the PATCH method in GOlang, I might have explained a bit too much
  for this, but there is a lot more to learn, which we will explore while creating
  a CRUD API in Golang later in this series.</p>\n<h2>Conclusion</h2>\n<p>I have also
  included some more examples of PATCH requests <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/blob/main/web/methods/patch/\"><strong>here</strong></a>.</p>\n<p>That's
  it from the 36th part of the series, all the source code for the examples are linked
  in the GitHub on the <a href=\"https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/methods/patch/\"><strong>100
  days of Golang</strong></a> repository.</p>\n<p><a href=\"https://github.com/Mr-Destructive/100-days-of-golang\"><strong>100-days-of-golang</strong></a></p>\n<h3>References:</h3>\n<ul>\n<li>\n<p><a
  href=\"https://en.wikipedia.org/wiki/HTTP#Request_methods\">HTTP Request Methods</a></p>\n</li>\n<li>\n<p><a
  href=\"https://rubyonrails.org/2012/2/26/edge-rails-patch-is-the-new-primary-http-method-for-updates\">Ruby
  on Rails Patch method</a></p>\n</li>\n<li>\n<p><a href=\"https://datatracker.ietf.org/doc/html/rfc5789\">HTTP
  PATCH RFC</a></p>\n</li>\n<li>\n<p><a href=\"https://datatracker.ietf.org/doc/html/rfc6902\">HTTP
  JSON PATCH RFC</a></p>\n</li>\n<li>\n<p><a href=\"https://datatracker.ietf.org/doc/html/rfc7386\">HTTP
  JSON MERGE PATCH</a></p>\n</li>\n</ul>\n"
image_url: https://meetgor-cdn.pages.dev/100-days-of-golang/golang-036-patch-method.png
long_description: 'In previous sections of this series, we In this section, we The
  PATCH method is often compared to the PUT method, but with one important distinction:
  PATCH is used to perform partial updates on a resource. Unlike PUT, which typically
  requires you to '
now: 2025-05-01 18:11:33.312355
path: blog/posts/2025-01-11-GO-36-PATCH-Method.md
prevnext: null
series:
- 100-days-of-golang
slug: golang-web-patch-method
status: published
tags:
- go
templateKey: blog-post
title: 'Golang Web: PATCH Method'
today: 2025-05-01
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