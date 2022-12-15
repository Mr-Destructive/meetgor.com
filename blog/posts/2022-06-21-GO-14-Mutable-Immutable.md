---
templateKey: blog-post
title: "Golang: Mutable and Immutable Data Types"
description: "Understanding and Exploring immutable and mutable data types in Golang. Learning why the variable behaves in golang as they do."
date: 2022-06-21 21:00:00
status: published
slug: golang-mutable-immutable
tags: ['go',]
series: ['100-days-of-golang',]
image_url: https://meetgor-cdn.pages.dev/golang-014-im-mutables.png
---

## Introduction 

In this 14th Post of the 100 days of GOlang, we will be understanding about the mutable and immutable data types in Golang. Firstly, we will understand the concept of mutability and understand the differences in mutable and immutable data types, further we will explore which data types in Golang are Mutable and Immutable.

## Mutable Data Type

Mutable data type is a data type which can be modified without reallocating any chunk of the memory assigned at the time of initialization. In simple words, a variable is mutable if its value can be altered without reallocating itself to a new memory space.

In mutable data type the value located in a memory address can be modified. This means we do not have to reallocate any memory or change the pointer of a variable to point to other address in order to change the value of the variable.

```
// Initialization 

int age = 12

// The value of age is 12 which is stored at some memory address let's sayy 0x12345


// Modifying the value

int age = 13

// If the data type is mutable, we can directly change the value in the memory address
```

Mutable data type looks quite simple by using a single variable to demonstrate, but if we scale the example a little, things start to escalate pretty quickly.

```
// Initialization 

int age = 12
// The value of age is 12 which is stored at some memory address let's sayy 0x12345

int experience = age

// This will make the experience variable point to the address which the age variable is pointing to.
// Yes, It will point to the same memory address

// Variable    |  Memory Address  |  Value

// age        ->    0x12345       ->   12
// experience ->    0x12345       ->   12

// Modifying the value

int experience = 17

// A change in either of the variable will change the value in both the variables 
// Since both the variables are poiting to the same memory address

// Variable    |  Memory Address  |  Value

// age        ->    0x12345       ->   17
// experience ->    0x12345       ->   17

// If the data type is mutable, we can directly change the value in the memory address

```   

## Immutable Data Types

Immutable data type is a data type which cannot be modified without allocating a new memory. So, the immutable data type has to reallocate memory for making changes to the value of a variable. This might be a downside if the variable is holding a large sets of values, it will require a lot of memory re-allocation for a slight change in the value. 

Immutable data types also mean that you cannot change the value in the memory address which the variable is pointing to, but you can make the variable point to a different memory location under the hood to change or modify the content of a variable.

```

// Initialization

string fruit = "apple"

// the value "apple" is stored in a memory location let's say 0x12345 originally



// Modifying the value

string fruit = "orange"

// the value "orange" will be located in a different memory location say 0x98765
// The memory address that the variable fruit points to is changed and not the value of the memory address
// This is called immutablility in data types
```

This is the basics of mutability and immutability, this might be a bit difficult to digest, but take your time and understand it thoroughly. This concept is critical for understanding the under the hood point of view while debugging in several occasions.


## Mutable data type in Golang

In golang there are a few mutable data types 

1. Slice
2. Array
3. Map
4. Channels

### Slice and Arrays

Slice and Arrays are mutable data types in golang, this means the value of the elements in slice or array can be changed after initialization without re-allocations of memory.

```	go
package main

import "fmt"

func main() {

	// Slice
	s := []int{1, 2, 3}
	fmt.Printf("S[1] -> %p\n", &s[1])
	p := s
	p[1] = 4
	fmt.Printf("S[1] -> %p\n", &s[1])

	fmt.Println("s =",s)
	fmt.Println("p =",p)
}
```

```
go run mutable.go

S[1] -> 0xc000018200
S[1] -> 0xc000018200
s = [1 4 3]
p = [1 4 3]
```

