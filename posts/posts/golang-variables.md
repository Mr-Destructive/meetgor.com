{
  "title": "Golang: Variables and Types",
  "status": "published",
  "slug": "golang-variables",
  "date": "2025-04-05T12:33:44Z"
}

<h2>Introduction</h2>
<p>In the third part of the series, we will be covering the fundamentals for learning any programming language i.e. variables and data types. We will be covering from data types to variable declaration. We won't be seeing each and every detail related to the data types as some of them require a knowledge of loops and other topics, so that can be left for the different part.</p>
<h2>Types in golang</h2>
<p>In Golang there are 3 major types : Numeric, Bool and String. Further we also have specific types for the three data types like int, float, rune, byte, etc. We will first see how to declare a simple variable and then explore the data types in Golang.</p>
<pre><code class="language-go">var name string
</code></pre>
<p>This the variable declaration in Golang, we have the keyword <code>var</code> similar to Javascript but we optionally have to specify the type of the variable if we are not immediately assigning/defining it a value.</p>
<p>To assign variable values, we can write the datatype of the assigned value but it is optional as the go compiler will know the datatype according to the assigned value. Further you cannot change the type of that variable once it is initialized.</p>
<pre><code class="language-go">var name string = &quot;Gopher&quot;

OR 

var name string
name = &quot;Gopher&quot;

OR

var name = &quot;Gopher&quot;

</code></pre>
<p>We also have <code>const</code> for assigning constant values to a particular value set. You cannot change the value for a constant type, doing this will result in compile time error. An important property of <code>const</code> can be noted here, if you simply declare a <code>const</code> without initializing the value and try to access that constant, the go-compiler will throw an compilation error.</p>
<pre><code class="language-go">const name string = &quot;David&quot;

OR

const name string
name = &quot;Calvin&quot;

OR

const name = &quot;Smith&quot;
</code></pre>
<p>By default, the values for string is an empty string<code>&quot;&quot;</code>, for integer and float it is <code>0</code> and for bool it is <code>false</code>.</p>
<p>Each of these are valid declaration of variables in golang. Let's now dive into the data types and follow up with variable declaration in detail later.</p>
<p>| Numeric    | String | Bool |
|------------|--------|------|
|            |        |      |<br>
|  int       | string | bool |
|            |        |      |
|  float     |        |      |
|            |        |      |
|  complex   |        |      |
|            |        |      |
|  rune      |        |      |
|            |        |      |<br>
|  byte      |        |      |</p>
<h3>Numeric</h3>
<p>Let's first explore the <code>numeric</code> data types in golang as you have guessed correctly, we have <code>int</code> and <code>float</code> as distinct categories but further we also have fine grained storage types for both of the types.</p>
<h4>Integer</h4>
<p>In integers as well we have two categories <code>signed</code> and <code>unsigned</code>, we can basically store only positive integers in <code>unsigned</code> integers giving us an extra bit to play with.</p>
<p>For Integers, we have specific data storage types depending on the bits it can store like <code>int8</code> for storing 8 bit integers, <code>int16</code> for storing 16 bit integer value, <code>int32</code> and <code>int64</code> for the given number of bits in the integer. Similarly we will have these storage based integer values for unsigned integers like <code>uint8</code>, <code>uint16</code>, <code>uint32</code> and <code>uint64</code>. We can basically store double amount of positive integers in unsigned integers as <code>uint</code> than in signed integers <code>int</code>, this is because the most significant bit is not used as a sign bit since all values in the variable are positive and hence no sign bit is required.</p>
<pre><code class="language-go">var likes int = 140
fmt.Println(likes)
</code></pre>
<pre><code>$ go run int.go
140
</code></pre>
<pre><code class="language-go">var age int8 = 140
fmt.Println(&quot;Age = &quot;,age) 
</code></pre>
<pre><code>$ go run int.go
# command-line-arguments
.\int.go:6:9: constant 140 overflows int8
</code></pre>
<p>This will give us an error as <code>140</code> is above the limit for <code>int8</code>. So, unless you have specific requirements as storage limitation, you should be using <code>int</code> as the default data type for storing integers.</p>
<p>So, we need to define variables as per the limits to which we are going to use them, if you just specify <code>int</code> the type will be selected based on your operating system, if it is <code>32bit</code>, it will take <code>int32</code>, for <code>64bit</code> OSes it will take as <code>int64</code> integer. If you define a variable with let say <code>16</code> bit storage and if it exceeds the limit for <code>16</code> bit storage, it would give a <code>overflow limit</code> error.</p>
<p>Below are the limits for all the integer types in Golang:</p>
<pre><code>uint8 -&gt;  0  to  255
uint16 -&gt;  0  to  65535
uint32 -&gt;  0  to  4294967295
uint64 -&gt;  0  to  18446744073709551615

