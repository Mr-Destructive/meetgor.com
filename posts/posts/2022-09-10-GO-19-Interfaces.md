{
  "type": "posts",
  "title": "Golang: Interfaces",
  "description": "Understanding the baiscs of interfaces in Golang. Exploring the concept of polymorphism in golang with the help of interfaces and structs.",
  "date": "2022-09-10 17:30:00",
  "status": "published",
  "slug": "golang-interfaces",
  "tags": [
    "go"
  ],
  "series": [
    "100-days-of-golang"
  ],
  "image_url": "https://meetgor-cdn.pages.dev/golang-019-interfaces.png"
}

## Introduction

In the 19th post of the series, we will be taking a look into interfaces in golang. Interfaces allow us to create function signatures common to different structs or types. So, we can allow multiple structs to have a common interface(method) that can have different implementations.

## What are Interfaces

Interface as the name suggests is a way to create methods that are common to different structures or types but can have different implementations. It's an interface to define the method or function signatures but not the implementation. Let's take an example of `Laptop` and `Phone` having the functionality of wifi. We can connect to wifi more or the less in a similar way on both devices. The implementation behind the functionality might be different but they share the same operation. The WiFi can act as an interface for both devices to connect to the internet.

## Define an Interface

To declare an interface in golang, we can use the `interface` keyword in golang. An interface is created with the type keyword, providing the name of the interface and defining the function declaration. Inside the interface which acts as a struct of general method signatures. The method signatures usually consist of the name of the function with its parameters if any and the return type of the function.

```go

package main

import "fmt"


type Player struct {
	name   string
	health int
}

type Mob struct {
	name     string
	health   int
	is_passive bool
}

type Creature interface {
	intro() string
	//attack() int
	//heal() int
}

func main() {
	player := Player{name: "Steve"}
	mob := Mob{name: "Zombie"}
	fmt.Println(player)
	fmt.Println(mob)
}
```

```
go run main.go

{Steve 0}
{Zombie 0 false}
```

In this above example, we have created an interface called `Creature`. There are a few structs that we have to define like `Player` and `Mob`, these two methods have a few attributes like `name` as `string` and `health` as `int` which are common in both structs but the `Mob` struct has an additional attribute `is_passive` as a `boolean` value. The `Creature` is an interface that will define certain function signatures, here we have declared `intro`, `attack`, and `heal` as the methods bound to the Creature interface. This means, that any object which satisfies the Creature interface, can define the methods associated with the interface.

## Defining Interfaces

Once we have declared the interface method signatures, we can move into defining the functionality of these methods depending on the struct. If we want a different working method for different types of struct objects passed we can define those for each type of struct that we want. Here, we have two types of structs namely `Creature` and `Mob`, based on the struct we can define the `intro` method for them individually.

```go
package main

import "fmt"

type Creature interface {
	intro() string
	attack(*int) int
}

type Player struct {
	name   string
	health int
}

type Mob struct {
	name     string
	health   int
	category bool
}

func (p Player) intro() string {
	fmt.Println("Player has spawned")
	return p.name
}

func (p Player) attack(m_health *int) int {
	fmt.Println("Player has attacked!")
	*m_health = *m_health - 50
	return *m_health
}

func (m Mob) intro() string {
	fmt.Printf("A wild %s has appeared!
", m.name)
	return m.name
}
func (m Mob) attack(p_health *int) int {
	fmt.Printf("%s has attacked you! -%d
", m.name, 30)
	*p_health = *p_health - 30
	return *p_health
}

func main() {
	player := Player{name: "Steve", health: 100}
	mob := Mob{name: "Zombie", health: 140}
	fmt.Println(player.intro())
	fmt.Println(mob.intro())
	fmt.Println(mob)
	fmt.Println(player)
	fmt.Println(player.attack(&mob.health))
	fmt.Println(mob.attack(&player.health))
	fmt.Println(mob)
	fmt.Println(player)
}
```

```
go run main.go

Player has spawned
Steve
A wild Zombie has appeared!
Zombie
{Zombie 140 false}
{Steve 100}
Player has attacked!
90
Zombie has attacked you! -30
70
{Zombie 90 false}
{Steve 70}
```

