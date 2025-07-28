{
  "title": "Setting up Vim for BASH Scripting",
  "status": "published",
  "slug": "vim-for-bash",
  "date": "2025-04-05T12:34:00Z"
}

<h2>Vim and BASH?</h2>
<p>Bash Scripting is a powerful skill to have as a programmer because we find Linux almost everywhere and to get through it you must have a command over its interface which is generally the BASH shell. Vim is a great option for doing this, or probably the best out there! Because Vim is pre-installed in almost every Linux distribution. This is not an in-depth setup for BASH on Vim, it is a simple editorial on starting up BASH scripting on the Vim editor. So without wasting time on &quot;Vim features&quot; let's dive in with the setup for BASH in Vim.</p>
<h2>Boilerplate macro</h2>
<p>Setting up a bash script doesn't require much code but still in some cases it can be a bit hassle and to avoid the repetitive task, one can easily set up a macro for the boilerplate BASH script.</p>
<pre><code class="language-vim">nnoremap bs i#!/bin/bash/&lt;ESC&gt;o
</code></pre>
<p>Ok that was pretty dumb but it can scale pretty quickly and it will be nice to tailor it as per needs, here's some snippet with function pre-loaded.</p>
<pre><code class="language-vim">nnoremap bs i#!/bin/bash/&lt;ESC&gt;o
nnoremap bs i#!/bin/bash/&lt;ESC&gt;o&lt;ESC&gt;ofunction main(){&lt;ESC&gt;o&lt;ESC&gt;o}&lt;ESC&gt;ki&lt;S-TAB&gt;

</code></pre>
<p><img src="https://s6.gifyu.com/images/bsclip.gif" alt="type bs to load boiler-plate code"></p>
<p>When the key bs is typed in normal mode, you enter into insert mode as per the command macro, then we type in the required text and escape to move to the next line and continue the same stuff. This could be extended further like making some input or printing out some text and any other formatted text that you could think it as repetition.</p>
<h2>Sourcing Scripts</h2>
<p>So, after creating the file, sourcing the script, and running it can be a bit slow for some people, as you have to go to the terminal and toggle in the permission to run the script and then run, But pull on your seatbelts as this is VIM! You can die due to slowness!</p>
<pre><code class="language-vim">nnoremap sh :!chmod +x % &amp;&amp; source %
</code></pre>
<p><img src="https://s6.gifyu.com/images/shclip.gif" alt="type sh to run script"></p>
<p>When the sh keys are typed in the normal mode, the preceding command after ! (bang) will be executed in the terminal, the &amp;&amp; keywords will execute the second command only when the first command is successfully executed.
That just can so fast! Imagine doing this for long scripts and especially for debugging, it will waste 2 minutes every time you leave the editor and for 10 times you do the debugging, you will carelessly was roughly 20 minutes! Improve your debugging skills surely :)</p>
<h2>Plugins</h2>
<p>There are very few plugins out there for BASH as for VIM, but it's quite to write scripts even without any plugins. One of the most supported and popular plugins for BASH in Vim is  <a href="https://www.vim.org/scripts/script.php?script_id=365">Bash-Support-Vim</a> for auto-completion and <a href="https://www.shellcheck.net">Shell-Check</a> for finding/correcting any bugs or error in the script .
The mentioned plugin is quite awesome and it can greatly improve the speed of scripting for BASH, some commands such as shortcuts for writing if-else, while, for loops, commenting and other aspects in the scripting. The thorough documentation for such commands is also provided by the  <a href="https://wolfgangmehner.github.io/vim-plugins/bashsupport.html">plugin website</a>.
This can be used for autocompleting keywords and writing nested if-else and other logical operators in BASH scripting. Again, you can do absolutely fine without plugins in Vim as it is heavily customizable to the user's need and can be very rewarding to set up your own configuration for BASH. You can use standard Vim(barebones) for auto-completion as well with the command CTRL+N and CTRL-P to move down and up respectively.</p>
<h2>Some More Tricks</h2>
<p>BASH in Vim can be quite versatile to use as it provides some custom addons to make the script more functional and easier to understand. Some tricks such as using autocompletion can be quite inconvenient to use at once but it can get really smooth after some runs at writing the scripts.</p>
<ul>
<li>In BASH Scripts there are quite a lot of brackets to play with that's why to jump around swiftly around such parentheses or brackets you can use <strong>% to move from the opened to closed brackets or vice versa</strong>.</li>
<li>You can execute any terminal command from Vim, be sure to be in command mode and press ! after the command, you would like to execute. This will run the command from the terminal and you don't have to leave the editor, it saves a ton of time and it's blazingly fast.</li>
<li>With the above trick, you kind of have a superpower within Vim to make, build, source, run the files or scripts within Vim, that is not repetition but it can run bash within bash. Ok! that's was pretty fast. Don't die of quickness now!</li>
</ul>
<p>Writing BASH scripts in Vim can be also boosted by using some built-in commands such as adding comments for multiple lines at once and some unexplored stuff which can be learned in the way to understanding the flow of Vim and BASH together. Happy Coding and Viming :)</p>
