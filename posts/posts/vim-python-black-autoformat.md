{
  "title": "Autoformat Python file with Black after saving in Vim",
  "status": "published",
  "slug": "vim-python-black-autoformat",
  "date": "2025-04-05T12:33:45Z"
}

<p>If you are like me who writes Python very badly, it has empty lines with whitespaces, no proper format in assigning variables, not formatted according to <a href="https://peps.python.org/pep-0008/">PEP 8</a> standards, and you use Vim as your text editor then my friend you need a autocmd badly for it.</p>
<h2>Install Black in Python</h2>
<p>Install the <a href="https://pypi.org/project/black/">black</a> package in python globally or locally as per your preferences.</p>
<pre><code>pip install black
</code></pre>
<p>OR with pipx</p>
<pre><code>pipx install black
</code></pre>
<p>For a detailed guide about running packages with pipx head toward my article on <a href="https://mr-destructive.github.io/techstructive-blog/pipx-intro/">pipx</a>.</p>
<h2>Set up Autocmd in Vim</h2>
<p>We can set up a autocmd. What is a autocmd? It is about running commands when certain events occur like running a command when a file is saved, a buffer is opened or closed, and so on. What we want is to run the black command from the shell when the current file is saved.</p>
<p>So, we can create a autocmd as follows:</p>
<pre><code class="language-vimscript">autocmd BufWritePost * !black %
</code></pre>
<p>Now, this will run when any type of file is saved, so we will make it specific to python by adding a <code>*.py</code> to add in the autocmd.</p>
<pre><code class="language-vimscript">autocmd BufWritePost *.py !black %
</code></pre>
<p>This works, but it gives a prompt after the command has been executed, to run the command silently we can simply add the silent keyword before the execution of the command from the shell.</p>
<pre><code class="language-vimscript">autocmd BufWritePost *.py silent !black %
</code></pre>
<p>This looks perfect!</p>
<p>But still, we need to add a auto-group(<code>augroup</code>) that groups the autocmds and by adding <code>autocmd!</code> it will clear all the commands from the group.</p>
<pre><code class="language-vimscript">augroup python_format
    autocmd!
    autocmd BufWritePost *.py silent !black %
augroup end
</code></pre>
<p>We can now add it to the vimrc to work all the time.</p>
<h2>Using pipx</h2>
<p>If you have used pipx to install black, you need to setup the autocmd a bit differently.</p>
<pre><code class="language-vimscript">autocmd BufWritePost *.py silent !pipx run black %
</code></pre>
<p>It might be a bit slower than running with global installation, but it is a neat way to run python package.</p>
<p>So, that's it we can now write clean and safe python code without breaking a sweat in Vim. Happy Coding :)</p>
