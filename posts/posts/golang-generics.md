{
  "title": "Golang: Generics",
  "status": "published",
  "slug": "golang-generics",
  "date": "2025-04-05T12:33:31Z"
}

<h2>Introduction</h2>
<p>In the 29th post of the series, we will be looking into generics in Golang. Generics were added in Golang version 1.18, so they are quite new in the world of Golang but the concept is quite old in other programming languages.</p>
<p>Generics provide a powerful toolset for writing more expressive and concise code that can handle a wide range of data types. With generics, we can write reusable algorithms, data structures, and functions that work seamlessly with various types, without sacrificing type safety.</p>
<p>We will learn how to create generic functions and work with generic types. Additionally, we will cover type constraints and interfaces, which allow us to specify requirements for the types used with our generics.</p>
<h2>Generic Type in Functions</h2>
<p>We can define a generic type with the keyword <code>any</code> that is going to replace the type <code>T</code> i.e. any data type with the inferred type at compilation. This makes the code reusable to any relevant data type to be used for that operation/task.</p>
<p>We can provide the type <code>any</code> after the name of the function/struct to make it generic like <code>func Name[T any](x T)</code>. Here, the Name is a function that takes in a generic type <code>T</code> it could be any type and the variable is named as <code>x</code> that could be used inside the function.</p>
<p>We could also make the function take specific types instead of <code>any</code> but we will eventually move into that. However, for now, let's ease the process of learning and then move on to the optimizations and adding constraints.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
)

func Print[T any](stuff T) {
	fmt.Println(stuff)
}

func main() {
	Print(&quot;hello&quot;)
	Print(123)
	Print(3.148)
}
</code></pre>
<p><a href="https://go.dev/play/p/t-ODmkHu5BJ">GO Playground Link</a></p>
<pre><code class="language-bash">$ go run main.go
hello
123
3.148
</code></pre>
<p>The above is the simplest example that could be used to demonstrate a generic function. The function <code>Print</code> takes a parameter <code>stuff</code> of a generic type denoted by a type parameter <code>T</code>. The type parameter <code>T</code> serves as a placeholder that represents a specific type determined at compile time when the function is invoked.</p>
<p>This means, if in my code I do not call the function with the type <code>[]int</code> it won't have the function with the signature as <code>Print(stuff []int)</code> rather only the types which the function is called with are inferred and written with that specific types.</p>
<h2>Creating a Generic Slice</h2>
<p>We can even create a slice with a generic type and allow any valid processing on the elements or the slice as a whole.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
)

func ForEach[T any](arr []T, f func(T)) {
	for _, v := range arr {
		f(v)
	}
}

func main() {

	strSlice := []string{&quot;b&quot;, &quot;e&quot;, &quot;a&quot;}
	ForEach(strSlice, func(v string) {
		fmt.Println(v)
	})

	slice := []int{10, 25, 33, 42, 50}
	var evenSlice []int
	ForEach(slice, func(v int) {
		isEven := v%2 == 0
		if isEven {
			evenSlice = append(evenSlice, v)
		}
	})
	fmt.Println(evenSlice)

}
</code></pre>
<pre><code class="language-bash">$ go run main.go

