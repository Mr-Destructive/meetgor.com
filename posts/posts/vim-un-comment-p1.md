{
  "title": "Comment/Uncomment Code: Vim for Programmers",
  "status": "published",
  "slug": "vim-un-comment-p1",
  "date": "2025-04-05T12:33:52Z"
}

<h2>Introduction</h2>
<p>We as programmers always fiddle with commenting out code for code testing, documenting the function of code, and most importantly debugging. So you can't wait to comment on a large chunk of code manually, as it is quite a tedious thing to do. Let's do it effectively in Vim.</p>
<p>In this part of the series, I'll cover how to comment/uncomment chunks/blocks of code effectively in Vim. We will see and use some commands, keybindings for doing so, and also we would add certain components to our vimrc file as well to design some custom key mappings.  Let's get faster with Vim.</p>
<h2>How to comment multiple lines effectively</h2>
<p>To comment on multiple lines of code, we can use the Visual Block mode to select the lines, and then after entering into insert mode, we can comment a single line and it would be reflected on all the selected lines.</p>
<ol>
<li>
<p>Press <code>CTRL+V</code> and Select the line using j and k</p>
</li>
<li>
<p>After Selecting the lines, Press <code>Escape</code></p>
</li>
<li>
<p>Press <code>Shift + I</code>, to enter insert mode</p>
</li>
<li>
<p>Enter the comment code (<code>//</code>, <code>#</code>, or other)</p>
</li>
</ol>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1633518136135/06dfBTq2T.gif" alt="vimcoment.gif"></p>
<p>So, using just simple steps you can comment out large chunks of code quite easily and effectively. If you are using some other language that has multiple characters for commenting like <code>//</code>, <code>- -</code>, etc, you can type in any number of characters while being in insert mode after selecting the lines.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1633520509953/0q-k2ZHC7.gif" alt="vimcppcom.gif"></p>
<p>This might look a bit wired on the first try but just try it every day, It is a life-saving and very satisfying experience once applied in a real-world scenario.</p>
<h2>How to uncomment multiple lines effectively</h2>
<p>Now, as we have seen to comment out a large chunk of code, we can even uncomment the code very easily. It's even simpler than commenting the code.</p>
<ol>
<li>
<p>Press <code>CTRL + V</code> to enter Visual Block mode</p>
</li>
<li>
<p>Select the commented characters</p>
</li>
<li>
<p>Press <code>d</code> to delete the comments</p>
</li>
<li>
<p>Press <code>Escape</code></p>
</li>
</ol>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1633518156818/GJzRPTI3I.gif" alt="vimuncoment.gif"></p>
<p>We can simply use the CTRL + V to select the comment, and then press d to delete all the comment characters.</p>
<p><strong>We are using the Visual Block mode as we only want the comment to be selected and not the entire code associated with the lines.</strong></p>
<h2>Using Multiline Comments for Programming languages</h2>
<p>Now you might say, why use multiple single-line comments when we can use multiline comments in almost all programming languages. Well, Of course, you can do that, it's easier for reading the code if syntax highlighting is accurate and greys out the commented part. We can simply add those characters to the start of the block and at the end of the block.</p>
<p>But in Vim, we can customize that too, just imagine when you just select the chunk/block of code that you need to comment out and then simply press a few keystrokes (just 2) and the multiline comments are automatically (programmatically) added as per the programming language extension of the file.</p>
<p>Isn't that cool? Well, you just need to copy-paste the below code to your Vimrc file and source it and you are good to go.</p>
<pre><code class="language-vim">function! Comment()
    let ext = tolower(expand('%:e'))
    if ext == 'py' 
        let cmt1 = &quot;'''&quot;
	    let cmt2 = &quot;'''&quot;   
    elseif ext == 'cpp' || ext =='java' || ext == 'css' || ext == 'js' || ext == 'c' || ext =='cs' || ext == 'rs' || ext == 'go'
	    let cmt1 = '/*'
	    let cmt2 = '*/'
    elseif ext == 'sh'
	    let cmt1 = &quot;:'&quot;
	    let cmt2 = &quot;'&quot;
    elseif ext == 'html'
	    let cmt1 = &quot;&lt;!--&quot;
	    let cmt2 = &quot;--&gt;&quot;
    elseif ext == 'hs'
	    let cmt1 = &quot;{-&quot;
	    let cmt2 = &quot;-}&quot;
    elseif ext == &quot;rb&quot;
	    let cmt1 = &quot;=begin&quot;
	    let cmt2 = &quot;=end&quot;
    endif
    exe line(&quot;'&lt;&quot;).&quot;normal O&quot;. cmt1 | exe line(&quot;'&gt;&quot;).&quot;normal o&quot;. cmt2 
endfunction

function! UnComment()
    exe line(&quot;'&lt;&quot;).&quot;normal dd&quot; | exe line(&quot;'&gt;&quot;).&quot;normal dd&quot;   
endfunction


vnoremap ,m :&lt;c-w&gt;&lt;c-w&gt;&lt;c-w&gt;&lt;c-w&gt;&lt;c-w&gt;call Comment()&lt;CR&gt;
vnoremap m, :&lt;c-w&gt;&lt;c-w&gt;&lt;c-w&gt;&lt;c-w&gt;&lt;c-w&gt;call UnComment()&lt;CR&gt;

</code></pre>
<p>The below screencast is an example of <code>HTML</code> snippet in a file that is getting commented using mapping with the keys <code>,m</code> you can put any other keybinding you like.
<img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1633595891674/hbhrbtRHd.gif" alt="htmcm.gif"></p>
<hr>
<p>Similarly for the next screencast is of an <code>Javascript</code> snippet in a file which is getting commented using a mapping with the keys <code>,m</code> and uncommented using <code>m,</code> again you can put any other keybinding you like.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1633595919104/xGTh5ztWu.gif" alt="jscom.gif"></p>
<hr>
<p>The following screencast is of a shell script(BASH) snippet.
<img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1633596156121/tbGHQBSSA.gif" alt="shcom.gif"></p>
<hr>
<h3>Multiline Comments in various Programming Languages:</h3>
<h4>1. C / C++ / Java / Javascript / CSS / C# / Rust / Go / PHP / Swift / Dart / Kotlin</h4>
<pre><code>/*
*/
</code></pre>
<h4>2. Python</h4>
<pre><code>'''
'''
</code></pre>
<p>You can even use <code>&quot;&quot;&quot;</code> double quotes instead of single quotes</p>
<h4>3. BASH (Shell Scripting)</h4>
<pre><code>: '
'
</code></pre>
<p>You can even use <code>: &quot;</code> and <code>&quot;</code> double quotes instead of single quotes</p>
<h4>4. Haskell</h4>
<pre><code>{-
-}
</code></pre>
<h4>5. Ruby</h4>
<pre><code>=begin
=end
</code></pre>
<h4>6. HTML</h4>
<pre><code>&lt;!--
--&gt;
</code></pre>
<h4>7. Julia</h4>
<pre><code>#=
=#
</code></pre>
<h3>Understanding the Commands / Keymapping</h3>
<p><strong>NOTE : You need to go from the top to bottom while commenting on the block of code, otherwise, it would be a mismatch in commenting for specific language syntax. While uncommenting the order doesn't matter.</strong></p>
<h4>Getting the extension (filetype)</h4>
<p>In Vim, we can get the file extension i.e. we can get the programming language associated with the current file. To do that we can use, <code>expand('%:e')</code>.</p>
<p>This will give us the file extension of the current file. Just for simplicity, <code>%</code> means the current file, added to it is <code>:e</code> for excluding the filename and keeping the extension. We convert the extension into lowercase just for keeping things safe and programmatic and store it in a variable <code>ext</code>.</p>
<h4>Checking for programming language</h4>
<p>We then can then use an if-else ladder to check for the programming languages and assign two variables <code>cmt1</code> for the initial characters in the multiline comment and <code>cmt2</code> for enclosing the comment.</p>
<h4>Typing in the characters</h4>
<p>We can use the function <code>line(&quot;'&lt;&quot;)</code> to get the line number of the previous visual selection. Similarly, <code>line(&quot;'&gt;&quot;)</code> for the ending line. We are using the <code>exe</code> command to execute the function <code>line</code> and so we have to use a concatenation of the commands even to write the raw commands like <code>i</code> to insert mode, <code>o</code> to insert mode but a line below the cursor. So, we use <code>normal</code> command for that. This command indicates the interpreter to execute the following commands from the normal mode.</p>
<p>We have to enclose the <code>normal</code> command in double-quotes/single quotes. We can simply use the variable again with concatenation.</p>
<pre><code class="language-vim">exe line(&quot;'&gt;&quot;).&quot;normal o&quot;. cmt2 
</code></pre>
<p>The above command will fetch the last line's number of the previous visual selection followed by entering <code>o</code> from the normal mode and concatenated with the value of the variable <code>cmt2</code> which we have already initialized in the <code>Comment</code> function. We are using <code>|</code> for running multiple commands as we also need to include the comment at the beginning of the visual selection.</p>
<p>For uncommenting the code, we are simply deleting the entire first and the last line in the visual selection. For that, we have used <code>dd</code> from the normal mode.</p>
<h3>Conclusion</h3>
<p>So, from the following type of tutorial, we were able to set up our Vim editor for efficient code commenting/ uncommenting using some commands, key shortcuts, and configuring the vimrc for making custom keymappings. We were also able to understand the multiline comments in various programming languages and use them in Vim very effectively with a simple addon to the config vimrc file.</p>
<p>Thank you for reading, hope you found this article helpful. If you have any queries or wanna add multiline comments for some more programming languages please let me know in the comments or contact section.</p>
<p>Happy Coding :)</p>
<h3>References</h3>
<ul>
<li>
<p><a href="https://stackoverflow.com/questions/1676632/whats-a-quick-way-to-comment-uncomment-lines-in-vim/1676690">StackOverflow - Commenting lines in Vim </a></p>
</li>
<li>
<p><a href="https://dev.to/grepliz/3-ways-to-comment-out-blocks-of-code-in-vi-6j4">Liz Lam - 3 ways to comment code in Vim</a></p>
</li>
<li>
<p><a href="https://vi.stackexchange.com/questions/9644/how-to-use-a-variable-in-the-expression-of-a-normal-command">StackExchange - Use variable in normal command</a></p>
</li>
</ul>
