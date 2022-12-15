---
templateKey: blog-post
title: "Golang: Variables and Types"
description: "Understanding and creating variables and their types in Golang"
date: 2022-04-03 10:45:00
status: published
slug: golang-variables
tags: ['go',]
image_url: https://meetgor-cdn.pages.dev/golang-002-variables-types.png
series: ['100-days-of-golang',]
---

## Introduction

In the third part of the series, we will be covering the fundamentals for learning any programming language i.e. variables and data types. We will be covering from data types to variable declaration. We won't be seeing each and every detail related to the data types as some of them require a knowledge of loops and other topics, so that can be left for the different part.

## Types in golang

In Golang there are 3 major types : Numeric, Bool and String. Further we also have specific types for the three data types like int, float, rune, byte, etc. We will first see how to declare a simple variable and then explore the data types in Golang. 

```go
var name string
```

This the variable declaration in Golang, we have the keyword `var` similar to Javascript but we optionally have to specify the type of the variable if we are not immediately assigning/defining it a value. 

To assign variable values, we can write the datatype of the assigned value but it is optional as the go compiler will know the datatype according to the assigned value. Further you cannot change the type of that variable once it is initialized.

```go
var name string = "Gopher"

OR 

var name string
name = "Gopher"

OR

var name = "Gopher"

```

We also have `const` for assigning constant values to a particular value set. You cannot change the value for a constant type, doing this will result in compile time error. An important property of `const` can be noted here, if you simply declare a `const` without initializing the value and try to access that constant, the go-compiler will throw an compilation error.

```go
const name string = "David"

OR

const name string
name = "Calvin"

OR

const name = "Smith"
```

By default, the values for string is an empty string`""`, for integer and float it is `0` and for bool it is `false`.

Each of these are valid declaration of variables in golang. Let's now dive into the data types and follow up with variable declaration in detail later.

| Numeric    | String | Bool |
|------------|--------|------|
|            |        |      |      
|  int       | string | bool |
|            |        |      |
|  float     |        |      |
|            |        |      |
|  complex   |        |      | 
|            |        |      | 
|  rune      |        |      |
|            |        |      |  
|  byte      |        |      | 
   
### Numeric

Let's first explore the `numeric` data types in golang as you have guessed correctly, we have `int` and `float` as distinct categories but further we also have fine grained storage types for both of the types. 

#### Integer

In integers as well we have two categories `signed` and `unsigned`, we can basically store only positive integers in `unsigned` integers giving us an extra bit to play with. 

For Integers, we have specific data storage types depending on the bits it can store like `int8` for storing 8 bit integers, `int16` for storing 16 bit integer value, `int32` and `int64` for the given number of bits in the integer. Similarly we will have these storage based integer values for unsigned integers like `uint8`, `uint16`, `uint32` and `uint64`. We can basically store double amount of positive integers in unsigned integers as `uint` than in signed integers `int`, this is because the most significant bit is not used as a sign bit since all values in the variable are positive and hence no sign bit is required. 

```go
var likes int = 140
fmt.Println(likes)
```

```
$ go run int.go
140
```

```go
var age int8 = 140
fmt.Println("Age = ",age) 
```

```
$ go run int.go
# command-line-arguments
.\int.go:6:9: constant 140 overflows int8
```

This will give us an error as `140` is above the limit for `int8`. So, unless you have specific requirements as storage limitation, you should be using `int` as the default data type for storing integers.

So, we need to define variables as per the limits to which we are going to use them, if you just specify `int` the type will be selected based on your operating system, if it is `32bit`, it will take `int32`, for `64bit` OSes it will take as `int64` integer. If you define a variable with let say `16` bit storage and if it exceeds the limit for `16` bit storage, it would give a `overflow limit` error. 

Below are the limits for all the integer types in Golang: 

```
uint8 ->  0  to  255
uint16 ->  0  to  65535
uint32 ->  0  to  4294967295
uint64 ->  0  to  18446744073709551615

int8 ->  -128  to  127
int16 ->  -32768  to  32767
int32 ->  -2147483648  to  2147483647
int64 ->  -9223372036854775808  to  9223372036854775807
```

If you want to reality check for the boundary values of this data types, you can code a program in `go` as below: 

- To find the maximum value of uint, we can flip all the bits in `0` to get all the set bits in the integer thus we use `^` operator.
- To find the maximum value for signed integers, we can right shit one bit so as to unset the sign bit.
- The minimum value for uint is the default value `0`.
- The minimum value for int can be calculated by subtracting one from the negative of the max limit.

