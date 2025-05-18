{
  "type": "posts",
  "title": "Golang: Anonymous Functions",
  "description": "Understanding the concept of the Anonymous functions in golang",
  "date": "2022-07-24 18:15:00",
  "status": "published",
  "slug": "golang-anonymous-functions",
  "tags": [
    "go"
  ],
  "series": [
    "100-days-of-golang"
  ],
  "image_url": "https://meetgor-cdn.pages.dev/golang-017-anonymous-functions.png"
}

## Introduction

We have looked at the defer keyword in golang in the [previous](https://www.meetgor.com/golang-defer/) part of the series, in this section, we will understand how we can use anonymous functions in golang. We will explore how to declare and use anonymous functions with a few examples.

## What are Anonymous Functions

Anonymous functions are quite simple to understand, we don't define a function, we declare it and call it instantly. An anonymous function doesn't have a name so hence it is called an anonymous function. As a normal function it can take in parameters and return values. With anonymous functions, we can bind the operations to a variable or a constant as a literal(value). If an anonymous function takes in a parameter, it needs to be parsed immediately after the end of the function body. We will see how we define the syntax and specifications of the anonymous functions in golang. 

## Simple Anonymous functions

To create a simple anonymous function we use the same function syntax without giving it a name like `func() {}`, inside the function body i.e. `{}`, you can define the operations that need to be performed. 

Here, we have created an anonymous function that simply calls the `fmt.Println` function with a string. So, we have made things a little too much as we can even directly call the `fmt.Println` function from the main function, instead we have called an anonymous function that in turn calls the `fmt.Println` function. It might not make sense to use an anonymous function here, but it can be helpful for other complex tasks that you want to isolate the logic without creating a dedicated function for the process.

```go
package main

import "fmt"

func main() {

  func() {
    fmt.Println("Hello, Anonymous functions")
  }()

}
```

```
go run anonymous_function.go

Hello, Anonymous functions
```

So, this is how we create a basic anonymous function in golang, this can be further used for returning values from the function or passing parameters into the function and also assigning the function call to a variable or a constant.

## Assigning anonymous function to variables

We can assign the call to the anonymous function to a variable or a constant and call the function as many times as we require. So, we can basically store the function logic in a variable and call it whenever we require the function with the `()` parenthesis as an indication to call the function.

In the following example, we have used the variable `draw` to store the function call which prints `Drawing` with the help of the `fmt.Println` function. The draw variable now contains the function and not its value. So be careful here, the anonymous function which we have defined as the `draw` variable's literal value, it's like we are giving a name to this anonymous function and the name will be the variable name so we have created the function `draw` which is an anonymous function stored in a variable.


```go
package main

import "fmt"

func main() {

  draw := func() {
    fmt.Println("Drawing")
  }
  draw()
  draw()
}
```

```
go run anonymous_function.go

Drawing
Drawing
```

We see that we call the variable `draw` as a function call by appending `()` parenthesis to it as `draw()` this thereby calls the anonymous function stored inside the variable value.

The main thing to note here is that we are not adding `()` at the time of declaring the anonymous function, as it will call the function directly. The problem then will arise because the function is not returning anything so we can't assign it to a variable.

```go
draw := func() {
  fmt.Println("Drawing")
}()
```

```
functions/anonymous_functions.go:10:11: func() {…}() (no value) used as value
```

So, we are trying to assign a variable to a function call that has no return value. This has no value, not even nil, so we get an error for assigning a variable to nothing.

Though if you had a return value from the function, we can directly assign the function call to the variable as it has returned a valid value defined in the function literal.

## Parsing parameters

We can even parse parameters to the anonymous functions as any other normal function. We define the name of the variable followed by the type of the variable in the parenthesis to use these parameters inside the anonymous function.

We need to keep in mind that these function parameters can either be passed with the variable call or directly at the time of defining the function.

In the below example, we have created a variable `call` and are assigning it to an anonymous function that takes in a parameter `name` which is a `string`, and prints some text to the console. So, we call the funciton `call` with a parameter as a string, as we have `"Meet"` and `person := "Chris"` as a string passed to the `call` function.

```go
package main

import "fmt"

func main() {

  call := func(name string) {
    fmt.Println("Calling,", name)
  }

  call("Meet")
  person := "Chris"
  call(person)

}
```

```
go run anonymous_function.go

Calling, Meet
Calling, Chris
```

Here, we can see that the function `call` prints the text that we have defined to call the `fmt.Println` function. We parse the function with the string literal as the anonymous function takes a parameter in the form of a string. We can parse multiple parameters to the anonymous function as we can with the normal function like slices, maps, arrays, struct, etc.

## Returning values

We can even return values from the anonymous function if we want to instantly call the function and save the `returned` value in the variable. We can return single or multiple values as per our requirements just like any normal function in golang.


```go
package main

import "fmt"

func main() {

  is_divisible := func(x int, y int) bool{
    var res bool 
    if x % y == 0 {
      res = true
    } else{
      res = false
    }
    fmt.Println(res)
    return res
  }

  fmt.Printf("%T
", is_divisible)
  is_divisible(10, 5)
  is_divisible(33, 5)

  divisblity_check := is_divisible(45, 5)
  fmt.Printf("%T : %t
", divisblity_check, divisblity_check) 

}
```

```
go run anonymous_function.go

func(int, int) bool 
true
false
bool : true
```

As we can see the function has returned a boolean value and we store it in a variable `divisblity_check`, the variable which contains the function i.e. `is_divisible` can then be called, and thereby we get the returned value in the variable as a boolean as we print the type and the value of the variable `divisblity_check`. We also can see that the type of the variable `is_divisible` is of type function literal which takes `int, int` and has a return value as `bool`.

We can also do one more interesting here, which we were restricted previously in the above examples. We can directly call the function and store it as a value rather than the function itself. So, we can only use the value returned from the function but can't call the function later.

```go
is_divisible := func(x int, y int) bool{
  var res bool 
  if x % y == 0 {
    res = true
  } else{
    res = false
  }
  fmt.Println(res)
  return res
}(13, 4)

fmt.Printf("%T
", is_divisible)
fmt.Printf(is_divisible) 
```

```
go run anonymous_function.go

bool
false
```

So, in the above-modified example, we have passed in the parameter instead of a callable function. This will store the returned value of the function call. So, we will store the boolean value in the `is_divisible` and we will have to pass the integer values to the function which we have parsed as `(13, 4)` to the anonymous function call.


In the below example, we have created an anonymous function that takes in three parameters like `(string, int, string)` and returns a string. We have used `fmt.Sprintf` function to format the variable and store it in a variable, we then return the string. This anonymous function is then directly called and we store the returned value instead of the function.

So, in this example, the `format_email` variable will store a string instead of acting it as a function as a callable object.

```go
package main

import "fmt"

func main() {

  format_email := func(name string, age int, company string) string{
    email := fmt.Sprintf("%s.%d@%s.com", name, age, company)
    return email
  }("john", 25, "gophersoft")

  fmt.Println(format_email)
  fmt.Printf("%T
",format_email)

}
```

```
go run anonymous_function.go

john.25@gophersoft.com   
string
```

As we can see the `format_email` variable only returns a string instead of a callable object. We have directly parsed the parameters to the anonymous function and hence it instantly calls it and we return the string.

## Passing Anonymous function to function parameters

We can even pass an anonymous function to a parameter to a function. This can be quite helpful for writing some simple logic inside a complex script. We can pass the function type as a parameter to the parameter and theere we can parse the actual data and perform the desired operation.

The below example is a bit of code to write but makes a lot of sense to understand anonymous functions in golang. The function `get_csv` is a function which takes in three parameters `int, string, func(string)[] string`. The third parameter is a function literal that takes in a string and returns a slice of string. So, we are basically converting a string `"kevin,21,john,33"` into a slice of the slice like `[[kevin 21] [john 33]]`, this is basically separating values with `,` comma as the delimiter and then processing slices to create a single slice.

```go
package main

import "fmt"
import "strings"

func get_csv(index int, str string, t func(csv string)[]string) [][]string{
  s := t(str)
  var res [][]string
  for i := 1; i<len(s); i+=2 {
    var data []string
    data = append(data, s[i-1], s[i])
    res = append(res, data)
  }
  return res
}

func main() {

  csv_slice := func (csv string) []string{
    return strings.Split(csv, ",")
  }
  csv_data := get_csv(2,"kevin,21,john,33,george,24", csv_slice)
  fmt.Println(csv_data)
  for _, s := range csv_data{
    fmt.Printf("%s - %s
",s[0],s[1])
  }
```

```
go run functions/anonymous_function.go

[[kevin 21] [john 33] [george 24]]

kevin - 21
john - 33
george - 24
```

Let's break down the code one by one, we will start with the main function, where we have `csv_slice` as a function literal and is an anonymous function that takes in a string and returns a slice of string. The function returns a call to the function [strings.Split](https://pkg.go.dev/strings#Split) taking in the string from the function parameter. We then call the function `get_csv` with parameters `(2, "kevin,21....", csv_slice)`, this function is defined outside the main. The function takes in three parameters as discussed and parsed from the main function and it returns a slice of type string. 

So, inside the `get_csv` function, we define `s` as the function call to `t(str)` which if you look carefully is a function call to `csv_slice` with parameter as a string. This function call returns a slice of strings separated by `,`. So that's all the logic we need to understand anonymous function from parameters. We have used a function literal to call the function from another function. In this case, the funciton is an anonymous function assigned to a variable. 

```
"kevin,21,john,33,george,24"
            ||
            \/
[kevin 21 john 33 george 24]
            ||
            \/
[[kevin 21] [john 33] [george 24]]

```

Further, after we have `s` which would look like `[kevin 21 john 33 george 24]` as each individual element. Thereafter we create an empty [slice of slice](https://www.geeksforgeeks.org/slice-of-slices-in-golang/) string as `res` and operate a loop through the slice jumping 2 indexes at a time. Why? because we want a slice of two elements combined. So, we also create a slice of string named `data` and we add the two components to it like `[kevin 21]` with the [append](https://pkg.go.dev/builtin#append) function, and this slice is appended to a slice of slices `res` as `[[kevin 21]]` thereby iterating over the entire slice and creating the desired slice data. We return the `res` from the function. This get's us back to the main function which simply iterates over the slice and we get the desired data from it.

So, this is how we convert an extremely easy task to extremely complicated code :)

That's it from this part. Reference for all the code examples and commands can be found in the [100 days of Golang](https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/functions/anonymous_function.go) GitHub repository.

## Conclusion

That is it from this part, we covered a little bit about anonymous functions in golang. Anonymous functions are simply function literals that can be used to do a lot of quick operations without needing to write an explicit function in the program. Further, in the next part look into closures which are a bit complex to understand and require the understanding of anonymous functions.

Thank you for reading, if you have any queries, feedback, or questions, you can freely ask me on my social handles. Happy Coding :)
