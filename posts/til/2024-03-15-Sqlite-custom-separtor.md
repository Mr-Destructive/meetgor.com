---
type: til
title: "SQLite importing CSV with custom separator"
description: "Explorng SQLite CLI with inline CSV import command with custom separator"
status: published
slug: sqlite-inline-custom-separator
tags: ["sqlite",]
date: 2024-03-15 22:30:00
---

explaination about how i learned how to write a command for importing a csv 
into a sqlite3 db with cusotm Separator with a single command inline

## Introduction

I was exploring some Issues that I can solve given the context of the problem and the intution for the solution.

I have some github repositories that I always look for to learn more. Some of them are:

- [Steampipe](https://github.com/turbot/steampipe)
- [MindsDB](https://github.com/mindsdb/mindsdb)
- [Turso](https://github.com/tursodatabase)

So, navigating around for a few minutes, I landed on this [issue](https://github.com/tursodatabase/turso-cli/issues/811).

The issue is really well explained in terms of what the feature is to be added, how the usage should be, which for a contributor is half the job done.

## What?

The [turso-cli](https://github.com/tursodatabase/turso-cli) is a CLI for [Turso](https://github.com/tursodatabase/turso) that is used to manage libsql databases with the [turso](https://turso.tech) platform.

This issue is about adding a flag to the `db create` command to allow the user to pass in a custom separator while imporing a csv file into a sqlite3 database.

The only puzzle piece left is to answer the question `how`?

## How?

So, I went in to the Codebase and found where the `db create` command has been handled and landed on this [file](https://sourcegraph.com/github.com/tursodatabase/turso-cli/-/blob/internal/cmd/db_create.go)

While controbuting to open source, I try to do the small steps and solve try to maintain the excitment with that progress. Because if you cannot find the solution, you try to procastinate and in the end nothing gets accomplished. So, breaking down the problem into small chunks is much helpful than solving the entire problem.

In this case, I  tried to add the flag in the cli which is pretty straight forward. We just add one more entry in the list of flags in the `db create` command. This step is just to add a functional CLI for taking the input of csv-separator string, however, this won't do anything for the functionality part of things, just allow the user to specify the separator/delimitor to use for parsing the CSV file.

Currently, the command that `turso db create` uses under the hood for creating a db from a csv file is:

```bash
sqlite3 "-csv" "dbName" ".import <FileName> <TableName>"
```

The above command is found in the [group flag](https://sourcegraph.com/github.com/tursodatabase/turso-cli/-/blob/internal/cmd/group_flag.go) file. To parse in the separator, we can use another string as `.separator ";"`, here the `;` is the character that should be used as the separator for the csv file into the db.

```bash
sqlite3 "-csv" "dbname" ".separator" ";" ".import <FileName> <TableName>"
```

Thank you, Happy Coding :)
