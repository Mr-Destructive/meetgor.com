---
article_html: "<h2>Introduction</h2>\n<p>Have you ever used Sourcegraph's Cody? It
  is a great tool for developers, it is not just another LLM, it is tailored specifically
  for developers. Cody has some good features that allow parsing of context for the
  prompt in a smarter way.</p>\n<h3>What is Sourcegraph's Cody</h3>\n<p>Cody is an
  AI assistant for developers that understands code context and can generate code.
  It goes beyond just answering questions - it can write code for you.</p>\n<p>The
  major features of Cody from the rest of the LLMs or Chatbots are:</p>\n<ul>\n<li>\n<p>Cody
  understands your code context - it reads your open files, repositories, etc. So
  it can answer questions specifically about your codebase, not just general programming
  questions.</p>\n</li>\n<li>\n<p>Cody can explain sections of code to you in plain
  English. This helps ramp up on unfamiliar code bases.</p>\n</li>\n<li>\n<p>Cody
  integrates into popular editors like VS Code, IntelliJ, Neovim, and others for frictionless
  use while coding.</p>\n</li>\n</ul>\n<p>For more insights, check out the blog <a
  href=\"https://about.sourcegraph.com/blog/all-you-need-is-cody\">all you need is
  Cody</a>. This is a great article about what and how Cody is tailored specifically
  to assist developers.</p>\n<h2>Prerequisites</h2>\n<p>To setup sourcegraph on neovim,
  you will require the following setup:</p>\n<ul>\n<li>\n<p>Neovim 0.9 or above</p>\n</li>\n<li>\n<p>Node.js
  &gt;= 18.17 (LTS)</p>\n</li>\n<li>\n<p>Cargo (Rust) (optional)</p>\n</li>\n</ul>\n<p>To
  install neovim latest/nightly release, you can follow the <a href=\"https://github.com/neovim/neovim/blob/master/INSTALL.md\">INSTALL</a>
  or <a href=\"https://github.com/neovim/neovim/blob/master/BUILD.md\">BUILD</a> documentation
  of the neovim project.</p>\n<p>Cargo is optional, as the plugin will install the
  binaries itself, however, if you prefer to have cargo, just install it in case something
  goes wrong.</p>\n<h2>Installing sg.nvim</h2>\n<p>There is a specific plugin for
  neovim for interacting with the Sourcegraph products, and Cody is one of them. The
  <a href=\"https://github.com/sourcegraph/sg.nvim\">sg.nvim</a> is a plugin for integrating
  sourcegraph search, Cody, and other features provided by soucegraph.</p>\n<h4>Using
  packer.nvim</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">use</span> <span class=\"p\">{</span> <span class=\"s1\">&#39;sourcegraph/sg.nvim&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">run</span> <span class=\"o\">=</span> <span
  class=\"s1\">&#39;nvim -l build/init.lua&#39;</span> <span class=\"p\">}</span>\n<span
  class=\"n\">use</span> <span class=\"p\">{</span> <span class=\"s1\">&#39;nvim-lua/plenary.nvim&#39;</span>
  <span class=\"p\">}</span>\n</pre></div>\n\n</pre>\n\n<p>Source your lua file where
  you have configured all the plugins and then enter the command <code>:PackerInstall</code>
  or <code>:PackerSync</code> to install the plugin.</p>\n<h4>Using vim-plug</h4>\n<p>If
  you are using vim-plug as the plugin manager, you can add the plugin to the configuration
  as below:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>Plug
  &#39;sourcegraph/sg.nvim&#39;, { &#39;do&#39;: &#39;nvim -l build/init.lua&#39;
  }\n\n&quot; Required for various utilities\nPlug &#39;nvim-lua/plenary.nvim&#39;\n\n&quot;
  Required if you want to use some of the search functionality\nPlug &#39;nvim-telescope/telescope.nvim&#39;\n</pre></div>\n\n</pre>\n\n<p>You
  can source the file and run the command <code>:PlugInstall</code> to install the
  plugin.</p>\n<h4>Using Lazy.nvim</h4>\n<p>If you are using Lazy.nvim as the plugin
  manager, you can add the plugin to the Configuration as below:</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kr\">return</span> <span class=\"p\">{</span>\n  <span class=\"p\">{</span>\n
  \   <span class=\"s2\">&quot;sourcegraph/sg.nvim&quot;</span><span class=\"p\">,</span>\n
  \   <span class=\"n\">dependencies</span> <span class=\"o\">=</span> <span class=\"p\">{</span>
  <span class=\"s2\">&quot;nvim-lua/plenary.nvim&quot;</span><span class=\"p\">,</span>
  <span class=\"s2\">&quot;nvim-telescope/telescope.nvim&quot;</span> <span class=\"p\">},</span>\n\n
  \   <span class=\"c1\">-- If you have a recent version of lazy.nvim, you don&#39;t
  need to add this!</span>\n    <span class=\"n\">build</span> <span class=\"o\">=</span>
  <span class=\"s2\">&quot;nvim -l build/init.lua&quot;</span><span class=\"p\">,</span>\n
  \ <span class=\"p\">},</span>\n<span class=\"p\">}</span>\n</pre></div>\n\n</pre>\n\n<p>You
  can source the file and run the command <code>:Lazy install</code> to install the
  plugin.</p>\n<p>For other installation instructions, you can refer to the <a href=\"https://github.com/sourcegraph/sg.nvim?tab=readme-ov-file#install\">README</a>
  of sg.nvim.</p>\n<h3>Installing Binaries and Building the Plugin</h3>\n<p>After
  the plugin is installed, you can move into the building and setup process of the
  sourcegraph Cody plugin.</p>\n<p>To install the binaries which are the dependencies
  of the plugin, you can run the command <code>:SourcegraphDownloadBinaries</code>
  which will force downloading the binaries, making sure that the plugin is properly
  installed.</p>\n<p><img src=\"https://meetgor-cdn.pages.dev/sg-nvim-build.png\"
  alt=\"SourcegraphDownloadBinaries Command Output\" /></p>\n<p>To build the plugin,
  you can simply run the command from within neovim as <code>:SourcegraphBuild</code>,
  this will rebuild the Sourcegraph rust crates and its dependencies (which might
  have failed during installation).</p>\n<h3>Sourcegraph Authentication</h3>\n<p>You
  need to now authenticate to your Sourcegraph account to use the sourcegraph features
  such as search and Cody.</p>\n<p>You can do that by running the command <code>:SourcegraphLogin</code>
  in neovim. This will require two inputs, the sourcegraph endpoint and the access
  token. If you are using sourcegraph cloud and not a self-hosted sourcegraph, you
  do not need to change the endpoint, just press enter and move ahead. This will redirect
  you to the browser for authentication and creating an access token. Log in with
  your credentials to sourcegraph and copy the access token.</p>\n<p>This will prompt
  you back to the neovim interface to provide the access token. Paste the access token
  there and you will be good to go.</p>\n<h3>Health Check</h3>\n<p>Once the plugins
  are installed then you can check the plugin is correctly downloaded by checking
  the health with the <code>:checkhealth sg</code> command.</p>\n<p>Below is the health
  check report on the sourcegraph plugin in neovim.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>sg:
  require(&quot;sg.health&quot;).check()\n\nsg.nvim report ~\n- Machine: x86_64, sysname:
  Linux\n- OK Valid nvim version: table: 0x7ffa0b7bce38\n- OK Found `cargo` (cargo
  1.70.0) is executable\n-     Use `:SourcegraphDownloadBinaries` to avoid building
  locally.\n- OK Found `sg-nvim-agent`: &quot;/home/meet/.local/share/nvim/site/pack/packer/start/sg.nvim/dist/sg-nvim-agent&quot;\n-
  OK Found `node` (config.node_executable) is executable.\n  Version: &#39;20.10.0&#39;\n-
  OK Found `cody-agent`: /home/meet/.local/share/nvim/site/pack/packer/start/sg.nvim/dist/cody-agent.js\n-
  OK   Authentication setup correctly\n- OK     endpoint set to: https://sourcegraph.com\n-
  OK Found correct binary versions: &quot;1.0.5&quot; = &quot;1.0.5&quot;\n- OK   Sourcegraph
  Connection info: {\n  access_token_set = true,\n  endpoint = &quot;https://sourcegraph.com&quot;,\n
  \ sg_nvim_version = &quot;1.0.5&quot;,\n  sourcegraph_version = {\n\t\t  build =
  &quot;256174_2023-12-30_5.2-dbb20677711c&quot;,\n\t\t  product = &quot;256174_2023-12-30_5.2-dbb20677711c&quot;\n
  \ }\n  }\n- To manage your Cody Account, navigate to: https://sourcegraph.com/cody/manage\n-
  OK Cody Account Information: {\n\t  chat_limit = 20,\n\t  chat_usage = 53,\n\t  code_limit
  = 500,\n\t  code_usage = 0,\n\t  cody_pro_enabled = false,\n\t  username = &quot;Mr-Destructive&quot;\n
  \ }\n- OK sg.nvim is ready to run\n</pre></div>\n\n</pre>\n\n<p>At this point, the
  sourcegraph plugin is ready to be used. However, we need to set up the plugin in
  neovim with the default configurations.</p>\n<h3>Configuration</h3>\n<p>In your
  lua setup files, you can set the plugin like this:</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">require</span><span class=\"p\">(</span><span class=\"s2\">&quot;sg&quot;</span><span
  class=\"p\">).</span><span class=\"n\">setup</span><span class=\"p\">()</span>\n</pre></div>\n\n</pre>\n\n<p>Source
  the lua file and restart neovim, this should properly make sourcegraph commands
  available in the editor. After these steps, Cody is right into neovim.</p>\n<h2>Usage</h2>\n<p>To
  use the plugin, there are multiple commands available within the editor, the complete
  list of them is given below:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>SourcegraphBuild\nSourcegraphClear\nSourcegraphDownloadBinaries\nSourcegraphInfo\nSourcegraphLink\nSourcegraphLogin\nSourcegraphSearch\n\nCodyAsk\nCodyChat\nCodyDo\nCodyRestart\nCodyTask\nCodyTaskAccept\nCodyTaskNext\nCodyTaskPrev\nCodyTaskView\nCodyToggle\n</pre></div>\n\n</pre>\n\n<p>You
  can get more info about these commands with the <code>help :commandname</code> command.
  The command are however self explanatory and simple to use.</p>\n<h3>Cody Ask</h3>\n<p>You
  can quickly parse a prompt as a string to the <code>:CodyAsk</code> command as <code>:CodyAsk
  &quot;what is neovim by the way?&quot;</code> and you will get Cody's response in
  the side vertical split.</p>\n<h3>Cody Chat</h3>\n<p>You can start Cody chats from
  neovim with the command <code>:CodyChat</code>. This will open a vertical split
  with the Cody chat interface, the bottom split has the user input prompt and the
  upper window will have the generated Cody response. You can enter the prompt in
  the bottom buffer get into normal mode and hit enter to send the prompt to generate
  a response from Cody.</p>\n<p><img src=\"https://meetgor-cdn.pages.dev/sg-nvim-cody-chat.png\"
  alt=\"Sourcegraph Cody Chat Interface\" /></p>\n<h2>Conclusion</h2>\n<p>Sourcegraph
  Cody is a great tool for getting quick solutions to trivial as well as specific
  problems to the current file/package or project. The context parsing of Cody makes
  it valuable for developers to ask for specific questions and it answers them really
  in a straightforward way, without the developer needing to parse the context for
  the prompt. Thank you for reading. Happy Coding :)</p>\n"
