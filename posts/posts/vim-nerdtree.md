{
  "title": "Vim: NERDTree",
  "status": "published",
  "slug": "vim-nerdtree",
  "date": "2025-04-05T12:33:55Z"
}

<h2>Introduction</h2>
<p><a href="https://github.com/preservim/nerdtree">NERDTree</a> is a great plugin in Vim for managing and navigating Files. Some might prefer fzf, telescope, and other plugins for navigation, NERDTree is not a bad option to begin within Vim. NERDTree allows you to even create/delete/move files and folders flawlessly without much effort, so it becomes a much more viable beginner's plugin.</p>
<h2>Installing NERDTree Plugin</h2>
<p>So, let's start with the Installation of the NERDTree Plugin, it's quite straightforward and simple.</p>
<p>You should have a Plugin-Manager for installing this plugin. It's not mandatory to have one but it becomes much easier to work with. You can choose any between <code>Vundle</code>, <code>Vim-Plug</code>, <code>Pathogen</code> to name a few. It does not really matter with what you use just stick to one but if you are stuck somewhere just switch and it's not a big trouble to use other Plugin Managers as they are quite similar to each other.</p>
<h4>Vundle</h4>
<p>To install a plugin using Vundle, you need to configure Vundle first if you have not already done it. You can find the installation docs <a href="https://github.com/VundleVim/Vundle.vim">here</a>.
After Vundle has been configured in your vimrc you can simply add <code>Plugin 'preservim/nerdtree'</code> between the call begin and end of Vundle, like :</p>
<pre><code class="language-vim">call vundle#begin()
  Plugin 'preservim/nerdtree'
call vundle#end()
</code></pre>
<p>All of your other Plugins will go in between those two statements, i.e. <code>call vundle#begin()</code> and <code>call vundle#end()</code>.
After saving and sourcing the vimrc file, you need to install the plugin using the command <code>:PluginInstall</code>, and there you are all Done!</p>
<h4>Vim-Plug</h4>
<p>To install a plugin using the Vim-Plug manager, you need to configure Vim-Plug if you have not already configured it in your vimrc. You can find the installation docs at the GitHub README of <a href="https://github.com/junegunn/vim-plug">Vim-Plug</a>.
After Vim-Plug has been configured in your vimrc you can simply add <code>Plug 'preservim/nerdtree'</code> between the call plug begin and end statements. Just like:</p>
<pre><code class="language-vim">call plug#begin()
  Plug 'preservim/nerdtree'
