{
  "title": "Techstructive Weekly #5",
  "status": "published",
  "slug": "techstructive-weekly-5",
  "date": "2025-04-05T12:33:24Z"
}

<h2>Week #5</h2>
<p>This week felt like a slog, with many challenges and frustrations. However, by the end of the week, I found my stride and got excited about the direction I’m heading. It's amazing how quickly things can shift from feeling like the end of the world to experiencing a burst of excitement.</p>
<p>I also managed to get the side project of the <a href="https://github.com/Mr-Destructive/meta-ai-golang">Meta AI API wrapper</a> in Golang correctly, over the weekend will polish this project and also fix the bug in the Chat GPT anonymous client in Python.</p>
<p>Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.</p>
<h2>Quote of the week</h2>
<blockquote>
<p><strong>&quot;Success is not final, failure is not fatal: It is the courage to continue that counts.&quot;</strong><br>
— Winston Churchill</p>
</blockquote>
<p>As a developer, you will experience both triumphs and challenges. Your journey might feel like a rollercoaster, swinging from highs to lows in a single day. Embrace the mantra: Build, Iterate, Release. There is no ultimate success or failure, and no product is ever perfect. What matters is the continuous improvement and learning you gain along the way.</p>
<h2>Read</h2>
<ul>
<li>
<p><a href="https://open.substack.com/pub/developingskills/p/dont-be-an-alpha-geek?utm_source=share&amp;utm_medium=android&amp;r=1hoe7f">Don't be an Alpha Geek - John Crickett</a><br>
Just be empathetic and thoughtful about your actions and feedback. This will gradually creep into your all issues and make you a better developer, this is great advice.</p>
</li>
<li>
<p><a href="https://sophiabits.com/blog/review-your-own-prs">Review your own PRs</a></p>
<p>This is good advice, I do it as a ritual on GitHub, 8 out of 10 times, I get the feedback myself. The editor myth is real, there are things that you don’t notice in your editor, in your environment, in your flow. But as soon as the context changes, the words, and the logic seem to be distant. Believe this advice, this saves a ton of time.</p>
</li>
<li>
<p><a href="https://dev.jimgrey.net/2024/07/03/lessons-learned-in-35-years-of-making-software/">Lessons learned in 35 years of making Software</a>: I have barely lived half of 35 years, the sheer amount of experience in 35 years of software is immense respect. But what is shared here, the words disheartened me that your code will eventually be deleted, will be replaced, will be evolved. This is a harsh reality but we need to accept and move ahead in life.</p>
</li>
</ul>
<h2>Watched</h2>
<ul>
<li>
<p>Recursion with <a href="http://Boot.Dev">Boot.Dev</a>:</p>
<p>This has been well explained and visualized as well. The key part in understanding recursion for beginners is visualizing the call stack and going through the code step by step. Must watch for begineers.</p>
</li>
<li>
<p>Arden Labs Podcast: Guest Samantha Coyle with host Bill Kennedy</p>
</li>
</ul>
<p>This had some good insights and provided some guidance on how to get internships through networking and building the way up the ladder as a software developer.</p>
<ul>
<li>
<p>Rust Command Line Arguemnts by Francesco Cuila:</p>
<p>I watched this on the Live stream on Saturday, was a nice chilling stream with learning a thing or two in Rust.</p>
</li>
</ul>
<h2>Learned</h2>
<ul>
<li>
<p>Difference between cookies Add and Set in URL Values:</p>
<p><a href="https://pkg.go.dev/net/url#Values.Add">Add</a>: Appends the value to the key without replacing existing values (useful for handling multiple values for a single key).</p>
<p><a href="https://pkg.go.dev/net/url#Values.Set">Set</a>: Replaces the existing value for the key (ensures that only one value is associated with the key).</p>
<p>I learnt this will working with the Meta AI API wrapper in Golang. The API uses payload as a URL encoded body and will append key-value pairs to the request body, the subtle difference can cause nil pointer access if not initialized and used the appropriate method correctly. I think I will write a blog post on this.</p>
</li>
<li>
<p>Shuffling Two Lists keeping the order of the corresponding index the same:</p>
<p>What I was doing was testing and evaluating some results on data, and that data was coming from a set of files in a folder, I wanted to randomly shuffle those values. I wanted to track the metrics from the data with the filename, so I created this little function that shuffles two or more lists in a random order and maintains a one-on-one index mapping.</p>
<pre><code class="language-go">import random
List of file names
file_names = [&quot;file1.txt&quot;, &quot;file2.txt&quot;, &quot;file3.txt&quot;, &quot;file4.txt&quot;, &quot;file5.txt&quot;]
Corresponding sample data for each file
sample_data = [&quot;Data from file 1&quot;, &quot;Data from file 2&quot;, &quot;Data from file 3&quot;, &quot;Data from file 4&quot;, &quot;Data from file 5&quot;]
Create a list of indices
indices = list(range(len(file_names)))
Shuffle the indices
random.shuffle(indices)
Reorder both lists using the shuffled indices
shuffled_file_names = [file_names[i] for i in indices]
shuffled_sample_data = [sample_data[i] for i in indices]
print(shuffled_file_names)
print(shuffled_sample_data)
&quot;&quot;&quot;
['file3.txt', 'file5.txt', 'file1.txt', 'file2.txt', 'file4.txt']
['Data from file 3', 'Data from file 5', 'Data from file 1', 'Data from file 2', 'Data from file 4']
&quot;&quot;&quot;
</code></pre>
</li>
</ul>
<p>This is the function, nice simple and modular.</p>
<pre><code class="language-go">import random
def shuffle_lists(*lists):
&quot;&quot;&quot;
Shuffles two or more lists while keeping the order of corresponding elements the same.
Args:
*lists: Two or more lists to be shuffled.
Returns:
A tuple of shuffled lists with the same order of corresponding elements.
&quot;&quot;&quot;
# Check if all lists have the same length
if len(set(len(lst) for lst in lists)) != 1:
raise ValueError(&quot;All input lists must have the same length.&quot;)
# Create a list of indices
indices = list(range(len(lists[0])))
# Shuffle the indices
random.shuffle(indices)
# Reorder all lists using the shuffled indices
shuffled_lists = tuple([lst[i] for i in indices] for lst in lists)
return shuffled_lists
</code></pre>
<ul>
<li>
<p>LLMs are good at one single thing, Don’t ask too many things in a shot, make it sequential like a pipeline.</p>
</li>
<li>
<p>Using Tabulate in Python to format a list of dictionaries pretty:</p>
<p>This is really handy to work if you want to quickly visualise dictionaries and list like stuff.</p>
<pre><code class="language-go">from tabulate import tabulate
List of books
books = [
{'Title': '1984', 'Author': 'George Orwell', 'Year': 1949, 'Genre': 'Dystopian'},
{'Title': 'To Kill a Mockingbird', 'Author': 'Harper Lee', 'Year': 1960, 'Genre': 'Fiction'},
{'Title': 'The Great Gatsby', 'Author': 'F. Scott Fitzgerald', 'Year': 1925, 'Genre': 'Classic'}
]
Generate and print the table
print(tabulate(books, headers=&quot;keys&quot;, tablefmt=&quot;grid&quot;))
</code></pre>
<p>Output:</p>
<pre><code class="language-go">+-----------------------+---------------------+--------+-----------+
| Title                 | Author              |   Year | Genre     |
+=======================+=====================+========+===========+
| 1984                  | George Orwell       |   1949 | Dystopian |
+-----------------------+---------------------+--------+-----------+
| To Kill a Mockingbird | Harper Lee          |   1960 | Fiction   |
+-----------------------+---------------------+--------+-----------+
| The Great Gatsby      | F. Scott Fitzgerald |   1925 | Classic   |
+-----------------------+---------------------+--------+-----------+
</code></pre>
</li>
</ul>
<h2>Tech News</h2>
<ul>
<li>
<p><a href="https://hashnode.com/blog/announcing-docs-by-hashnode-the-content-engine-for-api-and-product-documentation">Hashnode Docs</a>: The content engine for API and product documentation.<br>
This brings me in awe with the Hashnode speed and quality of development. The love with which they craft these products is truly commendable, they have the best free tiers for doing almost anything related to writing as a developer.</p>
</li>
<li>
<p><a href="https://www.anthropic.com/news/artifacts">Artifacts on Antrhropic</a>: Artifacts can be used to create, view, and iterate on platform-specific work directly within <a href="http://Claude.ai">Claude.ai</a>.</p>
</li>
<li>
<p>Buildspace is closing: <a href="https://buildspace.so/">https://buildspace.so/</a></p>
<p>In a tragic turn of events, the gradually growing community of builders is sadly shut down.</p>
</li>
</ul>
<p>For more news, follow the Hackernewsletter <a href="https://mailchi.mp/hackernewsletter/714">https://mailchi.mp/hackernewsletter/714</a> and for daily developer articles, join <a href="http://daily.dev">daily.dev</a></p>
<hr>
<p>That’s it from this week, hope you did well this week, and have a happy week and weekend ahead!</p>
<p>Thank you for reading, let’s catch up in the next week.</p>
<p>Happy Coding :)</p>
<p>Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.</p>
