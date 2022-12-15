---
templateKey: blog-post
title: "Golang: Operators"
description: "Understanding the basics of operators like arithmetic, logical, bitwise, assignment operators in Golang"
date: 2022-05-07 20:33:38
status: published
slug: golang-operators
tags: ['go',]
image_url: https://meetgor-cdn.pages.dev/golang-013-operators.png
series: ['100-days-of-golang',]
---

## Introduction 

In this 13th part of the series, we will be exploring the fundamentals of operators in Golang. We will be exploring the basics of operators and the various types like Arithmetic, Bitwise, Comparison, Assignment operators in Golang.

Operators are quite fundamentals in any programming language. Operators are basically expressions or a set of character(s) to perform certain fundamental tasks. They allow us to perform certain trivial operations with a simple expression or character. There are quite a few operators in Golang to perform various operations.

## Types of Operators

Golang has a few types of operators, each type providing particular aspect of forming expressions and evaluate conditions.

1. Bitwise Operators
2. Logical Operators
3. Arithmetic Operators
4. Assignment Operators
5. Comparison Operators

### Bitwise Operators

Bitwise Operators are used in performing operations on binary numbers. We can perform operation on a bit level and hence they are known as bitwise operators. Some fundamental bitwise operators include, `AND`, `OR`, `NOT`, and `EXOR`. Using this operators, the bits in the operands can be manipulated and certain logical operations can be performed. 

```go
package main

import "fmt"

func main() {
	x := 3
	y := 5
	// 3 -> 011
	// 5 -> 101
	fmt.Println("X AND Y = ", x & y)
	fmt.Println("X OR Y = ", x | y)
	fmt.Println("X EXOR Y = ", x ^ y)
	fmt.Println("X Right Shift 1  = ", x >> 1)
	fmt.Println("X Right Shift 2  = ", x >> 2)
	fmt.Println("Y Left Shift 1 = ", y << 1)
}
```

```
$ go run bitwise/main.go

X AND Y =  1
X OR Y =  7
X EXOR Y =  6
X Right Shift 1  =  1
X Right Shift 2  =  0
Y Left Shift 1 =  10

```

We use the `&` (AND operator) for performing AND operations on two operands. Here we are logically ANDing `3` and `5` i.e. `011` with `101` so it becomes `001` in binary or 1 in decimal.

Also, the `|` (OR operator) for performing logical OR operation on two operands. Here we are logically ORing `3` and `5` i.e. `011` with `101` so it becomes `111` in binary or 7 in decimal.

Also the `^` (EXOR operator) for performing logical EXOR operation on two operands. Here we are logically EXORing `3` and `5` i.e. `011` with `101` so it becomes `110` in binary or 6 in decimal.

We have a couple of more bitwise operators that allow us to shift bits in the binary representation of the number. We have two types of these shift operators, right sift and left shift operators. The main function of these operator is to shift a bit in either right or left direction. 

In the above example, we have shifted `3` i.e. `011` to right by one bit so it becomes `001`. If we would have given `x >> 2` it would have become `0` since the last bit was shifted to right and hence all bits were 0.

Similarly, the left shift operator sifts the bits in the binary representation of the number to the left. So, in the example above, `5` i.e. `101` is shifted left by one bit so it becomes `1010` in binary i.e. 10 in decimal. 

This was a basic overview of bitwise operators in Golang. We can use these basic operators to perform low level operations on numbers.

### Comparison Operators

This type of operators are quite important and widely used as they form the fundamentals of comparison of variables and forming boolean expressions. The comparison operator is used to compare two values or expressions. 

```go
package main

import "fmt"

func main() {
	a := 45
	b := 12
	fmt.Println("Is A equal to B ? ", a == b)
	fmt.Println("Is A not equal to B ? ", a != b)
	fmt.Println("Is A greater than B ? ", a > b)
	fmt.Println("Is A less than B ? ", a < b)
	fmt.Println("Is A greater than or equal to B ? ", a >= b)
	fmt.Println("Is A less than or equal to B ? ", a <= b)

```

```
$ go run comparison/main.go

Is A equal to B ?  false
Is A not equal to B ?  true
Is A greater than B ?  true
Is A less than B ?  false
Is A greater than or equal to B ?  true
Is A less than or equal to B ?  false
```

We use simple comparison operators like `==` or `!=` for comparing if two values are equal or not. The expression `a == b` will evaluate to `true` if the values of both variables or operands are equal. However, the expression `a != b` will evaluate to `true` if the values of both variables or operands are not equal.

Similarly, we have the `<` and `>` operators which allow us to evaluate expression by comparing if the values are less than or grater than the other operand. So, the expression `a > b` will evaluate to `true` if the value of `a` is greater than the value of `b`. Also the expression `a < b` will evaluate to `true` if the value of `a` is less than the value of `b`. 

