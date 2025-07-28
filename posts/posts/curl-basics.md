{
  "title": "Basics of curl command",
  "status": "published",
  "slug": "curl-basics",
  "date": "2025-04-05T12:33:51Z"
}

<h2>Introduction</h2>
<p>We all might have used the curl command but might be unaware of it. It's super old
and still serves a great purpose. It has been available since 1996 and still is
widely used in many embedded technologies, web API testing, CLI applications,
etc. In this article, we'll see some basics of using the curl command along with
its applications.</p>
<h2>What is the curl command?</h2>
<p>Curl or cURL command is the utility or tool to access the internet from the command
line interface using various protocols. This looks trivial but it can blow up
your mind! Most people use this tool for fetching and processing the
data from the servers/internet from their terminal without the browser but
there is a lot more to it. It is used in various embedded devices for accessing
the network in a lightweight and accessible way. Let's see how you can use the curl
command from the very basics.</p>
<h2>Why do we need it?</h2>
<p>Before we talk about how to use the curl command let's talk about why might we need
that? There are a lot of reasons and it even depends on the application you are
using.  You can use curl to test your API, well there are other tools like
POSTMAN, Insomnia, etc but for keeping things simple you can quickly get in
with curl and test some endpoints.  You might require curl for creating some
CLI applications that require fetching/posting to an URL over the internet.
If you are using the terminal, curl integrates really very well with the shell
programming languages like BASH, ZSH, etc So, after making WHY out of the way,
let's start with the actual content.</p>
<h2>Structure of curl command</h2>
<p><strong>curl or Client URL is a command-line utility that helps in accessing/posting
data with various protocols over the internet.</strong> It basically serves as a
bare-bones browser URL search bar.  You can't render those pages like the
actual GUI, and all but you can get is the HTML source code, JSON response,
etc.  That's still quite powerful and used in tons of applications.</p>
<pre><code>curl URL arguments 
</code></pre>
<p>The above is a basic structure of the curl command. We see the argument
structure in-depth in the next section. Firstly, let's take a simple curl command with just the URL is given.</p>
<pre><code class="language-bash">curl &quot;https://github.com&quot;   
</code></pre>
<p>From this query to <code>github.com</code>, you are literally going to <code>GitHub.com</code> and getting a response as the entire HTML source code of the page.
If you don't want to spam the output in the terminal, you can redirect the output to a file.</p>
<pre><code class="language-bash">curl &quot;https://github.com&quot; &gt;temp.html
</code></pre>
<p>With this command, we store the output of the command in the file temp.html, it can be any other file you like.</p>
<h3>Arguments</h3>
<p>It turns out that you can even parse in certain arguments to the <code>curl</code> command to get some desired and modified results. Let's take a look at some of them.
The <a href="https://curl.se/docs/manpage.html">entire list of arguments</a> is quite huge
and baffling, but this shows how customizable the command is.</p>
<ul>
<li><code>-s</code> (silent the progress bar)</li>
<li><code>-X</code> (web requests <code>POST, GET, etc</code> to the URL)</li>
<li><code>-o</code> (output to a file)</li>
<li><code>-H</code> ( provide Header to the request)</li>
<li><code>-d</code> (providing the data e.g. in POST request)</li>
</ul>
<pre><code class="language-bash">curl -s -o &quot;https://github.com&quot; temp.html
</code></pre>
<p>This command doesn't load the progress bar and simply outputs the response in a
file, making the execution process in the terminal clean.</p>
<h3>Integration with other commands</h3>
<p>As said, the <code>curl</code> command can be well integrated with the other commands using piping in shell, assigning to variables, and so on.</p>
<p>Let's see how we can convert the <code>JSON</code> response to a BASH variable.</p>
<pre><code class="language-bash">resp=$(curl -H &quot;api-key: N2vDzMyEeYGTxjUTePhC8bYd&quot; https://dev.to/api/users/me)

echo $resp
</code></pre>
<p>Here, we are fetching the <code>JSON</code> response from the <code>dev.to</code> <a href="https://developers.forem.com/api/">API</a>,The wired string <code>N2vDzMyEeYGTxjUTePhC8bYd</code> is my <a href="https://dev.to/settings/account">dev.to API token</a>(don't worry I have revoked it:) ) we have provided an argument <code>-H</code> that is a Header for accepting a <code>Json</code> response.
We can store the contents of the curl command by using the <code>$( )</code> and assigning that to the variable name of your choice.</p>
<pre><code class="language-bash">username=$(curl -H &quot;api-key: N2vDzMyEeYGTxjUTePhC8bYd&quot; https://dev.to/api/users/me | grep -o -P '(?&lt;=username&quot;:&quot;).*(?=&quot;,&quot;name)')
</code></pre>
<p>Here, we have stored the username from a <code>JSON</code> response to the variable username. We have piped the curl command so that we can work with that <code>JSON</code> response and modify the contents and then store the final results in a variable.
In this case, we are using <code>grep</code> to filter out the content between the key <code>username</code> and <code>name</code>, thus we get the value we desired. To see the value you can always run the echo command as below:</p>
<pre><code class="language-bash">echo $username
</code></pre>
<p>So, that's how the <code>curl</code> command integrates flawlessly with BASH and other shell programming languages.</p>
<h2>Where is it used?</h2>
<p><code>curl</code> is actually used in API testing, CLI applications, Web Scrapping, etc. It's a great tool for terminal lovers. Let's see where we can use the curl command actually to make some good projects.</p>
<h3>API Testing</h3>
<p>We can use, <code>curl</code> to test an API, it might be API you would have made or to simply test and play with other API available publicly. You can get an in-depth guide about <a href="https://www.codepedia.org/ama/how-to-test-a-rest-api-from-command-line-with-curl/">Testing a REST API with curl</a>.
Actually, curl can do more than just testing, I have made a <a href="https://gist.github.com/Mr-Destructive/80860664b1014ef0b94092d68ead1044">bash script</a> that actually posts some data over a database through the API so that I don't have to do that manually. That is the kind of automation I love to do and curl! Just did that.</p>
<h3>Web Scrapping</h3>
<p>Web-scrapping is usually trending with Python, but I have done that with BASH.
That might be an outdated idea but is a good task to learn the basics of
Web-scrapping with BASH ;). I must say that sed, awk, grep are the tools are
powerful like heck in doing these tricks. I have made this
<a href="https://mr-destructive.github.io/techstructive-blog/bash/2021/07/15/BASH-Crypto-Coingecko.html">crypto-currency</a>
and
<a href="https://mr-destructive.github.io/techstructive-blog/bash/2021/07/27/BASH-script-dictionary-scrap.html">dictionary</a>
scrapper with BASH. Web-scrapping can be done with the curl command by fetching to
an API if any or any website. We need to search and find the particular fields,
classes, or ids the elements the required data might be into and then extract
and filter using the tools like grep, sed or awk.</p>
<h3>CLI Applications</h3>
<p>We can make CLI applications like creating a terminal view of existing
applications using their APIs or website. I recently made a CLI for
<a href="https://github.com/Mr-Destructive/crossposter">cross-posting articles</a> to
dev. to, hashnode and medium. That is a project still in progress(tons of bugs)
but still serving a decent job. Definitely <code>curl</code> might not be the only command
that works here, but the project might look so incomplete without <code>curl</code>.</p>
<p><strong>There might be other applications as well, who knows there is a lot to do with this command.</strong> If you know one, please let everyone know in the comments.</p>
<h3>References:</h3>
<p>Special Thanks to the creator of the curl command - <a href="https://github.com/bagder">Magnus Daniel Stenberg</a> and the developers who are still contributing and maintaining the great project.</p>
<h3>Conclusion</h3>
<p>So, from this article, we were able to understand the basics of the <code>curl</code> command and understand its applications in actual programming stuff. Hope you liked it. Thanks for reading and until then Happy Coding :)</p>