b
e
a
[10 42 50]
</code></pre>
<p><a href="https://go.dev/play/p/tUwgbxnLc-1">Go Playground Link</a></p>
<p>The <code>ForEach</code> is a generic function that iterates over a slice of any type and calls a function on each element. Let's break it down:</p>
<ul>
<li>
<p><code>ForEach[T any]</code> declares this as a generic function that works on a slice of any type <code>T</code>.</p>
</li>
<li>
<p><code>arr []T</code> is the slice of elements we want to iterate over. It can be a slice of any type, ints, strings, <code>T</code> in general, etc. So it is a generic slice parameter.</p>
</li>
<li>
<p><code>f func(T)</code> is the function that will be called on each element. It takes a single parameter of type <code>T</code> which will be each element. So, this is a function parameter with a generic type as its parameter.</p>
</li>
</ul>
<p>In the body, we range over the slice arr:</p>
<pre><code class="language-go">for _, v := range arr {

}
</code></pre>
<p>On each iteration, <code>v</code> will be the next element. The underscore ignores the index. We call the provided function <code>f</code>, passing the element <code>v</code>: <code>f(v)</code></p>
<p>So in summary:</p>
<ul>
<li>
<p>It allows iterating over a slice of any type</p>
</li>
<li>
<p>This avoids having to know the specific slice type in the loop</p>
</li>
<li>
<p>The provided function <code>f</code> is called on each element</p>
</li>
<li>
<p>So it provides a reusable abstraction for processing slices generically.</p>
</li>
</ul>
<p>Now, we will discuss the example used in the main function. First, we create a slice of strings as <code>strSlice := []string{&quot;b&quot;, &quot;e&quot;, &quot;a&quot;}</code>. Then we call the generic <code>ForEach</code> function, passing the string slice and a function to handle each element.</p>
<pre><code class="language-bash">ForEach(strSlice, func(v string) {
  fmt.Println(v) 
})
</code></pre>
<p>Here, the <code>func(v string)</code> is the invocation of the <code>ForEach</code> function with the typed string and the variable name as v. The <code>v</code> is the element of the slice, so inside the function body(this is an anonymous function), we call the <code>fmt.Println(v)</code>, which will print each string in the slice.</p>
<pre><code class="language-go">slice := []int{10, 25, 33, 42, 50}
var evenSlice []int
ForEach(slice, func(v int) {
    isEven := v%2 == 0
    if isEven {
        evenSlice = append(evenSlice, v)
    }
})
fmt.Println(evenSlice)
</code></pre>
<p>In the next example, we create a new slice of int as <code>slice := []int{10, 25, 33, 42, 50}</code>. Then we call the generic <code>ForEach</code> function again, passing the slice and a function to handle each element just as before, however, we are processing some things and then appending to an external slice.</p>
<p>First, the <code>slice := []int{10, 25, 33, 42, 50}</code> is created with some numbers, we also initialize another slice as <code>evenSlice := []int{}</code> which is empty. Then we call the generic <code>ForEach</code> function again, passing the slice and a function to handle each <a href="http://element.Here">element.Here</a>, the ForEach is called with the <code>slice</code> slice and not the <code>evenSlice</code> slice, so we are going to access each element in the <code>slice</code> array. We first create a <code>isEven</code> boolean that checks if the element is even or odd by <code>v%2 == 0</code>. This expression will result in <code>true</code> if <code>v</code> is even and <code>false</code> otherwise. So, only if the <code>isEven</code> is true, we append the element <code>v</code> into the <code>evenSlice</code> slice.</p>
<p>So, that's how generic slices can be handy for doing type-specific processing without writing functions for those individual types. This avoids needing to write type-specific functions for each slice type.</p>
<p>NOTE: Make sure to only use generic functions with generic slice types with the appropriate and valid conditions and use it only when it looks legible to do so.</p>
<h2>Creating a Generic Map</h2>
<p>We can also create a generic map with the generic type of <code>map[K]T</code> where <code>K</code> is a generic type and <code>T</code> is the type of the key.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
)

func GetValue[K comparable, V any](m map[K]V, key K, defaultVal V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultVal
}

