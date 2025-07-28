{
  "title": "Golang: String Manipulation",
  "status": "published",
  "slug": "golang-strings",
  "date": "2025-04-05T12:33:39Z"
}

<h2>Introduction</h2>
<p>In the 15th post of the Series, we will be looking into the details of the String manipulation and performing types of operations in Golang. We will explore string manipulation, concatenation, helper functions, etc. which will help in working with strings in Golang.</p>
<h2>String Concatenation</h2>
<p>String Concatenation refers to the combining and formatting of strings in Golang. We can combine multiple strings and formating the way we display the strings in Golang. We have a few ways and functions to concatenate strings in Golang.</p>
<h3>Using the + operator</h3>
<p>We can simply concatenate strings using the <code>+</code> operator, though keep in mind you should only concatenate the string with a string and not any other data type like integer, or float, it will throw out errors for mismatched string types.</p>
<pre><code class="language-go">package main

import (
    &quot;fmt&quot;
)
func main() {
    s1 := &quot;Go&quot;
    s2 := &quot;Programming&quot;
    s3 := &quot;Language&quot;

    s := s1 + s2 + s3
    fmt.Println(s)
}
</code></pre>
<pre><code>$ go run concatenate.go

GoProgrammingLanguage

</code></pre>
<p>The <code>+</code> operator will literally join the strings as it is and form a string.</p>
<h3>Using the += operator</h3>
<p>The other way to continuously append a string to an existing string, we can use the <code>+=</code> operator. This operator will append the provided string to the end of the original string.</p>
<pre><code class="language-go">p := &quot;Story&quot;
p += &quot;Book&quot;
fmt.Println(p)
</code></pre>
<pre><code>go run concatenate.go

StoryBook
</code></pre>
<h3>Using the Join method</h3>
<p>The join method is a function available in the string package in Golang. We can join strings elements in a slice or an array using the <a href="https://pkg.go.dev/strings#Join">Join</a> method in the strings <a href="https://pkg.go.dev/strings">package</a> in golang. The Join method will combine all the elements in between the elements with a particular string. So, the function takes two parameters <code>Join(array, string)</code>, the array or a slice is parsed into the function which will be used to insert the provided string in between the elements of the slice.</p>
<h4>Join</h4>
<ul>
<li>Parameters   : array/slice, string</li>
<li>Return Value : string</li>
</ul>
<pre><code class="language-go">q := []string{&quot;meetgor.com&quot;, &quot;tags&quot;, &quot;golang&quot;, &quot;string&quot;}
r := strings.Join(q, &quot;/&quot;)
fmt.Println(r)
</code></pre>
<pre><code>go run concatenate.go

meetgor.com/tags/golang/string
</code></pre>
<p>In the above example, we use have used the <code>Join</code> method to insert a string in between the elements of a slice of strings. The string <code>&quot;/&quot;</code> has been inserted in between the elements, and the elements are combined as a single string. So, each individual element starting from the <code>0</code> index <code>meetgor.com</code> is appended the string <code>/</code> and further the next element <code>tags</code> have been appended and the procedure caries on till the last element. Note that the string is not inserted after the last element. The function <code>Join</code> returns a string and thereby we store the string in a variable.</p>
<h3>Using Sprintf method</h3>
<p>We can use the <a href="https://pkg.go.dev/fmt#Sprintf">Sprintf</a> function from the fmt package to format the string by storing the string rather than printing it to the console. The sprintf function is quite similar to the <code>Printf</code> but it only parses strings rather than printing them directly to the console.</p>
<pre><code class="language-go">// Using Sprintf function to format strings

name := &quot;peter&quot;
domain := &quot;telecom&quot;
service := &quot;ceo&quot;

email := fmt.Sprintf(&quot;%s.%s@%s.com&quot;, service, name, domain)
fmt.Println(email)
</code></pre>
<pre><code>go run concatenate.go

