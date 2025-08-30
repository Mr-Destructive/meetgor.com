---
type: "logsql"
title: "SQLite dot commands: run system shell commands"
slug: sqlite-shell-dot-command
date: 2025-08-30
tags: ["sqlite", "sql"]
---

## Shell dot command

If you are in middle of a sqlite shell session, and you don't want to quit the shell to run arbitrary shell command, you can simply use the `.shell <command>` to execute any shell commands right from within the sqlite shell. How handy is this!

```sqlite
.shell echo "hello, world!"
```

That is a lame example, but it shows you the power of the `.shell` command.

Let's say I want to run a golang project, I can do this:

```sqlite
.shell go run main.go
```

Its helpful if you want to do something but you don't want to quit the shell to do that:
- look up few files/datapoints from the local filesystem,
- run scripts to populate data
- populate database and then reopen the db shell

This is are the things that I have stumbled upon, so far. Need more experience to see if there are more.