cover: ''
date: 2024-01-06
datetime: 2024-01-06 00:00:00+00:00
description: Installing and Setting up Sourcegraph Plugin and Cody into Neovim. Enabling
  code generations, chat and sourcegraph search feature in Neovim.
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2024-01-06-Neovim-Cody-Plugin.md
html: "<h2>Introduction</h2>\n<p>Have you ever used Sourcegraph's Cody? It is a great
  tool for developers, it is not just another LLM, it is tailored specifically for
  developers. Cody has some good features that allow parsing of context for the prompt
  in a smarter way.</p>\n<h3>What is Sourcegraph's Cody</h3>\n<p>Cody is an AI assistant
  for developers that understands code context and can generate code. It goes beyond
  just answering questions - it can write code for you.</p>\n<p>The major features
  of Cody from the rest of the LLMs or Chatbots are:</p>\n<ul>\n<li>\n<p>Cody understands
  your code context - it reads your open files, repositories, etc. So it can answer
  questions specifically about your codebase, not just general programming questions.</p>\n</li>\n<li>\n<p>Cody
  can explain sections of code to you in plain English. This helps ramp up on unfamiliar
  code bases.</p>\n</li>\n<li>\n<p>Cody integrates into popular editors like VS Code,
  IntelliJ, Neovim, and others for frictionless use while coding.</p>\n</li>\n</ul>\n<p>For
  more insights, check out the blog <a href=\"https://about.sourcegraph.com/blog/all-you-need-is-cody\">all
  you need is Cody</a>. This is a great article about what and how Cody is tailored
  specifically to assist developers.</p>\n<h2>Prerequisites</h2>\n<p>To setup sourcegraph
  on neovim, you will require the following setup:</p>\n<ul>\n<li>\n<p>Neovim 0.9
  or above</p>\n</li>\n<li>\n<p>Node.js &gt;= 18.17 (LTS)</p>\n</li>\n<li>\n<p>Cargo
  (Rust) (optional)</p>\n</li>\n</ul>\n<p>To install neovim latest/nightly release,
  you can follow the <a href=\"https://github.com/neovim/neovim/blob/master/INSTALL.md\">INSTALL</a>
  or <a href=\"https://github.com/neovim/neovim/blob/master/BUILD.md\">BUILD</a> documentation
  of the neovim project.</p>\n<p>Cargo is optional, as the plugin will install the
  binaries itself, however, if you prefer to have cargo, just install it in case something
  goes wrong.</p>\n<h2>Installing sg.nvim</h2>\n<p>There is a specific plugin for
  neovim for interacting with the Sourcegraph products, and Cody is one of them. The
  <a href=\"https://github.com/sourcegraph/sg.nvim\">sg.nvim</a> is a plugin for integrating
  sourcegraph search, Cody, and other features provided by soucegraph.</p>\n<h4>Using
  packer.nvim</h4>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"n\">use</span> <span class=\"p\">{</span> <span class=\"s1\">&#39;sourcegraph/sg.nvim&#39;</span><span
  class=\"p\">,</span> <span class=\"n\">run</span> <span class=\"o\">=</span> <span
  class=\"s1\">&#39;nvim -l build/init.lua&#39;</span> <span class=\"p\">}</span>\n<span
  class=\"n\">use</span> <span class=\"p\">{</span> <span class=\"s1\">&#39;nvim-lua/plenary.nvim&#39;</span>
  <span class=\"p\">}</span>\n</pre></div>\n\n</pre>\n\n<p>Source your lua file where
  you have configured all the plugins and then enter the command <code>:PackerInstall</code>
  or <code>:PackerSync</code> to install the plugin.</p>\n<h4>Using vim-plug</h4>\n<p>If
  you are using vim-plug as the plugin manager, you can add the plugin to the configuration
  as below:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button class='copy'
  title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>Plug
  &#39;sourcegraph/sg.nvim&#39;, { &#39;do&#39;: &#39;nvim -l build/init.lua&#39;
  }\n\n&quot; Required for various utilities\nPlug &#39;nvim-lua/plenary.nvim&#39;\n\n&quot;
  Required if you want to use some of the search functionality\nPlug &#39;nvim-telescope/telescope.nvim&#39;\n</pre></div>\n\n</pre>\n\n<p>You
  can source the file and run the command <code>:PlugInstall</code> to install the
  plugin.</p>\n<h4>Using Lazy.nvim</h4>\n<p>If you are using Lazy.nvim as the plugin
  manager, you can add the plugin to the Configuration as below:</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"kr\">return</span> <span class=\"p\">{</span>\n  <span class=\"p\">{</span>\n
  \   <span class=\"s2\">&quot;sourcegraph/sg.nvim&quot;</span><span class=\"p\">,</span>\n
  \   <span class=\"n\">dependencies</span> <span class=\"o\">=</span> <span class=\"p\">{</span>
  <span class=\"s2\">&quot;nvim-lua/plenary.nvim&quot;</span><span class=\"p\">,</span>
  <span class=\"s2\">&quot;nvim-telescope/telescope.nvim&quot;</span> <span class=\"p\">},</span>\n\n
  \   <span class=\"c1\">-- If you have a recent version of lazy.nvim, you don&#39;t
  need to add this!</span>\n    <span class=\"n\">build</span> <span class=\"o\">=</span>
  <span class=\"s2\">&quot;nvim -l build/init.lua&quot;</span><span class=\"p\">,</span>\n
  \ <span class=\"p\">},</span>\n<span class=\"p\">}</span>\n</pre></div>\n\n</pre>\n\n<p>You
  can source the file and run the command <code>:Lazy install</code> to install the
  plugin.</p>\n<p>For other installation instructions, you can refer to the <a href=\"https://github.com/sourcegraph/sg.nvim?tab=readme-ov-file#install\">README</a>
  of sg.nvim.</p>\n<h3>Installing Binaries and Building the Plugin</h3>\n<p>After
  the plugin is installed, you can move into the building and setup process of the
  sourcegraph Cody plugin.</p>\n<p>To install the binaries which are the dependencies
  of the plugin, you can run the command <code>:SourcegraphDownloadBinaries</code>
  which will force downloading the binaries, making sure that the plugin is properly
  installed.</p>\n<p><img src=\"https://meetgor-cdn.pages.dev/sg-nvim-build.png\"
  alt=\"SourcegraphDownloadBinaries Command Output\" /></p>\n<p>To build the plugin,
  you can simply run the command from within neovim as <code>:SourcegraphBuild</code>,
  this will rebuild the Sourcegraph rust crates and its dependencies (which might
  have failed during installation).</p>\n<h3>Sourcegraph Authentication</h3>\n<p>You
  need to now authenticate to your Sourcegraph account to use the sourcegraph features
  such as search and Cody.</p>\n<p>You can do that by running the command <code>:SourcegraphLogin</code>
  in neovim. This will require two inputs, the sourcegraph endpoint and the access
  token. If you are using sourcegraph cloud and not a self-hosted sourcegraph, you
  do not need to change the endpoint, just press enter and move ahead. This will redirect
  you to the browser for authentication and creating an access token. Log in with
  your credentials to sourcegraph and copy the access token.</p>\n<p>This will prompt
  you back to the neovim interface to provide the access token. Paste the access token
  there and you will be good to go.</p>\n<h3>Health Check</h3>\n<p>Once the plugins
  are installed then you can check the plugin is correctly downloaded by checking
  the health with the <code>:checkhealth sg</code> command.</p>\n<p>Below is the health
  check report on the sourcegraph plugin in neovim.</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>sg:
  require(&quot;sg.health&quot;).check()\n\nsg.nvim report ~\n- Machine: x86_64, sysname:
  Linux\n- OK Valid nvim version: table: 0x7ffa0b7bce38\n- OK Found `cargo` (cargo
  1.70.0) is executable\n-     Use `:SourcegraphDownloadBinaries` to avoid building
  locally.\n- OK Found `sg-nvim-agent`: &quot;/home/meet/.local/share/nvim/site/pack/packer/start/sg.nvim/dist/sg-nvim-agent&quot;\n-
  OK Found `node` (config.node_executable) is executable.\n  Version: &#39;20.10.0&#39;\n-
  OK Found `cody-agent`: /home/meet/.local/share/nvim/site/pack/packer/start/sg.nvim/dist/cody-agent.js\n-
  OK   Authentication setup correctly\n- OK     endpoint set to: https://sourcegraph.com\n-
  OK Found correct binary versions: &quot;1.0.5&quot; = &quot;1.0.5&quot;\n- OK   Sourcegraph
  Connection info: {\n  access_token_set = true,\n  endpoint = &quot;https://sourcegraph.com&quot;,\n
  \ sg_nvim_version = &quot;1.0.5&quot;,\n  sourcegraph_version = {\n\t\t  build =
  &quot;256174_2023-12-30_5.2-dbb20677711c&quot;,\n\t\t  product = &quot;256174_2023-12-30_5.2-dbb20677711c&quot;\n
  \ }\n  }\n- To manage your Cody Account, navigate to: https://sourcegraph.com/cody/manage\n-
  OK Cody Account Information: {\n\t  chat_limit = 20,\n\t  chat_usage = 53,\n\t  code_limit
  = 500,\n\t  code_usage = 0,\n\t  cody_pro_enabled = false,\n\t  username = &quot;Mr-Destructive&quot;\n
  \ }\n- OK sg.nvim is ready to run\n</pre></div>\n\n</pre>\n\n<p>At this point, the
  sourcegraph plugin is ready to be used. However, we need to set up the plugin in
  neovim with the default configurations.</p>\n<h3>Configuration</h3>\n<p>In your
  lua setup files, you can set the plugin like this:</p>\n<pre class='wrapper'>\n\n<div
  class='copy-wrapper'>\n\n<button class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>lua</div>\n<div class=\"highlight\"><pre><span></span><span
  class=\"nb\">require</span><span class=\"p\">(</span><span class=\"s2\">&quot;sg&quot;</span><span
  class=\"p\">).</span><span class=\"n\">setup</span><span class=\"p\">()</span>\n</pre></div>\n\n</pre>\n\n<p>Source
  the lua file and restart neovim, this should properly make sourcegraph commands
  available in the editor. After these steps, Cody is right into neovim.</p>\n<h2>Usage</h2>\n<p>To
  use the plugin, there are multiple commands available within the editor, the complete
  list of them is given below:</p>\n<pre class='wrapper'>\n\n<div class='copy-wrapper'>\n\n<button
  class='copy' title='copy code to clipboard' onclick=\"navigator.clipboard.writeText(this.parentElement.parentElement.querySelector('pre').textContent)\"><svg
  version=\"1.1\" id=\"Layer_1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\"
  x=\"0px\" y=\"0px\" viewBox=\"0 0 115.77 122.88\" style=\"enable-background:new
  0 0 115.77 122.88\" xml:space=\"preserve\"><style type=\"text/css\">.st0{fill-rule:evenodd;clip-rule:evenodd;}</style><g><path
  class=\"st0\" d=\"M89.62,13.96v7.73h12.19h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02v0.02
  v73.27v0.01h-0.02c-0.01,3.84-1.57,7.33-4.1,9.86c-2.51,2.5-5.98,4.06-9.82,4.07v0.02h-0.02h-61.7H40.1v-0.02
  c-3.84-0.01-7.34-1.57-9.86-4.1c-2.5-2.51-4.06-5.98-4.07-9.82h-0.02v-0.02V92.51H13.96h-0.01v-0.02c-3.84-0.01-7.34-1.57-9.86-4.1
  c-2.5-2.51-4.06-5.98-4.07-9.82H0v-0.02V13.96v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07V0h0.02h61.7
  h0.01v0.02c3.85,0.01,7.34,1.57,9.86,4.1c2.5,2.51,4.06,5.98,4.07,9.82h0.02V13.96L89.62,13.96z
  M79.04,21.69v-7.73v-0.02h0.02 c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v64.59v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h12.19V35.65
  v-0.01h0.02c0.01-3.85,1.58-7.34,4.1-9.86c2.51-2.5,5.98-4.06,9.82-4.07v-0.02h0.02H79.04L79.04,21.69z
  M105.18,108.92V35.65v-0.02 h0.02c0-0.91-0.39-1.75-1.01-2.37c-0.61-0.61-1.46-1-2.37-1v0.02h-0.01h-61.7h-0.02v-0.02c-0.91,0-1.75,0.39-2.37,1.01
  c-0.61,0.61-1,1.46-1,2.37h0.02v0.01v73.27v0.02h-0.02c0,0.91,0.39,1.75,1.01,2.37c0.61,0.61,1.46,1,2.37,1v-0.02h0.01h61.7h0.02
  v0.02c0.91,0,1.75-0.39,2.37-1.01c0.61-0.61,1-1.46,1-2.37h-0.02V108.92L105.18,108.92z\"/></g></svg></button>\n</div>\n
  \       \n<div class='language-bar'>plaintext</div>\n<div class=\"highlight\"><pre><span></span>SourcegraphBuild\nSourcegraphClear\nSourcegraphDownloadBinaries\nSourcegraphInfo\nSourcegraphLink\nSourcegraphLogin\nSourcegraphSearch\n\nCodyAsk\nCodyChat\nCodyDo\nCodyRestart\nCodyTask\nCodyTaskAccept\nCodyTaskNext\nCodyTaskPrev\nCodyTaskView\nCodyToggle\n</pre></div>\n\n</pre>\n\n<p>You
  can get more info about these commands with the <code>help :commandname</code> command.
  The command are however self explanatory and simple to use.</p>\n<h3>Cody Ask</h3>\n<p>You
  can quickly parse a prompt as a string to the <code>:CodyAsk</code> command as <code>:CodyAsk
  &quot;what is neovim by the way?&quot;</code> and you will get Cody's response in
  the side vertical split.</p>\n<h3>Cody Chat</h3>\n<p>You can start Cody chats from
  neovim with the command <code>:CodyChat</code>. This will open a vertical split
  with the Cody chat interface, the bottom split has the user input prompt and the
  upper window will have the generated Cody response. You can enter the prompt in
  the bottom buffer get into normal mode and hit enter to send the prompt to generate
  a response from Cody.</p>\n<p><img src=\"https://meetgor-cdn.pages.dev/sg-nvim-cody-chat.png\"
  alt=\"Sourcegraph Cody Chat Interface\" /></p>\n<h2>Conclusion</h2>\n<p>Sourcegraph
  Cody is a great tool for getting quick solutions to trivial as well as specific
  problems to the current file/package or project. The context parsing of Cody makes
  it valuable for developers to ask for specific questions and it answers them really
  in a straightforward way, without the developer needing to parse the context for
  the prompt. Thank you for reading. Happy Coding :)</p>\n"
