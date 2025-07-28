{
  "title": "Neovim + Sourcegraph Cody Plugin Integration",
  "status": "published",
  "slug": "neovim-sourcegraph-cody",
  "date": "2025-04-05T12:33:28Z"
}

<h2>Introduction</h2>
<p>Have you ever used Sourcegraph's Cody? It is a great tool for developers, it is not just another LLM, it is tailored specifically for developers. Cody has some good features that allow parsing of context for the prompt in a smarter way.</p>
<h3>What is Sourcegraph's Cody</h3>
<p>Cody is an AI assistant for developers that understands code context and can generate code. It goes beyond just answering questions - it can write code for you.</p>
<p>The major features of Cody from the rest of the LLMs or Chatbots are:</p>
<ul>
<li>
<p>Cody understands your code context - it reads your open files, repositories, etc. So it can answer questions specifically about your codebase, not just general programming questions.</p>
</li>
<li>
<p>Cody can explain sections of code to you in plain English. This helps ramp up on unfamiliar code bases.</p>
</li>
<li>
<p>Cody integrates into popular editors like VS Code, IntelliJ, Neovim, and others for frictionless use while coding.</p>
</li>
</ul>
<p>For more insights, check out the blog <a href="https://about.sourcegraph.com/blog/all-you-need-is-cody">all you need is Cody</a>. This is a great article about what and how Cody is tailored specifically to assist developers.</p>
<h2>Prerequisites</h2>
<p>To setup sourcegraph on neovim, you will require the following setup:</p>
<ul>
<li>
<p>Neovim 0.9 or above</p>
</li>
<li>
<p>Node.js &gt;= 18.17 (LTS)</p>
</li>
<li>
<p>Cargo (Rust) (optional)</p>
</li>
</ul>
<p>To install neovim latest/nightly release, you can follow the <a href="https://github.com/neovim/neovim/blob/master/INSTALL.md">INSTALL</a> or <a href="https://github.com/neovim/neovim/blob/master/BUILD.md">BUILD</a> documentation of the neovim project.</p>
<p>Cargo is optional, as the plugin will install the binaries itself, however, if you prefer to have cargo, just install it in case something goes wrong.</p>
<h2>Installing sg.nvim</h2>
<p>There is a specific plugin for neovim for interacting with the Sourcegraph products, and Cody is one of them. The <a href="https://github.com/sourcegraph/sg.nvim">sg.nvim</a> is a plugin for integrating sourcegraph search, Cody, and other features provided by soucegraph.</p>
<h4>Using packer.nvim</h4>
<pre><code class="language-lua">use { 'sourcegraph/sg.nvim', run = 'nvim -l build/init.lua' }
use { 'nvim-lua/plenary.nvim' }
</code></pre>
<p>Source your lua file where you have configured all the plugins and then enter the command <code>:PackerInstall</code> or <code>:PackerSync</code> to install the plugin.</p>
<h4>Using vim-plug</h4>
<p>If you are using vim-plug as the plugin manager, you can add the plugin to the configuration as below:</p>
<pre><code class="language-plaintext">Plug 'sourcegraph/sg.nvim', { 'do': 'nvim -l build/init.lua' }

&quot; Required for various utilities
Plug 'nvim-lua/plenary.nvim'

&quot; Required if you want to use some of the search functionality
Plug 'nvim-telescope/telescope.nvim'
</code></pre>
<p>You can source the file and run the command <code>:PlugInstall</code> to install the plugin.</p>
<h4>Using Lazy.nvim</h4>
<p>If you are using Lazy.nvim as the plugin manager, you can add the plugin to the Configuration as below:</p>
<pre><code class="language-lua">return {
  {
    &quot;sourcegraph/sg.nvim&quot;,
    dependencies = { &quot;nvim-lua/plenary.nvim&quot;, &quot;nvim-telescope/telescope.nvim&quot; },

    -- If you have a recent version of lazy.nvim, you don't need to add this!
    build = &quot;nvim -l build/init.lua&quot;,
  },
}
</code></pre>
<p>You can source the file and run the command <code>:Lazy install</code> to install the plugin.</p>
<p>For other installation instructions, you can refer to the <a href="https://github.com/sourcegraph/sg.nvim?tab=readme-ov-file#install">README</a> of sg.nvim.</p>
<h3>Installing Binaries and Building the Plugin</h3>
<p>After the plugin is installed, you can move into the building and setup process of the sourcegraph Cody plugin.</p>
<p>To install the binaries which are the dependencies of the plugin, you can run the command <code>:SourcegraphDownloadBinaries</code> which will force downloading the binaries, making sure that the plugin is properly installed.</p>
<p><img src="https://meetgor-cdn.pages.dev/sg-nvim-build.png" alt="SourcegraphDownloadBinaries Command Output"></p>
<p>To build the plugin, you can simply run the command from within neovim as <code>:SourcegraphBuild</code>, this will rebuild the Sourcegraph rust crates and its dependencies (which might have failed during installation).</p>
<h3>Sourcegraph Authentication</h3>
<p>You need to now authenticate to your Sourcegraph account to use the sourcegraph features such as search and Cody.</p>
<p>You can do that by running the command <code>:SourcegraphLogin</code> in neovim. This will require two inputs, the sourcegraph endpoint and the access token. If you are using sourcegraph cloud and not a self-hosted sourcegraph, you do not need to change the endpoint, just press enter and move ahead. This will redirect you to the browser for authentication and creating an access token. Log in with your credentials to sourcegraph and copy the access token.</p>
<p>This will prompt you back to the neovim interface to provide the access token. Paste the access token there and you will be good to go.</p>
<h3>Health Check</h3>
<p>Once the plugins are installed then you can check the plugin is correctly downloaded by checking the health with the <code>:checkhealth sg</code> command.</p>
<p>Below is the health check report on the sourcegraph plugin in neovim.</p>
<pre><code class="language-plaintext">sg: require(&quot;sg.health&quot;).check()

