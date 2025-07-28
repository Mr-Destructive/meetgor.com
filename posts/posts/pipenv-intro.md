{
  "title": "Pipenv: Python's Official Package Management tool",
  "status": "published",
  "slug": "pipenv-intro",
  "date": "2025-04-05T12:33:46Z"
}

<h2>Introduction</h2>
<p>Pipenv is Python's officially recommended Package management tool, as the name suggests it combines the power of <code>pip</code> and <code>virtualenv</code>. It is really powerful and gives control of package management for a python project. It also has support for adding dev packages, python versions, smartly recognizing main and derived package installation/uninstallation, and so on. It is the <a href="https://packaging.python.org/en/latest/tutorials/managing-dependencies/#managing-dependencies">official package management tool for Python</a>.</p>
<p>It is quite similar to npm for Nodejs in Javascript, or bundle for Ruby, cargo for Rust, and so on. It really simple and easy to use as it manages the overhead for package management for us and hence it is also a high-level package management tool as opposed to pip, which is not as powerful as Pipenv. So, in this article, we'll explore <a href="https://pypi.org/project/pipenv/">Pipenv</a> package manager for Python and how you can use it in your next python project. Let's get started.</p>
<h2>Install Pipenv</h2>
<p><a href="https://pypi.org/project/pipenv/">Pipenv</a> is just like any other package in python, you can install it with pip as normally you install any other package with the command:</p>
<pre><code>pip install pipenv
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647193069/blog-media/hnhoeigfhx2hsypexgm5.png" alt="Pipenv Install"></p>
<p>You can refer to the documentation of Pipenv from <a href="https://pipenv.pypa.io/en/latest/">here</a>.</p>
<h2>Set up a Python Environment</h2>
<p>This step is not mandatory but it avoids any mistakes that you can make in the future while installing the package, so to simply create a new python environment for your project, you simply have to write the following command:</p>
<pre><code>pipenv shell
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647192853/blog-media/dahwaqnblvvvqyyw62uq.png" alt="Pipenv initialize">
This will create the virtual environment for Python in the local folder. It is quite similar to installing <code>virtualenv</code> and then activating the env/venv folder script. Though pipenv is quite powerful as it automatically detects the environment.</p>
<p>If you look at it carefully, there will be a file generated after the command has been executed successfully. The file called <code>Pipfile</code> without any extension will have been created in your current folder where you executed the command from. The file contains the version of python used in this project along with the list of dependencies(currently empty). Also the source from where it will download and manage the dependencies.</p>
<p>The Pipfile after simply creating the virtualenv via the command <code>pipenv shell</code> looks something like follows:</p>
<pre><code>[[source]]
url = &quot;https://pypi.org/simple&quot;
verify_ssl = true
name = &quot;pypi&quot;

[packages]

[dev-packages]

[requires]
python_version = &quot;3.8&quot;
</code></pre>
<p>If you want to set up a specific version of python for the virtual environment, you can do this using pipenv. The version should be installed in your system though for the pipenv to pick it up, if you have the specific version of python set up with all path variables configured, you can use the below command to set up the version of python in pipenv.</p>
<pre><code>pipenv --python version-number
</code></pre>
<p>For example : <code>pipenv --python 3.9</code> will set up the virtual environment with python version <code>3.9</code>.</p>
<h2>What is Pipfile</h2>
<p>Pipfile is basically a TOML file that has all the details of the different dependencies/packages and the version of Python used in the project/directory. A TOML is a simple configuration file that is reader-friendly, it is a map of keys and values as configuration data structures. In Pipenv, we can have keys like <code>package-names</code> and the value as the <code>version-number</code>, certain groups of dependencies like <code>dev-packages</code> and <code>packages</code>, etc. Pipenv is the file that Pipenv implements its package management environment. The file is really important and powerful as it can install all dependencies even if provided with <code>requirements.txt</code> file. Yes, it can automatically detect that if you provide the path to that file.</p>
<p>Pipenv also has additional features like adding dev dependencies/packages in the project as a separate dependency. So if you want to test a feature with a specific package you can add it as a dev package and it will be stored separately. The pipfile will segregate the dependencies so that Pipenv can install/uninstall from the virtual environment. In short, Pipfile lets you have great control over your project's packages management.</p>
<h2>Installing Python Packages</h2>
<p>Once your Pipenv is initialized as a virtual environment for Python, we can install dependencies with either <code>pip</code> or <code>pipenv</code>. This is the mistake that might get you trapped, if you already have not run the command <code>pipenv shell</code> and installed any dependencies with <code>pip</code>, you will install that dependency globally and make some mess of the project. So, it is advised to use <code>pipenv shell</code> in order to activate your virtual environment. If you do not wanna do that, you can use the command to install any dependency without activating the virtual environment.</p>
<pre><code>pipenv install &lt;package-name&gt;
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647192980/blog-media/dfsokm6e1easwvxtgehh.png" alt="Pipenv Install Flask"></p>
<p>This will detect any virtual environment in the project, if it's not created already it will create it and install that package in that environment.</p>
<p>Installing any package using <code>pipenv</code> will update the Pipenv file and the package to its packages list.</p>
<pre><code class="language-toml">[[source]]
url = &quot;https://pypi.org/simple&quot;
verify_ssl = true
name = &quot;pypi&quot;

