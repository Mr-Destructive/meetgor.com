{
  "title": "Techstructive Weekly #15",
  "status": "published",
  "slug": "techstructive-weekly-15",
  "date": "2025-04-05T12:33:21Z"
}

<h2>Week #15</h2>
<p>A wonderful week, bringing back the joy of programming, after some slumpy weeks, a refreshing and mind challenging week. The start was drowzy after a long weekend, but as the week progressed, I had to shift gears and complete the things that took some deep concentration and thinking, finally the results were worth the efforts. Everything in its place, the reward for the focus and hard work.</p>
<p>After a live stream and video creation in the end of week last week, I didn’t do much on Sunday, just took a break and spent the diwali removing the negativity in life. I am not mentioning it enough but I continue the writing routine, 35 days and counting, the writing streak is un-broken, I completed one section of writing of 21 days and now moving on to another section of 17 days, combining a total of (9+21+5(till date)) 35 articles. This is the motivation for me waking up everyday, writing will get me to the places I can dream of.</p>
<h3>Stats for the week</h3>
<ul>
<li>
<p>FIxed a ton of bugs</p>
</li>
<li>
<p>Brainstormed on refactoring code</p>
</li>
</ul>
<h3>Plans for the next week</h3>
<ul>
<li>
<p>Create 2 articles</p>
</li>
<li>
<p>Record 1 videos</p>
</li>
<li>
<p>Live Stream at least once</p>
</li>
</ul>
<h2>Quote of the week</h2>
<blockquote>
<p>&quot;Strength doesn’t come from what you can do. It comes from overcoming the things you once thought you couldn’t.&quot;</p>
<p>— Rikki Rogers</p>
</blockquote>
<p>I once thought before the week, that this week will be the make or break week for the things I have been doing for the past months, and finally in the end was able to break shakles and get into deep programming flow state, it felt good, it felt energetic, removing the guilt and imposter syndrome that has been creeping in.</p>
<p>Strength is not what you have, it is about learning what you don’t. It is about taking the challenge head on.</p>
<h2>Read</h2>
<ul>
<li>
<p><a href="https://rmoff.net/2023/07/19/blog-writing-for-developers/">Writing for developers</a>:<br>
This is really a good read, writing brings clarity and clarity is what shapes you to pick up correct directions while developing anything. All the time, I start to write an article, I ran off a tangent on one or the other interesting thing that I had no idea about and assumed I knew it.</p>
</li>
<li>
<p><a href="https://twitter.com/hnasr/status/1852537428227375482">How to become a Good Backend Engineer</a></p>
<p>This is really well summarised and insightful read for me at least, it gives a birds-eye view and ignites a curiosity to dive deeper into a specific topic or domain, which is essential for any developer.</p>
</li>
</ul>
<h2>Watched</h2>
<ul>
<li>
<p><a href="https://youtu.be/kCc8FmEb1nY?si=UV9hrQh2Uw8dc57Z">Building ChatGPT like LLM Model on a small scale from scratch</a>:</p>
<p>I have not completed it but just understanding the architecture and what the LLM does is vital to work and adapt to the changes in the evolving world.</p>
</li>
</ul>
<h2>Learned</h2>
<ul>
<li>
<p><a href="https://github.com/openai/tiktoken">TikToken</a> the tokeniser that powers GPT and its family of models: I want to explore this and create a project around it. I learnt that is the model that actually determines the number of tokens passed around the chats while interacting with the LLMs.</p>
</li>
<li>
<p>Using the <code>next</code> method in python to get the first non-empty or truth value from a dict or any expression.</p>
<pre><code class="language-go">row = {
    &quot;name&quot;: {
        &quot;value&quot;: &quot;123&quot;,
        &quot;position&quot;: [40, 40],
    },
    &quot;city&quot;: {
        &quot;value&quot;: &quot;mumbai&quot;,
        &quot;position&quot;: [50, 30],
    }
}
row_position = next((attrs[&quot;position&quot;] for attrs in row.values() if isinstance(attrs, dict) and attrs.get(&quot;value&quot;)), None)
</code></pre>
</li>
<li>
<p>Combining Pandas Dataframes</p>
</li>
</ul>
<p>Let’s say I have a list of dataframes, i want to combine them with certain number but not the contents, just append the next df to the end of the current df. I used the <a href="https://pandas.pydata.org/docs/reference/api/pandas.concat.html">pd.concat</a> function to combine the slices of pandas dataframe in a single list.</p>
<pre><code class="language-go">PAGE_LIMIT = 2
PAGES = 5
for i in range(0, PAGES, PAGE_LIMIT):
if i + PAGE_LIMIT &lt;= num_of_pages:
df_batch = pd.concat(df_list[i:i + PAGE_LIMIT], axis=0)
else:
df_batch = pd.concat(df_list[i:], axis=0)
</code></pre>
<ul>
<li>
<p>Deleting elements from the list given the indices</p>
<p>I know this could be easily gotten from GPT but feels good to do it yourself sometimes.</p>
</li>
</ul>
<pre><code class="language-go">def delete_elements_by_indices(lst, indices):
indices_set = set(indices)
return [item for idx, item in enumerate(lst) if idx not in indices_set]
lst = [10, 20, 30, 40, 50]
indices_to_delete = [1, 3]
result = delete_elements_by_indices(lst, indices_to_delete)
print(result)
Output: [10, 30, 50]
</code></pre>
<p>Fun tidbit,</p>
<p>Chat GPT search is just awesome, it recognises me? really?</p>
<p>![](https://substack-post-media.s3.amazonaws.com/public/images/ee8f7e63-6ee2-4257-823c-179855c99054_835x449.png align=&quot;left&quot;)</p>
<p>Even my youtube channel which is barely 2 months old</p>
<p>![](https://substack-post-media.s3.amazonaws.com/public/images/413b9325-6eff-4bda-bf79-eb862863e7a6_835x449.png align=&quot;left&quot;)</p>
<h2>Tech News</h2>
<ul>
<li>
<p><a href="http://Hertz.dev">Hertz.dev</a> launches the conversational audio generator model</p>
</li>
<li>
<p><a href="https://stripe.com/blog/workbench-a-new-way-to-debug-monitor-and-grow-your-stripe-integration">Stripe launches Workbench</a>: A new way to debug, monitor, and grow your Stripe integration</p>
</li>
</ul>
<p>For more news, follow the <a href="https://buttondown.com/hacker-newsletter/archive/hacker-newsletter-721">Hackernewsletter</a> and for daily developer articles, join <a href="http://daily.dev">daily.dev</a></p>
<p>That’s it from this week, hope you did well this week, and have a happy week and weekend ahead!</p>
<p><a href="https://techstructively.substack.com/p/techstructive-weekly-15/comments">Leave a comment</a></p>
<p>Thank you for reading, let’s catch up in the next week.</p>
<p>Happy Coding :)</p>
<p>Thanks for reading Techstructive Weekly! This post is public so feel free to share it.</p>
<p><a href="https://techstructively.substack.com/p/techstructive-weekly-15?utm_source=substack&amp;utm_medium=email&amp;utm_content=share&amp;action=share">Share</a></p>
<p>Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.</p>
