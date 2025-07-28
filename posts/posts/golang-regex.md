{
  "title": "Golang: Regex",
  "status": "published",
  "slug": "golang-regex",
  "date": "2025-04-05T12:33:33Z"
}

<h2>Introduction</h2>
<p>In this 26th part of the series, we will be covering the basics of using regular expressions in golang. This article will cover the basic operations like matching, finding, replacing, and sub-matches in a regular expression pattern from string source or file content. This will have examples for each of the concepts and similar variants will follow the same ideology in self-exploring the syntax.</p>
<h2>Regex in golang</h2>
<p>So, let's start with what are regular expressions.</p>
<blockquote>
<p>Regular expressions are basic building blocks for searching, pattern matching, and manipulation from a source of text using computational logic.</p>
</blockquote>
<p>This is not the formal definition, but have written it in words for my understanding of regular expressions till now. Might not be accurate, but makes sense to me after I had played and explored it (not fully).</p>
<p>So regular expressions use some pattern-matching techniques using basic logic operators like concatenation, quantifiers, etc. These relate to the study of the theory of computation quite closely but you don't need to get into too much theory in order to understand the working of regular expressions. However, it won't harm you if you are curious about it and want to explore it further.</p>
<p>Some resources to learn the fundamentals of regular expressions:</p>
<ul>
<li>
<p><a href="https://cs.lmu.edu/~ray/notes/regex/">Regular Expressions LMU notes</a></p>
</li>
<li>
<p><a href="https://regexone.com/">RegexOne</a></p>
</li>
</ul>
<h2>Regexp package</h2>
<p>We will be using the <a href="https://pkg.go.dev/regexp">regexp</a> package in the golang standard library to get some quick and important methods for quick and neat pattern matching and searching. It provides a <code>Regexp</code> type and a lot of methods on top of it to perform matching, finding, replacing, and sub-matches in the source text.</p>
<p>This package also supports two types of methods serving different purposes and use cases for string and slice of bytes, this can be useful for reading from a buffer, file, etc., and also flexible enough to search for a simple string.</p>
<h2>Matching Patterns</h2>
<p>One of the fundamental aspects of the regular expression is to check if a particular pattern is present or not in a source string. The <code>regexp</code> package provides some methods like <a href="https://pkg.go.dev/regexp#Match">Match</a>, <a href="https://pkg.go.dev/regexp#MatchString">MatchString</a> methods on a slice of bytes and string respectively from the pattern string.</p>
<h3>Matching Strings</h3>
<p>The basic operations with regex or regular expression can be performed to compare and match if the pattern matches a given string.</p>
<p>In golang, the <a href="https://pkg.go.dev/regexp">regexp</a> package provides a few functions to simply match expressions with strings or text. One of the easy-to-understand ones include the <a href="https://pkg.go.dev/regexp#MatchString">MtachString</a>, and <a href="https://pkg.go.dev/regexp#Match">Match</a> methods.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;regexp&quot;
)

func log_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	str := &quot;gophers of the goland&quot;
	is_matching, err := regexp.MatchString(&quot;go&quot;, str)
	log_error(err)
	log.Println(is_matching)
	is_matching, err := regexp.MatchString(&quot;no&quot;, str)
	log_error(err)
	log.Println(is_matching)

}
</code></pre>
<pre><code class="language-console">$ go run main.go

true
false
</code></pre>
<p>In the above code, we have used the <code>MatchString</code> method that takes in two parameters string/pattern to find, and the source string. The function returns a boolean as <code>true</code> or <code>false</code> if the pattern is present in the source string, also it might return an error if the pattern(first parameter) is parsed as an incorrect regular expression.</p>
<p>So, we can clearly see, the string <code>go</code> is present in the string <code>gophers of the goland</code> and the string <code>no</code> is not a substring.</p>
<p>We also have <code>Match</code> method which is a more general version of <code>MatchString</code> it excepts a slice of byte rather than a string as the source string. The first parameter is still a string, but the second parameter is a slice of bytes.</p>
<pre><code class="language-go">is_matching, err = regexp.Match(`.*land`, []byte(&quot;goland is a land of gophers&quot;))
log_error(err)
log.Println(is_matching)
</code></pre>
<pre><code class="language-console">$ go run main.go

