---
type: "sqlog"
title: "SQLite functions: generate_series"
date: 2025-08-21
---

The [generate_series](https://www.sqlite.org/series.html) is a table valued function in sqlite and is available via the [generate_series](https://sqlite.org/src/artifact?ci=trunk&filename=ext/misc/series.c) extension.

> The valued function is something that returns a table but is virtual (doesn't really have data or schema in it). It has hidden columns which are used as parameters to the function to constrain the output and generate the data according to those parameters to the function.

The basic operation with `generate_series` would look like this:

```sql
SELECT * FROM generate_series(1, 5);
```

```
sqlite> SELECT * FROM generate_series(1, 5);
+-------+
| value |
+-------+
| 1     |
| 2     |
| 3     |
| 4     |
| 5     |
+-------+
sqlite>
```

Pretty neat, it generated a series of numbers (integers) from 1 to 5 (5 included).

Now are they really integers or?

SQLite is wired, it doesn't really have types, but they are integers as default so to confirm:

```sql
SELECT typeof(value) FROM generate_series(1, 5);
```

```
sqlite> SELECT typeof(value) FROM generate_series(1, 5);
+---------------+
| typeof(value) |
+---------------+
| integer       |
| integer       |
| integer       |
| integer       |
| integer       |
+---------------+
```
OK! Cool, now let's take this to a next level

```sql
SELECT char(value+64) FROM generate_series(1, 5);
```

```
sqlite> SELECT char(value+64) FROM generate_series(1, 5);
+----------------+
| char(value+64) |
+----------------+
| A              |
| B              |
| C              |
| D              |
| E              |
+----------------+
```
OK! That is simple ascii and type casting shenanigans!

We cast the value which would be integer 1 to 5 and add 64 to it to make it in the range 65 to 70 which would be ASCII equivalent of A to E, you could do the same thing but different parameters:

```
SELECT char(value) FROM generate_series(65, 70);
```

```
sqlite> SELECT char(value) FROM generate_series(65, 70);
+-------------+
| char(value) |
+-------------+
| A           |
| B           |
| C           |
| D           |
| E           |
| F           |
+-------------+
```

Same result!

But, ok, this is other rabbit hole, let's focus back on generate_series, the parameters are start and stop right?
Well there is one more, step. This will be like the incremental number for each step.

If you are familiar with C styled for loops, you know this as the increment part, or the `i++` bit.

```sql
SELECT value FROM generate_series(0, 20, 5)
```

```
sqlite> SELECT value FROM generate_series(0, 20, 5);
+-------+
| value |
+-------+
| 0     |
| 5     |
| 10    |
| 15    |
| 20    |
+-------+
sqlite>
```
So, the third parameter increments the counter by that number. So we start from 0 and increment 5 for each step (row) till 20.

Default step value is 1, however the quirk is 0 is still considered as 1.

You can also go reverse

```sql
SELECT * FROM generate_series(0, 20, -5);
```

This will generate number series as 20,15,10,5,0: Note that the step is -1 not the stop value, this will basically flip the start and stop values, so we will start from 20 and decrement by 5 at each step.

If you want to generate a series in negative ranges, the start or stop needs to modified accordingly, the start and stop control which numbers you iterate and the step controls in which direction it moves from start to end.

```sql
SELECT * FROM generate_series(-20, 20, -5);
```

This will generate the number series from 20 to -20, start from 20 and decrement by 5 each step until we reach -20. So we had a start as -20 but why do we start from 20? well because the step is negative, we need to start from the higher value and end at the lower value, hence the start and stop will be flipped for a step value < 0. 

```sql
SELECT * FROM generate_series(10) LIMIT 5;
```

This gave 10,11,12,13,14 right, how? Because there is not just one column in the generate_series virtual table there are more

Apart from value that is the integer we get from the function as the row (column cell)

there are start, stop, step and rowid  columns which are hidden columns in the `generate_series` function or virtual table 

Let’s see

SELECT rowid, start, stop, step, value FROM generate_series(0,10,2);

```
+-------+-------+------+------+-------+
| rowid | start | stop | step | value |
+-------+-------+------+------+-------+
| 1     | 0     | 10   | 2    | 0     |
| 2     | 0     | 10   | 2    | 2     |
| 3     | 0     | 10   | 2    | 4     |
| 4     | 0     | 10   | 2    | 6     |
| 5     | 0     | 10   | 2    | 8     |
| 6     | 0     | 10   | 2    | 10    |
+-------+-------+------+------+-------+
```

The start, stop, step values remain same, as those are the parameters those can’t change but notice the rowid and the value the first and the last column in the result set do change.

Interestingly the start, stop, and step are the parameters to the function, so you can technically pass them however you like 

```sql
SELECT * FROM generate_series() WHERE start = 10 AND stop = 20 AND step = 5;
```
This will give 10,15,20 since the start is 10, increment by 5 until 20. The step is optional, note that stop is optional too, but never forget to provide it, it might cause a forever loop and keep on incrementing by 1 and never give the result back until your computer crashes.


