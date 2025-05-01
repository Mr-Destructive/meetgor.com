---
article_html: "<h2>Introduction</h2>\n<p>OK! Learning BASH can be quite confusing
  without a proper goal in hand. So this will be a pretty good idea to start learning
  BASH and have a ton of fun. In this little time, we'll make a Number game which
  I have designed myself last year in C++, which took about 3 months due to lazy research
  and wasting time. But I was surprised that I made this game within two hours in
  BASH. You can refer to the game instructions in this  <a href=\"https://github.com/Mr-Destructive/NumberJack\">repository
  at Github</a>.</p>\n<h2>Concept</h2>\n<p>The game will ask a number between 0 and
  9 to the user. Then a list of 10 numbers shuffled in a random order will appear
  from of the user along with another list used for indexing the numbers from the
  array. The user has to select the index beneath its chosen number to proceed ahead.
  The game loops until the user has failed to enter the correct index of the number
  or the time for input was exceeded by 5 seconds. The user will get a point for every
  successful hit. So that is probably the introduction of the game so, let's dive
  into the specifications.</p>\n<h2>Specifications of the Game in BASH</h2>\n<p>The
  game is a number based which means it will need Arithmetic operators a lot. In fact,
  we'll need a few complex functions such as shuf. We will very frequently use while
  and for loops to perform some tasks such as filling and printing array and the game
  loop. We'll use some flag variables to indicate the current situation in the game
  and finally some arithmetic on arrays and numbers.</p>\n<h2>Script Explanation</h2>\n<p>The
  game is quite simple to understand. You just have to select the number beneath your
  chosen number within 5 seconds in the shell script. We will create a menu-like display
  in the terminal by simple echo command and formatting. Before the menu, we will
  have a while loop that will iterate until the user enters 3 which is stored in variable
  <code>ch</code> which is initialized to 0 in the beginning so as to enter the loop
  for the first time. A while loop starts with the do statement and ends at the done
  statement.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">while</span> <span class=\"o\">[</span> condition <span class=\"o\">]</span><span
  class=\"p\">;</span>\n<span class=\"k\">do</span> \n<span class=\"c1\"># statements</span>\n<span
  class=\"k\">done</span>\n</pre></div>\n\n</pre>\n\n<p>For loop can be different
  based on the scenario. We'll use a range-based for loop to iterate over a range
  of numbers using the { } operators. For loop also has do as the beginning of the
  loop and done as the end of the loop.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">for</span> i <span class=\"k\">in</span> <span class=\"o\">{</span><span
  class=\"m\">1</span>..5<span class=\"o\">}</span><span class=\"p\">;</span>\n<span
  class=\"k\">do</span> \n<span class=\"c1\">#statements</span>\n<span class=\"k\">done</span>\n</pre></div>\n\n</pre>\n\n<p>We'll
  also use some If-else statements just to check for the correct user input and checking
  the exit status. The if statements have <code>then</code> to start the block and
  <code>fi</code> to end the if block.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">if</span> <span class=\"o\">[</span> condition <span class=\"o\">]</span><span
  class=\"p\">;</span>\n<span class=\"k\">then</span>\n    <span class=\"c1\">#statements</span>\n<span
  class=\"k\">elif</span>\n    <span class=\"c1\">#statements</span>\n<span class=\"k\">else</span>\n
  \   <span class=\"c1\">#statements</span>\n<span class=\"k\">fi</span>\n</pre></div>\n\n</pre>\n\n<p>We
  use a read statement with the argument -p to have a prompt to the user for some
  information on the input. The input of choice from the menu i.e 1 to play, 2 for
  Instructions, and 3 to exit are stored in the variable <code>ch</code>. If the input
  is 1, the game will start and it will ask the user for the number <code>n</code>,
  which is the number used throughout the loop until the game is over.</p>\n<p>Now
  we have the number for the rest of the game, we need to generate the list for the
  user to select the number from. We will have a flag sort of to check if the user
  has entered the correct number which is <code>c</code>, this will store 0 for correct
  input(number x) and 1 for incorrect input. It is initialized with 0, again to enter
  the while loop once before the generation of numbers.</p>\n<p>To generate and <strong>shuffle
  10 numbers which should not have any repeated numbers</strong>, as it can have multiple
  numbers which might be unfair also it might happen that the number chosen by the
  user might not be present due to repetition. So to avoid that mischief of pseudo-random
  numbers we have to generate distinct 10 numbers from 0 to 9 in this case. For that,
  we are gonna use a command in BASH called <code>shuf</code> which can create some
  permutation of the elements in a list/array or a sequence of numbers in an input
  stream. We are gonna use <code>shuf</code> to generate a random sequence of 10 numbers
  from 0 to 9 using the command <code>shuf -i 0-9 -n 10</code>.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625748675622/1Li6h3_vX.png\"
  alt=\"image.png\" /></p>\n<p>You can see it generated a list of shuffled numbers
  between 0 to 9 so there are 10 numbers. We'll store this in result an array to access
  and print them later. You can refer to  <a href=\"https://www.geeksforgeeks.org/shuf-command-in-linux-with-examples/\">this</a>
  \ and  <a href=\"https://www.howtoforge.com/linux-shuf-command/\">these</a>  articles
  for understanding shuf.</p>\n<p>The main thing is taken care of, now we need to
  print the list and also print another list to indicate the index of numbers to the
  user. We will print the list without a for loop using the <code>@</code> variable.
  If you are new to BASH and want a bit guide on BASH please do check out my series
  on  <a href=\"https://techstructiveblog.hashnode.dev/series/bash-scripting\">BASH
  scripting</a>, I have this all covered. So using <code>@</code> we can print the
  entire array in BASH.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625749273007/hf_y4Fm53.png\"
  alt=\"image.png\" /></p>\n<p>To print the lower list of indices, we'll use a range-based
  for loop i.e it will iterate over a range(in this case between 1 to 10), and assign
  each element the value of the counter i.e from 1 to 10. We are using <code>r</code>
  as the shuffled list and <code>a</code> as the indices list. And print this array
  with the same method.</p>\n<p>After the generation and printing of lists are complete,
  we'll take the input from the user for the index of his/her number. We'll use an
  argument in read known as timeout, which will give a stop to the input stream after
  several seconds provided in the argument. In this case, we will use 5 seconds as
  a timeout for the input of the index. <code>read -t 5 -p &quot;Enter the index of
  your number : &quot; x </code>\nWe'll store the input in <code>x</code> variable
  and access it later for verification.</p>\n<p>Next, we will check if the input was
  done before the timeout or not. For this, if the user input before timeout, we can
  proceed ahead but if the time was over, then we'll get an exit status above 128
  so we use this as a checker for the timeout in the input. I came to this via this
  \ <a href=\"https://www.linux.org/threads/exit-script-by-timeout-if-delay-of-read-input-in-command-line.15905/\">article</a>,
  really very helpful. We will break the loop and make the flag <code>c</code> as
  1 indicating an improper input and thus it'll show &quot;GAME OVER&quot;. But if
  you were fast enough then we'll check that the index of the shuffled array has your
  chosen number or not, we used this <code>${r[$(($x))-1]} -eq $n</code> to check
  for the correct number. Why -1? If you remember indexing in the array by default
  starts with 0, as we have started the second list from 1 hence every element will
  become offset by 1 hence to avoid that we'll subtract one to refer to that index.</p>\n<p>If
  the index of the number was equal and correct, well done we'll increment the counter
  of points <code>p</code> by one and if it was incorrect, the flag will be set to
  one as previously said and we'll break the loop. After coming out of the loop, we'll
  check if the status flag <code>c</code> was 1 if yes, then print the GAME OVER and
  display the points earned. And that is it. Let's take a look at some gameplay :)</p>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625753634816/CCUD8OD_K.gif\"
  alt=\"numbjackbash.gif\" /></p>\n<h2>BASH Script</h2>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">echo</span> -e <span class=\"s2\">&quot;\\n
  NumberJack \\n&quot;</span>\n<span class=\"nv\">ch</span><span class=\"o\">=</span><span
  class=\"m\">0</span>\n<span class=\"k\">while</span> <span class=\"o\">[</span>
  <span class=\"nv\">$ch</span> -ne <span class=\"m\">3</span> <span class=\"o\">]</span><span
  class=\"p\">;</span>\n<span class=\"k\">do</span>\n\t<span class=\"nb\">echo</span>
  \ <span class=\"s2\">&quot;  </span>\n<span class=\"s2\">\t\t PLAY : Hit 1 and enter.</span>\n<span
  class=\"s2\">\t\t HELP : Hit 2 and enter.</span>\n<span class=\"s2\">\t\t EXIT :
  Hit 3 and enter.</span>\n<span class=\"s2\">\t\t &quot;</span>\n\n\t<span class=\"nb\">read</span>
  -p <span class=\"s2\">&quot;Enter your choice : &quot;</span> ch\n\t<span class=\"k\">if</span>
  <span class=\"o\">[</span> <span class=\"nv\">$ch</span> -eq <span class=\"m\">1</span>
  <span class=\"o\">]</span><span class=\"p\">;</span><span class=\"k\">then</span>\n\t<span
  class=\"nv\">x</span><span class=\"o\">=</span><span class=\"m\">0</span>\n\t<span
  class=\"nv\">c</span><span class=\"o\">=</span><span class=\"m\">0</span>\n\t<span
  class=\"nv\">p</span><span class=\"o\">=</span><span class=\"m\">0</span>\n\t<span
  class=\"nb\">read</span> -p <span class=\"s2\">&quot;Enter any number between 0
  and 9 : &quot;</span> n\n\t<span class=\"k\">while</span> <span class=\"o\">[</span>
  <span class=\"nv\">$c</span> -eq <span class=\"m\">0</span> <span class=\"o\">]</span><span
  class=\"p\">;</span>\n\t<span class=\"k\">do</span>\n\t\t<span class=\"nv\">x</span><span
  class=\"o\">=</span><span class=\"m\">11</span>\n\t\t<span class=\"nv\">r</span><span
  class=\"o\">=(</span><span class=\"k\">$(</span>shuf -i <span class=\"m\">0</span>-9
  -n <span class=\"m\">10</span><span class=\"k\">)</span><span class=\"o\">)</span>\n\t\t<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"si\">${</span><span
  class=\"nv\">r</span><span class=\"p\">[@]</span><span class=\"si\">}</span><span
  class=\"s2\"> &quot;</span>\n\t\t<span class=\"k\">for</span> i <span class=\"k\">in</span>
  <span class=\"o\">{</span><span class=\"m\">1</span>..10<span class=\"o\">}</span><span
  class=\"p\">;</span>\n\t\t<span class=\"k\">do</span>\n\t\t\ta<span class=\"o\">[</span><span
  class=\"nv\">$i</span><span class=\"o\">]=</span><span class=\"nv\">$i</span>\t\n\t\t<span
  class=\"k\">done</span>\n\t\t<span class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span
  class=\"si\">${</span><span class=\"nv\">a</span><span class=\"p\">[@]</span><span
  class=\"si\">}</span><span class=\"s2\"> &quot;</span>\n\t\t<span class=\"nb\">read</span>
  -t <span class=\"m\">5</span> -p <span class=\"s2\">&quot;Enter the index of your
  number : &quot;</span> x\n\t\t<span class=\"k\">if</span> <span class=\"o\">[[</span>
  <span class=\"nv\">$?</span> -gt <span class=\"m\">128</span> <span class=\"o\">]]</span><span
  class=\"p\">;</span> <span class=\"k\">then</span> \n\t\t\t<span class=\"nv\">c</span><span
  class=\"o\">=</span><span class=\"m\">1</span>\n\t\t\t<span class=\"nb\">break</span>\n\t\t<span
  class=\"k\">fi</span>\n\t\t<span class=\"k\">if</span> <span class=\"o\">[</span>
  <span class=\"si\">${</span><span class=\"nv\">r</span><span class=\"p\">[</span><span
  class=\"k\">$((</span><span class=\"nv\">$x</span><span class=\"k\">))</span><span
  class=\"p\">-1]</span><span class=\"si\">}</span> -eq <span class=\"nv\">$n</span>
  <span class=\"o\">]</span><span class=\"p\">;</span><span class=\"k\">then</span>\n\t\t\t<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;Great&quot;</span>\n\t\t\t<span
  class=\"o\">((</span><span class=\"nv\">p</span><span class=\"o\">=</span>p+1<span
  class=\"o\">))</span>\n\t\t<span class=\"k\">else</span>\n\t\t\t<span class=\"nv\">c</span><span
  class=\"o\">=</span><span class=\"m\">1</span>\n\t\t\t<span class=\"nb\">break</span>\n\t\t<span
  class=\"k\">fi</span>\n\t<span class=\"k\">done</span>\n\t<span class=\"k\">elif</span>
  <span class=\"o\">[</span> <span class=\"nv\">$ch</span> -eq <span class=\"m\">2</span>
  <span class=\"o\">]</span><span class=\"p\">;</span><span class=\"k\">then</span>\n\t\t<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;HELP: INSTRUCTIONS TO PLAY THE
  GAME. &quot;</span>\n\t<span class=\"k\">else</span>\n\t\t<span class=\"nb\">break</span>\n<span
  class=\"k\">fi</span>\n\n<span class=\"k\">if</span> <span class=\"o\">[</span>
  <span class=\"nv\">$c</span> -eq <span class=\"m\">1</span> <span class=\"o\">]</span><span
  class=\"p\">;</span><span class=\"k\">then</span>\n\t\t\t<span class=\"nb\">echo</span>
  -e <span class=\"s2\">&quot;\\nGAME OVER\\n&quot;</span>\n\t\t\t<span class=\"nb\">echo</span>
  <span class=\"s2\">&quot;You scored </span><span class=\"nv\">$p</span><span class=\"s2\">
  points&quot;</span>\n<span class=\"k\">fi</span>\n\t\t<span class=\"k\">done</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625753738352/qBDF1PFQG.png\"
  alt=\"image.png\" /></p>\n<p>This is the final bare-bones script without any help
  instructions just keeping the script simple. I hope you learned something from the
  game development in BASH. This is just a fun little project and a cool way of learning
  certain concepts in BASH such as loops, conditional statements, and arithmetic.
  Have FUN. Happy CODING :)</p>\n"
