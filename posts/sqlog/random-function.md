---
type: "sqlog"
title: "SQLite Scalar Function: Random"
date: 2025-08-23
---

The random function in sqlite is quite handy to generate pseudo random numbers.

The random function returns a value between `-9223372036854775807` and `+9223372036854775807`

```sql
SELECT CAST(power(2,63) AS INTEGER);
SELECT CAST(-power(2,63) AS INTEGER);
```
Note that it doesn't return the max integer as `9223372036854775808` becuase using those might break where integer limit might overflow.

This was actually documentated incorrectly before 25th July 2025, that's quite recent.

We can use that to get absolute values, that is to avoid negative values and then mod (divide and get the remainder) it with the max number we want to generate upto.

Example

```sql
SELECT random() as random_number; 
```
```
random_number
7855057830251041076
```

If we want numbers between specific range then use this

```sql
SELECT abs(random() % 10) + 1 as random_number; 
```
```
random_number
6
```
Here 10 is the max number, 1 is the minimum number in the range, if you want negative as well as positive values, you can remove the abs function

```sql
SELECT (random() % 10) + 1 AS random_number
    FROM generate_series(1,20);
```

This will generae random numbers between - 10 and 10, 20 such numbers. We'll leverage the generate series function that I learned last day.


