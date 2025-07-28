{
  "title": "Python Pipreqs: Generate requirements file from the imported packages",
  "status": "published",
  "slug": "python-pipreqs",
  "date": "2025-04-05T12:33:36Z"
}

<h2>Introduction</h2>
<p><a href="https://pypi.org/project/pipreqs/">Pipreqs</a> is a python package that allows us to list all the pacakges which are imported in a python project. This is a great package for reducing the amount of redundant packages for a project.</p>
<h2>Install pipreqs</h2>
<p>You can install pipreqs with one of the many ways with pip, pipx, or any other pacakge management tool. I personally use pipreqs with <code>pipx</code> as it remains isolated from the rest of my project dependencies.</p>
<h3>Using simple pip install</h3>
<p>We can install with pip by creating a virtual environment or in a existing virtual environment.</p>
<pre><code>pip install virtualenv venv
source venv/bin/activate

pip install pipreqs
</code></pre>
<h3>Using pipx</h3>
<p>We can install pipreqs with pipx. <a href="https://pypi.org/project/pipx/">Pipx</a> is also a python package but used as a tool to install any cli specific tool with the isolated environment.</p>
<pre><code>pipx install pipreqs
pipx run pipreqs
</code></pre>
<h2>Using pipreqs</h2>
<p>We need to specify the encoding, which is used for reading the files while capturing the imports from the project.</p>
<pre><code>pipx run pipreqs --encoding=utf-8 .
</code></pre>
<p>Additionaly, we can specify the <code>path</code> or filename where it will be used to save the imported packages. The <code>--savepath</code> option takes in the path to the file where you want to generate the list of the packages to be installed.</p>
<pre><code>pipx run pipreqs --encoding=utf-8 --savepath reqs.txt . 
</code></pre>
<p>Though this doesn't guarentee all the requirements for a file, it is really helpful for explicitly used packages in the python project.</p>
