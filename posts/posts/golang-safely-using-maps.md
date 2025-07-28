{
  "title": "Safely using Maps in Golang: Differences in declaration and initialization",
  "status": "published",
  "slug": "golang-safely-using-maps",
  "date": "2025-04-05T12:33:25Z"
}

<h2>Introduction</h2>
<p>This week, I was working on one of the API wrapper packages for golang, and that dealt with sending post requests with URL encoded values, setting cookies, and all the fun stuff. However, while I was constructing the body, I was using <a href="https://pkg.go.dev/net/url#Values">url.Value</a> type to construct the body, and use that to add and set key-value pairs. However, I was getting a wired <code>nil</code> pointer reference error in some of the parts, I thought it was because of some of the variables I set manually. However, by debugging closer, I found out a common pitfall or bad practice of just declaring a type but initializing it and that causing <code>nil</code> reference errors.</p>
<p>In this post, I will cover, what are maps, how to create maps, and especially how to properly declare and initialize them. Create a proper distinction between the declaration and initialization of maps or any similar data type in golang.</p>
<h2>What is a map in Golang?</h2>
<p>A <a href="https://go.dev/src/runtime/map.go">map</a> or a hashmap in golang is a basic data type that allows us to store key-value pairs. Under the hood, it is a header map-like data structure that holds buckets, which are basically pointers to bucket arrays (contiguous memory). It has hash codes that store the actual key-value pairs, and pointers to new buckets if the current overflows with the number of keys. This is a really smart data structure that provides almost constant time access.</p>
<h2>How to create maps in Golang</h2>
<p>To create a simple map in golang, you can take an example of a letter frequency counter using a map of string and integer. The map will store the letters as keys and their frequency as values.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
    words := &quot;hello how are you&quot;
    letters := map[string]int{}

    for _, word := range words {
        wordCount[word]++
    }

    fmt.Println(&quot;Word counts:&quot;)
    for word, count := range wordCount {
        fmt.Printf(&quot;%s: %d
&quot;, word, count)
    }
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

Word counts:
e: 2
 : 3
w: 1
r: 1
y: 1
u: 1
h: 2
l: 2
o: 3
a: 1
</code></pre>
<p>So, by initializing the map as <code>map[string]int{}</code> you will get an empty map. This can be then used to populate the keys and values, we iterate over the string, and for each character (rune) we cast that byte of character into the string and increment the value, the zero value for int is <code>0</code>, so by default if the key is not present, it will be zero, it is a bit of double-edged swords though, we never know the key is present in the map with the value <code>0</code> or the key is not present but the default value is <code>0</code>. For that, you need to check if the <code>key</code> exists in the map or not.</p>
<p>For further details, you can check out my <a href="https://www.meetgor.com/golang-maps/">Golang Maps</a> post in detail.</p>
<h2>Difference between declaration and initialization</h2>
<p>There is a difference in declaring and initializing any variable in a programming language and has to do a lot more with the implementation of the underlying type. In the case of primary data types like <code>int</code>, <code>string</code>, <code>float</code>, etc. there are default/zero values, so that is the same as the declaration and initialization of the variables. However, in the case of maps and slices, the declaration is just making sure the variable is available to the scope of the program, however for initialization is setting it to its default/zero value or the actual value that should be assigned.</p>
<p>So, declaration simply makes the variable available within the scope of the program. For maps and slices, declaring a variable without initialization sets it to <code>nil</code>, meaning it points to no allocated memory and cannot be used directly.</p>
<p>Whereas the <code>initialization</code> allocates memory and sets the variable to a usable state. For maps and slices, you need to explicitly initialize them using syntax like <code>myMap = make(map[keyType]valueType)</code> or <code>slice = []type{}</code>. Without this initialization, attempting to use the map or slice will lead to runtime errors, such as panics for accessing or modifying a nil map or slice.</p>
<p>Let's looks at the values of a map when it is declared/initialized/not initialized.</p>
<p>Imagine you're building a configuration manager that reads settings from a map. The map will be declared globally but initialized only when the configuration is loaded.</p>
<ol>
<li>Declared but not initialized</li>
</ol>
<p>The below code demonstrates a map access that is not initialized.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;log&quot;
)

// Global map to store configuration settings
var configSettings map[string]string // Declared but not initialized

func main() {
	// Attempt to get a configuration setting before initializing the map
	serverPort := getConfigSetting(&quot;server_port&quot;)
	fmt.Printf(&quot;Server port: %s
&quot;, serverPort)
}

func getConfigSetting(key string) string {
	if configSettings == nil {
		log.Fatal(&quot;Configuration settings map is not initialized&quot;)
	}
	value, exists := configSettings[key]
	if !exists {
		return &quot;Setting not found&quot;
	}
	return value
}
</code></pre>
<pre><code class="language-bash">$ go run main.go
Server port: Setting not found
</code></pre>
<ol start="2">
<li>Declared and Initialized at the same time</li>
</ol>
<p>The below code demonstrates a map access that is initialized at the same time.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;log&quot;
)

