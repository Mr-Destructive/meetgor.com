{
  "type": "posts",
  "title": "Setting up Vim for Python",
  "subtitle": "For the python programmers who need speed!",
  "date": "2021-06-06 11:35:25 +0530",
  "status": "published",
  "tags": [
    "python",
    "vim"
  ],
  "slug": "vim-for-python",
  "image_url": "https://meetgor-cdn.pages.dev/vim-for-python.webp"
}

## Introduction
Vim is quite a powerful text editor which can add performance to the already fast typed language Python. Vim can be highly customizable and efficient to use as it has the power of **adding custom plugins and plugins managers, key mappings**, and the most critical weapon of vim - Access to the terminal straight away.
This is not a full-featured guide of using vim for python, it's just a quick setup for using python on vim blazingly fast!!

## Plugin Managers
So let us start making Vim, the text editor suitable for a python programmer. Firstly we'll need the vim plugin manager. There are different plugin managers out there, each of them has the same purpose to install, upgrade and manage the plugins for vim. You can install any one of them and get up and running.

-  [Vim Plug](https://www.vim.org/scripts/script.php?script_id=4828) 
-  [Vundle](https://github.com/VundleVim/Vundle.vim) 
-  [Pathogen](https://github.com/tpope/vim-pathogen) 

These are some of the finest and well-supported plugin managers in vim. You can use any of these plugin managers, and get started by installing some plugins.

## JEDI-VIM- Auto completion 
Firstly I will like to install Jedi for code completion in Python. The plugin can be simple and straightforward to install using any of the above plugin managers. Jedi-Vim provides some neat and clean** syntax analytics and autocompletion for Python in Vim**. You'll find the docs and installation process here  [JEDI-VIM ](https://github.com/davidhalter/jedi-vim) 

## NERDTree-File manager
Next, It would be great if we install a file manager for managing the files and folders in the code directories. We can simply use the Nerdtree plugin for this. NerdTree is quite a **fantastic plugin for file management in Vim**. It simply makes Vim feel like VS Code. The installation and docs can be found here  [NERDTree](https://github.com/preservim/nerdtree).

Nerdtree commands can be longer to write, for that let's start mapping and for that, we can start editing our Vimrc. 

```vim
set number
syntax enable
filetype indent on
set tabstop=4
set softtabstop=4
set autoindent 
set encoding=utf-8
``` 
This can be some addition to your existing vimrc as you might have a setup for plugin managers. You can choose the color scheme of your choice, don't waste time selecting the color scheme. Feel free and modify the vimrc according to your knowledge and choice. 

## Keymappings
We move on to the Key mappings for NERDTree and other features. You can make mappings generally in the normal mode but there might be cases where you need to use maps for visual mode or insert mode as well, that entirely depends on the user :)

To map in normal mode, we'll its command to be specific:


```vim
nnoremap <C-n> :NERDTree<CR>
``` 

This will map CTRL+n to open the NERDTree file manager to the left, saving a bit of time and avoiding frustration. Feel free to add any keymap of your choice, this is just for demonstration. 
You can further automate NERDTree for switching between tabs because it makes you type CTRL+w twice, you can reduce that to just typing w.

```vim
nnoremap w:<C-w><C-w>
``` 

## Integrated Terminal Macros
We can open a terminal window like a split between the editor. We can simply use the command :terminal to split the window horizontally, where the upper split will be terminal and the down window will have the editor. This is quite a neat feature of Vim in that it blends with the terminal so well so that we can switch between the terminal and the editor very quickly. For that, you can create a macro if you need to fire up a terminal again and again.
```vim
nnoremap <C-t> :terminal<CR>
```
If you place the above macro in your vimrc and then type Ctrl+t, the exact thing will happen to fire up a terminal split but with fewer keystrokes and without leaving the normal mode. 
Also, the NERDTree macro can be also fruitful with this as it will make a full-blown IDE-like feeling inside of Vim.
![Demonstrate macros for NERDTree and terminal split](https://s6.gifyu.com/images/screenrecording.gif)

## Running the Code with a snap

We can automate the process of running python scripts inside of vim. Instead of typing out the entire command for executing python script from vim. We can use keymaps for it as they can significantly boost the time required to run and debug the code. 


```vim
nnoremap py :!python %
``` 

This is a small map but can save a lot of time and give some motivation to use vim as you run the code blazingly faster than other editors. I have used py, but it can cause some problems as p is already mapped for pasting. So it's better to use other key combinations such as ty, yh, or any other key combination of your choice. Try it out and add your own flavor that's how we all learn.

So, that's the basic set-up for python on vim, you can make more custom mappings, find more plugins and test out which work out the best for your workflow. Happy Coding and Viming ;) 