ceo.peter@telecom.com
</code></pre>
<p>The sprintf function basically allows us to concatenate strings in a defined format just like we use <code>printf</code> to print formatted strings. In the above example, we have formatted three strings in the form of an email by assigning a placeholder for the string i.e. <code>%s</code>, and adding the required characters in the formatted string.</p>
<h3>Using Go string Builder method</h3>
<p>The <a href="https://pkg.go.dev/strings#Builder">Builder</a> type is provided by the strings package in Golang. The Builder type helps in building strings in an efficient way. By creating a string builder object, we can perform operations on a String.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;strings&quot;
)

func main() {
  // Using Builder function

  c := []string{&quot;j&quot;, &quot;a&quot;, &quot;v&quot;, &quot;a&quot;}
  var builder strings.Builder
  for _, item := range c {
    builder.WriteString(item)
  }
  fmt.Println(&quot;builder = &quot;, builder.String())
  b := []byte{'s', 'c', 'r', 'i', 'p', 't'}
  for _, item := range b {
    builder.WriteByte(item)
  }
  fmt.Println(&quot;builder = &quot;, builder.String())
  builder.WriteRune('s')
  fmt.Println(&quot;builder = &quot;, builder.String())
  fmt.Println(&quot;builder = &quot;, builder)
}
</code></pre>
<pre><code>go run concatenate.go

builder =  java
builder =  javascript
builder =  javascripts
builder =  {0xc000088dd8 [106 97 118 97 115 99 114 105 112 116 115]}
</code></pre>
<p>The builder structure provided by the strings package is quite important for working with strings in an efficient manner. Its usually used for string concatenation operations. We can perform write operations to the buffer which is a byte slice.  Here we have created the <code>builder</code> variable which is of type <code>strings.Builder</code>, further we have appended the string to it in a for a loop. So, we construct a string from the string list elements, they can be even rune slice or byte slice.
We have used three methods here, the <code>WriteString</code>, <code>WriteByte</code>, and <code>WriteRune</code> which are quite obliviously used for writing <code>string</code>, <code>byte</code>, and <code>rune</code>to the string builder object.</p>
<h3>Using the Bytes buffer method</h3>
<p>The <code>bytes</code> package also has something similar to <code>Builder</code> type in <code>string</code> as <a href="https://pkg.go.dev/bytes#Buffer.Bytes">Buffer</a>. It has almost the same set of methods and properties. The main difference is the efficiency, <code>strings.Builder</code> is comparatively faster than the <code>bytes.Buffer</code> type due to several low-level implementations. We can discuss those fine details in a separate article but right now we'll focus on the ways we can utilize this type for string concatenation.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;bytes&quot;
)

func main() {
	// Using bytes buffer method

	var buf bytes.Buffer

	for i := 0; i &lt; 2; i++ {
		buf.WriteString(&quot;ja&quot;)
	}
	fmt.Println(buf.String())

	buf.WriteByte('r')

	fmt.Println(buf.String())

	k := []rune{'s', 'c', 'r', 'i', 'p', 't'}
	for _, item := range k {
		buf.WriteRune(item)
	}
	fmt.Println(buf.String())
}
</code></pre>
<pre><code>go run concatenate.go