true
</code></pre>
<p>We can use the <code>Match</code> method to simply parse a slice of bytes to use as the source text to check the pattern.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
    &quot;os&quot;
	&quot;regexp&quot;
)

func log_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	file_content, err := os.ReadFile(&quot;temp.txt&quot;)
	log_error(err)
	is_matching, err = regexp.Match(`memory`, file_content)
	log_error(err)
	log.Println(is_matching)
	is_matching, err = regexp.Match(`text `, file_content)
	log_error(err)
	log.Println(is_matching)
}
</code></pre>
<pre><code class="language-bash"># temp.txt

One of the gophers used a slice,
the other one used a arrays.
Some gophers were idle in the memory.
</code></pre>
<pre><code class="language-console">$ go run main.go

true
false
</code></pre>
<p>We can now even parse the contents of a file as a slice of bytes. So, it would be really nice to compare and check for a string pattern in a file quickly. Here in the above example, we have checked if <code>memory</code> is present in the text which it is, and in the second call, we check if <code>text</code> string is present anywhere in the file content which it is not.</p>
<h2>Find Patterns</h2>
<p>We can even use regular expressions for searching text or string with the struct/type <a href="https://pkg.go.dev/regexp#Regexp">Regexp</a> provided by golang's regexp. We can create a regular expression and use other functions like <code>MatchString</code>, <code>Match</code>, and others that we will explore to match or find a pattern in the text.</p>
<h3>Find String from RegEx</h3>
<p>We can get a slice of byte from the <code>FindAll</code> method which will take in a slice of byte, the second parameter as -1 for all matches. The function returns a slice of slice of byte with the byte representation of the matched string in the source text.</p>
<pre><code class="language-bash">exp, err := regexp.Compile(`\d{5}(?:[-\s]\d{4})?`)
log_error(err)
pincode_file, err := os.ReadFile(&quot;pincode.txt&quot;)
log_error(err)
match := exp.FindAll(pincode_file, -1)
log.Println(match)
</code></pre>
<pre><code class="language-bash"># pincode.txt

Pincode: 12345-1234
City, 40084
State 123
</code></pre>
<pre><code class="language-console">$ go run main.go

[[49 50 51 52 53 45 49 50 51 52] [52 48 48 56 52]]
</code></pre>
<p>In the above example, we have used the <a href="https://pkg.go.dev/regexp#Compile">Compile</a> method to create a regular expression and <a href="https://pkg.go.dev/regexp#Regexp.FindAll">FindAll</a> to get all the occurrences of the matching patterns in the text. We have again read the contents from the file. In this example, the <code>exp</code> is a regular expression for a postal code, which can either have 5 digit or 5digits-4digit combination. We read the file <code>pincode.txt</code> as a slice of bytes and use the <code>FindAll</code> method. The FindAll method takes in the parameter as a slice of byte and the integer as the number of occurrences to search. If we use a negative number it will include all the matches.</p>
<p>We search for the pin code in the file and the funciton returns a list of bytes that match the regular expression in the provided object <code>exp</code>. Finally, we get the result as <code>12345-1234</code> and <code>40084</code> which are present in the file. It doesn't match the number <code>123</code> which is not a valid match for the given regular expression.</p>
<p>There is also a version of <code>FindAll</code> as <code>FindAllString</code> which will take in a string as the text source and return a slice of strings.</p>
<pre><code class="language-go">matches := exp.FindAllString(string(pincode_file), -1)
log.Println(matches)
</code></pre>
<pre><code class="language-console">$ go run main.go

[&quot;12345-1234&quot; &quot;40084&quot;]
</code></pre>
<p>So, the <code>FindAllString</code> would return a slice of strings of the matches in the text.</p>
<h3>Find the Index of String from RegEx</h3>
<p>We can even get the start and end index of the matched string in the text using the <a href="https://pkg.go.dev/regexp#Regexp.FindIndex">FindIndex</a> and <a href="https://pkg.go.dev/regexp#Regexp.FindAllIndex">FindAllIndex</a> methods to get the indexes of the match, or all the matches of the file content.</p>
<pre><code class="language-go">exp, err := regexp.Compile(`\d{5}(?:[-\s]\d{4})?`)
log_error(err)
pincode_file, err := os.ReadFile(&quot;pincode.txt&quot;)
log_error(err)

