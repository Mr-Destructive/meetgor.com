{
  "title": "Golang: Interfaces",
  "status": "published",
  "slug": "golang-interfaces",
  "date": "2025-04-05T12:33:36Z"
}

<h2>Introduction</h2>
<p>In the 19th post of the series, we will be taking a look into interfaces in golang. Interfaces allow us to create function signatures common to different structs or types. So, we can allow multiple structs to have a common interface(method) that can have different implementations.</p>
<h2>What are Interfaces</h2>
<p>Interface as the name suggests is a way to create methods that are common to different structures or types but can have different implementations. It's an interface to define the method or function signatures but not the implementation. Let's take an example of <code>Laptop</code> and <code>Phone</code> having the functionality of wifi. We can connect to wifi more or the less in a similar way on both devices. The implementation behind the functionality might be different but they share the same operation. The WiFi can act as an interface for both devices to connect to the internet.</p>
<h2>Define an Interface</h2>
<p>To declare an interface in golang, we can use the <code>interface</code> keyword in golang. An interface is created with the type keyword, providing the name of the interface and defining the function declaration. Inside the interface which acts as a struct of general method signatures. The method signatures usually consist of the name of the function with its parameters if any and the return type of the function.</p>
<pre><code class="language-go">
package main

import &quot;fmt&quot;


type Player struct {
	name   string
	health int
}

type Mob struct {
	name     string
	health   int
	is_passive bool
}

type Creature interface {
	intro() string
	//attack() int
	//heal() int
}

func main() {
	player := Player{name: &quot;Steve&quot;}
	mob := Mob{name: &quot;Zombie&quot;}
	fmt.Println(player)
	fmt.Println(mob)
}
</code></pre>
<pre><code>go run main.go

{Steve 0}
{Zombie 0 false}
</code></pre>
<p>In this above example, we have created an interface called <code>Creature</code>. There are a few structs that we have to define like <code>Player</code> and <code>Mob</code>, these two methods have a few attributes like <code>name</code> as <code>string</code> and <code>health</code> as <code>int</code> which are common in both structs but the <code>Mob</code> struct has an additional attribute <code>is_passive</code> as a <code>boolean</code> value. The <code>Creature</code> is an interface that will define certain function signatures, here we have declared <code>intro</code>, <code>attack</code>, and <code>heal</code> as the methods bound to the Creature interface. This means, that any object which satisfies the Creature interface, can define the methods associated with the interface.</p>
<h2>Defining Interfaces</h2>
<p>Once we have declared the interface method signatures, we can move into defining the functionality of these methods depending on the struct. If we want a different working method for different types of struct objects passed we can define those for each type of struct that we want. Here, we have two types of structs namely <code>Creature</code> and <code>Mob</code>, based on the struct we can define the <code>intro</code> method for them individually.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

type Creature interface {
	intro() string
	attack(*int) int
}

type Player struct {
	name   string
	health int
}

type Mob struct {
	name     string
	health   int
	category bool
}

func (p Player) intro() string {
	fmt.Println(&quot;Player has spawned&quot;)
	return p.name
}

func (p Player) attack(m_health *int) int {
	fmt.Println(&quot;Player has attacked!&quot;)
	*m_health = *m_health - 50
	return *m_health
}

func (m Mob) intro() string {
	fmt.Printf(&quot;A wild %s has appeared!
&quot;, m.name)
	return m.name
}
func (m Mob) attack(p_health *int) int {
	fmt.Printf(&quot;%s has attacked you! -%d
&quot;, m.name, 30)
	*p_health = *p_health - 30
	return *p_health
}

func main() {
	player := Player{name: &quot;Steve&quot;, health: 100}
	mob := Mob{name: &quot;Zombie&quot;, health: 140}
	fmt.Println(player.intro())
	fmt.Println(mob.intro())
	fmt.Println(mob)
	fmt.Println(player)
	fmt.Println(player.attack(&amp;mob.health))
	fmt.Println(mob.attack(&amp;player.health))
	fmt.Println(mob)
	fmt.Println(player)
}
</code></pre>
<pre><code>go run main.go

