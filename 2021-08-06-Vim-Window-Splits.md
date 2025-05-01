---
article_html: '<h2>Introduction</h2>

  <p>Have you ever been stuck in Vim opening multiple files within a single window?
  Didn''t anyone tell you, you can create multiple windows and split them within a
  single tab. Definitely, the window splits will be in separate buffers. In this way
  you can create multiple windows inside of a single Tab, what are Tabs? You can learn
  some basics about it from my previous article about  <a href="https://mr-destructive.github.io/techstructive-blog/vim/2021/08/03/Vim-Tabs.html">Tabs
  in Vim</a>. We can either create Vertical or Horizontal splits within the window
  making it flexible to work with multiple files in Vim. This article will look into
  the creation, navigation, closing, and rearrangement of Window Splits.</p>

  <h2>Creating a Window Split</h2>

  <p>Creating Window splits is quite straightforward. You should keep in mind the
  following things though:</p>

  <ul>

  <li>You can create a horizontal or a vertical split within a window.</li>

  <li>Creating a Split either vertically or horizontally can shorten the current window''s
  size, making it equally spaced.</li>

  </ul>

  <p>Let''s take a look at creating the vertical and horizontal splits one by one:</p>

  <h3>Vertical Splits</h3>

  <p>Vertical Split as the name suggests, it will split the current window into <strong>two
  halves vertically</strong> or a <strong>standing split between two windows</strong>.</p>

  <p>The below image clearly shows a vertical split between two windows. Here we are
  splitting a single window into two windows. We can also think it of in splitting
  the window from left to right.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628232885853/xtBgWb-Yg.png"
  alt="image.png" /></p>

  <p>To create a vertical split, you can use <code>:vsp</code> or <code>:vsplit</code>
  to create a split of the same file/ blank file.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628233753115/4seJbY-h9.gif"
  alt="vsp.gif" /></p>

  <p>If you already have a file open, it will open the same file in the split as long
  as you don''t specify which file to open. You can specify the name of the file after
  the command <code>:vsp filename</code> or <code>:vsplit filename</code></p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628233871768/B3D_3NNGo.gif"
  alt="vsp.gif" /></p>

  <p>It''s not like that you can create only a single split, you can create multiple
  vertical splits. That can get pretty wild pretty quickly.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628234228391/vmJxW5HOo.png"
  alt="image.png" /></p>

  <p>In the above screenshot, I have created 5 vertical splits from a single window,
  so making them equally wide and evenly spaced. This might not be useful every time
  but can get quite handy in some tricky situations.</p>

  <h3>Horizontal Splits</h3>

  <p>Similar to Vertical splits, we have horizontal Splits indicating to split from
  top to bottom. We can <strong>split a single window into two halves horizontally</strong>
  or a <strong>sleeping split between the windows</strong>.</p>

  <p>The below image clearly shows a horizontal split between two windows. Here we
  are splitting a single window into two windows. We can also think it of in splitting
  the window from top to bottom.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628233063400/5PVdEsGHZ.png"
  alt="image.png" /></p>

  <p>To create a horizontal split, you can use <code>:sp</code> or <code>:split</code>
  to create a horizontal split of the same file/ blank file. This will create a blank
  file inside a horizontal split.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628235156757/ckfDxh-1D.gif"
  alt="sp.gif" /></p>

  <p>Similar to the vertical splits, you can open files by creating the split. You
  can use the command <code>:sp filename</code> or <code>:split filename</code> to
  create the horizontal split between the windows and opening a specified file in
  it.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628235452142/eVGrEZmHVK.gif"
  alt="sp.gif" /></p>

  <p>Again as seen in the vertical split, we can create as many splits as we like.
  This looks very ugly but who knows when you may need this. We can basically create
  Splits in any order of vertical or horizontal, we''ll see it in the later section
  to the same.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628235679425/9dtK5TV6G.png"
  alt="image.png" /></p>

  <h3>Vertical and Horizontal Splits Together</h3>

  <p>We can also create Vertical split and horizontal splits independently like we
  are not restricted to only creating the only kinds of splits at a time. You can
  create a vertical split once and then can create a horizontal split within the vertical
  split created, in this way we can make good use of both the splits as per our needs.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628236573469/7FOJIgP-z.png"
  alt="image.png" /></p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628236677571/CCeVePLkp.png"
  alt="image.png" /></p>

  <h2>Moving around Splits</h2>

  <p>Now let''s talk about the navigation around these splits, it''s again a simple
  thing. We can use <code>Ctrl+w</code> as our primary command along with Vim navigation
  to move around the splits. Like for example <code>Ctrl + w + h</code> will take
  you to the left split to your current open window. You can also use <code>Ctrl+w</code>**
  twice** to hop around to the next window split in a repeating loop.</p>

  <p>So, we can use:</p>

  <p><code>Ctrl + w </code>+ <code>w</code> -&gt; Switch to the next split (in the
  order of creation of those splits).</p>

  <p><code>Ctrl + w </code> + <code>h</code> -&gt; Move to the left split.</p>

  <p><code>Ctrl + w </code> + <code>l</code> -&gt; Move to the right split.</p>

  <p><code>Ctrl + w </code> + <code>j</code>  -&gt; Move to the below split.</p>

  <p><code>Ctrl + w </code> + <code>k</code>  -&gt; Move to a upper split.</p>

  <p>These commands might be good enough to pull you through any splits from anywhere,
  it just becomes easy to use Vim navigation inside of these. You can use arrow keys
  if you are stuck somewhere but using Vim key bindings will work out of the box.</p>

  <h2>Creating Splits using Keyboard shortcuts</h2>

  <p>If you think you are wasting time going into the command mode and typing the
  commands to create splits, well there are some shortcuts for you.</p>

  <p><code>Ctrl + w</code> + <code>v</code> -&gt; Create a vertical split.</p>

  <p><code>Ctrl + w </code>+ <code>s</code> -&gt; Create a horizontal split.</p>

  <p>This will open the split with the current file in the original window, so if
  you need to change the file, you can use the edit command (<code>:e filename</code>)
  inside the split.</p>

  <h2>Rearranging the Window Splits</h2>

  <p>If you have a  specific set of splits of a kind open, you can rotate between
  those. Like for example, if we have a horizontal split, you can rotate the split
  to move the upper split down and below split up.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628240624930/HBIKPummz.gif"
  alt="cwr.gif" /></p>

  <p>From the above illustration, we can see we rotated(swapped) the splits. We can
  also use certain commands to arrange the splits into appropriate positions according
  to the user.</p>

  <p><code>Ctrl + w</code> + <code>r</code> -&gt; Swap the two splits(either horizontal
  or vertical)</p>

  <p><code>Ctrl + w</code> + <code>H</code> -&gt; Move the split to the left ( <code>Ctrl
  + w</code> + <code>Shift + h</code>).</p>

  <p><code>Ctrl + w</code> + <code>J</code> -&gt; Move the split down ( <code>Ctrl
  + w</code> + <code>Shift + j</code>).</p>

  <p><code>Ctrl + w</code> + <code>K</code> -&gt; Move the split up ( <code>Ctrl +
  w</code> + <code>Shift + k</code>). .</p>

  <p><code>Ctrl + w</code> + <code>L</code> -&gt; Move the split to the right ( <code>Ctrl
  + w</code> + <code>Shift + l</code>).</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628242295049/_HadBQPrs.gif"
  alt="spr.gif" /></p>

  <p>We can see that we were able to swap the splits to a location suitable according
  to our wish. We can definitely switch the splits internally as well.</p>

  <p>We can resize the splits as per the requirement and remove the equality in those
  splits. We have a couple of options to do this:</p>

  <p><code>Ctrl +w</code> + <code>+</code> -&gt;  Increase the height of the current
  split.</p>

  <p><code>Ctrl +w</code> + <code>-</code> -&gt;  Decrease the height of the current
  split.</p>

  <p><code>Ctrl +w</code> + <code>&gt;</code> -&gt;  Increase the width of the current
  split.</p>

  <p><code>Ctrl +w</code> +  <code>&lt;</code> -&gt;  Decrease the width of the current
  split.</p>

  <p><code>Ctrl +w</code> + <code>=</code> -&gt;  Make the splits equal in width and
  height.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628243827459/nL3mus88d.gif"
  alt="spr.gif" /></p>

  <p>We can also use <code>:resize {number}</code> to manually set the size of the
  horizontal split and <code>:vertical resize {number}</code> to manually set the
  size of the vertical split. This is really risky, like if you know what you are
  doing then it''s totally fine. The commands demonstrated earlier are really adjustable
  and user-friendly.</p>

  <p>We can also use <code>Ctrl + w</code> + <code>_</code> to minimize all the window
  split except the current one.</p>

  <h2>Closing the Splits</h2>

  <p>Now after doing all sorts of wizardry with the window splits the finishing touch
  is to close those splits after use. We can surely use <code>:q</code> to close the
  current window split but that is not intuitive to enter each split and manually
  close all of them, we can use:</p>

  <p><code>Ctrl +w</code> + <code>c</code> -&gt;  Close the current split.</p>

  <p><code>Ctrl +w</code> + <code>o</code> -&gt;  Quit all other splits except the
  current one.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628244089101/Ebdup7wNE.gif"
  alt="closesp.gif" /></p>

  <h2>Conclusion</h2>

  <p>So, we have seen the basics of using Window Splits in Vim. Please let me know
  if there is some important point that is missing. Thank you for reading through
  here. We have seen how to create, navigate, rearrange, closing and other basic stuff
  related to Window splits in Vim.</p>

  <h3>References:</h3>

  <ul>

  <li><a href="https://sodocumentation.net/vim/topic/1705/split-windows">Sodocumentation</a></li>

  <li><a href="https://linuxhint.com/vim_split_screen/">Linux Hint Vim Split Screen</a></li>

  <li><a href="https://gist.github.com/Starefossen/5957088">Starefossen</a></li>

  </ul>

  <p>Hopefully, this might have given you some good idea to deal with Vim in windows
  splits. Happy Coding and Viming :)</p>

  '