In the above example, we can see that the slice has been initialized to `{1,2,3}` another slice has been referenced to that string i.e. it will point to the same address as the original string `s`. If we modify the slice `p`, since it is pointing to the same memory address as the slice `s` is pointing it will modify the slice `s` as well (they are the same slice).

This is what mutability does, it modifies the value stored in a memory address directly, without allocating any extra memory to the variable.

### Arrays

You won't be able to see the concept of mutable data types with arrays as they are not referenced to any memory address, it is a collection of a single type of value and it is thus static. Since it is not a reference to any memory address, the value of the elements doesn't change if we change the value of an element in the copy of the array.

```go
package main

import "fmt"

func main() {

	// Arrays 
	a := [3]int{10, 20, 30}
	fmt.Printf("A[1] -> %p\n", &a[1])
	b := a
	b[1] = 40
	fmt.Printf("A[1] -> %p\n", &a[1])

	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
```

```
go run mutable.go

A[1] -> 0xc0000aa038
A[1] -> 0xc0000aa038
a = [10 20 30]
b = [10 40 30]
```

   In the above example, we can see the array `a` is initialized with a fixed length and initialized values. We then create another array by assigning the array `a` to it, this creates a copy of a collection of the elements in a different memory location. So, if we change an element in the array `b`, there won't be any change in the elements of array `a` as the elements are stored in a completely different memory location.

### Map