Player has spawned
Steve
A wild Zombie has appeared!
Zombie
{Zombie 140 false}
{Steve 100}
Player has attacked!
90
Zombie has attacked you! -30
70
{Zombie 90 false}
{Steve 70}
</code></pre>
<p>As we can see, the method <code>intro()</code> is bound to both the struct depending on what struct signature is associated with the method. The method <code>intro</code> takes in the object struct associated as per the call and returns <code>string</code> as defined in the method signature.</p>
<p>The <code>attack</code> method in the <code>Creature</code> interface is also implemented separately for the two structs. For the <code>Player</code> method, we simply take in a pointer to an integer and return an <code>int</code>. The parameter is the pointer to the mob health, and it returns the modified health. We take in a pointer to the mob or player health so as to parse in the actual value and not the copy of the value. If we modify the value, we want to reflect those changes in the actual object. So that is how we can use interfaces to construct dynamic operations on objects as well as different types of structs.</p>
<p>We have seen a simple example of how to declare and define interfaces for given type structs. Also, we can pass by value as well as by pointers so as to define the behavior of the method whether to dynamically modify or change the values of the object associated with it.</p>
<h2>Examples of Interfaces</h2>
<p>There are quite some use cases of interfaces, in object-oriented programming, the above example fits the polymorphism feature quite well. The ability to reuse certain method signatures and define the functions as per requirement brings flexibility to the code structure.</p>
<p>We will see a few examples for understanding how we can use interfaces in various ways.</p>
<h3>Type Switch Interface</h3>
<p>We can use an empty interface to check for the type of variable we have parsed. Using this empty interface we can create a kind of dynamic parameter to a function.</p>
<pre><code class="language-go">package main

import (
	&quot;fmt&quot;
	&quot;strconv&quot;
)

func parse_int(n interface{}) int {
	switch n.(type) {
	case int:
		return (n).(int) * (n).(int)
	case string:
		s, _ := strconv.Atoi(n.(string))
		return s
	case float64:
		return int(n.(float64))
	default:
		return n.(int)
	}
}

func main() {
	num := parse_int(4)
	fmt.Println(num)
	num = parse_int(&quot;4&quot;)
	fmt.Println(num)
	num = parse_int(4.1243)
	fmt.Println(num)
}

</code></pre>
<pre><code>go run main.go

16
4
4
</code></pre>
<p>Here, we can see we have an interface as a parameter to the function <code>parse_int</code>, the return type is <code>int</code>, so the incoming parameter can be any valid type. But if we don't convert the given type into an appropriate int, it will result in an error as we are returning the int value of the parsed parameter. We are taking the parameter as <code>interface{}</code> which is an empty interface, this will contain the parameter parsed as an interface type. That's why we need to convert the interface object into an int or the parsed type of the interface.</p>
<h3>Interface Slice</h3>
<p>We can even create a slice of interfaces, which means we can initialize or group together various objects of different structs in a single slice. This might be helpful for calling functions associated with different objects via the interface very easily and in a much cleaner way.</p>
<pre><code class="language-go">package main

import &quot;fmt&quot;

type Creature interface {
	intro() string
}

type Player struct {
	name   string
	health int
}

type Mob struct {
	name     string
	health   int
	category bool
}

func (p Player) intro() string {
	fmt.Println(&quot;Player has spawned&quot;)
	return p.name
}

func (m Mob) intro() string {
	var name string
	if m.name != &quot;&quot; {
		name = m.name
	} else {
		name = &quot;Mob&quot;
	}
	fmt.Printf(&quot;A wild %s has appeared!
&quot;, name)
	return m.name
}


func main() {
	entity := []Creature{Player{}, Mob{}, Mob{}, Player{}}

	for _, obj := range entity {
		fmt.Println(obj.intro())
	}
}
</code></pre>
<pre><code>Player has spawned
A wild Zombie has appeared!
A wild Zombie has appeared!
Player has spawned
</code></pre>
<p>In the above example, we can see that the entity variable is created as a slice of interfaces <code>Creature</code>, i.e. various objects associated with the <code>Creature</code> interface can be contained in a single slice. There are 2 instances of Player and Mob each in the slice. We can further iterate over the slice as a range-based loop and thereby the functions associated with the interfaces can be called. Here, we have called the <code>intro</code> function.</p>
<p>So, there are a lot of things that can be done with interfaces, we can create multiple interfaces for a single object struct and nest interfaces. Based on the use case of the program, interfaces can be used to reduce the boilerplate code as well as improve the readability of the code.</p>
<p>That's it from this part. Reference for all the code examples and commands can be found in the <a href="https://github.com/mr-destructive/100-days-of-golang/tree/main/scripts/interfaces/main.go">100 days of Golang</a> GitHub repository.</p>
<h2>Conclusion</h2>
<p>From this part of the series, we were able to understand the basics of interfaces using a few examples. We explored how interfaces can be used to bring in polymorphism in golang, also we can improve the readability of the code. The boilerplate code can be considerably reduced by using interfaces when dealing with structs and types. Hopefully, you found this post helpful and understood even the basics of interfaces in golang. Thank you for reading, if you have any queries, questions, or feedback, you can ping me on my social handles or in the comments. Happy Coding :)</p>
