{
  "title": "Golang: Error Handling",
  "status": "published",
  "slug": "golang-error-handling",
  "date": "2025-04-05T12:33:36Z"
}

<h2>Introduction</h2>
<p>Error handling is quite an important feature of any programming language to improve the quality and transparency between the user and the application. By raising appropriate error messages, the user can get a clear idea about the things happening in the interface as well as the application can handle the errors with appropriate actions.</p>
<p>In the 20th post of the series, we will be exploring the concept of error handling in golang. From this article, we will be able to learn the fundamentals of error or exception handling in golang, create custom error classes, raise and ignore error messages, and exit or redirect the flow of state of the application when an error is raised.</p>
<h2>Why we need Error Type</h2>
<p>We need <a href="https://go.dev/blog/error-handling-and-go">error handling</a> and catching in order to stop or divert the flow of the application which will restrict the dubious or unintentional execution of the code. Let's say, for example, we have a string as an input, and the user skipped the input and the string is returned as empty, we don't want to execute the further program as the execution might depend on the value of the string. So, in order to catch these kinds of events, we might use errors and log the results for better transparency.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
)

func main() {

	var s string
	n, err := fmt.Scanf(&quot;%s&quot;, &amp;s)

	if err != nil {
		fmt.Println(err)
		// panic(err)
        // OR
        // return
	} else {
		fmt.Println(n)
		if s[0] == 'a' {
			fmt.Println(s)
		}
	}
}
</code></pre>
<pre><code>$ go run main.go
asdf
1
asdf


$ go run main.go

unexpected newline


$ go run main.go
wsd
1
</code></pre>
<p>In the above example, we have a simple string <code>s</code> input, we will input the string using the <code>Scanf</code> function that will return an integer as the number of variables it has scanned and error if any. Here, as the function might return two values, we need two variables to call the function. The <code>n</code> variable stores the number of variables successfully scanned and the <code>err</code> as the err from the function generated. If there is an error, that is the value stored in <code>err</code> is not empty, we will log the error. And move into the rest of the program.</p>
<p>This might be looking cool, but it doesn't break out of the program if there is an error. We want it to log the error as well as exit from the program. We can do that using the panic function.</p>
<h2>Catching Errors</h2>
<p>We'll see a few examples, where we will catch errors in some of the regularly used functions. These error messages can be used in deciding the next procedure to be run.</p>
<h3>Comma OK/Error syntax</h3>
<p>We use the <a href="https://go.dev/doc/effective_go#:~:text=is%20called%20the%20%E2%80%9C-,comma%20ok,-%E2%80%9D%20idiom.%20In%20this">comma ok, error</a> syntax where we want multiple return values from a function. It is a narrowed syntax for a two-type return value function, we either return a value that we are expecting as <code>OK</code>, or we get an <code>error</code> from the function call.</p>
<pre><code class="language-go">
ok, err := function()
if err != nil {
    // handle error
    panic(err)
} else {
    // work with the ok object
    fmt.Println(ok)
}

</code></pre>
<p>In the above code, we have used the comma-ok, error syntax, the function call will return two objects either an expected object or an error object if there were errors in the processing. We handle the error if the error object is not empty i.e. it contains something, else we can do the rest of the required processing of the program.</p>
<p>We can even ignore the <code>err</code> or the <code>ok</code> object using the <code>_</code> i.e. to say a don't care variable. Remember you can ignore either of the values and not both. It is not recommended to ignore errors but if you know the obvious thing to process, you might as well sometimes.</p>
<p>A more compressed code might look something like below:</p>
<pre><code class="language-go">
if ok, err := function(); err != nil {
    // handle error
    panic(err)
} else {
    // work with the ok object
    fmt.Println(ok)
}
</code></pre>
<p>The above code wraps the initialization of <code>ok, err</code> objects by calling the function inside the if statement and checking for the error.</p>
<h3>Making HTTP requests</h3>
<p>Let's say we have a URL, and we want to check if the site exists or not, we can run the <a href="https://pkg.go.dev/net/http@go1.19.1#Client.Get">http.Get</a> function from the <a href="https://pkg.go.dev/net/http">net/http</a> package. We will parse a URL in the function and this function also returns a <a href="https://pkg.go.dev/net/http#Response">Response</a> type object and an error object if there are any errors generated during the function call. If there is an error while we call the function, we simply log the error and panic out of the program. Else we can log the status code and do further processing.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/http&quot;
)

