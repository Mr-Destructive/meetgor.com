{
  "title": "Golang: JSON YAML TOML (config) File Reading.",
  "status": "published",
  "slug": "golang-config-file-read",
  "date": "2025-04-05T12:33:35Z"
}

<h2>Reading specific file types (JSON, YAML, TOML)</h2>
<p>In the previous post, we have seen how to read files in golang, in this extended post of that part, we will look into reading some specific files used for configuration and storing data like JSON, YAML, TOML, CSV, etc.</p>
<p>We will see how to read files and get individual pieces in the files. We'll use packages like <code>os</code>, <code>ioutil</code> and <code>encoding</code> to perform reading operations on file and file objects.</p>
<h3>Reading a JSON File</h3>
<p>Golang has built-in support for reading JSON files, but still, we can and need to have low-level controls on how to parse and extract content from the file.</p>
<p>Let's say we have a <code>json</code> file named <code>blog.json</code>, we can use the <a href="https://pkg.go.dev/encoding/json">encoding/json</a> package to convert the JSON data into a GO object (an object that is native and understandable to go). The <a href="https://pkg.go.dev/encoding/json#Unmarshal">Unmarshal</a> function is used to convert the slice of bytes from the file, into a map object.</p>
<pre><code class="language-json">{
    &quot;title&quot;: &quot;Golang Blog Series&quot;,
    &quot;date&quot;: &quot;22nd October 2022&quot;,
    &quot;tags&quot;: [&quot;go&quot;, &quot;files&quot;],
    &quot;words&quot;: 1500,
    &quot;published&quot;: true
}
</code></pre>
<p>The above is a simple JSON file, this file has a few types of key-value pairs like string, list, integer, and boolean. But we can also have nested objects and a list of those nested objects.</p>
<pre><code class="language-go">package main

import (
	&quot;encoding/json&quot;
	&quot;log&quot;
    &quot;os&quot;
)

func main() {

	f, err := os.ReadFile(&quot;blog.json&quot;)
	if err != nil {
		log.Println(err)
	}
	var data map[string]interface{}
	json.Unmarshal([]byte(f), &amp;data)

	log.Println(data)
	for k, v := range data {
		log.Println(k, &quot;:&quot;, v)
	}

}
</code></pre>
<p><strong>I have removed the time stamp from the logs below so as to clearly see the output. We can use <code>fmt</code> to print the simple things while keeping consistent with the rest of the snippets in the series.</strong></p>
<pre><code>$ go run json.go

map[date:22nd October 2022 published:true tags:[go files] title:Golang Blog Series words:1500]
published : true
title : Golang Blog Series
date : 22nd October 2022
tags : [go files]
words : 1500
</code></pre>
<p>The file is read using the <a href="https://pkg.go.dev/os#ReadFile">os.ReadFile</a> method, that takes in a string as a path to the file and returns a slice of bytes or an error if there was an issue in reading the file. The parsed slice of byte is than passed as the first argument to the <code>Unmarshal</code> method in the <code>encoding/json</code> package. The second parameter is the output reference where the parsed JSON will be stored or returned. The function does not return the parsed content instead returns an error if there arose any while parsing the JSON content.</p>
<p>As we can see we have got a map of <code>string</code> with an <code>interface</code>. The interface is used because the value of the key can be anything. There is no fixed value like a <code>string</code>, <code>int</code>, <code>bool</code>, or a nested <code>map</code>, <code>slice</code>. Hence we have mapped the JSON object as a map of <code>string</code> with an <code>interface</code>. The type of the value is identified with the interface it has attached to it. Let's take a look what is the type of each value in the map.</p>
<pre><code>published : true
bool

title : Golang Blog Series
string

date : 22nd October 2022
string

tags : [go files]
[]interface {}