jaja
jajar
jajarscript
{[106 97 106 97 114 115 99 114 105 112 116] 0 0}
</code></pre>
<p>So, like for the <code>strings.Builder</code> type, we have <a href="https://pkg.go.dev/bytes#Buffer.WriteString">WriteString</a>, <a href="https://pkg.go.dev/bytes#Buffer.WriteByte">WriteByte</a>, and <a href="https://pkg.go.dev/bytes#Buffer.WriteRune">WriteRune</a> in the <code>bytes.Buffer</code> type. We can use it exactly the way we do with the previous example. Also, the <code>bytes.Buffer</code> type returns a slice of bytes so we will have to use the <a href="https://pkg.go.dev/strings#Builder.String">String()</a> method to format it as a string.</p>
<p>If we look at the <a href="https://go.dev/src/bytes/buffer.go">bytes.Buffer</a> type, it returns a slice of bytes and two more properties viz. <code>off</code> and <code>lastRead</code>. These two properties are used for indicating the index of the byte in the buffer and reallocation of the buffer. This is too low-level stuff that can be explored and explained in a separate section. For further readings on the bytes Buffer or String Builder types, you can follow up with these articles:</p>
<ul>
<li><a href="https://medium.com/@felipedutratine/string-concatenation-in-golang-since-1-10-bytes-buffer-vs-strings-builder-2b3081848c45">bytes.Buffer vs strings.Builder</a></li>
<li><a href="https://syslog.ravelin.com/bytes-buffer-i-thought-you-were-my-friend-4148fd001229">Bytes Buffer</a></li>
<li><a href="https://stackoverflow.com/a/49295215">Best ways to use bytes.Buffer</a></li>
</ul>
<h2>String Comparison</h2>
<p>Now, we can move into the comparison of Strings in Golang. We have quite a few ways to compare strings in golang. We cover some of them in this section.</p>
<h3>Using Comparison operators</h3>
<p>The basic comparison can be done with the comparison operators provided by Golang. Just like we compare numeric data we can compare strings. Though the factor with which we compare them is different. We compare them by the lexical order of the string characters.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
	s1 := &quot;gopher&quot;
	s2 := &quot;Gopher&quot;
	s3 := &quot;gopher&quot;

	isEqual := s1 == s2

  //&quot;gopher&quot; is not same as &quot;Gopher&quot; and hence `false`
	fmt.Printf(&quot;S1 and S2 are Equal? %t 
&quot;, isEqual)
	fmt.Println(s1 == s2)

  // &quot;gopher&quot; is not equal to &quot;Gopher&quot; and hence `true`
	fmt.Println(s1 != s2)

  // &quot;Gopher&quot; comes first lexicographically than &quot;gopher&quot; so return true 
  // G -&gt; 71 in ASCII and g -&gt; 103
	fmt.Println(s2 &lt; s3)
	fmt.Println(s2 &lt;= s3)

  // &quot;Gopher&quot; is not greater than &quot;gopher&quot; as `G` comes first in ASCII table
  // So G value is less than g i.e. 71 &gt; 103 which is false
	fmt.Println(s2 &gt; s3)
	fmt.Println(s2 &gt;= s3)

}
</code></pre>
<pre><code>go run comparison.go

S1 and S2 are Equal? false 
false
true
true
true
false
false
</code></pre>
<p>In the above examples, we are able to see the comparison of two strings. There are three strings, two of which are identical, and the third is identical as well but is not equal considering the case of the characters in the string. We have compared the strings in order of the ASCII value of the characters of the strings. For example, A (65) comes before a (97). Similarly, numbers come before letters. So accordingly the comparison of these string characters decides the result.</p>
<p>For the ASCII table, you can take a look over the below image:</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1656423519/blog-media/ascii_table.png" alt="ASCII Table"></p>
<h3>Using Compare method</h3>
<p>We also have the <a href="https://res.cloudinary.com/techstructive-blog/image/upload/v1656423519/blog-media/ascii_table.png">Compare</a> method in the strings package for comparing two strings. The comparison method returns an integer value of either -1, 0, or 1. If the two strings are equal, it will return 0. Else if the first string is lexicographically smaller than the second string, it will return -1, else it will return +1.</p>
<h4>strings.Compare</h4>
<ul>
<li>Return Type: Int (-1, 0, 1)</li>
<li>Parameters: string, string</li>
</ul>
<p>You can check out the <a href="https://go.dev/src/strings/compare.go">source code</a> for further clarity.</p>
<pre><code class="language-go">package main

import(
  &quot;fmt&quot;
  &quot;strings&quot;
)

func main() {
	s1 := &quot;gopher&quot;
	s2 := &quot;Gopher&quot;
	s3 := &quot;gopher&quot;

	fmt.Println(strings.Compare(s1, s2))
	fmt.Println(strings.Compare(s1, s3))
	fmt.Println(strings.Compare(s2, s3))
}
</code></pre>
<pre><code>go run comparison.go

