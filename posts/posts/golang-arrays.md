{
  "title": "Golang: Arrays",
  "status": "published",
  "slug": "golang-arrays",
  "date": "2025-04-05T12:33:43Z"
}

<h2>Introduction</h2>
<p>In this fifth section of Golang, we will be understanding the basics of arrays. We will be covering some basic operations on arrays like declaration, initialization, modifications, and iterating over an array.</p>
<h2>Declaring Arrays</h2>
<p>Arrays are type of data structure that allow us to store multiple items at continuous memory locations of the same type. In golang, we can create arrays similar to any variable but by adding a few bits and pieces like the <code>[]</code> square braces, length of the array, values, etc. In golang, we cannot resize the length once it is initialized.</p>
<p>To create a basic array in golang, we can use the following code:</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
    var languages[4]string
    languages[0] = &quot;Python&quot;
    fmt.Println(languages)
}
</code></pre>
<pre><code>$ go run basic.go
[Python   ]
</code></pre>
<p>Here, we have initialized a string array of size 4 and not initialized any values of the elements in the declaration. We later set the value of the 0th index or the first element in the array to a string and still rest of the elements are by default set to empty strings <code>&quot;&quot;</code>. The point ot be noted here, is that the size of the array cannot be changed later nor you can leave the size empty while declaring an array in Golang.</p>
<p>We can even initialize elements at the time of declaring the array as follows:</p>
<pre><code class="language-go">lang_array := [4]string {&quot;Python&quot;, &quot;Go&quot;, &quot;Javascript&quot;, &quot;C++&quot;}
fmt.Println(lang_array)
</code></pre>
<pre><code>$ go run basic.go
[Python Go Javascript C++]
</code></pre>
<p>So, we use the walrus operator <code>:=</code> to initialize an array with the values at the time of declaration.</p>
<h3>Letting Compiler makeout the array length</h3>
<p>We can even let the compiler decide the length of the array, using the <code>...</code> syntax inside the brackets. This is limited by using the array literal that is by initializing values in the <code>{}</code> braces. So, all the elements need to be declared in the array.</p>
<pre><code class="language-go">cart := [...]string {&quot;Bag&quot;, &quot;Shirt&quot;, &quot;Watch&quot;, &quot;Book&quot;}
fmt.Println(cart)
</code></pre>
<pre><code>$ go run basic.go
[Bag Shirt Watch Book]
</code></pre>
<h2>Access and Modify Elements</h2>
<p>To access an element in the array, we can use the index of that element which starts from 0 as usual in programming.</p>
<pre><code class="language-go">marks := [6]int {85, 89, 75, 93, 98, 60}
fmt.Println(marks[1])
fmt.Println(marks[5])
fmt.Println(marks[3])
</code></pre>
<pre><code>$ go run basic.go
89
60
93
</code></pre>
<p>We can now, access the element at a particular index in the array. Now, we will see how to modify or edit the elements which are already initialized.</p>
<pre><code class="language-go">name := [5]byte {'f','u','z','z','y'}
fmt.Printf(&quot;%s
&quot;,name)
name[0] = 'b'
name[4] = 'z'
fmt.Printf(&quot;%s
&quot;,name)
</code></pre>
<pre><code>$ go run basic.go
fuzzy
buzzz
</code></pre>
<p>By accessing the index of the element we can set a appropriate value to the element in the array and thus we have modified the contents of the array.</p>
<h2>Find Length of Array</h2>
<p>To find the length of the Array, we have the <code>len</code> function. The <a href="https://pkg.go.dev/builtin#len">len</a> function takes in the array as the parameter and returns the size of the array(int).</p>
<pre><code class="language-go">code := [7]rune {'#', '5', 'g', 't', 'm', 'y', '6'}
fmt.Println(&quot;The length of the array is :&quot;, len(code))
</code></pre>
<pre><code>$ go run basic.go
The length of the array is : 7
</code></pre>
<p>In the previous few section, we talked about letting the compiler make out the length of the array while declaring and initializing the array, we can use the len function to calculate the length of the array for further computation</p>
<pre><code class="language-go">cart := [...]string {&quot;Bag&quot;, &quot;Shirt&quot;, &quot;Watch&quot;, &quot;Book&quot;}
fmt.Printf(&quot;There are %d items in your cart
&quot;, len(cart))
</code></pre>
<pre><code>$ go run basic.go
There are 4 items in your cart
</code></pre>
<p>We can now get the length of the arrays even with <code>[...]</code> syntax using the len function.</p>
<h2>Iterate over an Array</h2>
<p>We can move on to the most important aspect when it comes to arrays i.e. to iterate over each element. We can use various types of for loops like the three statement for loops, range based loop or while loop.</p>
<h3>Three statement for loop</h3>
<p>We can use the three statement for loop, the initialization statement as to <code>0</code>, condition to be the counter (i) should be less than the length of the array by using the <code>len</code> function and increment each time by one.</p>
<pre><code class="language-go">code := [7]rune {'#', '5', 'g', 't', 'm', 'y', '6'}

for i := 0; i&lt;len(code); i++{
    fmt.Printf(&quot;%c
&quot;,code[i])
}
</code></pre>
<pre><code>$ go run basic.go
#
5
g
t
m
y
6
</code></pre>
<p>Thus, we can iterate over the array with three statement for loop in golang.</p>
<h3>Using range-based loop</h3>
<p>We can use the <code>range</code> keyword to iterate over the arrays in golang. The range keyword is used to iterate over the array by taking two variables i.e. the iterator and the copy of the element in the iterator. We don't have any use of the iterator so we say it as <code>_</code>, otherwise it gives a warning/error of not using declared variables. So, we only require the copy of the element in this case, so sayit as <code>s</code> or any other name you like.</p>
<pre><code class="language-go">cart := [...]string {&quot;Bag&quot;, &quot;Shirt&quot;, &quot;Watch&quot;, &quot;Book&quot;}

for _, s := range cart{
    fmt.Println(s)
}
</code></pre>
<pre><code>$ go run basic.go
Bag
Shirt
Watch
Book
</code></pre>
<p>Thus, using the range based for loops we were able to iterate over the array for each element without needing any check condition and incrementation of the counter/iterator.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So, from this part of the series, we were able to understand the baiscs of arrays in golang. We covered from declaration of arrays to iteration.
Thank you for reading. If you have any questions or feedback, please let me know in the comments or on social handles. Happy Coding :)</p>
