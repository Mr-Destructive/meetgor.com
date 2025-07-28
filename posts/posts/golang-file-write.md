{
  "title": "Golang: File Write",
  "status": "published",
  "slug": "golang-file-write",
  "date": "2025-04-05T12:33:34Z"
}

<h2>Introduction</h2>
<p>In the 24th post of the series, we will be taking a look at how we can perform write operations to a file using golang. We will be using the <code>os</code> package in most operations along with <code>bufio</code> text manipulations. We will be performing write operations like appending, deleting, and replacing a file using golang. We will be heavily leveraging standard library packages like <code>os</code>, <code>bufio</code>, <code>bytes</code> and <code>fmt</code>. We will also be looking into overwriting and string formatting to a file.</p>
<h2>Write to a File</h2>
<p>The first part of this section is the basic write operation to a file, we assume we are writing to a fresh file and overriding the contents of the existing file. The next section will cover the appending of content to the file and so on. In this example, we will see how we perform basic writing operations to write a string, a slice of string to a file.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;os&quot;
)

func HandleError(err){
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
	str := &quot;Hi, I am a gopher!
&quot;
	f, err := os.OpenFile(&quot;abc.txt&quot;, os.O_WRONLY, 0660)
    // f, err := os.Create(&quot;abc.txt&quot;)
	HandleError(err)
	_, err = f.Write([]byte(str))
	HandleError(err)
	defer f.Close()
}
</code></pre>
<pre><code class="language-bash">$ cat abc.txt

$ go run main.go

$ cat abc.txt
Hi, I am a gopher!
</code></pre>
<p>So, we have used a simple golang script to write to a file that exists/has already been created. If you don't want errors while having to write on a file that does not exist, use the <a href="https://pkg.go.dev/os#Create">Create</a> method instead which is similar to the <code>Open</code> method but creates a file if it doesn't exist. We use the <a href="https://pkg.go.dev/os#File.Write">Write</a> method to overwrite contents to the file, it takes in a parameter as a slice of byte, so we typecast the string <code>str</code> into <code>[]byte</code> using the <code>[]byte(str)</code> syntax. Thereby we write the contents of the string into the file. We use the defer keyword for closing the file at the end of the script or the end of the main function scope.</p>
<h3>Write slice of strings to file</h3>
<p>We can even write slice of string to file using a for loop and appending each string with a new line character.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;os&quot;
)

func HandleError(err){
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
	f, err := os.Create(&quot;abc.txt&quot;)
	//f, err := os.Open(&quot;abc.txt&quot;, os.O_WRONLY, 0660)
	langs := []string{&quot;golang&quot;, &quot;python&quot;, &quot;rust&quot;, &quot;javascript&quot;, &quot;ruby&quot;}
	for _, lang := range langs {
		_, err := f.WriteString(lang + &quot;
&quot;)
		HandleError(err)
	}
	defer f.Close()
}
</code></pre>
<pre><code class="language-plaintext">$ cat abc.txt

$ go run main.go

$ cat abc.txt
golang
python
rust
javascript
ruby
</code></pre>
<p>We have used the <a href="https://pkg.go.dev/os#File.WriteString">WriteString</a> method which will take in a string as a parameter instead of a slice of bytes. So, we don't have to type cast into slice of bytes. So, as we can see we have written the string slice into a file.</p>
<h3>Over Write</h3>
<p>The minimal code to write to a file is the <a href="https://pkg.go.dev/os#WriteFile">WriteFile</a> function in the <a href="https://pkg.go.dev/os">os</a> package, it overrides the content of the file with the provided slice of bytes, the name of the file, and the necessary permission to write. The funciton additionally creates a file if it does not exist, which is one less reason for the error. Though it returns an error object, the error might be created due to not right permissions to write to the file, encoding issues, etc.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;os&quot;
)

func main() {
	data := []byte{115, 111, 109, 101, 65}
	err := os.WriteFile(&quot;test.txt&quot;, data, 0660)
    log.Println(err)

	s := &quot;Hello&quot;
	err = os.WriteFile(&quot;test.txt&quot;, []byte(s), 0660)
    log.Println(err)
}
</code></pre>
<pre><code class="language-plaintext">$ go run main.go
2022/12/17 19:24:13 &lt;nil&gt;
2022/12/17 19:24:13 &lt;nil&gt;

