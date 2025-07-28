{
  "title": "BASH Scripting Guide - PART - 2",
  "status": "published",
  "slug": "bash-guide-p2",
  "date": "2025-04-05T12:33:58Z"
}

<h1>Introduction</h1>
<p>In this part, topics such as switch cases, mathematical expression, arrays, and strings. This won't be an in-depth guide to understand each and every concept, but to make users aware of the things and features in Bash. This also would not be an absolute basic guide, I expect to have some basic programming knowledge such as binary systems, logical and mathematical concepts. Don't worry, you won't be bombarded with commands, I'll just explain with easy examples to get started.</p>
<p>Topics to be covered in this part are as follows:</p>
<ul>
<li>
<p>User Input</p>
<ul>
<li>User Prompt</li>
<li>Changing the Delimiter</li>
<li>Password as Input</li>
<li>Limiting the length of Input</li>
</ul>
</li>
<li>
<p>Cases</p>
</li>
<li>
<p>Arrays</p>
<ul>
<li>Declaring and Printing Arrays</li>
<li>Number of elements in an array</li>
<li>Splicing the array</li>
<li>Inserting and Deleting elements</li>
</ul>
</li>
<li>
<p>Strings</p>
<ul>
<li>Making Substrings</li>
<li>String Concatenation</li>
<li>String Comparison</li>
</ul>
</li>
<li>
<p>Arithmetic in Bash</p>
<ul>
<li>Integer Arithmetic</li>
<li>Floating-Point Arithmetic</li>
</ul>
</li>
</ul>
<h1>User Input</h1>
<p>Taking user input in Bash is quite straightforward and quite readable as well. We can make use of <code>read</code> command to take in input from the user. We just specify the variable in which we want to store the input.<code> read x</code> Here, the input will be stored in x. We can also pass in certain arguments to the read command such as -p (prompt with string), -r ( delimiter variation), -a(pass to the array), and others as well. Each of them will make the foundation of various complicated tasks to perform in logical operations.</p>
<h3>User prompt argument</h3>
<p>The -p argument will prompt the user with a string before they input anything. It makes quite informative and useful user input. This becomes quite a useful argument/parameter to make it quite readable and understand what to do directly with much hassle. The below is the general syntax of passing the argument to the read function.</p>
<pre><code class="language-bash">read -p &quot;Enter the number &quot; n
</code></pre>
<pre><code class="language-bash">#!/bin/bash

read -p &quot;Enter the number &quot; n
echo &quot;The inputted number was $n&quot;
</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625118915300/NRF7Ci2rK.png" alt="bashs2.png"></p>
<p>In this example, we can prompt the user with the string <strong>Enter the number</strong>, and it gives certain information to the user about what to input.</p>
<h3>Changing the delimiter</h3>
<p>Next, we can make use of -r which depending on the use case, we can change the delimiter while taking the input.</p>
<pre><code class="language-bash">#!/bin/bash

IFS='/' read -p &quot;Enter the file path : &quot; user project app 
echo $user $project $app

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625202319183/R9Eo3lg0oU.png" alt="bashs2.png"></p>
<p>In the above script, we separated the path of a directory(user-entered) into separate components such as the username, project name, and the app name, this can get pretty important and a great tool for automation of making project and application structures. At the beginning of the command, we use IFS which stands for Internal Field Separator, which does the separation of variables based on the field provided, in this case it was <code>//</code>, you can use any other field characters appropriate to your needs.</p>
<p>This command will change the delimiter, by default it uses spaces or tabs etc to identify distinct input variables but we change it to other internal field separators such as / , \ ,- , |, etc. This can make the user input more customizable and dynamic.</p>
<h3>Password Typing</h3>
<p>We can hide the input from the screen so as to keep it confidential and keep sensitive information such as passwords and keys private and protected.</p>
<pre><code class="language-bash">#!/bin/bash

read -sp &quot;Password: &quot; pswd
echo &quot;the password was $pswd&quot;

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625161571993/MkCadbyKW.png" alt="bashs2.png"></p>
<p>We used the -s to keep the input hidden, the screen doesn't reflect what the user is typing, and -p for the prompt to offer the user some information on the text.</p>
<h3>Limiting Length of Input</h3>
<p>We can limit the user to only a certain number of characters as input. It becomes quite useful in constrained environments such as usernames and passwords to be restricted with a certain limit.</p>
<pre><code class="language-bash">#!/bin/bash