words : 1500
float64
</code></pre>
<p>Here, we can see it has correctly identified the string type of the fields like bool in case of true or false, a string for string type of values, the fourth field however has a list interface attached to it. The default precedence for <code>float64</code> over integer is the reason the <code>1500</code> value is of type <code>float64</code>.</p>
<h3>Reading a YAML File</h3>
<p>Though there is no standard package for parsing/unmarshaling <code>YAML</code> files in golang, it's quite easy to use a third-party package and use it to read YAML files.</p>
<p>The package <a href="https://pkg.go.dev/gopkg.in/yaml.v3">gopkg.in/yaml.v3</a> is used for encoding and decoding YAML files. We'll be just using it for decoding a YAML file by reading it and converting the file object into native Go objects like maps, lists, strings, etc.</p>
<p>The below steps can be used for setting up the package and installing the YAML package locally.</p>
<pre><code>go mod init &lt;your_project_package_name&gt;
go get gopkg.in/yaml.v3
</code></pre>
<p>This should create two files namely <code>go.mod</code> and <code>go.sum</code> with the dependency of the <code>gopkg.in/yaml.v3</code> package.</p>
<pre><code class="language-yml">title: &quot;Golang Blog Series&quot;
date: &quot;22nd October 2022&quot;
tags: [&quot;go&quot;, &quot;files&quot;]
published: false
words: 1500
</code></pre>
<p>The above file is a simple YAML config, we'll follow similar kind of examples for the dummy files used in the examples.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
    &quot;os&quot;

	yaml &quot;gopkg.in/yaml.v3&quot;
)

func main() {

	f, err := os.ReadFile(&quot;blog.yaml&quot;)

	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}

	err = yaml.Unmarshal(f, &amp;data)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(data)
	for k, v := range data {
		log.Println(k, &quot;:&quot;, v)
	}
}
</code></pre>
<pre><code>$ go run yaml.go

map[date:22nd October 2022 published:false tags:[go files] title:Golang Blog Series words:1500]
published : false
words : 1500
title : Golang Blog Series
date : 22nd October 2022
tags : [go files]
</code></pre>
<p>The above code and output demonstrate the usage of the <code>yaml.v3</code> package for reading a YAML file.</p>
<p>Firstly, we read the file into a single-string object with the <a href="https://pkg.go.dev/os#ReadFile">os.ReadFile()</a> method. The method will return a <code>[]byte</code> (slice of byte) or an error. If there is an error, we simply log or panic out of the program, else we can use the <a href="https://pkg.go.dev/gopkg.in/yaml.v3#Unmarshal">yaml.Unmarshal</a> method to convert the string/slice of the byte into a map or a pre-defined struct. In this example, we just keep things simple by storing the file content as <code>map [string, interface{}]</code>, i.e. a map of <code>string</code> and an <code>interface</code>. The key for YAML can be only a string or an integer. It can't have unrestricted data types like the value can have. Though if you want to be unrestrictive, you can use a map of <code>map[interface{}]interface{}</code> to make the key any shape you like to have.</p>
<p>So, we have created a variable called <code>data</code> as a map of <code>string</code> and <code>interface{}</code>, basically key can be a string and the value can be any type of interface depending on the parsed literally from the file object. The <code>Unmarshal</code> function takes in two parameters, the first being the slice of byte i.e. the file contents, and the second being the output variable. Now, the method does not return the parsed YAML, it can return an error if there arose, so we need to set the second parameter as a pointer to the object into which we want to store the parsed YAML.</p>
<p>In the example, we have called <code>Unmarshal(f, &amp;data)</code> which will fetch the contents from the slice of bytes <code>f</code> and output the parsed YAML from the slice of bytes into the memory location of <code>data</code> and hence using <code>&amp;data</code> indicating the pointer to the variable(fetch the memory address).</p>
<p>So, that is how we obtain the map of keys and values from the YAML config, thereafter, you can iterate on the map, access the keys and values, type caste them as per requirement, and basically have control over what processing needs to be done to the parsed YAML content.</p>
<pre><code>published : false
bool

words : 1500
int

title : Golang Blog Series
string

date : 22nd October 2022
string