sg.nvim report ~
- Machine: x86_64, sysname: Linux
- OK Valid nvim version: table: 0x7ffa0b7bce38
- OK Found `cargo` (cargo 1.70.0) is executable
-     Use `:SourcegraphDownloadBinaries` to avoid building locally.
- OK Found `sg-nvim-agent`: &quot;/home/meet/.local/share/nvim/site/pack/packer/start/sg.nvim/dist/sg-nvim-agent&quot;
- OK Found `node` (config.node_executable) is executable.
  Version: '20.10.0'
- OK Found `cody-agent`: /home/meet/.local/share/nvim/site/pack/packer/start/sg.nvim/dist/cody-agent.js
- OK   Authentication setup correctly
- OK     endpoint set to: https://sourcegraph.com
- OK Found correct binary versions: &quot;1.0.5&quot; = &quot;1.0.5&quot;
- OK   Sourcegraph Connection info: {
  access_token_set = true,
  endpoint = &quot;https://sourcegraph.com&quot;,
  sg_nvim_version = &quot;1.0.5&quot;,
  sourcegraph_version = {
		  build = &quot;256174_2023-12-30_5.2-dbb20677711c&quot;,
		  product = &quot;256174_2023-12-30_5.2-dbb20677711c&quot;
  }
  }
- To manage your Cody Account, navigate to: https://sourcegraph.com/cody/manage
- OK Cody Account Information: {
	  chat_limit = 20,
	  chat_usage = 53,
	  code_limit = 500,
	  code_usage = 0,
	  cody_pro_enabled = false,
	  username = &quot;Mr-Destructive&quot;
  }
- OK sg.nvim is ready to run
</code></pre>
<p>At this point, the sourcegraph plugin is ready to be used. However, we need to set up the plugin in neovim with the default configurations.</p>
<h3>Configuration</h3>
<p>In your lua setup files, you can set the plugin like this:</p>
<pre><code class="language-lua">require(&quot;sg&quot;).setup()
</code></pre>
<p>Source the lua file and restart neovim, this should properly make sourcegraph commands available in the editor. After these steps, Cody is right into neovim.</p>
<h2>Usage</h2>
<p>To use the plugin, there are multiple commands available within the editor, the complete list of them is given below:</p>
<pre><code class="language-plaintext">SourcegraphBuild
SourcegraphClear
SourcegraphDownloadBinaries
SourcegraphInfo
SourcegraphLink
SourcegraphLogin
SourcegraphSearch

CodyAsk
CodyChat
CodyDo
CodyRestart
CodyTask
CodyTaskAccept
CodyTaskNext
CodyTaskPrev
CodyTaskView
CodyToggle
</code></pre>
<p>You can get more info about these commands with the <code>help :commandname</code> command. The command are however self explanatory and simple to use.</p>
<h3>Cody Ask</h3>
<p>You can quickly parse a prompt as a string to the <code>:CodyAsk</code> command as <code>:CodyAsk &quot;what is neovim by the way?&quot;</code> and you will get Cody's response in the side vertical split.</p>
<h3>Cody Chat</h3>
<p>You can start Cody chats from neovim with the command <code>:CodyChat</code>. This will open a vertical split with the Cody chat interface, the bottom split has the user input prompt and the upper window will have the generated Cody response. You can enter the prompt in the bottom buffer get into normal mode and hit enter to send the prompt to generate a response from Cody.</p>
<p><img src="https://meetgor-cdn.pages.dev/sg-nvim-cody-chat.png" alt="Sourcegraph Cody Chat Interface"></p>
<h2>Conclusion</h2>
<p>Sourcegraph Cody is a great tool for getting quick solutions to trivial as well as specific problems to the current file/package or project. The context parsing of Cody makes it valuable for developers to ask for specific questions and it answers them really in a straightforward way, without the developer needing to parse the context for the prompt. Thank you for reading. Happy Coding :)</p>
