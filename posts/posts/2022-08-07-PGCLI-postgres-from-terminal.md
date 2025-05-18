{
  "type": "posts",
  "title": "PGCLI: Postgres from the terminal",
  "description": "Running SQL queries in a Postgres database with PGCLI. Exploring the python package PGCLI, that can run SQL queries for a postgres database environment from the command line.",
  "date": "2022-08-07 20:30:00",
  "status": "published",
  "slug": "pgcli-pg-from-terminal",
  "tags": [
    "postgres",
    "python",
    "DBMS"
  ],
  "image_url": "https://meetgor-cdn.pages.dev/python-pkg-pgcli-postgres-from-terminal.png"
}

## Introduction

Have you ever used the Postgres database and did you know you don't have to launch PGAdmin every time you want to write SQL queries, you can write down those queries even inside a terminal with a python package. It's PGCLI, that can act as an editor window for writing SQL queries for Postgres Databases. We can simply execute SQL queries in a PSQL-like shell with additional features like autocompletion, text-editing modes, file input/output options, etc.

If you are like CLIs and love to play with backend systems like Postgres(database) then chances are you will love this article. This article will cover the basics of using the PGCLI tool to perform simply as well as advanced options in interacting with your Postgres database.

## PGCLI 

PGCLi is a python package that acts as a CLI for executing SQL queries in a Postgres database. PGCLI allows us to interact with any Postgres database via the command line, it can be a remote database or a local database, you can access it right away from your terminal. The package acts as a CLI for accessing and executing any SQL queries, so we can simply provide the options/parameters to plug the database in and access the schema and its related data.

## Installation

There are a lot of ways pgcli can be installed, my preferred way is with `pipx`, it just works well and you don't have to install it every time or mess up your global python packages environment. 

```
pipx install pgcli
```

