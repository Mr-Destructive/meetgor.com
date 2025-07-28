{
  "title": "Golang: Maps",
  "status": "published",
  "slug": "golang-maps",
  "date": "2025-04-05T12:33:43Z"
}

<h2>Introduction</h2>
<p>In the seventh part of the series, we will be covering Maps. We have covered some basic data structures like arrays and slices, and now we can move into maps or hash tables. Maps allow us to store key-value pairs of a particular type. In this part of the series, we will be covering the basics of Maps in Golang like declaration, iteration, and Creating, updating, and deleting keys from the map.</p>
<h2>Maps in Golang</h2>
<p>Maps in golang are data structures that provide a way to store key-value pairs. It is also known as a hash table. Maps allow us to create unique keys which are associated with a value. It can be used to create a data structure that can have an item that is associated with a particular value, for example, the basic example of the map can be a frequency table of a list of numbers. We can store the frequency of each element occurring in the list. Let's say we have a list of numbers as <code>[3, 5, 9, 4, 9, 5, 5]</code>, we can create a map of the frequency of these elements as <code>[3:1, 5:3, 4:1, 9:2]</code>. Here, we have stored the information in the form of <code>key-value</code> pairs as a frequency. So, <code>3</code> has occurred one time, <code>5</code> 3 times, and so on.</p>
<p>Maps are not stored in order of the numbers they are unordered so we need to manually sort them in the order we want.</p>
<h2>Declaring Maps</h2>
<p>We can declare maps by defining the type of mapping like the two types we are mapping. We can map any type with any other, like a character with an integer, an integer with an integer as we saw earlier, etc. We have several ways to decalre maps in golang, like using map literal, make function, new function, and a few others. We'll look into each of them in a brief.</p>
<h3>Simple map literal</h3>
<p>As we saw in the array and slices, we used the slice literals to declare and initialize an array or a slice. Similarly, we can use the map literal to create a map in golang. Here, we use the <code>map</code> keyword followed by the two types of data we are going to map with.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {

	char_freq := map[string]int{
		&quot;M&quot;: 1,
		&quot;e&quot;: 2,
		&quot;t&quot;: 1,
	}
	fmt.Println(char_freq)
}
</code></pre>
<pre><code>$ go run map.go
map[M:1 e:2 t:1]
</code></pre>
<p>We have used the map keyword to initialize a map with a <code>string</code> with <code>int</code>. The first data type is declared inside the square brackets<code>[]</code> and the second data type outside the square brackets. We use the <code>{}</code> to define the map values. We can even leave the <code>{}</code> empty.</p>
<pre><code>char_freq := map[string]int{}
</code></pre>
<p>We initialize the values of the map by specifying the data for that data type in this example a string <code>&quot;&quot;</code> followed by a colon <code>:</code> and finally the value of the second pair data. Each value is separated by a comma(<code>,</code>).</p>
<h3>Using make function</h3>
<p>We can even use the <a href="https://pkg.go.dev/builtin#make">make</a> function to create a map in golang. The make function is used for allocating memory. The make function allocates memory which might be enough for the initial values provided. It allocates more memory as the map grows in size. We use the make function by providing the <code>map</code> keyword along with the data types of the key values pairs to be mapped. Optionally we can provide the capacity as we provided in the slice declaration. It basically doubles once it reaches the limit and is re-allocated.</p>
<pre><code class="language-go">marks := make(map[int]int)
marks[65] = 8
marks[95] = 3
marks[80] = 5
fmt.Println(marks)
</code></pre>
<pre><code>$ go run map.go
map[65:8 80:5 95:3]
</code></pre>
<p>We have used the <code>make</code> function for declaring the map, the initial size is around 7 if not mentioned. After it hits 7, the capacity is mostly doubled and increased as per the modifications.</p>
<h3>Using the new function</h3>
<p>We can even use the <a href="https://pkg.go.dev/builtin#new">new</a> function(a bit hacky) to crated a map in golang. The new function basically is used to allocate memory but is not the same as the <code>make</code> function, it returns the memory address to an allocated pointer. So, we can set the value of the returned function call of the new function with a pointer variable. A pointer in golang is simply a reference to a memory address, we'll dive into pointers in a different section. After the pointer is assigned a memory address, we can refer to the address of that pointer and thus access the original value which is the map itself.</p>
<pre><code class="language-go">name := new(map[byte]int)
*name = map[byte]int{}
name_map := *name

name_map['m'] = 1
name_map['e'] = 2
name_map['t'] = 1

fmt.Println(name_map)

