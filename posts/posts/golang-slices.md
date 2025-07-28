{
  "title": "Golang: Slices",
  "status": "published",
  "slug": "golang-slices",
  "date": "2025-04-05T12:33:43Z"
}

<h2>Introduction</h2>
<p>In the sixth part of the series, we will be covering slices. Slices are almost like arrays but have a lot of advantages over them, including flexibility and control over them. We can adjust the size and capacity of the data which we will store at a place using slices. We will be covering basic declaration, initialization, capacity in slices, iteration, and accessing the elements of the slices.</p>
<h2>Slices in Golang</h2>
<p>Slices are Arrays but can provide more control and are more flexible than arrays. In slices, we can increase the size of the array/list of elements whenever required. We can even have a capacity for slices i.e. the maximum length we wish to grow the initial slice.</p>
<p>Though slices are dynamic, it has a few disadvantages like compile safety, access time, comparability, etc. Everything has its pros and cons, you have to decide on the right data structure as per your problem statement and requirements.</p>
<h2>Declaring Slices</h2>
<p>There are a couple of different ways in which we can declare a slice that might be an uninitialized or initialized slice. Some of the standard ways include using the <code>make</code> function and the normal array-like declaration though there are other methods as well including using the <code>new</code> function. We'll explore some of the most easiest and handy ways to declare and initialize slices.</p>
<h3>Using array-like declaration</h3>
<p>Slices can be declared quite similar to arrays but we don't enter the initial size(length). As discussed in the array part, we can use the var keyword and the square brackets. Though you don't have to enter the length inside the <code>[]</code> brackets, the type of the slice needs to enter.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
    var marks [] int
    fmt.Println(marks)
}
</code></pre>
<pre><code>$ go run slices.go
[]
</code></pre>
<p>So, initially, it is empty without an element as we have not initialized any elements. If you enter the length in the <code>[]</code> brackets, it would be an array and not a slice.</p>
<p>We can also initialize the elements in the slice with the slice literal as we did with arrays using an array literal.</p>
<pre><code class="language-go">frameworks = []string{&quot;Django&quot;, &quot;Laravel&quot;, &quot;Flask&quot;, &quot;Rails&quot;}
fmt.Println(frameworks)
</code></pre>
<pre><code>$ go run slices.go
[Django Laravel Flask Rails]
</code></pre>
<h3>Using make function</h3>
<p>We can even use the <a href="https://pkg.go.dev/builtin#make">make</a> function to create a slice in golang. The make function basically allows us to create a slice by providing the length as well as the capacity. Let's clear the concept of Capacity and Slice first.</p>
<h4>Capacity in Slice</h4>
<p>Capacity in slices is the initial value provided during the declaration of a slice. It basically holds the capacity of the slice to grow beyond the length. OK, let's take an example, If you declare a slice using make with length 3 and capacity 5, you initially can access 3 elements but memory has been allocated for 5 elements, if your slice exceeds this capacity it will double its original capacity.</p>
<h4>Obtain Length and Capacity of Slice</h4>
<p>We can get the capacity by using the <a href="https://pkg.go.dev/builtin#cap">cap</a> function just like the <code>len</code> function. In slices, we can use the lens function to get the length and the cap function to get its underlying capacity to grow.</p>
<p>It's just a way for us to manage the memory reallocation for a slice. Slices under the hood are arrays with a more dynamic punch. So, now we can a bit confidently move to the make function for declaring slices.</p>
<h3>Back to make function</h3>
<p>The Make function is used to declare and initialize a slice (not only slice but maps and channels as well). The function primarily takes in 3 parameters namely, the type of slice, the initial length of the slice, and optionally the capacity of the slice. If we don't provide the capacity, the capacity is set the same as the length.</p>
<pre><code class="language-go">var langs = make([]string, 3, 5)

langs[0], langs[1], langs[2] = &quot;Python&quot;, &quot;Go&quot;, &quot;Javascript&quot;
fmt.Println(langs)

fmt.Printf(&quot;Length = %d 
Capacity = %d
&quot;, len(langs), cap(langs))

langs = append(langs, &quot;Java&quot;, &quot;Kotlin&quot;, &quot;PHP&quot;)