1
0
-1
</code></pre>
<p>In the above example, the two strings <code>s1</code> and <code>s2</code> are compared and it returns the integer value <code>+1</code>, indicating the first string is lexicographically greater than the second string which is true <code>&quot;gopher&quot;</code> will be lexicographically after <code>&quot;Gopher&quot;</code> due to the presence of <code>G</code>.</p>
<p>In the second example, we are comparing the strings <code>s1</code> and <code>s3</code> which are equal, and hence the function returns <code>0</code> as expected.</p>
<p>In the third example, we are comparing the strings <code>s2</code> and <code>s3</code> identical to the first case but here order matters. We are comparing <code>&quot;Gopher&quot;</code> with <code>&quot;gopher&quot;</code> so the first string is lexicographically smaller than the second string and thereby returning <code>-1</code> as discussed.</p>
<h3>Using strings EqualFold</h3>
<p>We also have another method in the strings library called <a href="https://pkg.go.dev/strings#EqualFold">EqualFold</a> which compares two strings lexicographically but without considering the case precedence. That is the upper case or lower case is ignored and considered equal. So we are truly case-insensitively comparing the strings.</p>
<h4>strings.EqualFold</h4>
<ul>
<li>Return Type: bool (true or false)</li>
<li>Parameters: string, string</li>
</ul>
<pre><code class="language-go">package main

import(
  &quot;fmt&quot;
  &quot;strings&quot;
)

func main() {

	s1 := &quot;gopher&quot;
	s2 := &quot;Gopher&quot;
	s3 := &quot;gophy&quot;

	fmt.Println(strings.EqualFold(s1, s2))
	fmt.Println(strings.EqualFold(s1, s3))
	fmt.Println(strings.EqualFold(s2, s3))
}
</code></pre>
<pre><code>go run comparison.go

true
false
false
</code></pre>
<p>So, in the above example, we are comparing the strings <code>&quot;gopher&quot;</code> and <code>&quot;Gopher&quot;</code> i.e. <code>s1</code> and <code>s2</code>, which are equal if we think case-insensitively. Hence the method returns true, they are equal.
In the next example, we compare the strings, <code>s1</code> and <code>s3</code> i.e. <code>&quot;gopher&quot;</code> <code>&quot;gophy&quot;</code> which are not equal, and hence we return <code>false</code>. Similar is the case for <code>&quot;Gopher&quot;</code> and <code>&quot;gophy&quot;</code> which is <code>false</code>. Also, if we consider two strings <code>&quot;gophy&quot;</code> and <code>&quot;gophy&quot;</code> it will quite obliviously return <code>true</code>.</p>
<h2>String Manipulation and utility methods</h2>
<p>The strings package in golang has some great utility methods for working with string or any form of text. We will explore some of the quite useful and widely used utilities in the strings package.</p>
<h3>ToLower, ToUpper and Title functions</h3>
<p>The strings package also provides some utility functions for operating on the case of the characters in the strings. We have functions like <a href="https://pkg.go.dev/strings#ToLower">ToLower</a>, <a href="https://pkg.go.dev/strings#ToUpper">ToUpper</a>, and <a href="https://pkg.go.dev/golang.org/x/text/cases#Title">Title</a> which can be used to convert the string into lower case, uppercased or Capitalised(Title) cased characters respectively.</p>
<h4>strings.ToLower</h4>
<ul>
<li>Return Type: string</li>
<li>Parameters: string</li>
</ul>
<h4>strings.ToUpper</h4>
<ul>
<li>Return Type: string</li>
<li>Parameters: string</li>
</ul>
<h4>cases.Title</h4>
<ul>
<li>Return Type: Caser</li>
<li>Parameters: Language Options</li>
</ul>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;strings&quot;

	&quot;golang.org/x/text/cases&quot;
	&quot;golang.org/x/text/language&quot;
)