call plug#end()
</code></pre>
<p>All of your other Plugins will go in between those two statements, i.e. <code>call plug#begin()</code> and <code>call plug#end()</code>.
After saving and sourcing your vimrc file, you need to now install those plugins using the command <code>:PlugInstall</code>, and that is it!</p>
<h4>Pathogen</h4>
<p>To install any plugin using Pathogen plugin manager, you need to configure Pathogen in your vimrc if you have not done it already. You can find the installation docs on <a href="https://github.com/tpope/vim-pathogen">Pathogen.vim</a>.
After Pathogen has been configured in your vimrc, you can clone the git repository of that plugin into your local machine and then activate it using Pathogen.</p>
<pre><code>git clone https://github.com/preservim/nerdtree.git ~/.vim/bundle/nerdtree
</code></pre>
<p>After cloning the repository, you can add this to your vimrc where you have configured it. It's a kind of DIY manager in terms of managing the folders of the plugin.</p>
<pre><code class="language-vim">call plug#begin()
call pathogen#infect()
syntax on
filetype plugin indent on
</code></pre>
<p>After this, you need to run this command to get docs and help with the plugins,</p>
<p><code>:help tags ~/.vim/bundle/nerdtree/doc/</code> or <code>:help tags</code></p>
<p>And there you are done with the plugin installed.</p>
<p>There are other Plugin managers as well, but these three are the most widely supported ones and they work out of the box, surely explore for yourself and find the perfect one for you.</p>
<h2>Activating and Using NERDTree</h2>
<p>Now, we actually need to use NERDTree, for that we can type in <code>:NERDTree</code> in any folder in our local machine, and there should be a window open a vertical split to the left, just like this:
<img src="https://i.imgur.com/KU2vMxO.png" alt="NERDTree activate"></p>
<p>After this, you can use <!-- raw HTML omitted --> that is <strong>CTRL+W twice</strong> twice to switch back and forth between windows. You can also use <strong>CTRL+W and HJKL</strong> to move in directions in the windows. For further guides and key-bindings in Window-Splits, you can read my article <a href="https://mr-destructive.github.io/techstructive-blog/vim/2021/08/06/Vim-Window-Splits.html">here</a>.</p>
<p>Now, you can navigate to the file/folders using HJKL or arrows keys(not preferred). You can even use the numbers before the HJKL to jump and hop around the large codebases, this integrates really well with base Vim key-bindings.
You can quiet the NERDTree window by just pressing <code>q</code> or <code>:q</code>, definitely the former is efficient. You can open/collapse the folders also using the enter key to open the file in the current buffer. But hey that's quite limited, what have you ever seen!</p>
<h4>Open File in Splits</h4>
<p>You can open a file in a horizontal split using the key <code>i</code> on the file. You can open a file in Vertical split using the <code>s</code> key keeping the current highlight in NERDTree on the file which you would like to open. This can be really a great feature to have while opening multiple files and file structures.</p>
<h4>Managing Files/Folders using NERDTree</h4>
<p>You can create files using the NERDTree window by pressing m inside the particular folder where you want to. If you want to create a file in the root folder, you can go to the topmost file location and press <code>m</code> inside the NERDTree window. If you press <code>m</code>, you will be able to see different kinds of options namely:</p>
<ol>
<li>
<p>Add a child node.(<code>a</code>)</p>
<p>We can create a file or a folder using the key <code>a</code> or simply <code>Enter</code> to create the file in the current highlighted location.</p>
</li>
<li>
<p>Move the Current Node. (<code>m</code>)</p>
<p>We can create a file or a folder using the key <code>a</code> or simply <code>Enter</code> to create the file in the current highlighted location.</p>
</li>
<li>
<p>Delete the current Node. (<code>d</code>)</p>
<p>We can move the currently highlighted file/folder to any other directory using the file manager itself.</p>
</li>
<li>
<p>Open the current Node in the system Text-Editor.(<code>o</code>)</p>
<p>We can delete the file/folder which is currently selected on the NERDTree menu.</p>
</li>
<li>
<p>Copy the current Node. (<code>c</code>)</p>
<p>We can open the file in the system-default text-editor using the key <code>o</code>.</p>
</li>
<li>
<p>Copy the Path to the clipboard.(<code>p</code>)</p>
<p>We can copy the current file/folder or a node using the command <code>c</code>.</p>
</li>
<li>
<p>List the Current Node. (<code>l</code>)</p>
<p>We can list the file/folder i.e to display its properties the read/write/execute permissions, date modified and created, etc.</p>
</li>
<li>
<p>Run system Command in this folder. (<code>s</code>)</p>
<p>We can run system commands or shell/terminal commands using the key <code>s</code>, For windows, we open the COMMAND PROMPT, and in Linux and macOS, it is terminal.</p>
</li>
</ol>
<p>You can quit that window by pressing <code>Esc</code>.</p>
<p>Here is some of the Screencast of me demonstrating the NERDTree plugin features and the edit options.
<img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1630423628366/zZE7R5aL7.gif" alt="vimnerd.gif"></p>
<p>This is just self-explanatory and beginner-friendly sets of commands, it becomes really easy to do this with some custom macros as we'll talk about in the next sections.</p>
<h2>Macros to open file tree</h2>
<p>You can make a key-binding to open the NERDTree,</p>
<pre><code class="language-vim">nnoremap &lt;C-n&gt; :NERDTree&lt;CR&gt;
</code></pre>
<p>You can map anything instead of <code>&lt;C-n&gt;</code>, most people use <code>&lt;leader&gt;</code> but it's easy to use <code>CTRL+N</code> for me, it's just personal preference.</p>
<p>If you do not like to open NERDTree again and again, you can keep it open whatsoever using the custom key-binding in your vimrc.</p>
<pre><code class="language-vim">autocmd VimEnter * NERDTree
</code></pre>
<p>This will open the NERDTree automatically for you when you open Vim, Ya I get it, it's not needed every time but most of the time a developer is switching between the files.</p>
<h2>Enabling Autoreload</h2>
<p>We can auto-reload the NERDTree window when there is a change in the file structure, i.e. a file/folder is deleted/created/moved/etc. We again need to set this in our vimrc:</p>
<pre><code class="language-vim">autocmd BufEnter NERD_tree_* | execute 'normal R'
au CursorHold * if exists(&quot;t:NerdTreeBufName&quot;) | call &lt;SNR&gt;15_refreshRoot() | endif
</code></pre>
<p>This will reload the NERDTree when the cursor is in the NERDTree's window. This could be really time-saving and a nice quick configuration to enhance the efficiency of your Text-editing.</p>
<h2>Enabling Autorefresh for change in the current directory</h2>
<p>We can also reload the NERDTree when we change the directory. The above-mentioned command is not sufficient to do that, we have to add another set of configurations.</p>
<pre><code class="language-vim">augroup DIRCHANGE
    au!
    autocmd DirChanged global :NERDTreeCWD