int8 -&gt;  -128  to  127
int16 -&gt;  -32768  to  32767
int32 -&gt;  -2147483648  to  2147483647
int64 -&gt;  -9223372036854775808  to  9223372036854775807
</code></pre>
<p>If you want to reality check for the boundary values of this data types, you can code a program in <code>go</code> as below:</p>
<ul>
<li>To find the maximum value of uint, we can flip all the bits in <code>0</code> to get all the set bits in the integer thus we use <code>^</code> operator.</li>
<li>To find the maximum value for signed integers, we can right shit one bit so as to unset the sign bit.</li>
<li>The minimum value for uint is the default value <code>0</code>.</li>
<li>The minimum value for int can be calculated by subtracting one from the negative of the max limit.</li>
</ul>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
)

func main() {
    var min_uint = 0
    var max_uint8 uint8 = ^uint8(0)
    var max_uint16 uint16 = ^uint16(0)
    var max_uint32 uint32 = ^uint32(0)
    var max_uint64 uint64 = ^uint64(0)

    var max_int8 int8 = int8(max_uint8&gt;&gt;1)
    var min_int8 int8 = -max_int8 - 1
    var max_int16 int16 = int16(max_uint16&gt;&gt;1)
    var min_int16 int16 = -max_int16 - 1
    var max_int32 int32 = int32(max_uint32&gt;&gt;1)
    var min_int32 int32 = -max_int32 - 1
    var max_int64 int64 = int64(max_uint64&gt;&gt;1)
    var min_int64 int64 = -max_int64 - 1

    fmt.Println(&quot;uint8 -&gt; &quot;, min_uint, &quot; to &quot;, max_uint8)
    fmt.Println(&quot;uint16 -&gt; &quot;, min_uint, &quot; to &quot;, max_uint16)
    fmt.Println(&quot;uint32 -&gt; &quot;, min_uint, &quot; to &quot;, max_uint32)
    fmt.Println(&quot;uint64 -&gt; &quot;, min_uint, &quot; to &quot;, max_uint64)
    fmt.Println(&quot;&quot;)
    fmt.Println(&quot;int8 -&gt; &quot;, min_int8, &quot; to &quot;, max_int8)
    fmt.Println(&quot;int16 -&gt; &quot;, min_int16, &quot; to &quot;, max_int16)
    fmt.Println(&quot;int32 -&gt; &quot;, min_int32, &quot; to &quot;, max_int32)
    fmt.Println(&quot;int64 -&gt; &quot;, min_int64, &quot; to &quot;, max_int64)
}
</code></pre>
<p>This was the basic overview of Integers in Golang.</p>
<p>Though rune and byte are Integer aliases, we will explore them in the String section as they make a lot of sense over there.</p>
<h4>Float</h4>
<p>Similar to integers, we also have floats in the numeric category. A float is a numeric data type that can allow numbers with decimal values. A simple float can be of either <code>float32</code> or <code>float64</code>. The two types are defined as the precision of the decimal values. Obliviously, the <code>float64</code> type is more precise than the counterpart and is also the default type assigned if not provided.</p>
<pre><code class="language-go">const percent = 30.5
fmt.Println(percent+50)
</code></pre>
<p>You optionally provide the <code>float32</code> type to have a bit less precision than usual using the <code>float32</code> keyword as follows:</p>
<pre><code class="language-go">const percent float32 = 46.34
fmt.Println(percent - 3.555)
</code></pre>
<p>The floating value precision of the float types in golang are as follows:</p>
<pre><code>float32	  --&gt;   -3.4e+38 to 3.4e+38.
float64	  --&gt;   -1.7e+308 to +1.7e+308.
</code></pre>
<p>As quite logical reasons, the precision is almost double in the <code>float64</code> compared to <code>float32</code>. If we try to add(any operation) a <code>float64</code> number with a <code>flaot32</code>, we get an error as performing operations on two differently stored types can't be operated.</p>
<h4>Complex Numbers</h4>
<p>We also have complex numbers in golang. This are the numbers which deal with a real and a imaginary numbers. We initialize the complex variable using the <code>complex</code> function which takes two parameters the <code>real</code> part and the <code>imagianry</code> part. Both the parts or numbers are stored as float in the complex data type.</p>
<p>So, since golang has <code>float32</code> and <code>float64</code> data types, we will have <code>complex64</code> and <code>complex128</code> as complex types. Since we are storing two <code>flaot64</code>, it has a <code>complex128</code> type and <code>complex64</code> for both parts as <code>float32</code>. Here as well, you cannot store the two parts(real and imaginary) as different float types i.e. You need to have both real and imaginary as either <code>flaot32</code> or <code>flaot64</code>.</p>
<pre><code class="language-go">var percent = 30.738
var f = 4.545
var comp1 = complex(f, percent)
var comp2 = complex(percent, f)
fmt.Println(comp1 - comp2)
</code></pre>
<pre><code>(-26.192999999999998+26.192999999999998i)
</code></pre>
<p>Golang automatically adds the <code>i</code> or iota in the complex/imaginary part for better readability.</p>
<h3>Strings</h3>
<p>We can now move onto the <code>string</code> data type in golang. It has several data types like <code>byte</code>, <code>rune</code>, <code>string</code>. In golang, <code>byte</code> and <code>rune</code> store individual characters whereas <code>string</code> can hold multiple characters.</p>
<h4>Byte</h4>
<p>A byte in golang is an unsigned 8 bit integer, which means it can hold numeric data from 0 to 255. So how is this displaying characters if it stores integer. Well, because each number it stores is mapped to the ASCII character set which is used to represent characters.</p>
<p>A byte can be stored in a single quote <code>''</code>, if we use double quotes<code>&quot;&quot;</code>, the variable is considered as string if we aren't specifying the data type.</p>
<pre><code class="language-go">const c byte = 't'
fmt.Println(c)
</code></pre>
<pre><code>$ go run byte.go
116
</code></pre>
<p>This gives the output as a number between 0 and 255 depending on the character which you have stored. To print the actual character you need to type caste into a string like:</p>
<pre><code class="language-go">const c byte = 't'
fmt.Printf(&quot;Character = %c 
Integer value = %d
&quot;, c, c)
</code></pre>
<pre><code>$ go run byte.go
Character = t
Integer Value = 116
</code></pre>
<p>We can get the character equivalent of the byte representation number using the <a href="https://cs.opensource.google/go/go/+/go1.18:src/fmt/print.go;l=212">Printf</a> function and the <code>%c</code> place holder for a character. The <code> </code> is used to end the line just for displaying proper output.</p>
<p>We can even store numbers between <code>0</code> and <code>255</code> as it is internally an <code>uint8</code>.</p>
<h4>Rune</h4>
<p>A rune is extended type of byte as it stores <code>int32</code> numbers and hence it represents <code>Unicode</code> characters. This is the default type if you do not specify <code>byte</code> and use single quotes to assign a character. Using rune, we can assign it an unicode characters to display some rich characters and literals like emoji or expressions.</p>
<pre><code class="language-go">var smiley_emoji = '☺'
fmt.Printf(&quot;Smiley Emoji --&gt; %c&quot;, smiley_emoji)
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1648962460/blog-media/obw9ihlxsvhytbe8ito3.png" alt="GO Rune Smiley Emoji"></p>
<p>So, rune is pretty amazing type to play with characters in golang. As it is a default type assigned against byte if not provided while assignment.</p>
<h4>String</h4>
<p>Strings are basically a slice(list) of bytes. Each character in a string is a byte. By default the string will be empty if you don't initialize it with a value.</p>
<pre><code class="language-go">const language string
language = &quot;C++&quot;

