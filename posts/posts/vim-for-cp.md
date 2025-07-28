{
  "title": "Vim for Competitive Programming",
  "status": "published",
  "slug": "vim-for-cp",
  "date": "2025-04-05T12:33:54Z"
}

<h2>Introduction</h2>
<p>Vim is not a bad text editor when it comes to using it for Competitive Programming. It's kind of one way or the other, you would love it or you could trash it as it can waste a tremendous amount of time to write code every time. But once you are in a decent setup including some key-bindings and plugins can improve your speed in using Vim. I personally have used it since the beginning and haven't regretted it even a tiny bit till date. It's a flawless and enjoyable experience. I've used C++ for my CP journey, C is quite similar as well, but C++ has more library support and is a bit easier to write comparatively.</p>
<p>Using Vim for CP is a bit challenging in the initial stage but just stay consistent and you'll be in flying colors on the other side. Here, I'll give some quick key-bindings and plugins to include to Vim to enhance the workflow for Competitive Programming.</p>
<h2>Quick Boilerplate code</h2>
<p>Firstly, let's create a key-binding for the boilerplate code. We just need to map the key with the code we need to write which will be integrated with Vim commands.</p>
<p>For C++</p>
<pre><code class="language-vim">nnoremap cpf i#include&lt;iostream&gt;&lt;Esc&gt;ousing namespace std;&lt;Esc&gt;o&lt;CR&gt;int main(){&lt;Esc&gt;o&lt;Esc&gt;oreturn 0;&lt;Esc&gt;o}&lt;Esc&gt;kki
</code></pre>
<p>For C:</p>
<pre><code class="language-vim">nnoremap cp i#include&lt;stdio.h&gt;&lt;Esc&gt;o&lt;CR&gt;int main(void){&lt;Esc&gt;o&lt;Esc&gt;oreturn 0;&lt;Esc&gt;o}&lt;Esc&gt;kki
</code></pre>
<p>For Java:</p>
<pre><code class="language-vim">nnoremap &lt;C-j&gt; iclass &lt;ESC&gt;&quot;%pxxxxxa {&lt;ESC&gt;opublic static void main(String args[]){&lt;Esc&gt;o&lt;Esc&gt;o}&lt;Esc&gt;o}&lt;Esc&gt;kki&lt;Tab&gt;&lt;Tab&gt;
</code></pre>
<p>Now, you get the idea of creating such macros for your own programming language. I don't think Python can have boilerplate code, even it has it's quite small compared to C/C++/Java/etc. You can include those codes into the key map. Let me explain the map to you.</p>
<p>Let's look at the Java example, the <code>nnoremap</code> is the map command with attributes like non-recursive and the mode of mapping. In this case <code>n</code> stands for normal mode map i.e. the map will get triggered in the normal mode, <code>nore</code> stands for non-recursive behavior of the map. For further readings, you can read my <a href="https://mr-destructive.github.io/techstructive-blog/vim/2021/06/14/Vim-Keymapping.html">Vim-keymapping Guide</a> article. The map itself has the right and a left part, kind of a map of the key and command to execute. You can use any key to trigger it, I have used CTRL+J or <code>&lt;C-j&gt;</code> to activate the map. The command is quite long but let's break it down into smaller chunks.</p>
<ul>
<li>Starting with <code>i</code>, we enter into insert mode and we type in class and hit <code>&lt;ESC&gt;</code> key and this is a special trick to deal with java for the class name as it should be the name of the file itself. We use the <code>%</code> register for copying the current file's entire name (with extension), so we have to remove the <code>.java</code> part so, you see five x's to delete those after we have pressed a to get into insert mode just ahead of the cursor. We then insert <code>{</code> for the class and again hit escape to go back to normal mode.</li>
<li>We type <code>o</code> to go into insert mode to the next line, type in <code>public static void main(String args[])</code> and this is the same stuff of escaping and going in next line.</li>
<li>Finally, we hit <code>kk</code> to move up twice and enter the insert mode with <code>i</code> and hit tab twice to get the indentation correct.</li>
</ul>
<p>So, you can add more as per your preference of pre-defined constants and functions and the rest stuff involved in your language of choice.</p>
<h2>Running Code by a Key-binding</h2>
<p>After the boilerplate code being taken care of, we can now move to the building and running of the code file in Vim. We can add some more custom key bindings to auto compile and run the code within Vim.</p>
<p>We can compile the code with the compilers or interpreters for the languages from the Vim using <code>!</code> which executes the commands from the terminal and pass in <code>%</code> i.e. the current file name and output file as <code>%:r</code> i.e. current file name without the extension in case of C/C++/Java.</p>
<p><strong>Compile -&gt; <code>c++ hello.cpp -o hello</code></strong></p>
<p><strong>Run -&gt; <code>%:r</code> for Windows</strong></p>
<p>and</p>
<p><strong><code>./%:r</code> for Linux/macOS</strong></p>
<p>We have added those two commands i.e (compile and run) using <code>&amp;&amp;</code> to make things more clear and compact.
For C++</p>
<pre><code class="language-vim">nnoremap cpp :!c++ % -o %:r &amp;&amp; %:r&lt;CR&gt;
</code></pre>
<p>For C</p>
<pre><code class="language-vim">nnoremap c, :!gcc % -o %:r &amp;&amp; %:r&lt;CR&gt;
</code></pre>
<p>For Java</p>
<pre><code class="language-vim">nnoremap ,j :!javac % &amp;&amp; java %:r&lt;CR&gt;
</code></pre>
<p>For Python</p>
<pre><code class="language-vim">nnoremap py :python %&lt;CR&gt;
</code></pre>
<p>Again, you can apply this to any programming language you use. Also kindly note that their could be some differences for compiling the code in Linux and Windows so just be careful which command you use in which Operating system.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631540728028/2puTZUXwK.gif" alt="vimjv.gif"></p>
<p>The above gif demonstrates the usage of those commands and key bindings to make the workflow of running and building process much easier and smoother.</p>
<h2>Opening Terminal to compile and run code</h2>
<p>We can use the terminal to compile and run code if you wish to do it this way as well. Vim is quite a terminal-friendly editor, you can open any terminal from Vim.</p>
<p>You can enter <code>:term</code> to open the default terminal in a horizontal split. You can explicitly enter the terminal/shell name to open it like <code>:term bash</code> to open bash, <code>:term powershell</code> to open PowerShell, <code>:term zsh</code> to open zsh, and so on. This is quite a great utility that vim provides with it. You can switch between windows using <code>&lt;C-w&gt;</code> or CTRL+W <strong>twice</strong> or use HJKL keys along with <code>&lt;C-w&gt;</code> to move with ease. More on Vim and Integrated Terminals in <a href="https://mr-destructive.github.io/techstructive-blog/vim/2021/06/29/Vim-Terminal.html">this article</a>.</p>
<p>You can now explore the terminal world yourself and execute the commands from therein.</p>
<h2>Plugins</h2>
<p>When it comes to Vim, there are a ton of plugins made for everything you can imagine and couldn't imagine. We can use Plugin managers to manage those plugins for us or manually manage those yourself. You can go with one of the three well-known plugin managers like <code>Vim-Plug</code>, <code>Vundle</code> or <code>Pathogen</code>.</p>
<p>Using those plugin managers, we can install plugins like some of the useful ones are:</p>
<h3>NERDTree</h3>
<p>NERDTree is a great plugin for beginners and also for people doing competitive programming as it saves a bit of time as well. You can install the NERDTree plugin from the GitHub docs. You can read about the detail of configuring the NERDTree plugin in <a href="https://mr-destructive.github.io/techstructive-blog/vim/2021/08/31/Vim-NERDTree.html">this article</a>. This plugin allows you to have a graphical representation of the files and folder structures in the current directory. You can extend its capabilities to auto-refresh and auto open/close using some custom configuration as provided in the above-mentioned article.</p>
<p>You basically open the NERDTree with the command <code>:NERDTree</code> after installing and configuring. You can make key-bindings and shortcuts to use as you like and prefer.</p>
<h3>Autocompletion and LSP</h3>
<p>Autocompletion is quite important when it comes to competitive programming as you have to save time on the <strong>stuff you know</strong> and let the application process it for you. We do have some plugins for auto-completion and also LSPs for this specific task. LSP is quite great and provides much more accurate predictions when it comes to autocompletion. You can check out the LSP <a href="https://github.com/prabirshrestha/vim-lsp">unofficial plugin for vim</a> for more information.</p>
<p>For normal plugins like <a href="https://github.com/vim-scripts/AutoComplPop">AuotoComplPop</a> are also decent and provide a great beginner experience.</p>
<p>Not many plugins would be required for competitive programming just some autocompletion and a nice interface with a file system will work for the basic set-up but that is not the end, you can extend it as per your needs and skills. Vim is quite epic in terms of the customization it can offer.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631768130787/ccDyyJ45y.gif" alt="vimcpt.gif"></p>
<p>Similarly, for C++, I have demonstrated the use of key mappings and also showed the usage of the terminal inside of Vim, this is quite a flawless experience to integrate and use the terminal as per choice and preference.</p>
<h2>Conclusion</h2>
<p>So from this article, we were able to set up Vim for competitive programming by adding some custom key mappings and plugins for different languages. This was not a complete guide, I could not include every language's configuration here, surely you can ask me anytime in the comments, GitHub, Twitter or just mail me I can try to configure a programming language on Vim. Thanks for reading.</p>
<p>Happy Viming and Happy Coding :)</p>
