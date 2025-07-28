{
  "title": "Golang: Pointers",
  "status": "published",
  "slug": "golang-pointers",
  "date": "2025-04-05T12:33:42Z"
}

<h2>Introduction</h2>
<p>In the tenth part of the series, we will be looking into an interesting concept in programming i.e. Pointer. It's a simple thing but a really powerful concept. Using pointers we can do several things very easily rather than writing a lot of code for a simple thing. We will be looking into basic concepts like declaration, referencing, de-referencing, and some examples on passing by reference, along with a pointer to struct instances.</p>
<h2>Pointers</h2>
<p>Pointers are simple, it's just their use case that makes it a big concept. Pointers are really powerful, they can do a lot of things that might seem impossible for a given problem. A pointer is a variable but unlike another variable which stores values in the form of integers, string, boolean, etc. pointers store the memory address. Memory address can be any valid location in memory that generally holds a variable.</p>
<p>So, using pointers we can play with the memory address of variables and modify the contents of the variable directly using the memory address rather than accessing the variable. In golang, we have ways to store pointers and perform operations for the same.</p>
<h2>Declaring Pointers</h2>
<p>To declare pointers in golang, we can use the <code>*</code> before the type of data type we want to refer to. This means a pointer needs to specify which data type it is referencing as a measure of caution to mismatch types in the variable. Initially, the pointer variable is mapped to <code>&lt;nil&gt;</code> that is it points to nothing but a null pointer.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
	var ptr *int
	fmt.Println(ptr)
}

</code></pre>
<pre><code>$ go run pointer.go
&lt;nil&gt;
</code></pre>
<p>As we can see, the pointer that references an integer is initialized to nil. We have used <code>*</code> before the data type, this can be anything like <code>*string</code>, <code>*bool</code>, <code>*float64</code>, etc.</p>
<h2>The * and &amp; in Pointers</h2>
<p>After declaring a pointer, we can now move into assigning a pointer a memory address. Using the <code>&amp;</code> or the ampersand operator we can get the memory address of a variable.</p>
<pre><code class="language-go">var n := 34
var a_pointer *int = &amp;n
fmt.Println(a_pointer)
</code></pre>
<pre><code>$ go run pointer.go
0xc0000a6080
</code></pre>
<p>Here, we can see that the pointer variable is storing the memory address of an integer variable. Don't worry about the value of the pointer variable, it is just a memory location on your machine. So, we use the <code>&amp;</code> to access the memory address of any variable.</p>
<p>We have seen that the <code>*</code> is used to declare a pointer variable, but it is also used for dereferencing a pointer. So, if we used <code>&amp;</code> to get the memory address of a variable, similarly we can use the <code>*</code> to get back the value from the memory address. Both are opposite in terms of accessing the value.</p>
<pre><code class="language-go">n := 34
var a_pointer *int = &amp;n
fmt.Println(a_pointer)
m := *a_pointer
fmt.Println(m)
</code></pre>
<pre><code>$ go run pointer.go
0xc0000a8080
34
</code></pre>
<p>As we can see, we have accessed the value stored in the pointer variable(<code>a_pointer</code>) by using the <code>*</code>. Here, the variable which we have created <code>m</code> will be of type whatever is stored in the memory address of the provided pointer variable. In this case, it is <code>int</code>, it can anything.</p>
<p>So, this is how <code>*</code> and the <code>&amp;</code> work in Golang. The <code>*</code> is used for declaring pointer variables as well as de-referencing pointer variables, and the <code>&amp;</code> operator is used for accessing the memory address of the variable.</p>
<p>That's basically the concept of pointers in golang. It's that simple. Using the simple concept of referencing and de-referencing, we can perform some operations like passing by reference to functions which will allow us to actually pass the value rather than the copy of the variable's value.</p>
<h2>Passing by Reference to Function</h2>
<p>Now we have the fundamentals of pointers cleared, we can move into actually using them to do some really useful operations. Generally, when we use parameters such as integers, strings, bool, etc. we are passing the copy of the variables into the function rather than the actual value of the variable. This is where pointers come in. By using pointers to pass the memory address of the variables we need to pass in we actually pass the location of the variables.</p>
<p>Let's take a look at a simple example of a function that swaps the value of two variables.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {

    x := 3
	y := 6
	k := &amp;x
	p := &amp;y
	fmt.Printf(&quot;Before swapping : x = %d and y = %d.
&quot;, x, y)
	swap(k, p)
	fmt.Printf(&quot;After swapping  : x = %d and y = %d.
&quot;, x, y)
}
</code></pre>
<pre><code>$ go run pointer.go
Before swapping : x = 3 and y = 6.
After swapping  : x = 6 and y = 3.
</code></pre>
<p>We can see here, that we have used pointers to pass the value of parameters to a function. Without using pointers, the value of the variable is passed as a copy but by using pointers, we are actually passing the memory address. In the main function, we first store the memory address of two variables <code>x</code> and <code>y</code> into two different pointer variables. We now can construct a function that accepts two memory addresses and perform further operations.</p>
<p>Inside the function, we have de-referenced the pointer variables as with <code>*</code>. Don't confuse <code>x *int</code> with <code>*x</code>. We use <code>x *int</code> to make the function realize that we are passing a pointer variable of an integer value, and <code>*x</code> is used to de-reference the memory address which is stored in <code>x</code>.</p>
<p>So, simply we</p>
<ul>
<li>store the value in the memory location stored at <code>x</code> in the temp variable</li>
<li>store the value at the memory address stored in <code>y</code> into the memory address <code>x</code>.</li>
<li>store the value of the temp variable into the memory address stored in <code>x</code>.</li>
</ul>
<p>We have successfully swapped two values without returning any values from the function.</p>
<h2>Pointer to a Struct Instance/Object</h2>
<p>We can now even modify the values of Struct objects/instances by referencing the instance to a pointer. By assigning the pointer variable to a struct instance, we have access to its associated properties and function. Thereby we can modify the contents directly from the pointer variable.</p>
<p>Let's take a look at a basic example of modifying properties using a pointer to a struct instance.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

