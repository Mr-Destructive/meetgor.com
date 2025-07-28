{
  "title": "Golang: Structs",
  "status": "published",
  "slug": "golang-structs",
  "date": "2025-04-05T12:33:42Z"
}

<h2>Introduction</h2>
<p>Moving on to the 9th part of the series, we will be understanding structs in golang. Structs are an important aspect of programming in Golang, they provide a way to define custom types and add functionality to them. We will be understanding the basics of operating on structs like declaration, initialization and adding functional logic into those structs.</p>
<h2>Structs in Golang</h2>
<p>Structs or Structures in Golang are the sequences or collections of built-in data types as a single type interface. Just like we have int, string, float, and complex, we can define our own data types in golang. They can consist of built-in data types as mentioned and also certain functions or methods which can be used to operate on them. Using structs we can create custom data types that can meet the specific requirements of our problem. We can define structs and later inside functions we can create instances of those structures.</p>
<p>Structures are like a template or blueprint representation of data. It doesn't hold the actual data in memory, it is just used to construct an object of that type. After defining a struct, we can create instances or objects of those structs. These instances actually hold data in memory in the run time, so we basically deal with objects in the actual program. We'll see certain concepts of creating instances, declaring and defining structs, accessing data from instances and so on in the following section of the article.</p>
<pre><code>Struct / Class 

Template / Structure for creating custom data types 

- Properties  (variables and constants defined inside a structure)
- Methods     (functions that are bound to a struct)

</code></pre>
<h2>Declaring Struct</h2>
<p>We can declare structs by using the keyword <code>type</code> followed by the name of the struct, after tha name, the <code>struct</code> keyword itself, and finally sets of parenthesis <code>{}</code>. Inside the parenthesis, we define the structure i.e. which type of data is to be stored and the name of those respective variables.</p>
<pre><code class="language-go">type Article struct {
    title string
    is_published bool
    words int
}
</code></pre>
<p>We have declared a struct or a custom data-type or a class(not really) in golang with the name <code>Article</code> that has few associated properties/variables inside of it. We have <code>title</code> as a string, <code>is_published</code> as a boolean, and <code>words</code> as an integer value. This constructs a simple type of golang which has a defined structure. We can further use this Article struct as a data type in the main function or any appropriate scope for actually assigning the structure memory at runtime.</p>
<h3>Struct Naming Convention</h3>
<p>There are a few things that we need to understand and make a note of, especially the naming convention.</p>
<ul>
<li>The struct name should be capitalized if you want to make it publicly accessible.</li>
<li>The variable/properties names i.e. <code>title</code>, <code>is_published</code>, and <code>words</code> should be also capitalized if you want to make them accessible from the struct instance.</li>
</ul>
<p>This might not be important right now but it is worth knowing for later use cases. Let's say we want to use a struct from other files or modules, for that the name of the struct in the file/script where the struct is defined should have the <code>Capitalized</code> convention. If you have a simple and single file script/program, you can keep it <code>lowercased</code> or <code>camelCased</code>.</p>
<p>Leaving that aside, for now, we will try to focus on the essence of the structs in golang.</p>
<h2>Creating Instances/Objects of Structs</h2>
<p>Now, after defining the struct we need to create instances or objects of them. This can be done in several ways like using Struct literal, Manual assignment, and using the new function. We'll look into each of them in this section.</p>
<h3>Using struct literal</h3>
<p>The most simplest and straightforward way to initialize a struct is to use the struct literal just like we did with Maps, Slices, and Arrays. We basically parse the values of the respective fields in the struct.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

type article struct {
	title        string
	is_published bool
	words        int
}

