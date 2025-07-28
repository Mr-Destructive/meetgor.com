{
  "title": "Configure Neovim in Lua",
  "status": "published",
  "slug": "neovim-vimscript-to-lua",
  "date": "2025-04-05T12:33:39Z"
}

<h2>Introduction</h2>
<p>It has been a while since I have written a Vim article. Finally, I got some ideas after configuring my Neovim setup for Lua. I recently migrated to Ubuntu a couple of months back and it has been a cool change from Windows 7!</p>
<p>In this article, we'll see how you can set up neovim for Lua. Since Neovim 0.5, it supports lua out of the box, so in the recent release 0.7, it added more native support to lua making it a lot easier to configure and play with neovim. So, we will see how we can use lua to convert all the 200 liner vimscript into lua (We can even have packages and modules:) We will cover how to configure your keymaps, pull up all the plugins, vim options, and other customizations.</p>
<h2>Why move to Lua?</h2>
<p>I have used Vimscript for quite a while now, configured it as per my needs, and also made a few plugins like <a href="https://github.com/Mr-Destructive/frontmatter.vim">frontmatter</a>, <a href="https://github.com/Mr-Destructive/dj.vim">dj.vim</a>, and <a href="https://github.com/Mr-Destructive/commenter.vim">commenter</a> which are quite clunky and not robust in terms of usage and customizability. Vimscript is good but it's a bit messy when you want extreme customization.</p>
<p>I recently wanted to go full Neovim, I was kind of stuck in Vimscript and some of my plugins work for me but it might not work for others, it might be a piece of gibberish vimscript dumped. So, Why not have full native experience in Neovim if you can. It has now baked-in support for LSP, Debugging, Autocommands, and so much more. If you have Neovim 0.5+ you should be good to go full lua.</p>
<h2>Backup Neovim Config</h2>
<p>Firstly, keep your original neovim/vim init files safe, back them up, make a copy and save it with a different name like <code>nvim_config.vim</code>. If you already have all your config files backed up as an ansible script or dotfiles GitHub repository, you can proceed ahead.</p>
<p>But don't keep the <code>init.vim</code> file as it is in the <code>~/.config/nvim</code> folder, it might override after we configure the lua scripts.</p>
<h2>Basic Configuration</h2>
<p>So, I assume you have set up Neovim, If not you need to follow some simple steps like downloading the package and making sure your neovim environment is working with vimscript first. The <a href="https://github.com/neovim/neovim/wiki/Installing-Neovim">Neovim Wiki</a> provides great documentation on how to install neovim on various platforms using different methods.</p>
<p>After your neovim is set up and you have a basic configuration, you can now start to migrate into lua.
Create a <code>init.lua</code> file in the same path as your <code>init.lua</code> file is i.e. at <code>~/.config/nvim</code> or <code>~/AppData/Local/nvim/</code> for Windows. That's why it is recommended to keep the initial configuration vimscript file in a safe place. While migrating from vimscript to lua, once the lua file is created and the next time you restart neovim, the default settings will be from <code>init.lua</code> and not <code>init.vim</code>, so be prepared.</p>
<p>Firstly, you need to configure some options like <code>number</code>, <code>syntax highlighting</code>, <code>tabs</code>, and some <code>keymaps</code> of course. We can use the <code>vim.opt</code> method to set options in vim using lua syntax. So, certain corresponding vim options would be converted as follows:</p>
<p>If you have the following kind of settings in your vimrc or init.vim:</p>
<pre><code class="language-vimscript">-- vimscript
set number
set tabstop=4 
set shiftwidth=4 
set softtabstop=0 
set expandtab 
set noswapfile
</code></pre>
<p>The above settings are migrated into lua syntax as follows:</p>
<pre><code class="language-lua">--lua
vim.opt.number = true
vim.opt.tabstop = 4
vim.opt.shiftwidth = 4
vim.opt.softtabstop = 0
vim.opt.expandtab = true
vim.opt.swapfile = false
</code></pre>
<p>You can set other options in your config file accordingly. If you get sick of writing <code>vim.opt.</code> again and again, you can use a variable set to <code>vim.opt</code> and then access that variable to set the options. Something of the lines as below:</p>
<pre><code class="language-lua">local set = vim.opt

