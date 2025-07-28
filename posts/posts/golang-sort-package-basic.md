{
  "title": "Golang: Sort Package Introduction",
  "status": "published",
  "slug": "golang-sort-package-basic",
  "date": "2025-04-05T12:33:28Z"
}

<p>I have been learning Golang for around 2 years now, and I have never paid attention to the sort package, can you believe that! Where was this package hiding?</p>
<p>The <code>sort</code> package provides convenient methods for sorting slices and user-defined collections in Go. Here's a quick overview of what it can do:</p>
<h2>Sorting Slices</h2>
<p>To sort a slice of builtin types like ints, strings, etc, you can simply call <code>sort.Slice()</code>:</p>
<pre><code class="language-go">nums := []int{5, 2, 6, 3, 1}

sort.Slice(nums, func(i, j int) bool {
    return nums[i] &lt; nums[j]
})

// or 
// sorts.Ints(nums)

fmt.Println(nums)
</code></pre>
<pre><code class="language-bash">$ go run main.go

[1, 2, 3, 5, 6]
</code></pre>
<p>The sort is in-place, mutating the original slice. You can customize the sort order by providing a less(i, j int) bool function.</p>
<h2>Sorting Struct Slices</h2>
<p>To sort a slice of custom structs, add a Len(), Less(i, j int) bool, and Swap(i, j int) method:</p>
<pre><code class="language-go">type Person struct {
  Name string
  Age  int
}

people := []Person{
  {&quot;Bob&quot;, 30},
  {&quot;John&quot;, 20},
  {&quot;Alice&quot;, 25},
}

sort.Slice(people, func(i, j int) bool {
  return people[i].Age &lt; people[j].Age
}) 

fmt.Println(person)
</code></pre>
<pre><code class="language-bash">$ go run main.go
[{&quot;Name&quot;:&quot;Alice&quot;,&quot;Age&quot;:25},{&quot;Name&quot;:&quot;Bob&quot;,&quot;Age&quot;:30},{&quot;Name&quot;:&quot;John&quot;,&quot;Age&quot;:20}]
</code></pre>
<p>This will sort people by age.</p>
<h2>Sorting Maps</h2>
<p>Maps are inherently unordered in Go. We can't sort them, but we can iterate them in a sorted way. We need a separate slice of keys or values(whicher required), we will sort those keys/values and iterate over them in that order.:</p>
<h3>Sort by Keys</h3>
<p>To sort a map by keys, we can use the <code>sort.Strings()</code> function or any other sort function as per the data structure, there are functions like <a href="https://pkg.go.dev/sort#Ints">sort.Ints</a>, <a href="https://pkg.go.dev/sort#Float64s">sort.Float64</a>, or <a href="https://pkg.go.dev/sort#Slice">sort.Slice</a>. We can create a new slice of keys from the map and then apply the sort on that newly created slice. After the slice of keys is created, we can iterate over it and access the map values in a order of sorted keys.</p>
<pre><code class="language-go">counts := map[string]int{
  &quot;hello&quot;: 5,
  &quot;world&quot;: 2,
  &quot;foo&quot;: 3,
}

keys := make([]string, 0, len(counts))
// extract keys 
for k := range counts {
  keys = append(keys, k)
} 

// sort keys
sort.Strings(keys) 

// iterate with the sorted keys slice
for _, k := range keys {
  fmt.Println(k, counts[k]) 
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

foo 3
hello 5
world 2
</code></pre>
<p>This prints the map ordered by key. You can sort by value too.</p>
<h3>Sort by Value</h3>
<p>To iterate the map with sorting order of values, we can approach it in a similar way. We create a slice of keys. This time, we don't sort the keys, instead, we change the position of the slice of keys depending on the values. This changes the key order based on the sorted values in the map. By similarly iterating over the key slice, we iterate over the map in a sorted order of the values.</p>
<pre><code class="language-go">counts := map[string]int{
  &quot;hello&quot;: 5,
  &quot;world&quot;: 2,
  &quot;foo&quot;: 3,
}

keys := make([]string, 0, len(counts))
// extract keys
for k := range counts {
  keys = append(keys, k)
}

// sort by value
// i.e. change the order of key based on values sort order
sort.SliceStable(keys, func(i, j int) bool {
    return counts[keys[i]] &lt; counts[keys[j]]
})

// iterate sorted
for _, k := range keys {
  fmt.Println(k, counts[k])
}
</code></pre>
<pre><code class="language-bash">$ go run main.go

world 2
foo 3
hello 5
</code></pre>
<p>The <code>sort.SliceStable</code> function is used to sort the slice in place. It takes in a slice and the less function which is the actual logic for comparison in the values. This function sorts the key elements based on the values in the map, and thereby the key slice is shuffled by the sorted order of the values in the map.</p>
<p>The sort package is super useful for quickly organizing Go data structures.
There are <a href="https://pkg.go.dev/sort#Find">Find</a> which tries to return a index of the first element that satisfies a comparison condition with a flag or found or not with a boolean. There is also a <a href="https://pkg.go.dev/sort#Search">Search</a> which is used specifically for searching elements in a sorted array, it also gives a first index of the occuring element. But that is a topic for another day!</p>
<p>Thank you, Happy Coding :)</p>