match_index := exp.FindIndex(pincode_file)
log.Printf(&quot;%T
&quot;, match_index)
log.Println(match_index)
</code></pre>
<pre><code class="language-console">$ go run main.go

[]int
[9 19]
</code></pre>
<p>The above code is using the <code>FindIndex</code> method to get the indexes of the first match of the regular expression. The output of the code gives a single slice of an integer with length two as the start and last index of the matched string in the text file. So, here, the <code>9</code> represents the position(index) of the first character match in the string, and <code>19</code> is the last index of the matched string.</p>
<pre><code class="language-bash">Pincode: 12345-1234
01234567890123456789

Assume 0 as 10 after 9, 11, and so on
It starts from the 9th character as `1` and ends at the `4` character
at position 18 but it returns the end position + 1 for the ease of slicing
 
City, 40084
012345678901

State 123
234567890
</code></pre>
<p>The character at 9 and 18 are the first and the last character position/index of the source string, so it returns the end position + 1 as the index. This makes retrieval of slicing source string much easier, as we won't be offset by one.</p>
<p>If we want to get the text in the source string, we can use the slicing as:</p>
<pre><code class="language-go">exp, err := regexp.Compile(`\d{5}(?:[-\s]\d{4})?`)
log_error(err)
pincode_file, err := os.ReadFile(&quot;pincode.txt&quot;)
log_error(err)

match_index := exp.FindIndex(pincode_file)
if len(match_index) &gt; 0 {
    
    // Get the slice of the original string from start to end index
    sliced_string := pincode_file[match_index[0]:match_index[1]]
    log.Printf(&quot;%q
&quot;, sliced_string)
}
</code></pre>
<pre><code class="language-console">$ go run main.go

&quot;12345-1234&quot;
</code></pre>
<p>So, we can access the string from the source text without calling other functions, simply slicing the original string will retrieve the expected results. This is the convention of the golang standard library to make the end index exclusive.</p>
<p>Similarly, the <code>FindAllIndex</code> method can be used to get a list of list of such indexes of matched strings.</p>
<pre><code class="language-go">exp, err := regexp.Compile(`\d{5}(?:[-\s]\d{4})?`)
log_error(err)
pincode_file, err := os.ReadFile(&quot;pincode.txt&quot;)
log_error(err)

match_indexes := exp.FindAllIndex(pincode_file)
log.Printf(&quot;%T
&quot;, match_indexes)
log.Println(match_indexes)
</code></pre>
<pre><code class="language-console">$ go run main.go

[][]int
[[9 19] [26 31]]
</code></pre>
<p>The above example gets the list of slices as for all the search pattern indexes in the source string/text. We can iterate over the list and get each matched string index.</p>
<h2>Find Submatch</h2>
<p>the <code>regexp</code> package also has a utility function for finding the sub-match of a given regular expression. The method returns a list of strings (or slices of bytes) containing the leftmost match and the sub-matches in that match. This also has the <code>All</code> version which instead of returning a single match i.e. the leftmost match it returns all the matches and the corresponding sub-matches.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;regexp&quot;
)