set.number = true
set.tabstop = 4
set.shiftwidth = 4
set.softtabstop = 0
set.expandtab = true
set.swapfile = false
</code></pre>
<p>We can create a variable in lua like <code>local variable_name = something</code> so, we have created a variable <code>set</code> which is assigned to the value of <code>vim.opt</code> which is a method(function) in lua to set the options from the vimscript environment. Finally, access the <code>set</code> keyword to set the options. Using the <code>set</code> word is not necessary, you can use whatever you want.</p>
<p>After setting up the basic options, you can source the file with <code>:so %</code> from the command mode. Just normally as you source the config files.</p>
<h3>Using Lua in Command Mode</h3>
<p>We can use the lua functions or any other commands from the command mode in neovim using the lua command. Just prefix the command with <code>:lua</code> and after that, you can use lua syntax like function calling or setting variables, logging things, etc.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1657380885/blog-media/lua_in_command_mode.gif" alt="Lua in Command Mode"></p>
<h2>Adding Keymaps</h2>
<p>Now, that we have some basic config setup, we can quickly get the keymaps. It's not that hard to make keymaps to set up in lua. To create keymaps, we have two options. One is compatible with Neovim and the other is compatible with Vim as well.</p>
<ol>
<li>vim.keymap.set OR</li>
<li>vim.api.nvim_set_keymap</li>
</ol>
<p>Personally, I don't see a difference in terms of usage but <a href="https://github.com/neovim/neovim/blob/master/runtime/lua/vim/keymap.lua#L51">vim.keymap.set</a> is a wrapper around <a href="https://github.com/neovim/neovim/blob/master/src/nvim/api/vim.c#L1451">nvim_set_keymap</a>. So, it is really a matter of personal preference which you want to use.</p>
<p>So, both the functions have quite similar parameters:</p>
<pre><code class="language-lua">vim.keymap.set({mode}, {lhs}, {rhs}, {options})

vim.api.nvim_set_keymap({mode}, {lhs}, {rhs}, {options})
</code></pre>
<p>The advantage of <code>vim.keymap.set</code> over <code>vim.api.nvim_set_keymap</code> is that it allows directly calling lua functions rather than calling vimscripty way like <code>:lua function()</code>, so we directly can use lua code in the RHS part of the function parameter.</p>
<p>Let's take a basic example mapping:</p>
<pre><code>vim.keymap.set('n', 'Y', 'yy', {noremap = false})
</code></pre>
<p>Here, we have mapped <code>Shift+y</code> with the keys <code>yy</code> in <code>n</code>ormal mode. The first argument is the mode, it can be a single-mode like <code>'n'</code>, <code>'v'</code>, <code>'i'</code>, etc., or a multi-mode table like <code>{'n', 'v'}</code>, <code>{'v', 'i'}</code>, etc.</p>
<p>The next parameter is also a string but it should be the key for triggering the mapping. In this case, we have used <code>Y</code> which is <code>Shift + y</code>, it can be any key shortcut you want to map.</p>
<p>The third parameter is the string which will be the command to be executed after the key has been used. Here we have used the keys <code>yy</code>, if the map is from a command mode, you will be using something like <code>':commands_to_be executed'</code> as the third parameter.</p>
<p>The fourth parameter which is optional can contain <a href="https://neovim.io/doc/user/api.html#:~:text=nvim_set_keymap(%7Bmode%7D%2C%20%7Blhs%7D%2C%20%7Brhs%7D%2C%20%7B*opts%7D)%20%20%20%20%20%20%20%20%20%20%20%20%20*nvim_set_keymap()*">special arguments</a>. We have set a default option which is <code>noremap</code> as true, the options are not string but lua tables instead, so it can be similar to python dictionary or a map kind of a structure with a key value pair.</p>
<p>One more important aspect in keymapping might about the leader key, you can set your leader key by using the global vim options with <code>vim.g</code> and access <code>mapleader</code> to set it to the key you wish. This will make the <code>leader</code> key available to us and thereafter we can map the leader key in custom mappings.</p>
<pre><code>vim.g.mapleader = &quot; &quot;
</code></pre>
<p>Here, I have set my leader key to the <code>&lt;Space&gt;</code> key. Now, we can map keys to the existing keymaps in the vimscript. Let's map some basic keymaps first and then after setting up the plugins,we can move into plugin-specific mappings.</p>
<p>You can also use <code>vim.api.nvim_set_keymap</code> function with the same parameters as well.</p>
<pre><code class="language-lua">vim.keymap.set('n', '&lt;leader&gt;w', ':w&lt;CR&gt;',{noremap = true})
vim.keymap.set('n', '&lt;leader&gt;q', ':q!&lt;CR&gt;',{noremap = true})
vim.keymap.set('n', '&lt;leader&gt;s', ':so %&lt;CR&gt;',{noremap = true})
vim.keymap.set('n', '&lt;leader&gt;ev', ':vsplit $MYVIMRC&lt;CR&gt;',{noremap = true})
vim.keymap.set('n', '&lt;leader&gt;sv', ':w&lt;CR&gt;:so %&lt;CR&gt;:q&lt;CR&gt;',{noremap = true})