for k, _ := range name_map {
    fmt.Printf(&quot;%c
&quot;, k)
}
</code></pre>
<pre><code>$ go run map.go
map[101:2 109:1 116:1]
m
e
t
</code></pre>
<p>So, we can see we created the map with the new function and stored the address into a pointer, later we initialized the empty map and stored the initial reference in the same pointer address. Then, we can finally store the map in another variable so that we can use it as a normal map. So, this is how we declare the map using the new function.</p>
<h2>Access Keys and Values in Maps</h2>
<p>We can access the values by simply accessing them with the keys. Using the square bracket and the key literal into the braces, we get the value associated with that key. For example, the map <code>[&quot;M&quot;: 1, &quot;E&quot;: 2, &quot;T&quot;:1]</code>, we can use the <code>map_name[&quot;E&quot;]</code> which will get the value as <code>3</code>.</p>
<h3>Length of Map</h3>
<p>The length of the map can be accessed using the len function, the len function returns the number of key-value pairs in the map.</p>
<pre><code class="language-go">char_freq := map[string]int{
    &quot;M&quot;: 1,
    &quot;e&quot;: 2,
    &quot;t&quot;: 1,
}
fmt.Println(char_freq)
fmt.Println(len(char_freq))
</code></pre>
<pre><code>$ go run map.go
map[M:1 e:2 t:1]
3
</code></pre>
<h2>Check for existing Keys in Map</h2>
<p>We can check if a key exists in the map by using the comma-ok syntax. The key can be accessed using the first variable and if the key doesn't exist, the second variable is set to false. So, we can verify the existence of a key in the map using the two-variable approach.</p>
<pre><code class="language-go">name_map := map[byte]int{
    'm': 1,
    'e': 2,
    't': 1,
}
var key byte = 't'
value, exist := name_map[key]
if exist == true {
    fmt.Printf(&quot;The key %c exist and has value %d
&quot;, key, value)
} else {
    fmt.Printf(&quot;The key %c does not exist.
&quot;, key)
}
</code></pre>
<pre><code>$ go run map.go
The key t exist and has value 1
</code></pre>
<p>So, we can see the exist value is true if the key exists and false if it doesn't. So, we can then verify if a particular key exists in a map or not.</p>
<h2>Adding and Modifying Keys/Values in Maps</h2>
<p>We can add a key-value pair in a map by just using the key as we did in the initialization process. We simply pass the key in the square braces <code>[]</code> and assign it a value appropriate to the data type used in the map.</p>
<pre><code class="language-go">cart_list := map[string]int{
    &quot;shirt&quot;: 2,
    &quot;mug&quot;: 4,
    &quot;shoes&quot;: 3,
}

fmt.Println(cart_list)

cart_list[&quot;jeans&quot;] = 1
cart_list[&quot;mug&quot;] = 3
fmt.Println(cart_list)
</code></pre>
<pre><code>$ go run map.go
map[mug:4 shirt:2 shoes:3]
map[jeans:1 mug:3 shirt:2 shoes:3]
</code></pre>
<p>We can access the keys in the map by just using the key as it is and altering the value it holds, the same thing applies to the addition of the key-value pairs, we can use the key and assign the value associated with it.</p>
<h2>Delete Keys in Maps</h2>
<p>We can delete the key-value pairs in the map, using the <code>delete</code> function. We pass in the <code>key</code> and the map to delete the key-value pair from the map.</p>
<pre><code class="language-go">cart_list := map[string]int{
    &quot;shirt&quot;: 2,
    &quot;mug&quot;:   4,
    &quot;shoes&quot;: 3,
}
fmt.Println(cart_list)

cart_list[&quot;jeans&quot;] = 1
cart_list[&quot;mug&quot;] = 3
delete(cart_list, &quot;shoes&quot;)

fmt.Println(cart_list)
</code></pre>
<pre><code>$ go run map.go
map[mug:4 shirt:2 shoes:3]
map[jeans:1 mug:3 shirt:2]
</code></pre>
<p>So, we can see the key-value pair was deleted from the map.</p>
<h2>Iterate over a Map</h2>
<p>We can iterate over a map similar to the range keyword iteration for slices and arrays, but the exception here, is that we use the key, value instead of the index, copy of an element in the map as the range.</p>
<pre><code class="language-go">is_prime := map[int]bool{
    7:  true,
    9:  false,
    13: true,
    15: false,
    16: false,
}

for key, value := range is_prime {
    fmt.printf(&quot;%d -&gt; %t
&quot;, key, value)
}
</code></pre>
<pre><code>$ go run map.go
9 -&gt; false
13 -&gt; true
15 -&gt; false
16 -&gt; false
7 -&gt; true
</code></pre>
<p>So, we can observe that we can access the keys and values in the map using the range keyword for iterating over the map. Inside the for loop, we can refer to the assigned values present in the map.</p>
<h3>Use only key or value while iterating</h3>
<p>If we don't use either of the variables like <code>key</code> or  <code>value</code>, the compiler might give us the unused variable error, so we have an alternative to use don't care variables namely the <code>_</code> underscore character.</p>
<pre><code class="language-go">is_prime := map[int]bool{
    7:  true,
    9:  false,
    13: true,
    15: false,
    16: false,
}

for key, _ := range is_prime {
    fmt.Printf(&quot;Key : %d
&quot;, key)
}

for _, value := range is_prime {
    fmt.Printf(&quot;Value: %t
&quot;, value)
}
</code></pre>
<pre><code>$ go run map.go
Key : 7
Key : 9
Key : 13
Key : 15
Key : 16
Value: true
Value: false
Value: true
Value: false
Value: false
</code></pre>
<p>So, we use the <code>_</code> to ignore the usage of the variable in the loop, if we are not sure of using any variable, we can ignore it completely with the underscore operator and thus prevent any compilation errors/warnings. So, here if we want to only access keys, we use <code>key, _</code> in order to fetch only keys and silence the values in the map. If we want to access only the values, we can use <code>_, value</code> so as to get all the values from the map. The variable name <code>key</code> or <code>value</code> can be anything but make sure to use those only inside the loop.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So, from this part of the series, we were able to understand the basics of maps in golang. We covered some basics stuff including the declaration, initialization, and iteration. Maps are quite simple but important for creating interesting applications.</p>
<p>Thank you for reading. If you have any questions or feedback, please let me know in the comments or on social handles. Happy Coding :)</p>
