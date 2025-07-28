{
  "title": "SQLite importing CSV with custom separator",
  "status": "published",
  "slug": "sqlite-inline-custom-separator",
  "date": "2025-04-05T12:33:27Z"
}

<p>explaination about how i learned how to write a command for importing a csv
into a sqlite3 db with cusotm Separator with a single command inline</p>
<h2>Introduction</h2>
<p>I was exploring some Issues that I can solve given the context of the problem and the intution for the solution.</p>
<p>I have some github repositories that I always look for to learn more. Some of them are:</p>
<ul>
<li><a href="https://github.com/turbot/steampipe">Steampipe</a></li>
<li><a href="https://github.com/mindsdb/mindsdb">MindsDB</a></li>
<li><a href="https://github.com/tursodatabase">Turso</a></li>
</ul>
<p>So, navigating around for a few minutes, I landed on this <a href="https://github.com/tursodatabase/turso-cli/issues/811">issue</a>.</p>
<p>The issue is really well explained in terms of what the feature is to be added, how the usage should be, which for a contributor is half the job done.</p>
<h2>What?</h2>
<p>The <a href="https://github.com/tursodatabase/turso-cli">turso-cli</a> is a CLI for <a href="https://github.com/tursodatabase/turso">Turso</a> that is used to manage libsql databases with the <a href="https://turso.tech">turso</a> platform.</p>
<p>This issue is about adding a flag to the <code>db create</code> command to allow the user to pass in a custom separator while imporing a csv file into a sqlite3 database.</p>
<p>The only puzzle piece left is to answer the question <code>how</code>?</p>
<h2>How?</h2>
<p>So, I went in to the Codebase and found where the <code>db create</code> command has been handled and landed on this <a href="https://sourcegraph.com/github.com/tursodatabase/turso-cli/-/blob/internal/cmd/db_create.go">file</a></p>
<p>While controbuting to open source, I try to do the small steps and solve try to maintain the excitment with that progress. Because if you cannot find the solution, you try to procastinate and in the end nothing gets accomplished. So, breaking down the problem into small chunks is much helpful than solving the entire problem.</p>
<p>In this case, I  tried to add the flag in the cli which is pretty straight forward. We just add one more entry in the list of flags in the <code>db create</code> command. This step is just to add a functional CLI for taking the input of csv-separator string, however, this won't do anything for the functionality part of things, just allow the user to specify the separator/delimitor to use for parsing the CSV file.</p>
<p>Currently, the command that <code>turso db create</code> uses under the hood for creating a db from a csv file is:</p>
<pre><code class="language-bash">sqlite3 &quot;-csv&quot; &quot;dbName&quot; &quot;.import &lt;FileName&gt; &lt;TableName&gt;&quot;
</code></pre>
<p>The above command is found in the <a href="https://sourcegraph.com/github.com/tursodatabase/turso-cli/-/blob/internal/cmd/group_flag.go">group flag</a> file. To parse in the separator, we can use another string as <code>.separator &quot;;&quot;</code>, here the <code>;</code> is the character that should be used as the separator for the csv file into the db.</p>
<pre><code class="language-bash">sqlite3 &quot;-csv&quot; &quot;dbname&quot; &quot;.separator&quot; &quot;;&quot; &quot;.import &lt;FileName&gt; &lt;TableName&gt;&quot;
</code></pre>
<p>Thank you, Happy Coding :)</p>