func main() {
	golang := article{&quot;Golang Intro&quot;, true, 2000}
	fmt.Println(golang)
}
</code></pre>
<pre><code>$ go run struct.go
{Golang Intro true 2000}
</code></pre>
<p>We have created the object or instance of the struct <code>Article</code> using the shorthand notation or the walrus <code>:=</code> operator. Inside the <code>{}</code> braces, we can assign values but those values need to be in the same order as defined in the struct definition, else it gives a compilation error of <code>type mismatch</code>. So, here we have assigned the value <code>title</code>, <code>is_published</code>, and <code>word</code> as <code>Golang Intro</code>, <code>true</code>, and <code>2000</code> respective in that order.</p>
<h3>Using Key-value pairs</h3>
<p>We can also use the <code>key-value</code> notation for assigning values in the instance. With the previous method, we need to specify and thus initialize all the properties at once, but using this method we have a bit more flexibility.</p>
<pre><code class="language-go">vim := Article{title: &quot;Vim: Keymapping&quot;, is_published: false}
fmt.Println(vim)
</code></pre>
<pre><code>$ go run struct.go
{Vim: Keymapping false 0}
</code></pre>
<p>Here, we have provided the key i.e. the variable name inside the struct, and then provided the value to it separated by a colon <code>:</code>. Using this way of initializing instances of struct we have better control and flexibility in providing a default value for that object. In the example above, we didn't initialize the property <code>words</code> but it already initialized to <code>0</code> since the object is created hence the memory allocation is completed, and thereby it needs to have a default value.</p>
<h3>Using the new function</h3>
<p>We can use the <a href="https://pkg.go.dev/builtin#new">new</a> function to create a new instance of a struct. Though we can't provide an initial value, using the new function all the properties are initialized with their respective default values. Further, if we want to modify the values, we can access each property (variables in struct) using the <code>dot operator</code> and assign the desired values.</p>
<pre><code class="language-go">django := *new(Article)
fmt.Println(django)
</code></pre>
<pre><code>$ go run struct.go
{ false 0}
</code></pre>
<p>We have used the new function to allocate memory for an instance of struct with the provided name. This function basically allocates all the properties of a default value and returns a pointer to that memory address. If we store the result of the new function in a variable object, we would get a pointer but we need the object itself, so we use <code>*</code> before the new function so as to de-reference the memory address from the pointer.</p>
<p>So, we have stored the default values in the newly created object of Article structure in <code>django</code>, this gives the default values like an empty string <code>&quot;&quot;</code>, default boolean value <code>false</code> and default integer value <code>0</code>. If we don't dereference the pointer and use it like <code>djagno := new(Article)</code>, thereby we get a pointer in that variable as <code>&amp;{ false 0}</code>. Hence we use <code>*</code> before the new keyword.</p>
<h4>Accessing/Assigning values to properties</h4>
<p>We can now change the values of the properties in the object of the struct using the dot operator. We basically use the instance object name followed by a <code>.</code> and the property name to set its value.</p>
<pre><code class="language-go">django := *new(Article)
fmt.Println(django)

django.title = &quot;Django View and URLs&quot;
django.words = 3500
django.is_published = true
fmt.Println(django)
</code></pre>
<pre><code>$ go run struct.go
{ false 0}
{Django View and URLs true 3500}
</code></pre>
<p>So, here we have used the object name which is <code>django</code>, and access any property by name with the <code>dot operator</code>, thereby we set the value as per the requirement. Note, we have not used the <code>:=</code> operator as the properties have already been initialized, we simply need to modify the default value.</p>
<h2>Creating Functions associated to Structs</h2>
<p>We can now move into creating functions in the struct, by adding functions/methods in structs we can incorporate a lot of functionality into the structure of our data type. For instance, we can set the value of a string as <code>&quot;Empty&quot;</code> or <code>&quot;NA&quot;</code> beforehand rather than empty string <code>&quot;&quot;</code>.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

type Mail struct {
	sender     string
	subject    string
	sent       bool
	word_count int
}

func (m Mail) check_spam() {
	if m.subject == &quot;&quot; {
		fmt.Println(&quot;Spam!&quot;)
	} else {
		fmt.Println(&quot;Safe!&quot;)
	}
}

