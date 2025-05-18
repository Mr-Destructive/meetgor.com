{
  "type": "posts",
  "title": "Setting up Vim for BASH Scripting",
  "subtitle": "VIM + BASH = Killer Speed Automation",
  "date": "2021-06-10 04:52:07 +0530",
  "status": "published",
  "tags": [
    "bash",
    "vim"
  ],
  "slug": "vim-for-bash",
  "image_url": "https://meetgor-cdn.pages.dev/vim-for-bash.webp"
}

## Vim and BASH? 
Bash Scripting is a powerful skill to have as a programmer because we find Linux almost everywhere and to get through it you must have a command over its interface which is generally the BASH shell. Vim is a great option for doing this, or probably the best out there! Because Vim is pre-installed in almost every Linux distribution. This is not an in-depth setup for BASH on Vim, it is a simple editorial on starting up BASH scripting on the Vim editor. So without wasting time on "Vim features" let's dive in with the setup for BASH in Vim.

## Boilerplate macro
Setting up a bash script doesn't require much code but still in some cases it can be a bit hassle and to avoid the repetitive task, one can easily set up a macro for the boilerplate BASH script.

```vim
nnoremap bs i#!/bin/bash/<ESC>o
```
Ok that was pretty dumb but it can scale pretty quickly and it will be nice to tailor it as per needs, here's some snippet with function pre-loaded.

```vim
nnoremap bs i#!/bin/bash/<ESC>o
nnoremap bs i#!/bin/bash/<ESC>o<ESC>ofunction main(){<ESC>o<ESC>o}<ESC>ki<S-TAB>

```
![type bs to load boiler-plate code](https://s6.gifyu.com/images/bsclip.gif)

When the key bs is typed in normal mode, you enter into insert mode as per the command macro, then we type in the required text and escape to move to the next line and continue the same stuff. This could be extended further like making some input or printing out some text and any other formatted text that you could think it as repetition. 

## Sourcing Scripts
So, after creating the file, sourcing the script, and running it can be a bit slow for some people, as you have to go to the terminal and toggle in the permission to run the script and then run, But pull on your seatbelts as this is VIM! You can die due to slowness!

```vim
nnoremap sh :!chmod +x % && source %
```

![type sh to run script](https://s6.gifyu.com/images/shclip.gif)

When the sh keys are typed in the normal mode, the preceding command after ! (bang) will be executed in the terminal, the && keywords will execute the second command only when the first command is successfully executed.
 That just can so fast! Imagine doing this for long scripts and especially for debugging, it will waste 2 minutes every time you leave the editor and for 10 times you do the debugging, you will carelessly was roughly 20 minutes! Improve your debugging skills surely :)

## Plugins
There are very few plugins out there for BASH as for VIM, but it's quite to write scripts even without any plugins. One of the most supported and popular plugins for BASH in Vim is  [Bash-Support-Vim](https://www.vim.org/scripts/script.php?script_id=365) for auto-completion and [Shell-Check](https://www.shellcheck.net) for finding/correcting any bugs or error in the script . 
The mentioned plugin is quite awesome and it can greatly improve the speed of scripting for BASH, some commands such as shortcuts for writing if-else, while, for loops, commenting and other aspects in the scripting. The thorough documentation for such commands is also provided by the  [plugin website](https://wolfgangmehner.github.io/vim-plugins/bashsupport.html). 
This can be used for autocompleting keywords and writing nested if-else and other logical operators in BASH scripting. Again, you can do absolutely fine without plugins in Vim as it is heavily customizable to the user's need and can be very rewarding to set up your own configuration for BASH. You can use standard Vim(barebones) for auto-completion as well with the command CTRL+N and CTRL-P to move down and up respectively.


## Some More Tricks
BASH in Vim can be quite versatile to use as it provides some custom addons to make the script more functional and easier to understand. Some tricks such as using autocompletion can be quite inconvenient to use at once but it can get really smooth after some runs at writing the scripts.
- In BASH Scripts there are quite a lot of brackets to play with that's why to jump around swiftly around such parentheses or brackets you can use **% to move from the opened to closed brackets or vice versa**.
- You can execute any terminal command from Vim, be sure to be in command mode and press ! after the command, you would like to execute. This will run the command from the terminal and you don't have to leave the editor, it saves a ton of time and it's blazingly fast.
- With the above trick, you kind of have a superpower within Vim to make, build, source, run the files or scripts within Vim, that is not repetition but it can run bash within bash. Ok! that's was pretty fast. Don't die of quickness now!

Writing BASH scripts in Vim can be also boosted by using some built-in commands such as adding comments for multiple lines at once and some unexplored stuff which can be learned in the way to understanding the flow of Vim and BASH together. Happy Coding and Viming :)
