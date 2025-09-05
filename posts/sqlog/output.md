---
type: "sqlog"
title: "SQLite Dot command: Output"
date: 2025-08-19
---

I will be starting to post something I learn daily about sql,sqlite or anything sql.

There is so much to learn!

Today I found you can output a result of query in sqlite shell with

```
 .output filename
```
This will start appending the result of all the queries executed in the shell to the file. It will preserve all the modes and configuration used for that specific instance of the shell. Basically just the output you would see after executing the query n the shell, it will dumo that to the mentioned file (it will not output to the screen/io as it will dump in the file)

More info here: [SQLite Dot command: Output](https://sqlite.org/cli.html#writing_results_to_a_file)