[packages]
flask = &quot;*&quot;

[dev-packages]

[requires]
python_version = &quot;3.8&quot;
</code></pre>
<p>OR</p>
<p>If you wish to install with pip, as usual, you need to be in the virtual subshell. If you already are, then Pipenv will add that dependency to the virtual environment. Note though, if you install any package with <code>pip</code> and not with <code>pipenv</code>, the package won't be added to Pipfile but would be installed in the virtual environment.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647193149/blog-media/suaykqsyrgn1w0mou7f0.png" alt="Pipenv pip install"></p>
<p>Here, if we type the command <code>pipenv graph</code> it will show us a nice detailed list of all the installed dependencies.</p>
<pre><code>pipenv graph
</code></pre>
<pre><code class="language-shell">$ pipenv graph
Flask==2.0.3
  - click [required: &gt;=7.1.2, installed: 8.0.4]
    - colorama [required: Any, installed: 0.4.4]
  - itsdangerous [required: &gt;=2.0, installed: 2.1.1]
  - Jinja2 [required: &gt;=3.0, installed: 3.0.3]
    - MarkupSafe [required: &gt;=2.0, installed: 2.1.0]
  - Werkzeug [required: &gt;=2.0, installed: 2.0.3]
requests==2.27.1
  - certifi [required: &gt;=2017.4.17, installed: 2021.10.8]
  - charset-normalizer [required: ~=2.0.0, installed: 2.0.12]
  - idna [required: &gt;=2.5,&lt;4, installed: 3.3]
  - urllib3 [required: &gt;=1.21.1,&lt;1.27, installed: 1.26.8]