As we can see, the method `intro()` is bound to both the struct depending on what struct signature is associated with the method. The method `intro` takes in the object struct associated as per the call and returns `string` as defined in the method signature. 

The `attack` method in the `Creature` interface is also implemented separately for the two structs. For the `Player` method, we simply take in a pointer to an integer and return an `int`. The parameter is the pointer to the mob health, and it returns the modified health. We take in a pointer to the mob or player health so as to parse in the actual value and not the copy of the value. If we modify the value, we want to reflect those changes in the actual object. So that is how we can use interfaces to construct dynamic operations on objects as well as different types of structs.

We have seen a simple example of how to declare and define interfaces for given type structs. Also, we can pass by value as well as by pointers so as to define the behavior of the method whether to dynamically modify or change the values of the object associated with it.

## Examples of Interfaces

There are quite some use cases of interfaces, in object-oriented programming, the above example fits the polymorphism feature quite well. The ability to reuse certain method signatures and define the functions as per requirement brings flexibility to the code structure.

We will see a few examples for understanding how we can use interfaces in various ways.

### Type Switch Interface

We can use an empty interface to check for the type of variable we have parsed. Using this empty interface we can create a kind of dynamic parameter to a function.  

```go
package main

import (
	"fmt"
	"strconv"
)

func parse_int(n interface{}) int {
	switch n.(type) {
	case int:
		return (n).(int) * (n).(int)
	case string:
		s, _ := strconv.Atoi(n.(string))
		return s
	case float64:
		return int(n.(float64))
	default:
		return n.(int)
	}
}

func main() {
	num := parse_int(4)
	fmt.Println(num)
	num = parse_int("4")
	fmt.Println(num)
	num = parse_int(4.1243)
	fmt.Println(num)
}

```

```
go run main.go

16
4
4
```

Here, we can see we have an interface as a parameter to the function `parse_int`, the return type is `int`, so the incoming parameter can be any valid type. But if we don't convert the given type into an appropriate int, it will result in an error as we are returning the int value of the parsed parameter. We are taking the parameter as `interface{}` which is an empty interface, this will contain the parameter parsed as an interface type. That's why we need to convert the interface object into an int or the parsed type of the interface.

### Interface Slice

We can even create a slice of interfaces, which means we can initialize or group together various objects of different structs in a single slice. This might be helpful for calling functions associated with different objects via the interface very easily and in a much cleaner way.

```go
package main

import "fmt"

type Creature interface {
	intro() string
}

type Player struct {
	name   string
	health int
}

type Mob struct {
	name     string
	health   int
	category bool
}

func (p Player) intro() string {
	fmt.Println("Player has spawned")
	return p.name
}

func (m Mob) intro() string {
	var name string
	if m.name != "" {
		name = m.name
	} else {
		name = "Mob"
	}
	fmt.Printf("A wild %s has appeared!
", name)
	return m.name
}


func main() {
	entity := []Creature{Player{}, Mob{}, Mob{}, Player{}}

	for _, obj := range entity {
		fmt.Println(obj.intro())
	}
}
```

```
Player has spawned
A wild Zombie has appeared!
A wild Zombie has appeared!
Player has spawned
```

In the above example, we can see that the entity variable is created as a slice of interfaces `Creature`, i.e. various objects associated with the `Creature` interface can be contained in a single slice. There are 2 instances of Player and Mob each in the slice. We can further iterate over the slice as a range-based loop and thereby the functions associated with the interfaces can be called. Here, we have called the `intro` function.

So, there are a lot of things that can be done with interfaces, we can create multiple interfaces for a single object struct and nest interfaces. Based on the use case of the program, interfaces can be used to reduce the boilerplate code as well as improve the readability of the code.

That's it from this part. Reference for all the code examples and commands can be found in the [100 days of Golang](https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/interfaces/main.go) GitHub repository.

## Conclusion

From this part of the series, we were able to understand the basics of interfaces using a few examples. We explored how interfaces can be used to bring in polymorphism in golang, also we can improve the readability of the code. The boilerplate code can be considerably reduced by using interfaces when dealing with structs and types. Hopefully, you found this post helpful and understood even the basics of interfaces in golang. Thank you for reading, if you have any queries, questions, or feedback, you can ping me on my social handles or in the comments. Happy Coding :)
