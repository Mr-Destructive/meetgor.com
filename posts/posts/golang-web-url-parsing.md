{
  "title": "Golang Web: URL Parsing",
  "status": "published",
  "slug": "golang-web-url-parsing",
  "date": "2025-04-05T12:33:30Z"
}

<h2>Introduction</h2>
<p>We have done around 32 posts on the fundamental concepts in golang, With that basic foundation, I'd like to start with the new section of this series which will be a major one as <code>web-development</code>. This section will have nearly 40-50 posts, this will cover the fundamental concepts for web development like APIs, Database integrations, Authentication and Authorizations, Web applications, static sites, etc.</p>
<h2>What is a URL?</h2>
<p>A URL is a Uniform Resource Locator. It is a string of characters that identifies a resource on the Internet. URLs are the building blocks of the web, allowing us to access websites, documents, and data with just a click. URLs are all over the place, if we want to build a strong foundation in web development, it's quite important to understand what URLs actually mean and what can they store.</p>
<p>A URL looks something like this:</p>
<pre><code>[scheme:][//[userinfo@]host][/]path[?query][#fragment]
</code></pre>
<p>Not all the URLs are like this, majority of the URLs that the common user sees are simply the ones with the <code>scheme</code>, <code>host</code>, and <code>paths</code>. However other components are equally important and are vital in the exchanging of information over the network.</p>
<ul>
<li>The <code>scheme</code> is the protocol used for accessing the resource like <code>http</code>, <code>https</code>, <code>ftp</code>, etc.</li>
<li>The <code>userinfo</code> is the username and password used to access the resource.</li>
<li>The <code>host</code> is the domain name of the resource.</li>
<li>The <code>path</code> is the path or folder to the resource.</li>
<li>The <code>query</code> is the query string of the resource. It is usually a key-value pair as a paramter to access resources.</li>
<li>The <code>fragment</code> is used as a reference within the resource.</li>
</ul>
<p>We will see the use cases of most of them throughout this series for example, the <code>userinfo</code> is commonly used in accessing databases over the internet/cloud. The query parameters will be used in making dynamic API calls, etc.</p>
<h2>Basic URL Parsing</h2>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/url&quot;
)

func main() {
	// simple url
	urlString := &quot;http://www.google.com&quot;
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedUrl)
}
</code></pre>
<pre><code>$ go run main.go

http://www.google.com
</code></pre>
<p>So, what is getting parsed here, we gave the URL as a string we get the URL back, the only difference is that instead of the URL being a string, it is now a structure of components. For instance, we want the protocol the host name, the port, etc. from the URL.</p>
<pre><code>fmt.Printf(&quot;%T
&quot;, parsedUrl)
// *url.URL
</code></pre>
<p>The <code>parsedUrl</code> is a pointer to a <a href="https://pkg.go.dev/net/url#URL">url.URL</a> structure. The structure <code>url.URL</code> has a lot of components to it like <code>Scheme</code>, <code>Host</code>, <code>User</code>, <code>Path</code>, <code>RawQuery</code>, etc. We will dive into each of these soon.</p>
<p>We could get those specific components ourselves, but that would be a bit tedious and might be even prone to bugs.</p>
<p>Let's try to get those components from the URL without any functions, just string manipulation.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
)

func main() {
    urlString := &quot;http://www.google.com&quot;
    protocol := urlString.split(&quot;:&quot;)[0]
    hostName := urlString.split(&quot;//&quot;)[1]
    fmt.Println(protocol)
    fmt.Println(hostName)
}
</code></pre>
<p>This might work for a simple URL, but what if we have more complex URLs which have paths, query parameters, fragments, username, port, etc? This could mess up quickly if we tried to get the parts of the URL ourselves. So, golang has a package called <a href="https://pkg.go.dev/net/url">net/url</a> explicitly for parsing URLs.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;strings&quot;
)