func main() {

	str := &quot;abc@def.com is the mail address of 8th user with id 12&quot;
	exp := regexp.MustCompile(
        `([a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,})` + 
            `|(email|mail)|` + 
            `(\d{1,3})`)`,
    )
	match := exp.FindStringSubmatch(str)
	log.Println(match)
	matches := exp.FindAllStringSubmatch(str, -1)
	log.Println(matches)
}
</code></pre>
<pre><code class="language-console">$ go run main.go

[abc@def.com abc@def.com  ]
[[abc@def.com abc@def.com  ] [mail  mail ] [8   8] [12   12]]
</code></pre>
<p>The above example uses a regex to compute a few things like email address, words like <code>mail</code> or <code>email</code>, also a number of up to 3 digits. The <code>|</code> between these expressions indicates any combination of these three things would be matched for the regular expression. The <a href="https://pkg.go.dev/regexp#Regexp.FindStringSubmatch">FindStringSubmatch</a> method, takes in a string as the source and returns a slice of the matching pattern. The first element is the leftmost match in the string source for the given regular expression, and the subsequent elements are the sub-matches in the matched string.</p>
<p>We can now move a little step ahead for actually understanding the sub-match in a regular expression.</p>
<pre><code class="language-go">str := &quot;abe21@def.com is the mail address of 8th user with id 124&quot;
exp := regexp.MustCompile(
    `([a-zA-Z]+(\d*)[a-zA-Z]*@[a-zA-Z]*(\d*)[a-zA-Z]+\.[a-z|A-Z]{2,})` +
        `|(mail|email)` +
        `|(\d{1,3})`,
)

match := exp.FindStringSubmatch(str)
log.Println(match)
matches := exp.FindAllStringSubmatch(str, -1)
log.Println(matches)
</code></pre>
<pre><code class="language-console">$ go run main.go

[abe21@def.com abe21@def.com 21   ]
[[abe21@def.com abe21@def.com 21   ] [mail    mail ] [8     8] [124     124]]
</code></pre>
<p>In the above example, there are a few things to take away, let us break it down into small pieces. We have a regex for matching either a mail address, the word <code>mail</code> or <code>email</code>, or a number up to 3 digits. The regex is a bit different from the previous example for understanding the sub-matches within an expression. We find the sub-matches in the string which will be handled by the <code>FindStringSubmatch</code>. This function takes in a string and returns a list of strings that are the leftmost matches in the source string.</p>
<p>First, we need to understand the regex to get a clear idea of the code snippet. The first sub-match is for the email address. However, we use <code>[a-zA-Z]</code> in the picking of the username and the domain name, we don't want to directly match numbers yet. The goal of this regex is to pick up numbers inside an email address. So, we can have 1 or more characters <code>[a-zA-z]+</code>, followed by 0 or more digits <code>(\d*)</code>, and again we can have 0 or more characters <code>[a-zA-Z]*</code>. The <code>+</code> is for 1 or more, <code>*</code> is for 0 or more, <code>\d</code> is for digits. After this, we have the <code>@</code> as a compulsory character in the mail, followed by the same sequence in the username i.e. 1 or more characters, 0 or more digits, 0 or more characters. Finally, we have the <code>.com</code> and the domain name extensions, as a group of 2 or more characters <code>[a-z|A-Z]{2,}</code>.</p>
<p>So, the regex accepts an email, with a sub-match of the number anywhere in the username or the domain name.</p>
<p>The <code>FindStringSubmatch</code> function lists out the sub-matches for the leftmost(first) match of the regex in the source string. So it finds the string <code>abc21@def.com</code> which is the email id. This regex <code>([a-zA-Z]+(\d*)[a-zA-Z]*@[a-zA-Z]+(\d*)[a-zA-Z]*\.[a-z|A-Z]{2,})</code> has two sub-matches <code>\d*</code> in the username part and also in the domain part. So, the list returns the email address as the match and the match as itself, as well as the number in the sub-match inside the mail address. So the result, <code>[abc21@def.com abc21@def.com 21 ]</code>, there are a few empty string sub-matches because the second sub-match for the domain name number returns an empty string.</p>
<p>Similarly, the <code>FindAllStringSubmatch</code> will return the list of all the matches in the source string. The other matches don't have any sub-matches in the regular expression so it just gets the match and the sub-match as itself in the case of string <code>mail</code>, digit <code>8</code>, and digit <code>124</code>.</p>
<p>We can also use this example from a file as a slice of bytes. This will return a list of slice of slice of bytes instead of slice of slice of strings.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;regexp&quot;
)

func main() {

    exp := regexp.MustCompile(
        `([a-zA-Z]+(\d*)[a-zA-Z]*@[a-zA-Z]*(\d*)[a-zA-Z]+\.[a-z|A-Z]{2,})` +
            `|(mail|email)` +
            `|(\d{1,3})`,
    )
	email_file, err := os.ReadFile(&quot;subtext.txt&quot;)
	log_error(err)
	mail_match := exp.FindSubmatch(email_file)
	log.Printf(&quot;%s
&quot;, mail_match)
	mail_matches := exp.FindAllSubmatch(email_file, -1)
	//log.Println(mail_matches)
	log.Printf(&quot;%s
&quot;, mail_matches)

}
</code></pre>
<pre><code class="language-txt"># submatch.txt

abc21@def.com is the mail address of user id 1234
The email address abe2def.com is of user name abc
a2be.@def.com
Email address: abe@de2f.com, User id: 45
johndoe@example.com
jane.doe123@example.com
janedoe@example.co.uk
john123@example.org
janedoe456@example.net
</code></pre>
<pre><code class="language-console">$ go run main.go

[][]unit8
[abc21@def.com abc21@def.com 21   ]

[][][]unit8
[
    [abc21@def.com abc21@def.com 21   ] [mail    mail ] [123 123] [4     4] 
    [email    email ] [2     2] [2     2] [mail    mail ] 
    [abe@de2f.com abe@de2f.com  2  ] [45     45] 
    [johndoe@example.com johndoe@example.com    ] 
    [doe123@example.com doe123@example.com 123   ] 
    [janedoe@example.co janedoe@example.co    ] 
    [john123@example.org john123@example.org 123   ] 
    [janedoe456@example.net janedoe456@example.net 456   ]
]
</code></pre>
<p>As we can see from the dummy email ids and some random text, we are able to match the email ids and the numbers in them as the sub-match of the string. This is a <code>[][]unit8</code> in case of the <code>FindSubmatch</code> and <code>[][][]unit8</code> in case of <code>FindAllSubmatch</code>. The working remains the same for the bytes as it was for the strings.</p>
<h3>Find Submatch Index</h3>
<p>We also have the <a href="https://pkg.go.dev/regexp#Regexp.FindSubmatchIndex">FindSubmatchIndex</a> and <a href="https://pkg.go.dev/regexp#Regexp.FindAllSubmatchIndex">FindAllSubmatchIndex</a>, and the string variants to get the index(es) of the sub-matches in the pattern picked from the regular expression.</p>
<pre><code class="language-go">str := &quot;abe21@def.com is the mail address of 8th user with id 124&quot;
exp := regexp.MustCompile(
    `([a-zA-Z]+(\d*)[a-zA-Z]*@[a-zA-Z]*(\d*)[a-zA-Z]+\.[a-z|A-Z]{2,})` +
        `|(mail|email)` +
        `|(\d{1,3})`,
)

match := exp.FindStringSubmatch(str)
match_index := exp.FindStringSubmatchIndex(str)
log.Println(match)
log.Println(match_index)
</code></pre>
<pre><code class="language-console">$ go run main.go

[abe21@def.com abe21@def.com 21   ]
[0 13 0 13 3 5 9 9 -1 -1 -1 -1]
</code></pre>
<p>So this returns a list of pairs, here consider the list as <code>[(0, 13) (0, 13), (3, 5), (9, 9), (-1, -1), (-1, -1)]</code>. Because the list represents the start and end indexes of the sub-matches in the source string. The match is the first pair, i.e. at the first character <code>0</code> -&gt; <code>a</code>, and ends at space i.e. <code>13</code> character off by one. Then we have the sub-match itself with the same indexes. Then the <code>3</code> and <code>5</code> indicating the number sub-match <code>21</code> in the <code>abc21@def.com</code> it starts at 3 and ends at 5, so it occupies 3 and 4 characters in the source string. Similarly, the domain level number doesn't have any number in the source string so it has returned the domain name end index as an empty string.</p>
<p>We had used <code>(\d*)</code> which can return 0 or more occurrences of the digit, so it returned no match in the case of the domain level name, hence we get the <code>9</code> as an empty string at the end of the sub-match for it. The rest are for the <code>email</code> or <code>mail</code>, and the lone digit sub-matches in the regular expression which we don't have in the first match in the source string.</p>
<pre><code class="language-go">str := &quot;abe21@def.com is the mail address of 8th user with id 124&quot;
exp := regexp.MustCompile(
    `([a-zA-Z]+(\d*)[a-zA-Z]*@[a-zA-Z]*(\d*)[a-zA-Z]+\.[a-z|A-Z]{2,})` +
        `|(mail|email)` +
        `|(\d{1,3})`,
)

match := exp.FindAllStringSubmatch(str, -1)
match_indexes := exp.FindAllStringSubmatchIndex(str, -1)
log.Println(match_indexes)
</code></pre>
<pre><code class="language-console">$ go run main.go

[[abe21@def.com abe21@def.com 21   ] [mail    mail ] [8     8] [124     124]]
[
    [0 13 0 13 3 5 9 9 -1 -1 -1 -1]
    [21 25 -1 -1 -1 -1 -1 -1 21 25 -1 -1]
    [37 38 -1 -1 -1 -1 -1 -1 -1 -1 37 38]
    [54 56 -1 -1 -1 -1 -1 -1 -1 -1 54 57]
]
</code></pre>
<p>Similarly, we have used the <a href="https://pkg.go.dev/regexp#Regexp.FindAllStringSubmatchIndex">FindAllStringSubmatchIndex</a> for getting the list of a slice of indexes(int) for the sub-match of the expression in the source string.</p>
<p>The first element is the same as the previous example, the next match in the source string is <code>mail</code> which comes at index 21 and ends at the 24th index which golang does 24+1 as a convention. Similarly, the number <code>8</code> is matched at index <code>37</code> and the number <code>124</code> at index <code>54</code>, rest of the sub-matches for these matches are not present so it turns up to be -1 for the rest of the sub-matches.</p>
<p>So this can also be used for the byte/uint8 type of variants with <a href="https://pkg.go.dev/regexp#Regexp.FindSubmatchIndex">FindSubmatchIndex</a> and <a href="https://pkg.go.dev/regexp#Regexp.FindAllSubmatchIndex">FindAllSubmatchIndex</a>.</p>
<h2>Replace Patterns</h2>
<p>The replace method is used for replacing the matched patterns.</p>
<p>The <a href="https://pkg.go.dev/regexp#Regexp.ReplaceAll">ReplaceAll</a> and <a href="https://pkg.go.dev/regexp#Regexp.ReplaceAllLiteral">ReplaceAllLiteral</a> with some string and byte slice variations can help us in replacing the source text with a replacement string against a regular expression.</p>
<p>Let's start with a simple example of strings.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;regexp&quot;
)

func main() {

	str := &quot;Gophy gophers go in the golang grounds&quot;
	exp := regexp.MustCompile(`(Go|go)`)
	replaced_str := exp.ReplaceAllString(str, &quot;to&quot;)
	log.Println(replaced_str)

}
</code></pre>
<p>The above code can replace the string with regex with the replacement string. The regex provided in the <code>exp</code> variable has a pattern matching for <code>Go</code> or <code>go</code>. So, the <code>ReplaceAllString</code> method takes in two arguments as strings and returns the string after replacing the character.</p>
<pre><code class="language-console">$ go run replace.go

Original:  Gophy gophers go in the golang grounds
Replaced:  tophy tophers to in the tolang grounds
</code></pre>
<p>So, this replaces all the characters with <code>Go</code> or <code>go</code> in the source string with <code>to</code>.</p>
<p>There is a special value available in the replacement string which can expand the regular expression literals. Since the regular expression consists of multiple literals/quantifiers, we can use those to replace or keep the string in a particular expression's evaluation.</p>
<pre><code class="language-go">str = &quot;Gophy gophers go in the golang grounds&quot;
exp2 := regexp.MustCompile(`(Go|go)|(phers)|(rounds)`)
log.Println(exp2.ReplaceAllString(str, &quot;hop&quot;))
log.Println(exp2.ReplaceAllString(str, &quot;$1&quot;))
log.Println(exp2.ReplaceAllString(str, &quot;$2&quot;))
log.Println(exp2.ReplaceAllString(str, &quot;$3&quot;))
</code></pre>
<pre><code class="language-console">$ go run replace.go

hopphy hophop hop in the hoplang ghop
Gophy go go in the golang g
phy phers  in the lang g
phy   in the lang grounds
</code></pre>
<p>The above code uses the regular expression match string have been replaced with the literals. So, the regex <code>(Go|go)|(phers)|(rounds)</code> has three parts <code>Go|go</code>, <code>phers</code>, and <code>rounds</code>. Each of the literals is separated by an operator, indicating either the match or all matches should be considered.</p>
<p>In the first statement, we replace the regex with <code>hop</code> as you can see all the matches are replaced with the replacement string. For instance, the word <code>gophers</code> is replaced by <code>hophop</code>, because <code>go</code> and <code>phers</code> is matched separately and replaced.</p>
<p>In the second statement, we replace the source string with <code>$1</code> indicating the first literal in the regex i.e. <code>Go|go</code>. These statements will expand the <code>$1</code> and keep the match only where the literal matches and rest are removed. So, <code>Gophy</code> is matched with <code>Go|go</code> so it is kept as is in the replacement. However, for <code>grounds</code>, the literal match for <code>$1</code> i.e. <code>Go|go</code> does not match and hence is not kept and removed, so the resulting string becomes <code>g</code> rest is substituted with an empty string.</p>
<p>In the third print statement, the source string is replaced with the second literal <code>$2</code> i.e. <code>phers</code>. So if any string matches the literal, only those are kept and the rest are substituted by an empty string. So, <code>Gophy</code> or <code>Go</code> doesn't match <code>phers</code> and hence is replaced, however, <code>gophers</code> match the <code>phers</code> part and is kept as it is, but the <code>go</code> part is substituted.</p>
<p>Similarly, for the fourth print statement, the source string is replaced with the third literal i.e. <code>rounds</code>. So if only the third literal is kept as is, the rest matching strings from the regex are substituted with an empty string. So, <code>grounds</code> remain as it is because it matches the <code>rounds</code> in the replacement literal.</p>
<p>In short, we replace the literal back after replacing the regex patterns in the source string. This can be used to fine-tune or access specific literals or expressions in the regular expressions.</p>
<pre><code class="language-go">str = &quot;Gophy gophers go in the golang grounds&quot;
exp2 := regexp.MustCompile(`(Go|go)|(phers)|(rounds)`)
log.Println(exp2.ReplaceAllString(str, &quot;$1$2&quot;))

str = &quot;Gophy gophers go in the golang cophers grounds&quot;
log.Println(exp2.ReplaceAllString(str, &quot;$1$3&quot;))
</code></pre>
<pre><code class="language-console">$ go run replace.go

Gophy gophers go in the golang g
Gophy go go in the golang co grounds
</code></pre>
<p>We can even concatenate and make some minor string literal adjustments to the replacement string. As we have done in the example, where both the <code>$1$2</code> are used as the replacement string. The two literals combine to make a string with <code>Go|go</code> and <code>phers</code>. So, we can see the result, in the first statement, <code>Gophy gophers go in the golang g</code>, all the character that have <code>Go|go</code> or <code>phers</code> are kept as it is(substituted the same string), the <code>grounds</code> however does not match and hence are replaced with the empty string(because the capture group <code>$1$2</code> does not match <code>rounds</code>).</p>
<p>Similarly, for the third example, the <code>$1$3</code> i.e. <code>Go|go</code> or <code>rounds</code> are matched with the source string. So, we the <code>phers</code> in <code>gophers</code> and <code>cophers</code> does not match the capture group <code>$1$3</code> and hence is replaced by an empty group(string). However, the <code>Gophy</code>, <code>golang</code>, and <code>grounds</code> match the capture group and are replaced by that match string (which is the same string).</p>
<p>If we want to avoid expansion of the strings as we did in the previous example with <code>$1</code> and other parameters, we can use the <a href="https://pkg.go.dev/regexp#Regexp.ReplaceAllLiteral">ReplaceAllLiteral</a> or <a href="https://pkg.go.dev/regexp#Regexp.ReplaceAllLiteralString">ReplaceAllLiteralString</a> to parse the string as it is.</p>
<pre><code class="language-go">str := &quot;Gophy gophers go in the golang cophers grounds&quot;
exp2 := regexp.MustCompile(`(Go|go)|(phers)|(rounds)`)
log.Println(exp2.ReplaceAllLiteralString(str, &quot;$1&quot;))
</code></pre>
<pre><code>$ go run replace.go

$1phy $1$1 $1 in the $1lang co$1 g$1
</code></pre>
<p>As we can see the <code>$1</code> is not expanded and parsed as it is for replacing the pattern in the regular expression. The <code>Go</code> is replaced with <code>$1</code> to look like <code>$1phy</code>, and similarly for the rest of the patterns.</p>
<p>That's it from the 26th part of the series, all the source code for the examples are linked in the GitHub on the <a href="https://github.com/Mr-Destructive/100-days-of-golang/tree/main/scripts/regex">100 days of Golang</a> repository.</p>
<h2>Conclusion</h2>
<p>This article covered the fundamentals of using the <code>regexp</code> package for working with regular expressions in golang. We explored the methods and <code>Regexp</code> type in the package with various methods available through the type interface. By exploring the examples and simple snippets, various ways for pattern matching, finding, and replacing were walked over and found.</p>
<p>So, hopefully, the article might have found useful to you, if you have any queries, questions, feedback, or mistakes in the article, you can let me know in the discussion or on my social handles. Happy Coding :)</p>
