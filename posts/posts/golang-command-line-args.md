{
  "title": "Golang: Command Line Arguments",
  "status": "published",
  "slug": "golang-command-line-args",
  "date": "2025-04-05T12:33:33Z"
}

<h2>Introduction</h2>
<p>In the 25th post of the series, we will be taking a look into parsing of command line arguments in golang. We will be exploring how to do the basics of parsing and using the positional parameters or arguments from the command line in the program. By using standard library packages like <code>os</code> and <code>flag</code>, we can make powerful yet easy-to-build CLI apps and programs.</p>
<h2>Parsing Arguments from the command line (os package)</h2>
<p>We can use the os package to get the arguments from the command line in a go script. We have to use the Args variable in the os package. The <code>Args</code> variable is a slice of strings which thereby is the parsed arguments from the command line.</p>
<ul>
<li>
<p>The first (0 index) Argument is the path to the program</p>
</li>
<li>
<p>The 1st index onwards are the actual arguments passed.</p>
</li>
</ul>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;os&quot;
)

func main() {
	args := os.Args
    fmt.Printf(&quot;Type of Args = %T
&quot;, args)
	fmt.Println(args[0], args[1])
}
</code></pre>
<pre><code class="language-bash">$ go run main.go hello
Type of Args = []string
/tmp/go-build1414795487/b001/exe/main hello
</code></pre>
<p>In the above example, we can see that the <code>Args</code> is a slice of string and we can get the indices as per the arguments passed from the command line.</p>
<p>If you don't parse any arguments and access the 1st argument as <code>os.Args[1]</code> it will result in an <code>index out of range</code> error. So, you need to first check if the argument is parsed and set a default value otherwise.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;os&quot;
    &quot;strconv&quot;
)

func main() {
	var port int
	var err error
	if len(os.Args) &gt; 1 {
		port, err = strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
	} else {
		port = 8000
	}
	fmt.Println(port)
}
</code></pre>
<pre><code class="language-bash">$ go run main.go
8000

$ go run main.go 7000
7090

$ go run main.go h
panic: strconv.Atoi: parsing &quot;h&quot;: invalid syntax
</code></pre>
<p>In the above example, we have declared the port variable as an integer and tried to see if we had an argument parsed from the command line using the len function and if there was a variable, we will simply cast it into an integer using the <a href="https://pkg.go.dev/strconv#Atoi">strconv.Atoi</a> function. If there are any errors in the process, we log an error message and panic out of the program. So, this is how we can set default values or check for any arguments from the command line in golang.</p>
<h3>Get the number of args</h3>
<p>We can use the len function with the <code>Args</code> slice to get the total number of arguments from the command line. To ignore the first argument which would be the path to the program, we simply can slice the first element as <code>os.Args[1:]</code>. This will slice the list of the arguments from the first index till the last element in the slice.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;os&quot;
)

func main() {
	total_args := len(os.Args[1:])
	fmt.Println(&quot;Total Args =&quot;, total_args)
}
</code></pre>
<pre><code class="language-bash">$ go run main.go hello world 56

Total Args = 3
</code></pre>
<p>This will simply give us the number of arguments passed from the command line, excluding the first(0th) argument which is the default argument as the execution path of the current program.</p>
<h3>Iterate over all arguments</h3>
<p>We can use the simple for loop with range over the <code>os.Args</code> or <code>os.Args[1:]</code> for iterating over each of the arguments passed from the command line.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;os&quot;
)