func main() {
	urlString := []string{&quot;http://www.google.com&quot;,
		&quot;http://www.google.com/about/&quot;,
		&quot;http://www.google.com/about?q=hello&quot;,
		&quot;http://www.google.com:8080/about?q=hello&quot;,
		&quot;http://user:password@example.com:8080/path/to/resource?query=value#fragment&quot;,
	}
	for _, url := range urlString {
		hostStr := strings.Split(url, &quot;//&quot;)[1]
        hostName := strings.Split(hostStr, &quot;/&quot;)[0]
		fmt.Println(hostName)
	}
}
</code></pre>
<pre><code>$ go run main.go

www.google.com
www.google.com
www.google.com
www.google.com:8080
user:password@example.com:8080
</code></pre>
<p>The above code might work for most of the URLs, but what if we have a more complex URL like the one with <code>port</code> or <code>user</code>, it doesn't give the thing we want exactly. In the above example, we have created a list of URLs as strings and simply iterated over each of the <code>url</code> in the <code>urlString</code>. Thereafter, we split the <code>url</code> on <code>//</code> so we get <code>https:</code> and <code>www.google.com</code>, if we want the host/domain name, we could simply get the <code>1</code> index in the slice since the <a href="https://pkg.go.dev/strings#Split">strings.Split</a> method returns a slice after splitting the string with the provided separator. The <code>hostName</code> could be fetched from the <code>1</code> index. However, this time for the 2nd element in the list, we have <code>https://www.google.com/about</code>, which would return <code>www.google.com/about</code> as the hostname which is not ideal, so we will again have to split this string with <code>/</code> and grab the first part i.e. 0th index.</p>
<p>The above code would work for <code>paths</code> and <code>query</code> parameters but if we had ports, and username, and password, it would not work as expected as evident from the last 2 examples in the list.</p>
<p>So, now we know the downsides of manually parsing the URLs, we can use the <a href="https://pkg.go.dev/net/url">net/url</a> package to do it for us.</p>
<h2>Parsing DB URLs</h2>
<p>Databases have a connection URL or connection string which provides a standard way to connect to a database/database server. The format of the URL is just the <code>URL</code> with all the components from the <code>scheme</code> to the <code>path</code>. The common examples of some database connection URLs might include:</p>
<pre><code># PostgreSQL DB Connection URL/string
postgresql://username:password@hostname:port/database_name

# MongoDB Connection URL/string
mongodb://username:password@hostname:port/database_name
</code></pre>
<p>The above are examples of the Postgres and MongoDB connection URLs, they have a <code>protocol</code> which usually for databases is their short name, the user credentials i.e. <code>username</code> and <code>password</code>, the <code>hostname</code> i.e. the server IP address, the <code>port</code> on which the database is running, and finally the path as the <code>database name</code>.</p>
<p>We can construct a simple snippet in golang to grab all the details from the simple connection URL string with the <code>net/url</code> package.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/url&quot;
)