type Book struct {
	pages int
	genre string
	title string
}

func main() {
	new_book := Book{120, &quot;fiction&quot;, &quot;Harry Potter&quot;}
	fmt.Println(new_book)
	fmt.Printf(&quot;Type of new_book -&gt; %T
&quot;, new_book)
	book_ptr := &amp;new_book
	book_ptr.title = &quot;Games of Thrones&quot;
	fmt.Println(new_book)
}

</code></pre>
<pre><code>$ go run pointer.go
{120 fiction Harry Potter}
Type of new_book -&gt; main.Book
{120 fiction Games of Thrones}
</code></pre>
<p>So, we have created a pointer variable of the type which is a struct <code>Book</code>, this gives us access to the memory addresses associated with various properties defined in the struct. Using the pointer variable, we can access properties and thereby change the value directly as we have the memory address stored in <code>book_ptr</code>. So, if we say <code>book_ptr.title = &quot;Games of Thrones&quot;</code>, we are storing the string directly into the memory address of the <code>new_book</code> object as <code>book_ptr</code> refers to the memory addresses to the struct object <code>new_book</code>.</p>
<p>Here, we have literally changed the value of a property in a struct object using pointers. This is really powerful and time-saving. If pointers were not a thing, you would have to write a separate function for doing the same.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So, that's it we have covered enough basics of pointers so that we are able to understand the working of simple scripts or programs. Even simple data structures like slices or strings can be understood by using pointers in golang. From this part of the series, we were able to understand the declaration, referencing, and de-referencing of pointers along with passing by reference to functions and creating pointers to struct instances.</p>
<p>Thank you for reading. If you have any questions or feedback, please let me know in the comments or on social handles. Happy Coding :)</p>