```go
package main

import (
    "fmt"
)

func main() {
    var min_uint = 0
    var max_uint8 uint8 = ^uint8(0)
    var max_uint16 uint16 = ^uint16(0)
    var max_uint32 uint32 = ^uint32(0)
    var max_uint64 uint64 = ^uint64(0)

    var max_int8 int8 = int8(max_uint8>>1)
    var min_int8 int8 = -max_int8 - 1
    var max_int16 int16 = int16(max_uint16>>1)
    var min_int16 int16 = -max_int16 - 1
    var max_int32 int32 = int32(max_uint32>>1)
    var min_int32 int32 = -max_int32 - 1
    var max_int64 int64 = int64(max_uint64>>1)
    var min_int64 int64 = -max_int64 - 1

    fmt.Println("uint8 -> ", min_uint, " to ", max_uint8)
    fmt.Println("uint16 -> ", min_uint, " to ", max_uint16)
    fmt.Println("uint32 -> ", min_uint, " to ", max_uint32)
    fmt.Println("uint64 -> ", min_uint, " to ", max_uint64)
    fmt.Println("")
    fmt.Println("int8 -> ", min_int8, " to ", max_int8)
    fmt.Println("int16 -> ", min_int16, " to ", max_int16)
    fmt.Println("int32 -> ", min_int32, " to ", max_int32)
    fmt.Println("int64 -> ", min_int64, " to ", max_int64)
}
```

This was the basic overview of Integers in Golang.

Though rune and byte are Integer aliases, we will explore them in the String section as they make a lot of sense over there.

#### Float

Similar to integers, we also have floats in the numeric category. A float is a numeric data type that can allow numbers with decimal values. A simple float can be of either `float32` or `float64`. The two types are defined as the precision of the decimal values. Obliviously, the `float64` type is more precise than the counterpart and is also the default type assigned if not provided.

```go
const percent = 30.5
fmt.Println(percent+50)
```

You optionally provide the `float32` type to have a bit less precision than usual using the `float32` keyword as follows:

```go
const percent float32 = 46.34
fmt.Println(percent - 3.555)
```

The floating value precision of the float types in golang are as follows:

```
float32	  -->   -3.4e+38 to 3.4e+38.
float64	  -->   -1.7e+308 to +1.7e+308.
```

As quite logical reasons, the precision is almost double in the `float64` compared to `float32`. If we try to add(any operation) a `float64` number with a `flaot32`, we get an error as performing operations on two differently stored types can't be operated. 

#### Complex Numbers

We also have complex numbers in golang. This are the numbers which deal with a real and a imaginary numbers. We initialize the complex variable using the `complex` function which takes two parameters the `real` part and the `imagianry` part. Both the parts or numbers are stored as float in the complex data type.

So, since golang has `float32` and `float64` data types, we will have `complex64` and `complex128` as complex types. Since we are storing two `flaot64`, it has a `complex128` type and `complex64` for both parts as `float32`. Here as well, you cannot store the two parts(real and imaginary) as different float types i.e. You need to have both real and imaginary as either `flaot32` or `flaot64`.

```go
var percent = 30.738
var f = 4.545
var comp1 = complex(f, percent)
var comp2 = complex(percent, f)
fmt.Println(comp1 - comp2)
```

```
(-26.192999999999998+26.192999999999998i)
```

Golang automatically adds the `i` or iota in the complex/imaginary part for better readability. 

### Strings 

We can now move onto the `string` data type in golang. It has several data types like `byte`, `rune`, `string`. In golang, `byte` and `rune` store individual characters whereas `string` can hold multiple characters. 

#### Byte

A byte in golang is an unsigned 8 bit integer, which means it can hold numeric data from 0 to 255. So how is this displaying characters if it stores integer. Well, because each number it stores is mapped to the ASCII character set which is used to represent characters. 

A byte can be stored in a single quote `''`, if we use double quotes`""`, the variable is considered as string if we aren't specifying the data type.

```go
const c byte = 't'
fmt.Println(c)
```

```
$ go run byte.go
116
```

This gives the output as a number between 0 and 255 depending on the character which you have stored. To print the actual character you need to type caste into a string like:

```go
const c byte = 't'
fmt.Printf("Character = %c \nInteger value = %d\n", c, c)
```

```
$ go run byte.go
Character = t
Integer Value = 116
```