cover: ''
date: 2021-08-06
datetime: 2021-08-06 00:00:00+00:00
description: Have you ever been stuck in Vim opening multiple files within a single
  window? Didn Creating Window splits is quite straightforward. You should keep in
  mind the
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-08-06-Vim-Window-Splits.md
html: '<h2>Introduction</h2>

  <p>Have you ever been stuck in Vim opening multiple files within a single window?
  Didn''t anyone tell you, you can create multiple windows and split them within a
  single tab. Definitely, the window splits will be in separate buffers. In this way
  you can create multiple windows inside of a single Tab, what are Tabs? You can learn
  some basics about it from my previous article about  <a href="https://mr-destructive.github.io/techstructive-blog/vim/2021/08/03/Vim-Tabs.html">Tabs
  in Vim</a>. We can either create Vertical or Horizontal splits within the window
  making it flexible to work with multiple files in Vim. This article will look into
  the creation, navigation, closing, and rearrangement of Window Splits.</p>

  <h2>Creating a Window Split</h2>

  <p>Creating Window splits is quite straightforward. You should keep in mind the
  following things though:</p>

  <ul>

  <li>You can create a horizontal or a vertical split within a window.</li>

  <li>Creating a Split either vertically or horizontally can shorten the current window''s
  size, making it equally spaced.</li>

  </ul>

  <p>Let''s take a look at creating the vertical and horizontal splits one by one:</p>

  <h3>Vertical Splits</h3>

  <p>Vertical Split as the name suggests, it will split the current window into <strong>two
  halves vertically</strong> or a <strong>standing split between two windows</strong>.</p>

  <p>The below image clearly shows a vertical split between two windows. Here we are
  splitting a single window into two windows. We can also think it of in splitting
  the window from left to right.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628232885853/xtBgWb-Yg.png"
  alt="image.png" /></p>

  <p>To create a vertical split, you can use <code>:vsp</code> or <code>:vsplit</code>
  to create a split of the same file/ blank file.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628233753115/4seJbY-h9.gif"
  alt="vsp.gif" /></p>

  <p>If you already have a file open, it will open the same file in the split as long
  as you don''t specify which file to open. You can specify the name of the file after
  the command <code>:vsp filename</code> or <code>:vsplit filename</code></p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628233871768/B3D_3NNGo.gif"
  alt="vsp.gif" /></p>

  <p>It''s not like that you can create only a single split, you can create multiple
  vertical splits. That can get pretty wild pretty quickly.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628234228391/vmJxW5HOo.png"
  alt="image.png" /></p>

  <p>In the above screenshot, I have created 5 vertical splits from a single window,
  so making them equally wide and evenly spaced. This might not be useful every time
  but can get quite handy in some tricky situations.</p>

  <h3>Horizontal Splits</h3>

  <p>Similar to Vertical splits, we have horizontal Splits indicating to split from
  top to bottom. We can <strong>split a single window into two halves horizontally</strong>
  or a <strong>sleeping split between the windows</strong>.</p>

  <p>The below image clearly shows a horizontal split between two windows. Here we
  are splitting a single window into two windows. We can also think it of in splitting
  the window from top to bottom.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628233063400/5PVdEsGHZ.png"
  alt="image.png" /></p>

  <p>To create a horizontal split, you can use <code>:sp</code> or <code>:split</code>
  to create a horizontal split of the same file/ blank file. This will create a blank
  file inside a horizontal split.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628235156757/ckfDxh-1D.gif"
  alt="sp.gif" /></p>

  <p>Similar to the vertical splits, you can open files by creating the split. You
  can use the command <code>:sp filename</code> or <code>:split filename</code> to
  create the horizontal split between the windows and opening a specified file in
  it.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628235452142/eVGrEZmHVK.gif"
  alt="sp.gif" /></p>

  <p>Again as seen in the vertical split, we can create as many splits as we like.
  This looks very ugly but who knows when you may need this. We can basically create
  Splits in any order of vertical or horizontal, we''ll see it in the later section
  to the same.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628235679425/9dtK5TV6G.png"
  alt="image.png" /></p>

  <h3>Vertical and Horizontal Splits Together</h3>

  <p>We can also create Vertical split and horizontal splits independently like we
  are not restricted to only creating the only kinds of splits at a time. You can
  create a vertical split once and then can create a horizontal split within the vertical
  split created, in this way we can make good use of both the splits as per our needs.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628236573469/7FOJIgP-z.png"
  alt="image.png" /></p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628236677571/CCeVePLkp.png"
  alt="image.png" /></p>

  <h2>Moving around Splits</h2>

  <p>Now let''s talk about the navigation around these splits, it''s again a simple
  thing. We can use <code>Ctrl+w</code> as our primary command along with Vim navigation
  to move around the splits. Like for example <code>Ctrl + w + h</code> will take
  you to the left split to your current open window. You can also use <code>Ctrl+w</code>**
  twice** to hop around to the next window split in a repeating loop.</p>

  <p>So, we can use:</p>

  <p><code>Ctrl + w </code>+ <code>w</code> -&gt; Switch to the next split (in the
  order of creation of those splits).</p>

  <p><code>Ctrl + w </code> + <code>h</code> -&gt; Move to the left split.</p>

  <p><code>Ctrl + w </code> + <code>l</code> -&gt; Move to the right split.</p>

  <p><code>Ctrl + w </code> + <code>j</code>  -&gt; Move to the below split.</p>

  <p><code>Ctrl + w </code> + <code>k</code>  -&gt; Move to a upper split.</p>

  <p>These commands might be good enough to pull you through any splits from anywhere,
  it just becomes easy to use Vim navigation inside of these. You can use arrow keys
  if you are stuck somewhere but using Vim key bindings will work out of the box.</p>

  <h2>Creating Splits using Keyboard shortcuts</h2>

  <p>If you think you are wasting time going into the command mode and typing the
  commands to create splits, well there are some shortcuts for you.</p>

  <p><code>Ctrl + w</code> + <code>v</code> -&gt; Create a vertical split.</p>

  <p><code>Ctrl + w </code>+ <code>s</code> -&gt; Create a horizontal split.</p>

  <p>This will open the split with the current file in the original window, so if
  you need to change the file, you can use the edit command (<code>:e filename</code>)
  inside the split.</p>

  <h2>Rearranging the Window Splits</h2>

  <p>If you have a  specific set of splits of a kind open, you can rotate between
  those. Like for example, if we have a horizontal split, you can rotate the split
  to move the upper split down and below split up.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628240624930/HBIKPummz.gif"
  alt="cwr.gif" /></p>

  <p>From the above illustration, we can see we rotated(swapped) the splits. We can
  also use certain commands to arrange the splits into appropriate positions according
  to the user.</p>

  <p><code>Ctrl + w</code> + <code>r</code> -&gt; Swap the two splits(either horizontal
  or vertical)</p>

  <p><code>Ctrl + w</code> + <code>H</code> -&gt; Move the split to the left ( <code>Ctrl
  + w</code> + <code>Shift + h</code>).</p>

  <p><code>Ctrl + w</code> + <code>J</code> -&gt; Move the split down ( <code>Ctrl
  + w</code> + <code>Shift + j</code>).</p>

  <p><code>Ctrl + w</code> + <code>K</code> -&gt; Move the split up ( <code>Ctrl +
  w</code> + <code>Shift + k</code>). .</p>

  <p><code>Ctrl + w</code> + <code>L</code> -&gt; Move the split to the right ( <code>Ctrl
  + w</code> + <code>Shift + l</code>).</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628242295049/_HadBQPrs.gif"
  alt="spr.gif" /></p>

  <p>We can see that we were able to swap the splits to a location suitable according
  to our wish. We can definitely switch the splits internally as well.</p>

  <p>We can resize the splits as per the requirement and remove the equality in those
  splits. We have a couple of options to do this:</p>

  <p><code>Ctrl +w</code> + <code>+</code> -&gt;  Increase the height of the current
  split.</p>

  <p><code>Ctrl +w</code> + <code>-</code> -&gt;  Decrease the height of the current
  split.</p>

  <p><code>Ctrl +w</code> + <code>&gt;</code> -&gt;  Increase the width of the current
  split.</p>

  <p><code>Ctrl +w</code> +  <code>&lt;</code> -&gt;  Decrease the width of the current
  split.</p>

  <p><code>Ctrl +w</code> + <code>=</code> -&gt;  Make the splits equal in width and
  height.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628243827459/nL3mus88d.gif"
  alt="spr.gif" /></p>

  <p>We can also use <code>:resize {number}</code> to manually set the size of the
  horizontal split and <code>:vertical resize {number}</code> to manually set the
  size of the vertical split. This is really risky, like if you know what you are
  doing then it''s totally fine. The commands demonstrated earlier are really adjustable
  and user-friendly.</p>

  <p>We can also use <code>Ctrl + w</code> + <code>_</code> to minimize all the window
  split except the current one.</p>

  <h2>Closing the Splits</h2>

  <p>Now after doing all sorts of wizardry with the window splits the finishing touch
  is to close those splits after use. We can surely use <code>:q</code> to close the
  current window split but that is not intuitive to enter each split and manually
  close all of them, we can use:</p>

  <p><code>Ctrl +w</code> + <code>c</code> -&gt;  Close the current split.</p>

  <p><code>Ctrl +w</code> + <code>o</code> -&gt;  Quit all other splits except the
  current one.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628244089101/Ebdup7wNE.gif"
  alt="closesp.gif" /></p>

  <h2>Conclusion</h2>

  <p>So, we have seen the basics of using Window Splits in Vim. Please let me know
  if there is some important point that is missing. Thank you for reading through
  here. We have seen how to create, navigate, rearrange, closing and other basic stuff
  related to Window splits in Vim.</p>

  <h3>References:</h3>

  <ul>

  <li><a href="https://sodocumentation.net/vim/topic/1705/split-windows">Sodocumentation</a></li>

  <li><a href="https://linuxhint.com/vim_split_screen/">Linux Hint Vim Split Screen</a></li>

  <li><a href="https://gist.github.com/Starefossen/5957088">Starefossen</a></li>

  </ul>

  <p>Hopefully, this might have given you some good idea to deal with Vim in windows
  splits. Happy Coding and Viming :)</p>

  '