read -n 6 -p &quot;Enter the name: &quot; n
echo $n
</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625161752440/8xI5Lovxo.png" alt="bashs2.png"></p>
<p>In the above script, we demonstrate the limit of characters of 6 in the variable n. This restricts the user with only the first 6 characters, it just doesn't exceed ahead, directly to the next command.</p>
<h3>Passing to the array</h3>
<p>Another important argument to be passed after read command is -a which inserts the value to the array elements.</p>
<pre><code class="language-bash">#!/bin/bash

read -a nums -p &quot;Enter the elements : &quot; 
for i in ${nums[*]};
do 
	echo -e &quot;$i
&quot;
done

</code></pre>
<p>In the above script, we have used array, don't worry, I'll explain it in the coming sections of this part. We have assigned an empty array and made the user enter those arrays, they are space-separated values. We have used the -a operator to insert the input to the elements of the array. The for loop is range-based which means it will do certain commands until there are no elements left in nums. The command <code>${nums[@]}</code> indicates every element in the array nums.</p>
<h1>Cases</h1>
<p>Cases are quite a good way of replacing nested if-else statements to make them nice and readable in the script.  Cases in Bash are quite powerful and easy to use compared with C/ C++ styled switch cases.</p>
<p>The general structure of using a case in Bash is as follows:</p>
<pre><code class="language-bash">case variable in
    pattern 1)
        statements
        ;;
    pattern 2)
        statements
        ;;
    pattern 3)
        statements
        ;;
    pattern 4)
        statements
        ;; 
    *)
        statements
        ;;
esac
</code></pre>
<p>It follows a particular pattern if it matches it stops the search and executes the statements, finally comes out of the block. If it doesn't find any match it redirects to a default condition if any.</p>
<pre><code class="language-bash">#!/bin/bash 

read -p &quot;Enter a name : &quot; n
case $n in 
	admin)
		echo &quot;You are logged as root&quot;
		;;
	unknown)
		echo &quot;A hacker probably&quot;
		;;
	manager)
		echo &quot;Weolcome Manager!&quot;
		;;
	*)
		echo &quot;A normal person&quot;
		;;
esac

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625160454707/YDYGkU34d.png" alt="bashs2.png"></p>
<p>Seeing the above example, it is quite clear, that it looks quite structured and readable than the nested ladder of If-else statements. Cases are rendered based on the variable after the <code>case</code> keyword. We use the patterns before <code>)</code> as making the match in the variable provided before. Once the interpreter finds a match it returns to <code>esac</code> command which is <code>case</code> spelled in reverse just like <code>fi</code> for <code>if</code> and <code>done</code> for <code>do</code> in loops :) If it doesn't match any pattern, we have a default case represented by <code>*)</code> and it executes for any non-matching expression.</p>
<h2>Arrays</h2>
<p>Arrays or a way to store a list of numbers is implemented in Bash identical to most of the general programming languages.</p>
<h3>Declaring and Printing Arrays</h3>
<p>We declare an array similar to a variable but we mention the index of the element in the array(0 based index).  We can also simply declare an empty array using the declare command <code>declare -A nums</code></p>
<pre><code class="language-bash">#!/bin/bash 

nums[0]=7
nums[1]=5
nums[2]=8 
for i in ${nums[@]}
do
echo -e &quot;$i 
&quot;
done
</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625124595563/TzBEaH1E4.png" alt="bashs2.png"></p>
<p>The above script initializes an array with some hard-coded elements, surely you can input those from the user. For printing and accessing those elements in the array, We can use a loop, here we have used a range-based for loop. You can use any other loop you prefer. The iterator is &quot; i &quot; and we use $ to access the values from the array, we use <code>{}</code> as we have nested expression for indexing the element and <code>*</code> for every element in the array ( <code>@</code> will also work fine ), that's why range-based for loops make it quite simple to use. And we have simply printed &quot; i &quot; as it holds a particular element based on the iteration.</p>
<p>OR</p>
<pre><code class="language-bash">#!/bin/bash