OR

vim.api.nvim_set_keymap('n', '&lt;leader&gt;w', ':w&lt;CR&gt;',{noremap = true})
vim.api.nvim_set_keymap('n', '&lt;leader&gt;q', ':q!&lt;CR&gt;',{noremap = true})
vim.api.nvim_set_keymap('n', '&lt;leader&gt;s', ':so %&lt;CR&gt;',{noremap = true})
vim.api.nvim_set_keymap('n', '&lt;leader&gt;ev', ':vsplit $MYVIMRC&lt;CR&gt;',{noremap = true})
vim.api.nvim_set_keymap('n', '&lt;leader&gt;sv', ':w&lt;CR&gt;:so %&lt;CR&gt;:q&lt;CR&gt;',{noremap = true})
</code></pre>
<p>If, you don't like writing <code>vim.keymap.set</code> or <code>vim.api.nvim_set_keymap</code> again and again, you can create a simpler function for it. In lua a function is created just like a variable by specifying the scope of the function i.e. local followed by the <code>function</code> keyword and finally the name of the function and parenthesis. The function body is terminated by the <code>end</code> keyword.</p>
<pre><code class="language-lua">function map(mode, lhs, rhs, opts)
    local options = { noremap = true }
    if opts then
        options = vim.tbl_extend(&quot;force&quot;, options, opts)
    end
    vim.api.nvim_set_keymap(mode, lhs, rhs, options)
end
</code></pre>
<p>Now, in this function <code>map</code>, we have passed in the same parameters like the <code>vim.keymap.set</code> function takes but we have just parsed the function in a shorter and safer way by adding <code>noremap = true</code> by default. So this is just a helper function or a verbose function for calling the vim.keymap.set function.</p>
<p>To use this function, we can simply call <code>map</code> with the same arguments as given to the prior functions.</p>
<pre><code class="language-lua">map('n', '&lt;leader&gt;w', ':w&lt;CR&gt;')
map('n', '&lt;leader&gt;q', ':q!&lt;CR&gt;')
map('n', '&lt;leader&gt;s', ':so %&lt;CR&gt;')
</code></pre>
<p>Notice here, though, we have not passed the <code>{noremap = true}</code> as it will be passed by default to the <code>vim.api.nvim_set_keymap</code> or <code>vim.keymap.set</code> function via the custom map function.</p>
<p>If you want some more examples, here are some additional mapping specific to languages, meant for compiling or running scripts with neovim instance.</p>
<pre><code class="language-lua">-- vimscript

nnoremap cpp :!c++ % -o %:r &amp;&amp; %:r&lt;CR&gt;
nnoremap c, :!gcc % -o %:r &amp;&amp; %:r&lt;CR&gt;
nnoremap py :!python %&lt;cr&gt;
nnoremap go :!go run %&lt;cr&gt;
nnoremap sh :!bash %&lt;CR&gt;