cover: ''
date: 2021-07-08
datetime: 2021-07-08 00:00:00+00:00
description: 'OK The game will ask a number between 0 and 9 to the user. Then a list
  of 10 numbers shuffled in a random order will appear from of the user along with
  another '
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-07-08-Learn-BASH-Number-Game.md
html: "<h2>Introduction</h2>\n<p>OK! Learning BASH can be quite confusing without
  a proper goal in hand. So this will be a pretty good idea to start learning BASH
  and have a ton of fun. In this little time, we'll make a Number game which I have
  designed myself last year in C++, which took about 3 months due to lazy research
  and wasting time. But I was surprised that I made this game within two hours in
  BASH. You can refer to the game instructions in this  <a href=\"https://github.com/Mr-Destructive/NumberJack\">repository
  at Github</a>.</p>\n<h2>Concept</h2>\n<p>The game will ask a number between 0 and
  9 to the user. Then a list of 10 numbers shuffled in a random order will appear
  from of the user along with another list used for indexing the numbers from the
  array. The user has to select the index beneath its chosen number to proceed ahead.
  The game loops until the user has failed to enter the correct index of the number
  or the time for input was exceeded by 5 seconds. The user will get a point for every
  successful hit. So that is probably the introduction of the game so, let's dive
  into the specifications.</p>\n<h2>Specifications of the Game in BASH</h2>\n<p>The
  game is a number based which means it will need Arithmetic operators a lot. In fact,
  we'll need a few complex functions such as shuf. We will very frequently use while
  and for loops to perform some tasks such as filling and printing array and the game
  loop. We'll use some flag variables to indicate the current situation in the game
  and finally some arithmetic on arrays and numbers.</p>\n<h2>Script Explanation</h2>\n<p>The
  game is quite simple to understand. You just have to select the number beneath your
  chosen number within 5 seconds in the shell script. We will create a menu-like display
  in the terminal by simple echo command and formatting. Before the menu, we will
  have a while loop that will iterate until the user enters 3 which is stored in variable
  <code>ch</code> which is initialized to 0 in the beginning so as to enter the loop
  for the first time. A while loop starts with the do statement and ends at the done
  statement.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">while</span> <span class=\"o\">[</span> condition <span class=\"o\">]</span><span
  class=\"p\">;</span>\n<span class=\"k\">do</span> \n<span class=\"c1\"># statements</span>\n<span
  class=\"k\">done</span>\n</pre></div>\n\n</pre>\n\n<p>For loop can be different
  based on the scenario. We'll use a range-based for loop to iterate over a range
  of numbers using the { } operators. For loop also has do as the beginning of the
  loop and done as the end of the loop.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">for</span> i <span class=\"k\">in</span> <span class=\"o\">{</span><span
  class=\"m\">1</span>..5<span class=\"o\">}</span><span class=\"p\">;</span>\n<span
  class=\"k\">do</span> \n<span class=\"c1\">#statements</span>\n<span class=\"k\">done</span>\n</pre></div>\n\n</pre>\n\n<p>We'll
  also use some If-else statements just to check for the correct user input and checking
  the exit status. The if statements have <code>then</code> to start the block and
  <code>fi</code> to end the if block.</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"k\">if</span> <span class=\"o\">[</span> condition <span class=\"o\">]</span><span
  class=\"p\">;</span>\n<span class=\"k\">then</span>\n    <span class=\"c1\">#statements</span>\n<span
  class=\"k\">elif</span>\n    <span class=\"c1\">#statements</span>\n<span class=\"k\">else</span>\n
  \   <span class=\"c1\">#statements</span>\n<span class=\"k\">fi</span>\n</pre></div>\n\n</pre>\n\n<p>We
  use a read statement with the argument -p to have a prompt to the user for some
  information on the input. The input of choice from the menu i.e 1 to play, 2 for
  Instructions, and 3 to exit are stored in the variable <code>ch</code>. If the input
  is 1, the game will start and it will ask the user for the number <code>n</code>,
  which is the number used throughout the loop until the game is over.</p>\n<p>Now
  we have the number for the rest of the game, we need to generate the list for the
  user to select the number from. We will have a flag sort of to check if the user
  has entered the correct number which is <code>c</code>, this will store 0 for correct
  input(number x) and 1 for incorrect input. It is initialized with 0, again to enter
  the while loop once before the generation of numbers.</p>\n<p>To generate and <strong>shuffle
  10 numbers which should not have any repeated numbers</strong>, as it can have multiple
  numbers which might be unfair also it might happen that the number chosen by the
  user might not be present due to repetition. So to avoid that mischief of pseudo-random
  numbers we have to generate distinct 10 numbers from 0 to 9 in this case. For that,
  we are gonna use a command in BASH called <code>shuf</code> which can create some
  permutation of the elements in a list/array or a sequence of numbers in an input
  stream. We are gonna use <code>shuf</code> to generate a random sequence of 10 numbers
  from 0 to 9 using the command <code>shuf -i 0-9 -n 10</code>.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625748675622/1Li6h3_vX.png\"
  alt=\"image.png\" /></p>\n<p>You can see it generated a list of shuffled numbers
  between 0 to 9 so there are 10 numbers. We'll store this in result an array to access
  and print them later. You can refer to  <a href=\"https://www.geeksforgeeks.org/shuf-command-in-linux-with-examples/\">this</a>
  \ and  <a href=\"https://www.howtoforge.com/linux-shuf-command/\">these</a>  articles
  for understanding shuf.</p>\n<p>The main thing is taken care of, now we need to
  print the list and also print another list to indicate the index of numbers to the
  user. We will print the list without a for loop using the <code>@</code> variable.
  If you are new to BASH and want a bit guide on BASH please do check out my series
  on  <a href=\"https://techstructiveblog.hashnode.dev/series/bash-scripting\">BASH
  scripting</a>, I have this all covered. So using <code>@</code> we can print the
  entire array in BASH.</p>\n<p><img src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625749273007/hf_y4Fm53.png\"
  alt=\"image.png\" /></p>\n<p>To print the lower list of indices, we'll use a range-based
  for loop i.e it will iterate over a range(in this case between 1 to 10), and assign
  each element the value of the counter i.e from 1 to 10. We are using <code>r</code>
  as the shuffled list and <code>a</code> as the indices list. And print this array
  with the same method.</p>\n<p>After the generation and printing of lists are complete,
  we'll take the input from the user for the index of his/her number. We'll use an
  argument in read known as timeout, which will give a stop to the input stream after
  several seconds provided in the argument. In this case, we will use 5 seconds as
  a timeout for the input of the index. <code>read -t 5 -p &quot;Enter the index of
  your number : &quot; x </code>\nWe'll store the input in <code>x</code> variable
  and access it later for verification.</p>\n<p>Next, we will check if the input was
  done before the timeout or not. For this, if the user input before timeout, we can
  proceed ahead but if the time was over, then we'll get an exit status above 128
  so we use this as a checker for the timeout in the input. I came to this via this
  \ <a href=\"https://www.linux.org/threads/exit-script-by-timeout-if-delay-of-read-input-in-command-line.15905/\">article</a>,
  really very helpful. We will break the loop and make the flag <code>c</code> as
  1 indicating an improper input and thus it'll show &quot;GAME OVER&quot;. But if
  you were fast enough then we'll check that the index of the shuffled array has your
  chosen number or not, we used this <code>${r[$(($x))-1]} -eq $n</code> to check
  for the correct number. Why -1? If you remember indexing in the array by default
  starts with 0, as we have started the second list from 1 hence every element will
  become offset by 1 hence to avoid that we'll subtract one to refer to that index.</p>\n<p>If
  the index of the number was equal and correct, well done we'll increment the counter
  of points <code>p</code> by one and if it was incorrect, the flag will be set to
  one as previously said and we'll break the loop. After coming out of the loop, we'll
  check if the status flag <code>c</code> was 1 if yes, then print the GAME OVER and
  display the points earned. And that is it. Let's take a look at some gameplay :)</p>\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625753634816/CCUD8OD_K.gif\"
  alt=\"numbjackbash.gif\" /></p>\n<h2>BASH Script</h2>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>bash</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"ch\">#!/bin/bash</span>\n\n<span class=\"nb\">echo</span> -e <span class=\"s2\">&quot;\\n
  NumberJack \\n&quot;</span>\n<span class=\"nv\">ch</span><span class=\"o\">=</span><span
  class=\"m\">0</span>\n<span class=\"k\">while</span> <span class=\"o\">[</span>
  <span class=\"nv\">$ch</span> -ne <span class=\"m\">3</span> <span class=\"o\">]</span><span
  class=\"p\">;</span>\n<span class=\"k\">do</span>\n\t<span class=\"nb\">echo</span>
  \ <span class=\"s2\">&quot;  </span>\n<span class=\"s2\">\t\t PLAY : Hit 1 and enter.</span>\n<span
  class=\"s2\">\t\t HELP : Hit 2 and enter.</span>\n<span class=\"s2\">\t\t EXIT :
  Hit 3 and enter.</span>\n<span class=\"s2\">\t\t &quot;</span>\n\n\t<span class=\"nb\">read</span>
  -p <span class=\"s2\">&quot;Enter your choice : &quot;</span> ch\n\t<span class=\"k\">if</span>
  <span class=\"o\">[</span> <span class=\"nv\">$ch</span> -eq <span class=\"m\">1</span>
  <span class=\"o\">]</span><span class=\"p\">;</span><span class=\"k\">then</span>\n\t<span
  class=\"nv\">x</span><span class=\"o\">=</span><span class=\"m\">0</span>\n\t<span
  class=\"nv\">c</span><span class=\"o\">=</span><span class=\"m\">0</span>\n\t<span
  class=\"nv\">p</span><span class=\"o\">=</span><span class=\"m\">0</span>\n\t<span
  class=\"nb\">read</span> -p <span class=\"s2\">&quot;Enter any number between 0
  and 9 : &quot;</span> n\n\t<span class=\"k\">while</span> <span class=\"o\">[</span>
  <span class=\"nv\">$c</span> -eq <span class=\"m\">0</span> <span class=\"o\">]</span><span
  class=\"p\">;</span>\n\t<span class=\"k\">do</span>\n\t\t<span class=\"nv\">x</span><span
  class=\"o\">=</span><span class=\"m\">11</span>\n\t\t<span class=\"nv\">r</span><span
  class=\"o\">=(</span><span class=\"k\">$(</span>shuf -i <span class=\"m\">0</span>-9
  -n <span class=\"m\">10</span><span class=\"k\">)</span><span class=\"o\">)</span>\n\t\t<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span class=\"si\">${</span><span
  class=\"nv\">r</span><span class=\"p\">[@]</span><span class=\"si\">}</span><span
  class=\"s2\"> &quot;</span>\n\t\t<span class=\"k\">for</span> i <span class=\"k\">in</span>
  <span class=\"o\">{</span><span class=\"m\">1</span>..10<span class=\"o\">}</span><span
  class=\"p\">;</span>\n\t\t<span class=\"k\">do</span>\n\t\t\ta<span class=\"o\">[</span><span
  class=\"nv\">$i</span><span class=\"o\">]=</span><span class=\"nv\">$i</span>\t\n\t\t<span
  class=\"k\">done</span>\n\t\t<span class=\"nb\">echo</span> <span class=\"s2\">&quot;</span><span
  class=\"si\">${</span><span class=\"nv\">a</span><span class=\"p\">[@]</span><span
  class=\"si\">}</span><span class=\"s2\"> &quot;</span>\n\t\t<span class=\"nb\">read</span>
  -t <span class=\"m\">5</span> -p <span class=\"s2\">&quot;Enter the index of your
  number : &quot;</span> x\n\t\t<span class=\"k\">if</span> <span class=\"o\">[[</span>
  <span class=\"nv\">$?</span> -gt <span class=\"m\">128</span> <span class=\"o\">]]</span><span
  class=\"p\">;</span> <span class=\"k\">then</span> \n\t\t\t<span class=\"nv\">c</span><span
  class=\"o\">=</span><span class=\"m\">1</span>\n\t\t\t<span class=\"nb\">break</span>\n\t\t<span
  class=\"k\">fi</span>\n\t\t<span class=\"k\">if</span> <span class=\"o\">[</span>
  <span class=\"si\">${</span><span class=\"nv\">r</span><span class=\"p\">[</span><span
  class=\"k\">$((</span><span class=\"nv\">$x</span><span class=\"k\">))</span><span
  class=\"p\">-1]</span><span class=\"si\">}</span> -eq <span class=\"nv\">$n</span>
  <span class=\"o\">]</span><span class=\"p\">;</span><span class=\"k\">then</span>\n\t\t\t<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;Great&quot;</span>\n\t\t\t<span
  class=\"o\">((</span><span class=\"nv\">p</span><span class=\"o\">=</span>p+1<span
  class=\"o\">))</span>\n\t\t<span class=\"k\">else</span>\n\t\t\t<span class=\"nv\">c</span><span
  class=\"o\">=</span><span class=\"m\">1</span>\n\t\t\t<span class=\"nb\">break</span>\n\t\t<span
  class=\"k\">fi</span>\n\t<span class=\"k\">done</span>\n\t<span class=\"k\">elif</span>
  <span class=\"o\">[</span> <span class=\"nv\">$ch</span> -eq <span class=\"m\">2</span>
  <span class=\"o\">]</span><span class=\"p\">;</span><span class=\"k\">then</span>\n\t\t<span
  class=\"nb\">echo</span> <span class=\"s2\">&quot;HELP: INSTRUCTIONS TO PLAY THE
  GAME. &quot;</span>\n\t<span class=\"k\">else</span>\n\t\t<span class=\"nb\">break</span>\n<span
  class=\"k\">fi</span>\n\n<span class=\"k\">if</span> <span class=\"o\">[</span>
  <span class=\"nv\">$c</span> -eq <span class=\"m\">1</span> <span class=\"o\">]</span><span
  class=\"p\">;</span><span class=\"k\">then</span>\n\t\t\t<span class=\"nb\">echo</span>
  -e <span class=\"s2\">&quot;\\nGAME OVER\\n&quot;</span>\n\t\t\t<span class=\"nb\">echo</span>
  <span class=\"s2\">&quot;You scored </span><span class=\"nv\">$p</span><span class=\"s2\">
  points&quot;</span>\n<span class=\"k\">fi</span>\n\t\t<span class=\"k\">done</span>\n</pre></div>\n\n</pre>\n\n<p><img
  src=\"https://cdn.hashnode.com/res/hashnode/image/upload/v1625753738352/qBDF1PFQG.png\"
  alt=\"image.png\" /></p>\n<p>This is the final bare-bones script without any help
  instructions just keeping the script simple. I hope you learned something from the
  game development in BASH. This is just a fun little project and a cool way of learning
  certain concepts in BASH such as loops, conditional statements, and arithmetic.
  Have FUN. Happy CODING :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/bash-script-game.webp