![PGCLI Install with pipx](https://res.cloudinary.com/techstructive-blog/image/upload/v1659878333/blog-media/pgcli-install-pipx.png)

We don't even require installing it when using pipx, as it will each time create a fresh instance of the pgcli in a virtual environment.

We can now run the `pgcli` with parameters using the `pipx run` command. In the below command, we have used the `--help` options to list out all the commands and other options available with the `pgcli` command.

```
pipx run pgcli --help
```

![PGCLI Help](https://res.cloudinary.com/techstructive-blog/image/upload/v1659878322/blog-media/pgcli-help.png)


If this command is working fine, it means we are ready to connect and interact with a database with pgcli.

## List local databases

We can list all the databases in our local system by providing the hostname and the username, in my case, I have to provide the `localhost` and `postgres` as the hostname and the username respectively. Further, by providing the `--list` option, we can list down all the available databases on our system, basically for a given host and user.

```
pipx run pgcli --host localhost --user postgres --list

OR

pipx run pgcli -h localhost -U postgres -l
```

![PGCLI List all databases](https://res.cloudinary.com/techstructive-blog/image/upload/v1659878596/blog-media/pgcli-list-db.png)

As, we can see, the command lists out all the available databases on the local server. This can help in getting an overview of the Postgres databases present in your local system. You need to enter the password for the username to access the database.

## Connect to a database

To connect to an individual Postgres database we have two options, we can either use a URL string of postgres database or parse parameters to the command individually. It's easier to use the URL option because it wraps a lot of things in a single string rather than writing each parameters.

### Connect with the Postgres Connection URI

We can simply pass the postgres connection URI for accessing a postgres database with pgcli. The URI has a specific pattern in the order where we need to specify the values like hostname, password, username, port, and database name. 

```
postgresql://username:password@hostname:port/database_name
```

The above is the structure for a postgres connection URI, we parse in the `username`, `password` for that `username`, `hostname` with the `port`, and the `database name`. These parameters if parsed correctly can directly access the database.

```
pipx run pgcli postgresql://postgres:postgres@localhost:5432/techstructive_blog
```

![PGCLI Postgres Connection URI](https://res.cloudinary.com/techstructive-blog/image/upload/v1659878955/blog-media/pgcli-connect-uri.png)

This command will leave us in a prompt where we are basically in a PSQL shell in the provided database or host. For connecting to a local database, the hostname will likely be `localhost` and postgres database port is generally `5432` there might be exceptions to this. We can quit out of PGCLI with `CTRL + D` or `exit`.

While using a remote database, the hostname and other details might be provided, for instance, we might have a Django application, in that project, we will have `database_name`, `port`, and details mentioned in the `DATABASES` settings, those fields can be checked and a postgres connection URI might be constructed.

Let's take an example.

```python
DATABASES = {
    "default": {
        "ENGINE": "django.db.backends.postgresql",
        "NAME": "blog",
        "USER": "meet",
        "PASSWORD": "postgres",
        "HOST": "localhost",
    }
}
```

We have all the fields provided in the django project' settings file. We can use this to construct a postgres database connection URI.

```
pgcli postgresql://meet:postgres@localhost:5432/blog
```

So, that is how you can extract parameters from other technologies and frameworks for working with postgres URIs.

### Connect with CLI options/paramters

We can parse the options like `hostname`, `password`, and `database name` manually with the CLI options that PGCLI provides. 

```
pipx run pgcli --username postgres --host localhost --port 5432 --dbname techstructive_blog

OR

pipx run pgcli -u postgres -h localhost -p 5432 -d techstructive_blog
```

![PGCLI Connect Paramaters](https://res.cloudinary.com/techstructive-blog/image/upload/v1659879642/blog-media/pgcli-connect-paramters.png)


You will be prompted for a password for the database, after entering the password, you will be inside the psql shell. In there, we can execute SQL queries and other specific psql or pgcli commands.

Once we have connected to a database, we can now try to work with the SQL queries and explore some of the commands and features provided by pgcli.

## List out all databases if connected to a server

Let's say we don't know the database name, we are just connected to a database server, we can get the list of the databases inside the shell with the `\l+` command. This command is similar to the `--list` parameter in the PGCLI command, but the only difference is that we are executing it from the psql shell rather than our terminal shell.

```
# Let's say we connected to a database server

pipx run pgcli --username postgres --host localhost --port 5432


# List out all databases in that server

\l+
```

![PGCLI List databases in a server](https://res.cloudinary.com/techstructive-blog/image/upload/v1659879988/blog-media/pgcli-list-databases.png)


So, that's how we would get the gist of all the databases inside the database server to which we are connected using the `\l+` command.


## List all tables in the database

We can get a list of all the tables existing in the database we are logged in with the `\dt` command, we can simply enter the command in the prompt we are directed to once we are logged in to the postgres database.

```
\dt
```

![PGCLI List tables in a database](https://res.cloudinary.com/techstructive-blog/image/upload/v1659880232/blog-media/pgcli-list-tables.png)

We can use `\d <table_name>` to describe the details of the table provided. We get back the list of the attributes inside the table and the indexes of the relationships of the table as well as the referenced table details.

![PGCLI describe a table in a database](https://res.cloudinary.com/techstructive-blog/image/upload/v1659880453/blog-media/pgcli-describe-table.png)

We also have other commands for displaying types, schemas, roles, views, etc. with the `\d` prefix, a detailed list can be found in the [PGCLI documentation](https://www.pgcli.com/commands).

## Executing queries

We can execute SQL queries in the PGCLI prompt as normally we do in a PSQL shell. There are added benefits like certain modes of editing text and auto-completion. We'll execute a few queries and also try out other options for editing and saving queries.

Normally for simple and easy tasks, we can write SQL queries right in the PSQL shell, it might be small and simple to construct a query. Let's take the simplest example to get all the attribute lists from a specific table in a database.

```SQL
SELECT * FROM table_name
```

![PGCLI executing SQL queries](https://res.cloudinary.com/techstructive-blog/image/upload/v1659880794/blog-media/pgcli-sql-query.png)

We get an excellent table display of output which we can view by scroll or arrow keys. We can even use `j` and `k` for moving vertically in an output table.

### Wriing modes in PGCLI

We can use two modes in PGCLI prompt, one is the Emacs mode which is the default one and the other is Vi-mode for Vim fanboys. We can switch between these modes with `F4` key. The `Emacs-mode` is just a normal typing experience without any modes or shortcut macros and other stuff. Though it's nice that PGCLI offers a VI-mode, it is great for people using Vim, when using the Shell for editing a Query, this mode can be utilized.

![PGCLI VI-mode](https://res.cloudinary.com/techstructive-blog/image/upload/v1659882586/blog-media/pgcli-vi-mode-demo.gif)

We can use `hjkl` keybinding for movement and other normal mode keymaps like `w` for moving a word, or `b` for moving a word in the backward direction, `e` to the end of the word, etc. We can use `i` or `a` for entering into insert mode, `dd` to delete the entire line, `cc` or `S` to delete the current line and enter into insert mode, `yy` to yank line, and `p` to paste, `u` for undo, etc.

We can even execute shell commands by using the `!` in normal mode where we will be prompted for the command executed in the same window.

### Autocompletion in PGCLI shell

PGCLI shell has an auto-completion feature that will provide some suggestions based on the recent queries or commands executed or the default options. You can toggle the autocompletion with the `F2` key, it is officially called as `Smart-Completion` as you would see in the bottom left corner of the shell.

### Select an Editor to write queries

We can use different editors for writing SQL queries. We can write a query from an editor by using the `` command. The default editor is chosen which will be set in one of your `$EDITOR` or `$VISUAL` environment variables. You can set them manually in the bashrc file by exporting the value of `VISUAL` or `EDITOR` as the name of your editor.

```bash
export VISUAL=vim

OR

export EDITOR=vim

OR

export VISUAL=gedit
```

Here, we can set `VISUAL` environment variable as the name of the editor or the path to it. The `VISUAL` environment variable is preferred first. For the details, we can check out the [editor command documentation](https://www.pgcli.com/editor).

Once, we have configured the editor, we will be inside the editor after we enter the `` command. We can write queries inside the editor and quit after we are done, once the editor is closed, the text is read from the editor and entered in the psql shell. This will make us wait for the enter command for executing the query. 

We can even execute more than one query at a time if we use the `;` at the end of each query. Also, we can press `v` in VI-normal mode, to open the current query in the default editor.

### Save the query output to a file

Let's say we have written a query inside the psql shell and we want the output to be saved in a file. We can achieve that with the `\o` command, we need to specify the filename in which the output of the queries will be saved. 

**NOTE: Once we use the `\o` command, all the queries will start appending the results in the provided file. To stop the behavior, you can use `\o` (without any filename) to stop appending the results to a file.**

```
# Start recording the output of queries to file
\o filename


# Stop recording the output of queries to file
\o
```

![PGCLI Query output to file](https://res.cloudinary.com/techstructive-blog/image/upload/v1659884242/blog-media/pgcli-output-query.gif)

### Execute query from a file

We can even execute queries stored in a file, we can use the `\i` command followed by the file name. This will load in the queries from the file and will display the output of each query.

```
\i file_name
```

![PGCLI query from a file](https://res.cloudinary.com/techstructive-blog/image/upload/v1659883225/blog-media/pgcli-file-query.gif)

## Summary

In this little article, we were able to explore the basics of the PGCLI package in python, with that package we can interact with the Postgres database that can be on your local server or somewhere remotely. With the PSQL-like shell provided by the PGCLI tool, we were able to write SQL queries, execute commands, write/read queries from files, etc. This tool is quite handy for people dealing with postgres databases and who need to view data or schema in that database.

I've personally used this tool in my current internship where I needed to create a few SQL queries to fetch particular data related to various tables and relationships. It saved a lot of time with auto-completion and integration with the text editor for writing longer queries, especially with JOINS and nested queries.

## Conclusion

This was my take and experience of using the PGCLI tool and exploring the various features it provides that can be leveraged for quick interactions with a Postgres database. Hopefully, you have learned something from this post, if you have any queries(NOT SQL queries) or feedback, please let me know in the comments or on my social handles. Thank you for reading and Happy Coding :)
