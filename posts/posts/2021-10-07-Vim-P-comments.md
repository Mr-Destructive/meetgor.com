{
  "type": "posts",
  "title": "Comment/Uncomment Code: Vim for Programmers",
  "subtitle": "A short guide to comment/uncomment chunks of code effectively in Vim",
  "date": "2021-10-07 16:45:42 +0530",
  "status": "published",
  "tags": [
    "vim"
  ],
  "slug": "vim-un-comment-p1",
  "image_url": "https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643288865/blogmedia/jbs2hcaiplfvwsgrcfgl.png"
}

## Introduction

We as programmers always fiddle with commenting out code for code testing, documenting the function of code, and most importantly debugging. So you can't wait to comment on a large chunk of code manually, as it is quite a tedious thing to do. Let's do it effectively in Vim.

In this part of the series, I'll cover how to comment/uncomment chunks/blocks of code effectively in Vim. We will see and use some commands, keybindings for doing so, and also we would add certain components to our vimrc file as well to design some custom key mappings.  Let's get faster with Vim.
  
## How to comment multiple lines effectively

To comment on multiple lines of code, we can use the Visual Block mode to select the lines, and then after entering into insert mode, we can comment a single line and it would be reflected on all the selected lines.

1. Press `CTRL+V` and Select the line using j and k

2. After Selecting the lines, Press `Escape`

3. Press `Shift + I`, to enter insert mode

4. Enter the comment code (`//`, `#`, or other)


![vimcoment.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1633518136135/06dfBTq2T.gif)

So, using just simple steps you can comment out large chunks of code quite easily and effectively. If you are using some other language that has multiple characters for commenting like `//`, `- -`, etc, you can type in any number of characters while being in insert mode after selecting the lines.

 
![vimcppcom.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1633520509953/0q-k2ZHC7.gif)

This might look a bit wired on the first try but just try it every day, It is a life-saving and very satisfying experience once applied in a real-world scenario.


## How to uncomment multiple lines effectively

Now, as we have seen to comment out a large chunk of code, we can even uncomment the code very easily. It's even simpler than commenting the code.

1. Press `CTRL + V` to enter Visual Block mode

2. Select the commented characters

3. Press `d` to delete the comments

4. Press `Escape`

![vimuncoment.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1633518156818/GJzRPTI3I.gif)

We can simply use the CTRL + V to select the comment, and then press d to delete all the comment characters. 

**We are using the Visual Block mode as we only want the comment to be selected and not the entire code associated with the lines.**

## Using Multiline Comments for Programming languages

Now you might say, why use multiple single-line comments when we can use multiline comments in almost all programming languages. Well, Of course, you can do that, it's easier for reading the code if syntax highlighting is accurate and greys out the commented part. We can simply add those characters to the start of the block and at the end of the block.  

But in Vim, we can customize that too, just imagine when you just select the chunk/block of code that you need to comment out and then simply press a few keystrokes (just 2) and the multiline comments are automatically (programmatically) added as per the programming language extension of the file.

Isn't that cool? Well, you just need to copy-paste the below code to your Vimrc file and source it and you are good to go. 

```vim
function! Comment()
    let ext = tolower(expand('%:e'))
    if ext == 'py' 
        let cmt1 = "'''"
	    let cmt2 = "'''"   
    elseif ext == 'cpp' || ext =='java' || ext == 'css' || ext == 'js' || ext == 'c' || ext =='cs' || ext == 'rs' || ext == 'go'
	    let cmt1 = '/*'
	    let cmt2 = '*/'
    elseif ext == 'sh'
	    let cmt1 = ":'"
	    let cmt2 = "'"
    elseif ext == 'html'
	    let cmt1 = "<!--"
	    let cmt2 = "-->"
    elseif ext == 'hs'
	    let cmt1 = "{-"
	    let cmt2 = "-}"
    elseif ext == "rb"
	    let cmt1 = "=begin"
	    let cmt2 = "=end"
    endif
    exe line("'<")."normal O". cmt1 | exe line("'>")."normal o". cmt2 
endfunction

function! UnComment()
    exe line("'<")."normal dd" | exe line("'>")."normal dd"   
endfunction


vnoremap ,m :<c-w><c-w><c-w><c-w><c-w>call Comment()<CR>
vnoremap m, :<c-w><c-w><c-w><c-w><c-w>call UnComment()<CR>

```
The below screencast is an example of `HTML` snippet in a file that is getting commented using mapping with the keys `,m` you can put any other keybinding you like. 
![htmcm.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1633595891674/hbhrbtRHd.gif)

