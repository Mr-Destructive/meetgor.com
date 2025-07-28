{
  "title": "Building Golang from Source v1.23 and Above",
  "status": "published",
  "slug": "golang-build-from-source-1-24-above",
  "date": "2025-04-05T12:33:21Z"
}

<h2>Introduction</h2>
<p>Are you excited to try out the latest golang version, or test out your changes (cooking some serious stuff?), or install some random golang version? Then let’s explore one of the easiest ways to install golang on your system (Linux).</p>
<h2>Building Go from source 1.23+</h2>
<p>The process for installing and building golang from source is quite simple.</p>
<ul>
<li>
<p>Clone the repo</p>
</li>
<li>
<p>Build the binary</p>
</li>
<li>
<p>Set Environment Variables</p>
</li>
<li>
<p>Export the binary to the system path</p>
</li>
</ul>
<p>For detailed other steps, you can follow <a href="https://go.dev/doc/install/source">this guide</a>.</p>
<h3>Clone the repo</h3>
<p>Just clone the repo from GitHub or Google Git repo.</p>
<pre><code class="language-bash">git clone https://github.com/golang/go

OR

git clone https://go.googlesource.com/go
</code></pre>
<p>This will install the golang source code required to build the golang binary and the ecosystem (gofmt + standard library + test suite).</p>
<p>Then let’s navigate to the cloned repo and we can build the golang from the source.</p>
<h3>Build it</h3>
<p>We need to run the bash script in the folder to build the binary. They <code>all.bash</code> can be run to build the binary which will be stored in the <code>go/bin</code> folder <code>go/bin/go</code> and <code>go/bin/gofmt</code> files. These two binaries will be generated and are required in the Golang ecosystem.</p>
<pre><code class="language-bash">cd src

./all.bash
</code></pre>
<p>Once we have the binaries in the specified folder, we can move ahead to make the environment understand where the actual binary is located.</p>
<h3>Environment Variables</h3>
<p>The <code>GOROOT</code> and <code>GOPATH</code> variables are required for the golang ecosystem to work as expected. The <code>GOROOT</code> is set as the path to the actual golang source repository, the cloned repository from which we built the binary. This <code>GOPATH</code> is the path where Golang stores the external repositories or modules for use anywhere in the system.</p>
<pre><code class="language-bash">export GOROOT=path_to_clone_repo

export GOPATH=$(HOME)/go

export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
</code></pre>
<p>The <code>PATH</code> environment needs to be updated with the <code>GOROOT</code> and <code>GOPATH</code> to make the binaries in those paths visible and accessible to the system.</p>
<blockquote>
<p>NOTE: If you are installing the golang from source when you have already a version of golang installed on your system, then you need to make sure you do not mess up the <code>GOROOT</code> and <code>GOPATH</code>.</p>
<p>Those could be juse set with the current shell session, as you do not want this golang version permantely on the system, if you do requrie the newly installed golang version as your default, then you can set this environment variables in your shell config.</p>
</blockquote>
<p>Finally, we can now set the binary as something different because we do not want it to clash with the default golang version.</p>
<h3>Run the binary</h3>
<p>The binary can be stored in the <code>/usr/local/bin/</code> to make any binary available in the system from anywhere. This is not necessary but handy if you are going to use it commonly but don’t need it as the default golang version.</p>
<pre><code class="language-bash"># The binary is stored in the 
# path_to_cloned_repo/bin
# Whatever you want to name the binary

cp bin/go /usr/local/bin/go-dev

OR

cp bin/go /usr/local/bin/go1.24
</code></pre>
<p>Once this is done. we can check the installed golang version</p>
<pre><code class="language-bash">go1.24 version
</code></pre>
<p>With this, you can use it as go1.24 or go-dev as the binary name.</p>
<p>So, that is how we install and build from source any golang version above 1.23.</p>
<h2>Conclusion</h2>
<p>For context, I wanted to check out the latest changes in 1.24, so I cloned the repo and after some trial and error of some commands to build the golang version, I was able to do it correctly. So just decided to share it here, hope you found it helpful.</p>
<p>Thank you for reading.</p>
<p>Happy Coding :)</p>
