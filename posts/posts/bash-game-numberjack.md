{
  "title": "Learning BASH by making a Number game",
  "status": "published",
  "slug": "bash-game-numberjack",
  "date": "2025-04-05T12:33:58Z"
}

<h2>Introduction</h2>
<p>OK! Learning BASH can be quite confusing without a proper goal in hand. So this will be a pretty good idea to start learning BASH and have a ton of fun. In this little time, we'll make a Number game which I have designed myself last year in C++, which took about 3 months due to lazy research and wasting time. But I was surprised that I made this game within two hours in BASH. You can refer to the game instructions in this  <a href="https://github.com/Mr-Destructive/NumberJack">repository at Github</a>.</p>
<h2>Concept</h2>
<p>The game will ask a number between 0 and 9 to the user. Then a list of 10 numbers shuffled in a random order will appear from of the user along with another list used for indexing the numbers from the array. The user has to select the index beneath its chosen number to proceed ahead. The game loops until the user has failed to enter the correct index of the number or the time for input was exceeded by 5 seconds. The user will get a point for every successful hit. So that is probably the introduction of the game so, let's dive into the specifications.</p>
<h2>Specifications of the Game in BASH</h2>
<p>The game is a number based which means it will need Arithmetic operators a lot. In fact, we'll need a few complex functions such as shuf. We will very frequently use while and for loops to perform some tasks such as filling and printing array and the game loop. We'll use some flag variables to indicate the current situation in the game and finally some arithmetic on arrays and numbers.</p>
<h2>Script Explanation</h2>
<p>The game is quite simple to understand. You just have to select the number beneath your chosen number within 5 seconds in the shell script. We will create a menu-like display in the terminal by simple echo command and formatting. Before the menu, we will have a while loop that will iterate until the user enters 3 which is stored in variable <code>ch</code> which is initialized to 0 in the beginning so as to enter the loop for the first time. A while loop starts with the do statement and ends at the done statement.</p>
<pre><code class="language-bash">while [ condition ];
do 
# statements
done
</code></pre>
<p>For loop can be different based on the scenario. We'll use a range-based for loop to iterate over a range of numbers using the { } operators. For loop also has do as the beginning of the loop and done as the end of the loop.</p>
<pre><code class="language-bash">for i in {1..5};
do 
#statements
done
</code></pre>
<p>We'll also use some If-else statements just to check for the correct user input and checking the exit status. The if statements have <code>then</code> to start the block and <code>fi</code> to end the if block.</p>
<pre><code class="language-bash">if [ condition ];
then
    #statements
elif
    #statements
else
    #statements
