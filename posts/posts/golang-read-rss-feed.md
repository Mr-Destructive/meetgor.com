{
  "title": "Read a Rss Feed with a URL in Golang",
  "status": "published",
  "slug": "golang-read-rss-feed",
  "date": "2025-04-05T12:33:34Z"
}

<h2>Reding Rss Feed</h2>
<p>We can use golang's <a href="https://pkg.go.dev/encoding/xml">encoding/xml</a> package to read a Rss feed. Though we have to be speicific of what type of structure the Rss feed has, so it is not dynamic but it works really well with structs. I have covered a few nuances of reading XML file in the <a href="https://www.meetgor.com/golang-config-file-read/#reading-xml-file">config file reading</a> post of the 100 days of golang series.</p>
<h3>Get request to Rss feed</h3>
<p>We first need to send a <code>GET</code> request to the Rss feed, we can use the <a href="https://pkg.go.dev/net/http">http</a> package to grab the response.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;net/http&quot;
)

func main() {

	url := &quot;https://meetgor.com/rss.xml&quot;
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

    log.Println(response.Body)
	defer response.Body.Close()

}
</code></pre>
<p>So, in the above example, we have used the <code>net/http</code> package to send a <code>GET</code> request with the <a href="https://pkg.go.dev/net/http#Get">Get</a> funciton. The function takes in a string as a <code>URL</code> and returns either the object as response or an error. If there arose any error, we simply exit out of the program and log the error. If the error is <code>nil</code>, we return the response in the <code>response</code> variable. This builds up a good foundation for the next step to read the response body and fetching the actual bytes from the response.</p>
<h3>Fetch the content from the Link</h3>
<p>Since we have the <code>response</code> object, we can use the <a href="https://pkg.go.dev/io#ReadAll">io.ReadAll</a> function to read the bytes in the response body. The function takes in the <a href="https://pkg.go.dev/io#Reader">Reader</a> object in this case it is <a href="https://pkg.go.dev/io#ReadCloser">ReadCloser</a> object as a http object. The function then returns the slice of bytes/int8. The slice then can be interpreted as string or other form data that can be used for parsing the xml from the response.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;net/http&quot;
    &quot;io&quot;
)

func main() {

	url := &quot;https://meetgor.com/rss.xml&quot;
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

    log.Println(string(data))
    log.Printf(&quot;Type -&gt; %T&quot;, data)
}
</code></pre>
<pre><code>&lt;rss&gt;
    &lt;channel&gt;
        &lt;item&gt;
        ...
        ...
        ...
        &lt;/item&gt;
    &lt;/channel&gt;
&lt;/rss&gt;


Type -&gt; []uint8 

</code></pre>
<p>So, we can see that the parsed content is indeed xml, it is type casted to string from the slice of bytes. This can be further used for the parsing the text as Rss structure and fetch the required details.</p>
<h2>Parsing Rss with a struct</h2>
<p>We can now move into creating a struct for individual tags required in the parsing.</p>
<pre><code class="language-go">package main

import (
    &quot;encoding/xml&quot;
	&quot;io&quot;
	&quot;log&quot;
	&quot;net/http&quot;
)

type Rss struct {
	XMLName xml.Name `xml:&quot;rss&quot;`
	Channel Channel  `xml:&quot;channel&quot;`
}

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

func main() {

	url := &quot;https://meetgor.com/rss.xml&quot;
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
</code></pre>
<p>If you would look at the <a href="https://meetgor.com/rss.xml">rss feed</a>, you can see it has a structure of tags and elements. The <code>rss</code> tag is the root tag, followed by <code>channel</code> and other types of nested tags speicific for the type of information to be stored like <code>title</code> for the title in the feed, <code>link</code> for the link to the feed, etc.</p>
<p>So, we create those as structure, the root structure is the <code>Rss</code> which we will create with a few attributes like <code>Channel</code> and the name of the current tag. In the <code>Rss</code> case the name of the tag/element is <code>rss</code>, so it is given the <code>xml.Name</code> as <code>xml:'rss'</code> in backticks indicating the type hint for the field. The next field is the <code>Channel</code> which is another type(custom type struct). We have defined <code>Channel</code> as a struct just after it that will hold information like the <code>title</code>, <code>description</code> of the website. We also have the <code>xml.Name</code> as <code>xml:&quot;channel&quot;</code> which indicates the current struct is representation of <code>channel</code> tag in the rss feed. Finally, we also have a custom type struct as <code>Item</code>. The <code>Item</code> struct has a few attributes like <code>Title</code>, <code>Link</code> and you can now start to see the pattern, you can customize it as per your requirements and speicifications.</p>
<pre><code class="language-go">package main

import (
    &quot;encoding/xml&quot;
	&quot;io&quot;
	&quot;log&quot;
	&quot;net/http&quot;
)

type Rss struct {
	XMLName xml.Name `xml:&quot;rss&quot;`
	Channel Channel  `xml:&quot;channel&quot;`
}

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

func main() {

	url := &quot;https://meetgor.com/rss.xml&quot;
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
	err = xml.Unmarshal(data, &amp;d)

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range d.Channel.Item {
		log.Println(item.Title)
	}
}
</code></pre>
<pre><code>$ go run main.go

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
</code></pre>
<p>So, here we have initialized the <code>Rss</code> struct as empty and then used the <a href="https://pkg.go.dev/encoding/xml#Unmarshal">Unmarshal</a> method in the <code>xml</code> package. The Unmarshal method will parse the data as per the type of either int, float, bool or string, any other type of data will be discarded as interface or struct. We can usually parse any valid type of data into <code>Unmarshal</code> method and it generally gives a proper expected outcome.</p>
<p>The Unmarshal method takes in the slice of byte and the second paramter as pointer to a struct or any variable that will store the parsed xml content from the slice of byte. The function just returns the error type, either <code>nil</code> in case of no errors, and returns the actual error obejct if there arise any type of error.</p>
<p>So we parse the <code>data</code> which is a slice of byte to the funciton and the reference to the <code>d</code> object which is a empty <code>Rss</code> object. This will get us the data in the <code>d</code> object. We can then iterate over the object as per the struct and use the perform operations like type casting or converting types, etc to get your required data back.</p>
<p>In the above example, we simply iterate over the <code>d.Channel.Item</code> which is a list of elements of tag <code>item</code> in the rss feed. Inside the for loop, we can access the object and simply print or perform any sort of operations. I have simply printed the list of articles with titles.</p>
<p>Links for the code available on the <a href="https://github.com/Mr-Destructive/100-days-of-golang/blob/main/scripts/files/read/config_files/xml/rss.go">100 days of golang</a> GitHub repository.</p>
<p>So, that's how we parse an XML feed in golang. Just plug and play if you have a similar type of structure of the Rss XML feed. Happy Coding :)</p>
