{
  "title": "Vim: Get the Text from Visual Selection",
  "status": "published",
  "slug": "vim-get-visual-text",
  "date": "2025-04-05T12:33:46Z"
}

<h2>Using Registers</h2>
<p>We can get the selected text in a variable in Vim Script using registers.</p>
<pre><code class="language-vimscript">normal gv&quot;xy
let context = getreg(&quot;x&quot;)
</code></pre>
<p>Lets break down the command</p>
<pre><code>normal mode -&gt; gv -&gt; (y)ank text -&gt; to the &quot;x&quot; register
               |                  |
               |               Copy the contents into x register(or any register you like)    
               |                  
            Select the previously selected text   
</code></pre>
<p>Here, we are entering normal mode and selecting text which was previously selected and yank the contents into a register in this case we are using (x) register, it can be any register.
Now to get the contents of that register we can use the function <code>getreg(&quot;register_name&quot;)</code> or use <code>&quot;xp&quot;</code> to paste the contents of the <code>x</code> register or more generally for any register(<code>&quot;&lt;register-name&gt;p</code>).</p>
<p>Hence we can store the contents of the selected text in a variable for further processing or manipulation.</p>
<p>To quickly test this snippet from command mode, you can try the following steps:</p>
<p>Select a text and press Escape, we just want the <code>gv</code> command to refresh and get it's contents to the latest visual selection.</p>
<pre><code class="language-vimscript">:normal! gv&quot;xy
</code></pre>
<pre><code class="language-vimscript">:let foo = getreg(&quot;x&quot;)
</code></pre>
<pre><code class="language-vimscript">:echo foo
</code></pre>
<p>The echo command will simply print the text which we have selected in the file.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1646483173/blog-media/wlrxgtmegtycilyhvyiz.gif" alt="Visual Select Text"></p>
