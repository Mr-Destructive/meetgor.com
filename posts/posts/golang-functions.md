{
  "title": "Golang: Functions",
  "status": "published",
  "slug": "golang-functions",
  "date": "2025-04-05T12:33:43Z"
}

<h2>Introduction</h2>
<p>In the eighth part of the series, we will be exploring functions in golang. We will be diving into some basics of functions in golang like declaration, definition and calling. We won't be exploring all the topics of functions as it is quite a large topic to cover in one shot. So, building from the base, we will be starting from the basic declaration to simple return statements.</p>
<h2>Functions in Golang</h2>
<p>Functions in golang are a simple way to structure a block of code that can be re-usable. Functions also allow us to process a piece of logic and return the output. Functions allow us to write readable and scalable code as we have to write the code once and we can re-use the functionality of it by calling it.</p>
<h2>Declaring Functions</h2>
<p>We have already defined a function, if you have followed the series so far, or even written a <code>hello-world</code> application. The <code>main</code> function is the most fundamental function we can define in golang. The main is complicated if dive deeper but in the simplest of term it acts as a entry point for the entire program.</p>
<pre><code class="language-go">package main

func main() {

}
</code></pre>
<p>We have written the above code a lot of times till now, but never really talked about it's significance. Here we will understand the terminologies related to the main function. A function is declared with the <code>func</code> keyword along with the name of the function. There needs to be the <code>()</code> parenthesis after the name of the function, optionally it can take parameters inside the parameters to be used inside the function.</p>
<p>We define the core functionality or the core logic of the function inside the braces <code>{}</code>. We also have the <code>return</code> keyword which can return values from the function to the block where we have called the function. Usually, we call a function from other function (most of the times it's the <code>main</code> function). The <code>return</code> keyword is not mandatory and it is usually added at the end of the function block, just before the closing braces <code>}</code>.</p>
<pre><code class="language-go">func hello_world_007() {

}
</code></pre>
<p>We can define a custom function outside the main function by giving it a appropriate name. For the time bwing we can leave it empty and further define the logic of the actual function.</p>
<p>The name of the function can be given as per the following standards:</p>
<ul>
<li>Using letters<code>a-z A-Z</code>, numbers<code>0-9</code>, underscore <code>_</code> as a name.</li>
<li>Should not contain any spaces in-between the name.</li>
<li>Should not begin with a number or underscore.</li>
</ul>
<h2>Defining Functions</h2>
<p>Inside the <code>{}</code> we define the actual functionality/logic of the function. The variables inside the function will remain local to the function and can't be accessed or altered from outside the function, though if we really want to access some global variables(from main or other functions) we can pass parameters, we will look into it in the next few sections. For time being, we will be focusing on the actual code block inside the function.</p>
<pre><code class="language-go">
func hello_world() {
    fmt.Println(&quot;Hello World&quot;)
}
</code></pre>
<p>This is a basic function that just calls another function <code>Println</code> from the fmt package, which basically prints text in the console. Though, we are using the function Println, it won't print the content to the string as we are not using/calling the function. Now, we can get a step ahead and start working with variables inside the function.</p>
<pre><code class="language-go">func hello_world() {
    name := &quot;Gopher&quot;
    fmt.Println(&quot;Hello&quot;, name)
}
</code></pre>
<p>We have now added the local variable <code>name </code>inside the function, so this variable can only we referred inside the particular function.</p>
<h2>Calling Functions</h2>
<p>We can call the function from the main function or any other function by just specifying the name along with the <code>()</code> and optionally the parameters inside the parenthesis.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
    hello_world()    
}

func hello_world() {
    name := &quot;Gopher&quot;
    fmt.Println(&quot;Hello&quot;, name)
}

</code></pre>
<pre><code>$ go run func.go
Hello Gopher
</code></pre>
<p>So, we define the function <code>hello_world</code> and call the function by using the statement <code>hello_world()</code> inside the main function and now, we are able to actually run the function.</p>
<h2>Passing Parameters</h2>
<p>We can optionally parse variables from a function to other and process it for further computation and programming. So, we can pass parameters in a function by specifying the name to be used inside the function followed by the type of that variable.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {

	greet_me(&quot;Meet&quot;)
	n := &quot;John&quot;
	greet_me(n)
}

func greet_me(name string) {
	fmt.Println(&quot;Hello,&quot;, name, &quot;!&quot;)
}
</code></pre>
<pre><code>$ go run func.go
Hello, Meet !
Hello, John !
</code></pre>
<p>We have used the parameter <code>name</code> as a string in the function and used it inside the function body. The parameter name which is to be called from the main function can be anything and not necessarily be the same as declared in the function declaration. For instance, we have used the variable in the main function <code>n</code> which is passed in the function call. We can even pass the value as it is to the function in golang.</p>
<h2>Return Keyword</h2>
<p>We can use the return keyword to actually return a value from the function and not just display the message. The returned value can be later used from other places in the program.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {

	// return value
	y := line_eq(3, 1, 2)
	fmt.println(&quot;for x = 3 , y = &quot;, y)
}

func line_eq(x int, m int, c int) int {
	return ((m * x) + c)
}
</code></pre>
<pre><code>$ go run func.go
for x = 3 , y =  5
</code></pre>
<p>So, here we are able to fetch the returned value from the function and store it in another variable and further compute the required logic. We also need to specify the return type of the function after the parameters like <code>func (parameters) return-type { }</code>. Here, we need to return the specified type of the return value from the function else it would give a compilation error.</p>
<p>So, we basically need to provide the return value and also the return statement to capture the value from the function call.</p>
<h3>Multiple return values</h3>
<p>We can also provide multiple return values by providing a list of return values like <code>(type1 type2 type3 ....)</code>. We can return the values by separating the values by a comma. So, while calling the function, we need to specify the variables again as comma-separated name and this will capture the value from the function call.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {

	// multiple return values
	s, c, odd := sqube(5)
	fmt.Println(&quot;for x = 5 , x^2 =&quot;, s, &quot;x^3 =&quot;, c)
	if odd == true {
		fmt.Println(&quot;x is odd&quot;)
	} else {
		fmt.Println(&quot;x is true&quot;)
	}
}

func sqube(x int) (int, int, bool) {
	square := x * x
	cube := x * x * x
	var is_odd bool
	if x%2 == 0 {
		is_odd = false
	} else {
		is_odd = true
	}
	return square, cube, is_odd
}
</code></pre>
<pre><code>$ go run func.go
for x = 5 , x^2 = 25 x^3 = 125
x is odd
</code></pre>
<p>So, we have returned multiple values from the function like two integers and one boolean. The parameter is a single integer, now we need to parse 3 variables in order to capture all the values from the function call. Thus, we are able to get all the values from the function.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So, from this part of the series, we are able to understand the basics of functions in golang. We covered from declaration, definition and simple return statements and function calling. Thank you for reading. If you have any questions or feedback, please let me know in the comments or on social handles. Happy Coding :)</p>