func main() {
	s1 := &quot;Ubuntu 22&quot;
	s2 := &quot;meet&quot;
	s3 := &quot;IND&quot;
	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToLower(s2))
	fmt.Println(strings.ToLower(s3))

	fmt.Printf(&quot;
&quot;)
	fmt.Println(strings.ToUpper(s1))
	fmt.Println(strings.ToUpper(s2))
	fmt.Println(strings.ToUpper(s3))

	fmt.Printf(&quot;
&quot;)
	cases := cases.Title(language.English)
	fmt.Println(cases.String(s1))
	fmt.Println(cases.String(s2))
	fmt.Println(cases.String(s3))
}
</code></pre>
<pre><code># 100-days-of-golang/scripts/strings

go mod init
go get golang.org/x/text/cases
go get golang.org/x/text/language

go run utility.go                                                                                             
ubuntu 22
meet
ind

UBUNTU 22
MEET
IND

Ubuntu 22
Meet
Ind
</code></pre>
<p>Here, we can see that the function, <code>ToLower</code> has converted all the characters of a string to the lower case of the alphabet. Similarly, the <code>ToUpper</code> function has turned the characters of the strings to their respective alphabetical upper case.</p>
<p>The <a href="https://pkg.go.dev/strings#Title">Title</a> method in the strings package has been deprecated due to incompatibility with certain languages and cases. So, we are using the <a href="https://pkg.go.dev/golang.org/x/text/cases">text/cases</a> package to get the <a href="https://pkg.go.dev/strings#Title">Title</a> method that appropriately converts a string to Title cased.
To set up this function, you need to perform a certain package installation process which is quite straightforward. Just create a go mod which is used for managing dependencies and packages for a project. So run the commands given below in the same order in your local setup:</p>
<pre><code>go mod init
go get golang.org/x/text/cases
go get golang.org/x/text/language
</code></pre>
<p>This will set up a go.mod file and install the packages namely the <code>cases</code> and <code>language</code> packages. After doing this you will be able to access the functions <code>Title</code> from the cases package which can be imported by the format <code>&quot;golang.org/x/text/cases&quot;</code> and <code>&quot;golang.org/x/text/language&quot;</code>. Now, we can use the Title function and parse the parameters which is the language type. Here we have used the <code>language.English</code> which is a <a href="https://pkg.go.dev/golang.org/x/text@v0.3.7/language#Tag">language Tag</a> to say use the semantics of English language while parsing the title cased characters. We now assign the value of the function <code>Title</code> to a variable as it will be of type <code>Caseer</code> and we want to still parse the string into the function. The <a href="https://pkg.go.dev/golang.org/x/text/cases#Caser">caser</a> object will have certain methods and properties attached to it, we will use the method <a href="https://pkg.go.dev/golang.org/x/text/cases#Caser.String">Strings</a> that will convert the given string into the title cased string. Hence we return the title cased string using the title function with the help of cases and language packages.</p>
<h3>Contains and ContainsAny functions</h3>
<p>In the strings package, we have the <a href="https://pkg.go.dev/strings#Contains">Contains</a> and <a href="https://pkg.go.dev/strings#ContainsAny">ContainsAny</a> method which checks for the presence of substrings within a string. This will help in looking up hidden patterns and find for substrings in a string.</p>
<h4>strings.Contains</h4>
<ul>
<li>Parameters: string, string</li>
<li>Return Type: bool (true, false)</li>
</ul>
<p>The Contains method helps in getting the exact match of the substring in the given string. Firstly, the method takes two parameters one being the actual string and the other being the substring that you want to find inside the string. Let's say we have the <code>string=&quot;bootcamp&quot;</code> and <code>substr=&quot;camp&quot;</code>, then the <code>Contains</code> method will return true because the string contains the substring <code>camp</code>.</p>
<h4>strings.ContainsAny</h4>
<ul>
<li>Parameters: string, string</li>
<li>Return Type: bool (true, false)</li>
</ul>
<p>The <code>ContainsAny</code> method just like the Contains method takes two parameters string and the other as the substring, but it would return true if it finds any character or a single byte(Unicode chars) inside the string. Let's say the <code>string=&quot;google photos&quot;</code> and <code>substring=&quot;abcde&quot;</code>, then the <code>ContainsAny</code> method will return true because the string contains at least one character in the substring which is <code>e</code>.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;strings&quot;
)

