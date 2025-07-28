{
  "title": "Podevcast: A single source for developer podcasts",
  "status": "published",
  "slug": "podevcast-project",
  "date": "2025-04-05T12:33:48Z"
}

<h2>Introduction</h2>
<p>Hello Developers! Want to listen to programming podcasts from a single place? Podevcast is the place you should be searching for.</p>
<p>I am Meet Gor and I present this project as a submission to the Netlify x Hashnode Hackathon. Podevcast is a webpage(static) for listening to podcasts centered around developers and programming. Just pick your favorite one and start listening straight away. Let's dive into the making of Podevcast. Head on to https://podevcast.netlify.app/ to check out the live app.</p>
<h2>What is Podevcast</h2>
<p>Podevcast is a web application or a static site that renders the top programming/development podcasts. You can listen to the top podcasts around the developer community from a single source.</p>
<blockquote>
<p>Listen to your favorite developer podcasts with Podevcast</p>
</blockquote>
<p>Podevcast is a static site generated using a script. There is a static site generator that is heavily done in Python and deployed to Netlify. You can simply listen to the podcasts on the web page or go to the canonical page of the podcast episode. From the canonical page, you can choose to hop to your chosen music player, but the default music player should be fine for casual listening. The core idea is to keep things in a single place for developer podcasts.</p>
<p><a href="https://podevcast.netlify.app/">Podevcast</a></p>
<p><a href="https://github.com/Mr-Destructive/podevcast">Source Code</a></p>
<h2>Preview</h2>
<p>Podevcast has popular developer podcasts like <code>Command Line Heroes</code>, <code>The Python Podcast</code>, <code>The freeCodeCamp Podcast</code>, and many others to choose from. You can go into categories for looking at a specific podcast.</p>
<h3>Application Demonstration</h3>
<p>Here's a small demonstration of the Podevcast application.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1645200224921/GC8gmxUzX.gif" alt="Podevcast Preview gif"></p>
<p>Podevcast has multiple pages like:</p>
<ol>
<li><a href="https://podevcast.netlify.app/">Home page</a></li>
<li><a href="https://podevcast.netlify.app/list">Podcast page</a></li>
<li><a href="https://podevcast.netlify.app/the_real_python_podcast/ep/1/">Episode page</a></li>
<li><a href="https://podevcast.netlify.app/command_line_heroes/">Podcast List page</a></li>
<li><a href="https://podevcast.netlify.app/category/">Categories page</a></li>
</ol>
<p>The Home page has the latest episode of all the podcasts. It also has an audio player to play on the go.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1645113477/blog-media/iafi8nthhj0vvvrcbhka.png" alt="Podevcast home page"></p>
<p>The Podcast List page has the list of all the Podcasts available in the project. It has the name of the podcast with the link to the podcast page that has the list of all the episodes of that podcast.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1645113598/blog-media/cnprgufs3lrouvgdl8jn.png" alt="Podevcast Podcast list"></p>
<p>The categories page has a list of categories of the podcasts like Web-development, backend, frontend, data science, DevOps, and so on. More categories will be added soon.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1645113626/blog-media/uloq4xi1d4zfo8sfl7bm.png" alt="Podevcast Categories"></p>
<p>The Episode page has the audio player, the summary of the episode, canonical episode, and podcast page.</p>
<p><img src="https://res.cloudinary.com/techstructive-blog/image/upload/v1645113654/blog-media/omqks44p8b3u7jclkhgz.png" alt="Podevcast Episode page"></p>
<h2>Why Podevcast?</h2>
<p>Listening to music is one thing and listening to podcasts is different. I wanted a place from where developers can listen to developer-specific podcasts from a single source not just give out the article <strong>&quot;Top 10 podcast you should be listening to as a developer&quot;</strong>. Having played around with python and some libraries like feedparser and jinga previously I saw this Hackathon as an opportunity to convert the idea into a project. It fits the JAMStack area well from the Hackathon and project perspective.</p>
<h2>Tech Stack</h2>
<ul>
<li>Python
<ul>
<li><a href="https://pypi.org/project/feedparser/">feedparser</a></li>
<li><a href="https://pypi.org/project/Jinja2/">jinga2</a></li>
</ul>
</li>
<li>GitHub Actions</li>
<li>HTML / CSS</li>
</ul>
<p>The data is extracted from various RSS Feeds using the feedparser library in Python.</p>
<p>Using GitHub Actions, the feed is refreshed every 24 hours to fetch the latest episodes from the respective podcast feeds. Basically, the GitHub action triggers a Netlify deployment that in turn generates the static site by running the script.</p>
<p>The command for running the script on Netlify and generating the <code>Podevcast</code> webpage is :</p>
<pre><code>pip install -r rquirements.txt &amp;&amp; python script.py
</code></pre>
<p>And the directory for deployed web pages (published directory) is <code>site</code> which contains all the <code>HTML</code> files that can be rendered as the website itself.</p>
<h3>Source Code</h3>
<p>The project is available on <a href="https://github.com/Mr-Destructive/podevcast">GitHub</a>. Feel free to open a PR to add a Podcast or a Category. The project only has a few python files, the main script is <code>script.py</code> which actually creates the home and the podcast list pages along with the episode pages. The <code>src</code> folder contains some extra bits of scripts like creating the categories and category podcast list pages. Also, it has certain config files like <code>runtime.txt</code>, <code>requirements.txt</code>, and so on. Finally, there is the <code>podlist.json</code> for the list of podcasts and <code>categorylist.json</code> for the categories of podcasts.</p>
<h3>Core Script Snippet</h3>
<p>The python script looks a lot bigger than the below snippet but it is doing the same process multiple times for different pages. There is also some type checking and tiny details that are added as per the requirement of the templates.</p>
<pre><code class="language-python">import feedparser
from jinja2 import Environment, FileSystemLoader
from pathlib import Path