func main() {
	//postgres db url
	dbUrl, err := url.Parse(&quot;postgres://admin:pass1234@localhost:5432/mydb&quot;)
	if err != nil {
		panic(err)
	}
	fmt.Println(dbUrl)
	fmt.Println(&quot;Scheme/Protocol = &quot;, dbUrl.Scheme)
	fmt.Println(&quot;User = &quot;, dbUrl.User)
	//fmt.Println(&quot;User = &quot;, dbUrl.User.String())
	fmt.Println(&quot;Username = &quot;, dbUrl.User.Username())
	password, _ := dbUrl.User.Password()
	fmt.Println(&quot;Password = &quot;, password)
	fmt.Println(&quot;Host = &quot;, dbUrl.Host)
	fmt.Println(&quot;HostName = &quot;, dbUrl.Hostname())
	fmt.Println(&quot;Port = &quot;, dbUrl.Port())
	fmt.Println(&quot;DB Name = &quot;, dbUrl.Path)
}
</code></pre>
<pre><code>$ go run main.go

postgres://postgres:pass1234@localhost:5432/mydb
Scheme/Protocol =  postgres
User = admin:pass1234
Username =  admin
Password =  pass1234
Host =  localhost:5432
HostName =  localhost
Port =  5432
DB Name =  mydb
</code></pre>
<p>In the above code, we have given the string <code>postgres://admin:pass1234@localhost:5432/mydb</code>, and we have parsed the URL using the <code>net/url</code> package. The result is we have a <code>parsedUrl</code> object which has all the components that can be accessed as either fields or methods. Let's break down each field/method we used in the above example:</p>
<ul>
<li>The <code>Scheme</code> is simply a string representing the protocol of the resource(URL).</li>
<li>The <code>User</code> is the <a href="https://pkg.go.dev/net/url#Userinfo">UserInfo</a> object having immutable username and password fields.</li>
<li>The <code>Username</code> is the method on <a href="https://pkg.go.dev/net/url#Userinfo.Username">UserInfo</a> that returns the string representing the username of the URL.</li>
<li>The <code>Password</code> is the method on <a href="https://pkg.go.dev/net/url#Userinfo.Password">UserInfo</a> that returns the string representing the password of the URL.</li>
<li>The <code>Host</code> is the field on <code>URL</code> as a string representing the host:port of the URL.</li>
<li>The <code>Hostname</code> is the method on <a href="https://pkg.go.dev/net/url#URL.Hostname">URL</a> that returns the string representing the hostname of the URL.</li>
<li>The <code>Port</code> is the method on <a href="https://pkg.go.dev/net/url#URL.Port">URL</a> that returns the string representing the port of the URL.</li>
<li>The <code>Path</code> is the field as the string representing the path of the URL.</li>
</ul>
<p>So, this is how we can get almost every detail like <code>db protocol</code>, <code>username</code>, <code>password</code>, <code>hostname</code>, <code>port</code>, and the <code>database name</code> from the database connection string URL.</p>
<h2>Parsing Query Parameters</h2>
<p>We can even get the query parameters of a URL using the <a href="https://pkg.go.dev/net/url#URL.Query">Query</a> method on the <code>URL</code> object. The <code>Query</code> method returns a <code>map[string][]string</code> which is to say a map with the key as <code>string</code> and the value as a <code>[]string</code> slice of string. For example, if the URL is something like <code>https://google.com/?q=hello</code>, the <code>Query</code> method will return <code>map[q:[hello]]</code> which means the key is <code>q</code> and the value is a list of strings of which the only element is <code>hello</code>.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/url&quot;
)

func main() {
	// a complex url with query params
	urlStr := &quot;http://www.google.com/?q=hello&amp;q=world&amp;lang=en&quot;
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedUrl)
	fmt.Println(parsedUrl.Query())
	for k, v := range parsedUrl.Query() {
		fmt.Println(&quot;KEY:&quot;, k)
		for _, vv := range v {
			fmt.Println(vv)
		}
	}
}
</code></pre>
<pre><code>$ go run main.go

http://www.google.com/?q=hello+world&amp;lang=en&amp;q=gopher
map[lang:[en] q:[hello world gopher]]
KEY: q
hello world
gopher
KEY: lang
en
</code></pre>
<p>We have taken a bit of a complex example that might cover many use cases of the <code>Query</code> method. We have a URL as <code>http://www.google.com/?q=hello&amp;q=world&amp;lang=en</code>, and the <code>Query</code> method returns <code>map[lang:[en] q:[hello world]]</code> which means the <code>q</code> key has a slice of a string with elements <code>hello world</code> and <code>gopher</code> and the <code>lang</code> key has a value of <code>en</code>. Here, the first paramter, <code>q=hello+world</code> is basically <code>hello world</code> or <code>hello%20world</code>, which is to say escaping the space in the URL. We can have multiple values for the same key, which is evident as we have added <code>q=gopher</code> at the end of the <code>URL</code>, the key <code>q</code> has two elements in the slice as <code>hello world</code> and <code>gopher</code>. The <code>lang=en</code> is simply a key as <code>lang</code> with the only element as <code>en</code> in the slice. We use <code>&amp;</code> to separate different query parameters in the URL.</p>
<h3>Checking Values in Query Parameters</h3>
<p>We can even check the values in the query parameters without requiring the construction of for loops to find a particular value in a key. The <a href="https://pkg.go.dev/net/url#Values">Values</a> is a type that stores the map as a return value from the <code>Query</code> method. It has a few handy methods like:</p>
<ul>
<li><a href="https://pkg.go.dev/net/url#Values.Has">Has</a> to check if the key exists in the map (paramter as key <code>string</code> and returns a <code>bool</code>).</li>
<li><a href="https://pkg.go.dev/net/url#Values.Get">Get</a> to fetch the first value of the given key as a string or if not present then returns an empty string (paramter as key <code>string</code> and returns <code>string</code>).</li>
<li><a href="https://pkg.go.dev/net/url#Values.Add">Add</a> method is used to append the value for a given key (paramter as key <code>string</code> and value to be added as <code>string</code>).</li>
<li><a href="https://pkg.go.dev/net/url#Values.Set">Set</a> method is used to replace the value for a given key if already exists (paramter as key <code>string</code> and value as <code>string</code>).</li>
<li><a href="https://pkg.go.dev/net/url#Values.Del">Del</a> method is used to delete the value for a given key (paramter as key <code>string</code>).</li>
</ul>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/url&quot;
)