$ cat test.txt
Hello
</code></pre>
<p>So, we have used the <code>WriteFile</code> method two times in the script, it first takes in a slice of bytes as it is defined as <code>data</code> which corresponds to <code>115 -&gt; s</code>, <code>111 -&gt; o</code>, <code>65 -&gt; A</code>, ASCII mapped to strings. The slice of bytes can be taken as a string like <code>someA</code> as the literal value of the underlying slice of the byte. So, we take that slice of byte and parse it to the second parameter of the WriteFile function. The first parameter is a string path of the file we want to write the contents to, the third parameter is the file permission. We have set it as <code>0660</code> indicating read(4) + write(2) to the group and the user and no permission to the other users. The function will return an error if any, or else it simply overwrites the data in the file.</p>
<p>In this case, we have called the <code>WriteFile</code> method with string <code>s</code> type cast to slice of bytes at the end of the script so we see the file has contents as <code>Hello</code> instead of <code>someA</code>. If we reverse the action, we don't see the <code>Hello</code> string in the file.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;os&quot;
)

func main() {
	s := &quot;Hello&quot;
	err := os.WriteFile(&quot;test.txt&quot;, []byte(s), 0660)
    log.Println(err)

	data := []byte{115, 111, 109, 101, 65}
	err = os.WriteFile(&quot;test.txt&quot;, data, 0660)
    log.Println(err)
}
</code></pre>
<pre><code class="language-plaintext">$ go run main.go
2022/12/17 19:24:13 &lt;nil&gt;
2022/12/17 19:24:13 &lt;nil&gt;

$ cat test.txt
someA
</code></pre>
<p>As we can see the <code>Hello</code> has been overwritten by <code>someA</code>.</p>
<h3>Write formatted string</h3>
<p>We can even use fmt to write formatted strings to a file. Just like we can take inputs with <code>Scanf</code>, we can use <a href="https://pkg.go.dev/fmt#Fprint">Fprint</a> and other similar functions like <a href="https://pkg.go.dev/fmt#Fprintf">Fprintf</a>, and <a href="https://pkg.go.dev/fmt#Fprintln">Fprintln</a> functions to print/add contents to the file.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;log&quot;
	&quot;os&quot;
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	f, err := os.Create(&quot;temp.txt&quot;)
	HandleErr(err)
	name, lang, exp := &quot;John&quot;, &quot;go&quot;, 2
	_, err = fmt.Fprint(f, &quot;Hi, I am &quot;, name, &quot;
&quot;)
	HandleErr(err)
	_, err = fmt.Fprintf(f, &quot;Language of choice: %s.
&quot;, lang)
	HandleErr(err)
	_, err = fmt.Fprintln(f, &quot;Having&quot;, exp, &quot;years of experience.&quot;)
	HandleErr(err)
	defer f.Close()
}
</code></pre>
<pre><code class="language-plaintext">$ cat temp.txt
cat: temp.txt: No such file or directory

$ go run format.go

$ cat test.txt
Hi, I am John.
Language of choice: go.
Having 2 years of experience.
</code></pre>
<p>So, we can see that we have used all three methods having their own use cases, we can use <code>Fprint</code> for simple strings, <code>Fprintf</code> for formatting the block of a string with multiple placeholders, and the <code>Fprintln</code> which works simply like <code>Fprint</code> but it adds a new line itself, we don't need to specify it explicitly.</p>
<h3>Append</h3>
<p>If we want to append text to a file, we can use the <a href="https://pkg.go.dev/os#OpenFile">OpenFile</a> function and provide a few parameters to append the contents instead of overwriting.</p>
<p>Here, we have two steps, open the file and then write the contents in the file. So while opening the file, we provide a few options as parameters to make the fine-tuned system call like only open for read, write or append modes. These options are defined as constant int values in the <a href="https://pkg.go.dev/os#pkg-constants">os package</a>.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;os&quot;
)

func HandleError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
	s := &quot;Hello&quot;
	err := os.WriteFile(&quot;test.txt&quot;, []byte(s), 0660)
    HandleError(err)

	s = &quot;World&quot;
	f, err := os.OpenFile(&quot;test.txt&quot;, os.O_APPEND|os.O_WRONLY, 0660)
	HandleError(err)
	_, err = f.WriteString(s)
	HandleError(err)
	defer f.Close()
}
</code></pre>
<pre><code class="language-plaintext">$ go run main.go

