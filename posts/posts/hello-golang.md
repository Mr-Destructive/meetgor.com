{
  "title": "Golang: Installation and Hello-World",
  "status": "published",
  "slug": "hello-golang",
  "date": "2025-04-05T12:33:44Z"
}

<h2>Introduction</h2>
<p>Moving on to the second day, we will be installing and setting up Go lang on our systems. The installation and setup are quite simple and not much demonstration is required, so further in the article, I will also make a hello-world program in GO. We will explore the basic program in GO and how to compile, run and build a GO program in this section.</p>
<h2>Installing Go</h2>
<p>Installing Go is pretty straightforward. You have to install the binaries from the official website as per your operating system.</p>
<p>Head on to the <a href="https://go.dev/">go.dev</a> which is the official website for the Go language. Click on the <a href="https://go.dev/doc/install">Download</a> Tab and there should be all the required information. Install the installer as per your specific operating system.</p>
<p>If you wish not to lead to any errors, keep the configuration for the installer as default and complete the installation process.</p>
<h2>Setting up Environment variables</h2>
<p>To make sure Go lang is perfectly installed in your system, type in CMD/Terminal/Powershell the following command:</p>
<pre><code>go version
</code></pre>
<p>If you get a specific version of golang along with the platform and architecture of your system, you have successfully installed Go lang in your system.</p>
<pre><code>$ go version
go version go1.17.8 windows/amd64
</code></pre>
<p>If you get an output as a command not found or anything else, this is an indication that your Go installation was not successful. You need to configure your Environment variables properly or re-install the installation script.</p>
<pre><code>$ go version
bash: go: command not found
</code></pre>
<h2>Hello Gophers</h2>
<p>Once the Go lang has been successfully installed in your system, we can start writing our first program. Yes, a <code>Hello World</code> program. It is not as simple as <code>print(&quot;hello world&quot;)</code> but a lot better than 10-15 lines of Java or C/C++.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
)

func main() {
	fmt.Println(&quot;Hello,Gophers!&quot;)
}
</code></pre>
<p>So, this is so-called the <code>hello-world</code> program in Go, we will see each of the terms in detail next. But before that, let's get an idea of the style of the code. It is definitely similar if you are coming from C/C++ or Java, the package the main function. It will even feel like Python or Javascript when we explore other aspects. It has really a unique style of programming but feels familiar to programmers coming from any programming language, this might not be true for all programming languages though.</p>
<h3>Packages</h3>
<p>A package in Go lang is a way to bundle up some useful functions or any other constructs. Using packages we can reuse components of a specific app in another. Packages in golang also help in optimizing the execution/compilation time by letting the compiler only compile the required packages.</p>
<p>Here, we have the main package, which provides the entry point for the entire project. This is mandatory for creating executable programs as we need to have a start point. The file name can be anything, but for executing the code, you need to have a main package where your main function resides. It provides an entry point as a package, when we will run the file, we provide the file which actually acts as a package and the file that has a tag of main is the entry point for the program.</p>
<h3>Main Function</h3>
<p>We have the main function where the main package is defined. It acts as a start point for a package. The main package will look for a main function inside the package. The main function doesn't take any parameter and nor does it return any value. When the function's scope ends, the program exits.</p>
<p>The main function has significance only in the main package, if you define the main function in other packages excluding <code>main</code>, it works as a normal function.</p>
<h3>Import Statements</h3>
<p>We have an <code>import</code> keyword for importing packages from the standard library or other external packages from the internet. There are a lot of ways to import packages in golang like single, nested, multiple, aliased, dot import, and blank imports. We will see these different import styles in a dedicated section. Right now, we are directly importing a package, a single package. The pacakge is called the <a href="https://pkg.go.dev/fmt">fmt</a> pacakge. It handles the format, input, and output format in the console. It is a standard package to perform some basic operations in golang.</p>
<p>We can import it as a single direct import like:</p>
<pre><code class="language-go">import &quot;fmt&quot;
</code></pre>
<p>OR</p>
<p>Make multiple imports like:</p>
<pre><code class="language-go">import (
    &quot;fmt&quot;
)
</code></pre>
<p>We can add multiple packages on each line, this way we do not have to write the keyword <code>import</code> again and again. It depends on what you want to do in the program.</p>
<h3>Println function</h3>
<p>We can access the functions from the imported packages, in our case we can use the functions from the <code>fmt</code> package. We have access to one of the functions like <code>Println</code>, which prints string on a new line. Syntactically, we access the function and call it by using the <code>dot</code> operator like:</p>
<pre><code class="language-go">fmt.Println(&quot;Hi there!&quot;)
</code></pre>
<p>The <a href="">Println</a> function takes in a parameter as a string and multiple optional parameters that can be strings or any variable. We will see how to declare variables and types in the next section.</p>
<p>Here, the <code>P</code> in <code>Println</code> has significance as it allows us to distinguish private methods(functions in structs aka classes) from public methods. If a function begins with an uppercase letter, it is a public function. In technical terms, if the first letter of a method is upper case, it can be exported to other packages.</p>
<h2>Running Scripts</h2>
<p>Let's run the code and GO programming language to our resume. You can run a go source file assuming it has a main package with the main function using the following command:</p>
<pre><code>go run &lt;filename.go&gt;
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1648833567/blog-media/o7i9spph2bfgemqonq8p.png" alt="GO run command"></p>
<p>This will simply display the string which we have passed to the <code>Println</code> function. If you didn't have a main package this command won't run and return you an error:</p>
<pre><code>package command-line-arguments is not the main package
</code></pre>
<p>By executing the run command, we can are creating a executable in a system's temp folder,</p>
<p>For Windows, it's usually at:</p>
<pre><code>C:\Userscer\AppData\Local
</code></pre>
<p>You can get the location of the temp directory using CMD/PowerShell:</p>
<pre><code>CMD:
echo %TEMP%