func main() {
	// a complex url with query params
	urlStr := &quot;http://www.google.com/?q=hello+world&amp;lang=en&amp;q=gopher&quot;
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedUrl)
	fmt.Println(parsedUrl.Query())

	queryParams := parsedUrl.Query()

	fmt.Println(queryParams.Get(&quot;q&quot;))

	fmt.Println(queryParams.Has(&quot;q&quot;))

	if queryParams.Has(&quot;lang&quot;) {
		fmt.Println(queryParams.Get(&quot;lang&quot;))
	}

	queryParams.Add(&quot;q&quot;, &quot;ferris&quot;)
	fmt.Println(queryParams)

	queryParams.Set(&quot;q&quot;, &quot;books&quot;)
	fmt.Println(queryParams)

	queryParams.Del(&quot;q&quot;)
	fmt.Println(queryParams)
}
</code></pre>
<pre><code>$ go run main.go

http://www.google.com/?q=hello+world&amp;lang=en&amp;q=gopher
hello world
true
en
map[lang:[en] q:[hello world gopher ferris]]
map[lang:[en] q:[books]]
map[lang:[en]]
</code></pre>
<p>The above code example demonstrates almost all the methods available on the <code>Values</code> type. The <code>Get</code> method is used to fetch the first value for a given key, so we parse the key as a <code>string</code> to the method and it would return a <code>string</code>. We checked for the <code>q</code> as the key and it returned the first element in the <code>queryParams</code> for key <code>q</code> as <code>hello world</code> from the list <code>[hello world, gopher]</code>. The <code>Has</code> method takes in the paramter as key as <code>string</code> and returns if the key exists in the <code>queryParams</code> or not as a bool. The <code>Add</code> method, we have used to <code>Add</code> a key with a particular value, we added the value <code>ferris</code> to the key <code>q</code> hence it appended to the list and the list of <code>queryParams[q]</code> became <code>[hello world, gopher, ferris]</code>. The <code>Set</code> method is used to override the existing key with a particular value, here we have set the value <code>books</code> to the key <code>q</code> and hence the list of <code>queryParams[q]</code> becomes <code>[books]</code>. We can use the <code>Del</code> method to remove the key from the <code>queryParams</code>, so we delete <code>q</code> from <code>queryParams</code>, then the <code>queryParams</code> simply has no key as <code>q</code> in it.</p>
<h3>Parsing Query Parameters to String</h3>
<p>Now, that you have manipulated the query parameters, let's say you want to construct back that string representation of the query particular or the URL for it. The <a href="">Encode</a> method is used to grab the <code>queryParams</code> i.e. <code>Values</code> object and convert it into the <code>string</code> representation of the encoded URL.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/url&quot;
)

func main() {
	// a complex url with query params
	urlStr := &quot;http://www.google.com/?q=hello+world&amp;lang=en&amp;q=gopher&quot;
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
    queryParams := parsedUrl.Query()
    queryParams.Add(&quot;name&quot;, &quot;ferris&quot;)

    q := queryParams.Encode()
    fmt.Println(q)
}
</code></pre>
<pre><code>$ go run main.go

lang=en&amp;name=ferris&amp;q=hello+world&amp;q=gopher
</code></pre>
<p>So, we can see the <code>Encode</code> method has given a query parameter in the form of a URL encoded string. We first grab the query parameters from the <code>parsedUrl</code> which is a <code>URL</code> object via the <code>Query</code> method, we then <code>Add</code> the key <code>name</code> with a value of <code>ferris</code> to the <code>queryParams</code>. This is then used to <code>Encode</code> the object back to a string representation. This could be useful to construct a query paramter for requesting other websites/APIs.</p>
<h2>Parsing URL object back to String</h2>
<p>We can even get the <code>URL</code> object back to a string representation using the <a href="https://pkg.go.dev/net/url#URL.String">String</a> method on the <code>URL</code> object. The <code>String</code> method returns a <code>string</code> representation of the URL object.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/url&quot;
)

