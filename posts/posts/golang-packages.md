{
  "title": "Golang: Packages",
  "status": "published",
  "slug": "golang-packages",
  "date": "2025-04-05T12:33:42Z"
}

<h2>Introduction</h2>
<p>In this 11th part of the series, we will be covering packages in golang. Package is a cool way to organize code in large projects. We can create a separate file which can include certain helper functions or variables from other files or scripts. There are couple of types of packages like packages from the standard library, open sourced community packages and custom packages that you can build of your own. In this particular section, we will be covering the fundamentals of packages and exploring the standard library in golang.</p>
<h2>What is a Package ?</h2>
<p>Package in simplest of terms is a collection of go scripts that can serve a purpose. Just like <code>fmt</code> has several functions like <code>Println</code>, <code>Printf</code>, <code>Scan</code>, <code>Scanf</code>, etc. most packages have functions that we can use in our own programs to solve a simple problem. We have already created many packages but none of them have been super useful so far, we just used the statement <code>package main</code> and didn't use the package anywhere. Every package has a entry point called <code>main</code> along with a entry function <code>main</code> which is triggered when we run the package.</p>
<pre><code>- package_name
    - script_1.go
    - script_2.go

    - sub_package_name
        - utility_1.go
    - go.mod
</code></pre>
<p>The above is a simple folder structure for a package in golang, we have the package itself as the name of the folder. Inside of the package folder, we would have the scripts or sub-packages if any. Also there is a <code>go.mod</code> file in all go source package folders which contain the meta information about the package and it's dependencies.</p>
<p>Let's take a look at the <a href="https://github.com/golang/go/tree/master/src/fmt">fmt</a> package source code, it has a <a href="https://github.com/golang/go/blob/master/src/fmt/print.go">print.go</a> file which has all the functions associated with printing, similarly separate files for different core functionality.</p>
<p>We will soon into the details of packages in this series. Right now, we only care about the fundamentals of a package in golang.</p>
<h2>Importing Packages</h2>
<p>We have been already importing a package since writing our hello world in go, the <code>fmt</code> package which holds some functions for formatting, printing, logging and various string input/output manipulation.</p>
<pre><code class="language-go">import &quot;fmt&quot;
</code></pre>
<p>So, we have used the simple <code>import &quot;package&quot;</code> statement, but there are a couple of more ways to import packages if we have multiple packages to import.</p>
<pre><code class="language-go">import (
    &quot;fmt&quot;
    &quot;math&quot;
)
</code></pre>
<p>Using the <code>()</code> and by specifying the name of the package we can import multiple packages at once. Also there is a harder way out there, if you really like to toil hard.</p>
<pre><code class="language-go">import &quot;fmt&quot;
import &quot;math&quot;
</code></pre>
<p>This is generally avoided as it just looks too pythonic, we are in golang.</p>
<h3>Aliasing Imported packages</h3>
<p>We can alias an package a name whatever we want for the usage in the rest of the script file. This allows a bit better semantics of longer package names into readable code.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	r &quot;math/rand&quot;
)

func main() {
	fmt.Println(r.Int())
}
</code></pre>
<pre><code>$ go run import.go
3454565657742387432
</code></pre>
<p>We have imported the package <code>math/rand</code>, here <a href="https://pkg.go.dev/math/rand@go1.18.1#Int">rand</a> is a sub package of the main package math. Hence we have aliased the rand package as <code>r</code> and thus, we can use r to access all the functions and other types from the package.</p>
<h3>Blank Package Import</h3>
<p>We can even import the package but not use it without getting a compilation error. So, the blank identifier is used in golang to ignore the initialized or returned values from any context and avoid the compilation warning or errors.</p>
<pre><code class="language-go">package main

import (
	_ &quot;fmt&quot;
)

