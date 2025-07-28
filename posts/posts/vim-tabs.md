{
  "title": "Vim: Tabs",
  "status": "published",
  "slug": "vim-tabs",
  "date": "2025-04-05T12:33:56Z"
}

<h2>Introduction</h2>
<p>So, you are playing with Vim and getting good at but something pulls you off. Basically can't get used to switching between files or windows, and that's totally fine. Some may prefer using Window splitting, file managers, and whatnot. But let me introduce you to TABS in Vim. A simple and elegant way to open multiple files in Vim.</p>
<h2>Opening Tabs</h2>
<p>To open a tab, you can press <code>:tabnew</code>  or  <code>:tabedit</code> to open a blank Tab with no file open in it. This basically works like the <code>:e</code> command, which opens a buffer for you with no named file.</p>
<p>If you already have an existing file in the current folder you are in, then you can press <code>:tabf filename</code> or <code>:tabnew filename</code> or <code>:tabedit filename</code>. This also applies to opening folders or directories, which will open the file structure in Vim buffer.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1627994196949/A-ZMWZefa.gif" alt="tabop.gif"></p>
<p>From the above illustration, we can see that the new tab was created using the filename with the <code>tabf</code> command and an empty tab was created with <code>tabnew</code> command. Ya, we can use <code>tabnew</code> for both cases but it saves time to write two more letters. It depends on the preference as you don't have to remember one more command in this case. You can also customize the commands if you feel they are too big to type in like a simple mapping would do the trick for opening the tabs for you.</p>
<h3>To open a Tab with a file specified.</h3>
<ul>
<li>
<p><code>:tabf filename</code></p>
</li>
<li>
<p><code>:tabnew filename</code></p>
</li>
<li>
<p><code>:tabedit filename</code></p>
</li>
</ul>
<h3>Open a Tab without any file specified.</h3>
<ul>
<li>
<p><code>:tabnew</code></p>
</li>
<li>
<p><code>:tabedit</code></p>
</li>
</ul>
<p>You can open the tabs as per your choice like it could be ideal if you are gonna use certain files for a longer duration of time. This could be very ideal for various programming cases especially in Web, Android, Application Development where we need to edit a few files again and again. If you prefer Window-Splitting, that's totally fine, this is just to tell that there exist other ways as well.</p>
<h2>Switching Tabs</h2>
<p>Now if you are comfortable with opening tabs, we can now move on to switching between tabs. If you just have few tabs open, you can easily switch to the next tab using <code>gt</code> and to the previous tab using <code>gT</code> commands. But if you are in a great mode and want to open ten-twenty tabs XD, then you can use the numbers before the <code>gt</code> command. Like you can type <code>5gt</code> to move to the 5th Tab. If you do not know which tab is which, you can type in <code>:tabs</code> and this will open up the currently open tabs along with the numbers.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1627996436129/vApSyRD6o.gif" alt="tabsw.gif"></p>
<p>You can see my keystrokes(except the last keystroke) in the lower right corner. We can easily switch between tabs using the three sets of commands and surely configure them as per your preference. We also saw the <code>:tabs</code> command which can be pretty handy if you are working with a number of tabs.</p>
<p>You can move around Tabs using some special commands like <code>:tablast</code> to move to the last tab and <code>:tabfirst</code> to move to the first tab.</p>
<ul>
<li>
<p><code>gt</code>  -&gt; Move to the <strong>NEXT</strong> Tab being in Normal mode.</p>
</li>
<li>
<p><code>gT</code> -&gt; Move to the <strong>PREVIOUS</strong> Tab being in Normal mode.</p>
</li>
<li>
<p><code>ngt</code>  -&gt; Move to the Nth Tab in Normal mode. (n is any number of tab which are opened).</p>
</li>
<li>
<p><code>:tablast</code> -&gt; Move to the <strong>LAST</strong> Tab.</p>
</li>
<li>
<p><code>:tabfirst</code> -&gt; Move to the <strong>FIRST</strong> Tab.</p>
</li>
<li>
<p><code>tabs</code>  -&gt; Get a list of Tabs which are currently opened. (includes file opened in the tab and the number)</p>
</li>
</ul>
<h2>Closing Tabs</h2>
<p>So, after opening tabs you want to close em right? That is quite simple as expected. Just type <code>:tabc</code>, this will delete the current tab. You can close the specific tab by prefixing <code>tabc</code> with the number of that tab. Like if you want to delete the 2nd tab, use <code>:2tabc</code> to close the 2nd tab.</p>
<p>If you want to reopen the closed tab, you can look out for the buffer name using <code>:ls</code> and then after finding the number of buffers in which your tab was open, you can type <code>:tabnew +nbuf</code>, here n is that number of the buffer.</p>
<p>If you want to close all the tabs except the current one, you can use <code>:tabo</code>. This will clear the tabs except in which you are in, hence it will also collapse the top tab bar showing the file opened in those tabs.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1627997764101/HhUmFFQLZ.gif" alt="tabce.gif"></p>
<p>As from the above GIF, we can see we located the number of the buffer last closed as we knew the name of the file which was opened in that tab. We also saw how to delete the specific tab using its number and the current tab.</p>
<ul>
<li>
<p><code>:tabc</code> -&gt; Close the current tab.</p>
</li>
<li>
<p><code>:ntabc</code>-&gt; Close the Nth Tab.</p>
</li>
<li>
<p><code>:tabo</code>  -&gt; Close all the tabs except the current Tab.</p>
</li>
</ul>
<h2>Re-ordering Tabs</h2>
<p>This is a very tiny little detail but becomes a super tool in many cases. Let's say you want some reference of some content in the file, again and again, it's quite likely you should make the tabs nearby instead of switching tabs again and again. You can use Window splitting in this case, though we will see how to reorder tabs just for having the grasp on using Tabs in Vim.</p>
<p>To reorder tabs, you are basically moving a tab from one position to other. Let's say you have a Tab at position <code>5</code> which is your current tab, you want it at position 2. So what you can do is move the current tab to position two, as simple as to speak <code>:tabm 1</code>. This will move the current tab which is at number 5 to the 2nd position. Remember the tab order is 0 based so just use the number you are thinking minus 1. So the command becomes <code>:tabm n</code>, where n is the index of the tab(starts from 0, the left-most tab). If you want to move to the last tab, you would not specify any number just type the <code>tabm</code> command, and that's it.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1628001405055/m0XQAfdQJ.gif" alt="tabmv.gif"></p>
<p>From the above example, we were able to move around the tabs to our desired location without a hassle. This is some basic stuff you can do with Tabs in Vim, surely you can add in your custom mappings to enhance the productivity and improve the workflow in Tabs in Vim.</p>
<ul>
<li><code>:tabm n</code> -&gt; Move the current opened Tab to the Nth position (Starts from 0).</li>
</ul>
<h2>Conclusion</h2>
<p>So, we have seen how we can use Tabs and move around in between files and folders, we are now able to open, close, move, navigate around the tabs in Vim. By using some custom mappings, this can be overhauled for much fewer keystrokes that get in it. There are many other navigation techniques in VIm, and using Tabs is one of them, surely it won't suit everyone but there will be someone who will prefer using this. Thank you for reading till here. I hope you learned something from this to enhance your grasp in Vim. Happy Coding and Viming :)</p>