$ cat test.txt
HelloWorld
</code></pre>
<p>So, from the above example, we are able to append text into a file, we have first added the <code>Hello</code> world string into the file using the <code>WriteFile</code> method to indicate we overwrite the previous contents of the file. We then use the <a href="https://pkg.go.dev/os#OpenFile">OpenFile</a> method to open a file provided in the first parameter as a string path. The second parameter is the options to be passed for performing operations on the opened file, we always have to use them <code>defer</code> to close the file or other resource-locking operations.</p>
<p>We have specified the <code>os.O_WRONLY</code> and the <code>os.O_APPEND</code> options indicating we want to write to the file while the file is open and specifically append to the file. So this is fine-tuning the opened file operation. We can use the ReadFile or WriteFile operation which is just used for simple read and write operations respectively.</p>
<p>We use the <a href="https://pkg.go.dev/os#File.WriteString">WriteString</a> method, but we can even use the <a href="https://pkg.go.dev/os#File.Write">Write</a> method to write a slice of byte instead. This is just used for exploring the different options in the file types of the os package.</p>
<h3>Append at a specific line</h3>
<p>We can also add content to a specific line or a portion of the file. There are no direct functions in golang to do the same, we will have to do some manual fine-tuning of file operations to append a particular text at a specific line.</p>
<pre><code class="language-go">package main

import (
    &quot;bufio&quot;
    &quot;bytes&quot;
    &quot;log&quot;
    &quot;os&quot;
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
	f, err = os.OpenFile(&quot;test.txt&quot;, os.O_RDWR, 0660)
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
		if i &lt; line_till {
			bytes_till += bytes.Count([]byte(line), []byte{})
			if i &gt; 0 {
				lines_till += &quot;
&quot;
			}
			lines_till += line
		} else {
			lines_after += &quot;
&quot; + line
		}
		i += 1
	}
	HandleError(m.Err())
	insert_text := lines_till + &quot;
Inserted content&quot;
	insert_text_bytes := bytes.Count([]byte(insert_text), []byte{})
	lines_after_bytes := bytes.Count([]byte(lines_after), []byte{})

	err = os.WriteFile(&quot;test.txt&quot;, []byte(insert_text), 0660)
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
</code></pre>
<pre><code class="language-plaintext">$ cat test.txt 
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
</code></pre>
<p>We have inserted <code>Inserted content</code> after the second line because the <code>line_till</code> the variable is set to <code>2</code></p>
<p>So, in the above example, we have used a bunch of packages to append a string or any form of text to a particular line. We first read the contents of the file, by using the <code>OpenFile</code> method which will open the file with certain permissions. We need to close the file at the end of the script so we simply use the <code>defer</code> keyword before the <code>f.Close()</code> method call. We then start to scan the file buffer, by creating a <code>Scanner</code> object with the <code>NewScanner</code> method. Then, with the scanner object of the file content, we then can use the <code>Scan()</code> method to scan the file contents line by line. By converting the content at each line from a slice of bytes to a string using the <code>Text</code>, we append to a string <code>line</code>, this will be used for keeping the count of bytes for appending text before the newly inserted text.</p>
<p>The <code>line_till</code> variable is used for the line number from which we want to append to the text after.</p>
<p>We count the bytes for the current line and add it to the <code>bytes_till</code> variable indicating the number of bytes there before appending content. We have a simple if-else check for the first line that is for appending a new line of characters. We append the lines into a single string <code>lines_till</code>. The string <code>insert_text</code> is created by appending all the lines before the line number <code>line_till</code> with the actual content to be inserted. We calculate the number of bytes using the <a href="https://pkg.go.dev/bytes#Count">Count</a> method in the bytes package. The separator is kept blank. The <code>lines_after</code> is also been created as a single string of lines after the line number in the file.</p>
<p>We add the <code>insert_text</code> (lines before + inserted text) into the file using the <code>WriteFile</code> which will override the contents of the file. Then we append the <code>lines_after</code> string as a slice of bytes to the <code>insert_text_bytes + lines_after_bytes</code> so we get the byte number position to append the <code>lines_after</code> string.</p>
<p>In short, we basically overwrite the file by creating two strings (slice of bytes) one which has the lines before the line number with the text to be inserted and the second string has all the lines after the line number.</p>
<h2>Replace text in a file</h2>
<p>Using the <a href="https://pkg.go.dev/bytes#Replace">bytes.Replace</a> method, we can first read all the bytes and replace the old with the new text, and store them as a slice of bytes. We then write these slices of bytes to the file again, so we first read the contents into slices of bytes, replace the content of the byte and then overwrite the contents with the slice of bytes. It's quite simple.</p>
<pre><code class="language-go">package main