image_url: https://meetgor-cdn.pages.dev/neovim-sourcegraph-cody.png
long_description: Have you ever used Sourcegraph Cody is an AI assistant for developers
  that understands code context and can generate code. It goes beyond just answering
  questions - it can write code for you. The major features of Cody from the rest
  of the LLMs or Ch
now: 2025-05-01 18:11:33.312874
path: blog/posts/2024-01-06-Neovim-Cody-Plugin.md
prevnext: null
slug: neovim-sourcegraph-cody
status: published
tags:
- vim
templateKey: blog-post
title: Neovim + Sourcegraph Cody Plugin Integration
today: 2025-05-01
---

## Introduction

Have you ever used Sourcegraph's Cody? It is a great tool for developers, it is not just another LLM, it is tailored specifically for developers. Cody has some good features that allow parsing of context for the prompt in a smarter way.

### What is Sourcegraph's Cody

Cody is an AI assistant for developers that understands code context and can generate code. It goes beyond just answering questions - it can write code for you.

The major features of Cody from the rest of the LLMs or Chatbots are:

* Cody understands your code context - it reads your open files, repositories, etc. So it can answer questions specifically about your codebase, not just general programming questions.

* Cody can explain sections of code to you in plain English. This helps ramp up on unfamiliar code bases.