image_url: https://cdn.hashnode.com/res/hashnode/image/upload/v1628151057227/gZey9TYHd.png?w=1600&h=840&fit=crop&crop=entropy&auto=compress
long_description: 'Have you ever been stuck in Vim opening multiple files within a
  single window? Didn Creating Window splits is quite straightforward. You should
  keep in mind the following things though: You can create a horizontal or a vertical
  split within a window.'
now: 2025-05-01 18:11:33.313259
path: blog/posts/2021-08-06-Vim-Window-Splits.md
prevnext: null
slug: vim-window-splits
status: published
subtitle: Learning to create, navigate in Window splits in Vim
tags:
- vim
templateKey: blog-post
title: 'Vim: Window Splits'
today: 2025-05-01
---

## Introduction

Have you ever been stuck in Vim opening multiple files within a single window? Didn't anyone tell you, you can create multiple windows and split them within a single tab. Definitely, the window splits will be in separate buffers. In this way you can create multiple windows inside of a single Tab, what are Tabs? You can learn some basics about it from my previous article about  [Tabs in Vim](https://mr-destructive.github.io/techstructive-blog/vim/2021/08/03/Vim-Tabs.html). We can either create Vertical or Horizontal splits within the window making it flexible to work with multiple files in Vim. This article will look into the creation, navigation, closing, and rearrangement of Window Splits.

