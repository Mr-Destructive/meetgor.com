{
  "type": "posts",
  "title": "Vim: Registers",
  "subtitle": "The guide in understanding Vim registers and where your text gets stored after copying or deleting.",
  "date": "2021-07-21 23:41:57 +0530",
  "status": "published",
  "tags": [
    "vim"
  ],
  "slug": "vim-registers",
  "image_url": "https://cdn.hashnode.com/res/hashnode/image/upload/v1626940723316/nFGNljJi0.png?w=1600&h=840&fit=crop&crop=entropy&auto=compress"
}

## Introduction

Have you ever found it difficult to manage and operate text in Vim, especially in Cut/Copy/Paste stuff. You cannot figure out how on earth am I supposed to retrieve my un-saved, un-committed changes in Vim. Then this my dear friend is an article for you. You will learn what are Vim registers and where your deletes and copied stuff resides. This won't be a comprehensive guide but enough to let you through in performing the wizardry of registers in Vim.

## What are Vim Registers?
Vim Registers are memory addresses that store data. Technically they are spaces of memory that Vim can store to retrieve later. It can store text, operations, commands, and anything you can do with text in Vim. 

## Access Vim Register
To store something in register `x` we can use the format `"[register-name][command]`. This will perform the command and store the output in the register specified.

A register can be accessed by using the double quotes `"` before its name. Let's say we have stored something in x register we can access its content using "x in the normal mode and lastly we can perform commands here.

![vimreg.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1626940084837/x99b92Wsq.gif)

For example, in the above gif, I tanked the selected text using the yank command but I stored it in the register x by prefixing the command with `"x`. Hence the command became `"xy` and similarly for pasting from the x register we can use `"xp`, quite simple, right? But what is the purpose of it? This can really shine in handling multiple files and imagine you lost your yanked text by using yank somewhere else. Vim registers add extra dimensions in storing and retrieving text, it is not used often but can shine really well in certain tricky situations. 

## Commonly used Vim Registers
You will use some Registers daily in Vim if you use it all the time. Some common ones are the clipboard registers/ selection registers `"+` and `"*`. These are used to perform operations on the content stored in the system clipboard. They come quite handy in copy and paste in Vim from other applications such as Stack Overflow XD. Just kidding, it can be used in any other application on the system. We can add `p` or `y` after the registers to paste and respectively in Vim, so to paste and yank text to and from Vim, you can use `"+p` and `"+y`.

## Types of Vim Registers
There are many different types of vim registers used for particular needs and also some for custom editing. 

- ### 26 Named registers (`a` to `z`)

These are the registers just for you. You can use this for whatever copy/pasting/editing and other creative stuff you could do with this. In the example where we utilized the `x` register, which was one of them. We can use a to z any register we want, it just boils down to the ease of the key combination as per preference. **Remember anything you yank into any register will also be stored in the unnamed register (discussed later)**. 

Also, if you try to use capital named registers (A-Z), the content will be appended to the lowercase named registers as well along with the unnamed register. Making it a bit safer to work with but messy at the same time. 

- ### 10 numbered registers (`0` to `9`)

These are the registers updated automatically filled in by Vim, they store the recently yanked or deleted content in the current buffer. 
- The `0` register stores the most recent yanked/copied text.
- The `1` register stores the most recent deletion of text.
- The `2` register stores the 2nd most recent deletion. 

.. and so on.

The content of the 9th register is flushed away every single time you delete something. The contents are shifted to the next numbered registers every time there is a deletion. For example, If we delete something the content is stored in the `1` register but its previous content gets shifted to the `2` register, and so on. So we have 9 clipboards for our deletion history. That is undoubtedly a powerful thing.

- ### Selection and drop registers (`*` , `+` , and `~`)

As discussed earlier, Selection registers are used for storing the contents of the system clipboard. There are differences in the `*` and `+` for Unix-like operating systems but they can do similar things. In Unix-like Operating systems, the `*` register stores the text from the mouse selection in the X Window, whereas the `+` register is used to store the text system clipboard. On Windows and other operating systems besides Unix-like, both of them function similarly.

The Drop register `~` is available only for the Gvim version as it stores the text from the last drag and drops/drag down operations performed. 

- ### Read-only registers (`:`, `.`, and `%`)

The registers `:`, `.`, and `%` are Read-only, which means you cannot change their content, they are modified and maintained by VIm automatically. 

