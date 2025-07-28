{
  "title": "Vim: Enhancing Movement Speed",
  "status": "published",
  "slug": "vim-movement-speed",
  "date": "2025-04-05T12:33:59Z"
}

<p>![]({{ page.image | relative_url }})</p>
<h2>Introduction</h2>
<p>OK! Vim and movement are like bread and butter or failed brakes. To become a proficient Vim user, you need to move in Vim very effectively. You don't have to think about doing certain things, your fingertips should automatically move without wasting time thinking about it. I am sure, it takes time and effort but OH! it is so rewarding in the end.</p>
<h3>Why H J K L?</h3>
<p>First things first, unmap arrow keys and make a habit of using h,j,k, and l. Yes, this would not make any sense in the initial stage but that will make no sense for not using it later.
The thing with H J K L is that they are quite flexible to work with and if you use it with numbers you would navigate much faster than arrow keys. Such as <code>10j</code> will move you 10 lines down at a time in normal mode. These keys are used with many of the other key shortcuts and commands which just make it most important to begin learning to move around in Vim with H J K L.</p>
<h2>Moving Horizontally.</h2>
<p>This is quite the common movement that every programmer uses most of the time. This is also a much-neglected part when movement in Vim is concerned. To the basics, we use the following commands:</p>
<p><code>w</code>   -&gt;  <strong>move forward by a word (considering punctuations as separate words).</strong></p>
<p><code>W</code>   -&gt;  <strong>move forward by a word (punctuations ignored).</strong></p>
<p><code>b</code>  -&gt;  <strong>move backward by a word (considering punctuations as separate words).</strong></p>
<p><code>B</code>  -&gt;  <strong>move backward by a word (punctuations ignored).</strong></p>
<p><code>e</code>  -&gt;  <strong>move to end of a word (considering punctuations as separate words).</strong></p>
<p><code>E</code>  -&gt;  <strong>move to end of a word (punctuations ignored).</strong></p>
<p><code>0</code>  -&gt;  <strong>move to the beginning of a sentence.</strong></p>
<p><code>$</code>  -   <strong>move to the end of a sentence.</strong></p>
<p>Those are the most useful and common commands for moving across the line. Don't forget to use the number before the command to repeat the task for that number of times. Like for example, if you would like to go 6 words ahead type in <code>6w</code>. This can improve your thinking and typing as well, good signs of a programmer Eh!</p>
<h2>Moving Vertically.</h2>
<p>To move vertically we can imagine moving within a file or the block of code. For moving in a file, the following are some useful commands.</p>
<p><code>gg</code>  -&gt; <strong>move to the beginning of a file.</strong></p>
<p><code>G</code>  -&gt;  <strong>move to the end of a file.</strong></p>
<p><code>Ctrl + e</code>  -&gt;  <strong>move the screen down without moving the cursor.</strong></p>
<p><code>Ctrl + y</code>  -&gt; <strong>move the screen up without moving the cursor.</strong></p>
<p><code>Ctrl + f</code>  -&gt; <strong>move forward one entire screen.</strong></p>
<p><code>Ctrl + b</code>  -&gt; <strong>move backward one entire screen.</strong></p>
<p><code>Ctrl + d</code>  -&gt; <strong>move forward half screen.</strong></p>
<p><code>Ctrl + u</code>  -&gt; <strong>move backward half screen.</strong></p>
<p>This just was moving around the screen and now a bit programmatic movement. We will see some keystrokes to move in code blocks or code snippets very efficiently.</p>
<p><code>gd</code>  -&gt; <strong>move to the local declaration of any code.</strong></p>
<p><code>gD</code>  -&gt; <strong>move to the global declaration of any code.</strong></p>
<p><code>%</code>  -&gt; <strong>move between pairs of ( ), { }, [ ] or any other type of such braces.</strong></p>
<p><code>{</code>  -&gt; <strong>move to the next paragraph/ code block/ function/ etc)</strong></p>
<p><code>}</code>  -&gt; <strong>move to the previous paragraph/ code block/ functions/ etc)</strong></p>
<p><code>fa</code>  -&gt; <strong>move to the next occurrence of the character 'a' in a sentence.</strong></p>
<p><code>Fa</code>  -&gt; <strong>move to the previous occurrence of the character 'a' in a sentence.</strong></p>
<p><code>ta</code>  -&gt; <strong>jump to before of the next occurrence of the character 'a' in a sentence.</strong></p>
<p><code>Ta</code>  -&gt; <strong>jump to after of the previous occurrence of the character 'a' in a sentence.</strong></p>
<p>The above might be quite handy key shortcuts in moving in a large code file. Even in files with complex variable names and structures, this can be quite handy.</p>
<h2>Search and navigation.</h2>
<p>Searching is quite a time-consuming task, especially when the code is quite complex and has a lot of variables and all. Vim shines in many of such aspects where people think it's dead. It rises from the ashes to produce a performance-driven experience like any other modern IDEs though it requires a bit of research:) Here are some commands that will make searching and navigating around it quite a lot easier.</p>
<p><code>*</code> -&gt; <strong>next occurrence of the word under the cursor.</strong></p>
<p><code>#</code> -&gt; <strong>previous occurrence of the word under the cursor.</strong></p>
<p><code>n</code>  -&gt; <strong>next occurrence of the word searched pattern.</strong></p>
<p><code>N</code>  -&gt; <strong>previous occurrence of the word searched pattern.</strong></p>
<p>The above commands will also work if you search the pattern from the command mode.
<code>/pattern</code>  or <code>?pattern</code> Enter and navigate to the next(<code>*</code> or <code>n</code>) and previous(<code>#</code> or <code>N</code>) occurrence of that pattern match.</p>
<h2>Moving across files.</h2>
<p>Moving across files without any plugins or file explorer is often considered tricky or impossible for some people but there is a way. You can switch between files using the following commands:</p>
<p><code>Ctrl + O</code>  -&gt;   <strong>move in the previously opened file.</strong></p>
<p>and</p>
<p><code>Ctrl + I</code>  -&gt;  <strong>move in the next file.</strong></p>
<p>We also can use <code>Ctrl + ^ </code> to move the previous two opened files.</p>
<p>If you want to switch from buffers, you can use <code>:bn</code> for moving into the next buffer, and <code>:bp</code> to move in the previous buffer. You always have an option to move from a buffer using the file name <code>:b filename</code> or using the index as <code>:bindex</code>.</p>
<h2>Moving between Tabs.</h2>
<p>People rarely use Tabs as far as I have seen, but they are quite useful and provide the polish off just as robust IDEs.</p>
<p><code>:tabnew filename</code>  -&gt;  <strong>create a Tab of a file.</strong></p>
<p><code>gt</code>  -&gt; <strong>move to the next tab.</strong></p>
<p><code>ngt</code>  -&gt; <strong>move to the nth tab.</strong></p>
<p><code>gT</code>  -&gt; <strong>move to the previous tab</strong></p>
<p><code>:tabo</code>  -&gt; <strong>close all the tabs except the current one.</strong></p>
<p><code>:tabc</code>  -&gt; <strong>close the tab.</strong></p>
<p><code>:tabm n</code>  -&gt; <strong>move the current tab to nth position.</strong></p>
<h2>Movement in Marks.</h2>
<p>Marks are used for some quite large files and code-bases. It is used to move from one mark(kind of a bookmark) to another using few key commands, marks are generally created when you would go to a particular code block or any part of the file again and again. Some of the quick navigation using maps are the following.</p>
<p><code>mn</code>  -&gt;  <strong>set the current position as mark 'n'.</strong></p>
<p><code>&lt;backtick&gt;n</code>  -&gt;  <strong>jump to the position of mark 'n'.</strong></p>
<p><code> </code>0`  -&gt; <strong>jump to the position where vim was last exited.</strong></p>
<p><code> </code>&quot;`  -&gt;  <strong>jump to the position when the last edit was made in the file.</strong></p>
<h2>Split Windows Movement</h2>
<p><code>Ctrl + w + r</code>  -&gt; <strong>move the split down.</strong></p>
<p><code>Ctrl + w + R</code>  -&gt; <strong>move the split up.</strong></p>
<p><code>Ctrl + w + h</code>  -&gt;  <strong>jump to the left split.</strong></p>
<p><code>Ctrl + w + j</code>  -&gt;  <strong>jump to the split down.</strong></p>
<p><code>Ctrl + w + k</code>  -&gt; <strong>jump to the upper split.</strong></p>
<p><code>Ctrl + w + l</code>  -&gt;  <strong>jump to the left split.</strong></p>
<p>You can use Caps H J K L to move the leftmost, bottom, uppermost, rightmost split respectively.
That just was quick to make you enough faster than previous hassles.</p>
<p>That was probably it, these were some tricks and shortcuts to move around Vim pretty effectively and smoothly. Moving around Vim can be quite complicated at once, but it's just finding the key shortcuts to make you feel and glide in VIm. There might be quite a lot of shortcuts missing, If you have any quicker shortcuts, Please let me know in the comments.  Happy Viming :)</p>