long_description: OK The game will ask a number between 0 and 9 to the user. Then
  a list of 10 numbers shuffled in a random order will appear from of the user along
  with another list used for indexing the numbers from the array. The user has to
  select the index beneat
now: 2025-05-01 18:11:33.312433
path: blog/posts/2021-07-08-Learn-BASH-Number-Game.md
prevnext: null
slug: bash-game-numberjack
status: published
subtitle: Understaning BASH concepts by making a number game
tags:
- bash
templateKey: blog-post
title: Learning BASH by making a Number game
today: 2025-05-01
---

## Introduction

OK! Learning BASH can be quite confusing without a proper goal in hand. So this will be a pretty good idea to start learning BASH and have a ton of fun. In this little time, we'll make a Number game which I have designed myself last year in C++, which took about 3 months due to lazy research and wasting time. But I was surprised that I made this game within two hours in BASH. You can refer to the game instructions in this  [repository at Github](https://github.com/Mr-Destructive/NumberJack).

## Concept 
The game will ask a number between 0 and 9 to the user. Then a list of 10 numbers shuffled in a random order will appear from of the user along with another list used for indexing the numbers from the array. The user has to select the index beneath its chosen number to proceed ahead. The game loops until the user has failed to enter the correct index of the number or the time for input was exceeded by 5 seconds. The user will get a point for every successful hit. So that is probably the introduction of the game so, let's dive into the specifications.

