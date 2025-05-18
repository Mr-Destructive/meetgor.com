{
  "type": "posts",
  "title": "Golang: File Reading",
  "description": "Using system calls to read files in Golang. Reading files using various styles like line by line, word, character, chunk and any other specifc delimiter. Using os, bufio, io, encoding packages for helper funcitons making system calls.",
  "date": "2022-10-23 23:15:00",
  "status": "published",
  "slug": "golang-file-read",
  "tags": [
    "go"
  ],
  "series": [
    "100-days-of-golang"
  ],
  "image_url": "https://meetgor-cdn.pages.dev/golang-022-file-read.png"
}

## Introduction

In the 22nd post of the series, we will be looking into the file-handling process in golang, in the next few posts, we will cover the operations on file using golang. In this first entry of the file handling sub-series, we will understand the `READ` operation with files. We will see different ways to read a file, it can be word by word, line by line, or even custom chink by chunk. 

While dealing with files, we will also use standard library packages such as `os`, `bufio`, etc. We'll also touch on how we can read files from a remote location. Using golang, we will have a low-level interaction with file management but golang also abstracts the most of heavy lifting and management of files for us, so it becomes quite easy to work with files.


## Read the file as a single string (using os.ReadFile)

We can use the [os](https://pkg.go.dev/os) package in golang, in which we have access to the [ReadFile](https://pkg.go.dev/os#ReadFile) funciton. The `ReadFile` function takes in a parameter as a string which should be a file name, it returns a slice of bytes or an error. We have discussed the error handling in the previous part of the series. So, we have to use the comma ok error syntax to get the appropriate return value from the funciton. We can grab the slice of bytes as the text we want or an error if there are errors like a file doesn't exist, it's a folder, etc. 

```go
package main

import (
	"os"
	"log"
)

func main() {

	text, err := os.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(text))
}
```

```
$ go run main.go                                                                                                                                                         
2022/10/23 22:39:11 Golang is a programming language.                                                                                                                    
created: 2007                                                                                                                                                            
type:static 
```

So, under the hood here, we have text as a slice of bytes. We can iterate over the text as a slice and get the character by character-in the content of the file. Though we don't directly interact with the file content, we are storing it in a variable. In technical words, the file is directly loaded into the memory at once. We thereby return a single string object containing the content of the file.


## Read file line by line 

We can even read a file line by line. Using the [bufio.NewScanner()](https://pkg.go.dev/bufio#NewScanner), the function takes in a [Reader](https://pkg.go.dev/io#Reader) object in our case it will be a file object. The function returns a scanner object that can be used to read the text with a particular scanner method. The returned object can be used in the loop to iterate over the content, in our case, we use the [Scan](https://pkg.go.dev/bufio#Scanner.Scan) method to split the file into lines. But we can use other methods like [ScanWords](https://pkg.go.dev/bufio#ScanWords) for scanning words, [ScanRunes](https://pkg.go.dev/bufio#ScanRunes) for scanning character by character, [ScanBytes](https://pkg.go.dev/bufio#ScanBytes) for scanning byte by byte. 

```go
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	f, err := os.Open("sample.txt")

	if err != nil {
		log.Fatal(err)
	}
	line_list := []string{}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line_list = append(line_list, line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, line := range line_list {
		log.Println(line)
	}
}
```

```
$ go run line.go

2022/10/23 22:39:50 Golang is a programming language.
2022/10/23 22:39:50 created: 2007
2022/10/23 22:39:50 type:static
```

In the above example, the file is read in with the `bufio.NewScanner` and is iterated line by line with the help of the `Scan` function. The text in the line is scanned and stored in the variable `line` which is a string, this is further appended to the string slice `line_list`. Hence we can iterate over the file content line by line and store the results as a string array.

Here we have used the `defer` keyword before calling the `f.Close()` method because we want to close the file after performing operations on the file. The defer will call the function at almost the end of the main function i.e. at the end of the program.

## Read File by a delimiter

We can even read a file with a custom delimiter which can be used to read CSV or other delimiters. The [csv](https://pkg.go.dev/encoding/csv) package has a [NewReader](https://pkg.go.dev/encoding/csv#NewReader) function which takes in a object of file content, and it will return a [Reader](https://pkg.go.dev/encoding/csv#Reader) object. We can alter the attribute `Comma` in the `Reader` object and set it to any value we want as a delimiter. Thereafter we can read the whole content or read it as words, lines, bytes, or chunks as per your criteria. The extracted content will be split as a slice of the data separated by the delimiter set in the `Comma` attribute.

```go
package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {

	f, err := os.Open("delimeter.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	reader.Comma = ':'
	
    data, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	log.Println(data)
}
```

```
$ cat delimiter.txt

10:22:2022
golang:21:read


$ go run delimiter.go

2022/10/23 22:40:44 [[10 22 2022] [golang 21 read]]
```

In the above example, the delimiter is set as `:` so by using the `Comma` attribute we can set the delimiter. By using the `NewReader` function, we fetch the reader object and by using the `ReadAll` function associated to the reader object, we read the contents. The content is fetched as a slice of strings which will be separated by the delimiter.


### Reading File word by word

We can even use [ScanWords](https://pkg.go.dev/bufio#ScanWords) to read a file word by word. A word can be a collection of characters that are separated by space. Instead of reading it line by line, this function reads the file content after a space. 

```go
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wordlist := []string{}
	for scanner.Scan() {
		word := scanner.Text()
		wordlist = append(wordlist, word)
		log.Println(word)
	}
	log.Println(wordlist)

}
```

```
$ go run word.go

2022/10/23 22:42:03 Golang
2022/10/23 22:42:03 is
2022/10/23 22:42:03 a
2022/10/23 22:42:03 programming
2022/10/23 22:42:03 language.
2022/10/23 22:42:03 created:
2022/10/23 22:42:03 2007
2022/10/23 22:42:03 type:static

2022/10/23 22:42:03 [Golang is a programming language. created: 2007 type:static] 
```

With the `ScanWords` function, we can read the contents of the file word by word. The scanner object which has the actual content of the file is split by the `Split` function, the split criteria are used as Word, where the delimiter will be used as space. The `wordlist` is a slice of strings to which we append the string `word` that in turn is read from the `scanner.Text()` function.

### Reading Files in chunks

We can even read files in chunks, a chunk is a collection/array of bytes. We can specify the number of bytes we want to read in one go and the file reader will scan the content as a slice of that number of bytes each iteration. The [Read](https://pkg.go.dev/bufio#Reader.Read) funciton takes in a slice of bytes and will return the number of bytes in the reader object.

```go
package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {

	f, err := os.Open("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	chunk_size := 16
	chunk_list := []string{}
	buf := make([]byte, chunk_size)

	for {
		n, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		chunk_list = append(chunk_list, string(buf[0:n]))
	}
	for _, chunk := range chunk_list {
		log.Print(chunk)
	}
}
```

```
$ go run chunks.go

2022/10/23 22:44:41 Golang is a prog
2022/10/23 22:44:41 ramming language
2022/10/23 22:44:41 .
created: 2007
2022/10/23 22:44:41 type:static
```

In the above example, we have opened the file and loaded the content into the `f` variable. the contents are read with the help of the `NewReader` function which returns a reader object which further can be used to read contents into chunks of bytes. The `chunk_size` defines the size we want to use for reading the content, `chunk_list` as a slice of strings which will hold the slice of chunks/bytes as a type caste into a slice of strings. With the `Read` function, the bytes are read into the function, and the buffer is split as per the chunk size obtained in the `Read` function. We append the slice of bytes into the sliced array and thereby we obtain the slice of strings.


### Read file character by character

We can even read file each character at a time, using the [ScanRunes](https://pkg.go.dev/bufio#ScanRunes) function, this function scans a single rune/byte at a time. So, we can scan these runes one at a time and store them as a slice of bytes. Thereby we will have the content of the file stored as a slice of bytes.

```go
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	charlist := []byte{}
	for scanner.Scan() {
		char := byte(scanner.Text()[0])
		charlist = append(charlist, char)
	}
	log.Println(charlist)

}
```

```
$ go run char.go

2022/10/23 22:48:55 [71 111 108 97 110 103 32 105 115 32 97 32 112 114 111 103 114 97 109 109 105 110 103 32 108 97 110 103 117 97 103 101 46 10 99 114 101 97 116 101 100 58 32 50 48 48 55 10 116 121 112 101 58 115 116 97 116 105 99 10]
```

We can see in the above example, the output is a slice of byte, hence these are `uint8`, we can cast them to `string` and obtain the equivalent ASCII representation of the bytes. The `ScanRunes` function allows us to read the content from the reader object as a rune as we split the reader object into the unit bytes/runes.

That's it from this part. Reference for all the code examples and commands can be found in the [100 days of Golang](https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/files/read/) GitHub repository.

## Conclusion

In this section, we explored the functions and packages related to file reading. We saw how we can use packages like `os`, `bufio`, `encoding`, etc. to read files in a different way. We saw how to read files as a single string, line by line, word by word, character by character, in chunks, and also with a custom delimiter. Hopefully, the basics of file reading will have been cleared and with the examples, the syntactical construct was understood. 

Thank you for reading. If you have any queries, questions, or feedback, you can let me know in the discussion below or on my social handles. Happy Coding :)