func main() {
	for n, args := range os.Args {
		fmt.Println(&quot;Arg&quot;, n, &quot;-&gt;&quot;, args)
	}

    /* 
    // For excluding the 0th argument
	for n, args := range os.Args[1:] {
		fmt.Println(&quot;Arg&quot;, n, &quot;-&gt;&quot;, args)
	}
    */
}
</code></pre>
<pre><code class="language-bash">$ go run main.go hello world 56
Arg 0 -&gt; /tmp/go-build2248205073/b001/exe/main
Arg 1 -&gt; hello
Arg 2 -&gt; world
Arg 3 -&gt; 56
</code></pre>
<p>We can now iterate over the arguments passed from the command line using a simple for loop. We can further process these arguments per the program's requirements and need.</p>
<h2>Using flags package</h2>
<p>Golang has a package in its standard library called <a href="https://pkg.go.dev/flag">flags</a> which allows us to parse flags and arguments from the command line with a lot of built-in features. For instance, a default value is easily parsed with a simple function parameter, help text in case of an error in parsing the arguments or flag, customization and freedom for choosing a data type for the type of argument, and so on. For a bare-bones and quick CLI program, the flag package is a great choice.</p>
<h3>Parse Typed Flags</h3>
<p>We can use typed flag values using the functions provided in the <code>flags</code> package like <a href="https://pkg.go.dev/flag#IntVar">IntVar</a> for an integer value, <a href="https://pkg.go.dev/flag#StringVar">StringVar</a> for string, <a href="https://pkg.go.dev/flag#BoolVar">BoolVar</a> for boolean values and so on. Each function takes in 4 parameters and they set the value of the parsed variable from the parsed argument/flag from the command line.</p>
<ul>
<li>
<p>The first parameter is a reference to the variable to store the value.</p>
</li>
<li>
<p>The second parameter is the name of the argument/flag to be read from the command line.</p>
</li>
<li>
<p>The third parameter is the default value of the variable.</p>
</li>
<li>
<p>The fourth parameter is the help text for that argument/flag.</p>
</li>
</ul>
<p>So, let's take the previous example of port number parsing from the command line. We can use the <code>flag.IntVar(&amp;port, &quot;p&quot;, 8000, &quot;Provide a port number&quot;)</code>, this will set the value of the variable port from the command line as the value of <code>-p 6789</code> or the default value as <code>8000</code>. The help text will be used if the user has provided a non-integer or an invalid value as an error message.</p>
<pre><code class="language-go">package main

import (
	&quot;flag&quot;
	&quot;fmt&quot;
)

func main() {
	var port int
	var dir string
	var publish bool

	flag.IntVar(&amp;port, &quot;p&quot;, 8000, &quot;Provide a port number&quot;)
	flag.StringVar(&amp;dir, &quot;dir&quot;, &quot;output_dir&quot;, &quot;Directory&quot;)
	flag.BoolVar(&amp;publish, &quot;publish&quot;, false, &quot;Publish the article&quot;)

	flag.Parse()

	fmt.Println(port)
	fmt.Println(dir)
	fmt.Println(publish)

	if publish {
		fmt.Println(&quot;Publishing article...&quot;)
	} else {
		fmt.Println(&quot;Article saved as Draft!&quot;)
	}
}
</code></pre>
<pre><code class="language-bash">$ go run flag.go

8000
output_dir
false
Article saved as Draft!


$ go run flag.go -p 1234

1234
output_dir
false
Article saved as Draft!


$ go run flag.go -p 1234 -dir site_out

1234
site_out
false
Article saved as Draft!


$ go run flag.go -publish

8000
output_dir
true
Publishing article...
</code></pre>
<p>So, in the above, example, we have used a few types of values like <code>IntegerVar</code> for <code>port</code>, <code>StringVar</code> for <code>dir</code>, and <code>BoolVar</code> for <code>publish</code>. As explained earlier, the functions take 4 parameters in the same format, the reference to the variable to hold the parsed value, the name of the argument/flag, the default value the variable will hold, and the help text or usage string. The <a href="https://pkg.go.dev/flag#BoolVar">BoolVar</a> is slightly different but it works logically well, if we parse <code>-publish</code> the value will be set as <code>true</code> and <code>false</code> otherwise. You can manually add the value like <code>-publish true</code> and so on but it is not mandatory and understood as true.</p>
<p>In the above example, we have parsed different arguments in the output and displayed the values of these flags. If we don't specify a value, we can see the default value being parsed, in the case of the <code>bool</code> variable, the default value is taken as <code>false</code>. Hence we can see how easily we can use and parse flags from the command line in golang, it's simple, quick, and also extensible.</p>
<p>For other data types, the flag package has functions like <a href="https://pkg.go.dev/flag#Float64Var">Float64Var</a> for float64 values, <a href="https://pkg.go.dev/flag#DurationVar">DurationVar</a> for time duration values and <a href="https://pkg.go.dev/flag#TextVar">TextVar</a> for other types as inferred by the unmarshalling of the text.</p>
<h3>Set flags from the script</h3>
<p>We can set the value of a flag/argument from the script rather than from the command line using the <a href="https://pkg.go.dev/flag#Set">Set</a> method in the flag package. The <code>Set</code> method takes in two values as parameters the name of the argument and the value of that argument to set as. It returns an error if any arise during the setting of the argument.</p>
<pre><code class="language-go">package main

import (
	&quot;flag&quot;
	&quot;fmt&quot;
)