Map is similar to slices in a way they are references to the memory address. A map as we have explored in the [seventh part](https://www.meetgor.com/golang-maps) of the series], they are a pair of key and value pairs. The map is internally a reference to a hash map, a hash map is an abstract data type or a structure in Golang, it basically is an array of buckets. Buckets contain high-order bits with a hash(random value) to make the keys distinct in the map. The number of buckets is initially 8, but it expands as required so it doubles the number of buckets and assigns the value to the map elements. For a detailed reference, you can look at the source implementation of [golang's map](https://github.com/golang/go/blob/master/src/runtime/map.go).

So, a map is mutable, which means if we change the value of a key it is changed directly to the memory rather than reallocating the memory space for the entire map.

```go
package main

import "fmt"

func main() {
	// Map
	m := map[string]int{"level": 5, "health": 9}
	fmt.Println("m =", m)
	n := m
	n["food"] = 12

	fmt.Println("m =", m)
	fmt.Println("n =", n)
}
```

```
go run mutable.go

m = map[health:9 level:5]
m = map[food:12 health:9 level:5]
n = map[food:12 health:9 level:5]
```

In the above example, we can see the map `m` is a string integer map with 2 keys. We create a new map called `n` and reference it as the map `m` this will make the map `n` point to the same hash map as the map `m` does. Hence, if we change the value of a key in the map `n` it will thereby change the map `m` as maps `m` and `n` both point to the same map (in memory address).

## Immutable data type in Golang

In golang, there are a few immutable data types as well like string, pointers, boolean, and core data types like integer, float, etc. As we discussed immutable data types, are data types that don't change the value of the variable directly into the provided memory address, it re-allocates the memory address with the new value/edited value. 

1. Boolean, Int, Float
3. Pointers
3. String
4. Interfaces

### Boolean

The boolean data type on golang is an immutable data type which means it re-allocates the memory for any change in the value of the boolean variable. Boolean variables are simple as they can have two values either `true` or `false`. If we declare a boolean variable initialize it with a value, if we further wanted to change the value of the variable, it is done by reallocating the memory address which was initially holding the value. 

```go
package main

import "fmt"

func main() {

	// bool
	boolean := true
	b := boolean
	b = false
	fmt.Println("boolean = ", boolean)
	fmt.Println("b = ", b)

}
```   

```
go run immutable.go

boolean =  true
b =  false
```

   In this example, we can see that the boolean variable `boolean` (I have literally named it boolean) is initialized as `true`, next we declare another variable `b` and set it to the value of `boolean` which is the value `true`. We then change the value of the variable `b` but the value of `boolean` does not change. It doesn't give much clarity on the immutability of the boolean data type. This is because it is an under an hood process and a low-level process in programming in golang.

Immutable doesn't mean you cannot change values, but the value is not directly changed, there happens some under-the-hood (low level) stuff to alter the value and change the memory address which is ben pointed as the location which holds the value. 

To get more clarity, you can refer to this [Stack-Overflow discussion](https://stackoverflow.com/questions/71589811/go-ints-and-strings-are-immutable-or-mutable/71590289#71590289).


### Pointers

Pointer as well is an immutable data type in golang, we cannot change the value in this case the memory address of a variable which we are pointing to, directly but we need to re-allocate the memory for changing the value of the variable.


```go
package main

import "fmt"

func main() {

  // Pointer
	n := 567
	t := 123
	ptr := &n
	ptr_new := ptr
	fmt.Println("ptr = ", ptr)
	fmt.Println("ptr_new = ", ptr_new)

	ptr_new = &t

	fmt.Println("ptr = ", ptr)
	fmt.Println("ptr_new = ", ptr_new)
}
```

```
go run immutable.go

ptr     =  0xc0000aa008
ptr_new =  0xc0000aa008
ptr     =  0xc0000aa008
ptr_new =  0xc0000aa020
```

   Clearly, in the example, we have initially created a pointer `ptr` which points to the variable `n`, then we create another pointer `ptr_new` and assign it to the value of the pointer `ptr`. This will make both the pointers point to the same variable `n`. We then point the `ptr_new` pointer to the variable `t`, this changes the value of pointer `ptr_new` but since the pointer stores a memory address, it was not holding the actual value at the memory address. So, the change in the pointing value of `ptr_new` does not change the value of the pointer `ptr`.

Again, immutable types cannot be seen actually changing the memory location of those variables, it is just the low-level implementation that sometimes needs to be kept in mind.

### String

Strings are the classical example of immutable data types in golang, this data type is quite commonly used and is quite important for creating a wide variety of applications. The value of the string variable can be changed but the process happens with/without changing the value of the memory address of the initial value, we have to change the memory address of the variable in order to change the value. This might not be evident by even using `&variable_name` as it doesn't give the actual insight about the internal working of how the memory address might be processed at run time.


```go
package main

import "fmt"

func main() {

	// String
	str := "gopher"
	str_copy := str
	str_copy = "cooper"
	fmt.Println("str = ", str)
	fmt.Println("str_copy = ", str_copy)
}
```

```
go run immutable.go

str =  gopher
str_copy =  cooper
```

Here, in the above example, we have created a variable `str` with a value `"gopher"` and then, another variable `str_copy` which is assigned the value of `str`. If we change the value of `str_copy`, the value of `str` is not changed, this is not giving any information on the immutability of a data type.

What actually gives a clear understanding of the immutability of data type in the string is the modification of the character of a string. We cannot change the character at a particular index of the string in golang.

```go
package main

import "fmt"

func main() {

	// Character at index cannot be changed in string

	s := "StarWars"
	s[4] = 'C'
	// s[4] = "C"
	// also won't work
	fmt.Println(s)

}
```

```
go run immutable.go

immutable.go:18:2: cannot assign to s[4] (value of type byte)
```
  
   So, we can see we cannot change the value of the internal characters of the string. This is why string data type is immutable, we cannot change the value of internal elements of the variable directly but we can change the value of the variable as a whole as we saw in the previous example. 

So this is how the immutable data types are perceived in golang, we cannot change the value directly at the memory address, we change the internal location for any change in the value which is dependent on the decision by the garbage collector.

That's it from this part. Reference for all the code examples and commands can be found in the [100 days of Golang](https://github.com/Mr-Destructive/100-days-of-golang/tree/main/scripts/im-mutable) GitHub repository.

## Conclusion

So, from this small post, we were able to understand the different behaviors of data types in golang. We can understand it at the beginning of the variables part of the series, but it becomes more clear when we play with all the data types and then understand how and why they behave as they do. This was not a deep dive into the core working of immutable and mutable data types but gave a fair bit of understanding of the behavior of different data types in golang.

Thank you for reading, if you have any queries, feedback, or questions, you can freely ask me on my social handles. Happy Coding :)
