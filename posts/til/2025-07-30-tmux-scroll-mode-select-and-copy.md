{
  "title": "Tmux scroll mode select and copy",
  "post_dir": "til",
  "type": "til",
  "status": "published",
  "date": "2025-07-30",
  "tags": ["tmux"]
}

I have found scrolling in tmux is a bit unconventional. Maybe its just me, but sometimes, the terminal interfaces are different, sometimes they are a bit wired, I use Ghostty right now, so selection without tmux works a charm, but the moment I am in tmux, ahmm, it kind of breaks.

The selection of text doesn't work with the mouse atleast, so I thought, I need a keyboard centric selection and copying mechanism, and dug in this rabit hole. Turns out, tmux is way better in selection then zellij, the multiplexer I had been using due to scroll and selection issues in terminal interfaces.

It's simple as `Prefix + [` to enter the scroll mode, `Shift+V` or `Ctrl+space` or `Ctrl+v` to start the selection in different ways, and finally, `Enter` or `Ctrl+j` to copy to clipboard. There is a dedicated mode for vi users, to feel at home. Tmux + Vim again a deadly combo.

## Entering Scroll Mode

To enter scroll mode, I had this config setup in my `.tmux.conf` file:

```
set -g mouse on
```

However if you want to enter without mouse on, you can use:

```
Ctrl + b + [
```

Here the `Ctrl + b` is for the prefix key, and `[` is the key for enterring into the  `Scroll Mode`

From here on you can use `hjkl` or arrow keys to scroll up or down.

To quit, you can press `q` or hit escape.

## Selecting 

Right now in scrolling mode, you can only navigate to the next or previous line. 

But you can use `Ctrl + <Space>` to select any arbitrary sections of the terminal output.
There are two modes in tmux for this the `copy-mode` i.e. the default, and the `copy-mode-vi`.

Without any config, this is be set to the default `copy-mode`, so we can use `Ctrl + Space` to begin selection in the scroll mode.

To get more keybindings for the copy-mode and copy-mode-vi, use 

For copy-mode
```
tmux list-keys -T copy-mode
```

For copy-mode-vi
```
tmux list-keys -T copy-mode-vi
```

This is a default and can be over-written by using the following config:

```
setw -g mode-keys vi
unbind-key -T copy-mode-vi Space

# replace <your-key> with the key combination you want
bind-key -T copy-mode-vi <your-key> send-keys -X begin-selection

# Example: using ctrl + b + v as the key to start selection
# bind-key -T copy-mode-vi v send-keys -X begin-selection
```

This all would be applied to copy-mode if you are not changing the keybindings for copy-mode-vi.

## Copying

Once, you have the selected region or piece of text, you can copy that to the clipboard (that is what you do, while logs-driven debugging :).

By default, you can use:

For copy-mode
- `Ctrl + w` 

or 

For copy-mode-vi
- `Ctrl + j` or `Enter` to copy the selected text to the clipboard.

This is default, and again can be over-written by using the following config:

```
setw -g mode-keys vi

bind-key -T copy-mode-vi <your-key> send-keys -X copy-pipe-and-cancel

# Example: Remap 'y' in copy-mode-vi to copy selection and exit
# bind-key -T copy-mode-vi y send-keys -X copy-pipe-and-cancel
```

This all would be applied to copy-mode if you are not changing the keybindings for copy-mode-vi.

That's it this should be more than enough for most of the log-driven or printf debugging tasks.
