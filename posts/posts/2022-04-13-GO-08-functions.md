{
  "type": "posts",
  "title": "Golang: Functions",
  "description": "Understanding the basics of functions in Golang",
  "date": "2022-04-13 09:00:00",
  "status": "published",
  "slug": "golang-functions",
  "tags": [
    "go"
  ],
  "image_url": "https://meetgor-cdn.pages.dev/golang-008-functions.png",
  "series": [
    "100-days-of-golang"
  ]
}

## Introduction

In the eighth part of the series, we will be exploring functions in golang. We will be diving into some basics of functions in golang like declaration, definition and calling. We won't be exploring all the topics of functions as it is quite a large topic to cover in one shot. So, building from the base, we will be starting from the basic declaration to simple return statements. 

## Functions in Golang

Functions in golang are a simple way to structure a block of code that can be re-usable. Functions also allow us to process a piece of logic and return the output. Functions allow us to write readable and scalable code as we have to write the code once and we can re-use the functionality of it by calling it. 

## Declaring Functions 

We have already defined a function, if you have followed the series so far, or even written a `hello-world` application. The `main` function is the most fundamental function we can define in golang. The main is complicated if dive deeper but in the simplest of term it acts as a entry point for the entire program. 

```go
package main

func main() {

}
```

We have written the above code a lot of times till now, but never really talked about it's significance. Here we will understand the terminologies related to the main function. A function is declared with the `func` keyword along with the name of the function. There needs to be the `()` parenthesis after the name of the function, optionally it can take parameters inside the parameters to be used inside the function. 

We define the core functionality or the core logic of the function inside the braces `{}`. We also have the `return` keyword which can return values from the function to the block where we have called the function. Usually, we call a function from other function (most of the times it's the `main` function). The `return` keyword is not mandatory and it is usually added at the end of the function block, just before the closing braces `}`. 

```go
func hello_world_007() {

}
```

We can define a custom function outside the main function by giving it a appropriate name. For the time bwing we can leave it empty and further define the logic of the actual function. 

The name of the function can be given as per the following standards:

- Using letters`a-z A-Z`, numbers`0-9`, underscore `_` as a name.
- Should not contain any spaces in-between the name.
- Should not begin with a number or underscore.

## Defining Functions

Inside the `{}` we define the actual functionality/logic of the function. The variables inside the function will remain local to the function and can't be accessed or altered from outside the function, though if we really want to access some global variables(from main or other functions) we can pass parameters, we will look into it in the next few sections. For time being, we will be focusing on the actual code block inside the function. 

```go

func hello_world() {
    fmt.Println("Hello World")
}
```

This is a basic function that just calls another function `Println` from the fmt package, which basically prints text in the console. Though, we are using the function Println, it won't print the content to the string as we are not using/calling the function. Now, we can get a step ahead and start working with variables inside the function.

```go
func hello_world() {
    name := "Gopher"
    fmt.Println("Hello", name)
}
```

We have now added the local variable `name `inside the function, so this variable can only we referred inside the particular function. 

## Calling Functions

We can call the function from the main function or any other function by just specifying the name along with the `()` and optionally the parameters inside the parenthesis. 

```go
package main

import "fmt"

func main() {
    hello_world()    
}

func hello_world() {
    name := "Gopher"
    fmt.Println("Hello", name)
}

```

```
$ go run func.go
Hello Gopher
```

So, we define the function `hello_world` and call the function by using the statement `hello_world()` inside the main function and now, we are able to actually run the function. 

## Passing Parameters

We can optionally parse variables from a function to other and process it for further computation and programming. So, we can pass parameters in a function by specifying the name to be used inside the function followed by the type of that variable. 

```go
package main

import "fmt"

func main() {

	greet_me("Meet")
	n := "John"
	greet_me(n)
}

func greet_me(name string) {
	fmt.Println("Hello,", name, "!")
}
```

```
$ go run func.go
Hello, Meet !
Hello, John !
```

We have used the parameter `name` as a string in the function and used it inside the function body. The parameter name which is to be called from the main function can be anything and not necessarily be the same as declared in the function declaration. For instance, we have used the variable in the main function `n` which is passed in the function call. We can even pass the value as it is to the function in golang.  

## Return Keyword

We can use the return keyword to actually return a value from the function and not just display the message. The returned value can be later used from other places in the program. 

```go
package main

import "fmt"

func main() {

	// return value
	y := line_eq(3, 1, 2)
	fmt.println("for x = 3 , y = ", y)
}

func line_eq(x int, m int, c int) int {
	return ((m * x) + c)
}
```

```
$ go run func.go
for x = 3 , y =  5
```

So, here we are able to fetch the returned value from the function and store it in another variable and further compute the required logic. We also need to specify the return type of the function after the parameters like `func (parameters) return-type { }`. Here, we need to return the specified type of the return value from the function else it would give a compilation error. 

So, we basically need to provide the return value and also the return statement to capture the value from the function call. 

### Multiple return values

We can also provide multiple return values by providing a list of return values like `(type1 type2 type3 ....)`. We can return the values by separating the values by a comma. So, while calling the function, we need to specify the variables again as comma-separated name and this will capture the value from the function call.

```go
package main

import "fmt"

func main() {

	// multiple return values
	s, c, odd := sqube(5)
	fmt.Println("for x = 5 , x^2 =", s, "x^3 =", c)
	if odd == true {
		fmt.Println("x is odd")
	} else {
		fmt.Println("x is true")
	}
}

func sqube(x int) (int, int, bool) {
	square := x * x
	cube := x * x * x
	var is_odd bool
	if x%2 == 0 {
		is_odd = false
	} else {
		is_odd = true
	}
	return square, cube, is_odd
}
```

```
$ go run func.go
for x = 5 , x^2 = 25 x^3 = 125
x is odd
```

So, we have returned multiple values from the function like two integers and one boolean. The parameter is a single integer, now we need to parse 3 variables in order to capture all the values from the function call. Thus, we are able to get all the values from the function.

That's it from this part. Reference for all the code examples and commands can be found in the [100 days of Golang](https://github.com/mr-destructive/100-days-of-golang/) GitHub repository.

## Conclusion

So, from this part of the series, we are able to understand the basics of functions in golang. We covered from declaration, definition and simple return statements and function calling. Thank you for reading. If you have any questions or feedback, please let me know in the comments or on social handles. Happy Coding :)