func main() {
	mail_one := *new(Mail)
	fmt.Printf(&quot;Mail one: &quot;)
	mail_one.check_spam()

	mail_two := Mail{&quot;xyz@xyz.com&quot;, &quot;Golang Structs&quot;, true, 100}
	fmt.Printf(&quot;Mail two: &quot;)
	mail_two.check_spam()
}
</code></pre>
<pre><code>$ go run methods.go
Mail one: Spam!
Mail two: Safe!
</code></pre>
<p>We define a function associated with a struct by providing the <code>struct-name</code> and a parameter name which can be just used inside of the function. Here, we have used <code>(m Mail)</code> so as to reference the object of the struct provided to it. This basically binds the function to the struct and hence it becomes a method of that struct.</p>
<p>Further, we can access the properties from the struct by their name using the dot separator. We are just checking whether the subject property in the instance is empty or not and simply printing text to the console. We are accessing the function and calling it with the syntax as <code>instance_name.function_name()</code>, here the function name is <code>check_spam</code> and the object name is <code>mail_one</code> for the first instance. Thereby we have called the function which is bounded to the instance of the struct. As we have accessed the function after the instance name the binding of the function i.e. the statements <code>(m Mail)</code> has taken the current instance and parsed it as the instance of the struct. Hence we are able to access the current instance's properties within the function/method.</p>
<h4>Adding a return statement</h4>
<p>By simply providing the return type and return statement with value, we can create functions of specific return types.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

type Mail struct {
	sender     string
	subject    string
	sent       bool
	word_count int
}

func (m Mail) check_spam() bool {
	if m.subject == &quot;&quot; {
		return true
	} else {
		return false
	}
}

func (m Mail) print_spam(spam bool) {
	if spam {
		fmt.Println(&quot;Spam!!&quot;)
	} else {
		fmt.Println(&quot;Safe!!&quot;)
	}
}

func main() {
	mail_one := *new(Mail)
	fmt.Printf(&quot;Mail one: &quot;)
	is_mail_1_spam := mail_one.check_spam()
	mail_one.print_spam(is_mail_1_spam)

	mail_two := Mail{&quot;xyz@xyz.com&quot;, &quot;Golang Structs&quot;, true, 100}
	fmt.Printf(&quot;Mail two: &quot;)
	is_mail_2_spam := mail_two.check_spam()
	mail_two.print_spam(is_mail_2_spam)
}
</code></pre>
<pre><code>$ go run methods.go
Mail one: Spam!!
Mail two: Safe!!
</code></pre>
<p>We have modified the <code>check_spam</code> function which returns a boolean value. If the subject is empty it returns true else it returns false. Also, we have added a function <code>print_spam</code> function which takes in a parameter as a boolean value and prints text according to the value. This is how we work with functions in structs. We have parsed the return value of the <code>check_spam</code> function as a parameter to the <code>print_spam</code> function.</p>
<h3>Constructor in Structs</h3>
<p>Constructors are special methods that are invoked when the instance of a struct is created i.e. the properties are assigned an initial value or default value. In this way, we can perform basic operations which we need to perform after the instantiation of the struct.</p>
<p>Golang does not have built-in constructors, but it is quite easy to create one. We simply need to create a function with an appropriate name(don't clash it with the struct name!!), by providing all the parameters that are in the struct so as to initialize them, and finally the return value as a reference to the struct instance.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

type Repository struct {
	name       string
	file_count int
}

func New_Repository(name string, file_count int) *Repository {
	file_count++
	name = &quot;Test&quot;
	return &amp;Repository{name, file_count}
}

func main() {
	blog := *New_Repository(&quot;&quot;, 0)
	fmt.Println(blog)
}

</code></pre>
<pre><code>$ go run constructor.go
{Test 1}
</code></pre>
<p>We have created a function that is technically acting like a constructor as it sets a default value to the properties in the structure. We have struct <code>Repository</code> containing <code>name</code> as a string and <code>file_count</code> as an integer. We created a Constructor function named <code>New_Repository</code> that basically takes in the properties in the struct, remember they haven't been initialized yet as we are writing the constructor for the very purpose. We have to parse the parameters with the initial value and let it modify once we have created the instance.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>So, from this part of the series, we are able to understand the basics of structs in golang. We covered declaration, definition, and adding methods in a struct. This gives a glimpse of Object-Oriented Programming in Golang. Thank you for reading. If you have any questions or feedback, please let me know in the comments or on social handles. Happy Coding :)</p>
