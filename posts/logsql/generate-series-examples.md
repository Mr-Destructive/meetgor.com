---
type: "logsql"
title: "SQLite functions: Generate_Series Examples"
date: 2025-08-22
---

After taking a look at what the Generate Series function does, now let's see some examples that might be handy.

## Simple sequence

To generate a series of numbers from 1 to 5.

```sql
SELECT * FROM generate_series(1, 5);
```

## Incremental sequence

To generate a series of numbers from 0 to 50 in steps of 10.

```sql
SELECT * FROM generate_series(0, 50, 10);
```

This will start from 0, increment by 10 till 50.

```
0
10
20
30
40
50
```


## Backwards sequence

To generate a series from 50 to 0 in steps of 10

```sql
SELECT * FROM generate_series(0, 50, -10);
```

This would start from 50 (becuase -ve step will flip the start and stop parameters) and decrement by 10 till 0

```
50
40
30
20
10
0
```

## Random Numbers

To generate 5 random number between 1 to 5.

```sql
SELECT * FROM generate_series(1, 5) ORDER BY RANDOM();
```

This would generate 5 numbers 1 to 5, and the ordering will be random. This is becuase for each row, it will generate a random number between -max int to +max int and then order them based on those numbers.

OR

To generate 5 random number between 1 and 10.

```
SELECT ABS((random()%10)) + 1 FROM generate_series(1, 5);
```

This would generate 5 random numbers between 1 and 10. The random function generates a integer between -max int to +max int and then mod it to 10 (or the range you want to generate up to) this will leave us with -9 to 9 numbers but it also be negative so we make it ABS which will make it positive, and add one to make the 0 offset to 1.

```
3
9
10
7
5
```


## Random Characters

To generate 5 random characters from A to Z.

```sql
SELECT char((random()%26)+65) FROM generate_series(1, 5);
```

This would generate a random integer between -max int and + max int, mod it with 26 to get that number between 0 to 25, add 65 which will then can be casted to char in order to render it as an ASCII String.


## Date Ranges

To generate dates for a month

```sql
SELECT date('2025-08-01', '+' || (value-1) || ' day') as date_val
FROM generate_series(1, 31);
``` 

This will take the start date 2025-08-01 and add the value -1 (the value will be from 1 so we need to omit the first day) as a day. This will then generate a series of dates from 2025-08-01 to 2025-08-31.

To generate first day of each month

```sql
SELECT date('2025-08-01', '+' || (value-1) || ' month') as date_val
FROM generate_series(1, 12);
```


## Gap filling in existing ids

To fill in the missing ids in a sequence of ids

```sql
SELECT * FROM generate_series(1, 5) WHERE value NOT IN (1, 3, 5);
```

Here you could imagine the not in list could be another table with the result set only with the ids.