Using the `:` register, we can get the content of any register in the Insert mode so that we can copy the contents and edit if we have stored it in some registers. This is quite a neat little feature, not used most of the time but it can become the most powerful tool in corner cases and tricky situations. 

![vimc-r.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1626946965392/KBSZmxOHvi.gif)
From the example in the gif, we deleted the line, and hence the content was stored in register 1 so when we press `Ctrl + R` in Insert mode, we can now paste in the contents of any register. Here we used the register `1` to get our content back. 

The `%` register holds the name of the current file, this is really a great and life-saving editor, as it is also used in automating the compilation and running of the source files from Vim itself. 

The `.` register contains the last inserted text, this can also be used in various scenarios. They make Vim a bit polished and well organized.

![vimc-r.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1626948017124/dQCky3kfC.gif)

From the above gif, we inserted some text, and using the Read-only register, we inserted the last inserted text to the screen using the `.` register and also we used the register `%` to get the current file name.

- ### Expression and Search pattern register (`=`, `/`)

The register `=` is used to perform operations and store them in the register. Vim has some built-in functions and some of them are listed  [here](https://renenyffenegger.ch/notes/development/vim/script/vimscript/functions/index). We can use these functions after we type `Ctrl + R` in insert mode and `=` to assign it to the register `=`. I have used some simple multiplication operations in the above gif, but you can take this further however you like and desire. 

The register `/` takes searching in Vim even further. We use /pattern in normal mode to search for the pattern. But the actual work is done with help of the register. We actually store the text in the register and Vim accesses it later t find the next pattern or any other operations performed with the pattern. This is again a read-only register but you can change its content using some methods and practices. It is not preferred to change the read-only registers as it can mess up the functions of Vim.
 
- ### small delete and the alternate file register (`-`, `#`)

Small delete is quite a meaningful and sensible word to put in the context of its functioning. The `-` register or small delete register stores the content of some small deletes like characters or words. Basically, anything smaller than a line is small for Vim. If you delete something let's say a word from a line, it'll be stored in the `-` register. 

The '#' register or alternate file register is the register that stores the alternate file in Vim. An alternate file is a file edited before the current file in Vim. If you edited a `file-x` and now you are in the `file-y`, the `file-x` becomes your alternate file. 

![vimc-r.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1626950574633/P7U01Ssy_.gif)

This example shows how we can use the small delete register to retrieve back the small deletions. Also, we saw the current file and alternate files in Vim and in which registers they are stored.

- ### Unnamed and black hole register (`"`, `_`)

Unnamed register `"`, simply to put in it is a register which stores everything from deletions to yank. It stores the text performed by the d, D, x, X, s, S, c, C, y, Y, some variants of these commands. This is quite good and a bad feature sometimes. It is nice to have such types of registers in hand. 

Wear your space goggles as we see will how Black holes work :) Blackhole register is an elegant touch to the series of registers we have. The register acts as a sink for our deletions. We can simply dump anything we don't care about in this register. This register helps in scenarios where we want to avoid adding stuff from deletions to the named or numbered registers. Isn't it like Blackholes? Just Swallow the things. 

![vimc-r.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1626951775215/OEanNfA4a.gif)
The above example makes it clear that we simply dump the deletion to the `_` register and avoid changing the rest of the registers which may replace the deletion and lose the saved configuration or set up of registers. We first deleted the line `solve(s);` and dumped the statement `string s = "0110` into the black hole register `_`, when we retrieve with thee unnamed register `"` we get back the deleted statement rather than the dumped statement. Quite neat and elegant.

## Where can you use Vim registers

Well, this may differ as per the experience with your programming journey, you may find it useless as a beginner but mind you this is some extraordinary stuff that can boost some development time. You may be tempted to use it the scenarios where there is a huge depth in file systems and the code to be edited is a lot for you. When you have a ton of things to be edited systematically, this concept can add 
 lights and dimensions to your work and make it comfortable and risk-free. 

## Conclusion

![vimregistertable.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1626953924487/2bPqejVkT.png)
The above chart is quite simple to remember if you are really into making Vim for complex projects and even for staying productive in Vim editor. Hope you find it interesting and easy to remember.
This was a small introduction to Vim-registers but is a small topic with extensive usage and skill-dependent concept. Vim is already a customizable editor, registers add an extra dimension to its customizability in editing and programming. Thank you for listening to me. Happy Coding :)

**References** :
 [baeldung.com](https://www.baeldung.com/linux/vim-registers), [brianstorti.com](https://www.brianstorti.com/vim-registers/)