## Specifications of the Game in BASH
The game is a number based which means it will need Arithmetic operators a lot. In fact, we'll need a few complex functions such as shuf. We will very frequently use while and for loops to perform some tasks such as filling and printing array and the game loop. We'll use some flag variables to indicate the current situation in the game and finally some arithmetic on arrays and numbers. 

## Script Explanation
The game is quite simple to understand. You just have to select the number beneath your chosen number within 5 seconds in the shell script. We will create a menu-like display in the terminal by simple echo command and formatting. Before the menu, we will have a while loop that will iterate until the user enters 3 which is stored in variable `ch` which is initialized to 0 in the beginning so as to enter the loop for the first time. A while loop starts with the do statement and ends at the done statement.
```bash
while [ condition ];
do 
# statements
done
```
For loop can be different based on the scenario. We'll use a range-based for loop to iterate over a range of numbers using the { } operators. For loop also has do as the beginning of the loop and done as the end of the loop.

```bash
for i in {1..5};
do 
#statements
done
```
We'll also use some If-else statements just to check for the correct user input and checking the exit status. The if statements have `then` to start the block and `fi` to end the if block. 
```bash
if [ condition ];
then
    #statements
elif
    #statements
else
    #statements
fi
```


We use a read statement with the argument -p to have a prompt to the user for some information on the input. The input of choice from the menu i.e 1 to play, 2 for Instructions, and 3 to exit are stored in the variable `ch`. If the input is 1, the game will start and it will ask the user for the number `n`, which is the number used throughout the loop until the game is over. 

