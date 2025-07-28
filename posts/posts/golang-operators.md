{
  "title": "Golang: Operators",
  "status": "published",
  "slug": "golang-operators",
  "date": "2025-04-05T12:33:41Z"
}

<h2>Introduction</h2>
<p>In this 13th part of the series, we will be exploring the fundamentals of operators in Golang. We will be exploring the basics of operators and the various types like Arithmetic, Bitwise, Comparison, Assignment operators in Golang.</p>
<p>Operators are quite fundamentals in any programming language. Operators are basically expressions or a set of character(s) to perform certain fundamental tasks. They allow us to perform certain trivial operations with a simple expression or character. There are quite a few operators in Golang to perform various operations.</p>
<h2>Types of Operators</h2>
<p>Golang has a few types of operators, each type providing particular aspect of forming expressions and evaluate conditions.</p>
<ol>
<li>Bitwise Operators</li>
<li>Logical Operators</li>
<li>Arithmetic Operators</li>
<li>Assignment Operators</li>
<li>Comparison Operators</li>
</ol>
<h3>Bitwise Operators</h3>
<p>Bitwise Operators are used in performing operations on binary numbers. We can perform operation on a bit level and hence they are known as bitwise operators. Some fundamental bitwise operators include, <code>AND</code>, <code>OR</code>, <code>NOT</code>, and <code>EXOR</code>. Using this operators, the bits in the operands can be manipulated and certain logical operations can be performed.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
	x := 3
	y := 5
	// 3 -&gt; 011
	// 5 -&gt; 101
	fmt.Println(&quot;X AND Y = &quot;, x &amp; y)
	fmt.Println(&quot;X OR Y = &quot;, x | y)
	fmt.Println(&quot;X EXOR Y = &quot;, x ^ y)
	fmt.Println(&quot;X Right Shift 1  = &quot;, x &gt;&gt; 1)
	fmt.Println(&quot;X Right Shift 2  = &quot;, x &gt;&gt; 2)
	fmt.Println(&quot;Y Left Shift 1 = &quot;, y &lt;&lt; 1)
}
</code></pre>
<pre><code>$ go run bitwise/main.go

X AND Y =  1
X OR Y =  7
X EXOR Y =  6
X Right Shift 1  =  1
X Right Shift 2  =  0
Y Left Shift 1 =  10

</code></pre>
<p>We use the <code>&amp;</code> (AND operator) for performing AND operations on two operands. Here we are logically ANDing <code>3</code> and <code>5</code> i.e. <code>011</code> with <code>101</code> so it becomes <code>001</code> in binary or 1 in decimal.</p>
<p>Also, the <code>|</code> (OR operator) for performing logical OR operation on two operands. Here we are logically ORing <code>3</code> and <code>5</code> i.e. <code>011</code> with <code>101</code> so it becomes <code>111</code> in binary or 7 in decimal.</p>
<p>Also the <code>^</code> (EXOR operator) for performing logical EXOR operation on two operands. Here we are logically EXORing <code>3</code> and <code>5</code> i.e. <code>011</code> with <code>101</code> so it becomes <code>110</code> in binary or 6 in decimal.</p>
<p>We have a couple of more bitwise operators that allow us to shift bits in the binary representation of the number. We have two types of these shift operators, right sift and left shift operators. The main function of these operator is to shift a bit in either right or left direction.</p>
<p>In the above example, we have shifted <code>3</code> i.e. <code>011</code> to right by one bit so it becomes <code>001</code>. If we would have given <code>x &gt;&gt; 2</code> it would have become <code>0</code> since the last bit was shifted to right and hence all bits were 0.</p>
<p>Similarly, the left shift operator sifts the bits in the binary representation of the number to the left. So, in the example above, <code>5</code> i.e. <code>101</code> is shifted left by one bit so it becomes <code>1010</code> in binary i.e. 10 in decimal.</p>
<p>This was a basic overview of bitwise operators in Golang. We can use these basic operators to perform low level operations on numbers.</p>
<h3>Comparison Operators</h3>
<p>This type of operators are quite important and widely used as they form the fundamentals of comparison of variables and forming boolean expressions. The comparison operator is used to compare two values or expressions.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
	a := 45
	b := 12
	fmt.Println(&quot;Is A equal to B ? &quot;, a == b)
	fmt.Println(&quot;Is A not equal to B ? &quot;, a != b)
	fmt.Println(&quot;Is A greater than B ? &quot;, a &gt; b)
	fmt.Println(&quot;Is A less than B ? &quot;, a &lt; b)
	fmt.Println(&quot;Is A greater than or equal to B ? &quot;, a &gt;= b)
	fmt.Println(&quot;Is A less than or equal to B ? &quot;, a &lt;= b)

