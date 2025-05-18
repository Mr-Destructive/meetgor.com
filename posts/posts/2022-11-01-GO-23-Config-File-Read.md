{
  "type": "posts",
  "title": "Golang: JSON YAML TOML (config) File Reading.",
  "description": "Reading specific files used generally for configuration and storing data, as well as for web development. Reading file formats like JSON, YAML, TOML, CSV, and XML.",
  "date": "2022-11-01 23:00:00",
  "status": "published",
  "slug": "golang-config-file-read",
  "tags": [
    "go"
  ],
  "series": [
    "100-days-of-golang"
  ],
  "image_url": "https://meetgor-cdn.pages.dev/golang-023-config-files.png"
}

## Reading specific file types (JSON, YAML, TOML)

In the previous post, we have seen how to read files in golang, in this extended post of that part, we will look into reading some specific files used for configuration and storing data like JSON, YAML, TOML, CSV, etc.

We will see how to read files and get individual pieces in the files. We'll use packages like `os`, `ioutil` and `encoding` to perform reading operations on file and file objects.

### Reading a JSON File

Golang has built-in support for reading JSON files, but still, we can and need to have low-level controls on how to parse and extract content from the file.

Let's say we have a `json` file named `blog.json`, we can use the [encoding/json](https://pkg.go.dev/encoding/json) package to convert the JSON data into a GO object (an object that is native and understandable to go). The [Unmarshal](https://pkg.go.dev/encoding/json#Unmarshal) function is used to convert the slice of bytes from the file, into a map object.


```json
{
    "title": "Golang Blog Series",
    "date": "22nd October 2022",
    "tags": ["go", "files"],
    "words": 1500,
    "published": true
}
```

The above is a simple JSON file, this file has a few types of key-value pairs like string, list, integer, and boolean. But we can also have nested objects and a list of those nested objects.

```go
package main

import (
	"encoding/json"
	"log"
    "os"
)

func main() {

	f, err := os.ReadFile("blog.json")
	if err != nil {
		log.Println(err)
	}
	var data map[string]interface{}
	json.Unmarshal([]byte(f), &data)

	log.Println(data)
	for k, v := range data {
		log.Println(k, ":", v)
	}

}
```

**I have removed the time stamp from the logs below so as to clearly see the output. We can use `fmt` to print the simple things while keeping consistent with the rest of the snippets in the series.**

