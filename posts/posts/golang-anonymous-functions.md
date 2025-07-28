{
  "title": "Golang: Anonymous Functions",
  "status": "published",
  "slug": "golang-anonymous-functions",
  "date": "2025-04-05T12:33:38Z"
}

<h2>Introduction</h2>
<p>We have looked at the defer keyword in golang in the <a href="https://www.meetgor.com/golang-defer/">previous</a> part of the series, in this section, we will understand how we can use anonymous functions in golang. We will explore how to declare and use anonymous functions with a few examples.</p>
<h2>What are Anonymous Functions</h2>
<p>Anonymous functions are quite simple to understand, we don't define a function, we declare it and call it instantly. An anonymous function doesn't have a name so hence it is called an anonymous function. As a normal function it can take in parameters and return values. With anonymous functions, we can bind the operations to a variable or a constant as a literal(value). If an anonymous function takes in a parameter, it needs to be parsed immediately after the end of the function body. We will see how we define the syntax and specifications of the anonymous functions in golang.</p>
<h2>Simple Anonymous functions</h2>
<p>To create a simple anonymous function we use the same function syntax without giving it a name like <code>func() {}</code>, inside the function body i.e. <code>{}</code>, you can define the operations that need to be performed.</p>
<p>Here, we have created an anonymous function that simply calls the <code>fmt.Println</code> function with a string. So, we have made things a little too much as we can even directly call the <code>fmt.Println</code> function from the main function, instead we have called an anonymous function that in turn calls the <code>fmt.Println</code> function. It might not make sense to use an anonymous function here, but it can be helpful for other complex tasks that you want to isolate the logic without creating a dedicated function for the process.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {

  func() {
    fmt.Println(&quot;Hello, Anonymous functions&quot;)
  }()

}
</code></pre>
<pre><code>go run anonymous_function.go

Hello, Anonymous functions
</code></pre>
<p>So, this is how we create a basic anonymous function in golang, this can be further used for returning values from the function or passing parameters into the function and also assigning the function call to a variable or a constant.</p>
<h2>Assigning anonymous function to variables</h2>
<p>We can assign the call to the anonymous function to a variable or a constant and call the function as many times as we require. So, we can basically store the function logic in a variable and call it whenever we require the function with the <code>()</code> parenthesis as an indication to call the function.</p>
<p>In the following example, we have used the variable <code>draw</code> to store the function call which prints <code>Drawing</code> with the help of the <code>fmt.Println</code> function. The draw variable now contains the function and not its value. So be careful here, the anonymous function which we have defined as the <code>draw</code> variable's literal value, it's like we are giving a name to this anonymous function and the name will be the variable name so we have created the function <code>draw</code> which is an anonymous function stored in a variable.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {

  draw := func() {
    fmt.Println(&quot;Drawing&quot;)
  }
  draw()
  draw()
}
</code></pre>
<pre><code>go run anonymous_function.go

Drawing
Drawing
</code></pre>
<p>We see that we call the variable <code>draw</code> as a function call by appending <code>()</code> parenthesis to it as <code>draw()</code> this thereby calls the anonymous function stored inside the variable value.</p>
<p>The main thing to note here is that we are not adding <code>()</code> at the time of declaring the anonymous function, as it will call the function directly. The problem then will arise because the function is not returning anything so we can't assign it to a variable.</p>
<pre><code class="language-go">draw := func() {
  fmt.Println(&quot;Drawing&quot;)
}()
</code></pre>
<pre><code>functions/anonymous_functions.go:10:11: func() {…}() (no value) used as value
</code></pre>
<p>So, we are trying to assign a variable to a function call that has no return value. This has no value, not even nil, so we get an error for assigning a variable to nothing.</p>
<p>Though if you had a return value from the function, we can directly assign the function call to the variable as it has returned a valid value defined in the function literal.</p>
<h2>Parsing parameters</h2>
<p>We can even parse parameters to the anonymous functions as any other normal function. We define the name of the variable followed by the type of the variable in the parenthesis to use these parameters inside the anonymous function.</p>
<p>We need to keep in mind that these function parameters can either be passed with the variable call or directly at the time of defining the function.</p>
<p>In the below example, we have created a variable <code>call</code> and are assigning it to an anonymous function that takes in a parameter <code>name</code> which is a <code>string</code>, and prints some text to the console. So, we call the funciton <code>call</code> with a parameter as a string, as we have <code>&quot;Meet&quot;</code> and <code>person := &quot;Chris&quot;</code> as a string passed to the <code>call</code> function.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {

  call := func(name string) {
    fmt.Println(&quot;Calling,&quot;, name)
  }

  call(&quot;Meet&quot;)
  person := &quot;Chris&quot;
  call(person)

}
</code></pre>
<pre><code>go run anonymous_function.go

