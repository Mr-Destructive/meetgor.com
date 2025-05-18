{
  "type": "posts",
  "title": "Golang: File Write",
  "description": "Using system calls to write files in Golang. Performig writing, appending, deleting, replacing operations to a file",
  "date": "2022-12-18 15:15:00",
  "status": "published",
  "slug": "golang-file-write",
  "tags": [
    "go"
  ],
  "series": [
    "100-days-of-golang"
  ],
  "image_url": "https://meetgor-cdn.pages.dev/golang-024-file-write.png"
}

## Introduction

In the 24th post of the series, we will be taking a look at how we can perform write operations to a file using golang. We will be using the `os` package in most operations along with `bufio` text manipulations. We will be performing write operations like appending, deleting, and replacing a file using golang. We will be heavily leveraging standard library packages like `os`, `bufio`, `bytes` and `fmt`. We will also be looking into overwriting and string formatting to a file.

## Write to a File

The first part of this section is the basic write operation to a file, we assume we are writing to a fresh file and overriding the contents of the existing file. The next section will cover the appending of content to the file and so on. In this example, we will see how we perform basic writing operations to write a string, a slice of string to a file.

```go
package main

import (
	"log"
	"os"
)

func HandleError(err){
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
	str := "Hi, I am a gopher!
"
	f, err := os.OpenFile("abc.txt", os.O_WRONLY, 0660)
    // f, err := os.Create("abc.txt")
	HandleError(err)
	_, err = f.Write([]byte(str))
	HandleError(err)
	defer f.Close()
}
```

```bash
$ cat abc.txt

$ go run main.go

$ cat abc.txt
Hi, I am a gopher!
```

