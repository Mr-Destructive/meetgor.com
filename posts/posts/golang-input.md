{
  "title": "Golang: Input",
  "status": "published",
  "slug": "golang-input",
  "date": "2025-04-05T12:33:44Z"
}

<h2>Introduction</h2>
<p>In this fourth section of Golang, we will be understanding the basics of user input. In golang, we can get user input with several functions most of which are similar to the <code>C</code> programming language like <code>scanf</code>. This type of input is quite powerful and gives more control on the input to be received.</p>
<h2>Scan Function</h2>
<p>The <a href="https://pkg.go.dev/fmt#Scan">Scan</a> function helps in getting a value with space as delimiter i.e. The input is stored before a space is encountered. This means the input is only limited to adding a space or a new line. We can use the function by passing the reference to the variable we are going to store the input value. So, we can have a basic input in Golang as follows:</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
    var pname string
    fmt.Println(&quot;Enter your favourite programming language: &quot;)
    fmt.Scan(&amp;pname)
    fmt.Println(&quot;So, your favourite programming language is&quot;,pname)
}
</code></pre>
<pre><code>$ go run scan.go
Enter your favorite programming language:
python
So, your favorite programming language is python
</code></pre>
<p>We need to declare the variable to take input as we need a reference of that variable to store the input. We will be talking about <code>&amp;</code> and pointers in a separate article. We use the <code>Scan</code> function by passing the reference to the variable <code>pname</code> like <code>&amp;pname</code> which means, fetch the memory address of the variable <code>name</code>, we just pass the address as <code>int</code> to the Scan function and it does the rest to store the input value in it. We then as usual access the variable and operations on it.</p>
<p>Here, if you add a space in the input, the value after the space won't be picked by the Scan function. It strictly stops accepting values input after it sees space. We can use this to input multiple variables at once. We know scan gets input before encountering space, so we can pass multiple variable references and add them as input.</p>
<pre><code class="language-go">var (
    name   string
    age    int
    gender rune
)
fmt.Println(&quot;Enter your name age and gender: &quot;)
fmt.Scan(&amp;name, &amp;age, &amp;gender)
fmt.Printf(&quot;Hello %s, you are a %c and %d years old&quot;, name, gender, age)
</code></pre>
<pre><code>$ go run scan.go
Enter your name age and gender:
Meet 19 77
Hello Meet, you are a M and 19 years old
</code></pre>
<p>Here, we are declaring multiple variables like <code>name</code>, <code>age</code>, and <code>gender</code> as <code>string</code>, <code>int</code>, and <code>rune</code> respectively. Then, we can input all of these in a single scan statement by comma-separated variables. Here, we need to input the <code>rune</code> as an int value because under the hood it is an integer alias. So, we inputted <code>77</code> which is equivalent to <code>M</code> in ASCII characters and even Unicode character sets. Thus, we were able to input multiple variables with the Scan function.</p>
<h2>Scanf functions</h2>
<p>The <a href="https://pkg.go.dev/fmt#Scanf">Scanf</a> function is quite similar to the <code>scanf</code> in C programming language as it allows to specify the type of the incoming input. This will solve the problem of us inputting <code>77</code> instead of <code>M</code> in the gender variable in the previous example. The Scanf function allows us to take input by specifying the placeholder types and the delimiters as well. The delimiter is basically the separator between two or more entities. We can either use space separation or <code> </code> as an input delimiter i.e. the way we want to separate inputs from each other while taking input.</p>
<pre><code class="language-go">var (
    name   string
    age    int
    gender rune
)
fmt.Println(&quot;Enter your name age and gender: &quot;)
fmt.Scanf(&quot;%s %d %c&quot;, &amp;name, &amp;age, &amp;gender)
fmt.Printf(&quot;Hello %s, you are a %c and %d years old&quot;, name, gender, age)
</code></pre>
<pre><code>$ go run scanf.go
Enter your name age and gender:
Meet 12 M
Hello Meet, you are a M and 12 years old


</code></pre>
<p>How cool is that? It definitely gives much more control on what and how to take input. We are taking input as only space-separated values. Let's now try to get more control over how the input will be taken and stored.</p>
<pre><code class="language-go">var (
    name   string
    age    int
    gender rune
)
fmt.Println(&quot;Enter your name age and gender: &quot;)
fmt.Scanf(&quot;%s 
 %d %c&quot;, &amp;name, &amp;age, &amp;gender)
fmt.Printf(&quot;Hello %s, you are a %c and %d years old&quot;, name, gender, age)
</code></pre>
<pre><code>$ go run scanf.go
Enter your name age and gender:
Meet
12 M
Hello Meet, you are a M and 12 years old
</code></pre>
<p>By adding <code> </code> between the <code>%s</code>(name) and <code>%d</code>(age), we want the user to type the name on one line and age with gender on a different line. The age and gender as before separated by space.</p>
<h2>Scanln function</h2>
<p>The <a href="https://pkg.go.dev/fmt#Scanln">Scanln</a> function is a modification of the Scan function as it only stops the input after a newline/enter is pressed.  So, using this we can input multiple variables which are space-separated in a single line.</p>
<pre><code class="language-go">var s string
fmt.Println(&quot;Enter a string: &quot;)
fmt.Scanln(&amp;s)
fmt.Println(s)
</code></pre>
<pre><code>$ go run scanln.go
Enter a string:



$ go run scanln.go
Enter a string:
Can't type
Can't

$ ype
-bash: ype: command not found
</code></pre>
<p>The Scanln function even accepts an empty string as input. It just needs to get the new line character and it will exit, it also only accepts space-separated values. The rest of the input after space is thrown away and is basically exited from the program stream. More specifically, the input <code>Can't Type</code> was treated only as <code>Can't</code> anything after the space is not considered in the input value.</p>
<p>The key difference between Scan and Scanln is that Scanln will not accept input that is space-separated, Scan function considers the newline/enter as a space if there are multiple inputs. The below example will make things absolutely clear.</p>
<pre><code class="language-go">// scan.go
package main

import &quot;fmt&quot;

func main() {
    var (
        name   string
        age    int
        gender rune
    )
    fmt.Println(&quot;Enter your name age and gender: &quot;)
    fmt.Scan(&amp;name, &amp;age, &amp;gender)
    fmt.Printf(&quot;Hello %s, you are a %c and %d years old&quot;, name, gender, age)
}

</code></pre>
<pre><code class="language-go">//scanln.go
package main

import &quot;fmt&quot;

func main() {
    var s string
    fmt.Println(&quot;Enter a string: &quot;)
    fmt.Scanln(&amp;s)
    fmt.Println(&quot;Inputted string : &quot;, s)
}
</code></pre>
<pre><code>$ go run scan.go
Enter your name age and gender:

Meet

14



77
Hello Meet, you are a M and 14 years old

$ go run scanln.go
Enter a string:

Inputted string :
</code></pre>
<p>We can see that, The Scan function won't exit until it has inputted all its input values even with newline and spaces. Whereas the Scanln function just waits for the newline character (Enter Key) to be pressed and it exits, thereby even allowing an empty string as input.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So, these are the basic input techniques in Golang. We saw functions in the <code>fmt</code> package like <code>Scan</code>, <code>Scanf</code>, and <code>Scanln</code> which allow us to get input in a specific pattern. Hopefully, from this article part, we can build a firm base for further exploration like Strings, Arrays, and the ways to input them. Thank you for reading. If you have any questions or feedback, please let me know in the comments or on social handles. Happy Coding :)</p>
