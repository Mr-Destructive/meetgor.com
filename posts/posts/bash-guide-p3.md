{
  "title": "BASH Scripting Guide - PART - 3",
  "status": "published",
  "slug": "bash-guide-p3",
  "date": "2025-04-05T12:33:58Z"
}

<p>Bash or shell won't be much popular and powerful if they didn't have some tools and utilities baked in. But even further they are supported natively in Bash, which just makes every task and challenge quite hassle-free to deal with. In this part of the series, I'll try to cover an overview of some quite powerful and robust tools and utilities in Bash(shell in general) and also some of the advanced topics like dictionaries and positional parameters. Enough talk let's dive in.</p>
<p>The topics to be covered in this part include the following:</p>
<ul>
<li>Hash tables/dictionaries in BASH</li>
<li>Positional parameters</li>
<li>Aliases in BASH</li>
<li>Some Tools and utilities
<ul>
<li>grep/sed/awk</li>
<li>cat/tac/head/tail</li>
<li>cURL</li>
<li>find</li>
<li>bc</li>
<li>wc</li>
</ul>
</li>
</ul>
<h2>Bash dictionaries</h2>
<p>Bash dictionaries or hash tables are just like any other hash tables or keymaps in other programming languages. Bash dictionaries are quite similar to arrays but they have a key instead of the index(0,1,2...) and a value just like arrays. This can be quite useful for storing passwords with emails or usernames or any other way in which a value can be accessed only via a unique key.</p>
<p>To declare a dictionary/ hash table, we can simply write <code>declare -A name</code>, this will declare an empty hash map for us. Further, we can populate the hash map with keys and values using the same syntax as of array just instead of numbers we can also have strings.</p>
<pre><code class="language-bash">#!/bin/bash

declare -A fruits
fruits[&quot;apple&quot;]=&quot;red&quot;
fruits[&quot;mango&quot;]=&quot;yellow&quot;
fruits[&quot;grapes&quot;]=&quot;green&quot;

read -p &quot;Enter the name of fruit : &quot; name 
echo &quot;The fruit is $name and its color is ${fruits[$name]} &quot;

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1626167875237/A2TxYPNoS.png" alt="possh.png"></p>
<p>The above example depicts a way to declare, define and access the key values in a dictionary. The example may look silly but you get the idea. We can also access every key or value using the <code>@</code> variable and access the number of key-value pairs using the <code>#</code> variable just like an array.</p>
<pre><code class="language-bash">#!/bin/bash

declare -A fruits
fruits[&quot;apple&quot;]=&quot;red&quot;
fruits[&quot;mango&quot;]=&quot;yellow&quot;
fruits[&quot;grapes&quot;]=&quot;green&quot;

for i in &quot;${!fruits[@]}&quot;;
do
	echo &quot;$i : ${fruits[$i]}&quot;
done