## Creating a Window Split

Creating Window splits is quite straightforward. You should keep in mind the following things though:
- You can create a horizontal or a vertical split within a window.
- Creating a Split either vertically or horizontally can shorten the current window's size, making it equally spaced.

Let's take a look at creating the vertical and horizontal splits one by one: 

### Vertical Splits

Vertical Split as the name suggests, it will split the current window into **two halves vertically** or a **standing split between two windows**.

The below image clearly shows a vertical split between two windows. Here we are splitting a single window into two windows. We can also think it of in splitting the window from left to right.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1628232885853/xtBgWb-Yg.png)

To create a vertical split, you can use `:vsp` or `:vsplit` to create a split of the same file/ blank file.

![vsp.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1628233753115/4seJbY-h9.gif)

If you already have a file open, it will open the same file in the split as long as you don't specify which file to open. You can specify the name of the file after the command `:vsp filename` or `:vsplit filename`

![vsp.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1628233871768/B3D_3NNGo.gif)

It's not like that you can create only a single split, you can create multiple vertical splits. That can get pretty wild pretty quickly.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1628234228391/vmJxW5HOo.png)

In the above screenshot, I have created 5 vertical splits from a single window, so making them equally wide and evenly spaced. This might not be useful every time but can get quite handy in some tricky situations.

