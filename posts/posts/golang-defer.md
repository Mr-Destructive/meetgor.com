{
  "title": "Golang: Defer",
  "status": "published",
  "slug": "golang-defer",
  "date": "2025-04-05T12:33:39Z"
}

<h2>Introduction</h2>
<p>In this part of the series, we will be taking a look at the <code>defer</code> keyword in golang. The defer keyword is used for delaying the function call in a particular block of program(function or a loop).</p>
<h2>Defer Keyword</h2>
<p>The <code>defer</code> keyword is an interesting keyword in golang, it basically holds up the execution of the statement until all the statements around the local scope has been executed. It is basically like a stack holding the execution of statements. You can have multiple defer keywords in the single code block(function or a loop), those will be called by the principle of <strong>first in last out</strong>.</p>
<p>So, let's take a simple example, the syntax is quite simple just add <code>defer</code> before the statement you want to hold up.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
	fmt.Println(&quot;First&quot;)
	defer fmt.Println(&quot;Second Ahhh..&quot;)
	fmt.Println(&quot;Third&quot;)
}
</code></pre>
<pre><code>go run defer.go                                                                                                               

First
Third
Second Ahhh..
</code></pre>
<p>As, you can see the function call <code>Second</code> was executed at the end of all the function calls in the main function. This is because of the defer keyword. It will halt or postpone the calling of a statement/function before all the statements have bee executed in the local scope. This keyword can be stacked for calling different statements at the specific time in the state of the program.</p>
<h3>Multiple defer keyword</h3>
<p>We can understand the <code>defer</code> keyword in a better way if we modify the previous example a bit. We will use multiple defer statements in it for understanding the flow of the program.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func main() {
	fmt.Println(&quot;bag&quot;)
	defer fmt.Println(&quot;book&quot;)
	defer fmt.Println(&quot;cap&quot;)
	fmt.Println(&quot;laptop&quot;)
	defer fmt.Println(&quot;wallet&quot;)
	fmt.Println(&quot;headphones&quot;)
}
</code></pre>
<pre><code>go run defer.go                                                                                                            

bag
laptop
headphones
wallet
cap
book
</code></pre>
<p>Here, we can see that the <code>bag</code> is printed out first, then <code>laptop</code> and then <code>headphones</code>, but after these, we have maintained a stack for calling out the defer statements.</p>
<p>Initially, consider a empty stack <code>[]</code> for storing defer statements. Firstly, the <code>bag</code> statement will be printed. After printing the <code>bag</code> keyword, there is a defer key word, so we add the statement to the stack <code>[&quot;book&quot;]</code>. Just for convenience, ignore the function syntax and the actual statements, just focus on what is printed when. So, we have printed <code>bag</code> and we have stack as <code>[&quot;book&quot;,]</code>. We again encounter a <code>defer</code> keyword, so we push the statement <code>cap</code> into the stack <code>[&quot;book&quot;, &quot;cap&quot;,]</code>. Further, we encounter the normal statement, so we print <code>&quot;laptop&quot;</code>. Next up, we encounter a defer keyword and thereby we push <code>&quot;wallet&quot;</code> into the stack, which then becomes <code>[&quot;book&quot;, &quot;cap&quot;, &quot;wallet&quot;]</code>. The last statement in the main function is <code>&quot;headphones&quot;</code>, so we print <code>&quot;headphones&quot;</code> with the <code>Println</code> function. Since, there are no further statements to execute in the main function, we start poping out the function calls from the stack. Remember last in first out, so we will print the statement which were pushed last into the stack. Since the stack is <code>[&quot;book&quot;, &quot;cap&quot;, &quot;wallet&quot;]</code>, so we will pop out <code>&quot;wallet&quot;</code> from the stack since we have pushed it last in the stack. Thereby we print <code>&quot;wallet&quot;</code> after <code>&quot;headphones&quot;</code> and the stack now becomes <code>[&quot;book&quot;, &quot;cap&quot;]</code>. Next, we have to pop out the second last element that we pushed into the stack which was <code>&quot;cap&quot;</code>, thus we print <code>&quot;cap&quot;</code>. The stack only has one element in the stack and we pop out that as well which is <code>&quot;book&quot;</code>. So, this is how the defer keyword stacks up in golang.</p>
<p><strong>NOTE: The defer keyword calls functions after the execution of all other functions in it's scope but the parameters are evaluated before handed which means just the function is executed later, but the parameters are already loaded</strong></p>
<h3>defer keyword in function call</h3>
<p>The defer keyword is quite similarly used while calling the functions. The <code>fmt.Println</code> is also a function, but writing custom functions gives a different feeling to us.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

func square(x int) int {
	fmt.Println(x * x)
	return x * x
}

func main() {
	// defer keyword in function
	defer square(12)
	defer square(10)
	square(15)
}
</code></pre>
<pre><code>go run defer.go                                                                                                            

225
100
144
</code></pre>
<p>So, we have created the <code>square</code> function and called it thrice and twice we have used the <code>defer</code> keyword. The defer keyword first pushes the function call <code>square(12)</code> to the stack, so the stack is <code>[square(12) ]</code>. Next, we again have the defer keyword, so we push the <code>square(10)</code> to the stack which becomes <code>[square(12) square(10)]</code>. The next statement is immediately called the function <code>square(15)</code> and thereby we end the main function. So, we have to pop the function calls from the stack. The last element <code>square(10)</code> is called first and then the only call to the <code>square(12)</code>. Thereby, the order of the calling becomes <code>225 100 144</code>.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/defer/defer.go">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So, from this post, we were able to understand the defer keyword in golang. We were able to explore how we can delay the calling of a function in a particular scope of the program.</p>
<p>Thank you for reading, if you have any queries, feedback, or questions, you can freely ask me on my social handles. Happy Coding :)</p>