tags : [go files]
[]interface {}
</code></pre>
<p>I have just printed the types of the values in the above output as <code>log.Printf(&quot;%T&quot;, v)</code>, we can see the types are being correctly recognized and being parsed. The last object is indeed a <code>slice</code> so it has an interface of the slice(array) attached to it.</p>
<h3>Reading a TOML file</h3>
<p>YAML and TOML are almost identical, though toml has more restrictive configuration options, but is more readable than YAML, as YAML can get complicated pretty quickly. Though both of them have their pros and cons, YAML is used everywhere in the DevOps world, configs, whereas TOML is the format of choice for python packaging, and static site generation configs.</p>
<p>Let's see how we can use golang to read TOML files.</p>
<pre><code>$ go mod init &lt;your_project_package_name&gt;
$ go get github.com/pelletier/go-toml
</code></pre>
<p>The above commands are used for setting up a golang package or project and installing the <a href="https://pkg.go.dev/github.com/pelletier/go-toml">go-toml</a> package. Once the above commands are done executing, it will generate <code>go.mod</code> and <code>go.sum</code> files used for storing dependencies and packages installed for the project locally.</p>
<pre><code class="language-toml">[blog]
name='techstructive-blog'
tags=['go','django','vim']
author='meet gor'
active=true

[author]
name='Meet Gor'
github='mr-destructive'
twitter='meetgor21'
posts=80
</code></pre>
<p>The above is the sample file <code>blog.toml</code> which we will use to read in the go script below. The toml file has a similar structure as we have seen in the previous examples. We have different data types like string, boolean, integer, and list.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
    &quot;os&quot;

	toml &quot;github.com/pelletier/go-toml&quot;
)

func main() {

	f, err := os.ReadFile(&quot;blog.toml&quot;)

	if err != nil {
		log.Fatal(err)
	}

	var data map[interface{}]interface{}

	err = toml.Unmarshal(f, &amp;data)
	log.Println(data)

	if err != nil {
		log.Fatal(err)
	}

	for k, v := range data {
		log.Println(k, &quot;:&quot;, v)

		switch t := v.(type) {
		case map[string]interface{}:
			for a, b := range t {
				log.Println(a, &quot;:&quot;, b)
			}
		}
	}
}
</code></pre>
<pre><code>$ go run toml.go

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
</code></pre>
<p>So, in the above example and output, the YAML file was read and the key-value pairs inside them were read. The first thing we do is read the file <code>blog.toml</code> with <code>ioutil</code> package, with the <code>ReadFile</code> function. The function takes in the string as the path to the file to be read and returns a slice of byte. We use this slice of byte as a parameter to the <a href="https://pkg.go.dev/github.com/pelletier/go-toml#Unmarshal">Unmarshal</a> method. The second paramter for the <code>Unmarshal</code> is the output variable(usually a pointer to a variable), we have created a map of <code>interface{]</code> with an <code>interface</code> as we see there can be nested keys which hold the name of the config.</p>
<p>The variable <code>data</code> is a map of <code>interface{}</code> to an <code>interface{}</code>, and we parse the memory address to the <code>data</code> variable to the <code>Unmarshal</code> method. Thereby the parsed <code>TOML</code> content is stored in the data variable.</p>
<pre><code>name : techstructive-blog
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
</code></pre>
<p>The above is a verbose output for the type of the values as parsed by golang, we have <code>string</code>, <code>bool</code>, <code>int64</code>, and a <code>slice</code> (list with interface{} attached with it). Only types like <code>string</code>, <code>bool</code>, <code>int</code>, <code>float64</code> can be parsed from the Unmarshal function, other than these types, the type will have an interface attached to it.</p>
<p>In such cases, where the type of value is not among the 4 types(string, bool, int float), we can use a pre-defined struct to parse the content from the file. Though it would require a strict structure and predictable response from the parsed file.</p>
<h3>Reading CSV file</h3>
<p>We can even read a CSV file in golang, we have seen in the previous post, we have used custom delimiters in the parsing of the file.</p>
<pre><code class="language-csv">id,name,posts,exp
21,jim,23,2
33,kevin,39,1
45,james,70,2
56,chris,89,3
</code></pre>
<p>The above file is a sample csv file, though the size is too small, we can use it as an example.</p>
<pre><code class="language-go">package main

