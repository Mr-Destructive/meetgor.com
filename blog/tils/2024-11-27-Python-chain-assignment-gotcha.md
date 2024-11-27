---
templateKey: til
title: "Gotcha with Chained Assignment in Python"
description: "A lesson learned about Python's chained assignment with mutable objects, where all variables store references to the same object, leading to unexpected behaviour when one is modified."
status: published-til
slug: python-chain-assignment-gotcha
tags: ["python",]
date: 2024-11-27 22:15:00
---

I was writing some Python code and wanted to initialize a few variables to an empty list. Instead of creating separate lists for each variable, I decided to use chained assignments like this:

```python
a = b = c = [1,2,3]
```

It seems okay, nothing new, we are just assigning a, b, and c as an empty list. But notice the problem if you can.

If we try to change the value of b, what do you think is going to happen?

```python
b.append(4)
```

If you didn’t know about how Python handles assignments, you might expect the variables to behave like this:

* `a` should still be `[1, 2, 3]`
    
* `b` should be `[1, 2, 3, 4]` (since I modified b)
    
* `c` should remain `[1, 2, 3]`
    

```python
# You might think it will be this
# a = [1,2,3]
# b = [1,2,3,4]
# c = [1,2,3]
```

But, hello python, it hides the pointer magic behind this statement

The actual and expected state of the variables is:

```python
# But it is actually this
# a = [1,2,3,4]
# b = [1,2,3,4]
# c = [1,2,3,4]
```

When you chain the assignment like this, all three variables refer to the same list object in memory. This means, that all the variables will hold the same object, this means that if you change any of the variables assigned that reference it will change the object, and that will result in changing all the variables since they are referring to the same variable.

You are not creating three independent lists. Instead, all three variables (a, b, and c) are referencing the same list object in memory. They don’t hold copies of the list, they all point to the same object. In Python, this is called **reference assignment**.

When you modify one of the variables (like appending to b), you’re not modifying just b, you’re modifying the single list object that all three variables are referencing. Since all three variables point to the same list, any change you make to the list will be reflected in all three variables.

On the other hand, if I do the same thing with strings, like this:

```python
a = b = c = "hello"
b = "world"

# a = "hello"
# b = "world"
# c = "hello"
```

This only mutates the b variable and not the a and c, since string is not a mutable object in Python.

In Python, the objects are either [mutable or immutable](https://realpython.com/python-mutable-vs-immutable-types/)

Some of the primitive data types are:

Mutable Types:

* List
    
* Dictionaries
    
* Set
    
* Byte Array
    

Immutable Types:

* Integer, Float, Complex
    
* String
    
* Tuple
    
* Bytes
    
* Boolean
    
* Frozenset
    

So, suppose you assign multiple variables with the same value of a mutable type. In that case, the change in one variable will mutate the other variables as well since the underlying object in memory is the same.

So, this is what I learned from the mistake, avoid the chining assignment when dealing with mutable objects

Instead do the following:

```python
a, b, c = [], [], []
b.append(256)

# a = []
# b = [256]
# c = []
```

This is safe and the right way to assign variables to individual values instead of the same value being referred to by all the variables.

Happy Coding :)