func main() {

	serverStats := map[string]int{
		&quot;port&quot;:      8000,
		&quot;pings&quot;:     47,
		&quot;status&quot;:    1,
		&quot;endpoints&quot;: 13,
	}
	v := GetValue(serverStats, &quot;status&quot;, -1)
	fmt.Println(v)
	v = GetValue(serverStats, &quot;cpu&quot;, 4)
	fmt.Println(v)

}
</code></pre>
<pre><code class="language-bash">$ go run main.go
1
4
</code></pre>
<p><a href="https://go.dev/play/p/ludlh6UwKaD">Go Playground Link</a></p>
<p>GetValue is a generic function that takes three type parameters: The map itself, the key to find the value for, and a default value if the key doesn't exist.</p>
<p>The <code>m</code> is a map with keys of type K and values of type V, the key is of type K, and defaultVal is of type V. So, we have two generics here, as the key and value need not be of the same type, hence we have distinct generics here. K has added a constraint of <code>comparable</code> and <code>V</code> as <code>any</code> type. The type constraint comparable restricts K to be a comparable type, and the type constraint <code>any</code> allows V to be any type.</p>
<ul>
<li>
<p>Inside the function, we use the ok variable to check if the given <code>key</code> exists in the map <code>m</code>.</p>
</li>
<li>
<p>If the key is present in the map (ok is true), we retrieve the corresponding value from the map and return it as <code>m[key]</code> which is stored in the variable <code>v</code>.</p>
</li>
<li>
<p>If the key is not present in the map (ok is false), we return the provided <code>defaultVal</code>.</p>
</li>
</ul>
<p>So, this is how we can use any type of map to retrieve the value of a key, the data type of key and value could be anything. It allows us to retrieve a value from a map irrespective of the pair type in the map.</p>
<p>NOTE: The type of <code>defaultVal</code> and the type of <code>v</code> should be the same since the map will need to have the value for the given key as the same type as defined in the map type(here <code>map[string]int</code> so <code>v</code> is <code>int</code> and so should be the <code>defaultVal</code>).</p>
<p>Moving into the main function, we create a map of <code>[string]int</code> i.e. the key is of type <code>string</code> and the value of type <code>int</code>. The map <code>serverStats</code> has a few keys like <code>port</code>, <code>pings</code>, <code>endpoints</code>, and <code>status</code>. We call the <code>GetValue</code> method on the map <code>serverStats</code> with the key <code>status</code> and provide a default value of <code>-1</code>. The function will readily return <code>1</code> since the key is present in the provided map. However, if we try to access the key <code>cpu</code>, the key is not present and the value is returned as the <code>defaultVal</code> which we passed as <code>4</code>.</p>
<p>So, this was a simple generic getter method on a map. We can get a value of a key in a map of any pair and even provide a default value if doesn't exist. However, it won't add it to the map, we will just return the value from the function that's it. We have to see that returned default value manually.</p>
<p>We can make another function to get or set the value of a key in a map. The function will take in a reference to the map rather than a copy of the map, we can then use that reference to set the key with the provided default value.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
)

func GetOrSetValue[K comparable, V any](m *map[K]V, key K, defaultVal V) V {
    // reference the original map
	ref := *m
	if v, ok := ref[key]; ok {
		return v
	} else {
        //mutate the original map
        ref[key] = defaultVal

        return defaultVal
    }
}

func main() {
    serverStats := map[string]int{
        &quot;port&quot;:      8000,
        &quot;pings&quot;:     47,
        &quot;status&quot;:    1,
        &quot;endpoints&quot;: 13,
    }
    fmt.Println(serverStats)
    v := GetOrSetValue(&amp;serverStats, &quot;cpu&quot;, 4)
    fmt.Println(v)
    fmt.Println(serverStats)
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

map[endpoints:13 pings:47 port:8000 status:1]
4
map[cpu:4 endpoints:13 pings:47 port:8000 status:1]
</code></pre>
<p><a href="https://go.dev/play/p/fYtjFQRaPCb">Go Playground Link</a></p>
<p>In the above code, we take a reference of the map as <code>*map[K]V</code>, this will give access to the actual place(memory address where the map is located so we could mutate/change it). The rest of the parameters are kept as is, the key will be taken as it was before, and so will the <code>defaultVal</code>. The only difference is that we will set the <code>key</code> doesn't exist, we set the <code>ref[key]</code> to the <code>defaultVal</code> and return the <code>defaultVal</code>.</p>
<p>For example, the <code>cpu</code> key is not present in the initial map <code>serverStats</code> so, when we call the <code>GetOrSetValue</code> with the reference to the map <code>serverStats</code>, key as <code>cpu</code> and the default value as <code>4</code>, the function returns <code>4</code> and the map is mutated with the key <code>cpu</code> with value <code>4</code>.</p>
<p>The takeaway is you can even use references to access the original data and mutate it based on your needs.</p>
<h2>Generic Type in Struct</h2>
<p>We can define custom structs with generic type as the field values. The name of the struct is followed by the <code>[T any]</code> which is the type parameter to be used in the struct fields, again this type could have multiple types(since a struct can have many fields), not necessary want a single type parameter, you could have multiple type parameters bunched up just like we saw in the map example.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
)

type Stack[T any] struct {
	items []T
}