fmt.Println(langs)
fmt.Printf(&quot;Length = %d 
Capacity = %d
&quot;, len(langs), cap(langs))
</code></pre>
<pre><code>$ go run slices.go
[Python Go Javascript]
Length = 3
Capacity = 5
[Python Go Javascript Java Kotlin PHP]
Length = 6
Capacity = 10
</code></pre>
<p>So, there are a lot of things to take in here. We use the make function by parsing in three parameters as said the type in the form of <code>[]type</code> as a distinction for creating slices because we also use the map to create maps and channels. The next two parameters are length and capacity. So, we have initialized a slice of type string, length 3 i.e. we are saying we will initially access only three elements from the array, finally, the third parameter is the capacity which will be the already initialized array under the hood for the slice. So, we have already created an array(under the hood) with 5 elements initialized but only 3 accessible from the slice interface.</p>
<p>Further, we initialize/modify the elements in the created slice. We set 3 elements to some strings and that way we have all elements filled with non-default values in the slice. Now if we say <code>langs[3] = &quot;Something&quot;</code> it would give an error like <code>panic: runtime error: index out of range [3] with length 3</code>. This is a panic in golang which can be triggered in one of the ways in the slice when you access an unreferenced element in it. We have initialized the element but not in the slice interface. So, we have a particular function called append which appends and thus grows the length beyond its current length and refers to the elements initialized from the capacity.</p>
<p>So initially the capacity was 5 and after adding the 6th element, it doubled down to <code>10</code>. So, under the hood, all 5 elements in the array would have been re-allocated into a new memory location and the sixth element would have been added to the next location. This is how we efficiently re-allocate memory for elements in slices. We have a bit more control over the length and are more flexible than arrays.</p>
<h3>Using new function</h3>
<p>The <a href="https://pkg.go.dev/builtin#new">new</a> function in golang is used to allocate a slice/any other type. We can use the new function so as to mimic the make function by adding a bit of value like the initial capacity and length. You can refer to the <a href="https://www.golangprograms.com/go-language/slices-in-golang-programming.html">article</a> for providing the original idea.</p>
<pre><code class="language-go">langs2 := new([3]string)[0:2]

langs2[0], langs2[1] = &quot;Python&quot;, &quot;Go&quot;
fmt.Println(langs2)

fmt.Printf(&quot;Length = %d 
Capacity = %d
&quot;, len(langs2), cap(langs2))

langs2 = append(langs2, &quot;Java&quot;, &quot;Kotlin&quot;, &quot;PHP&quot;)

fmt.Println(langs2)
fmt.Printf(&quot;Length = %d 
Capacity = %d
&quot;, len(langs2), cap(langs2))
</code></pre>
<pre><code>$ go run slices.go
Length = 2
Capacity = 3
[Python Go Java Kotlin PHP]
Length = 5
Capacity = 6
</code></pre>
<p>So, it would work almost similar to the make function. But by default, it would work as nil if you don't provide any length in the <code>[]</code> brackets like <code>new([]string)</code>. This will create a empty slice <code>[]</code> with zero capacity and zero length.</p>
<h2>Adding elements in Slice</h2>
<p>We had a few spoilers for this already with the <code>append</code> function. The <a href="https://pkg.go.dev/builtin#append">append</a> takes in the variable and then the list of values that we want to add. Here, if the capacity of the slice is exceeded, it re-allocates the slice to a new location, and the elements are moved to that location and then the provided elements are added.</p>
<pre><code class="language-go">var percentages = []float64{78.8, 85.7, 94.4, 79.8}
fmt.Println(percentages)
percentages = append(percentages, 60.5, 75.6)
fmt.Println(percentages)
</code></pre>
<pre><code>$ go run slices.go
[78.8 85.7 94.4 79.8]
[78.8 85.7 94.4 79.8 60.5 75.6]
</code></pre>
<p>So, here we can see the append function adding the elements in a slice. You can pass as many elements(MaxInt i.e. int64 or int32 elements precisely) you require in the append function. Calling the append function, again and again, might degrade the efficiency though, so make sure to add a right number of elements in a single call.</p>
<h2>Accessing and Modifying elements in Slice</h2>
<p>We can simply access the elements using the index in the <code>[]</code> brackets. But there is more to that in slices. We can actually get slices of slices. Even in arrays or slices, we can get the particular elements between a specific two indices like 2 and 4, so we can write <code>array[2:4]</code> to get elements at index <code>2</code>, and <code>3</code> the upper bound is non-inclusive. But if we want to have all the elements from a specific index to the last element, we can leave the number blank as <code>[2:]</code> would give elements from index 2 to the last index in the array/slice.</p>
<pre><code class="language-go">scores := []int{80, 85, 90, 75, 60, 56, 83}
fmt.Println(scores)
fmt.Println(&quot;From index 2 to 4&quot;, scores[2:5])
fmt.Println(&quot;From index 0 to 2&quot;, scores[:3])
fmt.Println(&quot;From index 3 to 5&quot;, scores[3:])
</code></pre>
<pre><code>$ go run slices.go
[80 85 90 75 60 56 83]
From index 2 to 4 [90 75 60]
From index 0 to 2 [80 85 90]
From index 3 to 5 [75 60 56 83]
</code></pre>
<p>So, we are able to perform index slicing in golang on arrays and slices.
Further, as for the array, we can also modify elements in slices. Using the index of that element, we can access the element and perform operations on it and thus change the literal value of the element.</p>
<pre><code class="language-go">word := []byte{'f', 'u', 'z', 'z', 'y'}
fmt.Printf(&quot;%s
&quot;, word)
word[0] = 'b'
word[len(word)-1] = 'z'
fmt.Printf(&quot;%s
&quot;, word)
</code></pre>
<pre><code>fuzzy
buzzz
</code></pre>
<p>So, now we can also modify existing values of elements in slices.</p>
<h2>Deleting elements from Slice</h2>
<p>We can also remove an element from the slice i.e. shrink the length of the slice. There is no function to remove an element from the slice, but we can work around with the append function in golang. So, in the slice before the element's index to be deleted is appended with all the elements after the index of the element to be deleted.</p>
<pre><code>10 20 30 40 50 60
0  1  2  3  4  5