Now we have the number for the rest of the game, we need to generate the list for the user to select the number from. We will have a flag sort of to check if the user has entered the correct number which is `c`, this will store 0 for correct input(number x) and 1 for incorrect input. It is initialized with 0, again to enter the while loop once before the generation of numbers. 

To generate and **shuffle 10 numbers which should not have any repeated numbers**, as it can have multiple numbers which might be unfair also it might happen that the number chosen by the user might not be present due to repetition. So to avoid that mischief of pseudo-random numbers we have to generate distinct 10 numbers from 0 to 9 in this case. For that, we are gonna use a command in BASH called `shuf` which can create some permutation of the elements in a list/array or a sequence of numbers in an input stream. We are gonna use `shuf` to generate a random sequence of 10 numbers from 0 to 9 using the command `shuf -i 0-9 -n 10`. 

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625748675622/1Li6h3_vX.png)

You can see it generated a list of shuffled numbers between 0 to 9 so there are 10 numbers. We'll store this in result an array to access and print them later. You can refer to  [this](https://www.geeksforgeeks.org/shuf-command-in-linux-with-examples/)  and  [these](https://www.howtoforge.com/linux-shuf-command/)  articles for understanding shuf.  

The main thing is taken care of, now we need to print the list and also print another list to indicate the index of numbers to the user. We will print the list without a for loop using the `@` variable. If you are new to BASH and want a bit guide on BASH please do check out my series on  [BASH scripting](https://techstructiveblog.hashnode.dev/series/bash-scripting), I have this all covered. So using `@` we can print the entire array in BASH. 

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625749273007/hf_y4Fm53.png)