func main() {
	var port int
	var dir string
	var publish bool

	flag.IntVar(&amp;port, &quot;p&quot;, 8000, &quot;Provide a port number&quot;)
	flag.StringVar(&amp;dir, &quot;dir&quot;, &quot;output_dir&quot;, &quot;Directory&quot;)

	flag.Parse()

    fmt.Println(port)
    fmt.Println(dir)
    flag.Set(&quot;dir&quot;, &quot;dumps&quot;)
    fmt.Println(dir)
}
</code></pre>
<pre><code class="language-bash">$ go run flag.go -p 8080
8080
output_dir
dumps
</code></pre>
<p>So, it is clearly visible that the value of an argument can be changed within the script, it also changes the value of the associated variable. Remember, we gave the two-parameter as strings so the first parameter is the name of the argument and not necessarily the variable name.</p>
<h3>Use Reference to arguments (pointers)</h3>
<p>Also, there are functions like <a href="https://pkg.go.dev/flag#Int">Int</a>, <a href="https://pkg.go.dev/flag#Float64">Float64</a>, <a href="https://pkg.go.dev/flag#String">String</a>, <a href="https://pkg.go.dev/flag#Bool">Bool</a> in the flag package that can allow getting the values of the arguments without using the <code>Parse</code> method. We use the reference of the value stored in as the arguments instead of defining the variables as a data value; we have a pointer to that value of data.</p>
<pre><code class="language-go">package main

import (
	&quot;flag&quot;
	&quot;fmt&quot;
)

func main() {
	port := flag.Int(&quot;p&quot;, 8000, &quot;Provide a port number&quot;)
	dir := flag.String(&quot;dir&quot;, &quot;output_dir&quot;, &quot;Directory&quot;)
	publish := flag.Bool(&quot;publish&quot;, false, &quot;Publish the article&quot;)
	help := flag.Bool(&quot;help&quot;, false, &quot;Help&quot;)

	if *help {
		flag.PrintDefaults()
	} else {
		fmt.Println(*port)
		fmt.Println(*dir)
		flag.Set(&quot;dir&quot;, &quot;dumps&quot;)
		fmt.Println(*dir)

		fmt.Println(flag.NFlag())
		fmt.Println(flag.NArg())

		fmt.Println(*publish)

		if *publish {
			fmt.Println(&quot;Publishing article...&quot;)
		} else {
			fmt.Println(&quot;Article saved as Draft!&quot;)
		}
		vals := flag.Args()
		fmt.Println(vals)
	}
}
</code></pre>
<pre><code class="language-bash">$ go run flag.go -p 80 -dir node_mods 1234
80
node_mods
dumps
2
1
false
Article saved as Draft!
[1234]
</code></pre>
<p>As we can it performs the same task, but we have to use pointers as references to the arguments instead of storing them in an actual memory address. We have performed the same set of operations on the arguments and flags as we do with the other examples.</p>
<p>We first, use the <code>Int</code> method or other methods appropriate that <code>String</code> can be used in general use cases, the function returns a reference (memory address) of the actual stored value of the arguments/flag. We can access the value from its memory address using the <code>*</code> operator. We have covered the <a href="https://www.meetgor.com/golang-pointers/">pointer</a> arithmetic in the last part of the series. When we use <code>*port</code> we get the value from the memory address and thereby we can use it for the required task in the program, we can also store a copy of the variable by creating a new variable with the value of that argument.</p>
<h3>Parse Arguments</h3>
<p>So, if we want to parse flags, with a single value, we have seen the use of the <a href="https://pkg.go.dev/flag#Args">flag.Args</a> function to get the values of the arguments passed from the command line which don't have any flag labels attached to them(just raw arguments from the CMD). Just as we used the <code>os.Args</code> variable but this function is much clean and filtered out the path to the program argument. So we can directly have the arguments which are clearly passed by the user from the command line.</p>
<pre><code class="language-go">package main

import (
	&quot;flag&quot;
	&quot;fmt&quot;
)

func main() {
	var port int
	flag.IntVar(&amp;port, &quot;p&quot;, 8000, &quot;Provide a port number&quot;)
	flag.Parse()
	fmt.Println(port)
	vals := flag.Args()
	fmt.Println(vals)
}
</code></pre>
<pre><code class="language-bash">$ go run flag.go -p 8123
8123
[]


$ go run flag.go -p 8123 1234 hello true
8123
[1234 hello true]


$ go run flag.go -p 8123 1234 hello true -p 9823 world
8123
[1234 hello true -p 9823 world]
</code></pre>
<p>In the above example, we can see that we have used a few non-flagged arguments from the command line. The return value of the <code>Args</code> function is a slice of string, we can then convert it into appropriate types using type casting and functions. Once the flagged arguments are parsed, if we use the <code>Args</code> function, it won't be possible to again use flagged arguments in the command line. It will be considered a simple string thereafter.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/files/write/">100 days of Golang</a> GitHub repository.</p>
<h3>Get Help text with PrintDefaults</h3>
<p>We can use the <a href="https://pkg.go.dev/flag#PrintDefaults">flag.PrintDefaults</a> method for just printing the default values and the help text for the expected arguments from the command line in the script. We can simply use it as a help flag or use it in error messages for guiding the user to the proper arguments and flags.</p>
<pre><code class="language-go">package main

