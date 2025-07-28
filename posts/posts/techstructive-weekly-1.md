{
  "title": "Techstructive Weekly #1",
  "status": "published",
  "slug": "techstructive-weekly-1",
  "date": "2025-04-05T12:33:26Z"
}

<h1>Week #1</h1>
<p>This week, I had a ton of fun on the side as I finally made the MVP of a side project. An SSG with a Content Management System-like interface. an SSG with an editor that syncs up the posts from a database.</p>
<p>In this process, I learned about Cloudflare workers and golang. At work, I had some thinking and experimentation with LLMs.</p>
<p>Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.</p>
<h2>Quote of the Week</h2>
<blockquote>
<p>&quot;The only limit to our realization of tomorrow is our doubts of today.&quot;</p>
<p>— Franklin D. Roosevelt</p>
</blockquote>
<p>The LLM thing is affecting the mindset of people, more of which people think that it is going to replace humans, but really? Can a computer program that computes some numbers based on some other numbers even compare to human thoughts? That’s way far from what you are capable of, don’t waste the energy in thinking this.</p>
<h2>Wrote</h2>
<p>Well, I didn’t write any articles this week, I am in the research phase for writing the Golang Patch Method, which is an interesting topic. It has way more things than I expected to be. This is what I love about writing articles/blogs; it helps me learn concepts in such detail.</p>
<p>Hopefully, next week, there will be an article to share in this newsletter.</p>
<h2>Read</h2>
<ul>
<li>Pandas’ Dataframe is Column Major Data Structures</li>
</ul>
<p>![](https://substack-post-media.s3.amazonaws.com/public/images/c5dc1fee-2d1e-4892-b219-4b96f6998ab5_288x288.png align=&quot;left&quot;)</p>
<p><a href="https://blog.dailydoseofds.com/p/what-happens-when-you-append-rows?utm_source=substack&amp;utm_campaign=post_embed&amp;utm_medium=web">Daily Dose of Data Science</a></p>
<p><a href="https://blog.dailydoseofds.com/p/what-happens-when-you-append-rows?utm_source=substack&amp;utm_campaign=post_embed&amp;utm_medium=web">What Happens When You Append Rows to a Pandas DataFrame</a></p>
<p><a href="https://blog.dailydoseofds.com/p/what-happens-when-you-append-rows?utm_source=substack&amp;utm_campaign=post_embed&amp;utm_medium=web">Daily Dose of Data Science Free Book | Deep Dives…</a></p>
<p><a href="https://blog.dailydoseofds.com/p/what-happens-when-you-append-rows?utm_source=substack&amp;utm_campaign=post_embed&amp;utm_medium=web">Read more</a></p>
<p><a href="https://blog.dailydoseofds.com/p/what-happens-when-you-append-rows?utm_source=substack&amp;utm_campaign=post_embed&amp;utm_medium=web">5 months ago · 24 likes · 4 comments · Avi Chawla</a></p>
<ul>
<li>
<p>Deno and HTTP imports: <a href="https://deno.com/blog/http-imports">Article</a></p>
</li>
<li>
<p>Pixel by Pixel, could create a vault of articles: <a href="https://www.pawlean.com/vault">Pauline’s Blog</a></p>
</li>
<li>
<p>Row Scanning in Golang with reflect package: <a href="https://stackoverflow.com/questions/56525471/how-to-use-rows-scan-of-gos-database-sql">Stackoverflow Question</a></p>
</li>
</ul>
<h2>Watched</h2>
<ul>
<li>
<p><strong>Cloudflare Workers 101: A Crash Course on Cloudflare Workers</strong></p>
</li>
<li>
<p><strong>Git Log</strong></p>
</li>
<li>
<p><strong>How good are LLM? Think about it</strong></p>
</li>
</ul>
<h2>Learnt</h2>
<ul>
<li>
<p><strong>Cloudflare Workers Setup</strong>: <a href="https://developers.cloudflare.com/workers/examples/">Documentation</a><br>
It has good examples of some standard use cases and applications.</p>
</li>
<li>
<p><strong>Getting table data using pymupdf</strong>: <a href="https://artifex.com/blog/table-recognition-extraction-from-pdfs-pymupdf-python">Article</a><br>
Getting some table position information of tables present in a pdf document. It can be really useful for the extraction of table-related data. For instance, I could get the width and height of the table to further use it in extraction techniques.</p>
</li>
<li>
<p><strong>HTTP Patch method is not what you think it is</strong>: <a href="https://imantumorang.com/posts/http-patch-method-ive-thought-the-wrong-way/">Article</a><br>
I am also surprised that the PATCH method is different from than rest of the methods, we actually need to send commands/instructions to the server in order to make it understand what and how to update the resource-specific fields or entities.</p>
</li>
<li>
<p><strong>Setting a Github Pages on GitHub Actions Workflow</strong>:<br>
I used this snippet in my <a href="https://github.com/Mr-Destructive/tuxo/blob/main/.github/workflows/cronjob.yml">Tuxo SSG</a>, which would serve as the output directory of the generated SSG files to GitHub Pages.</p>
<pre><code class="language-go">    - name: GitHub Pages
      uses: crazy-max/ghaction-github-pages@v3
      with:
        target_branch: output-branch
        build_dir: my_app/
        jekyll: false
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
</code></pre>
</li>
</ul>
<h2>Tech News</h2>
<ul>
<li>
<p><a href="https://fastht.ml/">FastHTML</a>: A Python framework for developing web applications. I would like to explore this in the upcoming weeks.</p>
</li>
<li>
<p><a href="https://ai.meta.com/ai-studio/">AI Studio</a>: Create fictional AI characters to chat with by Meta. This is pretty good use case for authors and fantasy writers. It could allow them to get to know their characters in a better way.</p>
</li>
</ul>
<p>As always, I recommend going through the <a href="https://mailchi.mp/hackernewsletter/710">HackerNews Newsletter</a> to get all the tech news.</p>
<p>That’s it from this week, hope you did well this week, and have a happy week and weekend ahead!</p>
<p>Thank you for reading, let’s catch up in the next week.</p>
<p>Happy Coding :)</p>
<p>Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.</p>