declare -A nums=(
[0]=44
[1]=45
[2]=46
)
echo &quot;${nums[@]}&quot;

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625128343654/NCKUzurXe.png" alt="bashs2.png"></p>
<p>The above script uses declare an array, it can be empty as well after the name declaration. We used the <code>()</code> to include the values in the array, using indices in the array we assigned the values to the particular index.</p>
<p>If you just want to print the elements, we can use <code>${nums[@]}</code> or <code>${nums[*]}</code>, this will print every element without using any iteration loops.</p>
<pre><code class="language-bash">#!/bin/bash

nums[0]=7
nums[1]=5
nums[2]=8 
echo &quot;${nums[@]}&quot;

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625125456166/85zXjapQ_.png" alt="bashs2.png"></p>
<h3>Number of Elements in the array</h3>
<p>To get the length of the array, we can use # in the expression <code>${nums[@]}</code>, like <code>${#nums[@]}</code> to get the number of elements from the array.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625125770005/VzIr8CD7-.png" alt="bashs2.png"></p>
<p>Since we had 4 elements in the array, it accurately printed 4.</p>
<h3>Inserting and Deleting elements from Array</h3>
<p>We can push elements to the array using the assignment operator.</p>
<p><code>nums=(${nums[@]} 76) </code></p>
<p>This will push 76 into the array, i.e. in the last index( length -1 index).</p>
<pre><code class="language-bash">#!/bin/bash

