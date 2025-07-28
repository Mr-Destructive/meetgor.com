{
  "title": "Techstructive Weekly #3",
  "status": "published",
  "slug": "techstructive-weekly-3",
  "date": "2025-04-05T12:33:25Z"
}

<h2>Techstructive Weekly #3</h2>
<p> The week was really tiring for me for some reasons, was not able to work on the side projects. I actually spent some time with my family and cousin who came over the weeks for festivals. However, on the work side of things, I was able to learn a few things about NGINX, pdf parsing and product specific features.</p>
<p>The upcoming week(end) will be actually fun, as I finally have some free time to write some blogs, hopefully will be able to have some things to showcase and share next week about the things I explored or failed at those things.</p>
<p>Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.</p>
<h2>Quote of the week</h2>
<blockquote>
<p>&quot;The secret to satisfaction is to find meaning in the journey, not just the destination.&quot;<br>
- Unknown</p>
</blockquote>
<p>Satisfaction is important in life, you need to understand the difference between your needs and wants. You absolutely have to give your best to satisfy your needs, but don't go too hard on your wants, understand your intentions and the feelings of others as well. Don't get too greedy on reaching the top, its should be about the process, about the art, about the journey, embrace it and focus on that instead.</p>
<hr>
<h2>Read</h2>
<ul>
<li>
<p>A Pragmatic Workflow for Technical Writing:</p>
<p>This article actually gave me the clarity to start my writing journey again, I seem to have lost touch with writing articles over the last few months, need to get back to the old routine of 1 week 1 article.</p>
</li>
</ul>
<p>![](https://substack-post-media.s3.amazonaws.com/public/images/c2b621db-62a6-452b-b433-ccc0f5c4f0b3_1024x1024.png align=&quot;left&quot;)</p>
<p><a href="https://blog.apiad.net/p/a-pragmatic-workflow-for-technical?utm_source=substack&amp;utm_campaign=post_embed&amp;utm_medium=web">Mostly Harmless Ideas</a></p>
<p><a href="https://blog.apiad.net/p/a-pragmatic-workflow-for-technical?utm_source=substack&amp;utm_campaign=post_embed&amp;utm_medium=web">A Pragmatic Workflow for Technical Writing</a></p>
<p><a href="https://blog.apiad.net/p/a-pragmatic-workflow-for-technical?utm_source=substack&amp;utm_campaign=post_embed&amp;utm_medium=web">Read more</a></p>
<p><a href="https://blog.apiad.net/p/a-pragmatic-workflow-for-technical?utm_source=substack&amp;utm_campaign=post_embed&amp;utm_medium=web">4 months ago · 16 likes · 6 comments · Alejandro Piad Morffis</a></p>
<ul>
<li>
<p>20 years of blogging: <a href="https://jeena.net/20-years-blogging">https://jeena.net/20-years-blogging</a></p>
<p>This shows how one must approach blogging as a developer, just be selfish, write for yourself first, if it helps others, that is the bonus and the reward for you. Don't make your expectations too high, the rewards come over time, make a routine, just dump your thoughts if feeling a bit overwhelmed, it will help.</p>
</li>
</ul>
<h2>Watched</h2>
<ul>
<li>
<p>NGINX Tutorial</p>
</li>
<li>
<p>This gave me a good refresher and a bit more concise explanation of most of the capabilities of NGINX, I am running the Nginx Survival guide blog series, which I will explain the most needed and common uses of NGINX.</p>
</li>
<li>
<p>RabitMQ from <a href="http://Boot.dev">Boot.dev</a>:</p>
</li>
<li>
<p>The video demonstrates the basic architecture for the usage of RabbitMQ for developers. If you are new to message brokers and want to learn about backend systems, I highly recommend watching the clips of <a href="http://boot.dev">boot.dev</a>.</p>
</li>
<li>
<p>Problem with Tutorials: <a href="https://youtube.com/shorts/ZFi-LTpUGHA?si=qt2i80-68I9CwABB">https://youtube.com/shorts/ZFi-LTpUGHA?si=qt2i80-68I9CwABB</a></p>
<p>I can relate as a technical writer, how much stuff is lost while writing the blog. So many examples discarded, so many attempts hidden, so many words deleted. The final guide/tutorial looks so perfect that it might overwhelm the reader/listener that how is it possible to get everything right in the first place, in reality it is a result of failed attempts and trail and error.</p>
</li>
<li>
<p>Tweet from Bill Kennedy about mindset of developers after the advent of AI assistants: <a href="https://x.com/goinggodotnet/status/1822298432813375760%EF%BF%BCReally">https://x.com/goinggodotnet/status/1822298432813375760<br>
Really</a>, this hit me hard, as a developer, we need to focus on solving architectures and problems and not limit ourselves in minor code fixes and feature implementations.</p>
</li>
</ul>
<h2>Learnt</h2>
<ul>
<li>
<p>If you want to quickly get a JSON response from a locally running API, never use <a href="http://subprocess.run"><code>subprocess.run</code></a> in python, this actually encodes the json response in a way that becomes almost impossible to get back, especially if the json response structure is a bit complicated. I was a bit in a hurry and lazy, so just prompted GPT to get me a script to call the api and store the response in a json file, but what happened was while returning the response from the stdout, the json string was encoded, and it became a mess to decode back. Always use requests or the native libraries in python. Do not rely on stdout responses, as it could be very different the way shells might interpret data from your native applications.</p>
</li>
<li>
<p>There is no apparent way of restoring files of VS Code file explorer after deleting, like yeah CTRL + Z would work if you just deleted. But after a day, I realized I needed the files back, I tried to into <code>/tmp</code> folder, there were files I deleted earlier this week, but not from VS Code apparently. Really a bummer. Let me know if you know the place where VS Code stores deleted files, if it does.</p>
</li>
</ul>
<h2>Tech News</h2>
<ul>
<li>
<p>Golang 1.23 released this week: <a href="https://tip.golang.org/doc/go1.23">https://tip.golang.org/doc/go1.23</a></p>
<p>This looks like a big release with some alternate approaches getting added to the language, which are kind of double edged swords from most of the professionals opinion I noticed.</p>
</li>
<li>
<p>X releases Grok AI 2: <a href="https://x.ai/blog/grok-2%EF%BF%BCThis">https://x.ai/blog/grok-2<br>
This</a> actually caused some trend on twitter (ok X) of floods of images generated by Grok AI.</p>
</li>
</ul>
<p>For more news, follow the Hackernewsletter:</p>
<p><a href="https://mailchi.mp/hackernewsletter/712">https://mailchi.mp/hackernewsletter/712</a></p>
<hr>
<p>That’s it from this week, I hope you did well this week, and have a happy week and weekend ahead!</p>
<p>Thank you for reading, let’s catch up in the next week.</p>
<p>Happy Coding :)</p>
<p>Thanks for reading Techstructive Weekly! Subscribe for free to receive new posts and support my work.</p>
