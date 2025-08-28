---
type: "logsql"
title: "SQLite dot commands: read file or standard output"
date: 2025-08-28
---

## Read dot Command 

The `.read` dot command is a quick handy command to import and run your SQL queries in the current session.

You can just pass a SQL file (usually a query ending with ; it will execute each query one by one)

```bash
.read filename
```

Let’s say this is a sql file containing the schema of a database, just one table users.

```sql
CREATE TABLE users (
  id   INTEGER PRIMARY KEY,
  name TEXT
);
```
Writing this sql query in the `schema.sql` file

If we use the .read command with the name of the file it will execute it line by line, (line meaning terminated by ; here or even any dot commands)

```bash
.read schema.sql
```

This will just execute the query, if there are SELECT statements possibly then will output the result set too.

Let’s add some users to the table with the `insert_users.sql`

```sql
INSERT INTO users(name) VALUES('abc'), ('def'), ('ghi');
SELECT * FROM users;
```

Saving this and running .read insert_users.sql should now have displayed all the 3 inserted users. 

## Reading with dot commands

Let’s create one more file with specific dot commands to render a nice table format of a result set from a select statement

```sql
.mode table
SELECT * FROM users;
```

I save the above contents into a file called `users.sql`

Now if I run the .read command with the users.sql you should see all your users but with the mode as table format. It even takes the dot commands and all of the nice options out there. Be cautious though!

```
.read users.sql
```

```
sqlite> .read users.sql
+----+---------+
| id |  name   |
+----+---------+
| 1  | User_1  |
| 2  | User_2  |
| 3  | User_3  |
| 4  | User_4  |
| 5  | User_5  |
| 6  | User_6  |
| 7  | User_7  |
| 8  | User_8  |
| 9  | User_9  |
| 10 | User_10 |
+----+---------+
```

## Read with standard output

You can also run a script that will generate some SQL commands and execute them. By using the | operator after .read and providing the name of the script. this will run the script and any output of that script (valid sql) will be executed on the shell.

For instance I have a script that would insert 10 users with a bash script like so called `inser_users.sh`

```bash
#!/bin/bash
for i in $(seq 1 10); do
  echo "INSERT INTO users (name) VALUES ('User_$i');"
done
```

If you would use this as is it won’t work as its a bash script, you can’t run those in a sqlite shell. However by using `.read | filename` (shell script), that will indeed do the trick.

The important thing to note here is that it would run all the things outputted from the script, so its kind of running the script and the std out of that script is taken as the input for the read command so like typical pipe operator in linux.

```
.read | insert_users.sh
```

This will run the script and the insert statements will be printed, so it will execute them. So, in essence we are basically piping the output of one command to the other in this case the read command.

```sql
SELECT * FROM users;
```

As you can see it has now inserted all the 10 users.

The `|` command can be followed by any script, not limited to bash, it could be python, or any script or program that can print to the standard output a correct and legal SQL statements.

This is a really great small yet highly important piece of functionality of sqlite shell.

Reference:
- [SQLite CLI: Reading SQL from a File](https://sqlite.org/cli.html#reading_sql_from_a_file)