</code></pre>
<p>As you might have guessed, the above command is equivalent to the <code>pip freeze</code> command, but just compare the details both tools have. Pipenv really shines here.</p>
<p>If you compare the output of Pipfile and <code>pipenv graph</code> you get a bit confused as to why is there are so fewer packages in the file. So, <strong>Pipfile doesn't store the sub-packages/dependencies of a certain base package</strong>. Let's take, for example, Flask here, we have Flask as the main package, and <code>click</code>, <code>Jinja2</code>, <code>Werkzeug</code> are its sub dependencies, further <code>colorama</code> and <code>MarkupSafe</code> are in turn dependencies of <code>click</code> and <code>Jinja2</code>. So, Pipfile just includes the top-level packages, pipenv is smart enough to install the rest. It also doesn't include <code>requests</code> but it is indeed included in the <code>graph</code>. This is because Pipenv will only put the dependencies in the Pipfile if it has been installed via the <code>pipenv</code> command.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647193333/blog-media/gzo95pbquaioujtqvntu.png" alt="Pipenv Graph vs pip freeze"></p>
<p>You can use <code>pip freeze</code> here as well as we are technically in a Python virtual environment. And you can clearly see the difference. Pipenv is a high-level tool compared to pip, it technically uses <code>pip</code> and <code>virtualenv</code> and leverages both of them to become a smart package management tool.</p>
<h2>What is the Pipfile.lock</h2>
<ul>
<li>If you are coming from <code>Javascript</code> world, it is similar to <code>package-lock.json</code> file.</li>
<li>If you are coming from <code>Ruby</code> world, it is similar to the <code>Gemfile.lock</code> file.</li>
<li>If you are coming from <code>Rust</code> world, it is similar to the <code>cargo.lock</code> file.</li>
</ul>
<p>Ok, you get the idea, it is a file that more robustly specifies the version of the packages without conflicting with the other version or the Python version itself. If you look at the Pipfile.lock also has hashes that store the sub-packages as well. The file format here is JSON as opposed to TOML for the Pipfile.</p>
<h2>Configuring the Pipenv environment</h2>
<p>Now, a question you might have is where is the virtual environment? Is it there? Of course, it will be there, it is configured to a different location though. By default, it will be stored in the <code>~\.virtualenvs\</code> folder.</p>
<p>You can get the location of the current virtualenv with the following command:</p>
<pre><code>pipenv --venv
</code></pre>
<pre><code>$ pipenv --venv
C:\Userscer\.virtualenvs\pipenv-blog-gHY6vF9t
</code></pre>
<p>For Windows, it is in the Admin user(in my case it is named <code>acer</code>) followed by the hidden folder <code>virtualenvs</code>, this folder will contain all the virtualenvs for different projects using <code>pipenv</code>.</p>
<p>If you wished to change this location and keep the virtual environment folder in the same directory as your project, you can set up the environment variable for it as follows:</p>
<p>For Linux/macOS:
Add the following to your <code>~/.bashrc</code> or other equivalent shell configuration file.</p>
<pre><code>export PIPENV_VENV_IN_PROJECT=&quot;enabled&quot;
</code></pre>
<p>For Windows:</p>
<p>Add it to the PATH Environment variable.</p>
<pre><code>set PIPENV_VENV_IN_PROJECT=&quot;enabled&quot;   
</code></pre>
<p>This will make sure the virtualenvs for the project in <code>pipenv</code> are created inside the current folder itself and not in a single <code>~\.virtualenvs\</code> folder.</p>
<h2>Creating the requirements.txt file</h2>
<p>Let's say you also want to create a requirements.txt file for distributing your project to someone else, as not everyone will use Pipenv to manage their dependencies. It is really straightforward and quick to create a requirements.txt file from the Pipenv environment.</p>
<p>The below command will make the <code>requirements.txt</code> file from the existing Pipenv project.</p>
<pre><code>pipenv lock -r &gt;requirements.txt
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647193388/blog-media/e6t68b7ckrsrvadvmeqa.png" alt="Pipenv to requirements.txt file"><br>
This will create the requirements.txt file, based on the Pipfile. Pipenv is smart again to provide all of the required dependencies to the requirements.txt in order that pip will be able to install all the required dependencies.</p>
<h2>Using requirements.txt in Pipenv</h2>
<p>We can install all the dependencies from the requirements.txt file while we are migrating from bare-bones virtualenv and pip to Pipenv. Pipenv will install all the mentioned dependencies and it will also add its checks for the appropriate checks for dependencies.</p>
<pre><code>pipenv install -r requirements.txt
</code></pre>
<p>This will install the dependencies mentioned in the requirements.txt file into the Pipenv virtual environment.</p>
<h2>Managing Dev Packages</h2>
<p>Let's say we want to install a package but we are not sure to use it in production or the actual script, just a test for how it will work. Thus we have dev packages to install optional or testing packages.</p>
<p>To install a dev-dependency or package, you can install via the following command:</p>
<pre><code>pipenv install &lt;package-name&gt; --dev
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647194653/blog-media/x5dimgfd2ikm2ercbzhv.png" alt="Pipenv dev package install"></p>
<p>If we see the Pipfile, the <code>django</code> package that we installed tagged with <code>--dev</code> will be in the <code>dev-packages</code></p>
<pre><code>$ cat Pipfile
[[source]]
url = &quot;https://pypi.org/simple&quot;
verify_ssl = true
name = &quot;pypi&quot;

[packages]
flask = &quot;*&quot;

[dev-packages]
django = &quot;*&quot;

[requires]
python_version = &quot;3.8&quot;
</code></pre>
<p>If we wanted to uninstall all the dev-packages, we can simply enter the command:</p>
<pre><code>pipenv uninstall --all-dev
</code></pre>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1647261079/blog-media/bg9n7aj6rfxsvdwflnah.gif" alt="Pipenv uninstall devpackages"></p>
<p>This will simply uninstall all the dev dependencies/packages from the pipenv environment and also remove them from the packages list in Pipfile. If you wished to uninstall a specific package in pipenv, you can uninstall it by the simple command:</p>
<pre><code>pipenv uninstall &lt;package-name&gt;
</code></pre>
<h2>Installing/Uninstalling all packages</h2>
<p>To install only the default packages and not dev-packages.</p>
<pre><code>pipenv install 
</code></pre>
<p>To install or configure a project, if you want to test the project with all the dev dependencies, you can install them with:</p>
<pre><code>pipenv install --dev
</code></pre>
<p>This will install all the packages both <code>packages</code> and <code>dev-packages</code>.</p>
<p>If you want to uninstall all the packages in pipenv, you can use the command :</p>
<pre><code>pipenv uninstall --all
</code></pre>
<p>This will uninstall all the default and dev packages from pipenv. This is like starting a fresh virtual environment.</p>
<h3>References:</h3>
<ul>
<li><a href="https://pipenv.pypa.io/en/latest/">Pipenv documentation</a></li>
<li><a href="https://realpython.com/pipenv-guide/">Real Python - Pipenv Guide</a></li>
</ul>
<h2>Conclusion</h2>
<p>So, this was a brief introduction to <code>pipenv</code> which is Python's recommended package management tool. We saw that we have control over the dependencies and packages in a virtual environment which is taken care of automatically for us. We don't have to activate the environment to install or manage any package for a project.</p>