template_env = Environment(loader=FileSystemLoader(searchpath='./layouts/'))
index_template = template_env.get_template('index.html')
episode_template = template_env.get_template('episode.html')

feed = feedparser.parse(&quot;https://freecodecamp.libsyn.com/rss&quot;)

pod_name = feed['feed']['title']

for i in range(0, len(feed['entries']):
    
    ep_title = feed['entries'][i]['title']
    audio = feed['entries'][i]['links'][1]['href']
    cover_image = feed['entries'][i]['image']['href']
    og_link = feed['entries'][i]['links'][0]['href']

    episode_obj = {}
    episode_obj['title'] = ep_title
    episode_obj['audiolink'] = audio
    episode_obj['cover'] = cover_image
    episode_obj['link'] = og_link

    with open(os.path.join(f&quot;site/{pod_name}/ep/{i}/index.html&quot;), 'w', encoding='utf-8') as ep_file:
         ep_file.write(
            episode_template.render(
            episode = episode_obj
            )
         )
</code></pre>
<p>Above is a simple snippet of the core functionality of the script. It basically takes the RSS Feed <code>URL</code> of the podcast and using <code>feedparser</code> the data is retrieved in the form of a dictionary in Python.</p>
<ul>
<li>Iterate over the <code>feed['entries']</code> which is a list of lengths same as the number of episodes of that podcast, we then assign a set of values like <code>episode title</code>, <code>audio link</code>, <code>cover image</code>, <code>canonical link for the episode</code>, <code>date</code> and so on.</li>
<li>Create a dictionary and store the key value as the mentioned data to access from the template.</li>
<li>Open a file in the structured file format and then parse the <code>episode_obj</code> which is a dictionary to the episode template.</li>
<li>Access the dictionary using jinga2 templating tags.</li>
</ul>
<pre><code class="language-html">&lt;html&gt;
    &lt;head&gt;
        &lt;title&gt;Podevcast&lt;/title&gt;
    &lt;/head&gt;
    &lt;body&gt;
        &lt;h3 class=&quot;ep-title&quot;&gt;{{ episode.title }}&lt;/h3&gt;
        &lt;img src=&quot;{{ episode.cover }}&quot;&gt;
        &lt;a class=&quot;ep-link&quot; href=&quot;{{ episode.link }}&quot;&gt;Episode &lt;/a&gt; 
        &lt;audio controls=&quot;enabled&quot; preload=&quot;none&quot;&gt;
            &lt;source src=&quot;{{ episode.audiolink }}&quot; type=&quot;audio/mpeg&quot;&gt;
        &lt;/audio&gt;   
    &lt;/body&gt;
&lt;/html&gt;
</code></pre>
<p>We can use <code>{{  }}</code> to access any value parsed to the template via the script. Also, we can make use of <code>{% %}</code> to run loops, conditionals, blocks, and other tags in the template.</p>
<p><img src="https://res.cloudinary.com/dgpxbrwoz/image/upload/v1645110268/blogmedia/uwdzcwn07oxhppiptem9.png" alt="Feedparser Illustration"></p>
<p>So, we can see the feed is basically a dictionary that has a key-value pair and further, it can be a nested dictionary or a list as a value of a key. As in the case of <code>feed['entries']</code> is a list with the length of the number of episodes of a podcast. And in the script, we use various keys to access various components, obviously, this requires a bit of exploration of the dictionary initially but it becomes easy thereafter to automate using Python.</p>
<h3>Episode List</h3>
<p>Currently, the episodes are added using the JSON file. It is not that user-friendly but still not a big task to simply add a link in a file. This is a #TODO that will require some external tooling to integrate into the webpage to ask for a form to submit a new Podcast.</p>
<pre><code class="language-json">{
    &quot;Command Line Heroes&quot;: &quot;https://feeds.pacific-content.com/commandlineheroes&quot;,
    &quot;Python Podcast__init__&quot;: &quot;https://www.pythonpodcast.com/feed/mp3/&quot;,
    &quot;Real Python Podcast&quot;: &quot;https://realpython.com/podcasts/rpp/feed&quot;,
    &quot;The freeCodeCamp Podcast&quot;: &quot;https://freecodecamp.libsyn.com/rss&quot;,
    &quot;CodeNewbie&quot;: &quot;http://feeds.codenewbie.org/cnpodcast.xml&quot;,
    &quot;Linux For Everyone&quot;: &quot;https://feeds.fireside.fm/linuxforeveryone/rss&quot;,
    &quot;JavaScript Jabber&quot; : &quot;https://feeds.fireside.fm/javascriptjabber/rss&quot;
}
</code></pre>
<p>The process requires a manual test to validate a given RSS Feed as not all feeds are generated the same way and thus there are a few exceptions that need to be sorted out manually. For example, the Python Podcast doesn't have a cover image parsed into the RSS Feed, so there needs to be a check for it in the script and also in the template to restrict parsing and displaying the cover image link.</p>
<h3>Episode Categories</h3>
<p>This is also a JSON file that holds the keys as the category and the value as a list of episode names (strictly the name from <code>feed['feed']['title']</code>). There needs to be a human decision to be taken to add the podcast into a specific category.</p>
<pre><code class="language-json">{
   &quot;Python&quot;:[
      &quot;Talk Python To Me&quot;,
      &quot;The Python Podcast.__init__&quot;,
      &quot;The Real Python Podcast&quot;,
      &quot;Python Bytes&quot;
   ],
   &quot;Javascript&quot;:[
      &quot;Full Stack Radio&quot;,
      &quot;JavaScript Jabber&quot;
   ],
   &quot;Linux&quot;:[
      &quot;Command Line Heroes&quot;,
      &quot;LINUX Unplugged&quot;,
      &quot;The Linux Cast&quot;,
      &quot;Linux For Everyone&quot;
   ],
   &quot;Data Science&quot;:[
      &quot;DataFramed&quot;,
      &quot;Data Skeptic&quot;,
      &quot;The Banana Data Podcast&quot;
   ],
   &quot;Dev Ops&quot;:[
      &quot;DevOps Cafe Podcast&quot;,
      &quot;Arrested DevOps&quot;,
      &quot;Pulling the Strings&quot;,
      &quot;Azure DevOps Podcast&quot;,
      &quot;DevOps and Docker Talk&quot;
   ]
}
</code></pre>
<p>Though the JSON file is managed manually the generation of the categories is automated. Please feel to add other categories of your choice.</p>
<h2>What's Coming?</h2>
<p>Certain features like adding podcast using a form, adding more podcasts, and categories for sure. Though what looks a bit cloudy in my opinion is adding accessibility links to music players because the RSS feed doesn't contain direct links to them. Though I still to explore and find out if it can be obtained from the feed itself.</p>
<ul>
<li>Search box for searching podcasts</li>
<li>Accessible Links to other platforms (Spotify, Itunes, etc)</li>
<li>More depth in categories (Languages/Frameworks/Niche-specific podcasts)</li>
</ul>
<p>I'll add these features after checking the feasibility of the ideas and the response from the community after releasing them.</p>
<h2>Final Words</h2>
<p>This project wouldn't have existed without this Hackathon as it gives a deadline to finish and hope to win something. Specially thanks to Hashnode and Netlify for organizing such a great opportunity in the form of a hackathon. Also, the maintainers of Python libraries like feedparser and jinja. The project would have been impossible without them.</p>
<p>If you like the project please give it a star on <a href="https://github.com/Mr-Destructive/podevcast">GitHub</a>. Have any feedback? Please let me know in the comments or on <a href="https://twitter.com/MeetGor21">Twitter</a>.  Thank you for reading, Hope you have a good time using Podevcast. Happy Coding :)</p>