func main() {
	str := &quot;javascript&quot;
	substr := &quot;script&quot;
	s := &quot;python&quot;

	fmt.Println(strings.Contains(str, substr))
	fmt.Println(strings.Contains(str, s))

	fmt.Println(strings.ContainsAny(str, &quot;joke&quot;))
	fmt.Println(strings.ContainsAny(str, &quot;xyz&quot;))
	fmt.Println(strings.ContainsAny(str, &quot;&quot;))
}
</code></pre>
<p>Here, we have used the string methods like <code>Contains</code> and <code>ContainsAny</code> to find the substring inside a string. In the first example, the <code>str</code> variable is assigned as <code>&quot;javascript&quot;</code> and <code>substr</code> string as <code>&quot;script&quot;</code>. We use the <code>Contains</code> function with parameters <code>(str, substr)</code>. This will return <code>true</code> as we can see the <code>&quot;script&quot;</code> is a substring of <code>&quot;javascript&quot;</code>. Also, we have initialized the variable <code>s</code> to <code>&quot;python&quot;</code>. We use the <code>Contains</code> method with the parameters <code>(str, s)</code> which will return <code>false</code> as <code>&quot;python&quot;</code> is not a substring of <code>&quot;javascript&quot;</code>.</p>
<p>The next set of examples is of the <code>ContainsAny</code> method which will return true for any character present in the substring is present in the string :). Quite a Simple right, Just understand that any character in your substring present in the string will return <code>true</code>. As we see in the example, <code>ContainsAny</code> method is used with the parameters  <code>(&quot;javascript&quot;, &quot;joke&quot;)</code>, It will return <code>true</code> as <code>j</code> is present inside the string, though the entire substring is not present.</p>
<p>The next example is by passing <code>(&quot;javascript&quot;, &quot;xyz&quot;)</code> to the method <code>ContainsAny</code> will return <code>false</code> as we don't have either <code>&quot;x&quot;</code>, <code>&quot;y&quot;</code>, or <code>&quot;z&quot;</code> in the string. So this is how the <code>ContainsAny</code> method works.</p>
<p>Other similar methods you might be interested in learning are: <a href="https://pkg.go.dev/strings#Index">Index</a>, <a href="https://pkg.go.dev/strings#IndexAny">IndexAny</a>, <a href="https://pkg.go.dev/strings#LastIndex">LastIndex</a>, etc. you can find the list of methods in the <a href="https://pkg.go.dev/strings#pkg-functions">strings package</a>.</p>
<h3>Split and SplitAfter functions</h3>
<p>We also have methods to manipulate the string for splitting the characters or bytes with certain patterns. In the strings package, the <a href="https://pkg.go.dev/strings#Split">Split</a> and <a href="https://pkg.go.dev/strings#SplitAfter">SplitAfter</a> methods are quite handy to know about.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;strings&quot;
)

func main() {
	// Split method for splitting string into slice of string
	link := &quot;meetgor.com/blog/golang/strings&quot;
	fmt.Println(strings.Split(link, &quot;/&quot;))
	fmt.Println(strings.SplitAfter(link, &quot;/&quot;))

	// SplitAfter method for splitting string into slice of string with the pattern
	data := &quot;200kms50kms120kms&quot;
	fmt.Println(strings.Split(data, &quot;kms&quot;))
	fmt.Println(strings.SplitAfter(data, &quot;kms&quot;))
}
</code></pre>
<pre><code>go run utility.go

