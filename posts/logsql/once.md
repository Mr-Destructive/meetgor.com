---
type: "logsql"
title: "SQLite Dot command: Once"
date: 2025-08-20
---

The [once](https://sqlite.org/cli.html#writing_results_to_a_file) dot command is similar to the [output](https://sqlite.org/cli.html#writing_results_to_a_file), however the distinction is that it is limited to the very next SQL query and not all the subsequent queries.

The primary way to use `once` is either by specifying the specific file name/path or opening the result set in the system editor.

### Output to a file

```
.once somefile.txt
```

Once this is set, any query you execute, the result set of it will be logged to the specified file.

```sql
SELECT printf("Hello %s! Bye, %s!", "world", "mars");
```

```
$ cat somefile.txt
Hello world! Bye, mars!
```

Pretty cool! But there's more

There are three options available to perform different things for where this output can go

- Into a temporary file opening in a system text editor with -e
- Into a temporary excel/csv/xlsx file in a system spreadsheet editor application with the -x option
- Into an excel/csv/xlsx file (compatible with Microsoft Excel) containing utf-8 character or symbols with the --bom option.


### Text Editor

You can open the result set to a temporary file in a system text editor with the -e option.

```
.once -e
```

You can set the system editor with 


### Spreadsheet Editing Application

Also you can open the result set into a temporary xlsx or csv file in a system spreadsheet editing application like Microsoft Excel, LibreOffice Calc, etc.

```
.once -x
```


### Unicode encoded Spreadsheet

This option is used to create an excel/csv/xlsx file (compatible with Microsoft Excel) containing utf-8 character or symbols. If you used -x on Linux, it would work fine there, but the file format won't be compatible in the Microsoft Excel or other xlsx formats. To make sure the utf-8 characters are rendered and parsed properly in the excel file, just use the --bom which stands for byte order mark. This options adds certain bytes at the beginning of the file to make the application understand which encoding to use while rendering to use, like `EF BB BF` for utf-8, `FE FF` for utf-16, and so on.

```
.mode csv
.once --bom filename.xlsx
```


