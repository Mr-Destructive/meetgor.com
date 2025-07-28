{
  "title": "Why and How to make and use Vim as a text editor and customizable IDE",
  "status": "published",
  "slug": "vim-text-editor-ide",
  "date": "2025-04-05T12:34:01Z"
}

<p>We all are asked to use visual studio code and other rich looking editors as a beginners text editor or an IDE but that makes a habit of a rich and pleasing interface of running and debugging programs, and that is the reason we don't feel good to use command line or the terminal. Terminal or Command line are quite important to learn and are powerful as well.</p>
<p>Vim is a text editor, not any text editor but a special and one of the fastest out there if you use it wisely. Vim can be used in the terminal which means it can run terminal commands inside the editor interface. That makes it quite a great option to improve your terminal command skills and also for enhancing your Linux knowledge. You can run program files from vim itself just with a few keystrokes. This text editor is preinstalled in most Linux distributions and hence you could quickly edit some files without going anywhere.</p>
<p>So that being said, vim is not easy to learn, you have to put some effort to reap the fruits of saving time in the longer run. It takes time to get used to the commands and key combination of Vim.</p>
<h2>Download VIM</h2>
<p>But, let us get our feet wet in learning a few commands in Vim.</p>
<p>To download vim you can go to the official  <a href="https://www.vim.org/download.php">Vim website</a> .</p>
<p>Firstly let us understand how to open vim from the command line or the terminal</p>
<p>If you are on windows you have to add the vim.exe file's path to the environment variable PATH. After that open CMD or terminal for windows 10 and type vim and enter. You are in Vim.</p>
<p>If you are in Linux or macOS, then open the terminal and type vim and enter.</p>
<p>After you are in Vim, you will surely be trying to type something right? else you are not a geek ; )</p>
<h2>VIM Basics</h2>
<p>You won't see anything happening even if you are typing. This is because are 4 basic modes in Vim, understand it like a room.  The modes are</p>
<ol>
<li>
<p><strong>Normal Mode</strong></p>
</li>
<li>
<p><strong>Insert Mode</strong></p>
</li>
<li>
<p><strong>Command Mode</strong></p>
</li>
<li>
<p><strong>Visual Mode</strong></p>
</li>
</ol>
<p>Wait for a second, you should learn how to move in vim first but you cannot learn everything at once, So take it bit by bit.</p>
<p>To navigate around vim like blazing fast you can use h,j,k, and l.  I know it's wired but you will get used to it.</p>
<p><strong>h</strong>  -&gt; move left.    ( hop left)</p>
<p><strong>l</strong>   -&gt; move right.   (opposite of left)</p>
<p><strong>k</strong>  -&gt; move up.       ( keep it up)</p>
<p><strong>j</strong>   -&gt; move down.   ( jump down!!)</p>
<p>If you want to apply certain movements several times you can use numbers before hjkl to move the number of times in that direction.</p>
<p>Let's say you type 23j , here you will go down 23 lines. You can also use : and number to jump to a particular line of that number, :12 moves you to the 12th line.</p>
<p>The default mode is the <strong>Normal mode</strong> where you can navigate in your file, like a passageway between different rooms. Each other mode is accessible via Normal mode. You have to enter normal mode before switching to the other three modes. By default, after opening vim you are in Normal mode. You start inserting text in the file while being in Normal mode but you can replace text and edit the existing text in the file. <strong>Press ESC to enter Normal mode</strong>.</p>
<p>The main mode in VIM is <strong>Insert mode</strong>, where the actual typing and text editing takes place, <strong>Enter i to enter insert mode</strong>. After entering i you can type anything you want. After you feel satisfied with writing in VIM, you can escape out of the insert mode (room) and enter the normal mode (passageway). Now the biggest question VIM beginners ask <strong>&quot;HOW ON EARTH SHOULD I EXIT OUT OF VIM ??&quot;</strong> . Take a breath dear, you need to save your file first. To do that let's explore a different mode.</p>
<p>The next crucial mode is <strong>Command mode</strong>, where you have to remember some commands to do something. Let's continue from where we left. To save our file,  <strong>Press :  and Voila!! you are in Command mode</strong>. After it depends you want to save the file, quit vim, save the file and quit vim or quit vim without saving the file. OK, that is too much to take in one go. Be sure to be in Command mode to press the following keys (i.e. press colon before any keys if you were in normal mode).</p>
<p><strong>w</strong>     -&gt; save file.</p>
<p><strong>q</strong>      -&gt; quit vim.</p>
<p><strong>wq</strong>     -&gt; save file and quit vim.</p>
<p><strong>q!</strong>      -&gt; quit vim without saving the file.</p>
<p>These are a few of the commands available in VIM :)</p>
<p>So now let's talk about <strong>Visual mode</strong>, the little helper mode (room) to do cut/ copy in VIM. There are many ways to edit or add text in Visual mode. You have options to choose from character by character, line by line, or block by block. <strong>To enter Visual mode press v</strong>. This gets you into character by character highlighting whether you want to cut the text or copy (yank) the text. After selecting the portion you can either cut that portion or copy it in vim register.</p>
<p><strong>y</strong>     -&gt; Copy the selected portion.</p>
<p><strong>d</strong>     -&gt; Cut the selected portion.</p>
<p>You can use <strong>V (Shift+V) to enter Line Visual mode</strong>, this selects or highlights the portion line by line and not single letter by letter.</p>
<p>You can also use <strong>Ctrl+V to enter Block Visual mode</strong>, this selects a block of code.</p>
<p>Here is a good way to remember switching between basic modes in VIM -&gt;</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1613912016216/SXavZAkPH.png" alt="image.png"></p>
<p>Switching between basic modes in VIM</p>
<p>It seems to be quite slower than VS Code, doesn't it? Ya, I know it's kinda slow but with time it picks up pace, just like a train. Spend time in VIM and it will be worth it. No Text editor is so powerful as VIM, it is so for a reason. If you are curious to dive deeper just dive in fully or stay in VS Code and your fancy.  There is a great community around VIM and you'll be happy surely to be a part of it. Oh! The quick fact even VS Code has an extension for VIM ;)</p>
<h2>Some more Basics</h2>
<p>You can do most of the basic programming using this but to be a bit faster you can use some tweaked commands to do stuff like,</p>
<p><strong>I</strong>     -&gt; Enter insert mode to the character at the beginning. ( i puts you in insert mode to the cursor)</p>
<p><strong>a</strong>  -&gt; Enter insert mode adjacent right to the cursor.</p>
<p><strong>A</strong>   -&gt; Enter insert mode at the end of the line.</p>
<p><strong>o</strong>    -&gt; Enter insert mode a line below cursor.</p>
<p>O   -&gt; Enter insert mode a line above cursor.</p>
<p>These commands are an optimized way to enter insert mode in a suitable way. These commands can be used as per needs and quite significantly improve editing speed.</p>
<p>You can also shift from one mode to Normal mode by pressing Ctrl+C or Ctrl+{. Whatever you feel good with.</p>
<p>Pasting is a programmer's everyday task. No, it's good as a beginner and only if you are not running as it is. In VIM you can paste from your previously copied text using Visual mode or deleted chunk just using p to paste. Remember if you are on a character it will paste on starting from where your cursor is, which means your character on the current cursor will be overwritten. You can also use &quot;0p to paste from the copied buffer.</p>
<p>If you are pasting from an external source i.e. out of VIM, you have to use &quot;+p . It pastes from the next character from the cursor.</p>
<p>If you want to search a word in a file, you can use navigation commands like hjkl but there's a faster way around, If you press / and type the word you are searching for, you will get the first instance of that word. If there are multiple instances of the same word, you can press n to got the next instance. You can also press Shift+N to move to the previous instance of the word.</p>
<p>Some more word searching ahead. Say if you are already on the word which you want to find its instance, then you could press * to move to its next instance and # will do the same backward. Oh ! that was smooth.</p>
<p>You can see that we are just using our keyboards while using vim. This is a great feature of old but gold text editors like vim and others that make you faster on the keyboard and avoid mouse traps.</p>
<p>If you want to jump to the end of the block or a paragraph you can use { and } to go to the beginning and the end of it respectively.</p>
<p><strong>G</strong>    -&gt; to the end of the file.</p>
<p><strong>gg</strong> -&gt; to the beginning of the file.</p>
<p><strong>L</strong>     -&gt; move the cursor to the end of the screen.</p>
<pre><code>:%s/old/new/g
</code></pre>
<p>in Command mode -&gt; replace the old word with a new word throughout the file(g)</p>
<p>Ctrl p   in Insert mode    -&gt; auto-complete reference from current file.</p>
<p>So these are some quick and pretty good commands to get started with and feel VIM. But this is just VIM used 25% efficiency, you would have to go to the command line or terminal to run and compile the program you have written. It could waste a lot of time. But as a beginner, this is pretty decent and helps to learn the core programming, behind the scenes of compiling and learning new stuff and commands. It gets pretty exciting if you have got the basics cleared. But be sure to make it through.</p>
<p>It's time to use VIM to its actual powers. VIM + terminal. It can get quite fast. So beware!!!</p>
<h1>Running Programs from VIM</h1>
<p>It turns out that you could go to the command mode and run terminal commands from there just use! before the command first.</p>
<p>Linux/Mac Users: use terminal from vim</p>
<pre><code>:!ls   
</code></pre>
<p>//-&gt; ls command from vim</p>
<p>Windows: use CMD from vim</p>
<pre><code>:!dir
</code></pre>
<p>-&gt; dir command from vim</p>
<p>You could also use other commands which are appropriate to be executed from vim.</p>
<p>Here is where the trick comes in, you don't have to quit vim now to compile or run the program.</p>
<p>Let's say you are writing a C++/C code and you have a clang compiler. Then you could use c++ and c respectively command with appropriate file handles to compile from vim. Here's the command,</p>
<pre><code>:!c++ % -o %:r
</code></pre>
<p>Here, % stands for the current file and :r removes the extension of the file. So we are creating an executable file without the extension and hence :r. And to execute the file, simply type,</p>
<pre><code>:!%:r
</code></pre>
<p>this will execute the program. If Windows users face an issue like a file is not identified as a batch file or executable file, you should add .exe after :r to make it an exe file.</p>
<p>For Python or other interpreted language, it is quite simple, just the interpreters name with the file name as it is</p>
<pre><code>:!python filename.py
</code></pre>
<pre><code>:!python %
</code></pre>
<p>You could find your preferred language's CLI tool and set it up for your environment and get it plugged with VIM.</p>
<p><strong>Mapping in VIM</strong></p>
<p>But this is not it! There is a long way to master VIM. It's not possible to know everything about anything. So for getting faster I'll introduce VIM's fastest tool and the part where it really shines. Mapping!!</p>
<p>You could map (assign) some keys to do a particular task in different modes. Just imagine typing cpp and running the program or py and running the program,no need to go to the mouse to click run button and close the window. It gets really blazing fast. Only your creativity is the limit. Mapping commands are also called macros.</p>
<pre><code>nmap cpp :!c++ % -o %:r  &amp;&amp;  %:r
</code></pre>
<pre><code>nmap py :!python %
</code></pre>
<p>In the above commands, nmap means in Normal mode MAP cpp and then followed by the command which for this case it is executed in the command mode. There can be various permutations you could make like imap or vmap in Insertion mode and Visual mode and so on.  From mapping,its basically left-hand side mapped to right-hand side. The key combination to the left and command to be executed to the right.</p>
<p>You will surely say, &quot;Do I have to do it every time I run the program or every time I use VIM?&quot;</p>
<p>Well, You can :) but you should not. Because there's a way around it and this is the fuel of VIM or the soul for its Speed.</p>
<p>Yes, I am talking about VIM CONFIGURATION FILE or vimrc</p>
<h1>VIM Run Commands File(vimrc)</h1>
<p>This is really a nasty feature of vim. You could tailor your vim editor as per your needs. You have to write the command once and it will run from the vimrc file. Making a vimrc file is a bit overwhelming in the beginning so, you should have a concrete knowledge of VIM commands and you should remember commands as well.</p>
<ul>
<li>
<p>Make a file called ~/.vimrc</p>
</li>
<li>
<p>You should be using VIM to edit it or write it of course! Type vim ~/.vimrc in terminal or cmd</p>
</li>
<li>
<p>Enter some commands such as set file type indent on , set number , you can find a  <a href="https://vim.fandom.com/wiki/Example_vimrc">sample vimrc file </a> on google and understand its commands and uses.
Paint your plain canvas vimrc with your own creativity.</p>
</li>
</ul>
<p>From here you can add the mappings and other key combinations that you might have been using while learning VIM to your vimrc. VIM also has plugins and more customizable tweaks that can enhance your productivity.</p>
<p>You could always get more help in VIM by using :help commandname. Also you can learn from vimtutor from terminal or CMD.</p>
<p>If you want to dive in more deeper , i link some few resources and video links to get your vim skills to new level.</p>
<p><a href="https://www.vim.org/docs.php">VIM Official Documentation. </a></p>
<p><a href="https://vim.fandom.com/wiki/Vim_Tips_Wiki">VIM fandom Wiki</a></p>
<p><a href="https://catswhocode.com/vim-commands/">130+ VIM commands. </a></p>
<p><a href="https://scotch.io/tutorials/getting-started-with-vim-an-interactive-guide">Vim Interactive Guide.</a></p>
<p><a href="https://www.youtube.com/watch?v=H3o4l4GVLW0&amp;list=PLm323Lc7iSW_wuxqmKx_xxNtJC_hJbQ7R">Vim as your editor.</a></p>
<p>Vim is a tool that can get quite handy for programmers and it stands out from rest of the text editors. By using Vim you will be in very rare people who can make something from almost nothing. Because in VIM you make everything from commands to mapping from customization to functionality. This cannot be done overnight for sure and you would  say you are wasting time in the initial stage but no, you are actually building a foundation for more efficiency and speed. If still you want to use VS Code ,you can use it freely . Because it won't matter much , at the end of the day its programming skills that win the race.</p>
<p><strong>This is quite a huge topic and everything could not be explained in a article, it should be self-explored and learnt from various sources to get its best functioning as per needs. This was a quick guide how and why you should be using VIM as a text editor or as a customizable IDE. As you might have seen its functionality is quite customizable and it is quite powerful tool for a programmer. Happy Coding and VIMing ;)</strong></p>