---
Similarly for the next screencast is of an `Javascript` snippet in a file which is getting commented using a mapping with the keys `,m` and uncommented using `m,` again you can put any other keybinding you like. 

![jscom.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1633595919104/xGTh5ztWu.gif)

---
The following screencast is of a shell script(BASH) snippet.
![shcom.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1633596156121/tbGHQBSSA.gif)

---
### Multiline Comments in various Programming Languages:

#### 1. C / C++ / Java / Javascript / CSS / C# / Rust / Go / PHP / Swift / Dart / Kotlin
```
/*
*/
```
#### 2. Python
```
'''
'''
```
You can even use `"""` double quotes instead of single quotes

#### 3. BASH (Shell Scripting)
```
: '
'
```
You can even use `: "` and `"` double quotes instead of single quotes

#### 4. Haskell

```
{-
-}
```
#### 5. Ruby

```
=begin
=end
```

#### 6. HTML

```
<!--
-->
```

#### 7. Julia

```
#=
=#
```

### Understanding the Commands / Keymapping

**NOTE : You need to go from the top to bottom while commenting on the block of code, otherwise, it would be a mismatch in commenting for specific language syntax. While uncommenting the order doesn't matter.**


#### Getting the extension (filetype)

In Vim, we can get the file extension i.e. we can get the programming language associated with the current file. To do that we can use, `expand('%:e')`.

This will give us the file extension of the current file. Just for simplicity, `%` means the current file, added to it is `:e` for excluding the filename and keeping the extension. We convert the extension into lowercase just for keeping things safe and programmatic and store it in a variable `ext`. 

#### Checking for programming language
We then can then use an if-else ladder to check for the programming languages and assign two variables `cmt1` for the initial characters in the multiline comment and `cmt2` for enclosing the comment. 

#### Typing in the characters

We can use the function `line("'<")` to get the line number of the previous visual selection. Similarly, `line("'>")` for the ending line. We are using the `exe` command to execute the function `line` and so we have to use a concatenation of the commands even to write the raw commands like `i` to insert mode, `o` to insert mode but a line below the cursor. So, we use `normal` command for that. This command indicates the interpreter to execute the following commands from the normal mode. 

We have to enclose the `normal` command in double-quotes/single quotes. We can simply use the variable again with concatenation.

```vim
exe line("'>")."normal o". cmt2 
```  

The above command will fetch the last line's number of the previous visual selection followed by entering `o` from the normal mode and concatenated with the value of the variable `cmt2` which we have already initialized in the `Comment` function. We are using `|` for running multiple commands as we also need to include the comment at the beginning of the visual selection. 

For uncommenting the code, we are simply deleting the entire first and the last line in the visual selection. For that, we have used `dd` from the normal mode.  


### Conclusion

So, from the following type of tutorial, we were able to set up our Vim editor for efficient code commenting/ uncommenting using some commands, key shortcuts, and configuring the vimrc for making custom keymappings. We were also able to understand the multiline comments in various programming languages and use them in Vim very effectively with a simple addon to the config vimrc file.

Thank you for reading, hope you found this article helpful. If you have any queries or wanna add multiline comments for some more programming languages please let me know in the comments or contact section. 

Happy Coding :)

### References

- [StackOverflow - Commenting lines in Vim ](https://stackoverflow.com/questions/1676632/whats-a-quick-way-to-comment-uncomment-lines-in-vim/1676690)

- [Liz Lam - 3 ways to comment code in Vim](https://dev.to/grepliz/3-ways-to-comment-out-blocks-of-code-in-vi-6j4)
 
- [StackExchange - Use variable in normal command](https://vi.stackexchange.com/questions/9644/how-to-use-a-variable-in-the-expression-of-a-normal-command)
