{
  "title": "Format JSON in Vim with JQ",
  "post_dir": "posts",
  "type": "til",
  "status": "published",
  "slug": "format-json-in-vim-with-jq",
  "date": "2025-03-21"
}

We can use 
```
:%!jq .
```

To instantly format a json file opened in vim/neovim.
I relied on online foramtter and opening vscode for formatting.
But this quick tool just saves a lot of hassle. Opening VS Co** == shutting down the laptop, it clogs the memory badly, so I avoid opening VS Co** as much as possible.

Yes, to use this, you need to have [jq](https://jqlang.org/) installed on your system.


If you are elsewhere and want to format json file using jq, you can quickly do this:
```
jq . input.json | sponge input.json

# OR

jq . input.json > tmp.json && mv tmp.json input.json

```
If you are in Vim and want to format other file, just add `!` at the start of the command and that will run in the shell for you, so you don't have to exit vim.
