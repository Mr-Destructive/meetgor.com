---
layout: page
title: Tags 
permalink: /tag/
---

<div id="archives">
{% for category in site.categories %}
  <div class="archive-group">
    {% capture category_name %}{{ category | first }}{% endcapture %}
    <div id="#{{ category_name | slugize }}"></div>
    <p></p>
<details>
	<summary>
		<div id="#{{ category_name | slugize }}"></div>
    		<h3 class="category-head">#{{ category_name }}</h3>
    		<a name="{{ category_name | slugize }}"></a>
	</summary>
    		{% for post in site.categories[category_name] %}
			<article class="archive-item">
				<h4><a href="{{ post.url | relative_url }}">{{post.title}}</a></h4>
			</article>
    		{% endfor %}
</details>
  </div>
{% endfor %}
</div>
