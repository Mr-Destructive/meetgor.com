{
  "title": "Object-Oriented Programming in C++: Classes and Objects",
  "status": "published",
  "slug": "oop-cpp-p1",
  "date": "2025-04-05T12:33:53Z"
}

<h2>Introduction</h2>
<p>We all know that C++ is famously known for Object-Oriented Programming, but what is Object-Oriented Programming? Well, this is the article, where we will explore the basics of Object-Oriented Programming, its building blocks i.e Classes and Objects, and get familiar with its basic semantics in C++. Let's get into it straight away!</p>
<h2>What is Object-Oriented Programming?</h2>
<p>Object-Oriented Programming(OOP) is a way to structure and design a program or an application. Object-Oriented Programming deals with creating classes and objects and it's related terminologies. OOP helps in following the principles like DRY(Don't Repeat Yourself) and KISS(Keep it Simple Stupid). Using OOP certain applications and problems can be solved in a simple and efficient way. It helps in better understanding and readability of the code.</p>
<blockquote>
<p>Virus is a class and COIVD-19 is an object of the Virus class XD</p>
</blockquote>
<p>In OOP we create a collection of the data and functionalities by organizing in a structure called <code>Class</code>. We then instantiate or create the actual data by creating an <code>object</code> of the particular class. Once we have created a class, we can simply create as many objects as we need to. Objects are basically the storage of the actual data in memory. Let's explore them in the next few sections.</p>
<p>NOTE: There are a lot of terminologies in OOP, just remember the basics of those terms and not the different names of those terms.</p>
<h2>What are Classes</h2>
<p>Classes are the structure or template of the data and its associated behavior. The data or the variables in the classes are called <code>properties</code> or <code>attributes</code>(also <code>data members</code>). We also have functions that define the behavior of the properties or the data present in the class and are called <code>methods</code>.</p>
<p>A Class in C++ should be defined outside of the main function. In C++, <code>class</code> is a reserved word and hence it is used to declare and define a class. It has a general structure like:</p>
<pre><code class="language-cpp">class Name
{
    // properties and methods
};
</code></pre>
<p>It's a convention to use the initial letter of a Class Name Uppercased. We can define class methods outside the class body as well, we will see it in the next few parts of the series.</p>
<p>Let's create a basic class in C++,</p>
<pre><code class="language-cpp">class Animal
{
	public:
		int legs;
		string name;
	void print()
    {
		cout&lt;&lt;name&lt;&lt;&quot; has &quot;&lt;&lt;legs&lt;&lt;&quot; legs.
&quot;;
	}
};

</code></pre>
<p>Here in the above code, we have a class called <code>Animal</code> and it has 2 properties called <code>legs</code> which is an int and <code>name</code> a string. We also have a method called <code>print</code> which displays the current object's name and the leg count. We will see what is the current object in the next section.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631797771422/6jCNpCYz_q.png" alt="oopcpp1-class.png"></p>
<p>We use those properties or the variables defined in the class in our application either by accessing them through the objects or in the methods of that class. The variables or properties in classes are not assigned to any memory address i.e. they are just kind of a blueprint or placeholder tags to match for the data.</p>
<h3>Access Specifier</h3>
<p>We have written <code>public</code> before the definition of the properties of the class, it is an access modifier. Don't worry we will see it in detail in the next few parts of the series, just for time being understand that we can change whether we want the main function(or any function globally) to access the class' properties or not.</p>
<p>We are saying <code>public</code> which means we can access the properties of this class anywhere outside the class. This is not a good practice but just for simplicity of understanding, we are using public. <strong>If you do not specify anything like <code>public</code>, it is by default <code>private</code>, which means the properties of the class are not accessed anywhere outside the class</strong>. Let keep it simple right now.</p>
<p>We also have other concepts like defining the methods outside/inside the class, header files, constructors, destructors, and many others related to a class definition, we will cover it in the next few parts.</p>
<h2>What are Objects</h2>
<p>Now, we have created a class but where is the data actually? It will be in <code>objects</code>, We actually assign the data or properties to a memory address by creating the objects of that particular class.</p>
<p>We can create objects of a particular class in C++ by writing the name of the class that this object will belong to and the name of the object,<code>classname objectname;</code>. This will create or assign memory to the properties of the class to the object.</p>
<p>After the object has been created, we can assign the value to the properties of the class in the object. We can access the properties of the class by referencing the name of the object with <code>.</code> and the name of the property or variable of the class, <code>objectname.propertyname</code>, we can assign the value simply by using the assignment operator <code>=</code> and give the properties the value, input from the user or however you like it.</p>
<pre><code class="language-cpp">#include&lt;iostream&gt;
using namespace std;

class Animal
{
	public:
		int legs;
		string name;
	void print()
    {
		cout&lt;&lt;name&lt;&lt;&quot; has &quot;&lt;&lt;legs&lt;&lt;&quot; legs.
&quot;;
	}
};

int main(){

	Animal dog;            // create an object of class Animal

	dog.name=&quot;Floyd&quot;;     // assign the class property to actual data in memory
	dog.legs=4;                 
	dog.print();          // call a method(function) associated to object's class
	
	return 0;
}

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631795000896/kUvcfVU7Y.png" alt="image.png"></p>
<p>If Animal is the class, then the dog here is the object. As simple as you can think. We are technically <strong>instantiating an object</strong> when we say <code>Animal dog</code>.  We assign the variables the memory as defined in the class. We give some value to the properties using the <code>.</code>(dot syntax), after the class name.</p>
<p>So <code>dog.legs = 4;</code> will assign the value of 4 in the memory address of the object's property(dog is the class and legs is the property).</p>
<p>Similarly, we can call the function associated with the class as we do for properties but we use <code>()</code> to indicate to calling the function.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631799152681/UhNDYdEXK.png" alt="oopcpp1-obj.png"></p>
<p>So that is how we create objects in C++.</p>
<h2>Why and When to use OOP</h2>
<p>Object-Oriented programming makes quite complex problems and applications structure and scale quite easily and enhances the readability of the codebase. We can use OOP to create the applications for bigger and real-life applications, it allows us to add new features quite easily and thereby improving the maintaining ability.</p>
<p>The following might be the reasons to use OOP in an application:</p>
<ul>
<li>
<p>If the application cannot be stated in a single sentence, you need to create the components into classes and hence use OOP.</p>
</li>
<li>
<p>It is quite easy to maintain OOP applications even if there are a lot of maintainers to the project.</p>
</li>
<li>
<p>If you see certain parts of the code repeating in many places, OOP is the way to go.</p>
</li>
<li>
<p>If you want your application to be secure in terms of data from the rest of the components, OOP is again a great option as you can use Encapsulation to the advantage.</p>
</li>
</ul>
<h2>Conclusion</h2>
<p>So, we have studied the concept of Object-Oriented Programming and how to define classes and create objects in C++. There are a lot of concepts in OOP to grab up in the next few parts of the series, but this should build the foundation of OOP as a paradigm or a way of thinking when it comes to problem-solving and programming.</p>
<p>I hope you have got a good introduction to the basics of Object-Oriented Programming and its semantics in C++. Thank you for reading. In the next part, we will explore the Access Modifiers/Specifiers in the Classes. Until then Happy Coding :)</p>