func main() {
	url := &quot;https://meetgor.com/&quot;
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println(resp.StatusCode)
	}
}
</code></pre>
<pre><code>$ go run web.go
URL: https://meetgor.com/
200



$ go run web.go

URL: htts://meetgor.com/
Get &quot;htts://meetgor.com/&quot;: unsupported protocol scheme &quot;htts&quot;
panic: Get &quot;htts://meetgor.com/&quot;: unsupported protocol scheme &quot;htts&quot;

goroutine 1 [running]:
main.main()
        /home/meet/code/100-days-of-golang/scripts/errors/https.go:14 +0x170
exit status 2
</code></pre>
<p>This is how we can validate a URL handling the error if the parsed URL is invalid or does not exist.</p>
<p>There is one more variation of the above code style, it is a bit compressed and might be just a syntactic change.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/http&quot;
)

func main() {
	url := &quot;https://meetgor.com/&quot;
	if resp, err := http.Get(url); err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println(resp.StatusCode)
	}
}
</code></pre>
<p>This can be used wherever you are using the <code>ok, err</code> kind of syntax, but I prefer the clean syntax so I won't move ahead with this.</p>
<h3>Opening or Handling of File</h3>
<p>We can even use error handling while dealing with Files or Folders. We can use the <a href="https://pkg.go.dev/os">os</a> package to read the file in golang. The <a href="https://pkg.go.dev/os#Open">Open</a> function will read the file if it exists or else it will return an error. We can catch the error from the comma <code>ok,error</code> syntax and do the required processing in the program.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;os&quot;
)

func main() {
	file_name := &quot;hello.txt&quot;
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Println(&quot;Error: &quot;, err)

        // Create a File 

		// _, err := os.Create(file_name)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(&quot;Created File&quot;, file_name)
		// file, _ = os.Open(file_name)
	}
	fmt.Println(file.Stat())
}
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1663425619/blog-media/golang-file-error.png" alt="File handling Error"></p>
<p>We can use the error as a hint that the file doesn't exist and create a file and then move toward the actual execution of the program. We can also ignore the file object while creating the file, as we are interested in only knowing that the file is just created without any errors, we use the <code>_</code> to ignore the variable in the assignment in the function call.</p>
<p>Inside the commented code, we use the <a href="https://pkg.go.dev/os#Create">Create</a> function to create a file and check for any errors in the process. We finally create the file and <code>Open</code> the newly created file.</p>
<h2>Custom Error</h2>
<p>We can create custom error types in golang with the help of interfaces and structs. An Error struct will simply consist of a string <code>message</code>, that string will display the error. By overriding or creating the <code>Error</code> method as an interface for the struct we can construct custom errors.</p>
<pre><code class="language-go">type Invalid_URL_Error struct {
	message string
}

func (e Invalid_URL_Error) Error() string {
	return &quot;Invalid URL&quot;
}

</code></pre>
<p>Here, we have the <code>Invalid_URL_Error</code> as the custom struct name and the <code>Error()</code> method as an interface that will print the error log. This Error method will be used while raising errors in the program. It might be called from another function while doing the actual processing of the URL while sending a GET request.</p>
<p>Further, we can call this custom error method when we wish, by using the package functions, we can override the function call with the custom method.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
    &quot;net/http&quot;
)

type Invalid_URL_Error struct {
	message string
}

func (e Invalid_URL_Error) Error() string {
	return &quot;Invalid URL&quot;
}

func main() {
	url := &quot;htt://meetgor.com&quot;
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(Invalid_URL_Error{})
		fmt.Println(err)
	} else {
		fmt.Println(response)
		defer response.Body.Close()
	}
}
</code></pre>
<pre><code>$ go run custom_error.go
Invalid URL 
Get &quot;htt://meetgor.com&quot;: unsupported protocol scheme &quot;htt&quot;

</code></pre>
<p>In the above code, we are basically calling the function <code>http.Get</code> that will return a <code>Response</code> or an <code>err</code> object. We can even call the custom error method with an empty <code>Invalid_URL_Error</code> object, this will call the function <code>Error</code> from that interface. The function will print the custom error message and thereby we are able to log the custom error message for the invalid URL example.</p>
<p>Also, we can parse the default error method to the custom error method which will get us additional information inside the error interface method.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/http&quot;
)