PowerShell:
$env:Temp
</code></pre>
<p>For Linux</p>
<pre><code>/tmp
</code></pre>
<p>You can get the location of the temp folder using Terminal in Linux/macOS:</p>
<pre><code>echo $TMPDIR
</code></pre>
<p>It doesn't create any executable in the current project or folder, it only runs the executable that it has built in the temp folder. The run command in simple terms <strong>compiles and executes the main package</strong>. As the file provided to the run command needs to have the main package with the main function, it will thus compile that source code in the provided file.</p>
<p>To get the location of the executable file that was generated by the <code>run</code> command, you can get the path using the following command:</p>
<pre><code>go run --work &lt;filename&gt;.go
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1648833643/blog-media/maqfd73xmiivrckc2acn.png" alt="GO Run TMP file"></p>
<p>This will print the path to the executable that it currently has compiled.</p>
<p>For further readings on the <code>run</code> command in Go, you can refer to the <a href="https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program">documentation</a>.</p>
<h2>Creating Executables</h2>
<p>We can go a step further by creating binary/executables with our source file using the <code>build</code> command:</p>
<pre><code>go build &lt;filename&gt;.go
</code></pre>
<p>If you run this you would get an error as building a package requires a few things. The most important is the mod file.</p>
<pre><code>go: cannot find main module, but found .git/config in D:\meet\Code\go\100-days-of-golang
    to create a module there, run:
    cd .. &amp;&amp; go mod init
</code></pre>
<p>We need to create a mod file first before we build our script.
A mod file in golang is the file that specifies the go version along with the packages and dependencies. It is like the <code>requirement.txt</code> but a bit different.</p>
<p>We use the following command with the file that contains the main package among the other packages in the folder. We can even provide other packages to add to the mod file(see in detail in the future)</p>
<pre><code>go mod init &lt;filename&gt;.go
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1648833683/blog-media/pdvkdpnanl3e0yzoenqb.png" alt="GO Mod Init"></p>
<p>This will generate a <code>go.mod</code> file, this is a file that contains the list of dependencies and packages in the project.
If you look at the mod file, it looks as follows:</p>
<pre><code class="language-go">module script.go

go 1.17
</code></pre>
<p>Currently, this is pretty simple and has very little detail, but as your project increases in complexity, this file populates with the modules and packages imported and used in the project.</p>
<p>So, after creating the mod file, we can build the script which we tried earlier.</p>
<pre><code>go build &lt;filename&gt;.go
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1648833745/blog-media/i0hnwsxxl0gglhh3rdoe.png" alt="GO Build Command"></p>
<p>So, this command generates an exe right in the current folder. This will generate the file named after the package which is mainly <code>filename.exe</code>.</p>
<p>If you have a <code>go.mod</code> file in your project, just running the command will generate the exe file:</p>
<pre><code>go build
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1648833832/blog-media/sdw9zrlff3odtnhallyq.png" alt="GO Build Command - Directory level"></p>
<p>NOTE: For the above command to work, you need to be in the directory which has the mod file for your project. It basically bundles the listed packages and creates the executable named after the package which is named <code>main</code>. Thus it generates a different file name as <code>filename.go.exe</code></p>
<p>We can also provide an output file as the exe file name, this can be done with the following command:</p>
<pre><code>go build -o &lt;filename&gt;
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1648833932/blog-media/t5dfhx0va7reyjjygfnq.png" alt="GO Build Output file"></p>
<p>For further reading on the <code>go build</code> command, head over to this <a href="https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies">documentation</a> page.</p>
<p>Link to all of the code and resources is mentioned in this <a href="https://github.com/Mr-Destructive/100-days-of-golang">GitHub</a> repository.</p>
<h2>Conclusion</h2>
<p>So, from this second post, we were able to set up the Go language in our system and write our first <code>hello-world</code> program. This was a bit more than the setup and installation guide as it is quite boring to demonstrate the installation procedure being quite straightforward. Hopefully, you followed so far and you were able to pick up things in the go landscape. Thank you for reading and if you have any questions, suggestions, or feedback, let me know in the comments or the provided social handles. See you tomorrow, until then Happy Coding :)</p>
