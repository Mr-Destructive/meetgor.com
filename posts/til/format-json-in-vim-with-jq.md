{
  "title": "Format JSON in Vim with JQ",
  "post_dir": "til",
  "type": "til",
  "status": "published",
  "slug": "format-json-in-vim-with-jq",
  "date": "2025-04-05T12:33:15Z"
}

<p>We can use</p>
<pre><code>:%!jq .
</code></pre>
<p>To instantly format a json file opened in vim/neovim.
I relied on online foramtter and opening vscode for formatting.
But this quick tool just saves a lot of hassle. Opening VS Co** == shutting down the laptop, it clogs the memory badly, so I avoid opening VS Co** as much as possible.</p>
<p>Yes, to use this, you need to have <a href="https://jqlang.org/">jq</a> installed on your system.</p>
<p>If you are elsewhere and want to format json file using jq, you can quickly do this:</p>
<pre><code>jq . input.json | sponge input.json

# OR

jq . input.json &gt; tmp.json &amp;&amp; mv tmp.json input.json

</code></pre>
<p>If you are in Vim and want to format other file, just add <code>!</code> at the start of the command and that will run in the shell for you, so you don't have to exit vim.</p>