-- lua

map('n', 'cpp' ':!c++ % -o %:r &amp;&amp; %:r&lt;CR&gt;')
map('n', 'c,' ':!gcc % -o %:r &amp;&amp; %:r&lt;CR&gt;')
map('n', 'py' ':!python %&lt;cr&gt;')
map('n', 'go' ':!go run %&lt;cr&gt;')
map('n', 'sh' ':!bash %&lt;cr&gt;')

</code></pre>
<p>So, this is how we can set up our keymaps in lua. You can customize this function as per your needs. These are just made for the understanding purpose of how to reduce the repetitive stuff in the setup.</p>
<p><strong>If you are really stuck up and not feeling to convert those mappings into lua then I have a function that can do it for you, check out my dotfile repo -&gt; <a href="https://github.com/Mr-Destructive/dotfiles/blob/master/nvim/lua/destructive/options.lua#L9">keymapper</a></strong></p>
<h2>Adding Plugin Manager</h2>
<p>Now, we really missing some plugins, aren't we? So, the neovim community has some good choices for using a new plugin manager written in core lua. It is usually a good idea to move into lua completely rather than switching to and fro between vimscript and lua.</p>
<p>So, <a href="https://github.com/wbthomason/packer.nvim">Packer</a> is the new plugin manager for Neovim in Lua, there is other plugin managers out there as well like <a href="https://github.com/savq/paq-nvim">paq</a>. If you don't want to switch with the plugin manager, you can still use vim-based plugin managers like <a href="https://dev.to/vonheikemen/neovim-using-vim-plug-in-lua-3oom">Vim-Plug</a>.</p>
<p>So, let's install the Packer plugin manager into Neovim. We simply have to run the following command in the console and make sure the plugin manager is configured correctly.</p>
<pre><code># Linux

git clone --depth 1 https://github.com/wbthomason/packer.nvim\
~/.local/share/nvim/site/pack/packer/start/packer.nvim


# Windows