So, we have used a simple golang script to write to a file that exists/has already been created. If you don't want errors while having to write on a file that does not exist, use the [Create](https://pkg.go.dev/os#Create) method instead which is similar to the `Open` method but creates a file if it doesn't exist. We use the [Write](https://pkg.go.dev/os#File.Write) method to overwrite contents to the file, it takes in a parameter as a slice of byte, so we typecast the string `str` into `[]byte` using the `[]byte(str)` syntax. Thereby we write the contents of the string into the file. We use the defer keyword for closing the file at the end of the script or the end of the main function scope.

### Write slice of strings to file

We can even write slice of string to file using a for loop and appending each string with a new line character.

```go
package main

import (
	"log"
	"os"
)

func HandleError(err){
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
	f, err := os.Create("abc.txt")
	//f, err := os.Open("abc.txt", os.O_WRONLY, 0660)
	langs := []string{"golang", "python", "rust", "javascript", "ruby"}
	for _, lang := range langs {
		_, err := f.WriteString(lang + "
")
		HandleError(err)
	}
	defer f.Close()
}
```

```plaintext
$ cat abc.txt

$ go run main.go

$ cat abc.txt
golang
python
rust
javascript
ruby
```

We have used the [WriteString](https://pkg.go.dev/os#File.WriteString) method which will take in a string as a parameter instead of a slice of bytes. So, we don't have to type cast into slice of bytes. So, as we can see we have written the string slice into a file.

### Over Write

The minimal code to write to a file is the [WriteFile](https://pkg.go.dev/os#WriteFile) function in the [os](https://pkg.go.dev/os) package, it overrides the content of the file with the provided slice of bytes, the name of the file, and the necessary permission to write. The funciton additionally creates a file if it does not exist, which is one less reason for the error. Though it returns an error object, the error might be created due to not right permissions to write to the file, encoding issues, etc.

```go
package main

import (
	"log"
	"os"
)

func main() {
	data := []byte{115, 111, 109, 101, 65}
	err := os.WriteFile("test.txt", data, 0660)
    log.Println(err)

	s := "Hello"
	err = os.WriteFile("test.txt", []byte(s), 0660)
    log.Println(err)
}
```

```plaintext
$ go run main.go
2022/12/17 19:24:13 <nil>
2022/12/17 19:24:13 <nil>

$ cat test.txt
Hello
```

So, we have used the `WriteFile` method two times in the script, it first takes in a slice of bytes as it is defined as `data` which corresponds to `115 -> s`, `111 -> o`, `65 -> A`, ASCII mapped to strings. The slice of bytes can be taken as a string like `someA` as the literal value of the underlying slice of the byte. So, we take that slice of byte and parse it to the second parameter of the WriteFile function. The first parameter is a string path of the file we want to write the contents to, the third parameter is the file permission. We have set it as `0660` indicating read(4) + write(2) to the group and the user and no permission to the other users. The function will return an error if any, or else it simply overwrites the data in the file.

In this case, we have called the `WriteFile` method with string `s` type cast to slice of bytes at the end of the script so we see the file has contents as `Hello` instead of `someA`. If we reverse the action, we don't see the `Hello` string in the file.

```go
package main

import (
	"log"
	"os"
)

func main() {
	s := "Hello"
	err := os.WriteFile("test.txt", []byte(s), 0660)
    log.Println(err)

	data := []byte{115, 111, 109, 101, 65}
	err = os.WriteFile("test.txt", data, 0660)
    log.Println(err)
}
```

```plaintext
$ go run main.go
2022/12/17 19:24:13 <nil>
2022/12/17 19:24:13 <nil>

$ cat test.txt
someA
```

As we can see the `Hello` has been overwritten by `someA`.

### Write formatted string

We can even use fmt to write formatted strings to a file. Just like we can take inputs with `Scanf`, we can use [Fprint](https://pkg.go.dev/fmt#Fprint) and other similar functions like [Fprintf](https://pkg.go.dev/fmt#Fprintf), and [Fprintln](https://pkg.go.dev/fmt#Fprintln) functions to print/add contents to the file.

```go
package main

import (
	"fmt"
	"log"
	"os"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	f, err := os.Create("temp.txt")
	HandleErr(err)
	name, lang, exp := "John", "go", 2
	_, err = fmt.Fprint(f, "Hi, I am ", name, "
")
	HandleErr(err)
	_, err = fmt.Fprintf(f, "Language of choice: %s.
", lang)
	HandleErr(err)
	_, err = fmt.Fprintln(f, "Having", exp, "years of experience.")
	HandleErr(err)
	defer f.Close()
}
```

```plaintext
$ cat temp.txt
cat: temp.txt: No such file or directory

$ go run format.go

$ cat test.txt
Hi, I am John.
Language of choice: go.
Having 2 years of experience.
```

So, we can see that we have used all three methods having their own use cases, we can use `Fprint` for simple strings, `Fprintf` for formatting the block of a string with multiple placeholders, and the `Fprintln` which works simply like `Fprint` but it adds a new line itself, we don't need to specify it explicitly.

### Append

If we want to append text to a file, we can use the [OpenFile](https://pkg.go.dev/os#OpenFile) function and provide a few parameters to append the contents instead of overwriting.

Here, we have two steps, open the file and then write the contents in the file. So while opening the file, we provide a few options as parameters to make the fine-tuned system call like only open for read, write or append modes. These options are defined as constant int values in the [os package](https://pkg.go.dev/os#pkg-constants).

```go
package main

import (
	"log"
	"os"
)

func HandleError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
	s := "Hello"
	err := os.WriteFile("test.txt", []byte(s), 0660)
    HandleError(err)

	s = "World"
	f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0660)
	HandleError(err)
	_, err = f.WriteString(s)
	HandleError(err)
	defer f.Close()
}
```

```plaintext
$ go run main.go

$ cat test.txt
HelloWorld
```

So, from the above example, we are able to append text into a file, we have first added the `Hello` world string into the file using the `WriteFile` method to indicate we overwrite the previous contents of the file. We then use the [OpenFile](https://pkg.go.dev/os#OpenFile) method to open a file provided in the first parameter as a string path. The second parameter is the options to be passed for performing operations on the opened file, we always have to use them `defer` to close the file or other resource-locking operations.

We have specified the `os.O_WRONLY` and the `os.O_APPEND` options indicating we want to write to the file while the file is open and specifically append to the file. So this is fine-tuning the opened file operation. We can use the ReadFile or WriteFile operation which is just used for simple read and write operations respectively.

We use the [WriteString](https://pkg.go.dev/os#File.WriteString) method, but we can even use the [Write](https://pkg.go.dev/os#File.Write) method to write a slice of byte instead. This is just used for exploring the different options in the file types of the os package.

### Append at a specific line

We can also add content to a specific line or a portion of the file. There are no direct functions in golang to do the same, we will have to do some manual fine-tuning of file operations to append a particular text at a specific line.

```go
package main

import (
    "bufio"
    "bytes"
    "log"
    "os"
)

func HandleError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    /* test.txt
    Hi
    Hello
    World
    Gopher
    */
	f, err = os.OpenFile("test.txt", os.O_RDWR, 0660)
	defer f.Close()
	HandleError(err)
	m := bufio.NewScanner(f)
	bytes_till := 0
	// line to be appended
	line_till := 2
	var lines_after string
	var lines_till string
	i := 0
	for m.Scan() {
		line := m.Text()
		if i < line_till {
			bytes_till += bytes.Count([]byte(line), []byte{})
			if i > 0 {
				lines_till += "
"
			}
			lines_till += line
		} else {
			lines_after += "
" + line
		}
		i += 1
	}
	HandleError(m.Err())
	insert_text := lines_till + "
Inserted content"
	insert_text_bytes := bytes.Count([]byte(insert_text), []byte{})
	lines_after_bytes := bytes.Count([]byte(lines_after), []byte{})

	err = os.WriteFile("test.txt", []byte(insert_text), 0660)
	HandleError(err)
	_, err = f.WriteAt([]byte(lines_after), int64(lines_after_bytes)+int64(insert_text_bytes))
	HandleError(err)
    /* test.txt
    Hi
    Hello
    Inserted content
    World
    Gopher
    */
}
```

```plaintext
$ cat test.txt 
Hi
Hello
World
Gophers

$ go run append.go 

$ cat test.txt 
Hi
Hello
Inserted content
World
Gophers
```

We have inserted `Inserted content` after the second line because the `line_till` the variable is set to `2`

So, in the above example, we have used a bunch of packages to append a string or any form of text to a particular line. We first read the contents of the file, by using the `OpenFile` method which will open the file with certain permissions. We need to close the file at the end of the script so we simply use the `defer` keyword before the `f.Close()` method call. We then start to scan the file buffer, by creating a `Scanner` object with the `NewScanner` method. Then, with the scanner object of the file content, we then can use the `Scan()` method to scan the file contents line by line. By converting the content at each line from a slice of bytes to a string using the `Text`, we append to a string `line`, this will be used for keeping the count of bytes for appending text before the newly inserted text.

The `line_till` variable is used for the line number from which we want to append to the text after.

We count the bytes for the current line and add it to the `bytes_till` variable indicating the number of bytes there before appending content. We have a simple if-else check for the first line that is for appending a new line of characters. We append the lines into a single string `lines_till`. The string `insert_text` is created by appending all the lines before the line number `line_till` with the actual content to be inserted. We calculate the number of bytes using the [Count](https://pkg.go.dev/bytes#Count) method in the bytes package. The separator is kept blank. The `lines_after` is also been created as a single string of lines after the line number in the file.

We add the `insert_text` (lines before + inserted text) into the file using the `WriteFile` which will override the contents of the file. Then we append the `lines_after` string as a slice of bytes to the `insert_text_bytes + lines_after_bytes` so we get the byte number position to append the `lines_after` string.

In short, we basically overwrite the file by creating two strings (slice of bytes) one which has the lines before the line number with the text to be inserted and the second string has all the lines after the line number.

## Replace text in a file

Using the [bytes.Replace](https://pkg.go.dev/bytes#Replace) method, we can first read all the bytes and replace the old with the new text, and store them as a slice of bytes. We then write these slices of bytes to the file again, so we first read the contents into slices of bytes, replace the content of the byte and then overwrite the contents with the slice of bytes. It's quite simple.

```go
package main

import (
	"bytes"
	"log"
	"os"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	filename := "test.txt"
	file, err := os.ReadFile(filename)
	HandleError(err)
	old_text := "Hello
World"
	new_text := "Bye"
	new_content := bytes.Replace(file, []byte(old_text), []byte(new_text), -1)
	err = os.WriteFile(filename, new_content, 0660)
	HandleError(err)
}
```

```plaintext
$ cat test.txt
Hi
Hello
World
Gophers

$ go run main.go


$ cat test.txt
Hi
Bye
Gophers
```

As we can see we have replaced `Hello
World` it with `Bye`. So, the \[Replace\] method in the bytes package takes in the parameters like the slice of bytes which should be the actual contents of the file, the old text to be replaced with again as a slice of bytes, and the new text to replace also as the slice of bytes, the final parameter is the number of replacements to be made. Here `-1` indicates there are no limits on how many replacements can be done, it can be `1`, `2` for replacing the first n occurrence of the old text, depending on how many times you want to replace the content in the file.

## Delete Text from a File

We can use the [os.Truncate](https://pkg.go.dev/os#File.Truncate) method to delete the contents of the file. The `Truncate` method takes in the parameters like the file path string and the size of the file to truncate or set to. If we set the second parameter to `0`, the file size will be zero and all the contents will be deleted or removed.

```go
package main

import (
	"log"
	"os"
)

func main() {
    /* test.txt
    Hi
    Hello
    World
    Gophers
    */
	err := os.Truncate("test.txt", 0)
	if err != nil {
		log.Fatal(err)
	}
    /* test.txt is empty
    */
}
```

```plaintext
$ cat test.txt
Hi
Hello
World
Gophers

$ go run delete.go

$ cat test.txt
```

As we can see that the contents of the file are emptied if we set the second parameter(size) of the `Truncate` method as 0.

We can also set the value of the size as the number of bytes to keep, so instead of `0` we can set it to `n` positive integer to only save the first `n` bytes in the file.

```go
package main

import (
	"log"
	"os"
)

func main() {
    /* test.txt
    Hi
    Hello
    World
    Gophers
    */
	err := os.Truncate("test.txt", 6)
	if err != nil {
		log.Fatal(err)
	}
    /* test.txt 
    Hi
    Hel
    */
}
```

```plaintext
$ cat test.txt
Hi
Hello
World
Gophers

$ go run delete.go

$ cat test.txt
Hi
Hel
```

So, in the above example, if we set the size parameter to the Truncate method as `6`, it will keep the size of the file to 6 bytes. So we only see `Hi
Hel`, the new line is a single byte. The rest of the content is deleted. This is how we previously deleted all the bytes from the file by setting the size to `0`.

That's it from this part. Reference for all the code examples and commands can be found in the [100 days of Golang](https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/files/write/) GitHub repository.

## Conclusion

So, from this section of the series, we were able to perform write operations on a file using golang. We used the packages from the standard library and performed operations like write, append, overwrite, delete and replace to a simple text file, but it could have been any file format.

Thank you for reading. If you have any queries, questions, or feedback, you can let me know in the discussion below or on my social handles. Happy Coding :)