Calling, Meet
Calling, Chris
</code></pre>
<p>Here, we can see that the function <code>call</code> prints the text that we have defined to call the <code>fmt.Println</code> function. We parse the function with the string literal as the anonymous function takes a parameter in the form of a string. We can parse multiple parameters to the anonymous function as we can with the normal function like slices, maps, arrays, struct, etc.</p>
<h2>Returning values</h2>
<p>We can even return values from the anonymous function if we want to instantly call the function and save the <code>returned</code> value in the variable. We can return single or multiple values as per our requirements just like any normal function in golang.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {

  is_divisible := func(x int, y int) bool{
    var res bool 
    if x % y == 0 {
      res = true
    } else{
      res = false
    }
    fmt.Println(res)
    return res
  }

  fmt.Printf(&quot;%T
&quot;, is_divisible)
  is_divisible(10, 5)
  is_divisible(33, 5)

  divisblity_check := is_divisible(45, 5)
  fmt.Printf(&quot;%T : %t
&quot;, divisblity_check, divisblity_check) 

}
</code></pre>
<pre><code>go run anonymous_function.go

func(int, int) bool 
true
false
bool : true
</code></pre>
<p>As we can see the function has returned a boolean value and we store it in a variable <code>divisblity_check</code>, the variable which contains the function i.e. <code>is_divisible</code> can then be called, and thereby we get the returned value in the variable as a boolean as we print the type and the value of the variable <code>divisblity_check</code>. We also can see that the type of the variable <code>is_divisible</code> is of type function literal which takes <code>int, int</code> and has a return value as <code>bool</code>.</p>
<p>We can also do one more interesting here, which we were restricted previously in the above examples. We can directly call the function and store it as a value rather than the function itself. So, we can only use the value returned from the function but can't call the function later.</p>
<pre><code class="language-go">is_divisible := func(x int, y int) bool{
  var res bool 
  if x % y == 0 {
    res = true
  } else{
    res = false
  }
  fmt.Println(res)
  return res
}(13, 4)

fmt.Printf(&quot;%T
&quot;, is_divisible)
fmt.Printf(is_divisible) 
</code></pre>
<pre><code>go run anonymous_function.go

bool
false
</code></pre>
<p>So, in the above-modified example, we have passed in the parameter instead of a callable function. This will store the returned value of the function call. So, we will store the boolean value in the <code>is_divisible</code> and we will have to pass the integer values to the function which we have parsed as <code>(13, 4)</code> to the anonymous function call.</p>
<p>In the below example, we have created an anonymous function that takes in three parameters like <code>(string, int, string)</code> and returns a string. We have used <code>fmt.Sprintf</code> function to format the variable and store it in a variable, we then return the string. This anonymous function is then directly called and we store the returned value instead of the function.</p>
<p>So, in this example, the <code>format_email</code> variable will store a string instead of acting it as a function as a callable object.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {

  format_email := func(name string, age int, company string) string{
    email := fmt.Sprintf(&quot;%s.%d@%s.com&quot;, name, age, company)
    return email
  }(&quot;john&quot;, 25, &quot;gophersoft&quot;)

  fmt.Println(format_email)
  fmt.Printf(&quot;%T
&quot;,format_email)

}
</code></pre>
<pre><code>go run anonymous_function.go