git clone https://github.com/wbthomason/packer.nvim &quot;$env:LOCALAPPDATA
vim-data\site\pack\packer\start\packer.nvim&quot;
</code></pre>
<p>Now, if you open a new neovim instance and try to run the command <code>:PackerClean</code>, and no error pops out that means you have configured it correctly. You can move ahead to installing plugins now. Yeah! PLUG-IN time!</p>
<pre><code class="language-lua">return require('packer').startup(function()
end)
</code></pre>
<p>First try to source the file, if it throws out errors it shouldn't try to fix the installation path of Packer. If the command succeded we can finally pull up some plugins.</p>
<p>Below are some of the plugins that you can use irrespective of what language preferences you would have. This includes basic dev-icons for the status line as well as the explorer window file icons. As usual, add your plugins and make them yours.</p>
<pre><code class="language-lua">
return require('packer').startup(function()
  use 'wbthomason/packer.nvim'
  use 'tpope/vim-fugitive'
  use {
    'nvim-lualine/lualine.nvim',
    requires = { 'kyazdani42/nvim-web-devicons', opt = true }
  }
  use 'tiagofumo/vim-nerdtree-syntax-highlight'
  use 'kyazdani42/nvim-web-devicons'
  use 'vim-airline/vim-airline'
  use 'vim-airline/vim-airline-themes'
end)
</code></pre>
<p>After adding the list of your plugins, you need to source the file and then install the plugins with the command <code>:PackerInstall</code>. This function will install all the plugins after the file has been sourced so make sure you don't forget it.</p>
<h2>Colors and Color Themes</h2>
<p>You might fancy some good-looking and aesthetic setup for neovim of course! In Neovim, we already have a wide variety of configurations to set up like color schemes, GUI colors, terminal colors, etc. You can pick up a color scheme from a large list of awesome color schemes from <a href="https://github.com/topics/neovim-colorscheme">GitHub</a>.</p>
<p>After choosing the theme, plug it in the packer plugin list which we just created and source the file and finally run <code>:PackerInstall</code>. This should install the plugin. You can then set the colorscheme as per your preference, first view the color scheme temporarily on the current instance with the command <code>:colorscheme colorscheme_name</code>.</p>
<pre><code class="language-lua">return require('packer').startup(function()
  use 'wbthomason/packer.nvim'
  -- 
  use 'Mofiqul/dracula.nvim'
  --
end)
</code></pre>
<p>You can then add the command to set it as the default color scheme for Neovim.</p>
<pre><code class="language-lua">vim.cmd [[silent! colorscheme dracula]]
</code></pre>
<p>You can change the background color, text color, icons style and terminal and gui style separately with the augroup and setting the colorscheme commands.</p>
<pre><code class="language-lua">vim.api.nvim_command([[
    augroup ChangeBackgroudColour
        autocmd colorscheme * :hi normal termbg=#000030 termfg=#ffffff
        autocmd colorscheme * :hi Directory ctermfg=#ffffff
    augroup END
]])
vim.o.termguicolors = true
</code></pre>
<p>Here, I have used the background and foreground colors of the terminal variant of Neovim. Also for the Directory Explorer i.e. netrw, I have changed the terminal foreground. This you can configure as per your needs, Though this is still vimscripty, there are Autocommands and autogroups available sooner in Neovim.</p>
<h2>Separating Configurations</h2>
<p>If you wish to keep all your config in one file i.e. <code>init.lua</code> file, you can, though you can separate out things like <code>keymaps</code>, <code>plugins</code>, <code>custom_options</code> or if you have <code>lsp</code> configurations into separate lua packages or creating a separate module. This helps in loading up things as per requirement and also looks readable, making it a lot easier to test out things without breaking the whole config.</p>
<p>Definitely, there will be personal preferences and pros and cons of both approaches, you can pick up whatever suits your style.</p>
<h3>Creating separate packages</h3>
<p>To create a separate package, we can simply add a file in the same folder as <code>init.vim</code> i.e. in the folder <code>~/.config/nvim</code> or equivalent for windows. The package name can be any valid filename with the <code>lua</code> extension.</p>
<p>For instance, you can create a package for all your keymaps and load it in the <code>init.lua</code> as per the order you want to load them. It can be at the top, or else if you have certain base settings lower in the init file, these might not reflect or load up in the keymap package so better to load them after some of the core settings have been set.</p>
<p>Let's create the package and dump all our maps into the keymap file package.</p>
<pre><code class="language-lua">-- ~/.config/nvim/keymap.lua

function map(mode, lhs, rhs, opts)
    local options = { noremap = true }
    if opts then
        options = vim.tbl_extend(&quot;force&quot;, options, opts)
    end
    vim.api.nvim_set_keymap(mode, lhs, rhs, options)
end

map('n', '&lt;leader&gt;w', ':w&lt;CR&gt;')
map('n', '&lt;leader&gt;q', ':q!&lt;CR&gt;')
map('n', '&lt;leader&gt;s', ':so %&lt;CR&gt;')
map('n', 'cpp' ':!c++ % -o %:r &amp;&amp; %:r&lt;CR&gt;')
map('n', 'c,' ':!gcc % -o %:r &amp;&amp; %:r&lt;CR&gt;')
map('n', 'py' ':!python %&lt;cr&gt;')
map('n', 'go' ':!go run %&lt;cr&gt;')
map('n', 'sh' ':!bash %&lt;cr&gt;')

-- more keymaps

</code></pre>
<p>So, this might work if you don't have any plugin-related keymaps as it would require those functions or objects available to execute. So, we might also want to separate plugins and load them first in our keymaps or in the init file.</p>
<p>Now, there needs to be a way for grabbing a package. Yes, there is basically like import in python or any other programming language, lua has <code>require</code> keyword for importing packages. Since the <code>init</code> file and the <code>keymaps</code> are in the same folder path, we can simply say, <code>require(&quot;keymap&quot;)</code> in our <code>init.lua</code> file. Now, it depends on your config where you want to load this package. At the top i.e. at the beginning of neovim instance or after loading the plugins down.</p>
<pre><code class="language-lua">-- init.lua