</code></pre>
<pre><code>$ go run comparison/main.go

Is A equal to B ?  false
Is A not equal to B ?  true
Is A greater than B ?  true
Is A less than B ?  false
Is A greater than or equal to B ?  true
Is A less than or equal to B ?  false
</code></pre>
<p>We use simple comparison operators like <code>==</code> or <code>!=</code> for comparing if two values are equal or not. The expression <code>a == b</code> will evaluate to <code>true</code> if the values of both variables or operands are equal. However, the expression <code>a != b</code> will evaluate to <code>true</code> if the values of both variables or operands are not equal.</p>
<p>Similarly, we have the <code>&lt;</code> and <code>&gt;</code> operators which allow us to evaluate expression by comparing if the values are less than or grater than the other operand. So, the expression <code>a &gt; b</code> will evaluate to <code>true</code> if the value of <code>a</code> is greater than the value of <code>b</code>. Also the expression <code>a &lt; b</code> will evaluate to <code>true</code> if the value of <code>a</code> is less than the value of <code>b</code>.</p>
<p>Finally, the operators <code>&lt;=</code> and <code>&gt;=</code> allow us to evaluate expression by comparing if the values are less than or equal to and greater than or equal to the other operand. So, the expression <code>a &gt;= b</code> will evaluate to <code>true</code> if the value of <code>a</code> is greater than or if it is equal to the value of <code>b</code>, else it would evaluate to <code>false</code>. Similarly, the expression <code>a &lt;= b</code> will evaluate to <code>true</code> if the value of <code>a</code> is less than or if it is equal to the value of <code>b</code>, else it would evaluate to <code>false</code>.</p>
<p>These was a basic overview of comparison operators in golang.</p>
<h3>Logical Operators</h3>
<p>Next, we move on to the logical operators in Golang which allow to perform logical operations like <code>AND</code>, <code>OR</code>, and <code>NOT</code> with conditional statements or storing boolean expressions.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
	a := 45
	b := &quot;Something&quot;
	fmt.Println(a &gt; 40 &amp;&amp; b == &quot;Something&quot;)
	fmt.Println(a &lt; 40 &amp;&amp; b == &quot;Something&quot;)
	fmt.Println(a &lt; 40 || b == &quot;Something&quot;)
	fmt.Println(a &lt; 40 || b != &quot;Something&quot;)
	fmt.Println(!(a &lt; 40 || b != &quot;Something&quot;))
}
</code></pre>
<pre><code>$ go run logical/main.go