Delete the element at index -&gt; 2

Copy from 3 to 5 into the slice from 0 to 1

// append(slice[:2], slice[2+1:]
            ^            ^
            |            | 
          10 20       40 50 60        

append 40 50 60 -&gt; 10 20

10 20 40 50 60
</code></pre>
<p>Here in the above example, we append the indices after the element to be deleted into the slice of elements before the <code>to be deleted element</code>.</p>
<pre><code class="language-go">marklist := []int{80, 85, 90, 75, 60}
fmt.Println(marklist)

var index int
fmt.Printf(&quot;Enter the index to be deleted: &quot;)
fmt.Scan(&amp;index)

elem := marklist[index]


// append in such a way that the element to be removed is excluded

marklist = append(marklist[:index], marklist[index+1:]...)


fmt.Printf(&quot;The element %d was deleted.
&quot;, elem)
fmt.Println(marklist)
</code></pre>
<pre><code>$ go run slices.go

[80 85 90 75 60]
Enter the index to be deleted: 3
The element 75 was deleted.
[80 85 90 60]
</code></pre>
<h2>Iterate through a slice</h2>
<p>As arrays are under the hood modifications of arrays, we have a quite similar approach to iterating over slices in golang.</p>
<h3>Using three statements for loop</h3>
<p>We can use the three statements for loop i.e. the initialization, condition, and incrementation procedure. The counter is set from 0 or any other starting value as the index of the slice, next we have the end loop condition i.e. a condition to check until when to exit, and finally the amount to which we need to increment the counter.</p>
<pre><code class="language-go">code := [7]rune{'g', 'o', 'l', 'a', 'n', 'g'}
for i := 0; i &lt; len(code); i++ {
    fmt.Printf(&quot;%c
&quot;, code[i])
}
</code></pre>
<pre><code>$ go run slices.go
g
o
l
a
n
g
</code></pre>
<h3>Using Range-based for loop</h3>
<p>We can use range-based for loops to iterate over the slice elements. The range keyword is passed with the slice name or the slice of an array to iterate over. Using the two variables i.e. the iterator and the copy of the element, we can access the index and the element in the slice.</p>
<pre><code class="language-go">scores := []int{80, 85, 90, 75, 60, 56, 83}
for _, s := range scores {
    fmt.Println(s)
}
</code></pre>
<pre><code>$ go run slices.go
80
85
90
75
60
56
83
</code></pre>
<p>We can also use a slice of slice i.e. scores[:4] to access a specific element in the range using index slicing.</p>
<pre><code class="language-go">scores := []int{80, 85, 90, 75, 60, 56, 83}
for _, s := range scores[1:4] {
    fmt.Println(s)
}
</code></pre>
<pre><code>$ go run slices.go
85
90
75
</code></pre>
<h3>Using for loop with range</h3>
<p>We can even use them for loop as a while loop to look and get a bit of both the above methods like the range method and the typical for loop access.</p>
<pre><code class="language-go">start, i, end := 2, 2, 5

modes := []string{&quot;normal&quot;, &quot;command&quot;, &quot;insert&quot;, &quot;visual&quot;, &quot;select&quot;, &quot;replace&quot;}

for range scores[start:end] {
    fmt.Printf(&quot;Element at index %d = %s 
&quot;, i, modes[i])
    i++
}
</code></pre>
<pre><code>$ go run slices.go
Element at index 2 = insert
Element at index 3 = visual
Element at index 4 = select
</code></pre>
<p>So by using the range keyword we were able to iterate over the slice but without assigning the iterator and the copy of the element, we manually set a counter <code>i</code> and increment it as per our liking. Using index slicing we were able to get the elements between particular indices.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So, from this part of the series, we were able to understand the basics of slices in golang. We covered some basics stuff including the declaration, initialization, and iteration. We also covered the under the hood working of slices and how to relate with the arrays.
Thank you for reading. If you have any questions or feedback, please let me know in the comments or on social handles. Happy Coding :)</p>