type Invalid_URL_Error struct {
	message string
}

func (e Invalid_URL_Error) Error() string {
	return &quot;Invalid URL : &quot; + e.message
}

func main() {
	url := &quot;htt://meetgor.com&quot;
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(Invalid_URL_Error{err.Error()})
	} else {
		fmt.Println(response)
		defer response.Body.Close()
	}
}
</code></pre>
<pre><code>$ go run custom_error.go
Invalid URL : Get &quot;htt://meetgor.com&quot;: unsupported protocol scheme &quot;htt&quot;

</code></pre>
<p>When the URL is invalid, we will call the custom error interface by parsing the default <code>err.Error</code> method. This will get us the error object from the main method to our custom interface. That is how we will be able to fetch additional information about the error from the interface with the <code>.</code> operator as <code>e.message</code>. So, the syntax is <code>Invalid_URL_Error{err.Error()}</code>, i.e. an object of type <code>Invalid_URL_Error</code> with the message set as the value returned from the default <code>Error()</code> function. hence we can implement the custom error message.</p>
<p>We also need to look for the response object and close the Response Body as it is mandatory to do so and the responsibility of the caller.</p>
<h3>Creating a function that returns two values (ok,error)</h3>
<p>We can even nest the calling of this error method inside another function. This will give us a good overview of how to deal with errors more thoroughly. We will construct a function with two return values one can be any normal desirable object (which we want from the function call) and the other as an error. This will check for any cases that we can call the custom error and return that error interface and the object which was to be returned will be nil if there is an error. If there are no errors, we will return the object and set the error as nil. This way, we can use the <code>ok, error</code> syntax while calling this function and thereby make it a lot easier in case of complex programs or multiple types of errors.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/http&quot;
)

type Invalid_URL_Error struct {
	message string
}

func (e Invalid_URL_Error) Error() string {
	return &quot;Invalid URL&quot;
}

func get_resp(url_link string) (http.Response, error) {

	resp, err := http.Get(url_link)

	if err != nil {
		return http.Response{}, &amp;Invalid_URL_Error{}
	} else {
		defer resp.Body.Close()
		return *resp, nil
	}

}

func main() {

	url := &quot;htts://meetgor.com&quot;
	resp, err := get_resp(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}

}
</code></pre>
<pre><code>$ go run errors.go
Invalid URL
</code></pre>
<p>This is the simple function <code>get_resp</code> which will either fetch the <code>Response</code> or <code>error</code>. We call the <code>http.Get</code> method internally in the function and if something is invalid, it will return an empty Response but the error will be an <code>Invalid_URL_Error</code> object which is the custom error class. This means, that if we have an error, we will return a string object from the <code>Error</code> method in the interface and if there is no error, we will return the <code>Response</code> object and the error will be set as <code>nil</code>. Hence, you can now realize, why we check for <code>err != nil</code>, it is used for checking if an error object has returned a string or not. As said earlier, we also need to close the request Body.</p>
<p>Further, we can pass the default error method to the custom error class as <code>Invalid_URL_Error{err.Error()}</code>. This will ensure, we get additional context from the custom error interface.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;net/http&quot;
)

type Invalid_URL_Error struct {
	message string
}

func (e Invalid_URL_Error) Error() string {
	return &quot;Invalid URL : &quot; + e.message
}

func get_resp(url_link string) (http.Response, error) {

	resp, err := http.Get(url_link)

	if err != nil {
		return http.Response{}, &amp;Invalid_URL_Error{err.Error()}
	} else {
		defer resp.Body.Close()
		return *resp, nil
	}

}

func main() {

	url := &quot;htts://meetgor.com&quot;
	resp, err := get_resp(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}

}
</code></pre>
<pre><code>$ go run custom_error
Invalid URL : Get &quot;htts://meetgor.com&quot;: unsupported protocol scheme &quot;htts&quot;
</code></pre>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/error-handling/main.go">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>From this article, we were able to understand the basics of error handling in golang. We can now work with handling errors in function calls and create custom error interfaces. Thank you for reading, if you have any queries, feedback, or questions, you can freely ask me on my social handles. Happy Coding :)</p>