func NewStack[T any]() *Stack[T] {
	return &amp;Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		panic(&quot;Stack is empty&quot;)
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

func main() {
	intStack := NewStack[int]()
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)
	fmt.Println(&quot;Integer Stack&quot;)
	fmt.Println(intStack)
	intStack.Pop()
	intStack.Pop()
	fmt.Println(intStack)

	// without the NewStack method
	strStack := Stack[string]{}
	strStack.Push(&quot;c&quot;)
	strStack.Push(&quot;python&quot;)
	strStack.Push(&quot;mojo&quot;)
	fmt.Println(&quot;String Stack:&quot;)
	fmt.Println(strStack)
	strStack.Pop()
	fmt.Println(strStack)
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

Integer Stack
&amp;{[10 20 30]}
&amp;{[10]}

String Stack:
{[c python mojo]}
{[c python]}
</code></pre>
<p><a href="https://go.dev/play/p/4t_P2mKtTZN">Go Playground Link</a></p>
<p>In this example, we have used the <code>Stack</code> example for doing a basic <code>Push</code> and <code>Pop</code> operation on the elements. Here the type of the underlying stack elements could be anything, hence the type parameter is defined for the <code>items</code> which is a list/slice of the type <code>T</code> as <code>[]T</code>. We have to specify the type before initializing the</p>
<p>We have created the <code>NewStack</code> method, it is not needed, it could be just used as <code>Stack[int]{}</code> to initialize a empty stack with int type(here <code>int</code> could be any other type you wish). I have just created it so that it shows the abstraction that could be built upon in real applications. The <code>NewStack</code> simply replaces the <code>T</code> with the provided <code>type</code> in the initialization.</p>
<p>The <code>Push</code> method is associated with the <code>Stack</code> struct, as we refer to the <code>*Stack[T]</code> indicating a reference to the Stack object with the type <code>T</code>. The method takes in the parameter <code>T</code> which would be the element to be added to the <code>Stack</code>. Since the function is tied to the Stack struct, we can simply mutate the underlying <code>items</code> by appending the provided value <code>item</code> in the parameter using <code>s.items = append(s.items, item)</code>. This appends the <code>item</code> to the underlying <code>items</code> list in the <code>Stack</code> object <code>s</code></p>
<p>Similarly, we can create <code>Pop</code> method as well, which will first check if the <code>Stack</code> is not empty(i.e. the s.items slice has a length greater than 0), then get the index of the last element using <code>len(s.items) - 1</code> and then slice the <code>items</code> at index <code>[:last_index]</code>. This will basically get you all the elements except the last. Before we remove the element from the slice, we also store the item in <code>item</code> variable and return it from the method.</p>
<p>You could see the case of generics here, you could build complex data structures without creating a separate implementation for each type.</p>
<h2>Adding Constraints to Generics</h2>
<p>We can add constraints to the generics to restrict the type of the generic parameter. For example, we can add a constraint for the generic type to be a slice of a specific type or we have seen in the map example the <code>comparable</code> constraint.</p>
<p>The <code>comparable</code> constraint is an interface that allows two instances of the same type to be compared i.e. support comparison operators like ==, &lt;, &gt;, !=, &gt;=, &lt;=, etc. It could be any numeric type like <code>int</code>, <code>float</code>, <code>uint</code> and variants, booleans, time duration, and any other struct that implements the <code>comparable</code> interface.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
)

func FindIndex[T comparable](arr []T, value T) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

func main() {

	strSlice := []string{&quot;m&quot;, &quot;e&quot;, &quot;e&quot;, &quot;t&quot;}
	fmt.Println(FindIndex(strSlice, &quot;e&quot;))
	fmt.Println(FindIndex(strSlice, &quot;t&quot;))
	fmt.Println(FindIndex(strSlice, &quot;a&quot;))

	intSlice := []int{10, 25, 33, 42, 50}
	fmt.Println(FindIndex(intSlice, 33))
	fmt.Println(FindIndex(intSlice, 90))

}
</code></pre>
<pre><code class="language-bash">$ go run main.go
1
3
-1

2
-1
</code></pre>
<p><a href="https://go.dev/play/p/fv9gzb8K4d7">Go Playground Link</a></p>
<p>In the above example, we have created the function <code>FindIndex</code> that takes in a generic slice, the type parameter <code>[T comparable]</code> indicates that the type used for calling this method needs to have a type that implements the comparable interface (for the elements of the slice). The method takes in two parameters, one the slice as <code>[]T</code> and the other the value to find the index for as type <code>T</code>. The method returns a type <code>int</code> since the index of the slice has to be an integer value.</p>
<p>Inside the function body, we simply iterate over the slice <code>arr</code> and check if the element is equal to the provided value. If the element exists, we return that index, else we return <code>-1</code></p>
<p>As we can see we have run a couple of slices with the function <code>FindIndex</code> with types <code>int</code> and <code>string</code>. The method returns an index value if the element exists, else it returns <code>-1</code>. The <code>comparable</code> is a built-in type constraint. We could even define custom constraints as interfaces that implement the types of the particular type(s).</p>
<p>Also, we could define custom constraints like numeric only, string only, etc.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
)