We can get the character equivalent of the byte representation number using the [Printf](https://cs.opensource.google/go/go/+/go1.18:src/fmt/print.go;l=212) function and the `%c` place holder for a character. The `\n` is used to end the line just for displaying proper output.

We can even store numbers between `0` and `255` as it is internally an `uint8`.

#### Rune

A rune is extended type of byte as it stores `int32` numbers and hence it represents `Unicode` characters. This is the default type if you do not specify `byte` and use single quotes to assign a character. Using rune, we can assign it an unicode characters to display some rich characters and literals like emoji or expressions.

```go
var smiley_emoji = '\u263A'
fmt.Printf("Smiley Emoji --> %c", smiley_emoji)
```

![GO Rune Smiley Emoji](https://res.cloudinary.com/techstructive-blog/image/upload/v1648962460/blog-media/obw9ihlxsvhytbe8ito3.png)

So, rune is pretty amazing type to play with characters in golang. As it is a default type assigned against byte if not provided while assignment. 

#### String

Strings are basically a slice(list) of bytes. Each character in a string is a byte. By default the string will be empty if you don't initialize it with a value. 

```go
const language string
language = "C++"

OR

const language string= "Python"

OR

const language = "Javascript"
```

We can even access particular character in the string using it's index.

```go
const code = "12AB34CD"
fmt.Println(code[6])
```

```
$ go run string.go
67
```

This gives us a integer as we are accessing the byte from the string using its index. Thus, we can use `%c` in the `Printf` function to format and print the equivalent characters of the byte.

```go
const code = "12AB34CD"
fmt.Printf("2nd Character in string = %c\n", code[4])
```

```
$ go run string.go
2nd Character in string = A
```

We can also declare strings using backticks/backquotes or whatever you call it (```), assigning string with this allows us to write multi line string.   

```go
var statement = `This is the first line
The next line
The last line`

fmt.Println(statement)
```

```
$ go run str-backticks.go
This is the first line
The next line
The last line
```

Further in the loop article we will see how to loop/iterate over a string.

### Boolean 

This type is used to store either `true` or `false` in golang. The default value of a boolean variable is `false`.

```go
var power bool
fmt.Println(power)
```

```
$ go run bool.go
false
```

We can assign the variable as either `true` or `false`.  

```go
const result = true
fmt.Printf("The statement is %t", result)
```

```
$ go run bool.go
The statement is true
```

So, using the `%t` we can print the value of a boolean value in golang in the `Printf` function.

## Creating Variables 

Now, we have familiar with data types in golang, we can more expressively create, declare, initialize variables in golang.

There are 3-4 primary ways to define a variable most of which we have already seen.

### Declaring a Variable 

We can declare a variable without assigning it any value but for that we need to then provide the data type, this can be done using the following command:

```go
var expereience int

expereience = 2
```

We can even use `const` for constant value in the given scope. 

Here, we can even declare multiple variables by separating each variable/constant with comma `,` which can be done as follows:

```go
var a, b, c int

OR

const i, j, k bool
```

### Defining and Initializing at the same time

We can initialize a variable/constant in golang by explicitly giving it a value. We can do that by using `var` for variable value or `const` for a constant value. We can optionally provide the data type at this moment as golang will automatically detect the type and assign the memory according to the value given.

```go
var place string = "home"
```

Here, there is no compulsion to provide the `datatype` as the compiler will be able to know it from the asisgned value. Though if you want to provide a non-default value you can specify the datatype. 

### Declaring Multiple Variables

We can assign multiple variables at once by separating them with comma`,`. The variable name to the left and the values to the right needs to separated with comm on both sides.

```go
var x, y, z = 100, '#', "days"

fmt.Printf(" x = %d \n y = %c \n z = %s \n",x,y,z)
```

```
$ go run multiplvar.go
 x = 100
 y = #
 z = daysofcode
```

We can are declaring and assigning multiple variables, the `x` variable is assigned an integer value, `y` with a `rune`(by default) and `z` with a string. We are using `Printf` function with place holders for int `%d`, rune/byte `%c` and string as `%s`. The `\n` is for a new line.

If we want to assign the variables with a particular data type, we can use the var keyword as a list of values.

```go
var (
    x int8 = 100
    y byte = '#'
    z =  "daysofcode"
)

fmt.Printf(" x = %T \n y = %T \n z = %T \n",x,y,z)
```

```
$ go run multiplvar.go
 x = int8
 y = uint8
 z = string
```

This is not only limited to `var` we can also use `const` to declare multiple constants with type constraint. Also, note we are using the `%T` placeholder for getting the type of the data stored in the variable.

Also, we can define(declare and initialize) multiple variable with same data type with command separated as follows:

```go
var pi, e, G float32 = 3.141, 2.718, 6.67e-11   
var start, end byte = 65, 90
fmt.Println(pi, e, G)
fmt.Printf("%c %c\n",start, end)
```

```
$ go run multp.go
3.141 2.718 6.67e-11
A Z
```

### Assigning Variable using Walrus Operator (Shorthand Declaration)

We can skip usign `var` or the `datatype` by using the `:=` walrus operator. This type of assignment using `walruns` operator can only be allowed in the function body and not anywhere else, in the global scope this type of assignment is not allowed.

```go
place := "school"
```

This is such a simple shorthand for assigning variables though only in a function body.

Also, multiple variable declaration is possible with walrus operator.

```go
x, y, z := "foo", 32, true
fmt.Println(x, y, z)
fmt.Printf("%T %T %T", x, y, z)
```

```shell
$ go run walrus.go
foo 32 true
string int bool
```

Links to all code and links are visible on the [GitHub](https://github.com/Mr-Destructive/100-days-of-golang) repository.

## Conclusion

So, from this part of the series, we were able to understand variables and the various data types in Golang. Though we didn't got too much in detail still we can find ourselves a bit comfortable in understanding basic go scripts. In the next section, we will looking into conditional statements and loops. This would give a good grasp on iterating over a string and even learn arrays(just the basics) we will explore Arrays and slices(remember strings?) after that. 

So, if you have any questions, suggestions, or feedback please let me know in the comments or on the social handles. See you next time, Happy Coding :)