import (
	&quot;bytes&quot;
	&quot;log&quot;
	&quot;os&quot;
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	filename := &quot;test.txt&quot;
	file, err := os.ReadFile(filename)
	HandleError(err)
	old_text := &quot;Hello
World&quot;
	new_text := &quot;Bye&quot;
	new_content := bytes.Replace(file, []byte(old_text), []byte(new_text), -1)
	err = os.WriteFile(filename, new_content, 0660)
	HandleError(err)
}
</code></pre>
<pre><code class="language-plaintext">$ cat test.txt
Hi
Hello
World
Gophers

$ go run main.go


$ cat test.txt
Hi
Bye
Gophers
</code></pre>
<p>As we can see we have replaced <code>Hello World</code> it with <code>Bye</code>. So, the [Replace] method in the bytes package takes in the parameters like the slice of bytes which should be the actual contents of the file, the old text to be replaced with again as a slice of bytes, and the new text to replace also as the slice of bytes, the final parameter is the number of replacements to be made. Here <code>-1</code> indicates there are no limits on how many replacements can be done, it can be <code>1</code>, <code>2</code> for replacing the first n occurrence of the old text, depending on how many times you want to replace the content in the file.</p>
<h2>Delete Text from a File</h2>
<p>We can use the <a href="https://pkg.go.dev/os#File.Truncate">os.Truncate</a> method to delete the contents of the file. The <code>Truncate</code> method takes in the parameters like the file path string and the size of the file to truncate or set to. If we set the second parameter to <code>0</code>, the file size will be zero and all the contents will be deleted or removed.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;os&quot;
)

func main() {
    /* test.txt
    Hi
    Hello
    World
    Gophers
    */
	err := os.Truncate(&quot;test.txt&quot;, 0)
	if err != nil {
		log.Fatal(err)
	}
    /* test.txt is empty
    */
}
</code></pre>
<pre><code class="language-plaintext">$ cat test.txt
Hi
Hello
World
Gophers

$ go run delete.go

$ cat test.txt
</code></pre>
<p>As we can see that the contents of the file are emptied if we set the second parameter(size) of the <code>Truncate</code> method as 0.</p>
<p>We can also set the value of the size as the number of bytes to keep, so instead of <code>0</code> we can set it to <code>n</code> positive integer to only save the first <code>n</code> bytes in the file.</p>
<pre><code class="language-go">package main

import (
	&quot;log&quot;
	&quot;os&quot;
)

func main() {
    /* test.txt
    Hi
    Hello
    World
    Gophers
    */
	err := os.Truncate(&quot;test.txt&quot;, 6)
	if err != nil {
		log.Fatal(err)
	}
    /* test.txt 
    Hi
    Hel
    */
}
</code></pre>
<pre><code class="language-plaintext">$ cat test.txt
Hi
Hello
World
Gophers

$ go run delete.go

$ cat test.txt
Hi
Hel
</code></pre>
<p>So, in the above example, if we set the size parameter to the Truncate method as <code>6</code>, it will keep the size of the file to 6 bytes. So we only see <code>Hi Hel</code>, the new line is a single byte. The rest of the content is deleted. This is how we previously deleted all the bytes from the file by setting the size to <code>0</code>.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/files/write/">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So, from this section of the series, we were able to perform write operations on a file using golang. We used the packages from the standard library and performed operations like write, append, overwrite, delete and replace to a simple text file, but it could have been any file format.</p>
<p>Thank you for reading. If you have any queries, questions, or feedback, you can let me know in the discussion below or on my social handles. Happy Coding :)</p>
