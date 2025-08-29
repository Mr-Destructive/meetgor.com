---
type: "logsql"
title: "SQLite dot commands: read command is recursive?"
slug: sqlite-read-dot-command-is-recursive
date: 2025-08-29
---

Yesterday, while exploring the `.read` command I discovered a little hack.

We know you can read arbitrary sqlite shell commands from the `.read | ` operator, even the dot commands.

So, you can even use `.read` itself again. Which means we can call the same file to read inside the script.

Ops!

Let's create a file called `loop.sh` this will emit certain valid sql/sqlite commands like so:

```bash
#!/bin/bash
echo "SELECT 'hello, world!';"
echo ".read |./loop.sh"
```

The first will print out "SELECT 'hello world!';" so when directed to the shell, it will be executed and print "hello, world!"

Then the 3rd line `echo ".read |./loop.sh"` is the magic, it will execute the script again right?

Because it's a valid sqlite shell command and it will execute it. It will take the read command and print everything when executing the `loop.sh` shell script.

What's the output?

```bash
$ cat loop.sh
#!/bin/bash
echo "SELECT 'hello, world!';"
echo ".read |./loop.sh"

$ sqlite3
SQLite version 3.50.4 2025-07-30 19:33:53
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
sqlite> .read |./loop.sh
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
hello, world!
Input nesting limit (25) reached at line 2. Check recursion.
sqlite>
```
Well, SQLite developers are smarter to handle the edge case I think. :)

This is a little fun, but it's kind of a quirk, an not-so easter egg but interesting thing to know.

