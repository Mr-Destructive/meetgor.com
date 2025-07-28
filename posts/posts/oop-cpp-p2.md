{
  "title": "Object-Oriented Programming in C++: Access Modifiers",
  "status": "published",
  "slug": "oop-cpp-p2",
  "date": "2025-04-05T12:33:53Z"
}

<h2>Introduction</h2>
<p>Moving on to Part 2 of Object-Oriented Programming in C++, we will look into the concept of access modifiers which we skipped in the first section. This won't be a comprehensive guide on access modifiers as it requires some other concepts like <code>Inheritance</code> which we will introduce in the next few sections. This series will be like connecting the pieces of OOP together and building a good foundation.</p>
<p>In this part, we will discuss the access modifiers and their related concepts. The below is a gist of what will be covered in this part.</p>
<ul>
<li>Understanding Access Modifiers</li>
<li>Type of Access Modifiers
<ol>
<li>Private
<ul>
<li>Getters</li>
<li>Setters</li>
</ul>
</li>
<li>Public</li>
<li>Protected
<ul>
<li>Introduction to Derived and Friend classes</li>
</ul>
</li>
</ol>
</li>
</ul>
<p>Without wasting any time, let's roll in.</p>
<h2>What are Access Modifiers</h2>
<p>Access Modifiers as the name suggests they are used to change the accessibility of properties and methods in a Class. That means if we want a certain scope of our code to access the class members we can specify that using <code>access modifiers</code>. Access Modifiers form the basics or the pillar for <code>Data Encapsulation</code> in OOP.</p>
<h3>Data Encapsulation</h3>
<p>Data Encapsulation is terminology in OOP to keep all the components in a single entity or unit. Here the components are <code>properties</code> and <code>methods</code> and they are contained in a single unit called <code>Class</code>.  We need to carefully use the class members i.e. properties and methods in order to avoid security and limiting certain members to be available for a specific scope or block of code.</p>
<p>Hence, the Encapsulation of class members helps in structuring the class in a secure and allows access in a particular expected way.</p>
<h2>Types of Access Modifiers in C++</h2>
<p>We do have certain modifiers or limiters for accessing class members in C++. We can either keep the class members limited to the class itself (and friend class), keep them open to the global scope, or restrict them to certain classes only(derived and friend classes).</p>
<p>Let's discuss them one by one.</p>
<ul>
<li>
<h3>Private</h3>
</li>
</ul>
<p>This access modifier allows us to access the class members only to the class methods and the friend functions and classes. We will discuss what friend classes are in the next few parts of the series. This restricts the usage of the class members directly from the main function or other scopes in our program. You can use <code>private:</code> to indicate the below declared properties and methods are private to the class.</p>
<pre><code class="language-cpp">class Name
{ 
    private:
        int name;
        //other properties and methods
}
</code></pre>
<p>OR (don't specify anything it <strong>is private by default</strong>)</p>
<pre><code class="language-cpp">class Name
{ 
    int name;
    //other properties and methods
}
</code></pre>
<p>This is assigned to every member of the class by default. So that is why we explicitly told to make it public in an example in the previous part of the series. Let's see what happens if we do not make it public.</p>
<pre><code class="language-cpp">#include&lt;iostream&gt;
using namespace std;

class Animal{
	int legs;
	string name;
	void print()
    	{
		cout&lt;&lt;name&lt;&lt;&quot; has &quot;&lt;&lt;legs&lt;&lt;&quot; legs.
&quot;;
	}
};

int main()
{

	Animal dog;
	dog.name=&quot;Floyd&quot;;
	dog.legs=4;
	dog.print();
	
	return 0;
}

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631940306089/l4zOKB-RY.png" alt="image.png"></p>
<p>You can see we cannot directly use those properties and methods which are set as private. It is a convention to keep properties private and create public methods to access those from the rest of the program. The public methods used to access and modify the value of the private properties are called <code>getters</code> and <code>setters</code> respectively.</p>
<h4>Getters</h4>
<p>Getters are the public methods of a class that are used to access a value to the private properties of that class. It is a function that returns the data of that particular property. We access the function as normally as we access the public functions.</p>
<p>We define the getter function as:</p>
<pre><code class="language-cpp">datatype getpropertyname()
{
    return propertyname;
}
</code></pre>
<p>We can access the getter function as a normal public function, but this function returns a value, so we can store it in a variable and do all sorts of things.</p>
<pre><code class="language-cpp">classname objname;
court&lt;&lt;objname.getpropertyname()&lt;&lt;endl;
</code></pre>
<h4>Setters</h4>
<p>Setters are the public methods of a class that are used to assign/modify the value of the private properties of that class. It is a function that simply assigns the private properties to data of the data which is passed in as an argument to the function. We access the function as normally as we access the public functions. We have to pass in the value to assign the property to the function.</p>
<p>We define the setter function as:</p>
<pre><code class="language-cpp">datatype setpropertyname(datatype x)
{
    propertyname = x;
}
</code></pre>
<p>We can access the setter function as a normal public function, but this function takes in a parameter, so need to pass in the value to assign it the same as the correspondent data type of that property.</p>
<pre><code class="language-cpp">classname objname;
objname.setpropertyname(data);
</code></pre>
<p>So, after applying the getter and setter concept to our example, we can use them and make the properties private without any issues.</p>
<pre><code class="language-cpp">#include&lt;iostream&gt;
using namespace std;

class Animal{
	int legs;
	string name;
	public:

    // take in a parameter of type same as of the property 
	void setName(string dogname)   
	{
		name=dogname; 
        // assign the property a value same as of the parameter       
	}

    // return type should be same as of the property 
	string getName()   
	{
		return name;
	}

	void setleg(int noflegs)
	{
		legs=noflegs;
	}

	int getleg()
	{
		return legs;
	}
};

int main()
{
	Animal dog;
	dog.setName(&quot;Flyod&quot;);
	dog.setleg(4);
	cout&lt;&lt;dog.getName()&lt;&lt;endl;
	cout&lt;&lt;dog.getleg()&lt;&lt;endl;
	
	return 0;
}

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631949135706/zVd5cMPKa.png" alt="image.png"></p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631956955383/MKp5YDMv5.png" alt="image.png"></p>
<p>The above code runs successfully, and hence we make our program more secure and provide limited access to the rest of the components. You can see how we have passed the data as the parameter to the setter function, it should be the appropriate data-type with the property you are trying to set and also the same return type for the getter function.</p>
<p>The code is also much readable and structured for others to read and understand. It might seem a silly thing but it really improves the maintainability in the longer run.</p>
<ul>
<li>
<h3>Public</h3>
</li>
</ul>
<p>This might be familiar till now, we have been using this access modifier till now and it is quite straightforward to understand. We have created the getter and setter function using this access modifier.</p>
<p>So, the <code>public</code> access modifier makes the class members be accessible anywhere in the program. This might be OK for many small applications but it is not ideal to use them for production-level applications as it might cause undesired consequences i.e. BUGS.</p>
<p>We need to explicitly write public in the class definition, as <code>private</code> is set by default. So, as a simple example, as we saw in the previous part, it can be a lot easier to go with the public to understand OOP but the main OOP really shines in aspects like Encapsulation and Inheritance of Classes, which we will surely see in the upcoming parts.</p>
<p>This is the example from the previous part:</p>
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

int main()
{
	Animal dog;            // create an object of class Animal

	dog.name=&quot;Floyd&quot;;     // assign the class property to actual data in memory
	dog.legs=4;                 
	dog.print();          // call a method(function) associated to object's class
	
	return 0;
}

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631795000896/kUvcfVU7Y.png" alt="public-class"></p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631958067712/yU8uvd5KQ.png" alt="image.png"></p>
<p>As explained earlier, it is accessible to the main function or other scopes as well. We can access them using the dot separator(<code>.</code>) to assign it or to call the method.</p>
<ul>
<li>
<h3>Protected</h3>
</li>
</ul>
<p>This is quite a handy access modifier, as it allows us to have the flexibility to keep the class members private and accessible to its derived or friend class. We will see the derived and friend classes in the next few parts. But for now, let's understand in an abstracted way.</p>
<h4>Derived Class(Child Class/ Sub Class)</h4>
<p>This is a concept in Inheritance, i.e. to pass the members of a class to another. So, there has to be two classes, the main(parent class) and another class that will inherit or take in the members from the parent class. So, the derived class has access to its public and protected members only.</p>
<h4>Friend class</h4>
<p>A friend class is a class that is allowed to access its <strong>private</strong> and protected properties or methods. It is kind of a special tag assigned to a class that it can access certain class' members. We will see it in detail afterward, right now, it's enough to understand, friend class is a class that can access a particular class' members may it be private or protected.</p>
<p>The difference between a friend and a derived class is that a friend class can access the private members of the class to which it is a friend, but a derived class can't. Friend class also can't be inherited. Again, we will see this is in detail.</p>
<p>So, it doesn't make sense for me to explain protected here. But just assume friend classes and derived classes are a thing.</p>
<pre><code class="language-cpp">#include&lt;iostream&gt;
using namespace std;

class Animal
{
	public:
		int legs;
		string type;
		
		void print()
		{
			cout&lt;&lt;type&lt;&lt;&quot; has &quot;&lt;&lt;legs&lt;&lt;&quot; legs.
&quot;;
		}

	protected:
		string name;
};


int main()
{

	Animal dog;
	dog.type=&quot;dog&quot;;
	dog.legs=4;
    	dog.name=&quot;Floyd&quot;;
	dog.print();
	
	return 0;
}

</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631956345308/epOckpl6X.png" alt="image.png"></p>
<p>This gives an error as protected members can be accessed only by derived or friend classes since we are accessing it from the main function, it's not allowed in the global scope. We can only access the protected members from the derived or friend classes.</p>
<p>So, let's derive a class from the base class(Animal), and after that, we can see protected members in action.</p>
<pre><code class="language-cpp">#include&lt;iostream&gt;
using namespace std;

class Animal
{
	public:
		int legs;
		string type;
		
		void print()
		{
			cout&lt;&lt;type&lt;&lt;&quot; has &quot;&lt;&lt;legs&lt;&lt;&quot; legs.
&quot;;
		}

	protected:
		string name;
};

class Pet:public Animal 
{
	public:
		void print()
		{
			name=&quot;Floyd&quot;;
			cout&lt;&lt;name&lt;&lt;&quot; is a &quot;&lt;&lt;type&lt;&lt;endl;
			cout&lt;&lt;type&lt;&lt;&quot; has &quot;&lt;&lt;legs&lt;&lt;&quot; legs.
&quot;;
		}

};

int main()
{

	Pet dog;
	dog.type=&quot;dog&quot;;
	dog.legs=4;
	dog.print();
	
	return 0;
}


</code></pre>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631956708698/HJTOZN9sw.png" alt="image.png"></p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631956815750/a2y5QLiPT.png" alt="image.png"></p>
<p>We have accessed the protected property <code>name</code> in the derived class <code>Pet</code>. Yes, this is how we derive a class in C++,</p>
<pre><code class="language-cpp">// A child class serived from base class
class NewClassName: public BaseClassName
{
     // properties and methods
};
</code></pre>
<p>So, after deriving the <code>Pet</code> class from <code>Animal</code> class, we have access to its public and protected class as well. We simply assigned one of the protected members from its base class i.e. <code>name</code> and used it in the function <code>print</code>.</p>
<p>This is how you can use protected members in a derived class, also you can do it in friend class, but we will look at that in a separate part. Remember you can <strong>use getters and setters as well to assign and access those protected members</strong>.  Keeping it simple and easy to understand is what a protected access modifier can provide.</p>
<p><strong>Protected access modifiers are quite great and provide much more flexibility than <code>private</code> and more privacy than <code>public</code> access modifiers.</strong></p>
<h2>Which to use when?</h2>
<p>It is often a good practice to use private members, but it might be not possible to make everything private, so we can use certain properties as private, some methods as public and protected as well as per requirement and complexity.</p>
<p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1631953460021/_4WlLGDI_.png" alt="acmod.png"></p>
<p>Looking at the above chart, you can analyze your problem and work a way around to find the perfect secure match for your application.</p>
<p>Remember to use public members only when you have to explicitly use through the program. If you find that, this member shouldn't be introduced in certain scope then use make it private or protected.</p>
<p>If there is some kind of hierarchy in the program and its components, analyze and make a structure of it to have a better understanding of the program flow.</p>
<h2>Conclusion</h2>
<p>So, from this part, we were able to understand the access modifiers and got somewhat of a dive into Inheritance. There were some concepts like Friend and Derived classes, Inheritance, Encapsulation which were just explained in short, but they are a topic that deserves separate attention. We'll discuss them in the upcoming parts.</p>
<p>Thanks for reading. Hope you understood the concepts in a better way. Until then, as always, Happy Coding :)</p>
