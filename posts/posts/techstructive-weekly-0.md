{
  "title": "Techstructive Weekly #0",
  "status": "published",
  "slug": "techstructive-weekly-0",
  "date": "2025-04-05T12:33:26Z"
}

<h2>Week #0</h2>
<p>This week was about writing one of the much-procrastinated blog posts, starting a YouTube channel, and going through some heavy debugging sessions with my colleagues at work.</p>
<h3>Quote of the week</h3>
<blockquote>
<p>It’s harder to say things which are easier to do, and easier to say things which are harder to do<br>
(It’s just my random thought converted to a quote)</p>
</blockquote>
<p>Seriously think about this quote, it is easier to say I will go to the gym daily, and read books, but working on them is it hard. Why is the harder to say that I am scrolling through Twitter, and Instagram, having burnout, but taking action against it is sheer a change in mindset?</p>
<p>Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.</p>
<h2>Wrote</h2>
<ul>
<li><a href="https://www.meetgor.com/nginx-02-web-servers/"><strong>Nginx Creating Web Applications</strong></a> - Part 2 of my Nginx Survival Guide<br>
I finally wrote my second post in the Nginx Survival guide, after procrastinating for a couple of months, this weekend, I’ll be writing a post on the Golang series. So, keep an eye out for that.</li>
</ul>
<h2>Read</h2>
<ol>
<li>
<p><a href="https://dev.to/syeo66/im-tired-of-it-5fe2"><strong>I am Tired of AI content</strong></a>: It is tiring to read AI-generated content spam on most of the blogging/social media platforms. Please, guys, write genuinely value-adding (authentic) content for yourself at least, don’t try to shine in the eyes of others, it would do more harm than good.</p>
</li>
<li>
<p><a href="https://survey.stackoverflow.co/2024/"><strong>StackOverflow Developer Survey 2024</strong></a>: This was important to reflect on the state of developers in 2024, especially after the wild layoff and AI-hype phase. Also one of the things I noticed was that remote and hybrid companies are still covering around 80% of the jobs, which is a good number, to be honest.</p>
</li>
<li>
<p><a href="https://saiyampathak.medium.com/2021-year-i-want-and-do-not-want-to-remember-4229c4a32015">Saiyam Pathak’s 2021 Year reflection</a>: There is always a year in someone’s life that they don’t want to remember, it is hard to lose someone very close to you. May life never cross such paths. But Saiyam Pathak’s journey is an inspiration to take from, he is an ambassador in CNCF, founder of Kubesimplify, and an online educator and mentor. What a journey, what a man.</p>
</li>
</ol>
<h2>Watched</h2>
<ol>
<li>
<p><strong>SQLite CTEs</strong>: A recursive CTE in SQLite is a handy way to run repeated queries to build up results step-by-step, like creating a sequence of numbers or working with hierarchical data. You start with a base query, then keep adding to it until you get the full set of results. It’s a neat trick for straightforwardly managing complex data patterns.</p>
</li>
<li>
<p><strong>Git Fork Vs Clone</strong>: This video is more than explaining Fork vs Clone, obliviously I know the difference between them but Primeagen goes beyond that and explains the workflow for contributing to an open-source project, and trying to matter more than merging :) Also, I can relate how much the importance of viewing the changes on GitHub before creating the PR is, have been there and I do it all the time, had saved me a lot of time from colleagues’ reviews.</p>
</li>
</ol>
<h2>Learnt</h2>
<ul>
<li>
<p><strong>Google Cloud Run Min and Max Instances</strong>: We can define certain number of containers to run at minimum or maximum. The minimum number prevents the time wasted on cold start if the service is inactive for long, and maximum number would allow us to scale the number of incoming requests(it largely depends on the implementation as well). GCP Cloud Run <a href="https://cloud.google.com/run/docs/configuring/min-instances">Minimum</a> and <a href="https://cloud.google.com/run/docs/configuring/max-instances">Maximum</a> Instances Config.</p>
</li>
<li>
<p><a href="https://docs.python.org/3/library/itertools.html#itertools.product">Python Itertools Product</a>: Imagine you're tasked with exploring every possible outcome of a scenario where each decision can be either negative, neutral, or positive. To systematically tackle this, Python’s <code>itertools.product</code> is handy. By setting up <code>itertools.product([-1, 0, 1], repeat=n)</code>, that generates every combination of these choices, repeated <code>n</code> times.</p>
<p>For instance, if <code>n</code> were 2, output (-1, -1), (-1, 0), (0, 1), and so forth—every pairing of -1s, 0s, and 1s. This method spares you from manually scripting nested loops, ensuring clarity and efficiency in your code. To further inject the output of this in the python code, cast the itertool object to list as <code>list(itertools.product([-1, 0, 1], repeat=n)).</code></p>
</li>
<li>
<p>Becoming a Never Nester: I kind of start with simple if-else but end up with nested if else conditions, I don’t know I am just hacking around, making the solution work. That’s the first phase right of make it work, make it correct, make it fast. Sometimes, after making it work, out of excitement I forget to make it readable and correct, which might lead to bad code down the line. So, I need to learn to make a habit to resist using nested if-else(or any nested code) at the first place.</p>
</li>
</ul>
<h2>Tech News</h2>
<ul>
<li>
<p><a href="https://ai.meta.com/blog/meta-llama-3-1/">Meta released its updated LLM Model</a>: Llama 3.1, It’s a 405B model, offering support for 128K context length and multiple languages. That is a big number but not sure if it is a big shift in the quality of the LLMs these days, the trend is moving a bit slowly now. Possibly a cap before something improves the results.</p>
</li>
<li>
<p><a href="https://chatgpt.com/search">OpenAI to launch AI search</a>: So there is finally a worthy rival to Google search it seems, but can’t comment till we see the results.</p>
</li>
<li>
<p><a href="https://mistral.ai/news/mistral-large-2407">Mistral Large 2</a>: It has 123 billion parameters, improved code generation, multilingual support, and better reasoning. It features a 128k context window and supports numerous languages and coding languages.</p>
</li>
</ul>
<p>For more updates, I follow <a href="https://mailchi.mp/hackernewsletter/567">hackernews</a>.</p>
<p>That’s it from this week, hope you did well this week, and have a happy week and weekend ahead!</p>
<p>Thank you for reading, let’s catch up in the next week.</p>
<p>Happy Coding :)</p>
<p>Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.</p>