OR

const language string= &quot;Python&quot;

OR

const language = &quot;Javascript&quot;
</code></pre>
<p>We can even access particular character in the string using it's index.</p>
<pre><code class="language-go">const code = &quot;12AB34CD&quot;
fmt.Println(code[6])
</code></pre>
<pre><code>$ go run string.go
67
</code></pre>
<p>This gives us a integer as we are accessing the byte from the string using its index. Thus, we can use <code>%c</code> in the <code>Printf</code> function to format and print the equivalent characters of the byte.</p>
<pre><code class="language-go">const code = &quot;12AB34CD&quot;
fmt.Printf(&quot;2nd Character in string = %c
&quot;, code[4])
</code></pre>
<pre><code>$ go run string.go
2nd Character in string = A
</code></pre>
<p>We can also declare strings using backticks/backquotes or whatever you call it (```), assigning string with this allows us to write multi line string.</p>
<pre><code class="language-go">var statement = `This is the first line
The next line
The last line`

fmt.Println(statement)
</code></pre>
<pre><code>$ go run str-backticks.go
This is the first line
The next line
The last line
</code></pre>
<p>Further in the loop article we will see how to loop/iterate over a string.</p>
<h3>Boolean</h3>
<p>This type is used to store either <code>true</code> or <code>false</code> in golang. The default value of a boolean variable is <code>false</code>.</p>
<pre><code class="language-go">var power bool
fmt.Println(power)
</code></pre>
<pre><code>$ go run bool.go
false
</code></pre>
<p>We can assign the variable as either <code>true</code> or <code>false</code>.</p>
<pre><code class="language-go">const result = true
fmt.Printf(&quot;The statement is %t&quot;, result)
</code></pre>
<pre><code>$ go run bool.go
The statement is true
</code></pre>
<p>So, using the <code>%t</code> we can print the value of a boolean value in golang in the <code>Printf</code> function.</p>
<h2>Creating Variables</h2>
<p>Now, we have familiar with data types in golang, we can more expressively create, declare, initialize variables in golang.</p>
<p>There are 3-4 primary ways to define a variable most of which we have already seen.</p>
<h3>Declaring a Variable</h3>
<p>We can declare a variable without assigning it any value but for that we need to then provide the data type, this can be done using the following command:</p>
<pre><code class="language-go">var expereience int

expereience = 2
</code></pre>
<p>We can even use <code>const</code> for constant value in the given scope.</p>
<p>Here, we can even declare multiple variables by separating each variable/constant with comma <code>,</code> which can be done as follows:</p>
<pre><code class="language-go">var a, b, c int

OR

const i, j, k bool
</code></pre>
<h3>Defining and Initializing at the same time</h3>
<p>We can initialize a variable/constant in golang by explicitly giving it a value. We can do that by using <code>var</code> for variable value or <code>const</code> for a constant value. We can optionally provide the data type at this moment as golang will automatically detect the type and assign the memory according to the value given.</p>
<pre><code class="language-go">var place string = &quot;home&quot;
</code></pre>
<p>Here, there is no compulsion to provide the <code>datatype</code> as the compiler will be able to know it from the asisgned value. Though if you want to provide a non-default value you can specify the datatype.</p>
<h3>Declaring Multiple Variables</h3>
<p>We can assign multiple variables at once by separating them with comma<code>,</code>. The variable name to the left and the values to the right needs to separated with comm on both sides.</p>
<pre><code class="language-go">var x, y, z = 100, '#', &quot;days&quot;

fmt.Printf(&quot; x = %d 
 y = %c 
 z = %s 
&quot;,x,y,z)
</code></pre>
<pre><code>$ go run multiplvar.go
 x = 100
 y = #
 z = daysofcode
</code></pre>
<p>We can are declaring and assigning multiple variables, the <code>x</code> variable is assigned an integer value, <code>y</code> with a <code>rune</code>(by default) and <code>z</code> with a string. We are using <code>Printf</code> function with place holders for int <code>%d</code>, rune/byte <code>%c</code> and string as <code>%s</code>. The <code> </code> is for a new line.</p>
<p>If we want to assign the variables with a particular data type, we can use the var keyword as a list of values.</p>
<pre><code class="language-go">var (
    x int8 = 100
    y byte = '#'
    z =  &quot;daysofcode&quot;
)

fmt.Printf(&quot; x = %T 
 y = %T 
 z = %T 
&quot;,x,y,z)
</code></pre>
<pre><code>$ go run multiplvar.go
 x = int8
 y = uint8
 z = string
</code></pre>
<p>This is not only limited to <code>var</code> we can also use <code>const</code> to declare multiple constants with type constraint. Also, note we are using the <code>%T</code> placeholder for getting the type of the data stored in the variable.</p>
<p>Also, we can define(declare and initialize) multiple variable with same data type with command separated as follows:</p>
<pre><code class="language-go">var pi, e, G float32 = 3.141, 2.718, 6.67e-11   
var start, end byte = 65, 90
fmt.Println(pi, e, G)
fmt.Printf(&quot;%c %c
&quot;,start, end)
</code></pre>
<pre><code>$ go run multp.go
3.141 2.718 6.67e-11
A Z
</code></pre>
<h3>Assigning Variable using Walrus Operator (Shorthand Declaration)</h3>
<p>We can skip usign <code>var</code> or the <code>datatype</code> by using the <code>:=</code> walrus operator. This type of assignment using <code>walruns</code> operator can only be allowed in the function body and not anywhere else, in the global scope this type of assignment is not allowed.</p>
<pre><code class="language-go">place := &quot;school&quot;
</code></pre>
<p>This is such a simple shorthand for assigning variables though only in a function body.</p>
<p>Also, multiple variable declaration is possible with walrus operator.</p>
<pre><code class="language-go">x, y, z := &quot;foo&quot;, 32, true
fmt.Println(x, y, z)
fmt.Printf(&quot;%T %T %T&quot;, x, y, z)
</code></pre>
<pre><code class="language-shell">$ go run walrus.go
foo 32 true
string int bool
</code></pre>
<p>Links to all code and links are visible on the <a href="https://github.com/Mr-Destructive/100-days-of-golang">GitHub</a> repository.</p>
<h2>Conclusion</h2>
<p>So, from this part of the series, we were able to understand variables and the various data types in Golang. Though we didn't got too much in detail still we can find ourselves a bit comfortable in understanding basic go scripts. In the next section, we will looking into conditional statements and loops. This would give a good grasp on iterating over a string and even learn arrays(just the basics) we will explore Arrays and slices(remember strings?) after that.</p>
<p>So, if you have any questions, suggestions, or feedback please let me know in the comments or on the social handles. See you next time, Happy Coding :)</p>