import (
	&quot;flag&quot;
	&quot;fmt&quot;
)

func main() {
	var port int
	var help bool
	flag.IntVar(&amp;port, &quot;p&quot;, 8000, &quot;Provide a port number&quot;)
	flag.BoolVar(&amp;help, &quot;help&quot;, false, &quot;Help&quot;)
	flag.Parse()
	if help {
		flag.PrintDefaults()
	} else {
		fmt.Println(port)
		vals := flag.Args()
		fmt.Println(vals)
	}
}
</code></pre>
<pre><code class="language-bash">$ go run help.go -h

Usage of /tmp/go-build121267600/b001/exe/help:
  -help
        Help
  -p int
        Provide a port number (default 8000)


$ go run help.go

8000
[]
</code></pre>
<p>So, we can see the <code>PrintDefaults</code> function will simply print the helper text for the flags expected in the script and the default value of those flags as well. This can be used to provide a good user-friendly interface for a simple terminal application.</p>
<h3>Get the number of arguments</h3>
<p>We can use the <a href="https://pkg.go.dev/flag#NFlag">NFlag</a> method in the <code>flag</code> package. The function returns an integer that indicates a count of the arguments that have been set from the command line.</p>
<pre><code class="language-go">package main

import (
	&quot;flag&quot;
	&quot;fmt&quot;
)

func main() {
	var port int
	var dir string
	var publish bool

	flag.IntVar(&amp;port, &quot;p&quot;, 8000, &quot;Provide a port number&quot;)
	flag.StringVar(&amp;dir, &quot;dir&quot;, &quot;output_dir&quot;, &quot;Directory&quot;)

	flag.Parse()

    fmt.Println(port)
    fmt.Println(dir)
    fmt.Println(flag.NFlag())
}
</code></pre>
<pre><code class="language-bash">$ go run flag.go
8000
output_dir
0


$ go run flag.go -p 8080 8999 false hello
8080
output_dir
1


$ go run flag.go -p 8080 -dir dumps hello 1234
8080
dumps
2
</code></pre>
<p>The <code>port</code> flag has been set from the command line, so we just have one argument set, hence the function <code>NFlag</code> returns <code>1</code> as the number of set flags.</p>
<p>Also, the <a href="https://pkg.go.dev/flag#NArg">NArg</a> method will return an integer that will count the number of arguments that have been provided leaving out the flag arguments.</p>
<pre><code class="language-go">package main

import (
	&quot;flag&quot;
	&quot;fmt&quot;
)

func main() {
	var port int
	var dir string
	var publish bool

	flag.IntVar(&amp;port, &quot;p&quot;, 8000, &quot;Provide a port number&quot;)
	flag.StringVar(&amp;dir, &quot;dir&quot;, &quot;output_dir&quot;, &quot;Directory&quot;)

	flag.Parse()

    fmt.Println(port)
    fmt.Println(dir)
    fmt.Println(flag.NArg())
}
</code></pre>
<pre><code class="language-bash">$ go run flag.go 1234
8000
output_dir
1


$ go run flag.go -p 8080 -dir dumps hello 1234
8080
dumps
2


$ go run flag.go -p 8080 hello 1234 false
8080
dumps
3
</code></pre>
<p>In the first example, we don't have any flag arguments set, we just have one unflagged argument as <code>1234</code>, hence the <code>NArg</code> function returns <code>1</code>. The second example has 2 values that are not flagged, we have set the values of <code>port</code> and <code>dir</code> as <code>8080</code> and <code>dumps</code> respectively, so the remaining unflagged values are <code>hello</code> and <code>1234</code> hence the return value as <code>2</code>. The third example has 3 unflagged values as <code>hello 1234 false</code>, hence we return <code>3</code>.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/cmd-args/">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>We have seen how to parse command line arguments in golang with the <code>os</code> and the <code>flag</code> packages. Though these two are not the only options for building CLI applications, they provide a clean and easy-to-start approach, also they come with the standard library which makes it even better as we don't have to mingle with third-party libraries. We saw the basics of parsing flags and arguments from a command line program.</p>
<p>Thank you for reading. If you have any queries, questions, or feedback, you can let me know in the discussion below or on my social handles. Happy Coding :)</p>