Finally, the operators `<=` and `>=` allow us to evaluate expression by comparing if the values are less than or equal to and greater than or equal to the other operand. So, the expression `a >= b` will evaluate to `true` if the value of `a` is greater than or if it is equal to the value of `b`, else it would evaluate to `false`. Similarly, the expression `a <= b` will evaluate to `true` if the value of `a` is less than or if it is equal to the value of `b`, else it would evaluate to `false`.

These was a basic overview of comparison operators in golang.

### Logical Operators

Next, we move on to the logical operators in Golang which allow to perform logical operations like `AND`, `OR`, and `NOT` with conditional statements or storing boolean expressions. 

```go
package main

import "fmt"

func main() {
	a := 45
	b := "Something"
	fmt.Println(a > 40 && b == "Something")
	fmt.Println(a < 40 && b == "Something")
	fmt.Println(a < 40 || b == "Something")
	fmt.Println(a < 40 || b != "Something")
	fmt.Println(!(a < 40 || b != "Something"))
}
```

```
$ go run logical/main.go

true
false
true
false
true
```

Here, we have used logical operators like `&&` for Logical AND, `||` for logical OR, and `!` for complementing the evaluated result. The `&&` operation only evaluates to `true` if both the expressions are `true` and `||` OR operator evaluates to `true` if either or both the expressions are `true`. The `!` operator is used to complement the evaluated expression from the preceding parenthesis.

### Arithmetic Operators

Arithmetic operators are used for performing Arithmetic operations. We have few basic arithmetic operators like `+`, `-`, `*`, `/`, and `%` for adding, subtracting, multiplication, division, and modulus operation in golang. 

```go
package main

import "fmt"

func main() {
	a := 30
	b := 50
	fmt.Println("A + B = ", a+b)
	fmt.Println("A - B = ", a-b)
	fmt.Println("A * B = ", a*b)
	fmt.Println("A / B = ", a/b)
	fmt.Println("A % B = ", a%b)
}
```

```
$ go run arithmetic/main.go
A + B =  80
A - B =  -20
A * B =  1500
A / B =  0
A % B =  30
```

These are the basic mathematical operators in any programming language. We can use `+` to add two values, `-` to subtract two values, `*` to multiply to values, `/` for division of two values and finally `%` to get the remainder of a division of two values i.e. if we divide 30 by 50, the remainder is 30 and the quotient is 0. 

We also have a few other operators like `++` and `--` that help in incrementing and decrementing values by a unit value. Let's say we have a variable `k` which is set to `4` and we want to increment it by one, so we can definitely use `k = k + 1` but it looks kind of too long, we have a short notation for the same `k++` to do the same.

```go
package main

import "fmt"

func main() {
	k := 3
	j := 20
	fmt.Println("k = ", k)
	fmt.Println("j = ", j)
	k++
	j--
	fmt.Println("k = ", k)
	fmt.Println("j = ", j)
}
```

```
$ go run arithmetic/main.go

k =  3
j =  20

k =  4
j =  19
```

So, we can see that the variable `k` is incremented by one and variable `j` is decremented by `1` using the `++` and `--` operator.

### Assignment Operators

These types of operators are quite handy and can condense down large operations into simple expressions. These types of operators allow us to perform operation on the same operand. Let's say we have the variable `k` set to `20` initially, we want to add `30` to the variable `k`, we can do that by using `k = k + 30` but a more sophisticated way would be to use `k += 30` which adds `30` or any value provided the same variable assigned and operated on.

```go
package main

import "fmt"

func main() {
	var a int = 100
	b := 20
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	a += 30
	fmt.Println("a = ", a)
	b -= 5
	fmt.Println("b = ", b)
	a *= b
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	a /= b
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	a %= b
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
```

```
$ go run assignment/main.go

a =  100
b =  20

a =  130
b =  15

a =  1950
b =  15

a =  130
b =  15

a =  10
b =  15
```

From the above example, we are able to perform operations by using shorthand notations like `+=` to add the value to the same operand. These also saves a bit of time and memory not much but considerable enough. This allow us to directly access and modify the contents of the provided operand in the register rather than assigning different registers and performing the operations.

That's it from this part. Reference for all the code examples and commands can be found in the [100 days of Golang](https://github.com/mr-destructive/100-days-of-golang/) GitHub repository.

## Conclusion

So, from the following part of the series, we were able to learn the basics of operators in golang. Using some simple and easy to understand examples, we were able to explore different types of operators like arithmetic, logical, assignment and bitwise operators in golang. These are quite fundamental in programming in general, this lays a good foundation for working with larger and complex projects that deal with any kind of logic in it, without a doubt almost all of the applications do have a bit of logic attached to it. So, we need to know the basics of operators in golang.