john.25@gophersoft.com   
string
</code></pre>
<p>As we can see the <code>format_email</code> variable only returns a string instead of a callable object. We have directly parsed the parameters to the anonymous function and hence it instantly calls it and we return the string.</p>
<h2>Passing Anonymous function to function parameters</h2>
<p>We can even pass an anonymous function to a parameter to a function. This can be quite helpful for writing some simple logic inside a complex script. We can pass the function type as a parameter to the parameter and theere we can parse the actual data and perform the desired operation.</p>
<p>The below example is a bit of code to write but makes a lot of sense to understand anonymous functions in golang. The function <code>get_csv</code> is a function which takes in three parameters <code>int, string, func(string)[] string</code>. The third parameter is a function literal that takes in a string and returns a slice of string. So, we are basically converting a string <code>&quot;kevin,21,john,33&quot;</code> into a slice of the slice like <code>[[kevin 21] [john 33]]</code>, this is basically separating values with <code>,</code> comma as the delimiter and then processing slices to create a single slice.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;
import &quot;strings&quot;

func get_csv(index int, str string, t func(csv string)[]string) [][]string{
  s := t(str)
  var res [][]string
  for i := 1; i&lt;len(s); i+=2 {
    var data []string
    data = append(data, s[i-1], s[i])
    res = append(res, data)
  }
  return res
}

func main() {

  csv_slice := func (csv string) []string{
    return strings.Split(csv, &quot;,&quot;)
  }
  csv_data := get_csv(2,&quot;kevin,21,john,33,george,24&quot;, csv_slice)
  fmt.Println(csv_data)
  for _, s := range csv_data{
    fmt.Printf(&quot;%s - %s
&quot;,s[0],s[1])
  }
</code></pre>
<pre><code>go run functions/anonymous_function.go

[[kevin 21] [john 33] [george 24]]

kevin - 21
john - 33
george - 24
</code></pre>
<p>Let's break down the code one by one, we will start with the main function, where we have <code>csv_slice</code> as a function literal and is an anonymous function that takes in a string and returns a slice of string. The function returns a call to the function <a href="https://pkg.go.dev/strings#Split">strings.Split</a> taking in the string from the function parameter. We then call the function <code>get_csv</code> with parameters <code>(2, &quot;kevin,21....&quot;, csv_slice)</code>, this function is defined outside the main. The function takes in three parameters as discussed and parsed from the main function and it returns a slice of type string.</p>
<p>So, inside the <code>get_csv</code> function, we define <code>s</code> as the function call to <code>t(str)</code> which if you look carefully is a function call to <code>csv_slice</code> with parameter as a string. This function call returns a slice of strings separated by <code>,</code>. So that's all the logic we need to understand anonymous function from parameters. We have used a function literal to call the function from another function. In this case, the funciton is an anonymous function assigned to a variable.</p>
<pre><code>&quot;kevin,21,john,33,george,24&quot;
            ||
            \/
[kevin 21 john 33 george 24]
            ||
            \/
[[kevin 21] [john 33] [george 24]]

</code></pre>
<p>Further, after we have <code>s</code> which would look like <code>[kevin 21 john 33 george 24]</code> as each individual element. Thereafter we create an empty <a href="https://www.geeksforgeeks.org/slice-of-slices-in-golang/">slice of slice</a> string as <code>res</code> and operate a loop through the slice jumping 2 indexes at a time. Why? because we want a slice of two elements combined. So, we also create a slice of string named <code>data</code> and we add the two components to it like <code>[kevin 21]</code> with the <a href="https://pkg.go.dev/builtin#append">append</a> function, and this slice is appended to a slice of slices <code>res</code> as <code>[[kevin 21]]</code> thereby iterating over the entire slice and creating the desired slice data. We return the <code>res</code> from the function. This get's us back to the main function which simply iterates over the slice and we get the desired data from it.</p>
<p>So, this is how we convert an extremely easy task to extremely complicated code :)</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/functions/anonymous_function.go">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>That is it from this part, we covered a little bit about anonymous functions in golang. Anonymous functions are simply function literals that can be used to do a lot of quick operations without needing to write an explicit function in the program. Further, in the next part look into closures which are a bit complex to understand and require the understanding of anonymous functions.</p>
<p>Thank you for reading, if you have any queries, feedback, or questions, you can freely ask me on my social handles. Happy Coding :)</p>