import (
	&quot;encoding/csv&quot;
    &quot;log&quot;
    &quot;os&quot;
)

func main() {
	f, err := os.Open(&quot;temp.csv&quot;)
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
</code></pre>
<pre><code>$ go run main.go
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
</code></pre>
<p>The CSV package has a <a href="https://pkg.go.dev/encoding/csv#NewReader">NewReader</a> method that accepts an <code>io.Reader</code> and returns a <code>Reader</code> object. After parsing the reader, we use the <a href="https://pkg.go.dev/encoding/csv#Reader.ReadAll">ReadAll</a> method to return a 2d string or an error if there exists an error while parsing the content. You can get a detailed explanation of the CSV parsing and reading in the <a href="https://www.meetgor.com/golang-file-read/#Read%20File%20by%20a%20delimiter">previous post</a>.</p>
<h3>Reading CSV from URL</h3>
<p>The CSV file can also be read from the URL, the content of the file is a <code>response.Body</code> in place of the file object reference, in the previous example, the <a href="https://pkg.go.dev/os#Open">os.Open()</a> method returns a <a href="https://pkg.go.dev/os#File">os.File</a> object.</p>
<p>We use the <code>http.Get(string)</code> method to get the response from the URL for reading the CSV file present on the web.</p>
<pre><code class="language-go">package main

import (
	&quot;encoding/csv&quot;
	&quot;log&quot;
	&quot;net/http&quot;
)

func main() {

	url := &quot;https://github.com/woocommerce/woocommerce/raw/master/sample-data/sample_products.csv&quot;
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
</code></pre>
<pre><code>$ go run csv.go
&lt;feff&gt;ID
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
</code></pre>
<p>So, that's how we can read a CSV file from the URL. By fetching the CSV URL <code>https://github.com/woocommerce/woocommerce/raw/master/sample-data/sample_products.csv</code> from the <a href="https://pkg.go.dev/net/http#Client.Get">http.Get</a> method, this will get us the <a href="https://pkg.go.dev/net/http#Response">response.Body</a> that contains the actual CSV file content. The response than can be parsed to the <code>csv.NewReader(*Os.File).ReadAll()</code> i.e. <a href="https://pkg.go.dev/encoding/csv#Reader.ReadAll">reader.ReadAll()</a>. The function returns a multidimensional slice <code>[][]slice</code> that can be iterated and parsed as per requirement.</p>
<h3>Reading XML file</h3>
<p>XML is the de facto standard for RSS feeds, it is widely used in many places and is still all over the web. We'll see an example to read an XML file locally, but as we saw in the above example, we can also read an RSS link from the web.</p>
<p>Just like CSV, we have <a href="https://pkg.go.dev/encoding/xml">encoding/xml</a>, and the standard library has all the functions used for parsing the XML files.</p>
<p>We will be using a local XML file called <code>rss.xml</code>, and reading the contents from the tags in the file.</p>
<pre><code class="language-xml">&lt;?xml version=&quot;1.0&quot; encoding=&quot;UTF-8&quot; ?&gt;
&lt;channel&gt;
&lt;title&gt;Meet Gor&lt;/title&gt;
&lt;description&gt;Techstructive Blog Feed&lt;/description&gt;
&lt;item&gt;
&lt;title&gt;Why and How to make and use Vim as a text editor and customizable IDE&lt;/title&gt;
&lt;link&gt;https://www.meetgor.com/vim-text-editor-ide&lt;/link&gt;
&lt;/item&gt;
&lt;item&gt;
&lt;title&gt;Setting up Vim for Python&lt;/title&gt;
&lt;link&gt;https://www.meetgor.com/vim-for-python&lt;/link&gt;
&lt;/item&gt;
&lt;/channel&gt;
</code></pre>
<p>The above example is a short part of my blog's <a href="https://www.meetgor.com/rss">rss</a> feed. I have just trimmed the unwanted part and will be just using the tags that we want to fetch.</p>
<pre><code class="language-go">package main

import (
	&quot;encoding/xml&quot;
	&quot;log&quot;
	&quot;os&quot;
)

type Channel struct {
	XMLName     xml.Name `xml:&quot;channel&quot;`
	Title       string   `xml:&quot;title&quot;`
	Description string   `xml:&quot;description&quot;`
	Item        []Item   `xml:&quot;item&quot;`
}

type Item struct {
	XMLName xml.Name `xml:&quot;item&quot;`
	Title   string   `xml:&quot;title&quot;`
	Link    string   `xml:&quot;link&quot;`
}

func check_error(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {

	f, err := os.ReadFile(&quot;rss.xml&quot;)
    check_error(err)
	defer f.Close()

	d := Channel{}
	err = xml.Unmarshal(f, &amp;d)
    check_error(err)

	for _, item := range d.Item {
		log.Println(item.Title)
	}
}
</code></pre>
<pre><code>$ go run xml.go

{{ channel} Meet Gor Techstructive Blog Feed [{{ item} Why and How to make and use Vim as a text editor and customizable IDE https://www.meetgor.com/vim-text-editor-ide} {{ item} Setting up Vim for Python https://www.meetgor.com/vim-for-python}]}

Why and How to make and use Vim as a text editor and customizable IDE
Setting up Vim for Python
</code></pre>
<p>The above example uses a couple of <code>struct</code> like <code>Channel</code> and <code>Item</code> that stores the tag data like <code>title</code>, <code>description</code>, <code>link</code>, etc. Unlike the JSON, YAML, and toml files, XML can't parse the content directly we need a structure to parse into. And that's fine as XML is not much dynamic in terms of config, we usually have standard tags and elements which can be pre-defined in a struct type.</p>
<p>In this example, the RSS feed has a <code>Channel</code> tag with <code>title</code>, <code>description</code>, and <code>item</code>.</p>
<p><strong>NOTE: Use Title case for the fields of the structs. It will make the fields public, I spent a few hours debugging that really :)</strong></p>
<p>So, we define the <code>Channel</code> struct with fields like <code>Title</code> as a string which is a tag in the file as <code>xml:&quot;title&quot;</code>. This means the title in the tag of the XML will be stored in the field as a string in the attribute name as <code>Title</code>. Similarly, we have fields like <code>Description</code> and <code>Item[]</code> this is a list or multiple of <code>item</code> tags that might be present in the XML file. The <code>XMLName</code> is used for identifying the parent tag for the struct, so we use <code>channel</code> for the first struct as it is the first tag appearing of the hierarchy in the XML file.</p>
<p>So, we create an object for the root structure as <code>Channel{}</code> (an empty object instantiated). The <code>xml.Unmarshal</code> function is parsed with the content of the file as <code>data</code> which is a slice of byte as we have seen in the previous examples. The slice of byte is then used in the <code>Unmarshal</code> method as the first parameter and the reference of an empty <code>Channel</code> object as the second parameter. The second parameter will be to store the parsed XML content from the file.</p>
<p>I have a few examples on the GitHub repository covering the reading of files from a URL for the CSV, and XML files. But, this concept in the example, can be applied to JSON, YAML, and other file formats as well.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/Mr-Destructive/100-days-of-golang/tree/main/scripts/files/read/config_files">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So, that's it from this post, we covered how to read specific configuration files like <code>JSON</code>, <code>CSV</code>, <code>YAML</code>, <code>TOML</code>, and <code>XML</code>. We saw how to read a local file and also touched on the concept to read contents from a file on the web with a URL. We also saw how we can use pre-defined structs to parse content from a file, especially for XML.</p>
<p>Thank you for reading. If you have any queries, questions, or feedback, you can let me know in the discussion below or on my social handles. Happy Coding :)</p>
