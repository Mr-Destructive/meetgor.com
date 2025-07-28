{
  "title": "Golang: Closures",
  "status": "published",
  "slug": "golang-closures",
  "date": "2025-04-05T12:33:36Z"
}

<h2>Introduction</h2>
<p>In the previous part of the series, we covered <a href="https://meetgor.com/golang-anonymous-functions/">anonymous functions</a> and in this section, we will look into <code>closures</code> which are quite a cool concept for various things. Closures are basically a function that returns a function instead of a value, so basically we will leverage anonymous functions for creating closures.</p>
<h2>Simple Closures</h2>
<p>A simple closure can be constructed for understanding how we can use closures in golang. We will return a function from a function, that is a simple closure. So, in the below code example, we have created a function <code>gophy()</code> which takes no parameters but returns a function that returns a <code>string</code>.  The function simply returns an anonymous function that returns a string.</p>
<p>We will initialize the variable <code>g</code> that is assigned to the function <code>gophy</code> which will simply return a function call. We are not calling the function simply returning the call to the function <code>gophy</code> that has the return value as the anonymous function. We will simply have the function in the variable <code>g</code> rather than the simple value string. So we will have to call the <code>g</code> variable for actually returning the string.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func gophy() func() string{
  return func() string{
    return &quot;Hello, Gophers!&quot;
  }
}

func main() {

	// using clousure/anonymous function to return a value
	// that value can be assigned to the variable
	g := gophy()
	fmt.Println(g())
}

</code></pre>
<pre><code>$ go run simple.go
Hello, Gophers!

</code></pre>
<p>So, that is how we can call the function <code>g</code> that will return a string, so we have the function body stored in the variable <code>g</code>. We can call it as many times as we want.</p>
<h2>Variable Scope in Closures</h2>
<p>We can even use variables that will remain in the function scope once it is initialized. So, let’s say we have a function that will increment the counter, but if we want to keep the counter the same throughout the program, we might have to use a global variable so as to maintain the context, but with closures, we will retain the value once we have initialized the function call.</p>
<p>In the below example, we are creating the function <code>incrementer</code> that is a closure with int as the return type. We are initializing the variable <code>counter</code> that will be acting as the counter in the program, the function returns an anonymous function that will increment the counter and return it.</p>
<p>Here, when we create an instance of the <code>increment</code> function it basically initializes the <code>counter</code> to <code>0</code> and returns the anonymous function as a call. Now, <code>c</code> will act as a function that has the counter variable bound to it and we can call <code>c</code> that will, in turn, call the anonymous function keeping the scope of the <code>counter</code> variable.  So, each time we call the function <code>c</code> it will increment the counter and thus we keep the counter inside the scope of the function <code>incrementer</code> in the <code>c</code> variable.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func inrementer() func() int{
  counter := 0
  return func() int{
    counter += 1
    return counter
  }
}

func main() {

  c := inrementer()

  fmt.Println(c())
  fmt.Println(c())
  fmt.Println(c())
  fmt.Println(c())
  fmt.Println(c())

}

</code></pre>
<pre><code>$go run simple.go
1
2
3
4
5

</code></pre>
<p>If we want to extend the functionality,  we can even assign the function call <code>c()</code> to a variable and access the returned value which will be the current state of the counter.</p>
<p>We can even use different scope or closures tied to a particular function, that is we can bind data to a different instances of a closure.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func inrementer() func() int{
  counter := 0
  return func() int{
    counter += 1
    return counter
  }
}

func main() {

  c1 := inrementer()

  fmt.Println(c1())
  fmt.Println(c1())
  fmt.Println(c1())

  c2 := inrementer()

  fmt.Println(c2())
  fmt.Println(c2())
  fmt.Println(c2())
  fmt.Println(c2())

}

</code></pre>
<pre><code>$go run simple.go
1
2
3
1
2
3
4

</code></pre>
<p>Here we have <code>c1</code> and <code>c2</code> forming different closures and thereby we can have different scopes of the variables associated with it. The variable is bound to the instance it which was initialized, so we can see the different closure instances having different values.</p>
<h2>Factorial of a function with Closures</h2>
<p>We can create some interesting programs with closures, we will implement the calculation of factorial with closures in golang.</p>
<p>This will be a <code>factorial</code> function that returns an anonymous function with the return type as <code>int</code>. The function will initialize the variable <code>fact</code> which will store the actual factorial value and <code>n</code> as the initial number for calculating the factorial of it.</p>
<p>Inside the anonymous function, we will calculate the factorial and increment the number and simply return the factorial value from the function. The <code>fact</code> variable will contain the factorial of the number n, so here we can leverage the use of closures as we will maintain the state of the variable <code>fact</code> and <code>n</code> from the previous calls or the initialization of the function.</p>
<p>Inside the <code>main</code> function, we have created the <code>f</code> variable and called the <code>factorial</code> function, so that will initialize the <code>fact</code> and <code>n</code> of the variable and thereby returning the anonymous function call. Now we can call the variable <code>f</code> as many times as we want that will simply return the factorial of the number incremented each time we call the function.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func factorial() func() int{
	fact, n := 1, 1
	return func() int{
    fact = fact * n
    n += 1
		return fact
	}
}

func main() {

  f := factorial()
  fmt.Println(f())
  fmt.Println(f())
  fmt.Println(f())
  fmt.Println(f())
  fmt.Println(f())

}

</code></pre>
<pre><code>$ go run simple.go
1
2
6
24
120

</code></pre>
<p>So, we can see that the factorial is getting printed for each call and the number is being incremented at each call.</p>
<p>So that's the basics of closures in golang, we can use closures to keep the content secured and encapsulated from different function calls. We can bind data with closures, with the help of anonymous functions a closure can be constructed and data can be bound to a particular function called scope.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/closures/main.go">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>From this post, we could understand the fundamentals of closures in golang. The basic concept of closures in golang was understood with a few examples. Thank you for reading, if you have any queries or feedback please leave them in the comments or on my social handles. Happy Coding :)</p>
