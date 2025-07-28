{
  "title": "Golang: Conditionals and Loops",
  "status": "published",
  "slug": "golang-conditionals-loops",
  "date": "2025-04-05T12:33:44Z"
}

<h2>Introduction</h2>
<p>Moving to the fourth part, we will be doing conditional statements and loops in golang. We will be seeing the basics of conditional statements like if-else and switch along with loops like for, while, and range-based loops. We won't be covering iterating over arrays in a loop as this requires an understanding of arrays.</p>
<h2>Conditional statements</h2>
<p>Conditional statements are quite a fundamental aspect of learning a programming language. In golang, we have if-else conditional statements as well as switch cases. We will be exploring both of them in this section. Firstly, we will dive into if-else statements which are quite easy to understand.</p>
<h3>if else</h3>
<p>An <code>if</code> statement is used for checking the validity of a condition. If the condition is true(condition is met), a particular sets of statements are executed else (condition is not satisfied) different statements gets executed. We can use a basic <code>if-else</code> statement in go with the following syntax:</p>
<pre><code class="language-go">if condition {
    // statements
}else{
    //statements
}
</code></pre>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
    age := 16
    if age &lt; 13{
        fmt.Println(&quot;Kid&quot;)
    }else{
        fmt.Println(&quot;Teenager&quot;)
    }
}
</code></pre>
<pre><code>$ go run if_else.go
Teenager
</code></pre>
<p>We can also use else if for evaluating more than one condition rather than using nested if and else.</p>
<pre><code class="language-go">if condition {
    // statements
}else if condtion {
    //statements
}else if condition {
    //statements
}else {
    //statements
}
</code></pre>
<pre><code class="language-go">year := 2

if year &lt; 1 {
    fmt.Println(&quot;Freshman&quot;)
} else if year == 2 {
    fmt.Println(&quot;Sophomore&quot;)
} else if year == 3 {
    fmt.Println(&quot;Junior&quot;)
} else if year == 4 {
    fmt.Println(&quot;Senior&quot;)
} else {
    fmt.Println(&quot;Probably a Graduate&quot;)
}
</code></pre>
<pre><code>$ go run if_else.go
Sophomore
</code></pre>
<p>Using <code>else if</code> we can evaluate multiple conditions. This style is much better than using nested <code>if else</code> statements as it becomes harder to read for complex cases.</p>
<h3>switch</h3>
<p>We also have switch statements in golang which allow us to write cases for a given state of a variable. We can simply add cases for a given variable, the case should be a valid value that the variable can take. If a case is matched it breaks out of the switch statement without executing any statements below the matched case.</p>
<p>The basic structure of the switch statements in golang is as follows:</p>
<pre><code class="language-go">switch variable{
case value1:
    //statements
case value2:
    //statements

}
</code></pre>
<pre><code class="language-go">age := 19
var result string
switch {
case age &lt; 13:
    result = &quot;Kid&quot;
case age &lt; 20:
    result = &quot;Teenager&quot;
case age &lt; 25:
    result = &quot;Adult&quot;
case age &lt; 35:
    result = &quot;Senior&quot;
}
fmt.Println(&quot;The person is a&quot;,result)
</code></pre>
<pre><code>$ go run switch.go
The person is a Senior with age 27.

$ go run switch.go
The person is a Teenager with age 19.

$ go run switch.go
The person is a Kid with age 11.
</code></pre>
<p>This gives a good understanding of switch-case statements. We can give a variable to the switch statement and pick its value in the respective case statements to evaluate the result accordingly. The <code>default</code> statement is evaluated when there is no match among the given cases.</p>
<pre><code class="language-go">language := &quot;&quot;
var devs string
switch language{
case &quot;go&quot;:
    devs = &quot;gopher&quot;
case &quot;rust&quot;:
    devs = &quot;rustacean&quot;
case &quot;python&quot;:
    devs = &quot;pythonista&quot;
case &quot;java&quot;:
    devs = &quot;Duke&quot;
default:
    language = &quot;javascript&quot;
    devs = &quot;developer&quot;
}
fmt.Println(&quot;A person who codes in&quot;,language,&quot;is called a&quot;,devs)
</code></pre>
<pre><code>$ go run switch.go
A person who codes in javascript is called a developer

$ go run switch.go
A person who codes in python is called a pythonista

