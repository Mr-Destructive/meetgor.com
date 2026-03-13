---


























type: links
title: 'How React works behind the scenes (under the hood)'
date: 2025-08-16
slug: how-react-works-behind-the-scenes-under-the-hood
tags:
  - links
link: 'https://www.deepintodev.com/blog/how-react-works-behind-the-scenes?ref=dailydev'
status: published
description: 'How React works behind the scenes (under the hood)'
source: newsletter
newsletter: 2025-08-16-techstructive-weekly-55


























---


## Commentary

- Wow! What a post, I can happily say, I know something about React now. There are a lot of moving parts. I thought it was just one step compilation, but man there are layers of compilation happening in various formats/data layers I think.
- The first one is JSON, the components and the apps are converted into JSON notation and then used to construct the DOM, that is fascinating, based on the reference to different components, it can decide which elements to render or re-render. There is a graph created so that it becomes easier to distinguish and make a hierarchy of the app. The virtual DOM as called is like the graph that helps in re-rendering, but the heavy lifting is done by the Fiber tree (which is a lower level abstraction) that does the actual replacement or rendering technique algorithms.
- The process is so quick that the magic is not even noticable, but yes that is one rabbit hole to dig into.