```
$ go run json.go

map[date:22nd October 2022 published:true tags:[go files] title:Golang Blog Series words:1500]
published : true
title : Golang Blog Series
date : 22nd October 2022
tags : [go files]
words : 1500
```
The file is read using the [os.ReadFile](https://pkg.go.dev/os#ReadFile) method, that takes in a string as a path to the file and returns a slice of bytes or an error if there was an issue in reading the file. The parsed slice of byte is than passed as the first argument to the `Unmarshal` method in the `encoding/json` package. The second parameter is the output reference where the parsed JSON will be stored or returned. The function does not return the parsed content instead returns an error if there arose any while parsing the JSON content.

As we can see we have got a map of `string` with an `interface`. The interface is used because the value of the key can be anything. There is no fixed value like a `string`, `int`, `bool`, or a nested `map`, `slice`. Hence we have mapped the JSON object as a map of `string` with an `interface`. The type of the value is identified with the interface it has attached to it. Let's take a look what is the type of each value in the map.

```
published : true
bool

title : Golang Blog Series
string

date : 22nd October 2022
string

tags : [go files]
[]interface {}

words : 1500
float64
```

Here, we can see it has correctly identified the string type of the fields like bool in case of true or false, a string for string type of values, the fourth field however has a list interface attached to it. The default precedence for `float64` over integer is the reason the `1500` value is of type `float64`. 


### Reading a YAML File

Though there is no standard package for parsing/unmarshaling `YAML` files in golang, it's quite easy to use a third-party package and use it to read YAML files.

The package [gopkg.in/yaml.v3](https://pkg.go.dev/gopkg.in/yaml.v3) is used for encoding and decoding YAML files. We'll be just using it for decoding a YAML file by reading it and converting the file object into native Go objects like maps, lists, strings, etc.

The below steps can be used for setting up the package and installing the YAML package locally.

```
go mod init <your_project_package_name>
go get gopkg.in/yaml.v3
```

This should create two files namely `go.mod` and `go.sum` with the dependency of the `gopkg.in/yaml.v3` package.

```yml
title: "Golang Blog Series"
date: "22nd October 2022"
tags: ["go", "files"]
published: false
words: 1500
```

The above file is a simple YAML config, we'll follow similar kind of examples for the dummy files used in the examples.

```go
package main

import (
	"log"
    "os"

	yaml "gopkg.in/yaml.v3"
)

func main() {

	f, err := os.ReadFile("blog.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}

	err = yaml.Unmarshal(f, &data)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(data)
	for k, v := range data {
		log.Println(k, ":", v)
	}
}
```

```
$ go run yaml.go

map[date:22nd October 2022 published:false tags:[go files] title:Golang Blog Series words:1500]
published : false
words : 1500
title : Golang Blog Series
date : 22nd October 2022
tags : [go files]
```

The above code and output demonstrate the usage of the `yaml.v3` package for reading a YAML file.

Firstly, we read the file into a single-string object with the [os.ReadFile()](https://pkg.go.dev/os#ReadFile) method. The method will return a `[]byte` (slice of byte) or an error. If there is an error, we simply log or panic out of the program, else we can use the [yaml.Unmarshal](https://pkg.go.dev/gopkg.in/yaml.v3#Unmarshal) method to convert the string/slice of the byte into a map or a pre-defined struct. In this example, we just keep things simple by storing the file content as `map [string, interface{}]`, i.e. a map of `string` and an `interface`. The key for YAML can be only a string or an integer. It can't have unrestricted data types like the value can have. Though if you want to be unrestrictive, you can use a map of `map[interface{}]interface{}` to make the key any shape you like to have.

So, we have created a variable called `data` as a map of `string` and `interface{}`, basically key can be a string and the value can be any type of interface depending on the parsed literally from the file object. The `Unmarshal` function takes in two parameters, the first being the slice of byte i.e. the file contents, and the second being the output variable. Now, the method does not return the parsed YAML, it can return an error if there arose, so we need to set the second parameter as a pointer to the object into which we want to store the parsed YAML.

In the example, we have called `Unmarshal(f, &data)` which will fetch the contents from the slice of bytes `f` and output the parsed YAML from the slice of bytes into the memory location of `data` and hence using `&data` indicating the pointer to the variable(fetch the memory address).

So, that is how we obtain the map of keys and values from the YAML config, thereafter, you can iterate on the map, access the keys and values, type caste them as per requirement, and basically have control over what processing needs to be done to the parsed YAML content.

```
published : false
bool

words : 1500
int

title : Golang Blog Series
string

date : 22nd October 2022
string

tags : [go files]
[]interface {}
```

I have just printed the types of the values in the above output as `log.Printf("%T", v)`, we can see the types are being correctly recognized and being parsed. The last object is indeed a `slice` so it has an interface of the slice(array) attached to it.

### Reading a TOML file

YAML and TOML are almost identical, though toml has more restrictive configuration options, but is more readable than YAML, as YAML can get complicated pretty quickly. Though both of them have their pros and cons, YAML is used everywhere in the DevOps world, configs, whereas TOML is the format of choice for python packaging, and static site generation configs.

Let's see how we can use golang to read TOML files.

```
$ go mod init <your_project_package_name>
$ go get github.com/pelletier/go-toml
```

The above commands are used for setting up a golang package or project and installing the [go-toml](https://pkg.go.dev/github.com/pelletier/go-toml) package. Once the above commands are done executing, it will generate `go.mod` and `go.sum` files used for storing dependencies and packages installed for the project locally.

```toml
[blog]
name='techstructive-blog'
tags=['go','django','vim']
author='meet gor'
active=true

[author]
name='Meet Gor'
github='mr-destructive'
twitter='meetgor21'
posts=80
```

The above is the sample file `blog.toml` which we will use to read in the go script below. The toml file has a similar structure as we have seen in the previous examples. We have different data types like string, boolean, integer, and list.

```go
package main

import (
	"log"
    "os"

	toml "github.com/pelletier/go-toml"
)

func main() {

	f, err := os.ReadFile("blog.toml")

	if err != nil {
		log.Fatal(err)
	}

	var data map[interface{}]interface{}

	err = toml.Unmarshal(f, &data)
	log.Println(data)

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range data {
		log.Println(k, ":", v)

		switch t := v.(type) {
		case map[string]interface{}:
			for a, b := range t {
				log.Println(a, ":", b)
			}
		}
	}
}
```

```
$ go run toml.go

map[author:map[github:mr-destructive name:Meet Gor posts:80 twitter:meetgor21] blog:map[active:true author:meet gor

name:techstructive-blog tags:[go django vim]]]

blog : map[active:true author:meet gor name:techstructive-blog tags:[go django vim]]
name : techstructive-blog
tags : [go django vim]
author : meet gor
active : true

author : map[github:mr-destructive name:Meet Gor posts:80 twitter:meetgor21]

name : Meet Gor
github : mr-destructive
twitter : meetgor21
posts : 80
```

So, in the above example and output, the YAML file was read and the key-value pairs inside them were read. The first thing we do is read the file `blog.toml` with `ioutil` package, with the `ReadFile` function. The function takes in the string as the path to the file to be read and returns a slice of byte. We use this slice of byte as a parameter to the [Unmarshal](https://pkg.go.dev/github.com/pelletier/go-toml#Unmarshal) method. The second paramter for the `Unmarshal` is the output variable(usually a pointer to a variable), we have created a map of `interface{]` with an `interface` as we see there can be nested keys which hold the name of the config.

The variable `data` is a map of `interface{}` to an `interface{}`, and we parse the memory address to the `data` variable to the `Unmarshal` method. Thereby the parsed `TOML` content is stored in the data variable.

```
name : techstructive-blog
string

tags : [go django vim]
[]interface{}

author : meet gor
string

active : true
bool

name : Meet Gor
string

github : mr-destructive
string

twitter : meetgor21
string

posts : 80
int64
```

The above is a verbose output for the type of the values as parsed by golang, we have `string`, `bool`, `int64`, and a `slice` (list with interface{} attached with it). Only types like `string`, `bool`, `int`, `float64` can be parsed from the Unmarshal function, other than these types, the type will have an interface attached to it.
 
 In such cases, where the type of value is not among the 4 types(string, bool, int float), we can use a pre-defined struct to parse the content from the file. Though it would require a strict structure and predictable response from the parsed file.

### Reading CSV file

We can even read a CSV file in golang, we have seen in the previous post, we have used custom delimiters in the parsing of the file.

```csv
id,name,posts,exp
21,jim,23,2
33,kevin,39,1
45,james,70,2
56,chris,89,3
```

The above file is a sample csv file, though the size is too small, we can use it as an example.

```go
package main

import (
	"encoding/csv"
    "log"
    "os"
)

func main() {
	f, err := os.Open("temp.csv")
	check_error(err)

	reader := csv.NewReader(f)

	n, err := reader.ReadAll()
	check_error(err)
	for _, line := range n {
		for _, text := range line {
			log.Println(text)
		}
	}
}
```

```
$ go run main.go
id
name
posts
exp
21
jim
23
2
33
kevin
39
1
45
james
70
2
56
chris
89
3
```

The CSV package has a [NewReader](https://pkg.go.dev/encoding/csv#NewReader) method that accepts an `io.Reader` and returns a `Reader` object. After parsing the reader, we use the [ReadAll](https://pkg.go.dev/encoding/csv#Reader.ReadAll) method to return a 2d string or an error if there exists an error while parsing the content. You can get a detailed explanation of the CSV parsing and reading in the [previous post](https://www.meetgor.com/golang-file-read/#Read%20File%20by%20a%20delimiter).


### Reading CSV from URL

The CSV file can also be read from the URL, the content of the file is a `response.Body` in place of the file object reference, in the previous example, the [os.Open()](https://pkg.go.dev/os#Open) method returns a [os.File](https://pkg.go.dev/os#File) object. 

We use the `http.Get(string)` method to get the response from the URL for reading the CSV file present on the web.


```go
package main

import (
	"encoding/csv"
	"log"
	"net/http"
)

func main() {

	url := "https://github.com/woocommerce/woocommerce/raw/master/sample-data/sample_products.csv"
	response, err := http.Get(url)

	if err != nil {
		log.Println(err)
		return
	}

	defer response.Body.Close()

	reader := csv.NewReader(response.Body)
	n, err := reader.ReadAll()

	if err != nil {
		log.Println(err)
	}

	for _, line := range n {
		for _, text := range line {
			log.Println(text)
		}
	}
}
```

```
$ go run csv.go
<feff>ID
Type
SKU
Name
Published
Is featured?
Visibility in catalog
Short description
Description
Date sale price starts
Date sale price ends
...
...
...
```
So, that's how we can read a CSV file from the URL. By fetching the CSV URL `https://github.com/woocommerce/woocommerce/raw/master/sample-data/sample_products.csv` from the [http.Get](https://pkg.go.dev/net/http#Client.Get) method, this will get us the [response.Body](https://pkg.go.dev/net/http#Response) that contains the actual CSV file content. The response than can be parsed to the `csv.NewReader(*Os.File).ReadAll()` i.e. [reader.ReadAll()](https://pkg.go.dev/encoding/csv#Reader.ReadAll). The function returns a multidimensional slice `[][]slice` that can be iterated and parsed as per requirement.

### Reading XML file

XML is the de facto standard for RSS feeds, it is widely used in many places and is still all over the web. We'll see an example to read an XML file locally, but as we saw in the above example, we can also read an RSS link from the web.

Just like CSV, we have [encoding/xml](https://pkg.go.dev/encoding/xml), and the standard library has all the functions used for parsing the XML files.

We will be using a local XML file called `rss.xml`, and reading the contents from the tags in the file.

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<channel>
<title>Meet Gor</title>
<description>Techstructive Blog Feed</description>
<item>
<title>Why and How to make and use Vim as a text editor and customizable IDE</title>
<link>https://www.meetgor.com/vim-text-editor-ide</link>
</item>
<item>
<title>Setting up Vim for Python</title>
<link>https://www.meetgor.com/vim-for-python</link>
</item>
</channel>
```

The above example is a short part of my blog's [rss](https://www.meetgor.com/rss) feed. I have just trimmed the unwanted part and will be just using the tags that we want to fetch. 

```go
package main

import (
	"encoding/xml"
	"log"
	"os"
)

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

func check_error(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {

	f, err := os.ReadFile("rss.xml")
    check_error(err)
	defer f.Close()

	d := Channel{}
	err = xml.Unmarshal(f, &d)
    check_error(err)

	for _, item := range d.Item {
		log.Println(item.Title)
	}
}
```

```
$ go run xml.go

{{ channel} Meet Gor Techstructive Blog Feed [{{ item} Why and How to make and use Vim as a text editor and customizable IDE https://www.meetgor.com/vim-text-editor-ide} {{ item} Setting up Vim for Python https://www.meetgor.com/vim-for-python}]}

Why and How to make and use Vim as a text editor and customizable IDE
Setting up Vim for Python
```

The above example uses a couple of `struct` like `Channel` and `Item` that stores the tag data like `title`, `description`, `link`, etc. Unlike the JSON, YAML, and toml files, XML can't parse the content directly we need a structure to parse into. And that's fine as XML is not much dynamic in terms of config, we usually have standard tags and elements which can be pre-defined in a struct type.

In this example, the RSS feed has a `Channel` tag with `title`, `description`, and `item`. 

**NOTE: Use Title case for the fields of the structs. It will make the fields public, I spent a few hours debugging that really :)**

So, we define the `Channel` struct with fields like `Title` as a string which is a tag in the file as `xml:"title"`. This means the title in the tag of the XML will be stored in the field as a string in the attribute name as `Title`. Similarly, we have fields like `Description` and `Item[]` this is a list or multiple of `item` tags that might be present in the XML file. The `XMLName` is used for identifying the parent tag for the struct, so we use `channel` for the first struct as it is the first tag appearing of the hierarchy in the XML file.

So, we create an object for the root structure as `Channel{}` (an empty object instantiated). The `xml.Unmarshal` function is parsed with the content of the file as `data` which is a slice of byte as we have seen in the previous examples. The slice of byte is then used in the `Unmarshal` method as the first parameter and the reference of an empty `Channel` object as the second parameter. The second parameter will be to store the parsed XML content from the file.
 
I have a few examples on the GitHub repository covering the reading of files from a URL for the CSV, and XML files. But, this concept in the example, can be applied to JSON, YAML, and other file formats as well.

That's it from this part. Reference for all the code examples and commands can be found in the [100 days of Golang](https://github.com/Mr-Destructive/100-days-of-golang/tree/main/scripts/files/read/config_files) GitHub repository.

## Conclusion

So, that's it from this post, we covered how to read specific configuration files like `JSON`, `CSV`, `YAML`, `TOML`, and `XML`. We saw how to read a local file and also touched on the concept to read contents from a file on the web with a URL. We also saw how we can use pre-defined structs to parse content from a file, especially for XML.

Thank you for reading. If you have any queries, questions, or feedback, you can let me know in the discussion below or on my social handles. Happy Coding :)