$ go run switch.go
A person who codes in go is called a gopher
</code></pre>
<p>This code will by default pick <code>javascript</code> and <code>developer</code> as the values for <code>language</code> and <code>devs</code> respectively if there is no match for the provided language or the language is left empty.</p>
<p>We also have <code>fallthrough</code> in the golang switch which allows evaluating more than one case if one of them is matched. This will allow the switch and check for all the cases sequentially and evaluate all the matched and satisfied cases.</p>
<pre><code class="language-go">character := 't'
fmt.Printf(&quot;The input character is = %c
&quot;, character)
switch {
case character == 97:
    fmt.Printf(&quot;Its %c
&quot;, character)
    fallthrough
case character &lt; 107 &amp;&amp; character &gt; 96:
    fmt.Println(&quot;It's between a and k&quot;)
    fallthrough
case character &lt; 117 &amp;&amp; character &gt; 98:
    fmt.Println(&quot;It's between a and t&quot;)
    fallthrough
case character &lt; 122 &amp;&amp; character &gt; 98:
    fmt.Println(&quot;It's between u and z&quot;)
default:
    fmt.Println(&quot;Its not a lowercase alphabet&quot;)
}
</code></pre>
<pre><code>$ go run switch.go
The input character is = f
It's between a and k
It's between a and t
It's between a and u

$ go run switch.go
The input character is = k
It's between a and t
It's between a and u

$ go run switch.go
The input character is = a
Its a
It's between a and k
It's between a and t
It's between a and u
</code></pre>
<p>So, here we can see that the fallthrough hits multiple cases. This is unlike the base case which exits the switch statement once a case has been satisfied. This can be helpful for situations where you really want to evaluate multiple conditions for a given variable.</p>
<h2>Loops</h2>
<p>We can now move on to loops in golang. We only have a <code>for</code> loop so to speak but this can be used as any other looping statement like the <code>while</code> loop or range-based loop. We will first see the most fundamental loop statement in golang which is a three-component loop.</p>
<h3>for loop</h3>
<p>We can have a simple for loop in golang by using the three statements namely <code>initialize</code>, <code>condition</code>, and <code>increment</code>. The structure of the loop is quite similar to the other programming languages.</p>
<pre><code class="language-go">for k := 0; k &lt; 4; k++ {
    fmt.Println(k)
}
</code></pre>
<h3>Range-based loop</h3>
<p>We can even iterate over a string, using the range keyword in golang. We need to have two variables for using a range-based for loop in golang one is the index or the 0 based position of the element and the copy of the element in the array or string. Using the range keyword, we can iterate over the string one by one.</p>
<pre><code class="language-go">name := &quot;GOLANG&quot;
for i, s := range name{
    fmt.Printf(&quot;%d -&gt; %c
&quot;,i, s)
}
</code></pre>
<pre><code>$ go run for.go
0 -&gt; G
1 -&gt; O
2 -&gt; L
3 -&gt; A
4 -&gt; N
5 -&gt; G
</code></pre>
<p>So, here we can see we have iterated over the string by each character. Using the range keyword in golang, The <code>i, s</code> is the index and the copy of the element at that index as discussed earlier. Using the index we get the value which we don't have to index the array for accessing it, that is already copied in the second variable while using the range loop.</p>
<h3>while loop (Go's while is for)</h3>
<p>There are no while loops as such in golang, but the for loop can also work similarly to the while loop. We can use a condition just after the for a keyword to make it act like a while loop.</p>
<pre><code class="language-go">for condition {
    // statements
}
</code></pre>
<pre><code class="language-go">count := 3
for count &lt; 9 {
	fmt.Println(count)
	count++
}
</code></pre>
<pre><code>$ go run while.go
3
4
5
6
7
8
</code></pre>
<p>We can see here that the condition is evaluated and the statements in the loop body are executed, if the condition evaluates to false, the flow is moved out of the loop and we exit the loop.</p>
<h3>Infinite loop</h3>
<p>We can run an infinite loop again using a keyword. We do not have any other keywords for loops in golang.</p>
<pre><code class="language-go">for {
    // statements
    // should have conditons to exit
}
</code></pre>
<pre><code class="language-go">flag := 4
for {
	flag++
	fmt.Println(flag)
}
</code></pre>
<p>This might be used with a base condition to exit the loop otherwise there should be a memory overflow and the program will exit with errors.</p>
<h3>Break</h3>
<p>If we want to exit out of a loop unconditionally, we can use the <code>break</code> keyword. This will break the loop and help us to exit out of an infinite or a condition-bound-based loop too.</p>
<pre><code class="language-go">flag := 1
for {
    fmt.Println(flag)
    flag++
    if flag == 7 {
        fmt.Println(&quot;It's time to break at&quot;, flag)
        break
    }
}
</code></pre>
<pre><code>$ go run infinite.go
1
2
3
4
5
6
It's time to break at 7
</code></pre>
<p>As, we can see inside an infinite loop, we were able to break out of it by using a conditional statement and <code>break</code> keyword. This also applies to switch cases it basically is the opposite of <code>fallthrough</code> in switch-case statements. But by default(without using fallthrough), the case statement breaks the switch after a match has been found or the default case has been encountered.</p>
<h3>Continue</h3>
<p>We also have the opposite of <code>break</code> i.e. <code>continue</code> which halts the execution of the loop and directs back to the post statement increment(in case of for loops) or evaluation(in case of while loop). We basically are starting to iterate over the loop again after we encounter the continue but by preserving the counter/iterator state values.</p>
<pre><code class="language-go">counter := 2
for counter &lt; 4 {
    counter++
    if counter &lt; 4 {
        continue
    }
    fmt.Println(&quot;Missed the Continue? at counter =&quot;, counter)
}
</code></pre>
<pre><code>$ go run infinite.go
Missed the Continue? at counter = 4
</code></pre>
<p>For following up with the code for this and all parts of the series, head over to the <a href="https://github.com/mr-destructive/100-days-of-golang">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So, from this section, we were able to understand the basics of conditional statements and loops in golang. We covered the things which are more important for understanding the core of the language than some specific things. There are certain parts that need to be explored further like iterating over arrays and slices which we'll cover after we have understood arrays and slices. Hopefully, you have understood the basics of the conditional statements and loops in golang. Thank you for reading, if you have any questions, or feedback please let me know in the comments or social handles. Until next time, Happy Coding :)</p>