func main() {
	urlStr := &quot;http://www.google.com/?q=hello+world&amp;lang=en&amp;q=gopher&quot;
	fmt.Println(&quot;URL:&quot;, urlStr)
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	queryParams := parsedUrl.Query()
	queryParams.Add(&quot;name&quot;, &quot;ferris&quot;)

	q := queryParams.Encode()
	fmt.Println(q)
	parsedUrl.RawQuery = q
	newUrl := parsedUrl.String()
	fmt.Println(&quot;New URL:&quot;, newUrl)
}
</code></pre>
<pre><code>$ go run main.go

URL: http://www.google.com/?q=hello+world&amp;lang=en&amp;q=gopher
lang=en&amp;name=ferris&amp;q=hello+world&amp;q=gopher
New URL: http://www.google.com/?lang=en&amp;name=ferris&amp;q=hello+world&amp;q=gopher
</code></pre>
<p>In the example above, we parse a URL string into a <code>URL</code> object as <code>parsedUrl</code>, then we <code>Add</code> the key <code>name</code> with a value of <code>ferris</code> to the <code>queryParams</code>. We then <code>Encode</code> the <code>URL</code> object back to a string representation. But this won't change the <code>parsedUrl</code> object we want to change the entire <code>URL</code> object. For that, we have overwritten the <code>RawQuery</code> field of the <code>URL</code> object with the query parameter encoded string as <code>q</code>. The <code>String</code> method returns a <code>string</code> representation of the <code>URL</code> object.</p>
<h2>Parsing Fragments</h2>
<p>The fragment in a URL is usually present in a static website like <code>#about</code>, <code>#contact</code>, <code>#blog</code>, etc. The <code>Fragment</code> is a string that is usually a reference to a specific section or anchor point within a web page or resource. When a URL with a fragment is accessed, the web browser or user agent will scroll the page to display the section identified by the fragment.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/url&quot;
)

func main() {
	// url with fragments
	urlStr := &quot;https://pkg.go.dev/math/rand#Norm Float64&quot;
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsedUrl)
	fmt.Println(parsedUrl.Fragment)
	fmt.Println(parsedUrl.RawFragment)
	fmt.Println(parsedUrl.EscapedFragment())
}
</code></pre>
<pre><code>$ go run main.go

https://pkg.go.dev/math/rand#Norm Float64

Norm Float64
Norm Float64
Norm%20Float64
</code></pre>
<p>The above code is used to fetch the <code>#Norm Float64</code> fragment from the URL <code>https://pkg.go.dev/math/rand#NormFloat64</code>. We can use the <a href="https://pkg.go.dev/net/url#URL">Fragment</a> field in the <code>URL</code> object to get the fragment text. There is <a href="https://pkg.go.dev/net/url#URL">RawFragment</a> field that is used to parse the fragment text as it is, without trying to escape any special characters in the URL. The <a href="https://pkg.go.dev/net/url#URL.EscapedFragment">EscapedFragment</a> is used to parse the fragment text by escaping the characters in the URL.</p>
<p>That's it from the 32nd part of the series, all the source code for the examples are linked in the GitHub on the <a href="https://github.com/Mr-Destructive/100-days-of-golang/tree/main/web/url-parsing">100 days of Golang</a> repository.</p>
<p><a href="https://github.com/Mr-Destructive/100-days-of-golang">100-days-of-golang</a></p>
<h2>Conclusion</h2>
<p>From the first post of the web development section, we covered the fundamentals of URL parsing and got a bit introduced to the <code>net</code> package, which will be heavily used for most of the core language's features for working with the web. We covered the concepts for parsing URLs, getting components of URLs from the parsed object, Database connection URL resolving, parsing query parameters, and some other URL-related concepts.</p>
<p>Hopefully, you found this section helpful, if you have any comments or feedback, please let me know in the comments section or on my social handles. Thank you for reading. Happy Coding :)</p>
