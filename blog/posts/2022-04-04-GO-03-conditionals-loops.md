---
templateKey: blog-post
title: "Golang: Conditionals and Loops"
description: "Getting familiar with condition statements and loops in golang."
date: 2022-04-04 09:00:00
status: published
slug: golang-conditionals-loops
tags: ['go',]
image_url: https://meetgor-cdn.pages.dev/golang-003-conditionals-loops.png
series: ['100-days-of-golang',]
---

## Introduction

Moving to the fourth part, we will be doing conditional statements and loops in golang. We will be seeing the basics of conditional statements like if-else and switch along with loops like for, while, and range-based loops. We won't be covering iterating over arrays in a loop as this requires an understanding of arrays. 

## Conditional statements

Conditional statements are quite a fundamental aspect of learning a programming language. In golang, we have if-else conditional statements as well as switch cases. We will be exploring both of them in this section. Firstly, we will dive into if-else statements which are quite easy to understand.

### if else

An `if` statement is used for checking the validity of a condition. If the condition is true(condition is met), a particular sets of statements are executed else (condition is not satisfied) different statements gets executed. We can use a basic `if-else` statement in go with the following syntax:

```go
if condition {
    // statements
}else{
    //statements
}
```

```go
package main

import "fmt"

func main() {
    age := 16
    if age < 13{
        fmt.Println("Kid")
    }else{
        fmt.Println("Teenager")
    }
}
```

```
$ go run if_else.go
Teenager
```

We can also use else if for evaluating more than one condition rather than using nested if and else. 

```go
if condition {
    // statements
}else if condtion {
    //statements
}else if condition {
    //statements
}else {
    //statements
}
```

```go
year := 2

if year < 1 {
    fmt.Println("Freshman")
} else if year == 2 {
    fmt.Println("Sophomore")
} else if year == 3 {
    fmt.Println("Junior")
} else if year == 4 {
    fmt.Println("Senior")
} else {
    fmt.Println("Probably a Graduate")
}
```

```
$ go run if_else.go
Sophomore
```

Using `else if` we can evaluate multiple conditions. This style is much better than using nested `if else` statements as it becomes harder to read for complex cases.


### switch 

We also have switch statements in golang which allow us to write cases for a given state of a variable. We can simply add cases for a given variable, the case should be a valid value that the variable can take. If a case is matched it breaks out of the switch statement without executing any statements below the matched case.

The basic structure of the switch statements in golang is as follows:

```go
switch variable{
case value1:
    //statements
case value2:
    //statements

}
```

```go
age := 19
var result string
switch {
case age < 13:
    result = "Kid"
case age < 20:
    result = "Teenager"
case age < 25:
    result = "Adult"
case age < 35:
    result = "Senior"
}
fmt.Println("The person is a",result)
```

```
$ go run switch.go
The person is a Senior with age 27.

$ go run switch.go
The person is a Teenager with age 19.

$ go run switch.go
The person is a Kid with age 11.
```

This gives a good understanding of switch-case statements. We can give a variable to the switch statement and pick its value in the respective case statements to evaluate the result accordingly. The `default` statement is evaluated when there is no match among the given cases. 

```go
language := ""
var devs string
switch language{
case "go":
    devs = "gopher"
case "rust":
    devs = "rustacean"
case "python":
    devs = "pythonista"
case "java":
    devs = "Duke"
default:
    language = "javascript"
    devs = "developer"
}
fmt.Println("A person who codes in",language,"is called a",devs)
```

```
$ go run switch.go
A person who codes in javascript is called a developer

$ go run switch.go
A person who codes in python is called a pythonista

$ go run switch.go
A person who codes in go is called a gopher
```

This code will by default pick `javascript` and `developer` as the values for `language` and `devs` respectively if there is no match for the provided language or the language is left empty. 

We also have `fallthrough` in the golang switch which allows evaluating more than one case if one of them is matched. This will allow the switch and check for all the cases sequentially and evaluate all the matched and satisfied cases. 

```go
character := 't'
fmt.Printf("The input character is = %c\n", character)
switch {
case character == 97:
    fmt.Printf("Its %c\n", character)
    fallthrough
case character < 107 && character > 96:
    fmt.Println("It's between a and k")
    fallthrough
case character < 117 && character > 98:
    fmt.Println("It's between a and t")
    fallthrough
case character < 122 && character > 98:
    fmt.Println("It's between u and z")
default:
    fmt.Println("Its not a lowercase alphabet")
}
```