true
false
true
false
true
</code></pre>
<p>Here, we have used logical operators like <code>&amp;&amp;</code> for Logical AND, <code>||</code> for logical OR, and <code>!</code> for complementing the evaluated result. The <code>&amp;&amp;</code> operation only evaluates to <code>true</code> if both the expressions are <code>true</code> and <code>||</code> OR operator evaluates to <code>true</code> if either or both the expressions are <code>true</code>. The <code>!</code> operator is used to complement the evaluated expression from the preceding parenthesis.</p>
<h3>Arithmetic Operators</h3>
<p>Arithmetic operators are used for performing Arithmetic operations. We have few basic arithmetic operators like <code>+</code>, <code>-</code>, <code>*</code>, <code>/</code>, and <code>%</code> for adding, subtracting, multiplication, division, and modulus operation in golang.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
	a := 30
	b := 50
	fmt.Println(&quot;A + B = &quot;, a+b)
	fmt.Println(&quot;A - B = &quot;, a-b)
	fmt.Println(&quot;A * B = &quot;, a*b)
	fmt.Println(&quot;A / B = &quot;, a/b)
	fmt.Println(&quot;A % B = &quot;, a%b)
}
</code></pre>
<pre><code>$ go run arithmetic/main.go
A + B =  80
A - B =  -20
A * B =  1500
A / B =  0
A % B =  30
</code></pre>
<p>These are the basic mathematical operators in any programming language. We can use <code>+</code> to add two values, <code>-</code> to subtract two values, <code>*</code> to multiply to values, <code>/</code> for division of two values and finally <code>%</code> to get the remainder of a division of two values i.e. if we divide 30 by 50, the remainder is 30 and the quotient is 0.</p>
<p>We also have a few other operators like <code>++</code> and <code>--</code> that help in incrementing and decrementing values by a unit value. Let's say we have a variable <code>k</code> which is set to <code>4</code> and we want to increment it by one, so we can definitely use <code>k = k + 1</code> but it looks kind of too long, we have a short notation for the same <code>k++</code> to do the same.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
	k := 3
	j := 20
	fmt.Println(&quot;k = &quot;, k)
	fmt.Println(&quot;j = &quot;, j)
	k++
	j--
	fmt.Println(&quot;k = &quot;, k)
	fmt.Println(&quot;j = &quot;, j)
}
</code></pre>
<pre><code>$ go run arithmetic/main.go

k =  3
j =  20

k =  4
j =  19
</code></pre>
<p>So, we can see that the variable <code>k</code> is incremented by one and variable <code>j</code> is decremented by <code>1</code> using the <code>++</code> and <code>--</code> operator.</p>
<h3>Assignment Operators</h3>
<p>These types of operators are quite handy and can condense down large operations into simple expressions. These types of operators allow us to perform operation on the same operand. Let's say we have the variable <code>k</code> set to <code>20</code> initially, we want to add <code>30</code> to the variable <code>k</code>, we can do that by using <code>k = k + 30</code> but a more sophisticated way would be to use <code>k += 30</code> which adds <code>30</code> or any value provided the same variable assigned and operated on.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
	var a int = 100
	b := 20
	fmt.Println(&quot;a = &quot;, a)
	fmt.Println(&quot;b = &quot;, b)
	a += 30
	fmt.Println(&quot;a = &quot;, a)
	b -= 5
	fmt.Println(&quot;b = &quot;, b)
	a *= b
	fmt.Println(&quot;a = &quot;, a)
	fmt.Println(&quot;b = &quot;, b)
	a /= b
	fmt.Println(&quot;a = &quot;, a)
	fmt.Println(&quot;b = &quot;, b)
	a %= b
	fmt.Println(&quot;a = &quot;, a)
	fmt.Println(&quot;b = &quot;, b)
}
</code></pre>
<pre><code>$ go run assignment/main.go

a =  100
b =  20

a =  130
b =  15

a =  1950
b =  15

a =  130
b =  15

a =  10
b =  15
</code></pre>
<p>From the above example, we are able to perform operations by using shorthand notations like <code>+=</code> to add the value to the same operand. These also saves a bit of time and memory not much but considerable enough. This allow us to directly access and modify the contents of the provided operand in the register rather than assigning different registers and performing the operations.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So, from the following part of the series, we were able to learn the basics of operators in golang. Using some simple and easy to understand examples, we were able to explore different types of operators like arithmetic, logical, assignment and bitwise operators in golang. These are quite fundamental in programming in general, this lays a good foundation for working with larger and complex projects that deal with any kind of logic in it, without a doubt almost all of the applications do have a bit of logic attached to it. So, we need to know the basics of operators in golang.</p>
