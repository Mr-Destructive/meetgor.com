---
type: "sqlog"
title: "SQLite dot commands: change directory command"
slug: sqlite-cd-dot-command
date: 2025-08-31
tags: ["sqlite", "sql"]
---

## Change directory dot command

If you are in a sqlite shell and forgot to change directory or want to navigate to a separate directory, you can do that with the `.cd` dot command.

```sqlite
.cd /path/to/directory
```

This is better than doing `.shell cd /path/to/directory` because it doesn't open a separate terminal process. So, the `.cd` is persistant throughout the session, whereas the `.shell cd <path>` will only within that command (subprocess).

The `.cd` command changes the working directory of the SQLite shell itself, so the change persists for the rest of the session. This means commands like `.import`, `.read`, or `.output` will automatically look for files in the new directory.

However, `.shell cd <path>` spawns a separate subprocess, and the directory change is discarded as soon as that command finishes. It does not affect SQLite’s own state of the current directory.

So if you plan to read or write multiple files from a different location during your SQLite session, prefer the built-in `.cd` command.