func main() {
}
</code></pre>
<pre><code>$ go run blank_import.go
</code></pre>
<p>Here, we have imported <code>fmt</code> package with a <code>_</code> blank identifier but we didn't use it anywhere, still we don't get any error messages and it compiled the code successfully.</p>
<h2>Standard Library Packages</h2>
<p>The Golang Standard Library has some good number of packages which can be used for some general tasks like formatting input/output, file handling, web requests, system commands, etc. You can check out the entire list along with the documentation on the <a href="https://pkg.go.dev/std">official website</a>.</p>
<p>We can import these standard library packages just by parsing their name in the string quotes like we did with <code>fmt</code> as <code>&quot;fmt&quot;</code>. We have previously used the <code>rand</code> package from the math standard library package as a sub-package by using the statement <code>&quot;math/rand&quot;</code>, if we want we can import the entire <code>math</code> package as <code>&quot;math&quot;</code> but that's unwanted and we import only the package which we really need.</p>
<p>There are other packages as well like <code>bufio</code> which is used for reading and performing operations with text, <code>os</code> for working with files systems and operating system level stuff, and other packages which are specific to use cases like rendering templates, time, sorting, math operations, encoding, etc. We will dive into some of them throughout this series.</p>
<h2>Installing Packages</h2>
<p>We can now get into installing other packages which are not in the standard library. You can get the documentation along with all references for a particular package on the official Golang <a href="https://pkg.go.dev/">package repository</a>. We use the CLI command to grab the packages into our <code>GOPATH</code>. OK, GOPATH, we have not covered this!</p>
<h3>GOPATH</h3>
<p>GOPATH is the path or the location in your system's disk where all the packages and modules are stored. You can get the default location of your GOPATH environment variable from the simple shell command.</p>
<pre><code>$ echo $GOPATH
C:\Userscer\go
</code></pre>
<p>It has a few folders namely, <code>bin</code>, <code>pkg</code>, and <code>src</code>. These folder server different purpose like:</p>
<ul>
<li><code>bin</code> for storing the binaries generated from <code>go install</code> command</li>
<li><code>pkg</code> for storing pre-compiled source files and objects for quicker generation of executables and compilation.</li>
<li><code>src</code> for storing all the go source files of packages and modules.</li>
</ul>
<h3>Go Get command</h3>
<p>Now, let's see how to install a package from the go community on GitHub.</p>
<pre><code>go get github.com/gorilla/mux
</code></pre>
<p>We have installed a package which is a powerful HTTP router and a URL dispatcher and it can also be used to make web applications. It's called <code>mux</code>, we won't be using it right away just to get a feel for installing and playing with packages at the moment.</p>
<p>After executing the command <code>go get</code> you should see a folder to be added in the <code>$GOPATH\pkg\mod</code> as <code>github.com\gorilla</code> and inside of it we should have a mux folder with the latest version. So, the <code>go get</code> command is used to download and install a package along with its all dependencies.</p>
<h3>Set up a project for using a package</h3>
<p>Now, we have got the package so we can import it from anywhere in our go environment.</p>
<h4>Create a new folder (any name)</h4>
<p>You can test a go package from a isolated environment from the GOPATH by creating using the mod command.
The <code>mod init</code> command is a official way to create modules in golang and it creates kind of a environment to work on a templated project and structure the project/module/package properly.</p>
<pre><code>go mod init
</code></pre>
<h4>Install the packages</h4>
<p>We have already installed the package but that was a global install in the GOPATH, so we need to install it in this module.</p>
<pre><code>go get github.com/gorilla/mux
</code></pre>
<h4>Use the package</h4>
<p>Now, we can move into actually using the package in our source go file. We won't do any thing complicated just a simple web server. It's too easy don't worry!</p>
<pre><code class="language-go">package main

import (
	&quot;net/http&quot;

	&quot;github.com/gorilla/mux&quot;
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc(&quot;/&quot;, Server)

	http.ListenAndServe(&quot;:8000&quot;, router)
}

func Server(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(&quot;Hello Mux!&quot;))
}
</code></pre>
<p>We firstly setup a router(a pair of HTTP route with some logic) from the <a href="https://pkg.go.dev/github.com/gorilla/mux#NewRouter">NewRouter</a> function provided by <code>mux</code>. We'll attach a function to this newly created router by pairing a URL with a function. So, in simple terms when the mentioned URL is visited or a GET request is sent(don't get into too much details) we want a function to be invoked or called which does something. Finally we will set up a web server that listens at a port with the created router.</p>
<p>The final piece is the function which we will call when the URL is visited i.e. the <code>Server</code> function it can be any name. The function needs to have two arguments as it is invoked by a router, the writer and the request. The Writer is a Response writer i.e. to write the message to the server. We will simply use the <a href="https://pkg.go.dev/net/http#Header.Write">Write</a> function to simply print a array of bytes.</p>
<p>The type of the writer is specifically <code>http.ResponseWriter</code> as we want to write a simple HTTP response header. Also the request is a <a href="https://pkg.go.dev/net/http#Request">http.Request</a> type as we simply accept a HTTP request.</p>
<p>So, on running the following script, we will be able to see a simple HTTP response on the localhost at port <code>8000</code> or on your provided port.</p>
<pre><code>go run main.go
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1650645111/blog-media/gzje6ukyystp2x27u83o.png" alt="GO Gorilla MUX web server"></p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So, we were able to dive a bit deeper in golang packages and modules. We covered from installation to importing packages in Golang, we also touched on basics of initializing a module in Golang. Hopefully, we were able to get the basics covered when it comes to packages in Golang.</p>
<p>Thank you for reading. If you have any questions or feedback, please let me know in the comments or on social handles. Happy Coding :)</p>