### Horizontal Splits

Similar to Vertical splits, we have horizontal Splits indicating to split from top to bottom. We can **split a single window into two halves horizontally** or a **sleeping split between the windows**. 

The below image clearly shows a horizontal split between two windows. Here we are splitting a single window into two windows. We can also think it of in splitting the window from top to bottom.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1628233063400/5PVdEsGHZ.png)

To create a horizontal split, you can use `:sp` or `:split` to create a horizontal split of the same file/ blank file. This will create a blank file inside a horizontal split.

![sp.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1628235156757/ckfDxh-1D.gif)

Similar to the vertical splits, you can open files by creating the split. You can use the command `:sp filename` or `:split filename` to create the horizontal split between the windows and opening a specified file in it.

![sp.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1628235452142/eVGrEZmHVK.gif)

Again as seen in the vertical split, we can create as many splits as we like. This looks very ugly but who knows when you may need this. We can basically create Splits in any order of vertical or horizontal, we'll see it in the later section to the same.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1628235679425/9dtK5TV6G.png)

### Vertical and Horizontal Splits Together
We can also create Vertical split and horizontal splits independently like we are not restricted to only creating the only kinds of splits at a time. You can create a vertical split once and then can create a horizontal split within the vertical split created, in this way we can make good use of both the splits as per our needs.