fi
</code></pre>
<p>We use a read statement with the argument -p to have a prompt to the user for some information on the input. The input of choice from the menu i.e 1 to play, 2 for Instructions, and 3 to exit are stored in the variable <code>ch</code>. If the input is 1, the game will start and it will ask the user for the number <code>n</code>, which is the number used throughout the loop until the game is over.</p>
<p>Now we have the number for the rest of the game, we need to generate the list for the user to select the number from. We will have a flag sort of to check if the user has entered the correct number which is <code>c</code>, this will store 0 for correct input(number x) and 1 for incorrect input. It is initialized with 0, again to enter the while loop once before the generation of numbers.</p>
<p>To generate and <strong>shuffle 10 numbers which should not have any repeated numbers</strong>, as it can have multiple numbers which might be unfair also it might happen that the number chosen by the user might not be present due to repetition. So to avoid that mischief of pseudo-random numbers we have to generate distinct 10 numbers from 0 to 9 in this case. For that, we are gonna use a command in BASH called <code>shuf</code> which can create some permutation of the elements in a list/array or a sequence of numbers in an input stream. We are gonna use <code>shuf</code> to generate a random sequence of 10 numbers from 0 to 9 using the command <code>shuf -i 0-9 -n 10</code>.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625748675622/1Li6h3_vX.png" alt="image.png"></p>
<p>You can see it generated a list of shuffled numbers between 0 to 9 so there are 10 numbers. We'll store this in result an array to access and print them later. You can refer to  <a href="https://www.geeksforgeeks.org/shuf-command-in-linux-with-examples/">this</a>  and  <a href="https://www.howtoforge.com/linux-shuf-command/">these</a>  articles for understanding shuf.</p>
<p>The main thing is taken care of, now we need to print the list and also print another list to indicate the index of numbers to the user. We will print the list without a for loop using the <code>@</code> variable. If you are new to BASH and want a bit guide on BASH please do check out my series on  <a href="https://techstructiveblog.hashnode.dev/series/bash-scripting">BASH scripting</a>, I have this all covered. So using <code>@</code> we can print the entire array in BASH.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625749273007/hf_y4Fm53.png" alt="image.png"></p>
<p>To print the lower list of indices, we'll use a range-based for loop i.e it will iterate over a range(in this case between 1 to 10), and assign each element the value of the counter i.e from 1 to 10. We are using <code>r</code> as the shuffled list and <code>a</code> as the indices list. And print this array with the same method.</p>
<p>After the generation and printing of lists are complete, we'll take the input from the user for the index of his/her number. We'll use an argument in read known as timeout, which will give a stop to the input stream after several seconds provided in the argument. In this case, we will use 5 seconds as a timeout for the input of the index. <code>read -t 5 -p &quot;Enter the index of your number : &quot; x </code>
We'll store the input in <code>x</code> variable and access it later for verification.</p>
<p>Next, we will check if the input was done before the timeout or not. For this, if the user input before timeout, we can proceed ahead but if the time was over, then we'll get an exit status above 128 so we use this as a checker for the timeout in the input. I came to this via this  <a href="https://www.linux.org/threads/exit-script-by-timeout-if-delay-of-read-input-in-command-line.15905/">article</a>, really very helpful. We will break the loop and make the flag <code>c</code> as 1 indicating an improper input and thus it'll show &quot;GAME OVER&quot;. But if you were fast enough then we'll check that the index of the shuffled array has your chosen number or not, we used this <code>${r[$(($x))-1]} -eq $n</code> to check for the correct number. Why -1? If you remember indexing in the array by default starts with 0, as we have started the second list from 1 hence every element will become offset by 1 hence to avoid that we'll subtract one to refer to that index.</p>
<p>If the index of the number was equal and correct, well done we'll increment the counter of points <code>p</code> by one and if it was incorrect, the flag will be set to one as previously said and we'll break the loop. After coming out of the loop, we'll check if the status flag <code>c</code> was 1 if yes, then print the GAME OVER and display the points earned. And that is it. Let's take a look at some gameplay :)</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625753634816/CCUD8OD_K.gif" alt="numbjackbash.gif"></p>
<h2>BASH Script</h2>
<pre><code class="language-bash">#!/bin/bash

echo -e &quot;
 NumberJack 
&quot;
ch=0
while [ $ch -ne 3 ];
do
	echo  &quot;  
		 PLAY : Hit 1 and enter.
		 HELP : Hit 2 and enter.
		 EXIT : Hit 3 and enter.
		 &quot;

	read -p &quot;Enter your choice : &quot; ch
	if [ $ch -eq 1 ];then
	x=0
	c=0
	p=0
	read -p &quot;Enter any number between 0 and 9 : &quot; n
	while [ $c -eq 0 ];
	do
		x=11
		r=($(shuf -i 0-9 -n 10))
		echo &quot;${r[@]} &quot;
		for i in {1..10};
		do
			a[$i]=$i	
		done
		echo &quot;${a[@]} &quot;
		read -t 5 -p &quot;Enter the index of your number : &quot; x
		if [[ $? -gt 128 ]]; then 
			c=1
			break
		fi
		if [ ${r[$(($x))-1]} -eq $n ];then
			echo &quot;Great&quot;
			((p=p+1))
		else
			c=1
			break
		fi
	done
	elif [ $ch -eq 2 ];then
		echo &quot;HELP: INSTRUCTIONS TO PLAY THE GAME. &quot;
	else
		break
fi

if [ $c -eq 1 ];then
			echo -e &quot;
GAME OVER
&quot;
			echo &quot;You scored $p points&quot;
fi
		done

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1625753738352/qBDF1PFQG.png" alt="image.png"></p>
<p>This is the final bare-bones script without any help instructions just keeping the script simple. I hope you learned something from the game development in BASH. This is just a fun little project and a cool way of learning certain concepts in BASH such as loops, conditional statements, and arithmetic. Have FUN. Happy CODING :)</p>