```
$ go run switch.go
The input character is = f
It's between a and k
It's between a and t
It's between a and u

$ go run switch.go
The input character is = k
It's between a and t
It's between a and u

$ go run switch.go
The input character is = a
Its a
It's between a and k
It's between a and t
It's between a and u
```

So, here we can see that the fallthrough hits multiple cases. This is unlike the base case which exits the switch statement once a case has been satisfied. This can be helpful for situations where you really want to evaluate multiple conditions for a given variable.

## Loops

We can now move on to loops in golang. We only have a `for` loop so to speak but this can be used as any other looping statement like the `while` loop or range-based loop. We will first see the most fundamental loop statement in golang which is a three-component loop. 

### for loop

We can have a simple for loop in golang by using the three statements namely `initialize`, `condition`, and `increment`. The structure of the loop is quite similar to the other programming languages.

```go
for k := 0; k < 4; k++ {
    fmt.Println(k)
}
```

### Range-based loop

We can even iterate over a string, using the range keyword in golang. We need to have two variables for using a range-based for loop in golang one is the index or the 0 based position of the element and the copy of the element in the array or string. Using the range keyword, we can iterate over the string one by one. 

```go
name := "GOLANG"
for i, s := range name{
    fmt.Printf("%d -> %c\n",i, s)
}
```

```
$ go run for.go
0 -> G
1 -> O
2 -> L
3 -> A
4 -> N
5 -> G
```

So, here we can see we have iterated over the string by each character. Using the range keyword in golang, The `i, s` is the index and the copy of the element at that index as discussed earlier. Using the index we get the value which we don't have to index the array for accessing it, that is already copied in the second variable while using the range loop. 

### while loop (Go's while is for)

There are no while loops as such in golang, but the for loop can also work similarly to the while loop. We can use a condition just after the for a keyword to make it act like a while loop. 


```go
for condition {
    // statements
}
```

```go
count := 3
for count < 9 {
	fmt.Println(count)
	count++
}
```

```
$ go run while.go
3
4
5
6
7
8
```

We can see here that the condition is evaluated and the statements in the loop body are executed, if the condition evaluates to false, the flow is moved out of the loop and we exit the loop. 

### Infinite loop

We can run an infinite loop again using a keyword. We do not have any other keywords for loops in golang. 

```go
for {
    // statements
    // should have conditons to exit
}
```

```go
flag := 4
for {
	flag++
	fmt.Println(flag)
}
```

This might be used with a base condition to exit the loop otherwise there should be a memory overflow and the program will exit with errors.

### Break 

If we want to exit out of a loop unconditionally, we can use the `break` keyword. This will break the loop and help us to exit out of an infinite or a condition-bound-based loop too.

```go
flag := 1
for {
    fmt.Println(flag)
    flag++
    if flag == 7 {
        fmt.Println("It's time to break at", flag)
        break
    }
}
```

```
$ go run infinite.go
1
2
3
4
5
6
It's time to break at 7
```

As, we can see inside an infinite loop, we were able to break out of it by using a conditional statement and `break` keyword. This also applies to switch cases it basically is the opposite of `fallthrough` in switch-case statements. But by default(without using fallthrough), the case statement breaks the switch after a match has been found or the default case has been encountered. 

### Continue

We also have the opposite of `break` i.e. `continue` which halts the execution of the loop and directs back to the post statement increment(in case of for loops) or evaluation(in case of while loop). We basically are starting to iterate over the loop again after we encounter the continue but by preserving the counter/iterator state values. 

```go
counter := 2
for counter < 4 {
    counter++
    if counter < 4 {
        continue
    }
    fmt.Println("Missed the Continue? at counter =", counter)
}
```

```
$ go run infinite.go
Missed the Continue? at counter = 4
```

For following up with the code for this and all parts of the series, head over to the [100 days of Golang](https://github.com/mr-destructive/100-days-of-golang) GitHub repository. 

## Conclusion

So, from this section, we were able to understand the basics of conditional statements and loops in golang. We covered the things which are more important for understanding the core of the language than some specific things. There are certain parts that need to be explored further like iterating over arrays and slices which we'll cover after we have understood arrays and slices. Hopefully, you have understood the basics of the conditional statements and loops in golang. Thank you for reading, if you have any questions, or feedback please let me know in the comments or social handles. Until next time, Happy Coding :)