* Cody integrates into popular editors like VS Code, IntelliJ, Neovim, and others for frictionless use while coding.

For more insights, check out the blog [all you need is Cody](https://about.sourcegraph.com/blog/all-you-need-is-cody). This is a great article about what and how Cody is tailored specifically to assist developers.

## Prerequisites

To setup sourcegraph on neovim, you will require the following setup:

* Neovim 0.9 or above

* Node.js &gt;= 18.17 (LTS)

* Cargo (Rust) (optional)

To install neovim latest/nightly release, you can follow the [INSTALL](https://github.com/neovim/neovim/blob/master/INSTALL.md) or [BUILD](https://github.com/neovim/neovim/blob/master/BUILD.md) documentation of the neovim project.

Cargo is optional, as the plugin will install the binaries itself, however, if you prefer to have cargo, just install it in case something goes wrong.

## Installing sg.nvim

There is a specific plugin for neovim for interacting with the Sourcegraph products, and Cody is one of them. The [sg.nvim](https://github.com/sourcegraph/sg.nvim) is a plugin for integrating sourcegraph search, Cody, and other features provided by soucegraph.

#### Using packer.nvim

```lua
use { 'sourcegraph/sg.nvim', run = 'nvim -l build/init.lua' }
use { 'nvim-lua/plenary.nvim' }
```

Source your lua file where you have configured all the plugins and then enter the command `:PackerInstall` or `:PackerSync` to install the plugin.

#### Using vim-plug

If you are using vim-plug as the plugin manager, you can add the plugin to the configuration as below:

```plaintext
Plug 'sourcegraph/sg.nvim', { 'do': 'nvim -l build/init.lua' }

" Required for various utilities
Plug 'nvim-lua/plenary.nvim'

" Required if you want to use some of the search functionality
Plug 'nvim-telescope/telescope.nvim'
```

You can source the file and run the command `:PlugInstall` to install the plugin.

#### Using Lazy.nvim

If you are using Lazy.nvim as the plugin manager, you can add the plugin to the Configuration as below:

```lua
return {
  {
    "sourcegraph/sg.nvim",
    dependencies = { "nvim-lua/plenary.nvim", "nvim-telescope/telescope.nvim" },

    -- If you have a recent version of lazy.nvim, you don't need to add this!
    build = "nvim -l build/init.lua",
  },
}
```

You can source the file and run the command `:Lazy install` to install the plugin.

For other installation instructions, you can refer to the [README](https://github.com/sourcegraph/sg.nvim?tab=readme-ov-file#install) of sg.nvim.

### Installing Binaries and Building the Plugin

After the plugin is installed, you can move into the building and setup process of the sourcegraph Cody plugin.

To install the binaries which are the dependencies of the plugin, you can run the command `:SourcegraphDownloadBinaries` which will force downloading the binaries, making sure that the plugin is properly installed.

![SourcegraphDownloadBinaries Command Output](https://meetgor-cdn.pages.dev/sg-nvim-build.png)

To build the plugin, you can simply run the command from within neovim as `:SourcegraphBuild`, this will rebuild the Sourcegraph rust crates and its dependencies (which might have failed during installation).

### Sourcegraph Authentication

You need to now authenticate to your Sourcegraph account to use the sourcegraph features such as search and Cody.

You can do that by running the command `:SourcegraphLogin` in neovim. This will require two inputs, the sourcegraph endpoint and the access token. If you are using sourcegraph cloud and not a self-hosted sourcegraph, you do not need to change the endpoint, just press enter and move ahead. This will redirect you to the browser for authentication and creating an access token. Log in with your credentials to sourcegraph and copy the access token.

This will prompt you back to the neovim interface to provide the access token. Paste the access token there and you will be good to go.

### Health Check

Once the plugins are installed then you can check the plugin is correctly downloaded by checking the health with the `:checkhealth sg` command.

Below is the health check report on the sourcegraph plugin in neovim.

```plaintext
sg: require("sg.health").check()

sg.nvim report ~
- Machine: x86_64, sysname: Linux
- OK Valid nvim version: table: 0x7ffa0b7bce38
- OK Found `cargo` (cargo 1.70.0) is executable
-     Use `:SourcegraphDownloadBinaries` to avoid building locally.
- OK Found `sg-nvim-agent`: "/home/meet/.local/share/nvim/site/pack/packer/start/sg.nvim/dist/sg-nvim-agent"
- OK Found `node` (config.node_executable) is executable.
  Version: '20.10.0'
- OK Found `cody-agent`: /home/meet/.local/share/nvim/site/pack/packer/start/sg.nvim/dist/cody-agent.js
- OK   Authentication setup correctly
- OK     endpoint set to: https://sourcegraph.com
- OK Found correct binary versions: "1.0.5" = "1.0.5"
- OK   Sourcegraph Connection info: {
  access_token_set = true,
  endpoint = "https://sourcegraph.com",
  sg_nvim_version = "1.0.5",
  sourcegraph_version = {
		  build = "256174_2023-12-30_5.2-dbb20677711c",
		  product = "256174_2023-12-30_5.2-dbb20677711c"
  }
  }
- To manage your Cody Account, navigate to: https://sourcegraph.com/cody/manage
- OK Cody Account Information: {
	  chat_limit = 20,
	  chat_usage = 53,
	  code_limit = 500,
	  code_usage = 0,
	  cody_pro_enabled = false,
	  username = "Mr-Destructive"
  }
- OK sg.nvim is ready to run
```

At this point, the sourcegraph plugin is ready to be used. However, we need to set up the plugin in neovim with the default configurations.

### Configuration

In your lua setup files, you can set the plugin like this:

```lua
require("sg").setup()
```

Source the lua file and restart neovim, this should properly make sourcegraph commands available in the editor. After these steps, Cody is right into neovim.

## Usage

To use the plugin, there are multiple commands available within the editor, the complete list of them is given below:

```plaintext
SourcegraphBuild
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
```

You can get more info about these commands with the `help :commandname` command. The command are however self explanatory and simple to use.

### Cody Ask

You can quickly parse a prompt as a string to the `:CodyAsk` command as `:CodyAsk "what is neovim by the way?"` and you will get Cody's response in the side vertical split.

### Cody Chat

You can start Cody chats from neovim with the command `:CodyChat`. This will open a vertical split with the Cody chat interface, the bottom split has the user input prompt and the upper window will have the generated Cody response. You can enter the prompt in the bottom buffer get into normal mode and hit enter to send the prompt to generate a response from Cody.

![Sourcegraph Cody Chat Interface](https://meetgor-cdn.pages.dev/sg-nvim-cody-chat.png)

## Conclusion

Sourcegraph Cody is a great tool for getting quick solutions to trivial as well as specific problems to the current file/package or project. The context parsing of Cody makes it valuable for developers to ask for specific questions and it answers them really in a straightforward way, without the developer needing to parse the context for the prompt. Thank you for reading. Happy Coding :)