type numeric interface {
	uint | uint8 | uint16 | uint32 | uint64 |
		int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Sum[T numeric](nums []T) T {
	var s T
	for _, n := range nums {
		s += n
	}
	return s
}

func main() {

	intSlice := []int{10, 20, 30, 40, 50}
	fmt.Println(Sum(intSlice))

	floatSlice := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(Sum(floatSlice))
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

150
16.5
</code></pre>
<p><a href="https://go.dev/play/p/_1eGl58qQ-2">Go Playground Link</a></p>
<p>In the above example, we have created the <code>numeric</code> constraints that allow the type <code>int</code>, <code>float</code> and their variants to be allowed in the numeric generic type. The function <code>Sum</code> is a generic function with the constraint of <code>numeric</code> type parameter. The method takes in the parameter as type <code>[]T</code> and returns the type as <code>T</code>. The method will simply iterate over the slice and return the sum of its elements.</p>
<p>This will allow any numeric type which can be added and the sum can be obtained, so if we try to call the method with other types like <code>string</code> or <code>maps</code>, it won't work, and give an error:</p>
<pre><code class="language-bash">$ go run constraints.go

# command-line-arguments                                                                                                               
scripts/generics/constraints.go:46:20: 
string does not satisfy numeric (string missing in uint | uint8 | uint16 | uint32 | uint64 | int
 | int8 | int16 | int32 | int64 | float32 | float64)

shell returned 1
</code></pre>
<p>So, we can use the constraint to restrict the type of the generic type parameter which will allow us to restrict the usage and avoid any unsafe type to be used in the generic function.</p>
<p>Also, if we have a custom type with the base types, we need to use <code>~</code> before the type to accept it into the generic constraint. This will allow any approximate type to be allowed in the constraint. Let's say we are implementing a custom string type, for that to work with a constraint, it won't be satisfied in the constraint since its type is <code>CustomString</code> and not <code>string</code>. So to avoid this we use <code>~string</code> that would approximate the type and allow abstracted base types.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
)

type string2 string

type strings interface {
	~string
}

func PrintEach[T strings](arr T) {
	for _, v := range arr {
		fmt.Printf(&quot;%c
&quot;, v)
	}
}

func main() {
	var s string2
	s = &quot;hello&quot;
	PrintEach(s)

}
</code></pre>
<pre><code class="language-bash">$ go run main.go

h
e
l
l
o
</code></pre>
<p><a href="https://go.dev/play/p/N-66A9C94ps">Go Playground Link</a></p>
<p>In the above example, we have used the type approximations in the type constraint <code>strings</code>, it allows any string type, whether a base <code>string</code> type or an abstract <code>string</code> type. If you try to remove the <code>~</code> in <code>~string</code>, it will result in the error that <code>string2 does not satisfy strings</code> interface. So, by adding <code>~</code> to the <code>string</code> type the abstract type <code>string2</code> can be satisfied in the generic constraint.</p>
<p>That's it from the 29th part of the series, all the source code for the examples are linked in the GitHub on the <a href="https://github.com/Mr-Destructive/100-days-of-golang/tree/main/scripts/generics">100 days of Golang</a> repository.</p>
<h3>References</h3>
<ul>
<li><a href="https://bitfieldconsulting.com/golang/generics">Generics in Go</a></li>
<li><a href="https://blog.logrocket.com/understanding-generics-go-1-18/">Understanding generics in Go</a></li>
</ul>
<h2>Conclusion</h2>
<p>From this section of the series, we have covered the basics of generics in Golang. By using generics in functions, slices, maps, and structs, and adding constraints to them the fundamental usage of generics was covered.</p>
<p>If you have any questions, feedback, or suggestions, please drop them in the comments/social handles or discuss them below. Thank you so much for reading. Happy Coding :)</p>