// Global map to store configuration settings
var configSettings = map[string]string{
	&quot;server_port&quot;:  &quot;8080&quot;,
	&quot;database_url&quot;: &quot;localhost:5432&quot;,
}

func main() {
	serverPort := getConfigSetting(&quot;server_port&quot;)
	fmt.Printf(&quot;Server port: %s
&quot;, serverPort)
}

func getConfigSetting(key string) string {
	value, exists := configSettings[key]
	if !exists {
		return &quot;Setting not found&quot;
	}
	return value
}
</code></pre>
<pre><code class="language-bash">$ go run main.go
Server port: 8080
</code></pre>
<ol start="3">
<li>Declared and later initialized</li>
</ol>
<p>The below code demonstrates a map access that is initialized later.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;log&quot;
)

// Global map to store configuration settings
var configSettings map[string]string // declared but not initialized

func main() {
	// Initialize configuration settings
	initializeConfigSettings()
    // if the function is not called, the map will be nil

	// Get a configuration setting safely
	serverPort := getConfigSetting(&quot;server_port&quot;)
	fmt.Printf(&quot;Server port: %s
&quot;, serverPort)
}

func initializeConfigSettings() {
	if configSettings == nil {
		configSettings = make(map[string]string) // Properly initialize the map
		configSettings[&quot;server_port&quot;] = &quot;8080&quot;
		configSettings[&quot;database_url&quot;] = &quot;localhost:5432&quot;
		fmt.Println(&quot;Configuration settings initialized&quot;)
	}
}

func getConfigSetting(key string) string {
	if configSettings == nil {
		log.Fatal(&quot;Configuration settings map is not initialized&quot;)
	}
	value, exists := configSettings[key]
	if !exists {
		return &quot;Setting not found&quot;
	}
	return value
}
</code></pre>
<pre><code class="language-bash">$ go run main.go
Configuration settings initialized
Server port: 8080
</code></pre>
<p>In the above code, we declared the global map <code>configSettings</code> but didn't initialize it at that point, until we wanted to access the map. We initialize the map in the main function, this main function could be other specific parts of the code, and the global variable <code>configSettings</code> a map from another part of the code, by initializing it in the required scope, we prevent it from causing nil pointer access errors. We only initialize the map if it is <code>nil</code> i.e. it has not been initialized elsewhere in the code. This prevents overriding the map/flushing out the config set from other parts of the scope.</p>
<h2>Pitfalls in access of un-initialized maps</h2>
<p>But since it deals with pointers, it comes with its own pitfalls like nil pointers access when the map is not initialized.</p>
<p>Let's take a look at an example, a real case where this might happen.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/url&quot;
)

func main() {
        var vals url.Values
        vals.Add(&quot;foo&quot;, &quot;bar&quot;)
        fmt.Println(vals)
}
</code></pre>
<p>This will result in a runtime panic.</p>
<pre><code class="language-nginx">$ go run main.go
panic: assignment to entry in nil map

goroutine 1 [running]:
net/url.Values.Add(...)
        /usr/local/go/src/net/url/url.go:902
main.main()
        /home/meet/code/playground/go/main.go:10 +0x2d
exit status 2
</code></pre>
<p>This is because the <a href="https://pkg.go.dev/net/url#Values">url.Values</a> is a map of string and a list of string values. Since the underlying type is a map for <code>Values</code>, and in the example, we only have declared the variable <code>vals</code> with the type <code>url.Values</code>, it will point to a <code>nil</code> reference, hence the message on adding the value to the type. So, it is a good practice to use <code>make</code> while declaring or initializing a map data type. If you are not sure the underlying type is <code>map</code> then you could use <code>Type{}</code> to initialize an empty value of that type.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/url&quot;
)

func main() {
        vals := make(url.Values)
        // OR
        // vals := url.Values{}
        vals.Add(&quot;foo&quot;, &quot;bar&quot;)
        fmt.Println(vals)
}
</code></pre>
<pre><code class="language-bash">$ go run urlvals.go
map[foo:[bar]]
foo=bar
</code></pre>
<p>It is also recommended by the <a href="https://go.dev/blog/maps">golang team</a> to use the make function while initializing a map. So, either use <code>make</code> for maps, slices, and channels, or initialize the empty value variable with <code>Type{}</code>. Both of them work similarly, but the latter is more generally applicable to structs as well.</p>
<h2>Conclusion</h2>
<p>Understanding the difference between declaring and initializing maps in Golang is essential for any developer, not just in golang, but in general. As we've explored, simply declaring a map variable without initializing it can lead to runtime errors, such as panics when attempting to access or modify a nil map. Initializing a map ensures that it is properly allocated in memory and ready for use, thereby avoiding these pitfalls.</p>
<p>By following best practices—such as using the make function or initializing with Type{}—you can prevent common issues related to uninitialized maps. Always ensure that maps and slices are explicitly initialized before use to safeguard against unexpected nil pointer dereferences</p>
<p>Thank you for reading this post, If you have any questions, feedback, and suggestions, feel free to drop them in the comments.</p>
<p>Happy Coding :)</p>