To print the lower list of indices, we'll use a range-based for loop i.e it will iterate over a range(in this case between 1 to 10), and assign each element the value of the counter i.e from 1 to 10. We are using `r` as the shuffled list and `a` as the indices list. And print this array with the same method.

After the generation and printing of lists are complete, we'll take the input from the user for the index of his/her number. We'll use an argument in read known as timeout, which will give a stop to the input stream after several seconds provided in the argument. In this case, we will use 5 seconds as a timeout for the input of the index. `read -t 5 -p "Enter the index of your number : " x `
We'll store the input in `x` variable and access it later for verification. 

Next, we will check if the input was done before the timeout or not. For this, if the user input before timeout, we can proceed ahead but if the time was over, then we'll get an exit status above 128 so we use this as a checker for the timeout in the input. I came to this via this  [article](https://www.linux.org/threads/exit-script-by-timeout-if-delay-of-read-input-in-command-line.15905/), really very helpful. We will break the loop and make the flag `c` as 1 indicating an improper input and thus it'll show "GAME OVER". But if you were fast enough then we'll check that the index of the shuffled array has your chosen number or not, we used this `${r[$(($x))-1]} -eq $n` to check for the correct number. Why -1? If you remember indexing in the array by default starts with 0, as we have started the second list from 1 hence every element will become offset by 1 hence to avoid that we'll subtract one to refer to that index. 