require(&quot;keymaps&quot;)

-- At the top
-- OR
-- After loading Packer plugins
</code></pre>
<p>So, now you can separate all your configs as per your requirement. It is like splitting up a puzzle and again combining them. Similar package can be created for <code>plugins</code>, <code>options</code> or <code>lsp</code> configurations.</p>
<h3>Creating a separate module</h3>
<p>Now, we have seen how to create a lua package and loading in neovim. We also can create modules in neovim configuration. For instance, first, the init file is loaded, Other files might not be required hence it is not loaded by default, it is only loaded when <code>require</code>ed. So, we can create a module in lua with a folder, and inside of it, we can have packages as we had in the previous method. Also, every module can have a init file loaded first when we require that module. The rest of the packages in the module are thus made available thereafter.</p>
<ul>
<li>Module is a component not loaded by default</li>
<li>Only loaded when required (literally require)</li>
<li>Every module can have a <code>init.lua</code> file loaded first when required.</li>
<li>Require a package in module from outside -&gt; <code>require('module_name.package_name')</code></li>
</ul>
<p>So, in neovim, we need to create a <code>lua</code> folder for placing all our modules so that the lua runtime is picked up correctly. Inside this lua folder, we can create a module basically a folder. This folder name can be anything you like, I like to keep it my nickname, you can use whatever you prefer.</p>
<pre><code># ~/.config/nvim

-- init.lua
-- lua/
    -- module_name/
        -- init.lua
        -- package_name.lua
        -- keymaps.lua
</code></pre>
<p>Now, we can create packages in this module. You can move your keymaps package inside this folder. The keymaps package is nothing here but an example provided in the previous section for creating a package. You can basically dump all your keymaps in a single lua file and then import it from the init file. Similarly you can create a package inside a module and import it from the init file of the module(local init file <code>~/.config/nvim/lua/module_name/init.lua</code>) or the global init file(<code>~/.config/nvim/init.lua</code>).</p>
<p>To load the packages, you have to use the same require statement irrespective of where you are importing from i.e. either from the module or from the global init file. The require statement would look like the following <code>require(&quot;module_name/package_name&quot;)</code>. Now, we can use the keymaps package to import from the module init file and then import the module from the global init file. To import a module, we can simply use the module name in the require string as <code>require(&quot;module_name&quot;)</code>.</p>
<pre><code class="language-lua">-- ~/.config/nvim

-- lua/module_name/options.lua

vim.opt.number = true
vim.opt.tabstop = 4
vim.opt.swapfile = false


-- lua/module_name/plugins.lua

require(&quot;module_name.options&quot;)
return require('packer').startup(function()
  use 'wbthomason/packer.nvim'
  --plugins
end)

-- lua/module_name/keymap.lua

require(&quot;module_name.plugins&quot;)
-- maps()


-- lua/module_name/init.lua

require(&quot;module_name.keymaps)


-- init.lua

require(&quot;module_name&quot;)

</code></pre>
<p>So, this is how we can create modules and packages for configurations in neovim using lua. This is also a kind of a structure for creating your own neovim plugin with lua!</p>
<p>For further references, you can check out my <a href="https://github.com/Mr-Destructive/dotfiles">dotfiles</a>.</p>
<h3>References</h3>
<ul>
<li><a href="https://vonheikemen.github.io/devlog/tools/configuring-neovim-using-lua/">Configure Neovim for Lua</a></li>
<li><a href="https://alpha2phi.medium.com/neovim-for-beginners-init-lua-45ff91f741cb">Neovim with Lua for beginners</a></li>
<li><a href="https://www.youtube.com/c/TJDeVries/videos">TJ Devries Youtube</a></li>
</ul>
<h2>Conclusion</h2>
<p>So, that is just a basic overview of how you can get your neovim configured for lua. It is a great experience to just create such a personalized environment and play with it and have fun. You might have hopefully configured Neovim for Lua from this little guide. Maybe in the next article, I'll set up LSP in Neovim. If you have any queries or feedback, please let me know. Thank you for reading.</p>
<p>Happy Viming :)</p>