![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1628236573469/7FOJIgP-z.png)


![image.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1628236677571/CCeVePLkp.png)

## Moving around Splits

Now let's talk about the navigation around these splits, it's again a simple thing. We can use `Ctrl+w` as our primary command along with Vim navigation to move around the splits. Like for example `Ctrl + w + h` will take you to the left split to your current open window. You can also use `Ctrl+w`** twice** to hop around to the next window split in a repeating loop. 

So, we can use:

`Ctrl + w `+ `w` -> Switch to the next split (in the order of creation of those splits).

`Ctrl + w ` + `h` -> Move to the left split.

`Ctrl + w ` + `l` -> Move to the right split.

`Ctrl + w ` + `j`  -> Move to the below split.

`Ctrl + w ` + `k`  -> Move to a upper split.

These commands might be good enough to pull you through any splits from anywhere, it just becomes easy to use Vim navigation inside of these. You can use arrow keys if you are stuck somewhere but using Vim key bindings will work out of the box.

## Creating Splits using Keyboard shortcuts

If you think you are wasting time going into the command mode and typing the commands to create splits, well there are some shortcuts for you. 

`Ctrl + w` + `v` -> Create a vertical split.

`Ctrl + w `+ `s` -> Create a horizontal split.

This will open the split with the current file in the original window, so if you need to change the file, you can use the edit command (`:e filename`) inside the split. 
 

