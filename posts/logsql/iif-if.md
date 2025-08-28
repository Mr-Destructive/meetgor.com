---
type: "logsql"
title: "SQLite Scalar function: 3 valued iif and if scalar function"
date: 2025-08-27
---

In the recent SQLite version specifically the 3.48.0, the support for pair of conditions was added to the iif function. And also an alias for iif as if was added for compatibility and ease of use for people coming from other databases.

## IIF Before 3.48
Before this release it was kind of 3 way /ternary operation kind of thing like this

### Simple IF THEN ELSE

We can use the IIF function by passing 3 paraeters, first is the condition to check, the second is the value to return if the condition is true and the third is the value to return if the condition is false.

```sql
SELECT iif(1=0, 'one is zero', 'one is not zero');
```

```
one is not zero
```

### Ternary like nested IIF

Optionally the third condition can be nested with another IIF as it will be a value in the end. So, we can use IIF like this

```sql
SELECT iif(7%2=0, 'even', iif(7%3=0, 'multiple of 3', 'prime')) AS is_prime;
```

```
+---------+
| is_prime|
+---------+
| prime   |
+---------+
```

That is quite simple but it gets little messed up for more than one check, take this for instance

```sql
SELECT iif(2025 % 400 = 0, 'leap',
       iif(2025 % 100 = 0, 'not leap',
       iif(2025 % 4 = 0, 'leap', 'not leap'))) AS  is_leap_year;
```

```
+--------------+
| is_leap_year |
+--------------+
| not leap     |
+--------------+
```

I have printed it in a nice way, but still might get harder to read and even write, notice the number of closing brackets at the end.

## IIF After 3.48

With the latest 3.48 version, this changes quite a bit, you can have iif or if as like a case when then replacement like so:

```sql
SELECT iif(2025 % 400 = 0, 'leap',
           2025 % 100 = 0, 'not leap',
           2025 % 4 = 0, 'leap', 'not leap') AS is_leap_year;
```

```
+--------------+
| is_leap_year |
+--------------+
| not leap     |
+--------------+
```

How clean is that, with this you can specify, Condition, Value as a pair and it can have N subsequent pairs of this condition - value.  All and this while maintaining backwards compatibility, that’s quite a remarkable thing to push I would say, not easy to handle these.

Now it has 3 variants
- 3 valued expression (single expression, could be nested too, but three expressions, like →  condition, true, false)
- 2 valued expression (assumes the false value is NULL)
- N condition - valued pair expression (two pair of condition and value)

These is neat and makes a lot of sense.

Let’s see the same example in all 3 variants

### 3 Valued Expression

This is the most simple one

```sql
SELECT iif(2025%4=0, 'probably leap', 'not leap') AS is_leap_year;
```

```
+--------------+
| is_leap_year |
+--------------+
| not leap     |
+--------------+
```
 
This returned the false value which happens when the condition evaluated to false.


### 2 Valued Expression

```sql
SELECT iif(2025%4=0, 'probably leap') AS is_leap_year;
```

```
+--------------+
| is_leap_year |
+--------------+
|              |
+--------------+
```
This returned NULL as 2025 is not divisible by 4. If there are 2N parameters to the function, and if all the conditions are false, it will return NULL. And the above example is the base case of it. Remember to have atleast 2 parameters, one the condition and one as the value if that condition evaluates to true. The last parameter is optional and if the condition is false, it would return NULL.


### N Pair Expression

So moving on the final one which is the most readable version of it and allows more stuff

```sql
SELECT iif(2025 % 400 = 0, 'leap',
           2025 % 100 = 0, 'not leap',
           2025 % 4 = 0, 'leap',
           'not leap')
|| ' year' AS is_leap_year;
```

```
+---------------+
| is_leap_year  |
+---------------+
| not leap year |
+---------------+
```

Elegant
Notice the last condition, if you don’t specify that, it will be NULL, so it would be 2N+1 parameters (where N is the number of conditions)

Other way to remember is 

- Even parameters and all conditions are
    - True (the first value) or equivalent corresponding value for that condition
    - False (it will return NULL)
- Odd parameters and all conditions are 
    - True (the first value) or equivalent corresponding value for that condition
    - False (it will return the last parameter which is the false condition)

You can use iif or if interchangeably. Both these do the same thing as the CASE WHEN THEN conditions, I think IIF now is really easy to write as now it supports arguments in pairs