augroup END
</code></pre>
<p>By adding this to your vimrc, you will refresh the NERDTree every time you enter or change the current directory. This is also a great addition to have to save time by reloading the Window for every change in the path, if you are looking for something among a huge code-base, this works a charm.</p>
<h2>Auto close</h2>
<p>You need to close the NERDTree manually each time you want to exit out of it, but this can also be automated just for the sake of simplicity and effectiveness in <strong>QUITTING VIM</strong>.</p>
<pre><code class="language-vim">autocmd bufenter * if (winnr(&quot;$&quot;) == 1 &amp;&amp; exists(&quot;b:NERDTree&quot;) &amp;&amp; b:NERDTree.isTabTree()) | q | endif
</code></pre>
<p>This will close the NERDTree window if it is the only open window. That can be frustrating at moments but the majority of the time this is a great addon indeed.</p>
<h2>Packing it together</h2>
<p>So, we have learned the basics of using and modifying NERDTree according to our needs, to put it together, you can use this snippet directly into your vimrc and enjoy the flawless experience.</p>
<pre><code class="language-vim">&quot; Open nerdtree window on opening Vim
autocmd VimEnter * NERDTree

&quot; Refresh the current folder if any changes
autocmd BufEnter NERD_tree_* | execute 'normal R'
au CursorHold * if exists(&quot;t:NerdTreeBufName&quot;) | call &lt;SNR&gt;15_refreshRoot() | endif

&quot;Reload the window if directory is changed
augroup DIRCHANGE
    au!
    autocmd DirChanged global :NERDTreeCWD
augroup END

&quot;Close nerdtree automatically if it is theonly window open
autocmd bufenter * if (winnr(&quot;$&quot;) == 1 &amp;&amp; exists(&quot;b:NERDTree&quot;) &amp;&amp; b:NERDTree.isTabTree()) | q | endif
</code></pre>
<h2>Conclusion:</h2>
<p>So, We were able to make Vim a better place to work with. Making it easier to navigate along with files and folders. Configuring the NERDTree Plugin, customizing the look and functionality of Vim as per the needs.
NERDTree is a great plugin no matter how you use it. It makes Vim more viable as a text editor for daily use and that also in an efficient and clean way. Surely there might be other plugins that are super powerful and blazing fast, NERDTree provides a good UI as well by providing a graphical representation of the File structure that enhances its usage.
That is what Vim is about, learning every day some things to change the way to edit. Thank you for reading. Happy Viming and Coding :)</p>
<h3>References:</h3>
<ul>
<li><a href="https://github.com/preservim/nerdtree">NERDTree - docs</a></li>
<li><a href="https://stackoverflow.com/questions/8793489/nerdtree-reload-new-files/8794468">Refresh NERDTree</a></li>
<li><a href="https://vi.stackexchange.com/questions/31050/how-can-i-make-nerdtree-update-root-to-the-current-directory-when-i-change-direc">Reload NERDTree on Directory change</a></li>
<li><a href="https://stackoverflow.com/questions/1447334/how-to-add-nerdtree-to-your-vimrc">Open NERDTree in Vim by default</a></li>
<li><a href="https://stackoverflow.com/questions/2066590/automatically-quit-vim-if-nerdtree-is-last-and-only-buffer">Close NERDTree automatically</a></li>
</ul>