## Rearranging the Window Splits

If you have a  specific set of splits of a kind open, you can rotate between those. Like for example, if we have a horizontal split, you can rotate the split to move the upper split down and below split up.  

![cwr.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1628240624930/HBIKPummz.gif)

From the above illustration, we can see we rotated(swapped) the splits. We can also use certain commands to arrange the splits into appropriate positions according to the user. 

`Ctrl + w` + `r` -> Swap the two splits(either horizontal or vertical)

`Ctrl + w` + `H` -> Move the split to the left ( `Ctrl + w` + `Shift + h`). 

`Ctrl + w` + `J` -> Move the split down ( `Ctrl + w` + `Shift + j`). 

`Ctrl + w` + `K` -> Move the split up ( `Ctrl + w` + `Shift + k`). . 

`Ctrl + w` + `L` -> Move the split to the right ( `Ctrl + w` + `Shift + l`).


![spr.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1628242295049/_HadBQPrs.gif)

We can see that we were able to swap the splits to a location suitable according to our wish. We can definitely switch the splits internally as well.

We can resize the splits as per the requirement and remove the equality in those splits. We have a couple of options to do this:

`Ctrl +w` + `+` ->  Increase the height of the current split.

`Ctrl +w` + `-` ->  Decrease the height of the current split.

`Ctrl +w` + `>` ->  Increase the width of the current split.

`Ctrl +w` +  `<` ->  Decrease the width of the current split.

`Ctrl +w` + `=` ->  Make the splits equal in width and height.

![spr.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1628243827459/nL3mus88d.gif)

We can also use `:resize {number}` to manually set the size of the horizontal split and `:vertical resize {number}` to manually set the size of the vertical split. This is really risky, like if you know what you are doing then it's totally fine. The commands demonstrated earlier are really adjustable and user-friendly. 

We can also use `Ctrl + w` + `_` to minimize all the window split except the current one.
 
## Closing the Splits

Now after doing all sorts of wizardry with the window splits the finishing touch is to close those splits after use. We can surely use `:q` to close the current window split but that is not intuitive to enter each split and manually close all of them, we can use:

`Ctrl +w` + `c` ->  Close the current split.

`Ctrl +w` + `o` ->  Quit all other splits except the current one. 

![closesp.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1628244089101/Ebdup7wNE.gif)

## Conclusion

So, we have seen the basics of using Window Splits in Vim. Please let me know if there is some important point that is missing. Thank you for reading through here. We have seen how to create, navigate, rearrange, closing and other basic stuff related to Window splits in Vim.

### References:

- [Sodocumentation](https://sodocumentation.net/vim/topic/1705/split-windows)
- [Linux Hint Vim Split Screen](https://linuxhint.com/vim_split_screen/)
- [Starefossen](https://gist.github.com/Starefossen/5957088)

Hopefully, this might have given you some good idea to deal with Vim in windows splits. Happy Coding and Viming :)