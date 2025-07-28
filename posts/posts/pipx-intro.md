{
  "title": "Pipx: A python package consumption tool for CLI packages",
  "status": "published",
  "slug": "pipx-intro",
  "date": "2025-04-05T12:33:46Z"
}

<h2>Introduction</h2>
<p>Previously, I explored the <a href="https://mr-destructive.github.io/techstructive-blog/pipenv-intro/">pipenv</a> as a python package management tool and this week it is time for exploring a python CLI package isolation tool called <code>pipx</code>. It is a really simple and powerful tool for running commands as an executable from an isolated environment. Those commands should be associated with a python package that has CLI. In pipx, we install the package once and we can use the package anywhere on our system isolated from other virtual environments.</p>
<p>In this article, We will explore the pipx package/tool in python. From installation to setting up your environment and removing some friction in your python workflow. This won't be an in-depth guide of pipx, though we would cover a few python packages that can be easily used in your environment.</p>
<h2>What is Pipx</h2>
<p>Pipx is a python package that works similar to <code>pip</code> but it is quite specific to the usage. It can run CLI commands of any python package from anywhere on your system. It uses its own virtual environment for managing the packages and the python version.</p>
<p><strong>NOTE</strong></p>
<ul>
<li>Pipx is a tool to install and execute CLI commands provided in the python packages</li>
<li>Pipx is not a python package management tool like pip or pipenv.</li>
</ul>
<p>Pipx is similar to <code>npx</code> for nodejs. Though the major difference is that pipx runs and installs all packages/tools globally unlike npx.</p>
<h2>Install Pipx</h2>
<p><strong>Pipx requires Python 3.6 or above.</strong></p>
<ul>
<li>You can install it normally like a python package with pip:</li>
</ul>
<pre><code>pip install pipx
</code></pre>
<p>To find out if pipx was successfully installed, you can simply type in the command:</p>
<pre><code>pipx --version
</code></pre>
<p>If it gives a number and no errors, it has been successfully installed in your system globally.</p>
<h2>Pipx Commands</h2>
<p>The primary commands that pipx can use are:</p>
<ol>
<li>pipx install</li>
<li>pipx list</li>
<li>pipx run</li>
<li>pipx inject</li>
<li>pipx uninstall</li>
</ol>
<p>For further command options for a specific use case, you can use the <code>pipx -h</code> command to get the detailed list of commands and options available in pipx.</p>
<p>Also, the documentation of pipx is really incredibly beginner-friendly. You can follow up with the article from the <a href="https://pypa.github.io/pipx/">documentation</a> webpage or the <a href="https://github.com/pypa/pipx/">GitHub</a> repository. The documentation and the webpage are simple single-page websites that really make learning the tool easier and it also has a well-documented package.</p>
<h3>Pipx Install Command</h3>
<p>Pipx can install packages in a virtual environment specific to the package provided. This means that you will have different virtual envs for each package you install and not a single virtual env for a project, it creates and makes the package venvs accessible globally. The next step after installing the package in a different venv is to make that package commands accessible via the PATH environment variable.</p>
<pre><code>pipx install &lt;package_name&gt;
</code></pre>
<p>If say we install <a href="https://pypi.org/project/pgcli/">pgcli</a>, a package for interacting with postgres database using python in the Command line. We will simply enter the following commands:</p>
<pre><code>pipx install pgcli
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647783261/blog-media/du830o3fbqogvkaesxnq.png" alt="install package with pipx"></p>
<p>We can use the package pgcli globally in our system. So, we have access to the databases in the local Postgres server. We can run the commands as normally as we want using the pgcli by just prefixing with <code>pipx run</code>.</p>
<pre><code>pipx run pgcli &lt;local_database_name&gt; -U postgres
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647782221/blog-media/dfpwlxkh5ybqj5pdos8d.gif" alt="PGCLI demonstration GIF"></p>
<p>This is a really cool way to interact with CLI applications without installing separately or globally, using pipx the virtual environment for the specific package makes it a clean and persistent behavior for running the commands anywhere in the system.</p>
<p>If you want to know more about pgcli tool/package in Python, you can it's <a href="https://www.pgcli.com/docs">documentation</a> webpage or the <a href="https://github.com/dbcli/pgcli">GitHub</a> repository.</p>
<h4>Pipx Default PATH</h4>
<p>The default path at which pipx stores the virtual environments for each package is by default the <code>~/.local/pipx/venvs</code>(here ~ refers to the root/default directory), the documentation says that it can be changed by editing the environment variable <code>PIPX_HOME</code>.</p>
<p>The default path at which pipx stores the binaries/command apps for each package is stored in the path <code>~/.local/bin</code>, even this path can be modified by editing the environment variable <code>PIPX_BIN_DIR</code></p>
<h3>Pipx List Command</h3>
<p>Pipx list command simply lists the executables or the commands you can run with pipx. This command will display the commands associated with the packages that are installed in the pipx environment.</p>
<pre><code>pipx list
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647782349/blog-media/qu95nynklbuceqqd9qke.png" alt="Pipx list output"></p>
<p>This command gives a detailed list of the commands associated with their respective packages. It also gives the meta-information like the Python version in which the package was installed.</p>
<h3>Pipx Run Command</h3>
<p>The most useful command in pipx has to be <code>pipx run</code>, this command can execute provide package along with the specified command associated with the package.</p>
<p>The structure of the command is as follows:</p>
<pre><code>pipx run &lt;package_name&gt; &lt;command&gt;
</code></pre>
<p>The pipx documentation refers to the <code>command</code> as an <code>APP</code>, as it is an executable created from the package specification.</p>
<p>If the APP/command name is not matching the <code>package_name</code>(most of the time it won't), you need to add an argument to the run command.</p>
<pre><code>pipx run --spec &lt;package_name&gt; &lt;command&gt;
</code></pre>
<p>The <code>--spec</code> option allows specifying a certain package to be used while running the command/APP. Using this option with the <code>pipx run</code> command, we can run package-specific commands. For example, if we want to run an ipython from the terminal, without messing up the current environment i.e. without installing any package. You can do that by using pipx.</p>
<pre><code>pipx install dailydotdev-bookmark-cli

pipx run --spec dailydotdev-bookmark-cli bookamrks
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647782547/blog-media/vsfmwlzazqtosuwttexb.gif" alt="dailydotdev-bookmark-cli demo"></p>
<p>This is my shameless plug of installing my <a href="https://github.com/Mr-Destructive/bookmarks-cli">first python package</a>. The command used for the CLI is totally not related to the package name and hence we have to use the <code>--spec</code> option.</p>
<h2>Install and Run packages from Pipx</h2>
<p>Any python package which provide a Command Line Interface can be installed and run with pipx. Some of the packages like <a href="https://docs.pytest.org/en/7.1.x/contents.html">pytest</a>, <a href="https://pipenv-fork.readthedocs.io/en/latest/install.html">pipenv</a>, <a href="https://github.com/Textualize/rich-cli">rich-cli</a>, <a href="https://github.com/psf/black">Black code formatter</a>, <a href="https://markata.dev/">markata</a> and many others.
You can install any of the packages which do have a cli to interact with on the go with pipx.</p>
<h3>Running Ipython shell</h3>
<p>We can use <a href="https://pypi.org/project/ipython/">Ipython</a> shell from pipx as an isolated environment. We simply install the package first, after installing the package it creates the virtual environment. Creating a separate virtual environment registers the app/command binaries into the PATH environment of the system so that they can be accessed globally.</p>
<pre><code>pipx install ipython

pipx run ipython
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647782651/blog-media/ssgymybn0dwi8ocs6xpf.png" alt="IPython pipx demo"></p>
<p>For further documentation on ipython using the CLI, you can refer to the <a href="https://github.com/ipython/ipython">GitHub</a> link or the <a href="https://ipython.org/ipython-doc/3/interactive/reference.html#command-line-usage">documentation</a> page.</p>
<h3>Reading an IPYNB file from pipx with JUT</h3>
<p>If you just want to read an ipynb file from the terminal, you can do that using <a href="https://pypi.org/project/jut/">jut</a> by specifying the command:</p>
<pre><code>pipx install jut

pipx run jut &lt;notebook.ipynb&gt;
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647782686/blog-media/ddm5uiqarjjmedhqvcsk.png" alt="jut pipx demo"></p>
<p>You can find more information on JUT via their official <a href="https://github.com/kracekumar/jut">GitHub</a> repository.</p>
<h3>Using rich-cli to display text using pipx</h3>
<p>We can even use <a href="https://pypi.org/project/rich-cli/">Rich-CLI</a> to print rich content in the project. Rich-CLI as the name suggests is a CLI for the Rich package in python. It is really intuitive and simple to use. If we want to take snaps of the rich content of a markdown file, source file, dataset, etc. rich-cli is a tool that can quickly do that. Use pipx to install the package globally and simply run wherever required in an isolated environment.</p>
<pre><code>pipx install rich-cli

pipx run --spec rich-cli rich &quot;[blue]Hello[/], [yellow]World[/]!&quot; --print
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647782738/blog-media/xvdyhxjpj2hsghl2x1ng.png" alt="rich-cli pipx demo"></p>
<p>This will give you a colorful display of &quot;Hello, World&quot; without you messing up your current project. Pipx has an isolated environment so it doesn't tinker with your local virtual environment or project. Though you can use your current project to use those CLI to execute commands which might require the source file in the project.</p>
<p>We can even print the markdown file in a rich format, simply using pipx and rich-cli. There are a lot of things we can do with rich-cli.</p>
<pre><code>pipx run rich-cli &lt;markdown_file.md&gt;
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647785474/blog-media/bwmvucrcgrtgwdv58ffj.png" alt="rich-cli markdown demo"></p>
<p>For further reading on the rich-cli package, you visit their <a href="https://github.com/Textualize/rich-cli">GitHub</a> link and the <a href="https://github.com/Textualize/rich-cli/blob/main/README.md">Documentation</a> Readme.</p>
<h3>Using pytest to perform a test for the current project</h3>
<p><a href="https://pypi.org/project/pytest/">Pytest</a> allows writing simple and scalable tests for Python apps, libraries, and packages. We can use it to write tests in the project without adding it as a dependency in the python environment. Simply install the package with pipx and run the tests in an isolated environment.</p>
<pre><code>pipx install pytest

pipx run pytest temp.py
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647782771/blog-media/qutb0bxzlnwctl9mbuon.png" alt="Pytest pipx demo"></p>
<p>Here, we can see the pytest was performed on the <code>temp.py</code> file which can be any application file for your project. This gave us results without installing pytest in our application's virtual environment. This becomes really convenient to run certain package commands whose package should not be a dependency on your project in the virtual environment.</p>
<p>You can visit Pytest's <a href="https://github.com/pytest-dev/pytest/">GitHub</a>, <a href="https://docs.pytest.org/en/stable/">Documentation</a> and their <a href="https://pytest.org/">Home Page</a>.</p>
<h3>Using Black (Python code formatter)</h3>
<p>We can use pipx to even format the python source files using <a href="https://pypi.org/project/black/">black</a>. We don't have to install <code>black</code> as a dependency in the project. Pipx will simply install black in its isolated environment and run the specified command.</p>
<pre><code>pipx install black

pipx run black temp.py
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647782823/blog-media/dtnae85fvgae2y4aiqkg.png" alt="Black pipx demo"></p>
<p>Black is a code formatter in Python, it basically edits your source python files for any incorrect python semantics and syntax, it corrects/formats them without a compromise. You can visit the <a href="https://github.com/psf/black">GitHub</a>, <a href="https://black.readthedocs.io/en/stable">Documentation</a> or the Black package <a href="https://black.vercel.app/">Playground webpage</a>.</p>
<h3>Using httpie (curl equivalent in Python)</h3>
<p>We can use the <a href="https://pypi.org/project/httpie/">httpie</a> package in python which is a CLI tool for HTTP clients. It is very similar to the <code>curl</code> command in Linux. We can even use that with pipx to test out API endpoints or any website that you might be working with without leaving the terminal. For further details on the httpie package, you can visit their <a href="https://httpie.io/">Home page</a>, <a href="https://github.com/httpie">GitHub</a>, or <a href="https://pypi.org/project/httpie/">PyPI webpage</a>.</p>
<pre><code>pipx install httpie

pipx run --spec httpie http -v dev.to
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647782857/blog-media/jbaudm3pbnnsjtzy5ok7.png" alt="httpie pipx demo"></p>
<p>Here, we can see the package simply gives a simple output of the headers of the request. It is smart enough to understand the web link as <code>https://dev.to</code> unlike CURL which needs to have an exact match.</p>
<h3>Running Markata blog with pipx</h3>
<p>As said, every package that bundles itself with a CLI can be installed and run from anywhere. So, <a href="https://pypi.org/project/markata/">markata</a> which is a plugin-driven static site generator which is a python package also bundles with a CLI. Using the base app(markata) we can run its associated commands like <code>build</code>, <code>list</code>, etc. I use this to build my blog.</p>
<pre><code>pipx install markata

pipx run markata build
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647783021/blog-media/cygzwixyzwv4wwsccbfc.gif" alt="Markata pipx demo"></p>
<p>You can visit the Markata's <a href="https://markata.dev/">documentation</a> web page or the <a href="https://github.com/WaylonWalker/markata">GitHub</a> repository for further insights on the Static site generator.</p>
<p><strong>If you want to explore more packages that you can use with pipx, then you can look at the docs guide <a href="https://pypa.github.io/pipx/programs-to-try/">programs-to-try</a> section.</strong></p>
<h2>Conclusion</h2>
<p>So, from this simple introduction, we were able to understand the basics of pipx which is a python package for isolating and running CLI-specific package commands. We saw how to setup pipx, install packages, run the commands from the package and interact with the virtual environment that it creates for each package we install.</p>
<p>Thank you for reading, if you have any comments, suggestions, feedback please let me know in the comments. Happy Coding :)</p>