If the index of the number was equal and correct, well done we'll increment the counter of points `p` by one and if it was incorrect, the flag will be set to one as previously said and we'll break the loop. After coming out of the loop, we'll check if the status flag `c` was 1 if yes, then print the GAME OVER and display the points earned. And that is it. Let's take a look at some gameplay :)

![numbjackbash.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1625753634816/CCUD8OD_K.gif)

## BASH Script
```bash
#!/bin/bash

echo -e "\n NumberJack \n"
ch=0
while [ $ch -ne 3 ];
do
	echo  "  
		 PLAY : Hit 1 and enter.
		 HELP : Hit 2 and enter.
		 EXIT : Hit 3 and enter.
		 "

	read -p "Enter your choice : " ch
	if [ $ch -eq 1 ];then
	x=0
	c=0
	p=0
	read -p "Enter any number between 0 and 9 : " n
	while [ $c -eq 0 ];
	do
		x=11
		r=($(shuf -i 0-9 -n 10))
		echo "${r[@]} "
		for i in {1..10};
		do
			a[$i]=$i	
		done
		echo "${a[@]} "
		read -t 5 -p "Enter the index of your number : " x
		if [[ $? -gt 128 ]]; then 
			c=1
			break
		fi
		if [ ${r[$(($x))-1]} -eq $n ];then
			echo "Great"
			((p=p+1))
		else
			c=1
			break
		fi
	done
	elif [ $ch -eq 2 ];then
		echo "HELP: INSTRUCTIONS TO PLAY THE GAME. "
	else
		break
fi

if [ $c -eq 1 ];then
			echo -e "\nGAME OVER\n"
			echo "You scored $p points"
fi
		done

```

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1625753738352/qBDF1PFQG.png)

This is the final bare-bones script without any help instructions just keeping the script simple. I hope you learned something from the game development in BASH. This is just a fun little project and a cool way of learning certain concepts in BASH such as loops, conditional statements, and arithmetic. Have FUN. Happy CODING :)