[meetgor.com blog golang strings]
[meetgor.com/ blog/ golang/ strings]
[200 50 120 ]
[200kms 50kms 120kms ]
</code></pre>
<p>So, from the above examples, we can see how <code>Split</code> and <code>SplitAfter</code> methods work. The Split method splits the text before and after the pattern string also removing the pattern string whereas the <code>SplitAfter</code> method keeps the pattern and splits it after it, hence we see the pattern string getting attached to the left string.</p>
<p>In the first example, we see the <code>link</code> variable being split into strings as <code>&quot;/&quot;</code> being the separator. The Split method returns the slice of the string elements which have been split. In the <code>data</code> variable, the <code>&quot;kms&quot;</code> is the separator, so we get the resultant slice as <code>[200, 50, 120]</code>, the <code>&quot;kms&quot;</code> string acts as a separator and it is ignored.</p>
<p>In the next example, we see the <code>link</code> variable being split into strings as <code>&quot;/&quot;</code> being the separator as previously but here, the splitting occurs after the separator has been parsed. So, we see <code>&quot;meetgor/&quot;</code> being split after the separator string and then <code>&quot;blog/&quot;</code> and so on. Also the next example, in the <code>data</code> variable, we see <code>&quot;200kms&quot;</code> being split instead of <code>200</code> using Split.</p>
<h4>Repeat and Fields functions</h4>
<p>In the strings package, we have methods like <a href="https://pkg.go.dev/strings#Repeat">Repeat</a> and <a href="https://pkg.go.dev/strings#Fields">Fields</a> for manipulating the text inside the string. These methods are used to populate or extract data from the string.</p>
<h5>strings.Repeat</h5>
<ul>
<li>Parameters: string, int</li>
<li>Return Type: string</li>
</ul>
<p>The <code>Repeat</code> method is used to create a string by repeating it n number of times and appending it to the string which is returned as the final string. So, the <code>Repeat</code> method takes in two parameters the string to repeat, an integer as a count to repeat the string, and returns a string.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;strings&quot;
)

func main() {
	// Repeat method for creating strings with given string and integer
	pattern := &quot;OK&quot;
	fmt.Println(strings.Repeat(pattern, 3))
}
</code></pre>
<pre><code>go run utility.go

OKOKOK
</code></pre>
<p>So in this example, we can see that the string <code>&quot;OK&quot;</code> is passed to the method <code>Repeat</code> with the integer <code>3</code> and thus it is appended into itself three times and the resultant string becomes <code>&quot;OKOKOK&quot;</code>.</p>
<h5>strings.Fields</h5>
<ul>
<li>Parameters: string</li>
<li>Return Type: []string</li>
</ul>
<p>The <code>Fields</code> method is used to extract the words from the string, that is the characters/bytes and group them with one or more consecutive white spaces. The function returns a slice of string.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;strings&quot;
)

func main() {
	// Fields method for extracting string from the given string with white space as delimiters
	text := &quot;Python is a prgramming language.   Go is not&quot;
	text_data := strings.Fields(text)
	fmt.Println(text_data)
	for _, d := range text_data {
		fmt.Println(&quot;data = &quot;, d)
	}
}
</code></pre>
<pre><code>go run utility.go

[Python is a prgramming language. Go is not]
data =  Python
data =  is
data =  a
data =  prgramming
data =  language.
data =  Go
data =  is
data =  not
</code></pre>
<p>The above example demonstrates the usage of <code>Fields</code> which will extract characters and split after encountering whitespace. So, we return a slice of string in which the string elements are split before white space. Thus, using the Fields method we get the words or characters as space-separated values in the slice.</p>
<p>You can even expand on this method with the <a href="https://pkg.go.dev/strings#FieldsFunc">FieldsFunc</a> method which allows you to write a custom delimiter option and extract data according to your requirement. There are tons of methods in the strings package for working with strings, you can see a detailed list of functions in the <a href="https://pkg.go.dev/strings#pkg-functions">documentation</a>.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So, from this post, we were able to understand the different methods and types to concatenate and interpolate strings in golang. We explored different types of concatenating strings, string comparison, and various methods for manipulating and interpolating strings.</p>
<p>Thank you for reading, if you have any queries, feedback, or questions, you can freely ask me on my social handles. Happy Coding :)</p>