nums[0]=7
nums[1]=5
nums[2]=8
nums[3]=19
nums=(${nums[@]} 76)
echo &quot;${nums[@]}&quot;
echo &quot;Length of nums = ${#nums[@]}&quot;

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625126198675/A8IAE-2FR.png" alt="bashs2.png"></p>
<p>As you can see the element was added at the end of the array, you can also specify the index you want to insert. We can use <code>unset nums[3] </code> to delete the element at the particular location or we can pop back (delete from the end) an element from the array using the index <code>-1</code> from the array using the following command.</p>
<pre><code class="language-bash">unset nums[-1]
</code></pre>
<p>Any index provided will delete the element at that location by using unset. By using -1, we intend to refer to the last element. This can be quite handy and important as well in certain cases.</p>
<pre><code class="language-bash">#!/bin/bash

nums[0]=7
nums[1]=5
nums[2]=8
nums[3]=19
unset nums[-1]
echo &quot;${nums[@]}&quot;
echo &quot;Length of nums = ${#nums[@]}&quot;

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625126770211/LYK2Q0Rp0.png" alt="bashs2.png"></p>
<p>There you can see we removed the element using unset.</p>
<h3>Splicing an Array</h3>
<p>We can splice the array to print/ copy a portion of the array to another one.</p>
<pre><code class="language-bash">echo &quot;${nums[@]:1:3}&quot;
</code></pre>
<p>Using two colons and numbers in between them, we can print in this case certain elements in the array from a particular range. Here the first number after the colon is the starting index to print from(inclusive) and the next number after the second colon is the length to which we would like to print the element, it is not the index but the number of elements after the start index to be spliced</p>
<pre><code class="language-bash">#!/bin/bash

nums[0]=7
nums[1]=5
nums[2]=8
nums[3]=19
nums[4]=76
newarr=${nums[@]:1:3}
echo &quot;${newarr[@]}&quot;
echo &quot;${nums[@]}&quot;

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625127387739/IH2Fc6ghk.png" alt="bashs2.png"></p>
<p>In this case, we have copied the slice of an array to another new array using the double colon operator. We added the elements from the 1st index till <code>1+3</code> indices i.e till 4. 3 is not the index but the length till you would like to copy or print.</p>
<p>This was a basic introduction to arrays, definitely, there will be much more stuff I didn't cover. Just to give an overview of how an array looks like in BASH scripting. Next, we move on to strings.</p>
<h1>Strings</h1>
<p>Strings are quite important as it forms the core of any script to deal with filenames, user information, etc all contain strings or array of characters. Let's take a closer look at how strings are declared, handled, and manipulated in Bash scripting.</p>
<pre><code class="language-bash">s=&quot;World&quot;
echo &quot;$s&quot;
</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625129318948/02V3bpP1I.png" alt="bashs2.png"></p>
<p>Strings are again declared as normal variables but are enclosed in double quotation marks.  And we access them in the exact same way as we do with variables. If you were to use single quotes instead of double quotes Bash would not interpret the variable name as a variable, it would print the name literally and not the value of the variable, So prefer using double quotes in echo and other commands to make variables accessible.</p>
<h3>Making Substrings</h3>
<p>We can even splice the string as we did with the arrays, in strings we can call it substrings. The syntax is almost identical as we just have to get the variable name.</p>
<pre><code class="language-bash">#!/bin/bash

s=&quot;Hello World&quot;
p=${s:6}
echo $p
q=${s::5}
echo $q
t=${s:4:3}
echo $t

</code></pre>
<p>In the above script, we had a base string 's' which was then sliced from the 6th index to the end, If we do not pass the second number and colon, it interprets as the end of the string and if we do not pass the first number, it will interpret as the first character in the string. We sliced s from the 6th index till the end of the string and copied it in the string 'p''. In the 'q' string, we sliced the first 5 characters from the string 's'. In the 't' string we sliced starting from the 4th index and 3 characters in length i.e till  7th index, not the 7th index.</p>
<p>We can use the # before the variable name to get the length of the string as we saw in the array section. So we can use the <code>echo ${#s}</code> command to print the length of the string where s is the string variable name.</p>
<h3>String Concatenation</h3>
<p>String concatenation on Bash is quite straightforward as it is just the matter of adding strings in a very simple way.</p>
<pre><code class="language-bash">#!/bin/bash

s=&quot;Hello&quot;
p=&quot;World&quot;
q=&quot;$s $p&quot;
echo $q
</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625135997314/_n05RIoTTM.png" alt="bashs2.png"></p>
<p>The space between the two variables is quite literal, anything you place between this space or the double quotes will get stored in the variable or get printed.</p>
<h3>String Comparison</h3>
<p>Moving on to the string comparison in Bash. String comparison is quite complex in certain programming languages but it's quite straightforward in some languages such as Bash. We can compare strings quite easily in Bash, either they are equal or they are not, it's just comparison operators to perform the heavy-lifting for us.</p>
<pre><code class="language-bash">#!/bin/bash

s=&quot;hello&quot;
p=&quot;Hello&quot;
if [ $s = $p ];
then
    echo &quot;Equal&quot;
else 
    echo &quot;Not equal&quot;
fi

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625138020711/UWlRN8aPq.png" alt="bashs2.png"></p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625137884154/7bNPkpGd3.png" alt="bashs2.png"></p>
<p>From the above code, it is quite clear that the strings as not equal and we compared them with the &quot;equality&quot; operator (=) and checked if that condition was true, and perform commands accordingly. We can also check if the strings are not equal using <code>!=</code> operator and we can perform commands based on the desired logic. We also have operators to compare the length of the strings. We can use <code>\&lt;</code> operator to check if the first string is less than the second string(compared characters in ASCII).  And check if the first string is greater than the second string using <code>\&gt;</code> operator.</p>
<pre><code class="language-bash">#!/bin/bash

s=&quot;hello&quot;
p=&quot;Hello&quot;
if [ $s \&lt; $p ];
then

	echo &quot;$s is Less than $p&quot;
elif [ $s \&gt; $p ];
then
	echo &quot;$s is greater than $p&quot;
else
echo &quot;Equal&quot;
fi

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625137393683/u3WbgDIrN.png" alt="bashs2.png"></p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625137467448/NP1UXZAbv.png" alt="bashs2.png"></p>
<p>Here, we are using the ASCII equivalent of strings to compare them as it gives an idea in terms of the value of the strings. We see that 'h'( 104)has a greater ASCII value than 'H&quot; (72) which is why we see the shown outcome.</p>
<p>We also have operators to check for the string being empty or not using the -z operator. Also, we have arguments to pass to the string comparison to check for non-empty strings as well, specifically for input validation and some error handling.</p>
<p>We can quite easily use -n to check for non-empty string and -z for the length of the string being zero.</p>
<pre><code class="language-bash">#!/bin/bash

read -p &quot;Enter a string : &quot; s
if [ -s $s ];
then 
    echo &quot;Empty Input&quot;
else
   echo &quot;Valid input&quot;
fi

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625138907051/tbjRDda0U1.png" alt="bashs2.png"></p>
<p>So the string topic is quite straightforward and self-explanatory as it doesn't how that much complexity but is still powerful and convenient to use.</p>
<h1>Arithmetic in Bash</h1>
<p>Performing any Arithmetic operations is the core for scripting. Without arithmetic, it feels incomplete to programmatically create something, it would be quite menial to write commands by hand without having the ability to perform arithmetic operations.</p>
<h3>Integer Arithmetic</h3>
<p>Firstly we quite commonly use operations on variables, so let us see how to perform an arithmetic operation on variables in Bash. We use double curly braces to evaluate certain results of the operations performed on variables.</p>
<pre><code class="language-bash">#!/bin/bash

x=4
y=9
z=$(($x * $y))
echo $z

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625139582135/Sb4hdE990.png" alt="bashs2.png"></p>
<p>We use double curly braces in order to evaluate the operations performed on the variables inside them. We definitely have to use the $ symbol to extract the value of the variable.</p>
<p>We can surely use operators such as addition(<code>+</code>), subtraction(<code>-</code>), multiplication(<code>*</code>), division(<code>/</code>), and modulus(<code>%</code>, it stores the remainder of the division,17%3 gets you 2) in statements. We can also perform operations such as <code>&lt;&lt;</code> to do left bitwise shift and <code>&gt;&gt;</code> right bitwise shift to shift the binary digits in left tor right respectively in a variable. There are also logical operations such as Bitwise and logical AND(<code>&amp;</code>), OR(<code>|</code>), EX-OR(<code>^</code>), and ternary expressions.</p>
<p>Alternative to double curly braces is <code>expr</code>, expr allows you to freely wherever you need to evaluate an arithmetic operation. Though this is not native in-built in shells, it uses a binary process to evaluate the arithmetic operations. It can also defer depending on the implementation of such commands in various environments.</p>
<p>We can also use the <code>let</code> command to initialize a variable and perform expressions in the initialization itself.</p>
<pre><code class="language-bash">#!/bin/bash

let a=4
let b=a*a
let c=&quot;b/(a*2)&quot;
echo $b
</code></pre>
<p>We can perform quite complex operations using simple implementation using <code>let</code>, this allows much readable and bug-free scripts.  If you would like to perform operations with brackets and other operations you can enclose the expression in double quotation marks.</p>
<h3>Floating-Point Arithmetic</h3>
<p>Performing floating-point arithmetic in Bash is not quite well though. We won't get 100% accurate answers in the expressions this is because it is <strong>not designed</strong> for such things. Doing <strong>things related to floating-point is a bad idea</strong>, Still, you can improve the precision to a little extent to do some basic things. I <strong>don't recommend this</strong> only do this if there are no other options.</p>
<pre><code class="language-bash">printf %.9f &quot;$((10/3))
</code></pre>
<p>The result of this is 3.0000000..064 roughly, which is pretty bad. Bash at its core doesn't support floating-point calculations. But there is good news, we have  <a href="https://en.wikipedia.org/wiki/AWK">awk</a>  and other tools such as  <a href="https://en.wikipedia.org/wiki/Bc_(programming_language)">bc</a>  and others which is planned for the next part in the series. I'll explain awk just for floating-point here, in the next part, I'll cover it in depth.</p>
<pre><code class="language-bash">#!/bin/bash
a=10
b=$(echo | awk -v a=&quot;$a&quot; '{print a/3}')
echo $b 

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625157391350/gHudsNntM4.png" alt="bashs2.png"></p>
<p>WOW! That is to the point, but that was a lot of hassle using echo but printing nothing! HH? OK, you see certain things can get really annoying when things aren't supported natively. Firstly, we use  | to pipe echo command with awk, the echo command doesn't do anything just a way to use awk command in assigning variables here. Then the general syntax of the awk command is <code> awk -options -commands</code>. In this case, we are using -v as an argument and passing in an as a variable which is equal to a, which is stupid and silly but that is what it is, you can name any variable name you want. Then we simply have to use the variable in the print function which generally evaluates the expressions or other operations and returns to the interpreter. And that is how we print the expression, Phew! That took a while to do some silly things, But hey! That's possible though.</p>
<p>That is the basic overview of Arithmetic in Bash, you can also perform logical operations in it which is very easy and can be understood on a quick run-through in the  <a href="https://www.gnu.org/savannah-checkouts/gnu/bash/manual/bash.html#Arithmetic-Expansion">documentation</a>.</p>
<p>I hope you understood the mentioned topics and what are their use cases depending on the requirements. Some topics such as positional parameters, tools and utilities, dictionaries, and some other important aspects of Bash scripting will be covered in the next part. Happy Coding.</p>
