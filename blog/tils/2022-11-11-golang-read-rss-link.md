---
templateKey: til
title: "Read a Rss Feed with a URL in Golang"
description: "Reading Rss Feed with a Rss XML Link/URL in golang using encoding package"
status: published-til
slug: golang-read-rss-feed
tags: ["go", ]
date: 2022-11-11 22:45:00
---

## Reding Rss Feed

We can use golang's [encoding/xml](https://pkg.go.dev/encoding/xml) package to read a Rss feed. Though we have to be speicific of what type of structure the Rss feed has, so it is not dynamic but it works really well with structs. I have covered a few nuances of reading XML file in the [config file reading](https://www.meetgor.com/golang-config-file-read/#reading-xml-file) post of the 100 days of golang series.

### Get request to Rss feed

We first need to send a `GET` request to the Rss feed, we can use the [http](https://pkg.go.dev/net/http) package to grab the response.

```go
package main

import (
	"log"
	"net/http"
)

func main() {

	url := "https://meetgor.com/rss.xml"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

    log.Println(response.Body)
	defer response.Body.Close()

}
```

So, in the above example, we have used the `net/http` package to send a `GET` request with the [Get](https://pkg.go.dev/net/http#Get) funciton. The function takes in a string as a `URL` and returns either the object as response or an error. If there arose any error, we simply exit out of the program and log the error. If the error is `nil`, we return the response in the `response` variable. This builds up a good foundation for the next step to read the response body and fetching the actual bytes from the response.

### Fetch the content from the Link

Since we have the `response` object, we can use the [io.ReadAll](https://pkg.go.dev/io#ReadAll) function to read the bytes in the response body. The function takes in the [Reader](https://pkg.go.dev/io#Reader) object in this case it is [ReadCloser](https://pkg.go.dev/io#ReadCloser) object as a http object. The function then returns the slice of bytes/int8. The slice then can be interpreted as string or other form data that can be used for parsing the xml from the response.

```go
package main

import (
	"log"
	"net/http"
    "io"
)

func main() {

	url := "https://meetgor.com/rss.xml"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

    log.Println(string(data))
    log.Printf("Type -> %T", data)
}
```

```
<rss>
    <channel>
        <item>
        ...
        ...
        ...
        </item>
    </channel>
</rss>


Type -> []uint8 

```

So, we can see that the parsed content is indeed xml, it is type casted to string from the slice of bytes. This can be further used for the parsing the text as Rss structure and fetch the required details.

## Parsing Rss with a struct

We can now move into creating a struct for individual tags required in the parsing.

```go
package main

import (
    "encoding/xml"
	"io"
	"log"
	"net/http"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Item        []Item   `xml:"item"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
}

func main() {

	url := "https://meetgor.com/rss.xml"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

    log.Println(string(data))
}
```

If you would look at the [rss feed](https://meetgor.com/rss.xml), you can see it has a structure of tags and elements. The `rss` tag is the root tag, followed by `channel` and other types of nested tags speicific for the type of information to be stored like `title` for the title in the feed, `link` for the link to the feed, etc. 

So, we create those as structure, the root structure is the `Rss` which we will create with a few attributes like `Channel` and the name of the current tag. In the `Rss` case the name of the tag/element is `rss`, so it is given the `xml.Name` as `xml:'rss'` in backticks indicating the type hint for the field. The next field is the `Channel` which is another type(custom type struct). We have defined `Channel` as a struct just after it that will hold information like the `title`, `description` of the website. We also have the `xml.Name` as `xml:"channel"` which indicates the current struct is representation of `channel` tag in the rss feed. Finally, we also have a custom type struct as `Item`. The `Item` struct has a few attributes like `Title`, `Link` and you can now start to see the pattern, you can customize it as per your requirements and speicifications.

```go
package main

import (
    "encoding/xml"
	"io"
	"log"
	"net/http"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Item        []Item   `xml:"item"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
}

func main() {

	url := "https://meetgor.com/rss.xml"
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

    // New Code

	d := Rss{}
	err = xml.Unmarshal(data, &d)

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range d.Channel.Item {
		log.Println(item.Title)
	}
}
```

```
$ go run main.go

Why and How to make and use Vim as a text editor and customizable IDE
Setting up Vim for Python
Setting up Vim for BASH Scripting
Vim: Keymapping Guide
...
...
...
Django + HTMX CRUD application
PGCLI: Postgres from the terminal
Golang: Closures
Golang: Interfaces
Golang: Error Handling
Golang: Paths
Golang: File Reading
Golang: JSON YAML TOML (config) File Reading.
```

So, here we have initialized the `Rss` struct as empty and then used the [Unmarshal](https://pkg.go.dev/encoding/xml#Unmarshal) method in the `xml` package. The Unmarshal method will parse the data as per the type of either int, float, bool or string, any other type of data will be discarded as interface or struct. We can usually parse any valid type of data into `Unmarshal` method and it generally gives a proper expected outcome.

The Unmarshal method takes in the slice of byte and the second paramter as pointer to a struct or any variable that will store the parsed xml content from the slice of byte. The function just returns the error type, either `nil` in case of no errors, and returns the actual error obejct if there arise any type of error.

So we parse the `data` which is a slice of byte to the funciton and the reference to the `d` object which is a empty `Rss` object. This will get us the data in the `d` object. We can then iterate over the object as per the struct and use the perform operations like type casting or converting types, etc to get your required data back.

In the above example, we simply iterate over the `d.Channel.Item` which is a list of elements of tag `item` in the rss feed. Inside the for loop, we can access the object and simply print or perform any sort of operations. I have simply printed the list of articles with titles. 

Links for the code available on the [100 days of golang](https://github.com/Mr-Destructive/100-days-of-golang/blob/main/scripts/files/read/config_files/xml/rss.go) GitHub repository.

So, that's how we parse an XML feed in golang. Just plug and play if you have a similar type of structure of the Rss XML feed. Happy Coding :)