echo &quot;There are ${#fruits[@]} key-value pairs.&quot;
</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1626171398570/14jJl2eFs.png" alt="possh.png"></p>
<p>OK, this is very tricky they are the same variables but used slightly differently. Firstly in the range-based for loop <code>${!fruits[@]}</code>, focus on the <code>!</code> before the hash table name, this will expand(depict) the hash map's keys. This is used to access every key from the hash table and we can also see <code>#</code> at the beginning of the hash map name as it is used to represent the entire hash map further in the <code>{#fruits[@]}</code> we can also use <code>{#fruits[*]}</code>.  Inside the for loop, <code>i</code> will be the key, and <code>{fruits[$i]}</code> will be the value for that <code>i</code> th key.</p>
<p>Also, you can notice the bash interpreter automatically arranges the map in the alphabetical order of the values and not keys. This is quite a neat little feature that can come in handy a lot of times.</p>
<p>If you want to delete or add any key-value pairs we can do that by the following commands:</p>
<pre><code class="language-bash">#!/bin/bash

declare -A fruits
fruits[&quot;apple&quot;]=&quot;red&quot;
fruits[&quot;mango&quot;]=&quot;yellow&quot;
fruits[&quot;grapes&quot;]=&quot;green&quot;

for i in &quot;${!fruits[@]}&quot;;
do
	echo &quot;$i : ${fruits[$i]}&quot;
done
echo &quot;There are ${#fruits[@]} key-value pairs.&quot;

unset fruits[&quot;mango&quot;] 
echo &quot;There are ${#fruits[@]} key-value pairs.&quot;
fruits[&quot;strawberry&quot;]=&quot;pink&quot;

for i in &quot;${!fruits[@]}&quot;;
do
	echo &quot;$i : ${fruits[$i]}&quot;
done
echo &quot;There are ${#fruits[@]} key-value pairs.&quot;

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1626172120966/yCeXpaM9w.png" alt="possh.png"></p>
<p>The above code example is a bit complex but easy enough to understand. We can delete the key-value pair using the unset command and pass in the key along with the hash map name. We can create a key-value pair by simple command as depicted in the above example. This was a basic overview of hash maps/dictionaries in BASH.</p>
<h2>Positional parameters</h2>
<p>We often use user input from within the script but there is another way to pass in parameters outside of the script using positional parameters. It basically allows us to pass in arguments or parameters from the command prompt/ shell and inside of the script, we can access them via Positional Parameters ( $1, $2, $3....$9, ${10} and so on).</p>
<pre><code class="language-bash">#!/bin/bash

echo &quot;first parameter : &quot; $1
echo &quot;second parameter : &quot; $2
echo &quot;eleventh parameter : &quot; ${11}

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1626159559317/XSyVwkC9U.png" alt="possh.png"></p>
<p>You can see from the above script we have passed the parameters from the command line just after typing the filename. The positional parameter $0 is the file name and from 1 onwards the parameters are optional to run if only your script needs any parameters or input to work with. The numbers are just random and just used for demonstration. The 11th parameter or double-digit parameter starting from 10 onwards, you need to encapsulate the number in {curly braces}<code>${number}</code> because it won't interpret <code>$10</code> or any other number as just <code>$1</code> and print 0.</p>
<p>So we can pass a list of parameters that should be space-separated. We can pass any relevant information such as a string, number, or file names as well.</p>
<p>If we want to access all the parameters passed to the script, we can use <code>@</code> variable. You may know this symbol from the array section of part-II, it is used to access every element in the array. But here it is used to access every parameter passed to the script just behaving like a list of values.</p>
<pre><code class="language-bash">#!/bin/bash

echo &quot;The parameters passed are: &quot; $@

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1626160205206/eH6BD1r_Yu.png" alt="possh.png"></p>
<p>To get the number of the parameters passed to the script, we can use <code>#</code> variable. This is also a variable used in the array section for accessing the number of elements in the array, in this case, the number of parameters in the list.</p>
<pre><code class="language-bash">#!/bin/bash

echo &quot;The parameters passed are: &quot; $#

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1626160630991/jVvJqtCqs.png" alt="possh.png"></p>
<p>Positional parameters allow to not take explicit input from the user from the script. This might not be used for the simple scripts but for administration purposes for the admins and users who know what does the script needs and it allows them to pass in arguments without designing the input system.</p>
<h2>Bash aliases</h2>
<p>Bash aliases are a great way of reducing the command length and making it much easier to type and work with the scripts or any development-related work. Alias is a file called bash_aliases inside the .bashrc folder that contains our shortcut commands, it has a particular order to map certain commands with others.</p>
<p>Let's see what is an alias first and then we'll see how to set it up.</p>
<pre><code class="language-bash">alias cdc='cd C:/Users/acer/Desktop/New\ folder/Code/'
</code></pre>
<p>This will make it viable to just type cdc and I will be in this directory instead of printing all of the jargon. The command we need to use to replace the big command is <code>cdc</code>. The right command is the variable assigned the command and the left or its value is the command to be replaced with it.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1626163036098/lDwlSdiry.gif" alt="shalias.gif"></p>
<p>This command will definitely defer on your machine and OS as the filesystems are different in each major operating system. We can quickly make other such alias or shortcuts so to speak for making the development process faster and efficient.</p>
<p>Now let us see how we set up this environment for bash alias, it's quite straightforward. You need to create a hidden file named &quot;bashrc&quot; i.e the file name will be <code>.bashrc</code>. This file has to be in the root directory (the folder to which bash defaults). I do not mean the <code>root</code> directory in Linux but the repository in which your bash interpreter opens. Once you have created the file put any alias in the file and source the file using the command:</p>
<pre><code class="language-bash">source .bashrc
</code></pre>
<p>And that would do the trick, you can now test your macro or shortcut by opening a new instance of the terminal. If this doesn't work for you, then you can check  <a href="https://opensource.com/article/19/7/bash-aliases">this article</a>  for a broader understanding of the setup.</p>
<h2>Bash tools &amp; utilities</h2>
<p>What would you call BASH without grep or sed man! It's a sad life:( BASH comes with some absolutely powerful and handy tools and utilities such as grep, sed, awk, at, wc, find, tar, gzip, which, make, ping, cURL, wget, ssh, .... my words there is an unstoppable long list. Really they are quite important and lay the foundation for some quite complex tasks. Some web servers can become redundant if some of the tools went missing. Let us find why they are so powerful.</p>
<h3>grep</h3>
<p>GREP or global regular expression print is a tool or command that can find patterns using regular expressions in files/strings or any other piece of data. It's not just printing or searching for the text, besides all that it can also edit the file and store the output in the desired file or any variable by proving some arguments to it. Grep supports Pearl's regular expression as well. There is a lot of customization options and arguments available in grep that can just do anything. It becomes an irreplaceable tool for some complex tasks.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1626179054053/9ttkV-MZZ.png" alt="possh.png"></p>
<p>The above code finds the pattern &quot;more text&quot; in the file specified and prints the line to the screen, but we can modify the output we want, extract the output in a file and do all kinds of wizardry with this tool. This is just a basic, to get started example but trust me it's more than you think, this tool is used widely for web scrapping and pattern matching in quite a lot of use cases.</p>
<h3>sed</h3>
<p>SED or stream editor is another beast in BASH's toolkit, this is just a flawless tool. No words for this. This is a great tool but still underrated. This can actually edit the text inside the terminal, no graphical environment, no interface at all just commands, but it can do what a huge text editor can't! Save time, just edit text without opening anything except a terminal, becomes unbeatable in large files. This is surely a tiny little application that can skyrocket the efficiency development process.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1626179410079/BkvdAqkfDS.png" alt="possh.png"></p>
<p>The above example replaces the word <code>more</code> with <code>less</code> using the sed command if you type 3g or nth line preceded by g, you will replace the word on the nth line only. In this case, only<code>g</code> will replace every occurrence of the word.
This is again a basic example of a sed command, its more if you go deeper, its more than a tool, its kind of a text-editor for wizards ;)</p>
<h3>awk</h3>
<p>awk or Aho, Weinberger, and Kernighan XD are the names of the developers of this application. This is another mind-blowing tool that can programmatically do a lot of stuff. This is like a programming language to a whole new level that can extrapolate and extract data from files and other forms of data. This is quite a great option if you want to quite neatly do something. It has great support libraries and functions that can even perform complex mathematical and scientific operations.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1626180322400/hWkEVhPl_.png" alt="possh.png"></p>
<p>These are the topics for separate articles because it is insufficient to explain everything at once.</p>
<h3>cat / tac / head / tail</h3>
<p>CAT or concatenate is a tool used for printing out files, create files, sorting the contents of files, editing files, and a plethora of stuff. This command is generally used for printing the file but there is more to it like creating a file directly in the terminal, merging two files, and a ton of other operations.</p>
<p>TAC or reverse of CAT is a tool used for everything that CAT can do but in reverse:) This is a wired tool but still quite useful sometimes.</p>
<p>Head is a tool that will print or edit the text in the first 10 lines of the file, it can be used to extrapolate multiple files with similar content.
Tail is a tool that will print or edit the text in the last 10 lines of the file, it can be used just like head but for the last few lines.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1626180451092/Z5VUpIxCm.png" alt="possh.png"></p>
<p>It turns out, you can not only print the first or last 10 lines but n lines by passing the -n as an argument, there is a ton of stuff to discover, this just drives me crazy.</p>
<h3>cURL</h3>
<p>cURL or client URL is a tool that can be used to transfer data via various network protocols. You might not believe but it is used in cars, televisions, routers, and other embedded systems for exchanging relevant data via appropriate protocols.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1626181263235/JPocJwoEd.png" alt="possh.png"></p>
<p>This example depicts how we can fetch data from an API using cURL and extract data in JSON format and use it for relevant tasks.
This is again one of the best utility out there as it becomes quite remarkable and vintage. Despite being almost 30 years old, it shines bright in the tech world.</p>
<h3>find</h3>
<p>Find as the name suggests it is used to find files among the folders and directories in a file system. it becomes quite helpful in complex projects where the directory structure is deep and large.
<img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1626181386566/zpP9Yaom4.png" alt="possh.png"></p>
<p>The command <code>find *.txt</code> finds every txt file available on this directory. As simple as it can get. This is surely looking silly and idiotic but it finds its glory in large and complicated codebases.</p>
<h3>bc</h3>
<p>bc or basic calculator is a utility tool for performing mathematical and arithmetical operations in the terminal, this commands gets integrated with other commands such as awk really well and can be used for further extending the limits of what the command line development can do.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1626182601487/z8X4KeDGG.png" alt="possh.png"></p>
<p>AWW! I could hear the excitement. That just added new dimensions into BASH. Just creativity and resonance to anything is the limit here. I am using  <a href="http://repl.it/">REPL.IT</a>  here for using bash as I do not have it on my windows machine :( But that command is truly insane.</p>
<h3>wc</h3>
<p>wc or word count is a utility tool for counting and analyzing the size or count of characters, words, lines, or files in a given file structure. This is quite a handy tool for monitoring and keeping track of a system, also for general development purposes.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1626182319602/r8UidHV2z.png" alt="possh.png"></p>
<p>The above command prints out the word and lines in the provided file. This command <code>wc</code> can even compute the size of files and even more properties of files.
Those were some of the quite powerful commands, tools, or utilities in BASH/shell. There are plenty of other commands not covered here because this an extremely large topic and even making separate articles or resources there will certainly and surely be some things that will get missed out, that's the beauty Linux or in general Computer Science has.
Ok, that was a lot, but I hope you got some insights for learning more BASH or Linux in general. This is a wide topic and can't be covered entirely in a single article.</p>
<p>Now that is it from this part, everything cannot be covered in any number of parts but at least it will help someone to get started in BASH scripting and its specifications for development. Have a Blast learning BASH. Happy Coding :)</p>
