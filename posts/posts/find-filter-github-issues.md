{
  "title": "Filter and Find an Issue on GitHub",
  "status": "published",
  "slug": "find-filter-github-issues",
  "date": "2025-04-05T12:33:54Z"
}

<h2>Introduction</h2>
<p>Are you stuck in finding an open-source project to contribute to?</p>
<p>We will see how you can pick up an issue on GitHub appropriate as per your preferences of languages, labels, complexity, and thus you can find a Community or a project to work and continue with further contributions.</p>
<p>This process might not be as efficient but is quite helpful for beginners or people getting started to contributing to Open Source.</p>
<h2>Understand the search bar</h2>
<p>I assume you have your GitHub account already created. If not go ahead at <a href="https://github.com/join">Github</a> and create one. On the Home page, you can easily navigate to the <code>Issues</code> tab and you will see something like this:</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631190578909/UBpq3rb0H.png" alt="Issues tab"></p>
<p>Now, you won't find any issues if you haven't created any. But if you look at the search bar, you will find the reason why it is empty or why there are only the issues that you have created. You will see that in the search bar there is a filter called <code>author:Username</code>, which filters the issues which are created by you. You definitely don't want this as you want to search and find other issues by other people/communities. So, simply remove the text <code>author:Username</code> from the search bar. Keep rest as it is for now. Now if you press enter after removing the author filter, you will see all the issues on GitHub.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631185853484/e0PyTbgip.png" alt="Issues removed author"></p>
<p>There will be a ton of them, very random in terms of programming languages, frameworks, projects, difficulty, type, etc. they are basically the issues created recently on GitHub.</p>
<p>In the next section, we will see how to filter those issues as per the programming languages/tools to which you might like to contribute to.</p>
<h2>Add languages</h2>
<p>We can add filters to the issues as <code>language:name</code>, this will filter all the Issues which have the languages in their codebase.</p>
<p>For Example:</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631190679194/8Od1tsdKp.png" alt="Issues language filter"></p>
<p>Here, I have filtered the issues which have language as <code>python</code>, you can use any language/tool you might want and would love to find some interesting projects to contribute and learn from.</p>
<p>If you want to search by multiple programming languages you can separate the names of those programming languages by a comma <code>,</code>.</p>
<p>You can also separate programming languages with space and enclosing all of them under double quotes <code>&quot;&quot;</code>.</p>
<p>For Example:</p>
<p>Let's search for issues with C, C++, and Java as their programming languages, we can use <code>language:c,cpp,java</code> or <code>language:&quot;c cpp java&quot;</code></p>
<p>The above filter will give out all the issues which are created from programming languages either C/C++/Java.</p>
<p>You can find more filter options on the <a href="https://docs.github.com/en/github/searching-for-information-on-github/searching-on-github/searching-issues-and-pull-requests">GitHub docs</a>.</p>
<h2>Add labels</h2>
<p>You can find issues as per labels marked on them, many issues have a label marked on them to improve their visibility and meta-information about the issue.</p>
<p>We have some labels which GitHub has created already for common scenarios in projects.</p>
<ol>
<li><code>bug</code></li>
<li><code>documentation</code></li>
<li><code>duplicate</code></li>
<li><code>enhancement</code></li>
<li><code>good first Issue</code></li>
<li><code>help wanted</code></li>
<li><code>invalid</code></li>
<li><code>question</code></li>
<li><code>wontfix</code></li>
</ol>
<p>We can even create our own labels by providing the label name and a description.</p>
<p>To search for labels, you can use <code>label:name of the label</code>. You can any of the above 9 label tags or any other tag name that you think is popular other than those 9.</p>
<p>You would have to use double quotes (<code>&quot;&quot;</code>) to add certain labels with multiple words like <code>good first issue</code> or <code>help wanted</code>.</p>
<p>For example:</p>
<p>If you search for <code>label:&quot;good first issue&quot;</code>, you will get all of the issues(newest first) which have a label <code>good first issues</code> tagged on them.</p>
<p>Similarly, for multiple issues, you can add comma-separated labels as well. Just like <code>label:bug,&quot;good first issue&quot;</code> will search for either <code>bug</code>, <code>good first issue</code> or both.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631190841185/vrYTLoaaNu.png" alt="Issues label"></p>
<h2>More Sorting Options</h2>
<p>In the rightmost part of the search bar, in the Sort button, you can click on there and find a couple of options like: <code>newest</code>, <code>oldest</code>, <code>least commented</code>, <code>recently updated</code>, and so on. If you click on any of them you will see the changes reflected on the list of issues as well as the search bar.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631189621396/jO58HkYxH.png" alt="Issues sort"></p>
<p><strong>After this the stage is yours, you can look at any issue and Understand its objective, then ask yourself can you solve this issue? If yes then read the contribution guidelines, and the rest is sheer skills like git, programming, documentation, etc.</strong></p>
<h2>Conclusion</h2>
<p>Now you can go ahead and start applying the filters on issues and make some contributions to Open-Source on GitHub. We covered some methods and tricks to find and filter out the issues on GitHub based on the programming languages/tools and the labels attached to them.</p>
<p>This technique can be good for beginners as well as people who want to find quick issues to solve. Feel free to explore and try out different filters and find the issue you are confident to work on. Good Luck!
Happy Coding :)</p>
