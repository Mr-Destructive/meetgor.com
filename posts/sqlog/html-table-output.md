---
type: "sqlog"
title: "SQLite dot commands: html tables with output and once"
date: 2025-08-25
---

The `.once -w` or `.www` was released in [SQLite 3.47](https://www.sqlite.org/releaselog/3_47_0.html) release last year.

I was trying this command

```shell
.once -w results.html
```

But was getting an error that it was not a valid argument, and the .www command doesn’t exist. I saw the git repo and found the source code and found out that it was released in the recent version. My laptop had SQLite 3.37 version installed.

I ran a docker image of sqlite to the latest version and check out the working of that .www command or .once -w subcommand option.

## What is .once -w or .output -w or .www

All of these are equivalent
- `.once -w`
- `.www`

The output is just for all the subsequent commands to be saved in a file, whereas the `.once -w` and `.www` are only for the next one query.

This command basically allows the resultset to be saved in a html format/file. Note that you cannot specify a file, it will be stored in a temporary file, and if that is accesible and readable to the browser and enough permissions are set to the temp folder (whatever that is for linux and windows or mac), it should render an html page (its simple)

However, if you are thinking of getting some specific file to write the results then you need to do some work

```
.output result.html
.print '<!DOCTYPE html><html><body><TABLE border="1" cellspacing="0" cellpadding="2">'
.mode html
.headers on
SELECT abs(random()%10)+1 as "some numbers", 'number' from generate_series(1,10);
.print '</table></body></html>'
```

This will store the html document in the result.html with proper tags and semantics not just the table element

Phew! That might not be a lot, but that is the thing, the .mode html will only get you the table element but it wont render properly, browsers won’t be able to read the raw html elemetnts without the html doctype tags.


